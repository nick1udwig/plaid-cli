package cmd

import "github.com/spf13/cobra"

const partnerDocPath = "docs/plaid/api/partner/index.md"

func newPartnerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "partner",
		Short: "Reseller partner commands",
		Long:  "Partner commands for Plaid reseller workflows. These are admin-style endpoints for managing end customers.",
	}

	cmd.AddCommand(newPartnerCustomerCmd())

	return cmd
}

func newPartnerCustomerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "customer",
		Short: "End customer commands",
		Long:  "Reseller end-customer lifecycle commands.",
	}

	cmd.AddCommand(newPartnerCustomerCreateCmd())
	cmd.AddCommand(newPartnerCustomerGetCmd())
	cmd.AddCommand(newPartnerCustomerOAuthInstitutionsGetCmd())
	cmd.AddCommand(newPartnerCustomerEnableCmd())
	cmd.AddCommand(newPartnerCustomerRemoveCmd())

	return cmd
}

func newPartnerCustomerCreateCmd() *cobra.Command {
	var companyName, legalEntityName, website, applicationName string
	var addressStreet, addressCity, addressRegion, addressPostalCode, addressCountryCode string
	var technicalContactGivenName, technicalContactFamilyName, technicalContactEmail string
	var billingContactGivenName, billingContactFamilyName, billingContactEmail string
	var supportEmail, supportPhoneNumber, supportContactURL, supportLinkUpdateURL string
	var products, redirectURIs []string
	var isDiligenceAttested, createLinkCustomization, isBankAddendumCompleted bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /partner/customer/create",
		Long:  "Capability: admin. Creates a reseller end customer. Use `--body` for the longer-tail registration fields such as assets under management or logo data.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"company_name":               "Example Company",
				"is_diligence_attested":      true,
				"legal_entity_name":          "Example Company LLC",
				"website":                    "https://example.com",
				"application_name":           "Example App",
				"is_bank_addendum_completed": true,
				"address": map[string]any{
					"street":       "123 Main St",
					"city":         "New York",
					"region":       "NY",
					"postal_code":  "10001",
					"country_code": "US",
				},
			}
			if len(products) > 0 {
				template["products"] = products
			}
			if createLinkCustomization {
				template["create_link_customization"] = true
			}
			if len(redirectURIs) > 0 {
				template["redirect_uris"] = redirectURIs
			}
			if handled, err := maybeWriteInfo(cmd, info, partnerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "company-name", companyName, "company_name"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "is-diligence-attested", isDiligenceAttested, "is_diligence_attested"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "product", products, "products"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "create-link-customization", createLinkCustomization, "create_link_customization"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "legal-entity-name", legalEntityName, "legal_entity_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "website", website, "website"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "application-name", applicationName, "application_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "address-street", addressStreet, "address", "street"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "address-city", addressCity, "address", "city"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "address-region", addressRegion, "address", "region"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "address-postal-code", addressPostalCode, "address", "postal_code"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "address-country-code", addressCountryCode, "address", "country_code"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "is-bank-addendum-completed", isBankAddendumCompleted, "is_bank_addendum_completed"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "technical-contact-given-name", technicalContactGivenName, "technical_contact", "given_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "technical-contact-family-name", technicalContactFamilyName, "technical_contact", "family_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "technical-contact-email", technicalContactEmail, "technical_contact", "email"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "billing-contact-given-name", billingContactGivenName, "billing_contact", "given_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "billing-contact-family-name", billingContactFamilyName, "billing_contact", "family_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "billing-contact-email", billingContactEmail, "billing_contact", "email"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "support-email", supportEmail, "customer_support_info", "email"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "support-phone-number", supportPhoneNumber, "customer_support_info", "phone_number"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "support-contact-url", supportContactURL, "customer_support_info", "contact_url"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "support-link-update-url", supportLinkUpdateURL, "customer_support_info", "link_update_url"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "redirect-uri", redirectURIs, "redirect_uris"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--company-name":               {"company_name"},
				"--is-diligence-attested":      {"is_diligence_attested"},
				"--legal-entity-name":          {"legal_entity_name"},
				"--website":                    {"website"},
				"--application-name":           {"application_name"},
				"--is-bank-addendum-completed": {"is_bank_addendum_completed"},
				"--address-street":             {"address", "street"},
				"--address-city":               {"address", "city"},
				"--address-region":             {"address", "region"},
				"--address-postal-code":        {"address", "postal_code"},
				"--address-country-code":       {"address", "country_code"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/partner/customer/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&companyName, "company-name", "", "Display name for the end customer in the Plaid Dashboard")
	cmd.Flags().BoolVar(&isDiligenceAttested, "is-diligence-attested", false, "Whether diligence attestation is complete for the end customer")
	cmd.Flags().StringSliceVar(&products, "product", nil, "Product to enable for the end customer (repeatable)")
	cmd.Flags().BoolVar(&createLinkCustomization, "create-link-customization", false, "Copy the partner's default Link customization to the end customer")
	cmd.Flags().StringVar(&legalEntityName, "legal-entity-name", "", "Legal entity name shared with financial institutions")
	cmd.Flags().StringVar(&website, "website", "", "End customer website")
	cmd.Flags().StringVar(&applicationName, "application-name", "", "Application name shown to end users during Link")
	cmd.Flags().StringVar(&addressStreet, "address-street", "", "End customer street address")
	cmd.Flags().StringVar(&addressCity, "address-city", "", "End customer address city")
	cmd.Flags().StringVar(&addressRegion, "address-region", "", "End customer address region or state")
	cmd.Flags().StringVar(&addressPostalCode, "address-postal-code", "", "End customer address postal code")
	cmd.Flags().StringVar(&addressCountryCode, "address-country-code", "", "End customer address country code")
	cmd.Flags().BoolVar(&isBankAddendumCompleted, "is-bank-addendum-completed", false, "Whether the bank addendum has been forwarded to the end customer")
	cmd.Flags().StringVar(&technicalContactGivenName, "technical-contact-given-name", "", "Technical contact first name")
	cmd.Flags().StringVar(&technicalContactFamilyName, "technical-contact-family-name", "", "Technical contact last name")
	cmd.Flags().StringVar(&technicalContactEmail, "technical-contact-email", "", "Technical contact email")
	cmd.Flags().StringVar(&billingContactGivenName, "billing-contact-given-name", "", "Billing contact first name")
	cmd.Flags().StringVar(&billingContactFamilyName, "billing-contact-family-name", "", "Billing contact last name")
	cmd.Flags().StringVar(&billingContactEmail, "billing-contact-email", "", "Billing contact email")
	cmd.Flags().StringVar(&supportEmail, "support-email", "", "Public support email for Plaid Portal")
	cmd.Flags().StringVar(&supportPhoneNumber, "support-phone-number", "", "Public support phone number for Plaid Portal")
	cmd.Flags().StringVar(&supportContactURL, "support-contact-url", "", "Public support URL for Plaid Portal")
	cmd.Flags().StringVar(&supportLinkUpdateURL, "support-link-update-url", "", "Public link update URL for Plaid Portal")
	cmd.Flags().StringSliceVar(&redirectURIs, "redirect-uri", nil, "OAuth redirect URI for the end customer (repeatable)")
	return cmd
}

