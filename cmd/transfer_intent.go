package cmd

import "github.com/spf13/cobra"

func newTransferIntentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "intent",
		Short: "Transfer intent commands",
		Long:  "Transfer intent commands for Transfer UI flows.",
	}

	cmd.AddCommand(newTransferIntentCreateCmd())
	cmd.AddCommand(newTransferIntentGetCmd())

	return cmd
}

func newTransferIntentCreateCmd() *cobra.Command {
	var accountID, mode, network, amount, description, achClass string
	var legalName, phoneNumber, emailAddress string
	var isoCurrencyCode, originationAccountID string
	var metadata map[string]string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /transfer/intent/create",
		Long:  "Capabilities: write, move-money. Creates a Transfer intent for Transfer UI flows.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"mode":        "PAYMENT",
				"amount":      "12.34",
				"description": "Desc",
				"user": map[string]any{
					"legal_name": "Jane Doe",
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, transferAccountLinkingDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "account-id", accountID, "account_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "mode", mode, "mode"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "network", network, "network"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "description", description, "description"); err != nil {
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
			if err := applyStringMapFlag(cmd, body, "metadata", metadata, "metadata"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "iso-currency-code", isoCurrencyCode, "iso_currency_code"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "origination-account-id", originationAccountID, "origination_account_id"); err != nil {
				return err
			}

			if err := requireBodyFields(body, map[string][]string{
				"--mode":        {"mode"},
				"--amount":      {"amount"},
				"--description": {"description"},
				"--legal-name":  {"user", "legal_name"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/intent/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id to preselect for Transfer UI")
	cmd.Flags().StringVar(&mode, "mode", "", "Transfer intent mode: PAYMENT or DISBURSEMENT")
	cmd.Flags().StringVar(&network, "network", "same-day-ach", "Transfer network")
	cmd.Flags().StringVar(&amount, "amount", "", "Transfer amount as a decimal string")
	cmd.Flags().StringVar(&description, "description", "", "Transfer description")
	cmd.Flags().StringVar(&achClass, "ach-class", "", "ACH SEC code for ACH transfers")
	cmd.Flags().StringVar(&legalName, "legal-name", "", "Account holder legal name")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Account holder phone number")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "Account holder email address")
	cmd.Flags().StringToStringVar(&metadata, "metadata", nil, "Transfer intent metadata as key=value pairs")
	cmd.Flags().StringVar(&isoCurrencyCode, "iso-currency-code", "USD", "Transfer currency code")
	cmd.Flags().StringVar(&originationAccountID, "origination-account-id", "", "Origination account ID for platform use cases")
	return cmd
}

func newTransferIntentGetCmd() *cobra.Command {
	var intentID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/intent/get",
		Long:  "Capability: read. Retrieves a Transfer intent by intent_id.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"intent_id": "<intent-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferAccountLinkingDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "intent-id", intentID, "intent_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--intent-id": {"intent_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/intent/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&intentID, "intent-id", "", "Transfer intent ID to retrieve")
	return cmd
}
