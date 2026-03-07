---
title: "Sandbox - Customize test data | Plaid Docs"
source_url: "https://plaid.com/docs/sandbox/user-custom/"
scraped_at: "2026-03-07T22:05:17+00:00"
---

# Create Sandbox test data

#### Use Sandbox accounts to create rich test data for Plaid products

Prefer to learn by watching? Check out this quick video guide to Custom Sandbox Users

=\*=\*=\*=

#### Customize Sandbox account data

In addition to a set of pre-populated [Sandbox test users](/docs/sandbox/test-credentials/), the Sandbox environment also provides the ability to create custom user accounts, which can be used in conjunction with [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) or Plaid Link to generate custom Sandbox data, or in conjunction with Plaid Link to test Link flows in the Sandbox.

Using these accounts, you can create your own testing data for the Assets, Auth, Balance, Identity, Investments, Liabilities, and Transactions products. You can also simulate an account with regular income or one that makes regular loan payments. For Link testing, custom accounts support multi-factor authentication, reCAPTCHA, and Link error flows.

To customize testing data for Document Income, see [Testing Document Income](https://plaid.com/docs/income/document-income/#testing-document-income).

=\*=\*=\*=

#### Configuring the custom user account

Custom user accounts can be configured and accessed in two ways. The easiest (and recommended) method is by using the [Sandbox Users](https://dashboard.plaid.com/developers/sandbox?tab=testUsers) tool, located in the Plaid Dashboard under Developers -> Sandbox. Set the username and [configuration object](/docs/sandbox/user-custom/#configuration-object-schema) in the Dashboard, then go through the Link flow on Sandbox with that username and any non-empty password. You can also use [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate) with `options.override_username` and `options.override_password` to create a public token for a custom user account while bypassing Link. Alternatively, you can log into Sandbox with the username `user_custom` and provide the [configuration object](/docs/sandbox/user-custom/#configuration-object-schema) as the password.

To aid in testing, Plaid maintains a [GitHub repo](https://github.com/plaid/sandbox-custom-users/) of pre-created custom user accounts that some users have found helpful.

##### Limitations of the custom user account

Very large configuration objects (larger than approximately 55kb, or approximately 250 transactions) are not supported and may cause the Link attempt to fail.

Using more than ten accounts on a single custom user is not supported and will cause the Link attempt to fail.

If you are using Consumer Report by Plaid Check, to use a custom user, you must click "add new account" in the Link flow rather than using one of the existing banks. Using an existing bank in the Plaid Passport flow will skip the ability to enter a custom username.

At OAuth institutions, custom users may not work properly: certain less frequently used customized fields may be overridden by the default values after the Link flow has completed, or the login may fail. If this occurs, retry the configuration using a non-OAuth institution, such as First Gingham Credit Union or First Platypus Bank.

=\*=\*=\*=

#### Configuration object schema

Custom test accounts are configured with a JSON configuration object formulated according to the schema below. All top level fields are optional. Sending an empty object as a configuration will result in an account configured with random balances and transaction history.

**Properties**

[`version`](/docs/sandbox/user-custom/#UserCustomPassword-version)

stringstring

The version of the password schema to use, possible values are 1 or 2. The default value is 2. You should only specify 1 if you know it is necessary for your test suite.

[`seed`](/docs/sandbox/user-custom/#UserCustomPassword-seed)

stringstring

A seed, in the form of a string, that will be used to randomly generate account and transaction data, if this data is not specified using the `override_accounts` argument. If no seed is specified, the randomly generated data will be different each time.  
Note that transactions data is generated relative to the Item's creation date. Different Items created on different dates with the same seed for transactions data will have different dates for the transactions. The number of days between each transaction and the Item creation will remain constant. For example, an Item created on December 15 might show a transaction on December 14. An Item created on December 20, using the same seed, would show that same transaction occurring on December 19.

[`override_accounts`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts)

[object][object]

An array of account overrides to configure the accounts for the Item. By default, if no override is specified, transactions and account data will be randomly generated based on the account type and subtype, and other products will have fixed or empty data.

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-type)

stringstring

`investment:` Investment account.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`payroll:` Payroll account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `payroll`, `other`

[`subtype`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-subtype)

stringstring

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`starting_balance`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-starting-balance)

numbernumber

If provided, the account will start with this amount as the current balance.  
  

Format: `double`

[`force_available_balance`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-force-available-balance)

numbernumber

If provided, the account will always have this amount as its available balance, regardless of current balance or changes in transactions over time. Cannot be set together with `has_null_available_balance`.  
  

Format: `double`

[`has_null_available_balance`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-has-null-available-balance)

booleanboolean

If set to `true`, the account will always have null as its available balance, regardless of current balance or changes in transactions over time. Cannot be set together with `force_available_balance`.

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-currency)

stringstring

ISO-4217 currency code. If provided, the account will be denominated in the given currency. Transactions will also be in this currency by default.

[`meta`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-meta)

objectobject

Allows specifying the metadata of the test account

[`name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-meta-name)

stringstring

The account's name

[`official_name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-meta-official-name)

stringstring

The account's official name

[`limit`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-meta-limit)

numbernumber

The account's limit  
  

Format: `double`

[`mask`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-meta-mask)

stringstring

The account's mask. Should be an empty string or a string of 2-4 alphanumeric characters. This allows you to model a mask which does not match the account number (such as with a virtual account number).  
  

Pattern: `^$|^[A-Za-z0-9]{2,4}$`

[`numbers`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers)

objectobject

Account and bank identifier number data used to configure the test account. All values are optional.

[`account`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers-account)

stringstring

Will be used for the account number.

[`ach_routing`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers-ach-routing)

stringstring

Must be a valid ACH routing number. To test `/transfer/capabilities/get`, set this to 322271627 to force a `true` result.

[`ach_wire_routing`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers-ach-wire-routing)

stringstring

Must be a valid wire transfer routing number.

[`eft_institution`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers-eft-institution)

stringstring

EFT institution number. Must be specified alongside `eft_branch`.

[`eft_branch`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers-eft-branch)

stringstring

EFT branch number. Must be specified alongside `eft_institution`.

[`international_bic`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers-international-bic)

stringstring

Bank identifier code (BIC). Must be specified alongside `international_iban`.

[`international_iban`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers-international-iban)

stringstring

International bank account number (IBAN). If no account number is specified via `account`, will also be used as the account number by default. Must be specified alongside `international_bic`.

[`bacs_sort_code`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-numbers-bacs-sort-code)

stringstring

BACS sort code

[`transactions`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-transactions)

[object][object]

Specify the list of transactions on the account.

[`date_transacted`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-transactions-date-transacted)

stringstring

The date of the transaction, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format. Transactions in Sandbox will move from pending to posted once their transaction date has been reached. If a `date_transacted` is not provided by the institution, a transaction date may be available in the [`authorized_date`](https://plaid.com/docs/api/products/transactions/#transactions-get-response-transactions-authorized-date) field.  
  

Format: `date`

[`date_posted`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-transactions-date-posted)

stringstring

The date the transaction posted, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format. Posted dates in the past or present will result in posted transactions; posted dates in the future will result in pending transactions.  
  

Format: `date`

[`amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-transactions-amount)

numbernumber

The transaction amount. Can be negative.  
  

Format: `double`

[`description`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-transactions-description)

stringstring

The transaction description.

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-transactions-currency)

stringstring

The ISO-4217 format currency code for the transaction.

[`holdings`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings)

objectobject

Specify the holdings on the account.

[`institution_price`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-institution-price)

numbernumber

The last price given by the institution for this security  
  

Format: `double`

[`institution_price_as_of`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-institution-price-as-of)

stringstring

The date at which `institution_price` was current. Must be formatted as an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) date.  
  

Format: `date`

[`cost_basis`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-cost-basis)

numbernumber

The total cost basis of the holding (e.g., the total amount spent to acquire all assets currently in the holding).  
  

Format: `double`

[`quantity`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-quantity)

numbernumber

The total quantity of the asset held, as reported by the financial institution.  
  

Format: `double`

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-currency)

stringstring

Either a valid `iso_currency_code` or `unofficial_currency_code`

[`security`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-security)

objectobject

Specify the security associated with the holding or investment transaction. When inputting custom security data to the Sandbox, Plaid will perform post-data-retrieval normalization and enrichment. These processes may cause the data returned by the Sandbox to be slightly different from the data you input. An ISO-4217 currency code and a security identifier (`ticker_symbol`, `cusip`, or `isin`) are required.

[`isin`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-security-isin)

stringstring

12-character ISIN, a globally unique securities identifier. A verified CUSIP Global Services license is required to receive this data. This field will be null by default for new customers, and null for existing customers starting March 12, 2024. If you would like access to this field, please [request ISIN/CUSIP access here](https://docs.google.com/forms/d/e/1FAIpQLSd9asHEYEfmf8fxJTHZTAfAzW4dugsnSu-HS2J51f1mxwd6Sw/viewform).

[`cusip`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-security-cusip)

stringstring

9-character CUSIP, an identifier assigned to North American securities. A verified CUSIP Global Services license is required to receive this data. This field will be null by default for new customers, and null for existing customers starting March 12, 2024. If you would like access to this field, please [request ISIN/CUSIP access here](https://docs.google.com/forms/d/e/1FAIpQLSd9asHEYEfmf8fxJTHZTAfAzW4dugsnSu-HS2J51f1mxwd6Sw/viewform).

[`name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-security-name)

stringstring

A descriptive name for the security, suitable for display.

[`ticker_symbol`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-security-ticker-symbol)

stringstring

The security’s trading symbol for publicly traded securities, and otherwise a short identifier if available.

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-holdings-security-currency)

stringstring

Either a valid `iso_currency_code` or `unofficial_currency_code`

[`investment_transactions`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions)

objectobject

Specify the list of investments transactions on the account.

[`date`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-date)

stringstring

Posting date for the transaction. Must be formatted as an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) date.  
  

Format: `date`

[`name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-name)

stringstring

The institution's description of the transaction.

[`quantity`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-quantity)

numbernumber

The number of units of the security involved in this transaction. Must be positive if the type is a buy and negative if the type is a sell.  
  

Format: `double`

[`price`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-price)

numbernumber

The price of the security at which this transaction occurred.  
  

Format: `double`

[`fees`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-fees)

numbernumber

The combined value of all fees applied to this transaction.  
  

Format: `double`

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-type)

stringstring

The type of the investment transaction. Possible values are:
`buy`: Buying an investment
`sell`: Selling an investment
`cash`: Activity that modifies a cash position
`fee`: A fee on the account
`transfer`: Activity that modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-currency)

stringstring

Either a valid `iso_currency_code` or `unofficial_currency_code`

[`security`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-security)

objectobject

Specify the security associated with the holding or investment transaction. When inputting custom security data to the Sandbox, Plaid will perform post-data-retrieval normalization and enrichment. These processes may cause the data returned by the Sandbox to be slightly different from the data you input. An ISO-4217 currency code and a security identifier (`ticker_symbol`, `cusip`, or `isin`) are required.

[`isin`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-security-isin)

stringstring

12-character ISIN, a globally unique securities identifier. A verified CUSIP Global Services license is required to receive this data. This field will be null by default for new customers, and null for existing customers starting March 12, 2024. If you would like access to this field, please [request ISIN/CUSIP access here](https://docs.google.com/forms/d/e/1FAIpQLSd9asHEYEfmf8fxJTHZTAfAzW4dugsnSu-HS2J51f1mxwd6Sw/viewform).

[`cusip`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-security-cusip)

stringstring

9-character CUSIP, an identifier assigned to North American securities. A verified CUSIP Global Services license is required to receive this data. This field will be null by default for new customers, and null for existing customers starting March 12, 2024. If you would like access to this field, please [request ISIN/CUSIP access here](https://docs.google.com/forms/d/e/1FAIpQLSd9asHEYEfmf8fxJTHZTAfAzW4dugsnSu-HS2J51f1mxwd6Sw/viewform).

[`name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-security-name)

stringstring

A descriptive name for the security, suitable for display.

[`ticker_symbol`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-security-ticker-symbol)

stringstring

The security’s trading symbol for publicly traded securities, and otherwise a short identifier if available.

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-investment-transactions-security-currency)

