---
title: "API - Reading Transfers | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/transfer/reading-transfers/"
scraped_at: "2026-03-07T22:04:25+00:00"
---

# Reading Transfers and Transfer events

#### API reference for Transfer read and Transfer event endpoints and webhooks

For how-to guidance, see the [Transfer events documentation](/docs/transfer/reconciling-transfers/).

| Reading Transfers |  |
| --- | --- |
| [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) | Retrieve information about a transfer |
| [`/transfer/list`](/docs/api/products/transfer/reading-transfers/#transferlist) | Retrieve a list of transfers and their statuses |
| [`/transfer/event/list`](/docs/api/products/transfer/reading-transfers/#transfereventlist) | Retrieve a list of transfer events |
| [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) | Sync transfer events |
| [`/transfer/sweep/get`](/docs/api/products/transfer/reading-transfers/#transfersweepget) | Retrieve information about a sweep |
| [`/transfer/sweep/list`](/docs/api/products/transfer/reading-transfers/#transfersweeplist) | Retrieve a list of sweeps |

| Webhooks |  |
| --- | --- |
| [`TRANSFER_EVENTS_UPDATE`](/docs/api/products/transfer/reading-transfers/#transfer_events_update) | New transfer events available |

### Endpoints

=\*=\*=\*=

#### `/transfer/get`

#### Retrieve a transfer

The [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) endpoint fetches information about the transfer corresponding to the given `transfer_id` or `authorization_id`. One of `transfer_id` or `authorization_id` must be populated but not both.

/transfer/get

**Request fields**

[`client_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/reading-transfers/#transfer-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-request-transfer-id)

stringstring

Plaid’s unique identifier for a transfer.

[`authorization_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-request-authorization-id)

stringstring

Plaid’s unique identifier for a transfer authorization.

/transfer/get

```
const request: TransferGetRequest = {
  transfer_id: '460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9',
};
try {
  const response = await plaidClient.transferGet(request);
  const transfer = response.data.transfer;
} catch (error) {
  // handle error
}
```

/transfer/get

**Response fields**

[`transfer`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer)

objectobject

Represents a transfer within the Transfers API.

[`id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-id)

stringstring

Plaid’s unique identifier for a transfer.

[`authorization_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-authorization-id)

stringstring

Plaid’s unique identifier for a transfer authorization.

[`ach_class`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`account_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-account-id)

stringstring

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-funding-account-id)

nullablestringnullable, string

The id of the associated funding account, available in the Plaid Dashboard. If present, this indicates which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`type`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-type)

stringstring

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`user`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`amount`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-description)

stringstring

The description of the transfer.

[`created`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-created)

stringstring

The datetime when this transfer was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`status`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-status)

stringstring

The status of the transfer.  
`pending`: A new transfer was created; it is in the pending state.
`posted`: The transfer has been successfully submitted to the payment network.
`settled`: The transfer was successfully completed by the payment network. Note that funds from received debits are not available to be moved out of the Ledger until the transfer reaches `funds_available` status. For credit transactions, `settled` means the funds have been delivered to the receiving bank account. This is the terminal state of a successful credit transfer.
`funds_available`: Funds from the transfer have been released from hold and applied to the ledger's available balance. (Only applicable to ACH debits.) This is the terminal state of a successful debit transfer.
`cancelled`: The transfer was cancelled by the client. This is the terminal state of a cancelled transfer.
`failed`: The transfer failed, no funds were moved. This is the terminal state of a failed transfer.
`returned`: A posted transfer was returned. This is the terminal state of a returned transfer.  
  

Possible values: `pending`, `posted`, `settled`, `funds_available`, `cancelled`, `failed`, `returned`

[`sweep_status`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-sweep-status)

nullablestringnullable, string

The status of the sweep for the transfer.  
`unswept`: The transfer hasn't been swept yet.
`swept`: The transfer was swept to the sweep account.
`swept_settled`: Credits are available to be withdrawn or debits have been deducted from the customer’s business checking account.
`return_swept`: The transfer was returned, funds were pulled back or pushed back to the sweep account.
`null`: The transfer will never be swept (e.g. if the transfer is cancelled or returned before being swept)  
  

Possible values: `null`, `unswept`, `swept`, `swept_settled`, `return_swept`

[`network`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-network)

stringstring

The network or rails used for the transfer.  
For transfers submitted as `ach` or `same-day-ach`, the Standard ACH cutoff is 8:30 PM Eastern Time.  
For transfers submitted as `same-day-ach`, the Same Day ACH cutoff is 3:00 PM Eastern Time. It is recommended to send the request 15 minutes prior to the cutoff to ensure that it will be processed in time for submission before the cutoff. If the transfer is processed after this cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur same-day charges; this will apply to both legs of the transfer if applicable. The transaction limit for a Same Day ACH transfer is $1,000,000. Authorization requests sent with an amount greater than $1,000,000 will fail.  
For transfers submitted as `rtp`, Plaid will automatically route between Real Time Payment rail by TCH or FedNow rails as necessary. If a transfer is submitted as `rtp` and the counterparty account is not eligible for RTP, the `/transfer/authorization/create` request will fail with an `INVALID_FIELD` error code. To pre-check to determine whether a counterparty account can support RTP, call `/transfer/capabilities/get` before calling `/transfer/authorization/create`.  
Wire transfers are currently in early availability. To request access to `wire` as a payment network, contact your Account Manager. For transfers submitted as `wire`, the `type` must be `credit`; wire debits are not supported. The cutoff to submit a wire payment is 6:30 PM Eastern Time on a business day; wires submitted after that time will be processed on the next business day. The transaction limit for a wire is $999,999.99. Authorization requests sent with an amount greater than $999,999.99 will fail.  
  

Possible values: `ach`, `same-day-ach`, `rtp`, `wire`

[`wire_details`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-wire-details)

nullableobjectnullable, object

Information specific to wire transfers.

[`message_to_beneficiary`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-wire-details-message-to-beneficiary)

nullablestringnullable, string

Additional information from the wire originator to the beneficiary. Max 140 characters.

[`wire_return_fee`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-wire-details-wire-return-fee)

nullablestringnullable, string

The fee amount deducted from the original transfer during a wire return, if applicable.

[`cancellable`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-cancellable)

booleanboolean

When `true`, you can still cancel this transfer.

[`failure_reason`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-failure-reason)

nullableobjectnullable, object

The failure reason if the event type for a transfer is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the transfer status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`ach_return_code`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-failure-reason-ach-return-code)

deprecatednullablestringdeprecated, nullable, string

The ACH return code, e.g. `R01`. A return code will be provided if and only if the transfer status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes).

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-failure-reason-description)

stringstring

A human-readable description of the reason for the failure or reversal.

[`metadata`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-metadata)

nullableobjectnullable, object

The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply:
The JSON values must be Strings (no nested JSON objects allowed)
Only ASCII characters may be used
Maximum of 50 key/value pairs
Maximum key length of 40 characters
Maximum value length of 500 characters

[`iso_currency_code`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

[`standard_return_window`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-standard-return-window)

nullablestringnullable, string

The date 3 business days from settlement date indicating the following ACH returns can no longer happen: R01, R02, R03, R29. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`unauthorized_return_window`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-unauthorized-return-window)

nullablestringnullable, string

The date 61 business days from settlement date indicating the following ACH returns can no longer happen: R05, R07, R10, R11, R51, R33, R37, R38, R51, R52, R53. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`expected_settlement_date`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-expected-settlement-date)

deprecatednullablestringdeprecated, nullable, string

Deprecated for Plaid Ledger clients, use `expected_funds_available_date` instead.  
  

Format: `date`

[`expected_funds_available_date`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-expected-funds-available-date)

nullablestringnullable, string

The expected date when funds from a transfer will be made available and can be withdrawn from the associated ledger balance, assuming the debit does not return before this date. If the transfer does return before this date, this field will be null. Only applies to debit transfers. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`originator_client_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-originator-client-id)

nullablestringnullable, string

The Plaid client ID that is the originator of this transfer. Only present if created on behalf of another client as a [Platform customer](https://plaid.com/docs/transfer/application/#originators-vs-platforms).

[`refunds`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds)

[object][object]

A list of refunds associated with this transfer.

[`id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-id)

stringstring

Plaid’s unique identifier for a refund.

[`transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-transfer-id)

stringstring

The ID of the transfer to refund.

[`amount`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-amount)

stringstring

The amount of the refund (decimal string with two digits of precision e.g. "10.00").

[`status`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-status)

stringstring

The status of the refund.  
`pending`: A new refund was created; it is in the pending state.
`posted`: The refund has been successfully submitted to the payment network.
`settled`: Credits have been refunded to the Plaid linked account.
`cancelled`: The refund was cancelled by the client.
`failed`: The refund has failed.
`returned`: The refund was returned.  
  

Possible values: `pending`, `posted`, `cancelled`, `failed`, `settled`, `returned`

[`failure_reason`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-failure-reason)

nullableobjectnullable, object

The failure reason if the event type for a refund is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the refund status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`ach_return_code`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-failure-reason-ach-return-code)

deprecatednullablestringdeprecated, nullable, string

The ACH return code, e.g. `R01`. A return code will be provided if and only if the refund status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes). This field is deprecated in favor of the more versatile `failure_code`, which encompasses non-ACH failure codes as well.

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-failure-reason-description)

stringstring

A human-readable description of the reason for the failure or reversal.

[`ledger_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`created`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-created)

stringstring

The datetime when this refund was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`network_trace_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-refunds-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`recurring_transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-recurring-transfer-id)

nullablestringnullable, string

The id of the recurring transfer if this transfer belongs to a recurring transfer.

[`expected_sweep_settlement_schedule`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-expected-sweep-settlement-schedule)

[object][object]

The expected sweep settlement schedule of this transfer, assuming this transfer is not `returned`. Only applies to ACH debit transfers.

[`sweep_settlement_date`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-expected-sweep-settlement-schedule-sweep-settlement-date)

stringstring

The settlement date of a sweep for this transfer.  
  

Format: `date`

[`swept_settled_amount`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-expected-sweep-settlement-schedule-swept-settled-amount)

stringstring

The accumulated amount that has been swept by `sweep_settlement_date`.

[`credit_funds_source`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-credit-funds-source)

deprecatednullablestringdeprecated, nullable, string

This field is now deprecated. You may ignore it for transfers created on and after 12/01/2023.  
Specifies the source of funds for the transfer. Only valid for `credit` transfers, and defaults to `sweep` if not specified. This field is not specified for `debit` transfers.  
`sweep` - Sweep funds from your funding account
`prefunded_rtp_credits` - Use your prefunded RTP credit balance with Plaid
`prefunded_ach_credits` - Use your prefunded ACH credit balance with Plaid  
  

Possible values: `sweep`, `prefunded_rtp_credits`, `prefunded_ach_credits`, `null`

[`facilitator_fee`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-facilitator-fee)

stringstring

The amount to deduct from `transfer.amount` and distribute to the platform’s Ledger balance as a facilitator fee (decimal string with two digits of precision e.g. "10.00"). The remainder will go to the end-customer’s Ledger balance. This must be value greater than 0 and less than or equal to the `transfer.amount`.

[`network_trace_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`request_id`](/docs/api/products/transfer/reading-transfers/#transfer-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transfer": {
    "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
    "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
    "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
    "ach_class": "ppd",
    "amount": "12.34",
    "cancellable": true,
    "created": "2020-08-06T17:27:15Z",
    "description": "Desc",
    "guarantee_decision": null,
    "guarantee_decision_rationale": null,
    "failure_reason": {
      "failure_code": "R13",
      "description": "Invalid ACH routing number"
    },
    "id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
    "authorization_id": "c9f90aa1-2949-c799-e2b6-ea05c89bb586",
    "metadata": {
      "key1": "value1",
      "key2": "value2"
    },
    "network": "ach",
    "origination_account_id": "",
    "originator_client_id": null,
    "refunds": [],
    "status": "pending",
    "type": "credit",
    "iso_currency_code": "USD",
    "standard_return_window": "2020-08-07",
    "unauthorized_return_window": "2020-10-07",
    "expected_settlement_date": "2020-08-04",
    "user": {
      "email_address": "acharleston@email.com",
      "legal_name": "Anne Charleston",
      "phone_number": "510-555-0128",
      "address": {
        "street": "123 Main St.",
        "city": "San Francisco",
        "region": "CA",
        "postal_code": "94053",
        "country": "US"
      }
    },
    "recurring_transfer_id": null,
    "credit_funds_source": "sweep",
    "facilitator_fee": "1.23",
    "network_trace_id": null
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/list`

#### List transfers

Use the [`/transfer/list`](/docs/api/products/transfer/reading-transfers/#transferlist) endpoint to see a list of all your transfers and their statuses. Results are paginated; use the `count` and `offset` query parameters to retrieve the desired transfers.

/transfer/list

**Request fields**

[`client_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/reading-transfers/#transfer-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`start_date`](/docs/api/products/transfer/reading-transfers/#transfer-list-request-start-date)

stringstring

The start `created` datetime of transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)  
  

Format: `date-time`

[`end_date`](/docs/api/products/transfer/reading-transfers/#transfer-list-request-end-date)

stringstring

The end `created` datetime of transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)  
  

Format: `date-time`

[`count`](/docs/api/products/transfer/reading-transfers/#transfer-list-request-count)

integerinteger

The maximum number of transfers to return.  
  

Minimum: `1`

Maximum: `25`

Default: `25`

[`offset`](/docs/api/products/transfer/reading-transfers/#transfer-list-request-offset)

integerinteger

The number of transfers to skip before returning results.  
  

Default: `0`

Minimum: `0`

[`originator_client_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-request-originator-client-id)

stringstring

Filter transfers to only those with the specified originator client.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-request-funding-account-id)

stringstring

Filter transfers to only those with the specified `funding_account_id`.

/transfer/list

```
const request: TransferListRequest = {
  start_date: '2019-12-06T22:35:49Z',
  end_date: '2019-12-12T22:35:49Z',
  count: 14,
  offset: 2,
  origination_account_id: '8945fedc-e703-463d-86b1-dc0607b55460',
};
try {
  const response = await plaidClient.transferList(request);
  const transfers = response.data.transfers;
  for (const transfer of transfers) {
    // iterate through transfers
  }
} catch (error) {
  // handle error
}
```

/transfer/list

**Response fields**

[`transfers`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers)

[object][object]

[`id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-id)

stringstring

Plaid’s unique identifier for a transfer.

[`authorization_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-authorization-id)

stringstring

Plaid’s unique identifier for a transfer authorization.

[`ach_class`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`account_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-account-id)

stringstring

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-funding-account-id)

nullablestringnullable, string

The id of the associated funding account, available in the Plaid Dashboard. If present, this indicates which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`type`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-type)

stringstring

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`user`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`amount`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-description)

stringstring

The description of the transfer.

[`created`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-created)

stringstring

The datetime when this transfer was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`status`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-status)

stringstring

The status of the transfer.  
`pending`: A new transfer was created; it is in the pending state.
`posted`: The transfer has been successfully submitted to the payment network.
`settled`: The transfer was successfully completed by the payment network. Note that funds from received debits are not available to be moved out of the Ledger until the transfer reaches `funds_available` status. For credit transactions, `settled` means the funds have been delivered to the receiving bank account. This is the terminal state of a successful credit transfer.
`funds_available`: Funds from the transfer have been released from hold and applied to the ledger's available balance. (Only applicable to ACH debits.) This is the terminal state of a successful debit transfer.
`cancelled`: The transfer was cancelled by the client. This is the terminal state of a cancelled transfer.
`failed`: The transfer failed, no funds were moved. This is the terminal state of a failed transfer.
`returned`: A posted transfer was returned. This is the terminal state of a returned transfer.  
  

Possible values: `pending`, `posted`, `settled`, `funds_available`, `cancelled`, `failed`, `returned`

[`sweep_status`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-sweep-status)

nullablestringnullable, string

The status of the sweep for the transfer.  
`unswept`: The transfer hasn't been swept yet.
`swept`: The transfer was swept to the sweep account.
`swept_settled`: Credits are available to be withdrawn or debits have been deducted from the customer’s business checking account.
`return_swept`: The transfer was returned, funds were pulled back or pushed back to the sweep account.
`null`: The transfer will never be swept (e.g. if the transfer is cancelled or returned before being swept)  
  

Possible values: `null`, `unswept`, `swept`, `swept_settled`, `return_swept`

[`network`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-network)

stringstring

The network or rails used for the transfer.  
For transfers submitted as `ach` or `same-day-ach`, the Standard ACH cutoff is 8:30 PM Eastern Time.  
For transfers submitted as `same-day-ach`, the Same Day ACH cutoff is 3:00 PM Eastern Time. It is recommended to send the request 15 minutes prior to the cutoff to ensure that it will be processed in time for submission before the cutoff. If the transfer is processed after this cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur same-day charges; this will apply to both legs of the transfer if applicable. The transaction limit for a Same Day ACH transfer is $1,000,000. Authorization requests sent with an amount greater than $1,000,000 will fail.  
For transfers submitted as `rtp`, Plaid will automatically route between Real Time Payment rail by TCH or FedNow rails as necessary. If a transfer is submitted as `rtp` and the counterparty account is not eligible for RTP, the `/transfer/authorization/create` request will fail with an `INVALID_FIELD` error code. To pre-check to determine whether a counterparty account can support RTP, call `/transfer/capabilities/get` before calling `/transfer/authorization/create`.  
Wire transfers are currently in early availability. To request access to `wire` as a payment network, contact your Account Manager. For transfers submitted as `wire`, the `type` must be `credit`; wire debits are not supported. The cutoff to submit a wire payment is 6:30 PM Eastern Time on a business day; wires submitted after that time will be processed on the next business day. The transaction limit for a wire is $999,999.99. Authorization requests sent with an amount greater than $999,999.99 will fail.  
  

Possible values: `ach`, `same-day-ach`, `rtp`, `wire`

[`wire_details`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-wire-details)

nullableobjectnullable, object

Information specific to wire transfers.

[`message_to_beneficiary`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-wire-details-message-to-beneficiary)

nullablestringnullable, string

Additional information from the wire originator to the beneficiary. Max 140 characters.

[`wire_return_fee`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-wire-details-wire-return-fee)

nullablestringnullable, string

The fee amount deducted from the original transfer during a wire return, if applicable.

[`cancellable`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-cancellable)

booleanboolean

When `true`, you can still cancel this transfer.

[`failure_reason`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-failure-reason)

nullableobjectnullable, object

The failure reason if the event type for a transfer is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the transfer status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`ach_return_code`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-failure-reason-ach-return-code)

