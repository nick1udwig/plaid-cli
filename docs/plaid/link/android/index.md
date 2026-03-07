---
title: "Link - Android | Plaid Docs"
source_url: "https://plaid.com/docs/link/android/"
scraped_at: "2026-03-07T22:05:03+00:00"
---

# Link Android SDK

#### Learn how to integrate your app with the Link Android SDK

#### Overview

The Plaid Link SDK is a quick and secure way to link bank accounts to Plaid in your Android app. Link
is a drop-in module that handles connecting a financial institution to your app (credential validation,
multi-factor authentication, error handling, etc.), without passing sensitive personal information to your server.

To get started with Plaid Link for Android, clone the
[GitHub repository](https://github.com/plaid/plaid-link-android) and try out the example application,
which provides a reference implementation in both Java and Kotlin. Youʼll want to sign up for
[free API keys](https://dashboard.plaid.com/developers/keys) through the Plaid Dashboard to get started.

Prefer to learn by watching? A [video guide](https://youtu.be/oM7vL49I5tc) is available for this content.

#### Initial Android setup

Before writing code using the SDK, you must first perform some setup steps to register your app with Plaid and configure your project.

##### Register your app ID

To register your Android app ID:

1. Sign in to the [Plaid Dashboard](https://dashboard.plaid.com/signin)
   and go to the [**Developers -> API**](https://dashboard.plaid.com/developers/api) page.
2. Next to **Allowed Android Package Names** click **Configure** then **Add New Android Package Name**.
3. Enter your package name, for example `com.plaid.example`.
4. Click **Save Changes**.

Your Android app is now set up and ready to start integrating with the Plaid SDK.

New versions of the Android SDK are released frequently, at least once every few months. Major releases occur annually.
You should keep your version up-to-date to provide the best Plaid Link experience in your application.

##### Update your project plugins

In your root-level (project-level) Gradle file (`build.gradle`), add rules to include
the Android Gradle plugin. Check that you have Google's Maven repository as well.

build.gradle (Project-level)

```
buildscript {
    repositories {
        // Check that you have the following line (if not, add it):
        google()  // Google's Maven repository
        mavenCentral() // Include to import Plaid Link Android SDK
    }
    dependencies {
        // ...
    }
}
```

##### Add the PlaidLink SDK to your app

In your module (app-level) Gradle file (usually `app/build.gradle`), add a line to the bottom of the file.
The latest version of the PlaidLink SDK is ![Maven Central](https://img.shields.io/maven-central/v/com.plaid.link/sdk-core)
and can be found on [Maven Central](https://search.maven.org/artifact/com.plaid.link/sdk-core).

build.gradle (App-level)

```
android {
  defaultConfig {
    minSdkVersion 21 // or greater
  }
}

dependencies {
  // ...
  implementation 'com.plaid.link:sdk-core:<insert latest version>'
}
```

##### Enable camera support (Identity Verification only)

If your app uses [Identity Verification](/docs/identity-verification/), a user may need to take a picture of identity documentation or a selfie during the Link flow. To support this workflow, the [`CAMERA`](https://developer.android.com/reference/android/Manifest.permission#CAMERA) , [`WRITE_EXTERNAL_STORAGE`](https://developer.android.com/reference/android/Manifest.permission#WRITE_EXTERNAL_STORAGE), [`RECORD_AUDIO`](https://developer.android.com/reference/android/Manifest.permission#RECORD_AUDIO), and [`MODIFY_AUDIO_SETTINGS`](https://developer.android.com/reference/android/Manifest.permission#MODIFY_AUDIO_SETTINGS) permissions need to be added to your application's `AndroidManifest.xml`. (While Plaid does not record any audio, some older Android devices require these last two permissions to use the camera.) The `WRITE_EXTERNAL_STORAGE` permission should be limited to < Android 9 (i.e. maxSdk=28). If these permissions are not granted in an app that uses Identity Verification, the app may crash during Link.

#### Opening Link

Before you can open Link, you need to first create a `link_token` by calling [`/link/token/create`](/docs/api/link/#linktokencreate) from your backend.
This call should **never** happen directly from the mobile client, as it risks exposing your API secret.
The [`/link/token/create`](/docs/api/link/#linktokencreate) call must include the `android_package_name` parameter, which should match the `applicationId` from your app-level
`build.gradle` file. You can learn more about `applicationId` in Google's [Android
developer documentation](https://developer.android.com/studio/build/application-id).

/link/token/create

```
const request: LinkTokenCreateRequest = {
  user: {
    client_user_id: 'user-id',
  },
  client_name: 'Plaid Test App',
  products: ['auth', 'transactions'],
  country_codes: ['GB'],
  language: 'en',
  webhook: 'https://sample-web-hook.com',
  android_package_name: 'com.plaid.example'
};
try {
  const response = await plaidClient.linkTokenCreate(request);
  const linkToken = response.data.link_token;
} catch (error) {
  // handle error
}
```

##### Create a LinkTokenConfiguration

Each time you open Link, you will need to get a new `link_token` from your server and create a new
`LinkTokenConfiguration` object with it.

openLink

```
val linkTokenConfiguration = linkTokenConfiguration {
  token = "LINK_TOKEN_FROM_SERVER"
}
```

The Link SDK runs as a separate `Activity` within your app. In order to return the result
to your app, it supports both the standard `startActivityForResult` and `onActivityResult`
and the `ActivityResultContract` [result APIs](https://developer.android.com/training/basics/intents/result).

Select group for content switcher

##### Register a callback for an Activity Result

```
private val linkAccountToPlaid =
registerForActivityResult(FastOpenPlaidLink()) {
  when (it) {
    is LinkSuccess -> /* handle LinkSuccess */
    is LinkExit -> /* handle LinkExit */
  }
}
```

##### Create a PlaidHandler

Create a `PlaidHandler` - A `PlaidHandler` is a one-time use object used to open a Link session.
It should be created as early as possible to warm up Link so that it opens quickly. We recommend
doing this as early as possible, since it must be completed before Link opens, and if you create
it just before opening Link, it can have a perceptible impact on Link startup time.

```
val plaidHandler: PlaidHandler = 
  Plaid.create(application, linkTokenConfiguration)
```

##### Open Link

```
linkAccountToPlaid.launch(plaidHandler)
```

At this point, Link will open, and will trigger the `onSuccess` callback if the user successfully completes the Link flow.

=\*=\*=\*=

#### onSuccess

The method is called when a user successfully links an Item. The onSuccess handler returns a `LinkConnection` class that includes the `public_token`, and additional Link metadata in the form of a `LinkConnectionMetadata` class.

onSuccess

**Properties**

[`publicToken`](/docs/link/android/#link-android-onsuccess-publicToken)

StringString

Displayed once a user has successfully completed Link. If using Identity Verification or Beacon, this field will be `null`. If using Document Income or Payroll Income, the `public_token` will be returned, but is not used.

[`metadata`](/docs/link/android/#link-android-onsuccess-metadata)

ObjectObject

Displayed once a user has successfully completed Link.

[`accounts`](/docs/link/android/#link-android-onsuccess-metadata-accounts)

List<LinkAccount>List<LinkAccount>

A list of accounts attached to the connected Item. If Account Select is enabled via the developer dashboard, `accounts` will only include selected accounts.

[`id`](/docs/link/android/#link-android-onsuccess-metadata-accounts-id)

stringstring

The Plaid `account_id`

[`name`](/docs/link/android/#link-android-onsuccess-metadata-accounts-name)

stringstring

The official account name

[`mask`](/docs/link/android/#link-android-onsuccess-metadata-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, it may also not match the mask that the bank displays to the user.

[`subtype`](/docs/link/android/#link-android-onsuccess-metadata-accounts-subtype)

LinkAccountSubtypeLinkAccountSubtype

The account subtype. See the [Account schema](/docs/api/accounts#account-type-schema) for a full list of possible values

[`type`](/docs/link/android/#link-android-onsuccess-metadata-accounts-subtype-type)

LinkAccountTypeLinkAccountType

The account type. See the [Account schema](/docs/api/accounts#account-type-schema) for a full list of possible values

[`verification_status`](/docs/link/android/#link-android-onsuccess-metadata-accounts-verification-status)

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

[`institution`](/docs/link/android/#link-android-onsuccess-metadata-institution)

nullableobjectnullable, object

An institution object. If the Item was created via Same-Day micro-deposit verification, will be `null`.

[`name`](/docs/link/android/#link-android-onsuccess-metadata-institution-name)

stringstring

The full institution name, such as `'Wells Fargo'`

[`institution_id`](/docs/link/android/#link-android-onsuccess-metadata-institution-institution-id)

stringstring

The Plaid institution identifier

[`linkSessionId`](/docs/link/android/#link-android-onsuccess-metadata-linkSessionId)

nullableStringnullable, String

A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.

[`metadataJson`](/docs/link/android/#link-android-onsuccess-metadata-metadataJson)

nullableMapnullable, Map

The data directly returned from the server with no client side changes.

```
val success = result as LinkSuccess

// Send public_token to your server, exchange for access_token 
// (if using Item-based products)
val publicToken = success.publicToken
val metadata = success.metadata
metadata.accounts.forEach { account ->
  val accountId = account.id
  val accountName = account.name
  val accountMask = account.mask
  val accountSubType = account.subtype
}
val institutionId = metadata.institution?.id
val institutionName = metadata.institution?.name
```

=\*=\*=\*=

#### onExit

The `onExit` handler is called when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. The `PlaidError` returned from the `onExit` handler is meant to help you guide your users after they have exited Link. We recommend storing the error and metadata information server-side in a way that can be associated with the user. You’ll also need to include this and any other relevant information in Plaid Support requests for the user.

onExit

**Properties**

[`error`](/docs/link/android/#link-android-onexit-error)

Map<String, Object>Map<String, Object>

An object that contains the error type, error code, and error message of the error that was last encountered by the user. If no error was encountered, `error` will be `null`.

[`displayMessage`](/docs/link/android/#link-android-onexit-error-displayMessage)

nullableStringnullable, String

A user-friendly representation of the error code. `null` if the error is not related to user action. This may change over time and is not safe for programmatic use.

[`errorCode`](/docs/link/android/#link-android-onexit-error-errorCode)

StringString

The particular error code. Each `errorType` has a specific set of `errorCodes`. A code of 499 indicates a client-side exception.

[`json`](/docs/link/android/#link-android-onexit-error-errorCode-json)

StringString

A string representation of the error code.

[`errorType`](/docs/link/android/#link-android-onexit-error-errorCode-errorType)

StringString

A broad categorization of the error.

[`errorMessage`](/docs/link/android/#link-android-onexit-error-errorMessage)

StringString

A developer-friendly representation of the error code.

[`errorJson`](/docs/link/android/#link-android-onexit-error-errorJson)

nullableStringnullable, String

The data directly returned from the server with no client side changes.

[`LinkExitMetadata`](/docs/link/android/#link-android-onexit-LinkExitMetadata)

Map<String, Object>Map<String, Object>

An object containing information about the exit event

[`linkSessionId`](/docs/link/android/#link-android-onexit-LinkExitMetadata-linkSessionId)

nullableStringnullable, String

A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.

[`institution`](/docs/link/android/#link-android-onexit-LinkExitMetadata-institution)

nullableobjectnullable, object

An institution object. If the Item was created via Same-Day micro-deposit verification, will be `null`.

[`name`](/docs/link/android/#link-android-onexit-LinkExitMetadata-institution-name)

stringstring

The full institution name, such as `'Wells Fargo'`

[`institution_id`](/docs/link/android/#link-android-onexit-LinkExitMetadata-institution-institution-id)

stringstring

The Plaid institution identifier

[`status`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status)

nullableStringnullable, String

The point at which the user exited the Link flow. One of the following values.

[`requires_questions`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-requires-questions)

User prompted to answer security questions

[`requires_selections`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-requires-selections)

User prompted to answer multiple choice question(s)

[`requires_recaptcha`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-requires-recaptcha)

User prompted to solve a reCAPTCHA challenge

[`requires_code`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-requires-code)

User prompted to provide a one-time passcode

[`choose_device`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-choose-device)

User prompted to select a device on which to receive a one-time passcode

[`requires_credentials`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-requires-credentials)

User prompted to provide credentials for the selected financial institution or has not yet selected a financial institution

[`requires_account_selection`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-requires-account-selection)

User prompted to select one or more financial accounts to share

[`institution_not_found`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-institution-not-found)

User exited the Link flow on the institution selection pane. Typically this occurs after the user unsuccessfully (no results returned) searched for a financial institution. Note that this status does not necessarily indicate that the user was unable to find their institution, as it is used for all user exits that occur from the institution selection pane, regardless of other user behavior.

[`institution_not_supported`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-institution-not-supported)

User exited the Link flow after discovering their selected institution is no longer supported by Plaid

[`unknown`](/docs/link/android/#link-android-onexit-LinkExitMetadata-status-unknown)

An exit status that is not handled by the current version of the SDK

[`requestId`](/docs/link/android/#link-android-onexit-LinkExitMetadata-requestId)

nullableStringnullable, String

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation

```
val exit = result as LinkExit

val error = exit.error
error?.let { err ->
  val errorCode = err.errorCode
  val errorMessage = err.errorMessage
  val displayMessage = err.displayMessage
}
val metadata = exit.metadata
val institutionId = metadata.institution?.id
val institutionName = metadata.institution?.name
val linkSessionId = metadata.linkSessionId;
val requestId = metadata.requestId;
```

=\*=\*=\*=

#### onEvent

The `onEvent` callback is called at certain points in the Link flow. Unlike the handlers for `onSuccess` and `onExit`, the `onEvent` handler is initialized as a global lambda passed to the `Plaid` class. `OPEN`, `LAYER_READY`, `LAYER_NOT_AVAILABLE`, and `LAYER_AUTOFILL_NOT_AVAILABLE` events will be sent in real-time, and remaining events will be sent when the Link session is finished and `onSuccess` or `onExit` is called. Callback ordering is not guaranteed; `onEvent` callbacks may fire before, after, or surrounding the `onSuccess` or `onExit` callback, and event callbacks are not guaranteed to fire in the order in which they occurred. If you need the exact time when an event happened, use the `timestamp` property.

The following `onEvent` callbacks are stable, which means that they are suitable for programmatic use in your application's logic: `OPEN`, `EXIT`, `HANDOFF`, `SELECT_INSTITUTION`, `ERROR`, `BANK_INCOME_INSIGHTS_COMPLETED`, `IDENTITY_VERIFICATION_PASS_SESSION`, `IDENTITY_VERIFICATION_FAIL_SESSION`, `LAYER_READY`, `LAYER_NOT_AVAILABLE`, `LAYER_AUTOFILL_NOT_AVAILABLE`. The remaining callback events are informational and subject to change, and should be used for analytics and troubleshooting purposes only.

onEvent

**Properties**

[`eventName`](/docs/link/android/#link-android-onevent-eventName)

StringString

A string representing the event that has just occurred in the Link flow.

[`AUTO_SUBMIT_PHONE`](/docs/link/android/#link-android-onevent-eventName-AUTO-SUBMIT-PHONE)

The user was automatically sent an OTP code without a UI prompt. This event can only occur if the user's phone phone number was provided to Link via the `/link/token/create` call and the user has previously consented to receive OTP codes from Plaid.

[`BANK_INCOME_INSIGHTS_COMPLETED`](/docs/link/android/#link-android-onevent-eventName-BANK-INCOME-INSIGHTS-COMPLETED)

The user has completed the Assets and Bank Income Insights flow.

[`CLOSE_OAUTH`](/docs/link/android/#link-android-onevent-eventName-CLOSE-OAUTH)

The user closed the third-party website or mobile app without completing the OAuth flow.

[`CONNECT_NEW_INSTITUTION`](/docs/link/android/#link-android-onevent-eventName-CONNECT-NEW-INSTITUTION)

The user has chosen to link a new institution instead of linking a saved institution. This event is only emitted in the Link Remember Me flow.

[`ERROR`](/docs/link/android/#link-android-onevent-eventName-ERROR)

A recoverable error occurred in the Link flow, see the `error_code` metadata.

[`EXIT`](/docs/link/android/#link-android-onevent-eventName-EXIT)

The user has exited without completing the Link flow and the [onExit](#onexit) callback is fired.

[`FAIL_OAUTH`](/docs/link/android/#link-android-onevent-eventName-FAIL-OAUTH)

The user encountered an error while completing the third-party's OAuth login flow.

[`HANDOFF`](/docs/link/android/#link-android-onevent-eventName-HANDOFF)

The user has exited Link after successfully linking an Item.

[`IDENTITY_MATCH_FAILED`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-MATCH-FAILED)

An Identity Match check configured via the Account Verification Dashboard failed the Identity Match rules and did not detect a match.

[`IDENTITY_MATCH_PASSED`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-MATCH-PASSED)

An Identity Match check configured via the Account Verification Dashboard passed the Identity Match rules and detected a match.

[`IDENTITY_VERIFICATION_START_STEP`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-START-STEP)

The user has started a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_PASS_STEP`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-PASS-STEP)

The user has passed a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_FAIL_STEP`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-FAIL-STEP)

The user has failed a step of the Identity Verification flow. The step is indicated by `view_name`.

[`IDENTITY_VERIFICATION_PENDING_REVIEW_STEP`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-PENDING-REVIEW-STEP)

The user has reached the pending review state.

[`IDENTITY_VERIFICATION_CREATE_SESSION`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-CREATE-SESSION)

The user has started a new Identity Verification session.

[`IDENTITY_VERIFICATION_RESUME_SESSION`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-RESUME-SESSION)

The user has resumed an existing Identity Verification session.

[`IDENTITY_VERIFICATION_PASS_SESSION`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-PASS-SESSION)

The user has passed their Identity Verification session.

[`IDENTITY_VERIFICATION_FAIL_SESSION`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-FAIL-SESSION)

The user has failed their Identity Verification session.

[`IDENTITY_VERIFICATION_PENDING_REVIEW_SESSION`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-PENDING-REVIEW-SESSION)

The user has completed their Identity Verification session, which is now in a pending review state.

[`IDENTITY_VERIFICATION_OPEN_UI`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-OPEN-UI)

The user has opened the UI of their Identity Verification session.

[`IDENTITY_VERIFICATION_RESUME_UI`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-RESUME-UI)

The user has resumed the UI of their Identity Verification session.

[`IDENTITY_VERIFICATION_CLOSE_UI`](/docs/link/android/#link-android-onevent-eventName-IDENTITY-VERIFICATION-CLOSE-UI)

The user has closed the UI of their Identity Verification session.

[`LAYER_AUTOFILL_NOT_AVAILABLE`](/docs/link/android/#link-android-onevent-eventName-LAYER-AUTOFILL-NOT-AVAILABLE)

The user's date of birth passed to Link is not eligible for Layer Extended Autofill.

[`LAYER_NOT_AVAILABLE`](/docs/link/android/#link-android-onevent-eventName-LAYER-NOT-AVAILABLE)

The user phone number passed to Link is not eligible for Layer.

[`LAYER_READY`](/docs/link/android/#link-android-onevent-eventName-LAYER-READY)

The user phone number passed to Link is eligible for Layer and `open()` may now be called.

[`MATCHED_SELECT_INSTITUTION`](/docs/link/android/#link-android-onevent-eventName-MATCHED-SELECT-INSTITUTION)

The user selected an institution that was presented as a matched institution. This event can be emitted either during the legacy version of the Returning User Experience flow or if the institution's `routing_number` was provided when calling `/link/token/create`. To distinguish between the two scenarios, see `LinkEventMetadata.match_reason`.

[`OPEN`](/docs/link/android/#link-android-onevent-eventName-OPEN)

The user has opened Link.

[`OPEN_MY_PLAID`](/docs/link/android/#link-android-onevent-eventName-OPEN-MY-PLAID)

The user has opened my.plaid.com. This event is only emitted when Link is initialized with Assets as a product.

[`OPEN_OAUTH`](/docs/link/android/#link-android-onevent-eventName-OPEN-OAUTH)

The user has navigated to a third-party website or mobile app in order to complete the OAuth login flow.

[`SEARCH_INSTITUTION`](/docs/link/android/#link-android-onevent-eventName-SEARCH-INSTITUTION)

The user has searched for an institution.

[`SKIP_SUBMIT_PHONE`](/docs/link/android/#link-android-onevent-eventName-SKIP-SUBMIT-PHONE)

The user has opted to not provide their phone number to Plaid. This event is only emitted in the Link Remember Me flow.

[`SELECT_BRAND`](/docs/link/android/#link-android-onevent-eventName-SELECT-BRAND)

The user selected a brand, e.g. Bank of America. The `SELECT_BRAND` event is only emitted for large financial institutions with multiple online banking portals.

[`SELECT_DEGRADED_INSTITUTION`](/docs/link/android/#link-android-onevent-eventName-SELECT-DEGRADED-INSTITUTION)

The user selected an institution with a `DEGRADED` health status and was shown a corresponding message.

[`SELECT_DOWN_INSTITUTION`](/docs/link/android/#link-android-onevent-eventName-SELECT-DOWN-INSTITUTION)

The user selected an institution with a `DOWN` health status and was shown a corresponding message.

[`SELECT_FILTERED_INSTITUTION`](/docs/link/android/#link-android-onevent-eventName-SELECT-FILTERED-INSTITUTION)

The user selected an institution Plaid does not support all requested products for and was shown a corresponding message.

[`SELECT_INSTITUTION`](/docs/link/android/#link-android-onevent-eventName-SELECT-INSTITUTION)

The user selected an institution.

[`SUBMIT_ACCOUNT_NUMBER`](/docs/link/android/#link-android-onevent-eventName-SUBMIT-ACCOUNT-NUMBER)

The user has submitted an account number. This event emits the `account_number_mask` metadata to indicate the mask of the account number the user provided.

[`SUBMIT_CREDENTIALS`](/docs/link/android/#link-android-onevent-eventName-SUBMIT-CREDENTIALS)

The user has submitted credentials.

[`SUBMIT_MFA`](/docs/link/android/#link-android-onevent-eventName-SUBMIT-MFA)

The user has submitted MFA.

[`SUBMIT_OTP`](/docs/link/android/#link-android-onevent-eventName-SUBMIT-OTP)

The user has submitted an OTP code during the phone number verification flow. This event is only emitted in the Link Returning User Experience (Remember Me) or Layer flow. This event will not be emitted if the phone number is verified via SNA.

[`SUBMIT_PHONE`](/docs/link/android/#link-android-onevent-eventName-SUBMIT-PHONE)

The user has submitted their phone number. This event is only emitted in the Link Remember Me flow.

[`SUBMIT_ROUTING_NUMBER`](/docs/link/android/#link-android-onevent-eventName-SUBMIT-ROUTING-NUMBER)

The user has submitted a routing number. This event emits the `routing_number` metadata to indicate user's routing number.

[`TRANSITION_VIEW`](/docs/link/android/#link-android-onevent-eventName-TRANSITION-VIEW)

The `TRANSITION_VIEW` event indicates that the user has moved from one view to the next.

[`UPLOAD_DOCUMENTS`](/docs/link/android/#link-android-onevent-eventName-UPLOAD-DOCUMENTS)

The user is asked to upload documents (for Income verification).

[`VERIFY_PHONE`](/docs/link/android/#link-android-onevent-eventName-VERIFY-PHONE)

The user has successfully verified their phone number using OTP or SNA. This event is only emitted in the Link Returning User Experience (Remember Me) flow or the Layer flow.

[`VIEW_DATA_TYPES`](/docs/link/android/#link-android-onevent-eventName-VIEW-DATA-TYPES)

The user has viewed data types on the data transparency consent pane.

[`UNKNOWN`](/docs/link/android/#link-android-onevent-eventName-UNKNOWN)

The `UNKNOWN` event indicates that the event is not handled by the current version of the SDK.

[`LinkEventMetadata`](/docs/link/android/#link-android-onevent-LinkEventMetadata)

Map<String, Object>Map<String, Object>

An object containing information about the event.

[`accountNumberMask`](/docs/link/android/#link-android-onevent-LinkEventMetadata-accountNumberMask)

StringString

The account number mask extracted from the user-provided account number. If the user-inputted account number is four digits long, `account_number_mask` is empty. Emitted by `SUBMIT_ACCOUNT_NUMBER`.

[`errorCode`](/docs/link/android/#link-android-onevent-LinkEventMetadata-errorCode)

StringString

The error code that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`errorMessage`](/docs/link/android/#link-android-onevent-LinkEventMetadata-errorMessage)

StringString

The error message that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`errorType`](/docs/link/android/#link-android-onevent-LinkEventMetadata-errorType)

StringString

The error type that the user encountered. Emitted by: `ERROR`, `EXIT`.

[`exitStatus`](/docs/link/android/#link-android-onevent-LinkEventMetadata-exitStatus)

StringString

The status key indicates the point at which the user exited the Link flow. Emitted by: `EXIT`.

[`institutionId`](/docs/link/android/#link-android-onevent-LinkEventMetadata-institutionId)

StringString

The ID of the selected institution. Emitted by: *all events*.

[`institutionName`](/docs/link/android/#link-android-onevent-LinkEventMetadata-institutionName)

StringString

The name of the selected institution. Emitted by: *all events*.

[`institutionSearchQuery`](/docs/link/android/#link-android-onevent-LinkEventMetadata-institutionSearchQuery)

StringString

The query used to search for institutions. Emitted by: `SEARCH_INSTITUTION`.

[`isUpdateMode`](/docs/link/android/#link-android-onevent-LinkEventMetadata-isUpdateMode)

StringString

Indicates if the current Link session is an update mode session. Emitted by: `OPEN`.

[`matchReason`](/docs/link/android/#link-android-onevent-LinkEventMetadata-matchReason)

nullablestringnullable, string

The reason this institution was matched.
This will be either `returning_user` or `routing_number` if emitted by: `MATCHED_SELECT_INSTITUTION`.
Otherwise, this will be `SAVED_INSTITUTION` or `AUTO_SELECT_SAVED_INSTITUTION` if emitted by: `SELECT_INSTITUTION`.

[`routingNumber`](/docs/link/android/#link-android-onevent-LinkEventMetadata-routingNumber)

Optional<String>Optional<String>

The routing number submitted by user at the micro-deposits routing number pane. Emitted by `SUBMIT_ROUTING_NUMBER`.

[`linkSessionId`](/docs/link/android/#link-android-onevent-LinkEventMetadata-linkSessionId)

StringString

The `link_session_id` is a unique identifier for a single session of Link. It's always available and will stay constant throughout the flow. Emitted by: *all events*.

[`mfaType`](/docs/link/android/#link-android-onevent-LinkEventMetadata-mfaType)

StringString

If set, the user has encountered one of the following MFA types: `code` `device` `questions` `selections`. Emitted by: `SUBMIT_MFA` and `TRANSITION_VIEW` when `view_name` is `MFA`.

[`requestId`](/docs/link/android/#link-android-onevent-LinkEventMetadata-requestId)

StringString

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation. Emitted by: *all events*.

[`selection`](/docs/link/android/#link-android-onevent-LinkEventMetadata-selection)

StringString

The Auth Type Select flow type selected by the user. Possible values are `flow_type_manual` or `flow_type_instant`. Emitted by: `SELECT_AUTH_TYPE`.

[`timestamp`](/docs/link/android/#link-android-onevent-LinkEventMetadata-timestamp)

StringString

An ISO 8601 representation of when the event occurred. For example `2017-09-14T14:42:19.350Z`. Emitted by: *all events*.

[`viewName`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName)

StringString

The name of the view that is being transitioned to. Emitted by: `TRANSITION_VIEW`.

[`ACCEPT_TOS`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-ACCEPT-TOS)

The view showing Terms of Service in the identity verification flow.

[`CONNECTED`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-CONNECTED)

The user has connected their account.

[`CONSENT`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-CONSENT)

We ask the user to consent to the privacy policy.

[`CREDENTIAL`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-CREDENTIAL)

Asking the user for their account credentials.

[`DOCUMENTARY_VERIFICATION`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-DOCUMENTARY-VERIFICATION)

The view requesting document verification in the identity verification flow (configured via "Fallback Settings" in the "Rulesets" section of the template editor).

[`ERROR`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-ERROR)

An error has occurred.

[`EXIT`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-EXIT)

Confirming if the user wishes to close Link.

[`INSTANT_MICRODEPOSIT_AUTHORIZED`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-INSTANT-MICRODEPOSIT-AUTHORIZED)

The user has authorized an instant micro-deposit to be sent to their account over the RTP or FedNow network with a 3-letter code to verify their account.

[`KYC_CHECK`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-KYC-CHECK)

The view representing the "know your customer" step in the identity verification flow.

[`LOADING`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-LOADING)

Link is making a request to our servers.

[`MFA`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-MFA)

The user is asked by the institution for additional MFA authentication.

[`NUMBERS`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-NUMBERS)

The user is asked to insert their account and routing numbers.

[`NUMBERS_SELECT_INSTITUTION`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-NUMBERS-SELECT-INSTITUTION)

The user goes through the Same Day micro-deposits flow with Reroute to Credentials.

[`OAUTH`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-OAUTH)

The user is informed they will authenticate with the financial institution via OAuth.

[`PROFILE_DATA_REVIEW`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-PROFILE-DATA-REVIEW)

The user is asked to review their profile data in the Layer flow.

[`RECAPTCHA`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-RECAPTCHA)

The user was presented with a Google reCAPTCHA to verify they are human.

[`RISK_CHECK`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-RISK-CHECK)

The risk check step in the identity verification flow (configured via "Risk Rules" in the "Rulesets" section of the template editor).

[`SAME_DAY_MICRODEPOSIT_AUTHORIZED`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SAME-DAY-MICRODEPOSIT-AUTHORIZED)

The user has authorized a same day micro-deposit to be sent to their account over the ACH network with a 3-letter code to verify their account.

[`SCREENING`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SCREENING)

The watchlist screening step in the identity verification flow.

[`SELECT_ACCOUNT`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SELECT-ACCOUNT)

We ask the user to choose an account.

[`SELECT_AUTH_TYPE`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SELECT-AUTH-TYPE)

The user is asked to choose whether to Link instantly or manually (i.e., with micro-deposits).

[`SELECT_BRAND`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SELECT-BRAND)

The user is asked to select a brand, e.g. Bank of America. The brand selection interface occurs before the institution select pane and is only provided for large financial institutions with multiple online banking portals.

[`SELECT_INSTITUTION`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SELECT-INSTITUTION)

We ask the user to choose their institution.

[`SELECT_SAVED_ACCOUNT`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SELECT-SAVED-ACCOUNT)

The user is asked to select their saved accounts and/or new accounts for linking in the Link Returning User Experience (Remember Me) flow.

[`SELECT_SAVED_INSTITUTION`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SELECT-SAVED-INSTITUTION)

The user is asked to pick a saved institution or link a new one in the Link Remember Me flow.

[`SELFIE_CHECK`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SELFIE-CHECK)

The view in the identity verification flow which uses the camera to confirm there is a real user present that matches their ID documents.

[`SUBMIT_PHONE`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-SUBMIT-PHONE)

The user is asked for their phone number in the Link Returning User Experience (Remember Me) flow.

[`UPLOAD_DOCUMENTS`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-UPLOAD-DOCUMENTS)

The user is asked to upload documents (for Income verification).

[`VERIFY_PHONE`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-VERIFY-PHONE)

The user is asked to verify their phone in the Link Returning User Experience (Remember Me) flow or Layer flow.

[`VERIFY_SMS`](/docs/link/android/#link-android-onevent-LinkEventMetadata-viewName-VERIFY-SMS)

The SMS verification step in the identity verification flow.

[`metadataJson`](/docs/link/android/#link-android-onevent-LinkEventMetadata-metadataJson)

nullableStringnullable, String

The data directly returned from the server with no client side changes.

MainActivity

```
Plaid.setLinkEventListener { event -> Log.i("Event", event.toString()) }
```

=\*=\*=\*=

#### submit()

The `submit` function is currently only used in the Layer product. It allows the client application to submit additional user-collected data to the Link flow (e.g. a user phone number).

submit

**Properties**

[`submissionData`](/docs/link/android/#link-android-submit-submissionData)

objectobject

Data to submit during a Link session.

[`phoneNumber`](/docs/link/android/#link-android-submit-submissionData-phoneNumber)

StringString

The end user phone number.

[`dateOfBirth`](/docs/link/android/#link-android-submit-submissionData-dateOfBirth)

StringString

The end user date of birth. To be provided in the format "yyyy-mm-dd".

Code

```
val submissionData = SubmissionData(phoneNumber)
plaidHandler.submit(submissionData)
```

#### Upgrading

The latest version of the SDK is available from [GitHub](https://github.com/plaid/plaid-link-android). New versions of the SDK are released frequently. Major releases occur annually. The Link SDK uses Semantic Versioning, ensuring that all non-major releases are non-breaking, backwards-compatible updates. We recommend you update regularly (at least once a quarter, and ideally once a month) to ensure the best Plaid Link experience in your application.

SDK versions are supported for two years; with each major SDK release, Plaid will stop officially supporting any previous SDK versions that are more than two years old. While these older versions are expected to continue to work without disruption, Plaid will not provide assistance with unsupported SDK versions.

#### Next steps

If you run into problems integrating with Plaid Link on Android, see [Troubleshooting the Plaid Link Android SDK](/docs/link/android/troubleshooting/).

Once you've gotten Link working, see [Link best practices](/docs/link/best-practices/) for recommendations on further improving the Link flow.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
