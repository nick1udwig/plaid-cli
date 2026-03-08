package cmd

import "github.com/spf13/cobra"

func newTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Transfer product commands",
		Long:  "Transfer commands for authorizing, creating, reading, refunding, and reconciling money movement operations.",
	}

	cmd.AddCommand(newTransferCreateCmd())
	cmd.AddCommand(newTransferGetCmd())
	cmd.AddCommand(newTransferListCmd())
	cmd.AddCommand(newTransferCancelCmd())
	cmd.AddCommand(newTransferAuthorizationCmd())
	cmd.AddCommand(newTransferEventCmd())
	cmd.AddCommand(newTransferRefundCmd())
	cmd.AddCommand(newTransferSweepCmd())
	cmd.AddCommand(newTransferRecurringCmd())
	cmd.AddCommand(newTransferMetricsCmd())
	cmd.AddCommand(newTransferCapabilitiesCmd())
	cmd.AddCommand(newTransferIntentCmd())

	return cmd
}
