package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const identityVerificationDocPath = "docs/plaid/api/products/identity-verification/index.md"

type identityVerificationUserFlags struct {
	emailAddress string
	phoneNumber  string
	dateOfBirth  string
	givenName    string
	familyName   string
	street       string
	street2      string
	city         string
	region       string
	postalCode   string
	country      string
	idNumber     string
	idNumberType string
}

type identityVerificationRetryStepFlags struct {
	verifySMS               bool
	kycCheck                bool
	documentaryVerification bool
	selfieCheck             bool
}

func newIdentityVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identity-verification",
		Short: "Identity Verification product commands",
		Long:  "Identity Verification commands. These create, inspect, and retry Plaid identity verification attempts.",
	}

	cmd.AddCommand(newIdentityVerificationCreateCmd())
	cmd.AddCommand(newIdentityVerificationGetCmd())
	cmd.AddCommand(newIdentityVerificationListCmd())
	cmd.AddCommand(newIdentityVerificationRetryCmd())

	return cmd
}

func newIdentityVerificationCreateCmd() *cobra.Command {
	var clientUserID, userID, templateID string
	var isShareable, gaveConsent, isIdempotent bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userFlags *identityVerificationUserFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /identity_verification/create",
		Long:  "Capability: write. Creates a new Plaid identity verification attempt.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"client_user_id": "<client-user-id>",
				"template_id":    "<template-id>",
				"is_shareable":   true,
				"gave_consent":   true,
			}
			if handled, err := maybeWriteInfo(cmd, info, identityVerificationDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "client-user-id", clientUserID, "client_user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "template-id", templateID, "template_id"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "is-shareable", isShareable, "is_shareable"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "gave-consent", gaveConsent, "gave_consent"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "is-idempotent", isIdempotent, "is_idempotent"); err != nil {
				return err
			}
			if err := requireAtLeastOneBodyField(body, map[string][]string{
				"--client-user-id": {"client_user_id"},
				"--user-id":        {"user_id"},
			}); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--template-id":  {"template_id"},
				"--is-shareable": {"is_shareable"},
				"--gave-consent": {"gave_consent"},
			}); err != nil {
				return err
			}
			if bodyHasValue(body, "user_id") && (identityVerificationUserFlagsProvided(cmd) || bodyHasValue(body, "user")) {
				return errors.New("user fields cannot be provided when user_id is present; use /user/update for stored user data")
			}
			if err := applyIdentityVerificationUserFlags(cmd, body, userFlags); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/identity_verification/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	userFlags = bindIdentityVerificationUserFlags(cmd)
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable client-side user identifier")
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id from /user/create")
	cmd.Flags().StringVar(&templateID, "template-id", "", "Plaid identity verification template ID")
	cmd.Flags().BoolVar(&isShareable, "is-shareable", false, "Whether Plaid should expose a shareable verification URL")
	cmd.Flags().BoolVar(&gaveConsent, "gave-consent", false, "Whether the user has already consented to share data with Plaid")
	cmd.Flags().BoolVar(&isIdempotent, "is-idempotent", false, "Reuse an existing active verification when possible")
	return cmd
}

func newIdentityVerificationGetCmd() *cobra.Command {
	var identityVerificationID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /identity_verification/get",
		Long:  "Capability: read. Retrieves a previously created identity verification attempt.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"identity_verification_id": "<identity-verification-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, identityVerificationDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "identity-verification-id", identityVerificationID, "identity_verification_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--identity-verification-id": {"identity_verification_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/identity_verification/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&identityVerificationID, "identity-verification-id", "", "Plaid identity verification attempt ID")
	return cmd
}

func newIdentityVerificationListCmd() *cobra.Command {
	var templateID, clientUserID, userID, cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /identity_verification/list",
		Long:  "Capability: read. Lists identity verification attempts for a template, optionally scoped to a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"template_id": "<template-id>",
			}
			if clientUserID != "" {
				template["client_user_id"] = clientUserID
			}
			if userID != "" {
				template["user_id"] = userID
			}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, identityVerificationDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "client-user-id", clientUserID, "client_user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--template-id": {"template_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/identity_verification/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&templateID, "template-id", "", "Plaid identity verification template ID")
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable client-side user identifier")
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id from /user/create")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Cursor returned by a previous identity-verification list call")
	return cmd
}

