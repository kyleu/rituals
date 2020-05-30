package database

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) Query(q string, tx *sqlx.Tx, values ...interface{}) (*sqlx.Rows, error) {
	if s.debug {
		logQuery(s, "running raw query", q, values)
	}
	if tx == nil {
		return s.db.Queryx(q, values...)
	}
	return tx.Queryx(q, values...)
}

func (s *Service) Select(dest interface{}, q string, tx *sqlx.Tx, values ...interface{}) error {
	if s.debug {
		logQuery(s, fmt.Sprintf("selecting rows of type [%T]", dest), q, values)
	}
	if tx == nil {
		return s.db.Select(dest, q, values...)
	}
	return tx.Select(dest, q, values...)
}

func (s *Service) Get(dto interface{}, q string, tx *sqlx.Tx, values ...interface{}) error {
	if s.debug {
		logQuery(s, fmt.Sprintf("getting single row of type [%T]", dto), q, values)
		util.LogDebug(s.logger, "getting single row\nSQL: %v\nValues: %v", q, util.ValueStrings(values))
	}
	if tx == nil {
		return s.db.Get(dto, q, values...)
	}
	return tx.Get(dto, q, values...)
}

type countResult struct {
	C int64 `db:"c"`
}

func (s *Service) Count(q string, tx *sqlx.Tx, values ...interface{}) (int64, error) {
	x := &countResult{}
	var err error
	if tx == nil {
		err = s.db.Get(x, q, values...)
	} else {
		err = tx.Get(x, q, values...)
	}
	if err != nil {
		return -1, errors.Wrap(err, "returned value is not an integer")
	}
	return x.C, nil
}
