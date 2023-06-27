package workspace

import (
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionRetro(p *Params) (*FullRetro, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.Profile, p.Accounts, nil, nil, p.Logger)
	fr, err := p.Svc.LoadRetro(lp, func() (team.Teams, error) {
		return p.Svc.t.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	}, func() (sprint.Sprints, error) {
		return p.Svc.s.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	})
	if err != nil {
		return nil, "", "", err
	}

	switch p.Act {
	case action.ActUpdate:
		return retroUpdate(p, fr)
	case action.ActChildAdd:
		return retroFeedbackAdd(p, fr)
	case action.ActChildUpdate:
		return retroFeedbackUpdate(p, fr)
	case action.ActChildRemove:
		return retroFeedbackRemove(p, fr)
	case action.ActMemberUpdate:
		return retroMemberUpdate(p, fr)
	case action.ActMemberRemove:
		return retroMemberRemove(p, fr)
	case action.ActMemberSelf:
		return retroUpdateSelf(p, fr)
	case action.ActComment:
		return retroComment(p, fr)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

//nolint:gocognit
func retroUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if !fr.Admin() {
		return nil, "", "", errors.New("you do not have permission to update this retro")
	}
	tgt := fr.Retro.Clone()
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
	tgt.Slug = p.Svc.r.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.rh, tx, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.Categories = util.StringSplitAndTrim(p.Frm.GetStringOpt("categories"), ",")
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = p.Frm.GetUUID(util.KeySprint, true)
	perms := loadPermissionsForm(p.Frm)
	modelChanged := len(fr.Retro.Diff(tgt)) > 0
	permsChanged := len(fr.Permissions.ToPermissions().Diff(perms)) > 0
	if !modelChanged && !permsChanged {
		return fr, MsgNoChangesNeeded, fr.Retro.PublicWebPath(), nil
	}
	if len(fr.Retro.Diff(tgt)) == 0 && len(fr.Permissions.ToPermissions().Diff(perms)) == 0 {
		return fr, MsgNoChangesNeeded, fr.Retro.PublicWebPath(), nil
	}
	if modelChanged {
		model, e := p.Svc.SaveRetro(p.Ctx, tgt, fr.Self.UserID, tx, p.Logger)
		if e != nil {
			return nil, "", "", e
		}
		err = updateTeam(
			util.KeyRetro, fr.Retro.TeamID, model.TeamID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fr.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		err = updateSprint(
			util.KeyRetro, fr.Retro.SprintID, model.SprintID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fr.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		fr.Retro = model
		err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActUpdate, model, &fr.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = sendTeamSprintUpdates(util.KeyRetro, model.TeamID, model.SprintID, model, &fr.Self.UserID, p.Svc, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if e := p.Svc.rp.DeleteWhere(p.Ctx, tx, "retro_id = $1", len(fr.Permissions), p.Logger, tgt.ID); e != nil {
			return nil, "", "", e
		}
		newPerms := lo.Map(perms, func(x *util.Permission, _ int) *rpermission.RetroPermission {
			return &rpermission.RetroPermission{RetroID: tgt.ID, Key: x.Key, Value: x.Value}
		})
		if e := p.Svc.rp.Save(p.Ctx, tx, p.Logger, newPerms...); e != nil {
			return nil, "", "", e
		}
		err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActPermissions, perms, &fr.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Retro updated", fr.Retro.PublicWebPath(), nil
}

func retroComment(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	c, u, err := commentFromForm(p.Frm, fr.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceRetro:
		if c.ModelID != fr.Retro.ID {
			return nil, "", "", errors.New("this comment refers to a different retro")
		}
	case enum.ModelServiceFeedback:
		if curr := fr.Feedbacks.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to a feedback that isn't part of this retro")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = sendComment(enum.ModelServiceRetro, fr.Retro.ID, c, &fr.Self.UserID, fr.Retro.TeamID, fr.Retro.SprintID, p.Svc.send, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, MsgCommentAdded, fr.Retro.PublicWebPath() + u, nil
}
