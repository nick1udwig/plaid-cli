---
title: "Errors - Signal errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/signal/"
scraped_at: "2026-03-07T22:04:53+00:00"
---

# Signal errors

#### Guide to troubleshooting Signal errors

#### **ADDENDUM\_NOT\_SIGNED**

##### Signal-only actions were taken, but the client has not signed the Signal Addendum.

##### Common causes

- A customer who was enabled for Balance before October 15, 2025 and is not enabled for Signal Transaction Scores called [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport), called [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport), or called [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) with the `device` or `user` fields populated.

##### Troubleshooting steps

To use Signal Transaction Scores, which enables rules based on the `device` and `user` fields, as well as Signal Rules, contact your Account Manager and request that Signal Transaction Scores be enabled for your account, or submit a [product add request in the Dashboard](https://dashboard.plaid.com/settings/team/products) for the Signal Transaction Scores product.

If you are not ready to upgrade to Signal Transaction Scores but want to use [Signal Rules](/docs/balance/migration-guide/), contact your Account Manager to request a copy of the Signal Addendum to sign. Note that, after signing the Addendum, the `device` and `user` fields will be accepted, but will be ignored until your account is enabled for Signal Transaction Scores.

If you do not plan to use Signal Transaction Scores or Signal Rules, use [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) instead of [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate).

API error response

```
http code 403
{
 "error_type": "SIGNAL_ERROR",
 "error_code": "ADDENDUM_NOT_SIGNED",
 "error_message": "The Signal Addendum has not been signed",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **CLIENT\_TRANSACTION\_ID\_ALREADY\_IN\_USE**

##### The `client_transaction_id` value specified is not unique

##### Common causes

- When calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), the specified `client_transaction_id` has already been used.

##### Troubleshooting steps

Ensure that you use a unique `client_transaction_id` for each evaluation, even if the transaction has the same transaction details (e.g. a transaction for the same amount, by the same end user) as a previous transaction attempt.

API error response

```
http code 400
{
 "error_type": "SIGNAL_ERROR",
 "error_code": "CLIENT_TRANSACTION_ID_ALREADY_IN_USE",
 "error_message": "Client transaction id has already been used in a different transaction",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_CONFIGURATION\_STATE**

##### /signal/evaluate was called by a client whose Signal Rules configuration settings are in an invalid state

##### Common causes

- A customer who is not enabled for Signal Transaction Scores called [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) with no ruleset key specified, while opted out of the default ruleset.

##### Troubleshooting steps

Only customers who are enabled for Signal Transaction Scores are allowed to opt out of the default ruleset, so this error should never occur and indicates that your account configuration is in an invalid state. Contact your Account Manager or file a support ticket to repair your account state.

As an interim workaround, either specify a `ruleset_key` when calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), or use [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) instead.

API error response

```
http code 400
{
 "error_type": "SIGNAL_ERROR",
 "error_code": "INVALID_CONFIGURATION_STATE",
 "error_message": "Client configuration state is missing default ruleset imputation, contact your Account Manager for assistance.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NOT\_ENABLED\_FOR\_SIGNAL\_TRANSACTION\_SCORE\_RULESETS**

##### A ruleset was specified that uses Signal Transaction Scores, but the customer is only enabled for Balance-only rulesets

##### Common causes

- A customer who has previously used Signal Transaction Scores, but has since downgraded to Balance-only, attempted to use a ruleset that requires Signal Transaction Scores when calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate).

##### Troubleshooting steps

To use Signal Transaction Scores, contact your Account Manager and request that Signal Transaction Scores be enabled for your account, or submit a [product add request in the Dashboard](https://dashboard.plaid.com/settings/team/products) for the Signal Transaction Scores product.

If you do not want to use Signal Transaction Scores, convert the ruleset to Balance-only or use a different ruleset key.

API error response

```
http code 403
{
 "error_type": "SIGNAL_ERROR",
 "error_code": "NOT_ENABLED_FOR_SIGNAL_TRANSACTION_SCORE_RULESETS",
 "error_message": "Your account is not enabled for Signal Transaction Score rulesets. Use a Balance-only ruleset or contact your Account Manager to enable Signal Transaction Scores.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **RULESET\_NOT\_FOUND**

##### The specified ruleset was not found

##### Common causes

- When calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), no `ruleset_key` was specified, and the default ruleset has not been created.
- When calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), a `ruleset_key` was specified, and no matching ruleset could be found.
- When switching between Sandbox and Production, a `ruleset_key` was specified that exists only in Sandbox or only in Production.

##### Troubleshooting steps

If no `ruleset_key` was specified, create a default Ruleset via the Dashboard under [**Signal > Rules**](https://dashboard.plaid.com/signal/risk-profiles).

If a `ruleset_key` was specified, check the Dashboard under [**Signal > Rules**](https://dashboard.plaid.com/signal/risk-profiles) to make sure that there is a ruleset matching the `ruleset_key`.

- Make sure you are checking the correct environment's ruleset list (i.e. if you are calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) in Production, set the slider in the Dashboard to show Production rulesets).
- Make sure `client_id` of the Plaid Team in the Dashboard, which can be found under [**Developers > Keys**](https://dashboard.plaid.com/developers/keys), matches the `client_id` used to make the API call.

If you do not want to use a ruleset and are a Signal Transaction Scores customer, contact your Account Manager to opt out of default Signal Rulesets.

If you do not want to use a ruleset and are not a Signal Transaction Scores customer, use [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) instead of [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) to get balance data.

API error response

```
http code 400
{
 "error_type": "SIGNAL_ERROR",
 "error_code": "RULESET_NOT_FOUND",
 "error_message": "Missing 'default' ruleset. Create the ruleset here: https://dashboard.plaid.com/signal/risk-profiles",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **SIGNAL\_TRANSACTION\_NOT\_INITIATED**

##### A return was reported on a transaction that was also reported never to have happened.

##### Common causes

- A return was reported on the transaction using [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport), but [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) was already called for the same transaction with a value of `initiated: false`.
- A return was reported on the transaction using [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport), but the Signal Ruleset evaluation indicated a value of `REROUTE` for the transaction and [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) was never called with a value of `initiated: true`.

##### Troubleshooting steps

If you change your verdict on whether to process a transaction after calling [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport), or if you make a decision that differs from the Signal Ruleset evaluation (processing a transaction that evaluated to `REROUTE`, or not processing a transaction that evaluated to `ACCEPT`), correct Plaid's records by calling [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) with the most up-to-date verdict about your decision.

API error response

```
http code 400
{
 "error_type": "SIGNAL_ERROR",
 "error_code": "SIGNAL_TRANSACTION_NOT_INITIATED",
 "error_message": "Transactions previously reported as not initiated cannot be reported as returned",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
