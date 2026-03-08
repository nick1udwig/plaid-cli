package cmd

import "github.com/spf13/cobra"

const incomeDocPath = "docs/plaid/api/products/income/index.md"

func newIncomeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "income",
		Short: "Income product commands",
		Long:  "Income verification commands. Most are read-only; `bank-income-refresh`, `payroll-income-refresh`, and `payroll-income-parsing-config-update` perform write-style requests. The current docs snapshot exposes these endpoints with `user_token` request fields.",
	}

	cmd.AddCommand(newIncomeSessionsGetCmd())
	cmd.AddCommand(newIncomeBankIncomeGetCmd())
	cmd.AddCommand(newIncomeBankIncomePDFGetCmd())
	cmd.AddCommand(newIncomeBankIncomeRefreshCmd())
	cmd.AddCommand(newIncomeBankStatementsUploadsGetCmd())
	cmd.AddCommand(newIncomePayrollIncomeGetCmd())
	cmd.AddCommand(newIncomePayrollIncomeRiskSignalsGetCmd())
	cmd.AddCommand(newIncomeEmploymentGetCmd())
	cmd.AddCommand(newIncomePayrollIncomeParsingConfigUpdateCmd())
	cmd.AddCommand(newIncomePayrollIncomeRefreshCmd())

	return cmd
}

func applyIncomeUserToken(cmd *cobra.Command, body map[string]any, userToken string) error {
	if err := applyStringFlag(cmd, body, "user-token", userToken, "user_token"); err != nil {
		return err
	}
	return requireBodyFields(body, map[string][]string{
		"--user-token": {"user_token"},
	})
}

func newIncomeSessionsGetCmd() *cobra.Command {
	var userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "sessions-get",
		Short: "Call /credit/sessions/get",
		Long:  "Capability: read. Retrieves completed Link sessions and result metadata for an Income user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_token": "<user-token>"}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/sessions/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	return cmd
}

func newIncomeBankIncomeGetCmd() *cobra.Command {
	var userToken string
	var count int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "bank-income-get",
		Short: "Call /credit/bank_income/get",
		Long:  "Capability: read. Retrieves bank income reports for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_token": "<user-token>",
				"options": map[string]any{
					"count": count,
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "options", "count"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/bank_income/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	cmd.Flags().IntVar(&count, "count", 1, "Number of bank income reports to return")
	return cmd
}

func newIncomeBankIncomePDFGetCmd() *cobra.Command {
	var userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var outputFlags *binaryOutputFlags

	cmd := &cobra.Command{
		Use:   "bank-income-pdf-get",
		Short: "Call /credit/bank_income/pdf/get",
		Long:  "Capability: read. Downloads the most recent bank income report PDF for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_token": "<user-token>"}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.CallBytes(ctx, "/credit/bank_income/pdf/get", body)
			if err != nil {
				return err
			}
			return writeBinaryOutput(cmd, outputFlags.out, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	outputFlags = bindBinaryOutputFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	return cmd
}

func newIncomeBankIncomeRefreshCmd() *cobra.Command {
	var userToken string
	var daysRequested int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "bank-income-refresh",
		Short: "Call /credit/bank_income/refresh",
		Long:  "Capability: write. Refreshes the most recent bank income report for a user.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_token": "<user-token>"}
			if daysRequested > 0 {
				template["options"] = map[string]any{"days_requested": daysRequested}
			}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}
			if daysRequested > 0 {
				if err := applyIntFlag(cmd, body, "days-requested", daysRequested, "options", "days_requested"); err != nil {
					return err
				}
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/bank_income/refresh", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	cmd.Flags().IntVar(&daysRequested, "days-requested", 0, "Override days of data to include in the refreshed report")
	return cmd
}

func newIncomeBankStatementsUploadsGetCmd() *cobra.Command {
	var userToken string
	var itemIDs []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "bank-statements-uploads-get",
		Short: "Call /credit/bank_statements/uploads/get",
		Long:  "Capability: read. Retrieves parsed data from bank statements uploaded during Document Income flows.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_token": "<user-token>"}
			if len(itemIDs) > 0 {
				template["options"] = map[string]any{"item_ids": itemIDs}
			}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "item-id", itemIDs, "options", "item_ids"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/bank_statements/uploads/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	cmd.Flags().StringSliceVar(&itemIDs, "item-id", nil, "Document Income item_id filter (repeatable)")
	return cmd
}

