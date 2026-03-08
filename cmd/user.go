package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const usersDocPath = "docs/plaid/api/users/index.md"

type userIdentityFlags struct {
	givenName    string
	familyName   string
	dateOfBirth  string
	email        string
	phoneNumber  string
	street1      string
	street2      string
	city         string
	region       string
	country      string
	postalCode   string
	idNumber     string
	idNumberType string
}

func newUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "User API commands",
		Long:  "Admin-style commands for Plaid user_id lifecycle. Typed flags target the new identity schema; use --body for legacy user_token-only payloads.",
	}

	cmd.AddCommand(newUserCreateCmd())
	cmd.AddCommand(newUserGetCmd())
	cmd.AddCommand(newUserUpdateCmd())
	cmd.AddCommand(newUserRemoveCmd())
	cmd.AddCommand(newUserItemsGetCmd())

	return cmd
}

func newUserCreateCmd() *cobra.Command {
	var clientUserID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var identity *userIdentityFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /user/create",
		Long:  "Capability: admin. Creates a Plaid user_id. Typed identity flags populate one primary email, phone number, address, and id_number in the new identity schema.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"client_user_id": "<client-user-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, usersDocPath, template); handled || err != nil {
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
			if err := applyUserIdentityFlags(cmd, body, identity); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--client-user-id": {"client_user_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/user/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	identity = bindUserIdentityFlags(cmd)
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable, non-PII user identifier from your system")
	return cmd
}

func newUserGetCmd() *cobra.Command {
	var userID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /user/get",
		Long:  "Capability: read. Retrieves a Plaid user created on the new user_id flow.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id": "<user-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, usersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--user-id": {"user_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/user/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id, e.g. usr_...")
	return cmd
}

func newUserUpdateCmd() *cobra.Command {
	var userID, userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var identity *userIdentityFlags

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Call /user/update",
		Long:  "Capability: admin. Updates user identity fields. Typed flags target the new identity schema; use --body for legacy consumer_report_user_identity payloads.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id": "<user-id>",
				"identity": map[string]any{
					"emails": []map[string]any{
						{
							"data":    "user@example.com",
							"primary": true,
						},
					},
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, usersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-token", userToken, "user_token"); err != nil {
				return err
			}
			if err := requireExactlyOneBodyField(body, map[string][]string{
				"--user-id":    {"user_id"},
				"--user-token": {"user_token"},
			}); err != nil {
				return err
			}
			if err := applyUserIdentityFlags(cmd, body, identity); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/user/update", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	identity = bindUserIdentityFlags(cmd)
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id for new user API flows")
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token for older integrations")
	return cmd
}

