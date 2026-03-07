---
title: "API - Payment Initiation (Europe) | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/payment-initiation/"
scraped_at: "2026-03-07T22:04:18+00:00"
---

# Payment Initiation (UK and Europe)

#### API reference for Payment Initiation endpoints and webhooks

Make payment transfers from your app. Plaid supports both domestic payments denominated in local currencies and international payments, generally denominated in Euro. Domestic payments can be made in pound sterling (typically via the Faster Payments network), Euro (via SEPA Credit Transfer or SEPA Instant) and other local currencies (Polish Zloty, Danish Krone, Swedish Krona, Norwegian Krone), typically via local payment schemes.

For payments in the US, see [Transfer](/docs/api/products/transfer/).

Looking for guidance on how to integrate using these endpoints? Check out the [Payment Initiation documentation](/docs/payment-initiation/).

| Endpoints |  |
| --- | --- |
| [`/payment_initiation/recipient/create`](/docs/api/products/payment-initiation/#payment_initiationrecipientcreate) | Create a recipient |
| [`/payment_initiation/recipient/get`](/docs/api/products/payment-initiation/#payment_initiationrecipientget) | Fetch recipient data |
| [`/payment_initiation/recipient/list`](/docs/api/products/payment-initiation/#payment_initiationrecipientlist) | List all recipients |
| [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) | Create a payment |
| [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) | Fetch a payment |
| [`/payment_initiation/payment/list`](/docs/api/products/payment-initiation/#payment_initiationpaymentlist) | List all payments |
| [`/payment_initiation/payment/reverse`](/docs/api/products/payment-initiation/#payment_initiationpaymentreverse) | Refund a payment from a virtual account |
| [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate) | Create a payment consent |
| [`/payment_initiation/consent/get`](/docs/api/products/payment-initiation/#payment_initiationconsentget) | Fetch a payment consent |
| [`/payment_initiation/consent/revoke`](/docs/api/products/payment-initiation/#payment_initiationconsentrevoke) | Revoke a payment consent |
| [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute) | Execute a payment using payment consent |

Users will be prompted to authorise the payment once you [initialise Link](/docs/link/#initializing-link). See [`/link/token/create`](/docs/api/link/#linktokencreate) for more information on how to obtain a payments `link_token`.

| See also |  |
| --- | --- |
| [`/sandbox/payment/simulate`](/docs/api/sandbox/#sandboxpaymentsimulate) | Simulate a payment in Sandbox |
| [`/wallet/create`](/docs/api/products/virtual-accounts/#walletcreate) | Create a virtual account |
| [`/wallet/get`](/docs/api/products/virtual-accounts/#walletget) | Fetch a virtual account |
| [`/wallet/list`](/docs/api/products/virtual-accounts/#walletlist) | List all virtual accounts |
| [`/wallet/transaction/execute`](/docs/api/products/virtual-accounts/#wallettransactionexecute) | Execute a transaction |
| [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) | Fetch a transaction |
| [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist) | List all transactions |
| [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) | The status of a transaction has changed |

| Webhooks |  |
| --- | --- |
| [`PAYMENT_STATUS_UPDATE`](/docs/api/products/payment-initiation/#payment_status_update) | The status of a payment has changed |
| [`CONSENT_STATUS_UPDATE`](/docs/api/products/payment-initiation/#consent_status_update) | The status of a consent has changed |

### Endpoints

=\*=\*=\*=

#### `/payment_initiation/recipient/create`

#### Create payment recipient

Create a payment recipient for payment initiation. The recipient must be in Europe, within a country that is a member of the Single Euro Payment Area (SEPA) or a non-Eurozone country [supported](https://plaid.com/global) by Plaid. For a standing order (recurring) payment, the recipient must be in the UK.

It is recommended to use `bacs` in the UK and `iban` in EU.

The endpoint is idempotent: if a developer has already made a request with the same payment details, Plaid will return the same `recipient_id`.

/payment\_initiation/recipient/create

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`name`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-name)

requiredstringrequired, string

The name of the recipient. We recommend using strings of length 18 or less and avoid special characters to ensure compatibility with all institutions.  
  

Min length: `1`

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-iban)

stringstring

The International Bank Account Number (IBAN) for the recipient. If BACS data is not provided, an IBAN is required.  
  

Min length: `15`

Max length: `34`

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-bacs)

objectobject

An object containing a BACS account number and sort code. If an IBAN is not provided or if this recipient needs to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`address`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-address)

objectobject

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-address-street)

required[string]required, [string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-address-city)

requiredstringrequired, string

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-address-postal-code)

requiredstringrequired, string

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-request-address-country)

requiredstringrequired, string

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

/payment\_initiation/recipient/create

```
// Using BACS, without IBAN or address
const request: PaymentInitiationRecipientCreateRequest = {
  name: 'John Doe',
  bacs: {
    account: '26207729',
    sort_code: '560029',
  },
};
try {
  const response = await plaidClient.paymentInitiationRecipientCreate(request);
  const recipientID = response.data.recipient_id;
} catch (error) {
  // handle error
}
```

/payment\_initiation/recipient/create

**Response fields**

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-response-recipient-id)

stringstring

A unique ID identifying the recipient

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "recipient_id": "recipient-id-sandbox-9b6b4679-914b-445b-9450-efbdb80296f6",
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/payment_initiation/recipient/get`

#### Get payment recipient

Get details about a payment recipient you have previously created.

/payment\_initiation/recipient/get

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-request-recipient-id)

requiredstringrequired, string

The ID of the recipient

```
const request: PaymentInitiationRecipientGetRequest = {
  recipient_id: recipientID,
};
try {
  const response = await plaidClient.paymentInitiationRecipientGet(request);
  const recipientID = response.data.recipient_id;
  const name = response.data.name;
  const iban = response.data.iban;
  const address = response.data.address;
} catch (error) {
  // handle error
}
```

/payment\_initiation/recipient/get

**Response fields**

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-recipient-id)

stringstring

The ID of the recipient.

[`name`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-name)

stringstring

The name of the recipient.

[`address`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-address)

nullableobjectnullable, object

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-address-street)

[string][string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-address-city)

stringstring

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-address-postal-code)

stringstring

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-address-country)

stringstring

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-iban)

nullablestringnullable, string

The International Bank Account Number (IBAN) for the recipient.

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if this recipient needs to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "recipient_id": "recipient-id-sandbox-9b6b4679-914b-445b-9450-efbdb80296f6",
  "name": "Wonder Wallet",
  "iban": "GB29NWBK60161331926819",
  "address": {
    "street": [
      "96 Guild Street",
      "9th Floor"
    ],
    "city": "London",
    "postal_code": "SE14 8JW",
    "country": "GB"
  },
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/payment_initiation/recipient/list`

#### List payment recipients

The [`/payment_initiation/recipient/list`](/docs/api/products/payment-initiation/#payment_initiationrecipientlist) endpoint list the payment recipients that you have previously created.

/payment\_initiation/recipient/list

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`count`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-request-count)

integerinteger

The maximum number of recipients to return. If `count` is not specified, a maximum of 100 recipients will be returned, beginning with the recipient at the cursor (if specified).  
  

Minimum: `1`

Maximum: `100`

Default: `100`

[`cursor`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-request-cursor)

stringstring

A value representing the latest recipient to be included in the response. Set this from `next_cursor` received from the previous `/payment_initiation/recipient/list` request. If provided, the response will only contain that recipient and recipients created before it. If omitted, the response will contain recipients starting from the most recent, and in descending order by the `created_at` time.  
  

Max length: `256`

```
try {
  const response = await plaidClient.paymentInitiationRecipientList({});
  const recipients = response.data.recipients;
} catch (error) {
  // handle error
}
```

/payment\_initiation/recipient/list

**Response fields**

[`recipients`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients)

[object][object]

An array of payment recipients created for Payment Initiation

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-recipient-id)

stringstring

The ID of the recipient.

[`name`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-name)

stringstring

The name of the recipient.

[`address`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-address)

nullableobjectnullable, object

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-address-street)

[string][string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-address-city)

stringstring

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-address-postal-code)

stringstring

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-address-country)

stringstring

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-iban)

nullablestringnullable, string

The International Bank Account Number (IBAN) for the recipient.

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if this recipient needs to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-recipients-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`next_cursor`](/docs/api/products/payment-initiation/#payment_initiation-recipient-list-response-next-cursor)

stringstring

The value that, when used as the optional `cursor` parameter to `/payment_initiation/recipient/list`, will return the corresponding recipient as its first recipient.  
  

Max length: `256`

Response Object

```
{
  "recipients": [
    {
      "recipient_id": "recipient-id-sandbox-9b6b4679-914b-445b-9450-efbdb80296f6",
      "name": "Wonder Wallet",
      "iban": "GB29NWBK60161331926819",
      "address": {
        "street": [
          "96 Guild Street",
          "9th Floor"
        ],
        "city": "London",
        "postal_code": "SE14 8JW",
        "country": "GB"
      }
    }
  ],
  "next_cursor": "YWJjMTIzIT8kKiYoKSctPUE",
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/payment_initiation/payment/create`

#### Create a payment

After creating a payment recipient, you can use the [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) endpoint to create a payment to that recipient. Payments can be one-time or standing order (recurring) and can be denominated in either EUR, GBP or other chosen [currency](https://plaid.com/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-amount-currency). If making domestic GBP-denominated payments, your recipient must have been created with BACS numbers. In general, EUR-denominated payments will be sent via SEPA Credit Transfer, GBP-denominated payments will be sent via the Faster Payments network and for non-Eurozone markets typically via the local payment scheme, but the payment network used will be determined by the institution. Payments sent via Faster Payments will typically arrive immediately, while payments sent via SEPA Credit Transfer or other local payment schemes will typically arrive in one business day.

Standing orders (recurring payments) must be denominated in GBP and can only be sent to recipients in the UK. Once created, standing order payments cannot be modified or canceled via the API. An end user can cancel or modify a standing order directly on their banking application or website, or by contacting the bank. Standing orders will follow the payment rules of the underlying rails (Faster Payments in UK). Payments can be sent Monday to Friday, excluding bank holidays. If the pre-arranged date falls on a weekend or bank holiday, the payment is made on the next working day. It is not possible to guarantee the exact time the payment will reach the recipient’s account, although at least 90% of standing order payments are sent by 6am.

In Limited Production, payments must be below 5 GBP or other chosen [currency](https://plaid.com/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-amount-currency), and standing orders, variable recurring payments, and Virtual Accounts are not supported.

/payment\_initiation/payment/create

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-recipient-id)

requiredstringrequired, string

The ID of the recipient the payment is for.

[`reference`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-reference)

requiredstringrequired, string

A reference for the payment. This must be an alphanumeric string with at most 18 characters and must not contain any special characters (since not all institutions support them).
In order to track settlement via Payment Confirmation, each payment must have a unique reference. If the reference provided through the API is not unique, Plaid will adjust it.
Some institutions may limit the reference to less than 18 characters. If necessary, Plaid will adjust the reference by truncating it to fit the institution's requirements.
Both the originally provided and automatically adjusted references (if any) can be found in the `reference` and `adjusted_reference` fields, respectively.  
  

Min length: `1`

Max length: `18`

[`amount`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-amount)

requiredobjectrequired, object

The amount and currency of a payment

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-amount-currency)

requiredstringrequired, string

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-amount-value)

requirednumberrequired, number

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.  
  

Format: `double`

[`schedule`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-schedule)

objectobject

The schedule that the payment will be executed on. If a schedule is provided, the payment is automatically set up as a standing order. If no schedule is specified, the payment will be executed only once.

[`interval`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-schedule-interval)

requiredstringrequired, string

The frequency interval of the payment.  
  

Possible values: `WEEKLY`, `MONTHLY`

Min length: `1`

[`interval_execution_day`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-schedule-interval-execution-day)

requiredintegerrequired, integer

The day of the interval on which to schedule the payment.  
If the payment interval is weekly, `interval_execution_day` should be an integer from 1 (Monday) to 7 (Sunday).  
If the payment interval is monthly, `interval_execution_day` should be an integer indicating which day of the month to make the payment on. Integers from 1 to 28 can be used to make a payment on that day of the month. Negative integers from -1 to -5 can be used to make a payment relative to the end of the month. To make a payment on the last day of the month, use -1; to make the payment on the second-to-last day, use -2, and so on.

[`start_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-schedule-start-date)

requiredstringrequired, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will begin on the first `interval_execution_day` on or after the `start_date`.  
If the first `interval_execution_day` on or after the start date is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`end_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-schedule-end-date)

stringstring

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will end on the last `interval_execution_day` on or before the `end_date`.
If the only `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`adjusted_start_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-schedule-adjusted-start-date)

stringstring

The start date sent to the bank after adjusting for holidays or weekends. Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, this field will be `null`.  
  

Format: `date`

[`options`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options)

objectobject

Additional payment options

[`request_refund_details`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options-request-refund-details)

booleanboolean

When `true`, Plaid will attempt to request refund details from the payee's financial institution. Support varies between financial institutions and will not always be available. If refund details could be retrieved, they will be available in the `/payment_initiation/payment/get` response.

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options-iban)

