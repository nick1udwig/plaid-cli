---
title: "Auth - Instant Auth, Instant Match, & Instant Micro-deposits | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/instant/"
scraped_at: "2026-03-07T22:04:32+00:00"
---

# Instant Auth, Instant Match, and Instant Micro-Deposits

#### Learn how to authenticate your users instantly

#### Instant Auth

Instant Auth supports more than 6,200 financial institutions
with credential-based login. Instant Auth is the default Auth flow and does not require extra configuration steps if Auth is already configured in your app. For clarity and completeness, the section below explains how to configure Instant Auth.

The user launches Link...

![The user launches Link...](/assets/img/docs/auth/ia_tour/link-1-rm.png)

![...and selects the institution to link.](/assets/img/docs/auth/ia_tour/select-institution.png)

![They are handed off to the institution...](/assets/img/docs/auth/ia_tour/oauth-handoff.png)

![...and log in.](/assets/img/docs/auth/ia_tour/link-gingham.png)

![Next, they select the account to link. (Single-account select is typically recommended for Auth.)](/assets/img/docs/auth/ia_tour/account-select.png)

![...and their account is linked!](/assets/img/docs/auth/ia_tour/link-success.png)

You can try out the Instant Auth flow in an [Interactive Demo](https://plaid.coastdemo.com/share/67d0ce0df465686c02cc4fd2?zoom=100&step=4).

##### Configure & Create a link\_token

Create a `link_token` with the following parameters:

- `products` array containing `auth` – If you are using only `auth` and no other products, `auth` must be specified in the Products array. Other products (such as `identity`) may be specified as well. If you are using multiple products, `auth` is not required to be specified in the products array, but including it is recommended for the best user experience.

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

##### Initialize Link with a link\_token

After creating a `link_token` for the `auth` product, use it to initialize Plaid Link.

When the user inputs their username and password for the financial institution,
the `onSuccess()` callback function will return a `public_token`.

App.js

```
Plaid.create({
  // Fetch a link_token configured for 'auth' from your app server
  token: (await $.post('/create_link_token')).link_token,
  onSuccess: (public_token, metadata) => {
    // Send the public_token and accounts to your app server
    $.post('/exchange_public_token', {
      publicToken: public_token,
      accounts: metadata.accounts,
    });
  },
});
```

##### Exchange the public\_token and fetch Auth data

In your own backend server, call the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange)
endpoint with the Link `public_token` received in the `onSuccess` callback to obtain an `access_token`.
Persist the returned `access_token` and `item_id` in your database in relation to the user. You will use
the `access_token` when making requests to the [`/auth/get`](/docs/api/products/auth/#authget) endpoint.

Exchange token and fetch Auth data

```
const publicToken = 'public-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d';

try {
  // Obtain an access_token from the Link public_token
  const tokenResponse = await client.itemPublicTokenExchange({
      public_token: publicToken});
  const accessToken = tokenResponse.access_token;

  // Instantly fetch Auth numbers
  const request: AuthGetRequest = {
    access_token: accessToken,
  };
  const response = await plaidClient.authGet(request);
  const numbers = response.data.numbers;
} catch (err) {
  // handle error
}
```

Check out the [`/auth/get`](/docs/api/products/auth/#authget) API reference documentation to see the full
Auth request and response schema.

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
  "accounts": [{ Object }],
  "item": { Object },
  "request_id": "m8MDnv9okwxFNBV"
}
```

#### Instant Match

Instant Match is available for approximately 800 U.S. additional financial institutions where Instant Auth is not available. Instant Match is enabled automatically for Auth customers and is automatically
provided at supported institutions as a fallback experience when Instant Auth is not available. When using Instant Match, Plaid Link will
prompt your user to enter their account number and routing number for a depository account. Plaid will then verify the last four digits of
the user-provided account number against the account mask retrieved from the financial institution.

The user launches Link...

![The user launches Link...](/assets/img/docs/auth/im_tour/link-1-rm.png)

![...and selects the institution to link.](/assets/img/docs/auth/im_tour/select-institution.png)

![They find their institution...](/assets/img/docs/auth/im_tour/fi-selected.png)

![...and log in.](/assets/img/docs/auth/im_tour/fi-login.png)

![Next, they select the account to link. (Single-account select is typically recommended for Auth.)](/assets/img/docs/auth/im_tour/account-select.png)

![They verify their bank's routing number...](/assets/img/docs/auth/im_tour/im-verify-routing.png)

![...and their account number.](/assets/img/docs/auth/im_tour/im-verify-account.png)

![If those numbers match the account masks that Plaid has retrieved, the account is verified](/assets/img/docs/auth/im_tour/link-success.png)

You can try out the Instant Match flow in an [Interactive Demo](https://plaid.coastdemo.com/share/67d0ce0df465686c02cc4fd2?zoom=100&step=4). See more details in our [testing guide](/docs/auth/coverage/testing/#testing-instant-match).

When using the Instant Match flow, the user can verify only a single account. Even if the Account Select properties allow selecting all or multiple accounts, the ability to select multiple depository accounts for Auth will be disabled in Link if the institution is using the Instant Match flow.

##### Configuring in Link

Instant Match will be enabled automatically if you configure the `link_token` with the following parameters:

- add `"auth"` to `products` array
- `country_codes` set to `['US']` (adding any other countries to the array will disable Instant Match)

Optionally, you can disable Instant Match on a per-session basis via the [`/link/token/create`](/docs/api/link/#linktokencreate) call, by setting `"auth.instant_match_enabled": false` in the request body. If you would like to disable Instant Match automatically for all Link sessions, contact your Account Manager or file a support ticket via the Dashboard.

Instant Match Configuration

```
const request: LinkTokenCreateRequest = {
  user: { client_user_id: new Date().getTime().toString() },
  client_name: 'Plaid App',
  products: [Products.Auth],
  country_codes: [CountryCode.Us],
  language: 'en',
};
try {
  const response = await plaidClient.linkTokenCreate(request);
  const linkToken = response.data.link_token;
} catch (error) {
  // handle error
}
```

##### Handling Link events

For a user who goes through the Instant Match flow, the `TRANSITION_VIEW (view_name = NUMBERS)` event will occur after `SUBMIT_CREDENTIALS`, and in the `onSuccess` callback the `verification_status` will be `null` because the user would have been verified instantly.

Sample Link events for Instant Match flow

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
onSuccess (verification_status: null)
```

#### Instant Micro-deposits

Instant Micro-deposits is the Plaid product term for our ability to authenticate any bank account in the US that is supported by RTP or FedNow. For over 80 Plaid-supported banks, Instant Micro-deposits is the fastest and highest-converting form of Auth support available. If both Instant Micro-deposits and Same Day Micro-deposits are enabled, any user who attempts a micro-deposit with one of the over 400 eligible RTP or FedNow routing numbers will automatically experience the Instant Micro-deposits flow and be able to verify instantly.

##### Instant Micro-deposit flow

The user launches Link...

![The user launches Link...](/assets/img/docs/auth/imd_tour/link-1-rm.png)

![...and selects the institution to link.](/assets/img/docs/auth/imd_tour/select-institution.png)

![They find their institution...](/assets/img/docs/auth/imd_tour/imd-select-fi.png)

![...enter their account info...](/assets/img/docs/auth/imd_tour/imd_enter_account_info.png)

![...their name...](/assets/img/docs/auth/imd_tour/imd_enter_name.png)

![...and their account type.](/assets/img/docs/auth/imd_tour/imd_enter_account_type.png)

![The user authorizes Plaid to make a deposit, and the deposit is made immediately.](/assets/img/docs/auth/imd_tour/imd_authorize_deposit.png)

![The user is then prompted to enter the code associated with the deposit.](/assets/img/docs/auth/imd_tour/imd_check_account.png)

![If the code matches the one that Plaid generated with the deposit description, the user's account is verified.](/assets/img/docs/auth/imd_tour/link-success.png)

1. Starting on a page in your app, the user clicks an action that opens Plaid Link.
2. Inside of Plaid Link, the user enters the micro-deposit initiation flow
   and provides their legal name, account and routing number.
3. Plaid sends a micro-deposit to the user's account that will post within 5 seconds, and directs the user to log into their bank account to obtain the code from the micro-deposit description.
4. The user enters the code from the micro-deposit description into Plaid Link.
5. Upon success, Link closes with a `public_token` and a `metadata` account status of `manually_verified`.

Plaid will not reverse the $0.01 micro-deposit credit.

When these steps are done, your user's Auth data is verified and ready to fetch.

You can try out the Instant Micro-deposits flow in an [Interactive Demo](https://plaid.coastdemo.com/share/67d0ce0df465686c02cc4fd2?zoom=100&step=7).

##### Configuring in Link

Instant Micro-deposits will be enabled if you configure the `link_token` with the following parameters:

- Set the `products` array to `["auth"]`. While in most cases additional products can be added to existing Plaid Items, Items created for micro-deposit verification cannot be used with any Plaid products other than Auth or Transfer, with the exception that approximately 30% of Items verified by Instant micro-deposits can also be verified by [Identity Match](/docs/identity/#using-identity-match-with-micro-deposit-or-database-items) and evaluated for [Signal Transaction Scores](/docs/signal/signal-rules/#data-availability-limitations).
- `country_codes` set to `['US']` (adding any other countries to the array will disable Instant Micro-deposits)
- `auth.instant_microdeposits_enabled` set to `true`. For Plaid teams created prior to November 2023 this setting is not required; for newer teams, it must be manually configured.

Optionally, you can disable Instant Match on a per-session basis via the [`/link/token/create`](/docs/api/link/#linktokencreate) call, by setting `"auth.instant_microdeposits_enabled": false` in the request body. If you would like to disable Instant Match automatically for all Link sessions, contact your Account Manager or file a support ticket via the Dashboard.

##### Entering the Instant Micro-deposit flow

Your user will enter the Instant Micro-deposit flow in the following scenarios:

- The user selects an eligible institution that is not enabled for Instant Auth, Instant Match, or Automated Micro-deposits.
- The Link session has [Same Day Micro-deposits](/docs/auth/coverage/same-day/) enabled and the user enters an eligible routing number during the Same Day Micro-deposits flow. In this case, the session will be "upgraded" to use Instant Micro-deposits rather than Same-Day Micro-deposits.

###### Instant Micro-deposit events

When a user goes through the Instant micro-deposits flow, the session will have the `TRANSITION_VIEW` (`view_name = NUMBERS`) event and a `TRANSITION_VIEW` (`view_name = INSTANT_MICRODEPOSIT_AUTHORIZED`) event after the user authorizes Plaid to send a micro-deposit to the submitted account and routing number. In the `onSuccess` callback the `verification_status` will be `manually_verified` since, unlike Same Day Micro-deposits, Instant Micro-deposits will resolve to either a success or fail state within a single Link session.

##### Testing the Instant Micro-deposit flow

For credentials that can be used to test Instant Micro-deposits in Sandbox, see [Auth testing flows](/docs/auth/coverage/testing/#testing-instant-micro-deposits).

[#### Automated Micro-deposits

Integrate the automated micro-deposit flow

View guide](/docs/auth/coverage/automated/)

#### Automated Micro-deposits

Integrate the automated micro-deposit flow

[View guide](/docs/auth/coverage/automated/)[#### Same Day Micro-deposits

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
