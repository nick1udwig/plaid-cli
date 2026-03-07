---
title: "API - Recurring Transfers | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/transfer/recurring-transfers/"
scraped_at: "2026-03-07T22:04:25+00:00"
---

# Recurring transfers

#### API reference for recurring transfer endpoints and webhooks

=\*=\*=\*=

#### Recurring transfers

For how-to guidance, see the [recurring transfers documentation](/docs/transfer/recurring-transfers/).

| Recurring Transfer endpoints |  |
| --- | --- |
| [`/transfer/recurring/create`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcreate) | Create a recurring transfer |
| [`/transfer/recurring/cancel`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcancel) | Cancel a recurring transfer |
| [`/transfer/recurring/get`](/docs/api/products/transfer/recurring-transfers/#transferrecurringget) | Retrieve information about a recurring transfer |
| [`/transfer/recurring/list`](/docs/api/products/transfer/recurring-transfers/#transferrecurringlist) | Retrieve a list of recurring transfers |

| Webhooks |  |
| --- | --- |
| [`RECURRING_CANCELLED`](/docs/api/products/transfer/recurring-transfers/#recurring_new_transfer) | A recurring transfer has been cancelled by Plaid |
| [`RECURRING_NEW_TRANSFER`](/docs/api/products/transfer/recurring-transfers/#recurring_new_transfer) | A new transfer of a recurring transfer has been originated |
| [`RECURRING_TRANSFER_SKIPPED`](/docs/api/products/transfer/recurring-transfers/#recurring_transfer_skipped) | An instance of a scheduled recurring transfer could not be created |

### Endpoints

=\*=\*=\*=

#### `/transfer/recurring/create`

#### Create a recurring transfer

Use the [`/transfer/recurring/create`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcreate) endpoint to initiate a new recurring transfer. This capability is not currently supported for Transfer UI or Transfer for Platforms (beta) customers.

/transfer/recurring/create

**Request fields**

[`client_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`access_token`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-access-token)

requiredstringrequired, string

The Plaid `access_token` for the account that will be debited or credited.

[`idempotency_key`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-idempotency-key)

requiredstringrequired, string

A random key provided by the client, per unique recurring transfer. Maximum of 50 characters.  
The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create a recurring fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single recurring transfer is created.  
  

Max length: `50`

[`account_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-account-id)

requiredstringrequired, string

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`type`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-type)

requiredstringrequired, string

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`network`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-network)

requiredstringrequired, string

Networks eligible for recurring transfers.  
  

Possible values: `ach`, `same-day-ach`, `rtp`

[`ach_class`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`amount`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-amount)

requiredstringrequired, string

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`user_present`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-present)

booleanboolean

If the end user is initiating the specific transfer themselves via an interactive UI, this should be `true`; for automatic recurring payments where the end user is not actually initiating each individual transfer, it should be `false`.

[`description`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-description)

requiredstringrequired, string

The description of the recurring transfer.

[`test_clock_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-test-clock-id)

stringstring

Plaid’s unique identifier for a test clock. This field may only be used when using the `sandbox` environment. If provided, the created `recurring_transfer` is associated with the `test_clock`. New originations are automatically generated when the associated `test_clock` advances. For more details, see [Simulating recurring transfers](https://plaid.com/docs/transfer/sandbox/#simulating-recurring-transfers).

[`schedule`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-schedule)

requiredobjectrequired, object

The schedule that the recurring transfer will be executed on.

[`interval_unit`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-schedule-interval-unit)

requiredstringrequired, string

The unit of the recurring interval.  
  

Possible values: `week`, `month`

Min length: `1`

[`interval_count`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-schedule-interval-count)

requiredintegerrequired, integer

The number of recurring `interval_units` between originations. The recurring interval (before holiday adjustment) is calculated by multiplying `interval_unit` and `interval_count`.
For example, to schedule a recurring transfer which originates once every two weeks, set `interval_unit` = `week` and `interval_count` = 2.

[`interval_execution_day`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-schedule-interval-execution-day)

requiredintegerrequired, integer

The day of the interval on which to schedule the transfer.  
If the `interval_unit` is `week`, `interval_execution_day` should be an integer from 1 (Monday) to 5 (Friday).  
If the `interval_unit` is `month`, `interval_execution_day` should be an integer indicating which day of the month to make the transfer on. Integers from 1 to 28 can be used to make a transfer on that day of the month. Negative integers from -1 to -5 can be used to make a transfer relative to the end of the month. To make a transfer on the last day of the month, use -1; to make the transfer on the second-to-last day, use -2, and so on.  
The transfer will be originated on the next available banking day if the designated day is a non banking day.

[`start_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-schedule-start-date)

requiredstringrequired, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). The recurring transfer will begin on the first `interval_execution_day` on or after the `start_date`.  
For `rtp` recurring transfers, `start_date` must be in the future.
Otherwise, if the first `interval_execution_day` on or after the start date is also the same day that `/transfer/recurring/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`end_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-schedule-end-date)