stringstring

The International Bank Account Number (IBAN) for the payer's account. Where possible, the end user will be able to send payments only from the specified bank account if provided.  
  

Min length: `15`

Max length: `34`

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options-bacs)

objectobject

An optional object used to restrict the accounts used for payments. If provided, the end user will be able to send payments only from the specified bank account.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`scheme`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options-scheme)

stringstring

Payment scheme. If not specified - the default in the region will be used (e.g. `SEPA_CREDIT_TRANSFER` for EU). In responses, if the scheme is not explicitly specified in the request, this value will be `null`. Using unsupported values will result in a failed payment.  
`LOCAL_DEFAULT`: The default payment scheme for the selected market and currency will be used.  
`LOCAL_INSTANT`: The instant payment scheme for the selected market and currency will be used (if applicable). Fees may be applied by the institution.  
`SEPA_CREDIT_TRANSFER`: The standard payment to a beneficiary within the SEPA area.  
`SEPA_CREDIT_TRANSFER_INSTANT`: Instant payment within the SEPA area. May involve additional fees and may not be available at some banks.  
  

Possible values: `null`, `LOCAL_DEFAULT`, `LOCAL_INSTANT`, `SEPA_CREDIT_TRANSFER`, `SEPA_CREDIT_TRANSFER_INSTANT`

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

/payment\_initiation/payment/create

**Response fields**

[`payment_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-response-payment-id)

stringstring

A unique ID identifying the payment

[`status`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-response-status)

stringstring

For a payment returned by this endpoint, there is only one possible value:  
`PAYMENT_STATUS_INPUT_NEEDED`: The initial phase of the payment  
  

Possible values: `PAYMENT_STATUS_INPUT_NEEDED`

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "payment_id": "payment-id-sandbox-feca8a7a-5591-4aef-9297-f3062bb735d3",
  "status": "PAYMENT_STATUS_INPUT_NEEDED",
  "request_id": "4ciYVmesrySiUAB"
}
```

=\*=\*=\*=

#### `/payment_initiation/payment/get`

#### Get payment details

The [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) endpoint can be used to check the status of a payment, as well as to receive basic information such as recipient and payment amount. In the case of standing orders, the [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) endpoint will provide information about the status of the overall standing order itself; the API cannot be used to retrieve payment status for individual payments within a standing order.

Polling for status updates in Production is highly discouraged. Repeatedly calling [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) to check a payment's status is unreliable and may trigger API rate limits. Only the `payment_status_update` webhook should be used to receive real-time status updates in Production.

In the case of standing orders, the [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) endpoint will provide information about the status of the overall standing order itself; the API cannot be used to retrieve payment status for individual payments within a standing order.

/payment\_initiation/payment/get

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`payment_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-request-payment-id)

requiredstringrequired, string

The `payment_id` returned from `/payment_initiation/payment/create`.

/payment\_initiation/payment/create

```
const request: PaymentInitiationPaymentGetRequest = {
  payment_id: paymentID,
};
try {
  const response = await plaidClient.paymentInitiationPaymentGet(request);
  const paymentID = response.data.payment_id;
  const paymentToken = response.data.payment_token;
  const reference = response.data.reference;
  const amount = response.data.amount;
  const status = response.data.status;
  const lastStatusUpdate = response.data.last_status_update;
  const paymentTokenExpirationTime =
    response.data.payment_token_expiration_time;
  const recipientID = response.data.recipient_id;
} catch (error) {
  // handle error
}
```

/payment\_initiation/payment/get

**Response fields**

[`payment_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-payment-id)

stringstring

The ID of the payment. Like all Plaid identifiers, the `payment_id` is case sensitive.

[`amount`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-amount)

objectobject

The amount and currency of a payment

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-amount-currency)

stringstring

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-amount-value)

numbernumber

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.  
  

Format: `double`

[`status`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-status)

stringstring

