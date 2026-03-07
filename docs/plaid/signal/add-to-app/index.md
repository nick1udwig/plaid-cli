---
title: "Signal - Add Signal Transaction Scores to your app | Plaid Docs"
source_url: "https://plaid.com/docs/signal/add-to-app/"
scraped_at: "2026-03-07T22:05:18+00:00"
---

# Add Signal Transaction Scores to your app

#### Learn how to add Signal Transaction Scores to your application

In this guide, we'll start from scratch and walk through how to use [Signal Transaction Scores](/docs/api/products/signal/) to perform risk analysis on proposed ACH transactions. If you are already familiar with using Plaid and are set up to make calls to the Plaid API, you can skip ahead to [creating rulesets](/docs/signal/add-to-app/#creating-rulesets).

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

Put `signal` in the `products` array when calling [`/link/token/create`](/docs/api/link/#linktokencreate).

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
    products: ['auth, signal'],
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

Save the `access_token` and `item_id` in a secure datastore, as they’re used to access Item data and identify webhooks, respectively. The `access_token` will remain valid unless you actively chose to expire it via rotation or remove the corresponding Item via [`/item/remove`](/docs/api/items/#itemremove). The `access_token` should be stored securely, and never in client-side code. A `public_token` is a one-time use token with a lifetime of 30 minutes, so there is no need to store it.

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

#### Adding Signal Transaction Scores to existing Items

If your flow involves adding Signal Transaction Scores to existing Items that weren't initialized with `signal` in the [`/link/token/create`](/docs/api/link/#linktokencreate) flow, call [`/signal/prepare`](/docs/api/products/signal/#signalprepare) on those Items. For more details, see [Creating Items](/docs/signal/creating-signal-items/#adding-signal-to-existing-items).

/signal/prepare

```
const prepare_request = {
  access_token: 'access-sandbox-71e02f71-0960-4a27-abd2-5631e04f2175',
};

try {
  const prepare_response = await plaidClient.signalPrepare(prepare_request);
  const request_id = prepare_response.data.request_id;
  console.log(request_id);
} catch (error) {
  // handle error
}
```

#### Creating rulesets

Now that the authentication and initialization step is out of the way, we can begin using Signal Transaction Scores to analyze proposed transactions. Go to the [Signal Rules Dashboard](https://dashboard.plaid.com/signal/risk-profiles) to create a ruleset. For details, see [Signal Rules](/docs/signal/signal-rules/).

Once you've created your ruleset, you can evaluate transactions with [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate).

#### Evaluating a proposed transaction

/signal/evaluate

```
const eval_request = {
  access_token: 'access-sandbox-71e02f71-0960-4a27-abd2-5631e04f2175',
  account_id: '3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr',
  client_transaction_id: 'txn12345',
  amount: 123.45,
  client_user_id: 'user1234',
  user: {
    name: {
      prefix: 'Ms.',
      given_name: 'Jane',
      middle_name: 'Leah',
      family_name: 'Doe',
      suffix: 'Jr.',
    },
    phone_number: '+14152223333',
    email_address: 'jane.doe@example.com',
    address: {
      street: '2493 Leisure Lane',
      city: 'San Matias',
      region: 'CA',
      postal_code: '93405-2255',
      country: 'US',
    },
  },
  device: {
    ip_address: '198.30.2.2',
    user_agent:
      'Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1',
  },
  user_present: true,
  ruleset_key: "recommended-risk-rules"
};

try {
  const eval_response = await plaidClient.signalEvaluate(eval_request);
  const result = eval_response.data.ruleset.result;
} catch (error) {
  // handle error
}
```

Based on the `result` of either `accept`, `review`, or `reroute`, proceed with the transfer, review it, or reroute the user to a different payment method.

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

#### Iterate on Signal Rules

Once you have have amassed enough transaction history in Production (a few hundred transactions), you should periodically review the Signal Performance Dashboard and tweak your rulesets as necessary. For more details, see [Tuning your rules](/docs/signal/signal-rules/#tuning-your-rules).

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
