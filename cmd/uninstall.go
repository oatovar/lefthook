package cmd

import (
	"github.com/spf13/cobra"

	"github.com/evilmartians/lefthook/internal/lefthook"
)

func newUninstallCmd(opts *lefthook.Options) *cobra.Command {
	args := lefthook.UninstallArgs{}

	uninstallCmd := cobra.Command{
		Use:   "uninstall",
		Short: "Revert install command",
		RunE: func(cmd *cobra.Command, _args []string) error {
			return lefthook.Uninstall(opts, &args)
		},
	}

	uninstallCmd.Flags().BoolVarP(
		&args.KeepConfiguration, "keep-config", "k", false,
		"keep configuration files and source directories present",
	)

	uninstallCmd.Flags().BoolVarP(
		&args.Aggressive, "aggressive", "a", false,
		"remove all git hooks even not lefthook-related",
	)

	return &uninstallCmd
}