The status of the payment.  
Core lifecycle statuses:  
**`PAYMENT_STATUS_INPUT_NEEDED`**: Transitional. The payment is awaiting user input to continue processing. It may re-enter this state if additional input is required.  
**`PAYMENT_STATUS_AUTHORISING`:** Transitional. The payment is being authorised by the financial institution. It will automatically move on once authorisation completes.  
**`PAYMENT_STATUS_INITIATED`:** Transitional. The payment has been authorised and accepted by the financial institution and is now in transit. A payment should be considered complete once it reaches the `PAYMENT_STATUS_EXECUTED` state or the funds settle in the recipient account.  
**`PAYMENT_STATUS_EXECUTED`: Terminal.** The funds have left the payer’s account and the payment is en route to settlement. Support is more common in the UK than in the EU; where unsupported, a successful payment remains in `PAYMENT_STATUS_INITIATED` before settling. When using Plaid Virtual Accounts, `PAYMENT_STATUS_EXECUTED` is not terminal—the payment will continue to `PAYMENT_STATUS_SETTLED` once funds are available.  
**`PAYMENT_STATUS_SETTLED`: Terminal.** The funds are available in the recipient’s account. Only available to customers using [Plaid Virtual Accounts](https://plaid.com/docs/payment-initiation/virtual-accounts/).  
Failure statuses:  
**`PAYMENT_STATUS_INSUFFICIENT_FUNDS`: Terminal.** The payment failed due to insufficient funds. No further retries will succeed until the payer’s balance is replenished.  
**`PAYMENT_STATUS_FAILED`: Terminal (retryable).** The payment could not be initiated due to a system error or outage. Retry once the root cause is resolved.  
**`PAYMENT_STATUS_BLOCKED`: Terminal (retryable).** The payment was blocked by Plaid (e.g., flagged as risky). Resolve any compliance or risk issues and retry.  
**`PAYMENT_STATUS_REJECTED`: Terminal.** The payment was rejected by the financial institution. No automatic retry is possible.  
**`PAYMENT_STATUS_CANCELLED`: Terminal.** The end user cancelled the payment during authorisation.  
Standing-order statuses:  
**`PAYMENT_STATUS_ESTABLISHED`: Terminal.** A recurring/standing order has been successfully created.  
Deprecated (to be removed in a future release):  
`PAYMENT_STATUS_UNKNOWN`: The payment status is unknown.  
`PAYMENT_STATUS_PROCESSING`: The payment is currently being processed.  
`PAYMENT_STATUS_COMPLETED`: Indicates that the standing order has been successfully established.  
  

Possible values: `PAYMENT_STATUS_INPUT_NEEDED`, `PAYMENT_STATUS_PROCESSING`, `PAYMENT_STATUS_INITIATED`, `PAYMENT_STATUS_COMPLETED`, `PAYMENT_STATUS_INSUFFICIENT_FUNDS`, `PAYMENT_STATUS_FAILED`, `PAYMENT_STATUS_BLOCKED`, `PAYMENT_STATUS_UNKNOWN`, `PAYMENT_STATUS_EXECUTED`, `PAYMENT_STATUS_SETTLED`, `PAYMENT_STATUS_AUTHORISING`, `PAYMENT_STATUS_CANCELLED`, `PAYMENT_STATUS_ESTABLISHED`, `PAYMENT_STATUS_REJECTED`

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-recipient-id)

stringstring

The ID of the recipient

[`reference`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-reference)

stringstring

A reference for the payment.

[`adjusted_reference`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-adjusted-reference)

nullablestringnullable, string

The value of the reference sent to the bank after adjustment to pass bank validation rules.

[`last_status_update`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-last-status-update)

stringstring

The date and time of the last time the `status` was updated, in IS0 8601 format  
  

Format: `date-time`

[`schedule`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-schedule)

nullableobjectnullable, object

The schedule that the payment will be executed on. If a schedule is provided, the payment is automatically set up as a standing order. If no schedule is specified, the payment will be executed only once.

[`interval`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-schedule-interval)

stringstring

The frequency interval of the payment.  
  

Possible values: `WEEKLY`, `MONTHLY`

Min length: `1`

[`interval_execution_day`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-schedule-interval-execution-day)

integerinteger

The day of the interval on which to schedule the payment.  
If the payment interval is weekly, `interval_execution_day` should be an integer from 1 (Monday) to 7 (Sunday).  
If the payment interval is monthly, `interval_execution_day` should be an integer indicating which day of the month to make the payment on. Integers from 1 to 28 can be used to make a payment on that day of the month. Negative integers from -1 to -5 can be used to make a payment relative to the end of the month. To make a payment on the last day of the month, use -1; to make the payment on the second-to-last day, use -2, and so on.

[`start_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-schedule-start-date)

stringstring

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will begin on the first `interval_execution_day` on or after the `start_date`.  
If the first `interval_execution_day` on or after the start date is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`end_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-schedule-end-date)

nullablestringnullable, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will end on the last `interval_execution_day` on or before the `end_date`.
If the only `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`adjusted_start_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-schedule-adjusted-start-date)

nullablestringnullable, string

The start date sent to the bank after adjusting for holidays or weekends. Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, this field will be `null`.  
  

Format: `date`

[`refund_details`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-refund-details)

nullableobjectnullable, object

Details about external payment refund

[`name`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-refund-details-name)

stringstring

The name of the account holder.

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-refund-details-iban)

nullablestringnullable, string

The International Bank Account Number (IBAN) for the account.

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-refund-details-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if this recipient needs to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-refund-details-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-refund-details-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if this recipient needs to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-iban)

nullablestringnullable, string

The International Bank Account Number (IBAN) for the sender, if specified in the `/payment_initiation/payment/create` call.

[`refund_ids`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-refund-ids)

nullable[string]nullable, [string]

Refund IDs associated with the payment.

[`amount_refunded`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-amount-refunded)

nullableobjectnullable, object

The amount and currency of a payment

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-amount-refunded-currency)

stringstring

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-amount-refunded-value)

numbernumber

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`.  
  

Format: `double`

Minimum: `0.01`

[`wallet_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-wallet-id)

nullablestringnullable, string

The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.

[`scheme`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-scheme)

nullablestringnullable, string

Payment scheme. If not specified - the default in the region will be used (e.g. `SEPA_CREDIT_TRANSFER` for EU). In responses, if the scheme is not explicitly specified in the request, this value will be `null`. Using unsupported values will result in a failed payment.  
`LOCAL_DEFAULT`: The default payment scheme for the selected market and currency will be used.  
`LOCAL_INSTANT`: The instant payment scheme for the selected market and currency will be used (if applicable). Fees may be applied by the institution.  
`SEPA_CREDIT_TRANSFER`: The standard payment to a beneficiary within the SEPA area.  
`SEPA_CREDIT_TRANSFER_INSTANT`: Instant payment within the SEPA area. May involve additional fees and may not be available at some banks.  
  

Possible values: `null`, `LOCAL_DEFAULT`, `LOCAL_INSTANT`, `SEPA_CREDIT_TRANSFER`, `SEPA_CREDIT_TRANSFER_INSTANT`

[`adjusted_scheme`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-adjusted-scheme)

nullablestringnullable, string

Payment scheme. If not specified - the default in the region will be used (e.g. `SEPA_CREDIT_TRANSFER` for EU). In responses, if the scheme is not explicitly specified in the request, this value will be `null`. Using unsupported values will result in a failed payment.  
`LOCAL_DEFAULT`: The default payment scheme for the selected market and currency will be used.  
`LOCAL_INSTANT`: The instant payment scheme for the selected market and currency will be used (if applicable). Fees may be applied by the institution.  
`SEPA_CREDIT_TRANSFER`: The standard payment to a beneficiary within the SEPA area.  
`SEPA_CREDIT_TRANSFER_INSTANT`: Instant payment within the SEPA area. May involve additional fees and may not be available at some banks.  
  

Possible values: `null`, `LOCAL_DEFAULT`, `LOCAL_INSTANT`, `SEPA_CREDIT_TRANSFER`, `SEPA_CREDIT_TRANSFER_INSTANT`

[`consent_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-consent-id)

nullablestringnullable, string

The payment consent ID that this payment was initiated with. Is present only when payment was initiated using the payment consent.

[`transaction_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-transaction-id)

nullablestringnullable, string

The transaction ID that this payment is associated with, if any. This is present only when a payment was initiated using virtual accounts.

[`end_to_end_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-end-to-end-id)

nullablestringnullable, string

A unique identifier assigned by Plaid to each payment for tracking and reconciliation purposes.  
Note: Not all banks handle `end_to_end_id` consistently. To ensure accurate matching, clients should convert both the incoming `end_to_end_id` and the one provided by Plaid to the same case (either lower or upper) before comparison. For virtual account payments, Plaid manages this field automatically.

[`error`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "payment_id": "payment-id-sandbox-feca8a7a-5591-4aef-9297-f3062bb735d3",
  "reference": "Account Funding 99744",
  "amount": {
    "currency": "GBP",
    "value": 100
  },
  "status": "PAYMENT_STATUS_INITIATED",
  "last_status_update": "2019-11-06T21:10:52Z",
  "recipient_id": "recipient-id-sandbox-9b6b4679-914b-445b-9450-efbdb80296f6",
  "bacs": {
    "account": "31926819",
    "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
    "sort_code": "601613"
  },
  "end_to_end_id": "sptch8cde8390bfd363888",
  "iban": null,
  "request_id": "aEAQmewMzlVa1k6"
}
```

=\*=\*=\*=

#### `/payment_initiation/payment/list`

#### List payments

The [`/payment_initiation/payment/list`](/docs/api/products/payment-initiation/#payment_initiationpaymentlist) endpoint can be used to retrieve all created payments. By default, the 10 most recent payments are returned. You can request more payments and paginate through the results using the optional `count` and `cursor` parameters.

/payment\_initiation/payment/list

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`count`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-request-count)

integerinteger

The maximum number of payments to return. If `count` is not specified, a maximum of 10 payments will be returned, beginning with the most recent payment before the cursor (if specified).  
  

Minimum: `1`

Maximum: `200`

Default: `10`

[`cursor`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-request-cursor)

stringstring

A string in RFC 3339 format (i.e. "2019-12-06T22:35:49Z"). Only payments created before the cursor will be returned.  
  

Format: `date-time`

[`consent_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-request-consent-id)

stringstring

The consent ID. If specified, only payments, executed using this consent, will be returned.

/payment\_initiation/payment/list

```
const request: PaymentInitiationPaymentListRequest = {
  count: 10,
  cursor: '2019-12-06T22:35:49Z',
};
try {
  const response = await plaidClient.paymentInitiationPaymentList(request);
  const payments = response.data.payments;
  const nextCursor = response.data.next_cursor;
} catch (error) {
  // handle error
}
```

/payment\_initiation/payment/list

**Response fields**

[`payments`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments)

[object][object]

An array of payments that have been created, associated with the given `client_id`.

[`payment_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-payment-id)

stringstring

The ID of the payment. Like all Plaid identifiers, the `payment_id` is case sensitive.

[`amount`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-amount)

objectobject

The amount and currency of a payment

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-amount-currency)

