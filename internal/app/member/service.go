package member

import (
	"database/sql"
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"strings"
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
	var dtos []Entry
	q := fmt.Sprintf("select user_id, name, role, created from %s where %s = $1", s.tableName, s.colName)
	err := s.db.Select(&dtos, q, id)
	if err != nil {
		return nil, err
	}
	return dtos, nil
}

func (s *Service) Get(modelID uuid.UUID, userID uuid.UUID) (*Entry, error) {
	dto := &Entry{}
	q := fmt.Sprintf("select user_id, name, role, created from %s where %s = $1 and user_id = $2", s.tableName, s.colName)
	err := s.db.Select(&dto, q, modelID, userID)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (s *Service) Register(modelID uuid.UUID, userID uuid.UUID, name string, role string) (*Entry, error) {
	_, err := s.Get(modelID, userID)
	if err == nil {
		q := fmt.Sprintf("update %s set name = $1, role = $2 where %s = $3 and user_id = $4)", s.tableName, s.colName)
		_, err = s.db.Exec(q, modelID, userID, name, role)
		if err != nil {
			return nil, err
		}
		return s.Get(modelID, userID)
	} else {
		if err == sql.ErrNoRows {
			q := fmt.Sprintf("insert into %s (%s, user_id, name, role) values ($1, $2, $3, $4)", s.tableName, s.colName)
			_, err = s.db.Exec(q, modelID, userID, name, role)
			if err != nil {
			  return nil, err
			}
			return s.Get(modelID, userID)
		} else {
			return nil, err
		}
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
