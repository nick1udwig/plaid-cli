---
title: "Errors - Institution errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/institution/"
scraped_at: "2026-03-07T22:04:50+00:00"
---

# Institution Errors

#### Guide to troubleshooting Institution errors

#### **INSTITUTION\_DOWN**

##### The financial institution is down, either for maintenance or due to an infrastructure issue with their systems.

##### Sample user-facing error message

Please try connecting a different account: There was a problem processing your request. Your account could not be connected at this time

##### Link user experience

Your user will be redirected to the `Institution Select` pane to retry connecting their Item or a different account.

##### Common causes

- The institution is undergoing scheduled maintenance.
- The institution is experiencing temporary technical problems.

##### Troubleshooting steps

Prompt your user to retry connecting their Item in a few hours, or the following day.

Check the status of the institution via the [Dashboard](https://dashboard.plaid.com/activity/status) or [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id).

If the error persists, please submit a 'Persistent HTTP 500 errors' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/persistent-500) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "INSTITUTION_ERROR",
 "error_code": "INSTITUTION_DOWN",
 "error_message": "this institution is not currently responding to this request. please try again soon",
 "display_message": "This financial institution is not currently responding to requests. We apologize for the inconvenience.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTION\_NOT\_ENABLED\_IN\_ENVIRONMENT**

##### Institution not enabled in this environment

##### Common causes

You're referencing an institution that exists, but is not enabled for this environment (e.g. calling a Sandbox endpoint with the ID of an institution that is not enabled there).

API error response

```
http code 400
{
 "error_type": "INSTITUTION_ERROR",
 "error_code": "INSTITUTION_NOT_ENABLED_IN_ENVIRONMENT",
 "error_message": "Institution not enabled in this environment",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTION\_NOT\_RESPONDING**

##### The financial institution is failing to respond to some of our requests. Your user may be successful if they retry another time.

##### Sample user-facing error message

Couldn't connect to your institution: If you need to use this app immediately, we recommend trying another institution. Errors like this may take some time to resolve, so you may want to try the same account again later.

##### Link user experience

Your user will be redirected to the `Institution Select` pane to retry connecting their Item or a different account.

##### Common causes

- The institution is experiencing temporary technical problems.

##### Troubleshooting steps

Prompt your user to retry connecting their Item in a few hours, or the following day.

Check the status of the institution via the [Dashboard](https://dashboard.plaid.com/activity/status) or [Institutions API](/docs/api/institutions/).

If the error persists, please submit a 'Persistent HTTP 500 errors' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/persistent-500) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "INSTITUTION_ERROR",
 "error_code": "INSTITUTION_NOT_RESPONDING",
 "error_message": "this institution is not currently responding to this request. please try again soon",
 "display_message": "This financial institution is not currently responding to requests. We apologize for the inconvenience.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTION\_NOT\_AVAILABLE**

##### The financial institution is not available this time.

##### Link user experience

Your user will be redirected to the `Institution Select` pane to retry connecting their Item or a different account.

##### Common causes

- Plaid’s connection to an institution is temporarily down for maintenance or other planned circumstances.

##### Troubleshooting steps

Prompt your user to retry connecting their Item in a few hours or the following day.

Check the status of the institution via the [Dashboard](https://dashboard.plaid.com/activity/status) or [API](/docs/api/institutions/).

If the error persists, please submit a 'Persistent HTTP 500 errors' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/persistent-500) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "INSTITUTION_ERROR",
 "error_code": "INSTITUTION_NOT_AVAILABLE",
 "error_message": "this institution is not currently responding to this request. please try again soon",
 "display_message": "This financial institution is not currently responding to requests. We apologize for the inconvenience.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTION\_NO\_LONGER\_SUPPORTED**

##### Plaid no longer supports this financial institution.

##### Sample user-facing error message

Institution Not Supported: Plaid does not support this institution yet, but we are working hard to add coverage based on your interest! In the meantime, please try another institution.

##### Common causes

- The financial institution is no longer supported on Plaid's platform.
- The institution has switched from supporting non-OAuth-based connections to requiring OAuth-based connections.

If the affected institution is Capital One, and you do not already have Production access, apply for Production access via the Plaid Dashboard.

If the affected institution is Capital One and you already have Production access, refer to the migration guide provided by Plaid via email for OAuth migration instructions. If you are unable to find your migration guide, [contact Support](https://dashboard.plaid.com/support).

For other institutions, this error is un-retryable and requires custom updates from Plaid to resolve. Submit a [Support](https://dashboard.plaid.com/support/new/financial-institutions) ticket with the failing `institution_id` for more detailed information.

API error response

```
http code 400
{
 "error_type": "INSTITUTION_ERROR",
 "error_code": "INSTITUTION_NO_LONGER_SUPPORTED",
 "error_message": "this institution is no longer supported",
 "display_message": "This financial institution is no longer supported. We apologize for the inconvenience.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **UNAUTHORIZED\_INSTITUTION**

##### You are not authorized to create Items for this institution at this time.

##### Common causes

- Access to this institution is subject to additional verification and must be manually enabled for your account.
- You have not yet completed the OAuth registration requirements for the institution.

##### Troubleshooting steps

Make sure your account is enabled for Production access.

See the [OAuth institutions](https://dashboard.plaid.com/settings/compliance/us-oauth-institutions) page in the Dashboard for a list of incomplete OAuth registration requirements.

If there are no incomplete OAuth registration requirements displayed for this institution in the Dashboard, or if you still cannot access the institution two weeks after completing the registration requirements, contact your Plaid Account Manager or [Plaid Support](https://dashboard.plaid.com/support).

API error response

```
http code 400
{
 "error_type": "INSTITUTION_ERROR",
 "error_code": "UNAUTHORIZED_INSTITUTION",
 "error_message": "not authorized to create items for this institution",
 "display_message": "You are not authorized to create items for this institution at this time.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTION\_REGISTRATION\_REQUIRED**

##### Your application is not yet registered with the institution.

##### Common causes

- Details about your application must be registered with this institution before use.
- If your account has not yet been enabled for Production access, some institutions that require registration may not be available to you.
- You have not completed the Security Questionnaire.

##### Troubleshooting steps

Make sure to complete the [Security Questionnaire](https://dashboard.plaid.com/overview/questionnaire).

If your account was recently enabled for Production or you recently completed the Security Questionnaire, please allow up to a week for registration to be completed with the institution. For some institutions, such as Chase and Charles Schwab, registration may take up to six weeks.

If you are still unable to access the institution after completing the above steps and waiting the appropriate time, please [contact Support](https://dashboard.plaid.com/support) regarding the status of your registration.

API error response

```
http code 400
{
 "error_type": "INSTITUTION_ERROR",
 "error_code": "INSTITUTION_REGISTRATION_REQUIRED",
 "error_message": "not yet registered to create items for this institution",
 "display_message": "You must register your application with this institution before you can create items for it.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **UNSUPPORTED\_RESPONSE**

##### The data returned from the financial institution is not valid.

##### Common causes

- The data returned from the financial institution is not supported or is invalid.

This error is un-retryable and requires custom updates from Plaid to resolve. Submit a [Support](https://dashboard.plaid.com/support/new/financial-institutions) ticket with the failing `institution_id` for more detailed information.

API error response

```
http code 400
{
 "error_type": "INSTITUTION_ERROR",
 "error_code": "UNSUPPORTED_RESPONSE",
 "error_message": "the data returned from the financial institution is not valid",
 "display_message": "The data returned from the financial institution is not valid.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
