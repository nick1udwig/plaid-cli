package cmd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestPaymentInitiationAddressHelpers(t *testing.T) {
	t.Parallel()

	t.Run("builds and validates an address object from typed flags", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		var street []string
		var city, postalCode, country string
		cmd.Flags().StringSliceVar(&street, "address-street", nil, "")
		cmd.Flags().StringVar(&city, "address-city", "", "")
		cmd.Flags().StringVar(&postalCode, "address-postal-code", "", "")
		cmd.Flags().StringVar(&country, "address-country", "", "")
		for _, flag := range []struct {
			name  string
			value string
		}{
			{name: "address-street", value: "96 Guild Street,9th Floor"},
			{name: "address-city", value: "London"},
			{name: "address-postal-code", value: "SE14 8JW"},
			{name: "address-country", value: "GB"},
		} {
			if err := cmd.Flags().Set(flag.name, flag.value); err != nil {
				t.Fatalf("Set(%q) error = %v", flag.name, err)
			}
		}

		body := map[string]any{}
		if err := applyPIAddressFlags(cmd, body, street, city, postalCode, country, "address-street", "address-city", "address-postal-code", "address-country", "address"); err != nil {
			t.Fatalf("applyPIAddressFlags() error = %v", err)
		}
		if err := validatePIAddress(body, "address", "address"); err != nil {
			t.Fatalf("validatePIAddress() error = %v", err)
		}

		want := map[string]any{
			"address": map[string]any{
				"street":      []string{"96 Guild Street", "9th Floor"},
				"city":        "London",
				"postal_code": "SE14 8JW",
				"country":     "GB",
			},
		}
		if !reflect.DeepEqual(body, want) {
			t.Fatalf("body = %#v, want %#v", body, want)
		}
	})

	t.Run("rejects partial address objects", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{
			"address": map[string]any{
				"city": "London",
			},
		}
		err := validatePIAddress(body, "address", "address")
		if err == nil {
			t.Fatal("validatePIAddress() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "address.street") {
			t.Fatalf("error = %q, want street guidance", err)
		}
	})
}

func TestPaymentInitiationConsentPeriodicHelpers(t *testing.T) {
	t.Parallel()

	t.Run("builds and validates the first periodic amount entry", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		var currency, value, interval, alignment string
		cmd.Flags().StringVar(&currency, "periodic-amount-currency", "", "")
		cmd.Flags().StringVar(&value, "periodic-amount-value", "", "")
		cmd.Flags().StringVar(&interval, "periodic-interval", "", "")
		cmd.Flags().StringVar(&alignment, "periodic-alignment", "", "")
		for _, flag := range []struct {
			name  string
			value string
		}{
			{name: "periodic-amount-currency", value: "GBP"},
			{name: "periodic-amount-value", value: "40.00"},
			{name: "periodic-interval", value: "MONTH"},
			{name: "periodic-alignment", value: "CALENDAR"},
		} {
			if err := cmd.Flags().Set(flag.name, flag.value); err != nil {
				t.Fatalf("Set(%q) error = %v", flag.name, err)
			}
		}

		body := map[string]any{}
		if err := applyPIConsentPeriodicFlags(cmd, body, currency, value, interval, alignment); err != nil {
			t.Fatalf("applyPIConsentPeriodicFlags() error = %v", err)
		}
		if err := validatePIConsentPeriodicFlags(body); err != nil {
			t.Fatalf("validatePIConsentPeriodicFlags() error = %v", err)
		}

		rawEntry, ok, err := firstObjectFromArrayPath(body, "constraints", "periodic_amounts")
		if err != nil || !ok {
			t.Fatalf("firstObjectFromArrayPath() = %#v, %v, %v; want entry, true, nil", rawEntry, ok, err)
		}
		amount := rawEntry["amount"].(map[string]any)
		if amount["currency"] != "GBP" || amount["value"] != 40.0 {
			t.Fatalf("amount = %#v, want GBP/40.0", amount)
		}
	})

	t.Run("rejects partial periodic entries", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{
			"constraints": map[string]any{
				"periodic_amounts": []any{
					map[string]any{
						"interval": "MONTH",
					},
				},
			},
		}
		err := validatePIConsentPeriodicFlags(body)
		if err == nil {
			t.Fatal("validatePIConsentPeriodicFlags() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "constraints.periodic_amounts[0].amount") {
			t.Fatalf("error = %q, want amount guidance", err)
		}
	})
}
