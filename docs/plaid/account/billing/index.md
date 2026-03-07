---
title: "Account - Pricing and billing | Plaid Docs"
source_url: "https://plaid.com/docs/account/billing/"
scraped_at: "2026-03-07T22:03:45+00:00"
---

# Plaid pricing and billing

#### Learn about pricing, what is considered a billable event, and how to monitor your bill

#### Pricing information

A price list is not available in the documentation. **To view pricing, [apply for Production access](https://dashboard.plaid.com/overview/request-products). Pricing information for Pay as you go and Growth plans will be displayed on the last page before you submit your request.** For Custom plans, select the Custom option and submit the form, and Sales will reach out to discuss pricing. Or, you can [contact Sales](https://plaid.com/contact/) to learn more about custom pricing.

#### Pricing plans

Plaid offers three types of pricing plans:

- **Pay as you go** – has no minimum spend or commitment. Most appropriate for hobbyist use, or for early stage small businesses without validated business models or investment capital.
- **Growth** – Minimum spend, annual commitment. Lower per-use costs than Pay as you go plans; includes business-focused features such as SSO, priority support, and a personal account manager. Most appropriate for small to medium size businesses with API usage volumes up to $6,000/month.
- **Custom (aka Scale)** – Higher minimum spend and annual commitment but provides access to the lowest per-use costs. Minimums can apply across the entire Plaid account or by product and can apply across both Plaid and Plaid Check. Most appropriate for businesses with API usage volumes over $2,000 / month or that require enterprise-level functionality. For a Custom plan, [contact Sales](https://plaid.com/contact/).

To change your plan, see [Viewing or changing pricing plans](/docs/account/billing/#viewing-or-changing-pricing-plans).

For customers based in the EU or UK, or who will be serving end users based in the EU or UK, only Custom plans are offered. In addition, the following products are only offered via Custom plans:

- Signal Transaction Scores
- Transfer
- Document Income with bank statements support
- All Plaid Check products except for Base Reports and Income Insights Reports\*
- Payment Initiation and Virtual Accounts
- Layer (pre-GA; contact Sales for details)
- Protect (pre-GA; contact Sales for details)

\*No Plaid Check products are available on Pay as you go plans. Base Reports and Income Insights Reports are available via either Growth or Custom Plans. All other Plaid Check products (LendScore, Partner Insights, Network Insights, etc.) are available via Custom plans exclusively.

#### Pricing models

Plaid has multiple different pricing models, depending on the product. These models only apply to Production traffic; usage of Sandbox is always free.

Note that the details outlined below are general information and guidance about typical pricing structures and policies for new customers. Your specific pricing and billing structure is governed by your agreement with Plaid. If you have questions about your bill, [contact Support](https://dashboard.plaid.com/support/new/admin/account-administration/pricing).

Endpoints that are not associated with any particular product are typically available to Plaid customers free of charge. Examples of these free-to-use endpoints include [Institutions endpoints](/docs/api/institutions/), [Item endpoints](/docs/api/items/), [`/accounts/get`](/docs/api/accounts/#accountsget) (not to be confused with [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget), which is associated with the Balance product), and [user management endpoints](/docs/api/users/).

##### One-time fee

Products that use one-time fee pricing are:

- Auth
- Identity
- Income (except for [`/credit/bank_income/refresh`](/docs/api/products/income/#creditbank_incomerefresh) and [`/credit/payroll_income/refresh`](/docs/api/products/income/#creditpayroll_incomerefresh), which use per-request pricing)
- Layer

You will incur charges for one-time fee products whenever the product is successfully added to an Item. This occurs when an Item has been successfully created by Link and this product was specified in [`/link/token/create`](/docs/api/link/#linktokencreate). An Item is typically considered to be successfully created if a `public_token` was exchanged, but sometimes (e.g. when using micro-deposit Auth flows, or the in-Link Identity Match flow) it may be considered created once the `public_token` was created. If the Item was not initialized with the product at the time of creation, the product can also be added to the Item later by calling a product endpoint belonging to that product on the Item.

For one-time fee products, it does not matter how many API calls are made for the Item (or user); the cost is the same regardless of whether the product's endpoints are called many times or zero times.

For the Auth and Identity products, if the product was added to the Item via the `optional_products` parameter in the [`/link/token/create`](/docs/api/link/#linktokencreate) call, you will not incur charges for those products until their corresponding endpoint is called. For example, if you initialize an Item with `auth` in the `products` parameter and `identity` in the `optional_products` parameter, you will not be charged for Identity on that Item until you call [`/identity/match`](/docs/api/products/identity/#identitymatch) or [`/identity/get`](/docs/api/products/identity/#identityget) for that access token. Note: Identity Match is a Per-request flat fee product, so calling [`/identity/match`](/docs/api/products/identity/#identitymatch) multiple times will result in repeated charges for Identity Match, but only one charge for Identity.

For Document Income and Payroll Income, charges will be incurred when data is available to be accessed. This will typically be indicated by the `INCOME_VERIFICATION` webhook firing with the status of `VERIFICATION_STATUS_PROCESSING_COMPLETE`. For Document Income, the fee depends on the documents processed and which processing options are enabled (i.e. fraud, document parsing, or both). For bank statements, the fee is based on the number of pages in the statement; for all other document types, there is a flat fee per document.

For Hosted Link, each session delivered via SMS or email will incur a one-time fee. There is no fee to use Hosted Link if you are not using Plaid to deliver sessions.

When using Transfer, each Item initialized with Transfer (unless created using the [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) endpoint) will also incur the one-time Auth fee, and no additional fee will be incurred for using Auth endpoints with that Item.

When using Layer, a billing event is incurred for each converted Link session (when `onSuccess` fires). You will not be billed for Layer eligibility checks or unconverted Layer sessions.

##### Subscription fee

Products that use subscription fee pricing are:

- Transactions (except for the [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) endpoint)
- Recurring Transactions
- Liabilities
- Investments (except for the [`/investments/refresh`](/docs/api/products/investments/#investmentsrefresh) endpoint)

Under the subscription fee model, an Item will incur a monthly subscription fee as long as a valid `access_token` exists for the Item. Removing the Item via [`/item/remove`](/docs/api/items/#itemremove) will end the subscription. The subscription will also be ended if the end user depermissions their Item, such as via [my.plaid.com](https://my.plaid.com/) or by contacting Plaid support. If the end user depermissions the Item via their financial institution's portal, this may also end the subscription, although this is not guaranteed, as not all institutions notify Plaid about permissions revocations performed via their portals.

Once a subscription fee product has been added to an Item, it is not possible to end the subscription and leave the Item in place. The Item must be deleted, which can be done by calling [`/item/remove`](/docs/api/items/#itemremove). If the Item's subscription is active, Plaid will charge for the subscription even if no API calls are made for the Item or API calls cannot be successfully made for the Item (e.g. because the Item is in an error state). Plaid's subscription cycle is based on calendar months in the UTC time zone; each month begins a new cycle. Fees for Items created or removed in the middle of the month are not pro-rated.

If you add a subscription-billed product to an Item using your free API request allocation (for example, by creating an Item initialized with Transactions, Investments, or Liabilities in Limited Production), you will be not be charged for that subscription unless you continue to use the Item after using up your free API requests and receiving full Production access.

When using subscription fee products in Production, it is important to persist the access token so that you can remove the Item when needed to avoid being billed indefinitely. If you have created access tokens for which you are being billed but did not store the tokens, [contact support](https://dashboard.plaid.com/support) for assistance.

A single Item can have multiple different subscriptions on it (for example, Transactions and Liabilities), but the same subscription will not be added multiple times to the same Item.

Investments has two separate subscriptions that can be associated with an Item: Investments Holdings and Investments Transactions. Adding Investments to an Item via [`/link/token/create`](/docs/api/link/#linktokencreate) or by calling [`/investments/holdings/get`](/docs/api/products/investments/#investmentsholdingsget) adds the Investments Holdings subscription. Calling [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) on an Item adds both the Investments Transactions and Investments Holdings subscriptions.

##### Per-request flat fee

For per-request flat fee products, a flat fee is charged for each successful API call to that product endpoint.

Products and endpoints that use per-request flat fee pricing are:

- Balance ([`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) or [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate))
- Signal Transaction Scores ([`/signal/evaluate`](/docs/api/products/signal/#signalevaluate))
- Investments Move ([`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget))
- Transfer Signal or Balance fee ([`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate)), see [Transfer fee model](/docs/account/billing/#transfer-fee-model) for more details
- Refresh endpoints, except for Statements Refresh ([`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh), [`/investments/refresh`](/docs/api/products/investments/#investmentsrefresh), [`/credit/bank_income/refresh`](/docs/api/products/income/#creditbank_incomerefresh), [`/credit/payroll_income/refresh`](/docs/api/products/income/#creditpayroll_incomerefresh))
- Asset report PDF and Audit Copy ([`/asset_report/audit_copy/create`](/docs/api/products/assets/#asset_reportaudit_copycreate), [`/asset_report/pdf/get`](/docs/api/products/assets/#asset_reportpdfget))
- Identity Match ([`/identity/match`](/docs/api/products/identity/#identitymatch)), see below for more details
- Beacon Account Insights ([`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget))

There is also a flat fee for each link delivered through SMS or email via Link Delivery (beta), when using Hosted Link with the optional Link Delivery feature enabled.

Identity Match can optionally be enabled via the Dashboard. In this integration mode, the Identity Match check is automatically performed during Link and you never call the [`/identity/match`](/docs/api/products/identity/#identitymatch) endpoint directly. The billing trigger for Identity Match in this situation is when Plaid successfully obtains Identity data from the Item. If Plaid could not obtain any Identity data for the Item, you will not be billed.

The [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) endpoint is a billable endpoint that is used by two different products, Balance and Signal Transaction Scores. It will be billed as a Balance API call if used with a Balance-only ruleset. It will be billed as a Signal Transaction Scores API call if used with a Signal Transaction Score-powered ruleset, or if used with no ruleset on a Signal Transaction Scores-enabled account.

Calls to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) that use a ruleset (including the default ruleset) will be billed when the ruleset is evaluated, even if some or all data could not be extracted from the Item.

##### Per-request flexible fee

Products that use per-request flexible fee pricing are:

- Assets
- Statements Refresh
- Enrich

For per-request flexible fee products, a fee is charged for each successful API call to that product endpoint. The fee will vary depending on the amount of information requested in the API request.

For Assets, the fee is calculated based on the number of Items in the report as well as the number of days of history requested. An additional fee is charged for an Asset Report containing more than 5 Items; this fee will be charged twice if the Asset Report contains more than 10 Items, and so on. An "Additional History" fee is also charged for each Item for which more than 61 days of history is requested. The Additional History fee is a flat fee regardless of how many days of additional history are requested; it does not matter how many days beyond 61 are requested.

For Enrich, the flexible fee is based on the number of transactions sent to be enriched.

For [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh), the flexible fee is based on the number of statements available between the provided start and end dates found when calling [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh). Note that you will be charged for any statement extracted, even those that you previously requested at an earlier date.

##### Per-Item flexible fee

Products that use per-Item flexible fee pricing are:

- Statements (excluding [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh), which is charged as a per-request flexible fee)

For per-Item flexible fee products, a fee is charged when the Item is created, which is deemed to occur when the `public_token` is created. The fee will vary depending on the amount of information requested when creating the Item.

For Statements, the flexible fee is calculated based on the number of statements available within the date range requested when calling [`/link/token/create`](/docs/api/link/#linktokencreate). The fee will be charged even if you do not call any Statements endpoints.

##### Payment Initiation fee model

A fee is charged for each payment initiated. A payment is considered initiated if the end user reached the end of the payment flow and received a confirmation success screen. Standing orders (recurring payments) are considered a single payment but are billed on a separate pricing schedule from one-time payments. Fees will vary depending on the payment amount and the network used to transfer the payment.

##### Plaid Check fee model

Most Plaid Check products are charged when the corresponding /get endpoint is called ([`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), [`/cra/check_report/network_insights/get`](/docs/api/products/check/#cracheck_reportnetwork_insightsget), and [`/cra/check_report/pdf/get`](/docs/api/products/check/#cracheck_reportpdfget)). The exception is Partner Insights, which is charged when the associated report is first generated and `partner_insights` is specified in the product array, which can happen during [`/link/token/create`](/docs/api/link/#linktokencreate), [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate), or [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget).

The Plaid Check fee model is similar to the one-time fee model. Like the one-time fee model, there is no additional charge to call an endpoint multiple times if the same Consumer Report is being retrieved each time. However, unlike the one-time fee model, if a new Consumer Report is generated for a given `user_id` or `user_token`, which can be done by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate), a new set of fees will be incurred for calling a `/get/` endpoint on the new Consumer Report.

In addition to any product-specific charges, each Plaid Check Report incurs a Base Report charge. This fee is charged when the first `/get` endpoint is called for a given `user_id` or `user_token` and report.

For example, if you called [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget) followed by [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), you would be charged the Base Report fee when [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget) is called, but only charged the Income Insights fee when [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget) is called. If you call the endpoints in the opposite order, calling [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget) followed by [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), both the Base Report fee and Income Insights fee would be charged when [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget) is called, and no fees would be charged when calling [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget).

For certain use cases, Plaid Check customers can select between different fee and billing models to accommodate different use cases and provide cost predictability. For details, [contact Sales](https://plaid.com/check/income-and-underwriting/#contact-form).

##### Transfer fee model

When working with Transfer, there are three fees: The Auth fee, the Signal fee, and the per-transfer fee.

Existing contracts created or renewed prior to October 1, 2025 may contain references to a Transfer Risk Engine fee, which is assessed for each call to [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate). For all new Transfer contracts going forward, this fee has been replaced with the Signal fee (which may also be called the Signal Transaction Scores fee or Balance fee, depending on which product(s) you are using) .

###### Auth fee

When using Transfer, each Item with the Transfer product (unless created by the [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) endpoint) will also have Auth added to the Item and will incur a one-time Auth fee in addition to the per-payment Transfer fee.

###### Signal fee

Calls to [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) or [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) will incur per-request flat fees for Balance or Signal Transaction Scores, depending on the ruleset invoked. A fee is incurred for every successful invocation of a ruleset. However, if the `authorization.decision` returned in the response is `user_action_required`, you will not be charged.

###### Per-transfer fee

A fee is charged for each transfer made. A transfer is considered made if it was successfully created. This occurs when [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) is successfully called, or if using Transfer UI instead of [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate), when the `onSuccess` callback is fired from the Link session where the transfer is authorized. The per-transfer fee is charged even if the payment is later reversed or clawed back. In such cases, a reversed payment fee may be charged as well. Fees will vary depending on the payment amount and the network used to make the payment.

A per-transfer fee is also applied to funds transfers between your bank account and your balance held with Plaid, such as sweeps, or payments to and from the Plaid Ledger.

##### Monitor fee model

Monitor uses a model with two fees:

- A base fee, which is incurred the first time a new user is scanned
- A rescanning fee, which is incurred based on the number of users rescanned each month

The rescanning fee is calculated based on the number of users on a rescanning program in a given month and works similarly to the subscription fee model. Plaid's subscription cycle is based on calendar months in the UTC time zone; each month begins a new cycle. Fees for users added to or removed from a rescanning program in the middle of the month are not pro-rated.

To avoid being rebilled for a user you no longer need recurring scans for (e.g., a user who has closed their account with you), create a program specifically for former users, make sure rescans are disabled, and move the user into that program when they are offboarded from your system.

##### Identity Verification fee model

For Identity Verification, a fee is charged based on the end user performing certain activities during the Link flow. The following events are billed:

- Anti-Fraud Engine (the first verification check run, typically when the end user enters their phone number to begin SMS verification. If this is skipped, the Anti-Fraud Engine fee will be incurred later on. For example, if the SMS check is skipped in favor of the document check, the document check will incur both the Anti-Fraud Engine fee and the document check fee.)
- Data Source verification (cost varies based on the end user's country, see your pricing contract for details)
- Document check
- Selfie check

If a retry is issued for a session, the retry will be billed like a new session, including the Anti-Fraud Engine and any other verification checks included in the retried session.

#### Partnerships pricing

If you are using a Plaid partner, you will be charged for Plaid API usage performed by the partner using your processor token in the same way as if the calls had been made by you. The partner's API usage will not be included in your Plaid Dashboard [usage report](https://dashboard.plaid.com/activity/usage). If you have questions about the partner's API activity, contact the Plaid partner.

#### Viewing billable activity and invoices

To view billable activity for your account, see the Production [Usage](https://dashboard.plaid.com/activity/usage) page on the Dashboard.

Invoices will be sent each month to the [billing contact](https://dashboard.plaid.com/settings/company/profile) configured for the team.

#### Updating payment information

Pay as you go customers can update [payment method](https://dashboard.plaid.com/settings/team/billing) and [billing contact](https://dashboard.plaid.com/settings/company/profile) information on the Plaid Dashboard. Custom plan customers should contact their account manager to update this information.

#### Updating products

To view your currently enabled products, see the [Products](https://dashboard.plaid.com/settings/team/products) page on the Dashboard under the Team Settings section. To add additional products to your Plaid account you can [submit the product request form](https://dashboard.plaid.com/overview/request-products) located on the Dashboard. This form is also accessible via the [Products](https://dashboard.plaid.com/settings/team/products) page.

#### Viewing or changing pricing plans

To view current pricing or to switch between pricing plans, customers on Pay as you go or Growth plans can use the [Plans page](https://dashboard.plaid.com/settings/team/plans) within the Dashboard. Customers with Pay as you go plans can upgrade their plan at any time. Customers with Growth plans can submit a request to upgrade or downgrade their plan at any time; requests will take effect after the minimum commitment period on the existing Growth plan has been satisfied. Customers on Custom plans should [contact Support](https://dashboard.plaid.com/support/new/product-and-development) or their Plaid Account Manager.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
