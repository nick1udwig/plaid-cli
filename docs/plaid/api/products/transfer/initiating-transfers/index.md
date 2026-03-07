---
title: "API - Initiating Transfers | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/transfer/initiating-transfers/"
scraped_at: "2026-03-07T22:04:22+00:00"
---

# Initiating transfers

#### API reference for Transfer initiation endpoints

For how-to guidance, see the [Transfer creation documentation](/docs/transfer/creating-transfers/).

| Initiating Transfers |  |
| --- | --- |
| [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) | Create a transfer authorization |
| [`/transfer/authorization/cancel`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcancel) | Cancel a transfer authorization |
| [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) | Create a transfer |
| [`/transfer/cancel`](/docs/api/products/transfer/initiating-transfers/#transfercancel) | Cancel a transfer |

=\*=\*=\*=

#### `/transfer/authorization/create`

#### Create a transfer authorization

Use the [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) endpoint to authorize a transfer. This endpoint must be called prior to calling [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate). The transfer authorization will expire if not used after one hour. (You can contact your account manager to change the default authorization lifetime.)

There are four possible outcomes to calling this endpoint:

- If the `authorization.decision` in the response is `declined`, the proposed transfer has failed the risk check and you cannot proceed with the transfer.

- If the `authorization.decision` is `user_action_required`, additional user input is needed, usually to fix a broken bank connection, before Plaid can properly assess the risk. You need to launch Link in update mode to complete the required user action. When calling [`/link/token/create`](/docs/api/link/#linktokencreate) to get a new Link token, instead of providing `access_token` in the request, you should set [`transfer.authorization_id`](https://plaid.com/docs/api/link/#link-token-create-request-transfer-authorization-id) as the `authorization.id`. After the Link flow is completed, you may re-attempt the authorization.

- If the `authorization.decision` is `approved`, and the `authorization.rationale_code` is `null`, the transfer has passed the risk check and you can proceed to call [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate).

- If the `authorization.decision` is `approved` and the `authorization.rationale_code` is non-`null`, the risk check could not be run: you may proceed with the transfer, but should perform your own risk evaluation. For more details, see the response schema.

In Plaid's Sandbox environment the decisions will be returned as follows:

- To approve a transfer with `null` rationale code, make an authorization request with an `amount` less than the available balance in the account.

- To approve a transfer with the rationale code `MANUALLY_VERIFIED_ITEM`, create an Item in Link through the [Same Day Micro-deposits flow](https://plaid.com/docs/auth/coverage/testing/#testing-same-day-micro-deposits).

- To get an authorization decision of `user_action_required`, [reset the login for an Item](https://plaid.com/docs/sandbox/#item_login_required).

- To decline a transfer with the rationale code `NSF`, the available balance on the account must be less than the authorization `amount`. See [Create Sandbox test data](https://plaid.com/docs/sandbox/user-custom/) for details on how to customize data in Sandbox.

- To decline a transfer with the rationale code `RISK`, the available balance on the account must be exactly $0. See [Create Sandbox test data](https://plaid.com/docs/sandbox/user-custom/) for details on how to customize data in Sandbox.

/transfer/authorization/create

**Request fields**

[`client_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`access_token`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-access-token)

requiredstringrequired, string

The Plaid `access_token` for the account that will be debited or credited.

[`account_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-account-id)

requiredstringrequired, string

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`ledger_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-ledger-id)

stringstring

Specify which ledger balance should be used to fund the transfer. You can find a list of `ledger_id`s in the Accounts page of your Plaid Dashboard. If this field is left blank, this will default to id of the default ledger balance.

[`type`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-type)

requiredstringrequired, string

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`network`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-network)

requiredstringrequired, string

The network or rails used for the transfer.  
For transfers submitted as `ach` or `same-day-ach`, the Standard ACH cutoff is 8:30 PM Eastern Time.  
For transfers submitted as `same-day-ach`, the Same Day ACH cutoff is 3:00 PM Eastern Time. It is recommended to send the request 15 minutes prior to the cutoff to ensure that it will be processed in time for submission before the cutoff. If the transfer is processed after this cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur same-day charges; this will apply to both legs of the transfer if applicable. The transaction limit for a Same Day ACH transfer is $1,000,000. Authorization requests sent with an amount greater than $1,000,000 will fail.  
For transfers submitted as `rtp`, Plaid will automatically route between Real Time Payment rail by TCH or FedNow rails as necessary. If a transfer is submitted as `rtp` and the counterparty account is not eligible for RTP, the `/transfer/authorization/create` request will fail with an `INVALID_FIELD` error code. To pre-check to determine whether a counterparty account can support RTP, call `/transfer/capabilities/get` before calling `/transfer/authorization/create`.  
Wire transfers are currently in early availability. To request access to `wire` as a payment network, contact your Account Manager. For transfers submitted as `wire`, the `type` must be `credit`; wire debits are not supported. The cutoff to submit a wire payment is 6:30 PM Eastern Time on a business day; wires submitted after that time will be processed on the next business day. The transaction limit for a wire is $999,999.99. Authorization requests sent with an amount greater than $999,999.99 will fail.  
  

Possible values: `ach`, `same-day-ach`, `rtp`, `wire`

[`amount`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-amount)

requiredstringrequired, string

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`ach_class`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`wire_details`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-wire-details)

objectobject

Information specific to wire transfers.

[`message_to_beneficiary`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-wire-details-message-to-beneficiary)

stringstring

Additional information from the wire originator to the beneficiary. Max 140 characters.

[`wire_return_fee`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-wire-details-wire-return-fee)

stringstring

The fee amount deducted from the original transfer during a wire return, if applicable.

[`user`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user)

requiredobjectrequired, object

The legal name and other information for the account holder. If the account has multiple account holders, provide the information for the account holder on whose behalf the authorization is being requested. The `user.legal_name` field is required. Other fields are not currently used and are present to support planned future functionality.

[`legal_name`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-legal-name)

requiredstringrequired, string

The user's legal name. If the user is a business, provide the business name.

[`phone_number`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-phone-number)

stringstring

The user's phone number.

[`email_address`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-email-address)

stringstring

The user's email address.

[`address`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-address)

objectobject

The address associated with the account holder.

[`street`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-address-street)

stringstring

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-address-city)

stringstring

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-address-region)

stringstring

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-address-postal-code)

stringstring

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-address-country)

stringstring

A two-letter country code (e.g., "US").

[`device`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-device)

objectobject

Information about the device being used to initiate the authorization. These fields are not currently incorporated into the risk check.

[`ip_address`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-device-ip-address)

stringstring

The IP address of the device being used to initiate the authorization.

[`user_agent`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-device-user-agent)

stringstring

The user agent of the device being used to initiate the authorization.

[`iso_currency_code`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-iso-currency-code)

stringstring

The currency of the transfer amount. The default value is "USD".

[`idempotency_key`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-idempotency-key)

stringstring

A random key provided by the client, per unique authorization, which expires after 48 hours. Maximum of 50 characters.  
The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create an authorization fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single authorization is created.  
Idempotency does not apply to authorizations whose decisions are `user_action_required`. Therefore you may re-attempt the authorization after completing the required user action without changing `idempotency_key`.  
This idempotency key expires after 48 hours, after which the same key can be reused. Failure to provide this key may result in duplicate charges.  
  

Max length: `50`

[`user_present`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-user-present)

booleanboolean

If the end user is initiating the specific transfer themselves via an interactive UI, this should be `true`; for automatic recurring payments where the end user is not actually initiating each individual transfer, it should be `false`. This field is not currently used and is present to support planned future functionality.

[`originator_client_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-originator-client-id)

stringstring

The Plaid client ID that is the originator of this transfer. Only needed if creating transfers on behalf of another client as a [Platform customer](https://plaid.com/docs/transfer/application/#originators-vs-platforms).

[`test_clock_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-test-clock-id)

stringstring

Plaid’s unique identifier for a test clock. This field may only be used when using `sandbox` environment. If provided, the `authorization` is created at the `virtual_time` on the provided test clock.

[`ruleset_key`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-request-ruleset-key)

stringstring

The key of the Ruleset for the transaction. This feature is currently in closed beta; to request access, contact your account manager.

/transfer/authorization/create

```
const request: TransferAuthorizationCreateRequest = {
  access_token: 'access-sandbox-71e02f71-0960-4a27-abd2-5631e04f2175',
  account_id: '3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr',
  type: 'debit',
  network: 'ach',
  amount: '12.34',
  ach_class: 'ppd',
  user: {
    legal_name: 'Anne Charleston',
  },
};

try {
  const response = await client.transferAuthorizationCreate(request);
  const authorizationId = response.data.authorization.id;
} catch (error) {
  // handle error
}
```

/transfer/authorization/create

**Response fields**

[`authorization`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization)

objectobject

Contains the authorization decision for a proposed transfer.

[`id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-id)

stringstring

Plaid’s unique identifier for a transfer authorization.

[`created`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-created)

stringstring

The datetime representing when the authorization was created, in the format `2006-01-02T15:04:05Z`.  
  

Format: `date-time`

[`decision`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-decision)

stringstring

A decision regarding the proposed transfer.  
`approved` – The proposed transfer has received the end user's consent and has been approved for processing by Plaid. The `decision_rationale` field is set if Plaid was unable to fetch the account information. You may proceed with the transfer, but further review is recommended. Refer to the `code` field in the `decision_rationale` object for details.  
`declined` – Plaid reviewed the proposed transfer and declined processing. Refer to the `code` field in the `decision_rationale` object for details.  
`user_action_required` – An action is required before Plaid can assess the transfer risk and make a decision. The most common scenario is to update authentication for an Item. To complete the required action, initialize Link by setting `transfer.authorization_id` in the request of `/link/token/create`. After Link flow is completed, you may re-attempt the authorization request.  
For `guarantee` requests, `approved` indicates the transfer is eligible for Plaid's guarantee, and `declined` indicates Plaid will not provide guarantee coverage for the transfer. `user_action_required` indicates you should follow the above guidance before re-attempting.  
  

Possible values: `approved`, `declined`, `user_action_required`

[`decision_rationale`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-decision-rationale)

nullableobjectnullable, object

The rationale for Plaid's decision regarding a proposed transfer. It is always set for `declined` decisions, and may or may not be null for `approved` decisions.

[`code`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-decision-rationale-code)

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

[`description`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-decision-rationale-description)

stringstring

A human-readable description of the code associated with a transfer approval or transfer decline.

[`proposed_transfer`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer)

objectobject

Details regarding the proposed transfer.

[`ach_class`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`account_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-account-id)

stringstring

The Plaid `account_id` for the account that will be debited or credited.

[`funding_account_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-funding-account-id)

nullablestringnullable, string

The id of the associated funding account, available in the Plaid Dashboard. If present, this indicates which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`type`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-type)

stringstring

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`user`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`amount`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`network`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-network)

stringstring

The network or rails used for the transfer.

[`wire_details`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-wire-details)

nullableobjectnullable, object

Information specific to wire transfers.

[`message_to_beneficiary`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-wire-details-message-to-beneficiary)

nullablestringnullable, string

Additional information from the wire originator to the beneficiary. Max 140 characters.

[`wire_return_fee`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-wire-details-wire-return-fee)

nullablestringnullable, string

The fee amount deducted from the original transfer during a wire return, if applicable.

[`iso_currency_code`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-iso-currency-code)

stringstring

The currency of the transfer amount. The default value is "USD".

[`originator_client_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-originator-client-id)

nullablestringnullable, string

The Plaid client ID that is the originator of this transfer. Only present if created on behalf of another client as a [Platform customer](https://plaid.com/docs/transfer/application/#originators-vs-platforms).

[`credit_funds_source`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-authorization-proposed-transfer-credit-funds-source)

deprecatednullablestringdeprecated, nullable, string

This field is now deprecated. You may ignore it for transfers created on and after 12/01/2023.  
Specifies the source of funds for the transfer. Only valid for `credit` transfers, and defaults to `sweep` if not specified. This field is not specified for `debit` transfers.  
`sweep` - Sweep funds from your funding account
`prefunded_rtp_credits` - Use your prefunded RTP credit balance with Plaid
`prefunded_ach_credits` - Use your prefunded ACH credit balance with Plaid  
  

Possible values: `sweep`, `prefunded_rtp_credits`, `prefunded_ach_credits`, `null`

[`request_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "authorization": {
    "id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
    "created": "2020-08-06T17:27:15Z",
    "decision": "approved",
    "decision_rationale": null,
    "guarantee_decision": null,
    "guarantee_decision_rationale": null,
    "payment_risk": null,
    "proposed_transfer": {
      "ach_class": "ppd",
      "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
      "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
      "type": "credit",
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
      "amount": "12.34",
      "network": "ach",
      "iso_currency_code": "USD",
      "origination_account_id": "",
      "originator_client_id": null,
      "credit_funds_source": "sweep"
    }
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/authorization/cancel`

#### Cancel a transfer authorization

Use the [`/transfer/authorization/cancel`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcancel) endpoint to cancel a transfer authorization. A transfer authorization is eligible for cancellation if it has not yet been used to create a transfer.

/transfer/authorization/cancel

**Request fields**

[`client_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-cancel-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-cancel-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`authorization_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-cancel-request-authorization-id)

requiredstringrequired, string

Plaid’s unique identifier for a transfer authorization.

/transfer/authorization/cancel

```
const request: TransferAuthorizationCancelRequest = {
  authorization_id: '123004561178933',
};
try {
  const response = await plaidClient.transferAuthorizationCancel(request);
  const request_id = response.data.request_id;
} catch (error) {
  // handle error
}
```

/transfer/authorization/cancel

**Response fields**

[`request_id`](/docs/api/products/transfer/initiating-transfers/#transfer-authorization-cancel-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/create`

#### Create a transfer

Use the [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) endpoint to initiate a new transfer. This endpoint is retryable and idempotent; if a transfer with the provided `transfer_id` has already been created, it will return the transfer details without creating a new transfer. A transfer may still be created if a 500 error is returned; to detect this scenario, use [Transfer events](https://plaid.com/docs/transfer/reconciling-transfers/).

/transfer/create

**Request fields**

[`client_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`access_token`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-access-token)

requiredstringrequired, string

The Plaid `access_token` for the account that will be debited or credited.

[`account_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-account-id)

requiredstringrequired, string

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`authorization_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-authorization-id)

requiredstringrequired, string

Plaid’s unique identifier for a transfer authorization. This parameter also serves the purpose of acting as an idempotency identifier.

[`amount`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`description`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-description)

requiredstringrequired, string

The transfer description, maximum of 15 characters (RTP transactions) or 10 characters (ACH transactions). Should represent why the money is moving, not your company name. For recommendations on setting the `description` field to avoid ACH returns, see [Description field recommendations](https://www.plaid.com/docs/transfer/creating-transfers/#description-field-recommendations).   
If reprocessing a returned transfer, the `description` field must be `"Retry 1"` or `"Retry 2"`. You may retry a transfer up to 2 times, within 180 days of creating the original transfer. Only transfers that were returned with code `R01` or `R09` may be retried.  
  

Max length: `15`

[`metadata`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-metadata)

objectobject

The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply:
The JSON values must be Strings (no nested JSON objects allowed)
Only ASCII characters may be used
Maximum of 50 key/value pairs
Maximum key length of 40 characters
Maximum value length of 500 characters

[`test_clock_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-test-clock-id)

stringstring

Plaid’s unique identifier for a test clock. This field may only be used when using `sandbox` environment. If provided, the `transfer` is created at the `virtual_time` on the provided `test_clock`.

[`facilitator_fee`](/docs/api/products/transfer/initiating-transfers/#transfer-create-request-facilitator-fee)

stringstring

The amount to deduct from `transfer.amount` and distribute to the platform’s Ledger balance as a facilitator fee (decimal string with two digits of precision e.g. "10.00"). The remainder will go to the end-customer’s Ledger balance. This must be value greater than 0 and less than or equal to the `transfer.amount`.

/transfer/create

```
const request: TransferCreateRequest = {
  access_token: 'access-sandbox-71e02f71-0960-4a27-abd2-5631e04f2175',
  account_id: '3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr',
  description: 'payment',
  authorization_id: '231h012308h3101z21909sw',
};
try {
  const response = await client.transferCreate(request);
  const transfer = response.data.transfer;
} catch (error) {
  // handle error
}
```

/transfer/create

**Response fields**

[`transfer`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer)

objectobject

Represents a transfer within the Transfers API.

[`id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-id)

stringstring

Plaid’s unique identifier for a transfer.

[`authorization_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-authorization-id)

stringstring

Plaid’s unique identifier for a transfer authorization.

[`ach_class`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-ach-class)

stringstring

Specifies the use case of the transfer. Required for transfers on an ACH network. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes).  
Codes supported for credits: `ccd`, `ppd`
Codes supported for debits: `ccd`, `tel`, `web`  
`"ccd"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  
`"ppd"` - Prearranged Payment or Deposit - The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained in writing either in person or via an electronic document signing, e.g. Docusign, by the consumer. Can be used for credits or debits.  
`"web"` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.  
`"tel"` - Telephone-Initiated Entry. The transfer debits a consumer. Debit authorization has been received orally over the telephone via a recorded call.  
  

Possible values: `ccd`, `ppd`, `tel`, `web`

[`account_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-account-id)

stringstring

The Plaid `account_id` corresponding to the end-user account that will be debited or credited.

[`funding_account_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-funding-account-id)

nullablestringnullable, string

The id of the associated funding account, available in the Plaid Dashboard. If present, this indicates which of your business checking accounts will be credited or debited.

[`ledger_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`type`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-type)

stringstring

The type of transfer. This will be either `debit` or `credit`. A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.  
  

Possible values: `debit`, `credit`

[`user`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user)

objectobject

The legal name and other information for the account holder.

[`legal_name`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-legal-name)

stringstring

The user's legal name.

[`phone_number`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-phone-number)

nullablestringnullable, string

The user's phone number.

[`email_address`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-email-address)

nullablestringnullable, string

The user's email address.

[`address`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-address)

nullableobjectnullable, object

The address associated with the account holder.

[`street`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-address-street)

nullablestringnullable, string

The street number and name (i.e., "100 Market St.").

[`city`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-address-city)

nullablestringnullable, string

Ex. "San Francisco"

[`region`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-address-region)

nullablestringnullable, string

The state or province (e.g., "CA").

[`postal_code`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-address-postal-code)

nullablestringnullable, string

The postal code (e.g., "94103").

[`country`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-user-address-country)

nullablestringnullable, string

A two-letter country code (e.g., "US").

[`amount`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-amount)

stringstring

The amount of the transfer (decimal string with two digits of precision e.g. "10.00"). When calling `/transfer/authorization/create`, specify the maximum amount to authorize. When calling `/transfer/create`, specify the exact amount of the transfer, up to a maximum of the amount authorized. If this field is left blank when calling `/transfer/create`, the maximum amount authorized in the `authorization_id` will be sent.

[`description`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-description)

stringstring

The description of the transfer.

[`created`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-created)

stringstring

The datetime when this transfer was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`status`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-status)

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

[`sweep_status`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-sweep-status)

nullablestringnullable, string

The status of the sweep for the transfer.  
`unswept`: The transfer hasn't been swept yet.
`swept`: The transfer was swept to the sweep account.
`swept_settled`: Credits are available to be withdrawn or debits have been deducted from the customer’s business checking account.
`return_swept`: The transfer was returned, funds were pulled back or pushed back to the sweep account.
`null`: The transfer will never be swept (e.g. if the transfer is cancelled or returned before being swept)  
  

Possible values: `null`, `unswept`, `swept`, `swept_settled`, `return_swept`

[`network`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-network)

stringstring

The network or rails used for the transfer.  
For transfers submitted as `ach` or `same-day-ach`, the Standard ACH cutoff is 8:30 PM Eastern Time.  
For transfers submitted as `same-day-ach`, the Same Day ACH cutoff is 3:00 PM Eastern Time. It is recommended to send the request 15 minutes prior to the cutoff to ensure that it will be processed in time for submission before the cutoff. If the transfer is processed after this cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur same-day charges; this will apply to both legs of the transfer if applicable. The transaction limit for a Same Day ACH transfer is $1,000,000. Authorization requests sent with an amount greater than $1,000,000 will fail.  
For transfers submitted as `rtp`, Plaid will automatically route between Real Time Payment rail by TCH or FedNow rails as necessary. If a transfer is submitted as `rtp` and the counterparty account is not eligible for RTP, the `/transfer/authorization/create` request will fail with an `INVALID_FIELD` error code. To pre-check to determine whether a counterparty account can support RTP, call `/transfer/capabilities/get` before calling `/transfer/authorization/create`.  
Wire transfers are currently in early availability. To request access to `wire` as a payment network, contact your Account Manager. For transfers submitted as `wire`, the `type` must be `credit`; wire debits are not supported. The cutoff to submit a wire payment is 6:30 PM Eastern Time on a business day; wires submitted after that time will be processed on the next business day. The transaction limit for a wire is $999,999.99. Authorization requests sent with an amount greater than $999,999.99 will fail.  
  

Possible values: `ach`, `same-day-ach`, `rtp`, `wire`

[`wire_details`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-wire-details)

nullableobjectnullable, object

Information specific to wire transfers.

[`message_to_beneficiary`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-wire-details-message-to-beneficiary)

nullablestringnullable, string

Additional information from the wire originator to the beneficiary. Max 140 characters.

[`wire_return_fee`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-wire-details-wire-return-fee)

nullablestringnullable, string

The fee amount deducted from the original transfer during a wire return, if applicable.

[`cancellable`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-cancellable)

booleanboolean

When `true`, you can still cancel this transfer.

[`failure_reason`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-failure-reason)

nullableobjectnullable, object

The failure reason if the event type for a transfer is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the transfer status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`ach_return_code`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-failure-reason-ach-return-code)

deprecatednullablestringdeprecated, nullable, string

The ACH return code, e.g. `R01`. A return code will be provided if and only if the transfer status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes).

[`description`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-failure-reason-description)

stringstring

A human-readable description of the reason for the failure or reversal.

[`metadata`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-metadata)

nullableobjectnullable, object

The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply:
The JSON values must be Strings (no nested JSON objects allowed)
Only ASCII characters may be used
Maximum of 50 key/value pairs
Maximum key length of 40 characters
Maximum value length of 500 characters

[`iso_currency_code`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-iso-currency-code)

stringstring

The currency of the transfer amount, e.g. "USD"

[`standard_return_window`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-standard-return-window)

nullablestringnullable, string

The date 3 business days from settlement date indicating the following ACH returns can no longer happen: R01, R02, R03, R29. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`unauthorized_return_window`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-unauthorized-return-window)

nullablestringnullable, string

The date 61 business days from settlement date indicating the following ACH returns can no longer happen: R05, R07, R10, R11, R51, R33, R37, R38, R51, R52, R53. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`expected_settlement_date`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-expected-settlement-date)

