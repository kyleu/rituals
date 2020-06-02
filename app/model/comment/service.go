package comment

import (
	"database/sql"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/database"

	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"
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
	var defaultOrdering = query.Orderings{{Column: util.KeyCreated, Asc: true}}
	params = query.ParamsWithDefaultOrdering(util.KeyMember, params, defaultOrdering...)
	var dtos []commentDTO
	q := query.SQLSelect("*", util.KeyComment, "svc = $1 and model_id = $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, s.svc.Key, modelID)
	if err != nil {
		util.LogError(s.logger, "error retrieving comment entries for [%v:%v]: %+v", s.svc.Key, modelID, err)
		return nil
	}
	ret := make(Comments, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.ToComment())
	}

	return ret
}

func (s *Service) GetByID(id uuid.UUID) *Comment {
	dto := commentDTO{}
	q := query.SQLSelectSimple("*", util.KeyComment, util.KeyID+" = $1")
	err := s.db.Get(&dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		util.LogError(s.logger, "error getting comment by id [%v]: %+v", id, err)
		return nil
	}

	return dto.ToComment()
}

func (s *Service) RemoveComment(commentID uuid.UUID) error {
	q := query.SQLDelete(util.KeyComment, util.KeyID+" = $1")
	err := s.db.DeleteOne(q, nil, commentID)
	return errors.Wrap(err, "unable to remove comment ["+commentID.String()+"]")
}

func (s *Service) Add(svc util.Service, modelID uuid.UUID, targetType string, targetID *uuid.UUID, content string, userID uuid.UUID) (*Comment, error) {
	id := util.UUID()
	html := util.ToHTML(content)
	q := query.SQLInsert(util.KeyComment, []string{"id", "svc", "model_id", "target_type", "target_id", "user_id", "content", "html"}, 1)
	err := s.db.Insert(q, nil, id, svc.Key, modelID, targetType, targetID, userID, content, html)
	if err != nil {
		return nil, errors.Wrap(err, "unable to add comment")
	}
	return s.GetByID(id), nil
}

func (s *Service) Update(id uuid.UUID, content string, userID uuid.UUID) (*Comment, error) {
	html := util.ToHTML(content)
	q := query.SQLUpdate(util.KeyComment, []string{"content", "html"}, "id = $3")
	err := s.db.UpdateOne(q, nil, content, html, id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to update comment")
	}
	return s.GetByID(id), nil
}
