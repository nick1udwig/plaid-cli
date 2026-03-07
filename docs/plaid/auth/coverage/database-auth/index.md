---
title: "Auth - Database Auth | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/database-auth/"
scraped_at: "2026-03-07T22:04:31+00:00"
---

# Database Auth

#### Evaluate a manually entered account using Plaid network data

#### Overview

Database Auth can increase conversion by providing instant account verification when end users cannot, or do not want to, log into their financial institution to verify their account. Examples include end users at banks that are not supported by Plaid, or users who need to verify a business bank account but who don't have access to manage the account via online banking. With Database Auth, users can instead enter their account and routing number manually.

![Link flow. First screen shows Plaid consent pane, second screen shows radio toggle option to connect bank instantly (recommended) or manually, manually option is selected. Third screen shows a form where user can enter account details: account and routing number, owner name, and checking or savings.](/assets/img/docs/auth/database-auth.png)

Example Link flow with Database Auth

Database Auth verifies account and routing numbers by checking the information provided against Plaid's known account numbers, leveraging Plaid's database of over 200 million verified accounts. If no match is found, Plaid will check the account number format against known usages by the institution associated with the given routing number. Database Auth will provide a verification status of 'pass', 'pass with caution', or 'fail' and a set of attributes that contributed to that status, such as whether a match was found or whether Plaid fell back to checking account number formats.

Database Auth does not verify that the user has access to the bank account and does also not fully guarantee that the account exists. Database Auth does not create a live connection to the bank account and cannot be used to check real-time balances (although it is compatible with anti-risk products such as Signal and Identity Match). For these reasons, Database Auth should not be enabled where there is a very high risk of fraud or ACH returns.