deprecatednullablestringdeprecated, nullable, string

Deprecated for Plaid Ledger clients, use `expected_funds_available_date` instead.  
  

Format: `date`

[`expected_funds_available_date`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-expected-funds-available-date)

nullablestringnullable, string

The expected date when funds from a transfer will be made available and can be withdrawn from the associated ledger balance, assuming the debit does not return before this date. If the transfer does return before this date, this field will be null. Only applies to debit transfers. This will be of the form YYYY-MM-DD.  
  

Format: `date`

[`originator_client_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-originator-client-id)

nullablestringnullable, string

The Plaid client ID that is the originator of this transfer. Only present if created on behalf of another client as a [Platform customer](https://plaid.com/docs/transfer/application/#originators-vs-platforms).

[`refunds`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds)

[object][object]

A list of refunds associated with this transfer.

[`id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-id)

stringstring

Plaid’s unique identifier for a refund.

[`transfer_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-transfer-id)

stringstring

The ID of the transfer to refund.

[`amount`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-amount)

stringstring

The amount of the refund (decimal string with two digits of precision e.g. "10.00").

[`status`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-status)

stringstring

The status of the refund.  
`pending`: A new refund was created; it is in the pending state.
`posted`: The refund has been successfully submitted to the payment network.
`settled`: Credits have been refunded to the Plaid linked account.
`cancelled`: The refund was cancelled by the client.
`failed`: The refund has failed.
`returned`: The refund was returned.  
  

Possible values: `pending`, `posted`, `cancelled`, `failed`, `settled`, `returned`

[`failure_reason`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-failure-reason)

nullableobjectnullable, object

The failure reason if the event type for a refund is `"failed"` or `"returned"`. Null value otherwise.

[`failure_code`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-failure-reason-failure-code)

nullablestringnullable, string

The failure code, e.g. `R01`. A failure code will be provided if and only if the refund status is `returned`. See [ACH return codes](https://plaid.com/docs/errors/transfer/#ach-return-codes) for a full listing of ACH return codes and [RTP/RfP error codes](https://plaid.com/docs/errors/transfer/#rtprfp-error-codes) for RTP error codes.

[`ach_return_code`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-failure-reason-ach-return-code)

deprecatednullablestringdeprecated, nullable, string

The ACH return code, e.g. `R01`. A return code will be provided if and only if the refund status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes). This field is deprecated in favor of the more versatile `failure_code`, which encompasses non-ACH failure codes as well.

[`description`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-failure-reason-description)

stringstring

A human-readable description of the reason for the failure or reversal.

[`ledger_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-ledger-id)

