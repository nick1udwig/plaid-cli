---
title: "Income - Add Income to your app | Plaid Docs"
source_url: "https://plaid.com/docs/income/add-to-app/"
scraped_at: "2026-03-07T22:04:58+00:00"
---

# Add Income to your app

#### Use Income to fetch income information about your users

This guide will walk you through how to use [Income](/docs/api/products/income/) to retrieve information about your users' current sources of income.

This is a basic quickstart guide that does not cover all features of the Income product. For more complete information about Income endpoints and capabilities, see the individual product pages for [Bank Income](/docs/income/bank-income/), [Document Income](/docs/income/document-income/), and [Payroll Income](/docs/income/payroll-income/).

#### Get Plaid API keys and complete application and company profile

If you don't already have one, [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile) in the Dashboard if you wish to retrieve real data in Production. The information in your profile will be shared with users of your application when they manage their connection on the [Plaid Portal](https://my.plaid.com), and must be completed before connecting to certain institutions.

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

#### Create a User Token

Unlike many other Plaid products, where you start by creating `Items` to represent user connections with individual banks, with Income you first start by creating a `user_token` that represents an individual user. As you add various sources of income, those will all be associated with this single `user_token`, which allows you to receive consolidated income from different sources by requesting an income report for that `user_token`.

To create a `user_token`, make a call to the [`/user/create`](/docs/api/users/#usercreate) endpoint. This endpoint requires a unique `client_user_id` value to represent the current user. Typically, this value is the same identifier you're using in your own application to represent the currently signed-in user. This call will return a randomly-generated string for the `user_token` as well as a separate user ID that Plaid includes in webhooks to represent this user.

You can only create one `user_token` per `client_user_id`. If you try to create a `user_token` with a `client_user_id` that you have previously sent to Plaid, you will receive an error.

Depending on your application, you might wish to create this `user_token` as soon as your user creates an account, or right before they begin the process of confirming their income data with you.

server.js

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

// The userId here represents the user's internal user ID with your
// application
const createUserToken = async (userId) => {
  const response = await plaidClient.userCreate({
    client_user_id: userId,
  });
  const newUserToken = response.data.user_token;
  const userIdForWebhooks = response.data.user_id;
  await updateUserRecordWithIncomeInfo(userId, newUserToken, userIdForWebhooks);
};
```

#### Decide what income verification method(s) you want to support

Income supports three different ways that a user can verify their sources of income.

- *Payroll Income* allows Plaid to retrieve information such as recent pay stubs and W-2 forms directly from the user's payroll provider. This tends to provide the most comprehensive data, but is dependent upon the user being able to sign in with their payroll provider.
- *Document Income* allows the user to upload images of income-related documents (such as pay stubs and W-2 forms), which Plaid can scan and interpret for you.
- *Bank Income* allows the user to connect with their bank to identify deposits that represent sources of income.

Depending on which methods you want to support (and it's perfectly acceptable to support all three), you will need to implement slightly different processes. While you can collect Payroll and Document Income in the same Link flow, you cannot collect Payroll or Document Income in the same Link flow as Bank Income.

### Fetching Payroll or Document Income data

The process for implementing Payroll and Document Income is very similar, so we'll cover these two scenarios next.

#### Create a Link token for Payroll or Document Income

Plaid Link is a drop-in module that provides a secure, elegant UI flow to guide your users through the process of connecting Plaid with their financial institutions. With Income, Link can help your user search for and connect to their payroll provider as well as assist them in uploading documents.

Note that these instructions cover Link on the web. For instructions on using Link within mobile apps, see the appropriate section in the [Link documentation](/docs/link/).

Before initializing Link, you will need to create a new `link_token` on the server side of your application. A `link_token` is a short-lived, single-use token that is used to authenticate your app with Link. You can create one using the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint. Then, on the client side of your application, initialize Link with the `link_token` that you just created.

When calling the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint, you'll include an object telling Plaid about your application as well as how to customize the Link experience. When creating a `link_token` for Payroll or Document Income, there are a few extra values you'll need to supply in this object:

- You will need to pass in the `user_token` that you created earlier.
- You can only list `"income_verification"` in the list of products
- You will need to create an object with `income_source_types` set to `["payroll"]` and pass that as the value for `income_verification`

Note that in the code sample below, in addition to passing in a `user_token`, we are also passing in the signed in user's ID as the `user.client_user_id` value. These are fundamentally different user IDs used for different purposes. The `user.client_user_id` value is primarily used for logging purposes and allows you to search for activity in the Plaid Dashboard based on the user ID. While its value is most likely the same user ID that you passed in when creating the `user_token` above, they are otherwise unrelated.

server.js

```
// Using Express

app.post('/api/create_link_token_for_payroll_income', async function (
  request,
  response,
) {
  // Get the client_user_id by searching for the current user.
  const clientUserId = await GetSignedInUserId();
  // Get the Plaid user token that we stored in an earlier step
  const userToken = await GetPlaidTokenForUser();
  const configs = {
    user: {
      // This should correspond to a unique id for the current user.
      client_user_id: clientUserId,
    },
    client_name: 'Plaid Test App',
    products: ['income_verification'],
    user_token: userToken,
    income_verification: {
      income_source_types: ['payroll'],
    },
    language: 'en',
    webhook: 'https://webhook.example.com',
    country_codes: ['US'],
  };
  try {
    const createTokenResponse = await client.linkTokenCreate(configs);
    response.json(createTokenResponse.data);
  } catch (error) {
    // handle error
  }
});
```

By default, a `link_token` created this way will allow the user to collect income information using both the Payroll and Document methods. However, in some cases you might want to permit the user to only use one of these methods. To do that, you can specify the `flow_types`, as in the example below.

```
income_verification: {
  income_source_types: ["payroll"],
  payroll_income: { flow_types: ["payroll_digital_income"] },
},
```

##### Install the Link library

You need to install the Link JavaScript library from Plaid in order to use Link in your web application.

index.html

```
<head>
  <title>Connect a bank</title>
  <script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
</head>
```

##### Configure the client-side Link handler

To run Link, you'll first need to retrieve the `link_token` from the server using the call to [`/link/token/create`](/docs/api/link/#linktokencreate) described above. Then, call `Plaid.create()`, passing along that token alongside other callbacks you might want to define. This will return a handler that has an `open` method. Call that method to open up the Link UI and start the connection process.

app.js

```
const linkHandler = Plaid.create({
  token: (await $.post('/api/create_link_token_for_payroll_income')).link_token,
  onSuccess: (public_token, metadata) => {
    // Typically, you'd exchange the public_token for an access token.
    // While you can still do that here, it's not strictly necessary.
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

If you've had experience with other Plaid products, you're probably used to taking the `public_token` received from Link and exchanging it on your server for a persistent `access_token`. However, because Payroll and Document Income are fetched once (and don't require persistent connections with your user's bank), and all of your reports are associated with the user's `user_token`, there's no need to retrieve or store an `access_token` in this case.

#### Listen for webhooks

After a user has finished using Link, it may take some time before Plaid has the user's income information available for you to download. In the case of Payroll Income, this typically takes a few seconds. For Document Income, this may take several minutes. Listen for the `INCOME: INCOME_VERIFICATION` webhook to know when this user's data is ready.

webhookServer.js

```
app.post('/server/receive_webhook', async (req, res, next) => {
  const product = req.body.webhook_type;
  const code = req.body.webhook_code;
  if (product === 'INCOME' && code === 'INCOME_VERIFICATION') {
    const plaidUserId = req.body.user_id;
    const verificationStatus = req.body.verification_status;
    if (verificationStatus === 'VERIFICATION_STATUS_PROCESSING_COMPLETE') {
      await retrieveIncomeDataForUser(plaidUserId);
    } else {
      // Handle other cases
    }
  }
  // Handle other types of webhooks
});
```

The `user_id` value included in this webhook is the same `user_id` that was returned by Plaid when you first created a `user_token` for this user.

#### Fetch income data

Now that Plaid has confirmed it has finished processing this user's data, you can fetch the data by making a call to the [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) endpoint. Pass in the `user_token`, and Plaid will return all Payroll and Document Income sources associated with that token.

server.js

```
try {
  const response = await plaidClient.creditPayrollIncomeGet({
    user_token: userToken,
  });
  const incomeData = response.data;
  // Do something interesting with the income data here.
} catch (error) {
  // Handle error
}
```

This call will typically return an array of `items`, each of which represents a connection with a payroll provider. These `items`, in turn, contain one or more `payroll_income` objects that contain detailed payroll information for the user. This can include individual pay stubs (with a full breakdown of gross pay, deductions, net pay, and more), as well as W-2 forms and all of their information. See the [sample response object](/docs/api/products/income/#creditpayroll_incomeget) in the documentation for a full example of what this response might look like.

You can distinguish between income that was retrieved through Payroll Income and income that was retrieved through Document Income by looking for an `account_id` value on the `payroll_income` object. If the data was retrieved through scanning user documents, this value will be null.

### Fetching Bank Income data

As an alternative to downloading income data from the user's payroll provider or W-2 forms, Plaid can also find income data in the user's bank account. Plaid will look for deposits that might be sources of income and ask the user to identify which deposits represent income that they wish to share with your application.

#### Create a Link token for Bank Income

Creating a `link_token` for Bank Income is similar to the process for creating a `link_token` for Payroll or Document Income. The main differences are:

- While you still must include `"income_verification"` as a product, you are allowed to combine this product with other products that you might want to use with this financial institution.
- Your `income_source_types` must be set to `["bank"]`.
- Your `income_verification` object must also include a `bank_income` object, where you have set the value of `days_requested`. This is the number of days back that Plaid will search for regularly occurring deposits.

The value of `days_requested` should a large enough number that you find all relevant deposits, but not so large that the user gets tired waiting for Plaid to retrieve all the relevant data. We recommend 120 days as good balance between the two.

server.js

```
app.post('/api/create_link_token_for_bank_income', async function (
  request,
  response,
) {
  // Get the client_user_id by searching for the current user.
  const clientUserId = await GetSignedInUserId();
  // Get the Plaid user token that we stored in an earlier step
  const userToken = await GetPlaidTokenForUser();
  const configs = {
    user: {
      // This should correspond to a unique id for the current user.
      client_user_id: clientUserId,
    },
    client_name: 'Plaid Test App',
    products: ['income_verification', 'transactions', 'assets'],
    user_token: userToken,
    income_verification: {
      income_source_types: ['bank'],
      bank_income: { days_requested: 120 },
    },
    language: 'en',
    webhook: 'https://webhook.example.com',
    country_codes: ['US'],
  };
  try {
    const createTokenResponse = await client.linkTokenCreate(configs);
    response.json(createTokenResponse.data);
  } catch (error) {
    // handle error
  }
});
```

##### Install the Link library

As described earlier, you need to install the Link JavaScript library from Plaid to use Link in your web application.

index.html

```
<head>
  <title>Connect a bank</title>
  <script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
</head>
```

##### Configure the client-side Link handler

Retrieve the `link_token` from the server, then call `Plaid.create()`, passing along the `link_token` alongside other callbacks you might want to define. This will return a handler that has an `open` method, which you can call to start the Link process. This is similar to the process described above, but with Bank Income you will probably want to keep the `public_token` and exchange it for an `access_token`, as described in the next step.

app.js

```
const linkHandler = Plaid.create({
  token: (await $.post('/api/create_link_token_for_bank_income')).link_token,
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

##### Retrieve metadata about the Link session

Call [`/link/token/get`](/docs/api/link/#linktokenget) in your server after receiving `onSuccess`, passing in the `link_token`. If you enabled [Multi-Item Link](/docs/link/multi-item-link/) when calling [`/link/token/create`](/docs/api/link/#linktokencreate), you should instead wait for the [`SESSION_FINISHED`](/docs/api/link/#session_finished) webhook rather than the `onSuccess` frontend callback, to ensure that the Link flow is complete, and all Items have been linked. For more details, see the [Multi-Item Link docs](/docs/link/multi-item-link/).

Previously, customers were instructed to call [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) to obtain the public token. While this flow is still supported, all customers are encouraged to use the [`/link/token/get`](/docs/api/link/#linktokenget) endpoint to obtain the public token instead, for improved consistency with other Plaid products and flows.

server.js

```
try {
  const response = await plaidClient.linkTokenGet({
    link_token: LinkToken,
  });
  const sessionData = response.data;
} catch (error) {
  // Handle this error
}
```

If you are using other Plaid products such as Auth or Balance alongside Bank Income, make sure to capture the `public_token` (or tokens) from the `results.item_add_result` field and exchange it for an `access_token`.

#### Get a persistent access token (for other products)

For the purpose of retrieving income data, you don't necessarily need an `access_token`. Plaid will retrieve the income data for you, and then that data will be associated with the `user_token` that you created when first configuring Link.

However, if you are using any other product with this financial institution (such as Transactions, Balance, or Assets), you will need to exchange your `public_token` for an `access_token` and `item_id`. The `access_token` lets you make authenticated calls to these other products in the Plaid API.

Exchanging a `public_token` for an `access_token` is as easy as calling the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) endpoint from the server-side handler. Store the `access_token` and `item_id` that get returned in a secure server-side location. Never store this information on the client.

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

#### Fetch income data

You can fetch income data by calling the [`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) endpoint using the Plaid client library, passing in the `user_token` you created earlier. Note that, similar to Document or Payroll Income, this data is static -- it is a snapshot of the user's income data that was retrieved by Plaid when the user ran Link. It does not update or change over time.

By default, Plaid will only return the most recent report in its response. However, you might want to retrieve more than one report if your user has intentionally run Link multiple times. You can do this by passing an `options` object with a `count` value set to an integer.

server.js

```
try {
  const response = await plaidClient.creditBankIncomeGet({
    user_token: userRecord[FIELD_USER_TOKEN],
    options: {
      count: 3,
    },
  });
  const bankIncomeData = response.data;
} catch (error) {
  // Handle this error
}
```

This call typically returns a `bank_income` array, which consists of a number of different objects that represent the report generated when the user ran Link. These individual reports contain an `items` array, which in turn contains one or more objects that represent the financial institutions associated with this report. These objects contain a `bank_income_accounts` array of objects, which gives you information about the individual accounts, as well as a `bank_income_sources` array of objects, which in turn contains the total amount of income reported, along with historical transactions about that income.

Each report also contains a `bank_income_summary` object, that summarizes the total bank income identified across all items in this report.

#### Example code

For an example of an app that incorporates Document, Payroll and Bank Income, see the React and Node-based [Income Sample app](https://github.com/plaid/income-sample). The Income Sample app is an application that uses Income and Liability data to help determine whether the user can qualify for financing on a pre-owned hoverboard. It supports the use of all three types of income data.

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
