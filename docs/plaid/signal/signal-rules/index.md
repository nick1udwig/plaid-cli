---
title: "Signal - Signal Rules | Plaid Docs"
source_url: "https://plaid.com/docs/signal/signal-rules/"
scraped_at: "2026-03-07T22:05:19+00:00"
---

# Signal Rules

#### Learn about configuring Signal Rules

#### Using Signal Rules

With Signal Rules, you can configure your decision logic directly within the Signal Dashboard using templates from Plaid, then use Dashboard tools to tailor the rules to your own traffic.

Signal Rules are available to users of both Signal Transaction Scores and Balance products.

Rules can be based on scores (higher scores indicate higher risk), attributes of the transaction, or a combination.

You can set up the following actions for your business to take as a result of a rule:

- `ACCEPT` and process the transaction
- `REVIEW` the transaction further before processing
- `REROUTE` the customer to another payment method, as the transaction is too risky

You can also create more complex action flows through the use of [custom action keys](/docs/signal/signal-rules/#using-a-custom-action-key-for-advanced-actions).

You can define a single ruleset for all your transactions, or you can build use-case specific rulesets, such as one for an initial deposit during customer onboarding and another for a returning customer depositing funds.

#### Creating a ruleset

Before using Signal Rules, you must create a Ruleset in the Dashboard, under [**Signal > Rules**](https://dashboard.plaid.com/signal/risk-profiles). You can create a ruleset from scratch, or Plaid can suggest an initial ruleset template relevant for your use case by utilizing context collected during your onboarding (such as transaction size).

Rules are evaluated sequentially, from top to bottom, typically ordered by the severity of the rule violation. If data that the rule depends on is not available, the rule will be skipped.

The `result` of the first rule that matches will be applied to the transaction and included in the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) response. You will be required to have a fallback rule that is applied if no other rules match.

![Example Plaid Signal Ruleset](/assets/img/docs/signal/ruleset-example.png)

Signal Ruleset example configuration

#### Signal Transaction Score versus Balance-only rulesets

When creating a ruleset, you will be given the choice between creating a **Signal Transaction Score** or **Balance-only** ruleset.

![](/assets/img/docs/signal/create-new-ruleset.png)

Balance customers who are not enabled for Signal Transaction Scores can create only **Balance-only** rulesets.

[`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) calls that use a **Signal Transaction Score ruleset**:

- Return the full `core_attributes` array with all 80+ core attributes, all of which can be used in the rule configuration
- Have ultra-low latency (p95 <2s)
- May use cached data
- Will not perform a real-time balance extraction
- Will be billed as Signal Transaction Score calls
- Are available only to Signal Transaction Scores customers

[`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) calls that use a **Balance-only ruleset**:

- Return only two `core_attributes`: `available_balance` and `current_balance`, which can be used in the rule configuration
- Have higher latency (p95 ~11s)
- Always perform a real-time fresh balance extraction
- Will be billed as Balance calls
- Are available to all customers on the Plaid Signal platform, including Balance-only customers

#### Using Signal Ruleset results

To evaluate a transaction's risk, call [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) and include the required request properties in addition to the ruleset's `ruleset_key`, which can be found in the Dashboard, in smaller type under the name of the ruleset.

The more properties you provide when calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), such as the `user` or `device` fields, the more accurate the results will be. Providing the `client_user_id` is also recommended to aid in your debugging process later.

In the response, use the `ruleset.result` property to determine the next steps.

Most integrations use the `result` to decide between approving a payment for processing (`result=APPROVE`), or rerouting the consumer to another payment method (`result=REROUTE`). If that's the case with your integration, ensure that your ruleset emits only these two result types.

