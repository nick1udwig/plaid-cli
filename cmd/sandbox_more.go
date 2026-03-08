package cmd

import (
	"errors"
	"fmt"
	"strings"

	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

type sandboxUserRefFlags struct {
	userID    string
	userToken string
}

func bindSandboxUserRefFlags(cmd *cobra.Command) *sandboxUserRefFlags {
	flags := &sandboxUserRefFlags{}
	cmd.Flags().StringVar(&flags.userID, "user-id", "", "Plaid user_id for the Sandbox user")
	cmd.Flags().StringVar(&flags.userToken, "user-token", "", "Legacy Plaid user_token for the Sandbox user")
	return flags
}

func applySandboxUserRefFlags(cmd *cobra.Command, body map[string]any, flags *sandboxUserRefFlags) error {
	if flags == nil {
		return nil
	}
	if err := applyStringFlag(cmd, body, "user-id", flags.userID, "user_id"); err != nil {
		return err
	}
	return applyStringFlag(cmd, body, "user-token", flags.userToken, "user_token")
}

func requireSandboxUserRef(body map[string]any) error {
	return requireAtLeastOneBodyField(body, map[string][]string{
		"--user-id":    {"user_id"},
		"--user-token": {"user_token"},
	})
}

func resolveStoredItemID(store *state.Store, itemID string) (string, error) {
	if itemID != "" {
		record, err := store.LoadItem(itemID)
		if err != nil {
			return "", err
		}
		return record.ItemID, nil
	}

	items, err := store.ListItems()
	if err != nil {
		return "", err
	}
	switch len(items) {
	case 0:
		return "", errors.New("no saved items found; run `plaid link connect` or provide --item-id in --body")
	case 1:
		return items[0].ItemID, nil
	default:
		ids := make([]string, 0, len(items))
		for _, item := range items {
			ids = append(ids, item.ItemID)
		}
		return "", fmt.Errorf("multiple saved items found; provide --item. available item_ids: %v", ids)
	}
}

func parseSandboxTransactionSpec(raw string) (map[string]any, error) {
	entry, err := loadRequestBody(raw)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func buildSandboxTransactionFromFlags(dateTransacted, datePosted, amount, description, currency string) (map[string]any, bool, error) {
	hasAny := strings.TrimSpace(dateTransacted) != "" ||
		strings.TrimSpace(datePosted) != "" ||
		strings.TrimSpace(amount) != "" ||
		strings.TrimSpace(description) != "" ||
		strings.TrimSpace(currency) != ""
	if !hasAny {
		return nil, false, nil
	}

	entry := map[string]any{}
	if strings.TrimSpace(dateTransacted) == "" {
		return nil, false, errors.New("--date-transacted is required when building a transaction from flags")
	}
	if strings.TrimSpace(datePosted) == "" {
		return nil, false, errors.New("--date-posted is required when building a transaction from flags")
	}
	if strings.TrimSpace(amount) == "" {
		return nil, false, errors.New("--amount is required when building a transaction from flags")
	}
	if strings.TrimSpace(description) == "" {
		return nil, false, errors.New("--description is required when building a transaction from flags")
	}

	if err := setBodyValue(entry, dateTransacted, "date_transacted"); err != nil {
		return nil, false, err
	}
	if err := setBodyValue(entry, datePosted, "date_posted"); err != nil {
		return nil, false, err
	}
	if err := setBodyValue(entry, description, "description"); err != nil {
		return nil, false, err
	}
	if strings.TrimSpace(currency) != "" {
		if err := setBodyValue(entry, currency, "iso_currency_code"); err != nil {
			return nil, false, err
		}
	}
	if err := applyDecimalStringFlag(&cobra.Command{}, entry, "amount", amount, "amount"); err != nil {
		return nil, false, err
	}

	return entry, true, nil
}

func newSandboxProcessorTokenCreateCmd() *cobra.Command {
	var institutionID, overrideUsername, overridePassword string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "processor-token-create",
		Short: "Call /sandbox/processor_token/create",
		Long:  "Capability: sandbox-write. Creates a Sandbox processor_token without running Link.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"institution_id": "ins_109508"}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "institution-id", institutionID, "institution_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "override-username", overrideUsername, "options", "override_username"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "override-password", overridePassword, "options", "override_password"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--institution-id": {"institution_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/processor_token/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&institutionID, "institution-id", "", "Sandbox institution_id")
	cmd.Flags().StringVar(&overrideUsername, "override-username", "", "Override the Sandbox username used when creating the Item")
	cmd.Flags().StringVar(&overridePassword, "override-password", "", "Override the Sandbox password used when creating the Item")
	return cmd
}

func newSandboxUserResetLoginCmd() *cobra.Command {
	var itemIDs []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userRefFlags *sandboxUserRefFlags

	cmd := &cobra.Command{
		Use:   "user-reset-login",
		Short: "Call /sandbox/user/reset_login",
		Long:  "Capability: sandbox-write. Forces one or more Sandbox user Items into ITEM_LOGIN_REQUIRED.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applySandboxUserRefFlags(cmd, body, userRefFlags); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "item-id", itemIDs, "item_ids"); err != nil {
				return err
			}
			if err := requireSandboxUserRef(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/user/reset_login", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	userRefFlags = bindSandboxUserRefFlags(cmd)
	cmd.Flags().StringSliceVar(&itemIDs, "item-id", nil, "Specific Plaid item_id values to reset (repeatable)")
	return cmd
}

func newSandboxItemSetVerificationStatusCmd() *cobra.Command {
	var itemID, accessToken, accountID, verificationStatus string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "item-set-verification-status",
		Short: "Call /sandbox/item/set_verification_status",
		Long:  "Capability: sandbox-write. Sets a Sandbox account verification status for Auth testing.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token":        "<access-token>",
				"account_id":          "<account-id>",
				"verification_status": "automatically_verified",
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "verification-status", verificationStatus, "verification_status"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--account-id":          {"account_id"},
				"--access-token":        {"access_token"},
				"--verification-status": {"verification_status"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/item/set_verification_status", body)
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
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id to update")
	cmd.Flags().StringVar(&verificationStatus, "verification-status", "", "Verification status to set, e.g. automatically_verified")
	return cmd
}

func newSandboxIncomeFireWebhookCmd() *cobra.Command {
	var savedItemID, explicitItemID, userID, webhook, webhookCode, verificationStatus string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "income-fire-webhook",
		Short: "Call /sandbox/income/fire_webhook",
		Long:  "Capability: sandbox-write. Fires an Income webhook for a Sandbox Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"item_id":      "<item-id>",
				"webhook":      "https://example.com/plaid/income",
				"webhook_code": "INCOME_VERIFICATION",
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if explicitItemID != "" {
				if err := applyStringFlag(cmd, body, "item-id", explicitItemID, "item_id"); err != nil {
					return err
				}
			} else if savedItemID != "" || !bodyHasValue(body, "item_id") {
				itemID, err := resolveStoredItemID(store, savedItemID)
				if err != nil {
					return err
				}
				if err := setBodyValue(body, itemID, "item_id"); err != nil {
					return err
				}
			}
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook-code", webhookCode, "webhook_code"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "verification-status", verificationStatus, "verification_status"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--item or --item-id": {"item_id"},
				"--webhook":           {"webhook"},
				"--webhook-code":      {"webhook_code"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/income/fire_webhook", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&savedItemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&explicitItemID, "item-id", "", "Explicit Plaid item_id override")
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id associated with the webhook")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive the sandbox Income webhook")
	cmd.Flags().StringVar(&webhookCode, "webhook-code", "", "Webhook code to fire, e.g. INCOME_VERIFICATION")
	cmd.Flags().StringVar(&verificationStatus, "verification-status", "", "Verification status to include, e.g. VERIFICATION_STATUS_PROCESSING_COMPLETE")
	return cmd
}

func newSandboxCRACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cra",
		Short: "Sandbox Plaid Check helpers",
		Long:  "Sandbox-only helpers for Plaid Check and Cash Flow Updates testing.",
	}
	cmd.AddCommand(newSandboxCRACashflowUpdatesUpdateCmd())
	return cmd
}

func newSandboxCRACashflowUpdatesUpdateCmd() *cobra.Command {
	var webhookCodes []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userRefFlags *sandboxUserRefFlags

	cmd := &cobra.Command{
		Use:   "cashflow-updates-update",
		Short: "Call /sandbox/cra/cashflow_updates/update",
		Long:  "Capability: sandbox-write. Triggers Cash Flow Updates webhook events for a Sandbox user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applySandboxUserRefFlags(cmd, body, userRefFlags); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "webhook-code", webhookCodes, "webhook_codes"); err != nil {
				return err
			}
			if err := requireSandboxUserRef(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/cra/cashflow_updates/update", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	userRefFlags = bindSandboxUserRefFlags(cmd)
	cmd.Flags().StringSliceVar(&webhookCodes, "webhook-code", nil, "Cash Flow Updates webhook code to simulate (repeatable)")
	return cmd
}

func newSandboxPaymentSimulateCmd() *cobra.Command {
	var paymentID, webhook, status string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "payment-simulate",
		Short: "Call /sandbox/payment/simulate",
		Long:  "Capability: sandbox-write. Simulates a Payment Initiation status transition in Sandbox.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"payment_id": "<payment-id>",
				"webhook":    "https://example.com/plaid/payment",
				"status":     "PAYMENT_STATUS_INITIATED",
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "status", status, "status"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--payment-id": {"payment_id"},
				"--webhook":    {"webhook"},
				"--status":     {"status"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/payment/simulate", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&paymentID, "payment-id", "", "Plaid payment_id to simulate")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive the sandbox payment status webhook")
	cmd.Flags().StringVar(&status, "status", "", "Payment status to simulate, e.g. PAYMENT_STATUS_EXECUTED")
	return cmd
}

func newSandboxTransactionsCreateCmd() *cobra.Command {
	var itemID, accessToken string
	var transactionSpecs []string
	var dateTransacted, datePosted, amount, description, currency string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "transactions-create",
		Short: "Call /sandbox/transactions/create",
		Long:  "Capability: sandbox-write. Creates custom Sandbox transactions for a dynamic transactions Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"transactions": []map[string]any{
					{
						"date_transacted":   "2026-03-08",
						"date_posted":       "2026-03-08",
						"amount":            12.34,
						"description":       "Coffee Shop",
						"iso_currency_code": "USD",
					},
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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

			entries := make([]map[string]any, 0, len(transactionSpecs)+1)
			for _, spec := range transactionSpecs {
				entry, err := parseSandboxTransactionSpec(spec)
				if err != nil {
					return fmt.Errorf("parse --transaction: %w", err)
				}
				entries = append(entries, entry)
			}
			if entry, ok, err := buildSandboxTransactionFromFlags(dateTransacted, datePosted, amount, description, currency); err != nil {
				return err
			} else if ok {
				entries = append(entries, entry)
			}
			if len(entries) > 0 {
				if err := setBodyValue(body, entries, "transactions"); err != nil {
					return err
				}
			}
			if err := requireBodyFields(body, map[string][]string{
				"--access-token":                       {"access_token"},
				"--transaction or --body.transactions": {"transactions"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transactions/create", body)
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
	cmd.Flags().StringSliceVar(&transactionSpecs, "transaction", nil, "Repeatable JSON object or @path describing a transaction to create")
	cmd.Flags().StringVar(&dateTransacted, "date-transacted", "", "Transaction date in YYYY-MM-DD for a single convenience transaction")
	cmd.Flags().StringVar(&datePosted, "date-posted", "", "Posted date in YYYY-MM-DD for a single convenience transaction")
	cmd.Flags().StringVar(&amount, "amount", "", "Transaction amount for a single convenience transaction")
	cmd.Flags().StringVar(&description, "description", "", "Transaction description for a single convenience transaction")
	cmd.Flags().StringVar(&currency, "currency", "", "ISO currency code for a single convenience transaction")
	return cmd
}
