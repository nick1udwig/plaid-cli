---
title: "Payments (Europe) - Refunds | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/payment-initiation-one-time/refunds/"
scraped_at: "2026-03-07T22:05:11+00:00"
---

# Refunds

#### Handle refunds for Payment Initiation transactions

#### Refunding payments

For Payment Initiation (without Virtual Accounts), refunds can be executed using your banking partner to send funds to a user's bank account.

When calling [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) you can set `request_refund_details: true`. If you do, the refund details (account number and sort code) will be available in the `/payment/get` response. You can then use your banking partner to push a refund to the end-user requesting a refund. Support varies across financial institutions and will not always be available.

Alternatively, if you have used Plaid's [Auth](/docs/auth/) product for account verification before making a payment, you will also have access to the user's bank account details.

For streamlined refunds via Plaid, see [Virtual Accounts](/docs/payment-initiation/virtual-accounts/).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
