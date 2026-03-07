---
title: "Link - Webview integrations | Plaid Docs"
source_url: "https://plaid.com/docs/link/webview/"
scraped_at: "2026-03-07T22:05:10+00:00"
---

# Link Webview SDK

#### Reference for integrating with the Link Webview JavaScript SDK

Using webviews to present Link is deprecated. If you're currently integrating this way, we recommend migrating to [Hosted Link](https://plaid.com/docs/link/hosted-link) or one Plaid's official SDKs for
[Android](https://plaid.com/docs/link/android/), [iOS](https://plaid.com/docs/link/ios/), and [React Native](https://plaid.com/docs/link/react-native/).

All Webview-based integrations need to extend the Webview handler for redirects in order to support connections to Chase. This can be accomplished with code samples for [iOS](https://github.com/plaid/plaid-link-examples/blob/master/webviews/wkwebview/wkwebview/LinkViewController.swift#L56-L72) and [Android](https://github.com/plaid/plaid-link-examples/blob/master/webviews/android/LinkWebview/app/src/main/java/com/example/linkwebview/MainActivity.kt#L89-L156). For more details, see [Extending webview instances to support certain institutions](/docs/link/oauth/#extending-webview-instances-to-support-certain-institutions) in the OAuth Guide and the [in-process Webview deprecation notice](https://plaid.docsend.com/view/h3qdupjusiwyjvv5). Webview-based integrations that do not extend the Webview handler are not supported by Chase and may be blocked by Chase in the future.

If your integration does not connect to Chase (for example, because you use only Identity Verification, Document Income, or Payroll Income, or you do not support end users in the US) you may ignore this warning.

=\*=\*=\*=

#### Overview

To integrate and use Plaid Link inside a Webview, we recommend starting with one of our sample Webview apps:

- [iOS WKWebView](https://github.com/plaid/plaid-link-examples/tree/master/webviews/wkwebview)
- [Android Webview](https://github.com/plaid/plaid-link-examples/tree/master/webviews/android)

Each example app is runnable (on both simulators and devices) and includes code to initialize Link and process events sent from Link to your app via HTTP redirects.

=\*=\*=\*=

#### Installation

Link is optimized to work within Webviews, including on iOS and Android. The Link initialization URL to use for Webviews is:

Link Webview initialization URL

```
https://cdn.plaid.com/link/v2/stable/link.html
```

The Link configuration options for a Webview integration are passed via querystring rather than via a client-side JavaScript call. See the [create](/docs/link/webview/#create) section below for details on the available initialization parameters.

##### Communication between Link and your app

Communication between the Webview and your app is handled by HTTP redirects rather than client-side JavaScript callbacks. These redirects should be intercepted by your app. The [example apps](https://github.com/plaid/plaid-link-examples) include sample code to do this.

All redirect URLs have the scheme `plaidlink`. The event type is communicated via the URL host and data is passed via the querystring.

HTTP redirect scheme

```
plaidlink://
```

There are three supported events, [`connected`](/docs/link/webview/#connected), [`exit`](/docs/link/webview/#exit), and [`event`](/docs/link/webview/#event), which are documented below.

=\*=\*=\*=

#### create()

create

**Properties**

[`isWebview`](/docs/link/webview/#link-webview-create-isWebview)

booleanboolean

Set to `true`, to trigger the Webview integration.

[`token`](/docs/link/webview/#link-webview-create-token)

stringstring

Specify a `link_token` to authenticate your app with Link. This is a short lived, one-time use token that should be unique for each Link session. In addition to the primary flow, a `link_token` can be configured to launch Link in [update mode](/docs/link/update-mode/). See the `/link/token/create` endpoint for a full list of configurations.

[`receivedRedirectUri`](/docs/link/webview/#link-webview-create-receivedRedirectUri)

stringstring

A `receivedRedirectUri` is required to support OAuth authentication flows when re-launching Link on a mobile device. Note that any unsafe ASCII characters in the `receivedRedirectUri` in the webview query string must be URL-encoded.

[`key`](/docs/link/webview/#link-webview-create-key)

deprecatedstringdeprecated, string

The `public_key` is no longer used for new implementations of Link. If your integration is still using a `public_key`, please contact Plaid support or your account manager.

Create example

```
https://cdn.plaid.com/link/v2/stable/link.html
  ?isWebview=true
  &token="GENERATED_LINK_TOKEN"
  &receivedRedirectUri=
```

=\*=\*=\*=

#### connected

The `connected` event is analogous to the `onSuccess` callback in [Link Web](/docs/link/web) and is sent when a user successfully completes Link. The following information is available from the querystring event:

connected

**Properties**

[`public_token`](/docs/link/webview/#link-webview-connected-public-token)

stringstring

Displayed once a user has successfully completed Link. If using Identity Verification or Beacon, this field will be `null`. If using Document Income or Payroll Income, the `public_token` will be returned, but is not used.

[`institution_name`](/docs/link/webview/#link-webview-connected-institution-name)

stringstring

The full institution name, such as `'Wells Fargo'`

[`institution_id`](/docs/link/webview/#link-webview-connected-institution-id)

stringstring

The Plaid institution identifier

[`accounts`](/docs/link/webview/#link-webview-connected-accounts)

objectobject

A JSON-stringified representation of the account(s) attached to the connected Item. If Account Select is enabled via the developer dashboard, `accounts` will only include selected accounts.

[`_id`](/docs/link/webview/#link-webview-connected-accounts--id)

stringstring

The Plaid `account_id`

[`meta`](/docs/link/webview/#link-webview-connected-accounts-meta)

objectobject

The account metadata

[`name`](/docs/link/webview/#link-webview-connected-accounts-meta-name)

stringstring

The official account name

[`number`](/docs/link/webview/#link-webview-connected-accounts-meta-number)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, it may also not match the mask that the bank displays to the user.

[`type`](/docs/link/webview/#link-webview-connected-accounts-type)

stringstring

The account type. See the [Account schema](/docs/api/accounts#account-type-schema) for a full list of possible values

[`subtype`](/docs/link/webview/#link-webview-connected-accounts-subtype)

stringstring

The account subtype. See the [Account schema](/docs/api/accounts#account-type-schema) for a full list of possible values

[`verification_status`](/docs/link/webview/#link-webview-connected-accounts-verification-status)

stringstring

When all Auth features are enabled by initializing Link with the user object, the accounts object includes an Item's `verification_status`. See Auth accounts for a full list of possible values.

[`class_type`](/docs/link/webview/#link-webview-connected-accounts-class-type)

nullablestringnullable, string

If micro-deposit verification is being used, indicates whether the account being verified is a `business` or `personal` account.

[`transfer_status`](/docs/link/webview/#link-webview-connected-transfer-status)

nullablestringnullable, string

The status of a transfer. Returned only when [Transfer UI](/docs/transfer/using-transfer-ui) is implemented.  

- `COMPLETE` – The transfer was completed.
- `INCOMPLETE` – The transfer could not be completed. For help, see [Troubleshooting Transfer UI](/docs/transfer/using-transfer-ui#troubleshooting-transfer-ui).
  
  

Possible values: `COMPLETE`, `INCOMPLETE`

[`link_session_id`](/docs/link/webview/#link-webview-connected-link-session-id)

stringstring

A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.

Connected event example

```
plaidlink://connected
  ?public_token=public-sandbox-fb7cca4a-82e6-4707
  &institution_id=ins_3
  &institution_name=Chase
  &accounts='[{"_id":"QPO8Jo8vdDHMepg41PBwckXm4KdK1yUdmXOwK", "meta": {"name":"Plaid Savings", "number": "0000"}, "subtype": "checking", "type": "depository" }]'
  &link_session_id=79e772be-547d-4c9c-8b76-4ac4ed4c441a
```

accounts schema example

```
"accounts": [
  {
    "_id": "ygPnJweommTWNr9doD6ZfGR6GGVQy7fyREmWy",
    "meta": {
      "name": "Plaid Checking",
      "number": "0000"
    },
    "type": "depository",
    "subtype": "checking",
    "verification_status": null
  },
  {
    "_id": "9ebEyJAl33FRrZNLBG8ECxD9xxpwWnuRNZ1V4",
    "meta": {
      "name": "Plaid Saving",
      "number": "1111"
    },
    "type": "depository",
    "subtype": "savings"
  }
  ...
]
```

=\*=\*=\*=

#### exit()

The `exit` event is analogous to the `onExit` callback and is sent when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. Note that on Android devices, an `exit` event will not be sent if the user exits Link through a system action, such as hitting the browser back button. The following information is available from the querystring:

exit

**Properties**

[`status`](/docs/link/webview/#link-webview-exit-status)

stringstring

The point at which the user exited the Link flow. One of the following values.

[`requires_questions`](/docs/link/webview/#link-webview-exit-status-requires-questions)

User prompted to answer security questions

[`requires_selections`](/docs/link/webview/#link-webview-exit-status-requires-selections)

User prompted to answer multiple choice question(s)

[`requires_code`](/docs/link/webview/#link-webview-exit-status-requires-code)

User prompted to provide a one-time passcode

[`choose_device`](/docs/link/webview/#link-webview-exit-status-choose-device)

User prompted to select a device on which to receive a one-time passcode

[`requires_credentials`](/docs/link/webview/#link-webview-exit-status-requires-credentials)

User prompted to provide credentials for the selected financial institution or has not yet selected a financial institution

[`requires_account_selection`](/docs/link/webview/#link-webview-exit-status-requires-account-selection)

User prompted to select one or more financial accounts to share

[`requires_oauth`](/docs/link/webview/#link-webview-exit-status-requires-oauth)

User prompted to enter an OAuth flow

[`institution_not_found`](/docs/link/webview/#link-webview-exit-status-institution-not-found)

User exited the Link flow on the institution selection pane. Typically this occurs after the user unsuccessfully (no results returned) searched for a financial institution. Note that this status does not necessarily indicate that the user was unable to find their institution, as it is used for all user exits that occur from the institution selection pane, regardless of other user behavior.

[`institution_not_supported`](/docs/link/webview/#link-webview-exit-status-institution-not-supported)

User exited the Link flow after discovering their selected institution is no longer supported by Plaid

[`error_type`](/docs/link/webview/#link-webview-exit-error-type)

StringString

A broad categorization of the error.

[`error_code`](/docs/link/webview/#link-webview-exit-error-code)

StringString

The particular error code. Each `error_type` has a specific set of `error_codes`.

[`error_message`](/docs/link/webview/#link-webview-exit-error-message)

StringString

A developer-friendly representation of the error code.

[`display_message`](/docs/link/webview/#link-webview-exit-display-message)

nullableStringnullable, String

A user-friendly representation of the error code. `null` if the error is not related to user action. This may change over time and is not safe for programmatic use.

[`institution_name`](/docs/link/webview/#link-webview-exit-institution-name)

stringstring

The full institution name, such as `Wells Fargo`

[`institution_id`](/docs/link/webview/#link-webview-exit-institution-id)

stringstring

The Plaid institution identifier

[`link_session_id`](/docs/link/webview/#link-webview-exit-link-session-id)

stringstring

A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.

[`request_id`](/docs/link/webview/#link-webview-exit-request-id)

stringstring

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation.

onExit example

```
plaidlink://exit
  ?status=requires_credentials
  &error_type=ITEM_ERROR
  &error_code=ITEM_LOGIN_REQUIRED
  &error_display_message=The%20credentials%20were%20not%20correct.%20Please%20try%20again.
  &error_message=the%20credentials%20were%20not%20correct
  &institution_id=ins_3
  &institution_name=Chase
  &link_session_id=79e772be-547d-4c9c-8b76-4ac4ed4c441a
  &request_id=m8MDnv9okwxFNBV
```

=\*=\*=\*=

#### event

The `event` message is analogous to the Link Web [`onEvent` callback](/docs/link/web#onevent) and is called as the user moves through the Link flow. The querystring will always contain all possible keys, though not all keys will have values. The `event_name` will dictate which keys are populated.

event

**Properties**

[`event_name`](/docs/link/webview/#link-webview-event-event-name)

stringstring

A string representing the event that has just occurred in the Link flow.

[`AUTO_SUBMIT_PHONE`](/docs/link/webview/#link-webview-event-event-name-AUTO-SUBMIT-PHONE)

The user was automatically sent an OTP code without a UI prompt. This event can only occur if the user's phone phone number was provided to Link via the `/link/token/create` call and the user has previously consented to receive OTP codes from Plaid.

[`BANK_INCOME_INSIGHTS_COMPLETED`](/docs/link/webview/#link-webview-event-event-name-BANK-INCOME-INSIGHTS-COMPLETED)

The user has completed the Assets and Bank Income Insights flow.

[`CLOSE_OAUTH`](/docs/link/webview/#link-webview-event-event-name-CLOSE-OAUTH)

The user closed the third-party website or mobile app without completing the OAuth flow.

[`CONNECT_NEW_INSTITUTION`](/docs/link/webview/#link-webview-event-event-name-CONNECT-NEW-INSTITUTION)

The user has chosen to link a new institution instead of linking a saved institution. This event is only emitted in the Link Returning User Experience flow.

[`ERROR`](/docs/link/webview/#link-webview-event-event-name-ERROR)

A recoverable error occurred in the Link flow, see the `error_code` metadata.

[`EXIT`](/docs/link/webview/#link-webview-event-event-name-EXIT)

The user has exited without completing the Link flow and the [exit](#exit) callback is fired.

[`FAIL_OAUTH`](/docs/link/webview/#link-webview-event-event-name-FAIL-OAUTH)

The user encountered an error while completing the third-party's OAuth login flow.

[`HANDOFF`](/docs/link/webview/#link-webview-event-event-name-HANDOFF)

The user has exited Link after successfully linking an Item.

[`IDENTITY_MATCH_FAILED`](/docs/link/webview/#link-webview-event-event-name-IDENTITY-MATCH-FAILED)

An Identity Match check configured via the Account Verification Dashboard failed the Identity Match rules and did not detect a match.

[`IDENTITY_MATCH_PASSED`](/docs/link/webview/#link-webview-event-event-name-IDENTITY-MATCH-PASSED)

An Identity Match check configured via the Account Verification Dashboard passed the Identity Match rules and detected a match.

[`INSTANT_MICRODEPOSIT_AUTHORIZED`](/docs/link/webview/#link-webview-event-event-name-INSTANT-MICRODEPOSIT-AUTHORIZED)

The user has authorized an instant micro-deposit to be sent to their account over the RTP or FedNow network with a 3-letter code to verify their account.

[`MATCHED_SELECT_INSTITUTION`](/docs/link/webview/#link-webview-event-event-name-MATCHED-SELECT-INSTITUTION)

The user selected an institution that was presented as a matched institution. This event can be emitted if [Embedded Institution Search](https://plaid.com/docs/link/embedded-institution-search/) is being used, if the institution was surfaced as a matched institution likely to have been linked to Plaid by a returning user, or if the institution's `routing_number` was provided when calling `/link/token/create`. For details on which scenario is triggering the event, see `metadata.matchReason`.

[`OPEN`](/docs/link/webview/#link-webview-event-event-name-OPEN)

The user has opened Link.

[`OPEN_MY_PLAID`](/docs/link/webview/#link-webview-event-event-name-OPEN-MY-PLAID)

The user has opened my.plaid.com. This event is only sent when Link is initialized with Assets as a product

[`OPEN_OAUTH`](/docs/link/webview/#link-webview-event-event-name-OPEN-OAUTH)

The user has navigated to a third-party website or mobile app in order to complete the OAuth login flow.

[`SAME_DAY_MICRODEPOSIT_AUTHORIZED`](/docs/link/webview/#link-webview-event-event-name-SAME-DAY-MICRODEPOSIT-AUTHORIZED)

The user has authorized a same day micro-deposit to be sent to their account over the ACH network with a 3-letter code to verify their account.

[`SEARCH_INSTITUTION`](/docs/link/webview/#link-webview-event-event-name-SEARCH-INSTITUTION)

The user has searched for an institution.

[`SELECT_BRAND`](/docs/link/webview/#link-webview-event-event-name-SELECT-BRAND)

The user selected a brand, e.g. Bank of America. The `SELECT_BRAND` event is only emitted for large financial institutions with multiple online banking portals.

[`SELECT_DEGRADED_INSTITUTION`](/docs/link/webview/#link-webview-event-event-name-SELECT-DEGRADED-INSTITUTION)

The user selected an institution with a `DEGRADED` health status and was shown a corresponding message.

[`SELECT_DOWN_INSTITUTION`](/docs/link/webview/#link-webview-event-event-name-SELECT-DOWN-INSTITUTION)

The user selected an institution with a `DOWN` health status and was shown a corresponding message.

[`SELECT_FILTERED_INSTITUTION`](/docs/link/webview/#link-webview-event-event-name-SELECT-FILTERED-INSTITUTION)

The user selected an institution Plaid does not support all requested products for and was shown a corresponding message.

[`SELECT_INSTITUTION`](/docs/link/webview/#link-webview-event-event-name-SELECT-INSTITUTION)

The user selected an institution.

[`SKIP_SUBMIT_PHONE`](/docs/link/webview/#link-webview-event-event-name-SKIP-SUBMIT-PHONE)

The user has opted to not provide their phone number to Plaid. This event is only emitted in the Link Returning User Experience flow.

[`SUBMIT_ACCOUNT_NUMBER`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-ACCOUNT-NUMBER)

The user has submitted an account number. This event emits the `account_number_mask` metadata to indicate the mask of the account number the user provided.

[`SUBMIT_CREDENTIALS`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-CREDENTIALS)

The user has submitted credentials.

[`SUBMIT_DOCUMENTS`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-DOCUMENTS)

The user is being prompted to submit documents for an Income verification flow.

[`SUBMIT_DOCUMENTS_ERROR`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-DOCUMENTS-ERROR)

The user encountered an error when submitting documents for an Income verification flow.

[`SUBMIT_DOCUMENTS_SUCCESS`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-DOCUMENTS-SUCCESS)

The user has successfully submitted documents for an Income verification flow.

[`SUBMIT_MFA`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-MFA)

The user has submitted MFA.

[`SUBMIT_OTP`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-OTP)

The user has submitted an OTP code during the phone number verification flow. This event is only emitted in the Link Returning User Experience (Remember Me) flow or Layer flow. This event will not be emitted if the phone number is verified via SNA.

[`SUBMIT_PHONE`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-PHONE)

The user has submitted their phone number. This event is only emitted in the Link Returning User Experience (Remember Me) flow.

[`SUBMIT_ROUTING_NUMBER`](/docs/link/webview/#link-webview-event-event-name-SUBMIT-ROUTING-NUMBER)

The user has submitted a routing number. This event emits the `routing_number` metadata to indicate user's routing number.

[`TRANSITION_VIEW`](/docs/link/webview/#link-webview-event-event-name-TRANSITION-VIEW)

The `TRANSITION_VIEW` event indicates that the user has moved from one view to the next.

[`VERIFY_PHONE`](/docs/link/webview/#link-webview-event-event-name-VERIFY-PHONE)

The user has successfully verified their phone number, via either OTP or SNA. This event is only emitted by the Link Returning User Experience (Remember Me) flow and the Layer flow.

[`VIEW_DATA_TYPES`](/docs/link/webview/#link-webview-event-event-name-VIEW-DATA-TYPES)

The user has viewed data types on the data transparency consent pane.

[`error_type`](/docs/link/webview/#link-webview-event-error-type)

nullablestringnullable, string

The error type that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`error_code`](/docs/link/webview/#link-webview-event-error-code)

nullablestringnullable, string

The error code that the user encountered. Emitted by `ERROR`, `EXIT`.

[`error_message`](/docs/link/webview/#link-webview-event-error-message)

nullablestringnullable, string

The error message that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`exit_status`](/docs/link/webview/#link-webview-event-exit-status)

nullablestringnullable, string

The status key indicates the point at which the user exited the Link flow. Emitted by: `EXIT`

[`institution_id`](/docs/link/webview/#link-webview-event-institution-id)

nullablestringnullable, string

The ID of the selected institution. Emitted by: *all events*.

[`institution_name`](/docs/link/webview/#link-webview-event-institution-name)

nullablestringnullable, string

The name of the selected institution. Emitted by: *all events*.

[`institution_search_query`](/docs/link/webview/#link-webview-event-institution-search-query)

nullablestringnullable, string

The query used to search for institutions. Emitted by: `SEARCH_INSTITUTION`.

[`is_update_mode`](/docs/link/webview/#link-webview-event-is-update-mode)

nullablestringnullable, string

Indicates if the current Link session is an update mode session. Emitted by: `OPEN`.

[`match_reason`](/docs/link/webview/#link-webview-event-match-reason)

nullablestringnullable, string

The reason this institution was matched.
This will be either `returning_user` or `routing_number` if emitted by: `MATCHED_SELECT_INSTITUTION`.
Otherwise, this will be `SAVED_INSTITUTION` or `AUTO_SELECT_SAVED_INSTITUTION` if emitted by: `SELECT_INSTITUTION`.

[`mfa_type`](/docs/link/webview/#link-webview-event-mfa-type)

nullablestringnullable, string

If set, the user has encountered one of the following MFA types: `code`, `device`, `questions`, `selections`. Emitted by: `SUBMIT_MFA` and `TRANSITION_VIEW` when `view_name` is `MFA`

[`view_name`](/docs/link/webview/#link-webview-event-view-name)

nullablestringnullable, string

The name of the view that is being transitioned to. Emitted by: `TRANSITION_VIEW`.

[`ACCEPT_TOS`](/docs/link/webview/#link-webview-event-view-name-ACCEPT-TOS)

The view showing Terms of Service in the identity verification flow.

[`CONNECTED`](/docs/link/webview/#link-webview-event-view-name-CONNECTED)

The user has connected their account.

[`CONSENT`](/docs/link/webview/#link-webview-event-view-name-CONSENT)

We ask the user to consent to the privacy policy.

[`CREDENTIAL`](/docs/link/webview/#link-webview-event-view-name-CREDENTIAL)

Asking the user for their account credentials.

[`DOCUMENTARY_VERIFICATION`](/docs/link/webview/#link-webview-event-view-name-DOCUMENTARY-VERIFICATION)

The view requesting document verification in the identity verification flow (configured via "Fallback Settings" in the "Rulesets" section of the template editor).

[`ERROR`](/docs/link/webview/#link-webview-event-view-name-ERROR)

An error has occurred.

[`EXIT`](/docs/link/webview/#link-webview-event-view-name-EXIT)

Confirming if the user wishes to close Link.

[`INSTANT_MICRODEPOSIT_AUTHORIZED`](/docs/link/webview/#link-webview-event-view-name-INSTANT-MICRODEPOSIT-AUTHORIZED)

The user has authorized an instant micro-deposit to be sent to their account over the RTP or FedNow network with a 3-letter code to verify their account.

[`KYC_CHECK`](/docs/link/webview/#link-webview-event-view-name-KYC-CHECK)

The view representing the "know your customer" step in the identity verification flow.

[`LOADING`](/docs/link/webview/#link-webview-event-view-name-LOADING)

Link is making a request to our servers.

[`MFA`](/docs/link/webview/#link-webview-event-view-name-MFA)

The user is asked by the institution for additional MFA authentication.

[`NUMBERS`](/docs/link/webview/#link-webview-event-view-name-NUMBERS)

The user is asked to insert their account and routing numbers.

[`OAUTH`](/docs/link/webview/#link-webview-event-view-name-OAUTH)

The user is informed they will authenticate with the financial institution via OAuth.

[`RECAPTCHA`](/docs/link/webview/#link-webview-event-view-name-RECAPTCHA)

The user was presented with a Google reCAPTCHA to verify they are human.

[`RISK_CHECK`](/docs/link/webview/#link-webview-event-view-name-RISK-CHECK)

The risk check step in the identity verification flow (configured via "Risk Rules" in the "Rulesets" section of the template editor).

[`SAME_DAY_MICRODEPOSIT_AUTHORIZED`](/docs/link/webview/#link-webview-event-view-name-SAME-DAY-MICRODEPOSIT-AUTHORIZED)

The user has authorized a same day micro-deposit to be sent to their account over the ACH network with a 3-letter code to verify their account.

[`SCREENING`](/docs/link/webview/#link-webview-event-view-name-SCREENING)

The watchlist screening step in the identity verification flow.

[`SELECT_ACCOUNT`](/docs/link/webview/#link-webview-event-view-name-SELECT-ACCOUNT)

We ask the user to choose an account.

[`SELECT_BRAND`](/docs/link/webview/#link-webview-event-view-name-SELECT-BRAND)

The user selected a brand, e.g. Bank of America. The brand selection interface occurs before the institution select pane and is only provided for large financial institutions with multiple online banking portals.

[`SELECT_INSTITUTION`](/docs/link/webview/#link-webview-event-view-name-SELECT-INSTITUTION)

We ask the user to choose their institution.

[`SELECT_SAVED_ACCOUNT`](/docs/link/webview/#link-webview-event-view-name-SELECT-SAVED-ACCOUNT)

The user is asked to select their saved accounts and/or new accounts for linking in the Link Returning User Experience (Remember Me) flow.

[`SELECT_SAVED_INSTITUTION`](/docs/link/webview/#link-webview-event-view-name-SELECT-SAVED-INSTITUTION)

The user is asked to pick a saved institution or link a new one in the Link Returning User Experience (Remember Me) flow.

[`SELFIE_CHECK`](/docs/link/webview/#link-webview-event-view-name-SELFIE-CHECK)

The view in the identity verification flow which uses the camera to confirm there is a real user present that matches their ID documents.

[`SUBMIT_PHONE`](/docs/link/webview/#link-webview-event-view-name-SUBMIT-PHONE)

The user is asked for their phone number in the Link Returning User Experience (Remember Me) flow.

[`VERIFY_PHONE`](/docs/link/webview/#link-webview-event-view-name-VERIFY-PHONE)

The user is asked to verify their phone in the Link Returning User Experience (Remember Me) flow or the Layer flow. This screen will appear even if a non-OTP verification method is used.

[`VERIFY_SMS`](/docs/link/webview/#link-webview-event-view-name-VERIFY-SMS)

The SMS verification step in the identity verification flow.

[`request_id`](/docs/link/webview/#link-webview-event-request-id)

stringstring

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation. Emitted by: *all events*.

[`link_session_id`](/docs/link/webview/#link-webview-event-link-session-id)

stringstring

The `link_session_id` is a unique identifier for a single session of Link. It's always available and will stay constant throughout the flow. Emitted by: *all events*.

[`timestamp`](/docs/link/webview/#link-webview-event-timestamp)

stringstring

An ISO 8601 representation of when the event occurred. For example `2017-09-14T14:42:19.350Z`. Emitted by: *all events*.

[`selection`](/docs/link/webview/#link-webview-event-selection)

nullablestringnullable, string

The Auth Type Select flow type selected by the user. Possible values are `flow_type_manual` or `flow_type_instant`. Emitted by: `SELECT_AUTH_TYPE`.

event example

```
plaidlink://event
  &event_name=SELECT_INSTITUTION
  ?error_type=ITEM_ERROR
  &error_code=ITEM_LOGIN_REQUIRED
  &error_message=the%20credentials%20were%20not%20correct
  &exit_status
  &institution_id=ins_55
  &institution_name=HSBC
  &institution_search_query=h
  &mfa_type
  &view_name=ERROR
  &request_id
  &link_session_id=821f45a8-854a-4dbb-8e5f-73f75350e7e7
  &timestamp=2018-10-05T15%3A22%3A50.542Z
```

=\*=\*=\*=

#### OAuth

Using Plaid Link with an OAuth flow requires some additional setup instructions. For details, see the [OAuth Guide](/docs/link/oauth/).

=\*=\*=\*=

#### Supported platforms

Plaid officially supports WKWebView on iOS 10 or later and Chrome WebView on Android 4.4 or later.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
