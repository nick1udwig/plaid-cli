package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

const beaconDocPath = "docs/plaid/api/products/beacon/index.md"

type beaconUserIdentityFlags struct {
	dateOfBirth string
	givenName   string
	familyName  string
	street      string
	street2     string
	city        string
	region      string
	postalCode  string
	country     string
	email       string
	phone       string
	idNumber    string
	idType      string
	ipAddress   string
}

func newBeaconCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beacon",
		Short: "Beacon product commands",
		Long:  "Beacon commands for user screening, account insights, fraud reports, syndications, and duplicate review workflows.",
	}

	cmd.AddCommand(newBeaconUserCmd())
	cmd.AddCommand(newBeaconReportCmd())
	cmd.AddCommand(newBeaconReportSyndicationCmd())
	cmd.AddCommand(newBeaconDuplicateCmd())

	return cmd
}

func newBeaconUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Beacon user commands",
		Long:  "Create, fetch, update, and inspect Beacon users.",
	}

	cmd.AddCommand(newBeaconUserCreateCmd())
	cmd.AddCommand(newBeaconUserGetCmd())
	cmd.AddCommand(newBeaconUserUpdateCmd())
	cmd.AddCommand(newBeaconUserAccountInsightsGetCmd())
	cmd.AddCommand(newBeaconUserHistoryListCmd())

	return cmd
}

func newBeaconReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Beacon report commands",
		Long:  "Create and retrieve Beacon fraud reports for a Beacon user.",
	}

	cmd.AddCommand(newBeaconReportCreateCmd())
	cmd.AddCommand(newBeaconReportGetCmd())
	cmd.AddCommand(newBeaconReportListCmd())

	return cmd
}

func newBeaconReportSyndicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-syndication",
		Short: "Beacon report syndication commands",
		Long:  "Retrieve Beacon report syndications attached to a Beacon user.",
	}

	cmd.AddCommand(newBeaconReportSyndicationGetCmd())
	cmd.AddCommand(newBeaconReportSyndicationListCmd())

	return cmd
}

func newBeaconDuplicateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "duplicate",
		Short: "Beacon duplicate commands",
		Long:  "Retrieve duplicate-review records for similar Beacon users.",
	}

	cmd.AddCommand(newBeaconDuplicateGetCmd())

	return cmd
}

func bindBeaconUserIdentityFlags(cmd *cobra.Command) *beaconUserIdentityFlags {
	flags := &beaconUserIdentityFlags{}
	cmd.Flags().StringVar(&flags.dateOfBirth, "date-of-birth", "", "Date of birth in YYYY-MM-DD")
	cmd.Flags().StringVar(&flags.givenName, "given-name", "", "User given name")
	cmd.Flags().StringVar(&flags.familyName, "family-name", "", "User family name")
	cmd.Flags().StringVar(&flags.street, "street", "", "User street address")
	cmd.Flags().StringVar(&flags.street2, "street2", "", "User address line 2")
	cmd.Flags().StringVar(&flags.city, "city", "", "User city")
	cmd.Flags().StringVar(&flags.region, "region", "", "User region or state code")
	cmd.Flags().StringVar(&flags.postalCode, "postal-code", "", "User postal code")
	cmd.Flags().StringVar(&flags.country, "country", "", "User country as ISO-3166-1 alpha-2")
	cmd.Flags().StringVar(&flags.email, "email", "", "User email address")
	cmd.Flags().StringVar(&flags.phone, "phone", "", "User phone number in E.164 format")
	cmd.Flags().StringVar(&flags.idNumber, "id-number", "", "User identity document number")
	cmd.Flags().StringVar(&flags.idType, "id-type", "", "User identity document type, e.g. us_ssn")
	cmd.Flags().StringVar(&flags.ipAddress, "ip-address", "", "User IPv4 or IPv6 address")
	return flags
}

func beaconUserIdentityFlagsChanged(cmd *cobra.Command) bool {
	return anyFlagChanged(
		cmd,
		"date-of-birth",
		"given-name",
		"family-name",
		"street",
		"street2",
		"city",
		"region",
		"postal-code",
		"country",
		"email",
		"phone",
		"id-number",
		"id-type",
		"ip-address",
	)
}

