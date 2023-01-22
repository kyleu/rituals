package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionSprint(p *Params) (*FullSprint, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.UserID, "", nil, nil, p.Logger)
	fs, err := p.Svc.LoadSprint(lp)
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
	tgt.Slug = p.Frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tgt.Slug = p.Svc.r.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.rh, nil, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.StartDate, _ = p.Frm.GetTime("startDate", false)
	tgt.EndDate, _ = p.Frm.GetTime("endDate", false)
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	model, err := p.Svc.SaveSprint(p.Ctx, tgt, fs.Self.UserID, nil, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	fs.Sprint = model
	err = p.Svc.send(enum.ModelServiceSprint, fs.Team.ID, action.ActUpdate, model, &fs.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Sprint saved", model.PublicWebPath(), nil
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
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberUpdate, curr, &fs.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Member updated", fs.Sprint.PublicWebPath(), nil
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
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberRemove, userID, &fs.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Member removed", fs.Sprint.PublicWebPath(), nil
}

func sprintUpdateSelf(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	choice := p.Frm.GetStringOpt("choice")
	name := p.Frm.GetStringOpt("name")
	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	fs.Self.Name = name
	err := p.Svc.sm.Update(p.Ctx, nil, fs.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == "global" {
		return nil, "", "", errors.New("can't change global name yet")
	}
	arg := util.ValueMap{"userID": fs.Self.UserID, "name": name}
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberUpdate, arg, &fs.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Profile edited", fs.Sprint.PublicWebPath(), nil
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
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActComment, c, &fs.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Comment added", fs.Sprint.PublicWebPath() + u, nil
}
