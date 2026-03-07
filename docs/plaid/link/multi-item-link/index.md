---
title: "Link - Multi-Item Link | Plaid Docs"
source_url: "https://plaid.com/docs/link/multi-item-link/"
scraped_at: "2026-03-07T22:05:07+00:00"
---

# Multi-Item Link

#### Allow end users to connect accounts from multiple institutions in a single Link session

#### Overview

Multi-Item Link allows end users to add multiple Items in the same Link session. This flow is designed for contexts where you expect the user to link multiple accounts; for example, to provide a complete financial picture for a lending or personal finance use case.

Multi-Item Link can result in more Items being connected per user, due to reduced friction, and can simplify your app's logic and user experience.

Multi-Item Link is compatible with the following Plaid products: Auth, Transfer, Transactions, Liabilities, Investments, Bank Income, Assets, and Plaid Check Consumer Reports (including Income and Partner Insights). It does not require special enablement and is automatically available to all customers.

Multi-Item Link is not compatible with [Embedded Institution Search](https://plaid.com/docs/link/embedded-institution-search). Multi-Item Link is also not currently compatible with [non-credential based Auth flows](/docs/auth/coverage/same-day/) such as Database Auth, Same Day Micro-Deposits, or Instant Micro-deposits.

![Success screen showing 2 connected institutions, Gingham Bank and Herringbone Treasury, with option to add another and Continue button.](/assets/img/docs/multi-item-link-add.png)

#### Integration process

##### Creating a Link token

First, create a `user_id` and/or `user_token` using the [`/user/create`](/docs/api/users/#usercreate) endpoint.

Sample /user/create request for Multi-Item Link

```
// client_user_id is your own internal identifier for the end user
{
    "client_id": "${PLAID_CLIENT_ID}",
    "secret": "${PLAID_SECRET}",
    "client_user_id" : "c0e2c4ee-b763-4af5-cfe9-46a46bce883d" 
}
```

Next, call [`/link/token/create`](/docs/api/link/#linktokencreate) with the user identifier you just created, and the `enable_multi_item_link` field set to `true`. If you received a `user_token` from [`/user/create`](/docs/api/users/#usercreate), you should use that as your identifier. If you are using the newer integration flow that returns only a `user_id` and does not return a `user_token`, you should use the `user_id`.

Sample /link/token/create request for Multi-Item Link

```
// Sample link/token/create request
{
    "client_id": "${PLAID_CLIENT_ID}",
    "secret": "${PLAID_SECRET}",
    "client_name" : "Plaid Test App",
    "user_id":  "usr_9nSp2KuZ2x4JDw",
    "enable_multi_item_link": true,
    "user": {
        "client_user_id": "c0e2c4ee-b763-4af5-cfe9-46a46bce883d",
    },
    "products": ["transactions"],
    "country_codes": ["US"],
    "language": "en",
}
```

Once you have created a Link token, you will launch [Link](/docs/link/) as normal, according to the standard instructions for your platform.

##### Obtaining a public token

In most other Plaid integration methods, upon completion of the Link flow, you will receive a public token via the `onSuccess` callback. In a Multi-Item Link session, the `onSuccess` callback will be empty, and you will instead receive information about the Link flow (including the `public_token` array) via the [`SESSION_FINISHED`](/docs/api/link/#session_finished) webhook.

When the entire Multi-Item Link flow is complete and the user has exited Link, Plaid will fire a [`SESSION_FINISHED`](/docs/api/link/#session_finished) webhook that contains information about what caused the session to end, as well as the public tokens if the session completed successfully.

Sample SESSION\_FINISHED webhook

```
{
  "webhook_type": "LINK",
  "webhook_code": "SESSION_FINISHED",
  "status": "SUCCESS",
  "link_session_id": "356dbb28-7f98-44d1-8e6d-0cec580f3171",
  "link_token": "link-sandbox-af1a0311-da53-4636-b754-dd15cc058176",
  "public_tokens": [
    "public-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d"
  ],
  "environment": "sandbox"
}
```

If you want to start getting results as soon as each Item is added, you can listen for the [`ITEM_ADD_RESULT`](/docs/api/link/#item_add_result) webhook, which will fire after each completed Item add within the Link session.

Sample ITEM\_ADD\_RESULT webhook

```
{
  "webhook_type": "LINK",
  "webhook_code": "ITEM_ADD_RESULT",
  "link_session_id": "356dbb28-7f98-44d1-8e6d-0cec580f3171",
  "link_token": "link-sandbox-af1a0311-da53-4636-b754-dd15cc058176",
  "public_token": "public-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d",
  "environment": "sandbox"
}
```

Detailed session data, including the public tokens and account metadata, will also be available from [`/link/token/get`](/docs/api/link/#linktokenget) for six hours after the session has completed. In general, it is recommended to use the webhook to obtain the public token, since it will allow you to get the public token more promptly, but you may want to use [`/link/token/get`](/docs/api/link/#linktokenget) if your integration does not use webhooks, or as a backup mechanism if your system missed webhooks due to an outage.

When using Multi-Item Link, make sure to obtain the `public_token` from either the `SESSION_FINISHED` webhook, the `ITEM_ADD_RESULT` webhook, or the `results.item_add_results` object returned by [`/link/token/get`](/docs/api/link/#linktokenget). The `on_success` object returned by [`/link/token/get`](/docs/api/link/#linktokenget) is deprecated and should not be used to obtain the public token in Multi-Item Link flows, as it will contain only one public token, rather than all of the public tokens from the session.

##### Frontend changes

When using Multi-Item Link, the frontend `onSuccess` callback will be empty, but other callbacks, such as `onExit` and `onEvent`, will remain populated as usual. Frontend callbacks can still be used to signal the end of the Link session. If you need the metadata that is normally found in `onSuccess`, such as the institution id or account mask, you can obtain it from the `results.item_add_results` object returned by [`/link/token/get`](/docs/api/link/#linktokenget).

#### Multi-Item Link with multiple products

If you are using Multi-Item Link, the same Link token settings will be used for every Item in the Link flow. This means, for example, that Auth with Same Day Micro-deposits cannot be used in the same Multi-Item Link flow as Income, because Auth must be initialized by itself to be used with Same Day Micro-deposits.

##### Using Auth or Transfer with Multi-Item Link

A common use case for including Auth or Transfer as part of a Multi-Item link process is when you want one of the linked accounts to be used for sending or receiving payments. For example, as a lender, you may want to disburse the user's loan to a linked account, or as a financial management app, you may want to use one of the linked accounts to collect a monthly subscription fee.

In this situation, your app will want to connect to multiple institutions, but only need one account for payment purposes. For this use case, the recommended flow is:

1. Put `auth` or `transfer` in the `optional_products` array when calling [`/link/token/create`](/docs/api/link/#linktokencreate); this will request Auth or Transfer permissions when possible, but will not block linking accounts that don't support Auth and won't cause you to be billed for Auth unless you use an Auth endpoint on the Item.
2. Locate eligible accounts by looking at the [`/link/token/get`](/docs/api/link/#linktokenget) response object and traversing `link_sessions.results.item_add_results[].account[].subtype` to check each account on each Item, to see if it is of subtype `checking` or `savings`.
3. For each Item on an account where the subtype is `checking` or `savings`, call [`/item/get`](/docs/api/items/#itemget) and check the `item.products[]` array for `auth`.
4. If the account subtype is `checking` or `savings`, and the `item.products[]` value contains `"auth"` as a value, then the account is eligible for use with Plaid Auth or Transfer. Present a UI in which the user can select an eligible account to use for sending or receiving funds. To avoid confusion, it is recommended that you present this UI even if there is only one eligible account.
5. (Optional) You can add a button to this UI to link a different account to use for funds transfer, and have that button launch Link with just `auth` or `transfer`, which could allow the end user to link a payment account via a non-credential-based flow, such as Same Day Micro-Deposits or Database Insights.
6. Once the user has selected an account to use for payments, call the appropriate endpoint, such as [`/auth/get`](/docs/api/products/auth/#authget), [`/processor/token/create`](/docs/api/processors/#processortokencreate), or [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate), on the corresponding Item.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