stringstring

Either a valid `iso_currency_code` or `unofficial_currency_code`

[`identity`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity)

objectobject

Data about the owner or owners of an account. Any fields not specified will be filled in with default Sandbox information.

[`names`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-names)

[string][string]

A list of names associated with the account by the financial institution. These should always be the names of individuals, even for business accounts. Note that the same name data will be used for all accounts associated with an Item.

[`phone_numbers`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-phone-numbers)

[object][object]

A list of phone numbers associated with the account.

[`data`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-phone-numbers-data)

stringstring

The phone number.

[`primary`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-phone-numbers-primary)

booleanboolean

When `true`, identifies the phone number as the primary number on an account.

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-phone-numbers-type)

stringstring

The type of phone number.  
  

Possible values: `home`, `work`, `office`, `mobile`, `mobile1`, `other`

[`emails`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-emails)

[object][object]

A list of email addresses associated with the account.

[`data`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-emails-data)

stringstring

The email address.

[`primary`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-emails-primary)

booleanboolean

When `true`, identifies the email address as the primary email on an account.

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-emails-type)

stringstring

The type of email account as described by the financial institution.  
  

Possible values: `primary`, `secondary`, `other`

[`addresses`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-addresses)

[object][object]

Data about the various addresses associated with the account.

[`data`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-addresses-data)

