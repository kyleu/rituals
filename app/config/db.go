package config

import (
	"fmt"
	"github.com/kyleu/rituals.dev/app/database"
	"strings"
	"time"

	"emperror.dev/errors"
	// load postgres driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/gen/queries"
	"golang.org/x/text/language"
	"logur.dev/logur"
)

func OpenDatabase(logger logur.Logger) (*database.Service, error) {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: "config"})

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

	svc := database.New(db)

	// remove when not needed
	err = dbWipe(svc, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error applying initial queries")
	}

	return svc, nil
}

func dbWipe(db *database.Service, logger logur.Logger) error {
	err := exec("reset", db, logger, func(sb *strings.Builder) { queries.ResetDatabase(sb) })
	if err != nil {
		return err
	}
	err = exec("create-types", db, logger, func(sb *strings.Builder) { queries.CreateTypes(sb) })
	if err != nil {
		return err
	}
	err = exec("create-tables", db, logger, func(sb *strings.Builder) { queries.CreateTables(sb) })
	if err != nil {
		return err
	}
	err = exec("seed-data", db, logger, func(sb *strings.Builder) { queries.SeedData(sb) })
	if err != nil {
		return err
	}

	return nil
}

func exec(name string, db *database.Service, logger logur.Logger, f func(*strings.Builder)) error {
	sb := &strings.Builder{}
	f(sb)
	sqls := strings.Split(sb.String(), ";")
	startNanos := time.Now().UnixNano()
	for _, q := range sqls {
		_, err := db.Exec(q, nil)
		if err != nil {
			return errors.Wrap(err, "cannot execute ["+name+"]")
		}
	}
	elapsed := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
	logger.Debug(fmt.Sprintf("ran initial query [%s] in [%v]", name, util.MicrosToMillis(language.AmericanEnglish, int(elapsed))))
	return nil
}
