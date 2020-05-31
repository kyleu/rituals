package database

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	debug  bool
	db     *sqlx.DB
	logger logur.Logger
}

func NewService(debug bool, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: "db"})
	return &Service{debug: debug, db: db, logger: logger}
}

func (s *Service) StartTransaction() (*sqlx.Tx, error) {
	if s.debug {
		s.logger.Debug("opening transaction")
	}
	return s.db.Beginx()
}

func errMessage(t string, q string, values []interface{}) string {
	return fmt.Sprintf("error running %v sql [%v] with values [%v]", t, strings.TrimSpace(q), util.ValueStrings(values))
}

func logQuery(s *Service, msg string, q string, values []interface{}) {
	util.LogDebug(s.logger, "%v {\n  SQL: %v\n  Values: %v\n}", msg, strings.TrimSpace(q), util.ValueStrings(values))
}
