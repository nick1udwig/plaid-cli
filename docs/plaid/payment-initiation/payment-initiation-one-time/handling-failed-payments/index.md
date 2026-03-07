---
title: "Payments (Europe) - Handling failed payments | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/payment-initiation-one-time/handling-failed-payments/"
scraped_at: "2026-03-07T22:05:11+00:00"
---

# Handling failed payments

#### Implement retry strategies and handle failures gracefully

#### Retrying a payment

To retry a failed payment, call [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) again with the same `recipient_id` and `amount`. Always ensure that retry logic doesn't create duplicate successful payments -- make sure to create a new payment with a new idempotency key to ensure the retry is processed as a separate transaction, and verify the status of previous attempts before initiating new ones.

If you have reached the limit for retry attempts, or the error is not immediately retryable, you should clearly explain the payment failure to your user and offer alternative payment methods or funding options.

#### Retryable payment failures

The following failures are temporary and can be successfully retried.

**`PAYMENT_STATUS_FAILED`** is typically due to a transient issue, such as temporary network problems or bank system maintenance. If you experience this status, it is recommended to make one or two retry attempts immediately. Limit the number of retry attempts to avoid causing a negative customer experience.

**`PAYMENT_STATUS_BLOCKED`** indicates the payment was blocked by Plaid. This is typically due to regulatory or compliance reasons. Once the compliance issue is resolved, the payment can be retried.

**`PAYMENT_STATUS_INSUFFICIENT_FUNDS`** indicates the payment failed due to insufficient funds. Once the user has sufficient funds, the payment can be retried. If you experience this status, it is recommended to tell the user what happened and to give them the option to trigger a retry, or make a new payment, once they have added funds.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
