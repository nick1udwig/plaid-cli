package cmd

import (
	"strings"
	"testing"
)

func TestRequireSandboxUserRef(t *testing.T) {
	t.Parallel()

	t.Run("accepts user_id", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{"user_id": "usr_123"}
		if err := requireSandboxUserRef(body); err != nil {
			t.Fatalf("requireSandboxUserRef() error = %v", err)
		}
	})

	t.Run("requires one identifier", func(t *testing.T) {
		t.Parallel()

		err := requireSandboxUserRef(map[string]any{})
		if err == nil {
			t.Fatal("requireSandboxUserRef() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--user-id") || !strings.Contains(err.Error(), "--user-token") {
			t.Fatalf("error = %q, want user reference guidance", err)
		}
	})
}

func TestBuildSandboxTransactionFromFlags(t *testing.T) {
	t.Parallel()

	t.Run("builds a transaction entry", func(t *testing.T) {
		t.Parallel()

		entry, ok, err := buildSandboxTransactionFromFlags("2026-03-08", "2026-03-08", "12.34", "Coffee Shop", "USD")
		if err != nil {
			t.Fatalf("buildSandboxTransactionFromFlags() error = %v", err)
		}
		if !ok {
			t.Fatal("buildSandboxTransactionFromFlags() ok = false, want true")
		}
		if got := entry["description"]; got != "Coffee Shop" {
			t.Fatalf("description = %#v, want Coffee Shop", got)
		}
		if got := entry["amount"]; got != 12.34 {
			t.Fatalf("amount = %#v, want 12.34", got)
		}
	})

	t.Run("requires all core fields when any are provided", func(t *testing.T) {
		t.Parallel()

		_, _, err := buildSandboxTransactionFromFlags("", "2026-03-08", "12.34", "Coffee Shop", "")
		if err == nil {
			t.Fatal("buildSandboxTransactionFromFlags() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--date-transacted") {
			t.Fatalf("error = %q, want date-transacted guidance", err)
		}
	})
}
