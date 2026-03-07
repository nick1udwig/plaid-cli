---
title: "Auth - Stripe | Plaid Docs"
source_url: "https://plaid.com/docs/auth/partnerships/stripe/"
scraped_at: "2026-03-07T22:04:43+00:00"
---

# Add Stripe to your app

#### Use Stripe with Plaid Auth to send and receive payments

![Plaid + Stripe logo image](/assets/img/products/plaid-link-connect-account-flow.png)

  

Plaid and Stripe have partnered to offer frictionless money transfers without the need to ever handle an account or routing number. Use [Plaid Link](/docs/link/) to instantly authenticate your customer's account and generate a Stripe [bank account token](https://stripe.com/docs/api#create_bank_account_token) so that you can accept ACH payments via Stripe.

This guide is designed for those who already have developer accounts with both Stripe and Plaid. If that's not you, [Sign up for a Plaid Dashboard account](https://dashboard.plaid.com/signup), then head over to the [Stripe docs](https://dashboard.stripe.com/register) to get a Stripe account.

This guide describes integrating Plaid with the [Stripe Payment Intents API](https://docs.stripe.com/payments/payment-intents), which is the modern way to send and receive ACH payments via Stripe.

Existing customers who have integrations using the older Stripe Charges and Sources API should contact their Plaid account manager with any questions. For more details on the deprecation of the Charges and Sources API, see the [changelog](/docs/changelog/#february-20-2026).

#### Getting Started

You'll first want to familiarize yourself with [Plaid Link](/docs/link/), a drop-in client-side integration for the Plaid API that handles input validation, error handling, and multi-factor authentication.

Your customers will use Link to authenticate with their financial institution and select the depository account they wish to use for ACH transactions. From there, you'll receive a Plaid `access_token`, allowing you to leverage real-time balance checks and transaction data, and a Stripe `bank_account_token`, which allows you to move money via Stripe's ACH API without ever handling an account or routing number.

#### Instructions

##### Set up your Plaid and Stripe accounts

You'll need accounts at both Plaid and Stripe in order to use the Plaid Link + Stripe integration. You'll also need to connect your Plaid and Stripe accounts so that Plaid can facilitate the creation of bank account tokens on your behalf.

First, [sign up for a Stripe account](https://dashboard.stripe.com/register) if you do not already have one, and then verify that it is enabled for ACH access. To verify that your Stripe account is ACH enabled, head to the Settings → Payment Methods section of the Stripe Dashboard and make sure that ACH Direct Debit is marked as enabled. If it is not, click the button to enable it.

If you do not have a Plaid account, [create one](https://dashboard.plaid.com/signup/stripe).

To verify that your Plaid account is enabled for the Stripe integration, go to the [Integrations](https://dashboard.plaid.com/developers/integrations) section of the account dashboard. If you see:

![Unlinked Stripe account](/assets/img/products/dashboard-stripe-account-unlinked.png)

your Plaid account is enabled for the integration but you have not connected your Stripe account.

Click the 'Connect With Stripe' button to connect your Plaid and Stripe accounts. This step is required so that Plaid can facilitate the creation of Stripe bank account tokens on your behalf.

Once your Stripe account is connected, you'll see:

![linked Stripe account](/assets/img/products/dashboard-stripe-account-linked.png)

Your Plaid account is now set up for the integration!

###### Contact Stripe to enable the Payment Intents API for Plaid

To use Plaid with the Stripe Payment Intents API, you will need to be manually enabled by Stripe.

The following is a template you can use to request access:

Email template to request Plaid / Stripe enablement

```
To: plaid-referral@stripe.com
Cc: partnerships@plaid.com
Subject: {{Your business name}} is requesting Plaid <> Stripe Payment Intents support

Stripe Account ID: {{your Stripe account id, beginning with acct_}}

Use Case:

{{A description of what you are trying to accomplish with your Stripe + Plaid integration.}}

Thank you for your assistance,
{{Your name}}
```

##### Complete your Plaid application profile and company profile

After connecting your Stripe and Plaid accounts, you'll need to complete your Plaid [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile) in the Dashboard, which involves filling out basic information about your app, such as your company name and website. This step helps your end-users learn more how your product uses their bank information and is also required for connecting to some banks.

##### Create a link\_token

In order to integrate with Plaid Link, you will first need to create a `link_token`. A `link_token`
is a short-lived, one-time use token that is used to authenticate your app with Link. To create one,
make a [`/link/token/create`](/docs/api/link/#linktokencreate) request with your `client_id`, `secret`, and a few other
required parameters from your app server. For a full list of parameters, see the [`/link/token/create`](/docs/api/link/#linktokencreate) reference.

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
    products: ['auth'],
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
for a Plaid `access_token` and a Stripe bank account token.

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
        // and the data about which account to use
        // to your backend app server. In this example,
        // we assume "Account Select" is set to
        // "enabled for one account".
        //
        // sendDataToBackendServer({
        //   public_token: public_token,
        //   account_id: metadata.accounts[0].id
        // });
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

The `accounts` array will contain information about bank accounts associated with the credentials entered by the user, and may contain multiple accounts if the user has more than one bank account at the institution. In order to avoid any confusion about which account your user wishes to use with Stripe, it is recommended to set [Account Select](https://dashboard.plaid.com/link/account-select) to "enabled for one account" in the Plaid Dashboard. When this setting is selected, the `accounts` array will always contain exactly one element.

Once you have identified the account you will use, you will call [`/processor/stripe/bank_account_token/create`](/docs/api/processors/#processorstripebank_account_tokencreate), sending Plaid the `account_id` property of the account along with the `access_token` you obtained from [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange). Plaid will create and return a Stripe [bank account token](https://stripe.com/docs/api#create_bank_account_token) for this account, which can then be used to move money via Stripe's ACH API. The bank account token will be linked to the Stripe account you linked in your [Plaid Dashboard](https://dashboard.plaid.com/).

Note that the Stripe bank account token is a one-time use token. To store bank account information for later use, you can use a Stripe [customer object](https://stripe.com/docs/api/customers/object) and create an associated [bank account](https://stripe.com/docs/api/customer_bank_accounts/create) from the token, or you can use a Stripe [Custom account](https://stripe.com/docs/api/accounts) and create an associated [external bank account](https://stripe.com/docs/api/external_account_bank_accounts/create) from the token. This bank account information should work indefinitely, unless the user's bank account information changes or they revoke Plaid's permissions to access their account (when using certain financial institutions). The Plaid `access_token`, on the other hand, does not expire and should be persisted securely.

Stripe bank account information cannot be modified once the bank account token has been created. If you ever need to change the bank account details used by Stripe for a specific customer, you should repeat this process from the beginning, including creating a `link_token`, sending the customer through the Link flow, generating a new bank account token, and creating a new customer object or Custom account.

```
// Change sandbox to production when you're ready to go live!
const {
  Configuration,
  PlaidApi,
  PlaidEnvironments,
  ProcessorStripeBankAccountTokenCreateRequest,
} = require('plaid');
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

  // Generate a bank account token
  const request: ProcessorStripeBankAccountTokenCreateRequest = {
    access_token: accessToken,
    account_id: accountID,
  };
  const stripeTokenResponse = await plaidClient.processorStripeBankAccountTokenCreate(
    request,
  );
  const bankAccountToken = stripeTokenResponse.data.stripe_bank_account_token;
} catch (error) {
  // handle error
}
```

For a valid request, the API will return a JSON response similar to:

```
{
  "stripe_bank_account_token": "btok_5oEetfLzPklE1fwJZ7SG",
  "request_id": "[Unique request ID]"
}
```

For possible error codes, see the full listing of Plaid [error codes](/docs/errors/).

Note: The `account_id` parameter is required if you wish to receive a Stripe bank account token.

##### Test with Sandbox credentials

You can create Stripe bank account tokens in all three API environments:

- Sandbox (<https://sandbox.plaid.com>): test simulated users with Stripe's "test mode" API
- Production (<https://production.plaid.com>): production environment for when you're ready to go live

Plaid's Sandbox API environment is compatible with Stripe's "test mode" API. To test the integration in Sandbox mode, simply use the Plaid [Sandbox credentials](/docs/sandbox/test-credentials/) when launching Link with a `link_token` created in the Sandbox environment. The Stripe bank account token created in the Sandbox environment will always match the Stripe bank test account with account number 000123456789 and routing number 110000000, and with the Stripe account that is linked in the Plaid Dashboard.

Use Stripe's [ACH API](https://stripe.com/docs/guides/ach) in test mode to create test transfers using the bank account tokens you retrieve from Plaid's Sandbox API environment.

Plaid's Sandbox is not compatible with the Stripe Sandbox environment. To test your Plaid - Stripe integration, you must use Stripe test mode, rather than Stripe Sandboxes.

When testing in the Sandbox, you have the option to use the [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) endpoint instead of the end-to-end Link flow to create a `public_token`. When using the [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate)-based flow, the Account Select flow will be bypassed and the `accounts` array will not be populated. On Sandbox, instead of using the `accounts` array, you can call [`/accounts/get`](/docs/api/accounts/#accountsget) and test with any returned account ID associated with an account with the subtype `checking` or `savings`.

##### Get ready for production

Your account is immediately enabled for our Sandbox environment (<https://sandbox.plaid.com>), which allows you to test with Sandbox API credentials. To move to Production, request access from the [Dashboard](https://dashboard.plaid.com/settings/team/products).

##### Next steps

Once you're successfully retrieving Stripe `bank_account_token`s, you're ready to make ACH transactions with Stripe. Head to [Stripe's ACH guide](https://stripe.com/docs/ach) to get started, and if you have any issues, please reach out to Stripe Support.

If you are using the [Charges and Sources API](https://stripe.com/docs/ach-deprecated), the Stripe documentation will walk you through the process.

If you are using the Payment Intents API, you will first need to create a Stripe [customer](https://stripe.com/docs/api/customers/create#create_customer-payment_method) using the bank account token, then follow the [migration guide](https://stripe.com/docs/payments/ach-debit/migrations#create-payment-intent) to create a Payment Intent based on the bank account, using the [manual migration flow](https://docs.stripe.com/payments/ach-direct-debit/migrating-from-another-processor#manual-bank-account-migration). Note that you will need to contact Stripe to request access to this flow.

#### Example code in Plaid Pattern

For a real-life example of an app that integrates a processor partner, see the Node-based [Plaid Pattern Account Funding](https://github.com/plaid/pattern-account-funding) sample app. Pattern Account Funding is a sample account funding app that creates a processor token to send to your payment partner. (Note that while the processor token creation call in Plaid Pattern's [items.js](https://github.com/plaid/pattern-account-funding/blob/master/server/routes/items.js#L126-L135) demonstrates usage of the [`/processor/token/create`](/docs/api/processors/#processortokencreate) endpoint for creating a processor token, Stripe integrations should use the [`/processor/stripe/bank_account_token/create`](/docs/api/processors/#processorstripebank_account_tokencreate) endpoint instead and create a bank account token.)

#### Troubleshooting

##### Stripe bank account token not returned

When a `stripe_bank_account_token` is not returned, a typical cause is that your Stripe and Plaid accounts have not yet been linked. First, be sure that your Stripe account has been [configured to accept ACH transfers](https://stripe.com/docs/guides/ach). Then, [link your Stripe account to your Plaid account](/docs/auth/partnerships/stripe/#set-up-your-plaid-and-stripe-accounts) and re-try your request.

##### Stripe::InvalidRequestError: No such token

The `Stripe::InvalidRequestError: No such token` error is returned by Stripe when the `bank_account_token` could not be used with the Stripe API. Make sure you are not mixing Plaid's Sandbox environment with Stripe's production environment and that, if you have more than one Stripe account or more than one Plaid `client_id`, you are using the Stripe account that is linked to the Plaid `client_id` you are using.

#### Support and questions

Find answers to many common integration questions and concerns, such as pricing, sandbox and test mode usage in our [docs](/docs/).

If you're still stuck, open a [support ticket](https://dashboard.plaid.com/support/new) with information describing the issue that you're experiencing and we'll get back to you as soon as we can.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
