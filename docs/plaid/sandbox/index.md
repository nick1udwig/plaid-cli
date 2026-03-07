---
title: "Sandbox - Overview | Plaid Docs"
source_url: "https://plaid.com/docs/sandbox/"
scraped_at: "2026-03-07T22:05:16+00:00"
---

# Sandbox Overview

#### Use the Sandbox to quickly develop and test your app

[API Reference](/docs/api/sandbox/)[Quickstart](/docs/quickstart/)

#### Sandbox overview

The Plaid Sandbox is a free and fully-featured environment for application development and testing. All Plaid functionality of both the Plaid API and Plaid Link is supported in the Sandbox environment. A variety of [test accounts](/docs/sandbox/test-credentials/) and [institutions](/docs/sandbox/institutions/) are available to test against, and you can create an unlimited number of test Items. Sandbox API keys can be obtained in the [Plaid Dashboard](https://dashboard.plaid.com).

The Sandbox environment provides capabilities for testing core use cases, but does not reflect the full scope and complexity of data that can exist in Production. After testing in Sandbox, it is recommended to test in Production or [Limited Production](/docs/sandbox/#testing-with-live-data-using-limited-production) to ensure your application can handle [institution-specific behaviors](/docs/link/oauth/#institution-specific-behaviors) and real-world data before launching.

#### Using Sandbox

The Sandbox can be reached by sending HTTPS POST requests to the endpoints on the `sandbox.plaid.com` domain. In order to use the Sandbox with a client library, specify Sandbox as your environment when initializing.

The default username/password combination for all Sandbox institutions is `user_good` / `pass_good`.

Most products can be immediately tested in Sandbox with no extra configuration. Some products, such as Payment Initiation (UK and Europe), may require your account be specially enabled. If the product you want to test against is not available in Sandbox, file a request for access ticket via the [Dashboard](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Bypassing Link

Because the Link UI can and will change from time to time, it is not recommended to write automated end-to-end tests that exercise the Link UI. You should always bypass Link if you are writing an automated test suite that can block your build.

When doing repeated manual tests, or for writing automated tests against Sandbox, completing the Link flow may be impractical. As an alternative, you can also create Items in Sandbox via the API, using a special Sandbox-only endpoint, [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate). This endpoint allows you to generate a `public_token` for an arbitrary institution ID, initial products, and test credentials.

#### Simulating data and events

The Sandbox provides rich test data, as well as the ability to [generate your own test data](/docs/sandbox/user-custom/). You can also simulate a number of scenarios for testing in Sandbox. Some of the most common Sandbox scenarios are listed below. For additional product-specific testing scenarios, see the documentation for the specific product you are using, or [Test credentials](/docs/sandbox/test-credentials/).

##### ITEM\_LOGIN\_REQUIRED

[Update mode](/docs/link/update-mode/) in Link is used to handle user logins that have become invalid (for example, due to the user changing their password at their financial institution). An Item that needs to be handled via update mode will enter the `ITEM_LOGIN_REQUIRED` error state. To test your update mode flow in Sandbox, you can use the [`/sandbox/item/reset_login`](/docs/api/sandbox/#sandboxitemreset_login) endpoint to force an Item into this state. Sandbox Items will also automatically enter the `ITEM_LOGIN_REQUIRED` state 30 days after being created.

An example of using update mode and testing in Sandbox can be found in the [Plaid Pattern](https://github.com/plaid/pattern) sample app. [Items.js](https://github.com/plaid/pattern/blob/master/server/routes/items.js#L195) illustrates how to incorporate the handling and testing of update mode using the [`/sandbox/item/reset_login`](/docs/api/sandbox/#sandboxitemreset_login) endpoint.

##### Instant Match and micro-deposit-based Auth

While the standard Instant Auth flow does not require special configuration to test, the additional Auth flows such as Instant Match, Same-day micro-deposits, and Automated micro-deposits require more complex steps such as verifying micro-deposits or routing numbers. The Sandbox environment provides several test institutions and endpoints that can be used to verify these flows. For more details, see [Testing Auth flows](/docs/auth/coverage/testing/).

##### Returning user experience

The Link returning user experience (formerly known as Remember Me) can be tested in Sandbox with several seeded phone numbers. More details can be found at [Testing returning user flow in Sandbox](/docs/link/returning-user/#testing-in-sandbox).

##### Transactions updates

To simulate updating transactions, see [Testing pending and posted transactions](/docs/transactions/transactions-data/#testing-pending-and-posted-transactions).

##### Webhooks

Plaid provides a special Sandbox-only endpoint, [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook), that can be used to trigger a number of webhooks on demand, allowing you to test that you are receiving webhooks successfully. Webhooks specific to the Income and Transfer products can be sent using the [`/sandbox/transfer/fire_webhook`](/docs/api/sandbox/#sandboxtransferfire_webhook) and [`/sandbox/income/fire_webhook`](/docs/api/sandbox/#sandboxincomefire_webhook) endpoints.

All webhooks that would be fired in Production will also fire in Sandbox, except when using the Transfer product.

#### Sandbox MCP server

Plaid has a [Sandbox MCP server](https://github.com/plaid/ai-coding-toolkit/tree/main/sandbox),
which, when run locally, allows your AI-powered agent to call many of these Sandbox endpoints, letting
you [trigger webhooks](/docs/sandbox/#webhooks) or [create public tokens](/docs/sandbox/#bypassing-link) to bypass Link.

#### Testing with live data using Limited Production

To use Link with bank data in Production or Limited Production in the US or Canada, all customers who created accounts after October 31, 2024 must [select a use case description](https://dashboard.plaid.com/link/data-transparency-v5) from the Link Customization section of the Dashboard.

Prior to obtaining Production access, you can make free API calls using live data in Limited Production.

You can Request Limited Production access via the Dashboard, either by clicking the "Test with Real Data" link on the Dashboard homepage, or by going directly to the [Limited Production request form](https://dashboard.plaid.com/overview/limited-production).

In Limited Production, the following restrictions apply:

- There is a cap on the number of API calls you can make for each product. Both failed and successful API calls count against this cap. API calls that do not relate to a specific product (such as [`/accounts/get`](/docs/api/accounts/#accountsget)) do not count against this cap.
- There is a cap on the total number of Items you can create, unless you have full Production access for at least one product.
- You cannot access certain major financial institutions, including Bank of America, Chase, and Wells Fargo, unless you have full Production access for at least one product and have also completed the OAuth registration process.
- The following products are supported in Limited Production: Assets, Auth, Balance, Identity (including Identity Match), Income, Investments, Liabilities, Payment Initiation, Transactions (including Transactions Refresh and Recurring Transactions), and Enrich. To test other products with live data, request Production access.
- If testing Payment Initiation in Limited Production, limitations on the type and quantity of payments apply. For details, see [Testing Payment Initiation](/docs/payment-initiation/payment-initiation-one-time/#testing-payment-initiation).

To remove the limits on your Limited Production access, apply for full Production access. Once you have full Production access for a product, usage limits for that product will be removed from your account. If you have any free API calls remaining, as long as you are direct customer of Plaid, you can continue to use them after receiving Production access; once those credits have been used up, your usage will be billed as normal Production traffic. If you are *not* a direct customer of Plaid and instead are accessing Plaid via a reseller partner, your free API credits will expire once you receive full Production access.

If you add a subscription-billed product to an Item in Limited Production (for example, by creating an Item initialized with Transactions, Investments, or Liabilities), you will not be charged for that subscription unless you continue to use the Item after using up your free API requests and receiving full Production access.

If you require additional API calls in Limited Production to evaluate a product, file a Support ticket.

#### Differences between Sandbox and Production

Aside from the fact that Sandbox uses test data and has [special endpoints](/docs/api/sandbox/) that can be used, Sandbox has the following special attributes:

- Does not always reflect institution-specific behaviors and quirks. For example, Sandbox does not reflect insititution-specific transaction history limits, and institutions that do not return pending transactions in Production may return pending transactions in Sandbox. Similarly, all Sandbox OAuth institutions use a single generic OAuth flow rather than institution-specific OAuth behavior. Certain institution-specific behaviors may be tested using [Sandbox test institutions](/docs/sandbox/institutions/).
- Does not send emails or text messages, such as confirmation emails after a user has linked an account, or text message reminders for the micro-deposit flow.
- Does not require a use case to be set, even if you are enrolled in [Data Transparency Messaging](/docs/link/data-transparency-messaging-migration-guide/).
- Allows `http` for redirect URIs, while Production requires `https`.
- Does not perform any OCR or image processing. To test flows that require these, such as Identity Verification Selfie Checks or Document Income, see the testing section for the specific product you are testing.
- Sandbox data is not always based on a consistent data source across different API calls, especially across different products. For example, Sandbox transaction history may show income and outflow inconsistent with the balance shown on the account, and some products, like Signal, generate random results in Sandbox. Triggering a transfer using a Transfer endpoint will not cause that transaction to show up when calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) or change the balance of the account when calling [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget).
- Sandbox may not show the latest updates to the Link UI until those updates have been rolled out to all eligible customers. This may occasionally result in temporary discrepancies between the Sandbox Link UI and the Production Link UI.
- In the OAuth flow, Sandbox will not enforce the checkboxes for "Select additional information you want to share" and will behave as though all boxes in that section are checked.

One exception to the rule that Sandbox does not use real data is Monitor. While Identity Verification checks cannot be made against real databases in Sandbox, Sandbox does use real Monitor watchlists, although they may not be as up-to-date as the watchlists used in Production.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
