---
title: "API - Income | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/income/"
scraped_at: "2026-03-07T22:04:13+00:00"
---

# Income

#### API reference for Income endpoints and webhooks

Verify a user's income via payroll or bank account data. For how-to guidance, see the [Income documentation](/docs/income/).

| Endpoints |  |
| --- | --- |
| [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) | Get Link session metadata and results for the end user |
| [`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) | Retrieve information from the bank accounts used for income verification |
| [`/credit/bank_income/pdf/get`](/docs/api/products/income/#creditbank_incomepdfget) | Retrieve information from the bank accounts used for income verification in PDF format |
| [`/credit/bank_income/refresh`](/docs/api/products/income/#creditbank_incomerefresh) | Refresh a user's bank income information |
| [`/credit/bank_statements/uploads/get`](/docs/api/products/income/#creditbank_statementsuploadsget) | Retrieve information from the bank statements used for income verification |
| [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) | Retrieve information from the pay stubs or tax forms used for income verification |
| [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) | Analyze uploaded income documents for indications of potential fraud |
| [`/credit/payroll_income/parsing_config/update`](/docs/api/products/income/#creditpayroll_incomeparsing_configupdate) | Update the parsing configuration for a document verification |
| [`/credit/employment/get`](/docs/api/products/income/#creditemploymentget) | (Beta) Retrieve employment information about the end user |
| [`/credit/payroll_income/refresh`](/docs/api/products/income/#creditpayroll_incomerefresh) | (Beta) Retrieve updated payroll income data on a linked account |

| See also |  |
| --- | --- |
| [`/sandbox/income/fire_webhook`](/docs/api/sandbox/#sandboxincomefire_webhook) | Manually fire an Income webhook (Sandbox only) |
| [`/user/create`](/docs/api/users/#usercreate) | Create a user for use with the income verification product |

| Webhooks |  |
| --- | --- |
| [`INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) | Income verification has completed |
| [`INCOME_VERIFICATION_RISK_SIGNALS`](/docs/api/products/income/#income_verification_risk_signals) | Risk evaluation of user-uploaded documents has completed |
| [`BANK_INCOME_REFRESH_COMPLETE`](/docs/api/products/income/#bank_income_refresh_complete) | The refreshed report has finished generating |
| [`INCOME_VERIFICATION_REFRESH_RECONNECT_NEEDED`](/docs/api/products/income/#income_verification_refresh_reconnect_needed) | A Payroll Income verification could not be refreshed |

### Endpoints

=\*=\*=\*=

#### `/user/create`

#### Create user

For Plaid products and flows that use the user object, [`/user/create`](/docs/api/users/#usercreate) provides you a single token to access all data associated with the user. You must call this endpoint before calling [`/link/token/create`](/docs/api/link/#linktokencreate) if you are using any of the following: Plaid Check, Income Verification, Multi-Item Link, or Plaid Protect.

For customers who began using this endpoint on or after December 10, 2025, this endpoint takes a `client_user_id` and an `identity` object and will return a `user_id`. For customers who began using it before that date, the endpoint takes a `client_user_id` and a `consumer_report_user_identity` object and will return a `user_token` and `user_id`. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

In order to create a Plaid Check Consumer Report for a user, the `identity` (new) or `consumer_report_user_identity` (legacy) object must be present. If it is not provided during the [`/user/create`](/docs/api/users/#usercreate) call, it can be added later by calling [`/user/update`](/docs/api/users/#userupdate).

In order to generate a Plaid Check Consumer Report, the following `identity` fields, at minimum, are required and must be non-empty: `name`, `date_of_birth`, `emails`, `phone_numbers`, and `addresses`, (with at least one email, phone number, and address designated as `primary`). Plaid Check Consumer Reports can only be created for US-based users; the user's address country must be `US`. If creating a report for sharing with a GSE such as Fannie or Freddie, the user's full SSN must be provided via the `id_numbers` field. Providing at least a partial SSN is also strongly recommended for all use cases, since it improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests.

When using Plaid Protect, it is highly recommended that you provide an `identity` object to better identify and block fraud across your Link sessions.

Plaid will normalize identity fields before storing them and utilize the same identity across different user-based products.

/user/create

**Request fields**

[`client_id`](/docs/api/products/income/#user-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#user-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_user_id`](/docs/api/products/income/#user-create-request-client-user-id)

requiredstringrequired, string

A unique ID representing the end user. Maximum of 128 characters. Typically this will be a user ID number from your application. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.  
  

Max length: `128`

Min length: `1`

[`identity`](/docs/api/products/income/#user-create-request-identity)

objectobject

The identity fields associated with a user. For a user to be eligible for a Plaid Check Consumer Report, all fields are required except `id_number`. Providing a partial SSN is strongly recommended, and improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests. If creating a report that will be shared with GSEs such as Fannie or Freddie, a full Social Security Number must be provided via the `id_number` field.

[`name`](/docs/api/products/income/#user-create-request-identity-name)

objectobject

User name information.

[`given_name`](/docs/api/products/income/#user-create-request-identity-name-given-name)

requiredstringrequired, string

User's given name.

[`family_name`](/docs/api/products/income/#user-create-request-identity-name-family-name)

requiredstringrequired, string

User's family name.

[`date_of_birth`](/docs/api/products/income/#user-create-request-identity-date-of-birth)

stringstring

The user's date of birth, to be provided in the format "yyyy-mm-dd".  
  

Format: `date`

[`emails`](/docs/api/products/income/#user-create-request-identity-emails)

[object][object]

The user's emails.

[`data`](/docs/api/products/income/#user-create-request-identity-emails-data)

requiredstringrequired, string

User's email.

[`primary`](/docs/api/products/income/#user-create-request-identity-emails-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary email for the User.

[`phone_numbers`](/docs/api/products/income/#user-create-request-identity-phone-numbers)

[object][object]

The user's phone numbers, in E.164 format: +{countrycode}{number}. For example: "+14157452130". Phone numbers provided in other formats will be parsed on a best-effort basis. Phone number input is validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`data`](/docs/api/products/income/#user-create-request-identity-phone-numbers-data)

requiredstringrequired, string

User's phone number.

[`primary`](/docs/api/products/income/#user-create-request-identity-phone-numbers-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary phone number for the User.

[`addresses`](/docs/api/products/income/#user-create-request-identity-addresses)

[object][object]

The user's addresses.

[`street_1`](/docs/api/products/income/#user-create-request-identity-addresses-street-1)

stringstring

First line of street address.

[`street_2`](/docs/api/products/income/#user-create-request-identity-addresses-street-2)

stringstring

Second line of street address.

[`city`](/docs/api/products/income/#user-create-request-identity-addresses-city)

stringstring

City name.

[`region`](/docs/api/products/income/#user-create-request-identity-addresses-region)

stringstring

State, province or region.

[`country`](/docs/api/products/income/#user-create-request-identity-addresses-country)

requiredstringrequired, string

Country code.

[`postal_code`](/docs/api/products/income/#user-create-request-identity-addresses-postal-code)

stringstring

Postal or ZIP code.

[`primary`](/docs/api/products/income/#user-create-request-identity-addresses-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary address for the User.

[`id_numbers`](/docs/api/products/income/#user-create-request-identity-id-numbers)

[object][object]

The user's ID numbers.

[`value`](/docs/api/products/income/#user-create-request-identity-id-numbers-value)

requiredstringrequired, string

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/income/#user-create-request-identity-id-numbers-type)

requiredstringrequired, string

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`consumer_report_user_identity`](/docs/api/products/income/#user-create-request-consumer-report-user-identity)

objectobject

This field is only used by integrations created before December 10, 2025. All other integrations must use the `identity` object instead. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).
To create a Plaid Check Consumer Report for a user when using a `user_token`, this field must be present. If this field is not provided during user token creation, you can add it to the user later by calling `/user/update`. Once the field has been added to the user, you will be able to call `/link/token/create` with a non-empty `consumer_report_permissible_purpose` (which will automatically create a Plaid Check Consumer Report), or call `/cra/check_report/create` for that user.

[`first_name`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-first-name)

requiredstringrequired, string

The user's first name

[`last_name`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-last-name)

requiredstringrequired, string

The user's last name

[`phone_numbers`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-phone-numbers)

required[string]required, [string]

The user's phone number, in E.164 format: +{countrycode}{number}. For example: "+14157452130". Phone numbers provided in other formats will be parsed on a best-effort basis. Phone number input is validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`emails`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-emails)

required[string]required, [string]

The user's emails

[`ssn_full`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-ssn-full)

stringstring

The user's full social security number. This field should only be provided by lenders intending to share the resulting consumer report with a Government-Sponsored Enterprise (GSE), such as Fannie Mae or Freddie Mac.  
Format: "ddd-dd-dddd"

[`ssn_last_4`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-ssn-last-4)

stringstring

The last 4 digits of the user's social security number.  
  

Max length: `4`

Min length: `4`

[`date_of_birth`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-date-of-birth)

requiredstringrequired, string

To be provided in the format "yyyy-mm-dd".
This field is required for all Plaid Check customers.  
  

Format: `date`

[`primary_address`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-primary-address)

requiredobjectrequired, object

Data about the components comprising an address.

[`city`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-primary-address-city)

requiredstringrequired, string

The full city name

[`region`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-primary-address-region)

requiredstringrequired, string

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-primary-address-street)

requiredstringrequired, string

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-primary-address-postal-code)

requiredstringrequired, string

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/income/#user-create-request-consumer-report-user-identity-primary-address-country)

requiredstringrequired, string

The ISO 3166-1 alpha-2 country code

/user/create

```
const request: UserCreateRequest = {
  client_user_id: 'c0e2c4ee-b763-4af5-cfe9-46a46bce883d',
};

try {
  const response = await client.userCreate(request);
} catch (error) {
  // handle error
}
```

/user/create

**Response fields**

[`user_token`](/docs/api/products/income/#user-create-response-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`user_id`](/docs/api/products/income/#user-create-response-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`request_id`](/docs/api/products/income/#user-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "user_id": "usr_9nSp2KuZ2x4JDw",
  "request_id": "Aim3b"
}
```

=\*=\*=\*=

#### `/credit/sessions/get`

#### Retrieve Link sessions for your user

This endpoint can be used for your end users after they complete the Link flow. This endpoint returns a list of Link sessions that your user completed, where each session includes the results from the Link flow.

These results include details about the Item that was created and some product related metadata (showing, for example, whether the user finished the bank income verification step).

/credit/sessions/get

**Request fields**

[`client_id`](/docs/api/products/income/#credit-sessions-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-sessions-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-sessions-get-request-user-token)

requiredstringrequired, string

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/credit/sessions/get

```
const request: CreditSessionsGetRequest = {
  user_token: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
};

try {
  const response = await client.creditSessionsGet(request);
} catch (error) {
  // handle error
}
```

/credit/sessions/get

**Response fields**

[`sessions`](/docs/api/products/income/#credit-sessions-get-response-sessions)

[object][object]

A list of Link sessions for the user. Sessions will be sorted in reverse chronological order.

[`link_session_id`](/docs/api/products/income/#credit-sessions-get-response-sessions-link-session-id)

stringstring

The unique identifier associated with the Link session. This identifier matches the `link_session_id` returned in the onSuccess/onExit callbacks.

[`session_start_time`](/docs/api/products/income/#credit-sessions-get-response-sessions-session-start-time)

stringstring

The time when the Link session started  
  

Format: `date-time`

[`results`](/docs/api/products/income/#credit-sessions-get-response-sessions-results)

objectobject

The set of results for a Link session.

[`item_add_results`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-item-add-results)

[object][object]

The set of Item adds for the Link session.

[`public_token`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-item-add-results-public-token)

stringstring

Returned once a user has successfully linked their Item.

[`item_id`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-item-add-results-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`institution_id`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-item-add-results-institution-id)

stringstring

The Plaid Institution ID associated with the Item.

[`bank_income_results`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-bank-income-results)

[object][object]

The set of bank income verifications for the Link session.

[`status`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-bank-income-results-status)

stringstring

Status of the Bank Income Link session.  
`APPROVED`: User has approved and verified their income  
`NO_DEPOSITS_FOUND`: We attempted, but were unable to find any income in the connected account.  
`USER_REPORTED_NO_INCOME`: The user explicitly indicated that they don't receive income in the connected account.  
`STARTED`: The user began the bank income portion of the link flow.  
`INTERNAL_ERROR`: The user encountered an internal error.  
  

Possible values: `APPROVED`, `NO_DEPOSITS_FOUND`, `USER_REPORTED_NO_INCOME`

[`item_id`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-bank-income-results-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`institution_id`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-bank-income-results-institution-id)

stringstring

The Plaid Institution ID associated with the Item.

[`bank_employment_results`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-bank-employment-results)

[object][object]

The set of bank employment verifications for the Link session.

[`status`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-bank-employment-results-status)

stringstring

Status of the Bank Employment Link session.  
`APPROVED`: User has approved and verified their employment.  
`NO_EMPLOYMENTS_FOUND`: We attempted, but were unable to find any employment in the connected account.  
`EMPLOYER_NOT_LISTED`: The user explicitly indicated that they did not see their current or previous employer in the list of employer names found.  
`STARTED`: The user began the bank income portion of the link flow.  
`INTERNAL_ERROR`: The user encountered an internal error.  
  

Possible values: `APPROVED`, `NO_EMPLOYERS_FOUND`, `EMPLOYER_NOT_LISTED`

[`item_id`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-bank-employment-results-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`institution_id`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-bank-employment-results-institution-id)

stringstring

The Plaid Institution ID associated with the Item.

[`payroll_income_results`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-payroll-income-results)

[object][object]

The set of payroll income verifications for the Link session.

[`num_paystubs_retrieved`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-payroll-income-results-num-paystubs-retrieved)

integerinteger

The number of paystubs retrieved from a payroll provider.

[`num_w2s_retrieved`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-payroll-income-results-num-w2s-retrieved)

integerinteger

The number of w2s retrieved from a payroll provider.

[`institution_id`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-payroll-income-results-institution-id)

stringstring

The Plaid Institution ID associated with the Item.

[`institution_name`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-payroll-income-results-institution-name)

stringstring

The Institution Name associated with the Item.

[`document_income_results`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-document-income-results)

nullableobjectnullable, object

The details of a document income verification in Link

[`num_paystubs_uploaded`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-document-income-results-num-paystubs-uploaded)

integerinteger

The number of paystubs uploaded by the user.

[`num_w2s_uploaded`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-document-income-results-num-w2s-uploaded)

integerinteger

The number of w2s uploaded by the user.

[`num_bank_statements_uploaded`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-document-income-results-num-bank-statements-uploaded)

integerinteger

The number of bank statements uploaded by the user.

[`num_1099s_uploaded`](/docs/api/products/income/#credit-sessions-get-response-sessions-results-document-income-results-num-1099s-uploaded)

integerinteger

The number of 1099s uploaded by the user

[`errors`](/docs/api/products/income/#credit-sessions-get-response-sessions-errors)

[object][object]

The set of errors that occurred during the Link session.

[`error_type`](/docs/api/products/income/#credit-sessions-get-response-sessions-errors-error-type)

stringstring

A broad categorization of the error.

[`error_code`](/docs/api/products/income/#credit-sessions-get-response-sessions-errors-error-code)

stringstring

The particular error code.

[`error_message`](/docs/api/products/income/#credit-sessions-get-response-sessions-errors-error-message)

stringstring

A developer-friendly representation of the error code.

[`display_message`](/docs/api/products/income/#credit-sessions-get-response-sessions-errors-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.

[`request_id`](/docs/api/products/income/#credit-sessions-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "Aim3b",
  "sessions": [
    {
      "link_session_id": "356dbb28-7f98-44d1-8e6d-0cec580f3171",
      "results": {
        "item_add_results": [
          {
            "public_token": "public-sandbox-5c224a01-8314-4491-a06f-39e193d5cddc",
            "item_id": "M5eVJqLnv3tbzdngLDp9FL5OlDNxlNhlE55op",
            "institution_id": "ins_56"
          }
        ],
        "bank_income_results": [
          {
            "status": "APPROVED",
            "item_id": "M5eVJqLnv3tbzdngLDp9FL5OlDNxlNhlE55op",
            "institution_id": "ins_56"
          }
        ]
      },
      "session_start_time": "2022-09-30T23:40:30.946225Z"
    },
    {
      "link_session_id": "f742cae8-31e4-49cc-a621-6cafbdb26fb9",
      "results": {
        "payroll_income_results": [
          {
            "num_paystubs_retrieved": 2,
            "num_w2s_retrieved": 1,
            "institution_id": "ins_92"
          }
        ]
      },
      "session_start_time": "2022-09-26T23:40:30.946225Z"
    }
  ]
}
```

=\*=\*=\*=

#### `/credit/bank_income/get`

#### Retrieve information from the bank accounts used for income verification

[`/credit/bank_income/get`](/docs/api/products/income/#creditbank_incomeget) returns the bank income report(s) for a specified user. A single report corresponds to all institutions linked in a single Link session. To include multiple institutions in a single report, use [Multi-Item Link](https://plaid.com/docs/link/multi-item-link). To return older reports, use the `options.count` field.

/credit/bank\_income/get

**Request fields**

[`client_id`](/docs/api/products/income/#credit-bank_income-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-bank_income-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-bank_income-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`options`](/docs/api/products/income/#credit-bank_income-get-request-options)

objectobject

An optional object for `/credit/bank_income/get` request options.

[`count`](/docs/api/products/income/#credit-bank_income-get-request-options-count)

integerinteger

How many Bank Income Reports should be fetched. Multiple reports may be available if the report has been re-created or refreshed. If more than one report is available, the most recent reports will be returned first.  
  

Default: `1`

/credit/bank\_income/get

```
const request: CreditBankIncomeGetRequest = {
  user_token: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
  options: {
    count: 1,
  },
};

try {
  const response = await client.creditBankIncomeGet(request);
} catch (error) {
  // handle error
}
```

/credit/bank\_income/get

**Response fields**

[`bank_income`](/docs/api/products/income/#credit-bank_income-get-response-bank-income)

[object][object]

[`bank_income_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-id)

stringstring

The unique identifier associated with the Bank Income Report.

[`generated_time`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-generated-time)

stringstring

The time when the report was generated.  
  

Format: `date-time`

[`days_requested`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-days-requested)

integerinteger

The number of days requested by the customer for the report.

[`items`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items)

[object][object]

The list of Items in the report along with the associated metadata about the Item.

[`bank_income_accounts`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts)

[object][object]

The Item's accounts that have Bank Income data.

[`account_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-account-id)

stringstring

Plaid's unique identifier for the account.

[`mask`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number.
Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.

[`name`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-name)

stringstring

The name of the bank account.

[`official_name`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-official-name)

nullablestringnullable, string

The official name of the bank account.

[`subtype`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-subtype)

stringstring

Valid account subtypes for depository accounts. For a list containing descriptions of each subtype, see [Account schemas](https://plaid.com/docs/api/accounts/#StandaloneAccountType-depository).  
  

Possible values: `checking`, `savings`, `hsa`, `cd`, `money market`, `paypal`, `prepaid`, `cash management`, `ebt`, `all`

[`type`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-type)

stringstring

The account type. This will always be `depository`.  
  

Possible values: `depository`

[`owners`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners)

[object][object]

Data returned by the financial institution about the account owner or owners. Identity information is optional, so field may return an empty array.

[`names`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-names)

[string][string]

A list of names associated with the account by the financial institution. In the case of a joint account, Plaid will make a best effort to report the names of all account holders.  
If an Item contains multiple accounts with different owner names, some institutions will report all names associated with the Item in each account's `names` array.

[`phone_numbers`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-phone-numbers)

[object][object]

A list of phone numbers associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-phone-numbers-data)

stringstring

The phone number.

[`primary`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-phone-numbers-primary)

booleanboolean

When `true`, identifies the phone number as the primary number on an account.

[`type`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-phone-numbers-type)

stringstring

The type of phone number.  
  

Possible values: `home`, `work`, `office`, `mobile`, `mobile1`, `other`

[`emails`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-emails)

[object][object]

A list of email addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-emails-data)

stringstring

The email address.

[`primary`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-emails-primary)

booleanboolean

When `true`, identifies the email address as the primary email on an account.

[`type`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-emails-type)

stringstring

The type of email account as described by the financial institution.  
  

Possible values: `primary`, `secondary`, `other`

[`addresses`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-addresses)

[object][object]

Data about the various addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.

[`data`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-addresses-data)

objectobject

Data about the components comprising an address.

[`city`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-addresses-data-city)

nullablestringnullable, string

The full city name

[`region`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-addresses-data-region)

nullablestringnullable, string

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-addresses-data-street)

stringstring

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-addresses-data-postal-code)

nullablestringnullable, string

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-addresses-data-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code

[`primary`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-accounts-owners-addresses-primary)

booleanboolean

When `true`, identifies the address as the primary address on an account.

[`bank_income_sources`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources)

[object][object]

The income sources for this Item. Each entry in the array is a single income source.

[`income_source_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-income-source-id)

stringstring

A unique identifier for an income source.

[`income_description`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-income-description)

stringstring

The most common name or original description for the underlying income transactions.

[`income_category`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-income-category)

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

[`account_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-account-id)

stringstring

Plaid's unique identifier for the account.

[`start_date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-start-date)

stringstring

Minimum of all dates within the specific income sources in the user's bank account for days requested by the client.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-end-date)

stringstring

Maximum of all dates within the specific income sources in the user’s bank account for days requested by the client.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`pay_frequency`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-pay-frequency)

stringstring

The income pay frequency.  
  

Possible values: `WEEKLY`, `BIWEEKLY`, `SEMI_MONTHLY`, `MONTHLY`, `DAILY`, `UNKNOWN`

[`total_amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-total-amount)

numbernumber

Total amount of earnings in the user’s bank account for the specific income source for days requested by the client.

[`transaction_count`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-transaction-count)

integerinteger

Number of transactions for the income source within the start and end date.

[`historical_summary`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary)

[object][object]

[`total_amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amount)

deprecatednumberdeprecated, number

Total amount of earnings for the income source(s) of the user for the month in the summary.
This may return an incorrect value if the summary includes income sources in multiple currencies.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts) instead.

[`iso_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-iso-currency-code)

deprecatednullablestringdeprecated, nullable, string

The ISO 4217 currency code of the amount or balance.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts) instead.

[`unofficial_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-unofficial-currency-code)

deprecatednullablestringdeprecated, nullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts) instead.

[`total_amounts`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts)

[object][object]

Total amount of earnings for the income source(s) of the user for the month in the summary.
This can contain multiple amounts, with each amount denominated in one unique currency.

[`amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`start_date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-start-date)

stringstring

The start date of the period covered in this monthly summary.
This date will be the first day of the month, unless the month being covered is a partial month because it is the first month included in the summary and the date range being requested does not begin with the first day of the month.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-end-date)

stringstring

The end date of the period included in this monthly summary.
This date will be the last day of the month, unless the month being covered is a partial month because it is the last month included in the summary and the date range being requested does not end with the last day of the month.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`transactions`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions)

[object][object]

[`amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-amount)

numbernumber

The settled value of the transaction, denominated in the transactions's currency as stated in `iso_currency_code` or `unofficial_currency_code`.
Positive values when money moves out of the account; negative values when money moves in.
For example, credit card purchases are positive; credit card payment, direct deposits, and refunds are negative.

[`date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-date)

stringstring

For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted.
Both dates are returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`name`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-name)

stringstring

The merchant name or transaction description.

[`original_description`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-original-description)

nullablestringnullable, string

The string returned by the financial institution to describe the transaction.

[`pending`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-pending)

booleanboolean

When true, identifies the transaction as pending or unsettled.
Pending transaction details (name, type, amount, category ID) may change before they are settled.

[`transaction_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-transaction-id)

stringstring

The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.

[`check_number`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-check-number)

nullablestringnullable, string

The check number of the transaction. This field is only populated for check transactions.

[`iso_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-transactions-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`last_updated_time`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-last-updated-time)

stringstring

The time when this Item's data was last retrieved from the financial institution.  
  

Format: `date-time`

[`institution_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-institution-id)

stringstring

The unique identifier of the institution associated with the Item.

[`institution_name`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-institution-name)

stringstring

The name of the institution associated with the Item.

[`item_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-item-id)

stringstring

The unique identifier for the Item.

[`bank_income_summary`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary)

objectobject

Summary for bank income across all income sources and items (max history of 730 days).

[`total_amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amount)

deprecatednumberdeprecated, number

Total amount of earnings across all the income sources in the end user's Items for the days requested by the client.
This may return an incorrect value if the summary includes income sources in multiple currencies.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts) instead.

[`iso_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-iso-currency-code)

deprecatednullablestringdeprecated, nullable, string

The ISO 4217 currency code of the amount or balance.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts) instead.

[`unofficial_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-unofficial-currency-code)

deprecatednullablestringdeprecated, nullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts) instead.

[`total_amounts`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts)

[object][object]

Total amount of earnings across all the income sources in the end user's Items for the days requested by the client.
This can contain multiple amounts, with each amount denominated in one unique currency.

[`amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`start_date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-start-date)

stringstring

The earliest date within the days requested in which all income sources identified by Plaid appear in a user's account.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-end-date)

stringstring

The latest date in which all income sources identified by Plaid appear in the user's account.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`income_sources_count`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-income-sources-count)

integerinteger

Number of income sources per end user.

[`income_categories_count`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-income-categories-count)

integerinteger

Number of income categories per end user.

[`income_transactions_count`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-income-transactions-count)

integerinteger

Number of income transactions per end user.

[`historical_summary`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary)

[object][object]

[`total_amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-total-amount)

deprecatednumberdeprecated, number

Total amount of earnings for the income source(s) of the user for the month in the summary.
This may return an incorrect value if the summary includes income sources in multiple currencies.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts) instead.

[`iso_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-iso-currency-code)

deprecatednullablestringdeprecated, nullable, string

The ISO 4217 currency code of the amount or balance.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts) instead.

[`unofficial_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-unofficial-currency-code)

deprecatednullablestringdeprecated, nullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.
Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-items-bank-income-sources-historical-summary-total-amounts) instead.

[`total_amounts`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-total-amounts)

[object][object]

Total amount of earnings for the income source(s) of the user for the month in the summary.
This can contain multiple amounts, with each amount denominated in one unique currency.

[`amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-total-amounts-amount)

numbernumber

Value of amount with up to 2 decimal places.

[`iso_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-total-amounts-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-total-amounts-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`start_date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-start-date)

stringstring

The start date of the period covered in this monthly summary.
This date will be the first day of the month, unless the month being covered is a partial month because it is the first month included in the summary and the date range being requested does not begin with the first day of the month.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-end-date)

stringstring

The end date of the period included in this monthly summary.
This date will be the last day of the month, unless the month being covered is a partial month because it is the last month included in the summary and the date range being requested does not end with the last day of the month.
The date will be returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`transactions`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions)

[object][object]

[`amount`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-amount)

numbernumber

The settled value of the transaction, denominated in the transactions's currency as stated in `iso_currency_code` or `unofficial_currency_code`.
Positive values when money moves out of the account; negative values when money moves in.
For example, credit card purchases are positive; credit card payment, direct deposits, and refunds are negative.

[`date`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-date)

stringstring

For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted.
Both dates are returned in an ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`name`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-name)

stringstring

The merchant name or transaction description.

[`original_description`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-original-description)

nullablestringnullable, string

The string returned by the financial institution to describe the transaction.

[`pending`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-pending)

booleanboolean

When true, identifies the transaction as pending or unsettled.
Pending transaction details (name, type, amount, category ID) may change before they are settled.

[`transaction_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-transaction-id)

stringstring

The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.

[`check_number`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-check-number)

nullablestringnullable, string

The check number of the transaction. This field is only populated for check transactions.

[`iso_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-iso-currency-code)

nullablestringnullable, string

The ISO 4217 currency code of the amount or balance.

[`unofficial_currency_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-historical-summary-transactions-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null.
Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.

[`warnings`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings)

[object][object]

If data from the report was unable to be retrieved, the warnings will contain information about the error that caused the data to be incomplete.

[`warning_type`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings-warning-type)

stringstring

The warning type which will always be `BANK_INCOME_WARNING`.  
  

Possible values: `BANK_INCOME_WARNING`

[`warning_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings-warning-code)

stringstring

The warning code identifies a specific kind of warning.
`IDENTITY_UNAVAILABLE`: Unable to extract identity for the Item
`TRANSACTIONS_UNAVAILABLE`: Unable to extract transactions for the Item
`ITEM_UNAPPROVED`: User exited flow before giving permission to share data for the Item
`REPORT_DELETED`: Report deleted due to customer or consumer request
`DATA_UNAVAILABLE`: No relevant data was found for the Item  
  

Possible values: `IDENTITY_UNAVAILABLE`, `TRANSACTIONS_UNAVAILABLE`, `ITEM_UNAPPROVED`, `REPORT_DELETED`, `DATA_UNAVAILABLE`

[`cause`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings-cause)

objectobject

An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.

[`error_type`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings-cause-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INTERNAL_SERVER_ERROR`, `INSUFFICIENT_CREDENTIALS`, `ITEM_LOCKED`, `USER_SETUP_REQUIRED`, `COUNTRY_NOT_SUPPORTED`, `INSTITUTION_DOWN`, `INSTITUTION_NO_LONGER_SUPPORTED`, `INSTITUTION_NOT_RESPONDING`, `INVALID_CREDENTIALS`, `INVALID_MFA`, `INVALID_SEND_METHOD`, `ITEM_LOGIN_REQUIRED`, `MFA_NOT_SUPPORTED`, `NO_ACCOUNTS`, `ITEM_NOT_SUPPORTED`, `ACCESS_NOT_GRANTED`

[`error_code`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings-cause-error-code)

stringstring

We use standard HTTP response codes for success and failure notifications, and our errors are further classified by `error_type`. In general, 200 HTTP codes correspond to success, 40X codes are for developer- or user-related failures, and 50X codes are for Plaid-related issues. Error fields will be `null` if no error has occurred.

[`error_message`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings-cause-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings-cause-display-message)

stringstring

A user-friendly representation of the error code. null if the error is not related to user action.
This may change over time and is not safe for programmatic use.

[`item_id`](/docs/api/products/income/#credit-bank_income-get-response-bank-income-warnings-cause-item-id)

stringstring

The `item_id` of the Item associated with this warning.

[`request_id`](/docs/api/products/income/#credit-bank_income-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "bank_income": [
    {
      "bank_income_id": "dacc92a0-cb59-43a5-ba24-1b1c07a03f28",
      "bank_income_summary": {
        "end_date": "2024-08-21",
        "historical_summary": [
          {
            "end_date": "2024-08-21",
            "iso_currency_code": "USD",
            "start_date": "2024-08-06",
            "total_amount": 4090.14,
            "total_amounts": [
              {
                "amount": 4090.14,
                "iso_currency_code": "USD",
                "unofficial_currency_code": null
              }
            ],
            "transactions": [
              {
                "amount": 120.12,
                "check_number": null,
                "date": "2024-08-07",
                "iso_currency_code": "USD",
                "name": "TEXAS OAG CHILD SUPPORT",
                "original_description": "TEXAS OAG CHILD SUPPORT",
                "transaction_id": "EZMmvwREqlSGmlRam7bzFKyBll3kJjU4xKm1w",
                "unofficial_currency_code": null
              },
              {
                "amount": 1525,
                "check_number": null,
                "date": "2024-08-08",
                "iso_currency_code": "USD",
                "name": "AIRBNB PAYMENTS PPD ID: 1234567890",
                "original_description": "AIRBNB PAYMENTS PPD ID: 1234567890",
                "transaction_id": "Wr6jzLwg1qs6ag9Xa8BrCpBAPPxnEXF6ZmjDR",
                "unofficial_currency_code": null
              },
              {
                "amount": 500,
                "check_number": null,
                "date": "2024-08-12",
                "iso_currency_code": "USD",
                "name": "TWC-BENEFITS/UI BENEFIT",
                "original_description": "TWC-BENEFITS/UI BENEFIT",
                "transaction_id": "Aj7Apx5bDyIA3VRl35yqC18wXXorBgI9rX5dp",
                "unofficial_currency_code": null
              },
              {
                "amount": 1000.7,
                "check_number": null,
                "date": "2024-08-15",
                "iso_currency_code": "USD",
                "name": "PLAID PAYROLL",
                "original_description": "PLAID PAYROLL",
                "transaction_id": "G1L9oybBrKSMPmBdPzXoFN8aGGE7gXC6MeoQB",
                "unofficial_currency_code": null
              },
              {
                "amount": 824.2,
                "check_number": null,
                "date": "2024-08-15",
                "iso_currency_code": "USD",
                "name": "SSI TREAS 310 XXSUPP SEC PPD ID: 1234567890",
                "original_description": "SSI TREAS 310 XXSUPP SEC PPD ID: 1234567890",
                "transaction_id": "nWLlwMm1qxi8DomvDXP3FaGjXX5bm9TAlyQnk",
                "unofficial_currency_code": null
              },
              {
                "amount": 120.12,
                "check_number": null,
                "date": "2024-08-21",
                "iso_currency_code": "USD",
                "name": "TEXAS OAG CHILD SUPPORT",
                "original_description": "TEXAS OAG CHILD SUPPORT",
                "transaction_id": "b7dkg6eQbPFQeRvVeZlxcqxZooa7nWSmb47dj",
                "unofficial_currency_code": null
              }
            ],
            "unofficial_currency_code": null
          }
        ],
        "income_categories_count": 5,
        "income_sources_count": 5,
        "income_transactions_count": 6,
        "iso_currency_code": "USD",
        "start_date": "2024-08-07",
        "total_amount": 4090.14,
        "total_amounts": [
          {
            "amount": 4090.14,
            "iso_currency_code": "USD",
            "unofficial_currency_code": null
          }
        ],
        "unofficial_currency_code": null
      },
      "days_requested": 15,
      "generated_time": "2024-08-21T18:10:46.293199Z",
      "items": [
        {
          "bank_income_accounts": [
            {
              "account_id": "G1L9oybBrKSMPmBdPzXoFN8oo16rqqC6PwkA5",
              "mask": "9217",
              "name": "Checking",
              "official_name": "Plaid checking",
              "owners": [
                {
                  "addresses": [],
                  "emails": [],
                  "names": [
                    "Jane Doe"
                  ],
                  "phone_numbers": []
                }
              ],
              "subtype": "checking",
              "type": "depository"
            }
          ],
          "bank_income_sources": [
            {
              "account_id": "G1L9oybBrKSMPmBdPzXoFN8oo16rqqC6PwkA5",
              "end_date": "2024-08-15",
              "historical_summary": [
                {
                  "end_date": "2024-08-21",
                  "iso_currency_code": "USD",
                  "start_date": "2024-08-06",
                  "total_amount": 1000.7,
                  "total_amounts": [
                    {
                      "amount": 1000.7,
                      "iso_currency_code": "USD",
                      "unofficial_currency_code": null
                    }
                  ],
                  "transactions": [
                    {
                      "amount": 1000.7,
                      "check_number": null,
                      "date": "2024-08-15",
                      "iso_currency_code": "USD",
                      "name": "PLAID PAYROLL",
                      "original_description": "PLAID PAYROLL",
                      "transaction_id": "G1L9oybBrKSMPmBdPzXoFN8aGGE7gXC6MeoQB",
                      "unofficial_currency_code": null
                    }
                  ],
                  "unofficial_currency_code": null
                }
              ],
              "income_category": "SALARY",
              "income_description": "PLAID PAYROLL",
              "income_source_id": "0e9d6fbc-29de-4225-9843-2f71e02a54d1",
              "pay_frequency": "UNKNOWN",
              "start_date": "2024-08-15",
              "total_amount": 1000.7,
              "transaction_count": 1
            },
            {
              "account_id": "G1L9oybBrKSMPmBdPzXoFN8oo16rqqC6PwkA5",
              "end_date": "2024-08-15",
              "historical_summary": [
                {
                  "end_date": "2024-08-21",
                  "iso_currency_code": "USD",
                  "start_date": "2024-08-06",
                  "total_amount": 824.2,
                  "total_amounts": [
                    {
                      "amount": 824.2,
                      "iso_currency_code": "USD",
                      "unofficial_currency_code": null
                    }
                  ],
                  "transactions": [
                    {
                      "amount": 824.2,
                      "check_number": null,
                      "date": "2024-08-15",
                      "iso_currency_code": "USD",
                      "name": "SSI TREAS 310 XXSUPP SEC PPD ID: 1234567890",
                      "original_description": "SSI TREAS 310 XXSUPP SEC PPD ID: 1234567890",
                      "transaction_id": "nWLlwMm1qxi8DomvDXP3FaGjXX5bm9TAlyQnk",
                      "unofficial_currency_code": null
                    }
                  ],
                  "unofficial_currency_code": null
                }
              ],
              "income_category": "LONG_TERM_DISABILITY",
              "income_description": "SSI TREAS 310 XXSUPP SEC PPD ID: 1234567890",
              "income_source_id": "88bc00d8-2bb1-42d0-a054-db3f20948283",
              "pay_frequency": "UNKNOWN",
              "start_date": "2024-08-15",
              "total_amount": 824.2,
              "transaction_count": 1
            },
            {
              "account_id": "G1L9oybBrKSMPmBdPzXoFN8oo16rqqC6PwkA5",
              "end_date": "2024-08-08",
              "historical_summary": [
                {
                  "end_date": "2024-08-21",
                  "iso_currency_code": "USD",
                  "start_date": "2024-08-06",
                  "total_amount": 1525,
                  "total_amounts": [
                    {
                      "amount": 1525,
                      "iso_currency_code": "USD",
                      "unofficial_currency_code": null
                    }
                  ],
                  "transactions": [
                    {
                      "amount": 1525,
                      "check_number": null,
                      "date": "2024-08-08",
                      "iso_currency_code": "USD",
                      "name": "AIRBNB PAYMENTS PPD ID: 1234567890",
                      "original_description": "AIRBNB PAYMENTS PPD ID: 1234567890",
                      "transaction_id": "Wr6jzLwg1qs6ag9Xa8BrCpBAPPxnEXF6ZmjDR",
                      "unofficial_currency_code": null
                    }
                  ],
                  "unofficial_currency_code": null
                }
              ],
              "income_category": "RENTAL",
              "income_description": "AIRBNB PAYMENTS PPD ID: 1234567890",
              "income_source_id": "063689af-7299-4327-b71f-9d8849a40c0e",
              "pay_frequency": "UNKNOWN",
              "start_date": "2024-08-08",
              "total_amount": 1525,
              "transaction_count": 1
            },
            {
              "account_id": "G1L9oybBrKSMPmBdPzXoFN8oo16rqqC6PwkA5",
              "end_date": "2024-08-12",
              "historical_summary": [
                {
                  "end_date": "2024-08-21",
                  "iso_currency_code": "USD",
                  "start_date": "2024-08-06",
                  "total_amount": 500,
                  "total_amounts": [
                    {
                      "amount": 500,
                      "iso_currency_code": "USD",
                      "unofficial_currency_code": null
                    }
                  ],
                  "transactions": [
                    {
                      "amount": 500,
                      "check_number": null,
                      "date": "2024-08-12",
                      "iso_currency_code": "USD",
                      "name": "TWC-BENEFITS/UI BENEFIT",
                      "original_description": "TWC-BENEFITS/UI BENEFIT",
                      "transaction_id": "Aj7Apx5bDyIA3VRl35yqC18wXXorBgI9rX5dp",
                      "unofficial_currency_code": null
                    }
                  ],
                  "unofficial_currency_code": null
                }
              ],
              "income_category": "UNEMPLOYMENT",
              "income_description": "TWC-BENEFITS/UI BENEFIT",
              "income_source_id": "ce160170-49d0-4811-b58e-cb4878d05f83",
              "pay_frequency": "UNKNOWN",
              "start_date": "2024-08-12",
              "total_amount": 500,
              "transaction_count": 1
            },
            {
              "account_id": "G1L9oybBrKSMPmBdPzXoFN8oo16rqqC6PwkA5",
              "end_date": "2024-08-21",
              "historical_summary": [
                {
                  "end_date": "2024-08-21",
                  "iso_currency_code": "USD",
                  "start_date": "2024-08-06",
                  "total_amount": 240.24,
                  "total_amounts": [
                    {
                      "amount": 240.24,
                      "iso_currency_code": "USD",
                      "unofficial_currency_code": null
                    }
                  ],
                  "transactions": [
                    {
                      "amount": 120.12,
                      "check_number": null,
                      "date": "2024-08-07",
                      "iso_currency_code": "USD",
                      "name": "TEXAS OAG CHILD SUPPORT",
                      "original_description": "TEXAS OAG CHILD SUPPORT",
                      "transaction_id": "EZMmvwREqlSGmlRam7bzFKyBll3kJjU4xKm1w",
                      "unofficial_currency_code": null
                    },
                    {
                      "amount": 120.12,
                      "check_number": null,
                      "date": "2024-08-21",
                      "iso_currency_code": "USD",
                      "name": "TEXAS OAG CHILD SUPPORT",
                      "original_description": "TEXAS OAG CHILD SUPPORT",
                      "transaction_id": "b7dkg6eQbPFQeRvVeZlxcqxZooa7nWSmb47dj",
                      "unofficial_currency_code": null
                    }
                  ],
                  "unofficial_currency_code": null
                }
              ],
              "income_category": "CHILD_SUPPORT",
              "income_description": "TEXAS OAG CHILD SUPPORT",
              "income_source_id": "c8e1576e-9de4-47b4-ad55-3f7b068cc863",
              "pay_frequency": "UNKNOWN",
              "start_date": "2024-08-07",
              "total_amount": 240.24,
              "transaction_count": 2
            }
          ],
          "institution_id": "ins_20",
          "institution_name": "Citizens Bank",
          "item_id": "L8EKo4GydxSKmJQGmXyPuDkeNn4rg9fP3MKLv",
          "last_updated_time": "2024-08-21T18:10:47.367335Z"
        }
      ]
    }
  ],
  "request_id": "MLM1fFu4fbVg7KR"
}
```

=\*=\*=\*=

#### `/credit/bank_income/pdf/get`

#### Retrieve information from the bank accounts used for income verification in PDF format

[`/credit/bank_income/pdf/get`](/docs/api/products/income/#creditbank_incomepdfget) returns the most recent bank income report for a specified user in PDF format. A single report corresponds to all institutions linked in a single Link session. To include multiple institutions in a single report, use [Multi-Item Link](https://plaid.com/docs/link/multi-item-link).

/credit/bank\_income/pdf/get

**Request fields**

[`client_id`](/docs/api/products/income/#credit-bank_income-pdf-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-bank_income-pdf-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-bank_income-pdf-get-request-user-token)

requiredstringrequired, string

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/credit/bank\_income/pdf/get

```
const request: CreditBankIncomePDFGetRequest = {
  user_token: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
};

try {
  const response = await client.creditBankIncomePdfGet(request, {
    responseType: 'arraybuffer',
  });
  const pdf = response.buffer.toString('base64');
} catch (error) {
  // handle error
}
```

##### Response

This endpoint returns binary PDF data. [View a sample Bank Income PDF.](https://plaid.com/documents/sample-bank-income.pdf)

=\*=\*=\*=

#### `/credit/bank_income/refresh`

#### Refresh a user's bank income information

[`/credit/bank_income/refresh`](/docs/api/products/income/#creditbank_incomerefresh) refreshes the most recent bank income report data for a specific user. If the most recent bank income report is no longer valid (i.e. deleted), the endpoint will refresh the most recent valid report instead.

/credit/bank\_income/refresh

**Request fields**

[`client_id`](/docs/api/products/income/#credit-bank_income-refresh-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-bank_income-refresh-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-bank_income-refresh-request-user-token)

requiredstringrequired, string

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`options`](/docs/api/products/income/#credit-bank_income-refresh-request-options)

objectobject

An optional object for `/credit/bank_income/refresh` request options.

[`days_requested`](/docs/api/products/income/#credit-bank_income-refresh-request-options-days-requested)

integerinteger

How many days of data to include in the refresh. If not specified, this will default to the days requested in the most recently generated bank income report for the user.

/credit/bank\_income/refresh

```
const request: CreditBankIncomeRefreshRequest = {
  user_token: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
  options: {
    days_requested: 120,
  },
};

try {
  const response = await client.creditBankIncomeRefresh(request);
} catch (error) {
  // handle error
}
```

/credit/bank\_income/refresh

**Response fields**

[`request_id`](/docs/api/products/income/#credit-bank_income-refresh-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm"
}
```

=\*=\*=\*=

#### `/credit/bank_statements/uploads/get`

#### Retrieve data for a user's uploaded bank statements

[`/credit/bank_statements/uploads/get`](/docs/api/products/income/#creditbank_statementsuploadsget) returns parsed data from bank statements uploaded by users as part of the Document Income flow. If your account is not enabled for Document Parsing, contact your account manager to request access.

/credit/bank\_statements/uploads/get

**Request fields**

[`client_id`](/docs/api/products/income/#credit-bank_statements-uploads-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-bank_statements-uploads-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-bank_statements-uploads-get-request-user-token)

requiredstringrequired, string

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`options`](/docs/api/products/income/#credit-bank_statements-uploads-get-request-options)

objectobject

An optional object for `/credit/bank_statements/uploads/get` request options.

[`item_ids`](/docs/api/products/income/#credit-bank_statements-uploads-get-request-options-item-ids)

[string][string]

An array of `item_id`s whose bank statements information is returned. Each `item_id` should uniquely identify a bank statements uploaded item. If this field is not provided, all `item_id`s associated with the `user_token` will returned in the response.

/credit/bank\_statements/uploads/get

```
const request: CreditBankStatementsUploadsGetRequest = {
  user_token: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
};

try {
  const response = await client.creditBankStatementsUploadsGet(request);
} catch (error) {
  // handle error
}
```

/credit/bank\_statements/uploads/get

**Response fields**

[`items`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items)

[object][object]

Array of bank statement upload items.

[`item_id`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`bank_statements`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements)

[object][object]

[`transactions`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-transactions)

[object][object]

An array of transactions appearing on the bank statement.

[`amount`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-transactions-amount)

nullablenumbernullable, number

The value of the transaction. A negative amount indicates that money moved into the account (such as a paycheck being deposited).

[`date`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-transactions-date)

nullablestringnullable, string

The date of when the transaction was made, in ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`original_description`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-transactions-original-description)

nullablestringnullable, string

The raw description of the transaction as it appears on the bank statement.

[`account_id`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-transactions-account-id)

nullablestringnullable, string

The unique id of the bank account that this transaction occurs in

[`document_metadata`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-document-metadata)

objectobject

Object representing metadata pertaining to the document.

[`name`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-document-metadata-name)

stringstring

The name of the document.

[`document_type`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-document-metadata-document-type)

nullablestringnullable, string

The type of document.  
`PAYSTUB`: A paystub.  
`BANK_STATEMENT`: A bank statement.  
`US_TAX_W2`: A W-2 wage and tax statement provided by a US employer reflecting wages earned by the employee.  
`US_MILITARY_ERAS`: An electronic Retirement Account Statement (eRAS) issued by the US military.  
`US_MILITARY_LES`: A Leave and Earnings Statement (LES) issued by the US military.  
`US_MILITARY_CLES`: A Civilian Leave and Earnings Statement (CLES) issued by the US military.  
`GIG`: Used to indicate that the income is related to gig work. Does not necessarily correspond to a specific document type.  
`PLAID_GENERATED_PAYSTUB_PDF`: Used to indicate that the PDF for the paystub was generated by Plaid.  
`NONE`: Used to indicate that there is no underlying document for the data.  
`UNKNOWN`: Document type could not be determined.  
  

Possible values: `UNKNOWN`, `PAYSTUB`, `BANK_STATEMENT`, `US_TAX_W2`, `US_MILITARY_ERAS`, `US_MILITARY_LES`, `US_MILITARY_CLES`, `GIG`, `PLAID_GENERATED_PAYSTUB_PDF`, `NONE`

[`download_url`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-document-metadata-download-url)

nullablestringnullable, string

Signed URL to retrieve the document(s). The payload will be a .zip file containing the document(s).  
For Payroll Income, the file type of the documents will always be PDF, and the documents may not be available, in which case the field will be `null`. If you would like Plaid to generate a PDF if the original is not available, contact your Account Manager. [Example generated pay stub](https://plaid.com/documents/plaid-generated-mock-paystub.pdf).  
For Document Income, this field will not be `null`, and the file type of the underlying document(s) will be the original file type uploaded by the user. For more details on available file types, see the [Document Income](https://plaid.com/docs/income/payroll-income) documentation.  
This download URL can only be used once and expires after two minutes. To generate a new download URL, call `/credit/payroll_income/get` again.

[`status`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-document-metadata-status)

nullablestringnullable, string

The processing status of the document.  
`PROCESSING_COMPLETE`: The document was successfully processed.  
`DOCUMENT_ERROR`: The document could not be processed. Possible causes include: The document was an unacceptable document type such as an offer letter or bank statement, the document image was cropped or blurry, or the document was corrupted.  
`UNKNOWN` or `null`: An internal error occurred. If this happens repeatedly, contact support or your Plaid account manager.  
  

Possible values: `UNKNOWN`, `PROCESSING_COMPLETE`, `DOCUMENT_ERROR`, `null`

[`page_count`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-document-metadata-page-count)

nullableintegernullable, integer

The number of pages of the uploaded document (if available).

[`error_message`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-document-metadata-error-message)

nullablestringnullable, string

The reason why a failure occurred during document processing (if available).

[`document_id`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-document-id)

nullablestringnullable, string

An identifier of the document referenced by the document metadata.

[`bank_accounts`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts)

[object][object]

An array of bank accounts associated with the uploaded bank statement.

[`name`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-name)

nullablestringnullable, string

The name of the bank account

[`bank_name`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-bank-name)

nullablestringnullable, string

The name of the bank institution.

[`account_type`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-account-type)

nullablestringnullable, string

The type of the bank account.

[`account_number`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-account-number)

nullablestringnullable, string

The bank account number.

[`owner`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-owner)

objectobject

An object containing data about the owner of the bank account for the uploaded bank statement.

[`name`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-owner-name)

nullablestringnullable, string

The name of the account owner

[`address`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-owner-address)

objectobject

Address on the uploaded bank statement

[`city`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-owner-address-city)

nullablestringnullable, string

The full city name.

[`country`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-owner-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-owner-address-postal-code)

nullablestringnullable, string

The postal code of the address.

[`region`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-owner-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-owner-address-street)

nullablestringnullable, string

The full street address.

[`periods`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-periods)

[object][object]

An array of period objects, containing more data on the overall period of the statement.

[`start_date`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-periods-start-date)

nullablestringnullable, string

The start date of the statement period in ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-periods-end-date)

nullablestringnullable, string

The end date of the statement period in ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`starting_balance`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-periods-starting-balance)

nullablenumbernullable, number

The starting balance of the bank account for the period.

[`ending_balance`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-periods-ending-balance)

nullablenumbernullable, number

The ending balance of the bank account for the period.

[`account_id`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-bank-statements-bank-accounts-account-id)

nullablestringnullable, string

The unique id of the bank account

[`status`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-status)

nullableobjectnullable, object

Details about the status of the payroll item.

[`processing_status`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-status-processing-status)

nullablestringnullable, string

Denotes the processing status for the verification.  
`UNKNOWN`: The processing status could not be determined.  
`PROCESSING_COMPLETE`: The processing has completed and the user has approved for sharing. The data is available to be retrieved.  
`PROCESSING`: The verification is still processing. The data is not available yet.  
`FAILED`: The processing failed to complete successfully.  
`APPROVAL_STATUS_PENDING`: The processing has completed but the user has not yet approved the sharing of the data.  
  

Possible values: `UNKNOWN`, `PROCESSING_COMPLETE`, `PROCESSING`, `FAILED`, `APPROVAL_STATUS_PENDING`

[`updated_at`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-items-updated-at)

nullablestringnullable, string

Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DDTHH:mm:ssZ) indicating the last time that the Item was updated.  
  

Format: `date-time`

[`request_id`](/docs/api/products/income/#credit-bank_statements-uploads-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "items": [
    {
      "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
      "bank_statements": [
        {
          "transactions": [
            {
              "amount": -1000,
              "date": "2023-01-01",
              "original_description": "PAYCHECK",
              "account_id": "c6778d3f-e44c-4348-874e-71507c1ac12d"
            }
          ],
          "document_metadata": {
            "document_type": "BANK_STATEMENT",
            "name": "statement_01.pdf",
            "status": "PROCESSING_COMPLETE",
            "download_url": null,
            "page_count": 2
          },
          "document_id": "2jkflanbd",
          "bank_accounts": [
            {
              "name": "CHASE CHECKING",
              "bank_name": "CHASE",
              "account_type": "CHECKING",
              "account_number": "000009752",
              "account_id": "c6778d3f-e44c-4348-874e-71507c1ac12d",
              "owner": {
                "name": "JANE DOE",
                "address": {
                  "postal_code": "94133",
                  "country": "US",
                  "region": "CA",
                  "city": "SAN FRANCISCO",
                  "street": "2140 TAYLOR ST"
                }
              },
              "periods": [
                {
                  "start_date": "2023-01-01",
                  "end_date": "2023-02-01",
                  "starting_balance": 2500,
                  "ending_balance": 3500
                }
              ]
            }
          ]
        }
      ],
      "status": {
        "processing_status": "PROCESSING_COMPLETE"
      },
      "updated_at": "2023-02-01T21:14:54Z"
    }
  ],
  "request_id": "LhQf0THi8SH1yJm"
}
```

=\*=\*=\*=

#### `/credit/payroll_income/get`

#### Retrieve a user's payroll information

This endpoint gets payroll income information for a specific user, either as a result of the user connecting to their payroll provider or uploading a pay related document.

/credit/payroll\_income/get

**Request fields**

[`client_id`](/docs/api/products/income/#credit-payroll_income-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-payroll_income-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-payroll_income-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`options`](/docs/api/products/income/#credit-payroll_income-get-request-options)

objectobject

An optional object for `/credit/payroll_income/get` request options.

[`item_ids`](/docs/api/products/income/#credit-payroll_income-get-request-options-item-ids)

[string][string]

An array of `item_id`s whose payroll information is returned. Each `item_id` should uniquely identify a payroll income item. If this field is not provided, all `item_id`s associated with the `user_token` will returned in the response.

/credit/payroll\_income/get

```
const request: CreditPayrollIncomeGetRequest = {
  user_token: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
};

try {
  const response = await client.creditPayrollIncomeGet(request);
} catch (error) {
  // handle error
}
```

/credit/payroll\_income/get

**Response fields**

[`items`](/docs/api/products/income/#credit-payroll_income-get-response-items)

[object][object]

Array of payroll items.

[`item_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`institution_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-institution-id)

stringstring

The unique identifier of the institution associated with the Item.

[`institution_name`](/docs/api/products/income/#credit-payroll_income-get-response-items-institution-name)

stringstring

The name of the institution associated with the Item.

[`accounts`](/docs/api/products/income/#credit-payroll_income-get-response-items-accounts)

[object][object]

[`account_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-accounts-account-id)

nullablestringnullable, string

ID of the payroll provider account.

[`rate_of_pay`](/docs/api/products/income/#credit-payroll_income-get-response-items-accounts-rate-of-pay)

objectobject

An object representing the rate at which an individual is paid.

[`pay_rate`](/docs/api/products/income/#credit-payroll_income-get-response-items-accounts-rate-of-pay-pay-rate)

nullablestringnullable, string

The rate at which an employee is paid.  
  

Possible values: `ANNUAL`, `HOURLY`, `CONTRACT`, `WEEKLY`, `BIWEEKLY`, `MONTHLY`, `SEMI_MONTHLY`, `DAILY`, `COMMISSION`, `OTHER`, `null`

[`pay_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-accounts-rate-of-pay-pay-amount)

nullablenumbernullable, number

The amount at which an employee is paid.  
  

Format: `double`

[`pay_frequency`](/docs/api/products/income/#credit-payroll_income-get-response-items-accounts-pay-frequency)

nullablestringnullable, string

The frequency at which an individual is paid.  
  

Possible values: `DAILY`, `WEEKLY`, `BIWEEKLY`, `SEMI_MONTHLY`, `MONTHLY`, `CONTRACT`, `QUARTERLY`, `SEMI_ANNUALLY`, `ANNUALLY`, `OTHER`, `null`

[`payroll_income`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income)

[object][object]

[`account_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-account-id)

nullablestringnullable, string

ID of the payroll provider account.

[`pay_stubs`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs)

[object][object]

Array of pay stubs for the user.

[`deductions`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions)

objectobject

An object with the deduction information found on a pay stub.

[`breakdown`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-breakdown)

[object][object]

[`current_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-breakdown-current-amount)

nullablenumbernullable, number

Raw amount of the deduction  
  

Format: `double`

[`description`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-breakdown-description)

nullablestringnullable, string

Description of the deduction line item

[`iso_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-breakdown-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-breakdown-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the line item. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`ytd_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-breakdown-ytd-amount)

nullablenumbernullable, number

The year-to-date amount of the deduction  
  

Format: `double`

[`total`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-total)

objectobject

An object representing the total deductions for the pay period

[`current_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-total-current-amount)

nullablenumbernullable, number

Raw amount of the deduction  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-total-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-total-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the line item. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`ytd_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-deductions-total-ytd-amount)

nullablenumbernullable, number

The year-to-date total amount of the deductions  
  

Format: `double`

[`document_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-document-id)

nullablestringnullable, string

An identifier of the document referenced by the document metadata.

[`document_metadata`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-document-metadata)

objectobject

Object representing metadata pertaining to the document.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-document-metadata-name)

stringstring

The name of the document.

[`document_type`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-document-metadata-document-type)

nullablestringnullable, string

The type of document.  
`PAYSTUB`: A paystub.  
`BANK_STATEMENT`: A bank statement.  
`US_TAX_W2`: A W-2 wage and tax statement provided by a US employer reflecting wages earned by the employee.  
`US_MILITARY_ERAS`: An electronic Retirement Account Statement (eRAS) issued by the US military.  
`US_MILITARY_LES`: A Leave and Earnings Statement (LES) issued by the US military.  
`US_MILITARY_CLES`: A Civilian Leave and Earnings Statement (CLES) issued by the US military.  
`GIG`: Used to indicate that the income is related to gig work. Does not necessarily correspond to a specific document type.  
`PLAID_GENERATED_PAYSTUB_PDF`: Used to indicate that the PDF for the paystub was generated by Plaid.  
`NONE`: Used to indicate that there is no underlying document for the data.  
`UNKNOWN`: Document type could not be determined.  
  

Possible values: `UNKNOWN`, `PAYSTUB`, `BANK_STATEMENT`, `US_TAX_W2`, `US_MILITARY_ERAS`, `US_MILITARY_LES`, `US_MILITARY_CLES`, `GIG`, `PLAID_GENERATED_PAYSTUB_PDF`, `NONE`

[`download_url`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-document-metadata-download-url)

nullablestringnullable, string

Signed URL to retrieve the document(s). The payload will be a .zip file containing the document(s).  
For Payroll Income, the file type of the documents will always be PDF, and the documents may not be available, in which case the field will be `null`. If you would like Plaid to generate a PDF if the original is not available, contact your Account Manager. [Example generated pay stub](https://plaid.com/documents/plaid-generated-mock-paystub.pdf).  
For Document Income, this field will not be `null`, and the file type of the underlying document(s) will be the original file type uploaded by the user. For more details on available file types, see the [Document Income](https://plaid.com/docs/income/payroll-income) documentation.  
This download URL can only be used once and expires after two minutes. To generate a new download URL, call `/credit/payroll_income/get` again.

[`status`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-document-metadata-status)

nullablestringnullable, string

The processing status of the document.  
`PROCESSING_COMPLETE`: The document was successfully processed.  
`DOCUMENT_ERROR`: The document could not be processed. Possible causes include: The document was an unacceptable document type such as an offer letter or bank statement, the document image was cropped or blurry, or the document was corrupted.  
`UNKNOWN` or `null`: An internal error occurred. If this happens repeatedly, contact support or your Plaid account manager.  
  

Possible values: `UNKNOWN`, `PROCESSING_COMPLETE`, `DOCUMENT_ERROR`, `null`

[`page_count`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-document-metadata-page-count)

nullableintegernullable, integer

The number of pages of the uploaded document (if available).

[`error_message`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-document-metadata-error-message)

nullablestringnullable, string

The reason why a failure occurred during document processing (if available).

[`earnings`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings)

objectobject

An object representing both a breakdown of earnings on a pay stub and the total earnings.

[`breakdown`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown)

[object][object]

[`canonical_description`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown-canonical-description)

nullablestringnullable, string

Commonly used term to describe the earning line item.  
  

Possible values: `BONUS`, `COMMISSION`, `OVERTIME`, `PAID_TIME_OFF`, `REGULAR_PAY`, `VACATION`, `BASIC_ALLOWANCE_HOUSING`, `BASIC_ALLOWANCE_SUBSISTENCE`, `OTHER`, `ALLOWANCE`, `BEREAVEMENT`, `HOLIDAY_PAY`, `JURY_DUTY`, `LEAVE`, `LONG_TERM_DISABILITY_PAY`, `MILITARY_PAY`, `PER_DIEM`, `REFERRAL_BONUS`, `REIMBURSEMENTS`, `RETENTION_BONUS`, `RETROACTIVE_PAY`, `SEVERANCE_PAY`, `SHIFT_DIFFERENTIAL`, `SHORT_TERM_DISABILITY_PAY`, `SICK_PAY`, `SIGNING_BONUS`, `TIPS_INCOME`, `RETIREMENT`, `GIG_ECONOMY`, `STOCK_COMPENSATION`, `null`

[`current_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown-current-amount)

nullablenumbernullable, number

Raw amount of the earning line item.  
  

Format: `double`

[`description`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown-description)

nullablestringnullable, string

Description of the earning line item.

[`hours`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown-hours)

nullablenumbernullable, number

Number of hours applicable for this earning.

[`iso_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.

[`rate`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown-rate)

nullablenumbernullable, number

Hourly rate applicable for this earning.  
  

Format: `double`

[`unofficial_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the line item. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`ytd_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-breakdown-ytd-amount)

nullablenumbernullable, number

The year-to-date amount of the deduction.  
  

Format: `double`

[`total`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-total)

objectobject

An object representing both the current pay period and year to date amount for an earning category.

[`current_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-total-current-amount)

nullablenumbernullable, number

Total amount of the earnings for this pay period.  
  

Format: `double`

[`hours`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-total-hours)

nullablenumbernullable, number

Total number of hours worked for this pay period.

[`iso_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-total-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-total-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the security. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`ytd_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-earnings-total-ytd-amount)

nullablenumbernullable, number

The total year-to-date amount of the earnings.  
  

Format: `double`

[`employee`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee)

objectobject

Data about the employee.

[`address`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-address)

objectobject

Address on the pay stub.

[`city`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-address-city)

nullablestringnullable, string

The full city name.

[`country`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-address-postal-code)

nullablestringnullable, string

The postal code of the address.

[`region`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-address-street)

nullablestringnullable, string

The full street address.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-name)

nullablestringnullable, string

The name of the employee.

[`marital_status`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-marital-status)

nullablestringnullable, string

Marital status of the employee - either `SINGLE` or `MARRIED` or `NOT LISTED`.  
  

Possible values: `SINGLE`, `MARRIED`, `NOT LISTED`, `null`

[`taxpayer_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-taxpayer-id)

objectobject

Taxpayer ID of the individual receiving the paystub.

[`id_type`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-taxpayer-id-id-type)

nullablestringnullable, string

Type of ID, e.g. 'SSN'.

[`id_mask`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employee-taxpayer-id-id-mask)

nullablestringnullable, string

ID mask; i.e. last 4 digits of the taxpayer ID.

[`employer`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employer)

objectobject

Information about the employer on the pay stub.

[`address`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employer-address)

objectobject

Address on the pay stub.

[`city`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employer-address-city)

nullablestringnullable, string

The full city name.

[`country`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employer-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employer-address-postal-code)

nullablestringnullable, string

The postal code of the address.

[`region`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employer-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employer-address-street)

nullablestringnullable, string

The full street address.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-employer-name)

nullablestringnullable, string

The name of the employer on the pay stub.

[`net_pay`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-net-pay)

objectobject

An object representing information about the net pay amount on the pay stub.

[`current_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-net-pay-current-amount)

nullablenumbernullable, number

Raw amount of the net pay for the pay period.  
  

Format: `double`

[`description`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-net-pay-description)

nullablestringnullable, string

Description of the net pay.

[`iso_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-net-pay-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.

[`unofficial_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-net-pay-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the net pay. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`ytd_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-net-pay-ytd-amount)

nullablenumbernullable, number

The year-to-date amount of the net pay.  
  

Format: `double`

[`pay_period_details`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details)

objectobject

Details about the pay period.

[`pay_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-pay-amount)

nullablenumbernullable, number

The amount of the paycheck.  
  

Format: `double`

[`distribution_breakdown`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-distribution-breakdown)

[object][object]

[`account_name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-distribution-breakdown-account-name)

nullablestringnullable, string

Name of the account for the given distribution.

[`bank_name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-distribution-breakdown-bank-name)

nullablestringnullable, string

The name of the bank that the payment is being deposited to.

[`current_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-distribution-breakdown-current-amount)

nullablenumbernullable, number

The amount distributed to this account.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-distribution-breakdown-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.

[`mask`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-distribution-breakdown-mask)

nullablestringnullable, string

The last 2-4 alphanumeric characters of an account's official account number.

[`type`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-distribution-breakdown-type)

nullablestringnullable, string

Type of the account that the paystub was sent to (e.g. 'checking').

[`unofficial_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-distribution-breakdown-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the net pay. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`end_date`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-end-date)

nullablestringnullable, string

The date on which the pay period ended, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`gross_earnings`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-gross-earnings)

nullablenumbernullable, number

Total earnings before tax/deductions.  
  

Format: `double`

[`iso_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-iso-currency-code)

nullablestringnullable, string

The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.

[`pay_date`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-pay-date)

nullablestringnullable, string

The date on which the pay stub was issued, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`pay_frequency`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-pay-frequency)

nullablestringnullable, string

The frequency at which an individual is paid.  
  

Possible values: `UNKNOWN`, `WEEKLY`, `BIWEEKLY`, `SEMI_MONTHLY`, `MONTHLY`, `null`

[`pay_basis`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-pay-basis)

stringstring

The explicit pay basis on the paystub (if present).  
  

Possible values: `SALARY`, `HOURLY`, `COMMISSION`

[`start_date`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-start-date)

nullablestringnullable, string

The date on which the pay period started, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ("yyyy-mm-dd").  
  

Format: `date`

[`unofficial_currency_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-pay-stubs-pay-period-details-unofficial-currency-code)

nullablestringnullable, string

The unofficial currency code associated with the net pay. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  
See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.

[`w2s`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s)

[object][object]

Array of tax form W-2s.

[`document_metadata`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-document-metadata)

objectobject

Object representing metadata pertaining to the document.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-document-metadata-name)

stringstring

The name of the document.

[`document_type`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-document-metadata-document-type)

nullablestringnullable, string

The type of document.  
`PAYSTUB`: A paystub.  
`BANK_STATEMENT`: A bank statement.  
`US_TAX_W2`: A W-2 wage and tax statement provided by a US employer reflecting wages earned by the employee.  
`US_MILITARY_ERAS`: An electronic Retirement Account Statement (eRAS) issued by the US military.  
`US_MILITARY_LES`: A Leave and Earnings Statement (LES) issued by the US military.  
`US_MILITARY_CLES`: A Civilian Leave and Earnings Statement (CLES) issued by the US military.  
`GIG`: Used to indicate that the income is related to gig work. Does not necessarily correspond to a specific document type.  
`PLAID_GENERATED_PAYSTUB_PDF`: Used to indicate that the PDF for the paystub was generated by Plaid.  
`NONE`: Used to indicate that there is no underlying document for the data.  
`UNKNOWN`: Document type could not be determined.  
  

Possible values: `UNKNOWN`, `PAYSTUB`, `BANK_STATEMENT`, `US_TAX_W2`, `US_MILITARY_ERAS`, `US_MILITARY_LES`, `US_MILITARY_CLES`, `GIG`, `PLAID_GENERATED_PAYSTUB_PDF`, `NONE`

[`download_url`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-document-metadata-download-url)

nullablestringnullable, string

Signed URL to retrieve the document(s). The payload will be a .zip file containing the document(s).  
For Payroll Income, the file type of the documents will always be PDF, and the documents may not be available, in which case the field will be `null`. If you would like Plaid to generate a PDF if the original is not available, contact your Account Manager. [Example generated pay stub](https://plaid.com/documents/plaid-generated-mock-paystub.pdf).  
For Document Income, this field will not be `null`, and the file type of the underlying document(s) will be the original file type uploaded by the user. For more details on available file types, see the [Document Income](https://plaid.com/docs/income/payroll-income) documentation.  
This download URL can only be used once and expires after two minutes. To generate a new download URL, call `/credit/payroll_income/get` again.

[`status`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-document-metadata-status)

nullablestringnullable, string

The processing status of the document.  
`PROCESSING_COMPLETE`: The document was successfully processed.  
`DOCUMENT_ERROR`: The document could not be processed. Possible causes include: The document was an unacceptable document type such as an offer letter or bank statement, the document image was cropped or blurry, or the document was corrupted.  
`UNKNOWN` or `null`: An internal error occurred. If this happens repeatedly, contact support or your Plaid account manager.  
  

Possible values: `UNKNOWN`, `PROCESSING_COMPLETE`, `DOCUMENT_ERROR`, `null`

[`page_count`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-document-metadata-page-count)

nullableintegernullable, integer

The number of pages of the uploaded document (if available).

[`error_message`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-document-metadata-error-message)

nullablestringnullable, string

The reason why a failure occurred during document processing (if available).

[`document_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-document-id)

stringstring

An identifier of the document referenced by the document metadata.

[`employer`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer)

objectobject

Information about the employer on the pay stub.

[`address`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer-address)

objectobject

Address on the pay stub.

[`city`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer-address-city)

nullablestringnullable, string

The full city name.

[`country`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer-address-postal-code)

nullablestringnullable, string

The postal code of the address.

[`region`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer-address-street)

nullablestringnullable, string

The full street address.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer-name)

nullablestringnullable, string

The name of the employer on the pay stub.

[`employee`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee)

objectobject

Data about the employee.

[`address`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-address)

objectobject

Address on the pay stub.

[`city`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-address-city)

nullablestringnullable, string

The full city name.

[`country`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-address-postal-code)

nullablestringnullable, string

The postal code of the address.

[`region`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-address-street)

nullablestringnullable, string

The full street address.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-name)

nullablestringnullable, string

The name of the employee.

[`marital_status`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-marital-status)

nullablestringnullable, string

Marital status of the employee - either `SINGLE` or `MARRIED` or `NOT LISTED`.  
  

Possible values: `SINGLE`, `MARRIED`, `NOT LISTED`, `null`

[`taxpayer_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-taxpayer-id)

objectobject

Taxpayer ID of the individual receiving the paystub.

[`id_type`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-taxpayer-id-id-type)

nullablestringnullable, string

Type of ID, e.g. 'SSN'.

[`id_mask`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employee-taxpayer-id-id-mask)

nullablestringnullable, string

ID mask; i.e. last 4 digits of the taxpayer ID.

[`tax_year`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-tax-year)

nullablestringnullable, string

The tax year of the W2 document.

[`employer_id_number`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-employer-id-number)

nullablestringnullable, string

An employee identification number or EIN.

[`wages_tips_other_comp`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-wages-tips-other-comp)

nullablestringnullable, string

Wages from tips and other compensation.

[`federal_income_tax_withheld`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-federal-income-tax-withheld)

nullablestringnullable, string

Federal income tax withheld for the tax year.

[`social_security_wages`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-social-security-wages)

nullablestringnullable, string

Wages from social security.

[`social_security_tax_withheld`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-social-security-tax-withheld)

nullablestringnullable, string

Social security tax withheld for the tax year.

[`medicare_wages_and_tips`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-medicare-wages-and-tips)

nullablestringnullable, string

Wages and tips from medicare.

[`medicare_tax_withheld`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-medicare-tax-withheld)

nullablestringnullable, string

Medicare tax withheld for the tax year.

[`social_security_tips`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-social-security-tips)

nullablestringnullable, string

Tips from social security.

[`allocated_tips`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-allocated-tips)

nullablestringnullable, string

Allocated tips.

[`box_9`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-box-9)

nullablestringnullable, string

Contents from box 9 on the W2.

[`dependent_care_benefits`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-dependent-care-benefits)

nullablestringnullable, string

Dependent care benefits.

[`nonqualified_plans`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-nonqualified-plans)

nullablestringnullable, string

Nonqualified plans.

[`box_12`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-box-12)

[object][object]

[`code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-box-12-code)

nullablestringnullable, string

W2 Box 12 code.

[`amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-box-12-amount)

nullablestringnullable, string

W2 Box 12 amount.

[`statutory_employee`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-statutory-employee)

nullablestringnullable, string

Statutory employee.

[`retirement_plan`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-retirement-plan)

nullablestringnullable, string

Retirement plan.

[`third_party_sick_pay`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-third-party-sick-pay)

nullablestringnullable, string

Third party sick pay.

[`other`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-other)

nullablestringnullable, string

Other.

[`state_and_local_wages`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-state-and-local-wages)

[object][object]

[`state`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-state-and-local-wages-state)

nullablestringnullable, string

State associated with the wage.

[`employer_state_id_number`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-state-and-local-wages-employer-state-id-number)

nullablestringnullable, string

State identification number of the employer.

[`state_wages_tips`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-state-and-local-wages-state-wages-tips)

nullablestringnullable, string

Wages and tips from the specified state.

[`state_income_tax`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-state-and-local-wages-state-income-tax)

nullablestringnullable, string

Income tax from the specified state.

[`local_wages_tips`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-state-and-local-wages-local-wages-tips)

nullablestringnullable, string

Wages and tips from the locality.

[`local_income_tax`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-state-and-local-wages-local-income-tax)

nullablestringnullable, string

Income tax from the locality.

[`locality_name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-w2s-state-and-local-wages-locality-name)

nullablestringnullable, string

Name of the locality.

[`form1099s`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s)

[object][object]

Array of tax form 1099s.

[`document_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-document-id)

nullablestringnullable, string

An identifier of the document referenced by the document metadata.

[`document_metadata`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-document-metadata)

objectobject

Object representing metadata pertaining to the document.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-document-metadata-name)

stringstring

The name of the document.

[`document_type`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-document-metadata-document-type)

nullablestringnullable, string

The type of document.  
`PAYSTUB`: A paystub.  
`BANK_STATEMENT`: A bank statement.  
`US_TAX_W2`: A W-2 wage and tax statement provided by a US employer reflecting wages earned by the employee.  
`US_MILITARY_ERAS`: An electronic Retirement Account Statement (eRAS) issued by the US military.  
`US_MILITARY_LES`: A Leave and Earnings Statement (LES) issued by the US military.  
`US_MILITARY_CLES`: A Civilian Leave and Earnings Statement (CLES) issued by the US military.  
`GIG`: Used to indicate that the income is related to gig work. Does not necessarily correspond to a specific document type.  
`PLAID_GENERATED_PAYSTUB_PDF`: Used to indicate that the PDF for the paystub was generated by Plaid.  
`NONE`: Used to indicate that there is no underlying document for the data.  
`UNKNOWN`: Document type could not be determined.  
  

Possible values: `UNKNOWN`, `PAYSTUB`, `BANK_STATEMENT`, `US_TAX_W2`, `US_MILITARY_ERAS`, `US_MILITARY_LES`, `US_MILITARY_CLES`, `GIG`, `PLAID_GENERATED_PAYSTUB_PDF`, `NONE`

[`download_url`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-document-metadata-download-url)

nullablestringnullable, string

Signed URL to retrieve the document(s). The payload will be a .zip file containing the document(s).  
For Payroll Income, the file type of the documents will always be PDF, and the documents may not be available, in which case the field will be `null`. If you would like Plaid to generate a PDF if the original is not available, contact your Account Manager. [Example generated pay stub](https://plaid.com/documents/plaid-generated-mock-paystub.pdf).  
For Document Income, this field will not be `null`, and the file type of the underlying document(s) will be the original file type uploaded by the user. For more details on available file types, see the [Document Income](https://plaid.com/docs/income/payroll-income) documentation.  
This download URL can only be used once and expires after two minutes. To generate a new download URL, call `/credit/payroll_income/get` again.

[`status`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-document-metadata-status)

nullablestringnullable, string

The processing status of the document.  
`PROCESSING_COMPLETE`: The document was successfully processed.  
`DOCUMENT_ERROR`: The document could not be processed. Possible causes include: The document was an unacceptable document type such as an offer letter or bank statement, the document image was cropped or blurry, or the document was corrupted.  
`UNKNOWN` or `null`: An internal error occurred. If this happens repeatedly, contact support or your Plaid account manager.  
  

Possible values: `UNKNOWN`, `PROCESSING_COMPLETE`, `DOCUMENT_ERROR`, `null`

[`page_count`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-document-metadata-page-count)

nullableintegernullable, integer

The number of pages of the uploaded document (if available).

[`error_message`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-document-metadata-error-message)

nullablestringnullable, string

The reason why a failure occurred during document processing (if available).

[`form_1099_type`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-form-1099-type)

stringstring

Form 1099 Type  
  

Possible values: `FORM_1099_TYPE_UNKNOWN`, `FORM_1099_TYPE_MISC`, `FORM_1099_TYPE_K`

[`recipient`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient)

objectobject

An object representing a recipient used in both 1099-K and 1099-MISC tax documents.

[`address`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-address)

objectobject

Address on the pay stub.

[`city`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-address-city)

nullablestringnullable, string

The full city name.

[`country`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-address-postal-code)

nullablestringnullable, string

The postal code of the address.

[`region`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-address-street)

nullablestringnullable, string

The full street address.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-name)

nullablestringnullable, string

Name of recipient.

[`tin`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-tin)

nullablestringnullable, string

Tax identification number of recipient.

[`account_number`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-account-number)

nullablestringnullable, string

Account number number of recipient.

[`facta_filing_requirement`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-facta-filing-requirement)

nullablestringnullable, string

Checked if FACTA is a filing requirement.  
  

Possible values: `CHECKED`, `NOT CHECKED`

[`second_tin_exists`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-recipient-second-tin-exists)

nullablestringnullable, string

Checked if 2nd TIN exists.  
  

Possible values: `CHECKED`, `NOT CHECKED`

[`payer`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer)

objectobject

An object representing a payer used by 1099-MISC tax documents.

[`address`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-address)

objectobject

Address on the pay stub.

[`city`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-address-city)

nullablestringnullable, string

The full city name.

[`country`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-address-postal-code)

nullablestringnullable, string

The postal code of the address.

[`region`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-address-street)

nullablestringnullable, string

The full street address.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-name)

nullablestringnullable, string

Name of payer.

[`tin`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-tin)

nullablestringnullable, string

Tax identification number of payer.

[`telephone_number`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-telephone-number)

nullablestringnullable, string

Telephone number of payer.

[`filer`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer)

objectobject

An object representing a filer used by 1099-K tax documents.

[`address`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-address)

objectobject

Address on the pay stub.

[`city`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-address-city)

nullablestringnullable, string

The full city name.

[`country`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-address-country)

nullablestringnullable, string

The ISO 3166-1 alpha-2 country code.

[`postal_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-address-postal-code)

nullablestringnullable, string

The postal code of the address.

[`region`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-address-region)

nullablestringnullable, string

The region or state.
Example: `"NC"`

[`street`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-address-street)

nullablestringnullable, string

The full street address.

[`name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-name)

nullablestringnullable, string

Name of filer.

[`tin`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-tin)

nullablestringnullable, string

Tax identification number of filer.

[`type`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-filer-type)

nullablestringnullable, string

One of the following values will be provided: Payment Settlement Entity (PSE), Electronic Payment Facilitator (EPF), Other Third Party  
  

Possible values: `Payment Settlement Entity (PSE)`, `Electronic Payment Facilitator (EPF)`, `Other Third Party`

[`tax_year`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-tax-year)

nullablestringnullable, string

Tax year of the tax form.

[`rents`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-rents)

nullablenumbernullable, number

Amount in rent by payer.  
  

Format: `double`

[`royalties`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-royalties)

nullablenumbernullable, number

Amount in royalties by payer.  
  

Format: `double`

[`other_income`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-other-income)

nullablenumbernullable, number

Amount in other income by payer.  
  

Format: `double`

[`federal_income_tax_withheld`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-federal-income-tax-withheld)

nullablenumbernullable, number

Amount of federal income tax withheld from payer.  
  

Format: `double`

[`fishing_boat_proceeds`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-fishing-boat-proceeds)

nullablenumbernullable, number

Amount of fishing boat proceeds from payer.  
  

Format: `double`

[`medical_and_healthcare_payments`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-medical-and-healthcare-payments)

nullablenumbernullable, number

Amount of medical and healthcare payments from payer.  
  

Format: `double`

[`nonemployee_compensation`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-nonemployee-compensation)

nullablenumbernullable, number

Amount of nonemployee compensation from payer.  
  

Format: `double`

[`substitute_payments_in_lieu_of_dividends_or_interest`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-substitute-payments-in-lieu-of-dividends-or-interest)

nullablenumbernullable, number

Amount of substitute payments made by payer.  
  

Format: `double`

[`payer_made_direct_sales_of_5000_or_more_of_consumer_products_to_buyer`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-made-direct-sales-of-5000-or-more-of-consumer-products-to-buyer)

nullablestringnullable, string

Whether or not payer made direct sales over $5000 of consumer products.

[`crop_insurance_proceeds`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-crop-insurance-proceeds)

nullablenumbernullable, number

Amount of crop insurance proceeds.  
  

Format: `double`

[`excess_golden_parachute_payments`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-excess-golden-parachute-payments)

nullablenumbernullable, number

Amount of golden parachute payments made by payer.  
  

Format: `double`

[`gross_proceeds_paid_to_an_attorney`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-gross-proceeds-paid-to-an-attorney)

nullablenumbernullable, number

Amount of gross proceeds paid to an attorney by payer.  
  

Format: `double`

[`section_409a_deferrals`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-section-409a-deferrals)

nullablenumbernullable, number

Amount of 409A deferrals earned by payer.  
  

Format: `double`

[`section_409a_income`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-section-409a-income)

nullablenumbernullable, number

Amount of 409A income earned by payer.  
  

Format: `double`

[`state_tax_withheld`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-state-tax-withheld)

nullablenumbernullable, number

Amount of state tax withheld of payer for primary state.  
  

Format: `double`

[`state_tax_withheld_lower`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-state-tax-withheld-lower)

nullablenumbernullable, number

Amount of state tax withheld of payer for secondary state.  
  

Format: `double`

[`payer_state_number`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-state-number)

nullablestringnullable, string

Primary state ID.

[`payer_state_number_lower`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-payer-state-number-lower)

nullablestringnullable, string

Secondary state ID.

[`state_income`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-state-income)

nullablenumbernullable, number

State income reported for primary state.  
  

Format: `double`

[`state_income_lower`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-state-income-lower)

nullablenumbernullable, number

State income reported for secondary state.  
  

Format: `double`

[`transactions_reported`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-transactions-reported)

nullablestringnullable, string

One of the values will be provided Payment card Third party network  
  

Possible values: `Payment card`, `Third party network`

[`pse_name`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-pse-name)

nullablestringnullable, string

Name of the PSE (Payment Settlement Entity).

[`pse_telephone_number`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-pse-telephone-number)

nullablestringnullable, string

Formatted (XXX) XXX-XXXX. Phone number of the PSE (Payment Settlement Entity).

[`gross_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-gross-amount)

nullablenumbernullable, number

Gross amount reported.  
  

Format: `double`

[`card_not_present_transaction`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-card-not-present-transaction)

nullablenumbernullable, number

Amount in card not present transactions.  
  

Format: `double`

[`merchant_category_code`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-merchant-category-code)

nullablestringnullable, string

Merchant category of filer.

[`number_of_payment_transactions`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-number-of-payment-transactions)

nullablestringnullable, string

Number of payment transactions made.

[`january_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-january-amount)

nullablenumbernullable, number

Amount reported for January.  
  

Format: `double`

[`february_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-february-amount)

nullablenumbernullable, number

Amount reported for February.  
  

Format: `double`

[`march_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-march-amount)

nullablenumbernullable, number

Amount reported for March.  
  

Format: `double`

[`april_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-april-amount)

nullablenumbernullable, number

Amount reported for April.  
  

Format: `double`

[`may_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-may-amount)

nullablenumbernullable, number

Amount reported for May.  
  

Format: `double`

[`june_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-june-amount)

nullablenumbernullable, number

Amount reported for June.  
  

Format: `double`

[`july_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-july-amount)

nullablenumbernullable, number

Amount reported for July.  
  

Format: `double`

[`august_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-august-amount)

nullablenumbernullable, number

Amount reported for August.  
  

Format: `double`

[`september_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-september-amount)

nullablenumbernullable, number

Amount reported for September.  
  

Format: `double`

[`october_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-october-amount)

nullablenumbernullable, number

Amount reported for October.  
  

Format: `double`

[`november_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-november-amount)

nullablenumbernullable, number

Amount reported for November.  
  

Format: `double`

[`december_amount`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-december-amount)

nullablenumbernullable, number

Amount reported for December.  
  

Format: `double`

[`primary_state`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-primary-state)

nullablestringnullable, string

Primary state of business.

[`secondary_state`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-secondary-state)

nullablestringnullable, string

Secondary state of business.

[`primary_state_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-primary-state-id)

nullablestringnullable, string

Primary state ID.

[`secondary_state_id`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-secondary-state-id)

nullablestringnullable, string

Secondary state ID.

[`primary_state_income_tax`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-primary-state-income-tax)

nullablenumbernullable, number

State income tax reported for primary state.  
  

Format: `double`

[`secondary_state_income_tax`](/docs/api/products/income/#credit-payroll_income-get-response-items-payroll-income-form1099s-secondary-state-income-tax)

nullablenumbernullable, number

State income tax reported for secondary state.  
  

Format: `double`

[`status`](/docs/api/products/income/#credit-payroll_income-get-response-items-status)

nullableobjectnullable, object

Details about the status of the payroll item.

[`processing_status`](/docs/api/products/income/#credit-payroll_income-get-response-items-status-processing-status)

nullablestringnullable, string

Denotes the processing status for the verification.  
`UNKNOWN`: The processing status could not be determined.  
`PROCESSING_COMPLETE`: The processing has completed and the user has approved for sharing. The data is available to be retrieved.  
`PROCESSING`: The verification is still processing. The data is not available yet.  
`FAILED`: The processing failed to complete successfully.  
`APPROVAL_STATUS_PENDING`: The processing has completed but the user has not yet approved the sharing of the data.  
  

Possible values: `UNKNOWN`, `PROCESSING_COMPLETE`, `PROCESSING`, `FAILED`, `APPROVAL_STATUS_PENDING`

[`updated_at`](/docs/api/products/income/#credit-payroll_income-get-response-items-updated-at)

nullablestringnullable, string

Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DDTHH:mm:ssZ) indicating the last time that the Item was updated.  
  

Format: `date-time`

[`error`](/docs/api/products/income/#credit-payroll_income-get-response-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/income/#credit-payroll_income-get-response-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/income/#credit-payroll_income-get-response-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/income/#credit-payroll_income-get-response-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/income/#credit-payroll_income-get-response-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/income/#credit-payroll_income-get-response-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/income/#credit-payroll_income-get-response-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/income/#credit-payroll_income-get-response-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/income/#credit-payroll_income-get-response-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/income/#credit-payroll_income-get-response-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/income/#credit-payroll_income-get-response-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/income/#credit-payroll_income-get-response-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/income/#credit-payroll_income-get-response-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`request_id`](/docs/api/products/income/#credit-payroll_income-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "items": [
    {
      "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
      "institution_id": "ins_92",
      "institution_name": "ADP",
      "accounts": [
        {
          "account_id": "GeooLPBGDEunl54q7N3ZcyD5aLPLEai1nkzM9",
          "rate_of_pay": {
            "pay_amount": 100000,
            "pay_rate": "ANNUAL"
          },
          "pay_frequency": "BIWEEKLY"
        }
      ],
      "payroll_income": [
        {
          "account_id": "GeooLPBGDEunl54q7N3ZcyD5aLPLEai1nkzM9",
          "pay_stubs": [
            {
              "deductions": {
                "breakdown": [
                  {
                    "current_amount": 123.45,
                    "description": "taxes",
                    "iso_currency_code": "USD",
                    "unofficial_currency_code": null,
                    "ytd_amount": 246.9
                  }
                ],
                "total": {
                  "current_amount": 123.45,
                  "iso_currency_code": "USD",
                  "unofficial_currency_code": null,
                  "ytd_amount": 246.9
                }
              },
              "document_metadata": {
                "document_type": "PAYSTUB",
                "name": "paystub.pdf",
                "status": "PROCESSING_COMPLETE",
                "download_url": null
              },
              "document_id": "2jkflanbd",
              "earnings": {
                "breakdown": [
                  {
                    "canonical_description": "REGULAR_PAY",
                    "current_amount": 200.22,
                    "description": "salary earned",
                    "hours": 80,
                    "iso_currency_code": "USD",
                    "rate": null,
                    "unofficial_currency_code": null,
                    "ytd_amount": 400.44
                  },
                  {
                    "canonical_description": "BONUS",
                    "current_amount": 100,
                    "description": "bonus earned",
                    "hours": null,
                    "iso_currency_code": "USD",
                    "rate": null,
                    "unofficial_currency_code": null,
                    "ytd_amount": 100
                  }
                ],
                "total": {
                  "current_amount": 300.22,
                  "hours": 160,
                  "iso_currency_code": "USD",
                  "unofficial_currency_code": null,
                  "ytd_amount": 500.44
                }
              },
              "employee": {
                "address": {
                  "city": "SAN FRANCISCO",
                  "country": "US",
                  "postal_code": "94133",
                  "region": "CA",
                  "street": "2140 TAYLOR ST"
                },
                "name": "ANNA CHARLESTON",
                "marital_status": "SINGLE",
                "taxpayer_id": {
                  "id_type": "SSN",
                  "id_mask": "3333"
                }
              },
              "employer": {
                "name": "PLAID INC",
                "address": {
                  "city": "SAN FRANCISCO",
                  "country": "US",
                  "postal_code": "94111",
                  "region": "CA",
                  "street": "1098 HARRISON ST"
                }
              },
              "net_pay": {
                "current_amount": 123.34,
                "description": "TOTAL NET PAY",
                "iso_currency_code": "USD",
                "unofficial_currency_code": null,
                "ytd_amount": 253.54
              },
              "pay_period_details": {
                "distribution_breakdown": [
                  {
                    "account_name": "Big time checking",
                    "bank_name": "bank of plaid",
                    "current_amount": 176.77,
                    "iso_currency_code": "USD",
                    "mask": "1223",
                    "type": "checking",
                    "unofficial_currency_code": null
                  }
                ],
                "end_date": "2020-12-15",
                "gross_earnings": 4500,
                "iso_currency_code": "USD",
                "pay_amount": 1490.21,
                "pay_date": "2020-12-15",
                "pay_frequency": "BIWEEKLY",
                "start_date": "2020-12-01",
                "unofficial_currency_code": null
              }
            }
          ],
          "w2s": [
            {
              "allocated_tips": "1000",
              "box_12": [
                {
                  "amount": "200",
                  "code": "AA"
                }
              ],
              "box_9": "box9",
              "dependent_care_benefits": "1000",
              "document_metadata": {
                "document_type": "US_TAX_W2",
                "download_url": null,
                "name": "w_2.pdf",
                "status": "PROCESSING_COMPLETE"
              },
              "document_id": "1pkflebk4",
              "employee": {
                "address": {
                  "city": "San Francisco",
                  "country": "US",
                  "postal_code": "94103",
                  "region": "CA",
                  "street": "1234 Grand St"
                },
                "name": "Josie Georgia Harrison",
                "marital_status": "SINGLE",
                "taxpayer_id": {
                  "id_type": "SSN",
                  "id_mask": "1234"
                }
              },
              "employer": {
                "address": {
                  "city": "New York",
                  "country": "US",
                  "postal_code": "10010",
                  "region": "NY",
                  "street": "456 Main St"
                },
                "name": "Acme Inc"
              },
              "employer_id_number": "12-1234567",
              "federal_income_tax_withheld": "1000",
              "medicare_tax_withheld": "1000",
              "medicare_wages_and_tips": "1000",
              "nonqualified_plans": "1000",
              "other": "other",
              "retirement_plan": "CHECKED",
              "social_security_tax_withheld": "1000",
              "social_security_tips": "1000",
              "social_security_wages": "1000",
              "state_and_local_wages": [
                {
                  "employer_state_id_number": "11111111111AAA",
                  "local_income_tax": "200",
                  "local_wages_and_tips": "200",
                  "locality_name": "local",
                  "state": "UT",
                  "state_income_tax": "200",
                  "state_wages_tips": "200"
                }
              ],
              "statutory_employee": "CHECKED",
              "tax_year": "2020",
              "third_party_sick_pay": "CHECKED",
              "wages_tips_other_comp": "1000"
            }
          ],
          "form1099s": [
            {
              "april_amount": null,
              "august_amount": null,
              "card_not_present_transaction": null,
              "crop_insurance_proceeds": 1000,
              "december_amount": null,
              "document_id": "mvMZ59Z2a5",
              "document_metadata": {
                "document_type": "US_TAX_1099_MISC",
                "download_url": null,
                "name": "form_1099_misc.pdf",
                "status": "PROCESSING_COMPLETE"
              },
              "excess_golden_parachute_payments": 1000,
              "february_amount": null,
              "federal_income_tax_withheld": 1000,
              "filer": {
                "address": {
                  "city": null,
                  "country": null,
                  "postal_code": null,
                  "region": null,
                  "street": null
                },
                "name": null,
                "tin": null,
                "type": null
              },
              "fishing_boat_proceeds": 1000,
              "form_1099_type": "FORM_1099_TYPE_MISC",
              "gross_amount": 1000,
              "gross_proceeds_paid_to_an_attorney": 1000,
              "january_amount": null,
              "july_amount": null,
              "june_amount": null,
              "march_amount": null,
              "may_amount": null,
              "medical_and_healthcare_payments": 1000,
              "merchant_category_code": null,
              "nonemployee_compensation": 1000,
              "november_amount": null,
              "number_of_payment_transactions": null,
              "october_amount": null,
              "other_income": 1000,
              "payer": {
                "address": {
                  "city": "SAN FRANCISCO",
                  "country": "US",
                  "postal_code": "94111",
                  "region": "CA",
                  "street": "1098 HARRISON ST"
                },
                "name": "PLAID INC",
                "telephone_number": "(123)456-7890",
                "tin": "12-3456789"
              },
              "payer_made_direct_sales_of_500_or_more_of_consumer_products_to_buyer": null,
              "payer_state_number": "CA 12345",
              "payer_state_number_lower": null,
              "primary_state": null,
              "primary_state_id": "CA 12345",
              "primary_state_income_tax": 1000,
              "pse_name": null,
              "pse_telephone_number": null,
              "recipient": {
                "account_number": "45678",
                "address": {
                  "city": "SAN FRANCISCO",
                  "country": "US",
                  "postal_code": "94133",
                  "region": "CA",
                  "street": "2140 TAYLOR ST"
                },
                "facta_filing_requirement": "CHECKED",
                "name": "Josie Georgia Harrison",
                "second_tin_exists": "NOT CHECKED",
                "tin": "12-3456789"
              },
              "rents": 1000,
              "royalties": 1000,
              "secondary_state": null,
              "secondary_state_id": null,
              "secondary_state_income_tax": null,
              "section_409a_deferrals": 1000,
              "section_409a_income": 1000,
              "september_amount": null,
              "state_income": 1000,
              "state_income_lower": null,
              "state_tax_withheld": 1000,
              "state_tax_withheld_lower": null,
              "substitute_payments_in_lieu_of_dividends_or_interest": null,
              "tax_year": "2022",
              "transactions_reported": null
            }
          ]
        }
      ],
      "status": {
        "processing_status": "PROCESSING_COMPLETE"
      },
      "updated_at": "2022-08-02T21:14:54Z"
    }
  ],
  "request_id": "2pxQ59buGdsHRef"
}
```

=\*=\*=\*=

#### `/credit/payroll_income/risk_signals/get`

#### Retrieve fraud insights for a user's manually uploaded document(s).

[`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) can be used as part of the Document Income flow to assess a user-uploaded document for signs of potential fraud or tampering. It returns a risk score for each uploaded document that indicates the likelihood of the document being fraudulent, in addition to details on the individual risk signals contributing to the score.

To trigger risk signal generation for an Item, call [`/link/token/create`](/docs/api/link/#linktokencreate) with `parsing_config` set to include `risk_signals`, or call [`/credit/payroll_income/parsing_config/update`](/docs/api/products/income/#creditpayroll_incomeparsing_configupdate). Once risk signal generation has been triggered, [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) can be called at any time after the `INCOME_VERIFICATION_RISK_SIGNALS` webhook has been fired.

[`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) is offered as an add-on to Document Income and is billed separately. To request access to this endpoint, submit a product access request or contact your Plaid account manager.

/credit/payroll\_income/risk\_signals/get

**Request fields**

[`client_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/credit/payroll\_income/risk\_signals/get

```
const request: CreditPayrollIncomeRiskSignalsGetRequest = {
  user_token: 'user-sandbox-dd4c42bd-4a81-4089-8146-40671e81dd12',
};
try {
  const response = await client.creditPayrollIncomeRiskSignalsGet(request);
} catch (error) {
  // handle error
}
```

/credit/payroll\_income/risk\_signals/get

**Response fields**

[`items`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items)

[object][object]

Array of payroll items.

[`item_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`verification_risk_signals`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals)

[object][object]

Array of payroll income document authenticity data retrieved for each of the user's accounts.

[`account_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-account-id)

nullablestringnullable, string

ID of the payroll provider account.

[`single_document_risk_signals`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals)

[object][object]

Array of document metadata and associated risk signals per document

[`document_reference`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-document-reference)

objectobject

Object containing metadata for the document

[`document_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-document-reference-document-id)

nullablestringnullable, string

An identifier of the document referenced by the document metadata.

[`document_name`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-document-reference-document-name)

stringstring

The name of the document

[`status`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-document-reference-status)

stringstring

Status of a document for risk signal analysis  
  

Possible values: `PROCESSING`, `PROCESSING_COMPLETE`, `PROCESSING_ERROR`, `PASSWORD_PROTECTED`, `VIRUS_DETECTED`

[`document_type`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-document-reference-document-type)

nullablestringnullable, string

Type of a document for risk signal analysis  
  

Possible values: `UNKNOWN`, `BANK_STATEMENT`, `BENEFITS_STATEMENT`, `BUSINESS_FILING`, `CHECK`, `DRIVING_LICENSE`, `FINANCIAL_STATEMENT`, `INVOICE`, `PAYSLIP`, `SOCIAL_SECURITY_CARD`, `TAX_FORM`, `UTILITY_BILL`

[`file_type`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-document-reference-file-type)

nullablestringnullable, string

The file type for risk signal analysis  
  

Possible values: `UNKNOWN`, `IMAGE_PDF`, `SCAN_OCR`, `TRUE_PDF`, `IMAGE`, `MIXED_PAGE_PDF`, `EMPTY_PDF`, `FLATTENED_PDF`

[`risk_signals`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals)

[object][object]

Array of attributes that indicate whether or not there is fraud risk with a document

[`type`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-type)

nullablestringnullable, string

The type of risk found in the risk signal check.  
  

Possible values: `FONT`, `MASKING`, `OVERLAID_TEXT`, `EDITED_TEXT`, `TEXT_COMPRESSION`, `ADDRESS_FORMAT_ANOMALY`, `DATE_FORMAT_ANOMALY`, `FONT_ANOMALY`, `NAME_FORMAT_ANOMALY`, `PDF_ALIGNMENT`, `BRUSH_DETECTION`, `METADATA_DATES_OUTSIDE_WINDOW`, `METADATA_DATES_INSIDE_WINDOW`, `METADATA_DATES_MISSING`, `METADATA_DATES_MATCH`, `ADOBE_FONTS`, `ANNOTATION_DATES`, `ANNOTATIONS`, `EDITED_WHILE_SCANNED`, `EXIF_DATA_MODIFIED`, `HIGH_USER_ACCESS`, `MALFORMED_DATE`, `QPDF`, `TEXT_LAYER_TEXT`, `TOUCHUP_TEXT`, `FLATTENED_PDF`, `BLACKLISTS`, `COPYCAT_IMAGE`, `COPYCAT_TEXT`, `REJECTED_CUSTOMER`, `TEMPLATES`, `SOFTWARE_BLACKLIST`

[`field`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-field)

nullablestringnullable, string

The field which the risk signal was computed for

[`has_fraud_risk`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-has-fraud-risk)

nullablebooleannullable, boolean

A flag used to quickly identify if the signal indicates that this field is authentic or fraudulent

[`institution_metadata`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-institution-metadata)

nullableobjectnullable, object

An object which contains additional metadata about the institution used to compute the verification attribute

[`item_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-institution-metadata-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`expected_value`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-expected-value)

nullablestringnullable, string

The expected value of the field, as seen on the document

[`actual_value`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-actual-value)

nullablestringnullable, string

The derived value obtained in the risk signal calculation process for this field

[`signal_description`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-signal-description)

nullablestringnullable, string

A human-readable explanation providing more detail into the particular risk signal

[`page_number`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-signals-page-number)

nullableintegernullable, integer

The relevant page associated with the risk signal. If the risk signal is not associated with a specific page, the value will be 0.

[`risk_summary`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-summary)

objectobject

A summary across all risk signals associated with a document

[`risk_score`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-single-document-risk-signals-risk-summary-risk-score)

nullablenumbernullable, number

A number between 0 and 100, inclusive, where a score closer to 0 indicates a document is likely to be trustworthy and a score closer to 100 indicates a document is likely to be fraudulent. You can automatically reject documents with a high risk score, automatically accept documents with a low risk score, and manually review documents in between. We suggest starting with a threshold of 80 for auto-rejection and 20 for auto-acceptance. As you gather more data points on typical risk scores for your use case, you can tune these parameters to reduce the number of documents undergoing manual review.

[`multi_document_risk_signals`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals)

[object][object]

Array of risk signals computed from a set of uploaded documents and the associated documents' metadata

[`document_references`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-document-references)

[object][object]

Array of objects containing attributes that could indicate if a document is fraudulent

[`document_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-document-references-document-id)

nullablestringnullable, string

An identifier of the document referenced by the document metadata.

[`document_name`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-document-references-document-name)

stringstring

The name of the document

[`status`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-document-references-status)

stringstring

Status of a document for risk signal analysis  
  

Possible values: `PROCESSING`, `PROCESSING_COMPLETE`, `PROCESSING_ERROR`, `PASSWORD_PROTECTED`, `VIRUS_DETECTED`

[`document_type`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-document-references-document-type)

nullablestringnullable, string

Type of a document for risk signal analysis  
  

Possible values: `UNKNOWN`, `BANK_STATEMENT`, `BENEFITS_STATEMENT`, `BUSINESS_FILING`, `CHECK`, `DRIVING_LICENSE`, `FINANCIAL_STATEMENT`, `INVOICE`, `PAYSLIP`, `SOCIAL_SECURITY_CARD`, `TAX_FORM`, `UTILITY_BILL`

[`file_type`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-document-references-file-type)

nullablestringnullable, string

The file type for risk signal analysis  
  

Possible values: `UNKNOWN`, `IMAGE_PDF`, `SCAN_OCR`, `TRUE_PDF`, `IMAGE`, `MIXED_PAGE_PDF`, `EMPTY_PDF`, `FLATTENED_PDF`

[`risk_signals`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals)

[object][object]

Array of attributes that indicate whether or not there is fraud risk with a set of documents

[`type`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-type)

nullablestringnullable, string

The type of risk found in the risk signal check.  
  

Possible values: `FONT`, `MASKING`, `OVERLAID_TEXT`, `EDITED_TEXT`, `TEXT_COMPRESSION`, `ADDRESS_FORMAT_ANOMALY`, `DATE_FORMAT_ANOMALY`, `FONT_ANOMALY`, `NAME_FORMAT_ANOMALY`, `PDF_ALIGNMENT`, `BRUSH_DETECTION`, `METADATA_DATES_OUTSIDE_WINDOW`, `METADATA_DATES_INSIDE_WINDOW`, `METADATA_DATES_MISSING`, `METADATA_DATES_MATCH`, `ADOBE_FONTS`, `ANNOTATION_DATES`, `ANNOTATIONS`, `EDITED_WHILE_SCANNED`, `EXIF_DATA_MODIFIED`, `HIGH_USER_ACCESS`, `MALFORMED_DATE`, `QPDF`, `TEXT_LAYER_TEXT`, `TOUCHUP_TEXT`, `FLATTENED_PDF`, `BLACKLISTS`, `COPYCAT_IMAGE`, `COPYCAT_TEXT`, `REJECTED_CUSTOMER`, `TEMPLATES`, `SOFTWARE_BLACKLIST`

[`field`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-field)

nullablestringnullable, string

The field which the risk signal was computed for

[`has_fraud_risk`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-has-fraud-risk)

nullablebooleannullable, boolean

A flag used to quickly identify if the signal indicates that this field is authentic or fraudulent

[`institution_metadata`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-institution-metadata)

nullableobjectnullable, object

An object which contains additional metadata about the institution used to compute the verification attribute

[`item_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-institution-metadata-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`expected_value`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-expected-value)

nullablestringnullable, string

The expected value of the field, as seen on the document

[`actual_value`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-actual-value)

nullablestringnullable, string

The derived value obtained in the risk signal calculation process for this field

[`signal_description`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-signal-description)

nullablestringnullable, string

A human-readable explanation providing more detail into the particular risk signal

[`page_number`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-items-verification-risk-signals-multi-document-risk-signals-risk-signals-page-number)

nullableintegernullable, integer

The relevant page associated with the risk signal. If the risk signal is not associated with a specific page, the value will be 0.

[`error`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`request_id`](/docs/api/products/income/#credit-payroll_income-risk_signals-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "items": [
    {
      "item_id": "testItemID",
      "verification_risk_signals": [
        {
          "account_id": null,
          "multi_document_risk_signals": [],
          "single_document_risk_signals": [
            {
              "document_reference": {
                "document_id": "lRepoQjxlJ1nz",
                "document_name": "Paystub.pdf",
                "file_type": "TRUE_PDF"
              },
              "risk_summary": {
                "risk_score": 70
              },
              "risk_signals": [
                {
                  "actual_value": "0.00",
                  "expected_value": "25.09",
                  "field": null,
                  "signal_description": null,
                  "has_fraud_risk": true,
                  "type": "MASKING",
                  "page_number": 1,
                  "institution_metadata": {
                    "item_id": "testItemID"
                  }
                },
                {
                  "actual_value": null,
                  "expected_value": null,
                  "field": null,
                  "signal_description": "Creation date and modification date do not match",
                  "has_fraud_risk": true,
                  "institution_metadata": null,
                  "type": "METADATA_DATES_OUTSIDE_WINDOW",
                  "page_number": 0
                }
              ]
            }
          ]
        }
      ]
    }
  ],
  "request_id": "LhQf0THi8SH1yJm"
}
```

=\*=\*=\*=

#### `/credit/employment/get`

#### Retrieve a summary of an individual's employment information

[`/credit/employment/get`](/docs/api/products/income/#creditemploymentget) returns a list of items with employment information from a user's payroll provider that was verified by an end user.

/credit/employment/get

**Request fields**

[`client_id`](/docs/api/products/income/#credit-employment-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-employment-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-employment-get-request-user-token)

requiredstringrequired, string

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/credit/employment/get

```
const request: CreditEmploymentGetRequest = {
  user_token: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
};

try {
  const response = await client.creditEmploymentGet(request);
} catch (error) {
  // handle error
}
```

/credit/employment/get

**Response fields**

[`items`](/docs/api/products/income/#credit-employment-get-response-items)

[object][object]

Array of employment items.

[`item_id`](/docs/api/products/income/#credit-employment-get-response-items-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`employments`](/docs/api/products/income/#credit-employment-get-response-items-employments)

[object][object]

[`account_id`](/docs/api/products/income/#credit-employment-get-response-items-employments-account-id)

nullablestringnullable, string

ID of the payroll provider account.

[`status`](/docs/api/products/income/#credit-employment-get-response-items-employments-status)

nullablestringnullable, string

Current employment status.  
  

Possible values: `ACTIVE`, `INACTIVE`, `null`

[`start_date`](/docs/api/products/income/#credit-employment-get-response-items-employments-start-date)

nullablestringnullable, string

Start of employment in ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`end_date`](/docs/api/products/income/#credit-employment-get-response-items-employments-end-date)

nullablestringnullable, string

End of employment, if applicable. Provided in ISO 8601 format (YYY-MM-DD).  
  

Format: `date`

[`employer`](/docs/api/products/income/#credit-employment-get-response-items-employments-employer)

objectobject

An object containing employer data.

[`name`](/docs/api/products/income/#credit-employment-get-response-items-employments-employer-name)

nullablestringnullable, string

Name of employer.

[`title`](/docs/api/products/income/#credit-employment-get-response-items-employments-title)

nullablestringnullable, string

Current title of employee.

[`platform_ids`](/docs/api/products/income/#credit-employment-get-response-items-employments-platform-ids)

objectobject

The object containing a set of ids related to an employee.

[`employee_id`](/docs/api/products/income/#credit-employment-get-response-items-employments-platform-ids-employee-id)

nullablestringnullable, string

The ID of an employee as given by their employer.

[`payroll_id`](/docs/api/products/income/#credit-employment-get-response-items-employments-platform-ids-payroll-id)

nullablestringnullable, string

The ID of an employee as given by their payroll.

[`position_id`](/docs/api/products/income/#credit-employment-get-response-items-employments-platform-ids-position-id)

nullablestringnullable, string

The ID of the position of the employee.

[`employee_type`](/docs/api/products/income/#credit-employment-get-response-items-employments-employee-type)

nullablestringnullable, string

The type of employment for the individual.
`"FULL_TIME"`: A full-time employee.
`"PART_TIME"`: A part-time employee.
`"CONTRACTOR"`: An employee typically hired externally through a contracting group.
`"TEMPORARY"`: A temporary employee.
`"OTHER"`: The employee type is not one of the above defined types.  
  

Possible values: `FULL_TIME`, `PART_TIME`, `CONTRACTOR`, `TEMPORARY`, `OTHER`, `null`

[`last_paystub_date`](/docs/api/products/income/#credit-employment-get-response-items-employments-last-paystub-date)

nullablestringnullable, string

The date of the employee's most recent paystub in ISO 8601 format (YYYY-MM-DD).  
  

Format: `date`

[`request_id`](/docs/api/products/income/#credit-employment-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "items": [
    {
      "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
      "employments": [
        {
          "account_id": "GeooLPBGDEunl54q7N3ZcyD5aLPLEai1nkzM9",
          "status": "ACTIVE",
          "start_date": "2020-01-01",
          "end_date": null,
          "employer": {
            "name": "Plaid Inc"
          },
          "title": "Software Engineer",
          "platform_ids": {
            "employee_id": "1234567",
            "position_id": "8888",
            "payroll_id": "1234567"
          },
          "employee_type": "FULL_TIME",
          "last_paystub_date": "2022-01-15"
        }
      ]
    }
  ],
  "request_id": "LhQf0THi8SH1yJm"
}
```

=\*=\*=\*=

#### `/credit/payroll_income/parsing_config/update`

#### Update the parsing configuration for a document income verification

[`/credit/payroll_income/parsing_config/update`](/docs/api/products/income/#creditpayroll_incomeparsing_configupdate) updates the parsing configuration for a document income verification.

/credit/payroll\_income/parsing\_config/update

**Request fields**

[`client_id`](/docs/api/products/income/#credit-payroll_income-parsing_config-update-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-payroll_income-parsing_config-update-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-payroll_income-parsing_config-update-request-user-token)

requiredstringrequired, string

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`item_id`](/docs/api/products/income/#credit-payroll_income-parsing_config-update-request-item-id)

stringstring

The `item_id` of the Item associated with this webhook, warning, or error

[`parsing_config`](/docs/api/products/income/#credit-payroll_income-parsing_config-update-request-parsing-config)

required[string]required, [string]

The types of analysis to enable for the document income verification session  
  

Possible values: `ocr`, `risk_signals`

/credit/payroll\_income/parsing\_config/update

```
const request: CreditPayrollIncomeParsingConfigUpdateRequest = {
  user_token: 'user-sandbox-dd4c42bd-4a81-4089-8146-40671e81dd12',
  parsing_config: ['fraud_risk'],
};
try {
  const response = await client.creditPayrollIncomeParsingConfigUpdate(request);
} catch (error) {
  // handle error
}
```

/credit/payroll\_income/parsing\_config/update

**Response fields**

[`request_id`](/docs/api/products/income/#credit-payroll_income-parsing_config-update-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "LhQf0THi8SH1yJm"
}
```

=\*=\*=\*=

#### `/credit/payroll_income/refresh`

#### Refresh a digital payroll income verification

[`/credit/payroll_income/refresh`](/docs/api/products/income/#creditpayroll_incomerefresh) refreshes a given digital payroll income verification.

/credit/payroll\_income/refresh

**Request fields**

[`client_id`](/docs/api/products/income/#credit-payroll_income-refresh-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/income/#credit-payroll_income-refresh-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/products/income/#credit-payroll_income-refresh-request-user-token)

requiredstringrequired, string

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`options`](/docs/api/products/income/#credit-payroll_income-refresh-request-options)

objectobject

An optional object for `/credit/payroll_income/refresh` request options.

[`item_ids`](/docs/api/products/income/#credit-payroll_income-refresh-request-options-item-ids)

[string][string]

An array of `item_id`s to be refreshed. Each `item_id` should uniquely identify a payroll income item. If this field is not provided, all `item_id`s associated with the `user_token` will be refreshed.

[`webhook`](/docs/api/products/income/#credit-payroll_income-refresh-request-options-webhook)

stringstring

The URL where Plaid will send the payroll income refresh webhook.

/credit/payroll\_income/refresh

```
const request: CreditPayrollIncomeRefreshRequest = {
  user_token: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
};

try {
  const response = await client.creditPayrollIncomeRefresh(request);
} catch (error) {
  // handle error
}
```

/credit/payroll\_income/refresh

**Response fields**

[`request_id`](/docs/api/products/income/#credit-payroll_income-refresh-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`verification_refresh_status`](/docs/api/products/income/#credit-payroll_income-refresh-response-verification-refresh-status)

stringstring

The verification refresh status. One of the following:  
`"USER_PRESENCE_REQUIRED"` User presence is required to refresh an income verification.
`"SUCCESSFUL"` The income verification refresh was successful.
`"NOT_FOUND"` No new data was found after the income verification refresh.  
  

Possible values: `USER_PRESENCE_REQUIRED`, `SUCCESSFUL`, `NOT_FOUND`

Response Object

```
{
  "request_id": "nTkbCH41HYmpbm5",
  "verification_refresh_status": "USER_PRESENCE_REQUIRED"
}
```

### Webhooks

Income webhooks are sent to indicate when an income verification or document fraud risk evaluation has finished processing.

=\*=\*=\*=

#### `INCOME_VERIFICATION`

Fired when the status of an income verification instance has changed. This webhook is fired for both the Document and Payroll Income flows, but not the Bank Income flow. It will typically take several minutes for this webhook to fire after the end user has uploaded their documents in the Document Income flow.

**Properties**

[`webhook_type`](/docs/api/products/income/#IncomeVerificationStatusWebhook-webhook-type)

stringstring

`"INCOME"`

[`webhook_code`](/docs/api/products/income/#IncomeVerificationStatusWebhook-webhook-code)

stringstring

`INCOME_VERIFICATION`

[`item_id`](/docs/api/products/income/#IncomeVerificationStatusWebhook-item-id)

stringstring

The Item ID associated with the verification.

[`user_id`](/docs/api/products/income/#IncomeVerificationStatusWebhook-user-id)

stringstring

The Plaid `user_id` of the User associated with this webhook, warning, or error.

[`verification_status`](/docs/api/products/income/#IncomeVerificationStatusWebhook-verification-status)

stringstring

`VERIFICATION_STATUS_PROCESSING_COMPLETE`: The income verification processing has completed. This indicates that the documents have been parsed successfully or that the documents were not parsable. If the user uploaded multiple documents, this webhook will fire when all documents have finished processing. Call the `/credit/payroll_income/get` endpoint and check the document metadata to see which documents were successfully parsed.  
`VERIFICATION_STATUS_PROCESSING_FAILED`: An unexpected internal error occurred when attempting to process the verification documentation.  
`VERIFICATION_STATUS_PENDING_APPROVAL`: (deprecated) The income verification has been sent to the user for review.

[`environment`](/docs/api/products/income/#IncomeVerificationStatusWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "INCOME",
  "webhook_code": "INCOME_VERIFICATION",
  "item_id": "gAXlMgVEw5uEGoQnnXZ6tn9E7Mn3LBc4PJVKZ",
  "user_id": "9eaba3c2fdc916bc197f279185b986607dd21682a5b04eab04a5a03e8b3f3334",
  "verification_status": "VERIFICATION_STATUS_PROCESSING_COMPLETE",
  "environment": "production"
}
```

=\*=\*=\*=

#### `INCOME_VERIFICATION_RISK_SIGNALS`

Fired when risk signals have been processed for documents uploaded via Document Income. It will typically take a minute or two for this webhook to fire after the end user has uploaded their documents in the Document Income flow. Once this webhook has fired, [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) may then be called to determine whether the documents were successfully processed and to retrieve risk data.

**Properties**

[`webhook_type`](/docs/api/products/income/#IncomeVerificationRiskSignalsStatusWebhook-webhook-type)

stringstring

`"INCOME"`

[`webhook_code`](/docs/api/products/income/#IncomeVerificationRiskSignalsStatusWebhook-webhook-code)

stringstring

`INCOME_VERIFICATION_RISK_SIGNALS`

[`item_id`](/docs/api/products/income/#IncomeVerificationRiskSignalsStatusWebhook-item-id)

stringstring

The Item ID associated with the verification.

[`user_id`](/docs/api/products/income/#IncomeVerificationRiskSignalsStatusWebhook-user-id)

stringstring

The Plaid `user_id` of the User associated with this webhook, warning, or error.

[`risk_signals_status`](/docs/api/products/income/#IncomeVerificationRiskSignalsStatusWebhook-risk-signals-status)

stringstring

`RISK_SIGNALS_PROCESSING_COMPLETE`: The income verification fraud detection processing has completed. If the user uploaded multiple documents, this webhook will fire when all documents have finished processing. Call the `/credit/payroll_income/risk_signals/get` endpoint to get all risk signal data.

[`environment`](/docs/api/products/income/#IncomeVerificationRiskSignalsStatusWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "INCOME",
  "webhook_code": "INCOME_VERIFICATION_RISK_SIGNALS",
  "item_id": "gAXlMgVEw5uEGoQnnXZ6tn9E7Mn3LBc4PJVKZ",
  "user_id": "9eaba3c2fdc916bc197f279185b986607dd21682a5b04eab04a5a03e8b3f3334",
  "risk_signals_status": "RISK_SIGNALS_PROCESSING_COMPLETE",
  "environment": "production"
}
```

=\*=\*=\*=

#### `BANK_INCOME_REFRESH_COMPLETE`

Fired when a refreshed bank income report has finished generating or failed to generate, triggered by calling [`/credit/bank_income/refresh`](/docs/api/products/income/#creditbank_incomerefresh). To get this webhook, subscribe via the [Dashboard](https://dashboard.plaid.com/developers/webhooks).

**Properties**

[`webhook_type`](/docs/api/products/income/#BankIncomeRefreshCompleteWebhook-webhook-type)

stringstring

`INCOME`

[`webhook_code`](/docs/api/products/income/#BankIncomeRefreshCompleteWebhook-webhook-code)

stringstring

`BANK_INCOME_REFRESH_COMPLETE`

[`user_id`](/docs/api/products/income/#BankIncomeRefreshCompleteWebhook-user-id)

stringstring

The `user_id` corresponding to the user the webhook has fired for.

[`result`](/docs/api/products/income/#BankIncomeRefreshCompleteWebhook-result)

stringstring

The result of the bank income refresh report generation  
`SUCCESS`: The refreshed report was successfully generated and can be retrieved via `/credit/bank_income/get`.  
`FAILURE`: The refreshed report failed to be generated  
  

Possible values: `SUCCESS`, `FAILURE`

[`environment`](/docs/api/products/income/#BankIncomeRefreshCompleteWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "INCOME",
  "webhook_code": "BANK_INCOME_REFRESH_COMPLETE",
  "user_id": "wz666MBjYWTp2PDzzggYhM6oWWmBb",
  "result": "SUCCESS",
  "environment": "production"
}
```

=\*=\*=\*=

#### `INCOME_VERIFICATION_REFRESH_RECONNECT_NEEDED`

Fired when the attempt to refresh Payroll Income data for a user via [`/credit/payroll_income/refresh`](/docs/api/products/income/#creditpayroll_incomerefresh) failed because the user must re-connect their payroll account.

**Properties**

[`webhook_type`](/docs/api/products/income/#IncomeVerificationRefreshReconnectNeededWebhook-webhook-type)

stringstring

`INCOME`

[`webhook_code`](/docs/api/products/income/#IncomeVerificationRefreshReconnectNeededWebhook-webhook-code)

stringstring

`INCOME_VERIFICATION_REFRESH_RECONNECT_NEEDED`

[`user_id`](/docs/api/products/income/#IncomeVerificationRefreshReconnectNeededWebhook-user-id)

stringstring

The `user_id` corresponding to the user the webhook has fired for.

[`environment`](/docs/api/products/income/#IncomeVerificationRefreshReconnectNeededWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "INCOME",
  "webhook_code": "INCOME_VERIFICATION_REFRESH_RECONNECT_NEEDED",
  "user_id": "wz666MBjYWTp2PDzzggYhM6oWWmBb",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
