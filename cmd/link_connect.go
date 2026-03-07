package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"plaid-cli/internal/plaid"
	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

func newLinkConnectCmd(opts *Options) *cobra.Command {
	var (
		products      []string
		countryCodes  []string
		clientUserID  string
		webhook       string
		redirectURI   string
		skipOpen      bool
		timeout       time.Duration
		pollInterval  time.Duration
		language      string
		printDocPath  bool
		printTemplate bool
	)

	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Open Hosted Link in the browser and save the resulting Item",
		Long:  "Create a Hosted Link session, open it in the browser, poll for completion, exchange the public token, and save the Item locally.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if printDocPath {
				return writeJSON(cmd, map[string]any{
					"doc_path": "docs/plaid/link/hosted-link/index.md",
				})
			}

			store, err := getStore(cmd)
			if err != nil {
				return err
			}

			profile, profileErr := store.LoadAppProfile()
			if printTemplate {
				clientName := "plaid-cli"
				templateLanguage := language
				templateCountries := countryCodes
				if profileErr == nil {
					if profile.ClientName != "" {
						clientName = profile.ClientName
					}
					if templateLanguage == "" {
						templateLanguage = profile.Language
					}
					if len(templateCountries) == 0 {
						templateCountries = profile.CountryCodes
					}
				}
				if templateLanguage == "" {
					templateLanguage = "en"
				}
				if len(templateCountries) == 0 {
					templateCountries = []string{"US"}
				}
				if len(products) == 0 {
					products = []string{"auth"}
				}
				if clientUserID == "" {
					clientUserID = "local-owner"
				}
				return writeJSON(cmd, map[string]any{
					"client_name":   clientName,
					"language":      templateLanguage,
					"country_codes": templateCountries,
					"user": map[string]any{
						"client_user_id": clientUserID,
					},
					"products":    products,
					"hosted_link": map[string]any{},
				})
			}

			if profileErr != nil {
				return fmt.Errorf("load app profile: %w\nrun `plaid init` first", profileErr)
			}
			if len(products) == 0 {
				return errors.New("at least one --product is required")
			}

			if language == "" {
				language = profile.Language
			}
			if len(countryCodes) == 0 {
				countryCodes = profile.CountryCodes
			}
			if clientUserID == "" {
				clientUserID = "local-owner"
			}

			client, err := plaid.NewClient(profile)
			if err != nil {
				return err
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			linkResp, err := client.CreateHostedLinkToken(ctx, plaid.CreateHostedLinkTokenInput{
				ClientName:   profile.ClientName,
				Language:     language,
				CountryCodes: countryCodes,
				ClientUserID: clientUserID,
				Products:     products,
				Webhook:      webhook,
				RedirectURI:  redirectURI,
			})
			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.ErrOrStderr(), "Hosted Link URL: %s\n", linkResp.HostedLinkURL)
			if !skipOpen {
				if err := opts.BrowserOpener(linkResp.HostedLinkURL); err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Could not open browser automatically: %v\n", err)
				}
			}
			fmt.Fprintln(cmd.ErrOrStderr(), "Waiting for Link to finish...")

			publicToken, err := plaid.WaitForPublicToken(ctx, client, linkResp.LinkToken, pollInterval)
			if err != nil {
				return err
			}

			exchangeResp, err := client.ExchangePublicToken(ctx, publicToken)
			if err != nil {
				return err
			}

			itemResp, err := client.GetItem(ctx, exchangeResp.AccessToken)
			if err != nil {
				return err
			}

			accountsResp, err := client.GetAccounts(ctx, exchangeResp.AccessToken)
			if err != nil {
				return err
			}

			institutionName := ""
			if strings.TrimSpace(itemResp.Item.InstitutionID) != "" {
				institutionName, err = client.GetInstitutionName(ctx, itemResp.Item.InstitutionID, countryCodes)
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Warning: could not resolve institution name: %v\n", err)
				}
			}

			record := state.ItemRecord{
				ItemID:          exchangeResp.ItemID,
				AccessToken:     exchangeResp.AccessToken,
				InstitutionID:   itemResp.Item.InstitutionID,
				InstitutionName: institutionName,
				LinkToken:       linkResp.LinkToken,
				Products:        products,
				Accounts:        state.AccountSummariesFromPlaid(accountsResp.Accounts),
			}

			if err := store.SaveItem(record); err != nil {
				return err
			}

			stateDir, err := getStateDir(cmd)
			if err != nil {
				return err
			}

			return writeJSON(cmd, map[string]any{
				"ok":        true,
				"state_dir": stateDir,
				"item": map[string]any{
					"item_id":           record.ItemID,
					"institution_id":    record.InstitutionID,
					"institution_name":  record.InstitutionName,
					"products":          record.Products,
					"accounts":          record.Accounts,
					"saved_item_record": store.ItemPath(record.ItemID),
				},
			})
		},
	}

	cmd.Flags().StringSliceVar(&products, "product", nil, "Plaid product to initialize in Link (repeatable)")
	cmd.Flags().StringSliceVar(&countryCodes, "country-code", nil, "Country code for Link (defaults to init profile)")
	cmd.Flags().StringVar(&clientUserID, "client-user-id", "", "Stable client_user_id for the single local owner")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Optional webhook for Hosted Link session events")
	cmd.Flags().StringVar(&redirectURI, "redirect-uri", "", "Optional redirect_uri for OAuth-capable flows")
	cmd.Flags().BoolVar(&skipOpen, "no-open", false, "Do not try to open the Hosted Link URL automatically")
	cmd.Flags().DurationVar(&timeout, "timeout", 10*time.Minute, "How long to wait for Link completion")
	cmd.Flags().DurationVar(&pollInterval, "poll-interval", 2*time.Second, "How often to poll /link/token/get")
	cmd.Flags().StringVar(&language, "language", "", "Override the Link language from the saved app profile")
	cmd.Flags().BoolVar(&printDocPath, "print-doc-path", false, "Print the local docs path backing this command and exit")
	cmd.Flags().BoolVar(&printTemplate, "print-request-template", false, "Print a minimal Link token request template and exit")

	return cmd
}
