package permission

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/database"

	"emperror.dev/errors"

	"github.com/kyleu/rituals.dev/app/model/member"
	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"
)

type Service struct {
	actions   *action.Service
	db        *database.Service
	logger    logur.Logger
	svc       util.Service
	tableName string
	colName   string
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger, svc util.Service) *Service {
	return &Service{
		actions:   actions,
		db:        db,
		logger:    logger,
		svc:       svc,
		tableName: svc.Key + "_permission",
		colName:   util.WithDBID(svc.Key),
	}
}

func (s *Service) GetByModelID(id uuid.UUID, params *query.Params) Permissions {
	params = query.ParamsWithDefaultOrdering(util.KeyMember, params, query.DefaultCreatedOrdering...)
	var dtos []permissionDTO
	where := fmt.Sprintf("%s = $1", s.colName)
	q := query.SQLSelect(fmt.Sprintf("k, v, access, created"), s.tableName, where, params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, id)
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
	q := query.SQLSelectSimple("k, v, access, created", s.tableName, where)
	err := s.db.Get(&dto, q, nil, modelID, modelID, k, v)
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
		q := query.SQLInsert(s.tableName, []string{s.colName, "k", "v", "access"}, 1)
		err = s.db.Insert(q, nil, modelID, k, v, access.Key)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error inserting permission for model [%v] and k/v [%v/%v]: %+v", modelID, k, v, err))
		}
	} else {
		q := query.SQLUpdate(s.tableName, []string{"access", "v"}, s.colName + " = $3 and k = $4")
		err = s.db.UpdateOne(q, nil, access.Key, v, modelID, k)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error updating permission for model [%v] and k/v [%v/%v]: %+v", modelID, k, v, err))
		}
	}

	dto, err = s.Get(modelID, k, v)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving updated permission for model [%v] and k/v [%v/%v]: %+v", modelID, k, v, err))
	}

	actionContent := map[string]interface{}{"k": k, "access": access}
	s.actions.Post(s.svc, modelID, userID, action.ActPermissions, actionContent, "")

	return dto
}

func (s *Service) SetAll(modelID uuid.UUID, perms Permissions, userID uuid.UUID) (Permissions, error) {
	tx, err := s.db.StartTransaction()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error starting transaction for model [%v] permissions: %+v", modelID, err))
	}
	defer func() { _ = tx.Commit() }()

	current := s.GetByModelID(modelID, nil)
	var i, u, r Permissions

	for _, p := range perms {
		if current.FindByK(p.K) == nil {
			i = append(i, p)
		} else {
			u = append(u, p)
		}
	}

	for _, c := range current {
		if perms.FindByK(c.K) == nil {
			r = append(r, c)
		}
	}

	for _, p := range u {
		q := query.SQLUpdate(s.tableName, []string{"access", "v"}, s.colName + " = $3 and k = $4")
		err := s.db.UpdateOne(q, tx, p.Access.Key, p.V, modelID, p.K)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error updating permission for model [%v] and k/v [%v/%v]: %+v", modelID, p.K, p.V, err))
		}
	}

	for _, p := range i {
		q := query.SQLInsert(s.tableName, []string{s.colName, "k", "v", "access"}, 1)
		err := s.db.Insert(q, tx, modelID, p.K, p.V, p.Access.Key)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error inserting permission for model [%v] and k/v [%v/%v]: %+v", modelID, p.K, p.V, err))
		}
	}

	for _, p := range r {
		q := query.SQLDelete(s.tableName, fmt.Sprintf("%v = $1 and k = $2", s.colName))
		_, err := s.db.Delete(q, tx, -1, modelID, p.K)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error removing existing permission from model [%v]: %+v", modelID, err))
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error committing permissions transaction for model [%v]: %+v", modelID, err))
	}

	s.actions.Post(s.svc, modelID, userID, action.ActPermissions, perms, "")

	return perms, nil
}
