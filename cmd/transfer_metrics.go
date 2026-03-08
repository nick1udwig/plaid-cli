package cmd

import "github.com/spf13/cobra"

func newTransferMetricsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metrics",
		Short: "Transfer metrics commands",
		Long:  "Read-only commands for transfer product metrics and configuration.",
	}

	cmd.AddCommand(newTransferMetricsGetCmd())
	cmd.AddCommand(newTransferConfigurationGetCmd())

	return cmd
}

func newTransferMetricsGetCmd() *cobra.Command {
	var originatorClientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/metrics/get",
		Long:  "Capability: read. Retrieves transfer product usage metrics.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{}
			if originatorClientID != "" {
				template["originator_client_id"] = originatorClientID
			}
			if handled, err := maybeWriteInfo(cmd, info, transferMetricsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/metrics/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID for platform use cases")
	return cmd
}

func newTransferConfigurationGetCmd() *cobra.Command {
	var originatorClientID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "configuration-get",
		Short: "Call /transfer/configuration/get",
		Long:  "Capability: read. Retrieves transfer product configuration and limits.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{}
			if originatorClientID != "" {
				template["originator_client_id"] = originatorClientID
			}
			if handled, err := maybeWriteInfo(cmd, info, transferMetricsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "originator-client-id", originatorClientID, "originator_client_id"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/configuration/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&originatorClientID, "originator-client-id", "", "Originator client ID for platform use cases")
	return cmd
}
