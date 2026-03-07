---
title: "Liabilities - Introduction to Liabilities | Plaid Docs"
source_url: "https://plaid.com/docs/liabilities/"
scraped_at: "2026-03-07T22:05:01+00:00"
---

# Introduction to Liabilities

#### Access data for private student loans, mortgages, and credit cards using the Liabilities product.

Get started with Liabilities

[API Reference](/docs/api/products/liabilities/)[Quickstart](/docs/quickstart/)

#### Overview

Plaid's Liabilities product allows you to access information about a user's debts. A common application is personal financial management tools to help customers manage or refinance debt. Liabilities is supported in the US and Canada (Canada coverage is limited).

#### Liabilities data

With Liabilities, you can view account information for credit cards, PayPal credit accounts, private student loans, and mortgages in the US. Available information includes balance, next payment date and amount, loan terms such as duration and interest rate, and originator information such as the origination date and initial loan amount. Liabilities data is refreshed approximately once per day, and the latest data can be accessed by calling [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget).

Note that liabilities data does not contain detailed transaction history for credit card accounts. For credit card account transactions, use the [Transactions](/docs/transactions/) product.

Sample Liabilities data: credit card account

```
"credit": [
  {
    "account_id": "dVzbVMLjrxTnLjX4G66XUp5GLklm4oiZy88yK",
    "aprs": [
      {
        "apr_percentage": 15.24,
        "apr_type": "balance_transfer_apr",
        "balance_subject_to_apr": 1562.32,
        "interest_charge_amount": 130.22
      },
      {
        "apr_percentage": 27.95,
        "apr_type": "cash_apr",
        "balance_subject_to_apr": 56.22,
        "interest_charge_amount": 14.81
      },
      {
        "apr_percentage": 12.5,
        "apr_type": "purchase_apr",
        "balance_subject_to_apr": 157.01,
        "interest_charge_amount": 25.66
      },
      {
        "apr_percentage": 0,
        "apr_type": "special",
        "balance_subject_to_apr": 1000,
        "interest_charge_amount": 0
      }
    ],
    "is_overdue": false,
    "last_payment_amount": 168.25,
    "last_payment_date": "2019-05-22",
    "last_statement_issue_date": "2019-05-28",
    "last_statement_balance": 1708.77,
    "minimum_payment_amount": 20,
    "next_payment_due_date": "2020-05-28"
  }
]
```

Sample Liabilities data: student loan account

```
"student": [
  {
    "account_id": "Pp1Vpkl9w8sajvK6oEEKtr7vZxBnGpf7LxxLE",
    "account_number": "4277075694",
    "disbursement_dates": [
      "2002-08-28"
    ],
    "expected_payoff_date": "2032-07-28",
    "guarantor": "DEPT OF ED",
    "interest_rate_percentage": 5.25,
    "is_overdue": false,
    "last_payment_amount": 138.05,
    "last_payment_date": "2019-04-22",
    "last_statement_issue_date": "2019-04-28",
    "loan_name": "Consolidation",
    "loan_status": {
      "end_date": "2032-07-28",
      "type": "repayment"
    },
    "minimum_payment_amount": 25,
    "next_payment_due_date": "2019-05-28",
    "origination_date": "2002-08-28",
    "origination_principal_amount": 25000,
    "outstanding_interest_amount": 6227.36,
    "payment_reference_number": "4277075694",
    "pslf_status": {
      "estimated_eligibility_date": "2021-01-01",
      "payments_made": 200,
      "payments_remaining": 160
    },
    "repayment_plan": {
      "description": "Standard Repayment",
      "type": "standard"
    },
    "sequence_number": "1",
    "servicer_address": {
      "city": "San Matias",
      "country": "US",
      "postal_code": "99415",
      "region": "CA",
      "street": "123 Relaxation Road"
    },
    "ytd_interest_paid": 280.55,
    "ytd_principal_paid": 271.65
  }
]
```

Sample Liabilities data: mortgage account

```
"mortgage": [
  {
    "account_id": "BxBXxLj1m4HMXBm9WZJyUg9XLd4rKEhw8Pb1J",
    "account_number": "3120194154",
    "current_late_fee": 25.0,
    "escrow_balance": 3141.54,
    "has_pmi": true,
    "has_prepayment_penalty": true,
    "interest_rate": {
      "percentage": 3.99,
      "type": "fixed"
    },
    "last_payment_amount": 3141.54,
    "last_payment_date": "2019-08-01",
    "loan_term": "30 year",
    "loan_type_description": "conventional",
    "maturity_date": "2045-07-31",
    "next_monthly_payment": 3141.54,
    "next_payment_due_date": "2019-11-15",
    "origination_date": "2015-08-01",
    "origination_principal_amount": 425000,
    "past_due_amount": 2304,
    "property_address": {
      "city": "Malakoff",
      "country": "US",
      "postal_code": "14236",
      "region": "NY",
      "street": "2992 Cameron Road"
    },
    "ytd_interest_paid": 12300.4,
    "ytd_principal_paid": 12340.5
  }
]
```

#### Payment history

The information returned by a [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget) request contains recent payment information, such as the date and amount of the most recent payment. To view further payment history, you can use Plaid's [Transactions](/docs/transactions/) product.

#### Liabilities webhooks

Plaid checks for updated Liabilities data approximately once per day and uses webhooks to inform you of any changes so you can keep your app up to date. For more detail on how to listen and respond to these webhooks, see [Liabilities webhooks](/docs/api/products/liabilities/#webhooks).

#### Testing Liabilities

Liabilities can be tested in Sandbox without any additional permissions.

Plaid also provides a [GitHub repo](https://github.com/plaid/sandbox-custom-users/) with test data for testing student loan accounts in Sandbox. For more information on configuring custom Sandbox data, see [Configuring the custom user account](/docs/sandbox/user-custom/#configuring-the-custom-user-account).

#### Liabilities pricing

Liabilities is billed on a [subscription model](/docs/account/billing/#subscription-fee). To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

#### Next steps

To get started building with Liabilities, see [Add Liabilities to your App](/docs/liabilities/add-to-app/).

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
