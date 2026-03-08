package cmd

import "github.com/spf13/cobra"

func bindFailureReasonFlags(cmd *cobra.Command, failureCode, achReturnCode, description *string) {
	cmd.Flags().StringVar(failureCode, "failure-code", "", "Failure code for failed or returned sandbox events")
	cmd.Flags().StringVar(achReturnCode, "ach-return-code", "", "Deprecated ACH return code for failed or returned sandbox events")
	cmd.Flags().StringVar(description, "failure-description", "", "Human-readable failure description for failed or returned sandbox events")
}

func applyFailureReasonFlags(cmd *cobra.Command, body map[string]any, failureCode, achReturnCode, description string) error {
	if err := applyStringFlag(cmd, body, "failure-code", failureCode, "failure_reason", "failure_code"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "ach-return-code", achReturnCode, "failure_reason", "ach_return_code"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "failure-description", description, "failure_reason", "description"); err != nil {
		return err
	}
	return nil
}
