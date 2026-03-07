---
title: "Assets - Ocrolus | Plaid Docs"
source_url: "https://plaid.com/docs/assets/partnerships/ocrolus/"
scraped_at: "2026-03-07T22:04:30+00:00"
---

# Add Ocrolus to your app

#### Use Ocrolus with Plaid Assets to digitize data collection

![Logos of Plaid and Ocrolus side by side.](/assets/img/docs/plaid-ocrolus-partnership.png)

  

Plaid and Ocrolus have partnered to offer lenders an easier way to access bank data to make informed loan decisions. Plaid enables businesses to instantly connect a customer's bank account, giving them the ability to authenticate and retrieve account details directly from the financial institution. Ocrolus digitizes bank and credit card statements from all US financial institutions to help lenders digitize their data collection for cash-flow analysis.

With the Plaid + Ocrolus integration, your users can verify their accounts in seconds by inputting their banking credentials in Plaid’s front-end module. Plaid will retrieve the relevant bank information and pass it to Ocrolus for further digestion and reporting in a seamless, secure fashion. Plaid’s mobile-friendly module handles input validation, error handling, and multi-factor authentication, providing a seamless onboarding experience to convert more users for your business.

#### Getting started

You'll first want to familiarize yourself with [Plaid Link](/docs/link/), a drop-in integration for the Plaid API that handles input validation, error handling, and multi-factor authentication. You will also need to create or be an existing [Ocrolus customer](https://www.ocrolus.com/get-api-keys/) in order to add a bank account.

Your customers will use Link to authenticate with their financial institution and select the bank account they wish to use for payment and verification of assets. From there, you'll receive a Plaid `access_token`, which you can use generate an Ocrolus `processor_token` and/or `audit_copy_token`, depending on your use case, which allow you to quickly and securely verify banking information via the [Ocrolus API](https://www.ocrolus.com/) without having to store that sensitive information yourself.

#### Instructions

##### Set up your Plaid and Ocrolus accounts

You'll need accounts at both Plaid and Ocrolus in order to use the Plaid + Ocrolus integration. You'll also need to enable your Plaid account for the Ocrolus integration.

