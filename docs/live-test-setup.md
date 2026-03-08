# Live Sandbox Test Setup

This repo has two layers of live Plaid sandbox coverage:

- `just test-live`: the default live suite
- opt-in product suites: processor, Payment Initiation, Income, and Plaid Check / Cash Flow Updates

The default live suite only needs valid sandbox app credentials. The opt-in suites need extra product access or special setup.

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

No extra Dashboard product enablement is required beyond normal sandbox API access.

## Processor Token Suite

Run:

```bash
just test-live-processor
```

This exercises `sandbox processor-token-create`.

Notes:

- No extra Dashboard product enablement is usually required.
- Plaid only returns a `processor_token`, not an `access_token`, so the test cannot remove the created sandbox Item afterward.
- Keep this suite opt-in for that reason.

## Payment Initiation Suite

Run:

```bash
just test-live-payment
```

This exercises:

- `payment-initiation recipient create`
- `payment-initiation payment create`
- `sandbox payment-simulate`

Dashboard setup:

1. Open `Team Settings -> Products` in Plaid Dashboard.
2. Enable or request `Payment Initiation`.
3. If the product is unavailable, file an access request from the Dashboard.

Notes:

- Plaid’s Payment Initiation docs say the product is enabled in Sandbox by default, but some teams still have account-level access constraints.
- If you need to adjust Payment Initiation sandbox behavior, Plaid documents additional controls on the Dashboard Sandbox page.

Official references:

- https://plaid.com/docs/payment-initiation/add-to-app/
- https://plaid.com/docs/dashboard/team-management/teams/

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
