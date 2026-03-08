package cmd

import "github.com/spf13/cobra"

func newTransferRecurringCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recurring",
		Short: "Recurring Transfer commands",
		Long:  "Recurring Transfer commands for creating, cancelling, and inspecting recurring money movement schedules.",
	}

	cmd.AddCommand(newTransferRecurringCreateCmd())
	cmd.AddCommand(newTransferRecurringCancelCmd())
	cmd.AddCommand(newTransferRecurringGetCmd())
	cmd.AddCommand(newTransferRecurringListCmd())

	return cmd
}

func newTransferRecurringCreateCmd() *cobra.Command {
	var itemID, accessToken, accountID string
	var idempotencyKey, transferType, network, achClass string
	var amount, description, testClockID string
	var legalName, phoneNumber, emailAddress string
	var ipAddress, userAgent string
	var intervalUnit, startDate, endDate string
	var intervalCount, intervalExecutionDay int
	var userPresent bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /transfer/recurring/create",
		Long:  "Capability: write. Creates a recurring transfer schedule and future transfer originations.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token":    "<access-token>",
				"account_id":      "<account-id>",
				"idempotency_key": "<idempotency-key>",
				"type":            "debit",
				"network":         "ach",
				"amount":          "12.34",
				"description":     "payment",
				"user": map[string]any{
					"legal_name": "Jane Doe",
				},
				"device": map[string]any{
					"ip_address": "203.0.113.10",
					"user_agent": "plaid-cli/0.1",
				},
				"schedule": map[string]any{
					"interval_unit":          "week",
					"interval_count":         1,
					"interval_execution_day": 5,
					"start_date":             "2026-03-15",
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, transferRecurringDocPath, template); handled || err != nil {
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
			if _, err := populateTransferAccess(cmd, store, body, itemID, accessToken, accountID); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "type", transferType, "type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "network", network, "network"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ach-class", achClass, "ach_class"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "user-present", userPresent, "user_present"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "description", description, "description"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "test-clock-id", testClockID, "test_clock_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "legal-name", legalName, "user", "legal_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "phone-number", phoneNumber, "user", "phone_number"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "email-address", emailAddress, "user", "email_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "ip-address", ipAddress, "device", "ip_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "user-agent", userAgent, "device", "user_agent"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "interval-unit", intervalUnit, "schedule", "interval_unit"); err != nil {
				return err
			}
			if cmd.Flags().Changed("interval-count") {
				if err := setBodyValue(body, intervalCount, "schedule", "interval_count"); err != nil {
					return err
				}
			}
			if cmd.Flags().Changed("interval-execution-day") {
				if err := setBodyValue(body, intervalExecutionDay, "schedule", "interval_execution_day"); err != nil {
					return err
				}
			}
			if err := applyStringFlag(cmd, body, "start-date", startDate, "schedule", "start_date"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-date", endDate, "schedule", "end_date"); err != nil {
				return err
			}

			if err := requireBodyFields(body, map[string][]string{
				"--access-token":           {"access_token"},
				"--account-id":             {"account_id"},
				"--idempotency-key":        {"idempotency_key"},
				"--type":                   {"type"},
				"--network":                {"network"},
				"--amount":                 {"amount"},
				"--description":            {"description"},
				"--legal-name":             {"user", "legal_name"},
				"--ip-address":             {"device", "ip_address"},
				"--user-agent":             {"device", "user_agent"},
				"--interval-unit":          {"schedule", "interval_unit"},
				"--interval-count":         {"schedule", "interval_count"},
				"--interval-execution-day": {"schedule", "interval_execution_day"},
				"--start-date":             {"schedule", "start_date"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/recurring/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id to debit or credit")
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency key for safely retrying recurring transfer creation")
	cmd.Flags().StringVar(&transferType, "type", "", "Recurring transfer type: debit or credit")
	cmd.Flags().StringVar(&network, "network", "", "Recurring transfer network: ach, same-day-ach, or rtp")
	cmd.Flags().StringVar(&achClass, "ach-class", "", "ACH SEC code for ACH recurring transfers")
	cmd.Flags().StringVar(&amount, "amount", "", "Recurring transfer amount as a decimal string")
	cmd.Flags().BoolVar(&userPresent, "user-present", false, "Whether the end user is actively initiating this recurring transfer")
	cmd.Flags().StringVar(&description, "description", "", "Recurring transfer description")
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Sandbox test clock ID to associate with the recurring transfer")
	cmd.Flags().StringVar(&legalName, "legal-name", "", "Account holder legal name")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Account holder phone number")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "Account holder email address")
	cmd.Flags().StringVar(&ipAddress, "ip-address", "", "Device IP address for the recurring transfer request")
	cmd.Flags().StringVar(&userAgent, "user-agent", "", "Device user agent for the recurring transfer request")
	cmd.Flags().StringVar(&intervalUnit, "interval-unit", "", "Schedule interval unit: week or month")
	cmd.Flags().IntVar(&intervalCount, "interval-count", 0, "Number of interval units between originations")
	cmd.Flags().IntVar(&intervalExecutionDay, "interval-execution-day", 0, "Day within the interval to execute the recurring transfer")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Schedule start date in YYYY-MM-DD")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Optional schedule end date in YYYY-MM-DD")
	return cmd
}

func newTransferRecurringCancelCmd() *cobra.Command {
	var recurringTransferID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "Call /transfer/recurring/cancel",
		Long:  "Capability: write. Cancels a recurring transfer and any future unsent originations.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"recurring_transfer_id": "<recurring-transfer-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferRecurringDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "recurring-transfer-id", recurringTransferID, "recurring_transfer_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--recurring-transfer-id": {"recurring_transfer_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/recurring/cancel", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&recurringTransferID, "recurring-transfer-id", "", "Recurring transfer ID to cancel")
	return cmd
}

func newTransferRecurringGetCmd() *cobra.Command {
	var recurringTransferID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/recurring/get",
		Long:  "Capability: read. Retrieves a recurring transfer by recurring_transfer_id.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"recurring_transfer_id": "<recurring-transfer-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferRecurringDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "recurring-transfer-id", recurringTransferID, "recurring_transfer_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--recurring-transfer-id": {"recurring_transfer_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/recurring/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&recurringTransferID, "recurring-transfer-id", "", "Recurring transfer ID to retrieve")
	return cmd
}

func newTransferRecurringListCmd() *cobra.Command {
	var startTime, endTime, fundingAccountID string
	var count, offset int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /transfer/recurring/list",
		Long:  "Capability: read. Lists recurring transfers and their statuses.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"count":  count,
				"offset": offset,
			}
			if startTime != "" {
				template["start_time"] = startTime
			}
			if endTime != "" {
				template["end_time"] = endTime
			}
			if handled, err := maybeWriteInfo(cmd, info, transferRecurringDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "start-time", startTime, "start_time"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-time", endTime, "end_time"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "offset", offset, "offset"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "funding-account-id", fundingAccountID, "funding_account_id"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/recurring/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&startTime, "start-time", "", "RFC3339 start created timestamp")
	cmd.Flags().StringVar(&endTime, "end-time", "", "RFC3339 end created timestamp")
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of recurring transfers to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Recurring transfer list offset")
	cmd.Flags().StringVar(&fundingAccountID, "funding-account-id", "", "Funding account ID filter")
	return cmd
}
