---
title: "Transactions - Transactions Sync migration guide | Plaid Docs"
source_url: "https://plaid.com/docs/transactions/sync-migration/"
scraped_at: "2026-03-07T22:05:22+00:00"
---

# Transactions Sync migration guide

#### Learn how to migrate from the /transactions/get endpoint to the /transactions/sync endpoint

#### Overview

[`/transactions/sync`](/docs/api/products/transactions/#transactionssync) is a newer endpoint that replaces [`/transactions/get`](/docs/api/products/transactions/#transactionsget) and provides a simpler and easier model for managing transactions updates. While [`/transactions/get`](/docs/api/products/transactions/#transactionsget) provides all transactions within a date range, [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) instead uses a cursor to provide all new, modified, and removed transactions that occurred since your previous request. With this cursor-based pagination, you do not need to worry about making redundant API calls to avoid missing transactions. Updates returned by [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) can be patched into your database, allowing you to avoid a complex transaction reconciliation process or having to keep track of which updates have already been applied.

This guide outlines how to update your existing [`/transactions/get`](/docs/api/products/transactions/#transactionsget) integration to use the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) endpoint and simplify your Plaid integration.

Looking for an example in code? Check out [Pattern on GitHub](https://github.com/plaid/pattern) for a complete, best-practice implementation of the Transactions Sync API within a sample app.

#### Update your client library

If you are using client libraries, you may need to update your current library to use [`/transactions/sync`](/docs/api/products/transactions/#transactionssync). The following are the minimum Plaid client library versions required to support [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) for each language:

- Python: 9.4.0
- Node: 10.4.0
- Ruby: 15.5.0
- Java: 11.3.0
- Go: 3.4.0

Detailed upgrade notes are language-specific may be found in the README and Changelog of the specific library. See the library's repo on the [Plaid GitHub](https://github.com/plaid) for more information.

#### Update callsites and pagination logic

Replace all instances of [`/transactions/get`](/docs/api/products/transactions/#transactionsget) with [`/transactions/sync`](/docs/api/products/transactions/#transactionssync). [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) has a slightly different call signature from [`/transactions/get`](/docs/api/products/transactions/#transactionsget) and does not have the `count` parameter inside the `options` object and uses a `cursor` instead of a `start_date` and `end_date`. Pagination logic is also different and relies on the `has_more` flag instead of the `transactions_count` value. Note that when requesting paginated updates with [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), unlike when using [`/transactions/get`](/docs/api/products/transactions/#transactionsget), it is important to retrieve all available updates before persisting the transactions updates to your database.

Unlike [`/transactions/get`](/docs/api/products/transactions/#transactionsget), [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) does not allow specifying a date range within which to retrieve transactions. If your implementation requires getting transactions within a certain date range, implement transaction filtering after calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync).

For copy-and-pastable examples of how to call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), including complete pagination logic, see the API reference code samples for [`/transactions/sync`](/docs/api/products/transactions/#transactionssync).

If a call to [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) fails when retrieving a paginated update as a result of the `TRANSACTIONS_SYNC_MUTATION_DURING_PAGINATION` error, the entire pagination request loop must be restarted beginning with the cursor for the first page of the update, rather than retrying only the single request that failed.

#### Update callsites for Item data

Unlike [`/transactions/get`](/docs/api/products/transactions/#transactionsget), [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) does not return an Item object. If your app relies on getting Item data, such as Item health status, use [`/item/get`](/docs/api/items/#itemget).

#### Update webhook handlers

When using [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), you should not listen for the webhooks `HISTORICAL_UPDATE`, `DEFAULT_UPDATE`, `INITIAL_UPDATE`, or `TRANSACTIONS_REMOVED`. While these webhooks will still be sent in order to maintain backwards compatibility, they are not required for the business logic used by [`/transactions/sync`](/docs/api/products/transactions/#transactionssync).

Instead, update your webhook handlers to listen for the [`SYNC_UPDATES_AVAILABLE`](/docs/api/products/transactions/#sync_updates_available) webhook and to call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) when this webhook is received.

#### Update initial call trigger

Unlike the [`/transactions/get`](/docs/api/products/transactions/#transactionsget) webhooks, the `SYNC_UPDATES_AVAILABLE` webhook will not be fired for an Item unless [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) has been called at least once for that Item. For this reason, you must call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) at least once before any sync webhook is received. After that point, rely on the `SYNC_UPDATES_AVAILABLE` webhook.

Unlike [`/transactions/get`](/docs/api/products/transactions/#transactionsget), [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) will not return the `PRODUCT_NOT_READY` error if transactions data is not yet ready when [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) is first called. Instead, you will receive a response with no transactions and a null cursor. Even if no transactions data is available, this call will still initialize the `SYNC_UPDATES_AVAILABLE` webhook, and it will fire once data becomes available.

The first call to [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) once historical updates are available will often have substantially higher latency (up to 8x) than the equivalent call in a [`/transactions/get`](/docs/api/products/transactions/#transactionsget)-based implementation. Depending on your application's logic, you may need to adjust user-facing messaging or hard-coded timeout settings.

#### Update transaction reconciliation logic

The response to [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) includes the patches you will need to apply in the `added`, `removed`, and `modified` arrays within its response. You should apply these to your transactions records. Any additional logic required to fetch or reconcile transactions data can be removed.

#### Migrating existing Items

You likely already have transactions stored for existing Items. If you onboard an existing Item onto [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) with `"cursor": ""` in the request body, the endpoint will return all historical transactions data associated with that Item up until the time of the API call (as "adds"). You may reconcile these with your stored copy of transactions to ensure that it reflects the Item's true state.

If you have a large number of Items to update, this reconciliation process may be slow and generate excessive system load. One other option for onboarding existing Items onto [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) is using `"cursor": "now"` in the request body. The endpoint will return a response containing no transaction updates, but only a cursor that will allow you to retrieve all transactions updates associated with that Item going forward, after the time of the API call. Accordingly, you should ensure that your local copy of transactions for an Item is up-to-date at the time you call [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) with `"cursor": "now"` for it, or else any transaction updates that occurred between the time that you last pulled fresh data and the time of your [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) call may be missing.

`"cursor": "now"` will work exactly like a cursor that was found by starting with `"cursor": ""` and paginating through all updates, with the only difference being that a transaction created before, but modified after, those requests would be returned as "added" if using `"cursor": "now"`, and "modified" if using `"cursor": ""`.

If you ever want to completely rebuild your local copy of transactions for an Item previously onboarded with `"cursor": "now"`, you may still do so with `"cursor": ""`.

Note that we strongly recommend that this cursor only be used with Items for which you've already used with [`/transactions/get`](/docs/api/products/transactions/#transactionsget), and not any new Items, which should always be onboarded with `"cursor": ""`.

#### Test your integration

You can perform basic testing of your integration's business logic in Sandbox, using the [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook) endpoint to simulate `SYNC_UPDATES_AVAILABLE`. If this testing succeeds, you should then test your integration with internal test accounts before releasing it to your full userbase.

#### Example code

For a full working example of a Plaid-powered app using [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), see [Plaid Pattern](https://github.com/plaid/pattern/tree/master/server).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
