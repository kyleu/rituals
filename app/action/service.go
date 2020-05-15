package action

import (
	"database/sql"
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/util"
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

func (s *Service) New(svc string, modelID uuid.UUID, authorID uuid.UUID, act string, content interface{}, note string) (*Action, error) {
	id := util.UUID()
	q := "insert into action (id, svc, model_id, author_id, act, content, note) values ($1, $2, $3, $4, $5, $6, $7)"
	contentJson, _ := json.Marshal(content)
	_, err := s.db.Exec(q, id, svc, modelID, authorID, act, string(contentJson), note)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new [" + svc + "] action"))
	}
	return s.GetByID(id)
}

func (s *Service) Post(svc string, modelID uuid.UUID, authorID uuid.UUID, act string, content interface{}, note string) {
	go func() {
		_, err := s.New(svc, modelID, authorID, act, content, note)
		if err != nil {
			s.logger.Warn(fmt.Sprintf("unable to save new action: %+v", err))
		}
	}()
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

func (s *Service) GetBySvcModel(svc string, modelID uuid.UUID) ([]*Action, error) {
	var dtos []actionDTO
	q := "select * from action where svc = $1 and model_id = $2 order by occurred desc"
	err := s.db.Select(&dtos, q, svc, modelID)
	if err != nil {
		return nil, err
	}
	ret := make([]*Action, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToAction())
	}
	return ret, nil
}
