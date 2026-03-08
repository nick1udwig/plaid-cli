package cmd

import "github.com/spf13/cobra"

func newProcessorSignalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signal",
		Short: "Processor Signal commands",
		Long:  "Signal evaluation and reporting commands for processor tokens.",
	}

	cmd.AddCommand(newProcessorSignalEvaluateCmd())
	cmd.AddCommand(newProcessorSignalDecisionReportCmd())
	cmd.AddCommand(newProcessorSignalReturnReportCmd())
	cmd.AddCommand(newProcessorTokenOnlyCmd(
		"prepare",
		"Call /processor/signal/prepare",
		"Capability: write. Opts a processor token into Signal data collection.",
		"/processor/signal/prepare",
	))

	return cmd
}

func newProcessorSignalEvaluateCmd() *cobra.Command {
	var processorToken, clientTransactionID, amount string
	var clientUserID, defaultPaymentMethod, rulesetKey string
	var prefix, givenName, middleName, familyName, suffix string
	var phoneNumber, emailAddress string
	var city, region, street, postalCode, country string
	var ipAddress, userAgent string
	var isRecurring, userPresent bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "evaluate",
		Short: "Call /processor/signal/evaluate",
		Long:  "Capability: read. Evaluates a proposed ACH transaction for a processor token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token":       "<processor-token>",
				"client_transaction_id": "<client-transaction-id>",
				"amount":                123.45,
			}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "client-transaction-id", clientTransactionID, "client_transaction_id"); err != nil {
				return err
			}
			if err := applyDecimalStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "user-present", userPresent, "user_present"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "client-user-id", clientUserID, "client_user_id"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "is-recurring", isRecurring, "is_recurring"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "default-payment-method", defaultPaymentMethod, "default_payment_method"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ruleset-key", rulesetKey, "ruleset_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "name-prefix", prefix, "user", "name", "prefix"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "given-name", givenName, "user", "name", "given_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "middle-name", middleName, "user", "name", "middle_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "family-name", familyName, "user", "name", "family_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "name-suffix", suffix, "user", "name", "suffix"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "phone-number", phoneNumber, "user", "phone_number"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "email-address", emailAddress, "user", "email_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "city", city, "user", "address", "city"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "region", region, "user", "address", "region"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "street", street, "user", "address", "street"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "postal-code", postalCode, "user", "address", "postal_code"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "country", country, "user", "address", "country"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ip-address", ipAddress, "device", "ip_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-agent", userAgent, "device", "user_agent"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token":       {"processor_token"},
				"--client-transaction-id": {"client_transaction_id"},
				"--amount":                {"amount"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/signal/evaluate", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringVar(&clientTransactionID, "client-transaction-id", "", "Unique client-side ID for this ACH transaction attempt")
	cmd.Flags().StringVar(&amount, "amount", "", "Transaction amount as a decimal number")
	cmd.Flags().BoolVar(&userPresent, "user-present", false, "Whether the end user is present while initiating the ACH transfer")
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable end-user ID from your system")
	cmd.Flags().BoolVar(&isRecurring, "is-recurring", false, "Whether the ACH transaction is recurring")
	cmd.Flags().StringVar(&defaultPaymentMethod, "default-payment-method", "", "Default payment method: SAME_DAY_ACH, STANDARD_ACH, or MULTIPLE_PAYMENT_METHODS")
	cmd.Flags().StringVar(&rulesetKey, "ruleset-key", "", "Optional Signal ruleset key")
	cmd.Flags().StringVar(&prefix, "name-prefix", "", "User name prefix")
	cmd.Flags().StringVar(&givenName, "given-name", "", "User given name")
	cmd.Flags().StringVar(&middleName, "middle-name", "", "User middle name")
	cmd.Flags().StringVar(&familyName, "family-name", "", "User family name")
	cmd.Flags().StringVar(&suffix, "name-suffix", "", "User name suffix")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "User phone number")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "User email address")
	cmd.Flags().StringVar(&city, "city", "", "User address city")
	cmd.Flags().StringVar(&region, "region", "", "User address region")
	cmd.Flags().StringVar(&street, "street", "", "User address street")
	cmd.Flags().StringVar(&postalCode, "postal-code", "", "User address postal code")
	cmd.Flags().StringVar(&country, "country", "", "User address country")
	cmd.Flags().StringVar(&ipAddress, "ip-address", "", "End-user device IP address")
	cmd.Flags().StringVar(&userAgent, "user-agent", "", "End-user device user agent")
	return cmd
}

func newProcessorSignalDecisionReportCmd() *cobra.Command {
	var processorToken, clientTransactionID string
	var daysFundsOnHold int
	var decisionOutcome, paymentMethod, amountInstantlyAvailable string
	var initiated bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "decision-report",
		Short: "Call /processor/signal/decision/report",
		Long:  "Capability: write. Reports whether a processor-token ACH transaction was initiated.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token":       "<processor-token>",
				"client_transaction_id": "<client-transaction-id>",
				"initiated":             true,
			}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "client-transaction-id", clientTransactionID, "client_transaction_id"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "initiated", initiated, "initiated"); err != nil {
				return err
			}
			if cmd.Flags().Changed("days-funds-on-hold") {
				if err := setBodyValue(body, daysFundsOnHold, "days_funds_on_hold"); err != nil {
					return err
				}
			}
			if err := applyStringFlag(cmd, body, "decision-outcome", decisionOutcome, "decision_outcome"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "payment-method", paymentMethod, "payment_method"); err != nil {
				return err
			}
			if err := applyDecimalStringFlag(cmd, body, "amount-instantly-available", amountInstantlyAvailable, "amount_instantly_available"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token":       {"processor_token"},
				"--client-transaction-id": {"client_transaction_id"},
				"--initiated":             {"initiated"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/signal/decision/report", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringVar(&clientTransactionID, "client-transaction-id", "", "Client transaction ID previously sent to processor signal evaluate")
	cmd.Flags().BoolVar(&initiated, "initiated", false, "Whether the ACH transaction was initiated")
	cmd.Flags().IntVar(&daysFundsOnHold, "days-funds-on-hold", 0, "Actual number of days funds were held before becoming available")
	cmd.Flags().StringVar(&decisionOutcome, "decision-outcome", "", "Decision outcome")
	cmd.Flags().StringVar(&paymentMethod, "payment-method", "", "Payment method used")
	cmd.Flags().StringVar(&amountInstantlyAvailable, "amount-instantly-available", "", "Amount made instantly available as a decimal number")
	return cmd
}

func newProcessorSignalReturnReportCmd() *cobra.Command {
	var processorToken, clientTransactionID, returnCode, returnedAt string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "return-report",
		Short: "Call /processor/signal/return/report",
		Long:  "Capability: write. Reports an ACH return for a processor-token transaction previously sent to Signal.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token":       "<processor-token>",
				"client_transaction_id": "<client-transaction-id>",
				"return_code":           "R01",
			}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "client-transaction-id", clientTransactionID, "client_transaction_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "return-code", returnCode, "return_code"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "returned-at", returnedAt, "returned_at"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token":       {"processor_token"},
				"--client-transaction-id": {"client_transaction_id"},
				"--return-code":           {"return_code"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/signal/return/report", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringVar(&clientTransactionID, "client-transaction-id", "", "Client transaction ID previously sent to processor signal evaluate")
	cmd.Flags().StringVar(&returnCode, "return-code", "", "ACH return code, e.g. R01")
	cmd.Flags().StringVar(&returnedAt, "returned-at", "", "Optional ISO 8601 timestamp for when the return was received")
	return cmd
}
