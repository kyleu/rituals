package estimate

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/action"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions   *action.Service
	db        *sqlx.DB
	Members   *member.Service
	logger    logur.Logger
}

func NewService(actions *action.Service, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcEstimate})

	return &Service{
		actions: actions,
		db:      db,
		Members: member.NewService(actions, db, util.SvcEstimate),
		logger:  logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcEstimate, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating estimate slug"))
	}

	e := NewSession(title, slug, userID, sprintID)

	q := "insert into estimate (id, slug, title, owner, status, choices, options) values ($1, $2, $3, $4, $5, $6, $7)"
	choiceString := "{" + strings.Join(e.Choices, ",") + "}"
	_, err = s.db.Exec(q, e.ID, slug, e.Title, e.Owner, e.Status.String(), choiceString, e.Options.ToJSON())
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new estimate session"))
	}

	s.actions.Post(util.SvcEstimate, e.ID, userID, "create", nil, "")
	return &e, nil
}

func (s *Service) List() ([]*Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from estimate order by created desc")
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
	err := s.db.Get(dto, "select * from estimate where id = $1", id)
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
	err := s.db.Get(dto, "select * from estimate where slug = $1", slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetByOwner(id uuid.UUID) ([]*Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from estimate where owner = $1 order by created desc", id)
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
	q := "select x.* from estimate x join estimate_member m on x.id = m.estimate_id where m.user_id = $1 order by m.created desc"
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

func (s *Service) GetBySprint(sprintID uuid.UUID, limit int) ([]*Session, error) {
	var dtos []sessionDTO
	q := "select * from estimate where sprint_id = $1 order by created desc"
	if limit > 0 {
		q += fmt.Sprint(" limit ", limit)
	}
	err := s.db.Select(&dtos, q, sprintID)
	if err != nil {
		return nil, err
	}
	ret := make([]*Session, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, choices []string, userID uuid.UUID) error {
	q := "update estimate set title = $1, choices = $2 where id = $3"
	choiceString := "{" + strings.Join(choices, ",") + "}"
	_, err := s.db.Exec(q, title, choiceString, sessionID)
	s.actions.Post(util.SvcEstimate, sessionID, userID, "update", nil, "")
	return errors.WithStack(errors.Wrap(err, "error updating estimate session"))
}
