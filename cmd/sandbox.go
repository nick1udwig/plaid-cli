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
	cmd.AddCommand(newSandboxItemResetLoginCmd())
	cmd.AddCommand(newSandboxItemFireWebhookCmd())
	cmd.AddCommand(newSandboxTransferCmd())
	return cmd
}

func newSandboxPublicTokenCreateCmd() *cobra.Command {
	var institutionID string
	var products []string
	info := bindInfoFlags(&cobra.Command{})

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

			body := map[string]any{
				"institution_id":   institutionID,
				"initial_products": products,
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
	cmd.Flags().StringVar(&institutionID, "institution-id", "", "Sandbox institution_id")
	cmd.Flags().StringSliceVar(&products, "product", nil, "Initial product to enable on the Sandbox Item (repeatable)")
	return cmd
}

func newSandboxItemResetLoginCmd() *cobra.Command {
	var itemID, accessToken string
	info := bindInfoFlags(&cobra.Command{})

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
			token, _, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/item/reset_login", map[string]any{"access_token": token})
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	return cmd
}

func newSandboxItemFireWebhookCmd() *cobra.Command {
	var itemID, accessToken, webhookCode string
	info := bindInfoFlags(&cobra.Command{})

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
			token, _, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/item/fire_webhook", map[string]any{
				"access_token": token,
				"webhook_code": webhookCode,
			})
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringVar(&webhookCode, "webhook-code", "", "Sandbox webhook code to fire")
	return cmd
}
