package cli

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/app/routing"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/gorilla/handlers"
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
func Configure(version string, commitHash string) cobra.Command {
	rootCmd := cobra.Command{
		Use:   npncore.AppName,
		Short: "Command line interface for " + npncore.AppName,
		RunE: func(cmd *cobra.Command, _ []string) error {
			info, err := InitApp(version, commitHash)
			if err != nil {
				return errors.Wrap(err, "error initializing application")
			}

			return MakeServer(info, addr, port)
		},
	}

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&redir, "redir", "r", "http://localhost:6660", "redirect url for signin, defaults to localhost")
	flags.StringVarP(&addr, "address", "a", "0.0.0.0", "interface address to listen on")
	flags.Uint16VarP(&port, "port", "p", 6660, "port for http server to listen on")
	flags.BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	flags.BoolVar(&authEnabled, "auth", true, "enable authentication")
	flags.BoolVarP(&wipeDatabase, "wipe", "w", false, "wipe and rebuild the database")

	return rootCmd
}

func InitApp(version string, commitHash string) (npnweb.AppInfo, error) {
	_ = os.Setenv("TZ", "UTC")

	npncore.AppKey = "rituals"
	npncore.AppName = "rituals.dev"
	npnweb.IconContent = "<span data-uk-icon=\"icon: git-fork; ratio: 1.6\"></span>"

	logger := npncore.InitLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose, "version": version, "commit": commitHash})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	ai, err := initAppInfo(logger, version, commitHash)
	if err != nil {
		return nil, err
	}

	return ai, nil
}

func initAppInfo(logger log.Logger, version string, commitHash string) (npnweb.AppInfo, error) {
	db, err := npndatabase.OpenDatabase(npndatabase.DBParams{
		Username: npncore.AppName, Password: npncore.AppName, DBName: npncore.AppName,
		Debug: verbose && debugSQL, Wipe: wipeDatabase, Migrate: true, Logger: logger,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error opening database pool")
	}

	return app.NewService(verbose, db, authEnabled, redir, version, commitHash, logger), nil
}

func MakeServer(info npnweb.AppInfo, address string, port uint16) error {
	r, err := routing.BuildRouter(info)
	if err != nil {
		return errors.WithMessage(err, "unable to construct routes")
	}

	var msg = "%v is starting on [%v:%v]"
	if info.Debug() {
		msg += " (verbose)"
	}
	info.Logger().Info(fmt.Sprintf(msg, npncore.AppName, address, port))
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), handlers.CORS()(r))
	return errors.Wrap(err, "unable to run http server")
}
