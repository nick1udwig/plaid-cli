package cmd

import (
	"fmt"
	"os"

	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

const assetsDocPath = "docs/plaid/api/products/assets/index.md"

func newAssetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assets",
		Short: "Assets product commands",
		Long:  "Assets commands for creating, retrieving, sharing, and removing Asset Reports.",
	}

	cmd.AddCommand(newAssetsReportCmd())
	cmd.AddCommand(newAssetsAuditCopyCmd())
	cmd.AddCommand(newAssetsRelayCmd())

	return cmd
}

func resolveAssetReportAccessTokens(store *state.Store, itemIDs, accessTokens []string) ([]string, error) {
	if len(itemIDs) == 0 {
		if len(accessTokens) == 0 {
			return nil, fmt.Errorf("provide at least one of --item or --access-token")
		}
		return accessTokens, nil
	}

	resolved := append([]string{}, accessTokens...)
	for _, itemID := range itemIDs {
		record, err := store.LoadItem(itemID)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf("load item %s: %w", itemID, err)
			}
			return nil, err
		}
		resolved = append(resolved, record.AccessToken)
	}
	if len(resolved) == 0 {
		return nil, fmt.Errorf("provide at least one of --item or --access-token")
	}
	return resolved, nil
}

func newAssetsReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Asset Report commands",
		Long:  "Commands for creating, retrieving, refreshing, filtering, and removing Asset Reports.",
	}

	cmd.AddCommand(newAssetsReportCreateCmd())
	cmd.AddCommand(newAssetsReportGetCmd())
	cmd.AddCommand(newAssetsReportPDFGetCmd())
	cmd.AddCommand(newAssetsReportRefreshCmd())
	cmd.AddCommand(newAssetsReportFilterCmd())
	cmd.AddCommand(newAssetsReportRemoveCmd())

	return cmd
}

func newAssetsAuditCopyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "audit-copy",
		Short: "Asset Report audit copy commands",
		Long:  "Commands for granting or revoking audit-copy access to an Asset Report.",
	}

	cmd.AddCommand(newAssetsAuditCopyCreateCmd())
	cmd.AddCommand(newAssetsAuditCopyRemoveCmd())

	return cmd
}

func newAssetsRelayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relay",
		Short: "Asset Report relay token commands",
		Long:  "Commands for creating and managing relay tokens that share Asset Reports with partner clients.",
	}

	cmd.AddCommand(newAssetsRelayCreateCmd())
	cmd.AddCommand(newAssetsRelayGetCmd())
	cmd.AddCommand(newAssetsRelayRefreshCmd())
	cmd.AddCommand(newAssetsRelayRemoveCmd())

	return cmd
}

