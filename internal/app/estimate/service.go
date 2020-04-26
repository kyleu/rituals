package estimate

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/internal/app/member"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"logur.dev/logur"
)

type Service struct {
	db      *sqlx.DB
	members member.Service
	logger  logur.Logger
}

func NewEstimateService(db *sqlx.DB, logger logur.LoggerFacade) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "estimate"})

	return Service{
		db:      db,
		members: member.NewMemberService(db, "estimate_member", "estimate_id"),
		logger:  logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID) (*Session, error) {

	slug, err := s.members.NewSlugFor(title)
	e := NewSession(title, slug, userID)

  if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating slug"))
	}

	q := "insert into estimate (id, slug, title, owner, status, choices, options) values ($1, $2, $3, $4, $5, $6, $7)"
	_, err = s.db.Exec(q, e.ID, slug, e.Title, e.Owner, e.Status.String(), util.ArrayToString(e.Choices), e.Options.ToJson())
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

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	err := s.db.Get(dto, "select * from estimate where id = $1", id)
	if err != nil {
		return nil, err
	}
	ret := dto.ToSession()
	return &ret, nil
}

func (s *Service) GetPolls(id uuid.UUID) ([]Poll, error) {
	var dtos []pollDTO
	err := s.db.Select(&dtos, "select * from poll where estimate_id = $1 order by idx", id)
	if err != nil {
		return nil, err
	}
	ret := make([]Poll, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToPoll())
	}
	return ret, nil
}

func (s *Service) GetPollByID(id uuid.UUID) (*Poll, error) {
	dto := &pollDTO{}
	err := s.db.Get(dto, "select * from poll where id = $1", id)
	if err != nil {
		return nil, err
	}
	ret := dto.ToPoll()
	return &ret, nil
}

func (s *Service) GetPollVotes(id uuid.UUID) ([]Vote, error) {
	var dtos []voteDTO
	err := s.db.Select(&dtos, "select * from vote where poll_id = $1", id)
	if err != nil {
		return nil, err
	}
	ret := make([]Vote, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToVote())
	}
	return ret, nil
}

func (s *Service) GetMembers(id uuid.UUID) ([]member.Entry, error) {
	return s.members.GetByModelID(id)
}
