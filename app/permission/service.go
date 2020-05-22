package permission

import (
	"database/sql"
	"fmt"

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

func (s *Service) GetByModelID(id uuid.UUID, params *query.Params) Permissions {
	params = query.ParamsWithDefaultOrdering(util.KeyMember, params, query.DefaultCreatedOrdering...)
	var dtos []permissionDTO
	where := fmt.Sprintf("%s = $1", s.colName)
	q := query.SQLSelect(fmt.Sprintf("k, v, access, created"), s.tableName, where, params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving permission entries for model [%v]: %+v", id, err))
		return nil
	}
	ret := make(Permissions, 0, len(dtos))
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
