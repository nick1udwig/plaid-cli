---
title: "Payments (Europe) - Payment Initiation | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/payment-initiation-one-time/"
scraped_at: "2026-03-07T22:05:10+00:00"
---

# Payment Initiation

#### Use Payment Initiation for one-time payments

Get started with Payment Initiation

[API Reference](/docs/api/products/payment-initiation/)[Integration Guide](/docs/payment-initiation/payment-initiation-one-time/add-to-app/)[Onboarding and account funding guide](/docs/payment-initiation/payment-initiation-one-time/user-onboarding-and-account-funding/)

#### Payment Initiation

Plaid's one-time Payment Initiation functionality enables your users to make secure, real-time payments without manually entering their account details or leaving your app. Perfect for account funding, marketplace transactions, or any single payment needs, one-time payments, also known in the UK as Single Immediate Payments, provide a faster, easier, and safer way to move money.

Benefits of Payment Initiation:

- **Faster, easier, safer:** Real-time payments that boost revenue while reducing fraud and chargebacks with identity and bank account verification.
- **Infrastructure built for scale:** Plaid handles the backend heavy lifting with instant notifications, seamless reconciliation, and real-time settlement across 20 markets.
- **Multi-rail support:** Underpinned by Faster Payments Service (FPS), SEPA Instant Credit Transfer, SEPA Credit Transfer and local country payment rails. Easily decide the payment rails with a single setting, using the same integration code.

By combining one-time payments with the optional [Virtual Accounts](/docs/payment-initiation/virtual-accounts/) product, you can enable even more functionality, such as sending refunds to users, allowing your users to payout a credit balance to a bank account, or gaining granular visibility into the settlement status of a payment.

#### Integration details

See the [Add Payment Initiation to your App](/docs/payment-initiation/payment-initiation-one-time/add-to-app/) guide to learn more about the product and how to implement it using Plaid Link.

For a complete integration, you will need to [track payment status](/docs/payment-initiation/payment-status/) and [handle failed payments](/docs/payment-initiation/payment-initiation-one-time/handling-failed-payments/).

If you plan to implement refund capabilities, see [refunds](/docs/payment-initiation/payment-initiation-one-time/refunds/).

#### Sample app

For a simple real-world implementation of Payment Initiation, see the [Payment Initiation Pattern App](https://github.com/plaid/payment-initiation-pattern-app) on GitHub, which uses Payment Initiation in an account funding use case.

#### Onboarding and account funding integration guide

See [User onboarding and account funding](/docs/payment-initiation/payment-initiation-one-time/user-onboarding-and-account-funding/) for an implementation guide to using Payment Initiation as part of a KYC and AML compliant onboarding flow.

#### Testing Payment Initiation

Payments can immediately be tried out with test data in the Sandbox environment. In order to test payments against live Items in Production, you will need to first request access by [contacting Sales](https://plaid.com/contact) or your Account Manager.

When testing in [Limited Production](/docs/sandbox/#testing-with-live-data-using-limited-production), payments must be below 5 GBP or other chosen [currency](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-amount-currency). For details on any payment limits in full Production, contact your Plaid Account Manager.

#### Pricing

Payment Initiation is billed on a [per-payment model](/docs/account/billing/#payment-initiation-fee-model). To view the exact pricing you may be eligible for, [contact Sales](https://plaid.com/contact/).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