deprecatednullablestringdeprecated, nullable, string

The ACH return code, e.g. `R01`. A return code will be provided if and only if the transfer status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes).

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-failure-reason-description)

stringstring

A human-readable description of the reason for the failure or reversal.

[`metadata`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-metadata)

nullableobjectnullable, object

The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply:
The JSON values must be Strings (no nested JSON objects allowed)
Only ASCII characters may be used
Maximum of 50 key/value pairs
Maximum key length of 40 characters
Maximum value length of 500 characters

[`iso_currency_code`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

[`standard_return_window`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-standard-return-window)

nullablestringnullable, string

The date 3 business days from settlement date indicating the following ACH returns can no longer happen: R01, R02, R03, R29. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`unauthorized_return_window`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-unauthorized-return-window)

nullablestringnullable, string

The date 61 business days from settlement date indicating the following ACH returns can no longer happen: R05, R07, R10, R11, R51, R33, R37, R38, R51, R52, R53. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`expected_settlement_date`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-expected-settlement-date)

deprecatednullablestringdeprecated, nullable, string

Deprecated for Plaid Ledger clients, use `expected_funds_available_date` instead.  
  

Format: `date`

[`expected_funds_available_date`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-expected-funds-available-date)

nullablestringnullable, string

The expected date when funds from a transfer will be made available and can be withdrawn from the associated ledger balance, assuming the debit does not return before this date. If the transfer does return before this date, this field will be null. Only applies to debit transfers. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`originator_client_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-originator-client-id)

nullablestringnullable, string

The Plaid client ID that is the originator of this transfer. Only present if created on behalf of another client as a [Platform customer](https://plaid.com/docs/transfer/application/#originators-vs-platforms).

[`refunds`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds)

[object][object]

A list of refunds associated with this transfer.

[`id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-id)

