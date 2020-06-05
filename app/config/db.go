package config

import (
	"fmt"

	"github.com/kyleu/rituals.dev/app/database"

	"emperror.dev/errors"
	// load postgres driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type DBParams struct {
	Debug   bool
	Wipe    bool
	Migrate bool
	Logger  logur.Logger
}

func OpenDatabase(params DBParams) (*database.Service, error) {
	params.Logger = logur.WithFields(params.Logger, map[string]interface{}{util.KeyService: "config"})

	// load from config
	host := "localhost"
	port := 5432
	user := util.AppName
	password := util.AppName
	dbname := util.AppName

	template := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	url := fmt.Sprintf(template, host, port, user, password, dbname)

	db, err := sqlx.Open("pgx", url)
	if err != nil {
		return nil, errors.Wrap(err, "error opening config database")
	}

	svc := database.NewService(params.Debug, db, params.Logger)

	if params.Wipe {
		err = database.DBWipe(svc, params.Logger)
		if err != nil {
			return nil, errors.Wrap(err, "error applying initial schema")
		}
	}

	if params.Migrate {
		err = database.Migrate(svc)
		if err != nil {
			return nil, errors.Wrap(err, "error applying database migrations")
		}
	}

	return svc, nil
}
