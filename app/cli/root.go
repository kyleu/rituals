package cli

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kyleu/rituals.dev/app/model/action"
	"github.com/kyleu/rituals.dev/app/model/auth"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/model/user"
	"github.com/kyleu/rituals.dev/app/socket"

	"github.com/kyleu/rituals.dev/app/web/routes"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/gorilla/handlers"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/util"
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
		Use:   util.AppName,
		Short: "Command line interface for " + util.AppName,
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
	flags.StringVarP(&addr, "address", "a", "127.0.0.1", "interface address to listen on")
	flags.Uint16VarP(&port, "port", "p", 6660, "port for http server to listen on")
	flags.BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	flags.BoolVar(&authEnabled, "auth", true, "enable authentication")
	flags.BoolVarP(&wipeDatabase, "wipe", "w", false, "wipe and rebuild the database")

	return rootCmd
}

func InitApp(version string, commitHash string) (*config.AppInfo, error) {
	_ = os.Setenv("TZ", "UTC")

	logger := initLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose, "version": version, "commit": commitHash})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	ai, err := initAppInfo(logger, version, commitHash)
	if err != nil {
		return nil, err
	}

	return ai, nil
}

func initAppInfo(logger log.Logger, version string, commitHash string) (*config.AppInfo, error) {
	db, err := config.OpenDatabase(config.DBParams{Debug: verbose && debugSQL, Wipe: wipeDatabase, Migrate: true, Logger: logger})
	if err != nil {
		return nil, errors.Wrap(err, "error opening database pool")
	}
	actionService := action.NewService(db, logger)
	commentService := comment.NewService(actionService, db, logger)
	userSvc := user.NewService(actionService, db, logger)
	authSvc := auth.NewService(authEnabled, redir, actionService, db, logger, userSvc)
	teamSvc := team.NewService(actionService, userSvc, commentService, db, logger)
	sprintSvc := sprint.NewService(actionService, userSvc, commentService, db, logger)
	estimateSvc := estimate.NewService(actionService, userSvc, commentService, db, logger)
	standupSvc := standup.NewService(actionService, userSvc, commentService, db, logger)
	retroSvc := retro.NewService(actionService, userSvc, commentService, db, logger)
	socketSvc := socket.NewService(logger, actionService, userSvc, commentService, authSvc, teamSvc, sprintSvc, estimateSvc, standupSvc, retroSvc)

	return &config.AppInfo{
		Debug:    verbose,
		Version:  version,
		Commit:   commitHash,
		Logger:   logger,
		User:     userSvc,
		Comment:  commentService,
		Auth:     authSvc,
		Action:   actionService,
		Team:     teamSvc,
		Sprint:   sprintSvc,
		Estimate: estimateSvc,
		Standup:  standupSvc,
		Retro:    retroSvc,
		Socket:   socketSvc,
		Database: db,
	}, nil
}

func MakeServer(info *config.AppInfo, address string, port uint16) error {
	r, err := routes.BuildRouter(info)
	if err != nil {
		return errors.WithMessage(err, "unable to construct routes")
	}
	var msg = "%v is starting on [%v:%v]"
	if info.Debug {
		msg += " (verbose)"
	}
	info.Logger.Info(fmt.Sprintf(msg, util.AppName, address, port))
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), handlers.CORS()(r))
	return errors.Wrap(err, "unable to run http server")
}
