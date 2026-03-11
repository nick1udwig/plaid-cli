# Plaid CLI Use Cases

Use these playbooks when the user wants concrete ways to exercise the CLI by hand. All commands are one-liners.

## Sandbox bootstrap

Use this when the user wants a safe scratch environment without touching `~/.plaid-cli`.

```bash
set -a; source .env.sandbox; set +a
plaid --state-dir /tmp/plaid-play init --env sandbox --client-id "$PLAID_SANDBOX_CLIENT_ID" --secret "$PLAID_SANDBOX_SECRET" --client-name "Plaid CLI Sandbox"
PUBLIC_TOKEN=$(plaid --state-dir /tmp/plaid-play sandbox public-token-create --institution-id "${PLAID_SANDBOX_INSTITUTION_ID:-ins_109508}" --product auth --product transactions | jq -r .public_token)
ITEM_ID=$(plaid --state-dir /tmp/plaid-play item public-token-exchange --public-token "$PUBLIC_TOKEN" --product auth --product transactions | jq -r .item_id)
ACCOUNT_ID=$(plaid --state-dir /tmp/plaid-play account get --item "$ITEM_ID" | jq -r '.accounts[0].account_id')
```

## Read account and routing data

Use this when the user wants to inspect a linked Item and verify the Auth product works.

```bash
plaid --state-dir /tmp/plaid-play item list
plaid --state-dir /tmp/plaid-play item get --item "$ITEM_ID"
plaid --state-dir /tmp/plaid-play account get --item "$ITEM_ID"
plaid --state-dir /tmp/plaid-play auth get --item "$ITEM_ID"
plaid --state-dir /tmp/plaid-play balance get --item "$ITEM_ID"
```

## Read and change transactions

Use this when the user wants an agent-style read loop plus a sandbox mutation that creates a visible delta.

```bash
plaid --state-dir /tmp/plaid-play transactions sync --item "$ITEM_ID"
plaid --state-dir /tmp/plaid-play transactions get --item "$ITEM_ID" --start-date 2025-01-01 --end-date 2026-12-31
plaid --state-dir /tmp/plaid-play transactions recurring-get --item "$ITEM_ID"
plaid --state-dir /tmp/plaid-play sandbox transactions-create --item "$ITEM_ID" --account-id "$ACCOUNT_ID" --transaction-id demo_txn_1 --date-transacted 2026-03-01 --date-posted 2026-03-01 --amount 12.34 --description "manual sandbox test" --currency USD
plaid --state-dir /tmp/plaid-play transactions sync --item "$ITEM_ID"
```

For a more mutation-friendly transactions fixture, create the Item with `--override-username user_transactions_dynamic`.

## Hosted Link browser flow

Use this when the user specifically wants to try Plaid's browser UX instead of a headless sandbox token flow.

```bash
plaid --state-dir /tmp/plaid-play link connect --product auth --product transactions
```

## ACH risk check with Signal

Use this when the user wants an agent to evaluate account risk before moving money.

```bash
plaid --state-dir /tmp/plaid-play signal evaluate --item "$ITEM_ID" --account-id "$ACCOUNT_ID" --client-transaction-id demo_signal_1 --amount 12.34
```

## Webhook smoke test

Use this when the user wants to verify that sandbox webhook triggers work. A disposable `webhook.site` URL is the easiest receiver.

```bash
WEBHOOK_URL="https://webhook.site/YOUR-UUID"
PUBLIC_TOKEN=$(plaid --state-dir /tmp/plaid-play sandbox public-token-create --institution-id "${PLAID_SANDBOX_INSTITUTION_ID:-ins_109508}" --product transactions --webhook "$WEBHOOK_URL" | jq -r .public_token)
ITEM_ID=$(plaid --state-dir /tmp/plaid-play item public-token-exchange --public-token "$PUBLIC_TOKEN" --product transactions | jq -r .item_id)
plaid --state-dir /tmp/plaid-play sandbox item-fire-webhook --item "$ITEM_ID" --webhook-code SYNC_UPDATES_AVAILABLE
```

If the user wants a full receiver plus signature-verification flow, point them to `just test-live` and `docs/live-test-setup.md`.

## Transfer and move-money flows

Use this only when the user has Transfer enabled and wants live balance-changing operations or sandbox transfer simulation.

```bash
plaid --state-dir /tmp/plaid-play transfer capabilities get --item "$ITEM_ID" --account-id "$ACCOUNT_ID"
plaid --state-dir /tmp/plaid-play transfer create --print-request-template
plaid --state-dir /tmp/plaid-play transfer authorization create --help
plaid --state-dir /tmp/plaid-play transfer create --help
```

Keep move-money flows separate from read-only flows. Mention that transfer setup is product-gated and may require additional dashboard configuration.

## Common user requests this skill should answer

- "Show me how to connect a sandbox bank account and read the balance."
- "Give me a safe sandbox flow for transactions and webhooks."
- "What can an agent do with this Plaid CLI?"
- "How would I use this CLI to inspect accounts before a transfer?"
- "What is the quickest way to try the browser-based Link flow?"
