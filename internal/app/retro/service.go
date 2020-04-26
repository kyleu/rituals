package retro

import (
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/internal/app/member"
	"logur.dev/logur"
)

type Service struct {
	db      *sqlx.DB
	members member.Service
	logger  logur.Logger
}

func NewRetroService(db *sqlx.DB, logger logur.LoggerFacade) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "retro"})

	return Service{
		db:      db,
		members: member.NewMemberService(db, "retro_member", "retro_id"),
		logger:  logger,
	}
}

func (s *Service) List() ([]Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from retro")
	if err != nil {
		return nil, err
	}
	ret := make([]Session, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) GetByOwner() ([]Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from retro where owner = $1")
	if err != nil {
		return nil, err
	}
	ret := make([]Session, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret, nil
}

func (s *Service) GetByMember(userID uuid.UUID) ([]Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select x.* from retro x join retro_member m on x.id = m.retro_id where m.user_id = $1", userID)
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
	err := s.db.Get(dto, "select * from retro where id = $1", id)
	if err != nil {
		return nil, err
	}
	ret := dto.ToSession()
	return &ret, nil
}

func (s *Service) GetMembers(id uuid.UUID) ([]member.Entry, error) {
	return s.members.GetByModelID(id)
}
