---
title: "Link - iOS | Plaid Docs"
source_url: "https://plaid.com/docs/link/ios/"
scraped_at: "2026-03-07T22:05:06+00:00"
---

# Link iOS SDK

#### Reference for integrating with the Link iOS SDK

#### Overview

The Plaid Link SDK is a quick and secure way to link bank accounts to Plaid from within your iOS app.
LinkKit is a drop-in framework that handles connecting a financial institution
to your app (credential validation, multi-factor authentication, error handling, etc.) without passing sensitive information to your server.

Prefer to learn by watching? Check out this quick guide to implementing Plaid on iOS devices

Want even more video lessons? A [full tutorial](https://youtu.be/9fgmW38usTo) for integrating Plaid with iOS is available on our YouTube channel.

To get started with Plaid Link for iOS, clone the
[GitHub repository](https://github.com/plaid/plaid-link-ios) and try out the example application,
which provides a reference implementation in both Swift and Objective-C. Youʼll want to sign up for
[free API keys](https://dashboard.plaid.com/developers/keys) through the Plaid Dashboard to get started.

#### Initial iOS setup

Before writing code using the SDK, you must first perform some setup steps to register your app
with Plaid and configure your project.

The "Register your redirect URI" and "Set up Universal Links" steps describe setting up your app for OAuth redirects. If your integration uses only Link flows that do not initiate connections to financial institutions, such as the Identity Verification or Document Income flows, they can be skipped; they are mandatory otherwise.

##### Register your redirect URI

1. Sign in to the [Plaid Dashboard](https://dashboard.plaid.com/signin)
   and go to the [**Team Settings -> API**](https://dashboard.plaid.com/developers/api) page.
2. Next to **Allowed redirect URIs** click **Configure** then **Add New URI**.
3. Enter your redirect URI, which you must also set up as a Universal Link
   for your application, for example: `https://app.example.com/plaid/`.
4. Click **Save Changes**.

These redirect URIs must be set up as [Universal Links](https://developer.apple.com/ios/universal-links/) in your application.
For details, see Apple's documentation on [Allowing Apps and Websites to Link to Your Content](https://developer.apple.com/documentation/xcode/allowing-apps-and-websites-to-link-to-your-content).

Plaid does not support registering URLs with custom URL schemes as redirect URIs
since they lack the security of Universal Links through the two-way association between
your app and your website. Any application can register custom URL schemes and there is no further
validation from iOS. If multiple applications have registered the same custom URL scheme,
a different application may be launched each time the URL is opened.
To complete Plaid OAuth flows, it is important that *your* application is
opened and not any arbitrary application that has registered the same URL scheme.

##### Set up Universal Links

In order for Plaid to return control back to your application after a user
completes a bank's OAuth flow, you must specify a redirect URI, which will be the URI from which
Link will be re-launched to complete the OAuth flow. The redirect URI must be a
Universal Link.
An example of a typical redirect URI is: `https://app.example.com/plaid`.

Universal Links consist of the following parts:

- An `applinks` entitlement for the Associated Domains capability in your app.
  For details, see Apple’s documentation on the
  [Associated Domains Entitlement](https://developer.apple.com/documentation/bundleresources/entitlements/com_apple_developer_associated-domains) and
  the [entitlements example](https://github.com/plaid/plaid-link-ios/blob/master/LinkDemo-Swift/LinkDemo-Swift/LinkDemo-Swift.entitlements) in the LinkDemo-Swift example app.
- An `apple-app-site-association` file on your website.
  For details, see Apple’s documentation on
  [Supporting Associated Domains](https://developer.apple.com/documentation/xcode/supporting-associated-domains) and see our minimal example below.

There are a few requirements for `apple-app-site-association` files:

- Must be a static JSON file
- Must be hosted using a `https://` scheme with a valid certificate and no redirects
- Must be hosted at `https://<my-fully-qualified-domain>/.well-known/apple-app-site-association`

Below is an example for `https://my-app.com` (`https://my-app.com/.well-known/apple-app-site-association`)

```
{
  "applinks": {
    "details": [
      {
        "appIDs": ["<My Application Identifier Prefix>.<My Bundle ID>"],
        "components": [
          {
            "/": "/plaid/*",
            "comment": "Matches any URL path whose path starts with /plaid/"
          }
        ]
      }
    ]
  }
}
```

Once you have enabled Universal Links, your iOS app is now set up and ready to start integrating with the Plaid SDK.

Ensure that the corresponding entry for the configured redirect URI in the `apple-app-site-association` file on your website continues to be available. If it is removed, OAuth sessions will fail until it is available again.

#### Installation

Plaid Link for iOS is an embeddable framework that is bundled and distributed with your application.
There are several ways to obtain the necessary files and keep them up-to-date; we
recommend using [Swift Package Manager](https://swiftpackageindex.com/plaid/plaid-link-ios) or [CocoaPods](https://cocoapods.org/pods/Plaid).
Regardless of what you choose, submitting a new version of your application with the updated
[`LinkKit.xcframework`](https://github.com/plaid/plaid-link-ios/releases) to the App Store is required.

##### Requirements

| LinkKit iOS version | Xcode toolchain minimum support | Supported iOS versions |
| --- | --- | --- |
| LinkKit 5.x.x | Xcode 15 | iOS 14 or greater |
| LinkKit 4.x.x | Xcode 14 | iOS 11 or greater |
| LinkKit 3.x.x | Xcode 13 | iOS 11 or greater |

Select group for content switcher

##### Swift Package Manager

1. To integrate LinkKit using Swift Package Manager, Swift version >= 5.3 is required.
2. In your Xcode project from the Project Navigator (**Xcode ❯ View ❯ Navigators ❯ Project ⌘ 1**)
   select your project, activate the **Package Dependencies** tab and click on the plus symbol **➕**
   to open the **Add Package** popup window:

   ![Xcode interface showing how to add a Swift package dependency; plus button highlighted.](/assets/img/docs/linkkit/spm_package_dependency.png)
3. Enter the LinkKit package URL `https://github.com/plaid/plaid-link-ios-spm` into the search bar
   in the top right corner of the **Add Package** popup window. The main repository with full git history is very large (~1 GB), and Swift Package Manager always downloads the full repository with all git history. This `plaid-link-ios-spm` repository is much smaller (less than 500kb), making the download faster.
4. Select the `plaid-link-ios-spm` package.
5. Choose your **Dependency Rule** (we recommend **Up to Next Major Version**).
6. Select the project to which you would like to add LinkKit, then click **Add Package**:

   ![Xcode screenshot showing input of Plaid Link iOS Swift package URL for package addition. 'Add Package' button highlighted.](/assets/img/docs/linkkit/spm_package_url.png)
7. Select the `LinkKit` package product and click **Add Package**:

   ![Dialog in Xcode for selecting Plaid link-ios package. 'LinkKit' checked, type 'Library', target 'MySampleApp'.](/assets/img/docs/linkkit/spm_add_package.png)
8. Verify that the LinkKit Swift package was properly added as a package dependency to your project:

   ![Xcode showing the Plaid Link iOS package dependency from GitHub in Package Dependencies tab.](/assets/img/docs/linkkit/spm_added_to_project.png)
9. Select your application target and ensure that the LinkKit library is embedded into your application:

   ![Xcode screenshot showing the Frameworks section with LinkKit listed under Package Dependencies for MySampleApp.](/assets/img/docs/linkkit/spm_embedded.png)

##### CocoaPods

1. If you haven’t already, install the latest version of [CocoaPods](https://guides.cocoapods.org/using/getting-started.html).
2. If you don’t have an existing Podfile, run the following command to create one:

   Create a new Podfile

   ```
   pod init
   ```
3. Add this line to your `Podfile`:

   Edit the Podfile to include the following line

   ```
   pod 'Plaid'
   ```
4. Run the following command:

   Install the Plaid and other CocoaPods

   ```
   pod install
   ```
5. To update to newer releases in the future, run:

   Update the Plaid and other CocoaPods

   ```
   pod install
   ```

##### Manual

Get the latest version of the [`LinkKit.xcframework`](https://github.com/plaid/plaid-link-ios/releases)
and embed it into your application, for example by dragging and dropping the XCFramework bundle
onto the **Embed Frameworks** build phase of your application target in Xcode as shown below.

![Xcode project with LinkKit.xcframework added under Frameworks, Libraries, and Embedded Content. LinkKit.xcframework folder in Downloads.](/assets/img/docs/linkkit/embed.jpg)

##### Camera Support (Identity Verification only)

When using the [Identity Verification](/docs/identity-verification/) product, the Link SDK may use the camera if a user
needs to take a picture of identity documentation. To support this workflow, add a [`NSCameraUsageDescription`](https://developer.apple.com/documentation/bundleresources/information_property_list/nscamerausagedescription) entry to your application's plist
with an informative string. This allows iOS to prompt the user for camera access. iOS will crash your application if
this string is not provided.

##### Upgrading

New versions of [`LinkKit.xcframework`](https://github.com/plaid/plaid-link-ios/releases)
are released frequently. Major releases occur annually. The Link SDK uses Semantic Versioning, ensuring that all non-major releases are non-breaking, backwards-compatible updates. We recommend you update regularly (at least once a quarter, and ideally once a month) to ensure the best Plaid Link experience in your application.

SDK versions are supported for two years; with each major SDK release, Plaid will stop officially supporting any previous SDK versions that are more than two years old. While these older versions are expected to continue to work without disruption, Plaid will not provide assistance with unsupported SDK versions.

For details on each release see the [GitHub releases](https://github.com/plaid/plaid-link-ios/releases) and [version release notes](https://github.com/plaid/plaid-link-ios/blob/master/CHANGELOG.md).

#### Opening Link

Before you can open Link, you need to first create a `link_token`. A `link_token` can be configured for
different Link flows and is used to control much of Link’s behavior. To learn how to create a new
`link_token`, see the API Reference entry for [`/link/token/create`](/docs/api/link/#linktokencreate).

For iOS the [`/link/token/create`](/docs/api/link/#linktokencreate) call must include the `redirect_uri` parameter and it must
match the redirect URI you have configured with Plaid
(see [**Register your redirect URI**](/docs/link/ios/#register-your-redirect-uri) above).

##### Create a Configuration

Starting the Plaid Link for iOS experience begins with creating a `link_token`. Once the `link_token` is passed to your app, create an instance of `LinkTokenConfiguration`, then create a `Handler` using `Plaid.create()` passing the previously created `LinkTokenConfiguration`. Plaid will begin to pre-load Link as soon as `Plaid.create()` is called, so to reduce UI latency when rendering Link, you should call `Plaid.create()` when initializing the screen where the user can enter the Link flow.

After calling `Plaid.create()`, to show the Link UI, call `open()` on the handler passing your preferred presentation method.

Note that each time you open Link, you will need to get a new `link_token` from your server and create a new `LinkTokenConfiguration` with it.

create

**Properties**

[`token`](/docs/link/ios/#link-ios-create-token)

StringString

Specify a `link_token` to authenticate your app with Link. This is a short lived, one-time use token that should be unique for each Link session. In addition to the primary flow, a `link_token` can be configured to launch Link in [update mode](/docs/link/update-mode/). See the `/link/token/create` endpoint for more details.

[`onSuccess`](/docs/link/ios/#link-ios-create-onSuccess)

closureclosure

A required closure that is called when a user successfully links an Item. The closure should expect a single `LinkSuccess` argument, containing the `publicToken` String and a `metadata` of type `SuccessMetadata`. See [onSuccess](#onsuccess) below.

[`onExit`](/docs/link/ios/#link-ios-create-onExit)

closureclosure

An optional closure that is called when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. The closure should expect a single `LinkExit` argument, containing an optional `error` and a `metadata` of type `ExitMetadata`. See [onExit](#onexit) below.

[`onEvent`](/docs/link/ios/#link-ios-create-onEvent)

closureclosure

An optional closure that is called when a user reaches certain points in the Link flow. The closure should expect a single `LinkEvent` argument, containing an `eventName` enum of type `EventType` and a `metadata` of type `EventMetadata`. See [onEvent](#onevent) below.

Create a Configuration

```
var linkConfiguration = LinkTokenConfiguration(
    token: "<#LINK_TOKEN_FROM_SERVER#>",
    onSuccess: { linkSuccess in
        // Send the linkSuccess.publicToken to your app server.
    }
)

// Optional: Configure additional Link presentation options. 

// If true, skips the default Link loading animation.
// Use this if you're displaying your own custom loading UI.
linkConfiguration.noLoadingState = true

// Optional: Configure additional Link presentation options -
// Available in SDK v6.2.0 and above:

// If false, displays a solid background instead of the default
// transparent gradient.
linkConfiguration.showGradientBackground = false
```

##### Create a Handler

A `Handler` is a one-time use object used to open a Link session. The `Handler` must
be retained for the duration of the Plaid SDK flow. It will also be needed to respond to OAuth Universal Link
redirects. For more details, see the [OAuth guide](/docs/link/oauth/#ios).

You can optionally provide an `onLoad` (available in **SDK v6.2.0 and above**) callback to know when Link has finished preloading and is ready to be presented. This is useful if you want to delay presenting Link until it’s fully loaded, or enable related UI elements.

Note for Layer customers: You should wait for the `LAYER_READY` event before presenting Link, rather than relying solely on `onLoad`.

Create a Handler

```
let result = Plaid.create(configuration, onLoad: { [weak self] in
    // Optional callback invoked once Plaid Link has finished preloading and is ready
    // to be presented.
    //
    // Example: Enable a button once Link has loaded
    // self?.openButton.isEnabled = true
    //
    // Example: Automatically present Link once it has loaded
    // guard let self = self else { return }
    // self.handler?.open(presentUsing: .viewController(self))
})

switch result {
case .failure(let error):
    logger.error("Unable to create Plaid handler due to: \(error)")
case .success(let handler):
    // Retain the handler for the duration of the Link flow (and OAuth redirects)
    self.handler = handler
}
```

##### Open Link

Finally, open Link by calling `open` on the `Handler` object.
This will usually be done in a button’s target action.

Open Link

```
let method: PresentationMethod = .viewController(self)
handler.open(presentUsing: method)
```

=\*=\*=\*=

#### onSuccess

The closure is called when a user successfully links an Item. It should take a single `LinkSuccess` argument, containing the `publicToken` String and a `metadata` of type `SuccessMetadata`.

onSuccess

**Properties**

[`linkSuccess`](/docs/link/ios/#link-ios-onsuccess-linkSuccess)

LinkSuccessLinkSuccess

Contains the `publicToken` and `metadata` for this successful flow.

[`publicToken`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-publicToken)

StringString

Displayed once a user has successfully completed Link. If using Identity Verification or Beacon, this field will be `null`. If using Document Income or Payroll Income, the `public_token` will be returned, but is not used.

[`metadata`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata)

LinkSuccessLinkSuccess

Displayed once a user has successfully completed Link.

[`institution`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-institution)

nullableInstitutionnullable, Institution

An institution object. If the Item was created via Same-Day micro-deposit verification, will be `null`.

[`name`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-institution-name)

StringString

The full institution name, such as 'Wells Fargo'

[`institutionID`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-institution-institutionID)

InstitutionID (String)InstitutionID (String)

The Plaid institution identifier

[`accounts`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-accounts)

Array<Account>Array<Account>

A list of accounts attached to the connected Item

[`id`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-accounts-id)

AccountID (String)AccountID (String)

The Plaid `account_id`

[`name`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-accounts-name)

StringString

The official account name

[`mask`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-accounts-mask)

Optional<AccountMask> (Optional<String>)Optional<AccountMask> (Optional<String>)

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, it may also not match the mask that the bank displays to the user.

[`subtype`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-accounts-subtype)

AccountSubtypeAccountSubtype

The account subtype and its type.
See [Account Types](/docs/api/accounts/#account-type-schema) for a full list of possible values.

[`verificationStatus`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-accounts-verificationStatus)

Optional<VerificationStatus>Optional<VerificationStatus>

Indicates an Item's micro-deposit-based verification or database verification status. Possible values are:  
`pending_automatic_verification`: The Item is pending automatic verification.  
`pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the deposit.  
`automatically_verified`: The Item has successfully been automatically verified.  
`manually_verified`: The Item has successfully been manually verified.  
`verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  
`verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.  
`database_matched`: The Item has successfully been verified using Plaid's data sources.  
`database_insights_pending`: The Database Insights result is pending and will be available upon Auth request.  
`nil`: Neither micro-deposit-based verification nor database verification are being used for the Item.

[`linkSessionID`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-linkSessionID)

StringString

A unique identifier associated with a user's actions and events through the Link flow.
Include this identifier when opening a support ticket for faster turnaround.

[`metadataJSON`](/docs/link/ios/#link-ios-onsuccess-linkSuccess-metadata-metadataJSON)

StringString

Unprocessed metadata, formatted as JSON, sent from the Plaid API.

Handle Link success

```
onSuccess: { linkSuccess in
  // Send the linkSuccess.publicToken to your app server.
}
```

=\*=\*=\*=

#### onExit

This optional closure is called when a user exits Link without successfully linking an Item, or when an error occurs during Link initialization. It should take a single `LinkExit` argument, containing an optional `error` and a `metadata` of type `ExitMetadata`.

onExit

**Properties**

[`linkExit`](/docs/link/ios/#link-ios-onexit-linkExit)

LinkExitLinkExit

Contains the optional `error` and `metadata` for when the flow was exited.

[`error`](/docs/link/ios/#link-ios-onexit-linkExit-error)

Optional<ExitError> (Swift), Optional<NSError> (Objective-C)Optional<ExitError> (Swift), Optional<NSError> (Objective-C)

An Error type that contains the `errorCode`, `errorMessage`, and `displayMessage` of the error that was last encountered by the user. If no error was encountered, `error` will be `nil`. In Objective-C, field names will match the `NSError` type.

[`errorCode`](/docs/link/ios/#link-ios-onexit-linkExit-error-errorCode)

ExitErrorCodeExitErrorCode

The error code and error type that the user encountered.
Each `errorCode` has an associated `errorType`, which is a broad categorization of the error.

[`errorMessage`](/docs/link/ios/#link-ios-onexit-linkExit-error-errorMessage)

StringString

A developer-friendly representation of the error code.

[`displayMessage`](/docs/link/ios/#link-ios-onexit-linkExit-error-displayMessage)

Optional<String>Optional<String>

A user-friendly representation of the error code or `nil` if the error is not related to user action. This may change over time and is not safe for programmatic use.

[`metadata`](/docs/link/ios/#link-ios-onexit-linkExit-metadata)

ExitMetadataExitMetadata

Displayed if a user exits Link without successfully linking an Item.

[`status`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status)

Optional<ExitStatus>Optional<ExitStatus>

The status key indicates the point at which the user exited the Link flow.

[`requiresQuestions`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-requiresQuestions)

User prompted to answer security question(s).

[`requiresSelections`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-requiresSelections)

User prompted to answer multiple choice question(s).

[`requiresCode`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-requiresCode)

User prompted to provide a one-time passcode.

[`chooseDevice`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-chooseDevice)

User prompted to select a device on which to receive a one-time passcode.

[`requiresCredentials`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-requiresCredentials)

User prompted to provide credentials for the selected financial institution or has not yet selected a financial institution.

[`requiresAccountSelection`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-requiresAccountSelection)

User prompted to select one or more financial accounts to share

[`institutionNotFound`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-institutionNotFound)

User exited the Link flow on the institution selection pane. Typically this occurs after the user unsuccessfully (no results returned) searched for a financial institution. Note that this status does not necessarily indicate that the user was unable to find their institution, as it is used for all user exits that occur from the institution selection pane, regardless of other user behavior.

[`institutionNotSupported`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-institutionNotSupported)

User exited the Link flow after discovering their selected institution is no longer supported by Plaid

[`unknown`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-status-unknown)

The exit status has not been defined in the current version of the SDK. The `unknown` case has an associated value carrying the original exit status as sent by the Plaid API.

[`institution`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-institution)

Optional<Institution>Optional<Institution>

An institution object. If the Item was created via Same-Day micro-deposit verification, will be omitted.

[`institutionID`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-institution-institutionID)

InstitutionID (String)InstitutionID (String)

The Plaid specific identifier for the institution.

[`name`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-institution-name)

StringString

The full institution name, such as 'Wells Fargo'.

[`linkSessionID`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-linkSessionID)

Optional<String>Optional<String>

A unique identifier associated with a user's actions and events through the Link flow.
Include this identifier when opening a support ticket for faster turnaround.

[`requestID`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-requestID)

Optional<String>Optional<String>

The request ID for the last request made by Link.
This can be shared with Plaid Support to expedite investigation.

[`metadataJSON`](/docs/link/ios/#link-ios-onexit-linkExit-metadata-metadataJSON)

RawJSONMetadata (String)RawJSONMetadata (String)

Unprocessed metadata, formatted as JSON, sent from Plaid API.

Handle Link exit

```
linkConfiguration.onExit = { linkExit in
  // Optionally handle linkExit data according to your application's needs
}
```

=\*=\*=\*=

#### onEvent

This closure is called when certain events in the Plaid Link flow have occurred. The `open`, `layerReady`, `layerNotAvailable`, `layerAutofillNotAvailable` events are guaranteed to fire in real time; other events will typically be fired when the Link session finishes, when `onSuccess` or `onExit` is called. Callback ordering is not guaranteed; `onEvent` callbacks may fire before, after, or surrounding the `onSuccess` or `onExit` callback, and event callbacks are not guaranteed to fire in the order in which they occurred.

The following `onEvent` callbacks are stable, which means that they are suitable for programmatic use in your application's logic: `open`, `exit`, `handoff`, `selectInstitution`, `error`, `bankIncomeInsightsCompleted`, `identityVerificationPassSession`, `identityVerificationFailSession`, `layerReady`, `layerNotAvailable`, `layerAutofillNotAvailable`. The remaining callback events are informational and subject to change, and should be used for analytics and troubleshooting purposes only.

onEvent

**Properties**

[`linkEvent`](/docs/link/ios/#link-ios-onevent-linkEvent)

LinkEventLinkEvent

Contains the `eventName` and `metadata` for the Link event.

[`eventName`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName)

EventNameEventName

An enum representing the event that has just occurred in the Link flow.

[`autoSubmitPhone`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-autoSubmitPhone)

The user was automatically sent an OTP code without a UI prompt. This event can only occur if the user's phone phone number was provided to Link via the `/link/token/create` call and the user has previously consented to receive OTP codes from Plaid.

[`bankIncomeInsightsCompleted`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-bankIncomeInsightsCompleted)

The user has completed the Assets and Bank Income Insights flow.

[`closeOAuth`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-closeOAuth)

The user closed the third-party website or mobile app without completing the OAuth flow.

[`connectNewInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-connectNewInstitution)

The user has chosen to link a new institution instead of linking a saved institution. This event is only emitted in the Link Remember Me flow.

[`error`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-error)

A recoverable error occurred in the Link flow, see the `errorCode` in the metadata.

[`exit`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-exit)

The user has exited without completing the Link flow and the [onExit](#onexit) callback is fired.

[`failOAuth`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-failOAuth)

The user encountered an error while completing the third-party's OAuth login flow.

[`handoff`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-handoff)

The user has exited Link after successfully linking an Item.

[`identityMatchFailed`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityMatchFailed)

An Identity Match check configured via the Account Verification Dashboard failed the Identity Match rules and did not detect a match.

[`identityMatchPassed`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityMatchPassed)

An Identity Match check configured via the Account Verification Dashboard passed the Identity Match rules and detected a match.

[`identityVerificationStartStep`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationStartStep)

The user has started a step of the Identity Verification flow. The step is indicated by `view_name`.

[`identityVerificationPassStep`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationPassStep)

The user has passed a step of the Identity Verification flow. The step is indicated by `view_name`.

[`identityVerificationFailStep`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationFailStep)

The user has failed a step of the Identity Verification flow. The step is indicated by `view_name`.

[`identityVerificationPendingReviewStep`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationPendingReviewStep)

The user has reached the pending review state.

[`identityVerificationCreateSession`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationCreateSession)

The user has started a new Identity Verification session.

[`identityVerificationResumeSession`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationResumeSession)

The user has resumed an existing Identity Verification session.

[`identityVerificationPassSession`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationPassSession)

The user has passed their Identity Verification session.

[`identityVerificationFailSession`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationFailSession)

The user has failed their Identity Verification session.

[`identityVerificationPendingReviewSession`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationPendingReviewSession)

The user has completed their Identity Verification session, which is now in a pending review state.

[`identityVerificationOpenUI`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationOpenUI)

The user has opened the UI of their Identity Verification session.

[`identityVerificationResumeUI`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationResumeUI)

The user has resumed the UI of their Identity Verification session.

[`identityVerificationCloseUI`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-identityVerificationCloseUI)

The user has closed the UI of their Identity Verification session.

[`layerAutofillNotAvailable`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-layerAutofillNotAvailable)

The user's date of birth passed to Link is not eligible for Layer Extended Autofill.

[`layerNotAvailable`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-layerNotAvailable)

The user phone number passed to Link is not eligible for Layer.

[`layerReady`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-layerReady)

The user phone number passed to Link is eligible for Layer and `open()` may now be called.

[`matchedSelectInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-matchedSelectInstitution)

The user selected an institution that was presented as a matched institution. This event can be emitted if [Embedded Institution Search](https://plaid.com/docs/link/embedded-institution-search/) is being used, if the institution was surfaced as a matched institution likely to have been linked to Plaid by a returning user, or if the institution's `routing_number` was provided when calling `/link/token/create`. For details on which scenario is triggering the event, see `metadata.matchReason`.

[`open`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-open)

The user has opened Link.

[`openMyPlaid`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-openMyPlaid)

The user has opened my.plaid.com. This event is only sent when Link is initialized with Assets as a product.

[`openOAuth`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-openOAuth)

The user has navigated to a third-party website or mobile app in order to complete the OAuth login flow.

[`searchInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-searchInstitution)

The user has searched for an institution.

[`selectAuthType`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-selectAuthType)

The user has chosen whether to Link instantly or manually (i.e., with micro-deposits). This event emits the `selection` metadata to indicate the user's selection.

[`selectBrand`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-selectBrand)

The user selected a brand, e.g. Bank of America. The brand selection interface occurs before the institution select pane and is only provided for large financial institutions with multiple online banking portals.

[`selectDegradedInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-selectDegradedInstitution)

The user selected an institution with a `DEGRADED` health status and was shown a corresponding message.

[`selectDownInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-selectDownInstitution)

The user selected an institution with a `DOWN` health status and was shown a corresponding message.

[`selectFilteredInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-selectFilteredInstitution)

The user selected an institution Plaid does not support all requested products for and was shown a corresponding message.

[`selectInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-selectInstitution)

The user selected an institution.

[`skipSubmitPhone`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-skipSubmitPhone)

The user has opted to not provide their phone number to Plaid. This event is only emitted in the Link Remember Me flow.

[`submitAccountNumber`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitAccountNumber)

The user has submitted an account number. This event emits the `accountNumberMask` metadata to indicate the mask of the account number the user provided.

[`submitCredentials`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitCredentials)

The user has submitted credentials.

[`submitDocuments`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitDocuments)

The user is being prompted to submit documents for an Income verification flow.

[`submitDocumentsError`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitDocumentsError)

The user encountered an error when submitting documents for an Income verification flow.

[`submitDocumentsSuccess`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitDocumentsSuccess)

The user has successfully submitted documents for an Income verification flow.

[`submitMFA`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitMFA)

The user has submitted MFA.

[`submitOTP`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitOTP)

The user has submitted an OTP code during the phone number verification flow. This event is only emitted in the Link Returning User Experience flow.

[`submitPhone`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitPhone)

The user has submitted their phone number. This event is only emitted in the Link Remember Me flow.

[`submitRoutingNumber`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-submitRoutingNumber)

The user has submitted a routing number. This event emits the `routingNumber` metadata to indicate user's routing number.

[`transitionView`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-transitionView)

The `transitionView` event indicates that the user has moved from one view to the next.

[`uploadDocuments`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-uploadDocuments)

The user is asked to upload documents (for Income verification).

[`verifyPhone`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-verifyPhone)

The user has successfully verified their phone number. This event is only emitted in the Link Remember Me flow.

[`viewDataTypes`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-viewDataTypes)

The user has viewed data types on the data transparency consent pane.

[`unknown`](/docs/link/ios/#link-ios-onevent-linkEvent-eventName-unknown)

The event has not been defined in the current version of the SDK.
The `unknown` case has an associated value carrying the original event name as sent by the Plaid API.

[`metadata`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata)

EventMetadata structEventMetadata struct

An object containing information about the event

[`accountNumberMask`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-accountNumberMask)

Optional<String>Optional<String>

The account number mask extracted from the user-provided account number. If the user-inputted account number is four digits long, `account_number_mask` is empty. Emitted by `SUBMIT_ACCOUNT_NUMBER`.

[`errorCode`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-errorCode)

Optional<ExitErrorCode>Optional<ExitErrorCode>

The error code that the user encountered. Emitted by: `error`, `exit`.

[`errorMessage`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-errorMessage)

Optional<String>Optional<String>

The error message that the user encountered. Emitted by: `error`, `exit`.

[`exitStatus`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-exitStatus)

Optional<ExitStatus>Optional<ExitStatus>

The status key indicates the point at which the user exited the Link flow. Emitted by: `exit`.

[`requiresQuestions`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-exitStatus-requiresQuestions)

User prompted to answer security question(s).

[`requiresSelections`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-exitStatus-requiresSelections)

User prompted to answer multiple choice question(s).

[`requiresCode`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-exitStatus-requiresCode)

User prompted to provide a one-time passcode.

[`chooseDevice`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-exitStatus-chooseDevice)

User prompted to select a device on which to receive a one-time passcode.

[`requiresCredentials`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-exitStatus-requiresCredentials)

User prompted to provide credentials for the selected financial institution or has not yet selected a financial institution.

[`institutionNotFound`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-exitStatus-institutionNotFound)

User exited the Link flow on the institution selection pane. Typically this occurs after the user unsuccessfully (no results returned) searched for a financial institution. Note that this status does not necessarily indicate that the user was unable to find their institution, as it is used for all user exits that occur from the institution selection pane, regardless of other user behavior.

[`unknown`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-exitStatus-unknown)

The exit status has not been defined in the current version of the SDK.
The unknown case has an associated value carrying the original exit status as sent by the Plaid API.

[`institutionID`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-institutionID)

Optional<InstitutionID> (Optional<String>)Optional<InstitutionID> (Optional<String>)

The ID of the selected institution. Emitted by: *all events*.

[`institutionName`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-institutionName)

Optional<String>Optional<String>

The name of the selected institution. Emitted by: *all events*.

[`institutionSearchQuery`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-institutionSearchQuery)

Optional<String>Optional<String>

The query used to search for institutions. Emitted by: `searchInstitution`.

[`isUpdateMode`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-isUpdateMode)

Optional<String>Optional<String>

Indicates if the current Link session is an update mode session. Emitted by: `OPEN`.

[`matchReason`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-matchReason)

nullablestringnullable, string

The reason this institution was matched.
This will be either `returning_user` or `routing_number` if emitted by: `MATCHED_SELECT_INSTITUTION`.
Otherwise, this will be `SAVED_INSTITUTION` or `AUTO_SELECT_SAVED_INSTITUTION` if emitted by: `SELECT_INSTITUTION`.

[`linkSessionID`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-linkSessionID)

StringString

The `link_session_id` is a unique identifier for a single session of Link. It's always available and will stay constant throughout the flow. Emitted by: *all events*.

[`mfaType`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-mfaType)

Optional<MFAType>Optional<MFAType>

If set, the user has encountered one of the following MFA types: `code`, `device`, `questions`, `selections`. Emitted by: `submitMFA` and `transitionView` when `viewName` is `mfa`

[`requestID`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-requestID)

Optional<String>Optional<String>

The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation. Emitted by: *all events*.

[`routingNumber`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-routingNumber)

Optional<String>Optional<String>

The routing number submitted by user at the micro-deposits routing number pane. Emitted by `SUBMIT_ROUTING_NUMBER`.

[`selection`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-selection)

Optional<String>Optional<String>

The Auth Type Select flow type selected by the user. Possible values are `flow_type_manual` or `flow_type_instant`. Emitted by: `SELECT_AUTH_TYPE`.

[`timestamp`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-timestamp)

DateDate

An ISO 8601 representation of when the event occurred. For example `2017-09-14T14:42:19.350Z`. Emitted by: *all events*.

[`viewName`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName)

Optional<ViewName>Optional<ViewName>

The name of the view that is being transitioned to. Emitted by: `transitionView`.

[`acceptTOS`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-acceptTOS)

The view showing Terms of Service in the identity verification flow.

[`connected`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-connected)

The user has connected their account.

[`consent`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-consent)

We ask the user to consent to the privacy policy.

[`credential`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-credential)

Asking the user for their account credentials.

[`documentaryVerification`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-documentaryVerification)

The view requesting document verification in the identity verification flow (configured via "Fallback Settings" in the "Rulesets" section of the template editor).

[`error`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-error)

An error has occurred.

[`exit`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-exit)

Confirming if the user wishes to close Link.

[`instantMicrodepositAuthorized`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-instantMicrodepositAuthorized)

The user has authorized an instant micro-deposit to be sent to their account over the RTP or FedNow network with a 3-letter code to verify their account.

[`kycCheck`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-kycCheck)

The view representing the "know your customer" step in the identity verification flow.

[`loading`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-loading)

Link is making a request to our servers.

[`mfa`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-mfa)

The user is asked by the institution for additional MFA authentication.

[`numbers`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-numbers)

The user is asked to insert their account and routing numbers.

[`numbersSelectInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-numbersSelectInstitution)

The user goes through the Same Day micro-deposits flow with Reroute to Credentials.

[`oauth`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-oauth)

The user is informed they will authenticate with the financial institution via OAuth.

[`profileDataReview`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-profileDataReview)

The user is asked to review their profile data in the Layer flow.

[`recaptcha`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-recaptcha)

The user was presented with a Google reCAPTCHA to verify they are human.

[`riskCheck`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-riskCheck)

The risk check step in the identity verification flow (configured via "Risk Rules" in the "Rulesets" section of the template editor).

[`sameDayMicrodepositAuthorized`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-sameDayMicrodepositAuthorized)

The user has authorized a same day micro-deposit to be sent to their account over the ACH network with a 3-letter code to verify their account.

[`screening`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-screening)

The watchlist screening step in the identity verification flow.

[`selectAccount`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-selectAccount)

We ask the user to choose an account.

[`selectAuthType`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-selectAuthType)

The user is asked to choose whether to Link instantly or manually (i.e., with micro-deposits).

[`selectBrand`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-selectBrand)

The user is asked to select a brand, e.g. Bank of America. The brand selection interface occurs before the institution select pane and is only provided for large financial institutions with multiple online banking portals.

[`selectInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-selectInstitution)

We ask the user to choose their institution.

[`selectSavedAccount`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-selectSavedAccount)

The user is asked to select their saved accounts and/or new accounts for linking in the Link Remember Me flow.

[`selectSavedInstitution`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-selectSavedInstitution)

The user is asked to pick a saved institution or link a new one in the Link Remember Me flow.

[`selfieCheck`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-selfieCheck)

The view in the identity verification flow which uses the camera to confirm there is real user that matches their ID documents.

[`submitPhone`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-submitPhone)

The user is asked for their phone number in the Link Remember Me flow.

[`uploadDocuments`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-uploadDocuments)

The user is asked to upload documents (for Income verification).

[`verifyPhone`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-verifyPhone)

The user is asked to verify their phone OTP in the Link Remember Me flow.

[`verifySMS`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-verifySMS)

The SMS verification step in the identity verification flow.

[`unknown`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-viewName-unknown)

The view has not been defined in the current version of the SDK.
The unknown case has an associated value carrying the original view name as sent by the
Plaid API.

[`metadataJSON`](/docs/link/ios/#link-ios-onevent-linkEvent-metadata-metadataJSON)

RawJSONMetadata (String)RawJSONMetadata (String)

Unprocessed metadata, formatted as JSON, sent from Plaid API.

Handle Link event

```
linkConfiguration.onEvent = { linkEvent in
  // Optionally handle linkEvent data according to your application's needs
}
```

=\*=\*=\*=

#### submit()

The `submit` function is currently only used in the Layer product. It allows the client application to submit additional user-collected data to the Link flow (e.g. a user phone number).

submit

**Properties**

[`submissionData`](/docs/link/ios/#link-ios-submit-submissionData)

objectobject

Data to submit during a Link session.

[`phoneNumber`](/docs/link/ios/#link-ios-submit-submissionData-phoneNumber)

StringString

The end user's phone number.

[`dateOfBirth`](/docs/link/ios/#link-ios-submit-submissionData-dateOfBirth)

StringString

The end user's date of birth. To be provided in the format "yyyy-mm-dd".

Submit phone number

```
// Create a model that conforms to the SubmissionData interface
struct PlaidSubmitData: SubmissionData {
    var phoneNumber: String?
}

let data = PlaidSubmitData(phoneNumber: "14155550015")

self.handler.submit(data)
```

#### Next steps

Once you've gotten Link working, see [Link best practices](https://plaid.com/docs/link/best-practices/) for recommendations on further improving the Link flow.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