func newPartnerCustomerGetCmd() *cobra.Command {
	var endCustomerClientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /partner/customer/get",
		Long:  "Capability: read. Retrieves a reseller end customer by client ID.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"end_customer_client_id": "<end-customer-client-id>"}
			if handled, err := maybeWriteInfo(cmd, info, partnerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "end-customer-client-id", endCustomerClientID, "end_customer_client_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--end-customer-client-id": {"end_customer_client_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/partner/customer/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&endCustomerClientID, "end-customer-client-id", "", "The end customer's Plaid client_id")
	return cmd
}

func newPartnerCustomerOAuthInstitutionsGetCmd() *cobra.Command {
	var endCustomerClientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "oauth-institutions-get",
		Short: "Call /partner/customer/oauth_institutions/get",
		Long:  "Capability: read. Retrieves OAuth institution registration details for a reseller end customer.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"end_customer_client_id": "<end-customer-client-id>"}
			if handled, err := maybeWriteInfo(cmd, info, partnerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "end-customer-client-id", endCustomerClientID, "end_customer_client_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--end-customer-client-id": {"end_customer_client_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/partner/customer/oauth_institutions/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&endCustomerClientID, "end-customer-client-id", "", "The end customer's Plaid client_id")
	return cmd
}

func newPartnerCustomerEnableCmd() *cobra.Command {
	var endCustomerClientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Call /partner/customer/enable",
		Long:  "Capability: admin. Enables an end customer in Production and returns its production secret.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"end_customer_client_id": "<end-customer-client-id>"}
			if handled, err := maybeWriteInfo(cmd, info, partnerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "end-customer-client-id", endCustomerClientID, "end_customer_client_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--end-customer-client-id": {"end_customer_client_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/partner/customer/enable", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&endCustomerClientID, "end-customer-client-id", "", "The end customer's Plaid client_id")
	return cmd
}

func newPartnerCustomerRemoveCmd() *cobra.Command {
	var endCustomerClientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Call /partner/customer/remove",
		Long:  "Capability: admin. Removes an end customer that has not yet been enabled in full Production.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"end_customer_client_id": "<end-customer-client-id>"}
			if handled, err := maybeWriteInfo(cmd, info, partnerDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "end-customer-client-id", endCustomerClientID, "end_customer_client_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--end-customer-client-id": {"end_customer_client_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/partner/customer/remove", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&endCustomerClientID, "end-customer-client-id", "", "The end customer's Plaid client_id")
	return cmd
}
