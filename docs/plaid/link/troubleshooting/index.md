---
title: "Link - Troubleshooting | Plaid Docs"
source_url: "https://plaid.com/docs/link/troubleshooting/"
scraped_at: "2026-03-07T22:05:08+00:00"
---

# Link errors and troubleshooting

#### Errors in the Link flow

#### Link error codes

Any errors during Link will be returned via the Link flow; for details, see the guide for your specific platform. Since any failed user attempt to log on to a financial institution for any reason, even an incorrect password, will result in a Link error, you should expect to see the occasional error during Link. In most cases, Link will guide your user through the steps required to resolve the error.

The most common case in which you will need a special flow for Link errors is the `ITEM_LOGIN_REQUIRED` error. For instructions on handling this error, see [Update mode](/docs/link/update-mode/).

For more details, or if you need to troubleshoot a problem involving an error message from the API, see [Item errors](/docs/errors/item/).

#### Missing institutions or "Connectivity not supported" error

##### Symptom

- Your user cannot find their financial institution in Link, even though it is one of the 11,500+ institutions supported by Plaid.
- Your user experiences a "Connectivity not supported" error message after selecting their institution in Link.

##### Common causes

- The institution does not support one of the products specified in Link initialization.
- The institution is associated with a country not specified in Link initialization.
- The institution is associated with a country your Plaid account hasn't been enabled for.

##### Troubleshooting steps

Make sure not to initialize Link with any products your application isn't using. For more details on which products to initialize Link with, see [Choosing how to initialize products](/docs/link/initializing-products/).

Make sure Link is initialized for the country in which the institution operates.

If the problem occurs in Production but not in Sandbox, your Plaid developer account may not have been enabled for the
country in which the institution operates. Submit a [product access
request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access)
via the Plaid Dashboard to gain access.

#### Institution unhealthy

##### Symptom

When a user tries to connect their account, they receive a message in Link that Plaid may be experiencing connectivity issues to that institution. For more details, see [Institution status in Link](/docs/link/institution-status/).

##### Common causes

- Plaid's connection to the institution is down or degraded.

##### Troubleshooting steps

Enroll in [Link Recovery (beta)](/docs/link/link-recovery/). Link Recovery allows your end users to sign up for automatic notifications from Plaid when the connection issue is resolved so that they can retry connecting their account.

If your user cannot link their account, ask them to try again later.

If possible, provide a manual, backup flow for the user to enter account information.

#### Duplicate Items

##### Symptom

Your user has multiple Items that seem to represent the same underlying set of accounts.

##### Common causes

- The user added the same Item more than once.

##### Troubleshooting steps

See [Preventing duplicate Items](/docs/link/duplicate-items/) for instructions on preventing this scenario.

#### Missing accounts

##### Symptom

Some of the end users' accounts are not present in the Item, or a new account added by the user does not appear in the Item.

##### Common causes

