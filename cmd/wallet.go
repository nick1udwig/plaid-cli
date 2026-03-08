package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

const walletDocPath = "docs/plaid/api/products/virtual-accounts/index.md"

type walletAddressFlags struct {
	street     []string
	city       string
	postalCode string
	country    string
}

func newWalletCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wallet",
		Short: "Virtual Accounts wallet commands",
		Long:  "Wallet commands for Plaid virtual accounts, including e-wallet creation and wallet transactions.",
	}

	cmd.AddCommand(newWalletCreateCmd())
	cmd.AddCommand(newWalletGetCmd())
	cmd.AddCommand(newWalletListCmd())
	cmd.AddCommand(newWalletTransactionCmd())

	return cmd
}

func newWalletCreateCmd() *cobra.Command {
	var isoCurrencyCode string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /wallet/create",
		Long:  "Capability: write. Creates a Plaid virtual account wallet.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"iso_currency_code": "GBP",
			}
			if handled, err := maybeWriteInfo(cmd, info, walletDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "iso-currency-code", isoCurrencyCode, "iso_currency_code"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--iso-currency-code": {"iso_currency_code"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/wallet/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&isoCurrencyCode, "iso-currency-code", "", "Wallet ISO-4217 currency code, e.g. GBP or EUR")
	return cmd
}

func newWalletGetCmd() *cobra.Command {
	var walletID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /wallet/get",
		Long:  "Capability: read. Fetches a single virtual account wallet.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"wallet_id": "<wallet-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, walletDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "wallet-id", walletID, "wallet_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--wallet-id": {"wallet_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/wallet/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&walletID, "wallet-id", "", "Plaid wallet_id")
	return cmd
}

func newWalletListCmd() *cobra.Command {
	var isoCurrencyCode, cursor string
	var count int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /wallet/list",
		Long:  "Capability: read. Lists virtual account wallets in reverse creation order.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"count": 10,
			}
			if isoCurrencyCode != "" {
				template["iso_currency_code"] = isoCurrencyCode
			}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, walletDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "iso-currency-code", isoCurrencyCode, "iso_currency_code"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/wallet/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&isoCurrencyCode, "iso-currency-code", "", "Filter wallets by ISO-4217 currency code")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Cursor returned by a previous wallet list call")
	cmd.Flags().IntVar(&count, "count", 10, "Number of wallets to return")
	return cmd
}

func newWalletTransactionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction",
		Short: "Wallet transaction commands",
		Long:  "Commands for executing and listing transactions against Plaid virtual account wallets.",
	}

	cmd.AddCommand(newWalletTransactionExecuteCmd())
	cmd.AddCommand(newWalletTransactionGetCmd())
	cmd.AddCommand(newWalletTransactionListCmd())

	return cmd
}

