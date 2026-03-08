package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"plaid-cli/internal/plaid"
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

type commandInfoFlags struct {
	printDocPath         bool
	printRequestTemplate bool
}

func bindInfoFlags(cmd *cobra.Command) *commandInfoFlags {
	flags := &commandInfoFlags{}
	cmd.Flags().BoolVar(&flags.printDocPath, "print-doc-path", false, "Print the local docs path backing this command and exit")
	cmd.Flags().BoolVar(&flags.printRequestTemplate, "print-request-template", false, "Print a minimal request template and exit")
	return flags
}

func maybeWriteInfo(cmd *cobra.Command, flags *commandInfoFlags, docPath string, template any) (bool, error) {
	if flags == nil {
		return false, nil
	}
	if flags.printDocPath {
		return true, writeJSON(cmd, map[string]any{"doc_path": docPath})
	}
	if flags.printRequestTemplate {
		if template == nil {
			template = map[string]any{}
		}
		return true, writeJSON(cmd, template)
	}
	return false, nil
}

func loadClientFromState(cmd *cobra.Command) (*state.Store, state.AppProfile, *plaid.Client, error) {
	store, err := getStore(cmd)
	if err != nil {
		return nil, state.AppProfile{}, nil, err
	}
	profile, err := store.LoadAppProfile()
	if err != nil {
		return nil, state.AppProfile{}, nil, fmt.Errorf("load app profile: %w\nrun `plaid init` first", err)
	}
	client, err := plaid.NewClient(profile)
	if err != nil {
		return nil, state.AppProfile{}, nil, err
	}
	return store, profile, client, nil
}

func commandContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 45*time.Second)
}

func defaultCountryCodes(profile state.AppProfile, provided []string) []string {
	if len(provided) != 0 {
		return provided
	}
	if len(profile.CountryCodes) != 0 {
		return profile.CountryCodes
	}
	return []string{"US"}
}

func resolveAccessToken(cmd *cobra.Command, store *state.Store, itemID, accessToken string) (string, *state.ItemRecord, error) {
	if accessToken != "" {
		record, err := store.FindItemByAccessToken(accessToken)
		if err == nil {
			return accessToken, record, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			return "", nil, err
		}
		return accessToken, nil, nil
	}

	if itemID != "" {
		record, err := store.LoadItem(itemID)
		if err != nil {
			return "", nil, err
		}
		return record.AccessToken, &record, nil
	}

	items, err := store.ListItems()
	if err != nil {
		return "", nil, err
	}
	switch len(items) {
	case 0:
		return "", nil, errors.New("no saved items found; run `plaid link connect` or provide --access-token")
	case 1:
		record := items[0]
		return record.AccessToken, &record, nil
	default:
		ids := make([]string, 0, len(items))
		for _, item := range items {
			ids = append(ids, item.ItemID)
		}
		return "", nil, fmt.Errorf("multiple saved items found; provide --item. available item_ids: %v", ids)
	}
}
