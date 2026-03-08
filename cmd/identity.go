package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newIdentityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identity",
		Short: "Identity product commands",
		Long:  "Read-only identity commands.",
	}
	cmd.AddCommand(newIdentityGetCmd())
	cmd.AddCommand(newIdentityMatchCmd())
	return cmd
}

func newIdentityGetCmd() *cobra.Command {
	var itemID, accessToken string
	var accountIDs []string
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /identity/get",
		Long:  "Capability: read. Retrieves identity data returned by the financial institution.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"access_token": "<access-token>"}
			if len(accountIDs) > 0 {
				template["options"] = map[string]any{"account_ids": accountIDs}
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/identity/index.md", template); handled || err != nil {
				return err
			}

			store, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			token, _, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			body := map[string]any{"access_token": token}
			if len(accountIDs) > 0 {
				body["options"] = map[string]any{"account_ids": accountIDs}
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/identity/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringSliceVar(&accountIDs, "account-id", nil, "Account ID to filter by (repeatable)")
	return cmd
}

func newIdentityMatchCmd() *cobra.Command {
	var itemID, accessToken, legalName, emailAddress, phoneNumber string
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "match",
		Short: "Call /identity/match",
		Long:  "Capability: read. Compares provided identity inputs against the institution-provided identity data.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
				"user": map[string]any{
					"legal_name": "Jane Doe",
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/products/identity/index.md", template); handled || err != nil {
				return err
			}
			if legalName == "" && emailAddress == "" && phoneNumber == "" {
				return errors.New("provide at least one of --legal-name, --email-address, or --phone-number")
			}

			store, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			token, _, err := resolveAccessToken(cmd, store, itemID, accessToken)
			if err != nil {
				return err
			}

			user := map[string]any{}
			if legalName != "" {
				user["legal_name"] = legalName
			}
			if emailAddress != "" {
				user["email_address"] = emailAddress
			}
			if phoneNumber != "" {
				user["phone_number"] = phoneNumber
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/identity/match", map[string]any{
				"access_token": token,
				"user":         user,
			})
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&itemID, "item", "", "Saved local item_id to use")
	cmd.Flags().StringVar(&accessToken, "access-token", "", "Explicit Plaid access_token override")
	cmd.Flags().StringVar(&legalName, "legal-name", "", "Legal name to match against bank identity")
	cmd.Flags().StringVar(&emailAddress, "email-address", "", "Email address to match against bank identity")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Phone number to match against bank identity")
	return cmd
}
