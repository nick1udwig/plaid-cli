---
title: "Changelog | Plaid Docs"
source_url: "https://plaid.com/docs/changelog/"
scraped_at: "2026-03-07T22:04:47+00:00"
---

# Plaid changelog

#### Track changes to the Plaid API and products

#### Overview

This changelog tracks updates to the Plaid API and changes to the Plaid mobile SDKs, Link flow, functionality, and APIs. The changelog is updated at least once per month. Updates that affect only products or features in beta or limited release may not always be reflected. Improvements to the Dashboard, institution connectivity / Item health, or data availability and quality may not always be reflected.

Link SDKs are released on GitHub. A summary of updates will be posted here. For a more complete and detailed record of Link SDK changes, see the GitHub changelogs: [iOS](https://github.com/plaid/plaid-link-ios/releases), [Android](https://github.com/plaid/plaid-link-android/releases), [React Native](https://github.com/plaid/react-native-plaid-link-sdk/releases), [React](https://github.com/plaid/react-plaid-link).

Plaid's officially supported libraries are updated frequently. For details, see the GitHub changelogs: [Python](https://github.com/plaid/plaid-python/blob/master/CHANGELOG.md), [Node](https://github.com/plaid/plaid-node/blob/master/CHANGELOG.md), [Ruby](https://github.com/plaid/plaid-ruby/blob/master/CHANGELOG.md), [Java](https://github.com/plaid/plaid-java/blob/master/CHANGELOG.md), [Go](https://github.com/plaid/plaid-go/blob/master/CHANGELOG.md).

#### Subscribe to updates

[Subscribe to this changelog via RSS](https://cdn.feedcontrol.net/8259/13425-kjU6T4Te5FUvO.xml).

##### February 26, 2026

- Updated the error object inside of webhooks to return `display_message: null` rather than omitting the `display_message` field if there is no associated display message, in order to match the documented behavior, as well as the behavior of error objects that are returned by the Plaid API outside of a webhook context.

For Auth:

- Stripe has announced that it will begin to migrate existing customers off of the older Charges and Sources API. Plaid has directly reached out to all impacted customers using the Stripe Charges and Sources x Plaid Auth integration with a list of options for moving forward. If you believe you may be impacted but did not receive an email, contact your Plaid account manager or support. Note that this change does not impact users of the newer Stripe Payment Intents integration (most customers who integrated Stripe x Plaid in 2024 or later are using this newer API). In addition, new customers integrating with Stripe x Plaid for the first time must use the Payment Intents API going forward and can no longer use Charges and Sources.
- Released access to Database Auth via API, using the [`/auth/verify`](/docs/api/products/auth/#authverify) endpoint. This feature is currently in Early Availability; to request access, contact your Plaid account manager.

For Plaid Check Consumer Report:

- Beginning March 3, 2026, the `consumer_report_permissible_purpose` field must be included in [`/link/token/create`](/docs/api/link/#linktokencreate) when creating a Link token for a Consumer Report product. Existing customers who are impacted by this change have been contacted by their account managers and may have arranged custom migration deadlines.
- Announced April 1, 2026 as the beginning date for the optional migration to [User APIs](https://plaid.com/docs/api/users/user-apis/). To enable this migration, on April 1, existing customers on the older APIs will begin to get both the new and old webhook versions for revised webhooks.

##### February 18, 2026

- Added [Domains](/docs/account/teams/#team-domains) functionality to Dashboard, enabling the ability to associate a domain with a Plaid team and automatically invite all new Dashboard users with a verified email address matching that domain to join your Plaid team.
- Added `user_id` field to webhook payloads for the following webhooks: `SYNC_UPDATES_AVAILABLE`, `PENDING_EXPIRATION`, `ERROR`, `LOGIN_REPAIRED`, `INVESTMENTS_TRANSACTIONS: DEFAULT_UPDATE`, `INVESTMENTS_TRANSACTIONS: HISTORICAL_UPDATE`, `HOLDINGS: DEFAULT_UPDATE`, `LIABILITIES: DEFAULT_UPDATE`, `PENDING_DISCONNECT`, `USER_PERMISSION_REVOKED`, `NEW_ACCOUNTS_AVAILABLE`, `USER_ACCOUNT_REVOKED`, `SESSION_FINISHED`.
- Released the [React Native SDK version 12.8](https://github.com/plaid/react-native-plaid-link-sdk/), with bug fixes for Layer events and Android Proguard rules.
- Announced the dates of the Bank of America API migration. From February through mid March 2026, the new API is being gradually rolled out for new Items. Throughout the period of mid March 2026 - late October 2026, existing Items will gradually be disconnected from the old Bank of America API and must be connected to the new Bank of America API by going through update mode. To minimize disruptions, listen for the `PENDING_DISCONNECT` webhook. Upon receiving the webhook, prompt the user to go through the Update Mode flow for the Item, which will migrate the Item to the new API. One week after the `PENDING_DISCONNECT` webhook was fired for a given Item, if the Item has not yet gone through update mode, it will be disconnected from the old API and will enter the `ITEM_LOGIN_REQUIRED` error state. Sending the Item through update mode will move it to the new API and restore it to a healthy state.

For Plaid Check Consumer Report:

- Added new permissible purpose for Consumer Report generation, `ELIGIBILITY_FOR_GOVT_BENEFITS`, to represent reports generated in connection with an eligibility determination for a government benefit where the entity is required to consider an applicant’s financial status pursuant to FCRA Section 604(a)(3)(D).

For Layer:

- In the `/user_accounts/session/get` response, added `edits_current` field to the `identity_edit_history` object, indicating how many times the user manually edited the prefilled data provided by Layer.

##### February 4, 2026

For Identity Verification:

- For the Document Verification flow, enhanced accuracy at detecting when the document presented is not a physical document. These enhancements will be applied automatically; customers do not need to take any action.

For Layer:

- Released Intelligent Account Sorting. By default, Layer orders accounts in Link based on which one is likeliest to have the highest conversion. With Intelligent Account Sorting, you can instead prioritize either the highest-balance account or the account that Plaid detects as being the user's primary bank account. The prioritized account will have a "Recommended" badge displayed next to it in Link. Intelligent Account Sorting may be a good fit for customers are using Layer in flows for cashflow underwriting, EWA, or cash advance. To enable Intelligent Account Sorting, contact your Plaid Account Manager, and specify which prioritization scheme (primary account or highest balance account) you prefer and which Link customizations the scheme should be applied to.

For Protect:

- Added `environment` field to Protect webhooks.

For Signal:

- Signal is now visible in the Dashboard left navigation for all customers. This change allows all customers to build with Signal Rules in Sandbox, even before they have been approved for Production access.

For Virtual Accounts:

- Added `payee_verification_status` field to `WalletTransaction` schema. This field indicates the result of payee verification checks for EUR payouts.

##### January 16, 2026

- Starting in late January or early February 2026, Plaid will begin enforcing data validity for user information included in [`/link/token/create`](/docs/api/link/#linktokencreate) API requests. If a phone number, address, name, or email sent in the user object is malformatted, the call will fail with the error `INVALID_USER_IDENTITY_DATA`. This will not impact integrations that do not send user PII in [`/link/token/create`](/docs/api/link/#linktokencreate).
- Released [iOS SDK](https://github.com/plaid/plaid-link-ios) v6.4.3, adding support for certain experimental and/or limited-availability Layer features.

For Transfer:

- Starting in early February 2026, the cutoff time for Same-Day ACH transfers will move from 3:30pm ET to 3:00pm ET.

##### January 2, 2026

For Identity Verification:

- Added "Explore" Dashboard view, allowing for backtesting and exploration of rule outcomes and the ability to build custom rules based on specific user and session attributes. To access Explore, go to the [Plaid Protect section of the Dashboard](https://dashboard.plaid.com/protect). This view is available to all Identity Verification customers based in the US or UK.

##### December 16, 2025

For Identity Verification:

- Added `facial_duplicates` object, containing details about facial duplicate results, to `risk_check` object.
- Made Trust Index scores accessible via the API, as well as via the Dashboard.

For Plaid Check Consumer Report:

- Redesigned the PDF report returned by [`/cra/check_report/verification/get`](/docs/api/products/check/#cracheck_reportverificationget), including a detailed overview of a borrower’s connected accounts with current and average balances, a breakout of deposits and withdrawals, and total inflows/outflows for each account.

##### December 10, 2025

- Announced upcoming Bank of America API migration and consent expiration. Existing Bank of America Items can be migrated to the new Bank of America API beginning in early 2026 and must eventually be migrated to avoid entering an `ITEM_LOGIN_REQUIRED` state. The exact beginning and end dates of the migration will be announced in 2026. Bank of America Items on the new API will have a 12-month consent expiration applied.
- Released Plaid's new User APIs, supporting all new integrations of Plaid Check, Multi-Item Link, and Plaid Protect. The User APIs create a single, simplified representation of a user via the `user_id` field and will be the representation used by the Plaid API going forward. Customers with existing integrations should continue to use user tokens; migration instructions will be released in 2026. See [New User APIs](/docs/api/users/user-apis/) for more details.

##### December 3, 2025

For Link:

- Ended support for Modular Link, a specific form of Link only available in Europe.
- Improved the ability of the returning user experience to recognize a returning user on the same device.

For Layer:

- Added support for the `SESSION_FINISHED` webhook, allowing backend detection of when a session is complete and `public_token` delivery.

Added the following template improvements:

- Layer template support for `required_if_supported_products`, `optional_products`, and `additional_consented_products`.
- A custom preferred institutions list, allowing you to select which institutions are prioritized.
- Customizing whether the SSN shown in Link is masked
- The ability to set criteria to pre-select certain accounts in Link, like those with the highest balance or recent income, to encourage users to share those. (This feature is in closed beta; contact your account manager for access.)

For Transactions:

Released Personal Finance Categories v2 (PFCv2). PFCv2 includes both improved transaction categorization accuracy (approximately 10% improvement at the category level and 20% improvement at the subcategory level) and a new, more granular transaction schema with additional categories specifically useful for Earned Wage Access (EWA) use cases, with six additional income subcategories, six additional loan disbursement subcategories, and three additional loan repayment subcategories. Additional subcategories have also been added for bank fees.

To opt in to PFCv2, existing customers should set `personal_finance_category_version` to `v2` in the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), [`/transactions/get`](/docs/api/products/transactions/#transactionsget), or [`/transactions/enrich`](/docs/api/products/enrich/#transactionsenrich) request. Opting in to PFCv2 is required for existing customers to receive the new, more granular subcategories. To receive the PFCv2 accuracy improvements, customers do not need to opt in to PFCv2 (with the exception of a small subset of customers in the EWA industry). All new customers enabled for Transactions after December 3, 2025 will be opted in to PFCv2 by default.

##### November 12, 2025

For Identity Verification:

- Added `latest_scored_protect_event` field to response objects in the Identity Verification endpoints.

For Plaid Check Consumer Report:

- Improved data normalization of the `employer_name` field.
- If the data retrieved by Plaid for a Consumer Report shows strong signs of being inaccurate (e.g. balance history and transaction history inconsistent with each other, or transaction history contains duplicate transactions), the report generation will now fail with a `DATA_QUALITY_CHECK_FAILED` error. This is estimated to impact approximately one in one thousand report generation attempts. If you experience this error, wait at least 24 hours and then try again.
- For [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), added `bank_income_source.status` and `bank_income_source.income_provider` fields.
- For [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), make `income_provider` nullable.

##### November 4, 2025

- Fixed an issue in which unhealthy Items would occasionally continue to be reported as healthy and never enter the `ITEM_LOGIN_REQUIRED` state.
- Delays with Chase OAuth registration have been resolved. Chase registration will now typically complete in approximately 1-2 business days for new clients. If you have received Production access and completed your Security Questionnaire over two weeks ago and do not yet have Chase access, contact your Account Manager or file a support ticket.

For Link:

- Released [iOS SDK](https://github.com/plaid/plaid-link-ios) v6.4.2, resolving a syncFinanceKit crash when running on iPad on compatibility mode.
- Released [Android SDK](https://github.com/plaid/plaid-link-android) v.5.5.0, with the following improvements: made `LinkErrorCode.errorType` public, fixed bug where Layer "auto" customization for light/dark mode was always dark, and added `onLoad` callback to `Plaid.create` for detecting when Link is ready to present.
- Released [React Native SDK](https://github.com/plaid/react-native-plaid-link-sdk) v12.6.1, updating the iOS SDK to v6.4.2 and improving internal logging and debugging to help diagnose customer issues more effectively.
- Released [React Native SDK](https://github.com/plaid/react-native-plaid-link-sdk) v12.7.0. This version adds `onLoad` to `LinkTokenConfiguration`, fired once when Link is fully loaded and ready to present. You can use it to manage your own loading UI or defer presentation until ready.

For Auth:

- PNC TAN expiration, previously scheduled to begin in January 2026, has been postponed until further notice. Further updates on PNC TAN expiration are not expected until 2026. The changelog will be updated when new information is available.

For Assets:

- Added Account Insights to Asset Reports in Europe. Account Insights includes many fields providing detailed information on a end user's credit risk and affordability, including insights about spending activity, income, atypical transactions, minimum and negative balances, loan payment activity, and gambling activity. Account Insights is only available in Europe. US customers seeking similar information should use Plaid Check Consumer Report.

For Plaid Check Consumer Report:

- Added support for investments holdings and accounts in reports, such as Home Lending Report.

For Identity Verification:

- For Document Verification, `issue_date` is now exposed in the API where relevant; previously this information was only available via the Dashboard.

For Transfer:

- Partners can now use access tokens belonging to one of their end customers to call [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate), allowing them to more easily initiate transfers even if the end customer performed the initial Item creation.
- PNC TAN expiration, previously scheduled to begin in January 2026, has been postponed until further notice. Further updates on PNC TAN expiration are not expected until 2026. The changelog will be updated when new information is available.

##### October 15, 2025

For Balance, Transfer, and Signal:

- Renamed "Signal" product to "Signal Transaction Scores". "Plaid Signal" now refers to Plaid's suite of capabilities used for assessing ACH return risk, which includes both Signal Transaction Scores and Balance.
- Streamlined Balance and Signal Transaction Scores integrations with the creation of Balance-only Signal rulesets and a single shared integration path for the Plaid Signal suite (Balance and Signal Transaction Scores) using the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) endpoint.
- Balance customers, including those who do not use Signal Transaction Scores, can now integrate using the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) endpoint and Signal Rules in the Dashboard to take advantage of no-code configuration, managing business logic, and tuning rules.
- Signal Transaction Scores customers can manage both Signal rule evaluations (now known as Signal Transaction Scores-powered rule evaluations) and Balance-only rule evaluations via a single Dashboard, consolidating business logic and performance tuning for all Plaid Signal products in a single no-code interface.
- All new Signal Transaction Scores customers who do not specify a `ruleset_key` when calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will now be assumed to be using the `default` ruleset, rather than no ruleset. Customers who do not want to use a ruleset can contact their Account Manager to opt out. This change does **not** impact existing Signal Transaction Scores customers; all existing customers will automatically be opted out of this behavior.

These changes will automatically take effect for all new customers who are enabled in Production for Balance or Signal Transaction Scores after October 15, 2025. A migration path will be available soon for existing customers. If you are an existing customer and would like to migrate early, contact your Account Manager for information.

For Plaid Check Consumer Report:

- Released the Plaid LendScore. The LendScore is a score from 1-99 that indicates an end user's creditworthiness based on data from their Plaid-linked accounts. LendScore is currently in closed beta. To request access, existing customers should contact their Plaid Account Manager; new customers should contact Sales.
- Released Cash Flow Insights, providing aggregated transaction data across all permissioned accounts. This data can be used for credit decisioning and forms the basis of the Plaid LendScore. Cash Flow Insights is currently in closed beta. To request access, existing customers should contact their Plaid Account Manager; new customers should contact Sales.

For Transfer:

- Released Transfer for Platforms (formerly Platform Payments) to wider beta for reseller partners (i.e. platforms that incorporate Plaid to facilitate pay-by-bank for their customers), with a set of redesigned endpoints for Platform use cases. If you are interested in using Transfer as a Platform, contact your Plaid Account Manager.

##### October 13, 2025

For Link:

- Released [Android SDK](https://github.com/plaid/plaid-link-android) v5.4.0, resolving an issue in which the `selection` event was missing in `LinkEvent.EventMetadata`, resolving an issue where `metadataJson` in `LinkEventMetadata` could be an empty string instead of `{}`, and improving internal logging and debugging to help diagnose customer issues more effectively.
- Released [React Native SDK](https://github.com/plaid/react-native-plaid-link-sdk) v12.6.0, updating Android to v5.4.0 and iOS to v6.4.1.

For Auth:

- Enabled known Instant Match incompatible depository accounts within the single account Account Select pane. While these accounts cannot be connected via Instant Match, end users can now select them in order to verify them via other methods, such as micro-deposits or Database Auth.

For Plaid Check Consumer Report:

- Base Reports and Income Insights Reports are now available to customers on Pay-as-you-go or Growth plans.

For Partners:

- Partners can now file tickets on behalf of their end customers directly from the Partner Dashboard without impersonating the end customer. This feature is not yet available for missing/faulty data categories.

##### October 1, 2025

For Link:

- Released [iOS SDK](https://github.com/plaid/plaid-link-ios) v6.4.1, resolving an issue in which the `selection` event was missing in `LinkEvent.EventMetadata` and improving internal logging and debugging to help diagnose customer issues more effectively.

For Auth and Identity:

- Released UI improvements to the [Account Verification Dashboard](https://dashboard.plaid.com/account-verification), including labels and session counts on the Overview page to show at a glance which flows each Link customization is enabled for and which customizations are actively used.

For Credit products:

- Released UI improvements to the Credit Dashboard (beta), including improved search, sort, and filter capabilities on the session overview page; breadcrumb navigation within the Dashboard; labeling of pending versus posted transactions; and color-coding / formatting to more clearly distinguish positive from negative transaction amounts.

For Layer:

- Shipped UI changes to the Layer flow to increase conversion.

For Transfer:

- Changed the default and maximum `count` values for [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) from 25 and 100 to 100 and 500, respectively.

##### September 24, 2025

Due to a change in Chase's permissioning UI, the `USER_ACCOUNT_REVOKED` webhook will no longer be fired for Chase Items, as Chase no longer allows account-level permissions revocation. Instead, the `USER_PERMISSION_REVOKED` webhook will be fired. PNC Items will continue to support the `USER_ACCOUNT_REVOKED` webhook.

For Plaid Check Consumer Report:

- Updated designs of certain panes in the Link flow. These are small visual updates that have resulted in improvements to both Link conversion and Plaid Passport enrollment.

For Transfer:

- Deprecated the `funds_settlement_date` field. All customers using Plaid Ledgers should use the `expected_funds_available_date` instead.

##### September 16, 2025

For Link:

- Released [React Native SDK](https://github.com/plaid/react-native-plaid-link-sdk) v12.5.1, containing bug fixes for Android build errors.
- Released [React Native SDK](https://github.com/plaid/react-native-plaid-link-sdk) v12.5.2, resolving iOS startup crash on React Native version 0.81+ introduced in v12.5.1. Also adds `metadataJson` key to event data to allow for all keys to be camelCase and updates Android SDK to 5.3.3.
- Released [React Native SDK](https://github.com/plaid/react-native-plaid-link-sdk) v12.5.3, upgrading Android SDK to v5.3.4.
- Released [Android SDK v5.3.3](https://github.com/plaid/plaid-link-android), containing the following changes: upgrade dependency `com.google.code.gson:gson` to 2.9.1, upgrade dependency `com.squareup.okhttp3:logging-interceptor` to 4.9.2.
- Released [Android SDK v5.3.4](https://github.com/plaid/plaid-link-android), fixing a retrofit reinitialization bug.

For Auth:

- TAN expirations at PNC have been postponed until January 2026. While PNC Items will still expire after one year if consent is not renewed, the TAN will continue to be valid until at least January 2026.

For Plaid Check Consumer Report:

- For Cashflow Monitoring, added the ability to subscribe to updates on a per-Item level.
- For Cashflow Monitoring, removed the fields `baseline_count` and `baseline_amount` in response to customer feedback. These fields will be omitted or returned with `null` values going forward.

For Identity Verification:

- Improvements to Trust Index Scores (US verifications only):
  - Added the Trust Index Thresholds feature, which allows customers to set a minimum Trust Index score in addition to attribute-based Risk Rules.
  - Improved and simplified the UI for displaying Trust Index Scores. Identity Verification will now display the primary drivers for the score as bullet points.
  - Improved the algorithm for calculating Trust Index Scores, resulting in substantially increased accuracy.

##### August 28, 2025

For Investments Move:

- Added a new error code, `MANUAL_VERIFICATION_REQUIRED`, when [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) is called on an Item that requires an interactive Link session for verification. To resolve this error, send the Item through update mode.

##### August 27, 2025

For the Dashboard:

- The "Team Management" permission is no longer required to access the Link Analytics page. Any verified user can now access this page.
- The "Link Customization Write Access" permission is now required to access Link Customization / Template editor pages instead of the "Team Management" permission.

For Link:

- Released [Plaid Android SDK 5.3.2](https://github.com/plaid/plaid-link-android/), including an upgrade to version 3.25.5 of the `com.google.protobuf:protobuf-kotlin-lite` library.

For Auth:

- Improved the Link UI on the "Are you sure?" screen displayed when an end user attempts to close Link without connecting an account. The new screen more prominently displays the option to link an account via non-credential based flows (if enabled), resulting in increased Link conversion.

For Plaid Check Consumer Report:

- Added support for Consumer Report in the [Embedded Institution Search](/docs/link/embedded-institution-search/) Link UI. This Link configuration is recommended when using Consumer Report alongside Auth as part of a single Link flow that supports bank account linking for both underwriting and repayment.

For Layer:

- Added [Extended Autofill](/docs/layer/#extended-autofill), which expands the number of users who can have their personal information prefilled with Layer, even if their phone number is not Layer eligible. To learn more and add Extended Autofill to your existing integration, see [Extended Autofill](/docs/layer/#extended-autofill).

For Transfer:

- Added monthly Transfer reconciliation report emails. This feature is currently opt-in; to request access, contact your Plaid Account Manager.

##### August 17, 2025

For the Plaid Dashboard:

- Added audit logs of Dashboard actions. Audit logs are available to admins via the Dashboard under Settings -> Audit Logs. Currently, core Dashboard actions are logged; logging for actions in product-specific Dashboards is not yet available. Audit logs are available only to customers on Premium Platform Support packages. To learn more about upgrading to a Premium Platform Support package, contact your Account Manager.

For all products:

- For [`/item/remove`](/docs/api/items/#itemremove), added optional `reason_code` and `reason_note` fields. Using these fields when removing an Item can help Plaid identify fraud and bad actors in the Plaid Network, improving anti-fraud data for all Plaid customers.

For Plaid Link:

- Released [Plaid Android SDK 5.3.1](https://github.com/plaid/plaid-link-android/), fixing a bug in which screens did not properly resize for the keyboard on Android 15+.

For Plaid Check Consumer Report:

- Added the ability to generate Home Lending Report and Employment Refresh reports suitable for providing to GSEs. To generate these reports, use the `base_report.gse_options` field when calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) or [`/link/token/create`](/docs/api/link/#linktokencreate).
- Added the ability to create and manage OAuth tokens to securely share consumer report data with Plaid partners Experian, Fannie Mae, and Freddie Mac. For more details, see [Consumer Report OAuth data sharing](/docs/check/oauth-sharing/).

For Identity Verification:

- In the Document Verification `analysis` object, added the `aamva_verification` object to report results from the American Association of Motor Vehicles Administrators (AAMVA) Drivers License Data Verification service. This object is currently in beta. To receive AAMVA data, you must enable the corresponding feature in your Identity Verification template under the "DMV/Secretary of State Validation" header in the Workflow tab of the Template Editor. This object will only be populated for US drivers licenses in [participating states](https://www.aamva.org/it-systems-participation-map?id=594).

For Investments:

- Removed the `sedol` field from responses, as this field was only available for stocks that trade on London Stock Exchange and to customers that had the appropriate licenses. The `sedol` field will now be `null` for all securities.
- Added `subtype` identifiers to the `security` object to specify the type of a security in more detail. New subtype values are: `asset backed security`, `bill`, `bond`, `bond with warrants`, `cash`, `cash management bill`, `common stock`, `convertible bond`, `convertible equity`, `cryptocurrency`, `depositary receipt`, `depositary receipt on debt`, `etf`, `float rating note`, `fund of funds`, `hedge fund`, `limited partnership unit`, `medium term note`, `money market debt`, `mortgage backed security`, `municipal bond`, `mutual fund`, `note`, `option`, `other`, `preferred convertible`, `preferred equity`, `private equity fund`, `real estate investment trust`, `structured equity product`, `treasury inflation protected securities`, `unit`, `warrant`.

For Transfer:

- In the [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget) response `rfp` object, added `max_amount` and `iso_currency_code` fields.
- Added [`/transfer/ledger/event/list`](/docs/api/products/transfer/ledger/#transferledgereventlist) endpoint.

##### August 6, 2025

For Plaid Link:

- Released a visual refresh of the Link UI to match the Plaid brand refresh. This update does not change the functionality of Link.
- Released [Plaid iOS SDK 6.4.0](https://github.com/plaid/plaid-link-ios), adding `issueDescription` and `issueDetectedAt` to `EventMetadata`.
- Released [Plaid Android SDK 5.3.0](https://github.com/plaid/plaid-link-ios), adding `issueDescription` and `issueDetectedAt` to `EventMetadata`, the event `LAYER_AUTOFILL_NOT_AVAILABLE`, and supporting the `dateOfBirth` parameter for Layer.
- Released [Plaid React Native SDK 12.4.0](https://github.com/plaid/react-native-plaid-link-sdk). This release updates the Android SDK to 5.3.0 and the iOS SDK to 6.4.0. It also includes updates to the example application to make it more clear how to use Layer and Sync Financekit.
- Released [Plaid React Link SDK 4.1.1](https://github.com/plaid/react-plaid-link), patching a build issue introduced in version 4.1.0.

For Plaid Check Consumer Report:

- Improved income forecasting calculations in the Income Insights report:
  - Tax refunds and inactive income streams will no longer contribute to forecasted income.
  - Income forecasts will now use all available income history (up to 5 years) rather than a 90 day maximum.
- During the week of August 18, 2025, Data Transparency Messaging (DTM) will be disabled for Consumer Report until further notice. No action is needed on the part of impacted customers. For Link sessions that include both Plaid Check products and Plaid Inc. products in a single Link flow, DTM will still be used for Plaid Inc. products if it is otherwise enabled for the Link flow.

For Signal:

- Beginning September 1, 2025, when [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) or [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport) is called with an unknown `client_transaction_id` (one that was never submitted to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate)), these endpoints will return an error instead of a success response. This change better reflects that these endpoints should only be used to report returns or decisions on previously-evaluated transactions.

For Transfer:

- Added more granular Transfer-specific permissions controls to the Dashboard. For more details, see [Transfer Dashboard permissions](/docs/transfer/dashboard/#dashboard-permissions). To manage your team's Transfer permissions, use the [Dashboard -> Team Settings -> Members page](https://dashboard.plaid.com/settings/team/members).

##### August 4, 2025

For Plaid Link:

- Released [Plaid React Link SDK 4.1](https://github.com/plaid/react-plaid-link), with support for passing date of birth for Layer autofill.

##### July 30, 2025

For Plaid Link:

- Released [Plaid iOS SDK 6.3.2](https://github.com/plaid/plaid-link-ios), with fixes to XCFramework signing issue.
- Released [Plaid iOS SDK 6.3.1](https://github.com/plaid/plaid-link-ios), with updated Layer Submit API, new Layer event `LAYER_AUTOFILL_NOT_AVAILABLE`, and exposing Finance Kit sync simulated behavior to Objective-C (React Native).
- Released [Plaid React Native SDK 12.3.2](https://github.com/plaid/react-native-plaid-link-sdk), updating the React Native SDK to use Plaid iOS SDK 6.3.2.
- Released [Plaid React Native SDK 12.3.1](https://github.com/plaid/react-native-plaid-link-sdk), updating the React Native SDK to use Plaid iOS SDK 6.3.1.

##### July 25, 2025

For Plaid Link:

- Released [Plaid Android SDK 5.2](https://github.com/plaid/plaid-link-android/), with new event names and a new event metadata field, improved behavior for `destroy()`, a smaller SDK size, and removal of unused dependencies. For more details, see the [release notes](https://github.com/plaid/plaid-link-android/releases/tag/v5.2.0)
- Released [Plaid React Native SDK 12.3.0](https://github.com/plaid/react-native-plaid-link-sdk), updating the React Native SDK to use Plaid Android SDK 5.2 and Plaid iOS SDK 6.3.
- Began rolling out a progress bar in the Link UI to all customers. During testing, Link sessions with the progress bar displayed showed a statistically significant increase in Link conversion.

##### July 24, 2025

For the Plaid Dashboard:

- Improved the user experience for OAuth registration and onboarding:
  - OAuth registration status is now displayed more prominently within the Dashboard
  - Added estimated timelines for institution enablement
  - Accelerated the OAuth registration process, resulting in shorter waits for OAuth enablement
  - Added more complete list of institutions that use OAuth and their enablement statuses
  - Added detailed and actionable error messages if an error occurred during registration
  - Customers on Growth or Custom plans will automatically be enabled for access to Charles Schwab; customers on Pay-as-you-go plans will still need to explicitly request Schwab access
- Added the ability to export the list of team members to CSV.
- Added the ability for customers with Premium Platform Support Packages to self-serve enable SCIM in the Dashboard without contacting Plaid Support or their Account Manager.
- Removed Data Quality related Known Issues the Dashboard, as these issues were often described in vague terms, causing customers to think they were encountering a known issue when they were not. Removing these from the Known Issues section of the Dashboard helps Plaid more accurately identify data quality issues and reduces customer confusion caused by following the wrong Known Issue. Connectivity-related Known Issues are still reported in the Dashboard.
- When users attempt a Dashboard action they don't have permissions for, they will now be shown a list of admins whom they can contact to request permission.

For Plaid Link:

- Released [React Native Plaid Link SDK 12.2.1](https://github.com/plaid/react-native-plaid-link-sdk), with bug fix for the error message "PLKEnvironmentDevelopment not defined".
- Released [Plaid iOS SDK 6.3](https://github.com/plaid/plaid-link-ios/), with new event names and a new event metadata field, enhancements for the SwiftUI API, and improved FinanceKit testing capabilities. For more details, see the [release notes](https://github.com/plaid/plaid-link-ios/releases/tag/6.3.0).

For Plaid Check Consumer Reports:

- Added support for Home Lending Report and Employment Refresh for mortgage lenders via [`/cra/check_report/verification/get`](/docs/api/products/check/#cracheck_reportverificationget). This feature is in Early Availability; contact Sales or your Account Manager to learn more.
- Improved the accuracy of insufficient funds (NSF) transaction calculations, leading to fewer transactions being reported under the `nsf_overdraft_transactions_count` fields.
- For Income Insights, improved the accuracy of income categorization, forecasted income, and historical income.
- Added institution details (`institution_name` and `institution_id`) to [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) response.
- Added `require_identity` boolean to options object in [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) and [`/link/token/create`](/docs/api/link/#linktokencreate) requests. If set to `true` and user identity data is not available, CRA report creation will fail, and you will not be charged.
- Updated and made consistent maximum values `days_required` and `days_requested` when generating a CRA report across multiple endpoints. The maximum value for `days_requested` is now 731 and the maximum value for `days_required` is 184.
- For Partner Insights, added `client_report_id` to partner insights response.

For Investments:

- Improvements and bug fixes to holdings valuations, such as better reflecting the value of unvested shares.

For Payment Initiation:

- Added `error` field to responses of [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) and [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute) containing error details when the payment fails.
- Added support for cursors to [`/payment_initiation/recipient/list`](/docs/api/products/payment-initiation/#payment_initiationrecipientlist).
- Added error field to `WALLET_TRANSACTION_STATUS_UPDATE` webhook and to responses of [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) and [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist).

For Transactions:

- Added [`/sandbox/transactions/create`](/docs/api/sandbox/#sandboxtransactionscreate) endpoint, which can be used to add custom transactions to the `user_transactions_dynamic` Sandbox test user.
- Added the `user_ewa_user`, `user_yuppie`, and `user_small_business` Sandbox test users for persona-based testing of realistic Transactions data.
- Added `account_numbers` field to `counterparty` object.

For Transfer:

- Added `wire_return_fee` to transfer and transfer event objects.

##### June 12, 2025

- Released [Plaid Protect](/docs/protect/), Plaid's new anti-fraud solution.
- Released [Premium Link Analytics](/docs/link/measuring-conversion/#premium-link-analytics), a new Dashboard feature with advanced analytics insights on Link usage and conversion. Premium Link Analytics is available at no extra charge to all customers with Premium Support Packages. To learn more about upgrading to a Premium Support Package, contact your Account Manager.
- Launched Plaid's AI Toolkit, including the Dashboard MCP Server, Sandbox MCP Server, LLM-friendly documentation, and vibe coding guides. For more details, see the [Resources page](https://plaid.com/docs/resources/#ai-developer-toolkit).
- Released version 5.1.1 of the [Plaid Android Link SDK](https://github.com/plaid/plaid-link-android) and version 6.2.1 of the [iOS Link SDK](https://github.com/plaid/plaid-link-ios). These updates both contain the ability to detect usage of [Plaid Link for Flutter](https://github.com/jorgefspereira/plaid_flutter), allowing Plaid to better troubleshoot integration issues by distinguishing between Flutter and non-Flutter usage.

##### June 9, 2025

For Link:

- Released version 5.1.0 of the [Plaid Android Link SDK](https://github.com/plaid/plaid-link-android). Improvements include a bug fix for edge-to-edge layout overlap issue on Android 15+, 20% smaller SDK size, and removal of the `org.bouncycastle:bcpkix-jdk15to18` dependency.
- Began rolling out a new pane added to the end of the Link flows for Assets and Bank Income, prompting end users to enroll in Plaid Passport for faster credit application flows in the future. This pane occurs after the end user has successfully completed Link and therefore has no impact on Link conversion. This pane will be released to all eligible customers by the end of August.

For Plaid Check Consumer Reports:

- Introduced a rate limit of 100 requests per minute for all Plaid Check endpoints.

For Transfer:

- Plaid Transfer now supports Instant Payments (pay-ins) via request-for-payment (RfP) on the RTP network.

##### June 3, 2025

- Announced the availability of the [Plaid Dashboard MCP server](/docs/resources/mcp/). This service allows you to access Dashboard-based services such as Link analytics data, API usage, and the Item debugger via an LLM agent. More tools and the ability to use additional platforms will be added soon.
- Released Plaid Link [iOS SDK 6.2.0](https://github.com/plaid/plaid-link-ios/releases/tag/6.2.0). This version features a 33% smaller package size, a `showGradientBackground` option for customizing Link appearance, and a new `onLoad` callback from `Plaid.create` to help detect when Plaid is ready to present.

For Auth:

- Released [Account Verification Analytics](https://dashboard.plaid.com/account-verification/analytics), Dashboard-based reports of Auth and Identity Match activity showing usage levels and conversion metrics by Auth method.

For Plaid Check Consumer Reports:

- For [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), deprecated `report.items.accounts.account_insights` in favor of the [`attributes` field](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes), which aggregates insights across accounts.
- Announced the availability of a new model version (v4) for their scores that are provided in Partner Insights. Starting on June 2, 2025, customers using Partner Insights can choose to use the latest version when calculating the scores by specifying it in the `prism_versions` object when calling [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget). If no version is specified, Plaid will default to the current version (v3) available now. Starting in November 2025, customers will be required to specify a `prism_version` when calling Partner Insights endpoints.

For Transfer:

- Added [automatic Sandbox simulations](https://plaid.com/docs/transfer/sandbox/#automatic-sandbox-simulations), allowing you to create test transactions in Sandbox that will automatically move through the transfer lifecycle, without requiring you to manually progress them to the next state.

##### May 20, 2025

For Enrich:

- Removed the legacy `category_id` and `category` fields for all customers enabled for the Transactions product on or after May 5, 2025. Customers should use the `personal_finance_category` field instead.

For Transactions:

- Removed the legacy `category_id` and `category` fields for all customers enabled for the Transactions product on or after May 5, 2025. Customers should use the `personal_finance_category` field instead.

##### May 14, 2025

For Auth:

- Announced PNC's new TAN expiration policy. All PNC Items with TANs will require the user to have authorized or re-authorized consent within the past 12 months. If the user has not done so, the TAN will stop working and the Item will enter an `ITEM_LOGIN_REQUIRED` state. For more details, see [PNC TAN expiration](/docs/auth/#pnc-tan-expiration).

For Layer:

- Added `android_package_name` to [`/session/token/create`](/docs/api/products/layer/#sessiontokencreate).

For Plaid Check Consumer Reports:

- Updated `average_inflow_amount` field to be positive.

For Transfer:

- Announced PNC's new TAN expiration policy. All PNC Items with TANs will require the user to have authorized or re-authorized consent within the past 12 months. If the user has not done so, the TAN will stop working and the Item will enter an `ITEM_LOGIN_REQUIRED` state. For more details, see [PNC TAN expiration](/docs/auth/#pnc-tan-expiration).

##### May 5, 2025

- Announced new `INSTITUTION_RATE_LIMIT` error and behavior for when an institution is experiencing excessive API traffic to realtime endpoints. This error will be rolled out to all customers by the end of May 2025.

For Transfer:

- Added `expected_funds_available_date` to the Transfer object.

##### April 24, 2025

- Released a new [migrations page](https://dashboard.plaid.com/activity/status/migrations) on the Dashboard, showing all institutions with planned or ongoing migrations to OAuth. Drilling in to each institution, you can now view the migration timeline and impact to existing Items.
- Added support for Austria and Finland. For details of supported institutions and products, see [European bank coverage](https://plaid.com/docs/institutions/europe).
- Added `USER_PERMISSION_REVOKED` webhook support for American Express.

For Layer:

- Added the ability to configure Layer sessions in the Dashboard.
- In [`/session/token/create`](/docs/api/products/layer/#sessiontokencreate), moved the `user_id` field into the `user` object.
- Released React Native SDK 12.1.1, with a new `destroy()` method to support Layer.

For Link:

- Released React Native SDK 12.1.1, with a new `destroy()` method to support Plaid Layer.
- Released React SDK 4.0.1, with support for React 19.
- Webview-based integration modes are no longer allowed for new customers. All new customers using mobile webview-based apps should use the Hosted Link integration mode if they can't use Plaid's SDKs. Existing customers using webview integration modes are not currently impacted.
- Announced a combined phone number and consent pane for the [returning user experience](https://plaid.com/docs/link/returning-user/), to be rolled out to all customers in May. This change increases conversion by reducing the number of screens the end user has to complete. It does not require any integration work on your part to adopt. Both the `CONSENT` and `SUBMIT_PHONE` `TRANSITION_VIEW` Link analytics events will fire when this screen is reached.

For Plaid Check Consumer Reports:

- Expanded support for Silent Network Authentication (SNA), a secure, lower friction way to verify a user’s phone number using their mobile network carrier. Once eligible users submit their phone number, they will see a loading pane in Link in lieu of the OTP pane. If SNA cannot verify a user’s phone number, then it will fall back to OTP for verification. SNA is available for Plaid Check Consumer Reports on both iOS and Android and is compatible with T-Mobile, Verizon, and AT&T. SNA will be rolled out to all eligible Consumer Report sessions by the end of May. No work is required on your part to support SNA. SNA is supported on Android SDK versions >=5.6.0 and iOS SDK versions >=4.7.9. To get the latest performance improvements for SNA, it is strongly recommended that all Plaid Check customers using Plaid's mobile SDKs upgrade to iOS SDK >=6.0.1 (or React Native iOS >=12.0.2).
- Date of birth is now required within the consumer report user identity object for [`/user/create`](/docs/api/users/#usercreate) and [`/user/update`](/docs/api/users/#userupdate).
- Added `warnings` to responses for [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), [`/cra/check_report/network_insights/get`](/docs/api/products/check/#cracheck_reportnetwork_insightsget) and [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget). `warnings` is populated when Plaid successfully generated a report, but encountered an error when extracting data from one or more Items.
- A user's phone number can now be pre-filled in Link using data from the `consumer_report_user_identity` object if the phone number is not pre-filled via [`/link/token/create`](/docs/api/link/#linktokencreate).
- A new toggle on the last pane of the Link flow enables users to opt-in to a faster way to share their consumer reports in the future.
- Updated the list of [Cash flow updates webhooks (beta)](https://plaid.com/docs/api/products/check/).

For Signal:

- Added `triggered_rule_details` to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) response.

For Transactions:

- Updated the response schema for [`/processor/transactions/sync`](/docs/api/processor-partners/#processortransactionssync) endpoint to allow `null` account objects. This change only impacts Transactions processor partners.

##### March 31, 2025

For Identity Verification:

- Announced the rollout of Age Estimation and Biometric Deduplication features. On May 1, 2025, Plaid will backfill historical ID document and liveness data and enable these features for all customers using IDV templates with ID document and/or liveness enabled, free of charge. Any Plaid customer who wants to opt out of this functionality should contact their Account Manager.

##### March 25, 2025

- Added a "copy link" button to the Team switcher page in the Dashboard, allowing Dashboard users to access a persistent link to their team.

For Assets:

- To minimize data access, Assets no longer requests the Investments data scope by default. To use the Investments add-on, `investments` must now be specified in the `optional_products` array when initializing Link.

For Layer:

- Added the [`LAYER_AUTHENTICATION_PASSED` webhook](https://plaid.com/docs/api/products/layer/#layer_authentication_passed).

For Investments:

- Added Vanguard access via OAuth. Access to Vanguard is no longer restricted during peak market hours.

For Plaid Check Consumer Reports:

- Broke up the existing `DATA_UNAVAILABLE` error into more specific error codes when possible, introducing the new codes `INSTITUTION_TRANSACTION_HISTORY_NOT_SUPPORTED`, `INSUFFICIENT_TRANSACTION_DATA`, and `DATA_QUALITY_CHECK_FAILED`. Also replaced `INVALID_FIELD` with `NETWORK_CONSENT_REQUIRED` for the use case where the end user has not given consent to share network data. For details on these new codes, as well as improved documentation of existing codes, see [Check Report Errors](https://plaid.com/docs/errors/check-report/).
- Improved the sensitivity of the `CHECK_REPORT_FAILED` webhook at alerting to partial failures. Previously, certain types of partial Check Report failures would not trigger the webhook.
- Added `successful_products` and `failed_products` arrays to the `CHECK_REPORT_READY` webhook to provide more granular detail on partial Check Report failures.
- Plaid can now pre-fill a user’s phone number in Link using information from the `consumer_report_user_identity` object if the phone number is not pre-filled via [`/link/token/create`](/docs/api/link/#linktokencreate). Pre-filling phone numbers can help boost conversion while reducing steps for users. Support for this feature is currently available in Production and will be added in Sandbox in the coming months.
- Increased the speed of report generation by approximately 25%.
- By the end of March, to improve security, Plaid will send all end users who complete the Plaid Check Link flow flow a confirmation email, similar to the change to Plaid Link announced in the [January 30 changelog entry](/docs/changelog/#january-30-2025).

For Signal:

- Added realtime backtesting support to the Signal Dashboard.
- Added a view-only permission to the Signal Dashboard.

##### March 17, 2025

For Link:

- Released iOS SDK 6.1.0, with improvements to Remember Me Silent Network Authentication.
- Released React Native SDK 12.0.3, which updates to iOS SDK 6.1.0.

For Auth and Transfer:

- Announced the impending migration of US Bank to tokenized account numbers (TANs), scheduled to occur May 2025.

For Identity:

- Added support for API-based detection of pass / fail results from Identity Match sessions when Identity Match rules are configured via the [Dashboard](https://dashboard.plaid.com/account-verification) via the `IDENTITY_MATCH_PASSED` and `IDENTITY_MATCH_FAILED` Link events.

For Layer:

- Publicly documented [`/session/token/create`](/docs/api/products/layer/#sessiontokencreate) as the preferred endpoint for creating a Layer session, rather than [`/link/token/create`](/docs/api/link/#linktokencreate), and added `redirect_uri` support to [`/session/token/create`](/docs/api/products/layer/#sessiontokencreate).

For Plaid Check Consumer Reports:

- Added support for `client_report_id` to Income Insights Reports. Moved `client_report_id` field to the top-level [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) request, rather than being inside the `base_report` request object, to reflect that the `client_report_id` is no longer Base Report-specific. Usage of the `client_report_id` within the `base_report` object is deprecated.
- Changed the `score` field from a float to an integer in the API spec, to reflect that it will always be an integer.

For Statements:

- Added `posted_date` field to [`/statements/list`](/docs/api/products/statements/#statementslist) response.

##### March 1, 2025

For Auth:

- Added the ability to configure micro-deposit and database verification-based Auth flows via the [Dashboard](https://dashboard.plaid.com/account-verification).
- Marked Database Match and Database Insights (beta) as deprecated in favor of the new Database Auth feature, which provides functionality similar to Database Insights. These features are still available, but customers should plan to migrate to Database Auth when possible.

For Identity:

- Added the ability to configure Identity Match rules via the [Dashboard](https://dashboard.plaid.com/account-verification) when using Identity and Auth as the only products.

For Plaid Check Consumer Report:

- Updated behavior when calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) to improve consistency. Previously, no reports would be eagerly generated upon calling this endpoint. Now, the base report will always be eagerly generated, along with products specified in the `products` array.
- When calling [`/link/token/create`](/docs/api/link/#linktokencreate), it is now possible to include Plaid Check products within the `optional_products`, `required_if_supported_products`, or `additional_consented_products` request fields when Plaid Inc. products are included in the `products` field, or vice-versa.
- Identity data is now provided in the Base Report only. Plaid will no longer return identity data in the Income and Partner Insights products, and the fields will be removed from API responses by June 30, 2025.
- Removed the Cashflow Updates `INSIGHTS_UPDATED` `reason` code. Instead, a separate webhook will be fired for each reason.
- The `score` field in [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget) is now nullable.

For Transfer:

- Added `intent_id` to the transfer event object returned by [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync), for transfers created via Transfer UI. This will currently only be populated for RfP transfers.

##### February 11, 2025

For Link:

- Released React Native SDK 12.0.3, which contains LinkKit 6.0.4, which fixes an issue where some sessions experienced delays in receiving the `LAYER_READY` event or did not receive it at all, and fixes an issue with XCFramework.

For Auth:

- Improved Sandbox behavior to better reflect institutions that use Tokenized Account Numbers (TANs). In Sandbox, Chase and PNC will now return `is_tokenized_account_number: true` and have populated `persistent_account_id` fields.

For Plaid Check Consumer Report:

- `consumer_report_user_identity.date_of_birth` is now a required field when creating or updating a user token with a `consumer_report_identity` field.
- Completed the renaming of several fields. The `longest_gap_between_transactions`, `average_inflow_amount`, and `average_outflow_amount` fields have been removed and replaced with pluralized versions `longest_gaps_between_transactions`, `average_inflow_amounts`, and `average_outflow_amounts` fields to better reflect the fact that they are arrays.

For Signal:

- Announced a behavior change for Items that were manually linked (e.g. using Same Day Micro-deposits or Database Insights). Beginning February 18, any Item that was manually linked but could not be given a Signal score and attributes by [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will no longer return a `PRODUCT_NOT_SUPPORTED` error. Instead, the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) call will succeed. This change allows customers to evaluate these Items using Signal Rules that do not require a Signal score and attributes. Billing will only be triggered for the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) call if either a score is returned or a Signal Rule is evaluated.

For Virtual Accounts:

- Added `originating_fund_source` to [`/wallet/transaction/execute`](/docs/api/products/virtual-accounts/#wallettransactionexecute) request. This field is required by local regulation for certain use cases (e.g. money remittance) to send payouts to recipients in the EU and UK.

##### February 3, 2025

For Link:

- Released iOS SDK 6.0.4, which fixes an issue where some sessions experienced delays in receiving the `LAYER_READY` event or did not receive it at all, and fixes an issue with XCFramework.

##### January 30, 2025

Notified all customers that they must sign an addendum to their agreement with Plaid to retain PNC data access after February 28th, 2025. If you do not know how to sign this addendum, or are not sure if you have signed it, contact your Plaid Account Manager or Plaid Support.

For Link:

- Released React Native SDK 12.0.1, which includes iOS SDK 6.0.2 and resolves issues where the SDK was not working with React Native 0.76 or later.
- Released React Native SDK 12.0.2, which results an issue where the `USE_FRAMEWORKS` preprocessor check was failing for projects using `use_frameworks! :linkage => :static`.
- To improve anti-fraud measures, began more broadly sending email confirmations to end users when they have linked an account to Plaid, allowing them to report unrecognized Link activity. Previously, Plaid would send these notifications only if the session was flagged as elevated risk.

For Payment Initiation (UK/EU):

- Added the `end_to_end_id` field to the payment object, providing a unique identifier for tracking and reconciliation purposes.

For Transfer:

- For [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget), added an `institution_supported_networks.rfp.debit` field to indicate an institution's support for debit request for payment (RfP).
- Added `funds_available` sweep status and event type.

##### January 17, 2025

SSO is now available, at no additional charge, to all customers except those on Pay-as-you-go plans. You can enable SSO via the [Dashboard > Team Settings > SSO](https://dashboard.plaid.com/settings/team/sso).

For Link:

- Released iOS SDK 6.0.1, with enhanced support for Silent Network Authentication in the Remember Me flow.
- Released iOS SDK 6.0.2, with small improvements to the UI of the Remember Me flow.
- Improved the UI for failed connection attempts at OAuth institutions to increase conversion. Rather than "Exit," the primary CTA on the Link error screen for this scenario is now a "Try Again" button that will directly re-start the institution's OAuth flow.
- Added phone number recycling detection to Link. The returning user experience will delete a user's Returning User profile if it detects that the phone number has been reassigned to a different customer since the user's last login, reducing the risk of fraud and account takeover.

For Plaid Check:

- Added `nsf_overdraft_transactions_count_30d/60d/90d` fields to `attributes` field of [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget).

For Payment Initiation:

- Added [`/sandbox/payment/simulate`](/docs/api/sandbox/#sandboxpaymentsimulate) endpoint to improve ease of testing in Sandbox.

For Signal:

- [Signal Rules](/docs/signal/signal-rules/) now supports creating rules based on the results of a [Database Insights](/docs/auth/coverage/database/) check.

For Transfer:

- Added Transfer Rules, allowing you to customize the logic used by the Transfer Risk Engine during the [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) call. You can configure your [Transfer Rules](https://dashboard.plaid.com/transfer/risk-profiles) via the Dashboard.
- To better represent errors from a variety of payments rails beyond just ACH, added `failure_code` and `description` fields to [`/transfer/sweep/get`](/docs/api/products/transfer/reading-transfers/#transfersweepget) and `transfer/sweep/list` endpoints. Developers should use these instead of the `ach_return_code` field.

##### January 2, 2025

For Link:

- For institutions such as PNC, Chase, and Charles Schwab that previously invalidated Items when a duplicate Item was added, prevented the old Item from becoming invalidated unless either Item is initialized with a Plaid Check product or the two Items have different sets of accounts associated. This is a delayed announcement of a change made in April 2024.
- Released React Native SDK 12.0.0. This incorporates iOS SDK 6.0.0 and Android SDK 5.0.0. Major updates include adding support for FinanceKit and AppleCard on iOS, removing the deprecated `PlaidLink` component and `openLink` function, and updating a number of Android libraries. For a full list of all changes in this release, see the [Release notes](https://github.com/plaid/react-native-plaid-link-sdk/releases/tag/v12.0.0).

##### December 23, 2024

For Link:

- Released Android SDK 5.0.0, containing numerous library and platform updates, including upgrading Kotlin to 1.9.25, and upgrading the compile version to SDK 35. This SDK version also adds the `AUTO_SUBMIT` event name and `INVALID_UPDATE_USERNAME` Item error, removing the `PROFILE_ELIGIBILITY_CHECK_ERROR` event name. For a full list of all libraries updated in this SDK release, see the [Release notes](https://github.com/plaid/plaid-link-android/releases/tag/v5.0.0).

##### December 17, 2024

For Auth:

- PNC now supports the `USER_ACCOUNT_REVOKED` and `USER_PERMISSION_REVOKED` webhooks, allowing you to reduce return risk. Upon receiving either of these webhooks, Auth customers should no longer make ACH transfers using the tokenized account numbers of the associated Items.

For Transactions:

- For the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) endpoint, added the ability to use an `account_id` filter.

##### December 13, 2024

For Plaid Check:

- Began the rollout of Plaid Passport, an update to the Plaid Check bank linking flow. Plaid Passport is an extension of the [returning user experience](/docs/link/returning-user/) specific to Plaid Check, and will enable conversion efficiencies across all financial institutions, improved Item health, and enhanced user-centric insights. You do not need to take any action to enable Plaid Passport.

##### December 10, 2024

For Link:

- Released [React Native SDK 11.13.3](https://github.com/plaid/react-native-plaid-link-sdk), fixing a regression introduced in version 11.13.1 that caused build failures for some customers.
- Completed the rollout of Passkey support for the returning user experience in Link to all eligible sessions. Passkey support increases security and provides a more streamlined experience for returning users. With Passkeys, end users on iOS devices who have previously logged in to a financial account through Link can opt in to using Face ID or Touch ID to authenticate to this account for subsequent Link sessions. To enable Passkey support, ensure you are using the Plaid iOS SDK 4.3.0 or later or the Plaid React Native SDK 10.2.0 or later. You do not need to take any other action to enable Passkey support.

##### December 4, 2024

- Added `institution_name` to the `Item` object schema. You no longer have to call [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) to translate the institution data in an Item to human-readable format.
- For [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook), added support for `USER_PERMISSION_REVOKED` and `USER_ACCOUNT_REVOKED` webhooks.

For Link:

- Increased the rollout of Passkey support for the returning user experience in Link. Passkey support increases security and provides a more streamlined experience for returning users. With Passkeys, end users on iOS devices who have previously logged in to a financial account through Link can opt in to using Face ID or Touch ID to authenticate to this account for subsequent Link sessions. Passkey support will be rolled out to all eligible sessions by the end of the year. To enable Passkey support, ensure you are using the Plaid iOS SDK 4.3.0 or later or the Plaid React Native SDK 10.2.0 or later. You do not need to take any other action to enable Passkey support.
- Added biometric authentication support for Citibank on Android (biometric authentication support for Citibank on iOS was added in June 2024). No SDK update is required.
- Released Plaid Link for React Native SDK 11.13.2, updating Android Compile Version from 31 to 34.

For Auth:

- Added `is_tokenized_account_number` field to [`/auth/get`](/docs/api/products/auth/#authget) to indicate whether an account number is tokenized. Tokenized account numbers may require special business logic to avoid ACH returns. For more details, see [Tokenized account numbers](/docs/auth/#tokenized-account-numbers).
- Added support for cash management account subtypes, rather than just checking and savings.
- Added `auth_method` enum to the `Item` object schema to indicate which verification method (e.g. Instant Auth, Same-Day Micro-deposits, Database Match...) was used to verify the Item.
- Added `SAME_DAY_MICRODEPOSIT_AUTHORIZED` and `INSTANT_MICRODEPOSIT_AUTHORIZED` transition view names to Link.

For Plaid Check Consumer Reports:

- Added the [`/sandbox/cra/cashflow_updates/update`](/docs/api/sandbox/#sandboxcracashflow_updatesupdate) endpoint to facilitate testing.
- Enhanced the model used for estimating gross income from net income, resulting in ~10% greater accuracy.
- Users with the Team Management permission can now download a copy of their submitted Plaid Check application form from the Dashboard under Settings > Compliance Center > Company Documents.

For Identity Verification:

- At the start of the Selfie Check flow, added a warning to users if insufficient ambient light is detected, prompting them to fix their lighting conditions before proceeding.

For Investments:

- For the `security` object, added the `fixed_income` object with the following details: `yield_rate`, `maturity_date`, `issue_date`, and `face_value` fields, to provide additional details on bonds and other fixed income securities.
- Added processor partner support for Investments with the new [`/processor/investments/transactions/get`](/docs/api/processor-partners/#processorinvestmentstransactionsget) and [`/processor/investments/holdings/get`](/docs/api/processor-partners/#processorinvestmentsholdingsget) endpoints.

For Transfer:

- Added support for cash management account subtypes, rather than just checking and savings.
- Added `auth_method` enum to the `Item` object schema to indicate which verification method (e.g. Instant Auth, Same-Day Micro-deposits, Database Match...) was used to verify the Item.
- Added `SAME_DAY_MICRODEPOSIT_AUTHORIZED` and `INSTANT_MICRODEPOSIT_AUTHORIZED` transition view names to Link.
- Added `webhook` field to requests for [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate), [`/sandbox/transfer/refund/simulate`](/docs/api/sandbox/#sandboxtransferrefundsimulate), [`/sandbox/transfer/ledger/simulate_available`](/docs/api/sandbox/#sandboxtransferledgersimulate_available) and [`/sandbox/transfer/sweep/simulate`](/docs/api/sandbox/#sandboxtransfersweepsimulate).
- Added `funds_available` support for [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate).

##### November 14, 2024

- Released [`/consent/events/get`](/docs/api/consent/#consenteventsget) to help customers obtain consent logs for auditing purposes.
- Added consent related-details to the `item` object in [`/item/get`](/docs/api/items/#itemget), including `consented_data_scopes`, `consented_use_cases`, and `consent_expiration_time`, to help customers manage consent records and consent expiration.

For Link:

- Released Plaid Link for iOS SDK 6.0.0, introducing FinanceKit and support for Apple Card integration.
- Released Plaid Link for React Native SDK 11.13.1, fixing a compiler error when using the New Architecture with Expo.

For Assets:

- Added a hard limit of 15,000 transactions per Asset Report, to avoid returning truncated Asset Reports. An Asset Report with more than 15,000 transactions will now trigger the error `ASSET_REPORT_GENERATION_FAILED` with an error message "asset report is too large to be generated, try again with a shorter date range".

##### November 7, 2024

For Link:

- Released Plaid Link for React SDK 3.6.1, fixing an issue that can occur when unmounting the `usePlaidLink` hook before the underlying script tag is loaded.
- Customers can now self-enable for [Hosted Link](/docs/link/hosted-link/) without contacting their Account Manager. To enable Hosted Link, provide a `hosted_link` object in the [`/link/token/create`](/docs/api/link/#linktokencreate) call. The object can be empty or can include configuration parameters for Hosted Link. This change does not apply to Link Delivery, which still requires Account Manager enablement.

##### October 29, 2024

For Link:

- Added support for Vietnamese and Hindi.
- Added `AUTO_SUBMIT_PHONE` event to capture scenarios where the user's phone number is supplied in the [`/link/token/create`](/docs/api/link/#linktokencreate) call and the user has previously enrolled in the returning user experience, allowing Plaid to send an OTP code without prompting.
- Improved conversion through improved error messaging for incorrect username / password at non-OAuth institutions. Previously, users would be taken to an error pane and required to press the back button to retry their credentials; the new UI provides inline error messages if the user's credentials are incorrect.
- Fixed bug in which users could enter an error state by clicking a button while the in-progress spinner was active on a different button.
- Minor improvements to consistency in font weight and styling.

For Consumer Report by Plaid Check:

- For the endpoint [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), added the `total_inflow_amount` and `total_outflow_amount` objects, with corresponding objects `total_inflow_amount_30d`, `total_inflow_amount_60d`, `total_inflow_amount_90d`, `total_outflow_amount_30d`, `total_outflow_amount_60d`, and `total_outflow_amount_90d`, to summarize inflows and outflows.
- For Partner Insights Reports, deprecated the `version` field (which used type `number`) and replaced it with the `model_version` field (which uses type `string`), to better support the not-strictly-numerical versioning scheme used by Prism.

For Identity Verification:

- For the response bodies of the endpoints [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate), [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget), [`/identity_verification/list`](/docs/api/products/identity-verification/#identity_verificationlist), and [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry), added the `verify_sms` object. This object contains details about SMS verification attempts, with sub-fields such as `status`, `attempt`, `phone_number`, `delivery_attempt_count`, `solve_attempt_count`, `initially_sent_at`, `last_sent_at`, and `redacted_at`.
- Restored the "Integration" button to the Dashboard. It can now be found next to the "Publish Changes" button in the upper right of the Template Editor.

For Payment Initiation:

- For [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate), added `payer_details` object. The fields `options.bacs` and `options.iban` have been deprecated in favor of `payer_details.numbers.bacs` and `payer_details.numbers.iban`; customers are encouraged to migrate to the new fields.

For Signal:

- Made the `scores` field nullable in the responses for [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) and [`/processor/signal/evaluate`](/docs/api/processor-partners/#processorsignalevaluate), as a pre-requisite for adding future Signal support to Items where a score cannot be calculated, such as Items added via Database Insights or Same-Day Micro-deposits.
- Changed Signal billing to behave similarly to Auth and Identity if it is specified in the `optional_products` array -- customers will not be billed for Signal if it is specified in `optional_products` until [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) is called.

##### October 28, 2024

- Fully removed legacy returning user flows (Institution Boosting, Pre-Matched RUX, and Pre-Authenticated RUX) in favor of the new returning user experience.
- Removed legacy returning user flow errors, `INVALID_PRODUCTS` and `PRODUCT_UNAVAILABLE`.

##### October 15, 2024

- For the `PENDING_DISCONNECT` webhook, added `INSTITUTION_TOKEN_EXPIRATION` as a `reason` code. This reason will be used if the user's access-consent on an OAuth token is about to expire. This webhook code will only be used for Items associated with institutions in the US and Canada; for Items in the UK and EU, Plaid will continue to send the `PENDING_EXPIRATION` webhook instead of the `PENDING_DISCONNECT` webhook in the event of pending OAuth access expiration.

For Plaid Check Consumer Reports:

- Reverted the change to historical balance data from the October 11 update. Updated the documentation to clarify the behavior and reduce customer confusion.
- Added `attributes` object with fields `nsf_overdraft_transactions_count`, `is_primary_account`, and `primary_account_score` to Base Report object, to help identify primary accounts and overdraft transactions.

For Payment Initiation:

- For the `/payment_initiation/consent/*` endpoints, deprecated the `scopes` field in favor of the `type` field, with possible values `SWEEPING` or `COMMERCIAL`. The `scopes` field will still be honored if used, but customers are encouraged to use the new `type` field instead, and the `scopes` field will be removed in the future.

##### October 11, 2024

For Plaid Check Consumer Reports:

- For [`/user/create`](/docs/api/users/#usercreate), added the fields `last_4_ssn` and `date_of_birth` to the `consumer_report_user_identity` object. All user tokens created after January 31, 2025 must have a `date_of_birth` populated in order to be compatible with Plaid Check Consumer Reports.
- Changed the behavior of returning historical balances. Previously, Consumer Reports would always return at least 90 days of historical balance information. This caused user confusion in the case of new accounts opened less than 90 days ago, since Consumer Report would report a historical balance on the account for dates when the account did not yet exist. Now, Consumer Reports will only return historical balances up to the date of the oldest transaction found in a given account.
- Released Cash Flow Updates and Network Insights to beta. For details, see [Plaid Check](/docs/check/#overview).

For Identity Verification:

Introduced several updates to Identity Verification. These changes are being gradually rolled out to customers. To request to be opted in to these changes early, or to delay receiving these changes, contact your Plaid Account Manager.

- Introduced Trust Index, a numerical scoring system rating the user on a scale of 1 to 100, with 11 sub-scores on different dimensions (e.g. email, phone, liveness), to help you understand users' relative risk across various dimensions and how different factors contribute to that risk.
- An overhauled session view that helps you digest a full user's verification without scrolling.

For Income:

- Fixed an issue in which customers could successfully call [`/link/token/create`](/docs/api/link/#linktokencreate) for Income without specifying a `user_token`. The `user_token` field is required when using Income.

##### October 9, 2024

- Began the automatic enrollment of some customers into [Data Transparency Messaging (DTM)](/docs/link/data-transparency-messaging-migration-guide/). Throughout Q4, customers serving end users in the US and Canada will be enrolled in DTM automatically if they have not already enrolled.
- Announced the elimination of public key support, effective January 31, 2025. Beginning in February 2025, it will no longer be possible to launch Link sessions using a public key. This impacts only a small number of customers, as the Plaid public key has been deprecated since the introduction of Link tokens in mid-2020, and no teams created after July 2020 have been issued public keys. All customers who are still using public keys were notified of this change via email in September. For instructions on migration from public keys to Link tokens, see the [Link token migration guide](/docs/link/link-token-migration-guide/).
- Announced the full removal of legacy returning user flows (Institution Boosting, Pre-Matched RUX, and Pre-Authenticated RUX) in favor of the new returning user experience, to be completed by October 28, 2024.
- Eligible customers (those with support packages) can enroll in SSO via the [Dashboard](https://dashboard.plaid.com/settings/team/sso). If you would like to use SSO and do not have a support package, contact your Plaid account manager.
- To reduce developer confusion, updated the [`/link/token/get`](/docs/api/link/#linktokenget) endpoint to return the various results arrays and objects (`item_add_results`, `payroll_income_results`, `document_income_results,` `bank_income_results`, `cra_item_add_results`) as empty arrays (if arrays) or as null (if objects) when not present, rather than being omitted from the schema.

For Link:

- Improved the user experience for Remember Me verification on eligible Android devices by automatically filling in the OTP code received on the device.

For Auth:

- PNC now returns Tokenized Account Numbers (TANs) with behavior similar to Chase.

For Balance:

- Discontinued the Balance Plus beta program. Balance Plus beta will no longer accept new enrollments; customers currently in the beta program have been contacted directly with more details and next steps.

For Plaid Check Consumer Report:

- Released the Plaid Check Third-Party User Token to beta, allowing customers to partner with select lenders to provide their customers access to credit. If you are interested in this feature, contact sales or your Plaid Account Manager.
- Integrated Plaid Check Consumer Reports with Layer to enable users to share cash flow insights alongside identity and bank data instantly, dramatically streamlining the loan application process. Users consent to share their consumer report in the final Layer pane. If you are interested in this feature, contact sales or your Plaid Account Manager.
- Updated the [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget) endpoint transaction history behavior for new customers. If no value is specified for the `days_requested` field, Plaid will default to requesting 365 days, and if a value under 180 is specified, Plaid will request 180 days. This change increases the quality of transaction history insights, as the more transaction history is requested, the more accurate the insights returned will be. This change impacts only customers obtaining Plaid Consumer Report Production access on or after October 1, 2024; the behavior for existing customers will not change.
- In the [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget) response `warnings` array, added a new possible warning code, `USER_FRAUD_ALERT`, which indicates that the user has placed a fraud alert on their Plaid Check consumer report due to suspected fraud.
- For the [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget) endpoint, added an `error_reason` to the response to surface Prism error codes.

For Identity:

- Released [Identity Document Upload](/docs/identity/identity-document-upload/) to beta. Identity Document Upload provides document-based account ownership identity verification for Items that do not support [`/identity/get`](/docs/api/products/identity/#identityget) or [`/identity/match`](/docs/api/products/identity/#identitymatch), including Items connected via Same-Day Micro-deposits or Database Insights.

For Identity Verification:

- Redesigned the fraud labeling system UI to allow 1-click fraud reports from the Dashboard.
- Fixed ordering of Link events callbacks when Hide Verification Outcome is enabled so that `IDENTITY_VERIFICATION_CLOSE_UI` fires and `onSuccess` is called after the event corresponding with the outcome of the session has fired.

For Layer:

- Integrated Plaid Check Consumer Reports with Layer to enable users to share cash flow insights alongside identity and bank data instantly, dramatically streamlining the loan application process. Users consent to share their consumer report in the final Layer pane. If you are interested in this feature, contact sales or your Plaid Account Manager.

For Payment Initiation:

- For [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute) endpoint, added the optional `processing_mode` parameter. This allows you to opt in to async payment execution processing, allowing for better performance and throughput when realtime payment processing results are not necessary, such as in user-not-present flows.

For Transfer:

- Improved [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) to consider more information about accounts, such as whether the account has previously returned codes R02 (account closed) or R16 (account frozen) when used with Plaid.

##### September 2024

- Released [Link Recovery](/docs/link/link-recovery/) to beta. With Link Recovery, your end users can sign up to be automatically notified by Plaid when an institution outage that prevented them from linking an account has been resolved. When the issue is over, Plaid will send users a link that can be used to connect their account to your app.
- Released [Investments Move](/docs/investments-move/) to Early Availability. Investments Move facilitates ACATS and ATON brokerage transfers by providing user-permissioned data including a user's account number, account information, and detailed holdings information. Investments Move can remove friction for end users in changing brokerages and reduce the frequency of transfer failures caused by data entry errors.
- Improved retry logic for missed webhooks. If a webhook sent by Plaid is not accepted by the webhook receiver endpoint with a `200` status within 10 seconds of being sent, Plaid will now retry up to six times over a 24-hour period, using an exponential backoff algorithm, rather than the previous behavior of retrying up to twice, a few minutes apart.
- Added a new `PENDING_DISCONNECT` Item webhook to alert customers of Items that are about to enter an unhealthy state. Unlike the existing `PENDING_EXPIRATION` webhook, `PENDING_DISCONNECT` covers Items that will break due to reasons other than consent expiration, such as a planned sunset by the bank of an old online banking platform, or a required migration to OAuth. Upon receiving a `PENDING_DISCONNECT` webhook, customers should direct the end user to the update mode flow for the Item.
- Added the ability to use update mode to reauthorize consent for a US or CA Item whose consent is expiring due to 1033 rules. If the consent expiration date is within 6 months, Plaid will automatically route the Item through the longer reauthorization update mode flow, which will cause the expiration date to be pushed to 12 months from the date that the user reauthorizes the Item. You can customize whether or not the Item goes through the reauthorization update flow by using the `update.reauthorization_enabled` parameter in [`/link/token/create`](/docs/api/link/#linktokencreate). Note that there is no reason to send Items through the reauthorization flow at this time, as Items will not be assigned consent expiration dates until approximately December 2024. This change only impacts Items in the US and Canada; the OAuth update flow for Items in the UK and EU has not changed.
- To help manage consent status for the upcoming wider Data Transparency Messaging rollout, added `consented_use_cases`, `consented_data_scopes`, and `created_at` fields to the Item object.
- Removed `deposit_switch` from the `products` array, as part of Deposit Switch deprecation.

For Link flow:

- Shortened and clarified text shown to end user during prompt to verify 2FA code.
- Added hover animations to UI buttons in Link.

For Plaid Link SDKs:

- Released iOS SDK 5.6.1, which adds adds haptics support, fixes embedded search view dynamic resizing, and adds missing event names and view names. For details, see the [iOS SDK changelog](https://github.com/plaid/plaid-link-ios/releases).
- Released Android SDK 5.6.1, which adds missing event names and view names. For details, see the [Android SDK changelog](https://github.com/plaid/plaid-link-android/releases).
- Released React Native SDK 11.13.0, which incorporates iOS SDK 5.6.1 and Android SDK 4.6.1. For details, see the [React Native SDK changelog](https://github.com/plaid/react-native-plaid-link-sdk/releases).

For Assets:

- Added [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) request parameter `options.require_all_items`, with a default value of `true`. By setting this to `false`, customers can optionally choose to generate an Asset Report even if one or more Items in the Asset Report encounters an error, as long as at least one Item succeeded. Otherwise, the endpoint will maintain its current behavior of only creating an Asset Report if all Items could be successfully retrieved.

For Consumer Report:

- Fixed confusing pluralization by renaming `longest_gap_between_transactions` to use `gaps` rather than `gap` as this field is array of gaps rather than a single gap. The old version of this field is deprecated and will be removed at the end of October.
- Removed the `products` field from the [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) request, as its usage there is deprecated.

For Identity Verification:

- Added `liveness_check` results to [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget) to improve parity between API and Dashboard data availability.
- Added `name` object to the `extracted_data` within each `documentary_verification.documents` object in the response of all of the Identity Verification endpoints.
- Design updates to the Identity Verification Dashboard.

For Investments:

- Added `sector` and `industry` fields to the securities object to help categorize holdings.

For Transfer:

- Updated [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) to accept either a `transfer_id` or a `transfer_authorization_id`.
- In [`/transfer/ledger/distribute`](/docs/api/products/transfer/ledger/#transferledgerdistribute) request, replaced `from_client_id` and `to_client_id` with `from_ledger_id` and `to_ledger_id`.

##### August 2024

- Added [Multi-Item Link](/docs/link/multi-item-link/), which allows end users to link multiple financial accounts in a single Link session.
- Added the `error_code_reason` field to the Plaid error object to provide more useful troubleshooting details for `ITEM_LOGIN_REQUIRED` errors at OAuth institutions. This field is currently in beta and may not be present for all customers.
- Added the ability to launch update mode with a user token and augmented [`/link/token/create`](/docs/api/link/#linktokencreate) with an `update.item_ids` field to optionally specify which Item ID(s) to update when launching Link in update mode with a user token. Launching update mode with a user token is the only supported update mode flow for Consumer Report. For products that use both access tokens and user tokens, like Income, update mode can be launched using either one.

For Plaid Link SDKs:

- Released an alternative installation repo to reduce download sizes for projects using Swift Package Manager. For details, see the [iOS SDK changelog](https://github.com/plaid/plaid-link-ios/releases).
- Released Android SDK 4.6.0, with dependency updates and improved Java compatibility. For more details, see the [Android SDK changelog](https://github.com/plaid/plaid-link-android/releases).
- Released React Native 11.2.1, with enhanced support for React Native New Architecture, and including upgrades to Android SDK 4.6.0 and iOS SDK 5.6.1. For more details, see the [React Native SDK changelog](https://github.com/plaid/react-native-plaid-link-sdk/releases).

For Plaid Link:

- The password input boxes in Link now allow the user to optionally reveal the password before submitting their credentials, in order to help reduce typos in password data entry.

For Beacon (beta):

- Added data breach reporting.

For Plaid Check Consumer Reports:

- Added a `consumer_disputes` field to the response for [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), showing the details of any disputes the consumer has submitted about the information in the report.

For Identity Verification:

- Added `is_shareable` parameter to [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry) to allow configuring different retry behavior for the original session and retry sessions. If set, this parameter will control whether a shareable link is generated for the retry session. If not set, the retry session will use the same shareable link behavior as the original session.

For Income:

- Added `file_type` field to [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget).
- Added the ability to customize the `ytd_amount` field in the Sandbox custom user for [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget).

For Investments:

- Added Investments Refresh. Investments Refresh is available as an add-on to Investments and allows you to make a real-time, on-demand update request for fresh Investments data by calling the [`/investments/refresh`](/docs/api/products/investments/#investmentsrefresh) endpoint.

For Liabilities:

- Plaid has discontinued support for federal student loan providers: Aidvantage, Central Research, Inc., EdFinancial Services - Federal Direct Loan Program (FDLP), Mohela and NelNet. Beginning in September, Liabilities subscriptions on Items from these institutions will no longer be billed, and end users will not be able to select these institutions in Link. Private student loan providers are not impacted. Impacted customers have been contacted directly by their Account Managers. For questions, contact your Plaid Account Manager.

For Transfer:

- Added support for multiple Ledgers. This improves the Plaid Transfer experience for customers who need to avoid co-mingling funds from different types of sources. As part of this change, several Transfer endpoints now have the optional ability to specify a `ledger_id` in the request and/or will return a `ledger_id` as part of the response. To request multiple Ledgers, contact your Plaid Account Manager.
- Improved the flow for authorizations that could not be evaluated due to the `ITEM_LOGIN_REQUIRED` error by enabling [update mode](/docs/link/update-mode/) for these authorizations, allowing them to be retried after the Item is fixed. Authorizations that previously would have had an `approved` decision and a `rationale_code` of `ITEM_LOGIN_REQUIRED` will now have the new `user_action_required` decision instead. If [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) returns an `user_action_required` decision, you can now launch Link in update mode by creating a Link token with the `authorization_id` passed to [`/link/token/create`](/docs/api/link/#linktokencreate). After the end user has completed the update mode flow, you can retry the [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) call.
- For [`/transfer/metrics/get`](/docs/api/products/transfer/metrics/#transfermetricsget), added an [`authorization_usage`](/docs/api/products/transfer/metrics/#transfermetricsget) response object allowing you to see your credit and debit utilization details.

For Virtual Accounts:

- Added the `related_transactions` array to the Wallet transaction object. This field makes it easier to identify associated transactions by showing the transaction ID of related transactions, as well as how they are related. For example, if a transaction is refunded, the original transaction and the refunded transaction will be linked via the `related_transactions` field.

##### July 2024

- Added support for Silent Network Auth (SNA) for the Android returning user experience. SNA is a faster and more secure method to verify an end user's phone number. To use SNA, you must be using at least version 3.14.3 or later of the Android SDK (if using 3.x) or version 4.1.1 or later (if using the 4.x). On React Native, the corresponding minimum versions are 10.12.0 and 11.4.0. Aside from updating your SDK, no action is required to enable SNA.
- Added support for biometric authentication for Citibank for integrations using the iOS SDK version 5.1.0 or later.
- Released account holder category (beta), indicating whether an account returned in the account object is a business or personal account. To request access to this feature, contact your account manager.

For Plaid Link SDKs:

- Released Android SDK 4.6.0, containing bug fixes and compatibility improvements. For details, see the [Android SDK changelog](https://github.com/plaid/plaid-link-android/releases).

For Auth:

- Released [AUTH: DEFAULT\_UPDATE](/docs/api/products/auth/#default_update) webhook to GA. This webhook fires if Plaid detects that a bank's Auth information has changed. (While rare, this can occur due to changes such as an acquisition.) To avoid ACH returns, after receiving this webhook, you should call [`/auth/get`](/docs/api/products/auth/#authget) or [`/processor/auth/get`](/docs/api/processor-partners/#processorauthget) to update your information on file for the user.

For Identity:

- In [`/identity/match`](/docs/api/products/identity/#identitymatch), `legal_name.is_business_name_detected` is no longer deprecated and can now be used for detecting business names.

For Transfer:

- Added [`return_rates`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-return-rates) field to [`/transfer/metrics/get`](/docs/api/products/transfer/metrics/#transfermetricsget), allowing you to get return rate information via the API.

##### June 2024

- Launched [Plaid Layer](https://plaid.com/docs/layer) to early availability. Layer provides a fast, high-converting onboarding experience powered by Plaid Link.
- Launched [Plaid Check](https://plaid.com/docs/check/), a subsidiary of Plaid Inc. that is a Consumer Reporting Agency. Through Plaid Check's Consumer Report API, you can retrieve ready-made credit risk insights from consumer-permissioned cash flow data.
- Decommissioned the Development Environment. Development has been replaced for testing purposes by Limited Production. For more details, see [Limited Production](https://plaid.com/docs/sandbox/#testing-with-live-data-using-limited-production).

For Plaid Link SDKs:

- Released iOS SDK 5.6.0, containing support for Plaid Layer, as well as bug fixes and improvements to the returning user experience.
- Released iOS SDK 4.7.9, containing bug fixes and improvements to the returning user experience.
- Released Android SDK 4.5.0, containing support for Plaid Layer, as well as bug fixes.
- Released React Native 11.11.0, containing support for Plaid Layer, as well as bug fixes and improved prefill support.

For Beacon:

- Added `access_tokens` field to [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate) and [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate) requests.
- Added `item_ids` to `/beacon/user/*` responses.
- Added Bank Account Insights to Beacon. Account Risk (beta) is being folded into Beacon

For Balance:

- Launched Balance Plus (beta). Balance Plus enhances Plaid's Balance product with additional insights and lower latency and is designed to require minimal integration work for existing Balance customers.

For Income:

- Launched [Plaid Check](https://plaid.com/docs/check/), a subsidiary of Plaid Inc. that is a Consumer Reporting Agency. Through Plaid Check's Consumer Report API, you can retrieve ready-made credit risk insights from consumer-permissioned cash flow data.
- Added [`/user/remove`](/docs/api/users/#userremove) to support the deletion of user tokens.

For Liabilities:

- Added `income-sensitive repayment` as a student loan repayment type.

For Transactions:

- As originally announced in the November 2023 changelog, the default number of days of transactions history requested for new transactions-enabled Items is now 90 days for all customers. To change this behavior, use the `days_requested` parameter in [`/link/token/create`](/docs/api/link/#linktokencreate) or (if initializing transactions post-Link) [`/transactions/get`](/docs/api/products/transactions/#transactionsget) or [`/transactions/sync`](/docs/api/products/transactions/#transactionssync).
- Improved behavior of Transactions-enabled Items being fixed via update mode. Now, if an Item was in an error state for over 24 errors before entering update mode, Plaid will immediately extract fresh transactions after update mode has completed, rather than waiting for the next scheduled update time.

For Transfer:

- Added support for RTP to Transfer UI.
- Added improved support for Transfer limit increase requests via the Transfer Dashboard.

For Virtual Accounts:

- Added support for `RECALL` as a transaction type where the sending bank has requested the return of funds due to a fraud claim, technical error, or other issue associated with the payment.

##### May 2024

- Released [Limited Production](/docs/quickstart/glossary/#limited-production), which will replace Development as a platform for testing for free with live data. The Development environment will be decommissioned on June 20, 2024, and all Development Items will be deleted.
- Released [Data Transparency Messaging](/docs/link/data-transparency-messaging-migration-guide/) to GA. To help ensure compliance with 1033 rulemaking, DTM will be automatically enabled for US and CA Link sessions later this year.
- Added new pre-seeded Sandbox test accounts for testing different credit profiles. These accounts are especially useful for testing Income, but are available to all customers. For details, see [Credit and Income testing credentials](/docs/sandbox/test-credentials/#credit-and-income-testing-credentials).
- Updated the response object for [`/link/token/get`](/docs/api/link/#linktokenget). To accommodate future multi-Item Link session support for additional products, the `on_success` object in the response has been deprecated; it is recommended to use the new `results` object instead. Old: `link_sessions[0].on_success.public_token`. New: `link_sessions[0].results.item_add_results[0].public_token`. The `on_exit` field has also been deprecated; it is recommended to use the new `exit` field.
- Updated the `SESSION_FINISHED` webhook. To accommodate future multi-Item Link session support for additional products, deprecated the `public_token` field in lieu of a new `public_tokens` array.

For Plaid Link SDKs:

- Released iOS SDK 5.5.0 and 4.7.7, containing bug fixes and (for iOS SDK 5.5.0) additional view names to support new functionality.
- Released iOS SDK 6.0.0 beta, containing support for Apple Card and FinanceKit.
- Released Android SDK 4.4.1, containing bug fixes and UI improvements.
- Released React Native SDK 11.10.1, with improved support for "pre-loading" Link for a lower-latency user experience.
- Released React Native SDK 12.0.0 beta, containing support for Apple Card and FinanceKit.

For Income:

- Added new values to `canonical_description` in [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget): `RETIREMENT`, `GIG_ECONOMY`, and `STOCK_COMPENSATION`.

For Signal:

- Added support for some Items added via Same Day Micro-deposits or Instant Micro-deposits.

For Transfer:

- Added [`/transfer/authorization/cancel`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcancel) endpoint.

##### April 2024

- Updated [`/link/token/get`](/docs/api/link/#linktokenget) response to include `user_token`.
- Rebranded Remember Me as returning user experience.

For Plaid Link SDKs:

- Released Android SDK 4.3.1, iOS SDK 4.5.2, and React Native SDK 11.8.2, containing bug fixes and UI enhancements.

For Auth:

- Released Database Match, which enables instant manual account verification using Plaid's database of known account numbers. When provided as an alternative to Same Day Micro-deposits, Database Match can increase conversion, as the user may be able to verify instantly, without having to return to Plaid to verify their micro-deposit codes.
- Released [Database Insights (beta)](/docs/auth/coverage/database/). Database Insights can be used as a manual alternative to credential-based Auth flows in low-risk use cases. Database Insights verifies account and routing numbers by checking the information provided against Plaid's known account numbers. If no known account number is found, Database Insights checks the information given against usages associated with the given routing number. For more information and to request access to the beta, see the [Database Insights (beta) documentation](/docs/auth/coverage/database/).
- Added support for the `SMS_MICRODEPOSITS_VERIFICATION` webhook to [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook) to support testing the [Text Message Verification flow](/docs/auth/coverage/same-day/#using-text-message-verification).

For Income:

- Added the ability to customize test data in the Sandbox environment for use with Document Income, by uploading a special JSON configuration object during the document upload step. For details, see [Testing Document Income](/docs/income/document-income/#testing-document-income).

For Payment Initiation:

- Added `supports_payment_consents` field to an institution's Payment Initiation metadata object.
- Added counterparty date of birth and address as optional request fields for [`/payment_initiation/payment/reverse`](/docs/api/products/payment-initiation/#payment_initiationpaymentreverse). Providing these fields increases the likelihood that the recipient's bank will process the transaction successfully.

For Statements:

- Released [Statements](/docs/statements/) to General Availability.
- Added Statements support to [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate).

For Transfer:

- Updated the possible values for `network` in recurring transfer contexts to reflect the fact that recurring transfers now support RTP.
- Added `funds_available` transfer status and transfer event type.
- Deprecated the `/transfer/balance/get` endpoint in favor of the Plaid Ledger.

##### March 2024

For Plaid Link SDKs:

- Multiple improvements to the Remember Me experience.
- Added the ability to "pre-load" Link for a lower-latency user experience.

For Auth:

- Enabled [Text Message Verification](/docs/auth/coverage/same-day/#using-text-message-verification) flow for Same-Day Micro-deposits by default. This new, non-breaking change can increase conversion of micro-deposit verification by up to 15%. This flow is now enabled for all Same Day Micro-deposit flows by default; to opt-out, use the `auth.sms_microdeposits_verification_enabled: false` setting when calling [`/link/token/create`](/docs/api/link/#linktokencreate).

For Transactions:

- Multiple improvements to the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) endpoint addressing common pain points:
  - [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) now returns the `accounts` array, eliminating the need for a separate call to [`/accounts/get`](/docs/api/accounts/#accountsget).
  - [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) now returns a `transactions_update_status`, reflecting the status of the transaction pulls, which was previously only available via webhook.
  - The list of `removed` transactions now includes the `account_id`s they were associated with, to make reconciliation and management easier.
- Granted all customers self-serve access to the `original_description` field. To access original transaction descriptions, set `options.include_original_description` to `true` when calling [`/transactions/get`](/docs/api/products/transactions/#transactionsget) or [`/transactions/sync`](/docs/api/products/transactions/#transactionssync). It is no longer necessary to request access from your Account Manager.
- Made `account_id` optional in the [`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget) endpoint.

For Transfer:

- Added `has_more` field to [`/transfer/event/list`](/docs/api/products/transfer/reading-transfers/#transfereventlist) and [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) to indicate there are more events to be pulled.
- Updated the possible values for `network` in recurring transfer contexts to reflect the fact that recurring transfers are only supported via ACH.

For Processor Partners:

- Added `institution_id` to [`/processor/account/get`](/docs/api/processor-partners/#processoraccountget) endpoint.

##### February 2024

For Auth:

- Released the [Text Message Verification](/docs/auth/coverage/same-day/#using-text-message-verification) flow for Same-Day Micro-deposits. This new, non-breaking change can increase conversion of micro-deposit verification by up to 15%. Currently, Text Message Verification is opt-in, and will transition to opt-out in March.

For Identity Verification:

- Added a large number of additional `linked_services` enum values.

For Investments:

- Added `vested_quantity` and `vested_value` fields to the `holdings` object.

For Liabilities:

- Added the `pending idr` status for student loans to reflect loans with a pending application for income-driven repayments.

For Transfer:

- Added fields including `wire_details` and `wire_routing_number`, and `wire` as a supported `network`, to support wire transfers. If you are interested in using wire transfers with Transfer, contact your Account Manager or (if you are not yet a customer) Sales.
- Increased maximum length of `description` field on [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) from 8 to 15.

For Signal:

- Added `warnings` array to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) response even if no warnings are present in order to match documented API behavior.

##### January 2024

- Expanded the Returning User Link flow to more customers. For more details on the returning user experience, including the associated Link events, see [returning user experience](/docs/link/returning-user/).
- Added `events` field to [`/link/token/get`](/docs/api/link/#linktokenget) to capture Link events.
- Made the `persistent_account_id` field available in GA to all customers. The `persistent_account_id` field is a special field, available only for Chase Items. Because Chase accounts have a different account number each time they are linked, this field may be used instead of the account number to uniquely identify a Chase account across multiple Items for payments use cases, helping to reduce duplicate Items or attempted fraud.
- Released the [`USER_ACCOUNT_REVOKED`](/docs/api/items/#user_account_revoked) webhook in GA to all customers, to complement the existing [`USER_PERMISSION_REVOKED`](/docs/api/items/#user_permission_revoked) webhook, but capturing account-level revocation rather than Item-level revocation. The `USER_ACCOUNT_REVOKED` webhook is sent only for Chase Items and is primarily intended for payments use cases. The webhook indicates that the TAN associated with the revoked account is no longer valid and cannot be used to create new transfers.

For Auth:

- Added support for the Stripe Payment Intents API. For more details, see [Add Stripe to your App](/docs/auth/partnerships/stripe/).

For Assets:

- Added `vested_quantity` and `vested_value` fields to the Investments object within the Asset Report.
- Added `margin_loan_amount` field to the Balance object within the Asset Report.

For Enrich:

- Added counterparty `phone_number` to the [`/transactions/enrich`](/docs/api/products/enrich/#transactionsenrich) response.

For Income:

- Added the ability to filter the results of [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) and [`/credit/bank_statements/uploads/get`](/docs/api/products/income/#creditbank_statementsuploadsget) to only certain Items, via an optional `item_ids` request field.
- Added `num_1099s_uploaded` to the `document_income_results` object in [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) response.

For Identity Verification:

- Added new `no_data` type to `name` and `date_of_birth` fields in `documentary_verification.documents[].analysis.extracted_data` in the response of all of the Identity Verification endpoints.
- Made `street` and `city` optional in the address attribute of [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate).

For Payment Initiation:

- Added optional `scope` and `reference` fields to [`/payment_initiation/consent/payment/execute`](/docs/api/products/payment-initiation/#payment_initiationconsentpaymentexecute).

For Transactions:

- Added a new Sandbox test user, `user_transactions_dynamic`. When logging into Sandbox with this username, you will see more realistic, dynamic Sandbox transactions data, including a wider variety of transaction types, and transactions moving between pending and posted state.

For Virtual Accounts:

- Added `available` balance to balance object in [`/wallet/get`](/docs/api/products/virtual-accounts/#walletget) and [`/wallet/list`](/docs/api/products/virtual-accounts/#walletlist).

For Partners:

- Added [`/processor/liabilities/get`](/docs/api/processor-partners/#processorliabilitiesget) endpoint.
- Added optional `registration_number` to [`/partner/customer/create`](/docs/api/partner/#partnercustomercreate) request.

##### December 2023

- Announced the upcoming removal of the Development environment. To simplify the development process, Plaid will be replacing the Development environment with the ability to test with real data for free directly in Production. On June 20, 2024, the Plaid Development environment will be decommissioned, and all Development Items will be removed. You may continue to test on the Development environment until this time.
- Capital One now provides real-time balance data for depository accounts, including checking and savings accounts. It is no longer necessary to include `min_last_updated_datetime` when making [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) calls for these accounts.

For Investments:

- Added `market_identifier_code` to the security object.
- Added `option_contract` to the security object, adding additional details `contract_type`, `expiration_date`, `strike_price`, and `underlying_security_ticker` for better insight into derivatives.

##### November 2023

For Identity:

- Released [Identity Match](/docs/identity/#identity-match) and the `/identity/match/` endpoint to General Availability.

For Income:

- Expanded [`/sandbox/income/fire_webhook`](/docs/api/sandbox/#sandboxincomefire_webhook) to support the `INCOME_VERIFICATION_RISK_SIGNALS` webhook, adding a new `webhook_code` parameter to the endpoint.

For Transfer Platform Payments (beta):

- Added [`/transfer/ledger/distribute`](/docs/api/products/transfer/ledger/#transferledgerdistribute) endpoint, as well as a `facilitator_fee` field to [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate), both of which can be used to collect payments from end customers.
- Removed support for doc and docx format files in `/transfer/diligence/document/upload`.

For Transfer:

- Updated default Sandbox behavior of Ledger. Ledgers in Sandbox will now have a $100 default balance.

For Transactions:

- Added the `days_requested` field to [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) and [`/transactions/get`](/docs/api/products/transactions/#transactionsget) under the `options` object, and to [`/link/token/create`](/docs/api/link/#linktokencreate) under the `transactions` object to support a behavior change in the Transactions product. For all new customers who created accounts after December 3, 2023 the maximum number of days of historical data requested from transactions endpoints will default to 90 days if the `days_requested` field is not specified. This change will also be applied to existing customers on June 24, 2024. To change the amount of historical data requested, use the `days_requested` field when Transactions is first added to your Item.

For Partners:

- Added [`/processor/signal/prepare`](/docs/api/processor-partners/#processorsignalprepare)
- Added [`/processor/transactions/get`](/docs/api/processor-partners/#processortransactionsget)
- Make `products` field in `/partner/customers/create` optional. If empty or `null`, this field will default to the products enabled for the partner.

##### October 2023

- Chase has extended the September 30 deadline for migrating away from in-process webviews. The new cutoff date will be in mid-January 2024. All impacted customers should migrate by January 1, 2024 to avoid a negative impact on Link conversion. See the [deprecation notice](https://docsend.com/view/h3qdupjusiwyjvv5) for more details.
- Added Belgium as a supported country.

For Auth:

- Added an `instant_microdeposits_enabled` flag to [`/link/token/create`](/docs/api/link/#linktokencreate) to allow customers to disable Instant Micro-deposits. By default, Instant Micro-deposits will be enabled for all sessions.

For Income:

- Released Document Fraud to General Availability.
- Added [`/credit/payroll_income/parsing_config/update`](/docs/api/products/income/#creditpayroll_incomeparsing_configupdate), allowing users of Document Fraud to update the parsing configuration used for document income.
- Deprecated the /credit/payroll\_income/precheck endpoint.

For Investments:

- Added the `investments.allow_manual_entry` parameter in [`/link/token/create`](/docs/api/link/#linktokencreate) and the corresponding `is_investments_fallback_item` field to [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) and [`/investments/holdings/get`](/docs/api/products/investments/#investmentsholdingsget). Enabling manual entry allows a user to create an investments fallback Item for a non-supported institution by manually entering their holdings in Link.
- Extended the deadline beyond which Plaid customers will need a CGS license to obtain `cusip` and `isin` data. The new deadline is March 2024.

For Transactions:

- Added `counterparties`, `logo_url`, and `website` information about merchants (previously available only via Enrich) to [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) and [`/transactions/get`](/docs/api/products/transactions/#transactionsget) data.

For Transfer:

- Added new methods to fund your Plaid Ledger. In addition to the [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit) and [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw) endpoints, you can now initiate a funds transfer via the Dashboard, set up a recurring schedule or minimum balance by contacting Support, or fund your ledger via ACH or wire transfer from your bank account.
- Made Plaid Ledger the default method for funding transfers for all new Transfer customers. Plaid Ledger allows faster funds transfers and is now the easiest way to move money between your bank account and Plaid. Existing customers who want to switch to Plaid Ledger should contact their Plaid Account Manager.
- Added [`/sandbox/transfer/refund/simulate`](/docs/api/sandbox/#sandboxtransferrefundsimulate) to test refunds in the Sandbox environment.
- Added refund-specific event types. These are existing event types with the prefix `refund.`, e.g. `refund.settled`.
- For [`/transfer/configuration/get`](/docs/api/products/transfer/metrics/#transferconfigurationget), stopped returning `max_single_transfer_amount` and `max_monthly_amount`. To maintain compatibility with older client libraries, these fields will still be present in the API, but will be blank strings.
- For [`/transfer/metrics/get`](/docs/api/products/transfer/metrics/#transfermetricsget), stopped returning `monthly_transfer_volume`. To maintain compatibility with older client libraries, this field will still be present in the API, but will be a blank string.

For Virtual Accounts:

- Added `RETURN` as possible Virtual Account wallet transaction type.

##### September 2023

- Added the `LOGIN_REPAIRED` webhook. This webhook will fire when an Item's status transitions from `ITEM_LOGIN_REQUIRED` to a healthy state, without the user having used the update mode flow in your app.
- Added [Update Mode Product Validations (beta)](/docs/link/update-mode/#resolving-access_not_granted-or-no_auth_accounts-errors-via-product-validations). Update Mode Product Validations (UMPV) allows you to validate that a user going through update mode at an OAuth institution has selected the appropriate authorizations to enable Auth and/or Identity. If they do not select the correct options, they will be prompted to update their selections.
- Added [Optional Products](/docs/link/initializing-products/#optional-products). If a product is specified in the `optional_products` field when calling [`/link/token/create`](/docs/api/link/#linktokencreate), the product will be added to the Item if possible. However, if the product cannot be added to the Item -- for example, because the Item is not supported by the institution, the user did not grant the correct OAuth permissions, or because of an institution connectivity error for that product -- Item creation will still succeed with the other products specified.
- Added `institution_not_supported` as a potential Link session exit metadata status.

  For Auth:
- Added [Instant Micro-Deposits](/docs/auth/coverage/instant/#instant-micro-deposits), allowing end users to verify their accounts in seconds, using micro-deposits sent over RTP or FedNow rails. Instant Micro-Deposits can be used as a faster alternative to Same-Day Micro-Deposits at supported institutions.

  For Assets:
- Deprecated `report_type` field used in [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate), and replaced it with `verification_report_type`. This change was made to avoid confusion with the existing `report_type` field used to by the `PRODUCT_READY` webhook to distinguish between Fast Asset Reports and Full Asset Reports.

For Signal:

- Launched the Signal Dashboard to help customers monitor return rates and tune Signal risk score thresholds.

For Transactions:

- Made [Personal Finance Categories](/docs/transactions/pfc-migration/) returned by default for all users. Personal Finance Categories are Plaid's newest and most accurate way of categorizing transactions, using more intuitive categories that correspond better to personal finance management use cases.

For Transfer:

- Launched Transfer from beta to Early Availability (EA) status.
- Added support for FedNow payments, which can be accessed by specifying `rtp` as the payment network.
- Added Plaid Ledger. Your Plaid Ledger allows you to maintain a funds balance with Plaid, which can be used to fund payouts over FedNow and RTP rails. You can use APIs to move money in and our of your Ledger.
- Removed Payment Profiles; access tokens are now the only supported mechanisms for connecting accounts on Transfer.
- Removed `beacon_id`, as it was used only to support Guarantee. Guarantee is no longer offered and will be replaced by a more sophisticated set of risk management tools within Transfer.

For Income:

- Updated API docs to reflect `"NOT LISTED"` as a potential value for `marital_status`.

For Statements (beta):

- Released [Statements](/docs/statements/) in beta. Statements allows you to retrieve an exact PDF copy of a customer's statement at supported institutions.

##### August 2023

For Enrich:

- Added `frequency` field to recurring transactions to indicate the recurrence frequency.

For Identity Verification:

- Added `area_code` match status to the responses of the [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget) and [`/identity_verification/list`](/docs/api/products/identity-verification/#identity_verificationlist) endpoints.

For Income:

- Added a number of new categories for Payroll Income `canonical_description` field: `ALLOWANCE`, `BEREAVEMENT`, `HOLIDAY_PAY`, `JURY_DUTY`, `LEAVE`, `LONG_TERM_DISABILITY_PAY`, `MILITARY_PAY`, `PER_DIEM`, `REFERRAL_BONUS`, `REIMBURSEMENTS`, `RETENTION_BONUS`, `RETROACTIVE_PAY`, `SEVERANCE_PAY`, `SHIFT_DIFFERENTIAL`, `SHORT_TERM_DISABILITY_PAY`, `SICK_PAY`, `SIGNING_BONUS`, and `TIPS_INCOME`. Also made `null` a valid value.
- Added `status` field to [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) to show the status of evaluated documents.
- Added `payroll_income.parsing_config` to [`/link/token/create`](/docs/api/link/#linktokencreate) to expose configuration options for risk signal evaluation.

For Liabilities:

- Added `saving on a valuable education` (SAVE) as a student loan repayment plan type.

For Signal:

- Added the ability to change an `initiated` status reported via [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport). Previously, attempting to change the status of a reported decision that was originally reported as `initiated` would result in an `INVALID_FIELD` error; now it is allowed.

For Transfer:

- Added `failure_reason` to refund objects to indicate why a Transfer refund failed.

For Virtual Accounts:

- Added `failure_reason` to [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) and [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist) endpoints to indicate why a wallet transaction failed.

##### July 2023

- Added [`/processor/token/webhook/update`](/docs/api/processor-partners/#processortokenwebhookupdate) endpoint to allow customers using processor partners to update the webhook endpoint associated with a processor token.

For Assets:

- Made `asset_report_token` optional in [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) request to support future development of Asset Reports based on other types of tokens.

For Investments:

- Added the `async_update` option for [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget), along with an accompanying `HISTORICAL_UPDATE` webhook. If `async_update` is enabled, the initial call to [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) for an Item will be made as an async call with a webhook, similar to [`/transactions/get`](/docs/api/products/transactions/#transactionsget).

For Enrich:

- Added `confidence_level` to `counterparty` and `personal_finance_category` fields.

For Identity Verification:

- Added `date_of_birth` and `address` fields to `documentary_verification.documents[].extracted_data` in the response of all Identity Verification endpoints.
- Mark all `region` and `postal_code` fields as nullable, to better support countries that do not use these fields.

For Transfer:

- Added `transfer_id` and `status` filters to [`/transfer/sweep/list`](/docs/api/products/transfer/reading-transfers/#transfersweeplist) request.
- Added `status` field to the sweep object.

##### June 2023

- Added the `required_if_supported_products` field to [`/link/token/create`](/docs/api/link/#linktokencreate). When using multiple Plaid products, you can use this field to specify products that must be enabled when the user's institution and account type support them, while still allowing users to add Items if the institution or linked account doesn't support these products. If a product in this field cannot be enabled at a compatible institution (for example, because the user failed to grant the required OAuth permissions, generating an `ACCESS_NOT_GRANTED` error), the Item will not be created, and the user will be prompted to try again in the Link flow.
- Changed the behavior of the `oauth` field on the institution object. Previously, `oauth` would be `true` only if *all* Items at an institution used OAuth connections. Now, `oauth` is true if *any* Items at an institution use OAuth connections.
- Changed the behavior of the `account_selection_enabled` flag in [`/link/token/create`](/docs/api/link/#linktokencreate). To improve conversion, at institutions that provide OAuth account selection flows, this flag will be overridden and treated as though it were always `true`. For other institutions, the flag behavior is unchanged.

For Auth:

- Added the ability to customize Same-day Microdeposit flows via the new `reroute_to_credentials` option in [`/link/token/create`](/docs/api/link/#linktokencreate). Using this field, you can detect whether a user is attempting to use Same-day Microdeposits to add an institution that is supported via a direct Plaid connection, and optionally either require or recommend that they use a direct connection instead.

For Enrich:

- Added `income_source` enum to `counterparty_type` field.

For Identity Verification:

- Added the ability to prefill user data when retrying a verification by providing a `user` object to [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry).
- Added `client_user_id` field to [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate). It is recommended to use this field instead of `user.client_user_id`, although both fields are supported.

For Income:

- Added `num_bank_statements_uploaded` to `document_income_results` object in [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget).

For Signal:

- Removed the maximum length constraint for `client_user_id` when calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) or [`/processor/signal/evaluate`](/docs/api/processor-partners/#processorsignalevaluate).

For Transfer:

- Added the ability to specify a `test_clock_id` when calling [`/sandbox/transfer/sweep/simulate`](/docs/api/sandbox/#sandboxtransfersweepsimulate) or [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate). If provided, the simulated action will take place at the `virtual_time` of the given test clock.
- Made the `virtual_time` field optional when calling [`/sandbox/transfer/test_clock/create`](/docs/api/sandbox/#sandboxtransfertest_clockcreate). If a time is not provided, the current time will be used.
- Updated the `sweep_id` format to be either a UUID or an 8-character hexadecimal string.
- Made `next_origination_date` nullable in the recurring transfer object.

For Virtual Accounts:

- Added address and date of birth to [`/wallet/transaction/execute`](/docs/api/products/virtual-accounts/#wallettransactionexecute).
- Added payment scheme (`scheme`) to [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) and [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist).

##### May 2023

For Assets:

- Added support for `ASSETS: PRODUCT_READY` and `ASSETS: ERROR` webhooks to [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook)

For Identity:

- Added `is_business_name_detected` to the [`/identity/match`](/docs/api/products/identity/#identitymatch) endpoint

For Identity Verification:

- Added the `selfie_check` field to the Income Verification object, to report the status of a selfie check.

For Document Income:

- Added the `INCOME_VERIFICATION_RISK_SIGNALS` webhook to indicate when [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) is ready to be called.

For Transfers:

- Increased the maximum length of the `description` field in [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) to 15 characters.
- Added the `/transfer/balance/get` endpoint.
- To support prefunded credit processing, added the `credit_funds_source` field to distinguish between credits based on sweeps, RTP credit balances, and ACH credit balances, and made the `funding_account_id` field nullable.

For Processor partners:

- Added [`/processor/token/permissions/set`](/docs/api/processors/#processortokenpermissionsset) and [`/processor/token/permissions/get`](/docs/api/processors/#processortokenpermissionsget) endpoints.
- Added [`/processor/identity/match`](/docs/api/processor-partners/#processoridentitymatch) endpoint to allow Processor partners to support the new beta Identity Match feature.

For Signal:

- Added a `warnings` object to the [`/processor/signal/evaluate`](/docs/api/processor-partners/#processorsignalevaluate) response.

##### April 2023

For Income:

- Added `STARTED` and `INTERNAL_ERROR` statuses to [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget).

For Document Income:

- Released the [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) endpoint. This endpoint can be used as part of the Document Income flow to assess a user-uploaded document for signs of potential fraud or tampering.

For Identity Verification:

- Documented several fields in the `user` object of [`/link/token/create`](/docs/api/link/#linktokencreate) as being officially supported methods of pre-filling user data in Link for the Identity Verification flow, as an alternative to [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate). (While this worked in the past, it was not previously an officially supported or documented flow.) As part of this change, un-deprecated the `user.date_of_birth` field.

For Enrich:

- Added `recurrence` and `is_recurring` fields to [`/transactions/enrich`](/docs/api/products/enrich/#transactionsenrich).

##### March 2023

- To maximize conversion, updated Link to use a pop-up window for OAuth authentication rather than redirects on non-webview mobile web flows.
- Clarified that `continue(from:)` is no longer required for iOS OAuth integrations using Plaid's iOS and React Native iOS SDKs.
- Updated error codes for most `RATE_LIMIT` errors to be specific to the endpoint for which they apply.

For Auth:

- Changed the Link UI for Instant Match and Automated Micro-deposits flows to increase conversion by reducing the number of screens in the flow and removing unnecessary input fields. No developer action is required as a result of this change.

For Transactions:

- Added `ANNUALLY` as a possible frequency for recurring transactions in [`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget).

For Identity:

- Added [`/identity/match`](/docs/api/products/identity/#identitymatch) (beta) to simplify the process of using Identity for account ownership verification. [`/identity/match`](/docs/api/products/identity/#identitymatch) allows you to send user information such as name, address, email, and phone number, and returns a match score indicating how closely that data matches information on file with the financial institution, so that you no longer have to implement similarity matching logic on your own. [`/identity/match`](/docs/api/products/identity/#identitymatch) is currently in closed beta; to request access, contact your Account Manager or [contact Sales](https://www.plaid.com/contact).

For Identity Verification:

- Added [`risk_check` object](https://plaid.com/docs/api/products/identity-verification/#identity_verificationget) to [Identity Verification API responses](/docs/api/products/identity-verification/) to provide more detail about the factors that influenced the risk check result.

For Assets:

- Released Fast Assets functionality to General Availability (GA). With Fast Assets, you can create a Fast Asset Report with partial information (current balance and identity data) in about half the time as a Full Asset Report, then request the Full Asset Report later. Fast Assets is available free of charge to all Assets customers. For more information, see the [Assets API reference](/docs/api/products/assets/#asset_report-create-request-options-add-ons).
- Added the optional `days_to_include` field to [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) requests. This field allows you to optionally reduce the number of days displayed in an Asset Report, to improve usability for human reviewers.
- Added Freddie Mac as an audit copy token partner with `auditor_id: freddie_mac`.
- Added `credit_category` (beta) to transactions data returned by Asset Reports with Insights, describing the category of the transaction. Access to this field is in closed beta; to request access, contact your Account Manager.

For Income:

- Added support for `bank_employment_results` data to [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget). Employment functionality is currently in closed beta; for more details, contact your Account Manager.

For Signal:

- Added `warnings` field to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) response.

For Virtual Accounts:

- Added `transaction_id` to the Payment Initiation `PAYMENT_STATUS_UPDATE` webhook. This field will be present only when a payment was initiated using Virtual Accounts.
- Added `payment_id` and `wallet_id` to the `WALLET_TRANSACTION_STATUS_UPDATE` webhook.

##### February 2023

- Added the `persistent_account_id` field (beta) to the `account` object. The `persistent_account_id` field identifies an account across multiple instances of the same Item, for use with Chase Items only, to simplify the management of Chase [duplicate Item behavior](/docs/link/oauth/#chase). Access to this field is currently in closed beta; for more details, contact your Account Manager.

For Assets:

- Beginning March 31, 2023, Assets will be updated to return `investment`, instead of `brokerage`, as the account type for investment accounts. `brokerage` may still be returned as an account sub-type under the `investment` account type.

For Income:

- Added `employment` value to the `products` array for [`/link/token/create`](/docs/api/link/#linktokencreate). Employment functionality is currently in closed beta; for more details, contact your account manager.

For Identity Verification:

- Added `redacted_at` field in Identity Verification response and Documentary Verification Document component.

For Transfer:

- Added the ability to specify a payment `network` (either `ach` or `same-day-ach`) when calling [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate).

For reseller partners:

- Added `income_verification` and `employment` as supported products for [`/partner/customer/create`](/docs/api/partner/#partnercustomercreate). Employment functionality is currently in closed beta.
- Added the ability to specify `redirect_uris` to [`/partner/customer/create`](/docs/api/partner/#partnercustomercreate).

For Virtual Accounts:

- Increased the minimum length of the `reference` field for [`/wallet/transaction/execute`](/docs/api/products/virtual-accounts/#wallettransactionexecute) to 6 characters.
- To improve usability, added `wallet_id` field to [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) and [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist) responses and to the [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) webhook.

For Payment Initiation:

- Increased the minimum length of the `reference` field for [`/payment_initiation/payment/reverse`](/docs/api/products/payment-initiation/#payment_initiationpaymentreverse) to 6 characters.
- To improve usability, added `transaction_id` field to the [`PAYMENT_STATUS_UPDATE`](/docs/api/products/payment-initiation/#payment_status_update) webhook. This field will be populated when the payment status is `PAYMENT_STATUS_SETTLED`.
- Removed the deprecated `options.wallet_id` field from [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) and [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate).

For Signal:

- Added `signal` to the products array. For most current Signal customers, Signal is automatically enabled for all of their Items; over time, Signal will be moving to a model similar to Plaid's other products, in which customers instead initialize an Item with signal by adding `signal` to the products array when calling [`/link/token/create`](/docs/api/link/#linktokencreate).

##### January 2023

- Released Plaid [React Native SDK 9.0.0](https://github.com/plaid/react-native-plaid-link-sdk) and Plaid [iOS SDK 4.1.0](https://github.com/plaid/plaid-link-ios). All iOS integrations must upgrade to these SDKs by June 30, 2023 in order to maintain support for Chase OAuth flows on iOS. Any integrations using webviews will also be required to update their webview handlers by June 30, 2023. For more details on required webview updates, see the [OAuth Guide](/docs/link/oauth/#extending-webview-instances-to-support-certain-institutions).
- Released improvements to the [returning user experience](/docs/link/returning-user/): removing the requirement to provide a verified time for phone numbers and email addresses, making RUX flows available for more products, and adding Institution Boosting, which automatically surfaces institutions a user has previously linked to Plaid in the Link flow. Enabling the returning user experience can increase conversion by 8% or more. For more details, see [Returning user experience documentation](/docs/link/returning-user/).

For Auth:

- Updated [Same-Day Micro-deposits](/docs/auth/coverage/same-day/) to make a single one-cent micro-deposit, which will not be reversed, rather than multiple micro-deposits. This change will reduce costs and decrease the incentive for micro-deposit fraud. No changes are required on the part of developers to support this change. This change will automatically be rolled out to customers over the next several months.

For Transfer:

- Added `expected_settlement_date` field to the Transfer object.
- Added `funding_account_id` field to clarify which account is used to fund a transfer. This field replaces the older `origination_account_id`.

For Virtual Accounts:

- Added `status` field to Wallet schema.

For Partners:

- Added [`PARTNER_END_CUSTOMER_OAUTH_STATUS_UPDATED`](/docs/api/partner/#partner_end_customer_oauth_status_updated) webhook.

##### December 2022

For Auth:

- Announced upcoming [Instant Match](/docs/auth/coverage/instant/#instant-match) automatic enrollment to existing customers who have not yet been enabled for Instant Match by default. Instant Match is a higher-converting experience for Link that expands Auth coverage to more end users. All customers will be enabled for Instant Match by default unless they opt-out by January 19, 2023, which they can do by contacting their Account Manager.

For Assets:

- Released relay tokens to beta, along with associated endpoints [`/credit/relay/create`](/docs/api/products/assets/#creditrelaycreate), [`/credit/relay/get`](/docs/api/products/assets/#creditrelayget), [`/credit/relay/refresh`](/docs/api/products/assets/#creditrelayrefresh), and
  [`/credit/relay/remove`](/docs/api/products/assets/#creditrelayremove). Relay tokens support sharing of Asset Reports with customers' authorized service providers, such as underwriters.

For Income:

- For [`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget), deprecated `amount`, `iso_currency_code`, and `unofficial_currency_code` at top levels in response in favor of new `total_amounts` field. This change enables more accurate reporting of income for users with income in multiple different currencies.
- Added new possible values to `rate_of_pay` field returned by [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget): `WEEKLY`, `BI_WEEKLY`, `MONTHLY`, `SEMI_MONTHLY`, `DAILY`, and `COMMISSION`.
- Added new `pay_basis` field to the response of [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget).

For Virtual Accounts:

- Added `FUNDS_SWEEP` as a possible transaction type.

For Transfer:

- Added [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget) endpoint to provide information about which Items support Real Time Payments (RTP)
- Added endpoints /transfer/originator/get, /transfer/originator/list, /transfer/originator/create, and /transfer/questionnaire/create to support marketplaces and reseller partners in creating and managing transfer originators.
- Added the ability to create and manage [recurring transfers](/docs/transfer/recurring-transfers/). API changes include the new endpoints `/transfer_recurring/create`, [`/transfer/recurring/cancel`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcancel), [`/transfer/recurring/get`](/docs/api/products/transfer/recurring-transfers/#transferrecurringget), and [`/transfer/recurring/list`](/docs/api/products/transfer/recurring-transfers/#transferrecurringlist), as well as the new webhooks `RECURRING_NEW_TRANSFER`, `RECURRING_TRANSFER_SKIPPED`, and `RECURRING_CANCELLED`, and the test endpoints [`/sandbox/transfer/test_clock/create`](/docs/api/sandbox/#sandboxtransfertest_clockcreate), [`/sandbox/transfer/test_clock/advance`](/docs/api/sandbox/#sandboxtransfertest_clockadvance), [`/sandbox/transfer/test_clock/get`](/docs/api/sandbox/#sandboxtransfertest_clockget), and [`/sandbox/transfer/test_clock/list`](/docs/api/sandbox/#sandboxtransfertest_clocklist).

For Enrich:

- Released Enrich to Early Availability. Enrich access is now available by default in Sandbox. For details on requesting access to use Enrich with real data, see the [Enrich documentation](https://plaid.com/docs/enrich/).
- Added support for legacy categories, using categories from [`/categories/get`](/docs/api/products/transactions/#categoriesget), in addition to the newer `personal_finance_category` schema.
- Renamed counterparty type `delivery_marketplace` to `marketplace` and added counterparty type `payment_terminal`.

For the Reseller Partner API:

- Added [`/partner/customer/oauth_institutions/get`](/docs/api/partner/#partnercustomeroauth_institutionsget) endpoint to provide OAuth-institution registration information to Reseller Partners.

##### November 2022

- Released [Signal](/docs/signal/) to general availability. Signal uses machine learning techniques to evaluate a proposed ACH transaction and determine the likelihood that the transaction will be reversed. It also provides fields and insights that you can incorporate into your own data models. By using the Signal API, you can release funds earlier to optimize your user experience while managing the risk of ACH returns.

For Auth:

- Enabled [Instant Match](/docs/auth/coverage/instant/#instant-match) by default for most existing customers.

For Investments:

- Notified customers that customers must hold a CGS (CUSIP Global Services) license to obtain CUSIP and ISIN data. Beginning in mid-September 2023, customers who do not have a record of this license with Plaid will receive `null` data in these fields. To maintain access to these fields, contact your Plaid Account Manager or [investments-vendors@plaid.com](mailto:investments-vendors@plaid.com).

For Payment Initiation:

- Added the ability to initiate partial refunds by specifying an `amount` when calling [`/payment_initiation/payment/reverse`](/docs/api/products/payment-initiation/#payment_initiationpaymentreverse). As part of this change, added `amount_refunded` field to [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) and [`/payment_initiation/payment/list`](/docs/api/products/payment-initiation/#payment_initiationpaymentlist).
- Added support for local payment schemes for all supported currencies (previously, only EUR and GBP local payment schemes been supported), and added the enum values `LOCAL_DEFAULT` and `LOCAL_INSTANT` to represent them. In the UK, the `FASTER_PAYMENTS` enum value has been replaced by `LOCAL_DEFAULT`.
- Removed support for currencies CHF and CZK.

For Virtual Accounts:

- For improved consistency, renamed `/wallet/transactions/list` to [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist).
- For clarity, renamed `start_date` to `start_time` in the [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist) response to reflect that the field is a date-time.

For Transfer:

- Made `account_id` nullable in the responses for [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) and [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate).
- Added `deleted_at` field to /payment\_profile/get response.
- Added `refunds` field to the transfer object.
- Added `refund_id` to the transfer event object.

For Reseller Partners:

- Added the [`/partner/customer/remove`](/docs/api/partner/#partnercustomerremove) endpoint to disable customers who have not yet been enabled in Production.

##### October 2022

- Added support for additional Link display languages: Danish, German, Estonian, Italian, Latvian, Lithuanian, Norwegian, Polish, Romanian, and Swedish.
- Added support for additional countries: Denmark, Estonia, Latvia, Lithuania, Poland, Norway, and Sweden.
- Added support for `USER_INPUT_TIMEOUT` as a value for `force_error` in the [Sandbox custom user](/docs/sandbox/user-custom/) schema.

For Auth:

- Enabled [Instant Match](/docs/auth/coverage/instant/#instant-match) by default for all new customers.

For Investments:

- Added `non-custodial wallet` to account subtypes.
- Added `trade` investment transaction subtype as a subtype of `transfer` investment transaction type

For Income:

- To improve reliability and developer ease of use, modified the multi-Item Link flow for Bank Income to use the new [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) endpoint rather than relying on Link events.
- Added Bank Income status `USER_REPORTED_NO_INCOME`.
- Added Income support to [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate).
- Deprecated `VERIFICATION_STATUS_PENDING_APPROVAL`.
- Established a 128 character limit for the `client_user_id` field in the [`/user/create`](/docs/api/users/#usercreate) request.
- Added `institution_name` and `institution_id` fields to [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) response.

For Virtual Accounts:

- Added `options.start_time` and `options.end_time` to [`/wallet/transaction/list`](/docs/api/products/virtual-accounts/#wallettransactionlist) request.
- Added `last_status_update` and `payment_id` field to wallet transactions.

For Payment Initiation:

- Added `transaction_id` field to the payment object.

For Transfer (beta):

- Added support for RTP networks
- Added decision rationale code `PAYMENT_PROFILE_LOGIN_REQUIRED` and the ability to update a Payment Profile via update mode. Also added /sandbox/payment\_profile/reset\_login to test this new update mode flow in Sandbox.
- To improve consistency, renamed `LOGIN_REQUIRED` decision rationale to `ITEM_LOGIN_REQUIRED`.
- Deprecated `origination_account_id` from [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) endpoint.
- Added new `originator_client_id` field to support Third-Party Senders (TPS).

For Wallet Onboard (beta):

- Released Wallet Onboard to beta.

For Partner Resellers:

- Added [Partner reseller endpoints](/docs/api/partner/) to improve the experience of onboarding new customers for delegated diligence partners.

##### September 2022

- Added the ability to simulate the `USER_INPUT_TIMEOUT` error in [Sandbox](/docs/sandbox/test-credentials/#error-testing-credentials).
- Added the ability to specify account mask when using [custom Sandbox users](/docs/sandbox/user-custom/).

For Investments:

- Added support for crypto wallet investment account subtype `non-custodial wallet` and crypto `trade` investment transaction type.
- Made `institution_price_as_of` nullable.

For Transfer (beta):

- Added support for client-side beacons. A beacon is now required when using Guarantee with web checkout flows.
- Removed the `idempotency_key` from the [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) response.
- Added `settled` and `swept_settled` transfer events and event statuses.
- Added `standard_return_window` and `unauthorized_return_window` fields to the Transfer object.

For Payment Initiation:

- Added `AUTHORISING` wallet status.
- Added `recipient_id` to the [`/wallet/create`](/docs/api/products/virtual-accounts/#walletcreate), [`/wallet/get`](/docs/api/products/virtual-accounts/#walletget), and [`/wallet/list`](/docs/api/products/virtual-accounts/#walletlist) responses.

For Assets:

- Added `days_to_include` and `options` to [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) request.

##### August 2022

- Released [Identity Verification](/docs/identity-verification/) and [Monitor](/docs/monitor/) to General Availability (GA).
- Added `issuing_region` as a field in the extracted data for Identity Verification documents
- Extended support for Android Link SDK versions prior to 3.5.0, iOS Link SDK versions prior to 2.2.2, and Link React Native SDK versions prior to 7.1.1. Previously, these SDK versions had been scheduled to be sunset on November 1, 2022. The sunset has been canceled and Plaid will now continue to support these versions past November 2022. We still recommend you use the latest SDK versions.

For Payment Initiation:

- Added support for [virtual accounts](/docs/payment-initiation/virtual-accounts/).
- Added support for additional currencies PLN, SEK, DKK, NOK, CHF, and CZK.
- Added `PAYMENT_STATUS_SETTLED` payment status.

For Income:

- Removed verification fields from [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) and /income/verification/paystubs/get
- Removed `pull_id` field from [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget)
- Added payroll institution to /credit/income/precheck

##### July 2022

For [`/link/token/create`](/docs/api/link/#linktokencreate) `user` object:

- Deprecated the unused fields `ssn`, `date_of_birth`, and `legal_name`.
- Added the `name`, `address`, and `id_number` parameters to support Identity Verification.

For Payment Initiation:

- Added `bin` as a field under `institution_metadata` for [`/link/token/create`](/docs/api/link/#linktokencreate).

For Income:

- Add `stated_income_sources` as a field under `income_verification` for [`/link/token/create`](/docs/api/link/#linktokencreate).
- Add 1099 data to [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) response.
- Deprecate `paystubs.verification` in [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) response.

For Transfer:

- Add `payment_profile_id` to [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate), [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate), and [`/link/token/create`](/docs/api/link/#linktokencreate) and make `account_number` and `access_token` optional for [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) and [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate). These changes will support forthcoming Transfer functionality.
- Add `user_present` a required field for [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) when using Guarantee (beta)
- Deprecated the no-longer-used `swept` and `reverse_swept` statuses.

##### June 2022

For Auth:

- Launched [Auth Type Select](/docs/auth/coverage/flow-options/#adding-manual-verification-entry-points-with-auth-type-select) (formerly known in beta as Flexible Auth or Flex Auth) to all customers. Auth Type Select allows you to optionally route end users directly to micro-deposit-based verification, even if they are eligible for Instant Match or Instant Auth.

For Transfer:

- Made Plaid Guarantee available in beta. Guarantee allows you to guarantee the settlement of an ACH transaction, protecting you against fraud and clawbacks.
- Added `TRANSFER_LIMIT_REACHED` to Transfer authorization decision rationale codes.

For Income:

- Added the `accounts` object with `rate_of_pay` data to the response for [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget).

For Payment Initiation (UK and Europe only):

- Launched payment consents, which can be used to initiate payments on behalf of the user.

##### May 2022

- Added [Identity Verification](/docs/identity-verification/) and [Monitor](/docs/monitor/) products in Early Access. Identity Verification checks the identity of users against identity databases and user-submitted ID documents, while Monitor checks user identities against government watchlists.

For Transactions:

- Added [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), which improves ease of use for handling transactions updates. While [`/transactions/get`](/docs/api/products/transactions/#transactionsget) continues to be supported and there are no plans to discontinue it, all new and existing integrations are encouraged to use [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) instead of the older [`/transactions/get`](/docs/api/products/transactions/#transactionsget) endpoint.

For Income:

- Added `employee_type` and `last_paystub_date` to [`/credit/employment/get`](/docs/api/products/income/#creditemploymentget) response.
- Removed `uploaded`, `created` and `APPROVAL_STATUS_APPROVED` enum strings, as these are no longer used.

##### April 2022

- Added the ability to default Link to highlighting a specific institution when launching Link, via the [`institution_data`](/docs/api/link/#link-token-create-request-institution-data) request field.
- Launched [Income](/docs/income/) to general availability. Income allows you to verify the income of end users for purposes of loan underwriting. Various updates were also made to Income interfaces prior to launch; Income beta customers have been contacted by their account managers with details on the differences between the beta API and the released API.

For Transactions:

- Added [`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget), which provides information about recurring transaction streams that can be used to help your users better manage cash flow, reduce spending, and stay on track with bill payments. This endpoint is not included by default as part of Transactions; to request access to this endpoint and learn more about pricing, contact your Plaid account manager.

For Auth:

- Added Highnote as a processor partner. For a full list of all Auth processor partners, see [Auth Payment Partners](/docs/auth/partnerships/).

##### March 2022

- Introduced the Institution Select shortcut, which enables you to highlight a matching institution on the Institution Select pane.
- Added `institution_data` request field to [`/link/token/create`](/docs/api/link/#linktokencreate) endpoint, which accepts `routing_number`.
- Added `match_reason` to metadata field for `MATCHED_SELECT_INSTITUTION`, which indicates whether `routing_number` or `returning_user` resulted in a matched institution.
- Added support for `AUTH_UPDATE` and `DEFAULT_UPDATE` webhooks to [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook). Also added `webhook_type` parameter to this endpoint to support different `DEFAULT_UPDATE` webhooks for Transactions, Liabilities, and Investments.

For Income:

- New set of API endpoints added and numerous updates made for General Availability release. Existing beta endpoints are still supported but marked as deprecated. For more details, see the [Income docs](/docs/income/) and [Income API reference](/docs/api/products/income/).

For Auth:

- Added Apex Clearing, Checkout.com, Marqeta, and Solid as processor partners.

For Payment Initiation:

- Added support for the `IT` country code.
- Added support for searching by `consent_id` to [`/institutions/search`](/docs/api/institutions/#institutionssearch).

For Bank Transfer (beta):

- Added `wire_routing_number` parameter to [`/bank_transfer/migrate_account`](/docs/bank-transfers/reference/#bank_transfermigrate_account).

For Transfer (beta):

- Removed `permitted` decision for [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate).
- Added [`/sandbox/transfer/fire_webhook`](/docs/api/sandbox/#sandboxtransferfire_webhook) endpoint.

##### February 2022

- For Transactions, released new personal finance categories to provide more intuitive and usable transaction categorization. Personal finance categories can now be accessed by calling [`/transactions/get`](/docs/api/products/transactions/#transactionsget) with the `include_personal_finance_category` option enabled.
- For Income, removed several unused fields and endpoints, including removing `income_verification_id` from [`/sandbox/income/fire_webhook`](/docs/api/sandbox/#sandboxincomefire_webhook) and removing the /income/verification/summary/get and /income/verification/paystub/get endpoints.

For Transfer (beta):

- Deprecated the `idempotency_key` parameter, as idempotency is now tracked via other identifiers.
- Made `repayment_id` a required parameter for /transfer/repayment/return/list.
- Made guaranteed fields required in Transfer endpoints.

##### January 2022

- Added the ability to test the `NEW_ACCOUNTS_AVAILABLE` webhook via [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook).
- Updated [`/item/webhook/update`](/docs/api/items/#itemwebhookupdate) to accept empty or `null` webhook URLs.
- Updated institutions endpoints to accept null values as input for optional input fields.

For Transfer (beta):

- Added publicly documented sweep endpoints to provide visibility into the status of Transfer sweeps.
- Added `iso_currency_code` throughout the API to future-proof for multi-currency support (currently only USD is supported).
- Made `repayment_id` required in /transfer/repayment/return/list endpoint

For Bank Transfer (beta):

- Removed `receiver_pending` and `receiver_posted` from bank transfer event types, and removed receiver details from events.

For Payment Initiation:

- Added payment scheme support for EU payments.
- Removed `scheme_automatic_downgrade` from [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate).
- Updated webhooks to use new statuses.

For Income:

- Added `DOCUMENT_TYPE_NONE` value for document type.
- Made employer address fields no longer required in /income/verification/precheck.

##### December 2021

- For Transfer, updated the schema definitions for [`/transfer/intent/get`](/docs/api/products/transfer/account-linking/#transferintentget) and [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate).
- For Income, deprecated the status `VERIFICATION_STATUS_DOCUMENT_REJECTED`.
- For Payment Initiation, added several new statuses, including `PAYMENT_STATUS_REJECTED`.

##### November 2021

- For Payment Initiation, added the new status `PAYMENT_STATUS_EXECUTED` and renamed `emi_account_id` to `wallet_id`.
- For [`/asset_report/get`](/docs/api/products/assets/#asset_reportget), added the fields `merchant_name` and `check_number` to the Transactions schema, to match the Transactions schema already being used by the Transactions API.
- For Income, added the new status `VERIFICATION_STATUS_PENDING_APPROVAL`.

##### October 2021

Multiple changes to the Income API:

- Added the ability to verify submitted paystubs against a user's transaction history by adding the `income_verification.access_tokens` parameter to [`/link/token/create`](/docs/api/link/#linktokencreate) and updating /income/verification/paystubs/get to return the income verification status for each paystub.
- Added a `precheck_id` field to /income/verification/create and an `income_verification.precheck_id` field to [`/link/token/create`](/docs/api/link/#linktokencreate) to fully support use of the /income/verification/precheck endpoint to check whether a given user is supportable by the Income product.
- Extensive changes to the paystub schema returned by /income/verification/paystubs/get, including adding new fields and deprecating old ones. For details, contact your Plaid Account Manager.
- Officially deprecated the `income_verification_id` in favor of the Link token-based income verification flow.
- Added an `item_id` field to the [`income_verification` webhook](/docs/api/products/income/#income_verification) and corresponding [`/sandbox/income/fire_webhook`](/docs/api/sandbox/#sandboxincomefire_webhook) endpoint.
- Added `employer.url` and `employer.address` fields to the /income/verification/precheck endpoint.
- Added a `doc_id` field to /income/verification/taxforms/get.

Other changes:

- For institutions endpoints, marked the `status` enum as deprecated in favor of the more detailed `breakdown` object.
- Added `DE` as a supported country to support Payment Initiation use cases.

##### September 2021

- Released Account Select v2, including the new [`NEW_ACCOUNTS_AVAILABLE`](/docs/api/items/#new_accounts_available) webhook, for improved end-user sharing controls. All implementations must migrate to Account Select v2 by March 2022.
- Added the ability to check which types of [Auth coverage](/docs/auth/coverage/) an institution supports via a new `auth_metadata` object now available from [Institutions endpoints](/docs/api/institutions/)
- Added the /income/verification/precheck endpoint to check whether a given user is supportable by the Income product.
- Added `initiated_refunds` field to [`/payment_initiation/payment/get`](/docs/api/products/payment-initiation/#payment_initiationpaymentget) and [`/payment_initiation/payment/list`](/docs/api/products/payment-initiation/#payment_initiationpaymentlist) to show refunds associated with payments.
- Added fields `include_personal_finance_category_beta` and `personal_finance_category_beta` to Transactions endpoints as part of a beta for new Transactions categorizations. To request beta access, contact [transactions-feedback@plaid.com](mailto:transactions-feedback@plaid.com)
- Removed some fields (`direction`, `custom_tag`, `iso_currency_code`, `receiver_details`) and receiver event types from [Bank Transfer (beta) endpoints](/docs/api/products/transfer/) and made `ach_class` a required field.
- Added support for the currency type Bitcoin SV.

##### August 2021

- Released [UX improvements to Link](https://plaid.com/blog/enhancements-to-link-for-an-improved-user-experience/).
- Launched [Plaid OpenAPI schema](https://github.com/plaid/plaid-openapi) and new, [auto-generated client libraries](https://plaid.com/blog/open-api-client-libraries/) for Python, Node, Java, Ruby, and Go to General Availability.
- Added the ability to use [custom Sandbox data](/docs/sandbox/user-custom/) with the Investments product.
- For [`/transactions/get`](/docs/api/products/transactions/#transactionsget), added the `check_number` field.
- Updated the list of [ACH partners](/docs/auth/partnerships/), adding Treasury Prime.
- For [`/processor/balance/get`](/docs/api/processor-partners/#processorbalanceget), added `min_last_updated_datetime` option to request.
- Multiple schema and endpoint updates for the [Bank Transfers (beta)](https://plaid.com/docs/bank-transfers/) and [Income (beta)](https://plaid.com/docs/income/) products. Beta participants can receive details on updates from their account managers.

##### July 2021

- Added new webhook for Deposit Switch.
- Added optional `country_code` and `options` parameters to deposit switch creation endpoints, as well as new values for the `state` enum in the response.
- For Deposit Switch, added additional response fields `employer_name`, `employer_id`, `institution_name`, `institution_id`, and `switch_method`.
- Updated the list of [ACH partners](/docs/auth/partnerships/), including adding Alpaca, Astra, and Moov.
- For [`/transactions/get`](/docs/api/products/transactions/#transactionsget), added the fields `include_original_description` and `original_description` in the request and response, respectively.

##### June 2021

- Chase now supports real-time payments and same-day ACH for OAuth-based connections.
- Added new Investment account subtypes `life insurance`, `other annuity`, and `other insurance`.
- Added new error codes [`INSTITUTION_NOT_ENABLED_IN_ENVIRONMENT`](/docs/errors/institution/#institution_not_enabled_in_environment), `INSTITUTION_NOT_FOUND`, and [`PRODUCT_NOT_ENABLED`](/docs/errors/item/#product_not_enabled). These codes replace more generic errors that could appear when working with OAuth institutions or API-based institution connections.
- Began rolling out improvements to reduce balance latency.
- Released a new representation of cryptocurrency, to represent crypto holdings more similarly to investments, rather than foreign currency, in order to match how most popular institutions represent crypto. `institution_price` will now reflect the USD price of the currency, as reported by the institution, and `institution_value` will reflect the value of the holding in USD. `close_price` and `close_price_as_of` will now be populated. `iso_currency_code` will be set to `USD` and `unofficial_currency_code` will be set to `null`.

##### May 2021

- Improved the UI for [Instant Match](/docs/auth/coverage/instant/#instant-match) and enabled it for all "Pay as you go" customers. Customers with a monthly contract who would like to use Instant Match should contact their Plaid Account Manager.
- Removed the requirement to be enabled by Support in order to request an [Asset Report with Insights](/docs/api/products/assets/#asset_reportget).
- For the Payment Initiation product, launched [Modular Link](https://plaid.com/blog/improving-the-payments-experience), allowing further UI customization.
- For the Payment Initiation product, added an `options` object to [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) to support restricting payments originating from specific accounts.
- For the Liabilities product, added a [`DEFAULT_UPDATE` webhook](/docs/api/products/liabilities/#default_update) to indicate changes to liabilities accounts.
- Made [Item Debugger and improved Institution Status](https://plaid.com/blog/item-debugger-institution-status/) available to 100% of developers.
- Added Link [Consent pane customization](/docs/link/customization/#consent-pane-customizations).
- Added the ability to use wildcards to specify a subdomain in an OAuth redirect URI.
- Clarified that the following Link onEvent callbacks are stable: `OPEN`, `EXIT`, `HANDOFF`, `SELECT_INSTITUTION`, `ERROR`, but the remaining are informational.

##### April 2021

- In order to support the beta generated client libraries, added the ability to provide `client_id` and `secret` [via headers](/docs/api/#api-protocols-and-headers) instead of as request parameters.
- Added the `min_last_updated_datetime` parameter to [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) to handle institutions that do not always provide real-time balances, and added the corresponding `LAST_UPDATED_DATETIME_OUT_OF_RANGE` error.
- Removed `last_statement_balance` from the official documentation for [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget), as it was not ever returned.
- Fixed the case of submitting an invalid client ID to return an [`INVALID_API_KEYS`](/docs/errors/invalid-input/#invalid_api_keys) error instead of an `INTERNAL_SERVER_ERROR`, in order to match documented behavior.

##### March 2021

- Launched [Plaid Income (beta)](https://plaid.com/income/) for verifying income and employment.
- Added the ability to specify an `end_date` for standing orders in the [Payment Initiation](/docs/payment-initiation/) product.
- Updated list of active processor partners.
- Added [webhook](/docs/bank-transfers/webhook-events/) to notify of new Bank Transfers events.
- Added [`/sandbox/bank_transfer/fire_webhook`](/docs/bank-transfers/reference/#sandboxbank_transferfire_webhook) endpoint to test Bank Transfers webhooks.
- Added `auth_flow` parameter to [`/link/token/create`](/docs/api/link/#linktokencreate) to support Flexible Auth (beta).
- Updated the [`/payment_initiation/recipient/create`](/docs/api/products/payment-initiation/#payment_initiationrecipientcreate) endpoint to accept `address.country_code` (preferred) instead of `address.country` (accepted, but deprecated).
- Removed the `MATCHED_INSTITUTION_SELECT` transition view event, after consolidating the returning user institution select view.
- Added the [`SELECT_BRAND`](/docs/link/web/#link-web-onevent-eventName-SELECT-BRAND) onEvent callback, after shipping a change that groups institution login portals within the same institution brand.
- Stopped sending `IS_MATCHED_USER` and `IS_MATCHED_USER_UI` because these events are duplicates. You should use `MATCHED_CONSENT` and `MATCHED_SELECT_INSTITUTION` to identify when a returning user is recognized and chooses an institution that we pre-matched.

##### February 2021

- Added additional [payment error codes](/docs/errors/payment/).
- Added the UK-only fields [`authorized_datetime`](/docs/api/products/transactions/#transactions-get-response-transactions-authorized-datetime) and [`datetime`](/docs/api/products/transactions/#transactions-get-response-transactions-datetime) fields to the `transaction` object, for more detailed information on when the transaction occurred.
- Added `update_type` field to the `item` object. This field will be used to support upcoming connectivity improvements to accounts with multi-factor authentication.
- Added optional `ssn` and `date_of_birth` fields to the `user` object in [`/link/token/create`](/docs/api/link/#linktokencreate) to support upcoming functionality enhancements.
- Released an [OpenAPI file](https://github.com/plaid/plaid-openapi) describing the Plaid API.

##### January 2021

- Launched Deposit Switch (beta) for transferring direct deposits from one account to another.
- Improved error handling to treat some invalid input errors that were previously treated as 500 errors as `INVALID_INPUT` errors instead.

##### December 2020

- Made Bank Transfers (beta) available for testing in the Development environment.
- Added `investments_updates` to the `status` object returned by Institutions endpoints.

##### November 2020

- Removed some internal-only, pre-beta products from appearing in the [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) response.
- Restored the `/item/public_token/create` and `/payment_initiation/payment/token/create` endpoints to API version 2020-09-14 to avoid disrupting update mode for end users on older mobile SDKs that do not support Link tokens. These endpoints are still deprecated, and it is recommended you update mobile apps to the latest SDKs as soon as possible.

##### October 2020

- Released API version 2020-09-14. See the [version changelog](/docs/api/versioning/#version-2020-09-14) for details.
- Added support for mortgages to [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget).
- Released [Bank Transfers](/docs/bank-transfers/) to closed beta.
- To improve consistency with other error types, changed the error type for Bank Transfers errors from `BANK_TRANSFER` to `BANK_TRANSFER_ERROR`.
- Added the fields `user.phone_number`, `user_email.address`, `user.phone_number_verified_time` and `user.email_address_verified_time` to [`/link/token/create`](/docs/api/link/#linktokencreate) to support the new [returning user experience](/docs/link/returning-user/), which allows users a more seamless Link flow.

##### August 2020

- Introduced [standing orders](https://blog.plaid.com/recurring-payments-with-standing-orders/) to [Payment Initiation](https://plaid.com/uk/products/payment-initiation/) for our UK and Europe customers. If applicable, your users will now be able to make recurring scheduled payments with a single authorization.
- Expanded access to [full Auth coverage](/docs/auth/coverage/) to more developers. If you would like to use micro-deposit based verification and don't have access to it, contact your Plaid Account Manager.
- Updated Link error messages to provide more actionable instructions for end users, making resolution troubleshooting easier.
- Made [account filtering](https://blog.plaid.com/improvements-to-plaid-link-with-account-filtering-and-updated-error-messaging/) available across all our products, so you can configure the Link flow to guide end users in selecting relevant institutions and accounts.

##### July 2020

- Added the [ITEM: USER\_PERMISSION\_REVOKED](https://plaid.com/docs/api/items/#user_permission_revoked) webhook, which will notify you when a user contacts Plaid directly to delete their data or change their data sharing preferences with your app.
- Released Link tokens, the new preferred way to use Link, replacing the public key. To learn how to migrate your application to use Link tokens, see the [Link token migration guide](https://plaid.com/docs/link/link-token-migration-guide/).

##### June 2020

- Added a new `merchant_name` field to the [`/transactions/get`](/docs/api/products/transactions/#transactionsget) endpoint for the US and Canada, providing clearer and more consistent merchant names for 95% of existing transactions.
- Added the [PAYMENT\_INITIATION:PAYMENT\_STATUS\_UPDATE](https://plaid.com/docs/api/products/payment-initiation/#payment_status_update) webhook, which pushes instant notifications when payment status changes for the [UK Payment Initiation product](https://plaid.com/uk/products/payment-initiation/).
- Added the ability to create payment recipients via sort codes and account numbers, not just IBANs.

##### May 2020

- We launched a new open finance platform called [Plaid Exchange](https://blog.plaid.com/introducing-plaid-exchange/) that enables financial institutions to build a consumer-permissioned data access strategy and strengthen the ability of end users to reliably access their financial data.

##### April 2020

- Launched [Payment Initiation](https://plaid.com/uk/products/payment-initiation/) in the UK, which offers an easy way for users to fund their accounts, make purchases, and pay invoices—all from their favorite apps or websites.

##### March 2020

- Added the ability to [filter by `account_subtype`](https://plaid.com/docs/api/link/#linktokencreate), allowing you to further optimize the Link flow.
- Added the `HOLDINGS: DEFAULT_UPDATE` and `INVESTMENT_TRANSACTIONS: DEFAULT_UPDATE` webhooks, which will fire each time data has successfully refreshed for Holdings and Investments Transactions.

##### February 2020

- Added [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh), which enables you to pull a user’s transactions on demand.
- Added [`/webhook_verification_key/get`](/docs/api/webhooks/webhook-verification/#get-webhook-verification-key), which allows you to verify the authenticity of incoming webhooks.

##### January 2020

- Added `store_number`, `authorized_date`, and `payment_channel` to the [`/transactions/get`](/docs/api/products/transactions/#transactionsget) response.
- Added the `investment_transactions.subtypes` field to provide more granular detail about the tax, performance, and fee impact of investments transactions.

##### November 2019

- Added the `status.investments.last_successful_update` and `status.investments.last_failed_update` fields to the data returned by [`/item/get`](/docs/api/items/#itemget).
- Launched official support for Link on React Native with a new [SDK](https://github.com/plaid/react-native-plaid-link-sdk), bringing unified support to React Native apps.

##### October 2019

- Added the `status.item_logins.breakdown` data to [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) and the Developer Dashboard.
- Added the `routing_numbers` field to the Institutions object. You can also filter institutions via the `options.routing_numbers` field in each Institutions API endpoint.

##### September 2019

- Added support for [credit card details](https://blog.plaid.com/liabilities-credit-cards) to the Liabilities product.
- Added Canada-specific account subtypes, including RRSP, RESP, and TFSA, to the Investments product.

##### August 2019

- Among numerous other improvements to Liabilities, such as expanded data access, added a new `loan_type`: `cancelled` to the Liabilities product.

##### July 2019

- We launched [Liabilities](https://blog.plaid.com/liabilities/), which enables developers to access a feed of standardized student loan details from the largest U.S. servicers including Navient, Nelnet, FedLoan, Great Lakes, and many more.

##### June 2019

- Launched [Investments](https://blog.plaid.com/investments/), which allows developers, fintech apps, and financial institutions to create a holistic view of their customers’ investments.
- Added the `status.transactions_updates` field, exposing Transactions health to both the [`/institutions/get`](/docs/api/institutions/#institutionsget) endpoint and the Dashboard.

##### May 2019

- We enhanced the [`/identity/get`](/docs/api/products/identity/#identityget) endpoint to now return [account-level identity information](https://plaid.com/docs/api/products/identity/#identityget) within the `accounts` object where available.
- We released [API version 2019.05.29](https://plaid.com/docs/api/versioning/#version-2019-05-29) to enable European institution coverage and provide nomenclature and schema updates required by Identity and future products.

##### March 2019

- We updated the [institutions endpoints](https://plaid.com/docs/api/institutions/) so you can now retrieve bank logos, colors, and website URLs to use to customize your in-app experience.
- We enabled triggering and testing of webhooks on demand via the new [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook) endpoint.

##### February 2019

- We launched [new features for Auth](https://blog.plaid.com/new-auth/), enabling developers to authenticate accounts from any bank or credit union in the U.S. Link automatically guides end-users to the best way to authenticate their account based on the bank they select.

##### November 2018

- Added a new [Insights](https://plaid.com/docs/api/products/assets/#asset_reportget) feature, which provides cleaned and categorized transaction data in an Asset Report. In addition to transaction categories, lenders will be able to retrieve merchant names and locations for transactions to use in building risk models.
- Improved the Link experience by informing users about connectivity issues with banks before connecting their account. When banks are experiencing significant issues, users will temporarily be directed to connect their account at a different bank to reduce frustration and drop-off during the onboarding process.

##### September 2018

- We added the [`/asset_report/refresh`](/docs/api/products/assets/#asset_reportrefresh) endpoint, so you can create a new Asset Report with the latest account balances and transactions for a borrower, based on the old report.

##### August 2018

- We added [account filtering](https://plaid.com/docs/api/products/assets/#asset_reportfilter), which gives you the ability to exclude unnecessary accounts from appearing in an Asset Report.

##### June 2018

- Added historical account balances to the PDF version of Asset Reports, bringing them closer in line with the core JSON endpoint.

##### May 2018

- We released the [2018-05-22 version](/docs/api/versioning/#version-2018-05-22) of the Plaid API.
- We enabled distinct API secrets [to now be be set](https://blog.plaid.com/api-secrets/) for the Sandbox, Development, and Production environments.
- Added the [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) endpoint, which enables the creation of Items in the Sandbox environment directly via the API (without Link).

##### April 2018

- Officially launched the [Assets](https://plaid.com/products/assets/) product. Assets is approved by Fannie Mae for [Day 1 Certainty™](https://www.fanniemae.com/singlefamily/day-1-certainty) asset verification.

##### March 2018

- Rolled out several Assets features, including webhooks for Asset Report generation and full support for Audit Copy token generation and Fannie Mae’s Day 1 Certainty™ program, while improving existing features like pending transaction support in PDF reports.

##### February 2018

- Released Assets as a beta product.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
