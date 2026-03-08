package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const layerDocPath = "docs/plaid/api/products/layer/index.md"

func newLayerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "layer",
		Short: "Layer product commands",
		Long:  "Layer commands for creating Layer Link sessions and resolving public tokens into user-permissioned account data.",
	}

	cmd.AddCommand(newLayerSessionTokenCreateCmd())
	cmd.AddCommand(newLayerUserAccountSessionGetCmd())

	return cmd
}

func newLayerSessionTokenCreateCmd() *cobra.Command {
	var templateID, clientUserID, userID, redirectURI, androidPackageName, webhook string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "session-token-create",
		Short: "Call /session/token/create",
		Long:  "Capability: write. Creates a Link token for a Layer session.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"template_id": "<template-id>",
				"user": map[string]any{
					"client_user_id": "<client-user-id>",
				},
				"redirect_uri": "https://example.com/oauth/callback",
			}
			if handled, err := maybeWriteInfo(cmd, info, layerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "template-id", templateID, "template_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "client-user-id", clientUserID, "user", "client_user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "redirect-uri", redirectURI, "redirect_uri"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "android-package-name", androidPackageName, "android_package_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--template-id": {"template_id"},
			}); err != nil {
				return err
			}
			if err := requireAtLeastOneBodyField(body, map[string][]string{
				"--client-user-id": {"user", "client_user_id"},
				"--user-id":        {"user_id"},
			}); err != nil {
				return err
			}

			hasRedirect := bodyHasValue(body, "redirect_uri")
			hasAndroidPackage := bodyHasValue(body, "android_package_name")
			if hasRedirect == hasAndroidPackage {
				return errors.New("provide exactly one of --redirect-uri or --android-package-name")
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/session/token/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&templateID, "template-id", "", "Layer template ID from the Plaid Dashboard")
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable client-side user identifier")
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id from /user/create")
	cmd.Flags().StringVar(&redirectURI, "redirect-uri", "", "Redirect URI for browser-based Layer flows")
	cmd.Flags().StringVar(&androidPackageName, "android-package-name", "", "Android package name for Android Layer flows")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Optional per-session webhook URL")
	return cmd
}

func newLayerUserAccountSessionGetCmd() *cobra.Command {
	var publicToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "user-account-session-get",
		Short: "Call /user_account/session/get",
		Long:  "Capability: read. Exchanges a Layer public token for the permissioned identity and Item data.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"public_token": "<public-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, layerDocPath, template); handled || err != nil {
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
			resp, err := client.Call(ctx, "/user_account/session/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&publicToken, "public-token", "", "Public token returned by a Layer session")
	return cmd
}
