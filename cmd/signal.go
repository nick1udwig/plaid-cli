package cmd

import "github.com/spf13/cobra"

const signalDocPath = "docs/plaid/api/products/signal/index.md"

func newSignalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signal",
		Short: "Signal product commands",
		Long:  "Signal commands for ACH return-risk evaluation and outcome reporting.",
	}

	cmd.AddCommand(newSignalEvaluateCmd())
	cmd.AddCommand(newSignalDecisionReportCmd())
	cmd.AddCommand(newSignalReturnReportCmd())
	cmd.AddCommand(newSignalPrepareCmd())

	return cmd
}

func newSignalEvaluateCmd() *cobra.Command {
	var itemID, accessToken, accountID string
	var clientTransactionID, amount string
	var clientUserID, defaultPaymentMethod, rulesetKey string
	var prefix, givenName, middleName, familyName, suffix string
	var phoneNumber, emailAddress string
	var city, region, street, postalCode, country string
	var ipAddress, userAgent string
	var isRecurring bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "evaluate",
		Short: "Call /signal/evaluate",
		Long:  "Capability: read. Evaluates a proposed ACH transaction for return risk and related Signal insights.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token":          "<access-token>",
				"account_id":            "<account-id>",
				"client_transaction_id": "<client-transaction-id>",
				"amount":                123.45,
			}
			if handled, err := maybeWriteInfo(cmd, info, signalDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "client-transaction-id", clientTransactionID, "client_transaction_id"); err != nil {
				return err
			}
			if err := applyDecimalStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
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
				"--access-token":          {"access_token"},
				"--account-id":            {"account_id"},
				"--client-transaction-id": {"client_transaction_id"},
				"--amount":                {"amount"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/signal/evaluate", body)
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
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id to evaluate")
	cmd.Flags().StringVar(&clientTransactionID, "client-transaction-id", "", "Unique client-side ID for this ACH transaction attempt")
	cmd.Flags().StringVar(&amount, "amount", "", "Transaction amount as a decimal number, e.g. 102.05")
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Optional stable end-user ID from your system")
	cmd.Flags().BoolVar(&isRecurring, "is-recurring", false, "Whether the ACH transaction is part of a recurring schedule")
	cmd.Flags().StringVar(&defaultPaymentMethod, "default-payment-method", "", "Default payment method: SAME_DAY_ACH, STANDARD_ACH, or MULTIPLE_PAYMENT_METHODS")
	cmd.Flags().StringVar(&rulesetKey, "ruleset-key", "", "Optional Signal ruleset key from the Plaid Dashboard")
	cmd.Flags().StringVar(&prefix, "name-prefix", "", "User name prefix, e.g. Mr. or Ms.")
	cmd.Flags().StringVar(&givenName, "given-name", "", "User given name")
	cmd.Flags().StringVar(&middleName, "middle-name", "", "User middle name")
	cmd.Flags().StringVar(&familyName, "family-name", "", "User family name or surname")
	cmd.Flags().StringVar(&suffix, "name-suffix", "", "User name suffix, e.g. II")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "User phone number in E.164 format")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "User email address")
	cmd.Flags().StringVar(&city, "city", "", "User address city")
	cmd.Flags().StringVar(&region, "region", "", "User address region or state")
	cmd.Flags().StringVar(&street, "street", "", "User street address")
	cmd.Flags().StringVar(&postalCode, "postal-code", "", "User postal code")
	cmd.Flags().StringVar(&country, "country", "", "User address ISO 3166-1 alpha-2 country code")
	cmd.Flags().StringVar(&ipAddress, "ip-address", "", "End-user device IP address")
	cmd.Flags().StringVar(&userAgent, "user-agent", "", "End-user device user agent")
	return cmd
}

func newSignalDecisionReportCmd() *cobra.Command {
	var clientTransactionID string
	var daysFundsOnHold int
	var decisionOutcome, paymentMethod, amountInstantlyAvailable string
	var initiated bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "decision-report",
		Short: "Call /signal/decision/report",
		Long:  "Capability: write. Reports whether you initiated a previously evaluated ACH transaction.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"client_transaction_id": "<client-transaction-id>",
				"initiated":             true,
			}
			if handled, err := maybeWriteInfo(cmd, info, signalDocPath, template); handled || err != nil {
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
				"--client-transaction-id": {"client_transaction_id"},
				"--initiated":             {"initiated"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/signal/decision/report", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&clientTransactionID, "client-transaction-id", "", "Client transaction ID previously sent to signal evaluate")
	cmd.Flags().BoolVar(&initiated, "initiated", false, "Whether the ACH transaction was initiated")
	cmd.Flags().IntVar(&daysFundsOnHold, "days-funds-on-hold", 0, "Actual number of days funds were held before becoming available")
	cmd.Flags().StringVar(&decisionOutcome, "decision-outcome", "", "Decision outcome: APPROVE, REVIEW, REJECT, TAKE_OTHER_RISK_MEASURES, or NOT_EVALUATED")
	cmd.Flags().StringVar(&paymentMethod, "payment-method", "", "Payment method used: SAME_DAY_ACH, STANDARD_ACH, or MULTIPLE_PAYMENT_METHODS")
	cmd.Flags().StringVar(&amountInstantlyAvailable, "amount-instantly-available", "", "Amount made instantly available as a decimal number")
	return cmd
}

func newSignalReturnReportCmd() *cobra.Command {
	var clientTransactionID, returnCode, returnedAt string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "return-report",
		Short: "Call /signal/return/report",
		Long:  "Capability: write. Reports an ACH return for a transaction previously sent to Signal or balance evaluation.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"client_transaction_id": "<client-transaction-id>",
				"return_code":           "R01",
			}
			if handled, err := maybeWriteInfo(cmd, info, signalDocPath, template); handled || err != nil {
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
				"--client-transaction-id": {"client_transaction_id"},
				"--return-code":           {"return_code"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/signal/return/report", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&clientTransactionID, "client-transaction-id", "", "Client transaction ID previously sent to signal evaluate or balance get")
	cmd.Flags().StringVar(&returnCode, "return-code", "", "ACH return code, e.g. R01")
	cmd.Flags().StringVar(&returnedAt, "returned-at", "", "Optional ISO 8601 timestamp for when the return was received")
	return cmd
}

func newSignalPrepareCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "prepare",
		Short: "Call /signal/prepare",
		Long:  "Capability: write. Opts an existing Item into the Signal data collection flow used for Signal Transaction Scores.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, signalDocPath, template); handled || err != nil {
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
			if _, err := populateAccessToken(cmd, store, body, itemID, accessToken); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/signal/prepare", body)
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
	return cmd
}
