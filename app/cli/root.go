package cli

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kyleu/rituals.dev/app/controllers/routes"

	"github.com/kyleu/rituals.dev/app/invitation"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/gorilla/handlers"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/user"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/spf13/cobra"
	log "logur.dev/logur"
)

var verbose bool
var redir string
var addr string
var port uint16
var authEnabled bool

// Configure configures a root command.
func Configure(version string, commitHash string) cobra.Command {
	rootCmd := cobra.Command{
		Use:   util.AppName,
		Short: "Command line interface for " + util.AppName + ", the database user interface",
		RunE: func(cmd *cobra.Command, _ []string) error {
			info, err := InitApp(version, commitHash)
			if err != nil {
				return errors.Wrap(err, "error initializing application")
			}

			return MakeServer(info, addr, port)
		},
	}

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&redir, "redir", "r", "http://localhost:6660", "redirect for signin, defaults to localhost")
	flags.StringVarP(&addr, "address", "a", "127.0.0.1", "interface address to listen on")
	flags.Uint16VarP(&port, "port", "p", 6660, "port for http server to listen on")
	flags.BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	flags.BoolVar(&authEnabled, "auth", true, "enable authentication")

	return rootCmd
}

func InitApp(version string, commitHash string) (*config.AppInfo, error) {
	_ = os.Setenv("TZ", "UTC")

	logger := initLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose, "version": version, "commit": commitHash})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	db, err := config.OpenDatabase(logger)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database pool")
	}

	actionService := action.NewService(db, logger)
	userSvc := user.NewService(actionService, db, logger)
	authSvc := auth.NewService(authEnabled, redir, actionService, db, logger, userSvc)
	invitationSvc := invitation.NewService(actionService, db, logger)
	teamSvc := team.NewService(actionService, db, logger)
	sprintSvc := sprint.NewService(actionService, db, logger)
	estimateSvc := estimate.NewService(actionService, db, logger)
	standupSvc := standup.NewService(actionService, db, logger)
	retroSvc := retro.NewService(actionService, db, logger)
	socketSvc := socket.NewService(logger, actionService, userSvc, authSvc, teamSvc, sprintSvc, estimateSvc, standupSvc, retroSvc)

	ai := config.AppInfo{
		Debug:      verbose,
		Version:    version,
		Commit:     commitHash,
		Logger:     logger,
		User:       userSvc,
		Auth:       authSvc,
		Action:     actionService,
		Invitation: invitationSvc,
		Team:       teamSvc,
		Sprint:     sprintSvc,
		Estimate:   estimateSvc,
		Standup:    standupSvc,
		Retro:      retroSvc,
		Socket:     socketSvc,
		Database:   db,
	}

	return &ai, nil
}

func MakeServer(info *config.AppInfo, address string, port uint16) error {
	r, err := routes.BuildRouter(info)
	if err != nil {
		return errors.WithMessage(err, "unable to construct routes")
	}
	var msg = fmt.Sprintf("%v is starting on [%v:%v]", util.AppName, address, port)
	if info.Debug {
		msg += " (verbose)"
	}
	info.Logger.Info(msg, map[string]interface{}{"address": address, "port": port})
	err = http.ListenAndServe(fmt.Sprint(address, ":", port), handlers.CORS()(r))
	return errors.Wrap(err, "unable to run http server")
}