objectobject

Data about the components comprising an address.

[`city`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-addresses-data-city)

stringstring

The full city name

[`region`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-addresses-data-region)

stringstring

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-addresses-data-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-addresses-data-postal-code)

stringstring

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-addresses-data-country)

stringstring

The ISO 3166-1 alpha-2 country code

[`primary`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-identity-addresses-primary)

booleanboolean

When `true`, identifies the address as the primary address on an account.

[`liability`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability)

objectobject

Used to configure Sandbox test data for the Liabilities product

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-type)

stringstring

The type of the liability object, either `credit` or `student`. Mortgages are not currently supported in the custom Sandbox.

[`purchase_apr`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-purchase-apr)

numbernumber

The purchase APR percentage value. For simplicity, this is the only interest rate used to calculate interest charges. Can only be set if `type` is `credit`.  
  

Format: `double`

[`cash_apr`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-cash-apr)

numbernumber

The cash APR percentage value. Can only be set if `type` is `credit`.  
  

Format: `double`

[`balance_transfer_apr`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-balance-transfer-apr)

numbernumber

The balance transfer APR percentage value. Can only be set if `type` is `credit`.  
  

Format: `double`

[`special_apr`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-special-apr)

numbernumber

The special APR percentage value. Can only be set if `type` is `credit`.  
  

