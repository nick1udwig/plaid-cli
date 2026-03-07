---
title: "Assets - Alloy | Plaid Docs"
source_url: "https://plaid.com/docs/assets/partnerships/alloy/"
scraped_at: "2026-03-07T22:04:30+00:00"
---

# Add Alloy to your app

#### Instantly authenticate your customers' bank accounts with Alloy's API and third-party data sources

![Partnership Alloy logo](/assets/img/docs/plaid-alloy-partnership.png)

  

#### Overview

Plaid and Alloy have partnered to offer identity data orchestration and credit decisioning to financial services companies looking to address fraud and uphold compliance. Plaid enables consumers to instantly and securely authenticate their identity associated with their bank account, including address, email, phone number, and name. Merchants can leverage Plaid Identity along with Alloy to pull third-party data and verify consumer identity at account opening, lending application review and on an ongoing basis.

This guide covers partnering with Alloy for credit decisioning. To learn about integrating with Alloy for Identity, see [Identity Partnerships: Alloy](/docs/identity/partnerships/alloy/).

#### Getting Started

You'll first want to familiarize yourself with [Plaid Link](/plaid-link/), a drop-in client-side integration for the Plaid API that handles input validation, error handling, and multi-factor authentication. You will also need to have a verified [Alloy account](https://www.alloy.com/contact) to add a bank funding source. Your customers will use Link to authenticate with their financial institution and select the bank account they wish to connect. From there, you'll receive a Plaid `access_token` and a Alloy `processor_token`, which allows you to quickly and securely retrieve the user's financial data via [Alloy's API](https://developer.alloy.com/public/docs). Utilizing Plaid + Alloy enables a seamless workflow for connecting external financial accounts to Alloy.

#### Instructions

##### Set up your accounts

You'll need accounts at both Plaid and Alloy in order to use the
Plaid + Alloy integration. You'll also need to enable your
Plaid account for the Alloy integration.

First, you will need to work with the Alloy team
to [sign up for an Alloy account](https://developer.alloy.com/public/docs)
if you do not already have one.

Next, verify that your Plaid account is enabled for the integration. If you do not have a Plaid account,
[create one](https://dashboard.plaid.com/signup).

To enable your Plaid account for the integration, go to the [Integrations](https://dashboard.plaid.com/developers/integrations)
section of the account dashboard. If the integration is off, simply click the 'Enable' button
for Alloy to enable the integration.

You'll need to complete your Plaid [Application Profile](https://dashboard.plaid.com/settings/company/app-branding) in the Dashboard, which involves filling out basic information about your app, such as your company name and website. This step helps your end-users learn more how your product uses their bank information and is also required for connecting to some banks.

Finally, you'll need to go to the [Link customization UI](https://dashboard.plaid.com/link/data-transparency-v5) and pick the use cases that you will be using Alloy to power, so that Plaid can request the appropriate authorization and consent from your end users. If you have any questions, contact Alloy.

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
the Link `public_token` for a Plaid `access_token`.

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

After creating the `access_token`(s) you need, you will create an Asset Report by calling [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate).

```
const daysRequested = 90;
const options = {
  client_report_id: '123',
  webhook: 'https://www.example.com',
  user: {
    client_user_id: '7f57eb3d2a9j6480121fx361',
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

After extracting the Asset Report token, call [`/credit/relay/create`](/docs/api/products/assets/#creditrelaycreate) to generate a Relay token to share with Alloy. You'll need to contact Alloy to obtain the `secondary_client_id` string belonging to Alloy.

```
const request: CreditRelayCreateRequest = {
  report_tokens: [createResponse.data.asset_report_token],
  secondary_client_id: secondaryClientId
};
try {
  const response = await plaidClient.creditRelayCreate(request);
  const relayToken = response.data.relay_token;
} catch (error) {
  // handle error
}
```

You can create Relay tokens in the following API environments:

- Sandbox (<https://sandbox.plaid.com>): test simulated users
- Production (<https://production.plaid.com>): production environment for when you're ready to go live and
  have valid Alloy Production credentials

#### Launching to Production

##### Test with Sandbox credentials

To test the integration in Sandbox mode, simply use the Plaid [Sandbox credentials](/docs/sandbox/test-credentials/)
along when launching Link with a `link_token` created in the Sandbox environment.

When testing in the Sandbox, you have the option to use the [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) endpoint instead of the end-to-end Link flow to create a `public_token`. When using the [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate)-based flow, the Account Select flow will be bypassed and the `accounts` array will not be populated. On Sandbox, instead of using the `accounts` array, you can call [`/accounts/get`](/docs/api/accounts/#accountsget) and test with any returned account ID associated with an account with the subtype `checking` or `savings`.

##### Get ready for production

Your account is immediately enabled for our Sandbox environment (<https://sandbox.plaid.com>).
To move to Production, please request access from the [Dashboard](https://dashboard.plaid.com/developers/keys).
You will need Alloy Production credentials prior to initiating live traffic
in the Alloy API with Plaid.

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
