---
title: "Transactions - Introduction to Transactions | Plaid Docs"
source_url: "https://plaid.com/docs/transactions/"
scraped_at: "2026-03-07T22:05:20+00:00"
---

# Introduction to Transactions

#### Retrieve up to 24 months of transaction data and stay up-to-date with webhooks.

Get started with Transactions

[API Reference](/docs/api/products/transactions/)[Quickstart](/docs/quickstart/)[Demo](https://plaid.coastdemo.com/share/67e1889391fb8841d04eb6ba?zoom=100)

#### Overview

Transactions data can be useful for many different applications, including
personal finance management, expense reporting, cash flow modeling, risk
analysis, and more. Plaid's Transactions product allows you to access a user's
transaction history for `depository` type accounts such as checking and savings
accounts, `credit` type accounts such as credit cards, and student loan
accounts. For transaction history from investment accounts, use Plaid's
[Investments](/docs/investments/) product.

Transactions data available via [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) includes transaction date, amount, category, merchant,
location, and more. Transaction data is lightly cleaned to populate the `name`
field, and more thoroughly processed to populate the `merchant_name` field. For example data, see the [Transactions API reference](https://plaid.com/docs/api/products/transactions/#transactions-sync-response-transactions-update-status).

##### Typical fill rates for selected fields

Below are typical fill rates for selected fields returned by Transactions. Not all fields are included in the table below.

| Field | Typical fill rate |
| --- | --- |
| Amount | 100% |
| Date | 100% |
| Description | 100% |
| Merchant name | 97%\* |
| Category (`personal_finance_category`) | 95% |

\*Denominator excludes transactions that do not have an associated merchant, such as cash transactions, direct deposits, or bank fees.

#### Integration overview

The steps below show an overview of how to integrate Transactions. For a detailed, step-by-step view, you can also watch our full-length, comprehensive tutorial walkthrough.

1. Call [`/link/token/create`](/docs/api/link/#linktokencreate). Along with any other parameters you specify, make sure to include the following:

   - The `products` array should include `transactions`.
   - Specify a `webhook` receiver endpoint to receive transactions updates.
   - Specify a value for `transactions.days_requested` corresponding to the amount of transaction history your integration needs. The more transaction history is requested, the longer the historical update poll will take. The default value is 90; the maximum value is 730.
2. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see the [Link documentation](/docs/link/).
3. Once the end user has completed the Link flow, exchange the `public_token` for an `access_token` by calling [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange).
4. Create a method to call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync). This method must do the following:

   - The first time [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) is called on an Item, call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) with no cursor value specified
   - Save the `next_cursor` value returned by [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) to use as an input parameter in the next call to [`/transactions/sync`](/docs/api/products/transactions/#transactionssync)
   - If `has_more` in the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) response is `true`, handle paginated results by temporarily preserving the current cursor, then calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) with the newly returned cursor. After each [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) call, call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) with the newly returned cursor from the previous [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) call until `has_more` is `false`.
   - When handling paginated results, if `TRANSACTIONS_SYNC_MUTATION_DURING_PAGINATION` error is returned during pagination, restart the process in the previous bullet with the old cursor that was temporarily preserved. Once `has_more` is `false`, it is safe to stop preserving the old cursor.

   For an example, see [fetchNewSyncData](https://github.com/plaid/tutorial-resources/blob/main/transactions/finished/server/routes/transactions.js#L31) and [syncTransactions](https://github.com/plaid/tutorial-resources/blob/main/transactions/finished/server/routes/transactions.js#L110).
5. Call the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) wrapper you created in the previous step, passing in the `access_token`, in order to activate the Item for the [`SYNC_UPDATES_AVAILABLE`](/docs/api/products/transactions/#sync_updates_available) webhook. It is common that no transactions will be returned during this first call, as it takes Plaid some time to fetch initial transactions.
6. (Optional) Wait for the [`SYNC_UPDATES_AVAILABLE`](/docs/api/products/transactions/#sync_updates_available) webhook to fire with the `initial_update_complete` field set to `true` and the `historical_update_complete` set to `false`. When this occurs, you may optionally call your [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) wrapper to obtain the most recent 30 days of data. If your end-user is expecting to see data load into your app in real-time, this can help improve the responsiveness of your app.
7. Wait for the [`SYNC_UPDATES_AVAILABLE`](/docs/api/products/transactions/#sync_updates_available) webhook to fire with the `historical_update_complete` field set to `true`. Once it does, call your [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) wrapper to obtain all available updates.
8. From this point forward, Plaid will periodically check for transactions data for your Item. When it detects that new transactions data is available, the [`SYNC_UPDATES_AVAILABLE`](/docs/api/products/transactions/#sync_updates_available) webhook will fire again. When it does, call your [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) wrapper to receive all updates since the `next_cursor` value. In addition to added transactions, these subsequent updates may also include `removed` or `modified` transactions. For details, see [Transactions updates](/docs/transactions/#transactions-updates).
9. (Optional) If you would like to check for updates more frequently than Plaid's default schedule, call [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) to force an on-demand update via the optional [Transactions Refresh](/docs/transactions/#transactions-refresh) add-on. If any new transaction data is available, a [`SYNC_UPDATES_AVAILABLE`](/docs/api/products/transactions/#sync_updates_available) webhook will fire.

#### Transactions updates

Transactions data is not static. As time passes, your users will make new
transactions, and transactions they made in the past will change as they are
processed by the financial institution. To learn more about how transactions are
processed and can change, see
[Transaction states](/docs/transactions/transactions-data/).

Plaid checks for updated transactions regularly, and uses
webhooks to notify you of any changes so you can keep your app up to date. For
more detail on how to listen and respond to transaction update webhooks, see
[Transaction webhooks](/docs/transactions/webhooks/).

The frequency of transactions update checks is typically one or more times a day. The exact frequency will depend on the institution. To
learn when an Item was last checked for updates, you can view the Item in the
[Item Debugger](https://dashboard.plaid.com/activity/debugger). If you would
like to display this information in your app's UI to help users understand the
freshness of their data, it can also be retrieved via API, using the [`/item/get`](/docs/api/items/#itemget)
endpoint.

##### Transactions refresh

You can also request an update on-demand via the [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh)
endpoint, which is available as an add-on for Transactions customers. To request
access to this endpoint, submit a
[product access request](https://dashboard.plaid.com/settings/team/products) or contact
your Plaid account manager.

#### Recurring transactions

If your app involves personal financial management functionality, you may want
to view a summary of a user's inflows and outflows. The
[`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget) endpoint provides a summary of the recurring
outflow and inflow streams and includes insights about each recurring stream
including the category, merchant, last amount, and more.
[`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget) is available as an add-on for Transactions
customers in the US, Canada, and UK. To request access to this endpoint, submit
a [product access request](https://dashboard.plaid.com/settings/team/products) or contact
your Plaid account manager.

#### Sample app demo and code

The [Transactions sample app](https://github.com/plaid/tutorial-resources/tree/main/transactions) is a simple app designed to accompany the [Transactions YouTube tutorial](https://www.youtube.com/watch?v=Pin0-ceDKcI).

For a more robust example of an app that incorporates transactions, along with
sample code for transactions reconciliation, see the Node-based
[Plaid Pattern](https://github.com/plaid/pattern) sample app. The Pattern app is also [hosted](http://pattern.plaid.com) and can be used as a demo app.

#### Testing Transactions in Sandbox

Plaid provides several special users for testing Transactions. One example is `user_transactions_dynamic` (password: any non-blank password). You can create a public token for this user using [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate); if using the interactive Link flow, you must use a non-OAuth test institution such as First Platypus Bank (`ins_109508`).

Unlike `user_good`, `user_transactions_dynamic` contains realistic, dynamic transactions data and can be used together with the [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) endpoint to simulate Transactions updates and trigger webhooks in Sandbox. For more details on how to simulate transaction activity with this user, see [Testing pending and posted transactions](/docs/transactions/transactions-data/#testing-pending-and-posted-transactions).

If you are using the [Recurring Transactions add-on](/docs/api/products/transactions/#transactionsrecurringget), you can also use `user_transactions_dynamic` to test recurring transactions in Sandbox. This user has six months of recurring transactions.

To add your own custom transactions to `user_transactions_dynamic`, call [`/sandbox/transactions/create`](/docs/api/sandbox/#sandboxtransactionscreate). This simulates a [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) call where Plaid will see those custom transactions as new transactions. All corresponding webhooks will fire.

For persona-based testing, Plaid also provides `user_ewa_user`, `user_yuppie`, and `user_small_business` users. These accounts simulate real life personas, so new transactions will appear at a more realistic rate and will not appear every time [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) is called. These users have three months of transactions, including some recurring transactions. These users also require a non-OAuth test institution.

All users below should be used with any non-blank password and a non-OAuth institution, such as First Platypus Bank (`ins_109508`).

| User Name | Description | Transaction Data | New txns on refresh | Recurring |
| --- | --- | --- | --- | --- |
| `user_transactions_dynamic` | Dynamic data | Triggers webhooks, allows adding custom transactions | Always | Yes (6 months) |
| `user_ewa_user` | Earned-wage access persona | 3 months of realistic data | Sometimes | Yes (some) |
| `user_yuppie` | Young affluent professional | 3 months of realistic data | Sometimes | Yes (some) |
| `user_small_business` | Small business persona | 3 months of realistic data | Sometimes | Yes (some) |

#### Transactions pricing

Transactions and the optional Recurring Transactions add-on are billed on a [subscription model](/docs/account/billing/#subscription-fee). The optional Transactions Refresh add-on is billed on a [per-request model](/docs/account/billing/#per-request-flat-fee). To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

#### Next steps

To get started building with Transactions, see
[Add Transactions to your App](/docs/transactions/add-to-app/).

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
