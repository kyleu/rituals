package cli

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/spf13/cobra"
)

func NewSandboxCommand(version string, commitHash string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sandbox",
		Aliases: []string{"x"},
		Short:   "Runs an internal test",
		RunE: func(cmd *cobra.Command, _ []string) error {
			_, err := InitApp(version, commitHash)
			if err != nil {
				return errors.WithStack(errors.Wrap(err, "error initializing application"))
			}
			fmt.Println("ok")
			return nil
		},
	}

	return cmd
}
