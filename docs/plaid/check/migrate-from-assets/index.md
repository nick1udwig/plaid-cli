---
title: "Consumer Report (by Plaid Check) - Migrate from Assets | Plaid Docs"
source_url: "https://plaid.com/docs/check/migrate-from-assets/"
scraped_at: "2026-03-07T22:04:47+00:00"
---

# Migrate from Asset Reports to Base Reports

#### Migrate to Base Reports to unlock access to the Consumer Report

This guide will walk you through how to migrate from generating Asset Reports with Plaid's [Assets](https://plaid.com/products/assets/) product to generating Base Reports with Plaid Check's [Consumer Report](https://plaid.com/check/income-and-underwriting/) product.

#### Prerequisites

1. To use Consumer Report, it is strongly recommended to update your Plaid client library to the latest version. The minimum required versions for new Consumer Report integrations are:
   - Python: 38.0.0
   - Go: 41.0.0
   - Java: 39.0.0
   - Node: 41.0.0
   - Ruby: 45.0.0
2. [Confirm](https://dashboard.plaid.com/settings/team/products) that you have access to all required Plaid Check products in the Production environment. If you don't have access to Plaid Check, [request access via the Dashboard](https://dashboard.plaid.com/overview/request-products). In order to migrate from Assets to Plaid Consumer Report, your end users must be in the US and you must be on a custom plan.

#### Changes to Plaid Link initialization

When using Plaid Check products, you must create a user prior to sending the user through Link, and you must initialize Link with the resulting `user_id`. This allows Plaid to associate multiple Items with a single user.

1. Call [`/user/create`](/docs/api/users/#usercreate) prior to creating a Link token
   - Include information about your user in the `identity` object. At minimum, the following fields must be provided and non-empty: `name`, `emails`, `phone_numbers`, and `addresses`. If you intend to share the report with a GSE (Government-Sponsored Entity) such as Fannie or Freddie, the full SSN is also required via the `id_numbers` field. For all use cases, providing at least a partial SSN is highly recommended, since it improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests.
   - Store the `user_id` in your database
2. Update your call to [`/link/token/create`](/docs/api/link/#linktokencreate)
   - Include the `user_id` string from [`/user/create`](/docs/api/users/#usercreate)
   - Replace the `assets` product string with `cra_base_report` in the `products` array
   - Add a `cra_options` object and specify your desired `days_requested`
   - Include a `consumer_report_permissible_purpose`
3. (Optional) Unlike Assets, by default, Plaid Check includes only one linked Item per Report. To include multiple Items in a single Consumer Report, set `enable_multi_item_link` to `true` in the [`/link/token/create`](/docs/api/link/#linktokencreate) request.

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

For more details, see the Plaid Check [Implementation guide](https://plaid.com/docs/check/add-to-app/).

##### Adding Consumer Report to existing Items

To enable an existing Assets-enabled Item for Consumer Report, call [`/user/create`](/docs/api/users/#usercreate) and [`/link/token/create`](/docs/api/link/#linktokencreate) as described above, but include the Item's `access_token` when calling [`/link/token/create`](/docs/api/link/#linktokencreate). When Link is launched, the end user will go through the Consumer Report consent flow, and on successful completion of the flow, the Item will be enabled for Consumer Report.

#### Changes to post-Link integration

Update your webhook listeners to listen for the new `USER_CHECK_REPORT_READY` and `USER_CHECK_REPORT_FAILED` webhooks.

- Upon receiving a `USER_CHECK_REPORT_READY` webhook, you should call [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), along with any other Consumer Report product endpoints that you would like to use.
- Upon receiving a `USER_CHECK_REPORT_FAILED` webhook, you should call [`/user/items/get`](/docs/api/users/#useritemsget) to determine why Items are in a bad state. If appropriate, send the user through [Update Mode](https://plaid.com/docs/link/update-mode/#using-update-mode) to repair them.

If you plan to continue using other Plaid Inc. products, such as [Auth](https://plaid.com/products/auth/) or [Balance](https://plaid.com/products/balance/), you should continue to retrieve the public token and exchange it for an access token. If not, you no longer need to obtain an access token.

#### Mapping API responses

Asset Reports and Base Reports have similar but not identical schemas. This [Google Sheet](https://docs.google.com/spreadsheets/d/1toj71Afemtu0rSjgh7AAGgBcZC7U0vYXKCDvKSSWTls/edit?usp=sharing) highlights their differences.

If you are using other Plaid Inc. products, note that the `account_id` returned in API responses from endpoints prefixed with `/cra/` will not match the `account_id` returned in responses from non-CRA Plaid endpoints.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
