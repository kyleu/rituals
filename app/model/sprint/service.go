package sprint

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/history"
	"github.com/kyleu/rituals.dev/app/model/session"
	"github.com/kyleu/rituals.dev/app/model/user"
	"time"

	"github.com/kyleu/rituals.dev/app/database"

	"github.com/kyleu/rituals.dev/app/model/permission"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	Data   *session.DataServices
	db     *database.Service
	logger logur.Logger
}

func NewService(actions *action.Service, users *user.Service, db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.SvcRetro.Key})

	data := session.DataServices{
		Members:     member.NewService(actions, users, db, logger, util.SvcSprint),
		Comments:    comment.NewService(actions, db, logger, util.SvcSprint),
		Permissions: permission.NewService(actions, db, logger, util.SvcSprint),
		History:     history.NewService(db, logger, util.SvcSprint),
		Actions:     actions,
	}

	return &Service{
		Data:   &data,
		db:     db,
		logger: logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, startDate *time.Time, endDate *time.Time, teamID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcSprint, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sprint slug")
	}

	model := NewSession(title, slug, userID, teamID, startDate, endDate)

	println(fmt.Sprintf("#@@@@@@@@@@@@@: %v", model.EndDate))

	q := query.SQLInsert(util.SvcSprint.Key, []string{util.KeyID, util.KeySlug, util.KeyTitle, util.WithDBID(util.SvcTeam.Key), util.KeyOwner, "start_date", "end_date"}, 1)
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.TeamID, model.Owner, model.StartDate, model.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new sprint session")
	}

	s.Data.Members.Register(model.ID, userID, member.RoleOwner)

	s.Data.Actions.Post(util.SvcSprint, model.ID, userID, action.ActCreate, nil, "")
	s.Data.Actions.PostRef(util.SvcTeam, model.TeamID, util.SvcSprint, model.ID, userID, action.ActContentAdd, "")

	return &model, nil
}

func (s *Service) List(params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcSprint.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprint sessions: %+v", err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Session {
	dto := &sessionDTO{}
	q := query.SQLSelectSimple("*", util.SvcSprint.Key, util.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		util.LogError(s.logger, "error getting sprint by id [%v]: %+v", id, err)
		return nil
	}
	return dto.ToSession()
}

func (s *Service) GetBySlug(slug string) *Session {
	var dto = &sessionDTO{}
	q := query.SQLSelectSimple("*", util.SvcSprint.Key, "slug = $1")
	err := s.db.Get(dto, q, nil, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		util.LogError(s.logger, "error getting sprint by slug [%v]: %+v", slug, err)
		return nil
	}
	return dto.ToSession()
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, query.DefaultMCreatedOrdering...)
	var dtos []sessionDTO
	t := "sprint join sprint_member m on id = m." + util.WithDBID(util.SvcSprint.Key)
	q := query.SQLSelect("sprint.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprints for user [%v]: %+v", userID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetIdsByMember(userID uuid.UUID) []uuid.UUID {
	var ids []uuid.UUID
	t := "sprint join sprint_member m on id = m." + util.WithDBID(util.SvcSprint.Key)
	q := query.SQLSelectSimple(util.KeyID, t, "m.user_id = $1")
	err := s.db.Select(&ids, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprint ids for user [%v]: %+v", userID, err))
		return nil
	}
	return ids
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcSprint.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, teamID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprints for team [%v]: %+v", teamID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, teamID *uuid.UUID, startDate *time.Time, endDate *time.Time, userID uuid.UUID) error {
	cols := []string{util.KeyTitle, "start_date", "end_date", util.WithDBID(util.SvcTeam.Key)}
	q := query.SQLUpdate(util.SvcSprint.Key, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, title, startDate, endDate, teamID, sessionID)
	s.Data.Actions.Post(util.SvcSprint, sessionID, userID, action.ActUpdate, nil, "")
	return errors.Wrap(err, "error updating sprint session")
}

func (s *Service) GetByIDPointer(sprintID *uuid.UUID) *Session {
	if sprintID == nil {
		return nil
	}
	return s.GetByID(*sprintID)
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
