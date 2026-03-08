package state

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoadAppProfile(t *testing.T) {
	t.Parallel()

	store := New(t.TempDir())

	profile := AppProfile{
		Env:          "sandbox",
		ClientID:     "client-id",
		Secret:       "secret",
		ClientName:   "plaid-cli",
		Language:     "en",
		CountryCodes: []string{"US"},
	}

	if err := store.SaveAppProfile(profile); err != nil {
		t.Fatalf("SaveAppProfile() error = %v", err)
	}

	got, err := store.LoadAppProfile()
	if err != nil {
		t.Fatalf("LoadAppProfile() error = %v", err)
	}

	if got.ClientID != profile.ClientID {
		t.Fatalf("ClientID = %q, want %q", got.ClientID, profile.ClientID)
	}
	if got.Secret != profile.Secret {
		t.Fatalf("Secret = %q, want %q", got.Secret, profile.Secret)
	}
	if got.CreatedAt.IsZero() || got.UpdatedAt.IsZero() {
		t.Fatalf("timestamps were not populated: %+v", got)
	}
}

func TestSaveAndListItems(t *testing.T) {
	t.Parallel()

	store := New(t.TempDir())
	record := ItemRecord{
		ItemID:      "item-123",
		AccessToken: "access-123",
		Products:    []string{"auth", "transactions"},
		Accounts: []AccountSummary{
			{AccountID: "acct-1", Name: "Checking"},
		},
	}

	if err := store.SaveItem(record); err != nil {
		t.Fatalf("SaveItem() error = %v", err)
	}

	items, err := store.ListItems()
	if err != nil {
		t.Fatalf("ListItems() error = %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("len(items) = %d, want 1", len(items))
	}
	if items[0].ItemID != record.ItemID {
		t.Fatalf("ItemID = %q, want %q", items[0].ItemID, record.ItemID)
	}
	if got, want := store.ItemPath(record.ItemID), filepath.Join(store.dir, "items", "item-123.json"); got != want {
		t.Fatalf("ItemPath() = %q, want %q", got, want)
	}
}

func TestFindAndDeleteItem(t *testing.T) {
	t.Parallel()

	store := New(t.TempDir())
	record := ItemRecord{
		ItemID:      "item-find-delete",
		AccessToken: "access-find-delete",
	}

	if err := store.SaveItem(record); err != nil {
		t.Fatalf("SaveItem() error = %v", err)
	}

	found, err := store.FindItemByAccessToken(record.AccessToken)
	if err != nil {
		t.Fatalf("FindItemByAccessToken() error = %v", err)
	}
	if found == nil || found.ItemID != record.ItemID {
		t.Fatalf("found = %#v, want item_id %q", found, record.ItemID)
	}

	if err := store.DeleteItem(record.ItemID); err != nil {
		t.Fatalf("DeleteItem() error = %v", err)
	}

	_, err = store.LoadItem(record.ItemID)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("LoadItem() error = %v, want os.ErrNotExist", err)
	}
}
