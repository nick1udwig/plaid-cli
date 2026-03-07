---
title: "Transactions - Troubleshooting | Plaid Docs"
source_url: "https://plaid.com/docs/transactions/troubleshooting/"
scraped_at: "2026-03-07T22:05:23+00:00"
---

# Troubleshooting Transactions

#### API error responses

If the problem you're encountering involves an API error message, see [Errors](/docs/errors/) for troubleshooting suggestions.

#### No transactions or cursor returned

When calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), a response with no cursor or transactions data is returned.

##### Common causes

- This is expected behavior when the Item has recently been initialized with `transactions` and Plaid has not yet completed fetching of transactions data.

##### Troubleshooting steps

If you have already called [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) on the Item, listen for the `SYNC_UPDATES_AVAILABLE` webhook and call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) again after it has been received. Note that you will need to call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) at least once on an Item to activate the `SYNC_UPDATES_AVAILABLE` webhook.

If you have not already called [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) on the Item, listen for the `INITIAL_UPDATE` webhook and call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) again after it has been received.

Wait a few seconds and try the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) call again.

#### Missing transactions

##### Common causes

- The institution is experiencing an error.
- Processing of the transactions has been delayed.
- The Item is not in a healthy state.
- The transaction is not available on the institution's online banking portal, or appears differently on the online banking portal.

##### Troubleshooting steps

Check the [Plaid status page](https://status.plaid.com) for any known system-wide outages.

Check the institution status via the [Plaid Dashboard](https://dashboard.plaid.com/activity/status) or the [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) endpoint. A degraded status indicates that the issue is known to Plaid and under investigation.

Ensure that the Item is not in an error state by examining the error returned by the failing endpoint, or making an [`/item/get`](/docs/api/items/#itemget) request to get more information. If the Item is in an error state, see the [error guide](/docs/errors/) for the specific error that was encountered.

Ask the user to ensure that the given transaction is visible in the institution's online banking portal. Other sources, such as paper statements, should not be used as the source of truth to compare to Plaid data. Also verify that the transaction is not being reported under a different description. Plaid cleans and standardizes transaction data, so it is possible for a transaction to be returned that does not have the exact description provided by the bank.

Wait two business days and try calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) or [`/transactions/get`](/docs/api/products/transactions/#transactionsget) again.

If the steps above do not resolve the issue, [file a ticket with Plaid support](https://dashboard.plaid.com/support/new/financial-institutions/missing-data/missing-transactions).

#### Incorrect data fields

In some cases, information about a merchant, such as its name or category, may be incorrect.

##### Troubleshooting steps

If the `name` field is incorrect, try the `merchant_name` field, or vice-versa. The `merchant_name` field has a higher level of processing applied to the raw transaction data and may be more likely to match your users' expectations of transaction data.

If the issue is not resolved, or if the category information is incorrect, [file a ticket with Plaid support](https://dashboard.plaid.com/support/new/financial-institutions/faulty-data/faulty-transaction-data) indicating which fields are incorrect, the correct data, and the `transaction_id` of the affected transaction.

#### Limited transaction history is returned

##### Common causes

- By default, Plaid will only return 90 days of transaction history. To request more transaction history, specify the `days_requested` field when initializing with Transactions. Note that this field will only take effect when Transactions is added to an Item for the first time; if you need more than 90 days of transactions history for an Item that has already been initialized with Transactions for 90 days of history, you will need to delete the Item and create a new one.
- The user's institution may provide limited transaction history. For example, Capital One provides only 90 days of transaction history and does not provide pending transactions.

##### Troubleshooting steps

- Specify `days_requested` when calling [`/link/token/create`](/docs/api/link/#linktokencreate) or (if not initializing with Transactions during Link) when calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) or [`/transactions/get`](/docs/api/products/transactions/#transactionsget) for the first time on an Item.
- Verify that the issue is not specific to a particular institution. For details of history available at certain major institutions, see the [Plaid guide to institution-specific OAuth experiences](https://plaid.com/documents/oauth-institution-ux.pdf).

#### Duplicate transactions are returned

##### Common causes

- An end user linked the same bank account to your app multiple times, causing the same transaction data to appear across multiple Items.
- The duplicate transaction data may be an accurate reflection of the end user's transaction activity (e.g. if the end user was double-charged).
- If a transaction is updated, modified, or posted, it will appear as a separate transaction in transaction history. This is expected behavior and your app should have logic to reconcile the two versions of the transaction. For more details, see [Transactions states](/docs/transactions/transactions-data/).
- There may be an error with the data provided by the institution to Plaid or with Plaid's processing of this data.

##### Troubleshooting steps

Confirm that the transaction and duplicate transaction are associated with the same Plaid `account_id`. If the account ids are different, your user likely linked the same account multiple times, causing data for that account to be duplicated in your app. For details on preventing and managing this situation, see [Duplicate Items](/docs/link/duplicate-items/).

Confirm with the end user that the transaction and duplicate transaction do not both appear in their online banking activity.

Confirm that the transaction and duplicate transaction are not pending and posted versions of the same transaction. A posted transaction will have a `pending_transaction_id` field linking it to the pending version of the transaction.

Confirm that the transaction is not a [modified](/docs/api/products/transactions/#transactions-sync-response-modified) version of an existing transaction. When using [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), a modified transaction will appear within the `modified` array.

If the issue is not resolved, [file a ticket with Plaid support](https://dashboard.plaid.com/support/new/financial-institutions/faulty-data/faulty-transaction-data) indicating which transactions are duplicated, and the `transaction_id`s of the affected transaction.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