Format: `double`

[`last_payment_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-last-payment-amount)

numbernumber

Override the `last_payment_amount` field. Can only be set if `type` is `credit`.  
  

Format: `double`

[`minimum_payment_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-minimum-payment-amount)

numbernumber

Override the `minimum_payment_amount` field. Can only be set if `type` is `credit` or `student`.  
  

Format: `double`

[`is_overdue`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-is-overdue)

booleanboolean

Override the `is_overdue` field

[`origination_date`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-origination-date)

stringstring

The date on which the loan was initially lent, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format. Can only be set if `type` is `student`.  
  

Format: `date`

[`principal`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-principal)

numbernumber

The original loan principal. Can only be set if `type` is `student`.  
  

Format: `double`

[`nominal_apr`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-nominal-apr)

numbernumber

The interest rate on the loan as a percentage. Can only be set if `type` is `student`.  
  

Format: `double`

[`interest_capitalization_grace_period_months`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-interest-capitalization-grace-period-months)

numbernumber

If set, interest capitalization begins at the given number of months after loan origination. By default interest is never capitalized. Can only be set if `type` is `student`.

[`repayment_model`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-repayment-model)

objectobject

Student loan repayment information used to configure Sandbox test data for the Liabilities product

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-repayment-model-type)

stringstring

The only currently supported value for this field is `standard`.

[`non_repayment_months`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-repayment-model-non-repayment-months)

numbernumber

Configures the number of months before repayment starts.

[`repayment_months`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-repayment-model-repayment-months)

numbernumber

Configures the number of months of repayments before the loan is paid off.

[`expected_payoff_date`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-expected-payoff-date)

stringstring

Override the `expected_payoff_date` field. Can only be set if `type` is `student`.  
  

Format: `date`

[`guarantor`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-guarantor)

stringstring

Override the `guarantor` field. Can only be set if `type` is `student`.

[`is_federal`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-is-federal)

booleanboolean

Override the `is_federal` field. Can only be set if `type` is `student`.

[`loan_name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-loan-name)

stringstring

Override the `loan_name` field. Can only be set if `type` is `student`.

[`loan_status`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-loan-status)

objectobject

An object representing the status of the student loan

[`end_date`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-loan-status-end-date)

stringstring

The date until which the loan will be in its current status. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  
  

Format: `date`

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-loan-status-type)

stringstring

The status type of the student loan  
  

Possible values: `cancelled`, `charged off`, `claim`, `consolidated`, `deferment`, `delinquent`, `discharged`, `extension`, `forbearance`, `in grace`, `in military`, `in school`, `not fully disbursed`, `other`, `paid in full`, `refunded`, `repayment`, `transferred`, `pending idr`

[`payment_reference_number`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-payment-reference-number)

stringstring

Override the `payment_reference_number` field. Can only be set if `type` is `student`.

[`repayment_plan_description`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-repayment-plan-description)

stringstring

Override the `repayment_plan.description` field. Can only be set if `type` is `student`.

[`repayment_plan_type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-repayment-plan-type)

stringstring

Override the `repayment_plan.type` field. Can only be set if `type` is `student`. Possible values are: `"extended graduated"`, `"extended standard"`, `"graduated"`, `"income-contingent repayment"`, `"income-based repayment"`, `"income-sensitive repayment"`, `"interest only"`, `"other"`, `"pay as you earn"`, `"revised pay as you earn"`, `"standard"`, or `"saving on a valuable education"`.

[`sequence_number`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-sequence-number)

stringstring

Override the `sequence_number` field. Can only be set if `type` is `student`.

[`servicer_address`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-servicer-address)

objectobject

A physical mailing address.

[`data`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-servicer-address-data)

objectobject

Data about the components comprising an address.

[`city`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-servicer-address-data-city)

