package sprint

import (
	"database/sql"
	"github.com/kyleu/rituals.dev/app/database"
	"time"

	"github.com/kyleu/rituals.dev/app/permission"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions     *action.Service
	db          *database.Service
	Members     *member.Service
	Permissions *permission.Service
	logger      logur.Logger
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.SvcRetro.Key})

	return &Service{
		actions:     actions,
		db:          db,
		Members:     member.NewService(actions, db, logger, util.SvcSprint.Key),
		Permissions: permission.NewService(actions, db, logger, util.SvcSprint.Key),
		logger:      logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID, teamID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcSprint.Key, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sprint slug")
	}

	model := NewSession(title, slug, userID, teamID, nil, nil)

	q := "insert into sprint (id, slug, title, team_id, owner, start_date, end_date) values ($1, $2, $3, $4, $5, $6, $7)"
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.TeamID, model.Owner, model.StartDate, model.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new sprint session")
	}

	s.Members.Register(model.ID, userID)

	s.actions.Post(util.SvcSprint.Key, model.ID, userID, action.ActCreate, nil, "")
	s.actions.PostRef(util.SvcTeam.Key, model.TeamID, util.SvcSprint.Key, model.ID, userID, action.ActContentAdd, "")

	return &model, nil
}

func (s *Service) List(params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcSprint.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	q := query.SQLSelect("*", util.SvcSprint.Key, "id = $1", "", 0, 0)
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
	q := query.SQLSelect("*", util.SvcSprint.Key, "slug = $1", "", 0, 0)
	err := s.db.Get(dto, q, nil, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetByOwner(userID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcSprint.Key, "owner = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, query.DefaultMCreatedOrdering...)
	var dtos []sessionDTO
	t := "sprint join sprint_member m on id = m.sprint_id"
	q := query.SQLSelect("sprint.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetIdsByMember(userID uuid.UUID) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	t := "sprint join sprint_member m on id = m.sprint_id"
	q := query.SQLSelect("id", t, "m.user_id = $1", "", 0, 0)
	err := s.db.Select(&ids, q, nil, userID)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *query.Params) (Sessions, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcSprint.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcSprint.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, teamID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, teamID *uuid.UUID, startDate *time.Time, endDate *time.Time, userID uuid.UUID) error {
	q := "update sprint set title = $1, team_id = $2, start_date = $3, end_date = $4 where id = $5"
	err := s.db.UpdateOne(q, nil, title, teamID, startDate, endDate, sessionID)
	s.actions.Post(util.SvcSprint.Key, sessionID, userID, action.ActUpdate, nil, "")
	return errors.Wrap(err, "error updating sprint session")
}

func (s *Service) GetByIDPointer(sprintID *uuid.UUID) *Session {
	if sprintID == nil {
		return nil
	}
	spr, _ := s.GetByID(*sprintID)
	return spr
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