nullablestringnullable, string

Plaid’s unique identifier for a Plaid Ledger Balance.

[`created`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-created)

stringstring

The datetime when this refund was created. This will be of the form `2006-01-02T15:04:05Z`  
  

Format: `date-time`

[`network_trace_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-refunds-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`recurring_transfer_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-recurring-transfer-id)

nullablestringnullable, string

The id of the recurring transfer if this transfer belongs to a recurring transfer.

[`expected_sweep_settlement_schedule`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-expected-sweep-settlement-schedule)

[object][object]

The expected sweep settlement schedule of this transfer, assuming this transfer is not `returned`. Only applies to ACH debit transfers.

[`sweep_settlement_date`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-expected-sweep-settlement-schedule-sweep-settlement-date)

stringstring

The settlement date of a sweep for this transfer.  
  

Format: `date`

[`swept_settled_amount`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-expected-sweep-settlement-schedule-swept-settled-amount)

stringstring

The accumulated amount that has been swept by `sweep_settlement_date`.

[`credit_funds_source`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-credit-funds-source)

deprecatednullablestringdeprecated, nullable, string

This field is now deprecated. You may ignore it for transfers created on and after 12/01/2023.  
Specifies the source of funds for the transfer. Only valid for `credit` transfers, and defaults to `sweep` if not specified. This field is not specified for `debit` transfers.  
`sweep` - Sweep funds from your funding account
`prefunded_rtp_credits` - Use your prefunded RTP credit balance with Plaid
`prefunded_ach_credits` - Use your prefunded ACH credit balance with Plaid  
  