stringstring

The full city name

[`region`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-servicer-address-data-region)

stringstring

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-servicer-address-data-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-servicer-address-data-postal-code)

stringstring

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-servicer-address-data-country)

stringstring

The ISO 3166-1 alpha-2 country code

[`primary`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-liability-servicer-address-primary)

booleanboolean

When `true`, identifies the address as the primary address on an account.

[`inflow_model`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-inflow-model)

objectobject

The `inflow_model` allows you to model a test account that receives regular income or make regular payments on a loan. Any transactions generated by the `inflow_model` will appear in addition to randomly generated test data or transactions specified by `override_accounts`.

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-inflow-model-type)

stringstring

Inflow model. One of the following:  
`none`: No income  
`monthly-income`: Income occurs once per month `monthly-balance-payment`: Pays off the balance on a liability account at the given statement day of month.  
`monthly-interest-only-payment`: Makes an interest-only payment on a liability account at the given statement day of month.  
Note that account types supported by Liabilities will accrue interest in the Sandbox. The types impacted are account type `credit` with subtype `credit` or `paypal`, and account type `loan` with subtype `student` or `mortgage`.

[`income_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-inflow-model-income-amount)

numbernumber

Amount of income per month. This value is required if `type` is `monthly-income`.  
  

Format: `double`

[`payment_day_of_month`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-inflow-model-payment-day-of-month)

numbernumber

Number between 1 and 28, or `last` meaning the last day of the month. The day of the month on which the income transaction will appear. The name of the income transaction. This field is required if `type` is `monthly-income`, `monthly-balance-payment` or `monthly-interest-only-payment`.

[`transaction_name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-inflow-model-transaction-name)

stringstring

The name of the income transaction. This field is required if `type` is `monthly-income`, `monthly-balance-payment` or `monthly-interest-only-payment`.

[`statement_day_of_month`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-inflow-model-statement-day-of-month)

stringstring

Number between 1 and 28, or `last` meaning the last day of the month. The day of the month on which the balance is calculated for the next payment. The name of the income transaction. This field is required if `type` is `monthly-balance-payment` or `monthly-interest-only-payment`.

[`income`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income)

objectobject

Specify payroll data on the account.

[`paystubs`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs)

[object][object]

A list of paystubs associated with the account.

[`employer`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employer)

objectobject

The employer on the paystub.

[`name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employer-name)

stringstring

The name of the employer.

[`address`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employer-address)

objectobject

The address of the employer.

[`city`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employer-address-city)

stringstring

The full city name.

[`region`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employer-address-region)

stringstring

The region or state
Example: `"NC"`

[`street`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employer-address-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employer-address-postal-code)

stringstring

5 digit postal code.

[`country`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employer-address-country)

stringstring

The country of the address.

[`employee`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee)

objectobject

The employee on the paystub.

[`name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-name)

stringstring

The name of the employee.

[`address`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-address)

objectobject

The address of the employee.

[`city`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-address-city)

stringstring

The full city name.

[`region`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-address-region)

stringstring

The region or state
Example: `"NC"`

[`street`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-address-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-address-postal-code)

stringstring

5 digit postal code.

[`country`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-address-country)

stringstring

The country of the address.

[`marital_status`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-marital-status)

stringstring

Marital status of the employee - either `single` or `married`.  
  

Possible values: `single`, `married`

[`taxpayer_id`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-taxpayer-id)

objectobject

Taxpayer ID of the individual receiving the paystub.

[`id_type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-taxpayer-id-id-type)

stringstring

Type of ID, e.g. 'SSN'

[`id_mask`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-employee-taxpayer-id-id-mask)

stringstring

ID mask; i.e. last 4 digits of the taxpayer ID

[`net_pay`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-net-pay)

objectobject

An object representing information about the net pay amount on the paystub.

[`description`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-net-pay-description)

stringstring

Description of the net pay

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-net-pay-currency)

stringstring

The ISO-4217 currency code of the net pay.

[`ytd_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-net-pay-ytd-amount)

numbernumber

The year-to-date amount of the net pay  
  

Format: `double`

[`deductions`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions)

objectobject

An object with the deduction information found on a paystub.

[`breakdown`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-breakdown)

[object][object]

[`current_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-breakdown-current-amount)

numbernumber

Raw amount of the deduction  
  

Format: `double`

[`description`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-breakdown-description)

stringstring

Description of the deduction line item

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-breakdown-currency)

stringstring

The ISO-4217 currency code of the line item.

[`ytd_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-breakdown-ytd-amount)

numbernumber

The year-to-date amount of the deduction  
  

Format: `double`

[`total`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-total)

objectobject

An object representing the total deductions for the pay period