func newUserRemoveCmd() *cobra.Command {
	var userID, userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Call /user/remove",
		Long:  "Capability: admin. Removes a Plaid user and any associated Items.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id": "<user-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, usersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-token", userToken, "user_token"); err != nil {
				return err
			}
			if err := requireExactlyOneBodyField(body, map[string][]string{
				"--user-id":    {"user_id"},
				"--user-token": {"user_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/user/remove", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id for new user API flows")
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token for older integrations")
	return cmd
}

func newUserItemsGetCmd() *cobra.Command {
	var userID, userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "items-get",
		Short: "Call /user/items/get",
		Long:  "Capability: read. Lists Items associated with a Plaid user_id or legacy user_token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id": "<user-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, usersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-token", userToken, "user_token"); err != nil {
				return err
			}
			if err := requireExactlyOneBodyField(body, map[string][]string{
				"--user-id":    {"user_id"},
				"--user-token": {"user_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/user/items/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userID, "user-id", "", "Plaid user_id for new user API flows")
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token for older integrations")
	return cmd
}

func bindUserIdentityFlags(cmd *cobra.Command) *userIdentityFlags {
	flags := &userIdentityFlags{}
	if cmd == nil {
		return flags
	}

	cmd.Flags().StringVar(&flags.givenName, "given-name", "", "Primary identity given name")
	cmd.Flags().StringVar(&flags.familyName, "family-name", "", "Primary identity family name")
	cmd.Flags().StringVar(&flags.dateOfBirth, "date-of-birth", "", "Date of birth in YYYY-MM-DD format")
	cmd.Flags().StringVar(&flags.email, "email", "", "Primary email address")
	cmd.Flags().StringVar(&flags.phoneNumber, "phone-number", "", "Primary phone number in E.164 format")
	cmd.Flags().StringVar(&flags.street1, "street-1", "", "Primary address line 1")
	cmd.Flags().StringVar(&flags.street2, "street-2", "", "Primary address line 2")
	cmd.Flags().StringVar(&flags.city, "city", "", "Primary address city")
	cmd.Flags().StringVar(&flags.region, "region", "", "Primary address region or state")
	cmd.Flags().StringVar(&flags.country, "country", "", "Primary address country code")
	cmd.Flags().StringVar(&flags.postalCode, "postal-code", "", "Primary address postal code")
	cmd.Flags().StringVar(&flags.idNumber, "id-number", "", "Primary identity document value, e.g. SSN last four digits")
	cmd.Flags().StringVar(&flags.idNumberType, "id-number-type", "", "Primary identity document type, e.g. us_ssn_last_4 or us_ssn")
	return flags
}

func applyUserIdentityFlags(cmd *cobra.Command, body map[string]any, flags *userIdentityFlags) error {
	if flags == nil {
		return nil
	}

	if err := applyStringFlag(cmd, body, "given-name", flags.givenName, "identity", "name", "given_name"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "family-name", flags.familyName, "identity", "name", "family_name"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "date-of-birth", flags.dateOfBirth, "identity", "date_of_birth"); err != nil {
		return err
	}
	if err := applyPrimaryIdentityValue(cmd, body, "email", flags.email, "emails"); err != nil {
		return err
	}
	if err := applyPrimaryIdentityValue(cmd, body, "phone-number", flags.phoneNumber, "phone_numbers"); err != nil {
		return err
	}
	if err := applyPrimaryAddress(cmd, body, flags); err != nil {
		return err
	}
	if err := applyPrimaryIDNumber(cmd, body, flags); err != nil {
		return err
	}
	return nil
}

func applyPrimaryIdentityValue(cmd *cobra.Command, body map[string]any, flagName, value, field string) error {
	shouldSet := cmd.Flags().Changed(flagName) || (value != "" && !bodyHasValue(body, "identity", field))
	if !shouldSet {
		return nil
	}
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("--%s requires a non-empty value", flagName)
	}

	return setBodyValue(body, []map[string]any{
		{
			"data":    value,
			"primary": true,
		},
	}, "identity", field)
}

func applyPrimaryAddress(cmd *cobra.Command, body map[string]any, flags *userIdentityFlags) error {
	shouldSet := anyFlagChanged(cmd, "street-1", "street-2", "city", "region", "country", "postal-code") ||
		((flags.street1 != "" || flags.street2 != "" || flags.city != "" || flags.region != "" || flags.country != "" || flags.postalCode != "") &&
			!bodyHasValue(body, "identity", "addresses"))
	if !shouldSet {
		return nil
	}

	address := map[string]any{
		"primary": true,
	}
	if flags.street1 != "" {
		address["street_1"] = flags.street1
	}
	if flags.street2 != "" {
		address["street_2"] = flags.street2
	}
	if flags.city != "" {
		address["city"] = flags.city
	}
	if flags.region != "" {
		address["region"] = flags.region
	}
	if flags.country != "" {
		address["country"] = flags.country
	}
	if flags.postalCode != "" {
		address["postal_code"] = flags.postalCode
	}
	if len(address) == 1 {
		return errors.New("address flags were provided without any address values")
	}

	return setBodyValue(body, []map[string]any{address}, "identity", "addresses")
}

func applyPrimaryIDNumber(cmd *cobra.Command, body map[string]any, flags *userIdentityFlags) error {
	shouldSet := anyFlagChanged(cmd, "id-number", "id-number-type") ||
		((flags.idNumber != "" || flags.idNumberType != "") && !bodyHasValue(body, "identity", "id_numbers"))
	if !shouldSet {
		return nil
	}
	if strings.TrimSpace(flags.idNumber) == "" || strings.TrimSpace(flags.idNumberType) == "" {
		return errors.New("--id-number and --id-number-type must be provided together")
	}

	return setBodyValue(body, []map[string]any{
		{
			"value": flags.idNumber,
			"type":  flags.idNumberType,
		},
	}, "identity", "id_numbers")
}

func anyFlagChanged(cmd *cobra.Command, names ...string) bool {
	for _, name := range names {
		if cmd.Flags().Changed(name) {
			return true
		}
	}
	return false
}
