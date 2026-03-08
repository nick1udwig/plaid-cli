package cmd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestApplyUserIdentityFlags(t *testing.T) {
	t.Parallel()

	t.Run("builds primary identity arrays from typed flags", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		flags := bindUserIdentityFlags(cmd)
		values := map[string]string{
			"given-name":     "Carmen",
			"family-name":    "Berzatto",
			"date-of-birth":  "1987-01-31",
			"email":          "carmy@example.com",
			"phone-number":   "+13125551212",
			"street-1":       "3200 W Armitage Ave",
			"city":           "Chicago",
			"region":         "IL",
			"country":        "US",
			"postal-code":    "60657",
			"id-number":      "1234",
			"id-number-type": "us_ssn_last_4",
		}
		for name, value := range values {
			if err := cmd.Flags().Set(name, value); err != nil {
				t.Fatalf("Set(%q) error = %v", name, err)
			}
		}

		body := map[string]any{}
		if err := applyUserIdentityFlags(cmd, body, flags); err != nil {
			t.Fatalf("applyUserIdentityFlags() error = %v", err)
		}

		want := map[string]any{
			"identity": map[string]any{
				"name": map[string]any{
					"given_name":  "Carmen",
					"family_name": "Berzatto",
				},
				"date_of_birth": "1987-01-31",
				"emails": []map[string]any{
					{
						"data":    "carmy@example.com",
						"primary": true,
					},
				},
				"phone_numbers": []map[string]any{
					{
						"data":    "+13125551212",
						"primary": true,
					},
				},
				"addresses": []map[string]any{
					{
						"street_1":    "3200 W Armitage Ave",
						"city":        "Chicago",
						"region":      "IL",
						"country":     "US",
						"postal_code": "60657",
						"primary":     true,
					},
				},
				"id_numbers": []map[string]any{
					{
						"value": "1234",
						"type":  "us_ssn_last_4",
					},
				},
			},
		}
		if !reflect.DeepEqual(body, want) {
			t.Fatalf("body = %#v, want %#v", body, want)
		}
	})

	t.Run("requires both id-number flags together", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		flags := bindUserIdentityFlags(cmd)
		if err := cmd.Flags().Set("id-number", "1234"); err != nil {
			t.Fatalf("Set() error = %v", err)
		}

		err := applyUserIdentityFlags(cmd, map[string]any{}, flags)
		if err == nil {
			t.Fatal("applyUserIdentityFlags() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--id-number") {
			t.Fatalf("error = %q, want id-number guidance", err)
		}
	})
}
