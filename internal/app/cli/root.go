package cli

import (
	"fmt"
	"net/http"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	logurhandler "emperror.dev/handler/logur"
	"github.com/gorilla/handlers"
	"github.com/kyleu/rituals.dev/internal/app/config"
	"github.com/kyleu/rituals.dev/internal/app/controllers"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"github.com/kyleu/rituals.dev/internal/app/invite"
	"github.com/kyleu/rituals.dev/internal/app/retro"
	"github.com/kyleu/rituals.dev/internal/app/socket"
	"github.com/kyleu/rituals.dev/internal/app/standup"
	"github.com/kyleu/rituals.dev/internal/app/user"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"github.com/spf13/cobra"
	"logur.dev/logur"
)

var verbose bool
var port uint16
var addr string

// Configure configures a root command.
func Configure(version string, commitHash string) cobra.Command {
	rootCmd := cobra.Command{
		Use:   util.AppName,
		Short: "Command line interface for " + util.AppName + ", the database user interface",
		RunE: func(cmd *cobra.Command, _ []string) error {
			info, err := InitApp(version, commitHash)
			if err != nil {
				return errors.WithStack(errors.Wrap(err, "error initializing application"))
			}

			return MakeServer(info, addr, port)
		},
	}

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&addr, "address", "a", "127.0.0.1", "interface address to listen on")
	flags.Uint16VarP(&port, "port", "p", 6660, "port for http server to listen on")
	flags.BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	return rootCmd
}

func InitApp(version string, commitHash string) (*config.AppInfo, error) {
	logger := util.InitLogging(verbose)
	logger = logur.WithFields(logger, map[string]interface{}{"debug": verbose, "version": version, "commit": commitHash})

	errorHandler := logurhandler.New(logger)
	defer emperror.HandleRecover(errorHandler)

	handler := emperror.WithDetails(util.AppErrorHandler{Logger: logger}, "key", "value")

	db, err := config.OpenDatabase(logger)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating config service"))
	}

	userSvc := user.NewUserService(db, logger)
	inviteSvc := invite.NewInviteService(db, logger)
	estimateSvc := estimate.NewEstimateService(db, logger)
	standupSvc := standup.NewStandupService(db, logger)
	retroSvc := retro.NewRetroService(db, logger)
	socketSvc := socket.NewSocketService(logger, &userSvc, &estimateSvc)

	ai := config.AppInfo{
		Debug:    verbose,
		Version:  version,
		Commit:   commitHash,
		Logger:   logger,
		Errors:   handler,
		User:     &userSvc,
		Invite:   &inviteSvc,
		Estimate: &estimateSvc,
		Standup:  &standupSvc,
		Retro:    &retroSvc,
		Socket:   &socketSvc,
	}

	return &ai, nil
}

func MakeServer(info *config.AppInfo, address string, port uint16) error {
	routes, err := controllers.BuildRouter(info)
	if err != nil {
		return errors.WithStack(errors.WithMessage(err, "unable to construct routes"))
	}
	var msg = fmt.Sprintf("%v is starting on [%v:%v]", util.AppName, address, port)
	if info.Debug {
		msg += " (verbose)"
	}
	info.Logger.Info(msg, map[string]interface{}{"address": address, "port": port})
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), handlers.CORS()(routes))
	return errors.WithStack(errors.Wrap(err, "unable to run http server"))
}
