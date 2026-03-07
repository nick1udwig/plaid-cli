package cmd

import (
	"errors"
	"fmt"
	"strings"

	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

func newInitCmd(_ *Options) *cobra.Command {
	var (
		env          string
		clientID     string
		secret       string
		clientName   string
		language     string
		countryCodes []string
	)

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Save Plaid app credentials and defaults to ~/.plaid-cli",
		Long: fmt.Sprintf(
			"Store the one-time Plaid app configuration required for this CLI.\n\n"+
				"The human operator must obtain the required values from the Plaid Dashboard.\n"+
				"See %s for the exact setup steps.",
			setupDocPath,
		),
		RunE: func(cmd *cobra.Command, _ []string) error {
			store, err := getStore(cmd)
			if err != nil {
				return err
			}

			if strings.TrimSpace(env) == "" {
				env = strings.TrimSpace(state.GetenvAny("PLAID_ENV"))
			}
			if strings.TrimSpace(clientID) == "" {
				clientID = strings.TrimSpace(state.GetenvAny("PLAID_CLIENT_ID"))
			}
			if strings.TrimSpace(secret) == "" {
				secret = strings.TrimSpace(state.GetenvAny("PLAID_SECRET"))
			}
			if strings.TrimSpace(clientName) == "" {
				clientName = strings.TrimSpace(state.GetenvAny("PLAID_CLIENT_NAME"))
			}

			if env == "" || clientID == "" || secret == "" || clientName == "" {
				return errors.New("missing required app profile fields; provide --env, --client-id, --secret, and --client-name, or set PLAID_ENV/PLAID_CLIENT_ID/PLAID_SECRET/PLAID_CLIENT_NAME")
			}

			profile := state.AppProfile{
				Env:          env,
				ClientID:     clientID,
				Secret:       secret,
				ClientName:   clientName,
				Language:     language,
				CountryCodes: countryCodes,
			}
			if err := profile.Validate(); err != nil {
				return err
			}

			if err := store.Ensure(); err != nil {
				return err
			}
			if err := store.SaveConfig(state.DefaultConfig()); err != nil {
				return err
			}
			if err := store.SaveAppProfile(profile); err != nil {
				return err
			}

			stateDir, err := getStateDir(cmd)
			if err != nil {
				return err
			}

			return writeJSON(cmd, map[string]any{
				"ok":        true,
				"state_dir": stateDir,
				"doc":       setupDocPath,
				"app_profile": map[string]any{
					"env":           profile.Env,
					"client_name":   profile.ClientName,
					"language":      profile.Language,
					"country_codes": profile.CountryCodes,
				},
			})
		},
	}

	cmd.Flags().StringVar(&env, "env", "", "Plaid environment: sandbox, development, or production")
	cmd.Flags().StringVar(&clientID, "client-id", "", "Plaid client_id")
	cmd.Flags().StringVar(&secret, "secret", "", "Plaid secret")
	cmd.Flags().StringVar(&clientName, "client-name", "plaid-cli", "Display name used in Plaid Link")
	cmd.Flags().StringVar(&language, "language", "en", "Language used for Link sessions")
	cmd.Flags().StringSliceVar(&countryCodes, "country-code", []string{"US"}, "Country code to enable for Link (repeatable)")

	return cmd
}
