---
title: "Errors - Income errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/income/"
scraped_at: "2026-03-07T22:04:49+00:00"
---

# Income errors

#### Guide to troubleshooting Income errors

#### **INCOME\_VERIFICATION\_DOCUMENT\_NOT\_FOUND**

##### The requested income verification document was not found.

##### Common causes

- The document URL is incorrect (for example, it may contain a typo) or has expired.
- The document URL has already been accessed. Document URLs can only be used once.

##### Troubleshooting steps

Check that the document URL is entered correctly.

Make a new call to [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) to generate a new document URL.

API error response

```
http code 400
{
 "error_type": "INCOME_VERIFICATION_ERROR",
 "error_code": "INCOME_VERIFICATION_DOCUMENT_NOT_FOUND",
 "error_message": "the requested data was not found. Please check the ID supplied.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INCOME\_VERIFICATION\_FAILED**

##### The income verification you are trying to retrieve could not be completed. please try creating a new income verification

##### Common causes

- The processing of the verification failed to complete successfully.

##### Troubleshooting steps

Have the user retry the verification.

If the problem persists, [contact Plaid Support](https://dashboard.plaid.com/support).

API error response

```
http code 400
{
 "error_type": "INCOME_VERIFICATION_ERROR",
 "error_code": "INCOME_VERIFICATION_FAILED",
 "error_message": "the income verification you are trying to retrieve could not be completed. please try creating a new income verification",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INCOME\_VERIFICATION\_NOT\_FOUND**

##### The requested income verification was not found.

##### Common causes

- The Link flow has not been completed with `income_verification` enabled; either the user has not yet completed the Link flow, or the link token was not initialized with the `income_verification` product.

##### Troubleshooting steps

Make sure to initialize Link with `income_verification` in the product array.

Listen for the [`INCOME: INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook with a corresponding `item_id`, which will fire once verification is complete.

API error response

```
http code 400
{
 "error_type": "INCOME_VERIFICATION_ERROR",
 "error_code": "INCOME_VERIFICATION_NOT_FOUND",
 "error_message": "the requested data was not found. Please check the ID supplied.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INCOME\_VERIFICATION\_UPLOAD\_ERROR**

##### There was an error uploading income verification documents.

##### Common causes

- The end user's Internet connection may have been interrupted during the upload attempt.

##### Troubleshooting steps

Plaid will prompt the end user to re-try uploading their documents.

API error response

```
http code 500
{
 "error_type": "INCOME_VERIFICATION_ERROR",
 "error_code": "INCOME_VERIFICATION_UPLOAD_ERROR",
 "error_message": "there was a problem uploading the document for verification. Please try again or recreate an income verification.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PRODUCT\_NOT\_ENABLED**

##### The Income product has not been enabled.

##### Common causes

- The Income product has not been enabled.

##### Troubleshooting steps

Make sure to initialize Link with `income_verification` in the `product` array.

Contact your Plaid Account Manager or [Plaid Support](https://dashboard.plaid.com/support) to request that your account be enabled for Income Verification.

API error response

```
http code 400
{
 "error_type": "INCOME_VERIFICATION_ERROR",
 "error_code": "PRODUCT_NOT_ENABLED",
 "error_message": "the 'income_verification' product is not enabled for the following client ID: <CLIENT_ID>. please ensure that the 'income_verification' is included in the 'product' array when initializing Link and try again.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PRODUCT\_NOT\_READY**

##### The Income product has not completed processing.

##### Common causes

- Parsing of the uploaded pay stubs has not yet completed.

##### Troubleshooting steps

Listen for the [`INCOME: INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook with a corresponding `item_id`, which will fire once verification is complete.

API error response

```
http code 400
{
 "error_type": "INCOME_VERIFICATION_ERROR",
 "error_code": "PRODUCT_NOT_READY",
 "error_message": "the requested product is not yet ready. please provide a webhook or try the request again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **VERIFICATION\_STATUS\_PENDING\_APPROVAL**

##### The user has not yet authorized the sharing of this data

##### Common causes

- The user has not yet authorized access to the data.

##### Troubleshooting steps

Listen for the [`INCOME: INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook with a corresponding `item_id`, which will fire once verification is complete.

Prompt the user to authorize access, then try again later.

API error response

```
http code 400
{
 "error_type": "INCOME_VERIFICATION_ERROR",
 "error_code": "VERIFICATION_STATUS_PENDING_APPROVAL",
 "error_message": "The user has not yet authorized the sharing of this data",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **EMPLOYMENT\_NOT\_FOUND**

##### The requested employment was not found.

##### Common causes

- The Link flow has not been completed with `employment` enabled; either the user has not yet completed the Link flow, or the link token was not initialized with the `employment` product.

##### Troubleshooting steps

Make sure to initialize Link with `employment` in the product array.

Listen for the [`INCOME: INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook with a corresponding `item_id`, which will fire once verification is complete.

API error response

```
http code 400
{
 "error_type": "INCOME_VERIFICATION_ERROR",
 "error_code": "EMPLOYMENT_NOT_FOUND",
 "error_message": "the requested employment data was not found. Please check the ID supplied.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
