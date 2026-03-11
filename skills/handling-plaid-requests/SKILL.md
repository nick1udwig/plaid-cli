---
name: handling-plaid-requests
description: Handles end-user Plaid requests by running the repo's `plaid` CLI and interpreting the results. Use when the user asks about linked accounts, balances, routing details, transactions, spending summaries, recurring payments, institution matches, risk checks, or money movement, including requests like "how much do I have in Chase?" or "summarize my last 3 months of spending by category."
---

# Handling Plaid Requests

Use this skill when the agent should satisfy a user's Plaid-related request by running `plaid` directly and answering in prose. Do not turn ordinary requests into "here are some commands to run" unless the user explicitly asks for manual instructions.

## Core Rule

- Run the CLI yourself for normal user requests.
- Ask the human to do work by hand only for one-time setup and browser-based account linking.
- Use `plaid` if it is on `PATH`; otherwise use `go run .`.
- Prefer the existing local state in `~/.plaid-cli`. Use `--state-dir` only for intentionally isolated scratch work.

## Setup Gate

If Plaid is not configured or no Items are linked, stop and route the human through setup:

- `plaid init` for one-time app credentials
- `plaid link connect` for the initial Hosted Link browser flow
- `docs/getting-started.md` for exact setup instructions

The agent may invoke `plaid link connect`, but the human still has to complete the browser interaction.

## Workflow

1. Start with local state:
   - `plaid item list`
2. Resolve the target Item or account:
   - If the user names an institution, inspect the saved Items and match institution metadata.
   - If the user names an account type, use `plaid account get --item ITEM_ID` and choose the matching account.
3. Call the narrowest Plaid command that answers the question.
4. Post-process the JSON locally when the user wants summaries, grouping, or comparisons.
5. Answer in prose with exact dates, amounts, and assumptions.

## Safety

- Execute read-only requests directly.
- For write, destructive, or move-money requests, require explicit user intent and clear target details before sending anything.
- Do not infer the source account, destination, amount, or recipient for a transfer.
- Do not expose access tokens, account numbers, or routing numbers more broadly than the user asked for.
- If multiple Items or accounts plausibly match, ask a short clarifying question.

## Command Mapping

### Balance and account questions

Use this for requests like:

- "How much do I have in my Chase checking account?"
- "What accounts do I have linked?"
- "Which account is my savings account?"

Run:

```bash
plaid item list
plaid account get --item ITEM_ID
plaid balance get --item ITEM_ID --account-id ACCOUNT_ID
```

Use `plaid item get --item ITEM_ID` when the institution metadata or Item status matters.

### Routing and account numbers

Use this for requests like:

- "What routing number is attached to my checking account?"
- "Show me the account and routing details for this Item."

Run:

```bash
plaid auth get --item ITEM_ID
```

Only reveal the fields the user actually asked for.

### Spending history and transaction summaries

Use this for requests like:

- "Create a spending history from the last 3 months and summarize by categories."
- "Show my biggest purchases this month."
- "What did I spend on groceries recently?"

Run:

```bash
plaid transactions get --item ITEM_ID --start-date START_DATE --end-date END_DATE
```

Or, if the user wants incremental updates or the newest delta:

```bash
plaid transactions sync --item ITEM_ID
```

Then group or summarize locally. Prefer Plaid's personal finance category fields when present. In the final answer, name the exact start and end dates you used.

### Recurring charges and subscriptions

Use this for requests like:

- "What subscriptions hit every month?"
- "Show my recurring bills."

Run:

```bash
plaid transactions recurring-get --item ITEM_ID
```

### Identity, liabilities, investments, statements, and institution lookups

Map the user's request directly to the matching product command and summarize the result:

- identity -> `plaid identity get`
- liabilities -> `plaid liabilities get`
- investments -> `plaid investments holdings-get`, `plaid investments transactions-get`
- statements -> `plaid statements list`, `plaid statements download`
- institution discovery -> `plaid institution get`, `plaid institution search`

Use `plaid <command> --help` or the mirrored docs in `docs/plaid/` only when the exact flags are unclear.

### Risk checks before moving money

Use this for requests like:

- "Can this account receive an ACH debit?"
- "Check the risk before sending money."

Run:

```bash
plaid transfer capabilities get --item ITEM_ID --account-id ACCOUNT_ID
plaid signal evaluate --item ITEM_ID --account-id ACCOUNT_ID --client-transaction-id CLIENT_TRANSACTION_ID --amount AMOUNT
```

Explain capability or product gating plainly if the account or app is not enabled.

### Transfers and move-money requests

Use this only after the user has clearly confirmed the amount, source account, and destination context.

Typical sequence:

```bash
plaid transfer capabilities get --item ITEM_ID --account-id ACCOUNT_ID
plaid transfer authorization create --item ITEM_ID --account-id ACCOUNT_ID --type debit --network ach --ach-class ppd --amount AMOUNT --legal-name LEGAL_NAME
plaid transfer create --item ITEM_ID --account-id ACCOUNT_ID --authorization-id AUTHORIZATION_ID --amount AMOUNT --description DESCRIPTION
```

For sandbox validation or event inspection:

```bash
plaid sandbox transfer simulate --transfer-id TRANSFER_ID --event-type posted
plaid transfer event list --transfer-id TRANSFER_ID
```

If Transfer is not enabled, say so directly instead of improvising a workaround.

## Example User Requests

### "How much do I have in my Chase account?"

1. `plaid item list`
2. Find the Chase Item.
3. `plaid account get --item ITEM_ID`
4. Choose the account that matches the user's wording.
5. `plaid balance get --item ITEM_ID --account-id ACCOUNT_ID`
6. Answer with the account name and the current and available balances.

### "Create a spending history from the last 3 months and summarize by categories."

1. Resolve the target Item or account.
2. Compute the date window at runtime.
3. `plaid transactions get --item ITEM_ID --start-date START_DATE --end-date END_DATE`
4. Aggregate locally by category.
5. Summarize totals, top categories, and notable merchants. Include the exact dates used.

### "What subscriptions recur each month?"

1. Resolve the relevant Item.
2. `plaid transactions recurring-get --item ITEM_ID`
3. Summarize likely recurring outflows and inflows, with frequency where available.

## Output Style

- Prefer a short prose answer over raw JSON.
- Include exact figures and date ranges.
- State assumptions and ambiguities.
- Offer a concise follow-up only when it is genuinely useful.

## Useful Paths

- `docs/getting-started.md`
- `docs/live-test-setup.md`
- `docs/plaid/`
