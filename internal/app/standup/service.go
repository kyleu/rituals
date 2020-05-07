package standup

import (
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/internal/app/member"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"logur.dev/logur"
)

type Service struct {
	db      *sqlx.DB
	Members member.Service
	logger  logur.Logger
}

func NewStandupService(db *sqlx.DB, logger logur.Logger) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcStandup})

	return Service{
		db:      db,
		Members: member.NewMemberService(db, util.SvcStandup, "standup_member", "standup_id"),
		logger:  logger,
	}
}

func (s *Service) List() ([]Session, error) {
	var dtos []sessionDTO
	err := s.db.Select(&dtos, "select * from standup")
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
	err := s.db.Select(&dtos, "select * from standup where owner = $1")
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
	err := s.db.Select(&dtos, "select x.* from standup x join standup_member m on x.id = m.standup_id where m.user_id = $1", userID)
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
	err := s.db.Get(dto, "select * from standup where id = $1", id)
	if err != nil {
		return nil, err
	}
	ret := dto.ToSession()
	return &ret, nil
}

func (s *Service) GetMembers(id uuid.UUID) ([]member.Entry, error) {
	return s.Members.GetByModelID(id)
}
