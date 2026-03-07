---
title: "Link - Hosted Link for embedded clients | Plaid Docs"
source_url: "https://plaid.com/docs/link/hosted-link/embedded-clients/"
scraped_at: "2026-03-07T22:05:04+00:00"
---

# Hosted Link for embedded clients

#### Hosted Link integration for PSPs and embedded clients

#### Overview

Some customers may not have apps themselves, but instead are embedded within other apps. For example, a payment service provider (PSP) may build a checkout flow which provides several payment options, one of which is pay by bank powered by Plaid. This checkout flow is then used by the PSP's customers within their apps. For these embedded clients, Hosted Link is the recommended approach. This guide will cover implementing [Hosted Link](/docs/link/hosted-link/) specifically for the embedded client use case, covering integration aspects that are unique to this scenario.

While not all customers who use the embedded flow are necessarily PSPs, for the sake of readability, this guide will refer to the client that is embedded within a third party app as the PSP, and the app that they are embedded within as the merchant.

Hosted Link has the following benefits for PSPs:

- **Simplicity:** Hosted Link is the easiest way to integrate with Plaid when Plaid's native mobile SDKs cannot be used.
- **Continuity:** With Hosted Link, Plaid manages the redirects between Plaid Link and the bank’s OAuth site. Configuring redirects to a merchant app may not be possible for some PSPs, so Hosted Link ensures the Link session can be completed wherever it is opened.
- **Easier access to Link callbacks and events:** When Link is launched via Plaid SDKs, Link events and metadata, including the public token, are typically returned via frontend callbacks. With Hosted Link, Link events are returned via backend API calls, making it easier to access information about the Link session without owning the frontend experience.

#### Overview of Hosted Link flow for embedded clients

##### In-app embeds

Hosted Link URLs are required to be opened in a secure context such as an `ASWebAuthenticationSession` on iOS or a Chrome Custom Tab on Android. Setting `hosted_link.is_mobile_app` to `true` when calling [`/link/token/create`](/docs/api/link/#linktokencreate) will indicate to Plaid that Hosted Link is being opened in a secure webview, so that Plaid will use the correct redirects.

##### Webpage embeds

If the flow is embedded in a webpage, there are two options for handling the Hosted Link URI.

###### Redirecting to Hosted Link

The merchant can redirect their entire page to the Plaid Hosted Link URI. Upon the user completing Link, Plaid will redirect that webpage to the `completion_redirect_uri` provided in the [`/link/token/create`](/docs/api/link/#linktokencreate) call.

###### Opening Hosted Link in a new window or tab

The merchant can open the Plaid Hosted Link URI in a new tab (for mobile web), or pop-up (for desktop web). This allows the user to link with Plaid while keeping the context of the checkout flow behind the pop-up. Upon completion, Plaid can close this pop-up / new tab, bringing the user back to the merchant checkout flow.

#### Hosted Link integration modes

There are two main types of flows that an embedded app may use. The instructions for integrating Hosted Link will be different for each type, so to integrate with Hosted Link, you will need to identify which one your app uses.

##### With intermediary steps

In the first type of flow, with intermediary steps, an end user is taken from the merchant to an intermediary page that is controlled directly by the PSP, where they will typically pick a payment method, complete the payment, and then be returned to the intermediary page, which will redirect the end user back to the merchant.

This flow provides slightly more complexity for the end customer, but does not require any additional integration work from the merchant.

##### Without intermediary steps

In the second type of flow, without intermediary steps, the end user never goes to a page controlled by the PSP. Instead, they will go directly from the merchant's app to the payment method, then return to the merchant's app.

This flow is simpler from the end customer's perspective, but may require additional integration steps by the merchant.

#### Integration process with intermediary steps

Follow the instructions for using [Hosted Link](/docs/link/hosted-link/). In the [`/link/token/create`](/docs/api/link/#linktokencreate) call, the `hosted_link.completion_redirect_uri` should be set to your website, which will then handle the process of redirecting the user back to the merchant page.

#### Integration process without intermediary steps

Follow the instructions for using [Hosted Link](/docs/link/hosted-link/), and see the sections below for instructions specific to the embedded client flow.

##### Completion redirect URIs

If you would like Plaid to redirect directly to the merchant app, you will use a separate `completion_redirect_uri` for each merchant, each of which must be added to the [Dashboard](https://dashboard.plaid.com/developers/api).

As an alternative, you may want the `completion_redirect_uri` to be a single URL that you host yourself, and manage the redirects back to merchant apps independently of Plaid. This way, you do not need to add each of the merchant’s redirect URIs to the Plaid Dashboard.

If no `completion_redirect_uri` is provided, the user will see a "Return to app" message on a Plaid.com domain in the browser and must manually navigate back to the merchant app.

###### Merchant integration steps

Because the application controls how URLs are handled, the merchant may need to add code to their app to open the Plaid hosted Link within the default browser, especially if the PSP's page is opened from within a webview.

[Example function for Android](https://github.com/plaid/plaid-link-examples/blob/master/webviews/android/LinkWebview/app/src/main/java/com/example/linkwebview/MainActivity.kt#L94-L157)

[Example function for iOS](https://github.com/plaid/plaid-link-examples/blob/master/webviews/wkwebview/wkwebview/LinkViewController.swift#L56-L72)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
