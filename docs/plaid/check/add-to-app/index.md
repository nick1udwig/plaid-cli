---
title: "Consumer Report (by Plaid Check) - Implementation | Plaid Docs"
source_url: "https://plaid.com/docs/check/add-to-app/"
scraped_at: "2026-03-07T22:04:47+00:00"
---

# Add Consumer Report to your app

#### Onboard your users to Plaid Check to generate cash flow insights

This guide will walk you through onboarding your users to Plaid Check, so you can generate cash flow insights for your customers by creating a Consumer Report.

This is designed as a granular, step-by-step walkthrough showing one option for a simple Plaid Check integration for customers new to Plaid. Plaid also offers a [no-code integration option](/docs/check/#no-code-integration-with-the-credit-dashboard) and a [high-level integration overview](/docs/check/#standard-integration-flow).

If you are migrating to Plaid Check Consumer Reports from a different Plaid product, see the corresponding migration guide: [Migrate from Assets](/docs/check/migrate-from-assets/), [Migrate from Income](/docs/check/migrate-from-income/), or [Migrate from Transactions](/docs/check/migrate-from-transactions/).

#### Install and initialize Plaid libraries

You can use our official server-side client libraries to connect to the Plaid Check API from your application:

Terminal

```
// Install via npm
npm install --save plaid
```

After you've installed these client libraries, initialize them by passing in your `client_id`, `secret`, and the environment you wish to connect to (Sandbox or Production). This will make sure the client libraries pass along your `client_id` and `secret` with each request, and you won't need to explicitly include them in any other calls.

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

#### Create a user

Your next step is to create a `user_id` that will represent the end user.

To create a `user_id`, make a call to the [`/user/create`](/docs/api/users/#usercreate) endpoint. This endpoint requires a unique `client_user_id` value to represent the current user. Typically, this value is the same identifier you're using in your own application to represent the currently signed-in user.

In addition, you will need to pass in an `identity` object to use Plaid Check products for this user. In this field, you should pass in user identity data that has been provided by your user for consumer report purposes. At minimum, the following fields must be provided:

- `name`
- `date_of_birth`
- `emails`
- `phone_numbers`
- `addresses`

At least one email, phone number, and address must be designated as `primary`.

If you intend to share the report with a GSE (Government-Sponsored Entity) such as Fannie or Freddie, the full SSN is also required via the `id_numbers` field. For all use cases, providing at least a partial SSN is highly recommended, since it improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests.

The [`/user/create`](/docs/api/users/#usercreate) endpoint will return a randomly generated string for the `user_id`. Make sure to save it. You send the `user_id` to Plaid Check to generate reports, and Plaid Check will use it to refer to this user when it sends you a webhook.

Depending on your application, you might wish to create this `user_id` as soon as your user creates an account, or wait until shortly before your application would want to generate a report for this user.

#### Create a Link token

Link is a client side UI widget that provides a secure, elegant flow to guide your users through the process of connecting Plaid Check with their financial institutions.

Before initializing Link, you will need to create a `link_token`. A `link_token` is a short-lived, single-use token that is used to authenticate your app with Link. You can create one using the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint from your server.

When calling [`/link/token/create`](/docs/api/link/#linktokencreate), you'll include an object telling Plaid Check about your application as well as how to customize the Link experience. In addition to the required parameters, you will need to provide the following:

- For `user_id`, provide the `user_id` you created in the previous step.
- For `products`, pass in `cra_base_report`, plus any additional Plaid Check products you wish to use (`cra_income_insights`, `cra_partner_insights`, and/or `cra_network_insights`), along with any additional Plaid Inc. products you are using. Note that Plaid Check products can't be used in the same Link flow with Income or Assets.
- Provide a `webhook` URL with an endpoint where you can receive webhooks from Plaid Check. Plaid Check will contact this endpoint when your report is ready.
- Include the appropriate options for the products you are using in the `cra_options` object.
- Include a `consumer_report_permissible_purpose` parameter specifying your purpose for generating an Consumer Report.
- On mobile, specify the `redirect_uri` and/or `android_package_name` fields as necessary, per the relevant [Link documentation](https://plaid.com/docs/link/) for your platform.

Send the `link_token` you get back in the response to your client application.

server.js

```
const request: LinkTokenCreateRequest = {
  loading_sample: true
};
try {
  const response = await plaidClient.linkTokenCreate(request);
  const linkToken = response.data.link_token;
} catch (error) {
  // handle error
}
```

##### Using Plaid Check products alongside Plaid products

If you wish to use Plaid Inc. products (Transactions, Auth, etc.) alongside Plaid Check products, you can. You can just add those products to the appropriate products array when sending a request to [`/link/token/create`](/docs/api/link/#linktokencreate). Link will take the user through a hybrid flow, where we establish connections to both Plaid Check and Plaid Inc. at the same time. One exception is that you cannot use the Plaid Inc. Income or Assets in the same session as a Plaid Check product. Instead, you should use the similar CRA Income Insights or CRA Base Consumer Report products.

#### Initialize Link

Note that these instructions cover Link on the web. For instructions on using Link within mobile apps, see the appropriate section in the [Link documentation](/docs/link/). To integrate with Embedded Link, visit these [integration steps](/docs/link/embedded-institution-search/#integration-steps).

##### Install the Link library

You will need to install the Link JavaScript library in order to use Link in your web application.

index.html

```
<head>
  <title>Extend my Credit</title>
  <script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
</head>
```

##### Configure the client-side Link handler

To run Link, you'll first need to retrieve the `link_token` from the server as described above. Then call `Plaid.create()`, passing along that token alongside other callbacks you might want to define. This will return a handler that has an `open` method. Call that method to open up the Link UI and start the connection process.

app.js

```
const linkHandler = Plaid.create({
  token: link_token_from_previous_step,
  onSuccess: (public_token, metadata) => {
    // Typically, you'd exchange the public_token for an access token.
    // You won't do that with Plaid Check products.
  },
  onExit: (err, metadata) => {
    // Optionally capture when your user exited the Link flow.
    // Storing this information can be helpful for support.
  },
  onEvent: (eventName, metadata) => {
    // Optionally capture Link flow events, streamed through
    // this callback as your users connect an Item to Plaid Check.
  },
});

linkHandler.open();
```

If you've used Link in the past with products created by Plaid, you're probably used to taking the `public_token` received from Link and exchanging it on your server for a persistent `access_token`. That is not necessary for products created by Plaid Check, and is only necessary if you are using the hybrid flow, where you are using Plaid Inc. products (Transactions, Auth, etc.) alongside Plaid Check products.

#### Generate a report

Plaid Check will start generating the Consumer Report for your user once they have completed the Link process.

Consumer Reports expire 24 hours after they have been created. After that point, you can [generate a new report](/docs/check/add-to-app/#generating-updated-reports) by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate), passing along the `user_id` you created earlier.

#### Listen for webhooks

After a user has finished using Link, it may take some time before the Consumer Report is available for you to retrieve. Listen for the `USER_CHECK_REPORT_READY` webhook to know when it is safe to request the appropriate set of product data.

webhookServer.js

```
app.post('/server/receive_webhook', async (req, res, next) => {
  const product = req.body.webhook_type;
  const code = req.body.webhook_code;
  if (product === 'CHECK_REPORT') {
    if (code === 'USER_CHECK_REPORT_READY') {
      const userId = req.body.user_id;
      // It is now okay to start fetching Consumer Report data for this user
      fetch_cra_report_for_user(userId);
    } else if (code === 'USER_CHECK_REPORT_FAILED') {
      const userId = req.body.user_id;
      // Handle this error appropriately
      report_error_with_generating_report(userId);
    }
  }
  // Handle other types of webhooks
});
```

#### Request product data

Once Plaid Check has informed you that your report is ready, you can use one of the Plaid Check product endpoints ([`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), or [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget)) to retrieve the appropriate cash flow insights. You can attempt to retrieve any type of product data, even if you didn't initialize it in the `products` array when you called [`/link/token/create`](/docs/api/link/#linktokencreate). However, retrieving new products might incur some latency as Plaid Check generates the new insight data.

Reports expire 24 hours after they have been created (either through a Link session that included the `consumer_report_permissible_purpose` parameter or through a [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) call). If you don't fetch any product data on top of this report within that time frame, you will need to generate a new Consumer Report.

#### Handling multiple institutions

If you would like your Consumer Report to include information from multiple institutions, the best way is to add the `enable_multi_item_link: true` flag to your [`/link/token/create`](/docs/api/link/#linktokencreate) call to enable [Multi-Item Link](/docs/link/multi-item-link/). Your user will then have the option to connect to as many institutions as they wish during their Link session, and the generated report will contain data from each institution that the user connects to.

##### Adding institutions after the initial Link session

If your user wants to add additional institutions *after* their initial Link session, you may do so by once again generating a Link token (using the same `user_id` as before, and including at least one Plaid Check product in the products array), and opening a new Link session on the client.

Plaid Check will start generating a new report for your user when they are done with a Link session. However, this report will *only* contain data from the institution that the user connected to during this most recent Link session.

When the user is done, re-generate your Consumer Report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) as described in the [Generating updated reports](/docs/check/add-to-app/#generating-updated-reports) section below.

#### Testing in Sandbox

Because Consumer Report information is heavily dependent on income data, Plaid recommends using the test accounts designed for producing realistic looking income data, such as `user_bank_income` (with the password `{}`). The basic Sandbox credentials (`user_good`/`pass_good`) do not return good income data.

Plaid also has additional test users to represent different levels of creditworthiness. For details, see [Credit and Income testing credentials](/docs/sandbox/test-credentials/#credit-and-income-testing-credentials).

When using special Sandbox test credentials (such as `user_bank_income` / `{}`), use a non-OAuth test institution, such as First Platypus Bank (ins\_109508) or many smaller community banks. Special test credentials may be ignored when using the Sandbox Link OAuth flow.

To use a [Sandbox custom user](https://plaid.com/docs/sandbox/user-custom/) with Plaid Check, click "add new account" rather than using one of the existing banks. Using an existing bank in the Plaid Passport flow will skip the ability to enter a custom username.

##### Testing Cash Flow Updates

After calling [`/cra/monitoring_insights/subscribe`](/docs/api/products/check/#cramonitoring_insightssubscribe), use the [`/sandbox/cra/cashflow_updates/update`](/docs/api/sandbox/#sandboxcracashflow_updatesupdate) endpoint to manually trigger an update for Cash Flow Updates (Monitoring) in the Sandbox environment.

To avoid unexpected behavior when passing `LOW_BALANCE_DETECTED` in the `webhook_codes` field of the [`/sandbox/cra/cashflow_updates/update`](/docs/api/sandbox/#sandboxcracashflow_updatesupdate) request body, use the default custom Sandbox user `user_good`.

##### Testing Income Insights

Use the [Sandbox credit testing credentials](/docs/sandbox/test-credentials/#credit-and-income-testing-credentials) with a non-OAuth test institution, such as First Platypus Bank, for more realistic and valid data.

To customize the employer name during testing, you can test with a custom Sandbox user, following the custom Sandbox user [documentation](/docs/sandbox/user-custom/) and [examples](https://github.com/plaid/sandbox-custom-users/). When creating an income transaction for this user:

- The `amount` must be negative, so that the transaction will represent income.
- `date_transacted` and `date_posted` must fall within the `cra_options.days_requested` range of the [`/link/token/create`](/docs/api/link/#linktokencreate) call.
- the `description` field must contain the desired employer name, followed by  `Direct Dep`.

The example transaction below will generate an income stream with the employer name `Weyland-Yutani`.

Sample income transaction for testing employer name

```
{
  "date_transacted": "2025-08-01", 
  "date_posted": "2025-08-03",
  "amount": -5500,
  "description": "Weyland-Yutani Direct Dep",
  "currency": "USD"
 }
```

#### Generating updated reports

If, after some time has passed, you wish to generate a new report with updated data (or a user has connected additional institutions), you can do so by performing the following steps:

1. Generate the new report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) with the same `user_id` you used earlier.
2. Wait for the `USER_CHECK_REPORT_READY` webhook.
3. Once you have received the webhook, use one of the Plaid Check product endpoints ([`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), or [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget)) to retrieve the corresponding product data.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
