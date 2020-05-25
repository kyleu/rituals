package action

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/database"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	db     *database.Service
	logger logur.Logger
}

func NewService(db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.KeyAction})
	svc := Service{
		db:     db,
		logger: logger,
	}

	return &svc
}

func (s *Service) New(svc string, modelID uuid.UUID, authorID uuid.UUID, act string, content interface{}, note string) (*Action, error) {
	id := util.UUID()
	q := "insert into action (id, svc, model_id, author_id, act, content, note) values ($1, $2, $3, $4, $5, $6, $7)"
	err := s.db.Insert(q, nil, id, svc, modelID, authorID, act, util.ToJSON(content), note)

	if err != nil {
		return nil, errors.Wrap(err, "error saving new ["+svc+"] action")
	}

	return s.GetByID(id)
}

func (s *Service) PostRef(svc string, modelID *uuid.UUID, refSvc string, refID uuid.UUID, userID uuid.UUID, act string, note string) {
	if modelID != nil {
		actionContent := map[string]interface{}{util.KeySvc: refSvc, util.KeyID: refID}
		s.Post(svc, *modelID, userID, act, actionContent, note)
	}
}

func (s *Service) Post(svc string, modelID uuid.UUID, authorID uuid.UUID, act string, content interface{}, note string) {
	go func() {
		_, err := s.New(svc, modelID, authorID, act, content, note)
		if err != nil {
			s.logger.Warn(fmt.Sprintf("unable to save new action: %+v", err))
		}
	}()
}

func (s *Service) List(params *query.Params) (Actions, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyAction, params, query.DefaultCreatedOrdering...)

	var dtos []actionDTO
	q := query.SQLSelect("*", util.KeyAction, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)

	if err != nil {
		return nil, err
	}

	return toActions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Action, error) {
	dto := actionDTO{}
	q := query.SQLSelect("*", util.KeyAction, "id = $1", "", 0, 0)
	err := s.db.Get(&dto, q, nil, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return dto.ToAction(), nil
}

func (s *Service) GetByAuthor(id uuid.UUID, params *query.Params) (Actions, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyAction, params, query.DefaultCreatedOrdering...)

	var dtos []actionDTO
	q := query.SQLSelect("*", util.KeyAction, "author_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, id)

	if err != nil {
		return nil, err
	}

	return toActions(dtos), nil
}

func (s *Service) GetBySvcModel(svc string, modelID uuid.UUID, params *query.Params) (Actions, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyAction, params, query.DefaultCreatedOrdering...)

	var dtos []actionDTO

	q := query.SQLSelect("*", util.KeyAction, "svc = $1 and model_id = $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, svc, modelID)

	if err != nil {
		return nil, err
	}

	return toActions(dtos), nil
}

func toActions(dtos []actionDTO) Actions {
	ret := make(Actions, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.ToAction())
	}

	return ret
}
