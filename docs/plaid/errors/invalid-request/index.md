---
title: "Errors - Invalid Request errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/invalid-request/"
scraped_at: "2026-03-07T22:04:50+00:00"
---

# Invalid Request Errors

#### Guide to troubleshooting invalid request errors

#### **INCOMPATIBLE\_API\_VERSION**

##### The request uses fields that are not compatible with the API version being used.

##### Common causes

- The API endpoint was called using a `public_key` for authentication rather than a `client_id` and `secret`.

##### Troubleshooting steps

Review the [API Reference](/docs/api/) for the endpoint to see what parameters are supported in the API version you are using.

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "INCOMPATIBLE_API_VERSION",
 "error_message": "The public_key cannot be used for this endpoint as of version {version-date} of the API. Please use the client_id and secret instead.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_ACCOUNT\_NUMBER**

##### The provided account number was invalid.

##### Sample user-facing error message

Account or routing number incorrect: Check with your bank and make sure that your account and routing numbers are entered correctly

##### Alternative user-facing error message

*This account does not support ACH debit. Please retry with a different account.*

##### Common causes

- While in the Instant Match, Automated Micro-deposit, or Same-Day Micro-deposit Link flows, the user provided an account number whose last four digits did not match the account mask of their bank account.
- If the user entered the correct account number, Plaid may have been unable to retrieve an account mask.
- If the user entered the correct account number, the account may be a non-debitable account or a non-supported account type. Common examples of non-debitable depository accounts include savings accounts at Chime or at Navy Federal Credit Union (NFCU).

##### Troubleshooting steps

Have the user confirm that they entered the correct account number.

Have the user confirm that the account is a debitable depository account, such as a checking, savings, or cash management account.

If the account number was correct and the account was a debitable supported type, please [contact Plaid Support](https://dashboard.plaid.com/support/new/).

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "INVALID_ACCOUNT_NUMBER",
 "error_message": "The provided account number was invalid.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_BODY**

##### The request body was invalid.

##### Common causes

- The JSON request body was malformed.
- The request `content-type` was not of type `application/json`. The Plaid API only accepts JSON text as the MIME media type,
  with `UTF-8` encoding, conforming to [RFC 4627](http://www.ietf.org/rfc/rfc4627.txt).

API request content-type

```
content-type: "application/json"
```

##### Troubleshooting steps

Resend the request with `content-type: 'application/json'`.

Resend the request with valid JSON in the body.

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "INVALID_BODY",
 "error_message": "body could not be parsed as JSON",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_CONFIGURATION**

##### /link/token/create was called with invalid configuration settings

##### Common causes

- One or more of the configuration objects provided to [`/link/token/create`](/docs/api/link/#linktokencreate) does not match the request schema for that endpoint.

##### Troubleshooting steps

Verify that all field names being provided to [`/link/token/create`](/docs/api/link/#linktokencreate) match the schema for that endpoint. In particular, when migrating from Link tokens, note that some field names have changed between those used for Link token style Link configuration and those used as parameters for [`/link/token/create`](/docs/api/link/#linktokencreate).

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "INVALID_CONFIGURATION",
 "error_message": "please ensure that the request body is formatted correctly",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_FIELD**

##### One or more of the request body fields were improperly formatted or invalid.

##### Common causes

- One or more fields in the request body were invalid, malformed, or used a wrong type. The `error_message` field will specify the erroneous field and how to resolve the error.
- Personally identifiable information (PII), such as an email address or phone number, was provided for a field where PII is not allowed, such as `user.client_user_id`.
- An unsupported country code was used in Production. Consult the [API Reference](/docs/api/) for the endpoint being used for a list of valid country codes.
- An request parameter that is optional in the API schema was not provided in a context where it is required. For example, [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) was called without specifying `options.min_last_updated_datetime` on a Capital One (`ins_128026`) Item with non-depository accounts.
- The value used in the field is not valid for business logic reasons. For example, [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) or [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport) endpoints were called with a `client_transaction_id` for which [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) was never called, or [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) was called using an `authorization_id` whose `decision` value is `declined`.
- A request to [`/link/token/create`](/docs/api/link/#linktokencreate) was sent specifying an OAuth redirect URI that was not added to the [allowed redirect URIs](https://dashboard.plaid.com/developers/api) list in the Dashboard.

##### Troubleshooting steps

Resend the request with the correctly formatted fields specified in the `error_message`.

Refer to the `error_message` field for instructions on the exact problem and how to resolve it.

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "INVALID_FIELD",
 "error_message": "{{ error message is specific to the given / missing request field }}",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_HEADERS**

##### The request was missing a required header.

##### Common causes

- The request was missing a `header`, typically the `Content-Type` header.

##### Troubleshooting steps

Resend the request with the missing headers.

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "INVALID_HEADERS",
 "error_message": "{{ error message is specific to the given / missing header }}",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **MISSING\_FIELDS**

##### The request was missing one or more required fields.

##### Common causes

- The request body is missing one or more required fields. The `error_message` field will list the missing field(s).

##### Troubleshooting steps

Resend the request with the missing required fields specified in the `error_message`.

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "MISSING_FIELDS",
 "error_message": "the following required fields are missing: {fields}",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NO\_LONGER\_AVAILABLE**

##### The endpoint requested is not available in the API version being used.

##### Common causes

- The endpoint you requested has been discontinued and no longer exists in the Plaid API.

##### Troubleshooting steps

Review the [API Reference](/docs/api/) to see what endpoints are supported in the API version you are using.

See the [Link Token migration guide](/docs/link/link-token-migration-guide/) for instructions on migrating away from endpoints that have been discontinued in recent API versions.

API error response

```
http code 404
{
 "error_type": "INVALID_REQUEST",
 "error_code": "NO_LONGER_AVAILABLE",
 "error_message": "This endpoint has been discontinued as of version {version-date} of the API.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NOT\_FOUND**

##### The endpoint requested does not exist.

##### Common causes

- The endpoint you requested does not exist in the Plaid API.

##### Troubleshooting steps

Navigate to the [API reference](/docs/api/) to find the correct API endpoint.

API error response

```
http code 404
{
 "error_type": "INVALID_REQUEST",
 "error_code": "NOT_FOUND",
 "error_message": "not found",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **SANDBOX\_ONLY**

##### The requested endpoint is only available in Sandbox.

##### Common causes

- The requested endpoint is only available in the [Sandbox API Environment](/docs/sandbox/).

##### Troubleshooting steps

Change your environment to [Sandbox](/docs/sandbox/) or remove usages of test endpoints.

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "SANDBOX_ONLY",
 "error_message": "access to {api/route} is only available in the sandbox environment at https://sandbox.plaid.com/",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **UNKNOWN\_FIELDS**

##### The request included a field that is not recognized by the endpoint.

##### Common causes

- The request body included one or more extraneous fields. The `error_message` field will list the unrecognized field(s).

##### Troubleshooting steps

Resend the request with the omitted fields specified in the `error_message`.

API error response

```
http code 400
{
 "error_type": "INVALID_REQUEST",
 "error_code": "UNKNOWN_FIELDS",
 "error_message": "the following fields are not recognized by this endpoint: {fields}",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
