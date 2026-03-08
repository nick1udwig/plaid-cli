package cmd

import "github.com/spf13/cobra"

const sandboxDocPath = "docs/plaid/api/sandbox/index.md"

func newSandboxTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Sandbox Transfer helpers",
		Long:  "Sandbox-only Transfer helpers for simulating events, sweeps, ledger movement, and test clocks.",
	}

	cmd.AddCommand(newSandboxTransferFireWebhookCmd())
	cmd.AddCommand(newSandboxTransferSimulateCmd())
	cmd.AddCommand(newSandboxTransferRefundSimulateCmd())
	cmd.AddCommand(newSandboxTransferSweepSimulateCmd())
	cmd.AddCommand(newSandboxTransferLedgerCmd())
	cmd.AddCommand(newSandboxTransferTestClockCmd())

	return cmd
}

func newSandboxTransferFireWebhookCmd() *cobra.Command {
	var webhook string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "fire-webhook",
		Short: "Call /sandbox/transfer/fire_webhook",
		Long:  "Capability: sandbox-write. Fires a TRANSFER_EVENTS_UPDATE webhook to the provided URL.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"webhook": "https://example.com/plaid/transfer",
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--webhook": {"webhook"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/fire_webhook", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive the sandbox Transfer webhook")
	return cmd
}

func newSandboxTransferSimulateCmd() *cobra.Command {
	var transferID, eventType, testClockID, webhook string
	var failureCode, achReturnCode, failureDescription string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "simulate",
		Short: "Call /sandbox/transfer/simulate",
		Long:  "Capability: sandbox-write. Simulates an asynchronous Transfer event for a transfer.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"transfer_id": "<transfer-id>",
				"event_type":  "posted",
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "transfer-id", transferID, "transfer_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "event-type", eventType, "event_type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "test-clock-id", testClockID, "test_clock_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := applyFailureReasonFlags(cmd, body, failureCode, achReturnCode, failureDescription); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--transfer-id": {"transfer_id"},
				"--event-type":  {"event_type"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/simulate", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindFailureReasonFlags(cmd, &failureCode, &achReturnCode, &failureDescription)
	cmd.Flags().StringVar(&transferID, "transfer-id", "", "Transfer ID to simulate an event for")
	cmd.Flags().StringVar(&eventType, "event-type", "", "Transfer event type to simulate")
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Sandbox Transfer test clock ID")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive a sandbox Transfer webhook")
	return cmd
}

func newSandboxTransferRefundSimulateCmd() *cobra.Command {
	var refundID, eventType, testClockID, webhook string
	var failureCode, achReturnCode, failureDescription string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "refund-simulate",
		Short: "Call /sandbox/transfer/refund/simulate",
		Long:  "Capability: sandbox-write. Simulates an asynchronous refund event for a Transfer refund.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"refund_id":  "<refund-id>",
				"event_type": "refund.posted",
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "refund-id", refundID, "refund_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "event-type", eventType, "event_type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "test-clock-id", testClockID, "test_clock_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := applyFailureReasonFlags(cmd, body, failureCode, achReturnCode, failureDescription); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--refund-id":  {"refund_id"},
				"--event-type": {"event_type"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/refund/simulate", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindFailureReasonFlags(cmd, &failureCode, &achReturnCode, &failureDescription)
	cmd.Flags().StringVar(&refundID, "refund-id", "", "Refund ID to simulate an event for")
	cmd.Flags().StringVar(&eventType, "event-type", "", "Refund event type to simulate")
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Sandbox Transfer test clock ID")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive a sandbox Transfer webhook")
	return cmd
}

func newSandboxTransferSweepSimulateCmd() *cobra.Command {
	var testClockID, webhook string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "sweep-simulate",
		Short: "Call /sandbox/transfer/sweep/simulate",
		Long:  "Capability: sandbox-write. Simulates a Transfer sweep and the associated sweep events.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "test-clock-id", testClockID, "test_clock_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/sweep/simulate", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Sandbox Transfer test clock ID")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive a sandbox Transfer webhook")
	return cmd
}