func newIdentityVerificationRetryCmd() *cobra.Command {
	var clientUserID, templateID, strategy string
	var isShareable bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userFlags *identityVerificationUserFlags
	var stepFlags *identityVerificationRetryStepFlags

	cmd := &cobra.Command{
		Use:   "retry",
		Short: "Call /identity_verification/retry",
		Long:  "Capability: write. Starts a new identity verification attempt for an existing user and template.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"client_user_id": "<client-user-id>",
				"template_id":    "<template-id>",
				"strategy":       "reset",
			}
			if handled, err := maybeWriteInfo(cmd, info, identityVerificationDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "client-user-id", clientUserID, "client_user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "template-id", templateID, "template_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "strategy", strategy, "strategy"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "is-shareable", isShareable, "is_shareable"); err != nil {
				return err
			}
			if err := applyIdentityVerificationRetrySteps(cmd, body, stepFlags); err != nil {
				return err
			}
			if err := applyIdentityVerificationUserFlags(cmd, body, userFlags); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--client-user-id": {"client_user_id"},
				"--template-id":    {"template_id"},
				"--strategy":       {"strategy"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/identity_verification/retry", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	userFlags = bindIdentityVerificationUserFlags(cmd)
	stepFlags = bindIdentityVerificationRetryStepFlags(cmd)
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable client-side user identifier")
	cmd.Flags().StringVar(&templateID, "template-id", "", "Plaid identity verification template ID")
	cmd.Flags().StringVar(&strategy, "strategy", "", "Retry strategy: reset, incomplete, infer, or custom")
	cmd.Flags().BoolVar(&isShareable, "is-shareable", false, "Override whether Plaid should expose a shareable verification URL")
	return cmd
}

func bindIdentityVerificationUserFlags(cmd *cobra.Command) *identityVerificationUserFlags {
	flags := &identityVerificationUserFlags{}
	if cmd == nil {
		return flags
	}

	cmd.Flags().StringVar(&flags.emailAddress, "email-address", "", "User email address")
	cmd.Flags().StringVar(&flags.phoneNumber, "phone-number", "", "User phone number in E.164 format")
	cmd.Flags().StringVar(&flags.dateOfBirth, "date-of-birth", "", "User date of birth in YYYY-MM-DD format")
	cmd.Flags().StringVar(&flags.givenName, "given-name", "", "User given name")
	cmd.Flags().StringVar(&flags.familyName, "family-name", "", "User family name")
	cmd.Flags().StringVar(&flags.street, "street", "", "User address line 1")
	cmd.Flags().StringVar(&flags.street2, "street-2", "", "User address line 2")
	cmd.Flags().StringVar(&flags.city, "city", "", "User address city")
	cmd.Flags().StringVar(&flags.region, "region", "", "User address region or state")
	cmd.Flags().StringVar(&flags.postalCode, "postal-code", "", "User address postal code")
	cmd.Flags().StringVar(&flags.country, "country", "", "User address country code")
	cmd.Flags().StringVar(&flags.idNumber, "id-number", "", "User identity document value")
	cmd.Flags().StringVar(&flags.idNumberType, "id-number-type", "", "User identity document type")
	return flags
}

func identityVerificationUserFlagsProvided(cmd *cobra.Command) bool {
	return anyFlagChanged(
		cmd,
		"email-address",
		"phone-number",
		"date-of-birth",
		"given-name",
		"family-name",
		"street",
		"street-2",
		"city",
		"region",
		"postal-code",
		"country",
		"id-number",
		"id-number-type",
	)
}

func applyIdentityVerificationUserFlags(cmd *cobra.Command, body map[string]any, flags *identityVerificationUserFlags) error {
	if flags == nil {
		return nil
	}
	if err := applyStringFlag(cmd, body, "email-address", flags.emailAddress, "user", "email_address"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "phone-number", flags.phoneNumber, "user", "phone_number"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "date-of-birth", flags.dateOfBirth, "user", "date_of_birth"); err != nil {
		return err
	}
	if err := applyIdentityVerificationUserNameFlags(cmd, body, flags); err != nil {
		return err
	}
	if err := applyIdentityVerificationUserAddressFlags(cmd, body, flags); err != nil {
		return err
	}
	if err := applyIdentityVerificationIDNumberFlags(cmd, body, flags); err != nil {
		return err
	}
	return nil
}

func applyIdentityVerificationUserNameFlags(cmd *cobra.Command, body map[string]any, flags *identityVerificationUserFlags) error {
	shouldSet := anyFlagChanged(cmd, "given-name", "family-name") ||
		((flags.givenName != "" || flags.familyName != "") && !bodyHasValue(body, "user", "name"))
	if !shouldSet {
		return nil
	}
	if flags.givenName == "" || flags.familyName == "" {
		return errors.New("--given-name and --family-name must be provided together")
	}
	return setBodyValue(body, map[string]any{
		"given_name":  flags.givenName,
		"family_name": flags.familyName,
	}, "user", "name")
}

func applyIdentityVerificationUserAddressFlags(cmd *cobra.Command, body map[string]any, flags *identityVerificationUserFlags) error {
	shouldSet := anyFlagChanged(cmd, "street", "street-2", "city", "region", "postal-code", "country") ||
		((flags.street != "" || flags.street2 != "" || flags.city != "" || flags.region != "" || flags.postalCode != "" || flags.country != "") &&
			!bodyHasValue(body, "user", "address"))
	if !shouldSet {
		return nil
	}

	address := map[string]any{}
	if flags.street != "" {
		address["street"] = flags.street
	}
	if flags.street2 != "" {
		address["street2"] = flags.street2
	}
	if flags.city != "" {
		address["city"] = flags.city
	}
	if flags.region != "" {
		address["region"] = flags.region
	}
	if flags.postalCode != "" {
		address["postal_code"] = flags.postalCode
	}
	if flags.country != "" {
		address["country"] = flags.country
	}
	if len(address) == 0 {
		return nil
	}

	if len(address) == 1 && flags.country != "" {
		return setBodyValue(body, address, "user", "address")
	}
	if flags.street == "" || flags.city == "" || flags.country == "" {
		return errors.New("user address must be country-only or include --street, --city, and --country")
	}

	return setBodyValue(body, address, "user", "address")
}

func applyIdentityVerificationIDNumberFlags(cmd *cobra.Command, body map[string]any, flags *identityVerificationUserFlags) error {
	shouldSet := anyFlagChanged(cmd, "id-number", "id-number-type") ||
		((flags.idNumber != "" || flags.idNumberType != "") && !bodyHasValue(body, "user", "id_number"))
	if !shouldSet {
		return nil
	}
	if flags.idNumber == "" || flags.idNumberType == "" {
		return errors.New("--id-number and --id-number-type must be provided together")
	}
	return setBodyValue(body, map[string]any{
		"value": flags.idNumber,
		"type":  flags.idNumberType,
	}, "user", "id_number")
}

func bindIdentityVerificationRetryStepFlags(cmd *cobra.Command) *identityVerificationRetryStepFlags {
	flags := &identityVerificationRetryStepFlags{}
	if cmd == nil {
		return flags
	}
	cmd.Flags().BoolVar(&flags.verifySMS, "verify-sms", false, "Custom retry: require the verify_sms step")
	cmd.Flags().BoolVar(&flags.kycCheck, "kyc-check", false, "Custom retry: require the kyc_check step")
	cmd.Flags().BoolVar(&flags.documentaryVerification, "documentary-verification", false, "Custom retry: require the documentary_verification step")
	cmd.Flags().BoolVar(&flags.selfieCheck, "selfie-check", false, "Custom retry: require the selfie_check step")
	return flags
}

func applyIdentityVerificationRetrySteps(cmd *cobra.Command, body map[string]any, flags *identityVerificationRetryStepFlags) error {
	if flags == nil {
		return nil
	}

	stepsChanged := anyFlagChanged(cmd, "verify-sms", "kyc-check", "documentary-verification", "selfie-check")
	strategyValue, _ := bodyValue(body, "strategy")
	strategy, _ := strategyValue.(string)
	if !stepsChanged && !bodyHasValue(body, "steps") {
		if strategy == "custom" {
			return errors.New("custom retry strategy requires steps; provide step flags or --body")
		}
		return nil
	}

	if strategy != "custom" {
		return errors.New("retry step flags are only valid with --strategy custom")
	}

	if stepsChanged {
		if !anyFlagChanged(cmd, "verify-sms") || !anyFlagChanged(cmd, "kyc-check") || !anyFlagChanged(cmd, "documentary-verification") || !anyFlagChanged(cmd, "selfie-check") {
			return errors.New("custom retry strategy requires all of --verify-sms, --kyc-check, --documentary-verification, and --selfie-check")
		}
		return setBodyValue(body, map[string]any{
			"verify_sms":               flags.verifySMS,
			"kyc_check":                flags.kycCheck,
			"documentary_verification": flags.documentaryVerification,
			"selfie_check":             flags.selfieCheck,
		}, "steps")
	}

	if !bodyHasValue(body, "steps", "verify_sms") ||
		!bodyHasValue(body, "steps", "kyc_check") ||
		!bodyHasValue(body, "steps", "documentary_verification") ||
		!bodyHasValue(body, "steps", "selfie_check") {
		return errors.New("custom retry strategy requires steps.verify_sms, steps.kyc_check, steps.documentary_verification, and steps.selfie_check")
	}

	return nil
}
