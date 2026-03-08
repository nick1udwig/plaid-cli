package cmd

import (
	"strings"
	"testing"

	"plaid-cli/internal/state"

	"github.com/spf13/cobra"
)

func TestResolveAccessToken(t *testing.T) {
	t.Parallel()

	t.Run("uses explicit access token without local record", func(t *testing.T) {
		t.Parallel()

		store := state.New(t.TempDir())

		token, record, err := resolveAccessToken(&cobra.Command{}, store, "", "access-explicit")
		if err != nil {
			t.Fatalf("resolveAccessToken() error = %v", err)
		}
		if token != "access-explicit" {
			t.Fatalf("token = %q, want %q", token, "access-explicit")
		}
		if record != nil {
			t.Fatalf("record = %#v, want nil", record)
		}
	})

	t.Run("uses explicit access token and returns matching local record", func(t *testing.T) {
		t.Parallel()

		store := state.New(t.TempDir())
		item := state.ItemRecord{
			ItemID:      "item-123",
			AccessToken: "access-123",
		}
		if err := store.SaveItem(item); err != nil {
			t.Fatalf("SaveItem() error = %v", err)
		}

		token, record, err := resolveAccessToken(&cobra.Command{}, store, "", "access-123")
		if err != nil {
			t.Fatalf("resolveAccessToken() error = %v", err)
		}
		if token != item.AccessToken {
			t.Fatalf("token = %q, want %q", token, item.AccessToken)
		}
		if record == nil || record.ItemID != item.ItemID {
			t.Fatalf("record = %#v, want item_id %q", record, item.ItemID)
		}
	})

	t.Run("loads access token from item id", func(t *testing.T) {
		t.Parallel()

		store := state.New(t.TempDir())
		item := state.ItemRecord{
			ItemID:      "item-456",
			AccessToken: "access-456",
		}
		if err := store.SaveItem(item); err != nil {
			t.Fatalf("SaveItem() error = %v", err)
		}

		token, record, err := resolveAccessToken(&cobra.Command{}, store, item.ItemID, "")
		if err != nil {
			t.Fatalf("resolveAccessToken() error = %v", err)
		}
		if token != item.AccessToken {
			t.Fatalf("token = %q, want %q", token, item.AccessToken)
		}
		if record == nil || record.ItemID != item.ItemID {
			t.Fatalf("record = %#v, want item_id %q", record, item.ItemID)
		}
	})

	t.Run("auto-selects the only saved item", func(t *testing.T) {
		t.Parallel()

		store := state.New(t.TempDir())
		item := state.ItemRecord{
			ItemID:      "item-solo",
			AccessToken: "access-solo",
		}
		if err := store.SaveItem(item); err != nil {
			t.Fatalf("SaveItem() error = %v", err)
		}

		token, record, err := resolveAccessToken(&cobra.Command{}, store, "", "")
		if err != nil {
			t.Fatalf("resolveAccessToken() error = %v", err)
		}
		if token != item.AccessToken {
			t.Fatalf("token = %q, want %q", token, item.AccessToken)
		}
		if record == nil || record.ItemID != item.ItemID {
			t.Fatalf("record = %#v, want item_id %q", record, item.ItemID)
		}
	})

	t.Run("fails when no items are saved", func(t *testing.T) {
		t.Parallel()

		store := state.New(t.TempDir())

		_, _, err := resolveAccessToken(&cobra.Command{}, store, "", "")
		if err == nil {
			t.Fatal("resolveAccessToken() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "no saved items found") {
			t.Fatalf("error = %q, want no saved items message", err)
		}
	})

	t.Run("fails when multiple items are saved", func(t *testing.T) {
		t.Parallel()

		store := state.New(t.TempDir())
		items := []state.ItemRecord{
			{ItemID: "item-a", AccessToken: "access-a"},
			{ItemID: "item-b", AccessToken: "access-b"},
		}
		for _, item := range items {
			if err := store.SaveItem(item); err != nil {
				t.Fatalf("SaveItem() error = %v", err)
			}
		}

		_, _, err := resolveAccessToken(&cobra.Command{}, store, "", "")
		if err == nil {
			t.Fatal("resolveAccessToken() error = nil, want error")
		}
		for _, item := range items {
			if !strings.Contains(err.Error(), item.ItemID) {
				t.Fatalf("error = %q, want to mention %q", err, item.ItemID)
			}
		}
	})
}

func TestLoadClientFromStateRequiresInit(t *testing.T) {
	t.Parallel()

	root := NewRootCmd(&Options{
		StateDir: t.TempDir(),
	})
	cmd := &cobra.Command{Use: "test"}
	root.AddCommand(cmd)
	cmd.SetArgs(nil)

	store, profile, client, err := loadClientFromState(cmd)
	if err == nil {
		t.Fatal("loadClientFromState() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "run `plaid init` first") {
		t.Fatalf("error = %q, want init guidance", err)
	}
	if store != nil || client != nil {
		t.Fatalf("unexpected non-nil return values: store=%v client=%v", store, client)
	}
	if profile.Env != "" || profile.ClientID != "" || profile.Secret != "" || profile.ClientName != "" || profile.Language != "" || len(profile.CountryCodes) != 0 {
		t.Fatalf("profile = %#v, want zero value", profile)
	}
}
