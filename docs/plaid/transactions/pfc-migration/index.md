---
title: "Transactions - Personal Finance Category migration guide | Plaid Docs"
source_url: "https://plaid.com/docs/transactions/pfc-migration/"
scraped_at: "2026-03-07T22:05:22+00:00"
---

# Personal Finance Category migration guide

#### Learn how to migrate from Plaid's legacy categories to the new personal finance categories

#### Migrating from Personal Finance Categories v1 to PFCv2

##### What's in PFCv2

In December 2025, Plaid released a new version of Personal Financial Categories, called PFCv2, including both improved transaction categorization accuracy and a new, more granular transaction schema with additional categories specifically useful for Earned Wage Access (EWA) use cases, with six additional income subcategories, six additional loan disbursement subcategories, and three additional loan repayment subcategories. Additional subcategories have also been added for bank fees. [See a full comparison of PFCv2 and PFCv1](https://plaid.com/documents/pfc-taxonomy-all.csv).

##### How to migrate

Opting in to PFCv2 is required for existing customers to receive the new, more granular subcategories. All customers, regardless of whether they opt in to PFCv2, will automatically receive the accuracy improvements, with the exception of a small subset of customers in the EWA industry, who will not be automatically opted in to receive the accuracy improvements until 2027 unless they opt in to PFCv2. All new customers enabled for Transactions after December 3, 2025 will be opted in to PFCv2 by default.

To opt in to PFCv2, existing customers should set `personal_finance_category_version` to `v2` in the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), [`/transactions/get`](/docs/api/products/transactions/#transactionsget), [`/transactions/enrich`](/docs/api/products/enrich/#transactionsenrich), or [`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget) request.

Before migrating, make sure your app can accept the new categories and subcategories used by PFCv2. The work involved will depend on your integration; for some customers, this may not require any action at all. If you maintain logic that maps PFCs to your own internal categories, you will need to update it to incorporate the new values. See the [PFCv2 to PFCv1 mapping specification](https://docs.google.com/spreadsheets/d/e/2PACX-1vQjAYBlds-bOGY6DfBW_Fl1lPD7xdwq7RjWWsVshwzWjUhIAxkIlroacVQKtIp-8xXY5DGHcsCq7Mse/pubhtml). You can also download the CSV [here](https://plaid.com/documents/pfc-taxonomy-all.csv).

##### Backfilling historical transactions with PFCv2

To apply the PFCv2 categories to existing transactions that you already fetched or that were already extracted prior to the migration, you can re-call any of the endpoints while passing in the `options.personal_finance_category_version: "v2"` parameter. The improvements to accuracy and confidence will not be backfilled to these pre-existing transactions; only the new taxonomy will be applied.

#### Migrating from legacy categories to Personal Finance Categories

This section is only relevant to customers who are still using Plaid's legacy `category` system, which was superseded by Personal Finance Categories in September 2023. Customers who migrate will receive PFCv1 by default and must set `personal_finance_category_version` to `v2` in the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync), [`/transactions/get`](/docs/api/products/transactions/#transactionsget), [`/transactions/enrich`](/docs/api/products/enrich/#transactionsenrich), or [`/transactions/recurring/get`](/docs/api/products/transactions/#transactionsrecurringget) request to migrate directly to the latest Personal Finance Categories, PFCv2.

Personal finance categories (PFCs), returned by all Transactions endpoints as of September 2023, provide more meaningful categorization at greater accuracy compared to the legacy category fields. All Transactions implementations are recommended to use PFCs rather than the legacy taxonomy. The PFC taxonomy is composed of fewer, better-defined categories based on customer needs and what users expect to see in a personal finance management app. Although PFCs are most relevant for personal finance use cases, they are not limited to those use cases.

PFCs are composed of the following fields:

- `personal_finance_category.primary`: A high level category that communicates the broad category of the transaction
- `personal_finance_category.detailed`: A granular category conveying the transaction's intent; this field can also be used as a unique identifier for the category
- `personal_finance_category.confidence_level`: A description of how confident we are that the provided categories accurately describe the transaction's intent
- `personal_finance_category_icon_url`: The URL of an icon associated with the primary personal finance category

See an example of these two types of categorization below:

Legacy category:

```
...
"category": [
  "Shops",
  "Computers and Electronics"
],
"category_id": "19013000",
...
```

Personal finance category:

```
...
"personal_finance_category": {
  "primary": "GENERAL_MERCHANDISE",
  "detailed": "GENERAL_MERCHANDISE_ELECTRONICS",
  "confidence_level": "VERY_HIGH"
},
"personal_finance_category_icon_url": "https://plaid-category-icons.plaid.com/PFC_GENERAL_MERCHANDISE.png",
...
```

See the [taxonomy CSV file](https://plaid.com/documents/pfc-taxonomy-all.csv) for a full list of PFCs.

Plaid has no plans to remove the legacy categories from the API response at this time. However, we are no longer developing improvements to those fields and strongly recommend using PFCs for the most accurate and up-to-date categorization.

Note that PFCs are not present in Assets endpoints.

##### Updating your internal categorization logic

Some customers maintain logic that translates Plaid's legacy categories into their own internal categories. In order to facilitate the updating of this translation layer, see the [mapping JSON file](https://plaid.com/documents/transactions-personal-finance-category-mapping.json) for a suggested 1-to-many mapping of our most common legacy categories to possible PFCs. Please note that this file is intended for general taxonomy-to-taxonomy mapping, and individual transactions may be assigned a PFC unaligned with its legacy category as we continue to refine the accuracy of our PFC predictions.

##### Backfilling historical transactions with PFCs

You may already have historical transactions stored with only the legacy category present. If you are interested in backfilling PFCs for those transactions, you can re-fetch those transactions with a populated PFC using the [`/transactions/get`](/docs/api/products/transactions/#transactionsget) or [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) endpoints.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
