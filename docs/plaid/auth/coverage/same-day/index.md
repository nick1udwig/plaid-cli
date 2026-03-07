---
title: "Auth - Same Day Micro-deposits | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/same-day/"
scraped_at: "2026-03-07T22:04:32+00:00"
---

# Same Day Micro-deposits

#### Learn how to authenticate your users with a manually verified micro-deposit

### Overview

Same Day Micro-deposits can be used to authenticate any bank account in the US, but especially for the ~2,000 institutions
that don't support Instant Auth, Instant Match, or Automated Micro-deposit verification. Plaid will make a deposit that will
post within one business day (using Same Day ACH, which is roughly two days faster than the standard micro-deposit experience
of two to three days). Users are instructed to manually verify the code in the transaction description deposited in the account.

#### The Same Day Micro-deposit flow

The user clicks an action that opens Plaid Link.

![The user clicks an action that opens Plaid Link.](/assets/img/docs/auth/smd_tour/smd_plaid_start.png)

![They optionally log in...](/assets/img/docs/auth/smd_tour/smd_remember_me.png)

![...select a new institution...](/assets/img/docs/auth/smd_tour/smd_select_fi_1.png)

![...and choose to connect by entering account numbers manually.](/assets/img/docs/auth/smd_tour/smd_select_fi.png)

![The user enters their routing number...](/assets/img/docs/auth/smd_tour/smd_enter_routing.png)

![...account number...](/assets/img/docs/auth/smd_tour/smd_enter_account.png)

![...full name...](/assets/img/docs/auth/smd_tour/smd_enter_name.png)

![...selects an account type...](/assets/img/docs/auth/smd_tour/smd_select_account.png)

![...and authorizes the transfer.](/assets/img/docs/auth/smd_tour/smd_authorize.png)

![The user enters their phone number to be notified when the transfer has arrived.](/assets/img/docs/auth/smd_tour/smd_enter_phone_number.png)

![The user is told to come back in a few days.](/assets/img/docs/auth/smd_tour/smd_deposit_initiated.png)

![A day or two later, the user receives an SMS notification that the deposit has arrived.](/assets/img/docs/auth/smd_tour/smd_sms.png)

![After getting the code from their bank account, they return to Plaid to verify the code.](/assets/img/docs/auth/smd_tour/smd_verify_1.png)

![And the account is verified](/assets/img/docs/auth/smd_tour/smd_verify_2.png)

A user connects their financial institution using the following connection flow:

