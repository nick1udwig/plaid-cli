---
title: "Income - Payroll Income | Plaid Docs"
source_url: "https://plaid.com/docs/income/payroll-income/"
scraped_at: "2026-03-07T22:04:58+00:00"
---

# Payroll Income

#### Learn about Payroll Income features and implementation

Get started with Payroll Income

[API Reference](/docs/api/products/income/)[Quickstart](https://github.com/plaid/income-sample)

#### Overview

Payroll Income (US only) allows you to instantly verify employment details and gross income information, including the information available on a pay stub, from a user-connected payroll account. Payroll Income supports approximately 80% of the US workforce, including gig income workers.

Prefer to learn by watching? Get an overview of how Income works in just 3 minutes!

#### Integration process

1. Call [`/user/create`](/docs/api/users/#usercreate) to create a `user_token` that will represent the end user interacting with your application. This step will only need to be done once per end user. If you are using multiple Income types, do not repeat this step when switching to a different Income type.
2. Call [`/link/token/create`](/docs/api/link/#linktokencreate). In addition to the required parameters, you will need to provide the following:
   - For `user_token`, provide the `user_token` from [`/user/create`](/docs/api/users/#usercreate).
   - For `products`, use `["income_verification"]`. Document Income cannot be used in the same Link session as any other Plaid products, except for Payroll Income.
   - For `income_verification.income_source_types`, use `payroll`.
   - (Optional) If you are only using Payroll Income and do not want customers to use Document Income, for `income_verification.payroll_income.flow_types`, use `["payroll_digital_income"]`.
   - Provide a `webhook` URI with the endpoint where you will receive Plaid webhooks.
3. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see [Link](/docs/link/).
4. Open Link in your web or mobile client and listen to the `onSuccess` and `onExit` callbacks, which will fire once the user has finished or exited the Link session.
5. Listen for the [`INCOME: INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook, which will fire within a few seconds, indicating that the Income data is ready.
6. Call [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) for income data, and/or [`/credit/employment/get`](/docs/api/products/income/#creditemploymentget) for employment details.

#### Payroll Income Refresh

To request access to Payroll Income Refresh, contact your account manager.

On-demand Payroll Income Refresh allows you to request updated information from a previously connected payroll account.

To trigger a refresh, call [`/credit/payroll_income/refresh`](/docs/api/products/income/#creditpayroll_incomerefresh) and specify the `user_token` to refresh.

If the refresh is successful, you will receive an [`INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook. The next time you call [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget), you will receive updated data.

If the refresh was not successful, you will receive a [`INCOME_VERIFICATION_REFRESH_RECONNECT_NEEDED`](/docs/api/products/income/#income_verification_refresh_reconnect_needed) webhook. To resolve this failure state, send the user through [update mode](/docs/link/update-mode/). The refresh attempt will automatically be retried once the user has completed update mode, and an [`INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook will be sent if the retry is successful.

#### Download original documents

To download PDF versions of the original payroll documents that were parsed to obtain payroll data, use the `document_metadata.download_url` returned by [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget). Not all integrations will return a `download_url`.

In cases where an original source cannot be obtained, Plaid can optionally generate a PDF pay stub based on the data obtained. To enable Plaid-generated PDF pay stubs, contact your account manager. [View an example generated pay stub](https://plaid.com/documents/plaid-generated-mock-paystub.pdf).

#### Employment data

Plaid provides employment data via the [`/credit/employment/get`](/docs/api/products/income/#creditemploymentget) endpoint. If you are on a Pay-as-you-go or Growth plan and want to use this endpoint, contact support to request access.

#### Testing Payroll Income

You can test Payroll Income in Sandbox using Link. For best results, select a payment provider that only requests a username, password, and/or SSN. (Good examples are Paycom, Paychex, or Workday.) Use `test-good` as the username, `test` as the password, and `1234` as the last 4 digits of your SSN.

You can also see an example Link flow for Payroll Income using the [Plaid Link demo site](https://plaid.com/link-demo/): select "Payroll and Document Income" from the drop-down and use the credentials above.

[`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) can optionally be tested in Sandbox without using Link. Call [`/user/create`](/docs/api/users/#usercreate), pass the returned `user_token` to [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) (use `ins_135842` when creating a public token for Payroll Income in Sandbox), and then call [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget). The output of [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) will not be used, but calling it initializes the user token for testing.

Payroll Income does not currently support the use of custom Sandbox data; only the default `user_good` user is compatible with Sandbox Payroll Income.

#### Payroll Income pricing

Payroll Income is billed on a [one-time fee model](/docs/account/billing/#one-time-fee). Payroll Income Refresh is billed on a [per-request fee model](/docs/account/billing/#per-request-flat-fee). To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

#### Next steps

If you're ready to launch to Production, see the Launch checklist.

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
