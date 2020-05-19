package member

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"
)

type Service struct {
	actions   *action.Service
	db        *sqlx.DB
	svc       string
	tableName string
	colName   string
}

func NewService(actions *action.Service, db *sqlx.DB, svc string) *Service {
	return &Service{
		actions:   actions,
		db:        db,
		svc:       svc,
		tableName: svc + "_member",
		colName:   svc + "_id",
	}
}

const nameClause = "case when name = '' then (select name from system_user su where su.id = user_id) else name end as name"

func (s *Service) GetByModelID(id uuid.UUID, params *query.Params) ([]*Entry, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyMember, params, &query.Ordering{Column: "lower(name)", Asc: true})
	var dtos []entryDTO
	q := query.SQLSelect(fmt.Sprintf("user_id, %s, role, created", nameClause), s.tableName, fmt.Sprintf("%s = $1", s.colName), params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, id)
	if err != nil {
		return nil, err
	}
	ret := make([]*Entry, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToEntry())
	}
	return ret, nil
}

func (s *Service) Get(modelID uuid.UUID, userID uuid.UUID) (*Entry, error) {
	dto := entryDTO{}
	q := query.SQLSelect(fmt.Sprintf("user_id, %s, role, created", nameClause), s.tableName, fmt.Sprintf("%s = $1 and user_id = $2", s.colName), "", 0, 0)
	err := s.db.Get(&dto, q, modelID, userID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return dto.ToEntry(), nil
}

func (s *Service) Register(modelID uuid.UUID, userID uuid.UUID) (*Entry, bool, error) {
	dto, err := s.Get(modelID, userID)
	if err != nil {
		return nil, false, err
	}
	if dto == nil {
		q := fmt.Sprintf(`insert into %s (%s, user_id, name, role) values ($1, $2, '', 'member')`, s.tableName, s.colName)
		_, err = s.db.Exec(q, modelID, userID)
		if err != nil {
			return nil, false, err
		}
		dto, err = s.Get(modelID, userID)

		s.actions.Post(s.svc, modelID, userID, action.ActMemberAdd, nil, "")

		return dto, true, err
	} else {
		return dto, false, nil
	}
}

func (s *Service) UpdateName(modelID uuid.UUID, userID uuid.UUID, name string) (*Entry, error) {
	q := fmt.Sprintf("update %s set name = $1 where %s = $2 and user_id = $3", s.tableName, s.colName)
	_, err := s.db.Exec(q, name, modelID, userID)
	if err != nil {
		return nil, err
	}
	return s.Get(modelID, userID)
}
