---
title: "Link - OAuth guide | Plaid Docs"
source_url: "https://plaid.com/docs/link/oauth/"
scraped_at: "2026-03-07T22:05:07+00:00"
---

# OAuth Guide

#### Configure Link to connect to institutions via OAuth

Prefer to learn by watching? Video guides are available for this topic.

- [iOS Video Tutorial: OAuth section](https://youtu.be/9fgmW38usTo&t=2608s)
- [OAuth overview and guide for web](https://youtu.be/E0GwNBFVGik)
- [OAuth video guide for Android](https://youtu.be/oM7vL49I5tc)

#### Introduction to OAuth

OAuth support is required in all Plaid integrations that connect to financial institutions in the US, EU, and UK. Without OAuth support, your end users will not be able to connect accounts from institutions that require OAuth, which includes many of the largest banks in the US, and all financial institutions in Europe. OAuth setup can be skipped only if your Plaid integration is limited to products and flows that do not connect to financial institutions (such as Enrich, Identity Verification, Monitor, Beacon, Protect, and Document Income).

OAuth is an industry-standard framework for authorization. With OAuth, end users can grant third parties access to their data without sharing their credentials directly with the third party.

Typically, end users authenticate and permission data directly within Plaid Link when connecting their financial accounts to third party applications. With OAuth, however, end users temporarily leave Link to authenticate and authorize data sharing using the institution's website or mobile app instead. Afterward, they're redirected back to Link to complete the Link flow and return control to the third party application.

In an OAuth flow, the user selects their financial institution...

![In an OAuth flow, the user selects their financial institution...](/assets/img/docs/link-tour/link-2.png)

![...they are directed to its OAuth flow...](/assets/img/docs/link-tour/link-3.png)

![...they sign in at institution's site or app and authorize Plaid's access...](/assets/img/docs/link-tour/link-4.png)

![...and are redirected to Plaid to complete the flow.](/assets/img/docs/link-tour/link-5.png)

In addition, Plaid integrations with OAuth have several benefits over the traditional, non-OAuth experience in Link, such as:

- **Familiar and trustworthy experiences** With OAuth, end users authenticate via the bank's website or mobile app, a familiar experience that can help with conversion.
- **Streamlined login experiences** Some OAuth-enabled institutions (e.g., Chase) provide an "App-to-App" experience for end users if the end user has the institution's mobile app installed on their device. App-to-App can provide alternative authentication methods to end users (e.g., Touch ID or Face ID) that can help simplify and accelerate the authentication process.
- **Greater connection uptime** You can generally expect greater connection uptime with OAuth-enabled institutions, which means fewer connection errors for end users when using Plaid Link.
- **Longer-lived connections** Items at OAuth-enabled institutions generally remain connected longer. This typically results in fewer re-authentication errors (e.g., `ITEM_LOGIN_REQUIRED`).
- **Improved MFA (multi/second-factor) support** OAuth-enabled institutions can support end user accounts that may be currently unsupported due to the end user's MFA settings.

#### OAuth support and compatibility

OAuth is supported on all platforms on which Link is supported.

Plaid supports the OAuth2 protocol. For a list of the largest Plaid-supported institutions that use OAuth in the US, consult the [OAuth institutions page](https://dashboard.plaid.com/settings/compliance/us-oauth-institutions). For a full list of institutions, call [`/institutions/get`](/docs/api/institutions/#institutionsget) endpoint with your desired `country_codes` and the `oauth` option set to `true`. Note that for an institution where a migration to OAuth is in progress, some Items may use OAuth, while other Items at the same institution may not.

In general, OAuth connections are used universally by financial institutions in the UK and EU, are used by a number of financial institutions, especially larger ones, in the United States, and are not currently used by financial institutions in Canada.

#### Prerequisites for adding OAuth support

##### Request Production access from Plaid

[Full Production access](https://dashboard.plaid.com/overview/production) is a prerequisite for supporting OAuth. Plaid will contact you once your account has been enabled for Production.

In the US, OAuth requires full Production access. You can test OAuth in the Sandbox environment, using Sandbox-only institutions, without needing Production approval.

##### Complete the registration requirements

Before implementing support for OAuth institutions, be sure to [complete the registration requirements](https://dashboard.plaid.com/settings/compliance/us-oauth-institutions) in the Plaid Dashboard.

**Application display information** – This is public information that end users of your application will see when managing connections between your application and their bank accounts, including during OAuth flows. This information helps end users understand what your application is and why it is requesting access, which can improve conversion. In addition, some US institutions require your profile to be completed and will not allow apps with an empty profile to access their OAuth implementations.

**Company information** – Information about your company. This information is not shared with end users of your application and is only accessible to Plaid, members of your team, and financial institutions you register with.

If you later need to update your company information or application display information, you can do so at any time via the Dashboard. If your company's name changes, in addition to updating the Dashboard, you will also need to contact your Account Manager in order for the change to be reflected in the OAuth flows for certain institutions, such as Charles Schwab and Capital One. If you do not have an Account Manager, you can [contact support](https://dashboard.plaid.com/support/new/admin/account-administration/oauth-registration).

**Plaid Master Services Agreement (MSA)** – (US/CA only) Your latest contract with Plaid. If this is marked as incomplete, please reach out to your account manager or [contact support](https://dashboard.plaid.com/support/new/product-and-development/account-administration/oauth-registration) for an updated version.

**Plaid security questionnaire** – (US/CA only) You must complete a questionnaire about your company's risk and security practices before accessing certain bank APIs. Because it may take some time for Plaid to review, it is recommended that you submit this questionnaire as early as possible in the integration process. If your Plaid integration process is otherwise complete but your security questionnaire has not yet been approved, contact your account manager or [submit a Support ticket](https://dashboard.plaid.com/support/new/product-and-development/account-administration/oauth-registration).

###### Legal Entity Identifier (LEI)

Plaid is not currently enforcing the requirement to have an LEI and will not do so until the first 1033 compliance deadline. This deadline has been moved from April 2026 until at least July 2026 and its enforcement is currently stayed pending the finalization of a revised 1033 rule.

As part of your company information, you will be requested to optionally provide a Legal Entity Identifier, or LEI. The LEI is required under the [section 1033 rule](https://view-su2.highspot.com/viewer/a28391b5577027178f61d40b03f9c466) for any Plaid customer using OAuth connections to connect to US financial institutions. If your organization has multiple client ids, a unique LEI will be requested for each client id used in Production. Plaid is not currently enforcing the requirement to have an LEI. You can obtain an LEI from a [registration agent](https://www.gleif.org/en/about-lei/get-an-lei-find-lei-issuing-organizations).

#### Implementing OAuth support

##### Required steps

1. Wait for OAuth approval from Plaid, which can be tracked on the [OAuth institution page](https://dashboard.plaid.com/settings/compliance/us-oauth-institutions).
2. Understand [institution-specific behaviors](/docs/link/oauth/#institution-specific-behaviors) and, if necessary, update your app to support them.

##### Additional required steps for mobile implementations

1. [Create and register a redirect URI](/docs/link/oauth/#create-and-register-a-redirect-uri)
2. [Generate a Link token and configure it with your redirect URI](/docs/link/oauth/#configure-your-link-token-with-your-redirect-uri)
3. [(Mobile web and webview integrations only) Support sessions launched in embedded browsers by reinitializing Link at your redirect URI](/docs/link/oauth/#reinitializing-link)

##### Recommended, optional steps

1. [Handle Link OAuth events](/docs/link/oauth/#handling-link-events)
2. [Listen for consent expiration webhooks](/docs/link/oauth/#refreshing-item-consent)
3. [Manage consent revocation](/docs/link/oauth/#managing-consent-revocation)
4. [(US/CA only) Enable OAuth and migrate users](/docs/link/oauth/#enabling-oauth-connections-and-migrating-users)
5. [(Europe only) Enable QR Code authentication](/docs/link/oauth/#qr-code-authentication)

#### Create and register a redirect URI

After successfully completing the OAuth flow via their bank's website or app, you'll need to redirect the end user back to your application. This is accomplished with a redirect URI that you'll need to set up and configure accordingly depending on your client platform.

Constructing valid redirect URIs

Redirect URIs must use HTTPS. The only exception is on Sandbox, where, for testing purposes, redirect URIs pointing to localhost are allowed over HTTP. Custom URI schemes are not supported in any environment. Subdomain wildcards are supported using a `*` character. For example, adding `https://*.example.com/oauth.html` to the allowlist permits `https://oauth1.example.com/oauth.html`, `https://oauth2.example.com/oauth.html`, etc. Subdomain wildcards can only be used for domains that you control and are not allowed for domains on the [Public Suffix List](https://publicsuffix.org/list/). For example, `https://*.co.uk/oauth.html` is not a valid subdomain wildcard. Redirect URIs do not support hash routing, so your URI cannot contain a '#' symbol.

Note: Do not enter a wildcard (\*) when specifying a `redirect_uri` in the call to [`/link/token/create`](/docs/api/link/#linktokencreate). Wildcards are reserved for the allowlist on the dashboard only.

##### Desktop web, mobile web, React, or Webview

For desktop web, mobile web, or React, the redirect URI is typically the address of a blank web page you'll need to create and host. This web page will be used to allow the end user to resume and complete the Link flow after completing the OAuth flow on their bank's website or app. `https://example.com/oauth-page.html` is an example of a typical redirect URI. After creating your redirect URI, add it to the [Allowed redirect URIs](https://dashboard.plaid.com/developers/api).

Redirect URIs and desktop & mobile web

OAuth flows will function properly on web even if you don't set up a redirect URI. Desktop web and mobile web integrations will always try to open the OAuth bank's website in a new pop-up window if possible (or in a new tab on mobile web), regardless of whether a `redirect_uri` is provided. However, not providing a redirect URI will prevent mobile web users from using your integration through a webview browser (a browser launched via Mail, Facebook, Google Maps, etc.) because those browsers often do not support pop-ups. To provide the best experience for end users on mobile web, always specify a redirect URI and reinitialize Link.  
Setting a `redirect_uri` is still required for Link web SDK integrations within a mobile application (e.g within a webview) because those integrations still use the redirect OAuth flow.

##### iOS SDK, React Native (iOS)

For iOS SDK or React Native (iOS), the redirect URI is typically the address of a blank web page you'll need to create and host. You'll need to configure an [Apple App Association File](https://developer.apple.com/documentation/security/password_autofill/setting_up_an_app_s_associated_domains) to associate your redirect URI with your application. To enable [App-to-App authentication flows](/docs/link/oauth/#app-to-app-authentication), you'll need to register the redirect URI as an app link. Custom URI schemes are not supported; a proper universal link **must** be used.

##### Android SDK, React Native (Android)

Register your Android package by adding the Android package name(s) to the [Allowed Android package names](https://dashboard.plaid.com/developers/api) list. Plaid will automatically create your redirect URI and its contents based on your package name. When specifying a redirect URI in the following steps, you will use [`android_package_name`](https://plaid.com/docs/api/link/#link-token-create-request-android-package-name).

#### Configure your Link token with your redirect URI

You'll need to specify your redirect URI via the `redirect_uri` field when generating a Link token with [`/link/token/create`](/docs/api/link/#linktokencreate) (on Android, use the `android_package_name` parameter to provide your Android package name instead). Use this Link token to initialize Link.

If you're using Link in [update mode](/docs/link/update-mode/), ensure you specify your redirect URI via the `redirect_uri` field (on Android, use the `android_package_name` parameter to provide your package name instead).

Do not use query parameters when specifying the `redirect_uri`. Make sure to specify the `user.client_user_id`.

Generating a Link token

```
const request = {
  user: {
    // This should correspond to a unique id for the current user.
    client_user_id: 'user-id',
  },
  client_name: 'Plaid Test App',
  products: [Products.Transactions],
  country_codes: [CountryCode.Us],
  language: 'en',
  webhook: 'https://sample-web-hook.com',
  redirect_uri: 'https://example.com/callback',
};

try {
  const createTokenResponse = await client.linkTokenCreate(request);
  const linkToken = createTokenResponse.data.link_token;
} catch (error) {
  // handle error
}
```

##### Using OAuth within an iFrame

Launching Link from within an iFrame is not recommended. Link conversion for OAuth institutions is typically up to 15 percentage points higher when using Plaid's SDKs than when using iFrames. If Link is launched from within an iFrame, you'll be unable to maintain user state. Page rendering, sizing, and data exchange may also be suboptimal.

#### Reinitializing Link

After completing the OAuth flow, the end user will be redirected to your redirect URI (e.g., `https://example.com/oauth-page.html`). This is where they'll resume and complete the Link flow and return to your application. To do this, you'll need to reinitialize Link at your redirect URI.

Depending on your client platform, Link may require additional configuration to work with OAuth. Detailed instructions for each platform are provided below.

| Client platform | Link reinitialization required? |
| --- | --- |
| [Desktop web](/docs/link/oauth/#desktop-web-mobile-web-or-react) | No |
| [Mobile web](/docs/link/oauth/#desktop-web-mobile-web-or-react) | Not required, but recommended in order to maximize Link conversion |
| [Webview](/docs/link/oauth/#webview) | Yes |
| [iOS SDK](/docs/link/oauth/#ios) | No |
| [React Native (iOS)](/docs/link/oauth/#react-native-on-ios) | No |
| Android SDK (version 3.2.3 or later required) | No, but [app package registration required](/docs/link/oauth/#android-sdk-react-native-android) |
| React Native (Android) | No, but [app package registration required](/docs/link/oauth/#android-sdk-react-native-android) |
| [Hosted Link](/docs/link/oauth/#hosted-link) | No |

##### Desktop web, mobile web, or React

A reference implementation for OAuth in React can be found in the [Plaid React GitHub](https://github.com/plaid/react-plaid-link#oauth--opening-link-without-a-button-click). If you are looking for a demonstration of a real-life app that incorporates the implementation of OAuth in React see [Plaid Pattern](https://github.com/plaid/pattern), a Node-based example app.

Desktop and mobile web sessions do not require Link reinitialization by default.

However, not supporting Link reinitialization will prevent mobile web users from using your integration through a webview (an embedded browser launched via Mail, Facebook, Google Maps, etc.). For these sessions, you'll need to launch Link twice, once before the OAuth redirect (i.e., the first Link initialization) and once after the OAuth redirect (i.e., Link reinitialization). The Link reinitialization should occur at your redirect URI.

When reinitializing Link, configure it using the same Link token you used when [initializing Link the first time](/docs/link/oauth/#configure-your-link-token-with-your-redirect-uri). It is up to you to determine the best way to provide the correct `link_token` upon redirect. As an example, the code sample below demonstrates the use of a browser's local storage to retrieve the Link token from the first Link initialization.

Select group for content switcher

/oauth-page: Reinitializing Link with receivedRedirectUri

```
import React, { useEffect } from 'react';
import { usePlaidLink } from 'react-plaid-link';

const OAuthLink = () => {
  // The Link token from the first Link initialization
  const linkToken = localStorage.getItem('link_token');

  const onSuccess = React.useCallback((public_token: string) => {
    // send public_token to server, retrieve access_token and item_id
    // return to "https://example.com" upon completion
  });

  const onExit = (err, metadata) => {
    // handle error...
  };

  const config: Parameters<typeof usePlaidLink>[0] = {
    token: linkToken!,
    // pass in the received redirect URI, which contains an OAuth state ID parameter that is required to
    // re-initialize Link
    receivedRedirectUri: window.location.href,
    onSuccess,
    onExit,
  };

  const { open, ready, error } = usePlaidLink(config);

  // automatically reinitialize Link
  useEffect(() => {
    if (ready) {
      open();
    }
  }, [ready, open]);

  return <></>;
};

export default OAuthLink;
```

oauth-page.html: Reinitialize Link with receivedRedirectUri

```
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.2.3/jquery.min.js"></script>
    <script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
    <script>
      (function ($) {
        var linkToken = localStorage.getItem("link_token");
        var handler = Plaid.create({
          token: linkToken,
          // pass in the received redirect URI, which contains an OAuth state ID parameter that is required to
          // re-initialize Link
          receivedRedirectUri: window.location.href,
          onSuccess: function (public_token) {
            $.post(
              "/api/set_access_token",
              { public_token: public_token },
              function (data) {
                location.href = "https://example.com";
              }
            );
          },
        });
        handler.open();
      })(jQuery);
    </script>
```

In addition, when reinitializing Link, you should configure it with the `receivedRedirectUri` field and pass in the full received redirect URI, as demonstrated in the code sample. The received redirect URI is your redirect URI appended with an OAuth state ID parameter. The OAuth state ID parameter will allow you to persist user state when reinitializing Link, allowing the end user to resume the Link flow where they left off. No extra configuration or setup is needed to generate the received redirect URI. The received redirect URI is programmatically generated for you by Plaid after the end user authenticates on their bank's website or mobile app. You can retrieve it using `window.location.href`.

The received redirect URI must not contain any extra query parameters or fragments other than what is provided upon redirect. The standard Link callback `onSuccess` will be triggered as usual once the user completes the Link flow.

An example received redirect URI

```
https://example.com/oauth-page.html?oauth_state_id=9d5feadd-a873-43eb-97ba-422f35ce849b
```

###### Optional methods for retrieving the initial Link token

If Link is reinitialized in the *same* browser session as the first Link initialization, you can store the Link token in a cookie or local storage in the browser for easy access when reinitializing Link. For example, the Plaid Quickstart uses `localStorage.setItem` to store the token.

If Link is reinitialized in a *different* browser session than the first Link initialization, you can store a mapping of the Link token associated with the user (server-side). Upon opening the second browser session, authenticate the user, fetch the corresponding Link token from the server, and use it to reinitialize Link.

##### Webview

All webview-based integrations need to extend the webview handler for redirects in order to support Chase OAuth. This can be accomplished with code samples for [iOS](https://github.com/plaid/plaid-link-examples/blob/master/webviews/wkwebview/wkwebview/LinkViewController.swift#L56-L72) and [Android](https://github.com/plaid/plaid-link-examples/blob/master/webviews/android/LinkWebview/app/src/main/java/com/example/linkwebview/MainActivity.kt#L89-L156) For more details, see [Extending webview instances to support certain institutions](/docs/link/oauth/#extending-webview-instances-to-support-certain-institutions).

For webview, you'll need to launch Link twice, once before the OAuth redirect (i.e., the first Link initialization) and once after the OAuth redirect (i.e., Link reinitialization). The Link reinitialization should occur at your redirect URI.

For the initial Link instance, first generate a Link token as described in [Configure your Link token with your redirect URI](/docs/link/oauth/#configure-your-link-token-with-your-redirect-uri), then set Link's `token` parameter to this Link token. After the end user successfully completes the OAuth flow via their bank's website or app, they'll be redirected to your redirect URI, where you'll reinitialize Link.

Initializing Link (first Link initialization)

```
https://cdn.plaid.com/link/v2/stable/link.html?isWebview=true
&token=GENERATED_LINK_TOKEN
```

When reinitializing Link, use the same Link token you generated when you first initialized Link. It is up to you to determine the best way to provide the correct `link_token` upon redirect.

In addition, when reinitializing Link, you should configure it with the `receivedRedirectUri` field. The received redirect URI is your redirect URI appended with an OAuth state ID parameter. The OAuth state ID parameter will allow you to persist user state when reinitializing Link, allowing the end user to resume the Link flow where they left off. No extra configuration or setup is needed to generate the received redirect URI. The received redirect URI is programmatically generated after the end user authenticates on their bank's website or mobile app. The received redirect URI must not contain any extra query parameters or fragments other than what is provided upon redirect. Note that any unsafe ASCII characters in the `receivedRedirectUri` in the webview query string must be URL-encoded; for improved readability, the example below is shown prior to URL encoding.

Reinitializing Link

```
https://cdn.plaid.com/link/v2/stable/link.html?isWebview=true
&token=SAME_GENERATED_LINK_TOKEN&receivedRedirectUri=https://example.com/oauth-page?oauth_state_id=9d5feadd-a873-43eb-97ba-422f35ce849b
```

###### Extending webview instances to support certain institutions

Some institutions require further modifications to work with mobile webviews. This applies to USAA on Android and Chase App-to-App on Android, as well as all Chase webview integrations. For these institutions, you'll need to extend your Webview instance to override the handler for redirects. [An example function for Android](https://github.com/plaid/plaid-link-examples/blob/master/webviews/android/LinkWebview/app/src/main/java/com/example/linkwebview/MainActivity.kt#L89-L156) and [an example function for iOS](https://github.com/plaid/plaid-link-examples/blob/master/webviews/wkwebview/wkwebview/LinkViewController.swift#L56-L72) can be found within Plaid's Link examples on GitHub and can be copied for use with your app.

On Android, you will also need to support Android App links to have a valid working `redirect_uri` to provide to [`/link/token/create`](/docs/api/link/#linktokencreate). This requires creating a `./well_known/assetlinks.json` and an `IntentFilter` on the Android app.

##### iOS

For the initial Link instance, first generate a Link token as described in [Configure your Link token with your redirect URI](/docs/link/oauth/#configure-your-link-token-with-your-redirect-uri). Your `redirect_uri` must also be added to the Plaid Dashboard and should be configured as a universal link (not a custom URI) using an Apple App Association File.

Initializing Link

```
// With custom configuration
let linkToken = "<#GENERATED_LINK_TOKEN#>"
let onSuccess: (LinkSuccess) -> Void = { (success) in
  // Read success.publicToken here
  // Log/handle success.metadata here
}
let linkConfiguration = LinkTokenConfiguration(linkToken: linkToken, onSuccess: onSuccess)
let handlerResult = Plaid.create(linkConfiguration)

switch handlerResult {
case .success(let handler):
  self.handler = handler
  handler.open(presentUsing: .viewController(self))
case .failure(let error):
  // Log and handle the error here.
}
```

OAuth on iOS devices can occur fully within the integrating application (In-App OAuth), or it can include a transition from your application to the bank's app ([App-to-App OAuth](/docs/link/oauth/#app-to-app-authentication)). App-to-App OAuth is initiated by the bank itself and is not controlled by the iOS SDK. In order to ensure your users can return to your application, you must support App-to-App OAuth.

##### App-to-App OAuth requirements

During App-to-App OAuth, the end user is directed from your application to the bank's app to authenticate. To return the end user back to your application after they authenticate, your redirect URI must be a universal link. Once the user returns to your app, UIKit will invoke a method within your application and provide the redirect URI that triggered this return to the app.

App-to-App behavior can be [tested in the Sandbox environment](/docs/link/oauth/#app-to-app-authentication). If App-to-App does not function as intended, validate that the redirect URI used to configure Link is a valid universal link and that your application has the associated-domains entitlement for that URI.

##### React Native on iOS

An example React Native client implementation for OAuth on iOS can be found in the [Tiny Quickstart](https://github.com/plaid/tiny-quickstart/tree/main/react_native/TinyQuickstartReactNative).

React Native on iOS uses universal links for OAuth. You will need to [create and register a redirect URI](/docs/link/oauth/#create-and-register-a-redirect-uri) and [configure your Link token with your redirect URI](/docs/link/oauth/#configure-your-link-token-with-your-redirect-uri) in order for OAuth to work correctly.

##### Android SDK and Android on React Native

Example code for implementing OAuth on Android can be found on GitHub in the [Android SDK](https://github.com/plaid/plaid-link-android). An example React Native client implementation for OAuth on Android can be found in the [Tiny Quickstart](https://github.com/plaid/tiny-quickstart/tree/main/react_native/TinyQuickstartReactNative).

When using the Android SDK or Android on React Native, you must generate a Link token by calling [`/link/token/create`](/docs/api/link/#linktokencreate) and passing in an `android_package_name`. Your package name (e.g., `com.example.testapp`) must also be added to the Plaid Dashboard. Do not pass a `redirect_uri` into the [`/link/token/create`](/docs/api/link/#linktokencreate) call. Proceed to initialize Link with the generated token.

Passing in an Android package name

```
String clientUserId = "user-id";

LinkTokenCreateRequestUser user = new LinkTokenCreateRequestUser()
  .clientUserId(clientUserId)
  .legalName("legal name")
  .phoneNumber("4155558888")
  .emailAddress("email@address.com");

LinkTokenCreateRequest request = new LinkTokenCreateRequest()
  .user(user)
  .clientName("Plaid Test App")
  .products(Arrays.asList(Products.AUTH))
  .countryCodes(Arrays.asList(CountryCode.US))
  .language("en")
  .webhook("https://example.com/webhook")
  .linkCustomizationName("default")
  .androidPackageName("com.plaid.example")

Response<LinkTokenCreateResponse> response = client()
  .linkTokenCreate(request)
  .execute();

String linkToken = response.body().getLinkToken();
```

##### Hosted Link

Hosted Link can use universal links for flows in which the Hosted Link URL is used in a secure context webview within a mobile application. In this scenario, the universal link is used to redirect the user back to your app after the [app-to-app](/docs/link/oauth/#app-to-app-authentication) flow is complete.

To do so, you will need to [create and register a redirect URI](/docs/link/oauth/#create-and-register-a-redirect-uri) for the universal link and [configure your Link token with your redirect URI](/docs/link/oauth/#configure-your-link-token-with-your-redirect-uri). For more details on configuring universal links with Hosted Link, see [App-to-app authentication](/docs/link/oauth/#app-to-app-authentication).

For details on configuring your redirect URI for Hosted Link on mobile, see [Integrating Hosted Link in Native Mobile applications](/docs/link/hosted-link/#integrating-hosted-link-in-native-mobile-applications).

#### Testing OAuth

You can test OAuth in Sandbox even if Plaid has not yet enabled OAuth flows for your account. To test out the OAuth flow in the Sandbox environment, you can select any bank that supports OAuth (this includes institutions like Chase, U.S. Bank, Wells Fargo, and others) or use the Platypus OAuth Bank (`ins_127287`). For Europe, you can select Flexible Platypus Open Banking (`ins_117181`).

These institutions will direct you to a Plaid sample OAuth flow that is similar to what you would see for a bank’s OAuth flow. When prompted, you can enter anything as credentials (including leaving the input fields blank) to proceed through the sample OAuth flow. Note that institution-specific OAuth flows cannot be tested in Sandbox; OAuth panes for Platypus institutions will be shown instead.

The Platypus OAuth flow includes two optional "consent" checkboxes that reflect what a user might need to select in order to opt in to Auth and Identity. However, these checkboxes are not functional -- your application will receive the proper Auth and Identity data, even if you don't select those boxes during testing.

To ensure your OAuth integration works across all platforms, test it in the following scenarios before deployment:

- On each client platform that is available to your users (e.g. desktop, iOS app, iOS mobile web, Android app, Android mobile web)
- With the OAuth institution app installed, for institutions that support [app-to-app authentication](/docs/link/oauth/#app-to-app-authentication) (e.g. Chase). To test app-to-app in Sandbox, use First Platypus Bank as the institution; see [App-to-App authentication](/docs/link/oauth/#app-to-app-authentication) for more details.
- Without the OAuth institution app installed
- In [update mode](/docs/link/update-mode/), by using [`/sandbox/item/reset_login`](/docs/api/sandbox/#sandboxitemreset_login) in the Sandbox environment

All environments, including Sandbox, use the pop-up flow on desktop and mobile web, unless the page is accessed through a mobile webview. To test the redirect flow, you can use your browser's developer tools to simulate running your application in a webview browser, such as Chrome WebView. On Chrome, for example, select the "Toggle Device Toolbar" option from within Chrome's Developer Tools and create a new virtual device configuration with a WebView user agent.

Example user agent

```
Mozilla/5.0 (Linux; Android 5.1.1; Nexus 5 Build/LMY48B; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/43.0.2357.65 Mobile Safari/537.36
```

#### Troubleshooting common OAuth problems

For guides to troubleshooting common OAuth issues, see [Link troubleshooting](/docs/link/troubleshooting/#oauth-not-working).

#### Enabling OAuth connections and migrating users

When an institution is newly enabled for OAuth connections, your integration will automatically convert to using OAuth connections for a given institution on the date that Plaid has indicated to you. You may also have the option to migrate to OAuth earlier; if this option is available, a button to enable OAuth for the institution will appear on the [OAuth institution page](https://dashboard.plaid.com/settings/compliance/us-oauth-institutions).

For some institutions, migrating to OAuth will cause disruptions to existing Items. In this case, existing Items will be moved into the `ITEM_LOGIN_REQUIRED` state. This will happen gradually over the migration period (by default, 90 days, but the duration may vary depending on the institution), starting on or shortly after the OAuth enablement date. Completing the [update mode](/docs/link/update-mode/) flow for these Items will convert them to use OAuth connections. To avoid any disruption in connectivity, you can also prompt your users to complete the update mode flow as soon as you have been enabled for OAuth with their institution.

To view which institutions are being migrated to OAuth connections, use the [migrations pane](https://dashboard.plaid.com/activity/status/migrations) within the Dashboard status page, or look up a specific institution in the status page to see its migration details, including the migration timeline and whether existing Items will be disrupted. For more details, see [Institution migration status](/docs/account/activity/#institution-migration-status).

#### Institution-specific behaviors

Some institutions have unique behaviors when used with OAuth connections. Note that these behaviors are standard for any connection using these institutions' APIs and not specific to their integration with Plaid.

The behaviors listed below are the ones most likely to require changes to your application's business logic; for a more exhaustive list of institution-specific OAuth details, including screenshots of the OAuth flows and summaries of data availability for several major institutions, see the [Plaid guide to institution-specific OAuth experiences](https://plaid.com/documents/oauth-institution-ux.pdf).

##### Bank of America

In 2026, Bank of America Items began an API migration. From February through March 2026, the new API is being gradually rolled out for new Items. Throughout the period of March 2026 - October 2026, existing Items will gradually be disconnected from the old Bank of America API and must be connected to the new Bank of America API by going through update mode. To minimize disruptions, listen for the `PENDING_DISCONNECT` webhook. Upon receiving the webhook, prompt the user to go through the Update Mode flow for the Item, which will migrate the Item to the new API. One week after the `PENDING_DISCONNECT` webhook was fired for a given Item, if the Item has not yet gone through update mode, it will be disconnected from the old API and will enter the `ITEM_LOGIN_REQUIRED` error state. Sending the Item through update mode will move it to the new API and restore it to a healthy state.

Once an Item is on the new Bank of America API, it will have a 12-month consent expiration policy applied to it. In the new API, Bank of America will also have its own hosted account select screen.

Note that some end users may see connection expiration dates of October 2026 in the Bank of America portal. These expiration dates represent expirations for Items that have not yet been migrated to use the new Bank of America API. Once the Item has been migrated to the new API, the expiration date will be updated.

##### Chase

You must complete the Security Questionnaire before gaining access to Chase in Production.

When used with Auth or Transfer, Chase will return tokenized routing and account numbers, which can be used for ACH transactions but are not the user's actual account and routing numbers. For important details on how to avoid user confusion and ACH returns, see [Tokenized account numbers](/docs/auth/#tokenized-account-numbers).

[Update mode](https://plaid.com/docs/link/update-mode/#using-update-mode-to-request-new-accounts) cannot be used to remove accounts or permissions from a Chase Item. If your end user wishes to revoke or limit Plaid's access to a Chase account, they must do so via the Chase Security Center within the Chase online banking portal. This occurs even across Items; if a Chase Item is deleted and a new Item is created for your app using the same credentials, the previously selected permissions will persist to the new Item.

Existing Chase OAuth Items will be invalidated when a new public token is created using the same credentials, if the two Items do not have exactly the same set of accounts associated with them, or if either Item is used with a [Plaid Check](/docs/check/) product. For more details, see [Preventing duplicate Items](/docs/link/duplicate-items/#preventing-duplicate-items).

##### Capital One

When calling [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) for a Capital One non-depository account, such as a credit card or loan account, you will need to specify how fresh you require the balance data to be. If balance data meeting your requirements is not available, the call will fail and you will not be billed. For more details, see the [API Reference](/docs/api/products/signal/#accounts-balance-get-request-options-min-last-updated-datetime). For similar reasons, [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) will result in a `PRODUCTS_NOT_SUPPORTED` error if used on a Capital One Item that includes only non-depository accounts.

Capital One Items require consent to be refreshed after one year.

Capital One (like several other institutions) does not provide pending transaction data.

Capital One does not allow end users to link credit cards whose payments are past due.

##### Charles Schwab

While most OAuth institutions grant OAuth access immediately upon full Production approval, Schwab has an additional waiting period and it may take up to five weeks from Production approval until Schwab Production access has been granted. Pay-as-you-go customers will need to explicitly request access to Schwab by filing a ticket in the Dashboard.

Existing Charles Schwab OAuth Items will be invalidated when a new public token is created using the same credentials, if the two Items do not have exactly the same set of accounts associated with them, or if either Item is used with a [Plaid Check](/docs/check/) product. For more details, see [Preventing duplicate Items](/docs/link/duplicate-items/#preventing-duplicate-items).

##### PNC

You must complete the Security Questionnaire before gaining access to PNC in Production.

When used with Auth or Transfer, PNC will return tokenized routing and account numbers (TANs), which can be used for ACH transactions but are not the user's actual account and routing numbers. For important details on how to avoid user confusion and ACH returns, see [Tokenized account numbers](/docs/auth/#tokenized-account-numbers). PNC TANs do not currently expire, even if the Item expires; for the latest information on this policy, see [PNC TAN expiration](/docs/auth/#pnc-tan-expiration).

PNC Items require consent refresh once per year.

Existing PNC OAuth Items will be invalidated when a new public token is created using the same credentials, if the two Items do not have exactly the same set of accounts associated with them, or if either Item is used with a [Plaid Check](/docs/check/) product. For more details, see [Preventing duplicate Items](/docs/link/duplicate-items/#preventing-duplicate-items).

##### US Bank

When used with Auth or Transfer, US Bank will return tokenized routing and account numbers, which can be used for ACH transactions but are not the user's actual account and routing numbers. For important details on how to avoid user confusion and ACH returns, see [Tokenized account numbers](/docs/auth/#tokenized-account-numbers).

#### Handling Link events

OAuth flows have a different sequence of [Link events](https://plaid.com/docs/link/web/#onevent) than non-OAuth flows. If you are using Link events to measure conversion metrics for completing the Link process, you may need to handle these events differently when using OAuth.

In addition, the flow itself may be different if you are initiating OAuth with a redirect URI or displaying the OAuth screen in a separate pop-up window.

The events fired for a non-OAuth flow might look something like this:

Typical non-OAuth event flow

```
OPEN (view_name = CONSENT)
TRANSITION_VIEW (view_name = SELECT_INSTITUTION)
SELECT_INSTITUTION
TRANSITION_VIEW (view_name = CREDENTIAL)
SUBMIT_CREDENTIALS
TRANSITION_VIEW (view_name = LOADING)
TRANSITION_VIEW (view_name = MFA, mfa_type = code)
SUBMIT_MFA (mfa_type = code)
TRANSITION_VIEW (view_name = LOADING)
TRANSITION_VIEW (view_name = CONNECTED)
HANDOFF
onSuccess
```

The events fired for a typical OAuth flow may look more like the following:

Typical OAuth event flow when using a redirect\_uri

```
OPEN (view_name = CONSENT)
TRANSITION_VIEW (view_name = SELECT_INSTITUTION)
SELECT_INSTITUTION
TRANSITION_VIEW (view_name = OAUTH)
OPEN_OAUTH
...
(The user completes the OAuth flow at their bank)
...
TRANSITION_VIEW (view_name = CONNECTED)
HANDOFF
onSuccess
```

Link does not issue the `SUBMIT_CREDENTIALS` event when a user authenticates with an institution that requires OAuth. Link issues the `OPEN_OAUTH` event when a user chooses to be redirected to the institution’s OAuth portal. It is recommended to track this event instead of `SUBMIT_CREDENTIALS`.

In most situations, if Plaid encounters an error from the bank's OAuth flow, or if the user chooses to not grant access via their OAuth login attempt, the user will return to your application and Link will fire an `EXIT` event with a `requires_oauth` exit status. You may also see an `ERROR` event, depending on the type of error that Plaid encountered.

Link event flow after returning with an error

```
SELECT_INSTITUTION
OPEN (view_name = null)
ERROR (error_code = INTERNAL_SERVER_ERROR)
TRANSITION_VIEW (view_name = ERROR)
TRANSITION_VIEW (view_name = EXIT)
EXIT (exit_status = requires_oauth)
```

If the user closes the bank's OAuth window without completing the OAuth flow, Link will fire a `CLOSE_OAUTH` event. This only happens in situations where the OAuth flow appears in a pop-up window.

If the OAuth flow times out while waiting for the user to sign in, Link will fire a `FAIL_OAUTH` event. This only happens in situations where the OAuth flow appears in a pop-up window.

Once you receive the `onSuccess` callback from an OAuth flow, the integration steps going forward are the same as for non-OAuth flows.

#### Refreshing Item consent

When using OAuth, consent may expire after a certain amount of time and need to be refreshed. This is common in Europe, where consent typically expires after 180 days. In the US, the following institutions require periodic OAuth consent refresh. Consent refresh at these institutions is required every 12 months, unless noted otherwise.

- American Express
- Bank of America (new API only, [see Bank of America for more details](/docs/link/oauth/#bank-of-america))
- Brex (3 months)
- Capital One
- Citibank
- Navy Federal Credit Union
- PNC
- TD Bank
- USAA (18 months)

To determine when consent expires, call [`/item/get`](/docs/api/items/#itemget) and note the `consent_expiration_time` field. Plaid will also send a [`PENDING_DISCONNECT`](/docs/api/items/#pending_disconnect) webhook (for US/CA Items) or a [`PENDING_EXPIRATION`](/docs/api/items/#pending_expiration) webhook (for UK/EU Items) one week before a user's consent is set to expire.

To refresh consent for an Item, send it through [update mode](/docs/link/update-mode/).

If consent is not refreshed before the Item expires, the Item will enter the `ITEM_LOGIN_REQUIRED` error state. Sending the Item through update mode will resolve the error.

For a real-life example of handling the `PENDING_EXPIRATION` webhook and update mode, see [handleItemWebhook.js](https://github.com/plaid/pattern/blob/master/server/webhookHandlers/handleItemWebhook.js#L69-L86), [linkTokens.js](https://github.com/plaid/pattern/blob/master/server/routes/linkTokens.js#L30-L35) and [launchLink.tsx](https://github.com/plaid/pattern/blob/master/client/src/components/LaunchLink.tsx#L42). These files illustrate the code for handling of webhooks and update mode for Plaid Link for React for the Node-based [Plaid Pattern](https://github.com/plaid/pattern) sample app.

#### Managing consent revocation

Many institutions that support OAuth provide a means for the end user to revoke consent via their website. If an end user revokes consent, the Item will enter an `ITEM_LOGIN_REQUIRED` state after approximately 24-48 hours.

A user may also revoke consent for a single account, without revoking consent for the entire Item. Accounts in this situation are treated the same as accounts that have been closed: no webhook will be fired, you will stop receiving transactions associated with the account, and Plaid will stop returning the account in API responses. Access can be re-authorized via the [update mode](/docs/link/update-mode/) flow.

When a user revokes access to an Item or an account, you should delete the associated data; the 1033 rule requires third parties to no longer use or retain covered data upon revocation unless use or retention of that covered data remains reasonably necessary to provide the consumer's requested product or service.

When using Link in update mode (for either case described in this section), be sure to specify your redirect URI via the `redirect_uri` field as described in [Configure your Link token with your redirect URI](/docs/link/oauth/#configure-your-link-token-with-your-redirect-uri).

#### Scoped consent

OAuth can provide the ability for end users to configure granular permissions on their Items. For example, an end user may allow access to a checking account but not a credit card account behind the same login, or may allow an institution to share only certain account information, such as identity data but not transaction history.

Before handing the user off to the institution's OAuth flow, Plaid will provide guidance within Link recommending which permissions the user needs to grant based on which products you have initialized Link with. In Production, this guidance will be tailored to use the same wording that the institution's OAuth flow uses. This guidance will appear both on initial link and during the [update mode flow](/docs/link/update-mode/).

If an end user chooses not to share data that is required by your Link token's `products` or `required_if_supported_products` configuration, or does not share access to any accounts, Link will show an error, and they will be prompted to restart the Link flow.

Note that if your app calls [`/link/token/create`](/docs/api/link/#linktokencreate) using an `account_filter` parameter to limit the account types that can be used with Link, the filter will only be applied after the OAuth flow has been completed and will not affect the permission selection interface within the OAuth flow.

If you do not have the product-specific OAuth permissions required to use a specific endpoint with an Item, you will receive an [`ACCESS_NOT_GRANTED`](/docs/errors/item/#access_not_granted) error. If you are missing permissions for an account on the Item, you may receive a product-specific error indicating that a compatible account could not be found, such as [`NO_AUTH_ACCOUNTS`](/docs/errors/item/#no_auth_accounts-or-no-depository-accounts), [`NO_INVESTMENT_ACCOUNTS`](/docs/errors/item/#no_investment_accounts), or [`NO_LIABILITY_ACCOUNTS`](/docs/errors/item/#no_liability_accounts).

If your app later needs to request access to a product or account that was not originally granted for that Item during Link, you can send the user to the [update mode](/docs/link/update-mode/) flow to authorize additional permissions. When using Link in update mode, be sure to specify your redirect URI via the `redirect_uri` field as described in [Configure your Link token with your redirect URI](/docs/link/oauth/#configure-your-link-token-with-your-redirect-uri).

#### App-to-App authentication

Some banks (e.g. Chase) support an App-to-App experience if the user is authenticating on their mobile device and has the bank's app installed. Instead of logging in via the bank's site, the bank's app will be launched instead, from which the user will be able to log in (including via TouchID or Face ID) before being redirected back to your app. Support for App-to-App should be automatic once you have implemented support for OAuth on mobile with any of Plaid's mobile SDKs. Note that on iOS, this requires configuring an Apple App Association file to associate your redirect URI with your app, as described under [Create and register a redirect URI](/docs/link/oauth/#create-and-register-a-redirect-uri). If using webviews, App-to-App support is not automatic; for an App-to-App experience, it is strongly recommended to use a Plaid mobile SDK or Hosted Link instead. If you are using Hosted Link, ensure you are passing in a Universal Link as the `redirect_uri` parameter when creating a Link token using [`/link/token/create`](/docs/api/link/#linktokencreate).

The full App-to-App flow can be tested in the Sandbox environment. You can test that your universal link works correctly using the First Platypus Bank - OAuth App2App (`ins_132241`) test institution. The bank's 'app' for this test institution will be a mobile web browser, but the universal link will function as expected if configured correctly, switching back to your app on handoff. On iOS, you must use the Plaid Link SDK version 5.1.0 or later to test App-to-App in Sandbox.

Chase and Chime are currently the only US banks that support App-to-App authentication on Plaid.

#### QR code authentication

For many European institutions, Plaid supports the ability for an end user to authenticate via their bank's mobile app – even if the user's journey begins in a desktop-based web session – in order to optimize for conversion. After the user selects an institution, they will be presented with the choice to scan a QR code and authenticate in the bank’s mobile app or to continue on desktop. When the user scans the QR code, they will be redirected to the bank’s app (or website, if the user does not have the app installed). After the user completes the OAuth flow, they will be redirected to a Plaid-owned page instructing them to return to their desktop to complete the flow.

To enable QR authentication, contact your Plaid account manager or [file a Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access). No changes to your Link OAuth implementation are required to enable this flow.

![Payment screen showing amount £21.23 for goods/services with 'Pay with bank' button on a Plaid OAuth integration page.](/assets/img/docs/qr-code-authentication.gif)

To test out the QR code flow in the Sandbox environment, you can use Flexible Platypus Open Banking (`ins_117181`). When you launch Link with this institution selected in Sandbox, the QR code authentication flow will be triggered. The Sandbox institution does not direct you to a real bank's mobile app, but allows you to grant, deny, or simulate errors from the placeholder OAuth page instead.

##### Supported institutions for QR code authentication

Below is a partial list of some of the institutions that support QR code authentication.

| Institution name | Institution ID |
| --- | --- |
| Bank of Scotland - Personal | `ins_118274` |
| Bank of Scotland - Business | `ins_118276` |
| Barclays (UK) - Mobile Banking: Business | `ins_118512` |
| Barclays (UK) - Mobile Banking: Personal | `ins_118511` |
| Barclays (UK) - Mobile Banking: Wealth Management | `ins_118513` |
| First Direct | `ins_81` |
| Halifax | `ins_117246` |
| HSBC (UK) - Business | `ins_118277` |
| HSBC (UK) - Personal | `ins_55` |
| Lloyds Bank - Business and Commercial | `ins_118275` |
| Lloyds Bank - Personal | `ins_61` |
| Monzo | `ins_117243` |
| Nationwide Building Society | `ins_60` |
| NatWest - Current Accounts | `ins_115643` |
| Revolut | `ins_63` |
| Royal Bank of Scotland - Current Accounts | `ins_115642` |
| Santander (UK) - Personal and Business | `ins_62` |
| Starling | `ins_117520` |
| Tesco (UK) | `ins_118393` |
| TSB | `ins_86` |
| Ulster Bank (UK) | `ins_117734` |

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
