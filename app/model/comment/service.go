package comment

import (
	"database/sql"
	"fmt"
	"time"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/database"

	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"
)

type Service struct {
	actions *action.Service
	db      *database.Service
	logger  logur.Logger
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger) *Service {
	return &Service{
		actions: actions,
		db:      db,
		logger:  logger,
	}
}

func (s *Service) List(params *query.Params) Comments {
	params = query.ParamsWithDefaultOrdering(util.KeyComment, params, query.DefaultCreatedOrdering...)

	var dtos []commentDTO
	q := query.SQLSelect("*", util.KeyComment, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving comments: %+v", err))
		return nil
	}

	return toComments(dtos)
}

func (s *Service) GetByModelID(svc util.Service, modelID uuid.UUID, params *query.Params) Comments {
	var defaultOrdering = query.Orderings{{Column: util.KeyCreated, Asc: true}}
	params = query.ParamsWithDefaultOrdering(util.KeyComment, params, defaultOrdering...)
	var dtos []commentDTO
	q := query.SQLSelect("*", util.KeyComment, "svc = $1 and model_id = $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, svc.Key, modelID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving comment entries for [%v:%v]: %+v", svc.Key, modelID, err))
		return nil
	}
	return toComments(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Comment {
	dto := commentDTO{}
	q := query.SQLSelectSimple("*", util.KeyComment, util.KeyID+" = $1")
	err := s.db.Get(&dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting comment by id [%v]: %+v", id, err))
		return nil
	}

	return dto.toComment()
}

func (s *Service) Add(svc util.Service, modelID uuid.UUID, targetType string, targetID *uuid.UUID, content string, userID uuid.UUID) (*Comment, error) {
	id := util.UUID()
	html := util.ToHTML(content)
	q := query.SQLInsert(util.KeyComment, []string{
		util.KeyID, util.KeySvc, util.WithDBID(util.KeyModel), "target_type", util.WithDBID("target"),
		util.WithDBID(util.KeyUser), util.KeyContent, util.KeyHTML,
	}, 1)
	err := s.db.Insert(q, nil, id, svc.Key, modelID, targetType, targetID, userID, content, html)
	if err != nil {
		return nil, errors.Wrap(err, "unable to add comment")
	}
	return s.GetByID(id), nil
}

func (s *Service) Update(id uuid.UUID, content string, userID uuid.UUID) (*Comment, error) {
	html := util.ToHTML(content)
	q := query.SQLUpdate(util.KeyComment, []string{util.KeyContent, util.KeyHTML}, "id = $3 and user_id = $4")
	err := s.db.UpdateOne(q, nil, content, html, id, userID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to update comment")
	}
	return s.GetByID(id), nil
}

func (s *Service) RemoveComment(commentID uuid.UUID) error {
	q := query.SQLDelete(util.KeyComment, util.KeyID+" = $1")
	err := s.db.DeleteOne(q, nil, commentID)
	return errors.Wrap(err, "unable to remove comment ["+commentID.String()+"]")
}

func (s *Service) GetByCreated(d *time.Time, params *query.Params) Comments {
	params = query.ParamsWithDefaultOrdering(util.KeyComment, params, query.DefaultCreatedOrdering...)
	var dtos []commentDTO
	q := query.SQLSelect("*", util.KeyComment, "created between $1 and $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, d, d.Add(util.HoursInDay*time.Hour))
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving comments created on [%v]: %+v", d, err))
		return nil
	}
	return toComments(dtos)
}

func toComments(dtos []commentDTO) Comments {
	ret := make(Comments, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toComment())
	}
	return ret
}
