package estimate

import (
	"database/sql"
	"github.com/kyleu/rituals.dev/app/database"
	"strings"

	"github.com/kyleu/rituals.dev/app/permission"

	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions     *action.Service
	db          *database.Service
	Members     *member.Service
	Permissions *permission.Service
	logger      logur.Logger
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.SvcEstimate.Key})

	return &Service{
		actions:     actions,
		db:          db,
		Members:     member.NewService(actions, db, logger, util.SvcEstimate.Key),
		Permissions: permission.NewService(actions, db, logger, util.SvcEstimate.Key),
		logger:      logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, teamID *uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcEstimate.Key, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating estimate slug")
	}

	model := NewSession(title, slug, userID, teamID, sprintID)

	q := "insert into estimate (id, slug, title, team_id, sprint_id, owner, status, choices) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	choiceString := "{" + strings.Join(model.Choices, ",") + "}"
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.TeamID, model.SprintID, model.Owner, model.Status.String(), choiceString)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new estimate session")
	}

	s.Members.Register(model.ID, userID)

	s.actions.Post(util.SvcEstimate.Key, model.ID, userID, action.ActCreate, nil, "")
	s.actions.PostRef(util.SvcSprint.Key, model.SprintID, util.SvcEstimate.Key, model.ID, userID, action.ActContentAdd, "")
	s.actions.PostRef(util.SvcTeam.Key, model.TeamID, util.SvcEstimate.Key, model.ID, userID, action.ActContentAdd, "")
	return &model, nil
}

func (s *Service) List(params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcEstimate.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcEstimate.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	q := query.SQLSelect("*", util.SvcEstimate.Key, "id = $1", "", 0, 0)
	err := s.db.Get(dto, q, nil, id)
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
	q := query.SQLSelect("*", util.SvcEstimate.Key, "slug = $1", "", 0, 0)
	err := s.db.Get(dto, q, nil, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetByOwner(userID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcEstimate.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcEstimate.Key, "owner = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcEstimate.Key, params, query.DefaultMCreatedOrdering...)
	var dtos []sessionDTO
	t := "estimate join estimate_member m on id = m.estimate_id"
	q := query.SQLSelect("estimate.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcEstimate.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcEstimate.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, teamID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetBySprint(sprintID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcEstimate.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcEstimate.Key, "sprint_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, sprintID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, choices []string, teamID *uuid.UUID, sprintID *uuid.UUID, userID uuid.UUID) error {
	q := "update estimate set title = $1, choices = $2, team_id = $3, sprint_id = $4 where id = $5"
	choiceString := "{" + strings.Join(choices, ",") + "}"
	err := s.db.UpdateOne(q, nil, title, choiceString, teamID, sprintID, sessionID)
	s.actions.Post(util.SvcEstimate.Key, sessionID, userID, "update", nil, "")
	return errors.Wrap(err, "error updating estimate session")
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
