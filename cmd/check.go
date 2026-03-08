package cmd

import "github.com/spf13/cobra"

const checkDocPath = "docs/plaid/api/products/check/index.md"

func newCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Plaid Check commands",
		Long:  "Plaid Check commands for consumer report generation, retrieval, and monitoring insights.",
	}

	cmd.AddCommand(newCheckReportCmd())
	cmd.AddCommand(newCheckMonitoringCmd())

	return cmd
}

func newCheckReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Consumer report commands",
		Long:  "Commands for creating and retrieving Plaid Check consumer reports and related derived reports.",
	}

	cmd.AddCommand(newCheckReportCreateCmd())
	cmd.AddCommand(newCheckReportBaseGetCmd())
	cmd.AddCommand(newCheckReportIncomeInsightsGetCmd())
	cmd.AddCommand(newCheckReportNetworkInsightsGetCmd())
	cmd.AddCommand(newCheckReportPartnerInsightsGetCmd())
	cmd.AddCommand(newCheckReportPDFGetCmd())
	cmd.AddCommand(newCheckReportCashflowInsightsGetCmd())
	cmd.AddCommand(newCheckReportLendScoreGetCmd())
	cmd.AddCommand(newCheckReportVerificationGetCmd())
	cmd.AddCommand(newCheckReportVerificationPDFGetCmd())

	return cmd
}

func newCheckMonitoringCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitoring",
		Short: "Monitoring insights commands",
		Long:  "Commands for fetching, subscribing to, and unsubscribing from Plaid Check monitoring insights.",
	}

	cmd.AddCommand(newCheckMonitoringGetCmd())
	cmd.AddCommand(newCheckMonitoringSubscribeCmd())
	cmd.AddCommand(newCheckMonitoringUnsubscribeCmd())

	return cmd
}