stringstring

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-amount-value)

numbernumber

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.  
  

Format: `double`

[`status`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-status)

stringstring

The status of the payment.  
Core lifecycle statuses:  
**`PAYMENT_STATUS_INPUT_NEEDED`**: Transitional. The payment is awaiting user input to continue processing. It may re-enter this state if additional input is required.  
**`PAYMENT_STATUS_AUTHORISING`:** Transitional. The payment is being authorised by the financial institution. It will automatically move on once authorisation completes.  
**`PAYMENT_STATUS_INITIATED`:** Transitional. The payment has been authorised and accepted by the financial institution and is now in transit. A payment should be considered complete once it reaches the `PAYMENT_STATUS_EXECUTED` state or the funds settle in the recipient account.  
**`PAYMENT_STATUS_EXECUTED`: Terminal.** The funds have left the payer’s account and the payment is en route to settlement. Support is more common in the UK than in the EU; where unsupported, a successful payment remains in `PAYMENT_STATUS_INITIATED` before settling. When using Plaid Virtual Accounts, `PAYMENT_STATUS_EXECUTED` is not terminal—the payment will continue to `PAYMENT_STATUS_SETTLED` once funds are available.  
**`PAYMENT_STATUS_SETTLED`: Terminal.** The funds are available in the recipient’s account. Only available to customers using [Plaid Virtual Accounts](https://plaid.com/docs/payment-initiation/virtual-accounts/).  
Failure statuses:  
**`PAYMENT_STATUS_INSUFFICIENT_FUNDS`: Terminal.** The payment failed due to insufficient funds. No further retries will succeed until the payer’s balance is replenished.  
**`PAYMENT_STATUS_FAILED`: Terminal (retryable).** The payment could not be initiated due to a system error or outage. Retry once the root cause is resolved.  
**`PAYMENT_STATUS_BLOCKED`: Terminal (retryable).** The payment was blocked by Plaid (e.g., flagged as risky). Resolve any compliance or risk issues and retry.  
**`PAYMENT_STATUS_REJECTED`: Terminal.** The payment was rejected by the financial institution. No automatic retry is possible.  
**`PAYMENT_STATUS_CANCELLED`: Terminal.** The end user cancelled the payment during authorisation.  
Standing-order statuses:  
**`PAYMENT_STATUS_ESTABLISHED`: Terminal.** A recurring/standing order has been successfully created.  
Deprecated (to be removed in a future release):  
`PAYMENT_STATUS_UNKNOWN`: The payment status is unknown.  
`PAYMENT_STATUS_PROCESSING`: The payment is currently being processed.  
`PAYMENT_STATUS_COMPLETED`: Indicates that the standing order has been successfully established.  
  

Possible values: `PAYMENT_STATUS_INPUT_NEEDED`, `PAYMENT_STATUS_PROCESSING`, `PAYMENT_STATUS_INITIATED`, `PAYMENT_STATUS_COMPLETED`, `PAYMENT_STATUS_INSUFFICIENT_FUNDS`, `PAYMENT_STATUS_FAILED`, `PAYMENT_STATUS_BLOCKED`, `PAYMENT_STATUS_UNKNOWN`, `PAYMENT_STATUS_EXECUTED`, `PAYMENT_STATUS_SETTLED`, `PAYMENT_STATUS_AUTHORISING`, `PAYMENT_STATUS_CANCELLED`, `PAYMENT_STATUS_ESTABLISHED`, `PAYMENT_STATUS_REJECTED`

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-recipient-id)

stringstring

The ID of the recipient

[`reference`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-reference)

stringstring

A reference for the payment.

[`adjusted_reference`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-adjusted-reference)

nullablestringnullable, string

The value of the reference sent to the bank after adjustment to pass bank validation rules.

[`last_status_update`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-last-status-update)

stringstring

The date and time of the last time the `status` was updated, in IS0 8601 format  
  

Format: `date-time`

[`schedule`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-schedule)

nullableobjectnullable, object

The schedule that the payment will be executed on. If a schedule is provided, the payment is automatically set up as a standing order. If no schedule is specified, the payment will be executed only once.

[`interval`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-schedule-interval)

stringstring

The frequency interval of the payment.  
  

Possible values: `WEEKLY`, `MONTHLY`

Min length: `1`

[`interval_execution_day`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-schedule-interval-execution-day)

integerinteger

The day of the interval on which to schedule the payment.  
If the payment interval is weekly, `interval_execution_day` should be an integer from 1 (Monday) to 7 (Sunday).  
If the payment interval is monthly, `interval_execution_day` should be an integer indicating which day of the month to make the payment on. Integers from 1 to 28 can be used to make a payment on that day of the month. Negative integers from -1 to -5 can be used to make a payment relative to the end of the month. To make a payment on the last day of the month, use -1; to make the payment on the second-to-last day, use -2, and so on.

[`start_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-schedule-start-date)

stringstring

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will begin on the first `interval_execution_day` on or after the `start_date`.  
If the first `interval_execution_day` on or after the start date is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`end_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-schedule-end-date)

nullablestringnullable, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will end on the last `interval_execution_day` on or before the `end_date`.
If the only `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`adjusted_start_date`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-schedule-adjusted-start-date)

nullablestringnullable, string

The start date sent to the bank after adjusting for holidays or weekends. Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, this field will be `null`.  
  

Format: `date`

[`refund_details`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-refund-details)

nullableobjectnullable, object

Details about external payment refund

[`name`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-refund-details-name)

stringstring

The name of the account holder.

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-refund-details-iban)

nullablestringnullable, string

The International Bank Account Number (IBAN) for the account.

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-refund-details-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if this recipient needs to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-refund-details-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-refund-details-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if this recipient needs to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-iban)

nullablestringnullable, string

The International Bank Account Number (IBAN) for the sender, if specified in the `/payment_initiation/payment/create` call.

[`refund_ids`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-refund-ids)

nullable[string]nullable, [string]

Refund IDs associated with the payment.

[`amount_refunded`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-amount-refunded)

nullableobjectnullable, object

The amount and currency of a payment

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-amount-refunded-currency)

stringstring

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-amount-refunded-value)

numbernumber

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`.  
  

Format: `double`

Minimum: `0.01`

[`wallet_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-wallet-id)

nullablestringnullable, string

The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.

[`scheme`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-scheme)

nullablestringnullable, string

Payment scheme. If not specified - the default in the region will be used (e.g. `SEPA_CREDIT_TRANSFER` for EU). In responses, if the scheme is not explicitly specified in the request, this value will be `null`. Using unsupported values will result in a failed payment.  
`LOCAL_DEFAULT`: The default payment scheme for the selected market and currency will be used.  
`LOCAL_INSTANT`: The instant payment scheme for the selected market and currency will be used (if applicable). Fees may be applied by the institution.  
`SEPA_CREDIT_TRANSFER`: The standard payment to a beneficiary within the SEPA area.  
`SEPA_CREDIT_TRANSFER_INSTANT`: Instant payment within the SEPA area. May involve additional fees and may not be available at some banks.  
  

Possible values: `null`, `LOCAL_DEFAULT`, `LOCAL_INSTANT`, `SEPA_CREDIT_TRANSFER`, `SEPA_CREDIT_TRANSFER_INSTANT`

[`adjusted_scheme`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-adjusted-scheme)

nullablestringnullable, string

Payment scheme. If not specified - the default in the region will be used (e.g. `SEPA_CREDIT_TRANSFER` for EU). In responses, if the scheme is not explicitly specified in the request, this value will be `null`. Using unsupported values will result in a failed payment.  
`LOCAL_DEFAULT`: The default payment scheme for the selected market and currency will be used.  
`LOCAL_INSTANT`: The instant payment scheme for the selected market and currency will be used (if applicable). Fees may be applied by the institution.  
`SEPA_CREDIT_TRANSFER`: The standard payment to a beneficiary within the SEPA area.  
`SEPA_CREDIT_TRANSFER_INSTANT`: Instant payment within the SEPA area. May involve additional fees and may not be available at some banks.  
  

Possible values: `null`, `LOCAL_DEFAULT`, `LOCAL_INSTANT`, `SEPA_CREDIT_TRANSFER`, `SEPA_CREDIT_TRANSFER_INSTANT`

[`consent_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-consent-id)

nullablestringnullable, string

The payment consent ID that this payment was initiated with. Is present only when payment was initiated using the payment consent.

