---
title: "API - Account Linking | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/transfer/account-linking/"
scraped_at: "2026-03-07T22:04:21+00:00"
---

# Transfer

#### API reference for Transfer account linking endpoints

For how-to guidance, see the [Transfer documentation](/docs/transfer/).

=\*=\*=\*=

#### Account Linking

| Account Linking |  |
| --- | --- |
| [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget) | Determine RTP eligibility for a Plaid Item |
| [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) | Create a transfer intent and invoke Transfer UI (Transfer UI only) |
| [`/transfer/intent/get`](/docs/api/products/transfer/account-linking/#transferintentget) | Retrieve information about a transfer intent (Transfer UI only) |
| [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) | Create an Item to use with Transfer from known account and routing numbers |

=\*=\*=\*=

#### `/transfer/capabilities/get`

#### Get RTP eligibility information of a transfer

Use the [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget) endpoint to determine the RTP eligibility information of an account to be used with Transfer. This endpoint works on all Transfer-capable Items, including those created by [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account).

/transfer/capabilities/get

**Request fields**

[`client_id`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`access_token`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-request-access-token)

requiredstringrequired, string

The Plaid `access_token` for the account that will be debited or credited.

[`account_id`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-request-account-id)

requiredstringrequired, string

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

/transfer/capabilities/get

```
const request: TransferCapabilitiesGetRequest = {
  access_token: 'access-sandbox-71e02f71-0960-4a27-abd2-5631e04f2175',
  account_id: '3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr',
};

try {
  const response = await client.transferCapabilitiesGet(request);
} catch (error) {
  // handle error
}
```

/transfer/capabilities/get

**Response fields**

[`institution_supported_networks`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-response-institution-supported-networks)

objectobject

Contains the RTP and RfP network and types supported by the linked Item's institution.

[`rtp`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-response-institution-supported-networks-rtp)

objectobject

Contains the supported service types in RTP

[`credit`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-response-institution-supported-networks-rtp-credit)

booleanboolean

When `true`, the linked Item's institution supports RTP credit transfer.  
  

Default: `false`

[`rfp`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-response-institution-supported-networks-rfp)

objectobject

Contains the supported service types in RfP

[`debit`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-response-institution-supported-networks-rfp-debit)

booleanboolean

When `true`, the linked Item's institution supports RfP debit transfer.  
  

Default: `false`

[`max_amount`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-response-institution-supported-networks-rfp-max-amount)

nullablestringnullable, string

The maximum amount (decimal string with two digits of precision e.g. "10.00") for originating RfP transfers with the given institution.

[`iso_currency_code`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-response-institution-supported-networks-rfp-iso-currency-code)

nullablestringnullable, string

The currency of the `max_amount`, e.g. "USD".

[`request_id`](/docs/api/products/transfer/account-linking/#transfer-capabilities-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "institution_supported_networks": {
    "rtp": {
      "credit": true
    },
    "rfp": {
      "debit": true
    }
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/intent/create`

#### Create a transfer intent object to invoke the Transfer UI

Use the [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) endpoint to generate a transfer intent object and invoke the Transfer UI.

/transfer/intent/create

**Request fields**

[`client_id`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`account_id`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-account-id)

stringstring

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`mode`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-mode)

requiredstringrequired, string

The direction of the flow of transfer funds.  
`PAYMENT`: Transfers funds from an end user's account to your business account.  
`DISBURSEMENT`: Transfers funds from your business account to an end user's account.  
  

Possible values: `PAYMENT`, `DISBURSEMENT`

[`network`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-network)

stringstring

The network or rails used for the transfer. Defaults to `same-day-ach`.  
For transfers submitted using `ach`, the Standard ACH cutoff is 8:30 PM Eastern Time.  
For transfers submitted using `same-day-ach`, the Same Day ACH cutoff is 3:00 PM Eastern Time. It is recommended to send the request 15 minutes prior to the cutoff to ensure that it will be processed in time for submission before the cutoff. If the transfer is processed after this cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur same-day charges.  
For transfers submitted using `rtp`, in the case that the account being credited does not support RTP, the transfer will be sent over ACH as long as an `ach_class` is provided in the request. If RTP isn't supported by the account and no `ach_class` is provided, the transfer will fail to be submitted.  
  

Possible values: `ach`, `same-day-ach`, `rtp`

Default: `same-day-ach`

[`amount`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-amount)

requiredstringrequired, string

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`description`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-description)

requiredstringrequired, string

A description for the underlying transfer. Maximum of 15 characters.  
  

Min length: `1`

Max length: `15`

[`ach_class`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`user`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user)

requiredobjectrequired, object

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-legal-name)

requiredstringrequired, string

The user's legal name.

[`phone_number`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-phone-number)

stringstring

The user's phone number. Phone number input may be validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`email_address`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-email-address)

stringstring

The user's email address.

[`address`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-address)

objectobject

The address associated with the account holder.

[`street`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-address-street)

stringstring

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-address-city)

stringstring

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-address-region)

stringstring

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-address-postal-code)

stringstring

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-user-address-country)

stringstring

A two-letter country code (e.g., "US").

[`metadata`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-metadata)

objectobject

The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply:
The JSON values must be Strings (no nested JSON objects allowed)
Only ASCII characters may be used
Maximum of 50 key/value pairs
Maximum key length of 40 characters
Maximum value length of 500 characters

[`iso_currency_code`](/docs/api/products/transfer/account-linking/#transfer-intent-create-request-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

/transfer/intent/create

```
const request: TransferIntentCreateRequest = {
  account_id: '3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr',
  mode: 'PAYMENT',
  amount: '12.34',
  description: 'Desc',
  ach_class: 'ppd',
  origination_account_id: '9853defc-e703-463d-86b1-dc0607a45359',
  user: {
    legal_name: 'Anne Charleston',
  },
};

try {
  const response = await client.transferIntentCreate(request);
} catch (error) {
  // handle error
}
```

/transfer/intent/create

**Response fields**

[`transfer_intent`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent)

objectobject

Represents a transfer intent within Transfer UI.

[`id`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-id)

stringstring

Plaid's unique identifier for the transfer intent object.

[`created`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-created)

stringstring

The datetime the transfer was created. This will be of the form `2006-01-02T15:04:05Z`.  
  

Format: `date-time`

[`status`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-status)

stringstring

The status of the transfer intent.  
`PENDING`: The transfer intent is pending.
`SUCCEEDED`: The transfer intent was successfully created.
`FAILED`: The transfer intent was unable to be created.  
  

Possible values: `PENDING`, `SUCCEEDED`, `FAILED`

[`account_id`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-account-id)

nullablestringnullable, string

The Plaid `account_id` corresponding to the end-user account that will be debited or credited. Returned only if `account_id` was set on intent creation.

[`funding_account_id`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`amount`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`mode`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-mode)

stringstring

The direction of the flow of transfer funds.  
`PAYMENT`: Transfers funds from an end user's account to your business account.  
`DISBURSEMENT`: Transfers funds from your business account to an end user's account.  
  

Possible values: `PAYMENT`, `DISBURSEMENT`

[`network`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-network)

stringstring

The network or rails used for the transfer. Defaults to `same-day-ach`.  
For transfers submitted using `ach`, the Standard ACH cutoff is 8:30 PM Eastern Time.  
For transfers submitted using `same-day-ach`, the Same Day ACH cutoff is 3:00 PM Eastern Time. It is recommended to send the request 15 minutes prior to the cutoff to ensure that it will be processed in time for submission before the cutoff. If the transfer is processed after this cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur same-day charges.  
For transfers submitted using `rtp`, in the case that the account being credited does not support RTP, the transfer will be sent over ACH as long as an `ach_class` is provided in the request. If RTP isn't supported by the account and no `ach_class` is provided, the transfer will fail to be submitted.  
  

Possible values: `ach`, `same-day-ach`, `rtp`

Default: `same-day-ach`

[`ach_class`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`user`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`description`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-description)

stringstring

A description for the underlying transfer. Maximum of 8 characters.

[`metadata`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-metadata)

nullableobjectnullable, object

The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply:
The JSON values must be Strings (no nested JSON objects allowed)
Only ASCII characters may be used
Maximum of 50 key/value pairs
Maximum key length of 40 characters
Maximum value length of 500 characters

[`iso_currency_code`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-transfer-intent-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

[`request_id`](/docs/api/products/transfer/account-linking/#transfer-intent-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transfer_intent": {
    "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
    "funding_account_id": "9853defc-e703-463d-86b1-dc0607a45359",
    "ach_class": "ppd",
    "amount": "12.34",
    "iso_currency_code": "USD",
    "created": "2020-08-06T17:27:15Z",
    "description": "Desc",
    "id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
    "metadata": {
      "key1": "value1",
      "key2": "value2"
    },
    "mode": "PAYMENT",
    "origination_account_id": "9853defc-e703-463d-86b1-dc0607a45359",
    "status": "PENDING",
    "user": {
      "address": {
        "street": "100 Market Street",
        "city": "San Francisco",
        "region": "CA",
        "postal_code": "94103",
        "country": "US"
      },
      "email_address": "acharleston@email.com",
      "legal_name": "Anne Charleston",
      "phone_number": "123-456-7890"
    }
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/intent/get`

#### Retrieve more information about a transfer intent

Use the [`/transfer/intent/get`](/docs/api/products/transfer/account-linking/#transferintentget) endpoint to retrieve more information about a transfer intent.

/transfer/intent/get

**Request fields**

[`client_id`](/docs/api/products/transfer/account-linking/#transfer-intent-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/account-linking/#transfer-intent-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`transfer_intent_id`](/docs/api/products/transfer/account-linking/#transfer-intent-get-request-transfer-intent-id)

requiredstringrequired, string

Plaid's unique identifier for a transfer intent object.

/transfer/intent/get

```
const request: TransferIntentGetRequest = {
  transfer_intent_id: '460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9',
};

try {
  const response = await client.transferIntentGet(request);
} catch (error) {
  // handle error
}
```

/transfer/intent/get

**Response fields**

[`transfer_intent`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent)

objectobject

Represents a transfer intent within Transfer UI.

[`id`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-id)

stringstring

Plaid's unique identifier for a transfer intent object.

[`created`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-created)

stringstring

The datetime the transfer was created. This will be of the form `2006-01-02T15:04:05Z`.  
  

Format: `date-time`

[`status`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-status)

stringstring

The status of the transfer intent.  
`PENDING`: The transfer intent is pending.
`SUCCEEDED`: The transfer intent was successfully created.
`FAILED`: The transfer intent was unable to be created.  
  

Possible values: `PENDING`, `SUCCEEDED`, `FAILED`

[`transfer_id`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-transfer-id)

nullablestringnullable, string

Plaid's unique identifier for the transfer created through the UI. Returned only if the transfer was successfully created. Null value otherwise.

[`failure_reason`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-failure-reason)

nullableobjectnullable, object

The reason for a failed transfer intent. Returned only if the transfer intent status is `failed`. Null otherwise.

[`error_type`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-failure-reason-error-type)

stringstring

A broad categorization of the error.

[`error_code`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-failure-reason-error-code)

stringstring

A code representing the reason for a failed transfer intent (i.e., an API error or the authorization being declined).

[`error_message`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-failure-reason-error-message)

stringstring

A human-readable description of the code associated with a failed transfer intent.

[`authorization_decision`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-authorization-decision)

nullablestringnullable, string

A decision regarding the proposed transfer.  
`APPROVED` – The proposed transfer has received the end user's consent and has been approved for processing by Plaid. The `decision_rationale` field is set if Plaid was unable to fetch the account information. You may proceed with the transfer, but further review is recommended (i.e., use Link in update mode to re-authenticate your user when `decision_rationale.code` is `ITEM_LOGIN_REQUIRED`). Refer to the `code` field in the `decision_rationale` object for details.  
`DECLINED` – Plaid reviewed the proposed transfer and declined processing. Refer to the `code` field in the `decision_rationale` object for details.  
  

Possible values: `APPROVED`, `DECLINED`

[`authorization_decision_rationale`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-authorization-decision-rationale)

nullableobjectnullable, object

The rationale for Plaid's decision regarding a proposed transfer. It is always set for `declined` decisions, and may or may not be null for `approved` decisions.

[`code`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-authorization-decision-rationale-code)

stringstring

A code representing the rationale for approving or declining the proposed transfer.  
If the `rationale_code` is `null`, the transfer passed the authorization check.  
Any non-`null` value for an `approved` transfer indicates that the the authorization check could not be run and that you should perform your own risk assessment on the transfer. The code will indicate why the check could not be run. Possible values for an `approved` transfer are:  
`MANUALLY_VERIFIED_ITEM` – Item created via a manual entry flow (i.e. Same Day Micro-deposit, Instant Micro-deposit, or database-based verification), limited information available.  
`ITEM_LOGIN_REQUIRED` – Unable to collect the account information due to Item staleness. Can be resolved by using Link and setting [`transfer.authorization_id`](https://plaid.com/docs/api/link/#link-token-create-request-transfer-authorization-id) in the request to `/link/token/create`.  
`MIGRATED_ACCOUNT_ITEM` - Item created via `/transfer/migrate_account` endpoint, limited information available.  
`ERROR` – Unable to collect the account information due to an unspecified error.  
The following codes indicate that the authorization decision was `declined`:  
`NSF` – Transaction likely to result in a return due to insufficient funds.  
`RISK` - Transaction is high-risk.  
`TRANSFER_LIMIT_REACHED` - One or several transfer limits are reached, e.g. monthly transfer limit. Check the accompanying `description` field to understand which limit has been reached.  
  

Possible values: `NSF`, `RISK`, `TRANSFER_LIMIT_REACHED`, `MANUALLY_VERIFIED_ITEM`, `ITEM_LOGIN_REQUIRED`, `PAYMENT_PROFILE_LOGIN_REQUIRED`, `ERROR`, `MIGRATED_ACCOUNT_ITEM`, `null`

[`description`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-authorization-decision-rationale-description)

stringstring

A human-readable description of the code associated with a transfer approval or transfer decline.

[`account_id`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-account-id)

nullablestringnullable, string

The Plaid `account_id` for the account that will be debited or credited. Returned only if `account_id` was set on intent creation.

[`funding_account_id`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`amount`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`mode`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-mode)

stringstring

The direction of the flow of transfer funds.  
`PAYMENT`: Transfers funds from an end user's account to your business account.  
`DISBURSEMENT`: Transfers funds from your business account to an end user's account.  
  

Possible values: `PAYMENT`, `DISBURSEMENT`

[`network`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-network)

stringstring

The network or rails used for the transfer. Defaults to `same-day-ach`.  
For transfers submitted using `ach`, the Standard ACH cutoff is 8:30 PM Eastern Time.  
For transfers submitted using `same-day-ach`, the Same Day ACH cutoff is 3:00 PM Eastern Time. It is recommended to send the request 15 minutes prior to the cutoff to ensure that it will be processed in time for submission before the cutoff. If the transfer is processed after this cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur same-day charges.  
For transfers submitted using `rtp`, in the case that the account being credited does not support RTP, the transfer will be sent over ACH as long as an `ach_class` is provided in the request. If RTP isn't supported by the account and no `ach_class` is provided, the transfer will fail to be submitted.  
  

Possible values: `ach`, `same-day-ach`, `rtp`

Default: `same-day-ach`

[`ach_class`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`user`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`description`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-description)

stringstring

A description for the underlying transfer. Maximum of 8 characters.

[`metadata`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-metadata)

nullableobjectnullable, object

The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply:
The JSON values must be Strings (no nested JSON objects allowed)
Only ASCII characters may be used
Maximum of 50 key/value pairs
Maximum key length of 40 characters
Maximum value length of 500 characters

[`iso_currency_code`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-transfer-intent-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

[`request_id`](/docs/api/products/transfer/account-linking/#transfer-intent-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transfer_intent": {
    "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
    "funding_account_id": "9853defc-e703-463d-86b1-dc0607a45359",
    "ach_class": "ppd",
    "amount": "12.34",
    "iso_currency_code": "USD",
    "authorization_decision": "APPROVED",
    "authorization_decision_rationale": null,
    "created": "2019-12-09T17:27:15Z",
    "description": "Desc",
    "failure_reason": null,
    "guarantee_decision": null,
    "guarantee_decision_rationale": null,
    "id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
    "metadata": {
      "key1": "value1",
      "key2": "value2"
    },
    "mode": "DISBURSEMENT",
    "origination_account_id": "9853defc-e703-463d-86b1-dc0607a45359",
    "status": "SUCCEEDED",
    "transfer_id": "590ecd12-1dcc-7eae-4ad6-c28d1ec90df2",
    "user": {
      "address": {
        "street": "123 Main St.",
        "city": "San Francisco",
        "region": "California",
        "postal_code": "94053",
        "country": "US"
      },
      "email_address": "acharleston@email.com",
      "legal_name": "Anne Charleston",
      "phone_number": "510-555-0128"
    }
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/migrate_account`

#### Migrate account into Transfers

As an alternative to adding Items via Link, you can also use the [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) endpoint to migrate previously-verified account and routing numbers to Plaid Items. This endpoint is also required when adding an Item for use with wire transfers; if you intend to create wire transfers on this account, you must provide `wire_routing_number`. Note that Items created in this way are not compatible with endpoints for other products, such as [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget), and can only be used with Transfer endpoints. If you require access to other endpoints, create the Item through Link instead. Access to [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) is not enabled by default; to obtain access, contact your Plaid Account Manager or [Support](https://dashboard.plaid.com/support).

/transfer/migrate\_account

**Request fields**

[`client_id`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`account_number`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-request-account-number)

requiredstringrequired, string

The user's account number.

[`routing_number`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-request-routing-number)

requiredstringrequired, string

The user's routing number.

[`wire_routing_number`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-request-wire-routing-number)

stringstring

The user's wire transfer routing number. This is the ABA number; for some institutions, this may differ from the ACH number used in `routing_number`. This field must be set for the created item to be eligible for wire transfers.

[`account_type`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-request-account-type)

requiredstringrequired, string

The type of the bank account (`checking` or `savings`).

/transfer/migrate\_account

```
const request: TransferMigrateAccountRequest = {
  account_number: '100000000',
  routing_number: '121122676',
  account_type: 'checking',
};
try {
  const response = await plaidClient.transferMigrateAccount(request);
  const access_token = response.data.access_token;
} catch (error) {
  // handle error
}
```

/transfer/migrate\_account

**Response fields**

[`access_token`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-response-access-token)

stringstring

The Plaid `access_token` for the newly created Item.

[`account_id`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-response-account-id)

stringstring

The Plaid `account_id` for the newly created Item.

[`request_id`](/docs/api/products/transfer/account-linking/#transfer-migrate_account-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "access_token": "access-sandbox-435beced-94e8-4df3-a181-1dde1cfa19f0",
  "account_id": "zvyDgbeeDluZ43AJP6m5fAxDlgoZXDuoy5gjN",
  "request_id": "mdqfuVxeoza6mhu"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
