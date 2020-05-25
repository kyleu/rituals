package retro

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
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.SvcRetro.Key})

	return &Service{
		actions:     actions,
		db:          db,
		Members:     member.NewService(actions, db, logger, util.SvcRetro.Key),
		Permissions: permission.NewService(actions, db, logger, util.SvcRetro.Key),
		logger:      logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, teamID *uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcRetro.Key, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating retro slug")
	}

	model := NewSession(title, slug, userID, teamID, sprintID)

	q := "insert into retro (id, slug, title, team_id, sprint_id, owner, status, categories) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	categoriesString := "{" + strings.Join(model.Categories, ",") + "}"
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.TeamID, model.SprintID, model.Owner, model.Status.String(), categoriesString)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new retro session")
	}

	s.Members.Register(model.ID, userID)

	s.actions.Post(util.SvcRetro.Key, model.ID, userID, action.ActCreate, nil, "")
	s.actions.PostRef(util.SvcSprint.Key, model.SprintID, util.SvcRetro.Key, model.ID, userID, action.ActContentAdd, "")
	s.actions.PostRef(util.SvcTeam.Key, model.TeamID, util.SvcRetro.Key, model.ID, userID, action.ActContentAdd, "")

	return &model, nil
}

func (s *Service) List(params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcRetro.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	q := query.SQLSelect("*", util.SvcRetro.Key, "id = $1", "", 0, 0)
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
	q := query.SQLSelect("*", util.SvcRetro.Key, "slug = $1", "", 0, 0)
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
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcRetro.Key, "owner = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultMCreatedOrdering...)
	var dtos []sessionDTO
	t := "retro join retro_member m on id = m.retro_id"
	q := query.SQLSelect("retro.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcRetro.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, teamID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetBySprint(sprintID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcRetro.Key, "sprint_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, sprintID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, categories []string, teamID *uuid.UUID, sprintID *uuid.UUID, userID uuid.UUID) error {
	q := "update retro set title = $1, categories = $2, team_id = $3, sprint_id = $4 where id = $5"
	categoriesString := "{" + strings.Join(categories, ",") + "}"
	err := s.db.UpdateOne(q, nil, title, categoriesString, teamID, sprintID, sessionID)
	s.actions.Post(util.SvcRetro.Key, sessionID, userID, action.ActUpdate, nil, "")
	return errors.Wrap(err, "error updating retro session")
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