[`current_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-total-current-amount)

numbernumber

Raw amount of the deduction  
  

Format: `double`

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-total-currency)

stringstring

The ISO-4217 currency code of the line item.

[`ytd_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-deductions-total-ytd-amount)

numbernumber

The year-to-date total amount of the deductions  
  

Format: `double`

[`earnings`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings)

objectobject

An object representing both a breakdown of earnings on a paystub and the total earnings.

[`breakdown`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-breakdown)

[object][object]

[`canonical_description`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-breakdown-canonical-description)

stringstring

Commonly used term to describe the earning line item.  
  

Possible values: `BONUS`, `COMMISSION`, `OVERTIME`, `PAID TIME OFF`, `REGULAR PAY`, `VACATION`, `BASIC ALLOWANCE HOUSING`, `BASIC ALLOWANCE SUBSISTENCE`, `OTHER`, `null`

[`current_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-breakdown-current-amount)

numbernumber

Raw amount of the earning line item.  
  

Format: `double`

[`description`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-breakdown-description)

stringstring

Description of the earning line item.

[`hours`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-breakdown-hours)

numbernumber

Number of hours applicable for this earning.

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-breakdown-currency)

stringstring

The ISO-4217 currency code of the line item.

[`rate`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-breakdown-rate)

numbernumber

Hourly rate applicable for this earning.  
  

Format: `double`

[`ytd_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-breakdown-ytd-amount)

numbernumber

The year-to-date amount of the deduction.  
  

Format: `double`

[`total`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-total)

objectobject

An object representing both the current pay period and year to date amount for an earning category.

[`hours`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-total-hours)

numbernumber

Total number of hours worked for this pay period

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-total-currency)

stringstring

The ISO-4217 currency code of the line item

[`ytd_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-earnings-total-ytd-amount)

numbernumber

The year-to-date amount for the total earnings  
  

Format: `double`

[`pay_period_details`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details)

objectobject

Details about the pay period.

[`check_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-check-amount)

numbernumber

The amount of the paycheck.  
  

Format: `double`

[`distribution_breakdown`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-distribution-breakdown)

[object][object]

[`account_name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-distribution-breakdown-account-name)

stringstring

Name of the account for the given distribution.

[`bank_name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-distribution-breakdown-bank-name)

stringstring

The name of the bank that the payment is being deposited to.

[`current_amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-distribution-breakdown-current-amount)

numbernumber

The amount distributed to this account.  
  

Format: `double`

[`currency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-distribution-breakdown-currency)

stringstring

The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.

[`mask`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-distribution-breakdown-mask)

stringstring

The last 2-4 alphanumeric characters of an account's official account number.

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-distribution-breakdown-type)

stringstring

Type of the account that the paystub was sent to (e.g. 'checking').

[`end_date`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-end-date)

stringstring

The pay period end date, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format: "yyyy-mm-dd".  
  

Format: `date`

[`gross_earnings`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-gross-earnings)

numbernumber

Total earnings before tax/deductions.  
  

Format: `double`

[`pay_date`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-pay-date)

stringstring

The date on which the paystub was issued, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`pay_frequency`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-pay-frequency)

stringstring

The frequency at which an individual is paid.  
  

Possible values: `PAY_FREQUENCY_UNKNOWN`, `PAY_FREQUENCY_WEEKLY`, `PAY_FREQUENCY_BIWEEKLY`, `PAY_FREQUENCY_SEMIMONTHLY`, `PAY_FREQUENCY_MONTHLY`, `null`

[`pay_day`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-pay-day)

deprecatedstringdeprecated, string

The date on which the paystub was issued, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`start_date`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-paystubs-pay-period-details-start-date)

stringstring

The pay period start date, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format: "yyyy-mm-dd".  
  

Format: `date`

[`w2s`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s)

[object][object]

A list of w2s associated with the account.

[`employer`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer)

objectobject

The employer on the paystub.

[`name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer-name)

stringstring

The name of the employer.

[`address`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer-address)

objectobject

The address of the employer.

[`city`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer-address-city)

stringstring

The full city name.

[`region`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer-address-region)

stringstring

The region or state
Example: `"NC"`

[`street`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer-address-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer-address-postal-code)

stringstring

5 digit postal code.

[`country`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer-address-country)

stringstring

The country of the address.

[`employee`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee)

objectobject

The employee on the paystub.

[`name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-name)

stringstring

The name of the employee.

[`address`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-address)

objectobject

The address of the employee.

[`city`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-address-city)

stringstring

The full city name.

[`region`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-address-region)

stringstring

The region or state
Example: `"NC"`

[`street`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-address-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-address-postal-code)

stringstring

5 digit postal code.

[`country`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-address-country)

stringstring

The country of the address.

[`marital_status`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-marital-status)

