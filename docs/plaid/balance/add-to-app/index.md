---
title: "Balance - Add Balance to your app | Plaid Docs"
source_url: "https://plaid.com/docs/balance/add-to-app/"
scraped_at: "2026-03-07T22:04:46+00:00"
---

# Add Balance to your app

#### Use Balance to fetch real-time balance data

In this guide, we'll start from scratch and walk through how to use Balance to retrieve real-time balance information and assess ACH return risk. If you are already familiar with using Plaid and are set up to make calls to the Plaid API, make sure to note that you should *not* include `balance` in the [`/link/token/create`](/docs/api/link/#linktokencreate) products array, but you *should* include `signal`, if you are using Balance for an ACH risk assessment use case; you can then skip ahead to [Fetching balance data](/docs/balance/add-to-app/#fetching-balance-data).

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

#### Create an Item in Link

Plaid Link is a drop-in module that provides a secure, elegant authentication flow
for each institution that Plaid supports. Link makes it secure and easy for users to
connect their bank accounts to Plaid. Note that these instructions cover Link on the web. For instructions on using Link within mobile apps, see the [Link documentation](/docs/link/).

Using Link, we will create a Plaid *Item*, which is a Plaid term for a login at a financial institution. An Item is not the same as a financial institution account, although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access both their checking account and their savings account, a single Item would be associated with both of those accounts. If you want to customize Link's look and feel, you can do so from the [Dashboard](https://dashboard.plaid.com/link).

