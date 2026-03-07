---
title: "API - Consumer Report (by Plaid Check) | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/check/"
scraped_at: "2026-03-07T22:04:08+00:00"
---

# Plaid Check

#### API reference for Plaid Check endpoints and webhooks

| Endpoints |  |
| --- | --- |
| [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget) | Retrieve the base Consumer Report for your user |
| [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget) | Retrieve cash flow insights from your user's banks |
| [`/cra/check_report/network_insights/get`](/docs/api/products/check/#cracheck_reportnetwork_insightsget) | Retrieve connection insights from the Plaid network (beta) |
| [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget) | Retrieve cash flow insights from our partners |
| [`/cra/check_report/pdf/get`](/docs/api/products/check/#cracheck_reportpdfget) | Retrieve Consumer Reports in PDF format |
| [`/cra/check_report/cashflow_insights/get`](/docs/api/products/check/#cracheck_reportcashflow_insightsget) | Retrieve Cash Flow Insights report |
| [`/cra/check_report/lend_score/get`](/docs/api/products/check/#cracheck_reportlend_scoreget) | Retrieve a Plaid LendScore generated from your user's banking data |
| [`/cra/check_report/verification/get`](/docs/api/products/check/#cracheck_reportverificationget) | Retrieve Verification Reports (Home Lending Report, Employment Refresh) for your user |
| [`/cra/check_report/verification/pdf/get`](/docs/api/products/check/#cracheck_reportverificationpdfget) | Retrieve Verification Reports in PDF format |
| [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) | Generate a new Consumer Report for your user with the latest data |
| [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) | Get Cash Flow Updates (beta) |
| [`/cra/monitoring_insights/subscribe`](/docs/api/products/check/#cramonitoring_insightssubscribe) | Subscribe to Cash Flow Updates (beta) |
| [`/cra/monitoring_insights/unsubscribe`](/docs/api/products/check/#cramonitoring_insightsunsubscribe) | Unsubscribe from Cash Flow Updates (beta) |

| See also |  |
| --- | --- |
| [`/link/token/create`](/docs/api/link/#linktokencreate) | Create a token for initializing a Link session with Plaid Check |
| [`/user/create`](/docs/api/users/#usercreate) | Create a user ID and token for use with Plaid Check |
| [`/user/update`](/docs/api/users/#userupdate) | Update an existing user token to work with Plaid Check, or change user details |
| [`/user/remove`](/docs/api/users/#userremove) | Removes the user and their relevant data |
| [`/user/items/get`](/docs/api/users/#useritemsget) | Returns Items associated with a user along with their corresponding statuses |
| [`/sandbox/cra/cashflow_updates/update`](/docs/api/sandbox/#sandboxcracashflow_updatesupdate) | Manually trigger a cashflow insights update for a user (Sandbox only) |

| Webhooks |  |
| --- | --- |
| [`USER_CHECK_REPORT_READY`](/docs/api/products/check/#user_check_report_ready) | A Consumer Report is ready to be retrieved |
| [`USER_CHECK_REPORT_FAILED`](/docs/api/products/check/#user_check_report_failed) | Plaid Check failed to create a report |
| [`CHECK_REPORT_READY`](/docs/api/products/check/#check_report_ready) | A Consumer Report is ready to be retrieved (legacy) |
| [`CHECK_REPORT_FAILED`](/docs/api/products/check/#check_report_failed) | Plaid Check failed to create a report (legacy) |

| Cash Flow Updates (beta) webhooks |  |
| --- | --- |
| [`CASH_FLOW_INSIGHTS_UPDATED`](/docs/api/products/check/#insights_updated) | Insights have been refreshed |
| [`INSIGHTS_UPDATED`](/docs/api/products/check/#insights_updated) | Insights have been refreshed (legacy) |
| [`LARGE_DEPOSIT_DETECTED`](/docs/api/products/check/#large_deposit_detected) | A large deposit over $5000 has been detected (legacy) |
| [`LOW_BALANCE_DETECTED`](/docs/api/products/check/#low_balance_detected) | Current balance has crossed below $100 (legacy) |
| [`NEW_LOAN_PAYMENT_DETECTED`](/docs/api/products/check/#new_loan_payment_detected) | A new loan payment has been detected (legacy) |
| [`NSF_OVERDRAFT_DETECTED`](/docs/api/products/check/#nsf_overdraft_detected) | An overdraft transaction has been detected (legacy) |

=\*=\*=\*=

#### `/cra/check_report/base_report/get`

#### Retrieve a Base Report

This endpoint allows you to retrieve the Base Report for your user, allowing you to receive comprehensive bank account and cash flow data. You should call this endpoint after you've received a `CHECK_REPORT_READY` or a `USER_CHECK_REPORT_READY` webhook, either after the Link session for the user or after calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate). If the most recent consumer report for the user doesn't have sufficient data to generate the base report, or the consumer report has expired, you will receive an error indicating that you should create a new consumer report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

/cra/check\_report/base\_report/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-base_report-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-base_report-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-base_report-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`user_token`](/docs/api/products/check/#cra-check_report-base_report-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/cra/check\_report/base\_report/get

```
try {
  const response = await client.craCheckReportBaseReportGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
  });
} catch (error) {
  // handle error
}
```

/cra/check\_report/base\_report/get

**Response fields**

[`report`](/docs/api/products/check/#cra-check_report-base_report-get-response-report)

objectobject

An object representing a Base Report

[`report_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-report-id)

stringstring

A unique ID identifying an Base Report. Like all Plaid identifiers, this ID is case sensitive.

[`date_generated`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-date-generated)

stringstring

The date and time when the Base Report was created, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (e.g. "2018-04-12T03:32:11Z").  
  

Format: `date-time`

[`days_requested`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-days-requested)

numbernumber

The number of days of transaction history requested.

[`client_report_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-client-report-id)

nullablestringnullable, string

Client-generated identifier, which can be used by lenders to track loan applications.

[`items`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items)

[object][object]

Data returned by Plaid about each of the Items included in the Base Report.

[`institution_name`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-institution-name)

stringstring

The full financial institution name associated with the Item.

[`institution_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-institution-id)

stringstring

The id of the financial institution associated with the Item.

[`date_last_updated`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-date-last-updated)

stringstring

The date and time when this Item’s data was last retrieved from the financial institution, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`item_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`accounts`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts)

[object][object]

Data about each of the accounts open on the Item.

[`account_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`balances`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances)

objectobject

Information about an account's balances.

[`available`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-available)

nullablenumbernullable, number

The amount of funds available to be withdrawn from the account, as determined by the financial institution.  
For `credit`-type accounts, the `available` balance typically equals the `limit` less the `current` balance, less any pending outflows plus any pending inflows.  
For `depository`-type accounts, the `available` balance typically equals the `current` balance less any pending outflows plus any pending inflows. For `depository`-type accounts, the `available` balance does not include the overdraft limit.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the `available` balance is the total cash available to withdraw as presented by the institution.  
Note that not all institutions calculate the `available` balance. In the event that `available` balance is unavailable, Plaid will return an `available` balance value of `null`.  
Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by `/accounts/balance/get`.  
If `current` is `null` this field is guaranteed not to be `null`.  
  

Format: `double`

[`current`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-current)

nullablenumbernullable, number

The total amount of funds in or owed by the account.  
For `credit`-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  
For `loan`-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (`ins_116944`). For Sallie Mae student loans, the account's balance includes both principal and any outstanding interest.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the current balance is the total value of assets as presented by the institution.  
Note that balance information may be cached unless the value was returned by `/accounts/balance/get`; if the Item is enabled for Transactions, the balance will be at least as recent as the most recent Transaction update. If you require realtime balance information, use the `available` balance as provided by `/accounts/balance/get`.  
When returned by `/accounts/balance/get`, this field may be `null`. When this happens, `available` is guaranteed not to be `null`.  
  

Format: `double`

[`limit`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-limit)

nullablenumbernullable, number

For `credit`-type accounts, this represents the credit limit.  
For `depository`-type accounts, this represents the pre-arranged overdraft limit, which is common for current (checking) accounts in Europe.  
In North America, this field is typically only available for `credit`-type accounts.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the balance. Always null if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the balance. Always null if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`last_updated_datetime`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-last-updated-datetime)

nullablestringnullable, string

Timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the oldest acceptable balance when making a request to `/accounts/balance/get`.  
This field is only used and expected when the institution is `ins_128026` (Capital One) and the Item contains one or more accounts with a non-depository account type, in which case a value must be provided or an `INVALID_REQUEST` error with the code of `INVALID_FIELD` will be returned. For Capital One depository accounts as well as all other account types on all other institutions, this field is ignored. See [account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full list of account types.  
If the balance that is pulled is older than the given timestamp for Items with this field required, an `INVALID_REQUEST` error with the code of `LAST_UPDATED_DATETIME_OUT_OF_RANGE` will be returned with the most recent timestamp for the requested account contained in the response.  
  

Format: `date-time`

[`average_balance`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-average-balance)

nullablenumbernullable, number

The average historical balance for the entire report  
  

Format: `double`

[`average_monthly_balances`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-average-monthly-balances)

[object][object]

The average historical balance of each calendar month

[`start_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-average-monthly-balances-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).

[`end_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-average-monthly-balances-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).

[`average_balance`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-average-monthly-balances-average-balance)

objectobject

This contains an amount, denominated in the currency specified by either `iso_currency_code` or `unofficial_currency_code`

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-average-monthly-balances-average-balance-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-average-monthly-balances-average-balance-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-average-monthly-balances-average-balance-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`most_recent_thirty_day_average_balance`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-balances-most-recent-thirty-day-average-balance)

nullablenumbernullable, number

The average historical balance from the most recent 30 days  
  

Format: `double`

[`consumer_disputes`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-consumer-disputes)

[object][object]

The information about previously submitted valid dispute statements by the consumer

[`consumer_dispute_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-consumer-disputes-consumer-dispute-id)

deprecatedstringdeprecated, string

(Deprecated) A unique identifier (UUID) of the consumer dispute that can be used for troubleshooting

[`dispute_field_create_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-consumer-disputes-dispute-field-create-date)

stringstring

Date of the disputed field (e.g. transaction date), in an ISO 8601 format (YYYY-MM-DD)  
  

Format: `date`

[`category`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-consumer-disputes-category)

stringstring

Type of data being disputed by the consumer  
  

Possible values: `TRANSACTION`, `BALANCE`, `IDENTITY`, `OTHER`

[`statement`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-consumer-disputes-statement)

stringstring

Text content of dispute

[`mask`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.

[`metadata`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-metadata)

objectobject

Metadata about the extracted account.

[`start_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-metadata-start-date)

nullablestringnullable, string

The beginning of the range of the financial institution provided data for the account, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-metadata-end-date)

nullablestringnullable, string

The end of the range of the financial institution provided data for the account, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`name`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself

[`official_name`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-official-name)

nullablestringnullable, string

The official name of the account as given by the financial institution

[`type`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`days_available`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-days-available)

numbernumber

The duration of transaction history available within this report for this Item, typically defined as the time since the date of the earliest transaction in that account.

[`transactions`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions)

[object][object]

Transaction history associated with the account. Transaction history returned by endpoints such as `/transactions/get` or `/investments/transactions/get` will be returned in the top-level `transactions` field instead. Some transactions may have their details masked in accordance to the FCRA. These will appear with a `credit_category` of `MASKED_TRANSACTION_CATEGORY`.

[`account_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-account-id)

stringstring

The ID of the account in which this transaction occurred.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-amount)

numbernumber

The settled value of the transaction, denominated in the transaction's currency, as stated in `iso_currency_code` or `unofficial_currency_code`. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the transaction. Always `null` if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the transaction. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`original_description`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-original-description)

nullablestringnullable, string

The string returned by the financial institution to describe the transaction.

[`credit_category`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-credit-category)

nullableobjectnullable, object

Information describing the intent of the transaction. Most relevant for credit use cases, but not limited to such use cases.  
See the [`taxonomy csv file`](https://plaid.com/documents/credit-category-taxonomy.csv) for a full list of credit categories.

[`primary`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-credit-category-primary)

stringstring

A high level category that communicates the broad category of the transaction.

[`detailed`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-credit-category-detailed)

stringstring

A granular category conveying the transaction's intent. This field can also be used as a unique identifier for the category.

[`check_number`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-check-number)

nullablestringnullable, string

The check number of the transaction. This field is only populated for check transactions.

[`date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-date)

stringstring

For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DD` ).  
  

Format: `date`

[`date_transacted`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-date-transacted)

nullablestringnullable, string

The date on which the transaction took place, in IS0 8601 format.

[`location`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location)

objectobject

A representation of where a transaction took place. Location data is provided only for transactions at physical locations, not for online transactions. Location data availability depends primarily on the merchant and is most likely to be populated for transactions at large retail chains; small, local businesses are less likely to have location data available.

[`address`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location-address)

nullablestringnullable, string

The street address where the transaction occurred.

[`city`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location-city)

nullablestringnullable, string

The city where the transaction occurred.

[`region`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location-region)

nullablestringnullable, string

The region or state where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called `state`.

[`postal_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location-postal-code)

nullablestringnullable, string

The postal code where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code where the transaction occurred.

[`lat`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location-lat)

nullablenumbernullable, number

The latitude where the transaction occurred.  
  

Format: `double`

[`lon`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location-lon)

nullablenumbernullable, number

The longitude where the transaction occurred.  
  

Format: `double`

[`store_number`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-location-store-number)

nullablestringnullable, string

The merchant defined store number where the transaction occurred.

[`merchant_name`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-merchant-name)

nullablestringnullable, string

The merchant name, as enriched by Plaid from the `name` field. This is typically a more human-readable version of the merchant counterparty in the transaction. For some bank transactions (such as checks or account transfers) where there is no meaningful merchant name, this value will be `null`.

[`pending`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-pending)

booleanboolean

When `true`, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled.

[`account_owner`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-account-owner)

nullablestringnullable, string

The name of the account owner. This field is not typically populated and only relevant when dealing with sub-accounts.

[`transaction_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-transactions-transaction-id)

stringstring

The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.

[`owners`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners)

[object][object]

Data returned by the financial institution about the account owner or owners. For business accounts, the name reported may be either the name of the individual or the name of the business, depending on the institution. Multiple owners on a single account will be represented in the same `owner` object, not in multiple owner objects within the array. This array can also be empty if no owners are found.

[`names`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-names)

[string][string]

A list of names associated with the account by the financial institution. In the case of a joint account, Plaid will make a best effort to report the names of all account holders.  
If an Item contains multiple accounts with different owner names, some institutions will report all names associated with the Item in each account's `names` array.

[`phone_numbers`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-phone-numbers)

[object][object]

A list of phone numbers associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-phone-numbers-data)

stringstring

The phone number.

[`primary`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-phone-numbers-primary)

booleanboolean

When `true`, identifies the phone number as the primary number on an account.

[`type`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-phone-numbers-type)

stringstring

The type of phone number.  
  

Possible values: `home`, `work`, `office`, `mobile`, `mobile1`, `other`

[`emails`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-emails)

[object][object]

A list of email addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-emails-data)

stringstring

The email address.

[`primary`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-emails-primary)

booleanboolean

When `true`, identifies the email address as the primary email on an account.

[`type`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-emails-type)

stringstring

The type of email account as described by the financial institution.  
  

Possible values: `primary`, `secondary`, `other`

[`addresses`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-addresses)

[object][object]

Data about the various addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-addresses-data)

objectobject

Data about the components comprising an address.

[`city`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-addresses-data-city)

nullablestringnullable, string

The full city name

[`region`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-addresses-data-region)

nullablestringnullable, string

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-addresses-data-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-addresses-data-postal-code)

nullablestringnullable, string

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-addresses-data-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code

[`primary`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-owners-addresses-primary)

booleanboolean

When `true`, identifies the address as the primary address on an account.

[`ownership_type`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-ownership-type)

nullablestringnullable, string

How an asset is owned.  
`association`: Ownership by a corporation, partnership, or unincorporated association, including for-profit and not-for-profit organizations.
`individual`: Ownership by an individual.
`joint`: Joint ownership by multiple parties.
`trust`: Ownership by a revocable or irrevocable trust.  
  

Possible values: `null`, `individual`, `joint`, `association`, `trust`

[`account_insights`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights)

deprecatedobjectdeprecated, object

Calculated insights derived from transaction-level data. This field has been deprecated in favor of [Base Report attributes aggregated across accounts](https://plaid.com/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes) and will be removed in a future release.

[`oldest_transaction_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-oldest-transaction-date)

nullablestringnullable, string

Date of the earliest transaction for the account.  
  

Format: `date`

[`most_recent_transaction_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-most-recent-transaction-date)

nullablestringnullable, string

Date of the most recent transaction for the account.  
  

Format: `date`

[`days_available`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-days-available)

integerinteger

Number of days days available for the account.

[`average_days_between_transactions`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-days-between-transactions)

numbernumber

Average number of days between sequential transactions

[`longest_gaps_between_transactions`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-longest-gaps-between-transactions)

[object][object]

Longest gap between sequential transactions in a time period. This array can include multiple time periods.

[`start_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-longest-gaps-between-transactions-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-longest-gaps-between-transactions-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`days`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-longest-gaps-between-transactions-days)

nullableintegernullable, integer

Largest number of days between sequential transactions for this time period.

[`number_of_inflows`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-inflows)

[object][object]

The number of debits into the account. This array will be empty for non-depository accounts.

[`start_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-inflows-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-inflows-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`count`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-inflows-count)

integerinteger

The number of credits or debits out of the account for this time period.

[`average_inflow_amounts`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-inflow-amounts)

[object][object]

Average amount of debit transactions into the account in a time period. This array will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`start_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-inflow-amounts-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-inflow-amounts-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`total_amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-inflow-amounts-total-amount)

objectobject

This contains an amount, denominated in the currency specified by either `iso_currency_code` or `unofficial_currency_code`

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-inflow-amounts-total-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-inflow-amounts-total-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-inflow-amounts-total-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`number_of_outflows`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-outflows)

[object][object]

The number of outflows from the account. This array will be empty for non-depository accounts.

[`start_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-outflows-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-outflows-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`count`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-outflows-count)

integerinteger

The number of credits or debits out of the account for this time period.

[`average_outflow_amounts`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-outflow-amounts)

[object][object]

Average amount of transactions out of the account in a time period. This array will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`start_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-outflow-amounts-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-outflow-amounts-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`total_amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-outflow-amounts-total-amount)

objectobject

This contains an amount, denominated in the currency specified by either `iso_currency_code` or `unofficial_currency_code`

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-outflow-amounts-total-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-outflow-amounts-total-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-average-outflow-amounts-total-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`number_of_days_no_transactions`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-account-insights-number-of-days-no-transactions)

integerinteger

Number of days with no transactions

[`attributes`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes)

objectobject

Calculated attributes derived from transaction-level data.

[`is_primary_account`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-is-primary-account)

nullablebooleannullable, boolean

Prediction indicator of whether the account is a primary account. Only one account per account type across the items connected will have a value of true.

[`primary_account_score`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-primary-account-score)

nullablenumbernullable, number

Value ranging from 0-1. The higher the score, the more confident we are of the account being the primary account.

[`nsf_overdraft_transactions_count`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-nsf-overdraft-transactions-count)

integerinteger

The number of net NSF fee transactions for a given account within the report time range (not counting any fees that were reversed within the time range).

[`nsf_overdraft_transactions_count_30d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-nsf-overdraft-transactions-count-30d)

integerinteger

The number of net NSF fee transactions within the last 30 days for a given account (not counting any fees that were reversed within the time range).

[`nsf_overdraft_transactions_count_60d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-nsf-overdraft-transactions-count-60d)

integerinteger

The number of net NSF fee transactions within the last 60 days for a given account (not counting any fees that were reversed within the time range).

[`nsf_overdraft_transactions_count_90d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-nsf-overdraft-transactions-count-90d)

integerinteger

The number of net NSF fee transactions within the last 90 days for a given account (not counting any fees that were reversed within the time range).

[`total_inflow_amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount)

nullableobjectnullable, object

Total amount of debit transactions into the account in the time period of the report. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_30d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-30d)

nullableobjectnullable, object

Total amount of debit transactions into the account in the last 30 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-30d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-30d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-30d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_60d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-60d)

nullableobjectnullable, object

Total amount of debit transactions into the account in the last 60 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-60d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-60d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-60d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_90d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-90d)

nullableobjectnullable, object

Total amount of debit transactions into the account in the last 90 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-90d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-90d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-inflow-amount-90d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount)

nullableobjectnullable, object

Total amount of credit transactions into the account in the time period of the report. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_30d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-30d)

nullableobjectnullable, object

Total amount of credit transactions into the account in the last 30 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-30d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-30d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-30d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_60d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-60d)

nullableobjectnullable, object

Total amount of credit transactions into the account in the last 60 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-60d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-60d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-60d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_90d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-90d)

nullableobjectnullable, object

Total amount of credit transactions into the account in the last 90 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-90d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-90d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-items-accounts-attributes-total-outflow-amount-90d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`attributes`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes)

objectobject

Calculated attributes derived from transaction-level data, aggregated across accounts.

[`nsf_overdraft_transactions_count`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-nsf-overdraft-transactions-count)

integerinteger

The number of net NSF fee transactions in the time range for the report (not counting any fees that were reversed within that time range).

[`nsf_overdraft_transactions_count_30d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-nsf-overdraft-transactions-count-30d)

integerinteger

The number of net NSF fee transactions in the last 30 days in the report (not counting any fees that were reversed within that time range).

[`nsf_overdraft_transactions_count_60d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-nsf-overdraft-transactions-count-60d)

integerinteger

The number of net NSF fee transactions in the last 60 days in the report (not counting any fees that were reversed within that time range).

[`nsf_overdraft_transactions_count_90d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-nsf-overdraft-transactions-count-90d)

integerinteger

The number of net NSF fee transactions in the last 90 days in the report (not counting any fees that were reversed within that time range).

[`total_inflow_amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount)

nullableobjectnullable, object

Total amount of debit transactions into the report's accounts in the time period of the report. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_30d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-30d)

nullableobjectnullable, object

Total amount of debit transactions into the report's accounts in the last 30 days. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-30d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-30d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-30d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_60d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-60d)

nullableobjectnullable, object

Total amount of debit transactions into the report's accounts in the last 60 days. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-60d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-60d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-60d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_90d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-90d)

nullableobjectnullable, object

Total amount of debit transactions into the report's account in the last 90 days. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-90d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-90d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-inflow-amount-90d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount)

nullableobjectnullable, object

Total amount of credit transactions into the report's accounts in the time period of the report. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_30d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-30d)

nullableobjectnullable, object

Total amount of credit transactions into the report's accounts in the last 30 days. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-30d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-30d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-30d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_60d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-60d)

nullableobjectnullable, object

Total amount of credit transactions into the report's accounts in the last 60 days. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-60d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-60d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-60d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_90d`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-90d)

nullableobjectnullable, object

Total amount of credit transactions into the report's accounts in the last 90 days. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-90d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-90d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes-total-outflow-amount-90d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`warnings`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings)

[object][object]

This array contains any information about errors or alerts related to the Base Report that did not block generation of the report.

[`warning_type`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-warning-type)

stringstring

The warning type, which will always be `BASE_REPORT_WARNING`

[`warning_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Account-owner information is not available.
`TRANSACTIONS_UNAVAILABLE`: Transactions information associated with Credit and Depository accounts are unavailable.
`USER_FRAUD_ALERT`: The User has placed a fraud alert on their Plaid Check consumer report due to suspected fraud. Note: when a fraud alert is in place, the recipient of the consumer report has an obligation to verify the consumer’s identity.  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `USER_FRAUD_ALERT`

[`cause`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause)

nullableobjectnullable, object

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`item_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-warnings-cause-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`request_id`](/docs/api/products/check/#cra-check_report-base_report-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "report": {
    "date_generated": "2024-07-16T01:52:42.912331716Z",
    "days_requested": 365,
    "attributes": {
      "total_inflow_amount": {
        "amount": -2500,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      },
      "total_inflow_amount_30d": {
        "amount": -1000,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      },
      "total_inflow_amount_60d": {
        "amount": -2500,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      },
      "total_inflow_amount_90d": {
        "amount": -2500,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      },
      "total_outflow_amount": {
        "amount": 2500,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      },
      "total_outflow_amount_30d": {
        "amount": 1000,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      },
      "total_outflow_amount_60d": {
        "amount": 2500,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      },
      "total_outflow_amount_90d": {
        "amount": 2500,
        "iso_currency_code": "USD",
        "unofficial_currency_code": null
      }
    },
    "items": [
      {
        "accounts": [
          {
            "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
            "account_insights": {
              "average_days_between_transactions": 0.15,
              "average_inflow_amount": [
                {
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01",
                  "total_amount": {
                    "amount": 1077.93,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                }
              ],
              "average_inflow_amounts": [
                {
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01",
                  "total_amount": {
                    "amount": 1077.93,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                },
                {
                  "end_date": "2024-08-31",
                  "start_date": "2024-08-01",
                  "total_amount": {
                    "amount": 1076.93,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                }
              ],
              "average_outflow_amount": [
                {
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01",
                  "total_amount": {
                    "amount": 34.95,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                }
              ],
              "average_outflow_amounts": [
                {
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01",
                  "total_amount": {
                    "amount": 34.95,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                },
                {
                  "end_date": "2024-08-31",
                  "start_date": "2024-08-01",
                  "total_amount": {
                    "amount": 0,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                }
              ],
              "days_available": 365,
              "longest_gap_between_transactions": [
                {
                  "days": 1,
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01"
                }
              ],
              "longest_gaps_between_transactions": [
                {
                  "days": 1,
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01"
                },
                {
                  "days": 2,
                  "end_date": "2024-08-31",
                  "start_date": "2024-08-01"
                }
              ],
              "most_recent_transaction_date": "2024-07-16",
              "number_of_days_no_transactions": 0,
              "number_of_inflows": [
                {
                  "count": 1,
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01"
                }
              ],
              "number_of_outflows": [
                {
                  "count": 27,
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01"
                }
              ],
              "oldest_transaction_date": "2024-07-12"
            },
            "balances": {
              "available": 5000,
              "average_balance": 4956.12,
              "average_monthly_balances": [
                {
                  "average_balance": {
                    "amount": 4956.12,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  },
                  "end_date": "2024-07-31",
                  "start_date": "2024-07-01"
                }
              ],
              "current": 5000,
              "iso_currency_code": "USD",
              "limit": null,
              "most_recent_thirty_day_average_balance": 4956.125,
              "unofficial_currency_code": null
            },
            "consumer_disputes": [],
            "days_available": 365,
            "mask": "1208",
            "metadata": {
              "start_date": "2024-01-01",
              "end_date": "2024-07-16"
            },
            "name": "Checking",
            "official_name": "Plaid checking",
            "owners": [
              {
                "addresses": [
                  {
                    "data": {
                      "city": "Malakoff",
                      "country": "US",
                      "postal_code": "14236",
                      "region": "NY",
                      "street": "2992 Cameron Road"
                    },
                    "primary": true
                  },
                  {
                    "data": {
                      "city": "San Matias",
                      "country": "US",
                      "postal_code": "93405-2255",
                      "region": "CA",
                      "street": "2493 Leisure Lane"
                    },
                    "primary": false
                  }
                ],
                "emails": [
                  {
                    "data": "accountholder0@example.com",
                    "primary": true,
                    "type": "primary"
                  },
                  {
                    "data": "accountholder1@example.com",
                    "primary": false,
                    "type": "secondary"
                  },
                  {
                    "data": "extraordinarily.long.email.username.123456@reallylonghostname.com",
                    "primary": false,
                    "type": "other"
                  }
                ],
                "names": [
                  "Alberta Bobbeth Charleson"
                ],
                "phone_numbers": [
                  {
                    "data": "+1 111-555-3333",
                    "primary": false,
                    "type": "home"
                  },
                  {
                    "data": "+1 111-555-4444",
                    "primary": false,
                    "type": "work"
                  },
                  {
                    "data": "+1 111-555-5555",
                    "primary": false,
                    "type": "mobile"
                  }
                ]
              }
            ],
            "ownership_type": null,
            "subtype": "checking",
            "transactions": [
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 37.07,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_ONLINE_MARKETPLACES",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-12",
                "date_posted": "2024-07-12T00:00:00Z",
                "date_transacted": "2024-07-12",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Amazon",
                "original_description": "AMZN Mktp US*11111111 Amzn.com/bill WA AM",
                "pending": false,
                "transaction_id": "XA7ZLy8rXzt7D3j9B6LMIgv5VxyQkAhbKjzmp",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 51.61,
                "check_number": null,
                "credit_category": {
                  "detailed": "DINING_DINING",
                  "primary": "DINING"
                },
                "date": "2024-07-12",
                "date_posted": "2024-07-12T00:00:00Z",
                "date_transacted": "2024-07-12",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Domino's",
                "original_description": "DOMINO's XXXX 111-222-3333",
                "pending": false,
                "transaction_id": "VEPeMbWqRluPVZLQX4MDUkKRw41Ljzf9gyLBW",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 7.55,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_FURNITURE_AND_HARDWARE",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-12",
                "date_posted": "2024-07-12T00:00:00Z",
                "date_transacted": "2024-07-12",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Chicago",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "IKEA",
                "original_description": "IKEA CHICAGO",
                "pending": false,
                "transaction_id": "6GQZARgvroCAE1eW5wpQT7w3oB6nvzi8DKMBa",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 12.87,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_SPORTING_GOODS",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-12",
                "date_posted": "2024-07-12T00:00:00Z",
                "date_transacted": "2024-07-12",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Redlands",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": "CA",
                  "state": "CA",
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Nike",
                "original_description": "NIKE REDLANDS CA",
                "pending": false,
                "transaction_id": "DkbmlP8BZxibzADqNplKTeL8aZJVQ1c3WR95z",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 44.21,
                "check_number": null,
                "credit_category": {
                  "detailed": "DINING_DINING",
                  "primary": "DINING"
                },
                "date": "2024-07-12",
                "date_posted": "2024-07-12T00:00:00Z",
                "date_transacted": "2024-07-12",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": null,
                "original_description": "POKE BROS * POKE BRO IL",
                "pending": false,
                "transaction_id": "RpdN7W8GmRSdjZB9Jm7ATj4M86vdnktapkrgL",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 36.82,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_DISCOUNT_STORES",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-13",
                "date_posted": "2024-07-13T00:00:00Z",
                "date_transacted": "2024-07-13",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Family Dollar",
                "original_description": "FAMILY DOLLAR",
                "pending": false,
                "transaction_id": "5AeQWvo5KLtAD9wNL68PTdAgPE7VNWf5Kye1G",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 13.27,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-13",
                "date_posted": "2024-07-13T00:00:00Z",
                "date_transacted": "2024-07-13",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Instacart",
                "original_description": "INSTACART HTTPSINSTACAR CA",
                "pending": false,
                "transaction_id": "Jjlr3MEVg1HlKbdkZj39ij5a7eg9MqtB6MWDo",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 36.03,
                "check_number": null,
                "credit_category": {
                  "detailed": "DINING_DINING",
                  "primary": "DINING"
                },
                "date": "2024-07-13",
                "date_posted": "2024-07-13T00:00:00Z",
                "date_transacted": "2024-07-13",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": null,
                "original_description": "POKE BROS * POKE BRO IL",
                "pending": false,
                "transaction_id": "kN9KV7yAZJUMPn93KDXqsG9MrpjlyLUL6Dgl8",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 54.74,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-13",
                "date_posted": "2024-07-13T00:00:00Z",
                "date_transacted": "2024-07-13",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Whittier",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": "CA",
                  "state": "CA",
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Smart & Final",
                "original_description": "POS SMART AND FINAL 111 WHITTIER CA",
                "pending": false,
                "transaction_id": "lPvrweZAMqHDar43vwWKs547kLZVEzfpogGVJ",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 37.5,
                "check_number": null,
                "credit_category": {
                  "detailed": "DINING_DINING",
                  "primary": "DINING"
                },
                "date": "2024-07-13",
                "date_posted": "2024-07-13T00:00:00Z",
                "date_transacted": "2024-07-13",
                "iso_currency_code": "USD",
                "location": {
                  "address": "1627 N 24th St",
                  "city": "Phoenix",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": "85008",
                  "region": "AZ",
                  "state": "AZ",
                  "store_number": null,
                  "zip": "85008"
                },
                "merchant_name": "Taqueria El Guerrerense",
                "original_description": "TAQUERIA EL GUERRERO PHOENIX AZ",
                "pending": false,
                "transaction_id": "wka74WKqngiyJ3pj7dl5SbpLGQBZqyCPZRDbP",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 41.42,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_ONLINE_MARKETPLACES",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-14",
                "date_posted": "2024-07-14T00:00:00Z",
                "date_transacted": "2024-07-14",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Amazon",
                "original_description": "AMZN Mktp US*11111111 Amzn.com/bill WA AM",
                "pending": false,
                "transaction_id": "BBGnV4RkerHjn8WVavGyiJbQ95VNDaC4M56bJ",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": -1077.93,
                "check_number": null,
                "credit_category": {
                  "detailed": "INCOME_OTHER",
                  "primary": "INCOME"
                },
                "date": "2024-07-14",
                "date_posted": "2024-07-14T00:00:00Z",
                "date_transacted": "2024-07-14",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Lyft",
                "original_description": "LYFT TRANSFER",
                "pending": false,
                "transaction_id": "3Ej78yKJlQu1Abw7xzo4U4JR6pmwzntZlbKDK",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 47.17,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-14",
                "date_posted": "2024-07-14T00:00:00Z",
                "date_transacted": "2024-07-14",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Whittier",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": "CA",
                  "state": "CA",
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Smart & Final",
                "original_description": "POS SMART AND FINAL 111 WHITTIER CA",
                "pending": false,
                "transaction_id": "rMzaBpJw8jSZRJQBabKdteQBwd5EaWc7J9qem",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 12.37,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-14",
                "date_posted": "2024-07-14T00:00:00Z",
                "date_transacted": "2024-07-14",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Whittier",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": "CA",
                  "state": "CA",
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Smart & Final",
                "original_description": "POS SMART AND FINAL 111 WHITTIER CA",
                "pending": false,
                "transaction_id": "zWPZjkmzynTyel89ZjExS59DV6WAaZflNBJ56",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 44.18,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-14",
                "date_posted": "2024-07-14T00:00:00Z",
                "date_transacted": "2024-07-14",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Portland",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": "OR",
                  "state": "OR",
                  "store_number": "1111",
                  "zip": null
                },
                "merchant_name": "Safeway",
                "original_description": "SAFEWAY #1111 PORTLAND OR            111111",
                "pending": false,
                "transaction_id": "K7qzx1nP8ptqgwaRMbxyI86XrqADMluRpkWx5",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 45.37,
                "check_number": null,
                "credit_category": {
                  "detailed": "DINING_DINING",
                  "primary": "DINING"
                },
                "date": "2024-07-14",
                "date_posted": "2024-07-14T00:00:00Z",
                "date_transacted": "2024-07-14",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Uber Eats",
                "original_description": "UBER EATS",
                "pending": false,
                "transaction_id": "qZrdzLRAgNHo5peMdD9xIzELl3a1NvcgrPAzL",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 15.22,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_ONLINE_MARKETPLACES",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-15",
                "date_posted": "2024-07-15T00:00:00Z",
                "date_transacted": "2024-07-15",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Amazon",
                "original_description": "AMZN Mktp US*11111111 Amzn.com/bill WA AM",
                "pending": false,
                "transaction_id": "NZzx4oRPkAHzyRekpG4PTZkWnBPqEyiy6pB1M",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 26.33,
                "check_number": null,
                "credit_category": {
                  "detailed": "DINING_DINING",
                  "primary": "DINING"
                },
                "date": "2024-07-15",
                "date_posted": "2024-07-15T00:00:00Z",
                "date_transacted": "2024-07-15",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Domino's",
                "original_description": "DOMINO's XXXX 111-222-3333",
                "pending": false,
                "transaction_id": "x84eNArKbESz8Woden6LT3nvyogeJXc64Pp35",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 39.8,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_DISCOUNT_STORES",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-15",
                "date_posted": "2024-07-15T00:00:00Z",
                "date_transacted": "2024-07-15",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Family Dollar",
                "original_description": "FAMILY DOLLAR",
                "pending": false,
                "transaction_id": "dzWnyxwZ4GHlZPGgrNyxiMG7qd5jDgCJEz5jL",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 45.06,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-15",
                "date_posted": "2024-07-15T00:00:00Z",
                "date_transacted": "2024-07-15",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Instacart",
                "original_description": "INSTACART HTTPSINSTACAR CA",
                "pending": false,
                "transaction_id": "4W7eE9rZqMToDArbPeLNIREoKpdgBMcJbVNQD",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 34.91,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-15",
                "date_posted": "2024-07-15T00:00:00Z",
                "date_transacted": "2024-07-15",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Whittier",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": "CA",
                  "state": "CA",
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Smart & Final",
                "original_description": "POS SMART AND FINAL 111 WHITTIER CA",
                "pending": false,
                "transaction_id": "j4yqDjb7QwS7woGzqrgDIEG1NaQVZwf6Wmz3D",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 49.78,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-15",
                "date_posted": "2024-07-15T00:00:00Z",
                "date_transacted": "2024-07-15",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Portland",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": "OR",
                  "state": "OR",
                  "store_number": "1111",
                  "zip": null
                },
                "merchant_name": "Safeway",
                "original_description": "SAFEWAY #1111 PORTLAND OR            111111",
                "pending": false,
                "transaction_id": "aqgWnze7xoHd6DQwLPnzT5dgPKjB1NfZ5JlBy",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 54.24,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-15",
                "date_posted": "2024-07-15T00:00:00Z",
                "date_transacted": "2024-07-15",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": "Portland",
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": "OR",
                  "state": "OR",
                  "store_number": "1111",
                  "zip": null
                },
                "merchant_name": "Safeway",
                "original_description": "SAFEWAY #1111 PORTLAND OR            111111",
                "pending": false,
                "transaction_id": "P13aP8b7nmS3WQoxg1PMsdvMK679RNfo65B4G",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 41.79,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_ONLINE_MARKETPLACES",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-16",
                "date_posted": "2024-07-16T00:00:00Z",
                "date_transacted": "2024-07-16",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Amazon",
                "original_description": "AMZN Mktp US*11111111 Amzn.com/bill WA AM",
                "pending": false,
                "transaction_id": "7nZMG6pXz8SADylMqzx7TraE4qjJm7udJyAGm",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 33.86,
                "check_number": null,
                "credit_category": {
                  "detailed": "FOOD_RETAIL_GROCERIES",
                  "primary": "FOOD_RETAIL"
                },
                "date": "2024-07-16",
                "date_posted": "2024-07-16T00:00:00Z",
                "date_transacted": "2024-07-16",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "Instacart",
                "original_description": "INSTACART HTTPSINSTACAR CA",
                "pending": false,
                "transaction_id": "MQr3ap7PWEIrQG7bLdaNsxyBV7g1KqCL6pwoy",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 27.08,
                "check_number": null,
                "credit_category": {
                  "detailed": "DINING_DINING",
                  "primary": "DINING"
                },
                "date": "2024-07-16",
                "date_posted": "2024-07-16T00:00:00Z",
                "date_transacted": "2024-07-16",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": null,
                "original_description": "POKE BROS * POKE BRO IL",
                "pending": false,
                "transaction_id": "eBAk9dvwNbHPZpr8W69dU3rekJz47Kcr9BRwl",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 25.94,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_FURNITURE_AND_HARDWARE",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-16",
                "date_posted": "2024-07-16T00:00:00Z",
                "date_transacted": "2024-07-16",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": "The Home Depot",
                "original_description": "THE HOME DEPOT",
                "pending": false,
                "transaction_id": "QLx4jEJZb9SxRm7aWbjAio3LrgZ5vPswm64dE",
                "unofficial_currency_code": null
              },
              {
                "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
                "account_owner": null,
                "amount": 27.57,
                "check_number": null,
                "credit_category": {
                  "detailed": "GENERAL_MERCHANDISE_OTHER_GENERAL_MERCHANDISE",
                  "primary": "GENERAL_MERCHANDISE"
                },
                "date": "2024-07-16",
                "date_posted": "2024-07-16T00:00:00Z",
                "date_transacted": "2024-07-16",
                "iso_currency_code": "USD",
                "location": {
                  "address": null,
                  "city": null,
                  "country": null,
                  "lat": null,
                  "lon": null,
                  "postal_code": null,
                  "region": null,
                  "state": null,
                  "store_number": null,
                  "zip": null
                },
                "merchant_name": null,
                "original_description": "The Press Club",
                "pending": false,
                "transaction_id": "ZnQ1ovqBldSQ6GzRbroAHLdQP68BrKceqmAjX",
                "unofficial_currency_code": null
              }
            ],
            "type": "depository"
          }
        ],
        "date_last_updated": "2024-07-16T01:52:42.912331716Z",
        "institution_id": "ins_109512",
        "institution_name": "Houndstooth Bank",
        "item_id": "NZzx4oRPkAHzyRekpG4PTZkDNkQW93tWnyGeA"
      }
    ],
    "report_id": "f3bb434f-1c9b-4ef2-b76c-3d1fd08156ec"
  },
  "warnings": [],
  "request_id": "FibfL8t3s71KJnj"
}
```

=\*=\*=\*=

#### `/cra/check_report/income_insights/get`

#### Retrieve cash flow information from your user's banks

This endpoint allows you to retrieve the Income Insights report for your user. You should call this endpoint after you've received a `CHECK_REPORT_READY` or a `USER_CHECK_REPORT_READY` webhook, either after the Link session for the user or after calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate). If the most recent consumer report for the user doesn’t have sufficient data to generate the base report, or the consumer report has expired, you will receive an error indicating that you should create a new consumer report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

/cra/check\_report/income\_insights/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-income_insights-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-income_insights-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/check/#cra-check_report-income_insights-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`user_id`](/docs/api/products/check/#cra-check_report-income_insights-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

/cra/check\_report/income\_insights/get

```
try {
  const response = await client.craCheckReportIncomeInsightsGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
  });
} catch (error) {
  // handle error
}
```

/cra/check\_report/income\_insights/get

**Response fields**

[`report`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report)

objectobject

The Check Income Insights Report for an end user.

[`report_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-report-id)

stringstring

The unique identifier associated with the Check Income Insights Report.

[`generated_time`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-generated-time)

stringstring

The time when the Check Income Insights Report was generated.  
  

Format: `date-time`

[`days_requested`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-days-requested)

integerinteger

The number of days requested by the customer for the Check Income Insights Report.

[`client_report_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-client-report-id)

nullablestringnullable, string

Client-generated identifier, which can be used by lenders to track loan applications.

[`items`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items)

[object][object]

The list of Items in the report along with the associated metadata about the Item.

[`item_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`bank_income_accounts`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts)

[object][object]

The Item's accounts that have bank income data.

[`account_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`mask`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number.
Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.

[`metadata`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-metadata)

objectobject

An object containing metadata about the extracted account.

[`start_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-metadata-start-date)

nullablestringnullable, string

The date of the earliest extracted transaction, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-metadata-end-date)

nullablestringnullable, string

The date of the most recent extracted transaction, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`name`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-name)

stringstring

The name of the bank account.

[`official_name`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-official-name)

nullablestringnullable, string

The official name of the bank account.

[`subtype`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-subtype)

stringstring

Valid account subtypes for depository accounts. For a list containing descriptions of each subtype, see [Account schemas](https://plaid.com/docs/api/accounts/#StandaloneAccountType-depository).  
  

Possible values: `checking`, `savings`, `hsa`, `cd`, `money market`, `paypal`, `prepaid`, `cash management`, `ebt`, `all`

[`type`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-accounts-type)

stringstring

The account type. This will always be `depository`.  
  

Possible values: `depository`

[`bank_income_sources`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources)

[object][object]

The income sources for this Item. Each entry in the array is a single income source.

[`account_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-account-id)

stringstring

The account ID with which this income source is associated.

[`income_source_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-income-source-id)

stringstring

A unique identifier for an income source. If the report is regenerated and a new `report_id` is created, the new report will have a new set of `income_source_id`s.

[`income_description`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-income-description)

stringstring

The most common name or original description for the underlying income transactions.

[`income_category`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-income-category)

stringstring

The income category.
`BANK_INTEREST`: Interest earned from a bank account.
`BENEFIT_OTHER`: Government benefits other than retirement, unemployment, child support, or disability. Currently used only in the UK, to represent benefits such as Cost of Living Payments.
`CASH`: Deprecated and used only for existing legacy implementations. Has been replaced by `CASH_DEPOSIT` and `TRANSFER_FROM_APPLICATION`.
`CASH_DEPOSIT`: A cash or check deposit.
`CHILD_SUPPORT`: Child support payments received.
`GIG_ECONOMY`: Income earned as a gig economy worker, e.g. driving for Uber, Lyft, Postmates, DoorDash, etc.
`LONG_TERM_DISABILITY`: Disability payments, including Social Security disability benefits.
`OTHER`: Income that could not be categorized as any other income category.
`MILITARY`: Veterans benefits. Income earned as salary for serving in the military (e.g. through DFAS) will be classified as `SALARY` rather than `MILITARY`.
`RENTAL`: Income earned from a rental property. Income may be identified as rental when the payment is received through a rental platform, e.g. Airbnb; rent paid directly by the tenant to the property owner (e.g. via cash, check, or ACH) will typically not be classified as rental income.
`RETIREMENT`: Payments from private retirement systems, pensions, and government retirement programs, including Social Security retirement benefits.
`SALARY`: Payment from an employer to an earner or other form of permanent employment.
`TAX_REFUND`: A tax refund.
`TRANSFER_FROM_APPLICATION`: Deposits from a money transfer app, such as Venmo, Cash App, or Zelle.
`UNEMPLOYMENT`: Unemployment benefits. In the UK, includes certain low-income benefits such as the Universal Credit.  
  

Possible values: `SALARY`, `UNEMPLOYMENT`, `CASH`, `GIG_ECONOMY`, `RENTAL`, `CHILD_SUPPORT`, `MILITARY`, `RETIREMENT`, `LONG_TERM_DISABILITY`, `BANK_INTEREST`, `CASH_DEPOSIT`, `TRANSFER_FROM_APPLICATION`, `TAX_REFUND`, `BENEFIT_OTHER`, `OTHER`

[`start_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-start-date)

stringstring

Minimum of all dates within the specific income sources in the user's bank account for days requested by the client.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-end-date)

stringstring

Maximum of all dates within the specific income sources in the user’s bank account for days requested by the client.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`pay_frequency`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-pay-frequency)

stringstring

The income pay frequency.  
  

Possible values: `WEEKLY`, `BIWEEKLY`, `SEMI_MONTHLY`, `MONTHLY`, `DAILY`, `UNKNOWN`

[`total_amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-total-amount)

numbernumber

Total amount of earnings in the user’s bank account for the specific income source for days requested by the client.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`transaction_count`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-transaction-count)

integerinteger

Number of transactions for the income source within the start and end date.

[`next_payment_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-next-payment-date)

nullablestringnullable, string

The expected date of the end user’s next paycheck for the income source.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`status`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-status)

stringstring

The status of the income sources.
`ACTIVE`: The income source is active.
`INACTIVE`: The income source is inactive.
`UNKNOWN`: The income source status is unknown.  
  

Possible values: `ACTIVE`, `INACTIVE`, `UNKNOWN`

[`historical_average_monthly_gross_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-average-monthly-gross-income)

nullablenumbernullable, number

An estimate of the average gross monthly income based on the historical net amount and income category for the income source(s).

[`historical_average_monthly_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-average-monthly-income)

nullablenumbernullable, number

The average monthly net income amount estimated based on the historical data for the income source(s).

[`forecasted_average_monthly_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-forecasted-average-monthly-income)

nullablenumbernullable, number

The predicted average monthly net income amount for the income source(s).

[`forecasted_average_monthly_income_prediction_intervals`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-forecasted-average-monthly-income-prediction-intervals)

[object][object]

The prediction interval(s) for the forecasted average monthly income.

[`lower_bound`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-forecasted-average-monthly-income-prediction-intervals-lower-bound)

nullablenumbernullable, number

The lower bound of the predicted attribute for the given probability.

[`upper_bound`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-forecasted-average-monthly-income-prediction-intervals-upper-bound)

nullablenumbernullable, number

The upper bound of the predicted attribute for the given probability.

[`probability`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-forecasted-average-monthly-income-prediction-intervals-probability)

nullablenumbernullable, number

The probability of the actual value of the attribute falling within the upper and lower bound.
This is a percentage represented as a value between 0 and 1.

[`employer`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-employer)

objectobject

The object containing employer data.

[`name`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-employer-name)

nullablestringnullable, string

The name of the employer.

[`income_provider`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-income-provider)

nullableobjectnullable, object

The object containing data about the income provider.

[`name`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-income-provider-name)

stringstring

The name of the income provider.

[`is_normalized`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-income-provider-is-normalized)

booleanboolean

Indicates whether the income provider name is normalized by comparing it against a canonical set of known providers.

[`historical_summary`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary)

[object][object]

[`total_amounts`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-total-amounts)

[object][object]

Total amount of earnings for the income source(s) of the user for the month in the summary.
This can contain multiple amounts, with each amount denominated in one unique currency.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-total-amounts-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-total-amounts-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-total-amounts-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`start_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-start-date)

stringstring

The start date of the period covered in this monthly summary.
This date will be the first day of the month, unless the month being covered is a partial month because it is the first month included in the summary and the date range being requested does not begin with the first day of the month.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-end-date)

stringstring

The end date of the period included in this monthly summary.
This date will be the last day of the month, unless the month being covered is a partial month because it is the last month included in the summary and the date range being requested does not end with the last day of the month.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`transactions`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions)

[object][object]

[`transaction_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-transaction-id)

stringstring

The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-amount)

numbernumber

The settled value of the transaction, denominated in the transaction's currency as stated in `iso_currency_code` or `unofficial_currency_code`.
Positive values when money moves out of the account; negative values when money moves in.
For example, credit card purchases are positive; credit card payment, direct deposits, and refunds are negative.

[`date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-date)

stringstring

For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted.
Both dates are returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`name`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-name)

deprecatedstringdeprecated, string

The merchant name or transaction description. This is a legacy field that is no longer maintained. For merchant name, use the `merchant_name` field; for description, use the `original_description` field.

[`original_description`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-original-description)

nullablestringnullable, string

The string returned by the financial institution to describe the transaction.

[`pending`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-pending)

booleanboolean

When true, identifies the transaction as pending or unsettled.
Pending transaction details (name, type, amount, category ID) may change before they are settled.

[`check_number`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-check-number)

nullablestringnullable, string

The check number of the transaction. This field is only populated for check transactions.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`bonus_type`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-bank-income-sources-historical-summary-transactions-bonus-type)

nullablestringnullable, string

The type of bonus that this transaction represents, if it is a bonus.
`BONUS_INCLUDED`: Bonus is included in this transaction along with the normal pay
`BONUS_ONLY`: This transaction is a standalone bonus  
  

Possible values: `BONUS_INCLUDED`, `BONUS_ONLY`, `null`

[`last_updated_time`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-last-updated-time)

stringstring

The time when this Item's data was last retrieved from the financial institution.  
  

Format: `date-time`

[`institution_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-institution-id)

stringstring

The unique identifier of the institution associated with the Item.

[`institution_name`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-items-institution-name)

stringstring

The name of the institution associated with the Item.

[`bank_income_summary`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary)

objectobject

Summary for income across all income sources and items (max history of 730 days).

[`total_amounts`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-total-amounts)

[object][object]

Total amount of earnings across all the income sources in the end user's Items for the days requested by the client.
This can contain multiple amounts, with each amount denominated in one unique currency.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-total-amounts-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-total-amounts-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-total-amounts-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`start_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-start-date)

stringstring

The earliest date within the days requested in which all income sources identified by Plaid appear in a user's account.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-end-date)

stringstring

The latest date in which all income sources identified by Plaid appear in the user's account.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`income_sources_count`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-income-sources-count)

integerinteger

Number of income sources per end user.

[`income_categories_count`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-income-categories-count)

integerinteger

Number of income categories per end user.

[`income_transactions_count`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-income-transactions-count)

integerinteger

Number of income transactions per end user.

[`historical_average_monthly_gross_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-average-monthly-gross-income)

[object][object]

An estimate of the average gross monthly income based on the historical net amount and income category for the income source(s). The average monthly income is calculated based on the lifetime of the income stream, rather than the entire historical period included in the scope of the report.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-average-monthly-gross-income-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-average-monthly-gross-income-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-average-monthly-gross-income-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`historical_average_monthly_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-average-monthly-income)

[object][object]

The average monthly income amount estimated based on the historical data for the income source(s). The average monthly income is calculated based on the lifetime of the income stream, rather than the entire historical period included in the scope of the report.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-average-monthly-income-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-average-monthly-income-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-average-monthly-income-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`forecasted_average_monthly_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-forecasted-average-monthly-income)

[object][object]

The predicted average monthly income amount for the income source(s).

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-forecasted-average-monthly-income-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-forecasted-average-monthly-income-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-forecasted-average-monthly-income-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`historical_annual_gross_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-annual-gross-income)

[object][object]

An estimate of the annual gross income for the income source, calculated by multiplying the `historical_average_monthly_gross_income` by 12.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-annual-gross-income-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-annual-gross-income-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-annual-gross-income-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`historical_annual_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-annual-income)

[object][object]

An estimate of the annual net income for the income source, calculated by multiplying the `historical_average_monthly_income` by 12.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-annual-income-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-annual-income-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-annual-income-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`forecasted_annual_income`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-forecasted-annual-income)

[object][object]

The predicted average annual income amount for the income source(s).

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-forecasted-annual-income-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-forecasted-annual-income-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-forecasted-annual-income-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`historical_summary`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary)

[object][object]

[`total_amounts`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-total-amounts)

[object][object]

Total amount of earnings for the income source(s) of the user for the month in the summary.
This can contain multiple amounts, with each amount denominated in one unique currency.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-total-amounts-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-total-amounts-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-total-amounts-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`start_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-start-date)

stringstring

The start date of the period covered in this monthly summary.
This date will be the first day of the month, unless the month being covered is a partial month because it is the first month included in the summary and the date range being requested does not begin with the first day of the month.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-end-date)

stringstring

The end date of the period included in this monthly summary.
This date will be the last day of the month, unless the month being covered is a partial month because it is the last month included in the summary and the date range being requested does not end with the last day of the month.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`transactions`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions)

[object][object]

[`transaction_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-transaction-id)

stringstring

The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.

[`amount`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-amount)

numbernumber

The settled value of the transaction, denominated in the transaction's currency as stated in `iso_currency_code` or `unofficial_currency_code`.
Positive values when money moves out of the account; negative values when money moves in.
For example, credit card purchases are positive; credit card payment, direct deposits, and refunds are negative.

[`date`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-date)

stringstring

For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted.
Both dates are returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`name`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-name)

