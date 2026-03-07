package cmd

import "github.com/spf13/cobra"

func newLinkCmd(opts *Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link",
		Short: "Create and manage Plaid Link sessions",
	}

	cmd.AddCommand(newLinkConnectCmd(opts))

	return cmd
}
