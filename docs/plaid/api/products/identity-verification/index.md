---
title: "API - Identity Verification | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/identity-verification/"
scraped_at: "2026-03-07T22:04:09+00:00"
---

# Identity Verification

#### API reference for Identity Verification endpoints and webhooks

For how-to guidance, see the [Identity Verification documentation](/docs/identity-verification/).

| Endpoints |  |
| --- | --- |
| [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) | Create a new identity verification |
| [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget) | Retrieve a previously created identity verification |
| [`/identity_verification/list`](/docs/api/products/identity-verification/#identity_verificationlist) | Filter and list identity verifications |
| [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry) | Allow a user to retry an identity verification |

| See also |  |
| --- | --- |
| [`/dashboard_user/get`](/docs/api/kyc-aml-users/#dashboard_userget) | Retrieve information about a dashboard user |
| [`/dashboard_user/list`](/docs/api/kyc-aml-users/#dashboard_userlist) | List dashboard users |

| Webhooks |  |
| --- | --- |
| [`STATUS_UPDATED`](/docs/api/products/identity-verification/#status_updated) | The status of an identity verification has been updated |
| [`STEP_UPDATED`](/docs/api/products/identity-verification/#step_updated) | A step in the identity verification process has been completed |
| [`RETRIED`](/docs/api/products/identity-verification/#retried) | An identity verification has been retried |

### Endpoints

=\*=\*=\*=

#### `/identity_verification/create`

#### Create a new Identity Verification

Create a new Identity Verification for the user specified by the `client_user_id` and/or `user_id` field. At least one of these two fields must be provided. The requirements and behavior of the verification are determined by the `template_id` provided. If `user_id` is provided, there must be an associated user otherwise an error will be returned.

If you don't know whether an active Identity Verification exists for a given `client_user_id` and/or `user_id`, you can specify `"is_idempotent": true` in the request body. With idempotency enabled, a new Identity Verification will only be created if one does not already exist for the associated `client_user_id` and/or `user_id`, and `template_id`. If an Identity Verification is found, it will be returned unmodified with a `200 OK` HTTP status code.

If `user_id` is not provided, you can also use this endpoint to supply information you already have collected about the user; if any of these fields are specified, the screens prompting the user to enter them will be skipped during the Link flow. If `user_id` is provided, user information can not be included in the request body. Please use the [`/user/update`](/docs/api/users/#userupdate) endpoint to update user data instead.

/identity\_verification/create

**Request fields**

[`client_user_id`](/docs/api/products/identity-verification/#identity_verification-create-request-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`user_id`](/docs/api/products/identity-verification/#identity_verification-create-request-user-id)

stringstring

Unique user identifier, created by calling `/user/create`. Either a `user_id` or the `client_user_id` must be provided. The `user_id` may only be used instead of the `client_user_id` if you were not a pre-existing user of `/user/create` as of December 10, 2025; for more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis). If both this field and a `client_user_id` are present in a request, the `user_id` must have been created from the provided `client_user_id`.

[`is_shareable`](/docs/api/products/identity-verification/#identity_verification-create-request-is-shareable)

requiredbooleanrequired, boolean

A flag specifying whether you would like Plaid to expose a shareable URL for the verification being created.

[`template_id`](/docs/api/products/identity-verification/#identity_verification-create-request-template-id)

requiredstringrequired, string

ID of the associated Identity Verification template. Like all Plaid identifiers, this is case-sensitive.

[`gave_consent`](/docs/api/products/identity-verification/#identity_verification-create-request-gave-consent)

requiredbooleanrequired, boolean

A flag specifying whether the end user has already agreed to a privacy policy specifying that their data will be shared with Plaid for verification purposes.  
If `gave_consent` is set to `true`, the `accept_tos` step will be marked as `skipped` and the end user's session will start at the next step requirement.  
  

Default: `false`

[`user`](/docs/api/products/identity-verification/#identity_verification-create-request-user)

objectobject

User information collected outside of Link, most likely via your own onboarding process.  
Each of the following identity fields are optional:  
`email_address`  
`phone_number`  
`date_of_birth`  
`name`  
`address`  
`id_number`  
Specifically, these fields are optional in that they can either be fully provided (satisfying every required field in their subschema) or omitted from the request entirely by not providing the key or value.
Providing these fields via the API will result in Link skipping the data collection process for the associated user. All verification steps enabled in the associated Identity Verification Template will still be run. Verification steps will either be run immediately, or once the user completes the `accept_tos` step, depending on the value provided to the `gave_consent` field.
If you are not using the shareable URL feature, you can optionally provide these fields via `/link/token/create` instead; both `/identity_verification/create` and `/link/token/create` are valid ways to provide this information. Note that if you provide a non-`null` user data object via `/identity_verification/create`, any user data fields entered via `/link/token/create` for the same `client_user_id` will be ignored when prefilling Link.

[`email_address`](/docs/api/products/identity-verification/#identity_verification-create-request-user-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-create-request-user-phone-number)

stringstring

A valid phone number in E.164 format.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-create-request-user-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`name`](/docs/api/products/identity-verification/#identity_verification-create-request-user-name)

objectobject

You can use this field to pre-populate the user's legal name; if it is provided here, they will not be prompted to enter their name in the identity verification attempt.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-create-request-user-name-given-name)

requiredstringrequired, string

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-create-request-user-name-family-name)

requiredstringrequired, string

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/identity-verification/#identity_verification-create-request-user-address)

objectobject

Home address for the user. Supported values are: not provided, address with only country code or full address.  
For more context on this field, see [Input Validation by Country](https://plaid.com/docs/identity-verification/hybrid-input-validation/#input-validation-by-country).

[`street`](/docs/api/products/identity-verification/#identity_verification-create-request-user-address-street)

stringstring

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/identity-verification/#identity_verification-create-request-user-address-street2)

stringstring

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/identity-verification/#identity_verification-create-request-user-address-city)

stringstring

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/identity-verification/#identity_verification-create-request-user-address-region)

stringstring

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-create-request-user-address-postal-code)

stringstring

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-create-request-user-address-country)

requiredstringrequired, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-create-request-user-id-number)

objectobject

ID number submitted by the user, currently used only for the Identity Verification product. If the user has not submitted this data yet, this field will be `null`. Otherwise, both fields are guaranteed to be filled.

[`value`](/docs/api/products/identity-verification/#identity_verification-create-request-user-id-number-value)

requiredstringrequired, string

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/identity-verification/#identity_verification-create-request-user-id-number-type)

requiredstringrequired, string

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`client_user_id`](/docs/api/products/identity-verification/#identity_verification-create-request-user-client-user-id)

deprecatedstringdeprecated, string

Specifying `user.client_user_id` is deprecated. Please provide `client_user_id` at the root level instead.

[`client_id`](/docs/api/products/identity-verification/#identity_verification-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/identity-verification/#identity_verification-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`is_idempotent`](/docs/api/products/identity-verification/#identity_verification-create-request-is-idempotent)

booleanboolean

An optional flag specifying how you would like Plaid to handle attempts to create an Identity Verification when an Identity Verification already exists for the provided `client_user_id` and `template_id`.
If idempotency is enabled, Plaid will return the existing Identity Verification. If idempotency is disabled, Plaid will reject the request with a `400 Bad Request` status code if an Identity Verification already exists for the supplied `client_user_id` and `template_id`.

/identity\_verification/create

```
const request: IdentityVerificationCreateRequest = {
  client_user_id: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
  is_shareable: true,
  template_id: 'idvtmp_52xR9LKo77r1Np',
  gave_consent: true,
  user: {
    email_address: 'acharleston@email.com',
    phone_number: '+12345678909',
    date_of_birth: '1975-01-18',
    name: {
      given_name: 'Anna',
      family_name: 'Charleston',
    },
    address: {
      street: '100 Market Street',
      street2: 'Apt 1A',
      city: 'San Francisco',
      region: 'CA',
      postal_code: '94103',
      country: 'US',
    },
    id_number: {
      value: '123456789',
      type: 'us_ssn',
    },
  },
};
try {
  const response = await client.identityVerificationCreate(request);
} catch (error) {
  // handle error
}
```

/identity\_verification/create

**Response fields**

[`id`](/docs/api/products/identity-verification/#identity_verification-create-response-id)

stringstring

ID of the associated Identity Verification attempt.

[`client_user_id`](/docs/api/products/identity-verification/#identity_verification-create-response-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`created_at`](/docs/api/products/identity-verification/#identity_verification-create-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`completed_at`](/docs/api/products/identity-verification/#identity_verification-create-response-completed-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`previous_attempt_id`](/docs/api/products/identity-verification/#identity_verification-create-response-previous-attempt-id)

nullablestringnullable, string

The ID for the Identity Verification preceding this session. This field will only be filled if the current Identity Verification is a retry of a previous attempt.

[`shareable_url`](/docs/api/products/identity-verification/#identity_verification-create-response-shareable-url)

nullablestringnullable, string

A shareable URL that can be sent directly to the user to complete verification

[`template`](/docs/api/products/identity-verification/#identity_verification-create-response-template)

objectobject

The resource ID and version number of the template configuring the behavior of a given Identity Verification.

[`id`](/docs/api/products/identity-verification/#identity_verification-create-response-template-id)

stringstring

ID of the associated Identity Verification template. Like all Plaid identifiers, this is case-sensitive.

[`version`](/docs/api/products/identity-verification/#identity_verification-create-response-template-version)

integerinteger

Version of the associated Identity Verification template.

[`user`](/docs/api/products/identity-verification/#identity_verification-create-response-user)

objectobject

The identity data that was either collected from the user or provided via API in order to perform an Identity Verification.

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-create-response-user-phone-number)

nullablestringnullable, string

A valid phone number in E.164 format.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-create-response-user-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`ip_address`](/docs/api/products/identity-verification/#identity_verification-create-response-user-ip-address)

nullablestringnullable, string

An IPv4 or IPV6 address.

[`email_address`](/docs/api/products/identity-verification/#identity_verification-create-response-user-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`name`](/docs/api/products/identity-verification/#identity_verification-create-response-user-name)

nullableobjectnullable, object

The full name provided by the user. If the user has not submitted their name, this field will be null. Otherwise, both fields are guaranteed to be filled.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-create-response-user-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-create-response-user-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/identity-verification/#identity_verification-create-response-user-address)

nullableobjectnullable, object

Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include postal code

[`street`](/docs/api/products/identity-verification/#identity_verification-create-response-user-address-street)

nullablestringnullable, string

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/identity-verification/#identity_verification-create-response-user-address-street2)

nullablestringnullable, string

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/identity-verification/#identity_verification-create-response-user-address-city)

nullablestringnullable, string

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/identity-verification/#identity_verification-create-response-user-address-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-create-response-user-address-postal-code)

nullablestringnullable, string

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-create-response-user-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-create-response-user-id-number)

nullableobjectnullable, object

ID number submitted by the user, currently used only for the Identity Verification product. If the user has not submitted this data yet, this field will be `null`. Otherwise, both fields are guaranteed to be filled.

[`value`](/docs/api/products/identity-verification/#identity_verification-create-response-user-id-number-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/identity-verification/#identity_verification-create-response-user-id-number-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-status)

stringstring

The status of this Identity Verification attempt.  
`active` - The Identity Verification attempt is incomplete. The user may have completed part of the session, but has neither failed or passed.  
`success` - The Identity Verification attempt has completed, passing all steps defined to the associated Identity Verification template  
`failed` - The user failed one or more steps in the session and was told to contact support.  
`expired` - The Identity Verification attempt was active for a long period of time without being completed and was automatically marked as expired. Note that sessions currently do not expire. Automatic expiration is expected to be enabled in the future.  
`canceled` - The Identity Verification attempt was canceled, either via the dashboard by a user, or via API. The user may have completed part of the session, but has neither failed or passed.  
`pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  
  

Possible values: `active`, `success`, `failed`, `expired`, `canceled`, `pending_review`

[`steps`](/docs/api/products/identity-verification/#identity_verification-create-response-steps)

objectobject

Each step will be one of the following values:  
`active` - This step is the user's current step. They are either in the process of completing this step, or they recently closed their Identity Verification attempt while in the middle of this step. Only one step will be marked as `active` at any given point.  
`success` - The Identity Verification attempt has completed this step.  
`failed` - The user failed this step. This can either call the user to fail the session as a whole, or cause them to fallback to another step depending on how the Identity Verification template is configured. A failed step does not imply a failed session.  
`waiting_for_prerequisite` - The user needs to complete another step first, before they progress to this step. This step may never run, depending on if the user fails an earlier step or if the step is only run as a fallback.  
`not_applicable` - This step will not be run for this session.  
`skipped` - The retry instructions that created this Identity Verification attempt specified that this step should be skipped.  
`expired` - This step had not yet been completed when the Identity Verification attempt as a whole expired.  
`canceled` - The Identity Verification attempt was canceled before the user completed this step.  
`pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  
`manually_approved` - The step was manually overridden to pass by a team member in the dashboard.  
`manually_rejected` - The step was manually overridden to fail by a team member in the dashboard.

[`accept_tos`](/docs/api/products/identity-verification/#identity_verification-create-response-steps-accept-tos)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-create-response-steps-verify-sms)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-create-response-steps-kyc-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-create-response-steps-documentary-verification)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-create-response-steps-selfie-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`watchlist_screening`](/docs/api/products/identity-verification/#identity_verification-create-response-steps-watchlist-screening)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`risk_check`](/docs/api/products/identity-verification/#identity_verification-create-response-steps-risk-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification)

nullableobjectnullable, object

Data, images, analysis, and results from the `documentary_verification` step. This field will be `null` unless `steps.documentary_verification` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-status)

stringstring

The outcome status for the associated Identity Verification attempt's `documentary_verification` step. This field will always have the same value as `steps.documentary_verification`.

[`documents`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents)

[object][object]

An array of documents submitted to the `documentary_verification` step. Each entry represents one user submission, where each submission will contain both a front and back image, or just a front image, depending on the document type.  
Note: Plaid will automatically let a user submit a new set of document images up to three times if we detect that a previous attempt might have failed due to user error. For example, if the first set of document images are blurry or obscured by glare, the user will be asked to capture their documents again, resulting in at least two separate entries within `documents`. If the overall `documentary_verification` is `failed`, the user has exhausted their retry attempts.

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-status)

stringstring

An outcome status for this specific document submission. Distinct from the overall `documentary_verification.status` that summarizes the verification outcome from one or more documents.  
  

Possible values: `success`, `failed`, `manually_approved`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-attempt)

integerinteger

The `attempt` field begins with 1 and increments with each subsequent document upload.

[`images`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-images)

objectobject

URLs for downloading original and cropped images for this document submission. The URLs are designed to only allow downloading, not hot linking, so the URL will only serve the document image for 60 seconds before expiring. The expiration time is 60 seconds after the `GET` request for the associated Identity Verification attempt. A new expiring URL is generated with each request, so you can always rerequest the Identity Verification attempt if one of your URLs expires.

[`original_front`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-images-original-front)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading the uncropped original image of the front of the document.

[`original_back`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-images-original-back)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading the original image of the back of the document. Might be null if the back of the document was not collected.

[`cropped_front`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-images-cropped-front)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a cropped image containing just the front of the document.

[`cropped_back`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-images-cropped-back)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a cropped image containing just the back of the document. Might be null if the back of the document was not collected.

[`face`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-images-face)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a crop of just the user's face from the document image. Might be null if the document does not contain a face photo.

[`extracted_data`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data)

nullableobjectnullable, object

Data extracted from a user-submitted document.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-id-number)

nullablestringnullable, string

Alpha-numeric ID number extracted via OCR from the user's document image.

[`category`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-category)

stringstring

The type of identity document detected in the images provided. Will always be one of the following values:  
 `drivers_license` - A driver's license issued by the associated country, establishing identity without any guarantee as to citizenship, and granting driving privileges  
 `id_card` - A general national identification card, distinct from driver's licenses as it only establishes identity  
 `passport` - A travel passport issued by the associated country for one of its citizens  
 `residence_permit_card` - An identity document issued by the associated country permitting a foreign citizen to *temporarily* reside there  
 `resident_card` - An identity document issued by the associated country permitting a foreign citizen to *permanently* reside there  
 `visa` - An identity document issued by the associated country permitting a foreign citizen entry for a short duration and for a specific purpose, typically no longer than 6 months  
Note: This value may be different from the ID type that the user selects within Link. For example, if they select "Driver's License" but then submit a picture of a passport, this field will say `passport`  
  

Possible values: `drivers_license`, `id_card`, `passport`, `residence_permit_card`, `resident_card`, `visa`

[`expiration_date`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-expiration-date)

nullablestringnullable, string

The expiration date of the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`issue_date`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-issue-date)

nullablestringnullable, string

The issue date of the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`issuing_country`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-issuing-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`issuing_region`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-issuing-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-date-of-birth)

nullablestringnullable, string

A date extracted from the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`address`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-address)

nullableobjectnullable, object

The address extracted from the document. The address must at least contain the following fields to be a valid address: `street`, `city`, `country`. If any are missing or unable to be extracted, the address will be null.  
`region`, and `postal_code` may be null based on the addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include postal code  
Note: Optical Character Recognition (OCR) technology may sometimes extract incorrect data from a document.

[`street`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-address-street)

stringstring

The full street address extracted from the document.

[`city`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-address-city)

stringstring

City extracted from the document.

[`region`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-address-region)

nullablestringnullable, string

A subdivision code extracted from the document. Related terms would be "state", "province", "prefecture", "zone", "subdivision", etc. For a full list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they can be inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-address-postal-code)

nullablestringnullable, string

The postal code extracted from the document. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country extracted from the document. Must be in ISO 3166-1 alpha-2 form.

[`name`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-name)

nullableobjectnullable, object

The individual's name extracted from the document.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-extracted-data-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`analysis`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis)

objectobject

High level descriptions of how the associated document was processed. If a document fails verification, the details in the `analysis` object should help clarify why the document was rejected.

[`authenticity`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-authenticity)

stringstring

High level summary of whether the document in the provided image matches the formatting rules and security checks for the associated jurisdiction.  
For example, most identity documents have formatting rules like the following:  
The image of the person's face must have a certain contrast in order to highlight skin tone  
The subject in the document's image must remove eye glasses and pose in a certain way  
The informational fields (name, date of birth, ID number, etc.) must be colored and aligned according to specific rules  
Security features like watermarks and background patterns must be present  
So a `match` status for this field indicates that the document in the provided image seems to conform to the various formatting and security rules associated with the detected document.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`image_quality`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-image-quality)

stringstring

A high level description of the quality of the image the user submitted.  
For example, an image that is blurry, distorted by glare from a nearby light source, or improperly framed might be marked as low or medium quality. Poor quality images are more likely to fail OCR and/or template conformity checks.  
Note: By default, Plaid will let a user recapture document images twice before failing the entire session if we attribute the failure to low image quality.  
  

Possible values: `high`, `medium`, `low`

[`extracted_data`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-extracted-data)

nullableobjectnullable, object

Analysis of the data extracted from the submitted document.

[`name`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-extracted-data-name)

stringstring

A match summary describing the cross comparison between the subject's name, extracted from the document image, and the name they separately provided to identity verification attempt.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-extracted-data-date-of-birth)

stringstring

A match summary describing the cross comparison between the subject's date of birth, extracted from the document image, and the date of birth they separately provided to the identity verification attempt.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`expiration_date`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-extracted-data-expiration-date)

stringstring

A description of whether the associated document was expired when the verification was performed.  
Note: In the case where an expiration date is not present on the document or failed to be extracted, this value will be `no_data`.  
  

Possible values: `not_expired`, `expired`, `no_data`

[`issuing_country`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-extracted-data-issuing-country)

stringstring

A binary match indicator specifying whether the country that issued the provided document matches the country that the user separately provided to Plaid.  
Note: You can configure whether a `no_match` on `issuing_country` fails the `documentary_verification` by editing your Plaid Template.  
  

Possible values: `match`, `no_match`

[`aamva_verification`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification)

nullableobjectnullable, object

Analyzed AAMVA data for the associated hit.  
Note: This field is only available for U.S. driver's licenses issued by participating states.

[`is_verified`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-is-verified)

booleanboolean

The overall outcome of checking the associated hit against the issuing state database.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-id-number)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`id_issue_date`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-id-issue-date)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`id_expiration_date`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-id-expiration-date)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`street`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-street)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`city`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-city)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-postal-code)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-date-of-birth)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`gender`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-gender)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`height`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-height)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`eye_color`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-eye-color)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`first_name`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-first-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`middle_name`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-middle-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`last_name`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-analysis-aamva-verification-last-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-create-response-documentary-verification-documents-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check)

nullableobjectnullable, object

Additional information for the `selfie_check` step. This field will be `null` unless `steps.selfie_check` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-status)

