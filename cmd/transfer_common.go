package cmd

import (
	"errors"
	"fmt"
	"os"

	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

const (
	transferInitiatingDocPath     = "docs/plaid/api/products/transfer/initiating-transfers/index.md"
	transferReadingDocPath        = "docs/plaid/api/products/transfer/reading-transfers/index.md"
	transferRefundsDocPath        = "docs/plaid/api/products/transfer/refunds/index.md"
	transferAccountLinkingDocPath = "docs/plaid/api/products/transfer/account-linking/index.md"
	transferRecurringDocPath      = "docs/plaid/api/products/transfer/recurring-transfers/index.md"
	transferMetricsDocPath        = "docs/plaid/api/products/transfer/metrics/index.md"
)

func populateAccessToken(cmd *cobra.Command, store *state.Store, body map[string]any, itemID, accessToken string) (*state.ItemRecord, error) {
	var record *state.ItemRecord

	if itemID != "" || accessToken != "" || !bodyHasValue(body, "access_token") {
		token, resolvedRecord, err := resolveAccessToken(cmd, store, itemID, accessToken)
		if err != nil {
			return nil, err
		}
		record = resolvedRecord
		if err := setBodyValue(body, token, "access_token"); err != nil {
			return nil, err
		}
		return record, nil
	}

	rawToken, ok := bodyValue(body, "access_token")
	if !ok {
		return nil, nil
	}

	token, ok := rawToken.(string)
	if !ok || token == "" {
		return nil, fmt.Errorf("request body field access_token must be a non-empty string")
	}

	resolvedRecord, err := store.FindItemByAccessToken(token)
	if err == nil {
		return resolvedRecord, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	return nil, err
}

func populateTransferAccess(cmd *cobra.Command, store *state.Store, body map[string]any, itemID, accessToken, accountID string) (*state.ItemRecord, error) {
	record, err := populateAccessToken(cmd, store, body, itemID, accessToken)
	if err != nil {
		return nil, err
	}

	if accountID != "" || !bodyHasValue(body, "account_id") {
		resolvedAccountID, err := resolveAccountID(record, accountID)
		if err != nil {
			return nil, err
		}
		if err := setBodyValue(body, resolvedAccountID, "account_id"); err != nil {
			return nil, err
		}
	}

	return record, nil
}

func requireBodyFields(body map[string]any, required map[string][]string) error {
	for label, path := range required {
		if !bodyHasValue(body, path...) {
			return fmt.Errorf("%s is required", label)
		}
	}
	return nil
}
