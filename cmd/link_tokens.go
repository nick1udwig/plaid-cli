package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newLinkTokenCreateCmd() *cobra.Command {
	var clientName, language string
	var itemID, accessToken string
	var userID, clientUserID string
	var products, requiredIfSupportedProducts, optionalProducts, additionalConsentedProducts []string
	var countryCodes []string
	var webhook, redirectURI, linkCustomizationName string
	var phoneNumber, emailAddress, legalName string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "token-create",
		Short: "Call /link/token/create",
		Long:  "Capability: write. Creates a raw Link token for agent-managed Link flows. Use `link connect` for the opinionated Hosted Link browser flow.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"client_name":   "plaid-cli",
				"language":      "en",
				"country_codes": []string{"US"},
				"user": map[string]any{
					"client_user_id": "local-owner",
				},
				"products": []string{"auth"},
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/link/index.md", template); handled || err != nil {
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
			if itemID != "" || accessToken != "" {
				if _, err := populateAccessToken(cmd, store, body, itemID, accessToken); err != nil {
					return err
				}
			}

			resolvedClientName := clientName
			if resolvedClientName == "" {
				resolvedClientName = profile.ClientName
			}
			if err := applyStringFlag(cmd, body, "client-name", resolvedClientName, "client_name"); err != nil {
				return err
			}
			resolvedLanguage := language
			if resolvedLanguage == "" {
				resolvedLanguage = profile.Language
			}
			if err := applyStringFlag(cmd, body, "language", resolvedLanguage, "language"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "country-code", defaultCountryCodes(profile, countryCodes), "country_codes"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if !bodyHasValue(body, "user_id") {
				defaultClientUserID := clientUserID
				if defaultClientUserID == "" {
					defaultClientUserID = "local-owner"
				}
				if err := applyStringFlag(cmd, body, "client-user-id", defaultClientUserID, "user", "client_user_id"); err != nil {
					return err
				}
			}
			if err := applyStringFlag(cmd, body, "phone-number", phoneNumber, "user", "phone_number"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "email-address", emailAddress, "user", "email_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "legal-name", legalName, "user", "legal_name"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "product", products, "products"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "required-if-supported-product", requiredIfSupportedProducts, "required_if_supported_products"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "optional-product", optionalProducts, "optional_products"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "additional-consented-product", additionalConsentedProducts, "additional_consented_products"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "redirect-uri", redirectURI, "redirect_uri"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "link-customization-name", linkCustomizationName, "link_customization_name"); err != nil {
				return err
			}

			if !bodyHasValue(body, "user_id") && !bodyHasValue(body, "user", "client_user_id") {
				return errors.New("--client-user-id or --user-id is required")
			}
			if !bodyHasValue(body, "access_token") && !bodyHasValue(body, "products") {
				return errors.New("at least one --product is required unless creating an update-mode token with --item or --access-token")
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/link/token/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&clientName, "client-name", "", "Application name shown in Link; defaults to the saved app profile")
	cmd.Flags().StringVar(&language, "language", "", "Link language; defaults to the saved app profile")
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use for update-mode Link tokens")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override for update-mode Link tokens")
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id for new user API flows")
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Fallback client_user_id for the local owner; defaults to local-owner")
	cmd.Flags().StringSliceVar(&products, "product", nil, "Primary Plaid product to initialize in Link (repeatable)")
	cmd.Flags().StringSliceVar(&requiredIfSupportedProducts, "required-if-supported-product", nil, "Product to request only if the selected institution supports it (repeatable)")
	cmd.Flags().StringSliceVar(&optionalProducts, "optional-product", nil, "Best-effort optional product to initialize in Link (repeatable)")
	cmd.Flags().StringSliceVar(&additionalConsentedProducts, "additional-consented-product", nil, "Additional product to collect consent for without immediate billing (repeatable)")
	cmd.Flags().StringSliceVar(&countryCodes, "country-code", nil, "Country code for Link; defaults to the saved app profile")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Optional webhook URL for the Link token")
	cmd.Flags().StringVar(&redirectURI, "redirect-uri", "", "Optional redirect_uri for OAuth-capable flows")
	cmd.Flags().StringVar(&linkCustomizationName, "link-customization-name", "", "Optional Link customization name from the Plaid Dashboard")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Optional user phone number for returning-user or identity flows")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "Optional user email address")
	cmd.Flags().StringVar(&legalName, "legal-name", "", "Optional user legal name")
	return cmd
}

func newLinkTokenGetCmd() *cobra.Command {
	var linkToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "token-get",
		Short: "Call /link/token/get",
		Long:  "Capability: read. Retrieves the state and results of a previously created Link token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"link_token": "<link-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/link/index.md", template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "link-token", linkToken, "link_token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--link-token": {"link_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/link/token/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&linkToken, "link-token", "", "Plaid link_token returned by /link/token/create")
	return cmd
}
