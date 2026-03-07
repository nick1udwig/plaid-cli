---
title: "API - Reseller partners | Plaid Docs"
source_url: "https://plaid.com/docs/api/partner/"
scraped_at: "2026-03-07T22:03:51+00:00"
---

# Partner endpoints and webhooks

#### Create and manage end customers

For general, non-reference documentation, see [Reseller partners](/docs/account/resellers/).

| Endpoints |  |
| --- | --- |
| [`/partner/customer/create`](/docs/api/partner/#partnercustomercreate) | Create an end customer |
| [`/partner/customer/get`](/docs/api/partner/#partnercustomerget) | Get the status of an end customer |
| [`/partner/customer/oauth_institutions/get`](/docs/api/partner/#partnercustomeroauth_institutionsget) | Get the OAuth-institution registration status for an end customer |
| [`/partner/customer/enable`](/docs/api/partner/#partnercustomerenable) | Enable an end customer in Production |
| [`/partner/customer/remove`](/docs/api/partner/#partnercustomerremove) | Remove an end customer |

| Webhooks |  |
| --- | --- |
| [`PARTNER_END_CUSTOMER_OAUTH_STATUS_UPDATED`](/docs/api/partner/#partner_end_customer_oauth_status_updated) | Customer OAuth status updated |

### Endpoints

=\*=\*=\*=

#### `/partner/customer/create`

#### Creates a new end customer for a Plaid reseller.

The [`/partner/customer/create`](/docs/api/partner/#partnercustomercreate) endpoint is used by reseller partners to create end customers. To create end customers, it should be called in the Production environment only, even when creating Sandbox API keys. If called in the Sandbox environment, it will return a sample response, but no customer will be created and the API keys will not be valid.

/partner/customer/create

**Request fields**

[`client_id`](/docs/api/partner/#partner-customer-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/partner/#partner-customer-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`company_name`](/docs/api/partner/#partner-customer-create-request-company-name)

requiredstringrequired, string

The company name of the end customer being created. This will be used to display the end customer in the Plaid Dashboard. It will not be shown to end users.

[`is_diligence_attested`](/docs/api/partner/#partner-customer-create-request-is-diligence-attested)

requiredbooleanrequired, boolean

Denotes whether or not the partner has completed attestation of diligence for the end customer to be created.

[`products`](/docs/api/partner/#partner-customer-create-request-products)

[string][string]

The products to be enabled for the end customer. If empty or `null`, this field will default to the products enabled for the reseller at the time this endpoint is called.  
  

Possible values: `assets`, `auth`, `balance`, `identity`, `income_verification`, `investments`, `investments_auth`, `liabilities`, `transactions`, `employment`, `cra_base_report`, `cra_income_insights`, `cra_partner_insights`

[`create_link_customization`](/docs/api/partner/#partner-customer-create-request-create-link-customization)

booleanboolean

If `true`, the end customer's default Link customization will be set to match the partner's. You can always change the end customer's Link customization in the Plaid Dashboard. See the [Link Customization docs](https://plaid.com/docs/link/customization/) for more information. If you require the ability to programmatically create end customers using multiple different Link customization profiles, contact your Plaid Account Manager for assistance.   
Important: Data Transparency Messaging (DTM) use cases will not be copied to the end customer's Link customization unless the **Publish changes** button is clicked after the use cases are applied. Link will not work in Production unless the end customer's DTM use cases are configured. For more details, see [Data Transparency Messaging](https://plaid.com/docs/link/data-transparency-messaging-migration-guide/).

[`logo`](/docs/api/partner/#partner-customer-create-request-logo)

stringstring

Base64-encoded representation of the end customer's logo. Must be a PNG of size 1024x1024 under 4MB. The logo will be shared with financial institutions and shown to the end user during Link flows. A logo is required if `create_link_customization` is `true`. If `create_link_customization` is `false` and the logo is omitted, the partner's logo will be used if one exists, otherwise a stock logo will be used.

[`legal_entity_name`](/docs/api/partner/#partner-customer-create-request-legal-entity-name)

requiredstringrequired, string

The end customer's legal name. This will be shared with financial institutions as part of the OAuth registration process. It will not be shown to end users.

[`website`](/docs/api/partner/#partner-customer-create-request-website)

requiredstringrequired, string

The end customer's website.

[`application_name`](/docs/api/partner/#partner-customer-create-request-application-name)

requiredstringrequired, string

The name of the end customer's application. This will be shown to end users when they go through the Plaid Link flow. The application name must be unique and cannot match the name of another application already registered with Plaid.

[`technical_contact`](/docs/api/partner/#partner-customer-create-request-technical-contact)

objectobject

The technical contact for the end customer. Defaults to partner's technical contact if omitted.

[`given_name`](/docs/api/partner/#partner-customer-create-request-technical-contact-given-name)

stringstring

[`family_name`](/docs/api/partner/#partner-customer-create-request-technical-contact-family-name)

stringstring

[`email`](/docs/api/partner/#partner-customer-create-request-technical-contact-email)

stringstring

[`billing_contact`](/docs/api/partner/#partner-customer-create-request-billing-contact)

objectobject

The billing contact for the end customer. Defaults to partner's billing contact if omitted.

[`given_name`](/docs/api/partner/#partner-customer-create-request-billing-contact-given-name)

stringstring

[`family_name`](/docs/api/partner/#partner-customer-create-request-billing-contact-family-name)

stringstring

[`email`](/docs/api/partner/#partner-customer-create-request-billing-contact-email)

stringstring

[`customer_support_info`](/docs/api/partner/#partner-customer-create-request-customer-support-info)

objectobject

This information is public. Users of your app will see this information when managing connections between your app and their bank accounts in Plaid Portal. Defaults to partner's customer support info if omitted. This field is mandatory for partners whose Plaid accounts were created after November 26, 2024 and will be mandatory for all partners by the 1033 compliance deadline.

[`email`](/docs/api/partner/#partner-customer-create-request-customer-support-info-email)

stringstring

This field is mandatory for partners whose Plaid accounts were created after November 26, 2024 and will be mandatory for all partners by the 1033 compliance deadline.

[`phone_number`](/docs/api/partner/#partner-customer-create-request-customer-support-info-phone-number)

stringstring

[`contact_url`](/docs/api/partner/#partner-customer-create-request-customer-support-info-contact-url)

stringstring

[`link_update_url`](/docs/api/partner/#partner-customer-create-request-customer-support-info-link-update-url)

stringstring

[`address`](/docs/api/partner/#partner-customer-create-request-address)

requiredobjectrequired, object

The end customer's address.

[`city`](/docs/api/partner/#partner-customer-create-request-address-city)

stringstring

[`street`](/docs/api/partner/#partner-customer-create-request-address-street)

stringstring

[`region`](/docs/api/partner/#partner-customer-create-request-address-region)

stringstring

[`postal_code`](/docs/api/partner/#partner-customer-create-request-address-postal-code)

stringstring

[`country_code`](/docs/api/partner/#partner-customer-create-request-address-country-code)

stringstring

ISO-3166-1 alpha-2 country code standard.

[`is_bank_addendum_completed`](/docs/api/partner/#partner-customer-create-request-is-bank-addendum-completed)

requiredbooleanrequired, boolean

Denotes whether the partner has forwarded the Plaid bank addendum to the end customer.

[`assets_under_management`](/docs/api/partner/#partner-customer-create-request-assets-under-management)

objectobject

Assets under management for the given end customer. Required for end customers with monthly service commitments.

[`amount`](/docs/api/partner/#partner-customer-create-request-assets-under-management-amount)

requirednumberrequired, number

[`iso_currency_code`](/docs/api/partner/#partner-customer-create-request-assets-under-management-iso-currency-code)

requiredstringrequired, string

[`redirect_uris`](/docs/api/partner/#partner-customer-create-request-redirect-uris)

[string][string]

A list of URIs indicating the destination(s) where a user can be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or another app. URIs should not contain any query parameters. When used in Production, URIs must use https. To specify any subdomain, use `*` as a wildcard character, e.g. `https://*.example.com/oauth.html`. To modify redirect URIs for an end customer after creating them, go to the end customer's [API page](https://dashboard.plaid.com/team/api) in the Dashboard.

[`registration_number`](/docs/api/partner/#partner-customer-create-request-registration-number)

stringstring

The unique identifier assigned to a financial institution by regulatory authorities, if applicable. For banks, this is the FDIC Certificate Number. For credit unions, this is the Credit Union Charter Number.

/partner/customer/create

```
const request: PartnerCustomerCreateRequest = {
  address: {
    city: city,
    country_code: countryCode,
    postal_code: postalCode,
    region: region,
    street: street,
  },
  application_name: applicationName,
  billing_contact: {
    email: billingEmail,
    given_name: billingGivenName,
    family_name: billingFamilyName,
  },
  customer_support_info: {
    email: supportEmail,
    phone_number: supportPhoneNumber,
    contact_url: supportContactUrl,
    link_update_url: linkUpdateUrl,
  },
  company_name: companyName,
  is_bank_addendum_completed: true,
  is_diligence_attested: true,
  legal_entity_name: legalEntityName,
  products: products,
  technical_contact: {
    email: technicalEmail,
    given_name: technicalGivenName,
    family_name: technicalFamilyName,
  },
  website: website,
};
try {
  const response = await plaidClient.partnerCustomerCreate(request);
  const endCustomer = response.data.end_customer;
} catch (error) {
  // handle error
}
```

/partner/customer/create

**Response fields**

[`end_customer`](/docs/api/partner/#partner-customer-create-response-end-customer)

objectobject

The details for the newly created end customer, including secrets for Sandbox and Limited Production.

[`client_id`](/docs/api/partner/#partner-customer-create-response-end-customer-client-id)

stringstring

The `client_id` of the end customer.

[`company_name`](/docs/api/partner/#partner-customer-create-response-end-customer-company-name)

stringstring

The company name associated with the end customer.

[`status`](/docs/api/partner/#partner-customer-create-response-end-customer-status)

stringstring

The status of the given end customer.  
`UNDER_REVIEW`: The end customer has been created and enabled in Sandbox and Limited Production. The end customer must be manually reviewed by the Plaid team before it can be enabled in full production, at which point its status will automatically transition to `PENDING_ENABLEMENT` or `DENIED`.  
`PENDING_ENABLEMENT`: The end customer is ready to be fully enabled in the Production environment. Call the `/partner/customer/enable` endpoint to enable the end customer in full Production.  
`ACTIVE`: The end customer has been fully enabled in all environments.  
`DENIED`: The end customer has been created and enabled in Sandbox and Limited Production, but it did not pass review by the Plaid team and therefore cannot be enabled for full Production access. Talk to your Account Manager for more information.  
  

Possible values: `UNDER_REVIEW`, `PENDING_ENABLEMENT`, `ACTIVE`, `DENIED`

[`secrets`](/docs/api/partner/#partner-customer-create-response-end-customer-secrets)

objectobject

The secrets for the newly created end customer.

[`sandbox`](/docs/api/partner/#partner-customer-create-response-end-customer-secrets-sandbox)

stringstring

The end customer's secret key for the Sandbox environment.

[`production`](/docs/api/partner/#partner-customer-create-response-end-customer-secrets-production)

stringstring

The end customer's secret key for the Production environment. The end customer will be provided with a limited number of credits to test in the Production environment before full enablement.

[`request_id`](/docs/api/partner/#partner-customer-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "end_customer": {
    "client_id": "7f57eb3d2a9j6480121fx361",
    "company_name": "Plaid",
    "status": "ACTIVE",
    "secrets": {
      "sandbox": "b60b5201d006ca5a7081d27c824d77",
      "production": "79g03eoofwl8240v776r2h667442119"
    }
  },
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/partner/customer/get`

#### Returns a Plaid reseller's end customer.

The [`/partner/customer/get`](/docs/api/partner/#partnercustomerget) endpoint is used by reseller partners to retrieve data about a single end customer.

/partner/customer/get

**Request fields**

[`client_id`](/docs/api/partner/#partner-customer-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/partner/#partner-customer-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`end_customer_client_id`](/docs/api/partner/#partner-customer-get-request-end-customer-client-id)

requiredstringrequired, string

/partner/customer/get

```
const request: PartnerCustomerGetRequest = {
  end_customer_client_id: clientId,
};
try {
  const response = await plaidClient.partnerCustomerGet(request);
  const endCustomer = response.data.end_customer;
} catch (error) {
  // handle error
}
```

/partner/customer/get

**Response fields**

[`end_customer`](/docs/api/partner/#partner-customer-get-response-end-customer)

objectobject

The details for an end customer.

[`client_id`](/docs/api/partner/#partner-customer-get-response-end-customer-client-id)

stringstring

The `client_id` of the end customer.

[`company_name`](/docs/api/partner/#partner-customer-get-response-end-customer-company-name)

stringstring

The company name associated with the end customer.

[`status`](/docs/api/partner/#partner-customer-get-response-end-customer-status)

stringstring

The status of the given end customer.  
`UNDER_REVIEW`: The end customer has been created and enabled in Sandbox and Limited Production. The end customer must be manually reviewed by the Plaid team before it can be enabled in full production, at which point its status will automatically transition to `PENDING_ENABLEMENT` or `DENIED`.  
`PENDING_ENABLEMENT`: The end customer is ready to be fully enabled in the Production environment. Call the `/partner/customer/enable` endpoint to enable the end customer in full Production.  
`ACTIVE`: The end customer has been fully enabled in all environments.  
`DENIED`: The end customer has been created and enabled in Sandbox and Limited Production, but it did not pass review by the Plaid team and therefore cannot be enabled for full Production access. Talk to your Account Manager for more information.  
  

Possible values: `UNDER_REVIEW`, `PENDING_ENABLEMENT`, `ACTIVE`, `DENIED`

[`request_id`](/docs/api/partner/#partner-customer-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "end_customer": {
    "client_id": "7f57eb3d2a9j6480121fx361",
    "company_name": "Plaid",
    "status": "ACTIVE"
  },
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/partner/customer/oauth_institutions/get`

#### Returns OAuth-institution registration information for a given end customer.

The [`/partner/customer/oauth_institutions/get`](/docs/api/partner/#partnercustomeroauth_institutionsget) endpoint is used by reseller partners to retrieve OAuth-institution registration information about a single end customer. To learn how to set up a webhook to listen to status update events, visit the [reseller documentation](https://plaid.com/docs/account/resellers/#enabling-end-customers).

/partner/customer/oauth\_institutions/get

**Request fields**

[`client_id`](/docs/api/partner/#partner-customer-oauth_institutions-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/partner/#partner-customer-oauth_institutions-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`end_customer_client_id`](/docs/api/partner/#partner-customer-oauth_institutions-get-request-end-customer-client-id)

requiredstringrequired, string

/partner/customer/oauth\_institutions/get

```
const request: PartnerCustomerOAuthInstitutionsGetRequest = {
  end_customer_client_id: clientId,
};
try {
  const response = await plaidClient.partnerCustomerOAuthInstitutionsGet(
    request,
  );
} catch (error) {
  // handle error
}
```

/partner/customer/oauth\_institutions/get

**Response fields**

[`flowdown_status`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-flowdown-status)

stringstring

The status of the addendum to the Plaid MSA ("flowdown") for the end customer.  
  

Possible values: `NOT_STARTED`, `IN_REVIEW`, `NEGOTIATION`, `COMPLETE`

[`questionnaire_status`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-questionnaire-status)

stringstring

The status of the end customer's security questionnaire.  
  

Possible values: `NOT_STARTED`, `RECEIVED`, `COMPLETE`

[`institutions`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions)

[object][object]

The OAuth institutions with which the end customer's application is being registered.

[`name`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-name)

stringstring

[`institution_id`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-institution-id)

stringstring

[`environments`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-environments)

objectobject

Registration statuses by environment.

[`development`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-environments-development)

stringstring

The registration status for the end customer's application.  
  

Possible values: `NOT_STARTED`, `PROCESSING`, `APPROVED`, `ENABLED`, `ATTENTION_REQUIRED`

[`production`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-environments-production)

stringstring

The registration status for the end customer's application.  
  

Possible values: `NOT_STARTED`, `PROCESSING`, `APPROVED`, `ENABLED`, `ATTENTION_REQUIRED`

[`production_enablement_date`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-production-enablement-date)

nullablestringnullable, string

The date on which the end customer's application was approved by the institution, or an empty string if their application has not yet been approved.

[`classic_disablement_date`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-classic-disablement-date)

nullablestringnullable, string

The date on which non-OAuth Item adds will no longer be supported for this institution, or an empty string if no such date has been set by the institution.

[`errors`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors)

[object][object]

The errors encountered while registering the end customer's application with the institutions.

[`error_type`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-error-code-reason)

nullablestringnullable, string

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-display-message)

nullablestringnullable, string

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-status)

nullableintegernullable, integer

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-suggested-action)

nullablestringnullable, string

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-institutions-errors-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`request_id`](/docs/api/partner/#partner-customer-oauth_institutions-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "flowdown_status": "COMPLETE",
  "questionnaire_status": "COMPLETE",
  "institutions": [
    {
      "name": "Chase",
      "institution_id": "ins_56",
      "environments": {
        "production": "PROCESSING"
      },
      "production_enablement_date": null,
      "classic_disablement_date": "2022-06-30"
    },
    {
      "name": "Capital One",
      "institution_id": "ins_128026",
      "environments": {
        "production": "ENABLED"
      },
      "production_enablement_date": "2022-12-19",
      "classic_disablement_date": null
    },
    {
      "name": "Bank of America",
      "institution_id": "ins_1",
      "environments": {
        "production": "ATTENTION_REQUIRED"
      },
      "production_enablement_date": null,
      "classic_disablement_date": null,
      "errors": [
        {
          "error_type": "PARTNER_ERROR",
          "error_code": "OAUTH_REGISTRATION_ERROR",
          "error_message": "Application logo is required",
          "display_message": null,
          "request_id": "4zlKapIkTm8p5KM"
        }
      ]
    }
  ],
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/partner/customer/enable`

#### Enables a Plaid reseller's end customer in the Production environment.

The [`/partner/customer/enable`](/docs/api/partner/#partnercustomerenable) endpoint is used by reseller partners to enable an end customer in the full Production environment.

/partner/customer/enable

**Request fields**

[`client_id`](/docs/api/partner/#partner-customer-enable-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/partner/#partner-customer-enable-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`end_customer_client_id`](/docs/api/partner/#partner-customer-enable-request-end-customer-client-id)

requiredstringrequired, string

/partner/customer/enable

```
const request: PartnerCustomerEnableRequest = {
  end_customer_client_id: clientId,
};
try {
  const response = await plaidClient.partnerCustomerEnable(request);
  const productionSecret = response.data.production_secret;
} catch (error) {
  // handle error
}
```

/partner/customer/enable

**Response fields**

[`production_secret`](/docs/api/partner/#partner-customer-enable-response-production-secret)

stringstring

The end customer's secret key for the Production environment.

[`request_id`](/docs/api/partner/#partner-customer-enable-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "production_secret": "79g03eoofwl8240v776r2h667442119",
  "request_id": "4zlKapIkTm8p5KM"
}
```

=\*=\*=\*=

#### `/partner/customer/remove`

#### Removes a Plaid reseller's end customer.

The [`/partner/customer/remove`](/docs/api/partner/#partnercustomerremove) endpoint is used by reseller partners to remove an end customer. Removing an end customer will remove it from view in the Plaid Dashboard and deactivate its API keys. This endpoint can only be used to remove an end customer that has not yet been enabled in full Production.

/partner/customer/remove

**Request fields**

[`client_id`](/docs/api/partner/#partner-customer-remove-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/partner/#partner-customer-remove-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`end_customer_client_id`](/docs/api/partner/#partner-customer-remove-request-end-customer-client-id)

requiredstringrequired, string

The `client_id` of the end customer to be removed.

/partner/customer/remove

```
const request: PartnerCustomerRemoveRequest = {
  end_customer_client_id: clientId,
};
try {
  const response = await plaidClient.partnerCustomerRemove(request);
} catch (error) {
  // handle error
}
```

/partner/customer/remove

**Response fields**

[`request_id`](/docs/api/partner/#partner-customer-remove-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "4zlKapIkTm8p5KM"
}
```

### Webhooks

=\*=\*=\*=

#### `PARTNER_END_CUSTOMER_OAUTH_STATUS_UPDATED`

The webhook of type `PARTNER` and code `END_CUSTOMER_OAUTH_STATUS_UPDATED` will be fired when a partner's end customer has an update on their OAuth registration status with an institution.

**Properties**

[`webhook_type`](/docs/api/partner/#PartnerEndCustomerOAuthStatusUpdatedWebhook-webhook-type)

stringstring

`PARTNER`

[`webhook_code`](/docs/api/partner/#PartnerEndCustomerOAuthStatusUpdatedWebhook-webhook-code)

stringstring

`END_CUSTOMER_OAUTH_STATUS_UPDATED`

[`end_customer_client_id`](/docs/api/partner/#PartnerEndCustomerOAuthStatusUpdatedWebhook-end-customer-client-id)

stringstring

The client ID of the end customer

[`environment`](/docs/api/partner/#PartnerEndCustomerOAuthStatusUpdatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

[`institution_id`](/docs/api/partner/#PartnerEndCustomerOAuthStatusUpdatedWebhook-institution-id)

stringstring

The institution ID

[`institution_name`](/docs/api/partner/#PartnerEndCustomerOAuthStatusUpdatedWebhook-institution-name)

stringstring

The institution name

[`status`](/docs/api/partner/#PartnerEndCustomerOAuthStatusUpdatedWebhook-status)

stringstring

The OAuth status of the update  
  

Possible values: `not-started`, `processing`, `approved`, `enabled`, `attention-required`

API Object

```
{
  "webhook_type": "PARTNER",
  "webhook_code": "END_CUSTOMER_OAUTH_STATUS_UPDATED",
  "end_customer_client_id": "634758733ebb4f00134b85ea",
  "environment": "production",
  "institution_id": "ins_127989",
  "institution_name": "Bank of America",
  "status": "attention-required"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
