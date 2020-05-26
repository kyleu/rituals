package database

import (
	"database/sql"
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
)

func (s *Service) Insert(q string, tx *sqlx.Tx, values ...interface{}) error {
	var err error
	var ret sql.Result
	if tx == nil {
		r, e := s.db.Exec(q, values...)
		ret = r
		err = e
	} else {
		r, e := tx.Exec(q, values...)
		ret = r
		err = e
	}
	if err != nil {
		return errors.Wrap(err, errMessage("insert", q, values))
	}
	aff, err := ret.RowsAffected()
	if err != nil || aff == 0 {
		return errors.Wrap(err, fmt.Sprintf("No rows affected by insert using sql [%v] and %v values", q, len(values)))
	}
	return nil
}

func (s *Service) Update(q string, tx *sqlx.Tx, values ...interface{}) (int, error) {
	return s.Exec(q, tx, values...)
}

func (s *Service) UpdateOne(q string, tx *sqlx.Tx, values ...interface{}) error {
	aff, err := s.Update(q, tx, values...)
	if err != nil {
		return err
	}
	if aff != 1 {
		valueStrings := make([]string, len(values))
		for i, v := range values {
			valueStrings[i] = fmt.Sprintf("\"%v\"", v)
		}
		vs := strings.Join(valueStrings, ", ")
		msg := fmt.Sprintf("expected one row, but [%v] records affected from sql [%v] with values [%s]", aff, q, vs)
		return errors.New(msg)
	}
	return nil
}

func (s *Service) Delete(q string, tx *sqlx.Tx, values ...interface{}) (int, error) {
	return s.Exec(q, tx, values...)
}

func (s *Service) Exec(q string, tx *sqlx.Tx, values ...interface{}) (int, error) {
	var err error
	var ret sql.Result
	if tx == nil {
		r, e := s.db.Exec(q, values...)
		ret = r
		err = e
	} else {
		r, e := tx.Exec(q, values...)
		ret = r
		err = e
	}
	if err != nil {
		return 0, errors.Wrap(err, errMessage("exec", q, values))
	}
	aff, err := ret.RowsAffected()
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return int(aff), nil
}
