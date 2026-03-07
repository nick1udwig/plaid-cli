---
title: "API - Institutions | Plaid Docs"
source_url: "https://plaid.com/docs/api/institutions/"
scraped_at: "2026-03-07T22:03:48+00:00"
---

# Institutions endpoints

#### Fetch data about supported institutions

=\*=\*=\*=

#### Institution coverage

For a user-friendly overview of which institutions Plaid supports, and the product coverage at each institution, see the [US and Canada Coverage Explorer](/docs/institutions/) or [European Coverage Explorer](/docs/institutions/europe/).

The [status dashboard](https://dashboard.plaid.com/activity/status) also provides a browsable view of institutions and supported products, with a focus on reporting institution health and downtimes.

For more detailed institution information, or to access this data programmatically, use the API endpoints described on this page.

##### Supported countries

For a list of which products are supported for each country, see [supported products by country](https://support.plaid.com/hc/en-us/articles/27895826947735-What-Plaid-products-are-supported-in-each-country-and-region) or the docs for the specific product you are interested in.

By default, customers in the United States and Canada receive access to institutions in all countries in Sandbox, and to United States and Canada in Production. To gain access to additional countries in Production, [file a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

| Endpoints |  |
| --- | --- |
| [`/institutions/get`](/docs/api/institutions/#institutionsget) | Get a list of all supported institutions meeting specified criteria |
| [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) | Get details about a specific institution |
| [`/institutions/search`](/docs/api/institutions/#institutionssearch) | Look up an institution by name |

The interface for these endpoints has changed in API version 2020-09-14. If you are using an older API version, see [API versioning](/docs/api/versioning/).

### Endpoints

=\*=\*=\*=

#### `/institutions/get`

#### Get details of all supported institutions

Returns a JSON response containing details on all financial institutions currently supported by Plaid. Because Plaid supports thousands of institutions, results are paginated.

If there is no overlap between an institution’s enabled products and a client’s enabled products, then the institution will be filtered out from the response. As a result, the number of institutions returned may not match the count specified in the call.

/institutions/get

**Request fields**

[`client_id`](/docs/api/institutions/#institutions-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/institutions/#institutions-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`count`](/docs/api/institutions/#institutions-get-request-count)

requiredintegerrequired, integer

The total number of Institutions to return.  
  

Minimum: `1`

Maximum: `500`

[`offset`](/docs/api/institutions/#institutions-get-request-offset)

requiredintegerrequired, integer

The number of Institutions to skip.  
  

Minimum: `0`

[`country_codes`](/docs/api/institutions/#institutions-get-request-country-codes)

required[string]required, [string]

Specify which country or countries to include institutions from, using the ISO-3166-1 alpha-2 country code standard.  
In API versions 2019-05-29 and earlier, the `country_codes` parameter is an optional parameter within the `options` object and will default to `[US]` if it is not supplied.  
  

Min items: `1`

Possible values: `US`, `GB`, `ES`, `NL`, `FR`, `IE`, `CA`, `DE`, `IT`, `PL`, `DK`, `NO`, `SE`, `EE`, `LT`, `LV`, `PT`, `BE`, `AT`, `FI`

[`options`](/docs/api/institutions/#institutions-get-request-options)

objectobject

An optional object to filter `/institutions/get` results.

[`products`](/docs/api/institutions/#institutions-get-request-options-products)

[string][string]

Filter the Institutions based on which products they support. Will only return institutions that support all listed products. When filtering based on `auth`, an institution must support Instant Auth to match the criterion. To filter for Signal Transaction Scores support, use `balance`. To filter for Transfer support, use `auth`.  
  

Min items: `1`

Possible values: `assets`, `auth`, `balance`, `employment`, `identity`, `cra_base_report`, `cra_income_insights`, `cra_cashflow_insights`, `cra_lend_score`, `cra_network_insights`, `cra_partner_insights`, `income_verification`, `identity_verification`, `investments`, `liabilities`, `payment_initiation`, `standing_orders`, `transactions`

[`routing_numbers`](/docs/api/institutions/#institutions-get-request-options-routing-numbers)

[string][string]

Specify an array of routing numbers to filter institutions. The response will only return institutions that match all of the routing numbers in the array. Routing number records used for this matching are generally comprehensive; however, failure to match a given routing number to an institution does not necessarily mean that the institution is unsupported by Plaid. Invalid routing numbers (numbers that are not 9 digits in length or do not have a valid checksum) will be filtered from the array before the response is processed. If all provided routing numbers are invalid, an `INVALID_REQUEST` error with the code of `INVALID_FIELD` will be returned.

[`oauth`](/docs/api/institutions/#institutions-get-request-options-oauth)

booleanboolean

Limit results to institutions with or without OAuth login flows. Note that institutions will have `oauth` set to `true` if some Items associated with that institution are required to use OAuth flows; institutions in a state of migration to OAuth will have the `oauth` attribute set to `true`.

[`include_optional_metadata`](/docs/api/institutions/#institutions-get-request-options-include-optional-metadata)

booleanboolean

When `true`, return the institution's homepage URL, logo and primary brand color. Not all institutions' logos are available.  
Note that Plaid does not own any of the logos shared by the API, and that by accessing or using these logos, you agree that you are doing so at your own risk and will, if necessary, obtain all required permissions from the appropriate rights holders and adhere to any applicable usage guidelines. Plaid disclaims all express or implied warranties with respect to the logos.

[`include_auth_metadata`](/docs/api/institutions/#institutions-get-request-options-include-auth-metadata)

booleanboolean

When `true`, returns metadata related to the Auth product indicating which auth methods are supported.  
  

Default: `false`

[`include_payment_initiation_metadata`](/docs/api/institutions/#institutions-get-request-options-include-payment-initiation-metadata)

booleanboolean

When `true`, returns metadata related to the Payment Initiation product indicating which payment configurations are supported.  
  

Default: `false`

/institutions/get

```
// Pull institutions
const request: InstitutionsGetRequest = {
  count: 10,
  offset: 0,
  country_codes: ['US'],
};
try {
  const response = await plaidClient.institutionsGet(request);
  const institutions = response.data.institutions;
} catch (error) {
  // Handle error
}
```

/institutions/get

**Response fields**

[`institutions`](/docs/api/institutions/#institutions-get-response-institutions)

[object][object]

A list of Plaid institutions

[`institution_id`](/docs/api/institutions/#institutions-get-response-institutions-institution-id)

stringstring

Unique identifier for the institution. Note that the same institution may have multiple records, each with different institution IDs; for example, if the institution has migrated to OAuth, there may be separate `institution_id`s for the OAuth and non-OAuth versions of the institution. Institutions that operate in different countries or with multiple login portals may also have separate `institution_id`s for each country or portal.

[`name`](/docs/api/institutions/#institutions-get-response-institutions-name)

stringstring

The official name of the institution.

[`products`](/docs/api/institutions/#institutions-get-response-institutions-products)

[string][string]

A list of the Plaid products supported by the institution. Note that only institutions that support Instant Auth will return `auth` in the product array; institutions that do not list `auth` may still support other Auth methods such as Instant Match or Automated Micro-deposit Verification. To identify institutions that support those methods, use the `auth_metadata` object. For more details, see [Full Auth coverage](https://plaid.com/docs/auth/coverage/). The `income_verification` product here indicates support for Bank Income. Note: For Signal Transaction Scores and Transfer, listed institutions may be incomplete or incorrect. Instead, use the following: `balance` support also indicates coverage of Signal Transaction Scores; `auth` support also indicates coverage of Transfer.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`country_codes`](/docs/api/institutions/#institutions-get-response-institutions-country-codes)

[string][string]

A list of the country codes supported by the institution.  
  

Possible values: `US`, `GB`, `ES`, `NL`, `FR`, `IE`, `CA`, `DE`, `IT`, `PL`, `DK`, `NO`, `SE`, `EE`, `LT`, `LV`, `PT`, `BE`, `AT`, `FI`

[`url`](/docs/api/institutions/#institutions-get-response-institutions-url)

nullablestringnullable, string

The URL for the institution's website

[`primary_color`](/docs/api/institutions/#institutions-get-response-institutions-primary-color)

nullablestringnullable, string

Hexadecimal representation of the primary color used by the institution

[`logo`](/docs/api/institutions/#institutions-get-response-institutions-logo)

nullablestringnullable, string

Base64 encoded representation of the institution's logo, returned as a base64 encoded 152x152 PNG. Not all institutions' logos are available.

[`routing_numbers`](/docs/api/institutions/#institutions-get-response-institutions-routing-numbers)

[string][string]

A list of routing numbers known to be associated with the institution. This list is provided for the purpose of looking up institutions by routing number. It is generally comprehensive but is not guaranteed to be a complete list of routing numbers for an institution.

[`dtc_numbers`](/docs/api/institutions/#institutions-get-response-institutions-dtc-numbers)

[string][string]

A partial list of DTC numbers associated with the institution.

[`oauth`](/docs/api/institutions/#institutions-get-response-institutions-oauth)

booleanboolean

Indicates that the institution has an OAuth login flow. This will be `true` if OAuth is supported for any Items associated with the institution, even if the institution also supports non-OAuth connections.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status)

nullableobjectnullable, object

The status of an institution is determined by the health of its Item logins, Transactions updates, Investments updates, Liabilities updates, Auth requests, Balance requests, Identity requests, Investments requests, and Liabilities requests. A login attempt is conducted during the initial Item add in Link. If there is not enough traffic to accurately calculate an institution's status, Plaid will return null rather than potentially inaccurate data.  
Institution status is accessible in the Dashboard and via the API using the `/institutions/get_by_id` endpoint with the `options.include_status` option set to true. Note that institution status is not available in the Sandbox environment.

[`item_logins`](/docs/api/institutions/#institutions-get-response-institutions-status-item-logins)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-item-logins-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get-response-institutions-status-item-logins-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get-response-institutions-status-item-logins-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get-response-institutions-status-item-logins-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get-response-institutions-status-item-logins-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get-response-institutions-status-item-logins-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get-response-institutions-status-item-logins-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`transactions_updates`](/docs/api/institutions/#institutions-get-response-institutions-status-transactions-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-transactions-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get-response-institutions-status-transactions-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get-response-institutions-status-transactions-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get-response-institutions-status-transactions-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get-response-institutions-status-transactions-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get-response-institutions-status-transactions-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get-response-institutions-status-transactions-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`auth`](/docs/api/institutions/#institutions-get-response-institutions-status-auth)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-auth-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get-response-institutions-status-auth-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get-response-institutions-status-auth-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get-response-institutions-status-auth-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get-response-institutions-status-auth-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get-response-institutions-status-auth-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get-response-institutions-status-auth-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`identity`](/docs/api/institutions/#institutions-get-response-institutions-status-identity)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-identity-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get-response-institutions-status-identity-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get-response-institutions-status-identity-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get-response-institutions-status-identity-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get-response-institutions-status-identity-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get-response-institutions-status-identity-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get-response-institutions-status-identity-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`investments_updates`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`liabilities_updates`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`liabilities`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get-response-institutions-status-liabilities-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`investments`](/docs/api/institutions/#institutions-get-response-institutions-status-investments)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get-response-institutions-status-investments-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`health_incidents`](/docs/api/institutions/#institutions-get-response-institutions-status-health-incidents)

nullable[object]nullable, [object]

Details of recent health incidents associated with the institution.

[`start_date`](/docs/api/institutions/#institutions-get-response-institutions-status-health-incidents-start-date)

stringstring

The start date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`end_date`](/docs/api/institutions/#institutions-get-response-institutions-status-health-incidents-end-date)

nullablestringnullable, string

The end date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`title`](/docs/api/institutions/#institutions-get-response-institutions-status-health-incidents-title)

stringstring

The title of the incident

[`incident_updates`](/docs/api/institutions/#institutions-get-response-institutions-status-health-incidents-incident-updates)

[object][object]

Updates on the health incident.

[`description`](/docs/api/institutions/#institutions-get-response-institutions-status-health-incidents-incident-updates-description)

stringstring

The content of the update.

[`status`](/docs/api/institutions/#institutions-get-response-institutions-status-health-incidents-incident-updates-status)

stringstring

The status of the incident.  
  

Possible values: `INVESTIGATING`, `IDENTIFIED`, `SCHEDULED`, `RESOLVED`, `UNKNOWN`

[`updated_date`](/docs/api/institutions/#institutions-get-response-institutions-status-health-incidents-incident-updates-updated-date)

stringstring

The date when the update was published, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`payment_initiation_metadata`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata)

nullableobjectnullable, object

Metadata that captures what specific payment configurations an institution supports when making Payment Initiation requests.

[`supports_international_payments`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-supports-international-payments)

booleanboolean

Indicates whether the institution supports payments from a different country.

[`supports_sepa_instant`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-supports-sepa-instant)

booleanboolean

Indicates whether the institution supports SEPA Instant payments.

[`maximum_payment_amount`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-maximum-payment-amount)

objectobject

A mapping of currency to maximum payment amount (denominated in the smallest unit of currency) supported by the institution.  
Example: `{"GBP": "10000"}`

[`supports_refund_details`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-supports-refund-details)

booleanboolean

Indicates whether the institution supports returning refund details when initiating a payment.

[`standing_order_metadata`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-standing-order-metadata)

nullableobjectnullable, object

Metadata specifically related to valid Payment Initiation standing order configurations for the institution.

[`supports_standing_order_end_date`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-standing-order-metadata-supports-standing-order-end-date)

booleanboolean

Indicates whether the institution supports closed-ended standing orders by providing an end date.

[`supports_standing_order_negative_execution_days`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-standing-order-metadata-supports-standing-order-negative-execution-days)

booleanboolean

This is only applicable to `MONTHLY` standing orders. Indicates whether the institution supports negative integers (-1 to -5) for setting up a `MONTHLY` standing order relative to the end of the month.

[`valid_standing_order_intervals`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-standing-order-metadata-valid-standing-order-intervals)

[string][string]

A list of the valid standing order intervals supported by the institution.  
  

Possible values: `WEEKLY`, `MONTHLY`

Min length: `1`

[`supports_payment_consents`](/docs/api/institutions/#institutions-get-response-institutions-payment-initiation-metadata-supports-payment-consents)

booleanboolean

Indicates whether the institution supports payment consents.

[`auth_metadata`](/docs/api/institutions/#institutions-get-response-institutions-auth-metadata)

nullableobjectnullable, object

Metadata that captures information about the Auth features of an institution.

[`supported_methods`](/docs/api/institutions/#institutions-get-response-institutions-auth-metadata-supported-methods)

nullableobjectnullable, object

Metadata specifically related to which auth methods an institution supports.

[`instant_auth`](/docs/api/institutions/#institutions-get-response-institutions-auth-metadata-supported-methods-instant-auth)

booleanboolean

Indicates if instant auth is supported.

[`instant_match`](/docs/api/institutions/#institutions-get-response-institutions-auth-metadata-supported-methods-instant-match)

booleanboolean

Indicates if instant match is supported.

[`automated_micro_deposits`](/docs/api/institutions/#institutions-get-response-institutions-auth-metadata-supported-methods-automated-micro-deposits)

booleanboolean

Indicates if automated microdeposits are supported.

[`instant_micro_deposits`](/docs/api/institutions/#institutions-get-response-institutions-auth-metadata-supported-methods-instant-micro-deposits)

booleanboolean

Indicates if instant microdeposits are supported.

[`total`](/docs/api/institutions/#institutions-get-response-total)

integerinteger

The total number of institutions available via this endpoint

[`request_id`](/docs/api/institutions/#institutions-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "institutions": [
    {
      "country_codes": [
        "US"
      ],
      "institution_id": "ins_1",
      "name": "Bank of America",
      "products": [
        "assets",
        "auth",
        "balance",
        "transactions",
        "identity",
        "liabilities"
      ],
      "routing_numbers": [
        "011000138",
        "011200365",
        "011400495"
      ],
      "dtc_numbers": [
        "2236",
        "0955",
        "1367"
      ],
      "oauth": false
    }
  ],
  "request_id": "tbFyCEqkU774ZGG",
  "total": 11384
}
```

=\*=\*=\*=

#### `/institutions/get_by_id`

#### Get details of an institution

Returns a JSON response containing details on a specified financial institution currently supported by Plaid.

Versioning note: API versions 2019-05-29 and earlier allow use of the `public_key` parameter instead of the `client_id` and `secret` to authenticate to this endpoint. The `public_key` has been deprecated; all customers are encouraged to use `client_id` and `secret` instead.

/institutions/get\_by\_id

**Request fields**

[`client_id`](/docs/api/institutions/#institutions-get_by_id-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/institutions/#institutions-get_by_id-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`institution_id`](/docs/api/institutions/#institutions-get_by_id-request-institution-id)

requiredstringrequired, string

The ID of the institution to get details about  
  

Min length: `1`

[`country_codes`](/docs/api/institutions/#institutions-get_by_id-request-country-codes)

required[string]required, [string]

Specify which country or countries to include institutions from, using the ISO-3166-1 alpha-2 country code standard. In API versions 2019-05-29 and earlier, the `country_codes` parameter is an optional parameter within the `options` object and will default to `[US]` if it is not supplied.  
  

Possible values: `US`, `GB`, `ES`, `NL`, `FR`, `IE`, `CA`, `DE`, `IT`, `PL`, `DK`, `NO`, `SE`, `EE`, `LT`, `LV`, `PT`, `BE`, `AT`, `FI`

[`options`](/docs/api/institutions/#institutions-get_by_id-request-options)

objectobject

Specifies optional parameters for `/institutions/get_by_id`. If provided, must not be `null`.

[`include_optional_metadata`](/docs/api/institutions/#institutions-get_by_id-request-options-include-optional-metadata)

booleanboolean

When `true`, return an institution's logo, brand color, and URL. When available, the bank's logo is returned as a base64 encoded 152x152 PNG, the brand color is in hexadecimal format. The default value is `false`.  
Note that Plaid does not own any of the logos shared by the API and that by accessing or using these logos, you agree that you are doing so at your own risk and will, if necessary, obtain all required permissions from the appropriate rights holders and adhere to any applicable usage guidelines. Plaid disclaims all express or implied warranties with respect to the logos.  
  

Default: `false`

[`include_status`](/docs/api/institutions/#institutions-get_by_id-request-options-include-status)

booleanboolean

If `true`, the response will include status information about the institution. Default value is `false`.  
  

Default: `false`

[`include_auth_metadata`](/docs/api/institutions/#institutions-get_by_id-request-options-include-auth-metadata)

booleanboolean

When `true`, returns metadata related to the Auth product indicating which auth methods are supported.  
  

Default: `false`

[`include_payment_initiation_metadata`](/docs/api/institutions/#institutions-get_by_id-request-options-include-payment-initiation-metadata)

booleanboolean

When `true`, returns metadata related to the Payment Initiation product indicating which payment configurations are supported.  
  

Default: `false`

/institutions/get\_by\_id

```
const request: InstitutionsGetByIdRequest = {
  institution_id: institutionID,
  country_codes: ['US'],
};
try {
  const response = await plaidClient.institutionsGetById(request);
  const institution = response.data.institution;
} catch (error) {
  // Handle error
}
```

/institutions/get\_by\_id

**Response fields**

[`institution`](/docs/api/institutions/#institutions-get_by_id-response-institution)

objectobject

Details relating to a specific financial institution

[`institution_id`](/docs/api/institutions/#institutions-get_by_id-response-institution-institution-id)

stringstring

Unique identifier for the institution. Note that the same institution may have multiple records, each with different institution IDs; for example, if the institution has migrated to OAuth, there may be separate `institution_id`s for the OAuth and non-OAuth versions of the institution. Institutions that operate in different countries or with multiple login portals may also have separate `institution_id`s for each country or portal.

[`name`](/docs/api/institutions/#institutions-get_by_id-response-institution-name)

stringstring

The official name of the institution.

[`products`](/docs/api/institutions/#institutions-get_by_id-response-institution-products)

[string][string]

A list of the Plaid products supported by the institution. Note that only institutions that support Instant Auth will return `auth` in the product array; institutions that do not list `auth` may still support other Auth methods such as Instant Match or Automated Micro-deposit Verification. To identify institutions that support those methods, use the `auth_metadata` object. For more details, see [Full Auth coverage](https://plaid.com/docs/auth/coverage/). The `income_verification` product here indicates support for Bank Income. Note: For Signal Transaction Scores and Transfer, listed institutions may be incomplete or incorrect. Instead, use the following: `balance` support also indicates coverage of Signal Transaction Scores; `auth` support also indicates coverage of Transfer.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`country_codes`](/docs/api/institutions/#institutions-get_by_id-response-institution-country-codes)

[string][string]

A list of the country codes supported by the institution.  
  

Possible values: `US`, `GB`, `ES`, `NL`, `FR`, `IE`, `CA`, `DE`, `IT`, `PL`, `DK`, `NO`, `SE`, `EE`, `LT`, `LV`, `PT`, `BE`, `AT`, `FI`

[`url`](/docs/api/institutions/#institutions-get_by_id-response-institution-url)

nullablestringnullable, string

The URL for the institution's website

[`primary_color`](/docs/api/institutions/#institutions-get_by_id-response-institution-primary-color)

nullablestringnullable, string

Hexadecimal representation of the primary color used by the institution

[`logo`](/docs/api/institutions/#institutions-get_by_id-response-institution-logo)

nullablestringnullable, string

Base64 encoded representation of the institution's logo, returned as a base64 encoded 152x152 PNG. Not all institutions' logos are available.

[`routing_numbers`](/docs/api/institutions/#institutions-get_by_id-response-institution-routing-numbers)

[string][string]

A list of routing numbers known to be associated with the institution. This list is provided for the purpose of looking up institutions by routing number. It is generally comprehensive but is not guaranteed to be a complete list of routing numbers for an institution.

[`dtc_numbers`](/docs/api/institutions/#institutions-get_by_id-response-institution-dtc-numbers)

[string][string]

A partial list of DTC numbers associated with the institution.

[`oauth`](/docs/api/institutions/#institutions-get_by_id-response-institution-oauth)

booleanboolean

Indicates that the institution has an OAuth login flow. This will be `true` if OAuth is supported for any Items associated with the institution, even if the institution also supports non-OAuth connections.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status)

nullableobjectnullable, object

The status of an institution is determined by the health of its Item logins, Transactions updates, Investments updates, Liabilities updates, Auth requests, Balance requests, Identity requests, Investments requests, and Liabilities requests. A login attempt is conducted during the initial Item add in Link. If there is not enough traffic to accurately calculate an institution's status, Plaid will return null rather than potentially inaccurate data.  
Institution status is accessible in the Dashboard and via the API using the `/institutions/get_by_id` endpoint with the `options.include_status` option set to true. Note that institution status is not available in the Sandbox environment.

[`item_logins`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-item-logins)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-item-logins-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-item-logins-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-item-logins-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-item-logins-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-item-logins-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-item-logins-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-item-logins-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`transactions_updates`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-transactions-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-transactions-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-transactions-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-transactions-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-transactions-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-transactions-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-transactions-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-transactions-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`auth`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-auth)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-auth-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-auth-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-auth-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-auth-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-auth-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-auth-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-auth-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`identity`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-identity)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-identity-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-identity-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-identity-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-identity-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-identity-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-identity-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-identity-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`investments_updates`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`liabilities_updates`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`liabilities`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-liabilities-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`investments`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-investments-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`health_incidents`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-health-incidents)

nullable[object]nullable, [object]

Details of recent health incidents associated with the institution.

[`start_date`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-health-incidents-start-date)

stringstring

The start date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`end_date`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-health-incidents-end-date)

nullablestringnullable, string

The end date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`title`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-health-incidents-title)

stringstring

The title of the incident

[`incident_updates`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-health-incidents-incident-updates)

[object][object]

Updates on the health incident.

[`description`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-health-incidents-incident-updates-description)

stringstring

The content of the update.

[`status`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-health-incidents-incident-updates-status)

stringstring

The status of the incident.  
  

Possible values: `INVESTIGATING`, `IDENTIFIED`, `SCHEDULED`, `RESOLVED`, `UNKNOWN`

[`updated_date`](/docs/api/institutions/#institutions-get_by_id-response-institution-status-health-incidents-incident-updates-updated-date)

stringstring

The date when the update was published, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`payment_initiation_metadata`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata)

nullableobjectnullable, object

Metadata that captures what specific payment configurations an institution supports when making Payment Initiation requests.

[`supports_international_payments`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-supports-international-payments)

booleanboolean

Indicates whether the institution supports payments from a different country.

[`supports_sepa_instant`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-supports-sepa-instant)

booleanboolean

Indicates whether the institution supports SEPA Instant payments.

[`maximum_payment_amount`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-maximum-payment-amount)

objectobject

A mapping of currency to maximum payment amount (denominated in the smallest unit of currency) supported by the institution.  
Example: `{"GBP": "10000"}`

[`supports_refund_details`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-supports-refund-details)

booleanboolean

Indicates whether the institution supports returning refund details when initiating a payment.

[`standing_order_metadata`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-standing-order-metadata)

nullableobjectnullable, object

Metadata specifically related to valid Payment Initiation standing order configurations for the institution.

[`supports_standing_order_end_date`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-standing-order-metadata-supports-standing-order-end-date)

booleanboolean

Indicates whether the institution supports closed-ended standing orders by providing an end date.

[`supports_standing_order_negative_execution_days`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-standing-order-metadata-supports-standing-order-negative-execution-days)

booleanboolean

This is only applicable to `MONTHLY` standing orders. Indicates whether the institution supports negative integers (-1 to -5) for setting up a `MONTHLY` standing order relative to the end of the month.

[`valid_standing_order_intervals`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-standing-order-metadata-valid-standing-order-intervals)

[string][string]

A list of the valid standing order intervals supported by the institution.  
  

Possible values: `WEEKLY`, `MONTHLY`

Min length: `1`

[`supports_payment_consents`](/docs/api/institutions/#institutions-get_by_id-response-institution-payment-initiation-metadata-supports-payment-consents)

booleanboolean

Indicates whether the institution supports payment consents.

[`auth_metadata`](/docs/api/institutions/#institutions-get_by_id-response-institution-auth-metadata)

nullableobjectnullable, object

Metadata that captures information about the Auth features of an institution.

[`supported_methods`](/docs/api/institutions/#institutions-get_by_id-response-institution-auth-metadata-supported-methods)

nullableobjectnullable, object

Metadata specifically related to which auth methods an institution supports.

[`instant_auth`](/docs/api/institutions/#institutions-get_by_id-response-institution-auth-metadata-supported-methods-instant-auth)

booleanboolean

Indicates if instant auth is supported.

[`instant_match`](/docs/api/institutions/#institutions-get_by_id-response-institution-auth-metadata-supported-methods-instant-match)

booleanboolean

Indicates if instant match is supported.

[`automated_micro_deposits`](/docs/api/institutions/#institutions-get_by_id-response-institution-auth-metadata-supported-methods-automated-micro-deposits)

booleanboolean

Indicates if automated microdeposits are supported.

[`instant_micro_deposits`](/docs/api/institutions/#institutions-get_by_id-response-institution-auth-metadata-supported-methods-instant-micro-deposits)

booleanboolean

Indicates if instant microdeposits are supported.

[`request_id`](/docs/api/institutions/#institutions-get_by_id-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "institution": {
    "country_codes": [
      "US"
    ],
    "institution_id": "ins_109512",
    "name": "Houndstooth Bank",
    "products": [
      "auth",
      "balance",
      "identity",
      "transactions"
    ],
    "routing_numbers": [
      "011000138",
      "011200365",
      "011400495"
    ],
    "dtc_numbers": [
      "2236",
      "0955",
      "1367"
    ],
    "oauth": false,
    "status": {
      "item_logins": {
        "status": "HEALTHY",
        "last_status_change": "2019-02-15T15:53:00Z",
        "breakdown": {
          "success": 0.9,
          "error_plaid": 0.01,
          "error_institution": 0.09
        }
      },
      "transactions_updates": {
        "status": "HEALTHY",
        "last_status_change": "2019-02-12T08:22:00Z",
        "breakdown": {
          "success": 0.95,
          "error_plaid": 0.02,
          "error_institution": 0.03,
          "refresh_interval": "NORMAL"
        }
      },
      "auth": {
        "status": "HEALTHY",
        "last_status_change": "2019-02-15T15:53:00Z",
        "breakdown": {
          "success": 0.91,
          "error_plaid": 0.01,
          "error_institution": 0.08
        }
      },
      "identity": {
        "status": "DEGRADED",
        "last_status_change": "2019-02-15T15:50:00Z",
        "breakdown": {
          "success": 0.42,
          "error_plaid": 0.08,
          "error_institution": 0.5
        }
      },
      "investments": {
        "status": "HEALTHY",
        "last_status_change": "2019-02-15T15:53:00Z",
        "breakdown": {
          "success": 0.89,
          "error_plaid": 0.02,
          "error_institution": 0.09
        },
        "liabilities": {
          "status": "HEALTHY",
          "last_status_change": "2019-02-15T15:53:00Z",
          "breakdown": {
            "success": 0.89,
            "error_plaid": 0.02,
            "error_institution": 0.09
          }
        }
      },
      "investments_updates": {
        "status": "HEALTHY",
        "last_status_change": "2019-02-12T08:22:00Z",
        "breakdown": {
          "success": 0.95,
          "error_plaid": 0.02,
          "error_institution": 0.03,
          "refresh_interval": "NORMAL"
        }
      },
      "liabilities_updates": {
        "status": "HEALTHY",
        "last_status_change": "2019-02-12T08:22:00Z",
        "breakdown": {
          "success": 0.95,
          "error_plaid": 0.02,
          "error_institution": 0.03,
          "refresh_interval": "NORMAL"
        }
      }
    },
    "primary_color": "#004966",
    "url": "https://plaid.com",
    "logo": null
  },
  "request_id": "m8MDnv9okwxFNBV"
}
```

=\*=\*=\*=

#### `/institutions/search`

#### Search institutions

Returns a JSON response containing details for institutions that match the query parameters, up to a maximum of ten institutions per query.

Versioning note: API versions 2019-05-29 and earlier allow use of the `public_key` parameter instead of the `client_id` and `secret` parameters to authenticate to this endpoint. The `public_key` parameter has since been deprecated; all customers are encouraged to use `client_id` and `secret` instead.

/institutions/search

**Request fields**

[`client_id`](/docs/api/institutions/#institutions-search-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/institutions/#institutions-search-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`query`](/docs/api/institutions/#institutions-search-request-query)

requiredstringrequired, string

The search query. Institutions with names matching the query are returned  
  

Min length: `1`

[`products`](/docs/api/institutions/#institutions-search-request-products)

[string][string]

Filter the Institutions based on whether they support all products listed in `products`. Provide `null` to get institutions regardless of supported products. Note that when `auth` is specified as a product, if you are enabled for Instant Match or Automated Micro-deposits, institutions that support those products will be returned even if `auth` is not present in their product array. To search for Transfer support, use `auth`; to search for Signal Transaction Scores support, use `balance`.  
  

Min items: `1`

Possible values: `assets`, `auth`, `balance`, `employment`, `identity`, `income_verification`, `investments`, `liabilities`, `identity_verification`, `payment_initiation`, `standing_orders`, `statements`, `transactions`

[`country_codes`](/docs/api/institutions/#institutions-search-request-country-codes)

required[string]required, [string]

Specify which country or countries to include institutions from, using the ISO-3166-1 alpha-2 country code standard. In API versions 2019-05-29 and earlier, the `country_codes` parameter is an optional parameter within the `options` object and will default to `[US]` if it is not supplied.  
  

Possible values: `US`, `GB`, `ES`, `NL`, `FR`, `IE`, `CA`, `DE`, `IT`, `PL`, `DK`, `NO`, `SE`, `EE`, `LT`, `LV`, `PT`, `BE`, `AT`, `FI`

[`options`](/docs/api/institutions/#institutions-search-request-options)

objectobject

An optional object to filter `/institutions/search` results.

[`oauth`](/docs/api/institutions/#institutions-search-request-options-oauth)

booleanboolean

Limit results to institutions with or without OAuth login flows. Note that institutions will have `oauth` set to `true` if some Items associated with that institution are required to use OAuth flows; institutions in a state of migration to OAuth will have the `oauth` attribute set to `true`.

[`include_optional_metadata`](/docs/api/institutions/#institutions-search-request-options-include-optional-metadata)

booleanboolean

When true, return the institution's homepage URL, logo and primary brand color.

[`include_auth_metadata`](/docs/api/institutions/#institutions-search-request-options-include-auth-metadata)

booleanboolean

When `true`, returns metadata related to the Auth product indicating which auth methods are supported.  
  

Default: `false`

[`include_payment_initiation_metadata`](/docs/api/institutions/#institutions-search-request-options-include-payment-initiation-metadata)

booleanboolean

When `true`, returns metadata related to the Payment Initiation product indicating which payment configurations are supported.  
  

Default: `false`

[`payment_initiation`](/docs/api/institutions/#institutions-search-request-options-payment-initiation)

objectobject

Additional options that will be used to filter institutions by various Payment Initiation configurations.

[`payment_id`](/docs/api/institutions/#institutions-search-request-options-payment-initiation-payment-id)

stringstring

A unique ID identifying the payment

[`consent_id`](/docs/api/institutions/#institutions-search-request-options-payment-initiation-consent-id)

stringstring

A unique ID identifying the payment consent

```
const request: InstitutionsSearchRequest = {
  query: institutionID,
  products: ['transactions'],
  country_codes: ['US'],
};
try {
  const response = await plaidClient.institutionsSearch(request);
  const institutions = response.data.institutions;
} catch (error) {
  // Handle error
}
```

/institutions/search

**Response fields**

[`institutions`](/docs/api/institutions/#institutions-search-response-institutions)

[object][object]

An array of institutions matching the search criteria

[`institution_id`](/docs/api/institutions/#institutions-search-response-institutions-institution-id)

stringstring

Unique identifier for the institution. Note that the same institution may have multiple records, each with different institution IDs; for example, if the institution has migrated to OAuth, there may be separate `institution_id`s for the OAuth and non-OAuth versions of the institution. Institutions that operate in different countries or with multiple login portals may also have separate `institution_id`s for each country or portal.

[`name`](/docs/api/institutions/#institutions-search-response-institutions-name)

stringstring

The official name of the institution.

[`products`](/docs/api/institutions/#institutions-search-response-institutions-products)

[string][string]

A list of the Plaid products supported by the institution. Note that only institutions that support Instant Auth will return `auth` in the product array; institutions that do not list `auth` may still support other Auth methods such as Instant Match or Automated Micro-deposit Verification. To identify institutions that support those methods, use the `auth_metadata` object. For more details, see [Full Auth coverage](https://plaid.com/docs/auth/coverage/). The `income_verification` product here indicates support for Bank Income. Note: For Signal Transaction Scores and Transfer, listed institutions may be incomplete or incorrect. Instead, use the following: `balance` support also indicates coverage of Signal Transaction Scores; `auth` support also indicates coverage of Transfer.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`country_codes`](/docs/api/institutions/#institutions-search-response-institutions-country-codes)

[string][string]

A list of the country codes supported by the institution.  
  

Possible values: `US`, `GB`, `ES`, `NL`, `FR`, `IE`, `CA`, `DE`, `IT`, `PL`, `DK`, `NO`, `SE`, `EE`, `LT`, `LV`, `PT`, `BE`, `AT`, `FI`

[`url`](/docs/api/institutions/#institutions-search-response-institutions-url)

nullablestringnullable, string

The URL for the institution's website

[`primary_color`](/docs/api/institutions/#institutions-search-response-institutions-primary-color)

nullablestringnullable, string

Hexadecimal representation of the primary color used by the institution

[`logo`](/docs/api/institutions/#institutions-search-response-institutions-logo)

nullablestringnullable, string

Base64 encoded representation of the institution's logo, returned as a base64 encoded 152x152 PNG. Not all institutions' logos are available.

[`routing_numbers`](/docs/api/institutions/#institutions-search-response-institutions-routing-numbers)

[string][string]

A list of routing numbers known to be associated with the institution. This list is provided for the purpose of looking up institutions by routing number. It is generally comprehensive but is not guaranteed to be a complete list of routing numbers for an institution.

[`dtc_numbers`](/docs/api/institutions/#institutions-search-response-institutions-dtc-numbers)

[string][string]

A partial list of DTC numbers associated with the institution.

[`oauth`](/docs/api/institutions/#institutions-search-response-institutions-oauth)

booleanboolean

Indicates that the institution has an OAuth login flow. This will be `true` if OAuth is supported for any Items associated with the institution, even if the institution also supports non-OAuth connections.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status)

nullableobjectnullable, object

The status of an institution is determined by the health of its Item logins, Transactions updates, Investments updates, Liabilities updates, Auth requests, Balance requests, Identity requests, Investments requests, and Liabilities requests. A login attempt is conducted during the initial Item add in Link. If there is not enough traffic to accurately calculate an institution's status, Plaid will return null rather than potentially inaccurate data.  
Institution status is accessible in the Dashboard and via the API using the `/institutions/get_by_id` endpoint with the `options.include_status` option set to true. Note that institution status is not available in the Sandbox environment.

[`item_logins`](/docs/api/institutions/#institutions-search-response-institutions-status-item-logins)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-item-logins-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-search-response-institutions-status-item-logins-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-search-response-institutions-status-item-logins-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-search-response-institutions-status-item-logins-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-search-response-institutions-status-item-logins-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-search-response-institutions-status-item-logins-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-search-response-institutions-status-item-logins-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`transactions_updates`](/docs/api/institutions/#institutions-search-response-institutions-status-transactions-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-transactions-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-search-response-institutions-status-transactions-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-search-response-institutions-status-transactions-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-search-response-institutions-status-transactions-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-search-response-institutions-status-transactions-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-search-response-institutions-status-transactions-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-search-response-institutions-status-transactions-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`auth`](/docs/api/institutions/#institutions-search-response-institutions-status-auth)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-auth-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-search-response-institutions-status-auth-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-search-response-institutions-status-auth-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-search-response-institutions-status-auth-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-search-response-institutions-status-auth-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-search-response-institutions-status-auth-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-search-response-institutions-status-auth-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`identity`](/docs/api/institutions/#institutions-search-response-institutions-status-identity)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-identity-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-search-response-institutions-status-identity-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-search-response-institutions-status-identity-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-search-response-institutions-status-identity-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-search-response-institutions-status-identity-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-search-response-institutions-status-identity-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-search-response-institutions-status-identity-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`investments_updates`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`liabilities_updates`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-updates)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-updates-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-updates-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-updates-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-updates-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-updates-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-updates-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-updates-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`liabilities`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-search-response-institutions-status-liabilities-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`investments`](/docs/api/institutions/#institutions-search-response-institutions-status-investments)

nullableobjectnullable, object

A representation of the status health of a request type. Auth requests, Balance requests, Identity requests, Investments requests, Liabilities requests, Transactions updates, Investments updates, Liabilities updates, and Item logins each have their own status object.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-status)

deprecatedstringdeprecated, string

This field is deprecated in favor of the `breakdown` object, which provides more granular institution health data.  
`HEALTHY`: the majority of requests are successful
`DEGRADED`: only some requests are successful
`DOWN`: all requests are failing  
  

Possible values: `HEALTHY`, `DEGRADED`, `DOWN`

[`last_status_change`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-last-status-change)

stringstring

[ISO 8601](https://wikipedia.org/wiki/ISO_8601) formatted timestamp of the last status change for the institution.  
  

Format: `date-time`

[`breakdown`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-breakdown)

objectobject

A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).

[`success`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-breakdown-success)

numbernumber

The percentage of login attempts that are successful, expressed as a decimal.  
  

Format: `double`

[`error_plaid`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-breakdown-error-plaid)

numbernumber

The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  
  

Format: `double`

[`error_institution`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-breakdown-error-institution)

numbernumber

The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.  
  

Format: `double`

[`refresh_interval`](/docs/api/institutions/#institutions-search-response-institutions-status-investments-breakdown-refresh-interval)

stringstring

How frequently data for subscription products like Investments, Transactions, and Liabilities, is being refreshed, relative to the institution's normal scheduling. The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high.  
  

Possible values: `NORMAL`, `DELAYED`, `STOPPED`

[`health_incidents`](/docs/api/institutions/#institutions-search-response-institutions-status-health-incidents)

nullable[object]nullable, [object]

Details of recent health incidents associated with the institution.

[`start_date`](/docs/api/institutions/#institutions-search-response-institutions-status-health-incidents-start-date)

stringstring

The start date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`end_date`](/docs/api/institutions/#institutions-search-response-institutions-status-health-incidents-end-date)

nullablestringnullable, string

The end date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`title`](/docs/api/institutions/#institutions-search-response-institutions-status-health-incidents-title)

stringstring

The title of the incident

[`incident_updates`](/docs/api/institutions/#institutions-search-response-institutions-status-health-incidents-incident-updates)

[object][object]

Updates on the health incident.

[`description`](/docs/api/institutions/#institutions-search-response-institutions-status-health-incidents-incident-updates-description)

stringstring

The content of the update.

[`status`](/docs/api/institutions/#institutions-search-response-institutions-status-health-incidents-incident-updates-status)

stringstring

The status of the incident.  
  

Possible values: `INVESTIGATING`, `IDENTIFIED`, `SCHEDULED`, `RESOLVED`, `UNKNOWN`

[`updated_date`](/docs/api/institutions/#institutions-search-response-institutions-status-health-incidents-incident-updates-updated-date)

stringstring

The date when the update was published, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2020-10-30T15:26:48Z"`.  
  

Format: `date-time`

[`payment_initiation_metadata`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata)

nullableobjectnullable, object

Metadata that captures what specific payment configurations an institution supports when making Payment Initiation requests.

[`supports_international_payments`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-supports-international-payments)

booleanboolean

Indicates whether the institution supports payments from a different country.

[`supports_sepa_instant`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-supports-sepa-instant)

booleanboolean

Indicates whether the institution supports SEPA Instant payments.

[`maximum_payment_amount`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-maximum-payment-amount)

objectobject

A mapping of currency to maximum payment amount (denominated in the smallest unit of currency) supported by the institution.  
Example: `{"GBP": "10000"}`

[`supports_refund_details`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-supports-refund-details)

booleanboolean

Indicates whether the institution supports returning refund details when initiating a payment.

[`standing_order_metadata`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-standing-order-metadata)

nullableobjectnullable, object

Metadata specifically related to valid Payment Initiation standing order configurations for the institution.

[`supports_standing_order_end_date`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-standing-order-metadata-supports-standing-order-end-date)

booleanboolean

Indicates whether the institution supports closed-ended standing orders by providing an end date.

[`supports_standing_order_negative_execution_days`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-standing-order-metadata-supports-standing-order-negative-execution-days)

booleanboolean

This is only applicable to `MONTHLY` standing orders. Indicates whether the institution supports negative integers (-1 to -5) for setting up a `MONTHLY` standing order relative to the end of the month.

[`valid_standing_order_intervals`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-standing-order-metadata-valid-standing-order-intervals)

[string][string]

A list of the valid standing order intervals supported by the institution.  
  

Possible values: `WEEKLY`, `MONTHLY`

Min length: `1`

[`supports_payment_consents`](/docs/api/institutions/#institutions-search-response-institutions-payment-initiation-metadata-supports-payment-consents)

booleanboolean

Indicates whether the institution supports payment consents.

[`auth_metadata`](/docs/api/institutions/#institutions-search-response-institutions-auth-metadata)

nullableobjectnullable, object

Metadata that captures information about the Auth features of an institution.

[`supported_methods`](/docs/api/institutions/#institutions-search-response-institutions-auth-metadata-supported-methods)

nullableobjectnullable, object

Metadata specifically related to which auth methods an institution supports.

[`instant_auth`](/docs/api/institutions/#institutions-search-response-institutions-auth-metadata-supported-methods-instant-auth)

booleanboolean

Indicates if instant auth is supported.

[`instant_match`](/docs/api/institutions/#institutions-search-response-institutions-auth-metadata-supported-methods-instant-match)

booleanboolean

Indicates if instant match is supported.

[`automated_micro_deposits`](/docs/api/institutions/#institutions-search-response-institutions-auth-metadata-supported-methods-automated-micro-deposits)

booleanboolean

Indicates if automated microdeposits are supported.

[`instant_micro_deposits`](/docs/api/institutions/#institutions-search-response-institutions-auth-metadata-supported-methods-instant-micro-deposits)

booleanboolean

Indicates if instant microdeposits are supported.

[`request_id`](/docs/api/institutions/#institutions-search-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "institutions": [
    {
      "country_codes": [
        "US"
      ],
      "institution_id": "ins_109513",
      "name": "Theoretical Bank",
      "oauth": true,
      "products": [
        "assets",
        "auth",
        "balance",
        "cra_lend_score",
        "cra_plaid_credit_score",
        "identity",
        "identity_match",
        "income",
        "pay_by_bank",
        "processor_payments",
        "recurring_transactions",
        "transactions",
        "transfer"
      ],
      "routing_numbers": [
        "031101270",
        "103100194",
        "103112357"
      ]
    }
  ],
  "request_id": "QheuqaazREmq9xp"
}
```

### Webhooks

Institution status alerts are configured within the [developer dashboard](https://dashboard.plaid.com/settings/team/notification-preferences). In the dashboard, you can choose to receive alerts as either emails or webhooks.

All dashboard-configured institution status alerts will have type `DASHBOARD_CONFIGURED_ALERT`.

=\*=\*=\*=

#### `INSTITUTION_STATUS_ALERT_TRIGGERED`

Fired when institution status meets the conditions configured in the developer dashboard.

**Properties**

[`webhook_type`](/docs/api/institutions/#InstitutionStatusAlertWebhook-webhook-type)

stringstring

`DASHBOARD_CONFIGURED_ALERT`

[`webhook_code`](/docs/api/institutions/#InstitutionStatusAlertWebhook-webhook-code)

stringstring

`INSTITUTION_STATUS_ALERT_TRIGGERED`

[`institution_id`](/docs/api/institutions/#InstitutionStatusAlertWebhook-institution-id)

stringstring

The ID of the associated institution.

[`institution_overall_success_rate`](/docs/api/institutions/#InstitutionStatusAlertWebhook-institution-overall-success-rate)

numbernumber

The global success rate of the institution, calculated based on item add health.  
  

Format: `double`

[`environment`](/docs/api/institutions/#InstitutionStatusAlertWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "DASHBOARD_CONFIGURED_ALERT",
  "webhook_code": "INSTITUTION_STATUS_ALERT_TRIGGERED",
  "institution_id": "ins_3",
  "institution_overall_success_rate": 0.9,
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
