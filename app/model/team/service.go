package team

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/history"
	"github.com/kyleu/rituals.dev/app/model/session"
	"github.com/kyleu/rituals.dev/app/model/user"

	"github.com/kyleu/rituals.dev/app/database"

	"github.com/kyleu/rituals.dev/app/model/permission"

	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	Data   *session.DataServices
	db     *database.Service
	logger logur.Logger
	svc    util.Service
}

func NewService(actions *action.Service, users *user.Service, db *database.Service, logger logur.Logger) *Service {
	svc := util.SvcTeam
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: svc.Key})

	data := session.DataServices{
		Members:     member.NewService(actions, users, db, logger, svc),
		Comments:    comment.NewService(actions, db, logger, svc),
		Permissions: permission.NewService(actions, db, logger, svc),
		History:     history.NewService(actions, db, logger, svc),
		Actions:     actions,
	}

	return &Service{Data: &data, db: db, logger: logger, svc: svc}
}

func (s *Service) New(title string, userID uuid.UUID) (*Session, error) {
	slug, err := s.Data.History.NewSlugFor(title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating team slug")
	}

	model := NewSession(title, slug, userID)

	q := query.SQLInsert(s.svc.Key, []string{util.KeyID, util.KeySlug, util.KeyTitle, util.KeyOwner}, 1)
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.Owner)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new team session")
	}

	s.Data.Members.Register(model.ID, userID, member.RoleOwner)

	s.Data.Actions.Post(s.svc, model.ID, userID, action.ActCreate, nil, "")
	return &model, nil
}

func (s *Service) List(params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(s.svc.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", s.svc.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving team sessions: %+v", err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Session {
	dto := &sessionDTO{}
	q := query.SQLSelectSimple("*", s.svc.Key, util.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Warn(fmt.Sprintf("error getting estimate by id: %+v", err))
		return nil
	}
	return dto.ToSession()
}

func (s *Service) GetBySlug(slug string) *Session {
	var dto = &sessionDTO{}
	q := query.SQLSelectSimple("*", s.svc.Key, "slug = $1")
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
	return dto.ToSession()
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(s.svc.Key, params, query.DefaultMCreatedOrdering...)
	var dtos []sessionDTO
	t := "team join team_member m on id = m." + util.WithDBID(s.svc.Key)
	q := query.SQLSelect("team.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving teams for user [%v]: %+v", userID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetIdsByMember(userID uuid.UUID) []uuid.UUID {
	var ids []uuid.UUID
	t := "team join team_member m on id = m." + util.WithDBID(s.svc.Key)
	q := query.SQLSelectSimple(util.KeyID, t, "m.user_id = $1")
	err := s.db.Select(&ids, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving team ids for user [%v]: %+v", userID, err))
		return nil
	}
	return ids
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, userID uuid.UUID) error {
	cols := []string{"title"}
	q := query.SQLUpdate(s.svc.Key, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, title, sessionID)
	s.Data.Actions.Post(s.svc, sessionID, userID, action.ActUpdate, nil, "")
	return errors.Wrap(err, "error updating team session")
}

func (s *Service) GetByIDPointer(teamID *uuid.UUID) *Session {
	if teamID == nil {
		return nil
	}
	return s.GetByID(*teamID)
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
