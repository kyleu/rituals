package cli

import (
	"github.com/spf13/cobra"
)

func NewVersionCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Displays the software version",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			println(version)
			return nil
		},
	}

	return cmd
}
