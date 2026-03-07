---
title: "Link - Choosing when to initialize products | Plaid Docs"
source_url: "https://plaid.com/docs/link/initializing-products/"
scraped_at: "2026-03-07T22:05:05+00:00"
---

# Choosing how to initialize products

#### Increase conversion and reduce costs by learning how to configure products in Link

#### Overview

When you create a Link token, you specify which products to initialize Link with by including those products in the [`/link/token/create`](/docs/api/link/#linktokencreate) request. Which products you specify in the request parameters will impact many aspects of your Plaid integration, including user conversion, institution availability, the latency your user experiences when connecting their accounts in Link, the latency you experience when retrieving data from the API, and which products you will be billed for.

You can also add products to an Item post-Link, which is discussed below.

#### Initializing products during Link

The table below summarizes the [`/link/token/create`](/docs/api/link/#linktokencreate) request parameters that determine product initialization in Link:

| Parameter | Required? | Summary | Will Link restrict institutions? | When will Plaid bill me? |
| --- | --- | --- | --- | --- |
| products | Yes | The user must connect an applicable institution and account. | Yes. Only institutions that support all products in this array will be available in Link. | For products billed under a [one-time or subscription fee](/docs/account/billing/) model, you will be billed upon Item creation. |
| required if supported products | No | If the institution supports the products and an applicable account is selected, Plaid will treat these products as required for Item creation to succeed. | No | For products billed under a [one-time or subscription fee](/docs/account/billing/) model, you will be billed upon Item creation. |
| optional products | No | Plaid will pull the data from the institution if possible, but in the event of a failure, Item creation will still succeed. | No | For Auth, Identity, and Signal Transaction Scores, you will only be billed if you use the product.  For other products billed under a [one-time or subscription fee](/docs/account/billing/) model, you will be billed upon Item creation. |
| additional consented products | No | Plaid will collect the user's consent to retrieve this data so that if you decide to use this product later on, you won't need to collect the user's consent again. | No | When you call the product endpoint for the first time |

#### Adding products post-Link

If a product was not initialized during Link, calling an endpoint for that product will automatically attempt to add that product to the Item. For example, if an Item was not initialized with Transactions during Link, calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) or [`/transactions/get`](/docs/api/products/transactions/#transactionsget) on that Item for the first time will attempt to add the Transactions product to the Item.

Data Transparency Messaging is automatically enabled for all new customers' US and Canada-based Link sessions as of October 2024. On an Item enabled for [Data Transparency Messaging](/docs/link/data-transparency-messaging-migration-guide/), you can only add products by calling an endpoint if you specified those products in the Additional Consented Products array when calling [`/link/token/create`](/docs/api/link/#linktokencreate), or if you already have the required [permissions scopes](/docs/link/data-transparency-messaging-migration-guide/#data-scopes-and-consent) for those products.

As another example, if the Item is at an OAuth institution that requires users to grant product-specific access, the end user will need to have granted that access to that product during the OAuth consent flow. To require that users provide all OAuth consent necessary to use a secondary product, use [Required if Supported Products](/docs/link/initializing-products/#required-if-supported-products) instead of adding the product by calling its endpoint post-Link.

If you do not have all the consent required to add a product, your API call attempting to add the product will fail and you will need to send the user through [update mode](/docs/link/update-mode/) to obtain that consent.

Some additional exceptions to the post-Link product add flow exist: Identity Verification and Payment Initiation cannot be added post-Link. Assets, Statements, and Bank Income can only be added post-Link via the update mode flow.

#### Impacts of product initialization on latency

If you initialize Link with a product in the Products, Required if Supported Products, or Optional Products array, Plaid will prepare relevant product data during or immediately after Link. Otherwise, Plaid will only begin to fetch data for that product when you call an endpoint for that product, such as [`/auth/get`](/docs/api/products/auth/#authget) or [`/transactions/sync`](/docs/api/products/transactions/#transactionssync). You can minimize product latency by initializing Link with a given product, as opposed to adding the product to the item post-Link.

Products that retrieve extended account history, such as Transactions and Investments, may take a minute or longer to prepare. If you specify one of these products when calling [`/link/token/create`](/docs/api/link/#linktokencreate), and then you add another product immediately after the user has exited Link, you may experience higher latency when calling the secondary product because the first product is still being prepared.

You can avoid this problem by including the secondary product as either an Optional Product or Required if Supported Product when calling [`/link/token/create`](/docs/api/link/#linktokencreate), which allows Plaid to optimize how product data is prepared.

#### Impacts of product initialization on billing

If you initialize Link with a product that is billed under a [one-time or subscription fee](/docs/account/billing/) model (such as Auth, Identity, Income, Investments, Liabilities, or Transactions) you will be billed when the Item is created, even if you are not yet using the product's endpoints. There are exceptions:

- If you specify Auth, Identity, and/or Signal in the Optional Products array, you will not be billed for them until you use their endpoints (e.g., [`/auth/get`](/docs/api/products/auth/#authget)).
- You will not be billed for any products in the Additional Consented Products array until you use their endpoints.
- Investments has two different subscriptions that can be associated with it: Investments Holdings and Investments Transactions. Initializing Link with Investments adds only Investments Holdings; the Investments Transactions subscription will not be added until you call [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) for the first time on the Item. Note that Plaid will still pre-fetch investment transaction history when you initialize with Investments, even though you are not yet being billed for Investments Transactions.

To avoid unwanted charges, do not initialize Link with a one-time or subscription fee product in Production unless you are sure you are going to use that product. This instruction applies to the Products, Required if Supported Products, and Optional Products arrays, with the exception of Auth, Identity, and Signal Transaction Scores when used as Optional Products.

To learn more about how Plaid bills each product, see [Billing](/docs/account/billing/).

#### Impacts of product initialization on conversion

The products you specify in the Products array will determine which institutions and accounts are available in Link. If you specify multiple products in the Products array, only institutions and accounts supporting *all* of those products will be available. See the [accounts / product support](/docs/api/accounts/#account-type--product-support-matrix) matrix to learn which account types support which products.

To avoid overly restricting the institution and account list, you can use the [Required if Supported Products](/docs/link/initializing-products/#required-if-supported-products) or [Optional Products](/docs/link/initializing-products/#optional-products) features, or initialize Link with only your required products in the Products array and then call your other products' endpoints later.

##### Non-connection based Auth

Because non-connection-based Auth methods (Database Auth, Same-Day Micro-deposits, and Instant Micro-deposits) limit an Item's ability to be used with other products, initializing with any product other than `auth` in the `products` array will disable the use of non-connection-based Auth methods, which may impact conversion. To retain the ability to use these alternative Auth verification methods, include other products in the `optional_products` or `required_if_supported_products` array. For more details, see [Expanding Auth coverage](/docs/auth/coverage/).

##### Signal and the institution list

Because Signal Transaction Scores will function even if not all of the 80+ core attributes can be calculated, and because Balance is supported at all relevant financial institutions, including `signal` in the `products` array will never limit the list of available institutions.

If `signal` is included in the `products` array, `auth` must also be included in the `products` array.

#### Required if Supported Products

Using the `required_if_supported_products` parameter with a Plaid client library requires the following client library minimum versions: Node: 14.1.0, Python: 14.1.0, Ruby: 19.1.0, Java: 15.1.0, Go: 13.0.0.

You may want to require one or more secondary products when possible, but avoid excluding users whose financial institution doesn't support those products. To do this, specify the secondary products in the Required if Supported Products array.

When a product is specified in the Required if Supported Products array, Plaid will require that product if possible. If the institution the end user is attempting to link supports the secondary product, and the user grants Plaid access to an account at the institution that is compatible with it, Plaid will treat the product as though it was specified in the Products array. If the institution doesn't support the product, if the user doesn't have accounts compatible with it, or if the user has compatible accounts but doesn't grant access to them, Plaid will ignore the product. (To determine whether a product was successfully added to the Item, you can call [`/item/get`](/docs/api/items/#itemget) after obtaining your access token and read the Item's Products array.)

For example, if you initialize Link with both Auth and Identity in the Products array, a user will only be able to link an account if both Auth and Identity are supported by their institution, and they will not be able to use flows like Same-day Micro-deposits that require Auth to be initialized by itself. But if you specify Auth in the Products array and Identity in the Required if Supported Products array, users will be able to link checking and savings accounts at any institution that supports Auth, even if the institution does not support Identity. They will also be able to access flows that require Auth to be initialized by itself, like Same-day Micro-deposits.

If the institution *does* support both Auth and Identity, Plaid will attempt to initialize the Item with both products. If the institution has an OAuth flow where Identity access can be configured as a separate permission, and the user does not grant access to Identity, the Link attempt will error and the user will be prompted to retry the OAuth flow.

As another example, if Transactions and Auth are both specified in the Products array, users will only be able to link checking or savings accounts, because specifying Auth limits acceptable accounts to checking and savings accounts. However, if Transactions is specified in the Products array and Auth is specified in the Required if Supported Products array, users will be able to link other Transactions-compatible account types, such as credit cards and money market accounts. If they do link a checking or savings account, Plaid will fetch Auth data for that account.

#### Optional Products

Using the `optional_products` parameter with a Plaid client library requires the following client library minimum versions: Node: 17.0.0, Python: 17.0.0, Ruby: 23.0.0, Java: 18.0.0, Go: 17.0.0.

Optional Products is similar to Required if Supported Products, except that products specified in the Optional Products array are treated as *best-effort*. If an institution supports the selected products but they cannot be added to the Item (e.g., because of a temporary institution error impacting only some products, or because the user did not grant the required product-specific OAuth permissions), Item creation will succeed in spite of these failures.

Like the Required if Supported Products array, the Optional Products array does not affect institution filtering or account filtering, which means there are no guarantees around product data availability. For example, if you initialize Link with Transactions in the Products array and Auth in the Optional Products array, the Item is guaranteed to support Transactions, but it will also allow linking account types that are incompatible with Auth, like credit cards.

Unlike Required if Supported Products, Optional Products will not fail the Link attempt if the necessary OAuth permissions are not granted. If the user does not grant access to the data needed to support an Optional Product in the OAuth linking flow, they will not be required to go through the OAuth flow again to update their permissions. Instead, the Item will be successfully created and the Optional Product will be unavailable until you have prompted the user to fix their permissions via [Update Mode](https://plaid.com/docs/link/update-mode/).

If Auth is specified in the Optional Products array, Auth will only be added to the Item if the institution supports [Instant Auth](https://plaid.com/docs/auth/coverage/instant/).

If you specify Auth, Identity, and/or Signal in the Optional Products array, you will only be billed for these products if you call their endpoints.

#### When to use Required if Supported Products

If your only reason for omitting a product from the Products array is to increase the number of institutions covered, you should specify the product in the Required if Supported Products array rather than adding the product by calling its endpoint. Initializing both products during Link will make it easier to obtain the consent required at institutions that use OAuth, because the user will be prompted to fix their selections during the OAuth flow if they do not grant the required consent for all products. Initializing both products upfront can also reduce latency, because it allows Plaid to begin fetching data immediately and to optimize the order in which different types of data are requested.

For example, one common pattern for payments customers is to use Auth in the Products array, and Identity in the Required if Supported Products array. This means you will get Identity coverage for all accounts where it's available, but you won't block customers from linking their accounts if their institution doesn't support Identity.

Required if Supported Products is recommended over Optional Products if it is important to your use case that your user not be able to link an account if they've opted out of using certain products in the OAuth flow. For example, if you're accepting ACH payments, you might not want to let a user link their account while opting out of products used for risk mitigation, like Identity or Signal Transaction Scores.

#### When to use Optional Products

If you think that you may want to use Auth, Identity, and/or Signal Transaction Scores for an Item, but it isn't necessary for all users, you should specify the product in the Optional Products array, because doing so will allow Plaid to pre-load required data for improved performance. For Auth and Identity, this means lower latency. In the case of Signal Transaction Scores, this means improved accuracy, as Signal Transaction Scores is designed to optimize for latency over accuracy when a tradeoff is necessary.

As an example of when to use the Optional Products array, if your app primarily uses Transactions for budgeting but also uses Auth to transfer funds for a small percentage of users (e.g., paid users), you should specify Transactions in the Products array and Auth in the Optional Products array. This will ensure that all Items support Transactions, while also ensuring that Plaid will collect Auth data for as many users as possible without affecting Link conversion or increasing your Auth bill.

As an example of when *not* to use the Optional Products array, if your app primarily uses Auth for funds transfers but also offers an optional budgeting add-on that uses Transactions, you should specify Auth in the Products array and Transactions in the Additional Consented Products array, and then add Transactions to the Item by calling [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) only once your user opts in to using the budgeting tool. This will prevent you from being billed for Transactions before you are sure you need to use it.

#### Special cases for Link initialization

There are some exceptions to the ability to initialize consented products by calling their endpoints:

- Bank Income and Assets can both be initialized post-Link, but may require re-launching Link to obtain additional consent. See [Verifying Bank Income for existing Items](/docs/income/bank-income/#verifying-bank-income-for-existing-items) and [Adding Assets to existing Items](/docs/assets/#getting-an-asset-report-for-an-existing-item).
- Identity Verification and Payment Initiation cannot be initialized post-Link.
- When Auth is initialized post-Link, the main Instant Auth flow can be used, but the less common Instant Match, Automated Micro-deposits, or Same Day Micro-deposits flows cannot be initialized post-Link, as these are interactive flows that require user input.

Some products have limitations on being initialized alongside other products in the same Link session:

- Plaid Check products (such as Consumer Report) cannot be initialized in the same Link session as Assets or Income Verification.
- To use Same Day Micro-deposits, Instant Micro-deposits, Database Auth, Database Insights, or Database Match, Auth must be the only product in the `products` array when initializing Link. However, you can add Signal Transaction Scores or Identity to these Items via the `optional_products`, `required_if_supported_products`, or `additional_consented_products` fields.
- Income can only be initialized alongside other products if using Bank Income; Payroll Income and Document Income cannot be initialized alongside other products.
- Identity Verification and Payment Initiation cannot be initialized alongside any other products. If you need to use these products along with other Plaid products, your user will need to launch the Link flow a second time.

#### Recommendations for initializing Link with specific product combinations

##### Auth with Identity and/or Signal Transaction Scores

Initialize Link with Auth and Signal in the Products array and Identity in the Required if Supported Products array. Specifying Identity in the Required if Supported Products array will increase conversion while minimizing latency when calling Identity endpoints.

For most use cases, it is recommended to put Identity in Required if Supported Products instead of Optional Products. Using Optional Products for Identity allows customers at some OAuth institutions to opt out of sharing identity data and still complete the Link flow. Bad actors are likely to opt out of data sharing, reducing the usefulness of Identity as a fraud detection tool.

If you are willing to take on more risk, you can initialize with Identity in the Optional Products array instead, which will increase conversion and prevent you from being billed for Identity when you don't use it.

If Signal is in the Products array, Auth *must* also be in the Products array.

If you initialize with `signal` in the `additional_consented_products` array, you will need to call [`/signal/prepare`](/docs/api/products/signal/#signalprepare) before calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) for the first time on an Item in order to get the most accurate results. For more details, see [Adding Signal to existing Items](/docs/signal/creating-signal-items/#adding-signal-to-existing-items).

##### Auth with Assets

Initialize Link with Assets in the Products array and Auth in the Required if Supported Products array. This will allow your user to link any account type, not just checking or savings accounts, which is important for getting their full financial picture.

##### Transactions with Auth and/or Identity

If Transactions is the primary product, initialize with Transactions in the Products array, Auth in the Optional Products array, and Identity in the Required if Supported Products array. This will reduce latency when calling the Auth and/or Identity endpoints while maximizing institution coverage. It will also prevent you from being charged for Auth unless you end up calling its endpoints.

For most use cases, it is recommended to put Identity in Required if Supported Products instead of Optional Products. Using Optional Products for Identity allows customers at some OAuth institutions to opt out of sharing identity data and still complete the Link flow. Bad actors are likely to opt out of data sharing, reducing the usefulness of Identity as a fraud detection tool.

If you are willing to take on more risk, you can initialize with Identity in the Optional Products array instead, which will increase conversion and prevent you from being billed for Identity when you don't use it.

##### Transactions with Liabilities and/or Investments

For personal finance use cases, initialize Link with Transactions, with Liabilities and Investments in the Additional Consented Products array. Then call Liabilities or Investments endpoints after the Item has been linked, based on the account type that was linked. Because virtually all Items that support Liabilities or Investments also support Transactions, this will maximize conversion. If you call an endpoint that is not applicable to the Item (for example, if you call [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget) on a checking account, or [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) on a credit card account), the call will fail and you will not be billed for it.

##### Bank Income with Transactions

Initialize with Income Verification, add Transactions to Additional Consented Products, then call Transactions endpoints later, after the Item has been linked. Because Bank Income must fetch transaction data in order to verify the user's income, initializing with Transactions is not necessary. (Even though Bank Income uses transaction data, you will not be billed for Transactions on a Bank Income Item unless you add Transactions to the Item, either during or after Link.)

##### Bank Income with Assets

Initialize Link with both Income Verification and Assets in the Products array. Because institutions that support one of these products almost always support both of them, and because you will not be billed for Assets until you call the Assets endpoints, this is safe to do even if you aren't sure if you will use Assets. Assets can't be added to an Item post-Link without sending the user through Update Mode, so including Assets upfront will increase conversion.

If you already have an Item enabled with one of these products and want to add the other: see the instructions on [Verifying Bank Income for existing Items](/docs/income/bank-income/#verifying-bank-income-for-existing-items) to add Bank Income to the Item. See the instructions on [Getting an Asset Report for an existing Item](/docs/assets/#getting-an-asset-report-for-an-existing-item) to add Assets.

##### Layer

For Layer, product initialization settings are configured in Layer templates via the Dashboard.

Products set to **Always require** in the template impact Layer eligibility -- users will only be eligible for Layer if they have saved accounts at institutions that are compatible with these products. If a product is configured as **Require if supported**, **Optional**, **Additional consented**, or **Disabled**, users will still be eligible for Layer even if they do not have saved Items that are compatible with the product.

If a user adds new accounts or institutions during Layer, the same restrictions impact which institutions and accounts they can add (e.g. if Auth is required, the user will only be able to add depository accounts that support ACH).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
