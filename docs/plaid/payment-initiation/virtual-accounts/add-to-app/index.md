---
title: "Payments (Europe) - Add Virtual Accounts to your app | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/virtual-accounts/add-to-app/"
scraped_at: "2026-03-07T22:05:13+00:00"
---

# Add Virtual Accounts to your app

#### Use virtual accounts to gain additional insight and control over the flow of funds

In this guide, we'll walk through how to set up and manage virtual accounts. After integrating you will be able to send and receive money, make withdrawals and refunds, and reconcile funds for faster reconciliation.

#### Integrate with Plaid Payment Initiation and enable Virtual Accounts

Virtual accounts are designed to be used in conjunction with the Payment Initiation and Variable Recurring Payments APIs. To get the most out of Plaid virtual accounts, first [add Payment Initiation to your app](/docs/payment-initiation/payment-initiation-one-time/add-to-app/).
In order to get started, you must first have a wallet set up. Virtual Accounts are not enabled by default in Sandbox. To get access in Sandbox or Production, [contact Sales](https://plaid.com/contact) or your Account Manager. Once enabled, create a virtual account via [`/wallet/create`](/docs/api/products/virtual-accounts/#walletcreate).

/wallet/create

```
const request: WalletCreateRequest = {
  iso_currency_code: isoCurrencyCode,
};
try {
  const response = await plaidClient.walletCreate(request);
  const walletID = response.data.wallet_id;
  const balance = response.data.balance;
  const numbers = response.data.numbers;
  const recipientID = response.data.recipient_id;
} catch (error) {
  // handle error
}
```

#### Enable Virtual Account Transaction webhooks

Enable [virtual account webhooks](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/#webhooks) to receive notifications when new transaction events occur. Plaid will send a webhook for each transaction event. These notifications can be used to stay updated on transaction status information.

#### Funding your Virtual Account through Payment Initiation

##### Create a Payment that will be initiated to your virtual account

/payment\_initiation/payment/create

```
const request: PaymentInitiationPaymentCreateRequest = {
  recipient_id: recipientID,
  reference: 'TestPayment',
  amount: {
    currency: 'GBP',
    value: 100.0,
  },
};
try {
  const response = await plaidClient.paymentInitiationPaymentCreate(request);
  const paymentID = response.data.payment_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

Use the `recipient_id` associated with your `wallet_id` so Plaid knows to route the payment to your virtual account.

##### Launch and complete the Payment Initiation flow in Link

For more details on how to do this, see the [Payment Initiation documentation](/docs/payment-initiation/payment-initiation-one-time/add-to-app/#launch-the-payment-initiation-flow-in-link).

##### Confirm that the payment has settled via a payment webhook

For more details, see the [`PAYMENT_STATUS_UPDATE`](/docs/api/products/payment-initiation/#payment_status_update) webhook.

PAYMENT\_STATUS\_UPDATE

```
{
  "webhook_type": "PAYMENT_INITIATION",
  "webhook_code": "PAYMENT_STATUS_UPDATE",
  "payment_id": "<PAYMENT_ID>",
  "new_payment_status": "PAYMENT_STATUS_SETTLED",
  "old_payment_status": "PAYMENT_STATUS_EXECUTED",
  "original_reference": "Account Funding 99744",
  "adjusted_reference": "Account Funding 99",
  "original_start_date": "2017-09-14",
  "adjusted_start_date": "2017-09-15",
  "timestamp": "2017-09-14T14:42:19.350Z"
}
```

##### Confirm that funds have arrived in your virtual account via a transaction webhook

For more details, see the [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) webhook.

WALLET\_TRANSACTION\_STATUS\_UPDATE

```
{
  "webhook_type": "WALLET",
  "webhook_code": "WALLET_TRANSACTION_STATUS_UPDATE",
  "wallet_id": "<WALLET_ID>",
  "new_status": "SETTLED",
  "old_status": "INITIATED",
  "timestamp": "2021-09-14T14:42:19.350Z"
}
```

#### Executing a Payout

##### Fetch your virtual account to confirm it has sufficient balance

/wallet/get

```
const request: WalletGetRequest = {
  wallet_id: walletID,
};
try {
  const response = await plaidClient.walletGet(request);
  const walletID = response.data.wallet_id;
  const balance = response.data.balance;
  const numbers = response.data.numbers;
} catch (error) {
  // handle error
}
```

##### Execute a payout

/wallet/transaction/execute

```
const request: WalletTransactionExecuteRequest = {
  wallet_id: walletID,
  counterparty: {
    name: 'Test',
    numbers: {
      bacs: {
        account: '12345678',
        sort_code: '123456',
      },
    },
  },
  amount: {
    value: 1,
    iso_currency_code: 'GBP',
  },
  reference: 'transaction ABC123',
  idempotency_key: '39fae5f2-b2b4-48b6-a363-5328995b2753',
};
try {
  const response = await plaidClient.walletTransactionExecute(request);
  const transactionID = response.data.transaction_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

##### Confirm the payout was executed, via a transaction webhook

For more details, see the [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) webhook.

WALLET\_TRANSACTION\_STATUS\_UPDATE

```
{
  "webhook_type": "WALLET",
  "webhook_code": "WALLET_TRANSACTION_STATUS_UPDATE",
  "wallet_id": "<WALLET_ID>",
  "new_status": "EXECUTED",
  "old_status": "INITIATED",
  "timestamp": "2021-09-14T14:42:19.350Z"
}
```

#### Refunding a Payment Initiation payment

##### Fetch your virtual account to confirm it has sufficient balance

/wallet/get

```
const request: WalletGetRequest = {
  wallet_id: walletID,
};
try {
  const response = await plaidClient.walletGet(request);
  const walletID = response.data.wallet_id;
  const balance = response.data.balance;
  const numbers = response.data.numbers;
} catch (error) {
  // handle error
}
```

##### Refund a Payment Initiation payment

/payment\_initiation/payment/reverse

```
const request: PaymentInitiationPaymentReverseRequest = {
  payment_id: paymentID,
  reference: 'Refund for purchase ABC123',
  idempotency_key: 'ae009325-df8d-4f52-b1e0-53ff26c23912',
};
try {
  const response = await plaidClient.paymentInitiationPaymentReverse(request);
  const refundID = response.data.refund_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

##### Confirm the refund was executed, via a transaction webhook

For more details, see the [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) webhook.

WALLET\_TRANSACTION\_STATUS\_UPDATE

```
{
  "webhook_type": "WALLET",
  "webhook_code": "WALLET_TRANSACTION_STATUS_UPDATE",
  "wallet_id": "<WALLET_ID>",
  "new_status": "EXECUTED",
  "old_status": "INITIATED",
  "timestamp": "2021-09-14T14:42:19.350Z"
}
```

#### Next steps

If you're ready to launch to Production, see the Launch checklist.

[#### Launch checklist

Recommended steps to take before launching in Production

Launch](/docs/launch-checklist/)

#### Launch checklist

Recommended steps to take before launching in Production

[Launch](/docs/launch-checklist/)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
