---
title: "Auth - Database Insights and Match (legacy) | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/database/"
scraped_at: "2026-03-07T22:04:31+00:00"
---

# Database Insights and Match (legacy)

#### Evaluate a manually entered account using Plaid network data

![Payment method selection page showing bank options with search. Car details, insurance info, and payment summary are displayed.](/assets/img/docs/auth/embedded-search-database-insights.gif)

The Database Insights end-user flow with Embedded Institution Search enabled.

#### Overview

Database Insights and Database Match have been replaced with Database Auth in the US. The documentation provided here is for the use of customers maintaining support of existing legacy Database Insights and Database Match integrations, or building integrations for Canada.

Database Insights (legacy) can increase conversion by providing instant account verification without requiring users to link a bank account via credentials. End users choosing the manual Database Insights path will not be required to log in to their financial institution and instead can enter their account and routing number manually.

Database Insights verifies account and routing numbers by checking the information provided against Plaid's known account numbers, leveraging Plaid's database of over 200 million verified accounts. If no match is found, Plaid will check the account number format against known usages by the institution associated with the given routing number. Database Insights will provide a verification status of 'pass', 'pass with caution', or 'fail' and a set of attributes that contributed to that status, such as whether a match was found or whether Plaid fell back to checking account number formats.

Database Insights does not verify that the user has access to the bank account and Database Insights does also not fully guarantee that the account exists, especially in a 'pass with caution' scenario. For these reasons, Database Insights should only be enabled where there is a low risk of fraud or ACH returns. Examples of use cases where Database Insights may be appropriate include bill payment, rent collection, business to business payments, or subscription payments.

