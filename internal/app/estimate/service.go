package estimate

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/internal/app/member"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"logur.dev/logur"
	"strings"
)

type Service struct {
	db      *sqlx.DB
	Members member.Service
	logger  logur.Logger
}

func NewEstimateService(db *sqlx.DB, logger logur.Logger) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcEstimate})

	return Service{
		db:      db,
		Members: member.NewMemberService(db, util.SvcEstimate, "estimate_member", "estimate_id"),
		logger:  logger,
	}
}

func (s *Service) NewSession(title string, userID uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcEstimate, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating slug"))
	}

	e := NewSession(title, slug, userID)

	q := "insert into estimate (id, slug, title, owner, status, choices, options) values ($1, $2, $3, $4, $5, $6, $7)"
	choiceString := "{" + strings.Join(e.Choices, ",") + "}"
	_, err = s.db.Exec(q, e.ID, slug, e.Title, e.Owner, e.Status.String(), choiceString, e.Options.ToJSON())
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new session"))
	}
	return &e, nil
}

func (s *Service) List() ([]Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from estimate")
	if err != nil {
		return nil, err
	}
	ret := make([]Session, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	err := s.db.Get(dto, "select * from estimate where id = $1", id)
	if err != nil {
		return nil, err
	}
	ret := dto.ToSession()
	return &ret, nil
}

func (s *Service) GetByOwner(id uuid.UUID) ([]Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from estimate where owner = $1", id)
	if err != nil {
		return nil, err
	}
	ret := make([]Session, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) GetBySlug(slug string) (*Session, error) {
	var dto = &sessionDTO{}
	err := s.db.Get(dto, "select * from estimate where slug = $1", slug)
	if err != nil {
		return nil, err
	}
	ret := dto.ToSession()
	return &ret, nil
}

func (s *Service) GetByMember(userID uuid.UUID) ([]Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select x.* from estimate x join estimate_member m on x.id = m.estimate_id where m.user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	ret := make([]Session, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, choices []string) error {
	q := "update estimate set title = $1, choices = $2 where id = $3"
	choiceString := "{" + strings.Join(choices, ",") + "}"
	_, err := s.db.Exec(q, title, choiceString, sessionID)
	return errors.WithStack(errors.Wrap(err, "error updating session"))
}
