package cmd

import (
	"github.com/spf13/cobra"
)

func newPaymentInitiationPaymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payment",
		Short: "Payment commands",
		Long:  "Payment Initiation payment commands for creating, reading, listing, and reversing payments.",
	}

	cmd.AddCommand(newPaymentInitiationPaymentCreateCmd())
	cmd.AddCommand(newPaymentInitiationPaymentGetCmd())
	cmd.AddCommand(newPaymentInitiationPaymentListCmd())
	cmd.AddCommand(newPaymentInitiationPaymentReverseCmd())

	return cmd
}

func newPaymentInitiationPaymentCreateCmd() *cobra.Command {
	var recipientID, reference string
	var amountCurrency, amountValue string
	var scheduleInterval, scheduleStartDate, scheduleEndDate string
	var scheduleExecutionDay int
	var requestRefundDetails bool
	var payerIBAN, payerBacsAccount, payerBacsSortCode, scheme string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /payment_initiation/payment/create",
		Long:  "Capabilities: write, move-money. Creates a Payment Initiation payment or standing order.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"recipient_id": "<recipient-id>",
				"reference":    "TestPayment",
				"amount": map[string]any{
					"currency": "GBP",
					"value":    100.0,
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
			if err := applyPIAmountFlags(cmd, body, amountCurrency, amountValue, "amount-currency", "amount-value", "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "schedule-interval", scheduleInterval, "schedule", "interval"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "schedule-start-date", scheduleStartDate, "schedule", "start_date"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "schedule-end-date", scheduleEndDate, "schedule", "end_date"); err != nil {
				return err
			}
			if cmd.Flags().Changed("schedule-execution-day") {
				if err := setBodyValue(body, scheduleExecutionDay, "schedule", "interval_execution_day"); err != nil {
					return err
				}
			}
			if err := applyBoolFlag(cmd, body, "request-refund-details", requestRefundDetails, "options", "request_refund_details"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "payer-iban", payerIBAN, "options", "iban"); err != nil {
				return err
			}
			if err := applyPIBACSFlags(cmd, body, payerBacsAccount, payerBacsSortCode, "payer-bacs-account", "payer-bacs-sort-code", "options", "bacs"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "scheme", scheme, "options", "scheme"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--recipient-id": {"recipient_id"},
				"--reference":    {"reference"},
			}); err != nil {
				return err
			}
			if err := validatePIAmount(body, "amount", "amount"); err != nil {
				return err
			}
			if bodyHasValue(body, "schedule") {
				if err := requireBodyFields(body, map[string][]string{
					"schedule.interval":               {"schedule", "interval"},
					"schedule.interval_execution_day": {"schedule", "interval_execution_day"},
					"schedule.start_date":             {"schedule", "start_date"},
				}); err != nil {
					return err
				}
			}
			if err := validatePIBACS(body, "options.bacs", "options", "bacs"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/payment/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&recipientID, "recipient-id", "", "Plaid payment recipient_id")
	cmd.Flags().StringVar(&reference, "reference", "", "Payment reference")
	cmd.Flags().StringVar(&amountCurrency, "amount-currency", "", "Payment currency, e.g. GBP or EUR")
	cmd.Flags().StringVar(&amountValue, "amount-value", "", "Payment amount as a decimal number")
	cmd.Flags().StringVar(&scheduleInterval, "schedule-interval", "", "Standing order interval: WEEKLY or MONTHLY")
	cmd.Flags().IntVar(&scheduleExecutionDay, "schedule-execution-day", 0, "Standing order execution day")
	cmd.Flags().StringVar(&scheduleStartDate, "schedule-start-date", "", "Standing order start date in YYYY-MM-DD format")
	cmd.Flags().StringVar(&scheduleEndDate, "schedule-end-date", "", "Standing order end date in YYYY-MM-DD format")
	cmd.Flags().BoolVar(&requestRefundDetails, "request-refund-details", false, "Request refund details from the payee institution when supported")
	cmd.Flags().StringVar(&payerIBAN, "payer-iban", "", "Restrict payment initiation to a specific payer IBAN")
	cmd.Flags().StringVar(&payerBacsAccount, "payer-bacs-account", "", "Restrict payment initiation to a specific payer BACS account")
	cmd.Flags().StringVar(&payerBacsSortCode, "payer-bacs-sort-code", "", "Restrict payment initiation to a specific payer BACS sort code")
	cmd.Flags().StringVar(&scheme, "scheme", "", "Preferred payment scheme, e.g. LOCAL_DEFAULT or SEPA_CREDIT_TRANSFER")
	return cmd
}

