---
title: "Signal - Reporting returns and decisions | Plaid Docs"
source_url: "https://plaid.com/docs/signal/reporting-returns/"
scraped_at: "2026-03-07T22:05:18+00:00"
---

# Reporting returns and decisions

#### Improve accuracy by reporting outcomes

It's critical that you report ACH activity (both decisions and returns) back to Plaid - reporting enables Plaid to help you tune your rule logic to optimize your integration.

For both decisions and returns, you can use API endpoints (recommended) or upload the transactions manually via the Signal Dashboard.

For Signal Transaction Scores customers, reporting decisions and returns is required in order to receive the most accurate results. For Balance-only customers, it is optional but recommended, as it's necessary in order to view transaction outcomes and tune rule logic in the Dashboard, and will "future proof" your integration if you later choose to augment Balance with Signal Transaction Scores.

#### Reporting decisions

Reporting a decision allows Plaid to know that the transaction was processed. Once it knows the transaction was processed, if Plaid doesn't hear about a corresponding return, it assumes the transaction was successful.

Plaid will automatically infer decisions for customers using rules with only `ACCEPT` or `REROUTE` results. There is no need to report decisions back to Plaid for these transactions unless the decision has changed (for example, your processor refused to process an `ACCEPT`ed transaction).

If your rules integration uses `REVIEW` results, then you'll need to call [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) once you've made the final determination for the transaction and it is convenient to do so.

If your integration is directly reading scores and attributes instead of using rules, then you'll need to call [`/signal/decision/report`](/docs/api/products/signal/#signaldecisionreport) for all transactions.

#### Reporting returns

To report an ACH return to Plaid, call [`/signal/return/report`](/docs/api/products/signal/#signalreturnreport) as soon as you're aware of the ACH return and it's convenient to do so.

#### Dashboard upload

If you would prefer to not report decisions and returns via API, you can upload them via the [data import](https://dashboard.plaid.com/signal/import) page in the Dashboard. It is recommended you upload reports on a weekly basis.

The CSV files for uploading decisions and returns must follow the schemas defined below. Ensure your data is properly formatted to avoid processing errors.

##### Decision File CSV Schema

| Column Name | Required? | Data Type | Description |
| --- | --- | --- | --- |
| `client_transaction_id` | Yes | String | Unique identifier for the transaction used in [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) |
| `initiated` | Yes | Boolean | Whether the transaction was initiated for processing or not. |
| `days_funds_on_hold` | No | Integer | Number of days the funds were held before being released, if applicable to your usecase. |
| `amount_instantly_available` | No | Integer | The amount made available instantly in cents, if relevant for your account funding usecase |

[Example decision file CSV](https://plaid.com/documents/signal-decision-import-example.csv)

##### Return File CSV Schema

| Column Name | Required? | Data Type | Description |
| --- | --- | --- | --- |
| `client_transaction_id` | Yes | String | Unique identifier for the transaction used in [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) |
| `return_code` | Yes | String | The ACH return code (e.g., R01, R02, R10) |
| `returned_at` | No | String | Timestamp of when the return occurred in ISO 8601 format (YYYY-MM-DDThh:mm:ssTZD) |

[Example return file CSV](https://plaid.com/documents/signal-return-import-example.csv)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