stringstring

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). The recurring transfer will end on the last `interval_execution_day` on or before the `end_date`.
If the `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/transfer/recurring/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`user`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user)

requiredobjectrequired, object

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-legal-name)

requiredstringrequired, string

The user's legal name.

[`phone_number`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-phone-number)

stringstring

The user's phone number. Phone number input may be validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`email_address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-email-address)

stringstring

The user's email address.

[`address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-address)

objectobject

The address associated with the account holder.

[`street`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-address-street)

stringstring

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-address-city)

stringstring

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-address-region)

stringstring

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-address-postal-code)

stringstring

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-user-address-country)

stringstring

A two-letter country code (e.g., "US").

[`device`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-device)

objectobject

Information about the device being used to initiate the authorization.

[`ip_address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-device-ip-address)

requiredstringrequired, string

The IP address of the device being used to initiate the authorization.

[`user_agent`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-request-device-user-agent)

requiredstringrequired, string

The user agent of the device being used to initiate the authorization.

/transfer/recurring/create

```
const request: TransferRecurringCreateRequest = {
  access_token: 'access-sandbox-71e02f71-0960-4a27-abd2-5631e04f2175',
  account_id: '3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr',
  type: 'credit',
  network: 'ach',
  amount: '12.34',
  ach_class: 'ppd',
  description: 'payment',
  idempotency_key: '12345',
  schedule: {
    start_date: '2022-10-01',
    end_date: '2023-10-01',
    interval_unit: 'week',
    interval_count: 1,
    interval_execution_day: 5
  },
  user: {
    legal_name: 'Anne Charleston',
  },
};

try {
  const response = await client.transferRecurringCreate(request);
  const recurringTransferId = response.data.recurring_transfer.recurring_transfer_id;
} catch (error) {
  // handle error
}
```

/transfer/recurring/create

**Response fields**

[`recurring_transfer`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer)

nullableobjectnullable, object

Represents a recurring transfer within the Transfers API.

[`recurring_transfer_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-recurring-transfer-id)

stringstring

Plaid’s unique identifier for a recurring transfer.

[`created`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-created)

stringstring

The datetime when this transfer was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`next_origination_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-next-origination-date)

nullablestringnullable, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
The next transfer origination date after bank holiday adjustment.  
  

Format: `date`

[`test_clock_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-test-clock-id)

nullablestringnullable, string

Plaid’s unique identifier for a test clock.

[`type`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-type)

stringstring

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`amount`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`status`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-status)

stringstring

The status of the recurring transfer.  
`active`: The recurring transfer is currently active.
`cancelled`: The recurring transfer was cancelled by the client or Plaid.
`expired`: The recurring transfer has completed all originations according to its recurring schedule.  
  

Possible values: `active`, `cancelled`, `expired`

[`ach_class`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`network`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-network)

stringstring

Networks eligible for recurring transfers.  
  

Possible values: `ach`, `same-day-ach`, `rtp`

[`account_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-account-id)

