---
title: "API - Beacon (beta) | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/beacon/"
scraped_at: "2026-03-07T22:04:01+00:00"
---

# Beacon

#### API reference for Beacon endpoints and webhooks

Add and report users on the Plaid Beacon network.

| Endpoints |  |
| --- | --- |
| [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate) | Create and scan a Beacon User against a Beacon Program |
| [`/beacon/user/get`](/docs/api/products/beacon/#beaconuserget) | Fetch a Beacon User |
| [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate) | Update and rescan a Beacon User |
| [`/beacon/user/account_insights/get`](/docs/api/products/beacon/#beaconuseraccount_insightsget) | Fetch a Beacon User's account insights |
| [`/beacon/user/history/list`](/docs/api/products/beacon/#beaconuserhistorylist) | List a Beacon User's history |
| [`/beacon/report/create`](/docs/api/products/beacon/#beaconreportcreate) | Create a fraud report for a given Beacon User |
| [`/beacon/report/get`](/docs/api/products/beacon/#beaconreportget) | Fetch a Beacon Report |
| [`/beacon/report/list`](/docs/api/products/beacon/#beaconreportlist) | List Beacon Reports for a given Beacon User |
| [`/beacon/report_syndication/get`](/docs/api/products/beacon/#beaconreport_syndicationget) | Fetch a Beacon Report Syndication |
| [`/beacon/report_syndication/list`](/docs/api/products/beacon/#beaconreport_syndicationlist) | List Beacon Report Syndications for a given Beacon User |
| [`/beacon/duplicate/get`](/docs/api/products/beacon/#beaconduplicateget) | Fetch a Beacon duplicate |

| See also |  |
| --- | --- |
| [`/dashboard_user/get`](/docs/api/kyc-aml-users/#dashboard_userget) | Retrieve information about a dashboard user |
| [`/dashboard_user/list`](/docs/api/kyc-aml-users/#dashboard_userlist) | List dashboard users |

| Webhooks |  |
| --- | --- |
| [`USER_STATUS_UPDATED`](/docs/api/products/beacon/#user_status_updated) | Beacon User status has changed |
| [`REPORT_CREATED`](/docs/api/products/beacon/#report_created) | Beacon Report has been created |
| [`REPORT_UPDATED`](/docs/api/products/beacon/#report_updated) | Beacon Report has been updated |
| [`REPORT_SYNDICATION_CREATED`](/docs/api/products/beacon/#report_syndication_created) | New Network Report matches one of your Users |
| [`DUPLICATE_DETECTED`](/docs/api/products/beacon/#duplicate_detected) | Duplicate Beacon User was created |

=\*=\*=\*=

#### `/beacon/user/create`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Create a Beacon User

Create and scan a Beacon User against your Beacon Program, according to your program's settings.

When you submit a new user to [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate), several checks are performed immediately:

- The user's PII (provided within the `user` object) is searched against all other users within the Beacon Program you specified. If a match is found that violates your program's "Duplicate Information Filtering" settings, the user will be returned with a status of `pending_review`.

- The user's PII is also searched against all fraud reports created by your organization across all of your Beacon Programs. If the user's data matches a fraud report that your team created, the user will be returned with a status of `rejected`.

- Finally, the user's PII is searched against all fraud report shared with the Beacon Network by other companies. If a matching fraud report is found, the user will be returned with a `pending_review` status if your program has enabled automatic flagging based on network fraud.

/beacon/user/create

**Request fields**

[`program_id`](/docs/api/products/beacon/#beacon-user-create-request-program-id)

requiredstringrequired, string

ID of the associated Beacon Program.

[`client_user_id`](/docs/api/products/beacon/#beacon-user-create-request-client-user-id)

requiredstringrequired, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`user`](/docs/api/products/beacon/#beacon-user-create-request-user)

requiredobjectrequired, object

A Beacon User's data which is used to check against duplicate records and the Beacon Fraud Network.  
In order to create a Beacon User, in addition to the `name`, *either* the `date_of_birth` *or* the `depository_accounts` field must be provided.

[`date_of_birth`](/docs/api/products/beacon/#beacon-user-create-request-user-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`name`](/docs/api/products/beacon/#beacon-user-create-request-user-name)

requiredobjectrequired, object

The full name for a given Beacon User.

[`given_name`](/docs/api/products/beacon/#beacon-user-create-request-user-name-given-name)

requiredstringrequired, string

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/beacon/#beacon-user-create-request-user-name-family-name)

requiredstringrequired, string

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/beacon/#beacon-user-create-request-user-address)

objectobject

Home address for the associated user. For more context on this field, see [Input Validation by Country](https://plaid.com/docs/identity-verification/hybrid-input-validation/#input-validation-by-country).

[`street`](/docs/api/products/beacon/#beacon-user-create-request-user-address-street)

requiredstringrequired, string

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/beacon/#beacon-user-create-request-user-address-street2)

stringstring

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/beacon/#beacon-user-create-request-user-address-city)

requiredstringrequired, string

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/beacon/#beacon-user-create-request-user-address-region)

stringstring

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/beacon/#beacon-user-create-request-user-address-postal-code)

stringstring

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/beacon/#beacon-user-create-request-user-address-country)

requiredstringrequired, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`email_address`](/docs/api/products/beacon/#beacon-user-create-request-user-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`phone_number`](/docs/api/products/beacon/#beacon-user-create-request-user-phone-number)

stringstring

A phone number in E.164 format.

[`id_number`](/docs/api/products/beacon/#beacon-user-create-request-user-id-number)

objectobject

The ID number associated with a Beacon User.

[`value`](/docs/api/products/beacon/#beacon-user-create-request-user-id-number-value)

requiredstringrequired, string

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/beacon/#beacon-user-create-request-user-id-number-type)

requiredstringrequired, string

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`ip_address`](/docs/api/products/beacon/#beacon-user-create-request-user-ip-address)

stringstring

An IPv4 or IPV6 address.

[`depository_accounts`](/docs/api/products/beacon/#beacon-user-create-request-user-depository-accounts)

[object][object]

Provide a list of bank accounts that are associated with this Beacon User. These accounts will be scanned across the Beacon Network and used to find duplicate records.  
Note: These accounts will not have Bank Account Insights. To receive Bank Account Insights please supply `access_tokens`.

[`account_number`](/docs/api/products/beacon/#beacon-user-create-request-user-depository-accounts-account-number)

requiredstringrequired, string

Must be a valid US Bank Account Number

[`routing_number`](/docs/api/products/beacon/#beacon-user-create-request-user-depository-accounts-routing-number)

requiredstringrequired, string

The routing number of the account.

[`access_tokens`](/docs/api/products/beacon/#beacon-user-create-request-access-tokens)

[string][string]

Send this array of access tokens to link accounts to the Beacon User and have them evaluated for Account Insights.
A maximum of 50 accounts total can be added to a single Beacon User.

[`client_id`](/docs/api/products/beacon/#beacon-user-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-user-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/user/create

```
const request: BeaconUserCreateRequest = {
  program_id: 'becprg_11111111111111',
  client_user_id: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
  access_tokens: [access_token],
  user: {
    email_address: 'user@example.com',
    date_of_birth: '1975-01-18',
    name: {
      given_name: 'Leslie',
      family_name: 'Knope',
    },
    address: {
      street: '123 Main St.',
      street2: 'Unit 42',
      city: 'Pawnee',
      region: 'IN',
      postal_code: '46001',
      country: 'US',
    },
  },
};

try {
  const response = await plaidClient.beaconUserCreate(request);
  console.log(response.status.value);
} catch (error) {
  // handle error
}
```

/beacon/user/create

**Response fields**

[`item_ids`](/docs/api/products/beacon/#beacon-user-create-response-item-ids)

[string][string]

An array of Plaid Item IDs corresponding to the Accounts associated with this Beacon User.

[`id`](/docs/api/products/beacon/#beacon-user-create-response-id)

stringstring

ID of the associated Beacon User.

[`version`](/docs/api/products/beacon/#beacon-user-create-response-version)

integerinteger

The `version` field begins with 1 and increments each time the user is updated.

[`created_at`](/docs/api/products/beacon/#beacon-user-create-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`updated_at`](/docs/api/products/beacon/#beacon-user-create-response-updated-at)

stringstring

An ISO8601 formatted timestamp. This field indicates the last time the resource was modified.  
  

Format: `date-time`

[`status`](/docs/api/products/beacon/#beacon-user-create-response-status)

stringstring

A status of a Beacon User.  
`rejected`: The Beacon User has been rejected for fraud. Users can be automatically or manually rejected.  
`pending_review`: The Beacon User has been marked for review.  
`cleared`: The Beacon User has been cleared of fraud.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`program_id`](/docs/api/products/beacon/#beacon-user-create-response-program-id)

stringstring

ID of the associated Beacon Program.

[`client_user_id`](/docs/api/products/beacon/#beacon-user-create-response-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`user`](/docs/api/products/beacon/#beacon-user-create-response-user)

objectobject

A Beacon User's data and resulting analysis when checked against duplicate records and the Beacon Fraud Network.

[`date_of_birth`](/docs/api/products/beacon/#beacon-user-create-response-user-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`name`](/docs/api/products/beacon/#beacon-user-create-response-user-name)

objectobject

The full name for a given Beacon User.

[`given_name`](/docs/api/products/beacon/#beacon-user-create-response-user-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/beacon/#beacon-user-create-response-user-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/beacon/#beacon-user-create-response-user-address)

objectobject

Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include a postal code

[`street`](/docs/api/products/beacon/#beacon-user-create-response-user-address-street)

stringstring

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/beacon/#beacon-user-create-response-user-address-street2)

nullablestringnullable, string

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/beacon/#beacon-user-create-response-user-address-city)

stringstring

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/beacon/#beacon-user-create-response-user-address-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/beacon/#beacon-user-create-response-user-address-postal-code)

nullablestringnullable, string

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/beacon/#beacon-user-create-response-user-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`email_address`](/docs/api/products/beacon/#beacon-user-create-response-user-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`phone_number`](/docs/api/products/beacon/#beacon-user-create-response-user-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`id_number`](/docs/api/products/beacon/#beacon-user-create-response-user-id-number)

nullableobjectnullable, object

The ID number associated with a Beacon User.

[`value`](/docs/api/products/beacon/#beacon-user-create-response-user-id-number-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/beacon/#beacon-user-create-response-user-id-number-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`ip_address`](/docs/api/products/beacon/#beacon-user-create-response-user-ip-address)

nullablestringnullable, string

An IPv4 or IPV6 address.

[`depository_accounts`](/docs/api/products/beacon/#beacon-user-create-response-user-depository-accounts)

[object][object]

[`account_mask`](/docs/api/products/beacon/#beacon-user-create-response-user-depository-accounts-account-mask)

stringstring

The last 2-4 numeric characters of this account’s account number.

[`routing_number`](/docs/api/products/beacon/#beacon-user-create-response-user-depository-accounts-routing-number)

stringstring

The routing number of the account.

[`added_at`](/docs/api/products/beacon/#beacon-user-create-response-user-depository-accounts-added-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`audit_trail`](/docs/api/products/beacon/#beacon-user-create-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/beacon/#beacon-user-create-response-audit-trail-source)

stringstring

A type indicating what caused a resource to be changed or updated.  
`dashboard` - The resource was created or updated by a member of your team via the Plaid dashboard.  
`api` - The resource was created or updated via the Plaid API.  
`system` - The resource was created or updated automatically by a part of the Plaid Beacon system. For example, if another business using Plaid Beacon created a fraud report that matched one of your users, your matching user's status would automatically be updated and the audit trail source would be `system`.  
`bulk_import` - The resource was created or updated as part of a bulk import process. For example, if your company provided a CSV of user data as part of your initial onboarding, the audit trail source would be `bulk_import`.  
  

Possible values: `dashboard`, `api`, `system`, `bulk_import`

[`dashboard_user_id`](/docs/api/products/beacon/#beacon-user-create-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/beacon/#beacon-user-create-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/beacon/#beacon-user-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "item_ids": [
    "515cd85321d3649aecddc015"
  ],
  "id": "becusr_42cF1MNo42r9Xj",
  "version": 1,
  "created_at": "2020-07-24T03:26:02Z",
  "updated_at": "2020-07-24T03:26:02Z",
  "status": "cleared",
  "program_id": "becprg_11111111111111",
  "client_user_id": "your-db-id-3b24110",
  "user": {
    "date_of_birth": "1990-05-29",
    "name": {
      "given_name": "Leslie",
      "family_name": "Knope"
    },
    "address": {
      "street": "123 Main St.",
      "street2": "Unit 42",
      "city": "Pawnee",
      "region": "IN",
      "postal_code": "46001",
      "country": "US"
    },
    "email_address": "user@example.com",
    "phone_number": "+19876543212",
    "id_number": {
      "value": "123456789",
      "type": "us_ssn"
    },
    "ip_address": "192.0.2.42",
    "depository_accounts": [
      {
        "account_mask": "4000",
        "routing_number": "021000021",
        "added_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/user/get`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Get a Beacon User

Fetch a Beacon User.

The Beacon User is returned with all of their associated information and a `status` based on the Beacon Network duplicate record and fraud checks.

/beacon/user/get

**Request fields**

[`beacon_user_id`](/docs/api/products/beacon/#beacon-user-get-request-beacon-user-id)

requiredstringrequired, string

ID of the associated Beacon User.

[`client_id`](/docs/api/products/beacon/#beacon-user-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-user-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/user/get

```
const request: BeaconUserGetRequest = {
  beacon_user_id: 'becusr_11111111111111',
};

try {
  const response = await plaidClient.beaconUserGet(request);
} catch (error) {
  // handle error
}
```

/beacon/user/get

**Response fields**

[`item_ids`](/docs/api/products/beacon/#beacon-user-get-response-item-ids)

[string][string]

An array of Plaid Item IDs corresponding to the Accounts associated with this Beacon User.

[`id`](/docs/api/products/beacon/#beacon-user-get-response-id)

stringstring

ID of the associated Beacon User.

[`version`](/docs/api/products/beacon/#beacon-user-get-response-version)

integerinteger

The `version` field begins with 1 and increments each time the user is updated.

[`created_at`](/docs/api/products/beacon/#beacon-user-get-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`updated_at`](/docs/api/products/beacon/#beacon-user-get-response-updated-at)

stringstring

An ISO8601 formatted timestamp. This field indicates the last time the resource was modified.  
  

Format: `date-time`

[`status`](/docs/api/products/beacon/#beacon-user-get-response-status)

stringstring

A status of a Beacon User.  
`rejected`: The Beacon User has been rejected for fraud. Users can be automatically or manually rejected.  
`pending_review`: The Beacon User has been marked for review.  
`cleared`: The Beacon User has been cleared of fraud.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`program_id`](/docs/api/products/beacon/#beacon-user-get-response-program-id)

stringstring

ID of the associated Beacon Program.

[`client_user_id`](/docs/api/products/beacon/#beacon-user-get-response-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`user`](/docs/api/products/beacon/#beacon-user-get-response-user)

objectobject

A Beacon User's data and resulting analysis when checked against duplicate records and the Beacon Fraud Network.

[`date_of_birth`](/docs/api/products/beacon/#beacon-user-get-response-user-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`name`](/docs/api/products/beacon/#beacon-user-get-response-user-name)

objectobject

The full name for a given Beacon User.

[`given_name`](/docs/api/products/beacon/#beacon-user-get-response-user-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/beacon/#beacon-user-get-response-user-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/beacon/#beacon-user-get-response-user-address)

objectobject

Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include a postal code

[`street`](/docs/api/products/beacon/#beacon-user-get-response-user-address-street)

stringstring

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/beacon/#beacon-user-get-response-user-address-street2)

nullablestringnullable, string

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/beacon/#beacon-user-get-response-user-address-city)

stringstring

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/beacon/#beacon-user-get-response-user-address-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/beacon/#beacon-user-get-response-user-address-postal-code)

nullablestringnullable, string

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/beacon/#beacon-user-get-response-user-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`email_address`](/docs/api/products/beacon/#beacon-user-get-response-user-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`phone_number`](/docs/api/products/beacon/#beacon-user-get-response-user-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`id_number`](/docs/api/products/beacon/#beacon-user-get-response-user-id-number)

nullableobjectnullable, object

The ID number associated with a Beacon User.

[`value`](/docs/api/products/beacon/#beacon-user-get-response-user-id-number-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/beacon/#beacon-user-get-response-user-id-number-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`ip_address`](/docs/api/products/beacon/#beacon-user-get-response-user-ip-address)

nullablestringnullable, string

An IPv4 or IPV6 address.

[`depository_accounts`](/docs/api/products/beacon/#beacon-user-get-response-user-depository-accounts)

[object][object]

[`account_mask`](/docs/api/products/beacon/#beacon-user-get-response-user-depository-accounts-account-mask)

stringstring

The last 2-4 numeric characters of this account’s account number.

[`routing_number`](/docs/api/products/beacon/#beacon-user-get-response-user-depository-accounts-routing-number)

stringstring

The routing number of the account.

[`added_at`](/docs/api/products/beacon/#beacon-user-get-response-user-depository-accounts-added-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`audit_trail`](/docs/api/products/beacon/#beacon-user-get-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/beacon/#beacon-user-get-response-audit-trail-source)

stringstring

A type indicating what caused a resource to be changed or updated.  
`dashboard` - The resource was created or updated by a member of your team via the Plaid dashboard.  
`api` - The resource was created or updated via the Plaid API.  
`system` - The resource was created or updated automatically by a part of the Plaid Beacon system. For example, if another business using Plaid Beacon created a fraud report that matched one of your users, your matching user's status would automatically be updated and the audit trail source would be `system`.  
`bulk_import` - The resource was created or updated as part of a bulk import process. For example, if your company provided a CSV of user data as part of your initial onboarding, the audit trail source would be `bulk_import`.  
  

Possible values: `dashboard`, `api`, `system`, `bulk_import`

[`dashboard_user_id`](/docs/api/products/beacon/#beacon-user-get-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/beacon/#beacon-user-get-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/beacon/#beacon-user-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "item_ids": [
    "515cd85321d3649aecddc015"
  ],
  "id": "becusr_42cF1MNo42r9Xj",
  "version": 1,
  "created_at": "2020-07-24T03:26:02Z",
  "updated_at": "2020-07-24T03:26:02Z",
  "status": "cleared",
  "program_id": "becprg_11111111111111",
  "client_user_id": "your-db-id-3b24110",
  "user": {
    "date_of_birth": "1990-05-29",
    "name": {
      "given_name": "Leslie",
      "family_name": "Knope"
    },
    "address": {
      "street": "123 Main St.",
      "street2": "Unit 42",
      "city": "Pawnee",
      "region": "IN",
      "postal_code": "46001",
      "country": "US"
    },
    "email_address": "user@example.com",
    "phone_number": "+19876543212",
    "id_number": {
      "value": "123456789",
      "type": "us_ssn"
    },
    "ip_address": "192.0.2.42",
    "depository_accounts": [
      {
        "account_mask": "4000",
        "routing_number": "021000021",
        "added_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/user/update`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Update the identity data of a Beacon User

Update the identity data for a Beacon User in your Beacon Program or add new accounts to the Beacon User.

Similar to [`/beacon/user/create`](/docs/api/products/beacon/#beaconusercreate), several checks are performed immediately when you submit an identity data change to [`/beacon/user/update`](/docs/api/products/beacon/#beaconuserupdate):

- The user's updated PII is searched against all other users within the Beacon Program you specified. If a match is found that violates your program's "Duplicate Information Filtering" settings, the user will be returned with a status of `pending_review`.

- The user's updated PII is also searched against all fraud reports created by your organization across all of your Beacon Programs. If the user's data matches a fraud report that your team created, the user will be returned with a status of `rejected`.

- Finally, the user's PII is searched against all fraud report shared with the Beacon Network by other companies. If a matching fraud report is found, the user will be returned with a `pending_review` status if your program has enabled automatic flagging based on network fraud.

Plaid maintains a version history for each Beacon User, so the Beacon User's identity data before and after the update is retained as separate versions.

/beacon/user/update

**Request fields**

[`beacon_user_id`](/docs/api/products/beacon/#beacon-user-update-request-beacon-user-id)

requiredstringrequired, string

ID of the associated Beacon User.

[`user`](/docs/api/products/beacon/#beacon-user-update-request-user)

objectobject

A subset of a Beacon User's data which is used to patch the existing identity data associated with a Beacon User. At least one field must be provided. If left unset or null, user data will not be patched.

[`date_of_birth`](/docs/api/products/beacon/#beacon-user-update-request-user-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`name`](/docs/api/products/beacon/#beacon-user-update-request-user-name)

objectobject

The full name for a given Beacon User.

[`given_name`](/docs/api/products/beacon/#beacon-user-update-request-user-name-given-name)

requiredstringrequired, string

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/beacon/#beacon-user-update-request-user-name-family-name)

requiredstringrequired, string

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/beacon/#beacon-user-update-request-user-address)

objectobject

Home address for the associated user. For more context on this field, see [Input Validation by Country](https://plaid.com/docs/identity-verification/hybrid-input-validation/#input-validation-by-country).

[`street`](/docs/api/products/beacon/#beacon-user-update-request-user-address-street)

requiredstringrequired, string

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/beacon/#beacon-user-update-request-user-address-street2)

stringstring

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/beacon/#beacon-user-update-request-user-address-city)

requiredstringrequired, string

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/beacon/#beacon-user-update-request-user-address-region)

stringstring

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/beacon/#beacon-user-update-request-user-address-postal-code)

stringstring

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/beacon/#beacon-user-update-request-user-address-country)

requiredstringrequired, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`email_address`](/docs/api/products/beacon/#beacon-user-update-request-user-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`phone_number`](/docs/api/products/beacon/#beacon-user-update-request-user-phone-number)

stringstring

A phone number in E.164 format.

[`id_number`](/docs/api/products/beacon/#beacon-user-update-request-user-id-number)

objectobject

The ID number associated with a Beacon User.

[`value`](/docs/api/products/beacon/#beacon-user-update-request-user-id-number-value)

requiredstringrequired, string

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/beacon/#beacon-user-update-request-user-id-number-type)

requiredstringrequired, string

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`ip_address`](/docs/api/products/beacon/#beacon-user-update-request-user-ip-address)

stringstring

An IPv4 or IPV6 address.

[`depository_accounts`](/docs/api/products/beacon/#beacon-user-update-request-user-depository-accounts)

[object][object]

[`account_number`](/docs/api/products/beacon/#beacon-user-update-request-user-depository-accounts-account-number)

requiredstringrequired, string

Must be a valid US Bank Account Number

[`routing_number`](/docs/api/products/beacon/#beacon-user-update-request-user-depository-accounts-routing-number)

requiredstringrequired, string

The routing number of the account.

[`access_tokens`](/docs/api/products/beacon/#beacon-user-update-request-access-tokens)

[string][string]

Send this array of access tokens to add accounts to this user for evaluation.
This will add accounts to this Beacon User. If left null only existing accounts will be returned in response.
A maximum of 50 accounts total can be added to a Beacon User.

[`client_id`](/docs/api/products/beacon/#beacon-user-update-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-user-update-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/user/update

```
const request: BeaconUserUpdateRequest = {
  beacon_user_id: 'becusr_11111111111111',
  user: {
    email_address: 'user@example.com',
    date_of_birth: '1975-01-18',
    name: {
      given_name: 'Leslie',
      family_name: 'Knope',
    },
    address: {
      street: '123 Main St.',
      street2: 'Unit 42',
      city: 'Pawnee',
      region: 'IN',
      postal_code: '46001',
      country: 'US',
    },
  },
};

try {
  const response = await plaidClient.beaconUserUpdate(request);
  console.log(response.status.value);
} catch (error) {
  // handle error
}
```

/beacon/user/update

**Response fields**

[`item_ids`](/docs/api/products/beacon/#beacon-user-update-response-item-ids)

[string][string]

An array of Plaid Item IDs corresponding to the Accounts associated with this Beacon User.

[`id`](/docs/api/products/beacon/#beacon-user-update-response-id)

stringstring

ID of the associated Beacon User.

[`version`](/docs/api/products/beacon/#beacon-user-update-response-version)

integerinteger

The `version` field begins with 1 and increments each time the user is updated.

[`created_at`](/docs/api/products/beacon/#beacon-user-update-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`updated_at`](/docs/api/products/beacon/#beacon-user-update-response-updated-at)

stringstring

An ISO8601 formatted timestamp. This field indicates the last time the resource was modified.  
  

Format: `date-time`

[`status`](/docs/api/products/beacon/#beacon-user-update-response-status)

stringstring

A status of a Beacon User.  
`rejected`: The Beacon User has been rejected for fraud. Users can be automatically or manually rejected.  
`pending_review`: The Beacon User has been marked for review.  
`cleared`: The Beacon User has been cleared of fraud.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`program_id`](/docs/api/products/beacon/#beacon-user-update-response-program-id)

stringstring

ID of the associated Beacon Program.

[`client_user_id`](/docs/api/products/beacon/#beacon-user-update-response-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`user`](/docs/api/products/beacon/#beacon-user-update-response-user)

objectobject

A Beacon User's data and resulting analysis when checked against duplicate records and the Beacon Fraud Network.

[`date_of_birth`](/docs/api/products/beacon/#beacon-user-update-response-user-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`name`](/docs/api/products/beacon/#beacon-user-update-response-user-name)

objectobject

The full name for a given Beacon User.

[`given_name`](/docs/api/products/beacon/#beacon-user-update-response-user-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/beacon/#beacon-user-update-response-user-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/beacon/#beacon-user-update-response-user-address)

objectobject

Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include a postal code

[`street`](/docs/api/products/beacon/#beacon-user-update-response-user-address-street)

stringstring

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/beacon/#beacon-user-update-response-user-address-street2)

nullablestringnullable, string

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/beacon/#beacon-user-update-response-user-address-city)

stringstring

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/beacon/#beacon-user-update-response-user-address-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/beacon/#beacon-user-update-response-user-address-postal-code)

nullablestringnullable, string

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/beacon/#beacon-user-update-response-user-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`email_address`](/docs/api/products/beacon/#beacon-user-update-response-user-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`phone_number`](/docs/api/products/beacon/#beacon-user-update-response-user-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`id_number`](/docs/api/products/beacon/#beacon-user-update-response-user-id-number)

nullableobjectnullable, object

The ID number associated with a Beacon User.

[`value`](/docs/api/products/beacon/#beacon-user-update-response-user-id-number-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/beacon/#beacon-user-update-response-user-id-number-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`ip_address`](/docs/api/products/beacon/#beacon-user-update-response-user-ip-address)

nullablestringnullable, string

An IPv4 or IPV6 address.

[`depository_accounts`](/docs/api/products/beacon/#beacon-user-update-response-user-depository-accounts)

[object][object]

[`account_mask`](/docs/api/products/beacon/#beacon-user-update-response-user-depository-accounts-account-mask)

stringstring

The last 2-4 numeric characters of this account’s account number.

[`routing_number`](/docs/api/products/beacon/#beacon-user-update-response-user-depository-accounts-routing-number)

stringstring

The routing number of the account.

[`added_at`](/docs/api/products/beacon/#beacon-user-update-response-user-depository-accounts-added-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`audit_trail`](/docs/api/products/beacon/#beacon-user-update-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/beacon/#beacon-user-update-response-audit-trail-source)

stringstring

A type indicating what caused a resource to be changed or updated.  
`dashboard` - The resource was created or updated by a member of your team via the Plaid dashboard.  
`api` - The resource was created or updated via the Plaid API.  
`system` - The resource was created or updated automatically by a part of the Plaid Beacon system. For example, if another business using Plaid Beacon created a fraud report that matched one of your users, your matching user's status would automatically be updated and the audit trail source would be `system`.  
`bulk_import` - The resource was created or updated as part of a bulk import process. For example, if your company provided a CSV of user data as part of your initial onboarding, the audit trail source would be `bulk_import`.  
  

Possible values: `dashboard`, `api`, `system`, `bulk_import`

[`dashboard_user_id`](/docs/api/products/beacon/#beacon-user-update-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/beacon/#beacon-user-update-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/beacon/#beacon-user-update-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "item_ids": [
    "515cd85321d3649aecddc015"
  ],
  "id": "becusr_42cF1MNo42r9Xj",
  "version": 1,
  "created_at": "2020-07-24T03:26:02Z",
  "updated_at": "2020-07-24T03:26:02Z",
  "status": "cleared",
  "program_id": "becprg_11111111111111",
  "client_user_id": "your-db-id-3b24110",
  "user": {
    "date_of_birth": "1990-05-29",
    "name": {
      "given_name": "Leslie",
      "family_name": "Knope"
    },
    "address": {
      "street": "123 Main St.",
      "street2": "Unit 42",
      "city": "Pawnee",
      "region": "IN",
      "postal_code": "46001",
      "country": "US"
    },
    "email_address": "user@example.com",
    "phone_number": "+19876543212",
    "id_number": {
      "value": "123456789",
      "type": "us_ssn"
    },
    "ip_address": "192.0.2.42",
    "depository_accounts": [
      {
        "account_mask": "4000",
        "routing_number": "021000021",
        "added_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/user/account_insights/get`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Get Account Insights for a Beacon User

Get Account Insights for all Accounts linked to this Beacon User. The insights for each account are computed based on the information that was last retrieved from the financial institution.

/beacon/user/account\_insights/get

**Request fields**

[`beacon_user_id`](/docs/api/products/beacon/#beacon-user-account_insights-get-request-beacon-user-id)

requiredstringrequired, string

ID of the associated Beacon User.

[`access_token`](/docs/api/products/beacon/#beacon-user-account_insights-get-request-access-token)

requiredstringrequired, string

The access token associated with the Item data is being requested for.

[`client_id`](/docs/api/products/beacon/#beacon-user-account_insights-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-user-account_insights-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/user/account\_insights/get

```
const request: BeaconUserAccountInsightsGetRequest = {
  beacon_user_id: 'becusr_11111111111111',
  access_token: 'access-sandbox-12345678',
  client_id: process.env.PLAID_CLIENT_ID,
  secret: process.env.PLAID_SECRET,
};

try {
  const response = await plaidClient.beaconUserAccountInsightsGet(request);
  console.log(response.status.value);
} catch (error) {
  // handle error
}
```

/beacon/user/account\_insights/get

**Response fields**

[`beacon_user_id`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-beacon-user-id)

stringstring

ID of the associated Beacon User.

[`created_at`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`updated_at`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-updated-at)

stringstring

An ISO8601 formatted timestamp. This field indicates the last time the resource was modified.  
  

Format: `date-time`

[`bank_account_insights`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights)

objectobject

A collection of Bank Accounts linked to an Item that is associated with this Beacon User.

[`item_id`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-item-id)

stringstring

The Plaid Item ID the Bank Accounts belong to.

[`accounts`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts)

[object][object]

[`account_id`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-account-id)

stringstring

The Plaid `account_id`

[`type`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

[`attributes`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes)

objectobject

The attributes object contains data that can be used to assess account risk. Examples of data include:
`days_since_first_plaid_connection`: The number of days since the first time the Item was connected to an application via Plaid
`plaid_connections_count_7d`: The number of times the Item has been connected to applications via Plaid over the past 7 days
`plaid_connections_count_30d`: The number of times the Item has been connected to applications via Plaid over the past 30 days
`total_plaid_connections_count`: The number of times the Item has been connected to applications via Plaid
For the full list and detailed documentation of core attributes available, or to request that core attributes not be returned, contact Sales or your Plaid account manager

[`days_since_first_plaid_connection`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-days-since-first-plaid-connection)

nullableintegernullable, integer

The number of days since the first time the Item was connected to an application via Plaid

[`is_account_closed`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-is-account-closed)

nullablebooleannullable, boolean

Indicates if the account has been closed by the financial institution or the consumer, or is at risk of being closed

[`is_account_frozen_or_restricted`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-is-account-frozen-or-restricted)

nullablebooleannullable, boolean

Indicates whether the account has withdrawals and transfers disabled or if access to the account is restricted. This could be due to a freeze by the credit issuer, legal restrictions (e.g., sanctions), or regulatory requirements limiting monthly withdrawals, among other reasons

[`total_plaid_connections_count`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-total-plaid-connections-count)

nullableintegernullable, integer

The total number of times the item has been connected to applications via Plaid

[`plaid_connections_count_7d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-plaid-connections-count-7d)

nullableintegernullable, integer

The number of times the Item has been connected to applications via Plaid over the past 7 days

[`plaid_connections_count_30d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-plaid-connections-count-30d)

nullableintegernullable, integer

The number of times the Item has been connected to applications via Plaid over the past 30 days

[`failed_plaid_non_oauth_authentication_attempts_count_3d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-failed-plaid-non-oauth-authentication-attempts-count-3d)

nullableintegernullable, integer

The number of failed non-OAuth authentication attempts via Plaid for this bank account over the past 3 days

[`plaid_non_oauth_authentication_attempts_count_3d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-plaid-non-oauth-authentication-attempts-count-3d)

nullableintegernullable, integer

The number of non-OAuth authentication attempts via Plaid for this bank account over the past 3 days

[`failed_plaid_non_oauth_authentication_attempts_count_7d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-failed-plaid-non-oauth-authentication-attempts-count-7d)

nullableintegernullable, integer

The number of failed non-OAuth authentication attempts via Plaid for this bank account over the past 7 days

[`plaid_non_oauth_authentication_attempts_count_7d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-plaid-non-oauth-authentication-attempts-count-7d)

nullableintegernullable, integer

The number of non-OAuth authentication attempts via Plaid for this bank account over the past 7 days

[`failed_plaid_non_oauth_authentication_attempts_count_30d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-failed-plaid-non-oauth-authentication-attempts-count-30d)

nullableintegernullable, integer

The number of failed non-OAuth authentication attempts via Plaid for this bank account over the past 30 days

[`plaid_non_oauth_authentication_attempts_count_30d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-plaid-non-oauth-authentication-attempts-count-30d)

nullableintegernullable, integer

The number of non-OAuth authentication attempts via Plaid for this bank account over the past 30 days

[`distinct_ip_addresses_count_3d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-distinct-ip-addresses-count-3d)

nullableintegernullable, integer

The number of distinct IP addresses linked to the same bank account during Plaid authentication in the last 3 days

[`distinct_ip_addresses_count_7d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-distinct-ip-addresses-count-7d)

nullableintegernullable, integer

The number of distinct IP addresses linked to the same bank account during Plaid authentication in the last 7 days

[`distinct_ip_addresses_count_30d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-distinct-ip-addresses-count-30d)

nullableintegernullable, integer

The number of distinct IP addresses linked to the same bank account during Plaid authentication in the last 30 days

[`distinct_ip_addresses_count_90d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-distinct-ip-addresses-count-90d)

nullableintegernullable, integer

The number of distinct IP addresses linked to the same bank account during Plaid authentication in the last 90 days

[`distinct_user_agents_count_3d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-distinct-user-agents-count-3d)

nullableintegernullable, integer

The number of distinct user agents linked to the same bank account during Plaid authentication in the last 3 days

[`distinct_user_agents_count_7d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-distinct-user-agents-count-7d)

nullableintegernullable, integer

The number of distinct user agents linked to the same bank account during Plaid authentication in the last 7 days

[`distinct_user_agents_count_30d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-distinct-user-agents-count-30d)

nullableintegernullable, integer

The number of distinct user agents linked to the same bank account during Plaid authentication in the last 30 days

[`distinct_user_agents_count_90d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-distinct-user-agents-count-90d)

nullableintegernullable, integer

The number of distinct user agents linked to the same bank account during Plaid authentication in the last 90 days

[`address_change_count_28d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-address-change-count-28d)

nullableintegernullable, integer

The number of times the account's addresses on file have changed over the past 28 days

[`email_change_count_28d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-email-change-count-28d)

nullableintegernullable, integer

The number of times the account's email addresses on file have changed over the past 28 days

[`phone_change_count_28d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-phone-change-count-28d)

nullableintegernullable, integer

The number of times the account's phone numbers on file have changed over the past 28 days

[`address_change_count_90d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-address-change-count-90d)

nullableintegernullable, integer

The number of times the account's addresses on file have changed over the past 90 days

[`email_change_count_90d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-email-change-count-90d)

nullableintegernullable, integer

The number of times the account's email addresses on file have changed over the past 90 days

[`phone_change_count_90d`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-phone-change-count-90d)

nullableintegernullable, integer

The number of times the account's phone numbers on file have changed over the past 90 days

[`days_since_account_opening`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-days-since-account-opening)

nullableintegernullable, integer

The number of days since the bank account was opened, as reported by the financial institution

[`days_since_first_observed_transaction`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-bank-account-insights-accounts-attributes-days-since-first-observed-transaction)

nullableintegernullable, integer

The number of days since the oldest transaction available to Plaid for this account. This measure, combined with Plaid connection history, can be used to infer the age of the account

[`request_id`](/docs/api/products/beacon/#beacon-user-account_insights-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "beacon_user_id": "becusr_42cF1MNo42r9Xj",
  "created_at": "2020-07-24T03:26:02Z",
  "updated_at": "2020-07-24T03:26:02Z",
  "bank_account_insights": {
    "item_id": "515cd85321d3649aecddc015",
    "accounts": [
      {
        "account_id": "blgvvBlXw3cq5GMPwqB6s6q4dLKB9WcVqGDGo",
        "type": "depository",
        "subtype": "checking",
        "attributes": {
          "days_since_first_plaid_connection": 1,
          "is_account_closed": false,
          "is_account_frozen_or_restricted": false,
          "total_plaid_connections_count": 1,
          "plaid_connections_count_7d": 1,
          "plaid_connections_count_30d": 1,
          "failed_plaid_non_oauth_authentication_attempts_count_3d": 1,
          "plaid_non_oauth_authentication_attempts_count_3d": 1,
          "failed_plaid_non_oauth_authentication_attempts_count_7d": 1,
          "plaid_non_oauth_authentication_attempts_count_7d": 1,
          "failed_plaid_non_oauth_authentication_attempts_count_30d": 1,
          "plaid_non_oauth_authentication_attempts_count_30d": 1,
          "distinct_ip_addresses_count_3d": 1,
          "distinct_ip_addresses_count_7d": 1,
          "distinct_ip_addresses_count_30d": 1,
          "distinct_ip_addresses_count_90d": 1,
          "distinct_user_agents_count_3d": 1,
          "distinct_user_agents_count_7d": 1,
          "distinct_user_agents_count_30d": 1,
          "distinct_user_agents_count_90d": 1,
          "address_change_count_28d": 1,
          "email_change_count_28d": 2,
          "phone_change_count_28d": 1,
          "address_change_count_90d": 3,
          "email_change_count_90d": 4,
          "phone_change_count_90d": 2,
          "days_since_account_opening": 365,
          "days_since_first_observed_transaction": 180
        }
      }
    ]
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/user/history/list`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### List a Beacon User's history

List all changes to the Beacon User in reverse-chronological order.

/beacon/user/history/list

**Request fields**

[`beacon_user_id`](/docs/api/products/beacon/#beacon-user-history-list-request-beacon-user-id)

requiredstringrequired, string

ID of the associated Beacon User.

[`cursor`](/docs/api/products/beacon/#beacon-user-history-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

[`client_id`](/docs/api/products/beacon/#beacon-user-history-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-user-history-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/user/history/list

```
const request: BeaconUserHistoryListRequest = {
  beacon_user_id: 'becusr_11111111111111',
};

try {
  const response = await plaidClient.beaconUserHistoryList(request);
} catch (error) {
  // handle error
}
```

/beacon/user/history/list

**Response fields**

[`beacon_users`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users)

[object][object]

[`item_ids`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-item-ids)

[string][string]

An array of Plaid Item IDs corresponding to the Accounts associated with this Beacon User.

[`id`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-id)

stringstring

ID of the associated Beacon User.

[`version`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-version)

integerinteger

The `version` field begins with 1 and increments each time the user is updated.

[`created_at`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`updated_at`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-updated-at)

stringstring

An ISO8601 formatted timestamp. This field indicates the last time the resource was modified.  
  

Format: `date-time`

[`status`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-status)

stringstring

A status of a Beacon User.  
`rejected`: The Beacon User has been rejected for fraud. Users can be automatically or manually rejected.  
`pending_review`: The Beacon User has been marked for review.  
`cleared`: The Beacon User has been cleared of fraud.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`program_id`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-program-id)

stringstring

ID of the associated Beacon Program.

[`client_user_id`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`user`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user)

objectobject

A Beacon User's data and resulting analysis when checked against duplicate records and the Beacon Fraud Network.

[`date_of_birth`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`name`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-name)

objectobject

The full name for a given Beacon User.

[`given_name`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-address)

objectobject

Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include a postal code

[`street`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-address-street)

stringstring

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-address-street2)

nullablestringnullable, string

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-address-city)

stringstring

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-address-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-address-postal-code)

nullablestringnullable, string

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`email_address`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`phone_number`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`id_number`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-id-number)

nullableobjectnullable, object

The ID number associated with a Beacon User.

[`value`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-id-number-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-id-number-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`ip_address`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-ip-address)

nullablestringnullable, string

An IPv4 or IPV6 address.

[`depository_accounts`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-depository-accounts)

[object][object]

[`account_mask`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-depository-accounts-account-mask)

stringstring

The last 2-4 numeric characters of this account’s account number.

[`routing_number`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-depository-accounts-routing-number)

stringstring

The routing number of the account.

[`added_at`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-user-depository-accounts-added-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`audit_trail`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-audit-trail-source)

stringstring

A type indicating what caused a resource to be changed or updated.  
`dashboard` - The resource was created or updated by a member of your team via the Plaid dashboard.  
`api` - The resource was created or updated via the Plaid API.  
`system` - The resource was created or updated automatically by a part of the Plaid Beacon system. For example, if another business using Plaid Beacon created a fraud report that matched one of your users, your matching user's status would automatically be updated and the audit trail source would be `system`.  
`bulk_import` - The resource was created or updated as part of a bulk import process. For example, if your company provided a CSV of user data as part of your initial onboarding, the audit trail source would be `bulk_import`.  
  

Possible values: `dashboard`, `api`, `system`, `bulk_import`

[`dashboard_user_id`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/beacon/#beacon-user-history-list-response-beacon-users-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`next_cursor`](/docs/api/products/beacon/#beacon-user-history-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/beacon/#beacon-user-history-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "beacon_users": [
    {
      "item_ids": [
        "515cd85321d3649aecddc015"
      ],
      "id": "becusr_42cF1MNo42r9Xj",
      "version": 1,
      "created_at": "2020-07-24T03:26:02Z",
      "updated_at": "2020-07-24T03:26:02Z",
      "status": "cleared",
      "program_id": "becprg_11111111111111",
      "client_user_id": "your-db-id-3b24110",
      "user": {
        "date_of_birth": "1990-05-29",
        "name": {
          "given_name": "Leslie",
          "family_name": "Knope"
        },
        "address": {
          "street": "123 Main St.",
          "street2": "Unit 42",
          "city": "Pawnee",
          "region": "IN",
          "postal_code": "46001",
          "country": "US"
        },
        "email_address": "user@example.com",
        "phone_number": "+19876543212",
        "id_number": {
          "value": "123456789",
          "type": "us_ssn"
        },
        "ip_address": "192.0.2.42",
        "depository_accounts": [
          {
            "account_mask": "4000",
            "routing_number": "021000021",
            "added_at": "2020-07-24T03:26:02Z"
          }
        ]
      },
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/report/create`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Create a Beacon Report

Create a fraud report for a given Beacon User.

/beacon/report/create

**Request fields**

[`beacon_user_id`](/docs/api/products/beacon/#beacon-report-create-request-beacon-user-id)

requiredstringrequired, string

ID of the associated Beacon User.

[`type`](/docs/api/products/beacon/#beacon-report-create-request-type)

requiredstringrequired, string

The type of Beacon Report.  
`first_party`: If this is the same individual as the one who submitted the KYC.  
`stolen`: If this is a different individual from the one who submitted the KYC.  
`synthetic`: If this is an individual using fabricated information.  
`account_takeover`: If this individual's account was compromised.  
`unknown`: If you aren't sure who committed the fraud.  
  

Possible values: `first_party`, `stolen`, `synthetic`, `account_takeover`, `unknown`

[`fraud_date`](/docs/api/products/beacon/#beacon-report-create-request-fraud-date)

requiredstringrequired, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`fraud_amount`](/docs/api/products/beacon/#beacon-report-create-request-fraud-amount)

objectobject

The amount and currency of the fraud or attempted fraud.
`fraud_amount` should be omitted to indicate an unknown fraud amount.

[`iso_currency_code`](/docs/api/products/beacon/#beacon-report-create-request-fraud-amount-iso-currency-code)

requiredstringrequired, string

An ISO-4217 currency code.  
  

Possible values: `USD`

[`value`](/docs/api/products/beacon/#beacon-report-create-request-fraud-amount-value)

requirednumberrequired, number

The amount value.
This value can be 0 to indicate no money was lost.
Must not contain more than two digits of precision (e.g., `1.23`).  
  

Format: `double`

[`client_id`](/docs/api/products/beacon/#beacon-report-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-report-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/report/create

```
const request: BeaconReportCreateRequest = {
  beacon_user_id: 'becusr_11111111111111',
  type: 'first_party',
  fraud_date: '1975-01-18',
};

try {
  const response = await plaidClient.beaconReportCreate(request);
} catch (error) {
  // handle error
}
```

/beacon/report/create

**Response fields**

[`id`](/docs/api/products/beacon/#beacon-report-create-response-id)

stringstring

ID of the associated Beacon Report.

[`beacon_user_id`](/docs/api/products/beacon/#beacon-report-create-response-beacon-user-id)

stringstring

ID of the associated Beacon User.

[`created_at`](/docs/api/products/beacon/#beacon-report-create-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`type`](/docs/api/products/beacon/#beacon-report-create-response-type)

stringstring

The type of Beacon Report.  
`first_party`: If this is the same individual as the one who submitted the KYC.  
`stolen`: If this is a different individual from the one who submitted the KYC.  
`synthetic`: If this is an individual using fabricated information.  
`account_takeover`: If this individual's account was compromised.  
`data_breach`: If this individual's data was compromised in a breach.  
`unknown`: If you aren't sure who committed the fraud.  
  

Possible values: `first_party`, `stolen`, `synthetic`, `account_takeover`, `data_breach`, `unknown`

[`fraud_date`](/docs/api/products/beacon/#beacon-report-create-response-fraud-date)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`event_date`](/docs/api/products/beacon/#beacon-report-create-response-event-date)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`fraud_amount`](/docs/api/products/beacon/#beacon-report-create-response-fraud-amount)

nullableobjectnullable, object

The amount and currency of the fraud or attempted fraud.
`fraud_amount` should be omitted to indicate an unknown fraud amount.

[`iso_currency_code`](/docs/api/products/beacon/#beacon-report-create-response-fraud-amount-iso-currency-code)

stringstring

An ISO-4217 currency code.  
  

Possible values: `USD`

[`value`](/docs/api/products/beacon/#beacon-report-create-response-fraud-amount-value)

numbernumber

The amount value.
This value can be 0 to indicate no money was lost.
Must not contain more than two digits of precision (e.g., `1.23`).  
  

Format: `double`

[`audit_trail`](/docs/api/products/beacon/#beacon-report-create-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/beacon/#beacon-report-create-response-audit-trail-source)

stringstring

A type indicating what caused a resource to be changed or updated.  
`dashboard` - The resource was created or updated by a member of your team via the Plaid dashboard.  
`api` - The resource was created or updated via the Plaid API.  
`system` - The resource was created or updated automatically by a part of the Plaid Beacon system. For example, if another business using Plaid Beacon created a fraud report that matched one of your users, your matching user's status would automatically be updated and the audit trail source would be `system`.  
`bulk_import` - The resource was created or updated as part of a bulk import process. For example, if your company provided a CSV of user data as part of your initial onboarding, the audit trail source would be `bulk_import`.  
  

Possible values: `dashboard`, `api`, `system`, `bulk_import`

[`dashboard_user_id`](/docs/api/products/beacon/#beacon-report-create-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/beacon/#beacon-report-create-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/beacon/#beacon-report-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "becrpt_11111111111111",
  "beacon_user_id": "becusr_42cF1MNo42r9Xj",
  "created_at": "2020-07-24T03:26:02Z",
  "type": "first_party",
  "fraud_date": "1990-05-29",
  "event_date": "1990-05-29",
  "fraud_amount": {
    "iso_currency_code": "USD",
    "value": 100
  },
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/report/get`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Get a Beacon Report

Returns a Beacon report for a given Beacon report id.

/beacon/report/get

**Request fields**

[`beacon_report_id`](/docs/api/products/beacon/#beacon-report-get-request-beacon-report-id)

requiredstringrequired, string

ID of the associated Beacon Report.

[`client_id`](/docs/api/products/beacon/#beacon-report-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-report-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/report/get

```
const request: BeaconReportGetRequest = {
  beacon_report_id: 'becrpt_11111111111111',
};

try {
  const response = await plaidClient.beaconReportGet(request);
} catch (error) {
  // handle error
}
```

/beacon/report/get

**Response fields**

[`id`](/docs/api/products/beacon/#beacon-report-get-response-id)

stringstring

ID of the associated Beacon Report.

[`beacon_user_id`](/docs/api/products/beacon/#beacon-report-get-response-beacon-user-id)

stringstring

ID of the associated Beacon User.

[`created_at`](/docs/api/products/beacon/#beacon-report-get-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`type`](/docs/api/products/beacon/#beacon-report-get-response-type)

stringstring

The type of Beacon Report.  
`first_party`: If this is the same individual as the one who submitted the KYC.  
`stolen`: If this is a different individual from the one who submitted the KYC.  
`synthetic`: If this is an individual using fabricated information.  
`account_takeover`: If this individual's account was compromised.  
`data_breach`: If this individual's data was compromised in a breach.  
`unknown`: If you aren't sure who committed the fraud.  
  

Possible values: `first_party`, `stolen`, `synthetic`, `account_takeover`, `data_breach`, `unknown`

[`fraud_date`](/docs/api/products/beacon/#beacon-report-get-response-fraud-date)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`event_date`](/docs/api/products/beacon/#beacon-report-get-response-event-date)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`fraud_amount`](/docs/api/products/beacon/#beacon-report-get-response-fraud-amount)

nullableobjectnullable, object

The amount and currency of the fraud or attempted fraud.
`fraud_amount` should be omitted to indicate an unknown fraud amount.

[`iso_currency_code`](/docs/api/products/beacon/#beacon-report-get-response-fraud-amount-iso-currency-code)

stringstring

An ISO-4217 currency code.  
  

Possible values: `USD`

[`value`](/docs/api/products/beacon/#beacon-report-get-response-fraud-amount-value)

numbernumber

The amount value.
This value can be 0 to indicate no money was lost.
Must not contain more than two digits of precision (e.g., `1.23`).  
  

Format: `double`

[`audit_trail`](/docs/api/products/beacon/#beacon-report-get-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/beacon/#beacon-report-get-response-audit-trail-source)

stringstring

A type indicating what caused a resource to be changed or updated.  
`dashboard` - The resource was created or updated by a member of your team via the Plaid dashboard.  
`api` - The resource was created or updated via the Plaid API.  
`system` - The resource was created or updated automatically by a part of the Plaid Beacon system. For example, if another business using Plaid Beacon created a fraud report that matched one of your users, your matching user's status would automatically be updated and the audit trail source would be `system`.  
`bulk_import` - The resource was created or updated as part of a bulk import process. For example, if your company provided a CSV of user data as part of your initial onboarding, the audit trail source would be `bulk_import`.  
  

Possible values: `dashboard`, `api`, `system`, `bulk_import`

[`dashboard_user_id`](/docs/api/products/beacon/#beacon-report-get-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/beacon/#beacon-report-get-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/beacon/#beacon-report-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "becrpt_11111111111111",
  "beacon_user_id": "becusr_42cF1MNo42r9Xj",
  "created_at": "2020-07-24T03:26:02Z",
  "type": "first_party",
  "fraud_date": "1990-05-29",
  "event_date": "1990-05-29",
  "fraud_amount": {
    "iso_currency_code": "USD",
    "value": 100
  },
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/report/list`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### List Beacon Reports for a Beacon User

Use the [`/beacon/report/list`](/docs/api/products/beacon/#beaconreportlist) endpoint to view all Beacon Reports you created for a specific Beacon User. The reports returned by this endpoint are exclusively reports you created for a specific user. A Beacon User can only have one active report at a time, but a new report can be created if a previous report has been deleted. The results from this endpoint are paginated; the `next_cursor` field will be populated if there is another page of results that can be retrieved. To fetch the next page, pass the `next_cursor` value as the `cursor` parameter in the next request.

/beacon/report/list

**Request fields**

[`beacon_user_id`](/docs/api/products/beacon/#beacon-report-list-request-beacon-user-id)

requiredstringrequired, string

ID of the associated Beacon User.

[`cursor`](/docs/api/products/beacon/#beacon-report-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

[`client_id`](/docs/api/products/beacon/#beacon-report-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-report-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/report/list

```
const request: BeaconReportListRequest = {
  beacon_user_id: 'becusr_11111111111111',
};

try {
  const response = await plaidClient.beaconReportList(request);
} catch (error) {
  // handle error
}
```

/beacon/report/list

**Response fields**

[`beacon_reports`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports)

[object][object]

[`id`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-id)

stringstring

ID of the associated Beacon Report.

[`beacon_user_id`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-beacon-user-id)

stringstring

ID of the associated Beacon User.

[`created_at`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`type`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-type)

stringstring

The type of Beacon Report.  
`first_party`: If this is the same individual as the one who submitted the KYC.  
`stolen`: If this is a different individual from the one who submitted the KYC.  
`synthetic`: If this is an individual using fabricated information.  
`account_takeover`: If this individual's account was compromised.  
`data_breach`: If this individual's data was compromised in a breach.  
`unknown`: If you aren't sure who committed the fraud.  
  

Possible values: `first_party`, `stolen`, `synthetic`, `account_takeover`, `data_breach`, `unknown`

[`fraud_date`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-fraud-date)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`event_date`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-event-date)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`fraud_amount`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-fraud-amount)

nullableobjectnullable, object

The amount and currency of the fraud or attempted fraud.
`fraud_amount` should be omitted to indicate an unknown fraud amount.

[`iso_currency_code`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-fraud-amount-iso-currency-code)

stringstring

An ISO-4217 currency code.  
  

Possible values: `USD`

[`value`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-fraud-amount-value)

numbernumber

The amount value.
This value can be 0 to indicate no money was lost.
Must not contain more than two digits of precision (e.g., `1.23`).  
  

Format: `double`

[`audit_trail`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-audit-trail-source)

stringstring

A type indicating what caused a resource to be changed or updated.  
`dashboard` - The resource was created or updated by a member of your team via the Plaid dashboard.  
`api` - The resource was created or updated via the Plaid API.  
`system` - The resource was created or updated automatically by a part of the Plaid Beacon system. For example, if another business using Plaid Beacon created a fraud report that matched one of your users, your matching user's status would automatically be updated and the audit trail source would be `system`.  
`bulk_import` - The resource was created or updated as part of a bulk import process. For example, if your company provided a CSV of user data as part of your initial onboarding, the audit trail source would be `bulk_import`.  
  

Possible values: `dashboard`, `api`, `system`, `bulk_import`

[`dashboard_user_id`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/beacon/#beacon-report-list-response-beacon-reports-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`next_cursor`](/docs/api/products/beacon/#beacon-report-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/beacon/#beacon-report-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "beacon_reports": [
    {
      "id": "becrpt_11111111111111",
      "beacon_user_id": "becusr_42cF1MNo42r9Xj",
      "created_at": "2020-07-24T03:26:02Z",
      "type": "first_party",
      "fraud_date": "1990-05-29",
      "event_date": "1990-05-29",
      "fraud_amount": {
        "iso_currency_code": "USD",
        "value": 100
      },
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/report_syndication/get`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Get a Beacon Report Syndication

Returns a Beacon Report Syndication for a given Beacon Report Syndication id.

/beacon/report\_syndication/get

**Request fields**

[`beacon_report_syndication_id`](/docs/api/products/beacon/#beacon-report_syndication-get-request-beacon-report-syndication-id)

requiredstringrequired, string

ID of the associated Beacon Report Syndication.

[`client_id`](/docs/api/products/beacon/#beacon-report_syndication-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-report_syndication-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/report\_syndication/get

```
const request: BeaconReportSyndicationGetRequest = {
  beacon_report_syndication_id: 'becrsn_11111111111111',
};

try {
  const response = await plaidClient.beaconReportSyndicationGet(request);
} catch (error) {
  // handle error
}
```

/beacon/report\_syndication/get

**Response fields**

[`id`](/docs/api/products/beacon/#beacon-report_syndication-get-response-id)

stringstring

ID of the associated Beacon Report Syndication.

[`beacon_user_id`](/docs/api/products/beacon/#beacon-report_syndication-get-response-beacon-user-id)

stringstring

ID of the associated Beacon User.

[`report`](/docs/api/products/beacon/#beacon-report_syndication-get-response-report)

objectobject

A subset of information from a Beacon Report that has been syndicated to a matching Beacon User in your program.  
The `id` field in the response is the ID of the original report that was syndicated. If the original report was created by your organization, the field will be filled with the ID of the report. Otherwise, the field will be `null` indicating that the original report was created by another Beacon customer.

[`id`](/docs/api/products/beacon/#beacon-report_syndication-get-response-report-id)

nullablestringnullable, string

ID of the associated Beacon Report.

[`created_at`](/docs/api/products/beacon/#beacon-report_syndication-get-response-report-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`type`](/docs/api/products/beacon/#beacon-report_syndication-get-response-report-type)

stringstring

The type of Beacon Report.  
`first_party`: If this is the same individual as the one who submitted the KYC.  
`stolen`: If this is a different individual from the one who submitted the KYC.  
`synthetic`: If this is an individual using fabricated information.  
`account_takeover`: If this individual's account was compromised.  
`data_breach`: If this individual's data was compromised in a breach.  
`unknown`: If you aren't sure who committed the fraud.  
  

Possible values: `first_party`, `stolen`, `synthetic`, `account_takeover`, `data_breach`, `unknown`

[`fraud_date`](/docs/api/products/beacon/#beacon-report_syndication-get-response-report-fraud-date)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`event_date`](/docs/api/products/beacon/#beacon-report_syndication-get-response-report-event-date)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`analysis`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis)

objectobject

Analysis of which fields matched between the originally reported Beacon User and the Beacon User that the report was syndicated to.

[`address`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-date-of-birth)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`email_address`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-email-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`name`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-name)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`id_number`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-id-number)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`ip_address`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-ip-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`phone_number`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-phone-number)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`depository_accounts`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-depository-accounts)

[object][object]

[`account_mask`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-depository-accounts-account-mask)

stringstring

The last 2-4 numeric characters of this account’s account number.

[`routing_number`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-depository-accounts-routing-number)

stringstring

The routing number of the account.

[`match_status`](/docs/api/products/beacon/#beacon-report_syndication-get-response-analysis-depository-accounts-match-status)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`request_id`](/docs/api/products/beacon/#beacon-report_syndication-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "becrsn_11111111111111",
  "beacon_user_id": "becusr_42cF1MNo42r9Xj",
  "report": {
    "id": "becrpt_11111111111111",
    "created_at": "2020-07-24T03:26:02Z",
    "type": "first_party",
    "fraud_date": "1990-05-29",
    "event_date": "1990-05-29"
  },
  "analysis": {
    "address": "match",
    "date_of_birth": "match",
    "email_address": "match",
    "name": "match",
    "id_number": "match",
    "ip_address": "match",
    "phone_number": "match",
    "depository_accounts": [
      {
        "account_mask": "4000",
        "routing_number": "021000021",
        "match_status": "match"
      }
    ]
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/report_syndication/list`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### List Beacon Report Syndications for a Beacon User

Use the [`/beacon/report_syndication/list`](/docs/api/products/beacon/#beaconreport_syndicationlist) endpoint to view all Beacon Reports that have been syndicated to a specific Beacon User. This endpoint returns Beacon Report Syndications which are references to Beacon Reports created either by you, or another Beacon customer, that matched the specified Beacon User. A Beacon User can have multiple active Beacon Report Syndications at once. The results from this endpoint are paginated; the `next_cursor` field will be populated if there is another page of results that can be retrieved. To fetch the next page, pass the `next_cursor` value as the `cursor` parameter in the next request.

/beacon/report\_syndication/list

**Request fields**

[`beacon_user_id`](/docs/api/products/beacon/#beacon-report_syndication-list-request-beacon-user-id)

requiredstringrequired, string

ID of the associated Beacon User.

[`cursor`](/docs/api/products/beacon/#beacon-report_syndication-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

[`client_id`](/docs/api/products/beacon/#beacon-report_syndication-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-report_syndication-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/report\_syndication/list

```
const request: BeaconReportSyndicationListRequest = {
  beacon_user_id: 'becusr_11111111111111',
};

try {
  const response = await plaidClient.beaconReportSyndicationList(request);
} catch (error) {
  // handle error
}
```

/beacon/report\_syndication/list

**Response fields**

[`beacon_report_syndications`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications)

[object][object]

[`id`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-id)

stringstring

ID of the associated Beacon Report Syndication.

[`beacon_user_id`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-beacon-user-id)

stringstring

ID of the associated Beacon User.

[`report`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-report)

objectobject

A subset of information from a Beacon Report that has been syndicated to a matching Beacon User in your program.  
The `id` field in the response is the ID of the original report that was syndicated. If the original report was created by your organization, the field will be filled with the ID of the report. Otherwise, the field will be `null` indicating that the original report was created by another Beacon customer.

[`id`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-report-id)

nullablestringnullable, string

ID of the associated Beacon Report.

[`created_at`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-report-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`type`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-report-type)

stringstring

The type of Beacon Report.  
`first_party`: If this is the same individual as the one who submitted the KYC.  
`stolen`: If this is a different individual from the one who submitted the KYC.  
`synthetic`: If this is an individual using fabricated information.  
`account_takeover`: If this individual's account was compromised.  
`data_breach`: If this individual's data was compromised in a breach.  
`unknown`: If you aren't sure who committed the fraud.  
  

Possible values: `first_party`, `stolen`, `synthetic`, `account_takeover`, `data_breach`, `unknown`

[`fraud_date`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-report-fraud-date)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`event_date`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-report-event-date)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`analysis`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis)

objectobject

Analysis of which fields matched between the originally reported Beacon User and the Beacon User that the report was syndicated to.

[`address`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-date-of-birth)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`email_address`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-email-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`name`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-name)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`id_number`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-id-number)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`ip_address`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-ip-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`phone_number`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-phone-number)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`depository_accounts`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-depository-accounts)

[object][object]

[`account_mask`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-depository-accounts-account-mask)

stringstring

The last 2-4 numeric characters of this account’s account number.

[`routing_number`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-depository-accounts-routing-number)

stringstring

The routing number of the account.

[`match_status`](/docs/api/products/beacon/#beacon-report_syndication-list-response-beacon-report-syndications-analysis-depository-accounts-match-status)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`next_cursor`](/docs/api/products/beacon/#beacon-report_syndication-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/beacon/#beacon-report_syndication-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "beacon_report_syndications": [
    {
      "id": "becrsn_11111111111111",
      "beacon_user_id": "becusr_42cF1MNo42r9Xj",
      "report": {
        "id": "becrpt_11111111111111",
        "created_at": "2020-07-24T03:26:02Z",
        "type": "first_party",
        "fraud_date": "1990-05-29",
        "event_date": "1990-05-29"
      },
      "analysis": {
        "address": "match",
        "date_of_birth": "match",
        "email_address": "match",
        "name": "match",
        "id_number": "match",
        "ip_address": "match",
        "phone_number": "match",
        "depository_accounts": [
          {
            "account_mask": "4000",
            "routing_number": "021000021",
            "match_status": "match"
          }
        ]
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/beacon/duplicate/get`

This feature is in currently in beta; Your account must be enabled for this feature in order to test it in Sandbox. To enable this feature or check your status, contact your account manager or [submit a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

#### Get a Beacon Duplicate

Returns a Beacon Duplicate for a given Beacon Duplicate id.

A Beacon Duplicate represents a pair of similar Beacon Users within your organization.

Two Beacon User revisions are returned for each Duplicate record in either the `beacon_user1` or `beacon_user2` response fields.

The `analysis` field in the response indicates which fields matched between `beacon_user1` and `beacon_user2`.

/beacon/duplicate/get

**Request fields**

[`beacon_duplicate_id`](/docs/api/products/beacon/#beacon-duplicate-get-request-beacon-duplicate-id)

requiredstringrequired, string

ID of the associated Beacon Duplicate.

[`client_id`](/docs/api/products/beacon/#beacon-duplicate-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/beacon/#beacon-duplicate-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/beacon/duplicate/get

```
const request: BeaconDuplicateGetRequest = {
  beacon_duplicate_id: 'becdup_11111111111111',
};

try {
  const response = await plaidClient.beaconDuplicateGet(request);
} catch (error) {
  // handle error
}
```

/beacon/duplicate/get

**Response fields**

[`id`](/docs/api/products/beacon/#beacon-duplicate-get-response-id)

stringstring

ID of the associated Beacon Duplicate.

[`beacon_user1`](/docs/api/products/beacon/#beacon-duplicate-get-response-beacon-user1)

objectobject

A Beacon User Revision identifies a Beacon User at some point in its revision history.

[`id`](/docs/api/products/beacon/#beacon-duplicate-get-response-beacon-user1-id)

stringstring

ID of the associated Beacon User.

[`version`](/docs/api/products/beacon/#beacon-duplicate-get-response-beacon-user1-version)

integerinteger

The `version` field begins with 1 and increments with each subsequent revision.

[`beacon_user2`](/docs/api/products/beacon/#beacon-duplicate-get-response-beacon-user2)

objectobject

A Beacon User Revision identifies a Beacon User at some point in its revision history.

[`id`](/docs/api/products/beacon/#beacon-duplicate-get-response-beacon-user2-id)

stringstring

ID of the associated Beacon User.

[`version`](/docs/api/products/beacon/#beacon-duplicate-get-response-beacon-user2-version)

integerinteger

The `version` field begins with 1 and increments with each subsequent revision.

[`analysis`](/docs/api/products/beacon/#beacon-duplicate-get-response-analysis)

objectobject

Analysis of which fields matched between one Beacon User and another.

[`address`](/docs/api/products/beacon/#beacon-duplicate-get-response-analysis-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/beacon/#beacon-duplicate-get-response-analysis-date-of-birth)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`email_address`](/docs/api/products/beacon/#beacon-duplicate-get-response-analysis-email-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`name`](/docs/api/products/beacon/#beacon-duplicate-get-response-analysis-name)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`id_number`](/docs/api/products/beacon/#beacon-duplicate-get-response-analysis-id-number)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`ip_address`](/docs/api/products/beacon/#beacon-duplicate-get-response-analysis-ip-address)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`phone_number`](/docs/api/products/beacon/#beacon-duplicate-get-response-analysis-phone-number)

stringstring

An enum indicating the match type between two Beacon Users.  
`match` indicates that the provided input data was a strong match against the other Beacon User.  
`partial_match` indicates the data approximately matched the other Beacon User. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to compare this field against the other Beacon User and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to compare this field against the original Beacon User because the field was not present in one of the Beacon Users.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`request_id`](/docs/api/products/beacon/#beacon-duplicate-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "becdup_11111111111111",
  "beacon_user1": {
    "id": "becusr_42cF1MNo42r9Xj",
    "version": 1
  },
  "beacon_user2": {
    "id": "becusr_42cF1MNo42r9Xj",
    "version": 1
  },
  "analysis": {
    "address": "match",
    "date_of_birth": "match",
    "email_address": "match",
    "name": "match",
    "id_number": "match",
    "ip_address": "match",
    "phone_number": "match"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

### Webhooks

A user in `cleared` status may change to a `pending_review` status after a Beacon Report Syndication is received, which would trigger both a `USER_STATUS_UPDATED` as well as a `REPORT_SYNDICATION_CREATED` webhook. New Beacon Users may also be flagged as duplicates of another user, which triggers a `DUPLICATE_DETECTED` webhook. Beacon Reports created and managed by your account will trigger `REPORT_CREATED` and `REPORT_UPDATED` webhooks, and may also result in a `USER_STATUS_UPDATED` if the user status is changed from `cleared` to `rejected` at that time.

=\*=\*=\*=

#### `USER_STATUS_UPDATED`

Fired when a Beacon User status has changed, which can occur manually via the dashboard or when information is reported to the Beacon network.

**Properties**

[`webhook_type`](/docs/api/products/beacon/#BeaconUserStatusUpdatedWebhook-webhook-type)

stringstring

`BEACON`

[`webhook_code`](/docs/api/products/beacon/#BeaconUserStatusUpdatedWebhook-webhook-code)

stringstring

`USER_STATUS_UPDATED`

[`beacon_user_id`](/docs/api/products/beacon/#BeaconUserStatusUpdatedWebhook-beacon-user-id)

stringstring

The ID of the associated Beacon user.

[`environment`](/docs/api/products/beacon/#BeaconUserStatusUpdatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "BEACON",
  "webhook_code": "USER_STATUS_UPDATED",
  "beacon_user_id": "becusr_4WciCrtbxF76T8",
  "environment": "production"
}
```

=\*=\*=\*=

#### `REPORT_CREATED`

Fired when one of your Beacon Users is first reported to the Beacon network.

**Properties**

[`webhook_type`](/docs/api/products/beacon/#BeaconReportCreatedWebhook-webhook-type)

stringstring

`BEACON`

[`webhook_code`](/docs/api/products/beacon/#BeaconReportCreatedWebhook-webhook-code)

stringstring

`REPORT_CREATED`

[`beacon_report_id`](/docs/api/products/beacon/#BeaconReportCreatedWebhook-beacon-report-id)

stringstring

The ID of the associated Beacon Report.

[`environment`](/docs/api/products/beacon/#BeaconReportCreatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "BEACON",
  "webhook_code": "REPORT_CREATED",
  "beacon_report_id": "becrpt_2zugxV6hWQZG91",
  "environment": "production"
}
```

=\*=\*=\*=

#### `REPORT_UPDATED`

Fired when one of your existing Beacon Reports has been modified or removed from the Beacon Network.

**Properties**

[`webhook_type`](/docs/api/products/beacon/#BeaconReportUpdatedWebhook-webhook-type)

stringstring

`BEACON`

[`webhook_code`](/docs/api/products/beacon/#BeaconReportUpdatedWebhook-webhook-code)

stringstring

`REPORT_UPDATED`

[`beacon_report_id`](/docs/api/products/beacon/#BeaconReportUpdatedWebhook-beacon-report-id)

stringstring

The ID of the associated Beacon Report.

[`environment`](/docs/api/products/beacon/#BeaconReportUpdatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "BEACON",
  "webhook_code": "REPORT_UPDATED",
  "beacon_report_id": "becrpt_2zugxV6hWQZG91",
  "environment": "production"
}
```

=\*=\*=\*=

#### `REPORT_SYNDICATION_CREATED`

Fired when a report created on the Beacon Network matches with one of your Beacon Users.

**Properties**

[`webhook_type`](/docs/api/products/beacon/#BeaconReportSyndicationCreatedWebhook-webhook-type)

stringstring

`BEACON`

[`webhook_code`](/docs/api/products/beacon/#BeaconReportSyndicationCreatedWebhook-webhook-code)

stringstring

`REPORT_SYNDICATION_CREATED`

[`beacon_report_syndication_id`](/docs/api/products/beacon/#BeaconReportSyndicationCreatedWebhook-beacon-report-syndication-id)

stringstring

The ID of the associated Beacon Report Syndication.

[`environment`](/docs/api/products/beacon/#BeaconReportSyndicationCreatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "BEACON",
  "webhook_code": "REPORT_SYNDICATION_CREATED",
  "beacon_report_syndication_id": "becrsn_eZPgiiv3JH8rfT",
  "environment": "production"
}
```

=\*=\*=\*=

#### `DUPLICATE_DETECTED`

Fired when a Beacon User created within your organization matches one of your existing users.

**Properties**

[`webhook_type`](/docs/api/products/beacon/#BeaconDuplicateDetectedWebhook-webhook-type)

stringstring

`BEACON`

[`webhook_code`](/docs/api/products/beacon/#BeaconDuplicateDetectedWebhook-webhook-code)

stringstring

`DUPLICATE_DETECTED`

[`beacon_duplicate_id`](/docs/api/products/beacon/#BeaconDuplicateDetectedWebhook-beacon-duplicate-id)

stringstring

The ID of the associated Beacon Duplicate.

[`environment`](/docs/api/products/beacon/#BeaconDuplicateDetectedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "BEACON",
  "webhook_code": "DUPLICATE_DETECTED",
  "beacon_duplicate_id": "becdup_erJcFn97r9sugZ",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
