---
title: "Assets - Introduction to Assets | Plaid Docs"
source_url: "https://plaid.com/docs/assets/"
scraped_at: "2026-03-07T22:04:29+00:00"
---

# Introduction to Assets

#### Learn how to use Assets to access users' financial information for loan underwriting.

Get started with Assets

[API Reference](/docs/api/products/assets/)[Quickstart](/docs/quickstart/)[Demo](https://plaid.coastdemo.com/share/66fb0a180582208ffa82103e?zoom=100)

#### Overview

Most new customers should use [Consumer Report by Plaid Check](/docs/check/) instead of Assets. Consumer Report is an FCRA-compliant product providing underwriting scores and insights, as well as Day 1 Certainty compatible Asset Verification reports for mortgage lending. Assets is currently recommended only for use cases not supported by Consumer Report, such as underwriting outside the US.

Assets data is used to determine whether a user has enough assets to qualify for a loan. Plaid's Assets product allows you to access a user's Asset Report via the [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) (JSON) and [`/asset_report/pdf/get`](/docs/api/products/assets/#asset_reportpdfget) (PDF) endpoints. Asset Reports are created on a per-user basis and include a snapshot of information about a user's identity, transaction history, account balances, and more.

Prefer to learn by watching? Get an overview of how Assets works in just 3 minutes!

For a fuller view of a borrower's financial situation, Assets can be used with [Bank Income](/docs/income/), which provides user-permissioned income streams derived from deposit activity.

If you need exact copies of a user's statements, see [Statements](/docs/statements/).

#### No-code integration with the Credit Dashboard

The Credit Dashboard is currently in beta. To request access, contact your Account Manager. To use the Credit Dashboard for Assets, you must not be using any [Plaid Check](/docs/check/) products.

Assets data is available via a no-code integration through the [Credit Dashboard](https://dashboard.plaid.com/credit). You can fill out a form with basic information about your end user, such as their name and email address, and indicate how much history you would like to collect and whether you would like to enable [Bank Income](/docs/income/bank-income/) for the session as well. Plaid will then generate a hyperlink that can be used to launch a Plaid-hosted Link session. Plaid can email the link directly to your user (Production only, not available in Sandbox), or you can send it to them yourself.

After the end user successfully completes the Link session, their data will be available in the Dashboard for you to view, archive, or delete. The data shown will be the same data returned by a regular Assets call, in a user-friendly web-based session.

![Plaid Assets Dashboard: Summary tab with Profile, Income, and Accounts info. Shows user's reference ID, income, and bank accounts.](/assets/img/docs/assets-dashboard-view.png)

This integration is also compatible with Income, and you can enable Assets and Income in the same session.

Once you are enabled for the Credit Dashboard, all new sessions, including ones created via the API, will be displayed in the Dashboard.

Data from Link sessions created via the Credit Dashboard cannot be accessed via the Plaid API. If you require programmatic access to data, use the [code-based integration flow](/docs/assets/#integration-overview) instead.

#### Integration overview

There are a few steps you will need to take to obtain an Asset Report for a particular user.

1. First, using the Link flow, [create Items](/docs/api/items/#token-exchange-flow) for each of the user's financial institutions; in doing so, you will obtain an `access_token` for each institution. You can use these tokens to obtain an Asset Report as long as the `access_token` was generated with `assets` included in Link's `product` array. With Assets, the user, upon using Link to authenticate with their financial institution, will need to provide consent for Plaid to access account balances, account holder information, and transaction history for that institution. Should the user revoke access via [my.plaid.com](https://my.plaid.com/), Plaid will notify you via a [webhook](/docs/api/items/#user_permission_revoked).
2. Once you have the required `access_token`s, call the [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) endpoint to create an Asset Report. Note that Asset Reports are not created instantly. When the Asset Report is ready, you will be notified via a [`PRODUCT_READY`](/docs/api/products/assets/#product_ready) webhook. If low latency is important to your use case, you can call [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) with `options.add_ons` set to `["fast_assets"]`. This will cause Plaid to create two versions of the Asset Report: one with only current and available balance and identity information, and then later on the complete Asset Report. You will receive separate webhooks for each version of the Asset Report.
3. After receiving the [`PRODUCT_READY`](/docs/api/products/assets/#product_ready) webhook, you can retrieve an Asset Report in JSON format by calling the [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) endpoint. To retrieve an Asset Report in PDF format, call the [`/asset_report/pdf/get`](/docs/api/products/assets/#asset_reportpdfget) endpoint. To retrieve an Asset Report with cleaned and categorized transactions as well as additional information about merchants and locations in JSON format, call the [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) endpoint with `include_insights=true`. If you enabled Fast Assets when calling [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate), you can request the Fast version of the Asset Report by calling [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) with `fast_report=true`.
4. Sometimes, third party auditors need to see Asset Reports to audit lenders' decision making processes. To provide an Audit Copy of an Asset Report to a requesting third party, call the [`/asset_report/audit_copy/create`](/docs/api/products/assets/#asset_reportaudit_copycreate) endpoint and provide the auditor with an `audit_copy_token`. The auditor needs to be integrated with Plaid in order to view Audit Copies. Currently, Fannie Mae, Freddie Mac, and Ocrolus are the only auditors integrated with Plaid.

Sample asset report

```
{
    "report": {
        "asset_report_id": "ebb8f490-8f45-4c93-a6c3-5801bf92c3ff",
        "client_report_id": null,
        "date_generated": "2023-12-18T07:20:25Z",
        "days_requested": 2,
        "items": [
            {
                "accounts": [
                    {
                        "account_id": "1Gd3X8NmgLFVn47RorabTJ7Bvy8vBqfpaG5Ky",
                        "balances": {
                            "available": null,
                            "current": 320.76,
                            "iso_currency_code": "USD",
                            "limit": null,
                            "unofficial_currency_code": null
                        },
                        "days_available": 0,
                        "historical_balances": [],
                        "mask": "5555",
                        "name": "Plaid IRA",
                        "official_name": null,
                        "owners": [
                            {
                                "addresses": [
                                    {
                                        "data": {
                                            "city": "Malakoff",
                                            "country": "US",
                                            "postal_code": "14236",
                                            "region": "NY",
                                            "street": "2992 Cameron Road"
                                        },
                                        "primary": true
                                    }
                                ],
                                "emails": [
                                    {
                                        "data": "accountholder0@example.com",
                                        "primary": true,
                                        "type": "primary"
                                    }
                                ],
                                "names": [
                                    "Alberta Bobbeth Charleson"
                                ],
                                "phone_numbers": [
                                    {
                                        "data": "1112225555",
                                        "primary": false,
                                        "type": "mobile"
                                    }
                                ]
                            }
                        ],
                        "ownership_type": null,
                        "subtype": "ira",
                        "transactions": [],
                        "type": "investment"
                    },
                    {
                        "account_id": "7B73zkZlnLfeWMvkAlxpslbR4PM4RECdaKgll",
                        "balances": {
                            "available": 100,
                            "current": 110,
                            "iso_currency_code": "USD",
                            "limit": null,
                            "unofficial_currency_code": null
                        },
                        "days_available": 2,
                        "historical_balances": [
                            {
                                "current": 110,
                                "date": "2023-12-17",
                                "iso_currency_code": "USD",
                                "unofficial_currency_code": null
                            },
                            {
                                "current": 116.33,
                                "date": "2023-12-16",
                                "iso_currency_code": "USD",
                                "unofficial_currency_code": null
                            }
                        ],
                        "mask": "0000",
                        "name": "Plaid Checking",
                        "official_name": "Plaid Gold Standard 0% Interest Checking",
                        "owners": [
                            {
                                "addresses": [
                                    {
                                        "data": {
                                            "city": "Malakoff",
                                            "country": "US",
                                            "postal_code": "14236",
                                            "region": "NY",
                                            "street": "2992 Cameron Road"
                                        },
                                        "primary": true
                                    },
                                    {
                                        "data": {
                                            "city": "San Matias",
                                            "country": "US",
                                            "postal_code": "93405-2255",
                                            "region": "CA",
                                            "street": "2493 Leisure Lane"
                                        },
                                        "primary": false
                                    }
                                ],
                                "emails": [
                                    {
                                        "data": "accountholder0@example.com",
                                        "primary": true,
                                        "type": "primary"
                                    }
                                ],
                                "names": [
                                    "Alberta Bobbeth Charleson"
                                ],
                                "phone_numbers": [
                                    {
                                        "data": "1112223333",
                                        "primary": false,
                                        "type": "home"
                                    }
                                ]
                            }
                        ],
                        "ownership_type": null,
                        "subtype": "checking",
                        "transactions": [
                            {
                                "account_id": "7B73zkZlnLfeWMvkAlxpslbR4PM4RECdaKgll",
                                "amount": 6.33,
                                "date": "2023-12-09",
                                "iso_currency_code": "USD",
                                "original_description": "Uber 072515 SF**POOL**",
                                "pending": false,
                                "transaction_id": "mqdZ1xRbjau73pER6NrPSZQXZKQdN8CgLKrXA",
                                "unofficial_currency_code": null
                            }
                        ],
                        "type": "depository"
                    }
                ],
                "date_last_updated": "2023-12-18T07:16:21Z",
                "institution_id": "ins_3",
                "institution_name": "Chase",
                "item_id": "KEv98XqZ7yudGRpDqaxJsZ3PKDGzR8CRM13nr"
            }
        ],
        "user": {
            "client_user_id": null,
            "email": null,
            "first_name": null,
            "last_name": null,
            "middle_name": null,
            "phone_number": null,
            "ssn": null
        }
    },
    "request_id": "AYRjla7HgW4nKM9",
    "warnings": []
}
```

The `report` field of the above object contains the body of the Asset Report. This field consists of several subfields that contain information about the date and time the Asset Report was created; data submitted about the user; and information about Items, containing the user's historical balances, identity information, and more for each of the user's financial institutions. For a full explanation of the `report` field and its subfields, consult the [API Reference](/docs/api/products/assets/).

#### Assets webhooks

When you create an Asset Report, Plaid aggregates account balances, account holder identity information, and transaction history for the duration specified. If you attempt to retrieve an Asset Report before the requested data has been collected, you’ll receive a response with the HTTP status code 400 and a Plaid error code of `PRODUCT_NOT_READY`.

To remove the need for polling, Plaid sends a webhook to the URL you supplied when creating the Asset Report once the Report has been generated. If generation fails, the webhook will indicate why. For examples of webhooks corresponding to successful and unsuccessful Asset Report generation, see the [API Reference for Assets webhooks](/docs/api/products/assets/#webhooks).

#### Refreshing Asset Reports

Asset Reports can either be created or removed; they cannot be updated. To retrieve up-to-date information about users' transactions and balances, call the [`/asset_report/refresh`](/docs/api/products/assets/#asset_reportrefresh) endpoint. This endpoint, when used, creates a new Asset Report based on the corresponding existing Asset Report.

#### Getting an Asset Report for an existing Item

If your user has already connected their account to your application for a different product, you can add Assets to the existing Item via [update mode](/docs/link/update-mode/).

To do so, in the [`/link/token/create`](/docs/api/link/#linktokencreate) request, populate the `access_token` field with the access token for the Item, and set the `products` array to `["assets"]`. If the user connected their account less than two years ago, they can bypass the Link credentials pane and complete just the Asset Report consent step. Otherwise, they will be prompted to complete the full Link flow.

#### Financial Account Matching

If you are already using [Identity Verification](/docs/identity-verification/), you can enhance your verification process by enabling [Financial Account Matching](/docs/identity-verification/#financial-account-matching) in your Identity Verification template.
This feature allows you to compare the data collected during KYC with the bank account data collected when creating an Asset Report.
To ensure accurate matching, it's important to maintain consistency in the `client_user_id` used for end users across all products.

#### Financial Insights Reports (UK only)

In the UK, the Financial Insights Report, including risk and affordability insights, is available as an optional add-on to Assets. Risk insights provides details on instances of potential risks such as loan disbursements, loan payments, gambling, bank penalties, and negative balances. Affordability insights provides insights based on spending / income ratios, considering factors such as essential vs. non-essential spend and other financial commitments.

Like Asset Reports, Financial Insights Reports are available in both JSON and PDF format. [View a sample Financial Insights Report (UK only) PDF](https://plaid.com/documents/sample-financial-insights-report.pdf).

To learn more about Financial Insights Reports or to request pricing details, [contact Sales](https://plaid.com/contact/) or your Plaid Account Manager.

#### Testing Assets

Assets can be tested in [Sandbox](/docs/sandbox/) without any additional permissions and can be tested in Limited Production once your Production request is approved.

Plaid provides a [GitHub repo](https://github.com/plaid/sandbox-custom-users/) with test data for testing Assets in Sandbox, beyond that provided by the default Sandbox user. For more information on configuring custom Sandbox data, see [Configuring the custom user account](/docs/sandbox/user-custom/#configuring-the-custom-user-account).

#### Attributes library

Plaid provides a GitHub [Credit attributes library](https://github.com/plaid/credit-attributes) with Python scripts that can be used to derive various useful information from Asset Reports.

#### Assets pricing

Assets is billed on a [flexible fee model](/docs/account/billing/#per-request-flexible-fee); the cost will depend on the number of Items in the Asset Report (there is an additional cost for Asset Reports with more than 5 Items), as well as whether Additional History (over 61 days of history) has been requested for the report. Audit copies and PDF Asset Reports are billed based on a [per-request model](/docs/account/billing/#per-request-flat-fee). To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

#### Next steps

To learn more about building with Assets, see [Create an Asset Report](/docs/assets/add-to-app/) and the [API Reference](/docs/api/products/assets/).

For a more detailed overview of the steps involved in creating a full implementation of Plaid Assets, see the [Plaid Assets Solution Guide](/documents/plaid-assets-solution-guide.pdf).

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
