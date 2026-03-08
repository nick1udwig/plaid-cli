package cmd

import "github.com/spf13/cobra"

const enrichDocPath = "docs/plaid/api/products/enrich/index.md"

func newEnrichCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enrich",
		Short: "Enrich product commands",
		Long:  "Enrich commands for classifying locally held transaction data that did not originate from Plaid.",
	}

	cmd.AddCommand(newEnrichTransactionsCmd())

	return cmd
}

func newEnrichTransactionsCmd() *cobra.Command {
	var accountType, pfcVersion string
	var includeLegacyCategory bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Call /transactions/enrich",
		Long:  "Capability: read. Enriches locally held transaction data. Use `--body` to provide the transaction array and any structured location fields.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"account_type": "depository",
				"transactions": []map[string]any{
					{
						"id":                "txn-1",
						"description":       "PURCHASE WM SUPERCENTER #1700",
						"amount":            72.10,
						"direction":         "OUTFLOW",
						"iso_currency_code": "USD",
					},
				},
			}
			options := map[string]any{}
			if includeLegacyCategory {
				options["include_legacy_category"] = true
			}
			if pfcVersion != "" {
				options["personal_finance_category_version"] = pfcVersion
			}
			if len(options) > 0 {
				template["options"] = options
			}
			if handled, err := maybeWriteInfo(cmd, info, enrichDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "account-type", accountType, "account_type"); err != nil {
				return err
			}
			if err := applyBoolFlag(cmd, body, "include-legacy-category", includeLegacyCategory, "options", "include_legacy_category"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "pfc-version", pfcVersion, "options", "personal_finance_category_version"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--account-type": {"account_type"},
				"--body":         {"transactions"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transactions/enrich", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&accountType, "account-type", "", "Account type for the provided transactions, e.g. depository or credit")
	cmd.Flags().BoolVar(&includeLegacyCategory, "include-legacy-category", false, "Include legacy category fields in the response")
	cmd.Flags().StringVar(&pfcVersion, "pfc-version", "", "Personal finance category taxonomy version, e.g. v1 or v2")
	return cmd
}
