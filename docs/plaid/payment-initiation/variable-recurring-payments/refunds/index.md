---
title: "Payments (Europe) - Refunds | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/variable-recurring-payments/refunds/"
scraped_at: "2026-03-07T22:05:12+00:00"
---

# Refunds

#### Process refunds for Variable Recurring Payments

For Variable Recurring Payments (without Virtual Accounts), refunds can be executed using your banking partner to send funds to a user's bank account.

When calling [`/payment_initiation/consent/get`](/docs/api/products/payment-initiation/#payment_initiationconsentget) you can retrieve the payer details, such as name, and account numbers in the response. You can then use your banking partner to push a refund to the end user requesting a refund.

For streamlined refunds via Plaid, see [Virtual Accounts](/docs/payment-initiation/virtual-accounts/).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
