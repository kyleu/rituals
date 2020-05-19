package sprint

import (
	"database/sql"
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/query"
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
		Members: member.NewService(actions, db, util.SvcSprint.Key),
		logger:  logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcSprint.Key, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating sprint slug"))
	}

	model := NewSession(title, slug, userID, nil, nil)

	q := "insert into sprint (id, slug, title, owner, start_date, end_date) values ($1, $2, $3, $4, $5, $6)"
	_, err = s.db.Exec(q, model.ID, slug, model.Title, model.Owner, model.StartDate, model.EndDate)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new sprint session"))
	}

	s.actions.Post(util.SvcSprint.Key, model.ID, userID, action.ActCreate, nil, "")
	return &model, nil
}

func (s *Service) List(params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, &query.Ordering{Column: "created", Asc: false})
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.SvcSprint.Key, "", params.OrderByString(), params.Limit, params.Offset))
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.SvcSprint.Key, "id = $1", "", 0, 0), id)
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
	err := s.db.Get(dto, query.SQLSelect("*", util.SvcSprint.Key, "slug = $1", "", 0, 0), slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetByOwner(userID uuid.UUID, params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, &query.Ordering{Column: "created", Asc: false})
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.SvcSprint.Key, "owner = $1", params.OrderByString(), params.Limit, params.Offset), userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, &query.Ordering{Column: "m.created", Asc: false})
	var dtos []sessionDTO
	q := query.SQLSelect("x.*", "sprint x join sprint_member m on x.id = m.sprint_id", "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, userID uuid.UUID) error {
	q := "update sprint set title = $1 where id = $2"
	_, err := s.db.Exec(q, title, sessionID)
	s.actions.Post(util.SvcSprint.Key, sessionID, userID, action.ActUpdate, nil, "")
	return errors.WithStack(errors.Wrap(err, "error updating sprint session"))
}

func (s *Service) AssignSprint(svc string, sessionID *uuid.UUID, userID uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	q := "update " + svc + " set sprint_id = $1 where id = $2"
	_, err := s.db.Exec(q, sprintID, sessionID)
	s.actions.Post(svc, *sessionID, userID, action.ActAssignSprint, sprintID, "")
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error updating sprint for " + svc + " session"))
	}
	if sprintID == nil {
		return nil, nil
	}
	return s.GetByID(*sprintID)
}

func (s *Service) GetByIDPointer(sprintID *uuid.UUID) *Session {
	if sprintID == nil {
		return nil
	}
	spr, _ := s.GetByID(*sprintID)
	return spr
}

func toSessions(dtos []sessionDTO) []*Session {
	ret := make([]*Session, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
