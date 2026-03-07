---
title: "Transactions - Add Transactions to your app | Plaid Docs"
source_url: "https://plaid.com/docs/transactions/add-to-app/"
scraped_at: "2026-03-07T22:05:20+00:00"
---

# Add Transactions to your app

#### Learn how to fetch Transactions data for your users

Try out the [Pattern Demo](https://pattern.plaid.com) for a demonstration of a sample app that uses Plaid's Transactions product for the personal financial management use case.

In this guide, we'll start from scratch and walk through how to use Transactions to perform an initial fetch of a user's transaction history. If you are already familiar with using Plaid and are set up to make calls to the Plaid API, you can skip ahead to [Fetching transaction data](/docs/transactions/add-to-app/#fetching-transaction-data).

For a detailed, step-by-step view, you can also watch our full-length, comprehensive tutorial walkthrough on integrating transactions.

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

#### Create an Item with Link

Plaid Link is a drop-in module that provides a secure, elegant authentication flow
for each institution that Plaid supports. Link makes it secure and easy for users to
connect their bank accounts to Plaid. Note that these instructions cover Link on the web. For instructions on using Link within mobile apps, see the [Link documentation](/docs/link/).

Using Link, we will create a Plaid *Item*, which is a Plaid term for a login at a financial institution. An Item is not the same as a financial institution account, although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access both their checking account and their savings account, a single Item would be associated with both of those accounts. If you want to customize Link's look and feel, you can do so from the [Dashboard](https://dashboard.plaid.com/link).

Before initializing Link, you will need to create a new `link_token` on the server side of your application.
A `link_token` is a short-lived, one-time use token that is used to authenticate your app with Link.
You can create one using the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint. Then, on the client side of your application, you'll need to initialize Link with the `link_token` that you just created.

The [`/link/token/create`](/docs/api/link/#linktokencreate) sample code below will create an Item with a maximum of 90 days of transaction history. To request more, set the `transactions.days_requested` parameter in the [`/link/token/create`](/docs/api/link/#linktokencreate) request.

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
    products: ['transactions'],
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

#### Get a persistent access token

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

#### Fetching transaction data

Now that the authentication step is out of the way, we can begin using authenticated endpoints from the Plaid API and fetch transaction data using the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) endpoint.

The [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) endpoint is used to both initialize your view of transactions data, and keep you current with any changes that have occurred. When you first call it on an Item with no `cursor` parameter, transactions data available at that time is returned. If more updates are available than requested with the `count` parameter (maximum of 500), `has_more` will be set to `true`, indicating the endpoint should be called again, using the `next_cursor` from the previous response in the `cursor` field of the next request, to receive another page of data. After successfully pulling all currently available pages, you can store the cursor for later requests, allowing Plaid to send you new updates from when you last queried the endpoint.

Note that if you encounter an error during pagination, it's important to restart the pagination loop from the beginning. For more details, see the documentation for [`TRANSACTIONS_SYNC_MUTATION_DURING_PAGINATION`](/docs/errors/transactions/#transactions_sync_mutation_during_pagination). For sample code for handling the error, see the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) API reference.

Typically, the first 30 days of transaction history is available to be fetched almost immediately, but full transaction history may take a minute or more to become available. If you get an empty response when calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) shortly after linking an Item, it's likely that the first 30 days of transaction history has not yet been pulled. You will need to call the endpoint when the data is pulled. Similarly, if you only get the first 30 days of transaction history, you will need to wait until it is complete, and call the endpoint again.

To be notified whenever additional data becomes available, see [Transaction webhooks](/docs/transactions/webhooks/).

/transactions/sync

```
// Provide a cursor from your database if you've previously
// received one for the Item. Leave null if this is your
// first sync call for this Item. The first request will
// return a cursor.
let cursor = database.getLatestCursorOrNull(itemId);

// New transaction updates since "cursor"
let added: Array<Transaction> = [];
let modified: Array<Transaction> = [];
// Removed transaction ids
let removed: Array<RemovedTransaction> = [];
let hasMore = true;

// Iterate through each page of new transaction updates for item
while (hasMore) {
  const request: TransactionsSyncRequest = {
    access_token: accessToken,
    cursor: cursor,
  };
  const response = await client.transactionsSync(request);
  const data = response.data;

  // Add this page of results
  added = added.concat(data.added);
  modified = modified.concat(data.modified);
  removed = removed.concat(data.removed);

  hasMore = data.has_more;

  // Update cursor to the next cursor
  cursor = data.next_cursor;
}

// Persist cursor and updated data
database.applyUpdates(itemId, added, modified, removed, cursor);
```

#### Updating transaction data

After your initial [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) request, you may want your application to be notified when any transactions are added, removed, or modified in order to immediately fetch them from [`/transactions/sync`](/docs/api/products/transactions/#transactionssync). To learn how, see [Transaction Webhooks](/docs/transactions/webhooks/).

#### Example code in Plaid Pattern

For a real-life example of an app that incorporates transactions, see the Node-based [Plaid Pattern](https://github.com/plaid/pattern) sample app. Pattern is a sample financial management app that fetches transactions data upon receipt of transactions webhooks. Transactions code in Plaid Pattern can be found in [handleTransactionsWebhook.js](https://github.com/plaid/pattern/blob/master/server/webhookHandlers/handleTransactionsWebhook.js).

#### Fetching by date

If you want to fetch transactions data by date range, you can use the [`/transactions/get`](/docs/api/products/transactions/#transactionsget) endpoint.

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
