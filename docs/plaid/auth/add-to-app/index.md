---
title: "Auth - Add Auth to your app | Plaid Docs"
source_url: "https://plaid.com/docs/auth/add-to-app/"
scraped_at: "2026-03-07T22:04:30+00:00"
---

# Add Auth to your app

#### Use Auth to connect user bank accounts

In this guide, we'll demonstrate how to add [Auth](/docs/api/products/auth/) to your app so that you can connect to your users' bank accounts and obtain the information needed to set up funds transfers.

If you're already familiar with using Plaid and are set up to make calls to the Plaid API, see [Getting Auth data](/docs/auth/add-to-app/#getting-auth-data). If you're interested in using a Plaid partner, such as Stripe or Dwolla, to process payments, see [Moving funds with a payment partner](/docs/auth/add-to-app/#moving-funds-with-a-payment-partner).

Prefer to learn by watching? A [video guide](https://youtu.be/FlZ5nzlIq74) is
available for this topic.

#### Get Plaid API keys and complete application and company profile

If you don't already have one, you'll need to [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You will also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile) in the Dashboard. The information in your profile will be shared with users of your application when they manage their connection on the [Plaid Portal](https://my.plaid.com). Your application profile and company profile must be completed before connecting to certain institutions in Production.

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

#### Create an Item using Link

These instructions cover Link for web applications. For instructions on using Link in mobile apps, see the [Link documentation](/docs/link/).

Plaid Link is the client-side component that your users interact with to securely connect their bank accounts to your app. An [Item](/docs/quickstart/glossary/#item) is created when a user successfully logs into their financial institution using Link. An Item represents a single login at a financial institution. Items do not represent individual bank accounts, although all bank accounts are associated with an Item. For example, if a user has a single login at a bank that allows them to access both a checking account and a savings account, only a single Item would be associated with both of the accounts.

When using Auth, you will typically only need access to the specific bank account that will be used to transfer funds, rather than all of the accounts a user may have at an institution. Because of this, it is recommended that you configure Link with [Account Select](/docs/link/customization/#account-select) when using Auth. Configuring Link with Account Select will limit unnecessary access to user accounts. You can [enable Account Select from the Dashboard](https://dashboard.plaid.com/link/account-select).

To create an Item, you'll first need to create a [Link token](/docs/quickstart/glossary/#link-token) on your application server. You can create a Link token by calling the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint. Then, on the client-side of your application, you'll initialize Link with using the Link token you created.

The code samples below demonstrate how to create a Link token and how to initialize Link using the token.

##### Create a Link token

```
app.post('/api/create_link_token', async function (request, response) {
  // Get the client_user_id by searching for the current user
  const user = await User.find(...);
  const clientUserId = user.id;
  const linkTokenRequest = {
    user: {
      // This should correspond to a unique id for the current user.
      client_user_id: clientUserId,
    },
    client_name: 'Plaid Test App',
    products: ['auth'],
    language: 'en',
    webhook: 'https://webhook.example.com',
    redirect_uri: 'https://domainname.com/oauth-page.html',
    country_codes: ['US'],
  };
  try {
    const createTokenResponse = await client.linkTokenCreate(linkTokenRequest);
    response.json(createTokenResponse.data);
  } catch (error) {
    // handle error
  }
});
```

##### Initialize Link with the Link token

If your use case involves a pay-by-bank payments flow where the end user can choose between paying via a credit card and paying via a bank account, it is highly recommended to use the [Embedded experience](https://plaid.com/docs/link/embedded-institution-search/) for Link to increase uptake of pay-by-bank. If your use case is an account opening or funding flow that requires the customer to use a bank account, or has a surcharge for credit card use, use the standard Link experience.

Select group for content switcher

app.js

```
const handler = Plaid.create({
  token: (await $.post('/create_link_token')).link_token,
  onSuccess: (public_token, metadata) => {
    // Upon successful linking of a bank account,
    // Link will return a public token.
    // Exchange the public token for an access token
    // to make calls to the Plaid API.
    $.post('/exchange_public_token', {
      public_token: public_token,
    });
  },
  onLoad: () => {},
  onExit: (err, metadata) => {
    // Optionally capture when your user exited the Link flow.
    // Storing this information can be helpful for support.
  },
  onEvent: (eventName, metadata) => {
    // Optionally capture Link flow events, streamed through
    // this callback as your users connect an Item to Plaid.
  },
});

handler.open();
```

The sample code below is for JavaScript. For more examples for other platforms, see [Embedded Link](/docs/link/embedded-institution-search/#integration-steps).

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

#### Obtain an access token

When a user successfully links an Item via Link, the [`onSuccess`](/docs/link/web/#onsuccess) callback will be called. The `onSuccess` callback returns a [public token](/docs/quickstart/glossary/#public-token). On your application server, exchange the public token for an [access token](/docs/quickstart/glossary/#access-token) and an [Item ID](/docs/quickstart/glossary/#item-id) by calling `/item/public_token/exchange/`. The access token will allow you to make authenticated calls to the Plaid API.

Store the access token and Item ID in a secure datastore, as they’re used to access Item data and identify webhooks, respectively. The access token will remain valid unless you actively expire it via rotation or remove it by calling [`/item/remove`](/docs/api/items/#itemremove) on the corresponding Item. For security purposes, never store the access token in client-side code. The public token is a one-time use token with a lifetime of 30 minutes, so there is no need to store it.

##### Exchange public token for an access token

```
app.post('/api/exchange_public_token', async function (
  request,
  response,
  next,
) {
  const publicToken = request.body.public_token;
  try {
    const tokenResponse = await client.itemPublicTokenExchange({
      public_token: publicToken,
    });

    // These values should be saved to a persistent database and
    // associated with the currently signed-in user
    const accessToken = tokenResponse.data.access_token;
    const itemID = tokenResponse.data.item_id;

    response.json({ public_token_exchange: 'complete' });
  } catch (error) {
    // handle error
  }
});
```

#### Getting Auth data

Now that you have an access token, you can begin making authenticated calls to the Plaid API. If you are processing payments yourself, the [`/auth/get`](/docs/api/products/auth/#authget) endpoint to retrieve the bank account and bank identification numbers (such as routing numbers, for US accounts) associated with an Item's accounts. You can then supply these to your payment system. If you are using a Plaid partner to move funds, you will not need to call [`/auth/get`](/docs/api/products/auth/#authget). Instead, see [Moving funds with a payment partner](/docs/auth/add-to-app/#moving-funds-with-a-payment-partner).

```
const { AuthGetRequest } = require('plaid');

// Call /auth/get and retrieve account numbers for an Item
const request = {
  access_token: access_token,
};
try {
  const response = await plaidClient.authGet(request);
  const accountData = response.data.accounts;
  if (response.data.numbers.ach.length > 0) {
    // Handle ACH numbers (US accounts)
    const achNumbers = response.data.numbers.ach;
  }
  if (response.data.numbers.eft.length > 0) {
    // Handle EFT numbers (Canadian accounts)
    const eftNumbers = response.data.numbers.eft;
  }
  if (response.data.numbers.international.length > 0) {
    // Handle International numbers
    const internationalNumbers = response.data.numbers.international;
  }
  if (response.data.numbers.bacs.length > 0) {
    // Handle BACS numbers (British accounts)
    const bacsNumbers = response.data.numbers.bacs;
  }
} catch (error) {
  //handle error
}
```

Example response data is below. Note that this is test account data; real accounts would not include all four sets of numbers.

/auth/get sample response

```
{
  "accounts": [
    {
      "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
      "balances": {
        "available": 100,
        "current": 110,
        "limit": null,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      },
      "mask": "9606",
      "name": "Plaid Checking",
      "official_name": "Plaid Gold Checking",
      "subtype": "checking",
      "type": "depository"
    }
  ],
  "numbers": {
    "ach": [
      {
        "account": "9900009606",
        "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
        "routing": "011401533",
        "wire_routing": "021000021"
      }
    ],
    "eft": [
      {
        "account": "111122223333",
        "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
        "institution": "021",
        "branch": "01140"
      }
    ],
    "international": [
      {
        "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
        "bic": "NWBKGB21",
        "iban": "GB29NWBK60161331926819"
      }
    ],
    "bacs": [
      {
        "account": "31926819",
        "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
        "sort_code": "601613"
      }
    ]
  },
  "item": {
    "available_products": [
      "balance",
      "identity",
      "payment_initiation",
      "transactions"
    ],
    "billed_products": ["assets", "auth"],
    "consent_expiration_time": null,
    "error": null,
    "institution_id": "ins_117650",
    "item_id": "DWVAAPWq4RHGlEaNyGKRTAnPLaEmo8Cvq7na6",
    "webhook": "https://www.genericwebhookurl.com/webhook"
  },
  "request_id": "m8MDnv9okwxFNBV"
}
```

#### Moving funds with a payment partner

You can move money via ACH transfer by pairing Auth with one of Plaid's payment partners. When using a partner to move money, the partner's payments service will initiate the transfer; Plaid does not function as the payment processor. For the full list of payments platforms who have partnered with Plaid to provide ACH money movement, see [Auth Partnerships](/docs/auth/partnerships/).

To move money using a Plaid partner, first create an Item using Link and obtain an access token as described above. Then, instead of calling Plaid's Auth endpoints, call one of Plaid's [processor token endpoints](/docs/api/processors/) to create a processor token. You can then send this processor token to one of Plaid's partners by using endpoints that are specific to the payment platform. Refer to the partner's technical documentation for more information. Using a partner to transfer funds gives you access to payment functionality while freeing you from having to securely store sensitive bank account information.

The sample code below demonstrates a call to [`/processor/token/create`](/docs/api/processors/#processortokencreate) using Dwolla as the payment processor.

Sample /processor/token/create call

```
const { ProcessorTokenCreateRequest } = require('plaid');

try {
  // Create a processor token for a specific account id
  const request: ProcessorTokenCreateRequest = {
    access_token: accessToken,
    account_id: accountID,
    processor: 'dwolla',
  };
  const processorTokenResponse = await plaidClient.processorTokenCreate(
    request,
  );
  const processorToken = processorTokenResponse.data.processor_token;
} catch (error) {
  // handle error
}
```

#### Tutorial and example code in Plaid Pattern

For a real-life example of an app that incorporates Auth, see the Node-based [Plaid Pattern Account Funding](https://github.com/plaid/pattern-account-funding) sample app. Pattern Account Funding is a sample account funding app that fetches Auth data in order to set up funds transfers. The Auth code can be found in [items.js](https://github.com/plaid/pattern-account-funding/blob/master/server/routes/items.js#L81-L135).

For a step-by-step tutorial on how to implement account funding, [Account funding tutorial](https://github.com/plaid/account-funding-tutorial).

#### Next steps

Once Auth is implemented in your app, see [Micro-deposit and database verification](/docs/auth/coverage/) to make sure your app is supporting the maximum number of institutions and verification methods (US only).

If your use case is an account funding use case, see [the Account funding solutions guide](https://plaid.com/documents/plaid-account-funding-guide.pdf) for a set of recommendations on how to implement Auth.

If you're ready to launch to Production, see the Launch checklist.

[#### Launch checklist

Recommended steps to take before launching in Production

Launch](/docs/launch-checklist/)

#### Launch checklist

Recommended steps to take before launching in Production

[Launch](/docs/launch-checklist/)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
