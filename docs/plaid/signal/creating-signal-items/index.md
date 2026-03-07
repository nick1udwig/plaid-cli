---
title: "Signal - Creating Items | Plaid Docs"
source_url: "https://plaid.com/docs/signal/creating-signal-items/"
scraped_at: "2026-03-07T22:05:18+00:00"
---

# Initialize Plaid Items

#### Create Plaid Items to be used with Signal Transaction Scores

#### Creating Plaid Items

Before getting a Signal Transaction Score for a proposed transaction, your end users need to link a bank account to your app using [Link](/docs/link/), Plaid's client-side widget. Link will connect the user's bank account and obtain the consent required to perform this evaluation.

See the [Link documentation](/docs/link/) for more details on setting up a Plaid Link session. At a high level, the steps are:

1. Call [`/link/token/create`](/docs/api/link/#linktokencreate).

   - Put `signal` and `auth` in the `products` array, along with any other Plaid products (except Balance) you will be requiring, e.g. `products: [signal, auth]`.
   - Put any other Plaid products you plan to use in the `optional_products` or `required_if_supported_products` arrays, e.g. `required_if_supported_products: [identity]`.
2. Initialize Link using the `link_token` created in the previous step. For more details for your specific platform, see the [Link documentation](/docs/link/). The user will now go through the Link flow.
3. Call [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to exchange the `public_token` for an `access_token`.
4. Obtain the `account_id` of the account used for the transaction you wish to perform the evaluation on; this can be obtained from the `metadata.accounts` field in the `onSuccess` callback, or by calling [`/accounts/get`](/docs/api/accounts/#accountsget) or [`/link/token/get`](/docs/api/link/#linktokenget).

Once you have your Plaid Item, continue to [evaluate](/docs/signal/signal-rules/) the risk of the transaction.

#### Adding Signal to existing Items

You may have Items that were not initialized with `signal`; for example, if you are adding Signal support to an existing Plaid integration. In this case, for best results, call [`/signal/prepare`](/docs/api/products/signal/#signalprepare) on the Item before your first call to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate). If you skip calling [`/signal/prepare`](/docs/api/products/signal/#signalprepare), the Item's first call to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will take longer and be less accurate, because Plaid will not have had the opportunity to pre-load certain data about the Item. Subsequent calls to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) on the Item will have higher accuracy.

If you intend to add Signal to an existing Item and have enabled [Data Transparency Messaging](/docs/link/data-transparency-messaging-migration-guide/) on the Item, you may need to send the Item through [update mode](/docs/link/update-mode/#data-transparency-messaging). If you have a large number of existing Items that require update mode for this reason, contact your Plaid Account Manager for more details and assistance.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
