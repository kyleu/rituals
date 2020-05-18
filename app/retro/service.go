package retro

import (
	"database/sql"
	"strings"

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
	actions *action.Service
	db      *sqlx.DB
	Members *member.Service
	logger  logur.Logger
}

func NewService(actions *action.Service, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcRetro.Key})

	return &Service{
		actions: actions,
		db:      db,
		Members: member.NewService(actions, db, util.SvcRetro.Key),
		logger:  logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcRetro.Key, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating retro slug"))
	}

	model := NewSession(title, slug, userID, sprintID)

	q := "insert into retro (id, slug, title, sprint_id, owner, status, categories) values ($1, $2, $3, $4, $5, $6, $7)"
	categoriesString := "{" + strings.Join(model.Categories, ",") + "}"
	_, err = s.db.Exec(q, model.ID, slug, model.Title, model.SprintID, model.Owner, model.Status.String(), categoriesString)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new retro session"))
	}

	s.actions.Post(util.SvcRetro.Key, model.ID, userID, "create", nil, "")
	if model.SprintID != nil {
		s.actions.Post(util.SvcSprint.Key, model.ID, userID, "add-retro", nil, "")
	}
	return &model, nil
}

func (s *Service) List(params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering("retro", params, &query.Ordering{Column: "created", Asc: false})
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", "retro", "", params.OrderByString(), params.Limit, params.Offset))
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", "retro", "id = $1", "", 0, 0), id)
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
	err := s.db.Get(dto, query.SQLSelect("*", "retro", "slug = $1", "", 0, 0), slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetByOwner(userID uuid.UUID, params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering("retro", params, &query.Ordering{Column: "created", Asc: false})
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", "retro", "owner = $1", params.OrderByString(), params.Limit, params.Offset), userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering("retro", params, &query.Ordering{Column: "m.created", Asc: false})
	var dtos []sessionDTO
	q := query.SQLSelect("x.*", "retro x join retro_member m on x.id = m.retro_id", "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetBySprint(sprintID uuid.UUID, params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering("retro", params, &query.Ordering{Column: "created", Asc: false})
	var dtos []sessionDTO
	q := query.SQLSelect("*", "retro", "sprint_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, sprintID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, categories []string, userID uuid.UUID) error {
	q := "update retro set title = $1, categories = $2 where id = $3"
	categoriesString := "{" + strings.Join(categories, ",") + "}"
	_, err := s.db.Exec(q, title, categoriesString, sessionID)
	s.actions.Post(util.SvcRetro.Key, sessionID, userID, "update", nil, "")
	return errors.WithStack(errors.Wrap(err, "error updating retro session"))
}

func toSessions(dtos []sessionDTO) []*Session {
	ret := make([]*Session, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