func newCheckReportCreateCmd() *cobra.Command {
	var userID, userToken string
	var webhook, clientReportID string
	var products, gseReportTypes []string
	var daysRequested, daysRequired int
	var requireIdentity bool
	var cashflowAttributesVersion string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /cra/check_report/create",
		Long:  "Capability: write. Creates or refreshes a Plaid Check consumer report and optionally pre-generates add-on report products.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id":        "<user-id>",
				"webhook":        "https://example.com/webhook",
				"days_requested": 365,
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "days-requested", daysRequested, "days_requested"); err != nil {
				return err
			}
			if cmd.Flags().Changed("days-required") {
				if err := setBodyValue(body, daysRequired, "days_required"); err != nil {
					return err
				}
			}
			if err := applyStringFlag(cmd, body, "client-report-id", clientReportID, "client_report_id"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "product", products, "products"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "require-identity", requireIdentity, "base_report", "require_identity"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "gse-report-type", gseReportTypes, "base_report", "gse_options", "report_types"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cashflow-attributes-version", cashflowAttributesVersion, "cashflow_insights", "attributes_version"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--webhook":        {"webhook"},
				"--days-requested": {"days_requested"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/check_report/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL for report readiness notifications")
	cmd.Flags().IntVar(&daysRequested, "days-requested", 365, "Number of days of bank data to request")
	cmd.Flags().IntVar(&daysRequired, "days-required", 0, "Optional minimum number of days required for report generation")
	cmd.Flags().StringVar(&clientReportID, "client-report-id", "", "Client-generated identifier for this report")
	cmd.Flags().StringSliceVar(&products, "product", nil, "Plaid Check add-on product to pre-generate (repeatable)")
	cmd.Flags().BoolVar(&requireIdentity, "require-identity", false, "Require identity information in the base report")
	cmd.Flags().StringSliceVar(&gseReportTypes, "gse-report-type", nil, "GSE report type to include under base_report.gse_options.report_types (repeatable)")
	cmd.Flags().StringVar(&cashflowAttributesVersion, "cashflow-attributes-version", "", "Optional Cashflow Insights attributes version, e.g. CFI1")
	return cmd
}

func newCheckReportBaseGetCmd() *cobra.Command {
	var userID, userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "base-get",
		Short: "Call /cra/check_report/base_report/get",
		Long:  "Capability: read. Retrieves the base Plaid Check consumer report for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/check_report/base_report/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	return cmd
}

func newCheckReportIncomeInsightsGetCmd() *cobra.Command {
	var userID, userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "income-insights-get",
		Short: "Call /cra/check_report/income_insights/get",
		Long:  "Capability: read. Retrieves the Income Insights report for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/check_report/income_insights/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	return cmd
}

func newCheckReportNetworkInsightsGetCmd() *cobra.Command {
	var userID, userToken, version string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "network-insights-get",
		Short: "Call /cra/check_report/network_insights/get",
		Long:  "Capability: read. Retrieves Plaid Check network insights for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if version != "" {
				template["options"] = map[string]any{"network_insights_version": version}
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "version", version, "options", "network_insights_version"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/check_report/network_insights/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringVar(&version, "version", "", "Network insights version, e.g. NI1")
	return cmd
}

func newCheckReportPartnerInsightsGetCmd() *cobra.Command {
	var userID, userToken string
	var firstDetectVersion, detectVersion, cashScoreVersion, extendVersion, insightsVersion string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "partner-insights-get",
		Short: "Call /cra/check_report/partner_insights/get",
		Long:  "Capability: read. Retrieves partner-derived Plaid Check insights for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "firstdetect-version", firstDetectVersion, "partner_insights", "prism_versions", "firstdetect"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "detect-version", detectVersion, "partner_insights", "prism_versions", "detect"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "cashscore-version", cashScoreVersion, "partner_insights", "prism_versions", "cashscore"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "extend-version", extendVersion, "partner_insights", "prism_versions", "extend"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "insights-version", insightsVersion, "partner_insights", "prism_versions", "insights"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/check_report/partner_insights/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringVar(&firstDetectVersion, "firstdetect-version", "", "Partner insights FirstDetect version")
	cmd.Flags().StringVar(&detectVersion, "detect-version", "", "Partner insights Detect version")
	cmd.Flags().StringVar(&cashScoreVersion, "cashscore-version", "", "Partner insights CashScore version")
	cmd.Flags().StringVar(&extendVersion, "extend-version", "", "Partner insights Extend version")
	cmd.Flags().StringVar(&insightsVersion, "insights-version", "", "Partner insights Insights version")
	return cmd
}

func newCheckReportPDFGetCmd() *cobra.Command {
	var userID, userToken string
	var addOns []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var outFlags *binaryOutputFlags

	cmd := &cobra.Command{
		Use:   "pdf-get",
		Short: "Call /cra/check_report/pdf/get",
		Long:  "Capability: read. Downloads the latest Plaid Check consumer report as a PDF.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if len(addOns) > 0 {
				template["add_ons"] = addOns
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "add-on", addOns, "add_ons"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.CallBytes(ctx, "/cra/check_report/pdf/get", body)
			if err != nil {
				return err
			}
			return writeBinaryOutput(cmd, outFlags.out, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	outFlags = bindBinaryOutputFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringSliceVar(&addOns, "add-on", nil, "Report add-on to include in the PDF, e.g. cra_income_insights (repeatable)")
	return cmd
}

func newCheckReportCashflowInsightsGetCmd() *cobra.Command {
	var userID, userToken, attributesVersion string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "cashflow-insights-get",
		Short: "Call /cra/check_report/cashflow_insights/get",
		Long:  "Capability: read. Retrieves Plaid Check cashflow insights for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if attributesVersion != "" {
				template["options"] = map[string]any{"attributes_version": attributesVersion}
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "attributes-version", attributesVersion, "options", "attributes_version"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/check_report/cashflow_insights/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringVar(&attributesVersion, "attributes-version", "", "Cashflow Insights attributes version, e.g. CFI1")
	return cmd
}

func newCheckReportLendScoreGetCmd() *cobra.Command {
	var userID, userToken, lendScoreVersion string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "lend-score-get",
		Short: "Call /cra/check_report/lend_score/get",
		Long:  "Capability: read. Retrieves the Plaid Check LendScore report for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_id": "<user-id>"}
			if lendScoreVersion != "" {
				template["options"] = map[string]any{"lend_score_version": lendScoreVersion}
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "version", lendScoreVersion, "options", "lend_score_version"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/check_report/lend_score/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringVar(&lendScoreVersion, "version", "", "LendScore version, e.g. LS1")
	return cmd
}

func newCheckReportVerificationGetCmd() *cobra.Command {
	var userID, userToken string
	var reportsRequested []string
	var employmentRefreshDays int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "verification-get",
		Short: "Call /cra/check_report/verification/get",
		Long:  "Capability: read. Retrieves home lending verification reports such as VOA and Employment Refresh.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id":           "<user-id>",
				"reports_requested": []string{"VOA"},
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "report-requested", reportsRequested, "reports_requested"); err != nil {
				return err
			}
			if cmd.Flags().Changed("employment-refresh-days-requested") {
				if err := setBodyValue(body, employmentRefreshDays, "employment_refresh_options", "days_requested"); err != nil {
					return err
				}
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--report-requested": {"reports_requested"},
			}); err != nil {
				return err
			}
			if bodyStringSliceContains(body, "EMPLOYMENT_REFRESH", "reports_requested") &&
				!bodyHasValue(body, "employment_refresh_options", "days_requested") {
				return requireBodyFields(body, map[string][]string{
					"--employment-refresh-days-requested": {"employment_refresh_options", "days_requested"},
				})
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/check_report/verification/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringSliceVar(&reportsRequested, "report-requested", nil, "Verification report to retrieve: VOA or EMPLOYMENT_REFRESH (repeatable)")
	cmd.Flags().IntVar(&employmentRefreshDays, "employment-refresh-days-requested", 0, "Days of payroll data to request when EMPLOYMENT_REFRESH is requested")
	return cmd
}

func newCheckReportVerificationPDFGetCmd() *cobra.Command {
	var userID, userToken, reportRequested string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var outFlags *binaryOutputFlags

	cmd := &cobra.Command{
		Use:   "verification-pdf-get",
		Short: "Call /cra/check_report/verification/pdf/get",
		Long:  "Capability: read. Downloads a home-lending verification PDF such as VOA or Employment Refresh.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id":          "<user-id>",
				"report_requested": "voa",
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "report-requested", reportRequested, "report_requested"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--report-requested": {"report_requested"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.CallBytes(ctx, "/cra/check_report/verification/pdf/get", body)
			if err != nil {
				return err
			}
			return writeBinaryOutput(cmd, outFlags.out, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	outFlags = bindBinaryOutputFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringVar(&reportRequested, "report-requested", "", "Verification PDF to fetch: voa or employment_refresh")
	return cmd
}

func newCheckMonitoringGetCmd() *cobra.Command {
	var userID, userToken, purpose string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /cra/monitoring_insights/get",
		Long:  "Capability: read. Retrieves the latest monitoring insights report for a Plaid Check user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id":                             "<user-id>",
				"consumer_report_permissible_purpose": "ACCOUNT_REVIEW_CREDIT",
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "purpose", purpose, "consumer_report_permissible_purpose"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--purpose": {"consumer_report_permissible_purpose"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/monitoring_insights/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringVar(&purpose, "purpose", "", "Consumer report permissible purpose, e.g. ACCOUNT_REVIEW_CREDIT")
	return cmd
}

func newCheckMonitoringSubscribeCmd() *cobra.Command {
	var userID, userToken, itemID, webhook string
	var incomeCategories []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: "Call /cra/monitoring_insights/subscribe",
		Long:  "Capability: write. Subscribes a user and optionally a specific item to monitoring insights webhooks.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_id": "<user-id>",
				"webhook": "https://example.com/webhook",
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyCheckUserReference(cmd, body, userID, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "item-id", itemID, "item_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "webhook"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "income-category", incomeCategories, "income_categories"); err != nil {
				return err
			}
			if err := requireCheckUserReference(body); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--webhook": {"webhook"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/monitoring_insights/subscribe", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	bindCheckUserReferenceFlags(cmd, &userID, &userToken)
	cmd.Flags().StringVar(&itemID, "item-id", "", "Optional Plaid item_id to scope the subscription to")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL for monitoring updates")
	cmd.Flags().StringSliceVar(&incomeCategories, "income-category", nil, "Income category to include in monitoring insights (repeatable)")
	return cmd
}

func newCheckMonitoringUnsubscribeCmd() *cobra.Command {
	var subscriptionID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "unsubscribe",
		Short: "Call /cra/monitoring_insights/unsubscribe",
		Long:  "Capability: write. Unsubscribes a prior monitoring insights subscription.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"subscription_id": "<subscription-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, checkDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "subscription-id", subscriptionID, "subscription_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--subscription-id": {"subscription_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/cra/monitoring_insights/unsubscribe", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&subscriptionID, "subscription-id", "", "Plaid monitoring insights subscription_id")
	return cmd
}

func bindCheckUserReferenceFlags(cmd *cobra.Command, userID, userToken *string) {
	cmd.Flags().StringVar(userID, "user-id", "", "Plaid user_id for new Plaid Check integrations")
	cmd.Flags().StringVar(userToken, "user-token", "", "Legacy Plaid user_token for older Plaid Check integrations")
}

func applyCheckUserReference(cmd *cobra.Command, body map[string]any, userID, userToken string) error {
	if err := applyStringFlag(cmd, body, "user-id", userID, "user_id"); err != nil {
		return err
	}
	return applyStringFlag(cmd, body, "user-token", userToken, "user_token")
}

func requireCheckUserReference(body map[string]any) error {
	return requireAtLeastOneBodyField(body, map[string][]string{
		"--user-id":    {"user_id"},
		"--user-token": {"user_token"},
	})
}

func bodyStringSliceContains(body map[string]any, target string, path ...string) bool {
	raw, ok := bodyValue(body, path...)
	if !ok {
		return false
	}

	switch values := raw.(type) {
	case []string:
		for _, value := range values {
			if value == target {
				return true
			}
		}
	case []any:
		for _, value := range values {
			text, ok := value.(string)
			if ok && text == target {
				return true
			}
		}
	}
	return false
}
