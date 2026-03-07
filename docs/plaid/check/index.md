---
title: "Consumer Report (by Plaid Check) - Introduction | Plaid Docs"
source_url: "https://plaid.com/docs/check/"
scraped_at: "2026-03-07T22:04:47+00:00"
---

# Introduction to Consumer Report

#### Use Consumer Report for smarter credit and lending decisions powered by Plaid Check

Get started with Consumer Report

[API Reference](/docs/api/products/check/)[Request Access](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access)[Demo](https://plaid.coastdemo.com/share/67ed96db47d17b02bdfc4a88?zoom=100)

#### Overview

Evaluate a consumer’s income, assets, employment, and financial stability in seconds based on insights and attributes derived from cash flow data. Get comprehensive, up-to-date bank account data to grow qualified borrowers, manage risk, and make credit and verification decisions with confidence - all through a single, high-converting UX.

To see a full financial picture, Consumer Report, offered through Plaid Check (Plaid's CRA subsidiary), provides three different product bundles, with two optional add-ons. Plaid Check products are available only to US entities and can be used only to evaluate US consumers.

Plaid Check includes a variety of different modules, broken into three use cases: Income, Home Lending, and Underwriting. In addition, two modules are available as add-ons for any of the three use-case bundles.

#### Income modules

Income modules are designed for customers performing income or employment verification.

##### Base Report

Provides comprehensive bank account and cash flow data: See up to 24 months of consumer-permissioned bank account data (including account holder identity) along with categorized inflows and outflows. Get account-based attributes like balance averages and trends, as well as indicators of whether this is the borrower’s primary account based on transaction frequency.

##### Income Insights

Get details on over a dozen categorized income streams along with ready-made, model-driven attributes like historical gross and net income to streamline debt-to-income calculations and income verification. Assess financial stability with insights like forecasted income, employer name, income frequency, and predicted next payment date.

#### Underwriting modules

Underwriting modules are designed for customers performing non-mortgage underwriting, such as BNPL, personal lending, or second-look underwriting.

##### Base Report

Provides comprehensive bank account and cash flow data: See up to 24 months of consumer-permissioned bank account data (including account holder identity) along with categorized inflows and outflows. Get account-based attributes like balance averages and trends, as well as indicators of whether this is the borrower’s primary account based on transaction frequency.

##### LendScore (beta)

Based on data from both Cash Flow Insights and Network Insights, the Plaid LendScore is a score from 1-99, indicating how likely a borrower is to default over the next 12 months, across a broad range of credit products; a higher score indicates greater likelihood of repayment. The LendScore is a single, market-ready score that lenders can use to better understand the creditworthiness of applicants by layering in ability to pay and behavioral insights. LendScore also returns the top five reasons the score is not higher, for use in adverse action notices.

##### Network Insights (beta)

Receive insights based on a user's connection history, including for services that aren't consistently reported to traditional credit bureaus, such as rent, buy-now-pay-later, cash advances, and earned wage access applications. With credit insights powered by the Plaid Network, lenders receive differentiated risk signals that are complementary to traditional credit and cash flow.

##### Cash Flow Insights (beta)

Receive aggregated transaction data across all permissioned accounts. This data can be used for credit decisioning and forms the basis of the Plaid LendScore. It also includes measures of volatility across income and expenditure categories, allowing lenders to better understand how an applicant is managing their finances.

#### Home Lending modules

Home Lending modules are designed for customers performing home lending underwriting.

##### Home Lending Report

Designed specifically for home lending, Home Lending Report provides an FCRA-compliant asset verification report approved for use with Fannie Mae Day 1 Certainty. With Home Lending Report, lenders can validate asset balances, account ownership, and transaction behavior; submit reports to GSEs for rep and warrant relief; refresh reports for employment verification pre-close; and replace document uploads with a single Plaid Link session.

To create a Home Lending Report and (optionally) share it with a GSE, follow the [Home Lending integration flow](/docs/check/#home-lending-integration-flow).

#### Add-on modules

The following modules are available as add-ons for any of the above use cases.

##### Partner Insights

For scores and attributes: Get off-the-shelf cash flow risk scores based on de-identified deposit account data. Offered through [Prism Data](https://prismdata.com), these scores and attributes assess a consumer’s credit default risk and are available through a single integration and API call.

##### Cash Flow Updates (beta)

Receive regular updates about changes to cash flow. See how income and loan exposure have changed for a borrower over time to improve post-decision servicing.

#### Integration overview

If you already use one of Plaid's other products, see the product-specific migration guides for [Assets](https://plaid.com/docs/check/migrate-from-assets/), [Income](https://plaid.com/docs/check/migrate-from-income/), or [Transactions](https://plaid.com/docs/check/migrate-from-transactions/).

##### Standard integration flow

Below is a high-level integration overview. For detailed step-by-step instructions with sample code, see [Implementation](/docs/check/add-to-app/).

The flow below describes the latest integration pattern for Plaid Check Consumer Reports. If you are an existing customer who integrated before December 10, 2025, your API calls have a slightly different interface. For details on the differences, see [New User APIs](/docs/api/users/user-apis/).

1. Call [`/user/create`](/docs/api/users/#usercreate) to create a `user_id`. You must include an `identity` object when calling [`/user/create`](/docs/api/users/#usercreate). At minimum, the following fields must be provided and non-empty: `name`, `date_of_birth`, `emails`, `phone_numbers`, and `addresses` (with at least one email address, phone number, and address designated as `primary`). If you intend to share the report with a GSE (Government-Sponsored Entity) such as Fannie or Freddie, the full SSN is also required via the `id_numbers` field. Providing at least a partial SSN via `id_numbers` is also highly recommended for all use cases since it improves the accuracy of matching user records during compliance processes. The user creation step will only need to be done once per end user.
2. Call [`/link/token/create`](/docs/api/link/#linktokencreate). In addition to the required parameters, you will need to provide the following:
   - For `user_id`, provide the `user_id` you created in the previous step.
   - For `products`, pass in `cra_base_report`, plus any additional Plaid Check products you wish to use (`cra_income_insights`, `cra_partner_insights`, and/or `cra_network_insights`), along with any additional Plaid Inc. products you are using. Note that Plaid Check products can't be used in the same Link flow with Income or Assets.
   - Provide a `webhook` URL with an endpoint where you can receive webhooks from Plaid Check. Plaid Check will contact this endpoint when your reports are ready.
   - Include the appropriate options for the products you are using in the `cra_options` object, including the required [`days_requested`](/docs/api/link/#link-token-create-request-cra-options-days-requested) field.
   - Include a `consumer_report_permissible_purpose` specifying your purpose for generating an insights report.
   - On mobile, specify the `redirect_uri` and/or `android_package_name` fields as necessary, per the relevant [Link documentation](https://plaid.com/docs/link/) for your platform.
   - If you already have an existing Item and need to add Plaid Check to the Item, provide the Item's `access_token` in the `options.access_token` field. Otherwise, omit the `access_token` field in your request.
3. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see the [Link documentation](/docs/link/).
4. Launch Link with the Link token and send the user through the Link flow. Once they have completed the flow and the report is ready, you will receive a [`USER_CHECK_REPORT_READY`](/docs/api/products/check/#user_check_report_ready) webhook.
5. Upon receiving the `USER_CHECK_REPORT_READY` webhook, fetch the report data using an endpoint such as [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget), [`/cra/check_report/network_insights/get`](/docs/api/products/check/#cracheck_reportnetwork_insightsget), or [`/cra/check_report/pdf/get`](/docs/api/products/check/#cracheck_reportpdfget). This data must be fetched within 24 hours of report creation.
6. (Optional) To refresh a Consumer Report with new data, call [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) to refresh the report. Upon refresh, a new `USER_CHECK_REPORT_READY` webhook will fire and you can repeat the step above.

##### Home Lending integration flow

Below is a variation on the standard flow with a focus on home lending-related use cases and sharing reports with a GSE.

1. Call [`/user/create`](/docs/api/users/#usercreate) to create a `user_id`. You must include an `identity` object when calling [`/user/create`](/docs/api/users/#usercreate). The user creation step will only need to be done once per end user.
   - Make sure to include the user's full SSN, as it is required for GSE integrations.
2. Call [`/link/token/create`](/docs/api/link/#linktokencreate). In addition to the required parameters, you will need to provide the following:
   - Set [`gse_options.report_types`](/docs/api/link/#link-token-create-request-cra-options-base-report-gse-options-report-types) to contain `VOA`.
   - For `user_id`, provide the `user_id` you created in the previous step.
   - For `products`, pass in `cra_base_report`.
   - Provide a `webhook` URL with an endpoint where you can receive webhooks from Plaid Check. Plaid Check will contact this endpoint when your reports are ready.
   - Include the appropriate options for the products you are using in the `cra_options` object, including the required [`days_requested`](/docs/api/link/#link-token-create-request-cra-options-days-requested) field.
   - Include a `consumer_report_permissible_purpose` specifying your purpose for generating an insights report.
   - On mobile, specify the `redirect_uri` and/or `android_package_name` fields as necessary, per the relevant [Link documentation](https://plaid.com/docs/link/) for your platform.
3. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see the [Link documentation](/docs/link/).
4. Launch Link with the Link token and send the user through the Link flow. Once they have completed the flow and the report is ready, you will receive a [`CHECK_REPORT_READY`](/docs/api/products/check/#check_report_ready) webhook.
5. Upon receiving the `CHECK_REPORT_READY` webhook, fetch a Home Lending Report using [`/cra/check_report/verification/get`](/docs/api/products/check/#cracheck_reportverificationget) or a PDF copy using [`/cra/check_report/verification/pdf/get`](/docs/api/products/check/#cracheck_reportverificationpdfget). This data must be fetched within 24 hours of report creation.
6. Use the [`/oauth/token`](/docs/api/oauth/#oauthtoken) endpoint to create a sharing token for the GSE partner you want to share report data with. For more details, see [Sharing Consumer Report data with partners](/docs/check/oauth-sharing/).
7. Send the OAuth token to your GSE partner and continue through their flow.
8. (Optional) To refresh a Consumer Report with new data, call [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) to refresh the report. Upon refresh, a new `CHECK_REPORT_READY` webhook will fire and you can repeat step 5 above.
   - After refreshing the report, to share the refreshed report with a GSE, you must create a new sharing token by repeating steps 6-7 above.

Example /link/token/create call for home lending use case

```
curl -X POST https://sandbox.plaid.com/link/token/create -H 'Content-Type: application/json' -d '{
  "client_id": "${PLAID_CLIENT_ID}",
  "secret": "${PLAID_SECRET}",
  "user": {
    "client_user_id": "user-abc",
    "email_address": "user@example.com"
  },
  "user_id": "usr_9nSp2KuZ2x4JDw"",
  "products": ["cra_base_report"],
  "consumer_report_permissible_purpose": "ACCOUNT_REVIEW_CREDIT",
  "cra_options": {
    "days_requested": 365,
    "base_report": {
      "gse_options": {
        "report_types": ["voa"]
      }
    }
  },
  "gse_options": {
    "report_type": ["VOA"]
  },
  "client_name": "The Loan Arranger",
  "language": "en",
  "country_codes": ["US"],
  "webhook": "https://sample-web-hook.com"
}'
```

##### Cash Flow Updates integration

To receive Cash Flow Updates (beta) on a Plaid Check-enabled user, follow the steps in the [Standard integration flow](/docs/check/#standard-integration-flow), then perform the following additional steps:

1. Call [`/cra/monitoring_insights/subscribe`](/docs/api/products/check/#cramonitoring_insightssubscribe), passing in the `user_id`.
2. Listen for the [`CASH_FLOW_INSIGHTS: CASH_FLOW_INSIGHTS_UPDATED`](/docs/api/products/check/#cash_flow_insights_updated) webhook, which will be sent between one and four times a day for a user's Item.
3. Call [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) to retrieve the updated insights.
4. (Optional) To stop receiving `CASH_FLOW_INSIGHTS_UPDATED` webhooks for a user, call [`/cra/monitoring_insights/unsubscribe`](/docs/api/products/check/#cramonitoring_insightsunsubscribe).

##### No-code integration with the Credit Dashboard

The Credit Dashboard is currently in beta. To request access, contact your Account Manager.

Consumer Report data is available via a no-code integration through the [Credit Dashboard](https://dashboard.plaid.com/credit). You can fill out a form with information about your end user, such as their name, email address, address, and phone number; and indicate your basis for processing this data and how much history you would like to collect. Plaid will then generate a hyperlink that can be used to launch a Plaid-hosted Link session. Plaid can email the link directly to your user (Production only, not available in Sandbox), or you can send it to them yourself.

After the end user successfully completes the Link session, their data will be available in the Dashboard for you to view, archive, or delete. The data shown will be the same data returned by a regular Consumer Report API call, in a user-friendly web-based session.

![Example dashboard credit report for end user, showing forecasted annual and monthly incomes. Income sources listed with types and frequency.](/assets/img/docs/credit-dashboard-view.png)

Once you are enabled for the Credit Dashboard, all new sessions going forward, including ones created via the API, will be displayed in the Dashboard.

Data from Link sessions created via the Credit Dashboard cannot be accessed via the Plaid API. If you require programmatic access to data, use the code-based [Standard integration flow](/docs/check/#standard-integration-flow) instead.

##### Integrating with other Plaid products

If you are using other Plaid Inc. products, note that the `account_id` returned in API responses from endpoints prefixed with `/cra/` will not match the `account_id` returned in responses from non-CRA Plaid endpoints. You can map accounts between Plaid CRA endpoints and Plaid endpoints by matching other fields such as `mask`, account `type` and `subtype`, and `institution_id`. If all of these fields match, and the accounts are associated with the same user in your records, the accounts are very likely the same.

##### Sharing reports with partners via OAuth

To share a report with a partner such as Experian, Fannie Mae, or Freddie Mac, see [Share reports with partners](/docs/check/oauth-sharing/).

#### Testing Consumer Report

Access to Consumer Report in Sandbox is not granted by default. To request access, [contact Sales](https://plaid.com/check/income-and-underwriting/#contact-form) or, for existing Plaid customers, contact your account manager or submit a [product access request](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access). Customers with [Production access to Consumer Report](https://dashboard.plaid.com/settings/team/product) will also automatically be granted Sandbox access.

For details on how to test specific scenarios, including multiple income streams, specific employer names, or cash flow update events, see [Add Consumer Report to your App: Testing in Sandbox](/docs/check/add-to-app/#testing-in-sandbox).

##### Testing Income Insights

Use the [Sandbox credit testing credentials](/docs/sandbox/test-credentials/#credit-and-income-testing-credentials) with a non-OAuth test institution, such as First Platypus Bank, for more realistic and valid data.

To customize the employer name during testing, you can test with a custom Sandbox user, following the custom Sandbox user [documentation](/docs/sandbox/user-custom/) and [examples](https://github.com/plaid/sandbox-custom-users/). When creating an income transaction for this user:

- The `amount` must be negative, so that the transaction will represent income.
- `date_transacted` and `date_posted` must fall within the `cra_options.days_requested` range of the [`/link/token/create`](/docs/api/link/#linktokencreate) call.
- the `description` field must contain the desired employer name, followed by  `Direct Dep`.

The example transaction below will generate an income stream with the employer name `Weyland-Yutani`.

Sample income transaction for testing employer name

```
{
  "date_transacted": "2025-08-01", 
  "date_posted": "2025-08-03",
  "amount": -5500,
  "description": "Weyland-Yutani Direct Dep",
  "currency": "USD"
 },
```

#### Billing

For details, see [Plaid Check fee model](/docs/account/billing/#plaid-check-fee-model).

#### Next steps

To get started building with Plaid Check, see [Add Consumer Report to your App](/docs/check/add-to-app/).

If you're migrating from another product, see the [Assets migration guide](/docs/check/migrate-from-assets/), [Income migration guide](/docs/check/migrate-from-income/), or [Transactions migration guide](/docs/check/migrate-from-transactions/).

If you're ready to launch to Production, see the Launch Center.

[#### Launch Center

See next steps to launch in Production

Launch](https://dashboard.plaid.com/developers/launch-center)

#### Launch Center

See next steps to launch in Production

[Launch](https://dashboard.plaid.com/developers/launch-center)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