stringstring

The outcome status for the associated Identity Verification attempt's `selfie_check` step. This field will always have the same value as `steps.selfie_check`.  
  

Possible values: `success`, `failed`

[`selfies`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies)

[object][object]

An array of selfies submitted to the `selfie_check` step. Each entry represents one user submission.

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies-status)

stringstring

An outcome status for this specific selfie. Distinct from the overall `selfie_check.status` that summarizes the verification outcome from one or more selfies.  
  

Possible values: `success`, `failed`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies-attempt)

integerinteger

The `attempt` field begins with 1 and increments with each subsequent selfie upload.

[`capture`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies-capture)

objectobject

The image or video capture of a selfie. Only one of image or video URL will be populated per selfie.

[`image_url`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies-capture-image-url)

nullablestringnullable, string

Temporary URL for downloading an image selfie capture.

[`video_url`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies-capture-video-url)

nullablestringnullable, string

Temporary URL for downloading a video selfie capture.

[`analysis`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies-analysis)

objectobject

High level descriptions of how the associated selfie was processed. If a selfie fails verification, the details in the `analysis` object should help clarify why the selfie was rejected.

[`document_comparison`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies-analysis-document-comparison)

stringstring

Information about the comparison between the selfie and the document (if documentary verification also ran).  
  

Possible values: `match`, `no_match`, `no_input`

[`liveness_check`](/docs/api/products/identity-verification/#identity_verification-create-response-selfie-check-selfies-analysis-liveness-check)

stringstring

Assessment of whether the selfie capture is of a real human being, as opposed to a picture of a human on a screen, a picture of a paper cut out, someone wearing a mask, or a deepfake.  
  

Possible values: `success`, `failed`

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check)

nullableobjectnullable, object

Additional information for the `kyc_check` (Data Source Verification) step. This field will be `null` unless `steps.kyc_check` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-status)

stringstring

The outcome status for the associated Identity Verification attempt's `kyc_check` step. This field will always have the same value as `steps.kyc_check`.

[`address`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-address)

objectobject

Result summary object specifying how the `address` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-address-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`po_box`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-address-po-box)

stringstring

Field describing whether the associated address is a post office box. Will be `yes` when a P.O. box is detected, `no` when Plaid confirmed the address is not a P.O. box, and `no_data` when Plaid was not able to determine if the address is a P.O. box.  
  

Possible values: `yes`, `no`, `no_data`

[`type`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-address-type)

stringstring

Field describing whether the associated address is being used for commercial or residential purposes.  
Note: This value will be `no_data` when Plaid does not have sufficient data to determine the address's use.  
  

Possible values: `residential`, `commercial`, `no_data`

[`name`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-name)

objectobject

Result summary object specifying how the `name` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-name-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-date-of-birth)

objectobject

Result summary object specifying how the `date_of_birth` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-date-of-birth-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`id_number`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-id-number)

objectobject

Result summary object specifying how the `id_number` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-id-number-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-phone-number)

objectobject

Result summary object specifying how the `phone` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-phone-number-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`area_code`](/docs/api/products/identity-verification/#identity_verification-create-response-kyc-check-phone-number-area-code)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`risk_check`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check)

nullableobjectnullable, object

Additional information for the `risk_check` step.

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-status)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`behavior`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-behavior)

nullableobjectnullable, object

Result summary object specifying values for `behavior` attributes of risk check, when available.

[`user_interactions`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-behavior-user-interactions)

stringstring

Field describing the overall user interaction signals of a behavior risk check. This value represents how familiar the user is with the personal data they provide, based on a number of signals that are collected during their session.  
`genuine` indicates the user has high familiarity with the data they are providing, and that fraud is unlikely.  
`neutral` indicates some signals are present in between `risky` and `genuine`, but there are not enough clear signals to determine an outcome.  
`risky` indicates the user has low familiarity with the data they are providing, and that fraud is likely.  
`no_data` indicates there is not sufficient information to give an accurate signal.  
  

Possible values: `genuine`, `neutral`, `risky`, `no_data`

[`fraud_ring_detected`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-behavior-fraud-ring-detected)

stringstring

Field describing the outcome of a fraud ring behavior risk check.  
`yes` indicates that fraud ring activity was detected.  
`no` indicates that fraud ring activity was not detected.  
`no_data` indicates there was not enough information available to give an accurate signal.  
  

Possible values: `yes`, `no`, `no_data`

[`bot_detected`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-behavior-bot-detected)

stringstring

Field describing the outcome of a bot detection behavior risk check.  
`yes` indicates that automated activity was detected.  
`no` indicates that automated activity was not detected.  
`no_data` indicates there was not enough information available to give an accurate signal.  
  

Possible values: `yes`, `no`, `no_data`

[`email`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email)

nullableobjectnullable, object

Result summary object specifying values for `email` attributes of risk check.

[`is_deliverable`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-is-deliverable)

stringstring

SMTP-MX check to confirm the email address exists if known.  
  

Possible values: `yes`, `no`, `no_data`

[`breach_count`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-breach-count)

nullableintegernullable, integer

Count of all known breaches of this email address if known.

[`first_breached_at`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-first-breached-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`last_breached_at`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-last-breached-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`domain_registered_at`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-domain-registered-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`domain_is_free_provider`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-domain-is-free-provider)

stringstring

Indicates whether the email address domain is a free provider such as Gmail or Hotmail if known.  
  

Possible values: `yes`, `no`, `no_data`

[`domain_is_custom`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-domain-is-custom)

stringstring

Indicates whether the email address domain is custom if known, i.e. a company domain and not free or disposable.  
  

Possible values: `yes`, `no`, `no_data`

[`domain_is_disposable`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-domain-is-disposable)

stringstring

Indicates whether the email domain is listed as disposable if known. Disposable domains are often used to create email addresses that are part of a fake set of user details.  
  

Possible values: `yes`, `no`, `no_data`

[`top_level_domain_is_suspicious`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-top-level-domain-is-suspicious)

stringstring

Indicates whether the email address top level domain, which is the last part of the domain, is fraudulent or risky if known. In most cases, a suspicious top level domain is also associated with a disposable or high-risk domain.  
  

Possible values: `yes`, `no`, `no_data`

[`linked_services`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-email-linked-services)

[string][string]

A list of online services where this email address has been detected to have accounts or other activity.  
  

Possible values: `aboutme`, `adobe`, `adult_sites`, `airbnb`, `altbalaji`, `amazon`, `apple`, `archiveorg`, `atlassian`, `bitmoji`, `bodybuilding`, `booking`, `bukalapak`, `codecademy`, `deliveroo`, `diigo`, `discord`, `disneyplus`, `duolingo`, `ebay`, `envato`, `eventbrite`, `evernote`, `facebook`, `firefox`, `flickr`, `flipkart`, `foursquare`, `freelancer`, `gaana`, `giphy`, `github`, `google`, `gravatar`, `hubspot`, `imgur`, `instagram`, `jdid`, `kakao`, `kommo`, `komoot`, `lastfm`, `lazada`, `line`, `linkedin`, `mailru`, `microsoft`, `myspace`, `netflix`, `nike`, `ok`, `patreon`, `pinterest`, `plurk`, `quora`, `qzone`, `rambler`, `rappi`, `replit`, `samsung`, `seoclerks`, `shopclues`, `skype`, `snapchat`, `snapdeal`, `soundcloud`, `spotify`, `starz`, `strava`, `taringa`, `telegram`, `tiki`, `tokopedia`, `treehouse`, `tumblr`, `twitter`, `venmo`, `viber`, `vimeo`, `vivino`, `vkontakte`, `wattpad`, `weibo`, `whatsapp`, `wordpress`, `xing`, `yahoo`, `yandex`, `zalo`, `zoho`

[`phone`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-phone)

nullableobjectnullable, object

Result summary object specifying values for `phone` attributes of risk check.

[`linked_services`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-phone-linked-services)

[string][string]

A list of online services where this phone number has been detected to have accounts or other activity.  
  

Possible values: `aboutme`, `adobe`, `adult_sites`, `airbnb`, `altbalaji`, `amazon`, `apple`, `archiveorg`, `atlassian`, `bitmoji`, `bodybuilding`, `booking`, `bukalapak`, `codecademy`, `deliveroo`, `diigo`, `discord`, `disneyplus`, `duolingo`, `ebay`, `envato`, `eventbrite`, `evernote`, `facebook`, `firefox`, `flickr`, `flipkart`, `foursquare`, `freelancer`, `gaana`, `giphy`, `github`, `google`, `gravatar`, `hubspot`, `imgur`, `instagram`, `jdid`, `kakao`, `kommo`, `komoot`, `lastfm`, `lazada`, `line`, `linkedin`, `mailru`, `microsoft`, `myspace`, `netflix`, `nike`, `ok`, `patreon`, `pinterest`, `plurk`, `quora`, `qzone`, `rambler`, `rappi`, `replit`, `samsung`, `seoclerks`, `shopclues`, `skype`, `snapchat`, `snapdeal`, `soundcloud`, `spotify`, `starz`, `strava`, `taringa`, `telegram`, `tiki`, `tokopedia`, `treehouse`, `tumblr`, `twitter`, `venmo`, `viber`, `vimeo`, `vivino`, `vkontakte`, `wattpad`, `weibo`, `whatsapp`, `wordpress`, `xing`, `yahoo`, `yandex`, `zalo`, `zoho`

[`devices`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-devices)

[object][object]

Array of result summary objects specifying values for `device` attributes of risk check.

[`ip_proxy_type`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-devices-ip-proxy-type)

nullablestringnullable, string

An enum indicating whether a network proxy is present and if so what type it is.  
`none_detected` indicates the user is not on a detectable proxy network.  
`tor` indicates the user was using a Tor browser, which sends encrypted traffic on a decentralized network and is somewhat similar to a VPN (Virtual Private Network).  
`vpn` indicates the user is on a VPN (Virtual Private Network)  
`web_proxy` indicates the user is on a web proxy server, which may allow them to conceal information such as their IP address or other identifying information.  
`public_proxy` indicates the user is on a public web proxy server, which is similar to a web proxy but can be shared by multiple users. This may allow multiple users to appear as if they have the same IP address for instance.  
  

Possible values: `none_detected`, `tor`, `vpn`, `web_proxy`, `public_proxy`

[`ip_spam_list_count`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-devices-ip-spam-list-count)

nullableintegernullable, integer

Count of spam lists the IP address is associated with if known.

[`ip_timezone_offset`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-devices-ip-timezone-offset)

nullablestringnullable, string

UTC offset of the timezone associated with the IP address.

[`identity_abuse_signals`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-identity-abuse-signals)

nullableobjectnullable, object

Result summary object capturing abuse signals related to `identity abuse`, e.g. stolen and synthetic identity fraud. These attributes are only available for US identities and some signals may not be available depending on what information was collected.

[`synthetic_identity`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-identity-abuse-signals-synthetic-identity)

nullableobjectnullable, object

Field containing the data used in determining the outcome of the synthetic identity risk check.  
Contains the following fields:  
`score` - A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.

[`score`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-identity-abuse-signals-synthetic-identity-score)

integerinteger

A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.

[`stolen_identity`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-identity-abuse-signals-stolen-identity)

nullableobjectnullable, object

Field containing the data used in determining the outcome of the stolen identity risk check.  
Contains the following fields:  
`score` - A score from 0 to 100 indicating the likelihood that the user is a stolen identity.

[`score`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-identity-abuse-signals-stolen-identity-score)

integerinteger

A score from 0 to 100 indicating the likelihood that the user is a stolen identity.

[`facial_duplicates`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-facial-duplicates)

[object][object]

The attributes related to the facial duplicates captured in risk check.

[`id`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-facial-duplicates-id)

stringstring

ID of the associated Identity Verification attempt.

[`similarity`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-facial-duplicates-similarity)

integerinteger

Similarity score of the match. Ranges from 0 to 100.

[`matched_after_completed`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-facial-duplicates-matched-after-completed)

booleanboolean

Whether this match occurred after the session was complete. For example, this would be `true` if a later session ended up matching this one.

[`trust_index_score`](/docs/api/products/identity-verification/#identity_verification-create-response-risk-check-trust-index-score)

nullableintegernullable, integer

The trust index score for the `risk_check` step.

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms)

nullableobjectnullable, object

Additional information for the `verify_sms` step.

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-status)

stringstring

The outcome status for the associated Identity Verification attempt's `verify_sms` step. This field will always have the same value as `steps.verify_sms`.  
  

Possible values: `success`, `failed`

[`verifications`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications)

[object][object]

An array where each entry represents a verification attempt for the `verify_sms` step. Each entry represents one user-submitted phone number. Phone number edits, and in some cases error handling due to edge cases like rate limiting, may generate additional verifications.

[`status`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications-status)

stringstring

The outcome status for the individual SMS verification.  
  

Possible values: `pending`, `success`, `failed`, `canceled`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications-attempt)

integerinteger

The attempt field begins with 1 and increments with each subsequent SMS verification.

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`delivery_attempt_count`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications-delivery-attempt-count)

integerinteger

The number of delivery attempts made within the verification to send the SMS code to the user. Each delivery attempt represents the user taking action from the front end UI to request creation and delivery of a new SMS verification code, or to resend an existing SMS verification code. There is a limit of 3 delivery attempts per verification.

[`solve_attempt_count`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications-solve-attempt-count)

integerinteger

The number of attempts made by the user within the verification to verify the SMS code by entering it into the front end UI. There is a limit of 3 solve attempts per verification.