Approximately 30% of Items verified by Database Insights can also be verified by [`/identity/match`](/docs/api/products/identity/#identitymatch) or [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate). For more details, see [Identity Match](/docs/identity/#identity-match) and [Signal Transaction Scores](/docs/signal/). All Items verified with Database Insights are also compatible with account ownership identity verification via [Identity Document Upload](/docs/identity/identity-document-upload/).

Database Insights and [Embedded Institution Search](/docs/link/embedded-institution-search/) are both designed to increase adoption of ACH payment methods and are frequently used together. Database Insights is also fully compatible with the standard, non-Embedded Institution Search Link flow.

##### Database Insights flow

1. Starting on a page in your app, the user clicks an action that opens Plaid Link.
2. Inside of Plaid Link, the user selects an option to enter the manual verification flow
   and provides their account and routing number.
3. Once the user has submitted their information, Link closes and returns a `public_token` within the `onSuccess` callback.
4. Call `/item/token/exchange` to exchange the `public_token` for an `access_token`, then call [`/auth/get`](/docs/api/products/auth/#authget) to obtain the account numbers and the verification results.
5. Based on the values of the `verification_status` and `verification_insights` fields returned by [`/auth/get`](/docs/api/products/auth/#authget), make a decision whether to proceed with the account information for ACH or to reject the account information as unverified.

When using other flows, customers using a [processor partner](/docs/auth/partnerships/) do not typically need to call [`/auth/get`](/docs/api/products/auth/#authget), and can directly call [`/processor/token/create`](/docs/api/processors/#processortokencreate) instead. However, if you are using Database Insights with a processor partner, you must call [`/auth/get`](/docs/api/products/auth/#authget) and check the value of the `verification_status` and/or `verification_insights` fields before passing a processor token to the partner.

#### Implementation steps

##### Create a link\_token

Create a `link_token` with the following parameters:

- The `products` array should include only `auth` or `transfer` as a product when using Database Insights. While in most cases additional products can be added to existing Plaid Items, Items created with Database Insights are an exception and cannot be used with any Plaid products other than Auth, Transfer, Signal, or Identity Match.

Approximately 30% of Items verified by Database Insights can also be verified by [`/identity/match`](/docs/api/products/identity/#identitymatch) or Signal. For more details, see [Identity Match](/docs/identity/#identity-match) and [Signal](/docs/signal/signal-rules/#data-availability-limitations). If using Identity Match or Signal in this way, they should be added to the Item via the `required_if_supported_products`, `optional_products`, or `additional_consented_products` fields rather than the `products` array.

- `country_codes` should be set to `['US']` or `['CA']`– Database Insights is currently only available in the United States or Canada.
- The `auth` object should specify `"database_insights_enabled": true`.
- (Optional) Within the `auth` object, specify `"auth_type_select_enabled": true` in order to enable [Auth Type Select](/docs/auth/coverage/flow-options/#adding-manual-verification-entry-points-with-auth-type-select), which will surface the manual entry point for Database Insights on the default Link screen. If Auth Type Select is not enabled, credential-based flows will be the primary UI flow and the Database Insights entry point will only appear as a fallback, except when using [Embedded Institution Search](/docs/link/embedded-institution-search/).

Database Insights cannot be used in the same session as Database Match, Same-Day Micro-deposits, or Instant Micro-deposits. If any of those flows are explicitly enabled in the `auth` object alongside Database Insights, an error will occur when calling [`/link/token/create`](/docs/api/link/#linktokencreate). If they are implicitly enabled due to account default settings, they will be overridden by enabling Database Insights.

Database Insights Configuration

```
const request: LinkTokenCreateRequest = {
  user: { client_user_id: new Date().getTime().toString() },
  client_name: 'Plaid App',
  products: [Products.Auth],
  country_codes: [CountryCode.Us],
  webhook: 'https://webhook.sample.com',
  language: 'en',
  auth: {
    database_insights_enabled: true
  },
};
try {
  const response = await plaidClient.linkTokenCreate(request);
  const linkToken = response.data.link_token;
} catch (error) {
  // handle error
}
```

##### Initialize Link with a link\_token

After creating a `link_token` for the `auth` product, use it to initialize Plaid Link.

When the user successfully inputs their account and routing numbers, the `onSuccess()` callback
function will return a `public_token`.

App.js

```
const linkHandler = Plaid.create({
  // Fetch a link_token configured for 'auth' from your app server
  token: (await $.post('/create_link_token')).link_token,
  onSuccess: (public_token, metadata) => {
    // Send the public_token and connected accounts to your app server
    $.post('/exchange_public_token', {
      publicToken: public_token,
      accounts: metadata.accounts,
    });

    metadata = {
      ...,
      link_session_id: String,
      institution: {
        name: null, // name is always null for Database Insights
        institution_id: null // institution_id is always null for Database Insights
      },
      accounts: [{
        id: 'vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D',
        mask: '1234',
        name: null,
        type: 'depository',
        subtype: 'checking',
        verification_status: ''
      }]
    }
  },
  // ...
});

// Open Link on user-action
linkHandler.open();
```

##### Exchange the public token

In your own backend server, call the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange)
endpoint with the Link `public_token` received in the `onSuccess` callback to
obtain an `access_token`. Persist the returned `access_token` and `item_id` in your database
in relation to the user.

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

##### Fetch Auth data and verification results

Next, we can retrieve Auth data, along with the results of the Database Insights verification check, by calling [`/auth/get`](/docs/api/products/auth/#authget).

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
        "account": "9900009606",
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
      "name": null,
      "official_name": null,
      "verification_status": "database_insights_pass",
      "verification_insights": { 
        "network_status": {
          "has_numbers_match": true,
          "is_numbers_match_verified": true
        },
        "previous_returns": {
          "previous_administrative_return": false
        },
        "account_number_format": "valid"
      },
      "subtype": "checking",
      "type": "depository"
    }
  ],
  "item": { Object },
  "request_id": "m8MDnv9okwxFNBV"
}
```

#### Making a risk determination

The results of the Database Insights verification can be found in the `verification_status` and `verification_insights` fields returned by [`/auth/get`](/docs/api/products/auth/#authget). Based on the values in these fields, you will make a business decision on whether to accept the account numbers as verified, take additional risk mitigation steps, or reject the account numbers as unverified.

#### Testing Database Insights in Sandbox

For test credentials that can be used to test Database Insights in the Sandbox environment, see [Testing Database Auth or Database Insights](/docs/auth/coverage/testing/#testing-database-auth-or-database-insights).

#### Database Match

Database Insights and Database Match have been replaced with Database Auth. The documentation provided here is for the use of customers maintaining existing legacy Database Insights and Database Match integrations.

[Database Match](/docs/auth/coverage/database/#database-match) enables instant manual account verification without the need for micro-deposits, instead relying on Plaid's database of known account numbers. When provided as an alternative to Same-Day Micro-deposits, Database Match can increase conversion, as the user may be able to verify instantly, without having to return to Plaid to verify their micro-deposit codes.

Database Match will present the user with the option to manually add a bank account in Link by providing their name, account number, and routing number. Plaid will then check this information against its network of over 200 million known bank accounts. Approximately 30% of user bank accounts can be verified via Database Match. If a match is not found, Database Match will fall back to routing the user to a manual micro-deposit flow, either Instant Micro-deposits or Same-Day Micro-deposits.

Database Match can be used to verify the validity of the bank account being linked, but it does not verify that the end user has access to the bank account. To mitigate account takeover risk with Database Match, it can be paired with [Identity Match](/docs/identity/) to verify the user's information, such as phone number and address, on the linked account.

##### Database Match flow

1. Starting on a page in your app, the user clicks an action that opens Plaid Link.
2. Inside of Plaid Link, the user selects an option to enter the manual verification flow
   and provides their legal name, account and routing number.
3. Plaid will confirm the account number, routing number, and name match a previously verified account.
4. If a match is confirmed, Link closes with a `public_token` and a verification status of `database_matched`.
5. If there is no match, the user will be prompted to enter the Instant or Same Day Micro-deposits flow.

When these steps are complete, you can call [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to obtain an `access_token` for calling endpoints such as [`/auth/get`](/docs/api/products/auth/#authget) or [`/processor/token/create`](/docs/api/processors/#processortokencreate).

##### Database Match implementation steps

To enable Database Match, use the same settings as Same-Day Micro-deposits. You can then enable the feature in one of two ways:

- When calling [`/link/token/create`](/docs/api/link/#linktokencreate), set `auth.database_match_enabled: true`.
- [In the Dashboard Account Verification pane](https://dashboard.plaid.com/account-verification), enable Database Match.

Database Match cannot be used in the same session as [Database Insights](/docs/auth/coverage/database/). If Database Insights is explicitly enabled in the `auth` object alongside Database Match, an error will occur when calling [`/link/token/create`](/docs/api/link/#linktokencreate).

Database Match Configuration

```
const request: LinkTokenCreateRequest = {
  user: { client_user_id: new Date().getTime().toString() },
  client_name: 'Plaid App',
  products: [Products.Auth],
  country_codes: [CountryCode.Us],
  webhook: 'https://webhook.sample.com',
  language: 'en',
  auth: {
    database_match_enabled: true,
    same_day_microdeposits_enabled: true
  },
};
try {
  const response = await plaidClient.linkTokenCreate(request);
  const linkToken = response.data.link_token;
} catch (error) {
  // handle error
}
```

When calling [`/auth/get`](/docs/api/products/auth/#authget), the returned object will have a `verification_status` value of `database_matched` as an indication that the user's data was verified through Database Match.

Auth response for Database Match

```
{
  "numbers": {
    "ach": [
      {
        "account": "1111222233330000",
        "routing": "011401533",
        "wire_routing": "021000021"
      }
    ],
    ...
  },
  "accounts": [
    {
      "account_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
      "balances": { Object },
      "mask": "0000",
      "name": "Checking...0000",
      "official_name": null,
      "verification_status": "database_matched",
      "subtype": "checking",
      "type": "depository"
    }
  ],
  ...
}
```

##### Testing Database Match in Sandbox

For test credentials that can be used to test Database Match in the Sandbox environment, see [Testing Database Match](/docs/auth/coverage/testing/#testing-database-match).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
