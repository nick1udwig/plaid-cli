---
title: "Transactions - Pinwheel | Plaid Docs"
source_url: "https://plaid.com/docs/transactions/partnerships/pinwheel/"
scraped_at: "2026-03-07T22:05:22+00:00"
---

# Add Pinwheel to your app

#### Activate new accounts with the industry’s top performing direct deposit and bill switching solution, The Pinwheel Switch Kit.

![Partnership Pinwheel logo](/assets/img/docs/plaid-pinwheel-partnership.png)

  

#### Overview

Pinwheel and Plaid have partnered to revolutionize how financial institutions activate new customers day 1, with in-app direct deposit and bill switching at account opening. By integrating The Pinwheel Switch Kit through Plaid Link, financial institutions can offer users a secure, frictionless experience to switch their direct deposit and recurring bills to their new account. The Pinwheel Switch Kit is the only solution in market to combine both credential-less payroll authentication and automatic bill detection with 1-click switching.

Switch Kit leverages Plaid’s trusted infrastructure to handle authentication and transaction data retrieval of external accounts. Embedding Pinwheel Switch Kit powered by Plaid enables banks and credit unions to achieve primacy on day 1.

#### Getting Started

You'll first want to familiarize yourself with [Plaid Link](/plaid-link/), a drop-in client-side integration for the Plaid API that handles input validation, error handling, and multi-factor authentication. You will also need to have a verified [Pinwheel account](https://www.pinwheelapi.com/company/contact-us) to add a bank funding source. Your customers will use Link to authenticate with their financial institution and select the bank account they wish to connect. From there, you'll receive a Plaid `access_token` and a Pinwheel `processor_token`, which allows you to quickly and securely verify a bank funding source via [Pinwheel's API](https://docs.pinwheelapi.com/public/docs/switch-kit) without having to store any sensitive banking information. Utilizing Plaid + Pinwheel enables a seamless workflow for connecting external financial accounts to Pinwheel.

#### Instructions

##### Set up your accounts

You'll need accounts at both Plaid and Pinwheel in order to use the
Plaid + Pinwheel integration. You'll also need to enable your
Plaid account for the Pinwheel integration.

First, you will need to work with the Pinwheel team
to [sign up for a Pinwheel account](https://www.pinwheelapi.com/company/contact-us),
if you do not already have one.

Next, verify that your Plaid account is enabled for the integration. If you do not have a Plaid account,
[create one](https://dashboard.plaid.com/signup).

To enable your Plaid account for the integration, go to the [Integrations](https://dashboard.plaid.com/developers/integrations)
section of the account dashboard. If the integration is off, simply click the 'Enable' button
for Pinwheel to enable the integration.

You'll need to complete your Plaid [Application Profile](https://dashboard.plaid.com/settings/company/app-branding) in the Dashboard, which involves filling out basic information about your app, such as your company name and website. This step helps your end-users learn more how your product uses their bank information and is also required for connecting to some banks.

Finally, you'll need to go to the [Link customization UI](https://dashboard.plaid.com/link/data-transparency-v5) and pick the use cases that you will be using Pinwheel to power, so that Plaid can request the appropriate authorization and consent from your end users. If you have any questions, contact Pinwheel.

##### Create a link\_token

In order to integrate with Plaid Link, you will first need to create a `link_token`. A `link_token`
is a short-lived, one-time use token that is used to authenticate your app with Link. To create one,
make a [`/link/token/create`](/docs/api/link/#linktokencreate) request with your `client_id`, `secret`, and a few other
required parameters from your app server. For a full list of `link_token` configurations, see [`/link/token/create`](/docs/api/link/#linktokencreate).

To see your `client_id` and `secret`, visit the [Plaid Dashboard](https://dashboard.plaid.com/developers/keys).

/link/token/create

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

##### Integrate with Plaid Link

Once you have a `link_token`, all it takes is a few lines of client-side JavaScript to launch
Link. Then, in the `onSuccess` callback, you can call a simple server-side handler to exchange
the Link `public_token` for a Plaid `access_token` and a Pinwheel `processor_token`.

Integrate Link

```
<button id="linkButton">Open Link - Institution Select</button>
<script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
<script>
  (async function(){
    var linkHandler = Plaid.create({
      // Make a request to your server to fetch a new link_token.
      token: (await $.post('/create_link_token')).link_token,
      onSuccess: function(public_token, metadata) {
        // The onSuccess function is called when the user has successfully
        // authenticated and selected an account to use.
        //
        // When called, you will send the public_token and the selected accounts,
        // metadata.accounts, to your backend app server.
        sendDataToBackendServer({
           public_token: public_token,
           accounts: metadata.accounts
        });
      },
      onExit: function(err, metadata) {
        // The user exited the Link flow.
        if (err != null) {
            // The user encountered a Plaid API error prior to exiting.
        }
        // metadata contains information about the institution
        // that the user selected and the most recent API request IDs.
        // Storing this information can be helpful for support.
      },
    });
  })();

  // Trigger the authentication view
  document.getElementById('linkButton').onclick = function() {
    // Link will automatically detect the institution ID
    // associated with the public token and present the
    // credential view to your user.
    linkHandler.open();
  };
</script>
```

See the [Link parameter reference](/docs/link/web/#create) for complete documentation on possible configurations.

`Plaid.create` accepts one argument, a configuration `Object`, and returns an `Object` with three functions, [`open`](/docs/link/web/#open), [`exit`](/docs/link/web/#exit), and [`destroy`](/docs/link/web/#destroy). Calling `open` will display the "Institution Select" view, calling `exit` will close Link, and calling `destroy` will clean up the iframe.

##### Write server-side handler

The Link module handles the entire onboarding flow securely and quickly, but does not actually retrieve account
data for a user. Instead, the Link module returns a `public_token` and an `accounts` array, which is a property
on the `metadata` object, via the `onSuccess` callback. Exchange this `public_token` for a Plaid `access_token`
using the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) API endpoint.

The `accounts` array will contain information about bank accounts associated with the credentials entered by the
user, and may contain multiple accounts if the user has more than one bank account at the institution. If you want the user to specify only a single account to link so you know which account to use with Pinwheel, set [Account Select](https://dashboard.plaid.com/link/account-select) to "enabled for one account" in the Plaid Dashboard. When this setting is selected, the `accounts` array will always contain exactly one account.

Once you have identified the account you will use, you will send the `access_token` and `account_id` property of the account to Plaid via the [`/processor/token/create`](/docs/api/processors/#processortokencreate) endpoint in order to create a Pinwheel `processor_token`. You'll send this token
to Pinwheel and they will use it to securely retrieve transaction data from Plaid.

You can create Pinwheel `processor_tokens` in both API environments:

- Sandbox (<https://sandbox.plaid.com>): test simulated users
- Production (<https://production.plaid.com>): production environment for when you're ready to go live and
  have valid Pinwheel Production credentials

Processor token create request

```
const {
  Configuration,
  PlaidApi,
  PlaidEnvironments,
  ProcessorTokenCreateRequest,
} = require('plaid');

// Change sandbox to production when you're ready to go live!
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
    public_token: publicToken,
  });
  const accessToken = tokenResponse.data.access_token;

  // Create a processor token for a specific account id.
  const request: ProcessorTokenCreateRequest = {
    access_token: accessToken,
    account_id: accountID,
    processor: 'pinwheel',
  };
  const processorTokenResponse = await plaidClient.processorTokenCreate(
    request,
  );
  const processorToken = processorTokenResponse.data.processor_token;
} catch (error) {
  // handle error
}
```

For a valid request, the API will return a JSON response similar to:

Processor token create response

```
{
  "processor_token": "processor-sandbox-0asd1-a92nc",
  "request_id": "m8MDnv9okwxFNBV"
}
```

For possible error codes, see the full listing of Plaid [error codes](/docs/errors/).

#### Launching to Production

##### Test with Sandbox credentials

To test the integration in Sandbox mode, simply use the Plaid [Sandbox credentials](/docs/sandbox/test-credentials/)
along when launching Link with a `link_token` created in the Sandbox environment.

When testing in the Sandbox, you have the option to use the [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) endpoint instead of the end-to-end Link flow to create a `public_token`. When using the [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate)-based flow, the Account Select flow will be bypassed and the `accounts` array will not be populated. On Sandbox, instead of using the `accounts` array, you can call [`/accounts/get`](/docs/api/accounts/#accountsget) and test with any returned account ID.

##### Get ready for production

Your account is immediately enabled for our Sandbox environment (<https://sandbox.plaid.com>).
To move to Production, please request access from the [Dashboard](https://dashboard.plaid.com/developers/keys).
You will need Pinwheel Production credentials prior to initiating live traffic
in the Pinwheel API with Plaid.

##### Support and questions

Find answers to many common integration questions and concerns—such as pricing, sandbox and test mode usage,
and more, in our [docs](/docs/).

If you're still stuck, open a [support ticket](https://dashboard.plaid.com/support/new) with information
describing the issue that you're experiencing and we'll get back to you as soon as we can.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
