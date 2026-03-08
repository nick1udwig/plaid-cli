package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newSandboxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sandbox",
		Short: "Sandbox helpers for testing Plaid integrations",
		Long:  "Sandbox-only commands for testing write paths and webhook behavior.",
	}
	cmd.AddCommand(newSandboxPublicTokenCreateCmd())
	cmd.AddCommand(newSandboxProcessorTokenCreateCmd())
	cmd.AddCommand(newSandboxItemResetLoginCmd())
	cmd.AddCommand(newSandboxUserResetLoginCmd())
	cmd.AddCommand(newSandboxItemFireWebhookCmd())
	cmd.AddCommand(newSandboxItemSetVerificationStatusCmd())
	cmd.AddCommand(newSandboxIncomeFireWebhookCmd())
	cmd.AddCommand(newSandboxCRACmd())
	cmd.AddCommand(newSandboxPaymentSimulateCmd())
	cmd.AddCommand(newSandboxTransactionsCreateCmd())
	cmd.AddCommand(newSandboxTransferCmd())
	return cmd
}

func newSandboxPublicTokenCreateCmd() *cobra.Command {
	var institutionID string
	var products []string
	var webhook string
	var overrideUsername string
	var overridePassword string
	var userID string
	var userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "public-token-create",
		Short: "Call /sandbox/public_token/create",
		Long:  "Capability: sandbox-write. Creates a Sandbox public token for a test institution.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"institution_id": "ins_109508",
				"initial_products": []string{
					"auth",
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/sandbox/index.md", template); handled || err != nil {
				return err
			}
			if institutionID == "" {
				return errors.New("--institution-id is required")
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}

			if len(products) == 0 {
				products = []string{"auth"}
			}

			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "institution-id", institutionID, "institution_id"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "product", products, "initial_products"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "options", "webhook"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "override-username", overrideUsername, "options", "override_username"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "override-password", overridePassword, "options", "override_password"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-token", userToken, "user_token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--institution-id": {"institution_id"},
				"--product":        {"initial_products"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/public_token/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&institutionID, "institution-id", "", "Sandbox institution_id")
	cmd.Flags().StringSliceVar(&products, "product", nil, "Initial product to enable on the Sandbox Item (repeatable)")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Optional webhook URL to attach to the Sandbox Item")
	cmd.Flags().StringVar(&overrideUsername, "override-username", "", "Override the Sandbox username used when creating the Item")
	cmd.Flags().StringVar(&overridePassword, "override-password", "", "Override the Sandbox password used when creating the Item")
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id to associate with the Sandbox Item")
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token to associate with the Sandbox Item")
	return cmd
}

func newSandboxItemResetLoginCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "item-reset-login",
		Short: "Call /sandbox/item/reset_login",
		Long:  "Capability: sandbox. Forces an Item into ITEM_LOGIN_REQUIRED for testing.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/sandbox/index.md", map[string]any{
				"access_token": "<access-token>",
			}); handled || err != nil {
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
			resp, err := client.Call(ctx, "/sandbox/item/reset_login", body)
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

func newSandboxItemFireWebhookCmd() *cobra.Command {
	var itemID, accessToken, webhookType, webhookCode string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "item-fire-webhook",
		Short: "Call /sandbox/item/fire_webhook",
		Long:  "Capability: sandbox. Fires a test webhook for the selected Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"webhook_code": "DEFAULT_UPDATE",
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/sandbox/index.md", template); handled || err != nil {
				return err
			}
			if webhookCode == "" {
				return errors.New("--webhook-code is required")
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
			if err := applyStringFlag(cmd, body, "webhook-type", webhookType, "webhook_type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook-code", webhookCode, "webhook_code"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--webhook-code": {"webhook_code"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/item/fire_webhook", body)
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
	cmd.Flags().StringVar(&webhookType, "webhook-type", "", "Sandbox webhook type such as ITEM, AUTH, TRANSACTIONS, or ASSETS")
	cmd.Flags().StringVar(&webhookCode, "webhook-code", "", "Sandbox webhook code to fire")
	return cmd
}
