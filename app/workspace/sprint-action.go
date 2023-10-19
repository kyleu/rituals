package workspace

import (
	"github.com/pkg/errors"
	"github.com/samber/lo"

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
	if !fs.Admin() {
		return nil, "", "", errors.New("you do not have permission to update this sprint")
	}
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
		model, e := p.Svc.SaveSprint(p.Ctx, tgt, fs.Self.UserID, tx, p.Logger)
		if e != nil {
			return nil, "", "", e
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
		err = sendTeamSprintUpdates(enum.ModelServiceSprint, model.TeamID, nil, model, &fs.Self.UserID, p.Svc, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if e := p.Svc.sp.DeleteWhere(p.Ctx, tx, "sprint_id = $1", len(fs.Permissions), p.Logger, tgt.ID); e != nil {
			return nil, "", "", e
		}
		newPerms := lo.Map(perms, func(x *util.Permission, _ int) *spermission.SprintPermission {
			return &spermission.SprintPermission{SprintID: tgt.ID, Key: x.Key, Value: x.Value}
		})
		if e := p.Svc.sp.Save(p.Ctx, tx, p.Logger, newPerms...); e != nil {
			return nil, "", "", e
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

func sprintComment(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
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