func applyBeaconUserIdentityFlags(cmd *cobra.Command, body map[string]any, flags *beaconUserIdentityFlags) error {
	if flags == nil {
		return nil
	}

	if err := applyStringFlag(cmd, body, "date-of-birth", flags.dateOfBirth, "user", "date_of_birth"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "email", flags.email, "user", "email_address"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "phone", flags.phone, "user", "phone_number"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "ip-address", flags.ipAddress, "user", "ip_address"); err != nil {
		return err
	}

	nameChanged := anyFlagChanged(cmd, "given-name", "family-name")
	if nameChanged || ((flags.givenName != "" || flags.familyName != "") && !bodyHasValue(body, "user", "name")) {
		name := map[string]any{}
		if flags.givenName != "" {
			name["given_name"] = flags.givenName
		}
		if flags.familyName != "" {
			name["family_name"] = flags.familyName
		}
		if len(name) > 0 {
			if err := setBodyValue(body, name, "user", "name"); err != nil {
				return err
			}
		}
	}

	addressChanged := anyFlagChanged(cmd, "street", "street2", "city", "region", "postal-code", "country")
	if addressChanged || ((flags.street != "" || flags.street2 != "" || flags.city != "" || flags.region != "" || flags.postalCode != "" || flags.country != "") && !bodyHasValue(body, "user", "address")) {
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
		if len(address) > 0 {
			if err := setBodyValue(body, address, "user", "address"); err != nil {
				return err
			}
		}
	}

	idChanged := anyFlagChanged(cmd, "id-number", "id-type")
	if idChanged || ((flags.idNumber != "" || flags.idType != "") && !bodyHasValue(body, "user", "id_number")) {
		idNumber := map[string]any{}
		if flags.idNumber != "" {
			idNumber["value"] = flags.idNumber
		}
		if flags.idType != "" {
			idNumber["type"] = flags.idType
		}
		if len(idNumber) > 0 {
			if err := setBodyValue(body, idNumber, "user", "id_number"); err != nil {
				return err
			}
		}
	}

	return validateBeaconUserIdentity(body)
}

func validateBeaconUserIdentity(body map[string]any) error {
	if bodyHasValue(body, "user", "name") {
		for _, field := range []string{"given_name", "family_name"} {
			if bodyHasValue(body, "user", "name", field) {
				continue
			}
			return fmt.Errorf("user.name.%s is required", field)
		}
	}

	if bodyHasValue(body, "user", "address") {
		for _, field := range []string{"street", "city", "country"} {
			if bodyHasValue(body, "user", "address", field) {
				continue
			}
			return fmt.Errorf("user.address.%s is required", field)
		}
	}

	if bodyHasValue(body, "user", "id_number") {
		for _, field := range []string{"value", "type"} {
			if bodyHasValue(body, "user", "id_number", field) {
				continue
			}
			return fmt.Errorf("user.id_number.%s is required", field)
		}
	}

	return nil
}

func validateBeaconUserCreateBody(body map[string]any) error {
	if err := requireBodyFields(body, map[string][]string{
		"--program-id":     {"program_id"},
		"--client-user-id": {"client_user_id"},
	}); err != nil {
		return err
	}
	if !bodyHasValue(body, "user") {
		return errors.New("user is required")
	}
	if !bodyHasValue(body, "user", "name") {
		return errors.New("user.name is required")
	}
	if err := validateBeaconUserIdentity(body); err != nil {
		return err
	}
	if bodyHasValue(body, "user", "date_of_birth") || bodyHasValue(body, "user", "depository_accounts") {
		return nil
	}
	return errors.New("provide at least one of user.date_of_birth or user.depository_accounts")
}

func validateBeaconUserUpdateBody(body map[string]any) error {
	if err := requireBodyFields(body, map[string][]string{
		"--beacon-user-id": {"beacon_user_id"},
	}); err != nil {
		return err
	}
	if err := validateBeaconUserIdentity(body); err != nil {
		return err
	}
	if bodyHasValue(body, "user") || bodyHasValue(body, "access_tokens") {
		return nil
	}
	return errors.New("provide at least one of user fields or access_tokens")
}

