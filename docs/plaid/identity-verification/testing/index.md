---
title: "Identity Verification - Testing in Sandbox | Plaid Docs"
source_url: "https://plaid.com/docs/identity-verification/testing/"
scraped_at: "2026-03-07T22:04:56+00:00"
---

# Testing in Sandbox

#### Test values and options for Identity Verification

Identity Verification is not available by default in the Sandbox environment. To obtain Sandbox access, [request full Production access for Identity Verification](https://dashboard.plaid.com/settings/team/products), which will automatically grant Sandbox access, or [contact Sales](https://plaid.com/products/identity-verification/#contact-form). To obtain Sandbox access for Identity Verification if you are already using another Plaid product, you can also contact your account manager or submit an [access request ticket](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access).

#### Sandbox users for Identity Verification

##### Primary test user (Leslie Knope)

In Sandbox mode, Identity Verification accepts the fixed set of inputs below in order to result in successful Data Source and Document Verification checks.

| Form Field | Test Value |
| --- | --- |
| Mobile number | +1 234-567-8909 |
| First name | Leslie |
| Last name | Knope |
| Verification code | 11111 |
| Address | 123 Main St. |
| City | Pawnee |
| State | Indiana |
| ZIP code | 46001 |
| Month | January |
| Day | 18 |
| Year | 1975 |
| SSN | 123-45-6788 or 123-45-6789 |

##### Custom test users

For US verifications, Plaid supports importing custom test users from your organization into Sandbox, allowing you to test your own sets of PII to result in successful Data Source and Document Verification checks. To learn more about creating custom test users for Identity Verification, contact your Plaid Account Manager.

#### Identity verification behavior in Sandbox

##### Data Source checks

Data Source checks in Sandbox will be compared against the information above. To simulate different test cases and results, you can adjust your Data Source input at the attribute level (e.g. an incorrect birthdate), and set up different Identity Rules to simulate different verification results.

##### Document checks

Documents uploaded in Sandbox will always be interpreted as genuine (not fake) documents reflecting the name and date of birth matching the Leslie Knope test user. Document check will pass if the data provided in the Data Source check matches this name and date of birth, and will fail otherwise.

Users have three attempts to pass the Document check, so the status will not be updated in the Dashboard to `failed` until you have provided mismatching data three times.

##### Selfie checks

Selfie checks will not be run in Sandbox mode, even if they are enabled in your template.

##### AML Screening

In Sandbox, Monitor screens against a real-world dataset, so the Sandbox user will not return any screening hits. For more information, including example data you can use to trigger screening hits, see [Testing Monitor](/docs/monitor/#testing-monitor).

##### Risk Check

The risk check is fully functional in Sandbox. It’s common to trigger high risk scores when testing. For more information and advice, see [Risk rules for testing](/docs/identity-verification/testing/#risk-rules-for-testing).

##### Auto-fill

Auto-fill behavior is fully functional in Sandbox.

##### Financial Account Matching

Financial Account Matching is fully functional in Sandbox. If you want to change the Identity data that is being compared to the Identity Verification session, see [how to create Sandbox test data](/docs/sandbox/user-custom/).

#### Risk rules for testing

Many common testing behaviors (e.g. attempting to enter the Identity Verification flow repeatedly on the same device using different credentials) may be flagged as risky behavior and cause your verification attempt to fail. For testing purposes, you can temporarily set the Acceptable Risk Level for any checks you are failing (typically Network Risk and Device Risk) to **High**, under the **Rulesets -> Risk Rules** section of the template editor. Make sure to set the Acceptable Risk Level back to your desired setting before launching in Production.

If you want to force a check to fail in Sandbox, an easy way to do this is to set Acceptable Risk Level for the Network Risk and Device Risk fields to **Low**. Checks will begin failing after your first few attempts.

#### Triggering specific test outcomes

Using the rules above, it is easy to create the outcome of passing both Data Source checks and Document checks, or of failing both checks. If you need to simulate specific combinations of pass / fail (passing one check but not the other), you can use the guides below.

##### Data Source check passes, Document Verification check fails

1. Set up your template to always be enabled for Data Source checks and Document Checks.
2. Set up your template to not require a match or partial match on name or date of birth.
3. Enter Leslie Knope's correct test data for the address, SSN, and phone number fields.
4. Enter incorrect data for the name and date of birth fields.

If you would like this test to also trigger a Monitor hit, you can use the name and date of birth of a sanctioned individual. For examples, see [Monitor test data](/docs/monitor/#testing-monitor).

##### Data Source check fails, Document Verification check passes

1. Set up your template to be enabled for Data Source checks, with Document checks as a fallback.
2. Set up your template to require a match for address, SSN, and phone number fields.
3. Enter Leslie Knope's correct test data for the name and date of birth fields.
4. Enter incorrect data for the address, SSN, and phone number fields.

#### Testing Financial Account Matching

1. Create a template with Financial Account Matching enabled.
2. Call [`/link/token/create`](/docs/api/link/#linktokencreate) with `identity_verification` in the `products` array. Make sure to save the `client_user_id` you are using with this call, as you will need it later.
3. Launch Link and complete the Identity Verification session, entering Leslie Knope's correct test data.
4. See the [Sandbox custom user repo](https://github.com/plaid/sandbox-custom-users/tree/main) for instructions on configuring custom Sandbox users and create a custom user using the [Financial Account Matching custom user data](https://github.com/plaid/sandbox-custom-users/blob/main/identity/leslie_knope_financial_account_matching.json) from that repo.
5. Call [`/link/token/create`](/docs/api/link/#linktokencreate) with `identity` (or `assets`, for underwriting use cases) in the `products` array and the same `client_user_id` as was used in the first step, and launch a Link session.
6. Use the custom user you created above to log in to Link, then call [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to exchange the public token for an access token.
7. Call [`/identity/get`](/docs/api/products/identity/#identityget) or [`/identity/match`](/docs/api/products/identity/#identitymatch) (or [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate), if using `assets` instead of `identity`) on the access token.
8. Check the Identity Verification session in the Dashboard. You should see the linked account and a successful match in the Financial Account Matching section.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
