---
title: "Payments (Europe) - Variable Recurring Payments | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/variable-recurring-payments/"
scraped_at: "2026-03-07T22:05:12+00:00"
---

# Variable Recurring Payments

#### Use Variable Recurring Payments to power automated bill payments

#### Overview

Variable Recurring Payments (VRP) allows you to obtain bank account details and end-user consent for payments within a set of limits. You can then use VRP to automatically initiate payments from your customer's bank account, without requiring additional end user interaction for each transfer.

VRP is ideal for creating automated recurring bank payments, especially when the payment amount can change from payment to payment, such as for usage-based subscription services, utility payments, or credit card payments.

By combining one-time payments with the optional [Virtual Accounts](/docs/payment-initiation/virtual-accounts/) product, you can enable additional functionality, such as sending refunds to users, allowing your users to payout a credit balance to a bank account, or gaining granular visibility into the settlement status of a payment.

To use Variable Recurring Payments, your end users must bank in the UK (your company can be based elsewhere). To request access, contact your Account Manager. If you are not yet a Plaid customer, [contact sales](https://plaid.com/contact/).

For more information on Variable Recurring Payments, see the [VRP FAQ](https://support.plaid.com/hc/en-us/articles/24622039958295-What-are-Variable-Recurring-Payments-VRPs).

![Plaid Link VRP setup showing bank selection, authentication, and success confirmation for recurring payments.](/assets/img/docs/payment_initiation/vrp.png)

#### Implementation Process

See the [Add Variable Recurring Payments to your App](/docs/payment-initiation/variable-recurring-payments/add-to-app/) guide to learn more about the product and how to implement it using Plaid Link.

#### Testing Variable Recurring Payments

Variable Recurring Payments can immediately be tried out with test data in the Sandbox environment. In order to test payments against live Items in Production, you will need to first request access by [contacting Sales](https://plaid.com/contact) or your Account Manager.

When testing in [Limited Production](/docs/sandbox/#testing-with-live-data-using-limited-production), payments must be below 5 GBP or other chosen [currency](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-amount-currency). For details on any payment limits in full Production, contact your Plaid Account Manager.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
