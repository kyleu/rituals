package cli

import (
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app/routing"

	"emperror.dev/errors"
	"github.com/spf13/cobra"
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
