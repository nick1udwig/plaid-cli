---
title: "Income - Bank Income | Plaid Docs"
source_url: "https://plaid.com/docs/income/bank-income/"
scraped_at: "2026-03-07T22:04:58+00:00"
---

# Bank Income

#### Learn about Bank Income features and implementation

Get started with Bank Income

[API Reference](/docs/api/products/income/)[Quickstart](https://github.com/plaid/income-sample)[Income Verification Demo](https://plaid.coastdemo.com/share/66fb0a180582208ffa82103e?zoom=100)

#### Overview

Bank Income allows you to instantly retrieve net income information from a user-connected bank account, supporting both irregular or gig income and W-2 income. Data available includes a breakdown of income streams to the account, as well as recent and historical income.

Prefer to learn by watching? Get an overview of how Income works in just 3 minutes!

For all new integrations supporting end users based in the US, Plaid recommends the use of [Consumer Report](/docs/check/) instead of Bank Income. Consumer Report's income solution provides FCRA compliant insights and bundles income calculations (historical average income, forecasted income, predicted next payment date) in a single report. For more information, see [Plaid Check Consumer Report](/docs/check/).

#### No-code Income integration with the Credit Dashboard

Bank Income can be used as part of a no-code integration flow. This integration mode is available only when using Income alongside Assets. For more details, see the [Assets no-code integration guide](/docs/assets/#no-code-integration-with-the-credit-dashboard).

#### Integration process

New customers integrating with Bank Income after December 10, 2025 must request access to user tokens by contacting Sales or their Account Manager, or filing a Support ticket via the Dashboard. If access is not requested, [`/user/create`](/docs/api/users/#usercreate) will not generate a user token, and will generate only a `user_id` instead. Later in Q1 2026, Plaid will add an integration option for Income that does not require user tokens. For details, see [New User APIs](/docs/api/users/user-apis/).

1. Call [`/user/create`](/docs/api/users/#usercreate) to create a `user_token` that will represent the end user interacting with your application. This step will only need to be done once per end user. If you are using multiple Income types, do not repeat this step when switching to a different Income type. If you do not receive a `user_token` when calling [`/user/create`](/docs/api/users/#usercreate), contact your Account Manager or file a Support ticket to request access. For more details, see [New User APIs](/docs/api/users/user-apis/).
2. Call [`/link/token/create`](/docs/api/link/#linktokencreate). In addition to the required parameters, you will need to provide the following:
   - For `user_token`, provide the `user_token` from [`/user/create`](/docs/api/users/#usercreate).
   - For `products`, use `["income_verification"]`. You can also specify additional products. For more details, see [Using Bank Income with other products](/docs/income/bank-income/#using-bank-income-with-other-products).
   - For `income_verification.income_source_types`, use `bank`.
   - Set `income_verification.bank_income.days_requested` to the desired number of days.
   - Provide a `webhook` URI with the endpoint where you will receive Plaid webhooks.
3. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see the [Link documentation](/docs/link/).
4. Open Link in your web or mobile client and listen to the `onSuccess` and `onExit` callbacks, which will fire once the user has finished or exited the Link session.
5. If you are using other Plaid products such as Auth or Balance alongside Bank Income, call [`/link/token/get`](/docs/api/link/#linktokenget) and make sure to capture each `public_token` from the `results.item_add_results` array. Exchange each `public_token` for an `access_token` using [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange). For more details on token exchange, see the [Token exchange flow](/docs/api/items/#token-exchange-flow).
6. To retrieve data, call [`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) with the `user_token`.

#### Multi-Item sessions

Many users get income deposited into multiple institutions. To help capture a user’s full income profile, you can allow your users to link multiple accounts within the same link session on web integrations.

![Select bank, enter credentials, choose income transactions, and review income in Plaid's bank income flow.](/assets/img/docs/income/bank_income_multi_item.png)

Bank Income Multi Item Link flow (some panes excluded)

To enable this flow, see [Multi-Item Link](/docs/link/multi-item-link/). The previous Income-specific flow, which used the `income_verification.bank_income.enable_multiple_items` setting in [`/link/token/create`](/docs/api/link/#linktokencreate) and obtained public tokens from the [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) endpoint, has been deprecated and replaced with a generic Multi-Item Link flow that supports all Plaid and Plaid Check products.

#### Using Bank Income with other products

Bank Income Items are fully compatible with other Plaid Item-based products, including Auth, Transactions, Balance, and Assets.

When initializing Link, if you plan to use Bank Income and Assets in the same session, it is recommended to put both `income_verification` and `assets` in the [`/link/token/create`](/docs/api/link/#linktokencreate) `products` array. If you plan to use Bank Income with Transactions, you should not put `transactions` in the `products` array, as this may increase latency. For more details, see [Choosing how to initialize Link](/docs/link/initializing-products/).

To capture an `access_token` for use with other products, call [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) after receiving the `onExit` or `onSuccess` callback from Link. This endpoint will return all `public_token`s for every Item linked to a given `user_token`. For details on the API schema, see the [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) documentation. You can then exchange these `public_tokens` for `access_tokens` using [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange).

#### Verifying Bank Income for existing Items

If your user has already connected their depository bank account to your application for a different product, you can add Bank Income to the existing Item via [update mode](/docs/link/update-mode/).

To do so, in the [`/link/token/create`](/docs/api/link/#linktokencreate) request described above, populate the `access_token` field with the access token for the Item. If the user connected their account less than two years ago, they can bypass the Link credentials pane and complete just the Income Verification step. Otherwise, they will be prompted to complete the full Plaid Bank Income Link flow.

#### Bank Income Refresh

Bank Income Refresh is available as an optional add-on to Bank Income. With Bank Income Refresh, you will be able to get updated income data for a user. Existing income sources will be updated with new transactions, and new income sources will be added if detected.

To implement Bank Income Refresh:

1. On the [webhooks page in the Dashboard](https://dashboard.plaid.com/developers/webhooks), enable Bank Income Refresh webhooks and (optionally) Bank Income Refresh Update webhooks.
2. Call [`/credit/bank_income/refresh`](/docs/api/products/income/#creditbank_incomerefresh) with the `user_token` for which you want an updated report.
3. A [`BANK_INCOME_REFRESH_COMPLETE`](/docs/api/products/income/#bank_income_refresh_complete) webhook will notify you when the process has finished. If the value of the `result` field in the webhook body is `SUCCESS`, the report was updated. If it was `FAILURE`, send the user through [update mode](/docs/link/update-mode/) and then try calling [`/credit/bank_income/refresh`](/docs/api/products/income/#creditbank_incomerefresh) again.
4. If the report was updated, when you next call [`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) or [`/credit/bank_income/pdf/get`](/docs/api/products/income/#creditbank_incomepdfget), the refreshed version of the report will be returned. To see old versions of the report, call [`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) with the `options.count` parameter set to a number greater than 1.

#### Institution coverage

To determine whether an institution supports Bank Income, you can use the Dashboard status page (look for "Bank Income") or the [`/institutions/get`](/docs/api/institutions/#institutionsget) endpoint (look for or filter by the `income_verification` product). The `income` product returned by these surfaces represents a legacy product and does not indicate coverage for Bank Income.

#### Testing Bank Income

In Sandbox, you can use the user `user_bank_income` with the password `{}`. The basic [Sandbox credentials](/docs/sandbox/test-credentials/#sandbox-simple-test-credentials) (`user_good`/`pass_good`) will not return data when used to test Bank Income.

Plaid also has additional test users for scenarios such as joint accounts, bonus income, and different levels of creditworthiness: for details, see [Credit and Income testing credentials](/docs/sandbox/test-credentials/#credit-and-income-testing-credentials).

When using special Sandbox test credentials (such as `user_bank_income` / `{}`), use the [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) flow or a non-OAuth test institution, such as First Platypus Bank (`ins_109508`). Special test credentials may be ignored when using the Sandbox Link OAuth flow.

If you’d like to test Bank Income with custom data, Plaid provides several [test JSON configuration objects](https://github.com/plaid/sandbox-custom-users/tree/main/income). To load this data into Sandbox, copy and paste the JSON into a new custom user via the [Sandbox Users pane](https://dashboard.plaid.com/developers/sandbox) in the Dashboard.

[`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) can optionally be tested in Sandbox without using Link. Call [`/user/create`](/docs/api/users/#usercreate) and pass the returned `user_token` to [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate). [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) must be called with the following request body:

/sandbox\_public/token/create request for Bank Income testing

```
{
  "client_id": "${PLAID_CLIENT_ID}",
  "secret": "${PLAID_SECRET}",
  "institution_id": "ins_20", //any valid institution id is fine
  "initial_products": ["income_verification"],
  "user_token": "user-token-goes-here", //use the user_token from the `/user/create` call made earlier
  "options": {
    "override_username": "user_bank_income", //or other test user from Credit and Income testing credentials
    "override_password": "{}",
    "income_verification": {
      "income_source_types": ["bank"],
      "bank_income": {
        "days_requested": 120 //any number of days under 730 is valid
      }
    },
  }
}
```

After calling [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate), call [`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) using the same `user_token`. The output of [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) will not be used, but calling it initializes the user token for testing.

#### Bank Income pricing

Bank Income is billed on a [one-time fee model](/docs/account/billing/#one-time-fee). Bank Income Refresh is billed on a [per-request fee model](/docs/account/billing/#per-request-flat-fee). To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

#### Next steps

If you're ready to launch to Production, see the Launch checklist.

[#### Launch Center

See next steps to launch in Production

Launch](https://dashboard.plaid.com/developers/launch-center)

#### Launch Center

See next steps to launch in Production

[Launch](https://dashboard.plaid.com/developers/launch-center)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
