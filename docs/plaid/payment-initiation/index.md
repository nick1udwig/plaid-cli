---
title: "Payments (Europe) | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/"
scraped_at: "2026-03-07T22:05:10+00:00"
---

# Introduction to Payments (Europe)

#### Initiate real-time payments for account funding and transfers

Get started with Payments (Europe)

[API Reference](/docs/api/products/payment-initiation/)[Integration Guide](/docs/payment-initiation/payment-initiation-one-time/add-to-app/)

#### Overview

Plaid's European Payments suite enables your users to make real-time payments without manually entering their account details or leaving your app. Common use cases include:

- **Account funding:** Allow users to easily transfer money into your app or wallet.
- **Pay-by-bank:** Payments supports both one-time payments and (in the UK) recurring payments. Accept payments, track settlement, and issue refunds.
- **Payouts:** Pay out stored balances to your end users.

![Plaid Link payment flow: view payment details, select bank, authenticate with Gingham Bank, confirm payment with details.](/assets/img/docs/payment_initiation/payment-initiation-uk-link-screens.png)

Benefits of Payments include:

- **Easy integration with a single API:** Verify user accounts, manage risk, and move money across 20 markets, all through a single Plaid integration.
- **Real-time settlement, simplified reconciliation:** Sweep transaction funds into your treasury account quickly and balance your books with an intuitive reconciliation report.
- **Multi-rail support:** Payments supports Faster Payments Service (FPS), SEPA Instant Credit Transfer, SEPA Credit Transfer and local country payment rails. Easily switch payment rails with a single request parameter, no complex code changes required.
- **Streamlined operational support:** Manage daily operations with the Plaid Dashboard to easily monitor payment activity.

For enhanced capabilities, you can add the optional [Virtual Accounts](/docs/payment-initiation/virtual-accounts/) feature to unlock advanced features such as sending refunds to users, allowing your users to payout a credit balance to a bank account, or gaining granular visibility into the settlement status of a payment.

Payments is Europe-only. Looking for similar capabilities in the US? Check out the [Transfer](/docs/transfer/) docs.

#### Payments solutions

The Payments suite includes the following solutions. You can "mix and match" these solutions, incorporating all three for a robust, fully-integrated solution, or purchase them separately, incorporating only the ones you need. (Exception: Payouts requires Virtual Accounts and cannot be used separately.)

**[One-time Payment Initiation](/docs/payment-initiation/payment-initiation-one-time/)**: Enable your users to make single real-time payments, such as SEPA Instant Credit Transfers or Single Immediate Payments (UK), without manually entering their account number and sort code.

**[Variable Recurring Payments (VRP)](/docs/payment-initiation/variable-recurring-payments/)** *(UK only)*: Establish a single recurring payment consent, which can then be used for a series of ongoing payments, with no end user interaction required. For background information on what Variable Recurring Payments (VRPs) are and how they work in the UK banking system, see the [Plaid VRP FAQ](https://support.plaid.com/hc/en-us/articles/24622039958295-What-are-Variable-Recurring-Payments-VRPs). For instructions on integrating support for Variable Recurring Payments using Plaid, see [Add Variable Recurring Payments to your app](/docs/payment-initiation/variable-recurring-payments/add-to-app/).

**[Payouts](/docs/payment-initiation/virtual-accounts/payouts/)** (requires Virtual Accounts add-on): Payouts enables users to seamlessly and instantly make withdrawals or get paid.

Any of the solutions above can also be enhanced with the Virtual Accounts add on:

**[Virtual Accounts](/docs/payment-initiation/virtual-accounts/)**: Collect payments, get confirmation when payments settle, initiate payouts and refunds, and streamline reconciliation processes. Virtual Accounts provides granular control and visibility throughout the entire lifecycle of a payment.

Payments can also be used with other Plaid solutions such as [Auth](/docs/auth/), [Identity](/docs/identity/), [Identity Verification](/docs/identity-verification/), and [Monitor](/docs/monitor/) for enhanced anti-fraud protections and KYC/AML compliance.

#### Onboarding and account funding integration guide

See [User onboarding and account funding](/docs/payment-initiation/payment-initiation-one-time/#onboarding-and-account-funding-integration-guide) for an implementation guide to using Plaid's Europe Payments suite as part of a KYC and AML compliant onboarding flow.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