[`transaction_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-transaction-id)

nullablestringnullable, string

The transaction ID that this payment is associated with, if any. This is present only when a payment was initiated using virtual accounts.

[`end_to_end_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-end-to-end-id)

nullablestringnullable, string

A unique identifier assigned by Plaid to each payment for tracking and reconciliation purposes.  
Note: Not all banks handle `end_to_end_id` consistently. To ensure accurate matching, clients should convert both the incoming `end_to_end_id` and the one provided by Plaid to the same case (either lower or upper) before comparison. For virtual account payments, Plaid manages this field automatically.

[`error`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-payments-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`next_cursor`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-next-cursor)

nullablestringnullable, string

The value that, when used as the optional `cursor` parameter to `/payment_initiation/payment/list`, will return the next unreturned payment as its first payment.  
  

Format: `date-time`

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "payments": [
    {
      "payment_id": "payment-id-sandbox-feca8a7a-5581-4aef-9297-f3062bb735d3",
      "reference": "Account Funding 99744",
      "amount": {
        "currency": "GBP",
        "value": 100
      },
      "status": "PAYMENT_STATUS_EXECUTED",
      "last_status_update": "2019-11-06T21:10:52Z",
      "recipient_id": "recipient-id-sandbox-9b6b4679-914b-445b-9450-efbdb80296f6",
      "bacs": {
        "account": "31926819",
        "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
        "sort_code": "601613"
      },
      "iban": "null,",
      "end_to_end_id": "sptch8cde8390bfd363888,"
    }
  ],
  "next_cursor": "2020-01-01T00:00:00Z",
  "request_id": "aEAQmewMzlVa1k6"
}
```

=\*=\*=\*=

#### `/payment_initiation/payment/reverse`

#### Reverse an existing payment

Reverse a settled payment from a Plaid virtual account.

The original payment must be in a settled state to be refunded.
To refund partially, specify the amount as part of the request.
If the amount is not specified, the refund amount will be equal to all
of the remaining payment amount that has not been refunded yet.

The refund will go back to the source account that initiated the payment.
The original payment must have been initiated to a Plaid virtual account
so that this account can be used to initiate the refund.

Providing counterparty information such as date of birth and address increases
the likelihood of refund being successful without human intervention.

/payment\_initiation/payment/reverse

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`payment_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-payment-id)

requiredstringrequired, string

The ID of the payment to reverse

[`idempotency_key`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-idempotency-key)

requiredstringrequired, string

A random key provided by the client, per unique wallet transaction. Maximum of 128 characters.  
The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. If a request to execute a wallet transaction fails due to a network connection error, then after a minimum delay of one minute, you can retry the request with the same idempotency key to guarantee that only a single wallet transaction is created. If the request was successfully processed, it will prevent any transaction that uses the same idempotency key, and was received within 24 hours of the first request, from being processed.  
  

Max length: `128`

Min length: `1`

[`reference`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-reference)

requiredstringrequired, string

A reference for the refund. This must be an alphanumeric string with 6 to 18 characters and must not contain any special characters or spaces.  
  

Max length: `18`

Min length: `6`

[`amount`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-amount)

objectobject

The amount and currency of a payment

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-amount-currency)

requiredstringrequired, string

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-amount-value)

requirednumberrequired, number

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`.  
  

Format: `double`

Minimum: `0.01`

[`counterparty_date_of_birth`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-counterparty-date-of-birth)

stringstring

The counterparty's birthdate, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format.  
  

Format: `date`

[`counterparty_address`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-counterparty-address)

objectobject

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-counterparty-address-street)

required[string]required, [string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-counterparty-address-city)

requiredstringrequired, string

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-counterparty-address-postal-code)

requiredstringrequired, string

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-request-counterparty-address-country)

requiredstringrequired, string

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

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

/payment\_initiation/payment/reverse

**Response fields**

[`refund_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-response-refund-id)

stringstring

A unique ID identifying the refund

[`status`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-response-status)

stringstring

The status of the transaction.  
`AUTHORISING`: The transaction is being processed for validation and compliance.  
`INITIATED`: The transaction has been initiated and is currently being processed.  
`EXECUTED`: The transaction has been successfully executed and is considered complete. This is only applicable for debit transactions.  
`SETTLED`: The transaction has settled and funds are available for use. This is only applicable for credit transactions. A transaction will typically settle within seconds to several days, depending on which payment rail is used.  
`FAILED`: The transaction failed to process successfully. This is a terminal status.  
`BLOCKED`: The transaction has been blocked for violating compliance rules. This is a terminal status.  
  

Possible values: `AUTHORISING`, `INITIATED`, `EXECUTED`, `SETTLED`, `BLOCKED`, `FAILED`

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-reverse-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "refund_id": "wallet-transaction-id-production-c5f8cd31-6cae-4cad-9b0d-f7c10be9cc4b",
  "request_id": "HtlKzBX0fMeF7mU",
  "status": "INITIATED"
}
```

=\*=\*=\*=

#### `/payment_initiation/consent/create`

#### Create payment consent

The [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate) endpoint is used to create a payment consent, which can be used to initiate payments on behalf of the user. Payment consents are created with `UNAUTHORISED` status by default and must be authorised by the user before payments can be initiated.

Consents can be limited in time and scope, and have constraints that describe limitations for payments.

/payment\_initiation/consent/create

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-recipient-id)

requiredstringrequired, string

The ID of the recipient the payment consent is for. The created consent can be used to transfer funds to this recipient only.

[`reference`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-reference)

requiredstringrequired, string

A reference for the payment consent. This must be an alphanumeric string with at most 18 characters and must not contain any special characters.  
  

Min length: `1`

Max length: `18`

[`scopes`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-scopes)

deprecated[string]deprecated, [string]

An array of payment consent scopes.  
  

Min items: `1`

Possible values: `ME_TO_ME`, `EXTERNAL`

[`type`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-type)

stringstring

Payment consent type. Defines possible use case for payments made with the given consent.  
`SWEEPING`: Allows moving money between accounts owned by the same user.  
`COMMERCIAL`: Allows initiating payments from the user's account to third parties.  
  

Possible values: `SWEEPING`, `COMMERCIAL`

[`constraints`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints)

requiredobjectrequired, object

Limitations that will be applied to payments initiated using the payment consent.

[`valid_date_time`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-valid-date-time)

objectobject

Life span for the payment consent. After the `to` date the payment consent expires and can no longer be used for payment initiation.

[`from`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-valid-date-time-from)

stringstring

The date and time from which the consent should be active, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`to`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-valid-date-time-to)

stringstring

The date and time at which the consent expires, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`max_payment_amount`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-max-payment-amount)

requiredobjectrequired, object

Maximum amount of a single payment initiated using the payment consent.

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-max-payment-amount-currency)

requiredstringrequired, string

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-max-payment-amount-value)

requirednumberrequired, number

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.  
  

Format: `double`

[`periodic_amounts`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-periodic-amounts)

required[object]required, [object]

A list of amount limitations per period of time.  
  

Min items: `1`

[`amount`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-periodic-amounts-amount)

requiredobjectrequired, object

Maximum cumulative amount for all payments in the specified interval.

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-periodic-amounts-amount-currency)

requiredstringrequired, string

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-periodic-amounts-amount-value)

requirednumberrequired, number

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.  
  

Format: `double`

[`interval`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-periodic-amounts-interval)

requiredstringrequired, string

Payment consent periodic interval.  
  

Possible values: `DAY`, `WEEK`, `MONTH`, `YEAR`

[`alignment`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-constraints-periodic-amounts-alignment)

requiredstringrequired, string

Where the payment consent period should start.  
If the institution is Monzo, only `CONSENT` alignments are supported.  
`CALENDAR`: line up with a calendar.  
`CONSENT`: on the date of consent creation.  
  

Possible values: `CALENDAR`, `CONSENT`

[`options`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-options)

deprecatedobjectdeprecated, object

(Deprecated) Additional payment consent options. Please use `payer_details` to specify the account.

[`request_refund_details`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-options-request-refund-details)

booleanboolean

When `true`, Plaid will attempt to request refund details from the payee's financial institution. Support varies between financial institutions and will not always be available. If refund details could be retrieved, they will be available in the `/payment_initiation/payment/get` response.

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-options-iban)

stringstring

The International Bank Account Number (IBAN) for the payer's account. Where possible, the end user will be able to set up payment consent using only the specified bank account if provided.  
  

Min length: `15`

Max length: `34`

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-options-bacs)

objectobject

An optional object used to restrict the accounts used for payments. If provided, the end user will be able to send payments only from the specified bank account.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-options-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-options-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`payer_details`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details)

objectobject

An object representing the payment consent payer details.
Payer `name` and account `numbers` are required to lock the account to which the consent can be created.

[`name`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-name)

requiredstringrequired, string

The name of the payer as it appears in their bank account  
  

Min length: `1`

[`numbers`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-numbers)

requiredobjectrequired, object

The counterparty's bank account numbers. Exactly one of IBAN or BACS data is required.

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-numbers-bacs)

objectobject

An optional object used to restrict the accounts used for payments. If provided, the end user will be able to send payments only from the specified bank account.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-numbers-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-numbers-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-numbers-iban)

stringstring

International Bank Account Number (IBAN).  
  

Min length: `15`

Max length: `34`

[`address`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-address)

objectobject

The optional address of the payment recipient's bank account. Required by most institutions outside of the UK.

[`street`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-address-street)

required[string]required, [string]

An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.  
  

Min items: `1`

Min length: `1`

[`city`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-address-city)

requiredstringrequired, string

The city where the recipient is located. Maximum of 35 characters.  
  

Min length: `1`

Max length: `35`

[`postal_code`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-address-postal-code)

requiredstringrequired, string

The postal code where the recipient is located. Maximum of 16 characters.  
  

Min length: `1`

Max length: `16`

[`country`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-address-country)

requiredstringrequired, string

The ISO 3166-1 alpha-2 country code where the recipient is located.  
  

Min length: `2`

Max length: `2`

[`date_of_birth`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-date-of-birth)

stringstring

The payer's birthdate, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format.  
  

Format: `date`

[`phone_numbers`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-phone-numbers)

[string][string]

The payer's phone numbers in E.164 format: +{countrycode}{number}

[`emails`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-request-payer-details-emails)

[string][string]

The payer's emails

/payment\_initiation/consent/create

```
const request: PaymentInitiationConsentCreateRequest = {
  recipient_id: recipientID,
  reference: 'TestPaymentConsent',
  type: PaymentInitiationConsentType.Commercial,
  constraints: {
    valid_date_time: {
      to: '2024-12-31T23:59:59Z',
    },
    max_payment_amount: {
      currency: PaymentAmountCurrency.Gbp,
      value: 15,
    },
    periodic_amounts: [
      {
        amount: {
          currency: PaymentAmountCurrency.Gbp,
          value: 40,
        },
        alignment: PaymentConsentPeriodicAlignment.Calendar,
        interval: PaymentConsentPeriodicInterval.Month,
      },
    ],
  },
};

