---
title: "Credit and Underwriting | Plaid Docs"
source_url: "https://plaid.com/docs/underwriting/"
scraped_at: "2026-03-07T22:05:27+00:00"
---

# Underwriting products

#### Review and compare underwriting solutions

This page provides overviews of Plaid's underwriting and lending products to help you find the right one for your needs.

#### Consumer Report by Plaid Check

[Plaid Check](/docs/check/) products include scores and analyses that you can use directly for underwriting purposes, along with transaction data categorized into lender-relevant categories. The Consumer Report provides raw data similar to Assets, with additional analysis such as overall inflow and outflow counts and averages. Consumer Report is enhanced by two optional add-ons: Income Insights and Partner Insights. Income Insights provides historical monthly and annual gross income data as well as future income predictions. Partner Insights provides pre-generated cashflow underwriting risk scores and analysis in partnership with Prism Data.

Consumer Report also provides an optional Home Lending Report, a FCRA-compliant and Fannie Mae Day 1 Certainty-approved report that delivers a comprehensive view of a borrower's assets, tailored for mortgage lenders. Home Lending Report is currently in early availability; for more details, [contact Sales](https://plaid.com/check/income-and-underwriting/#contact-form).

Plaid Check can be used together with most Plaid Inc. products, such as Layer or Balance, but cannot be used with Assets or Income.

[See a live, interactive demo](https://plaid.coastdemo.com/share/67ed96db47d17b02bdfc4a88?zoom=100) of a Plaid-powered underwriting flow, using Consumer Report by Plaid Check.

#### Assets

Most new customers should use Consumer Report rather than Assets. Assets is currently recommended only for use cases not supported by Consumer Report, such as underwriting outside the US.

[Assets](/docs/assets/) gets details on an end user's assets and balances, including transaction history, with inflows and outflows categorized into lender-relevant categories. Asset reports can be provided either in JSON format for programmatic use or PDF format for human-powered verification workflows. Assets is optimized for verification and underwriting for traditional loan origination.

#### Income

Most new customers should use the [Consumer Report Income solution](/docs/check/) instead of Income. Income is currently recommended only for use cases not supported by Consumer Report, such as supporting end users outside the US, or Payroll or Document Income based flows.

[Income](/docs/income/) gets income details suitable for cash flow underwriting and income or employment verification. Where Assets provides all transactions and requires you to analyze them yourself, Income identifies specific income streams and provides details such as amount and frequency. Income supports both gig and payroll income types and can verify income based on deposits to a linked account (Bank Income), user-uploaded documentation such as W-2s or pay stubs (Document Income), or data imported from a linked payroll provider (Payroll Income).

[See a live, interactive demo](https://plaid.coastdemo.com/share/66fb0a180582208ffa82103e?zoom=100) of a Plaid-powered underwriting flow, using Assets and Income.

#### Consumer Report, Assets, and Income comparison

Consumer Report by Plaid Check is recommended for most new customers. Assets and Income may be appropriate for customers who have use cases that aren't supported by Plaid Check, such as supporting non-US-based customers. Assets and Income can't be used together in the same Link flow with Consumer Report by Plaid Check.

Assets provides a broader financial picture as part of traditional loan underwriting and asset verification, while Income focuses on providing rich income data optimized for cash flow underwriting and employment verification and supports a variety of data sources.

Consumer Report by Plaid Check provides data similar to Assets and Income, plus risk scoring and income insights. With its off-the-shelf insights, Consumer Report is ideal for customers who want to do cash flow underwriting without building their own credit risk attributes and scores from raw bank transaction data. When used with the optional asset verification feature, it's also the best solution for mortgage underwriting.

|  | Assets | Income | Consumer Report by Plaid Check |
| --- | --- | --- | --- |
| Summary | Provides financial data for underwriting decisions and asset verification | Provides analyzed and parsed income streams for cash flow underwriting decisions and income/employment verification | Provides financial data for underwriting decisions, asset verification, and cash flow underwriting, plus underwriting risk analysis and optional Home Lending Report for mortgage underwriting |
| Provides detailed balance and transaction history that can be used for underwriting purposes | Yes | Income transactions only | Yes |
| Provides off-the-shelf risk scores for cash flow underwriting | No | No | Yes, with Partner Insights |
| Identifies income streams and provides historical income dates and amounts | No | Yes | Yes, with Income Insights |
| Provides future income predictions and historical income summaries | No | No | Yes, with Income Insights |
| Provides reports in either PDF or JSON format | Yes | Yes | Yes |
| Day 1 Certainty report provider for Fannie Mae Desktop Underwriting | Yes | No | Yes |
| Provides identity data from linked financial institution accounts to reduce fraud risk | Yes | Yes, with Bank Income | Yes |
| Supports data from linked financial institution accounts | Yes | Yes, with Bank Income | Yes |
| Supports data from payroll providers | No | Yes, with Payroll Income | No |
| Supports data from user-uploaded documents | No | Yes, with Document Income | No |
| Can be used in a single Link session with other Plaid products such as Auth | Yes | Yes, with Bank Income | No |
| Supported countries | US, CA, UK, Europe | US, CA, UK | US |
| Billing plans available | Pay-as-you-go or 12-month contract | Pay-as-you-go or 12-month contract | Pay-as-you-go or 12-month contract |

#### Statements

[Statements](/docs/statements/) (US only) gets exact copies of PDF statements as provided by the financial institution, to streamline your verification processes that require statement copies. Because Statements does not cover all major or long-tail institutions, it should be used with a fallback option in place if data is not available. Statements supports only depository (e.g. checking, savings, or money market) accounts.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
