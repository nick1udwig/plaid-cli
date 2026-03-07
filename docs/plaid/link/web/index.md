---
title: "Link - Web | Plaid Docs"
source_url: "https://plaid.com/docs/link/web/"
scraped_at: "2026-03-07T22:05:09+00:00"
---

# Link Web SDK

#### Reference for integrating with the Link JavaScript SDK and React SDK

Prefer to learn with code examples? Check out our [GitHub repo](https://github.com/plaid/tiny-quickstart) with working example Link implementations for both [JavaScript](https://github.com/plaid/tiny-quickstart/tree/main/vanilla_js) and [React](https://github.com/plaid/tiny-quickstart/tree/main/react).

#### Installation

Select group for content switcher

Include the Plaid Link initialize script on each page of your site. It must always be loaded directly from `https://cdn.plaid.com`, rather than included in a bundle or hosted yourself. Unlike Plaid's other SDKs, the JavaScript web SDK is not versioned; `cdn.plaid.com` will automatically provide the latest available SDK.

index.html

```
<script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
```

To get started with Plaid Link for React, clone the [GitHub repository](https://github.com/plaid/react-plaid-link) and review the example application and README, which provide reference implementations.

Next, you'll need to install the react-plaid-link package.

With npm:

```
npm install --save react-plaid-link
```

With yarn:

```
yarn add react-plaid-link
```

Then import the necessary components and types:

```
import {
  usePlaidLink,
  PlaidLinkOptions,
  PlaidLinkOnSuccess,
} from 'react-plaid-link';
```

#### CSP directives

If you are using a Content Security Policy (CSP), use the following directives to allow Link traffic:

CSP Directives

```
default-src https://cdn.plaid.com/;
script-src 'unsafe-inline' https://cdn.plaid.com/link/v2/stable/link-initialize.js;
frame-src https://cdn.plaid.com/;
connect-src https://production.plaid.com/;
```

If using Sandbox instead of Production, make sure to update the `connect-src` directive to point to the appropriate server (`https://sandbox.plaid.com`).

If your organization does not allow the use of `unsafe-inline`, you can use a [CSP nonce](https://content-security-policy.com/nonce/) instead.

#### Creating a Link token

Before you can create an instance of Link, you need to first create a `link_token`. A `link_token` can be configured for
different Link flows and is used to control much of Link's behavior. To learn how to create a new
`link_token`, see the API Reference entry for [`/link/token/create`](/docs/api/link/#linktokencreate).

=\*=\*=\*=

#### create()

`Plaid.create` accepts one argument, a configuration `Object`, and returns an `Object` with three functions, [`open`](#open), [`exit`](#exit), and [`destroy`](#destroy). Calling `open` will open Link and display the Consent Pane view, calling `exit` will close Link, and calling `destroy` will clean up the iframe.

It is recommended to call `Plaid.create` when initializing the view that is responsible for loading Plaid, as this will allow Plaid to pre-initialize Link, resulting in lower UI latency upon calling `open`, which can increase Link conversion.

When using the React SDK, this method is called `usePlaidLink` and returns an object with four values, [`open`](#open), [`exit`](#exit), `ready`, and `error`. The values `open` and `exit` behave as described above. `ready` is a passthrough for `onLoad` and will be `true` when Link is ready to be opened. `error` is populated only if Plaid fails to load the Link JavaScript. There is no separate method to destroy Link in the React SDK, as unmount will automatically destroy the Link instance.

**Note:** Control whether or not your Link integration uses the Account Select view from the [Dashboard](https://dashboard.plaid.com/signin?redirect=%2Flink%2Faccount-select).

create

**Properties**

[`token`](/docs/link/web/#link-web-create-token)

stringstring

Specify a `link_token` to authenticate your app with Link. This is a short lived, one-time use token that should be unique for each Link session. In addition to the primary flow, a `link_token` can be configured to launch Link in [update mode](/docs/link/update-mode/). See the `/link/token/create` endpoint for a full list of configurations.

[`onSuccess`](/docs/link/web/#link-web-create-onSuccess)

callbackcallback

A function that is called when a user successfully links an Item. The function should expect two arguments, the `public_token` and a `metadata` object. See [onSuccess](#onsuccess).

[`onExit`](/docs/link/web/#link-web-create-onExit)

callbackcallback

A function that is called when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. The function should expect two arguments, a nullable `error` object and a `metadata` object. See [onExit](#onexit).

[`onEvent`](/docs/link/web/#link-web-create-onEvent)

callbackcallback

A function that is called when a user reaches certain points in the Link flow. The function should expect two arguments, an `eventName` string and a `metadata` object. See [onEvent](#onevent).

[`onLoad`](/docs/link/web/#link-web-create-onLoad)

callbackcallback

A function that is called when the Link module has finished loading. Calls to `plaidLinkHandler.open()` prior to the `onLoad` callback will be delayed until the module is fully loaded.

[`receivedRedirectUri`](/docs/link/web/#link-web-create-receivedRedirectUri)

stringstring

A `receivedRedirectUri` is required to support OAuth authentication flows when re-launching Link on a mobile device.

[`key`](/docs/link/web/#link-web-create-key)

deprecatedstringdeprecated, string

The `public_key` is no longer used for new implementations of Link. If your integration is still using a `public_key`, please contact Plaid support or your account manager.

Create example

```
// The usePlaidLink hook manages Plaid Link creation
// It does not return a destroy function;
// instead, on unmount it automatically destroys the Link instance
const config: PlaidLinkOptions = {
  onSuccess: (public_token, metadata) => {},
  onExit: (err, metadata) => {},
  onEvent: (eventName, metadata) => {},
  token: 'GENERATED_LINK_TOKEN',
};

const { open, exit, ready } = usePlaidLink(config);
```

=\*=\*=\*=

#### onSuccess

The `onSuccess` callback is called when a user successfully links an Item. It takes two arguments: the `public_token` and a `metadata` object.

onSuccess

**Properties**

[`public_token`](/docs/link/web/#link-web-onsuccess-public-token)

stringstring

Displayed once a user has successfully completed Link. If using Identity Verification or Beacon, this field will be `null`. If using Document Income or Payroll Income, the `public_token` will be returned, but is not used.

[`metadata`](/docs/link/web/#link-web-onsuccess-metadata)

objectobject

Displayed once a user has successfully completed Link.

[`institution`](/docs/link/web/#link-web-onsuccess-metadata-institution)

nullableobjectnullable, object

An institution object. If the Item was created via Same-Day micro-deposit verification, will be `null`.

[`name`](/docs/link/web/#link-web-onsuccess-metadata-institution-name)

stringstring

The full institution name, such as `'Wells Fargo'`

[`institution_id`](/docs/link/web/#link-web-onsuccess-metadata-institution-institution-id)

stringstring

The Plaid institution identifier

[`accounts`](/docs/link/web/#link-web-onsuccess-metadata-accounts)

objectobject

A list of accounts attached to the connected Item. If Account Select is enabled via the developer dashboard, `accounts` will only include selected accounts.

[`id`](/docs/link/web/#link-web-onsuccess-metadata-accounts-id)

stringstring

The Plaid `account_id`

[`name`](/docs/link/web/#link-web-onsuccess-metadata-accounts-name)

stringstring

The official account name

[`mask`](/docs/link/web/#link-web-onsuccess-metadata-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts. It may also not match the mask that the bank displays to the user.

[`type`](/docs/link/web/#link-web-onsuccess-metadata-accounts-type)

stringstring

The account type. See the [Account schema](/docs/api/accounts#account-type-schema) for a full list of possible values

[`subtype`](/docs/link/web/#link-web-onsuccess-metadata-accounts-subtype)

stringstring

The account subtype. See the [Account schema](/docs/api/accounts#account-type-schema) for a full list of possible values

[`verification_status`](/docs/link/web/#link-web-onsuccess-metadata-accounts-verification-status)

nullablestringnullable, string

Indicates an Item's micro-deposit-based verification or database verification status. Possible values are:  
`pending_automatic_verification`: The Item is pending automatic verification  
`pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the deposit.  
`automatically_verified`: The Item has successfully been automatically verified  
`manually_verified`: The Item has successfully been manually verified  
`verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  
`verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.  
`database_matched`: The Item has successfully been verified using Plaid's data sources.  
`database_insights_pending`: The Database Insights result is pending and will be available upon Auth request.  
`null`: Neither micro-deposit-based verification nor database verification are being used for the Item.

[`class_type`](/docs/link/web/#link-web-onsuccess-metadata-accounts-class-type)

nullablestringnullable, string

If micro-deposit verification is being used, indicates whether the account being verified is a `business` or `personal` account.

[`account`](/docs/link/web/#link-web-onsuccess-metadata-account)

deprecatednullableobjectdeprecated, nullable, object

Deprecated. Use `accounts` instead.

[`link_session_id`](/docs/link/web/#link-web-onsuccess-metadata-link-session-id)

stringstring

A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.

[`transfer_status`](/docs/link/web/#link-web-onsuccess-metadata-transfer-status)

nullablestringnullable, string

The status of a transfer. Returned only when [Transfer UI](/docs/transfer/using-transfer-ui) is implemented.  

- `COMPLETE` – The transfer was completed.
- `INCOMPLETE` – The transfer could not be completed. For help, see [Troubleshooting Transfer UI](/docs/transfer/using-transfer-ui#troubleshooting-transfer-ui).
  
  

Possible values: `COMPLETE`, `INCOMPLETE`

onSuccess example

```
import {
  PlaidLinkOnSuccess,
  PlaidLinkOnSuccessMetadata,
} from 'react-plaid-link';

const onSuccess = useCallback<PlaidLinkOnSuccess>(
  (public_token: string, metadata: PlaidLinkOnSuccessMetadata) => {
    // log and save metadata
    // exchange public token (if using Item-based products)
    fetch('//yourserver.com/exchange-public-token', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        public_token,
      }),
    });
  },
  [],
);
```

Metadata schema

```
{
  institution: {
    name: 'Wells Fargo',
    institution_id: 'ins_4'
  },
  accounts: [
    {
      id: 'ygPnJweommTWNr9doD6ZfGR6GGVQy7fyREmWy',
      name: 'Plaid Checking',
      mask: '0000',
      type: 'depository',
      subtype: 'checking',
      verification_status: null
    },
    {
      id: '9ebEyJAl33FRrZNLBG8ECxD9xxpwWnuRNZ1V4',
      name: 'Plaid Saving',
      mask: '1111',
      type: 'depository',
      subtype: 'savings'
    }
    ...
  ],
  link_session_id: '79e772be-547d-4c9c-8b76-4ac4ed4c441a'
}
```

=\*=\*=\*=

#### onExit

The `onExit` callback is called when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. `onExit` takes two arguments, a nullable `error` object and a `metadata` object. The `metadata` parameter is always present, though some values may be `null`. Note that `onExit` will not be called when Link is destroyed in some other way than closing Link, such as the user hitting the browser back button or closing the browser tab on which the Link session is present.

onExit

**Properties**

[`error`](/docs/link/web/#link-web-onexit-error)

nullableobjectnullable, object

A nullable object that contains the error type, code, and message of the error that was last encountered by the user. If no error was encountered, error will be null.

[`error_type`](/docs/link/web/#link-web-onexit-error-error-type)

StringString

A broad categorization of the error.

[`error_code`](/docs/link/web/#link-web-onexit-error-error-code)

StringString

The particular error code. Each `error_type` has a specific set of `error_codes`.

[`error_message`](/docs/link/web/#link-web-onexit-error-error-message)

StringString

A developer-friendly representation of the error code.

[`display_message`](/docs/link/web/#link-web-onexit-error-display-message)

nullableStringnullable, String

A user-friendly representation of the error code. `null` if the error is not related to user action. This may change over time and is not safe for programmatic use.

[`metadata`](/docs/link/web/#link-web-onexit-metadata)

objectobject

Displayed if a user exits Link without successfully linking an Item.

[`institution`](/docs/link/web/#link-web-onexit-metadata-institution)

nullableobjectnullable, object

An institution object. If the Item was created via Same-Day micro-deposit verification, will be `null`.

[`name`](/docs/link/web/#link-web-onexit-metadata-institution-name)

stringstring

The full institution name, such as `Wells Fargo`

[`institution_id`](/docs/link/web/#link-web-onexit-metadata-institution-institution-id)

stringstring

The Plaid institution identifier

[`status`](/docs/link/web/#link-web-onexit-metadata-status)

stringstring

The point at which the user exited the Link flow. One of the following values.

[`requires_questions`](/docs/link/web/#link-web-onexit-metadata-status-requires-questions)

User prompted to answer security questions

[`requires_selections`](/docs/link/web/#link-web-onexit-metadata-status-requires-selections)

User prompted to answer multiple choice question(s)

[`requires_code`](/docs/link/web/#link-web-onexit-metadata-status-requires-code)

User prompted to provide a one-time passcode

[`choose_device`](/docs/link/web/#link-web-onexit-metadata-status-choose-device)

User prompted to select a device on which to receive a one-time passcode

[`requires_credentials`](/docs/link/web/#link-web-onexit-metadata-status-requires-credentials)

User prompted to provide credentials for the selected financial institution or has not yet selected a financial institution

[`requires_account_selection`](/docs/link/web/#link-web-onexit-metadata-status-requires-account-selection)

User prompted to select one or more financial accounts to share

[`requires_oauth`](/docs/link/web/#link-web-onexit-metadata-status-requires-oauth)

User prompted to enter an OAuth flow

[`institution_not_found`](/docs/link/web/#link-web-onexit-metadata-status-institution-not-found)

User exited the Link flow on the institution selection pane. Typically this occurs after the user unsuccessfully (no results returned) searched for a financial institution. Note that this status does not necessarily indicate that the user was unable to find their institution, as it is used for all user exits that occur from the institution selection pane, regardless of other user behavior.

[`institution_not_supported`](/docs/link/web/#link-web-onexit-metadata-status-institution-not-supported)

User exited the Link flow after discovering their selected institution is no longer supported by Plaid

[`link_session_id`](/docs/link/web/#link-web-onexit-metadata-link-session-id)

stringstring

A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.

[`request_id`](/docs/link/web/#link-web-onexit-metadata-request-id)

stringstring

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation.

onExit example

```
import {
  PlaidLinkOnExit,
  PlaidLinkOnExitMetadata,
  PlaidLinkError,
} from 'react-plaid-link';

const onExit = useCallback<PlaidLinkOnExit>(
  (error: PlaidLinkError, metadata: PlaidLinkOnExitMetadata) => {
    // log and save error and metadata
    // handle invalid link token
    if (error != null && error.error_code === 'INVALID_LINK_TOKEN') {
      // generate new link token
    }
    // to handle other error codes, see https://plaid.com/docs/errors/
  },
  [],
);
```

Error schema

```
{
  error_type: 'ITEM_ERROR',
  error_code: 'INVALID_CREDENTIALS',
  error_message: 'the credentials were not correct',
  display_message: 'The credentials were not correct.',
}
```

Metadata schema

```
{
  institution: {
    name: 'Wells Fargo',
    institution_id: 'ins_4'
  },
  status: 'requires_credentials',
  link_session_id: '36e201e0-2280-46f0-a6ee-6d417b450438',
  request_id: '8C7jNbDScC24THu'
}
```

=\*=\*=\*=

#### onEvent

The `onEvent` callback is called at certain points in the Link flow. It takes two arguments, an `eventName` string and a `metadata` object.

The `metadata` parameter is always present, though some values may be `null`. Note that new `eventNames`, `metadata` keys, or view names may be added without notice.

The `OPEN`, `LAYER_READY`, `LAYER_NOT_AVAILABLE`, and `LAYER_AUTOFILL_NOT_AVAILABLE` events will fire in real time; subsequent events will fire at the end of the Link flow, along with the `onSuccess` or `onExit` callback. Callback ordering is not guaranteed; `onEvent` callbacks may fire before, after, or surrounding the `onSuccess` or `onExit` callback, and event callbacks are not guaranteed to fire in the order in which they occurred. If you need to determine the exact time when an event happened, use the `timestamp` in the metadata.

The following callback events are stable, which means that they are suitable for programmatic use in your application's logic: `OPEN`, `EXIT`, `HANDOFF`, `SELECT_INSTITUTION`, `ERROR`, `BANK_INCOME_INSIGHTS_COMPLETED`, `IDENTITY_VERIFICATION_PASS_SESSION`, `IDENTITY_VERIFICATION_FAIL_SESSION`, `IDENTITY_MATCH_FAILED`, `IDENTITY_MATCH_PASSED`, `LAYER_READY`, `LAYER_NOT_AVAILABLE`, `LAYER_AUTOFILL_NOT_AVAILABLE`. The remaining callback events are informational and subject to change and should be used for analytics and troubleshooting purposes only.

onEvent

**Properties**

[`eventName`](/docs/link/web/#link-web-onevent-eventName)

stringstring

A string representing the event that has just occurred in the Link flow.

[`AUTO_SUBMIT_PHONE`](/docs/link/web/#link-web-onevent-eventName-AUTO-SUBMIT-PHONE)

The user was automatically sent an OTP code without a UI prompt. This event can only occur if the user's phone phone number was provided to Link via the `/link/token/create` call and the user has previously consented to receive OTP codes from Plaid.

[`BANK_INCOME_INSIGHTS_COMPLETED`](/docs/link/web/#link-web-onevent-eventName-BANK-INCOME-INSIGHTS-COMPLETED)

The user has completed the Assets and Bank Income Insights flow.

[`CLOSE_OAUTH`](/docs/link/web/#link-web-onevent-eventName-CLOSE-OAUTH)

The user closed the third-party website or mobile app without completing the OAuth flow.

[`CONNECT_NEW_INSTITUTION`](/docs/link/web/#link-web-onevent-eventName-CONNECT-NEW-INSTITUTION)

The user has chosen to link a new institution instead of linking a saved institution. This event is only emitted in the Link Returning User Experience flow.

[`ERROR`](/docs/link/web/#link-web-onevent-eventName-ERROR)

A recoverable error occurred in the Link flow, see the `error_code` metadata.

[`EXIT`](/docs/link/web/#link-web-onevent-eventName-EXIT)

The user has exited without completing the Link flow and the [onExit](#onexit) callback is fired.

[`FAIL_OAUTH`](/docs/link/web/#link-web-onevent-eventName-FAIL-OAUTH)

The user encountered an error while completing the third-party's OAuth login flow.

[`HANDOFF`](/docs/link/web/#link-web-onevent-eventName-HANDOFF)

The user has exited Link after successfully linking an Item.

[`IDENTITY_MATCH_FAILED`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-MATCH-FAILED)

An Identity Match check configured via the Account Verification Dashboard failed the Identity Match rules and did not detect a match.

[`IDENTITY_MATCH_PASSED`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-MATCH-PASSED)

An Identity Match check configured via the Account Verification Dashboard passed the Identity Match rules and detected a match.

[`IDENTITY_VERIFICATION_START_STEP`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-START-STEP)

The user has started a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_PASS_STEP`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-PASS-STEP)

The user has passed a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_FAIL_STEP`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-FAIL-STEP)

The user has failed a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_PENDING_REVIEW_STEP`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-PENDING-REVIEW-STEP)

The user has reached the pending review state.

[`IDENTITY_VERIFICATION_CREATE_SESSION`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-CREATE-SESSION)

The user has started a new Identity Verification session.

[`IDENTITY_VERIFICATION_RESUME_SESSION`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-RESUME-SESSION)

The user has resumed an existing Identity Verification session.

[`IDENTITY_VERIFICATION_PASS_SESSION`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-PASS-SESSION)

The user has passed their Identity Verification session.

[`IDENTITY_VERIFICATION_FAIL_SESSION`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-FAIL-SESSION)

The user has failed their Identity Verification session.

[`IDENTITY_VERIFICATION_PENDING_REVIEW_SESSION`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-PENDING-REVIEW-SESSION)

The user has completed their Identity Verification session, which is now in a pending review state.

[`IDENTITY_VERIFICATION_OPEN_UI`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-OPEN-UI)

The user has opened the UI of their Identity Verification session.

[`IDENTITY_VERIFICATION_RESUME_UI`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-RESUME-UI)

The user has resumed the UI of their Identity Verification session.

[`IDENTITY_VERIFICATION_CLOSE_UI`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-CLOSE-UI)

The user has closed the UI of their Identity Verification session.

[`LAYER_AUTOFILL_NOT_AVAILABLE`](/docs/link/web/#link-web-onevent-eventName-LAYER-AUTOFILL-NOT-AVAILABLE)

The user's date of birth passed to Link is not eligible for Layer Extended Autofill.

[`LAYER_NOT_AVAILABLE`](/docs/link/web/#link-web-onevent-eventName-LAYER-NOT-AVAILABLE)

The user phone number passed to Link is not eligible for Layer.

[`LAYER_READY`](/docs/link/web/#link-web-onevent-eventName-LAYER-READY)

The user phone number passed to Link is eligible for Layer and `open()` may now be called.

[`MATCHED_SELECT_INSTITUTION`](/docs/link/web/#link-web-onevent-eventName-MATCHED-SELECT-INSTITUTION)

The user selected an institution that was presented as a matched institution. This event can be emitted if [Embedded Institution Search](https://plaid.com/docs/link/embedded-institution-search/) is being used, if the institution was surfaced as a matched institution likely to have been linked to Plaid by a returning user, or if the institution's `routing_number` was provided when calling `/link/token/create`. For details on which scenario is triggering the event, see `metadata.matchReason`.

[`OPEN`](/docs/link/web/#link-web-onevent-eventName-OPEN)

The user has opened Link.

[`OPEN_MY_PLAID`](/docs/link/web/#link-web-onevent-eventName-OPEN-MY-PLAID)

The user has opened my.plaid.com. This event is only sent when Link is initialized with Assets as a product

[`OPEN_OAUTH`](/docs/link/web/#link-web-onevent-eventName-OPEN-OAUTH)

The user has navigated to a third-party website or mobile app in order to complete the OAuth login flow.

[`SEARCH_INSTITUTION`](/docs/link/web/#link-web-onevent-eventName-SEARCH-INSTITUTION)

The user has searched for an institution.

[`SELECT_AUTH_TYPE`](/docs/link/web/#link-web-onevent-eventName-SELECT-AUTH-TYPE)

The user has chosen whether to Link instantly or manually (i.e., with micro-deposits). This event emits the `selection` metadata to indicate the user's selection.

[`SELECT_BRAND`](/docs/link/web/#link-web-onevent-eventName-SELECT-BRAND)

The user selected a brand, e.g. Bank of America. The `SELECT_BRAND` event is only emitted for large financial institutions with multiple online banking portals.

[`SELECT_DEGRADED_INSTITUTION`](/docs/link/web/#link-web-onevent-eventName-SELECT-DEGRADED-INSTITUTION)

The user selected an institution with a `DEGRADED` health status and was shown a corresponding message.

[`SELECT_DOWN_INSTITUTION`](/docs/link/web/#link-web-onevent-eventName-SELECT-DOWN-INSTITUTION)

The user selected an institution with a `DOWN` health status and was shown a corresponding message.

[`SELECT_FILTERED_INSTITUTION`](/docs/link/web/#link-web-onevent-eventName-SELECT-FILTERED-INSTITUTION)

The user selected an institution Plaid does not support all requested products for and was shown a corresponding message.

[`SELECT_INSTITUTION`](/docs/link/web/#link-web-onevent-eventName-SELECT-INSTITUTION)

The user selected an institution.

[`SKIP_SUBMIT_PHONE`](/docs/link/web/#link-web-onevent-eventName-SKIP-SUBMIT-PHONE)

The user has opted to not provide their phone number to Plaid. This event is only emitted in the Link Returning User Experience flow.

[`SUBMIT_ACCOUNT_NUMBER`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-ACCOUNT-NUMBER)

The user has submitted an account number. This event emits the `account_number_mask` metadata to indicate the mask of the account number the user provided.

[`SUBMIT_CREDENTIALS`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-CREDENTIALS)

The user has submitted credentials.

[`SUBMIT_DOCUMENTS`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-DOCUMENTS)

The user is being prompted to submit documents for an Income verification flow.

[`SUBMIT_DOCUMENTS_ERROR`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-DOCUMENTS-ERROR)

The user encountered an error when submitting documents for an Income verification flow.

[`SUBMIT_DOCUMENTS_SUCCESS`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-DOCUMENTS-SUCCESS)

The user has successfully submitted documents for an Income verification flow.

[`SUBMIT_MFA`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-MFA)

The user has submitted MFA.

[`SUBMIT_OTP`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-OTP)

The user has submitted an OTP code during the phone number verification flow. This event is only emitted in the Link Returning User Experience (Remember Me) flow or Layer flow. This event will not be emitted if the phone number is verified via SNA.

[`SUBMIT_PHONE`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-PHONE)

The user has submitted their phone number. This event is only emitted in the Link Returning User Experience (Remember Me) flow.

[`SUBMIT_ROUTING_NUMBER`](/docs/link/web/#link-web-onevent-eventName-SUBMIT-ROUTING-NUMBER)

The user has submitted a routing number. This event emits the `routing_number` metadata to indicate user's routing number.

[`TRANSITION_VIEW`](/docs/link/web/#link-web-onevent-eventName-TRANSITION-VIEW)

The `TRANSITION_VIEW` event indicates that the user has moved from one view to the next.

[`VERIFY_PHONE`](/docs/link/web/#link-web-onevent-eventName-VERIFY-PHONE)

The user has successfully verified their phone number via either OTP or SNA. This event is only emitted in the Link Returning User Experience (Remember Me) flow or the Layer flow.

[`VIEW_DATA_TYPES`](/docs/link/web/#link-web-onevent-eventName-VIEW-DATA-TYPES)

The user has viewed data types on the data transparency consent pane.

[`metadata`](/docs/link/web/#link-web-onevent-metadata)

objectobject

An object containing information about the event.

[`account_number_mask`](/docs/link/web/#link-web-onevent-metadata-account-number-mask)

nullablestringnullable, string

The account number mask extracted from the user-provided account number. If the user-inputted account number is four digits long, `account_number_mask` is empty. Emitted by `SUBMIT_ACCOUNT_NUMBER`.

[`error_type`](/docs/link/web/#link-web-onevent-metadata-error-type)

nullablestringnullable, string

The error type that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`error_code`](/docs/link/web/#link-web-onevent-metadata-error-code)

nullablestringnullable, string

The error code that the user encountered. Emitted by `ERROR`, `EXIT`.

[`error_message`](/docs/link/web/#link-web-onevent-metadata-error-message)

nullablestringnullable, string

The error message that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`exit_status`](/docs/link/web/#link-web-onevent-metadata-exit-status)

nullablestringnullable, string

The status key indicates the point at which the user exited the Link flow. Emitted by: `EXIT`

[`institution_id`](/docs/link/web/#link-web-onevent-metadata-institution-id)

nullablestringnullable, string

The ID of the selected institution. Emitted by: *all events*.

[`institution_name`](/docs/link/web/#link-web-onevent-metadata-institution-name)

nullablestringnullable, string

The name of the selected institution. Emitted by: *all events*.

[`institution_search_query`](/docs/link/web/#link-web-onevent-metadata-institution-search-query)

nullablestringnullable, string

The query used to search for institutions. Emitted by: `SEARCH_INSTITUTION`.

[`is_update_mode`](/docs/link/web/#link-web-onevent-metadata-is-update-mode)

nullablestringnullable, string

Indicates if the current Link session is an update mode session. Emitted by: `OPEN`.

[`match_reason`](/docs/link/web/#link-web-onevent-metadata-match-reason)

nullablestringnullable, string

The reason this institution was matched.
This will be either `returning_user` or `routing_number` if emitted by: `MATCHED_SELECT_INSTITUTION`.
Otherwise, this will be `SAVED_INSTITUTION` or `AUTO_SELECT_SAVED_INSTITUTION` if emitted by: `SELECT_INSTITUTION`.

[`routing_number`](/docs/link/web/#link-web-onevent-metadata-routing-number)

nullablestringnullable, string

The routing number submitted by user at the micro-deposits routing number pane. Emitted by `SUBMIT_ROUTING_NUMBER`.

[`mfa_type`](/docs/link/web/#link-web-onevent-metadata-mfa-type)

nullablestringnullable, string

If set, the user has encountered one of the following MFA types: `code`, `device`, `questions`, `selections`. Emitted by: `SUBMIT_MFA` and `TRANSITION_VIEW` when `view_name` is `MFA`

[`view_name`](/docs/link/web/#link-web-onevent-metadata-view-name)

nullablestringnullable, string

The name of the view that is being transitioned to. Emitted by: `TRANSITION_VIEW`.

[`ACCEPT_TOS`](/docs/link/web/#link-web-onevent-metadata-view-name-ACCEPT-TOS)

The view showing Terms of Service in the identity verification flow.

[`CONNECTED`](/docs/link/web/#link-web-onevent-metadata-view-name-CONNECTED)

The user has connected their account.

[`CONSENT`](/docs/link/web/#link-web-onevent-metadata-view-name-CONSENT)

We ask the user to consent to the privacy policy.

[`CREDENTIAL`](/docs/link/web/#link-web-onevent-metadata-view-name-CREDENTIAL)

Asking the user for their account credentials.

[`DOCUMENTARY_VERIFICATION`](/docs/link/web/#link-web-onevent-metadata-view-name-DOCUMENTARY-VERIFICATION)

The view requesting document verification in the identity verification flow (configured via "Fallback Settings" in the "Rulesets" section of the template editor).

[`ERROR`](/docs/link/web/#link-web-onevent-metadata-view-name-ERROR)

An error has occurred.

[`EXIT`](/docs/link/web/#link-web-onevent-metadata-view-name-EXIT)

Confirming if the user wishes to close Link.

[`INSTANT_MICRODEPOSIT_AUTHORIZED`](/docs/link/web/#link-web-onevent-metadata-view-name-INSTANT-MICRODEPOSIT-AUTHORIZED)

The user has authorized an instant micro-deposit to be sent to their account over the RTP or FedNow network with a 3-letter code to verify their account.

[`KYC_CHECK`](/docs/link/web/#link-web-onevent-metadata-view-name-KYC-CHECK)

The view representing the "know your customer" step in the identity verification flow.

[`LOADING`](/docs/link/web/#link-web-onevent-metadata-view-name-LOADING)

Link is making a request to our servers.

[`MFA`](/docs/link/web/#link-web-onevent-metadata-view-name-MFA)

The user is asked by the institution for additional MFA authentication.

[`NUMBERS`](/docs/link/web/#link-web-onevent-metadata-view-name-NUMBERS)

The user is asked to insert their account and routing numbers.

[`NUMBERS_SELECT_INSTITUTION`](/docs/link/web/#link-web-onevent-metadata-view-name-NUMBERS-SELECT-INSTITUTION)

The user goes through the Same Day micro-deposits flow with Reroute to Credentials.

[`OAUTH`](/docs/link/web/#link-web-onevent-metadata-view-name-OAUTH)

The user is informed they will authenticate with the financial institution via OAuth.

[`PROFILE_DATA_REVIEW`](/docs/link/web/#link-web-onevent-metadata-view-name-PROFILE-DATA-REVIEW)

The user is asked to review their profile data in the Layer flow.

[`RECAPTCHA`](/docs/link/web/#link-web-onevent-metadata-view-name-RECAPTCHA)

The user was presented with a Google reCAPTCHA to verify they are human.

[`RISK_CHECK`](/docs/link/web/#link-web-onevent-metadata-view-name-RISK-CHECK)

The risk check step in the identity verification flow (configured via "Risk Rules" in the "Rulesets" section of the template editor).

[`SAME_DAY_MICRODEPOSIT_AUTHORIZED`](/docs/link/web/#link-web-onevent-metadata-view-name-SAME-DAY-MICRODEPOSIT-AUTHORIZED)

The user has authorized a same day micro-deposit to be sent to their account over the ACH network with a 3-letter code to verify their account.

[`SCREENING`](/docs/link/web/#link-web-onevent-metadata-view-name-SCREENING)

The watchlist screening step in the identity verification flow.

[`SELECT_ACCOUNT`](/docs/link/web/#link-web-onevent-metadata-view-name-SELECT-ACCOUNT)

We ask the user to choose an account.

[`SELECT_AUTH_TYPE`](/docs/link/web/#link-web-onevent-metadata-view-name-SELECT-AUTH-TYPE)

The user is asked to choose whether to Link instantly or manually (i.e., with micro-deposits).

[`SELECT_BRAND`](/docs/link/web/#link-web-onevent-metadata-view-name-SELECT-BRAND)

The user is asked to select a brand, e.g. Bank of America. The brand selection interface occurs before the institution select pane and is only provided for large financial institutions with multiple online banking portals.

[`SELECT_INSTITUTION`](/docs/link/web/#link-web-onevent-metadata-view-name-SELECT-INSTITUTION)

We ask the user to choose their institution.

[`SELECT_SAVED_ACCOUNT`](/docs/link/web/#link-web-onevent-metadata-view-name-SELECT-SAVED-ACCOUNT)

The user is asked to select their saved accounts and/or new accounts for linking in the Link Returning User Experience (Remember Me) flow.

[`SELECT_SAVED_INSTITUTION`](/docs/link/web/#link-web-onevent-metadata-view-name-SELECT-SAVED-INSTITUTION)

The user is asked to pick a saved institution or link a new one in the Link Returning User Experience (Remember Me) flow.

[`SELFIE_CHECK`](/docs/link/web/#link-web-onevent-metadata-view-name-SELFIE-CHECK)

The view in the identity verification flow which uses the camera to confirm there is a real user present that matches their ID documents.

[`SUBMIT_PHONE`](/docs/link/web/#link-web-onevent-metadata-view-name-SUBMIT-PHONE)

The user is asked for their phone number in the Link Returning User Experience (Remember Me) flow.

[`UPLOAD_DOCUMENTS`](/docs/link/web/#link-web-onevent-metadata-view-name-UPLOAD-DOCUMENTS)

The user is asked to upload documents (for Income verification).

[`VERIFY_PHONE`](/docs/link/web/#link-web-onevent-metadata-view-name-VERIFY-PHONE)

The user is asked to verify their phone in the Link Returning User Experience (Remember Me) flow or the Layer flow. This screen will appear even if a non-OTP verification method is used.

[`VERIFY_SMS`](/docs/link/web/#link-web-onevent-metadata-view-name-VERIFY-SMS)

The SMS verification step in the identity verification flow.

[`request_id`](/docs/link/web/#link-web-onevent-metadata-request-id)

stringstring

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation. Emitted by: *all events*.

[`link_session_id`](/docs/link/web/#link-web-onevent-metadata-link-session-id)

stringstring

The `link_session_id` is a unique identifier for a single session of Link. It's always available and will stay constant throughout the flow. Emitted by: *all events*.

[`timestamp`](/docs/link/web/#link-web-onevent-metadata-timestamp)

stringstring

An ISO 8601 representation of when the event occurred. For example `2017-09-14T14:42:19.350Z`. Emitted by: *all events*.

[`selection`](/docs/link/web/#link-web-onevent-metadata-selection)

nullablestringnullable, string

The Auth Type Select flow type selected by the user. Possible values are `flow_type_manual` or `flow_type_instant`. Emitted by: `SELECT_AUTH_TYPE`.

onEvent example

```
import {
  PlaidLinkOnEvent,
  PlaidLinkOnEventMetadata,
  PlaidLinkStableEvent,
} from 'react-plaid-link';

const onEvent = useCallback<PlaidLinkOnEvent>(
  (
    eventName: PlaidLinkStableEvent | string,
    metadata: PlaidLinkOnEventMetadata,
  ) => {
    // log eventName and metadata
  },
  [],
);
```

Metadata schema

```
{
  error_type: 'ITEM_ERROR',
  error_code: 'INVALID_CREDENTIALS',
  error_message: 'the credentials were not correct',
  exit_status: null,
  institution_id: 'ins_4',
  institution_name: 'Wells Fargo',
  institution_search_query: 'wellsf',
  mfa_type: null,
  view_name: 'ERROR',
  request_id: 'm8MDnv9okwxFNBV',
  link_session_id: '30571e9b-d6c6-42ee-a7cf-c34768a8f62d',
  timestamp: '2017-09-14T14:42:19.350Z',
  selection: null,
}
```

=\*=\*=\*=

#### open()

Calling `open` will display the Consent Pane view to your user, starting the Link flow. Once `open` is called, you will begin receiving events via the [`onEvent` callback](#onevent).

open

**Properties**

This endpoint or method takes an empty request body.

open example

```
const { open, exit, ready } = usePlaidLink(config);

// Open Link
if (ready) {
  open();
}
```

=\*=\*=\*=

#### exit()

The `exit` function allows you to programmatically close Link. Calling `exit` will trigger either the [`onExit`](#onexit) or [`onSuccess`](#onsuccess) callbacks.

The `exit` function takes a single, optional argument, a configuration `Object`.

exit

**Properties**

[`force`](/docs/link/web/#link-web-exit-force)

booleanboolean

If `true`, Link will exit immediately. If `false`, or the option is not provided, an exit confirmation screen may be presented to the user.

Graceful exit example

```
const { open, exit, ready } = usePlaidLink(config);

// Graceful exit - Link may display a confirmation screen
// depending on how far the user is in the flow
exit();
```

Forced exit example

```
const { open, exit, ready } = usePlaidLink(config);

// Force exit - Link exits immediately
exit({ force: true });
```

=\*=\*=\*=

#### destroy()

The `destroy` function allows you to destroy the Link handler instance, properly removing any DOM artifacts that were created by it. Use `destroy()` when creating new replacement Link handler instances in the `onExit` callback.

destroy

**Properties**

This endpoint or method takes an empty request body.

Destroy example

```
// On unmount usePlaidLink hook automatically destroys any
// existing link instance
```

=\*=\*=\*=

#### submit()

The `submit` function is currently only used in the Layer product. It allows the client application to submit additional user-collected data to the Link flow (e.g. a user phone number).

submit

**Properties**

[`phone_number`](/docs/link/web/#link-web-submit-phone-number)

stringstring

The user's phone number.

[`date_of_birth`](/docs/link/web/#link-web-submit-date-of-birth)

stringstring

The user's date of birth. To be provided in the format "yyyy-mm-dd".

Submit example

```
const { open, exit, submit } = usePlaidLink(config);

// After collecting a user phone number...
submit({
    "phone_number": "+14155550123"
});
```

=\*=\*=\*=

#### OAuth

Using Plaid Link with an OAuth flow requires some additional setup instructions. For details, see the [OAuth Guide](/docs/link/oauth/).

=\*=\*=\*=

#### Supported browsers

Plaid officially supports Link on the latest versions of Chrome, Firefox, Safari, and Edge. Browsers are supported on Windows, Mac, Linux, iOS, and Android. Previous browser versions are also supported, as long as they are actively maintained; Plaid does not support browser versions that are no longer receiving patch updates, or that have been assigned official end of life (EOL) or end of support (EOS) status.

Ad-blocking software is not officially supported with Link web, and some ad-blockers have known to cause conflicts with Link.

=\*=\*=\*=

#### Example code in Plaid Pattern

For a real-life example of using Plaid Link for React, see [LaunchLink.tsx](https://github.com/plaid/pattern/blob/master/client/src/components/LaunchLink.tsx). This file illustrates the code for implementation of Plaid Link for React for the Node-based [Plaid Pattern](https://github.com/plaid/pattern) sample app.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