try {
  const response = await plaidClient.paymentInitiationConsentCreate(request);
  const consentID = response.data.consent_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

/payment\_initiation/consent/create

**Response fields**

[`consent_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-response-consent-id)

stringstring

A unique ID identifying the payment consent.

[`status`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-response-status)

stringstring

The status of the payment consent.  
`UNAUTHORISED`: Consent created, but requires user authorisation.  
`REJECTED`: Consent authorisation was rejected by the bank.  
`AUTHORISED`: Consent is active and ready to be used.  
`REVOKED`: Consent has been revoked and can no longer be used.  
`EXPIRED`: Consent is no longer valid.  
  

Possible values: `UNAUTHORISED`, `AUTHORISED`, `REVOKED`, `REJECTED`, `EXPIRED`

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "consent_id": "consent-id-production-feca8a7a-5491-4444-9999-f3062bb735d3",
  "status": "UNAUTHORISED",
  "request_id": "4ciYmmesdqSiUAB"
}
```

=\*=\*=\*=

#### `/payment_initiation/consent/get`

#### Get payment consent

The [`/payment_initiation/consent/get`](/docs/api/products/payment-initiation/#payment_initiationconsentget) endpoint can be used to check the status of a payment consent, as well as to receive basic information such as recipient and constraints.

/payment\_initiation/consent/get

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`consent_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-request-consent-id)

requiredstringrequired, string

The `consent_id` returned from `/payment_initiation/consent/create`.

```
const request: PaymentInitiationConsentGetRequest = {
  consent_id: consentID,
};

try {
  const response = await plaidClient.paymentInitiationConsentGet(request);
  const consentID = response.data.consent_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

/payment\_initiation/consent/get

**Response fields**

[`consent_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-consent-id)

stringstring

The consent ID.  
  

Min length: `1`

[`status`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-status)

stringstring

The status of the payment consent.  
`UNAUTHORISED`: Consent created, but requires user authorisation.  
`REJECTED`: Consent authorisation was rejected by the bank.  
`AUTHORISED`: Consent is active and ready to be used.  
`REVOKED`: Consent has been revoked and can no longer be used.  
`EXPIRED`: Consent is no longer valid.  
  

Possible values: `UNAUTHORISED`, `AUTHORISED`, `REVOKED`, `REJECTED`, `EXPIRED`

[`created_at`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-created-at)

stringstring

Consent creation timestamp, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`recipient_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-recipient-id)

stringstring

The ID of the recipient the payment consent is for.  
  

Min length: `1`

[`reference`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-reference)

stringstring

A reference for the payment consent.

[`constraints`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints)

objectobject

Limitations that will be applied to payments initiated using the payment consent.

[`valid_date_time`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-valid-date-time)

nullableobjectnullable, object

Life span for the payment consent. After the `to` date the payment consent expires and can no longer be used for payment initiation.

[`from`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-valid-date-time-from)

nullablestringnullable, string

The date and time from which the consent should be active, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`to`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-valid-date-time-to)

nullablestringnullable, string

The date and time at which the consent expires, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`max_payment_amount`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-max-payment-amount)

objectobject

Maximum amount of a single payment initiated using the payment consent.

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-max-payment-amount-currency)

stringstring

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-max-payment-amount-value)

numbernumber

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.  
  

Format: `double`

[`periodic_amounts`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-periodic-amounts)

[object][object]

A list of amount limitations per period of time.  
  

Min items: `1`

[`amount`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-periodic-amounts-amount)

objectobject

Maximum cumulative amount for all payments in the specified interval.

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-periodic-amounts-amount-currency)

stringstring

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-periodic-amounts-amount-value)

numbernumber

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.  
  

Format: `double`

[`interval`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-periodic-amounts-interval)

stringstring

Payment consent periodic interval.  
  

Possible values: `DAY`, `WEEK`, `MONTH`, `YEAR`

[`alignment`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-constraints-periodic-amounts-alignment)

stringstring

Where the payment consent period should start.  
If the institution is Monzo, only `CONSENT` alignments are supported.  
`CALENDAR`: line up with a calendar.  
`CONSENT`: on the date of consent creation.  
  

Possible values: `CALENDAR`, `CONSENT`

[`scopes`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-scopes)

deprecated[string]deprecated, [string]

Deprecated, use the 'type' field instead.  
  

Possible values: `ME_TO_ME`, `EXTERNAL`

[`type`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-type)

stringstring

Payment consent type. Defines possible use case for payments made with the given consent.  
`SWEEPING`: Allows moving money between accounts owned by the same user.  
`COMMERCIAL`: Allows initiating payments from the user's account to third parties.  
  

Possible values: `SWEEPING`, `COMMERCIAL`

[`payer_details`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-payer-details)

nullableobjectnullable, object

Details about external payment refund

[`name`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-payer-details-name)

stringstring

The name of the account holder.

[`iban`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-payer-details-iban)

nullablestringnullable, string

The International Bank Account Number (IBAN) for the account.

[`bacs`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-payer-details-bacs)

nullableobjectnullable, object

An object containing a BACS account number and sort code. If an IBAN is not provided or if this recipient needs to accept domestic GBP-denominated payments, BACS data is required.

[`account`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-payer-details-bacs-account)

stringstring

The account number of the account. Maximum of 10 characters.  
  

Min length: `1`

Max length: `10`

[`sort_code`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-payer-details-bacs-sort-code)

stringstring

The 6-character sort code of the account.  
  

Min length: `6`

Max length: `6`

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "4ciYuuesdqSiUAB",
  "consent_id": "consent-id-production-feca8a7a-5491-4aef-9298-f3062bb735d3",
  "status": "AUTHORISED",
  "created_at": "2021-10-30T15:26:48Z",
  "recipient_id": "recipient-id-production-9b6b4679-914b-445b-9450-efbdb80296f6",
  "reference": "ref-00001",
  "constraints": {
    "valid_date_time": {
      "from": "2021-12-25T11:12:13Z",
      "to": "2022-12-31T15:26:48Z"
    },
    "max_payment_amount": {
      "currency": "GBP",
      "value": 100
    },
    "periodic_amounts": [
      {
        "amount": {
          "currency": "GBP",
          "value": 300
        },
        "interval": "WEEK",
        "alignment": "CALENDAR"
      }
    ]
  },
  "type": "SWEEPING"
}
```

=\*=\*=\*=

#### `/payment_initiation/consent/revoke`

#### Revoke payment consent

The [`/payment_initiation/consent/revoke`](/docs/api/products/payment-initiation/#payment_initiationconsentrevoke) endpoint can be used to revoke the payment consent. Once the consent is revoked, it is not possible to initiate payments using it.

/payment\_initiation/consent/revoke

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-revoke-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-consent-revoke-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`consent_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-revoke-request-consent-id)

requiredstringrequired, string

The consent ID.

```
const request: PaymentInitiationConsentRevokeRequest = {
  consent_id: consentID,
};
try {
  const response = await plaidClient.paymentInitiationConsentRevoke(request);
} catch (error) {
  // handle error
}
```

/payment\_initiation/consent/revoke

