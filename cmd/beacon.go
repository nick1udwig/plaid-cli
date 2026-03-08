package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

const beaconDocPath = "docs/plaid/api/products/beacon/index.md"

type beaconUserFlags struct {
	givenName         string
	familyName        string
	dateOfBirth       string
	emailAddress      string
	phoneNumber       string
	addressStreet     string
	addressStreet2    string
	addressCity       string
	addressRegion     string
	addressPostalCode string
	addressCountry    string
	depositoryAccount string
	depositoryRouting string
	accessTokens      []string
}

func newBeaconCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beacon",
		Short: "Beacon product commands",
		Long:  "Beacon commands for user screening, fraud reporting, and network syndication lookups.",
	}

	cmd.AddCommand(newBeaconUserCmd())
	cmd.AddCommand(newBeaconReportCmd())
	cmd.AddCommand(newBeaconReportSyndicationCmd())
	cmd.AddCommand(newBeaconDuplicateGetCmd())

	return cmd
}

func newBeaconUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Beacon user commands",
		Long:  "Create, update, inspect, and review Beacon users.",
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
		Long:  "Create and inspect Beacon fraud reports.",
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
		Long:  "Inspect Beacon report syndications matched to a Beacon user.",
	}

	cmd.AddCommand(newBeaconReportSyndicationGetCmd())
	cmd.AddCommand(newBeaconReportSyndicationListCmd())

	return cmd
}

func bindBeaconUserFlags(cmd *cobra.Command) *beaconUserFlags {
	flags := &beaconUserFlags{}
	cmd.Flags().StringVar(&flags.givenName, "given-name", "", "User given name")
	cmd.Flags().StringVar(&flags.familyName, "family-name", "", "User family name")
	cmd.Flags().StringVar(&flags.dateOfBirth, "date-of-birth", "", "User date of birth in YYYY-MM-DD")
	cmd.Flags().StringVar(&flags.emailAddress, "email-address", "", "User email address")
	cmd.Flags().StringVar(&flags.phoneNumber, "phone-number", "", "User phone number in E.164 format")
	cmd.Flags().StringVar(&flags.addressStreet, "address-street", "", "User street address")
	cmd.Flags().StringVar(&flags.addressStreet2, "address-street2", "", "User address line 2")
	cmd.Flags().StringVar(&flags.addressCity, "address-city", "", "User address city")
	cmd.Flags().StringVar(&flags.addressRegion, "address-region", "", "User address region or state")
	cmd.Flags().StringVar(&flags.addressPostalCode, "address-postal-code", "", "User address postal code")
	cmd.Flags().StringVar(&flags.addressCountry, "address-country", "", "User address ISO country code")
	cmd.Flags().StringVar(&flags.depositoryAccount, "depository-account-number", "", "User depository account number")
	cmd.Flags().StringVar(&flags.depositoryRouting, "depository-routing-number", "", "User depository routing number")
	cmd.Flags().StringSliceVar(&flags.accessTokens, "access-token", nil, "Plaid access_token to associate with the Beacon user (repeatable)")
	return flags
}

func applyBeaconUserFlags(cmd *cobra.Command, body map[string]any, flags *beaconUserFlags) error {
	if flags == nil {
		return nil
	}

	if err := applyStringFlag(cmd, body, "given-name", flags.givenName, "user", "name", "given_name"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "family-name", flags.familyName, "user", "name", "family_name"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "date-of-birth", flags.dateOfBirth, "user", "date_of_birth"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "email-address", flags.emailAddress, "user", "email_address"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "phone-number", flags.phoneNumber, "user", "phone_number"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "address-street", flags.addressStreet, "user", "address", "street"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "address-street2", flags.addressStreet2, "user", "address", "street2"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "address-city", flags.addressCity, "user", "address", "city"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "address-region", flags.addressRegion, "user", "address", "region"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "address-postal-code", flags.addressPostalCode, "user", "address", "postal_code"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "address-country", flags.addressCountry, "user", "address", "country"); err != nil {
		return err
	}
	if err := applyStringSliceFlag(cmd, body, "access-token", flags.accessTokens, "access_tokens"); err != nil {
		return err
	}

	if !cmd.Flags().Changed("depository-account-number") && !cmd.Flags().Changed("depository-routing-number") &&
		!bodyHasValue(body, "user", "depository_accounts") {
		return nil
	}

	entry := map[string]any{}
	if flags.depositoryAccount != "" {
		entry["account_number"] = flags.depositoryAccount
	}
	if flags.depositoryRouting != "" {
		entry["routing_number"] = flags.depositoryRouting
	}
	if len(entry) == 0 && !bodyHasValue(body, "user", "depository_accounts") {
		return nil
	}
	if len(entry) > 0 {
		return setBodyValue(body, []any{entry}, "user", "depository_accounts")
	}
	return nil
}