[`initially_sent_at`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications-initially-sent-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`last_sent_at`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications-last-sent-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-create-response-verify-sms-verifications-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`watchlist_screening_id`](/docs/api/products/identity-verification/#identity_verification-create-response-watchlist-screening-id)

nullablestringnullable, string

ID of the associated screening.

[`beacon_user_id`](/docs/api/products/identity-verification/#identity_verification-create-response-beacon-user-id)

nullablestringnullable, string

ID of the associated Beacon User.

[`user_id`](/docs/api/products/identity-verification/#identity_verification-create-response-user-id)

nullablestringnullable, string

Unique user identifier, created by calling `/user/create`. Either a `user_id` or the `client_user_id` must be provided. The `user_id` may only be used instead of the `client_user_id` if you were not a pre-existing user of `/user/create` as of December 10, 2025; for more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis). If both this field and a `client_user_id` are present in a request, the `user_id` must have been created from the provided `client_user_id`.

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-create-response-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`latest_scored_protect_event`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event)

nullableobjectnullable, object

Information about a Protect event including Trust Index score and fraud attributes.

[`event_id`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-event-id)

stringstring

The event ID.

[`timestamp`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-timestamp)

stringstring

The timestamp of the event, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2017-09-14T14:42:19.350Z"`  
  

Format: `date-time`

[`trust_index`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-trust-index)

nullableobjectnullable, object

Represents a calculate Trust Index Score.

[`score`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-trust-index-score)

integerinteger

The overall trust index score.

[`model`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-trust-index-model)

stringstring

The versioned name of the Trust Index model used for scoring.

[`subscores`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-trust-index-subscores)

nullableobjectnullable, object

Contains sub-score metadata.

[`device_and_connection`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-trust-index-subscores-device-and-connection)

nullableobjectnullable, object

Represents Trust Index Subscore.

[`score`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-trust-index-subscores-device-and-connection-score)

integerinteger

The subscore score.

[`bank_account_insights`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-trust-index-subscores-bank-account-insights)

nullableobjectnullable, object

Represents Trust Index Subscore.

[`score`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-trust-index-subscores-bank-account-insights-score)

integerinteger

The subscore score.

[`fraud_attributes`](/docs/api/products/identity-verification/#identity_verification-create-response-latest-scored-protect-event-fraud-attributes)

nullableobjectnullable, object

Event fraud attributes as an arbitrary set of key-value pairs.

[`request_id`](/docs/api/products/identity-verification/#identity_verification-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "idv_52xR9LKo77r1Np",
  "client_user_id": "your-db-id-3b24110",
  "created_at": "2020-07-24T03:26:02Z",
  "completed_at": "2020-07-24T03:26:02Z",
  "previous_attempt_id": "idv_42cF1MNo42r9Xj",
  "shareable_url": "https://flow.plaid.com/verify/idv_4FrXJvfQU3zGUR?key=e004115db797f7cc3083bff3167cba30644ef630fb46f5b086cde6cc3b86a36f",
  "template": {
    "id": "idvtmp_4FrXJvfQU3zGUR",
    "version": 2
  },
  "user": {
    "phone_number": "+12345678909",
    "date_of_birth": "1990-05-29",
    "ip_address": "192.0.2.42",
    "email_address": "user@example.com",
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
    "id_number": {
      "value": "123456789",
      "type": "us_ssn"
    }
  },
  "status": "success",
  "steps": {
    "accept_tos": "success",
    "verify_sms": "success",
    "kyc_check": "success",
    "documentary_verification": "success",
    "selfie_check": "success",
    "watchlist_screening": "success",
    "risk_check": "success"
  },
  "documentary_verification": {
    "status": "success",
    "documents": [
      {
        "status": "success",
        "attempt": 1,
        "images": {
          "original_front": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/original_front.jpeg",
          "original_back": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/original_back.jpeg",
          "cropped_front": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/cropped_front.jpeg",
          "cropped_back": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/cropped_back.jpeg",
          "face": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/face.jpeg"
        },
        "extracted_data": {
          "id_number": "AB123456",
          "category": "drivers_license",
          "expiration_date": "2030-05-29",
          "issue_date": "2020-05-29",
          "issuing_country": "US",
          "issuing_region": "IN",
          "date_of_birth": "1990-05-29",
          "address": {
            "street": "123 Main St. Unit 42",
            "city": "Pawnee",
            "region": "IN",
            "postal_code": "46001",
            "country": "US"
          },
          "name": {
            "given_name": "Leslie",
            "family_name": "Knope"
          }
        },
        "analysis": {
          "authenticity": "match",
          "image_quality": "high",
          "extracted_data": {
            "name": "match",
            "date_of_birth": "match",
            "expiration_date": "not_expired",
            "issuing_country": "match"
          },
          "aamva_verification": {
            "is_verified": true,
            "id_number": "match",
            "id_issue_date": "match",
            "id_expiration_date": "match",
            "street": "match",
            "city": "match",
            "postal_code": "match",
            "date_of_birth": "match",
            "gender": "match",
            "height": "match",
            "eye_color": "match",
            "first_name": "match",
            "middle_name": "match",
            "last_name": "match"
          }
        },
        "redacted_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "selfie_check": {
    "status": "success",
    "selfies": [
      {
        "status": "success",
        "attempt": 1,
        "capture": {
          "image_url": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/selfie/liveness.jpeg",
          "video_url": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/selfie/liveness.webm"
        },
        "analysis": {
          "document_comparison": "match",
          "liveness_check": "success"
        }
      }
    ]
  },
  "kyc_check": {
    "status": "success",
    "address": {
      "summary": "match",
      "po_box": "yes",
      "type": "residential"
    },
    "name": {
      "summary": "match"
    },
    "date_of_birth": {
      "summary": "match"
    },
    "id_number": {
      "summary": "match"
    },
    "phone_number": {
      "summary": "match",
      "area_code": "match"
    }
  },
  "risk_check": {
    "status": "success",
    "behavior": {
      "user_interactions": "risky",
      "fraud_ring_detected": "yes",
      "bot_detected": "yes"
    },
    "email": {
      "is_deliverable": "yes",
      "breach_count": 1,
      "first_breached_at": "1990-05-29",
      "last_breached_at": "1990-05-29",
      "domain_registered_at": "1990-05-29",
      "domain_is_free_provider": "yes",
      "domain_is_custom": "yes",
      "domain_is_disposable": "yes",
      "top_level_domain_is_suspicious": "yes",
      "linked_services": [
        "apple"
      ]
    },
    "phone": {
      "linked_services": [
        "apple"
      ]
    },
    "devices": [
      {
        "ip_proxy_type": "none_detected",
        "ip_spam_list_count": 1,
        "ip_timezone_offset": "+06:00:00"
      }
    ],
    "identity_abuse_signals": {
      "synthetic_identity": {
        "score": 0
      },
      "stolen_identity": {
        "score": 0
      }
    },
    "facial_duplicates": [
      {
        "id": "idv_52xR9LKo77r1Np",
        "similarity": 95,
        "matched_after_completed": true
      }
    ],
    "trust_index_score": 86
  },
  "verify_sms": {
    "status": "success",
    "verifications": [
      {
        "status": "success",
        "attempt": 1,
        "phone_number": "+12345678909",
        "delivery_attempt_count": 1,
        "solve_attempt_count": 1,
        "initially_sent_at": "2020-07-24T03:26:02Z",
        "last_sent_at": "2020-07-24T03:26:02Z",
        "redacted_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "watchlist_screening_id": "scr_52xR9LKo77r1Np",
  "beacon_user_id": "becusr_42cF1MNo42r9Xj",
  "user_id": "usr_dddAs9ewdcDQQQ",
  "redacted_at": "2020-07-24T03:26:02Z",
  "latest_scored_protect_event": {
    "event_id": "ptevt_7AJYTMFxRUgJ",
    "timestamp": "2020-07-24T03:26:02Z",
    "trust_index": {
      "score": 86,
      "model": "trust_index.2.0.0",
      "subscores": {
        "device_and_connection": {
          "score": 87
        },
        "bank_account_insights": {
          "score": 85
        }
      }
    },
    "fraud_attributes": {
      "fraud_attributes": {
        "link_session.linked_bank_accounts.user_pi_matches_owners": true,
        "link_session.linked_bank_accounts.connected_apps.days_since_first_connection": 582,
        "session.challenged_with_mfa": false,
        "user.bank_accounts.num_of_frozen_or_restricted_accounts": 0,
        "user.linked_bank_accounts.num_family_names": 1,
        "user.linked_bank_accounts.num_of_connected_banks": 1,
        "user.link_sessions.days_since_first_link_session": 150,
        "user.pi.email.history_yrs": 7.03,
        "user.pi.email.num_social_networks_linked": 12,
        "user.pi.ssn.user_likely_has_better_ssn": false
      }
    }
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/identity_verification/get`

#### Retrieve Identity Verification

Retrieve a previously created Identity Verification.

/identity\_verification/get

**Request fields**

[`identity_verification_id`](/docs/api/products/identity-verification/#identity_verification-get-request-identity-verification-id)

requiredstringrequired, string

ID of the associated Identity Verification attempt.

[`secret`](/docs/api/products/identity-verification/#identity_verification-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/identity-verification/#identity_verification-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

/identity\_verification/get

```
const request: IdentityVerificationGetRequest = {
  identity_verification_id: 'idv_52xR9LKo77r1Np',
};
try {
  const response = await plaidClient.identityVerificationGet(request);
} catch (error) {
  // handle error
}
```

/identity\_verification/get

**Response fields**

[`id`](/docs/api/products/identity-verification/#identity_verification-get-response-id)

stringstring

ID of the associated Identity Verification attempt.

[`client_user_id`](/docs/api/products/identity-verification/#identity_verification-get-response-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`created_at`](/docs/api/products/identity-verification/#identity_verification-get-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`completed_at`](/docs/api/products/identity-verification/#identity_verification-get-response-completed-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`previous_attempt_id`](/docs/api/products/identity-verification/#identity_verification-get-response-previous-attempt-id)

nullablestringnullable, string

The ID for the Identity Verification preceding this session. This field will only be filled if the current Identity Verification is a retry of a previous attempt.

[`shareable_url`](/docs/api/products/identity-verification/#identity_verification-get-response-shareable-url)

nullablestringnullable, string

A shareable URL that can be sent directly to the user to complete verification

[`template`](/docs/api/products/identity-verification/#identity_verification-get-response-template)

objectobject

The resource ID and version number of the template configuring the behavior of a given Identity Verification.

[`id`](/docs/api/products/identity-verification/#identity_verification-get-response-template-id)

stringstring

ID of the associated Identity Verification template. Like all Plaid identifiers, this is case-sensitive.

[`version`](/docs/api/products/identity-verification/#identity_verification-get-response-template-version)

integerinteger

Version of the associated Identity Verification template.

[`user`](/docs/api/products/identity-verification/#identity_verification-get-response-user)

objectobject

The identity data that was either collected from the user or provided via API in order to perform an Identity Verification.

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-get-response-user-phone-number)

nullablestringnullable, string

A valid phone number in E.164 format.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-get-response-user-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`ip_address`](/docs/api/products/identity-verification/#identity_verification-get-response-user-ip-address)

nullablestringnullable, string

An IPv4 or IPV6 address.

[`email_address`](/docs/api/products/identity-verification/#identity_verification-get-response-user-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`name`](/docs/api/products/identity-verification/#identity_verification-get-response-user-name)

nullableobjectnullable, object

The full name provided by the user. If the user has not submitted their name, this field will be null. Otherwise, both fields are guaranteed to be filled.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-get-response-user-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-get-response-user-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/identity-verification/#identity_verification-get-response-user-address)

nullableobjectnullable, object

Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include postal code

[`street`](/docs/api/products/identity-verification/#identity_verification-get-response-user-address-street)

nullablestringnullable, string

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/identity-verification/#identity_verification-get-response-user-address-street2)

nullablestringnullable, string

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/identity-verification/#identity_verification-get-response-user-address-city)

nullablestringnullable, string

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/identity-verification/#identity_verification-get-response-user-address-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-get-response-user-address-postal-code)

nullablestringnullable, string

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-get-response-user-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-get-response-user-id-number)

nullableobjectnullable, object

ID number submitted by the user, currently used only for the Identity Verification product. If the user has not submitted this data yet, this field will be `null`. Otherwise, both fields are guaranteed to be filled.

[`value`](/docs/api/products/identity-verification/#identity_verification-get-response-user-id-number-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/identity-verification/#identity_verification-get-response-user-id-number-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-status)

stringstring

The status of this Identity Verification attempt.  
`active` - The Identity Verification attempt is incomplete. The user may have completed part of the session, but has neither failed or passed.  
`success` - The Identity Verification attempt has completed, passing all steps defined to the associated Identity Verification template  
`failed` - The user failed one or more steps in the session and was told to contact support.  
`expired` - The Identity Verification attempt was active for a long period of time without being completed and was automatically marked as expired. Note that sessions currently do not expire. Automatic expiration is expected to be enabled in the future.  
`canceled` - The Identity Verification attempt was canceled, either via the dashboard by a user, or via API. The user may have completed part of the session, but has neither failed or passed.  
`pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  
  

Possible values: `active`, `success`, `failed`, `expired`, `canceled`, `pending_review`

[`steps`](/docs/api/products/identity-verification/#identity_verification-get-response-steps)

objectobject

Each step will be one of the following values:  
`active` - This step is the user's current step. They are either in the process of completing this step, or they recently closed their Identity Verification attempt while in the middle of this step. Only one step will be marked as `active` at any given point.  
`success` - The Identity Verification attempt has completed this step.  
`failed` - The user failed this step. This can either call the user to fail the session as a whole, or cause them to fallback to another step depending on how the Identity Verification template is configured. A failed step does not imply a failed session.  
`waiting_for_prerequisite` - The user needs to complete another step first, before they progress to this step. This step may never run, depending on if the user fails an earlier step or if the step is only run as a fallback.  
`not_applicable` - This step will not be run for this session.  
`skipped` - The retry instructions that created this Identity Verification attempt specified that this step should be skipped.  
`expired` - This step had not yet been completed when the Identity Verification attempt as a whole expired.  
`canceled` - The Identity Verification attempt was canceled before the user completed this step.  
`pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  
`manually_approved` - The step was manually overridden to pass by a team member in the dashboard.  
`manually_rejected` - The step was manually overridden to fail by a team member in the dashboard.

[`accept_tos`](/docs/api/products/identity-verification/#identity_verification-get-response-steps-accept-tos)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-get-response-steps-verify-sms)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-get-response-steps-kyc-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-get-response-steps-documentary-verification)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-get-response-steps-selfie-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`watchlist_screening`](/docs/api/products/identity-verification/#identity_verification-get-response-steps-watchlist-screening)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`risk_check`](/docs/api/products/identity-verification/#identity_verification-get-response-steps-risk-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification)

nullableobjectnullable, object

Data, images, analysis, and results from the `documentary_verification` step. This field will be `null` unless `steps.documentary_verification` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-status)

stringstring

The outcome status for the associated Identity Verification attempt's `documentary_verification` step. This field will always have the same value as `steps.documentary_verification`.

[`documents`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents)

[object][object]

An array of documents submitted to the `documentary_verification` step. Each entry represents one user submission, where each submission will contain both a front and back image, or just a front image, depending on the document type.  
Note: Plaid will automatically let a user submit a new set of document images up to three times if we detect that a previous attempt might have failed due to user error. For example, if the first set of document images are blurry or obscured by glare, the user will be asked to capture their documents again, resulting in at least two separate entries within `documents`. If the overall `documentary_verification` is `failed`, the user has exhausted their retry attempts.

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-status)

stringstring

An outcome status for this specific document submission. Distinct from the overall `documentary_verification.status` that summarizes the verification outcome from one or more documents.  
  

Possible values: `success`, `failed`, `manually_approved`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-attempt)

integerinteger

The `attempt` field begins with 1 and increments with each subsequent document upload.

[`images`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-images)

objectobject

URLs for downloading original and cropped images for this document submission. The URLs are designed to only allow downloading, not hot linking, so the URL will only serve the document image for 60 seconds before expiring. The expiration time is 60 seconds after the `GET` request for the associated Identity Verification attempt. A new expiring URL is generated with each request, so you can always rerequest the Identity Verification attempt if one of your URLs expires.

[`original_front`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-images-original-front)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading the uncropped original image of the front of the document.

[`original_back`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-images-original-back)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading the original image of the back of the document. Might be null if the back of the document was not collected.

[`cropped_front`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-images-cropped-front)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a cropped image containing just the front of the document.

[`cropped_back`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-images-cropped-back)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a cropped image containing just the back of the document. Might be null if the back of the document was not collected.

[`face`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-images-face)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a crop of just the user's face from the document image. Might be null if the document does not contain a face photo.

[`extracted_data`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data)

nullableobjectnullable, object

Data extracted from a user-submitted document.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-id-number)

nullablestringnullable, string

Alpha-numeric ID number extracted via OCR from the user's document image.

[`category`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-category)

stringstring

The type of identity document detected in the images provided. Will always be one of the following values:  
 `drivers_license` - A driver's license issued by the associated country, establishing identity without any guarantee as to citizenship, and granting driving privileges  
 `id_card` - A general national identification card, distinct from driver's licenses as it only establishes identity  
 `passport` - A travel passport issued by the associated country for one of its citizens  
 `residence_permit_card` - An identity document issued by the associated country permitting a foreign citizen to *temporarily* reside there  
 `resident_card` - An identity document issued by the associated country permitting a foreign citizen to *permanently* reside there  
 `visa` - An identity document issued by the associated country permitting a foreign citizen entry for a short duration and for a specific purpose, typically no longer than 6 months  
Note: This value may be different from the ID type that the user selects within Link. For example, if they select "Driver's License" but then submit a picture of a passport, this field will say `passport`  
  

Possible values: `drivers_license`, `id_card`, `passport`, `residence_permit_card`, `resident_card`, `visa`

[`expiration_date`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-expiration-date)

nullablestringnullable, string

The expiration date of the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`issue_date`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-issue-date)

nullablestringnullable, string

The issue date of the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`issuing_country`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-issuing-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`issuing_region`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-issuing-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-date-of-birth)

nullablestringnullable, string

A date extracted from the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`address`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-address)

nullableobjectnullable, object

The address extracted from the document. The address must at least contain the following fields to be a valid address: `street`, `city`, `country`. If any are missing or unable to be extracted, the address will be null.  
`region`, and `postal_code` may be null based on the addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include postal code  
Note: Optical Character Recognition (OCR) technology may sometimes extract incorrect data from a document.

[`street`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-address-street)

stringstring

The full street address extracted from the document.

[`city`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-address-city)

stringstring

City extracted from the document.

[`region`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-address-region)

nullablestringnullable, string

A subdivision code extracted from the document. Related terms would be "state", "province", "prefecture", "zone", "subdivision", etc. For a full list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they can be inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-address-postal-code)

nullablestringnullable, string

The postal code extracted from the document. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country extracted from the document. Must be in ISO 3166-1 alpha-2 form.

[`name`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-name)

nullableobjectnullable, object

The individual's name extracted from the document.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-extracted-data-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`analysis`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis)

objectobject

High level descriptions of how the associated document was processed. If a document fails verification, the details in the `analysis` object should help clarify why the document was rejected.

[`authenticity`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-authenticity)

stringstring

High level summary of whether the document in the provided image matches the formatting rules and security checks for the associated jurisdiction.  
For example, most identity documents have formatting rules like the following:  
The image of the person's face must have a certain contrast in order to highlight skin tone  
The subject in the document's image must remove eye glasses and pose in a certain way  
The informational fields (name, date of birth, ID number, etc.) must be colored and aligned according to specific rules  
Security features like watermarks and background patterns must be present  
So a `match` status for this field indicates that the document in the provided image seems to conform to the various formatting and security rules associated with the detected document.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`image_quality`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-image-quality)

stringstring

A high level description of the quality of the image the user submitted.  
For example, an image that is blurry, distorted by glare from a nearby light source, or improperly framed might be marked as low or medium quality. Poor quality images are more likely to fail OCR and/or template conformity checks.  
Note: By default, Plaid will let a user recapture document images twice before failing the entire session if we attribute the failure to low image quality.  
  

Possible values: `high`, `medium`, `low`

[`extracted_data`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-extracted-data)

nullableobjectnullable, object

Analysis of the data extracted from the submitted document.

[`name`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-extracted-data-name)

stringstring

A match summary describing the cross comparison between the subject's name, extracted from the document image, and the name they separately provided to identity verification attempt.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-extracted-data-date-of-birth)

stringstring

A match summary describing the cross comparison between the subject's date of birth, extracted from the document image, and the date of birth they separately provided to the identity verification attempt.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`expiration_date`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-extracted-data-expiration-date)

stringstring

A description of whether the associated document was expired when the verification was performed.  
Note: In the case where an expiration date is not present on the document or failed to be extracted, this value will be `no_data`.  
  

Possible values: `not_expired`, `expired`, `no_data`

[`issuing_country`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-extracted-data-issuing-country)

stringstring

A binary match indicator specifying whether the country that issued the provided document matches the country that the user separately provided to Plaid.  
Note: You can configure whether a `no_match` on `issuing_country` fails the `documentary_verification` by editing your Plaid Template.  
  

Possible values: `match`, `no_match`

[`aamva_verification`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification)

nullableobjectnullable, object

Analyzed AAMVA data for the associated hit.  
Note: This field is only available for U.S. driver's licenses issued by participating states.

[`is_verified`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-is-verified)

booleanboolean

The overall outcome of checking the associated hit against the issuing state database.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-id-number)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`id_issue_date`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-id-issue-date)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`id_expiration_date`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-id-expiration-date)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`street`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-street)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`city`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-city)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-postal-code)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-date-of-birth)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`gender`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-gender)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`height`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-height)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`eye_color`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-eye-color)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`first_name`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-first-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`middle_name`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-middle-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`last_name`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-analysis-aamva-verification-last-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-get-response-documentary-verification-documents-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check)

nullableobjectnullable, object

Additional information for the `selfie_check` step. This field will be `null` unless `steps.selfie_check` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-status)

stringstring

The outcome status for the associated Identity Verification attempt's `selfie_check` step. This field will always have the same value as `steps.selfie_check`.  
  

Possible values: `success`, `failed`

[`selfies`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies)

[object][object]

An array of selfies submitted to the `selfie_check` step. Each entry represents one user submission.

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies-status)

stringstring

An outcome status for this specific selfie. Distinct from the overall `selfie_check.status` that summarizes the verification outcome from one or more selfies.  
  

Possible values: `success`, `failed`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies-attempt)

integerinteger

The `attempt` field begins with 1 and increments with each subsequent selfie upload.

[`capture`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies-capture)

objectobject

The image or video capture of a selfie. Only one of image or video URL will be populated per selfie.

[`image_url`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies-capture-image-url)

nullablestringnullable, string

Temporary URL for downloading an image selfie capture.

[`video_url`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies-capture-video-url)

nullablestringnullable, string

Temporary URL for downloading a video selfie capture.

[`analysis`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies-analysis)

objectobject

High level descriptions of how the associated selfie was processed. If a selfie fails verification, the details in the `analysis` object should help clarify why the selfie was rejected.

[`document_comparison`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies-analysis-document-comparison)

stringstring

Information about the comparison between the selfie and the document (if documentary verification also ran).  
  

Possible values: `match`, `no_match`, `no_input`

[`liveness_check`](/docs/api/products/identity-verification/#identity_verification-get-response-selfie-check-selfies-analysis-liveness-check)

stringstring

Assessment of whether the selfie capture is of a real human being, as opposed to a picture of a human on a screen, a picture of a paper cut out, someone wearing a mask, or a deepfake.  
  

Possible values: `success`, `failed`

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check)

nullableobjectnullable, object

Additional information for the `kyc_check` (Data Source Verification) step. This field will be `null` unless `steps.kyc_check` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-status)

stringstring

The outcome status for the associated Identity Verification attempt's `kyc_check` step. This field will always have the same value as `steps.kyc_check`.

[`address`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-address)

objectobject

Result summary object specifying how the `address` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-address-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`po_box`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-address-po-box)

stringstring

Field describing whether the associated address is a post office box. Will be `yes` when a P.O. box is detected, `no` when Plaid confirmed the address is not a P.O. box, and `no_data` when Plaid was not able to determine if the address is a P.O. box.  
  

Possible values: `yes`, `no`, `no_data`

[`type`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-address-type)

stringstring

Field describing whether the associated address is being used for commercial or residential purposes.  
Note: This value will be `no_data` when Plaid does not have sufficient data to determine the address's use.  
  

Possible values: `residential`, `commercial`, `no_data`

[`name`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-name)

objectobject

Result summary object specifying how the `name` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-name-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-date-of-birth)

objectobject

Result summary object specifying how the `date_of_birth` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-date-of-birth-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`id_number`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-id-number)

objectobject

Result summary object specifying how the `id_number` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-id-number-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-phone-number)

objectobject

Result summary object specifying how the `phone` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-phone-number-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`area_code`](/docs/api/products/identity-verification/#identity_verification-get-response-kyc-check-phone-number-area-code)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`risk_check`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check)

nullableobjectnullable, object

Additional information for the `risk_check` step.

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-status)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`behavior`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-behavior)

nullableobjectnullable, object

Result summary object specifying values for `behavior` attributes of risk check, when available.

[`user_interactions`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-behavior-user-interactions)

stringstring

Field describing the overall user interaction signals of a behavior risk check. This value represents how familiar the user is with the personal data they provide, based on a number of signals that are collected during their session.  
`genuine` indicates the user has high familiarity with the data they are providing, and that fraud is unlikely.  
`neutral` indicates some signals are present in between `risky` and `genuine`, but there are not enough clear signals to determine an outcome.  
`risky` indicates the user has low familiarity with the data they are providing, and that fraud is likely.  
`no_data` indicates there is not sufficient information to give an accurate signal.  
  

Possible values: `genuine`, `neutral`, `risky`, `no_data`

[`fraud_ring_detected`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-behavior-fraud-ring-detected)

stringstring

Field describing the outcome of a fraud ring behavior risk check.  
`yes` indicates that fraud ring activity was detected.  
`no` indicates that fraud ring activity was not detected.  
`no_data` indicates there was not enough information available to give an accurate signal.  
  

Possible values: `yes`, `no`, `no_data`

[`bot_detected`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-behavior-bot-detected)

stringstring

Field describing the outcome of a bot detection behavior risk check.  
`yes` indicates that automated activity was detected.  
`no` indicates that automated activity was not detected.  
`no_data` indicates there was not enough information available to give an accurate signal.  
  

Possible values: `yes`, `no`, `no_data`

[`email`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email)

nullableobjectnullable, object

Result summary object specifying values for `email` attributes of risk check.

[`is_deliverable`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-is-deliverable)

stringstring

SMTP-MX check to confirm the email address exists if known.  
  

Possible values: `yes`, `no`, `no_data`

[`breach_count`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-breach-count)

nullableintegernullable, integer

Count of all known breaches of this email address if known.

[`first_breached_at`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-first-breached-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`last_breached_at`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-last-breached-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`domain_registered_at`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-domain-registered-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`domain_is_free_provider`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-domain-is-free-provider)

stringstring

Indicates whether the email address domain is a free provider such as Gmail or Hotmail if known.  
  

Possible values: `yes`, `no`, `no_data`

[`domain_is_custom`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-domain-is-custom)

stringstring

Indicates whether the email address domain is custom if known, i.e. a company domain and not free or disposable.  
  

Possible values: `yes`, `no`, `no_data`

[`domain_is_disposable`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-domain-is-disposable)

stringstring

Indicates whether the email domain is listed as disposable if known. Disposable domains are often used to create email addresses that are part of a fake set of user details.  
  

Possible values: `yes`, `no`, `no_data`

[`top_level_domain_is_suspicious`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-top-level-domain-is-suspicious)

stringstring

Indicates whether the email address top level domain, which is the last part of the domain, is fraudulent or risky if known. In most cases, a suspicious top level domain is also associated with a disposable or high-risk domain.  
  

Possible values: `yes`, `no`, `no_data`

[`linked_services`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-email-linked-services)

[string][string]

A list of online services where this email address has been detected to have accounts or other activity.  
  

Possible values: `aboutme`, `adobe`, `adult_sites`, `airbnb`, `altbalaji`, `amazon`, `apple`, `archiveorg`, `atlassian`, `bitmoji`, `bodybuilding`, `booking`, `bukalapak`, `codecademy`, `deliveroo`, `diigo`, `discord`, `disneyplus`, `duolingo`, `ebay`, `envato`, `eventbrite`, `evernote`, `facebook`, `firefox`, `flickr`, `flipkart`, `foursquare`, `freelancer`, `gaana`, `giphy`, `github`, `google`, `gravatar`, `hubspot`, `imgur`, `instagram`, `jdid`, `kakao`, `kommo`, `komoot`, `lastfm`, `lazada`, `line`, `linkedin`, `mailru`, `microsoft`, `myspace`, `netflix`, `nike`, `ok`, `patreon`, `pinterest`, `plurk`, `quora`, `qzone`, `rambler`, `rappi`, `replit`, `samsung`, `seoclerks`, `shopclues`, `skype`, `snapchat`, `snapdeal`, `soundcloud`, `spotify`, `starz`, `strava`, `taringa`, `telegram`, `tiki`, `tokopedia`, `treehouse`, `tumblr`, `twitter`, `venmo`, `viber`, `vimeo`, `vivino`, `vkontakte`, `wattpad`, `weibo`, `whatsapp`, `wordpress`, `xing`, `yahoo`, `yandex`, `zalo`, `zoho`

[`phone`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-phone)

nullableobjectnullable, object

Result summary object specifying values for `phone` attributes of risk check.

[`linked_services`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-phone-linked-services)

[string][string]

A list of online services where this phone number has been detected to have accounts or other activity.  
  

Possible values: `aboutme`, `adobe`, `adult_sites`, `airbnb`, `altbalaji`, `amazon`, `apple`, `archiveorg`, `atlassian`, `bitmoji`, `bodybuilding`, `booking`, `bukalapak`, `codecademy`, `deliveroo`, `diigo`, `discord`, `disneyplus`, `duolingo`, `ebay`, `envato`, `eventbrite`, `evernote`, `facebook`, `firefox`, `flickr`, `flipkart`, `foursquare`, `freelancer`, `gaana`, `giphy`, `github`, `google`, `gravatar`, `hubspot`, `imgur`, `instagram`, `jdid`, `kakao`, `kommo`, `komoot`, `lastfm`, `lazada`, `line`, `linkedin`, `mailru`, `microsoft`, `myspace`, `netflix`, `nike`, `ok`, `patreon`, `pinterest`, `plurk`, `quora`, `qzone`, `rambler`, `rappi`, `replit`, `samsung`, `seoclerks`, `shopclues`, `skype`, `snapchat`, `snapdeal`, `soundcloud`, `spotify`, `starz`, `strava`, `taringa`, `telegram`, `tiki`, `tokopedia`, `treehouse`, `tumblr`, `twitter`, `venmo`, `viber`, `vimeo`, `vivino`, `vkontakte`, `wattpad`, `weibo`, `whatsapp`, `wordpress`, `xing`, `yahoo`, `yandex`, `zalo`, `zoho`

[`devices`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-devices)

[object][object]

Array of result summary objects specifying values for `device` attributes of risk check.

[`ip_proxy_type`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-devices-ip-proxy-type)

nullablestringnullable, string

An enum indicating whether a network proxy is present and if so what type it is.  
`none_detected` indicates the user is not on a detectable proxy network.  
`tor` indicates the user was using a Tor browser, which sends encrypted traffic on a decentralized network and is somewhat similar to a VPN (Virtual Private Network).  
`vpn` indicates the user is on a VPN (Virtual Private Network)  
`web_proxy` indicates the user is on a web proxy server, which may allow them to conceal information such as their IP address or other identifying information.  
`public_proxy` indicates the user is on a public web proxy server, which is similar to a web proxy but can be shared by multiple users. This may allow multiple users to appear as if they have the same IP address for instance.  
  

Possible values: `none_detected`, `tor`, `vpn`, `web_proxy`, `public_proxy`

[`ip_spam_list_count`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-devices-ip-spam-list-count)

nullableintegernullable, integer

Count of spam lists the IP address is associated with if known.

[`ip_timezone_offset`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-devices-ip-timezone-offset)

nullablestringnullable, string

UTC offset of the timezone associated with the IP address.

[`identity_abuse_signals`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-identity-abuse-signals)

nullableobjectnullable, object

Result summary object capturing abuse signals related to `identity abuse`, e.g. stolen and synthetic identity fraud. These attributes are only available for US identities and some signals may not be available depending on what information was collected.

[`synthetic_identity`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-identity-abuse-signals-synthetic-identity)

nullableobjectnullable, object

Field containing the data used in determining the outcome of the synthetic identity risk check.  
Contains the following fields:  
`score` - A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.

[`score`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-identity-abuse-signals-synthetic-identity-score)

integerinteger

A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.

[`stolen_identity`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-identity-abuse-signals-stolen-identity)

nullableobjectnullable, object

Field containing the data used in determining the outcome of the stolen identity risk check.  
Contains the following fields:  
`score` - A score from 0 to 100 indicating the likelihood that the user is a stolen identity.

[`score`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-identity-abuse-signals-stolen-identity-score)

integerinteger

A score from 0 to 100 indicating the likelihood that the user is a stolen identity.

[`facial_duplicates`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-facial-duplicates)

[object][object]

The attributes related to the facial duplicates captured in risk check.

[`id`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-facial-duplicates-id)

stringstring

ID of the associated Identity Verification attempt.

[`similarity`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-facial-duplicates-similarity)

integerinteger

Similarity score of the match. Ranges from 0 to 100.

[`matched_after_completed`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-facial-duplicates-matched-after-completed)

booleanboolean

Whether this match occurred after the session was complete. For example, this would be `true` if a later session ended up matching this one.

[`trust_index_score`](/docs/api/products/identity-verification/#identity_verification-get-response-risk-check-trust-index-score)

nullableintegernullable, integer

The trust index score for the `risk_check` step.

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms)

nullableobjectnullable, object

Additional information for the `verify_sms` step.

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-status)

stringstring

The outcome status for the associated Identity Verification attempt's `verify_sms` step. This field will always have the same value as `steps.verify_sms`.  
  

Possible values: `success`, `failed`

[`verifications`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications)

[object][object]

An array where each entry represents a verification attempt for the `verify_sms` step. Each entry represents one user-submitted phone number. Phone number edits, and in some cases error handling due to edge cases like rate limiting, may generate additional verifications.

[`status`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications-status)

stringstring

The outcome status for the individual SMS verification.  
  

Possible values: `pending`, `success`, `failed`, `canceled`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications-attempt)

integerinteger

The attempt field begins with 1 and increments with each subsequent SMS verification.

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`delivery_attempt_count`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications-delivery-attempt-count)

integerinteger

The number of delivery attempts made within the verification to send the SMS code to the user. Each delivery attempt represents the user taking action from the front end UI to request creation and delivery of a new SMS verification code, or to resend an existing SMS verification code. There is a limit of 3 delivery attempts per verification.

[`solve_attempt_count`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications-solve-attempt-count)

integerinteger

The number of attempts made by the user within the verification to verify the SMS code by entering it into the front end UI. There is a limit of 3 solve attempts per verification.

[`initially_sent_at`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications-initially-sent-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`last_sent_at`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications-last-sent-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-get-response-verify-sms-verifications-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`watchlist_screening_id`](/docs/api/products/identity-verification/#identity_verification-get-response-watchlist-screening-id)

nullablestringnullable, string

ID of the associated screening.

[`beacon_user_id`](/docs/api/products/identity-verification/#identity_verification-get-response-beacon-user-id)

nullablestringnullable, string

ID of the associated Beacon User.

[`user_id`](/docs/api/products/identity-verification/#identity_verification-get-response-user-id)

nullablestringnullable, string

Unique user identifier, created by calling `/user/create`. Either a `user_id` or the `client_user_id` must be provided. The `user_id` may only be used instead of the `client_user_id` if you were not a pre-existing user of `/user/create` as of December 10, 2025; for more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis). If both this field and a `client_user_id` are present in a request, the `user_id` must have been created from the provided `client_user_id`.

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-get-response-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`latest_scored_protect_event`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event)

nullableobjectnullable, object

Information about a Protect event including Trust Index score and fraud attributes.

[`event_id`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-event-id)

stringstring

The event ID.

[`timestamp`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-timestamp)

stringstring

The timestamp of the event, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2017-09-14T14:42:19.350Z"`  
  

Format: `date-time`

[`trust_index`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-trust-index)

nullableobjectnullable, object

Represents a calculate Trust Index Score.

[`score`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-trust-index-score)

integerinteger

The overall trust index score.

[`model`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-trust-index-model)

stringstring

The versioned name of the Trust Index model used for scoring.

[`subscores`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-trust-index-subscores)

nullableobjectnullable, object

Contains sub-score metadata.

[`device_and_connection`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-trust-index-subscores-device-and-connection)

nullableobjectnullable, object

Represents Trust Index Subscore.

[`score`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-trust-index-subscores-device-and-connection-score)

integerinteger

The subscore score.

[`bank_account_insights`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-trust-index-subscores-bank-account-insights)

nullableobjectnullable, object

Represents Trust Index Subscore.

[`score`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-trust-index-subscores-bank-account-insights-score)

integerinteger

The subscore score.

[`fraud_attributes`](/docs/api/products/identity-verification/#identity_verification-get-response-latest-scored-protect-event-fraud-attributes)

nullableobjectnullable, object

Event fraud attributes as an arbitrary set of key-value pairs.

[`request_id`](/docs/api/products/identity-verification/#identity_verification-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "idv_52xR9LKo77r1Np",
  "client_user_id": "your-db-id-3b24110",
  "created_at": "2020-07-24T03:26:02Z",
  "completed_at": "2020-07-24T03:26:02Z",
  "previous_attempt_id": "idv_42cF1MNo42r9Xj",
  "shareable_url": "https://flow.plaid.com/verify/idv_4FrXJvfQU3zGUR?key=e004115db797f7cc3083bff3167cba30644ef630fb46f5b086cde6cc3b86a36f",
  "template": {
    "id": "idvtmp_4FrXJvfQU3zGUR",
    "version": 2
  },
  "user": {
    "phone_number": "+12345678909",
    "date_of_birth": "1990-05-29",
    "ip_address": "192.0.2.42",
    "email_address": "user@example.com",
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
    "id_number": {
      "value": "123456789",
      "type": "us_ssn"
    }
  },
  "status": "success",
  "steps": {
    "accept_tos": "success",
    "verify_sms": "success",
    "kyc_check": "success",
    "documentary_verification": "success",
    "selfie_check": "success",
    "watchlist_screening": "success",
    "risk_check": "success"
  },
  "documentary_verification": {
    "status": "success",
    "documents": [
      {
        "status": "success",
        "attempt": 1,
        "images": {
          "original_front": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/original_front.jpeg",
          "original_back": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/original_back.jpeg",
          "cropped_front": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/cropped_front.jpeg",
          "cropped_back": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/cropped_back.jpeg",
          "face": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/face.jpeg"
        },
        "extracted_data": {
          "id_number": "AB123456",
          "category": "drivers_license",
          "expiration_date": "2030-05-29",
          "issue_date": "2020-05-29",
          "issuing_country": "US",
          "issuing_region": "IN",
          "date_of_birth": "1990-05-29",
          "address": {
            "street": "123 Main St. Unit 42",
            "city": "Pawnee",
            "region": "IN",
            "postal_code": "46001",
            "country": "US"
          },
          "name": {
            "given_name": "Leslie",
            "family_name": "Knope"
          }
        },
        "analysis": {
          "authenticity": "match",
          "image_quality": "high",
          "extracted_data": {
            "name": "match",
            "date_of_birth": "match",
            "expiration_date": "not_expired",
            "issuing_country": "match"
          },
          "aamva_verification": {
            "is_verified": true,
            "id_number": "match",
            "id_issue_date": "match",
            "id_expiration_date": "match",
            "street": "match",
            "city": "match",
            "postal_code": "match",
            "date_of_birth": "match",
            "gender": "match",
            "height": "match",
            "eye_color": "match",
            "first_name": "match",
            "middle_name": "match",
            "last_name": "match"
          }
        },
        "redacted_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "selfie_check": {
    "status": "success",
    "selfies": [
      {
        "status": "success",
        "attempt": 1,
        "capture": {
          "image_url": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/selfie/liveness.jpeg",
          "video_url": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/selfie/liveness.webm"
        },
        "analysis": {
          "document_comparison": "match",
          "liveness_check": "success"
        }
      }
    ]
  },
  "kyc_check": {
    "status": "success",
    "address": {
      "summary": "match",
      "po_box": "yes",
      "type": "residential"
    },
    "name": {
      "summary": "match"
    },
    "date_of_birth": {
      "summary": "match"
    },
    "id_number": {
      "summary": "match"
    },
    "phone_number": {
      "summary": "match",
      "area_code": "match"
    }
  },
  "risk_check": {
    "status": "success",
    "behavior": {
      "user_interactions": "risky",
      "fraud_ring_detected": "yes",
      "bot_detected": "yes"
    },
    "email": {
      "is_deliverable": "yes",
      "breach_count": 1,
      "first_breached_at": "1990-05-29",
      "last_breached_at": "1990-05-29",
      "domain_registered_at": "1990-05-29",
      "domain_is_free_provider": "yes",
      "domain_is_custom": "yes",
      "domain_is_disposable": "yes",
      "top_level_domain_is_suspicious": "yes",
      "linked_services": [
        "apple"
      ]
    },
    "phone": {
      "linked_services": [
        "apple"
      ]
    },
    "devices": [
      {
        "ip_proxy_type": "none_detected",
        "ip_spam_list_count": 1,
        "ip_timezone_offset": "+06:00:00"
      }
    ],
    "identity_abuse_signals": {
      "synthetic_identity": {
        "score": 0
      },
      "stolen_identity": {
        "score": 0
      }
    },
    "facial_duplicates": [
      {
        "id": "idv_52xR9LKo77r1Np",
        "similarity": 95,
        "matched_after_completed": true
      }
    ],
    "trust_index_score": 86
  },
  "verify_sms": {
    "status": "success",
    "verifications": [
      {
        "status": "success",
        "attempt": 1,
        "phone_number": "+12345678909",
        "delivery_attempt_count": 1,
        "solve_attempt_count": 1,
        "initially_sent_at": "2020-07-24T03:26:02Z",
        "last_sent_at": "2020-07-24T03:26:02Z",
        "redacted_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "watchlist_screening_id": "scr_52xR9LKo77r1Np",
  "beacon_user_id": "becusr_42cF1MNo42r9Xj",
  "user_id": "usr_dddAs9ewdcDQQQ",
  "redacted_at": "2020-07-24T03:26:02Z",
  "latest_scored_protect_event": {
    "event_id": "ptevt_7AJYTMFxRUgJ",
    "timestamp": "2020-07-24T03:26:02Z",
    "trust_index": {
      "score": 86,
      "model": "trust_index.2.0.0",
      "subscores": {
        "device_and_connection": {
          "score": 87
        },
        "bank_account_insights": {
          "score": 85
        }
      }
    },
    "fraud_attributes": {
      "fraud_attributes": {
        "link_session.linked_bank_accounts.user_pi_matches_owners": true,
        "link_session.linked_bank_accounts.connected_apps.days_since_first_connection": 582,
        "session.challenged_with_mfa": false,
        "user.bank_accounts.num_of_frozen_or_restricted_accounts": 0,
        "user.linked_bank_accounts.num_family_names": 1,
        "user.linked_bank_accounts.num_of_connected_banks": 1,
        "user.link_sessions.days_since_first_link_session": 150,
        "user.pi.email.history_yrs": 7.03,
        "user.pi.email.num_social_networks_linked": 12,
        "user.pi.ssn.user_likely_has_better_ssn": false
      }
    }
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/identity_verification/list`

#### List Identity Verifications

Filter and list Identity Verifications created by your account

/identity\_verification/list

**Request fields**

[`secret`](/docs/api/products/identity-verification/#identity_verification-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/identity-verification/#identity_verification-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`template_id`](/docs/api/products/identity-verification/#identity_verification-list-request-template-id)

requiredstringrequired, string

ID of the associated Identity Verification template. Like all Plaid identifiers, this is case-sensitive.

[`client_user_id`](/docs/api/products/identity-verification/#identity_verification-list-request-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`user_id`](/docs/api/products/identity-verification/#identity_verification-list-request-user-id)

stringstring

A unique user identifier, created by calling `/user/create`. Either a `user_id` or the `client_user_id` must be provided. The `user_id` may only be used instead of the `client_user_id` if you were not a pre-existing user of `/user/create` as of December 10, 2025; for more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis). If both this field and the `client_user_id` are present in the request, the `user_id` must have been created from the provided `client_user_id`.

[`cursor`](/docs/api/products/identity-verification/#identity_verification-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/identity\_verification/list

```
const request: ListIdentityVerificationRequest = {
  template_id: 'idvtmp_52xR9LKo77r1Np',
  client_user_id: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
};
try {
  const response = await plaidClient.identityVerificationList(request);
} catch (error) {
  // handle error
}
```

/identity\_verification/list

**Response fields**

[`identity_verifications`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications)

[object][object]

List of Plaid sessions

[`id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-id)

stringstring

ID of the associated Identity Verification attempt.

[`client_user_id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`created_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`completed_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-completed-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`previous_attempt_id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-previous-attempt-id)

nullablestringnullable, string

The ID for the Identity Verification preceding this session. This field will only be filled if the current Identity Verification is a retry of a previous attempt.

[`shareable_url`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-shareable-url)

nullablestringnullable, string

A shareable URL that can be sent directly to the user to complete verification

[`template`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-template)

objectobject

The resource ID and version number of the template configuring the behavior of a given Identity Verification.

[`id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-template-id)

stringstring

ID of the associated Identity Verification template. Like all Plaid identifiers, this is case-sensitive.

[`version`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-template-version)

integerinteger

Version of the associated Identity Verification template.

[`user`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user)

objectobject

The identity data that was either collected from the user or provided via API in order to perform an Identity Verification.

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-phone-number)

nullablestringnullable, string

A valid phone number in E.164 format.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`ip_address`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-ip-address)

nullablestringnullable, string

An IPv4 or IPV6 address.

[`email_address`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-name)

nullableobjectnullable, object

The full name provided by the user. If the user has not submitted their name, this field will be null. Otherwise, both fields are guaranteed to be filled.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-address)

nullableobjectnullable, object

Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include postal code

[`street`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-address-street)

nullablestringnullable, string

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-address-street2)

nullablestringnullable, string

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-address-city)

nullablestringnullable, string

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-address-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-address-postal-code)

nullablestringnullable, string

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-id-number)

nullableobjectnullable, object

ID number submitted by the user, currently used only for the Identity Verification product. If the user has not submitted this data yet, this field will be `null`. Otherwise, both fields are guaranteed to be filled.

[`value`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-id-number-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-id-number-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-status)

stringstring

The status of this Identity Verification attempt.  
`active` - The Identity Verification attempt is incomplete. The user may have completed part of the session, but has neither failed or passed.  
`success` - The Identity Verification attempt has completed, passing all steps defined to the associated Identity Verification template  
`failed` - The user failed one or more steps in the session and was told to contact support.  
`expired` - The Identity Verification attempt was active for a long period of time without being completed and was automatically marked as expired. Note that sessions currently do not expire. Automatic expiration is expected to be enabled in the future.  
`canceled` - The Identity Verification attempt was canceled, either via the dashboard by a user, or via API. The user may have completed part of the session, but has neither failed or passed.  
`pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  
  

Possible values: `active`, `success`, `failed`, `expired`, `canceled`, `pending_review`

[`steps`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-steps)

objectobject

Each step will be one of the following values:  
`active` - This step is the user's current step. They are either in the process of completing this step, or they recently closed their Identity Verification attempt while in the middle of this step. Only one step will be marked as `active` at any given point.  
`success` - The Identity Verification attempt has completed this step.  
`failed` - The user failed this step. This can either call the user to fail the session as a whole, or cause them to fallback to another step depending on how the Identity Verification template is configured. A failed step does not imply a failed session.  
`waiting_for_prerequisite` - The user needs to complete another step first, before they progress to this step. This step may never run, depending on if the user fails an earlier step or if the step is only run as a fallback.  
`not_applicable` - This step will not be run for this session.  
`skipped` - The retry instructions that created this Identity Verification attempt specified that this step should be skipped.  
`expired` - This step had not yet been completed when the Identity Verification attempt as a whole expired.  
`canceled` - The Identity Verification attempt was canceled before the user completed this step.  
`pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  
`manually_approved` - The step was manually overridden to pass by a team member in the dashboard.  
`manually_rejected` - The step was manually overridden to fail by a team member in the dashboard.

[`accept_tos`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-steps-accept-tos)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-steps-verify-sms)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-steps-kyc-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-steps-documentary-verification)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-steps-selfie-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`watchlist_screening`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-steps-watchlist-screening)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`risk_check`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-steps-risk-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification)

nullableobjectnullable, object

Data, images, analysis, and results from the `documentary_verification` step. This field will be `null` unless `steps.documentary_verification` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-status)

stringstring

The outcome status for the associated Identity Verification attempt's `documentary_verification` step. This field will always have the same value as `steps.documentary_verification`.

[`documents`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents)

[object][object]

An array of documents submitted to the `documentary_verification` step. Each entry represents one user submission, where each submission will contain both a front and back image, or just a front image, depending on the document type.  
Note: Plaid will automatically let a user submit a new set of document images up to three times if we detect that a previous attempt might have failed due to user error. For example, if the first set of document images are blurry or obscured by glare, the user will be asked to capture their documents again, resulting in at least two separate entries within `documents`. If the overall `documentary_verification` is `failed`, the user has exhausted their retry attempts.

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-status)

stringstring

An outcome status for this specific document submission. Distinct from the overall `documentary_verification.status` that summarizes the verification outcome from one or more documents.  
  

Possible values: `success`, `failed`, `manually_approved`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-attempt)

integerinteger

The `attempt` field begins with 1 and increments with each subsequent document upload.

[`images`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-images)

objectobject

URLs for downloading original and cropped images for this document submission. The URLs are designed to only allow downloading, not hot linking, so the URL will only serve the document image for 60 seconds before expiring. The expiration time is 60 seconds after the `GET` request for the associated Identity Verification attempt. A new expiring URL is generated with each request, so you can always rerequest the Identity Verification attempt if one of your URLs expires.

[`original_front`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-images-original-front)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading the uncropped original image of the front of the document.

[`original_back`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-images-original-back)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading the original image of the back of the document. Might be null if the back of the document was not collected.

[`cropped_front`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-images-cropped-front)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a cropped image containing just the front of the document.

[`cropped_back`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-images-cropped-back)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a cropped image containing just the back of the document. Might be null if the back of the document was not collected.

[`face`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-images-face)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a crop of just the user's face from the document image. Might be null if the document does not contain a face photo.

[`extracted_data`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data)

nullableobjectnullable, object

Data extracted from a user-submitted document.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-id-number)

nullablestringnullable, string

Alpha-numeric ID number extracted via OCR from the user's document image.

[`category`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-category)

stringstring

The type of identity document detected in the images provided. Will always be one of the following values:  
 `drivers_license` - A driver's license issued by the associated country, establishing identity without any guarantee as to citizenship, and granting driving privileges  
 `id_card` - A general national identification card, distinct from driver's licenses as it only establishes identity  
 `passport` - A travel passport issued by the associated country for one of its citizens  
 `residence_permit_card` - An identity document issued by the associated country permitting a foreign citizen to *temporarily* reside there  
 `resident_card` - An identity document issued by the associated country permitting a foreign citizen to *permanently* reside there  
 `visa` - An identity document issued by the associated country permitting a foreign citizen entry for a short duration and for a specific purpose, typically no longer than 6 months  
Note: This value may be different from the ID type that the user selects within Link. For example, if they select "Driver's License" but then submit a picture of a passport, this field will say `passport`  
  

Possible values: `drivers_license`, `id_card`, `passport`, `residence_permit_card`, `resident_card`, `visa`

[`expiration_date`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-expiration-date)

nullablestringnullable, string

The expiration date of the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`issue_date`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-issue-date)

nullablestringnullable, string

The issue date of the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`issuing_country`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-issuing-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`issuing_region`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-issuing-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-date-of-birth)

nullablestringnullable, string

A date extracted from the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`address`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-address)

nullableobjectnullable, object

The address extracted from the document. The address must at least contain the following fields to be a valid address: `street`, `city`, `country`. If any are missing or unable to be extracted, the address will be null.  
`region`, and `postal_code` may be null based on the addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include postal code  
Note: Optical Character Recognition (OCR) technology may sometimes extract incorrect data from a document.

[`street`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-address-street)

stringstring

The full street address extracted from the document.

[`city`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-address-city)

stringstring

City extracted from the document.

[`region`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-address-region)

nullablestringnullable, string

A subdivision code extracted from the document. Related terms would be "state", "province", "prefecture", "zone", "subdivision", etc. For a full list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they can be inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-address-postal-code)

nullablestringnullable, string

The postal code extracted from the document. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country extracted from the document. Must be in ISO 3166-1 alpha-2 form.

[`name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-name)

nullableobjectnullable, object

The individual's name extracted from the document.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-extracted-data-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`analysis`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis)

objectobject

High level descriptions of how the associated document was processed. If a document fails verification, the details in the `analysis` object should help clarify why the document was rejected.

[`authenticity`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-authenticity)

stringstring

High level summary of whether the document in the provided image matches the formatting rules and security checks for the associated jurisdiction.  
For example, most identity documents have formatting rules like the following:  
The image of the person's face must have a certain contrast in order to highlight skin tone  
The subject in the document's image must remove eye glasses and pose in a certain way  
The informational fields (name, date of birth, ID number, etc.) must be colored and aligned according to specific rules  
Security features like watermarks and background patterns must be present  
So a `match` status for this field indicates that the document in the provided image seems to conform to the various formatting and security rules associated with the detected document.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`image_quality`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-image-quality)

stringstring

A high level description of the quality of the image the user submitted.  
For example, an image that is blurry, distorted by glare from a nearby light source, or improperly framed might be marked as low or medium quality. Poor quality images are more likely to fail OCR and/or template conformity checks.  
Note: By default, Plaid will let a user recapture document images twice before failing the entire session if we attribute the failure to low image quality.  
  

Possible values: `high`, `medium`, `low`

[`extracted_data`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-extracted-data)

nullableobjectnullable, object

Analysis of the data extracted from the submitted document.

[`name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-extracted-data-name)

stringstring

A match summary describing the cross comparison between the subject's name, extracted from the document image, and the name they separately provided to identity verification attempt.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-extracted-data-date-of-birth)

stringstring

A match summary describing the cross comparison between the subject's date of birth, extracted from the document image, and the date of birth they separately provided to the identity verification attempt.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`expiration_date`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-extracted-data-expiration-date)

stringstring

A description of whether the associated document was expired when the verification was performed.  
Note: In the case where an expiration date is not present on the document or failed to be extracted, this value will be `no_data`.  
  

Possible values: `not_expired`, `expired`, `no_data`

[`issuing_country`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-extracted-data-issuing-country)

stringstring

A binary match indicator specifying whether the country that issued the provided document matches the country that the user separately provided to Plaid.  
Note: You can configure whether a `no_match` on `issuing_country` fails the `documentary_verification` by editing your Plaid Template.  
  

Possible values: `match`, `no_match`

[`aamva_verification`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification)

nullableobjectnullable, object

Analyzed AAMVA data for the associated hit.  
Note: This field is only available for U.S. driver's licenses issued by participating states.

[`is_verified`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-is-verified)

booleanboolean

The overall outcome of checking the associated hit against the issuing state database.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-id-number)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`id_issue_date`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-id-issue-date)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`id_expiration_date`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-id-expiration-date)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`street`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-street)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`city`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-city)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-postal-code)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-date-of-birth)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`gender`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-gender)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`height`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-height)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`eye_color`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-eye-color)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`first_name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-first-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`middle_name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-middle-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`last_name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-analysis-aamva-verification-last-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-documentary-verification-documents-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check)

nullableobjectnullable, object

Additional information for the `selfie_check` step. This field will be `null` unless `steps.selfie_check` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-status)

stringstring

The outcome status for the associated Identity Verification attempt's `selfie_check` step. This field will always have the same value as `steps.selfie_check`.  
  

Possible values: `success`, `failed`

[`selfies`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies)

[object][object]

An array of selfies submitted to the `selfie_check` step. Each entry represents one user submission.

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies-status)

stringstring

An outcome status for this specific selfie. Distinct from the overall `selfie_check.status` that summarizes the verification outcome from one or more selfies.  
  

Possible values: `success`, `failed`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies-attempt)

integerinteger

The `attempt` field begins with 1 and increments with each subsequent selfie upload.

[`capture`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies-capture)

objectobject

The image or video capture of a selfie. Only one of image or video URL will be populated per selfie.

[`image_url`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies-capture-image-url)

nullablestringnullable, string

Temporary URL for downloading an image selfie capture.

[`video_url`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies-capture-video-url)

nullablestringnullable, string

Temporary URL for downloading a video selfie capture.

[`analysis`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies-analysis)

objectobject

High level descriptions of how the associated selfie was processed. If a selfie fails verification, the details in the `analysis` object should help clarify why the selfie was rejected.

[`document_comparison`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies-analysis-document-comparison)

stringstring

Information about the comparison between the selfie and the document (if documentary verification also ran).  
  

Possible values: `match`, `no_match`, `no_input`

[`liveness_check`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-selfie-check-selfies-analysis-liveness-check)

stringstring

Assessment of whether the selfie capture is of a real human being, as opposed to a picture of a human on a screen, a picture of a paper cut out, someone wearing a mask, or a deepfake.  
  

Possible values: `success`, `failed`

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check)

nullableobjectnullable, object

Additional information for the `kyc_check` (Data Source Verification) step. This field will be `null` unless `steps.kyc_check` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-status)

stringstring

The outcome status for the associated Identity Verification attempt's `kyc_check` step. This field will always have the same value as `steps.kyc_check`.

[`address`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-address)

objectobject

Result summary object specifying how the `address` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-address-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`po_box`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-address-po-box)

stringstring

Field describing whether the associated address is a post office box. Will be `yes` when a P.O. box is detected, `no` when Plaid confirmed the address is not a P.O. box, and `no_data` when Plaid was not able to determine if the address is a P.O. box.  
  

Possible values: `yes`, `no`, `no_data`

[`type`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-address-type)

stringstring

Field describing whether the associated address is being used for commercial or residential purposes.  
Note: This value will be `no_data` when Plaid does not have sufficient data to determine the address's use.  
  

Possible values: `residential`, `commercial`, `no_data`

[`name`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-name)

objectobject

Result summary object specifying how the `name` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-name-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-date-of-birth)

objectobject

Result summary object specifying how the `date_of_birth` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-date-of-birth-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`id_number`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-id-number)

objectobject

Result summary object specifying how the `id_number` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-id-number-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-phone-number)

objectobject

Result summary object specifying how the `phone` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-phone-number-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`area_code`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-kyc-check-phone-number-area-code)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`risk_check`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check)

nullableobjectnullable, object

Additional information for the `risk_check` step.

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-status)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`behavior`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-behavior)

nullableobjectnullable, object

Result summary object specifying values for `behavior` attributes of risk check, when available.

[`user_interactions`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-behavior-user-interactions)

stringstring

Field describing the overall user interaction signals of a behavior risk check. This value represents how familiar the user is with the personal data they provide, based on a number of signals that are collected during their session.  
`genuine` indicates the user has high familiarity with the data they are providing, and that fraud is unlikely.  
`neutral` indicates some signals are present in between `risky` and `genuine`, but there are not enough clear signals to determine an outcome.  
`risky` indicates the user has low familiarity with the data they are providing, and that fraud is likely.  
`no_data` indicates there is not sufficient information to give an accurate signal.  
  

Possible values: `genuine`, `neutral`, `risky`, `no_data`

[`fraud_ring_detected`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-behavior-fraud-ring-detected)

stringstring

Field describing the outcome of a fraud ring behavior risk check.  
`yes` indicates that fraud ring activity was detected.  
`no` indicates that fraud ring activity was not detected.  
`no_data` indicates there was not enough information available to give an accurate signal.  
  

Possible values: `yes`, `no`, `no_data`

[`bot_detected`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-behavior-bot-detected)

stringstring

Field describing the outcome of a bot detection behavior risk check.  
`yes` indicates that automated activity was detected.  
`no` indicates that automated activity was not detected.  
`no_data` indicates there was not enough information available to give an accurate signal.  
  

Possible values: `yes`, `no`, `no_data`

[`email`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email)

nullableobjectnullable, object

Result summary object specifying values for `email` attributes of risk check.

[`is_deliverable`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-is-deliverable)

stringstring

SMTP-MX check to confirm the email address exists if known.  
  

Possible values: `yes`, `no`, `no_data`

[`breach_count`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-breach-count)

nullableintegernullable, integer

Count of all known breaches of this email address if known.

[`first_breached_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-first-breached-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`last_breached_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-last-breached-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`domain_registered_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-domain-registered-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`domain_is_free_provider`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-domain-is-free-provider)

stringstring

Indicates whether the email address domain is a free provider such as Gmail or Hotmail if known.  
  

Possible values: `yes`, `no`, `no_data`

[`domain_is_custom`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-domain-is-custom)

stringstring

Indicates whether the email address domain is custom if known, i.e. a company domain and not free or disposable.  
  

Possible values: `yes`, `no`, `no_data`

[`domain_is_disposable`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-domain-is-disposable)

stringstring

Indicates whether the email domain is listed as disposable if known. Disposable domains are often used to create email addresses that are part of a fake set of user details.  
  

Possible values: `yes`, `no`, `no_data`

[`top_level_domain_is_suspicious`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-top-level-domain-is-suspicious)

stringstring

Indicates whether the email address top level domain, which is the last part of the domain, is fraudulent or risky if known. In most cases, a suspicious top level domain is also associated with a disposable or high-risk domain.  
  

Possible values: `yes`, `no`, `no_data`

[`linked_services`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-email-linked-services)

[string][string]

A list of online services where this email address has been detected to have accounts or other activity.  
  

Possible values: `aboutme`, `adobe`, `adult_sites`, `airbnb`, `altbalaji`, `amazon`, `apple`, `archiveorg`, `atlassian`, `bitmoji`, `bodybuilding`, `booking`, `bukalapak`, `codecademy`, `deliveroo`, `diigo`, `discord`, `disneyplus`, `duolingo`, `ebay`, `envato`, `eventbrite`, `evernote`, `facebook`, `firefox`, `flickr`, `flipkart`, `foursquare`, `freelancer`, `gaana`, `giphy`, `github`, `google`, `gravatar`, `hubspot`, `imgur`, `instagram`, `jdid`, `kakao`, `kommo`, `komoot`, `lastfm`, `lazada`, `line`, `linkedin`, `mailru`, `microsoft`, `myspace`, `netflix`, `nike`, `ok`, `patreon`, `pinterest`, `plurk`, `quora`, `qzone`, `rambler`, `rappi`, `replit`, `samsung`, `seoclerks`, `shopclues`, `skype`, `snapchat`, `snapdeal`, `soundcloud`, `spotify`, `starz`, `strava`, `taringa`, `telegram`, `tiki`, `tokopedia`, `treehouse`, `tumblr`, `twitter`, `venmo`, `viber`, `vimeo`, `vivino`, `vkontakte`, `wattpad`, `weibo`, `whatsapp`, `wordpress`, `xing`, `yahoo`, `yandex`, `zalo`, `zoho`

[`phone`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-phone)

nullableobjectnullable, object

Result summary object specifying values for `phone` attributes of risk check.

[`linked_services`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-phone-linked-services)

[string][string]

A list of online services where this phone number has been detected to have accounts or other activity.  
  

Possible values: `aboutme`, `adobe`, `adult_sites`, `airbnb`, `altbalaji`, `amazon`, `apple`, `archiveorg`, `atlassian`, `bitmoji`, `bodybuilding`, `booking`, `bukalapak`, `codecademy`, `deliveroo`, `diigo`, `discord`, `disneyplus`, `duolingo`, `ebay`, `envato`, `eventbrite`, `evernote`, `facebook`, `firefox`, `flickr`, `flipkart`, `foursquare`, `freelancer`, `gaana`, `giphy`, `github`, `google`, `gravatar`, `hubspot`, `imgur`, `instagram`, `jdid`, `kakao`, `kommo`, `komoot`, `lastfm`, `lazada`, `line`, `linkedin`, `mailru`, `microsoft`, `myspace`, `netflix`, `nike`, `ok`, `patreon`, `pinterest`, `plurk`, `quora`, `qzone`, `rambler`, `rappi`, `replit`, `samsung`, `seoclerks`, `shopclues`, `skype`, `snapchat`, `snapdeal`, `soundcloud`, `spotify`, `starz`, `strava`, `taringa`, `telegram`, `tiki`, `tokopedia`, `treehouse`, `tumblr`, `twitter`, `venmo`, `viber`, `vimeo`, `vivino`, `vkontakte`, `wattpad`, `weibo`, `whatsapp`, `wordpress`, `xing`, `yahoo`, `yandex`, `zalo`, `zoho`

[`devices`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-devices)

[object][object]

Array of result summary objects specifying values for `device` attributes of risk check.

[`ip_proxy_type`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-devices-ip-proxy-type)

nullablestringnullable, string

An enum indicating whether a network proxy is present and if so what type it is.  
`none_detected` indicates the user is not on a detectable proxy network.  
`tor` indicates the user was using a Tor browser, which sends encrypted traffic on a decentralized network and is somewhat similar to a VPN (Virtual Private Network).  
`vpn` indicates the user is on a VPN (Virtual Private Network)  
`web_proxy` indicates the user is on a web proxy server, which may allow them to conceal information such as their IP address or other identifying information.  
`public_proxy` indicates the user is on a public web proxy server, which is similar to a web proxy but can be shared by multiple users. This may allow multiple users to appear as if they have the same IP address for instance.  
  

Possible values: `none_detected`, `tor`, `vpn`, `web_proxy`, `public_proxy`

[`ip_spam_list_count`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-devices-ip-spam-list-count)

nullableintegernullable, integer

Count of spam lists the IP address is associated with if known.

[`ip_timezone_offset`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-devices-ip-timezone-offset)

nullablestringnullable, string

UTC offset of the timezone associated with the IP address.

[`identity_abuse_signals`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-identity-abuse-signals)

nullableobjectnullable, object

Result summary object capturing abuse signals related to `identity abuse`, e.g. stolen and synthetic identity fraud. These attributes are only available for US identities and some signals may not be available depending on what information was collected.

[`synthetic_identity`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-identity-abuse-signals-synthetic-identity)

nullableobjectnullable, object

Field containing the data used in determining the outcome of the synthetic identity risk check.  
Contains the following fields:  
`score` - A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.

[`score`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-identity-abuse-signals-synthetic-identity-score)

integerinteger

A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.

[`stolen_identity`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-identity-abuse-signals-stolen-identity)

nullableobjectnullable, object

Field containing the data used in determining the outcome of the stolen identity risk check.  
Contains the following fields:  
`score` - A score from 0 to 100 indicating the likelihood that the user is a stolen identity.

[`score`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-identity-abuse-signals-stolen-identity-score)

integerinteger

A score from 0 to 100 indicating the likelihood that the user is a stolen identity.

[`facial_duplicates`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-facial-duplicates)

[object][object]

The attributes related to the facial duplicates captured in risk check.

[`id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-facial-duplicates-id)

stringstring

ID of the associated Identity Verification attempt.

[`similarity`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-facial-duplicates-similarity)

integerinteger

Similarity score of the match. Ranges from 0 to 100.

[`matched_after_completed`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-facial-duplicates-matched-after-completed)

booleanboolean

Whether this match occurred after the session was complete. For example, this would be `true` if a later session ended up matching this one.

[`trust_index_score`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-risk-check-trust-index-score)

nullableintegernullable, integer

The trust index score for the `risk_check` step.

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms)

nullableobjectnullable, object

Additional information for the `verify_sms` step.

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-status)

stringstring

The outcome status for the associated Identity Verification attempt's `verify_sms` step. This field will always have the same value as `steps.verify_sms`.  
  

Possible values: `success`, `failed`

[`verifications`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications)

[object][object]

An array where each entry represents a verification attempt for the `verify_sms` step. Each entry represents one user-submitted phone number. Phone number edits, and in some cases error handling due to edge cases like rate limiting, may generate additional verifications.

[`status`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications-status)

stringstring

The outcome status for the individual SMS verification.  
  

Possible values: `pending`, `success`, `failed`, `canceled`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications-attempt)