**Response fields**

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-revoke-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "4ciYaaesdqSiUAB"
}
```

=\*=\*=\*=

#### `/payment_initiation/consent/payment/execute`

#### Execute a single payment using consent

The [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute) endpoint can be used to execute payments using payment consent.

/payment\_initiation/consent/payment/execute

**Request fields**

[`client_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`consent_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-consent-id)

requiredstringrequired, string

The consent ID.

[`amount`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-amount)

requiredobjectrequired, object

The amount and currency of a payment

[`currency`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-amount-currency)

requiredstringrequired, string

The ISO-4217 currency code of the payment. For standing orders and payment consents, `"GBP"` must be used. For Poland, Denmark, Sweden and Norway, only the local currency is currently supported.  
  

Possible values: `GBP`, `EUR`, `PLN`, `SEK`, `DKK`, `NOK`

Min length: `3`

Max length: `3`

[`value`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-amount-value)

requirednumberrequired, number

The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.  
  

Format: `double`

[`idempotency_key`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-idempotency-key)

requiredstringrequired, string

A random key provided by the client, per unique consent payment. Maximum of 128 characters.  
The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. If a request to execute a consent payment fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single payment is created. If the request was successfully processed, it will prevent any payment that uses the same idempotency key, and was received within 48 hours of the first request, from being processed.  
  

Max length: `128`

Min length: `1`

[`reference`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-reference)

stringstring

A reference for the payment. This must be an alphanumeric string with at most 18 characters and must not contain any special characters (since not all institutions support them).
If not provided, Plaid will automatically fall back to the reference from consent. In order to track settlement via Payment Confirmation, each payment must have a unique reference. If the reference provided through the API is not unique, Plaid will adjust it.
Some institutions may limit the reference to less than 18 characters. If necessary, Plaid will adjust the reference by truncating it to fit the institution's requirements.
Both the originally provided and automatically adjusted references (if any) can be found in the `reference` and `adjusted_reference` fields, respectively.  
  

Min length: `1`

Max length: `18`

[`scope`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-scope)

deprecatedstringdeprecated, string

Deprecated, payments will be executed within the type of the consent.  
A scope of the payment. Must be one of the scopes mentioned in the consent.
Optional if the appropriate consent has only one scope defined, required otherwise.  
  

Possible values: `ME_TO_ME`, `EXTERNAL`

[`processing_mode`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-request-processing-mode)

stringstring

Decides the mode under which the payment processing should be performed, using `IMMEDIATE` as default.  
`IMMEDIATE`: Will immediately execute the payment, waiting for a response from the ASPSP before returning the result of the payment initiation. This is ideal for user-present flows.  
 `ASYNC`: Will accept a payment execution request and schedule it for processing, immediately returning the new `payment_id`. Listen for webhooks to obtain real-time updates on the payment status. This is ideal for non user-present flows.  
  

Possible values: `ASYNC`, `IMMEDIATE`

/payment\_initiation/consent/payment/execute

```
const request: PaymentInitiationConsentPaymentExecuteRequest = {
  consent_id: consentID,
  amount: {
    currency: PaymentAmountCurrency.Gbp,
    value: 7.99,
  },
  reference: 'Payment1',
  idempotency_key: idempotencyKey,
};
try {
  const response = await plaidClient.paymentInitiationConsentPaymentExecute(
    request,
  );
  const paymentID = response.data.payment_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

/payment\_initiation/consent/payment/execute

**Response fields**

[`payment_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-payment-id)

stringstring

A unique ID identifying the payment

[`status`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-status)

stringstring

The status of the payment.  
Core lifecycle statuses:  
**`PAYMENT_STATUS_INPUT_NEEDED`**: Transitional. The payment is awaiting user input to continue processing. It may re-enter this state if additional input is required.  
**`PAYMENT_STATUS_AUTHORISING`:** Transitional. The payment is being authorised by the financial institution. It will automatically move on once authorisation completes.  
**`PAYMENT_STATUS_INITIATED`:** Transitional. The payment has been authorised and accepted by the financial institution and is now in transit. A payment should be considered complete once it reaches the `PAYMENT_STATUS_EXECUTED` state or the funds settle in the recipient account.  
**`PAYMENT_STATUS_EXECUTED`: Terminal.** The funds have left the payer’s account and the payment is en route to settlement. Support is more common in the UK than in the EU; where unsupported, a successful payment remains in `PAYMENT_STATUS_INITIATED` before settling. When using Plaid Virtual Accounts, `PAYMENT_STATUS_EXECUTED` is not terminal—the payment will continue to `PAYMENT_STATUS_SETTLED` once funds are available.  
**`PAYMENT_STATUS_SETTLED`: Terminal.** The funds are available in the recipient’s account. Only available to customers using [Plaid Virtual Accounts](https://plaid.com/docs/payment-initiation/virtual-accounts/).  
Failure statuses:  
**`PAYMENT_STATUS_INSUFFICIENT_FUNDS`: Terminal.** The payment failed due to insufficient funds. No further retries will succeed until the payer’s balance is replenished.  
**`PAYMENT_STATUS_FAILED`: Terminal (retryable).** The payment could not be initiated due to a system error or outage. Retry once the root cause is resolved.  
**`PAYMENT_STATUS_BLOCKED`: Terminal (retryable).** The payment was blocked by Plaid (e.g., flagged as risky). Resolve any compliance or risk issues and retry.  
**`PAYMENT_STATUS_REJECTED`: Terminal.** The payment was rejected by the financial institution. No automatic retry is possible.  
**`PAYMENT_STATUS_CANCELLED`: Terminal.** The end user cancelled the payment during authorisation.  
Standing-order statuses:  
**`PAYMENT_STATUS_ESTABLISHED`: Terminal.** A recurring/standing order has been successfully created.  
Deprecated (to be removed in a future release):  
`PAYMENT_STATUS_UNKNOWN`: The payment status is unknown.  
`PAYMENT_STATUS_PROCESSING`: The payment is currently being processed.  
`PAYMENT_STATUS_COMPLETED`: Indicates that the standing order has been successfully established.  
  

Possible values: `PAYMENT_STATUS_INPUT_NEEDED`, `PAYMENT_STATUS_PROCESSING`, `PAYMENT_STATUS_INITIATED`, `PAYMENT_STATUS_COMPLETED`, `PAYMENT_STATUS_INSUFFICIENT_FUNDS`, `PAYMENT_STATUS_FAILED`, `PAYMENT_STATUS_BLOCKED`, `PAYMENT_STATUS_UNKNOWN`, `PAYMENT_STATUS_EXECUTED`, `PAYMENT_STATUS_SETTLED`, `PAYMENT_STATUS_AUTHORISING`, `PAYMENT_STATUS_CANCELLED`, `PAYMENT_STATUS_ESTABLISHED`, `PAYMENT_STATUS_REJECTED`

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`error`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/payment-initiation/#payment_initiation-consent-payment-execute-response-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

Response Object

```
{
  "payment_id": "payment-id-sandbox-feca8a7a-5591-4aef-9297-f3062bb735d3",
  "request_id": "4ciYccesdqSiUAB",
  "status": "PAYMENT_STATUS_INITIATED"
}
```

### Webhooks

Updates are sent to indicate that the status of an initiated payment has changed. All Payment Initiation webhooks have a `webhook_type` of `PAYMENT_INITIATION`.

=\*=\*=\*=

#### `PAYMENT_STATUS_UPDATE`

Fired when the status of a payment has changed.

**Properties**

[`webhook_type`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-webhook-type)

stringstring

`PAYMENT_INITIATION`

[`webhook_code`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-webhook-code)

stringstring

`PAYMENT_STATUS_UPDATE`

[`payment_id`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-payment-id)

stringstring

The `payment_id` for the payment being updated

[`transaction_id`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-transaction-id)

stringstring

The transaction ID that this payment is associated with, if any. This is present only when a payment was initiated using virtual accounts.

[`new_payment_status`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-new-payment-status)

stringstring

The status of the payment.  
Core lifecycle statuses:  
**`PAYMENT_STATUS_INPUT_NEEDED`**: Transitional. The payment is awaiting user input to continue processing. It may re-enter this state if additional input is required.  
**`PAYMENT_STATUS_AUTHORISING`:** Transitional. The payment is being authorised by the financial institution. It will automatically move on once authorisation completes.  
**`PAYMENT_STATUS_INITIATED`:** Transitional. The payment has been authorised and accepted by the financial institution and is now in transit. A payment should be considered complete once it reaches the `PAYMENT_STATUS_EXECUTED` state or the funds settle in the recipient account.  
**`PAYMENT_STATUS_EXECUTED`: Terminal.** The funds have left the payer’s account and the payment is en route to settlement. Support is more common in the UK than in the EU; where unsupported, a successful payment remains in `PAYMENT_STATUS_INITIATED` before settling. When using Plaid Virtual Accounts, `PAYMENT_STATUS_EXECUTED` is not terminal—the payment will continue to `PAYMENT_STATUS_SETTLED` once funds are available.  
**`PAYMENT_STATUS_SETTLED`: Terminal.** The funds are available in the recipient’s account. Only available to customers using [Plaid Virtual Accounts](https://plaid.com/docs/payment-initiation/virtual-accounts/).  
Failure statuses:  
**`PAYMENT_STATUS_INSUFFICIENT_FUNDS`: Terminal.** The payment failed due to insufficient funds. No further retries will succeed until the payer’s balance is replenished.  
**`PAYMENT_STATUS_FAILED`: Terminal (retryable).** The payment could not be initiated due to a system error or outage. Retry once the root cause is resolved.  
**`PAYMENT_STATUS_BLOCKED`: Terminal (retryable).** The payment was blocked by Plaid (e.g., flagged as risky). Resolve any compliance or risk issues and retry.  
**`PAYMENT_STATUS_REJECTED`: Terminal.** The payment was rejected by the financial institution. No automatic retry is possible.  
**`PAYMENT_STATUS_CANCELLED`: Terminal.** The end user cancelled the payment during authorisation.  
Standing-order statuses:  
**`PAYMENT_STATUS_ESTABLISHED`: Terminal.** A recurring/standing order has been successfully created.  
Deprecated (to be removed in a future release):  
`PAYMENT_STATUS_UNKNOWN`: The payment status is unknown.  
`PAYMENT_STATUS_PROCESSING`: The payment is currently being processed.  
`PAYMENT_STATUS_COMPLETED`: Indicates that the standing order has been successfully established.  
  

Possible values: `PAYMENT_STATUS_INPUT_NEEDED`, `PAYMENT_STATUS_PROCESSING`, `PAYMENT_STATUS_INITIATED`, `PAYMENT_STATUS_COMPLETED`, `PAYMENT_STATUS_INSUFFICIENT_FUNDS`, `PAYMENT_STATUS_FAILED`, `PAYMENT_STATUS_BLOCKED`, `PAYMENT_STATUS_UNKNOWN`, `PAYMENT_STATUS_EXECUTED`, `PAYMENT_STATUS_SETTLED`, `PAYMENT_STATUS_AUTHORISING`, `PAYMENT_STATUS_CANCELLED`, `PAYMENT_STATUS_ESTABLISHED`, `PAYMENT_STATUS_REJECTED`

[`old_payment_status`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-old-payment-status)

stringstring

The status of the payment.  
Core lifecycle statuses:  
**`PAYMENT_STATUS_INPUT_NEEDED`**: Transitional. The payment is awaiting user input to continue processing. It may re-enter this state if additional input is required.  
**`PAYMENT_STATUS_AUTHORISING`:** Transitional. The payment is being authorised by the financial institution. It will automatically move on once authorisation completes.  
**`PAYMENT_STATUS_INITIATED`:** Transitional. The payment has been authorised and accepted by the financial institution and is now in transit. A payment should be considered complete once it reaches the `PAYMENT_STATUS_EXECUTED` state or the funds settle in the recipient account.  
**`PAYMENT_STATUS_EXECUTED`: Terminal.** The funds have left the payer’s account and the payment is en route to settlement. Support is more common in the UK than in the EU; where unsupported, a successful payment remains in `PAYMENT_STATUS_INITIATED` before settling. When using Plaid Virtual Accounts, `PAYMENT_STATUS_EXECUTED` is not terminal—the payment will continue to `PAYMENT_STATUS_SETTLED` once funds are available.  
**`PAYMENT_STATUS_SETTLED`: Terminal.** The funds are available in the recipient’s account. Only available to customers using [Plaid Virtual Accounts](https://plaid.com/docs/payment-initiation/virtual-accounts/).  
Failure statuses:  
**`PAYMENT_STATUS_INSUFFICIENT_FUNDS`: Terminal.** The payment failed due to insufficient funds. No further retries will succeed until the payer’s balance is replenished.  
**`PAYMENT_STATUS_FAILED`: Terminal (retryable).** The payment could not be initiated due to a system error or outage. Retry once the root cause is resolved.  
**`PAYMENT_STATUS_BLOCKED`: Terminal (retryable).** The payment was blocked by Plaid (e.g., flagged as risky). Resolve any compliance or risk issues and retry.  
**`PAYMENT_STATUS_REJECTED`: Terminal.** The payment was rejected by the financial institution. No automatic retry is possible.  
**`PAYMENT_STATUS_CANCELLED`: Terminal.** The end user cancelled the payment during authorisation.  
Standing-order statuses:  
**`PAYMENT_STATUS_ESTABLISHED`: Terminal.** A recurring/standing order has been successfully created.  
Deprecated (to be removed in a future release):  
`PAYMENT_STATUS_UNKNOWN`: The payment status is unknown.  
`PAYMENT_STATUS_PROCESSING`: The payment is currently being processed.  
`PAYMENT_STATUS_COMPLETED`: Indicates that the standing order has been successfully established.  
  

Possible values: `PAYMENT_STATUS_INPUT_NEEDED`, `PAYMENT_STATUS_PROCESSING`, `PAYMENT_STATUS_INITIATED`, `PAYMENT_STATUS_COMPLETED`, `PAYMENT_STATUS_INSUFFICIENT_FUNDS`, `PAYMENT_STATUS_FAILED`, `PAYMENT_STATUS_BLOCKED`, `PAYMENT_STATUS_UNKNOWN`, `PAYMENT_STATUS_EXECUTED`, `PAYMENT_STATUS_SETTLED`, `PAYMENT_STATUS_AUTHORISING`, `PAYMENT_STATUS_CANCELLED`, `PAYMENT_STATUS_ESTABLISHED`, `PAYMENT_STATUS_REJECTED`

[`original_reference`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-original-reference)

stringstring

The original value of the reference when creating the payment.

[`adjusted_reference`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-adjusted-reference)

stringstring

The value of the reference sent to the bank after adjustment to pass bank validation rules.

[`original_start_date`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-original-start-date)

stringstring

The original value of the `start_date` provided during the creation of a standing order. If the payment is not a standing order, this field will be `null`.  
  

Format: `date`

[`adjusted_start_date`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-adjusted-start-date)

stringstring

The start date sent to the bank after adjusting for holidays or weekends. Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, or if the payment is not a standing order, this field will be `null`.  
  

Format: `date`

[`timestamp`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-timestamp)

stringstring

The timestamp of the update, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2017-09-14T14:42:19.350Z"`  
  

Format: `date-time`

[`error`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error)

objectobject

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-error-code-reason)

stringstring

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-display-message)

