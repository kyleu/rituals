package workspace

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullSprint struct {
	Sprint      *sprint.Sprint                `json:"sprint"`
	Histories   shistory.SprintHistories      `json:"histories,omitempty"`
	Members     smember.SprintMembers         `json:"members,omitempty"`
	Self        *smember.SprintMember         `json:"self,omitempty"`
	Permissions spermission.SprintPermissions `json:"permissions,omitempty"`
	Team        *team.Team                    `json:"team,omitempty"`
	Estimates   estimate.Estimates            `json:"estimates,omitempty"`
	Standups    standup.Standups              `json:"standups,omitempty"`
	Retros      retro.Retros                  `json:"retros,omitempty"`
	Comments    comment.Comments              `json:"comments,omitempty"`
	Actions     action.Actions                `json:"actions,omitempty"`
}

func (s *Service) LoadSprint(
	ctx context.Context, slug string, userID uuid.UUID, username string, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger,
) (*FullSprint, error) {
	spr, err := s.s.GetBySlug(ctx, tx, slug, logger)
	if err != nil {
		if hist, _ := s.sh.Get(ctx, tx, slug, logger); hist != nil {
			spr, err = s.s.Get(ctx, tx, hist.SprintID, logger)
			if err != nil {
				return nil, errors.Errorf("no sprint found with slug [%s]", slug)
			}
		}
	}
	if spr == nil {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no sprint found with slug [%s]", slug)
		}
		spr, err = s.s.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no sprint found with id [%s]", slug)
		}
	}
	ret := &FullSprint{Sprint: spr}

	ret.Histories, err = s.sh.GetBySprintID(ctx, tx, spr.ID, params.Get("shistory", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Members, ret.Self, err = s.membersSprint(ctx, tx, spr.ID, userID, username, params.Get("smember", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Permissions, err = s.sp.GetBySprintID(ctx, tx, spr.ID, params.Get("spermission", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	if spr.TeamID != nil {
		ret.Team, err = s.t.Get(ctx, tx, *spr.TeamID, logger)
		if err != nil {
			return nil, err
		}
	}

	ret.Estimates, err = s.e.GetBySprintID(ctx, tx, &spr.ID, params.Get(util.KeyEstimate, nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Standups, err = s.u.GetBySprintID(ctx, tx, &spr.ID, params.Get(util.KeyStandup, nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Retros, err = s.r.GetBySprintID(ctx, tx, &spr.ID, params.Get(util.KeyRetro, nil, logger), logger)
	if err != nil {
		return nil, err
	}

	args := make([]any, 0, (len(ret.Estimates)*2)+(len(ret.Standups)*2)+(len(ret.Retros)*2)+2)
	args = append(args, util.KeySprint, spr.ID)
	for _, x := range ret.Estimates {
		args = append(args, util.KeyEstimate, x.ID)
	}
	for _, x := range ret.Standups {
		args = append(args, util.KeyStandup, x.ID)
	}
	for _, x := range ret.Retros {
		args = append(args, util.KeyRetro, x.ID)
	}

	ret.Comments, err = s.c.GetByModels(ctx, tx, logger, args...)
	if err != nil {
		return nil, err
	}
	ret.Actions, err = s.a.GetByModels(ctx, tx, logger, enum.ModelServiceSprint, ret.Sprint.ID)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Service) membersSprint(
	ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, userID uuid.UUID, username string, params *filter.Params, logger util.Logger,
) (smember.SprintMembers, *smember.SprintMember, error) {
	members, err := s.sm.GetBySprintID(ctx, tx, sprintID, params, logger)
	if err != nil {
		return nil, nil, err
	}
	self := members.Get(sprintID, userID)
	if self == nil && username != "" {
		err = s.us.CreateIfNeeded(ctx, userID, username, tx, logger)
		if err != nil {
			return nil, nil, err
		}
		_, err = s.sm.Register(ctx, sprintID, userID, username, enum.MemberStatusMember, nil, s.a, s.send, logger)
		if err != nil {
			return nil, nil, err
		}
		members, err = s.sm.GetBySprintID(ctx, tx, sprintID, params, logger)
		if err != nil {
			return nil, nil, err
		}
		self = members.Get(sprintID, userID)
	}
	return members, self, nil
}