func validateBeaconUser(body map[string]any, requireIdentifiers bool) error {
	if requireIdentifiers {
		if !bodyHasValue(body, "user", "name", "given_name") {
			return errors.New("user.name.given_name is required")
		}
		if !bodyHasValue(body, "user", "name", "family_name") {
			return errors.New("user.name.family_name is required")
		}
		if !bodyHasValue(body, "user", "date_of_birth") && !bodyHasValue(body, "user", "depository_accounts") {
			return errors.New("provide at least one of user.date_of_birth or user.depository_accounts")
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

	if bodyHasValue(body, "user", "depository_accounts") {
		entry, ok, err := firstObjectFromArrayPath(body, "user", "depository_accounts")
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("user.depository_accounts must contain at least one object")
		}
		for _, field := range []string{"account_number", "routing_number"} {
			if _, ok := entry[field]; ok {
				continue
			}
			return fmt.Errorf("user.depository_accounts[0].%s is required", field)
		}
	}

	return nil
}

func beaconUserHasPatch(body map[string]any) bool {
	return bodyHasValue(body, "user") || bodyHasValue(body, "access_tokens")
}

func newBeaconUserCreateCmd() *cobra.Command {
	var programID, clientUserID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userFlags *beaconUserFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /beacon/user/create",
		Long:  "Capability: write. Creates and immediately screens a Beacon user against your Beacon program.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"program_id":     "<program-id>",
				"client_user_id": "<client-user-id>",
				"user": map[string]any{
					"name": map[string]any{
						"given_name":  "Jane",
						"family_name": "Doe",
					},
					"date_of_birth": "1990-01-01",
				},
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
			if err := applyBeaconUserFlags(cmd, body, userFlags); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--program-id":     {"program_id"},
				"--client-user-id": {"client_user_id"},
			}); err != nil {
				return err
			}
			if err := validateBeaconUser(body, true); err != nil {
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
	userFlags = bindBeaconUserFlags(cmd)
	cmd.Flags().StringVar(&programID, "program-id", "", "Beacon program ID")
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable client-side identifier for the end user")
	return cmd
}

func newBeaconUserGetCmd() *cobra.Command {
	var beaconUserID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /beacon/user/get",
		Long:  "Capability: read. Retrieves a Beacon user by ID.",
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
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user ID")
	return cmd
}

func newBeaconUserUpdateCmd() *cobra.Command {
	var beaconUserID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userFlags *beaconUserFlags

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Call /beacon/user/update",
		Long:  "Capability: write. Updates Beacon user identity data or linked access tokens and immediately re-screens the user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"beacon_user_id": "<beacon-user-id>",
				"user": map[string]any{
					"email_address": "jane@example.com",
				},
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
			if err := applyBeaconUserFlags(cmd, body, userFlags); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-user-id": {"beacon_user_id"},
			}); err != nil {
				return err
			}
			if !beaconUserHasPatch(body) {
				return errors.New("provide at least one Beacon user field or --access-token")
			}
			if err := validateBeaconUser(body, false); err != nil {
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
	userFlags = bindBeaconUserFlags(cmd)
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user ID")
	return cmd
}

func newBeaconUserAccountInsightsGetCmd() *cobra.Command {
	var beaconUserID, itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "account-insights-get",
		Short: "Call /beacon/user/account_insights/get",
		Long:  "Capability: read. Retrieves account insights for accounts linked to a Beacon user.",
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
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user ID")
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use for the access token")
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
		Long:  "Capability: read. Lists the revision history for a Beacon user.",
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
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user ID")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor")
	return cmd
}

func newBeaconReportCreateCmd() *cobra.Command {
	var beaconUserID, reportType, fraudDate, amountCurrency, amountValue string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /beacon/report/create",
		Long:  "Capability: write. Creates a Beacon fraud report for a Beacon user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"beacon_user_id": "<beacon-user-id>",
				"type":           "stolen",
				"fraud_date":     "2026-01-15",
			}
			if amountCurrency != "" || amountValue != "" {
				template["fraud_amount"] = map[string]any{
					"iso_currency_code": "USD",
					"value":             100.0,
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
			if err := applyStringFlag(cmd, body, "amount-currency", amountCurrency, "fraud_amount", "iso_currency_code"); err != nil {
				return err
			}
			if err := applyDecimalStringFlag(cmd, body, "amount-value", amountValue, "fraud_amount", "value"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--beacon-user-id": {"beacon_user_id"},
				"--type":           {"type"},
				"--fraud-date":     {"fraud_date"},
			}); err != nil {
				return err
			}
			if bodyHasValue(body, "fraud_amount") {
				for _, field := range []string{"iso_currency_code", "value"} {
					if bodyHasValue(body, "fraud_amount", field) {
						continue
					}
					return fmt.Errorf("fraud_amount.%s is required", field)
				}
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
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user ID")
	cmd.Flags().StringVar(&reportType, "type", "", "Beacon report type, e.g. first_party, stolen, synthetic, account_takeover, or unknown")
	cmd.Flags().StringVar(&fraudDate, "fraud-date", "", "Fraud date in YYYY-MM-DD")
	cmd.Flags().StringVar(&amountCurrency, "amount-currency", "", "Fraud amount currency code")
	cmd.Flags().StringVar(&amountValue, "amount-value", "", "Fraud amount value")
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
	cmd.Flags().StringVar(&beaconReportID, "beacon-report-id", "", "Beacon report ID")
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
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user ID")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor")
	return cmd
}

func newBeaconReportSyndicationGetCmd() *cobra.Command {
	var beaconReportSyndicationID string
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
			if err := applyStringFlag(cmd, body, "beacon-report-syndication-id", beaconReportSyndicationID, "beacon_report_syndication_id"); err != nil {
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
	cmd.Flags().StringVar(&beaconReportSyndicationID, "beacon-report-syndication-id", "", "Beacon report syndication ID")
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
	cmd.Flags().StringVar(&beaconUserID, "beacon-user-id", "", "Beacon user ID")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor")
	return cmd
}

func newBeaconDuplicateGetCmd() *cobra.Command {
	var beaconDuplicateID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "duplicate-get",
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
			if err := applyStringFlag(cmd, body, "beacon-duplicate-id", beaconDuplicateID, "beacon_duplicate_id"); err != nil {
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
	cmd.Flags().StringVar(&beaconDuplicateID, "beacon-duplicate-id", "", "Beacon duplicate ID")
	return cmd
}
