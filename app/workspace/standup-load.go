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
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullStandup struct {
	Standup     *standup.Standup               `json:"standup"`
	Histories   uhistory.StandupHistories      `json:"histories,omitempty"`
	Members     umember.StandupMembers         `json:"members,omitempty"`
	UtilMembers util.Members                   `json:"-"`
	Self        *umember.StandupMember         `json:"self,omitempty"`
	Permissions upermission.StandupPermissions `json:"permissions,omitempty"`
	Team        *team.Team                     `json:"team,omitempty"`
	Sprint      *sprint.Sprint                 `json:"sprint,omitempty"`
	Reports     report.Reports                 `json:"reports,omitempty"`
	Comments    comment.Comments               `json:"comments,omitempty"`
	Actions     action.Actions                 `json:"actions,omitempty"`
}

func (s *Service) LoadStandup(
	ctx context.Context, slug string, userID uuid.UUID, username string, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger,
) (*FullStandup, error) {
	u, err := s.u.GetBySlug(ctx, tx, slug, logger)
	if err != nil {
		if hist, _ := s.uh.Get(ctx, tx, slug, logger); hist != nil {
			u, err = s.u.Get(ctx, tx, hist.StandupID, logger)
			if err != nil {
				return nil, errors.Errorf("no standup found with slug [%s]", slug)
			}
		}
	}
	if u == nil {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no standup found with slug [%s]", slug)
		}
		u, err = s.u.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no standup found with id [%s]", slug)
		}
	}
	ret := &FullStandup{Standup: u}

	ret.Histories, err = s.uh.GetByStandupID(ctx, tx, u.ID, params.Get("uhistory", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Members, ret.Self, err = s.membersStandup(ctx, tx, u.ID, userID, username, params.Get("umember", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.UtilMembers = ret.Members.ToMembers()
	ret.Permissions, err = s.up.GetByStandupID(ctx, tx, u.ID, params.Get("upermission", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	if u.TeamID != nil {
		ret.Team, err = s.t.Get(ctx, tx, *u.TeamID, logger)
		if err != nil {
			return nil, err
		}
	}
	if u.SprintID != nil {
		ret.Sprint, err = s.s.Get(ctx, tx, *u.SprintID, logger)
		if err != nil {
			return nil, err
		}
	}

	ret.Reports, err = s.rt.GetByStandupID(ctx, tx, u.ID, params.Get("report", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	args := make([]any, 0, (len(ret.Reports)*2)+2)
	args = append(args, util.KeyStandup, u.ID)
	for _, rpt := range ret.Reports {
		args = append(args, util.KeyReport, rpt.ID)
	}

	ret.Comments, err = s.c.GetByModels(ctx, tx, logger, args...)
	if err != nil {
		return nil, err
	}
	ret.Actions, err = s.a.GetByModels(ctx, tx, logger, enum.ModelServiceStandup, ret.Standup.ID)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Service) membersStandup(
	ctx context.Context, tx *sqlx.Tx, standupID uuid.UUID, userID uuid.UUID, username string, params *filter.Params, logger util.Logger,
) (umember.StandupMembers, *umember.StandupMember, error) {
	members, err := s.um.GetByStandupID(ctx, tx, standupID, params, logger)
	if err != nil {
		return nil, nil, err
	}
	self := members.Get(standupID, userID)
	if self == nil && username != "" {
		err = s.us.CreateIfNeeded(ctx, userID, username, tx, logger)
		if err != nil {
			return nil, nil, err
		}
		_, err = s.um.Register(ctx, standupID, userID, username, enum.MemberStatusMember, nil, s.a, s.send, logger)
		if err != nil {
			return nil, nil, err
		}
		members, err = s.um.GetByStandupID(ctx, tx, standupID, params, logger)
		if err != nil {
			return nil, nil, err
		}
		self = members.Get(standupID, userID)
	}
	return members, self, nil
}
