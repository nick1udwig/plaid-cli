package cmd

import "github.com/spf13/cobra"

const consentDocPath = "docs/plaid/api/consent/index.md"

func newConsentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "consent",
		Short: "Consent history commands",
		Long:  "Read-only commands for Plaid Item consent history.",
	}

	cmd.AddCommand(newConsentEventsGetCmd())

	return cmd
}

func newConsentEventsGetCmd() *cobra.Command {
	var itemID, accessToken string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "events-get",
		Short: "Call /consent/events/get",
		Long:  "Capability: read. Retrieves the historical consent event log for an Item.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"access_token": "<access-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, consentDocPath, template); handled || err != nil {
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
			if _, err := populateAccessToken(cmd, store, body, itemID, accessToken); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/consent/events/get", body)
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
	return cmd
}
