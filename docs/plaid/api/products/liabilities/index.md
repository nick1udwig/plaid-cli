---
title: "API - Liabilities | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/liabilities/"
scraped_at: "2026-03-07T22:04:15+00:00"
---

# Liabilities

#### API reference for Liabilities endpoints and webhooks

For how-to guidance, see the [Liabilities documentation](/docs/liabilities/).

| Endpoints |  |
| --- | --- |
| [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget) | Fetch liabilities data |

| Webhooks |  |
| --- | --- |
| [`DEFAULT_UPDATE`](/docs/api/products/liabilities/#default_update) | New or updated liabilities available |

### Endpoints

=\*=\*=\*=

#### `/liabilities/get`

#### Retrieve Liabilities data

The [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget) endpoint returns various details about an Item with loan or credit accounts. Liabilities data is available primarily for US financial institutions, with some limited coverage of Canadian institutions. Currently supported account types are account type `credit` with account subtype `credit card` or `paypal`, and account type `loan` with account subtype `student` or `mortgage`. To limit accounts listed in Link to types and subtypes supported by Liabilities, you can use the `account_filters` parameter when [creating a Link token](https://plaid.com/docs/api/link/#linktokencreate).

The types of information returned by Liabilities can include balances and due dates, loan terms, and account details such as original loan amount and guarantor. Data is refreshed approximately once per day; the latest data can be retrieved by calling [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget).

/liabilities/get

**Request fields**

[`client_id`](/docs/api/products/liabilities/#liabilities-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/liabilities/#liabilities-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`access_token`](/docs/api/products/liabilities/#liabilities-get-request-access-token)

requiredstringrequired, string

The access token associated with the Item data is being requested for.

[`options`](/docs/api/products/liabilities/#liabilities-get-request-options)

objectobject

An optional object to filter `/liabilities/get` results. If provided, `options` cannot be null.

[`account_ids`](/docs/api/products/liabilities/#liabilities-get-request-options-account-ids)

[string][string]

A list of accounts to retrieve for the Item.  
An error will be returned if a provided `account_id` is not associated with the Item

/liabilities/get

```
// Retrieve Liabilities data for an Item
const request: LiabilitiesGetRequest = {
  access_token: accessToken,
};
try {
  const response = await plaidClient.liabilitiesGet(request);
  const liabilities = response.data.liabilities;
} catch (error) {
  // handle error
}
```

/liabilities/get

**Response fields**

[`accounts`](/docs/api/products/liabilities/#liabilities-get-response-accounts)

[object][object]

An array of accounts associated with the Item

[`account_id`](/docs/api/products/liabilities/#liabilities-get-response-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
When using a CRA endpoint (an endpoint associated with Plaid Check Consumer Report, i.e. any endpoint beginning with `/cra/`), the `account_id` returned will not match the `account_id` returned by a non-CRA endpoint.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`balances`](/docs/api/products/liabilities/#liabilities-get-response-accounts-balances)

objectobject

A set of fields describing the balance for an account. Balance information may be cached unless the balance object was returned by `/accounts/balance/get` or `/signal/evaluate` (using a Balance-only ruleset).

[`available`](/docs/api/products/liabilities/#liabilities-get-response-accounts-balances-available)

nullablenumbernullable, number

The amount of funds available to be withdrawn from the account, as determined by the financial institution.  
For `credit`-type accounts, the `available` balance typically equals the `limit` less the `current` balance, less any pending outflows plus any pending inflows.  
For `depository`-type accounts, the `available` balance typically equals the `current` balance less any pending outflows plus any pending inflows. For `depository`-type accounts, the `available` balance does not include the overdraft limit.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the `available` balance is the total cash available to withdraw as presented by the institution.  
Note that not all institutions calculate the `available` balance. In the event that `available` balance is unavailable, Plaid will return an `available` balance value of `null`.  
Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by `/accounts/balance/get`, or by `/signal/evaluate` with a Balance-only ruleset.  
If `current` is `null` this field is guaranteed not to be `null`.  
  

Format: `double`

[`current`](/docs/api/products/liabilities/#liabilities-get-response-accounts-balances-current)

nullablenumbernullable, number

The total amount of funds in or owed by the account.  
For `credit`-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  
For `loan`-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (`ins_116944`). For Sallie Mae student loans, the account's balance includes both principal and any outstanding interest. Similar to `credit`-type accounts, a positive balance is typically expected, while a negative amount indicates the lender owing the account holder.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the current balance is the total value of assets as presented by the institution.  
Note that balance information may be cached unless the value was returned by `/accounts/balance/get` or by `/signal/evaluate` with a Balance-only ruleset; if the Item is enabled for Transactions, the balance will be at least as recent as the most recent Transaction update. If you require realtime balance information, use the `available` balance as provided by `/accounts/balance/get` or `/signal/evaluate` called with a Balance-only `ruleset_key`.  
When returned by `/accounts/balance/get`, this field may be `null`. When this happens, `available` is guaranteed not to be `null`.  
  

Format: `double`

[`limit`](/docs/api/products/liabilities/#liabilities-get-response-accounts-balances-limit)

nullablenumbernullable, number

For `credit`-type accounts, this represents the credit limit.  
For `depository`-type accounts, this represents the pre-arranged overdraft limit, which is common for current (checking) accounts in Europe.  
In North America, this field is typically only available for `credit`-type accounts.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/liabilities/#liabilities-get-response-accounts-balances-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the balance. Always null if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/liabilities/#liabilities-get-response-accounts-balances-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the balance. Always null if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`last_updated_datetime`](/docs/api/products/liabilities/#liabilities-get-response-accounts-balances-last-updated-datetime)

nullablestringnullable, string

Timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the last time the balance was updated.  
This field is returned only when the institution is `ins_128026` (Capital One).  
  

Format: `date-time`

[`mask`](/docs/api/products/liabilities/#liabilities-get-response-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of either the account’s displayed mask or the account’s official account number. Note that the mask may be non-unique between an Item’s accounts.

[`name`](/docs/api/products/liabilities/#liabilities-get-response-accounts-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself

[`official_name`](/docs/api/products/liabilities/#liabilities-get-response-accounts-official-name)

nullablestringnullable, string

The official name of the account as given by the financial institution

[`type`](/docs/api/products/liabilities/#liabilities-get-response-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/products/liabilities/#liabilities-get-response-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`verification_status`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-status)

stringstring

Indicates an Item's micro-deposit-based verification or database verification status. This field is only populated when using Auth and falling back to micro-deposit or database verification. Possible values are:  
`pending_automatic_verification`: The Item is pending automatic verification.  
`pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the code.  
`automatically_verified`: The Item has successfully been automatically verified.  
`manually_verified`: The Item has successfully been manually verified.  
`verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  
`verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.  
`unsent`: The Item is pending micro-deposit verification, but Plaid has not yet sent the micro-deposit.  
`database_insights_pending`: The Database Auth result is pending and will be available upon Auth request.  
`database_insights_fail`: The Item's numbers have been verified using Plaid's data sources and have signal for being invalid and/or have no signal for being valid. Typically this indicates that the routing number is invalid, the account number does not match the account number format associated with the routing number, or the account has been reported as closed or frozen. Only returned for Auth Items created via Database Auth.  
`database_insights_pass`: The Item's numbers have been verified using Plaid's data sources: the routing and account number match a routing and account number of an account recognized on the Plaid network, and the account is not known by Plaid to be frozen or closed. Only returned for Auth Items created via Database Auth.  
`database_insights_pass_with_caution`: The Item's numbers have been verified using Plaid's data sources and have some signal for being valid: the routing and account number were not recognized on the Plaid network, but the routing number is valid and the account number is a potential valid account number for that routing number. Only returned for Auth Items created via Database Auth.  
`database_matched`: (deprecated) The Item has successfully been verified using Plaid's data sources. Only returned for Auth Items created via Database Match.  
`null` or empty string: Neither micro-deposit-based verification nor database verification are being used for the Item.  
  

Possible values: `automatically_verified`, `pending_automatic_verification`, `pending_manual_verification`, `unsent`, `manually_verified`, `verification_expired`, `verification_failed`, `database_matched`, `database_insights_pass`, `database_insights_pass_with_caution`, `database_insights_fail`

[`verification_name`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-name)

stringstring

The account holder name that was used for micro-deposit and/or database verification. Only returned for Auth Items created via micro-deposit or database verification. This name was manually-entered by the user during Link, unless it was otherwise provided via the `user.legal_name` request field in `/link/token/create` for the Link session that created the Item.

[`verification_insights`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-insights)

objectobject

Insights from performing database verification for the account. Only returned for Auth Items using Database Auth.

[`name_match_score`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-insights-name-match-score)

nullableintegernullable, integer

Indicates the score of the name match between the given name provided during database verification (available in the [`verification_name`](https://plaid.com/docs/api/products/auth/#auth-get-response-accounts-verification-name) field if using standard Database Auth, or provided in the request if using `/auth/verify`) and matched Plaid network accounts. If defined, will be a value between 0 and 100. Will be undefined if name matching was not enabled for the database verification session or if there were no eligible Plaid network matches to compare the given name with.

[`network_status`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-insights-network-status)

objectobject

Status information about the account and routing number in the Plaid network.

[`has_numbers_match`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-insights-network-status-has-numbers-match)

booleanboolean

Indicates whether we found at least one matching account for the ACH account and routing number.

[`is_numbers_match_verified`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-insights-network-status-is-numbers-match-verified)

booleanboolean

Indicates if at least one matching account for the ACH account and routing number is already verified.

[`previous_returns`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-insights-previous-returns)

objectobject

Information about known ACH returns for the account and routing number.

[`has_previous_administrative_return`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-insights-previous-returns-has-previous-administrative-return)

booleanboolean

Indicates whether Plaid's data sources include a known administrative ACH return for this account and routing number.

[`account_number_format`](/docs/api/products/liabilities/#liabilities-get-response-accounts-verification-insights-account-number-format)

stringstring

Indicator of account number format validity for institution.  
`valid`: indicates that the account number has a correct format for the institution.  
`invalid`: indicates that the account number has an incorrect format for the institution.  
`unknown`: indicates that there was not enough information to determine whether the format is correct for the institution.  
  

Possible values: `valid`, `invalid`, `unknown`

[`persistent_account_id`](/docs/api/products/liabilities/#liabilities-get-response-accounts-persistent-account-id)

stringstring

A unique and persistent identifier for accounts that can be used to trace multiple instances of the same account across different Items for depository accounts. This field is currently supported only for Items at institutions that use Tokenized Account Numbers (i.e., Chase and PNC, and in May 2025 US Bank). Because these accounts have a different account number each time they are linked, this field may be used instead of the account number to uniquely identify an account across multiple Items for payments use cases, helping to reduce duplicate Items or attempted fraud. In Sandbox, this field is populated for TAN-based institutions (`ins_56`, `ins_13`) as well as the OAuth Sandbox institution (`ins_127287`); in Production, it will only be populated for accounts at applicable institutions.

[`holder_category`](/docs/api/products/liabilities/#liabilities-get-response-accounts-holder-category)

nullablestringnullable, string

Indicates the account's categorization as either a personal or a business account. This field is currently in beta; to request access, contact your account manager.  
  

Possible values: `business`, `personal`, `unrecognized`

[`item`](/docs/api/products/liabilities/#liabilities-get-response-item)

objectobject

Metadata about the Item.

[`item_id`](/docs/api/products/liabilities/#liabilities-get-response-item-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`institution_id`](/docs/api/products/liabilities/#liabilities-get-response-item-institution-id)

nullablestringnullable, string

The Plaid Institution ID associated with the Item. Field is `null` for Items created without an institution connection, such as Items created via Same Day Micro-deposits.

[`institution_name`](/docs/api/products/liabilities/#liabilities-get-response-item-institution-name)

nullablestringnullable, string

The name of the institution associated with the Item. Field is `null` for Items created without an institution connection, such as Items created via Same Day Micro-deposits.

[`webhook`](/docs/api/products/liabilities/#liabilities-get-response-item-webhook)

nullablestringnullable, string

The URL registered to receive webhooks for the Item.

[`auth_method`](/docs/api/products/liabilities/#liabilities-get-response-item-auth-method)

nullablestringnullable, string

The method used to populate Auth data for the Item. This field is only populated for Items that have had Auth numbers data set on at least one of its accounts, and will be `null` otherwise. For info about the various flows, see our [Auth coverage documentation](https://plaid.com/docs/auth/coverage/).  
`INSTANT_AUTH`: The Item's Auth data was provided directly by the user's institution connection.  
`INSTANT_MATCH`: The Item's Auth data was provided via the Instant Match fallback flow.  
`AUTOMATED_MICRODEPOSITS`: The Item's Auth data was provided via the Automated Micro-deposits flow.  
`SAME_DAY_MICRODEPOSITS`: The Item's Auth data was provided via the Same Day Micro-deposits flow.  
`INSTANT_MICRODEPOSITS`: The Item's Auth data was provided via the Instant Micro-deposits flow.  
`DATABASE_MATCH`: The Item's Auth data was provided via the Database Match flow.  
`DATABASE_INSIGHTS`: The Item's Auth data was provided via the Database Insights flow.  
`TRANSFER_MIGRATED`: The Item's Auth data was provided via [`/transfer/migrate_account`](https://plaid.com/docs/api/products/transfer/account-linking/#migrate-account-into-transfers).  
`INVESTMENTS_FALLBACK`: The Item's Auth data for Investments Move was provided via a [fallback flow](https://plaid.com/docs/investments-move/#fallback-flows).  
  

Possible values: `INSTANT_AUTH`, `INSTANT_MATCH`, `AUTOMATED_MICRODEPOSITS`, `SAME_DAY_MICRODEPOSITS`, `INSTANT_MICRODEPOSITS`, `DATABASE_MATCH`, `DATABASE_INSIGHTS`, `TRANSFER_MIGRATED`, `INVESTMENTS_FALLBACK`, `null`

[`error`](/docs/api/products/liabilities/#liabilities-get-response-item-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/liabilities/#liabilities-get-response-item-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/liabilities/#liabilities-get-response-item-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/liabilities/#liabilities-get-response-item-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/liabilities/#liabilities-get-response-item-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/liabilities/#liabilities-get-response-item-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/liabilities/#liabilities-get-response-item-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/liabilities/#liabilities-get-response-item-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/liabilities/#liabilities-get-response-item-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/liabilities/#liabilities-get-response-item-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/liabilities/#liabilities-get-response-item-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/liabilities/#liabilities-get-response-item-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/liabilities/#liabilities-get-response-item-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`available_products`](/docs/api/products/liabilities/#liabilities-get-response-item-available-products)

[string][string]

A list of products available for the Item that have not yet been accessed. The contents of this array will be mutually exclusive with `billed_products`.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`billed_products`](/docs/api/products/liabilities/#liabilities-get-response-item-billed-products)

[string][string]

A list of products that have been billed for the Item. The contents of this array will be mutually exclusive with `available_products`. Note - `billed_products` is populated in all environments but only requests in Production are billed. Also note that products that are billed on a pay-per-call basis rather than a pay-per-Item basis, such as `balance`, will not appear here.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`products`](/docs/api/products/liabilities/#liabilities-get-response-item-products)

[string][string]

A list of products added to the Item. In almost all cases, this will be the same as the `billed_products` field. For some products, it is possible for the product to be added to an Item but not yet billed (e.g. Assets, before `/asset_report/create` has been called, or Auth or Identity when added as Optional Products but before their endpoints have been called), in which case the product may appear in `products` but not in `billed_products`.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`consented_products`](/docs/api/products/liabilities/#liabilities-get-response-item-consented-products)

[string][string]

A list of products that the user has consented to for the Item via [Data Transparency Messaging](https://plaid.com/docs/link/data-transparency-messaging-migration-guide). This will consist of all products where both of the following are true: the user has consented to the required data scopes for that product and you have Production access for that product.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `transactions`, `income`, `income_verification`, `transfer`, `employment`, `recurring_transactions`, `signal`, `statements`, `processor_payments`, `processor_identity`, `cra_base_report`, `cra_income_insights`, `cra_lend_score`, `cra_partner_insights`, `cra_cashflow_insights`, `cra_monitoring`, `layer`

[`consent_expiration_time`](/docs/api/products/liabilities/#liabilities-get-response-item-consent-expiration-time)

nullablestringnullable, string

The date and time at which the Item's access consent will expire, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. If the Item does not have consent expiration scheduled, this field will be `null`. Currently, only institutions in Europe and a small number of institutions in the US have expiring consent. For a list of US institutions that currently expire consent, see the [OAuth Guide](https://plaid.com/docs/link/oauth/#refreshing-item-consent).  
  

Format: `date-time`

[`update_type`](/docs/api/products/liabilities/#liabilities-get-response-item-update-type)

stringstring

Indicates whether an Item requires user interaction to be updated, which can be the case for Items with some forms of two-factor authentication.  
`background` - Item can be updated in the background  
`user_present_required` - Item requires user interaction to be updated  
  

Possible values: `background`, `user_present_required`

[`liabilities`](/docs/api/products/liabilities/#liabilities-get-response-liabilities)

objectobject

An object containing liability accounts

[`credit`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit)

nullable[object]nullable, [object]

The credit accounts returned.

[`account_id`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-account-id)

nullablestringnullable, string

The ID of the account that this liability belongs to.

[`aprs`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-aprs)

[object][object]

The various interest rates that apply to the account. APR information is not provided by all card issuers; if APR data is not available, this array will be empty.

[`apr_percentage`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-aprs-apr-percentage)

numbernumber

Annual Percentage Rate applied.  
  

Format: `double`

[`apr_type`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-aprs-apr-type)

stringstring

The type of balance to which the APR applies.  
  

Possible values: `balance_transfer_apr`, `cash_apr`, `purchase_apr`, `special`

[`balance_subject_to_apr`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-aprs-balance-subject-to-apr)

nullablenumbernullable, number

Amount of money that is subjected to the APR if a balance was carried beyond payment due date. How it is calculated can vary by card issuer. It is often calculated as an average daily balance.  
  

Format: `double`

[`interest_charge_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-aprs-interest-charge-amount)

nullablenumbernullable, number

Amount of money charged due to interest from last statement.  
  

Format: `double`

[`is_overdue`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-is-overdue)

nullablebooleannullable, boolean

true if a payment is currently overdue. Availability for this field is limited.

[`last_payment_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-last-payment-amount)

nullablenumbernullable, number

The amount of the last payment.  
  

Format: `double`

[`last_payment_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-last-payment-date)

nullablestringnullable, string

The date of the last payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Availability for this field is limited.  
  

Format: `date`

[`last_statement_issue_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-last-statement-issue-date)

nullablestringnullable, string

The date of the last statement. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`last_statement_balance`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-last-statement-balance)

nullablenumbernullable, number

The total amount owed as of the last statement issued  
  

Format: `double`

[`minimum_payment_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-minimum-payment-amount)

nullablenumbernullable, number

The minimum payment due for the next billing cycle.  
  

Format: `double`

[`next_payment_due_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-credit-next-payment-due-date)

nullablestringnullable, string

The due date for the next payment. The due date is `null` if a payment is not expected. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`mortgage`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage)

nullable[object]nullable, [object]

The mortgage accounts returned.

[`account_id`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-account-id)

stringstring

The ID of the account that this liability belongs to.

[`account_number`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-account-number)

nullablestringnullable, string

The account number of the loan.

[`current_late_fee`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-current-late-fee)

nullablenumbernullable, number

The current outstanding amount charged for late payment.  
  

Format: `double`

[`escrow_balance`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-escrow-balance)

nullablenumbernullable, number

Total amount held in escrow to pay taxes and insurance on behalf of the borrower.  
  

Format: `double`

[`has_pmi`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-has-pmi)

nullablebooleannullable, boolean

Indicates whether the borrower has private mortgage insurance in effect.

[`has_prepayment_penalty`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-has-prepayment-penalty)

nullablebooleannullable, boolean

Indicates whether the borrower will pay a penalty for early payoff of mortgage.

[`interest_rate`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-interest-rate)

objectobject

Object containing metadata about the interest rate for the mortgage.

[`percentage`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-interest-rate-percentage)

nullablenumbernullable, number

Percentage value (interest rate of current mortgage, not APR) of interest payable on a loan.  
  

Format: `double`

[`type`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-interest-rate-type)

nullablestringnullable, string

The type of interest charged (fixed or variable).

[`last_payment_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-last-payment-amount)

nullablenumbernullable, number

The amount of the last payment.  
  

Format: `double`

[`last_payment_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-last-payment-date)

nullablestringnullable, string

The date of the last payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`loan_type_description`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-loan-type-description)

nullablestringnullable, string

Description of the type of loan, for example `conventional`, `fixed`, or `variable`. This field is provided directly from the loan servicer and does not have an enumerated set of possible values.

[`loan_term`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-loan-term)

nullablestringnullable, string

Full duration of mortgage as at origination (e.g. `10 year`).

[`maturity_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-maturity-date)

nullablestringnullable, string

Original date on which mortgage is due in full. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`next_monthly_payment`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-next-monthly-payment)

nullablenumbernullable, number

The amount of the next payment.  
  

Format: `double`

[`next_payment_due_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-next-payment-due-date)

nullablestringnullable, string

The due date for the next payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`origination_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-origination-date)

nullablestringnullable, string

The date on which the loan was initially lent. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`origination_principal_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-origination-principal-amount)

nullablenumbernullable, number

The original principal balance of the mortgage.  
  

Format: `double`

[`past_due_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-past-due-amount)

nullablenumbernullable, number

Amount of loan (principal + interest) past due for payment.  
  

Format: `double`

[`property_address`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-property-address)

objectobject

Object containing fields describing property address.

[`city`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-property-address-city)

nullablestringnullable, string

The city name.

[`country`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-property-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-property-address-postal-code)

nullablestringnullable, string

The five or nine digit postal code.

[`region`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-property-address-region)

nullablestringnullable, string

The region or state (example "NC").

[`street`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-property-address-street)

nullablestringnullable, string

The full street address (example "564 Main Street, Apt 15").

[`ytd_interest_paid`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-ytd-interest-paid)

nullablenumbernullable, number

The year to date (YTD) interest paid.  
  

Format: `double`

[`ytd_principal_paid`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-mortgage-ytd-principal-paid)

nullablenumbernullable, number

The YTD principal paid.  
  

Format: `double`

[`student`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student)

nullable[object]nullable, [object]

The student loan accounts returned.

[`account_id`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-account-id)

nullablestringnullable, string

The ID of the account that this liability belongs to. Each account can only contain one liability.

[`account_number`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-account-number)

nullablestringnullable, string

The account number of the loan. For some institutions, this may be a masked version of the number (e.g., the last 4 digits instead of the entire number).

[`disbursement_dates`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-disbursement-dates)

nullable[string]nullable, [string]

The dates on which loaned funds were disbursed or will be disbursed. These are often in the past. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`expected_payoff_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-expected-payoff-date)

nullablestringnullable, string

The date when the student loan is expected to be paid off. Availability for this field is limited. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`guarantor`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-guarantor)

nullablestringnullable, string

The guarantor of the student loan.

[`interest_rate_percentage`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-interest-rate-percentage)

numbernumber

The interest rate on the loan as a percentage.  
  

Format: `double`

[`is_overdue`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-is-overdue)

nullablebooleannullable, boolean

`true` if a payment is currently overdue. Availability for this field is limited.

[`last_payment_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-last-payment-amount)

nullablenumbernullable, number

The amount of the last payment.  
  

Format: `double`

[`last_payment_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-last-payment-date)

nullablestringnullable, string

The date of the last payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`last_statement_balance`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-last-statement-balance)

nullablenumbernullable, number

The total amount owed as of the last statement issued  
  

Format: `double`

[`last_statement_issue_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-last-statement-issue-date)

nullablestringnullable, string

The date of the last statement. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`loan_name`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-loan-name)

nullablestringnullable, string

The type of loan, e.g., "Consolidation Loans".

[`loan_status`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-loan-status)

objectobject

An object representing the status of the student loan

[`end_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-loan-status-end-date)

nullablestringnullable, string

The date until which the loan will be in its current status. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`type`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-loan-status-type)

nullablestringnullable, string

The status type of the student loan  
  

Possible values: `cancelled`, `charged off`, `claim`, `consolidated`, `deferment`, `delinquent`, `discharged`, `extension`, `forbearance`, `in grace`, `in military`, `in school`, `not fully disbursed`, `other`, `paid in full`, `refunded`, `repayment`, `transferred`, `pending idr`

[`minimum_payment_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-minimum-payment-amount)

nullablenumbernullable, number

The minimum payment due for the next billing cycle. There are some exceptions:
Some institutions require a minimum payment across all loans associated with an account number. Our API presents that same minimum payment amount on each loan. The institutions that do this are: Great Lakes ( `ins_116861`), Firstmark (`ins_116295`), Commonbond Firstmark Services (`ins_116950`), Granite State (`ins_116308`), and Oklahoma Student Loan Authority (`ins_116945`).
Firstmark (`ins_116295` ) and Navient (`ins_116248`) will display as $0 if there is an autopay program in effect.  
  

Format: `double`

[`next_payment_due_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-next-payment-due-date)

nullablestringnullable, string

The due date for the next payment. The due date is `null` if a payment is not expected. A payment is not expected if `loan_status.type` is `deferment`, `in_school`, `consolidated`, `paid in full`, or `transferred`. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`origination_date`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-origination-date)

nullablestringnullable, string

The date on which the loan was initially lent. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`origination_principal_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-origination-principal-amount)

nullablenumbernullable, number

The original principal balance of the loan.  
  

Format: `double`

[`outstanding_interest_amount`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-outstanding-interest-amount)

nullablenumbernullable, number

The total dollar amount of the accrued interest balance. For Sallie Mae ( `ins_116944`), this amount is included in the current balance of the loan, so this field will return as `null`.  
  

Format: `double`

[`payment_reference_number`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-payment-reference-number)

nullablestringnullable, string

The relevant account number that should be used to reference this loan for payments. In the majority of cases, `payment_reference_number` will match `account_number,` but in some institutions, such as Great Lakes (`ins_116861`), it will be different.

[`repayment_plan`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-repayment-plan)

objectobject

An object representing the repayment plan for the student loan

[`description`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-repayment-plan-description)

nullablestringnullable, string

The description of the repayment plan as provided by the servicer.

[`type`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-repayment-plan-type)

nullablestringnullable, string

The type of the repayment plan.  
  

Possible values: `extended graduated`, `extended standard`, `graduated`, `income-contingent repayment`, `income-based repayment`, `income-sensitive repayment`, `interest-only`, `other`, `pay as you earn`, `revised pay as you earn`, `standard`, `saving on a valuable education`, `null`

[`sequence_number`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-sequence-number)

nullablestringnullable, string

The sequence number of the student loan. Heartland ECSI (`ins_116948`) does not make this field available.

[`servicer_address`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-servicer-address)

objectobject

The address of the student loan servicer. This is generally the remittance address to which payments should be sent.

[`city`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-servicer-address-city)

nullablestringnullable, string

The full city name

[`region`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-servicer-address-region)

nullablestringnullable, string

The region or state
Example: `"NC"`

[`street`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-servicer-address-street)

nullablestringnullable, string

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-servicer-address-postal-code)

nullablestringnullable, string

The postal code

[`country`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-servicer-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code

[`ytd_interest_paid`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-ytd-interest-paid)

nullablenumbernullable, number

The year to date (YTD) interest paid. Availability for this field is limited.  
  

Format: `double`

[`ytd_principal_paid`](/docs/api/products/liabilities/#liabilities-get-response-liabilities-student-ytd-principal-paid)

nullablenumbernullable, number

The year to date (YTD) principal paid. Availability for this field is limited.  
  

Format: `double`

[`request_id`](/docs/api/products/liabilities/#liabilities-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "accounts": [
    {
      "account_id": "BxBXxLj1m4HMXBm9WZZmCWVbPjX16EHwv99vp",
      "balances": {
        "available": 100,
        "current": 110,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "0000",
      "name": "Plaid Checking",
      "official_name": "Plaid Gold Standard 0% Interest Checking",
      "subtype": "checking",
      "type": "depository"
    },
    {
      "account_id": "dVzbVMLjrxTnLjX4G66XUp5GLklm4oiZy88yK",
      "balances": {
        "available": null,
        "current": 410,
        "iso_currency_code": "USD",
        "limit": 2000,
        "unofficial_currency_code": null
      },
      "mask": "3333",
      "name": "Plaid Credit Card",
      "official_name": "Plaid Diamond 12.5% APR Interest Credit Card",
      "subtype": "credit card",
      "type": "credit"
    },
    {
      "account_id": "Pp1Vpkl9w8sajvK6oEEKtr7vZxBnGpf7LxxLE",
      "balances": {
        "available": null,
        "current": 65262,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "7777",
      "name": "Plaid Student Loan",
      "official_name": null,
      "subtype": "student",
      "type": "loan"
    },
    {
      "account_id": "BxBXxLj1m4HMXBm9WZJyUg9XLd4rKEhw8Pb1J",
      "balances": {
        "available": null,
        "current": 56302.06,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "8888",
      "name": "Plaid Mortgage",
      "official_name": null,
      "subtype": "mortgage",
      "type": "loan"
    }
  ],
  "item": {
    "available_products": [
      "balance",
      "investments"
    ],
    "billed_products": [
      "assets",
      "auth",
      "identity",
      "liabilities",
      "transactions"
    ],
    "consent_expiration_time": null,
    "error": null,
    "institution_id": "ins_3",
    "institution_name": "Chase",
    "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
    "update_type": "background",
    "webhook": "https://www.genericwebhookurl.com/webhook",
    "auth_method": "INSTANT_AUTH"
  },
  "liabilities": {
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
    ],
    "mortgage": [
      {
        "account_id": "BxBXxLj1m4HMXBm9WZJyUg9XLd4rKEhw8Pb1J",
        "account_number": "3120194154",
        "current_late_fee": 25,
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
    ],
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
        "last_statement_balance": 1708.77,
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
  },
  "request_id": "dTnnm60WgKGLnKL"
}
```

### Webhooks

Liabilities webhooks are sent to indicate that new loans or updated loan fields
for existing accounts are available.

=\*=\*=\*=

#### `DEFAULT_UPDATE`

The webhook of type `LIABILITIES` and code `DEFAULT_UPDATE` will be fired when new or updated liabilities have been detected on a liabilities item.

**Properties**

[`webhook_type`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-webhook-type)

stringstring

`LIABILITIES`

[`webhook_code`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-webhook-code)

stringstring

`DEFAULT_UPDATE`

[`item_id`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`user_id`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-user-id)

stringstring

The Plaid `user_id` of the User associated with this webhook, warning, or error.

[`error`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error)

objectobject

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-error-code-reason)

stringstring

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-display-message)

stringstring

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-status)

integerinteger

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-suggested-action)

stringstring

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`account_ids_with_new_liabilities`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-account-ids-with-new-liabilities)

[string][string]

An array of `account_id`'s for accounts that contain new liabilities.'

[`account_ids_with_updated_liabilities`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-account-ids-with-updated-liabilities)

objectobject

An object with keys of `account_id`'s that are mapped to their respective liabilities fields that changed.  
Example: `{ "XMBvvyMGQ1UoLbKByoMqH3nXMj84ALSdE5B58": ["past_amount_due"] }`

[`environment`](/docs/api/products/liabilities/#LiabilitiesDefaultUpdateWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "LIABILITIES",
  "webhook_code": "DEFAULT_UPDATE",
  "item_id": "wz666MBjYWTp2PDzzggYhM6oWWmBb",
  "user_id": "usr_9nSp2KuZ2x4JDw",
  "error": null,
  "account_ids_with_new_liabilities": [
    "XMBvvyMGQ1UoLbKByoMqH3nXMj84ALSdE5B58",
    "BxBXxLj1m4HMXBm9WZZmCWVbPjX16EHwv99vp"
  ],
  "account_ids_with_updated_liabilities": {
    "XMBvvyMGQ1UoLbKByoMqH3nXMj84ALSdE5B58": [
      "past_amount_due"
    ]
  },
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
