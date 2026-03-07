---
title: "Payments (Europe) - Payment Status | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/payment-status/"
scraped_at: "2026-03-07T22:05:11+00:00"
---

# Payment Status

#### Verify payment status and track settlement using webhooks

#### Tracking payment status

Once the payment has been authorised by the end user and the Link flow is completed, the `onSuccess` callback is invoked, signaling that the payment status is now `PAYMENT_STATUS_AUTHORISING`.

From this point on, you can track the payment status using the [`PAYMENT_STATUS_UPDATE` webhook](/docs/api/products/payment-initiation/#payment_status_update), which Plaid sends to the [configured webhook URL](/docs/api/link/#link-token-create-request-webhook) in real time to indicate that the status of an initiated payment has changed. (For more information on how to implement your webhook listener, see the [webhook documentation](/docs/api/webhooks/)).

Sample PAYMENT\_STATUS\_UPDATE webhook

```
{
  "webhook_type": "PAYMENT_INITIATION",
  "webhook_code": "PAYMENT_STATUS_UPDATE",
  "payment_id": "payment-id-production-2ba30780-d549-4335-b1fe-c2a938aa39d2",
  "new_payment_status": "PAYMENT_STATUS_AUTHORISING",
  "old_payment_status": "PAYMENT_STATUS_INPUT_NEEDED",
  "original_reference": "Account Funding 99744",
  "adjusted_reference": "Account Funding 99",
  "original_start_date": "2017-09-14",
  "adjusted_start_date": "2017-09-15",
  "timestamp": "2017-09-14T14:42:19.350Z",
  "environment": "production"
}
```

While you could repeatedly call [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) to check a payment's status instead of listening for webhooks, this approach may [trigger API rate limit errors](/docs/errors/rate-limit-exceeded/) and is strongly discouraged. Only consider polling as a fallback when webhooks are unavailable or significantly delayed.

Once you receive the [`PAYMENT_STATUS_UPDATE` webhook](/docs/api/products/payment-initiation/#payment_status_update), use the [`payment_id`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-payment-id) field to retrieve the payment's metadata from your database.

From the status update object, use the [`new_payment_status`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-new-payment-status) field to decide what action needs to be taken for this payment. For example, you may decide to fund the account once the payment status is `PAYMENT_STATUS_EXECUTED`, or display a user-facing error if the status is `PAYMENT_STATUS_REJECTED`.

Note: In order to protect against webhook forgery attacks, before funding an account in response to a webhook, it's recommended to ensure the status is legitimate. One way is to implement [webhook verification](/docs/api/webhooks/webhook-verification/). An easier alternative is to verify the new status by calling [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) a single time after receiving the webhook.

#### Terminal payment statuses

For many payments, including most payments in the UK, the terminal status (if not using Virtual Accounts) is `PAYMENT_STATUS_EXECUTED`, which indicates that funds have left the payer's account. `PAYMENT_STATUS_INITIATED` is the terminal state for payments at some non-UK banks, as well as smaller UK banks.

Funds typically settle (i.e., arrive in the payee's account) within a few seconds of the payment execution, although settlement of an executed payment is not guaranteed.

If using [Virtual Accounts](/docs/payment-initiation/virtual-accounts/payment-confirmation/), the successful terminal status will always be `PAYMENT_STATUS_SETTLED`, regardless of bank.

To gain access to the `PAYMENT_STATUS_SETTLED` terminal status and track whether a payment has settled, you can use the [Payment Confirmation feature](/docs/payment-initiation/virtual-accounts/payment-confirmation/), available via the [Virtual Accounts](/docs/payment-initiation/virtual-accounts/) product. To request access, contact your Account Manager.

##### Terminal payment status timeframes

The most common payment schemes and their settlement times are:

- **Faster Payments Service (FPS):** An instant payment scheme used in the UK. The money is usually available in the receiving customer's account almost immediately, although it can sometimes take up to two hours.
- **SEPA Instant Credit Transfer:** A pan-European instant payment scheme. Funds will be made available in less than ten seconds.
- **SEPA Credit Transfer:** A pan-European payment scheme where payments are processed and settled within one business day.

#### Payment status transitions

Below is the status transitions for Payment Initiation. If using Virtual Accounts, see [Payment Confirmation](/docs/payment-initiation/virtual-accounts/payment-confirmation/), which adds the `SETTLED` status.

Payment status transitions

```
INPUT_NEEDED 
    |
    +--> AUTHORISING 
    |        |
    |        +--> INITIATED [terminal state for successful payments at some banks]
    |        |        |
    |        |        +--> EXECUTED [For some banks only, mostly in UK]
    |        |        |
    |        |        +--> BLOCKED 
    |        |        +--> REJECTED
    |        |        +--> CANCELLED
    |        |        +--> INSUFFICIENT_FUNDS
    |        |        +--> FAILED
    |        |
    |        +--> BLOCKED
    |
    +--> BLOCKED
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