stringstring

Plaid’s unique identifier for a refund.

[`transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-transfer-id)

stringstring

The ID of the transfer to refund.

[`amount`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-amount)

stringstring

The amount of the refund (decimal string with two digits of precision e.g. "10.00").

[`status`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-status)

stringstring

The status of the refund.  
`pending`: A new refund was created; it is in the pending state.
`posted`: The refund has been successfully submitted to the payment network.
`settled`: Credits have been refunded to the Plaid linked account.
`cancelled`: The refund was cancelled by the client.
`failed`: The refund has failed.
`returned`: The refund was returned.  
  

Possible values: `pending`, `posted`, `cancelled`, `failed`, `settled`, `returned`

[`failure_reason`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-failure-reason)

nullableobjectnullable, object

The failure reason if the event type for a refund is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the refund status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`ach_return_code`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-failure-reason-ach-return-code)

deprecatednullablestringdeprecated, nullable, string

The ACH return code, e.g. `R01`. A return code will be provided if and only if the refund status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes). This field is deprecated in favor of the more versatile `failure_code`, which encompasses non-ACH failure codes as well.

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-failure-reason-description)

stringstring

A human-readable description of the reason for the failure or reversal.

[`ledger_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`created`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-created)

stringstring

The datetime when this refund was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`network_trace_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-refunds-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`recurring_transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-recurring-transfer-id)

