---
title: "Statements - Introduction to Statements | Plaid Docs"
source_url: "https://plaid.com/docs/statements/"
scraped_at: "2026-03-07T22:05:19+00:00"
---

# Statements

#### Retrieve a PDF copy of a user's financial statement

Get started with Statements

[API Reference](/docs/api/products/statements/)[Quickstart](/docs/quickstart/)

#### Overview

Statements (US only) allows you to retrieve an exact, bank-branded, PDF copy of an end user's bank statement, directly from their bank. The Statements product simplifies the
process of collecting documents for loan verification purposes or rental application review purposes, among others.

Prefer to learn by watching? Get an overview of how Statements works in just 3 minutes!

#### Integration process

1. Call [`/link/token/create`](/docs/api/link/#linktokencreate) to create a link token.
   - Include a statements object containing fields `start_date` and `end_date`. Plaid allows extracting up to 2 years
     of statements.
   - If your integration uses multiple Plaid products, such as Assets and Statements, we recommend that you put your
     other products in the `products` array and put `statements` in the `required_if_supported_products` array. This
     configuration will require statements to be extracted if the financial institution supports the product, but
     will not block the user from progressing if statements are not supported. Instead, if Statements is not supported by
     the user's institution, you will receive a [`PRODUCTS_NOT_SUPPORTED`](/docs/errors/item/#products_not_supported) error
     when calling Statements endpoints. For more details, see [Choosing how to initialize products](/docs/link/initializing-products/).
   - If you are using only the Statements product, put statements in the products array when calling [`/link/token/create`](/docs/api/link/#linktokencreate).

/link/token/create

```
const request: LinkTokenCreateRequest = {
  loading_sample: true
};
try {
  const response = await plaidClient.linkTokenCreate(request);
  const linkToken = response.data.link_token;
} catch (error) {
  // handle error
}
```

1. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details,
   see the [Link documentation](/docs/link/).
2. Once the user has successfully finished the Link session, the client-side `onSuccess` callback will fire. Extract the
   `public_token` from the callback and exchange it for an `access_token` by calling [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange).
3. Call [`/statements/list`](/docs/api/products/statements/#statementslist), passing the `access_token` obtained above. This will return a list of statements, including a
   `statement_id` for each.

   Sample successful /statements/list response

   ```
   {
         "accounts": [
           {
              "account_id": "1qKRXQjk8xUWDJojNwPXTj8gEmR48piqRNye8",
              "account_name": "Plaid Checking",
              "account_type": "depository",
              "statements": [
                 {
                   "month": 1,
                   "statement_id": "efgh12e3-gh1c-56d6-e7e9-923bc64d80a5",
                   "year": 2024
                 },
                 {
                   "month": 2,
                   "statement_id": "jklh12e3-ab3e-87y3-f8a0-908bc64d80a5",
                   "year": 2024
                 },
                 {
                   "month": 3,
                   "statement_id": "4710abc-af28-481a-991a-48387a7345ddf",
                   "year": 2024
                 }
              ]
           }
         ],
         "institution_id": "ins_118923",
         "institution_name":  "First Platypus Bank",
         "item_id": "wz666MBjYWTp2PDzzggYhM6oWWmBb",
         "request_id": "NBZaq"
      }
   ```
4. Call [`/statements/download`](/docs/api/products/statements/#statementsdownload), passing in the `access_token` and desired `statement_id`, to download a specific statement.
   The statement will be provided in PDF format, exactly as provided by the financial institution.

   ![Mock bank statement from First Platypus Bank with account summary and February 2022 transactions. Retrieved via /statements/download.](/assets/img/docs/asset-reports/statements-bank-statement.png)

   Sample Sandbox (mock) bank statement retrieved with /statements/download.
5. (Optional) If you would like to re-check for new statements generated after the end user linked their account, you can call [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh). When the [`STATEMENTS_REFRESH_COMPLETE`](/docs/api/products/statements/#statements_refresh_complete) webhook has been received, call [`/statements/list`](/docs/api/products/statements/#statementslist) again for an updated list of statements.

#### Adding Statements to an existing Item

If your user has already connected their account to your application for a different product, you can add Statements to the existing Item via [update mode](/docs/link/update-mode/).

To do so, in the [`/link/token/create`](/docs/api/link/#linktokencreate) request described in the [Integration process](/docs/statements/#integration-process), populate the `access_token` field with the access token for the Item, and set the `products` array to `["statements"]`. If the user connected their account less than two years ago, they can bypass the Link credentials pane and complete just the Statements consent step. Otherwise, they will be prompted to complete the full Link flow.

#### Supported accounts and institutions

Statements currently supports only bank depository accounts (e.g. checking and savings accounts).

Statements support includes the following major institutions, constituting ~40% of US depository accounts:

- Bank of America
- Chase (Early Availability; contact your Plaid Account Manager to request access)
- Citibank
- Fifth Third Bank
- Huntington Bank
- Navy FCU
- Regions Bank
- Truist
- US Bank
- Wells Fargo

In addition to the institutions named above, Statements also supports several smaller banks and credit unions. Statements does not currently support all major or long tail-institutions, and should be used with a fallback option in case data is not available.

For a full list of institutions that support Statements, see the [Institution Coverage Explorer](/docs/institutions/).

#### Testing Statements

Statements can be tested in [Sandbox](/docs/sandbox/), where a mock statement is returned. In Production, the statement is retrieved from the financial institution. Existing customers whose Plaid teams were created in 2023 or earlier may need to file a [product access request](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access) to access Statements in the Sandbox.

#### Statements pricing

Statements is billed when [`/link/token/create`](/docs/api/link/#linktokencreate) is successfully called for Statements, with the cost based on the number of statements between the provided start and end dates, according to a [flexible per-Item fee model](/docs/account/billing/#per-item-flexible-fee). Statements Refresh is billed based on the number of statements extracted between the provided start and end dates when calling [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh), based on a [flexible per-request fee model](/docs/account/billing/#per-request-flexible-fee). To view the exact pricing you may be eligible for, [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

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