- The user linked their account but did not grant your app permission to access the account, or created the account after linking the Item and did not grant your app permission to access newly created accounts.
- The account is not compatible with settings specified during the [`/link/token/create`](/docs/api/link/#linktokencreate) call. For example, if `auth` was one of the required products specified when calling [`/link/token/create`](/docs/api/link/#linktokencreate), accounts that do not support Auth, such as credit card accounts, will not appear on the Item.

##### Troubleshooting steps

Prompt the user to enter the [update mode flow](/docs/link/update-mode/), which will allow them to change the permissions they have granted to your app.

Ensure that the [`/link/token/create`](/docs/api/link/#linktokencreate) call is configured in a way that is compatible with all the desired accounts.

If accounts disappeared after the Item was linked, see the Troubleshooting guide for [`INVALID_ACCOUNT_ID`](/docs/errors/invalid-input/#invalid_account_id).

#### Link flow is failing

##### Symptom

The Link flow is not working correctly for the user -- the `onSuccess` callback may not be sent, Link may not load, or redirects may not function correctly.

##### Common causes

- Third party software, such as ad-blocking browser extensions, may be interfering with Link.
- The end user linked their account via an OAuth flow and OAuth is not configured correctly.

##### Troubleshooting steps

Ask the end user to temporarily disable ad-blockers or other third-party browser extensions that may be interfering with Link and try again.

If the Item is at an institution that uses OAuth, see [OAuth not working](/docs/link/troubleshooting/#oauth-not-working).

#### "You may need to update your app", "Couldn't connect to your institution", or "Something went wrong: an internal error occurred" error message

##### Symptom

After selecting a bank in Link, the user gets an error message saying "You may need to update your app in order to connect to Chase" (or another institution), "Couldn't connect to your institution", or "Something went wrong: an internal error occurred".

##### Common causes

- You haven't selected a use case for Link.
- The full list of OAuth prerequisites have not been completed to enable your client id for a given institution's OAuth connection.
- A backend error occurred during the Link flow.

##### Troubleshooting steps

Ensure you have selected use cases for Link, which you can do from the Dashboard, under [Link Customization -> Data Transparency](https://dashboard.plaid.com/link/data-transparency-v5).

See the troubleshooting guide for the [`UNAUTHORIZED_INSTITUTION`](/docs/errors/institution/#unauthorized_institution) error.

The "something went wrong" message indicates a 500 error. Viewing the logs (for example, via [**Dashboard>Developers>Logs**](https://dashboard.plaid.com/developers/logs)) will show you a more detailed error message and code.

#### Interactive Brokers customer prompted to enter a query id and token

##### Symptom

When linking an Interactive Brokers account, the end user is prompted to enter a "token" and/or "query id" but does not know where or how to get one.

##### Common causes

This is part of the standard Interactive Brokers Link flow. Interactive Brokers requires this step for any customer linking a third party service, such as Plaid, to their account.

###### Troubleshooting steps

See the [Interactive Brokers help center](https://www.interactivebrokers.com/lib/cstools/faq/#/content/41567214) for instructions on obtaining a query id and token.

#### OAuth not working

First, verify these settings before reviewing the more detailed OAuth troubleshooting guides below:

Ensure your redirect URI is configured in the [Plaid Dashboard](https://dashboard.plaid.com/developers/api)

For web, iOS SDK, or React Native iOS integrations, ensure the `redirect_uri` parameter is set in the [`/link/token/create`](/docs/api/link/#linktokencreate) call for all Link sessions, including Link in update mode.

For Android SDK integrations, ensure that your Android package name is registered in the [Plaid Dashboard](https://dashboard.plaid.com/developers/api) and that the `android_package_name` parameter is set in the [`/link/token/create`](/docs/api/link/#linktokencreate) call for all Link sessions, including Link in update mode.

For iOS SDK integrations, check that the redirect URI has been associated with an Apple App Association file and that your application has the associated-domains entitlement for that URI. For more details, see the [OAuth guide](/docs/link/oauth/#ios).

If none of the steps above resolve your problem, see the specific OAuth troubleshooting guides below. If you do not see your problem listed, or following the guide does not fix the problem, [contact Support](https://dashboard.plaid.com/support/new/product-and-development/developer-lifecycle/sdk).

#### OAuth institutions not appearing in Link

##### Symptom

Institutions that use OAuth, such as Capital One or Wells Fargo, are not appearing in the Link UI.

##### Common causes

- A `redirect_uri` or `android_package_name` was not specified in the call to [`/link/token/create`](/docs/api/link/#linktokencreate).
- The redirect URI or Android package name was not allowlisted on the [Plaid Dashboard](https://dashboard.plaid.com/developers/api).

##### Troubleshooting steps

Make sure the `redirect_uri` (or, for Android, `android_package_name`) is specified in the call to [`/link/token/create`](/docs/api/link/#linktokencreate).

Make sure to list your redirect URI or Android package name in the [Plaid Dashboard](https://dashboard.plaid.com/developers/api). This setting can be found under **Team settings -> API -> Allowed redirect URIs** (or **Allowed Android package names**) -> **Configure**.

#### OAuth flow shows unsupported browser warning

##### Symptom

Upon entering their institution's OAuth flow, users see a warning that their browser is not supported.

##### Common causes

- The Plaid integration is using a webview with an incorrectly configured user agent.

##### Troubleshooting steps

Use an official Plaid mobile SDK or Hosted Link rather than a webview-based integration.

Ensure that the webview uses a valid user agent matching one commonly found in the wild.

#### OAuth redirects not working

##### Symptom

Users are not being redirected back to Link after completing their institution's OAuth flow.

##### Common causes

###### Android

- An `android_package_name` was not specified in the call to [`/link/token/create`](/docs/api/link/#linktokencreate).
- An incorrect `android_package_name` was specified in the call to [`/link/token/create`](/docs/api/link/#linktokencreate).

###### iOS

- Universal Links are not correctly configured.

##### Troubleshooting steps

###### Android

Make sure the `android_package_name` is correctly specified in the call to [`/link/token/create`](/docs/api/link/#linktokencreate).

###### iOS

Ensure you are testing on a physical device and not an iOS simulator. Universal Links may not always work correctly on the simulator.

Visit your redirect URI manually and confirm that it opens your app. You can do this by either clicking on this URI in a link (for instance, by adding it to the Reminders app) or by using the command line tool to open the URL in your currently running simulator: `xcrun simctl openurl booted '<your URI>'`.

If this does successfully open your app, confirm that the `redirect_uri` is correctly set in your [`/link/token/create`](/docs/api/link/#linktokencreate) call. If does not, there is likely a problem with the `apple-app-site-association` file or the entitlements for the app; view the troubleshooting steps below.

Test the `apple-app-site-association` file by taking the domain (not the full path) of the `redirect_uri` and accessing
the corresponding location on Apple's CDN using a web browser. For example, if the `redirect_uri` is `https://app.wonderwallet.com/plaid`, then visit `https://app-site-association.cdn-apple.com/a/v1/app.wonderwallet.com`.

If the contents of the `apple-app-site-association` blob do not appear, then Apple does not have a copy of the `apple-app-site-association` file at the correct location needed to invoke Universal Links. See [Apple's Universal Links documentation](https://developer.apple.com/library/archive/documentation/General/Conceptual/AppSearch/UniversalLinks.html) for information on how to create, name, and locate this file.

Confirm that the `apple-app-site-association` contains the exact and correct path for the `redirect_uri`.

Confirm that the `apple-app-site-association` is not malformatted, using a tool such as the [ASAA validator](https://branch.io/resources/aasa-validator/). If combining iOS 12 style (`paths:`) and iOS 13 or higher style (`components:`) formats, the two should be in separate sections of the file, as in this [example app site association file](https://cdn-testing.plaid.com/.well-known/apple-app-site-association).

Ensure that the associated domains have been added as entitlements to your app. For instructions, see [Apple Universal Links documentation](https://developer.apple.com/library/archive/documentation/General/Conceptual/AppSearch/UniversalLinks.html#/apple_ref/doc/uid/TP40016308-CH12-SW2).

#### "Current version not supported" or `PUBLIC_KEY_NOT_ENABLED`

##### Symptom

- Link displays a "Current version not supported" message on launch, or you see a `PUBLIC_KEY_NOT_ENABLED` error on the backend.

##### Common causes

- The integration is using a public key to launch Link instead of a Link token. As of February 2025, public key integrations are no longer allowed.

##### Troubleshooting steps

See the [Link token migration guide](/docs/link/link-token-migration-guide/) to migrate to Link tokens.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