nullablestringnullable, string

The id of the recurring transfer if this transfer belongs to a recurring transfer.

[`expected_sweep_settlement_schedule`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-expected-sweep-settlement-schedule)

[object][object]

The expected sweep settlement schedule of this transfer, assuming this transfer is not `returned`. Only applies to ACH debit transfers.

[`sweep_settlement_date`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-expected-sweep-settlement-schedule-sweep-settlement-date)

stringstring

The settlement date of a sweep for this transfer.  
  

Format: `date`

[`swept_settled_amount`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-expected-sweep-settlement-schedule-swept-settled-amount)

stringstring

The accumulated amount that has been swept by `sweep_settlement_date`.

[`credit_funds_source`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-credit-funds-source)

deprecatednullablestringdeprecated, nullable, string

This field is now deprecated. You may ignore it for transfers created on and after 12/01/2023.  
Specifies the source of funds for the transfer. Only valid for `credit` transfers, and defaults to `sweep` if not specified. This field is not specified for `debit` transfers.  
`sweep` - Sweep funds from your funding account
`prefunded_rtp_credits` - Use your prefunded RTP credit balance with Plaid
`prefunded_ach_credits` - Use your prefunded ACH credit balance with Plaid  
  

Possible values: `sweep`, `prefunded_rtp_credits`, `prefunded_ach_credits`, `null`

[`facilitator_fee`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-facilitator-fee)

stringstring

The amount to deduct from `transfer.amount` and distribute to the platform’s Ledger balance as a facilitator fee (decimal string with two digits of precision e.g. "10.00"). The remainder will go to the end-customer’s Ledger balance. This must be value greater than 0 and less than or equal to the `transfer.amount`.

