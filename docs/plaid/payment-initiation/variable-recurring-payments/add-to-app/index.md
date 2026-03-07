---
title: "Payments (Europe) - Add Variable Recurring Payments to your app | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/variable-recurring-payments/add-to-app/"
scraped_at: "2026-03-07T22:05:12+00:00"
---

# Add Variable Recurring Payments to your app

#### Learn how to use Variable Recurring Payments in your application

In this guide, we start from scratch and walk through how to set up a Variable Recurring Payments (VRP) flow. For a high-level overview of all Plaid's European payment offerings, see [Introduction to Payments (Europe)](/docs/payment-initiation/).

#### Variable Recurring Payments flow overview

Variable Recurring Payments consist of two main phases: creating a consent, and making payments using the authorised consent. The Plaid flow begins when your user wants to set up a consent with your app to make regular payments.

The sections below outline the general flow for VRP. The rest of the guide will cover the end-to-end process in more detail, including sample code and instructions for creating your Plaid developer account.

##### Creating a consent

1. If you haven't already done so, call [`/payment_initiation/recipient/create`](/docs/api/products/payment-initiation/#payment_initiationrecipientcreate), specifying at least a `name` and a `bacs` or `iban`, to create the `recipient_id` of the funds. You can re-use this `recipient_id` for other payment consents in the future.
2. Call [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate), specifying the `recipient_id` of the payment and the type and limitations of the consent, including the maximum amount and frequency, and when the consent expires. This endpoint will return a `consent_id`, which you should store associated with the end user who created it. You will need this value when making a payment.
3. Call [`/link/token/create`](/docs/api/link/#linktokencreate), passing in the `consent_id`. This creates a `link_token` containing all the information needed to display the correct details in Link to your end user.
4. Pass the `link_token` to your client and launch Link on the client side.
5. The end user goes through the Link flow. During this flow, they will consent to the terms as specified in the [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate) call in step 2.

##### Making payments using the consent

Call [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute), passing in the details of the specific payment, such as the amount and the `consent_id`.
Plaid will fire `PAYMENT_STATUS_UPDATE` webhooks as the payment is processed, to allow you to track the status of the payment.

##### Revoking consent (optional)

End users can revoke consent via either their bank or the [my.plaid.com](https://my.plaid.com) Portal. You can also allow them to revoke consent within your app via [`/payment_initiation/consent/revoke`](/docs/api/products/payment-initiation/#payment_initiationconsentrevoke).

#### Get Plaid API keys and complete application and company profile

If you don't already have one, you'll need to [create a Plaid developer account](https://dashboard.plaid.com/signup). After creating your account, you can find your [API keys](https://dashboard.plaid.com/developers/keys) under the Team Settings menu on the Plaid Dashboard.

You will also need to complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile) on the Dashboard. The information in your profile will be shared with users of your application when they manage their connection on the [Plaid Portal](https://my.plaid.com), and must be completed before connecting to certain institutions.

#### Access Variable Recurring Payments

Variable Recurring Payments is enabled in Sandbox by default. It uses test data and does not interact with financial institutions. For Production access, contact your account manager or [sales](https://plaid.com/contact/).

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

#### Setting up the payment

##### Creating a recipient

To create a recipient, call [`/payment_initiation/recipient/create`](/docs/api/products/payment-initiation/#payment_initiationrecipientcreate). You must provide a `name` and either an `iban` or `bacs` for the recipient. You'll receive a `recipient_id`, which you can re-use for future payments.

/payment\_initiation/recipient/create request

```
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

##### Creating a consent

Create a consent by calling [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate). You must provide the `recipient_id`, as well as a payment type and a set of constraints that align to your billing use case. This includes the maximum amount for a single payment, the frequency of payments, and the maximum cumulative amount for all payments over the consent period.

[`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate) will return a `consent_id`, which you should store associated with this end user, as it will be needed every time you execute a payment. If you forget to store the `consent_id`, there is no way to retrieve it; you will need to make a new call to [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate) and send your user back through the Link flow to grant consent.

Once the `consent_id` is created, it will have an initial status of `UNAUTHORISED`.

/payment\_initiation/consent/create request

```
const request: PaymentInitiationConsentCreateRequest = {
  recipient_id: recipientID,
  reference: 'TestPaymentConsent',
  type: PaymentInitiationConsentType.Commercial,
  constraints: {
    valid_date_time: {
      to: '2026-12-31T23:59:59Z',
    },
    max_payment_amount: {
      currency: PaymentAmountCurrency.Gbp,
      value: 15,
    },
    periodic_amounts: [
      {
        amount: {
          currency: PaymentAmountCurrency.Gbp,
          value: 40,
        },
        alignment: PaymentConsentPeriodicAlignment.Calendar,
        interval: PaymentConsentPeriodicInterval.Month,
      },
    ],
  },
};

try {
  const response = await plaidClient.paymentInitiationConsentCreate(request);
  const consentID = response.data.consent_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

#### Create a Link token

From your backend, call [`/link/token/create`](/docs/api/link/#linktokencreate), passing in the `consent_id`, to create a `link_token`.

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

#### Launch the Payment flow in Link

Plaid Link is a drop-in module that provides a secure, elegant authentication flow for the many financial institutions that Plaid supports. Link makes it secure and easy for users to connect their bank accounts to Plaid.

Because Link has access to all the details of the payment consent at the time of initialisation, it will display a screen with the consent details already populated. All your end user has to do is log in to their financial institution through a Link-initiated OAuth flow, select a funding account, and consent to the VRP details. When the end user has successfully done this, you will receive an `onSuccess` callback.

Note that these instructions cover Link on the web. For instructions on using Link within mobile apps, see the [Link documentation](/docs/link/). If you want to customize Link's look and feel, you can do so from the [Dashboard](https://dashboard.plaid.com/link).

##### Install Link dependency

index.html

```
<head>
  <title>Link for VRP</title>
  <script src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
</head>
```

##### Configure the client-side Link handler

Plaid communicates to you certain events that relate to how the user is interacting with Link. What you do with each of these event triggers depends on your particular use case, but a basic scaffolding might look like this:

app.js

```
const linkHandler = Plaid.create({
  // Use the link_token created in the previous step to initialize Link
  token: (await $.post('/create_link_token')).link_token,
  onSuccess: (public_token, metadata) => {
    // Show a success page to your user confirming that the
    // payment consent and bank account details were received.
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

Unlike most other Plaid products, it is not necessary to exchange the `public_token` you receive from the `onSuccess` callback for an `access_token` when using Variable Recurring Payments.

#### Making payments using the consent

Once a user has completed the consent flow in Link, the consent status will update to `AUTHORISED` and Plaid will send a [`CONSENT_STATUS_UPDATE`](/docs/api/products/payment-initiation/#consent_status_update) webhook to the webhook listener endpoint that you specified during the [`/link/token/create`](/docs/api/link/#linktokencreate) call.

At this point you can make payments within the consent parameters, with no user input required.

To make a payment, call [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute). You will need to provide the `consent_id`, as well as the `amount` of the payment, a `reference` string of your choice to identify the payment, and an `idempotency_key`.

The `idempotency_key` should be a string that is unique per payment and is used to ensure that you do not accidentally make the same payment twice when re-trying a payment attempt (e.g., when retrying after receiving a 500 error that does not guarantee whether or not the payment was successful); if you have already made a payment with the same `idempotency_key`, the payment attempt will fail.

/payment\_initiation/consent/payment/execute request

```
const request: PaymentInitiationConsentPaymentExecuteRequest = {
  consent_id: consentID,
  amount: {
    currency: PaymentAmountCurrency.Gbp,
    value: 7.99,
  },
  reference: 'Payment1',
  idempotency_key: idempotencyKey,
};
try {
  const response = await plaidClient.paymentInitiationConsentPaymentExecute(
    request,
  );
  const paymentID = response.data.payment_id;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

This will complete the process of initiating the payment. An individual instance of a VRP series is treated the same as a single payment created via One-Time Payment Initiation and has the same possible statuses and transitions.

To learn about tracking payments, see [Payment Status](/docs/payment-initiation/payment-status/).

#### Next steps

Review the following sections for information about additional steps that may be required for your integration.

[Managing consent](/docs/payment-initiation/variable-recurring-payments/managing-consent/): For building in-app functionality allowing users to terminate recurring payment consents, and to learn more about the consent lifecycle.

[Payment status](/docs/payment-initiation/payment-status/): To learn more about the payment status lifecycle and how to track the status of an initiated payment.

[Handling failed payments](/docs/payment-initiation/variable-recurring-payments/handling-failed-payments/): For details on error handling, including retryable versus non-retryable errors.

[Refunds](/docs/payment-initiation/variable-recurring-payments/refunds/): For details on refund initiation.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
