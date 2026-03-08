package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"plaid-cli/internal/plaid"
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

func TestResolveAccountID(t *testing.T) {
	t.Parallel()

	t.Run("uses explicit account id", func(t *testing.T) {
		t.Parallel()

		accountID, err := resolveAccountID(nil, "acct-explicit")
		if err != nil {
			t.Fatalf("resolveAccountID() error = %v", err)
		}
		if accountID != "acct-explicit" {
			t.Fatalf("accountID = %q, want %q", accountID, "acct-explicit")
		}
	})

	t.Run("auto-selects the only saved account", func(t *testing.T) {
		t.Parallel()

		record := &state.ItemRecord{
			ItemID: "item-1",
			Accounts: []state.AccountSummary{
				{AccountID: "acct-1"},
			},
		}

		accountID, err := resolveAccountID(record, "")
		if err != nil {
			t.Fatalf("resolveAccountID() error = %v", err)
		}
		if accountID != "acct-1" {
			t.Fatalf("accountID = %q, want %q", accountID, "acct-1")
		}
	})

	t.Run("fails when no saved item record exists", func(t *testing.T) {
		t.Parallel()

		_, err := resolveAccountID(nil, "")
		if err == nil {
			t.Fatal("resolveAccountID() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--account-id is required") {
			t.Fatalf("error = %q, want account-id guidance", err)
		}
	})

	t.Run("fails when multiple saved accounts exist", func(t *testing.T) {
		t.Parallel()

		record := &state.ItemRecord{
			ItemID: "item-2",
			Accounts: []state.AccountSummary{
				{AccountID: "acct-a"},
				{AccountID: "acct-b"},
			},
		}

		_, err := resolveAccountID(record, "")
		if err == nil {
			t.Fatal("resolveAccountID() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "acct-a") || !strings.Contains(err.Error(), "acct-b") {
			t.Fatalf("error = %q, want both account IDs", err)
		}
	})
}

func TestLoadRequestBody(t *testing.T) {
	t.Parallel()

	t.Run("returns empty object for empty input", func(t *testing.T) {
		t.Parallel()

		body, err := loadRequestBody("")
		if err != nil {
			t.Fatalf("loadRequestBody() error = %v", err)
		}
		if len(body) != 0 {
			t.Fatalf("body = %#v, want empty object", body)
		}
	})

	t.Run("loads inline JSON object", func(t *testing.T) {
		t.Parallel()

		body, err := loadRequestBody(`{"access_token":"access-123","user":{"legal_name":"Jane"}}`)
		if err != nil {
			t.Fatalf("loadRequestBody() error = %v", err)
		}
		want := map[string]any{
			"access_token": "access-123",
			"user": map[string]any{
				"legal_name": "Jane",
			},
		}
		if !reflect.DeepEqual(body, want) {
			t.Fatalf("body = %#v, want %#v", body, want)
		}
	})

	t.Run("loads JSON object from file", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()
		path := filepath.Join(dir, "request.json")
		if err := os.WriteFile(path, []byte(`{"amount":"10.00"}`), 0o600); err != nil {
			t.Fatalf("WriteFile() error = %v", err)
		}

		body, err := loadRequestBody("@" + path)
		if err != nil {
			t.Fatalf("loadRequestBody() error = %v", err)
		}
		want := map[string]any{"amount": "10.00"}
		if !reflect.DeepEqual(body, want) {
			t.Fatalf("body = %#v, want %#v", body, want)
		}
	})

	t.Run("rejects non-object json", func(t *testing.T) {
		t.Parallel()

		_, err := loadRequestBody(`["not","an","object"]`)
		if err == nil {
			t.Fatal("loadRequestBody() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "JSON object") {
			t.Fatalf("error = %q, want JSON object message", err)
		}
	})
}

func TestWriteBinaryOutput(t *testing.T) {
	t.Parallel()

	t.Run("writes file and emits metadata json", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()
		outPath := filepath.Join(dir, "reports", "report.pdf")
		cmd := &cobra.Command{Use: "test"}
		buf := &bytes.Buffer{}
		cmd.SetOut(buf)

		resp := &plaid.BinaryResponse{
			Body: []byte("%PDF-1.7"),
			Headers: map[string][]string{
				"Plaid-Content-Hash": {"sha256:abc123"},
				"Plaid-Request-ID":   {"req-123"},
			},
		}

		if err := writeBinaryOutput(cmd, outPath, resp); err != nil {
			t.Fatalf("writeBinaryOutput() error = %v", err)
		}

		got, err := os.ReadFile(outPath)
		if err != nil {
			t.Fatalf("ReadFile() error = %v", err)
		}
		if string(got) != "%PDF-1.7" {
			t.Fatalf("file contents = %q, want PDF bytes", got)
		}

		output := buf.String()
		if !strings.Contains(output, `"path": "`+outPath+`"`) {
			t.Fatalf("output = %q, want path metadata", output)
		}
		if !strings.Contains(output, `"plaid_content_hash": "sha256:abc123"`) {
			t.Fatalf("output = %q, want content hash metadata", output)
		}
		if !strings.Contains(output, `"request_id": "req-123"`) {
			t.Fatalf("output = %q, want request_id metadata", output)
		}
	})

	t.Run("requires output path", func(t *testing.T) {
		t.Parallel()

		err := writeBinaryOutput(&cobra.Command{}, "", &plaid.BinaryResponse{})
		if err == nil {
			t.Fatal("writeBinaryOutput() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--out is required") {
			t.Fatalf("error = %q, want out guidance", err)
		}
	})
}

func TestSetBodyValue(t *testing.T) {
	t.Parallel()

	t.Run("creates nested objects", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{}
		if err := setBodyValue(body, "Jane Doe", "user", "legal_name"); err != nil {
			t.Fatalf("setBodyValue() error = %v", err)
		}

		want := map[string]any{
			"user": map[string]any{
				"legal_name": "Jane Doe",
			},
		}
		if !reflect.DeepEqual(body, want) {
			t.Fatalf("body = %#v, want %#v", body, want)
		}
	})

	t.Run("preserves existing nested object", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{
			"user": map[string]any{
				"legal_name": "Jane Doe",
			},
		}
		if err := setBodyValue(body, "jane@example.com", "user", "email_address"); err != nil {
			t.Fatalf("setBodyValue() error = %v", err)
		}

		if got, ok := bodyValue(body, "user", "email_address"); !ok || got != "jane@example.com" {
			t.Fatalf("bodyValue() = %#v, %v; want %q, true", got, ok, "jane@example.com")
		}
	})

	t.Run("rejects non-object intermediate field", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{"user": "not-an-object"}
		err := setBodyValue(body, "Jane Doe", "user", "legal_name")
		if err == nil {
			t.Fatal("setBodyValue() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "user") {
			t.Fatalf("error = %q, want field name", err)
		}
	})
}

