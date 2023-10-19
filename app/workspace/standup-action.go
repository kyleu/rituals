package workspace

import (
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionStandup(p *Params) (*FullStandup, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.Profile, p.Accounts, nil, nil, p.Logger)
	fu, err := p.Svc.LoadStandup(lp, func() (team.Teams, error) {
		return p.Svc.t.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	}, func() (sprint.Sprints, error) {
		return p.Svc.s.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	})
	if err != nil {
		return nil, "", "", err
	}
	switch p.Act {
	case action.ActUpdate:
		return standupUpdate(p, fu)
	case action.ActChildAdd:
		return standupReportAdd(p, fu)
	case action.ActChildUpdate:
		return standupReportUpdate(p, fu)
	case action.ActChildRemove:
		return standupReportRemove(p, fu)
	case action.ActMemberUpdate:
		return standupMemberUpdate(p, fu)
	case action.ActMemberRemove:
		return standupMemberRemove(p, fu)
	case action.ActMemberSelf:
		return standupUpdateSelf(p, fu)
	case action.ActComment:
		return standupComment(p, fu)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

func standupUpdate(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	if !fu.Admin() {
		return nil, "", "", errors.New("you do not have permission to update this standup")
	}
	tgt := fu.Standup.Clone()
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
	tgt.Slug = p.Svc.u.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.uh, tx, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = p.Frm.GetUUID(util.KeySprint, true)
	perms := loadPermissionsForm(p.Frm)
	modelChanged := len(fu.Standup.Diff(tgt)) > 0
	permsChanged := len(fu.Permissions.ToPermissions().Diff(perms)) > 0
	if !modelChanged && !permsChanged {
		return fu, MsgNoChangesNeeded, fu.Standup.PublicWebPath(), nil
	}
	if modelChanged {
		model, e := p.Svc.SaveStandup(p.Ctx, tgt, fu.Self.UserID, tx, p.Logger)
		if e != nil {
			return nil, "", "", e
		}
		err = updateTeam(
			util.KeyStandup, fu.Standup.TeamID, model.TeamID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fu.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		err = updateSprint(
			util.KeyStandup, fu.Standup.SprintID, model.SprintID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fu.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		fu.Standup = model
		err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActUpdate, model, &fu.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = sendTeamSprintUpdates(enum.ModelServiceStandup, model.TeamID, model.SprintID, model, &fu.Self.UserID, p.Svc, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if e := p.Svc.up.DeleteWhere(p.Ctx, tx, "standup_id = $1", len(fu.Permissions), p.Logger, tgt.ID); e != nil {
			return nil, "", "", e
		}
		newPerms := lo.Map(perms, func(x *util.Permission, _ int) *upermission.StandupPermission {
			return &upermission.StandupPermission{StandupID: tgt.ID, Key: x.Key, Value: x.Value}
		})
		if e := p.Svc.up.Save(p.Ctx, tx, p.Logger, newPerms...); e != nil {
			return nil, "", "", e
		}
		err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActPermissions, perms, &fu.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, "", "", err
	}
	return fu, "Standup updated", fu.Standup.PublicWebPath(), nil
}

func standupComment(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	c, u, err := commentFromForm(p.Frm, fu.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceStandup:
		if c.ModelID != fu.Standup.ID {
			return nil, "", "", errors.New("this comment refers to a different standup")
		}
	case enum.ModelServiceReport:
		if curr := fu.Reports.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to a report that isn't part of this standup")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = sendComment(enum.ModelServiceStandup, fu.Standup.ID, c, &fu.Self.UserID, fu.Standup.TeamID, fu.Standup.SprintID, p.Svc.send, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fu, MsgCommentAdded, fu.Standup.PublicWebPath() + u, nil
}
