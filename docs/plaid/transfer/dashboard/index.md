---
title: "Transfer - Transfer Dashboard | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/dashboard/"
scraped_at: "2026-03-07T22:05:24+00:00"
---

# Transfer Dashboard

#### Learn about the features available in the Plaid Transfer dashboard.

#### Overview

You can complete most transfer operational activities in the Transfer Dashboard. Many of these actions are also available via the [API](/docs/api/products/transfer/).

A link to the Transfer Dashboard will appear in the Dashboard navigation after you have applied for access to [Transfer in the Production environment](https://dashboard.plaid.com/overview/production). You can also access the Transfer Dashboard before applying for Production access by navigating directly to [dashboard.plaid.com/transfer](https://dashboard.plaid.com/transfer).

The main overview page displays the following:

- **Aggregate metrics**: total transfers, return rates, current volumes
- **Account information**: limits (daily, monthly) and utilization to date. For ACH transfers, this also includes SEC codes you are approved for.
- **Transfer activity**: table displaying all transfer, refund and sweep activity

![Plaid Transfer Dashboard showing aggregate metrics, account info, and transfer activity with search, filter, and export options.](/assets/img/docs/transfer/dashboard_home.png)

Home page of Transfer dashboard

Features available through the Dashboard include:

- [Viewing and changing funding account information](/docs/transfer/flow-of-funds/#adding-or-changing-funding-accounts): You can self-serve linking a new funding account, deleting an account, or changing the default status of a funding account in the "Account details" page of the dashboard
- [Moving money between your funding account and your Plaid Ledger](/docs/transfer/flow-of-funds/#moving-money-via-the-dashboard)
- [Refunding](/docs/transfer/refunds/) or [canceling](/docs/transfer/creating-transfers/#canceling-transfers) transfers
- Viewing and searching transfer details including events, statuses, and associated refunds, and viewing the trace numbers and event history of a transfer
- Configuring [Signal rules](/docs/transfer/signal-rules/)

Creating and authorizing a transfer *cannot* be completed from the Dashboard, nor can you link an end user's bank account from the Dashboard. These operations can only be completed using the API.

![Transfer details page showing $5 debit, options to simulate events, refund, or cancel. State: Settled. Activity log included.](/assets/img/docs/transfer/transfer-details.png)

Transfer details page, including options to refund and cancel transfer. Because this transaction occurred in the Sandbox, options to manually change its state are shown.

#### Limits, reserves, and hold time

Your initial limits, reserve requirements, and debit hold time (if applicable) are set at onboarding. As your business grows over time, you can request changes to your debit limits, credit limits, and/or debit hold time through the dashboard.

![Plaid Transfer Dashboard. Enter projected debit and credit volumes for limit change request. Options include single, daily, monthly.](/assets/img/docs/transfer/request-limit-change.png)

Plaid Transfer Dashboard Limits Change Request

If your requested changes require additional reserves, the amount will be indicated to you. Once the additional reserves are sent to Plaid (if applicable), the approved changes will be automatically implemented.

You can view your current limits, reserves, and debit hold time in the “Account Details” section of the dashboard:

![Plaid Transfer Dashboard showing account details, funding accounts, approved limits, hold times, and options to deposit or withdraw.](/assets/img/docs/transfer/limits-and-hold-times.png)

Plaid Transfer Dashboard Limits and Hold Times

Wire instructions for funding the reserve account can also be found by clicking "Wire instructions" on the "Account details" section.

#### Ledger management

From the Dashboard, you can see more details of your Ledger, add or withdraw money, and configure settings such as automated cashouts. For details, see [Moving money via the Dashboard](/docs/transfer/flow-of-funds/#moving-money-via-the-dashboard).

#### Report extraction

##### Transfer Activity Report

In the Overview page of the Transfer section of the Plaid Dashboard, you can export a report with a list of transfers and sweeps as a .csv file by clicking "Export CSV."

![Popup for downloading transfers as CSV. Options: 'Yes, continue' and 'Cancel'. Background: transfer metrics and activity log.](/assets/img/docs/transfer/csv.png)

Plaid Transfer Dashboard CSV Export

Filters can be applied to the report by selecting **Filter**.

- **Date**: Only transfers and sweeps that were created during the specified period (have `created_at` timestamps within the date range) will be included
- **Transaction type**: Only the selected transaction types (transfers, sweeps, and/or refunds) will be included
- **Status**: Only transfers or sweeps that currently (as of the report creation time) have the selected status(es) (`pending`, `posted`, `settled`, `funds available`, `returned`, `failed`, and/or `cancelled`) will be included

**In this report, you will see the following columns:**
*Note:* All timestamps utilize the timezone (EST or UTC) applied to the report.

| Column | Description |
| --- | --- |
| ID | The ID of the transaction type that occurred. For example, if the Type is `transfer`, then this field is the transfer ID. |
| Status | The status of the transaction type as of the time of the report export. Note that `funds available` is a status that is only applicable to debit transfers or debit sweeps. |
| Type | Whether the transaction was a `debit` or `credit` payment instruction. Transfers and sweeps can be debits or credits. Refunds are credits only. |
| Amount | The amount of the transaction. Amounts are signed. Debits have positive signed amounts (as debit transfers and sweeps increase the ledger balance). Credits have negative signed amounts (this activity decreases the ledger balance). |
| ACH class | For activity that was executed via the ACH network, the ACH Standard Entry Class (SEC) code of the transaction. For more details, see [ACH SEC codes](https://plaid.com/docs/transfer/creating-transfers/#ach-sec-codes). |
| Network | The payment network used to fulfill the transaction. For debits, this will be either `same-day-ach` or `ach` (Standard ACH). For credits, this can be `rtp`, `wire`, `same-day-ach`, or `ach`. Note that if a Same Day ACH transfer was submitted after the cutoff, Transfer will downgrade the network to Standard ACH. Sweeps are not downgraded, so if a Same Day ACH cutoff is missed, the sweep will be retried the next morning for the first Same Day ACH cutoff of the day. |
| Account\_id | Plaid’s unique identifier for the account. |
| Account type | The type of bank account being debited or credited for a transfer. Either `checking` or `savings`. |
| User Legal Name | The legal name of the end user being debited or credited. Only populated for transfers and refunds. |
| User Email Address | The email address of the end user being debited or credited. Only populated for transfers and refunds. |
| User Routing Number | The routing number of the bank account being debited or credited. |
| Failure reason - ACH return code | If an ACH or Same Day ACH transfer has failed (returned), the return code (RXX). |
| Failure reason - Description | If an ACH or Same Day ACH transfer has failed (returned), a description for the return code. |
| Timestamps: Created / Pending / Posted / Settled / Funds Available / Returned / Failed / Cancelled | The timestamp of when the transfer, refund, or sweep progressed to this status. |
| Metadata | The contents of the key-value pair(s) included for the transaction (if applicable). |
| Refund Of | If the transaction is a refund, the ID of the debit transfer that the transaction is refunding |
| Refunded Amount | Refund amount issued. Refunds may be issued partially or completely. |
| Refund 1, 2, 3 | ID of the refund(s) issued. Multiple refunds can only be issued if the initial refund was not a complete refund of the debit transfer. |

*Note:* The following columns are not applicable for customers on Plaid Ledger: `swept`, `return_swept`, `swept_settled`, `signed_sweep_amount`, `reserve_swept`, `signed_reserve_sweep_amount`, `signed_return_sweep_amount`.

##### Ledger balance Report

Plaid provides reconciliation reports of all money moved into and out of your Plaid Ledger settlement balance. You can generate a custom report for any date range (up to 31 days) via the [Transfer account details](https://dashboard.plaid.com/transfer/account-details) Dashboard page.

In this report, you will see the following columns:

| Column | Description |
| --- | --- |
| Timestamp | The datetime, in the Eastern time zone, when the balance was impacted. |
| Originator Client ID / Client Name | Denotes the owner of the Plaid Ledger balance. For customers running a payments [platform](https://plaid.com/docs/transfer/platform-payments/), this will be your subcustomer’s client ID and business name. Otherwise, it’s your own client ID and business name. |
| Type | The type of activity impacting the balance: `transfer` indicates activity between the end consumer’s bank account (i.e accounts linked in your application’s product flows) and the Ledger balance. `sweep` indicates activity between your business checking accounts and the Ledger balance. `refund`s can occur on ACH debits to return funds to the consumer account. `distribute` indicates a movement of funds from one Plaid Ledger’s available balance to another Plaid Ledger’s available balance. `manual_adjustment` indicates this balance change was entered manually by Plaid. |
| ID | The ID of the activity that occurred. For example, if the Type is `transfer`, then this field is the transfer ID. |
| Description | A summary of the action that took place on this Type. `originated`: A request to originate the transfer, sweep, refund or distribute was received by Plaid. `returned`: An ACH return occurred. `canceled`: Payment execution was canceled via the API or dashboard. `converted_to_available`: Funds that were originally placed in a pending balance were moved to the available balance. `failed`: The attempt to move money failed. `dishonor`: a previous ACH return that came in was invalidated. When a dishonored return occurs, the transfer status will not change, but a row will be added to the report showing the funds from the dishonored return being added back to the Ledger balance.  `facilitator_fee`: A fee was charged by the platform parent to its subcustomer. `distribute`: Funds were moved from one Plaid Ledger’s `available` balance to another Plaid Ledger’s `available` balance. `manual_adjustment`: A balance adjustment was entered manually. `incentive`: If “incentive” is prepended to a value above, it is referencing a pay by bank incentive that occurred on the transfer (for example, `incentive originated`). |
| Impacted Balance Type | The Ledger balance type that was impacted: `available`, `pending`, or `pending_to_available`. The latter denotes that both balances were impacted during a move of funds from `pending` to `available`. |
| Amount | The signed amount of the balance impact. `pending_to_available` is always a positive amount, noting the amount removed from `pending` and allocated to `available`. |
| New Pending Balance | The resulting value of the `pending` balance after this activity was applied to the relevant balance(s). |
| New Available Balance | The resulting value of the `available` balance after this activity was applied to the relevant balance(s). |

Using these columns, you can complete comprehensive financial reconciliation. For example: to see all the activity between your business checking account and the Ledger balance, filter for all rows where the column "Type" is `sweep`. To see refunds issued, filter for all rows where the column "Type" is `refund`. To see exactly when a specific ACH debit transfer’s funds became available in the Ledger balance, search for the row of type `transfer`, specifying the relevant `ID`, and look for its `pending_to_available` balance impact type.

##### Email reconciliation reports

Plaid also provides optional monthly reconciliation reports via email, including aggregated details of funds movement activity for the month. To opt in to receiving these reports, contact your account manager.

#### Rules

See [Customizing Signal rules](/docs/transfer/signal-rules/) to learn how to configure the rules that will be run during authorization checks for debit transfers.

#### Tasks

The Tasks pane will show any onboarding tasks you have not yet completed, such as selecting a hold time or sending your initial reserve amounts.

If Plaid is requesting that you provide [Proof of Authorization](/docs/transfer/creating-transfers/#managing-proof-of-authorization) for any disputed transactions, this will also be shown on the Tasks pane.

#### Dashboard permissions

Transfer-specific Dashboard permissions can be set via the [team settings page](https://dashboard.plaid.com/settings/team/). Transfer-specific permissions are as follows:

- **Transfer Read & Write all** - Grants access to Transfer features.
  - **Transfer Read** - Grants access to view Transfer Dashboard
  - **Manage funding accounts** - Grants ability to manage Transfer funding accounts
  - **Manage Ledgers** - Grants ability to manage Ledgers, including automated balance management settings
  - **Manage transfer limits** - Grants ability to request changes to Transfer limits and hold times
  - **Initiate refunds** - Grants ability to initiate refunds
  - **Manage transfer Ledger funds** - Grants ability to deposit and withdraw funds from a Transfer Ledger

Note that these permissions only impact Dashboard access and do not apply to actions initiated via the API. A user who has the "Production Keys" permission or otherwise has access to Plaid Production API keys can still use the API to move funds in and out of the Ledger or initiate a refund, regardless of their Transfer Dashboard permissions.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
