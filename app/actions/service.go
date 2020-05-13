package actions

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"logur.dev/logur"
)

type Service struct {
	db         *sqlx.DB
	logger     logur.Logger
}

func NewService(db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "action"})
	svc := Service{
		db:     db,
		logger: logger,
	}
	return &svc
}

func (s *Service) List() ([]*Action, error) {
	var dtos []actionDTO
	err := s.db.Select(&dtos, "select * from action order by occurred desc")
	if err != nil {
		return nil, err
	}
	ret := make([]*Action, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToAction())
	}
	return ret, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Action, error) {
	dto := &actionDTO{}
	err := s.db.Get(dto, "select * from action where id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToAction(), nil
}

func (s *Service) GetByAuthor(id uuid.UUID) ([]*Action, error) {
	var dtos []actionDTO
	err := s.db.Select(&dtos, "select * from action where author_id = $1 order by occurred desc", id)
	if err != nil {
		return nil, err
	}
	ret := make([]*Action, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToAction())
	}
	return ret, nil
}
