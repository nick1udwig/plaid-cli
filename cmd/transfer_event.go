package cmd

import "github.com/spf13/cobra"

func newTransferEventCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "event",
		Short: "Transfer event commands",
		Long:  "Read-only transfer event commands for listing and syncing transfer events.",
	}

	cmd.AddCommand(newTransferEventListCmd())
	cmd.AddCommand(newTransferEventSyncCmd())

	return cmd
}

func newTransferEventListCmd() *cobra.Command {
	var startDate, endDate, transferID, accountID, transferType string
	var sweepID, originatorClientID, fundingAccountID string
	var eventTypes []string
	var count, offset int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /transfer/event/list",
		Long:  "Capability: read. Lists transfer events matching the provided filter criteria.",
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
			if err := applyStringFlag(cmd, body, "transfer-id", transferID, "transfer_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "account-id", accountID, "account_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "transfer-type", transferType, "transfer_type"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "event-type", eventTypes, "event_types"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "sweep-id", sweepID, "sweep_id"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "offset", offset, "offset"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "funding-account-id", fundingAccountID, "funding_account_id"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/event/list", body)
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
	cmd.Flags().StringVar(&transferID, "transfer-id", "", "Transfer ID filter")
	cmd.Flags().StringVar(&accountID, "account-id", "", "Account ID filter")
	cmd.Flags().StringVar(&transferType, "transfer-type", "", "Transfer type filter: debit or credit")
	cmd.Flags().StringSliceVar(&eventTypes, "event-type", nil, "Transfer event type filter (repeatable)")
	cmd.Flags().StringVar(&sweepID, "sweep-id", "", "Sweep ID filter")
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of events to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Transfer event list offset")
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID filter")
	cmd.Flags().StringVar(&fundingAccountID, "funding-account-id", "", "Funding account ID filter")
	return cmd
}

func newTransferEventSyncCmd() *cobra.Command {
	var afterID, count int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Call /transfer/event/sync",
		Long:  "Capability: read. Incrementally syncs transfer events after the specified event ID.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"after_id": 0,
				"count":    count,
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
			if err := applyIntFlag(cmd, body, "after-id", afterID, "after_id"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--after-id": {"after_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/event/sync", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().IntVar(&afterID, "after-id", 0, "Largest previously seen transfer event_id, or 0 initially")
	cmd.Flags().IntVar(&count, "count", 100, "Maximum number of events to return")
	return cmd
}
