package cli

import (
	"github.com/kyleu/npn/npnasset"
	"os"
	"strings"

	"github.com/kyleu/rituals.dev/app/gql"
	"github.com/kyleu/rituals.dev/gen/query"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	log "logur.dev/logur"
)

func init() {
	npncore.AppKey = "rituals"
	npncore.AppName = "rituals.dev"
	npnasset.AssetBase = "../npn/" + npnasset.AssetBase
}

func InitApp() (npnweb.AppInfo, error) {
	_ = os.Setenv("TZ", "UTC")

	npnweb.IconContent = `<span data-uk-icon="icon: git-fork; ratio: 1.6"></span>`

	logger := npncore.InitLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	InitCols()

	ai, err := initAppInfo(logger)
	if err != nil {
		return nil, err
	}

	err = gql.InitService(ai)
	if err != nil {
		return nil, err
	}

	return ai, nil
}

func initAppInfo(logger log.Logger) (npnweb.AppInfo, error) {
	npndatabase.InitialSchemaMigrations = npndatabase.MigrationFiles{
		{Title: "reset", F: func(sb *strings.Builder) { query.ResetDatabase(sb) }},
		{Title: "create-types", F: func(sb *strings.Builder) { query.CreateTypes(sb) }},
		{Title: "create-tables", F: func(sb *strings.Builder) { query.CreateTables(sb) }},
		{Title: "seed-data", F: func(sb *strings.Builder) { query.SeedData(sb) }},
	}

	npndatabase.DatabaseMigrations = npndatabase.MigrationFiles{
		{Title: "first-migration", F: func(sb *strings.Builder) { query.Migration1(sb) }},
	}

	db, err := npndatabase.OpenDatabase(npndatabase.DBParams{
		Username: npncore.AppName, Password: npncore.AppName, DBName: npncore.AppName,
		Debug: verbose && debugSQL, Wipe: wipeDatabase, Migrate: true, Logger: logger,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error opening database pool")
	}

	return app.NewService(verbose, db, authEnabled, redir, logger), nil
}
