package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) CreateTeam(
	ctx context.Context, id uuid.UUID, title string, user uuid.UUID, name string, logger util.Logger,
) (*team.Team, *tmember.TeamMember, error) {
	slug := s.e.Slugify(ctx, id, title, "", s.eh, nil, logger)
	model := &team.Team{ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, Owner: user, Created: time.Now()}
	err := s.t.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save team")
	}

	err = s.a.Post(ctx, util.KeyTeam, model.ID, user, action.ActCreate, util.ValueMap{"payload": model}, nil, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save team activity")
	}

	member, err := s.tm.Register(ctx, model.ID, user, name, enum.MemberStatusOwner, nil, s.a, s.send, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save team owner")
	}

	return model, member, nil
}

func (s *Service) SaveTeam(ctx context.Context, t *team.Team, user uuid.UUID, tx *sqlx.Tx, logger util.Logger) (*team.Team, error) {
	curr, err := s.t.Get(ctx, tx, t.ID, logger)
	if err != nil {
		return nil, err
	}
	if curr == nil {
		return nil, errors.Errorf("no existing team found with id [%s]", t.ID.String())
	}

	if curr.Slug != t.Slug {
		_ = s.th.DeleteWhere(ctx, tx, "team_id = $1 and slug = $2", -1, logger, t.ID, curr.Slug)
		hist := &thistory.TeamHistory{Slug: curr.Slug, TeamID: t.ID, TeamName: t.TitleString(), Created: time.Now()}
		err = s.th.Create(ctx, tx, logger, hist)
		if err != nil {
			return nil, err
		}
		curr.Slug = t.Slug
	}

	err = s.t.Update(ctx, tx, t, logger)
	if err != nil {
		return nil, err
	}

	act := action.NewAction(enum.ModelServiceTeam, t.ID, user, "update", util.ValueMap{}, "")
	err = s.a.Save(ctx, tx, logger, act)
	if err != nil {
		return nil, err
	}

	return curr, nil
}
