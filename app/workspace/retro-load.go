package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

var RetroDefaultCategories = []string{"good", "bad"}

type FullRetro struct {
	Retro       *retro.Retro                 `json:"retro"`
	Histories   rhistory.RetroHistories      `json:"histories,omitempty"`
	Members     rmember.RetroMembers         `json:"members,omitempty"`
	Self        *rmember.RetroMember         `json:"self,omitempty"`
	Permissions rpermission.RetroPermissions `json:"permissions,omitempty"`
	Team        *team.Team                   `json:"team,omitempty"`
	Sprint      *sprint.Sprint               `json:"sprint,omitempty"`
	Feedbacks   feedback.Feedbacks           `json:"feedbacks,omitempty"`
	Comments    comment.Comments             `json:"comments,omitempty"`
	Actions     action.Actions               `json:"actions,omitempty"`
}

func (s *Service) LoadRetro(ctx context.Context, slug string, user uuid.UUID, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger) (*FullRetro, error) {
	r, err := s.r.GetBySlug(ctx, tx, slug, logger)
	if err != nil {
		if hist, _ := s.rh.Get(ctx, tx, slug, logger); hist != nil {
			r, err = s.r.Get(ctx, tx, hist.RetroID, logger)
			if err != nil {
				return nil, errors.Errorf("no retro found with slug [%s]", slug)
			}
		}
	}
	if r == nil {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no retro found with slug [%s]", slug)
		}
		r, err = s.r.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no retro found with id [%s]", slug)
		}
	}
	ret := &FullRetro{Retro: r}

	ret.Histories, err = s.rh.GetByRetroID(ctx, tx, r.ID, params.Get("rhistory", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Members, err = s.rm.GetByRetroID(ctx, tx, r.ID, params.Get("rmember", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Self = ret.Members.Get(r.ID, user)
	ret.Permissions, err = s.rp.GetByRetroID(ctx, tx, r.ID, params.Get("rpermission", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	if r.TeamID != nil {
		ret.Team, err = s.t.Get(ctx, tx, *r.TeamID, logger)
		if err != nil {
			return nil, err
		}
	}
	if r.SprintID != nil {
		ret.Sprint, err = s.s.Get(ctx, tx, *r.SprintID, logger)
		if err != nil {
			return nil, err
		}
	}

	ret.Feedbacks, err = s.f.GetByRetroID(ctx, tx, r.ID, params.Get("feedback", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	args := make([]any, 0, (len(ret.Feedbacks)*2)+2)
	args = append(args, util.KeyRetro, r.ID)
	for _, f := range ret.Feedbacks {
		args = append(args, util.KeyFeedback, f.ID)
	}

	ret.Comments, err = s.c.GetByModels(ctx, tx, logger, args...)
	if err != nil {
		return nil, err
	}
	ret.Actions, err = s.a.GetByModels(ctx, tx, logger, enum.ModelServiceRetro, ret.Retro.ID)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
