---
title: "Transfer - Refunds | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/refunds/"
scraped_at: "2026-03-07T22:05:25+00:00"
---

# Refunds

#### Issue a refund for an ACH debit

Refunds allow you to quickly and easily refund customers for ACH debit transfers.

#### Creating a refund

Use [`/transfer/refund/create`](/docs/api/products/transfer/refunds/#transferrefundcreate) or the Plaid Dashboard transfer details pane to create a refund for a debit transfer. Pass the ID of the transfer you'd like to refund as `transfer_id` in the request, and specify the amount of the refund with the `amount` field. The amount of a refund can't exceed the amount of the refunded transfer.

![Plaid Dashboard UI showing refund initiation form with fields for net amount, refund amount, confirmation, and a Confirm button.](/assets/img/docs/transfer/refunds_dashboard.png)

Initiate refunds from the Plaid Dashboard UI.

You can create multiple refunds for a transfer as long as the total amount does not exceed the amount of the original transfer.

Refunds come out of the available balance of the ledger used for the original debit transfer. If there are not enough funds in the available balance to cover the refund amount, the refund will be rejected. You can create a refund at any time. Plaid does not impose any hold time on refunds.

A refund can still be issued even if the Item's `access_token` is no longer valid (e.g. if the user revoked OAuth consent or the Item was deleted via [`/item/remove`](/docs/api/items/#itemremove)), as long as the account and routing number pair used to make the original transaction is still valid. A refund cannot be issued if the Item has an [invalidated TAN](/docs/auth/#tokenized-account-numbers), which can occur with Chase or PNC.

If you issue a refund for a debit transfer within 2 business days of the settlement date of the original debit transfer, the debit could still return for insufficient funds. Debits to consumers can also return as “unauthorized” for up to 60 calendar days.

If a debit transfer that you have refunded is later returned by the payment network, you may be debited for both the return and for the refund. Plaid will not reimburse you for these costs.

Transfers in a `cancelled`, `failed`, or `returned` state cannot be refunded.

#### Canceling a refund

Use [`/transfer/refund/cancel`](/docs/api/products/transfer/refunds/#transferrefundcancel) or the Dashboard to cancel a refund. You can only cancel refunds before they've been submitted to the ACH network for processing.

#### Refund events

Plaid creates an event any time the status of a refund changes. Refund events are prefixed with `refund.` and have a non-null `refund_id` in the event object. See [event monitoring](https://plaid.com/docs/transfer/reconciling-transfers/#event-monitoring) for more information on monitoring events.

Refund events will have many of the same characteristics as the original payment. They will have the same `account_id`, as well as having the same `transfer_type` and a positive transfer amount. For example, if you were to refund $5 from a customer's `debit` transfer of $100.00, that refund will be displayed in the events queue with a `debit` transfer type and a `"5.00"` transfer amount.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
