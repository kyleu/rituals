package retro

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/history"
	"github.com/kyleu/rituals.dev/app/model/session"
	"github.com/kyleu/rituals.dev/app/model/user"
	"strings"

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

func NewService(actions *action.Service, users *user.Service, db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.SvcRetro.Key})

	data := session.DataServices{
		Members:     member.NewService(actions, users, db, logger, util.SvcRetro),
		Comments:    comment.NewService(actions, db, logger, util.SvcRetro),
		Permissions: permission.NewService(actions, db, logger, util.SvcRetro),
		History:     history.NewService(db, logger, util.SvcRetro),
		Actions:     actions,
	}

	return &Service{Data: &data, db: db, logger: logger}
}

func (s *Service) New(title string, userID uuid.UUID, categories []string, teamID *uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcRetro, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating retro slug")
	}

	model := NewSession(title, slug, userID, categories, teamID, sprintID)

	q := query.SQLInsert(util.SvcRetro.Key, []string{util.KeyID, util.KeySlug, util.KeyTitle, util.WithDBID(util.SvcTeam.Key), util.WithDBID(util.SvcSprint.Key), util.KeyOwner, util.KeyStatus, util.Plural(util.KeyCategory)}, 1)
	categoriesString := "{" + strings.Join(model.Categories, ",") + "}"
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.TeamID, model.SprintID, model.Owner, model.Status.String(), categoriesString)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new retro session")
	}

	s.Data.Members.Register(model.ID, userID, member.RoleOwner)

	s.Data.Actions.Post(util.SvcRetro, model.ID, userID, action.ActCreate, nil, "")
	s.Data.Actions.PostRef(util.SvcSprint, model.SprintID, util.SvcRetro, model.ID, userID, action.ActContentAdd, "")
	s.Data.Actions.PostRef(util.SvcTeam, model.TeamID, util.SvcRetro, model.ID, userID, action.ActContentAdd, "")

	return &model, nil
}

func (s *Service) List(params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcRetro.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retro sessions: %+v", err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Session {
	dto := &sessionDTO{}
	q := query.SQLSelectSimple("*", util.SvcRetro.Key, util.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		util.LogError(s.logger, "error getting retro by id [%v]: %+v", id, err)
		return nil
	}
	return dto.ToSession()
}

func (s *Service) GetBySlug(slug string) *Session {
	var dto = &sessionDTO{}
	q := query.SQLSelectSimple("*", util.SvcRetro.Key, "slug = $1")
	err := s.db.Get(dto, q, nil, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		util.LogError(s.logger, "error getting retro by slug [%v]: %+v", slug, err)
		return nil
	}
	return dto.ToSession()
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultMCreatedOrdering...)
	var dtos []sessionDTO
	t := "retro join retro_member m on id = m." + util.WithDBID(util.SvcRetro.Key)
	q := query.SQLSelect("retro.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retros for user [%v]: %+v", userID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcRetro.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, teamID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retros for team [%v]: %+v", teamID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetBySprint(sprintID uuid.UUID, params *query.Params) Sessions {
	params = query.ParamsWithDefaultOrdering(util.SvcRetro.Key, params, query.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcRetro.Key, "sprint_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, sprintID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retros for sprint [%v]: %+v", sprintID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, categories []string, teamID *uuid.UUID, sprintID *uuid.UUID, userID uuid.UUID) error {
	cols := []string{"title", util.Plural(util.KeyCategory), util.WithDBID(util.SvcTeam.Key), util.WithDBID(util.SvcSprint.Key)}
	q := query.SQLUpdate(util.SvcRetro.Key, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	categoriesString := "{" + strings.Join(categories, ",") + "}"
	err := s.db.UpdateOne(q, nil, title, categoriesString, teamID, sprintID, sessionID)
	s.Data.Actions.Post(util.SvcRetro, sessionID, userID, action.ActUpdate, nil, "")
	return errors.Wrap(err, "error updating retro session")
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
