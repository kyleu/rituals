package team

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"github.com/kyleu/rituals.dev/app/comment"

	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/rituals.dev/app/session"

	"github.com/kyleu/rituals.dev/app/action"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	Data   *session.DataServices
	db     *npndatabase.Service
	logger logur.Logger
	svc    util.Service
}

func NewService(actions *action.Service, users user.Service, comments *comment.Service, db *npndatabase.Service, logger logur.Logger) *Service {
	svc := util.SvcTeam
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: svc.Key})

	data := session.NewDataServices(svc, actions, users, comments, db, logger)
	return &Service{Data: data, db: db, logger: logger, svc: svc}
}

func (s *Service) New(title string, userID uuid.UUID, memberName string) (*Session, error) {
	slug, err := s.Data.History.NewSlugFor(nil, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating team slug")
	}

	model := NewSession(title, slug, userID)

	q := npndatabase.SQLInsert(s.svc.Key, []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.KeyStatus, npncore.KeyOwner}, 1)
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.Status.String(), model.Owner)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new team session")
	}

	s.Data.Members.Register(model.ID, userID, memberName, member.RoleOwner)

	s.Data.Actions.Post(s.svc.Key, model.ID, userID, action.ActCreate, nil)
	return &model, nil
}

func (s *Service) List(params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving team sessions: %+v", err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Session {
	dto := &sessionDTO{}
	q := npndatabase.SQLSelectSimple("*", s.svc.Key, npncore.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Warn(fmt.Sprintf("error getting estimate by id: %+v", err))
		return nil
	}
	return dto.toSession()
}

func (s *Service) GetByIDPointer(teamID *uuid.UUID) *Session {
	if teamID == nil {
		return nil
	}
	return s.GetByID(*teamID)
}

func (s *Service) GetBySlug(slug string) *Session {
	var dto = &sessionDTO{}
	q := npndatabase.SQLSelectSimple("*", s.svc.Key, "slug = $1")
	err := s.db.Get(dto, q, nil, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			h := s.Data.History.Get(slug)
			if h == nil {
				return nil
			}
			return s.GetByID(h.ModelID)
		}
		s.logger.Warn(fmt.Sprintf("error getting estimate by slug: %+v", err))
		return nil
	}
	return dto.toSession()
}

func (s *Service) GetByMember(userID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	t := "team join team_member m on id = m." + npncore.WithDBID(s.svc.Key)
	q := npndatabase.SQLSelect("team.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving teams for user [%v]: %+v", userID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetIdsByMember(userID uuid.UUID) []uuid.UUID {
	var ids []uuid.UUID
	t := "team join team_member m on id = m." + npncore.WithDBID(s.svc.Key)
	q := npndatabase.SQLSelectSimple(npncore.KeyID, t, "m.user_id = $1")
	err := s.db.Select(&ids, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving team ids for user [%v]: %+v", userID, err))
		return nil
	}
	return ids
}

func (s *Service) GetByCreated(d *time.Time, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "created between $1 and $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, d, d.Add(npncore.HoursInDay*time.Hour))
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving teams created on [%v]: %+v", d, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, userID uuid.UUID) error {
	cols := []string{"title"}
	q := npndatabase.SQLUpdate(s.svc.Key, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, title, sessionID)
	s.Data.Actions.Post(s.svc.Key, sessionID, userID, action.ActUpdate, nil)
	return errors.Wrap(err, "error updating team session")
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toSession())
	}
	return ret
}