[`network_trace_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-transfers-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`request_id`](/docs/api/products/transfer/reading-transfers/#transfer-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transfers": [
    {
      "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
      "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
      "ach_class": "ppd",
      "amount": "12.34",
      "cancellable": true,
      "created": "2019-12-09T17:27:15Z",
      "description": "Desc",
      "guarantee_decision": null,
      "guarantee_decision_rationale": null,
      "failure_reason": {
        "failure_code": "R13",
        "description": "Invalid ACH routing number"
      },
      "id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
      "authorization_id": "c9f90aa1-2949-c799-e2b6-ea05c89bb586",
      "metadata": {
        "key1": "value1",
        "key2": "value2"
      },
      "network": "ach",
      "origination_account_id": "",
      "originator_client_id": null,
      "refunds": [],
      "status": "pending",
      "type": "credit",
      "iso_currency_code": "USD",
      "standard_return_window": "2020-08-07",
      "unauthorized_return_window": "2020-10-07",
      "expected_settlement_date": "2020-08-04",
      "user": {
        "email_address": "acharleston@email.com",
        "legal_name": "Anne Charleston",
        "phone_number": "510-555-0128",
        "address": {
          "street": "123 Main St.",
          "city": "San Francisco",
          "region": "CA",
          "postal_code": "94053",
          "country": "US"
        }
      },
      "recurring_transfer_id": null,
      "credit_funds_source": "sweep",
      "facilitator_fee": "1.23",
      "network_trace_id": null
    }
  ],
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/event/list`

#### List transfer events

Use the [`/transfer/event/list`](/docs/api/products/transfer/reading-transfers/#transfereventlist) endpoint to get a list of transfer events based on specified filter criteria.

/transfer/event/list

**Request fields**

[`client_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`start_date`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-start-date)

stringstring

The start `created` datetime of transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)  
  

Format: `date-time`

[`end_date`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-end-date)

stringstring

The end `created` datetime of transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)  
  

Format: `date-time`

[`transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-transfer-id)

stringstring

Plaid’s unique identifier for a transfer.

[`account_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-account-id)

stringstring

The account ID to get events for all transactions to/from an account.

[`transfer_type`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-transfer-type)

stringstring

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into your origination account; a `credit` indicates a transfer of money out of your origination account.  
  

Possible values: `debit`, `credit`, `null`

[`event_types`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-event-types)

[string][string]

Filter events by event type.  
  

Possible values: `pending`, `cancelled`, `failed`, `posted`, `settled`, `funds_available`, `returned`, `swept`, `swept_settled`, `return_swept`, `sweep.pending`, `sweep.posted`, `sweep.settled`, `sweep.returned`, `sweep.failed`, `sweep.funds_available`, `refund.pending`, `refund.cancelled`, `refund.failed`, `refund.posted`, `refund.settled`, `refund.returned`, `refund.swept`, `refund.return_swept`

[`sweep_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-sweep-id)

stringstring

Plaid’s unique identifier for a sweep.

[`count`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-count)

integerinteger

The maximum number of transfer events to return. If the number of events matching the above parameters is greater than `count`, the most recent events will be returned.  
  

Default: `25`

Maximum: `25`

Minimum: `1`

[`offset`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-offset)

integerinteger

The offset into the list of transfer events. When `count`=25 and `offset`=0, the first 25 events will be returned. When `count`=25 and `offset`=25, the next 25 events will be returned.  
  

Default: `0`

Minimum: `0`

[`originator_client_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-originator-client-id)

stringstring

Filter transfer events to only those with the specified originator client.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-request-funding-account-id)

stringstring

Filter transfer events to only those with the specified `funding_account_id`.

/transfer/event/list

```
const request: TransferEventListRequest = {
  start_date: '2019-12-06T22:35:49Z',
  end_date: '2019-12-12T22:35:49Z',
  transfer_id: '460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9',
  account_id: '3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr',
  transfer_type: 'credit',
  event_types: ['pending', 'posted'],
  count: 14,
  offset: 2,
  origination_account_id: '8945fedc-e703-463d-86b1-dc0607b55460',
};
try {
  const response = await plaidClient.transferEventList(request);
  const events = response.data.transfer_events;
  for (const event of events) {
    // iterate through events
  }
} catch (error) {
  // handle error
}
```

/transfer/event/list

**Response fields**

[`transfer_events`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events)

[object][object]

[`event_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-event-id)

integerinteger

Plaid’s unique identifier for this event. IDs are sequential unsigned 64-bit integers.  
  

Minimum: `0`

[`timestamp`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-timestamp)

stringstring

The datetime when this event occurred. This will be of the form `2006-01-02T15:04:05Z`.  
  

Format: `date-time`

[`event_type`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-event-type)

stringstring

The type of event that this transfer represents. Event types with prefix `sweep` represents events for Plaid Ledger sweeps.  
`pending`: A new transfer was created; it is in the pending state.  
`cancelled`: The transfer was cancelled by the client.  
`failed`: The transfer failed, no funds were moved.  
`posted`: The transfer has been successfully submitted to the payment network.  
`settled`: The transfer has been successfully completed by the payment network.  
`funds_available`: Funds from the transfer have been released from hold and applied to the ledger's available balance. (Only applicable to ACH debits.)  
`returned`: A posted transfer was returned.  
`swept`: The transfer was swept to / from the sweep account.  
`swept_settled`: Credits are available to be withdrawn or debits have been deducted from the customer’s business checking account.  
`return_swept`: Due to the transfer being returned, funds were pulled from or pushed back to the sweep account.  
`sweep.pending`: A new ledger sweep was created; it is in the pending state.  
`sweep.posted`: The ledger sweep has been successfully submitted to the payment network.  
`sweep.settled`: The transaction has settled in the funding account. This means that funds withdrawn from Plaid Ledger balance have reached the funding account, or funds to be deposited into the Plaid Ledger Balance have been pulled, and the hold period has begun.  
`sweep.returned`: A posted ledger sweep was returned.  
`sweep.failed`: The ledger sweep failed, no funds were moved.  
`sweep.funds_available`: Funds from the ledger sweep have been released from hold and applied to the ledger's available balance. This is only applicable to debits.  
`refund.pending`: A new refund was created; it is in the pending state.  
`refund.cancelled`: The refund was cancelled.  
`refund.failed`: The refund failed, no funds were moved.  
`refund.posted`: The refund has been successfully submitted to the payment network.  
`refund.settled`: The refund transaction has settled in the Plaid linked account.  
`refund.returned`: A posted refund was returned.  
`refund.swept`: The refund was swept from the sweep account.  
`refund.return_swept`: Due to the refund being returned, funds were pushed back to the sweep account.  
  

Possible values: `pending`, `cancelled`, `failed`, `posted`, `settled`, `funds_available`, `returned`, `swept`, `swept_settled`, `return_swept`, `sweep.pending`, `sweep.posted`, `sweep.settled`, `sweep.returned`, `sweep.failed`, `sweep.funds_available`, `refund.pending`, `refund.cancelled`, `refund.failed`, `refund.posted`, `refund.settled`, `refund.returned`, `refund.swept`, `refund.return_swept`

[`account_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-account-id)

stringstring

The account ID associated with the transfer. This field is omitted for Plaid Ledger Sweep events.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-funding-account-id)

nullablestringnullable, string

The id of the associated funding account, available in the Plaid Dashboard. If present, this indicates which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-transfer-id)

stringstring

Plaid's unique identifier for a transfer. This field is an empty string for Plaid Ledger Sweep events.

[`transfer_type`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-transfer-type)

stringstring

The type of transfer. Valid values are `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account. This field is omitted for Plaid Ledger Sweep events.  
  

Possible values: `debit`, `credit`

[`transfer_amount`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-transfer-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). This field is omitted for Plaid Ledger Sweep events.

[`failure_reason`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-failure-reason)

nullableobjectnullable, object

The failure reason if the event type for a transfer is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the transfer status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`ach_return_code`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-failure-reason-ach-return-code)

deprecatednullablestringdeprecated, nullable, string

The ACH return code, e.g. `R01`. A return code will be provided if and only if the transfer status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes).

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-failure-reason-description)

stringstring

A human-readable description of the reason for the failure or reversal.

[`sweep_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-sweep-id)

nullablestringnullable, string

Plaid’s unique identifier for a sweep.

[`sweep_amount`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-sweep-amount)

nullablestringnullable, string

A signed amount of how much was `swept` or `return_swept` for this transfer (decimal string with two digits of precision e.g. "-5.50").

[`refund_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-refund-id)

nullablestringnullable, string

Plaid’s unique identifier for a refund. A non-null value indicates the event is for the associated refund of the transfer.

