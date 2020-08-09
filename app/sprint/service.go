package sprint

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"time"

	"github.com/kyleu/rituals.dev/app/comment"

	"github.com/kyleu/rituals.dev/app/history"
	"github.com/kyleu/rituals.dev/app/session"
	"github.com/kyleu/npn/npnservice/user"

	"github.com/kyleu/rituals.dev/app/permission"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
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

func NewService(actions *action.Service, users *user.Service, comments *comment.Service, db *npndatabase.Service, logger logur.Logger) *Service {
	svc := util.SvcSprint
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: svc.Key})

	data := session.DataServices{
		Svc:         svc,
		Members:     member.NewService(actions, users, db, logger, svc),
		Comments:    comments,
		Permissions: permission.NewService(actions, db, logger, svc),
		History:     history.NewService(actions, db, logger, svc),
		Actions:     actions,
	}

	return &Service{Data: &data, db: db, logger: logger, svc: svc}
}

func (s *Service) New(title string, userID uuid.UUID, memberName string, startDate *time.Time, endDate *time.Time, teamID *uuid.UUID) (*Session, error) {
	slug, err := s.Data.History.NewSlugFor(nil, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sprint slug")
	}

	model := NewSession(title, slug, userID, teamID, startDate, endDate)

	q := npndatabase.SQLInsert(s.svc.Key, []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.WithDBID(util.SvcTeam.Key), npncore.KeyOwner, "start_date", "end_date"}, 1)
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.TeamID, model.Owner, model.StartDate, model.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new sprint session")
	}

	s.Data.Members.Register(model.ID, userID, memberName, member.RoleOwner)

	s.Data.Actions.Post(s.svc, model.ID, userID, action.ActCreate, nil)
	s.Data.Actions.PostRef(util.SvcTeam, model.TeamID, s.svc, model.ID, userID, action.ActContentAdd)

	return &model, nil
}

func (s *Service) List(params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprint sessions: %+v", err))
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
		s.logger.Error(fmt.Sprintf("error getting sprint by id [%v]: %+v", id, err))
		return nil
	}
	return dto.toSession()
}

func (s *Service) GetBySlug(slug string) *Session {
	var dto = &sessionDTO{}
	q := npndatabase.SQLSelectSimple("*", s.svc.Key, "slug = $1")
	err := s.db.Get(dto, q, nil, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting sprint by slug [%v]: %+v", slug, err))
		return nil
	}
	return dto.toSession()
}

func (s *Service) GetByMember(userID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	t := "sprint join sprint_member m on id = m." + npncore.WithDBID(s.svc.Key)
	q := npndatabase.SQLSelect("sprint.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprints for user [%v]: %+v", userID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetIdsByMember(userID uuid.UUID) []uuid.UUID {
	var ids []uuid.UUID
	t := "sprint join sprint_member m on id = m." + npncore.WithDBID(s.svc.Key)
	q := npndatabase.SQLSelectSimple(npncore.KeyID, t, "m.user_id = $1")
	err := s.db.Select(&ids, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprint ids for user [%v]: %+v", userID, err))
		return nil
	}
	return ids
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, teamID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprints for team [%v]: %+v", teamID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByCreated(d *time.Time, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "created between $1 and $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, d, d.Add(npncore.HoursInDay*time.Hour))
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving sprints created on [%v]: %+v", d, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, teamID *uuid.UUID, startDate *time.Time, endDate *time.Time, userID uuid.UUID) error {
	cols := []string{npncore.KeyTitle, "start_date", "end_date", npncore.WithDBID(util.SvcTeam.Key)}
	q := npndatabase.SQLUpdate(s.svc.Key, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, title, startDate, endDate, teamID, sessionID)
	s.Data.Actions.Post(s.svc, sessionID, userID, action.ActUpdate, nil)
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
		ret = append(ret, dto.toSession())
	}
	return ret
}