stringstring

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-status)

integerinteger

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-suggested-action)

stringstring

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`environment`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "PAYMENT_INITIATION",
  "webhook_code": "PAYMENT_STATUS_UPDATE",
  "payment_id": "payment-id-production-2ba30780-d549-4335-b1fe-c2a938aa39d2",
  "new_payment_status": "PAYMENT_STATUS_INITIATED",
  "old_payment_status": "PAYMENT_STATUS_PROCESSING",
  "original_reference": "Account Funding 99744",
  "adjusted_reference": "Account Funding 99",
  "original_start_date": "2017-09-14",
  "adjusted_start_date": "2017-09-15",
  "timestamp": "2017-09-14T14:42:19.350Z",
  "environment": "production"
}
```

=\*=\*=\*=

#### `CONSENT_STATUS_UPDATE`

Fired when the status of a payment consent has changed.

**Properties**

[`webhook_type`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-webhook-type)

stringstring

`PAYMENT_INITIATION`

[`webhook_code`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-webhook-code)

stringstring

`CONSENT_STATUS_UPDATE`

[`consent_id`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-consent-id)

stringstring

The `id` for the consent being updated

[`old_status`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-old-status)

stringstring

The status of the payment consent.  
`UNAUTHORISED`: Consent created, but requires user authorisation.  
`REJECTED`: Consent authorisation was rejected by the bank.  
`AUTHORISED`: Consent is active and ready to be used.  
`REVOKED`: Consent has been revoked and can no longer be used.  
`EXPIRED`: Consent is no longer valid.  
  

Possible values: `UNAUTHORISED`, `AUTHORISED`, `REVOKED`, `REJECTED`, `EXPIRED`

[`new_status`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-new-status)

stringstring

The status of the payment consent.  
`UNAUTHORISED`: Consent created, but requires user authorisation.  
`REJECTED`: Consent authorisation was rejected by the bank.  
`AUTHORISED`: Consent is active and ready to be used.  
`REVOKED`: Consent has been revoked and can no longer be used.  
`EXPIRED`: Consent is no longer valid.  
  

Possible values: `UNAUTHORISED`, `AUTHORISED`, `REVOKED`, `REJECTED`, `EXPIRED`

[`timestamp`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-timestamp)

stringstring

The timestamp of the update, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2017-09-14T14:42:19.350Z"`  
  

Format: `date-time`

[`error`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error)

objectobject

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-error-code-reason)

stringstring

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-display-message)

stringstring

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-status)

integerinteger

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-suggested-action)

stringstring

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`environment`](/docs/api/products/payment-initiation/#PaymentInitiationConsentStatusUpdateWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "PAYMENT_INITIATION",
  "webhook_code": "CONSENT_STATUS_UPDATE",
  "consent_id": "payment-consent-id-production-e7258765-69f9-46b1-9c67-d2800448e5ff",
  "old_status": "UNAUTHORISED",
  "new_status": "AUTHORISED",
  "timestamp": "2017-09-14T14:42:19.350Z",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