If your integration is taking more complex actions before deciding whether to process the transaction (for example, sending certain transactions to a manual review queue), use the `REVIEW` result type. Remember to [report the final decision](/docs/signal/reporting-returns/#reporting-decisions) for these reviewed transactions later.

Signal Transaction Score results can never be used as an input to determine a consumer's eligibility to access your services. Signal Transaction Scores evaluates the risk of a transaction from a specific bank account and can be used to determine if you should require the user to use a different payment method, according to your risk tolerance.

#### Opting out of the default ruleset

If you do not specify a `ruleset_key`, your `default` ruleset will be used. If you do not have a `default` ruleset and the `ruleset_key` is not specified, [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will return an error unless you have opted out of the default ruleset. If you omit the `ruleset_key` and have opted out of the default ruleset, [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will return `core_attributes` but will not return a `ruleset` or a `ruleset.result`.

All customers who were enabled for Signal Transaction Scores in Production before October 15, 2025, are automatically opted out of the default ruleset, in order to maintain parity with earlier behavior (before ruleset functionality was introduced).

Signal Transaction Scores customers who use [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) exclusively to incorporate raw insights from the `core_attributes` into their own risk models, and do not want to use Signal rulesets at all, can choose to opt out of the default ruleset by contacting their Account Manager. For more details on using Signal Transaction Scores in this mode, see [Accessing scores and attributes directly via API](/docs/signal/signal-rules/#accessing-scores-and-attributes-directly-via-api).

Opting out of the default ruleset only impacts behavior when no ruleset is specified in the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) call; if you explicitly specify a ruleset when calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), the ruleset will be used.

Balance-only customers cannot opt out of the default ruleset. If you wish to check real-time balance without a ruleset, use [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) instead of [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate).

To re-enable using default rulesets, contact your Account Manager.

##### For existing customers prior to May 2025: Migrating away from the outcome field

Customers who used Signal Rules before May 2025 used the `ruleset.outcome` field. This field has been replaced by [`ruleset.result`](/docs/api/products/signal/#signal-evaluate-response-ruleset-result). Update your integration to use the new values from the `ruleset.result` field:

| Outcome (old) | Result (new) |
| --- | --- |
| `block` | `REROUTE` |
| `accept` | `ACCEPT` |
| `review` | `REVIEW` |

For most integrations, that's all you need to change. Once you've updated your integration, inform your account manager.

If you have an advanced use case where you take different actions for transactions with the same `result`, you can use the `custom_action_key`. For example, to apply a variable hold time, make sure your Signal rules emit an `result=ACCEPT` value for all transactions you intend to process. Then, use the `custom_action_key` to define a `3-day-hold` for some rules and `5-day-hold` for others. For more details, see [Using a custom action key](/docs/signal/signal-rules/#using-a-custom-action-key-for-advanced-actions).

#### Using a custom action key for advanced actions

Some integrations may need more options for handling a transaction than the three choices of `ACCEPT`, `REVIEW`, or `REROUTE`. Common use cases include:

- Applying a variable hold time to transactions that were accepted and processed, depending on the level of risk.
- Applying labeling to transactions to denote a cohort for data analysis purposes.
- Running a Signal Transaction Score ruleset check on a proposed transaction, and falling back to a higher-latency, Balance-only ruleset check if the transaction is higher risk.

To support more complex actions, check the "Configure advanced options" box on the ruleset editing screen and specify a `custom_action_key`.

You can then create multiple rules with the same result but different custom actions. For example:

- If medium risk, `ACCEPT` with `custom_action_key=5-day-hold`
- If low risk, `ACCEPT` with `custom_action_key=3-day-hold`

When calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), use the [`custom_action_key`](/docs/api/products/signal/#signal-evaluate-response-ruleset-triggered-rule-details-custom-action-key) in the response to determine the action to take.

#### Data availability limitations

The amount of consumer-permissioned data available can vary. Plaid attempts to prefetch this data at a regular cadence and will also attempt to fetch fresh data during the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) call if needed. Most Signal Transaction Score evaluations will return information about the Item's use in the network, the latest balance, historical balance information, or account activity. The full set of data is usually available to Plaid within 15 seconds of Plaid Link being completed. However, partial data may be returned if [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) is called before this process completes, the Item has been inactive for a long time, or the financial institution's connectivity to Plaid is degraded or down. Plaid may pause background refreshes for Items that have been inactive for an extended period. If this occurs, the next [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) call may take longer as Plaid attempts to fetch fresh data before returning a response.

When used with Items that do not have a connection to a financial institution, (i.e., Items created verified using [alternative methods](/docs/auth/coverage/) such as micro-deposits or Database Auth), Plaid can generate Signal Transaction Scores for approximately 30% of these Items, although Items created in this way will have no data in the `core_attributes` response property. If there is insufficient data, the `scores` property will be `null`, and any rules using `score` will be skipped (potentially triggering your fallback rule). You will still be billed for the evaluation (unless you have no Signal Ruleset enabled).

#### Testing Signal Rules in Sandbox

In the Sandbox environment, scores and attributes returned by [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will be randomly generated. You can influence the score returned by [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) in Sandbox by providing particular values in the `amount` request field. Use the values from the following table to control the returned score:

| Amount | Score |
| --- | --- |
| 3.53 | 10 |
| 12.17 | 60 |
| 27.53 | 90 |

You can use these hardcoded responses and create simple rules in Sandbox (such as `REROUTE if bank_score == 90`) to test various result types.

#### Debugging transactions

All evaluations can be viewed in realtime on the [transactions viewer](https://dashboard.plaid.com/signal/transactions-viewer) Dashboard page. Look up your transaction using its `client_transaction_id`. You can see its scores, the ruleset's `result`, and the data used to build the score.

#### Debugging rules

You can view all transactions that matched a specific rule by clicking `View Performance` in the `...` icon on the rule.

![Navigation to the rule performance page](/assets/img/docs/signal/view-rule-performance-nav.png)

Navigation to the rule performance page

![Example performance page for a single Signal rule](/assets/img/docs/signal/one-rule-performance.png)

Performance page of a single rule

#### Tuning your rules

After your integration is live, Plaid will report metrics such as approval rate, and ACH return rate, in the Signal Dashboard. You will use these Dashboard tools to [tune your rule logic](/docs/signal/tuning-rules/), which is a critical step in a successful launch.

You must integrate the [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport) endpoint to use the Plaid Dashboard performance data features. If you do not report returns, performance data will not be available.

#### Shadow tests and proof of concept testing

If you're using Signal as part of a shadow test or proof of concept, you should still follow all the of the above integration guidance as though you were launching to production (including use of Signal Rules) - you just won't use the results to make the real-world payment decision. Log any relevant rule results, scores, or attributes as necessary.

It is particularly critical to [report decisions and returns](/docs/signal/reporting-returns/) in shadows and proof of concepts, in order for Plaid to measure performance and provide optimized rule suggestions tailored to your traffic.

#### Accessing scores and attributes directly via API

Plaid strongly recommends using Signal Rules. However, if you have your own risk management system in which you are using Signal Transaction Scores as one of multiple data sources, you might choose to ingest this data directly via the API instead. You can also ingest this data directly via the API while using Signal Rules as well. If you don't use Signal Rules, Plaid may be limited in the support it can offer you in tuning external logic.

To directly access scores and data attributes, call [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate). The `scores` and `core_attributes` properties in the response will contain the raw data you can provide to your model. For a full listing of `core_attributes`, contact your Account Manager. If you are not using Signal Rules, you should also contact your Account Manager to [opt out of the default ruleset](/docs/signal/signal-rules/#opting-out-of-the-default-ruleset).

Even if you aren't using Signal Rules, it's still recommended to [report decisions and returns](/docs/signal/reporting-returns/) in order to use the [Signal Performance pane](/docs/signal/signal-rules/#tuning-your-rules), which helps you understand performance and find the right score threshold.

#### For Reseller Partners

A reseller partner will have a global ruleset that all their end customers inherit. If the reseller makes a change to that global default ruleset, all the end-customers also automatically adopt that change as well. End customers can also create their own rulesets, which will not propagate back to the reseller partner.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
