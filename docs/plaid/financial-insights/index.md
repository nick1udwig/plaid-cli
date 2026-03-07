---
title: "Financial Insights | Plaid Docs"
source_url: "https://plaid.com/docs/financial-insights/"
scraped_at: "2026-03-07T22:04:54+00:00"
---

# Financial insights products

#### Review and compare solutions

This page provides overviews of Plaid's financial insights solutions to help you find the right one for your needs.

These products serve use cases such as personal finance and budgeting, loan payback and debt management, expense and accounting applications, and wealth management and investing. Note that none of these products may be used as part of a credit or underwriting decisioning process; for underwriting use cases, see [credit underwriting products](/docs/underwriting/).

#### Transactions

[Transactions](/docs/transactions/) gets transaction data for an end user's account, such as a checking, savings, or credit card account. Transaction data includes details such as date, merchant, and category; for some transactions, it will also include location data and a merchant name that has been enhanced for greater human readability. Any transaction data fields not provided to Plaid directly by the financial institution are derived via Plaid's ML-powered transactions enrichment engine. New and updated transactions are typically extracted between one and four times per day, depending on the institution.

##### Transactions Refresh

[Transactions Refresh](/docs/transactions/) is an optional add-on for Transactions that can be used to trigger an on-demand, real-time update of transactions data. Transactions Refresh provides on-demand updates only; Plaid does not offer a real-time feed of transactions updates.

##### Recurring Transactions

[Recurring Transactions](/docs/transactions/#recurring-transactions), available in the US, CA, and UK, is an optional add-on for Transactions that can be used to identify recurring transaction streams. Recurring Transactions is typically used by financial management apps to power features such as budget analysis and helping users identify and cancel unwanted subscriptions.

#### Enrich

[Enrich](/docs/enrich/) allows you to provide your own transactions data and uses the same ML-powered enrichment engine used by the Transactions product to enhance this data with more details, such as category, location, and merchant name. Unlike most other Plaid products, Enrich does not use Link to connect to an end user's financial account; Enrich is designed for customers who already have transactions data from a non-Plaid source and would like to enhance it to provide better insights to their customers.

#### Liabilities

[Liabilities](/docs/liabilities/) gets details about an end user's debt, including interest rate, outstanding balance, repayment schedule, and details of the most recent repayment. Supported loan types are credit cards, private student loans, mortgages, and PayPal loans. Liabilities does not provide transaction history. For full transaction and loan details, use both Liabilities and Transactions together.

#### Investments

[Investments](/docs/investments/) gets details about an end user's investment account, such as a brokerage account or 401(k). Investments comes with access to two main features: Investments Holdings, which provides details on the specific assets held, including type, description, value, cost basis, and acquisition date; and Investments Transactions, which provides details on transactions (such as trades, transfers, or dividends) within an investment account.

##### Investments Refresh

[Investments Refresh](/docs/investments/) is an optional add-on for Investments that can be used to trigger an on-demand, real-time update of investments data.

#### Transactions, Investments, Liabilities, and Enrich comparison table

|  | Transactions | Investments | Liabilities | Enrich |
| --- | --- | --- | --- | --- |
| Summary | Details on transactions | Details on investments and investments transactions | Details on loans, payments, balances, and interest | Enrich your existing transactions data |
| Main supported account types | Checking, savings, credit cards | Brokerage accounts, including retirement | Credit cards, private student loans, mortgages | Checking, savings, credit cards |
| Supports brokerage accounts | No | Yes | No | No |
| Provides transaction details | Yes | Yes | No (only most recent payment) | N/A |
| Provides ML-enhanced transaction descriptions and details | Yes | No | No | Yes |
| Provides interest rates | No | No | Yes | No |
| Typical update frequency | 1-4 times per day | Once per day, after market close | Once per day | N/A |
| Optional Refresh add-on for on-demand updates | Yes | Yes | No | N/A |
| Supported countries | US, CA, UK, Europe | US | US, CA (CA coverage limited) | US, CA |
| Billing plans available | Pay-as-you-go or 12-month contract | Pay-as-you-go or 12-month contract | Pay-as-you-go or 12-month contract | Pay-as-you-go or 12-month contract |

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
