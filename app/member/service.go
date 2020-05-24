package member

import (
	"database/sql"
	"fmt"

	"emperror.dev/errors"

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
		tableName: svc + "_member",
		colName:   svc + "_id",
	}
}

const nameClause = "case when name = '' then (select name from system_user su where su.id = user_id) else name end as name"

func (s *Service) GetByModelID(id uuid.UUID, params *query.Params) Entries {
	var defaultOrdering = query.Orderings{{Column: "name", Asc: true}}
	params = query.ParamsWithDefaultOrdering(util.KeyMember, params, defaultOrdering...)
	var dtos []entryDTO
	where := fmt.Sprintf("%s = $1", s.colName)
	cols := fmt.Sprintf("user_id, %s, role, created", nameClause)
	q := query.SQLSelect(cols, s.tableName, where, params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving member entries for model [%v]: %+v", id, err))
		return nil
	}
	ret := make(Entries, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.ToEntry())
	}

	return ret
}

func (s *Service) Get(modelID uuid.UUID, userID uuid.UUID) (*Entry, error) {
	dto := entryDTO{}
	cols := fmt.Sprintf("user_id, %s, role, created", nameClause)
	where := fmt.Sprintf("%s = $1 and user_id = $2", s.colName)
	q := query.SQLSelect(cols, s.tableName, where, "", 0, 0)
	err := s.db.Get(&dto, q, modelID, userID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return dto.ToEntry(), nil
}

func (s *Service) UpdateName(modelID uuid.UUID, userID uuid.UUID, name string) (*Entry, error) {
	q := fmt.Sprintf("update %s set name = $1 where %s = $2 and user_id = $3", s.tableName, s.colName)
	_, err := s.db.Exec(q, name, modelID, userID)
	if err != nil {
		return nil, err
	}
	return s.Get(modelID, userID)
}

func (s *Service) RemoveMember(modelID uuid.UUID, target uuid.UUID) error {
	q := fmt.Sprintf("delete from %s where %s = $1 and user_id = $2", s.tableName, s.colName)
	_, err := s.db.Exec(q, modelID, target)
	return errors.WithStack(errors.Wrap(err, "unable to remove member ["+target.String()+"]"))
}
