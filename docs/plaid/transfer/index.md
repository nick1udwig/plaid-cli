---
title: "Transfer - Transfer Overview | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/"
scraped_at: "2026-03-07T22:05:23+00:00"
---

# Introduction to Transfer

#### Intelligently process transfers between accounts

Get started with Transfer

[API Reference](/docs/api/products/transfer/)[Quickstart](https://github.com/plaid/transfer-quickstart)[Demo](https://plaid.coastdemo.com/share/66d783135f0f8b245d4153fe?zoom=100)[Application Process](/docs/transfer/application/)

#### Overview

Plaid Transfer (US only) is a flexible multi-rail payment platform designed for companies looking to add or improve their bank payment solution. Transfer provides all of the necessary tools to easily send and manage ACH, RTP, RfP, wire transfer, and FedNow transactions, including:

- Easy integration with a single API: Avoid the need for multiple service providers. Connect user accounts, make smarter transaction decisions, manage risk, and move money—all through a single Plaid integration.
- Fast settlement, simplified reconciliation: Sweep transaction funds into your treasury account quickly and balance your books with an intuitive reconciliation report.
- Multi-rail routing: Dynamic routing between RTP, RfP, and FedNow. Fall back to same day ACH if needed. Easily enable new payment rails with a single setting, using the same integration code.
- Streamlined operational support: Manage daily operations with dashboards to easily monitor transfer activity.
- Payment risk reduction: Minimize your payment failure and ACH return rates by using Plaid's risk engine.

Plaid Transfer provides a full-service funds transfer solution for payers and recipients within the US. If you prefer to use a third party payment processor, your use case is not supported by Transfer (e.g. a marketplace or money transfer app), or if you or your counterparties are outside the US, see [Plaid Auth](/docs/auth/) instead for a bring-your-own-processor funds transfer solution. See [Auth and Transfer comparison](/docs/payments/#auth-and-transfer-comparison) for a side-by-side product comparison.

For similar end-to-end payment capabilities in Europe, see [Payments (Europe)](/docs/payment-initiation/).

### Integration overview

Prefer to learn by watching? Watch this introduction to Transfer in 3(-ish) minutes!

The process below defines the highest level steps to move money via Transfer and monitor for updates.

Prerequisite: Prior to beginning your integration, [complete the Transfer Application](/docs/transfer/application/) and receive approval. You can begin integrating with Sandbox while waiting for approval.

1. [Initialize](/docs/transfer/creating-transfers/#account-linking) a Link session to link a consumer bank account
2. [Authorize](/docs/transfer/creating-transfers/#authorizing-a-transfer) and [create](/docs/transfer/creating-transfers/#initiating-a-transfer) the transfer
3. (Optional) [Monitor](/docs/transfer/reconciling-transfers/) for updates to the transfer
4. (Optional) [Customize Signal Rules](/docs/transfer/signal-rules/) to adjust authorization thresholds for your use case

Additionally, implement additional features such as [refunds](/docs/transfer/refunds/), [recurring transfers](/docs/transfer/recurring-transfers/), or [Transfer UI](/docs/transfer/using-transfer-ui/) as needed.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
