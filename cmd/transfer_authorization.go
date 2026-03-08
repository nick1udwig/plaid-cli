package cmd

import (
	"github.com/spf13/cobra"
)

func newTransferAuthorizationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authorization",
		Short: "Transfer authorization commands",
		Long:  "Transfer authorization commands. `create` is a write operation; `cancel` revokes an unused authorization.",
	}

	cmd.AddCommand(newTransferAuthorizationCreateCmd())
	cmd.AddCommand(newTransferAuthorizationCancelCmd())

	return cmd
}

func newTransferAuthorizationCreateCmd() *cobra.Command {
	var itemID, accessToken, accountID string
	var transferType, network, amount, achClass string
	var legalName, phoneNumber, emailAddress string
	var ipAddress, userAgent string
	var ledgerID, isoCurrencyCode, idempotencyKey string
	var originatorClientID, testClockID, rulesetKey string
	var userPresent bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /transfer/authorization/create",
		Long:  "Capabilities: write, move-money. Creates a Transfer authorization, which is required before creating a transfer.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"account_id":   "<account-id>",
				"type":         "debit",
				"network":      "ach",
				"amount":       "12.34",
				"user": map[string]any{
					"legal_name": "Jane Doe",
				},
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
			if err := applyStringFlag(cmd, body, "type", transferType, "type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "network", network, "network"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ach-class", achClass, "ach_class"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "legal-name", legalName, "user", "legal_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "phone-number", phoneNumber, "user", "phone_number"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "email-address", emailAddress, "user", "email_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ip-address", ipAddress, "device", "ip_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-agent", userAgent, "device", "user_agent"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ledger-id", ledgerID, "ledger_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "iso-currency-code", isoCurrencyCode, "iso_currency_code"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "test-clock-id", testClockID, "test_clock_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ruleset-key", rulesetKey, "ruleset_key"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "user-present", userPresent, "user_present"); err != nil {
				return err
			}

			if err := requireBodyFields(body, map[string][]string{
				"--type":         {"type"},
				"--network":      {"network"},
				"--amount":       {"amount"},
				"--legal-name":   {"user", "legal_name"},
				"--account-id":   {"account_id"},
				"--access-token": {"access_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/authorization/create", body)
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
	cmd.Flags().StringVar(&transferType, "type", "", "Transfer type: debit or credit")
	cmd.Flags().StringVar(&network, "network", "", "Transfer network: ach, same-day-ach, rtp, or wire")
	cmd.Flags().StringVar(&amount, "amount", "", "Transfer amount as a decimal string, e.g. 12.34")
	cmd.Flags().StringVar(&achClass, "ach-class", "", "ACH SEC code for ACH transfers")
	cmd.Flags().StringVar(&legalName, "legal-name", "", "Account holder legal name")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Account holder phone number")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "Account holder email address")
	cmd.Flags().StringVar(&ipAddress, "ip-address", "", "Device IP address for the authorization request")
	cmd.Flags().StringVar(&userAgent, "user-agent", "", "Device user agent for the authorization request")
	cmd.Flags().StringVar(&ledgerID, "ledger-id", "", "Plaid ledger_id to fund the transfer from")
	cmd.Flags().StringVar(&isoCurrencyCode, "iso-currency-code", "USD", "Transfer currency code")
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Optional idempotency key for safely retrying the authorization")
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID for platform use cases")
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Sandbox test clock ID")
	cmd.Flags().StringVar(&rulesetKey, "ruleset-key", "", "Optional Transfer ruleset key")
	cmd.Flags().BoolVar(&userPresent, "user-present", false, "Whether the end user is actively initiating the transfer")

	return cmd
}

func newTransferAuthorizationCancelCmd() *cobra.Command {
	var authorizationID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "Call /transfer/authorization/cancel",
		Long:  "Capabilities: write, move-money. Cancels an unused transfer authorization.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"authorization_id": "<authorization-id>",
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
			if err := applyStringFlag(cmd, body, "authorization-id", authorizationID, "authorization_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--authorization-id": {"authorization_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/authorization/cancel", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&authorizationID, "authorization-id", "", "Plaid authorization_id to cancel")
	return cmd
}