func newPaymentInitiationPaymentGetCmd() *cobra.Command {
	var paymentID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /payment_initiation/payment/get",
		Long:  "Capability: read. Retrieves a Payment Initiation payment by payment_id.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"payment_id": "<payment-id>"}
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
			if err := applyStringFlag(cmd, body, "payment-id", paymentID, "payment_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--payment-id": {"payment_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/payment/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&paymentID, "payment-id", "", "Plaid payment_id")
	return cmd
}

func newPaymentInitiationPaymentListCmd() *cobra.Command {
	var count int
	var cursor, consentID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /payment_initiation/payment/list",
		Long:  "Capability: read. Lists Payment Initiation payments with optional cursor pagination.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"count": count}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if consentID != "" {
				template["consent_id"] = consentID
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
			if err := applyStringFlag(cmd, body, "consent-id", consentID, "consent_id"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/payment/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().IntVar(&count, "count", 10, "Maximum number of payments to return")
	cmd.Flags().StringVar(&cursor, "cursor", "", "RFC3339 cursor from the previous payment list response")
	cmd.Flags().StringVar(&consentID, "consent-id", "", "Limit results to payments executed from the given consent_id")
	return cmd
}

func newPaymentInitiationPaymentReverseCmd() *cobra.Command {
	var paymentID, idempotencyKey, reference string
	var amountCurrency, amountValue string
	var counterpartyDateOfBirth string
	var addressStreet []string
	var addressCity, addressPostalCode, addressCountry string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "reverse",
		Short: "Call /payment_initiation/payment/reverse",
		Long:  "Capabilities: write, move-money. Reverses a settled payment from a Plaid virtual account.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"payment_id":      "<payment-id>",
				"idempotency_key": "<idempotency-key>",
				"reference":       "Refund123",
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
			if err := applyStringFlag(cmd, body, "payment-id", paymentID, "payment_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "reference", reference, "reference"); err != nil {
				return err
			}
			if err := applyPIAmountFlags(cmd, body, amountCurrency, amountValue, "amount-currency", "amount-value", "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "counterparty-date-of-birth", counterpartyDateOfBirth, "counterparty_date_of_birth"); err != nil {
				return err
			}
			if err := applyPIAddressFlags(cmd, body, addressStreet, addressCity, addressPostalCode, addressCountry, "counterparty-address-street", "counterparty-address-city", "counterparty-address-postal-code", "counterparty-address-country", "counterparty_address"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--payment-id":      {"payment_id"},
				"--idempotency-key": {"idempotency_key"},
				"--reference":       {"reference"},
			}); err != nil {
				return err
			}
			if err := validatePIAmount(body, "amount", "amount"); err != nil {
				return err
			}
			if err := validatePIAddress(body, "counterparty_address", "counterparty_address"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/payment_initiation/payment/reverse", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&paymentID, "payment-id", "", "Plaid payment_id to reverse")
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency key for the reversal request")
	cmd.Flags().StringVar(&reference, "reference", "", "Refund reference")
	cmd.Flags().StringVar(&amountCurrency, "amount-currency", "", "Optional refund currency")
	cmd.Flags().StringVar(&amountValue, "amount-value", "", "Optional refund amount as a decimal number")
	cmd.Flags().StringVar(&counterpartyDateOfBirth, "counterparty-date-of-birth", "", "Optional counterparty date of birth in YYYY-MM-DD format")
	cmd.Flags().StringSliceVar(&addressStreet, "counterparty-address-street", nil, "Counterparty address street line (repeatable, max 2)")
	cmd.Flags().StringVar(&addressCity, "counterparty-address-city", "", "Counterparty address city")
	cmd.Flags().StringVar(&addressPostalCode, "counterparty-address-postal-code", "", "Counterparty address postal code")
	cmd.Flags().StringVar(&addressCountry, "counterparty-address-country", "", "Counterparty address country code")
	return cmd
}
