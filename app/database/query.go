package database

import (
	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
)

func (s *Service) Query(q string, tx *sqlx.Tx, values ...interface{}) (*sqlx.Rows, error) {
	if tx == nil {
		return s.db.Queryx(q, values...)
	}
	return tx.Queryx(q, values...)
}

func (s *Service) Select(dest interface{}, q string, tx *sqlx.Tx, values ...interface{}) error {
	if tx == nil {
		return s.db.Select(dest, q, values...)
	}
	return tx.Select(dest, q, values...)
}

func (s *Service) Get(dto interface{}, q string, tx *sqlx.Tx, values ...interface{}) error {
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
