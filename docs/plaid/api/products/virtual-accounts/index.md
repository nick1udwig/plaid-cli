---
title: "API - Virtual Accounts | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/virtual-accounts/"
scraped_at: "2026-03-07T22:04:26+00:00"
---

# Virtual Accounts (UK and Europe)

#### API reference for Virtual Accounts endpoints and webhooks

Manage the entire lifecycle of a payment. For how-to guidance, see the [Virtual Accounts documentation](/docs/payment-initiation/virtual-accounts/).

| Endpoints |  |
| --- | --- |
| [`/wallet/create`](/docs/api/products/virtual-accounts/#walletcreate) | Create a virtual account |
| [`/wallet/get`](/docs/api/products/virtual-accounts/#walletget) | Fetch a virtual account |
| [`/wallet/list`](/docs/api/products/virtual-accounts/#walletlist) | List all virtual accounts |
| [`/wallet/transaction/execute`](/docs/api/products/virtual-accounts/#wallettransactionexecute) | Execute a transaction |
| [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) | Fetch a transaction |
| [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist) | List all transactions |

| See also |  |
| --- | --- |
| [`/payment_initiation/payment/reverse`](/docs/api/products/payment-initiation/#payment_initiationpaymentreverse) | Refund a payment from a virtual account |

| Webhooks |  |
| --- | --- |
| [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) | The status of a transaction has changed |

### Endpoints

=\*=\*=\*=

#### `/wallet/create`

#### Create an e-wallet

Create an e-wallet. The response is the newly created e-wallet object.

/wallet/create

**Request fields**

[`client_id`](/docs/api/products/virtual-accounts/#wallet-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/virtual-accounts/#wallet-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`iso_currency_code`](/docs/api/products/virtual-accounts/#wallet-create-request-iso-currency-code)

requiredstringrequired, string

An ISO-4217 currency code, used with e-wallets and transactions.  
  

Possible values: `GBP`, `EUR`

Min length: `3`

Max length: `3`

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

/wallet/create

**Response fields**

[`wallet_id`](/docs/api/products/virtual-accounts/#wallet-create-response-wallet-id)

stringstring

A unique ID identifying the e-wallet

[`balance`](/docs/api/products/virtual-accounts/#wallet-create-response-balance)

objectobject

An object representing the e-wallet balance

[`iso_currency_code`](/docs/api/products/virtual-accounts/#wallet-create-response-balance-iso-currency-code)

stringstring

The ISO-4217 currency code of the balance

[`current`](/docs/api/products/virtual-accounts/#wallet-create-response-balance-current)

numbernumber

The total amount of funds in the account  
  

Format: `double`

[`available`](/docs/api/products/virtual-accounts/#wallet-create-response-balance-available)

numbernumber

The total amount of funds in the account after subtracting pending debit transaction amounts  
  

Format: `double`

[`numbers`](/docs/api/products/virtual-accounts/#wallet-create-response-numbers)

objectobject

An object representing the e-wallet account numbers

[`bacs`](/docs/api/products/virtual-accounts/#wallet-create-response-numbers-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if you need to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/virtual-accounts/#wallet-create-response-numbers-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/virtual-accounts/#wallet-create-response-numbers-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`international`](/docs/api/products/virtual-accounts/#wallet-create-response-numbers-international)

nullableobjectnullable, object

Account numbers using the International Bank Account Number and BIC/SWIFT code format.

[`iban`](/docs/api/products/virtual-accounts/#wallet-create-response-numbers-international-iban)

stringstring

International Bank Account Number (IBAN).  
  

Min length: `15`

Max length: `34`

[`bic`](/docs/api/products/virtual-accounts/#wallet-create-response-numbers-international-bic)

stringstring

The Business Identifier Code, also known as SWIFT code, for this bank account.  
  

Min length: `8`

Max length: `11`

[`recipient_id`](/docs/api/products/virtual-accounts/#wallet-create-response-recipient-id)

stringstring

The ID of the recipient that corresponds to the e-wallet account numbers

[`status`](/docs/api/products/virtual-accounts/#wallet-create-response-status)

stringstring

The status of the wallet.  
`UNKNOWN`: The wallet status is unknown.  
`ACTIVE`: The wallet is active and ready to send money to and receive money from.  
`CLOSED`: The wallet is closed. Any transactions made to or from this wallet will error.  
  

Possible values: `UNKNOWN`, `ACTIVE`, `CLOSED`

[`request_id`](/docs/api/products/virtual-accounts/#wallet-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "wallet_id": "wallet-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a88",
  "recipient_id": "recipient-id-production-9b6b4679-914b-445b-9450-efbdb80296f6",
  "balance": {
    "iso_currency_code": "GBP",
    "current": 123.12,
    "available": 100.96
  },
  "request_id": "4zlKapIkTm8p5KM",
  "numbers": {
    "bacs": {
      "account": "12345678",
      "sort_code": "123456"
    }
  },
  "status": "ACTIVE"
}
```

=\*=\*=\*=

#### `/wallet/get`

#### Fetch an e-wallet

Fetch an e-wallet. The response includes the current balance.

/wallet/get

**Request fields**

[`client_id`](/docs/api/products/virtual-accounts/#wallet-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/virtual-accounts/#wallet-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`wallet_id`](/docs/api/products/virtual-accounts/#wallet-get-request-wallet-id)

requiredstringrequired, string

The ID of the e-wallet  
  

Min length: `1`

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
  const recipientID = response.data.recipient_id;
} catch (error) {
  // handle error
}
```

/wallet/get

**Response fields**

[`wallet_id`](/docs/api/products/virtual-accounts/#wallet-get-response-wallet-id)

stringstring

A unique ID identifying the e-wallet

[`balance`](/docs/api/products/virtual-accounts/#wallet-get-response-balance)

objectobject

An object representing the e-wallet balance

[`iso_currency_code`](/docs/api/products/virtual-accounts/#wallet-get-response-balance-iso-currency-code)

stringstring

The ISO-4217 currency code of the balance

[`current`](/docs/api/products/virtual-accounts/#wallet-get-response-balance-current)

numbernumber

The total amount of funds in the account  
  

Format: `double`

[`available`](/docs/api/products/virtual-accounts/#wallet-get-response-balance-available)

numbernumber

The total amount of funds in the account after subtracting pending debit transaction amounts  
  

Format: `double`

[`numbers`](/docs/api/products/virtual-accounts/#wallet-get-response-numbers)

objectobject

An object representing the e-wallet account numbers

[`bacs`](/docs/api/products/virtual-accounts/#wallet-get-response-numbers-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if you need to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/virtual-accounts/#wallet-get-response-numbers-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/virtual-accounts/#wallet-get-response-numbers-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`international`](/docs/api/products/virtual-accounts/#wallet-get-response-numbers-international)

nullableobjectnullable, object

Account numbers using the International Bank Account Number and BIC/SWIFT code format.

[`iban`](/docs/api/products/virtual-accounts/#wallet-get-response-numbers-international-iban)

stringstring

International Bank Account Number (IBAN).  
  

Min length: `15`

Max length: `34`

[`bic`](/docs/api/products/virtual-accounts/#wallet-get-response-numbers-international-bic)

stringstring

The Business Identifier Code, also known as SWIFT code, for this bank account.  
  

Min length: `8`

Max length: `11`

[`recipient_id`](/docs/api/products/virtual-accounts/#wallet-get-response-recipient-id)

stringstring

The ID of the recipient that corresponds to the e-wallet account numbers

[`status`](/docs/api/products/virtual-accounts/#wallet-get-response-status)

stringstring

The status of the wallet.  
`UNKNOWN`: The wallet status is unknown.  
`ACTIVE`: The wallet is active and ready to send money to and receive money from.  
`CLOSED`: The wallet is closed. Any transactions made to or from this wallet will error.  
  

Possible values: `UNKNOWN`, `ACTIVE`, `CLOSED`

[`request_id`](/docs/api/products/virtual-accounts/#wallet-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "wallet_id": "wallet-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a88",
  "recipient_id": "recipient-id-production-9b6b4679-914b-445b-9450-efbdb80296f6",
  "balance": {
    "iso_currency_code": "GBP",
    "current": 123.12,
    "available": 100.96
  },
  "request_id": "4zlKapIkTm8p5KM",
  "numbers": {
    "bacs": {
      "account": "12345678",
      "sort_code": "123456"
    },
    "international": {
      "iban": "GB33BUKB20201555555555",
      "bic": "BUKBGB22"
    }
  },
  "status": "ACTIVE"
}
```

=\*=\*=\*=

#### `/wallet/list`

#### Fetch a list of e-wallets

This endpoint lists all e-wallets in descending order of creation.

/wallet/list

**Request fields**

[`client_id`](/docs/api/products/virtual-accounts/#wallet-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/virtual-accounts/#wallet-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`iso_currency_code`](/docs/api/products/virtual-accounts/#wallet-list-request-iso-currency-code)

stringstring

An ISO-4217 currency code, used with e-wallets and transactions.  
  

Possible values: `GBP`, `EUR`

Min length: `3`

Max length: `3`

[`cursor`](/docs/api/products/virtual-accounts/#wallet-list-request-cursor)

stringstring

A base64 value representing the latest e-wallet that has already been requested. Set this to `next_cursor` received from the previous `/wallet/list` request. If provided, the response will only contain e-wallets created before that e-wallet. If omitted, the response will contain e-wallets starting from the most recent, and in descending order.  
  

Max length: `1024`

[`count`](/docs/api/products/virtual-accounts/#wallet-list-request-count)

integerinteger

The number of e-wallets to fetch  
  

Minimum: `1`

Maximum: `20`

Default: `10`

/wallet/list

```
const request: WalletListRequest = {
  iso_currency_code: 'GBP',
  count: 10,
};
try {
  const response = await plaidClient.walletList(request);
  const wallets = response.data.wallets;
  const nextCursor = response.data.next_cursor;
} catch (error) {
  // handle error
}
```

/wallet/list

**Response fields**

[`wallets`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets)

[object][object]

An array of e-wallets

[`wallet_id`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-wallet-id)

stringstring

A unique ID identifying the e-wallet

[`balance`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-balance)

objectobject

An object representing the e-wallet balance

[`iso_currency_code`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-balance-iso-currency-code)

stringstring

The ISO-4217 currency code of the balance

[`current`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-balance-current)

numbernumber

The total amount of funds in the account  
  

Format: `double`

[`available`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-balance-available)

numbernumber

The total amount of funds in the account after subtracting pending debit transaction amounts  
  

Format: `double`

[`numbers`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-numbers)

objectobject

An object representing the e-wallet account numbers

[`bacs`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-numbers-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if you need to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-numbers-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-numbers-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`international`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-numbers-international)

nullableobjectnullable, object

Account numbers using the International Bank Account Number and BIC/SWIFT code format.

[`iban`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-numbers-international-iban)

stringstring

International Bank Account Number (IBAN).  
  

Min length: `15`

Max length: `34`

[`bic`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-numbers-international-bic)

stringstring

The Business Identifier Code, also known as SWIFT code, for this bank account.  
  

Min length: `8`

Max length: `11`

[`recipient_id`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-recipient-id)

stringstring

The ID of the recipient that corresponds to the e-wallet account numbers

[`status`](/docs/api/products/virtual-accounts/#wallet-list-response-wallets-status)

stringstring

The status of the wallet.  
`UNKNOWN`: The wallet status is unknown.  
`ACTIVE`: The wallet is active and ready to send money to and receive money from.  
`CLOSED`: The wallet is closed. Any transactions made to or from this wallet will error.  
  

Possible values: `UNKNOWN`, `ACTIVE`, `CLOSED`

[`next_cursor`](/docs/api/products/virtual-accounts/#wallet-list-response-next-cursor)

stringstring

Cursor used for fetching e-wallets created before the latest e-wallet provided in this response

[`request_id`](/docs/api/products/virtual-accounts/#wallet-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "wallets": [
    {
      "wallet_id": "wallet-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a88",
      "recipient_id": "recipient-id-production-9b6b4679-914b-445b-9450-efbdb80296f6",
      "balance": {
        "iso_currency_code": "GBP",
        "current": 123.12,
        "available": 100.96
      },
      "numbers": {
        "bacs": {
          "account": "12345678",
          "sort_code": "123456"
        }
      },
      "status": "ACTIVE"
    },
    {
      "wallet_id": "wallet-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a999",
      "recipient_id": "recipient-id-production-9b6b4679-914b-445b-9450-efbdb80296f7",
      "balance": {
        "iso_currency_code": "EUR",
        "current": 456.78,
        "available": 100.96
      },
      "numbers": {
        "international": {
          "iban": "GB22HBUK40221241555626",
          "bic": "HBUKGB4B"
        }
      },
      "status": "ACTIVE"
    }
  ],
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/wallet/transaction/execute`

#### Execute a transaction using an e-wallet

Execute a transaction using the specified e-wallet.
Specify the e-wallet to debit from, the counterparty to credit to, the idempotency key to prevent duplicate transactions, the amount and reference for the transaction.
Transactions will settle in seconds to several days, depending on the underlying payment rail.

/wallet/transaction/execute

**Request fields**

[`client_id`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`idempotency_key`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-idempotency-key)

requiredstringrequired, string

A random key provided by the client, per unique wallet transaction. Maximum of 128 characters.  
The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. If a request to execute a wallet transaction fails due to a network connection error, then after a minimum delay of one minute, you can retry the request with the same idempotency key to guarantee that only a single wallet transaction is created. If the request was successfully processed, it will prevent any transaction that uses the same idempotency key, and was received within 24 hours of the first request, from being processed.  
  

Max length: `128`

Min length: `1`

[`wallet_id`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-wallet-id)

requiredstringrequired, string

The ID of the e-wallet to debit from  
  

Min length: `1`

[`counterparty`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty)

requiredobjectrequired, object

An object representing the e-wallet transaction's counterparty

[`name`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-name)

requiredstringrequired, string

The name of the counterparty  
  

Min length: `1`

[`numbers`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers)

requiredobjectrequired, object

The counterparty's bank account numbers. Exactly one of IBAN or BACS data is required.

[`bacs`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers-bacs)

objectobject

The account number and sort code of the counterparty's account

[`account`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`international`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers-international)

objectobject

International Bank Account Number for a Wallet Transaction

[`iban`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-numbers-international-iban)

stringstring

International Bank Account Number (IBAN).  
  

Min length: `15`

Max length: `34`

[`address`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-address)

objectobject

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-address-street)

required[string]required, [string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-address-city)

requiredstringrequired, string

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-address-postal-code)

requiredstringrequired, string

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-address-country)

requiredstringrequired, string

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

[`date_of_birth`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-counterparty-date-of-birth)

stringstring

The counterparty's birthdate, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format.  
  

Format: `date`

[`amount`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-amount)

requiredobjectrequired, object

The amount and currency of a transaction

[`iso_currency_code`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-amount-iso-currency-code)

requiredstringrequired, string

An ISO-4217 currency code, used with e-wallets and transactions.  
  

Possible values: `GBP`, `EUR`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-amount-value)

requirednumberrequired, number

The amount of the transaction. Must contain at most two digits of precision e.g. `1.23`.  
  

Format: `double`

Minimum: `0.01`

[`reference`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-reference)

requiredstringrequired, string

A reference for the transaction. This must be an alphanumeric string with 6 to 18 characters and must not contain any special characters or spaces.
Ensure that the `reference` field is unique for each transaction.  
  

Max length: `18`

Min length: `6`

[`originating_fund_source`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source)

objectobject

The original source of the funds. This field is required by local regulation for certain businesses (e.g. money remittance) to send payouts to recipients in the EU and UK.

[`full_name`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source-full-name)

requiredstringrequired, string

The full name associated with the source of the funds.

[`address`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source-address)

requiredobjectrequired, object

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source-address-street)

required[string]required, [string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source-address-city)

requiredstringrequired, string

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source-address-postal-code)

requiredstringrequired, string

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source-address-country)

requiredstringrequired, string

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

[`account_number`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source-account-number)

requiredstringrequired, string

The account number from which the funds are sourced.

[`bic`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-request-originating-fund-source-bic)

requiredstringrequired, string

The Business Identifier Code, also known as SWIFT code, for this bank account.  
  

Min length: `8`

Max length: `11`

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

/wallet/transaction/execute

**Response fields**

[`transaction_id`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-response-transaction-id)

stringstring

A unique ID identifying the transaction

[`status`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-response-status)

stringstring

The status of the transaction.  
`AUTHORISING`: The transaction is being processed for validation and compliance.  
`INITIATED`: The transaction has been initiated and is currently being processed.  
`EXECUTED`: The transaction has been successfully executed and is considered complete. This is only applicable for debit transactions.  
`SETTLED`: The transaction has settled and funds are available for use. This is only applicable for credit transactions. A transaction will typically settle within seconds to several days, depending on which payment rail is used.  
`FAILED`: The transaction failed to process successfully. This is a terminal status.  
`BLOCKED`: The transaction has been blocked for violating compliance rules. This is a terminal status.  
  

Possible values: `AUTHORISING`, `INITIATED`, `EXECUTED`, `SETTLED`, `BLOCKED`, `FAILED`

[`request_id`](/docs/api/products/virtual-accounts/#wallet-transaction-execute-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transaction_id": "wallet-transaction-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a88",
  "status": "EXECUTED",
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/wallet/transaction/get`

#### Fetch an e-wallet transaction

Fetch a specific e-wallet transaction

/wallet/transaction/get

**Request fields**

[`client_id`](/docs/api/products/virtual-accounts/#wallet-transaction-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/virtual-accounts/#wallet-transaction-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`transaction_id`](/docs/api/products/virtual-accounts/#wallet-transaction-get-request-transaction-id)

requiredstringrequired, string

The ID of the transaction to fetch  
  

Min length: `1`

/wallet/transaction/get

```
const request: WalletTransactionGetRequest = {
  transaction_id: transactionID,
};
try {
  const response = await plaidClient.walletTransactionGet(request);
  const transactionID = response.data.transaction_id;
  const reference = response.data.reference;
  const type = response.data.type;
  const amount = response.data.amount;
  const counterparty = response.data.counterparty;
  const status = response.data.status;
  const createdAt = response.data.created_at;
} catch (error) {
  // handle error
}
```

/wallet/transaction/get

**Response fields**

[`transaction_id`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-transaction-id)

stringstring

A unique ID identifying the transaction

[`wallet_id`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-wallet-id)

stringstring

The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.

[`reference`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-reference)

stringstring

A reference for the transaction

[`type`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-type)

stringstring

The type of the transaction. The supported transaction types that are returned are:
`BANK_TRANSFER:` a transaction which credits an e-wallet through an external bank transfer.  
`PAYOUT:` a transaction which debits an e-wallet by disbursing funds to a counterparty.  
`PIS_PAY_IN:` a payment which credits an e-wallet through Plaid's Payment Initiation Services (PIS) APIs. For more information see the [Payment Initiation endpoints](https://plaid.com/docs/api/products/payment-initiation/).  
`REFUND:` a transaction which debits an e-wallet by refunding a previously initiated payment made through Plaid's [PIS APIs](https://plaid.com/docs/api/products/payment-initiation/).  
`FUNDS_SWEEP`: an automated transaction which debits funds from an e-wallet to a designated client-owned account.  
`RETURN`: an automated transaction where a debit transaction was reversed and money moved back to originating account.  
`RECALL`: a transaction where the sending bank has requested the return of funds due to a fraud claim, technical error, or other issue associated with the payment.  
  

Possible values: `BANK_TRANSFER`, `PAYOUT`, `PIS_PAY_IN`, `REFUND`, `FUNDS_SWEEP`, `RETURN`, `RECALL`

[`scheme`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-scheme)

nullablestringnullable, string

The payment scheme used to execute this transaction. This is present only for transaction types `PAYOUT` and `REFUND`.  
`FASTER_PAYMENTS`: The standard payment scheme within the UK.  
`SEPA_CREDIT_TRANSFER`: The standard payment to a beneficiary within the SEPA area.  
`SEPA_CREDIT_TRANSFER_INSTANT`: Instant payment to a beneficiary within the SEPA area.  
  

Possible values: `null`, `FASTER_PAYMENTS`, `SEPA_CREDIT_TRANSFER`, `SEPA_CREDIT_TRANSFER_INSTANT`

[`amount`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-amount)

objectobject

The amount and currency of a transaction

[`iso_currency_code`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-amount-iso-currency-code)

stringstring

An ISO-4217 currency code, used with e-wallets and transactions.  
  

Possible values: `GBP`, `EUR`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-amount-value)

numbernumber

The amount of the transaction. Must contain at most two digits of precision e.g. `1.23`.  
  

Format: `double`

Minimum: `0.01`

[`counterparty`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty)

objectobject

An object representing the e-wallet transaction's counterparty

[`name`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-name)

stringstring

The name of the counterparty  
  

Min length: `1`

[`numbers`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-numbers)

objectobject

The counterparty's bank account numbers. Exactly one of IBAN or BACS data is required.

[`bacs`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-numbers-bacs)

nullableobjectnullable, object

The account number and sort code of the counterparty's account

[`account`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-numbers-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-numbers-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`international`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-numbers-international)

nullableobjectnullable, object

International Bank Account Number for a Wallet Transaction

[`iban`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-numbers-international-iban)

stringstring

International Bank Account Number (IBAN).  
  

Min length: `15`

Max length: `34`

[`address`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-address)

nullableobjectnullable, object

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-address-street)

[string][string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-address-city)

stringstring

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-address-postal-code)

stringstring

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-address-country)

stringstring

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

[`date_of_birth`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-counterparty-date-of-birth)

nullablestringnullable, string

The counterparty's birthdate, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format.  
  

Format: `date`

[`status`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-status)

stringstring

The status of the transaction.  
`AUTHORISING`: The transaction is being processed for validation and compliance.  
`INITIATED`: The transaction has been initiated and is currently being processed.  
`EXECUTED`: The transaction has been successfully executed and is considered complete. This is only applicable for debit transactions.  
`SETTLED`: The transaction has settled and funds are available for use. This is only applicable for credit transactions. A transaction will typically settle within seconds to several days, depending on which payment rail is used.  
`FAILED`: The transaction failed to process successfully. This is a terminal status.  
`BLOCKED`: The transaction has been blocked for violating compliance rules. This is a terminal status.  
  

Possible values: `AUTHORISING`, `INITIATED`, `EXECUTED`, `SETTLED`, `BLOCKED`, `FAILED`

[`created_at`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-created-at)

stringstring

Timestamp when the transaction was created, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`last_status_update`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-last-status-update)

stringstring

The date and time of the last time the `status` was updated, in IS0 8601 format  
  

Format: `date-time`

[`payee_verification_status`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-payee-verification-status)

nullablestringnullable, string

Result of payee verification check for EUR payouts. Payee verification checks whether the payee name provided matches the account holder name at the destination institution.  
`FULL_MATCH`: The payee name fully matches the account holder.  
`PARTIAL_MATCH`: The payee name partially matches the account holder.  
`NO_MATCH`: The payee name does not match the account holder.  
`ERROR`: An error occurred during payee verification.  
`CHECK_NOT_POSSIBLE`: Payee verification could not be performed.  
This field is only populated for applicable EUR payout transactions and will be `null` for other transaction types.  
  

Possible values: `FULL_MATCH`, `PARTIAL_MATCH`, `NO_MATCH`, `ERROR`, `CHECK_NOT_POSSIBLE`

[`payment_id`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-payment-id)

nullablestringnullable, string

The payment id that this transaction is associated with, if any. This is present only for transaction types `PIS_PAY_IN` and `REFUND`.

[`failure_reason`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-failure-reason)

nullablestringnullable, string

The error code of a failed transaction. Error codes include:
`EXTERNAL_SYSTEM`: The transaction was declined by an external system.
`EXPIRED`: The transaction request has expired.
`CANCELLED`: The transaction request was rescinded.
`INVALID`: The transaction did not meet certain criteria, such as an inactive account or no valid counterparty, etc.
`UNKNOWN`: The transaction was unsuccessful, but the exact cause is unknown.  
  

Possible values: `EXTERNAL_SYSTEM`, `EXPIRED`, `CANCELLED`, `INVALID`, `UNKNOWN`

[`error`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`related_transactions`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-related-transactions)

[object][object]

A list of wallet transactions that this transaction is associated with, if any.

[`id`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-related-transactions-id)

stringstring

The ID of the related transaction.

[`type`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-related-transactions-type)

stringstring

The type of the transaction.  
  

Possible values: `PAYOUT`, `RETURN`, `REFUND`, `FUNDS_SWEEP`

[`request_id`](/docs/api/products/virtual-accounts/#wallet-transaction-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transaction_id": "wallet-transaction-id-sandbox-feca8a7a-5591-4aef-9297-f3062bb735d3",
  "wallet_id": "wallet-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a88",
  "type": "PAYOUT",
  "reference": "Payout 99744",
  "amount": {
    "iso_currency_code": "GBP",
    "value": 123.12
  },
  "status": "EXECUTED",
  "created_at": "2020-12-02T21:14:54Z",
  "last_status_update": "2020-12-02T21:15:01Z",
  "counterparty": {
    "numbers": {
      "bacs": {
        "account": "31926819",
        "sort_code": "601613"
      }
    },
    "name": "John Smith"
  },
  "request_id": "4zlKapIkTm8p5KM",
  "related_transactions": [
    {
      "id": "wallet-transaction-id-sandbox-2ba30780-d549-4335-b1fe-c2a938aa39d2",
      "type": "RETURN"
    }
  ]
}
```

=\*=\*=\*=

#### `/wallet/transaction/list`

#### List e-wallet transactions

This endpoint lists the latest transactions of the specified e-wallet. Transactions are returned in descending order by the `created_at` time.

/wallet/transaction/list

**Request fields**

[`client_id`](/docs/api/products/virtual-accounts/#wallet-transaction-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/virtual-accounts/#wallet-transaction-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`wallet_id`](/docs/api/products/virtual-accounts/#wallet-transaction-list-request-wallet-id)

requiredstringrequired, string

The ID of the e-wallet to fetch transactions from  
  

Min length: `1`

[`cursor`](/docs/api/products/virtual-accounts/#wallet-transaction-list-request-cursor)

stringstring

A value representing the latest transaction to be included in the response. Set this from `next_cursor` received in the previous `/wallet/transaction/list` request. If provided, the response will only contain that transaction and transactions created before it. If omitted, the response will contain transactions starting from the most recent, and in descending order by the `created_at` time.  
  

Max length: `256`

[`count`](/docs/api/products/virtual-accounts/#wallet-transaction-list-request-count)

integerinteger

The number of transactions to fetch  
  

Minimum: `1`

Maximum: `200`

Default: `10`

[`options`](/docs/api/products/virtual-accounts/#wallet-transaction-list-request-options)

objectobject

Additional wallet transaction options

[`start_time`](/docs/api/products/virtual-accounts/#wallet-transaction-list-request-options-start-time)

stringstring

Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DDThh:mm:ssZ) for filtering transactions, inclusive of the provided date.  
  

Format: `date-time`

[`end_time`](/docs/api/products/virtual-accounts/#wallet-transaction-list-request-options-end-time)

stringstring

Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DDThh:mm:ssZ) for filtering transactions, inclusive of the provided date.  
  

Format: `date-time`

/wallet/transaction/list

```
const request: WalletTransactionListRequest = {
  wallet_id: walletID,
  count: 10,
};
try {
  const response = await plaidClient.walletTransactionList(request);
  const transactions = response.data.transactions;
  const nextCursor = response.data.next_cursor;
} catch (error) {
  // handle error
}
```

/wallet/transaction/list

**Response fields**

[`transactions`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions)

[object][object]

An array of transactions of an e-wallet, associated with the given `wallet_id`

[`transaction_id`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-transaction-id)

stringstring

A unique ID identifying the transaction

[`wallet_id`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-wallet-id)

stringstring

The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.

[`reference`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-reference)

stringstring

A reference for the transaction

[`type`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-type)

stringstring

The type of the transaction. The supported transaction types that are returned are:
`BANK_TRANSFER:` a transaction which credits an e-wallet through an external bank transfer.  
`PAYOUT:` a transaction which debits an e-wallet by disbursing funds to a counterparty.  
`PIS_PAY_IN:` a payment which credits an e-wallet through Plaid's Payment Initiation Services (PIS) APIs. For more information see the [Payment Initiation endpoints](https://plaid.com/docs/api/products/payment-initiation/).  
`REFUND:` a transaction which debits an e-wallet by refunding a previously initiated payment made through Plaid's [PIS APIs](https://plaid.com/docs/api/products/payment-initiation/).  
`FUNDS_SWEEP`: an automated transaction which debits funds from an e-wallet to a designated client-owned account.  
`RETURN`: an automated transaction where a debit transaction was reversed and money moved back to originating account.  
`RECALL`: a transaction where the sending bank has requested the return of funds due to a fraud claim, technical error, or other issue associated with the payment.  
  

Possible values: `BANK_TRANSFER`, `PAYOUT`, `PIS_PAY_IN`, `REFUND`, `FUNDS_SWEEP`, `RETURN`, `RECALL`

[`scheme`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-scheme)

nullablestringnullable, string

The payment scheme used to execute this transaction. This is present only for transaction types `PAYOUT` and `REFUND`.  
`FASTER_PAYMENTS`: The standard payment scheme within the UK.  
`SEPA_CREDIT_TRANSFER`: The standard payment to a beneficiary within the SEPA area.  
`SEPA_CREDIT_TRANSFER_INSTANT`: Instant payment to a beneficiary within the SEPA area.  
  

Possible values: `null`, `FASTER_PAYMENTS`, `SEPA_CREDIT_TRANSFER`, `SEPA_CREDIT_TRANSFER_INSTANT`

[`amount`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-amount)

objectobject

The amount and currency of a transaction

[`iso_currency_code`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-amount-iso-currency-code)

stringstring

An ISO-4217 currency code, used with e-wallets and transactions.  
  

Possible values: `GBP`, `EUR`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-amount-value)

numbernumber

The amount of the transaction. Must contain at most two digits of precision e.g. `1.23`.  
  

Format: `double`

Minimum: `0.01`

[`counterparty`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty)

objectobject

An object representing the e-wallet transaction's counterparty

[`name`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-name)

stringstring

The name of the counterparty  
  

Min length: `1`

[`numbers`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-numbers)

objectobject

The counterparty's bank account numbers. Exactly one of IBAN or BACS data is required.

[`bacs`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-numbers-bacs)

nullableobjectnullable, object

The account number and sort code of the counterparty's account

[`account`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-numbers-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-numbers-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`international`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-numbers-international)

nullableobjectnullable, object

International Bank Account Number for a Wallet Transaction

[`iban`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-numbers-international-iban)

stringstring

International Bank Account Number (IBAN).  
  

Min length: `15`

Max length: `34`

[`address`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-address)

nullableobjectnullable, object

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-address-street)

[string][string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-address-city)

stringstring

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-address-postal-code)

stringstring

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-address-country)

stringstring

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

[`date_of_birth`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-counterparty-date-of-birth)

nullablestringnullable, string

The counterparty's birthdate, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format.  
  

Format: `date`

[`status`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-status)

stringstring

The status of the transaction.  
`AUTHORISING`: The transaction is being processed for validation and compliance.  
`INITIATED`: The transaction has been initiated and is currently being processed.  
`EXECUTED`: The transaction has been successfully executed and is considered complete. This is only applicable for debit transactions.  
`SETTLED`: The transaction has settled and funds are available for use. This is only applicable for credit transactions. A transaction will typically settle within seconds to several days, depending on which payment rail is used.  
`FAILED`: The transaction failed to process successfully. This is a terminal status.  
`BLOCKED`: The transaction has been blocked for violating compliance rules. This is a terminal status.  
  

Possible values: `AUTHORISING`, `INITIATED`, `EXECUTED`, `SETTLED`, `BLOCKED`, `FAILED`

[`created_at`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-created-at)

stringstring

Timestamp when the transaction was created, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`last_status_update`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-last-status-update)

stringstring

The date and time of the last time the `status` was updated, in IS0 8601 format  
  

Format: `date-time`

[`payee_verification_status`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-payee-verification-status)

nullablestringnullable, string

Result of payee verification check for EUR payouts. Payee verification checks whether the payee name provided matches the account holder name at the destination institution.  
`FULL_MATCH`: The payee name fully matches the account holder.  
`PARTIAL_MATCH`: The payee name partially matches the account holder.  
`NO_MATCH`: The payee name does not match the account holder.  
`ERROR`: An error occurred during payee verification.  
`CHECK_NOT_POSSIBLE`: Payee verification could not be performed.  
This field is only populated for applicable EUR payout transactions and will be `null` for other transaction types.  
  

Possible values: `FULL_MATCH`, `PARTIAL_MATCH`, `NO_MATCH`, `ERROR`, `CHECK_NOT_POSSIBLE`

[`payment_id`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-payment-id)

nullablestringnullable, string

The payment id that this transaction is associated with, if any. This is present only for transaction types `PIS_PAY_IN` and `REFUND`.

[`failure_reason`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-failure-reason)

nullablestringnullable, string

The error code of a failed transaction. Error codes include:
`EXTERNAL_SYSTEM`: The transaction was declined by an external system.
`EXPIRED`: The transaction request has expired.
`CANCELLED`: The transaction request was rescinded.
`INVALID`: The transaction did not meet certain criteria, such as an inactive account or no valid counterparty, etc.
`UNKNOWN`: The transaction was unsuccessful, but the exact cause is unknown.  
  

Possible values: `EXTERNAL_SYSTEM`, `EXPIRED`, `CANCELLED`, `INVALID`, `UNKNOWN`

[`error`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`related_transactions`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-related-transactions)

[object][object]

A list of wallet transactions that this transaction is associated with, if any.

[`id`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-related-transactions-id)

stringstring

The ID of the related transaction.

[`type`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-transactions-related-transactions-type)

stringstring

The type of the transaction.  
  

Possible values: `PAYOUT`, `RETURN`, `REFUND`, `FUNDS_SWEEP`

[`next_cursor`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-next-cursor)

stringstring

The value that, when used as the optional `cursor` parameter to `/wallet/transaction/list`, will return the corresponding transaction as its first entry.

[`request_id`](/docs/api/products/virtual-accounts/#wallet-transaction-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "next_cursor": "YWJjMTIzIT8kKiYoKSctPUB",
  "transactions": [
    {
      "transaction_id": "wallet-transaction-id-sandbox-feca8a7a-5591-4aef-9297-f3062bb735d3",
      "wallet_id": "wallet-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a88",
      "type": "PAYOUT",
      "reference": "Payout 99744",
      "amount": {
        "iso_currency_code": "GBP",
        "value": 123.12
      },
      "status": "EXECUTED",
      "created_at": "2020-12-02T21:14:54Z",
      "last_status_update": "2020-12-02T21:15:01Z",
      "counterparty": {
        "numbers": {
          "bacs": {
            "account": "31926819",
            "sort_code": "601613"
          }
        },
        "name": "John Smith"
      },
      "related_transactions": [
        {
          "id": "wallet-transaction-id-sandbox-2ba30780-d549-4335-b1fe-c2a938aa39d2",
          "type": "RETURN"
        }
      ]
    },
    {
      "transaction_id": "wallet-transaction-id-sandbox-feca8a7a-5591-4aef-9297-f3062bb735d3",
      "wallet_id": "wallet-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a88",
      "type": "PAYOUT",
      "reference": "Payout 99744",
      "amount": {
        "iso_currency_code": "EUR",
        "value": 456.78
      },
      "status": "EXECUTED",
      "created_at": "2020-12-02T21:14:54Z",
      "last_status_update": "2020-12-02T21:15:01Z",
      "counterparty": {
        "numbers": {
          "international": {
            "iban": "GB33BUKB20201555555555"
          }
        },
        "name": "John Smith"
      },
      "related_transactions": []
    }
  ],
  "request_id": "4zlKapIkTm8p5KM"
}
```

### Webhooks

Updates are sent to indicate that the status of transaction has changed. All virtual account webhooks have a `webhook_type` of `WALLET`.

=\*=\*=\*=

#### `WALLET_TRANSACTION_STATUS_UPDATE`

Fired when the status of a wallet transaction has changed.

**Properties**

[`webhook_type`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-webhook-type)

stringstring

`WALLET`

[`webhook_code`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-webhook-code)

stringstring

`WALLET_TRANSACTION_STATUS_UPDATE`

[`transaction_id`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-transaction-id)

stringstring

The `transaction_id` for the wallet transaction being updated

[`payment_id`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-payment-id)

stringstring

The `payment_id` associated with the transaction. This will be present in case of `REFUND` and `PIS_PAY_IN`.

[`wallet_id`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-wallet-id)

stringstring

The EMI (E-Money Institution) wallet that this payment is associated with. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.

[`new_status`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-new-status)

stringstring

The status of the transaction.  
`AUTHORISING`: The transaction is being processed for validation and compliance.  
`INITIATED`: The transaction has been initiated and is currently being processed.  
`EXECUTED`: The transaction has been successfully executed and is considered complete. This is only applicable for debit transactions.  
`SETTLED`: The transaction has settled and funds are available for use. This is only applicable for credit transactions. A transaction will typically settle within seconds to several days, depending on which payment rail is used.  
`FAILED`: The transaction failed to process successfully. This is a terminal status.  
`BLOCKED`: The transaction has been blocked for violating compliance rules. This is a terminal status.  
  

Possible values: `AUTHORISING`, `INITIATED`, `EXECUTED`, `SETTLED`, `BLOCKED`, `FAILED`

[`old_status`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-old-status)

stringstring

The status of the transaction.  
`AUTHORISING`: The transaction is being processed for validation and compliance.  
`INITIATED`: The transaction has been initiated and is currently being processed.  
`EXECUTED`: The transaction has been successfully executed and is considered complete. This is only applicable for debit transactions.  
`SETTLED`: The transaction has settled and funds are available for use. This is only applicable for credit transactions. A transaction will typically settle within seconds to several days, depending on which payment rail is used.  
`FAILED`: The transaction failed to process successfully. This is a terminal status.  
`BLOCKED`: The transaction has been blocked for violating compliance rules. This is a terminal status.  
  

Possible values: `AUTHORISING`, `INITIATED`, `EXECUTED`, `SETTLED`, `BLOCKED`, `FAILED`

[`timestamp`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-timestamp)

stringstring

The timestamp of the update, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2017-09-14T14:42:19.350Z"`  
  

Format: `date-time`

[`error`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error)

objectobject

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-error-code-reason)

stringstring

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-display-message)

stringstring

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-status)

integerinteger

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-suggested-action)

stringstring

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`environment`](/docs/api/products/virtual-accounts/#WalletTransactionStatusUpdateWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "WALLET",
  "webhook_code": "WALLET_TRANSACTION_STATUS_UPDATE",
  "transaction_id": "wallet-transaction-id-production-2ba30780-d549-4335-b1fe-c2a938aa39d2",
  "payment_id": "payment-id-production-feca8a7a-5591-4aef-9297-f3062bb735d3",
  "wallet_id": "wallet-id-production-53e58b32-fc1c-46fe-bbd6-e584b27a88",
  "new_status": "SETTLED",
  "old_status": "INITIATED",
  "timestamp": "2017-09-14T14:42:19.350Z",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