stringstring

Marital status of the employee - either `single` or `married`.  
  

Possible values: `single`, `married`

[`taxpayer_id`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-taxpayer-id)

objectobject

Taxpayer ID of the individual receiving the paystub.

[`id_type`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-taxpayer-id-id-type)

stringstring

Type of ID, e.g. 'SSN'

[`id_mask`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employee-taxpayer-id-id-mask)

stringstring

ID mask; i.e. last 4 digits of the taxpayer ID

[`tax_year`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-tax-year)

stringstring

The tax year of the W2 document.

[`employer_id_number`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-employer-id-number)

stringstring

An employer identification number or EIN.

[`wages_tips_other_comp`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-wages-tips-other-comp)

stringstring

Wages from tips and other compensation.

[`federal_income_tax_withheld`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-federal-income-tax-withheld)

stringstring

Federal income tax withheld for the tax year.

[`social_security_wages`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-social-security-wages)

stringstring

Wages from social security.

[`social_security_tax_withheld`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-social-security-tax-withheld)

stringstring

Social security tax withheld for the tax year.

[`medicare_wages_and_tips`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-medicare-wages-and-tips)

stringstring

Wages and tips from medicare.

[`medicare_tax_withheld`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-medicare-tax-withheld)

stringstring

Medicare tax withheld for the tax year.

[`social_security_tips`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-social-security-tips)

stringstring

Tips from social security.

[`allocated_tips`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-allocated-tips)

stringstring

Allocated tips.

[`box_9`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-box-9)

stringstring

Contents from box 9 on the W2.

[`dependent_care_benefits`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-dependent-care-benefits)

stringstring

Dependent care benefits.

[`nonqualified_plans`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-nonqualified-plans)

stringstring

Nonqualified plans.

[`box_12`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-box-12)

[object][object]

[`code`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-box-12-code)

stringstring

W2 Box 12 code.

[`amount`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-box-12-amount)

stringstring

W2 Box 12 amount.

[`statutory_employee`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-statutory-employee)

stringstring

Statutory employee.

[`retirement_plan`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-retirement-plan)

stringstring

Retirement plan.

[`third_party_sick_pay`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-third-party-sick-pay)

stringstring

Third party sick pay.

[`other`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-other)

stringstring

Other.

[`state_and_local_wages`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-state-and-local-wages)

[object][object]

[`state`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-state-and-local-wages-state)

stringstring

State associated with the wage.

[`employer_state_id_number`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-state-and-local-wages-employer-state-id-number)

stringstring

State identification number of the employer.

[`state_wages_tips`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-state-and-local-wages-state-wages-tips)

stringstring

Wages and tips from the specified state.

[`state_income_tax`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-state-and-local-wages-state-income-tax)

stringstring

Income tax from the specified state.

[`local_wages_tips`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-state-and-local-wages-local-wages-tips)

stringstring

Wages and tips from the locality.

[`local_income_tax`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-state-and-local-wages-local-income-tax)

stringstring

Income tax from the locality.

