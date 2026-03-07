---
title: "Errors - API errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/api/"
scraped_at: "2026-03-07T22:04:49+00:00"
---

# API Errors

#### Guide to troubleshooting API errors

#### **INTERNAL\_SERVER\_ERROR** or **plaid-internal-error**

##### Plaid was unable to process the request

##### Sample user-facing error message

Something went wrong: You can try again later or find another bank to continue

##### Link user experience

Your user will be redirected to the Institution Select pane to retry connecting their Item or a different account.

##### Common causes

- Plaid received an unsupported response from a financial institution, which frequently corresponds to an [institution](/docs/errors/institution/) error.
- Plaid is experiencing internal system issues.
- A product endpoint request was made for an Item at an OAuth-based institution, but the end user did not authorize the Item for the specific product, or has revoked Plaid's access to the product. Note that for some institutions, the end user may need to specifically opt-in during the OAuth flow to share specific details, such as identity data, or account and routing number information, even if they have already opted in to sharing information about a specific account.

##### Troubleshooting steps

Retry the request. If the endpoint supports the use of an idempotency key parameter, ensure you are using one before retrying. If the endpoint is [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate), see [Handling errors in transfer creation](/docs/transfer/creating-transfers/#handling-errors-in-transfer-creation).

If the Item is at an OAuth-based institution, prompt the end user to allow Plaid to access identity data and/or account and routing number data. The end user should do this during the Link flow if they were unable to successfully complete the Link flow for the Item, or at their institution's online banking portal if the Item has already been added.

Prompt your user to retry connecting their Item in a few hours or the following day.

Check the status of the institution via the [Dashboard](https://dashboard.plaid.com/activity/status) or [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id).

If the error persists, please submit a [Support](https://dashboard.plaid.com/support/new/) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 500
{
 "error_type": "API_ERROR",
 "error_code": "INTERNAL_SERVER_ERROR",
 "error_message": "an unexpected error occurred",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PLANNED\_MAINTENANCE**

##### Plaid's systems are undergoing maintenance and API operations are disabled

##### Sample user-facing error message

Please try connecting a different account: There was a problem processing your request. Your account could not be connected at this time

##### Link user experience

Your user will be redirected to the Institution Select pane to retry connecting their Item or a different account.

##### Common causes

- Plaid's systems are under maintenance and API operations are temporarily disabled. Advance notice will be provided when a maintenance window is planned.

##### Troubleshooting steps

Check Plaid's [System status](https://status.plaid.com/) for any recent maintenance updates.

If you have not been previously informed of planned maintenance, please reach out to Plaid [Support](https://dashboard.plaid.com/support/new/) for more information.

API error response

```
http code 503
{
 "error_type": "API_ERROR",
 "error_code": "PLANNED_MAINTENANCE",
 "error_message": "the Plaid API is temporarily unavailable due to planned maintenance. visit https://status.plaid.com/ for more information",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
