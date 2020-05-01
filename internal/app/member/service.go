package member

import (
	"database/sql"
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/internal/app/util"
)

type Service struct {
	db        *sqlx.DB
	tableName string
	colName   string
}

func NewMemberService(db *sqlx.DB, table string, col string) Service {
	return Service{
		db:        db,
		tableName: table,
		colName:   col,
	}
}

func (s *Service) GetByModelID(id uuid.UUID) ([]Entry, error) {
	var dtos []entryDTO
	q := fmt.Sprintf("select user_id, name, role, created from %s where %s = $1", s.tableName, s.colName)
	err := s.db.Select(&dtos, q, id)
	if err != nil {
		return nil, err
	}
	ret := make([]Entry, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToEntry())
	}
	return ret, nil
}

func (s *Service) Get(modelID uuid.UUID, userID uuid.UUID) (*Entry, error) {
	dto := entryDTO{}
	q := fmt.Sprintf("select user_id, name, role, created from %s where %s = $1 and user_id = $2", s.tableName, s.colName)
	err := s.db.Get(&dto, q, modelID, userID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	ret := dto.ToEntry()
	return &ret, nil
}

func (s *Service) Register(modelID uuid.UUID, userID uuid.UUID) (*Entry, bool, error) {
	dto, err := s.Get(modelID, userID)
	if err != nil {
		return nil, false, err
	}
	if dto == nil {
		q := fmt.Sprintf(strings.TrimSpace(`
			insert into %s (%s, user_id, name, role) 
			values ($1, $2, (select name from system_user where id = $2), 'member')
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

func (s *Service) NewSlugFor(str string) (string, error) {
	if len(str) == 0 {
		str = util.RandomString(4)
	}
	slug := strings.ReplaceAll(strings.ToLower(str), " ", "-")
	q := "select id from estimate where slug = $1"
	x, err := s.db.Queryx(q, slug)
	if err != nil {
		return slug, errors.WithStack(errors.Wrap(err, "error fetching existing session"))
	}
	if x.Next() {
		slug, err = s.NewSlugFor(slug + "-" + util.RandomString(4))
		if err != nil {
			return slug, errors.WithStack(errors.Wrap(err, "error recursing for new session"))
		}
	}
	return slug, nil
}
