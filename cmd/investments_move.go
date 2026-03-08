package cmd

import "github.com/spf13/cobra"

const investmentsMoveDocPath = "docs/plaid/api/products/investments-move/index.md"

func newInvestmentsMoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "investments-move",
		Short: "Investments Move commands",
		Long:  "Investments Move commands for retrieving authorization data needed for holdings transfers.",
	}

	cmd.AddCommand(newInvestmentsMoveAuthGetCmd())

	return cmd
}

func newInvestmentsMoveAuthGetCmd() *cobra.Command {
	var itemID, accessToken string
	var accountIDs []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "auth-get",
		Short: "Call /investments/auth/get",
		Long:  "Capability: read. Retrieves investments authorization data for ACATS or ATON transfer workflows.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"access_token": "<access-token>"}
			if len(accountIDs) > 0 {
				template["options"] = map[string]any{"account_ids": accountIDs}
			}
			if handled, err := maybeWriteInfo(cmd, info, investmentsMoveDocPath, template); handled || err != nil {
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
			resp, err := client.Call(ctx, "/investments/auth/get", body)
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
	cmd.Flags().StringSliceVar(&accountIDs, "account-id", nil, "Investment account_id filter (repeatable)")
	return cmd
}
