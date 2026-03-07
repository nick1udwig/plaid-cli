package cmd

import (
	"encoding/json"
	"fmt"

	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

func getStateDir(cmd *cobra.Command) (string, error) {
	return cmd.Root().PersistentFlags().GetString("state-dir")
}

func getStore(cmd *cobra.Command) (*state.Store, error) {
	stateDir, err := getStateDir(cmd)
	if err != nil {
		return nil, err
	}
	return state.New(stateDir), nil
}

func writeJSON(cmd *cobra.Command, value any) error {
	encoder := json.NewEncoder(cmd.OutOrStdout())
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}
