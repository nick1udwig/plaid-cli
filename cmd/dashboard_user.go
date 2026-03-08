package cmd

import "github.com/spf13/cobra"

const dashboardUserDocPath = "docs/plaid/api/kyc-aml-users/index.md"

func newDashboardUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboard-user",
		Short: "Dashboard user lookup commands",
		Long:  "Dashboard user lookup commands used by Monitor, Beacon, and Identity Verification audit trails.",
	}

	cmd.AddCommand(newDashboardUserGetCmd())
	cmd.AddCommand(newDashboardUserListCmd())

	return cmd
}

func newDashboardUserGetCmd() *cobra.Command {
	var dashboardUserID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /dashboard_user/get",
		Long:  "Capability: read. Retrieves a single Plaid Dashboard user referenced in audit trails.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"dashboard_user_id": "<dashboard-user-id>"}
			if handled, err := maybeWriteInfo(cmd, info, dashboardUserDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "dashboard-user-id", dashboardUserID, "dashboard_user_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--dashboard-user-id": {"dashboard_user_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/dashboard_user/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&dashboardUserID, "dashboard-user-id", "", "Plaid Dashboard user identifier from an audit_trail object")
	return cmd
}

func newDashboardUserListCmd() *cobra.Command {
	var cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /dashboard_user/list",
		Long:  "Capability: read. Lists Plaid Dashboard users associated with your account.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, dashboardUserDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/dashboard_user/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor from a previous dashboard-user list call")
	return cmd
}
