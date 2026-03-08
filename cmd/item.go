package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newItemCmd(_ *Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "item",
		Short: "Inspect and manage Plaid Items",
		Long:  "Item commands. Includes local-read, read, write, and admin-style operations depending on the subcommand.",
	}

	cmd.AddCommand(newItemListCmd())
	cmd.AddCommand(newItemGetCmd())
	cmd.AddCommand(newItemRemoveCmd())
	cmd.AddCommand(newItemPublicTokenExchangeCmd())
	cmd.AddCommand(newItemInvalidateAccessTokenCmd())
	cmd.AddCommand(newItemWebhookUpdateCmd())

	return cmd
}

func newItemListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List locally saved Items",
		Long:  "Capability: local-read. Lists saved Item records from ~/.plaid-cli without calling the Plaid API.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			store, err := getStore(cmd)
			if err != nil {
				return err
			}
			items, err := store.ListItems()
			if err != nil {
				return err
			}
			return writeJSON(cmd, map[string]any{"items": items})
		},
	}
}

func newItemGetCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve an Item from the Plaid API",
		Long:  "Capability: read. Retrieves Item metadata from the Plaid API.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/items/index.md", map[string]any{
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

			resp, err := client.Call(ctx, "/item/get", body)
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

func newItemRemoveCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove an Item from Plaid and local state",
		Long:  "Capability: write. Removes the Item remotely via /item/remove. If the Item was loaded from local state, its local record is deleted too.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/items/index.md", map[string]any{
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
			record, err := populateAccessToken(cmd, store, body, itemID, accessToken)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()

			resp, err := client.Call(ctx, "/item/remove", body)
			if err != nil {
				return err
			}

			localDeleted := false
			if record != nil {
				if err := store.DeleteItem(record.ItemID); err != nil {
					return err
				}
				localDeleted = true
			}
			resp["local_item_deleted"] = localDeleted
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to remove")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	return cmd
}

func newItemPublicTokenExchangeCmd() *cobra.Command {
	var publicToken, linkToken string
	var products []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "public-token-exchange",
		Short: "Exchange a Link public token for an access token",
		Long:  "Capability: write. Calls /item/public_token/exchange and saves the resulting Item locally under ~/.plaid-cli/items.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"public_token": "<public-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/items/index.md", template); handled || err != nil {
				return err
			}

			store, profile, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "public-token", publicToken, "public_token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--public-token": {"public_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/item/public_token/exchange", body)
			if err != nil {
				return err
			}

			accessTokenValue, _ := resp["access_token"].(string)
			if accessTokenValue == "" {
				return errors.New("Plaid response did not include access_token")
			}

			record, err := saveItemFromAccessToken(ctx, cmd, store, client, accessTokenValue, linkToken, products, defaultCountryCodes(profile, nil))
			if err != nil {
				return err
			}

			resp["saved_item_record"] = store.ItemPath(record.ItemID)
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&publicToken, "public-token", "", "Plaid public_token returned from Link or Sandbox")
	cmd.Flags().StringVar(&linkToken, "link-token", "", "Optional link_token to save alongside the Item metadata")
	cmd.Flags().StringSliceVar(&products, "product", nil, "Optional product metadata to save with the local Item record (repeatable)")
	return cmd
}

func newItemInvalidateAccessTokenCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "access-token-invalidate",
		Short: "Rotate an Item access token",
		Long:  "Capability: write. Calls /item/access_token/invalidate and updates the local saved Item when possible.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/items/index.md", map[string]any{
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
			record, err := populateAccessToken(cmd, store, body, itemID, accessToken)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()

			resp, err := client.Call(ctx, "/item/access_token/invalidate", body)
			if err != nil {
				return err
			}

			newToken, _ := resp["new_access_token"].(string)
			if record != nil && newToken != "" {
				record.AccessToken = newToken
				if err := store.SaveItem(*record); err != nil {
					return err
				}
			}

			resp["local_item_updated"] = record != nil && newToken != ""
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to update")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	return cmd
}

func newItemWebhookUpdateCmd() *cobra.Command {
	var itemID, accessToken, webhookURL string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "webhook-update",
		Short: "Update an Item webhook URL",
		Long:  "Capability: admin. Calls /item/webhook/update for the selected Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/items/index.md", map[string]any{
				"access_token": "<access-token>",
				"webhook":      "https://example.com/plaid",
			}); handled || err != nil {
				return err
			}
			if webhookURL == "" {
				return errors.New("--webhook-url is required")
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
			if err := applyStringFlag(cmd, body, "webhook-url", webhookURL, "webhook"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--webhook-url": {"webhook"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/item/webhook/update", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to update")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringVar(&webhookURL, "webhook-url", "", "Webhook URL to set on the Item")
	return cmd
}
