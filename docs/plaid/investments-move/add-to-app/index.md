---
title: "Investments Move - Add Investments Move to your app | Plaid Docs"
source_url: "https://plaid.com/docs/investments-move/add-to-app/"
scraped_at: "2026-03-07T22:04:59+00:00"
---

# Add Investments Move to your app

#### Use Investments Move to streamline brokerage-to-brokerage account transfers

In this guide, we'll start from scratch and walk through how to use Investments Move to get the data required to set up an ACATS transfer. If you are already familiar with using Plaid and are set up to make calls to the Plaid API, make sure to initialize Link with the `investments_auth` product; you can then skip ahead to [Fetching Investments Move data](/docs/investments-move/add-to-app/#fetching-investments-move-data).

#### Get Plaid API keys and complete application profile

If you don't already have one, you'll need to [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You will also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) on the Dashboard. The information in your profile will be shared with users of your application when they manage their connection on the [Plaid Portal](https://my.plaid.com). Your application profile must be completed before connecting to certain institutions in Production.

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

#### Create an Item in Link

Plaid Link is a drop-in module that provides a secure, elegant authentication flow
for each institution that Plaid supports. Link makes it secure and easy for users to
connect their bank accounts to Plaid. Note that these instructions cover Link on the web. For instructions on using Link within mobile apps, see the [Link documentation](/docs/link/).

Using Link, we will create a Plaid *Item*, which is a Plaid term for a login at a financial institution. An Item is not the same as a financial institution account, although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access both their checking account and their savings account, a single Item would be associated with both of those accounts.

First, on the client side of your application, you'll need to set up and configure Link. If you want to customize Link's look and feel, you can do so from the [Dashboard](https://dashboard.plaid.com/link).

When initializing Link, you will need to specify the products you will be using in the product array.

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
    products: ['investments_auth'],
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

When using Investments Move, you can also configure options in the [`/link/token/create`](/docs/api/link/#linktokencreate) call to allow more brokerage accounts to be added, with the tradeoff that Plaid may not be able to verify all of the information. For details, see [Fallback flows](/docs/investments-move/#fallback-flows).

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

#### Fetching Investments Move data

Now that the authentication step is out of the way, we can begin using authenticated endpoints from the Plaid API. For more detailed information on the schema for account information returned, see [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget).

```
const request: InvestmentsAuthGetRequest = {
  access_token: accessToken,
};
try {
  const response = await plaidClient.investmentsAuthGet(request);
  const investmentsAuthData = response.data;
} catch (error) {
  // handle error
}
```

The results of the [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) call return the information you need to submit an ACATS transfer, using verified data derived directly from the brokerage. Example response data is below. For more details on the schema of data returned, see the [API Reference](/docs/api/products/investments-move/#investments-auth-get-response-accounts).

Investments move sample response

```
{
  "accounts": [
    {
      "account_id": "31qEA6LPwGumkA4Z5mGbfyGwr4mL6nSZlQqpZ",
      "balances": {
        "available": 43200,
        "current": 43200,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "4444",
      "name": "Plaid Money Market",
      "official_name": "Plaid Platinum Standard 1.85% Interest Money Market",
      "subtype": "money market",
      "type": "depository"
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "balances": {
        "available": null,
        "current": 320.76,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "5555",
      "name": "Plaid IRA",
      "official_name": null,
      "subtype": "ira",
      "type": "investment"
    }
  ],
  "holdings": [
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "cost_basis": 1,
      "institution_price": 1,
      "institution_price_as_of": "2021-05-25",
      "institution_price_datetime": null,
      "institution_value": 0.01,
      "iso_currency_code": "USD",
      "quantity": 0.01,
      "security_id": "d6ePmbPxgWCWmMVv66q9iPV94n91vMtov5Are",
      "unofficial_currency_code": null,
      "vested_quantity": 1,
      "vested_value": 1
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "cost_basis": 0.01,
      "institution_price": 0.011,
      "institution_price_as_of": "2021-05-25",
      "institution_price_datetime": null,
      "institution_value": 110,
      "iso_currency_code": "USD",
      "quantity": 10000,
      "security_id": "8E4L9XLl6MudjEpwPAAgivmdZRdBPJuvMPlPb",
      "unofficial_currency_code": null,
      "vested_quantity": null,
      "vested_value": null
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "cost_basis": 40,
      "institution_price": 42.15,
      "institution_price_as_of": "2021-05-25",
      "institution_price_datetime": null,
      "institution_value": 210.75,
      "iso_currency_code": "USD",
      "quantity": 5,
      "security_id": "abJamDazkgfvBkVGgnnLUWXoxnomp5up8llg4",
      "unofficial_currency_code": null,
      "vested_quantity": 7,
      "vested_value": 66
    }
  ],
  "item": {
    "available_products": [
      "assets",
      "balance",
      "beacon",
      "cra_base_report",
      "cra_income_insights",
      "signal",
      "identity",
      "identity_match",
      "income",
      "income_verification",
      "investments",
      "processor_identity",
      "recurring_transactions",
      "transactions"
    ],
    "billed_products": [
      "investments_auth"
    ],
    "consent_expiration_time": null,
    "error": null,
    "institution_id": "ins_115616",
    "item_id": "7qBnDwLP3aIZkD7NKZ5ysk5X9xVxDWHg65oD5",
    "products": [
      "investments_auth"
    ],
    "update_type": "background",
    "webhook": "https://www.genericwebhookurl.com/webhook"
  },
  "numbers": {
    "acats": [
      {
        "account": "TR5555",
        "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
        "dtc_numbers": [
          "1111",
          "2222",
          "3333"
        ]
      }
    ]
  },
  "owners": [
    {
      "account_id": "31qEA6LPwGumkA4Z5mGbfyGwr4mL6nSZlQqpZ",
      "names": [
        "Alberta Bobbeth Charleson"
      ]
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "names": [
        "Alberta Bobbeth Charleson"
      ]
    }
  ],
  "request_id": "hPCXou4mm9Qwzzu",
  "securities": [
    {
      "close_price": 0.011,
      "close_price_as_of": null,
      "cusip": null,
      "industry": null,
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": false,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "Nflx Feb 01'18 $355 Call",
      "option_contract": null,
      "proxy_security_id": null,
      "sector": null,
      "security_id": "8E4L9XLl6MudjEpwPAAgivmdZRdBPJuvMPlPb",
      "sedol": null,
      "ticker_symbol": "NFLX180201C00355000",
      "type": "derivative",
      "unofficial_currency_code": null,
      "update_datetime": null
    },
    {
      "close_price": 9.08,
      "close_price_as_of": "2024-09-09",
      "cusip": null,
      "industry": "Investment Trusts or Mutual Funds",
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": false,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "DoubleLine Total Return Bond I",
      "option_contract": null,
      "proxy_security_id": null,
      "sector": "Miscellaneous",
      "security_id": "AE5rBXra1AuZLE34rkvvIyG8918m3wtRzElnJ",
      "sedol": "B5ND9B4",
      "ticker_symbol": "DBLTX",
      "type": "mutual fund",
      "unofficial_currency_code": null,
      "update_datetime": null
    },
    {
      "close_price": 42.15,
      "close_price_as_of": null,
      "cusip": null,
      "industry": null,
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": false,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "iShares Inc MSCI Brazil",
      "option_contract": null,
      "proxy_security_id": null,
      "sector": null,
      "security_id": "abJamDazkgfvBkVGgnnLUWXoxnomp5up8llg4",
      "sedol": null,
      "ticker_symbol": "EWZ",
      "type": "etf",
      "unofficial_currency_code": null,
      "update_datetime": null
    },
    {
      "close_price": 1,
      "close_price_as_of": null,
      "cusip": null,
      "industry": null,
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": true,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "U S Dollar",
      "option_contract": null,
      "proxy_security_id": null,
      "sector": null,
      "security_id": "d6ePmbPxgWCWmMVv66q9iPV94n91vMtov5Are",
      "sedol": null,
      "ticker_symbol": null,
      "type": "cash",
      "unofficial_currency_code": null,
      "update_datetime": null
    }
  ]
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
