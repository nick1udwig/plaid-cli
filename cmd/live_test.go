package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"plaid-cli/internal/plaid"
	"plaid-cli/internal/state"
)

const (
	liveTestGateEnv           = "PLAID_RUN_LIVE_TESTS"
	liveSandboxClientIDEnv    = "PLAID_SANDBOX_CLIENT_ID"
	liveSandboxSecretEnv      = "PLAID_SANDBOX_SECRET"
	liveFallbackClientIDEnv   = "PLAID_CLIENT_ID"
	liveFallbackSecretEnv     = "PLAID_SECRET"
	liveSandboxInstitutionEnv = "PLAID_SANDBOX_INSTITUTION_ID"
	liveMicrodepositTokenEnv  = "PLAID_LIVE_AUTOMATED_MICRODEPOSIT_ACCESS_TOKEN"
	liveMicrodepositAcctEnv   = "PLAID_LIVE_AUTOMATED_MICRODEPOSIT_ACCOUNT_ID"
)

type liveSandboxConfig struct {
	ClientID      string
	Secret        string
	InstitutionID string
	ClientName    string
}

type liveSandboxItem struct {
	ItemID      string
	AccessToken string
}

type liveSandboxHarness struct {
	t            *testing.T
	stateDir     string
	cleanupItems []liveSandboxItem
}

type commandRunError struct {
	Args   []string
	Stdout string
	Stderr string
	Err    error
}

func (e *commandRunError) Error() string {
	return fmt.Sprintf(
		"command %q failed: %v\nstdout:\n%s\nstderr:\n%s",
		strings.Join(e.Args, " "),
		e.Err,
		e.Stdout,
		e.Stderr,
	)
}

func (e *commandRunError) Unwrap() error {
	return e.Err
}

func TestLiveSandboxSmokeSuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)
	harness := newLiveSandboxHarness(t)
	cleanupClient := newLiveSandboxClient(t, cfg)
	t.Cleanup(func() {
		harness.cleanup(t, cleanupClient)
	})

	initResp := harness.mustRunJSON("init", "--env", "sandbox", "--client-id", cfg.ClientID, "--secret", cfg.Secret, "--client-name", cfg.ClientName)
	if got := requireStringField(t, initResp, "app_profile", "env"); got != "sandbox" {
		t.Fatalf("init app_profile.env = %q, want sandbox", got)
	}

	institutionResp := harness.mustRunJSON("institution", "get-by-id", "--institution-id", cfg.InstitutionID)
	if requireStringField(t, institutionResp, "institution", "name") == "" {
		t.Fatal("institution response did not include institution.name")
	}
	institutionListResp := harness.mustRunJSON("institution", "get", "--count", "5")
	if len(requireArrayField(t, institutionListResp, "institutions")) == 0 {
		t.Fatal("institution.get returned no institutions")
	}
	institutionSearchResp := harness.mustRunJSON("institution", "search", "--query", "Chase", "--product", "auth")
	if len(requireArrayField(t, institutionSearchResp, "institutions")) == 0 {
		t.Fatal("institution.search returned no institutions")
	}

	linkResp := harness.mustRunJSON("link", "token-create", "--product", "auth", "--client-user-id", "plaid-cli-live-test")
	linkToken := requireStringField(t, linkResp, "link_token")
	if linkToken == "" {
		t.Fatal("link token response did not include link_token")
	}
	linkGetResp := harness.mustRunJSON("link", "token-get", "--link-token", linkToken)
	if requireStringField(t, linkGetResp, "request_id") == "" {
		t.Fatal("link token-get response did not include request_id")
	}

	publicTokenResp := harness.mustRunJSON(
		"sandbox",
		"public-token-create",
		"--institution-id", cfg.InstitutionID,
		"--product", "auth",
		"--product", "transactions",
	)
	publicToken := requireStringField(t, publicTokenResp, "public_token")

	exchangeResp := harness.mustRunJSON(
		"item",
		"public-token-exchange",
		"--public-token", publicToken,
		"--product", "auth",
		"--product", "transactions",
	)
	itemID := requireStringField(t, exchangeResp, "item_id")
	accessToken := requireStringField(t, exchangeResp, "access_token")
	harness.trackItem(itemID, accessToken)

	store := state.New(harness.stateDir)
	record, err := store.LoadItem(itemID)
	if err != nil {
		t.Fatalf("LoadItem(%q) error = %v", itemID, err)
	}
	if record.AccessToken == "" {
		t.Fatal("saved item record did not include access_token")
	}
	if len(record.Accounts) == 0 {
		t.Fatal("saved item record did not include any accounts")
	}

	listResp := harness.mustRunJSON("item", "list")
	items := requireArrayField(t, listResp, "items")
	if len(items) != 1 {
		t.Fatalf("item list length = %d, want 1", len(items))
	}

	itemResp := harness.mustRunJSON("item", "get", "--item", itemID)
	if got := requireStringField(t, itemResp, "item", "item_id"); got != itemID {
		t.Fatalf("item.get item_id = %q, want %q", got, itemID)
	}

	accountResp := harness.mustRunJSON("account", "get", "--item", itemID)
	accounts := requireArrayField(t, accountResp, "accounts")
	if len(accounts) == 0 {
		t.Fatal("account.get returned no accounts")
	}

	authResp := harness.mustRunJSON("auth", "get", "--item", itemID)
	authAccounts := requireArrayField(t, authResp, "accounts")
	if len(authAccounts) == 0 {
		t.Fatal("auth.get returned no accounts")
	}

	balanceResp := harness.mustRunJSON("balance", "get", "--item", itemID)
	if len(requireArrayField(t, balanceResp, "accounts")) == 0 {
		t.Fatal("balance.get returned no accounts")
	}

	transactionsResp := harness.mustRunJSONRetryProductReady(
		10,
		3*time.Second,
		"transactions",
		"sync",
		"--item", itemID,
		"--count", "25",
	)
	if requireStringField(t, transactionsResp, "next_cursor") == "" {
		t.Fatal("transactions.sync did not include next_cursor")
	}
	startDate := time.Now().UTC().AddDate(0, 0, -365).Format("2006-01-02")
	endDate := time.Now().UTC().Format("2006-01-02")
	transactionsGetResp := harness.mustRunJSONRetryProductReady(
		10,
		3*time.Second,
		"transactions",
		"get",
		"--item", itemID,
		"--start-date", startDate,
		"--end-date", endDate,
		"--count", "25",
	)
	if len(requireArrayField(t, transactionsGetResp, "transactions")) == 0 {
		t.Fatal("transactions.get returned no transactions")
	}
	transactionsRecurringResp := harness.mustRunJSONRetryProductReady(
		10,
		3*time.Second,
		"transactions",
		"recurring-get",
		"--item", itemID,
	)
	requireArrayField(t, transactionsRecurringResp, "inflow_streams")
	requireArrayField(t, transactionsRecurringResp, "outflow_streams")

	invalidateResp := harness.mustRunJSON("item", "access-token-invalidate", "--item", itemID)
	newAccessToken := requireStringField(t, invalidateResp, "new_access_token")
	if newAccessToken == accessToken {
		t.Fatal("item.access-token-invalidate returned the existing access token")
	}
	harness.updateItemAccessToken(itemID, newAccessToken)
	updatedRecord, err := store.LoadItem(itemID)
	if err != nil {
		t.Fatalf("LoadItem(%q) after invalidate error = %v", itemID, err)
	}
	if updatedRecord.AccessToken != newAccessToken {
		t.Fatalf("saved item access token = %q, want %q", updatedRecord.AccessToken, newAccessToken)
	}

	webhookUpdateResp := harness.mustRunJSON(
		"item",
		"webhook-update",
		"--item", itemID,
		"--webhook-url", "https://example.com/plaid-cli-live-test",
	)
	if requireStringField(t, webhookUpdateResp, "request_id") == "" {
		t.Fatal("item.webhook-update did not include request_id")
	}

	fireWebhookResp := harness.mustRunJSON(
		"sandbox",
		"item-fire-webhook",
		"--item", itemID,
		"--webhook-type", "TRANSACTIONS",
		"--webhook-code", "DEFAULT_UPDATE",
	)
	if requireStringField(t, fireWebhookResp, "request_id") == "" {
		t.Fatal("sandbox item-fire-webhook did not include request_id")
	}

	microdepositAccessToken := strings.TrimSpace(os.Getenv(liveMicrodepositTokenEnv))
	microdepositAccountID := strings.TrimSpace(os.Getenv(liveMicrodepositAcctEnv))
	if microdepositAccessToken != "" && microdepositAccountID != "" {
		setVerificationResp := harness.mustRunJSON(
			"sandbox",
			"item-set-verification-status",
			"--access-token", microdepositAccessToken,
			"--account-id", microdepositAccountID,
			"--verification-status", "automatically_verified",
		)
		if requireStringField(t, setVerificationResp, "request_id") == "" {
			t.Fatal("sandbox item-set-verification-status did not include request_id")
		}
	} else {
		t.Logf(
			"skipping sandbox item-set-verification-status smoke test; set %s and %s to run it against a pre-created automated micro-deposit item",
			liveMicrodepositTokenEnv,
			liveMicrodepositAcctEnv,
		)
	}

	resetResp := harness.mustRunJSON("sandbox", "item-reset-login", "--item", itemID)
	if requireStringField(t, resetResp, "request_id") == "" {
		t.Fatal("sandbox item-reset-login did not include request_id")
	}

	removeResp := harness.mustRunJSON("item", "remove", "--item", itemID)
	localDeleted, ok := bodyValue(removeResp, "local_item_deleted")
	if !ok {
		t.Fatal("item.remove did not include local_item_deleted")
	}
	if deleted, ok := localDeleted.(bool); !ok || !deleted {
		t.Fatalf("item.remove local_item_deleted = %#v, want true", localDeleted)
	}
	harness.untrackItem(itemID)

	listAfterResp := harness.mustRunJSON("item", "list")
	itemsAfter := requireArrayField(t, listAfterResp, "items")
	if len(itemsAfter) != 0 {
		t.Fatalf("item list length after remove = %d, want 0", len(itemsAfter))
	}
}

