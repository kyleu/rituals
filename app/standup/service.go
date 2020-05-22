package standup

import (
	"database/sql"

	"github.com/kyleu/rituals.dev/app/permission"

	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions     *action.Service
	db          *sqlx.DB
	Members     *member.Service
	Permissions *permission.Service
	logger      logur.Logger
}

func NewService(actions *action.Service, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcStandup.Key})

	return &Service{
		actions:     actions,
		db:          db,
		Members:     member.NewService(actions, db, logger, util.SvcStandup.Key),
		Permissions: permission.NewService(actions, db, logger, util.SvcStandup.Key),
		logger:      logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, teamID *uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcStandup.Key, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating standup slug"))
	}

	model := NewSession(title, slug, userID, teamID, sprintID)

	q := "insert into standup (id, slug, title, team_id, sprint_id, owner, status) values ($1, $2, $3, $4, $5, $6, $7)"
	_, err = s.db.Exec(q, model.ID, slug, model.Title, model.TeamID, model.SprintID, model.Owner, model.Status.String())
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new standup session"))
	}

	s.Members.Register(model.ID, userID)

	s.actions.Post(util.SvcStandup.Key, model.ID, userID, action.ActCreate, nil, "")
	s.actions.PostRef(util.SvcSprint.Key, model.SprintID, util.SvcStandup.Key, model.ID, userID, action.ActContentAdd, "")
	s.actions.PostRef(util.SvcTeam.Key, model.TeamID, util.SvcStandup.Key, model.ID, userID, action.ActContentAdd, "")

	return &model, nil
}

func (s *Service) List(params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcStandup.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.SvcStandup.Key, "", params.OrderByString(), params.Limit, params.Offset))
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.SvcStandup.Key, "id = $1", "", 0, 0), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetBySlug(slug string) (*Session, error) {
	var dto = &sessionDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.SvcStandup.Key, "slug = $1", "", 0, 0), slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetByOwner(userID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcStandup.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.SvcStandup.Key, "owner = $1", params.OrderByString(), params.Limit, params.Offset), userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcStandup.Key, params, query.DefaultMCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("x.*", "standup x join standup_member m on x.id = m.standup_id", "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcStandup.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.SvcStandup.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset), teamID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetBySprint(sprintID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcStandup.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcStandup.Key, "sprint_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, sprintID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, teamID *uuid.UUID, sprintID *uuid.UUID, userID uuid.UUID) error {
	q := "update standup set title = $1, team_id = $2, sprint_id = $3 where id = $4"
	_, err := s.db.Exec(q, title, teamID, sprintID, sessionID)
	s.actions.Post(util.SvcStandup.Key, sessionID, userID, action.ActUpdate, nil, "")
	return errors.WithStack(errors.Wrap(err, "error updating standup session"))
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