deprecatedstringdeprecated, string

The merchant name or transaction description. This is a legacy field that is no longer maintained. For merchant name, use the `merchant_name` field; for description, use the `original_description` field.

[`original_description`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-original-description)

nullablestringnullable, string

The string returned by the financial institution to describe the transaction.

[`pending`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-pending)

booleanboolean

When true, identifies the transaction as pending or unsettled.
Pending transaction details (name, type, amount, category ID) may change before they are settled.

[`check_number`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-check-number)

nullablestringnullable, string

The check number of the transaction. This field is only populated for check transactions.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`bonus_type`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-bank-income-summary-historical-summary-transactions-bonus-type)

nullablestringnullable, string

The type of bonus that this transaction represents, if it is a bonus.
`BONUS_INCLUDED`: Bonus is included in this transaction along with the normal pay
`BONUS_ONLY`: This transaction is a standalone bonus  
  

Possible values: `BONUS_INCLUDED`, `BONUS_ONLY`, `null`

[`warnings`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-warnings)

[object][object]

If data from the report was unable to be retrieved, the warnings object will contain information about the error that caused the data to be incomplete.

[`warning_type`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-warnings-warning-type)

stringstring

The warning type which will always be `BANK_INCOME_WARNING`.  
  

Possible values: `BANK_INCOME_WARNING`

