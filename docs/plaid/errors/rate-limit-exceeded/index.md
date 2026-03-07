---
title: "Errors - Rate Limit Exceeded errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/rate-limit-exceeded/"
scraped_at: "2026-03-07T22:04:52+00:00"
---

# Rate Limit Exceeded Errors

#### Guide to troubleshooting rate limit exceeded errors

#### Rate limit table

Errors of type `RATE_LIMIT_EXCEEDED` will occur when the rate limit for a
particular endpoint has been exceeded. Default rate limit thresholds for some of
the most commonly rate-limited endpoints are shown below. Note that these tables
are not an exhaustive listing of all Plaid rate limits or rate-limited
endpoints, that some customers may experience different rate limit thresholds from
those shown, and that rate limits are subject to change at any time.

In general, Plaid default rate limits are set such that using the API as
designed should typically not cause a rate limit to be encountered. If your use
case requires a higher rate limit, contact your Account Manager or file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

##### Production rate limits

| Endpoint | Max requests per Item | Max requests per client |
| --- | --- | --- |
| [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) | 5 per minute | 1,200 per minute |
| [`/accounts/get`](/docs/api/accounts/#accountsget) | 15 per minute | 15,000 per minute |
| [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) | 5 per minute | 50 per minute |
| [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) | 15 per minute (per asset report) | 1,000 per minute |
| [`/asset_report/pdf/get`](/docs/api/products/assets/#asset_reportpdfget) | 15 per minute (per asset report) | 50 per minute |
| [`/asset_report/refresh`](/docs/api/products/assets/#asset_reportrefresh) | 5 per minute (per asset report) | 50 per minute |
| [`/auth/get`](/docs/api/products/auth/#authget) | 15 per minute | 12,000 per minute |
| `/cra/check_report/*` | N/A | 100 per minute |
| `/cra/monitoring_insights/*` | N/A | 100 per minute |
| [`/identity/get`](/docs/api/products/identity/#identityget) | 15 per minute | 2,000 per minute |
| [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) | N/A | 120 per minute |
| [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget) | N/A | 420 per minute |
| [`/identity_verification/list`](/docs/api/products/identity-verification/#identity_verificationlist) | N/A | 300 per minute |
| [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry) | N/A | 120 per minute |
| [`/institutions/get`](/docs/api/institutions/#institutionsget) | N/A | 50 per minute |
| [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) | N/A | 400 per minute |
| [`/investments/holdings/get`](/docs/api/products/investments/#investmentsholdingsget) | 15 per minute | 2,000 per minute |
| [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) | 30 per minute | 20,000 per minute |
| [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) | 15 per minute | 500 per minute |
| [`/investments/refresh`](/docs/api/products/investments/#investmentsrefresh) | 1 per minute | 100 per minute |
| [`/item/get`](/docs/api/items/#itemget) | 15 per minute | 5,000 per minute |
| [`/item/remove`](/docs/api/items/#itemremove) | 30 per minute | 2,000 per minute |
| [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget) | 15 per minute | 1,000 per minute |
| [`/network/status/get`](/docs/api/network/#networkstatusget) | N/A | 1,000 per minute |
| [`/processor/token/create`](/docs/api/processors/#processortokencreate) | N/A | 500 per minute |
| [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) | 70 per hour, 10 per transaction per hour | 20 per second |
| [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) | N/A | 4,000 per minute |
| [`/signal/prepare`](/docs/api/products/signal/#signalprepare) | N/A | 2,000 per minute |
| [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport) | N/A | 4,000 per minute |
| [`/transactions/enrich`](/docs/api/products/enrich/#transactionsenrich) | N/A | 100 per minute |
| [`/transactions/get`](/docs/api/products/transactions/#transactionsget) | 30 per minute | 20,000 per minute |
| [`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget) | 20 per minute | 1,000 per minute |
| [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) | 2 per minute | 100 per minute |
| [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) | 50 per minute | 2,500 per minute |
| [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) | 100 per hour | 2,500 per minute |
| [`/transfer/cancel`](/docs/api/products/transfer/initiating-transfers/#transfercancel) | N/A | 250 per minute |
| [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) | N/A | 2,500 per minute |
| [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) | N/A | 5,000 per minute |
| [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) | N/A | 5,000 per minute |
| [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) | N/A | 250 per minute |
| [`/transfer/recurring/cancel`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcancel) | N/A | 250 per minute |
| [`/transfer/recurring/create`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcreate) | N/A | 5,000 per minute |
| [`/transfer/refund/cancel`](/docs/api/products/transfer/refunds/#transferrefundcancel) | N/A | 250 per minute |
| [`/transfer/refund/create`](/docs/api/products/transfer/refunds/#transferrefundcreate) | N/A | 5,000 per minute |
| `/transfer/*/get` | N/A | 5,000 per minute |
| `/transfer/*/list` | N/A | 100 per minute |
| [`/watchlist_screening/individual/get`](/docs/api/products/monitor/#watchlist_screeningindividualget) | N/A | 300 per minute |
| [`/payment_initiation/recipient/create`](/docs/api/products/payment-initiation/#payment_initiationrecipientcreate) | N/A | 240 per minute |
| [`/payment_initiation/recipient/get`](/docs/api/products/payment-initiation/#payment_initiationrecipientget) | N/A | 240 per minute |
| [`/payment_initiation/recipient/list`](/docs/api/products/payment-initiation/#payment_initiationrecipientlist) | N/A | 240 per minute |
| [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) | N/A | 240 per minute |
| [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) | N/A | 240 per minute |
| [`/payment_initiation/payment/list`](/docs/api/products/payment-initiation/#payment_initiationpaymentlist) | N/A | 240 per minute |
| [`/payment_initiation/payment/reverse`](/docs/api/products/payment-initiation/#payment_initiationpaymentreverse) | N/A | 240 per minute |
| [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate) | N/A | 100 per minute |
| [`/payment_initiation/consent/get`](/docs/api/products/payment-initiation/#payment_initiationconsentget) | N/A | 240 per minute |
| [`/payment_initiation/consent/revoke`](/docs/api/products/payment-initiation/#payment_initiationconsentrevoke) | N/A | 100 per minute |
| [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute) | N/A | 100 per minute (5 per consent) |
| [`/statements/download`](/docs/api/products/statements/#statementsdownload) | N/A | 50 per minute |
| [`/statements/list`](/docs/api/products/statements/#statementslist) | N/A | 100 per minute |
| [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh) | N/A | 50 per minute |

##### Sandbox rate limits

| Endpoint | Max requests per Item | Max requests per client |
| --- | --- | --- |
| [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) | 25 per minute | 100 per minute |
| [`/accounts/get`](/docs/api/accounts/#accountsget) | 100 per minute | 5,000 per minute |
| [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) | N/A | 100 per minute |
| [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) | 1,000 per minute | 1,000 per minute |
| [`/asset_report/refresh`](/docs/api/products/assets/#asset_reportrefresh) | N/A | 100 per minute |
| [`/auth/get`](/docs/api/products/auth/#authget) | 100 per minute | 500 per minute |
| `/cra/check_report/*` | N/A | 100 per minute |
| `/cra/monitoring_insights/*` | N/A | 100 per minute |
| [`/identity/get`](/docs/api/products/identity/#identityget) | 100 per minute | 1,000 per minute |
| `/identity_verification/*` | N/A | 60 per minute |
| [`/institutions/get`](/docs/api/institutions/#institutionsget) | N/A | 10 per minute |
| [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) | N/A | 400 per minute |
| [`/investments/holdings/get`](/docs/api/products/investments/#investmentsholdingsget) | 100 per minute | 1,000 per minute |
| [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) | 100 per minute | 1,000 per minute |
| [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) | 15 per minute | 500 per minute |
| [`/investments/refresh`](/docs/api/products/investments/#investmentsrefresh) | 1 per minute | 100 per minute |
| [`/item/get`](/docs/api/items/#itemget) | 40 per minute | 5,000 per minute |
| [`/item/remove`](/docs/api/items/#itemremove) | 100 per minute | 500 per minute |
| [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget) | 10 per minute | 1,000 per minute |
| [`/network/status/get`](/docs/api/network/#networkstatusget) | N/A | 1,000 per minute |
| [`/processor/token/create`](/docs/api/processors/#processortokencreate) | N/A | 500 per minute |
| [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) | 70 per hour, 10 per transaction per hour | 30 per second |
| [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) | N/A | 4,000 per minute |
| [`/signal/prepare`](/docs/api/products/signal/#signalprepare) | N/A | 2,000 per minute |
| [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport) | N/A | 4,000 per minute |
| [`/transactions/enrich`](/docs/api/products/enrich/#transactionsenrich) | N/A | 100 per minute |
| [`/transactions/get`](/docs/api/products/transactions/#transactionsget) | 80 per minute | 1,000 per minute |
| [`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget) | 20 per minute | 1,000 per minute |
| [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) | 2 per minute | 100 per minute |
| [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) | 50 per minute | 1,000 per minute |
| [`/sandbox/transactions/create`](/docs/api/sandbox/#sandboxtransactionscreate) | 2 per minute | 50 per minute |
| [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) | 100 per hour | 100 per minute |
| [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) | N/A | 100 per minute |
| [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) | N/A | 100 per minute |
| [`/transfer/recurring/create`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcreate) | N/A | 100 per minute |
| [`/transfer/refund/create`](/docs/api/products/transfer/refunds/#transferrefundcreate) | N/A | 100 per minute |
| [`/watchlist_screening/individual/get`](/docs/api/products/monitor/#watchlist_screeningindividualget) | N/A | 300 per minute |
| [`/payment_initiation/recipient/create`](/docs/api/products/payment-initiation/#payment_initiationrecipientcreate) | N/A | 240 per minute |
| [`/payment_initiation/recipient/get`](/docs/api/products/payment-initiation/#payment_initiationrecipientget) | N/A | 240 per minute |
| [`/payment_initiation/recipient/list`](/docs/api/products/payment-initiation/#payment_initiationrecipientlist) | N/A | 240 per minute |
| [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) | N/A | 100 per minute |
| [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) | N/A | 240 per minute |
| [`/payment_initiation/payment/list`](/docs/api/products/payment-initiation/#payment_initiationpaymentlist) | N/A | 240 per minute |
| [`/payment_initiation/payment/reverse`](/docs/api/products/payment-initiation/#payment_initiationpaymentreverse) | N/A | 240 per minute |
| [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate) | N/A | 100 per minute |
| [`/payment_initiation/consent/get`](/docs/api/products/payment-initiation/#payment_initiationconsentget) | N/A | 240 per minute |
| [`/payment_initiation/consent/revoke`](/docs/api/products/payment-initiation/#payment_initiationconsentrevoke) | N/A | 100 per minute |
| [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute) | N/A | 100 per minute (10 per consent) |
| [`/statements/download`](/docs/api/products/statements/#statementsdownload) | N/A | 50 per minute |
| [`/statements/list`](/docs/api/products/statements/#statementslist) | N/A | 100 per minute |
| [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh) | N/A | 50 per minute |

#### **ACCOUNTS\_LIMIT**

##### Too many requests were made to the `/accounts/get` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/accounts/get`](/docs/api/accounts/#accountsget) in Production are rate limited at
  a maximum of 15 requests per Item per minute and 15,000 per client per minute.
  In the Sandbox, they are limited at a maximum of 100 per Item per minute and
  5,000 per client per minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/accounts/get`](/docs/api/accounts/#accountsget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "ACCOUNTS_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ACCOUNTS\_BALANCE\_GET\_LIMIT**

##### Too many requests were made to the `/accounts/balance/get` endpoint.

##### Common causes

- Too many requests were made in a short period of time by a single client.
  Requests to `/account/balance/get` in Production are client rate limited to 1,200 requests per client per minute. In the
  Sandbox environment, they are client rate limited at a maximum of 100 requests
  per client per minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "ACCOUNTS_BALANCE_GET_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **AUTH\_LIMIT**

##### Too many requests were made to the `/auth/get` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to [`/auth/get`](/docs/api/products/auth/#authget)
  in Production are rate limited at a maximum of 15
  requests per Item per minute and 12,000 per client per minute. In the Sandbox,
  they are rate limited at a maximum of 100 requests per Item per minute and 500
  requests per client per minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/auth/get`](/docs/api/products/auth/#authget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "AUTH_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **BALANCE\_LIMIT**

##### Too many requests were made to the `/accounts/balance/get` endpoint.

##### Common causes

- Too many requests were made for a single Item in a short period of time.
  Requests to `/account/balance/get` in Production
  are Item rate limited at a maximum of 5 requests per Item per minute. In the
  Sandbox environment, they are Item rate limited at a maximum of 25 requests
  per Item per minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "BALANCE_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **CREDITS\_EXHAUSTED**

##### You have used up your free API usage allocation in Limited Production

##### Common causes

- You ran out of free API calls for a given product in Limited Production.
- You do not yet have Production access and hit the Item creation cap in Limited Production.

##### Troubleshooting steps

Request full Production access from the
[Dashboard](https://dashboard.plaid.com/onboarding/).

If you need more Limited Production API calls in order to test your use case, contact your Account Manager or file a ticket with
[Plaid Support](https://dashboard.plaid.com/support/new/product-and-development/account-administration/addition-limit-exceeded)
to request additional usage.

Test in a different environment, such as Sandbox.

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "CREDITS_EXHAUSTED",
 "error_message": "Free usage exhausted, please request full Production access to continue using this product",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **IDENTITY\_LIMIT**

##### Too many requests were made to the `/identity/get` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/identity/get`](/docs/api/products/identity/#identityget) in Production are rate limited at
  a maximum of 15 requests per Item per minute and 2,000 per client per minute.
  In the Sandbox environment, they are rate limited at a maximum of 100 requests
  per Item per minute and 1,000 requests per client per minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/identity/get`](/docs/api/products/identity/#identityget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "IDENTITY_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTIONS\_GET\_LIMIT**

##### Too many requests were made to the `/institutions/get` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/institutions/get`](/docs/api/institutions/#institutionsget) in Production are rate
  limited at a maximum of 25 per client per minute. In the Sandbox environment,
  they are rate limited at a maximum of 10 requests per client per minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/institutions/get`](/docs/api/institutions/#institutionsget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "INSTITUTIONS_GET_LIMIT",
 "error_message": "rate limit exceeded for attempts to access \"institutions get by id\". please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTIONS\_GET\_BY\_ID\_LIMIT**

##### Too many requests were made to the `/institutions/get_by_id` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) are rate limited at a maximum of 400 requests per
  client per minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "INSTITUTIONS_GET_BY_ID_LIMIT",
 "error_message": "rate limit exceeded for attempts to access \"institutions get by id\". please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTION\_RATE\_LIMIT**

##### Too many requests were made to a given institution.

##### Common causes

- Too many requests were made by Plaid in a short period of time to a given institution. Because each institution has unique rate limiting behavior, Plaid cannot provide exact details of how many requests are necessary to trigger this behavior.
- This error will only trigger when calling API endpoints that request realtime data from the institution, such as [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) or [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh).
- This institution-level limit is distinct from per-Item or per-client limits and is not applied to user-present traffic (e.g., within Link).
- In some cases, the cause of the rate limit may be another client and may not be your client application's behavior.

##### Troubleshooting steps

If your client made a very large number of requests in a short time to a realtime endpoint such as [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) at a single institution, consider adding logic to spread these requests over a longer time window.

Use an exponential backoff retry algorithm to try your request again later.

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "INSTITUTION_RATE_LIMIT",
 "error_message": "The institution is currently receiving too many requests. Please try again later.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVESTMENT\_HOLDINGS\_GET\_LIMIT**

##### Too many requests were made to the `/investments/holdings/get` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/investments/holdings/get`](/docs/api/products/investments/#investmentsholdingsget) in Production are
  rate limited at a maximum of 15 requests per Item per minute and 2,000 per
  client per minute. In the Sandbox environment, they are rate limited at a
  maximum of 100 requests per Item per minute and 1,000 requests per client per
  minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/investments/holdings/get`](/docs/api/products/investments/#investmentsholdingsget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "INVESTMENT_HOLDINGS_GET_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVESTMENT\_TRANSACTIONS\_LIMIT**

##### Too many requests were made to the `/investments/transactions/get` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) in Production are
  rate limited at a maximum of 30 requests per Item per minute and 20,000 per
  client per minute. In the Sandbox environment, they are rate limited at a
  maximum of 100 requests per Item per minute and 1,000 requests per client per
  minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "INVESTMENT_TRANSACTIONS_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ITEM\_GET\_LIMIT**

##### Too many requests were made to the `/item/get` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to [`/item/get`](/docs/api/items/#itemget)
  in Production are rate limited at a maximum of 15
  requests per Item per minute and 5,000 per client per minute. In the Sandbox
  environment, they are rate limited at a maximum of 40 requests per Item per
  minute and 5,000 requests per client per minute.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/item/get`](/docs/api/items/#itemget).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "ITEM_GET_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **RATE\_LIMIT**

##### Too many requests were made.

##### Common causes

- Too many requests were made in a short period of time.
- Sandbox credentials (the username `user_good` or `user_custom`) were used to
  attempt to log in to Production. Because using these
  credentials in a live environment is a common misconfiguration, they are
  frequently subject to rate limiting.
- A Link attempt was detected as potential attempted abuse. For example, this
  may occur if a user enters their credentials incorrectly too many times in a
  row.

##### Troubleshooting steps

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to a Plaid
endpoint.

Verify that you are not using Sandbox credentials in Production.

Have your user wait 1-2 days and attempt to log in again later.

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "RATE_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTIONS\_LIMIT**

##### Too many requests were made to a transactions endpoint such as`/transactions/get` or `/transactions/refresh`.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/transactions/get`](/docs/api/products/transactions/#transactionsget) in Production are rate
  limited at a maximum of 30 requests per Item per minute and 20,000 per client
  per minute. In the Sandbox environment, they are rate limited at a maximum of
  80 requests per Item per minute and 1,000 requests per client per minute.

##### Troubleshooting steps

Use the `count` request parameter in the [`/transactions/get`](/docs/api/products/transactions/#transactionsget) call to increase the number of transactions retrieved per request. By default, transactions endpoints will retrieve 100 transactions at a time, but you can retrieve up to 500 at a time by setting `count` to the desired number. The more transactions you retrieve per API call, the fewer total API calls you will need to make.

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to a transactions endpoint.

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "TRANSACTIONS_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTIONS\_REFRESH\_LIMIT**

##### Too many requests were made to the `/transactions/refresh` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) are rate
  limited at a maximum of 2 requests per Item per minute and 100 per client
  per minute.

##### Troubleshooting steps

If you are sending [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) requests based on end user behavior (such as a "refresh" button in your UI), limit the frequency at which you will send these requests or the end user can initiate them. Since [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) is billed on a per API call basis, unneeded API calls to this endpoint can lead to excess billing.

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh).

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "TRANSACTIONS_REFRESH_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTIONS\_SYNC\_LIMIT**

##### Too many requests were made to the `/transactions/sync` endpoint.

##### Common causes

- Too many requests were made in a short period of time. Requests to
  [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) in Production are rate
  limited at a maximum of 50 requests per Item per minute and 2,500 per client
  per minute. In the Sandbox environment, they are rate limited at a maximum of
  50 requests per Item per minute and 1,000 requests per client per minute.

##### Troubleshooting steps

Use the `count` request parameter in the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) call to increase the number of transactions retrieved per request. By default, transactions endpoints will retrieve 100 transactions at a time, but you can retrieve up to 500 at a time by setting `count` to the desired number. The more transactions you retrieve per API call, the fewer total API calls you will need to make.

Check the [activity log](https://dashboard.plaid.com/activity/logs) in the
Dashboard to view a history of all requests made with your API keys and verify
that you are not accidentally sending an excessive number of requests to
[`/transactions/sync`](/docs/api/products/transactions/#transactionssync). If you are not already, you may be able to reduce your
number of requests to [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) by specifying 500 for the `count`
request field.

If your use case requires a higher rate limit, contact your Account Manager or
file a
[Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 429
{
 "error_type": "RATE_LIMIT_EXCEEDED",
 "error_code": "TRANSACTIONS_SYNC_LIMIT",
 "error_message": "rate limit exceeded for attempts to access this item. please try again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
