package action

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	db     *npndatabase.Service
	logger logur.Logger
}

func NewService(db *npndatabase.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeyAction})
	svc := Service{
		db:     db,
		logger: logger,
	}

	return &svc
}

func (s *Service) New(svc string, modelID uuid.UUID, userID uuid.UUID, act string, content interface{}, note string) (*Action, error) {
	id := npncore.UUID()
	q := npndatabase.SQLInsert(npncore.KeyAction, []string{npncore.KeyID, npncore.KeySvc, npncore.WithDBID(npncore.KeyModel), npncore.WithDBID(npncore.KeyUser), npncore.KeyAct, npncore.KeyContent, npncore.KeyNote}, 1)
	err := s.db.Insert(q, nil, id, svc, modelID, userID, act, npncore.ToJSON(content, s.logger), note)

	if err != nil {
		return nil, errors.Wrap(err, "error saving new ["+svc+"] action")
	}

	return s.GetByID(id), nil
}

func (s *Service) PostRef(svc string, modelID *uuid.UUID, refSvc util.Service, refID uuid.UUID, userID uuid.UUID, act string, notes ...string) {
	if modelID != nil {
		actionContent := map[string]interface{}{npncore.KeySvc: refSvc.Key, npncore.KeyID: refID}
		s.Post(svc, *modelID, userID, act, actionContent, notes...)
	}
}

func (s *Service) Post(svc string, modelID uuid.UUID, userID uuid.UUID, act string, content interface{}, notes ...string) {
	go func() {
		_, err := s.New(svc, modelID, userID, act, content, strings.Join(notes, "\n\n"))
		if err != nil {
			s.logger.Error(fmt.Sprintf("unable to save new action: %+v", err))
		}
	}()
}

func (s *Service) List(params *npncore.Params) Actions {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyAction, params, npncore.DefaultCreatedOrdering...)

	var dtos []actionDTO
	q := npndatabase.SQLSelect("*", npncore.KeyAction, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving actions: %+v", err))
		return nil
	}

	return toActions(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Action {
	dto := actionDTO{}
	q := npndatabase.SQLSelectSimple("*", npncore.KeyAction, npncore.KeyID+" = $1")
	err := s.db.Get(&dto, q, nil, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting action by id [%v]: %+v", id, err))
		return nil
	}

	return dto.toAction()
}

func (s *Service) GetByUser(userID uuid.UUID, params *npncore.Params) Actions {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyAction, params, npncore.DefaultCreatedOrdering...)

	var dtos []actionDTO
	q := npndatabase.SQLSelect("*", npncore.KeyAction, "user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving actions for user [%v]: %+v", userID, err))
		return nil
	}
	return toActions(dtos)
}

func (s *Service) GetBySvcModel(svc string, modelID uuid.UUID, params *npncore.Params) Actions {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyAction, params, npncore.DefaultCreatedOrdering...)

	var dtos []actionDTO

	q := npndatabase.SQLSelect("*", npncore.KeyAction, "svc = $1 and model_id = $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, svc, modelID)

	if err != nil {
		s.logger.Error(fmt.Sprintf("unable to get actions for [%v:%v]: %+v", svc, modelID, err))
		return nil
	}

	return toActions(dtos)
}

func toActions(dtos []actionDTO) Actions {
	ret := make(Actions, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.toAction())
	}

	return ret
}
