package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestIdentityVerificationUserAddressValidation(t *testing.T) {
	cmd := &cobra.Command{Use: "idv"}
	flags := bindIdentityVerificationUserFlags(cmd)
	body := map[string]any{}

	if err := cmd.Flags().Set("country", "US"); err != nil {
		t.Fatalf("set country flag: %v", err)
	}
	flags.country = "US"

	if err := applyIdentityVerificationUserAddressFlags(cmd, body, flags); err != nil {
		t.Fatalf("apply country-only address: %v", err)
	}
	if !bodyHasValue(body, "user", "address", "country") {
		t.Fatalf("expected country-only user address to be set")
	}
}

func TestIdentityVerificationRetryStepsValidation(t *testing.T) {
	cmd := &cobra.Command{Use: "idv-retry"}
	flags := bindIdentityVerificationRetryStepFlags(cmd)
	body := map[string]any{"strategy": "custom"}

	if err := cmd.Flags().Set("verify-sms", "true"); err != nil {
		t.Fatalf("set verify-sms: %v", err)
	}
	flags.verifySMS = true

	err := applyIdentityVerificationRetrySteps(cmd, body, flags)
	if err == nil {
		t.Fatal("expected missing custom retry steps to fail")
	}
}

func TestWalletAddressValidation(t *testing.T) {
	cmd := &cobra.Command{Use: "wallet"}
	body := map[string]any{}
	flags := walletAddressFlags{
		street:  []string{"1 Main St"},
		city:    "Paris",
		country: "FR",
	}

	cmd.Flags().StringSlice("counterparty-address-street", nil, "")
	cmd.Flags().String("counterparty-address-city", "", "")
	cmd.Flags().String("counterparty-address-postal-code", "", "")
	cmd.Flags().String("counterparty-address-country", "", "")
	if err := cmd.Flags().Set("counterparty-address-street", "1 Main St"); err != nil {
		t.Fatalf("set street flag: %v", err)
	}
	if err := cmd.Flags().Set("counterparty-address-city", "Paris"); err != nil {
		t.Fatalf("set city flag: %v", err)
	}
	if err := cmd.Flags().Set("counterparty-address-country", "FR"); err != nil {
		t.Fatalf("set country flag: %v", err)
	}

	err := applyWalletAddress(cmd, body, "counterparty-address", flags, "counterparty", "address")
	if err == nil {
		t.Fatal("expected partial wallet address to fail")
	}
}
