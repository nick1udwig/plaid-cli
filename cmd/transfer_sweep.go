package cmd

import "github.com/spf13/cobra"

func newTransferSweepCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sweep",
		Short: "Transfer sweep commands",
		Long:  "Read-only sweep commands for retrieving sweep records associated with Transfer.",
	}

	cmd.AddCommand(newTransferSweepGetCmd())
	cmd.AddCommand(newTransferSweepListCmd())

	return cmd
}

func newTransferSweepGetCmd() *cobra.Command {
	var sweepID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/sweep/get",
		Long:  "Capability: read. Retrieves a transfer sweep by sweep_id.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"sweep_id": "<sweep-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferReadingDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "sweep-id", sweepID, "sweep_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--sweep-id": {"sweep_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/sweep/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&sweepID, "sweep-id", "", "Sweep ID to retrieve")
	return cmd
}

func newTransferSweepListCmd() *cobra.Command {
	var startDate, endDate, amount, status, originatorClientID string
	var fundingAccountID, transferID, trigger string
	var count, offset int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /transfer/sweep/list",
		Long:  "Capability: read. Lists sweeps matching the provided filter criteria.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"count":  count,
				"offset": offset,
			}
			if handled, err := maybeWriteInfo(cmd, info, transferReadingDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "start-date", startDate, "start_date"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-date", endDate, "end_date"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "offset", offset, "offset"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "status", status, "status"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "funding-account-id", fundingAccountID, "funding_account_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "transfer-id", transferID, "transfer_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "trigger", trigger, "trigger"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/sweep/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&startDate, "start-date", "", "RFC3339 start timestamp")
	cmd.Flags().StringVar(&endDate, "end-date", "", "RFC3339 end timestamp")
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of sweeps to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Sweep list offset")
	cmd.Flags().StringVar(&amount, "amount", "", "Signed sweep amount filter")
	cmd.Flags().StringVar(&status, "status", "", "Sweep status filter")
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID filter")
	cmd.Flags().StringVar(&fundingAccountID, "funding-account-id", "", "Funding account ID filter")
	cmd.Flags().StringVar(&transferID, "transfer-id", "", "Transfer ID filter")
	cmd.Flags().StringVar(&trigger, "trigger", "", "Sweep trigger filter")
	return cmd
}
