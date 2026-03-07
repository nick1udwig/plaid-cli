---
title: "Assets - Create an Asset Report | Plaid Docs"
source_url: "https://plaid.com/docs/assets/add-to-app/"
scraped_at: "2026-03-07T22:04:29+00:00"
---

# Create an Asset Report

#### Learn how to create Asset Reports with the Assets product

In this guide, we'll start from scratch and walk through how to use [Assets](/docs/api/products/assets/) to generate and retrieve Asset Reports. If a user's Asset Report involves data from multiple financial institutions, the user will need to allow access to each institution, which will in turn enable you to access their data from each institution. If you are already familiar with using Plaid and are set up to make calls to the Plaid API, you can skip ahead to [Creating Asset Reports](/docs/assets/add-to-app/#creating-asset-reports).

#### Get Plaid API keys and complete application and company profile

If you don't already have one, you'll need to [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You will also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile) in the Dashboard. This will provide basic information about your app to help users manage their connection on the Plaid Portal at [my.plaid.com](https://my.plaid.com). The application profile and company profile must be completed before connecting to certain institutions in Production.

#### Install and initialize Plaid libraries

You can use our official server-side client libraries to connect to the Plaid API from your application:

Terminal

```
// Install via npm
npm install --save plaid
```

After you've installed Plaid's client libraries, you can initialize them by passing in your `client_id`, `secret`, and the environment you wish to connect to (Sandbox or Production). This will make sure the client libraries pass along your `client_id` and `secret` with each request, and you won't need to explicitly include it in any other calls.

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

Using Link, we will create a Plaid *Item*, which is a Plaid term for a login at a financial institution. An Item is not the same as a financial institution account, although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access both their checking account and their savings account, a single Item would be associated with both of those accounts. Asset Reports can consist of user data from multiple financial institutions; users will need to use Link to provide access to each financial institution, providing you with multiple Items. If you want to customize Link's look and feel, you can do so from the [Dashboard](https://dashboard.plaid.com/link).

Before initializing Link, you will need to create a new `link_token` on the server side of your application.
A `link_token` is a short-lived, one-time use token that is used to authenticate your app with Link.
You can create one using the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint. Then, on the client side of your application, you'll need to initialize Link with the `link_token` that you just created. Since, for Assets, users may need to grant access to more than one financial institution via Link, you may need to initialize Link more than once.

In the code samples below, you will need to replace `PLAID_CLIENT_ID` and `PLAID_SECRET` with your own keys, which you can obtain from the [Dashboard](https://dashboard.plaid.com/developers/keys).

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
    products: ['assets'],
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

Next, on the server side, we need to exchange our `public_token` for an `access_token` and `item_id` for each Item the user provided you with via Link. The `access_token` will allow us to make authenticated calls to the Plaid API for its corresponding financial institution. Doing so is as easy as calling the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) endpoint from our server-side handler. We'll use the client library we configured earlier to make the API call.

Save the `access_token`s and `item_id`s in a secure datastore, as they’re used to access `Item` data and identify `webhooks`, respectively. An `access_token` will remain valid unless you actively chose to expire it via rotation or remove the corresponding Item via [`/item/remove`](/docs/api/items/#itemremove). An `access_token` should be stored securely and never in client-side code. A `public_token` is a one-time use token with a lifetime of 30 minutes, so there is no need to store it.

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

#### Creating Asset Reports

Now that the authentication step is out of the way, we can begin using authenticated endpoints from the Plaid API and create an Asset Report using the [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) endpoint.

##### Assets sample request: create

```
const { AssetReportCreateRequest } = require('plaid');

const daysRequested = 90;
const options = {
  client_report_id: '123',
  webhook: 'https://www.example.com',
  user: {
    client_user_id: '789',
    first_name: 'Jane',
    middle_name: 'Leah',
    last_name: 'Doe',
    ssn: '123-45-6789',
    phone_number: '(555) 123-4567',
    email: 'jane.doe@example.com',
  },
};
const request: AssetReportCreateRequest = {
  access_tokens: [accessToken],
  days_requested,
  options,
};
// accessTokens is an array of Item access tokens.
// Note that the assets product must be enabled for all Items.
// All fields on the options object are optional.
try {
  const response = await plaidClient.assetReportCreate(request);
  const assetReportId = response.data.asset_report_id;
  const assetReportToken = response.data.asset_report_token;
} catch (error) {
  // handle error
}
```

Sample response data is below.

Assets response: create

```
{
  "asset_report_token": "assets-sandbox-6f12f5bb-22dd-4855-b918-f47ec439198a",
  "asset_report_id": "1f414183-220c-44f5-b0c8-bc0e6d4053bb",
  "request_id": "Iam3b"
}
```

#### Fetching asset data

Once an Asset Report has been created, it can be retrieved to analyze the user's loan eligibility. For more detailed information on the schema for Asset Reports, see [`/asset_report/get`](/docs/api/products/assets/#asset_reportget).

Asset Reports are not generated instantly. If you receive a `PRODUCT_NOT_READY` error when calling [`/asset_report/get`](/docs/api/products/assets/#asset_reportget), the requested Asset Report has not yet been generated. To be alerted when the requested Asset Report has been generated, listen to [Assets webhooks](/docs/api/products/assets/#product_ready).

##### Assets sample request: get

```
const { AssetReportGetRequest } = require('plaid');

const request: AssetReportGetRequest = {
  asset_report_token: assetReportToken,
  include_insights: true,
};
try {
  const response = await plaidClient.assetReportGet(request);
  const assetReportId = response.data.asset_report_id;
} catch (error) {
  // handle error
}
```

Sample response data is below.

Assets response: get

```
{
  "report": {
    "asset_report_id": "bf3a0490-344c-4620-a219-2693162e4b1d",
    "client_report_id": "123abc",
    "date_generated": "2020-06-05T22:47:53Z",
    "days_requested": 2,
    "items": [
      {
        "accounts": [
          {
            "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
            "balances": {
              "available": 200,
              "current": 210,
              "iso_currency_code": "USD",
              "limit": null,
              "unofficial_currency_code": null
            },
            "days_available": 2,
            "historical_balances": [
              {
                "current": 210,
                "date": "2020-06-04",
                "iso_currency_code": "USD",
                "unofficial_currency_code": null
              },
              {
                "current": 210,
                "date": "2020-06-03",
                "iso_currency_code": "USD",
                "unofficial_currency_code": null
              }
            ],
            "mask": "1111",
            "name": "Plaid Saving",
            "official_name": "Plaid Silver Standard 0.1% Interest Saving",
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
                    "data": "extraordinarily.long.email.username.123456@reallylonghostname.com",
                    "primary": false,
                    "type": "other"
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
                    "data": "1112225555",
                    "primary": false,
                    "type": "mobile1"
                  }
                ]
              }
            ],
            "ownership_type": null,
            "subtype": "savings",
            "transactions": [],
            "type": "depository"
          },
          {
            "account_id": "BxBXxLj1m4HMXBm9WZJyUg9XLd4rKEhw8Pb1J",
            "balances": {
              "available": null,
              "current": 56302.06,
              "iso_currency_code": "USD",
              "limit": null,
              "unofficial_currency_code": null
            },
            "days_available": 2,
            "historical_balances": [],
            "mask": "8888",
            "name": "Plaid Mortgage",
            "official_name": null,
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
                    "data": "extraordinarily.long.email.username.123456@reallylonghostname.com",
                    "primary": false,
                    "type": "other"
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
                    "data": "1112225555",
                    "primary": false,
                    "type": "mobile1"
                  }
                ]
              }
            ],
            "ownership_type": null,
            "subtype": "mortgage",
            "transactions": [],
            "type": "loan"
          },
          {
            "account_id": "dVzbVMLjrxTnLjX4G66XUp5GLklm4oiZy88yK",
            "balances": {
              "available": null,
              "current": 410,
              "iso_currency_code": "USD",
              "limit": null,
              "unofficial_currency_code": null
            },
            "days_available": 2,
            "historical_balances": [
              {
                "current": 410,
                "date": "2020-06-04",
                "iso_currency_code": "USD",
                "unofficial_currency_code": null
              },
              {
                "current": 410,
                "date": "2020-06-03",
                "iso_currency_code": "USD",
                "unofficial_currency_code": null
              }
            ],
            "mask": "3333",
            "name": "Plaid Credit Card",
            "official_name": "Plaid Diamond 12.5% APR Interest Credit Card",
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
                    "data": "extraordinarily.long.email.username.123456@reallylonghostname.com",
                    "primary": false,
                    "type": "other"
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
                    "data": "1112225555",
                    "primary": false,
                    "type": "mobile1"
                  }
                ]
              }
            ],
            "ownership_type": null,
            "subtype": "credit card",
            "transactions": [],
            "type": "credit"
          }
        ],
        "date_last_updated": "2020-06-05T22:47:52Z",
        "institution_id": "ins_3",
        "institution_name": "Chase",
        "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6"
      }
    ],
    "user": {
      "client_user_id": "123456789",
      "email": "accountholder0@example.com",
      "first_name": "Alberta",
      "last_name": "Charleson",
      "middle_name": "Bobbeth",
      "phone_number": "111-222-3333",
      "ssn": "123-45-6789"
    }
  },
  "request_id": "eYupqX1mZkEuQRx",
  "warnings": []
}
```

#### Working with Assets data

After your initial [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) and [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) requests, you may want your application to fetch updated Asset Reports, provide Audit Copies of Asset Reports, and more. Consult the [API Reference](/docs/api/products/assets/) to explore these and other options.

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
