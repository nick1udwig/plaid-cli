---
title: "API - Accounts | Plaid Docs"
source_url: "https://plaid.com/docs/api/accounts/"
scraped_at: "2026-03-07T22:03:46+00:00"
---

# Accounts

#### API reference for retrieving account information and seeing all possible account types and subtypes

=\*=\*=\*=

#### `/accounts/get`

#### Retrieve accounts

The [`/accounts/get`](/docs/api/accounts/#accountsget) endpoint can be used to retrieve a list of accounts associated with any linked Item. Plaid will only return active bank accounts — that is, accounts that are not closed and are capable of carrying a balance.
To return new accounts that were created after the user linked their Item, you can listen for the [`NEW_ACCOUNTS_AVAILABLE`](https://plaid.com/docs/api/items/#new_accounts_available) webhook and then use Link's [update mode](https://plaid.com/docs/link/update-mode/) to request that the user share this new account with you.

[`/accounts/get`](/docs/api/accounts/#accountsget) is free to use and retrieves cached information, rather than extracting fresh information from the institution. The balance returned will reflect the balance at the time of the last successful Item update. If the Item is enabled for a regularly updating product, such as Transactions, Investments, or Liabilities, the balance will typically update about once a day, as long as the Item is healthy. If the Item is enabled only for products that do not frequently update, such as Auth or Identity, balance data may be much older.

For realtime balance information, use the paid endpoints [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) or [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) instead.

/accounts/get

**Request fields**

[`client_id`](/docs/api/accounts/#accounts-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/accounts/#accounts-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`access_token`](/docs/api/accounts/#accounts-get-request-access-token)

requiredstringrequired, string

The access token associated with the Item data is being requested for.

[`options`](/docs/api/accounts/#accounts-get-request-options)

objectobject

An optional object to filter `/accounts/get` results.

[`account_ids`](/docs/api/accounts/#accounts-get-request-options-account-ids)

[string][string]

An array of `account_ids` to retrieve for the Account.

/accounts/get

```
const request: AccountsGetRequest = {
  access_token: ACCESS_TOKEN,
};
try {
  const response = await plaidClient.accountsGet(request);
  const accounts = response.data.accounts;
} catch (error) {
  // handle error
}
```

/accounts/get

**Response fields**

[`accounts`](/docs/api/accounts/#accounts-get-response-accounts)

[object][object]

An array of financial institution accounts associated with the Item.
If `/accounts/balance/get` was called, each account will include real-time balance information.

[`account_id`](/docs/api/accounts/#accounts-get-response-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
When using a CRA endpoint (an endpoint associated with Plaid Check Consumer Report, i.e. any endpoint beginning with `/cra/`), the `account_id` returned will not match the `account_id` returned by a non-CRA endpoint.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`balances`](/docs/api/accounts/#accounts-get-response-accounts-balances)

objectobject

A set of fields describing the balance for an account. Balance information may be cached unless the balance object was returned by `/accounts/balance/get` or `/signal/evaluate` (using a Balance-only ruleset).

[`available`](/docs/api/accounts/#accounts-get-response-accounts-balances-available)

nullablenumbernullable, number

The amount of funds available to be withdrawn from the account, as determined by the financial institution.  
For `credit`-type accounts, the `available` balance typically equals the `limit` less the `current` balance, less any pending outflows plus any pending inflows.  
For `depository`-type accounts, the `available` balance typically equals the `current` balance less any pending outflows plus any pending inflows. For `depository`-type accounts, the `available` balance does not include the overdraft limit.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the `available` balance is the total cash available to withdraw as presented by the institution.  
Note that not all institutions calculate the `available` balance. In the event that `available` balance is unavailable, Plaid will return an `available` balance value of `null`.  
Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by `/accounts/balance/get`, or by `/signal/evaluate` with a Balance-only ruleset.  
If `current` is `null` this field is guaranteed not to be `null`.  
  

Format: `double`

[`current`](/docs/api/accounts/#accounts-get-response-accounts-balances-current)

nullablenumbernullable, number

The total amount of funds in or owed by the account.  
For `credit`-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  
For `loan`-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (`ins_116944`). For Sallie Mae student loans, the account's balance includes both principal and any outstanding interest. Similar to `credit`-type accounts, a positive balance is typically expected, while a negative amount indicates the lender owing the account holder.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the current balance is the total value of assets as presented by the institution.  
Note that balance information may be cached unless the value was returned by `/accounts/balance/get` or by `/signal/evaluate` with a Balance-only ruleset; if the Item is enabled for Transactions, the balance will be at least as recent as the most recent Transaction update. If you require realtime balance information, use the `available` balance as provided by `/accounts/balance/get` or `/signal/evaluate` called with a Balance-only `ruleset_key`.  
When returned by `/accounts/balance/get`, this field may be `null`. When this happens, `available` is guaranteed not to be `null`.  
  

Format: `double`

[`limit`](/docs/api/accounts/#accounts-get-response-accounts-balances-limit)

nullablenumbernullable, number

For `credit`-type accounts, this represents the credit limit.  
For `depository`-type accounts, this represents the pre-arranged overdraft limit, which is common for current (checking) accounts in Europe.  
In North America, this field is typically only available for `credit`-type accounts.  
  

Format: `double`

[`iso_currency_code`](/docs/api/accounts/#accounts-get-response-accounts-balances-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the balance. Always null if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/accounts/#accounts-get-response-accounts-balances-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the balance. Always null if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`last_updated_datetime`](/docs/api/accounts/#accounts-get-response-accounts-balances-last-updated-datetime)

nullablestringnullable, string

Timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the last time the balance was updated.  
This field is returned only when the institution is `ins_128026` (Capital One).  
  

Format: `date-time`

[`mask`](/docs/api/accounts/#accounts-get-response-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of either the account’s displayed mask or the account’s official account number. Note that the mask may be non-unique between an Item’s accounts.

[`name`](/docs/api/accounts/#accounts-get-response-accounts-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself

[`official_name`](/docs/api/accounts/#accounts-get-response-accounts-official-name)

nullablestringnullable, string

The official name of the account as given by the financial institution

[`type`](/docs/api/accounts/#accounts-get-response-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/accounts/#accounts-get-response-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`verification_status`](/docs/api/accounts/#accounts-get-response-accounts-verification-status)

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

[`verification_name`](/docs/api/accounts/#accounts-get-response-accounts-verification-name)

stringstring

The account holder name that was used for micro-deposit and/or database verification. Only returned for Auth Items created via micro-deposit or database verification. This name was manually-entered by the user during Link, unless it was otherwise provided via the `user.legal_name` request field in `/link/token/create` for the Link session that created the Item.

[`verification_insights`](/docs/api/accounts/#accounts-get-response-accounts-verification-insights)

objectobject

Insights from performing database verification for the account. Only returned for Auth Items using Database Auth.

[`name_match_score`](/docs/api/accounts/#accounts-get-response-accounts-verification-insights-name-match-score)

nullableintegernullable, integer

Indicates the score of the name match between the given name provided during database verification (available in the [`verification_name`](https://plaid.com/docs/api/products/auth/#auth-get-response-accounts-verification-name) field if using standard Database Auth, or provided in the request if using `/auth/verify`) and matched Plaid network accounts. If defined, will be a value between 0 and 100. Will be undefined if name matching was not enabled for the database verification session or if there were no eligible Plaid network matches to compare the given name with.

[`network_status`](/docs/api/accounts/#accounts-get-response-accounts-verification-insights-network-status)

objectobject

Status information about the account and routing number in the Plaid network.

[`has_numbers_match`](/docs/api/accounts/#accounts-get-response-accounts-verification-insights-network-status-has-numbers-match)

booleanboolean

Indicates whether we found at least one matching account for the ACH account and routing number.

[`is_numbers_match_verified`](/docs/api/accounts/#accounts-get-response-accounts-verification-insights-network-status-is-numbers-match-verified)

booleanboolean

Indicates if at least one matching account for the ACH account and routing number is already verified.

[`previous_returns`](/docs/api/accounts/#accounts-get-response-accounts-verification-insights-previous-returns)

objectobject

Information about known ACH returns for the account and routing number.

[`has_previous_administrative_return`](/docs/api/accounts/#accounts-get-response-accounts-verification-insights-previous-returns-has-previous-administrative-return)

booleanboolean

Indicates whether Plaid's data sources include a known administrative ACH return for this account and routing number.

[`account_number_format`](/docs/api/accounts/#accounts-get-response-accounts-verification-insights-account-number-format)

stringstring

Indicator of account number format validity for institution.  
`valid`: indicates that the account number has a correct format for the institution.  
`invalid`: indicates that the account number has an incorrect format for the institution.  
`unknown`: indicates that there was not enough information to determine whether the format is correct for the institution.  
  

Possible values: `valid`, `invalid`, `unknown`

[`persistent_account_id`](/docs/api/accounts/#accounts-get-response-accounts-persistent-account-id)

stringstring

A unique and persistent identifier for accounts that can be used to trace multiple instances of the same account across different Items for depository accounts. This field is currently supported only for Items at institutions that use Tokenized Account Numbers (i.e., Chase and PNC, and in May 2025 US Bank). Because these accounts have a different account number each time they are linked, this field may be used instead of the account number to uniquely identify an account across multiple Items for payments use cases, helping to reduce duplicate Items or attempted fraud. In Sandbox, this field is populated for TAN-based institutions (`ins_56`, `ins_13`) as well as the OAuth Sandbox institution (`ins_127287`); in Production, it will only be populated for accounts at applicable institutions.

[`holder_category`](/docs/api/accounts/#accounts-get-response-accounts-holder-category)

nullablestringnullable, string

Indicates the account's categorization as either a personal or a business account. This field is currently in beta; to request access, contact your account manager.  
  

Possible values: `business`, `personal`, `unrecognized`

[`item`](/docs/api/accounts/#accounts-get-response-item)

objectobject

Metadata about the Item.

[`item_id`](/docs/api/accounts/#accounts-get-response-item-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`institution_id`](/docs/api/accounts/#accounts-get-response-item-institution-id)

nullablestringnullable, string

The Plaid Institution ID associated with the Item. Field is `null` for Items created without an institution connection, such as Items created via Same Day Micro-deposits.

[`institution_name`](/docs/api/accounts/#accounts-get-response-item-institution-name)

nullablestringnullable, string

The name of the institution associated with the Item. Field is `null` for Items created without an institution connection, such as Items created via Same Day Micro-deposits.

[`webhook`](/docs/api/accounts/#accounts-get-response-item-webhook)

nullablestringnullable, string

The URL registered to receive webhooks for the Item.

[`auth_method`](/docs/api/accounts/#accounts-get-response-item-auth-method)

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

[`error`](/docs/api/accounts/#accounts-get-response-item-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/accounts/#accounts-get-response-item-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/accounts/#accounts-get-response-item-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/accounts/#accounts-get-response-item-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/accounts/#accounts-get-response-item-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/accounts/#accounts-get-response-item-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/accounts/#accounts-get-response-item-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/accounts/#accounts-get-response-item-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/accounts/#accounts-get-response-item-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/accounts/#accounts-get-response-item-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/accounts/#accounts-get-response-item-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/accounts/#accounts-get-response-item-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/accounts/#accounts-get-response-item-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`available_products`](/docs/api/accounts/#accounts-get-response-item-available-products)

[string][string]

A list of products available for the Item that have not yet been accessed. The contents of this array will be mutually exclusive with `billed_products`.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`billed_products`](/docs/api/accounts/#accounts-get-response-item-billed-products)

[string][string]

A list of products that have been billed for the Item. The contents of this array will be mutually exclusive with `available_products`. Note - `billed_products` is populated in all environments but only requests in Production are billed. Also note that products that are billed on a pay-per-call basis rather than a pay-per-Item basis, such as `balance`, will not appear here.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`products`](/docs/api/accounts/#accounts-get-response-item-products)

[string][string]

A list of products added to the Item. In almost all cases, this will be the same as the `billed_products` field. For some products, it is possible for the product to be added to an Item but not yet billed (e.g. Assets, before `/asset_report/create` has been called, or Auth or Identity when added as Optional Products but before their endpoints have been called), in which case the product may appear in `products` but not in `billed_products`.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`consented_products`](/docs/api/accounts/#accounts-get-response-item-consented-products)

[string][string]

A list of products that the user has consented to for the Item via [Data Transparency Messaging](https://plaid.com/docs/link/data-transparency-messaging-migration-guide). This will consist of all products where both of the following are true: the user has consented to the required data scopes for that product and you have Production access for that product.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `transactions`, `income`, `income_verification`, `transfer`, `employment`, `recurring_transactions`, `signal`, `statements`, `processor_payments`, `processor_identity`, `cra_base_report`, `cra_income_insights`, `cra_lend_score`, `cra_partner_insights`, `cra_cashflow_insights`, `cra_monitoring`, `layer`

[`consent_expiration_time`](/docs/api/accounts/#accounts-get-response-item-consent-expiration-time)

nullablestringnullable, string

The date and time at which the Item's access consent will expire, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. If the Item does not have consent expiration scheduled, this field will be `null`. Currently, only institutions in Europe and a small number of institutions in the US have expiring consent. For a list of US institutions that currently expire consent, see the [OAuth Guide](https://plaid.com/docs/link/oauth/#refreshing-item-consent).  
  

Format: `date-time`

[`update_type`](/docs/api/accounts/#accounts-get-response-item-update-type)

stringstring

Indicates whether an Item requires user interaction to be updated, which can be the case for Items with some forms of two-factor authentication.  
`background` - Item can be updated in the background  
`user_present_required` - Item requires user interaction to be updated  
  

Possible values: `background`, `user_present_required`

[`request_id`](/docs/api/accounts/#accounts-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "accounts": [
    {
      "account_id": "blgvvBlXw3cq5GMPwqB6s6q4dLKB9WcVqGDGo",
      "balances": {
        "available": 100,
        "current": 110,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "holder_category": "personal",
      "mask": "0000",
      "name": "Plaid Checking",
      "official_name": "Plaid Gold Standard 0% Interest Checking",
      "subtype": "checking",
      "type": "depository"
    },
    {
      "account_id": "6PdjjRP6LmugpBy5NgQvUqpRXMWxzktg3rwrk",
      "balances": {
        "available": null,
        "current": 23631.9805,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "6666",
      "name": "Plaid 401k",
      "official_name": null,
      "subtype": "401k",
      "type": "investment"
    },
    {
      "account_id": "XMBvvyMGQ1UoLbKByoMqH3nXMj84ALSdE5B58",
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
    }
  ],
  "item": {
    "available_products": [
      "balance",
      "identity",
      "payment_initiation",
      "transactions"
    ],
    "billed_products": [
      "assets",
      "auth"
    ],
    "consent_expiration_time": null,
    "error": null,
    "institution_id": "ins_117650",
    "institution_name": "Royal Bank of Plaid",
    "item_id": "DWVAAPWq4RHGlEaNyGKRTAnPLaEmo8Cvq7na6",
    "update_type": "background",
    "webhook": "https://www.genericwebhookurl.com/webhook",
    "auth_method": "INSTANT_AUTH"
  },
  "request_id": "bkVE1BHWMAZ9Rnr"
}
```

=\*=\*=\*=

#### Account type schema

The schema below describes the various `types` and corresponding `subtypes` that Plaid recognizes and reports for financial institution accounts. For a mapping of supported types and subtypes to Plaid products, see the [Account type / product support matrix](https://plaid.com/docs/api/accounts/#account-type--product-support-matrix).

**Properties**

[`depository`](/docs/api/accounts/#StandaloneAccountType-depository)

stringstring

An account type holding cash, in which funds are deposited.

[`cash management`](/docs/api/accounts/#StandaloneAccountType-depository-cash%20management)

stringstring

A cash management account, typically a cash account at a brokerage

[`cd`](/docs/api/accounts/#StandaloneAccountType-depository-cd)

stringstring

Certificate of deposit account

[`checking`](/docs/api/accounts/#StandaloneAccountType-depository-checking)

stringstring

Checking account

[`ebt`](/docs/api/accounts/#StandaloneAccountType-depository-ebt)

stringstring

An Electronic Benefit Transfer (EBT) account, used by certain public assistance programs to distribute funds (US only)

[`hsa`](/docs/api/accounts/#StandaloneAccountType-depository-hsa)

stringstring

Health Savings Account (US only) that can only hold cash

[`money market`](/docs/api/accounts/#StandaloneAccountType-depository-money%20market)

stringstring

Money market account

[`paypal`](/docs/api/accounts/#StandaloneAccountType-depository-paypal)

stringstring

PayPal depository account

[`prepaid`](/docs/api/accounts/#StandaloneAccountType-depository-prepaid)

stringstring

Prepaid debit card

[`savings`](/docs/api/accounts/#StandaloneAccountType-depository-savings)

stringstring

Savings account

[`credit`](/docs/api/accounts/#StandaloneAccountType-credit)

stringstring

A credit card type account.

[`credit card`](/docs/api/accounts/#StandaloneAccountType-credit-credit%20card)

stringstring

Bank-issued credit card

[`paypal`](/docs/api/accounts/#StandaloneAccountType-credit-paypal)

stringstring

PayPal-issued credit card

[`loan`](/docs/api/accounts/#StandaloneAccountType-loan)

stringstring

A loan type account.

[`auto`](/docs/api/accounts/#StandaloneAccountType-loan-auto)

stringstring

Auto loan

[`business`](/docs/api/accounts/#StandaloneAccountType-loan-business)

stringstring

Business loan

[`commercial`](/docs/api/accounts/#StandaloneAccountType-loan-commercial)

stringstring

Commercial loan

[`construction`](/docs/api/accounts/#StandaloneAccountType-loan-construction)

stringstring

Construction loan

[`consumer`](/docs/api/accounts/#StandaloneAccountType-loan-consumer)

stringstring

Consumer loan

[`home equity`](/docs/api/accounts/#StandaloneAccountType-loan-home%20equity)

stringstring

Home Equity Line of Credit (HELOC)

[`line of credit`](/docs/api/accounts/#StandaloneAccountType-loan-line%20of%20credit)

stringstring

Pre-approved line of credit

[`loan`](/docs/api/accounts/#StandaloneAccountType-loan-loan)

stringstring

General loan

[`mortgage`](/docs/api/accounts/#StandaloneAccountType-loan-mortgage)

stringstring

Mortgage loan

[`other`](/docs/api/accounts/#StandaloneAccountType-loan-other)

stringstring

Other loan type or unknown loan type

[`overdraft`](/docs/api/accounts/#StandaloneAccountType-loan-overdraft)

stringstring

Pre-approved overdraft account, usually tied to a checking account

[`student`](/docs/api/accounts/#StandaloneAccountType-loan-student)

stringstring

Student loan

[`investment`](/docs/api/accounts/#StandaloneAccountType-investment)

stringstring

An investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage`.

[`529`](/docs/api/accounts/#StandaloneAccountType-investment-529)

stringstring

Tax-advantaged college savings and prepaid tuition 529 plans (US)

[`401a`](/docs/api/accounts/#StandaloneAccountType-investment-401a)

stringstring

Employer-sponsored money-purchase 401(a) retirement plan (US)

[`401k`](/docs/api/accounts/#StandaloneAccountType-investment-401k)

stringstring

Standard 401(k) retirement account (US)

[`403B`](/docs/api/accounts/#StandaloneAccountType-investment-403B)

stringstring

403(b) retirement savings account for non-profits and schools (US)

[`457b`](/docs/api/accounts/#StandaloneAccountType-investment-457b)

stringstring

Tax-advantaged deferred-compensation 457(b) retirement plan for governments and non-profits (US)

[`brokerage`](/docs/api/accounts/#StandaloneAccountType-investment-brokerage)

stringstring

Standard brokerage account

[`cash isa`](/docs/api/accounts/#StandaloneAccountType-investment-cash%20isa)

stringstring

Individual Savings Account (ISA) that pays interest tax-free (UK)

[`crypto exchange`](/docs/api/accounts/#StandaloneAccountType-investment-crypto%20exchange)

stringstring

Standard cryptocurrency exchange account

[`education savings account`](/docs/api/accounts/#StandaloneAccountType-investment-education%20savings%20account)

stringstring

Tax-advantaged Coverdell Education Savings Account (ESA) (US)

[`fixed annuity`](/docs/api/accounts/#StandaloneAccountType-investment-fixed%20annuity)

stringstring

Fixed annuity

[`gic`](/docs/api/accounts/#StandaloneAccountType-investment-gic)

stringstring

Guaranteed Investment Certificate (Canada)

[`health reimbursement arrangement`](/docs/api/accounts/#StandaloneAccountType-investment-health%20reimbursement%20arrangement)

stringstring

Tax-advantaged Health Reimbursement Arrangement (HRA) benefit plan (US)

[`hsa`](/docs/api/accounts/#StandaloneAccountType-investment-hsa)

stringstring

Non-cash tax-advantaged medical Health Savings Account (HSA) (US)

[`ira`](/docs/api/accounts/#StandaloneAccountType-investment-ira)

stringstring

Traditional Individual Retirement Account (IRA) (US)

[`isa`](/docs/api/accounts/#StandaloneAccountType-investment-isa)

stringstring

Non-cash Individual Savings Account (ISA) (UK)

[`keogh`](/docs/api/accounts/#StandaloneAccountType-investment-keogh)

stringstring

Keogh self-employed retirement plan (US)

[`lif`](/docs/api/accounts/#StandaloneAccountType-investment-lif)

stringstring

Life Income Fund (LIF) retirement account (Canada)

[`life insurance`](/docs/api/accounts/#StandaloneAccountType-investment-life%20insurance)

stringstring

Life insurance account

[`lira`](/docs/api/accounts/#StandaloneAccountType-investment-lira)

stringstring

Locked-in Retirement Account (LIRA) (Canada)

[`lrif`](/docs/api/accounts/#StandaloneAccountType-investment-lrif)

stringstring

Locked-in Retirement Income Fund (LRIF) (Canada)

[`lrsp`](/docs/api/accounts/#StandaloneAccountType-investment-lrsp)

stringstring

Locked-in Retirement Savings Plan (Canada)

[`mutual fund`](/docs/api/accounts/#StandaloneAccountType-investment-mutual%20fund)

stringstring

Mutual fund account

[`non-custodial wallet`](/docs/api/accounts/#StandaloneAccountType-investment-non-custodial%20wallet)

stringstring

A cryptocurrency wallet where the user controls the private key

[`non-taxable brokerage account`](/docs/api/accounts/#StandaloneAccountType-investment-non-taxable%20brokerage%20account)

stringstring

A non-taxable brokerage account that is not covered by a more specific subtype

[`other`](/docs/api/accounts/#StandaloneAccountType-investment-other)

stringstring

An account whose type could not be determined

[`other annuity`](/docs/api/accounts/#StandaloneAccountType-investment-other%20annuity)

stringstring

An annuity account not covered by other subtypes

[`other insurance`](/docs/api/accounts/#StandaloneAccountType-investment-other%20insurance)

stringstring

An insurance account not covered by other subtypes

[`pension`](/docs/api/accounts/#StandaloneAccountType-investment-pension)

stringstring

Standard pension account

[`prif`](/docs/api/accounts/#StandaloneAccountType-investment-prif)

stringstring

Prescribed Registered Retirement Income Fund (Canada)

[`profit sharing plan`](/docs/api/accounts/#StandaloneAccountType-investment-profit%20sharing%20plan)

stringstring

Plan that gives employees share of company profits

[`qshr`](/docs/api/accounts/#StandaloneAccountType-investment-qshr)

stringstring

Qualifying share account

[`rdsp`](/docs/api/accounts/#StandaloneAccountType-investment-rdsp)

stringstring

Registered Disability Savings Plan (RSDP) (Canada)

[`resp`](/docs/api/accounts/#StandaloneAccountType-investment-resp)

stringstring

Registered Education Savings Plan (Canada)

[`retirement`](/docs/api/accounts/#StandaloneAccountType-investment-retirement)

stringstring

Retirement account not covered by other subtypes

[`rlif`](/docs/api/accounts/#StandaloneAccountType-investment-rlif)

stringstring

Restricted Life Income Fund (RLIF) (Canada)

[`roth`](/docs/api/accounts/#StandaloneAccountType-investment-roth)

stringstring

Roth IRA (US)

[`roth 401k`](/docs/api/accounts/#StandaloneAccountType-investment-roth%20401k)

stringstring

Employer-sponsored Roth 401(k) plan (US)

[`rrif`](/docs/api/accounts/#StandaloneAccountType-investment-rrif)

stringstring

Registered Retirement Income Fund (RRIF) (Canada)

[`rrsp`](/docs/api/accounts/#StandaloneAccountType-investment-rrsp)

stringstring

Registered Retirement Savings Plan (Canadian, similar to US 401(k))

[`sarsep`](/docs/api/accounts/#StandaloneAccountType-investment-sarsep)

stringstring

Salary Reduction Simplified Employee Pension Plan (SARSEP), discontinued retirement plan (US)

[`sep ira`](/docs/api/accounts/#StandaloneAccountType-investment-sep%20ira)

stringstring

Simplified Employee Pension IRA (SEP IRA), retirement plan for small businesses and self-employed (US)

[`simple ira`](/docs/api/accounts/#StandaloneAccountType-investment-simple%20ira)

stringstring

Savings Incentive Match Plan for Employees IRA, retirement plan for small businesses (US)

[`sipp`](/docs/api/accounts/#StandaloneAccountType-investment-sipp)

stringstring

Self-Invested Personal Pension (SIPP) (UK)

[`stock plan`](/docs/api/accounts/#StandaloneAccountType-investment-stock%20plan)

stringstring

Standard stock plan account

[`tfsa`](/docs/api/accounts/#StandaloneAccountType-investment-tfsa)

stringstring

Tax-Free Savings Account (TFSA), a retirement plan similar to a Roth IRA (Canada)

[`thrift savings plan`](/docs/api/accounts/#StandaloneAccountType-investment-thrift%20savings%20plan)

stringstring

Thrift Savings Plan, a retirement savings and investment plan for Federal employees and members of the uniformed services.

[`trust`](/docs/api/accounts/#StandaloneAccountType-investment-trust)

stringstring

Account representing funds or assets held by a trustee for the benefit of a beneficiary. Includes both revocable and irrevocable trusts.

[`ugma`](/docs/api/accounts/#StandaloneAccountType-investment-ugma)

stringstring

'Uniform Gift to Minors Act' (brokerage account for minors, US)

[`utma`](/docs/api/accounts/#StandaloneAccountType-investment-utma)

stringstring

'Uniform Transfers to Minors Act' (brokerage account for minors, US)

[`variable annuity`](/docs/api/accounts/#StandaloneAccountType-investment-variable%20annuity)

stringstring

Tax-deferred capital accumulation annuity contract

[`payroll`](/docs/api/accounts/#StandaloneAccountType-payroll)

stringstring

A payroll account.

[`payroll`](/docs/api/accounts/#StandaloneAccountType-payroll-payroll)

stringstring

Standard payroll account

[`other`](/docs/api/accounts/#StandaloneAccountType-other)

stringstring

Other or unknown account type.

=\*=\*=\*=

#### Account type / product support matrix

The chart below indicates which products can be used with which account types. Note that some products can only be used with certain subtypes:

- Auth and Signal require a debitable `checking`, `savings`, or `cash management` account.
- Liabilities does not support `loan` types other than `student` or `mortgage`.
- Transactions does not support `loan` types other than `student`. (For Canadian institutions, Transactions also supports the `mortgage` loan type.)
- Investments does not support `depository` types other than `cash management`.

Also note that not all institutions support all products; for details on which products a given institution supports, use [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) or look up the institution on the [Plaid Dashboard status page](https://dashboard.plaid.com/activity/status) or the [Coverage Explorer](/docs/institutions/).

| Product | Depository | Credit | Investments | Loan | Other |
| --- | --- | --- | --- | --- | --- |
| [Balance](/docs/balance/) |  |  | \* |  |  |
| [Transactions](/docs/transactions/) |  |  |  |  |  |
| [Identity](/docs/identity/) |  |  |  |  |  |
| [Assets](/docs/assets/) |  |  |  |  |  |
| [Consumer Reports by Plaid Check](/docs/check/) |  |  |  |  |  |
| [Investments](/docs/investments/) |  |  |  |  |  |
| [Investments Move](/docs/investments/) |  |  |  |  |  |
| [Liabilities](/docs/liabilities/) |  |  |  |  |  |
| [Auth](/docs/auth/) |  |  |  |  |  |
| [Transfer](/docs/transfer/) |  |  |  |  |  |
| [Income (Bank Income flow)](/docs/income/) |  |  |  |  |  |
| [Statements](/docs/statements/) |  |  |  |  |  |
| [Signal](/docs/signal/) |  |  |  |  |  |
| [Payment Initiation (UK and Europe)](/docs/payment-initiation/) |  |  |  |  |  |
| [Virtual Accounts (UK and Europe)](/docs/payment-initiation/) |  |  |  |  |  |

\* Investments holdings data is not priced intra-day.

=\*=\*=\*=

#### Currency code schema

The following currency codes are supported by Plaid.

**Properties**

[`iso_currency_code`](/docs/api/accounts/#StandaloneCurrencyCodeList-iso-currency-code)

stringstring

Plaid supports all ISO 4217 currency codes.

[`unofficial_currency_code`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code)

stringstring

List of unofficial currency codes

[`ADA`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-ADA)

stringstring

Cardano

[`BAT`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-BAT)

stringstring

Basic Attention Token

[`BCH`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-BCH)

stringstring

Bitcoin Cash

[`BNB`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-BNB)

stringstring

Binance Coin

[`BTC`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-BTC)

stringstring

Bitcoin

[`BTG`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-BTG)

stringstring

Bitcoin Gold

[`BSV`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-BSV)

stringstring

Bitcoin Satoshi Vision

[`CNH`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-CNH)

stringstring

Chinese Yuan (offshore)

[`DASH`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-DASH)

stringstring

Dash

[`DOGE`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-DOGE)

stringstring

Dogecoin

[`ETC`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-ETC)

stringstring

Ethereum Classic

[`ETH`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-ETH)

stringstring

Ethereum

[`GBX`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-GBX)

stringstring

Pence sterling, i.e. British penny

[`LSK`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-LSK)

stringstring

Lisk

[`NEO`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-NEO)

stringstring

Neo

[`OMG`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-OMG)

stringstring

OmiseGO

[`QTUM`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-QTUM)

stringstring

Qtum

[`USDT`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-USDT)

stringstring

Tether

[`XLM`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-XLM)

stringstring

Stellar Lumen

[`XMR`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-XMR)

stringstring

Monero

[`XRP`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-XRP)

stringstring

Ripple

[`ZEC`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-ZEC)

stringstring

Zcash

[`ZRX`](/docs/api/accounts/#StandaloneCurrencyCodeList-unofficial-currency-code-ZRX)

stringstring

0x

=\*=\*=\*=

#### Investment transaction types schema

Valid values for investment transaction types and subtypes. Note that transactions representing inflow of cash will appear as negative amounts, outflow of cash will appear as positive amounts.

**Properties**

[`buy`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy)

stringstring

Buying an investment

[`assignment`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy-assignment)

stringstring

Assignment of short option holding

[`contribution`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy-contribution)

stringstring

Inflow of assets into a tax-advantaged account

[`buy`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy-buy)

stringstring

Purchase to open or increase a position

[`buy to cover`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy-buy%20to%20cover)

stringstring

Purchase to close a short position

[`dividend reinvestment`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy-dividend%20reinvestment)

stringstring

Purchase using proceeds from a cash dividend

[`interest reinvestment`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy-interest%20reinvestment)

stringstring

Purchase using proceeds from a cash interest payment

[`long-term capital gain reinvestment`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy-long-term%20capital%20gain%20reinvestment)

stringstring

Purchase using long-term capital gain cash proceeds

[`short-term capital gain reinvestment`](/docs/api/accounts/#StandaloneInvestmentTransactionType-buy-short-term%20capital%20gain%20reinvestment)

stringstring

Purchase using short-term capital gain cash proceeds

[`sell`](/docs/api/accounts/#StandaloneInvestmentTransactionType-sell)

stringstring

Selling an investment

[`distribution`](/docs/api/accounts/#StandaloneInvestmentTransactionType-sell-distribution)

stringstring

Outflow of assets from a tax-advantaged account

[`exercise`](/docs/api/accounts/#StandaloneInvestmentTransactionType-sell-exercise)

stringstring

Exercise of an option or warrant contract

[`sell`](/docs/api/accounts/#StandaloneInvestmentTransactionType-sell-sell)

stringstring

Sell to close or decrease an existing holding

[`sell short`](/docs/api/accounts/#StandaloneInvestmentTransactionType-sell-sell%20short)

stringstring

Sell to open a short position

[`cancel`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cancel)

stringstring

A cancellation of a pending transaction

[`cash`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash)

stringstring

Activity that modifies a cash position

[`account fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-account%20fee)

stringstring

Fees paid for account maintenance

[`contribution`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-contribution)

stringstring

Inflow of assets into a tax-advantaged account

[`deposit`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-deposit)

stringstring

Inflow of cash into an account

[`dividend`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-dividend)

stringstring

Inflow of cash from a dividend

[`stock distribution`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-stock%20distribution)

stringstring

Inflow of stock from a distribution

[`interest`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-interest)

stringstring

Inflow of cash from interest

[`legal fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-legal%20fee)

stringstring

Fees paid for legal charges or services

[`long-term capital gain`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-long-term%20capital%20gain)

stringstring

Long-term capital gain received as cash

[`management fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-management%20fee)

stringstring

Fees paid for investment management of a mutual fund or other pooled investment vehicle

[`margin expense`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-margin%20expense)

stringstring

Fees paid for maintaining margin debt

[`non-qualified dividend`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-non-qualified%20dividend)

stringstring

Inflow of cash from a non-qualified dividend

[`non-resident tax`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-non-resident%20tax)

stringstring

Taxes paid on behalf of the investor for non-residency in investment jurisdiction

[`pending credit`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-pending%20credit)

stringstring

Pending inflow of cash

[`pending debit`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-pending%20debit)

stringstring

Pending outflow of cash

[`qualified dividend`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-qualified%20dividend)

stringstring

Inflow of cash from a qualified dividend

[`short-term capital gain`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-short-term%20capital%20gain)

stringstring

Short-term capital gain received as cash

[`tax`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-tax)

stringstring

Taxes paid on behalf of the investor

[`tax withheld`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-tax%20withheld)

stringstring

Taxes withheld on behalf of the customer

[`transfer fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-transfer%20fee)

stringstring

Fees incurred for transfer of a holding or account

[`trust fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-trust%20fee)

stringstring

Fees related to administration of a trust account

[`unqualified gain`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-unqualified%20gain)

stringstring

Unqualified capital gain received as cash

[`withdrawal`](/docs/api/accounts/#StandaloneInvestmentTransactionType-cash-withdrawal)

stringstring

Outflow of cash from an account

[`fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee)

stringstring

Fees on the account, e.g. commission, bookkeeping, options-related.

[`account fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-account%20fee)

stringstring

Fees paid for account maintenance

[`adjustment`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-adjustment)

stringstring

Increase or decrease in quantity of item

[`dividend`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-dividend)

stringstring

Inflow of cash from a dividend

[`interest`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-interest)

stringstring

Inflow of cash from interest

[`interest receivable`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-interest%20receivable)

stringstring

Inflow of cash from interest receivable

[`long-term capital gain`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-long-term%20capital%20gain)

stringstring

Long-term capital gain received as cash

[`legal fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-legal%20fee)

stringstring

Fees paid for legal charges or services

[`management fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-management%20fee)

stringstring

Fees paid for investment management of a mutual fund or other pooled investment vehicle

[`margin expense`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-margin%20expense)

stringstring

Fees paid for maintaining margin debt

[`non-qualified dividend`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-non-qualified%20dividend)

stringstring

Inflow of cash from a non-qualified dividend

[`non-resident tax`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-non-resident%20tax)

stringstring

Taxes paid on behalf of the investor for non-residency in investment jurisdiction

[`qualified dividend`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-qualified%20dividend)

stringstring

Inflow of cash from a qualified dividend

[`return of principal`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-return%20of%20principal)

stringstring

Repayment of loan principal

[`short-term capital gain`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-short-term%20capital%20gain)

stringstring

Short-term capital gain received as cash

[`stock distribution`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-stock%20distribution)

stringstring

Inflow of stock from a distribution

[`tax`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-tax)

stringstring

Taxes paid on behalf of the investor

[`tax withheld`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-tax%20withheld)

stringstring

Taxes withheld on behalf of the customer

[`transfer fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-transfer%20fee)

stringstring

Fees incurred for transfer of a holding or account

[`trust fee`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-trust%20fee)

stringstring

Fees related to administration of a trust account

[`unqualified gain`](/docs/api/accounts/#StandaloneInvestmentTransactionType-fee-unqualified%20gain)

stringstring

Unqualified capital gain received as cash

[`transfer`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer)

stringstring

Activity that modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer

[`assignment`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-assignment)

stringstring

Assignment of short option holding

[`adjustment`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-adjustment)

stringstring

Increase or decrease in quantity of item

[`exercise`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-exercise)

stringstring

Exercise of an option or warrant contract

[`expire`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-expire)

stringstring

Expiration of an option or warrant contract

[`merger`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-merger)

stringstring

Stock exchanged at a pre-defined ratio as part of a merger between companies

[`request`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-request)

stringstring

Request fiat or cryptocurrency to an address or email

[`send`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-send)

stringstring

Inflow or outflow of fiat or cryptocurrency to an address or email

[`spin off`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-spin%20off)

stringstring

Inflow of stock from spin-off transaction of an existing holding

[`split`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-split)

stringstring

Inflow of stock from a forward split of an existing holding

[`trade`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-trade)

stringstring

Trade of one cryptocurrency for another

[`transfer`](/docs/api/accounts/#StandaloneInvestmentTransactionType-transfer-transfer)

stringstring

Movement of assets into or out of an account

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
