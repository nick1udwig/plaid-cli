---
title: "Link - Optimizing Link conversion | Plaid Docs"
source_url: "https://plaid.com/docs/link/best-practices/"
scraped_at: "2026-03-07T22:05:03+00:00"
---

# Optimizing Link conversion

#### Discover best practices for improving Link conversion

#### Overview

This guide contains tips for optimizing your existing Link implementation to help increase conversion and improve your users’ experiences with Link. If you are new to Link or don’t yet have a working Link integration, see [Link overview](/docs/link/) for instructions on getting started with Link.

If you are using pay-by-bank flows, also see [Increasing pay-by-bank adoption](/docs/auth/pay-by-bank-ux/) for tips on getting users to use the Link-based payments flow rather than paying via card.

#### Measuring Link conversion

Before making changes to your integration to improve conversion, you should first understand your current conversion rates. To learn how, see [Tracking Link conversion](/docs/link/measuring-conversion/).

#### Improving Link conversion

Many different steps can be taken to maximize Link conversion, and the exact impact of these steps will vary for each customer and use case. The recommendations below are provided in general priority order.

##### Initialize with a minimal set of products

In general, calling [`/link/token/create`](/docs/api/link/#linktokencreate) with a minimal set of products will both increase conversion and reduce your costs. For more details, see [Choosing how to initialize products](/docs/link/initializing-products/).

##### Implement full OAuth support on mobile, including app-to-app

While supporting mobile app-to-app impacts only a small number of banks (currently only Chase), the impact on conversion for eligible users is very large, as app-to-app flows can allow users to authenticate with biometrics instead of a username and password. For iOS, supporting app-to-app requires [creating an `apple-app-site-association` file](/docs/link/oauth/#ios-sdk-react-native-ios) to support universal links. On Android, supporting app-to-app requires [registering your Android package name](/docs/link/oauth/#android-sdk-react-native-android). For more details, see the [OAuth Guide](/docs/link/oauth/).

##### Pre-initialize Link for lower UI latency

All Plaid SDKs offer a method to pre-initialize Link before opening it -- for example, by calling the `create()` method on web, React Native, or iOS, or creating a `PlaidHandler` on Android. It is recommended to call this method as soon as you load the view or screen from which the end user can open Link. Pre-initializing Link as soon as possible, before the `open()` method is called, results in lower latency for your end-user in loading the Link UI, leading to increased conversion.

Support for pre-initialization was added in March 2024. You must be using an up-to-date version of Plaid Link SDKs to benefit from the performance enhancements of this approach.

##### Provide pre-Link messaging

Link conversion is highest when users have the right expectations set going into Link. Your app should explain why you use Plaid, the benefits for the user of connecting their account, and that the user's information will be secure. It should also set the user's expectations around what information they will need to provide during the Link flow. Plaid should be configured as the default experience. For more details, including visual examples, see [Pre-Link messaging for optimizing conversion](/docs/link/messaging/).

##### Update the Link SDK regularly

Plaid makes continual updates to Plaid Link to improve conversion. If you are using Link on web, you will automatically be on the latest version; if you are using one of Plaid's SDKs for [iOS](https://github.com/plaid/plaid-link-ios), [Android](https://github.com/plaid/plaid-link-android), or [React Native](https://github.com/plaid/react-native-plaid-link-sdk), we recommend you update regularly (ideally once a month, and at least once per quarter) to ensure you have access to the latest Link conversion improvements.

##### Provide a phone number during Link token creation

If you have the user's phone number, provide it in the [`/link/token/create`](/docs/api/link/#linktokencreate) request in the `user.phone_number` field. Plaid will pre-fill the phone number in Link for the user, allowing them to access the Returning User Experience without any data entry. For more details, see [Returning user experience](/docs/link/returning-user/).

##### (For Auth customers) Implement alternative Auth flows

While the vast majority of users can use Plaid's default Auth flows, some users, especially those at smaller banks and credit unions, have accounts that do not support those flows. Implementing micro-deposit-based or Database Auth flows will increase conversion by allowing these users to link their accounts. For more details, see [Auth coverage](/docs/auth/coverage/).

##### (For pay by bank use cases) Implement Embedded Institution Search

If you are building pay by bank flows, implementing the [Embedded Institution Search](/docs/link/embedded-institution-search/) UI for Link has been shown to heavily increase the percentage of customers who choose to pay by bank rather than via a different means of payment.

##### Use Link Recovery

[Link Recovery (beta)](/docs/link/link-recovery/) allows you to recapture end users who abandoned Link due to temporary institution outages and notify them when the issue is resolved. To learn more and request access to the beta, see [Link Recovery](/docs/link/link-recovery/).

##### Configure Link for your user's country and language

For apps with multi-language experiences, [custom profiles](/docs/link/customization/) improve conversion by allowing you to display Link in your user's preferred language. Calling [`/link/token/create`](/docs/api/link/#linktokencreate) with the `country` parameter set to your user's specific country, rather than all countries your app supports, will allow Link to show a more accurately targeted list of institutions. For Auth customers, calling [`/link/token/create`](/docs/api/link/#linktokencreate) with `country` set to just `US` is also required to enable the conversion-maximizing Instant Match and micro-deposit based [Auth flows](/docs/auth/coverage/).

##### Customize Link with your organization's branding

Plaid allows you to customize certain aspects of the Link UI. Customizing these in a way that matches your app -- for example, [uploading your organization's logo](/docs/link/customization/#consent-pane-customizations) to be used on the Link consent pane or [matching your brand colors](/docs/link/customization/#color-scheme) can increase conversion. For more details, see [Customizing Link](/docs/link/).

##### (If applicable) Use the Institution Select shortcut

If you already know the institution your user plans to connect before the Link flow is launched, you can highlight this institution in the Institution Select UI by specifying the [routing number](/docs/api/link/#link-token-create-request-institution-data-routing-number) during the [`/link/token/create`](/docs/api/link/#linktokencreate) call. For more details, see [Institution Select shortcut](/docs/link/customization/#institution-select-shortcut).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