func applyBeaconFraudAmountFlags(cmd *cobra.Command, body map[string]any, currency, value string) error {
	if err := applyStringFlag(cmd, body, "fraud-amount-currency", currency, "fraud_amount", "iso_currency_code"); err != nil {
		return err
	}
	if value != "" || cmd.Flags().Changed("fraud-amount-value") {
		if err := applyDecimalStringFlag(cmd, body, "fraud-amount-value", value, "fraud_amount", "value"); err != nil {
			return err
		}
	}
	if !bodyHasValue(body, "fraud_amount") {
		return nil
	}
	for _, field := range []string{"iso_currency_code", "value"} {
		if bodyHasValue(body, "fraud_amount", field) {
			continue
		}
		return fmt.Errorf("fraud_amount.%s is required", field)
	}
	return nil
}

func newBeaconUserCreateCmd() *cobra.Command {
	var programID, clientUserID string
	var accessTokens []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userFlags *beaconUserIdentityFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /beacon/user/create",
		Long:  "Capability: write. Creates and immediately screens a Beacon user. Use `--body` for depository_accounts or other advanced user fields.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"program_id":     "<program-id>",
				"client_user_id": "<client-user-id>",
				"user": map[string]any{
					"date_of_birth": "1975-01-18",
					"name": map[string]any{
						"given_name":  "Leslie",
						"family_name": "Knope",
					},
				},
			}
			if len(accessTokens) > 0 {
				template["access_tokens"] = accessTokens
			}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "program-id", programID, "program_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "client-user-id", clientUserID, "client_user_id"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "access-token", accessTokens, "access_tokens"); err != nil {
				return err
			}
			if err := applyBeaconUserIdentityFlags(cmd, body, userFlags); err != nil {
				return err
			}
			if err := validateBeaconUserCreateBody(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/user/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	userFlags = bindBeaconUserIdentityFlags(cmd)
	cmd.Flags().StringVar(&programID, "program-id", "", "Beacon program_id")
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable client-side user identifier")
	cmd.Flags().StringSliceVar(&accessTokens, "access-token", nil, "Access token to associate with the Beacon user for Account Insights (repeatable)")
	return cmd
}

func newBeaconUserGetCmd() *cobra.Command {
	var beaconUserID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /beacon/user/get",
		Long:  "Capability: read. Retrieves a Beacon user and current screening status.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"beacon_user_id": "<beacon-user-id>"}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-user-id", beaconUserID, "beacon_user_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-user-id": {"beacon_user_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/user/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user identifier")
	return cmd
}

func newBeaconUserUpdateCmd() *cobra.Command {
	var beaconUserID string
	var accessTokens []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userFlags *beaconUserIdentityFlags

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Call /beacon/user/update",
		Long:  "Capability: write. Updates Beacon user identity data or adds new linked accounts. Use `--body` for depository_accounts or other advanced patch fields.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"beacon_user_id": "<beacon-user-id>",
				"user": map[string]any{
					"email_address": "user@example.com",
				},
			}
			if len(accessTokens) > 0 {
				template["access_tokens"] = accessTokens
			}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-user-id", beaconUserID, "beacon_user_id"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "access-token", accessTokens, "access_tokens"); err != nil {
				return err
			}
			if err := applyBeaconUserIdentityFlags(cmd, body, userFlags); err != nil {
				return err
			}
			if err := validateBeaconUserUpdateBody(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/user/update", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	userFlags = bindBeaconUserIdentityFlags(cmd)
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user identifier")
	cmd.Flags().StringSliceVar(&accessTokens, "access-token", nil, "Access token to add to the Beacon user for evaluation (repeatable)")
	return cmd
}

func newBeaconUserAccountInsightsGetCmd() *cobra.Command {
	var beaconUserID, itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "account-insights-get",
		Short: "Call /beacon/user/account_insights/get",
		Long:  "Capability: read. Retrieves Bank Account Insights for an access token linked to a Beacon user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"beacon_user_id": "<beacon-user-id>",
				"access_token":   "<access-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-user-id", beaconUserID, "beacon_user_id"); err != nil {
				return err
			}
			if _, err := populateAccessToken(cmd, store, body, itemID, accessToken); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-user-id": {"beacon_user_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/user/account_insights/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user identifier")
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	return cmd
}

