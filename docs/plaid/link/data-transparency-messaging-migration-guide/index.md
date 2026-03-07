---
title: "Link - Data Transparency Messaging migration | Plaid Docs"
source_url: "https://plaid.com/docs/link/data-transparency-messaging-migration-guide/"
scraped_at: "2026-03-07T22:05:04+00:00"
---

# Data Transparency Messaging Guide

#### How to enable and manage Data Transparency Messaging for your application

As of October 31, 2024, all new Plaid Inc. customers launching to end users in the US and/or Canada are automatically enrolled for Data Transparency Messaging and must [select a use case](https://dashboard.plaid.com/link/data-transparency-v5) to use Link in Production, including Limited Production. This does not apply customers who use only products that don't connect to financial institutions, like Identity Verification, or to customers using only Consumer Report or other Plaid Check products.

#### Introduction

Plaid has introduced Data Transparency Messaging (DTM) into the Link flow to help you stay in compliance with the [1033 rule](https://plaid.com/resources/compliance/section-1033/). This experience provides end users (consumers) with a greater understanding of the types of data that they share with you and Plaid. When using DTM, a user is informed of the specific data types that you are requesting and the reason that you are requesting them (use cases). If you want access to additional data from a user after their initial Link, or to use the data for additional use cases, they must consent to sharing that data through a separate consent flow.

This guide covers the implementation changes required to implement the Data Transparency Messaging feature.

![Plaid's Data Transparency Messaging: Log in at Gingham Bank, disclose account info, select accounts, with product-specific disclosures.](/assets/img/docs/dtm.png)

DTM introduces product-specific disclosures above the "Continue" buttons on key screens such as the OAuth handoff pane and "Account Select" pane (pictured). If the user clicks the Learn More link within the disclosure, the "Why is this needed" half-pane appears.

##### Who is affected

Data Transparency Messaging (DTM) will be required for all Plaid customers, unless you meet one of the criteria below:

- You use only products that don't connect to financial institutions: Identity Verification, Beacon, Monitor, or Enrich.
- You serve end users only in Europe. If you serve end users in Canada but not the United States and would like to opt out of Data Transparency Messaging, contact your Account Manager.
- You use only Plaid Check products (i.e. Consumer Report) rather than Plaid Inc. products.

##### What's new

- The API will only return data for products that users have consented to through Link. In addition to the existing `products` field, we have added a new configuration field called `additional_consented_products` that can be used to gather consent for additional products.
- Link can be configured to display DTM. The language will appear as a footer on the OAuth pane (shown to consumers before they are directed to an institution’s OAuth flow), or the Account Select pane for sessions that do not go through OAuth.
- If enabled via the Plaid Dashboard, DTM must be configured to include at least one use case, which explains to users the purpose of the data being requested. You may choose up to 3 use cases for each Link customization that will be shown in Link to your users. You can select these use cases via the [Dashboard Link Customization page](https://dashboard.plaid.com/link/data-transparency-v5).
- Products that require data not consented to by users through Link cannot be accessed without going through [update mode](/docs/link/update-mode/).

Data Transparency Messaging is available to all customers who are enabled for US or Canada on an opt-in basis. New customers will be automatically enabled for Data Transparency Messaging, and some existing customers have been automatically enabled for Data Transparency Messaging. For existing customers who have not yet been enabled for Data Transparency Messaging, Plaid will provide more details on the enablement timeline once it has been finalized.

#### Migration overview

1. Review your API integration and, if necessary, [populate the `additional_consented_products` field](/docs/link/data-transparency-messaging-migration-guide/#additional-consented-products) with products you want to gather consent for. Every product that you currently use should be included in either the `products`, `optional_products`, `required_if_supported_products` or `additional_consented_products` fields. For more information on the differences between these fields, see [Initializing products](/docs/link/initializing-products/#overview).
2. Update the use cases for each Link customization via the [Dashboard](https://plaid.com/docs/link/customization/).
3. (Optional) [Update Link customizations](/docs/link/data-transparency-messaging-migration-guide/#updating-link-customizations) to enable DTM via the [Dashboard](https://plaid.com/docs/link/customization/).
4. (Future) Configure [update mode](/docs/link/update-mode/) to renew consent.

#### Additional consented products

Previously, as long as you were approved and enabled for a product during your Plaid Dashboard Production Request process, you could request and start using it by calling the corresponding product API route on any Item. Going forward, any Item that goes through Link with DTM enabled restricts your access to new product data unless the Item has the required consent; if you do not have the required consent, you will receive the error [`ADDITIONAL_CONSENT_REQUIRED`](/docs/errors/invalid-input/#additional_consent_required). Additional consent will need to be collected through update mode.

If you need to use a product but don't want Plaid to try to extract that data immediately, you can create a Link token with [`additional_consented_products`](/docs/api/link/#link-token-create-request-additional-consented-products) specified. These products will be shown in Link as part of DTM, but will not be fetched or billed until you request them. Specifying `additional_consented_products` will not have any impact on Items created before DTM is enabled, but it is recommended that you start using this field so you can be ready when DTM is enabled for your Link flows.

The products in the `additional_consented_products` field may not have any overlap with products listed in the `products`, `optional_products` or `required_if_supported_products`. The `products` array must contain at least one product. In the DTM language, Link will show data scopes corresponding to the union of `products`, `optional_products`, `required_if_supported_products`, and `additional_consented_products`.

Products placed in the `additional_consented_products` field do not have any impact on institution filtering or account subtype filtering. For example, if `auth` is an `additional_consented_product`, users will be allowed to select any account, not just checking/savings accounts.

When your existing Items are migrated to DTM, Items in the US or Canada created without DTM will have the following migration applied: the `billed_products`, along with any additional products that use the same data scopes, will be automatically added to the `consented_products` field for those Items. You will be notified when your Items are scheduled to be migrated.

#### Updating Link customizations

To opt into DTM, go to your [Link customization](https://dashboard.plaid.com/link/privacy-interstitial) or create a new one (you may duplicate an existing customization by using the [duplicate feature](https://plaid.com/docs/link/customization/#link-customization-overview)), click on the Data Transparency section, where you can customize the use case. You must have at least one selected use case to save the customization. You can preview the primary footer language and the detailed panel. This language will appear on the OAuth pane for sessions that go through OAuth, and on Account Select otherwise.

Use cases will only be shown in Link if the end user clicks the "Learn more" link within the DTM disclosure text.

| Category | Use cases |
| --- | --- |
| Payments | Send and receive money, Pay your bills, Make a purchase online, Facilitate business-to-business payments, Fund your account |
| Identity verification and fraud | Verify your identity and prevent fraud, Verify your account, Protect against fraud |
| Personal / Business finance management | Track and manage your finances, Prepare your taxes, Get rewards, Invest your money, Do business accounting and tax preparation, Prepare and categorize invoices, Manage employee expense reporting, Track, manage and build your credit, Access your paycheck sooner, Pay down debt |
| Credit underwriting | Get considered for a loan, Verify your income |

There is no restriction enforced on which use case(s) you select beyond the limit of three use cases per customization. For example, you may still select a payments-related use case even if you are not Production-enabled for any payments-related products and not initializing Link with any payments-related products.

In the Production environment, the use cases you select will also visible to consumers via the Plaid Portal ([my.plaid.com](https://my.plaid.com/)) and to Plaid financial institution partners.

Once you've selected how you want to present DTM to users and selected at least one use case, click Publish. Your changes will go into effect the next time you initialize Link with this customization. Only newly created Items created with the updated customization will have the new API behavior; existing Items will not be affected.

If you want to prepare for the DTM automatic enablement date without enabling DTM earlier than necessary, you can add use cases to your [Link customization](https://dashboard.plaid.com/link/privacy-interstitial) without enabling DTM.

If you currently use additional products after Link that you do not initialize via the `products` field, it is recommended that you update your [`/link/token/create`](/docs/api/link/#linktokencreate) code to include `additional_consented_products` before enabling this feature in the Plaid Dashboard to avoid losing access to the additional API endpoints on your new Items.

#### Authorization records

You may need to review authorization records to meet audit requirements or to make decisions about collecting additional consent.

Current authorization details such as an Item's consented products and consented use cases are available via the [`/item/get`](/docs/api/items/#itemget) endpoint. These are updated in near real time and may be used to configure update mode to request additional consent from a user.

Historical authorization records are available via [`/consent/events/get`](/docs/api/consent/#consenteventsget). These records will be retained and available for at least 3 years (though the earliest records will begin in November 2024).

Authorization details and historical authorization records will only be available for Items that have been enabled for Data Transparency Messaging.

#### API changes

For Items that have gone through Link with Data Transparency Messaging enabled, there will be new `consented_products` and `consented_data_scopes` fields returned by the [`/item/get`](/docs/api/items/#itemget) endpoint.

You will see the error code [`ADDITIONAL_CONSENT_REQUIRED`](/docs/errors/invalid-input/#additional_consent_required) if you do not have consent to a product you are trying to request. This is the same error you would get if you are not enabled for the product.

To resolve `ADDITIONAL_CONSENT_REQUIRED` errors, you will need to send the user through [update mode](/docs/link/update-mode/#requesting-additional-consented-products).

#### Data scopes and consent

The products you request for Link Initialization are not mapped 1:1 with the different data scopes displayed. Our goal is to allow consumers to understand what types of data they are permissioning, and product names alone can be confusing to consumers. Because data scopes can be broad, in certain cases, you will get consent for more than just the products you initialize Link with. For example, the Assets product requires consent to both the "Transactions" and the "Account Holder Information" data scopes, so the Transactions and Identity products will both be consented to by default. To understand which products you have consent to access for a given Item, you can use the `consented_products` field on the [`/item/get`](/docs/api/items/#itemget) or product endpoints as detailed above.

Note that the data scopes used by Plaid should not be confused with OAuth data scopes. Some institutions that use OAuth define their own data scopes, and users at OAuth institutions may be required to grant consent to both Plaid's data scopes and institution-specific OAuth data scopes during the Link flow. For more information on OAuth data scopes, see [Scoped consent](/docs/link/oauth/#scoped-consent) in the OAuth documentation.

##### Data scope descriptions

|  |  |
| --- | --- |
| Account and balance info | May include account details such as account name, type, description, balances, and masked account number. |
| Contact info | May include account owner name, email, phone, and address. |
| Account and routing number | Includes account number and financial institution routing numbers. |
| Transactions | May include transaction details such as amounts, dates, price, location, spending categories and descriptions, and insights like recurring transactions. |
| Credit and loans | May include details about your credit and loan accounts, such as balance, payment dates and amounts due, credit limit, repayment status, interest rate, loan terms, and originator info. |
| Bank Statements | May include account info such as activity and usage, and contact details. |
| Investments | May include info about investment accounts, such as securities details, quantity, price, and transactions. |
| Risk info | May use account info such as activity and usage, contact details, transaction-related information, and Plaid connection history to provide an assessment of fraud and/or minimize overdrafts |

##### Data scopes by product

| Plaid Product | Data scopes required |
| --- | --- |
| Account Risk (beta) | Account and balance info, Contact info, Transactions, Risk info |
| Assets | Account and balance info, Contact info, Transactions |
| Auth: Instant Auth or Instant Match | Account and balance info, Contact info, Account and routing number |
| Auth: Database Auth, Database Insights, or Database Match | Account and routing number |
| Auth: Same-Day Micro-deposit or Instant Micro-deposit | Account and routing number |
| Auth: Automatic Micro-deposit | Account and balance info, Contact info, Transactions, Account and routing number |
| Balance | Account and balance info |
| Bank Income | Account and balance info, Contact info, Transactions |
| Consumer Report by Plaid Check | Account and balance info, Contact info, Transactions |
| Statements | Account and balance info, Contact info, Statements |
| Identity (includes Identity Match) | Account and balance info, Contact info |
| Investments | Account and balance info, Contact info, Investments |
| Liabilities | Account and balance info, Contact info, Credit and Loans |
| Signal Transaction Scores | Account and balance info, Contact info, Transactions, Risk info |
| Transactions (includes Recurring Transactions) | Account and balance info, Contact info, Transactions |
| Transfer | Account and balance info, Contact info, Account and routing number |
| Investments Move (fka Investments Auth) (beta) | Account and balance info, Contact info, Account and routing number, Investments |

Account and balance info and Contact info are almost always requested; there are some exceptions for products like Database Match or Database Insights where Plaid does not pull data from the bank, and does not require this information.

Plaid will disclose all the underlying data scopes needed to deliver the product. For example, Signal Transaction Scores uses Transactions data as one input to derive the score, so we disclose that the data scope Transactions was accessed, even if you do not request the product Transactions in `products` or `additional_consented_products`. In this case, Plaid will record the user’s consent to the Transactions data scope. If you request a product requiring the Transactions data scope in the future, you will not need to use update mode. (But if you are adding a new use case as well, you will need to use update mode to request consent for that use case, even if you already have the required product consent.)

Some Plaid products do not have any associated bank data scopes because they do not rely on access to bank data. For example, Identity Verification, Document Income, and Payroll Income do not rely on bank data.

#### Requesting consent for additional products or use cases in update mode

If you would like to start collecting product data for an Item that you did not collect enough consent for during initial creation, you are able to send that Item through update mode to collect new consent. For details, see [Requesting additional consented products](/docs/link/update-mode/#requesting-additional-consented-products) and [Requesting additional use cases](/docs/link/update-mode/#requesting-additional-use-cases) in the update mode documentation.

#### Automatic migration path

If you do not choose to follow any of the steps in this guide, your eligible Link sessions (those including a country code of `US` and/or `CA`) will be automatically enabled for Data Transparency Messaging. No date has been finalized for this automatic migration; Plaid will provide details when it has been determined.

When your team is automatically enabled for Data Transparency Messaging, the following migrations will occur:

- Plaid will enable DTM for your Items and Link flows. If you have not customized your use cases as described in [Updating Link customizations](/docs/link/data-transparency-messaging-migration-guide/#updating-link-customizations), the use cases presented in the Link UI will be automatically populated based on information you have previously provided to us (e.g. in the Production Request Form).
- If you do not specify `additional_consented_products` in your [`/link/token/create`](/docs/api/link/#linktokencreate) call, Plaid will perform an assessment of your historical product add behavior. If, in the last 6 months before the enforcement date, you had billable activity for a product, and that product is not included in the `products`, `required_if_supported_products`, or `optional_products` arrays, Plaid will automatically update your Link session to request consent for that product as though that product had been added to the `additional_consented_products` array. To opt out of this behavior, contact your Account Manager.

DTM is mandatory for all customers who are enabled for Plaid in the US. While Canada-based Link sessions are not subject to 1033 enforcement, they will receive DTM to provide enhanced consumer transparency. If you are enabled for Canada but not for the US and want to opt out of automatic DTM enablement, contact your Account Manager.

#### Minimum client library version for using additional consented products

- [Python](https://github.com/plaid/plaid-python/blob/master/CHANGELOG.md) - 9.3.0
- [Node](https://github.com/plaid/plaid-node/blob/master/CHANGELOG.md) - 10.4.0
- [Ruby](https://github.com/plaid/plaid-ruby/blob/master/CHANGELOG.md) - 15.5.0
- [Java](https://github.com/plaid/plaid-java/blob/master/CHANGELOG.md) - 11.3.0
- [Go](https://github.com/plaid/plaid-go/blob/master/CHANGELOG.md) - 3.4.0

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
