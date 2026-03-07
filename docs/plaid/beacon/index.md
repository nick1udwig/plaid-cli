---
title: "Beacon - Introduction to Beacon | Plaid Docs"
source_url: "https://plaid.com/docs/beacon/"
scraped_at: "2026-03-07T22:04:46+00:00"
---

# Introduction to Beacon (beta)

#### Learn how to fight fraud with the Plaid Beacon network

[#### API Reference

View Beacon requests, responses, and example code

View Beacon API](/docs/api/products/beacon/)

#### API Reference

View Beacon requests, responses, and example code

[View Beacon API](/docs/api/products/beacon/)

#### Overview

The Beacon beta program is now closed. Customers who are not already using Beacon should instead use the newer [Plaid Protect](https://plaid.com/products/protect/), which incorporates Beacon's capabilities and enhances them with additional fraud insights.

Beacon (US only) helps prevent fraud by detecting data associated with stolen or synthetic identities, account takeover fraud, or breaches. You can create blocklists based on data reported by the Beacon network, or based on first-party fraud attempts on your own app. Beacon is available free of charge.

An optional paid feature, Beacon Account Insights, is available to customers who want to build their own risk analysis on top of Plaid's Beacon data. Beacon Account Insights includes risk-related details such as account age, recent PII changes on the account, and recent failed login attempts. For a full list of data reported by Bank Account Insights, see the [`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget) response schema.

#### Integration setup

Select group for content switcher

##### Initial setup

1. [Create a Beacon program](https://dashboard.plaid.com/sandbox/beacon/programs/new) via the Dashboard.
   - For best results, disable "auto flag reported user" under the "data breach hits" section, unless you have a specific business reason to enable it. Because so many users have been exposed to data breaches, enabling auto-flagging for data breach hits will result in a large number of users entering the review queue. All other automatically-selected options should be left enabled for most use cases.
2. Take note of the beacon program ID, which is the portion of the URL after the `/` when viewing the program in the Dashboard, e.g. `becprg_7Fn5XcPhXnJJyU`. You will need this id when calling [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate) and other endpoints.
3. (Optional) For best results, backfill six months of existing data into Beacon. This should include both known-good and known-fraudulent users. To backfill, call [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate) for each user, making sure to include any known instances of fraud in the `report` object.

Make sure to always use the same `client_user_id` when referring to the same end user, whenever you call [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate), [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate), or [`/link/token/create`](/docs/api/link/#linktokencreate).

##### User creation

1. To create a new user, call [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate).
2. (Optional) To be alerted for the accidental creation of potential duplicate users, listen for the [`DUPLICATE_DETECTED`](/docs/api/products/beacon/#duplicate_detected) webhook. You can investigate potential duplicates via the Dashboard (when configuring your Beacon template, you can set all detected duplicates to enter the review queue by default) or by calling [`/beacon/duplicate/get`](/docs/api/products/beacon/#beaconduplicateget).

##### Reviewing and reporting fraud signals

1. Review any Beacon hits via the Dashboard, to clear or reject the user.
2. If you become aware of any fraud committed by a user, report it to the network by clicking the "Report Fraud" button in the Beacon Dashboard or by calling [`/beacon/report/create`](/docs/api/products/beacon/#beaconreportcreate). To learn more about how to properly categorize fraud, see the [Fraud Reporting Guide](https://view-su2.highspot.com/viewer/3ba173969f637a61e6e9d8e41eba083b).
3. (Optional) To monitor for fraud on an ongoing basis, listen for the [`REPORT_SYNDICATION_CREATED`](/docs/api/products/beacon/#report_syndication_created) webhook to be alerted to new fraud reports for any Beacon user in your system. When this webhook fires, to get more details on the report, use the Dashboard, or call [`/beacon/report_syndication/get`](/docs/api/products/beacon/#beaconreport_syndicationget).
4. (Optional) If you are monitoring for fraud on an ongoing basis, call [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate) when you are aware of any changes to user data (such as name, address or phone number) or if the user links a new account. Calling [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate) will immediately re-run fraud checks for the user and may generate Beacon hits.

##### Using Beacon Account Insights (optional)

Beacon Account Insights is an optional paid feature that allows you to build your own risk analysis on top of Plaid-reported risk signals. Reported risk signals include account age, recent PII changes on the account, and recent failed login attempts. For a full list of returned fields, see the API Reference for [`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget).

1. Create a Link token and use it to launch Link, following the [Link documentation](/docs/link/) for your platform. When calling [`/link/token/create`](/docs/api/link/#linktokencreate), be sure to include `beacon` in the `products` array.
2. The end user will then complete a Link session, resulting in a `public_token`, which can be obtained by either the `onSuccess` client-side callback or by calling [`/link/token/get`](/docs/api/link/#linktokenget).
3. Use [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to exchange the `public_token` for an `access_token`.
4. Call [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate) to associate the `access_token` with the Beacon user.
5. Call [`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget) to get account insights.

Before following these steps, set up [Identity Verification](/docs/identity-verification/) following the [Identity Verification integration process](/docs/identity-verification/#identity-verification-integration-process).

##### Initial setup

1. [Create a Beacon program](https://dashboard.plaid.com/sandbox/beacon/programs/new) via the Dashboard.
   - For best results, disable "auto flag reported user" under the "data breach hits" section, unless you have a specific business reason to enable it. Because so many users have been exposed to data breaches, enabling auto-flagging for data breach hits will result in a large number of users entering the review queue. All other automatically-selected options should be left enabled for most use cases.
2. Open the Identity Verification template editor and select your Beacon program under Setup > Beacon Fraud Screening, then click "publish".
3. (Optional) For best results, backfill six months of existing data into Beacon.
   - If you are adding Beacon to an Identity Verification deployment that has already been running in Production, then to backfill, call [`/beacon/report/create`](/docs/api/products/beacon/#beaconreportcreate) to report each known incident of fraud, using the `beacon_user_id` obtained from calling [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget), and making sure to include any known instances of fraud in the `report` object.
   - To backfill users that have *not* been through an Identity Verification session, call [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate) for each user, making sure to include any known instances of fraud in the `report` object.

Make sure to always use the same `client_user_id` when referring to the same end user, whenever you call [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate), [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate), or [`/link/token/create`](/docs/api/link/#linktokencreate).

##### User creation

1. To add a new user, launch Link as normal and have the user go through an Identity Verification session. You should *not* update the [`/link/token/create`](/docs/api/link/#linktokencreate) `products` array to include `beacon`.
2. To get the `beacon_id` for the user, call [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget). There is no need to call [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate).

##### Reviewing and reporting fraud signals

1. Review any Beacon hits via the Dashboard, to clear or reject the user.
2. If you become aware of any fraud committed by a user, report it to the network by clicking the "Report Fraud" button in the Beacon Dashboard or by calling [`/beacon/report/create`](/docs/api/products/beacon/#beaconreportcreate). To learn more about how to properly categorize fraud, see the [Fraud Reporting Guide](https://view-su2.highspot.com/viewer/3ba173969f637a61e6e9d8e41eba083b).
3. (Optional) To monitor for fraud on an ongoing basis, listen for the [`REPORT_SYNDICATION_CREATED`](/docs/api/products/beacon/#report_syndication_created) webhook to be alerted to any new fraud reports.
4. (Optional) If you are monitoring for fraud on an ongoing basis, call [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate) when you are aware of any changes to user data (such as name, address or phone number) or if the user links a new account. Calling [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate) will immediately re-run fraud checks for the user and may generate Beacon hits.

In the Link flow, a user will see a "successfully verified" screen if there were no Beacon (or Monitor, if using) hits and all required Identity Verification checks passed. If there were any hits, or if Identity Verification checks failed, the user will see the "verification failed" screen and be placed in the pending review state, where you can review their session from the Dashboard.

##### Using Beacon Account Insights (optional)

Beacon Account Insights is an optional paid feature that allows you to build your own risk analysis on top of Plaid-reported risk signals. Reported risk signals include account age, recent PII changes on the account, and recent failed login attempts. For a full list of returned fields, see the API Reference for [`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget).

1. Create a Link token and use it to launch Link, following the [Link documentation](/docs/link/) for your platform. When calling [`/link/token/create`](/docs/api/link/#linktokencreate), be sure to include `beacon` in the `products` array.
2. The end user will then complete a Link session, resulting in a `public_token`, which can be obtained by either the `onSuccess` client-side callback or by calling [`/link/token/get`](/docs/api/link/#linktokenget).
3. Use [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to exchange the `public_token` for an `access_token`.
4. Call [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate) to associate the `access_token` with the Beacon user.
5. Call [`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget) to get account insights.

#### Adding Beacon Account Insights to existing Items

1. Send the Item through [update mode](/docs/link/update-mode/). Beacon does not yet support `additional_consented_products`; you will set `"products": ["beacon"]` when initializing update mode.
2. Call [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate) (or, if the user was already created, [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate)), providing the Item's `access_token`(s), to associate the Beacon user with the Item.
3. Call [`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget) to get account insights.

#### Beacon workflows

All Beacon hits generate a `pending_review` status and must be reviewed in the Dashboard to resolve into a `cleared` or `rejected` state, unless the hit is for a first-party fraud report generated by your organization. Those hits will move the user directly to a `rejected` status.

When creating your Beacon template, you can optionally configure whether a detected duplicate should be moved into `pending_review` state.

#### Testing Beacon in Sandbox

In Sandbox, use the [Leslie Knope test user](/docs/identity-verification/testing/#primary-test-user-leslie-knope) to generate Beacon hits and test the Pending Review state.

To generate a Cleared status, you can use any user data that you have not already screened.

To generate a Rejected status, run a Beacon screening, report a `first_party` fraud incident associated with the user, then run another screening on the user.

When testing in Sandbox, re-using the same test user repeatedly will degrade performance over time. If running automated test suites, use randomly generated user data, rather than re-using the same data, in order to avoid performance issues. Similarly, it is not recommended to create large number of duplicate users to test duplicate user detection.

#### Beacon pricing

Participation in the Plaid Beacon Network is free of charge. The optional Beacon Account Insights feature, using the [`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget) endpoint, is billed under [a per-request flat fee model](/docs/account/billing/#per-request-flat-fee).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
