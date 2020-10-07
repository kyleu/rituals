package cli

import (
	"os"
	"strings"

	"github.com/kyleu/npn/npnasset"

	"github.com/kyleu/rituals.dev/app/gql"
	"github.com/kyleu/rituals.dev/gen/query"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/app/routing"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/spf13/cobra"
	log "logur.dev/logur"
)

var debugSQL bool
var verbose bool
var redir string
var addr string
var port uint16
var authEnabled bool
var wipeDatabase bool

// Configure configures a root command.
func Configure() cobra.Command {
	npncore.AppKey = "rituals"
	npncore.AppName = "rituals.dev"
	npnasset.AssetBase = "../npn/" + npnasset.AssetBase

	rootCmd := cobra.Command{
		Use:   npncore.AppName,
		Short: "Command line interface for " + npncore.AppName,
		RunE: func(cmd *cobra.Command, _ []string) error {
			info, err := InitApp()
			if err != nil {
				return errors.Wrap(err, "error initializing application")
			}

			r, err := routing.BuildRouter(info)
			if err != nil {
				return errors.WithMessage(err, "unable to construct routes")
			}
			actualPort, err := npnweb.MakeServer(info, r, addr, port)
			if actualPort > 0 {
				port = actualPort
			}
			return err
		},
	}

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&redir, "redir", "r", "http://localhost:6660", "redirect url for signin, defaults to localhost")
	flags.StringVarP(&addr, "address", "a", "127.0.0.1", "interface address to listen on")
	flags.Uint16VarP(&port, "port", "p", 6660, "port for http server to listen on")
	flags.BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	flags.BoolVar(&authEnabled, "auth", true, "enable authentication")
	flags.BoolVarP(&wipeDatabase, "wipe", "w", false, "wipe and rebuild the database")

	return rootCmd
}

func InitApp() (npnweb.AppInfo, error) {
	_ = os.Setenv("TZ", "UTC")

	npnweb.IconContent = `<span data-uk-icon="icon: git-fork; ratio: 1.6"></span>`

	logger := npncore.InitLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

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
