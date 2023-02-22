package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionSprint(p *Params) (*FullSprint, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.Profile, p.Accounts, nil, nil, p.Logger)
	fs, err := p.Svc.LoadSprint(lp, func() (team.Teams, error) {
		return p.Svc.t.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	})
	if err != nil {
		return nil, "", "", err
	}

	switch p.Act {
	case action.ActUpdate:
		return sprintUpdate(p, fs)
	case action.ActMemberUpdate:
		return sprintMemberUpdate(p, fs)
	case action.ActMemberRemove:
		return sprintMemberRemove(p, fs)
	case action.ActMemberSelf:
		return sprintUpdateSelf(p, fs)
	case action.ActComment:
		return sprintComment(p, fs)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

func sprintUpdate(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	tgt := fs.Sprint.Clone()
	tgt.Title = p.Frm.GetStringOpt("title")
	if tgt.Title == "" {
		return nil, "", "", errors.New("title may not be empty")
	}
	tgt.Slug = p.Frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tx, err := p.Svc.db.StartTransaction(p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	tgt.Slug = p.Svc.s.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.sh, tx, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.StartDate, _ = p.Frm.GetTime("startDate", false)
	tgt.StartDate = util.TimeTruncate(tgt.StartDate)
	tgt.EndDate, _ = p.Frm.GetTime("endDate", false)
	tgt.EndDate = util.TimeTruncate(tgt.EndDate)
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	perms := loadPermissionsForm(p.Frm)
	modelChanged := len(fs.Sprint.Diff(tgt)) > 0
	permsChanged := len(fs.Permissions.ToPermissions().Diff(perms)) > 0
	if !modelChanged && !permsChanged {
		return fs, MsgNoChangesNeeded, fs.Sprint.PublicWebPath(), nil
	}
	if modelChanged {
		model, err := p.Svc.SaveSprint(p.Ctx, tgt, fs.Self.UserID, tx, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = updateTeam(
			util.KeySprint, fs.Sprint.TeamID, model.TeamID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fs.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		fs.Sprint = model
		err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActUpdate, model, &fs.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = sendTeamSprintUpdates(util.KeySprint, model.TeamID, nil, model, &fs.Self.UserID, p.Svc, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if err := p.Svc.sp.DeleteWhere(p.Ctx, tx, "sprint_id = $1", len(fs.Permissions), p.Logger, tgt.ID); err != nil {
			return nil, "", "", err
		}
		newPerms := make(spermission.SprintPermissions, 0, len(perms))
		for _, x := range perms {
			newPerms = append(newPerms, &spermission.SprintPermission{SprintID: tgt.ID, Key: x.Key, Value: x.Value})
		}
		if err = p.Svc.sp.Save(p.Ctx, tx, p.Logger, newPerms...); err != nil {
			return nil, "", "", err
		}
		err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActPermissions, perms, &fs.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Sprint updated", fs.Sprint.PublicWebPath(), nil
}

func sprintMemberUpdate(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	if fs.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this sprint")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := p.Frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fs.Members.Get(fs.Sprint.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this sprint", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := p.Svc.sm.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberUpdate, curr, &fs.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, MsgMemberUpdated, fs.Sprint.PublicWebPath(), nil
}

func sprintMemberRemove(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	if fs.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this sprint")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := fs.Members.Get(fs.Sprint.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this sprint", userID.String())
	}
	err := p.Svc.sm.Delete(p.Ctx, nil, curr.SprintID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberRemove, userID, &fs.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, MsgMemberRemoved, fs.Sprint.PublicWebPath(), nil
}

func sprintUpdateSelf(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == fs.Self.Name && picture == fs.Self.Picture {
		return fs, MsgNoChangesNeeded, fs.Sprint.PublicWebPath(), nil
	}

	fs.Self.Picture = picture
	fs.Self.Name = name
	err := p.Svc.sm.Update(p.Ctx, nil, fs.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.setName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": fs.Self.UserID, "name": name, "role": fs.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberUpdate, arg, &fs.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, MsgProfileEdited, fs.Sprint.PublicWebPath(), nil
}

func sprintComment(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	c, u, err := commentFromForm(p.Frm, fs.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceSprint:
		if c.ModelID != fs.Sprint.ID {
			return nil, "", "", errors.New("this comment refers to a different sprint")
		}
	case enum.ModelServiceEstimate:
		if curr := fs.Estimates.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an estimate that isn't part of this sprint")
		}
	case enum.ModelServiceStandup:
		if curr := fs.Standups.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an standup that isn't part of this sprint")
		}
	case enum.ModelServiceRetro:
		if curr := fs.Retros.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an retro that isn't part of this sprint")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = sendComment(enum.ModelServiceSprint, fs.Sprint.ID, c, &fs.Self.UserID, fs.Sprint.TeamID, nil, p.Svc.send, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, MsgCommentAdded, fs.Sprint.PublicWebPath() + u, nil
}
