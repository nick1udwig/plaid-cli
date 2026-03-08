package cmd

import "github.com/spf13/cobra"

func newProcessorTransactionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Processor transactions commands",
		Long:  "Transactions commands for processor tokens.",
	}
	cmd.AddCommand(newProcessorTransactionsGetCmd())
	cmd.AddCommand(newProcessorTransactionsSyncCmd())
	cmd.AddCommand(newProcessorTransactionsRecurringGetCmd())
	cmd.AddCommand(newProcessorTokenOnlyCmd(
		"refresh",
		"Call /processor/transactions/refresh",
		"Capability: write. Triggers an on-demand transactions refresh for a processor token.",
		"/processor/transactions/refresh",
	))
	return cmd
}

func newProcessorTransactionsGetCmd() *cobra.Command {
	var processorToken, startDate, endDate string
	var count, offset int
	var includeOriginalDescription bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /processor/transactions/get",
		Long:  "Capability: read. Retrieves transactions within a date range for a processor token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token": "<processor-token>",
				"start_date":      "2026-01-01",
				"end_date":        "2026-02-01",
			}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "start-date", startDate, "start_date"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-date", endDate, "end_date"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "options", "count"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "offset", offset, "options", "offset"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "include-original-description", includeOriginalDescription, "options", "include_original_description"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token": {"processor_token"},
				"--start-date":      {"start_date"},
				"--end-date":        {"end_date"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/transactions/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Start date in YYYY-MM-DD")
	cmd.Flags().StringVar(&endDate, "end-date", "", "End date in YYYY-MM-DD")
	cmd.Flags().IntVar(&count, "count", 100, "Maximum number of transactions to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Transaction list offset")
	cmd.Flags().BoolVar(&includeOriginalDescription, "include-original-description", false, "Include original institution transaction descriptions")
	return cmd
}

func newProcessorTransactionsSyncCmd() *cobra.Command {
	var processorToken, cursor, accountID, pfcVersion string
	var count, daysRequested int
	var includeOriginalDescription bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Call /processor/transactions/sync",
		Long:  "Capability: read. Incrementally syncs transaction changes for a processor token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token": "<processor-token>",
				"count":           count,
			}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
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
			if err := applyOptionalIntFlag(cmd, body, "days-requested", daysRequested, "options", "days_requested"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token": {"processor_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/transactions/sync", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Previously saved sync cursor")
	cmd.Flags().IntVar(&count, "count", 100, "Number of transaction updates to fetch")
	cmd.Flags().StringVar(&accountID, "account-id", "", "Single account ID to scope the sync stream to")
	cmd.Flags().BoolVar(&includeOriginalDescription, "include-original-description", false, "Include original institution transaction descriptions")
	cmd.Flags().StringVar(&pfcVersion, "pfc-version", "", "Personal finance category version: v1 or v2")
	cmd.Flags().IntVar(&daysRequested, "days-requested", 0, "Days of transaction history to request when initializing Transactions")
	return cmd
}

func newProcessorTransactionsRecurringGetCmd() *cobra.Command {
	var processorToken, pfcVersion string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "recurring-get",
		Short: "Call /processor/transactions/recurring/get",
		Long:  "Capability: read. Retrieves recurring inflow and outflow streams for a processor token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"processor_token": "<processor-token>"}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "pfc-version", pfcVersion, "options", "personal_finance_category_version"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token": {"processor_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/transactions/recurring/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringVar(&pfcVersion, "pfc-version", "", "Personal finance category version: v1 or v2")
	return cmd
}
