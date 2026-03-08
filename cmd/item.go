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
	info := bindInfoFlags(&cobra.Command{})

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
			token, _, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()

			resp, err := client.GetItem(ctx, token)
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

func newItemRemoveCmd() *cobra.Command {
	var itemID, accessToken string
	info := bindInfoFlags(&cobra.Command{})

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
			token, record, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()

			resp, err := client.Call(ctx, "/item/remove", map[string]any{"access_token": token})
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
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to remove")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	return cmd
}

func newItemInvalidateAccessTokenCmd() *cobra.Command {
	var itemID, accessToken string
	info := bindInfoFlags(&cobra.Command{})

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
			token, record, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()

			resp, err := client.Call(ctx, "/item/access_token/invalidate", map[string]any{"access_token": token})
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
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to update")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	return cmd
}

func newItemWebhookUpdateCmd() *cobra.Command {
	var itemID, accessToken, webhookURL string
	info := bindInfoFlags(&cobra.Command{})

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
			token, _, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()

			resp, err := client.Call(ctx, "/item/webhook/update", map[string]any{
				"access_token": token,
				"webhook":      webhookURL,
			})
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to update")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringVar(&webhookURL, "webhook-url", "", "Webhook URL to set on the Item")
	return cmd
}
