package comment

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/database"

	"emperror.dev/errors"

	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"
)

type Service struct {
	svc     util.Service
	actions *action.Service
	db      *database.Service
	logger  logur.Logger
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger, svc util.Service) *Service {
	return &Service{
		svc:     svc,
		actions: actions,
		db:      db,
		logger:  logger,
	}
}

func (s *Service) GetByModelID(modelID uuid.UUID, params *query.Params) Comments {
	var defaultOrdering = query.Orderings{{Column: util.KeyCreated, Asc: false}}
	params = query.ParamsWithDefaultOrdering(util.KeyMember, params, defaultOrdering...)
	var dtos []commentDTO
	q := query.SQLSelect("*", util.KeyComment, "svc = $1 and model_id = $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, s.svc.Key, modelID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving comment entries", util.ToMap(util.KeyModel, modelID, util.KeyError, err)))
		return nil
	}
	ret := make(Comments, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.ToComment())
	}

	return ret
}

func (s *Service) GetByID(id uuid.UUID) (*Comment, error) {
	dto := commentDTO{}
	q := query.SQLSelectSimple("*", util.KeyComment, util.KeyID + " = $1")
	err := s.db.Get(&dto, q, nil, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return dto.ToComment(), nil
}

func (s *Service) RemoveComment(commentID uuid.UUID) error {
	q := query.SQLDelete(util.KeyComment, util.KeyID + " = $1")
	err := s.db.DeleteOne(q, nil, commentID)
	return errors.Wrap(err, "unable to remove comment ["+commentID.String()+"]")
}
