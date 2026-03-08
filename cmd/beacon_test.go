package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestBeaconCreateValidation(t *testing.T) {
	t.Parallel()

	t.Run("requires birth date or depository accounts", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{
			"program_id":     "becprg_123",
			"client_user_id": "user-123",
			"user": map[string]any{
				"name": map[string]any{
					"given_name":  "Leslie",
					"family_name": "Knope",
				},
			},
		}

		err := validateBeaconUserCreateBody(body)
		if err == nil {
			t.Fatal("validateBeaconUserCreateBody() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "user.date_of_birth") {
			t.Fatalf("error = %q, want date_of_birth guidance", err)
		}
	})
}

func TestBeaconUserIdentityValidation(t *testing.T) {
	t.Parallel()

	t.Run("rejects partial address objects", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "beacon"}
		flags := bindBeaconUserIdentityFlags(cmd)
		body := map[string]any{}

		if err := cmd.Flags().Set("street", "123 Main St"); err != nil {
			t.Fatalf("Set(street) error = %v", err)
		}
		if err := cmd.Flags().Set("city", "Pawnee"); err != nil {
			t.Fatalf("Set(city) error = %v", err)
		}
		flags.street = "123 Main St"
		flags.city = "Pawnee"

		err := applyBeaconUserIdentityFlags(cmd, body, flags)
		if err == nil {
			t.Fatal("applyBeaconUserIdentityFlags() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "user.address.country") {
			t.Fatalf("error = %q, want country guidance", err)
		}
	})

	t.Run("rejects partial id numbers", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "beacon"}
		flags := bindBeaconUserIdentityFlags(cmd)
		body := map[string]any{}

		if err := cmd.Flags().Set("id-number", "1234"); err != nil {
			t.Fatalf("Set(id-number) error = %v", err)
		}
		flags.idNumber = "1234"

		err := applyBeaconUserIdentityFlags(cmd, body, flags)
		if err == nil {
			t.Fatal("applyBeaconUserIdentityFlags() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "user.id_number.type") {
			t.Fatalf("error = %q, want type guidance", err)
		}
	})
}
