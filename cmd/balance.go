package cmd

import "github.com/spf13/cobra"

func newBalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance",
		Short: "Retrieve real-time balances",
		Long:  "Read-only balance commands.",
	}
	cmd.AddCommand(newBalanceGetCmd())
	return cmd
}

func newBalanceGetCmd() *cobra.Command {
	var itemID, accessToken string
	var accountIDs []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /accounts/balance/get",
		Long:  "Capability: read. Retrieves real-time balances for accounts on a linked Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"access_token": "<access-token>"}
			if len(accountIDs) > 0 {
				template["options"] = map[string]any{"account_ids": accountIDs}
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/signal/index.md", template); handled || err != nil {
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
			if _, err := populateAccessToken(cmd, store, body, itemID, accessToken); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "account-id", accountIDs, "options", "account_ids"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/accounts/balance/get", body)
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
	cmd.Flags().StringSliceVar(&accountIDs, "account-id", nil, "Account ID to filter by (repeatable)")
	return cmd
}
