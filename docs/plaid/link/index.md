---
title: "Link - Overview | Plaid Docs"
source_url: "https://plaid.com/docs/link/"
scraped_at: "2026-03-07T22:05:02+00:00"
---

# Link overview

#### Use Link to connect to your users' financial accounts with the Plaid API

=\*=\*=\*=

#### Introduction to Link

Plaid Link is the client-side component that your users will interact with in order to link their accounts to Plaid and allow you to access their accounts via the Plaid API. Using Link is mandatory for all Plaid integrations, except for ones using only the handful of products that do not require end user interaction, such as [Enrich](/docs/enrich/) or the [Identity Verification backend only-flow](/docs/identity-verification/#data-source-checks-without-ui-backend-flow).

An example Link flow. Your exact flow may differ.

![An example Link flow. Your exact flow may differ.](/assets/img/docs/link-tour/link-1-rm.png)

![...the user selects their bank...](/assets/img/docs/link-tour/link-2-original-backup.png)

![...is directed to its OAuth flow...](/assets/img/docs/link-tour/link-3-original-backup.png)

![...signs in via OAuth...](/assets/img/docs/link-tour/link-4.png)

![...selects which accounts to share...](/assets/img/docs/link-tour/link-select-accounts.png)

![...and links their account!](/assets/img/docs/link-tour/link-5.png)

![In the guest flow, they can optionally save their account for fast access next time.](/assets/img/docs/link-tour/link-6.png)

![Once the phone number is verified, Link will close automatically.](/assets/img/docs/link-tour/link-7.png)

Plaid Link handles all aspects of the login and authentication experience, including credential validation, multi-factor authentication, error handling, and sending account linking confirmation emails. For institutions that use OAuth, Link also manages the OAuth handoff flow, bringing the user to their institution to log in, and then returning them to the Plaid Link experience within your app. Link is supported via SDKs for all modern browsers and platforms, including [web](/docs/link/web/), [iOS](/docs/link/ios/), [Android](/docs/link/android/), as well as via [React Native](/docs/link/react-native/), along with community-supported wrappers for [Flutter](https://github.com/jorgefspereira/plaid_flutter), [Angular](https://github.com/mike-roberts/ngx-plaid-link), and [Vue](https://github.com/jclaessens97/vue-plaid-link/).

For webview-based integrations or integrations that don't have a frontend, Plaid also provides a drop-in [Hosted Link](/docs/link/hosted-link/) integration mode.

To try Link, see [Plaid Link Demo](https://plaid.com/demo/).

Link is the only available method for connecting accounts and authenticating users in Production. In the Sandbox test environment, Link can optionally be bypassed for testing purposes via [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate).

=\*=\*=\*=

#### Initializing Link

Link is initialized by passing the `link_token` to Link. The exact implementation details for passing the `link_token` will vary by platform. For detailed instructions, see the page for your specific platform: [web](/docs/link/web/), [iOS](/docs/link/ios/), [Android](/docs/link/android/), [React Native](/docs/link/react-native/), [mobile webview](/docs/link/webview/), or [Plaid-hosted](/docs/link/hosted-link/).

For recommendations on configuring the `link_token` for your use case, see [Choosing how to initialize products](/docs/link/initializing-products/).

=\*=\*=\*=

#### Link flow overview

Most Plaid products use Link to generate `public_tokens`. The diagram below shows a model of how Link is used to obtain a `public_token`, which can then be exchanged for an `access_token`, which is used to authenticate requests to the Plaid API.

Note that some products (including Plaid Check, Identity Verification, Monitor, Document Income, and Payroll Income) do not use a `public_token` or `access_token`. For those products, you will call product endpoints once the end user has completed Link; see product-specific documentation for details on the flow.

**The Plaid flow** begins when your user wants to connect their bank account to your app.

![Step  diagram](/assets/img/docs/link-tokens/link-token-row-1.png)

**1**Call [`/link/token/create`](/docs/api/link/#linktokencreate) to create a `link_token` and pass the temporary token to your app's client.

![Step 1 diagram](/assets/img/docs/link-tokens/link-token-row-2.png)

**2**Use the `link_token` to open Link for your user. In the [`onSuccess` callback](/docs/link/web/#onsuccess), Link will provide a temporary `public_token`. This token can also be obtained on the backend via `/link/token/get`.

![Step 2 diagram](/assets/img/docs/link-tokens/link-token-row-3.png)

**3**Call [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to exchange the `public_token` for a permanent `access_token` and `item_id` for the new `Item`.

![Step 3 diagram](/assets/img/docs/link-tokens/link-token-row-4.png)

**4**Store the `access_token` and use it to make product requests for your user's `Item`.

![Step 4 diagram](/assets/img/docs/link-tokens/link-token-row-5.png)

In code, this flow is initiated by creating a `link_token` and using it to initialize Link. The
`link_token` can be configured with the Plaid products you will be using and the countries you will
need to support.

Once the user has logged in via Link, Link will issue a `public_token`. You can obtain the `public_token` through either the frontend or the backend:

- On the frontend: From the client-side `onSuccess` callback returned by Link after a successful session. For more details on this method, see the Link frontend documentation for your specific platform.
- On the backend: From the [`/link/token/get`](/docs/api/link/#linktokenget) endpoint or opt-in [`SESSION_FINISHED`](/docs/api/link/#session_finished) webhook after the Link session has been completed successfully. For more details on this method, see the [Hosted Link](/docs/link/hosted-link/) documentation.

The `public_token` can then be exchanged for an `access_token` via [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange).

=\*=\*=\*=

#### Supporting OAuth

Many institutions use an OAuth authentication flow, in which Plaid Link redirects the end user to their bank's website or mobile app to authenticate. To learn how to connect to an institution that uses OAuth, see the [OAuth guide](/docs/link/oauth/).

=\*=\*=\*=

#### Customizing Link

You can customize parts of Link's flow, including some text elements, the institution select view, and the background color, and enable additional features like the Account Select view straight from the [Dashboard](https://dashboard.plaid.com/link). You can preview your changes in realtime and then publish them instantly once you're ready to go live. For more details, see [Link customization](/docs/link/customization/).

To help you take advantage of the options available for customizing and configuring Link, Plaid offers a [best practices guide](/docs/link/best-practices/) with recommendations for how to initialize and configure Link within your app.

Link's appearance will also automatically change if the institution selected is not in a healthy state. For more details, see [Institution status in Link](/docs/link/institution-status/).

=\*=\*=\*=

#### Returning user flows

The returning user flow allows you to enable a faster Plaid Link experience for your users who already use Plaid. To learn more, see [Returning user experience](/docs/link/returning-user/).

=\*=\*=\*=

#### Error-handling flows

If your application will access an Item on a recurring basis, rather than just once, it should support [update mode](/docs/link/update-mode/). Update mode allows you to refresh an Item if it enters an error state, such as when a user changes their password or MFA information. For more information, see [Updating an Item](/docs/link/update-mode/).

It's also recommended to have special handling for when a user attempts to link the same Item twice. Requesting access tokens for duplicate Items can lead to higher bills and end-user confusion. To learn more, see [preventing duplicate Items](/docs/link/duplicate-items/).

Occasionally, Link itself can enter an error state if the user takes over 30 minutes to complete the Link process. For information on handling this flow, see [Handling invalid Link tokens](/docs/link/handle-invalid-link-token/).

=\*=\*=\*=

#### Optimizing Link conversion

How you configure Link can have a huge impact on the percentage of users who successfully complete the Link flow. To ensure you're maximizing conversion, see [Best practices for Link conversion](/docs/link/best-practices/).

=\*=\*=\*=

#### Troubleshooting

Since all your users will go through Link, it's important to build as robust an integration as possible. For details on dealing with common problems, see the [Troubleshooting](/docs/link/troubleshooting/) section.

=\*=\*=\*=

#### Link updates

Plaid periodically updates Link to add new functionality and improve conversion. These changes will be automatically deployed. Any test suites and business logic in your app should be robust to the possibility of changes to the user-facing Link flow.

Users of Plaid's SDKs for React, React Native, iOS, and Android should regularly update to ensure support for the latest client platforms and Plaid functionality.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