[`warning_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Unable to extract identity for the Item
`TRANSACTIONS_UNAVAILABLE`: Unable to extract transactions for the Item
`REPORT_DELETED`: Report deleted due to customer or consumer request
`DATA_UNAVAILABLE`: No relevant data was found for the Item  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `REPORT_DELETED`, `DATA_UNAVAILABLE`

[`cause`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-warnings-cause)

objectobject

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INTERNAL_SERVER_ERROR`, `INSUFFICIENT_CREDENTIALS`, `ITEM_LOCKED`, `USER_SETUP_REQUIRED`, `COUNTRY_NOT_SUPPORTED`, `INSTITUTION_DOWN`, `INSTITUTION_NO_LONGER_SUPPORTED`, `INSTITUTION_NOT_RESPONDING`, `INVALID_CREDENTIALS`, `INVALID_MFA`, `INVALID_SEND_METHOD`, `ITEM_LOGIN_REQUIRED`, `MFA_NOT_SUPPORTED`, `NO_ACCOUNTS`, `ITEM_NOT_SUPPORTED`, `ACCESS_NOT_GRANTED`

[`error_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-warnings-cause-error-code)

stringstring

We use standard HTTP response codes for success and failure notifications, and our errors are further classified by `error_type`. In general, 200 HTTP codes correspond to success, 40X codes are for developer- or user-related failures, and 50X codes are for Plaid-related issues. Error fields will be `null` if no error has occurred.

[`error_message`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/check/#cra-check_report-income_insights-get-response-report-warnings-cause-display-message)

stringstring

A user-friendly representation of the error code. null if the error is not related to user action.
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`warnings`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings)

[object][object]

If the Income Insights generation was successful but a subset of data could not be retrieved, this array will contain information about the errors causing information to be missing

[`warning_type`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-warning-type)

stringstring

The warning type, which will always be `CHECK_REPORT_WARNING`

[`warning_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Account-owner information is not available.
`TRANSACTIONS_UNAVAILABLE`: Transactions information associated with Credit and Depository accounts are unavailable.
`USER_FRAUD_ALERT`: The user has placed a fraud alert on their Plaid Check consumer report due to suspected fraud. Please note that when a fraud alert is in place, the recipient of the consumer report has an obligation to verify the consumer’s identity.  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `USER_FRAUD_ALERT`

[`cause`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause)

nullableobjectnullable, object

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`item_id`](/docs/api/products/check/#cra-check_report-income_insights-get-response-warnings-cause-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm",
  "report": {
    "report_id": "bbfc5174-5433-4648-8d93-9fec6a0c0966",
    "generated_time": "2022-01-31T22:47:53Z",
    "days_requested": 365,
    "items": [
      {
        "item_id": "AZMP7JrGXgtPd3AQMeg7hwMKgk5E8qU1V5ME7",
        "last_updated_time": "2022-01-31T22:47:53Z",
        "institution_id": "ins_0",
        "institution_name": "Plaid Bank",
        "bank_income_accounts": [
          {
            "account_id": "1qKRXQjk8xUWDJojNwPXTj8gEmR48piqRNye8",
            "mask": "8888",
            "metadata": {
              "start_date": "2024-01-01",
              "end_date": "2024-07-16"
            },
            "name": "Plaid Checking Account",
            "official_name": "Plaid Checking Account",
            "type": "depository",
            "subtype": "checking",
            "owners": []
          }
        ],
        "bank_income_sources": [
          {
            "account_id": "1qKRXQjk8xUWDJojNwPXTj8gEmR48piqRNye8",
            "income_source_id": "f17efbdd-caab-4278-8ece-963511cd3d51",
            "income_description": "PLAID_INC_DIRECT_DEP_PPD",
            "income_category": "SALARY",
            "start_date": "2021-11-15",
            "end_date": "2022-01-15",
            "pay_frequency": "MONTHLY",
            "total_amount": 300,
            "iso_currency_code": "USD",
            "unofficial_currency_code": null,
            "transaction_count": 1,
            "next_payment_date": "2022-12-15",
            "status": "ACTIVE",
            "historical_average_monthly_gross_income": 390,
            "historical_average_monthly_income": 300,
            "forecasted_average_monthly_income": 300,
            "forecasted_average_monthly_income_prediction_intervals": [
              {
                "lower_bound": 200,
                "upper_bound": 400,
                "probability": 0.8
              }
            ],
            "employer": {
              "name": "Plaid Inc"
            },
            "income_provider": {
              "name": "Plaid Inc",
              "is_normalized": true
            },
            "historical_summary": [
              {
                "start_date": "2021-11-02",
                "end_date": "2021-11-30",
                "total_amounts": [
                  {
                    "amount": 100,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                ],
                "transactions": [
                  {
                    "transaction_id": "aH5klwqG3B19OMT7D6F24Syv8pdnJXmtZoKQ5",
                    "amount": 100,
                    "bonus_type": null,
                    "date": "2021-11-15",
                    "name": "PLAID_INC_DIRECT_DEP_PPD",
                    "original_description": "PLAID_INC_DIRECT_DEP_PPD 123A",
                    "pending": false,
                    "check_number": null,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                ]
              },
              {
                "start_date": "2021-12-01",
                "end_date": "2021-12-31",
                "total_amounts": [
                  {
                    "amount": 100,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                ],
                "transactions": [
                  {
                    "transaction_id": "mN3rQ5iH8BC41T6UjKL9oD2vWJpZqXFomGwY1",
                    "amount": 100,
                    "bonus_type": null,
                    "date": "2021-12-15",
                    "name": "PLAID_INC_DIRECT_DEP_PPD",
                    "original_description": "PLAID_INC_DIRECT_DEP_PPD 123B",
                    "pending": false,
                    "check_number": null,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                ]
              },
              {
                "start_date": "2022-01-01",
                "end_date": "2022-01-31",
                "total_amounts": [
                  {
                    "amount": 100,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                ],
                "transactions": [
                  {
                    "transaction_id": "zK9lDoR8uBH51PNQ3W4T6Mjy2VFXpGtJwsL4",
                    "amount": 100,
                    "bonus_type": null,
                    "date": "2022-01-31",
                    "name": "PLAID_INC_DIRECT_DEP_PPD",
                    "original_description": "PLAID_INC_DIRECT_DEP_PPD 123C",
                    "pending": false,
                    "check_number": null,
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                ]
              }
            ]
          }
        ]
      }
    ],
    "bank_income_summary": {
      "total_amounts": [
        {
          "amount": 300,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        }
      ],
      "start_date": "2021-11-15",
      "end_date": "2022-01-15",
      "income_sources_count": 1,
      "income_categories_count": 1,
      "income_transactions_count": 1,
      "historical_average_monthly_gross_income": [
        {
          "amount": 390,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        }
      ],
      "historical_average_monthly_income": [
        {
          "amount": 300,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        }
      ],
      "forecasted_average_monthly_income": [
        {
          "amount": 300,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        }
      ],
      "historical_annual_gross_income": [
        {
          "amount": 4680,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        }
      ],
      "historical_annual_income": [
        {
          "amount": 3600,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        }
      ],
      "forecasted_annual_income": [
        {
          "amount": 3600,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        }
      ],
      "historical_summary": [
        {
          "start_date": "2021-11-02",
          "end_date": "2021-11-30",
          "total_amounts": [
            {
              "amount": 100,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            }
          ],
          "transactions": [
            {
              "transaction_id": "aH5klwqG3B19OMT7D6F24Syv8pdnJXmtZoKQ5",
              "amount": 100,
              "bonus_type": null,
              "date": "2021-11-15",
              "name": "PLAID_INC_DIRECT_DEP_PPD",
              "original_description": "PLAID_INC_DIRECT_DEP_PPD 123A",
              "pending": false,
              "check_number": null,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            }
          ]
        },
        {
          "start_date": "2021-12-01",
          "end_date": "2021-12-31",
          "total_amounts": [
            {
              "amount": 100,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            }
          ],
          "transactions": [
            {
              "transaction_id": "mN3rQ5iH8BC41T6UjKL9oD2vWJpZqXFomGwY1",
              "amount": 100,
              "bonus_type": null,
              "date": "2021-12-15",
              "name": "PLAID_INC_DIRECT_DEP_PPD",
              "original_description": "PLAID_INC_DIRECT_DEP_PPD 123B",
              "pending": false,
              "check_number": null,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            }
          ]
        },
        {
          "start_date": "2022-01-01",
          "end_date": "2022-01-31",
          "total_amounts": [
            {
              "amount": 100,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            }
          ],
          "transactions": [
            {
              "transaction_id": "zK9lDoR8uBH51PNQ3W4T6Mjy2VFXpGtJwsL4",
              "amount": 100,
              "bonus_type": null,
              "date": "2022-01-31",
              "name": "PLAID_INC_DIRECT_DEP_PPD",
              "original_description": "PLAID_INC_DIRECT_DEP_PPD 123C",
              "pending": false,
              "check_number": null,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            }
          ]
        }
      ],
      "warnings": []
    }
  }
}
```

=\*=\*=\*=

#### `/cra/check_report/network_insights/get`

#### Retrieve network attributes for the user

This endpoint allows you to retrieve the Network Insights product for your user. You should call this endpoint after you've received a `CHECK_REPORT_READY` or a `USER_CHECK_REPORT_READY` webhook, either after the Link session for the user or after calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate). If the most recent consumer report for the user doesn’t have sufficient data to generate the report, or the consumer report has expired, you will receive an error indicating that you should create a new consumer report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

If you did not initialize Link with the `cra_network_attributes` product or have generated a report using [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate), Plaid will generate the attributes when you call this endpoint.

/cra/check\_report/network\_insights/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-network_insights-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-network_insights-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-network_insights-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`options`](/docs/api/products/check/#cra-check_report-network_insights-get-request-options)

objectobject

Defines configuration options to generate Network Insights

[`network_insights_version`](/docs/api/products/check/#cra-check_report-network_insights-get-request-options-network-insights-version)

stringstring

The version of network insights  
  

Possible values: `NI1`

[`user_token`](/docs/api/products/check/#cra-check_report-network_insights-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/cra/check\_report/network\_insights/get

```
try {
  const response = await client.craCheckReportNetworkInsightsGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
  });
} catch (error) {
  // handle error
}
```

/cra/check\_report/network\_insights/get

**Response fields**

[`report`](/docs/api/products/check/#cra-check_report-network_insights-get-response-report)

objectobject

Contains data for the CRA Network Attributes Report.

[`report_id`](/docs/api/products/check/#cra-check_report-network_insights-get-response-report-report-id)

stringstring

The unique identifier associated with the report object.

[`generated_time`](/docs/api/products/check/#cra-check_report-network_insights-get-response-report-generated-time)

stringstring

The time when the report was generated.  
  

Format: `date-time`

[`network_attributes`](/docs/api/products/check/#cra-check_report-network_insights-get-response-report-network-attributes)

objectobject

A map of network attributes, where the key is a string, and the value is a float, int, or boolean. For a full list of attributes, contact your account manager.

[`items`](/docs/api/products/check/#cra-check_report-network_insights-get-response-report-items)

[object][object]

The Items the end user connected in Link.

[`institution_id`](/docs/api/products/check/#cra-check_report-network_insights-get-response-report-items-institution-id)

stringstring

The ID for the institution the user linked.

[`institution_name`](/docs/api/products/check/#cra-check_report-network_insights-get-response-report-items-institution-name)

stringstring

The name of the institution the user linked.

[`item_id`](/docs/api/products/check/#cra-check_report-network_insights-get-response-report-items-item-id)

stringstring

The identifier for the Item.

[`request_id`](/docs/api/products/check/#cra-check_report-network_insights-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`warnings`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings)

[object][object]

If the Network Insights generation was successful but a subset of data could not be retrieved, this array will contain information about the errors causing information to be missing

[`warning_type`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-warning-type)

stringstring

The warning type, which will always be `CHECK_REPORT_WARNING`

[`warning_code`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Account-owner information is not available.
`TRANSACTIONS_UNAVAILABLE`: Transactions information associated with Credit and Depository accounts are unavailable.
`USER_FRAUD_ALERT`: The user has placed a fraud alert on their Plaid Check consumer report due to suspected fraud. Please note that when a fraud alert is in place, the recipient of the consumer report has an obligation to verify the consumer’s identity.  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `USER_FRAUD_ALERT`

[`cause`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause)

nullableobjectnullable, object

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`item_id`](/docs/api/products/check/#cra-check_report-network_insights-get-response-warnings-cause-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm",
  "report": {
    "report_id": "ee093cb0-e3f2-42d1-9dbc-8d8408964194",
    "generated_time": "2022-01-31T22:47:53Z",
    "network_attributes": {
      "plaid_conn_user_lifetime_lending_count": 5,
      "plaid_conn_user_lifetime_personal_lending_flag": 1,
      "plaid_conn_user_lifetime_cash_advance_primary_count": 0
    },
    "items": [
      {
        "institution_id": "ins_0",
        "institution_name": "Plaid Bank",
        "item_id": "AZMP7JrGXgtPd3AQMeg7hwMKgk5E8qU1V5ME7"
      }
    ]
  }
}
```

=\*=\*=\*=

#### `/cra/check_report/partner_insights/get`

#### Retrieve cash flow insights from partners

This endpoint allows you to retrieve the Partner Insights report for your user. You should call this endpoint after you've received a `CHECK_REPORT_READY` or a `USER_CHECK_REPORT_READY` webhook, either after the Link session for the user or after calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate). If the most recent consumer report for the user doesn’t have sufficient data to generate the base report, or the consumer report has expired, you will receive an error indicating that you should create a new consumer report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

If you did not initialize Link with the `credit_partner_insights` product or have generated a report using [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate), we will call our partners to generate the insights when you call this endpoint. In this case, you may optionally provide parameters under `options` to configure which insights you want to receive.

/cra/check\_report/partner\_insights/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`user_token`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`partner_insights`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-partner-insights)

objectobject

Defines configuration to generate Partner Insights

[`prism_versions`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-partner-insights-prism-versions)

objectobject

The versions of Prism products to evaluate

[`firstdetect`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-partner-insights-prism-versions-firstdetect)

stringstring

The version of Prism FirstDetect. If not specified, will default to v3.  
  

Possible values: `3`, `null`

[`detect`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-partner-insights-prism-versions-detect)

stringstring

The version of Prism Detect  
  

Possible values: `4`, `null`

[`cashscore`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-partner-insights-prism-versions-cashscore)

stringstring

The version of Prism CashScore. If not specified, will default to v3.  
  

Possible values: `4`, `3`, `null`

[`extend`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-partner-insights-prism-versions-extend)

stringstring

The version of Prism Extend  
  

Possible values: `4`, `null`

[`insights`](/docs/api/products/check/#cra-check_report-partner_insights-get-request-partner-insights-prism-versions-insights)

stringstring

The version of Prism Insights. If not specified, will default to v3.  
  

Possible values: `4`, `3`, `null`

/cra/check\_report/partner\_insights/get

```
try {
  const response = await client.craCheckReportPartnerInsightsGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
  });
} catch (error) {
  // handle error
}
```

/cra/check\_report/partner\_insights/get

**Response fields**

[`report`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report)

objectobject

The Partner Insights report of the bank data for an end user.

[`report_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-report-id)

stringstring

A unique identifier associated with the Partner Insights object.

[`generated_time`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-generated-time)

stringstring

The time when the Partner Insights report was generated.  
  

Format: `date-time`

[`client_report_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-client-report-id)

nullablestringnullable, string

Client-generated identifier, which can be used by lenders to track loan applications.

[`prism`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism)

objectobject

The Prism Data insights for the user.

[`insights`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-insights)

nullableobjectnullable, object

The data from the Insights product returned by Prism Data.

[`version`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-insights-version)

integerinteger

The version of Prism Data's insights model used.

[`result`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-insights-result)

objectobject

The Insights Result object is a map of cash flow attributes, where the key is a string, and the value is a float or string. For a full list of attributes, contact your account manager. The attributes may vary depending on the Prism version used.

[`error_reason`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-insights-error-reason)

stringstring

The error returned by Prism for this product.

[`cash_score`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score)

nullableobjectnullable, object

The data from the CashScore® product returned by Prism Data.

[`version`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-version)

deprecatedintegerdeprecated, integer

The version of Prism Data's cash score model used. This field is deprecated in favor of `model_version`.

[`model_version`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-model-version)

stringstring

The version of Prism Data's cash score model used.

[`score`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-score)

nullableintegernullable, integer

The score returned by Prism Data. Ranges from 1-999, with higher score indicating lower risk.

[`reason_codes`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-reason-codes)

[string][string]

The reasons for an individual having risk according to the cash score.

[`metadata`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata)

objectobject

An object containing metadata about the provided transactions.

[`max_age`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-max-age)

nullableintegernullable, integer

Number of days since the oldest transaction.

[`min_age`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-min-age)

nullableintegernullable, integer

Number of days since the latest transaction.

[`min_age_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-min-age-credit)

nullableintegernullable, integer

Number of days since the latest credit transaction.

[`min_age_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-min-age-debit)

nullableintegernullable, integer

Number of days since the latest debit transaction.

[`max_age_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-max-age-debit)

nullableintegernullable, integer

Number of days since the oldest debit transaction.

[`max_age_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-max-age-credit)

nullableintegernullable, integer

Number of days since the oldest credit transaction.

[`num_trxn_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-num-trxn-credit)

nullableintegernullable, integer

Number of credit transactions.

[`num_trxn_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-num-trxn-debit)

nullableintegernullable, integer

Number of debit transactions.

[`l1m_credit_value_cnt`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-l1m-credit-value-cnt)

nullableintegernullable, integer

Number of credit transactions in the last 30 days.

[`l1m_debit_value_cnt`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-metadata-l1m-debit-value-cnt)

nullableintegernullable, integer

Number of debit transactions in the last 30 days.

[`error_reason`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-cash-score-error-reason)

stringstring

The error returned by Prism for this product.

[`extend`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend)

nullableobjectnullable, object

The data from the CashScore® Extend product returned by Prism Data.

[`model_version`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-model-version)

stringstring

The version of Prism Data's CashScore® Extend model used.

[`score`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-score)

nullableintegernullable, integer

The score returned by Prism Data. Ranges from 1-999, with higher score indicating lower risk.

[`reason_codes`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-reason-codes)

[string][string]

The reasons for an individual having risk according to the CashScore® Extend score.

[`metadata`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata)

objectobject

An object containing metadata about the provided transactions.

[`max_age`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-max-age)

nullableintegernullable, integer

Number of days since the oldest transaction.

[`min_age`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-min-age)

nullableintegernullable, integer

Number of days since the latest transaction.

[`min_age_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-min-age-credit)

nullableintegernullable, integer

Number of days since the latest credit transaction.

[`min_age_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-min-age-debit)

nullableintegernullable, integer

Number of days since the latest debit transaction.

[`max_age_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-max-age-debit)

nullableintegernullable, integer

Number of days since the oldest debit transaction.

[`max_age_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-max-age-credit)

nullableintegernullable, integer

Number of days since the oldest credit transaction.

[`num_trxn_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-num-trxn-credit)

nullableintegernullable, integer

Number of credit transactions.

[`num_trxn_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-num-trxn-debit)

nullableintegernullable, integer

Number of debit transactions.

[`l1m_credit_value_cnt`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-l1m-credit-value-cnt)

nullableintegernullable, integer

Number of credit transactions in the last 30 days.

[`l1m_debit_value_cnt`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-metadata-l1m-debit-value-cnt)

nullableintegernullable, integer

Number of debit transactions in the last 30 days.

[`error_reason`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-extend-error-reason)

stringstring

The error returned by Prism for this product.

[`first_detect`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect)

nullableobjectnullable, object

The data from the FirstDetect product returned by Prism Data.

[`version`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-version)

deprecatedintegerdeprecated, integer

The version of Prism Data's FirstDetect model used. This field is deprecated in favor of `model_version`.

[`model_version`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-model-version)

stringstring

The version of Prism Data's FirstDetect model used.

[`score`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-score)

nullableintegernullable, integer

The score returned by Prism Data. Ranges from 1-999, with higher score indicating lower risk.

[`reason_codes`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-reason-codes)

[string][string]

The reasons for an individual having risk according to the FirstDetect score.

[`metadata`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata)

objectobject

An object containing metadata about the provided transactions.

[`max_age`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-max-age)

nullableintegernullable, integer

Number of days since the oldest transaction.

[`min_age`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-min-age)

nullableintegernullable, integer

Number of days since the latest transaction.

[`min_age_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-min-age-credit)

nullableintegernullable, integer

Number of days since the latest credit transaction.

[`min_age_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-min-age-debit)

nullableintegernullable, integer

Number of days since the latest debit transaction.

[`max_age_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-max-age-debit)

nullableintegernullable, integer

Number of days since the oldest debit transaction.

[`max_age_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-max-age-credit)

nullableintegernullable, integer

Number of days since the oldest credit transaction.

[`num_trxn_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-num-trxn-credit)

nullableintegernullable, integer

Number of credit transactions.

[`num_trxn_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-num-trxn-debit)

nullableintegernullable, integer

Number of debit transactions.

[`l1m_credit_value_cnt`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-l1m-credit-value-cnt)

nullableintegernullable, integer

Number of credit transactions in the last 30 days.

[`l1m_debit_value_cnt`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-metadata-l1m-debit-value-cnt)

nullableintegernullable, integer

Number of debit transactions in the last 30 days.

[`error_reason`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-first-detect-error-reason)

stringstring

The error returned by Prism for this product.

[`detect`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect)

nullableobjectnullable, object

The data from the CashScore® Detect product returned by Prism Data.

[`model_version`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-model-version)

stringstring

The version of Prism Data's CashScore® Detect model used.

[`score`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-score)

nullableintegernullable, integer

The score returned by Prism Data. Ranges from 1-999, with higher score indicating lower risk.

[`reason_codes`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-reason-codes)

[string][string]

The reasons for an individual having risk according to the CashScore® Detect score.

[`metadata`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata)

objectobject

An object containing metadata about the provided transactions.

[`max_age`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-max-age)

nullableintegernullable, integer

Number of days since the oldest transaction.

[`min_age`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-min-age)

nullableintegernullable, integer

Number of days since the latest transaction.

[`min_age_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-min-age-credit)

nullableintegernullable, integer

Number of days since the latest credit transaction.

[`min_age_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-min-age-debit)

nullableintegernullable, integer

Number of days since the latest debit transaction.

[`max_age_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-max-age-debit)

nullableintegernullable, integer

Number of days since the oldest debit transaction.

[`max_age_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-max-age-credit)

nullableintegernullable, integer

Number of days since the oldest credit transaction.

[`num_trxn_credit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-num-trxn-credit)

nullableintegernullable, integer

Number of credit transactions.

[`num_trxn_debit`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-num-trxn-debit)

nullableintegernullable, integer

Number of debit transactions.

[`l1m_credit_value_cnt`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-l1m-credit-value-cnt)

nullableintegernullable, integer

Number of credit transactions in the last 30 days.

[`l1m_debit_value_cnt`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-metadata-l1m-debit-value-cnt)

nullableintegernullable, integer

Number of debit transactions in the last 30 days.

[`error_reason`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-detect-error-reason)

stringstring

The error returned by Prism for this product.

[`status`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-prism-status)

stringstring

Details on whether the Prism Data attributes succeeded or failed to be generated.

[`items`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items)

[object][object]

The list of Items used in the report along with the associated metadata about the Item.

[`institution_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-institution-id)

stringstring

The ID for the institution that the user linked.

[`institution_name`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-institution-name)

stringstring

The name of the institution the user linked.

[`item_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-item-id)

stringstring

The identifier for the item.

[`accounts`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts)

[object][object]

A list of accounts in the item

[`account_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-account-id)

stringstring

Plaid's unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`mask`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number.
Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.

[`metadata`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-metadata)

objectobject

An object containing metadata about the extracted account.

[`start_date`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-metadata-start-date)

nullablestringnullable, string

The date of the earliest extracted transaction, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-metadata-end-date)

nullablestringnullable, string

The date of the most recent extracted transaction, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`name`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-name)

stringstring

The name of the account

[`official_name`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-official-name)

nullablestringnullable, string

The official name of the bank account.

[`subtype`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-subtype)

stringstring

Valid account subtypes for depository accounts. For a list containing descriptions of each subtype, see [Account schemas](https://plaid.com/docs/api/accounts/#StandaloneAccountType-depository).  
  

Possible values: `checking`, `savings`, `hsa`, `cd`, `money market`, `paypal`, `prepaid`, `cash management`, `ebt`, `all`

[`type`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-report-items-accounts-type)

stringstring

The account type. This will always be `depository`.  
  

Possible values: `depository`

[`request_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`warnings`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings)

[object][object]

If the Partner Insights generation was successful but a subset of data could not be retrieved, this array will contain information about the errors causing information to be missing

[`warning_type`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-warning-type)

stringstring

The warning type, which will always be `CHECK_REPORT_WARNING`

[`warning_code`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Account-owner information is not available.
`TRANSACTIONS_UNAVAILABLE`: Transactions information associated with Credit and Depository accounts are unavailable.
`USER_FRAUD_ALERT`: The user has placed a fraud alert on their Plaid Check consumer report due to suspected fraud. Please note that when a fraud alert is in place, the recipient of the consumer report has an obligation to verify the consumer’s identity.  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `USER_FRAUD_ALERT`

[`cause`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause)

nullableobjectnullable, object

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`item_id`](/docs/api/products/check/#cra-check_report-partner_insights-get-response-warnings-cause-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm",
  "report": {
    "report_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
    "client_report_id": "client_report_id_1221",
    "generated_time": "2022-01-31T22:47:53Z",
    "items": [
      {
        "institution_id": "ins_109508",
        "institution_name": "Plaid Bank",
        "item_id": "Ed6bjNrDLJfGvZWwnkQlfxwoNz54B5C97ejBr",
        "accounts": [
          {
            "account_id": "1qKRXQjk8xUWDJojNwPXTj8gEmR48piqRNye8",
            "mask": "8888",
            "metadata": {
              "start_date": "2022-01-01",
              "end_date": "2022-01-31"
            },
            "name": "Plaid Checking Account",
            "official_name": "Plaid Checking Account",
            "type": "depository",
            "subtype": "checking",
            "owners": []
          }
        ]
      }
    ],
    "prism": {
      "insights": {
        "version": 3,
        "result": {
          "l6m_cumbal_acc": 1
        }
      },
      "cash_score": {
        "version": 3,
        "model_version": "3",
        "score": 900,
        "reason_codes": [
          "CS03038"
        ],
        "metadata": {
          "max_age": 20,
          "min_age": 1,
          "min_age_credit": 0,
          "min_age_debit": 1,
          "max_age_debit": 20,
          "max_age_credit": 0,
          "num_trxn_credit": 0,
          "num_trxn_debit": 40,
          "l1m_credit_value_cnt": 0,
          "l1m_debit_value_cnt": 40
        }
      },
      "first_detect": {
        "version": 3,
        "model_version": "3",
        "score": 900,
        "reason_codes": [
          "CS03038"
        ],
        "metadata": {
          "max_age": 20,
          "min_age": 1,
          "min_age_credit": 0,
          "min_age_debit": 1,
          "max_age_debit": 20,
          "max_age_credit": 0,
          "num_trxn_credit": 0,
          "num_trxn_debit": 40,
          "l1m_credit_value_cnt": 0,
          "l1m_debit_value_cnt": 40
        }
      },
      "status": "SUCCESS"
    }
  }
}
```

=\*=\*=\*=

#### `/cra/check_report/pdf/get`

#### Retrieve Consumer Reports as a PDF

[`/cra/check_report/pdf/get`](/docs/api/products/check/#cracheck_reportpdfget) retrieves the most recent Consumer Report in PDF format. By default, the most recent Base Report (if it exists) for the user will be returned. To request that the most recent Partner Insights or Income Insights report be included in the PDF as well, use the `add-ons` field.

/cra/check\_report/pdf/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-pdf-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-pdf-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-pdf-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`add_ons`](/docs/api/products/check/#cra-check_report-pdf-get-request-add-ons)

[string][string]

Use this field to include other reports in the PDF.  
  

Possible values: `cra_income_insights`, `cra_partner_insights`

[`user_token`](/docs/api/products/check/#cra-check_report-pdf-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/cra/check\_report/pdf/get

```
try {
  const response = await client.craCheckReportPDFGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
  });
  const pdf = response.buffer.toString('base64');
} catch (error) {
  // handle error
}
```

##### Response

This endpoint returns binary PDF data. [View a sample Check Report PDF.](https://plaid.com/documents/sample-check-report.pdf)
[View a sample Check Report PDF containing Income Insights.](https://plaid.com/documents/sample-check-report-with-income.pdf)

=\*=\*=\*=

#### `/cra/check_report/cashflow_insights/get`

#### Retrieve cash flow insights from your user's banking data

This endpoint allows you to retrieve the Cashflow Insights report for your user. You should call this endpoint after you've received a `CHECK_REPORT_READY` or a `USER_CHECK_REPORT_READY` webhook, either after the Link session for the user or after calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate). If the most recent consumer report for the user doesn’t have sufficient data to generate the insights, or the consumer report has expired, you will receive an error indicating that you should create a new consumer report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

If you did not initialize Link with the `cra_cashflow_insights` product or have generated a report using [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate), we will generate the insights when you call this endpoint. In this case, you may optionally provide parameters under `options` to configure which insights you want to receive.

/cra/check\_report/cashflow\_insights/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`user_token`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`options`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-request-options)

objectobject

Defines configuration options to generate Cashflow Insights

[`attributes_version`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-request-options-attributes-version)

stringstring

The version of cashflow attributes  
  

Possible values: `CFI1`

/cra/check\_report/cashflow\_insights/get

```
try {
  const response = await client.craCheckReportCashflowInsightsGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
  });
} catch (error) {
  // handle error
}
```

/cra/check\_report/cashflow\_insights/get

**Response fields**

[`report`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-report)

objectobject

Contains data for the CRA Cashflow Insights Report.

[`report_id`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-report-report-id)

stringstring

The unique identifier associated with the report object.

[`generated_time`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-report-generated-time)

stringstring

The time when the report was generated.  
  

Format: `date-time`

[`attributes`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-report-attributes)

objectobject

A map of cash flow attributes, where the key is a string, and the value is a float, int, or boolean. The specific list of attributes will depend on the cash flow attributes version used. For a full list of attributes, contact your account manager.

[`request_id`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`warnings`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings)

[object][object]

If the Cashflow Insights generation was successful but a subset of data could not be retrieved, this array will contain information about the errors causing information to be missing

[`warning_type`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-warning-type)

stringstring

The warning type, which will always be `CHECK_REPORT_WARNING`

[`warning_code`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Account-owner information is not available.
`TRANSACTIONS_UNAVAILABLE`: Transactions information associated with Credit and Depository accounts are unavailable.
`USER_FRAUD_ALERT`: The user has placed a fraud alert on their Plaid Check consumer report due to suspected fraud. Please note that when a fraud alert is in place, the recipient of the consumer report has an obligation to verify the consumer’s identity.  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `USER_FRAUD_ALERT`

[`cause`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause)

nullableobjectnullable, object

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`item_id`](/docs/api/products/check/#cra-check_report-cashflow_insights-get-response-warnings-cause-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm",
  "report": {
    "report_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
    "generated_time": "2022-01-31T22:47:53Z",
    "attributes": {
      "cash_reliance_atm_withdrawal_amt_cv_90d": 180.1
    }
  }
}
```

=\*=\*=\*=

#### `/cra/check_report/lend_score/get`

#### Retrieve the LendScore from your user's banking data

This endpoint allows you to retrieve the LendScore report for your user. You should call this endpoint after you've received a `CHECK_REPORT_READY` or a `USER_CHECK_REPORT_READY` webhook, either after the Link session for the user or after calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate). If the most recent consumer report for the user doesn’t have sufficient data to generate the insights, or the consumer report has expired, you will receive an error indicating that you should create a new consumer report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

If you did not initialize Link with the `cra_lend_score` product or call [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) with the `cra_lend_score` product, Plaid will generate the insights when you call this endpoint. In this case, you may optionally provide parameters under `options` to configure which insights you want to receive.

/cra/check\_report/lend\_score/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-lend_score-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-lend_score-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-lend_score-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`user_token`](/docs/api/products/check/#cra-check_report-lend_score-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`options`](/docs/api/products/check/#cra-check_report-lend_score-get-request-options)

objectobject

Defines configuration options to generate the LendScore

[`lend_score_version`](/docs/api/products/check/#cra-check_report-lend_score-get-request-options-lend-score-version)

stringstring

The version of the LendScore  
  

Possible values: `LS1`

/cra/check\_report/lend\_score/get

```
try {
  const response = await client.craCheckReportLendScoreGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
  });
} catch (error) {
  // handle error
}
```

/cra/check\_report/lend\_score/get

**Response fields**

[`report`](/docs/api/products/check/#cra-check_report-lend_score-get-response-report)

objectobject

Contains data for the CRA LendScore Report.

[`report_id`](/docs/api/products/check/#cra-check_report-lend_score-get-response-report-report-id)

stringstring

The unique identifier associated with the report object.

[`generated_time`](/docs/api/products/check/#cra-check_report-lend_score-get-response-report-generated-time)

stringstring

The time when the report was generated.  
  

Format: `date-time`

[`lend_score`](/docs/api/products/check/#cra-check_report-lend_score-get-response-report-lend-score)

nullableobjectnullable, object

The results of the LendScore

[`score`](/docs/api/products/check/#cra-check_report-lend_score-get-response-report-lend-score-score)

nullableintegernullable, integer

The score returned by the LendScore model. Will be an integer in the range 1 to 99. Higher scores indicate lower credit risk.

[`reason_codes`](/docs/api/products/check/#cra-check_report-lend_score-get-response-report-lend-score-reason-codes)

[string][string]

The reasons for an individual having risk according to the LendScore. For a full list of possible reason codes, contact your Plaid Account Manager. Different LendScore versions will use different sets of reason codes.

[`error_reason`](/docs/api/products/check/#cra-check_report-lend_score-get-response-report-lend-score-error-reason)

nullablestringnullable, string

Human-readable description of why the LendScore could not be computed.

[`request_id`](/docs/api/products/check/#cra-check_report-lend_score-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`warnings`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings)

[object][object]

If the LendScore generation was successful but a subset of data could not be retrieved, this array will contain information about the errors causing information to be missing

[`warning_type`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-warning-type)

stringstring

The warning type, which will always be `CHECK_REPORT_WARNING`

[`warning_code`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Account-owner information is not available.
`TRANSACTIONS_UNAVAILABLE`: Transactions information associated with Credit and Depository accounts are unavailable.
`USER_FRAUD_ALERT`: The user has placed a fraud alert on their Plaid Check consumer report due to suspected fraud. Please note that when a fraud alert is in place, the recipient of the consumer report has an obligation to verify the consumer’s identity.  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `USER_FRAUD_ALERT`

[`cause`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause)

nullableobjectnullable, object

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`item_id`](/docs/api/products/check/#cra-check_report-lend_score-get-response-warnings-cause-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm",
  "report": {
    "report_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
    "generated_time": "2022-01-31T22:47:53Z",
    "lend_score": {
      "score": 80,
      "reason_codes": [
        "Bank Balance Volatility"
      ]
    }
  }
}
```

=\*=\*=\*=

#### `/cra/check_report/verification/get`

#### Retrieve various home lending reports for a user.

This endpoint allows you to retrieve home lending reports for a user. To obtain a VoA or Employment Refresh report, you need to make sure that `cra_base_report` is included in the `products` parameter when calling [`/link/token/create`](/docs/api/link/#linktokencreate) or [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

You should call this endpoint after you've received a `CHECK_REPORT_READY` or a `USER_CHECK_REPORT_READY` webhook, either after the Link session for the user or after calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

If the most recent consumer report for the user doesn’t have sufficient data to generate the report, or the consumer report has expired, you will receive an error indicating that you should create a new consumer report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate)."

/cra/check\_report/verification/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-verification-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-verification-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-verification-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`reports_requested`](/docs/api/products/check/#cra-check_report-verification-get-request-reports-requested)

required[string]required, [string]

Specifies which types of home lending reports are expected in the response  
  

Possible values: `VOA`, `EMPLOYMENT_REFRESH`

[`employment_refresh_options`](/docs/api/products/check/#cra-check_report-verification-get-request-employment-refresh-options)

objectobject

Defines configuration options for the Employment Refresh Report.

[`days_requested`](/docs/api/products/check/#cra-check_report-verification-get-request-employment-refresh-options-days-requested)

requiredintegerrequired, integer

The number of days of data to request for the report. This field is required if an Employment Refresh Report is requested. Maximum is 731.  
  

Maximum: `731`

[`user_token`](/docs/api/products/check/#cra-check_report-verification-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/cra/check\_report/verification/get

```
try {
  const response = await client.craCheckReportVerificationGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
    reports_requested: ['VOA', 'EMPLOYMENT_REFRESH'],
    employment_refresh_options: {
      days_requested: 60,
    },
  });
} catch (error) {
  // handle error
}
```

/cra/check\_report/verification/get

**Response fields**

[`report`](/docs/api/products/check/#cra-check_report-verification-get-response-report)

objectobject

Contains data for the CRA Home Lending Report.

[`report_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-report-id)

stringstring

The unique identifier associated with the Home Lending Report object. This ID will be the same as the Base Report ID.

[`client_report_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-client-report-id)

nullablestringnullable, string

Client-generated identifier, which can be used by lenders to track loan applications.

[`voa`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa)

nullableobjectnullable, object

An object representing a VOA report.

[`generated_time`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-generated-time)

stringstring

The date and time when the VOA Report was created, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (e.g. "2018-04-12T03:32:11Z").  
  

Format: `date-time`

[`days_requested`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-days-requested)

numbernumber

The number of days of transaction history that the VOA report covers.

[`items`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items)

[object][object]

Data returned by Plaid about each of the Items included in the Base Report.

[`accounts`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts)

[object][object]

Data about each of the accounts open on the Item.

[`account_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`balances`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances)

objectobject

VOA Report information about an account's balances.

[`available`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-available)

nullablenumbernullable, number

The amount of funds available to be withdrawn from the account, as determined by the financial institution.  
For `credit`-type accounts, the `available` balance typically equals the `limit` less the `current` balance, less any pending outflows plus any pending inflows.  
For `depository`-type accounts, the `available` balance typically equals the `current` balance less any pending outflows plus any pending inflows. For `depository`-type accounts, the `available` balance does not include the overdraft limit.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the `available` balance is the total cash available to withdraw as presented by the institution.  
Note that not all institutions calculate the `available` balance. In the event that `available` balance is unavailable, Plaid will return an `available` balance value of `null`.  
Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by `/accounts/balance/get`.  
If `current` is `null` this field is guaranteed not to be `null`.  
  

Format: `double`

[`current`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-current)

nullablenumbernullable, number

The total amount of funds in or owed by the account.  
For `credit`-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  
For `loan`-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (`ins_116944`). For Sallie Mae student loans, the account's balance includes both principal and any outstanding interest.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the current balance is the total value of assets as presented by the institution.  
Note that balance information may be cached unless the value was returned by `/accounts/balance/get`; if the Item is enabled for Transactions, the balance will be at least as recent as the most recent Transaction update. If you require realtime balance information, use the `available` balance as provided by `/accounts/balance/get`.  
When returned by `/accounts/balance/get`, this field may be `null`. When this happens, `available` is guaranteed not to be `null`.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the balance. Always null if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the balance. Always null if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`historical_balances`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-historical-balances)

[object][object]

Calculated data about the historical balances on the account.  
Available for `credit` and `depository` type accounts.

[`current`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-historical-balances-current)

numbernumber

The total amount of funds in the account, calculated from the `current` balance in the `balance` object by subtracting inflows and adding back outflows according to the posted date of each transaction.  
  

Format: `double`

[`date`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-historical-balances-date)

stringstring

The date of the calculated historical balance, in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-historical-balances-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the balance. Always `null` if `unofficial_currency_code` is non-`null`.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-historical-balances-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the balance. Always `null` if `iso_currency_code` is non-`null`.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`average_balance_30_days`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-average-balance-30-days)

nullablenumbernullable, number

The average balance in the account over the last 30 days. Calculated using the derived historical balances.  
  

Format: `double`

[`average_balance_60_days`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-average-balance-60-days)

nullablenumbernullable, number

The average balance in the account over the last 60 days. Calculated using the derived historical balances.  
  

Format: `double`

[`nsf_overdraft_transactions_count`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-balances-nsf-overdraft-transactions-count)

numbernumber

The number of net NSF fee transactions in the time range for the report in the given account (not counting any fees that were reversed within the time range).

[`consumer_disputes`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-consumer-disputes)

[object][object]

The information about previously submitted valid dispute statements by the consumer

[`consumer_dispute_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-consumer-disputes-consumer-dispute-id)

deprecatedstringdeprecated, string

(Deprecated) A unique identifier (UUID) of the consumer dispute that can be used for troubleshooting

[`dispute_field_create_date`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-consumer-disputes-dispute-field-create-date)

stringstring

Date of the disputed field (e.g. transaction date), in an ISO 8601 format (YYYY-MM-DD)  
  

Format: `date`

[`category`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-consumer-disputes-category)

stringstring

Type of data being disputed by the consumer  
  

Possible values: `TRANSACTION`, `BALANCE`, `IDENTITY`, `OTHER`

[`statement`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-consumer-disputes-statement)

stringstring

Text content of dispute

[`mask`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.

[`name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself.

[`official_name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-official-name)

nullablestringnullable, string

The official name of the account as given by the financial institution.

[`type`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`days_available`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-days-available)

numbernumber

The duration of transaction history available within this report for this Item, typically defined as the time since the date of the earliest transaction in that account.

[`transactions_insights`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights)

objectobject

Transaction data associated with the account.

[`all_transactions`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions)

[object][object]

Transaction history associated with the account.

[`account_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-account-id)

stringstring

The ID of the account in which this transaction occurred.

[`amount`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-amount)

numbernumber

The settled value of the transaction, denominated in the transaction's currency, as stated in `iso_currency_code` or `unofficial_currency_code`. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the transaction. Always `null` if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the transaction. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`original_description`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-original-description)

nullablestringnullable, string

The string returned by the financial institution to describe the transaction.

[`credit_category`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-credit-category)

nullableobjectnullable, object

Information describing the intent of the transaction. Most relevant for credit use cases, but not limited to such use cases.  
See the [`taxonomy csv file`](https://plaid.com/documents/credit-category-taxonomy.csv) for a full list of credit categories.

[`primary`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-credit-category-primary)

stringstring

A high level category that communicates the broad category of the transaction.

[`detailed`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-credit-category-detailed)

stringstring

A granular category conveying the transaction's intent. This field can also be used as a unique identifier for the category.

[`check_number`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-check-number)

nullablestringnullable, string

The check number of the transaction. This field is only populated for check transactions.

[`date`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-date)

stringstring

For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DD` ).  
  

Format: `date`

[`date_transacted`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-date-transacted)

nullablestringnullable, string

The date on which the transaction took place, in IS0 8601 format.

[`location`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location)

objectobject

A representation of where a transaction took place. Location data is provided only for transactions at physical locations, not for online transactions. Location data availability depends primarily on the merchant and is most likely to be populated for transactions at large retail chains; small, local businesses are less likely to have location data available.

[`address`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location-address)

nullablestringnullable, string

The street address where the transaction occurred.

[`city`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location-city)

nullablestringnullable, string

The city where the transaction occurred.

[`region`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location-region)

nullablestringnullable, string

The region or state where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called `state`.

[`postal_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location-postal-code)

nullablestringnullable, string

The postal code where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code where the transaction occurred.

[`lat`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location-lat)

nullablenumbernullable, number

The latitude where the transaction occurred.  
  

Format: `double`

[`lon`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location-lon)

nullablenumbernullable, number

The longitude where the transaction occurred.  
  

Format: `double`

[`store_number`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-location-store-number)

nullablestringnullable, string

The merchant defined store number where the transaction occurred.

[`merchant_name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-merchant-name)

nullablestringnullable, string

The merchant name, as enriched by Plaid from the `name` field. This is typically a more human-readable version of the merchant counterparty in the transaction. For some bank transactions (such as checks or account transfers) where there is no meaningful merchant name, this value will be `null`.

[`pending`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-pending)

booleanboolean

When `true`, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled.

[`account_owner`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-account-owner)

nullablestringnullable, string

The name of the account owner. This field is not typically populated and only relevant when dealing with sub-accounts.

[`transaction_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-all-transactions-transaction-id)

stringstring

The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.

[`end_date`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-end-date)

nullablestringnullable, string

The latest timeframe provided by the FI, in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`start_date`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-transactions-insights-start-date)

nullablestringnullable, string

The earliest timeframe provided by the FI, in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`owners`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners)

[object][object]

Data returned by the financial institution about the account owner or owners.

[`names`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-names)

[string][string]

A list of names associated with the account by the financial institution. In the case of a joint account, Plaid will make a best effort to report the names of all account holders.  
If an Item contains multiple accounts with different owner names, some institutions will report all names associated with the Item in each account's `names` array.

[`phone_numbers`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-phone-numbers)

[object][object]

A list of phone numbers associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-phone-numbers-data)

stringstring

The phone number.

[`primary`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-phone-numbers-primary)

booleanboolean

When `true`, identifies the phone number as the primary number on an account.

[`type`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-phone-numbers-type)

stringstring

The type of phone number.  
  

Possible values: `home`, `work`, `office`, `mobile`, `mobile1`, `other`

[`emails`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-emails)

[object][object]

A list of email addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-emails-data)

stringstring

The email address.

[`primary`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-emails-primary)

booleanboolean

When `true`, identifies the email address as the primary email on an account.

[`type`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-emails-type)

stringstring

The type of email account as described by the financial institution.  
  

Possible values: `primary`, `secondary`, `other`

[`addresses`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-addresses)

[object][object]

Data about the various addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-addresses-data)

objectobject

Data about the components comprising an address.

[`city`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-addresses-data-city)

nullablestringnullable, string

The full city name

[`region`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-addresses-data-region)

nullablestringnullable, string

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-addresses-data-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-addresses-data-postal-code)

nullablestringnullable, string

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-addresses-data-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code

[`primary`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-owners-addresses-primary)

booleanboolean

When `true`, identifies the address as the primary address on an account.

[`ownership_type`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-ownership-type)

nullablestringnullable, string

How an asset is owned.  
`association`: Ownership by a corporation, partnership, or unincorporated association, including for-profit and not-for-profit organizations.
`individual`: Ownership by an individual.
`joint`: Joint ownership by multiple parties.
`trust`: Ownership by a revocable or irrevocable trust.  
  

Possible values: `null`, `individual`, `joint`, `association`, `trust`

[`investments`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments)

nullableobjectnullable, object

A set of fields describing the investments data on an account.

[`holdings`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings)

[object][object]

Quantities and values of securities held in the investment account. Map to the `securities` array for security details.

[`account_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-account-id)

stringstring

The Plaid `account_id` associated with the holding.

[`security_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-security-id)

stringstring

The Plaid `security_id` associated with the holding. Security data is not specific to a user's account; any user who held the same security at the same financial institution at the same time would have identical security data. The `security_id` for the same security will typically be the same across different institutions, but this is not guaranteed. The `security_id` does not typically change, but may change if inherent details of the security change due to a corporate action, for example, in the event of a ticker symbol change or CUSIP change.

[`institution_price`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-institution-price)

numbernumber

The last price given by the institution for this security.  
  

Format: `double`

[`institution_price_as_of`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-institution-price-as-of)

nullablestringnullable, string

The date at which `institution_price` was current.  
  

Format: `date`

[`institution_value`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-institution-value)

numbernumber

The value of the holding, as reported by the institution.  
  

Format: `double`

[`cost_basis`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-cost-basis)

nullablenumbernullable, number

The original total value of the holding. This field is calculated by Plaid as the sum of the purchase price of all of the shares in the holding.  
  

Format: `double`

[`quantity`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-quantity)

numbernumber

The total quantity of the asset held, as reported by the financial institution. If the security is an option, `quantity` will reflect the total number of options (typically the number of contracts multiplied by 100), not the number of contracts.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the holding. Always `null` if `unofficial_currency_code` is non-`null`.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-holdings-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the holding. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`securities`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities)

[object][object]

Details of specific securities held in the investment account.

[`security_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities-security-id)

stringstring

A unique, Plaid-specific identifier for the security, used to associate securities with holdings. Like all Plaid identifiers, the `security_id` is case sensitive. The `security_id` may change if inherent details of the security change due to a corporate action, for example, in the event of a ticker symbol change or CUSIP change.

[`name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities-name)

nullablestringnullable, string

A descriptive name for the security, suitable for display.

[`isin`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities-isin)

nullablestringnullable, string

12-character ISIN, a globally unique securities identifier. A verified CUSIP Global Services license is required to receive this data. This field will be null by default for new customers, and null for existing customers starting March 12, 2024. If you would like access to this field, please start the verification process [here](https://docs.google.com/forms/d/e/1FAIpQLSd9asHEYEfmf8fxJTHZTAfAzW4dugsnSu-HS2J51f1mxwd6Sw/viewform).

[`cusip`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities-cusip)

nullablestringnullable, string

9-character CUSIP, an identifier assigned to North American securities. A verified CUSIP Global Services license is required to receive this data. This field will be null by default for new customers, and null for existing customers starting March 12, 2024. If you would like access to this field, please start the verification process [here](https://docs.google.com/forms/d/e/1FAIpQLSd9asHEYEfmf8fxJTHZTAfAzW4dugsnSu-HS2J51f1mxwd6Sw/viewform).

[`institution_security_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities-institution-security-id)

nullablestringnullable, string

An identifier given to the security by the institution.

[`institution_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities-institution-id)

nullablestringnullable, string

If `institution_security_id` is present, this field indicates the Plaid `institution_id` of the institution to whom the identifier belongs.

[`ticker_symbol`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities-ticker-symbol)

nullablestringnullable, string

The security’s trading symbol for publicly traded securities, and otherwise a short identifier if available.

[`type`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-securities-type)

nullablestringnullable, string

The security type of the holding. Valid security types are:  
`cash`: Cash, currency, and money market funds  
`cryptocurrency`: Digital or virtual currencies  
`derivative`: Options, warrants, and other derivative instruments  
`equity`: Domestic and foreign equities  
`etf`: Multi-asset exchange-traded investment funds  
`fixed income`: Bonds and certificates of deposit (CDs)  
`loan`: Loans and loan receivables  
`mutual fund`: Open- and closed-end vehicles pooling funds of multiple investors  
`other`: Unknown or other investment types

[`investment_transactions`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions)

[object][object]

Transaction history on the investment account.

[`investment_transaction_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-investment-transaction-id)

stringstring

The ID of the Investment transaction, unique across all Plaid transactions. Like all Plaid identifiers, the `investment_transaction_id` is case sensitive.

[`account_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-account-id)

stringstring

The `account_id` of the account against which this transaction posted.

[`security_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-security-id)

nullablestringnullable, string

The `security_id` to which this transaction is related.

[`date`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-date)

stringstring

The [ISO 8601](https://wikipedia.org/wiki/ISO_8601) posting date for the transaction.  
  

Format: `date`

[`name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-name)

stringstring

The institution’s description of the transaction.

[`quantity`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-quantity)

numbernumber

The number of units of the security involved in this transaction. Positive for buy transactions; negative for sell transactions.  
  

Format: `double`

[`amount`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-amount)

numbernumber

The complete value of the transaction. Positive values when cash is debited, e.g. purchases of stock; negative values when cash is credited, e.g. sales of stock. Treatment remains the same for cash-only movements unassociated with securities.  
  

Format: `double`

[`price`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-price)

numbernumber

The price of the security at which this transaction occurred.  
  

Format: `double`

[`fees`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-fees)

nullablenumbernullable, number

The combined value of all fees applied to this transaction  
  

Format: `double`

[`type`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-type)

stringstring

Value is one of the following:
`buy`: Buying an investment
`sell`: Selling an investment
`cancel`: A cancellation of a pending transaction
`cash`: Activity that modifies a cash position
`fee`: A fee on the account
`transfer`: Activity which modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer  
For descriptions of possible transaction types and subtypes, see the [Investment transaction types schema](https://plaid.com/docs/api/accounts/#investment-transaction-types-schema).  
  

Possible values: `buy`, `sell`, `cancel`, `cash`, `fee`, `transfer`

[`subtype`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-subtype)

stringstring

For descriptions of possible transaction types and subtypes, see the [Investment transaction types schema](https://plaid.com/docs/api/accounts/#investment-transaction-types-schema).  
  

Possible values: `account fee`, `adjustment`, `assignment`, `buy`, `buy to cover`, `contribution`, `deposit`, `distribution`, `dividend`, `dividend reinvestment`, `exercise`, `expire`, `fund fee`, `interest`, `interest receivable`, `interest reinvestment`, `legal fee`, `loan payment`, `long-term capital gain`, `long-term capital gain reinvestment`, `management fee`, `margin expense`, `merger`, `miscellaneous fee`, `non-qualified dividend`, `non-resident tax`, `pending credit`, `pending debit`, `qualified dividend`, `rebalance`, `return of principal`, `request`, `sell`, `sell short`, `send`, `short-term capital gain`, `short-term capital gain reinvestment`, `spin off`, `split`, `stock distribution`, `tax`, `tax withheld`, `trade`, `transfer`, `transfer fee`, `trust fee`, `unqualified gain`, `withdrawal`

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the transaction. Always `null` if `unofficial_currency_code` is non-`null`.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-accounts-investments-investment-transactions-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the holding. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`institution_name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-institution-name)

stringstring

The full financial institution name associated with the Item.

[`institution_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-institution-id)

stringstring

The id of the financial institution associated with the Item.

[`item_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`last_update_time`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-items-last-update-time)

stringstring

The date and time when this Item’s data was last retrieved from the financial institution, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`attributes`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes)

objectobject

Attributes for the VOA report.

[`total_inflow_amount`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes-total-inflow-amount)

nullableobjectnullable, object

Total amount of debit transactions into the report's accounts in the time period of the report. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes-total-inflow-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes-total-inflow-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes-total-inflow-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes-total-outflow-amount)

nullableobjectnullable, object

Total amount of credit transactions into the report's accounts in the time period of the report. This field only takes into account USD transactions from the accounts.

[`amount`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes-total-outflow-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes-total-outflow-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-check_report-verification-get-response-report-voa-attributes-total-outflow-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`employment_refresh`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh)

nullableobjectnullable, object

An object representing an Employment Refresh Report.

[`generated_time`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-generated-time)

stringstring

The date and time when the Employment Refresh Report was created, in ISO 8601 format (e.g. "2018-04-12T03:32:11Z").  
  

Format: `date-time`

[`days_requested`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-days-requested)

numbernumber

The number of days of transaction history that the Employment Refresh Report covers.

[`items`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items)

[object][object]

Data returned by Plaid about each of the Items included in the Employment Refresh Report.

[`accounts`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts)

[object][object]

Data about each of the accounts open on the Item.

[`account_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself.

[`official_name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-official-name)

nullablestringnullable, string

The official name of the account as given by the financial institution.

[`type`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`transactions`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-transactions)

[object][object]

Transaction history associated with the account for the Employment Refresh Report. Note that this transaction differs from a Base Report transaction in that it will only be deposits, and the amounts will be omitted.

[`account_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-transactions-account-id)

stringstring

The ID of the account in which this transaction occurred.

[`original_description`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-transactions-original-description)

stringstring

The string returned by the financial institution to describe the transaction.

[`date`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-transactions-date)

stringstring

For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an ISO 8601 format ( `YYYY-MM-DD` ).  
  

Format: `date`

[`pending`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-transactions-pending)

booleanboolean

When `true`, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled.

[`transaction_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-accounts-transactions-transaction-id)

stringstring

The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.

[`institution_name`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-institution-name)

stringstring

The full financial institution name associated with the Item.

[`institution_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-institution-id)

stringstring

The id of the financial institution associated with the Item.

[`item_id`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`last_update_time`](/docs/api/products/check/#cra-check_report-verification-get-response-report-employment-refresh-items-last-update-time)

stringstring

The date and time when this Item’s data was last retrieved from the financial institution, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`request_id`](/docs/api/products/check/#cra-check_report-verification-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`warnings`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings)

[object][object]

If the home lending report generation was successful but a subset of data could not be retrieved, this array will contain information about the errors causing information to be missing.

[`warning_type`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-warning-type)

stringstring

The warning type, which will always be `CHECK_REPORT_WARNING`

[`warning_code`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Account-owner information is not available.
`TRANSACTIONS_UNAVAILABLE`: Transactions information associated with Credit and Depository accounts are unavailable.
`USER_FRAUD_ALERT`: The user has placed a fraud alert on their Plaid Check consumer report due to suspected fraud. Please note that when a fraud alert is in place, the recipient of the consumer report has an obligation to verify the consumer’s identity.  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `USER_FRAUD_ALERT`

[`cause`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause)

nullableobjectnullable, object

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`item_id`](/docs/api/products/check/#cra-check_report-verification-get-response-warnings-cause-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm",
  "report": {
    "report_id": "028e8404-a013-4a45-ac9e-002482f9cafc",
    "client_report_id": "client_report_id_1221",
    "voa": {
      "generated_time": "2023-03-30T18:27:37Z",
      "days_requested": 90,
      "attributes": {
        "total_inflow_amount": {
          "amount": -345.12,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        },
        "total_outflow_amount": {
          "amount": 235.12,
          "iso_currency_code": "USD",
          "unofficial_currency_code": null
        }
      },
      "items": [
        {
          "accounts": [
            {
              "account_id": "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
              "balances": {
                "available": 100,
                "current": 110,
                "iso_currency_code": "USD",
                "unofficial_currency_code": null,
                "historical_balances": [
                  {
                    "current": 110,
                    "date": "2023-03-29",
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  },
                  {
                    "current": 125.55,
                    "date": "2023-03-28",
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  },
                  {
                    "current": 80.13,
                    "date": "2023-03-27",
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  },
                  {
                    "current": 246.11,
                    "date": "2023-03-26",
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  },
                  {
                    "current": 182.71,
                    "date": "2023-03-25",
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null
                  }
                ],
                "average_balance_30_days": 200,
                "average_balance_60_days": 150,
                "average_balance_90_days": 125,
                "nsf_overdraft_transactions_count": 0
              },
              "consumer_disputes": [],
              "mask": "0000",
              "name": "Plaid Checking",
              "official_name": "Plaid Gold Standard 0% Interest Checking",
              "type": "depository",
              "subtype": "checking",
              "days_available": 90,
              "transactions_insights": {
                "all_transactions": [
                  {
                    "account_id": "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
                    "amount": 89.4,
                    "date": "2023-03-27",
                    "iso_currency_code": "USD",
                    "original_description": "SparkFun",
                    "pending": false,
                    "transaction_id": "4zBRq1Qem4uAPnoyKjJNTRQpQddM4ztlo1PLD",
                    "unofficial_currency_code": null
                  },
                  {
                    "account_id": "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
                    "amount": 12,
                    "date": "2023-03-28",
                    "iso_currency_code": "USD",
                    "original_description": "McDonalds #3322",
                    "pending": false,
                    "transaction_id": "dkjL41PnbKsPral79jpxhMWdW55gkPfBkWpRL",
                    "unofficial_currency_code": null
                  },
                  {
                    "account_id": "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
                    "amount": 4.33,
                    "date": "2023-03-28",
                    "iso_currency_code": "USD",
                    "original_description": "Starbucks",
                    "pending": false,
                    "transaction_id": "a84ZxQaWDAtDL3dRgmazT57K7jjN3WFkNWMDy",
                    "unofficial_currency_code": null
                  },
                  {
                    "account_id": "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
                    "amount": -500,
                    "date": "2023-03-29",
                    "iso_currency_code": "USD",
                    "original_description": "United Airlines **** REFUND ****",
                    "pending": false,
                    "transaction_id": "xG9jbv3eMoFWepzB7wQLT3LoLggX5Duy1Gbe5",
                    "unofficial_currency_code": null
                  }
                ],
                "end_date": "2024-07-31",
                "start_date": "2024-07-01"
              },
              "owners": [
                {
                  "addresses": [
                    {
                      "data": {
                        "city": "Malakoff",
                        "country": "US",
                        "region": "NY",
                        "street": "2992 Cameron Road",
                        "postal_code": "14236"
                      },
                      "primary": true
                    },
                    {
                      "data": {
                        "city": "San Matias",
                        "country": "US",
                        "region": "CA",
                        "street": "2493 Leisure Lane",
                        "postal_code": "93405-2255"
                      },
                      "primary": false
                    }
                  ],
                  "emails": [
                    {
                      "data": "accountholder0@example.com",
                      "primary": true,
                      "type": "primary"
                    },
                    {
                      "data": "accountholder1@example.com",
                      "primary": false,
                      "type": "secondary"
                    },
                    {
                      "data": "extraordinarily.long.email.username.123456@reallylonghostname.com",
                      "primary": false,
                      "type": "other"
                    }
                  ],
                  "names": [
                    "Alberta Bobbeth Charleson"
                  ],
                  "phone_numbers": [
                    {
                      "data": "+1 111-555-3333",
                      "primary": false,
                      "type": "home"
                    },
                    {
                      "data": "+1 111-555-4444",
                      "primary": false,
                      "type": "work"
                    },
                    {
                      "data": "+1 111-555-5555",
                      "primary": false,
                      "type": "mobile"
                    }
                  ]
                }
              ],
              "ownership_type": null
            }
          ],
          "institution_name": "First Platypus Bank",
          "institution_id": "ins_109508",
          "item_id": "AZMP7JrGXgtPd3AQMeg7hwMKgk5E8qU1V5ME7",
          "last_update_time": "2023-03-30T18:25:26Z"
        }
      ]
    },
    "employment_refresh": {
      "generated_time": "2023-03-30T18:27:37Z",
      "days_requested": 60,
      "items": [
        {
          "accounts": [
            {
              "account_id": "1qKRXQjk8xUWDJojNwPXTj8gEmR48piqRNye8",
              "name": "Plaid Money Market",
              "official_name": "Plaid Platinum Standard 1.85% Interest Money Market",
              "type": "depository",
              "subtype": "money market",
              "transactions": [
                {
                  "account_id": "1qKRXQjk8xUWDJojNwPXTj8gEmR48piqRNye8",
                  "original_description": "ACH Electronic CreditGUSTO PAY 123456",
                  "date": "2023-03-30",
                  "pending": false,
                  "transaction_id": "gGQgjoeyqBF89PND6K14Sow1wddZBmtLomJ78"
                }
              ]
            },
            {
              "account_id": "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
              "name": "Plaid Checking",
              "official_name": "Plaid Gold Standard 0% Interest Checking",
              "type": "depository",
              "subtype": "checking",
              "transactions": [
                {
                  "account_id": "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
                  "original_description": "United Airlines **** REFUND ****",
                  "date": "2023-03-29",
                  "pending": false,
                  "transaction_id": "xG9jbv3eMoFWepzB7wQLT3LoLggX5Duy1Gbe5"
                }
              ]
            }
          ],
          "institution_name": "First Platypus Bank",
          "institution_id": "ins_109508",
          "item_id": "AZMP7JrGXgtPd3AQMeg7hwMKgk5E8qU1V5ME7",
          "last_update_time": "2023-03-30T18:25:26Z"
        }
      ]
    }
  },
  "warnings": []
}
```

=\*=\*=\*=

#### `/cra/check_report/verification/pdf/get`

#### Retrieve Consumer Reports as a Verification PDF

The [`/cra/check_report/verification/pdf/get`](/docs/api/products/check/#cracheck_reportverificationpdfget) endpoint retrieves the most recent Consumer Report in PDF format, specifically formatted for Home Lending verification use cases. Before calling this endpoint, ensure that you've created a VOA report through Link or the [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) endpoint, and have received a `CHECK_REPORT_READY` or a `USER_CHECK_REPORT_READY` webhook.

The response to [`/cra/check_report/verification/pdf/get`](/docs/api/products/check/#cracheck_reportverificationpdfget) is the PDF binary data. The `request_id` is returned in the `Plaid-Request-ID` header.

/cra/check\_report/verification/pdf/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-verification-pdf-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-verification-pdf-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-verification-pdf-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`report_requested`](/docs/api/products/check/#cra-check_report-verification-pdf-get-request-report-requested)

requiredstringrequired, string

The type of verification PDF report to fetch.  
  

Possible values: `voa`, `employment_refresh`

[`user_token`](/docs/api/products/check/#cra-check_report-verification-pdf-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/cra/check\_report/verification/pdf/get

```
try {
  const response = await client.craCheckReportVerificationPDFGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
    report_requested: 'VOA',
  });
  const pdf = response.buffer.toString('base64');
} catch (error) {
  // handle error
}
```

##### Response

This endpoint returns binary PDF data. View a sample [Home Lending Report (aka VoA Report)](https://plaid.com/documents/sample-mortgage-verification-voa.pdf) or [Employment Refresh](https://plaid.com/documents/sample-mortgage-verification-voe.pdf) report.

=\*=\*=\*=

#### `/cra/check_report/create`

#### Refresh or create a Consumer Report

Use [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) to refresh data in an existing report. A Consumer Report will last for 24 hours before expiring; you should call any `/get` endpoints on the report before it expires. If a report expires, you can call [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate) again to re-generate it and refresh the data in the report.

/cra/check\_report/create

**Request fields**

[`client_id`](/docs/api/products/check/#cra-check_report-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-check_report-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-check_report-create-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`user_token`](/docs/api/products/check/#cra-check_report-create-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`webhook`](/docs/api/products/check/#cra-check_report-create-request-webhook)

requiredstringrequired, string

The destination URL to which webhooks will be sent  
  

Format: `url`

[`days_requested`](/docs/api/products/check/#cra-check_report-create-request-days-requested)

requiredintegerrequired, integer

The number of days of data to request for the report. Default value is 365; maximum is 731; minimum is 180. If a value lower than 180 is provided, a minimum of 180 days of history will be requested.  
  

Maximum: `731`

[`days_required`](/docs/api/products/check/#cra-check_report-create-request-days-required)

integerinteger

The minimum number of days of data required for the report to be successfully generated.  
  

Maximum: `184`

[`client_report_id`](/docs/api/products/check/#cra-check_report-create-request-client-report-id)

stringstring

Client-generated identifier, which can be used by lenders to track loan applications.

[`products`](/docs/api/products/check/#cra-check_report-create-request-products)

[string][string]

Specifies a list of products that will be eagerly generated when creating the report (in addition to the Base Report, which is always eagerly generated). These products will be made available before a success webhook is sent. Use this option to minimize response latency for product `/get` endpoints. Note that specifying `cra_partner_insights` in this field will trigger a billable event. Other products are not billed until the respective reports are fetched via product-specific `/get` endpoints.  
  

Min items: `1`

Possible values: `cra_income_insights`, `cra_cashflow_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_lend_score`

[`base_report`](/docs/api/products/check/#cra-check_report-create-request-base-report)

objectobject

Defines configuration options to generate a Base Report

[`client_report_id`](/docs/api/products/check/#cra-check_report-create-request-base-report-client-report-id)

deprecatedstringdeprecated, string

Client-generated identifier, which can be used by lenders to track loan applications. This field is deprecated. Use the `client_report_id` field at the top level of the request instead.

[`gse_options`](/docs/api/products/check/#cra-check_report-create-request-base-report-gse-options)

objectobject

Specifies options for creating reports that can be shared with GSEs for mortgage verification.

[`report_types`](/docs/api/products/check/#cra-check_report-create-request-base-report-gse-options-report-types)

required[string]required, [string]

Specifies which types of reports should be made available to GSEs.  
  

Possible values: `VOA`, `EMPLOYMENT_REFRESH`

[`require_identity`](/docs/api/products/check/#cra-check_report-create-request-base-report-require-identity)

booleanboolean

Indicates that the report must include identity information. If identity information is not available, the report will fail.

[`cashflow_insights`](/docs/api/products/check/#cra-check_report-create-request-cashflow-insights)

objectobject

Defines configuration options to generate Cashflow Insights

[`attributes_version`](/docs/api/products/check/#cra-check_report-create-request-cashflow-insights-attributes-version)

stringstring

The version of cashflow attributes  
  

Possible values: `CFI1`

[`partner_insights`](/docs/api/products/check/#cra-check_report-create-request-partner-insights)

objectobject

Defines configuration to generate Partner Insights.

[`prism_products`](/docs/api/products/check/#cra-check_report-create-request-partner-insights-prism-products)

deprecated[string]deprecated, [string]

The specific Prism Data products to return. If none are passed in, then all products will be returned.  
  

Possible values: `insights`, `scores`

[`prism_versions`](/docs/api/products/check/#cra-check_report-create-request-partner-insights-prism-versions)

objectobject

The versions of Prism products to evaluate

[`firstdetect`](/docs/api/products/check/#cra-check_report-create-request-partner-insights-prism-versions-firstdetect)

stringstring

The version of Prism FirstDetect. If not specified, will default to v3.  
  

Possible values: `3`, `null`

[`detect`](/docs/api/products/check/#cra-check_report-create-request-partner-insights-prism-versions-detect)

stringstring

The version of Prism Detect  
  

Possible values: `4`, `null`

[`cashscore`](/docs/api/products/check/#cra-check_report-create-request-partner-insights-prism-versions-cashscore)

stringstring

The version of Prism CashScore. If not specified, will default to v3.  
  

Possible values: `4`, `3`, `null`

[`extend`](/docs/api/products/check/#cra-check_report-create-request-partner-insights-prism-versions-extend)

stringstring

The version of Prism Extend  
  

Possible values: `4`, `null`

[`insights`](/docs/api/products/check/#cra-check_report-create-request-partner-insights-prism-versions-insights)

stringstring

The version of Prism Insights. If not specified, will default to v3.  
  

Possible values: `4`, `3`, `null`

[`lend_score`](/docs/api/products/check/#cra-check_report-create-request-lend-score)

objectobject

Defines configuration options to generate the LendScore

[`lend_score_version`](/docs/api/products/check/#cra-check_report-create-request-lend-score-lend-score-version)

stringstring

The version of the LendScore  
  

Possible values: `LS1`

[`network_insights`](/docs/api/products/check/#cra-check_report-create-request-network-insights)

objectobject

Defines configuration options to generate Network Insights

[`network_insights_version`](/docs/api/products/check/#cra-check_report-create-request-network-insights-network-insights-version)

stringstring

The version of network insights  
  

Possible values: `NI1`

[`include_investments`](/docs/api/products/check/#cra-check_report-create-request-include-investments)

booleanboolean

Indicates that investment data should be extracted from the linked account(s).

[`consumer_report_permissible_purpose`](/docs/api/products/check/#cra-check_report-create-request-consumer-report-permissible-purpose)

requiredstringrequired, string

Describes the reason you are generating a Consumer Report for this user. When calling `/link/token/create`, this field is required when using Plaid Check (CRA) products; invalid if not using Plaid Check (CRA) products.  
`ACCOUNT_REVIEW_CREDIT`: In connection with a consumer credit transaction for the review or collection of an account pursuant to FCRA Section 604(a)(3)(A).  
`ACCOUNT_REVIEW_NON_CREDIT`: For a legitimate business need of the information to review a non-credit account provided primarily for personal, family, or household purposes to determine whether the consumer continues to meet the terms of the account pursuant to FCRA Section 604(a)(3)(F)(2).  
`EXTENSION_OF_CREDIT`: In connection with a credit transaction initiated by and involving the consumer pursuant to FCRA Section 604(a)(3)(A).  
`LEGITIMATE_BUSINESS_NEED_TENANT_SCREENING`: For a legitimate business need in connection with a business transaction initiated by the consumer primarily for personal, family, or household purposes in connection with a property rental assessment pursuant to FCRA Section 604(a)(3)(F)(i).  
`LEGITIMATE_BUSINESS_NEED_OTHER`: For a legitimate business need in connection with a business transaction made primarily for personal, family, or household initiated by the consumer pursuant to FCRA Section 604(a)(3)(F)(i).  
`WRITTEN_INSTRUCTION_PREQUALIFICATION`: In accordance with the written instructions of the consumer pursuant to FCRA Section 604(a)(2), to evaluate an application’s profile to make an offer to the consumer.  
`WRITTEN_INSTRUCTION_OTHER`: In accordance with the written instructions of the consumer pursuant to FCRA Section 604(a)(2), such as when an individual agrees to act as a guarantor or assumes personal liability for a consumer, business, or commercial loan.  
`ELIGIBILITY_FOR_GOVT_BENEFITS`: In connection with an eligibility determination for a government benefit where the entity is required to consider an applicant’s financial status pursuant to FCRA Section 604(a)(3)(D).  
  

Possible values: `ACCOUNT_REVIEW_CREDIT`, `ACCOUNT_REVIEW_NON_CREDIT`, `EXTENSION_OF_CREDIT`, `LEGITIMATE_BUSINESS_NEED_TENANT_SCREENING`, `LEGITIMATE_BUSINESS_NEED_OTHER`, `WRITTEN_INSTRUCTION_PREQUALIFICATION`, `WRITTEN_INSTRUCTION_OTHER`, `ELIGIBILITY_FOR_GOVT_BENEFITS`

/cra/check\_report/create

```
try {
  const response = await client.craCheckReportCreate({
    user_id: 'usr_9nSp2KuZ2x4JDw',
    webhook: 'https://sample-web-hook.com',
    days_requested: 365,
    consumer_report_permissible_purpose: 'LEGITIMATE_BUSINESS_NEED_OTHER',
  });
} catch (error) {
  // handle error
}
```

/cra/check\_report/create

**Response fields**

[`request_id`](/docs/api/products/check/#cra-check_report-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm"
}
```

=\*=\*=\*=

#### `/cra/monitoring_insights/get`

#### Retrieve a Monitoring Insights Report

This endpoint allows you to retrieve a Cash Flow Updates report by passing in the `user_id` referred to in the webhook you received.

/cra/monitoring\_insights/get

**Request fields**

[`client_id`](/docs/api/products/check/#cra-monitoring_insights-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-monitoring_insights-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-monitoring_insights-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`consumer_report_permissible_purpose`](/docs/api/products/check/#cra-monitoring_insights-get-request-consumer-report-permissible-purpose)

requiredstringrequired, string

Describes the reason you are generating a Consumer Report for this user.  
`ACCOUNT_REVIEW_CREDIT`: In connection with a consumer credit transaction for the review or collection of an account pursuant to FCRA Section 604(a)(3)(A).  
`WRITTEN_INSTRUCTION_OTHER`: In accordance with the written instructions of the consumer pursuant to FCRA Section 604(a)(2), such as when an individual agrees to act as a guarantor or assumes personal liability for a consumer, business, or commercial loan.  
  

Possible values: `ACCOUNT_REVIEW_CREDIT`, `WRITTEN_INSTRUCTION_OTHER`

[`user_token`](/docs/api/products/check/#cra-monitoring_insights-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/cra/monitoring\_insights/get

```
try {
  const response = await client.craMonitoringInsightsGet({
    user_id: 'usr_9nSp2KuZ2x4JDw',
    consumer_report_permissible_purpose: 'EXTENSION_OF_CREDIT',
  });
} catch (error) {
  // handle error
}
```

/cra/monitoring\_insights/get

**Response fields**

[`request_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`user_insights_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-user-insights-id)

stringstring

A unique ID identifying a User Monitoring Insights Report. Like all Plaid identifiers, this ID is case sensitive.

[`items`](/docs/api/products/check/#cra-monitoring_insights-get-response-items)

[object][object]

An array of Monitoring Insights Items associated with the user.

[`date_generated`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-date-generated)

stringstring

The date and time when the specific insights were generated (per-item), in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (e.g. "2018-04-12T03:32:11Z").  
  

Format: `date-time`

[`item_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-item-id)

stringstring

The `item_id` of the Item associated with the insights

[`institution_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-institution-id)

stringstring

The id of the financial institution associated with the Item.

[`institution_name`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-institution-name)

stringstring

The full financial institution name associated with the Item.

[`status`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-status)

objectobject

An object with details of the Monitoring Insights Item's status.

[`status_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-status-status-code)

stringstring

Enum for the status of the Item's insights  
  

Possible values: `AVAILABLE`, `FAILED`, `PENDING`

[`reason`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-status-reason)

nullablestringnullable, string

A reason for why a Monitoring Insights Report is not available.
This field will only be populated when the `status_code` is not `AVAILABLE`

[`insights`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights)

nullableobjectnullable, object

An object representing the Monitoring Insights for the given Item

[`income`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income)

objectobject

An object representing the income subcategory of the report

[`total_monthly_income`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-total-monthly-income)

objectobject

Details about about the total monthly income

[`current_amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-total-monthly-income-current-amount)

numbernumber

The aggregated income of the last 30 days

[`income_sources_counts`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-income-sources-counts)

objectobject

Details about the number of income sources

[`current_count`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-income-sources-counts-current-count)

numbernumber

The number of income sources currently detected

[`forecasted_monthly_income`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-forecasted-monthly-income)

objectobject

An object representing the predicted average monthly net income amount. This amount reflects the funds deposited into the account and may not include any withheld income such as taxes or other payroll deductions

[`current_amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-forecasted-monthly-income-current-amount)

numbernumber

The current forecasted monthly income

[`historical_annual_income`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-historical-annual-income)

objectobject

An object representing the historical annual income amount.

[`current_amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-historical-annual-income-current-amount)

numbernumber

The current historical annual income

[`income_sources`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-income-sources)

[object][object]

The income sources for this Item. Each entry in the array is a single income source

[`income_source_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-income-sources-income-source-id)

stringstring

A unique identifier for an income source

[`income_description`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-income-sources-income-description)

stringstring

The most common name or original description for the underlying income transactions

[`income_category`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-income-sources-income-category)

stringstring

The income category.
`BANK_INTEREST`: Interest earned from a bank account.
`BENEFIT_OTHER`: Government benefits other than retirement, unemployment, child support, or disability. Currently used only in the UK, to represent benefits such as Cost of Living Payments.
`CASH`: Deprecated and used only for existing legacy implementations. Has been replaced by `CASH_DEPOSIT` and `TRANSFER_FROM_APPLICATION`.
`CASH_DEPOSIT`: A cash or check deposit.
`CHILD_SUPPORT`: Child support payments received.
`GIG_ECONOMY`: Income earned as a gig economy worker, e.g. driving for Uber, Lyft, Postmates, DoorDash, etc.
`LONG_TERM_DISABILITY`: Disability payments, including Social Security disability benefits.
`OTHER`: Income that could not be categorized as any other income category.
`MILITARY`: Veterans benefits. Income earned as salary for serving in the military (e.g. through DFAS) will be classified as `SALARY` rather than `MILITARY`.
`RENTAL`: Income earned from a rental property. Income may be identified as rental when the payment is received through a rental platform, e.g. Airbnb; rent paid directly by the tenant to the property owner (e.g. via cash, check, or ACH) will typically not be classified as rental income.
`RETIREMENT`: Payments from private retirement systems, pensions, and government retirement programs, including Social Security retirement benefits.
`SALARY`: Payment from an employer to an earner or other form of permanent employment.
`TAX_REFUND`: A tax refund.
`TRANSFER_FROM_APPLICATION`: Deposits from a money transfer app, such as Venmo, Cash App, or Zelle.
`UNEMPLOYMENT`: Unemployment benefits. In the UK, includes certain low-income benefits such as the Universal Credit.  
  

Possible values: `SALARY`, `UNEMPLOYMENT`, `CASH`, `GIG_ECONOMY`, `RENTAL`, `CHILD_SUPPORT`, `MILITARY`, `RETIREMENT`, `LONG_TERM_DISABILITY`, `BANK_INTEREST`, `CASH_DEPOSIT`, `TRANSFER_FROM_APPLICATION`, `TAX_REFUND`, `BENEFIT_OTHER`, `OTHER`

[`last_transaction_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-income-income-sources-last-transaction-date)

stringstring

The last detected transaction date for this income source  
  

Format: `date`

[`loans`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-loans)

objectobject

An object representing the loan exposure subcategory of the report

[`loan_payments_counts`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-loans-loan-payments-counts)

objectobject

Details regarding the number of loan payments

[`current_count`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-loans-loan-payments-counts-current-count)

numbernumber

The current number of loan payments made in the last 30 days

[`loan_disbursements_count`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-loans-loan-disbursements-count)

numbernumber

The number of loan disbursements detected in the last 30 days

[`loan_payment_merchants_counts`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-loans-loan-payment-merchants-counts)

objectobject

Details regarding the number of unique loan payment merchants

[`current_count`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-insights-loans-loan-payment-merchants-counts-current-count)

numbernumber

The current number of unique loan payment merchants detected in the last 30 days

[`accounts`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts)

[object][object]

Data about each of the accounts open on the Item.

[`account_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`balances`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances)

objectobject

Information about an account's balances.

[`available`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-available)

nullablenumbernullable, number

The amount of funds available to be withdrawn from the account, as determined by the financial institution.  
For `credit`-type accounts, the `available` balance typically equals the `limit` less the `current` balance, less any pending outflows plus any pending inflows.  
For `depository`-type accounts, the `available` balance typically equals the `current` balance less any pending outflows plus any pending inflows. For `depository`-type accounts, the `available` balance does not include the overdraft limit.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the `available` balance is the total cash available to withdraw as presented by the institution.  
Note that not all institutions calculate the `available` balance. In the event that `available` balance is unavailable, Plaid will return an `available` balance value of `null`.  
Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by `/accounts/balance/get`.  
If `current` is `null` this field is guaranteed not to be `null`.  
  

Format: `double`

[`current`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-current)

nullablenumbernullable, number

The total amount of funds in or owed by the account.  
For `credit`-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  
For `loan`-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (`ins_116944`). For Sallie Mae student loans, the account's balance includes both principal and any outstanding interest.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the current balance is the total value of assets as presented by the institution.  
Note that balance information may be cached unless the value was returned by `/accounts/balance/get`; if the Item is enabled for Transactions, the balance will be at least as recent as the most recent Transaction update. If you require realtime balance information, use the `available` balance as provided by `/accounts/balance/get`.  
When returned by `/accounts/balance/get`, this field may be `null`. When this happens, `available` is guaranteed not to be `null`.  
  

Format: `double`

[`limit`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-limit)

nullablenumbernullable, number

For `credit`-type accounts, this represents the credit limit.  
For `depository`-type accounts, this represents the pre-arranged overdraft limit, which is common for current (checking) accounts in Europe.  
In North America, this field is typically only available for `credit`-type accounts.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the balance. Always null if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the balance. Always null if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`last_updated_datetime`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-last-updated-datetime)

nullablestringnullable, string

Timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the oldest acceptable balance when making a request to `/accounts/balance/get`.  
This field is only used and expected when the institution is `ins_128026` (Capital One) and the Item contains one or more accounts with a non-depository account type, in which case a value must be provided or an `INVALID_REQUEST` error with the code of `INVALID_FIELD` will be returned. For Capital One depository accounts as well as all other account types on all other institutions, this field is ignored. See [account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full list of account types.  
If the balance that is pulled is older than the given timestamp for Items with this field required, an `INVALID_REQUEST` error with the code of `LAST_UPDATED_DATETIME_OUT_OF_RANGE` will be returned with the most recent timestamp for the requested account contained in the response.  
  

Format: `date-time`

[`average_balance`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-average-balance)

nullablenumbernullable, number

The average historical balance for the entire report  
  

Format: `double`

[`average_monthly_balances`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-average-monthly-balances)

[object][object]

The average historical balance of each calendar month

[`start_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-average-monthly-balances-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).

[`end_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-average-monthly-balances-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).

[`average_balance`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-average-monthly-balances-average-balance)

objectobject

This contains an amount, denominated in the currency specified by either `iso_currency_code` or `unofficial_currency_code`

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-average-monthly-balances-average-balance-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-average-monthly-balances-average-balance-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-average-monthly-balances-average-balance-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`most_recent_thirty_day_average_balance`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-balances-most-recent-thirty-day-average-balance)

nullablenumbernullable, number

The average historical balance from the most recent 30 days  
  

Format: `double`

[`consumer_disputes`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-consumer-disputes)

[object][object]

The information about previously submitted valid dispute statements by the consumer

[`consumer_dispute_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-consumer-disputes-consumer-dispute-id)

deprecatedstringdeprecated, string

(Deprecated) A unique identifier (UUID) of the consumer dispute that can be used for troubleshooting

[`dispute_field_create_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-consumer-disputes-dispute-field-create-date)

stringstring

Date of the disputed field (e.g. transaction date), in an ISO 8601 format (YYYY-MM-DD)  
  

Format: `date`

[`category`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-consumer-disputes-category)

stringstring

Type of data being disputed by the consumer  
  

Possible values: `TRANSACTION`, `BALANCE`, `IDENTITY`, `OTHER`

[`statement`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-consumer-disputes-statement)

stringstring

Text content of dispute

[`mask`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.

[`metadata`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-metadata)

objectobject

Metadata about the extracted account.

[`start_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-metadata-start-date)

nullablestringnullable, string

The beginning of the range of the financial institution provided data for the account, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-metadata-end-date)

nullablestringnullable, string

The end of the range of the financial institution provided data for the account, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`name`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself

[`official_name`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-official-name)

nullablestringnullable, string

The official name of the account as given by the financial institution

[`type`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`days_available`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-days-available)

numbernumber

The duration of transaction history available within this report for this Item, typically defined as the time since the date of the earliest transaction in that account.

[`transactions`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions)

[object][object]

Transaction history associated with the account. Transaction history returned by endpoints such as `/transactions/get` or `/investments/transactions/get` will be returned in the top-level `transactions` field instead. Some transactions may have their details masked in accordance to the FCRA. These will appear with a `credit_category` of `MASKED_TRANSACTION_CATEGORY`.

[`account_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-account-id)

stringstring

The ID of the account in which this transaction occurred.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-amount)

numbernumber

The settled value of the transaction, denominated in the transaction's currency, as stated in `iso_currency_code` or `unofficial_currency_code`. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the transaction. Always `null` if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the transaction. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`original_description`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-original-description)

nullablestringnullable, string

The string returned by the financial institution to describe the transaction.

[`credit_category`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-credit-category)

nullableobjectnullable, object

Information describing the intent of the transaction. Most relevant for credit use cases, but not limited to such use cases.  
See the [`taxonomy csv file`](https://plaid.com/documents/credit-category-taxonomy.csv) for a full list of credit categories.

[`primary`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-credit-category-primary)

stringstring

A high level category that communicates the broad category of the transaction.

[`detailed`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-credit-category-detailed)

stringstring

A granular category conveying the transaction's intent. This field can also be used as a unique identifier for the category.

[`check_number`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-check-number)

nullablestringnullable, string

The check number of the transaction. This field is only populated for check transactions.

[`date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-date)

stringstring

For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DD` ).  
  

Format: `date`

[`date_transacted`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-date-transacted)

nullablestringnullable, string

The date on which the transaction took place, in IS0 8601 format.

[`location`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location)

objectobject

A representation of where a transaction took place. Location data is provided only for transactions at physical locations, not for online transactions. Location data availability depends primarily on the merchant and is most likely to be populated for transactions at large retail chains; small, local businesses are less likely to have location data available.

[`address`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location-address)

nullablestringnullable, string

The street address where the transaction occurred.

[`city`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location-city)

nullablestringnullable, string

The city where the transaction occurred.

[`region`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location-region)

nullablestringnullable, string

The region or state where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called `state`.

[`postal_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location-postal-code)

nullablestringnullable, string

The postal code where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code where the transaction occurred.

[`lat`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location-lat)

nullablenumbernullable, number

The latitude where the transaction occurred.  
  

Format: `double`

[`lon`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location-lon)

nullablenumbernullable, number

The longitude where the transaction occurred.  
  

Format: `double`

[`store_number`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-location-store-number)

nullablestringnullable, string

The merchant defined store number where the transaction occurred.

[`merchant_name`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-merchant-name)

nullablestringnullable, string

The merchant name, as enriched by Plaid from the `name` field. This is typically a more human-readable version of the merchant counterparty in the transaction. For some bank transactions (such as checks or account transfers) where there is no meaningful merchant name, this value will be `null`.

[`pending`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-pending)

booleanboolean

When `true`, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled.

[`account_owner`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-account-owner)

nullablestringnullable, string

The name of the account owner. This field is not typically populated and only relevant when dealing with sub-accounts.

[`transaction_id`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-transactions-transaction-id)

stringstring

The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.

[`owners`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners)

[object][object]

Data returned by the financial institution about the account owner or owners. For business accounts, the name reported may be either the name of the individual or the name of the business, depending on the institution. Multiple owners on a single account will be represented in the same `owner` object, not in multiple owner objects within the array. This array can also be empty if no owners are found.

[`names`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-names)

[string][string]

A list of names associated with the account by the financial institution. In the case of a joint account, Plaid will make a best effort to report the names of all account holders.  
If an Item contains multiple accounts with different owner names, some institutions will report all names associated with the Item in each account's `names` array.

[`phone_numbers`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-phone-numbers)

[object][object]

A list of phone numbers associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-phone-numbers-data)

stringstring

The phone number.

[`primary`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-phone-numbers-primary)

booleanboolean

When `true`, identifies the phone number as the primary number on an account.

[`type`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-phone-numbers-type)

stringstring

The type of phone number.  
  

Possible values: `home`, `work`, `office`, `mobile`, `mobile1`, `other`

[`emails`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-emails)

[object][object]

A list of email addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-emails-data)

stringstring

The email address.

[`primary`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-emails-primary)

booleanboolean

When `true`, identifies the email address as the primary email on an account.

[`type`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-emails-type)

stringstring

The type of email account as described by the financial institution.  
  

Possible values: `primary`, `secondary`, `other`

[`addresses`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-addresses)

[object][object]

Data about the various addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-addresses-data)

objectobject

Data about the components comprising an address.

[`city`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-addresses-data-city)

nullablestringnullable, string

The full city name

[`region`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-addresses-data-region)

nullablestringnullable, string

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-addresses-data-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-addresses-data-postal-code)

nullablestringnullable, string

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-addresses-data-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code

[`primary`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-owners-addresses-primary)

booleanboolean

When `true`, identifies the address as the primary address on an account.

[`ownership_type`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-ownership-type)

nullablestringnullable, string

How an asset is owned.  
`association`: Ownership by a corporation, partnership, or unincorporated association, including for-profit and not-for-profit organizations.
`individual`: Ownership by an individual.
`joint`: Joint ownership by multiple parties.
`trust`: Ownership by a revocable or irrevocable trust.  
  

Possible values: `null`, `individual`, `joint`, `association`, `trust`

[`account_insights`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights)

deprecatedobjectdeprecated, object

Calculated insights derived from transaction-level data. This field has been deprecated in favor of [Base Report attributes aggregated across accounts](https://plaid.com/docs/api/products/check/#cra-check_report-base_report-get-response-report-attributes) and will be removed in a future release.

[`oldest_transaction_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-oldest-transaction-date)

nullablestringnullable, string

Date of the earliest transaction for the account.  
  

Format: `date`

[`most_recent_transaction_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-most-recent-transaction-date)

nullablestringnullable, string

Date of the most recent transaction for the account.  
  

Format: `date`

[`days_available`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-days-available)

integerinteger

Number of days days available for the account.

[`average_days_between_transactions`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-days-between-transactions)

numbernumber

Average number of days between sequential transactions

[`longest_gaps_between_transactions`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-longest-gaps-between-transactions)

[object][object]

Longest gap between sequential transactions in a time period. This array can include multiple time periods.

[`start_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-longest-gaps-between-transactions-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-longest-gaps-between-transactions-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`days`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-longest-gaps-between-transactions-days)

nullableintegernullable, integer

Largest number of days between sequential transactions for this time period.

[`number_of_inflows`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-inflows)

[object][object]

The number of debits into the account. This array will be empty for non-depository accounts.

[`start_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-inflows-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-inflows-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`count`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-inflows-count)

integerinteger

The number of credits or debits out of the account for this time period.

[`average_inflow_amounts`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-inflow-amounts)

[object][object]

Average amount of debit transactions into the account in a time period. This array will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`start_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-inflow-amounts-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-inflow-amounts-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`total_amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-inflow-amounts-total-amount)

objectobject

This contains an amount, denominated in the currency specified by either `iso_currency_code` or `unofficial_currency_code`

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-inflow-amounts-total-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-inflow-amounts-total-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-inflow-amounts-total-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`number_of_outflows`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-outflows)

[object][object]

The number of outflows from the account. This array will be empty for non-depository accounts.

[`start_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-outflows-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-outflows-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`count`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-outflows-count)

integerinteger

The number of credits or debits out of the account for this time period.

[`average_outflow_amounts`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-outflow-amounts)

[object][object]

Average amount of transactions out of the account in a time period. This array will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`start_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-outflow-amounts-start-date)

stringstring

The start date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-outflow-amounts-end-date)

stringstring

The end date of this time period.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`total_amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-outflow-amounts-total-amount)

objectobject

This contains an amount, denominated in the currency specified by either `iso_currency_code` or `unofficial_currency_code`

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-outflow-amounts-total-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-outflow-amounts-total-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-average-outflow-amounts-total-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`number_of_days_no_transactions`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-account-insights-number-of-days-no-transactions)

integerinteger

Number of days with no transactions

[`attributes`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes)

objectobject

Calculated attributes derived from transaction-level data.

[`is_primary_account`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-is-primary-account)

nullablebooleannullable, boolean

Prediction indicator of whether the account is a primary account. Only one account per account type across the items connected will have a value of true.

[`primary_account_score`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-primary-account-score)

nullablenumbernullable, number

Value ranging from 0-1. The higher the score, the more confident we are of the account being the primary account.

[`nsf_overdraft_transactions_count`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-nsf-overdraft-transactions-count)

integerinteger

The number of net NSF fee transactions for a given account within the report time range (not counting any fees that were reversed within the time range).

[`nsf_overdraft_transactions_count_30d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-nsf-overdraft-transactions-count-30d)

integerinteger

The number of net NSF fee transactions within the last 30 days for a given account (not counting any fees that were reversed within the time range).

[`nsf_overdraft_transactions_count_60d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-nsf-overdraft-transactions-count-60d)

integerinteger

The number of net NSF fee transactions within the last 60 days for a given account (not counting any fees that were reversed within the time range).

[`nsf_overdraft_transactions_count_90d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-nsf-overdraft-transactions-count-90d)

integerinteger

The number of net NSF fee transactions within the last 90 days for a given account (not counting any fees that were reversed within the time range).

[`total_inflow_amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount)

nullableobjectnullable, object

Total amount of debit transactions into the account in the time period of the report. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_30d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-30d)

nullableobjectnullable, object

Total amount of debit transactions into the account in the last 30 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-30d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-30d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-30d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_60d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-60d)

nullableobjectnullable, object

Total amount of debit transactions into the account in the last 60 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-60d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-60d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-60d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_inflow_amount_90d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-90d)

nullableobjectnullable, object

Total amount of debit transactions into the account in the last 90 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-90d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-90d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-inflow-amount-90d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount)

nullableobjectnullable, object

Total amount of credit transactions into the account in the time period of the report. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_30d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-30d)

nullableobjectnullable, object

Total amount of credit transactions into the account in the last 30 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-30d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-30d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-30d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_60d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-60d)

nullableobjectnullable, object

Total amount of credit transactions into the account in the last 60 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-60d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-60d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-60d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`total_outflow_amount_90d`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-90d)

nullableobjectnullable, object

Total amount of credit transactions into the account in the last 90 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.

[`amount`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-90d-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-90d-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/check/#cra-monitoring_insights-get-response-items-accounts-attributes-total-outflow-amount-90d-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

Response Object

```
{
  "user_insights_id": "028e8404-a013-4a45-ac9e-002482f9cafc",
  "items": [
    {
      "date_generated": "2023-03-30T18:27:37Z",
      "item_id": "AZMP7JrGXgtPd3AQMeg7hwMKgk5E8qU1V5ME7",
      "institution_id": "ins_0",
      "institution_name": "Plaid Bank",
      "status": {
        "status_code": "AVAILABLE"
      },
      "insights": {
        "income": {
          "income_sources": [
            {
              "income_source_id": "f17efbdd-caab-4278-8ece-963511cd3d51",
              "income_description": "PLAID_INC_DIRECT_DEP_PPD",
              "income_category": "SALARY",
              "last_transaction_date": "2023-03-30"
            }
          ],
          "forecasted_monthly_income": {
            "current_amount": 12000
          },
          "total_monthly_income": {
            "current_amount": 20000.31
          },
          "historical_annual_income": {
            "current_amount": 144000
          },
          "income_sources_counts": {
            "current_count": 1
          }
        },
        "loans": {
          "loan_payments_counts": {
            "current_count": 1
          },
          "loan_payment_merchants_counts": {
            "current_count": 1
          },
          "loan_disbursements_count": 1
        }
      },
      "accounts": [
        {
          "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
          "attributes": {
            "total_inflow_amount": {
              "amount": -2500,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            },
            "total_inflow_amount_30d": {
              "amount": -1000,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            },
            "total_inflow_amount_60d": {
              "amount": -2500,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            },
            "total_inflow_amount_90d": {
              "amount": -2500,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            },
            "total_outflow_amount": {
              "amount": 2500,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            },
            "total_outflow_amount_30d": {
              "amount": 1000,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            },
            "total_outflow_amount_60d": {
              "amount": 2500,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            },
            "total_outflow_amount_90d": {
              "amount": 2500,
              "iso_currency_code": "USD",
              "unofficial_currency_code": null
            }
          },
          "balances": {
            "available": 5000,
            "average_balance": 4956.12,
            "average_monthly_balances": [
              {
                "average_balance": {
                  "amount": 4956.12,
                  "iso_currency_code": "USD",
                  "unofficial_currency_code": null
                },
                "end_date": "2024-07-31",
                "start_date": "2024-07-01"
              }
            ],
            "current": 5000,
            "iso_currency_code": "USD",
            "limit": null,
            "most_recent_thirty_day_average_balance": 4956.125,
            "unofficial_currency_code": null
          },
          "consumer_disputes": [],
          "days_available": 365,
          "mask": "1208",
          "metadata": {
            "start_date": "2024-01-01",
            "end_date": "2024-07-16"
          },
          "name": "Checking",
          "official_name": "Plaid checking",
          "owners": [
            {
              "addresses": [
                {
                  "data": {
                    "city": "Malakoff",
                    "country": "US",
                    "postal_code": "14236",
                    "region": "NY",
                    "street": "2992 Cameron Road"
                  },
                  "primary": true
                },
                {
                  "data": {
                    "city": "San Matias",
                    "country": "US",
                    "postal_code": "93405-2255",
                    "region": "CA",
                    "street": "2493 Leisure Lane"
                  },
                  "primary": false
                }
              ],
              "emails": [
                {
                  "data": "accountholder0@example.com",
                  "primary": true,
                  "type": "primary"
                },
                {
                  "data": "accountholder1@example.com",
                  "primary": false,
                  "type": "secondary"
                },
                {
                  "data": "extraordinarily.long.email.username.123456@reallylonghostname.com",
                  "primary": false,
                  "type": "other"
                }
              ],
              "names": [
                "Alberta Bobbeth Charleson"
              ],
              "phone_numbers": [
                {
                  "data": "+1 111-555-3333",
                  "primary": false,
                  "type": "home"
                },
                {
                  "data": "+1 111-555-4444",
                  "primary": false,
                  "type": "work"
                },
                {
                  "data": "+1 111-555-5555",
                  "primary": false,
                  "type": "mobile"
                }
              ]
            }
          ],
          "ownership_type": null,
          "subtype": "checking",
          "transactions": [
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 37.07,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_ONLINE_MARKETPLACES",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-12",
              "date_posted": "2024-07-12T00:00:00Z",
              "date_transacted": "2024-07-12",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Amazon",
              "original_description": "AMZN Mktp US*11111111 Amzn.com/bill WA AM",
              "pending": false,
              "transaction_id": "XA7ZLy8rXzt7D3j9B6LMIgv5VxyQkAhbKjzmp",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 51.61,
              "check_number": null,
              "credit_category": {
                "detailed": "DINING_DINING",
                "primary": "DINING"
              },
              "date": "2024-07-12",
              "date_posted": "2024-07-12T00:00:00Z",
              "date_transacted": "2024-07-12",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Domino's",
              "original_description": "DOMINO's XXXX 111-222-3333",
              "pending": false,
              "transaction_id": "VEPeMbWqRluPVZLQX4MDUkKRw41Ljzf9gyLBW",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 7.55,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_FURNITURE_AND_HARDWARE",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-12",
              "date_posted": "2024-07-12T00:00:00Z",
              "date_transacted": "2024-07-12",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Chicago",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "IKEA",
              "original_description": "IKEA CHICAGO",
              "pending": false,
              "transaction_id": "6GQZARgvroCAE1eW5wpQT7w3oB6nvzi8DKMBa",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 12.87,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_SPORTING_GOODS",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-12",
              "date_posted": "2024-07-12T00:00:00Z",
              "date_transacted": "2024-07-12",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Redlands",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": "CA",
                "state": "CA",
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Nike",
              "original_description": "NIKE REDLANDS CA",
              "pending": false,
              "transaction_id": "DkbmlP8BZxibzADqNplKTeL8aZJVQ1c3WR95z",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 44.21,
              "check_number": null,
              "credit_category": {
                "detailed": "DINING_DINING",
                "primary": "DINING"
              },
              "date": "2024-07-12",
              "date_posted": "2024-07-12T00:00:00Z",
              "date_transacted": "2024-07-12",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": null,
              "original_description": "POKE BROS * POKE BRO IL",
              "pending": false,
              "transaction_id": "RpdN7W8GmRSdjZB9Jm7ATj4M86vdnktapkrgL",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 36.82,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_DISCOUNT_STORES",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-13",
              "date_posted": "2024-07-13T00:00:00Z",
              "date_transacted": "2024-07-13",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Family Dollar",
              "original_description": "FAMILY DOLLAR",
              "pending": false,
              "transaction_id": "5AeQWvo5KLtAD9wNL68PTdAgPE7VNWf5Kye1G",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 13.27,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-13",
              "date_posted": "2024-07-13T00:00:00Z",
              "date_transacted": "2024-07-13",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Instacart",
              "original_description": "INSTACART HTTPSINSTACAR CA",
              "pending": false,
              "transaction_id": "Jjlr3MEVg1HlKbdkZj39ij5a7eg9MqtB6MWDo",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 36.03,
              "check_number": null,
              "credit_category": {
                "detailed": "DINING_DINING",
                "primary": "DINING"
              },
              "date": "2024-07-13",
              "date_posted": "2024-07-13T00:00:00Z",
              "date_transacted": "2024-07-13",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": null,
              "original_description": "POKE BROS * POKE BRO IL",
              "pending": false,
              "transaction_id": "kN9KV7yAZJUMPn93KDXqsG9MrpjlyLUL6Dgl8",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 54.74,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-13",
              "date_posted": "2024-07-13T00:00:00Z",
              "date_transacted": "2024-07-13",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Whittier",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": "CA",
                "state": "CA",
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Smart & Final",
              "original_description": "POS SMART AND FINAL 111 WHITTIER CA",
              "pending": false,
              "transaction_id": "lPvrweZAMqHDar43vwWKs547kLZVEzfpogGVJ",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 37.5,
              "check_number": null,
              "credit_category": {
                "detailed": "DINING_DINING",
                "primary": "DINING"
              },
              "date": "2024-07-13",
              "date_posted": "2024-07-13T00:00:00Z",
              "date_transacted": "2024-07-13",
              "iso_currency_code": "USD",
              "location": {
                "address": "1627 N 24th St",
                "city": "Phoenix",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": "85008",
                "region": "AZ",
                "state": "AZ",
                "store_number": null,
                "zip": "85008"
              },
              "merchant_name": "Taqueria El Guerrerense",
              "original_description": "TAQUERIA EL GUERRERO PHOENIX AZ",
              "pending": false,
              "transaction_id": "wka74WKqngiyJ3pj7dl5SbpLGQBZqyCPZRDbP",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 41.42,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_ONLINE_MARKETPLACES",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-14",
              "date_posted": "2024-07-14T00:00:00Z",
              "date_transacted": "2024-07-14",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Amazon",
              "original_description": "AMZN Mktp US*11111111 Amzn.com/bill WA AM",
              "pending": false,
              "transaction_id": "BBGnV4RkerHjn8WVavGyiJbQ95VNDaC4M56bJ",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": -1077.93,
              "check_number": null,
              "credit_category": {
                "detailed": "INCOME_OTHER",
                "primary": "INCOME"
              },
              "date": "2024-07-14",
              "date_posted": "2024-07-14T00:00:00Z",
              "date_transacted": "2024-07-14",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Lyft",
              "original_description": "LYFT TRANSFER",
              "pending": false,
              "transaction_id": "3Ej78yKJlQu1Abw7xzo4U4JR6pmwzntZlbKDK",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 47.17,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-14",
              "date_posted": "2024-07-14T00:00:00Z",
              "date_transacted": "2024-07-14",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Whittier",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": "CA",
                "state": "CA",
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Smart & Final",
              "original_description": "POS SMART AND FINAL 111 WHITTIER CA",
              "pending": false,
              "transaction_id": "rMzaBpJw8jSZRJQBabKdteQBwd5EaWc7J9qem",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 12.37,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-14",
              "date_posted": "2024-07-14T00:00:00Z",
              "date_transacted": "2024-07-14",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Whittier",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": "CA",
                "state": "CA",
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Smart & Final",
              "original_description": "POS SMART AND FINAL 111 WHITTIER CA",
              "pending": false,
              "transaction_id": "zWPZjkmzynTyel89ZjExS59DV6WAaZflNBJ56",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 44.18,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-14",
              "date_posted": "2024-07-14T00:00:00Z",
              "date_transacted": "2024-07-14",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Portland",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": "OR",
                "state": "OR",
                "store_number": "1111",
                "zip": null
              },
              "merchant_name": "Safeway",
              "original_description": "SAFEWAY #1111 PORTLAND OR            111111",
              "pending": false,
              "transaction_id": "K7qzx1nP8ptqgwaRMbxyI86XrqADMluRpkWx5",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 45.37,
              "check_number": null,
              "credit_category": {
                "detailed": "DINING_DINING",
                "primary": "DINING"
              },
              "date": "2024-07-14",
              "date_posted": "2024-07-14T00:00:00Z",
              "date_transacted": "2024-07-14",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Uber Eats",
              "original_description": "UBER EATS",
              "pending": false,
              "transaction_id": "qZrdzLRAgNHo5peMdD9xIzELl3a1NvcgrPAzL",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 15.22,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_ONLINE_MARKETPLACES",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-15",
              "date_posted": "2024-07-15T00:00:00Z",
              "date_transacted": "2024-07-15",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Amazon",
              "original_description": "AMZN Mktp US*11111111 Amzn.com/bill WA AM",
              "pending": false,
              "transaction_id": "NZzx4oRPkAHzyRekpG4PTZkWnBPqEyiy6pB1M",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 26.33,
              "check_number": null,
              "credit_category": {
                "detailed": "DINING_DINING",
                "primary": "DINING"
              },
              "date": "2024-07-15",
              "date_posted": "2024-07-15T00:00:00Z",
              "date_transacted": "2024-07-15",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Domino's",
              "original_description": "DOMINO's XXXX 111-222-3333",
              "pending": false,
              "transaction_id": "x84eNArKbESz8Woden6LT3nvyogeJXc64Pp35",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 39.8,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_DISCOUNT_STORES",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-15",
              "date_posted": "2024-07-15T00:00:00Z",
              "date_transacted": "2024-07-15",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Family Dollar",
              "original_description": "FAMILY DOLLAR",
              "pending": false,
              "transaction_id": "dzWnyxwZ4GHlZPGgrNyxiMG7qd5jDgCJEz5jL",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 45.06,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-15",
              "date_posted": "2024-07-15T00:00:00Z",
              "date_transacted": "2024-07-15",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Instacart",
              "original_description": "INSTACART HTTPSINSTACAR CA",
              "pending": false,
              "transaction_id": "4W7eE9rZqMToDArbPeLNIREoKpdgBMcJbVNQD",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 34.91,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-15",
              "date_posted": "2024-07-15T00:00:00Z",
              "date_transacted": "2024-07-15",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Whittier",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": "CA",
                "state": "CA",
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Smart & Final",
              "original_description": "POS SMART AND FINAL 111 WHITTIER CA",
              "pending": false,
              "transaction_id": "j4yqDjb7QwS7woGzqrgDIEG1NaQVZwf6Wmz3D",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 49.78,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-15",
              "date_posted": "2024-07-15T00:00:00Z",
              "date_transacted": "2024-07-15",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Portland",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": "OR",
                "state": "OR",
                "store_number": "1111",
                "zip": null
              },
              "merchant_name": "Safeway",
              "original_description": "SAFEWAY #1111 PORTLAND OR            111111",
              "pending": false,
              "transaction_id": "aqgWnze7xoHd6DQwLPnzT5dgPKjB1NfZ5JlBy",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 54.24,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-15",
              "date_posted": "2024-07-15T00:00:00Z",
              "date_transacted": "2024-07-15",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": "Portland",
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": "OR",
                "state": "OR",
                "store_number": "1111",
                "zip": null
              },
              "merchant_name": "Safeway",
              "original_description": "SAFEWAY #1111 PORTLAND OR            111111",
              "pending": false,
              "transaction_id": "P13aP8b7nmS3WQoxg1PMsdvMK679RNfo65B4G",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 41.79,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_ONLINE_MARKETPLACES",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-16",
              "date_posted": "2024-07-16T00:00:00Z",
              "date_transacted": "2024-07-16",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Amazon",
              "original_description": "AMZN Mktp US*11111111 Amzn.com/bill WA AM",
              "pending": false,
              "transaction_id": "7nZMG6pXz8SADylMqzx7TraE4qjJm7udJyAGm",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 33.86,
              "check_number": null,
              "credit_category": {
                "detailed": "FOOD_RETAIL_GROCERIES",
                "primary": "FOOD_RETAIL"
              },
              "date": "2024-07-16",
              "date_posted": "2024-07-16T00:00:00Z",
              "date_transacted": "2024-07-16",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "Instacart",
              "original_description": "INSTACART HTTPSINSTACAR CA",
              "pending": false,
              "transaction_id": "MQr3ap7PWEIrQG7bLdaNsxyBV7g1KqCL6pwoy",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 27.08,
              "check_number": null,
              "credit_category": {
                "detailed": "DINING_DINING",
                "primary": "DINING"
              },
              "date": "2024-07-16",
              "date_posted": "2024-07-16T00:00:00Z",
              "date_transacted": "2024-07-16",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": null,
              "original_description": "POKE BROS * POKE BRO IL",
              "pending": false,
              "transaction_id": "eBAk9dvwNbHPZpr8W69dU3rekJz47Kcr9BRwl",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 25.94,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_FURNITURE_AND_HARDWARE",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-16",
              "date_posted": "2024-07-16T00:00:00Z",
              "date_transacted": "2024-07-16",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": "The Home Depot",
              "original_description": "THE HOME DEPOT",
              "pending": false,
              "transaction_id": "QLx4jEJZb9SxRm7aWbjAio3LrgZ5vPswm64dE",
              "unofficial_currency_code": null
            },
            {
              "account_id": "NZzx4oRPkAHzyRekpG4PTZkoGpNAR4uypaj1E",
              "account_owner": null,
              "amount": 27.57,
              "check_number": null,
              "credit_category": {
                "detailed": "GENERAL_MERCHANDISE_OTHER_GENERAL_MERCHANDISE",
                "primary": "GENERAL_MERCHANDISE"
              },
              "date": "2024-07-16",
              "date_posted": "2024-07-16T00:00:00Z",
              "date_transacted": "2024-07-16",
              "iso_currency_code": "USD",
              "location": {
                "address": null,
                "city": null,
                "country": null,
                "lat": null,
                "lon": null,
                "postal_code": null,
                "region": null,
                "state": null,
                "store_number": null,
                "zip": null
              },
              "merchant_name": null,
              "original_description": "The Press Club",
              "pending": false,
              "transaction_id": "ZnQ1ovqBldSQ6GzRbroAHLdQP68BrKceqmAjX",
              "unofficial_currency_code": null
            }
          ],
          "type": "depository"
        }
      ]
    }
  ],
  "request_id": "m8MDnv9okwxFNBV"
}
```

=\*=\*=\*=

#### `/cra/monitoring_insights/subscribe`

#### Subscribe to Monitoring Insights

This endpoint allows you to subscribe to insights for a user's linked CRA items, which are updated between one and four times per day (best-effort).

/cra/monitoring\_insights/subscribe

**Request fields**

[`client_id`](/docs/api/products/check/#cra-monitoring_insights-subscribe-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-monitoring_insights-subscribe-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/products/check/#cra-monitoring_insights-subscribe-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`item_id`](/docs/api/products/check/#cra-monitoring_insights-subscribe-request-item-id)

stringstring

The item ID to subscribe for Cash Flow Updates.

[`webhook`](/docs/api/products/check/#cra-monitoring_insights-subscribe-request-webhook)

requiredstringrequired, string

URL to which Plaid will send Cash Flow Updates webhooks, for example when the requested Cash Flow Updates report is ready.  
  

Format: `url`

[`income_categories`](/docs/api/products/check/#cra-monitoring_insights-subscribe-request-income-categories)

[string][string]

Income categories to include in Cash Flow Updates. If empty or `null`, this field will default to including all possible categories.  
  

Possible values: `SALARY`, `UNEMPLOYMENT`, `CASH`, `GIG_ECONOMY`, `RENTAL`, `CHILD_SUPPORT`, `MILITARY`, `RETIREMENT`, `LONG_TERM_DISABILITY`, `BANK_INTEREST`, `CASH_DEPOSIT`, `TRANSFER_FROM_APPLICATION`, `TAX_REFUND`, `BENEFIT_OTHER`, `OTHER`

[`user_token`](/docs/api/products/check/#cra-monitoring_insights-subscribe-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/cra/monitoring\_insights/subscribe

```
try {
  const response = await client.craMonitoringInsightsSubscribe({
    user_id: 'usr_9nSp2KuZ2x4JDw',
    item_id: 'eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6',
    webhook: 'https://example.com/webhook',
    income_categories: [CreditBankIncomeCategory.Salary],
  });
} catch (error) {
  // handle error
}
```

/cra/monitoring\_insights/subscribe

**Response fields**

[`request_id`](/docs/api/products/check/#cra-monitoring_insights-subscribe-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`subscription_id`](/docs/api/products/check/#cra-monitoring_insights-subscribe-response-subscription-id)

stringstring

A unique identifier for the subscription.

Response Object

```
{
  "subscription_id": "f17efbdd-caab-4278-8ece-963511cd3d51",
  "request_id": "GVzMdiDd8DDAQK4"
}
```

=\*=\*=\*=

#### `/cra/monitoring_insights/unsubscribe`

#### Unsubscribe from Monitoring Insights

This endpoint allows you to unsubscribe from previously subscribed Monitoring Insights.

/cra/monitoring\_insights/unsubscribe

**Request fields**

[`client_id`](/docs/api/products/check/#cra-monitoring_insights-unsubscribe-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/check/#cra-monitoring_insights-unsubscribe-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`subscription_id`](/docs/api/products/check/#cra-monitoring_insights-unsubscribe-request-subscription-id)

requiredstringrequired, string

A unique identifier for the subscription.

/cra/monitoring\_insights/unsubscribe

```
try {
  const response = await client.craMonitoringInsightsUnsubscribe({
    subscription_id: '"f17efbdd-caab-4278-8ece-963511cd3d51"',
  });
} catch (error) {
  // handle error
}
```

/cra/monitoring\_insights/unsubscribe

**Response fields**

[`request_id`](/docs/api/products/check/#cra-monitoring_insights-unsubscribe-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "GVzMdiDd8DDAQK4"
}
```

### Webhooks

When you create a new report, either by creating a Link token with a Plaid Check product, or by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate), Plaid Check will start generating a report for you. When the report has been created (or the report creation fails), Plaid Check will let you know by sending you either a `CHECK_REPORT: USER_CHECK_REPORT_READY` or `CHECK_REPORT: USER_CHECK_REPORT_FAILED` webhook.

Customers who first called [`/user/create`](/docs/api/users/#usercreate) after December 10, 2025 will receive the `USER_CHECK_REPORT_READY` / `USER_CHECK_REPORT_FAILED` webhooks. Customers who integrated before this date will receive the older `CHECK_REPORT_READY` / `CHECK_REPORT_FAILED` webhooks. For more details, see [new User APIs](/docs/api/users/user-apis/).

=\*=\*=\*=

#### `USER_CHECK_REPORT_READY`

Fired when the Check Report are ready to be retrieved. Once this webhook has fired, the report will be available to retrieve for 24 hours.

**Properties**

[`webhook_type`](/docs/api/products/check/#CraUserCheckReportReadyWebhook-webhook-type)

stringstring

`CHECK_REPORT`

[`webhook_code`](/docs/api/products/check/#CraUserCheckReportReadyWebhook-webhook-code)

stringstring

`USER_CHECK_REPORT_READY`

[`user_id`](/docs/api/products/check/#CraUserCheckReportReadyWebhook-user-id)

stringstring

The `user_id` associated with the user whose data is being requested. This is received by calling `user/create`.

[`successful_products`](/docs/api/products/check/#CraUserCheckReportReadyWebhook-successful-products)

[string][string]

Specifies a list of products that have successfully been generated for the report.  
  

Possible values: `cra_base_report`, `cra_income_insights`, `cra_cashflow_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_monitoring`, `cra_lend_score`

[`failed_products`](/docs/api/products/check/#CraUserCheckReportReadyWebhook-failed-products)

[string][string]

Specifies a list of products that have failed to generate for the report. Additional detail on what caused the failure can be found by calling the product /get endpoint.  
  

Possible values: `cra_base_report`, `cra_income_insights`, `cra_cashflow_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_monitoring`, `cra_lend_score`

[`environment`](/docs/api/products/check/#CraUserCheckReportReadyWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CHECK_REPORT",
  "webhook_code": "USER_CHECK_REPORT_READY",
  "user_id": "usr_8c3ZbDBYjaqUXZ",
  "successful_products": [
    "cra_base_report"
  ],
  "environment": "production"
}
```

=\*=\*=\*=

#### `USER_CHECK_REPORT_FAILED`

Fired when a Check Report has failed to generate

**Properties**

[`webhook_type`](/docs/api/products/check/#CraUserCheckReportFailedWebhook-webhook-type)

stringstring

`CHECK_REPORT`

[`webhook_code`](/docs/api/products/check/#CraUserCheckReportFailedWebhook-webhook-code)

stringstring

`USER_CHECK_REPORT_FAILED`

[`user_id`](/docs/api/products/check/#CraUserCheckReportFailedWebhook-user-id)

stringstring

The `user_id` associated with the user whose data is being requested. This is received by calling user/create.

[`environment`](/docs/api/products/check/#CraUserCheckReportFailedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CHECK_REPORT",
  "webhook_code": "USER_CHECK_REPORT_FAILED",
  "user_id": "usr_8c3ZbDBYjaqUXZ",
  "environment": "production"
}
```

=\*=\*=\*=

#### `CHECK_REPORT_READY`

Fired when the Check Report are ready to be retrieved. Once this webhook has fired, the report will be available to retrieve for 24 hours.

**Properties**

[`webhook_type`](/docs/api/products/check/#CraCheckReportReadyWebhook-webhook-type)

stringstring

`CHECK_REPORT`

[`webhook_code`](/docs/api/products/check/#CraCheckReportReadyWebhook-webhook-code)

stringstring

`CHECK_REPORT_READY`

[`user_id`](/docs/api/products/check/#CraCheckReportReadyWebhook-user-id)

stringstring

The `user_id` corresponding to the user the webhook has fired for.

[`successful_products`](/docs/api/products/check/#CraCheckReportReadyWebhook-successful-products)

[string][string]

Specifies a list of products that have successfully been generated for the report.  
  

Possible values: `cra_base_report`, `cra_income_insights`, `cra_cashflow_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_monitoring`, `cra_lend_score`

[`failed_products`](/docs/api/products/check/#CraCheckReportReadyWebhook-failed-products)

[string][string]

Specifies a list of products that have failed to generate for the report. Additional detail on what caused the failure can be found by calling the product /get endpoint.  
  

Possible values: `cra_base_report`, `cra_income_insights`, `cra_cashflow_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_monitoring`, `cra_lend_score`

[`environment`](/docs/api/products/check/#CraCheckReportReadyWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CHECK_REPORT",
  "webhook_code": "CHECK_REPORT_READY",
  "user_id": "wz666MBjYWTp2PDzzggYhM6oWWmBb",
  "successful_products": [
    "cra_base_report"
  ],
  "environment": "production"
}
```

=\*=\*=\*=

#### `CHECK_REPORT_FAILED`

Fired when a Check Report has failed to generate. To get more details, call [`/user/items/get`](/docs/api/users/#useritemsget) and check for non-null `error` objects on the associated Items in the response. These `error` objects will contain more details on why the Item is in an error state and how to resolve it. After resolving the errors, you can try to re-generate the report.

**Properties**

[`webhook_type`](/docs/api/products/check/#CraCheckReportFailedWebhook-webhook-type)

stringstring

`CHECK_REPORT`

[`webhook_code`](/docs/api/products/check/#CraCheckReportFailedWebhook-webhook-code)

stringstring

`CHECK_REPORT_FAILED`

[`user_id`](/docs/api/products/check/#CraCheckReportFailedWebhook-user-id)

stringstring

The `user_id` corresponding to the user the webhook has fired for.

[`environment`](/docs/api/products/check/#CraCheckReportFailedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CHECK_REPORT",
  "webhook_code": "CHECK_REPORT_FAILED",
  "user_id": "wz666MBjYWTp2PDzzggYhM6oWWmBb",
  "environment": "production"
}
```

### Cash flow updates webhooks

These webhooks are specific to the Cash Flow Updates (beta) product.

All webhooks in this section except for `CASH_FLOW_INSIGHTS_UPDATED` are legacy webhooks and will only be fired for customers who integrated with Plaid Check before December 10, 2025. For newer integrations, `CASH_FLOW_INSIGHTS_UPDATED` has replaced the other webhooks. For more details, see [New user APIs](/docs/api/users/user-apis/).

=\*=\*=\*=

#### `CASH_FLOW_INSIGHTS_UPDATED`

For each item on an enabled user, this webhook will fire up to four times a day with status information. This webhook will not fire immediately upon enrollment in Cash Flow Updates. The payload may contain an `insights` array with insights that have been detected, if any (e.g. `LOW_BALANCE_DETECTED`, `LARGE_DEPOSIT_DETECTED`). Upon receiving the webhook, call [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) to retrieve the updated insights.

**Properties**

[`webhook_type`](/docs/api/products/check/#CashFlowUpdatesInsightsV2Webhook-webhook-type)

stringstring

`CASH_FLOW_UPDATES`

[`webhook_code`](/docs/api/products/check/#CashFlowUpdatesInsightsV2Webhook-webhook-code)

stringstring

`CASH_FLOW_INSIGHTS_UPDATED`

[`status`](/docs/api/products/check/#CashFlowUpdatesInsightsV2Webhook-status)

stringstring

Enum for the status of the insights  
  

Possible values: `AVAILABLE`, `FAILED`

[`user_id`](/docs/api/products/check/#CashFlowUpdatesInsightsV2Webhook-user-id)

stringstring

The `user_id` associated with the user whose data is being requested. This is received by calling `user/create`.

[`insights`](/docs/api/products/check/#CashFlowUpdatesInsightsV2Webhook-insights)

[string][string]

Array containing the insights detected within the generated report, if any. Possible values include:
`LARGE_DEPOSIT_DETECTED`: signaling a deposit over $5,000
`LOW_BALANCE_DETECTED`: signaling a balance below $100
`NEW_LOAN_PAYMENT_DETECTED`: signaling a new loan payment
`NSF_OVERDRAFT_DETECTED`: signaling an NSF overdraft  
  

Possible values: `LARGE_DEPOSIT_DETECTED`, `LOW_BALANCE_DETECTED`, `NEW_LOAN_PAYMENT_DETECTED`, `NSF_OVERDRAFT_DETECTED`

[`environment`](/docs/api/products/check/#CashFlowUpdatesInsightsV2Webhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CASH_FLOW_UPDATES",
  "webhook_code": "CASH_FLOW_INSIGHTS_UPDATED",
  "status": "AVAILABLE",
  "user_id": "usr_6009db6e",
  "insights": [
    "LARGE_DEPOSIT_DETECTED",
    "LOW_BALANCE_DETECTED",
    "NEW_LOAN_PAYMENT_DETECTED",
    "NSF_OVERDRAFT_DETECTED"
  ],
  "environment": "sandbox"
}
```

=\*=\*=\*=

#### `INSIGHTS_UPDATED`

For each user's Item enabled for Cash Flow Updates, this webhook will fire between one and four times a day with information on the status of the update. This webhook will not fire immediately upon enrollment in Cash Flow Updates. Upon receiving the webhook, call [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) to retrieve the updated insights. At approximately the same time as the `INSIGHTS_UPDATED` webhook, any event-driven `CASH_FLOW_UPDATES` webhooks (e.g. `LOW_BALANCE_DETECTED`, `LARGE_DEPOSIT_DETECTED`) that were triggered by the update will also fire. This webhook has been replaced by the `CASH_FLOW_INSIGHTS_UPDATED` webhook for all customers who began using Plaid Check on or after December 10, 2025.

**Properties**

[`webhook_type`](/docs/api/products/check/#CashFlowUpdatesInsightsWebhook-webhook-type)

stringstring

`CASH_FLOW_UPDATES`

[`webhook_code`](/docs/api/products/check/#CashFlowUpdatesInsightsWebhook-webhook-code)

stringstring

`INSIGHTS_UPDATED`

[`status`](/docs/api/products/check/#CashFlowUpdatesInsightsWebhook-status)

stringstring

Enum for the status of the insights  
  

Possible values: `AVAILABLE`, `FAILED`

[`user_id`](/docs/api/products/check/#CashFlowUpdatesInsightsWebhook-user-id)

stringstring

The `user_id` that the report is associated with

[`environment`](/docs/api/products/check/#CashFlowUpdatesInsightsWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CASH_FLOW_UPDATES",
  "webhook_code": "INSIGHTS_UPDATED",
  "status": "AVAILABLE",
  "user_id": "9eaba3c2fdc916bc197f279185b986607dd21682a5b04eab04a5a03e8b3f3334",
  "environment": "production"
}
```

=\*=\*=\*=

#### `LARGE_DEPOSIT_DETECTED`

For each user's item enabled for Cash Flow Updates, this webhook will fire when an update detects a deposit over $5,000. Upon receiving the webhook, call [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) to retrieve the updated insights.

**Properties**

[`webhook_type`](/docs/api/products/check/#CashFlowUpdatesLargeDepositWebhook-webhook-type)

stringstring

`CASH_FLOW_UPDATES`

[`webhook_code`](/docs/api/products/check/#CashFlowUpdatesLargeDepositWebhook-webhook-code)

stringstring

`LARGE_DEPOSIT_DETECTED`

[`status`](/docs/api/products/check/#CashFlowUpdatesLargeDepositWebhook-status)

stringstring

Enum for the status of the insights  
  

Possible values: `AVAILABLE`, `FAILED`

[`user_id`](/docs/api/products/check/#CashFlowUpdatesLargeDepositWebhook-user-id)

stringstring

The `user_id` that the report is associated with

[`environment`](/docs/api/products/check/#CashFlowUpdatesLargeDepositWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CASH_FLOW_UPDATES",
  "webhook_code": "LARGE_DEPOSIT_DETECTED",
  "status": "AVAILABLE",
  "user_id": "9eaba3c2fdc916bc197f279185b986607dd21682a5b04eab04a5a03e8b3f3334",
  "environment": "production"
}
```

=\*=\*=\*=

#### `LOW_BALANCE_DETECTED`

For each user's item enabled for Cash Flow Updates, this webhook will fire when an update detects a balance below $100. Upon receiving the webhook, call [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) to retrieve the updated insights.

**Properties**

[`webhook_type`](/docs/api/products/check/#CashFlowUpdatesLowBalanceWebhook-webhook-type)

stringstring

`CASH_FLOW_UPDATES`

[`webhook_code`](/docs/api/products/check/#CashFlowUpdatesLowBalanceWebhook-webhook-code)

stringstring

`LOW_BALANCE_DETECTED`

[`status`](/docs/api/products/check/#CashFlowUpdatesLowBalanceWebhook-status)

stringstring

Enum for the status of the insights  
  

Possible values: `AVAILABLE`, `FAILED`

[`user_id`](/docs/api/products/check/#CashFlowUpdatesLowBalanceWebhook-user-id)

stringstring

The `user_id` that the report is associated with

[`environment`](/docs/api/products/check/#CashFlowUpdatesLowBalanceWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CASH_FLOW_UPDATES",
  "webhook_code": "LOW_BALANCE_DETECTED",
  "status": "AVAILABLE",
  "user_id": "9eaba3c2fdc916bc197f279185b986607dd21682a5b04eab04a5a03e8b3f3334",
  "environment": "production"
}
```

=\*=\*=\*=

#### `NEW_LOAN_PAYMENT_DETECTED`

For each user's item enabled for Cash Flow Updates, this webhook will fire when an update detects a new loan payment. Upon receiving the webhook, call [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) to retrieve the updated insights.

**Properties**

[`webhook_type`](/docs/api/products/check/#CashFlowUpdatesNewLoanPaymentWebhook-webhook-type)

stringstring

`CASH_FLOW_UPDATES`

[`webhook_code`](/docs/api/products/check/#CashFlowUpdatesNewLoanPaymentWebhook-webhook-code)

stringstring

`NEW_LOAN_PAYMENT_DETECTED`

[`status`](/docs/api/products/check/#CashFlowUpdatesNewLoanPaymentWebhook-status)

stringstring

Enum for the status of the insights  
  

Possible values: `AVAILABLE`, `FAILED`

[`user_id`](/docs/api/products/check/#CashFlowUpdatesNewLoanPaymentWebhook-user-id)

stringstring

The `user_id` that the report is associated with

[`environment`](/docs/api/products/check/#CashFlowUpdatesNewLoanPaymentWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CASH_FLOW_UPDATES",
  "webhook_code": "NEW_LOAN_PAYMENT_DETECTED",
  "status": "AVAILABLE",
  "user_id": "9eaba3c2fdc916bc197f279185b986607dd21682a5b04eab04a5a03e8b3f3334",
  "environment": "production"
}
```

=\*=\*=\*=

#### `NSF_OVERDRAFT_DETECTED`

For each user's item enabled for Cash Flow Updates, this webhook will fire when an update includes an NSF overdraft transaction. Upon receiving the webhook, call [`/cra/monitoring_insights/get`](/docs/api/products/check/#cramonitoring_insightsget) to retrieve the updated insights.

**Properties**

[`webhook_type`](/docs/api/products/check/#CashFlowUpdatesNSFWebhook-webhook-type)

stringstring

`CASH_FLOW_UPDATES`

[`webhook_code`](/docs/api/products/check/#CashFlowUpdatesNSFWebhook-webhook-code)

stringstring

`NSF_OVERDRAFT_DETECTED`

[`status`](/docs/api/products/check/#CashFlowUpdatesNSFWebhook-status)

stringstring

Enum for the status of the insights  
  

Possible values: `AVAILABLE`, `FAILED`

[`user_id`](/docs/api/products/check/#CashFlowUpdatesNSFWebhook-user-id)

stringstring

The `user_id` that the report is associated with

[`environment`](/docs/api/products/check/#CashFlowUpdatesNSFWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "CASH_FLOW_UPDATES",
  "webhook_code": "NSF_OVERDRAFT_DETECTED",
  "status": "AVAILABLE",
  "user_id": "9eaba3c2fdc916bc197f279185b986607dd21682a5b04eab04a5a03e8b3f3334",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
