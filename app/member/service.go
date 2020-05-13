package member

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db        *sqlx.DB
	svc       string
	tableName string
	colName   string
}

func NewService(db *sqlx.DB, svc string, table string, col string) *Service {
	return &Service{
		db:        db,
		svc:       svc,
		tableName: table,
		colName:   col,
	}
}

var nameClause = "case when name = '' then (select name from system_user su where su.id = user_id) else name end as name"

func (s *Service) GetByModelID(id uuid.UUID) ([]*Entry, error) {
	var dtos []entryDTO
	q := fmt.Sprintf("select user_id, %s, role, created from %s where %s = $1 order by lower(name)", nameClause, s.tableName, s.colName)
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
	q := fmt.Sprintf("select user_id, %s, role, created from %s where %s = $1 and user_id = $2", nameClause, s.tableName, s.colName)
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
		q := fmt.Sprintf(strings.TrimSpace(`
			insert into %s (%s, user_id, name, role) 
			values ($1, $2, '', 'member')
		`), s.tableName, s.colName)
		_, err = s.db.Exec(q, modelID, userID)
		if err != nil {
			return nil, false, err
		}
		entry, err := s.Get(modelID, userID)
		return entry, true, err
	} else {
		// q := fmt.Sprintf("update %s set name = $1, role = $2 where %s = $3 and user_id = $4)", s.tableName, s.colName)
		// _, err = s.db.Exec(q, modelID, userID, name, role)
		// if err != nil {
		// 	return nil, err
		// }
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
