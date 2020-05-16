package standup

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/action"

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
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcStandup.Key})

	return &Service{
		actions: actions,
		db:      db,
		Members: member.NewService(actions, db, util.SvcStandup.Key),
		logger:  logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcStandup.Key, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating standup slug"))
	}

	model := NewSession(title, slug, userID, sprintID)

	q := "insert into standup (id, slug, title, sprint_id, owner, status) values ($1, $2, $3, $4, $5, $6)"
	_, err = s.db.Exec(q, model.ID, slug, model.Title, model.SprintID, model.Owner, model.Status.String())
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new standup session"))
	}

	s.actions.Post(util.SvcStandup.Key, model.ID, userID, "create", nil, "")
	if model.SprintID != nil {
		s.actions.Post(util.SvcSprint.Key, model.ID, userID, "add-standup", nil, "")
	}
	return &model, nil
}

func (s *Service) List() ([]*Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from standup order by created desc")
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
	err := s.db.Get(dto, "select * from standup where id = $1", id)
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
	err := s.db.Get(dto, "select * from standup where slug = $1", slug)
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
	err := s.db.Select(&dtos, "select * from standup where owner = $1 order by created desc")
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
	q := "select x.* from standup x join standup_member m on x.id = m.standup_id where m.user_id = $1 order by m.created desc"
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
	q := "select * from standup where sprint_id = $1 order by created desc"
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

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, userID uuid.UUID) error {
	q := "update standup set title = $1 where id = $2"
	_, err := s.db.Exec(q, title, sessionID)
	s.actions.Post(util.SvcStandup.Key, sessionID, userID, "update", nil, "")
	return errors.WithStack(errors.Wrap(err, "error updating standup session"))
}
