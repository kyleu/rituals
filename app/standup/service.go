package standup

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

func NewService(actions *action.Service, users *user.Service, comments *comment.Service, db *npndatabase.Service, logger logur.Logger) *Service {
	svc := util.SvcStandup
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: svc.Key})

	data := session.NewDataServices(svc, actions, users, comments, db, logger)
	return &Service{Data: data, db: db, logger: logger, svc: svc}
}

func (s *Service) New(title string, userID uuid.UUID, memberName string, teamID *uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := s.Data.History.NewSlugFor(nil, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating standup slug")
	}

	model := NewSession(title, slug, userID, teamID, sprintID)

	q := npndatabase.SQLInsert(s.svc.Key, []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.WithDBID(util.SvcTeam.Key), npncore.WithDBID(util.SvcSprint.Key), npncore.KeyOwner, npncore.KeyStatus}, 1)
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.TeamID, model.SprintID, model.Owner, model.Status.String())
	if err != nil {
		return nil, errors.Wrap(err, "error saving new standup session")
	}

	s.Data.Members.Register(model.ID, userID, memberName, member.RoleOwner)

	s.Data.Actions.Post(s.svc.Key, model.ID, userID, action.ActCreate, nil)
	s.Data.Actions.PostRef(util.SvcSprint.Key, model.SprintID, s.svc, model.ID, userID, action.ActContentAdd)
	s.Data.Actions.PostRef(util.SvcTeam.Key, model.TeamID, s.svc, model.ID, userID, action.ActContentAdd)

	return &model, nil
}

func (s *Service) List(params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving standup sessions: %+v", err))
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
		s.logger.Error(fmt.Sprintf("error getting standup by slug [%v]: %+v", slug, err))
		return nil
	}
	return dto.toSession()
}

func (s *Service) GetByMember(userID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	t := "standup join standup_member m on id = m." + npncore.WithDBID(s.svc.Key)
	q := npndatabase.SQLSelect("standup.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving standups for user [%v]: %+v", userID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, teamID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving standups for team [%v]: %+v", teamID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetBySprintID(sprintID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "sprint_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, sprintID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving standups for sprint [%v]: %+v", sprintID, err))
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
		s.logger.Error(fmt.Sprintf("error retrieving standups created on [%v]: %+v", d, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, teamID *uuid.UUID, sprintID *uuid.UUID, userID uuid.UUID) error {
	cols := []string{npncore.KeyTitle, npncore.WithDBID(util.SvcTeam.Key), npncore.WithDBID(util.SvcSprint.Key)}
	q := npndatabase.SQLUpdate(s.svc.Key, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, title, teamID, sprintID, sessionID)
	s.Data.Actions.Post(s.svc.Key, sessionID, userID, action.ActUpdate, nil)
	return errors.Wrap(err, "error updating standup session")
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toSession())
	}
	return ret
}