1. Starting on a page in your app, the user clicks an action that opens Plaid Link with the correct
   Auth [configuration](/docs/auth/coverage/same-day/#create-a-link_token).
2. Inside of Plaid Link, the user enters the micro-deposit initiation flow
   and provides their legal name, account and routing number.
3. Upon [successful authentication](/docs/auth/coverage/same-day/#exchange-the-public-token), Link closes with a `public_token`
   and a `metadata` account status of `pending_manual_verification`.
4. Behind the scenes, Plaid sends a micro-deposit to the user's account that will post within one to two business days.
5. After one to two days, the user is prompted to verify the code in the transaction description in their account, by
   [opening Link with a generated `link_token`](/docs/auth/coverage/same-day/#prompt-user-to-verify-micro-deposit-code-in-link).

Plaid will not reverse the $0.01 micro-deposit credit.

When these steps are done, your user's Auth data is verified and [ready to fetch](/docs/auth/coverage/same-day/#fetch-auth-data).

##### Demoing the flow in Link

You can try out the Same Day Micro-deposits flow in an [Interactive Demo](https://plaid.coastdemo.com/share/67d0ce0df465686c02cc4fd2?zoom=100&step=9). For more details, see the [testing guide](/docs/auth/coverage/testing/#testing-same-day-micro-deposits).

#### Implementation steps

##### Enable Same Day micro-deposits

Enable Same Day micro-deposits via the [Account Verification Dashboard](https://dashboard.plaid.com/account-verification). Alternatively, you can also enable this flow by setting the `auth.same_day_microdeposits_enabled: true` parameter when calling [`/link/token/create`](/docs/api/link/#linktokencreate).

##### Create a link\_token

Create a `link_token` with the following parameters:

- `products` array should include only `auth` or `transfer` as a product when using same-day manual micro-deposit verification. While in most cases additional products can be added to existing Plaid Items, Items created for Same Day manual micro-deposit verification are an exception and cannot be used with any Plaid products other than Auth or Transfer.

Approximately 30% of Items verified by Same Day micro-deposits can also be verified by [`/identity/match`](/docs/api/products/identity/#identitymatch) or [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate). If using Identity Match or Signal Transaction Scores in this way, they should be added to the Item via the `required_if_supported_products`, `optional_products`, or `additional_consented_products` fields rather than the `products` array. For more details, see [Identity Match](/docs/identity/#identity-match) and [Signal Transaction Scores](/docs/signal/signal-rules/#data-availability-limitations). All Items verified by Same Day micro-deposits are also compatible with statement-based verification via [Identity Document Upload](/docs/identity/identity-document-upload/).

- `country_codes` set to `['US']` – Micro-deposit verification is currently only available in the United States.

##### Initialize Link with a link\_token

After creating a `link_token` for the `auth` product, use it to initialize Plaid Link.

When the user successfully inputs their account and routing numbers, the `onSuccess()` callback
function (or the equivalent field in [`/link/token/get`](/docs/api/link/#linktokenget), if using the [Hosted Link](/docs/link/hosted-link/) integration method) will return a `public_token`, with `verification_status` equal to `'pending_manual_verification'`.

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
        name: null,          // name is always null for Same Day Micro-deposits
        institution_id: null // institution_id is always null for Same Day Micro-deposits
      },
      accounts: [{
        id: 'vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D',
        mask: '1234',
        name: "Checking...1234",
        type: 'depository',
        subtype: 'checking',
        verification_status: 'pending_manual_verification'
      }]
    }
  },
  // ...
});

// Open Link on user-action
linkHandler.open();
```

##### Display a "pending" status in your app

Because Same Day verification usually takes one business day to complete, we recommend displaying a UI
in your app that communicates to a user that verification is currently pending.

You can use the `verification_status` key returned in the `onSuccess` `metadata.accounts` object once
Plaid Link closes successfully.

Metadata verification\_status

```
verification_status: 'pending_manual_verification';
```

You can also [fetch the `verification_status`](/docs/auth/coverage/same-day/#check-the-account-verification-status-optional) for an
Item's account via the Plaid API, to obtain the latest account status.

##### Exchange the public token

In your own backend server, call the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange)
endpoint with the Link `public_token` to
obtain an `access_token`.

When using same-day micro-deposit verification, only one account can be associated with each access token. If you want to allow a user to link multiple accounts at the same institution using same-day micro-deposits, you will need to create a new Link flow and generate a separate access token for each account.

To test your integration outside of Production, see [Testing Same Day Micro-deposits in Sandbox](/docs/auth/coverage/testing/#testing-same-day-micro-deposits).

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

##### Check the account verification status *(optional)*

In some cases you may want to implement logic in your app to display the `verification_status` of
an Item that is pending manual verification. The [`/accounts/get`](/docs/api/accounts/#accountsget)
API endpoint allows you to query this information.

To be notified via webhook when Plaid has sent the micro-deposit to your end user, see [micro-deposit events](/docs/auth/coverage/microdeposit-events/).

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
      "name": "Checking...0000",
      "official_name": null,
      "type": "depository",
      "subtype": "checking",
      "verification_status":
        "pending_manual_verification" |
        "manually_verified" |
        "verification_failed"
    },
    ...
  ],
  "item": { Object },
  "request_id": String
}
```

##### Prompt user to verify micro-deposit code in Link

After one to two business days, the micro-deposit sent to the user's account is expected to be posted.
To securely verify a Same Day Micro-deposits account, your user needs to come back into Link to verify
the code in the transaction description.

When the micro-deposit posts to your end user's bank account, the transaction description will be written with the format:

Micro-deposit post description

```
#XXX <clientName> ACCTVERIFY
```

The `#` will be followed with the three letter code required for verification. The `<clientName>` is defined by the value of
the `client_name` parameter that was used to create the `link_token` that initialized Link. Due to network requirements, the `client_name` will be truncated to the first 11 characters and `ACCTVERIFY` will be added to signify the deposit is for account verification.

Users with business or corporate accounts that have ACH debit blocks enabled on
their account may need to authorize Plaid's Company / Tax ID, `1460820571`, to
avoid any issues with linking their accounts.

To optimize conversion, we strongly recommend sending
your user a notification (e.g. email, SMS, push notification) prompting them to come back into your
app and verify the micro-deposit code. To be notified via webhook when Plaid has sent the micro-deposit
to your end user, see [micro-deposit events](/docs/auth/coverage/microdeposit-events/).

![Plaid Instant Match process: Enter 3-letter code from deposit, confirmation screen, success message linking to Wonderwallet.](/assets/img/docs/auth/manual_md_2.png)

Verification of Same Day Micro-deposits is performed in two steps:

1. In your backend server, create a new `link_token` from the associated `access_token` for
   the given user.
2. Pass the generated `link_token` into your client-side app, using the `token` parameter in
   the Link configuration. This will automatically trigger the micro-deposit verification flow in Link.

##### Create a new link\_token from a persistent access\_token

Generate a `link_token` for verifying micro-deposits by passing the user's associated `access_token` to the
[`/link/token/create`](/docs/api/link/#linktokencreate) API endpoint. Note that the `products` field should not be set because the micro-deposits verification flow does not change the products associated with the given `access_token`.

```
// Using Express
app.post('/api/create_link_token', async function (request, response) {
  // Get the client_user_id by searching for the current user
  const user = await User.find(...);
  const clientUserId = user.id;
  const linkTokenRequest = {
    user: {
      client_user_id: clientUserId,
    },
    client_name: 'Plaid Test App',
    language: 'en',
    webhook: 'https://webhook.sample.com',
    country_codes: [CountryCode.Us],
    access_token: 'ENTER_YOUR_ACCESS_TOKEN',
  };
  try {
    const createTokenResponse = await client.linkTokenCreate(linkTokenRequest);
    response.json(createTokenResponse.data);
  } catch (error) {
    // handle error
  }
});
```

##### Initialize Link with the generated `link_token`

In your client-side app, pass the generated `link_token` into the Link `token` parameter. Link will
automatically detect that Same Day verification is required for the Item and will open directly into
the verification flow (see the image above).

In Link, the user will be prompted to log in to their personal banking portal to confirm the code in the
micro-deposit transaction description. Upon successful entry of the code, the `onSuccess` callback will be fired, with an
updated `verification_status: 'manually_verified'`. The verification code will be case-insensitive.

There is no time limit for the user to verify the deposit. A user has three attempts to enter the code correctly, after which the Item will be permanently locked for security reasons.
See [INCORRECT\_DEPOSIT\_VERIFICATION](/docs/errors/invalid-input/#incorrect_deposit_verification) and [PRODUCT\_NOT\_READY](/docs/errors/item/#product_not_ready) for errors that may
occur during the micro-deposit initiation and verification flow.

App.js

```
const linkHandler = Plaid.create({
  token: await fetchLinkTokenForMicrodepositsVerification(),
  onSuccess: (public_token, metadata) => {
    metadata = {
      accounts: [{
        ...,
        verification_status: 'manually_verified',
      }],
    };
  },
  // ...
});

// Open Link to verify micro-deposit amounts
linkHandler.open();
```

An Item's `access_token` does not change when verifying micro-deposits, so there is no need to repeat
the exchange token process.

##### Fetch Auth data

Finally, we can retrieve Auth data once the user has manually verified their account through Same Day Micro-deposits:

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
      "name": "Checking ...0000",
      "official_name": null,
      "verification_status": "manually_verified",
      "subtype": "checking",
      "type": "depository"
    }
  ],
  "item": { Object },
  "request_id": "m8MDnv9okwxFNBV"
}
```

Check out the [`/auth/get`](/docs/api/products/auth/#authget) API reference documentation to see the full
Auth request and response schema.

#### Using Text Message Verification

Text Message Verification is an alternative verification method for the Same Day Micro-deposit flow. With Text Message Verification, Plaid will send your user a one-time SMS message, directing them to a Plaid-hosted website where they can complete the micro-deposit verification process. When the user is done verifying their micro-deposit code, you will receive a [`SMS_MICRODEPOSITS_VERIFICATION`](/docs/api/products/auth/#sms_microdeposits_verification) webhook, telling you that the user has completed the process and that it is now safe to retrieve Auth information.

Text Message Verification can and should be used alongside the usual verification flow of prompting your user to verify their code inside your app through Link. The user may not be prompted to receive the SMS message (if they clicked "continue as guest"), may choose not to receive an SMS message, or they might simply ignore the message, so it is important for your app to still provide a way for your user to complete the process.

##### Implementation steps

Text Message Verification is enabled by default as long as Same Day Micro-deposits have been enabled. To opt out of Text Message Verification, use the [Dashboard Account Verification pane](https://dashboard.plaid.com/account-verification) to disable it, or, if not using the Account Verification Dashboard, set `auth.sms_microdeposits_verification_enabled: false` in your [`/link/token/create`](/docs/api/link/#linktokencreate) call.

1. When calling [`/link/token/create`](/docs/api/link/#linktokencreate), make sure you have specified a URL for your webhook receiver, so you can receive the [`SMS_MICRODEPOSITS_VERIFICATION`](/docs/api/products/auth/#sms_microdeposits_verification) webhook.
2. Listen for the [`SMS_MICRODEPOSITS_VERIFICATION`](/docs/api/products/auth/#sms_microdeposits_verification) webhook.

   When the user completes the verification process, Plaid will send a [`SMS_MICRODEPOSITS_VERIFICATION`](/docs/api/products/auth/#sms_microdeposits_verification) webhook to the webhook receiver URL that you specified earlier. When you receive this webhook, review the value of the `status` field.

   Example webhook

   ```
   {
     "webhook_type": "AUTH",
     "webhook_code": "SMS_MICRODEPOSITS_VERIFICATION",
     "status": "MANUALLY_VERIFIED",
     "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
     "account_id": "dVzbVMLjrxTnLjX4G66XUp5GLklm4oiZy88yK",
     "environment": "sandbox"
   }
   ```

   A value of `MANUALLY_VERIFIED` indicates that the user successfully entered the micro-deposit code and has verified their account information. You can now retrieve Auth information on behalf of this user, and you should remove any pending in-app messages asking the user to complete the verification process.

   If you re-open Link and ask the user to verify their code after they have already verified it using Text Message Verification, Link will close immediately and fire the `onSuccess` callback. So even if you don't act on this webhook, your application will continue to function normally.

   A `status` field of `VERIFICATION_FAILED` indicates that the user failed the verification process. Verification cannot be retried once this status has been triggered; you will need to create a new Item.

##### User experience

When the user goes through the Same Day Micro-deposit flow in Link, they will be prompted to enter their phone number, as long as they did not opt out of using or creating a Plaid account. Users who skip the Plaid account process and click "Continue as guest" instead will not be prompted to enter their phone number and cannot participate in Text Message Verification. After the micro-deposit has been placed in their account, if the user authorized Plaid to send the text message, Plaid will contact the user via SMS with a URL pointing to a Plaid-hosted page where the user can complete the verification process. The text message itself will contain the following message:

Sample SMS message

```
Plaid: On behalf of [client_name], a $0.01 deposit was sent to your account ending in 1234. Verify this deposit here: https://hosted.plaid.com/link/lp1234. Then, return to [client_name] to complete your account setup.
```

Currently, the text message is only provided in English and will not be localized according to your Link customization settings.

##### Testing text message verification

Text message verification cannot be tested in the Sandbox environment. Text messages will only be sent in Production, and will only be sent to users who do not click "Continue as guest" on the Link consent pane.

#### Same Day Micro-deposit flow configuration options

In addition to the default flow, Same Day Micro-deposits has several optional flow settings you can enable.

To expose more users to the Same Day micro-deposit flow, you can enable [Auth Type Select](/docs/auth/coverage/flow-options/#adding-manual-verification-entry-points-with-auth-type-select), or to limit users' exposure to the flow, you can enable [Reroute to Credentials](/docs/auth/coverage/flow-options/#removing-manual-verification-entry-points-with-reroute-to-credentials).

To provide an alternative flow that allows users to skip micro-deposit verification and instead relies on recognizing a known bank account within the Plaid network, you can enable [Database Auth](/docs/auth/coverage/database-auth/).

The setting that is best for you will depend on your use case, your risk exposure, and which other Plaid products you use. Learn more about how to optimize your configuration and manage risk under [best practices](/docs/auth/coverage/same-day-link-best-practices/).

Same Day Micro-deposit flow options are configured on a Link customization level (if using the Account Verification Dashboard) or on a Link token level (if configuring the options directly in the [`/link/token/create`](/docs/api/link/#linktokencreate) call). This enables you to decide which sessions are enabled for which flows; for example, you can enable different flows based on users' risk profiles.

#### Handling Link events

When a user goes through the Same Day micro-deposits flow, the session will have the `TRANSITION_VIEW (view_name = NUMBERS)` event and a `TRANSITION_VIEW` (`view_name = SAME_DAY_MICRODEPOSIT_AUTHORIZED`) event after the user authorizes Plaid to send a micro-deposit to the submitted account and routing number. In the `onSuccess` callback the `verification_status` will be `pending_manual_verification` because the user will have to return to Link to verify their micro-deposit at a later Link session.

Sample Link events for Same Day micro-deposits where user enters flow from empty Search state

```
OPEN (view_name = CONSENT)
TRANSITION_VIEW (view_name = SELECT_INSTITUTION)
SEARCH_INSTITUTION
TRANSITION_VIEW (view_name = NUMBERS)
TRANSITION_VIEW (view_name = LOADING)
TRANSITION_VIEW (view_name = CONNECTED)
HANDOFF
onSuccess (verification_status: pending_manual_verification)
```

When a user goes through the Same Day micro-deposits flow with Reroute to Credentials, you will additionally see `TRANSITION_VIEW (view_name = NUMBERS_SELECT_INSTITUTION)` with `view_variant = SINGLE_INSTITUTION` or `view_variant = MULTI_INSTITUTION`.

Sample Link events for Same Day micro-deposits flow where user encounters Reroute to Credentials

```
OPEN (view_name = CONSENT)
TRANSITION_VIEW (view_name = SELECT_INSTITUTION)
SEARCH_INSTITUTION
TRANSITION_VIEW (view_name = NUMBERS)
TRANSITION_VIEW (view_name = NUMBERS_SELECT_INSTITUTION, view_variant = SINGLE_INSTITUTION)
TRANSITION_VIEW (view_name = LOADING)
TRANSITION_VIEW (view_name = CONNECTED)
HANDOFF
onSuccess (verification_status: pending_manual_verification)
```

When a user goes through the Same Day micro-deposits flow with the Auth Type Select configuration, you will additionally see `TRANSITION_VIEW (view_name = SELECT_AUTH_TYPE)` and also `SELECT_AUTH_TYPE (selection = flow_type_manual)`

Sample Link events for Same Day micro-deposits flow where user enters flow from Auth Type Select

```
OPEN (view_name = CONSENT)
TRANSITION_VIEW (view_name = SELECT_AUTH_TYPE)
SELECT_AUTH_TYPE (selection = flow_type_manual)
TRANSITION_VIEW (view_name = NUMBERS)
TRANSITION_VIEW (view_name = LOADING)
TRANSITION_VIEW (view_name = CONNECTED)
HANDOFF
onSuccess (verification_status: pending_manual_verification)
```

[#### Testing in Sandbox

Learn how to test each Auth flow in the Sandbox

View guide](/docs/auth/coverage/testing/)

#### Testing in Sandbox

Learn how to test each Auth flow in the Sandbox

[View guide](/docs/auth/coverage/testing/)[#### Manual verification flow best practices

Minimize fraud by following best practices

View guide](/docs/auth/coverage/same-day-link-best-practices/)

#### Manual verification flow best practices

Minimize fraud by following best practices

[View guide](/docs/auth/coverage/same-day-link-best-practices/)[#### Micro-deposit events

Learn how to use webhooks to receive micro-deposit status updates

View guide](/docs/auth/coverage/microdeposit-events/)

#### Micro-deposit events

Learn how to use webhooks to receive micro-deposit status updates

[View guide](/docs/auth/coverage/microdeposit-events/)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
