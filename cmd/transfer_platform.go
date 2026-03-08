package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

type transferRequirementSubmissionFlags struct {
	requirementType string
	value           string
	personID        string
}

func parseTransferRequirementSubmissionSpec(raw string) (map[string]any, error) {
	return loadRequestBody(raw)
}

func buildTransferRequirementSubmissionFromFlags(flags transferRequirementSubmissionFlags) (map[string]any, bool, error) {
	hasAny := strings.TrimSpace(flags.requirementType) != "" ||
		strings.TrimSpace(flags.value) != "" ||
		strings.TrimSpace(flags.personID) != ""
	if !hasAny {
		return nil, false, nil
	}
	if strings.TrimSpace(flags.requirementType) == "" {
		return nil, false, errors.New("--requirement-type is required when building a requirement submission from flags")
	}
	if strings.TrimSpace(flags.value) == "" {
		return nil, false, errors.New("--value is required when building a requirement submission from flags")
	}

	entry := map[string]any{
		"requirement_type": flags.requirementType,
		"value":            flags.value,
	}
	if strings.TrimSpace(flags.personID) != "" {
		entry["person_id"] = flags.personID
	}
	return entry, true, nil
}

func applyTransferPlatformPersonAddress(cmd *cobra.Command, body map[string]any, street, street2, city, region, postalCode, country string) error {
	shouldSet := anyFlagChanged(cmd, "street", "street2", "city", "region", "postal-code", "country") ||
		((street != "" || street2 != "" || city != "" || region != "" || postalCode != "" || country != "") &&
			!bodyHasValue(body, "address"))
	if !shouldSet {
		return nil
	}

	address := map[string]any{}
	if street != "" {
		address["street"] = street
	}
	if street2 != "" {
		address["street2"] = street2
	}
	if city != "" {
		address["city"] = city
	}
	if region != "" {
		address["region"] = region
	}
	if postalCode != "" {
		address["postal_code"] = postalCode
	}
	if country != "" {
		address["country"] = country
	}
	if len(address) == 0 {
		return nil
	}
	return setBodyValue(body, address, "address")
}

func newTransferPlatformCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "platform",
		Short: "Transfer for Platforms write commands",
		Long:  "Commands for onboarding transfer platform originators, people, requirements, and documents.",
	}

	cmd.AddCommand(newTransferPlatformOriginatorCreateCmd())
	cmd.AddCommand(newTransferPlatformPersonCreateCmd())
	cmd.AddCommand(newTransferPlatformRequirementSubmitCmd())
	cmd.AddCommand(newTransferPlatformDocumentSubmitCmd())

	return cmd
}

func newTransferPlatformOriginatorCreateCmd() *cobra.Command {
	var originatorClientID, originatorIPAddress, agreementAcceptedAt, originatorReviewedAt, webhook string
	var agreementAccepted bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "originator-create",
		Short: "Call /transfer/platform/originator/create",
		Long:  "Capability: write. Starts Transfer for Platforms onboarding for an originator.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"originator_client_id": "<originator-client-id>",
				"tos_acceptance_metadata": map[string]any{
					"agreement_accepted":    true,
					"originator_ip_address": "203.0.113.10",
					"agreement_accepted_at": "2026-03-08T00:00:00Z",
				},
				"originator_reviewed_at": "2026-03-08T00:00:00Z",
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "agreement-accepted", agreementAccepted, "tos_acceptance_metadata", "agreement_accepted"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "originator-ip-address", originatorIPAddress, "tos_acceptance_metadata", "originator_ip_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "agreement-accepted-at", agreementAcceptedAt, "tos_acceptance_metadata", "agreement_accepted_at"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "originator-reviewed-at", originatorReviewedAt, "originator_reviewed_at"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--originator-client-id":   {"originator_client_id"},
				"--agreement-accepted":     {"tos_acceptance_metadata", "agreement_accepted"},
				"--originator-ip-address":  {"tos_acceptance_metadata", "originator_ip_address"},
				"--agreement-accepted-at":  {"tos_acceptance_metadata", "agreement_accepted_at"},
				"--originator-reviewed-at": {"originator_reviewed_at"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/platform/originator/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Client ID for the originator being onboarded")
	cmd.Flags().BoolVar(&agreementAccepted, "agreement-accepted", false, "Whether the originator accepted the Transfer terms of service")
	cmd.Flags().StringVar(&originatorIPAddress, "originator-ip-address", "", "IP address used when the originator accepted the terms")
	cmd.Flags().StringVar(&agreementAcceptedAt, "agreement-accepted-at", "", "RFC3339 timestamp when the originator accepted the terms")
	cmd.Flags().StringVar(&originatorReviewedAt, "originator-reviewed-at", "", "RFC3339 timestamp when you last collected onboarding data")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive PLATFORM_ONBOARDING_UPDATE events")
	return cmd
}

func newTransferPlatformPersonCreateCmd() *cobra.Command {
	var originatorClientID string
	var givenName, familyName, emailAddress, phoneNumber string
	var street, street2, city, region, postalCode, country string
	var idNumberType, idNumberValue, dateOfBirth, relationshipToOriginator, title string
	var ownershipPercentage int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "person-create",
		Short: "Call /transfer/platform/person/create",
		Long:  "Capability: write. Creates a person associated with a transfer originator.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"originator_client_id": "<originator-client-id>",
				"name": map[string]any{
					"given_name":  "Jane",
					"family_name": "Doe",
				},
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "given-name", givenName, "name", "given_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "family-name", familyName, "name", "family_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "email-address", emailAddress, "email_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "phone-number", phoneNumber, "phone_number"); err != nil {
				return err
			}
			if err := applyTransferPlatformPersonAddress(cmd, body, street, street2, city, region, postalCode, country); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "id-number-type", idNumberType, "id_number", "type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "id-number-value", idNumberValue, "id_number", "value"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "date-of-birth", dateOfBirth, "date_of_birth"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "relationship-to-originator", relationshipToOriginator, "relationship_to_originator"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "title", title, "title"); err != nil {
				return err
			}
			if err := applyOptionalIntFlag(cmd, body, "ownership-percentage", ownershipPercentage, "ownership_percentage"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--originator-client-id": {"originator_client_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/platform/person/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID that this person belongs to")
	cmd.Flags().StringVar(&givenName, "given-name", "", "Person given name")
	cmd.Flags().StringVar(&familyName, "family-name", "", "Person family name")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "Person email address")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Person phone number in E.164 format")
	cmd.Flags().StringVar(&street, "street", "", "Primary street address")
	cmd.Flags().StringVar(&street2, "street2", "", "Secondary street address")
	cmd.Flags().StringVar(&city, "city", "", "City")
	cmd.Flags().StringVar(&region, "region", "", "Region or state")
	cmd.Flags().StringVar(&postalCode, "postal-code", "", "Postal code")
	cmd.Flags().StringVar(&country, "country", "", "Two-letter ISO country code")
	cmd.Flags().StringVar(&idNumberType, "id-number-type", "", "ID number type, e.g. us_ssn")
	cmd.Flags().StringVar(&idNumberValue, "id-number-value", "", "ID number value with formatting stripped")
	cmd.Flags().StringVar(&dateOfBirth, "date-of-birth", "", "Date of birth in YYYY-MM-DD")
	cmd.Flags().StringVar(&relationshipToOriginator, "relationship-to-originator", "", "Relationship to the originator, e.g. BENEFICIAL_OWNER")
	cmd.Flags().IntVar(&ownershipPercentage, "ownership-percentage", 0, "Ownership percentage for beneficial owners")
	cmd.Flags().StringVar(&title, "title", "", "Business title for control persons")
	return cmd
}

