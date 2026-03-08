package cmd

import "github.com/spf13/cobra"

func newSandboxTransferTestClockCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test-clock",
		Short: "Sandbox Transfer test clock helpers",
		Long:  "Sandbox-only helpers for creating, advancing, and inspecting Transfer test clocks.",
	}

	cmd.AddCommand(newSandboxTransferTestClockCreateCmd())
	cmd.AddCommand(newSandboxTransferTestClockAdvanceCmd())
	cmd.AddCommand(newSandboxTransferTestClockGetCmd())
	cmd.AddCommand(newSandboxTransferTestClockListCmd())

	return cmd
}

func newSandboxTransferTestClockCreateCmd() *cobra.Command {
	var virtualTime string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /sandbox/transfer/test_clock/create",
		Long:  "Capability: sandbox-write. Creates a Transfer sandbox test clock.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "virtual-time", virtualTime, "virtual_time"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/test_clock/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&virtualTime, "virtual-time", "", "Initial RFC3339 virtual time for the test clock")
	return cmd
}

func newSandboxTransferTestClockAdvanceCmd() *cobra.Command {
	var testClockID, newVirtualTime string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "advance",
		Short: "Call /sandbox/transfer/test_clock/advance",
		Long:  "Capability: sandbox-write. Advances a Transfer sandbox test clock forward in time.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"test_clock_id":    "<test-clock-id>",
				"new_virtual_time": "2026-03-08T00:00:00Z",
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "test-clock-id", testClockID, "test_clock_id"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "new-virtual-time", newVirtualTime, "new_virtual_time"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--test-clock-id":    {"test_clock_id"},
				"--new-virtual-time": {"new_virtual_time"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/test_clock/advance", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Transfer test clock ID to advance")
	cmd.Flags().StringVar(&newVirtualTime, "new-virtual-time", "", "New RFC3339 virtual time")
	return cmd
}

func newSandboxTransferTestClockGetCmd() *cobra.Command {
	var testClockID string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /sandbox/transfer/test_clock/get",
		Long:  "Capability: sandbox-read. Retrieves a Transfer sandbox test clock by ID.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"test_clock_id": "<test-clock-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "test-clock-id", testClockID, "test_clock_id"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--test-clock-id": {"test_clock_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/test_clock/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&testClockID, "test-clock-id", "", "Transfer test clock ID to retrieve")
	return cmd
}

func newSandboxTransferTestClockListCmd() *cobra.Command {
	var startVirtualTime, endVirtualTime string
	var count, offset int
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /sandbox/transfer/test_clock/list",
		Long:  "Capability: sandbox-read. Lists Transfer sandbox test clocks.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"count":  count,
				"offset": offset,
			}
			if handled, err := maybeWriteInfo(cmd, info, sandboxDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "start-virtual-time", startVirtualTime, "start_virtual_time"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "end-virtual-time", endVirtualTime, "end_virtual_time"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "count", count, "count"); err != nil {
				return err
			}
			if err := applyIntFlag(cmd, body, "offset", offset, "offset"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/sandbox/transfer/test_clock/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&startVirtualTime, "start-virtual-time", "", "RFC3339 start virtual time")
	cmd.Flags().StringVar(&endVirtualTime, "end-virtual-time", "", "RFC3339 end virtual time")
	cmd.Flags().IntVar(&count, "count", 25, "Maximum number of test clocks to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Test clock list offset")
	return cmd
}
