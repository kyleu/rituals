package comment

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"emperror.dev/errors"
	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
)

type Service struct {
	actions *action.Service
	db      *npndatabase.Service
	logger  logur.Logger
}

func NewService(actions *action.Service, db *npndatabase.Service, logger logur.Logger) *Service {
	return &Service{
		actions: actions,
		db:      db,
		logger:  logger,
	}
}

func (s *Service) List(params *npncore.Params) Comments {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyComment, params, npncore.DefaultCreatedOrdering...)

	var dtos []commentDTO
	q := npndatabase.SQLSelect("*", npncore.KeyComment, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving comments: %+v", err))
		return nil
	}

	return toComments(dtos)
}

func (s *Service) GetByModelID(svc util.Service, modelID uuid.UUID, params *npncore.Params) Comments {
	var defaultOrdering = npncore.Orderings{{Column: npncore.KeyCreated, Asc: true}}
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyComment, params, defaultOrdering...)
	var dtos []commentDTO
	q := npndatabase.SQLSelect("*", npncore.KeyComment, "svc = $1 and model_id = $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, svc.Key, modelID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving comment entries for [%v:%v]: %+v", svc.Key, modelID, err))
		return nil
	}
	return toComments(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Comment {
	dto := commentDTO{}
	q := npndatabase.SQLSelectSimple("*", npncore.KeyComment, npncore.KeyID+" = $1")
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

func (s *Service) Add(svc string, modelID uuid.UUID, targetType string, targetID *uuid.UUID, content string, userID uuid.UUID) (*Comment, error) {
	id := npncore.UUID()
	html := util.ToHTML(content)
	q := npndatabase.SQLInsert(npncore.KeyComment, []string{
		npncore.KeyID, npncore.KeySvc, npncore.WithDBID(npncore.KeyModel), "target_type", npncore.WithDBID("target"),
		npncore.WithDBID(npncore.KeyUser), npncore.KeyContent, npncore.KeyHTML,
	}, 1)
	err := s.db.Insert(q, nil, id, svc, modelID, targetType, targetID, userID, content, html)
	if err != nil {
		return nil, errors.Wrap(err, "unable to add comment")
	}
	return s.GetByID(id), nil
}

func (s *Service) Update(id uuid.UUID, content string, userID uuid.UUID) (*Comment, error) {
	html := util.ToHTML(content)
	q := npndatabase.SQLUpdate(npncore.KeyComment, []string{npncore.KeyContent, npncore.KeyHTML}, "id = $3 and user_id = $4")
	err := s.db.UpdateOne(q, nil, content, html, id, userID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to update comment")
	}
	return s.GetByID(id), nil
}

func (s *Service) RemoveComment(commentID uuid.UUID, userID uuid.UUID) error {
	q := npndatabase.SQLDelete(npncore.KeyComment, npncore.KeyID+" = $1")
	err := s.db.DeleteOne(q, nil, commentID)
	return errors.Wrap(err, "unable to remove comment ["+commentID.String()+"] for user [" + userID.String() + "]")
}

func (s *Service) GetByCreated(d *time.Time, params *npncore.Params) Comments {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyComment, params, npncore.DefaultCreatedOrdering...)
	var dtos []commentDTO
	q := npndatabase.SQLSelect("*", npncore.KeyComment, "created between $1 and $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, d, d.Add(npncore.HoursInDay*time.Hour))
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
