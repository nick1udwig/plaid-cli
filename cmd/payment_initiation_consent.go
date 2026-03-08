package cmd

import (
	"github.com/spf13/cobra"
)

func newPaymentInitiationConsentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "consent",
		Short: "Payment consent commands",
		Long:  "Payment Initiation consent commands for creating, reading, revoking, and executing consent-backed payments.",
	}

	cmd.AddCommand(newPaymentInitiationConsentCreateCmd())
	cmd.AddCommand(newPaymentInitiationConsentGetCmd())
	cmd.AddCommand(newPaymentInitiationConsentRevokeCmd())
	cmd.AddCommand(newPaymentInitiationConsentPaymentExecuteCmd())

	return cmd
}

func newPaymentInitiationConsentCreateCmd() *cobra.Command {
	var recipientID, reference, consentType string
	var validFrom, validTo string
	var maxAmountCurrency, maxAmountValue string
	var periodicAmountCurrency, periodicAmountValue, periodicInterval, periodicAlignment string
	var payerName, payerIBAN, payerBacsAccount, payerBacsSortCode, payerDateOfBirth string
	var payerPhones, payerEmails []string
	var payerAddressStreet []string
	var payerAddressCity, payerAddressPostalCode, payerAddressCountry string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /payment_initiation/consent/create",
		Long:  "Capabilities: write, move-money. Creates a payment consent that can later be used to initiate payments on behalf of the user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"recipient_id": "<recipient-id>",
				"reference":    "TestConsent",
				"constraints": map[string]any{
					"max_payment_amount": map[string]any{
						"currency": "GBP",
						"value":    15.0,
					},
					"periodic_amounts": []any{
						map[string]any{
							"amount": map[string]any{
								"currency": "GBP",
								"value":    40.0,
							},
							"interval":  "MONTH",
							"alignment": "CALENDAR",
						},
					},
				},
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
			if err := applyStringFlag(cmd, body, "recipient-id", recipientID, "recipient_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "reference", reference, "reference"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "type", consentType, "type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "valid-from", validFrom, "constraints", "valid_date_time", "from"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "valid-to", validTo, "constraints", "valid_date_time", "to"); err != nil {
				return err
			}
			if err := applyPIAmountFlags(cmd, body, maxAmountCurrency, maxAmountValue, "max-amount-currency", "max-amount-value", "constraints", "max_payment_amount"); err != nil {
				return err
			}
			if err := applyPIConsentPeriodicFlags(cmd, body, periodicAmountCurrency, periodicAmountValue, periodicInterval, periodicAlignment); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "payer-name", payerName, "payer_details", "name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "payer-iban", payerIBAN, "payer_details", "numbers", "iban"); err != nil {
				return err
			}
			if err := applyPIBACSFlags(cmd, body, payerBacsAccount, payerBacsSortCode, "payer-bacs-account", "payer-bacs-sort-code", "payer_details", "numbers", "bacs"); err != nil {
				return err
			}
			if err := applyPIAddressFlags(cmd, body, payerAddressStreet, payerAddressCity, payerAddressPostalCode, payerAddressCountry, "payer-address-street", "payer-address-city", "payer-address-postal-code", "payer-address-country", "payer_details", "address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "payer-date-of-birth", payerDateOfBirth, "payer_details", "date_of_birth"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "payer-phone", payerPhones, "payer_details", "phone_numbers"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "payer-email", payerEmails, "payer_details", "emails"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--recipient-id": {"recipient_id"},
				"--reference":    {"reference"},
			}); err != nil {
				return err
			}
			if err := validatePIAmount(body, "constraints.max_payment_amount", "constraints", "max_payment_amount"); err != nil {
				return err
			}
			if err := validatePIConsentPeriodicFlags(body); err != nil {
				return err
			}
			if bodyHasValue(body, "payer_details") {
				if err := requireBodyFields(body, map[string][]string{
					"payer_details.name": {"payer_details", "name"},
				}); err != nil {
					return err
				}
				if err := requireAtLeastOneBodyField(body, map[string][]string{
					"payer_details.numbers.iban":         {"payer_details", "numbers", "iban"},
					"payer_details.numbers.bacs.account": {"payer_details", "numbers", "bacs", "account"},
				}); err != nil {
					return err
				}
				if err := validatePIBACS(body, "payer_details.numbers.bacs", "payer_details", "numbers", "bacs"); err != nil {
					return err
				}
				if err := validatePIAddress(body, "payer_details.address", "payer_details", "address"); err != nil {
					return err
				}
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/consent/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&recipientID, "recipient-id", "", "Plaid payment recipient_id")
	cmd.Flags().StringVar(&reference, "reference", "", "Payment consent reference")
	cmd.Flags().StringVar(&consentType, "type", "", "Consent type: SWEEPING or COMMERCIAL")
	cmd.Flags().StringVar(&validFrom, "valid-from", "", "Optional consent activation timestamp in RFC3339 format")
	cmd.Flags().StringVar(&validTo, "valid-to", "", "Optional consent expiration timestamp in RFC3339 format")
	cmd.Flags().StringVar(&maxAmountCurrency, "max-amount-currency", "", "Maximum single-payment currency")
	cmd.Flags().StringVar(&maxAmountValue, "max-amount-value", "", "Maximum single-payment amount as a decimal number")
	cmd.Flags().StringVar(&periodicAmountCurrency, "periodic-amount-currency", "", "Periodic limit currency")
	cmd.Flags().StringVar(&periodicAmountValue, "periodic-amount-value", "", "Periodic limit amount as a decimal number")
	cmd.Flags().StringVar(&periodicInterval, "periodic-interval", "", "Periodic interval: DAY, WEEK, MONTH, or YEAR")
	cmd.Flags().StringVar(&periodicAlignment, "periodic-alignment", "", "Periodic alignment: CALENDAR or CONSENT")
	cmd.Flags().StringVar(&payerName, "payer-name", "", "Optional payer name for locking the consent to a specific account")
	cmd.Flags().StringVar(&payerIBAN, "payer-iban", "", "Optional payer IBAN for locking the consent to a specific account")
	cmd.Flags().StringVar(&payerBacsAccount, "payer-bacs-account", "", "Optional payer BACS account number")
	cmd.Flags().StringVar(&payerBacsSortCode, "payer-bacs-sort-code", "", "Optional payer BACS sort code")
	cmd.Flags().StringSliceVar(&payerAddressStreet, "payer-address-street", nil, "Payer address street line (repeatable, max 2)")
	cmd.Flags().StringVar(&payerAddressCity, "payer-address-city", "", "Payer address city")
	cmd.Flags().StringVar(&payerAddressPostalCode, "payer-address-postal-code", "", "Payer address postal code")
	cmd.Flags().StringVar(&payerAddressCountry, "payer-address-country", "", "Payer address country code")
	cmd.Flags().StringVar(&payerDateOfBirth, "payer-date-of-birth", "", "Optional payer date of birth in YYYY-MM-DD format")
	cmd.Flags().StringSliceVar(&payerPhones, "payer-phone", nil, "Optional payer phone number in E.164 format (repeatable)")
	cmd.Flags().StringSliceVar(&payerEmails, "payer-email", nil, "Optional payer email address (repeatable)")
	return cmd
}

