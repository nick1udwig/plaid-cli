package cmd

import "github.com/spf13/cobra"

func newTransferCapabilitiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "capabilities",
		Short: "Transfer capability commands",
		Long:  "Read-only transfer capability commands for determining network eligibility.",
	}

	cmd.AddCommand(newTransferCapabilitiesGetCmd())

	return cmd
}

func newTransferCapabilitiesGetCmd() *cobra.Command {
	var itemID, accessToken, accountID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/capabilities/get",
		Long:  "Capability: read. Retrieves Transfer network eligibility information for an account.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"account_id":   "<account-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferAccountLinkingDocPath, template); handled || err != nil {
				return err
			}

			store, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			if _, err := populateTransferAccess(cmd, store, body, itemID, accessToken, accountID); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/capabilities/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id to evaluate")
	return cmd
}