[`originator_client_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-originator-client-id)

nullablestringnullable, string

The Plaid client ID that is the originator of the transfer that this event applies to. Only present if the transfer was created on behalf of another client as a third-party sender (TPS).

[`intent_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-intent-id)

nullablestringnullable, string

The `id` returned by the /transfer/intent/create endpoint, for transfers created via Transfer UI. For transfers not created by Transfer UI, the value is `null`. This will currently only be populated for RfP transfers.

[`wire_return_fee`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-transfer-events-wire-return-fee)

nullablestringnullable, string

The fee amount deducted from the original transfer during a wire return, if applicable.

[`has_more`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-has-more)

booleanboolean

Whether there are more events to be pulled from the endpoint that have not already been returned

[`request_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transfer_events": [
    {
      "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
      "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
      "transfer_amount": "12.34",
      "transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
      "transfer_type": "credit",
      "event_id": 1,
      "event_type": "posted",
      "failure_reason": null,
      "origination_account_id": "",
      "originator_client_id": "569ed2f36b3a3a021713abc1",
      "refund_id": null,
      "sweep_amount": null,
      "sweep_id": null,
      "timestamp": "2019-12-09T17:27:15Z"
    }
  ],
  "has_more": true,
  "request_id": "mdqfuVxeoza6mhu"
}
```

=\*=\*=\*=

#### `/transfer/event/sync`

#### Sync transfer events

[`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) allows you to request up to the next 500 transfer events that happened after a specific `event_id`. Use the [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) endpoint to guarantee you have seen all transfer events.

/transfer/event/sync

**Request fields**

[`client_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`after_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-request-after-id)

requiredintegerrequired, integer

The latest (largest) `event_id` fetched via the sync endpoint, or 0 initially.  
  

Minimum: `0`

[`count`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-request-count)

integerinteger

The maximum number of transfer events to return.  
  

Default: `100`

Minimum: `1`

Maximum: `500`

/transfer/event/sync

```
const request: TransferEventSyncRequest = {
  after_id: 4,
  count: 22,
};
try {
  const response = await plaidClient.transferEventSync(request);
  const events = response.data.transfer_events;
  for (const event of events) {
    // iterate through events
  }
} catch (error) {
  // handle error
}
```

/transfer/event/sync

**Response fields**

[`transfer_events`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events)

[object][object]

[`event_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-event-id)

integerinteger

Plaid’s unique identifier for this event. IDs are sequential unsigned 64-bit integers.  
  

Minimum: `0`

[`timestamp`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-timestamp)

stringstring

The datetime when this event occurred. This will be of the form `2006-01-02T15:04:05Z`.  
  

Format: `date-time`

[`event_type`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-event-type)

stringstring

The type of event that this transfer represents. Event types with prefix `sweep` represents events for Plaid Ledger sweeps.  
`pending`: A new transfer was created; it is in the pending state.  
`cancelled`: The transfer was cancelled by the client.  
`failed`: The transfer failed, no funds were moved.  
`posted`: The transfer has been successfully submitted to the payment network.  
`settled`: The transfer has been successfully completed by the payment network.  
`funds_available`: Funds from the transfer have been released from hold and applied to the ledger's available balance. (Only applicable to ACH debits.)  
`returned`: A posted transfer was returned.  
`swept`: The transfer was swept to / from the sweep account.  
`swept_settled`: Credits are available to be withdrawn or debits have been deducted from the customer’s business checking account.  
`return_swept`: Due to the transfer being returned, funds were pulled from or pushed back to the sweep account.  
`sweep.pending`: A new ledger sweep was created; it is in the pending state.  
`sweep.posted`: The ledger sweep has been successfully submitted to the payment network.  
`sweep.settled`: The transaction has settled in the funding account. This means that funds withdrawn from Plaid Ledger balance have reached the funding account, or funds to be deposited into the Plaid Ledger Balance have been pulled, and the hold period has begun.  
`sweep.returned`: A posted ledger sweep was returned.  
`sweep.failed`: The ledger sweep failed, no funds were moved.  
`sweep.funds_available`: Funds from the ledger sweep have been released from hold and applied to the ledger's available balance. This is only applicable to debits.  
`refund.pending`: A new refund was created; it is in the pending state.  
`refund.cancelled`: The refund was cancelled.  
`refund.failed`: The refund failed, no funds were moved.  
`refund.posted`: The refund has been successfully submitted to the payment network.  
`refund.settled`: The refund transaction has settled in the Plaid linked account.  
`refund.returned`: A posted refund was returned.  
`refund.swept`: The refund was swept from the sweep account.  
`refund.return_swept`: Due to the refund being returned, funds were pushed back to the sweep account.  
  

Possible values: `pending`, `cancelled`, `failed`, `posted`, `settled`, `funds_available`, `returned`, `swept`, `swept_settled`, `return_swept`, `sweep.pending`, `sweep.posted`, `sweep.settled`, `sweep.returned`, `sweep.failed`, `sweep.funds_available`, `refund.pending`, `refund.cancelled`, `refund.failed`, `refund.posted`, `refund.settled`, `refund.returned`, `refund.swept`, `refund.return_swept`

[`account_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-account-id)

stringstring

The account ID associated with the transfer. This field is omitted for Plaid Ledger Sweep events.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-funding-account-id)

nullablestringnullable, string

The id of the associated funding account, available in the Plaid Dashboard. If present, this indicates which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-transfer-id)

stringstring

Plaid's unique identifier for a transfer. This field is an empty string for Plaid Ledger Sweep events.

[`transfer_type`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-transfer-type)

stringstring

The type of transfer. Valid values are `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account. This field is omitted for Plaid Ledger Sweep events.  
  

Possible values: `debit`, `credit`

[`transfer_amount`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-transfer-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). This field is omitted for Plaid Ledger Sweep events.

[`failure_reason`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-failure-reason)

nullableobjectnullable, object

The failure reason if the event type for a transfer is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the transfer status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`ach_return_code`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-failure-reason-ach-return-code)

deprecatednullablestringdeprecated, nullable, string

The ACH return code, e.g. `R01`. A return code will be provided if and only if the transfer status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes).

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-failure-reason-description)

stringstring

A human-readable description of the reason for the failure or reversal.

[`sweep_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-sweep-id)

nullablestringnullable, string

Plaid’s unique identifier for a sweep.

[`sweep_amount`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-sweep-amount)

nullablestringnullable, string

A signed amount of how much was `swept` or `return_swept` for this transfer (decimal string with two digits of precision e.g. "-5.50").

[`refund_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-refund-id)

