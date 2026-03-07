---
title: "Link - Embedded Link (pay by bank) | Plaid Docs"
source_url: "https://plaid.com/docs/link/embedded-institution-search/"
scraped_at: "2026-03-07T22:05:04+00:00"
---

# Link Embedded Institution Search

#### Enhance uptake of Pay-By-Bank by embedding institution search directly into your app

With Embedded Institution Search (also known as Embedded Link), you can provide a more seamless transition to the account connectivity experience. For payment experiences in particular, embedding institution search can help guide users to choosing their bank account as a payment method.

Embedded Institution Search and [Database Auth](/docs/auth/coverage/database-auth/) are both designed to increase adoption of ACH payment methods and are frequently used together. Embedded Institution Search is also fully compatible with other Auth flows, including micro-deposit based flows.

It is highly recommended to use Embedded Institution Search for Link to increase uptake of pay-by-bank if your use case involves a pay-by-bank payments flow where the end user can choose between paying via a credit card and paying via a bank account. If your use case is an account opening or funding flow that requires the customer to use a bank account, or has a surcharge for credit card use, use the standard Link experience instead.

In order to realize the benefits of Embedded Institution Search, you must show it by default when the payments view loads, without requiring the user to first select 'Pay by bank'. If you cannot do this, use the standard Link experience instead.

Embedded Institution Search can be used with the following products: Auth, Identity, Balance, Signal, Transfer, Transactions, Investments, Assets, Liabilities, and all Plaid Check Consumer Report products.

