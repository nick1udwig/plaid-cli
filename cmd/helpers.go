package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
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

type requestBodyFlags struct {
	body string
}

func bindInfoFlags(cmd *cobra.Command) *commandInfoFlags {
	flags := &commandInfoFlags{}
	cmd.Flags().BoolVar(&flags.printDocPath, "print-doc-path", false, "Print the local docs path backing this command and exit")
	cmd.Flags().BoolVar(&flags.printRequestTemplate, "print-request-template", false, "Print a minimal request template and exit")
	return flags
}

func bindBodyFlag(cmd *cobra.Command) *requestBodyFlags {
	flags := &requestBodyFlags{}
	cmd.Flags().StringVar(&flags.body, "body", "", "Base JSON request body as inline JSON or @path/to/file.json; explicit flags override it")
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

func resolveAccountID(record *state.ItemRecord, accountID string) (string, error) {
	if accountID != "" {
		return accountID, nil
	}
	if record == nil {
		return "", errors.New("--account-id is required when no saved item record is available")
	}

	switch len(record.Accounts) {
	case 0:
		return "", fmt.Errorf("--account-id is required; no saved accounts found for item %s", record.ItemID)
	case 1:
		return record.Accounts[0].AccountID, nil
	default:
		ids := make([]string, 0, len(record.Accounts))
		for _, account := range record.Accounts {
			ids = append(ids, account.AccountID)
		}
		return "", fmt.Errorf("multiple saved accounts found for item %s; provide --account-id. available account_ids: %v", record.ItemID, ids)
	}
}

func loadRequestBody(raw string) (map[string]any, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return map[string]any{}, nil
	}

	source := "inline JSON"
	payload := raw
	if strings.HasPrefix(raw, "@") {
		path := strings.TrimSpace(strings.TrimPrefix(raw, "@"))
		if path == "" {
			return nil, errors.New("body path after @ must not be empty")
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read request body file %s: %w", path, err)
		}
		source = path
		payload = string(content)
	}

	var decoded any
	if err := json.Unmarshal([]byte(payload), &decoded); err != nil {
		return nil, fmt.Errorf("decode request body from %s: %w", source, err)
	}

	body, ok := decoded.(map[string]any)
	if !ok {
		return nil, errors.New("request body must be a JSON object")
	}
	return body, nil
}

func bodyValue(body map[string]any, path ...string) (any, bool) {
	if len(path) == 0 {
		return body, body != nil
	}

	current := any(body)
	for _, segment := range path {
		node, ok := current.(map[string]any)
		if !ok {
			return nil, false
		}
		value, ok := node[segment]
		if !ok {
			return nil, false
		}
		current = value
	}
	return current, true
}

func bodyHasValue(body map[string]any, path ...string) bool {
	_, ok := bodyValue(body, path...)
	return ok
}

func setBodyValue(body map[string]any, value any, path ...string) error {
	if len(path) == 0 {
		return errors.New("setBodyValue requires at least one path segment")
	}
	if body == nil {
		return errors.New("setBodyValue requires a non-nil body")
	}

	current := body
	for _, segment := range path[:len(path)-1] {
		next, ok := current[segment]
		if !ok {
			child := map[string]any{}
			current[segment] = child
			current = child
			continue
		}
		child, ok := next.(map[string]any)
		if !ok {
			return fmt.Errorf("request body field %q is not an object", strings.Join(path[:len(path)-1], "."))
		}
		current = child
	}

	current[path[len(path)-1]] = value
	return nil
}

func applyStringFlag(cmd *cobra.Command, body map[string]any, flagName, value string, path ...string) error {
	if !cmd.Flags().Changed(flagName) {
		if value == "" || bodyHasValue(body, path...) {
			return nil
		}
	}
	return setBodyValue(body, value, path...)
}

func applyIntFlag(cmd *cobra.Command, body map[string]any, flagName string, value int, path ...string) error {
	if !cmd.Flags().Changed(flagName) && bodyHasValue(body, path...) {
		return nil
	}
	return setBodyValue(body, value, path...)
}

func applyBoolFlag(cmd *cobra.Command, body map[string]any, flagName string, value bool, path ...string) error {
	if !cmd.Flags().Changed(flagName) {
		return nil
	}
	return setBodyValue(body, value, path...)
}

func applyStringSliceFlag(cmd *cobra.Command, body map[string]any, flagName string, value []string, path ...string) error {
	if !cmd.Flags().Changed(flagName) {
		if len(value) == 0 || bodyHasValue(body, path...) {
			return nil
		}
	}
	return setBodyValue(body, value, path...)
}

func applyStringMapFlag(cmd *cobra.Command, body map[string]any, flagName string, value map[string]string, path ...string) error {
	if !cmd.Flags().Changed(flagName) {
		if len(value) == 0 || bodyHasValue(body, path...) {
			return nil
		}
	}
	return setBodyValue(body, stringMapToAny(value), path...)
}

func stringMapToAny(value map[string]string) map[string]any {
	if len(value) == 0 {
		return map[string]any{}
	}

	out := make(map[string]any, len(value))
	for key, item := range value {
		out[key] = item
	}
	return out
}
