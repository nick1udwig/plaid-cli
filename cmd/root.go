package cmd

import (
	"fmt"
	"io"
	"os"

	"plaid-cli/internal/browser"
	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

const setupDocPath = "docs/getting-started.md"

type Options struct {
	Stdout        io.Writer
	Stderr        io.Writer
	BrowserOpener browser.Opener
	StateDir      string
}

func defaultOptions() *Options {
	defaultStateDir, err := state.DefaultDir()
	if err != nil {
		defaultStateDir = "~/.plaid-cli"
	}

	return &Options{
		Stdout:        os.Stdout,
		Stderr:        os.Stderr,
		BrowserOpener: browser.OpenURL,
		StateDir:      defaultStateDir,
	}
}

func Execute() {
	rootCmd := NewRootCmd(defaultOptions())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(rootCmd.ErrOrStderr(), err)
		os.Exit(1)
	}
}

func NewRootCmd(opts *Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "plaid",
		Short:         "Agent-friendly CLI for the Plaid API",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       "0.1.0",
		Long: fmt.Sprintf(
			"plaid is a single-owner CLI for working with the Plaid API.\n\n"+
				"Persistent state is stored under ~/.plaid-cli by default.\n"+
				"Humans should complete one-time setup with `plaid init` before using Link.\n"+
				"See %s in this repo for setup details.",
			setupDocPath,
		),
	}

	cmd.SetOut(opts.Stdout)
	cmd.SetErr(opts.Stderr)
	cmd.PersistentFlags().String("state-dir", opts.StateDir, "Directory for persistent Plaid CLI state")

	cmd.AddCommand(newInitCmd(opts))
	cmd.AddCommand(newLinkCmd(opts))
	cmd.AddCommand(newItemCmd(opts))
	cmd.AddCommand(newAccountCmd())
	cmd.AddCommand(newAssetsCmd())
	cmd.AddCommand(newInstitutionCmd())
	cmd.AddCommand(newLiabilitiesCmd())
	cmd.AddCommand(newInvestmentsCmd())
	cmd.AddCommand(newStatementsCmd())
	cmd.AddCommand(newPartnerCmd())
	cmd.AddCommand(newProcessorCmd())
	cmd.AddCommand(newUserCmd())
	cmd.AddCommand(newDashboardUserCmd())
	cmd.AddCommand(newOAuthCmd())
	cmd.AddCommand(newConsentCmd())
	cmd.AddCommand(newNetworkCmd())
	cmd.AddCommand(newAssetsCmd())
	cmd.AddCommand(newLiabilitiesCmd())
	cmd.AddCommand(newInvestmentsCmd())
	cmd.AddCommand(newStatementsCmd())
	cmd.AddCommand(newEnrichCmd())
	cmd.AddCommand(newInvestmentsMoveCmd())
	cmd.AddCommand(newPaymentInitiationCmd())
	cmd.AddCommand(newCheckCmd())
	cmd.AddCommand(newIncomeCmd())
	cmd.AddCommand(newIdentityVerificationCmd())
	cmd.AddCommand(newLayerCmd())
	cmd.AddCommand(newWalletCmd())
	cmd.AddCommand(newAuthCmd())
	cmd.AddCommand(newBalanceCmd())
	cmd.AddCommand(newSignalCmd())
	cmd.AddCommand(newIdentityCmd())
	cmd.AddCommand(newTransactionsCmd())
	cmd.AddCommand(newTransferCmd())
	cmd.AddCommand(newSandboxCmd())

	return cmd
}
