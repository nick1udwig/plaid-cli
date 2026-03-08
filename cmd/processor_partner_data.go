package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newProcessorTokenOnlyCmd(use, short, long, endpoint string) *cobra.Command {
	var processorToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"processor_token": "<processor-token>"}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token": {"processor_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, endpoint, body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	return cmd
}

func newProcessorAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Processor account data commands",
		Long:  "Read account data associated with a processor token.",
	}
	cmd.AddCommand(newProcessorTokenOnlyCmd(
		"get",
		"Call /processor/account/get",
		"Capability: read. Retrieves the account associated with a processor token.",
		"/processor/account/get",
	))
	return cmd
}

func newProcessorAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Processor Auth commands",
		Long:  "Read Auth data associated with a processor token.",
	}
	cmd.AddCommand(newProcessorTokenOnlyCmd(
		"get",
		"Call /processor/auth/get",
		"Capability: read. Retrieves Auth numbers and account metadata for a processor token.",
		"/processor/auth/get",
	))
	return cmd
}

func newProcessorBalanceCmd() *cobra.Command {
	var processorToken, minLastUpdatedDatetime string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "balance",
		Short: "Processor balance commands",
		Long:  "Read real-time balance data associated with a processor token.",
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Call /processor/balance/get",
		Long:  "Capability: read. Retrieves real-time balance data for a processor token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"processor_token": "<processor-token>"}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "min-last-updated-datetime", minLastUpdatedDatetime, "options", "min_last_updated_datetime"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token": {"processor_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/balance/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(getCmd)
	bodyFlags = bindBodyFlag(getCmd)
	getCmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	getCmd.Flags().StringVar(&minLastUpdatedDatetime, "min-last-updated-datetime", "", "Minimum acceptable balance timestamp in RFC3339 format")
	cmd.AddCommand(getCmd)
	return cmd
}

func newProcessorIdentityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identity",
		Short: "Processor identity commands",
		Long:  "Read identity and identity-match data associated with a processor token.",
	}
	cmd.AddCommand(newProcessorTokenOnlyCmd(
		"get",
		"Call /processor/identity/get",
		"Capability: read. Retrieves identity data for a processor token.",
		"/processor/identity/get",
	))
	cmd.AddCommand(newProcessorIdentityMatchCmd())
	return cmd
}

func newProcessorIdentityMatchCmd() *cobra.Command {
	var processorToken, legalName, emailAddress, phoneNumber string
	var city, region, street, postalCode, country string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "match",
		Short: "Call /processor/identity/match",
		Long:  "Capability: read. Compares provided identity inputs against processor-token identity data.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token": "<processor-token>",
				"user": map[string]any{
					"legal_name": "Jane Doe",
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
				return err
			}
			if legalName == "" && emailAddress == "" && phoneNumber == "" && street == "" && city == "" && region == "" && postalCode == "" && country == "" {
				return errors.New("provide at least one identity input flag")
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "legal-name", legalName, "user", "legal_name"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "email-address", emailAddress, "user", "email_address"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "phone-number", phoneNumber, "user", "phone_number"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "city", city, "user", "address", "city"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "region", region, "user", "address", "region"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "street", street, "user", "address", "street"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "postal-code", postalCode, "user", "address", "postal_code"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "country", country, "user", "address", "country"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token": {"processor_token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/identity/match", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringVar(&legalName, "legal-name", "", "Legal name to match")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "Email address to match")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Phone number to match")
	cmd.Flags().StringVar(&city, "city", "", "Address city to match")
	cmd.Flags().StringVar(&region, "region", "", "Address region to match")
	cmd.Flags().StringVar(&street, "street", "", "Address street to match")
	cmd.Flags().StringVar(&postalCode, "postal-code", "", "Address postal code to match")
	cmd.Flags().StringVar(&country, "country", "", "Address country to match")
	return cmd
}

func newProcessorLiabilitiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liabilities",
		Short: "Processor liabilities commands",
		Long:  "Read liabilities data associated with a processor token.",
	}
	cmd.AddCommand(newProcessorTokenOnlyCmd(
		"get",
		"Call /processor/liabilities/get",
		"Capability: read. Retrieves liabilities data for a processor token.",
		"/processor/liabilities/get",
	))
	return cmd
}

func newProcessorInvestmentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "investments",
		Short: "Processor investments commands",
		Long:  "Read investments data associated with a processor token.",
	}

	cmd.AddCommand(newProcessorTokenOnlyCmd(
		"holdings-get",
		"Call /processor/investments/holdings/get",
		"Capability: read. Retrieves investment holdings for a processor token.",
		"/processor/investments/holdings/get",
	))
	cmd.AddCommand(newProcessorInvestmentsTransactionsGetCmd())

	return cmd
}

func newProcessorInvestmentsTransactionsGetCmd() *cobra.Command {
	var processorToken, startDate, endDate string
	var accountIDs []string
	var count, offset int
	var asyncUpdate bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "transactions-get",
		Short: "Call /processor/investments/transactions/get",
		Long:  "Capability: read. Retrieves investment transactions for a processor token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token": "<processor-token>",
				"start_date":      "2026-01-01",
				"end_date":        "2026-02-01",
			}
			if handled, err := maybeWriteInfo(cmd, info, processorPartnersDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor-token", processorToken, "processor_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "start-date", startDate, "start_date"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-date", endDate, "end_date"); err != nil {
				return err
			}
			if err := applyStringSliceFlag(cmd, body, "account-id", accountIDs, "options", "account_ids"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "options", "count"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "offset", offset, "options", "offset"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "async-update", asyncUpdate, "options", "async_update"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token": {"processor_token"},
				"--start-date":      {"start_date"},
				"--end-date":        {"end_date"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/investments/transactions/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Start date in YYYY-MM-DD")
	cmd.Flags().StringVar(&endDate, "end-date", "", "End date in YYYY-MM-DD")
	cmd.Flags().StringSliceVar(&accountIDs, "account-id", nil, "Account IDs to filter by (repeatable)")
	cmd.Flags().IntVar(&count, "count", 100, "Maximum number of transactions to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Investment transactions offset")
	cmd.Flags().BoolVar(&asyncUpdate, "async-update", false, "Whether to enable async investment transaction refresh")
	return cmd
}
