package sprint

import (
	"database/sql"
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/action"
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
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcRetro})

	return &Service{
		actions: actions,
		db:      db,
		Members: member.NewService(actions, db, util.SvcSprint),
		logger:  logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcSprint, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating sprint slug"))
	}

	e := NewSession(title, slug, userID, nil)

	q := "insert into sprint (id, slug, title, owner, end_date) values ($1, $2, $3, $4, $5)"
	_, err = s.db.Exec(q, e.ID, slug, e.Title, e.Owner, e.EndDate)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new sprint session"))
	}

	s.actions.Post(util.SvcSprint, e.ID, userID, "create", nil, "")
	return &e, nil
}

func (s *Service) List() ([]*Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from sprint order by created desc")
	if err != nil {
		return nil, err
	}
	ret := make([]*Session, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	err := s.db.Get(dto, "select * from sprint where id = $1", id)
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
	err := s.db.Get(dto, "select * from sprint where slug = $1", slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetByOwner() ([]*Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from sprint where owner = $1 order by created desc")
	if err != nil {
		return nil, err
	}
	ret := make([]*Session, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) GetByMember(userID uuid.UUID, limit int) ([]*Session, error) {
	var dtos []sessionDTO
	q := "select x.* from sprint x join sprint_member m on x.id = m.sprint_id where m.user_id = $1 order by m.created desc"
	if limit > 0 {
		q += fmt.Sprint(" limit ", limit)
	}
	err := s.db.Select(&dtos, q, userID)
	if err != nil {
		return nil, err
	}
	ret := make([]*Session, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, userID uuid.UUID) error {
	q := "update sprint set title = $1 where id = $2"
	_, err := s.db.Exec(q, title, sessionID)
	s.actions.Post(util.SvcSprint, sessionID, userID, "update", nil, "")
	return errors.WithStack(errors.Wrap(err, "error updating sprint session"))
}
