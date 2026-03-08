---
title: "Payments (Europe) - User onboarding and account funding guide | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/payment-initiation-one-time/user-onboarding-and-account-funding/"
scraped_at: "2026-03-07T22:05:11+00:00"
---

# User onboarding and account funding

#### Use Plaid to simplify your onboarding and funding flow in Europe

In this guide, we first walk through setting up an account linking flow that uses Plaid to connect your user's bank account. We'll then demonstrate how to verify the identity of your users which may help comply with KYC and AML regulations.

Finally, we’ll show how how to initiate payments while requiring that bank accounts involved in money movements are verified.

If you're looking for a guide that is specific to Plaid's One-Time Payment Initiation product and does not incorporate other features such as anti-fraud measures or Identity Verification, see [Add One-time Payment Initiation to your App](/docs/payment-initiation/payment-initiation-one-time/).

This guide is specific to integrations serving end users in Europe (including non-EU countries such as the UK and Norway). For US payment solutions, see [Transfer](/docs/transfer/).

#### Before you get started

The [Auth](/docs/auth/), [Identity](/docs/identity/), and [Payments (Europe)](/docs/payment-initiation/) products covered in this guide can immediately be tried out in the Sandbox environment, which uses test data and does not interact with financial institutions.

To gain access to these products in the Production environments, [request product access](https://dashboard.plaid.com/settings/team/products) or [contact Sales](https://plaid.com/contact/). If you already are using Plaid products, you can also request access by contacting your account manager or submitting a [Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

### User onboarding

The **Plaid flow** begins when the user registers with your app. The flow will connect your app to your user's financial institution.

![Step  diagram](/assets/img/docs/payment_initiation/PIS-overview-row-0.png)

**1**Your backend creates a `link_token` by using the[`/link/token/create`](/docs/api/link/#create-link-token) endpoint and passes the temporary token to your Client app.

![Step 1 diagram](/assets/img/docs/link-tokens/link-token-row-2.png)

**2**Your Client app uses the `link_token` to initiate a Link flow for your user. The [`onSuccess` callback](/docs/link/web/#onsuccess) signals that the user has granted your app access to their financial institution.

![Step 2 diagram](/assets/img/docs/payment_initiation/PIS-overview-row-2.png)

**3**The `onSuccess` payload contains a [`public_token`](/docs/link/web/#link-web-onsuccess-public-token) which, after sending it to your backend, needs to be exchanged for a long-lived access token using the [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) endpoint.

![Step 3 diagram](/assets/img/docs/link-tokens/link-token-row-4.png)

**4**Using this access token your backend can now retrieve the user’s bank account numbers via [`/auth/get`](/docs/api/products/auth/#authget) or their identity information (such as their full name) via [`/identity/get`](/docs/api/products/identity/#identityget).

![Step 4 diagram](/assets/img/docs/payment_initiation/PIS-overview-row-no-downward-arrow.png)

#### Adding the Auth and Identity products to your app

The steps listed above are broken down into more detail in the guides linked below. The first guide covers adding the [Auth](/docs/api/products/auth/) product (e.g. for obtaining the user’s bank account numbers). The second covers the [Identity](/docs/api/products/identity/) product (e.g. for obtaining the user’s legal name). Both guides describe almost identical steps, so you can follow either one.

The next section will highlight the single difference between the two and will show how you can use both products at the same time.

[#### Auth

Add Auth to your app

Detailed steps](/docs/auth/add-to-app/)

#### Auth

Add Auth to your app

[Detailed steps](/docs/auth/add-to-app/)[#### Identity

Add Identity to your app

Detailed steps](/docs/identity/add-to-app/)

#### Identity

Add Identity to your app

[Detailed steps](/docs/identity/add-to-app/)

#### Creating a Link token

Link token creation is the only step requires changes when requesting both Auth and Identity-related information. To do so, pass both products as part of the [`/link/token/create`](/docs/api/link/#linktokencreate) request.

Create link token

```
curl -X POST https://sandbox.plaid.com/link/token/create -H 'Content-Type: application/json' -d '{ "client_id": "${PLAID_CLIENT_ID}", "secret": "${PLAID_SECRET}", "client_name": "Plaid Test App", "user": { "client_user_id": "${UNIQUE_USER_ID}" }, "products": ["auth", "identity"], // using both products "country_codes": ["GB", "NL", "DE"], "language": "en", "webhook": "https://webhook.sample.com", }'
```

#### End user experience

During the Link flow, your user grants permission to all the data products that you request. The images below demonstrate examples of the panes Link will present when requesting access to one or multiple products.

![Plaid consent screen for sharing financial data with WonderWallet, featuring an 'Allow' button and 'Account details' section.](/assets/img/docs/payment_initiation/ais-consent-pane-auth.png)

![Plaid data sharing consent screen, listing Contact and Account details. Buttons: Allow, Cancel. Terms link present.](/assets/img/docs/payment_initiation/ais-consent-pane-auth-identity.png)

### Optional steps and best practices

The sections below in this guide describe optional steps you can use to improve link conversion or achieve KYC/AML compliance.

#### Preselecting a financial institution

By default, Plaid will ask the user to manually select their financial institution. However, there might be cases where you already know which institution your user wants to use. For example, when your user has previously completed a sign up flow, you can increase conversion by skipping this part of the payment flow. This requires two steps:

1. Call [`/item/get`](/docs/api/items/#itemget) with the `access_token` as an input parameter. This endpoint will return the [`institution_id`](/docs/api/items/#item-get-response-item-institution-id) as part of the response.
2. Provide the `institution_id` as part of the payment [`/link/token/create`](/docs/api/link/#linktokencreate) request.

Create link token with institution preselected

```
curl -X POST https://sandbox.plaid.com/link/token/create -H 'Content-Type: application/json' -d '{ "client_id": "${PLAID_CLIENT_ID}", "secret": "${PLAID_SECRET}", "client_name": "Plaid Test App", "user": { "client_user_id": "${UNIQUE_USER_ID}" }, "products": ["payment_initiation"], "country_codes": ["GB", "NL", "DE"], "language": "en", "webhook": "https://webhook.sample.com", "payment_initiation": { "payment_id": "${PAYMENT_ID}" }, "institution_id": "${INSTITUTION_ID}" // preselect institution_id }'
```

#### Compliant account funding

To help comply with KYC & AML regulations you may choose to restrict payments to your app from bank accounts whose owner was previously verified.

One-Time Payment Initiation can support this flow by requiring that a payment is to be made from a specific account. The process has three steps:

1. Call [`/auth/get`](/docs/api/products/auth/#authget) with the `access_token` as an input parameter. This endpoint will return the user’s account numbers as part of the response (Use [`bacs`](/docs/api/products/auth/#auth-get-response-numbers-bacs) for the UK and [`international`](/docs/api/products/auth/#auth-get-response-numbers-international) for the rest of Europe).
2. Provide one of the account numbers under [`options`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options) as part of the [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) request. See the snippet below as an example.
3. Provide the `institution_id` as part of the [`/link/token/create`](/docs/api/link/#linktokencreate) request, as described in the previous section.

Not all financial institutions support this feature. In the UK there is full support. In the rest of Europe, support may be limited. Plaid will let the user choose any account when their financial institution does not support this feature. It is recommended to also verify the origin of payments as part of your payment reconciliation process.

Create payment request

```
curl -X POST https://sandbox.plaid.com/payment_initiation/payment/create -H 'Content-Type: application/json' -d '{ "client_id": "${PLAID_CLIENT_ID}", "secret": "${PLAID_SECRET}", "recipient_id": "${RECIPIENT_ID}", "reference": "Sample reference", "amount": { "currency": "GBP", "amount": 1.99 }, "options": { // additional payee account restriction "bacs": { "account": "26207729", "sort_code": "560029" } } }'
```

#### Compliant withdrawals

To comply with KYC and AML regulations, you may want to restrict outbound money movements to accounts whose owner is verified. You can do this in the following ways:

##### Using Virtual Accounts

Make sure your virtual account is set up before following these steps. For more information on setting up an account, see [Managing virtual accounts](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/).

1. Follow the [Payment Confirmation flow](/docs/payment-initiation/virtual-accounts/payment-confirmation/) to confirm that funds have settled in your virtual account.
2. Fetch the payment object by calling [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) with the `payment_id` of the payment you initiated. The payment object will contain a [`transaction_id`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-transaction-id), which corresponds to the payment's underlying virtual account transaction.
3. Fetch the virtual account transaction by calling [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) with the `transaction_id`.
4. The virtual account transaction will contain account details in the `counterparty` field.
5. [Execute a Payout](/docs/payment-initiation/virtual-accounts/payouts/) using the `counterparty` details.

##### Using Auth data

1. Call [`/auth/get`](/docs/api/products/auth/#authget) with the `access_token` as an input parameter. This endpoint will return the user’s verified account numbers as part of the response (Use [`bacs`](/docs/api/products/auth/#auth-get-response-numbers-bacs) for the UK and [`international`](/docs/api/products/auth/#auth-get-response-numbers-international) for the rest of Europe).
2. Select accounts that can receive funds by filtering for `account.subtype = "checking"`.
3. Initiate the withdrawal to any of the filtered accounts. You may want to let the user choose a preferred account in case there are multiple.

##### Using the `request_refund_details` flag

This approach is not recommended since it is an optional flag with limited incomplete coverage dependent on the financial institution a payment was initiated to.

1. Go through the [One-Time Payment Initiation flow](/docs/payment-initiation/payment-initiation-one-time/) and initiate a payment with the [`request_refund_details`](/docs/api/products/payment-initiation/#payment_initiation-payment-create-request-options-request-refund-details) flag enabled.
2. Fetch the payment object by calling [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) and parse the [`refund_details`](/docs/api/products/payment-initiation/#payment_initiation-payment-get-response-refund-details) field.
3. Initiate the withdrawal to account specified in the `refund_details` field.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