func loadLiveSandboxConfig(t *testing.T) liveSandboxConfig {
	t.Helper()

	if testing.Short() {
		t.Skip("skipping live Plaid sandbox smoke tests in -short mode")
	}
	if !envTruthy(os.Getenv(liveTestGateEnv)) {
		t.Skipf("set %s=1 to run live Plaid sandbox smoke tests", liveTestGateEnv)
	}

	clientID := strings.TrimSpace(state.GetenvAny(liveSandboxClientIDEnv, liveFallbackClientIDEnv))
	secret := strings.TrimSpace(state.GetenvAny(liveSandboxSecretEnv, liveFallbackSecretEnv))
	institutionID := strings.TrimSpace(os.Getenv(liveSandboxInstitutionEnv))
	if institutionID == "" {
		institutionID = "ins_109508"
	}

	clientName := "plaid-cli live tests"
	if clientID == "" || secret == "" {
		defaultDir, err := state.DefaultDir()
		if err != nil {
			t.Fatalf("resolve default state dir: %v", err)
		}
		profile, err := state.New(defaultDir).LoadAppProfile()
		if err != nil {
			t.Fatalf("live sandbox tests require %s/%s, %s/%s, or a saved sandbox profile in %s", liveSandboxClientIDEnv, liveSandboxSecretEnv, liveFallbackClientIDEnv, liveFallbackSecretEnv, defaultDir)
		}
		if strings.TrimSpace(strings.ToLower(profile.Env)) != "sandbox" {
			t.Fatalf("saved app profile in %s uses env %q, want sandbox", defaultDir, profile.Env)
		}
		clientID = profile.ClientID
		secret = profile.Secret
		if strings.TrimSpace(profile.ClientName) != "" {
			clientName = profile.ClientName
		}
	}

	return liveSandboxConfig{
		ClientID:      clientID,
		Secret:        secret,
		InstitutionID: institutionID,
		ClientName:    clientName,
	}
}

func newLiveSandboxHarness(t *testing.T) *liveSandboxHarness {
	t.Helper()
	return &liveSandboxHarness{
		t:        t,
		stateDir: t.TempDir(),
	}
}

func (h *liveSandboxHarness) trackItem(itemID, accessToken string) {
	h.cleanupItems = append(h.cleanupItems, liveSandboxItem{
		ItemID:      itemID,
		AccessToken: accessToken,
	})
}