nullablestringnullable, string

Plaid’s unique identifier for a refund. A non-null value indicates the event is for the associated refund of the transfer.

[`originator_client_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-originator-client-id)

nullablestringnullable, string

The Plaid client ID that is the originator of the transfer that this event applies to. Only present if the transfer was created on behalf of another client as a third-party sender (TPS).

[`intent_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-intent-id)

nullablestringnullable, string

The `id` returned by the /transfer/intent/create endpoint, for transfers created via Transfer UI. For transfers not created by Transfer UI, the value is `null`. This will currently only be populated for RfP transfers.

[`wire_return_fee`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-transfer-events-wire-return-fee)

nullablestringnullable, string

The fee amount deducted from the original transfer during a wire return, if applicable.

[`has_more`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-has-more)

booleanboolean

Whether there are more events to be pulled from the endpoint that have not already been returned

[`request_id`](/docs/api/products/transfer/reading-transfers/#transfer-event-sync-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transfer_events": [
    {
      "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
      "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
      "transfer_amount": "12.34",
      "transfer_id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
      "transfer_type": "credit",
      "event_id": 1,
      "event_type": "pending",
      "failure_reason": null,
      "origination_account_id": "",
      "originator_client_id": null,
      "refund_id": null,
      "sweep_amount": null,
      "sweep_id": null,
      "timestamp": "2019-12-09T17:27:15Z"
    }
  ],
  "has_more": true,
  "request_id": "mdqfuVxeoza6mhu"
}
```

=\*=\*=\*=

#### `/transfer/sweep/get`

#### Retrieve a sweep

The [`/transfer/sweep/get`](/docs/api/products/transfer/reading-transfers/#transfersweepget) endpoint fetches a sweep corresponding to the given `sweep_id`.

/transfer/sweep/get

**Request fields**

[`client_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`sweep_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-request-sweep-id)

requiredstringrequired, string

Plaid's unique identifier for the sweep (UUID) or a shortened form consisting of the first 8 characters of the identifier (8-digit hexadecimal string).

/transfer/sweep/get

```
const request: TransferSweepGetRequest = {
  sweep_id: '8c2fda9a-aa2f-4735-a00f-f4e0d2d2faee',
};
try {
  const response = await plaidClient.transferSweepGet(request);
  const sweep = response.data.sweep;
} catch (error) {
  // handle error
}
```

/transfer/sweep/get

**Response fields**

[`sweep`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep)

objectobject

Describes a sweep of funds to / from the sweep account.  
A sweep is associated with many sweep events (events of type `swept` or `return_swept`) which can be retrieved by invoking the `/transfer/event/list` endpoint with the corresponding `sweep_id`.  
`swept` events occur when the transfer amount is credited or debited from your sweep account, depending on the `type` of the transfer. `return_swept` events occur when a transfer is returned and Plaid undoes the credit or debit.  
The total sum of the `swept` and `return_swept` events is equal to the `amount` of the sweep Plaid creates and matches the amount of the entry on your sweep account ledger.

[`id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-id)

stringstring

Identifier of the sweep.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`created`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-created)

stringstring

The datetime when the sweep occurred, in RFC 3339 format.  
  

Format: `date-time`

[`amount`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-amount)

stringstring

Signed decimal amount of the sweep as it appears on your sweep account ledger (e.g. "-10.00")  
If amount is not present, the sweep was net-settled to zero and outstanding debits and credits between the sweep account and Plaid are balanced.

[`iso_currency_code`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-iso-currency-code)

stringstring

The currency of the sweep, e.g. "USD".

[`settled`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-settled)

nullablestringnullable, string

The date when the sweep settled, in the YYYY-MM-DD format.  
  

Format: `date`

[`expected_funds_available_date`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-expected-funds-available-date)

nullablestringnullable, string

The expected date when funds from a ledger deposit will be made available and can be withdrawn from the associated ledger balance. Only applies to deposits. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`status`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-status)

nullablestringnullable, string

The status of a sweep transfer  
`"pending"` - The sweep is currently pending
`"posted"` - The sweep has been posted
`"settled"` - The sweep has settled. This is the terminal state of a successful credit sweep.
`"returned"` - The sweep has been returned. This is the terminal state of a returned sweep. Returns of a sweep are extremely rare, since sweeps are money movement between your own bank account and your own Ledger.
`"funds_available"` - Funds from the sweep have been released from hold and applied to the ledger's available balance. (Only applicable to deposits.) This is the terminal state of a successful deposit sweep.
`"failed"` - The sweep has failed. This is the terminal state of a failed sweep.  
  

Possible values: `pending`, `posted`, `settled`, `funds_available`, `returned`, `failed`, `null`

[`trigger`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-trigger)

nullablestringnullable, string

The trigger of the sweep  
`"manual"` - The sweep is created manually by the customer
`"incoming"` - The sweep is created by incoming funds flow (e.g. Incoming Wire)
`"balance_threshold"` - The sweep is created by balance threshold setting
`"automatic_aggregate"` - The sweep is created by the Plaid automatic aggregation process. These funds did not pass through the Plaid Ledger balance.  
  

Possible values: `manual`, `incoming`, `balance_threshold`, `automatic_aggregate`

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-description)

stringstring

The description of the deposit that will be passed to the receiving bank (up to 10 characters). Note that banks utilize this field differently, and may or may not show it on the bank statement.