func newWalletTransactionExecuteCmd() *cobra.Command {
	var idempotencyKey, walletID, counterpartyName string
	var counterpartyIBAN, counterpartyBACSAccount, counterpartyBACSSortCode string
	var counterpartyDateOfBirth string
	var amountCurrency, amountValue, reference string
	var fundSourceFullName, fundSourceAccountNumber, fundSourceBIC string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var counterpartyAddress walletAddressFlags
	var fundSourceAddress walletAddressFlags

	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Call /wallet/transaction/execute",
		Long:  "Capability: write. Executes a payout or related wallet transaction from a virtual account.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"idempotency_key": "<idempotency-key>",
				"wallet_id":       "<wallet-id>",
				"counterparty": map[string]any{
					"name": "<counterparty-name>",
					"numbers": map[string]any{
						"bacs": map[string]any{
							"account":   "12345678",
							"sort_code": "123456",
						},
					},
				},
				"amount": map[string]any{
					"iso_currency_code": "GBP",
					"value":             1.23,
				},
				"reference": "ABC12345",
			}
			if handled, err := maybeWriteInfo(cmd, info, walletDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "wallet-id", walletID, "wallet_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "counterparty-name", counterpartyName, "counterparty", "name"); err != nil {
				return err
			}
			if err := applyWalletCounterpartyNumbers(cmd, body, counterpartyIBAN, counterpartyBACSAccount, counterpartyBACSSortCode); err != nil {
				return err
			}
			if err := applyWalletAddress(cmd, body, "counterparty-address", counterpartyAddress, "counterparty", "address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "counterparty-date-of-birth", counterpartyDateOfBirth, "counterparty", "date_of_birth"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount-currency", amountCurrency, "amount", "iso_currency_code"); err != nil {
				return err
			}
			if err := applyDecimalStringFlag(cmd, body, "amount-value", amountValue, "amount", "value"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "reference", reference, "reference"); err != nil {
				return err
			}
			if err := applyWalletFundSource(cmd, body, fundSourceFullName, fundSourceAddress, fundSourceAccountNumber, fundSourceBIC); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--idempotency-key":   {"idempotency_key"},
				"--wallet-id":         {"wallet_id"},
				"--counterparty-name": {"counterparty", "name"},
				"--amount-currency":   {"amount", "iso_currency_code"},
				"--amount-value":      {"amount", "value"},
				"--reference":         {"reference"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/wallet/transaction/execute", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Unique idempotency key for this wallet transaction")
	cmd.Flags().StringVar(&walletID, "wallet-id", "", "Plaid wallet_id to debit")
	cmd.Flags().StringVar(&counterpartyName, "counterparty-name", "", "Counterparty name")
	cmd.Flags().StringVar(&counterpartyIBAN, "counterparty-iban", "", "Counterparty IBAN")
	cmd.Flags().StringVar(&counterpartyBACSAccount, "counterparty-bacs-account", "", "Counterparty BACS account number")
	cmd.Flags().StringVar(&counterpartyBACSSortCode, "counterparty-bacs-sort-code", "", "Counterparty BACS sort code")
	cmd.Flags().StringSliceVar(&counterpartyAddress.street, "counterparty-address-street", nil, "Counterparty street lines (repeatable, 1-2)")
	cmd.Flags().StringVar(&counterpartyAddress.city, "counterparty-address-city", "", "Counterparty city")
	cmd.Flags().StringVar(&counterpartyAddress.postalCode, "counterparty-address-postal-code", "", "Counterparty postal code")
	cmd.Flags().StringVar(&counterpartyAddress.country, "counterparty-address-country", "", "Counterparty ISO 3166-1 alpha-2 country code")
	cmd.Flags().StringVar(&counterpartyDateOfBirth, "counterparty-date-of-birth", "", "Counterparty birthdate in YYYY-MM-DD format")
	cmd.Flags().StringVar(&amountCurrency, "amount-currency", "", "Transaction ISO-4217 currency code, e.g. GBP or EUR")
	cmd.Flags().StringVar(&amountValue, "amount-value", "", "Transaction amount as a decimal string")
	cmd.Flags().StringVar(&reference, "reference", "", "Unique alphanumeric payment reference")
	cmd.Flags().StringVar(&fundSourceFullName, "fund-source-full-name", "", "Originating fund source full name")
	cmd.Flags().StringSliceVar(&fundSourceAddress.street, "fund-source-address-street", nil, "Originating fund source street lines (repeatable, 1-2)")
	cmd.Flags().StringVar(&fundSourceAddress.city, "fund-source-address-city", "", "Originating fund source city")
	cmd.Flags().StringVar(&fundSourceAddress.postalCode, "fund-source-address-postal-code", "", "Originating fund source postal code")
	cmd.Flags().StringVar(&fundSourceAddress.country, "fund-source-address-country", "", "Originating fund source ISO 3166-1 alpha-2 country code")
	cmd.Flags().StringVar(&fundSourceAccountNumber, "fund-source-account-number", "", "Originating fund source account number")
	cmd.Flags().StringVar(&fundSourceBIC, "fund-source-bic", "", "Originating fund source BIC/SWIFT code")
	return cmd
}

