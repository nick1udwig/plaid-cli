---
title: "Payments (Europe) - Payment Confirmation | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/virtual-accounts/payment-confirmation/"
scraped_at: "2026-03-07T22:05:14+00:00"
---

# Payment Confirmation

#### Confirm when funds have settled in a virtual account

Plaid offers Payment Confirmation on top of our [Payment Initiation](/docs/payment-initiation/) APIs. Payment Confirmation provides confirmation of payments settling in your virtual account. Plaid will notify you when funds have settled so you can confidently release those funds to your end users. Each virtual account can receive or send funds on payment rails supported by the currency of the account.

The Payment Confirmation flow builds on the Payment Initiation flow, with two main differences:

- Payments are initiated with the client-owned virtual account as the recipient.
- Optionally, the current balance of the virtual account can be automatically swept (transferred) periodically to your bank account. To use this functionality, reach out to your Account Manager.

To your end user, there is no difference between initiating a payment to a virtual account and initiating a payment directly to the recipient bank account.

#### Confirming funds have settled

Make sure your virtual account is set up before following these steps. For more information on setting up an account, see [Managing virtual accounts](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/).

1. Go through the [Payment Initiation flow](/docs/payment-initiation/payment-initiation-one-time/) and initiate a payment.

   - Use the `recipient_id` associated with your `wallet_id` so Plaid knows to route the payment to your virtual account.
2. The payment will settle in your virtual account. You will receive the following webhooks:

   - If you included a webhook URL when creating your `link_token` you will receive a [payment status update](/docs/api/products/payment-initiation/#payment_status_update) that the payment is now `SETTLED`. This is the primary webhook that indicates payment confirmation.
   - If you configured a webhook through the [Plaid Dashboard](https://dashboard.plaid.com/developers/webhooks), you will also receive a [status update](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) that a transaction of type `PIS_PAY_IN` transitioned from `INITIATED` to `SETTLED` in your virtual account. This webhook indicates an update to your virtual account with the funds that have settled (i.e., the transaction that corresponds to the payment that settled).
   - In addition to webhooks, you can check if the payment has settled by calling [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) with the `payment_id` from Step 1.
3. Once the payment has settled, you can safely release funds to your end user.

In the UK, payments settle nearly instantly and run on the Faster Payment Services rails. In the Eurozone, Plaid defaults to SEPA Credit Transfer, which settles within one business day (approximately 12 hours later, if both the payment initiation date and the next day are business days). If you want to use SEPA Instant Credit Transfer for faster settlement times (within seconds, 24/7), you can specify the payment scheme when creating the payment.

#### Testing Payment Confirmation

You can begin testing Payment Confirmation in Sandbox by following the steps listed in the [Add Virtual Accounts to your App](/docs/payment-initiation/virtual-accounts/add-to-app/) guide. For Production access you will first need to [submit a product access request Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
