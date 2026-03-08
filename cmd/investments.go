package cmd

import "github.com/spf13/cobra"

const investmentsDocPath = "docs/plaid/api/products/investments/index.md"

func newInvestmentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "investments",
		Short: "Investments product commands",
		Long:  "Investments commands. Holdings and transactions are read-only; refresh triggers a write-style extraction request.",
	}

	cmd.AddCommand(newInvestmentsHoldingsGetCmd())
	cmd.AddCommand(newInvestmentsTransactionsGetCmd())
	cmd.AddCommand(newInvestmentsRefreshCmd())

	return cmd
}

func newInvestmentsHoldingsGetCmd() *cobra.Command {
	var itemID, accessToken string
	var accountIDs []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "holdings-get",
		Short: "Call /investments/holdings/get",
		Long:  "Capability: read. Retrieves investment account holdings and security metadata.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"access_token": "<access-token>"}
			if len(accountIDs) > 0 {
				template["options"] = map[string]any{"account_ids": accountIDs}
			}
			if handled, err := maybeWriteInfo(cmd, info, investmentsDocPath, template); handled || err != nil {
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
			resp, err := client.Call(ctx, "/investments/holdings/get", body)
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

func newInvestmentsTransactionsGetCmd() *cobra.Command {
	var itemID, accessToken, startDate, endDate string
	var accountIDs []string
	var count, offset int
	var asyncUpdate bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "transactions-get",
		Short: "Call /investments/transactions/get",
		Long:  "Capability: read. Retrieves investment transactions for a date range.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"start_date":   "2026-01-01",
				"end_date":     "2026-03-01",
			}
			options := map[string]any{}
			if len(accountIDs) > 0 {
				options["account_ids"] = accountIDs
			}
			if count != 100 {
				options["count"] = count
			}
			if offset != 0 {
				options["offset"] = offset
			}
			if asyncUpdate {
				options["async_update"] = true
			}
			if len(options) > 0 {
				template["options"] = options
			}
			if handled, err := maybeWriteInfo(cmd, info, investmentsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "start-date", startDate, "start_date"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-date", endDate, "end_date"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "account-id", accountIDs, "options", "account_ids"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "options", "count"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "offset", offset, "options", "offset"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "async-update", asyncUpdate, "options", "async_update"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--start-date": {"start_date"},
				"--end-date":   {"end_date"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/investments/transactions/get", body)
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
	cmd.Flags().StringVar(&startDate, "start-date", "", "Earliest transaction date in YYYY-MM-DD")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Latest transaction date in YYYY-MM-DD")
	cmd.Flags().StringSliceVar(&accountIDs, "account-id", nil, "Investment account_id filter (repeatable)")
	cmd.Flags().IntVar(&count, "count", 100, "Number of transactions to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Transaction pagination offset")
	cmd.Flags().BoolVar(&asyncUpdate, "async-update", false, "Allow Plaid to initialize investment extraction asynchronously if needed")
	return cmd
}

func newInvestmentsRefreshCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Call /investments/refresh",
		Long:  "Capability: write. Triggers an on-demand investment holdings and transactions refresh for an Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"access_token": "<access-token>"}
			if handled, err := maybeWriteInfo(cmd, info, investmentsDocPath, template); handled || err != nil {
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

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/investments/refresh", body)
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
	return cmd
}
