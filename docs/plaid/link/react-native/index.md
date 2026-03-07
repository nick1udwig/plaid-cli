---
title: "Link - React Native | Plaid Docs"
source_url: "https://plaid.com/docs/link/react-native/"
scraped_at: "2026-03-07T22:05:08+00:00"
---

# Link React Native SDK

#### Reference for integrating with the Link React Native SDK

This guide covers the latest major version of the Link React Native SDK, which is version 12.x.x. For information on migrating from older versions, see [Migration guides](/docs/link/react-native/#migration-guides).

#### Overview

Prefer to learn with code examples? A [GitHub repo](https://github.com/plaid/tiny-quickstart/tree/main/react_native) showing a working example Link implementation is available for this topic.

To get started with [Plaid Link for React Native](https://github.com/plaid/react-native-plaid-link-sdk) youʼll want to sign up
for [free API keys](https://dashboard.plaid.com/developers/keys) through the Plaid Dashboard.

#### Requirements

- React Native Version `0.66.0` or higher

New versions of the [React Native SDK](https://github.com/plaid/react-native-plaid-link-sdk) are released frequently. Major releases occur annually. The Link SDK uses Semantic Versioning, ensuring that all non-major releases are non-breaking, backwards-compatible updates. We recommend you update regularly (at least once a quarter, and ideally once a month) to ensure the best Plaid Link experience in your application.

SDK versions are supported for two years; with each major SDK release, Plaid will stop officially supporting any previous SDK versions that are more than two years old. While these older versions are expected to continue to work without disruption, Plaid will not provide assistance with unsupported SDK versions.

##### Version Compatibility

| React Native SDK | Android SDK | iOS SDK | Status |
| --- | --- | --- | --- |
| 12.x.x | 5.0+ | >=6.0.0 | Active, supports Xcode 16 |
| 11.x.x | 4.1.1+ | >=5.1.0 | Active, supports Xcode 15 |
| 10.x.x | 3.10.1+ | >=4.1.0 | Active, supports Xcode 14 |
| 9.x.x | 3.10.1+ | >=4.1.0 | Deprecated, supports Xcode 14 |

#### Getting Started

##### Installing the SDK

In your react-native project directory, run:

```
npm install --save react-native-plaid-link-sdk
```

##### iOS Setup

Add Plaid to your project's Podfile as follows (likely located at ios/Podfile).

```
pod 'Plaid', '~> <insert latest version>'
```

Autolinking should install the CocoaPods dependencies for iOS project. If it fails you can run

```
cd ios && bundle install && bundle exec pod install
```

##### Android Setup

Requirements:

- Android 5.0 (API level 21) and above.
- Your compileSdkVersion must be 35.
- Android gradle plugin 4.x and above.

Autolinking should handle all of the Android setup.

- Register your Android package name in the [Dashboard](https://dashboard.plaid.com/developers/api). This is required in order to connect to OAuth institutions (which includes most major banks).

###### Sample app

For a sample app that demonstrates a minimal integration with the React Native Plaid Link SDK, see the [Tiny Quickstart (React Native)](https://github.com/plaid/tiny-quickstart/tree/main/react_native).

#### Opening Link

Before you can open Link, you need to first create a `link_token`. A `link_token` can be configured for
different Link flows and is used to control much of Link's behavior. To see how to create a new
`link_token`, see the API Reference entry for [`/link/token/create`](/docs/api/link/#linktokencreate).
If your React Native application will be used on Android, the `link/token/create` call should include the `android_package_name` parameter.
Each time you open Link, you will need to get a new `link_token` from your server.

Next, open Link via the `create` and `open` methods. These functions require version 11.6 or later of the Plaid React Native SDK.

If using a version earlier than 11.6, you must open Link with the legacy `PlaidLink` component, which configures Link and registers a callback in a single component. This approach has higher user-facing latency than using the `create` and `open` methods.

=\*=\*=\*=

#### create()

You can initiate the Link preloading process by invoking the `create` function. After calling `create`, call `open` to open Link. This function requires SDK version 11.6 or later.

create

**Properties**

[`linkTokenConfiguration`](/docs/link/react-native/#link-react_native-create-linkTokenConfiguration)

LinkTokenConfigurationLinkTokenConfiguration

A configuration used to open Link with a Link Token.

[`token`](/docs/link/react-native/#link-react_native-create-linkTokenConfiguration-token)

stringstring

The `link_token` to be used to authenticate your app with Link. The `link_token` is created by calling `/link/token/create` and is a short lived, one-time use token that should be unique for each Link session. In addition to the primary flow, a `link_token` can be configured to launch Link in [update mode](/docs/link/update-mode/). See the `/link/token/create` endpoint for a full list of configurations.

[`noLoadingState`](/docs/link/react-native/#link-react_native-create-linkTokenConfiguration-noLoadingState)

booleanboolean

Hides native activity indicator if true.

[`logLevel`](/docs/link/react-native/#link-react_native-create-linkTokenConfiguration-logLevel)

LinkLogLevelLinkLogLevel

Set the level at which to log statements  
  

Possible values: `DEBUG`, `INFO`, `WARN`, `ERROR`

```
<TouchableOpacity
  style={styles.button}
  onPress={() => {
      create({token: linkToken});
      setDisabled(false);
    }
  }>
  <Text style={styles.button}>Create Link</Text>
</TouchableOpacity>
```

=\*=\*=\*=

#### open()

After calling `create`, you can subsequently invoke the `open` function. Note that maximizing the delay between these two calls will reduce latency for your users by allowing Link more time to load. This function requires SDK version 11.6 or later.

open

**Properties**

[`onSuccess`](/docs/link/react-native/#link-react_native-open-onSuccess)

LinkSuccessListenerLinkSuccessListener

A function that is called when a user successfully links an Item. The function should expect one argument. See [onSuccess](#onsuccess).

[`onExit`](/docs/link/react-native/#link-react_native-open-onExit)

LinkExitListenerLinkExitListener

A function that is called when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. The function should expect one argument. See [onExit](#onexit).

[`iosPresentationStyle`](/docs/link/react-native/#link-react_native-open-iosPresentationStyle)

LinkIOSPresentationStyleLinkIOSPresentationStyle

The presentation style to use on iOS. Defaults to `MODAL`.  
  

Possible values: `MODAL`, `FULL_SCREEN`

[`logLevel`](/docs/link/react-native/#link-react_native-open-logLevel)

LinkLogLevelLinkLogLevel

Set the level at which to log statements  
  

Possible values: `DEBUG`, `INFO`, `WARN`, `ERROR`

```
<TouchableOpacity
  disabled={disabled}
  style={disabled ? styles.disabledButton : styles.button}
  onPress={() => {
    const openProps = {
      onSuccess: (success: LinkSuccess) => {
      console.log(success);
      },
      onExit: (linkExit: LinkExit) => {
        console.log(linkExit);
      },
    };
    open(openProps);
    setDisabled(true);
  }}>
  <Text style={styles.button}>Open Link</Text>
</TouchableOpacity>
```

=\*=\*=\*=

#### PlaidLink

PlaidLink is a React component used to open Link from a React Native application. PlaidLink renders a [Pressable](https://reactnative.dev/docs/pressable) component, which wraps the component you provide and intercepts onPress events to open Link. PlaidLink is an older alternative to the `create` and `open` methods, which offer reduced Link latency and improved performance.

plaidLink

**Properties**

[`onSuccess`](/docs/link/react-native/#link-react_native-plaidlink-onSuccess)

LinkSuccessListenerLinkSuccessListener

A function that is called when a user successfully links an Item. The function should expect one argument. See [onSuccess](#onsuccess).

[`tokenConfig`](/docs/link/react-native/#link-react_native-plaidlink-tokenConfig)

LinkTokenConfigurationLinkTokenConfiguration

A configuration used to open Link with a Link Token.

[`token`](/docs/link/react-native/#link-react_native-plaidlink-tokenConfig-token)

stringstring

The `link_token` to be used to authenticate your app with Link. The `link_token` is created by calling `/link/token/create` and is a short lived, one-time use token that should be unique for each Link session. In addition to the primary flow, a `link_token` can be configured to launch Link in [update mode](/docs/link/update-mode/). See the `/link/token/create` endpoint for a full list of configurations.

[`logLevel`](/docs/link/react-native/#link-react_native-plaidlink-tokenConfig-logLevel)

LinkLogLevelLinkLogLevel

Set the level at which to log statements  
  

Possible values: `DEBUG`, `INFO`, `WARN`, `ERROR`

[`extraParams`](/docs/link/react-native/#link-react_native-plaidlink-tokenConfig-extraParams)

Record<string, any>Record<string, any>

You do not need to use this field unless specifically instructed to by Plaid.

[`onExit`](/docs/link/react-native/#link-react_native-plaidlink-onExit)

LinkExitListenerLinkExitListener

A function that is called when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. The function should expect one argument. See [onExit](#onexit).

[`children`](/docs/link/react-native/#link-react_native-plaidlink-children)

React.ReactNodeReact.ReactNode

The underlying component to render

```
<PlaidLink
  tokenConfig={{
    token: '#GENERATED_LINK_TOKEN#',
  }}
  onSuccess={(success: LinkSuccess) => {
    console.log(success);
  }}
  onExit={(exit: LinkExit) => {
    console.log(exit);
  }}
>
  <Text>Add Account</Text>
</PlaidLink>
```

=\*=\*=\*=

#### onSuccess

The method is called when a user successfully links an Item. The onSuccess handler returns a `LinkConnection` class that includes the `public_token`, and additional Link metadata in the form of a `LinkConnectionMetadata` class.

onSuccess

**Properties**

[`publicToken`](/docs/link/react-native/#link-react_native-onsuccess-publicToken)

StringString

Displayed once a user has successfully completed Link. If using Identity Verification or Beacon, this field will be `null`. If using Document Income or Payroll Income, the `public_token` will be returned, but is not used.

[`metadata`](/docs/link/react-native/#link-react_native-onsuccess-metadata)

LinkSuccessMetadataLinkSuccessMetadata

Displayed once a user has successfully completed Link.

[`accounts`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts)

Array<LinkAccount>Array<LinkAccount>

A list of accounts attached to the connected Item. If Account Select is enabled via the developer dashboard, `accounts` will only include selected accounts.

[`id`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-id)

stringstring

The Plaid `account_id`

[`name`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-name)

nullablestringnullable, string

The official account name

[`mask`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, it may also not match the mask that the bank displays to the user.

[`type`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-type)

LinkAccountTypeLinkAccountType

The account type. See the [Account schema](/docs/api/accounts#account-type-schema) for a full list of possible values

[`subtype`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-subtype)

LinkAccountSubtypeLinkAccountSubtype

The account subtype. See the [Account schema](/docs/api/accounts#account-type-schema) for a full list of possible values

[`verification_status`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status)

nullableLinkAccountVerificationStatusnullable, LinkAccountVerificationStatus

Indicates an Item's micro-deposit-based verification or database verification status.

[`pending_automatic_verification`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-pending-automatic-verification)

The Item is pending automatic verification

[`pending_manual_verification`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-pending-manual-verification)

The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the deposit.

[`automatically_verified`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-automatically-verified)

The Item has successfully been automatically verified

[`manually_verified`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-manually-verified)

The Item has successfully been manually verified

[`verification_expired`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-verification-expired)

Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.

[`verification_failed`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-verification-failed)

The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.

[`database_matched`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-database-matched)

The Item has successfully been verified using Plaid's data sources.

[`database_insights_pending`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-database-insights-pending)

The Database Insights result is pending and will be available upon Auth request.

[`null`](/docs/link/react-native/#link-react_native-onsuccess-metadata-accounts-verification-status-null)

Neither micro-deposit-based verification nor database verification are being used for the Item.

[`institution`](/docs/link/react-native/#link-react_native-onsuccess-metadata-institution)

nullableobjectnullable, object

An institution object. If the Item was created via Same-Day micro-deposit verification, will be `null`.

[`name`](/docs/link/react-native/#link-react_native-onsuccess-metadata-institution-name)

stringstring

The full institution name, such as `'Wells Fargo'`

[`id`](/docs/link/react-native/#link-react_native-onsuccess-metadata-institution-id)

stringstring

The Plaid institution identifier

[`linkSessionId`](/docs/link/react-native/#link-react_native-onsuccess-metadata-linkSessionId)

nullableStringnullable, String

A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.

[`metadataJson`](/docs/link/react-native/#link-react_native-onsuccess-metadata-metadataJson)

nullableMapnullable, Map

The data directly returned from the server with no client side changes.

onSuccess example

```
const onSuccess = (success: LinkSuccess) => {
  // If using Item-based products, exchange public_token
  // for access_token
  fetch('https://yourserver.com/exchange_public_token', {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      publicToken: linkSuccess.publicToken,
      accounts: linkSuccess.metadata.accounts,
      institution: linkSuccess.metadata.institution,
      linkSessionId: linkSuccess.metadata.linkSessionId,
    }),
  });
};
```

=\*=\*=\*=

#### onExit

The `onExit` handler is called when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. The `PlaidError` returned from the `onExit` handler is meant to help you guide your users after they have exited Link. We recommend storing the error and metadata information server-side in a way that can be associated with the user. You’ll also need to include this and any other relevant info in Plaid Support requests for the user.

onExit

**Properties**

[`error`](/docs/link/react-native/#link-react_native-onexit-error)

LinkErrorLinkError

An object that contains the error type, error code, and error message of the error that was last encountered by the user. If no error was encountered, `error` will be `null`.

[`displayMessage`](/docs/link/react-native/#link-react_native-onexit-error-displayMessage)

nullableStringnullable, String

A user-friendly representation of the error code. `null` if the error is not related to user action. This may change over time and is not safe for programmatic use.

[`errorCode`](/docs/link/react-native/#link-react_native-onexit-error-errorCode)

LinkErrorCodeLinkErrorCode

The particular error code. Each `errorType` has a specific set of `errorCodes`. A code of 499 indicates a client-side exception.

[`errorType`](/docs/link/react-native/#link-react_native-onexit-error-errorType)

LinkErrorTypeLinkErrorType

A broad categorization of the error.

[`errorMessage`](/docs/link/react-native/#link-react_native-onexit-error-errorMessage)

StringString

A developer-friendly representation of the error code.

[`errorJson`](/docs/link/react-native/#link-react_native-onexit-error-errorJson)

nullableStringnullable, String

The data directly returned from the server with no client side changes.

[`metadata`](/docs/link/react-native/#link-react_native-onexit-metadata)

Map<String, Object>Map<String, Object>

An object containing information about the exit event

[`linkSessionId`](/docs/link/react-native/#link-react_native-onexit-metadata-linkSessionId)

nullableStringnullable, String

A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.

[`institution`](/docs/link/react-native/#link-react_native-onexit-metadata-institution)

nullableLinkInstitutionnullable, LinkInstitution

An institution object. If the Item was created via Same-Day micro-deposit verification, will be `null`.

[`name`](/docs/link/react-native/#link-react_native-onexit-metadata-institution-name)

stringstring

The full institution name, such as `'Wells Fargo'`

[`id`](/docs/link/react-native/#link-react_native-onexit-metadata-institution-id)

stringstring

The Plaid institution identifier

[`status`](/docs/link/react-native/#link-react_native-onexit-metadata-status)

nullableLinkExitMetadataStatusnullable, LinkExitMetadataStatus

The point at which the user exited the Link flow. One of the following values

[`requires_questions`](/docs/link/react-native/#link-react_native-onexit-metadata-status-requires-questions)

User prompted to answer security questions

[`requires_selections`](/docs/link/react-native/#link-react_native-onexit-metadata-status-requires-selections)

User prompted to answer multiple choice question(s)

[`requires_recaptcha`](/docs/link/react-native/#link-react_native-onexit-metadata-status-requires-recaptcha)

User prompted to solve a reCAPTCHA challenge

[`requires_code`](/docs/link/react-native/#link-react_native-onexit-metadata-status-requires-code)

User prompted to provide a one-time passcode

[`choose_device`](/docs/link/react-native/#link-react_native-onexit-metadata-status-choose-device)

User prompted to select a device on which to receive a one-time passcode

[`requires_credentials`](/docs/link/react-native/#link-react_native-onexit-metadata-status-requires-credentials)

User prompted to provide credentials for the selected financial institution or has not yet selected a financial institution

[`requires_account_selection`](/docs/link/react-native/#link-react_native-onexit-metadata-status-requires-account-selection)

User prompted to select one or more financial accounts to share

[`requires_oauth`](/docs/link/react-native/#link-react_native-onexit-metadata-status-requires-oauth)

User prompted to enter an OAuth flow

[`institution_not_found`](/docs/link/react-native/#link-react_native-onexit-metadata-status-institution-not-found)

User exited the Link flow on the institution selection pane. Typically this occurs after the user unsuccessfully (no results returned) searched for a financial institution. Note that this status does not necessarily indicate that the user was unable to find their institution, as it is used for all user exits that occur from the institution selection pane, regardless of other user behavior.

[`institution_not_supported`](/docs/link/react-native/#link-react_native-onexit-metadata-status-institution-not-supported)

User exited the Link flow after discovering their selected institution is no longer supported by Plaid

[`unknown`](/docs/link/react-native/#link-react_native-onexit-metadata-status-unknown)

An exit status that is not handled by the current version of the SDK

[`requestId`](/docs/link/react-native/#link-react_native-onexit-metadata-requestId)

nullableStringnullable, String

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation

[`metadataJson`](/docs/link/react-native/#link-react_native-onexit-metadata-metadataJson)

nullableMapnullable, Map

The data directly returned from the server with no client side changes.

onExit example

```
const onExit = (linkExit: LinkExit) => {
  supportHandler.report({
    error: linkExit.error,
    institution: linkExit.metadata.institution,
    linkSessionId: linkExit.metadata.linkSessionId,
    requestId: linkExit.metadata.requestId,
    status: linkExit.metadata.status,
  });
};
```

=\*=\*=\*=

#### onEvent

The React Native Plaid module emits `onEvent` events throughout the account linking process.
To receive these events use the `usePlaidEmitter` hook.

The `onEvent` callback is called at certain points in the Link flow. Unlike the handlers for `onSuccess` and `onExit`, the `onEvent` handler is initialized as a global lambda passed to the `Plaid` class. `OPEN`, `LAYER_READY`, `LAYER_NOT_AVAILABLE`, and `LAYER_AUTOFILL_NOT_AVAILABLE` events will be sent immediately in real-time, and remaining events will be sent when the Link session is finished and `onSuccess` or `onExit` is called. Callback ordering is not guaranteed; `onEvent` callbacks may fire before, after, or surrounding the `onSuccess` or `onExit` callback, and event callbacks are not guaranteed to fire in the order in which they occurred. If you need the exact time when an event happened, use the `timestamp` property.

The following `onEvent` callbacks are stable, which means that they are suitable for programmatic use in your application's logic: `OPEN`, `EXIT`, `HANDOFF`, `SELECT_INSTITUTION`, `ERROR`, `BANK_INCOME_INSIGHTS_COMPLETED`, `IDENTITY_VERIFICATION_PASS_SESSION`, `IDENTITY_VERIFICATION_FAIL_SESSION`, `LAYER_READY`, `LAYER_NOT_AVAILABLE`, `LAYER_AUTOFILL_NOT_AVAILABLE`. The remaining callback events are informational and subject to change, and should be used for analytics and troubleshooting purposes only.

onEvent

**Properties**

[`eventName`](/docs/link/react-native/#link-react_native-onevent-eventName)

LinkEventNameLinkEventName

A string representing the event that has just occurred in the Link flow.

[`AUTO_SUBMIT_PHONE`](/docs/link/react-native/#link-react_native-onevent-eventName-AUTO-SUBMIT-PHONE)

The user was automatically sent an OTP code without a UI prompt. This event can only occur if the user's phone phone number was provided to Link via the `/link/token/create` call and the user has previously consented to receive OTP codes from Plaid.

[`BANK_INCOME_INSIGHTS_COMPLETED`](/docs/link/react-native/#link-react_native-onevent-eventName-BANK-INCOME-INSIGHTS-COMPLETED)

The user has completed the Assets and Bank Income Insights flow.

[`CLOSE_OAUTH`](/docs/link/react-native/#link-react_native-onevent-eventName-CLOSE-OAUTH)

The user closed the third-party website or mobile app without completing the OAuth flow.

[`CONNECT_NEW_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-eventName-CONNECT-NEW-INSTITUTION)

The user has chosen to link a new institution instead of linking a saved institution. This event is only emitted in the Link Remember Me flow.

[`ERROR`](/docs/link/react-native/#link-react_native-onevent-eventName-ERROR)

A recoverable error occurred in the Link flow, see the `error_code` metadata.

[`EXIT`](/docs/link/react-native/#link-react_native-onevent-eventName-EXIT)

The user has exited without completing the Link flow and the [onExit](#onexit) callback is fired.

[`FAIL_OAUTH`](/docs/link/react-native/#link-react_native-onevent-eventName-FAIL-OAUTH)

The user encountered an error while completing the third-party's OAuth login flow.

[`HANDOFF`](/docs/link/react-native/#link-react_native-onevent-eventName-HANDOFF)

The user has exited Link after successfully linking an Item.

[`IDENTITY_VERIFICATION_START_STEP`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-START-STEP)

The user has started a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_PASS_STEP`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-PASS-STEP)

The user has passed a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_FAIL_STEP`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-FAIL-STEP)

The user has failed a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_PENDING_REVIEW_STEP`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-PENDING-REVIEW-STEP)

The user has reached the pending review state.

[`IDENTITY_VERIFICATION_CREATE_SESSION`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-CREATE-SESSION)

The user has started a new Identity Verification session.

[`IDENTITY_VERIFICATION_RESUME_SESSION`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-RESUME-SESSION)

The user has resumed an existing Identity Verification session.

[`IDENTITY_VERIFICATION_PASS_SESSION`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-PASS-SESSION)

The user has passed their Identity Verification session.

[`IDENTITY_VERIFICATION_PENDING_REVIEW_SESSION`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-PENDING-REVIEW-SESSION)

The user has completed their Identity Verification session, which is now in a pending review state.

[`IDENTITY_VERIFICATION_FAIL_SESSION`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-FAIL-SESSION)

The user has failed their Identity Verification session.

[`IDENTITY_VERIFICATION_OPEN_UI`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-OPEN-UI)

The user has opened the UI of their Identity Verification session.

[`IDENTITY_VERIFICATION_RESUME_UI`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-RESUME-UI)

The user has resumed the UI of their Identity Verification session.

[`IDENTITY_VERIFICATION_CLOSE_UI`](/docs/link/react-native/#link-react_native-onevent-eventName-IDENTITY-VERIFICATION-CLOSE-UI)

The user has closed the UI of their Identity Verification session.

[`LAYER_AUTOFILL_NOT_AVAILABLE`](/docs/link/react-native/#link-react_native-onevent-eventName-LAYER-AUTOFILL-NOT-AVAILABLE)

The user's date of birth passed to Link is not eligible for Layer Extended Autofill.

[`LAYER_NOT_AVAILABLE`](/docs/link/react-native/#link-react_native-onevent-eventName-LAYER-NOT-AVAILABLE)

The user phone number passed to Link is not eligible for Layer.

[`LAYER_READY`](/docs/link/react-native/#link-react_native-onevent-eventName-LAYER-READY)

The user phone number passed to Link is eligible for Layer and `open()` may now be called.

[`MATCHED_SELECT_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-eventName-MATCHED-SELECT-INSTITUTION)

The user selected an institution that was presented as a matched institution. This event can be emitted if [Embedded Institution Search](https://plaid.com/docs/link/embedded-institution-search/) is being used, if the institution was surfaced as a matched institution likely to have been linked to Plaid by a returning user, or if the institution's `routing_number` was provided when calling `/link/token/create`. For details on which scenario is triggering the event, see `metadata.matchReason`.

[`OPEN`](/docs/link/react-native/#link-react_native-onevent-eventName-OPEN)

The user has opened Link.

[`OPEN_MY_PLAID`](/docs/link/react-native/#link-react_native-onevent-eventName-OPEN-MY-PLAID)

The user has opened my.plaid.com. This event is only sent when Link is initialized with Assets as a product.

[`OPEN_OAUTH`](/docs/link/react-native/#link-react_native-onevent-eventName-OPEN-OAUTH)

The user has navigated to a third-party website or mobile app in order to complete the OAuth login flow.

[`SEARCH_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-eventName-SEARCH-INSTITUTION)

The user has searched for an institution.

[`SELECT_AUTH_TYPE`](/docs/link/react-native/#link-react_native-onevent-eventName-SELECT-AUTH-TYPE)

The user has chosen whether to Link instantly or manually (i.e., with micro-deposits). This event emits the `selection` metadata to indicate the user's selection.

[`SELECT_BRAND`](/docs/link/react-native/#link-react_native-onevent-eventName-SELECT-BRAND)

The user selected a brand, e.g. Bank of America. The `SELECT_BRAND` event is only emitted for large financial institutions with multiple online banking portals.

[`SELECT_DEGRADED_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-eventName-SELECT-DEGRADED-INSTITUTION)

The user selected an institution with a `DEGRADED` health status and was shown a corresponding message.

[`SELECT_DOWN_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-eventName-SELECT-DOWN-INSTITUTION)

The user selected an institution with a `DOWN` health status and was shown a corresponding message.

[`SELECT_FILTERED_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-eventName-SELECT-FILTERED-INSTITUTION)

The user selected an institution Plaid does not support all requested products for and was shown a corresponding message.

[`SELECT_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-eventName-SELECT-INSTITUTION)

The user selected an institution.

[`SKIP_SUBMIT_PHONE`](/docs/link/react-native/#link-react_native-onevent-eventName-SKIP-SUBMIT-PHONE)

The user has opted to not provide their phone number to Plaid. This event is only emitted in the Link Remember Me flow.

[`SUBMIT_ACCOUNT_NUMBER`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-ACCOUNT-NUMBER)

The user has submitted an account number. This event emits the `account_number_mask` metadata to indicate the mask of the account number the user provided.

[`SUBMIT_CREDENTIALS`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-CREDENTIALS)

The user has submitted credentials.

[`SUBMIT_DOCUMENTS`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-DOCUMENTS)

The user is being prompted to submit documents for an Income verification flow.

[`SUBMIT_DOCUMENTS_ERROR`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-DOCUMENTS-ERROR)

The user encountered an error when submitting documents for an Income verification flow.

[`SUBMIT_DOCUMENTS_SUCCESS`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-DOCUMENTS-SUCCESS)

The user has successfully submitted documents for an Income verification flow.

[`SUBMIT_MFA`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-MFA)

The user has submitted MFA.

[`SUBMIT_OTP`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-OTP)

The user has submitted an OTP code during the phone number verification flow. This event is only emitted in the Link Returning User Experience (Remember Me) flow or the Layer flow. This event will not be emitted if the phone number is verified via SNA.

[`SUBMIT_PHONE`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-PHONE)

The user has submitted their phone number. This event is only emitted in the Link Returning User Experience (Remember Me) flow.

[`SUBMIT_ROUTING_NUMBER`](/docs/link/react-native/#link-react_native-onevent-eventName-SUBMIT-ROUTING-NUMBER)

The user has submitted a routing number. This event emits the `routing_number` metadata to indicate user's routing number.

[`TRANSITION_VIEW`](/docs/link/react-native/#link-react_native-onevent-eventName-TRANSITION-VIEW)

The `TRANSITION_VIEW` event indicates that the user has moved from one view to the next.

[`VERIFY_PHONE`](/docs/link/react-native/#link-react_native-onevent-eventName-VERIFY-PHONE)

The user has successfully verified their phone number using OTP or SNA. This event is only emitted in the Link Returning User Experience (Remember Me) flow or the Layer flow.

[`VIEW_DATA_TYPES`](/docs/link/react-native/#link-react_native-onevent-eventName-VIEW-DATA-TYPES)

The user has viewed data types on the data transparency consent pane.

[`UNKNOWN`](/docs/link/react-native/#link-react_native-onevent-eventName-UNKNOWN)

The `UNKNOWN` event indicates that the event is not handled by the current version of the SDK.

[`metadata`](/docs/link/react-native/#link-react_native-onevent-metadata)

LinkEventMetadataLinkEventMetadata

An object containing information about the event.

[`submitAccountNumber`](/docs/link/react-native/#link-react_native-onevent-metadata-submitAccountNumber)

StringString

The account number mask extracted from the user-provided account number. If the user-inputted account number is four digits long, `account_number_mask` is empty. Emitted by `SUBMIT_ACCOUNT_NUMBER`.

[`errorCode`](/docs/link/react-native/#link-react_native-onevent-metadata-errorCode)

StringString

The error code that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`errorMessage`](/docs/link/react-native/#link-react_native-onevent-metadata-errorMessage)

StringString

The error message that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`errorType`](/docs/link/react-native/#link-react_native-onevent-metadata-errorType)

StringString

The error type that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`exitStatus`](/docs/link/react-native/#link-react_native-onevent-metadata-exitStatus)

StringString

The status key indicates the point at which the user exited the Link flow. Emitted by: `EXIT`.

[`institutionId`](/docs/link/react-native/#link-react_native-onevent-metadata-institutionId)

StringString

The ID of the selected institution. Emitted by: *all events*.

[`institutionName`](/docs/link/react-native/#link-react_native-onevent-metadata-institutionName)

StringString

The name of the selected institution. Emitted by: *all events*.

[`institutionSearchQuery`](/docs/link/react-native/#link-react_native-onevent-metadata-institutionSearchQuery)

StringString

The query used to search for institutions. Emitted by: `SEARCH_INSTITUTION`.

[`isUpdateMode`](/docs/link/react-native/#link-react_native-onevent-metadata-isUpdateMode)

StringString

Indicates if the current Link session is an update mode session. Emitted by: `OPEN`.

[`matchReason`](/docs/link/react-native/#link-react_native-onevent-metadata-matchReason)

nullablestringnullable, string

The reason this institution was matched.
This will be either `returning_user` or `routing_number` if emitted by: `MATCHED_SELECT_INSTITUTION`.
Otherwise, this will be `SAVED_INSTITUTION` or `AUTO_SELECT_SAVED_INSTITUTION` if emitted by: `SELECT_INSTITUTION`.

[`routingNumber`](/docs/link/react-native/#link-react_native-onevent-metadata-routingNumber)

Optional<String>Optional<String>

The routing number submitted by user at the micro-deposits routing number pane. Emitted by `SUBMIT_ROUTING_NUMBER`.

[`linkSessionId`](/docs/link/react-native/#link-react_native-onevent-metadata-linkSessionId)

StringString

The `linkSessionId` is a unique identifier for a single session of Link. It's always available and will stay constant throughout the flow. Emitted by: *all events*.

[`mfaType`](/docs/link/react-native/#link-react_native-onevent-metadata-mfaType)

StringString

If set, the user has encountered one of the following MFA types: `code` `device` `questions` `selections`. Emitted by: `SUBMIT_MFA` and `TRANSITION_VIEW` when `view_name` is `MFA`.

[`requestId`](/docs/link/react-native/#link-react_native-onevent-metadata-requestId)

StringString

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation. Emitted by: *all events*.

[`selection`](/docs/link/react-native/#link-react_native-onevent-metadata-selection)

StringString

The Auth Type Select flow type selected by the user. Possible values are `flow_type_manual` or `flow_type_instant`. Emitted by: `SELECT_AUTH_TYPE`.

[`timestamp`](/docs/link/react-native/#link-react_native-onevent-metadata-timestamp)

StringString

An ISO 8601 representation of when the event occurred. For example, `2017-09-14T14:42:19.350Z`. Emitted by: *all events*.

[`viewName`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName)

LinkEventViewNameLinkEventViewName

The name of the view that is being transitioned to. Emitted by: `TRANSITION_VIEW`.

[`ACCEPT_TOS`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-ACCEPT-TOS)

The view showing Terms of Service in the identity verification flow.

[`CONNECTED`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-CONNECTED)

The user has connected their account.

[`CONSENT`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-CONSENT)

We ask the user to consent to the privacy policy.

[`CREDENTIAL`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-CREDENTIAL)

Asking the user for their account credentials.

[`DOCUMENTARY_VERIFICATION`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-DOCUMENTARY-VERIFICATION)

The view requesting document verification in the identity verification flow (configured via "Fallback Settings" in the "Rulesets" section of the template editor).

[`ERROR`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-ERROR)

An error has occurred.

[`EXIT`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-EXIT)

Confirming if the user wishes to close Link.

[`INSTANT_MICRODEPOSIT_AUTHORIZED`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-INSTANT-MICRODEPOSIT-AUTHORIZED)

The user has authorized an instant micro-deposit to be sent to their account over the RTP or FedNow network with a 3-letter code to verify their account.

[`KYC_CHECK`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-KYC-CHECK)

The view representing the "know your customer" step in the identity verification flow.

[`LOADING`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-LOADING)

Link is making a request to our servers.

[`MFA`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-MFA)

The user is asked by the institution for additional MFA authentication.

[`NUMBERS`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-NUMBERS)

The user is asked to insert their account and routing numbers.

[`NUMBERS_SELECT_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-NUMBERS-SELECT-INSTITUTION)

The user goes through the Same Day micro-deposits flow with Reroute to Credentials.

[`OAUTH`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-OAUTH)

The user is informed they will authenticate with the financial institution via OAuth.

[`PROFILE_DATA_REVIEW`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-PROFILE-DATA-REVIEW)

The user is asked to review their profile data in the Layer flow.

[`RECAPTCHA`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-RECAPTCHA)

The user was presented with a Google reCAPTCHA to verify they are human.

[`RISK_CHECK`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-RISK-CHECK)

The risk check step in the identity verification flow (configured via "Risk Rules" in the "Rulesets" section of the template editor).

[`SAME_DAY_MICRODEPOSIT_AUTHORIZED`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SAME-DAY-MICRODEPOSIT-AUTHORIZED)

The user has authorized a same day micro-deposit to be sent to their account over the ACH network with a 3-letter code to verify their account.

[`SCREENING`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SCREENING)

The watchlist screening step in the identity verification flow.

[`SELECT_ACCOUNT`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SELECT-ACCOUNT)

We ask the user to choose an account.

[`SELECT_AUTH_TYPE`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SELECT-AUTH-TYPE)

The user is asked to choose whether to Link instantly or manually (i.e., with micro-deposits).

[`SELECT_BRAND`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SELECT-BRAND)

The user is asked to select a brand, e.g. Bank of America. The brand selection interface occurs before the institution select pane and is only provided for large financial institutions with multiple online banking portals.

[`SELECT_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SELECT-INSTITUTION)

We ask the user to choose their institution.

[`SELECT_SAVED_ACCOUNT`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SELECT-SAVED-ACCOUNT)

The user is asked to select their saved accounts and/or new accounts for linking in the Link Returning User Experience (Remember Me) flow.

[`SELECT_SAVED_INSTITUTION`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SELECT-SAVED-INSTITUTION)

The user is asked to pick a saved institution or link a new one in the Link Returning User Experience (Remember Me) flow.

[`SELFIE_CHECK`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SELFIE-CHECK)

The view in the identity verification flow which uses the camera to confirm there is a real user present that matches their ID documents.

[`SUBMIT_PHONE`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-SUBMIT-PHONE)

The user is asked for their phone number in the Link Returning User Experience (Remember Me) flow.

[`UPLOAD_DOCUMENTS`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-UPLOAD-DOCUMENTS)

The user is asked to upload documents (for Income verification).

[`VERIFY_PHONE`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-VERIFY-PHONE)

The user is asked to verify their phone in the Link Returning User Experience (Remember Me) flow or the Layer flow. This screen will appear even if a non-OTP verification method is used.

[`VERIFY_SMS`](/docs/link/react-native/#link-react_native-onevent-metadata-viewName-VERIFY-SMS)

The SMS verification step in the identity verification flow.

```
usePlaidEmitter((event) => {
  console.log(event);
});
```

=\*=\*=\*=

#### submit()

The `submit` function is currently only used in the Layer product. It allows the client application to submit additional user-collected data to the Link flow (e.g. a user phone number).

submit

**Properties**

[`submissionData`](/docs/link/react-native/#link-react_native-submit-submissionData)

objectobject

Data to submit during a Link session.

[`phoneNumber`](/docs/link/react-native/#link-react_native-submit-submissionData-phoneNumber)

StringString

The end user's phone number.

[`dateOfBirth`](/docs/link/react-native/#link-react_native-submit-submissionData-dateOfBirth)

StringString

The end user's date of birth. To be provided in the format "yyyy-mm-dd".

```
submit({
  "phone_number": "+14155550123"
})
```

=\*=\*=\*=

#### destroy()

The `destroy` function clears state and resources from a previously opened session. The `destroy` function is only available on Android, as this state clearing behavior occurs automatically on iOS. `destroy` is intended for use with the Layer product and should be used if you are making multiple calls to `create` before calling `submit`. By calling `destroy` between the `create` calls, you can avoid unexpected behavior on the `submit` call.

destroy

**Properties**

This endpoint or method takes an empty request body.

```
create(tokenConfiguration1);
(async () => {
  try {
    await destroy(); // Clear previous session state
    create(tokenConfiguration2);
    submit(phoneNumber);
  } catch (e) {
    console.error('Error during flow:', e);
  }
})();
```

#### OAuth

Using Plaid Link with an OAuth flow requires some additional setup instructions. For details, see the [OAuth guide](/docs/link/oauth/).

#### Upgrading

The latest version of the SDK is available from [GitHub](https://github.com/plaid/react-native-plaid-link-sdk). New versions of the SDK are released frequently. Major releases occur annually. The Link SDK uses Semantic Versioning, ensuring that all non-major releases are non-breaking, backwards-compatible updates. We recommend you update regularly (at least once a quarter, and ideally once a month) to ensure the best Plaid Link experience in your application.

SDK versions are supported for two years; with each major SDK release, Plaid will stop officially supporting any previous SDK versions that are more than two years old. While these older versions are expected to continue to work without disruption, Plaid will not provide assistance with unsupported SDK versions.

##### Migration guides

- Version 12.x removes the `PlaidLink` component and `openLink` function, which were deprecated in version 11.6.0. If you are using this method of opening Link, replace it with the new process that uses `create` and then `open`. For details, see [Opening Link](/docs/link/react-native/#opening-link) or the [README](https://github.com/plaid/react-native-plaid-link-sdk?tab=readme-ov-file#version--1160). Version 12.x also removes the `PROFILE_ELIGIBILITY_CHECK_ERROR`.
- Version 12.x updates the Android target SDK version and compile version from 33 to 35 and requires use of an Xcode 16 toolchain.
- Version 11.x contains several breaking changes from previous major versions. For details, see the [migration guide on GitHub](https://github.com/plaid/react-native-plaid-link-sdk/blob/master/v11-migration-guide.md).
- When upgrading from 9.x to 10.x or later, you should remove any invocation of `useDeepLinkRedirector` on iOS, as it has been removed from the SDK, since it is no longer required for handling Universal Links. You must make sure you are using a compatible version of React Native (0.66.0 or higher) and the Plaid iOS SDK (see the [version compatibility chart](https://github.com/plaid/react-native-plaid-link-sdk#version-compatibility) on GitHub).
- No code changes are required to upgrade from 8.x to 9.x, although you must make sure you are using a compatible version of the Plaid iOS SDK (4.1.0 or higher).
- The only difference between version 7.x and 8.x is that 8.x adds support for Xcode 14. No code changes are required to upgrade from 7.x to 8.x, although you must convert to an Xcode 14 toolchain.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
