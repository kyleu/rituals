package team

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/comment"

	"github.com/kyleu/rituals.dev/app/database"

	"github.com/kyleu/rituals.dev/app/permission"

	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions     *action.Service
	db          *database.Service
	Members     *member.Service
	Permissions *permission.Service
	Comments    *comment.Service
	logger      logur.Logger
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.SvcTeam.Key})

	return &Service{
		actions:     actions,
		db:          db,
		Members:     member.NewService(actions, db, logger, util.SvcTeam),
		Permissions: permission.NewService(actions, db, logger, util.SvcTeam),
		Comments:    comment.NewService(actions, db, logger, util.SvcTeam),
		logger:      logger,
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

	s.Members.Register(model.ID, userID, member.RoleOwner)

	s.actions.Post(util.SvcTeam, model.ID, userID, action.ActCreate, nil, "")
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

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	q := query.SQLSelectSimple("*", util.SvcTeam.Key, util.KeyID + " = $1")
	err := s.db.Get(dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetBySlug(slug string) (*Session, error) {
	var dto = &sessionDTO{}
	q := query.SQLSelectSimple("*", util.SvcTeam.Key, "slug = $1")
	err := s.db.Get(dto, q, nil, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
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
	s.actions.Post(util.SvcTeam, sessionID, userID, action.ActUpdate, nil, "")
	return errors.Wrap(err, "error updating team session")
}

func (s *Service) GetByIDPointer(teamID *uuid.UUID) *Session {
	if teamID == nil {
		return nil
	}
	tm, _ := s.GetByID(*teamID)
	return tm
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