[`locality_name`](/docs/sandbox/user-custom/#UserCustomPassword-override-accounts-income-w2s-state-and-local-wages-locality-name)

stringstring

Name of the locality.

[`mfa`](/docs/sandbox/user-custom/#UserCustomPassword-mfa)

objectobject

Specifies the multi-factor authentication settings to use with this test account

[`type`](/docs/sandbox/user-custom/#UserCustomPassword-mfa-type)

stringstring

Possible values are `device`, `selections`, or `questions`.  
If value is `device`, the MFA answer is `1234`.  
If value is `selections`, the MFA answer is always the first option.  
If value is `questions`, the MFA answer is `answer_<i>_<j>` for the j-th question in the i-th round, starting from 0. For example, the answer to the first question in the second round is `answer_1_0`.

[`question_rounds`](/docs/sandbox/user-custom/#UserCustomPassword-mfa-question-rounds)

numbernumber

Number of rounds of questions. Required if value of `type` is `questions`.

[`questions_per_round`](/docs/sandbox/user-custom/#UserCustomPassword-mfa-questions-per-round)

numbernumber

Number of questions per round. Required if value of `type` is `questions`. If value of type is `selections`, default value is 2.

[`selection_rounds`](/docs/sandbox/user-custom/#UserCustomPassword-mfa-selection-rounds)

numbernumber

Number of rounds of selections, used if `type` is `selections`. Defaults to 1.

[`selections_per_question`](/docs/sandbox/user-custom/#UserCustomPassword-mfa-selections-per-question)

numbernumber

Number of available answers per question, used if `type` is `selection`. Defaults to 2.

[`recaptcha`](/docs/sandbox/user-custom/#UserCustomPassword-recaptcha)

stringstring

You may trigger a reCAPTCHA in Plaid Link in the Sandbox environment by using the recaptcha field. Possible values are `good` or `bad`. A value of `good` will result in successful Item creation and `bad` will result in a `RECAPTCHA_BAD` error to simulate a failed reCAPTCHA. Both values require the reCAPTCHA to be manually solved within Plaid Link.

[`force_error`](/docs/sandbox/user-custom/#UserCustomPassword-force-error)

stringstring

An error code to force on Item creation. Possible values are:  
`"INSTITUTION_NOT_RESPONDING"`
`"INSTITUTION_NO_LONGER_SUPPORTED"`
`"INVALID_CREDENTIALS"`
`"INVALID_MFA"`
`"ITEM_LOCKED"`
`"ITEM_LOGIN_REQUIRED"`
`"ITEM_NOT_SUPPORTED"`
`"INVALID_LINK_TOKEN"`
`"MFA_NOT_SUPPORTED"`
`"NO_ACCOUNTS"`
`"PLAID_ERROR"`
`"USER_INPUT_TIMEOUT"`
`"USER_SETUP_REQUIRED"`

```
{
  "seed": "my-seed-string-3",
  "override_accounts": [
    {
      "type": "depository",
      "subtype": "checking",
      "identity": {
        "names": [
          "John Doe"
        ],
        "phone_numbers": [
          {
            "primary": true,
            "type": "home",
            "data": "4673956022"
          }
        ],
        "emails": [
          {
            "primary": true,
            "type": "primary",
            "data": "accountholder0@example.com"
          }
        ],
        "addresses": [
          {
            "primary": true,
            "data": {
              "city": "Malakoff",
              "region": "NY",
              "street": "2992 Cameron Road",
              "postal_code": "14236",
              "country": "US"
            }
          }
        ]
      },
      "transactions": [
        {
          "date_transacted": "2023-10-01",
          "date_posted": "2023-10-03",
          "currency": "USD",
          "amount": 100,
          "description": "1 year Netflix subscription"
        },
        {
          "date_transacted": "2023-10-01",
          "date_posted": "2023-10-02",
          "currency": "USD",
          "amount": 100,
          "description": "1 year mobile subscription"
        }
      ]
    },
    {
      "type": "loan",
      "subtype": "student",
      "liability": {
        "type": "student",
        "origination_date": "2023-01-01",
        "principal": 10000,
        "nominal_apr": 6.25,
        "loan_name": "Plaid Student Loan",
        "repayment_model": {
          "type": "standard",
          "non_repayment_months": 12,
          "repayment_months": 120
        }
      }
    },
    {
      "type": "credit",
      "subtype": "credit card",
      "starting_balance": 10000,
      "inflow_model": {
        "type": "monthly-interest-only-payment",
        "payment_day_of_month": 15,
        "statement_day_of_month": 13,
        "transaction_name": "Interest Payment"
      },
      "liability": {
        "type": "credit",
        "purchase_apr": 12.9,
        "balance_transfer_apr": 15.24,
        "cash_apr": 28.45,
        "special_apr": 0,
        "last_payment_amount": 500,
        "minimum_payment_amount": 10
      }
    },
    {
      "type": "investment",
      "subtype": "brokerage",
      "investment_transactions": [
        {
          "date": "2023-07-01",
          "name": "buy stock",
          "quantity": 10,
          "price": 10,
          "fees": 20,
          "type": "buy",
          "currency": "USD",
          "security": {
            "ticker_symbol": "PLAID",
            "currency": "USD"
          }
        }
      ],
      "holdings": [
        {
          "institution_price": 10,
          "institution_price_as_of": "2023-08-01",
          "cost_basis": 10,
          "quantity": 10,
          "currency": "USD",
          "security": {
            "ticker_symbol": "PLAID",
            "currency": "USD"
          }
        }
      ]
    },
    {
      "type": "payroll",
      "subtype": "payroll",
      "income": {
        "paystubs": [
          {
            "employer": {
              "name": "Heartland Toy Company"
            },
            "employee": {
              "name": "Chip Hazard",
              "address": {
                "city": "Burbank",
                "region": "CA",
                "street": "411 N Hollywood Way",
                "postal_code": "91505",
                "country": "US"
              }
            },
            "income_breakdown": [
              {
                "type": "regular",
                "rate": 20,
                "hours": 40,
                "total": 800
              },
              {
                "type": "overtime",
                "rate": 30,
                "hours": 6.68,
                "total": 200.39
              }
            ],
            "pay_period_details": {
              "start_date": "2021-05-04",
              "end_date": "2021-05-18",
              "gross_earnings": 1000.39,
              "check_amount": 499.28
            }
          }
        ]
      }
    }
  ]
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
