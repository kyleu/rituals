package database

import (
	"database/sql"
	"emperror.dev/errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Service struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) StartTransaction() (*sqlx.Tx, error) {
	return s.db.Beginx()
}

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
		return 0, errors.Wrap(err, errMessage("delete", q, values))
	}
	aff, err := ret.RowsAffected()
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return int(aff), nil
}

func errMessage(t string, q string, values []interface{}) string {
	valueStrings := make([]string, len(values))
	for i, v := range values {
		valueStrings[i] = fmt.Sprintf("\"%v\"", v)
	}
	vs := strings.Join(valueStrings, ", ")
	return fmt.Sprintf("error running %v sql [%v] with values [%v]", t, q, vs)
}
