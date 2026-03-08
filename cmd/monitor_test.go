package cmd

import (
	"strings"
	"testing"
)

func TestValidateMonitorIndividualCreate(t *testing.T) {
	t.Parallel()

	body := map[string]any{
		"search_terms": map[string]any{
			"watchlist_program_id": "program-1",
			"legal_name":           "Jane Doe",
		},
	}

	err := validateMonitorIndividualCreate(body)
	if err == nil {
		t.Fatal("validateMonitorIndividualCreate() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "--client-user-id") || !strings.Contains(err.Error(), "--user-id") {
		t.Fatalf("error = %q, want user reference guidance", err)
	}
}

func TestValidateMonitorIndividualUpdate(t *testing.T) {
	t.Parallel()

	t.Run("requires at least one update field", func(t *testing.T) {
		t.Parallel()

		err := validateMonitorIndividualUpdate(map[string]any{})
		if err == nil {
			t.Fatal("validateMonitorIndividualUpdate() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "provide at least one update field") {
			t.Fatalf("error = %q, want update field guidance", err)
		}
	})

	t.Run("rejects search terms with status", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{
			"search_terms": map[string]any{
				"watchlist_program_id": "program-1",
				"legal_name":           "Jane Doe",
			},
			"status": "cleared",
		}

		err := validateMonitorIndividualUpdate(body)
		if err == nil {
			t.Fatal("validateMonitorIndividualUpdate() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "search_terms and status cannot be updated in the same request") {
			t.Fatalf("error = %q, want status conflict guidance", err)
		}
	})
}

func TestValidateMonitorEntityUpdate(t *testing.T) {
	t.Parallel()

	body := map[string]any{
		"search_terms": map[string]any{
			"legal_name": "Example LLC",
		},
	}

	err := validateMonitorEntityUpdate(body)
	if err == nil {
		t.Fatal("validateMonitorEntityUpdate() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "search_terms.entity_watchlist_program_id") {
		t.Fatalf("error = %q, want entity watchlist program guidance", err)
	}
}
