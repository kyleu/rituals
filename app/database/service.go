package database

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
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

func errMessage(t string, q string, values []interface{}) string {
	valueStrings := make([]string, len(values))
	for i, v := range values {
		valueStrings[i] = fmt.Sprintf("\"%v\"", v)
	}
	vs := strings.Join(valueStrings, ", ")
	return fmt.Sprintf("error running %v sql [%v] with values [%v]", t, q, vs)
}
