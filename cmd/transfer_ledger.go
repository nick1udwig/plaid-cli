package cmd

import "github.com/spf13/cobra"

func newTransferLedgerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ledger",
		Short: "Plaid Ledger commands",
		Long:  "Commands for depositing, withdrawing, moving, and inspecting Plaid Ledger balances.",
	}

	cmd.AddCommand(newTransferLedgerDepositCmd())
	cmd.AddCommand(newTransferLedgerDistributeCmd())
	cmd.AddCommand(newTransferLedgerGetCmd())
	cmd.AddCommand(newTransferLedgerWithdrawCmd())
	cmd.AddCommand(newTransferLedgerEventListCmd())

	return cmd
}

func newTransferLedgerDepositCmd() *cobra.Command {
	var originatorClientID, fundingAccountID, ledgerID, amount, description, idempotencyKey, network string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "deposit",
		Short: "Call /transfer/ledger/deposit",
		Long:  "Capabilities: write, move-money. Deposits funds into a Plaid Ledger balance.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"amount":          "12.34",
				"idempotency_key": "<idempotency-key>",
				"network":         "ach",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferLedgerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "funding-account-id", fundingAccountID, "funding_account_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ledger-id", ledgerID, "ledger_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "description", description, "description"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "network", network, "network"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--amount":          {"amount"},
				"--idempotency-key": {"idempotency_key"},
				"--network":         {"network"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/ledger/deposit", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID for platform ledger flows")
	cmd.Flags().StringVar(&fundingAccountID, "funding-account-id", "", "Funding account ID to use for the ledger deposit")
	cmd.Flags().StringVar(&ledgerID, "ledger-id", "", "Ledger ID to deposit into")
	cmd.Flags().StringVar(&amount, "amount", "", "Ledger deposit amount as a decimal string")
	cmd.Flags().StringVar(&description, "description", "", "Bank statement description for the deposit")
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency key for safely retrying the deposit")
	cmd.Flags().StringVar(&network, "network", "", "Transfer network to use, e.g. ach or same-day-ach")
	return cmd
}

func newTransferLedgerDistributeCmd() *cobra.Command {
	var fromLedgerID, toLedgerID, amount, idempotencyKey, description string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "distribute",
		Short: "Call /transfer/ledger/distribute",
		Long:  "Capabilities: write, move-money. Moves available balance between Plaid Ledger balances.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"from_ledger_id":  "<from-ledger-id>",
				"to_ledger_id":    "<to-ledger-id>",
				"amount":          "12.34",
				"idempotency_key": "<idempotency-key>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferLedgerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "from-ledger-id", fromLedgerID, "from_ledger_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "to-ledger-id", toLedgerID, "to_ledger_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "description", description, "description"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--from-ledger-id":  {"from_ledger_id"},
				"--to-ledger-id":    {"to_ledger_id"},
				"--amount":          {"amount"},
				"--idempotency-key": {"idempotency_key"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/ledger/distribute", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&fromLedgerID, "from-ledger-id", "", "Ledger ID to debit")
	cmd.Flags().StringVar(&toLedgerID, "to-ledger-id", "", "Ledger ID to credit")
	cmd.Flags().StringVar(&amount, "amount", "", "Amount to move as a decimal string")
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency key for safely retrying the distribution")
	cmd.Flags().StringVar(&description, "description", "", "Optional distribution description")
	return cmd
}

func newTransferLedgerGetCmd() *cobra.Command {
	var ledgerID, originatorClientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/ledger/get",
		Long:  "Capability: read. Retrieves a Plaid Ledger balance.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{}
			if handled, err := maybeWriteInfo(cmd, info, transferLedgerDocPath, template); handled || err != nil {
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

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/ledger/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&ledgerID, "ledger-id", "", "Ledger ID to retrieve")
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID for platform ledger flows")
	return cmd
}

func newTransferLedgerWithdrawCmd() *cobra.Command {
	var originatorClientID, fundingAccountID, ledgerID, amount, description, idempotencyKey, network string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "withdraw",
		Short: "Call /transfer/ledger/withdraw",
		Long:  "Capabilities: write, move-money. Withdraws funds from a Plaid Ledger balance.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"amount":          "12.34",
				"idempotency_key": "<idempotency-key>",
				"network":         "ach",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferLedgerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "funding-account-id", fundingAccountID, "funding_account_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ledger-id", ledgerID, "ledger_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "description", description, "description"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "network", network, "network"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--amount":          {"amount"},
				"--idempotency-key": {"idempotency_key"},
				"--network":         {"network"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/ledger/withdraw", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID for platform ledger flows")
	cmd.Flags().StringVar(&fundingAccountID, "funding-account-id", "", "Funding account ID to receive the withdrawal")
	cmd.Flags().StringVar(&ledgerID, "ledger-id", "", "Ledger ID to withdraw from")
	cmd.Flags().StringVar(&amount, "amount", "", "Ledger withdrawal amount as a decimal string")
	cmd.Flags().StringVar(&description, "description", "", "Bank statement description for the withdrawal")
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency key for safely retrying the withdrawal")
	cmd.Flags().StringVar(&network, "network", "", "Transfer network to use, e.g. ach, same-day-ach, rtp, or wire")
	return cmd
}

func newTransferLedgerEventListCmd() *cobra.Command {
	var originatorClientID, startDate, endDate, ledgerID, ledgerEventID, sourceType, sourceID string
	var count, offset int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "event-list",
		Short: "Call /transfer/ledger/event/list",
		Long:  "Capability: read. Lists Plaid Ledger events with optional filters.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"count":  count,
				"offset": offset,
			}
			if handled, err := maybeWriteInfo(cmd, info, transferLedgerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "start-date", startDate, "start_date"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-date", endDate, "end_date"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ledger-id", ledgerID, "ledger_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ledger-event-id", ledgerEventID, "ledger_event_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "source-type", sourceType, "source_type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "source-id", sourceID, "source_id"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "offset", offset, "offset"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/ledger/event/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID filter")
	cmd.Flags().StringVar(&startDate, "start-date", "", "RFC3339 start timestamp filter")
	cmd.Flags().StringVar(&endDate, "end-date", "", "RFC3339 end timestamp filter")
	cmd.Flags().StringVar(&ledgerID, "ledger-id", "", "Ledger ID filter")
	cmd.Flags().StringVar(&ledgerEventID, "ledger-event-id", "", "Ledger event ID filter")
	cmd.Flags().StringVar(&sourceType, "source-type", "", "Source type filter: TRANSFER, SWEEP, or REFUND")
	cmd.Flags().StringVar(&sourceID, "source-id", "", "Transfer, sweep, or refund ID filter")
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of ledger events to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Ledger event list offset")
	return cmd
}