stringstring

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`funding_account_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`iso_currency_code`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

[`description`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-description)

stringstring

The description of the recurring transfer.

[`transfer_ids`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-transfer-ids)

[string][string]

The created transfer instances associated with this `recurring_transfer_id`. If the recurring transfer has been newly created, this array will be empty.

[`user`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`schedule`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-schedule)

objectobject

The schedule that the recurring transfer will be executed on.

[`interval_unit`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-schedule-interval-unit)

stringstring

The unit of the recurring interval.  
  

Possible values: `week`, `month`

Min length: `1`

[`interval_count`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-schedule-interval-count)

integerinteger

The number of recurring `interval_units` between originations. The recurring interval (before holiday adjustment) is calculated by multiplying `interval_unit` and `interval_count`.
For example, to schedule a recurring transfer which originates once every two weeks, set `interval_unit` = `week` and `interval_count` = 2.

[`interval_execution_day`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-schedule-interval-execution-day)

integerinteger

The day of the interval on which to schedule the transfer.  
If the `interval_unit` is `week`, `interval_execution_day` should be an integer from 1 (Monday) to 5 (Friday).  
If the `interval_unit` is `month`, `interval_execution_day` should be an integer indicating which day of the month to make the transfer on. Integers from 1 to 28 can be used to make a transfer on that day of the month. Negative integers from -1 to -5 can be used to make a transfer relative to the end of the month. To make a transfer on the last day of the month, use -1; to make the transfer on the second-to-last day, use -2, and so on.  
The transfer will be originated on the next available banking day if the designated day is a non banking day.

[`start_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-schedule-start-date)

stringstring

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). The recurring transfer will begin on the first `interval_execution_day` on or after the `start_date`.  
For `rtp` recurring transfers, `start_date` must be in the future.
Otherwise, if the first `interval_execution_day` on or after the start date is also the same day that `/transfer/recurring/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`end_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-recurring-transfer-schedule-end-date)

nullablestringnullable, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). The recurring transfer will end on the last `interval_execution_day` on or before the `end_date`.
If the `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/transfer/recurring/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`decision`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-decision)

stringstring

A decision regarding the proposed transfer.  
`approved` – The proposed transfer has received the end user's consent and has been approved for processing by Plaid. The `decision_rationale` field is set if Plaid was unable to fetch the account information. You may proceed with the transfer, but further review is recommended. Refer to the `code` field in the `decision_rationale` object for details.  
`declined` – Plaid reviewed the proposed transfer and declined processing. Refer to the `code` field in the `decision_rationale` object for details.  
`user_action_required` – An action is required before Plaid can assess the transfer risk and make a decision. The most common scenario is to update authentication for an Item. To complete the required action, initialize Link by setting `transfer.authorization_id` in the request of `/link/token/create`. After Link flow is completed, you may re-attempt the authorization request.  
For `guarantee` requests, `approved` indicates the transfer is eligible for Plaid's guarantee, and `declined` indicates Plaid will not provide guarantee coverage for the transfer. `user_action_required` indicates you should follow the above guidance before re-attempting.  
  

Possible values: `approved`, `declined`, `user_action_required`

[`decision_rationale`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-decision-rationale)

nullableobjectnullable, object

The rationale for Plaid's decision regarding a proposed transfer. It is always set for `declined` decisions, and may or may not be null for `approved` decisions.

[`code`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-decision-rationale-code)

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

[`description`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-decision-rationale-description)

stringstring

A human-readable description of the code associated with a transfer approval or transfer decline.

[`request_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "recurring_transfer": {
    "recurring_transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
    "created": "2022-07-05T12:48:37Z",
    "next_origination_date": "2022-10-28",
    "test_clock_id": "b33a6eda-5e97-5d64-244a-a9274110151c",
    "status": "active",
    "amount": "12.34",
    "description": "payment",
    "type": "debit",
    "ach_class": "ppd",
    "network": "ach",
    "origination_account_id": "",
    "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
    "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
    "iso_currency_code": "USD",
    "transfer_ids": [],
    "user": {
      "legal_name": "Anne Charleston",
      "phone_number": "510-555-0128",
      "email_address": "acharleston@email.com",
      "address": {
        "street": "123 Main St.",
        "city": "San Francisco",
        "region": "CA",
        "postal_code": "94053",
        "country": "US"
      }
    },
    "schedule": {
      "start_date": "2022-10-01",
      "end_date": "2023-10-01",
      "interval_unit": "week",
      "interval_count": 1,
      "interval_execution_day": 5
    }
  },
  "decision": "approved",
  "decision_rationale": null,
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/recurring/cancel`

#### Cancel a recurring transfer.

Use the [`/transfer/recurring/cancel`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcancel) endpoint to cancel a recurring transfer. Scheduled transfer that hasn't been submitted to bank will be cancelled.

/transfer/recurring/cancel

**Request fields**

[`client_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-cancel-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-cancel-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`recurring_transfer_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-cancel-request-recurring-transfer-id)

requiredstringrequired, string

Plaid’s unique identifier for a recurring transfer.

/transfer/recurring/cancel

```
const request: TransferRecurringCancelRequest = {
  recurring_transfer_id: '460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9',
};

