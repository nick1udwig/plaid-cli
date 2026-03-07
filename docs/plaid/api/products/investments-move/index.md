---
title: "API - Investments Move | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/investments-move/"
scraped_at: "2026-03-07T22:04:13+00:00"
---

# Investments Move

#### API reference for Investments Move endpoints and webhooks

For how-to guidance, see the [Investments Move documentation](/docs/investments-move/).

| Endpoints |  |
| --- | --- |
| [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) | Fetch investments data required for ACATS or ATON transfer |

=\*=\*=\*=

#### `/investments/auth/get`

#### Get data needed to authorize an investments transfer

The [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) endpoint allows developers to receive user-authorized data to facilitate the transfer of holdings

/investments/auth/get

**Request fields**

[`client_id`](/docs/api/products/investments-move/#investments-auth-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/investments-move/#investments-auth-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`access_token`](/docs/api/products/investments-move/#investments-auth-get-request-access-token)

requiredstringrequired, string

The access token associated with the Item data is being requested for.

[`options`](/docs/api/products/investments-move/#investments-auth-get-request-options)

objectobject

An optional object to filter `/investments/auth/get` results.

[`account_ids`](/docs/api/products/investments-move/#investments-auth-get-request-options-account-ids)

[string][string]

An array of `account_id`s to retrieve for the Item. An error will be returned if a provided `account_id` is not associated with the Item.

/investments/auth/get

```
const request: InvestmentsAuthGetRequest = {
  access_token: accessToken,
};
try {
  const response = await plaidClient.investmentsAuthGet(request);
  const investmentsAuthData = response.data;
} catch (error) {
  // handle error
}
```

/investments/auth/get

**Response fields**

[`accounts`](/docs/api/products/investments-move/#investments-auth-get-response-accounts)

[object][object]

The accounts for which data is being retrieved

[`account_id`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  
The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  
If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  
When using a CRA endpoint (an endpoint associated with Plaid Check Consumer Report, i.e. any endpoint beginning with `/cra/`), the `account_id` returned will not match the `account_id` returned by a non-CRA endpoint.  
Like all Plaid identifiers, the `account_id` is case sensitive.

[`balances`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-balances)

objectobject

A set of fields describing the balance for an account. Balance information may be cached unless the balance object was returned by `/accounts/balance/get` or `/signal/evaluate` (using a Balance-only ruleset).

[`available`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-balances-available)

nullablenumbernullable, number

The amount of funds available to be withdrawn from the account, as determined by the financial institution.  
For `credit`-type accounts, the `available` balance typically equals the `limit` less the `current` balance, less any pending outflows plus any pending inflows.  
For `depository`-type accounts, the `available` balance typically equals the `current` balance less any pending outflows plus any pending inflows. For `depository`-type accounts, the `available` balance does not include the overdraft limit.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the `available` balance is the total cash available to withdraw as presented by the institution.  
Note that not all institutions calculate the `available` balance. In the event that `available` balance is unavailable, Plaid will return an `available` balance value of `null`.  
Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by `/accounts/balance/get`, or by `/signal/evaluate` with a Balance-only ruleset.  
If `current` is `null` this field is guaranteed not to be `null`.  
  

Format: `double`

[`current`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-balances-current)

nullablenumbernullable, number

The total amount of funds in or owed by the account.  
For `credit`-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  
For `loan`-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (`ins_116944`). For Sallie Mae student loans, the account's balance includes both principal and any outstanding interest. Similar to `credit`-type accounts, a positive balance is typically expected, while a negative amount indicates the lender owing the account holder.  
For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the current balance is the total value of assets as presented by the institution.  
Note that balance information may be cached unless the value was returned by `/accounts/balance/get` or by `/signal/evaluate` with a Balance-only ruleset; if the Item is enabled for Transactions, the balance will be at least as recent as the most recent Transaction update. If you require realtime balance information, use the `available` balance as provided by `/accounts/balance/get` or `/signal/evaluate` called with a Balance-only `ruleset_key`.  
When returned by `/accounts/balance/get`, this field may be `null`. When this happens, `available` is guaranteed not to be `null`.  
  

Format: `double`

[`limit`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-balances-limit)

nullablenumbernullable, number

For `credit`-type accounts, this represents the credit limit.  
For `depository`-type accounts, this represents the pre-arranged overdraft limit, which is common for current (checking) accounts in Europe.  
In North America, this field is typically only available for `credit`-type accounts.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-balances-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the balance. Always null if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-balances-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the balance. Always null if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.

[`last_updated_datetime`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-balances-last-updated-datetime)

nullablestringnullable, string

Timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the last time the balance was updated.  
This field is returned only when the institution is `ins_128026` (Capital One).  
  

Format: `date-time`

[`mask`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of either the account’s displayed mask or the account’s official account number. Note that the mask may be non-unique between an Item’s accounts.

[`name`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself

[`official_name`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-official-name)

nullablestringnullable, string

The official name of the account as given by the financial institution

[`type`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`verification_status`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-status)

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

[`verification_name`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-name)

stringstring

The account holder name that was used for micro-deposit and/or database verification. Only returned for Auth Items created via micro-deposit or database verification. This name was manually-entered by the user during Link, unless it was otherwise provided via the `user.legal_name` request field in `/link/token/create` for the Link session that created the Item.

[`verification_insights`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-insights)

objectobject

Insights from performing database verification for the account. Only returned for Auth Items using Database Auth.

[`name_match_score`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-insights-name-match-score)

nullableintegernullable, integer

Indicates the score of the name match between the given name provided during database verification (available in the [`verification_name`](https://plaid.com/docs/api/products/auth/#auth-get-response-accounts-verification-name) field if using standard Database Auth, or provided in the request if using `/auth/verify`) and matched Plaid network accounts. If defined, will be a value between 0 and 100. Will be undefined if name matching was not enabled for the database verification session or if there were no eligible Plaid network matches to compare the given name with.

[`network_status`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-insights-network-status)

objectobject

Status information about the account and routing number in the Plaid network.

[`has_numbers_match`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-insights-network-status-has-numbers-match)

booleanboolean

Indicates whether we found at least one matching account for the ACH account and routing number.

[`is_numbers_match_verified`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-insights-network-status-is-numbers-match-verified)

booleanboolean

Indicates if at least one matching account for the ACH account and routing number is already verified.

[`previous_returns`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-insights-previous-returns)

objectobject

Information about known ACH returns for the account and routing number.

[`has_previous_administrative_return`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-insights-previous-returns-has-previous-administrative-return)

booleanboolean

Indicates whether Plaid's data sources include a known administrative ACH return for this account and routing number.

[`account_number_format`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-verification-insights-account-number-format)

stringstring

Indicator of account number format validity for institution.  
`valid`: indicates that the account number has a correct format for the institution.  
`invalid`: indicates that the account number has an incorrect format for the institution.  
`unknown`: indicates that there was not enough information to determine whether the format is correct for the institution.  
  

Possible values: `valid`, `invalid`, `unknown`

[`persistent_account_id`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-persistent-account-id)

stringstring

A unique and persistent identifier for accounts that can be used to trace multiple instances of the same account across different Items for depository accounts. This field is currently supported only for Items at institutions that use Tokenized Account Numbers (i.e., Chase and PNC, and in May 2025 US Bank). Because these accounts have a different account number each time they are linked, this field may be used instead of the account number to uniquely identify an account across multiple Items for payments use cases, helping to reduce duplicate Items or attempted fraud. In Sandbox, this field is populated for TAN-based institutions (`ins_56`, `ins_13`) as well as the OAuth Sandbox institution (`ins_127287`); in Production, it will only be populated for accounts at applicable institutions.

[`holder_category`](/docs/api/products/investments-move/#investments-auth-get-response-accounts-holder-category)

nullablestringnullable, string

Indicates the account's categorization as either a personal or a business account. This field is currently in beta; to request access, contact your account manager.  
  

Possible values: `business`, `personal`, `unrecognized`

[`holdings`](/docs/api/products/investments-move/#investments-auth-get-response-holdings)

[object][object]

The holdings belonging to investment accounts associated with the Item. Details of the securities in the holdings are provided in the `securities` field.

[`account_id`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-account-id)

stringstring

The Plaid `account_id` associated with the holding.

[`security_id`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-security-id)

stringstring

The Plaid `security_id` associated with the holding. Security data is not specific to a user's account; any user who held the same security at the same financial institution at the same time would have identical security data. The `security_id` for the same security will typically be the same across different institutions, but this is not guaranteed. The `security_id` does not typically change, but may change if inherent details of the security change due to a corporate action, for example, in the event of a ticker symbol change or CUSIP change.

[`institution_price`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-institution-price)

numbernumber

The last price given by the institution for this security.  
  

Format: `double`

[`institution_price_as_of`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-institution-price-as-of)

nullablestringnullable, string

The date at which `institution_price` was current.  
  

Format: `date`

[`institution_price_datetime`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-institution-price-datetime)

nullablestringnullable, string

Date and time at which `institution_price` was current, in ISO 8601 format (YYYY-MM-DDTHH:mm:ssZ).  
This field is returned for select financial institutions and comes as provided by the institution. It may contain default time values (such as 00:00:00).  
  

Format: `date-time`

[`institution_value`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-institution-value)

numbernumber

The value of the holding, as reported by the institution.  
  

Format: `double`

[`cost_basis`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-cost-basis)

nullablenumbernullable, number

The total cost basis of the holding (e.g., the total amount spent to acquire all assets currently in the holding).  
  

Format: `double`

[`quantity`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-quantity)

numbernumber

The total quantity of the asset held, as reported by the financial institution. If the security is an option, `quantity` will reflect the total number of options (typically the number of contracts multiplied by 100), not the number of contracts.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the holding. Always `null` if `unofficial_currency_code` is non-`null`.

[`unofficial_currency_code`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the holding. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`vested_quantity`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-vested-quantity)

nullablenumbernullable, number

The total quantity of vested assets held, as reported by the financial institution. Vested assets are only associated with [equities](https://plaid.com/docs/api/products/investments/#investments-holdings-get-response-securities-type).  
  

Format: `double`

[`vested_value`](/docs/api/products/investments-move/#investments-auth-get-response-holdings-vested-value)

nullablenumbernullable, number

The value of the vested holdings as reported by the institution.  
  

Format: `double`

[`securities`](/docs/api/products/investments-move/#investments-auth-get-response-securities)

[object][object]

Objects describing the securities held in the accounts associated with the Item.

[`security_id`](/docs/api/products/investments-move/#investments-auth-get-response-securities-security-id)

stringstring

A unique, Plaid-specific identifier for the security, used to associate securities with holdings. Like all Plaid identifiers, the `security_id` is case sensitive. The `security_id` may change if inherent details of the security change due to a corporate action, for example, in the event of a ticker symbol change or CUSIP change.

[`isin`](/docs/api/products/investments-move/#investments-auth-get-response-securities-isin)

nullablestringnullable, string

12-character ISIN, a globally unique securities identifier. A verified CUSIP Global Services license is required to receive this data. This field will be null by default for new customers, and null for existing customers starting March 12, 2024. If you would like access to this field, please start the verification process [here](https://docs.google.com/forms/d/e/1FAIpQLSd9asHEYEfmf8fxJTHZTAfAzW4dugsnSu-HS2J51f1mxwd6Sw/viewform).

[`cusip`](/docs/api/products/investments-move/#investments-auth-get-response-securities-cusip)

nullablestringnullable, string

9-character CUSIP, an identifier assigned to North American securities. A verified CUSIP Global Services license is required to receive this data. This field will be null by default for new customers, and null for existing customers starting March 12, 2024. If you would like access to this field, please start the verification process [here](https://docs.google.com/forms/d/e/1FAIpQLSd9asHEYEfmf8fxJTHZTAfAzW4dugsnSu-HS2J51f1mxwd6Sw/viewform).

[`sedol`](/docs/api/products/investments-move/#investments-auth-get-response-securities-sedol)

deprecatednullablestringdeprecated, nullable, string

(Deprecated) 7-character SEDOL, an identifier assigned to securities in the UK.

[`institution_security_id`](/docs/api/products/investments-move/#investments-auth-get-response-securities-institution-security-id)

nullablestringnullable, string

An identifier given to the security by the institution

[`institution_id`](/docs/api/products/investments-move/#investments-auth-get-response-securities-institution-id)

nullablestringnullable, string

If `institution_security_id` is present, this field indicates the Plaid `institution_id` of the institution to whom the identifier belongs.

[`proxy_security_id`](/docs/api/products/investments-move/#investments-auth-get-response-securities-proxy-security-id)

nullablestringnullable, string

In certain cases, Plaid will provide the ID of another security whose performance resembles this security, typically when the original security has low volume, or when a private security can be modeled with a publicly traded security.

[`name`](/docs/api/products/investments-move/#investments-auth-get-response-securities-name)

nullablestringnullable, string

A descriptive name for the security, suitable for display.

[`ticker_symbol`](/docs/api/products/investments-move/#investments-auth-get-response-securities-ticker-symbol)

nullablestringnullable, string

The security’s trading symbol for publicly traded securities, and otherwise a short identifier if available.

[`is_cash_equivalent`](/docs/api/products/investments-move/#investments-auth-get-response-securities-is-cash-equivalent)

nullablebooleannullable, boolean

Indicates that a security is a highly liquid asset and can be treated like cash.

[`type`](/docs/api/products/investments-move/#investments-auth-get-response-securities-type)

nullablestringnullable, string

The security type of the holding.  
In rare instances, a null value is returned when institutional data is insufficient to determine the security type.  
Valid security types are:  
`cash`: Cash, currency, and money market funds  
`cryptocurrency`: Digital or virtual currencies  
`derivative`: Options, warrants, and other derivative instruments  
`equity`: Domestic and foreign equities  
`etf`: Multi-asset exchange-traded investment funds  
`fixed income`: Bonds and certificates of deposit (CDs)  
`loan`: Loans and loan receivables  
`mutual fund`: Open- and closed-end vehicles pooling funds of multiple investors  
`other`: Unknown or other investment types

[`subtype`](/docs/api/products/investments-move/#investments-auth-get-response-securities-subtype)

nullablestringnullable, string

The security subtype of the holding.  
In rare instances, a null value is returned when institutional data is insufficient to determine the security subtype.  
Possible values: `asset backed security`, `bill`, `bond`, `bond with warrants`, `cash`, `cash management bill`, `common stock`, `convertible bond`, `convertible equity`, `cryptocurrency`, `depositary receipt`, `depositary receipt on debt`, `etf`, `float rating note`, `fund of funds`, `hedge fund`, `limited partnership unit`, `medium term note`, `money market debt`, `mortgage backed security`, `municipal bond`, `mutual fund`, `note`, `option`, `other`, `preferred convertible`, `preferred equity`, `private equity fund`, `real estate investment trust`, `structured equity product`, `treasury inflation protected securities`, `unit`, `warrant`.

[`close_price`](/docs/api/products/investments-move/#investments-auth-get-response-securities-close-price)

nullablenumbernullable, number

Price of the security at the close of the previous trading session. Null for non-public securities.  
If the security is a foreign currency this field will be updated daily and will be priced in USD.  
If the security is a cryptocurrency, this field will be updated multiple times a day. As crypto prices can fluctuate quickly and data may become stale sooner than other asset classes, refer to `update_datetime` with the time when the price was last updated.  
  

Format: `double`

[`close_price_as_of`](/docs/api/products/investments-move/#investments-auth-get-response-securities-close-price-as-of)

nullablestringnullable, string

Date for which `close_price` is accurate. Always `null` if `close_price` is `null`.  
  

Format: `date`

[`update_datetime`](/docs/api/products/investments-move/#investments-auth-get-response-securities-update-datetime)

nullablestringnullable, string

Date and time at which `close_price` is accurate, in ISO 8601 format (YYYY-MM-DDTHH:mm:ssZ). Always `null` if `close_price` is `null`.  
  

Format: `date-time`

[`iso_currency_code`](/docs/api/products/investments-move/#investments-auth-get-response-securities-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the price given. Always `null` if `unofficial_currency_code` is non-`null`.

[`unofficial_currency_code`](/docs/api/products/investments-move/#investments-auth-get-response-securities-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the security. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`market_identifier_code`](/docs/api/products/investments-move/#investments-auth-get-response-securities-market-identifier-code)

nullablestringnullable, string

The ISO-10383 Market Identifier Code of the exchange or market in which the security is being traded.

[`sector`](/docs/api/products/investments-move/#investments-auth-get-response-securities-sector)

nullablestringnullable, string

The sector classification of the security, such as Finance, Health Technology, etc.  
For a complete list of possible values, please refer to the ["Sectors and Industries" spreadsheet](https://docs.google.com/spreadsheets/d/1L7aXUdqLhxgM8qe7hK67qqKXiUdQqILpwZ0LpxvCVnc).

[`industry`](/docs/api/products/investments-move/#investments-auth-get-response-securities-industry)

nullablestringnullable, string

The industry classification of the security, such as Biotechnology, Airlines, etc.  
For a complete list of possible values, please refer to the ["Sectors and Industries" spreadsheet](https://docs.google.com/spreadsheets/d/1L7aXUdqLhxgM8qe7hK67qqKXiUdQqILpwZ0LpxvCVnc).

[`option_contract`](/docs/api/products/investments-move/#investments-auth-get-response-securities-option-contract)

nullableobjectnullable, object

Details about the option security.  
For the Sandbox environment, this data is currently only available if the Item is using a [custom Sandbox user](https://plaid.com/docs/sandbox/user-custom/) and the `ticker` field of the custom security follows the [OCC Option Symbol](https://en.wikipedia.org/wiki/Option_symbol#The_OCC_Option_Symbol) standard with no spaces. For an example of simulating this in Sandbox, see the [custom Sandbox GitHub](https://github.com/plaid/sandbox-custom-users).

[`contract_type`](/docs/api/products/investments-move/#investments-auth-get-response-securities-option-contract-contract-type)

stringstring

The type of this option contract. It is one of:  
`put`: for Put option contracts  
`call`: for Call option contracts

[`expiration_date`](/docs/api/products/investments-move/#investments-auth-get-response-securities-option-contract-expiration-date)

stringstring

The expiration date for this option contract, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date`

[`strike_price`](/docs/api/products/investments-move/#investments-auth-get-response-securities-option-contract-strike-price)

numbernumber

The strike price for this option contract, per share of security.  
  

Format: `double`

[`underlying_security_ticker`](/docs/api/products/investments-move/#investments-auth-get-response-securities-option-contract-underlying-security-ticker)

stringstring

The ticker of the underlying security for this option contract.

[`fixed_income`](/docs/api/products/investments-move/#investments-auth-get-response-securities-fixed-income)

nullableobjectnullable, object

Details about the fixed income security.

[`yield_rate`](/docs/api/products/investments-move/#investments-auth-get-response-securities-fixed-income-yield-rate)

nullableobjectnullable, object

Details about a fixed income security's expected rate of return.

[`percentage`](/docs/api/products/investments-move/#investments-auth-get-response-securities-fixed-income-yield-rate-percentage)

numbernumber

The fixed income security's expected rate of return.  
  

Format: `double`

[`type`](/docs/api/products/investments-move/#investments-auth-get-response-securities-fixed-income-yield-rate-type)

nullablestringnullable, string

The type of rate which indicates how the predicted yield was calculated. It is one of:  
`coupon`: the annualized interest rate for securities with a one-year term or longer, such as treasury notes and bonds.  
`coupon_equivalent`: the calculated equivalent for the annualized interest rate factoring in the discount rate and time to maturity, for shorter term, non-interest-bearing securities such as treasury bills.  
`discount`: the rate at which the present value or cost is discounted from the future value upon maturity, also known as the face value.  
`yield`: the total predicted rate of return factoring in both the discount rate and the coupon rate, applicable to securities such as exchange-traded bonds which can both be interest-bearing as well as sold at a discount off its face value.  
  

Possible values: `coupon`, `coupon_equivalent`, `discount`, `yield`, `null`

[`maturity_date`](/docs/api/products/investments-move/#investments-auth-get-response-securities-fixed-income-maturity-date)

nullablestringnullable, string

The maturity date for this fixed income security, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date`

[`issue_date`](/docs/api/products/investments-move/#investments-auth-get-response-securities-fixed-income-issue-date)

nullablestringnullable, string

The issue date for this fixed income security, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date`

[`face_value`](/docs/api/products/investments-move/#investments-auth-get-response-securities-fixed-income-face-value)

nullablenumbernullable, number

The face value that is paid upon maturity of the fixed income security, per unit of security.  
  

Format: `double`

[`owners`](/docs/api/products/investments-move/#investments-auth-get-response-owners)

[object][object]

Information about the account owners for the accounts associated with the Item.

[`account_id`](/docs/api/products/investments-move/#investments-auth-get-response-owners-account-id)

stringstring

The ID of the account that this identity information pertains to

[`names`](/docs/api/products/investments-move/#investments-auth-get-response-owners-names)

[string][string]

A list of names associated with the account by the financial institution. In the case of a joint account, Plaid will make a best effort to report the names of all account holders.  
If an Item contains multiple accounts with different owner names, some institutions will report all names associated with the Item in each account's `names` array.

[`numbers`](/docs/api/products/investments-move/#investments-auth-get-response-numbers)

objectobject

Identifying information for transferring holdings to an investments account.

[`acats`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-acats)

[object][object]

[`account_id`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-acats-account-id)

stringstring

The Plaid account ID associated with the account numbers

[`account`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-acats-account)

stringstring

The full account number for the account

[`dtc_numbers`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-acats-dtc-numbers)

[string][string]

Identifiers for the clearinghouses that are associated with the account in order of relevance. If this array is empty, call `/institutions/get_by_id` with the `item.institution_id` to get the DTC number.

[`aton`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-aton)

[object][object]

[`account_id`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-aton-account-id)

stringstring

The Plaid account ID associated with the account numbers

[`account`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-aton-account)

stringstring

The full account number for the account

[`retirement_401k`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-retirement-401k)

[object][object]

[`account_id`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-retirement-401k-account-id)

stringstring

The Plaid account ID associated with the account numbers

[`plan`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-retirement-401k-plan)

stringstring

The plan number for the employer's 401k retirement plan

[`account`](/docs/api/products/investments-move/#investments-auth-get-response-numbers-retirement-401k-account)

stringstring

The full account number for the account

[`data_sources`](/docs/api/products/investments-move/#investments-auth-get-response-data-sources)

objectobject

Object with metadata pertaining to the source of data for the account numbers, owners, and holdings that are returned.

[`numbers`](/docs/api/products/investments-move/#investments-auth-get-response-data-sources-numbers)

stringstring

A description of the source of data for a given product/data type.  
`INSTITUTION`: The institution supports this product, and the data was provided by the institution.
`INSTITUTION_MASK`: The user manually provided the full account number, which was matched to the account mask provided by the institution. Only applicable to the `numbers` data type.
`USER`: The institution does not support this product, and the data was manually provided by the user.  
  

Possible values: `INSTITUTION`, `INSTITUTION_MASK`, `USER`

[`owners`](/docs/api/products/investments-move/#investments-auth-get-response-data-sources-owners)

stringstring

A description of the source of data for a given product/data type.  
`INSTITUTION`: The institution supports this product, and the data was provided by the institution.
`INSTITUTION_MASK`: The user manually provided the full account number, which was matched to the account mask provided by the institution. Only applicable to the `numbers` data type.
`USER`: The institution does not support this product, and the data was manually provided by the user.  
  

Possible values: `INSTITUTION`, `INSTITUTION_MASK`, `USER`

[`holdings`](/docs/api/products/investments-move/#investments-auth-get-response-data-sources-holdings)

stringstring

A description of the source of data for a given product/data type.  
`INSTITUTION`: The institution supports this product, and the data was provided by the institution.
`INSTITUTION_MASK`: The user manually provided the full account number, which was matched to the account mask provided by the institution. Only applicable to the `numbers` data type.
`USER`: The institution does not support this product, and the data was manually provided by the user.  
  

Possible values: `INSTITUTION`, `INSTITUTION_MASK`, `USER`

[`item`](/docs/api/products/investments-move/#investments-auth-get-response-item)

objectobject

Metadata about the Item.

[`item_id`](/docs/api/products/investments-move/#investments-auth-get-response-item-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`institution_id`](/docs/api/products/investments-move/#investments-auth-get-response-item-institution-id)

nullablestringnullable, string

The Plaid Institution ID associated with the Item. Field is `null` for Items created without an institution connection, such as Items created via Same Day Micro-deposits.

[`institution_name`](/docs/api/products/investments-move/#investments-auth-get-response-item-institution-name)

nullablestringnullable, string

The name of the institution associated with the Item. Field is `null` for Items created without an institution connection, such as Items created via Same Day Micro-deposits.

[`webhook`](/docs/api/products/investments-move/#investments-auth-get-response-item-webhook)

nullablestringnullable, string

The URL registered to receive webhooks for the Item.

[`auth_method`](/docs/api/products/investments-move/#investments-auth-get-response-item-auth-method)

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

[`error`](/docs/api/products/investments-move/#investments-auth-get-response-item-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/investments-move/#investments-auth-get-response-item-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`available_products`](/docs/api/products/investments-move/#investments-auth-get-response-item-available-products)

[string][string]

A list of products available for the Item that have not yet been accessed. The contents of this array will be mutually exclusive with `billed_products`.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`billed_products`](/docs/api/products/investments-move/#investments-auth-get-response-item-billed-products)

[string][string]

A list of products that have been billed for the Item. The contents of this array will be mutually exclusive with `available_products`. Note - `billed_products` is populated in all environments but only requests in Production are billed. Also note that products that are billed on a pay-per-call basis rather than a pay-per-Item basis, such as `balance`, will not appear here.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`products`](/docs/api/products/investments-move/#investments-auth-get-response-item-products)

[string][string]

A list of products added to the Item. In almost all cases, this will be the same as the `billed_products` field. For some products, it is possible for the product to be added to an Item but not yet billed (e.g. Assets, before `/asset_report/create` has been called, or Auth or Identity when added as Optional Products but before their endpoints have been called), in which case the product may appear in `products` but not in `billed_products`.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`consented_products`](/docs/api/products/investments-move/#investments-auth-get-response-item-consented-products)

[string][string]

A list of products that the user has consented to for the Item via [Data Transparency Messaging](https://plaid.com/docs/link/data-transparency-messaging-migration-guide). This will consist of all products where both of the following are true: the user has consented to the required data scopes for that product and you have Production access for that product.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `transactions`, `income`, `income_verification`, `transfer`, `employment`, `recurring_transactions`, `signal`, `statements`, `processor_payments`, `processor_identity`, `cra_base_report`, `cra_income_insights`, `cra_lend_score`, `cra_partner_insights`, `cra_cashflow_insights`, `cra_monitoring`, `layer`

[`consent_expiration_time`](/docs/api/products/investments-move/#investments-auth-get-response-item-consent-expiration-time)

nullablestringnullable, string

The date and time at which the Item's access consent will expire, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. If the Item does not have consent expiration scheduled, this field will be `null`. Currently, only institutions in Europe and a small number of institutions in the US have expiring consent. For a list of US institutions that currently expire consent, see the [OAuth Guide](https://plaid.com/docs/link/oauth/#refreshing-item-consent).  
  

Format: `date-time`

[`update_type`](/docs/api/products/investments-move/#investments-auth-get-response-item-update-type)

stringstring

Indicates whether an Item requires user interaction to be updated, which can be the case for Items with some forms of two-factor authentication.  
`background` - Item can be updated in the background  
`user_present_required` - Item requires user interaction to be updated  
  

Possible values: `background`, `user_present_required`

[`request_id`](/docs/api/products/investments-move/#investments-auth-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "accounts": [
    {
      "account_id": "31qEA6LPwGumkA4Z5mGbfyGwr4mL6nSZlQqpZ",
      "balances": {
        "available": 43200,
        "current": 43200,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "4444",
      "name": "Plaid Money Market",
      "official_name": "Plaid Platinum Standard 1.85% Interest Money Market",
      "subtype": "money market",
      "type": "depository"
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "balances": {
        "available": null,
        "current": 415.57,
        "iso_currency_code": "USD",
        "limit": null,
        "unofficial_currency_code": null
      },
      "mask": "5555",
      "name": "Plaid IRA",
      "official_name": null,
      "subtype": "ira",
      "type": "investment"
    }
  ],
  "holdings": [
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "cost_basis": 1,
      "institution_price": 1,
      "institution_price_as_of": "2021-05-25",
      "institution_price_datetime": null,
      "institution_value": 0.01,
      "iso_currency_code": "USD",
      "quantity": 0.01,
      "security_id": "d6ePmbPxgWCWmMVv66q9iPV94n91vMtov5Are",
      "unofficial_currency_code": null,
      "vested_quantity": 1,
      "vested_value": 1
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "cost_basis": 0.01,
      "institution_price": 0.011,
      "institution_price_as_of": "2021-05-25",
      "institution_price_datetime": null,
      "institution_value": 110,
      "iso_currency_code": "USD",
      "quantity": 10000,
      "security_id": "8E4L9XLl6MudjEpwPAAgivmdZRdBPJuvMPlPb",
      "unofficial_currency_code": null,
      "vested_quantity": null,
      "vested_value": null
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "cost_basis": 94.808,
      "institution_price": 94.808,
      "institution_price_as_of": "2021-04-13",
      "institution_price_datetime": null,
      "institution_value": 94.808,
      "iso_currency_code": "USD",
      "quantity": 1,
      "security_id": "Lxe4yz4XQEtwb2YArO7RFMpPDvPxy7FALRyea",
      "unofficial_currency_code": null,
      "vested_quantity": null,
      "vested_value": null
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "cost_basis": 40,
      "institution_price": 42.15,
      "institution_price_as_of": "2021-05-25",
      "institution_price_datetime": null,
      "institution_value": 210.75,
      "iso_currency_code": "USD",
      "quantity": 5,
      "security_id": "abJamDazkgfvBkVGgnnLUWXoxnomp5up8llg4",
      "unofficial_currency_code": null,
      "vested_quantity": 7,
      "vested_value": 66
    }
  ],
  "item": {
    "available_products": [
      "assets",
      "balance",
      "beacon",
      "cra_base_report",
      "cra_income_insights",
      "signal",
      "identity",
      "identity_match",
      "income",
      "income_verification",
      "investments",
      "processor_identity",
      "recurring_transactions",
      "transactions"
    ],
    "billed_products": [
      "investments_auth"
    ],
    "consent_expiration_time": null,
    "error": null,
    "institution_id": "ins_115616",
    "institution_name": "Vanguard",
    "item_id": "7qBnDwLP3aIZkD7NKZ5ysk5X9xVxDWHg65oD5",
    "products": [
      "investments_auth"
    ],
    "update_type": "background",
    "webhook": "https://www.genericwebhookurl.com/webhook"
  },
  "numbers": {
    "acats": [
      {
        "account": "TR5555",
        "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
        "dtc_numbers": [
          "1111",
          "2222",
          "3333"
        ]
      }
    ]
  },
  "owners": [
    {
      "account_id": "31qEA6LPwGumkA4Z5mGbfyGwr4mL6nSZlQqpZ",
      "names": [
        "Alberta Bobbeth Charleson"
      ]
    },
    {
      "account_id": "xlP8npRxwgCj48LQbjxWipkeL3gbyXf64knoy",
      "names": [
        "Alberta Bobbeth Charleson"
      ]
    }
  ],
  "request_id": "hPCXou4mm9Qwzzu",
  "securities": [
    {
      "close_price": 0.011,
      "close_price_as_of": null,
      "cusip": null,
      "industry": null,
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": false,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "Nflx Feb 01'18 $355 Call",
      "option_contract": null,
      "fixed_income": null,
      "proxy_security_id": null,
      "sector": null,
      "security_id": "8E4L9XLl6MudjEpwPAAgivmdZRdBPJuvMPlPb",
      "sedol": null,
      "ticker_symbol": "NFLX180201C00355000",
      "type": "derivative",
      "subtype": "option",
      "unofficial_currency_code": null,
      "update_datetime": null
    },
    {
      "close_price": 94.808,
      "close_price_as_of": "2023-11-02",
      "cusip": "912797HE0",
      "fixed_income": {
        "face_value": 100,
        "issue_date": "2023-11-02",
        "maturity_date": "2024-10-31",
        "yield_rate": {
          "percentage": 5.43,
          "type": "coupon_equivalent"
        }
      },
      "industry": "Sovereign Government",
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": false,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "US Treasury Bill - 5.43% 31/10/2024 USD 100",
      "option_contract": null,
      "proxy_security_id": null,
      "sector": "Government",
      "security_id": "Lxe4yz4XQEtwb2YArO7RFMpPDvPxy7FALRyea",
      "sedol": null,
      "ticker_symbol": null,
      "type": "fixed income",
      "subtype": "bill",
      "unofficial_currency_code": null,
      "update_datetime": null
    },
    {
      "close_price": 9.08,
      "close_price_as_of": "2024-09-09",
      "cusip": null,
      "fixed_income": null,
      "industry": "Investment Trusts or Mutual Funds",
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": false,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "DoubleLine Total Return Bond I",
      "option_contract": null,
      "proxy_security_id": null,
      "sector": "Miscellaneous",
      "security_id": "AE5rBXra1AuZLE34rkvvIyG8918m3wtRzElnJ",
      "sedol": null,
      "ticker_symbol": "DBLTX",
      "type": "mutual fund",
      "subtype": "mutual fund",
      "unofficial_currency_code": null,
      "update_datetime": null
    },
    {
      "close_price": 42.15,
      "close_price_as_of": null,
      "cusip": null,
      "fixed_income": null,
      "industry": null,
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": false,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "iShares Inc MSCI Brazil",
      "option_contract": null,
      "proxy_security_id": null,
      "sector": null,
      "security_id": "abJamDazkgfvBkVGgnnLUWXoxnomp5up8llg4",
      "sedol": null,
      "ticker_symbol": "EWZ",
      "type": "etf",
      "subtype": "etf",
      "unofficial_currency_code": null,
      "update_datetime": null
    },
    {
      "close_price": 1,
      "close_price_as_of": null,
      "cusip": null,
      "fixed_income": null,
      "industry": null,
      "institution_id": null,
      "institution_security_id": null,
      "is_cash_equivalent": true,
      "isin": null,
      "iso_currency_code": "USD",
      "market_identifier_code": null,
      "name": "U S Dollar",
      "option_contract": null,
      "proxy_security_id": null,
      "sector": null,
      "security_id": "d6ePmbPxgWCWmMVv66q9iPV94n91vMtov5Are",
      "sedol": null,
      "ticker_symbol": null,
      "type": "cash",
      "subtype": "cash",
      "unofficial_currency_code": null,
      "update_datetime": null
    }
  ],
  "data_sources": {
    "numbers": "INSTITUTION",
    "owners": "INSTITUTION",
    "holdings": "INSTITUTION"
  }
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
