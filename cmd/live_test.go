package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"plaid-cli/internal/plaid"
	"plaid-cli/internal/state"
)

const (
	liveTestGateEnv           = "PLAID_RUN_LIVE_TESTS"
	liveIncomeTestGateEnv     = "PLAID_RUN_LIVE_INCOME_TESTS"
	liveCheckTestGateEnv      = "PLAID_RUN_LIVE_CHECK_TESTS"
	liveSandboxClientIDEnv    = "PLAID_SANDBOX_CLIENT_ID"
	liveSandboxSecretEnv      = "PLAID_SANDBOX_SECRET"
	liveFallbackClientIDEnv   = "PLAID_CLIENT_ID"
	liveFallbackSecretEnv     = "PLAID_SECRET"
	liveSandboxInstitutionEnv = "PLAID_SANDBOX_INSTITUTION_ID"
	liveMicrodepositTokenEnv  = "PLAID_LIVE_AUTOMATED_MICRODEPOSIT_ACCESS_TOKEN"
	liveMicrodepositAcctEnv   = "PLAID_LIVE_AUTOMATED_MICRODEPOSIT_ACCOUNT_ID"
	liveIncomeItemIDEnv       = "PLAID_LIVE_INCOME_ITEM_ID"
	liveCheckUserIDEnv        = "PLAID_LIVE_CHECK_USER_ID"
	liveCheckItemIDEnv        = "PLAID_LIVE_CHECK_ITEM_ID"
	liveWebhookSiteTokenURL   = "https://webhook.site/token"
	liveWebhookSiteBaseURL    = "https://webhook.site"
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

type liveSandboxUser struct {
	UserID string
}

type webhookSiteTokenResponse struct {
	UUID string `json:"uuid"`
}

type webhookSiteRequestsResponse struct {
	Data []webhookSiteRequest `json:"data"`
}

type webhookSiteRequest struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Content string              `json:"content"`
}

type liveSandboxHarness struct {
	t                    *testing.T
	stateDir             string
	cleanupItems         []liveSandboxItem
	cleanupUsers         []liveSandboxUser
	cleanupSubscriptions []string
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

	harness.initializeAppProfile(cfg)

	institutionResp := harness.mustRunJSON("institution", "get-by-id", "--institution-id", cfg.InstitutionID)
	if requireStringField(t, institutionResp, "institution", "name") == "" {
		t.Fatal("institution response did not include institution.name")
	}
	institutionListResp := harness.mustRunJSONRetryProductReady(
		5,
		3*time.Second,
		"institution",
		"get",
		"--count", "5",
	)
	if len(requireArrayField(t, institutionListResp, "institutions")) == 0 {
		t.Fatal("institution.get returned no institutions")
	}
	institutionSearchResp := harness.mustRunJSONRetryProductReady(
		5,
		3*time.Second,
		"institution",
		"search",
		"--query", "Chase",
		"--product", "auth",
	)
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

	userID := harness.createUser("plaid-cli-live-user-reset", false)
	userGetResp := harness.mustRunJSON("user", "get", "--user-id", userID)
	if got := requireStringField(t, userGetResp, "user_id"); got != userID {
		t.Fatalf("user.get user_id = %q, want %q", got, userID)
	}
	userItem := harness.createSandboxItem(cfg, []string{"auth"}, "--user-id", userID)
	userResetResp := harness.mustRunJSON(
		"sandbox",
		"user-reset-login",
		"--user-id", userID,
		"--item-id", userItem.ItemID,
	)
	if requireStringField(t, userResetResp, "request_id") == "" {
		t.Fatal("sandbox user-reset-login did not include request_id")
	}
	userRemoveResp := harness.mustRunJSON("item", "remove", "--item", userItem.ItemID)
	localUserItemDeleted, ok := bodyValue(userRemoveResp, "local_item_deleted")
	if !ok {
		t.Fatal("user-linked item.remove did not include local_item_deleted")
	}
	if deleted, ok := localUserItemDeleted.(bool); !ok || !deleted {
		t.Fatalf("user-linked item.remove local_item_deleted = %#v, want true", localUserItemDeleted)
	}
	harness.untrackItem(userItem.ItemID)

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

func TestLiveSandboxDynamicTransactionsSuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)
	harness := newLiveSandboxHarness(t)
	cleanupClient := newLiveSandboxClient(t, cfg)
	t.Cleanup(func() {
		harness.cleanup(t, cleanupClient)
	})

	harness.initializeAppProfile(cfg)
	item := harness.createSandboxItem(
		cfg,
		[]string{"transactions"},
		"--override-username", "user_transactions_dynamic",
		"--override-password", "plaid-cli-live-test",
	)

	beforeSyncResp := harness.mustRunJSONRetryProductReady(
		10,
		3*time.Second,
		"transactions",
		"sync",
		"--item", item.ItemID,
		"--count", "100",
	)
	cursor := requireStringField(t, beforeSyncResp, "next_cursor")

	today := time.Now().UTC().Format("2006-01-02")
	refreshResp := harness.mustRunJSON("transactions", "refresh", "--item", item.ItemID)
	if requireStringField(t, refreshResp, "request_id") == "" {
		t.Fatal("transactions.refresh did not include request_id")
	}

	createResp := harness.mustRunJSON(
		"sandbox",
		"transactions-create",
		"--item", item.ItemID,
		"--date-transacted", today,
		"--date-posted", today,
		"--amount", "12.34",
		"--description", "plaid-cli live test transaction",
		"--currency", "USD",
	)
	if requireStringField(t, createResp, "request_id") == "" {
		t.Fatal("sandbox transactions-create did not include request_id")
	}

	foundAddedTransaction := false
	for attempt := 1; attempt <= 10; attempt++ {
		afterSyncResp := harness.mustRunJSONRetryProductReady(
			10,
			3*time.Second,
			"transactions",
			"sync",
			"--item", item.ItemID,
			"--cursor", cursor,
			"--count", "100",
		)
		if len(requireArrayField(t, afterSyncResp, "added")) > 0 {
			foundAddedTransaction = true
			break
		}
		if attempt < 10 {
			t.Logf("transactions.sync after sandbox transactions-create returned no added transactions yet (attempt %d/10); retrying", attempt)
			time.Sleep(2 * time.Second)
		}
	}
	if !foundAddedTransaction {
		t.Fatal("transactions.sync after sandbox transactions-create returned no added transactions")
	}
}

func TestLiveSandboxProcessorSuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)

	harness := newLiveSandboxHarness(t)
	harness.initializeAppProfile(cfg)

	resp := harness.mustRunJSON(
		"sandbox",
		"processor-token-create",
		"--institution-id", cfg.InstitutionID,
	)
	processorToken := requireStringField(t, resp, "processor_token")
	if processorToken == "" {
		t.Fatal("sandbox processor-token-create did not include processor_token")
	}

	permissionsResp := harness.mustRunJSON("processor", "token-permissions-get", "--processor-token", processorToken)
	if requireStringField(t, permissionsResp, "request_id") == "" {
		t.Fatal("processor token-permissions-get did not include request_id")
	}

	accountResp := harness.mustRunJSON("processor", "account", "get", "--processor-token", processorToken)
	if requireStringField(t, accountResp, "account", "account_id") == "" {
		t.Fatal("processor account.get did not include account.account_id")
	}

	authResp := harness.mustRunJSON("processor", "auth", "get", "--processor-token", processorToken)
	if requireStringField(t, authResp, "numbers", "ach", "account") == "" {
		t.Fatal("processor auth.get did not include numbers.ach.account")
	}

	balanceResp := harness.mustRunJSON("processor", "balance", "get", "--processor-token", processorToken)
	if requireStringField(t, balanceResp, "account", "account_id") == "" {
		t.Fatal("processor balance.get did not include account.account_id")
	}

	identityResp := harness.mustRunJSON("processor", "identity", "get", "--processor-token", processorToken)
	if requireStringField(t, identityResp, "account", "account_id") == "" {
		t.Fatal("processor identity.get did not include account.account_id")
	}

	startDate := time.Now().UTC().AddDate(0, 0, -365).Format("2006-01-02")
	endDate := time.Now().UTC().Format("2006-01-02")
	transactionsGetResp := harness.mustRunJSONRetryProductReady(
		10,
		3*time.Second,
		"processor",
		"transactions",
		"get",
		"--processor-token", processorToken,
		"--start-date", startDate,
		"--end-date", endDate,
		"--count", "25",
	)
	if len(requireArrayField(t, transactionsGetResp, "transactions")) == 0 {
		t.Fatal("processor transactions.get returned no transactions")
	}

	transactionsSyncResp := harness.mustRunJSONRetryProductReady(
		10,
		3*time.Second,
		"processor",
		"transactions",
		"sync",
		"--processor-token", processorToken,
		"--count", "25",
	)
	if requireStringField(t, transactionsSyncResp, "next_cursor") == "" {
		t.Fatal("processor transactions.sync did not include next_cursor")
	}

	transactionsRecurringResp := harness.mustRunJSONRetryProductReady(
		10,
		3*time.Second,
		"processor",
		"transactions",
		"recurring-get",
		"--processor-token", processorToken,
	)
	requireArrayField(t, transactionsRecurringResp, "inflow_streams")
	requireArrayField(t, transactionsRecurringResp, "outflow_streams")

	transactionsRefreshResp := harness.mustRunJSON(
		"processor",
		"transactions",
		"refresh",
		"--processor-token", processorToken,
	)
	if requireStringField(t, transactionsRefreshResp, "request_id") == "" {
		t.Fatal("processor transactions.refresh did not include request_id")
	}
}

func TestLiveSandboxAdditionalReadProductsSuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)
	harness := newLiveSandboxHarness(t)
	cleanupClient := newLiveSandboxClient(t, cfg)
	t.Cleanup(func() {
		harness.cleanup(t, cleanupClient)
	})

	harness.initializeAppProfile(cfg)

	t.Run("identity get", func(t *testing.T) {
		item, err := harness.tryCreateSandboxItem(cfg, []string{"identity"})
		if err != nil {
			skipUnavailableLiveProduct(t, "identity", err)
		}

		resp, err := harness.runJSON("identity", "get", "--item", item.ItemID)
		if err != nil {
			skipUnavailableLiveProduct(t, "identity", err)
		}
		if len(requireArrayField(t, resp, "accounts")) == 0 {
			t.Fatal("identity.get returned no accounts")
		}
	})

	t.Run("liabilities get", func(t *testing.T) {
		item, err := harness.tryCreateSandboxItem(cfg, []string{"liabilities"})
		if err != nil {
			skipUnavailableLiveProduct(t, "liabilities", err)
		}

		resp, err := harness.runJSON("liabilities", "get", "--item", item.ItemID)
		if err != nil {
			skipUnavailableLiveProduct(t, "liabilities", err)
		}
		if len(requireArrayField(t, resp, "accounts")) == 0 {
			t.Fatal("liabilities.get returned no accounts")
		}
		if _, ok := bodyValue(resp, "liabilities"); !ok {
			t.Fatal("liabilities.get did not include liabilities")
		}
	})

	t.Run("investments holdings and transactions", func(t *testing.T) {
		item, err := harness.tryCreateSandboxItem(cfg, []string{"investments"})
		if err != nil {
			skipUnavailableLiveProduct(t, "investments", err)
		}

		holdingsResp, err := harness.runJSON("investments", "holdings-get", "--item", item.ItemID)
		if err != nil {
			skipUnavailableLiveProduct(t, "investments holdings", err)
		}
		if len(requireArrayField(t, holdingsResp, "accounts")) == 0 {
			t.Fatal("investments holdings-get returned no accounts")
		}
		if _, ok := bodyValue(holdingsResp, "holdings"); !ok {
			t.Fatal("investments holdings-get did not include holdings")
		}

		startDate := time.Now().UTC().AddDate(0, -1, 0).Format("2006-01-02")
		endDate := time.Now().UTC().Format("2006-01-02")
		transactionsResp, err := harness.runJSON(
			"investments",
			"transactions-get",
			"--item", item.ItemID,
			"--start-date", startDate,
			"--end-date", endDate,
		)
		if err != nil {
			skipUnavailableLiveProduct(t, "investments transactions", err)
		}
		if len(requireArrayField(t, transactionsResp, "accounts")) == 0 {
			t.Fatal("investments transactions-get returned no accounts")
		}
		if _, ok := bodyValue(transactionsResp, "investment_transactions"); !ok {
			t.Fatal("investments transactions-get did not include investment_transactions")
		}
	})

	t.Run("signal evaluate", func(t *testing.T) {
		item, err := harness.tryCreateSandboxItem(cfg, []string{"auth"})
		if err != nil {
			skipUnavailableLiveProduct(t, "signal bootstrap auth item", err)
		}
		accountID := harness.requireItemAccountID(item.ItemID)

		resp, err := harness.runJSON(
			"signal",
			"evaluate",
			"--item", item.ItemID,
			"--account-id", accountID,
			"--client-transaction-id", fmt.Sprintf("signal-live-%d", time.Now().UTC().UnixNano()),
			"--amount", "12.34",
		)
		if err != nil {
			skipUnavailableLiveProduct(t, "signal", err)
		}
		if requireStringField(t, resp, "request_id") == "" {
			t.Fatal("signal.evaluate did not include request_id")
		}
	})

	t.Run("statements list", func(t *testing.T) {
		statementStartDate := time.Now().UTC().AddDate(0, -3, 0).Format("2006-01-02")
		statementEndDate := time.Now().UTC().Format("2006-01-02")
		item, err := harness.tryCreateSandboxItem(
			cfg,
			[]string{"statements"},
			"--body", marshalBodyArg(t, map[string]any{
				"options": map[string]any{
					"statements": map[string]any{
						"start_date": statementStartDate,
						"end_date":   statementEndDate,
					},
				},
			}),
		)
		if err != nil {
			skipUnavailableLiveProduct(t, "statements", err)
		}

		resp, err := harness.runJSON("statements", "list", "--item", item.ItemID)
		if err != nil {
			skipUnavailableLiveProduct(t, "statements", err)
		}
		if len(requireArrayField(t, resp, "accounts")) == 0 {
			t.Fatal("statements.list returned no accounts")
		}
	})
}

func TestLiveSandboxPaymentInitiationSuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)

	harness := newLiveSandboxHarness(t)
	harness.initializeAppProfile(cfg)

	recipientResp := harness.mustRunJSON(
		"payment-initiation",
		"recipient",
		"create",
		"--name", "Plaid CLI Live Test Recipient",
		"--iban", "GB29NWBK60161331926819",
	)
	recipientID := requireStringField(t, recipientResp, "recipient_id")
	recipientGetResp := harness.mustRunJSON(
		"payment-initiation",
		"recipient",
		"get",
		"--recipient-id", recipientID,
	)
	if got := requireStringField(t, recipientGetResp, "recipient_id"); got != recipientID {
		t.Fatalf("payment-initiation recipient.get recipient_id = %q, want %q", got, recipientID)
	}
	recipientListResp := harness.mustRunJSON(
		"payment-initiation",
		"recipient",
		"list",
		"--count", "10",
	)
	if !arrayContainsMapField(requireArrayField(t, recipientListResp, "recipients"), "recipient_id", recipientID) {
		t.Fatalf("payment-initiation recipient.list did not include recipient %q", recipientID)
	}

	paymentResp := harness.mustRunJSON(
		"payment-initiation",
		"payment",
		"create",
		"--recipient-id", recipientID,
		"--reference", "PlaidCliLiveTest",
		"--amount-currency", "GBP",
		"--amount-value", "10.50",
	)
	paymentID := requireStringField(t, paymentResp, "payment_id")
	paymentGetResp := harness.mustRunJSON(
		"payment-initiation",
		"payment",
		"get",
		"--payment-id", paymentID,
	)
	if got := requireStringField(t, paymentGetResp, "payment_id"); got != paymentID {
		t.Fatalf("payment-initiation payment.get payment_id = %q, want %q", got, paymentID)
	}
	paymentListResp := harness.mustRunJSON(
		"payment-initiation",
		"payment",
		"list",
		"--count", "10",
	)
	if !arrayContainsMapField(requireArrayField(t, paymentListResp, "payments"), "payment_id", paymentID) {
		t.Fatalf("payment-initiation payment.list did not include payment %q", paymentID)
	}

	consentResp := harness.mustRunJSON(
		"payment-initiation",
		"consent",
		"create",
		"--recipient-id", recipientID,
		"--reference", "Consent123",
		"--type", "COMMERCIAL",
		"--max-amount-currency", "GBP",
		"--max-amount-value", "15.00",
		"--periodic-amount-currency", "GBP",
		"--periodic-amount-value", "40.00",
		"--periodic-interval", "MONTH",
		"--periodic-alignment", "CALENDAR",
	)
	consentID := requireStringField(t, consentResp, "consent_id")
	consentStatus := requireStringField(t, consentResp, "status")
	consentGetResp := harness.mustRunJSON(
		"payment-initiation",
		"consent",
		"get",
		"--consent-id", consentID,
	)
	if got := requireStringField(t, consentGetResp, "consent_id"); got != consentID {
		t.Fatalf("payment-initiation consent.get consent_id = %q, want %q", got, consentID)
	}

	paymentWebhookURL := "https://example.com/plaid-cli-live-payment"
	var paymentWebhookInbox *liveWebhookInbox
	inbox, err := newLiveWebhookInbox()
	if err != nil {
		t.Logf("skipping payment webhook delivery check: %v", err)
	} else {
		paymentWebhookInbox = inbox
		paymentWebhookURL = paymentWebhookInbox.url
	}
	simulateResp := harness.mustRunJSON(
		"sandbox",
		"payment-simulate",
		"--payment-id", paymentID,
		"--webhook", paymentWebhookURL,
		"--status", "PAYMENT_STATUS_EXECUTED",
	)
	if requireStringField(t, simulateResp, "request_id") == "" {
		t.Fatal("sandbox payment-simulate did not include request_id")
	}

	executedSeen := false
	for attempt := 1; attempt <= 10; attempt++ {
		paymentGetAfterSimulateResp := harness.mustRunJSON(
			"payment-initiation",
			"payment",
			"get",
			"--payment-id", paymentID,
		)
		if got := requireStringField(t, paymentGetAfterSimulateResp, "payment_id"); got != paymentID {
			t.Fatalf("payment-initiation payment.get payment_id = %q, want %q", got, paymentID)
		}
		if got := requireStringField(t, paymentGetAfterSimulateResp, "status"); got == "PAYMENT_STATUS_EXECUTED" {
			executedSeen = true
			break
		}
		if attempt < 10 {
			t.Logf("payment.get has not shown PAYMENT_STATUS_EXECUTED yet (attempt %d/10); retrying", attempt)
			time.Sleep(2 * time.Second)
		}
	}
	if !executedSeen {
		t.Fatal("payment.get did not report PAYMENT_STATUS_EXECUTED after sandbox payment-simulate")
	}

	if paymentWebhookInbox != nil {
		request, payload := paymentWebhookInbox.mustWaitForJSONRequest(t, 45*time.Second, func(body map[string]any) bool {
			paymentValue, ok := bodyValue(body, "payment_id")
			if !ok || paymentValue != paymentID {
				return false
			}
			webhookType, ok := bodyValue(body, "webhook_type")
			if !ok || webhookType != "PAYMENT_INITIATION" {
				return false
			}
			webhookCode, ok := bodyValue(body, "webhook_code")
			if !ok || webhookCode != "PAYMENT_STATUS_UPDATE" {
				return false
			}
			newStatus, ok := bodyValue(body, "new_payment_status")
			return ok && newStatus == "PAYMENT_STATUS_EXECUTED"
		})
		if got := requireStringField(t, payload, "payment_id"); got != paymentID {
			t.Fatalf("captured payment webhook payment_id = %q, want %q", got, paymentID)
		}
		if got := requireStringField(t, payload, "webhook_type"); got != "PAYMENT_INITIATION" {
			t.Fatalf("captured payment webhook webhook_type = %q, want PAYMENT_INITIATION", got)
		}
		if got := requireStringField(t, payload, "webhook_code"); got != "PAYMENT_STATUS_UPDATE" {
			t.Fatalf("captured payment webhook webhook_code = %q, want PAYMENT_STATUS_UPDATE", got)
		}
		if got := requireStringField(t, payload, "new_payment_status"); got != "PAYMENT_STATUS_EXECUTED" {
			t.Fatalf("captured payment webhook new_payment_status = %q, want PAYMENT_STATUS_EXECUTED", got)
		}
		plaidVerification := request.headerValue("plaid-verification")
		if plaidVerification == "" {
			t.Fatal("captured payment webhook did not include plaid-verification header")
		}
		verifyResp := harness.mustRunJSON(
			"webhook",
			"verification-key",
			"get",
			"--plaid-verification", plaidVerification,
		)
		if requireStringField(t, verifyResp, "key", "kid") == "" {
			t.Fatal("payment webhook verification-key.get did not include key.kid")
		}
	}

	if consentStatus == "AUTHORISED" {
		executeResp, err := harness.runJSON(
			"payment-initiation",
			"consent",
			"payment-execute",
			"--consent-id", consentID,
			"--idempotency-key", fmt.Sprintf("consent-exec-%d", time.Now().UTC().UnixNano()),
			"--amount-currency", "GBP",
			"--amount-value", "5.25",
			"--reference", "ConsentExec123",
		)
		if err != nil {
			t.Logf("skipping consent payment-execute coverage after sandbox error: %v", err)
		} else {
			executedPaymentID := requireStringField(t, executeResp, "payment_id")
			if executedPaymentID == "" {
				t.Fatal("payment-initiation consent.payment-execute did not include payment_id")
			}
			executedPaymentGetResp := harness.mustRunJSON(
				"payment-initiation",
				"payment",
				"get",
				"--payment-id", executedPaymentID,
			)
			if got := requireStringField(t, executedPaymentGetResp, "payment_id"); got != executedPaymentID {
				t.Fatalf("payment-initiation payment.get payment_id = %q, want %q", got, executedPaymentID)
			}
		}
	} else {
		t.Logf("skipping consent payment-execute coverage because consent status is %q", consentStatus)
	}
}

func TestLiveSandboxWebhookDeliverySuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)
	harness := newLiveSandboxHarness(t)
	cleanupClient := newLiveSandboxClient(t, cfg)
	t.Cleanup(func() {
		harness.cleanup(t, cleanupClient)
	})

	harness.initializeAppProfile(cfg)
	inbox, err := newLiveWebhookInbox()
	if err != nil {
		t.Skipf("skipping live webhook delivery suite: %v", err)
	}

	item := harness.createSandboxItem(cfg, []string{"transactions"})
	webhookUpdateResp := harness.mustRunJSON(
		"item",
		"webhook-update",
		"--item", item.ItemID,
		"--webhook-url", inbox.url,
	)
	if requireStringField(t, webhookUpdateResp, "request_id") == "" {
		t.Fatal("item.webhook-update did not include request_id")
	}

	fireResp := harness.mustRunJSON(
		"sandbox",
		"item-fire-webhook",
		"--item", item.ItemID,
		"--webhook-type", "TRANSACTIONS",
		"--webhook-code", "SYNC_UPDATES_AVAILABLE",
	)
	if requireStringField(t, fireResp, "request_id") == "" {
		t.Fatal("sandbox item-fire-webhook did not include request_id")
	}

	request, payload := inbox.mustWaitForJSONRequest(t, 45*time.Second, func(body map[string]any) bool {
		itemValue, ok := bodyValue(body, "item_id")
		if !ok || itemValue != item.ItemID {
			return false
		}
		webhookType, ok := bodyValue(body, "webhook_type")
		if !ok || webhookType != "TRANSACTIONS" {
			return false
		}
		webhookCode, ok := bodyValue(body, "webhook_code")
		return ok && webhookCode == "SYNC_UPDATES_AVAILABLE"
	})

	if got := requireStringField(t, payload, "item_id"); got != item.ItemID {
		t.Fatalf("captured webhook item_id = %q, want %q", got, item.ItemID)
	}
	if got := requireStringField(t, payload, "webhook_type"); got != "TRANSACTIONS" {
		t.Fatalf("captured webhook webhook_type = %q, want TRANSACTIONS", got)
	}

	plaidVerification := request.headerValue("plaid-verification")
	if plaidVerification == "" {
		t.Fatal("captured webhook did not include plaid-verification header")
	}

	verifyResp := harness.mustRunJSON(
		"webhook",
		"verification-key",
		"get",
		"--plaid-verification", plaidVerification,
	)
	if requireStringField(t, verifyResp, "key", "kid") == "" {
		t.Fatal("webhook verification-key.get did not include key.kid")
	}

	recurringWebhookResp := harness.mustRunJSON(
		"sandbox",
		"item-fire-webhook",
		"--item", item.ItemID,
		"--webhook-code", "RECURRING_TRANSACTIONS_UPDATE",
	)
	if requireStringField(t, recurringWebhookResp, "request_id") == "" {
		t.Fatal("sandbox item-fire-webhook for recurring transactions did not include request_id")
	}
	recurringRequest, recurringPayload := inbox.mustWaitForJSONRequest(t, 45*time.Second, func(body map[string]any) bool {
		itemValue, ok := bodyValue(body, "item_id")
		if !ok || itemValue != item.ItemID {
			return false
		}
		webhookType, ok := bodyValue(body, "webhook_type")
		if !ok || webhookType != "TRANSACTIONS" {
			return false
		}
		webhookCode, ok := bodyValue(body, "webhook_code")
		return ok && webhookCode == "RECURRING_TRANSACTIONS_UPDATE"
	})
	if got := requireStringField(t, recurringPayload, "webhook_code"); got != "RECURRING_TRANSACTIONS_UPDATE" {
		t.Fatalf("captured recurring transactions webhook webhook_code = %q, want RECURRING_TRANSACTIONS_UPDATE", got)
	}
	if recurringRequest.headerValue("plaid-verification") == "" {
		t.Fatal("captured recurring transactions webhook did not include plaid-verification header")
	}

	newAccountsWebhookResp, err := harness.runJSON(
		"sandbox",
		"item-fire-webhook",
		"--item", item.ItemID,
		"--webhook-code", "NEW_ACCOUNTS_AVAILABLE",
	)
	if err != nil {
		var apiErr *plaid.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode == "SANDBOX_ACCOUNT_SELECT_V2_NOT_ENABLED" {
			t.Logf("skipping NEW_ACCOUNTS_AVAILABLE webhook delivery check: %v", err)
		} else {
			t.Fatal(err)
		}
	} else {
		if requireStringField(t, newAccountsWebhookResp, "request_id") == "" {
			t.Fatal("sandbox item-fire-webhook for new accounts did not include request_id")
		}
		_, newAccountsPayload := inbox.mustWaitForJSONRequest(t, 45*time.Second, func(body map[string]any) bool {
			itemValue, ok := bodyValue(body, "item_id")
			if !ok || itemValue != item.ItemID {
				return false
			}
			webhookType, ok := bodyValue(body, "webhook_type")
			if !ok || webhookType != "ITEM" {
				return false
			}
			webhookCode, ok := bodyValue(body, "webhook_code")
			return ok && webhookCode == "NEW_ACCOUNTS_AVAILABLE"
		})
		if got := requireStringField(t, newAccountsPayload, "webhook_code"); got != "NEW_ACCOUNTS_AVAILABLE" {
			t.Fatalf("captured item webhook webhook_code = %q, want NEW_ACCOUNTS_AVAILABLE", got)
		}
	}

	authItem := harness.createSandboxItem(cfg, []string{"auth"})
	authWebhookUpdateResp := harness.mustRunJSON(
		"item",
		"webhook-update",
		"--item", authItem.ItemID,
		"--webhook-url", inbox.url,
	)
	if requireStringField(t, authWebhookUpdateResp, "request_id") == "" {
		t.Fatal("item.webhook-update for auth item did not include request_id")
	}
	authWebhookResp := harness.mustRunJSON(
		"sandbox",
		"item-fire-webhook",
		"--item", authItem.ItemID,
		"--webhook-type", "AUTH",
		"--webhook-code", "DEFAULT_UPDATE",
	)
	if requireStringField(t, authWebhookResp, "request_id") == "" {
		t.Fatal("sandbox item-fire-webhook for auth did not include request_id")
	}
	_, authPayload := inbox.mustWaitForJSONRequest(t, 45*time.Second, func(body map[string]any) bool {
		itemValue, ok := bodyValue(body, "item_id")
		if !ok || itemValue != authItem.ItemID {
			return false
		}
		webhookType, ok := bodyValue(body, "webhook_type")
		if !ok || webhookType != "AUTH" {
			return false
		}
		webhookCode, ok := bodyValue(body, "webhook_code")
		return ok && webhookCode == "DEFAULT_UPDATE"
	})
	if got := requireStringField(t, authPayload, "webhook_type"); got != "AUTH" {
		t.Fatalf("captured auth webhook webhook_type = %q, want AUTH", got)
	}
	if got := requireStringField(t, authPayload, "webhook_code"); got != "DEFAULT_UPDATE" {
		t.Fatalf("captured auth webhook webhook_code = %q, want DEFAULT_UPDATE", got)
	}
}

func TestLiveSandboxTransferSuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)
	harness := newLiveSandboxHarness(t)
	cleanupClient := newLiveSandboxClient(t, cfg)
	t.Cleanup(func() {
		harness.cleanup(t, cleanupClient)
	})

	harness.initializeAppProfile(cfg)
	item := harness.createSandboxItem(cfg, []string{"auth"})
	accountID := harness.requireItemAccountID(item.ItemID)

	capabilitiesResp, err := harness.runJSON("transfer", "capabilities", "get", "--item", item.ItemID, "--account-id", accountID)
	if err != nil {
		if isTransferUnavailableError(err) {
			t.Skipf("skipping live Transfer suite: %v", err)
		}
		t.Fatal(err)
	}
	if requireStringField(t, capabilitiesResp, "request_id") == "" {
		t.Fatal("transfer capabilities.get did not include request_id")
	}

	clockVirtualTime := time.Now().UTC().Truncate(time.Second)
	for clockVirtualTime.Weekday() != time.Sunday {
		clockVirtualTime = clockVirtualTime.AddDate(0, 0, 1)
	}
	clockResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"test-clock",
		"create",
		"--virtual-time", clockVirtualTime.Format(time.RFC3339),
	)
	testClockID := requireStringField(t, clockResp, "test_clock", "test_clock_id")
	clockGetResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"test-clock",
		"get",
		"--test-clock-id", testClockID,
	)
	if got := requireStringField(t, clockGetResp, "test_clock", "test_clock_id"); got != testClockID {
		t.Fatalf("sandbox transfer test-clock.get test_clock_id = %q, want %q", got, testClockID)
	}
	clockListStart := clockVirtualTime.Add(-1 * time.Second).Format(time.RFC3339)
	clockListEnd := clockVirtualTime.Add(1 * time.Second).Format(time.RFC3339)
	foundClockInList := false
	for attempt := 1; attempt <= 5; attempt++ {
		clockListResp := harness.mustRunJSON(
			"sandbox",
			"transfer",
			"test-clock",
			"list",
			"--start-virtual-time", clockListStart,
			"--end-virtual-time", clockListEnd,
			"--count", "25",
		)
		if arrayContainsMapField(requireArrayField(t, clockListResp, "test_clocks"), "test_clock_id", testClockID) {
			foundClockInList = true
			break
		}
		if attempt < 5 {
			t.Logf("sandbox transfer test-clock.list has not shown test clock %q yet (attempt %d/5); retrying", testClockID, attempt)
			time.Sleep(2 * time.Second)
		}
	}
	if !foundClockInList {
		t.Fatalf("sandbox transfer test-clock.list did not include test clock %q", testClockID)
	}

	authorizationResp := harness.mustRunJSON(
		"transfer",
		"authorization",
		"create",
		"--item", item.ItemID,
		"--account-id", accountID,
		"--type", "debit",
		"--network", "ach",
		"--ach-class", "ppd",
		"--amount", "1.00",
		"--legal-name", "Plaid CLI Live Test",
	)
	authorizationID := requireStringField(t, authorizationResp, "authorization", "id")
	if got := requireStringField(t, authorizationResp, "authorization", "decision"); got != "approved" {
		t.Fatalf("transfer authorization.decision = %q, want approved", got)
	}

	createResp := harness.mustRunJSON(
		"transfer",
		"create",
		"--item", item.ItemID,
		"--account-id", accountID,
		"--authorization-id", authorizationID,
		"--amount", "1.00",
		"--description", "live transfer",
	)
	transferCreatedAt := time.Now().UTC()
	transferID := requireStringField(t, createResp, "transfer", "id")
	refundSimResp := harness.mustRunJSON(
		"transfer",
		"refund",
		"create",
		"--transfer-id", transferID,
		"--amount", "0.50",
		"--idempotency-key", fmt.Sprintf("refund-sim-%d", time.Now().UTC().UnixNano()),
	)
	refundSimID := requireStringField(t, refundSimResp, "refund", "id")
	refundSimGetResp := harness.mustRunJSON(
		"transfer",
		"refund",
		"get",
		"--refund-id", refundSimID,
	)
	if got := requireStringField(t, refundSimGetResp, "refund", "id"); got != refundSimID {
		t.Fatalf("transfer refund.get refund.id = %q, want %q", got, refundSimID)
	}

	getResp := harness.mustRunJSON("transfer", "get", "--transfer-id", transferID)
	if got := requireStringField(t, getResp, "transfer", "id"); got != transferID {
		t.Fatalf("transfer.get transfer.id = %q, want %q", got, transferID)
	}

	listResp := harness.mustRunJSON(
		"transfer",
		"list",
		"--count", "25",
		"--start-date", transferCreatedAt.Add(-1*time.Hour).Format(time.RFC3339),
		"--end-date", transferCreatedAt.Add(1*time.Hour).Format(time.RFC3339),
	)
	if !arrayContainsMapField(requireArrayField(t, listResp, "transfers"), "id", transferID) {
		t.Fatalf("transfer.list did not include transfer %q", transferID)
	}

	recurringStartDate := clockVirtualTime.AddDate(0, 0, 1)
	for recurringStartDate.Weekday() != time.Monday {
		recurringStartDate = recurringStartDate.AddDate(0, 0, 1)
	}
	recurringResp := harness.mustRunJSON(
		"transfer",
		"recurring",
		"create",
		"--item", item.ItemID,
		"--account-id", accountID,
		"--idempotency-key", fmt.Sprintf("recur-%d", time.Now().UTC().UnixNano()),
		"--type", "debit",
		"--network", "ach",
		"--ach-class", "ppd",
		"--amount", "1.00",
		"--description", "live recur",
		"--legal-name", "Plaid CLI Live Test",
		"--ip-address", "203.0.113.10",
		"--user-agent", "plaid-cli/1",
		"--interval-unit", "week",
		"--interval-count", "1",
		"--interval-execution-day", "1",
		"--start-date", recurringStartDate.Format("2006-01-02"),
		"--test-clock-id", testClockID,
	)
	recurringCreatedAt := time.Now().UTC()
	recurringTransferID := requireStringField(t, recurringResp, "recurring_transfer", "recurring_transfer_id")
	recurringGetResp := harness.mustRunJSON(
		"transfer",
		"recurring",
		"get",
		"--recurring-transfer-id", recurringTransferID,
	)
	if got := requireStringField(t, recurringGetResp, "recurring_transfer", "recurring_transfer_id"); got != recurringTransferID {
		t.Fatalf("transfer recurring.get recurring_transfer_id = %q, want %q", got, recurringTransferID)
	}
	recurringListResp := harness.mustRunJSON(
		"transfer",
		"recurring",
		"list",
		"--count", "25",
		"--start-time", recurringCreatedAt.Add(-1*time.Hour).Format(time.RFC3339),
		"--end-time", recurringCreatedAt.Add(1*time.Hour).Format(time.RFC3339),
	)
	if !arrayContainsMapField(requireArrayField(t, recurringListResp, "recurring_transfers"), "recurring_transfer_id", recurringTransferID) {
		t.Fatalf("transfer recurring.list did not include recurring transfer %q", recurringTransferID)
	}
	clockAdvanceTime := recurringStartDate.AddDate(0, 0, 8).Format(time.RFC3339)
	clockAdvanceResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"test-clock",
		"advance",
		"--test-clock-id", testClockID,
		"--new-virtual-time", clockAdvanceTime,
	)
	if requireStringField(t, clockAdvanceResp, "request_id") == "" {
		t.Fatal("sandbox transfer test-clock.advance did not include request_id")
	}
	clockGetAfterAdvanceResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"test-clock",
		"get",
		"--test-clock-id", testClockID,
	)
	if got := requireStringField(t, clockGetAfterAdvanceResp, "test_clock", "virtual_time"); got != clockAdvanceTime {
		t.Fatalf("sandbox transfer test-clock.get virtual_time = %q, want %q", got, clockAdvanceTime)
	}
	recurringCancelResp := harness.mustRunJSON(
		"transfer",
		"recurring",
		"cancel",
		"--recurring-transfer-id", recurringTransferID,
	)
	if requireStringField(t, recurringCancelResp, "request_id") == "" {
		t.Fatal("transfer recurring.cancel did not include request_id")
	}

	transferWebhookInbox, err := newLiveWebhookInbox()
	if err != nil {
		t.Logf("skipping sandbox transfer fire-webhook delivery check: %v", err)
	} else {
		fireWebhookResp := harness.mustRunJSON(
			"sandbox",
			"transfer",
			"fire-webhook",
			"--webhook", transferWebhookInbox.url,
		)
		if requireStringField(t, fireWebhookResp, "request_id") == "" {
			t.Fatal("sandbox transfer fire-webhook did not include request_id")
		}
		request, payload := transferWebhookInbox.mustWaitForJSONRequest(t, 45*time.Second, func(body map[string]any) bool {
			webhookType, ok := bodyValue(body, "webhook_type")
			if !ok || webhookType != "TRANSFER" {
				return false
			}
			webhookCode, ok := bodyValue(body, "webhook_code")
			return ok && webhookCode == "TRANSFER_EVENTS_UPDATE"
		})
		if got := requireStringField(t, payload, "webhook_type"); got != "TRANSFER" {
			t.Fatalf("captured transfer webhook webhook_type = %q, want TRANSFER", got)
		}
		if got := requireStringField(t, payload, "webhook_code"); got != "TRANSFER_EVENTS_UPDATE" {
			t.Fatalf("captured transfer webhook webhook_code = %q, want TRANSFER_EVENTS_UPDATE", got)
		}
		plaidVerification := request.headerValue("plaid-verification")
		if plaidVerification == "" {
			t.Fatal("captured transfer webhook did not include plaid-verification header")
		}
		verifyResp := harness.mustRunJSON(
			"webhook",
			"verification-key",
			"get",
			"--plaid-verification", plaidVerification,
		)
		if requireStringField(t, verifyResp, "key", "kid") == "" {
			t.Fatal("transfer webhook verification-key.get did not include key.kid")
		}
	}

	simulateResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"simulate",
		"--transfer-id", transferID,
		"--event-type", "posted",
	)
	if requireStringField(t, simulateResp, "request_id") == "" {
		t.Fatal("sandbox transfer simulate did not include request_id")
	}

	foundPostedEvent := false
	for attempt := 1; attempt <= 10; attempt++ {
		eventListResp := harness.mustRunJSON("transfer", "event", "list", "--transfer-id", transferID, "--count", "25")
		if arrayContainsMapFieldValue(requireArrayField(t, eventListResp, "transfer_events"), "event_type", "posted") {
			foundPostedEvent = true
			break
		}
		if attempt < 10 {
			t.Logf("transfer event.list has not shown posted yet (attempt %d/10); retrying", attempt)
			time.Sleep(2 * time.Second)
		}
	}
	if !foundPostedEvent {
		t.Fatal("transfer event.list did not include a posted event after sandbox transfer simulate")
	}

	settledResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"simulate",
		"--transfer-id", transferID,
		"--event-type", "settled",
	)
	if requireStringField(t, settledResp, "request_id") == "" {
		t.Fatal("sandbox transfer simulate settled did not include request_id")
	}
	foundSettledEvent := false
	for attempt := 1; attempt <= 10; attempt++ {
		eventListResp := harness.mustRunJSON("transfer", "event", "list", "--transfer-id", transferID, "--count", "25")
		if arrayContainsMapFieldValue(requireArrayField(t, eventListResp, "transfer_events"), "event_type", "settled") {
			foundSettledEvent = true
			break
		}
		if attempt < 10 {
			t.Logf("transfer event.list has not shown settled yet (attempt %d/10); retrying", attempt)
			time.Sleep(2 * time.Second)
		}
	}
	if !foundSettledEvent {
		t.Fatal("transfer event.list did not include a settled event after sandbox transfer simulate")
	}

	fundsAvailableResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"simulate",
		"--transfer-id", transferID,
		"--event-type", "funds_available",
	)
	if requireStringField(t, fundsAvailableResp, "request_id") == "" {
		t.Fatal("sandbox transfer simulate funds_available did not include request_id")
	}
	foundFundsAvailableEvent := false
	for attempt := 1; attempt <= 10; attempt++ {
		eventListResp := harness.mustRunJSON("transfer", "event", "list", "--transfer-id", transferID, "--count", "25")
		if arrayContainsMapFieldValue(requireArrayField(t, eventListResp, "transfer_events"), "event_type", "funds_available") {
			foundFundsAvailableEvent = true
			break
		}
		if attempt < 10 {
			t.Logf("transfer event.list has not shown funds_available yet (attempt %d/10); retrying", attempt)
			time.Sleep(2 * time.Second)
		}
	}
	if !foundFundsAvailableEvent {
		t.Fatal("transfer event.list did not include a funds_available event after sandbox transfer simulate")
	}

	sweepSimResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"sweep-simulate",
	)
	if sweepID, ok := bodyValue(sweepSimResp, "sweep", "id"); ok {
		sweepIDString, ok := sweepID.(string)
		if !ok || strings.TrimSpace(sweepIDString) == "" {
			t.Fatalf("sandbox transfer sweep-simulate sweep.id = %#v, want non-empty string", sweepID)
		}
		sweepGetResp := harness.mustRunJSON(
			"transfer",
			"sweep",
			"get",
			"--sweep-id", sweepIDString,
		)
		if got := requireStringField(t, sweepGetResp, "sweep", "id"); got != sweepIDString {
			t.Fatalf("transfer sweep.get sweep.id = %q, want %q", got, sweepIDString)
		}
		sweepListResp := harness.mustRunJSON(
			"transfer",
			"sweep",
			"list",
			"--count", "25",
			"--transfer-id", transferID,
		)
		if !arrayContainsMapField(requireArrayField(t, sweepListResp, "sweeps"), "id", sweepIDString) {
			t.Fatalf("transfer sweep.list did not include sweep %q", sweepIDString)
		}
	} else {
		t.Log("sandbox transfer sweep-simulate returned no sweep; skipping sweep get/list assertions")
	}

	refundSimulateResp := harness.mustRunJSON(
		"sandbox",
		"transfer",
		"refund-simulate",
		"--refund-id", refundSimID,
		"--event-type", "refund.failed",
	)
	if requireStringField(t, refundSimulateResp, "request_id") == "" {
		t.Fatal("sandbox transfer refund-simulate did not include request_id")
	}
	foundRefundFailedEvent := false
	for attempt := 1; attempt <= 10; attempt++ {
		eventListResp := harness.mustRunJSON("transfer", "event", "list", "--transfer-id", transferID, "--count", "25")
		if arrayContainsMapFieldValue(requireArrayField(t, eventListResp, "transfer_events"), "event_type", "refund.failed") {
			foundRefundFailedEvent = true
			break
		}
		if attempt < 10 {
			t.Logf("transfer event.list has not shown refund.failed yet (attempt %d/10); retrying", attempt)
			time.Sleep(2 * time.Second)
		}
	}
	if !foundRefundFailedEvent {
		t.Fatal("transfer event.list did not include refund.failed after sandbox transfer refund-simulate")
	}

	refundCancelResp := harness.mustRunJSON(
		"transfer",
		"refund",
		"create",
		"--transfer-id", transferID,
		"--amount", "0.25",
		"--idempotency-key", fmt.Sprintf("refund-cancel-%d", time.Now().UTC().UnixNano()),
	)
	refundCancelID := requireStringField(t, refundCancelResp, "refund", "id")
	refundCancelAckResp := harness.mustRunJSON(
		"transfer",
		"refund",
		"cancel",
		"--refund-id", refundCancelID,
	)
	if requireStringField(t, refundCancelAckResp, "request_id") == "" {
		t.Fatal("transfer refund.cancel did not include request_id")
	}

	latestTransferEventsResp := harness.mustRunJSON("transfer", "event", "list", "--transfer-id", transferID, "--count", "25")
	latestTransferEventID, ok := maxMapIntField(requireArrayField(t, latestTransferEventsResp, "transfer_events"), "event_id")
	if !ok {
		t.Fatal("transfer event.list did not include any event_id values for event.sync validation")
	}
	afterID := latestTransferEventID - 1
	if afterID < 0 {
		afterID = 0
	}
	eventSyncResp := harness.mustRunJSON("transfer", "event", "sync", "--after-id", fmt.Sprintf("%d", afterID), "--count", "25")
	if !arrayContainsMapField(requireArrayField(t, eventSyncResp, "transfer_events"), "transfer_id", transferID) {
		t.Fatalf("transfer event.sync did not include transfer %q", transferID)
	}
}

func TestLiveSandboxIncomeSuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)
	requireOptionalLiveSuite(t, liveIncomeTestGateEnv, "live Income sandbox tests")

	harness := newLiveSandboxHarness(t)
	harness.initializeAppProfile(cfg)
	incomeItemID := strings.TrimSpace(os.Getenv(liveIncomeItemIDEnv))
	if incomeItemID == "" {
		t.Skipf(
			"set %s to run the Income webhook suite against a pre-created income verification item; see docs/live-test-setup.md",
			liveIncomeItemIDEnv,
		)
	}

	fireResp := harness.mustRunJSON(
		"sandbox",
		"income-fire-webhook",
		"--item-id", incomeItemID,
		"--webhook", "https://example.com/plaid-cli-live-income",
		"--webhook-code", "INCOME_VERIFICATION",
		"--verification-status", "VERIFICATION_STATUS_PROCESSING_COMPLETE",
	)
	if requireStringField(t, fireResp, "request_id") == "" {
		t.Fatal("sandbox income-fire-webhook did not include request_id")
	}
}

func TestLiveSandboxCheckMonitoringSuite(t *testing.T) {
	cfg := loadLiveSandboxConfig(t)
	requireOptionalLiveSuite(t, liveCheckTestGateEnv, "live Plaid Check sandbox tests")

	harness := newLiveSandboxHarness(t)
	cleanupClient := newLiveSandboxClient(t, cfg)
	t.Cleanup(func() {
		harness.cleanup(t, cleanupClient)
	})

	harness.initializeAppProfile(cfg)
	userID := strings.TrimSpace(os.Getenv(liveCheckUserIDEnv))
	itemID := strings.TrimSpace(os.Getenv(liveCheckItemIDEnv))
	if userID == "" || itemID == "" {
		t.Skipf(
			"set %s and %s to run the Plaid Check monitoring suite against a pre-created CRA user and item; see docs/live-test-setup.md",
			liveCheckUserIDEnv,
			liveCheckItemIDEnv,
		)
	}

	subscribeResp := harness.mustRunJSON(
		"check",
		"monitoring",
		"subscribe",
		"--user-id", userID,
		"--item-id", itemID,
		"--webhook", "https://example.com/plaid-cli-live-check",
	)
	subscriptionID := requireStringField(t, subscribeResp, "subscription_id")
	harness.trackSubscription(subscriptionID)

	updateResp := harness.mustRunJSON(
		"sandbox",
		"cra",
		"cashflow-updates-update",
		"--user-id", userID,
		"--webhook-code", "LOW_BALANCE_DETECTED",
	)
	if requireStringField(t, updateResp, "request_id") == "" {
		t.Fatal("sandbox cra cashflow-updates-update did not include request_id")
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

func requireOptionalLiveSuite(t *testing.T, gateEnv, description string) {
	t.Helper()

	if !envTruthy(os.Getenv(gateEnv)) {
		t.Skipf("set %s=1 to run %s", gateEnv, description)
	}
}

func marshalBodyArg(t *testing.T, body map[string]any) string {
	t.Helper()

	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	return string(raw)
}

func newLiveWebhookInbox() (*liveWebhookInbox, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	req, err := http.NewRequest(http.MethodPost, liveWebhookSiteTokenURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create webhook.site token request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create webhook.site inbox: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return nil, fmt.Errorf("create webhook.site inbox: status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var token webhookSiteTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, fmt.Errorf("decode webhook.site token response: %w", err)
	}
	if strings.TrimSpace(token.UUID) == "" {
		return nil, errors.New("webhook.site token response did not include uuid")
	}

	return &liveWebhookInbox{
		client:  client,
		tokenID: token.UUID,
		url:     fmt.Sprintf("%s/%s", liveWebhookSiteBaseURL, token.UUID),
	}, nil
}

func (h *liveSandboxHarness) initializeAppProfile(cfg liveSandboxConfig) {
	h.t.Helper()

	initResp := h.mustRunJSON(
		"init",
		"--env", "sandbox",
		"--client-id", cfg.ClientID,
		"--secret", cfg.Secret,
		"--client-name", cfg.ClientName,
	)
	if got := requireStringField(h.t, initResp, "app_profile", "env"); got != "sandbox" {
		h.t.Fatalf("init app_profile.env = %q, want sandbox", got)
	}
}

func (h *liveSandboxHarness) trackItem(itemID, accessToken string) {
	h.cleanupItems = append(h.cleanupItems, liveSandboxItem{
		ItemID:      itemID,
		AccessToken: accessToken,
	})
}

func (h *liveSandboxHarness) trackUser(userID string) {
	h.cleanupUsers = append(h.cleanupUsers, liveSandboxUser{
		UserID: userID,
	})
}

func (h *liveSandboxHarness) trackSubscription(subscriptionID string) {
	h.cleanupSubscriptions = append(h.cleanupSubscriptions, subscriptionID)
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

func (h *liveSandboxHarness) createUser(prefix string, withIdentity bool) string {
	h.t.Helper()

	clientUserID := fmt.Sprintf("%s-%d", prefix, time.Now().UTC().UnixNano())
	args := []string{"user", "create", "--client-user-id", clientUserID}
	if withIdentity {
		args = append(
			args,
			"--given-name", "Plaid",
			"--family-name", "CLI",
			"--date-of-birth", "1990-01-01",
			"--email", clientUserID+"@example.com",
			"--phone-number", "+14155550199",
			"--street-1", "123 Main St",
			"--city", "San Francisco",
			"--region", "CA",
			"--country", "US",
			"--postal-code", "94105",
			"--id-number", "1234",
			"--id-number-type", "us_ssn_last_4",
		)
	}

	resp := h.mustRunJSON(args...)
	userID := requireStringField(h.t, resp, "user_id")
	h.trackUser(userID)
	return userID
}

func (h *liveSandboxHarness) createSandboxItem(cfg liveSandboxConfig, products []string, extraArgs ...string) liveSandboxItem {
	h.t.Helper()

	item, err := h.tryCreateSandboxItem(cfg, products, extraArgs...)
	if err != nil {
		h.t.Fatal(err)
	}
	return item
}

func (h *liveSandboxHarness) tryCreateSandboxItem(cfg liveSandboxConfig, products []string, extraArgs ...string) (liveSandboxItem, error) {
	h.t.Helper()

	createArgs := []string{
		"sandbox",
		"public-token-create",
		"--institution-id", cfg.InstitutionID,
	}
	createArgs = append(createArgs, extraArgs...)
	for _, product := range products {
		createArgs = append(createArgs, "--product", product)
	}
	publicTokenResp, err := h.runJSON(createArgs...)
	if err != nil {
		return liveSandboxItem{}, err
	}
	publicToken := requireStringField(h.t, publicTokenResp, "public_token")

	exchangeArgs := []string{
		"item",
		"public-token-exchange",
		"--public-token", publicToken,
	}
	for _, product := range products {
		exchangeArgs = append(exchangeArgs, "--product", product)
	}
	exchangeResp, err := h.runJSON(exchangeArgs...)
	if err != nil {
		return liveSandboxItem{}, err
	}
	item := liveSandboxItem{
		ItemID:      requireStringField(h.t, exchangeResp, "item_id"),
		AccessToken: requireStringField(h.t, exchangeResp, "access_token"),
	}
	h.trackItem(item.ItemID, item.AccessToken)
	return item, nil
}

func (h *liveSandboxHarness) requireItemAccountID(itemID string) string {
	h.t.Helper()

	record, err := state.New(h.stateDir).LoadItem(itemID)
	if err != nil {
		h.t.Fatalf("LoadItem(%q) error = %v", itemID, err)
	}
	if len(record.Accounts) == 0 {
		h.t.Fatalf("saved item %q did not include any accounts", itemID)
	}
	if strings.TrimSpace(record.Accounts[0].AccountID) == "" {
		h.t.Fatalf("saved item %q has an empty first account_id", itemID)
	}
	return record.Accounts[0].AccountID
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
		if errors.As(err, &apiErr) && (apiErr.ErrorCode == "PRODUCT_NOT_READY" || apiErr.ErrorType == "RATE_LIMIT_EXCEEDED") && attempt < attempts {
			h.t.Logf("%s returned transient Plaid error %s/%s (attempt %d/%d); retrying in %s", strings.Join(args, " "), apiErr.ErrorType, apiErr.ErrorCode, attempt, attempts, delay)
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

	for i := len(h.cleanupSubscriptions) - 1; i >= 0; i-- {
		subscriptionID := h.cleanupSubscriptions[i]
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		_, err := client.Call(ctx, "/cra/monitoring_insights/unsubscribe", map[string]any{
			"subscription_id": subscriptionID,
		})
		cancel()
		if err != nil {
			t.Logf("cleanup: unsubscribe monitoring subscription %s: %v", subscriptionID, err)
		}
	}

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

	for i := len(h.cleanupUsers) - 1; i >= 0; i-- {
		user := h.cleanupUsers[i]
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		_, err := client.Call(ctx, "/user/remove", map[string]any{
			"user_id": user.UserID,
		})
		cancel()
		if err != nil {
			t.Logf("cleanup: remove sandbox user %s: %v", user.UserID, err)
		}
	}
}

type liveWebhookInbox struct {
	client  *http.Client
	tokenID string
	url     string
}

func (i *liveWebhookInbox) mustWaitForJSONRequest(t *testing.T, timeout time.Duration, match func(map[string]any) bool) (webhookSiteRequest, map[string]any) {
	t.Helper()

	deadline := time.Now().Add(timeout)
	for attempt := 1; time.Now().Before(deadline); attempt++ {
		requests, err := i.listRequests()
		if err == nil {
			for _, request := range requests {
				var payload map[string]any
				if err := json.Unmarshal([]byte(request.Content), &payload); err != nil {
					continue
				}
				if match == nil || match(payload) {
					return request, payload
				}
			}
		}

		if attempt < 20 {
			time.Sleep(2 * time.Second)
		}
	}

	t.Fatalf("timed out waiting for webhook delivery to %s", i.url)
	return webhookSiteRequest{}, nil
}

func (i *liveWebhookInbox) listRequests() ([]webhookSiteRequest, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/token/%s/requests?sorting=newest", liveWebhookSiteBaseURL, i.tokenID), nil)
	if err != nil {
		return nil, fmt.Errorf("create webhook.site list request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("list webhook.site requests: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return nil, fmt.Errorf("list webhook.site requests: status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var payload webhookSiteRequestsResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("decode webhook.site requests response: %w", err)
	}
	return payload.Data, nil
}

func (r webhookSiteRequest) headerValue(name string) string {
	for key, values := range r.Headers {
		if strings.EqualFold(key, name) && len(values) > 0 {
			return values[0]
		}
	}
	return ""
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

func arrayContainsMapField(values []any, field, want string) bool {
	for _, raw := range values {
		entry, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		got, ok := entry[field].(string)
		if ok && got == want {
			return true
		}
	}
	return false
}

func arrayContainsMapFieldValue(values []any, field, want string) bool {
	return arrayContainsMapField(values, field, want)
}

func maxMapIntField(values []any, field string) (int, bool) {
	found := false
	maxValue := 0
	for _, raw := range values {
		entry, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		rawValue, ok := entry[field]
		if !ok {
			continue
		}

		var value int
		switch typed := rawValue.(type) {
		case float64:
			value = int(typed)
		case float32:
			value = int(typed)
		case int:
			value = typed
		case int64:
			value = int(typed)
		case json.Number:
			parsed, err := typed.Int64()
			if err != nil {
				continue
			}
			value = int(parsed)
		default:
			continue
		}

		if !found || value > maxValue {
			maxValue = value
			found = true
		}
	}
	return maxValue, found
}

func isTransferUnavailableError(err error) bool {
	var apiErr *plaid.APIError
	if !errors.As(err, &apiErr) {
		return false
	}

	switch apiErr.ErrorCode {
	case "PRODUCT_NOT_ENABLED", "PRODUCTS_NOT_SUPPORTED", "INVALID_PRODUCT", "UNAUTHORIZED_ROUTE_ACCESS":
		return true
	}

	message := strings.ToLower(strings.TrimSpace(strings.Join([]string{
		apiErr.ErrorType,
		apiErr.ErrorCode,
		apiErr.ErrorMessage,
		apiErr.DisplayMessage,
	}, " ")))
	if !strings.Contains(message, "transfer") {
		return false
	}
	return strings.Contains(message, "not enabled") ||
		strings.Contains(message, "not supported") ||
		strings.Contains(message, "not available") ||
		strings.Contains(message, "not configured") ||
		strings.Contains(message, "request access") ||
		strings.Contains(message, "product")
}

func isProductUnavailableError(err error) bool {
	var apiErr *plaid.APIError
	if !errors.As(err, &apiErr) {
		return false
	}

	switch apiErr.ErrorCode {
	case "PRODUCT_NOT_ENABLED", "PRODUCTS_NOT_SUPPORTED", "INVALID_PRODUCT", "UNAUTHORIZED_ROUTE_ACCESS", "SANDBOX_PRODUCT_NOT_ENABLED":
		return true
	}

	message := strings.ToLower(strings.TrimSpace(strings.Join([]string{
		apiErr.ErrorType,
		apiErr.ErrorCode,
		apiErr.ErrorMessage,
		apiErr.DisplayMessage,
	}, " ")))

	return strings.Contains(message, "not enabled") ||
		strings.Contains(message, "not supported") ||
		strings.Contains(message, "not available") ||
		strings.Contains(message, "not configured") ||
		strings.Contains(message, "request access") ||
		(strings.Contains(message, "product") && strings.Contains(message, "sandbox"))
}

func skipUnavailableLiveProduct(t *testing.T, label string, err error) {
	t.Helper()

	if isProductUnavailableError(err) {
		t.Skipf("skipping %s live coverage: %v", label, err)
	}
	t.Fatal(err)
}

func envTruthy(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "y", "on":
		return true
	default:
		return false
	}
}
