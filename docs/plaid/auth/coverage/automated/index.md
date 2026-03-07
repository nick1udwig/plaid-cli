---
title: "Auth - Automated Micro-deposits | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/automated/"
scraped_at: "2026-03-07T22:04:31+00:00"
---

# Automated Micro-deposits

#### Learn how to authenticate your users in a secure and frictionless micro-deposit flow

#### The Automated Micro-deposit flow

The Automated Micro-deposits authentication flow is supported for an additional 1,900 financial institutions in the US where Instant Auth is not available, accounting for approximately <1% of depository accounts.
Plaid will make a single micro-deposit and then automatically verify it within one to two business days.

The user launches Link...

![The user launches Link...](/assets/img/docs/auth/amd_tour/link-1-rm.png)

![...and selects the institution to link.](/assets/img/docs/auth/amd_tour/select-institution.png)

![They find their institution...](/assets/img/docs/auth/amd_tour/fi-selected.png)

![...and log in.](/assets/img/docs/auth/amd_tour/fi-login.png)

![Next, they select the account to link. (Single-account select is typically recommended for Auth.)](/assets/img/docs/auth/amd_tour/account-select.png)

![They verify their bank's routing number...](/assets/img/docs/auth/amd_tour/verify-routing.png)

![...their account number...](/assets/img/docs/auth/amd_tour/verify-account.png)

![...and their name.](/assets/img/docs/auth/amd_tour/enter-name.png)

![Plaid will send a micro-deposit. Once it lands, Plaid will automatically detect it and verify the account.](/assets/img/docs/auth/amd_tour/await-verification.png)

You can try out the Automated Micro-deposits flow in an [Interactive Demo](https://plaid.coastdemo.com/share/67d0ce0df465686c02cc4fd2?zoom=100&step=6). See more details in our [testing guide](/docs/auth/coverage/testing/#testing-automated-micro-deposits).

A user connects their financial institution using the following connection flow:

1. Starting on a page in your app, the user clicks an action that opens Plaid Link, with the correct
   Auth [configuration](/docs/auth/coverage/automated/#configure--create-a-link_token).
2. Inside of Plaid Link, the user selects their institution, authenticates with their credentials,
   provides their account and routing number, and enters in their legal name.
3. Upon [successful authentication](/docs/auth/coverage/automated/#exchange-the-public-token), Link closes with a
   `public_token` and a `metadata` account status of `pending_automatic_verification`.
4. Behind the scenes, Plaid sends a single micro-deposit to the user's account and will automatically
   verify the deposited amounts within one to two business days.
5. When verification succeeds or fails, Plaid sends an [Auth webhook](/docs/auth/coverage/automated/#handle-auth-webhooks),
   which you can use to notify the user that their account is ready to move money. Once this step is done, your user's Auth data is verified and [ready to fetch](/docs/auth/coverage/automated/#fetch-auth-data).

#### Configure & Create a link\_token

Create a `link_token` with the following parameters:

- `products` array containing `auth` or `transfer` -- unlike with Same-Day Micro-deposits, you can also include other products besides `auth` or `transfer` when creating a Link token for use with Automated Micro-deposits, but `auth` or `transfer` must be present.
- `country_codes` set to `['US']` – Micro-deposit verification is currently only available in the United States.
- A `webhook` URL to receive a POST HTTPS request sent from Plaid's servers to your application
  server, after Automated Micro-deposits succeeds or fails verification of a user's micro-deposits.
- `auth` object should specify `"automated_microdeposits_enabled": true`

Automated Micro-deposits Configuration

```
const request: LinkTokenCreateRequest = {
  user: { client_user_id: new Date().getTime().toString() },
  client_name: 'Plaid App',
  products: [Products.Auth],
  country_codes: [CountryCode.Us],
  language: 'en',
  auth: {
    automated_microdeposits_enabled: true,
  },
};
try {
  const response = await plaidClient.linkTokenCreate(request);
  const linkToken = response.data.link_token;
} catch (error) {
  // handle error
}
```

#### Initialize Link with a link\_token

After creating a `link_token` for the `auth` product, use it to initialize Plaid Link.

When the user inputs their username and password, and account and routing numbers for the financial institution,
the `onSuccess()` callback function will return a `public_token`, with `verification_status` equal to `'pending_automatic_verification'`.

App.js

```
const linkHandler = Plaid.create({
  // Fetch a link_token configured for 'auth' from your app server
  token: (await $.post('/create_link_token')).link_token,
  onSuccess: (public_token, metadata) => {
    // Send the public_token and accounts to your app server
    $.post('/exchange_public_token', {
      publicToken: public_token,
      accounts: metadata.accounts,
    });

    metadata = {
      ...,
      link_session_id: String,
      institution: { name: 'Bank of the West', institution_id: 'ins_100017' },
      accounts: [{
        id: 'vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D',
        mask: '1234',
        name: null,
        type: 'depository',
        subtype: 'checking',
        verification_status: 'pending_automatic_verification'
      }]
    }
  },
  // ...
});

// Open Link on user-action
linkHandler.open();
```

##### Display a "pending" status in your app

Because Automated verification usually takes between one to two days to complete, we recommend displaying a UI
in your app that communicates to a user that verification will occur automatically and is currently pending.

You can use the `verification_status` key returned in the `onSuccess` `metadata.accounts` object once
Plaid Link closes successfully.

Metadata verification\_status

```
verification_status: 'pending_automatic_verification';
```

You can also [fetch the `verification_status`](/docs/auth/coverage/automated/#check-the-account-verification-status-optional) for an
Item's account via the Plaid API to obtain the latest account status.

#### Exchange the public token

In your own backend server, call the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange)
endpoint with the Link `public_token` received in the `onSuccess` callback to obtain an `access_token`.
Persist the returned `access_token` and `item_id` in your database in relation to the user.

Note that micro-deposits will only be delivered to the ACH network in the Production environment. To test your integration outside of Production, see [Testing automated micro-deposits in Sandbox](/docs/auth/coverage/testing/#testing-automated-micro-deposits).

Exchange token request

```
// publicToken and accountID are sent from your app to your backend-server
const accountID = 'vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D';
const publicToken = 'public-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d';

// Obtain an access_token from the Link public_token
try {
  const response = await client.itemPublicTokenExchange({
    public_token: publicToken,
  });
  const accessToken = response.data.access_token;
} catch (err) {
  // handle error
}
```

Exchange token response

```
{
  "access_token": "access-sandbox-5cd6e1b1-1b5b-459d-9284-366e2da89755",
  "item_id": "M5eVJqLnv3tbzdngLDp9FL5OlDNxlNhlE55op",
  "request_id": "m8MDnv9okwxFNBV"
}
```

#### Handle Auth webhooks

Before you can call [`/auth/get`](/docs/api/products/auth/#authget) to fetch Auth data for a user's `access_token`, a micro-deposit first
need to post successfully to the user's bank account. Because Plaid uses Same Day ACH to send a
single micro-deposit amount, this process usually takes between one to two days.

Once the deposit has arrived in the user's account, Plaid will automatically verify the deposit
transaction and send an [`AUTOMATICALLY_VERIFIED`](/docs/api/products/auth/#automatically_verified)
webhook to confirm the account and routing numbers have been successfully verified.

Attempting to call [`/auth/get`](/docs/api/products/auth/#authget) on an unverified `access_token` will result
in a [`PRODUCT_NOT_READY`](/docs/errors/item/#product_not_ready) error.

Auth AUTOMATICALLY\_VERIFIED webhook

```
> POST https://your_app_url.com/webhook

{
  "webhook_type": "AUTH",
  "webhook_code": "AUTOMATICALLY_VERIFIED",
  "item_id": "zeWoWyv84xfkGg1w4ox5iQy5k6j75xu8QXMEm",
  "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D"
}
```

Occasionally automatic verification may fail, likely due to erroneous user input, such as an incorrect
account and routing number pair. If the Item is unable to be verified within seven days, Plaid will send
a [`VERIFICATION_EXPIRED`](/docs/api/products/auth/#verification_expired)
webhook. When verification fails, the Item is permanently locked; we recommend prompting your user to
retry connecting their institution via Link.

Auth VERIFICATION\_EXPIRED webhook

```
> POST https://your_app_url.com/webhook

{
  "webhook_type": "AUTH",
  "webhook_code": "VERIFICATION_EXPIRED",
  "item_id": "zeWoWyv84xfkGg1w4ox5iQy5k6j75xu8QXMEm",
  "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D"
}
```

If Plaid encounters an `ITEM_LOGIN_REQUIRED` error during attempted validation, this may mean that Plaid lost access to the user's account after sending this micro-deposit but before being able to verify it. If this occurs, send the user through the [update mode](/docs/link/update-mode/) flow to re-verify their account.

The example code below shows how to handle `AUTOMATICALLY_VERIFIED` and `VERIFICATION_EXPIRED` webhooks
and call [`/auth/get`](/docs/api/products/auth/#authget) to retrieve account and routing data.

If you are using the Sandbox environment, you can use
the [`/sandbox/item/set_verification_status`](/docs/api/sandbox/#sandboxitemset_verification_status)
endpoint to test your integration.

Webhook example code

```
// This example uses Express to receive webhooks
const app = require('express')();
const bodyParser = require('body-parser');
app.use(bodyParser);

app.post('/webhook', async (request, response) => {
  const event = request.body;

  // Handle the event
  switch (event.webhook_code) {
    case 'AUTOMATICALLY_VERIFIED':
      const accessToken = lookupAccessToken(event.item_id);
      const request: AuthGetRequest = { access_token: accessToken };
      const authResponse = await client.authGet(request);
      const numbers = authResponse.numbers;
      break;
    case 'VERIFICATION_EXPIRED':
      // handle verification failure; prompt user to re-authenticate
      console.error('Verification failed for', event.item_id);
      break;
    default:
      // Unexpected event type
      return response.status(400).end();
  }

  // Return a response to acknowledge receipt of the event
  response.json({ received: true });
});

app.listen(8000, () => console.log('Running on port 8000'));
```

#### Check the account verification status *(optional)*

In some cases you may want to implement logic to display the `verification_status` of an Item
that is pending automated verification in your app. The [`/accounts/get`](/docs/api/accounts/#accountsget)
API endpoint allows you to query this information.

For non-programmatic access to this information, you can use the [Account Verification Dashboard](https://dashboard.plaid.com/account-verification/).

Accounts get request

```
// Fetch the accountID and accessToken from your database
const accountID = 'vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D';
const accessToken = 'access-sandbox-5cd6e1b1-1b5b-459d-9284-366e2da89755';
const request: AccountsGetRequest = {
  access_token: accessToken,
};
try {
  const response = await client.accountsGet(request);
  const account = response.data.accounts.find((a) => a.account_id === accountID);
  const verificationStatus = account.verification_status;
} catch (err) {
  // handle error
}
```

Account get response

```
{
  "accounts": [
    {
      "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
      "balances": { Object },
      "mask": "0000",
      "name": "Plaid Checking",
      "official_name": "Plaid Gold Checking",
      "type": "depository",
      "subtype": "checking",
      "verification_status":
        "pending_automatic_verification" |
        "automatically_verified" |
        "verification_expired"
    },
    ...
  ],
  "item": { Object },
  "request_id": String
}
```

#### Fetch Auth data

Finally, we can retrieve Auth data once automated verification has succeeded:

Auth request

```
const accessToken = 'access-sandbox-5cd6e1b1-1b5b-459d-9284-366e2da89755';

// Instantly fetch Auth numbers
const request: AuthGetRequest = {
  access_token: accessToken,
};
try {
  const response = await client.authGet(request);
  const numbers = response.data.numbers;
} catch (err) {
  // handle error
}
```

Auth response

```
{
  "numbers": {
    "ach": [
      {
        "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
        "account": "1111222233330000",
        "routing": "011401533",
        "wire_routing": "021000021"
      }
    ],
    "eft": [],
    "international": [],
    "bacs": []
  },
  "accounts": [
    {
      "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
      "balances": { Object },
      "mask": "0000",
      "name": "Plaid Checking",
      "official_name": "Plaid Gold Standard 0% Interest Checking",
      "verification_status": "automatically_verified",
      "subtype": "checking" | "savings",
      "type": "depository"
    }
  ],
  "item": { Object },
  "request_id": "m8MDnv9okwxFNBV"
}
```

Check out the [`/auth/get`](/docs/api/products/auth/#authget) API reference documentation to see the full
Auth request and response schema.

#### Handling Link events

For a user who goes through the Automated Micro-deposit flow, the `TRANSITION_VIEW (view_name = NUMBERS)`
event will occur after `SUBMIT_CREDENTIALS`, and in the `onSuccess` callback the
`verification_status` will be `pending_automatic_verification`.

Sample Link events for Automated Micro-deposits flow

```
OPEN (view_name = CONSENT)
TRANSITION_VIEW (view_name = SELECT_INSTITUTION)
SEARCH_INSTITUTION
SELECT_INSTITUTION
TRANSITION_VIEW (view_name = CREDENTIAL)
SUBMIT_CREDENTIALS
TRANSITION_VIEW (view_name = LOADING)
TRANSITION_VIEW (view_name = MFA, mfa_type = code)
SUBMIT_MFA (mfa_type = code)
TRANSITION_VIEW (view_name = LOADING)
TRANSITION_VIEW (view_name = SELECT_ACCOUNT)
TRANSITION_VIEW (view_name = NUMBERS)
TRANSITION_VIEW (view_name = LOADING)
TRANSITION_VIEW (view_name = CONNECTED)
HANDOFF
onSuccess (verification_status: pending_automatic_verification)
```

[#### Same Day Micro-deposits

Integrate the manual micro-deposit flow

View guide](/docs/auth/coverage/same-day/)

#### Same Day Micro-deposits

Integrate the manual micro-deposit flow

[View guide](/docs/auth/coverage/same-day/)[#### Testing in Sandbox

Learn how to test each Auth flow in the Sandbox

View guide](/docs/auth/coverage/testing/)

#### Testing in Sandbox

Learn how to test each Auth flow in the Sandbox

[View guide](/docs/auth/coverage/testing/)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
