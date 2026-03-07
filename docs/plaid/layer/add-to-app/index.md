---
title: "Layer - Add Layer to your app | Plaid Docs"
source_url: "https://plaid.com/docs/layer/add-to-app/"
scraped_at: "2026-03-07T22:05:01+00:00"
---

# Plaid Layer integration guide

#### Use Plaid Layer to instantly onboard new customers

In this guide, we'll start from scratch and walk through how to use [Plaid Layer](/docs/api/products/layer/) to create a fast, frictionless onboarding experience for your customers.

If you're already familiar with Link, you can skip to [Create a Link token](/docs/layer/add-to-app/#create-a-link-token).

#### Get Plaid API keys and complete application and company profile

If you don't already have one, you'll need to [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You will also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile) in the Dashboard. The information in your profile will be shared with users of your application when they manage their connection on the [Plaid Portal](https://my.plaid.com).

#### Install and initialize Plaid libraries

You can use our official server-side client libraries to connect to the Plaid API from your application:

Terminal

```
// Install via npm
npm install --save plaid
```

After you've installed Plaid's client libraries, you can initialize them by passing in your `client_id`, `secret`, and the environment you wish to connect to (Sandbox or Production). This will make sure the client libraries pass along your `client_id` and `secret` with each request, and you won't need to explicitly include them in any other calls.

```
// Using Express
const express = require('express');
const app = express();
app.use(express.json());

const { Configuration, PlaidApi, PlaidEnvironments } = require('plaid');

const configuration = new Configuration({
  basePath: PlaidEnvironments.sandbox,
  baseOptions: {
    headers: {
      'PLAID-CLIENT-ID': process.env.PLAID_CLIENT_ID,
      'PLAID-SECRET': process.env.PLAID_SECRET,
    },
  },
});

const client = new PlaidApi(configuration);
```

#### Launch Link

Plaid Link is a drop-in module that provides a secure, elegant flow for Layer. Before initializing Link, you will need to create a new `link_token` on the server side of your application. A `link_token` is a short-lived, one-time use token that is used to authenticate your app with Link.
You can create one using the [`/session/token/create`](/docs/api/products/layer/#sessiontokencreate) endpoint. Then, on the client side of your application, you'll need to initialize Link with the `link_token` that you just created.

##### Create a Link token

Unlike a regular Link flow, the starting point for Layer is a call to [`/session/token/create`](/docs/api/products/layer/#sessiontokencreate) with a specified `TEMPLATE_ID`. The template should be created and configured in the dashboard ahead of time. If you want to use other products in addition to the `layer` product, make sure your template has them enabled.

```
app.post('/api/create_session_token', async function (request, response) {
  // Get the client_user_id by searching for the current user
  const user = await User.find(...);
  const clientUserId = user.id;
  const sessionTokenRequest = {
    user: {
      // This should correspond to a unique id for the current user.
      client_user_id: clientUserId,
    },
    template_id: TEMPLATE_ID,
  };
  try {
    const createTokenResponse = await client.sessionTokenCreate(sessionTokenRequest);
    response.json(createTokenResponse.data);
  } catch (error) {
    // handle error
  }
});
```

##### Install the Link SDKs

For instructions on installing Link SDKs, see the [Link documentation](/docs/link/) for your platform: [iOS](/docs/link/ios/), [Android](/docs/link/android/), [React Native](/docs/link/react-native/), or [React](/docs/link/web/). Layer is not supported on mobile webview integrations.

As Plaid is actively adding new Layer functionality, it is strongly recommended that you use the latest SDK version if your app uses Layer.

If you are developing a native mobile application, Layer requires SDK versions from June 2024 or later. Minimum versions are 6.0.4 (iOS), 4.5.0 (Android), and 11.11.0 (React Native). For Extended Autofill support, minimum versions are 6.3.1 (iOS), 5.3.0 (Android), 12.4.0 (React Native), and 4.1.1 (React).

Layer is not compatible with the [Hosted Link](https://plaid.com/docs/link/hosted-link/) integration mode.

##### Create the Link Handler

Basic sample code for creating the Link handler is shown below. For more details, see the [Link documentation](/docs/link/) for your platform: [iOS](/docs/link/ios/), [Android](/docs/link/android/), [React Native](/docs/link/react-native/), or [React](/docs/link/web/).

Ensure you are creating a Link token and passing it to the Link SDK as early as possible. The more time between when you create your handler and when you open Link, the better the UX will be for your users.

If you already have an existing Android or React Native integration created prior to June 2024, you will likely need to update your client-side Link opening code to support Layer. If your Android integration does not use a `PlaidHandler` or uses `OpenPlaidLink` instead of `FastOpenPlaidLink`, or if your React Native integration uses `PlaidLink` instead of `create` and `open`, you will need to add a separate handler creation step, as shown in the sample code below. For more details, see [React Native: opening Link](https://plaid.com/docs/link/react-native/#opening-link) and [Android: opening Link](/docs/link/android/#create-a-linktokenconfiguration).

Select group for content switcher

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

```
val plaidHandler: PlaidHandler = 
  Plaid.create(application, linkTokenConfiguration)
```

##### Open Link

```
linkAccountToPlaid.launch(plaidHandler)
```

##### Create a Configuration

Once the Link token is passed to your app, you will create an instance of `LinkTokenConfiguration`, then create a Handler using `Plaid.create()` passing the previously created `LinkTokenConfiguration`.

Create a Configuration

```
var linkConfiguration = LinkTokenConfiguration(
    token: "<#LINK_TOKEN_FROM_SERVER#>",
    onSuccess: { linkSuccess in
        // Send the linkSuccess.publicToken to your app server.
    }
)
```

##### Create a Handler

A `Handler` is a one-time use object used to open a Link session. The `Handler` must
be retained for the duration of the Plaid SDK flow. It will also be needed to respond to OAuth Universal Link
redirects. For more details, see the [OAuth guide](/docs/link/oauth/#ios).

Create a Handler

```
let result = Plaid.create(configuration)
switch result {
  case .failure(let error):
      logger.error("Unable to create Plaid handler due to: \(error)")
  case .success(let handler):
      self.handler = handler
}
```

Initiate the Link preloading process by invoking the `create` function.

Invoke Link create method

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

Call `Plaid.create` (or, if using React, `usePlaidLink`) when initializing the view that is responsible for loading Plaid.

Create example

```
// The usePlaidLink hook manages Plaid Link creation
// It does not return a destroy function;
// instead, on unmount it automatically destroys the Link instance
const config: PlaidLinkOptions = {
  onSuccess: (public_token, metadata) => {},
  onExit: (err, metadata) => {},
  onEvent: (eventName, metadata) => {},
  token: 'GENERATED_LINK_TOKEN',
};

const { open, exit, ready, submit } = usePlaidLink(config);
```

#### Submit the user’s phone number

Call the `submit` method on the Plaid handler you created earlier with the user's phone number. The semantics depend on the language/platform, but all methods are called `submit`.

Select group for content switcher

```
val submissionData = SubmissionData(phoneNumber)
plaidHandler.submit(submissionData)
```

Swift sample code

```
// Create a model that conforms to the SubmissionData interface
struct PlaidSubmitData: SubmissionData {
    var phoneNumber: String?
}

let data = PlaidSubmitData(phoneNumber: "14155550015")

self.handler.submit(data)
```

```
submit({
  "phone_number": "+14155550123"
})
```

```
handler.submit({
  "phone_number": "+14155550123"
})
```

#### Submit the user's date of birth

Upon receiving `LAYER_NOT_AVAILABLE` after phone number submission, collect the user's date of birth and call the `submit` method on the existing Plaid handler. For more details, see [Extended Autofill](/docs/layer/#extended-autofill).

This functionality requires minimum SDK versions 6.3.1 (iOS), 5.3.0 (Android), 12.4.0 (React Native), and 4.1.1 (React). If you are using an older SDK version or do not wish to use Extended Autofill, you can skip this step.

Select group for content switcher

```
val submissionData = SubmissionData(dateOfBirth = "1975-01-18")
plaidHandler.submit(submissionData)
```

Swift sample code

```
// Create a model that conforms to the SubmissionData interface
struct PlaidSubmitData: SubmissionData {
    var dateOfBirth: String?
    var phoneNumber: String?
}

let data = PlaidSubmitData(dateOfBirth: "1975-01-18")

self.handler.submit(data)
```

```
submit({
  "dateOfBirth": "1975-01-18"
})
```

```
handler.submit({
  "date_of_birth": "1975-01-18"
})
```

#### Open Layer UI on the LAYER\_READY event

Listen to the events on the Plaid handler via `onEvent`. For platform-specific details, see the [Link documentation](/docs/link/) for your platform.

If you receive `LAYER_READY`, the user is eligible for the Layer flow and you should proceed to open Link according to the [Link documentation](/docs/link/) for your platform.

If you receive `LAYER_AUTOFILL_NOT_AVAILABLE` (or if you receive `LAYER_NOT_AVAILABLE` and have not built Extended Autofill support for your integration), the user is not eligible for the Layer flow. You can clean up the handler you created earlier and fall back to whatever non-Layer onboarding flow fits your application (e.g. a traditional Link session, or other custom flow for your app).

Select group for content switcher

```
Plaid.setLinkEventListener { event ->
    when(event.eventName) {
        is LAYER_READY -> {
            // open link
            linkAccountToPlaid.launch(plaidHandler)
        }
        is LAYER_NOT_AVAILABLE -> {
            // run fall back flow
        }
        else -> { Log.i("Event", event.toString()) }
    }
}
```

Swift sample code

```
var linkSessionID: String?

linkTokenConfiguration.onEvent = { [weak self] linkEvent in
    guard let self = self else { return }
    switch linkEvent.eventName {
        case .layerReady:
            self.handler.open(presentUsing: .viewController(self))
            break
        case .layerNotAvailable:
            // Fall back on non-Layer flows, clean up
            break
        default:
            // Other cases ignored in this use case.
            break
    }
}
```

JavaScript sample code

```
usePlaidEmitter((event: LinkEvent) => {
  switch (event.eventName) {
    case LinkEventName.LAYER_READY:
      // Open Link
      open({...})
      break;
    case LinkEventName.LAYER_NOT_AVAILABLE:
      // Run another fallback flow
      break;
    default:
      // Other cases ignored in this use case
      break;
  }
});
```

JavaScript sample code

```
//Same onEvent handler from Link create sample
onEvent: (eventName, metadata) => {
  switch(eventName) {
    case "LAYER_READY":
      // Open Link
      open({...})
      break;
  case "LAYER_NOT_AVAILABLE":
    // Run another fallback flow
    break;
  default:
    //Other cases ignored in this use case
    break;
  }
}
```

#### Get the public token from the onSuccess callback

On successful completion, an `onSuccess` callback will be invoked, similar to the standard Link flow. Capture the `public_token` from `onSuccess`.

Select group for content switcher

```
val profileToken = success.publicToken
// Send the public token to your backend
```

Swift sample code

```
onSuccess: { linkSuccess in
}
// Send the public token to your backend
```

JavaScript sample code

```
const onSuccess = (linkSuccess: LinkSuccess) => {
  let publicToken = linkSuccess.publicToken;
  // Send the public token to your backend
};
```

JavaScript sample code

```
onSuccess: (public_token) => {
  const publicToken = public_token;
  // Send the public token to your backend
};
```

#### Get user account data

Call [`/user_account/session/get`](/docs/api/products/layer/#user_accountsessionget) to retrieve user-permissioned identity information as well as Item access tokens. Unlike typical Plaid Link sessions, where you must first exchange your public token for an access token in order to talk to the Plaid API, the [`/user_account/session/get`](/docs/api/products/layer/#user_accountsessionget) endpoint allows you to retrieve user-permissioned identity information as well as Item access tokens in a single call. You can optionally use Plaid products such as [Identity Match](/docs/identity/#identity-match) or [Identity Verification](/docs/identity-verification/) if you wish to verify this data.

Because Layer already verifies your user's phone number via OTP or SNA, for a low-friction experience, you should *not* perform additional OTP phone number verification as long as Layer has verified the number.

The best indication that the number was verified is the [`LAYER_AUTHENTICATION_PASSED`](/docs/api/products/layer/#layer_authentication_passed) webhook. Before skipping OTP verification based on this webhook, be sure to implement [webhook verification](https://plaid.com/docs/api/webhooks/webhook-verification/) to protect against webhook spoofing attacks.

Alternatively, you can use Link completion to indicate that the number was verified, but this method will exclude sessions where the user verified their number without fully completing Link. To use this method, call `/user_account/session_get`; the `phone_number` returned is the verified number. Do not rely on the `onSuccess` callback, as this can be vulnerable to client-side spoofing attacks.

/user\_account/session/get

```
const request: UserAccountSessionGetRequest = {
  public_token: 'profile-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce992d',
};
try {
  const response = await client.userAccountSessionGet(request);
} catch (error) {
  // handle error
}
```

Layer sample response

```
{
  "identity": {
    "phone_number": "+14155550015",
    "name": {
      "first_name": "Leslie",
      "last_name": "Knope"
    },
    "address": {
      "street": "123 Main St.",
      "street2": "Apt 123",
      "city": "Pawnee",
      "region": "Indiana",
      "postal_code": "46001",
      "country": "US"
    },
    "email": "leslie@knope.com",
    "date_of_birth": "1979-01-01",
    "ssn": "987654321",
    "ssn_last4": "4321"
  },
  "items": [
    {
      "item_id": <external_item_id>,
      "access_token": "access-token-<UUID>"
    }
  ],
  "request_id": "j0LkqT9OPdVwjwh"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
