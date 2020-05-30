package team

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/history"
	"github.com/kyleu/rituals.dev/app/model/session"

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
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.SvcTeam.Key})

	data := session.DataServices{
		Members:     member.NewService(actions, db, logger, util.SvcTeam),
		Comments:    comment.NewService(actions, db, logger, util.SvcTeam),
		Permissions: permission.NewService(actions, db, logger, util.SvcTeam),
		History:     history.NewService(db, logger, util.SvcTeam),
		Actions:     actions,
	}

	return &Service{
		Data:   &data,
		db:     db,
		logger: logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcTeam, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating team slug")
	}

	model := NewSession(title, slug, userID)

	q := query.SQLInsert(util.SvcTeam.Key, []string{util.KeyID, util.KeySlug, util.KeyTitle, util.KeyOwner}, 1)
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.Owner)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new team session")
	}

	s.Data.Members.Register(model.ID, userID, member.RoleOwner)

	s.Data.Actions.Post(util.SvcTeam, model.ID, userID, action.ActCreate, nil, "")
	return &model, nil
}

func (s *Service) List(params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(util.SvcTeam.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcTeam.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving team sessions: %+v", err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Session {
	dto := &sessionDTO{}
	q := query.SQLSelectSimple("*", util.SvcTeam.Key, util.KeyID+" = $1")
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
	q := query.SQLSelectSimple("*", util.SvcTeam.Key, "slug = $1")
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
	params = query.ParamsWithDefaultOrdering(util.SvcTeam.Key, params, query.DefaultMCreatedOrdering...)
	var dtos []sessionDTO
	t := "team join team_member m on id = m." + util.WithDBID(util.SvcTeam.Key)
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
	t := "team join team_member m on id = m." + util.WithDBID(util.SvcTeam.Key)
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
	q := query.SQLUpdate(util.SvcTeam.Key, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, title, sessionID)
	s.Data.Actions.Post(util.SvcTeam, sessionID, userID, action.ActUpdate, nil, "")
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
