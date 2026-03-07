---
title: "Payments (Europe) - Add Payment Initiation to your app | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/payment-initiation-one-time/add-to-app/"
scraped_at: "2026-03-07T22:05:11+00:00"
---

# Add One-Time Payment Initiation to your app

#### Learn how to use one-time Payment Initiation in your application

In this guide, we start from scratch and walk through how to set up a one-time [Payment Initiation](/docs/api/products/payment-initiation/) flow. For a high-level overview of all our payment related offerings, see the [Introduction to Payments (Europe)](/docs/payment-initiation/).

For an example guide to using one-time Payment Initiation as part of an integrated onboarding flow with other Plaid products, see [Onboarding and account funding guide](/docs/payment-initiation/payment-initiation-one-time/user-onboarding-and-account-funding/).

#### Get Plaid API keys and complete application and company profile

If you don't already have one, you'll need to [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You will also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile) on the Dashboard. The information in your profile will be shared with users of your application when they manage their connection on the [Plaid Portal](https://my.plaid.com), and must be completed before connecting to certain institutions.

#### Access Payment Initiation

Payments (Europe) is enabled in Sandbox by default. It uses test data and does not interact with financial institutions. You may need to request access to Payment Initiation via the [Dashboard](https://dashboard.plaid.com) before using it in Production.

#### Install and initialize Plaid libraries

You can use our official server-side client libraries to connect to the Plaid API from your application:

Terminal

```
// Install via npm
npm install --save plaid
```

After you've installed Plaid's client libraries, you can initialize them by passing in your `client_id`, `secret`, and the environment you wish to connect to (Sandbox or Production). This will make sure the client libraries pass along your `client_id` and `secret` with each request, and you won't need to include them in any other calls.

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

#### One-time Payment Initiation flow overview

Before we jump into the code, let's see an overview of the steps needed to set up your payment with Plaid:

**1**Your app collects information of the sender and the recipient, as well as the payment amount.

![Step 1 diagram](/assets/img/docs/payment_initiation/setup_flow/payment-setup-row-1.png)

**2**With this information, [**create a recipient**](/docs/api/products/payment-initiation/#payment_initiationrecipientcreate) and obtain a `recipient_id`. You can reuse this `recipient_id` for future payments.

![Step 2 diagram](/assets/img/docs/payment_initiation/setup_flow/payment-setup-row-2.png)

**3**Provide the `recipient_id` to Plaid when you [**create a payment**](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate). Store the resulting `payment_id` along with your payment metadata.

![Step 3 diagram](/assets/img/docs/payment_initiation/setup_flow/payment-setup-row-3.png)

**4**From your backend, use the `payment_id` to [**create a link\_token**](/docs/api/link/#linktokencreate).

![Step 4 diagram](/assets/img/docs/payment_initiation/setup_flow/payment-setup-row-4.png)

**5**Your client app uses the `link_token` to initiate a Link flow for your user. The onSuccess callback signals that the payment has been initiated.

![Step 5 diagram](/assets/img/docs/payment_initiation/PIS-overview-row-2.png)

**6**Your backend listens for [`PAYMENT_STATUS_UPDATE`](/docs/api/products/payment-initiation#payment_status_update) webhooks to keep track of the payment's status.

![Step 6 diagram](/assets/img/docs/payment_initiation/setup_flow/payment-setup-row-5.png)

##### Create a recipient

Before beginning the payment initiation process, you will need to know the name and account information of the recipient. With this information in hand, you can then call [`/payment_initiation/recipient/create`](/docs/api/products/payment-initiation/#payment_initiationrecipientcreate) to create a payment recipient and receive a `recipient_id`. You can later reuse this `recipient_id` for future payments to the same account.

```
const { PaymentInitiationRecipientCreateRequest } = require('plaid');

// Using BACS, without IBAN or address
const request: PaymentInitiationRecipientCreateRequest = {
  name: 'John Doe',
  bacs: {
    account: '26207729',
    sort_code: '560029',
  },
};
try {
  const response = await plaidClient.paymentInitiationRecipientCreate(request);
  const recipientID = response.data.recipient_id;
} catch (error) {
  // handle error
}
```

##### Create a payment

Now that you have the `recipient_id`, you can provide it together with the payment amount to [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate), which returns a `payment_id`.

```
const { PaymentInitiationPaymentCreateRequest } = require('plaid');

const request: PaymentInitiationPaymentCreateRequest = {
  recipient_id: recipientID,
  reference: 'TestPayment',
  amount: {
    currency: 'GBP',
    value: 100.0,
  },
};
try {
  const response = await plaidClient.paymentInitiationPaymentCreate(request);
  const paymentID = response.data.payment_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

Make sure to store the `payment_id` in your backend together with your own metadata (e.g. your internal `user_id`). You can use this later to process the [payment status updates](/docs/api/products/payment-initiation/#webhooks)

##### Create a link\_token

Before initializing Link, you need to create a new `link_token` on the server side of your application.
A `link_token` is a short-lived, one-time use token that is used to authenticate your app with Link.
You can create one using the [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint. Then, on the client side of your application, you'll need to initialize Link with the `link_token` that you created. This will bring up the payment initiation flow in Link that allows your end user to confirm the payment.

When calling [`/link/token/create`](/docs/api/link/#linktokencreate) to create a token for use with Payment Initiation, you provide the `payment_id` and specify `payment_initiation` as the product.

While it is possible to initialize Link with many products, `payment_initiation` cannot be specified along with any other products and must be the only product in Link's product array if it is being used.

In the code samples below, you will need to replace `PLAID_CLIENT_ID` and `PLAID_SECRET` with your own keys, which you can get from the [Dashboard](https://dashboard.plaid.com/developers/keys).

```
// Using Express
app.post('/api/create_link_token', async function (request, response) {
  const configs = {
    user: {
      // This should correspond to a unique id for the current user.
      client_user_id: 'user-id',
    },
    client_name: 'Plaid Test App',
    products: [Products.PaymentInitiation],
    language: 'en',
    webhook: 'https://webhook.sample.com',
    country_codes: [CountryCode.Gb],
    payment_initiation: {
      payment_id: paymentID,
    },
  };
  try {
    const createTokenResponse = await client.linkTokenCreate(configs);
    response.json(createTokenResponse.data);
  } catch (error) {
    // handle error
  }
});
```

Once you have exposed an endpoint to create a `link_token` in your application server, you now need to configure your client application to import and use Link.

#### Launch the payment initiation flow in Link

Plaid Link is a drop-in module that provides a secure, elegant authentication flow for the many financial institutions that Plaid supports. Link makes it secure and easy for users to connect their bank accounts to Plaid.
Because Link has access to all the details of the payment at the time of initialization, it will display a screen with the payment details already populated. All your end user has to do is log in to their financial institution through a Link-initiated OAuth flow, select a funding account, and confirm the payment details.

![Plaid Link payment flow: View payment details, select bank, authenticate with bank, payment confirmation with bank and transaction info.](/assets/img/docs/payment_initiation/payment-initiation-uk-link-screens.png)

Note that these instructions cover Link on the web. For instructions on using Link within mobile apps, see the [Link documentation](/docs/link/). If you want to customize Link's look and feel, you can do so from the [Dashboard](https://dashboard.plaid.com/link).

##### Install Link dependency

index.html

```
<head>
  <title>Link for Payment Initiation</title>
  <script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
</head>
```

##### Configure the client-side Link handler

Plaid communicates to you certain events that relate to how the user is interacting with Link. What you do with each of these event triggers depends on your particular use case, but a basic scaffolding might look like this:

app.js

```
const linkHandler = Plaid.create({
  // Create a new link_token to initialize Link
  token: (await $.post('/create_link_token')).link_token,
  onSuccess: (public_token, metadata) => {
    // Show a success page to your user confirming that the
    // payment will be processed.
    //
    // The 'metadata' object contains info about the institution
    // the user selected.
    // For example:
    //  metadata  = {
    //    link_session_id: "123-abc",
    //    institution: {
    //      institution_id: "ins_117243",
    //      name:"Monzo"
    //    }
    //  }
  },
  onExit: (err, metadata) => {
    // The user exited the Link flow.
    if (err != null) {
      // The user encountered a Plaid API error prior to exiting.
    }
    // 'metadata' contains information about the institution
    // that the user selected and the most recent API request IDs.
    // Storing this information can be helpful for support.
  },
  onEvent: (eventName, metadata) => {
    // Optionally capture Link flow events, streamed through
    // this callback as your users connect with Plaid.
    // For example:
    //  eventName = "TRANSITION_VIEW",
    //  metadata  = {
    //    link_session_id: "123-abc",
    //    mfa_type:        "questions",
    //    timestamp:       "2017-09-14T14:42:19.350Z",
    //    view_name:       "MFA",
    //  }
  },
});

linkHandler.open();
```

Unlike other products, for `payment_initiation` it is not necessary to exchange the `public_token` for an `access_token`. You only need the `payment_id` to interact with the `payment_initiation` endpoints.

#### Verify payment status

##### **Listening for status update webhooks**

Once the payment has been authorised by the end user and the Link flow is completed, the `onSuccess` callback is invoked, signaling that the payment status is now `PAYMENT_STATUS_INITIATED`.

From this point on, you can track the payment status using the [`PAYMENT_STATUS_UPDATE`](/docs/api/products/payment-initiation/#payment_status_update) webhook that is triggered by Plaid when updates occur:

- Updates are sent by Plaid to the [configured webhook URL](/docs/api/link/#link-token-create-request-webhook) to indicate that the status of an initiated payment has changed. All Payment Initiation webhooks have a [`webhook_type`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-webhook-type) of `PAYMENT_INITIATION`.
- Once you receive the status update, use the [`payment_id`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-payment-id) field to retrieve the payment's metadata from your database.
- From the status update object, use the [`new_payment_status`](/docs/api/products/payment-initiation/#PaymentStatusUpdateWebhook-new-payment-status) field to decide what action needs to be taken for this payment. For example, you may decide to fund the account once the payment status is `PAYMENT_STATUS_EXECUTED` or notify the user that their payment got rejected (`PAYMENT_STATUS_REJECTED`).

`PAYMENT_STATUS_SETTLED` is the only status that can be used for releasing funds to end users or merchants. For all other statuses, including `PAYMENT_STATUS_INITIATED` and `PAYMENT_STATUS_EXECUTED`, you must wait until you have confirmed the funds have arrived in your own bank account, or until the payment transitions to `PAYMENT_STATUS_SETTLED`.

For more information, see [Payment Status](/docs/payment-initiation/payment-status/).

For more information on how to implement your webhook listener, see the [webhooks documentation](/docs/api/webhooks/).

#### Sample code in Plaid Pattern

For an example implementation of Payment Initiation, see the [Payment Initiation Pattern App](https://github.com/plaid/payment-initiation-pattern-app) on GitHub, which uses Payment Initiation in an account funding use case.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