Database Auth is available for Auth flows in the United States. In Canada, Database Auth functionality is provided by the [Database Insights](https://plaid.com/docs/auth/coverage/database/) feature.

You can try out the Database Auth flow in an [Interactive Demo](https://plaid.coastdemo.com/share/67d0ce0df465686c02cc4fd2?zoom=100&step=7).

##### Database Auth flow

1. Starting on a page in your app, the user clicks an action that opens Plaid Link.
2. Inside of Plaid Link, the user selects an option that triggers the manual verification flow. This could occur because the user opted in to the manual flow (if Auth Type Select is enabled) or because their institution does not support Instant Auth.
3. The user will be prompted to enter their name, routing number, and account number.
4. Based on rules you have configured, the user will either see a success screen or a failure screen, or be prompted to fall back to micro-deposit-based verification.

#### Implementation steps

1. Create a `link_token` with the following parameters:
   - The `products` array should include `auth` or `transfer` only, and no other products.
   - If you are enabling Identity Match, put `identity` in the `required_if_supported_products` array, as putting it in the `products` array will prevent Database Auth from activating. Approximately one-third of all Items verified via Database Auth are compatible with Identity Match.
   - If you plan to use Signal Transaction Scores, put `signal` in the `additional_consented_products` or `required_if_supported_products` array. Putting it in the `products` array will prevent Database Auth from activating. Approximately one-third of all Items verified via Database Auth are compatible with Signal Transaction Scores.
   - No other products besides Signal Transaction Scores and Identity Match are compatible with Database Auth Items. To enable other products for Items where Database Auth is not used, put those products in the `additional_consented_products` array.
   - `country_codes` should be set to `['US']` – Database Auth is currently only available in the United States.
2. In the [Dashboard Account Verification pane](https://dashboard.plaid.com/account-verification), ensure "Manage via Dashboard" is selected, and enable "Verify with Database Verification".
3. Configure the rules:
   - (Recommended for most customers) To use Plaid's rules-based, customizable, no-code evaluation logic, select "For certain Database Auth results, fallback".
   - To use your own evaluation logic, select "For any Database Auth result, complete Link". This option is designed for customers who require more fine-grained control of the Link flow beyond what the ruleset configuration allows, or who have highly sophisticated fraud evaluation systems and require more custom logic. See [Using your own evaluation logic](/docs/auth/coverage/database-auth/#using-your-own-evaluation-logic-advanced).
4. [Open Link](/docs/link/) according to the instructions for your platform, using the `link_token` created earlier.
5. After the end user completes the Link session, if the session was successful, or if you are using manual evaluation logic, Link will return a `public_token` that you can exchange for an `access_token`. You can then proceed to call [`/auth/get`](/docs/api/products/auth/#authget) and/or [`/processor/auth/get`](/docs/api/processor-partners/#processorauthget), as described in the main [Auth integration process](/docs/auth/#auth-integration-process).
6. If the session was unsuccessful, Link will fall back to either an error screen or to a micro-deposit flow, depending on which options you selected in the Account Verification pane.

#### Backend-only implementation steps

The backend-only [`/auth/verify`](/docs/api/products/auth/#authverify)-based flow for Database Auth is in Early Availability. To request access, contact Sales or your Plaid account manager.

Database Auth can be used to verify account numbers you have already collected without the Plaid Link flow. Call [`/auth/verify`](/docs/api/products/auth/#authverify) and provide the account and routing number and, optionally, the account owner name. The response will include the same Database Auth-related results as [`/auth/get`](/docs/api/products/auth/#authget): the verification status, name match score (if a name was provided in the request and name matching data is available), and verification insights. For more details, see the [`/auth/verify`](/docs/api/products/auth/#authverify) API reference.

#### Understanding verification status and name match score

Your business logic will be based on the results of the verification status and the name match score.

##### Verification status

The possible values and most common causes for [`verification_status`](/docs/api/products/auth/#auth-get-response-accounts-verification-status) for an Item going through the Database Auth flow are:

`database_insights_pass`: Plaid was able to match the account and routing number provided to an existing Item on the Plaid network and has high confidence that the account exists. The account is not known to be closed or frozen.

`database_insights_pass_with_caution`: Plaid was unable to match the account and routing number provided to an existing Item on the Plaid network, but the routing number is valid and the account number matches Plaid's known rules for account number formatting at the associated financial institution.

`database_insights_fail`: The routing number is not recognized, or the account number does not appear to match the formatting rules at the associated financial institution, or Plaid has received data indicating the account is closed or frozen.

In addition to the common causes above, additional risk factors may impact `verification_status` results.

Balance is not compatible with Items verified via Database Auth. If you need to check for potential ACH return risk for items verified via Database Auth, you must use Signal Transaction Scores.

###### Understanding the pass with caution verification status

Because `database_insights_pass_with_caution` is a common Database Insights result, how you choose to handle it is important.

In a typical test scenario, accepting only `pass` results reduced administrative return rates by by over 80%, while accepting both `pass` and `pass_with_caution` reduced administrative return rates by approximately 50%. Note that results will vary based on your own user base and use case and are not guarantees.

If you do not want to accept `database_insights_pass_with_caution`, you can configure Link to use a [Fallback flow](/docs/auth/coverage/database-auth/#fallback-flows).

Signal Transaction Scores and Identity Match are only compatible with Database Auth Items whose verification status is `database_insights_pass`. If you need to use Signal Transaction Scores to check for non-administrative ACH return risk (e.g. R01 "insufficient funds" returns), or need to check for account takeover fraud using Identity Match, you should not accept the `database_insights_pass_with_caution` status.

##### Name match score

The possible values for the name match score range from 0-100. In general, you should require a name match of at least 70 to effectively protect against account takeover attacks. A name match score will only be returned if the session configuration was enabled for name matching and a name could be found to match against.

###### Example of how to interpret name match score

| Range | Meaning | Example |
| --- | --- | --- |
| 100 | Exact match | Andrew Smith, Andrew Smith |
| 85-99 | Strong match, likely spelling error, nickname, or a missing middle name, prefix or suffix | Andrew Smith, Andrew Simth |
| 70-84 | Possible match, likely alias or nickname and spelling error | Andrew Smith, Andy Simth |
| 50-69 | Unlikely match, likely relative | Andrew Smith, Betty Smith |
| 0-49 | Unlikely match | Andrew Smith, Ray Charles |

##### Using your own evaluation logic (advanced)

If you selected "For any Database Auth result, complete Link", you must write your own code for handling the results of the Database Auth verification.

To get the results, call [`/auth/get`](/docs/api/products/auth/#authget) and examine the `verification_status`, `verification_insights`, and `verification_name` fields. Based on the values in these fields, you will make a decision on whether to accept the account data as verified, take additional risk mitigation steps, or reject the account data as unverified.

The `verification_insights` object will contain more detailed analysis showing exactly which checks passed or failed to result in the given `verification_status`. It will also contain the `name_score`.

For more details on the `verification_insights` values, see the [API Reference](/docs/api/products/auth/#auth-get-response-accounts-verification-insights).

When using your own evaluation logic, you can fallback only to Automated Micro-deposits, and not to Instant Micro-deposits or Same-Day Micro-deposits.

When using other flows, customers using a [processor partner](/docs/auth/partnerships/) do not typically need to call [`/auth/get`](/docs/api/products/auth/#authget), and can directly call [`/processor/token/create`](/docs/api/processors/#processortokencreate) instead. However, if you are using Database Auth with a processor partner and using your own evaluation logic, you must call [`/auth/get`](/docs/api/products/auth/#authget) and check the value of the `verification_status` and/or `verification_insights` fields before passing a processor token to the partner.

##### Fallback flows

If Database Auth fails, you can configure fallback to [Instant Micro-deposits](/docs/auth/coverage/instant/#instant-micro-deposits), [Same-Day Micro-deposits](/docs/auth/coverage/same-day/), or [Automated Micro-deposits](/docs/auth/coverage/automated/) via the Account Verification Dashboard.

#### Testing Database Auth in Sandbox

For test credentials that can be used to test Database Insights in the Sandbox environment, see [Testing Database Auth](/docs/auth/coverage/testing/#testing-database-auth-or-database-insights).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