try {
  const response = await client.transferRecurringCancel(request);
} catch (error) {
  // handle error
}
```

/transfer/recurring/cancel

**Response fields**

[`request_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-cancel-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/recurring/get`

#### Retrieve a recurring transfer

The [`/transfer/recurring/get`](/docs/api/products/transfer/recurring-transfers/#transferrecurringget) fetches information about the recurring transfer corresponding to the given `recurring_transfer_id`.

/transfer/recurring/get

**Request fields**

[`client_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`recurring_transfer_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-request-recurring-transfer-id)

requiredstringrequired, string

Plaid’s unique identifier for a recurring transfer.

/transfer/recurring/get

```
const request: TransferRecurringGetRequest = {
  recurring_transfer_id: '460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9',
};

try {
  const response = await client.transferRecurringGet(request);
  const recurringTransferId =
    response.data.recurring_transfer.recurring_transfer_id;
} catch (error) {
  // handle error
}
```

/transfer/recurring/get

**Response fields**

[`recurring_transfer`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer)

objectobject

Represents a recurring transfer within the Transfers API.

[`recurring_transfer_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-recurring-transfer-id)

stringstring

Plaid’s unique identifier for a recurring transfer.

[`created`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-created)

stringstring

The datetime when this transfer was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`next_origination_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-next-origination-date)

nullablestringnullable, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
The next transfer origination date after bank holiday adjustment.  
  

Format: `date`

[`test_clock_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-test-clock-id)

nullablestringnullable, string

Plaid’s unique identifier for a test clock.

[`type`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-type)

stringstring

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`amount`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`status`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-status)

stringstring

The status of the recurring transfer.  
`active`: The recurring transfer is currently active.
`cancelled`: The recurring transfer was cancelled by the client or Plaid.
`expired`: The recurring transfer has completed all originations according to its recurring schedule.  
  

Possible values: `active`, `cancelled`, `expired`

[`ach_class`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`network`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-network)

stringstring

Networks eligible for recurring transfers.  
  

Possible values: `ach`, `same-day-ach`, `rtp`

