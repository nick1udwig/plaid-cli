package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Auth product commands",
		Long:  "Auth commands. Most are read-only; `verify` performs a write-style verification request.",
	}
	cmd.AddCommand(newAuthGetCmd())
	cmd.AddCommand(newAuthVerifyCmd())
	cmd.AddCommand(newAuthBankTransferEventListCmd())
	cmd.AddCommand(newAuthBankTransferEventSyncCmd())
	return cmd
}

func newAuthGetCmd() *cobra.Command {
	var itemID, accessToken string
	var accountIDs []string
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /auth/get",
		Long:  "Capability: read. Retrieves account and routing details for eligible accounts on a linked Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"access_token": "<access-token>"}
			if len(accountIDs) > 0 {
				template["options"] = map[string]any{"account_ids": accountIDs}
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/auth/index.md", template); handled || err != nil {
				return err
			}

			store, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			token, _, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			body := map[string]any{"access_token": token}
			if len(accountIDs) > 0 {
				body["options"] = map[string]any{"account_ids": accountIDs}
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/auth/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringSliceVar(&accountIDs, "account-id", nil, "Account ID to filter by (repeatable)")
	return cmd
}

func newAuthVerifyCmd() *cobra.Command {
	var accountNumber, routingNumber, legalName string
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Call /auth/verify",
		Long:  "Capability: write. Verifies an account and routing number through Database Auth.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"numbers": map[string]any{
					"ach": map[string]any{
						"account": "<account-number>",
						"routing": "<routing-number>",
					},
				},
			}
			if legalName != "" {
				template["legal_name"] = legalName
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/auth/index.md", template); handled || err != nil {
				return err
			}
			if accountNumber == "" || routingNumber == "" {
				return errors.New("--account-number and --routing-number are required")
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body := map[string]any{
				"numbers": map[string]any{
					"ach": map[string]any{
						"account": accountNumber,
						"routing": routingNumber,
					},
				},
			}
			if legalName != "" {
				body["legal_name"] = legalName
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/auth/verify", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&accountNumber, "account-number", "", "ACH account number to verify")
	cmd.Flags().StringVar(&routingNumber, "routing-number", "", "ACH routing number to verify")
	cmd.Flags().StringVar(&legalName, "legal-name", "", "Optional account owner legal name")
	return cmd
}

func newAuthBankTransferEventListCmd() *cobra.Command {
	var startDate, endDate, bankTransferID, accountID, bankTransferType, originationAccountID, direction string
	var eventTypes []string
	var count, offset int
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "bank-transfer-event-list",
		Short: "Call /bank_transfer/event/list",
		Long:  "Capability: read. Lists historical bank transfer events.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			body := map[string]any{"count": count, "offset": offset}
			if startDate != "" {
				body["start_date"] = startDate
			}
			if endDate != "" {
				body["end_date"] = endDate
			}
			if bankTransferID != "" {
				body["bank_transfer_id"] = bankTransferID
			}
			if accountID != "" {
				body["account_id"] = accountID
			}
			if bankTransferType != "" {
				body["bank_transfer_type"] = bankTransferType
			}
			if len(eventTypes) > 0 {
				body["event_types"] = eventTypes
			}
			if originationAccountID != "" {
				body["origination_account_id"] = originationAccountID
			}
			if direction != "" {
				body["direction"] = direction
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/auth/index.md", body); handled || err != nil {
				return err
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/bank_transfer/event/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&startDate, "start-date", "", "RFC3339 start timestamp")
	cmd.Flags().StringVar(&endDate, "end-date", "", "RFC3339 end timestamp")
	cmd.Flags().StringVar(&bankTransferID, "bank-transfer-id", "", "Specific bank_transfer_id to filter by")
	cmd.Flags().StringVar(&accountID, "account-id", "", "Account ID to filter by")
	cmd.Flags().StringVar(&bankTransferType, "bank-transfer-type", "", "Transfer type filter: debit or credit")
	cmd.Flags().StringSliceVar(&eventTypes, "event-type", nil, "Event type filter (repeatable)")
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of events to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Event list offset")
	cmd.Flags().StringVar(&originationAccountID, "origination-account-id", "", "Origination account ID to filter by")
	cmd.Flags().StringVar(&direction, "direction", "", "Direction filter: inbound or outbound")
	return cmd
}

func newAuthBankTransferEventSyncCmd() *cobra.Command {
	var afterID, count int
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "bank-transfer-event-sync",
		Short: "Call /bank_transfer/event/sync",
		Long:  "Capability: read. Incrementally syncs bank transfer events.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			body := map[string]any{
				"after_id": afterID,
				"count":    count,
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/auth/index.md", body); handled || err != nil {
				return err
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/bank_transfer/event/sync", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().IntVar(&afterID, "after-id", 0, "Largest previously seen event_id, or 0 initially")
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of events to return")
	return cmd
}
