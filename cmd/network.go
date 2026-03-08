package cmd

import "github.com/spf13/cobra"

const networkDocPath = "docs/plaid/api/network/index.md"

func newNetworkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network",
		Short: "Plaid Network commands",
		Long:  "Read-only commands for checking Plaid Network and Layer-related eligibility status.",
	}

	cmd.AddCommand(newNetworkStatusGetCmd())

	return cmd
}

func newNetworkStatusGetCmd() *cobra.Command {
	var phoneNumber, templateID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "status-get",
		Short: "Call /network/status/get",
		Long:  "Capability: read. Checks whether Plaid has a matching network profile for a user and whether Layer metadata is available.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"user": map[string]any{
					"phone_number": "+14155550015",
				},
			}
			if handled, err := maybeWriteInfo(cmd, info, networkDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "phone-number", phoneNumber, "user", "phone_number"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "template-id", templateID, "template_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--phone-number": {"user", "phone_number"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/network/status/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Phone number in E.164 format")
	cmd.Flags().StringVar(&templateID, "template-id", "", "Optional Plaid Dashboard template_id")
	return cmd
}
