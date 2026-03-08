package cmd

import "github.com/spf13/cobra"

func newTransferRefundCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refund",
		Short: "Transfer refund commands",
		Long:  "Refund commands. `create` and `cancel` are write operations; `get` is read-only.",
	}

	cmd.AddCommand(newTransferRefundCreateCmd())
	cmd.AddCommand(newTransferRefundGetCmd())
	cmd.AddCommand(newTransferRefundCancelCmd())

	return cmd
}

func newTransferRefundCreateCmd() *cobra.Command {
	var transferID, amount, idempotencyKey string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /transfer/refund/create",
		Long:  "Capabilities: write, move-money. Creates a refund for a transfer.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"transfer_id":     "<transfer-id>",
				"amount":          "12.34",
				"idempotency_key": "<idempotency-key>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferRefundsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "transfer-id", transferID, "transfer_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "amount", amount, "amount"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "idempotency-key", idempotencyKey, "idempotency_key"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--transfer-id":     {"transfer_id"},
				"--amount":          {"amount"},
				"--idempotency-key": {"idempotency_key"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/refund/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&transferID, "transfer-id", "", "Transfer ID to refund")
	cmd.Flags().StringVar(&amount, "amount", "", "Refund amount as a decimal string")
	cmd.Flags().StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency key for safely retrying the refund")
	return cmd
}

func newTransferRefundGetCmd() *cobra.Command {
	var refundID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /transfer/refund/get",
		Long:  "Capability: read. Retrieves a refund by refund_id.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"refund_id": "<refund-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferRefundsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "refund-id", refundID, "refund_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--refund-id": {"refund_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/refund/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&refundID, "refund-id", "", "Refund ID to retrieve")
	return cmd
}

func newTransferRefundCancelCmd() *cobra.Command {
	var refundID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "Call /transfer/refund/cancel",
		Long:  "Capabilities: write, move-money. Cancels a refund that has not yet been submitted to the payment network.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"refund_id": "<refund-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, transferRefundsDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "refund-id", refundID, "refund_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--refund-id": {"refund_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/transfer/refund/cancel", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&refundID, "refund-id", "", "Refund ID to cancel")
	return cmd
}