First, you will need to work with the Ocrolus team to [sign up for an Ocrolus account](https://www.ocrolus.com/get-api-keys/), if you do not already have one.

Next, verify that your Plaid account is enabled for the integration. If you do not have a Plaid account, [create one](https://dashboard.plaid.com/signup/ocrolus). Your account will be automatically enabled for integration access.

To verify that your Plaid account is enabled for the integration, go to the [Integrations](https://dashboard.plaid.com/developers/integrations) section of the account dashboard. If the integration is off, simply click the 'Enable' button for Ocrolus to enable the integration.

##### Create a link\_token

In order to integrate with Plaid Link, you will first need to create a `link_token`. A `link_token`
is a short-lived, one-time use token that is used to authenticate your app with Link. To create one,
make a [`/link/token/create`](/docs/api/link/#linktokencreate) request with your `client_id`, `secret`, and a few other
required parameters from your app server. View the [documentation](/docs/api/link/#linktokencreate) for a full list of `link_token` configurations.

To see your `client_id` and `secret`, visit the [Plaid Dashboard](https://dashboard.plaid.com/developers/keys).

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

##### Integrate with Plaid Link

Once you have a `link_token`, all it takes is a few lines of client-side JavaScript to launch Link.
Then, in the `onSuccess` callback, you can call a simple server-side handler to exchange the Link `public_token`
for a Plaid `access_token` and a Ocrolus `processor_token`.

Integrate Link

```
<button id="linkButton">Open Link - Institution Select</button>
<script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
<script>
  (async function(){
    var linkHandler = Plaid.create({
      // Make a request to your server to fetch a new link_token.
      token: (await $.post('/create_link_token')).link_token,
      onLoad: function() {
          // The Link module finished loading.
      },
      onSuccess: function(public_token, metadata) {
        // The onSuccess function is called when the user has
        // successfully authenticated and selected an account to
        // use.
        //
        // When called, you will send the public_token
        // and the selected accounts, metadata.accounts,
        // to your backend app server.
        //
        // sendDataToBackendServer({
        //   public_token: public_token,
        //   accounts: metadata.accounts
        // });
        console.log('Public Token: ' + public_token);
        console.log('Customer-selected account IDs: ' + metadata.accounts.map(account => account.id).join(', '));
      },
      onExit: function(err, metadata) {
        // The user exited the Link flow.
        if (err != null) {
            // The user encountered a Plaid API error
            // prior to exiting.
        }
        // metadata contains information about the institution
        // that the user selected and the most recent
        // API request IDs.
        // Storing this information can be helpful for support.
      },
    });
  })();

  // Trigger the authentication view
  document.getElementById('linkButton').onclick = function() {
    linkHandler.open();
  };
</script>
```

See the [Link parameter reference](/docs/link/web/#create) for complete documentation on possible configurations.

`Plaid.create` accepts one argument, a configuration `Object`, and returns an `Object` with three functions, [`open`](/docs/link/web/#open), [`exit`](/docs/link/web/#exit), and [`destroy`](/docs/link/web/#destroy). Calling `open` will display the "Institution Select" view, calling `exit` will close Link, and calling `destroy` will clean up the iframe.

##### Write server-side handler

The Link module handles the entire onboarding flow securely and quickly, but does not actually retrieve account data for a user. Instead, the Link module returns a `public_token` and an `accounts` array, which is a property on the `metadata` object, via the `onSuccess` callback. Exchange this `public_token` for a Plaid `access_token` using the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) API endpoint.

Once you have the `access_token` for the Item, you can create an Ocrolus `processor_token` and/or `audit_copy_token`. You'll send these tokens to Ocrolus and they will use them to securely retrieve banking information from Plaid.

```
const {
  Configuration,
  PlaidApi,
  PlaidEnvironments,
  ProcessorTokenCreateRequest,
  AssetReportAuditCopyCreateRequest,
} = require('plaid');
// Change sandbox to production to test with live users or to go live!
const configuration = new Configuration({
  basePath: PlaidEnvironments[process.env.PLAID_ENV],
  baseOptions: {
    headers: {
      'PLAID-CLIENT-ID': process.env.PLAID_CLIENT_ID,
      'PLAID-SECRET': process.env.PLAID_SECRET,
      'Plaid-Version': '2020-09-14',
    },
  },
});

const plaidClient = new PlaidApi(configuration);

try {
  // Exchange the public_token from Plaid Link for an access token.
  const tokenResponse = await plaidClient.itemPublicTokenExchange({
    public_token: PUBLIC_TOKEN,
  });
  const accessToken = tokenResponse.data.access_token;

  // Create a processor token for a specific account id.
  const request: ProcessorTokenCreateRequest = {
    access_token: accessToken,
    account_id: accountID,
    processor: 'ocrolus',
  };
  const processorTokenResponse = await plaidClient.processorTokenCreate(
    request,
  );
  const processorToken = processorTokenResponse.data.processor_token;

  // Create an Asset Report for the specific access token.
  const request: AssetReportCreateRequest = {
    access_tokens: [accessToken],
    days_requested: 90,
    options: {},
  };
  const response = await plaidClient.assetReportCreate(request);
  const assetReportId = response.data.asset_report_id;
  const assetReportToken = response.data.asset_report_token;

  // Create an audit copy token for the Asset Report.
  const auditCopyRequest: AssetReportAuditCopyCreateRequest = {
    asset_report_token: response.data.asset_report_token,
    auditor_id: 'ocrolus',
  };
  const auditCopyResponse = await plaidClient.assetReportAuditCopyCreate(auditCopyRequest);
  const auditCopyToken = auditCopyResponse.data.audit_copy_token;
} catch (error) {
  // handle error
}
```

For a valid request, the API will return a JSON response similar to:

```
{
  "processor_token": "processor-sandbox-0asd1-a92nc",
  "request_id": "UNIQUE_REQUEST_ID"
}
```

For a valid `audit_copy_token` request, the API will return a JSON response similar to:

```
{
  "audit_copy_token": "a-sandbox-3TAU2CWVYBDVRHUCAAAI27ULU4",
  "request_id": "UNIQUE_REQUEST_ID"
}
```

For more information on creating Asset Report `audit_copy_tokens`, see the documentation for the [Assets](/docs/assets/) product.

##### Testing your Ocrolus integration

You can create Ocrolus `processor_tokens` in Sandbox (sandbox.plaid.com, allows testing with simulated users). To test the integration in Sandbox mode, use the Plaid [Sandbox credentials](/docs/sandbox/test-credentials/) when launching Link with a `link_token` created in the Sandbox environment.

To move to Production, [request access](https://dashboard.plaid.com/settings/team/products) from the Dashboard. You will want to ensure that you have valid Ocrolus Production credentials prior to connecting bank accounts in the Ocrolus API with Plaid.

#### Support and questions

Find answers to many common integration questions and concerns—such as pricing, sandbox and test mode usage, and more, in our [docs](/docs/).

If you're still stuck, open a [support ticket](https://dashboard.plaid.com/support/new) with information describing the issue that you're experiencing and we'll get back to you as soon as we can.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
