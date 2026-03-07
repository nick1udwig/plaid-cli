---
title: "Payments (Europe) - Handling Unexpected Payments | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/virtual-accounts/handling-unexpected-payments/"
scraped_at: "2026-03-07T22:05:14+00:00"
---

# Handling Unexpected Payments

#### Automatically refund unexpected bank transfers to virtual accounts

Plaid automatically detects and refunds unexpected payments sent directly to your virtual accounts via bank transfer. This reduces operational overhead and improves end user experience by handling these transfers without manual intervention.

#### Overview

Unexpected payments are transfers sent directly to your virtual account via bank transfer, not through a Plaid-initiated payment flow. These transfers arrive with a `settled` status but have no associated Plaid payment, making them difficult to reconcile. When automatic refunds are enabled, Plaid detects these unexpected payments and automatically initiates a full refund to the sender. You receive webhook notifications throughout the process, and transactions are linked for easy reconciliation. You can also configure an allowlist to prevent refunds from specific accounts, such as your own funding accounts.

#### How automatic refunds work

When an unexpected payment is received:

1. **Detection**: Plaid identifies the incoming transfer as an unexpected payment and classifies it as a `BANK_TRANSFER` transaction.
2. **Notification**: A [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) webhook is sent for the incoming transaction, including the transaction type. You can retrieve counterparty details by calling [[`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget)](/docs/api/products/virtual-accounts/#wallettransactionget).
3. **Allowlist check**: Plaid checks if the sender is on your [allowlist](/docs/payment-initiation/virtual-accounts/handling-unexpected-payments/#allowlist-configuration). If yes, the transaction is classified as `ACCOUNT_FUNDING` and no refund is initiated.
4. **Automatic refund**: If the sender is not allowlisted, Plaid automatically initiates a full refund using the counterparty details and classifies it as an `AUTO_REFUND` transaction.
5. **Status updates**: [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) webhooks are sent as the refund progresses to `EXECUTED` status.
6. **Transaction linking**: The original transfer and refund are linked together for [reconciliation](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/#reconciling-transactions) purposes.

#### Transaction types

Three transaction types are used for automatic refunds:

| Type | Description |
| --- | --- |
| `BANK_TRANSFER` | Incoming bank transfer not initiated through Plaid. Triggers `AUTO_REFUND` if automatic refunds are enabled. |
| `AUTO_REFUND` | Outgoing refund automatically initiated by Plaid in response to an unexpected `BANK_TRANSFER`. |
| `ACCOUNT_FUNDING` | Incoming transfer from an allowlisted account. Not automatically refunded. |

#### Webhooks

You receive [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) webhooks for both the incoming `BANK_TRANSFER` and as the `AUTO_REFUND` progresses. The refund follows the same status flow as other outgoing payments, reaching `EXECUTED` when complete.

Configure your webhook endpoint in the [Plaid Dashboard](https://dashboard.plaid.com/developers/webhooks) using the `Wallet transaction event` type.

#### Allowlist configuration

You can exclude specific accounts from automatic refunds by adding them to an allowlist. This is useful for accounts you control that fund your virtual account. Transfers from allowlisted accounts are classified as `ACCOUNT_FUNDING` instead of `BANK_TRANSFER` and are not automatically refunded.

Contact your Account Manager to configure the allowlist for your virtual accounts.

#### Handling failures

Plaid automatically retries failed refunds. If a refund continues to fail after multiple attempts, you receive a webhook indicating the failure, and manual intervention is required.

Common failure reasons include:

- Insufficient balance in the virtual account
- Sender's account closed or invalid
- Refund fails sanctions screening

The failed `AUTO_REFUND` remains linked to the original `BANK_TRANSFER`, making it easy to identify which transfer needs manual attention.

#### Testing automatic refunds

You can begin testing automatic refunds for unexpected payments in Sandbox by following the steps listed in the [Add Virtual Accounts to your app](/docs/payment-initiation/virtual-accounts/add-to-app/) guide. For Production access you will first need to [submit a product access request Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Next steps

- Contact your Account Manager to enable automatic refunds for your virtual accounts.
- See [Managing virtual accounts](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/) for transaction reconciliation.
- See [Virtual Accounts API reference](/docs/api/products/virtual-accounts/) for technical details.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
