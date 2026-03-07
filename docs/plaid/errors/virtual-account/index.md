---
title: "Errors - Virtual Accounts errors (Europe) | Plaid Docs"
source_url: "https://plaid.com/docs/errors/virtual-account/"
scraped_at: "2026-03-07T22:04:54+00:00"
---

# Virtual Account(UK/EU) Errors

#### **TRANSACTION\_INSUFFICIENT\_FUNDS**

##### Common Causes

- The account does not have sufficient funds to complete the transaction.

##### Troubleshooting steps

Check the [available balance](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-balance-available) on the Virtual Account.

If the Virtual Account has insufficient balance, [fund the Virtual Account](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/#fund-your-virtual-account) and try again.

Try again with a lower [transaction amount](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-amount).

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "TRANSACTION_INSUFFICIENT_FUNDS",
 "error_message": "There are insufficient funds to complete the transaction",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTION\_AMOUNT\_EXCEEDED**

##### Common Causes

- The transaction amount exceeds the allowed threshold configured for this client.

##### Troubleshooting steps

Contact your Plaid Account Manager for your compliance rules.

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "TRANSACTION_AMOUNT_EXCEEDED",
 "error_message": "Transaction amount exceeds the allowed threshold for this client",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTION\_ON\_SAME\_ACCOUNT**

##### Common Causes

- The recipient bank account details on the transaction are incorrect and refer to the source Virtual Account.

##### Troubleshooting steps

Confirm that the [bank account details](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty) do not route back to the Virtual Account.

Try again with a different recipient bank account.

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "TRANSACTION_ON_SAME_ACCOUNT",
 "error_message": "Payment to the same account is not allowed",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTION\_CURRENCY\_MISMATCH**

##### Common Causes

- The currency on the recipient bank account is different than that of the wallet.

##### Troubleshooting steps

Check the [currency](/docs/api/products/virtual-accounts/#wallet-get-response-balance-iso-currency-code) of the Virtual Account

Confirm that the provided [bank account details](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty) accepts the same currency as the Virtual Account.

Try again with a different recipient bank account.

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "TRANSACTION_CURRENCY_MISMATCH",
 "error_message": "The currency between the wallet and recipient account is different",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTION\_IBAN\_INVALID**

##### Common Causes

- The provided IBAN on the recipient bank account is invalid.

##### Troubleshooting steps

Check that a valid [IBAN](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers-international-iban) was provided in the request.

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "TRANSACTION_IBAN_INVALID",
 "error_message": "The provided IBAN is invalid",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTION\_BACS\_INVALID**

##### Common Causes

- The provided Account Number and/or Sort Code on the recipient account is invalid.

##### Troubleshooting steps

Check that a valid [Account Number](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers-bacs-account) was provided in the request.

Check that a valid [Sort Code](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers-bacs-sort-code) was provided in the request.

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "TRANSACTION_BACS_INVALID",
 "error_message": "The provided BACS account number and/or sort code is invalid",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTION\_FAST\_PAY\_DISABLED**

##### Common Causes

- The provided Sort Code on the recipient account is not enabled for Faster Payments. GBP transactions out of Virtual Accounts in GB are made via the Faster Payments rail.

##### Troubleshooting steps

Check with the bank of the recipient bank account that Faster Payments has been enabled.

Try again using a different recipient bank account if Faster Payments is not supported.

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "TRANSACTION_FAST_PAY_DISABLED",
 "error_message": "The recipient sort code is not enabled for faster payments",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSACTION\_EXECUTION\_FAILED**

##### Common Causes

- The transaction failed to execute due to an internal error.
- The transaction might have been declined by the receiving bank.
- Technical issues with the payment processor.

##### Troubleshooting steps

Check the request body against the API reference for [`/wallet/transaction/execute`](/docs/api/products/virtual-accounts/#wallettransactionexecute)

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "TRANSACTION_EXECUTION_FAILED",
 "error_message": "There was a problem executing the transaction. Please retry.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NONIDENTICAL\_REQUEST**

##### Common Causes

- The request parameters have changed compared to the original request with the same idempotency key.

##### Troubleshooting steps

Check that [idempotency keys](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/#idempotency) are not being reused for different requests.

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "NONIDENTICAL_REQUEST",
 "error_message": "Request body does not match previous request with this idempotency key",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **REQUEST\_CONFLICT**

##### Common Causes

- The original request is still processing and has not completed.
- A network or system issue caused a delay in processing the original request, resulting in a conflict when the new request is sent.

##### Troubleshooting steps

Wait and retry the request again

API error response

```
http code 400
{
 "error_type": "TRANSACTION_ERROR",
 "error_code": "REQUEST_CONFLICT",
 "error_message": "Original request is still processing",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
