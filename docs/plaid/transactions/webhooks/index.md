---
title: "Transactions - Transactions webhooks | Plaid Docs"
source_url: "https://plaid.com/docs/transactions/webhooks/"
scraped_at: "2026-03-07T22:05:23+00:00"
---

# Transactions webhooks

#### Listen for Transaction webhooks to learn when transactions are ready for retrieval or when transactions have been updated.

#### Introduction

Webhooks are a useful part of the Transactions product that notifies you when Plaid has new or updated transaction information. This guide will explain how to use webhooks to make sure you have up-to-date transaction history.

#### Configuring Link for transactions webhooks

Before you can listen to webhooks, you must first set up an endpoint and tell Plaid where to find it.
To tell Plaid where to send its webhooks, send your webhook endpoint URL as an optional argument via the
`webhook` parameter to [`/link/token/create`](/docs/api/link/#linktokencreate).

You must also initialize your Item with Transactions by including `transactions` in the `products` array provided to [`/link/token/create`](/docs/api/link/#linktokencreate). If you do not do this, Plaid will not attempt to retrieve any transactions for your Item until after [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) or [`/transactions/get`](/docs/api/products/transactions/#transactionsget) is called for the first time. For more information, see [Choosing how to initialize products](/docs/link/initializing-products/).

#### Integrating the update notification webhook

After [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) is called for the first time on an Item, [`SYNC_UPDATES_AVAILABLE`](https://plaid.com/docs/api/products/transactions/#sync_updates_available) webhooks will begin to be sent to the configured destination endpoint.

This webhook will fire whenever any change has happened to the Item's transactions. The changes can then be retrieved by calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) with the `cursor` from your last sync call to this Item.

If at least 30 days of history is available with an update, the `initial_update_complete` parameter in the body of the `SYNC_UPDATES_AVAILABLE` webhook will be `true`. Similarly, `historical_update_complete` will be `true` if the full history (up to 24 months) is available.

For a real-life example that illustrates how to handle this webhook, see [handleTransactionsWebhook.js](https://github.com/plaid/pattern/blob/master/server/webhookHandlers/handleTransactionsWebhook.js), which contains the webhook handling code for the Node-based [Plaid Pattern](https://github.com/plaid/pattern) app.

#### Forcing transactions refresh

Sometimes, checking for transactions a few times a day is not good enough. For example, you might want to build a refresh button in your app that allows your user to check for updated transactions on-demand. To accomplish this, you can use the [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) product. After a successful call to [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh), if there are new updates, `SYNC_UPDATES_AVAILABLE` will be fired (along with `DEFAULT_UPDATE` and, if applicable, `TRANSACTIONS_REMOVED`).

#### Instructions for integrations using /transactions/get

The content in this section and below applies only to existing integrations using the [`/transactions/get`](/docs/api/products/transactions/#transactionsget) endpoint. It is recommended that any new integrations use [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) instead of [`/transactions/get`](/docs/api/products/transactions/#transactionsget), for easier and simpler handling of transaction state changes. For information on migrating an existing [`/transactions/get`](/docs/api/products/transactions/#transactionsget) integration to [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), see the [Transactions Sync migration guide](/docs/transactions/sync-migration/).

When you first connect an Item in Link, transactions data will not immediately be available. [`INITIAL_UPDATE`](/docs/api/products/transactions/#initial_update)
and [`HISTORICAL_UPDATE`](/docs/api/products/transactions/#historical_update) are both webhooks that fire shortly after an Item has been initially linked
and initialized with the Transactions product. These webhooks will let you know when your transactions
are ready. `INITIAL_UPDATE` fires first, after Plaid has successfully pulled 30 days of transactions for an Item.
The `HISTORICAL_UPDATE` webhook fires next, once all historical transactions data is available. `INITIAL_UPDATE` typically fires within 10 seconds, and `HISTORICAL_UPDATE` within 1 minute, although these webhooks may take 2 minutes or more. The time required for the webhooks to fire will depend on the institution, as well as on the number of transactions being processed.

If you attempt to call [`/transactions/get`](/docs/api/products/transactions/#transactionsget) before `INITIAL_UPDATE` has fired, you will get a
[`PRODUCT_NOT_READY`](/docs/errors/item/#product_not_ready) error. If you attempt to call [`/transactions/get`](/docs/api/products/transactions/#transactionsget) after `INITIAL_UPDATE`
has fired, but before `HISTORICAL_UPDATE` has fired, you will only be able to receive the last 30 days of
transaction data. If you did not initialize the Item with Transactions, your first call to [`/transactions/get`](/docs/api/products/transactions/#transactionsget)
will result in a `PRODUCT_NOT_READY` error and kick off the process of readying transactions. You can then
listen for the `INITIAL_UPDATE` or `HISTORICAL_UPDATE` webhooks to begin receiving transactions.

**Updating transactions**

Plaid fires two types of webhooks that provide information about changes to transaction data: [`DEFAULT_UPDATE`](/docs/api/products/transactions/#default_update) and [`TRANSACTIONS_REMOVED`](/docs/api/products/transactions/#transactions_removed).

**Adding new transactions**

The [`DEFAULT_UPDATE`](/docs/api/products/transactions/#default_update) webhook fires when new transactions are available. Typically, Plaid will check for transactions at a frequency ranging from one to four times per day, depending on factors such as the institution and account type. If new transactions are available, the `DEFAULT_UPDATE` webhook will fire.

To reflect up-to-date transactions for a user in your app, handle the `DEFAULT_UPDATE` webhook by fetching more transactions. We recommend fetching about 7-14 days of transactions in response to `DEFAULT_UPDATE`. This is typically enough history to ensure that you haven't missed any transactions, but not so much that performance or rate limiting is likely to be a problem.

Once you've fetched these transactions, you will need to identify which transactions are new and which are duplicates of existing data that you have. You should not rely on the number in the webhook's `new_transactions` field to identify duplicates, since it can be unreliable. For example, new transactions may arrive between your receipt of the webhook and your call to [`/transactions/get`](/docs/api/products/transactions/#transactionsget). Instead, compare the `transaction_id` field of each newly fetched transaction to the `transaction_id` fields of your existing transactions, and skip the ones that you already have. For an example, see [Plaid Pattern](https://github.com/plaid/pattern/blob/master/server/webhookHandlers/handleTransactionsWebhook.js#L176).

**Removing stale transactions**

The [`TRANSACTIONS_REMOVED`](/docs/api/products/transactions/#transactions_removed) webhook fires when transactions have been removed. The most common reason for this is in the case of pending transactions. In general, transactions start out as pending transactions and then move to the posted state one to two business days later. When Plaid detects that a transaction has moved from pending to posted state, the pending transaction as returned by [`/transactions/get`](/docs/api/products/transactions/#transactionsget) is not modified. Instead, the pending transaction is removed, and a new transaction is added, representing the posted transaction. For a detailed explanation of pending and posted transactions and how they are handled by Plaid, see [Transaction states](/docs/transactions/transactions-data/).

Pending transactions can also be removed when they are canceled by the bank or payment processor. A transaction may be removed if its details are changed by the bank so extensively that Plaid can no longer recognize the new and old versions of the transaction as being the same (e.g., a transaction amount and description both being changed simultaneously). In this case, the old transaction will be deleted and a new transaction with the new details will be added. This "transaction churn" can affect both pending and posted transactions.

The `TRANSACTIONS_REMOVED` webhook contains the transaction IDs of the removed transactions, which you can use to identify and remove the corresponding transactions in your own application to avoid presenting duplicated or inaccurate data. If you encounter any problems with the webhook, you can also manually query transaction history for deleted transactions using logic similar to that recommended for handling the `DEFAULT_UPDATE` webhook.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