[`network_trace_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`failure_reason`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-failure-reason)

nullableobjectnullable, object

The failure reason if the status for a sweep is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the sweep status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-sweep-failure-reason-description)

nullablestringnullable, string

A human-readable description of the reason for the failure or reversal.

[`request_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "sweep": {
    "id": "8c2fda9a-aa2f-4735-a00f-f4e0d2d2faee",
    "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
    "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
    "created": "2020-08-06T17:27:15Z",
    "amount": "12.34",
    "iso_currency_code": "USD",
    "settled": "2020-08-07",
    "status": "settled",
    "network_trace_id": "123456789012345"
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/sweep/list`

#### List sweeps

The [`/transfer/sweep/list`](/docs/api/products/transfer/reading-transfers/#transfersweeplist) endpoint fetches sweeps matching the given filters.

/transfer/sweep/list

**Request fields**

[`client_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`start_date`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-start-date)

stringstring

The start `created` datetime of sweeps to return (RFC 3339 format).  
  

Format: `date-time`

[`end_date`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-end-date)

stringstring

The end `created` datetime of sweeps to return (RFC 3339 format).  
  

Format: `date-time`

[`count`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-count)

integerinteger

The maximum number of sweeps to return.  
  

Minimum: `1`

Maximum: `25`

Default: `25`

[`offset`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-offset)

integerinteger

The number of sweeps to skip before returning results.  
  

Default: `0`

Minimum: `0`

[`amount`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-amount)

stringstring

Filter sweeps to only those with the specified amount.

[`status`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-status)

stringstring

The status of a sweep transfer  
`"pending"` - The sweep is currently pending
`"posted"` - The sweep has been posted
`"settled"` - The sweep has settled. This is the terminal state of a successful credit sweep.
`"returned"` - The sweep has been returned. This is the terminal state of a returned sweep. Returns of a sweep are extremely rare, since sweeps are money movement between your own bank account and your own Ledger.
`"funds_available"` - Funds from the sweep have been released from hold and applied to the ledger's available balance. (Only applicable to deposits.) This is the terminal state of a successful deposit sweep.
`"failed"` - The sweep has failed. This is the terminal state of a failed sweep.  
  

Possible values: `pending`, `posted`, `settled`, `funds_available`, `returned`, `failed`, `null`

[`originator_client_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-originator-client-id)

stringstring

Filter sweeps to only those with the specified originator client.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-funding-account-id)

stringstring

Filter sweeps to only those with the specified `funding_account_id`.

[`transfer_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-transfer-id)

stringstring

Filter sweeps to only those with the included `transfer_id`.

[`trigger`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-request-trigger)

stringstring

The trigger of the sweep  
`"manual"` - The sweep is created manually by the customer
`"incoming"` - The sweep is created by incoming funds flow (e.g. Incoming Wire)
`"balance_threshold"` - The sweep is created by balance threshold setting
`"automatic_aggregate"` - The sweep is created by the Plaid automatic aggregation process. These funds did not pass through the Plaid Ledger balance.  
  

Possible values: `manual`, `incoming`, `balance_threshold`, `automatic_aggregate`

/transfer/sweep/list

```
const request: TransferSweepListRequest = {
  start_date: '2019-12-06T22:35:49Z',
  end_date: '2019-12-12T22:35:49Z',
  count: 14,
  offset: 2,
};
try {
  const response = await plaidClient.transferSweepList(request);
  const sweeps = response.data.sweeps;
  for (const sweep of sweeps) {
    // iterate through sweeps
  }
} catch (error) {
  // handle error
}
```

/transfer/sweep/list

**Response fields**

[`sweeps`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps)

[object][object]

[`id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-id)

stringstring

Identifier of the sweep.

[`funding_account_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`created`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-created)

stringstring

The datetime when the sweep occurred, in RFC 3339 format.  
  

Format: `date-time`

[`amount`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-amount)

stringstring

Signed decimal amount of the sweep as it appears on your sweep account ledger (e.g. "-10.00")  
If amount is not present, the sweep was net-settled to zero and outstanding debits and credits between the sweep account and Plaid are balanced.

[`iso_currency_code`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-iso-currency-code)

stringstring

The currency of the sweep, e.g. "USD".

[`settled`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-settled)

nullablestringnullable, string

The date when the sweep settled, in the YYYY-MM-DD format.  
  

Format: `date`

[`expected_funds_available_date`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-expected-funds-available-date)

nullablestringnullable, string

The expected date when funds from a ledger deposit will be made available and can be withdrawn from the associated ledger balance. Only applies to deposits. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`status`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-status)

nullablestringnullable, string

The status of a sweep transfer  
`"pending"` - The sweep is currently pending
`"posted"` - The sweep has been posted
`"settled"` - The sweep has settled. This is the terminal state of a successful credit sweep.
`"returned"` - The sweep has been returned. This is the terminal state of a returned sweep. Returns of a sweep are extremely rare, since sweeps are money movement between your own bank account and your own Ledger.
`"funds_available"` - Funds from the sweep have been released from hold and applied to the ledger's available balance. (Only applicable to deposits.) This is the terminal state of a successful deposit sweep.
`"failed"` - The sweep has failed. This is the terminal state of a failed sweep.  
  

Possible values: `pending`, `posted`, `settled`, `funds_available`, `returned`, `failed`, `null`

[`trigger`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-trigger)

nullablestringnullable, string

The trigger of the sweep  
`"manual"` - The sweep is created manually by the customer
`"incoming"` - The sweep is created by incoming funds flow (e.g. Incoming Wire)
`"balance_threshold"` - The sweep is created by balance threshold setting
`"automatic_aggregate"` - The sweep is created by the Plaid automatic aggregation process. These funds did not pass through the Plaid Ledger balance.  
  

Possible values: `manual`, `incoming`, `balance_threshold`, `automatic_aggregate`

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-description)

stringstring

The description of the deposit that will be passed to the receiving bank (up to 10 characters). Note that banks utilize this field differently, and may or may not show it on the bank statement.

[`network_trace_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`failure_reason`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-failure-reason)

nullableobjectnullable, object

The failure reason if the status for a sweep is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the sweep status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`description`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-sweeps-failure-reason-description)

nullablestringnullable, string

A human-readable description of the reason for the failure or reversal.

[`request_id`](/docs/api/products/transfer/reading-transfers/#transfer-sweep-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "sweeps": [
    {
      "id": "d5394a4d-0b04-4a02-9f4a-7ca5c0f52f9d",
      "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
      "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
      "created": "2019-12-09T17:27:15Z",
      "amount": "-12.34",
      "iso_currency_code": "USD",
      "settled": "2019-12-10",
      "status": "settled",
      "originator_client_id": null
    }
  ],
  "request_id": "saKrIBuEB9qJZno"
}
```

### Webhooks

=\*=\*=\*=

#### `TRANSFER_EVENTS_UPDATE`

Fired when new transfer events are available. Receiving this webhook indicates you should fetch the new events from [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync). If multiple transfer events occur within a single minute, only one webhook will be fired, so a single webhook instance may correspond to multiple transfer events.

**Properties**

[`webhook_type`](/docs/api/products/transfer/reading-transfers/#TransferEventsUpdateWebhook-webhook-type)

stringstring

`TRANSFER`

[`webhook_code`](/docs/api/products/transfer/reading-transfers/#TransferEventsUpdateWebhook-webhook-code)

stringstring

`TRANSFER_EVENTS_UPDATE`

[`environment`](/docs/api/products/transfer/reading-transfers/#TransferEventsUpdateWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "TRANSFER",
  "webhook_code": "TRANSFER_EVENTS_UPDATE",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