[`account_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-account-id)

stringstring

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`funding_account_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`iso_currency_code`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

[`description`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-description)

stringstring

The description of the recurring transfer.

[`transfer_ids`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-transfer-ids)

[string][string]

The created transfer instances associated with this `recurring_transfer_id`. If the recurring transfer has been newly created, this array will be empty.

[`user`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`schedule`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-schedule)

objectobject

The schedule that the recurring transfer will be executed on.

[`interval_unit`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-schedule-interval-unit)

stringstring

The unit of the recurring interval.  
  

Possible values: `week`, `month`

Min length: `1`

[`interval_count`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-schedule-interval-count)

integerinteger

The number of recurring `interval_units` between originations. The recurring interval (before holiday adjustment) is calculated by multiplying `interval_unit` and `interval_count`.
For example, to schedule a recurring transfer which originates once every two weeks, set `interval_unit` = `week` and `interval_count` = 2.

[`interval_execution_day`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-schedule-interval-execution-day)

integerinteger

The day of the interval on which to schedule the transfer.  
If the `interval_unit` is `week`, `interval_execution_day` should be an integer from 1 (Monday) to 5 (Friday).  
If the `interval_unit` is `month`, `interval_execution_day` should be an integer indicating which day of the month to make the transfer on. Integers from 1 to 28 can be used to make a transfer on that day of the month. Negative integers from -1 to -5 can be used to make a transfer relative to the end of the month. To make a transfer on the last day of the month, use -1; to make the transfer on the second-to-last day, use -2, and so on.  
The transfer will be originated on the next available banking day if the designated day is a non banking day.

[`start_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-schedule-start-date)

stringstring

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). The recurring transfer will begin on the first `interval_execution_day` on or after the `start_date`.  
For `rtp` recurring transfers, `start_date` must be in the future.
Otherwise, if the first `interval_execution_day` on or after the start date is also the same day that `/transfer/recurring/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`end_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-recurring-transfer-schedule-end-date)

nullablestringnullable, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). The recurring transfer will end on the last `interval_execution_day` on or before the `end_date`.
If the `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/transfer/recurring/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`request_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "recurring_transfer": {
    "recurring_transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
    "created": "2022-07-05T12:48:37Z",
    "next_origination_date": "2022-10-28",
    "test_clock_id": null,
    "status": "active",
    "amount": "12.34",
    "description": "payment",
    "type": "debit",
    "ach_class": "ppd",
    "network": "ach",
    "origination_account_id": "",
    "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
    "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
    "iso_currency_code": "USD",
    "transfer_ids": [
      "271ef220-dbf8-caeb-a7dc-a2b3e8a80963",
      "c8dbaf75-2abb-e2dc-4171-12448e13b848"
    ],
    "user": {
      "legal_name": "Anne Charleston",
      "phone_number": "510-555-0128",
      "email_address": "acharleston@email.com",
      "address": {
        "street": "123 Main St.",
        "city": "San Francisco",
        "region": "CA",
        "postal_code": "94053",
        "country": "US"
      }
    },
    "schedule": {
      "start_date": "2022-10-01",
      "end_date": "2023-10-01",
      "interval_unit": "week",
      "interval_count": 1,
      "interval_execution_day": 5
    }
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/recurring/list`

#### List recurring transfers

Use the [`/transfer/recurring/list`](/docs/api/products/transfer/recurring-transfers/#transferrecurringlist) endpoint to see a list of all your recurring transfers and their statuses. Results are paginated; use the `count` and `offset` query parameters to retrieve the desired recurring transfers.

/transfer/recurring/list

**Request fields**

[`client_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`start_time`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-request-start-time)

stringstring

The start `created` datetime of recurring transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)  
  

Format: `date-time`

[`end_time`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-request-end-time)

stringstring

The end `created` datetime of recurring transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)  
  

Format: `date-time`

[`count`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-request-count)

integerinteger

The maximum number of recurring transfers to return.  
  

Minimum: `1`

Maximum: `25`

Default: `25`

[`offset`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-request-offset)

integerinteger

The number of recurring transfers to skip before returning results.  
  

Default: `0`

Minimum: `0`

[`funding_account_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-request-funding-account-id)

stringstring

Filter recurring transfers to only those with the specified `funding_account_id`.

/transfer/recurring/list

```
const request: TransferRecurringListRequest = {
  start_time: '2022-09-29T20:35:49Z',
  end_time: '2022-10-29T20:35:49Z',
  count: 1,
};

try {
  const response = await client.transferRecurringList(request);
} catch (error) {
  // handle error
}
```

/transfer/recurring/list

**Response fields**

[`recurring_transfers`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers)

[object][object]

[`recurring_transfer_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-recurring-transfer-id)

stringstring

Plaid’s unique identifier for a recurring transfer.

[`created`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-created)

stringstring

The datetime when this transfer was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`next_origination_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-next-origination-date)

nullablestringnullable, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
The next transfer origination date after bank holiday adjustment.  
  

Format: `date`

[`test_clock_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-test-clock-id)

nullablestringnullable, string

Plaid’s unique identifier for a test clock.

[`type`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-type)

stringstring

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`amount`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`status`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-status)

stringstring

The status of the recurring transfer.  
`active`: The recurring transfer is currently active.
`cancelled`: The recurring transfer was cancelled by the client or Plaid.
`expired`: The recurring transfer has completed all originations according to its recurring schedule.  
  

Possible values: `active`, `cancelled`, `expired`

[`ach_class`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`network`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-network)

stringstring

Networks eligible for recurring transfers.  
  

Possible values: `ach`, `same-day-ach`, `rtp`

[`account_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-account-id)

stringstring

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`funding_account_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`iso_currency_code`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

[`description`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-description)

stringstring

The description of the recurring transfer.

[`transfer_ids`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-transfer-ids)

[string][string]

The created transfer instances associated with this `recurring_transfer_id`. If the recurring transfer has been newly created, this array will be empty.

[`user`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`schedule`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-schedule)

objectobject

The schedule that the recurring transfer will be executed on.

[`interval_unit`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-schedule-interval-unit)

stringstring

The unit of the recurring interval.  
  

Possible values: `week`, `month`

Min length: `1`

[`interval_count`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-schedule-interval-count)

integerinteger

The number of recurring `interval_units` between originations. The recurring interval (before holiday adjustment) is calculated by multiplying `interval_unit` and `interval_count`.
For example, to schedule a recurring transfer which originates once every two weeks, set `interval_unit` = `week` and `interval_count` = 2.

[`interval_execution_day`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-schedule-interval-execution-day)

integerinteger

The day of the interval on which to schedule the transfer.  
If the `interval_unit` is `week`, `interval_execution_day` should be an integer from 1 (Monday) to 5 (Friday).  
If the `interval_unit` is `month`, `interval_execution_day` should be an integer indicating which day of the month to make the transfer on. Integers from 1 to 28 can be used to make a transfer on that day of the month. Negative integers from -1 to -5 can be used to make a transfer relative to the end of the month. To make a transfer on the last day of the month, use -1; to make the transfer on the second-to-last day, use -2, and so on.  
The transfer will be originated on the next available banking day if the designated day is a non banking day.

[`start_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-schedule-start-date)

stringstring

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). The recurring transfer will begin on the first `interval_execution_day` on or after the `start_date`.  
For `rtp` recurring transfers, `start_date` must be in the future.
Otherwise, if the first `interval_execution_day` on or after the start date is also the same day that `/transfer/recurring/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`end_date`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-recurring-transfers-schedule-end-date)

nullablestringnullable, string

A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). The recurring transfer will end on the last `interval_execution_day` on or before the `end_date`.
If the `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/transfer/recurring/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.  
  

Format: `date`

[`request_id`](/docs/api/products/transfer/recurring-transfers/#transfer-recurring-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "recurring_transfers": [
    {
      "recurring_transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
      "created": "2022-07-05T12:48:37Z",
      "next_origination_date": "2022-10-28",
      "test_clock_id": null,
      "status": "active",
      "amount": "12.34",
      "description": "payment",
      "type": "debit",
      "ach_class": "ppd",
      "network": "ach",
      "origination_account_id": "",
      "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
      "iso_currency_code": "USD",
      "transfer_ids": [
        "4242fc8d-3ec6-fb38-fa0c-a8e37d03cd57"
      ],
      "user": {
        "legal_name": "Anne Charleston",
        "phone_number": "510-555-0128",
        "email_address": "acharleston@email.com",
        "address": {
          "street": "123 Main St.",
          "city": "San Francisco",
          "region": "CA",
          "postal_code": "94053",
          "country": "US"
        }
      },
      "schedule": {
        "start_date": "2022-10-01",
        "end_date": "2023-10-01",
        "interval_unit": "week",
        "interval_count": 1,
        "interval_execution_day": 5
      }
    }
  ],
  "request_id": "saKrIBuEB9qJZno"
}
```

### Webhooks

=\*=\*=\*=

#### `RECURRING_NEW_TRANSFER`

Fired when a new transfer of a recurring transfer is originated.

**Properties**

[`webhook_type`](/docs/api/products/transfer/recurring-transfers/#RecurringNewTransferWebhook-webhook-type)

stringstring

`TRANSFER`

[`webhook_code`](/docs/api/products/transfer/recurring-transfers/#RecurringNewTransferWebhook-webhook-code)

stringstring

`RECURRING_NEW_TRANSFER`

[`recurring_transfer_id`](/docs/api/products/transfer/recurring-transfers/#RecurringNewTransferWebhook-recurring-transfer-id)

stringstring

Plaid’s unique identifier for a recurring transfer.

[`transfer_id`](/docs/api/products/transfer/recurring-transfers/#RecurringNewTransferWebhook-transfer-id)

stringstring

Plaid’s unique identifier for a transfer.

[`environment`](/docs/api/products/transfer/recurring-transfers/#RecurringNewTransferWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "TRANSFER",
  "webhook_code": "RECURRING_NEW_TRANSFER",
  "recurring_transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
  "transfer_id": "271ef220-dbf8-caeb-a7dc-a2b3e8a80963",
  "environment": "production"
}
```

