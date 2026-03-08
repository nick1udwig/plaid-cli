package cmd

import "github.com/spf13/cobra"

const statementsDocPath = "docs/plaid/api/products/statements/index.md"

func newStatementsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "statements",
		Short: "Statements product commands",
		Long:  "Statements commands for listing, downloading, and refreshing bank statements on linked Items.",
	}

	cmd.AddCommand(newStatementsListCmd())
	cmd.AddCommand(newStatementsDownloadCmd())
	cmd.AddCommand(newStatementsRefreshCmd())

	return cmd
}

func newStatementsListCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /statements/list",
		Long:  "Capability: read. Lists the bank statements available for a linked Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"access_token": "<access-token>"}
			if handled, err := maybeWriteInfo(cmd, info, statementsDocPath, template); handled || err != nil {
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
			resp, err := client.Call(ctx, "/statements/list", body)
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

func newStatementsDownloadCmd() *cobra.Command {
	var itemID, accessToken, statementID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var outputFlags *binaryOutputFlags

	cmd := &cobra.Command{
		Use:   "download",
		Short: "Call /statements/download",
		Long:  "Capability: read. Downloads a single bank statement PDF.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"statement_id": "<statement-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, statementsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "statement-id", statementID, "statement_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--statement-id": {"statement_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.CallBytes(ctx, "/statements/download", body)
			if err != nil {
				return err
			}
			return writeBinaryOutput(cmd, outputFlags.out, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	outputFlags = bindBinaryOutputFlag(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringVar(&statementID, "statement-id", "", "Statement identifier returned by `plaid statements list`")
	return cmd
}

func newStatementsRefreshCmd() *cobra.Command {
	var itemID, accessToken, startDate, endDate string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Call /statements/refresh",
		Long:  "Capability: write. Triggers on-demand statement extraction for a date range.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"start_date":   "2025-01-01",
				"end_date":     "2025-03-31",
			}
			if handled, err := maybeWriteInfo(cmd, info, statementsDocPath, template); handled || err != nil {
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
			if err := requireBodyFields(body, map[string][]string{
				"--start-date": {"start_date"},
				"--end-date":   {"end_date"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/statements/refresh", body)
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
	return cmd
}
