package cmd

import (
	"context"
	"fmt"
	"strings"

	"plaid-cli/internal/plaid"
	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

func saveItemFromAccessToken(ctx context.Context, cmd *cobra.Command, store *state.Store, client *plaid.Client, accessToken, linkToken string, products, countryCodes []string) (state.ItemRecord, error) {
	itemResp, err := client.GetItem(ctx, accessToken)
	if err != nil {
		return state.ItemRecord{}, err
	}

	accountsResp, err := client.GetAccounts(ctx, accessToken)
	if err != nil {
		return state.ItemRecord{}, err
	}

	institutionName := ""
	if strings.TrimSpace(itemResp.Item.InstitutionID) != "" {
		institutionName, err = client.GetInstitutionName(ctx, itemResp.Item.InstitutionID, countryCodes)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Warning: could not resolve institution name: %v\n", err)
		}
	}

	record := state.ItemRecord{
		ItemID:          itemResp.Item.ItemID,
		AccessToken:     accessToken,
		InstitutionID:   itemResp.Item.InstitutionID,
		InstitutionName: institutionName,
		LinkToken:       linkToken,
		Products:        products,
		Accounts:        state.AccountSummariesFromPlaid(accountsResp.Accounts),
	}
	if err := store.SaveItem(record); err != nil {
		return state.ItemRecord{}, err
	}
	return record, nil
}
