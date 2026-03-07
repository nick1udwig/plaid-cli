---
title: "Transfer - Transfer UI | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/using-transfer-ui/"
scraped_at: "2026-03-07T22:05:26+00:00"
---

# Receiving Funds Using Transfer UI

#### Facilitate transfers with an intuitive user interface

Plaid Transfer UI is a drop-in user interface that makes it easy for end users to authorize one-time transfers. Transfer UI is compliant with Nacha WEB guidelines and automatically captures and manages [Proof of Authorization](https://plaid.com/docs/transfer/creating-transfers/#managing-proof-of-authorization) on your behalf.

If your use case involves one-time debit transfers and you do not have the resources to build your own authorization UX, Transfer UI may be a good fit for you.

With Transfer UI, users are able to authorize payments or disbursements. (While Proof of Authorization requirements are not mandatory for credits, Transfer UI supports both credit and debit transfers.) Before authorizing a transfer, users are able to review transfer details, such as amount, fund origination account, fund target account, and more.

Transfer UI does not currently support [recurring transfers](https://plaid.com/docs/transfer/recurring-transfers/), [Database Auth](https://plaid.com/docs/auth/coverage/database/), or [Transfer for Platforms](https://plaid.com/docs/transfer/platform-payments/). Transfer UI cannot be used to capture standing authorizations such as those required for variable recurring payments.

![Mobile app transfer UI flow showing loan payment process, payment method selection, transfer confirmation, and success message.](/assets/img/docs/transfer/transfer_ui.png)

Example Transfer UI flow.

Prefer to learn by watching? Get an overview of how Transfer UI works in just 3 minutes!

#### Integration

To integrate with Transfer UI, follow the following steps:

1. In the Plaid Dashboard, create a [Link customization](/docs/link/customization/) with "Account Select" set to the "Enabled for one account" selection.
2. Call [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) to obtain a `transfer_intent.id`.
3. Call [`/link/token/create`](/docs/api/link/#linktokencreate), specifying `["transfer"]` in the `products` parameter, the `transfer_intent.id` in the `transfer.intent_id` parameter, and the name of your Link customization in the `link_customization_name` parameter. If you already have an `access_token` for this user, you can provide it to [`/link/token/create`](/docs/api/link/#linktokencreate) to streamline the Link flow, otherwise, the user will authenticate their bank account within the same Transfer UI session.
4. Initialize a Link instance using the `link_token` created in the previous step. For more details for your specific platform, see the [Link documentation](/docs/link/).
5. The user will now go through the Link flow to perform their transfer. The `onSuccess` callback will indicate they have completed the Link flow.
6. (Optional) You will receive a `public_token` from the `onSuccess` callback. If you do not already have an access token for this user's account, call [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to exchange this public token for an `access_token` in order to streamline future transfers for this user.

#### Transfer UI user flows

Depending on whether you provided an access token in the [`/link/token/create`](/docs/api/link/#linktokencreate) step, the Transfer UI experience will differ.

Bank on file flow (`access_token` provided):

![Plaid transfer UI flow: 1. Sending credentials. 2. Confirm transfer from Bank to customer. 3. Transfer success.](/assets/img/docs/transfer/bank_on_file_user.png)

Example Transfer UI flow if the access\_token is provided.

Net new user flow (`access_token` not provided):

![6-step UI flow for Plaid transfer without access_token: connecting account, selecting bank, entering credentials, account selection, transfer confirmation, success.](/assets/img/docs/transfer/net_new_user.png)

Example Transfer UI flow if the access\_token is not provided.

#### Tracking transfer creation

The instructions in this section correspond to the Web and Webview libraries for Link. If you're using a mobile SDK, information about the transfer intent status can be found in the `metadataJson` field in the SDK's `onSuccess` callback. As an example, see [Android: metadataJson](https://plaid.com/docs/link/android/#link-android-onsuccess-metadata-metadataJson).

You can determine whether a transfer was successfully created by referring to the `transfer_status` field in the `metadata` object returned by `onSuccess`. A value of `complete` indicates that the transfer was successfully originated. A value of `incomplete` indicates that the transfer was not originated. Note that this field only indicates the status of a transfer creation. It does not indicate the status of a transfer (i.e., funds movement). For more information on transfer intents, see [Retrieving additional information about transfer intents](/docs/transfer/using-transfer-ui/#retrieving-additional-information-about-transfer-intents). For help troubleshooting incomplete transfers, see [Troubleshooting Transfer UI](/docs/transfer/using-transfer-ui/#troubleshooting-transfer-ui).

When using Transfer UI, the `onSuccess` callback is called at a different point in time in the Link flow. Typically, the `onSuccess` callback is called after an account is successfully linked using Link. When using Transfer UI, however, the `onSuccess` callback is called only after both of the following conditions are met: an account is successfully linked using Link and a transfer is confirmed via the UI. Note that the `onSuccess` callback only indicates that an account was successfully linked. It does not indicate a successful transfer (i.e., funds movement).

Example onSuccess metadata object, incomplete transfer creation

```
{
  institution: {
    name: 'Wells Fargo',
    institution_id: 'ins_4'
  },
  accounts: [
    {
      id: 'ygPnJweommTWNr9doD6ZfGR6GGVQy7fyREmWy',
      name: 'Plaid Checking',
      mask: '0000',
      type: 'depository',
      subtype: 'checking',
      verification_status: ''
    },
    {
      id: '9ebEyJAl33FRrZNLBG8ECxD9xxpwWnuRNZ1V4',
      name: 'Plaid Saving',
      mask: '1111',
      type: 'depository',
      subtype: 'savings'
    }
    ...
  ],
  transfer_status: 'incomplete',
  link_session_id: '79e772be-547d-4c9c-8b76-4ac4ed4c441a'
}
```

#### Retrieving additional information about transfer intents

To retrieve more information about a transfer intent, call the [`/transfer/intent/get`](/docs/api/products/transfer/account-linking/#transferintentget) endpoint and pass in a transfer intent `id` (returned by the [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) endpoint).

Retrieving more information about a transfer intent

```
const request: TransferIntentGetRequest = {
  transfer_intent_id: '460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9',
};

try {
  const response = await client.transferIntentGet(request);
} catch (error) {
  // handle error
}
```

Example /transfer/intent/get response

```
{
  "transfer_intent": {
    "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
    "ach_class": "ppd",
    "amount": "15.75",
    "authorization_decision": "APPROVED",
    "authorization_decision_rationale": null,
    "created": "2020-08-06T17:27:15Z",
    "description": "Desc",
    "failure_reason": null,
    "id": "460cbe92-2dcc-8eae-5ad6-b37d0ec90fd9",
    "metadata": {
      "key1": "value1",
      "key2": "value2"
    },
    "mode": "DISBURSEMENT",
    "origination_account_id": "9853defc-e703-463d-86b1-dc0607a45359",
    "status": "SUCCEEDED",
    "transfer_id": "8945fedc-e703-463d-86b1-dc0607b55460",
    "user": {
      "address": {
        "street": "100 Market Street",
        "city": "San Francisco",
        "region": "California",
        "postal_code": "94103",
        "country": "US"
      },
      "email_address": "lknope@email.com",
      "legal_name": "Leslie Knope",
      "phone_number": "123-456-7890"
    }
  },
  "request_id": "saKrIBuEB9qJZno"
}
```

The response is a `transfer_intent` object with more information about the transfer intent.

The `status` field in the response indicates whether the transfer intent was successfully captured. It can have one of the following values: `FAILED`, `SUCCEEDED`, or `PENDING`.

If the value of `status` is `FAILED`, the transfer intent was not captured. The `transfer_id` field will be `"null"` and the `failure_reason` object will contain information about why the transfer intent failed. Some possible reasons may include: the authorization was declined, the account is blocked, or the origination account was invalid.

If the value of `status` is `SUCCEEDED`, the transfer intent was successfully captured. The accompanying `transfer_id` field in the response will be set to the ID of the originated transfer. This ID can be used with other [Transfer API endpoints](/docs/api/products/transfer/).

A value of `PENDING` can mean one of the following:

- The transfer intent has not yet been processed by the authorization engine (`authorization_decision` is `"null"`). This is the initial state of all transfer intents.
- The transfer intent was processed and approved by the authorization engine (`authorization_decision` is `"approved"`), but has not yet been processed by the transfer creation processor.
- The transfer intent was processed by the authorization engine and was declined (`authorization_decision` is `"declined"`) due to insufficient funds (`authorization_decision_rationale.code` is `"NSF"`). If this is the case, the end user can retry the transfer intent up to three times. The transfer intent status will remain as `"PENDING"` throughout the retry intents. After three unsuccessful retries, the transfer intent status will be `"FAILED"`.

#### Transfer UI best practices

While transfer UI provides a Nacha-compliant UI for your transfers, to maximize conversion, it is recommended to still provide context in your own app to your end users about the transfer. We recommend that your app implement a UI where the customer selects or confirms the transfer amount and, if applicable, the funds source or destination. To maximize user confidence, your app should also provide a post-Link success pane confirming the transfer details, including amount and effective date.

#### Sample implementation

For a real-life example of an app that incorporates both Transfer and Transfer UI, see the Node-based [Plaid Pattern Transfer](https://github.com/plaid/pattern-transfers) sample app. Pattern Transfer is a sample subscriptions payment app that enables ACH bank transfers.

#### Troubleshooting Transfer UI

##### The onExit callback is called

If the `onExit` callback is called, the Link session failed to link an account. The transfer intent could not be initiated.

##### onSuccess is called, but the transfer intent is incomplete

The `onSuccess` callback is not an indication of a successful transfer creation or a successful transfer (i.e., funds movement). It only indicates the successful linking of an account. To fully diagnose incomplete transfer intents (i.e., when the `transfer_status` field in the `metadata` object returned by `onSuccess` is `incomplete`), call the [`/transfer/intent/get`](/docs/api/products/transfer/account-linking/#transferintentget) endpoint and pass in the corresponding transfer intent ID as an argument. Information about the transfer intent can be found in the `status`, `failure_reason`, `authorization_decision`, and `authorization_decision_rationale` fields in the response.

##### Network errors or institution connection issues

Network errors or institution connection issues can also result in incomplete transfer intents. In such cases, generate a new Link token using the existing `access_token` and invoke Transfer UI to allow the end user to attempt the transfer intent again.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
