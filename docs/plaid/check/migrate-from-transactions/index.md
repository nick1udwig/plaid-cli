---
title: "Consumer Report (by Plaid Check) - Migrate from Transactions | Plaid Docs"
source_url: "https://plaid.com/docs/check/migrate-from-transactions/"
scraped_at: "2026-03-07T22:04:48+00:00"
---

# Migrate from Transactions to Plaid Consumer Report

#### Switch to CRA Base Report for FCRA-compliant reporting with enhanced data coverage.

This guide will walk you through how to migrate from accessing bank data with Plaid's [Transactions](https://plaid.com/products/transactions/) product to generating Base Reports with Plaid's [Consumer Report](https://plaid.com/check/income-and-underwriting/) product.

The CRA Base Report includes account balance, account ownership, and transaction data, and is designed for use cases that require FCRA-compliant consumer reporting, such as tenant screening, income verification, and credit decisioning.

#### Why migrate?

While the Transactions product gives you historical transaction data, the CRA Base Report enhances this by:

- Supporting FCRA-compliant usage
- Bundling additional data (identity, balance, account metadata) in a single report
- Managing the explicit consumer consent and disclosures required for underwriting use cases in the US

If your use case involves evaluating the financial standing of US-based end users for underwriting, credit, or leasing, you should use the CRA Base Report instead of Transactions.

#### Prerequisites

1. To use Consumer Report, it is strongly recommended to update your Plaid client library to the latest version. The minimum required versions for new Consumer Report integrations are:

   - Python: 38.0.0
   - Go: 41.0.0
   - Java: 39.0.0
   - Node: 41.0.0
   - Ruby: 45.0.0
2. [Confirm](https://dashboard.plaid.com/settings/team/products) you have access to Plaid Check products in Production. If not, [request access via the Dashboard](https://dashboard.plaid.com/overview/request-products).

#### Changes to Plaid Link initialization

When using Plaid Check products, you must create a user before initializing Link. This allows Plaid to associate multiple Items with a single user.

1. Call [`/user/create`](/docs/api/users/#usercreate) before [`/link/token/create`](/docs/api/link/#linktokencreate) .

   - Include the `identity` object. At minimum, the following fields must be provided and non-empty: `name`, `date_of_birth`, `emails`, `phone_numbers`, and `addresses` (with at least one email address, phone number, and address designated as `primary`). If you intend to share the report with a GSE (Government-Sponsored Entity) such as Fannie or Freddie, the full SSN is also required via the `id_numbers` field. For all use cases, providing at least a partial SSN is highly recommended, since it improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests.
   - Store the `user_id` in your database.
2. Update your [`/link/token/create`](/docs/api/link/#linktokencreate) call:

   - Remove the `transactions` object (if present) from the call.
   - Include the `user_id` from [`/user/create`](/docs/api/users/#usercreate).
   - Replace `transactions` with `cra_base_report` (and any other CRA products) in the `products` array.
   - Add a `cra_options` object with `days_requested` (min 180).
   - Provide a `consumer_report_permissible_purpose`.
   - (Optional) To allow multi-institution linking, set `enable_multi_item_link` to `true`.

server.js

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

##### Adding Consumer Report to existing Items

To enable an existing Transactions-enabled Item for Consumer Report, call [`/user/create`](/docs/api/users/#usercreate) and [`/link/token/create`](/docs/api/link/#linktokencreate) as described above, but include the Item's `access_token` when calling [`/link/token/create`](/docs/api/link/#linktokencreate). When Link is launched, the end user will go through the Consumer Report consent flow, and on successful completion of the flow, the Item will be enabled for Consumer Report.

#### Changes to post-Link integration

#### Update product API calls and webhook listeners

1. Update your webhook listener endpoints, adding listeners for Plaid Check.

   - Upon receiving the [`USER_CHECK_REPORT_READY`](/docs/api/products/check/#user_check_report_ready) webhook, call product endpoints such as [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget) or (if using) [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget).
   - Upon receiving the [`USER_CHECK_REPORT_FAILED`](/docs/api/products/check/#user_check_report_failed) webhook, call [`/user/items/get`](/docs/api/users/#useritemsget) to determine why Items are in a bad state. If appropriate, send the user through [update mode](/docs/link/update-mode/) to repair the Item.
2. If you're still using other Plaid products like [Auth](https://plaid.com/products/auth/) or [Balance](https://plaid.com/products/balance/), continue retrieving and exchanging the public token. Otherwise, it's no longer needed.

#### Mapping API responses

This [Google Sheet](https://docs.google.com/spreadsheets/d/1toj71Afemtu0rSjgh7AAGgBcZC7U0vYXKCDvKSSWTls/edit?gid=820738585#gid=820738585) highlights the correspondences between Transactions and CRA Base Report schemas.

If you are using other Plaid Inc. products, note that the `account_id` returned in API responses from endpoints prefixed with `/cra/` will not match the `account_id` returned in responses from non-CRA Plaid endpoints.

#### Migrating existing Items

There is no way to directly migrate an existing Transactions-enabled Item to Plaid Check. Instead, you must delete the old Item using [`/item/remove`](/docs/api/items/#itemremove) and prompt the user to go through the Link flow again, using a Link token that is enabled for Plaid Check.

If you are ready to migrate all Items, use [`/item/remove`](/docs/api/items/#itemremove) to delete the Transactions-enabled Items, and prompt your users to return to your app to re-add their accounts using the new Plaid Check-enabled flow. You can also perform this process in stages, disabling only a percentage of Items at a time.

The most conservative gradual migration strategy is to not delete the Transactions-enabled Item until it is in an unhealthy state that requires user interaction to fix (e.g. `ITEM_LOGIN_REQUIRED`). At that time, instead of sending the Item through update mode, delete the Item using [`/item/remove`](/docs/api/items/#itemremove) and prompt your users to re-add their accounts using the new Plaid-Check enabled flow.

#### Removing Transactions logic for existing Items

Depending on your use case and business, you may want to leave existing Transactions-enabled Items in place and perform a gradual migration, as described above. Do not remove Transactions logic unless you are no longer using any Transactions-enabled Items.

1. Remove handling for Transactions webhooks, including `SYNC_UPDATES_AVAILABLE`, `INITIAL_UPDATE`, `HISTORICAL_UPDATE`, `DEFAULT_UPDATE`, `RECURRING_TRANSACTIONS_UPDATE`
2. Remove calls to `/transactions/` endpoints such as [`/transactions/get`](/docs/api/products/transactions/#transactionsget), [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), or [`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget), as well as any handling for transactions-specific logic, such as transactions updates and pagination.
3. Remove handling for the `TRANSACTIONS_SYNC_MUTATION_DURING_PAGINATION` error.
4. If you have existing Transactions-enabled Items you will no longer be using, call [`/item/remove`](/docs/api/items/#itemremove) to delete these Items so you will no longer be billed for Transactions.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
