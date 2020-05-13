package retro

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/actions"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions *actions.Service
	db      *sqlx.DB
	Members *member.Service
	logger  logur.Logger
}

func NewService(actions *actions.Service, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcRetro})

	return &Service{
		actions: actions,
		db:      db,
		Members: member.NewService(db, util.SvcRetro, "retro_member", "retro_id"),
		logger:  logger,
	}
}

func (s *Service) NewSession(title string, userID uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcRetro, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating retro slug"))
	}

	e := NewSession(title, slug, userID)

	q := "insert into retro (id, slug, title, owner, status, categories) values ($1, $2, $3, $4, $5, $6)"
	categoriesString := "{" + strings.Join(e.Categories, ",") + "}"
	_, err = s.db.Exec(q, e.ID, slug, e.Title, e.Owner, e.Status.String(), categoriesString)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new retro session"))
	}
	return &e, nil
}

func (s *Service) List() ([]*Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from retro order by created desc")
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
	err := s.db.Get(dto, "select * from retro where id = $1", id)
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
	err := s.db.Get(dto, "select * from retro where slug = $1", slug)
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
	err := s.db.Select(&dtos, "select * from retro where owner = $1 order by created desc")
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
	q := "select x.* from retro x join retro_member m on x.id = m.retro_id where m.user_id = $1 order by m.created desc"
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

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, categories []string) error {
	q := "update retro set title = $1, categories = $2 where id = $3"
	categoriesString := "{" + strings.Join(categories, ",") + "}"
	_, err := s.db.Exec(q, title, categoriesString, sessionID)
	return errors.WithStack(errors.Wrap(err, "error updating retro session"))
}
