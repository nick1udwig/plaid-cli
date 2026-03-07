---
title: "Launch checklist | Plaid Docs"
source_url: "https://plaid.com/docs/launch-checklist/"
scraped_at: "2026-03-07T22:05:00+00:00"
---

# Launch in Production checklist

#### Check off these recommended steps before launching your app in Production

The Launch Checklist has been replaced by the [Launch Center](https://dashboard.plaid.com/developers/launch-center). The Launch Center shows a personalized list of steps for your integration, with links to supporting materials like videos, sample code, and docs. The Launch Checklist is no longer being updated.

Below is a list of recommended steps to take before launching your Plaid integration in production. While they might not all be required for the minimal operation of your application, the steps below will help to make your Plaid integration more robust, secure, efficient, and maintainable.

For a similar list of steps in a PDF guide format, see the [Plaid implementation handbook](https://plaid.com/documents/plaid-implementation-handbook.pdf).

#### Production setup

If you haven't already done so, [request Production access](https://dashboard.plaid.com/overview/production).

Complete your [application profile](https://dashboard.plaid.com/settings/company/app-branding) and [company profile](https://dashboard.plaid.com/settings/company/profile), which are required to access certain institutions that use OAuth-based connections.

Complete your [security questionnaire](https://dashboard.plaid.com/settings/company/compliance?tab=dataSecurity), which is required to access certain US institutions that use OAuth-based connections.

For European Union countries and the UK, there is a separate compliance process required to access financial institutions in Production. If your Plaid integration will support financial institutions in Europe (including the UK), and your business is not based in Europe, [file a support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access) to request Production access for Europe. Allow at least one week for your request to be processed.

Configure both Plaid Link and your server-side code to use the Production
environment by setting the appropriate value when configuring your client object or (if not using a client library) your HTTP request. For language-specific details, see the GitHub page for your client library, or [API host information](https://plaid.com/docs/api/#api-host) if not using a library.

Ensure you are using your Production secret, which can be found on the [dashboard](https://dashboard.plaid.com/team/api).

If migrating from Sandbox, ensure that you remove any usage of Sandbox-specific functionality, such as the `user_good` test user or calls to `/sandbox/` endpoints.

Add teammates to the dashboard to give other users access. For more details, see [Plaid teams](/docs/account/teams/).

#### Link setup

Follow the steps below to assist users in completing the Link flow, help ensure compliance with Plaid's policies, and avoid being billed for unneeded products.

[Implement OAuth support](https://plaid.com/docs/link/oauth/). Some institutions in both Europe and the US require the use of OAuth-based connections. While OAuth will typically "just work" on desktop, as long as you have completed your [application profile](https://dashboard.plaid.com/settings/company/app-branding), [company profile](https://dashboard.plaid.com/settings/company/profile), and [security questionnaire](https://dashboard.plaid.com/overview/questionnaire-start), you may need to implement client-side redirect logic for users on mobile devices.

Make sure your OAuth integration works with all the [test cases](/docs/link/oauth/#testing-oauth) recommended in the OAuth documentation.

To avoid unnecessary billing and user confusion, implement logic for preventing [duplicate Items](/docs/link/duplicate-items/).

Configure Link customizations within the
[dashboard](https://dashboard.plaid.com/link) and ensure that the countries and languages selected for the customizations match the settings provided to [`/link/token/create`](/docs/api/link/#linktokencreate). For more details, see [Link customization](/docs/link/customization/).

Ensure the `client_name` parameter of [`/link/token/create`](/docs/api/link/#linktokencreate) is set
to display your company's name as you'd like it to appear within Link.

Ensure that the `products` parameter of [`/link/token/create`](/docs/api/link/#linktokencreate) includes only the
products you intend to use. The products listed here will influence which
institutions and accounts appear in Link (only institutions and account types that support all specified products will
appear in Link) and will trigger a billing event for listed products upon
successful token exchange.

If you are using multiple Plaid products, see the recommendations in [Choosing how to initialize products](/docs/link/initializing-products/) to make sure your Link sessions are optimized for performance, conversion, and billing.

Provide any notice and obtain any consent required for Plaid to process end user
information in accordance with Plaid's [End User Privacy Policy](/legal/#end-user-privacy-policy).

You may find it helpful to know that some of our customers link to Plaid's privacy policy within their own privacy policy, while others surface a separate just-in-time notice and consent page during their onboarding flow. Ultimately, it is up to you to determine how to obtain any legally required consents.

To maximize the number of users who complete the Link flow, review Plaid's recommendations for [optimizing Link conversion](/docs/link/best-practices/) and implement any practices relevant to your use case, such as providing [pre-Link messaging](/docs/link/messaging/).

##### Callbacks

Handle callbacks beyond just `onSuccess` in order to gracefully handle errors and build analytics around Link.

Listen to the [`onExit()`](/docs/link/web/#onexit) and [`onEvent()`](/docs/link/web/#onevent) callbacks for `error_type` and `error_code` in order to implement error handling.

Listen to the [`onEvent()`](/docs/link/web/#onevent) callback for
`exit_status` or `timestamp` in order to implement Link conversion analytics. Alternatively, you can use Plaid's built-in [Link conversion analytics](https://dashboard.plaid.com/link-analytics) in the Dashboard.

#### Webhook configuration

Plaid uses webhooks for many common flows. If your integration does any of the following, webhooks are required or strongly recommended:

- (Any product) Your application is calling a Plaid endpoint for an Item repeatedly, over a period of time, not just immediately after Link.
- (Transactions or Investments) You are accessing transactions made after the end user linked their Item.
- (Auth) You are using automated micro-deposits for account verification, or making transfers over a year after the Item was initially linked.
- (Assets) You are creating Asset Reports.
- (Income) You are verifying a user's income.
- (Payment Initiation (UK and Europe)) You are initiating payments.
- (Virtual Accounts (UK and Europe) or Transfer) You are sending or receiving funds.

Make sure Link has been initialized with your URL for receiving webhooks via [`/link/token/create`](/docs/api/link/#linktokencreate).

(For Identity Verification, Transfer or Payment Initiation, as well as [optional Auth micro-deposit events](/docs/auth/coverage/microdeposit-events/)) Make sure to configure account level webhook URLs via the [Plaid Dashboard](https://dashboard.plaid.com/developers/webhooks).

Make sure you can receive webhooks from [Plaid's webhook IPs](/docs/api/webhooks/).

Review Plaid's [webhook best practices](/docs/api/webhooks/#best-practices-for-applications-using-webhooks) to ensure your webhook handling logic is robust to outages and traffic spikes.

For all Item-based products, listen for the [`PENDING_DISCONNECT`](/docs/api/items/#pending_disconnect) webhook and send the Item through update mode when receiving it.

#### Error handling

Sometimes, Plaid API calls may fail due to intermittent outages or connectivity errors at supported institutions. Implement retry logic or error handling as necessary for product API calls.

#### Link in update mode

Update mode is used to fix Items that have entered an error state (for example, because a user changed their password). If your application needs to access an Item repeatedly over a period of time, rather than just immediately after Link, implementing update mode logic is strongly recommended.

Handle the `ITEM_LOGIN_REQUIRED`, `PENDING_DISCONNECT`, and `PENDING_EXPIRATION` Item errors or webhooks by launching update mode to ensure your users retain access to their Items. For more information, see [Updating Items via Link](/docs/link/update-mode/).

Listen for the [`NEW_ACCOUNTS_AVAILABLE`](/docs/api/items/#new_accounts_available) webhook to learn when an Item has a new account associated with it. To request access to that account, launch Link in [update mode](/docs/link/update-mode/).

(Optional, Auth and Identity only) Implement [Product Validations](/docs/link/update-mode/#resolving-access_not_granted-or-no_auth_accounts-errors-via-product-validations) in update mode to prevent customers from accidentally revoking required permissions during the update mode flow.

#### Storage & logging

Log Plaid identifiers and IDs properly to enhance security, when contacting Support about a specific request or callback, and for finding specific entries in the [Activity Log](https://dashboard.plaid.com/activity/logs). For more information, see [Dashboard logs and troubleshooting](/docs/account/activity/).

[Access tokens](/docs/quickstart/glossary/#access-token) and [Item IDs](/docs/quickstart/glossary/#item-id) are the core identifiers that map your users to
their financial institutions. Store them securely and associate them with
users of your application. Make sure, however, that these identifiers are
never exposed client-side. Keep in mind that one user can create multiple
Items if they have accounts with multiple financial institutions.

  

The same storage requirements apply to other types of tokens used instead of access tokens by certain products, such as asset report tokens (used with Assets), payment profile tokens (used with Transfer) and user tokens (used with Income).

Ensure that the following identifiers are securely logged, as they will be needed when contacting Support about a specific request or callback.

`link_session_id`: Included in the `onExit`, `onEvent`, and `onSuccess` callback of a Link integration.

`request_id`: Included in all Plaid API responses.

`account_id`: Included in all successful Plaid API responses that relate to a specific Item or account.

`item_id`: Included in all successful Plaid API responses that relate to a specific Item.

#### Item management

Delete Items using [`/item/remove`](/docs/api/items/#itemremove) when they are no longer being used. For example, you may wish to allow users to remove linked Items through your app's account management interface, or you may want to delete Items when a user deletes their account with your service or becomes inactive, or if the Item has been in an error state for an extended period. Deleting unneeded Items is a security best practice and will also prevent you from being charged for these Items when using a subscription-based product such as Transactions.

#### Multi-app use cases

If you anticipate having multiple apps that use Plaid or having multiple teams at your company using it, consider creating a single shared backend service to handle functions such as logging and token storage in a consolidated place.

### Product-specific recommendations

#### Auth

If you are using Auth for an account funding use case, see the [Plaid account funding guide](/documents/plaid-account-funding-guide.pdf) for use case specific recommendations.

If launching in the US, support the [Automated micro-deposit](/docs/auth/coverage/automated/) and [Same-day micro-deposit](/docs/auth/coverage/same-day/) flows for maximum institution coverage.

If supporting the Instant Match, Automated micro-deposit, or Same-Day micro-deposit flows, make sure that the country codes parameter provided to [`/link/token/create`](/docs/api/link/#linktokencreate) (or directly to Link, if using a legacy public key implementation) for these flows includes only the US and no other countries.

If supporting the Automated micro-deposit flow, make sure to listen for [Auth webhooks](/docs/api/products/auth/#webhooks) to know when the transaction is completed.

Make sure you are not displaying account numbers in your app's UI, even if truncated or masked, to avoid user confusion when working with institutions that provide virtualized or temporary account numbers. Always use values from the `mask` field instead.

Listen for the [`USER_ACCOUNT_REVOKED`](/docs/api/items/#user_account_revoked) and [`USER_PERMISSION_REVOKED`](/docs/api/items/#user_permission_revoked) webhooks to be notified when a temporary account number has become invalidated. To avoid R04 errors, do not make a transfer using an account number, processor token, or Stripe bank account token that has been revoked. Instead, send the user through update mode (for `USER_ACCOUNT_REVOKED`) or create a new Item (for `USER_PERMISSION_REVOKED`), then call the auth endpoint again to obtain a new account number or token.

Listen for the [`AUTH: DEFAULT_UPDATE`](/docs/api/products/auth/#default_update) webhook to be notified when an Item's account or routing number has changed, then call [`/auth/get`](/docs/api/products/auth/#authget) or [`/processor/auth/get`](/docs/api/processor-partners/#processorauthget) again to refresh your information on file.

#### Balance

If you are using Balance for an account funding use case, see the [Plaid account funding guide](/documents/plaid-account-funding-guide.pdf) for use case specific recommendations.

If you are using Balance with non-depository accounts, such as credit card or loan accounts, make sure to specify `min_last_updated_datetime` when calling [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) to ensure balance calls to Capital One can succeed.

#### Identity

If you are using Identity for an account funding use case, see the [Plaid account funding guide](/documents/plaid-account-funding-guide.pdf) for use case specific recommendations.

#### Transactions

If fetching historical Transactions data using [`/transactions/get`](/docs/api/products/transactions/#transactionsget) or [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), make sure your app implements pagination logic. For sample code, see the [integration instructions](/docs/transactions/add-to-app/#fetching-transaction-data) or the API Reference for the endpoint.

If you use [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), make sure your app properly handles errors that occur during pagination by restarting the pagination loop. For more details, see [`TRANSACTIONS_SYNC_MUTATION_DURING_PAGINATION`](/docs/errors/transactions/#transactions_sync_mutation_during_pagination).

Handle Transactions webhooks, see [Handling Transaction webhooks](/docs/transactions/webhooks/).

#### Investments

Handle the `HOLDINGS`-type [`DEFAULT_UPDATE`](/docs/api/products/investments/#holdings-default_update) webhook if your app needs to keep holdings, values, and balances up-to-date.

Handle the `INVESTMENTS_TRANSACTIONS`-type [`DEFAULT_UPDATE`](/docs/api/products/investments/#investments_transactions-default_update) webhook if your app needs to keep Investments transactions up-to-date.

If fetching historical Investments transactions using [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget), make sure your app implements pagination logic. For sample code, see the [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) API Reference (pagination examples available in Python and Ruby code samples).

#### Assets

Handle the `ASSETS`-type [`PRODUCT_READY`](/docs/api/products/assets/#product_ready) webhook to know when to call [`/asset_report/get`](/docs/api/products/assets/#asset_reportget).

Handle the `ASSETS`-type [`ERROR`](/docs/api/products/assets/#error) webhook to gracefully handle errors from failures in [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate).

#### Payment Initiation

Handle the [`PAYMENT_STATUS_UPDATE`](/docs/api/products/payment-initiation/#payment_status_update) webhook to keep updated on payment status information.

#### Virtual Accounts

Handle the [`WALLET_TRANSACTION_STATUS_UPDATE`](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) webhook to keep updated on transaction status information.

#### Income

See the [Plaid income verification solution guide](/documents/plaid-income-verification-solution-guide.pdf) for use case specific recommendations.

Handle the `INCOME`-type [`income_verification`](/docs/api/products/income/#income_verification) webhook to know when the verification is complete and endpoints that show income or employment related data are ready to be called.

#### Identity Verification

See the [Plaid Identity Verification and Monitor solution guide](/documents/plaid-idv-monitor-solution-guide.pdf) for use case specific recommendations.

Make sure to update your template IDs when moving to Production; the Sandbox and Production environments have different template IDs.

If you've set any of your Acceptable Risk Levels to "high" in your Risk Rules in order to prevent checks from failing during testing, make sure to change them back to the levels you plan to use in Production.

Implement logic for users to retry identity verification if they have failed.
You can integrate this process directly using the `/identity_verification/retry/`
endpoint or from the Dashboard.

If enabling Selfie Check, make sure that you are [requesting the required camera permissions](/docs/identity-verification/#mobile-support) on mobile.

When performing your matching logic, ensure make sure to take into account that `no_data` is
different from `no_match`. `no_data` means that the list issuer didn’t supply data
against which to match. `no_match` means that data was provided by the list issuer
and it did not match the information provided by your customer.

#### Monitor

See the [Plaid Identity Verification and Monitor solution guide](/documents/plaid-idv-monitor-solution-guide.pdf) for use case specific recommendations.

Make sure to implement logic for both onboarding (adding users to a program when they are created) and offboarding (moving users to a program with rescans disabled when they close their account).

Ensure that the Monitor review queue is checked regularly. For more details, see [Building a reviewer workflow](/docs/monitor/#preparing-for-ongoing-screening) and [preparing for ongoing screening](/docs/monitor/#preparing-for-ongoing-screening).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