func newAssetsReportCreateCmd() *cobra.Command {
	var itemIDs, accessTokens, addOns []string
	var daysRequested int
	var webhook, clientReportID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /asset_report/create",
		Long:  "Capability: write. Creates a new Asset Report from one or more linked Items.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_tokens":  []string{"<access-token>"},
				"days_requested": daysRequested,
			}
			options := map[string]any{}
			if webhook != "" {
				options["webhook"] = webhook
			}
			if clientReportID != "" {
				options["client_report_id"] = clientReportID
			}
			if len(addOns) > 0 {
				options["add_ons"] = addOns
			}
			if len(options) > 0 {
				template["options"] = options
			}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if cmd.Flags().Changed("item") || cmd.Flags().Changed("access-token") || !bodyHasValue(body, "access_tokens") {
				resolved, err := resolveAssetReportAccessTokens(store, itemIDs, accessTokens)
				if err != nil {
					return err
				}
				if err := setBodyValue(body, resolved, "access_tokens"); err != nil {
					return err
				}
			}
			if err := applyIntFlag(cmd, body, "days-requested", daysRequested, "days_requested"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "options", "webhook"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "client-report-id", clientReportID, "options", "client_report_id"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "add-on", addOns, "options", "add_ons"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--item/--access-token": {"access_tokens"},
				"--days-requested":      {"days_requested"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/asset_report/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringSliceVar(&itemIDs, "item", nil, "Saved local item_id to include in the report (repeatable)")
	cmd.Flags().StringSliceVar(&accessTokens, "access-token", nil, "Explicit Plaid access_token to include in the report (repeatable)")
	cmd.Flags().IntVar(&daysRequested, "days-requested", 61, "Number of days of history to include in the report")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL for Asset Report readiness events")
	cmd.Flags().StringVar(&clientReportID, "client-report-id", "", "Client-defined report identifier")
	cmd.Flags().StringSliceVar(&addOns, "add-on", nil, "Asset Report add-on, e.g. fast_assets or investments (repeatable)")
	return cmd
}

func newAssetsReportGetCmd() *cobra.Command {
	var assetReportToken string
	var includeInsights, fastReport bool
	var daysToInclude int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /asset_report/get",
		Long:  "Capability: read. Retrieves an Asset Report in JSON form.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"asset_report_token": "<asset-report-token>"}
			if includeInsights {
				template["include_insights"] = true
			}
			if fastReport {
				template["fast_report"] = true
			}
			if daysToInclude > 0 {
				template["options"] = map[string]any{"days_to_include": daysToInclude}
			}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "asset-report-token", assetReportToken, "asset_report_token"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "include-insights", includeInsights, "include_insights"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "fast-report", fastReport, "fast_report"); err != nil {
				return err
			}
			if daysToInclude > 0 {
				if err := applyIntFlag(cmd, body, "days-to-include", daysToInclude, "options", "days_to_include"); err != nil {
					return err
				}
			}
			if err := requireBodyFields(body, map[string][]string{
				"--asset-report-token": {"asset_report_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/asset_report/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&assetReportToken, "asset-report-token", "", "Asset Report token")
	cmd.Flags().BoolVar(&includeInsights, "include-insights", false, "Include Asset Report Insights in the response")
	cmd.Flags().BoolVar(&fastReport, "fast-report", false, "Fetch the fast-assets variant of the report")
	cmd.Flags().IntVar(&daysToInclude, "days-to-include", 0, "Restrict the report to this many recent days")
	return cmd
}

func newAssetsReportPDFGetCmd() *cobra.Command {
	var assetReportToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var outputFlags *binaryOutputFlags

	cmd := &cobra.Command{
		Use:   "pdf-get",
		Short: "Call /asset_report/pdf/get",
		Long:  "Capability: read. Downloads an Asset Report PDF.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"asset_report_token": "<asset-report-token>"}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "asset-report-token", assetReportToken, "asset_report_token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--asset-report-token": {"asset_report_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.CallBytes(ctx, "/asset_report/pdf/get", body)
			if err != nil {
				return err
			}
			return writeBinaryOutput(cmd, outputFlags.out, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	outputFlags = bindBinaryOutputFlag(cmd)
	cmd.Flags().StringVar(&assetReportToken, "asset-report-token", "", "Asset Report token")
	return cmd
}

func newAssetsReportRefreshCmd() *cobra.Command {
	var assetReportToken, webhook, clientReportID string
	var daysRequested int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Call /asset_report/refresh",
		Long:  "Capability: write. Creates a refreshed Asset Report based on an existing report token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"asset_report_token": "<asset-report-token>"}
			if daysRequested > 0 {
				template["days_requested"] = daysRequested
			}
			options := map[string]any{}
			if clientReportID != "" {
				options["client_report_id"] = clientReportID
			}
			if webhook != "" {
				options["webhook"] = webhook
			}
			if len(options) > 0 {
				template["options"] = options
			}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "asset-report-token", assetReportToken, "asset_report_token"); err != nil {
				return err
			}
			if daysRequested > 0 {
				if err := applyIntFlag(cmd, body, "days-requested", daysRequested, "days_requested"); err != nil {
					return err
				}
			}
			if err := applyStringFlag(cmd, body, "client-report-id", clientReportID, "options", "client_report_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "options", "webhook"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--asset-report-token": {"asset_report_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/asset_report/refresh", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&assetReportToken, "asset-report-token", "", "Existing Asset Report token to refresh")
	cmd.Flags().IntVar(&daysRequested, "days-requested", 0, "Override days of history for the refreshed report")
	cmd.Flags().StringVar(&clientReportID, "client-report-id", "", "Client-defined report identifier for the refreshed report")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL for refreshed report readiness events")
	return cmd
}

func newAssetsReportFilterCmd() *cobra.Command {
	var assetReportToken string
	var accountIDsToExclude []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "filter",
		Short: "Call /asset_report/filter",
		Long:  "Capability: write. Creates a filtered Asset Report that excludes specific accounts.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"asset_report_token":     "<asset-report-token>",
				"account_ids_to_exclude": []string{"<account-id>"},
			}
			if len(accountIDsToExclude) > 0 {
				template["account_ids_to_exclude"] = accountIDsToExclude
			}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "asset-report-token", assetReportToken, "asset_report_token"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "account-id-to-exclude", accountIDsToExclude, "account_ids_to_exclude"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--asset-report-token":    {"asset_report_token"},
				"--account-id-to-exclude": {"account_ids_to_exclude"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/asset_report/filter", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&assetReportToken, "asset-report-token", "", "Existing Asset Report token to filter")
	cmd.Flags().StringSliceVar(&accountIDsToExclude, "account-id-to-exclude", nil, "Account ID to exclude from the new report (repeatable)")
	return cmd
}

func newAssetsReportRemoveCmd() *cobra.Command {
	var assetReportToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Call /asset_report/remove",
		Long:  "Capability: write. Removes an Asset Report and invalidates its token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"asset_report_token": "<asset-report-token>"}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "asset-report-token", assetReportToken, "asset_report_token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--asset-report-token": {"asset_report_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/asset_report/remove", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&assetReportToken, "asset-report-token", "", "Asset Report token to remove")
	return cmd
}

func newAssetsAuditCopyCreateCmd() *cobra.Command {
	var assetReportToken, auditorID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /asset_report/audit_copy/create",
		Long:  "Capability: write. Creates an audit-copy token for sharing an Asset Report with an auditor.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"asset_report_token": "<asset-report-token>",
				"auditor_id":         "<auditor-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "asset-report-token", assetReportToken, "asset_report_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "auditor-id", auditorID, "auditor_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--asset-report-token": {"asset_report_token"},
				"--auditor-id":         {"auditor_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/asset_report/audit_copy/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&assetReportToken, "asset-report-token", "", "Asset Report token to share")
	cmd.Flags().StringVar(&auditorID, "auditor-id", "", "Plaid auditor ID, e.g. fannie_mae")
	return cmd
}

func newAssetsAuditCopyRemoveCmd() *cobra.Command {
	var auditCopyToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Call /asset_report/audit_copy/remove",
		Long:  "Capability: write. Revokes an audit-copy token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"audit_copy_token": "<audit-copy-token>"}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "audit-copy-token", auditCopyToken, "audit_copy_token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--audit-copy-token": {"audit_copy_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/asset_report/audit_copy/remove", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&auditCopyToken, "audit-copy-token", "", "Audit-copy token to revoke")
	return cmd
}

func newAssetsRelayCreateCmd() *cobra.Command {
	var reportTokens []string
	var secondaryClientID, webhook string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /credit/relay/create",
		Long:  "Capability: write. Creates a relay token for sharing an Asset Report with a partner client.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"report_tokens":       []string{"<asset-report-token>"},
				"secondary_client_id": "<secondary-client-id>",
			}
			if webhook != "" {
				template["webhook"] = webhook
			}
			if len(reportTokens) > 0 {
				template["report_tokens"] = reportTokens
			}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringSliceFlag(cmd, body, "report-token", reportTokens, "report_tokens"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "secondary-client-id", secondaryClientID, "secondary_client_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--report-token":        {"report_tokens"},
				"--secondary-client-id": {"secondary_client_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/relay/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringSliceVar(&reportTokens, "report-token", nil, "Report token to relay, usually an Asset Report token (repeatable)")
	cmd.Flags().StringVar(&secondaryClientID, "secondary-client-id", "", "Partner client_id that will receive the shared report")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL for relay refresh events")
	return cmd
}

func newAssetsRelayGetCmd() *cobra.Command {
	var relayToken, reportType string
	var includeInsights bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /credit/relay/get",
		Long:  "Capability: read. Retrieves a relayed Asset Report using a relay token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"relay_token": "<relay-token>",
				"report_type": "asset",
			}
			if includeInsights {
				template["include_insights"] = true
			}
			if reportType != "" {
				template["report_type"] = reportType
			}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "relay-token", relayToken, "relay_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "report-type", reportType, "report_type"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "include-insights", includeInsights, "include_insights"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--relay-token": {"relay_token"},
				"--report-type": {"report_type"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/relay/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&relayToken, "relay-token", "", "Relay token created by /credit/relay/create")
	cmd.Flags().StringVar(&reportType, "report-type", "asset", "Report type to fetch, currently asset")
	cmd.Flags().BoolVar(&includeInsights, "include-insights", false, "Include Asset Report Insights in the relayed report")
	return cmd
}

func newAssetsRelayRefreshCmd() *cobra.Command {
	var relayToken, reportType, webhook string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Call /credit/relay/refresh",
		Long:  "Capability: write. Refreshes the report referenced by a relay token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"relay_token": "<relay-token>",
				"report_type": "asset",
			}
			if webhook != "" {
				template["webhook"] = webhook
			}
			if reportType != "" {
				template["report_type"] = reportType
			}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "relay-token", relayToken, "relay_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "report-type", reportType, "report_type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--relay-token": {"relay_token"},
				"--report-type": {"report_type"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/relay/refresh", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&relayToken, "relay-token", "", "Relay token created by /credit/relay/create")
	cmd.Flags().StringVar(&reportType, "report-type", "asset", "Report type to refresh, currently asset")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL for refreshed relay reports")
	return cmd
}

func newAssetsRelayRemoveCmd() *cobra.Command {
	var relayToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Call /credit/relay/remove",
		Long:  "Capability: write. Removes a relay token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"relay_token": "<relay-token>"}
			if handled, err := maybeWriteInfo(cmd, info, assetsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "relay-token", relayToken, "relay_token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--relay-token": {"relay_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/relay/remove", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&relayToken, "relay-token", "", "Relay token to remove")
	return cmd
}
