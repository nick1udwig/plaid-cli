---
title: "API - Plaid Ledger | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/transfer/ledger/"
scraped_at: "2026-03-07T22:04:22+00:00"
---

# Plaid Ledger

#### API reference for Plaid Ledger

For how-to guidance, see the [Ledger documentation](/docs/transfer/flow-of-funds/).

| Plaid Ledger |  |
| --- | --- |
| [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit) | Deposit funds into a ledger balance held with Plaid |
| [`/transfer/ledger/distribute`](/docs/api/products/transfer/ledger/#transferledgerdistribute) | Move available balance between platform and its originator |
| [`/transfer/ledger/get`](/docs/api/products/transfer/ledger/#transferledgerget) | Retrieve information about the ledger balance held with Plaid |
| [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw) | Withdraw funds from a ledger balance held with Plaid |
| [`/transfer/ledger/event/list`](/docs/api/products/transfer/ledger/#transferledgereventlist) | Retrieve a list of ledger balance events |

=\*=\*=\*=

#### `/transfer/ledger/deposit`

#### Deposit funds into a Plaid Ledger balance

Use the [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit) endpoint to deposit funds into Plaid Ledger.

/transfer/ledger/deposit

**Request fields**

[`client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`originator_client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-originator-client-id)

stringstring

Client ID of the customer that owns the Ledger balance. This is so Plaid knows which of your customers to payout or collect funds. Only applicable for [Platform customers](https://plaid.com/docs/transfer/application/#originators-vs-platforms). Do not include if you’re paying out to yourself.

[`funding_account_id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-funding-account-id)

stringstring

Specify which funding account to use. Customers can find a list of `funding_account_id`s in the Accounts page of the Plaid Dashboard, under the "Account ID" column. If this field is left blank, the funding account associated with the specified Ledger will be used. If an `originator_client_id` is specified, the `funding_account_id` must belong to the specified originator.

[`ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-ledger-id)

stringstring

Specify which ledger balance to deposit to. Customers can find a list of `ledger_id`s in the Accounts page of your Plaid Dashboard. If this field is left blank, this will default to id of the default ledger balance.

[`amount`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-amount)

requiredstringrequired, string

A positive amount of how much will be deposited into ledger (decimal string with two digits of precision e.g. "5.50").

[`description`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-description)

stringstring

The description of the deposit that will be passed to the receiving bank (up to 10 characters). Note that banks utilize this field differently, and may or may not show it on the bank statement.  
  

Max length: `10`

[`idempotency_key`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-idempotency-key)

requiredstringrequired, string

A unique key provided by the client, per unique ledger deposit. Maximum of 50 characters.  
The API supports idempotency for safely retrying the request without accidentally performing the same operation twice. For example, if a request to create a ledger deposit fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single deposit is created.  
  

Max length: `50`

[`network`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-request-network)

requiredstringrequired, string

The ACH networks used for the funds flow.  
For requests submitted as either `ach` or `same-day-ach` the cutoff for Same Day ACH is 3:00 PM Eastern Time and the cutoff for Standard ACH transfers is 8:30 PM Eastern Time. It is recommended to submit a request at least 15 minutes before the cutoff time in order to ensure that it will be processed before the cutoff. Any request that is indicated as `same-day-ach` and that misses the Same Day ACH cutoff, but is submitted in time for the Standard ACH cutoff, will be sent over Standard ACH rails and will not incur same-day charges.  
  

Possible values: `ach`, `same-day-ach`

/transfer/ledger/deposit

```
const request: TransferLedgerDepositRequest = {
  amount: '12.34',
  network: 'ach',
  idempotency_key: 'test_deposit_abc',
  description: 'deposit',
};
try {
  const response = await client.transferLedgerDeposit(request);
  const sweep = response.data.sweep;
} catch (error) {
  // handle error
}
```

/transfer/ledger/deposit

**Response fields**

[`sweep`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep)

objectobject

Describes a sweep of funds to / from the sweep account.  
A sweep is associated with many sweep events (events of type `swept` or `return_swept`) which can be retrieved by invoking the `/transfer/event/list` endpoint with the corresponding `sweep_id`.  
`swept` events occur when the transfer amount is credited or debited from your sweep account, depending on the `type` of the transfer. `return_swept` events occur when a transfer is returned and Plaid undoes the credit or debit.  
The total sum of the `swept` and `return_swept` events is equal to the `amount` of the sweep Plaid creates and matches the amount of the entry on your sweep account ledger.

[`id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-id)

stringstring

Identifier of the sweep.

[`funding_account_id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`created`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-created)

stringstring

The datetime when the sweep occurred, in RFC 3339 format.  
  

Format: `date-time`

[`amount`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-amount)

stringstring

Signed decimal amount of the sweep as it appears on your sweep account ledger (e.g. "-10.00")  
If amount is not present, the sweep was net-settled to zero and outstanding debits and credits between the sweep account and Plaid are balanced.

[`iso_currency_code`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-iso-currency-code)

stringstring

The currency of the sweep, e.g. "USD".

[`settled`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-settled)

nullablestringnullable, string

The date when the sweep settled, in the YYYY-MM-DD format.  
  

Format: `date`

[`expected_funds_available_date`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-expected-funds-available-date)

nullablestringnullable, string

The expected date when funds from a ledger deposit will be made available and can be withdrawn from the associated ledger balance. Only applies to deposits. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`status`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-status)

nullablestringnullable, string

The status of a sweep transfer  
`"pending"` - The sweep is currently pending
`"posted"` - The sweep has been posted
`"settled"` - The sweep has settled. This is the terminal state of a successful credit sweep.
`"returned"` - The sweep has been returned. This is the terminal state of a returned sweep. Returns of a sweep are extremely rare, since sweeps are money movement between your own bank account and your own Ledger.
`"funds_available"` - Funds from the sweep have been released from hold and applied to the ledger's available balance. (Only applicable to deposits.) This is the terminal state of a successful deposit sweep.
`"failed"` - The sweep has failed. This is the terminal state of a failed sweep.  
  

Possible values: `pending`, `posted`, `settled`, `funds_available`, `returned`, `failed`, `null`

[`trigger`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-trigger)

nullablestringnullable, string

The trigger of the sweep  
`"manual"` - The sweep is created manually by the customer
`"incoming"` - The sweep is created by incoming funds flow (e.g. Incoming Wire)
`"balance_threshold"` - The sweep is created by balance threshold setting
`"automatic_aggregate"` - The sweep is created by the Plaid automatic aggregation process. These funds did not pass through the Plaid Ledger balance.  
  

Possible values: `manual`, `incoming`, `balance_threshold`, `automatic_aggregate`

[`description`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-description)

stringstring

The description of the deposit that will be passed to the receiving bank (up to 10 characters). Note that banks utilize this field differently, and may or may not show it on the bank statement.

[`network_trace_id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`failure_reason`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-failure-reason)

nullableobjectnullable, object

The failure reason if the status for a sweep is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the sweep status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`description`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-sweep-failure-reason-description)

nullablestringnullable, string

A human-readable description of the reason for the failure or reversal.

[`request_id`](/docs/api/products/transfer/ledger/#transfer-ledger-deposit-response-request-id)

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
    "amount": "-12.34",
    "iso_currency_code": "USD",
    "settled": null,
    "status": "pending",
    "trigger": "manual",
    "description": "deposit",
    "network_trace_id": null
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/ledger/distribute`

#### Move available balance between ledgers

Use the [`/transfer/ledger/distribute`](/docs/api/products/transfer/ledger/#transferledgerdistribute) endpoint to move available balance between ledgers, if you have multiple. If you’re a platform, you can move funds between one of your ledgers and one of your customer’s ledger.

/transfer/ledger/distribute

**Request fields**

[`client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-distribute-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/ledger/#transfer-ledger-distribute-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`from_ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-distribute-request-from-ledger-id)

requiredstringrequired, string

The Ledger to pull money from.

[`to_ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-distribute-request-to-ledger-id)

requiredstringrequired, string

The Ledger to credit money to.

[`amount`](/docs/api/products/transfer/ledger/#transfer-ledger-distribute-request-amount)

requiredstringrequired, string

The amount to move (decimal string with two digits of precision e.g. "10.00"). Amount must be positive.

[`idempotency_key`](/docs/api/products/transfer/ledger/#transfer-ledger-distribute-request-idempotency-key)

requiredstringrequired, string

A unique key provided by the client, per unique ledger distribute. Maximum of 50 characters.  
The API supports idempotency for safely retrying the request without accidentally performing the same operation twice. For example, if a request to create a ledger distribute fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single distribute is created.  
  

Max length: `50`

[`description`](/docs/api/products/transfer/ledger/#transfer-ledger-distribute-request-description)

stringstring

An optional description for the ledger distribute operation.

/transfer/ledger/distribute

```
const request: TransferLedgerDistributeRequest = {
   from_client_id: '6a65dh3d1h0d1027121ak184',
   to_client_id: '415ab64b87ec47401d000002',
   amount: '12.34',
   idempotency_key: 'test_distribute_abc',
   description: 'distribute',
};
try {
  const response = await client.transferLedgerDistribute(request);
} catch (error) {
  // handle error
}
```

/transfer/ledger/distribute

**Response fields**

[`request_id`](/docs/api/products/transfer/ledger/#transfer-ledger-distribute-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/ledger/get`

#### Retrieve Plaid Ledger balance

Use the [`/transfer/ledger/get`](/docs/api/products/transfer/ledger/#transferledgerget) endpoint to view a balance on the ledger held with Plaid.

/transfer/ledger/get

**Request fields**

[`client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/ledger/#transfer-ledger-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-get-request-ledger-id)

stringstring

Specify which ledger balance to get. Customers can find a list of `ledger_id`s in the Accounts page of your Plaid Dashboard. If this field is left blank, this will default to id of the default ledger balance.

[`originator_client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-get-request-originator-client-id)

stringstring

Client ID of the end customer.

/transfer/ledger/get

```
try {
  const response = await client.transferLedgerGet({});
  const available_balance = response.data.balance.available;
  const pending_balance = response.data.balance.pending;
} catch (error) {
  // handle error
}
```

/transfer/ledger/get

**Response fields**

[`ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-get-response-ledger-id)

stringstring

The unique identifier of the Ledger that was returned.

[`balance`](/docs/api/products/transfer/ledger/#transfer-ledger-get-response-balance)

objectobject

Information about the balance of the ledger held with Plaid.

[`available`](/docs/api/products/transfer/ledger/#transfer-ledger-get-response-balance-available)

stringstring

The amount of this balance available for use (decimal string with two digits of precision e.g. "10.00").

[`pending`](/docs/api/products/transfer/ledger/#transfer-ledger-get-response-balance-pending)

stringstring

The amount of pending funds that are in processing (decimal string with two digits of precision e.g. "10.00").

[`name`](/docs/api/products/transfer/ledger/#transfer-ledger-get-response-name)

stringstring

The name of the Ledger

[`is_default`](/docs/api/products/transfer/ledger/#transfer-ledger-get-response-is-default)

booleanboolean

Whether this Ledger is the client's default ledger.

[`request_id`](/docs/api/products/transfer/ledger/#transfer-ledger-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
  "name": "Default",
  "is_default": true,
  "balance": {
    "available": "1721.70",
    "pending": "123.45"
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/ledger/withdraw`

#### Withdraw funds from a Plaid Ledger balance

Use the [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw) endpoint to withdraw funds from a Plaid Ledger balance.

/transfer/ledger/withdraw

**Request fields**

[`client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`originator_client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-originator-client-id)

stringstring

Client ID of the customer that owns the Ledger balance. This is so Plaid knows which of your customers to payout or collect funds. Only applicable for [Platform customers](https://plaid.com/docs/transfer/application/#originators-vs-platforms). Do not include if you’re paying out to yourself.

[`funding_account_id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-funding-account-id)

stringstring

Specify which funding account to use. Customers can find a list of `funding_account_id`s in the Accounts page of the Plaid Dashboard, under the "Account ID" column. If this field is left blank, the funding account associated with the specified Ledger will be used. If an `originator_client_id` is specified, the `funding_account_id` must belong to the specified originator.

[`ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-ledger-id)

stringstring

Specify which ledger balance to withdraw from. Customers can find a list of `ledger_id`s in the Accounts page of your Plaid Dashboard. If this field is left blank, this will default to id of the default ledger balance.

[`amount`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-amount)

requiredstringrequired, string

A positive amount of how much will be withdrawn from the ledger balance (decimal string with two digits of precision e.g. "5.50").

[`description`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-description)

stringstring

The description of the deposit that will be passed to the receiving bank (up to 10 characters). Note that banks utilize this field differently, and may or may not show it on the bank statement.  
  

Max length: `10`

[`idempotency_key`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-idempotency-key)

requiredstringrequired, string

A unique key provided by the client, per unique ledger withdraw. Maximum of 50 characters.  
The API supports idempotency for safely retrying the request without accidentally performing the same operation twice. For example, if a request to create a ledger withdraw fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single withdraw is created.  
  

Max length: `50`

[`network`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-request-network)

requiredstringrequired, string

The network or rails used for the transfer.  
For transfers submitted as `ach` or `same-day-ach`, the Standard ACH cutoff is 8:30 PM Eastern Time.  
For transfers submitted as `same-day-ach`, the Same Day ACH cutoff is 3:00 PM Eastern Time. It is recommended to send the request 15 minutes prior to the cutoff to ensure that it will be processed in time for submission before the cutoff. If the transfer is processed after this cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur same-day charges; this will apply to both legs of the transfer if applicable. The transaction limit for a Same Day ACH transfer is $1,000,000. Authorization requests sent with an amount greater than $1,000,000 will fail.  
For transfers submitted as `rtp`, Plaid will automatically route between Real Time Payment rail by TCH or FedNow rails as necessary. If a transfer is submitted as `rtp` and the counterparty account is not eligible for RTP, the `/transfer/authorization/create` request will fail with an `INVALID_FIELD` error code. To pre-check to determine whether a counterparty account can support RTP, call `/transfer/capabilities/get` before calling `/transfer/authorization/create`.  
Wire transfers are currently in early availability. To request access to `wire` as a payment network, contact your Account Manager. For transfers submitted as `wire`, the `type` must be `credit`; wire debits are not supported. The cutoff to submit a wire payment is 6:30 PM Eastern Time on a business day; wires submitted after that time will be processed on the next business day. The transaction limit for a wire is $999,999.99. Authorization requests sent with an amount greater than $999,999.99 will fail.  
  

Possible values: `ach`, `same-day-ach`, `rtp`, `wire`

/transfer/ledger/withdraw

```
const request: TransferLedgerWithdrawRequest = {
  amount: '12.34',
  network: 'ach',
  idempotency_key: 'test_deposit_abc',
  description: 'withdraw',
};
try {
  const response = await client.transferLedgerWithdraw(request);
  const sweep = response.data.sweep;
} catch (error) {
  // handle error
}
```

/transfer/ledger/withdraw

**Response fields**

[`sweep`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep)

objectobject

Describes a sweep of funds to / from the sweep account.  
A sweep is associated with many sweep events (events of type `swept` or `return_swept`) which can be retrieved by invoking the `/transfer/event/list` endpoint with the corresponding `sweep_id`.  
`swept` events occur when the transfer amount is credited or debited from your sweep account, depending on the `type` of the transfer. `return_swept` events occur when a transfer is returned and Plaid undoes the credit or debit.  
The total sum of the `swept` and `return_swept` events is equal to the `amount` of the sweep Plaid creates and matches the amount of the entry on your sweep account ledger.

[`id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-id)

stringstring

Identifier of the sweep.

[`funding_account_id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-funding-account-id)

stringstring

The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`created`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-created)

stringstring

The datetime when the sweep occurred, in RFC 3339 format.  
  

Format: `date-time`

[`amount`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-amount)

stringstring

Signed decimal amount of the sweep as it appears on your sweep account ledger (e.g. "-10.00")  
If amount is not present, the sweep was net-settled to zero and outstanding debits and credits between the sweep account and Plaid are balanced.

[`iso_currency_code`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-iso-currency-code)

stringstring

The currency of the sweep, e.g. "USD".

[`settled`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-settled)

nullablestringnullable, string

The date when the sweep settled, in the YYYY-MM-DD format.  
  

Format: `date`

[`expected_funds_available_date`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-expected-funds-available-date)

nullablestringnullable, string

The expected date when funds from a ledger deposit will be made available and can be withdrawn from the associated ledger balance. Only applies to deposits. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`status`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-status)

nullablestringnullable, string

The status of a sweep transfer  
`"pending"` - The sweep is currently pending
`"posted"` - The sweep has been posted
`"settled"` - The sweep has settled. This is the terminal state of a successful credit sweep.
`"returned"` - The sweep has been returned. This is the terminal state of a returned sweep. Returns of a sweep are extremely rare, since sweeps are money movement between your own bank account and your own Ledger.
`"funds_available"` - Funds from the sweep have been released from hold and applied to the ledger's available balance. (Only applicable to deposits.) This is the terminal state of a successful deposit sweep.
`"failed"` - The sweep has failed. This is the terminal state of a failed sweep.  
  

Possible values: `pending`, `posted`, `settled`, `funds_available`, `returned`, `failed`, `null`

[`trigger`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-trigger)

nullablestringnullable, string

The trigger of the sweep  
`"manual"` - The sweep is created manually by the customer
`"incoming"` - The sweep is created by incoming funds flow (e.g. Incoming Wire)
`"balance_threshold"` - The sweep is created by balance threshold setting
`"automatic_aggregate"` - The sweep is created by the Plaid automatic aggregation process. These funds did not pass through the Plaid Ledger balance.  
  

Possible values: `manual`, `incoming`, `balance_threshold`, `automatic_aggregate`

[`description`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-description)

stringstring

The description of the deposit that will be passed to the receiving bank (up to 10 characters). Note that banks utilize this field differently, and may or may not show it on the bank statement.

[`network_trace_id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`failure_reason`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-failure-reason)

nullableobjectnullable, object

The failure reason if the status for a sweep is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the sweep status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`description`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-sweep-failure-reason-description)

nullablestringnullable, string

A human-readable description of the reason for the failure or reversal.

[`request_id`](/docs/api/products/transfer/ledger/#transfer-ledger-withdraw-response-request-id)

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
    "settled": null,
    "status": "pending",
    "trigger": "manual",
    "description": "withdraw",
    "network_trace_id": null
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/ledger/event/list`

#### List transfer ledger events

Use the [`/transfer/ledger/event/list`](/docs/api/products/transfer/ledger/#transferledgereventlist) endpoint to get a list of ledger events for a specific ledger based on specified filter criteria.

/transfer/ledger/event/list

**Request fields**

[`client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`originator_client_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-originator-client-id)

stringstring

Filter transfer events to only those with the specified originator client. (This field is specifically for resellers. Caller's client ID will be used if this field is not specified.)

[`secret`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`start_date`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-start-date)

stringstring

The start created datetime of transfers to list. This should be in RFC 3339 format (i.e. 2019-12-06T22:35:49Z)  
  

Format: `date-time`

[`end_date`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-end-date)

stringstring

The end created datetime of transfers to list. This should be in RFC 3339 format (i.e. 2019-12-06T22:35:49Z)  
  

Format: `date-time`

[`ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-ledger-id)

stringstring

Plaid's unique identifier for a Plaid Ledger Balance.

[`ledger_event_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-ledger-event-id)

stringstring

Plaid's unique identifier for the ledger event.

[`source_type`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-source-type)

stringstring

Source of the ledger event.  
`"TRANSFER"` - The source of the ledger event is a transfer
`"SWEEP"` - The source of the ledger event is a sweep
`"REFUND"` - The source of the ledger event is a refund  
  

Possible values: `TRANSFER`, `SWEEP`, `REFUND`

[`source_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-source-id)

stringstring

Plaid's unique identifier for a transfer, sweep, or refund.

[`count`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-count)

integerinteger

The maximum number of transfer events to return. If the number of events matching the above parameters is greater than `count`, the most recent events will be returned.  
  

Default: `25`

Maximum: `25`

Minimum: `1`

[`offset`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-request-offset)

integerinteger

The offset into the list of transfer events. When `count`=25 and `offset`=0, the first 25 events will be returned. When `count`=25 and `offset`=25, the next 25 events will be returned.  
  

Default: `0`

Minimum: `0`

/transfer/ledger/event/list

```
const request: TransferLedgerEventListRequest = {
    originator_client_id: "8945fedc-e703-463d-86b1-dc0607b55460",
    start_date: '2019-12-06T22:35:49Z',
    end_date: '2019-12-12T22:35:49Z',
    count: 14,
    offset: 2,
    ledger_id: "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
    ledger_event_id: "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
    source_type: "transfer",
    source_id: "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
 };
try {
  const response = await plaidClient.transferLedgerEventList(request);
  const events = response.data.ledger_events;
  for (const event of events) {
    // iterate through events
  }
} catch (error) {
  // handle error
}
```

/transfer/ledger/event/list

**Response fields**

[`ledger_events`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events)

[object][object]

[`ledger_event_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-ledger-event-id)

stringstring

Plaid's unique identifier for this ledger event.

[`ledger_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-ledger-id)

stringstring

The ID of the ledger this event belongs to.

[`amount`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-amount)

stringstring

The amount of the ledger event as a decimal string.

[`transfer_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-transfer-id)

nullablestringnullable, string

The ID of the transfer source that triggered this ledger event.

[`refund_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-refund-id)

nullablestringnullable, string

The ID of the refund source that triggered this ledger event.

[`sweep_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-sweep-id)

nullablestringnullable, string

The ID of the sweep source that triggered this ledger event.

[`description`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-description)

stringstring

A description of the ledger event.

[`pending_balance`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-pending-balance)

stringstring

The new pending balance after this event.

[`available_balance`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-available-balance)

stringstring

The new available balance after this event.

[`type`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-type)

stringstring

The type of balance that was impacted by this event.

[`timestamp`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-ledger-events-timestamp)

stringstring

The datetime when this ledger event occurred.  
  

Format: `date-time`

[`has_more`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-has-more)

booleanboolean

Whether there are more events to be pulled from the endpoint that have not already been returned

[`request_id`](/docs/api/products/transfer/ledger/#transfer-ledger-event-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "ledger_events": [
    {
      "ledger_event_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
      "amount": "100.00",
      "type": "deposit",
      "transfer_id": "460cbe92-2dcc-8eae",
      "description": "Converted to available",
      "pending_balance": "100.00",
      "available_balance": "100.00",
      "timestamp": "2023-12-01T10:00:00Z"
    }
  ],
  "has_more": false,
  "request_id": "mdqfuVxeoza6mhu"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
