package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestValidateBeaconUser(t *testing.T) {
	t.Parallel()

	t.Run("requires name and one identifier on create", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{
			"user": map[string]any{
				"name": map[string]any{
					"given_name": "Jane",
				},
			},
		}

		err := validateBeaconUser(body, true)
		if err == nil {
			t.Fatal("validateBeaconUser() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "user.name.family_name") {
			t.Fatalf("error = %q, want family_name guidance", err)
		}
	})

	t.Run("rejects partial depository account entries", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{
			"user": map[string]any{
				"name": map[string]any{
					"given_name":  "Jane",
					"family_name": "Doe",
				},
				"depository_accounts": []any{
					map[string]any{
						"routing_number": "021000021",
					},
				},
			},
		}

		err := validateBeaconUser(body, true)
		if err == nil {
			t.Fatal("validateBeaconUser() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "account_number") {
			t.Fatalf("error = %q, want account_number guidance", err)
		}
	})
}

func TestBeaconUserHasPatch(t *testing.T) {
	t.Parallel()

	cmd := &cobra.Command{Use: "beacon"}
	flags := bindBeaconUserFlags(cmd)
	body := map[string]any{}

	if err := cmd.Flags().Set("access-token", "access-1"); err != nil {
		t.Fatalf("Set(access-token) error = %v", err)
	}
	flags.accessTokens = []string{"access-1"}

	if err := applyBeaconUserFlags(cmd, body, flags); err != nil {
		t.Fatalf("applyBeaconUserFlags() error = %v", err)
	}
	if !beaconUserHasPatch(body) {
		t.Fatal("beaconUserHasPatch() = false, want true")
	}
}
