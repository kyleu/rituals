package permission

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/member"
	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"
)

type Service struct {
	actions   *action.Service
	db        *sqlx.DB
	logger    logur.Logger
	svc       string
	tableName string
	colName   string
}

func NewService(actions *action.Service, db *sqlx.DB, logger logur.Logger, svc string) *Service {
	return &Service{
		actions:   actions,
		db:        db,
		logger:    logger,
		svc:       svc,
		tableName: svc + "_permission",
		colName:   svc + "_id",
	}
}

func (s *Service) GetByModelID(id uuid.UUID, params *query.Params) []*Permission {
	params = query.ParamsWithDefaultOrdering(util.KeyMember, params, &query.Ordering{Column: "created", Asc: false})
	var dtos []permissionDTO
	where := fmt.Sprintf("%s = $1", s.colName)
	q := query.SQLSelect(fmt.Sprintf("k, v, access, created"), s.tableName, where, params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving permission entries for model [%v]: %+v", id, err))
		return nil
	}
	ret := make([]*Permission, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToPermission())
	}
	return ret
}

func (s *Service) Get(modelID uuid.UUID, k string, v string) (*Permission, error) {
	dto := permissionDTO{}
	where := fmt.Sprintf("%s = $1 and k = $2 and v = $3", s.colName)
	q := query.SQLSelect("k, v, access, created", s.tableName, where, "", 0, 0)
	err := s.db.Get(&dto, q, modelID, modelID, k, v)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return dto.ToPermission(), nil
}

func (s *Service) Set(modelID uuid.UUID, k string, v string, access member.Role, userID uuid.UUID) *Permission {
	dto, err := s.Get(modelID, k, v)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting existing permission for model [%v] and k/v [%v/%v]: %+v", modelID, k, v, err))
		return nil
	}
	if dto == nil {
		q := fmt.Sprintf(`insert into %s (%s, k, v, access) values ($1, $2, $3, $4)`, s.tableName, s.colName)
		_, err = s.db.Exec(q, modelID, k, v, access.Key)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error inserting permission for model [%v] and k/v [%v/%v]: %+v", modelID, k, v, err))
		}
	} else {
		q := fmt.Sprintf(`update %s set access = $1 where %s = $2 and k = $3 and v = $4`, s.tableName, s.colName)
		_, err = s.db.Exec(q, access.Key, modelID, k, v)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error updating permission for model [%v] and k/v [%v/%v]: %+v", modelID, k, v, err))
		}
	}

	dto, err = s.Get(modelID, k, v)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving updated permission for model [%v] and k/v [%v/%v]: %+v", modelID, k, v, err))
	}

	actionContent := map[string]interface{}{"k": k, "access": access}
	s.actions.Post(s.svc, modelID, userID, action.ActPermissionAdd, actionContent, "")

	return dto
}

func (s *Service) Check(svc util.Service, modelID uuid.UUID, auths []*auth.Record, teamID *uuid.UUID, teamTitle string, teams []uuid.UUID) Errors {
	a := s.CheckAuths(svc, modelID, auths)
	t := s.CheckTeam(svc, teamID, teamTitle, teams)

	if t != nil {
		return append(a, *t)
	}
	return a
}

func (s *Service) CheckAuths(svc util.Service, modelID uuid.UUID, auths []*auth.Record) Errors {
	var ret []Error

	perms := s.GetByModelID(modelID, nil)
	for _, p := range perms {
		switch p.K {
		case auth.ProviderGitHub.Key:
			ret = append(ret, providerError(svc, p.V, auth.ProviderGitHub))
		case auth.ProviderGoogle.Key:
			ret = append(ret, providerError(svc, p.V, auth.ProviderGoogle))
		case auth.ProviderSlack.Key:
			ret = append(ret, providerError(svc, p.V, auth.ProviderSlack))
		default:
			ret = append(ret, Error{K: p.K, V: p.V, Code: "missing", Message: "unhandled permission key [" + p.K + "]"})
		}
	}
	return ret
}

func (s *Service) CheckTeam(svc util.Service, teamID *uuid.UUID, teamTitle string, teams []uuid.UUID) *Error {
	if teamID != nil {
		hasTeam := false

		for _, t := range teams {
			if t == *teamID {
				hasTeam = true
				break
			}
		}
		if !hasTeam {
			msg := fmt.Sprintf("you are not a member of [%v], this %v's team", teamTitle, svc.Key)
			return &Error{K: "team", V: teamID.String(), Code: "team", Message: msg}
		}
	}
	return nil
}

func providerError(svc util.Service, v string, p auth.Provider) Error {
	msg := "you must log in with " + p.Title + " to access this " + svc.Key
	return Error{K: svc.Key, V: v, Code: p.Key, Message: msg}
}