func (h *liveSandboxHarness) untrackItem(itemID string) {
	filtered := h.cleanupItems[:0]
	for _, item := range h.cleanupItems {
		if item.ItemID == itemID {
			continue
		}
		filtered = append(filtered, item)
	}
	h.cleanupItems = filtered
}

func (h *liveSandboxHarness) updateItemAccessToken(itemID, accessToken string) {
	for i := range h.cleanupItems {
		if h.cleanupItems[i].ItemID == itemID {
			h.cleanupItems[i].AccessToken = accessToken
			return
		}
	}
	h.trackItem(itemID, accessToken)
}

func (h *liveSandboxHarness) mustRunJSON(args ...string) map[string]any {
	h.t.Helper()

	resp, err := h.runJSON(args...)
	if err != nil {
		h.t.Fatal(err)
	}
	return resp
}

func (h *liveSandboxHarness) mustRunJSONRetryProductReady(attempts int, delay time.Duration, args ...string) map[string]any {
	h.t.Helper()

	for attempt := 1; attempt <= attempts; attempt++ {
		resp, err := h.runJSON(args...)
		if err == nil {
			return resp
		}

		var apiErr *plaid.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode == "PRODUCT_NOT_READY" && attempt < attempts {
			h.t.Logf("%s not ready yet (attempt %d/%d); retrying in %s", strings.Join(args, " "), attempt, attempts, delay)
			time.Sleep(delay)
			continue
		}

		h.t.Fatal(err)
	}

	h.t.Fatalf("%s did not become ready after %d attempts", strings.Join(args, " "), attempts)
	return nil
}

func (h *liveSandboxHarness) runJSON(args ...string) (map[string]any, error) {
	h.t.Helper()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	root := NewRootCmd(&Options{
		Stdout:        stdout,
		Stderr:        stderr,
		BrowserOpener: func(string) error { return nil },
		StateDir:      h.stateDir,
	})
	root.SetArgs(args)
	if _, err := root.ExecuteC(); err != nil {
		return nil, &commandRunError{
			Args:   append([]string(nil), args...),
			Stdout: stdout.String(),
			Stderr: stderr.String(),
			Err:    err,
		}
	}

	var resp map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &resp); err != nil {
		return nil, fmt.Errorf("decode command output for %q: %w\nstdout:\n%s\nstderr:\n%s", strings.Join(args, " "), err, stdout.String(), stderr.String())
	}
	return resp, nil
}

func (h *liveSandboxHarness) cleanup(t *testing.T, client *plaid.Client) {
	t.Helper()

	for i := len(h.cleanupItems) - 1; i >= 0; i-- {
		item := h.cleanupItems[i]
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		_, err := client.Call(ctx, "/item/remove", map[string]any{
			"access_token": item.AccessToken,
		})
		cancel()
		if err != nil {
			t.Logf("cleanup: remove sandbox item %s: %v", item.ItemID, err)
		}
	}
}

func newLiveSandboxClient(t *testing.T, cfg liveSandboxConfig) *plaid.Client {
	t.Helper()

	client, err := plaid.NewClient(state.AppProfile{
		Env:      "sandbox",
		ClientID: cfg.ClientID,
		Secret:   cfg.Secret,
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	return client
}

func requireStringField(t *testing.T, body map[string]any, path ...string) string {
	t.Helper()

	value, ok := bodyValue(body, path...)
	if !ok {
		t.Fatalf("response missing %s", strings.Join(path, "."))
	}
	valueString, ok := value.(string)
	if !ok {
		t.Fatalf("response field %s = %#v, want string", strings.Join(path, "."), value)
	}
	return valueString
}

func requireArrayField(t *testing.T, body map[string]any, path ...string) []any {
	t.Helper()

	value, ok := bodyValue(body, path...)
	if !ok {
		t.Fatalf("response missing %s", strings.Join(path, "."))
	}
	valueArray, ok := value.([]any)
	if !ok {
		t.Fatalf("response field %s = %#v, want array", strings.Join(path, "."), value)
	}
	return valueArray
}

func envTruthy(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "y", "on":
		return true
	default:
		return false
	}
}
