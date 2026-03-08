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
plaid transfer create \
  --authorization-id YOUR_AUTHORIZATION_ID \
  --description "payment" \
  --body @transfer.json
```

Create a raw Link token for an agent-managed Link flow:

```bash
plaid link token-create \
  --product auth \
  --product identity \
  --client-user-id local-owner
```

Exchange a Link public token and save the resulting Item locally:

```bash
plaid item public-token-exchange \
  --public-token YOUR_PUBLIC_TOKEN \
  --product auth
```

Evaluate ACH return risk with Signal:

```bash
plaid signal evaluate \
  --item YOUR_ITEM_ID \
  --account-id YOUR_ACCOUNT_ID \
  --client-transaction-id txn_123 \
  --amount 102.05
```

Create a Plaid Check report for a user:

```bash
plaid check report create \
  --user-id YOUR_USER_ID \
  --webhook https://example.com/webhooks/plaid-check \
  --days-requested 365
```

Download the latest bank income PDF to disk:

```bash
plaid income bank-income-pdf-get \
  --user-token YOUR_USER_TOKEN \
  --out reports/bank-income.pdf
```

Create a wallet transfer using typed flags plus JSON for less common fields:

```bash
plaid wallet transaction execute \
  --wallet-id YOUR_WALLET_ID \
  --idempotency-key txn_123 \
  --counterparty-name "Jane Doe" \
  --amount-currency GBP \
  --amount-value 10.50 \
  --reference "PAYOUT-123" \
  --body @wallet-transaction.json
```

## Local Docs Snapshot

Plaid's docs are mirrored under [docs/plaid/](docs/plaid/) for implementation reference.
