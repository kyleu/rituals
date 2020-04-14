package cli

import (
	"emperror.dev/emperror"
	"emperror.dev/errors"
	logurhandler "emperror.dev/handler/logur"
	"github.com/kyleu/rituals.dev/internal/app/config"
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

	cfg, err := config.NewService(logger)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating config service"))
	}

	ai := config.AppInfo{
		Debug:         verbose,
		Version:       version,
		CommitHash:    commitHash,
		Logger:        logger,
		ErrorHandler:  handler,
		ConfigService: cfg,
	}

	return &ai, nil
}
