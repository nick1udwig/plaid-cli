package cmd

import (
	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

func newTransferOriginatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "originator",
		Short: "Transfer originator commands",
		Long:  "Commands for inspecting originator onboarding state and funding accounts for Transfer for Platforms.",
	}

	cmd.AddCommand(newTransferOriginatorGetCmd())
	cmd.AddCommand(newTransferOriginatorListCmd())
	cmd.AddCommand(newTransferOriginatorFundingAccountCreateCmd())

	return cmd
}

func applyTransferFundingAccountFlags(cmd *cobra.Command, store *state.Store, body map[string]any, itemID, accessToken, accountID, displayName string) error {
	record, err := populateAccessToken(cmd, store, body, itemID, accessToken)
	if err != nil {
		return err
	}
	if err := setBodyValue(body, body["access_token"], "funding_account", "access_token"); err != nil {
		return err
	}
	delete(body, "access_token")

	resolvedAccountID, err := resolveAccountID(record, accountID)
	if err != nil {
		return err
	}
	if err := setBodyValue(body, resolvedAccountID, "funding_account", "account_id"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "display-name", displayName, "funding_account", "display_name"); err != nil {
		return err
	}
	return nil
}

func newTransferOriginatorGetCmd() *cobra.Command {
	var originatorClientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/originator/get",
		Long:  "Capability: read. Retrieves onboarding status for a single transfer originator.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"originator_client_id": "<originator-client-id>"}
			if handled, err := maybeWriteInfo(cmd, info, transferPlatformDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--originator-client-id": {"originator_client_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/originator/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID to retrieve")
	return cmd
}

func newTransferOriginatorListCmd() *cobra.Command {
	var count, offset int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /transfer/originator/list",
		Long:  "Capability: read. Lists transfer originator onboarding records.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"count":  count,
				"offset": offset,
			}
			if handled, err := maybeWriteInfo(cmd, info, transferPlatformDocPath, template); handled || err != nil {
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
			if err := applyIntFlag(cmd, body, "offset", offset, "offset"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/originator/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of originators to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Originator list offset")
	return cmd
}

func newTransferOriginatorFundingAccountCreateCmd() *cobra.Command {
	var originatorClientID, itemID, accessToken, accountID, displayName string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "funding-account-create",
		Short: "Call /transfer/originator/funding_account/create",
		Long:  "Capability: write. Creates a funding account for a transfer originator.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"originator_client_id": "<originator-client-id>",
				"funding_account": map[string]any{
					"access_token": "<access-token>",
					"account_id":   "<account-id>",
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, transferPlatformDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyTransferFundingAccountFlags(cmd, store, body, itemID, accessToken, accountID, displayName); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--originator-client-id": {"originator_client_id"},
				"--access-token":         {"funding_account", "access_token"},
				"--account-id":           {"funding_account", "account_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/originator/funding_account/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID that owns this funding account")
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id to link as the funding account")
	cmd.Flags().StringVar(&displayName, "display-name", "", "Display name for the funding account in the Plaid dashboard")
	return cmd
}
