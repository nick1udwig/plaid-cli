---
title: "Identity - Add Identity to your app | Plaid Docs"
source_url: "https://plaid.com/docs/identity/add-to-app/"
scraped_at: "2026-03-07T22:04:57+00:00"
---

# Add Identity to your app

#### Use Identity to verify user data

In this guide, we'll start from scratch and walk through how to use [Identity](/docs/api/products/identity/) to retrieve identity data. If you are already familiar with using Plaid and are set up to make calls to the Plaid API, you can skip ahead to [Matching identity data](/docs/identity/add-to-app/#matching-identity-data) (for [`/identity/match`](/docs/api/products/identity/#identitymatch)) or [Fetching identity data](/docs/identity/add-to-app/#fetching-identity-data) (for [`/identity/get`](/docs/api/products/identity/#identityget)).

#### Get Plaid API keys and complete application and company profile

If you don't already have one, you'll need to [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You will also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile) on the Dashboard. The information in your profile will be shared with users of your application when they manage their connection on the [Plaid Portal](https://my.plaid.com). Your application profile and company profile must be completed before connecting to certain institutions in Production.

#### Install and initialize Plaid libraries

You can use our official server-side client libraries to connect to the Plaid API from your application:

Terminal

```
// Install via npm
npm install --save plaid
```

After you've installed Plaid's client libraries, you can initialize them by passing in your `client_id`, `secret`, and the environment you wish to connect to (Sandbox or Production). This will make sure the client libraries pass along your `client_id` and `secret` with each request, and you won't need to explicitly include them in any other calls.

In the code samples below, you will need to replace `PLAID_CLIENT_ID` and `PLAID_SECRET` with your own keys, which you can obtain from the [Dashboard](https://dashboard.plaid.com/developers/keys). These code samples also demonstrate starting up a server commonly used in each framework (such as Express or Flask).

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

#### Create an Item in Link

Plaid Link is a drop-in module that provides a secure, elegant authentication flow
for each institution that Plaid supports. Link makes it secure and easy for users to
connect their bank accounts to Plaid. Note that these instructions cover Link on the web. For instructions on using Link within mobile apps, see the [Link documentation](/docs/link/).

Using Link, we will create a Plaid *Item*, which is a Plaid term for a login at a financial institution. An Item is not the same as a financial institution account, although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access both their checking account and their savings account, a single Item would be associated with both of those accounts. If you want to customize Link's look and feel, you can do so from the [Dashboard](https://dashboard.plaid.com/link).

Before initializing Link, you will need to create a new `link_token` on the server side of your application.
A `link_token` is a short-lived, one-time use token that is used to authenticate your app with Link.
You can create one using the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint. Then, on the client side of your application, you'll need to initialize Link with the `link_token` that you just created.

##### Create a link\_token

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
    products: ['identity'],
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

##### Install Link dependency

index.html

```
<head>
  <title>Connect a bank</title>
  <script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
</head>
```

##### Configure the client-side Link handler

app.js

```
const linkHandler = Plaid.create({
  token: (await $.post('/create_link_token')).link_token,
  onSuccess: (public_token, metadata) => {
    // Send the public_token to your app server.
    $.post('/exchange_public_token', {
      public_token: public_token,
    });
  },
  onExit: (err, metadata) => {
    // Optionally capture when your user exited the Link flow.
    // Storing this information can be helpful for support.
  },
  onEvent: (eventName, metadata) => {
    // Optionally capture Link flow events, streamed through
    // this callback as your users connect an Item to Plaid.
  },
});

linkHandler.open();
```

#### Get a persistent access\_token

Next, on the server side, we need to exchange our `public_token` for an `access_token` and `item_id`. The `access_token` will allow us to make authenticated calls to the Plaid API. Doing so is as easy as calling the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) endpoint from our server-side handler. We'll use the client library we configured earlier to make the API call.

Save the `access_token` and `item_id` in a secure datastore, as they’re used to access `Item` data and identify `webhooks`, respectively. The `access_token` will remain valid unless you actively chose to expire it via rotation or remove the corresponding Item via [`/item/remove`](/docs/api/items/#itemremove). The `access_token` should be stored securely, and never in client-side code. A `public_token` is a one-time use token with a lifetime of 30 minutes, so there is no need to store it.

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

Now that the authentication step is out of the way, we can begin using authenticated endpoints from the Plaid API.

#### Matching Identity data

To match Identity data, call [`/identity/match`](/docs/api/products/identity/#identitymatch).

If you are using [Identity Verification](/docs/api/products/identity-verification/), you can automatically match data from the linked account against data collected during the Identity Verification flow. To enable this setting, from the Identity Verification section of the Dashboard, access the template editor and on the "Setup" pane of the template, check the box under the "Financial Account Matching" header. If this option is selected, you should call [`/identity/match`](/docs/api/products/identity/#identitymatch) with only an `access_token` to obtain match scores.

If you are not using [Identity Verification](/docs/api/products/identity-verification/), you will need to send the identity information that you have on file and would like to match against, such as name, phone number, and address, as part of your call to [`/identity/match`](/docs/api/products/identity/#identitymatch).

```
const request = {
  access_token: accessToken,
  // Omit user object if using Identity Verification / Financial Account Matching
  user: {
    legal_name: 'Jane Smith',
    phone_number: '+14155552671',
    email_address: 'jane.smith@example.com',
    address: {
      street: '500 Market St',
      city: 'San Francisco',
      region: 'CA',
      postal_code: '94105',
      country: 'US',
    },
  },
};

try {
  const response = await plaidClient.identityMatch(request);
  const accounts = response.data.accounts;
  for (const account of accounts) {
    const legalNameScore = account.legal_name?.score;
    const phoneScore = account.phone_number?.score;
    const emailScore = account.email_address?.score;
    const addressScore = account.address?.score;
  }
} catch (error) {
  // handle error
}
```

The call to [`/identity/match`](/docs/api/products/identity/#identitymatch) will return a match score for each field that was evaluated. Typically, your threshold to accept the field as a match should be set to at least 70. For more details, see the [match score table](/docs/identity/#sample-identity-match-data).

Identity Match sample response

```
{
  "accounts": [
    {
      "account_id": "BxBXxLj1m4HMXBm9WZZmCWVbPjX16EHwv99vp",
      "balances": {
        "available": null,
        "current": null,
        "iso_currency_code": null,
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "0000",
      "name": "Plaid Checking",
      "official_name": "Plaid Gold Standard 0% Interest Checking",
      "legal_name": {
        "score": 90,
        "is_nickname_match": true,
        "is_first_name_or_last_name_match": true,
        "is_business_name_detected": false
      },
      "phone_number": {
        "score": 100
      },
      "email_address": {
        "score": 100
      },
      "address": {
        "score": 100,
        "is_postal_code_match": true
      },
      "subtype": "checking",
      "type": "depository"
    },
    {
      "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "balances": {
        "available": null,
        "current": null,
        "iso_currency_code": null,
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "1111",
      "name": "Plaid Saving",
      "official_name": "Plaid Silver Standard 0.1% Interest Saving",
      "legal_name": {
        "score": 30,
        "is_first_name_or_last_name_match": false
      },
      "phone_number": {
        "score": 100
      },
      "email_address": null,
      "address": {
        "score": 100,
        "is_postal_code_match": true
      },
      "subtype": "savings",
      "type": "depository"
    }
  ],
  ...
}
```

#### Fetching Identity data

If you are not using Identity Match, call [`/identity/get`](/docs/api/products/identity/#identityget) to obtain Identity data. You will need to implement your own matching algorithm to determine whether the data returned matches the information that you have on file about the user. For more detailed information on the schema returned, see [`/identity/get`](/docs/api/products/identity/#identityget).

```
const { IdentityGetRequest } = require('plaid');

// Pull Identity data for an Item
const request: IdentityGetRequest = {
  access_token: accessToken,
};
try {
  const response = await plaidClient.identityGet(request);
  const identities = response.data.accounts.flatMap(
    (account) => account.owners,
  );
} catch (error) {
  // handle error
}
```

Example response data is below.

Identity sample data

```
{
  "accounts": [
    {
      "account_id": "BxBXxLj1m4HMXBm9WZZmCWVbPjX16EHwv99vp",
      "balances": {
        "available": 100,
        "current": 110,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "0000",
      "name": "Plaid Checking",
      "official_name": "Plaid Gold Standard 0% Interest Checking",
      "owners": [
        {
          "addresses": [
            {
              "data": {
                "city": "Malakoff",
                "country": "US",
                "postal_code": "14236",
                "region": "NY",
                "street": "2992 Cameron Road"
              },
              "primary": true
            },
            {
              "data": {
                "city": "San Matias",
                "country": "US",
                "postal_code": "93405-2255",
                "region": "CA",
                "street": "2493 Leisure Lane"
              },
              "primary": false
            }
          ],
          "emails": [
            {
              "data": "accountholder0@example.com",
              "primary": true,
              "type": "primary"
            },
            {
              "data": "accountholder1@example.com",
              "primary": false,
              "type": "secondary"
            }
          ],
          "names": ["Alberta Bobbeth Charleson"],
          "phone_numbers": [
            {
              "data": "1112223333",
              "primary": false,
              "type": "home"
            },
            {
              "data": "1112224444",
              "primary": false,
              "type": "work"
            },
            {
              "data": "1112225555",
              "primary": false,
              "type": "mobile1"
            }
          ]
        }
      ],
      "subtype": "checking",
      "type": "depository"
    }
  ],
  "item": {
    "available_products": ["balance", "credit_details", "investments"],
    "billed_products": [
      "assets",
      "auth",
      "identity",
      "liabilities",
      "transactions"
    ],
    "consent_expiration_time": null,
    "error": null,
    "institution_id": "ins_3",
    "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
    "webhook": "https://www.genericwebhookurl.com/webhook"
  },
  "request_id": "3nARps6TOYtbACO"
}
```

#### Tutorial and example code in Plaid Pattern

For a real-life example of an app that incorporates Identity, see the Node-based [Plaid Pattern Account Funding](https://github.com/plaid/pattern-account-funding) sample app. Pattern Account Funding is a sample account funding app that fetches Identity data in order verify identity prior to a funds transfer. The Identity code can be found in [items.js](https://github.com/plaid/pattern-account-funding/blob/master/server/routes/items.js#L81-L116).

For a tutorial walkthrough of creating a similar app, see [Account funding tutorial](https://github.com/plaid/account-funding-tutorial).

#### Next steps

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
