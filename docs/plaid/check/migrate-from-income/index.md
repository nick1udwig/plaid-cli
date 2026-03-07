---
title: "Consumer Report (by Plaid Check) - Migrate from Income | Plaid Docs"
source_url: "https://plaid.com/docs/check/migrate-from-income/"
scraped_at: "2026-03-07T22:04:48+00:00"
---

# Migrate from Bank Income to Plaid Consumer Report

#### Switch to CRA Income Insights for model-driven income attributes and FCRA compliance

This guide will walk you through how to migrate from accessing bank data with Plaid's [Bank Income](https://plaid.com/docs/income/) product to generating Income Insights Reports with Plaid Check's Consumer Reports.

The CRA Income Insights Report offers model-driven attributes such as historical gross and net income, facilitating streamlined debt-to-income calculations and income verification. It enables assessment of financial stability through insights like forecasted income, employer name, income frequency, and predicted next payment date.

#### Advantages of Plaid Consumer Report

Bank Income gives you net income data from both irregular or gig income and W-2 income. The CRA Income Insights Report enhances this by:

- Providing FCRA-compliant insights
- Bundling income calculations (historical average income, forecasted income, predicted next payment date) in a single report

Upgrading to Plaid Consumer Report is recommended for all Bank Income customers whose end users are based in the US. (CRA Income Insights is not available for end users outside the US.)

Income Insights does not replace Document Income or Payroll Income. Plaid Check products can't be used in the same Link flow as Plaid Inc. credit products such as Plaid Document Income, Payroll Income, Assets, or Statements. If you want to continue using any of those products, create separate [`/link/token/create`](/docs/api/link/#linktokencreate) requests: one for Plaid Check products, and one for the Plaid Inc. credit products.

Non-credit Plaid products, such as Balance and Auth, can continue to be used in the same Link flow as Plaid Check products, and do not require any special changes to your integration.

#### Prerequisites

1. To use Consumer Report, it is strongly recommended to update your Plaid client library to the latest version. The minimum required versions for new Consumer Report integrations are:

   - Python: 38.0.0
   - Go: 41.0.0
   - Java: 39.0.0
   - Node: 41.0.0
   - Ruby: 45.0.0
2. [Confirm](https://dashboard.plaid.com/settings/team/products) that you have access to all required Plaid Check products in the Production environment. If you don't have access to Plaid Check, [request access via the Dashboard](https://dashboard.plaid.com/overview/request-products). In order to use Plaid Check, your end users must be in the US and you must be on a custom pricing plan.

#### Changes to Plaid Link initialization

When using Plaid Check products, you must create a user prior to sending the user through Link and initialize Link with the resulting `user_id`. This allows Plaid to associate multiple Items with a single user.

1. Call [`/user/create`](/docs/api/users/#usercreate) prior to creating a Link token:

   - Include identity info in the `identity` object. At minimum, the following fields must be provided: `name`, `date_of_birth`, `emails`, `phone_numbers`, and `addresses` (with at least one email address, phone number, and address designated as `primary`). If you intend to share the report with a GSE (Government-Sponsored Entity) such as Fannie or Freddie, the full SSN is also required via the `id_numbers` field. For all use cases, providing at least a partial SSN is highly recommended, since it improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests.
   - Store the `user_id` in your database.
2. Update your [`/link/token/create`](/docs/api/link/#linktokencreate) request:

   - Include the `user_id` string from [`/user/create`](/docs/api/users/#usercreate) .
   - In the `products` array, replace `income_verification` with `cra_base_report` and `cra_income_insights`.
   - Remove the `income_verification` object.
   - Add a `cra_options` object and specify `days_requested` (minimum 180).
   - If you haven't already, set `webhook` to your webhook endpoint listener URL. Unlike Bank Income, Plaid Check requires the use of webhooks.
   - Provide a `consumer_report_permissible_purpose`
   - (Optional) To allow linking multiple Items in one session, set `enable_multi_item_link` to `true`.

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

To enable an existing Income-enabled Item for Consumer Report, call [`/user/create`](/docs/api/users/#usercreate) and [`/link/token/create`](/docs/api/link/#linktokencreate) as described above, but include the Item's `access_token` when calling [`/link/token/create`](/docs/api/link/#linktokencreate). When Link is launched, the end user will go through the Consumer Report consent flow, and on successful completion of the flow, the Item will be enabled for Consumer Report.

#### Update product API calls and webhook listeners

Unlike Bank Income, Plaid Check uses webhooks with async report generation.

1. Add webhook listener endpoints for [`USER_CHECK_REPORT_READY`](/docs/api/products/check/#user_check_report_ready) and [`USER_CHECK_REPORT_FAILED`](/docs/api/products/check/#user_check_report_failed).

   - Upon receiving the [`USER_CHECK_REPORT_READY`](/docs/api/products/check/#user_check_report_ready) webhook, call product endpoints such as [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget).
   - Upon receiving the [`USER_CHECK_REPORT_FAILED`](/docs/api/products/check/#user_check_report_failed) webhook, call [`/user/items/get`](/docs/api/users/#useritemsget) to determine why Items are in a bad state. If appropriate, send the user through [update mode](/docs/link/update-mode/) to repair the Item.
2. Replace any Income endpoints and webhook listeners with the equivalent Plaid Check endpoint or webhook listener, using the [API response comparison sheet](https://docs.google.com/spreadsheets/d/1sjzpbPNjz0bF0Ndaly0tShMFgAZ2M9zdvBcQMlDKJbE/edit?gid=727344192#gid=727344192) to map response fields between the old endpoints and new endpoints.

| Income endpoint | Equivalent Plaid Check endpoint |
| --- | --- |
| [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) | None (remove) |
| [`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) | [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget) and/or [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget) |
| [`/credit/bank_income/pdf/get`](/docs/api/products/income/#creditbank_incomepdfget) | [`/cra/check_report/pdf/get`](/docs/api/products/check/#cracheck_reportpdfget) with `add_ons: ["cra_income_insights"]` |
| [`/credit/bank_income/refresh`](/docs/api/products/income/#creditbank_incomerefresh) | [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) |

| Income webhook | Equivalent Plaid Check webhook |
| --- | --- |
| `BANK_INCOME_REFRESH_COMPLETE` | `USER_CHECK_REPORT_READY` |

#### API response comparison

This [Google Sheet](https://docs.google.com/spreadsheets/d/1sjzpbPNjz0bF0Ndaly0tShMFgAZ2M9zdvBcQMlDKJbE/edit?gid=727344192#gid=727344192) highlights the correspondences between Bank Income and Plaid Check Income Insights schemas.

If you are using other Plaid Inc. products, note that the `account_id` returned in API responses from endpoints prefixed with `/cra/` will not match the `account_id` returned in responses from non-CRA Plaid endpoints.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
