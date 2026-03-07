---
title: "Balance - Introduction to Balance | Plaid Docs"
source_url: "https://plaid.com/docs/balance/"
scraped_at: "2026-03-07T22:04:45+00:00"
---

# Introduction to Balance

#### Retrieve real-time balance information

Get started with Balance

[API Reference](/docs/api/products/signal/)[Quickstart](/docs/quickstart/)[Demo](https://plaid.coastdemo.com/share/6786ccc5a048a5f1cf748cb5?zoom=100)

#### Overview

Balance is Plaid's product for receiving real-time balance information. This data is commonly used to tell if an account has sufficient funds before using it as a funding source for a payment transaction, helping you to avoid ACH returns. Balance is available for use exclusively in combination with other Plaid products, such as [Auth](/docs/auth/) for money movement or account funding use cases or [Transactions](/docs/transactions/) for personal finance use cases.

Prefer to learn by watching? Get an overview of how Balance works in just 3 minutes!

#### Cached and realtime balance

While you can retrieve balance data via many endpoints, including by calling the free-to-use [`/accounts/get`](/docs/api/accounts/#accountsget) endpoint, this data is cached, making it unsuitable for situations where you need real-time balance information, except immediately after the Item has been added.

An Item with Transactions may update cached balances once a day or more, while an Item with only Auth may update cached balances every 30 days or even less frequently.

For retrieving real-time balance for payments use cases, only Balance should be used.

Because Balance always performs a real-time data extraction, latency is higher than other Plaid endpoints (p50 ~3s, p95 ~11s). To check for ACH return risk in critical user-present flows that require low latency, consider [Signal Transaction Scores](/docs/signal/). For more details, see [Reducing latency and getting better insights with Signal Transaction Scores](/docs/balance/#reducing-latency-and-getting-better-insights-with-signal-transaction-scores).

#### Balance integration options

There are two options you can use for retrieving balance data: using Signal Rules (with [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate)) or using [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget). While they are priced the same, return the same account data, and have the same latency, each method works differently and is designed for different use cases.

##### Balance using /signal/evaluate

The [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) endpoint is the recommended way to retrieve balance data whenever you are checking balance to evaluate the insufficient funds risk of a proposed ACH transaction. [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) checks balance and returns current and available balance, as well as a recommended action to take (i.e., accept the transaction, review it further, or request a different payment method) based on a ruleset evaluation using [Signal Rules](/docs/signal/signal-rules/). With Signal Rules, you can easily customize your business logic around these recommended actions in a Dashboard-based UI, without making any code changes.

![](/assets/img/docs/balance-rule-editor.png)

Using Balance with Signal Rules in the Dashboard

[`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) is available for all new Balance customers who received Production access after October 15, 2025. Existing Balance customers will be enabled for [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) in the near future. If you are an existing Balance customer and would like immediate access to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), contact your Account Manager; for more details, see the [migration guide](/docs/balance/migration-guide/).

##### Balance using /accounts/balance/get

[`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) retrieves balance data, including current and available balance, but does not run Signal Rules. This endpoint is recommended for all use cases where you are *not* checking balance to evaluate a proposed ACH transaction: for example, if you need real-time balance data for a Personal Financial Management (PFM) use case, or for treasury management, or for any non-US bank account. [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) is also the legacy endpoint for checking balance that many customers use in existing integrations.

#### Integration overview using /signal/evaluate

1. [Create a Balance-only Rule](/docs/signal/signal-rules/) in the [Signal Rules Dashboard](https://dashboard.plaid.com/signal/risk-profiles).
2. Call [`/link/token/create`](/docs/api/link/#linktokencreate). Along with any other parameters you specify, make sure to include the following:

   - Include `signal` and `auth` in the `products` array, along with any other Plaid product(s) you plan to use. Do *not* include `balance` in the `products` array.
3. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see the [Link documentation](/docs/link/).

   - If you will be using the linked Item as a payment source or destination, it is recommended to [configure your Link customization](/docs/link/customization/) to use the setting [Account Select: Enabled for one account](/docs/link/customization/#account-select). Otherwise, you will need to build a UI within your app for the user to indicate which account they want to use for the payment, if their bank login is associated with more than one bank account.
4. Once the end user has completed the Link flow, exchange the `public_token` for an `access_token` by calling [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange).
5. Call [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) and [determine the next steps based on results](/docs/signal/signal-rules/#using-signal-ruleset-results).
6. [Report ACH returns and decisions](/docs/signal/reporting-returns/) to Plaid.
7. (Optional) After launch, periodically [review and tune your Signal Rules](/docs/signal/tuning-rules/) using the Dashboard.

#### Integration overview using /accounts/balance/get

1. Call [`/link/token/create`](/docs/api/link/#linktokencreate) with the product(s) you plan to use. Do *not* include `balance` in the `products` array.
2. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see the [Link documentation](/docs/link/).
3. Once the end user has completed the Link flow, exchange the `public_token` for an `access_token` by calling [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange).
4. Call [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget), passing in the `access_token` returned in step 3.

#### Reducing latency and getting better insights with Signal Transaction Scores

For use cases that involve evaluating a proposed ACH transaction to reduce return risk, Balance can be used alongside [Signal Transaction Scores](/docs/signal/). Signal Transaction Scores has lower latency (p95 <2s versus Balance's 11s), making it ideal for user-present transactions in critical flows such as onboarding or purchasing. With Signal Transaction Scores, you can access 80+ criteria for rule generation and ML-powered recommendations and insights to help you optimize and fine-tune your transaction approval logic.

Signal Transaction Scores and Balance are both part of Plaid Signal, Plaid's solution for ACH risk management. You can purchase and use either Balance or Signal Transaction Scores by itself, or combine them for a more comprehensive ACH risk management approach.

Once you have integrated Balance using [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), Signal Transaction Scores requires no additional engineering work to integrate. Just use the Signal Rule editor in the Dashboard to upgrade your existing Balance-only rule into a Signal Transaction Score-powered rule.

To learn more, or to request access to Signal Transaction Scores, contact [Sales](https://plaid.com/contact/) or your account manager.

#### Migration guide

If you are an existing customer and would like to migrate from [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) to take advantage of Signal Rules, see the [Balance migration guide](/docs/balance/migration-guide/).

#### Current and available balance

Banks represent balance with two separate values, current balance and available balance. For assessing ACH risk, available balance is often more important than current balance, as it represents predicted balance net of any pending transactions, while current balance does not take pending transactions into account.

For credit accounts, available balance indicates the amount that can be spent without hitting the credit limit (net of any pending transactions), while in investment accounts, it indicates the total available cash.

In some cases, a financial institution may not provide available balance information. If necessary, you can often calculate available balance by starting with the current balance, then using the [Transactions](/docs/transactions/) product to detect any pending transactions and adjusting the balance accordingly.

##### Typical balance fill rates by field

| Field | Typical fill rate |
| --- | --- |
| Current balance | 99% |
| Available balance | 91% |

#### Balance pricing

Balance is billed on a [flat fee model](/docs/account/billing/#per-request-flat-fee), meaning you will be billed once for each successful call to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) or [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget).

[`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) (when used with a Balance-only ruleset) and [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) are billed at the same rate.

When using [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), you will be billed if the API call is successful, even if balance extraction fails.

To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

#### Next steps

To get started building with Balance, see [Add Balance to your App](/docs/balance/add-to-app/).

If you're ready to launch to Production, see the Launch Center.

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
