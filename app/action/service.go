package action

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	db     *sqlx.DB
	logger logur.Logger
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
	contentJSON, _ := json.Marshal(content)
	_, err := s.db.Exec(q, id, svc, modelID, authorID, act, string(contentJSON), note)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new ["+svc+"] action"))
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

func (s *Service) List(params *query.Params) ([]*Action, error) {
	params = query.ParamsWithDefaultOrdering("action", params, &query.Ordering{Column: "occurred", Asc: false})
	var dtos []actionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", "action", "", params.OrderByString(), params.Limit, params.Offset))
	if err != nil {
		return nil, err
	}
	return toActions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Action, error) {
	dto := &actionDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", "action", "id = $1", "", 0, 0), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToAction(), nil
}

func (s *Service) GetByAuthor(id uuid.UUID, params *query.Params) ([]*Action, error) {
	params = query.ParamsWithDefaultOrdering("action", params, &query.Ordering{Column: "occurred", Asc: false})
	var dtos []actionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", "action", "author_id = $1", params.OrderByString(), params.Limit, params.Offset), id)
	if err != nil {
		return nil, err
	}
	return toActions(dtos), nil
}

func (s *Service) GetBySvcModel(svc string, modelID uuid.UUID, params *query.Params) ([]*Action, error) {
	params = query.ParamsWithDefaultOrdering("action", params, &query.Ordering{Column: "occurred", Asc: false})
	var dtos []actionDTO
	q := query.SQLSelect("*", "action", "svc = $1 and model_id = $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, svc, modelID)
	if err != nil {
		return nil, err
	}
	return toActions(dtos), nil
}

func toActions(dtos []actionDTO) []*Action {
	ret := make([]*Action, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToAction())
	}
	return ret
}
