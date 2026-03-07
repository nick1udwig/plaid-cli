---
title: "Transfer - Tracking transfer status | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/reconciling-transfers/"
scraped_at: "2026-03-07T22:05:25+00:00"
---

# Tracking transfer status

#### Monitor for updates to transfers, sweeps, and refunds

#### Transfer status lifecycle

All transfers initiated via ACH (including refunds and sweeps) move through a set of `status`es that correspond to phases in the ACH lifecycle. Understanding these lets you forecast funds availability and surface accurate messaging to your end users.

| Status | Typical timing after /transfer/create | ACH Network Phase | Business significance / recommended actions |
| --- | --- | --- | --- |
| pending | Instant (<1s) | Payment instruction received by Plaid and queued for processing | Transfer can still be canceled until it moves to posted. Safe to show "processing" to users. |
| posted | At cutoff of applicable ACH window (e.g. 9:35 AM, 1:50 PM, 3:00 PM, or 8:30 PM ET) | Payment instruction released to the ACH Network via ODFI | Transfer has been submitted to the payment network and can no longer be canceled. |
| settled | Same Day ACH: ~8:00 PM ET T; Standard ACH: ~9:00 AM ET T+1 | The Federal Reserve settles funds between sending and receiving banks | For debits: funds have been withdrawn from the debited account and settled to Plaid's FBO bank account. For credits: funds have been withdrawn from Plaid's FBO bank account and settled to the receiving institution. |
| Funds available | 2-5 business days after settlement | N/A | For debits only. Funds released from pending and applied to ledger's available balance. Debits that do not return will remain in this status. |
| failed | Seconds to minutes after a failure | Entry rejected before leaving Plaid/ODFI (e.g. compliance or validation failure) | No funds movement occurred. Investigate cause and optionally retry. |
| cancelled | Seconds after calling [`/transfer/cancel`](/docs/api/products/transfer/initiating-transfers/#transfercancel) | Client-initiated cancellation accepted | The transfer will not be submitted to the payment network. |
| returned | Typically T+2 - T+5 (can be up to 60 days for consumer debits) | Receiving FI returns the debit or credit with a return code | For ACH debits: funds are withdrawn from Ledger (pending or available depending on timing). For ACH credits: funds are returned to the Ledger’s available balance. Take action per return code. |

#### Event monitoring

Plaid creates a transfer event any time the `transfer.status` changes. For example, when a transfer is sent to the payment network, the `transfer.status` moves to `posted` and a `posted` event is emitted. By monitoring transfer events, you can stay informed about their current status and notify customers in case of a canceled, failed, or returned transfer. When `transfer.status` moves to `settled`, you can expect that the consumer can see the transaction reflected in their personal bank account. Likewise, if a settled or posted transaction experiences an ACH return, a `returned` event will be emitted as the `transfer.status` is updated; for more details, see [Troubleshooting ACH returns](https://plaid.com/docs/transfer/troubleshooting/#ach-returns).

Events are also emitted for changes to a sweep's and refund's `status` property. These events follow an `[object].[status]` format, such as `sweep.posted` and `refund.posted`.

For a list of all event types and descriptions, see the [API Reference](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-timestamp).

#### Ingesting event updates

Most integrations proactively monitor all events for every transfer. This allows you to respond to transfer events with business logic operations, such as:

- Kicking off the fulfillment of an order once the transfer has settled
- Making funds available to your end consumers for use in your application
- Monitoring returns to know when to claw these services back, or retry the transfer

To do this, set up Transfer webhooks to listen for updates as they happen. You must register a URL to enable webhooks to be sent.

You can do this in the [webhook settings page](https://dashboard.plaid.com/team/webhooks) of the Plaid Dashboard. Click **New Webhook** and specify a webhook URL for a "Transfer Webhook". You must be enabled for Production access to Transfer in order to access this option. To confirm that your endpoint has been correctly configured, you can trigger a test webhook via [`/sandbox/transfer/fire_webhook`](/docs/api/sandbox/#sandboxtransferfire_webhook). Only one webhook URL can be set per environment at a time; if you register multiple Transfer webhook receiver endpoints for a given environment, the webhooks will be sent to only one of the registered URLs.

Now, every time there are new transfer events, Plaid will fire a notification webhook.

Example webhook payload

```
{
  "webhook_type": "TRANSFER",
  "webhook_code": "TRANSFER_EVENTS_UPDATE",
  "environment": "production"
}
```

To receive details about the event, call [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync).

Example request payload

```
{
  # Return the next 20 transfer events after the transfer event with id 4
  "after_id": 4,
  "count": 20
}
```

You can then store the highest `event_id` returned by the response and use that value as the `after_id` the next time you call [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) to get only the new events.

Note that webhooks don't contain any identifying information about what transfer has updated; only that an update happened. As an alternative to listening to webhooks, your application could also call [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) on a regular basis to process the most recent batch of Transfer events.

For a real-life example of an app that incorporates the transfer webhook and tests it using the [`/sandbox/transfer/fire_webhook`](/docs/api/sandbox/#sandboxtransferfire_webhook) endpoint, see the Node-based [Plaid Pattern Transfer](https://github.com/plaid/pattern-transfers) sample app. Pattern Transfer is a sample subscriptions payment app that enables ACH bank transfers. The Transfer webhook handler can be found in [handleTransferWebhook.js](https://github.com/plaid/pattern-transfers/blob/master/server/webhookHandlers/handleTransferWebhook.js) and the test which fires the webhook can be found at [events.js](https://github.com/plaid/pattern-transfers/blob/master/server/routes/events.js).

#### Filtering for specific events

Calling [`/transfer/event/list`](/docs/api/products/transfer/reading-transfers/#transfereventlist) will get a list of transfer events based on specified filter criteria. For example, you could search for all events for a specific `transfer_id`. If you do not specify any filter criteria, this endpoint will return the latest 25 transfer events.

You can apply filters to only fetch specific event types, events for a specific transfer type, a specific sweep, etc.

Example transfer event payload

```
{
  "transfer_events": [
    {
      "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "transfer_amount": "12.34",
      "transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
      "transfer_type": "credit",
      "event_id": 1,
      "event_type": "pending",
      "failure_reason": null,
      "origination_account_id": null,
      "originator_client_id": null,
      "refund_id": null,
      "sweep_amount": null,
      "sweep_id": null,
      "timestamp": "2019-12-09T17:27:15Z"
    }
  ],
  "request_id": "mdqfuVxeoza6mhu"
}
```

#### Reconciling sweeps with your bank account

As Plaid moves money in and out of your business account as you process transfers and cashout the Plaid Ledger balance, you might want to match the account activity in your bank account with the associated transfers.

Plaid will deposit or draw money from your business checking account in the form of a [sweep](/docs/transfer/creating-transfers/#sweeping-funds-to-funding-accounts). This means that any time you are interacting with your bank statement, you are viewing sweeps, not specific transfers.

To match an entry in your bank account with a sweep in Plaid's records, Plaid ensures the first 8 characters of the sweep's `sweep_id` will show up on your bank statements. For example, consider the following entries in your bank account from Plaid:

| Entry | Amount | Date |
| --- | --- | --- |
| PLAID 6c036ea0 CCD | -$5,264.62 | November 18, 2022 |
| PLAID ae42c210 CCD | $2,367.80 | November 16, 2022 |
| PLAID 550c85fc CCD | $6,007.49 | November 10, 2022 |

You can use this 8 character string from your bank statement to search for the sweep via [`/transfer/sweep/get`](/docs/api/products/transfer/reading-transfers/#transfersweepget), or in your Plaid Transfer dashboard.

To follow the lifecycle of a sweep, and monitor funds coming into and out of your business checking account due to Plaid Transfer activity, observe the `sweep.*` events in the [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) endpoint. To view all sweeps, use [`/transfer/sweep/list`](/docs/api/products/transfer/reading-transfers/#transfersweeplist).

To see if a given debit transfer has been included in a sweep transaction, check the transfer's `status` field. If it has been swept, it will have the `funds_available` status.

#### Performing financial reconciliation audits

For information on performing financial reconciliation audits, see [Report extraction](/docs/transfer/dashboard/#report-extraction).

[#### Flow of funds

Understand how your money moves through the Plaid network

Learn more](/docs/transfer/flow-of-funds/)

#### Flow of funds

Understand how your money moves through the Plaid network

[Learn more](/docs/transfer/flow-of-funds/)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