=\*=\*=\*=

#### `RECURRING_TRANSFER_SKIPPED`

Fired when Plaid is unable to originate a new ACH transaction of the recurring transfer on the planned date.

**Properties**

[`webhook_type`](/docs/api/products/transfer/recurring-transfers/#RecurringTransferSkippedWebhook-webhook-type)

stringstring

`TRANSFER`

[`webhook_code`](/docs/api/products/transfer/recurring-transfers/#RecurringTransferSkippedWebhook-webhook-code)

stringstring

`RECURRING_TRANSFER_SKIPPED`

[`recurring_transfer_id`](/docs/api/products/transfer/recurring-transfers/#RecurringTransferSkippedWebhook-recurring-transfer-id)

stringstring

Plaid’s unique identifier for a recurring transfer.

[`authorization_decision`](/docs/api/products/transfer/recurring-transfers/#RecurringTransferSkippedWebhook-authorization-decision)

stringstring

A decision regarding the proposed transfer.  
`approved` – The proposed transfer has received the end user's consent and has been approved for processing by Plaid. The `decision_rationale` field is set if Plaid was unable to fetch the account information. You may proceed with the transfer, but further review is recommended. Refer to the `code` field in the `decision_rationale` object for details.  
`declined` – Plaid reviewed the proposed transfer and declined processing. Refer to the `code` field in the `decision_rationale` object for details.  
`user_action_required` – An action is required before Plaid can assess the transfer risk and make a decision. The most common scenario is to update authentication for an Item. To complete the required action, initialize Link by setting `transfer.authorization_id` in the request of `/link/token/create`. After Link flow is completed, you may re-attempt the authorization request.  
For `guarantee` requests, `approved` indicates the transfer is eligible for Plaid's guarantee, and `declined` indicates Plaid will not provide guarantee coverage for the transfer. `user_action_required` indicates you should follow the above guidance before re-attempting.  
  

Possible values: `approved`, `declined`, `user_action_required`

[`authorization_decision_rationale_code`](/docs/api/products/transfer/recurring-transfers/#RecurringTransferSkippedWebhook-authorization-decision-rationale-code)

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

[`skipped_origination_date`](/docs/api/products/transfer/recurring-transfers/#RecurringTransferSkippedWebhook-skipped-origination-date)

stringstring

The planned date on which Plaid is unable to originate a new ACH transaction of the recurring transfer. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`environment`](/docs/api/products/transfer/recurring-transfers/#RecurringTransferSkippedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "TRANSFER",
  "webhook_code": "RECURRING_TRANSFER_SKIPPED",
  "recurring_transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
  "authorization_decision": "declined",
  "authorization_decision_rationale_code": "NSF",
  "skipped_origination_date": "2022-11-30",
  "environment": "production"
}
```

=\*=\*=\*=

#### `RECURRING_CANCELLED`

Fired when a recurring transfer is cancelled by Plaid.

**Properties**

[`webhook_type`](/docs/api/products/transfer/recurring-transfers/#RecurringCancelledWebhook-webhook-type)

stringstring

`TRANSFER`

[`webhook_code`](/docs/api/products/transfer/recurring-transfers/#RecurringCancelledWebhook-webhook-code)

stringstring

`RECURRING_CANCELLED`

[`recurring_transfer_id`](/docs/api/products/transfer/recurring-transfers/#RecurringCancelledWebhook-recurring-transfer-id)

stringstring

Plaid’s unique identifier for a recurring transfer.

[`environment`](/docs/api/products/transfer/recurring-transfers/#RecurringCancelledWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "TRANSFER",
  "webhook_code": "RECURRING_CANCELLED",
  "recurring_transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
