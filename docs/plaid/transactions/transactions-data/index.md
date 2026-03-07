---
title: "Transactions - Transaction states | Plaid Docs"
source_url: "https://plaid.com/docs/transactions/transactions-data/"
scraped_at: "2026-03-07T22:05:23+00:00"
---

# Transaction states

#### Learn about the differences between pending and posted transactions

#### Pending and posted transactions

There are two types of transactions: pending and posted. A transaction begins its life as a pending transaction, then becomes posted once the funds have actually been transferred. It typically takes about one to five business days for a transaction to move from pending to posted, although it can take up to fourteen days in rare situations.

When a transaction posts, the transition from a pending to posted transaction will be represented through the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) endpoint with the pending transaction's id in the `removed` field of the response and the new posted transaction in the `added` section of the response -- note that these aren't guaranteed to be in the same page, but should happen within the same overall update. If Plaid matches the pending transaction to the new posted transaction, the pending transaction's id will be marked in the `pending_transaction_id` of the posted transaction.

Some institutions, such as Capital One and USAA, do not provide pending transaction data. If no pending transaction is provided, or the pending transaction cannot be matched, the `pending_transaction_id` will be `null`.

Example of /transactions/sync response for transaction posting

```
{
  "added": [
    {
      "transaction_id": "lPNjeW1nR6CDn5okmGQ6hEpMo4lLNoSrzqDje",
      "pending_transaction_id": "no86Eox18VHMvaOVL7gPUM9ap3aR1LsAVZ5nc",
      "account_id": "BxBXxLj1m4HMXBm9WZZmCWVbPjX16EHwv99vp",
      "pending": false,
      "name": "Apple Store",
      "amount": 2307.21,
      /* ... */
    }
  ],
  "removed": [
    {
      "transaction_id": "no86Eox18VHMvaOVL7gPUM9ap3aR1LsAVZ5nc",
      "account_id": "BxBXxLj1m4HMXBm9WZZmCWVbPjX16EHwv99vp",

      /* ... */
    }
  ],
  "modified": []
  /* ... */
}
```

The pending and posted versions of a transaction may not necessarily share the same details: their name and amount may change. For example, the pending charge for a meal at a restaurant may not include a tip, but the posted version will include the final amount spent, including the tip. In some cases, a pending transaction may not convert to a posted transaction at all and will simply disappear; this can happen, for example, if the pending transaction was used as an "authorization hold," which is a sort of a deposit for a potential future transaction, frequently used by gas stations, hotels, and rental-car companies. Pending transactions are short-lived and frequently altered or removed by the institution before finally settling as a posted transaction.

Note that while transactions will rarely change once they have posted, a posted transaction cannot necessarily be considered immutable. For example, a refund or a recategorization of a transaction by the institution could cause a previously posted transaction to change. This is why it's important to apply all `modified`, `added`, and `removed` updates surfaced through [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) in order to maintain consistency with the underlying account data.

#### Testing pending and posted transactions

To test dynamic transactions behavior in Sandbox, use the test user `user_transactions_dynamic` with any non-blank password. You can create a public token for this user using [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate); if using the interactive Link flow, you must use a non-OAuth test institution such as First Platypus Bank (`ins_109508`). An Item associated with this user will be created with 50 transactions in both the `pending` and `posted` state. You can then simulate Plaid receiving more data for this Item by calling [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh). This will generate new `pending` transactions, all previously `pending` transactions will be moved to `posted`, and the amount of one previous transaction will be incremented by $1.00. All appropriate transaction webhooks will also be fired at this time.

Remember that in Production, you don't need to call [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) unless you want to force Plaid to update its transaction data outside of its usual schedule.

#### Example code in Plaid Pattern

For a real-life example, see [update\_transactions.js](https://github.com/plaid/pattern/blob/master/server/update_transactions.js). This file demonstrates code for handling transaction states in the Node-based [Plaid Pattern](https://github.com/plaid/pattern) sample app.

#### Transaction state changes with /transactions/get

The content in this section and below applies only to existing integrations using the [`/transactions/get`](/docs/api/products/transactions/#transactionsget) endpoint. It is recommended that any new integrations use [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) instead of [`/transactions/get`](/docs/api/products/transactions/#transactionsget), for easier and simpler handling of Transaction state changes.

**Reconciling transactions**

[`/transactions/get`](/docs/api/products/transactions/#transactionsget) returns both pending and posted transactions; however, some institutions do not provide pending transaction data and will only supply posted transactions. The `pending` boolean field in the transaction indicates whether the transaction is pending or posted.

Plaid does not model the transition of a pending to posted transaction as a state change for an existing transaction; instead, the posted transaction is a new transaction with a `pending_transaction_id` field that matches it to a corresponding pending transaction. When a pending transaction is converted to a posted transaction, Plaid removes the pending transaction, sends a [`TRANSACTIONS_REMOVED`](/docs/api/products/transactions/#transactions_removed) webhook, and returns the new, posted transaction. The posted transaction will have a `pending_transaction_id` field whose value is the `transaction_id` of the now-removed pending transaction. The posted transaction’s date will reflect the date the transaction was posted, which may differ from the date on which the transaction actually occurred.

In some rare cases, Plaid will fail to match a posted transaction to its pending counterpart. On such occasions, the posted transaction will be returned without a `pending_transaction_id` field, and its pending transaction is removed.

**Handling pending and posted transactions**

To manage the movement of a transaction from pending to posted state, you will need to handle the `TRANSACTIONS_REMOVED` webhook to identify the removed transactions, then delete them from your records. For detailed instructions, see [Transactions webhooks](/docs/api/products/transactions/#transactions_removed).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