Possible values: `sweep`, `prefunded_rtp_credits`, `prefunded_ach_credits`, `null`

[`facilitator_fee`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-facilitator-fee)

stringstring

The amount to deduct from `transfer.amount` and distribute to the platform’s Ledger balance as a facilitator fee (decimal string with two digits of precision e.g. "10.00"). The remainder will go to the end-customer’s Ledger balance. This must be value greater than 0 and less than or equal to the `transfer.amount`.

[`network_trace_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-transfer-network-trace-id)

nullablestringnullable, string

The trace identifier for the transfer based on its network. This will only be set after the transfer has posted.  
For `ach` or `same-day-ach` transfers, this is the ACH trace number.
For `rtp` transfers, this is the Transaction Identification number.
For `wire` transfers, this is the IMAD (Input Message Accountability Data) number.

[`request_id`](/docs/api/products/transfer/initiating-transfers/#transfer-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "transfer": {
    "id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
    "authorization_id": "c9f90aa1-2949-c799-e2b6-ea05c89bb586",
    "ach_class": "ppd",
    "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
    "funding_account_id": "8945fedc-e703-463d-86b1-dc0607b55460",
    "ledger_id": "563db5f8-4c95-4e17-8c3e-cb988fb9cf1a",
    "type": "credit",
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
    "amount": "12.34",
    "description": "payment",
    "created": "2020-08-06T17:27:15Z",
    "refunds": [],
    "status": "pending",
    "network": "ach",
    "cancellable": true,
    "guarantee_decision": null,
    "guarantee_decision_rationale": null,
    "failure_reason": null,
    "metadata": {
      "key1": "value1",
      "key2": "value2"
    },
    "origination_account_id": "",
    "iso_currency_code": "USD",
    "standard_return_window": "2023-08-07",
    "unauthorized_return_window": "2023-10-07",
    "expected_settlement_date": "2023-08-04",
    "originator_client_id": "569ed2f36b3a3a021713abc1",
    "recurring_transfer_id": null,
    "credit_funds_source": "sweep",
    "facilitator_fee": "1.23",
    "network_trace_id": null
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

=\*=\*=\*=

#### `/transfer/cancel`

#### Cancel a transfer

Use the [`/transfer/cancel`](/docs/api/products/transfer/initiating-transfers/#transfercancel) endpoint to cancel a transfer. A transfer is eligible for cancellation if the `cancellable` property returned by [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) is `true`.

/transfer/cancel

**Request fields**

[`client_id`](/docs/api/products/transfer/initiating-transfers/#transfer-cancel-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/initiating-transfers/#transfer-cancel-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`transfer_id`](/docs/api/products/transfer/initiating-transfers/#transfer-cancel-request-transfer-id)

requiredstringrequired, string

Plaid’s unique identifier for a transfer.

[`reason_code`](/docs/api/products/transfer/initiating-transfers/#transfer-cancel-request-reason-code)

stringstring

Specifies the reason for cancelling transfer. This is required for RfP transfers, and will be ignored for other networks.  
`"AC03"` - Invalid Creditor Account Number  
`"AM09"` - Incorrect Amount  
`"CUST"` - Requested By Customer - Cancellation requested  
`"DUPL"` - Duplicate Payment  
`"FRAD"` - Fraudulent Payment - Unauthorized or fraudulently induced  
`"TECH"` - Technical Problem - Cancellation due to system issues  
`"UPAY"` - Undue Payment - Payment was made through another channel  
`"AC14"` - Invalid or Missing Creditor Account Type  
`"AM06"` - Amount Too Low  
`"BE05"` - Unrecognized Initiating Party  
`"FOCR"` - Following Refund Request  
`"MS02"` - No Specified Reason - Customer  
`"MS03"` - No Specified Reason - Agent  
`"RR04"` - Regulatory Reason  
`"RUTA"` - Return Upon Unable To Apply  
  

Possible values: `AC03`, `AM09`, `CUST`, `DUPL`, `FRAD`, `TECH`, `UPAY`, `AC14`, `AM06`, `BE05`, `FOCR`, `MS02`, `MS03`, `RR04`, `RUTA`

/transfer/cancel

```
const request: TransferCancelRequest = {
  transfer_id: '123004561178933',
};
try {
  const response = await plaidClient.transferCancel(request);
  const request_id = response.data.request_id;
} catch (error) {
  // handle error
}
```

/transfer/cancel

**Response fields**

[`request_id`](/docs/api/products/transfer/initiating-transfers/#transfer-cancel-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "saKrIBuEB9qJZno"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
