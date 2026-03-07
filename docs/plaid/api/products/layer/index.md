---
title: "API - Layer | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/layer/"
scraped_at: "2026-03-07T22:04:14+00:00"
---

# Layer

#### API reference for Layer endpoints

For how-to guidance, see the [Layer documentation](/docs/layer/).

| Endpoints |  |
| --- | --- |
| [`/session/token/create`](/docs/api/products/layer/#sessiontokencreate) | Creates a Link token for a Layer session |
| [`/user_account/session/get`](/docs/api/products/layer/#user_accountsessionget) | Returns user permissioned account data |

| Webhooks |  |
| --- | --- |
| [`LAYER_AUTHENTICATION_PASSED`](/docs/api/products/layer/#layer_authentication_passed) | A user has been authenticated |
| [`SESSION_FINISHED`](/docs/api/products/layer/#session_finished) | A Layer session has finished |

### Endpoints

=\*=\*=\*=

#### `/session/token/create`

#### Create a Link token for Layer

[`/session/token/create`](/docs/api/products/layer/#sessiontokencreate) is used to create a Link token for Layer. The returned Link token is used as an parameter when initializing the Link SDK. For more details, see the [Link flow overview](https://plaid.com/docs/link/#link-flow-overview).

/session/token/create

**Request fields**

[`client_id`](/docs/api/products/layer/#session-token-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/layer/#session-token-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`template_id`](/docs/api/products/layer/#session-token-create-request-template-id)

requiredstringrequired, string

The id of a template defined in Plaid Dashboard

[`user`](/docs/api/products/layer/#session-token-create-request-user)

objectobject

Details about the end user. Required if a root-level `user_id` is not provided.

[`client_user_id`](/docs/api/products/layer/#session-token-create-request-user-client-user-id)

requiredstringrequired, string

A unique ID representing the end user. Typically this will be a user ID number from your application. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`. It is currently used as a means of searching logs for the given user in the Plaid Dashboard.

[`user_id`](/docs/api/products/layer/#session-token-create-request-user-user-id)

stringstring

The `user_id` created by calling `/user/create`. Provide this field only if you are using Plaid Check Report with Layer and have a `user_token`.

[`redirect_uri`](/docs/api/products/layer/#session-token-create-request-redirect-uri)

stringstring

A URI indicating the destination where a user should be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or another app. The `redirect_uri` should not contain any query parameters. When used in Production, must be an https URI. To specify any subdomain, use `*` as a wildcard character, e.g. `https://*.example.com/oauth.html`. Note that any redirect URI must also be added to the Allowed redirect URIs list in the [developer dashboard](https://dashboard.plaid.com/team/api). If initializing on Android, `android_package_name` must be specified instead and `redirect_uri` should be left blank.

[`android_package_name`](/docs/api/products/layer/#session-token-create-request-android-package-name)

stringstring

The name of your app's Android package. Required if using the session token to initialize Layer on Android. Any package name specified here must also be added to the Allowed Android package names setting on the [developer dashboard](https://dashboard.plaid.com/team/api). When creating a session token for initializing Layer on other platforms, `android_package_name` must be left blank and `redirect_uri` should be used instead.

[`webhook`](/docs/api/products/layer/#session-token-create-request-webhook)

stringstring

The destination URL to which any webhooks should be sent. If you use the same webhook listener for all Sandbox or all Production activity, set this value in the Layer template editor in the Dashboard instead. Only provide a value in this field if you need to use multiple webhook URLs per environment (an uncommon use case). If provided, a value in this field will take priority over webhook values set in the Layer template editor.  
  

Format: `url`

[`user_id`](/docs/api/products/layer/#session-token-create-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

/session/token/create

```
const request: SessionTokenCreateRequest = {
  user: {
    client_user_id: 'user-abc'
  },
  template_id: 'template_4uinBNe4B2x9'
};
try {
  const response = await client.sessionTokenCreate(request);
  const linkToken = response.data.link.link_token;
} catch (error) {
  // handle error
}
```

/session/token/create

**Response fields**

[`request_id`](/docs/api/products/layer/#session-token-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`link`](/docs/api/products/layer/#session-token-create-response-link)

objectobject

Response data for `/session/token/create` intended for use with the Link SDK.

[`link_token`](/docs/api/products/layer/#session-token-create-response-link-link-token)

stringstring

A Link token, which can be supplied to Link in order to initialize it and receive a `public_token`.

[`expiration`](/docs/api/products/layer/#session-token-create-response-link-expiration)

stringstring

The expiration date for the `link_token`, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. A `link_token` created to generate a `public_token` that will be exchanged for a new `access_token` expires after 4 hours. A `link_token` created for an existing Item (such as when updating an existing `access_token` by launching Link in update mode) expires after 30 minutes.  
  

Format: `date-time`

[`user_id`](/docs/api/products/layer/#session-token-create-response-link-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

Response Object

```
{
  "link": {
    "link_token": "link-sandbox-af1a0311-da53-4636-b754-dd15cc058176",
    "expiration": "2020-03-27T12:56:34Z"
  },
  "request_id": "XQVgFigpGHXkb0b"
}
```

=\*=\*=\*=

#### `/user_account/session/get`

#### Retrieve User Account

This endpoint returns user permissioned account data, including identity and Item access tokens, for use with [Plaid Layer](https://plaid.com/docs/layer). Note that end users are permitted to edit the prefilled identity data in the Link flow before sharing it with you; you should treat any identity data returned by this endpoint as user-submitted, unverified data. For a verification layer, you can add [Identity Verification](https://plaid.com/docs/identity-verification/) to your flow, or check the submitted identity data against bank account data from linked accounts using [Identity Match](https://plaid.com/docs/identity/#identity-match).

/user\_account/session/get

**Request fields**

[`client_id`](/docs/api/products/layer/#user_account-session-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/layer/#user_account-session-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`public_token`](/docs/api/products/layer/#user_account-session-get-request-public-token)

requiredstringrequired, string

The public token generated by the end user Layer session.

/user\_account/session/get

```
const request: UserAccountSessionGetRequest = {
  public_token: 'profile-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce992d',
};
try {
  const response = await client.userAccountSessionGet(request);
} catch (error) {
  // handle error
}
```

/user\_account/session/get

**Response fields**

[`identity`](/docs/api/products/layer/#user_account-session-get-response-identity)

nullableobjectnullable, object

The identity data permissioned by the end user during the authorization flow.

[`name`](/docs/api/products/layer/#user_account-session-get-response-identity-name)

nullableobjectnullable, object

The user's first name and last name.

[`first_name`](/docs/api/products/layer/#user_account-session-get-response-identity-name-first-name)

stringstring

[`last_name`](/docs/api/products/layer/#user_account-session-get-response-identity-name-last-name)

stringstring

[`address`](/docs/api/products/layer/#user_account-session-get-response-identity-address)

nullableobjectnullable, object

The user's address.

[`city`](/docs/api/products/layer/#user_account-session-get-response-identity-address-city)

nullablestringnullable, string

The full city name

[`region`](/docs/api/products/layer/#user_account-session-get-response-identity-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/layer/#user_account-session-get-response-identity-address-street)

nullablestringnullable, string

The full street address
Example: `"564 Main Street, APT 15"`

[`street2`](/docs/api/products/layer/#user_account-session-get-response-identity-address-street2)

nullablestringnullable, string

The second line street address

[`postal_code`](/docs/api/products/layer/#user_account-session-get-response-identity-address-postal-code)

nullablestringnullable, string

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/layer/#user_account-session-get-response-identity-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code

[`phone_number`](/docs/api/products/layer/#user_account-session-get-response-identity-phone-number)

stringstring

The user's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format

[`email`](/docs/api/products/layer/#user_account-session-get-response-identity-email)

nullablestringnullable, string

The user's email address.  
Note: email is currently not returned for users, and will be added later in 2026.

[`date_of_birth`](/docs/api/products/layer/#user_account-session-get-response-identity-date-of-birth)

nullablestringnullable, string

The user's date of birth.

[`ssn`](/docs/api/products/layer/#user_account-session-get-response-identity-ssn)

nullablestringnullable, string

The user's social security number.

[`ssn_last_4`](/docs/api/products/layer/#user_account-session-get-response-identity-ssn-last-4)

nullablestringnullable, string

The last 4 digits of the user's social security number.

[`items`](/docs/api/products/layer/#user_account-session-get-response-items)

[object][object]

[`item_id`](/docs/api/products/layer/#user_account-session-get-response-items-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`access_token`](/docs/api/products/layer/#user_account-session-get-response-items-access-token)

stringstring

The access token associated with the Item data is being requested for.

[`identity_edit_history`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history)

nullableobjectnullable, object

Statistics tracking the number of edits made to identity fields over various time periods.

[`name`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-name)

objectobject

Edit counts over various time periods.

[`edits_current`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-name-edits-current)

integerinteger

Number of edits in the current session

[`edits_1d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-name-edits-1d)

integerinteger

Number of edits in the last 1 day

[`edits_30d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-name-edits-30d)

integerinteger

Number of edits in the last 30 days

[`edits_365d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-name-edits-365d)

integerinteger

Number of edits in the last 365 days

[`edits_all_time`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-name-edits-all-time)

integerinteger

Total number of edits

[`address`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-address)

objectobject

Edit counts over various time periods.

[`edits_current`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-address-edits-current)

integerinteger

Number of edits in the current session

[`edits_1d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-address-edits-1d)

integerinteger

Number of edits in the last 1 day

[`edits_30d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-address-edits-30d)

integerinteger

Number of edits in the last 30 days

[`edits_365d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-address-edits-365d)

integerinteger

Number of edits in the last 365 days

[`edits_all_time`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-address-edits-all-time)

integerinteger

Total number of edits

[`email`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-email)

objectobject

Edit counts over various time periods.

[`edits_current`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-email-edits-current)

integerinteger

Number of edits in the current session

[`edits_1d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-email-edits-1d)

integerinteger

Number of edits in the last 1 day

[`edits_30d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-email-edits-30d)

integerinteger

Number of edits in the last 30 days

[`edits_365d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-email-edits-365d)

integerinteger

Number of edits in the last 365 days

[`edits_all_time`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-email-edits-all-time)

integerinteger

Total number of edits

[`date_of_birth`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-date-of-birth)

objectobject

Edit counts over various time periods.

[`edits_current`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-date-of-birth-edits-current)

integerinteger

Number of edits in the current session

[`edits_1d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-date-of-birth-edits-1d)

integerinteger

Number of edits in the last 1 day

[`edits_30d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-date-of-birth-edits-30d)

integerinteger

Number of edits in the last 30 days

[`edits_365d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-date-of-birth-edits-365d)

integerinteger

Number of edits in the last 365 days

[`edits_all_time`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-date-of-birth-edits-all-time)

integerinteger

Total number of edits

[`official_document`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-official-document)

nullableobjectnullable, object

Official identity document edit statistics.

[`ssn`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-official-document-ssn)

objectobject

Edit counts over various time periods.

[`edits_current`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-official-document-ssn-edits-current)

integerinteger

Number of edits in the current session

[`edits_1d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-official-document-ssn-edits-1d)

integerinteger

Number of edits in the last 1 day

[`edits_30d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-official-document-ssn-edits-30d)

integerinteger

Number of edits in the last 30 days

[`edits_365d`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-official-document-ssn-edits-365d)

integerinteger

Number of edits in the last 365 days

[`edits_all_time`](/docs/api/products/layer/#user_account-session-get-response-identity-edit-history-official-document-ssn-edits-all-time)

integerinteger

Total number of edits

[`request_id`](/docs/api/products/layer/#user_account-session-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "identity": {
    "name": {
      "first_name": "Leslie",
      "last_name": "Knope"
    },
    "address": {
      "street": "123 Main St.",
      "street2": "",
      "city": "Pawnee",
      "region": "IN",
      "postal_code": "41006",
      "country": "US"
    },
    "email": "leslie@knope.com",
    "phone_number": "+14157452130",
    "date_of_birth": "1975-01-18",
    "ssn": "987654321",
    "ssn_last_4": "4321"
  },
  "items": [
    {
      "item_id": "Ed6bjNrDLJfGvZWwnkQlfxwoNz54B5C97ejBr",
      "access_token": "access-sandbox-435beced-94e8-4df3-a181-1dde1cfa19f0"
    }
  ],
  "request_id": "m8MDnv9okwxFNBV"
}
```

### Webhooks

=\*=\*=\*=

#### `LAYER_AUTHENTICATION_PASSED`

Indicates that Plaid's authentication process has completed for a user and that Plaid has verified that the user owns their phone number. If you receive this webhook, you should skip your own OTP phone number verification flow for the user, even if the user does not complete the entire Link flow. If the user doesn't complete the full Link flow (as verified by your being able to successfully call [`/user_account/session/get`](/docs/api/products/layer/#user_accountsessionget) using the `public_token` from the `onSuccess` callback) it is recommended that you implement [webhook verification](https://plaid.com/docs/api/webhooks/webhook-verification/) or another technique to avoid webhook spoofing attacks.

**Properties**

[`webhook_type`](/docs/api/products/layer/#LayerAuthenticationPassedWebhook-webhook-type)

stringstring

`LAYER`

[`webhook_code`](/docs/api/products/layer/#LayerAuthenticationPassedWebhook-webhook-code)

stringstring

`LAYER_AUTHENTICATION_PASSED`

[`environment`](/docs/api/products/layer/#LayerAuthenticationPassedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

[`link_session_id`](/docs/api/products/layer/#LayerAuthenticationPassedWebhook-link-session-id)

stringstring

An identifier for the Link session these events occurred in

[`link_token`](/docs/api/products/layer/#LayerAuthenticationPassedWebhook-link-token)

stringstring

The Link token used to create the Link session these events are from

API Object

```
{
  "webhook_type": "LAYER",
  "webhook_code": "LAYER_AUTHENTICATION_PASSED",
  "environment": "production",
  "link_session_id": "1daca4d5-9a0d-4e85-a2e9-1e905ecaa32e",
  "link_token": "link-sandbox-79e723b0-0e04-4248-8a33-15ceb6828a45"
}
```

=\*=\*=\*=

#### `SESSION_FINISHED`

Contains the state of a completed Link session, along with the public token(s) if available.

By default, this webhook is sent only for sessions enabled for the Hosted Link flow (including Link Recovery flows), a Multi-Item Link flow, or a Layer flow. If you would like to receive this webhook for other sessions, contact your Account Manager or Support. This enablement will also enable the `EVENTS` webhook for all Link sessions and the ability to use [`/link/token/get`](/docs/api/link/#linktokenget) to retrieve events for non-Hosted-Link sessions.

**Properties**

[`webhook_type`](/docs/api/products/layer/#LinkSessionFinishedWebhook-webhook-type)

stringstring

`LINK`

[`webhook_code`](/docs/api/products/layer/#LinkSessionFinishedWebhook-webhook-code)

stringstring

`SESSION_FINISHED`

[`status`](/docs/api/products/layer/#LinkSessionFinishedWebhook-status)

stringstring

The final status of the Link session. Will always be "SUCCESS" or "EXITED".

[`link_session_id`](/docs/api/products/layer/#LinkSessionFinishedWebhook-link-session-id)

stringstring

The identifier for the Link session.

[`link_token`](/docs/api/products/layer/#LinkSessionFinishedWebhook-link-token)

stringstring

The link token used to create the Link session.

[`public_token`](/docs/api/products/layer/#LinkSessionFinishedWebhook-public-token)

deprecatedstringdeprecated, string

The public token generated by the Link session. This field has been deprecated; please use `public_tokens` instead.

[`public_tokens`](/docs/api/products/layer/#LinkSessionFinishedWebhook-public-tokens)

[string][string]

The public tokens generated by the Link session.

[`user_id`](/docs/api/products/layer/#LinkSessionFinishedWebhook-user-id)

stringstring

The Plaid `user_id` of the User associated with this webhook, warning, or error.

[`environment`](/docs/api/products/layer/#LinkSessionFinishedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "LINK",
  "webhook_code": "SESSION_FINISHED",
  "status": "SUCCESS",
  "link_session_id": "356dbb28-7f98-44d1-8e6d-0cec580f3171",
  "link_token": "link-sandbox-af1a0311-da53-4636-b754-dd15cc058176",
  "public_tokens": [
    "public-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d"
  ],
  "environment": "sandbox"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