integerinteger

The attempt field begins with 1 and increments with each subsequent SMS verification.

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`delivery_attempt_count`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications-delivery-attempt-count)

integerinteger

The number of delivery attempts made within the verification to send the SMS code to the user. Each delivery attempt represents the user taking action from the front end UI to request creation and delivery of a new SMS verification code, or to resend an existing SMS verification code. There is a limit of 3 delivery attempts per verification.

[`solve_attempt_count`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications-solve-attempt-count)

integerinteger

The number of attempts made by the user within the verification to verify the SMS code by entering it into the front end UI. There is a limit of 3 solve attempts per verification.

[`initially_sent_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications-initially-sent-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`last_sent_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications-last-sent-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-verify-sms-verifications-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`watchlist_screening_id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-watchlist-screening-id)

nullablestringnullable, string

ID of the associated screening.

[`beacon_user_id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-beacon-user-id)

nullablestringnullable, string

ID of the associated Beacon User.

[`user_id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-user-id)

nullablestringnullable, string

Unique user identifier, created by calling `/user/create`. Either a `user_id` or the `client_user_id` must be provided. The `user_id` may only be used instead of the `client_user_id` if you were not a pre-existing user of `/user/create` as of December 10, 2025; for more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis). If both this field and a `client_user_id` are present in a request, the `user_id` must have been created from the provided `client_user_id`.

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`latest_scored_protect_event`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event)

nullableobjectnullable, object

Information about a Protect event including Trust Index score and fraud attributes.

[`event_id`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-event-id)

stringstring

The event ID.

[`timestamp`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-timestamp)

stringstring

The timestamp of the event, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2017-09-14T14:42:19.350Z"`  
  

Format: `date-time`

[`trust_index`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-trust-index)

nullableobjectnullable, object

Represents a calculate Trust Index Score.

[`score`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-trust-index-score)

integerinteger

The overall trust index score.

[`model`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-trust-index-model)

stringstring

The versioned name of the Trust Index model used for scoring.

[`subscores`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-trust-index-subscores)

nullableobjectnullable, object

Contains sub-score metadata.

[`device_and_connection`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-trust-index-subscores-device-and-connection)

nullableobjectnullable, object

Represents Trust Index Subscore.

[`score`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-trust-index-subscores-device-and-connection-score)

integerinteger

The subscore score.

[`bank_account_insights`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-trust-index-subscores-bank-account-insights)

nullableobjectnullable, object

Represents Trust Index Subscore.

[`score`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-trust-index-subscores-bank-account-insights-score)

integerinteger

The subscore score.

[`fraud_attributes`](/docs/api/products/identity-verification/#identity_verification-list-response-identity-verifications-latest-scored-protect-event-fraud-attributes)

nullableobjectnullable, object

Event fraud attributes as an arbitrary set of key-value pairs.

[`next_cursor`](/docs/api/products/identity-verification/#identity_verification-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/identity-verification/#identity_verification-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "identity_verifications": [
    {
      "id": "idv_52xR9LKo77r1Np",
      "client_user_id": "your-db-id-3b24110",
      "created_at": "2020-07-24T03:26:02Z",
      "completed_at": "2020-07-24T03:26:02Z",
      "previous_attempt_id": "idv_42cF1MNo42r9Xj",
      "shareable_url": "https://flow.plaid.com/verify/idv_4FrXJvfQU3zGUR?key=e004115db797f7cc3083bff3167cba30644ef630fb46f5b086cde6cc3b86a36f",
      "template": {
        "id": "idvtmp_4FrXJvfQU3zGUR",
        "version": 2
      },
      "user": {
        "phone_number": "+12345678909",
        "date_of_birth": "1990-05-29",
        "ip_address": "192.0.2.42",
        "email_address": "user@example.com",
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
        "id_number": {
          "value": "123456789",
          "type": "us_ssn"
        }
      },
      "status": "success",
      "steps": {
        "accept_tos": "success",
        "verify_sms": "success",
        "kyc_check": "success",
        "documentary_verification": "success",
        "selfie_check": "success",
        "watchlist_screening": "success",
        "risk_check": "success"
      },
      "documentary_verification": {
        "status": "success",
        "documents": [
          {
            "status": "success",
            "attempt": 1,
            "images": {
              "original_front": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/original_front.jpeg",
              "original_back": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/original_back.jpeg",
              "cropped_front": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/cropped_front.jpeg",
              "cropped_back": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/cropped_back.jpeg",
              "face": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/face.jpeg"
            },
            "extracted_data": {
              "id_number": "AB123456",
              "category": "drivers_license",
              "expiration_date": "2030-05-29",
              "issue_date": "2020-05-29",
              "issuing_country": "US",
              "issuing_region": "IN",
              "date_of_birth": "1990-05-29",
              "address": {
                "street": "123 Main St. Unit 42",
                "city": "Pawnee",
                "region": "IN",
                "postal_code": "46001",
                "country": "US"
              },
              "name": {
                "given_name": "Leslie",
                "family_name": "Knope"
              }
            },
            "analysis": {
              "authenticity": "match",
              "image_quality": "high",
              "extracted_data": {
                "name": "match",
                "date_of_birth": "match",
                "expiration_date": "not_expired",
                "issuing_country": "match"
              },
              "aamva_verification": {
                "is_verified": true,
                "id_number": "match",
                "id_issue_date": "match",
                "id_expiration_date": "match",
                "street": "match",
                "city": "match",
                "postal_code": "match",
                "date_of_birth": "match",
                "gender": "match",
                "height": "match",
                "eye_color": "match",
                "first_name": "match",
                "middle_name": "match",
                "last_name": "match"
              }
            },
            "redacted_at": "2020-07-24T03:26:02Z"
          }
        ]
      },
      "selfie_check": {
        "status": "success",
        "selfies": [
          {
            "status": "success",
            "attempt": 1,
            "capture": {
              "image_url": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/selfie/liveness.jpeg",
              "video_url": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/selfie/liveness.webm"
            },
            "analysis": {
              "document_comparison": "match",
              "liveness_check": "success"
            }
          }
        ]
      },
      "kyc_check": {
        "status": "success",
        "address": {
          "summary": "match",
          "po_box": "yes",
          "type": "residential"
        },
        "name": {
          "summary": "match"
        },
        "date_of_birth": {
          "summary": "match"
        },
        "id_number": {
          "summary": "match"
        },
        "phone_number": {
          "summary": "match",
          "area_code": "match"
        }
      },
      "risk_check": {
        "status": "success",
        "behavior": {
          "user_interactions": "risky",
          "fraud_ring_detected": "yes",
          "bot_detected": "yes"
        },
        "email": {
          "is_deliverable": "yes",
          "breach_count": 1,
          "first_breached_at": "1990-05-29",
          "last_breached_at": "1990-05-29",
          "domain_registered_at": "1990-05-29",
          "domain_is_free_provider": "yes",
          "domain_is_custom": "yes",
          "domain_is_disposable": "yes",
          "top_level_domain_is_suspicious": "yes",
          "linked_services": [
            "apple"
          ]
        },
        "phone": {
          "linked_services": [
            "apple"
          ]
        },
        "devices": [
          {
            "ip_proxy_type": "none_detected",
            "ip_spam_list_count": 1,
            "ip_timezone_offset": "+06:00:00"
          }
        ],
        "identity_abuse_signals": {
          "synthetic_identity": {
            "score": 0
          },
          "stolen_identity": {
            "score": 0
          }
        },
        "facial_duplicates": [
          {
            "id": "idv_52xR9LKo77r1Np",
            "similarity": 95,
            "matched_after_completed": true
          }
        ],
        "trust_index_score": 86
      },
      "verify_sms": {
        "status": "success",
        "verifications": [
          {
            "status": "success",
            "attempt": 1,
            "phone_number": "+12345678909",
            "delivery_attempt_count": 1,
            "solve_attempt_count": 1,
            "initially_sent_at": "2020-07-24T03:26:02Z",
            "last_sent_at": "2020-07-24T03:26:02Z",
            "redacted_at": "2020-07-24T03:26:02Z"
          }
        ]
      },
      "watchlist_screening_id": "scr_52xR9LKo77r1Np",
      "beacon_user_id": "becusr_42cF1MNo42r9Xj",
      "user_id": "usr_dddAs9ewdcDQQQ",
      "redacted_at": "2020-07-24T03:26:02Z",
      "latest_scored_protect_event": {
        "event_id": "ptevt_7AJYTMFxRUgJ",
        "timestamp": "2020-07-24T03:26:02Z",
        "trust_index": {
          "score": 86,
          "model": "trust_index.2.0.0",
          "subscores": {
            "device_and_connection": {
              "score": 87
            },
            "bank_account_insights": {
              "score": 85
            }
          }
        },
        "fraud_attributes": {
          "fraud_attributes": {
            "link_session.linked_bank_accounts.user_pi_matches_owners": true,
            "link_session.linked_bank_accounts.connected_apps.days_since_first_connection": 582,
            "session.challenged_with_mfa": false,
            "user.bank_accounts.num_of_frozen_or_restricted_accounts": 0,
            "user.linked_bank_accounts.num_family_names": 1,
            "user.linked_bank_accounts.num_of_connected_banks": 1,
            "user.link_sessions.days_since_first_link_session": 150,
            "user.pi.email.history_yrs": 7.03,
            "user.pi.email.num_social_networks_linked": 12,
            "user.pi.ssn.user_likely_has_better_ssn": false
          }
        }
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/identity_verification/retry`

#### Retry an Identity Verification

Allow a customer to retry their Identity Verification

/identity\_verification/retry

**Request fields**

[`client_user_id`](/docs/api/products/identity-verification/#identity_verification-retry-request-client-user-id)

requiredstringrequired, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`template_id`](/docs/api/products/identity-verification/#identity_verification-retry-request-template-id)

requiredstringrequired, string

ID of the associated Identity Verification template. Like all Plaid identifiers, this is case-sensitive.

[`strategy`](/docs/api/products/identity-verification/#identity_verification-retry-request-strategy)

requiredstringrequired, string

An instruction specifying what steps the new Identity Verification attempt should require the user to complete:  
`reset` - Restart the user at the beginning of the session, regardless of whether they successfully completed part of their previous session.  
`incomplete` - Start the new session at the step that the user failed in the previous session, skipping steps that have already been successfully completed.  
`infer` - If the most recent Identity Verification attempt associated with the given `client_user_id` has a status of `failed` or `expired`, retry using the `incomplete` strategy. Otherwise, use the `reset` strategy.  
`custom` - Start the new session with a custom configuration, specified by the value of the `steps` field  
Note:  
The `incomplete` strategy cannot be applied if the session's failing step is `screening` or `risk_check`.  
The `infer` strategy cannot be applied if the session's status is still `active`  
  

Possible values: `reset`, `incomplete`, `infer`, `custom`

[`user`](/docs/api/products/identity-verification/#identity_verification-retry-request-user)

objectobject

User information collected outside of Link, most likely via your own onboarding process.  
Each of the following identity fields are optional:  
`email_address`  
`phone_number`  
`date_of_birth`  
`name`  
`address`  
`id_number`  
Specifically, these fields are optional in that they can either be fully provided (satisfying every required field in their subschema) or omitted from the request entirely by not providing the key or value.
Providing these fields via the API will result in Link skipping the data collection process for the associated user. All verification steps enabled in the associated Identity Verification Template will still be run. Verification steps will either be run immediately, or once the user completes the `accept_tos` step, depending on the value provided to the `gave_consent` field.

[`email_address`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-phone-number)

stringstring

A valid phone number in E.164 format.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`name`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-name)

objectobject

You can use this field to pre-populate the user's legal name; if it is provided here, they will not be prompted to enter their name in the identity verification attempt.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-name-given-name)

requiredstringrequired, string

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-name-family-name)

requiredstringrequired, string

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-address)

objectobject

Home address for the user. Supported values are: not provided, address with only country code or full address.  
For more context on this field, see [Input Validation by Country](https://plaid.com/docs/identity-verification/hybrid-input-validation/#input-validation-by-country).

[`street`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-address-street)

stringstring

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-address-street2)

stringstring

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-address-city)

stringstring

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-address-region)

stringstring

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-address-postal-code)

stringstring

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-address-country)

requiredstringrequired, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-id-number)

objectobject

ID number submitted by the user, currently used only for the Identity Verification product. If the user has not submitted this data yet, this field will be `null`. Otherwise, both fields are guaranteed to be filled.

[`value`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-id-number-value)

requiredstringrequired, string

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/identity-verification/#identity_verification-retry-request-user-id-number-type)

requiredstringrequired, string

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`steps`](/docs/api/products/identity-verification/#identity_verification-retry-request-steps)

objectobject

Instructions for the `custom` retry strategy specifying which steps should be required or skipped.  
Note:  
This field must be provided when the retry strategy is `custom` and must be omitted otherwise.  
Custom retries override settings in your Plaid Template. For example, if your Plaid Template has `verify_sms` disabled, a custom retry with `verify_sms` enabled will still require the step.  
The `selfie_check` step is currently not supported on the sandbox server. Sandbox requests will silently disable the `selfie_check` step when provided.

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-retry-request-steps-verify-sms)

requiredbooleanrequired, boolean

A boolean field specifying whether the new session should require or skip the `verify_sms` step.

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-retry-request-steps-kyc-check)

requiredbooleanrequired, boolean

A boolean field specifying whether the new session should require or skip the `kyc_check` (Data Source Verification) step.

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-retry-request-steps-documentary-verification)

requiredbooleanrequired, boolean

A boolean field specifying whether the new session should require or skip the `documentary_verification` step.

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-retry-request-steps-selfie-check)

requiredbooleanrequired, boolean

A boolean field specifying whether the new session should require or skip the `selfie_check` step. If a previous session has already passed the `selfie_check` step, the new selfie check will be a Selfie Reauthentication check, in which the selfie is tested for liveness and for consistency with the previous selfie.

[`is_shareable`](/docs/api/products/identity-verification/#identity_verification-retry-request-is-shareable)

booleanboolean

A flag specifying whether you would like Plaid to expose a shareable URL for the verification being retried. If a value for this flag is not specified, the `is_shareable` setting from the original verification attempt will be used.

[`client_id`](/docs/api/products/identity-verification/#identity_verification-retry-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/identity-verification/#identity_verification-retry-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/identity\_verification/retry

```
const request: IdentityVerificationRetryRequest = {
  client_user_id: 'user-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d',
  template_id: 'idvtmp_52xR9LKo77r1Np',
  strategy: 'reset',
  user: {
    email_address: 'acharleston@email.com',
    phone_number: '+12345678909',
    date_of_birth: '1975-01-18',
    name: {
      given_name: 'Anna',
      family_name: 'Charleston',
    },
    address: {
      street: '100 Market Street',
      street2: 'Apt 1A',
      city: 'San Francisco',
      region: 'CA',
      postal_code: '94103',
      country: 'US',
    },
    id_number: {
      value: '123456789',
      type: 'us_ssn',
    },
  },
};
try {
  const response = await plaidClient.identityVerificationRetry(request);
} catch (error) {
  // handle error
}
```

/identity\_verification/retry

**Response fields**

[`id`](/docs/api/products/identity-verification/#identity_verification-retry-response-id)

stringstring

ID of the associated Identity Verification attempt.

[`client_user_id`](/docs/api/products/identity-verification/#identity_verification-retry-response-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`created_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`completed_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-completed-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`previous_attempt_id`](/docs/api/products/identity-verification/#identity_verification-retry-response-previous-attempt-id)

nullablestringnullable, string

The ID for the Identity Verification preceding this session. This field will only be filled if the current Identity Verification is a retry of a previous attempt.

[`shareable_url`](/docs/api/products/identity-verification/#identity_verification-retry-response-shareable-url)

nullablestringnullable, string

A shareable URL that can be sent directly to the user to complete verification

[`template`](/docs/api/products/identity-verification/#identity_verification-retry-response-template)

objectobject

The resource ID and version number of the template configuring the behavior of a given Identity Verification.

[`id`](/docs/api/products/identity-verification/#identity_verification-retry-response-template-id)

stringstring

ID of the associated Identity Verification template. Like all Plaid identifiers, this is case-sensitive.

[`version`](/docs/api/products/identity-verification/#identity_verification-retry-response-template-version)

integerinteger

Version of the associated Identity Verification template.

[`user`](/docs/api/products/identity-verification/#identity_verification-retry-response-user)

objectobject

The identity data that was either collected from the user or provided via API in order to perform an Identity Verification.

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-phone-number)

nullablestringnullable, string

A valid phone number in E.164 format.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`ip_address`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-ip-address)

nullablestringnullable, string

An IPv4 or IPV6 address.

[`email_address`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`name`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-name)

nullableobjectnullable, object

The full name provided by the user. If the user has not submitted their name, this field will be null. Otherwise, both fields are guaranteed to be filled.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`address`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-address)

nullableobjectnullable, object

Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include postal code

[`street`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-address-street)

nullablestringnullable, string

The primary street portion of an address. If an address is provided, this field will always be filled. A string with at least one non-whitespace alphabetical character, with a max length of 80 characters.

[`street2`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-address-street2)

nullablestringnullable, string

Extra street information, like an apartment or suite number. If provided, a string with at least one non-whitespace character, with a max length of 50 characters.

[`city`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-address-city)

nullablestringnullable, string

City from the address. A string with at least one non-whitespace alphabetical character, with a max length of 100 characters.

[`region`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-address-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-address-postal-code)

nullablestringnullable, string

The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-id-number)

nullableobjectnullable, object

ID number submitted by the user, currently used only for the Identity Verification product. If the user has not submitted this data yet, this field will be `null`. Otherwise, both fields are guaranteed to be filled.

[`value`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-id-number-value)

stringstring

Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped. For specific format requirements by ID type, see [Input Validation Rules](https://plaid.com/docs/identity-verification/hybrid-input-validation/#id-numbers).

[`type`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-id-number-type)

stringstring

A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Input Validation Rules](https://cognitohq.com/docs/flow/flow-hybrid-input-validation#id-numbers).  
  

Possible values: `ar_dni`, `au_drivers_license`, `au_passport`, `br_cpf`, `ca_sin`, `cl_run`, `cn_resident_card`, `co_nit`, `dk_cpr`, `eg_national_id`, `es_dni`, `es_nie`, `hk_hkid`, `in_pan`, `it_cf`, `jo_civil_id`, `jp_my_number`, `ke_huduma_namba`, `kw_civil_id`, `mx_curp`, `mx_rfc`, `my_nric`, `ng_nin`, `nz_drivers_license`, `om_civil_id`, `ph_psn`, `pl_pesel`, `ro_cnp`, `sa_national_id`, `se_pin`, `sg_nric`, `tr_tc_kimlik`, `us_ssn`, `us_ssn_last_4`, `za_smart_id`

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-status)

stringstring

The status of this Identity Verification attempt.  
`active` - The Identity Verification attempt is incomplete. The user may have completed part of the session, but has neither failed or passed.  
`success` - The Identity Verification attempt has completed, passing all steps defined to the associated Identity Verification template  
`failed` - The user failed one or more steps in the session and was told to contact support.  
`expired` - The Identity Verification attempt was active for a long period of time without being completed and was automatically marked as expired. Note that sessions currently do not expire. Automatic expiration is expected to be enabled in the future.  
`canceled` - The Identity Verification attempt was canceled, either via the dashboard by a user, or via API. The user may have completed part of the session, but has neither failed or passed.  
`pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  
  

Possible values: `active`, `success`, `failed`, `expired`, `canceled`, `pending_review`

[`steps`](/docs/api/products/identity-verification/#identity_verification-retry-response-steps)

objectobject

Each step will be one of the following values:  
`active` - This step is the user's current step. They are either in the process of completing this step, or they recently closed their Identity Verification attempt while in the middle of this step. Only one step will be marked as `active` at any given point.  
`success` - The Identity Verification attempt has completed this step.  
`failed` - The user failed this step. This can either call the user to fail the session as a whole, or cause them to fallback to another step depending on how the Identity Verification template is configured. A failed step does not imply a failed session.  
`waiting_for_prerequisite` - The user needs to complete another step first, before they progress to this step. This step may never run, depending on if the user fails an earlier step or if the step is only run as a fallback.  
`not_applicable` - This step will not be run for this session.  
`skipped` - The retry instructions that created this Identity Verification attempt specified that this step should be skipped.  
`expired` - This step had not yet been completed when the Identity Verification attempt as a whole expired.  
`canceled` - The Identity Verification attempt was canceled before the user completed this step.  
`pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  
`manually_approved` - The step was manually overridden to pass by a team member in the dashboard.  
`manually_rejected` - The step was manually overridden to fail by a team member in the dashboard.

[`accept_tos`](/docs/api/products/identity-verification/#identity_verification-retry-response-steps-accept-tos)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-retry-response-steps-verify-sms)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-retry-response-steps-kyc-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-retry-response-steps-documentary-verification)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-retry-response-steps-selfie-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`watchlist_screening`](/docs/api/products/identity-verification/#identity_verification-retry-response-steps-watchlist-screening)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`risk_check`](/docs/api/products/identity-verification/#identity_verification-retry-response-steps-risk-check)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`documentary_verification`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification)

nullableobjectnullable, object

Data, images, analysis, and results from the `documentary_verification` step. This field will be `null` unless `steps.documentary_verification` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-status)

stringstring

The outcome status for the associated Identity Verification attempt's `documentary_verification` step. This field will always have the same value as `steps.documentary_verification`.

[`documents`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents)

[object][object]

An array of documents submitted to the `documentary_verification` step. Each entry represents one user submission, where each submission will contain both a front and back image, or just a front image, depending on the document type.  
Note: Plaid will automatically let a user submit a new set of document images up to three times if we detect that a previous attempt might have failed due to user error. For example, if the first set of document images are blurry or obscured by glare, the user will be asked to capture their documents again, resulting in at least two separate entries within `documents`. If the overall `documentary_verification` is `failed`, the user has exhausted their retry attempts.

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-status)

stringstring

An outcome status for this specific document submission. Distinct from the overall `documentary_verification.status` that summarizes the verification outcome from one or more documents.  
  

Possible values: `success`, `failed`, `manually_approved`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-attempt)

integerinteger

The `attempt` field begins with 1 and increments with each subsequent document upload.

[`images`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-images)

objectobject

URLs for downloading original and cropped images for this document submission. The URLs are designed to only allow downloading, not hot linking, so the URL will only serve the document image for 60 seconds before expiring. The expiration time is 60 seconds after the `GET` request for the associated Identity Verification attempt. A new expiring URL is generated with each request, so you can always rerequest the Identity Verification attempt if one of your URLs expires.

[`original_front`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-images-original-front)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading the uncropped original image of the front of the document.

[`original_back`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-images-original-back)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading the original image of the back of the document. Might be null if the back of the document was not collected.

[`cropped_front`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-images-cropped-front)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a cropped image containing just the front of the document.

[`cropped_back`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-images-cropped-back)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a cropped image containing just the back of the document. Might be null if the back of the document was not collected.

[`face`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-images-face)

nullablestringnullable, string

Temporary URL that expires after 60 seconds for downloading a crop of just the user's face from the document image. Might be null if the document does not contain a face photo.

[`extracted_data`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data)

nullableobjectnullable, object

Data extracted from a user-submitted document.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-id-number)

nullablestringnullable, string

Alpha-numeric ID number extracted via OCR from the user's document image.

[`category`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-category)

stringstring

The type of identity document detected in the images provided. Will always be one of the following values:  
 `drivers_license` - A driver's license issued by the associated country, establishing identity without any guarantee as to citizenship, and granting driving privileges  
 `id_card` - A general national identification card, distinct from driver's licenses as it only establishes identity  
 `passport` - A travel passport issued by the associated country for one of its citizens  
 `residence_permit_card` - An identity document issued by the associated country permitting a foreign citizen to *temporarily* reside there  
 `resident_card` - An identity document issued by the associated country permitting a foreign citizen to *permanently* reside there  
 `visa` - An identity document issued by the associated country permitting a foreign citizen entry for a short duration and for a specific purpose, typically no longer than 6 months  
Note: This value may be different from the ID type that the user selects within Link. For example, if they select "Driver's License" but then submit a picture of a passport, this field will say `passport`  
  

Possible values: `drivers_license`, `id_card`, `passport`, `residence_permit_card`, `resident_card`, `visa`

[`expiration_date`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-expiration-date)

nullablestringnullable, string

The expiration date of the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`issue_date`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-issue-date)

nullablestringnullable, string

The issue date of the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`issuing_country`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-issuing-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`issuing_region`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-issuing-region)

nullablestringnullable, string

A subdivision code. "Subdivision" is a generic term for "state", "province", "prefecture", "zone", etc. For the list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they are inferred from the `country` field.

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-date-of-birth)

nullablestringnullable, string

A date extracted from the document in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`address`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-address)

nullableobjectnullable, object

The address extracted from the document. The address must at least contain the following fields to be a valid address: `street`, `city`, `country`. If any are missing or unable to be extracted, the address will be null.  
`region`, and `postal_code` may be null based on the addressing system. For example:  
Addresses from the United Kingdom will not include a region  
Addresses from Hong Kong will not include postal code  
Note: Optical Character Recognition (OCR) technology may sometimes extract incorrect data from a document.

[`street`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-address-street)

stringstring

The full street address extracted from the document.

[`city`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-address-city)

stringstring

City extracted from the document.

[`region`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-address-region)

nullablestringnullable, string

A subdivision code extracted from the document. Related terms would be "state", "province", "prefecture", "zone", "subdivision", etc. For a full list of valid codes, see [country subdivision codes](https://plaid.com/documents/country_subdivision_codes.json). Country prefixes are omitted, since they can be inferred from the `country` field.

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-address-postal-code)

nullablestringnullable, string

The postal code extracted from the document. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.

[`country`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-address-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country extracted from the document. Must be in ISO 3166-1 alpha-2 form.

[`name`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-name)

nullableobjectnullable, object

The individual's name extracted from the document.

[`given_name`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-name-given-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`family_name`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-extracted-data-name-family-name)

stringstring

A string with at least one non-whitespace character, with a max length of 100 characters.

[`analysis`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis)

objectobject

High level descriptions of how the associated document was processed. If a document fails verification, the details in the `analysis` object should help clarify why the document was rejected.

[`authenticity`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-authenticity)

stringstring

High level summary of whether the document in the provided image matches the formatting rules and security checks for the associated jurisdiction.  
For example, most identity documents have formatting rules like the following:  
The image of the person's face must have a certain contrast in order to highlight skin tone  
The subject in the document's image must remove eye glasses and pose in a certain way  
The informational fields (name, date of birth, ID number, etc.) must be colored and aligned according to specific rules  
Security features like watermarks and background patterns must be present  
So a `match` status for this field indicates that the document in the provided image seems to conform to the various formatting and security rules associated with the detected document.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`image_quality`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-image-quality)

stringstring

A high level description of the quality of the image the user submitted.  
For example, an image that is blurry, distorted by glare from a nearby light source, or improperly framed might be marked as low or medium quality. Poor quality images are more likely to fail OCR and/or template conformity checks.  
Note: By default, Plaid will let a user recapture document images twice before failing the entire session if we attribute the failure to low image quality.  
  

Possible values: `high`, `medium`, `low`

[`extracted_data`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-extracted-data)

nullableobjectnullable, object

Analysis of the data extracted from the submitted document.

[`name`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-extracted-data-name)

stringstring

A match summary describing the cross comparison between the subject's name, extracted from the document image, and the name they separately provided to identity verification attempt.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-extracted-data-date-of-birth)

stringstring

A match summary describing the cross comparison between the subject's date of birth, extracted from the document image, and the date of birth they separately provided to the identity verification attempt.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`expiration_date`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-extracted-data-expiration-date)

stringstring

A description of whether the associated document was expired when the verification was performed.  
Note: In the case where an expiration date is not present on the document or failed to be extracted, this value will be `no_data`.  
  

Possible values: `not_expired`, `expired`, `no_data`

[`issuing_country`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-extracted-data-issuing-country)

stringstring

A binary match indicator specifying whether the country that issued the provided document matches the country that the user separately provided to Plaid.  
Note: You can configure whether a `no_match` on `issuing_country` fails the `documentary_verification` by editing your Plaid Template.  
  

Possible values: `match`, `no_match`

[`aamva_verification`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification)

nullableobjectnullable, object

Analyzed AAMVA data for the associated hit.  
Note: This field is only available for U.S. driver's licenses issued by participating states.

[`is_verified`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-is-verified)

booleanboolean

The overall outcome of checking the associated hit against the issuing state database.

[`id_number`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-id-number)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`id_issue_date`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-id-issue-date)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`id_expiration_date`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-id-expiration-date)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`street`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-street)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`city`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-city)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`postal_code`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-postal-code)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-date-of-birth)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`gender`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-gender)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`height`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-height)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`eye_color`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-eye-color)

stringstring

The outcome of checking the particular field against state databases.  
`match` - The field is an exact match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `no_match`, `no_data`

[`first_name`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-first-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`middle_name`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-middle-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`last_name`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-analysis-aamva-verification-last-name)

stringstring

The outcome of checking the associated hit against state databases.  
`match` - The field is an exact match with the state database.  
`partial_match` - The field is a partial match with the state database.  
`no_match` - The field is not an exact match with the state database.  
`no_data` - The field was unable to be checked against state databases.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-documentary-verification-documents-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`selfie_check`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check)

nullableobjectnullable, object

Additional information for the `selfie_check` step. This field will be `null` unless `steps.selfie_check` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-status)

stringstring

The outcome status for the associated Identity Verification attempt's `selfie_check` step. This field will always have the same value as `steps.selfie_check`.  
  

Possible values: `success`, `failed`

[`selfies`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies)

[object][object]

An array of selfies submitted to the `selfie_check` step. Each entry represents one user submission.

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies-status)

stringstring

An outcome status for this specific selfie. Distinct from the overall `selfie_check.status` that summarizes the verification outcome from one or more selfies.  
  

Possible values: `success`, `failed`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies-attempt)

integerinteger

The `attempt` field begins with 1 and increments with each subsequent selfie upload.

[`capture`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies-capture)

objectobject

The image or video capture of a selfie. Only one of image or video URL will be populated per selfie.

[`image_url`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies-capture-image-url)

nullablestringnullable, string

Temporary URL for downloading an image selfie capture.

[`video_url`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies-capture-video-url)

nullablestringnullable, string

Temporary URL for downloading a video selfie capture.

[`analysis`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies-analysis)

objectobject

High level descriptions of how the associated selfie was processed. If a selfie fails verification, the details in the `analysis` object should help clarify why the selfie was rejected.

[`document_comparison`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies-analysis-document-comparison)

stringstring

Information about the comparison between the selfie and the document (if documentary verification also ran).  
  

Possible values: `match`, `no_match`, `no_input`

[`liveness_check`](/docs/api/products/identity-verification/#identity_verification-retry-response-selfie-check-selfies-analysis-liveness-check)

stringstring

Assessment of whether the selfie capture is of a real human being, as opposed to a picture of a human on a screen, a picture of a paper cut out, someone wearing a mask, or a deepfake.  
  

Possible values: `success`, `failed`

[`kyc_check`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check)

nullableobjectnullable, object

Additional information for the `kyc_check` (Data Source Verification) step. This field will be `null` unless `steps.kyc_check` has reached a terminal state of either `success` or `failed`.

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-status)

stringstring

The outcome status for the associated Identity Verification attempt's `kyc_check` step. This field will always have the same value as `steps.kyc_check`.

[`address`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-address)

objectobject

Result summary object specifying how the `address` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-address-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`po_box`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-address-po-box)

stringstring

Field describing whether the associated address is a post office box. Will be `yes` when a P.O. box is detected, `no` when Plaid confirmed the address is not a P.O. box, and `no_data` when Plaid was not able to determine if the address is a P.O. box.  
  

Possible values: `yes`, `no`, `no_data`

[`type`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-address-type)

stringstring

Field describing whether the associated address is being used for commercial or residential purposes.  
Note: This value will be `no_data` when Plaid does not have sufficient data to determine the address's use.  
  

Possible values: `residential`, `commercial`, `no_data`

[`name`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-name)

objectobject

Result summary object specifying how the `name` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-name-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`date_of_birth`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-date-of-birth)

objectobject

Result summary object specifying how the `date_of_birth` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-date-of-birth-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`id_number`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-id-number)

objectobject

Result summary object specifying how the `id_number` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-id-number-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-phone-number)

objectobject

Result summary object specifying how the `phone` field matched.

[`summary`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-phone-number-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`area_code`](/docs/api/products/identity-verification/#identity_verification-retry-response-kyc-check-phone-number-area-code)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`risk_check`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check)

nullableobjectnullable, object

Additional information for the `risk_check` step.

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-status)

stringstring

The status of a step in the Identity Verification process.  
  

Possible values: `success`, `active`, `failed`, `waiting_for_prerequisite`, `not_applicable`, `skipped`, `expired`, `canceled`, `pending_review`, `manually_approved`, `manually_rejected`

[`behavior`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-behavior)

nullableobjectnullable, object

Result summary object specifying values for `behavior` attributes of risk check, when available.

[`user_interactions`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-behavior-user-interactions)

stringstring

Field describing the overall user interaction signals of a behavior risk check. This value represents how familiar the user is with the personal data they provide, based on a number of signals that are collected during their session.  
`genuine` indicates the user has high familiarity with the data they are providing, and that fraud is unlikely.  
`neutral` indicates some signals are present in between `risky` and `genuine`, but there are not enough clear signals to determine an outcome.  
`risky` indicates the user has low familiarity with the data they are providing, and that fraud is likely.  
`no_data` indicates there is not sufficient information to give an accurate signal.  
  

Possible values: `genuine`, `neutral`, `risky`, `no_data`

[`fraud_ring_detected`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-behavior-fraud-ring-detected)

stringstring

Field describing the outcome of a fraud ring behavior risk check.  
`yes` indicates that fraud ring activity was detected.  
`no` indicates that fraud ring activity was not detected.  
`no_data` indicates there was not enough information available to give an accurate signal.  
  

Possible values: `yes`, `no`, `no_data`

[`bot_detected`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-behavior-bot-detected)

stringstring

Field describing the outcome of a bot detection behavior risk check.  
`yes` indicates that automated activity was detected.  
`no` indicates that automated activity was not detected.  
`no_data` indicates there was not enough information available to give an accurate signal.  
  

Possible values: `yes`, `no`, `no_data`

[`email`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email)

nullableobjectnullable, object

Result summary object specifying values for `email` attributes of risk check.

[`is_deliverable`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-is-deliverable)

stringstring

SMTP-MX check to confirm the email address exists if known.  
  

Possible values: `yes`, `no`, `no_data`

[`breach_count`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-breach-count)

nullableintegernullable, integer

Count of all known breaches of this email address if known.

[`first_breached_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-first-breached-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`last_breached_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-last-breached-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`domain_registered_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-domain-registered-at)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`domain_is_free_provider`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-domain-is-free-provider)

stringstring

Indicates whether the email address domain is a free provider such as Gmail or Hotmail if known.  
  

Possible values: `yes`, `no`, `no_data`

[`domain_is_custom`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-domain-is-custom)

stringstring

Indicates whether the email address domain is custom if known, i.e. a company domain and not free or disposable.  
  

Possible values: `yes`, `no`, `no_data`

[`domain_is_disposable`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-domain-is-disposable)

stringstring

Indicates whether the email domain is listed as disposable if known. Disposable domains are often used to create email addresses that are part of a fake set of user details.  
  

Possible values: `yes`, `no`, `no_data`

[`top_level_domain_is_suspicious`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-top-level-domain-is-suspicious)

stringstring

Indicates whether the email address top level domain, which is the last part of the domain, is fraudulent or risky if known. In most cases, a suspicious top level domain is also associated with a disposable or high-risk domain.  
  

Possible values: `yes`, `no`, `no_data`

[`linked_services`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-email-linked-services)

[string][string]

A list of online services where this email address has been detected to have accounts or other activity.  
  

Possible values: `aboutme`, `adobe`, `adult_sites`, `airbnb`, `altbalaji`, `amazon`, `apple`, `archiveorg`, `atlassian`, `bitmoji`, `bodybuilding`, `booking`, `bukalapak`, `codecademy`, `deliveroo`, `diigo`, `discord`, `disneyplus`, `duolingo`, `ebay`, `envato`, `eventbrite`, `evernote`, `facebook`, `firefox`, `flickr`, `flipkart`, `foursquare`, `freelancer`, `gaana`, `giphy`, `github`, `google`, `gravatar`, `hubspot`, `imgur`, `instagram`, `jdid`, `kakao`, `kommo`, `komoot`, `lastfm`, `lazada`, `line`, `linkedin`, `mailru`, `microsoft`, `myspace`, `netflix`, `nike`, `ok`, `patreon`, `pinterest`, `plurk`, `quora`, `qzone`, `rambler`, `rappi`, `replit`, `samsung`, `seoclerks`, `shopclues`, `skype`, `snapchat`, `snapdeal`, `soundcloud`, `spotify`, `starz`, `strava`, `taringa`, `telegram`, `tiki`, `tokopedia`, `treehouse`, `tumblr`, `twitter`, `venmo`, `viber`, `vimeo`, `vivino`, `vkontakte`, `wattpad`, `weibo`, `whatsapp`, `wordpress`, `xing`, `yahoo`, `yandex`, `zalo`, `zoho`

[`phone`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-phone)

nullableobjectnullable, object

Result summary object specifying values for `phone` attributes of risk check.

[`linked_services`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-phone-linked-services)

[string][string]

A list of online services where this phone number has been detected to have accounts or other activity.  
  

Possible values: `aboutme`, `adobe`, `adult_sites`, `airbnb`, `altbalaji`, `amazon`, `apple`, `archiveorg`, `atlassian`, `bitmoji`, `bodybuilding`, `booking`, `bukalapak`, `codecademy`, `deliveroo`, `diigo`, `discord`, `disneyplus`, `duolingo`, `ebay`, `envato`, `eventbrite`, `evernote`, `facebook`, `firefox`, `flickr`, `flipkart`, `foursquare`, `freelancer`, `gaana`, `giphy`, `github`, `google`, `gravatar`, `hubspot`, `imgur`, `instagram`, `jdid`, `kakao`, `kommo`, `komoot`, `lastfm`, `lazada`, `line`, `linkedin`, `mailru`, `microsoft`, `myspace`, `netflix`, `nike`, `ok`, `patreon`, `pinterest`, `plurk`, `quora`, `qzone`, `rambler`, `rappi`, `replit`, `samsung`, `seoclerks`, `shopclues`, `skype`, `snapchat`, `snapdeal`, `soundcloud`, `spotify`, `starz`, `strava`, `taringa`, `telegram`, `tiki`, `tokopedia`, `treehouse`, `tumblr`, `twitter`, `venmo`, `viber`, `vimeo`, `vivino`, `vkontakte`, `wattpad`, `weibo`, `whatsapp`, `wordpress`, `xing`, `yahoo`, `yandex`, `zalo`, `zoho`

[`devices`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-devices)

[object][object]

Array of result summary objects specifying values for `device` attributes of risk check.

[`ip_proxy_type`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-devices-ip-proxy-type)

nullablestringnullable, string

An enum indicating whether a network proxy is present and if so what type it is.  
`none_detected` indicates the user is not on a detectable proxy network.  
`tor` indicates the user was using a Tor browser, which sends encrypted traffic on a decentralized network and is somewhat similar to a VPN (Virtual Private Network).  
`vpn` indicates the user is on a VPN (Virtual Private Network)  
`web_proxy` indicates the user is on a web proxy server, which may allow them to conceal information such as their IP address or other identifying information.  
`public_proxy` indicates the user is on a public web proxy server, which is similar to a web proxy but can be shared by multiple users. This may allow multiple users to appear as if they have the same IP address for instance.  
  

Possible values: `none_detected`, `tor`, `vpn`, `web_proxy`, `public_proxy`

[`ip_spam_list_count`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-devices-ip-spam-list-count)

nullableintegernullable, integer

Count of spam lists the IP address is associated with if known.

[`ip_timezone_offset`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-devices-ip-timezone-offset)

nullablestringnullable, string

UTC offset of the timezone associated with the IP address.

[`identity_abuse_signals`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-identity-abuse-signals)

nullableobjectnullable, object

Result summary object capturing abuse signals related to `identity abuse`, e.g. stolen and synthetic identity fraud. These attributes are only available for US identities and some signals may not be available depending on what information was collected.

[`synthetic_identity`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-identity-abuse-signals-synthetic-identity)

nullableobjectnullable, object

Field containing the data used in determining the outcome of the synthetic identity risk check.  
Contains the following fields:  
`score` - A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.

[`score`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-identity-abuse-signals-synthetic-identity-score)

integerinteger

A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.

[`stolen_identity`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-identity-abuse-signals-stolen-identity)

nullableobjectnullable, object

Field containing the data used in determining the outcome of the stolen identity risk check.  
Contains the following fields:  
`score` - A score from 0 to 100 indicating the likelihood that the user is a stolen identity.

[`score`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-identity-abuse-signals-stolen-identity-score)

integerinteger

A score from 0 to 100 indicating the likelihood that the user is a stolen identity.

[`facial_duplicates`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-facial-duplicates)

[object][object]

The attributes related to the facial duplicates captured in risk check.

[`id`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-facial-duplicates-id)

stringstring

ID of the associated Identity Verification attempt.

[`similarity`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-facial-duplicates-similarity)

integerinteger

Similarity score of the match. Ranges from 0 to 100.

[`matched_after_completed`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-facial-duplicates-matched-after-completed)

booleanboolean

Whether this match occurred after the session was complete. For example, this would be `true` if a later session ended up matching this one.

[`trust_index_score`](/docs/api/products/identity-verification/#identity_verification-retry-response-risk-check-trust-index-score)

nullableintegernullable, integer

The trust index score for the `risk_check` step.

[`verify_sms`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms)

nullableobjectnullable, object

Additional information for the `verify_sms` step.

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-status)

stringstring

The outcome status for the associated Identity Verification attempt's `verify_sms` step. This field will always have the same value as `steps.verify_sms`.  
  

Possible values: `success`, `failed`

[`verifications`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications)

[object][object]

An array where each entry represents a verification attempt for the `verify_sms` step. Each entry represents one user-submitted phone number. Phone number edits, and in some cases error handling due to edge cases like rate limiting, may generate additional verifications.

[`status`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications-status)

stringstring

The outcome status for the individual SMS verification.  
  

Possible values: `pending`, `success`, `failed`, `canceled`

[`attempt`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications-attempt)

integerinteger

The attempt field begins with 1 and increments with each subsequent SMS verification.

[`phone_number`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`delivery_attempt_count`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications-delivery-attempt-count)

integerinteger

The number of delivery attempts made within the verification to send the SMS code to the user. Each delivery attempt represents the user taking action from the front end UI to request creation and delivery of a new SMS verification code, or to resend an existing SMS verification code. There is a limit of 3 delivery attempts per verification.

[`solve_attempt_count`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications-solve-attempt-count)

integerinteger

The number of attempts made by the user within the verification to verify the SMS code by entering it into the front end UI. There is a limit of 3 solve attempts per verification.

[`initially_sent_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications-initially-sent-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`last_sent_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications-last-sent-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-verify-sms-verifications-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`watchlist_screening_id`](/docs/api/products/identity-verification/#identity_verification-retry-response-watchlist-screening-id)

nullablestringnullable, string

ID of the associated screening.

[`beacon_user_id`](/docs/api/products/identity-verification/#identity_verification-retry-response-beacon-user-id)

nullablestringnullable, string

ID of the associated Beacon User.

[`user_id`](/docs/api/products/identity-verification/#identity_verification-retry-response-user-id)

nullablestringnullable, string

Unique user identifier, created by calling `/user/create`. Either a `user_id` or the `client_user_id` must be provided. The `user_id` may only be used instead of the `client_user_id` if you were not a pre-existing user of `/user/create` as of December 10, 2025; for more details, see [New User APIs](https://plaid.com/docs/api/users/user-apis). If both this field and a `client_user_id` are present in a request, the `user_id` must have been created from the provided `client_user_id`.

[`redacted_at`](/docs/api/products/identity-verification/#identity_verification-retry-response-redacted-at)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`latest_scored_protect_event`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event)

nullableobjectnullable, object

Information about a Protect event including Trust Index score and fraud attributes.

[`event_id`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-event-id)

stringstring

The event ID.

[`timestamp`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-timestamp)

stringstring

The timestamp of the event, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `"2017-09-14T14:42:19.350Z"`  
  

Format: `date-time`

[`trust_index`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-trust-index)

nullableobjectnullable, object

Represents a calculate Trust Index Score.

[`score`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-trust-index-score)

integerinteger

The overall trust index score.

[`model`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-trust-index-model)

stringstring

The versioned name of the Trust Index model used for scoring.

[`subscores`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-trust-index-subscores)

nullableobjectnullable, object

Contains sub-score metadata.

[`device_and_connection`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-trust-index-subscores-device-and-connection)

nullableobjectnullable, object

Represents Trust Index Subscore.

[`score`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-trust-index-subscores-device-and-connection-score)

integerinteger

The subscore score.

[`bank_account_insights`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-trust-index-subscores-bank-account-insights)

nullableobjectnullable, object

Represents Trust Index Subscore.

[`score`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-trust-index-subscores-bank-account-insights-score)

integerinteger

The subscore score.

[`fraud_attributes`](/docs/api/products/identity-verification/#identity_verification-retry-response-latest-scored-protect-event-fraud-attributes)

nullableobjectnullable, object

Event fraud attributes as an arbitrary set of key-value pairs.

[`request_id`](/docs/api/products/identity-verification/#identity_verification-retry-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "idv_52xR9LKo77r1Np",
  "client_user_id": "your-db-id-3b24110",
  "created_at": "2020-07-24T03:26:02Z",
  "completed_at": "2020-07-24T03:26:02Z",
  "previous_attempt_id": "idv_42cF1MNo42r9Xj",
  "shareable_url": "https://flow.plaid.com/verify/idv_4FrXJvfQU3zGUR?key=e004115db797f7cc3083bff3167cba30644ef630fb46f5b086cde6cc3b86a36f",
  "template": {
    "id": "idvtmp_4FrXJvfQU3zGUR",
    "version": 2
  },
  "user": {
    "phone_number": "+12345678909",
    "date_of_birth": "1990-05-29",
    "ip_address": "192.0.2.42",
    "email_address": "user@example.com",
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
    "id_number": {
      "value": "123456789",
      "type": "us_ssn"
    }
  },
  "status": "success",
  "steps": {
    "accept_tos": "success",
    "verify_sms": "success",
    "kyc_check": "success",
    "documentary_verification": "success",
    "selfie_check": "success",
    "watchlist_screening": "success",
    "risk_check": "success"
  },
  "documentary_verification": {
    "status": "success",
    "documents": [
      {
        "status": "success",
        "attempt": 1,
        "images": {
          "original_front": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/original_front.jpeg",
          "original_back": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/original_back.jpeg",
          "cropped_front": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/cropped_front.jpeg",
          "cropped_back": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/cropped_back.jpeg",
          "face": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/documents/1/face.jpeg"
        },
        "extracted_data": {
          "id_number": "AB123456",
          "category": "drivers_license",
          "expiration_date": "2030-05-29",
          "issue_date": "2020-05-29",
          "issuing_country": "US",
          "issuing_region": "IN",
          "date_of_birth": "1990-05-29",
          "address": {
            "street": "123 Main St. Unit 42",
            "city": "Pawnee",
            "region": "IN",
            "postal_code": "46001",
            "country": "US"
          },
          "name": {
            "given_name": "Leslie",
            "family_name": "Knope"
          }
        },
        "analysis": {
          "authenticity": "match",
          "image_quality": "high",
          "extracted_data": {
            "name": "match",
            "date_of_birth": "match",
            "expiration_date": "not_expired",
            "issuing_country": "match"
          },
          "aamva_verification": {
            "is_verified": true,
            "id_number": "match",
            "id_issue_date": "match",
            "id_expiration_date": "match",
            "street": "match",
            "city": "match",
            "postal_code": "match",
            "date_of_birth": "match",
            "gender": "match",
            "height": "match",
            "eye_color": "match",
            "first_name": "match",
            "middle_name": "match",
            "last_name": "match"
          }
        },
        "redacted_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "selfie_check": {
    "status": "success",
    "selfies": [
      {
        "status": "success",
        "attempt": 1,
        "capture": {
          "image_url": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/selfie/liveness.jpeg",
          "video_url": "https://example.plaid.com/verifications/idv_52xR9LKo77r1Np/selfie/liveness.webm"
        },
        "analysis": {
          "document_comparison": "match",
          "liveness_check": "success"
        }
      }
    ]
  },
  "kyc_check": {
    "status": "success",
    "address": {
      "summary": "match",
      "po_box": "yes",
      "type": "residential"
    },
    "name": {
      "summary": "match"
    },
    "date_of_birth": {
      "summary": "match"
    },
    "id_number": {
      "summary": "match"
    },
    "phone_number": {
      "summary": "match",
      "area_code": "match"
    }
  },
  "risk_check": {
    "status": "success",
    "behavior": {
      "user_interactions": "risky",
      "fraud_ring_detected": "yes",
      "bot_detected": "yes"
    },
    "email": {
      "is_deliverable": "yes",
      "breach_count": 1,
      "first_breached_at": "1990-05-29",
      "last_breached_at": "1990-05-29",
      "domain_registered_at": "1990-05-29",
      "domain_is_free_provider": "yes",
      "domain_is_custom": "yes",
      "domain_is_disposable": "yes",
      "top_level_domain_is_suspicious": "yes",
      "linked_services": [
        "apple"
      ]
    },
    "phone": {
      "linked_services": [
        "apple"
      ]
    },
    "devices": [
      {
        "ip_proxy_type": "none_detected",
        "ip_spam_list_count": 1,
        "ip_timezone_offset": "+06:00:00"
      }
    ],
    "identity_abuse_signals": {
      "synthetic_identity": {
        "score": 0
      },
      "stolen_identity": {
        "score": 0
      }
    },
    "facial_duplicates": [
      {
        "id": "idv_52xR9LKo77r1Np",
        "similarity": 95,
        "matched_after_completed": true
      }
    ],
    "trust_index_score": 86
  },
  "verify_sms": {
    "status": "success",
    "verifications": [
      {
        "status": "success",
        "attempt": 1,
        "phone_number": "+12345678909",
        "delivery_attempt_count": 1,
        "solve_attempt_count": 1,
        "initially_sent_at": "2020-07-24T03:26:02Z",
        "last_sent_at": "2020-07-24T03:26:02Z",
        "redacted_at": "2020-07-24T03:26:02Z"
      }
    ]
  },
  "watchlist_screening_id": "scr_52xR9LKo77r1Np",
  "beacon_user_id": "becusr_42cF1MNo42r9Xj",
  "user_id": "usr_dddAs9ewdcDQQQ",
  "redacted_at": "2020-07-24T03:26:02Z",
  "latest_scored_protect_event": {
    "event_id": "ptevt_7AJYTMFxRUgJ",
    "timestamp": "2020-07-24T03:26:02Z",
    "trust_index": {
      "score": 86,
      "model": "trust_index.2.0.0",
      "subscores": {
        "device_and_connection": {
          "score": 87
        },
        "bank_account_insights": {
          "score": 85
        }
      }
    },
    "fraud_attributes": {
      "fraud_attributes": {
        "link_session.linked_bank_accounts.user_pi_matches_owners": true,
        "link_session.linked_bank_accounts.connected_apps.days_since_first_connection": 582,
        "session.challenged_with_mfa": false,
        "user.bank_accounts.num_of_frozen_or_restricted_accounts": 0,
        "user.linked_bank_accounts.num_family_names": 1,
        "user.linked_bank_accounts.num_of_connected_banks": 1,
        "user.link_sessions.days_since_first_link_session": 150,
        "user.pi.email.history_yrs": 7.03,
        "user.pi.email.num_social_networks_linked": 12,
        "user.pi.ssn.user_likely_has_better_ssn": false
      }
    }
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

### Webhooks

=\*=\*=\*=

#### `STATUS_UPDATED`

Fired when the status of an identity verification has been updated, which can be triggered via the dashboard or the API.

**Properties**

[`webhook_type`](/docs/api/products/identity-verification/#IdentityVerificationStatusUpdatedWebhook-webhook-type)

stringstring

`IDENTITY_VERIFICATION`

[`webhook_code`](/docs/api/products/identity-verification/#IdentityVerificationStatusUpdatedWebhook-webhook-code)

stringstring

`STATUS_UPDATED`

[`identity_verification_id`](/docs/api/products/identity-verification/#IdentityVerificationStatusUpdatedWebhook-identity-verification-id)

stringstring

The ID of the associated Identity Verification attempt.

[`environment`](/docs/api/products/identity-verification/#IdentityVerificationStatusUpdatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "IDENTITY_VERIFICATION",
  "webhook_code": "STATUS_UPDATED",
  "identity_verification_id": "idv_52xR9LKo77r1Np",
  "environment": "production"
}
```

=\*=\*=\*=

#### `STEP_UPDATED`

Fired when an end user has completed a step of the Identity Verification process.

**Properties**

[`webhook_type`](/docs/api/products/identity-verification/#IdentityVerificationStepUpdatedWebhook-webhook-type)

stringstring

`IDENTITY_VERIFICATION`

[`webhook_code`](/docs/api/products/identity-verification/#IdentityVerificationStepUpdatedWebhook-webhook-code)

stringstring

`STEP_UPDATED`

[`identity_verification_id`](/docs/api/products/identity-verification/#IdentityVerificationStepUpdatedWebhook-identity-verification-id)

stringstring

The ID of the associated Identity Verification attempt.

[`environment`](/docs/api/products/identity-verification/#IdentityVerificationStepUpdatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "IDENTITY_VERIFICATION",
  "webhook_code": "STEP_UPDATED",
  "identity_verification_id": "idv_52xR9LKo77r1Np",
  "environment": "production"
}
```

=\*=\*=\*=

#### `RETRIED`

Fired when identity verification has been retried, which can be triggered via the dashboard or the API.

**Properties**

[`webhook_type`](/docs/api/products/identity-verification/#IdentityVerificationRetriedWebhook-webhook-type)

stringstring

`IDENTITY_VERIFICATION`

[`webhook_code`](/docs/api/products/identity-verification/#IdentityVerificationRetriedWebhook-webhook-code)

stringstring

`RETRIED`

[`identity_verification_id`](/docs/api/products/identity-verification/#IdentityVerificationRetriedWebhook-identity-verification-id)

stringstring

The ID of the associated Identity Verification attempt.

[`environment`](/docs/api/products/identity-verification/#IdentityVerificationRetriedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "IDENTITY_VERIFICATION",
  "webhook_code": "RETRIED",
  "identity_verification_id": "idv_52xR9LKo77r1Np",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
