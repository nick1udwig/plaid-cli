---
title: "Balance - Migrate to Signal Rules | Plaid Docs"
source_url: "https://plaid.com/docs/balance/migration-guide/"
scraped_at: "2026-03-07T22:04:46+00:00"
---

# Balance migration guide

#### Why migrate?

If you use Balance to reduce ACH return risk, migrating your integration from [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) and Signal Rules provides the following benefits:

- Configure and customize business logic for accepting transactions in the [Signal Dashboard](https://dashboard.plaid.com/signal/), without making any code changes.
- View your transaction return activity and performance via the Dashboard.
- Access [personalized recommendations and scenario modeling](/docs/signal/tuning-rules/) to help tune and optimize your rules.
- If you need to upgrade certain balance checks to using Signal Transaction Scores for lower latency, deeper insights, or more sophisticated rules logic, you can do so without any code changes.

Note: Signal Rules and [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) are only available for use cases involving evaluating potential ACH transactions for return risk. Other Balance use cases should continue to use [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget).

#### Migration steps

Signal Rules will be available soon for all existing Balance customers. In the meantime, if you would like to migrate to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), contact your Account Manager to request access.

1. Configure the [Signal Rules](https://dashboard.plaid.com/signal/risk-profiles) in the Dashboard to create a **Balance-only Ruleset** that matches your existing business logic. For more details, see [Signal Rules](/docs/signal/signal-rules/).
2. Replace your [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) calls with calls to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate). At a minimum, you will need to send the following parameters:

   - The `access_token`
   - The `account_id` of the account you will be debiting
   - The transaction `amount`
   - A `client_transaction_id` of your choosing to uniquely identify the proposed transaction
   - The `ruleset_key` you want to use, if you are using multiple rulesets (will default to using the `default` ruleset if not specified)
3. [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will return a `ruleset.result` such as `ACCEPT` or `REROUTE`, based on your ruleset logic. In your application's code, replace the balance-based logic around accepting or rerouting a proposed transaction to instead use the value of this field.
4. If you still need direct access to real-time balance data (for example, to display in your UI) you can access it via the `core_attributes.available_balance` and `core_attributes.current_balance` fields returned by [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate).
5. [Report whether you proceeded with the transaction](/docs/signal/reporting-returns/) by calling [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport). You only need to do this if the result was `REVIEW`, or if you did not follow the recommendation in the result (e.g. accepted a transaction even if the result was `REROUTE`).
6. If you later receive an ACH return for an accepted transaction, call [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