func newIncomePayrollIncomeGetCmd() *cobra.Command {
	var userToken string
	var itemIDs []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "payroll-income-get",
		Short: "Call /credit/payroll_income/get",
		Long:  "Capability: read. Retrieves payroll income data from linked payroll providers or uploaded documents.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_token": "<user-token>"}
			if len(itemIDs) > 0 {
				template["options"] = map[string]any{"item_ids": itemIDs}
			}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "item-id", itemIDs, "options", "item_ids"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/payroll_income/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	cmd.Flags().StringSliceVar(&itemIDs, "item-id", nil, "Payroll item_id filter (repeatable)")
	return cmd
}

func newIncomePayrollIncomeRiskSignalsGetCmd() *cobra.Command {
	var userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "payroll-income-risk-signals-get",
		Short: "Call /credit/payroll_income/risk_signals/get",
		Long:  "Capability: read. Retrieves fraud-risk signals for manually uploaded income documents.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_token": "<user-token>"}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/payroll_income/risk_signals/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	return cmd
}

func newIncomeEmploymentGetCmd() *cobra.Command {
	var userToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "employment-get",
		Short: "Call /credit/employment/get",
		Long:  "Capability: read. Retrieves employment information verified through payroll providers.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_token": "<user-token>"}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/employment/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	return cmd
}

func newIncomePayrollIncomeParsingConfigUpdateCmd() *cobra.Command {
	var userToken, itemID string
	var parsingConfig []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "payroll-income-parsing-config-update",
		Short: "Call /credit/payroll_income/parsing_config/update",
		Long:  "Capability: write. Updates the parsing configuration for a document income verification session.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user_token":     "<user-token>",
				"parsing_config": []string{"ocr"},
			}
			if itemID != "" {
				template["item_id"] = itemID
			}
			if len(parsingConfig) > 0 {
				template["parsing_config"] = parsingConfig
			}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "item-id", itemID, "item_id"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "parsing-config", parsingConfig, "parsing_config"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--parsing-config": {"parsing_config"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/payroll_income/parsing_config/update", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	cmd.Flags().StringVar(&itemID, "item-id", "", "Payroll income item_id to target")
	cmd.Flags().StringSliceVar(&parsingConfig, "parsing-config", nil, "Parsing analyses to enable, e.g. ocr or risk_signals (repeatable)")
	return cmd
}

func newIncomePayrollIncomeRefreshCmd() *cobra.Command {
	var userToken, webhook string
	var itemIDs []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "payroll-income-refresh",
		Short: "Call /credit/payroll_income/refresh",
		Long:  "Capability: write. Refreshes payroll income data for one or more payroll items.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"user_token": "<user-token>"}
			options := map[string]any{}
			if len(itemIDs) > 0 {
				options["item_ids"] = itemIDs
			}
			if webhook != "" {
				options["webhook"] = webhook
			}
			if len(options) > 0 {
				template["options"] = options
			}
			if handled, err := maybeWriteInfo(cmd, info, incomeDocPath, template); handled || err != nil {
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
			if err := applyIncomeUserToken(cmd, body, userToken); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "item-id", itemIDs, "options", "item_ids"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "webhook", webhook, "options", "webhook"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/credit/payroll_income/refresh", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&userToken, "user-token", "", "Legacy Plaid user_token required by the current Income docs")
	cmd.Flags().StringSliceVar(&itemIDs, "item-id", nil, "Payroll item_id to refresh (repeatable)")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Webhook URL to receive payroll refresh events")
	return cmd
}