If you are using Balance for a payments use case and you do not set the [Link Account Select UI](https://dashboard.plaid.com/link/account-select) to "enabled for one account", your UI flow when creating a payment must handle the scenario of an end user linking an Item with multiple accounts and allow them to specify which account to use for the payment.

In order to launch Link in Production, you must [select use cases for your Link customization](https://dashboard.plaid.com/link/data-transparency-v5). This requirement is not enforced in Sandbox.

Before initializing Link, you will need to create a new `link_token` on the server side of your application.
A `link_token` is a short-lived, one-time use token that is used to authenticate your app with Link.
You can create one using the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint. Then, on the client side of your application, you'll need to initialize Link with the `link_token` that you just created.

If using Balance for ACH return risk assessment, include `signal` in the `products` array, along with all other Plaid product(s) you will be using with Balance.

For all other Balance use cases, omit `signal`, but instead include the Plaid product(s) you will be using with Balance in the `products` array.

`balance` cannot be included in the `products` array when calling [`/link/token/create`](/docs/api/link/#linktokencreate).

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
    products: ['signal, auth'],
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

#### Fetching Balance data

Now that the authentication step is out of the way, we can begin using authenticated endpoints from the Plaid API.

There are two ways to use Balance: by using [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), or by using [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget).

If you are using Balance to evaluate a proposed ACH transaction for return risk, use [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate).

For all other use cases, use [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget).

##### Using /signal/evaluate

###### Creating a Balance-only ruleset

In the Dashboard, navigate to [**Signal->Rules**](https://dashboard.plaid.com/signal/risk-profiles) to create a Balance-only ruleset. When you do, Plaid will pre-populate a suggested ruleset; you can either use it as-is or customize it.

In Sandbox, you will be offered a choice between Balance-only and Signal Transaction Score-powered rulesets. To use Balance, select "Balance-only". In Production, Signal Transaction Score rulesets will only be provided as an option if you are Production-enabled for the Signal Transaction Scores product.

###### Getting the account\_id

If you are using [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), you will need to know the `account_id` of the account that is being debited. You can get this in multiple ways, the simplest of which are: by using the `onSuccess` callback in Link (`metadata.accounts[].id`) or by calling [`/accounts/get`](/docs/api/accounts/#accountsget).

Method 1: The metadata object from Link onSuccess callback

```
{
  ...
  "accounts": [
    {
      "id": "ygPnJweommTWNr9doD6ZfGR6GGVQy7fyREmWy",
      "name": "Plaid Checking",
      "mask": "0000",
      "type": "depository",
      "subtype": "checking",
      "verification_status": null
    },
    {
      "id": "9ebEyJAl33FRrZNLBG8ECxD9xxpwWnuRNZ1V4",
      "name": "Plaid Saving",
      "mask": "1111",
      "type": "depository",
      "subtype": "savings"
    }
  ],
  ...
}
```

Method 2: Calling /accounts/get

```
const request: AccountsGetRequest = {
  access_token: ACCESS_TOKEN,
};
try {
  const response = await plaidClient.accountsGet(request);
  const accounts = response.data.accounts;
} catch (error) {
  // handle error
}
```

A sample response is below (note that [`/accounts/get`](/docs/api/accounts/#accountsget) returns balance data, but it is cached and not updated in real time):

/accounts/get response

```
{
  "accounts": [
    {
      "account_id": "blgvvBlXw3cq5GMPwqB6s6q4dLKB9WcVqGDGo",
      "balances": {
        "available": 100, 
        "current": 110,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "holder_category": "personal",
      "mask": "0000",
      "name": "Plaid Checking",
      "official_name": "Plaid Gold Standard 0% Interest Checking",
      "subtype": "checking",
      "type": "depository"
    },
    {
      "account_id": "6PdjjRP6LmugpBy5NgQvUqpRXMWxzktg3rwrk",
      "balances": {
        "available": null,
        "current": 23631.9805,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "6666",
      "name": "Plaid 401k",
      "official_name": null,
      "subtype": "401k",
      "type": "investment"
    }
  ],
  ...
```

Once you have obtained both an `access_token` and an `account_id`, you can call [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) to evaluate the proposed transaction. If you do not specify a `ruleset_key`, the transaction will be evaluated based on the ruleset named `default`.

/signal/evaluate

```
const eval_request = {
  access_token: 'access-sandbox-71e02f71-0960-4a27-abd2-5631e04f2175',
  account_id: '3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr',
  client_transaction_id: 'txn12345',
  amount: 123.45,
}

try {
  const eval_response = await plaidClient.signalEvaluate(eval_request);
  const core_attributes = eval_response.data.core_attributes;
} catch (error) {
  // handle error
}
```

To determine which action to take, use the `ruleset.result` value of `accept`, `reroute`, or `review`. To learn more about these values, see [Signal rules](/docs/signal/signal-rules/#using-signal-ruleset-results).

You can also view the real-time current and available balances in `core_attributes.current_balance` and `core_attributes.available_balance`.

#### Reporting outcomes

If the result was `review`, report your final review decision to Plaid using [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) or by uploading a CSV. For more details, see [Reporting decisions and returns](/docs/signal/reporting-returns/).

/signal/decision/report

```
const decision_report_request = {
  client_transaction_id: 'txn12345',
  initiated: true,
  days_funds_on_hold: 3,
};

try {
  const decision_report_response = await plaidClient.signalDecisionReport(
    decision_report_request,
  );
  const decision_request_id = decision_report_response.data.request_id;
} catch (error) {
  // handle error
}
```

If you allow a transfer that does end up returned, report that result back to Plaid as well. You can do this by calling [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport) or by uploading a CSV. For more details, see [Reporting decisions and returns](/docs/signal/reporting-returns/).

/signal/return/report

```
const return_report_request = {
  client_transaction_id: 'txn12345',
  return_code: 'R01',
};

try {
  const return_report_response = await plaidClient.signalReturnReport(
    return_report_request,
  );
  const request_id = return_report_response.data.request_id;
  console.log(request_id);
} catch (error) {
  // handle error
}
```

##### Using /accounts/balance/get

For more detailed information on the schema returned, see [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget).

/accounts/balance/get

```
const { AccountsGetRequest } = require('plaid');

// Pull real-time balance information for each account associated
// with the Item
const request: AccountsGetRequest = {
  access_token: accessToken,
};
try {
  const response = await plaidClient.accountsBalanceGet(request);
  const accounts = response.data.accounts;
} catch (error) {
  // handle error
}
```

Example response data is below.

Balance sample response

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
      "subtype": "checking",
      "type": "depository"
    },
    {
      "account_id": "dVzbVMLjrxTnLjX4G66XUp5GLklm4oiZy88yK",
      "balances": {
        "available": null,
        "current": 410,
        "iso_currency_code": "USD",
        "limit": 2000,
        "unofficial_currency_code": null
      },
      "mask": "3333",
      "name": "Plaid Credit Card",
      "official_name": "Plaid Diamond 12.5% APR Interest Credit Card",
      "subtype": "credit card",
      "type": "credit"
    },
    {
      "account_id": "Pp1Vpkl9w8sajvK6oEEKtr7vZxBnGpf7LxxLE",
      "balances": {
        "available": null,
        "current": 65262,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "7777",
      "name": "Plaid Student Loan",
      "official_name": null,
      "subtype": "student",
      "type": "loan"
    }
  ],
  "item": {
    "available_products": [
      "balance",
      "credit_details",
      "identity",
      "investments"
    ],
    "billed_products": ["assets", "auth", "liabilities", "transactions"],
    "consent_expiration_time": null,
    "error": null,
    "institution_id": "ins_3",
    "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
    "webhook": "https://www.genericwebhookurl.com/webhook"
  },
  "request_id": "qk5Bxes3gDfv4F2"
}
```

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
