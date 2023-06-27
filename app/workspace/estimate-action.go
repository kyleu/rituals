package workspace

import (
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionEstimate(p *Params) (*FullEstimate, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.Profile, p.Accounts, nil, nil, p.Logger)
	fe, err := s.LoadEstimate(lp, func() (team.Teams, error) {
		return p.Svc.t.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	}, func() (sprint.Sprints, error) {
		return p.Svc.s.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	})
	if err != nil {
		return nil, "", "", err
	}
	switch p.Act {
	case action.ActUpdate:
		return estimateUpdate(p, fe)
	case action.ActChildAdd:
		return estimateStoryAdd(p, fe)
	case action.ActChildUpdate:
		return estimateStoryUpdate(p, fe)
	case action.ActChildStatus:
		return estimateStoryStatus(p, fe)
	case action.ActVote:
		return estimateStoryVote(p, fe)
	case action.ActChildRemove:
		return estimateStoryRemove(p, fe)
	case action.ActMemberUpdate:
		return estimateMemberUpdate(p, fe)
	case action.ActMemberRemove:
		return estimateMemberRemove(p, fe)
	case action.ActMemberSelf:
		return estimateUpdateSelf(p, fe)
	case action.ActComment:
		return estimateComment(p, fe)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

func estimateUpdate(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	if !fe.Admin() {
		return nil, "", "", errors.New("you do not have permission to edit this estimate")
	}
	tgt := fe.Estimate.Clone()
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
	tgt.Slug = p.Svc.e.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.eh, tx, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.Choices = util.StringSplitAndTrim(p.Frm.GetStringOpt("choices"), ",")
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = p.Frm.GetUUID(util.KeySprint, true)
	perms := loadPermissionsForm(p.Frm)
	modelChanged := len(fe.Estimate.Diff(tgt)) > 0
	permsChanged := len(fe.Permissions.ToPermissions().Diff(perms)) > 0
	if !modelChanged && !permsChanged {
		return fe, MsgNoChangesNeeded, fe.Estimate.PublicWebPath(), nil
	}
	if modelChanged {
		model, e := p.Svc.SaveEstimate(p.Ctx, tgt, fe.Self.UserID, tx, p.Logger)
		if e != nil {
			return nil, "", "", e
		}
		err = updateTeam(
			util.KeyEstimate, fe.Estimate.TeamID, model.TeamID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fe.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		err = updateSprint(
			util.KeyEstimate, fe.Estimate.SprintID, model.SprintID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fe.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		fe.Estimate = model
		err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActUpdate, model, &fe.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = sendTeamSprintUpdates(enum.ModelServiceEstimate, model.TeamID, model.SprintID, model, &fe.Self.UserID, p.Svc, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if e := p.Svc.ep.DeleteWhere(p.Ctx, tx, "estimate_id = $1", len(fe.Permissions), p.Logger, tgt.ID); e != nil {
			return nil, "", "", e
		}
		newPerms := lo.Map(perms, func(x *util.Permission, _ int) *epermission.EstimatePermission {
			return &epermission.EstimatePermission{EstimateID: tgt.ID, Key: x.Key, Value: x.Value}
		})
		if e := p.Svc.ep.Save(p.Ctx, tx, p.Logger, newPerms...); e != nil {
			return nil, "", "", e
		}
		err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActPermissions, perms, &fe.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Estimate updated", fe.Estimate.PublicWebPath(), nil
}

func estimateComment(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	c, u, err := commentFromForm(p.Frm, fe.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceEstimate:
		if c.ModelID != fe.Estimate.ID {
			return nil, "", "", errors.New("this comment refers to a different estimate")
		}
	case enum.ModelServiceStory:
		if curr := fe.Stories.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to a story that isn't part of this estimate")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = sendComment(enum.ModelServiceEstimate, fe.Estimate.ID, c, &fe.Self.UserID, fe.Estimate.TeamID, fe.Estimate.SprintID, p.Svc.send, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, MsgCommentAdded, fe.Estimate.PublicWebPath() + u, nil
}