Embedded Institution Search is not compatible with [Multi-Item Link](https://plaid.com/docs/link/multi-item-link/). Embedded Institution Search sessions will not be reported in [Link Analytics](/docs/link/measuring-conversion/#link-analytics).

For recommendations on how to improve the uptake of pay-by-bank and optimize your Embedded Institution Search UX, see [Increasing Pay-by-bank adoption](/docs/auth/pay-by-bank-ux/).

#### Institution Search user experience

The user will be able to select from a set of institution logos. This set of logos is customized based on user data, such as location. Alternatively, you can choose which logos will be displayed via the Plaid Dashboard. If a user selects an institution, they will be taken to Link to connect their account.

![Payment options screen with monthly plan selected, pay by bank method, search bar, and bank logos. Pay button at the bottom.](/assets/img/docs/embedded-search-desktop.png)

Example of Embedded Institution Search on Desktop

![Mobile flow showing institution search, login to Gingham Bank, and success confirmation via Plaid for payment setup.](/assets/img/docs/embedded-search-mobile.png)

Example of Embedded Institution Search flow on mobile when user selects an institution

If the user's institution is not one of the featured institutions, they can search for their institution using the search bar.

![Search for a bank in the checkout payment method section using the search bar or select from listed institutions; option to connect manually.](/assets/img/docs/embedded-search-search-bar.gif)

Example of searching for an institution in Embedded Institution Search on desktop

#### Returning user experience

The Embedded Institution Search experience can be further optimized for users who have already connected a financial account with Plaid.

![Returning user form in Embedded Institution Search with fields for name, address, city, state, and zip code, and a 'Continue' button.](/assets/img/docs/embedded-search-returning-user.gif)

Example of the returning user experience in Embedded Institution Search

When a user's device is recognized or the phone number is provided via [`/link/token/create`](/docs/api/link/#linktokencreate), Plaid will check to see if they are a returning user. If they are detected as a returning user, Link will run a security check on the session in the background. If this check passes, the user will be shown a list of previously connected accounts in the embedded module.

The user can then verify their phone number using a one-time password and complete the account linking process if the previously linked institutions do not require re-authentication. For more details on the returning user flow, see [Returning user experience](/docs/link/returning-user/).

Note that if Plaid Check Consumer Report products are included in the Link token request, then returning users will see the default Institution Search user experience.

##### Testing in Sandbox

In order to test the returning user experience in Embedded Institution Search, use the phone numbers listed in [Testing returning user experience in Sandbox](/docs/link/returning-user/#testing-in-sandbox).

#### Integration steps

Embedded Institution Search is available for all supported integration modes except webviews and Hosted Link.

Embedded Institution Search cannot be rendered inside an iFrame.

Before integrating, make sure you are using a version of the SDK that supports Embedded Institution Search.

| Platform | Minimum SDK version required |
| --- | --- |
| [Android](https://github.com/plaid/plaid-link-android) | 3.14.0 |
| [iOS](https://github.com/plaid/plaid-link-ios) | 4.5.0 |
| [React Native](https://github.com/plaid/react-native-plaid-link-sdk) | 10.6.0 |
| [React web](https://github.com/plaid/react-plaid-link) | 3.5.1 |
| JavaScript web | N/A (all customers are on the latest version) |

1. Create a Link token as normal, using [`/link/token/create`](/docs/api/link/#linktokencreate); no special parameters are required when making this call to enable Embedded Institution Search. The 'Connect Manually' button rendering is configured with the `auth_type_select_enabled` boolean in the `auth` object.
2. Create an embedded view. This should be done before it will be displayed in your app, to minimize latency. Plaid will track when this view is requested and activate the Embedded Search UI.
   If you are using iOS, call `createEmbeddedView`, which will return a Result containing a `UIView`. Once you have the `UIView`, add it to your `ViewController`'s view.
   If you are using the web SDK, call `Plaid.createEmbedded` instead of `Plaid.create` to open Link.
   For other platforms, you will create the view as normal.
3. Lay out the view Plaid returns from the configuration. Using a **minimum size of 350x300px or 300x350px** is strongly recommended. Smaller sizes will significantly degrade conversion. If you are migrating to Embedded Institution Search from traditional Link, delete your old Link presentation in favor of this new one.

Select group for content switcher

Embedded Institution Search - iOS (Swift)

```
private func setupEmbeddedSearchView(token: String) {

        // Create a configuration like normal.
        var configuration = LinkTokenConfiguration(
            token: token,
            onSuccess: { success in
                print("success: \(success)")
            }
        )

        configuration.onEvent = { event in
            print("Event: \(event)")
        }

        configuration.onExit = { exit in
            print("Exit: \(exit)")
        }

        // Create a handler with your configuration like normal.
        let handlerResult = Plaid.create(configuration)
        switch handlerResult {
        case .success(let handler):

            // Save a reference to your handler.
            self.handler = handler

            // Create an embedded view.
            let embeddedSearchView = handler.createEmbeddedView(presentUsing: .viewController(self))

            // Layout this view
            embeddedSearchView.translatesAutoresizingMaskIntoConstraints = false
            view.addSubview(embeddedSearchView)

            NSLayoutConstraint.activate([
                embeddedSearchView.topAnchor.constraint(equalTo: view.safeAreaLayoutGuide.topAnchor, constant: 8),
                embeddedSearchView.leadingAnchor.constraint(equalTo: view.leadingAnchor, constant: 25),
                embeddedSearchView.trailingAnchor.constraint(equalTo: view.trailingAnchor, constant: -25),
                embeddedSearchView.heightAnchor.constraint(equalToConstant: 360),
            ])

        case .failure(let error):
            // Error creating handler. Handle like normal.
            fatalError("\(error)")
        }
    }
```

Embedded Institution Search - Android (Kotlin)

```
// Register a callback for an Activity Result like normal (must be done from an Activity)
 private val linkAccountToPlaid =
   registerForActivityResult(OpenPlaidLink()) { result ->
      when (result) {
        is LinkSuccess -> /* handle LinkSuccess */
        is LinkExit -> /* handle LinkExit (from LinkActivity) */
      }
    }

// Create a linkTokenConfiguration like normal
val linkTokenConfiguration = LinkTokenConfiguration.Builder().token(token).build()

// Create the view with a trailing lambda for handling LinkExits from the Embedded View
val embeddedView = Plaid.createLinkEmbeddedView(
this /*Activity context*/,
linkTokenConfiguration,
linkAccountToPlaid) {
    exit: LinkExit -> /* handle LinkExit (from Embedded View) */
}

// Add this embeddedView to a view in your layout
binding.embeddedViewContainer.addView(embeddedView)
```

Embedded Institution search - web (HTML)

```
<div id="plaid-embedded-link-container"></div>
```

Embedded Institution Search - web (JavaScript)

```
// The container at `#plaid-embedded-link-container` will need to be sized in order to
// control the size of the embedded Plaid module
const embeddedLinkOpenTarget = document.querySelector('#plaid-embedded-link-container');

Plaid.createEmbedded(
    {
      token: 'GENERATED_LINK_TOKEN',
      onSuccess: (public_token, metadata) => {},
      onLoad: () => {},
      onExit: (err, metadata) => {},
      onEvent: (eventName, metadata) => {},
    },
    embeddedLinkOpenTarget,
);
```

Embedded Institution Search - web (React)

```
import React, { useCallback } from 'react';
import { PlaidEmbeddedLink } from 'react-plaid-link';

const App = props => {
  const onSuccess = useCallback(
    (token, metadata) => console.log('onSuccess', token, metadata),
    []
  );

  const onEvent = useCallback(
    (eventName, metadata) => console.log('onEvent', eventName, metadata),
    []
  );

  const onExit = useCallback(
    (err, metadata) => console.log('onExit', err, metadata),
    []
  );

  const config = {
    token: "plaid-token-123",
    onSuccess,
    onEvent,
    onExit,
  };

  return (
    <PlaidEmbeddedLink
      {...config}
      style={{
        height: '350px',
        width: '350px',
      }}
    />
  );
};

export default App;
```

For an end-to-end example app, see the [Transfer Quickstart](https://github.com/plaid/transfer-quickstart/), which incorporates Embedded Institution Search in a JavaScript frontend.

##### Update mode

Embedded Institution Search cannot be used in [update mode](/docs/link/update-mode/); to update a user's Item, launch a regular update mode session instead.

#### Customization Options

You can customize the Embedded Institution Search user experience to match your application's needs.

##### Module responsiveness and tile count

The embedded Link module will responsively scale to display between two and fifteen institutions, depending on the height and width of the module. Using a **minimum size of 350x300px or 300x350px** is strongly recommended, as smaller sizes degrade conversion particularly on returning user flows. Testing both embedded search and returning user experiences at your chosen module size to ensure usability is also strongly recommended.

On the web, the embedded Plaid module will attempt to use 100% of the height and width of its container. To modify the size of the Plaid module, or to render the Plaid module responsively, you should set the size of the container.

Plaid recommends using the largest container that will comfortably fit in your UI. The larger the container, the more institution logos will be shown. Users are more likely to adopt pay-by-bank when they see their institution's logo on the embedded Link screen.

##### Customizing institutions displayed

The institutions displayed in Link Embedded Search are based on your [Institution Select settings](https://dashboard.plaid.com/link/institution-select), which you can optionally customize in the [Dashboard](https://dashboard.plaid.com/link/institution-select). For most customers, it is recommended to use the default settings.

Embedded Institution Search is compatible with the [Institution Select Shortcut](/docs/link/customization/#institution-select-shortcut): If you already know which institution the user wants to connect to before initializing Link, you can pass `routing_number` into the `institution_data` request field in the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint. The matched institution will be listed first (top left position) in the embedded institution grid.

#### Event callbacks emitted by Embedded Institution Search

###### User chooses an institution directly from embedded search

- `onEvent: OPEN` – `view_name: "CONSENT"`
- `onEvent: SELECT_INSTITUTION`
- `onEvent: TRANSITION_VIEW` – `view_name: "CONSENT"`
- `onEvent: TRANSITION_VIEW` – `view_name: "CREDENTIAL" or "OAUTH"` - user selects Continue on the ConsentPane

#### UI Recommendations for Embedded Institution Search

See [Increasing pay-by-bank adoption](/docs/auth/pay-by-bank-ux/) for recommendations on displaying Embedded Institution Search within your app for pay-by-bank use cases.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
