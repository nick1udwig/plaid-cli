# plaid-cli

Single-owner, agent-friendly CLI for the Plaid API.

## Setup

Humans must complete one-time setup before an agent can use the CLI.

See [docs/getting-started.md](docs/getting-started.md) for:

- how to get the required Plaid app credentials
- how to run `plaid init`
- how to connect a bank with Hosted Link

## Examples

Print a minimal request template for a transfer authorization:

```bash
plaid transfer authorization create --print-request-template
```

Check whether a linked account supports Transfer capabilities:

```bash
plaid transfer capabilities get --item YOUR_ITEM_ID --account-id YOUR_ACCOUNT_ID
```

Create a transfer using typed flags plus JSON as an escape hatch for less common fields:

```bash
plaid transfer create --authorization-id YOUR_AUTHORIZATION_ID --description "payment" --body @transfer.json
```

Create a raw Link token for an agent-managed Link flow:

```bash
plaid link token-create --product auth --product identity --client-user-id local-owner
```

Exchange a Link public token and save the resulting Item locally:

```bash
plaid item public-token-exchange --public-token YOUR_PUBLIC_TOKEN --product auth
```

Evaluate ACH return risk with Signal:

```bash
plaid signal evaluate --item YOUR_ITEM_ID --account-id YOUR_ACCOUNT_ID --client-transaction-id txn_123 --amount 102.05
```

Create a Plaid Check report for a user:

```bash
plaid check report create --user-id YOUR_USER_ID --webhook https://example.com/webhooks/plaid-check --days-requested 365
```

Download the latest bank income PDF to disk:

```bash
plaid income bank-income-pdf-get --user-token YOUR_USER_TOKEN --out reports/bank-income.pdf
```

Create a wallet transfer using typed flags plus JSON for less common fields:

```bash
plaid wallet transaction execute --wallet-id YOUR_WALLET_ID --idempotency-key txn_123 --counterparty-name "Jane Doe" --amount-currency GBP --amount-value 10.50 --reference "PAYOUT-123" --body @wallet-transaction.json
```

## Local Docs Snapshot

Plaid's docs are mirrored under [docs/plaid/](docs/plaid/) for implementation reference.

## Testing

Run the standard test suite:

```bash
just test
```

Run the live Plaid sandbox smoke suite:

```bash
just test-live
```

Run the manual-fixture live suites:

```bash
just test-live-income
```

```bash
just test-live-check
```

The live suite is gated behind `PLAID_RUN_LIVE_TESTS=1`, uses a temporary state directory instead of `~/.plaid-cli`, and removes the sandbox Items it can clean up. It reads sandbox creds from `PLAID_SANDBOX_CLIENT_ID` / `PLAID_SANDBOX_SECRET`, falls back to `PLAID_CLIENT_ID` / `PLAID_SECRET`, and finally falls back to a saved sandbox app profile in `~/.plaid-cli/app-profile.json`.

The `sandbox item-set-verification-status` smoke path is opt-in because Plaid requires a pre-created automated micro-deposit Item for that endpoint. To include it in `just test-live`, set `PLAID_LIVE_AUTOMATED_MICRODEPOSIT_ACCESS_TOKEN` and `PLAID_LIVE_AUTOMATED_MICRODEPOSIT_ACCOUNT_ID`.

`just test-live` also includes dynamic `transactions refresh`, processor downstream reads, webhook delivery plus `webhook verification-key get`, broader Payment Initiation coverage (`recipient get/list`, `payment get/list`, `consent create/get`), and deeper Transfer coverage (`refund create/get/cancel`, recurring transfers, and sandbox test clocks). The webhook delivery test uses a disposable `webhook.site` inbox and will skip if that service is unavailable. The Transfer path self-skips when the sandbox app does not have Transfer access.

The Income suite requires `PLAID_LIVE_INCOME_ITEM_ID`. The Plaid Check suite requires `PLAID_LIVE_CHECK_USER_ID` and `PLAID_LIVE_CHECK_ITEM_ID`.

See [docs/live-test-setup.md](docs/live-test-setup.md) for the Dashboard setup and manual fixture requirements for the opt-in live suites.
