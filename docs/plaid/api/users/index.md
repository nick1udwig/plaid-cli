---
title: "API - Users | Plaid Docs"
source_url: "https://plaid.com/docs/api/users/"
scraped_at: "2026-03-07T22:04:27+00:00"
---

# User endpoints

#### API reference for user management endpoints

This page covers API endpoints related to user ids and user tokens, which are used by [Plaid Check](/docs/check/), Plaid Protect, and [Income](/docs/income/), as well as by the [Multi-Item Link flow](/docs/link/multi-item-link/).

| Endpoints |  |
| --- | --- |
| [`/user/create`](/docs/api/users/#usercreate) | Create a user ID |
| [`/user/get`](/docs/api/users/#userget) | Get user details |
| [`/user/update`](/docs/api/users/#userupdate) | Update user data or enable an existing user for Plaid Check |
| [`/user/remove`](/docs/api/users/#userremove) | Remove the user and their associated Items |
| [`/user/items/get`](/docs/api/users/#useritemsget) | Return Items associated with a user |

| See also |  |
| --- | --- |
| [`/sandbox/user/reset_login`](/docs/api/sandbox/#sandboxuserreset_login) | Force user into an error state for testing |

Plaid has switched from using the `user_token` to the `user_id` as of December 10, 2025. If you have an existing integration that uses the `user_token`, you should continue to use the `user_token`. New customers must use the `user_id`.

To accommodate both flows, many Plaid API endpoints use either a `user_id` or a `user_token`. In these API requests, you should provide a `user_token` only if you have one; otherwise, provide a `user_id`. You do not need to provide both fields.

For more details on this change, see [New User APIs](/docs/api/users/user-apis/).

=\*=\*=\*=

#### User identifiers

There are three identifiers Plaid uses when working with users:

`client_user_id`: The unique identifier you provide to Plaid for each end user in your application. This is determined by you and will typically correspond to your application's primary key for a user record. It must not contain PII, such as a phone number, Social Security number, or email address.

`user_id`: Plaid-generated identifier, prefixed by `usr_` and returned in responses from the User APIs and [`/link/token/create`](/docs/api/link/#linktokencreate). Each `user_id` corresponds to a single `client_user_id`; calling [`/user/create`](/docs/api/users/#usercreate) repeatedly with the same `client_user_id` returns the same `user_id`. The `user_id` is used by customers who integrated on December 10, 2025 or later.

`user_token` — A Plaid-generated token for accessing the user. Only applicable for customers who integrated with [`/user/create`](/docs/api/users/#usercreate) prior to December 10, 2025. Note that `user_tokens` can also have corresponding `user_id`s; these are not equivalent to the new `usr_` `user_id`s mentioned above. For more details, see [New User APIs](/docs/api/users/user-apis/).

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

[`client_id`](/docs/api/users/#user-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/users/#user-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_user_id`](/docs/api/users/#user-create-request-client-user-id)

requiredstringrequired, string

A unique ID representing the end user. Maximum of 128 characters. Typically this will be a user ID number from your application. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.  
  

Max length: `128`

Min length: `1`

[`identity`](/docs/api/users/#user-create-request-identity)

objectobject

The identity fields associated with a user. For a user to be eligible for a Plaid Check Consumer Report, all fields are required except `id_number`. Providing a partial SSN is strongly recommended, and improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests. If creating a report that will be shared with GSEs such as Fannie or Freddie, a full Social Security Number must be provided via the `id_number` field.

[`name`](/docs/api/users/#user-create-request-identity-name)

objectobject

User name information.

[`given_name`](/docs/api/users/#user-create-request-identity-name-given-name)

requiredstringrequired, string

User's given name.

[`family_name`](/docs/api/users/#user-create-request-identity-name-family-name)

requiredstringrequired, string

User's family name.

[`date_of_birth`](/docs/api/users/#user-create-request-identity-date-of-birth)

stringstring

The user's date of birth, to be provided in the format "yyyy-mm-dd".  
  

Format: `date`

[`emails`](/docs/api/users/#user-create-request-identity-emails)

[object][object]

The user's emails.

[`data`](/docs/api/users/#user-create-request-identity-emails-data)

requiredstringrequired, string

User's email.

[`primary`](/docs/api/users/#user-create-request-identity-emails-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary email for the User.

[`phone_numbers`](/docs/api/users/#user-create-request-identity-phone-numbers)

[object][object]

The user's phone numbers, in E.164 format: +{countrycode}{number}. For example: "+14157452130". Phone numbers provided in other formats will be parsed on a best-effort basis. Phone number input is validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`data`](/docs/api/users/#user-create-request-identity-phone-numbers-data)

requiredstringrequired, string

User's phone number.

[`primary`](/docs/api/users/#user-create-request-identity-phone-numbers-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary phone number for the User.

[`addresses`](/docs/api/users/#user-create-request-identity-addresses)

[object][object]

The user's addresses.

[`street_1`](/docs/api/users/#user-create-request-identity-addresses-street-1)

stringstring

First line of street address.

[`street_2`](/docs/api/users/#user-create-request-identity-addresses-street-2)

stringstring

Second line of street address.

[`city`](/docs/api/users/#user-create-request-identity-addresses-city)

stringstring

City name.

[`region`](/docs/api/users/#user-create-request-identity-addresses-region)

stringstring

State, province or region.

[`country`](/docs/api/users/#user-create-request-identity-addresses-country)

requiredstringrequired, string

Country code.

[`postal_code`](/docs/api/users/#user-create-request-identity-addresses-postal-code)

stringstring

Postal or ZIP code.

[`primary`](/docs/api/users/#user-create-request-identity-addresses-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary address for the User.

[`id_numbers`](/docs/api/users/#user-create-request-identity-id-numbers)

[object][object]

The user's ID numbers.

[`value`](/docs/api/users/#user-create-request-identity-id-numbers-value)

requiredstringrequired, string

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/users/#user-create-request-identity-id-numbers-type)

requiredstringrequired, string

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`consumer_report_user_identity`](/docs/api/users/#user-create-request-consumer-report-user-identity)

objectobject

This field is only used by integrations created before December 10, 2025. All other integrations must use the `identity` object instead. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).
To create a Plaid Check Consumer Report for a user when using a `user_token`, this field must be present. If this field is not provided during user token creation, you can add it to the user later by calling `/user/update`. Once the field has been added to the user, you will be able to call `/link/token/create` with a non-empty `consumer_report_permissible_purpose` (which will automatically create a Plaid Check Consumer Report), or call `/cra/check_report/create` for that user.

[`first_name`](/docs/api/users/#user-create-request-consumer-report-user-identity-first-name)

requiredstringrequired, string

The user's first name

[`last_name`](/docs/api/users/#user-create-request-consumer-report-user-identity-last-name)

requiredstringrequired, string

The user's last name

[`phone_numbers`](/docs/api/users/#user-create-request-consumer-report-user-identity-phone-numbers)

required[string]required, [string]

The user's phone number, in E.164 format: +{countrycode}{number}. For example: "+14157452130". Phone numbers provided in other formats will be parsed on a best-effort basis. Phone number input is validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`emails`](/docs/api/users/#user-create-request-consumer-report-user-identity-emails)

required[string]required, [string]

The user's emails

[`ssn_full`](/docs/api/users/#user-create-request-consumer-report-user-identity-ssn-full)

stringstring

The user's full social security number. This field should only be provided by lenders intending to share the resulting consumer report with a Government-Sponsored Enterprise (GSE), such as Fannie Mae or Freddie Mac.  
Format: "ddd-dd-dddd"

[`ssn_last_4`](/docs/api/users/#user-create-request-consumer-report-user-identity-ssn-last-4)

stringstring

The last 4 digits of the user's social security number.  
  

Max length: `4`

Min length: `4`

[`date_of_birth`](/docs/api/users/#user-create-request-consumer-report-user-identity-date-of-birth)

requiredstringrequired, string

To be provided in the format "yyyy-mm-dd".
This field is required for all Plaid Check customers.  
  

Format: `date`

[`primary_address`](/docs/api/users/#user-create-request-consumer-report-user-identity-primary-address)

requiredobjectrequired, object

Data about the components comprising an address.

[`city`](/docs/api/users/#user-create-request-consumer-report-user-identity-primary-address-city)

requiredstringrequired, string

The full city name

[`region`](/docs/api/users/#user-create-request-consumer-report-user-identity-primary-address-region)

requiredstringrequired, string

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/api/users/#user-create-request-consumer-report-user-identity-primary-address-street)

requiredstringrequired, string

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/api/users/#user-create-request-consumer-report-user-identity-primary-address-postal-code)

requiredstringrequired, string

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/users/#user-create-request-consumer-report-user-identity-primary-address-country)

requiredstringrequired, string

The ISO 3166-1 alpha-2 country code

/user/create

```
const request = {
  client_user_id: "c0e2c4ee-b763-4af5-cfe9-46a46bce883d",
  identity: {
    name: {
      given_name: "Carmen",
      family_name: "Berzatto"
    },
    date_of_birth: "1987-01-31",
    emails: [
      { data: "carmy@example.com", primary: true },
      { data: "bear@example.com", primary: false }
    ],
    phone_numbers: [
      { data: "+13125551212", primary: true }
    ],
    addresses: [
      {
        street_1: "3200 W Armitage Ave",
        street_2: null,
        city: "Chicago",
        region: "IL",
        country: "US",
        postal_code: "60657",
        primary: true
      }
    ],
    id_numbers: [
      {
        value: "1234",
        type: "us_ssn_last_4"
      }
    ]
  }
};

try {
  const response = await client.userCreate(request);
} catch (error) {
}
```

/user/create

**Response fields**

[`user_token`](/docs/api/users/#user-create-response-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`user_id`](/docs/api/users/#user-create-response-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`request_id`](/docs/api/users/#user-create-response-request-id)

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

#### `/user/get`

#### Retrieve user identity and information

Get user details using a `user_id`. This endpoint only supports users that were created on the new user API flow, without a `user_token`. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/user/get

**Request fields**

[`client_id`](/docs/api/users/#user-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/users/#user-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/users/#user-get-request-user-id)

requiredstringrequired, string

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

/user/get

```
const request = {
  user_id: "usr_9nSp2KuZ2x4JDw"
};

try {
  const response = await client.userGet(request);
} catch (error) {
}
```

/user/get

**Response fields**

[`request_id`](/docs/api/users/#user-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`user_id`](/docs/api/users/#user-get-response-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`client_user_id`](/docs/api/users/#user-get-response-client-user-id)

nullablestringnullable, string

Client provided user ID.

[`created_at`](/docs/api/users/#user-get-response-created-at)

stringstring

Timestamp of user creation.  
  

Format: `date-time`

[`updated_at`](/docs/api/users/#user-get-response-updated-at)

stringstring

Timestamp of last user update.  
  

Format: `date-time`

[`identity`](/docs/api/users/#user-get-response-identity)

nullableobjectnullable, object

The identity fields associated with a user. For a user to be eligible for a Plaid Check Consumer Report, all fields are required except `id_number`. Providing a partial SSN is strongly recommended, and improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests. If creating a report that will be shared with GSEs such as Fannie or Freddie, a full Social Security Number must be provided via the `id_number` field.

[`name`](/docs/api/users/#user-get-response-identity-name)

nullableobjectnullable, object

User name information.

[`given_name`](/docs/api/users/#user-get-response-identity-name-given-name)

stringstring

User's given name.

[`family_name`](/docs/api/users/#user-get-response-identity-name-family-name)

stringstring

User's family name.

[`date_of_birth`](/docs/api/users/#user-get-response-identity-date-of-birth)

nullablestringnullable, string

The user's date of birth, to be provided in the format "yyyy-mm-dd".  
  

Format: `date`

[`emails`](/docs/api/users/#user-get-response-identity-emails)

[object][object]

The user's emails.

[`data`](/docs/api/users/#user-get-response-identity-emails-data)

stringstring

User's email.

[`primary`](/docs/api/users/#user-get-response-identity-emails-primary)

booleanboolean

Indicates whether this is the primary email for the User.

[`phone_numbers`](/docs/api/users/#user-get-response-identity-phone-numbers)

[object][object]

The user's phone numbers, in E.164 format: +{countrycode}{number}. For example: "+14157452130". Phone numbers provided in other formats will be parsed on a best-effort basis. Phone number input is validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`data`](/docs/api/users/#user-get-response-identity-phone-numbers-data)

stringstring

User's phone number.

[`primary`](/docs/api/users/#user-get-response-identity-phone-numbers-primary)

booleanboolean

Indicates whether this is the primary phone number for the User.

[`addresses`](/docs/api/users/#user-get-response-identity-addresses)

[object][object]

The user's addresses.

[`street_1`](/docs/api/users/#user-get-response-identity-addresses-street-1)

nullablestringnullable, string

First line of street address.

[`street_2`](/docs/api/users/#user-get-response-identity-addresses-street-2)

nullablestringnullable, string

Second line of street address.

[`city`](/docs/api/users/#user-get-response-identity-addresses-city)

nullablestringnullable, string

City name.

[`region`](/docs/api/users/#user-get-response-identity-addresses-region)

nullablestringnullable, string

State, province or region.

[`country`](/docs/api/users/#user-get-response-identity-addresses-country)

stringstring

Country code.

[`postal_code`](/docs/api/users/#user-get-response-identity-addresses-postal-code)

nullablestringnullable, string

Postal or ZIP code.

[`primary`](/docs/api/users/#user-get-response-identity-addresses-primary)

booleanboolean

Indicates whether this is the primary address for the User.

[`id_numbers`](/docs/api/users/#user-get-response-identity-id-numbers)

[object][object]

The user's ID numbers.

[`value`](/docs/api/users/#user-get-response-identity-id-numbers-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/users/#user-get-response-identity-id-numbers-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

Response Object

```
{
  "user_id": "usr_8c6ZbDAYjacUXF",
  "client_user_id": "uid_12345",
  "created_at": "2019-02-15T15:51:39Z",
  "updated_at": "2019-02-15T15:52:39Z",
  "request_id": "m8MDnv9okwxFNBV",
  "identity": {
    "name": {
      "given_name": "Alice",
      "family_name": "Johnson"
    },
    "date_of_birth": "1988-07-22",
    "emails": [
      {
        "data": "alice.johnson@example.com",
        "primary": true
      },
      {
        "data": "alice.j@workmail.com",
        "primary": false
      }
    ],
    "phone_numbers": [
      {
        "data": "+15551234567",
        "primary": true
      },
      {
        "data": "+15559876543",
        "primary": false
      }
    ],
    "addresses": [
      {
        "street_1": "123 Main St",
        "street_2": "Apt 4B",
        "city": "Anytown",
        "region": "CA",
        "country": "US",
        "postal_code": "90210",
        "primary": true
      }
    ],
    "id_numbers": [
      {
        "value": "1234",
        "type": "us_ssn_last_4"
      }
    ]
  }
}
```

=\*=\*=\*=

#### `/user/update`

#### Update user information

This endpoint updates user information for an existing `user_id` or `user_token`. If an existing `user_id` or `user_token` is missing fields required for a given use case (e.g. creating a Consumer Report) use [`/user/update`](/docs/api/users/#userupdate) to add values for those fields.

Identity updates use merge semantics: provided fields overwrite existing ones; omitted fields remain unchanged.

/user/update

**Request fields**

[`client_id`](/docs/api/users/#user-update-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/users/#user-update-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/users/#user-update-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`identity`](/docs/api/users/#user-update-request-identity)

objectobject

The identity fields associated with a user. For a user to be eligible for a Plaid Check Consumer Report, all fields are required except `id_number`. Providing a partial SSN is strongly recommended, and improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests. If creating a report that will be shared with GSEs such as Fannie or Freddie, a full Social Security Number must be provided via the `id_number` field.

[`name`](/docs/api/users/#user-update-request-identity-name)

objectobject

User name information.

[`given_name`](/docs/api/users/#user-update-request-identity-name-given-name)

requiredstringrequired, string

User's given name.

[`family_name`](/docs/api/users/#user-update-request-identity-name-family-name)

requiredstringrequired, string

User's family name.

[`date_of_birth`](/docs/api/users/#user-update-request-identity-date-of-birth)

stringstring

The user's date of birth, to be provided in the format "yyyy-mm-dd".  
  

Format: `date`

[`emails`](/docs/api/users/#user-update-request-identity-emails)

[object][object]

The user's emails.

[`data`](/docs/api/users/#user-update-request-identity-emails-data)

requiredstringrequired, string

User's email.

[`primary`](/docs/api/users/#user-update-request-identity-emails-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary email for the User.

[`phone_numbers`](/docs/api/users/#user-update-request-identity-phone-numbers)

[object][object]

The user's phone numbers, in E.164 format: +{countrycode}{number}. For example: "+14157452130". Phone numbers provided in other formats will be parsed on a best-effort basis. Phone number input is validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`data`](/docs/api/users/#user-update-request-identity-phone-numbers-data)

requiredstringrequired, string

User's phone number.

[`primary`](/docs/api/users/#user-update-request-identity-phone-numbers-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary phone number for the User.

[`addresses`](/docs/api/users/#user-update-request-identity-addresses)

[object][object]

The user's addresses.

[`street_1`](/docs/api/users/#user-update-request-identity-addresses-street-1)

stringstring

First line of street address.

[`street_2`](/docs/api/users/#user-update-request-identity-addresses-street-2)

stringstring

Second line of street address.

[`city`](/docs/api/users/#user-update-request-identity-addresses-city)

stringstring

City name.

[`region`](/docs/api/users/#user-update-request-identity-addresses-region)

stringstring

State, province or region.

[`country`](/docs/api/users/#user-update-request-identity-addresses-country)

requiredstringrequired, string

Country code.

[`postal_code`](/docs/api/users/#user-update-request-identity-addresses-postal-code)

stringstring

Postal or ZIP code.

[`primary`](/docs/api/users/#user-update-request-identity-addresses-primary)

requiredbooleanrequired, boolean

Indicates whether this is the primary address for the User.

[`id_numbers`](/docs/api/users/#user-update-request-identity-id-numbers)

[object][object]

The user's ID numbers.

[`value`](/docs/api/users/#user-update-request-identity-id-numbers-value)

requiredstringrequired, string

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/users/#user-update-request-identity-id-numbers-type)

requiredstringrequired, string

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`user_token`](/docs/api/users/#user-update-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`consumer_report_user_identity`](/docs/api/users/#user-update-request-consumer-report-user-identity)

objectobject

This field is only used by integrations created before December 10, 2025. All other integrations must use the `identity` object instead. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).
To create a Plaid Check Consumer Report for a user when using a `user_token`, this field must be present. If this field is not provided during user token creation, you can add it to the user later by calling `/user/update`. Once the field has been added to the user, you will be able to call `/link/token/create` with a non-empty `consumer_report_permissible_purpose` (which will automatically create a Plaid Check Consumer Report), or call `/cra/check_report/create` for that user.

[`first_name`](/docs/api/users/#user-update-request-consumer-report-user-identity-first-name)

requiredstringrequired, string

The user's first name

[`last_name`](/docs/api/users/#user-update-request-consumer-report-user-identity-last-name)

requiredstringrequired, string

The user's last name

[`phone_numbers`](/docs/api/users/#user-update-request-consumer-report-user-identity-phone-numbers)

required[string]required, [string]

The user's phone number, in E.164 format: +{countrycode}{number}. For example: "+14157452130". Phone numbers provided in other formats will be parsed on a best-effort basis. Phone number input is validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.

[`emails`](/docs/api/users/#user-update-request-consumer-report-user-identity-emails)

required[string]required, [string]

The user's emails

[`ssn_full`](/docs/api/users/#user-update-request-consumer-report-user-identity-ssn-full)

stringstring

The user's full social security number. This field should only be provided by lenders intending to share the resulting consumer report with a Government-Sponsored Enterprise (GSE), such as Fannie Mae or Freddie Mac.  
Format: "ddd-dd-dddd"

[`ssn_last_4`](/docs/api/users/#user-update-request-consumer-report-user-identity-ssn-last-4)

stringstring

The last 4 digits of the user's social security number.  
  

Max length: `4`

Min length: `4`

[`date_of_birth`](/docs/api/users/#user-update-request-consumer-report-user-identity-date-of-birth)

requiredstringrequired, string

To be provided in the format "yyyy-mm-dd".
This field is required for all Plaid Check customers.  
  

Format: `date`

[`primary_address`](/docs/api/users/#user-update-request-consumer-report-user-identity-primary-address)

requiredobjectrequired, object

Data about the components comprising an address.

[`city`](/docs/api/users/#user-update-request-consumer-report-user-identity-primary-address-city)

requiredstringrequired, string

The full city name

[`region`](/docs/api/users/#user-update-request-consumer-report-user-identity-primary-address-region)

requiredstringrequired, string

The region or state. In API versions 2018-05-22 and earlier, this field is called `state`.
Example: `"NC"`

[`street`](/docs/api/users/#user-update-request-consumer-report-user-identity-primary-address-street)

requiredstringrequired, string

The full street address
Example: `"564 Main Street, APT 15"`

[`postal_code`](/docs/api/users/#user-update-request-consumer-report-user-identity-primary-address-postal-code)

requiredstringrequired, string

The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.

[`country`](/docs/api/users/#user-update-request-consumer-report-user-identity-primary-address-country)

requiredstringrequired, string

The ISO 3166-1 alpha-2 country code

/user/update

```
const request = {
  user_id: 'usr_9nSp2KuZ2x4JDw',
  identity: {
    name: {
      given_name: "Carmen",
      family_name: "Berzatto"
    },
    date_of_birth: "1987-01-31",
    emails: [
      { data: "carmy@example.com", primary: true },
      { data: "bear@example.com", primary: false }
    ],
    phone_numbers: [
      { data: "+13125551212", primary: true }
    ],
    addresses: [
      {
        street_1: "3200 W Armitage Ave",
        street_2: null,
        city: "Chicago",
        region: "IL",
        country: "US",
        postal_code: "60657",
        primary: true
      }
    ],
    id_numbers: [
      {
        value: "1234",
        type: "us_ssn_last_4"
      }
    ]
  }
};

try {
  const response = await client.userUpdate(request);
} catch (error) {
  // handle error
}
```

/user/update

**Response fields**

[`request_id`](/docs/api/users/#user-update-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "Aim3b"
}
```

=\*=\*=\*=

#### `/user/remove`

#### Remove user

[`/user/remove`](/docs/api/users/#userremove) deletes a `user_id` or `user_token` and and associated information, including any Items associated with the user.

/user/remove

**Request fields**

[`client_id`](/docs/api/users/#user-remove-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/users/#user-remove-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_id`](/docs/api/users/#user-remove-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

[`user_token`](/docs/api/users/#user-remove-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

/user/remove

```
const request = {
  user_id: 'usr_9nSp2KuZ2x4JDw'
};

try {
  const response = await client.userRemove(request);
} catch (error) {
  // handle error
}
```

/user/remove

**Response fields**

[`request_id`](/docs/api/users/#user-remove-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "Aim3b"
}
```

=\*=\*=\*=

#### `/user/items/get`

#### Get Items associated with a User

Returns Items associated with a `user_id`, along with their corresponding statuses. Plaid associates an Item with a User when it has been successfully connected within a Link session initialized with that `user_id`.

/user/items/get

**Request fields**

[`client_id`](/docs/api/users/#user-items-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/users/#user-items-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user_token`](/docs/api/users/#user-items-get-request-user-token)

stringstring

The user token associated with the User data is being requested for. This field is used only by customers with pre-existing integrations that already use the `user_token` field. All other customers should use the `user_id` instead. For more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis).

[`user_id`](/docs/api/users/#user-items-get-request-user-id)

stringstring

A unique user identifier, created by `/user/create`. Integrations that began using `/user/create` after December 10, 2025 use this field to identify a user instead of the `user_token`. For more details, see [new user APIs](https://plaid.com/docs/api/users/user-apis).

/user/items/get

```
const request = {
  user_id: 'usr_9nSp2KuZ2x4JDw'
};

try {
  const response = await client.userItemsGet(request);
} catch (error) {
  // handle error
}
```

/user/items/get

**Response fields**

[`items`](/docs/api/users/#user-items-get-response-items)

[object][object]

[`item_id`](/docs/api/users/#user-items-get-response-items-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`institution_id`](/docs/api/users/#user-items-get-response-items-institution-id)

nullablestringnullable, string

The Plaid Institution ID associated with the Item. Field is `null` for Items created without an institution connection, such as Items created via Same Day Micro-deposits.

[`institution_name`](/docs/api/users/#user-items-get-response-items-institution-name)

nullablestringnullable, string

The name of the institution associated with the Item. Field is `null` for Items created without an institution connection, such as Items created via Same Day Micro-deposits.

[`webhook`](/docs/api/users/#user-items-get-response-items-webhook)

nullablestringnullable, string

The URL registered to receive webhooks for the Item.

[`auth_method`](/docs/api/users/#user-items-get-response-items-auth-method)

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

[`error`](/docs/api/users/#user-items-get-response-items-error)

nullableobjectnullable, object

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling `/item/get` to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

[`error_type`](/docs/api/users/#user-items-get-response-items-error-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/users/#user-items-get-response-items-error-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/users/#user-items-get-response-items-error-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/users/#user-items-get-response-items-error-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/users/#user-items-get-response-items-error-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/users/#user-items-get-response-items-error-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/users/#user-items-get-response-items-error-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/users/#user-items-get-response-items-error-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/users/#user-items-get-response-items-error-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/users/#user-items-get-response-items-error-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/users/#user-items-get-response-items-error-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/users/#user-items-get-response-items-error-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`available_products`](/docs/api/users/#user-items-get-response-items-available-products)

[string][string]

A list of products available for the Item that have not yet been accessed. The contents of this array will be mutually exclusive with `billed_products`.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`billed_products`](/docs/api/users/#user-items-get-response-items-billed-products)

[string][string]

A list of products that have been billed for the Item. The contents of this array will be mutually exclusive with `available_products`. Note - `billed_products` is populated in all environments but only requests in Production are billed. Also note that products that are billed on a pay-per-call basis rather than a pay-per-Item basis, such as `balance`, will not appear here.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`products`](/docs/api/users/#user-items-get-response-items-products)

[string][string]

A list of products added to the Item. In almost all cases, this will be the same as the `billed_products` field. For some products, it is possible for the product to be added to an Item but not yet billed (e.g. Assets, before `/asset_report/create` has been called, or Auth or Identity when added as Optional Products but before their endpoints have been called), in which case the product may appear in `products` but not in `billed_products`.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `payment_initiation`, `identity_verification`, `transactions`, `credit_details`, `income`, `income_verification`, `standing_orders`, `transfer`, `employment`, `recurring_transactions`, `transactions_refresh`, `signal`, `statements`, `processor_payments`, `processor_identity`, `profile`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`, `cra_network_insights`, `cra_cashflow_insights`, `cra_monitoring`, `cra_lend_score`, `cra_plaid_credit_score`, `layer`, `pay_by_bank`, `protect_linked_bank`

[`consented_products`](/docs/api/users/#user-items-get-response-items-consented-products)

[string][string]

A list of products that the user has consented to for the Item via [Data Transparency Messaging](https://plaid.com/docs/link/data-transparency-messaging-migration-guide). This will consist of all products where both of the following are true: the user has consented to the required data scopes for that product and you have Production access for that product.  
  

Possible values: `assets`, `auth`, `balance`, `balance_plus`, `beacon`, `identity`, `identity_match`, `investments`, `investments_auth`, `liabilities`, `transactions`, `income`, `income_verification`, `transfer`, `employment`, `recurring_transactions`, `signal`, `statements`, `processor_payments`, `processor_identity`, `cra_base_report`, `cra_income_insights`, `cra_lend_score`, `cra_partner_insights`, `cra_cashflow_insights`, `cra_monitoring`, `layer`

[`consent_expiration_time`](/docs/api/users/#user-items-get-response-items-consent-expiration-time)

nullablestringnullable, string

The date and time at which the Item's access consent will expire, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. If the Item does not have consent expiration scheduled, this field will be `null`. Currently, only institutions in Europe and a small number of institutions in the US have expiring consent. For a list of US institutions that currently expire consent, see the [OAuth Guide](https://plaid.com/docs/link/oauth/#refreshing-item-consent).  
  

Format: `date-time`

[`update_type`](/docs/api/users/#user-items-get-response-items-update-type)

stringstring

Indicates whether an Item requires user interaction to be updated, which can be the case for Items with some forms of two-factor authentication.  
`background` - Item can be updated in the background  
`user_present_required` - Item requires user interaction to be updated  
  

Possible values: `background`, `user_present_required`

[`request_id`](/docs/api/users/#user-items-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "items": [
    {
      "available_products": [
        "balance",
        "auth"
      ],
      "billed_products": [
        "identity",
        "transactions"
      ],
      "error": null,
      "institution_id": "ins_109508",
      "institution_name": "First Platypus Bank",
      "item_id": "Ed6bjNrDLJfGvZWwnkQlfxwoNz54B5C97ejBr",
      "update_type": "background",
      "webhook": "https://plaid.com/example/hook",
      "consent_expiration_time": null
    },
    {
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
      "error": null,
      "institution_id": "ins_109508",
      "institution_name": "First Platypus Bank",
      "item_id": "DWVAAPWq4RHGlEaNyGKRTAnPLaEmo8Cvq7na6",
      "update_type": "background",
      "webhook": "https://plaid.com/example/hook",
      "consent_expiration_time": null
    }
  ],
  "request_id": "m8MDnv9okwxFNBV"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
