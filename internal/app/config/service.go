package config

import (
	"fmt"
	"strings"

	"github.com/kyleu/rituals.dev/internal/gen/queries"

	"emperror.dev/errors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/internal/app/conn"
	_ "github.com/mattn/go-sqlite3"
	"logur.dev/logur"
)

type Service struct {
	logger          logur.LoggerFacade
}

func NewService(logger logur.LoggerFacade) (*Service, error) {
	svc := Service{logger: logger}

	err = initIfNeeded(db, logger)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error initializing config database"))
	}

	err = pr.Refresh(db)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error initializing project registry"))
	}

	logger.Debug("Config service started at [" + root.URL + "]")
	return &svc, nil
}

func initIfNeeded(db *sqlx.DB, logger logur.LoggerFacade) error {
	exec("burn-it-down", db, logger, func(sb *strings.Builder) { queries.ResetConfigDatabase(sb) })
	exec("create-table-project", db, logger, func(sb *strings.Builder) { queries.CreateTableProject(sb) })
	exec("insert-data-project", db, logger, func(sb *strings.Builder) { queries.InsertDataProject(sb) })
	return nil
}

func exec(name string, db *sqlx.DB, logger logur.LoggerFacade, f func(*strings.Builder)) {
	sb := &strings.Builder{}
	f(sb)
	result, err := conn.ExecuteNoTx(logger, db, conn.Adhoc(sb.String()))
	if err != nil {
		panic(errors.WithStack(err))
	}
	logger.Debug(fmt.Sprintf("Ran [%s] in [%vms], [%v] rows affected", name, result.Timing.Elapsed, result.RowsAffected))
}
