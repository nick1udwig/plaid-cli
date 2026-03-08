package cmd

import "github.com/spf13/cobra"

func newTransactionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Transactions product commands",
		Long:  "Transactions commands. Most are read-only; `refresh` triggers a write-style refresh request.",
	}
	cmd.AddCommand(newTransactionsGetCmd())
	cmd.AddCommand(newTransactionsSyncCmd())
	cmd.AddCommand(newTransactionsRecurringGetCmd())
	cmd.AddCommand(newTransactionsRefreshCmd())
	cmd.AddCommand(newTransactionsCategoriesGetCmd())
	return cmd
}

func newTransactionsGetCmd() *cobra.Command {
	var itemID, accessToken, startDate, endDate string
	var accountIDs []string
	var count, offset int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transactions/get",
		Long:  "Capability: read. Retrieves transactions within a date range.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"start_date":   "2026-01-01",
				"end_date":     "2026-02-01",
			}
			if len(accountIDs) > 0 || count != 100 || offset != 0 {
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
				if len(options) > 0 {
					template["options"] = options
				}
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/transactions/index.md", template); handled || err != nil {
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
			if err := requireBodyFields(body, map[string][]string{
				"--start-date": {"start_date"},
				"--end-date":   {"end_date"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transactions/get", body)
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
	cmd.Flags().StringVar(&startDate, "start-date", "", "Start date in YYYY-MM-DD")
	cmd.Flags().StringVar(&endDate, "end-date", "", "End date in YYYY-MM-DD")
	cmd.Flags().StringSliceVar(&accountIDs, "account-id", nil, "Account ID to filter by (repeatable)")
	cmd.Flags().IntVar(&count, "count", 100, "Maximum number of transactions to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Transaction list offset")
	return cmd
}

func newTransactionsSyncCmd() *cobra.Command {
	var itemID, accessToken, cursor, accountID, pfcVersion string
	var count, daysRequested int
	var includeOriginalDescription bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Call /transactions/sync",
		Long:  "Capability: read. Incrementally syncs transaction changes using Plaid cursors.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"count":        count,
			}
			if cursor != "" {
				template["cursor"] = cursor
			}
			options := map[string]any{}
			if accountID != "" {
				options["account_id"] = accountID
			}
			if includeOriginalDescription {
				options["include_original_description"] = true
			}
			if pfcVersion != "" {
				options["personal_finance_category_version"] = pfcVersion
			}
			if daysRequested != 0 {
				options["days_requested"] = daysRequested
			}
			if len(options) > 0 {
				template["options"] = options
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/transactions/index.md", template); handled || err != nil {
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
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "account-id", accountID, "options", "account_id"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "include-original-description", includeOriginalDescription, "options", "include_original_description"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "pfc-version", pfcVersion, "options", "personal_finance_category_version"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "days-requested", daysRequested, "options", "days_requested"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transactions/sync", body)
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
	cmd.Flags().StringVar(&cursor, "cursor", "", "Previously saved sync cursor")
	cmd.Flags().IntVar(&count, "count", 100, "Number of transaction updates to fetch")
	cmd.Flags().StringVar(&accountID, "account-id", "", "Single account ID to scope the sync stream to")
	cmd.Flags().BoolVar(&includeOriginalDescription, "include-original-description", false, "Include original institution transaction descriptions")
	cmd.Flags().StringVar(&pfcVersion, "pfc-version", "", "Personal finance category version: v1 or v2")
	cmd.Flags().IntVar(&daysRequested, "days-requested", 0, "Days of transaction history to request when initializing Transactions")
	return cmd
}

func newTransactionsRecurringGetCmd() *cobra.Command {
	var itemID, accessToken, pfcVersion string
	var accountIDs []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "recurring-get",
		Short: "Call /transactions/recurring/get",
		Long:  "Capability: read. Retrieves recurring inflow and outflow streams inferred by Plaid.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"access_token": "<access-token>"}
			if len(accountIDs) > 0 {
				template["account_ids"] = accountIDs
			}
			if pfcVersion != "" {
				template["options"] = map[string]any{"personal_finance_category_version": pfcVersion}
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/transactions/index.md", template); handled || err != nil {
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
			if err := applyStringSliceFlag(cmd, body, "account-id", accountIDs, "account_ids"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "pfc-version", pfcVersion, "options", "personal_finance_category_version"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transactions/recurring/get", body)
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
	cmd.Flags().StringVar(&pfcVersion, "pfc-version", "", "Personal finance category version: v1 or v2")
	return cmd
}

func newTransactionsRefreshCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Call /transactions/refresh",
		Long:  "Capability: write. Triggers an on-demand transactions refresh for an Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/transactions/index.md", map[string]any{
				"access_token": "<access-token>",
			}); handled || err != nil {
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
			resp, err := client.Call(ctx, "/transactions/refresh", body)
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

func newTransactionsCategoriesGetCmd() *cobra.Command {
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "categories-get",
		Short: "Call /categories/get",
		Long:  "Capability: read. Retrieves Plaid transaction categories.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/transactions/index.md", map[string]any{}); handled || err != nil {
				return err
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/categories/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	return cmd
}
