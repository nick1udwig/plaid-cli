---
title: "Liabilities - Add Liabilities to your app | Plaid Docs"
source_url: "https://plaid.com/docs/liabilities/add-to-app/"
scraped_at: "2026-03-07T22:05:02+00:00"
---

# Add Liabilities to your app

#### Use Liabilities to fetch data for credit, student loan, and mortgage accounts

In this guide, we'll start from scratch and walk through how to use Liabilities to retrieve information about credit card, PayPal credit, student loan, and mortgage accounts. If you are already familiar with using Plaid and are set up to make calls to the Plaid API, make sure to initialize Link with the `liabilities` product and to note that Liabilities supports the use of [account subtype filtering](/docs/liabilities/add-to-app/#create-an-item-in-link); you can then skip ahead to [Fetching Liabilities data](/docs/liabilities/add-to-app/#fetching-liabilities-data).

#### Get Plaid API keys and complete application profile

If you don't already have one, you'll need to [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You will also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) on the Dashboard. The information in your profile will be shared with users of your application when they manage their connection on the [Plaid Portal](https://my.plaid.com). Your application profile must be completed before connecting to certain institutions in Production.

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

Using Link, we will create a Plaid *Item*, which is a Plaid term for a login at a financial institution. An Item is not the same as a financial institution account, although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access both their checking account and their savings account, a single Item would be associated with both of those accounts.

First, on the client side of your application, you'll need to set up and configure Link. If you want to customize Link's look and feel, you can do so from the [Dashboard](https://dashboard.plaid.com/link).

When initializing Link, you will need to specify the products you will be using in the product array.

Liabilities supports the use of the `account_filters` parameter to [`/link/token/create`](/docs/api/link/#linktokencreate), which you can use to limit the account types that will be shown in Link to either only credit card accounts, only student loan accounts, or only mortgage accounts. For student loans, use the type `loan` and subtype `student`. For credit cards, use the type `credit` and subtype `credit card`. For mortgages, use the type `loan` and subtype `mortgage`.

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
    products: ['liabilities'],
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

#### Fetching Liabilities data

Now that the authentication step is out of the way, we can begin using authenticated endpoints from the Plaid API. For more detailed information on the schema for account information returned, see [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget).

```
const { LiabilitiesGetRequest } = require('plaid');

// Retrieve Liabilities data for an Item
const request: LiabilitiesGetRequest = {
  access_token: accessToken,
};
try {
  const response = await plaidClient.liabilitiesGet(request);
  const liabilities = response.data.liabilities;
} catch (error) {
  // handle error
}
```

Example response data is below. The data provided is for illustrative purposes; in a Production environment, it would be very unusual for a single institution to provide both credit card and student loan data.

Liabilities sample response

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
    },
    {
      "account_id": "BxBXxLj1m4HMXBm9WZJyUg9XLd4rKEhw8Pb1J",
      "balances": {
        "available": null,
        "current": 56302.06,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "8888",
      "name": "Plaid Mortgage",
      "official_name": null,
      "subtype": "mortgage",
      "type": "loan"
    }
  ],
  "item": {
    "available_products": ["balance", "credit_details", "investments"],
    "billed_products": [
      "assets",
      "auth",
      "identity",
      "liabilities",
      "transactions"
    ],
    "consent_expiration_time": null,
    "error": null,
    "institution_id": "ins_3",
    "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
    "webhook": "https://www.genericwebhookurl.com/webhook"
  },
  "liabilities": {
    "credit": [
      {
        "account_id": "dVzbVMLjrxTnLjX4G66XUp5GLklm4oiZy88yK",
        "aprs": [
          {
            "apr_percentage": 15.24,
            "apr_type": "balance_transfer_apr",
            "balance_subject_to_apr": 1562.32,
            "interest_charge_amount": 130.22
          },
          {
            "apr_percentage": 27.95,
            "apr_type": "cash_apr",
            "balance_subject_to_apr": 56.22,
            "interest_charge_amount": 14.81
          },
          {
            "apr_percentage": 12.5,
            "apr_type": "purchase_apr",
            "balance_subject_to_apr": 157.01,
            "interest_charge_amount": 25.66
          },
          {
            "apr_percentage": 0,
            "apr_type": "special",
            "balance_subject_to_apr": 1000,
            "interest_charge_amount": 0
          }
        ],
        "is_overdue": false,
        "last_payment_amount": 168.25,
        "last_payment_date": "2025-05-22",
        "last_statement_issue_date": "2025-05-28",
        "last_statement_balance": 1708.77,
        "minimum_payment_amount": 20,
        "next_payment_due_date": "2025-05-28"
      }
    ],
    "mortgage": [
      {
        "account_id": "BxBXxLj1m4HMXBm9WZJyUg9XLd4rKEhw8Pb1J",
        "account_number": "3120194154",
        "current_late_fee": 25.0,
        "escrow_balance": 3141.54,
        "has_pmi": true,
        "has_prepayment_penalty": true,
        "interest_rate": {
          "percentage": 3.99,
          "type": "fixed"
        },
        "last_payment_amount": 3141.54,
        "last_payment_date": "2025-08-01",
        "loan_term": "30 year",
        "loan_type_description": "conventional",
        "maturity_date": "2045-07-31",
        "next_monthly_payment": 3141.54,
        "next_payment_due_date": "2025-11-15",
        "origination_date": "2015-08-01",
        "origination_principal_amount": 425000,
        "past_due_amount": 2304,
        "property_address": {
          "city": "Malakoff",
          "country": "US",
          "postal_code": "14236",
          "region": "NY",
          "street": "2992 Cameron Road"
        },
        "ytd_interest_paid": 12300.4,
        "ytd_principal_paid": 12340.5
      }
    ],
    "student": [
      {
        "account_id": "Pp1Vpkl9w8sajvK6oEEKtr7vZxBnGpf7LxxLE",
        "account_number": "4277075694",
        "disbursement_dates": ["2002-08-28"],
        "expected_payoff_date": "2032-07-28",
        "guarantor": "DEPT OF ED",
        "interest_rate_percentage": 5.25,
        "is_overdue": false,
        "last_payment_amount": 138.05,
        "last_payment_date": "2025-04-22",
        "last_statement_issue_date": "2025-04-28",
        "loan_name": "Consolidation",
        "loan_status": {
          "end_date": "2032-07-28",
          "type": "repayment"
        },
        "minimum_payment_amount": 25,
        "next_payment_due_date": "2025-05-28",
        "origination_date": "2002-08-28",
        "origination_principal_amount": 25000,
        "outstanding_interest_amount": 6227.36,
        "payment_reference_number": "4277075694",
        "pslf_status": {
          "estimated_eligibility_date": "2028-01-01",
          "payments_made": 200,
          "payments_remaining": 160
        },
        "repayment_plan": {
          "description": "Standard Repayment",
          "type": "standard"
        },
        "sequence_number": "1",
        "servicer_address": {
          "city": "San Matias",
          "country": "US",
          "postal_code": "99415",
          "region": "CA",
          "street": "123 Relaxation Road"
        },
        "ytd_interest_paid": 280.55,
        "ytd_principal_paid": 271.65
      }
    ]
  },
  "request_id": "dTnnm60WgKGLnKL"
}
```

#### Set up webhook listener

To be alerted to changes in Liabilities data, listen for the [`DEFAULT_UPDATE`](/docs/api/products/liabilities/#default_update) webhook. Upon receiving this webhook, you can call [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget) again to retrieve updated liabilities information.

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
