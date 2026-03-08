set shell := ["bash", "-euo", "pipefail", "-c"]

# Build the plaid binary into the repo root.
build:
	go build -o plaid .

# Run the standard Go test suite.
test:
	go test ./...

# Install the CLI binary to ~/.local/bin.
install: build
	mkdir -p "$HOME/.local/bin"
	install -m 0755 plaid "$HOME/.local/bin/plaid"

# Run live Plaid sandbox smoke tests. Requires sandbox creds in env.
test-live:
	PLAID_RUN_LIVE_TESTS=1 go test -count=1 -run '^TestLive' ./...

# Run the opt-in Income sandbox smoke test.
test-live-income:
	PLAID_RUN_LIVE_TESTS=1 PLAID_RUN_LIVE_INCOME_TESTS=1 go test -count=1 -run '^TestLiveSandboxIncomeSuite$' ./...

# Run the opt-in Plaid Check / Cash Flow Updates sandbox smoke test.
test-live-check:
	PLAID_RUN_LIVE_TESTS=1 PLAID_RUN_LIVE_CHECK_TESTS=1 go test -count=1 -run '^TestLiveSandboxCheckMonitoringSuite$' ./...