func TestApplyStringFlag(t *testing.T) {
	t.Parallel()

	t.Run("preserves body value when flag is not changed", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		cmd.Flags().String("network", "same-day-ach", "")
		body := map[string]any{"network": "ach"}

		if err := applyStringFlag(cmd, body, "network", "same-day-ach", "network"); err != nil {
			t.Fatalf("applyStringFlag() error = %v", err)
		}
		if got, _ := bodyValue(body, "network"); got != "ach" {
			t.Fatalf("body[network] = %#v, want %q", got, "ach")
		}
	})

	t.Run("sets default when body does not already define the field", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		cmd.Flags().String("network", "same-day-ach", "")
		body := map[string]any{}

		if err := applyStringFlag(cmd, body, "network", "same-day-ach", "network"); err != nil {
			t.Fatalf("applyStringFlag() error = %v", err)
		}
		if got, _ := bodyValue(body, "network"); got != "same-day-ach" {
			t.Fatalf("body[network] = %#v, want %q", got, "same-day-ach")
		}
	})
}

func TestApplyOptionalIntFlag(t *testing.T) {
	t.Parallel()

	t.Run("does not set zero-value default when flag is unchanged", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		cmd.Flags().Int("days-requested", 0, "")
		body := map[string]any{}

		if err := applyOptionalIntFlag(cmd, body, "days-requested", 0, "options", "days_requested"); err != nil {
			t.Fatalf("applyOptionalIntFlag() error = %v", err)
		}
		if _, ok := bodyValue(body, "options", "days_requested"); ok {
			t.Fatalf("body unexpectedly contains options.days_requested: %#v", body)
		}
	})

	t.Run("sets zero when the flag was explicitly provided", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		cmd.Flags().Int("days-requested", 0, "")
		if err := cmd.Flags().Set("days-requested", "0"); err != nil {
			t.Fatalf("Set() error = %v", err)
		}
		body := map[string]any{}

		if err := applyOptionalIntFlag(cmd, body, "days-requested", 0, "options", "days_requested"); err != nil {
			t.Fatalf("applyOptionalIntFlag() error = %v", err)
		}
		if got, ok := bodyValue(body, "options", "days_requested"); !ok || got != 0 {
			t.Fatalf("bodyValue() = %#v, %v; want 0, true", got, ok)
		}
	})
}