func newBeaconUserHistoryListCmd() *cobra.Command {
	var beaconUserID, cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "history-list",
		Short: "Call /beacon/user/history/list",
		Long:  "Capability: read. Lists Beacon user revisions in reverse chronological order.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"beacon_user_id": "<beacon-user-id>"}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-user-id", beaconUserID, "beacon_user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-user-id": {"beacon_user_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/user/history/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user identifier")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor from a previous history-list call")
	return cmd
}

func newBeaconReportCreateCmd() *cobra.Command {
	var beaconUserID, reportType, fraudDate, fraudAmountCurrency, fraudAmountValue string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /beacon/report/create",
		Long:  "Capability: write. Creates a Beacon fraud report for a Beacon user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"beacon_user_id": "<beacon-user-id>",
				"type":           "first_party",
				"fraud_date":     "2026-03-08",
			}
			if fraudAmountCurrency != "" || fraudAmountValue != "" {
				template["fraud_amount"] = map[string]any{
					"iso_currency_code": fraudAmountCurrency,
					"value":             fraudAmountValue,
				}
			}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-user-id", beaconUserID, "beacon_user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "type", reportType, "type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "fraud-date", fraudDate, "fraud_date"); err != nil {
				return err
			}
			if err := applyBeaconFraudAmountFlags(cmd, body, fraudAmountCurrency, fraudAmountValue); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-user-id": {"beacon_user_id"},
				"--type":           {"type"},
				"--fraud-date":     {"fraud_date"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/report/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user identifier")
	cmd.Flags().StringVar(&reportType, "type", "", "Beacon report type, e.g. first_party, stolen, synthetic, account_takeover, or unknown")
	cmd.Flags().StringVar(&fraudDate, "fraud-date", "", "Fraud date in YYYY-MM-DD")
	cmd.Flags().StringVar(&fraudAmountCurrency, "fraud-amount-currency", "", "Fraud amount ISO-4217 currency code")
	cmd.Flags().StringVar(&fraudAmountValue, "fraud-amount-value", "", "Fraud amount decimal value")
	return cmd
}

func newBeaconReportGetCmd() *cobra.Command {
	var beaconReportID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /beacon/report/get",
		Long:  "Capability: read. Retrieves a Beacon report by ID.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"beacon_report_id": "<beacon-report-id>"}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-report-id", beaconReportID, "beacon_report_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-report-id": {"beacon_report_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/report/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&beaconReportID, "beacon-report-id", "", "Beacon report identifier")
	return cmd
}

func newBeaconReportListCmd() *cobra.Command {
	var beaconUserID, cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /beacon/report/list",
		Long:  "Capability: read. Lists Beacon reports created for a Beacon user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"beacon_user_id": "<beacon-user-id>"}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-user-id", beaconUserID, "beacon_user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-user-id": {"beacon_user_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/report/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user identifier")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor from a previous report list call")
	return cmd
}

func newBeaconReportSyndicationGetCmd() *cobra.Command {
	var syndicationID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /beacon/report_syndication/get",
		Long:  "Capability: read. Retrieves a Beacon report syndication by ID.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"beacon_report_syndication_id": "<beacon-report-syndication-id>"}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-report-syndication-id", syndicationID, "beacon_report_syndication_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-report-syndication-id": {"beacon_report_syndication_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/report_syndication/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&syndicationID, "beacon-report-syndication-id", "", "Beacon report syndication identifier")
	return cmd
}

func newBeaconReportSyndicationListCmd() *cobra.Command {
	var beaconUserID, cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /beacon/report_syndication/list",
		Long:  "Capability: read. Lists Beacon report syndications for a Beacon user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"beacon_user_id": "<beacon-user-id>"}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-user-id", beaconUserID, "beacon_user_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-user-id": {"beacon_user_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/report_syndication/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user identifier")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor from a previous syndication list call")
	return cmd
}

func newBeaconDuplicateGetCmd() *cobra.Command {
	var duplicateID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /beacon/duplicate/get",
		Long:  "Capability: read. Retrieves a Beacon duplicate record by ID.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"beacon_duplicate_id": "<beacon-duplicate-id>"}
			if handled, err := maybeWriteInfo(cmd, info, beaconDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "beacon-duplicate-id", duplicateID, "beacon_duplicate_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-duplicate-id": {"beacon_duplicate_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/beacon/duplicate/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&duplicateID, "beacon-duplicate-id", "", "Beacon duplicate identifier")
	return cmd
}
