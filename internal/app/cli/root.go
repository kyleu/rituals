package cli

import (
	"emperror.dev/emperror"
	"emperror.dev/errors"
	logurhandler "emperror.dev/handler/logur"
	"github.com/kyleu/rituals.dev/internal/app/config"
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

// Configure configures a root command.
func Configure(version string, commitHash string) cobra.Command {
	rootCmd := cobra.Command{
		Use:   util.AppName,
		Short: "Command line interface for " + util.AppName + ", the database user interface",
	}

	flags := rootCmd.PersistentFlags()
	flags.BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(
		NewServerCommand(version, commitHash),
		NewSandboxCommand(version, commitHash),
		NewVersionCommand(version),
	)

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
		User:     userSvc,
		Invite:   inviteSvc,
		Estimate: estimateSvc,
		Standup:  standupSvc,
		Retro:    retroSvc,
		Socket:   socketSvc,
	}

	return &ai, nil
}