func newWalletTransactionGetCmd() *cobra.Command {
	var transactionID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /wallet/transaction/get",
		Long:  "Capability: read. Fetches a single wallet transaction.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"transaction_id": "<transaction-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, walletDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "transaction-id", transactionID, "transaction_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--transaction-id": {"transaction_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/wallet/transaction/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&transactionID, "transaction-id", "", "Plaid wallet transaction_id")
	return cmd
}

func newWalletTransactionListCmd() *cobra.Command {
	var walletID, cursor, startTime, endTime string
	var count int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /wallet/transaction/list",
		Long:  "Capability: read. Lists transactions for a virtual account wallet.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"wallet_id": "<wallet-id>",
				"count":     10,
			}
			if cursor != "" {
				template["cursor"] = cursor
			}
			options := map[string]any{}
			if startTime != "" {
				options["start_time"] = startTime
			}
			if endTime != "" {
				options["end_time"] = endTime
			}
			if len(options) > 0 {
				template["options"] = options
			}
			if handled, err := maybeWriteInfo(cmd, info, walletDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "wallet-id", walletID, "wallet_id"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "start-time", startTime, "options", "start_time"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-time", endTime, "options", "end_time"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--wallet-id": {"wallet_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/wallet/transaction/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&walletID, "wallet-id", "", "Plaid wallet_id")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Cursor returned by a previous wallet transaction list call")
	cmd.Flags().IntVar(&count, "count", 10, "Number of wallet transactions to return")
	cmd.Flags().StringVar(&startTime, "start-time", "", "Filter start time in ISO 8601 format")
	cmd.Flags().StringVar(&endTime, "end-time", "", "Filter end time in ISO 8601 format")
	return cmd
}

func applyWalletCounterpartyNumbers(cmd *cobra.Command, body map[string]any, iban, bacsAccount, bacsSortCode string) error {
	if err := applyStringFlag(cmd, body, "counterparty-iban", iban, "counterparty", "numbers", "international", "iban"); err != nil {
		return err
	}
	bacsChanged := anyFlagChanged(cmd, "counterparty-bacs-account", "counterparty-bacs-sort-code")
	if bacsChanged || ((bacsAccount != "" || bacsSortCode != "") && !bodyHasValue(body, "counterparty", "numbers", "bacs")) {
		if bacsAccount == "" || bacsSortCode == "" {
			return errors.New("--counterparty-bacs-account and --counterparty-bacs-sort-code must be provided together")
		}
		if err := setBodyValue(body, map[string]any{
			"account":   bacsAccount,
			"sort_code": bacsSortCode,
		}, "counterparty", "numbers", "bacs"); err != nil {
			return err
		}
	}

	hasIBAN := bodyHasValue(body, "counterparty", "numbers", "international", "iban")
	hasBACS := bodyHasValue(body, "counterparty", "numbers", "bacs")
	if hasIBAN == hasBACS {
		return errors.New("provide exactly one of --counterparty-iban or the BACS counterparty flags")
	}
	return nil
}

func applyWalletAddress(cmd *cobra.Command, body map[string]any, label string, flags walletAddressFlags, path ...string) error {
	flagStreet := fmt.Sprintf("%s-street", label)
	flagCity := fmt.Sprintf("%s-city", label)
	flagPostalCode := fmt.Sprintf("%s-postal-code", label)
	flagCountry := fmt.Sprintf("%s-country", label)

	shouldSet := anyFlagChanged(cmd, flagStreet, flagCity, flagPostalCode, flagCountry) ||
		((len(flags.street) > 0 || flags.city != "" || flags.postalCode != "" || flags.country != "") && !bodyHasValue(body, path...))
	if !shouldSet {
		return nil
	}

	if len(flags.street) == 0 || flags.city == "" || flags.postalCode == "" || flags.country == "" {
		return fmt.Errorf("%s requires street, city, postal code, and country together", label)
	}

	address := map[string]any{
		"street":      flags.street,
		"city":        flags.city,
		"postal_code": flags.postalCode,
		"country":     flags.country,
	}
	return setBodyValue(body, address, path...)
}

func applyWalletFundSource(cmd *cobra.Command, body map[string]any, fullName string, address walletAddressFlags, accountNumber, bic string) error {
	shouldSet := anyFlagChanged(
		cmd,
		"fund-source-full-name",
		"fund-source-address-street",
		"fund-source-address-city",
		"fund-source-address-postal-code",
		"fund-source-address-country",
		"fund-source-account-number",
		"fund-source-bic",
	) || ((fullName != "" || len(address.street) > 0 || address.city != "" || address.postalCode != "" || address.country != "" || accountNumber != "" || bic != "") &&
		!bodyHasValue(body, "originating_fund_source"))
	if !shouldSet {
		return nil
	}

	if fullName == "" || accountNumber == "" || bic == "" {
		return errors.New("originating fund source requires --fund-source-full-name, --fund-source-account-number, and --fund-source-bic")
	}
	if err := applyWalletAddress(cmd, body, "fund-source-address", address, "originating_fund_source", "address"); err != nil {
		return err
	}
	if err := setBodyValue(body, fullName, "originating_fund_source", "full_name"); err != nil {
		return err
	}
	if err := setBodyValue(body, accountNumber, "originating_fund_source", "account_number"); err != nil {
		return err
	}
	return setBodyValue(body, bic, "originating_fund_source", "bic")
}