func newTransferPlatformRequirementSubmitCmd() *cobra.Command {
	var originatorClientID string
	var submissions []string
	var single transferRequirementSubmissionFlags
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "requirement-submit",
		Short: "Call /transfer/platform/requirement/submit",
		Long:  "Capability: write. Submits onboarding requirements for a transfer originator.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"originator_client_id": "<originator-client-id>",
				"requirement_submissions": []map[string]any{
					{
						"requirement_type": "BUSINESS_NAME",
						"value":            "Example LLC",
					},
				},
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}

			entries := make([]map[string]any, 0, len(submissions)+1)
			for _, spec := range submissions {
				entry, err := parseTransferRequirementSubmissionSpec(spec)
				if err != nil {
					return err
				}
				entries = append(entries, entry)
			}
			if entry, ok, err := buildTransferRequirementSubmissionFromFlags(single); err != nil {
				return err
			} else if ok {
				entries = append(entries, entry)
			}
			if len(entries) > 0 {
				if err := setBodyValue(body, entries, "requirement_submissions"); err != nil {
					return err
				}
			}
			if err := requireBodyFields(body, map[string][]string{
				"--originator-client-id":                         {"originator_client_id"},
				"--submission or --body.requirement_submissions": {"requirement_submissions"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/platform/requirement/submit", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID that the requirement submissions apply to")
	cmd.Flags().StringSliceVar(&submissions, "submission", nil, "Repeatable JSON object or @path describing one requirement submission")
	cmd.Flags().StringVar(&single.requirementType, "requirement-type", "", "Requirement type for a single convenience submission")
	cmd.Flags().StringVar(&single.value, "value", "", "Requirement value for a single convenience submission")
	cmd.Flags().StringVar(&single.personID, "person-id", "", "Optional person_id for a single convenience submission")
	return cmd
}

func newTransferPlatformDocumentSubmitCmd() *cobra.Command {
	var originatorClientID, filePath, requirementType, personID string
	var info *commandInfoFlags

	cmd := &cobra.Command{
		Use:   "document-submit",
		Short: "Call /transfer/platform/document/submit",
		Long:  "Capability: write. Uploads a document needed for transfer originator onboarding.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"originator_client_id": "<originator-client-id>",
				"document_submission":  "/path/to/document.pdf",
				"requirement_type":     "BUSINESS_ADDRESS_VALIDATION",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferPlatformDocPath, template); handled || err != nil {
				return err
			}
			if strings.TrimSpace(originatorClientID) == "" {
				return errors.New("--originator-client-id is required")
			}
			if strings.TrimSpace(filePath) == "" {
				return errors.New("--file is required")
			}
			if strings.TrimSpace(requirementType) == "" {
				return errors.New("--requirement-type is required")
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}

			fields := map[string]string{
				"originator_client_id": originatorClientID,
				"requirement_type":     requirementType,
				"person_id":            personID,
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.CallMultipart(ctx, "/transfer/platform/document/submit", fields, "document_submission", filePath)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID that owns the submitted document")
	cmd.Flags().StringVar(&filePath, "file", "", "Path to the file to upload")
	cmd.Flags().StringVar(&requirementType, "requirement-type", "", "Requirement type fulfilled by the uploaded document")
	cmd.Flags().StringVar(&personID, "person-id", "", "Optional person_id associated with the document")
	return cmd
}
