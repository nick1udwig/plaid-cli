---
title: "Auth - Micro-deposit Events | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/microdeposit-events/"
scraped_at: "2026-03-07T22:04:32+00:00"
---

# Micro-deposit events

#### Learn how to use Bank Transfers webhooks to receive micro-deposit status updates

#### Overview

If you are using the optional [Same Day Micro-deposits](/docs/auth/coverage/same-day/) verification flow for Auth and want to receive updates on the state of your micro-deposit transfers, you can enable webhooks for Bank Transfer events that notify you of transfer status updates for Plaid-initiated transfers within the ACH network.

#### Bank Transfers webhooks

To enable Bank Transfers webhooks, add your endpoint on the [account webhooks](https://dashboard.plaid.com/developers/webhooks) page of the dashboard. You will need to have received production approval for Auth before being able to add an endpoint.

Bank Transfers webhooks are not part of Plaid's [Transfer](/docs/transfer/) product. All Auth customers have access to Bank Transfer webhooks; it is not required to sign up for Plaid Transfer to use these webhooks.

To confirm that your endpoint has been correctly configured, you can trigger a test webhook via [`/sandbox/bank_transfer/fire_webhook`](/docs/bank-transfers/reference/#sandboxbank_transferfire_webhook). You should receive the payload body specified below.

Bank Transfers webhook body

```
{
  "webhook_type": "BANK_TRANSFERS",
  "webhook_code": "BANK_TRANSFERS_EVENTS_UPDATE"
}
```

Once you have enabled Bank Transfers webhooks, the [`/bank_transfer/event/sync`](/docs/api/products/auth/#bank_transfereventsync) endpoint can be called to discover new ACH events. To know when you should call this endpoint, listen for the [`BANK_TRANSFERS_EVENTS_UPDATE`](/docs/bank-transfers/reference/#bank_transfers_events_update) webhook.
You will receive a Bank Transfers webhook any time you have posted or returned ACH micro-deposit events available. You can also search or filter micro-deposit events using the [`/bank_transfer/event/list`](/docs/api/products/auth/#bank_transfereventlist) endpoint.

Bank Transfers webhooks and the Bank Transfers endpoint will reflect any micro-deposit sent by Plaid, including both Same Day Micro-deposits and Automated Micro-deposits, if enabled. Bank Transfers webhooks and endpoints will only reflect data about ACH events initiated through Plaid. They do not reflect other ACH activity on a linked account.

#### Event types

Once your user successfully completes the micro-deposit Link flow, you will receive an `account_id` in the success callback.
Bank Transfers events also contain an `account_id`, which you should use to connect events to the corresponding user.

##### Pending

A pending event type means that we have a record of the micro-deposit in our systems, but it has not yet been sent.
You may assume that micro-deposits are in a pending state once you [exchange](/docs/api/items/#itempublic_tokenexchange) the corresponding Item’s public token for an access token.
Note that Plaid does not send webhooks for new pending events, but you will still see pending events in event sync responses.

Example pending event

```
{
  "account_id": "3MqrrrP5pWUGgmnvaQlPu19R6DvwPRHwNbLr9",
  "bank_transfer_amount": "0.01",
  "bank_transfer_id": "5645fe0e-bd5e-d8da-828b-e2c7540c69d8",
  "bank_transfer_iso_currency_code": "USD",
  "bank_transfer_type": "credit",
  "direction": "outbound",
  "event_id": 5,
  "event_type": "pending",
  "failure_reason": null,
  "origination_account_id": null,
  "receiver_details": null,
  "timestamp": "2021-03-22T18:52:02Z"
}
```

###### Recommended action

No action needed

##### Posted

For successful micro-deposit transfers, `posted` events are the terminal event type, and no more events will be issued.
Note that the end user may not receive the micro-deposit until several banking hours after the `posted` event.
In addition, a `posted` event does not guarantee that the micro-deposit was successful; if the micro-deposit fails, a `reversed` event will eventually occur some time after the `posted` event.

Example posted event

```
{
  "account_id": "3MqrrrP5pWUGgmnvaQlPu19R6DvwPRHwNbLr9",
  "bank_transfer_amount": "0.01",
  "bank_transfer_id": "5645fe0e-bd5e-d8da-828b-e2c7540c69d8",
  "bank_transfer_iso_currency_code": "USD",
  "bank_transfer_type": "credit",
  "direction": "outbound",
  "event_id": 5,
  "event_type": "posted",
  "failure_reason": null,
  "origination_account_id": null,
  "receiver_details": null,
  "timestamp": "2021-03-22T18:52:02Z"
}
```

###### Recommended Action

If the micro-deposit succeeds, all posted transfers will land in the end user’s account by 8:30am EST on the following banking day.

After the micro-deposit settlement time, Plaid recommends sending the end user an alert (through SMS, email, or push notification) to verify the micro-deposit amounts.

If you have enabled both Same Day and Automated micro-deposits, then before notifying the user, you should first confirm that the event corresponds to a Same Day micro-deposit by checking that the verification status is `pending_manual_verification` (and not `pending_automatic_verification`, which would correspond to an Automated Micro-deposit). Because Plaid handles Automated Micro-deposits without user interaction, it is not necessary to prompt the user during the Automatic Micro-deposit flow.

If you have not already stored the verification status, you can obtain it by calling [`/accounts/get`](/docs/api/accounts/#accountsget).

##### Reversed

A `reversed` event indicates that a micro-deposit attempt has failed.
Reversed events will contain an [ACH return code](/docs/errors/transfer/#ach-return-codes) that indicates why the micro-deposit failed.

Example reversed event

```
{
  "account_id": "bV8WNn73rLI5Ln1MmdErsDn9jv7w37uGMaQvP",
  "bank_transfer_amount": "0.01",
  "bank_transfer_id": "826712b2-c707-cf98-5ba9-13bd3cc2b2f0",
  "bank_transfer_iso_currency_code": "USD",
  "bank_transfer_type": "credit",
  "direction": "outbound",
  "event_id": 5,
  "event_type": "reversed",
  "failure_reason": {
    "ach_return_code": "R03",
    "description": "No account or unable to locate account"
  },
  "origination_account_id": null,
  "receiver_details": null,
  "timestamp": "2021-03-25T21:35:47Z"
}
```

###### Recommended Action

Contact the user with a notification that authentication has failed. Once they return to your application, restart the Link flow to begin another authentication attempt.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
