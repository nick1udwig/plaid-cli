package cmd

import "github.com/spf13/cobra"

func newSandboxTransferLedgerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ledger",
		Short: "Sandbox Transfer ledger helpers",
		Long:  "Sandbox-only helpers for simulating Plaid Ledger sweep and balance events used by Transfer.",
	}

	cmd.AddCommand(newSandboxTransferLedgerDepositSimulateCmd())
	cmd.AddCommand(newSandboxTransferLedgerSimulateAvailableCmd())
	cmd.AddCommand(newSandboxTransferLedgerWithdrawSimulateCmd())

	return cmd
}

func newSandboxTransferLedgerDepositSimulateCmd() *cobra.Command {
	var sweepID, eventType string
	var failureCode, achReturnCode, failureDescription string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "deposit-simulate",
		Short: "Call /sandbox/transfer/ledger/deposit/simulate",
		Long:  "Capability: sandbox-write. Simulates a ledger deposit sweep event.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"sweep_id":   "<sweep-id>",
				"event_type": "sweep.posted",
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
			if err := applyStringFlag(cmd, body, "sweep-id", sweepID, "sweep_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "event-type", eventType, "event_type"); err != nil {
				return err
			}
			if err := applyFailureReasonFlags(cmd, body, failureCode, achReturnCode, failureDescription); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--sweep-id":   {"sweep_id"},
				"--event-type": {"event_type"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/ledger/deposit/simulate", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindFailureReasonFlags(cmd, &failureCode, &achReturnCode, &failureDescription)
	cmd.Flags().StringVar(&sweepID, "sweep-id", "", "Sweep ID to simulate a ledger deposit event for")
	cmd.Flags().StringVar(&eventType, "event-type", "", "Ledger deposit event type to simulate")
	return cmd
}

func newSandboxTransferLedgerSimulateAvailableCmd() *cobra.Command {
	var ledgerID, originatorClientID, testClockID, webhook string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "simulate-available",
		Short: "Call /sandbox/transfer/ledger/simulate_available",
		Long:  "Capability: sandbox-write. Simulates converting pending Transfer ledger balance into available balance.",
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
			if err := applyStringFlag(cmd, body, "ledger-id", ledgerID, "ledger_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
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
			resp, err := client.Call(ctx, "/sandbox/transfer/ledger/simulate_available", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&ledgerID, "ledger-id", "", "Ledger ID to simulate available balance for")
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID for platform use cases")
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Sandbox Transfer test clock ID")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive a sandbox Transfer webhook")
	return cmd
}

func newSandboxTransferLedgerWithdrawSimulateCmd() *cobra.Command {
	var sweepID, eventType string
	var failureCode, achReturnCode, failureDescription string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "withdraw-simulate",
		Short: "Call /sandbox/transfer/ledger/withdraw/simulate",
		Long:  "Capability: sandbox-write. Simulates a ledger withdrawal sweep event.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"sweep_id":   "<sweep-id>",
				"event_type": "sweep.posted",
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
			if err := applyStringFlag(cmd, body, "sweep-id", sweepID, "sweep_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "event-type", eventType, "event_type"); err != nil {
				return err
			}
			if err := applyFailureReasonFlags(cmd, body, failureCode, achReturnCode, failureDescription); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--sweep-id":   {"sweep_id"},
				"--event-type": {"event_type"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/ledger/withdraw/simulate", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindFailureReasonFlags(cmd, &failureCode, &achReturnCode, &failureDescription)
	cmd.Flags().StringVar(&sweepID, "sweep-id", "", "Sweep ID to simulate a ledger withdrawal event for")
	cmd.Flags().StringVar(&eventType, "event-type", "", "Ledger withdrawal event type to simulate")
	return cmd
}
