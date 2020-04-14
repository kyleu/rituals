package cli

import (
	"fmt"

	"github.com/kyleu/rituals.dev/internal/app/conn/output"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/internal/app/conn"
	"github.com/spf13/cobra"
)

func NewSandboxCommand(version string, commitHash string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sandbox",
		Aliases: []string{"x"},
		Short:   "Runs an internal test",
		RunE: func(cmd *cobra.Command, _ []string) error {
			info, err := InitApp(version, commitHash)
			if err != nil {
				return errors.WithStack(errors.Wrap(err, "error initializing application"))
			}
			connection, ms, err := info.ConfigService.GetConnection("")
			if err != nil {
				return errors.WithStack(errors.Wrap(err, "error opening connection"))
			}
			rs, err := conn.RunQueryNoTx(info.Logger, connection, ms, conn.Adhoc(""))
			if err != nil {
				return errors.WithStack(errors.Wrap(err, "error retrieving result"))
			}
			out, err := output.OutputFor(rs, "table")
			if err != nil {
				return errors.WithStack(errors.Wrap(err, "error formatting output"))
			}
			fmt.Println(out)
			return nil
		},
	}

	return cmd
}