func newPaymentInitiationConsentGetCmd() *cobra.Command {
	var consentID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /payment_initiation/consent/get",
		Long:  "Capability: read. Retrieves a Payment Initiation consent by consent_id.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"consent_id": "<consent-id>"}
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
			if err := applyStringFlag(cmd, body, "consent-id", consentID, "consent_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--consent-id": {"consent_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/consent/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&consentID, "consent-id", "", "Plaid consent_id")
	return cmd
}

func newPaymentInitiationConsentRevokeCmd() *cobra.Command {
	var consentID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "revoke",
		Short: "Call /payment_initiation/consent/revoke",
		Long:  "Capabilities: write, move-money. Revokes a Payment Initiation consent so it can no longer be used.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"consent_id": "<consent-id>"}
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
			if err := applyStringFlag(cmd, body, "consent-id", consentID, "consent_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--consent-id": {"consent_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/consent/revoke", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&consentID, "consent-id", "", "Plaid consent_id")
	return cmd
}

func newPaymentInitiationConsentPaymentExecuteCmd() *cobra.Command {
	var consentID, amountCurrency, amountValue, idempotencyKey, reference, processingMode string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "payment-execute",
		Short: "Call /payment_initiation/consent/payment/execute",
		Long:  "Capabilities: write, move-money. Executes a payment using a previously authorised payment consent.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"consent_id":      "<consent-id>",
				"idempotency_key": "<idempotency-key>",
				"amount": map[string]any{
					"currency": "GBP",
					"value":    7.99,
				},
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
			if err := applyStringFlag(cmd, body, "consent-id", consentID, "consent_id"); err != nil {
				return err
			}
			if err := applyPIAmountFlags(cmd, body, amountCurrency, amountValue, "amount-currency", "amount-value", "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "reference", reference, "reference"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "processing-mode", processingMode, "processing_mode"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--consent-id":      {"consent_id"},
				"--idempotency-key": {"idempotency_key"},
			}); err != nil {
				return err
			}
			if err := validatePIAmount(body, "amount", "amount"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/consent/payment/execute", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&consentID, "consent-id", "", "Plaid consent_id")
	cmd.Flags().StringVar(&amountCurrency, "amount-currency", "", "Payment currency, e.g. GBP or EUR")
	cmd.Flags().StringVar(&amountValue, "amount-value", "", "Payment amount as a decimal number")
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency key for the payment execution request")
	cmd.Flags().StringVar(&reference, "reference", "", "Optional payment reference; falls back to the consent reference when omitted")
	cmd.Flags().StringVar(&processingMode, "processing-mode", "", "Processing mode: IMMEDIATE or ASYNC")
	return cmd
}
