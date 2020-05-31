package action

import (
	"database/sql"
	"github.com/kyleu/rituals.dev/app/database"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
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

func (s *Service) New(svc util.Service, modelID uuid.UUID, userID uuid.UUID, act string, content interface{}, note string) (*Action, error) {
	id := util.UUID()
	q := query.SQLInsert(util.KeyAction, []string{util.KeyID, util.KeySvc, util.WithDBID(util.KeyModel), util.WithDBID(util.KeyUser), util.KeyAct, util.KeyContent, util.KeyNote}, 1)
	err := s.db.Insert(q, nil, id, svc.Key, modelID, userID, act, util.ToJSON(content, s.logger), note)

	if err != nil {
		return nil, errors.Wrap(err, "error saving new ["+svc.Key+"] action")
	}

	return s.GetByID(id), nil
}

func (s *Service) PostRef(svc util.Service, modelID *uuid.UUID, refSvc util.Service, refID uuid.UUID, userID uuid.UUID, act string, note string) {
	if modelID != nil {
		actionContent := map[string]interface{}{util.KeySvc: refSvc.Key, util.KeyID: refID}
		s.Post(svc, *modelID, userID, act, actionContent, note)
	}
}

func (s *Service) Post(svc util.Service, modelID uuid.UUID, userID uuid.UUID, act string, content interface{}, note string) {
	go func() {
		_, err := s.New(svc, modelID, userID, act, content, note)
		if err != nil {
			util.LogError(s.logger, "unable to save new action: %+v", err)
		}
	}()
}

func (s *Service) List(params *query.Params) Actions {
	params = query.ParamsWithDefaultOrdering(util.KeyAction, params, query.DefaultCreatedOrdering...)

	var dtos []actionDTO
	q := query.SQLSelect("*", util.KeyAction, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)

	if err != nil {
		util.LogError(s.logger, "error retrieving actions: %+v", err)
		return nil
	}

	return toActions(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Action {
	dto := actionDTO{}
	q := query.SQLSelectSimple("*", util.KeyAction, util.KeyID + " = $1")
	err := s.db.Get(&dto, q, nil, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		util.LogError(s.logger, "error getting action by id [%v]: %+v", id, err)
		return nil
	}

	return dto.ToAction()
}

func (s *Service) GetByUser(userID uuid.UUID, params *query.Params) Actions {
	params = query.ParamsWithDefaultOrdering(util.KeyAction, params, query.DefaultCreatedOrdering...)

	var dtos []actionDTO
	q := query.SQLSelect("*", util.KeyAction, "user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)

	if err != nil {
		util.LogError(s.logger, "error retrieving actions for user [%v]: %+v", userID, err)
		return nil
	}
	return toActions(dtos)
}

func (s *Service) GetBySvcModel(svc util.Service, modelID uuid.UUID, params *query.Params) Actions {
	params = query.ParamsWithDefaultOrdering(util.KeyAction, params, query.DefaultCreatedOrdering...)

	var dtos []actionDTO

	q := query.SQLSelect("*", util.KeyAction, "svc = $1 and model_id = $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, svc.Key, modelID)

	if err != nil {
		util.LogError(s.logger, "unable to get actions for [%v:%v]: %+v", svc, modelID, err)
		return nil
	}

	return toActions(dtos)
}

func toActions(dtos []actionDTO) Actions {
	ret := make(Actions, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.ToAction())
	}

	return ret
}
