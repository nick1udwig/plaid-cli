---
title: "Payments (Europe) - Refunds | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/virtual-accounts/refunds/"
scraped_at: "2026-03-07T22:05:15+00:00"
---

# Refunds

#### Refund a Payment Initiation payment

Refunds enable you to refund your end users. You can issue refunds for any [Payment Initiation](/docs/payment-initiation/) payments that have settled in your virtual account.

- The original payment must be in a settled state to be refunded.
- To refund partially, specify the amount as part of the request.
- If the amount is not specified, the refund amount will be equal to all of the remaining payment amount that has not been refunded yet.
- If the remaining payment amount is less than one unit of currency (e.g. 1 GBP or 1 EUR), the refund will fail.

#### Execute a Refund

Make sure your virtual account is set up before following these steps. For more information on setting up an account, see [Managing virtual accounts](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/).

1. Follow the [Payment Confirmation](/docs/payment-initiation/virtual-accounts/payment-confirmation/#confirming-funds-have-settled) flow to pay into your virtual account.
2. Call [`/wallet/get`](/docs/api/products/virtual-accounts/#walletget) to check your virtual account balance for sufficient funds.

   - If you have insufficient funds, [fund your virtual account](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/#fund-your-virtual-account) before proceeding. After funding your virtual account, you can check the updated balance by calling [`/wallet/get`](/docs/api/products/virtual-accounts/#walletget).
3. Issue a refund by calling [`/payment_initiation/payment/reverse`](/docs/api/products/payment-initiation/#payment_initiationpaymentreverse), providing the `payment_id` from the payment made in Step 3, and optionally an `amount` for partial refunds. Store the `refund_id` and `status` from the response.

   - If you have [configured transaction webhooks](https://dashboard.plaid.com/developers/webhooks), you will receive a status update that the Refund transitioned from `INITIATED` to `EXECUTED`.
   - Alternatively, if not using webhooks, you can confirm the transaction has been executed by calling [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) with the `transaction_id`.
   - You can also confirm that the original payment was refunded by calling [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) with the `payment_id` to see the `refund_id` as part of the payment’s details.
4. If you have previously executed any partial refunds for the payment, you can still issue subsequent refunds if this payment has sufficient remaining amount.

   - To check if the payment has sufficient remaining amount, call [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) and fetch the `amount_refunded` field, which represents the amount that has been refunded already. Subtract this from the payment amount to calculate the amount still available to refund.

A successful refund will transition to the `EXECUTED` status. If a refund is in a `FAILED` status, you may try again, with a new `idempotency_key`, until you have a refund in an `EXECUTED` status for the payment. Refunds should transition to `EXECUTED` or `FAILED` automatically. If this does not occur,

Retry the request with the same `idempotency_key`. This will address any network connection issues.

If it consistently takes longer than expected for the underlying payment rails (within seconds to several days), contact your Account Manager.

#### Testing Refunds

You can begin testing Refunds in Sandbox by following the steps listed in the [Add Virtual Accounts to your App](/docs/payment-initiation/virtual-accounts/add-to-app/) guide. For Production access you will first need to [submit a product access request Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
