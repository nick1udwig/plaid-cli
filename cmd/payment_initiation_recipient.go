package cmd

import (
	"github.com/spf13/cobra"
)

func newPaymentInitiationRecipientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recipient",
		Short: "Payment recipient commands",
		Long:  "Payment Initiation recipient commands for creating and listing payout destinations.",
	}

	cmd.AddCommand(newPaymentInitiationRecipientCreateCmd())
	cmd.AddCommand(newPaymentInitiationRecipientGetCmd())
	cmd.AddCommand(newPaymentInitiationRecipientListCmd())

	return cmd
}

func newPaymentInitiationRecipientCreateCmd() *cobra.Command {
	var name, iban, bacsAccount, bacsSortCode string
	var addressStreet []string
	var addressCity, addressPostalCode, addressCountry string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /payment_initiation/recipient/create",
		Long:  "Capability: write. Creates a payment recipient for UK or European Payment Initiation flows.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"name": "Wonder Wallet",
				"iban": "GB29NWBK60161331926819",
			}
			if handled, err := maybeWriteInfo(cmd, info, paymentInitiationDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "name", name, "name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "iban", iban, "iban"); err != nil {
				return err
			}
			if err := applyPIBACSFlags(cmd, body, bacsAccount, bacsSortCode, "bacs-account", "bacs-sort-code", "bacs"); err != nil {
				return err
			}
			if err := applyPIAddressFlags(cmd, body, addressStreet, addressCity, addressPostalCode, addressCountry, "address-street", "address-city", "address-postal-code", "address-country", "address"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--name": {"name"},
			}); err != nil {
				return err
			}
			if err := requireAtLeastOneBodyField(body, map[string][]string{
				"--iban":         {"iban"},
				"--bacs-account": {"bacs", "account"},
			}); err != nil {
				return err
			}
			if err := validatePIBACS(body, "bacs", "bacs"); err != nil {
				return err
			}
			if err := validatePIAddress(body, "address", "address"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/recipient/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&name, "name", "", "Recipient name")
	cmd.Flags().StringVar(&iban, "iban", "", "Recipient IBAN")
	cmd.Flags().StringVar(&bacsAccount, "bacs-account", "", "Recipient BACS account number")
	cmd.Flags().StringVar(&bacsSortCode, "bacs-sort-code", "", "Recipient BACS sort code")
	cmd.Flags().StringSliceVar(&addressStreet, "address-street", nil, "Recipient address street line (repeatable, max 2)")
	cmd.Flags().StringVar(&addressCity, "address-city", "", "Recipient address city")
	cmd.Flags().StringVar(&addressPostalCode, "address-postal-code", "", "Recipient address postal code")
	cmd.Flags().StringVar(&addressCountry, "address-country", "", "Recipient address country code")
	return cmd
}

func newPaymentInitiationRecipientGetCmd() *cobra.Command {
	var recipientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /payment_initiation/recipient/get",
		Long:  "Capability: read. Retrieves a previously created Payment Initiation recipient.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"recipient_id": "<recipient-id>"}
			if handled, err := maybeWriteInfo(cmd, info, paymentInitiationDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "recipient-id", recipientID, "recipient_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--recipient-id": {"recipient_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/recipient/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&recipientID, "recipient-id", "", "Plaid payment recipient_id")
	return cmd
}

func newPaymentInitiationRecipientListCmd() *cobra.Command {
	var count int
	var cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /payment_initiation/recipient/list",
		Long:  "Capability: read. Lists Payment Initiation recipients with optional cursor pagination.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"count": count}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, paymentInitiationDocPath, template); handled || err != nil {
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
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/recipient/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().IntVar(&count, "count", 100, "Maximum number of recipients to return")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor from the previous recipient list response")
	return cmd
}
