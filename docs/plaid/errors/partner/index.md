---
title: "Errors - Partner errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/partner/"
scraped_at: "2026-03-07T22:04:52+00:00"
---

# Partner Errors

#### Guide to troubleshooting partner errors

#### **CUSTOMER\_NOT\_FOUND**

##### Customer not found

##### Common causes

- The end customer could not be found.

##### Troubleshooting steps

Check the end customer ID.

Ensure that you're making the request in the correct Plaid environment.

API error response

```
http code 404
{
 "error_type": "PARTNER_ERROR",
 "error_code": "CUSTOMER_NOT_FOUND",
 "error_message": "the customer was not found",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **FLOWDOWN\_NOT\_COMPLETE**

##### Flowdown not complete

##### Common causes

- A flowdown, which is a prerequisite for creating an end customer, has not been completed.

##### Troubleshooting steps

Complete a flowdown in the [Plaid Dashboard](https://dashboard.plaid.com/settings/compliance/us-oauth-institutions).

API error response

```
http code 412
{
 "error_type": "PARTNER_ERROR",
 "error_code": "FLOWDOWN_NOT_COMPLETE",
 "error_message": "you must complete a flowdown to create and enable customers",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **QUESTIONNAIRE\_NOT\_COMPLETE**

##### Questionnaire not complete

##### Common causes

- A security questionnaire, which is a prerequisite for creating an end customer, has not been submitted.

##### Troubleshooting steps

Submit a security questionnaire in the [Plaid Dashboard](https://dashboard.plaid.com/overview/questionnaire).

API error response

```
http code 412
{
 "error_type": "PARTNER_ERROR",
 "error_code": "QUESTIONNAIRE_NOT_COMPLETE",
 "error_message": "you must complete the security questionnaire to create and enable customers",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **CUSTOMER\_NOT\_READY\_FOR\_ENABLEMENT**

##### Customer not ready for enablement

##### Common causes

- The end customer requires manual approval from Plaid's Partnerships team.
- The end customer has been denied access to the Plaid API.

##### Troubleshooting steps

Submit your request again after the end customer's status has been updated to `READY FOR ENABLEMENT`.

Talk to your Partner Account Manager.

API error response

```
http code 412
{
 "error_type": "PARTNER_ERROR",
 "error_code": "CUSTOMER_NOT_READY_FOR_ENABLEMENT",
 "error_message": "this customer is not ready for enablement",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **CUSTOMER\_ALREADY\_ENABLED**

##### Customer already enabled

##### Common causes

- The end customer has already been enabled.

##### Troubleshooting steps

Determine why your integration is attempting to enable end customers more than once.

API error response

```
http code 400
{
 "error_type": "PARTNER_ERROR",
 "error_code": "CUSTOMER_ALREADY_ENABLED",
 "error_message": "this customer has already been enabled",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **CUSTOMER\_ALREADY\_CREATED**

##### Customer already created

##### Common causes

- The end customer has already been created.

##### Troubleshooting steps

Determine why your integration is attempting to create end customers more than once.

API error response

```
http code 400
{
 "error_type": "PARTNER_ERROR",
 "error_code": "CUSTOMER_ALREADY_CREATED",
 "error_message": "this customer has already been created",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **LOGO\_REQUIRED**

##### Logo required

##### Common causes

- A logo is required for this customer because the `create_link_customization` field was set to `true` and the co-branded consent field is in use.

##### Troubleshooting steps

Add a logo to the request.

API error response

```
http code 400
{
 "error_type": "PARTNER_ERROR",
 "error_code": "LOGO_REQUIRED",
 "error_message": "a logo is required because create_link_customization is set to true and the co-branded consent pane is in use",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_LOGO**

##### Invalid logo

##### Common causes

- The logo is not a valid base64-encoded string of a PNG of size 1024x1024.

##### Troubleshooting steps

Check the logo size, file type, and encoding.

API error response

```
http code 400
{
 "error_type": "PARTNER_ERROR",
 "error_code": "INVALID_LOGO",
 "error_message": "the logo must be a base64-encoded string of a PNG of size 1024x1024",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **CONTACT\_REQUIRED**

##### Contact required

##### Common causes

- A billing or technical contact is required, either in the API request or in the Dashboard.

##### Troubleshooting steps

Add a billing and/or technical contact to your API request or in the [Dashboard](https://dashboard.plaid.com/settings/company/profile).

API error response

```
http code 400
{
 "error_type": "PARTNER_ERROR",
 "error_code": "CONTACT_REQUIRED",
 "error_message": "billing or technical contacts must be submitted either in the request or filled in for your team in the dashboard",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ASSETS\_UNDER\_MANAGEMENT\_REQUIRED**

##### Assets Under Management required

##### Common causes

- The `assets_under_management` field is required because you have a monthly service commitment, but it was omitted from your API request.

##### Troubleshooting steps

Set the `assets_under_management` field.

API error response

```
http code 400
{
 "error_type": "PARTNER_ERROR",
 "error_code": "ASSETS_UNDER_MANAGEMENT_REQUIRED",
 "error_message": "assets under management must be submitted because your team has a monthly service commitment",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **CUSTOMER\_REMOVAL\_NOT\_ALLOWED**

##### Removal not allowed

##### Common causes

- You have attempted to remove an end customer that has already been enabled in the Production environment, which is not allowed.

##### Troubleshooting steps

If you need to remove a Production-enabled end customer, talk to your Account Manager.

API error response

```
http code 400
{
 "error_type": "PARTNER_ERROR",
 "error_code": "CUSTOMER_REMOVAL_NOT_ALLOWED",
 "error_message": "removal of a production-enabled end customer is not allowed. talk to your account manager to remove this end customer.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **OAUTH\_REGISTRATION\_ERROR**

##### OAuth Registration Error

##### Common causes

- End customer error encountered while registering with OAuth institutions.

##### Troubleshooting steps

Log in to the end customer’s Dashboard from the [Partner End Customer Summary](https://dashboard.plaid.com/partner/customers-summary) page, and update the fields flagged in the error in the Compliance Center.

API error response

```
http code 400
{
 "error_type": "PARTNER_ERROR",
 "error_code": "OAUTH_REGISTRATION_ERROR",
 "error_message": "application logo is required",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
