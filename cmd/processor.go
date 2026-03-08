package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const processorDocPath = "docs/plaid/api/processors/index.md"

func newProcessorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "processor",
		Short: "Processor token commands",
		Long:  "Commands for creating processor tokens and managing processor token permissions.",
	}

	cmd.AddCommand(newProcessorTokenCreateCmd())
	cmd.AddCommand(newProcessorStripeBankAccountTokenCreateCmd())
	cmd.AddCommand(newProcessorTokenPermissionsGetCmd())
	cmd.AddCommand(newProcessorTokenPermissionsSetCmd())

	return cmd
}

func newProcessorTokenCreateCmd() *cobra.Command {
	var itemID, accessToken, accountID, processor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "token-create",
		Short: "Call /processor/token/create",
		Long:  "Capability: write. Creates a processor token for a specific account on a linked Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"account_id":   "<account-id>",
				"processor":    "dwolla",
			}
			if handled, err := maybeWriteInfo(cmd, info, processorDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "processor", processor, "processor"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--access-token": {"access_token"},
				"--account-id":   {"account_id"},
				"--processor":    {"processor"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/token/create", body)
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
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id for the processor token")
	cmd.Flags().StringVar(&processor, "processor", "", "Processor integration name, e.g. dwolla, modern_treasury, or unit")
	return cmd
}

func newProcessorStripeBankAccountTokenCreateCmd() *cobra.Command {
	var itemID, accessToken, accountID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "stripe-bank-account-token-create",
		Short: "Call /processor/stripe/bank_account_token/create",
		Long:  "Capability: write. Creates a one-time Stripe bank account token for a linked account.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"account_id":   "<account-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, processorDocPath, template); handled || err != nil {
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

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/stripe/bank_account_token/create", body)
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
	cmd.Flags().StringVar(&accountID, "account-id", "", "Plaid account_id for the Stripe token")
	return cmd
}

func newProcessorTokenPermissionsGetCmd() *cobra.Command {
	var processorToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "token-permissions-get",
		Short: "Call /processor/token/permissions/get",
		Long:  "Capability: admin. Retrieves the product permission set on a processor token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token": "<processor-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, processorDocPath, template); handled || err != nil {
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
			resp, err := client.Call(ctx, "/processor/token/permissions/get", body)
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

func newProcessorTokenPermissionsSetCmd() *cobra.Command {
	var processorToken string
	var products []string
	var allowAll bool
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "token-permissions-set",
		Short: "Call /processor/token/permissions/set",
		Long:  "Capability: admin. Sets the product permission list for a processor token. Use --allow-all to restore access to all available products.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"processor_token": "<processor-token>",
				"products":        []string{"auth", "balance", "identity"},
			}
			if handled, err := maybeWriteInfo(cmd, info, processorDocPath, template); handled || err != nil {
				return err
			}
			if allowAll && len(products) > 0 {
				return errors.New("--allow-all cannot be combined with --product")
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
			if allowAll {
				if err := setBodyValue(body, []string{}, "products"); err != nil {
					return err
				}
			} else {
				if err := applyStringSliceFlag(cmd, body, "product", products, "products"); err != nil {
					return err
				}
			}
			if err := requireBodyFields(body, map[string][]string{
				"--processor-token": {"processor_token"},
				"--product":         {"products"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/processor/token/permissions/set", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&processorToken, "processor-token", "", "Plaid processor_token")
	cmd.Flags().StringSliceVar(&products, "product", nil, "Product the processor token should have access to (repeatable)")
	cmd.Flags().BoolVar(&allowAll, "allow-all", false, "Set an empty products list, restoring access to all available products")
	return cmd
}
