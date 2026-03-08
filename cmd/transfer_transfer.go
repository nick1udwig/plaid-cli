package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newTransferCreateCmd() *cobra.Command {
	var itemID, accessToken, accountID string
	var authorizationID, amount, description string
	var facilitatorFee, testClockID string
	var metadata map[string]string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /transfer/create",
		Long:  "Capabilities: write, move-money. Creates a transfer from an approved transfer authorization.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token":     "<access-token>",
				"account_id":       "<account-id>",
				"authorization_id": "<authorization-id>",
				"description":      "payment",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferInitiatingDocPath, template); handled || err != nil {
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

			if _, err := populateTransferAccess(cmd, store, body, itemID, accessToken, accountID); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "authorization-id", authorizationID, "authorization_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "description", description, "description"); err != nil {
				return err
			}
			if err := applyStringMapFlag(cmd, body, "metadata", metadata, "metadata"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "test-clock-id", testClockID, "test_clock_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "facilitator-fee", facilitatorFee, "facilitator_fee"); err != nil {
				return err
			}

			if err := requireBodyFields(body, map[string][]string{
				"--authorization-id": {"authorization_id"},
				"--description":      {"description"},
				"--account-id":       {"account_id"},
				"--access-token":     {"access_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/create", body)
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
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id to debit or credit")
	cmd.Flags().StringVar(&authorizationID, "authorization-id", "", "Plaid authorization_id returned by transfer authorization")
	cmd.Flags().StringVar(&amount, "amount", "", "Exact transfer amount as a decimal string; defaults to the authorized amount when omitted")
	cmd.Flags().StringVar(&description, "description", "", "Transfer description")
	cmd.Flags().StringToStringVar(&metadata, "metadata", nil, "Transfer metadata as key=value pairs")
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Sandbox test clock ID")
	cmd.Flags().StringVar(&facilitatorFee, "facilitator-fee", "", "Facilitator fee amount to deduct from the transfer")
	return cmd
}

func newTransferGetCmd() *cobra.Command {
	var transferID, authorizationID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/get",
		Long:  "Capability: read. Retrieves a transfer by transfer_id or authorization_id.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"transfer_id": "<transfer-id>",
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
			if err := applyStringFlag(cmd, body, "transfer-id", transferID, "transfer_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "authorization-id", authorizationID, "authorization_id"); err != nil {
				return err
			}

			hasTransferID := bodyHasValue(body, "transfer_id")
			hasAuthorizationID := bodyHasValue(body, "authorization_id")
			if hasTransferID == hasAuthorizationID {
				return errors.New("provide exactly one of --transfer-id or --authorization-id")
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&transferID, "transfer-id", "", "Plaid transfer_id")
	cmd.Flags().StringVar(&authorizationID, "authorization-id", "", "Plaid authorization_id")
	return cmd
}

func newTransferListCmd() *cobra.Command {
	var startDate, endDate, originatorClientID, fundingAccountID string
	var count, offset int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /transfer/list",
		Long:  "Capability: read. Lists transfers, optionally filtered by date range and originator or funding account.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"count":  count,
				"offset": offset,
			}
			if startDate != "" {
				template["start_date"] = startDate
			}
			if endDate != "" {
				template["end_date"] = endDate
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "funding-account-id", fundingAccountID, "funding_account_id"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/list", body)
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
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of transfers to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Transfer list offset")
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID filter")
	cmd.Flags().StringVar(&fundingAccountID, "funding-account-id", "", "Funding account ID filter")
	return cmd
}

func newTransferCancelCmd() *cobra.Command {
	var transferID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "Call /transfer/cancel",
		Long:  "Capabilities: write, move-money. Cancels a transfer that has not yet been submitted to the payment network.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"transfer_id": "<transfer-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferInitiatingDocPath, template); handled || err != nil {
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
			if err := requireBodyFields(body, map[string][]string{
				"--transfer-id": {"transfer_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/cancel", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&transferID, "transfer-id", "", "Plaid transfer_id to cancel")
	return cmd
}