func TestApplyDecimalStringFlag(t *testing.T) {
	t.Parallel()

	t.Run("parses decimal value from changed flag", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		cmd.Flags().String("amount", "", "")
		if err := cmd.Flags().Set("amount", "102.05"); err != nil {
			t.Fatalf("Set() error = %v", err)
		}

		body := map[string]any{}
		if err := applyDecimalStringFlag(cmd, body, "amount", "102.05", "amount"); err != nil {
			t.Fatalf("applyDecimalStringFlag() error = %v", err)
		}
		if got, _ := bodyValue(body, "amount"); got != 102.05 {
			t.Fatalf("body[amount] = %#v, want %v", got, 102.05)
		}
	})

	t.Run("preserves existing numeric body value when flag is not changed", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		cmd.Flags().String("amount", "", "")

		body := map[string]any{"amount": 55.5}
		if err := applyDecimalStringFlag(cmd, body, "amount", "", "amount"); err != nil {
			t.Fatalf("applyDecimalStringFlag() error = %v", err)
		}
		if got, _ := bodyValue(body, "amount"); got != 55.5 {
			t.Fatalf("body[amount] = %#v, want %v", got, 55.5)
		}
	})

	t.Run("rejects invalid decimal values", func(t *testing.T) {
		t.Parallel()

		cmd := &cobra.Command{Use: "test"}
		cmd.Flags().String("amount", "", "")
		if err := cmd.Flags().Set("amount", "abc"); err != nil {
			t.Fatalf("Set() error = %v", err)
		}

		err := applyDecimalStringFlag(cmd, map[string]any{}, "amount", "abc", "amount")
		if err == nil {
			t.Fatal("applyDecimalStringFlag() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--amount") {
			t.Fatalf("error = %q, want flag name", err)
		}
	})
}

func TestPopulateTransferAccess(t *testing.T) {
	t.Parallel()

	t.Run("fills account id from saved item matched by body access token", func(t *testing.T) {
		t.Parallel()

		store := state.New(t.TempDir())
		record := state.ItemRecord{
			ItemID:      "item-transfer",
			AccessToken: "access-transfer",
			Accounts: []state.AccountSummary{
				{AccountID: "acct-transfer"},
			},
		}
		if err := store.SaveItem(record); err != nil {
			t.Fatalf("SaveItem() error = %v", err)
		}

		body := map[string]any{
			"access_token": "access-transfer",
		}
		gotRecord, err := populateTransferAccess(&cobra.Command{Use: "test"}, store, body, "", "", "")
		if err != nil {
			t.Fatalf("populateTransferAccess() error = %v", err)
		}
		if gotRecord == nil || gotRecord.ItemID != record.ItemID {
			t.Fatalf("record = %#v, want item_id %q", gotRecord, record.ItemID)
		}
		if got, ok := bodyValue(body, "account_id"); !ok || got != "acct-transfer" {
			t.Fatalf("body account_id = %#v, %v; want %q, true", got, ok, "acct-transfer")
		}
	})
}

func TestRequireExactlyOneBodyField(t *testing.T) {
	t.Parallel()

	fields := map[string][]string{
		"--user-id":    {"user_id"},
		"--user-token": {"user_token"},
	}

	t.Run("accepts exactly one field", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{"user_id": "usr_123"}
		if err := requireExactlyOneBodyField(body, fields); err != nil {
			t.Fatalf("requireExactlyOneBodyField() error = %v", err)
		}
	})

	t.Run("rejects zero fields", func(t *testing.T) {
		t.Parallel()

		err := requireExactlyOneBodyField(map[string]any{}, fields)
		if err == nil {
			t.Fatal("requireExactlyOneBodyField() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--user-id") || !strings.Contains(err.Error(), "--user-token") {
			t.Fatalf("error = %q, want both field labels", err)
		}
	})

	t.Run("rejects multiple fields", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{
			"user_id":    "usr_123",
			"user_token": "user-token-123",
		}
		err := requireExactlyOneBodyField(body, fields)
		if err == nil {
			t.Fatal("requireExactlyOneBodyField() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--user-id") || !strings.Contains(err.Error(), "--user-token") {
			t.Fatalf("error = %q, want both field labels", err)
		}
	})
}

func TestRequireAtLeastOneBodyField(t *testing.T) {
	t.Parallel()

	fields := map[string][]string{
		"--iban":         {"iban"},
		"--bacs-account": {"bacs", "account"},
	}

	t.Run("accepts when a field is present", func(t *testing.T) {
		t.Parallel()

		body := map[string]any{"iban": "GB33BUKB20201555555555"}
		if err := requireAtLeastOneBodyField(body, fields); err != nil {
			t.Fatalf("requireAtLeastOneBodyField() error = %v", err)
		}
	})

	t.Run("rejects when no field is present", func(t *testing.T) {
		t.Parallel()

		err := requireAtLeastOneBodyField(map[string]any{}, fields)
		if err == nil {
			t.Fatal("requireAtLeastOneBodyField() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--iban") || !strings.Contains(err.Error(), "--bacs-account") {
			t.Fatalf("error = %q, want both field labels", err)
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
