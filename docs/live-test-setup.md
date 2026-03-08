# Live Sandbox Test Setup

This repo has two layers of live Plaid sandbox coverage:

- `just test-live`: the default live suite
- opt-in manual-fixture suites: Income and Plaid Check / Cash Flow Updates

The default live suite only needs valid sandbox app credentials. The opt-in suites need extra fixture env vars.

## Dashboard Access

Before enabling product-gated suites, make sure the human doing setup can manage product access in Plaid Dashboard.

- Open Plaid Dashboard.
- Go to `Team Settings -> Products`.
- If you do not see the Products page or cannot edit it, your Dashboard role likely needs the `Team Management` permission.
- If a product is not available for your team, file an access request from the Dashboard or contact your Plaid account team.

Official references:

- https://plaid.com/docs/dashboard/team-management/teams/
- https://plaid.com/docs/sandbox/

## Base Suite

Run:

```bash
just test-live
```

This covers the core headless sandbox flow, plus:

- user-linked `sandbox user-reset-login`
- dynamic `sandbox transactions-create`
- `transactions refresh`
- `sandbox processor-token-create`
- processor `account/auth/balance/identity/transactions/*` reads
- webhook delivery plus `webhook verification-key get`
- `payment-initiation recipient create`
- `payment-initiation recipient get/list`
- `payment-initiation payment create`
- `payment-initiation payment get/list`
- `payment-initiation consent create/get`
- `sandbox payment-simulate`
- Transfer capabilities, authorization, create/get/list, refund create/get/cancel, recurring create/get/list/cancel, sandbox test-clock create/get/list/advance, sandbox simulate, and event reads when Transfer is enabled

Notes:

- `sandbox processor-token-create` returns only a `processor_token`, so the Item it creates cannot be cleaned up through this CLI.
- Webhook delivery coverage uses a disposable `webhook.site` inbox. If `webhook.site` is unavailable, that one test skips and the rest of `just test-live` still runs.
- The Payment Initiation path runs as part of `just test-live`. On this sandbox account it worked without extra env vars. If it fails on another account with a product-access error, enable or request `Payment Initiation` under `Team Settings -> Products`.
- The Transfer path runs as part of `just test-live` and self-skips when the app lacks Transfer access. If you want that coverage, enable or request `Transfer` under `Team Settings -> Products`.

## Income Suite

Run:

```bash
just test-live-income
```

This exercises:

- `sandbox income-fire-webhook`

Dashboard setup:

1. Open `Team Settings -> Products`.
2. Enable or request `Income Verification`.
3. If the product is not available, file an access request from the Dashboard.

Test fixture details:

1. Create a sandbox Income Verification item manually in a human-run flow after the product is enabled.
2. Export the resulting `item_id` as `PLAID_LIVE_INCOME_ITEM_ID`.
3. Run `just test-live-income`.

Observed constraint:

- On March 8, 2026, the current sandbox `/sandbox/public_token/create` flow returned `user token required` when bootstrapping `income_verification` with a new-flow `user_id`.
- Because this Plaid account cannot create legacy `user_token` users, the suite does not attempt a fully headless Income bootstrap.

Sandbox data guidance:

- For realistic income data, use the special sandbox username `user_bank_income` with password `{}`.
- Use a non-OAuth sandbox institution such as `ins_109508`.
- The Sandbox page in Dashboard is where Plaid documents custom Sandbox user tooling.

Official references:

- https://plaid.com/docs/income/bank-income/
- https://plaid.com/docs/dashboard/team-management/teams/
- https://plaid.com/docs/sandbox/

## Plaid Check / Cash Flow Updates Suite

Run:

```bash
just test-live-check
```

This exercises:

- `check monitoring subscribe`
- `sandbox cra cashflow-updates-update`

Dashboard setup:

1. Open `Team Settings -> Products`.
2. Enable or request Plaid Check access.
3. Make sure the Cash Flow Updates / monitoring product family is included for your team.
4. If the product is unavailable, file an access request from the Dashboard.

Test fixture details:

1. Create a Plaid Check user and a linked CRA item manually in a human-run flow after the product is enabled.
2. Export the user as `PLAID_LIVE_CHECK_USER_ID`.
3. Export the linked CRA item as `PLAID_LIVE_CHECK_ITEM_ID`.
4. Run `just test-live-check`.

Observed constraints:

- On March 8, 2026, `cra_monitoring` was rejected in `sandbox public-token-create initial_products` for this account.
- On the same date, a `cra_base_report` public token created through `sandbox public-token-create` could not be exchanged through `/item/public_token/exchange`.
- Because of those sandbox bootstrap constraints, the suite uses pre-created CRA fixtures instead of a fully headless setup.

Sandbox data guidance:

- For more realistic Plaid Check cashflow behavior, use the special sandbox username `user_bank_income` with password `{}` when creating the manual fixture.
- Use a non-OAuth sandbox institution such as `ins_109508`.

Official references:

- https://plaid.com/docs/check/
- https://plaid.com/docs/check/add-to-app/
- https://plaid.com/docs/dashboard/team-management/teams/
- https://plaid.com/docs/sandbox/

## Failure Modes

If an opt-in suite fails immediately with a Plaid product-access error, the usual cause is missing Dashboard product enablement rather than a CLI bug.

Common signals:

- `PRODUCT_NOT_ENABLED`
- `PRODUCTS_NOT_SUPPORTED`
- `INVALID_PRODUCT`

When that happens:

1. Verify the product is enabled under `Team Settings -> Products`.
2. Verify the app profile stored in `~/.plaid-cli/app-profile.json` is using `sandbox`.
3. Re-run the suite with the matching `just test-live-*` command.
