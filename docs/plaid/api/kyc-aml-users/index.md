---
title: "API - Look up Dashboard users | Plaid Docs"
source_url: "https://plaid.com/docs/api/kyc-aml-users/"
scraped_at: "2026-03-07T22:03:49+00:00"
---

# Dashboard User Audit API

#### API reference for viewing Dashboard users for Monitor and Beacon

These endpoints are used to look up a Dashboard user, as referenced in an `audit_trail` object from the [Monitor](/docs/api/products/monitor/) or [Beacon](/docs/api/products/beacon/) APIs.

| Endpoints |  |
| --- | --- |
| [`/dashboard_user/get`](/docs/api/kyc-aml-users/#dashboard_userget) | Retrieve information about Dashboard user |
| [`/dashboard_user/list`](/docs/api/kyc-aml-users/#dashboard_userlist) | List Dashboard users |

=\*=\*=\*=

#### `/dashboard_user/get`

#### Retrieve a dashboard user

The [`/dashboard_user/get`](/docs/api/kyc-aml-users/#dashboard_userget) endpoint provides details (such as email address) about a specific Dashboard user based on the `dashboard_user_id` field, which is returned in the `audit_trail` object of certain Monitor and Beacon endpoints. This can be used to identify the specific reviewer who performed a Dashboard action.

/dashboard\_user/get

**Request fields**

[`dashboard_user_id`](/docs/api/kyc-aml-users/#dashboard_user-get-request-dashboard-user-id)

requiredstringrequired, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`secret`](/docs/api/kyc-aml-users/#dashboard_user-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/kyc-aml-users/#dashboard_user-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

/dashboard\_user/get

```
const request: DashboardUserGetRequest = {
  dashboard_user_id: 'usr_1SUuwqBdK75GKi',
};

try {
  const response = await client.dashboardUserGet(request);
} catch (error) {
  // handle error
}
```

/dashboard\_user/get

**Response fields**

[`id`](/docs/api/kyc-aml-users/#dashboard_user-get-response-id)

stringstring

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`created_at`](/docs/api/kyc-aml-users/#dashboard_user-get-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`email_address`](/docs/api/kyc-aml-users/#dashboard_user-get-response-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`status`](/docs/api/kyc-aml-users/#dashboard_user-get-response-status)

stringstring

The current status of the user.  
  

Possible values: `invited`, `active`, `deactivated`

[`request_id`](/docs/api/kyc-aml-users/#dashboard_user-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "54350110fedcbaf01234ffee",
  "created_at": "2020-07-24T03:26:02Z",
  "email_address": "user@example.com",
  "status": "active",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/dashboard_user/list`

#### List dashboard users

The [`/dashboard_user/list`](/docs/api/kyc-aml-users/#dashboard_userlist) endpoint provides details (such as email address) all Dashboard users associated with your account. This can use used to audit or track the list of reviewers for Monitor, Beacon, and Identity Verification products.

/dashboard\_user/list

**Request fields**

[`secret`](/docs/api/kyc-aml-users/#dashboard_user-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/kyc-aml-users/#dashboard_user-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`cursor`](/docs/api/kyc-aml-users/#dashboard_user-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/dashboard\_user/list

```
try {
  const response = await client.dashboardUserList({});
} catch (error) {
  // handle error
}
```

/dashboard\_user/list

**Response fields**

[`dashboard_users`](/docs/api/kyc-aml-users/#dashboard_user-list-response-dashboard-users)

[object][object]

List of dashboard users

[`id`](/docs/api/kyc-aml-users/#dashboard_user-list-response-dashboard-users-id)

stringstring

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`created_at`](/docs/api/kyc-aml-users/#dashboard_user-list-response-dashboard-users-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`email_address`](/docs/api/kyc-aml-users/#dashboard_user-list-response-dashboard-users-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`status`](/docs/api/kyc-aml-users/#dashboard_user-list-response-dashboard-users-status)

stringstring

The current status of the user.  
  

Possible values: `invited`, `active`, `deactivated`

[`next_cursor`](/docs/api/kyc-aml-users/#dashboard_user-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/kyc-aml-users/#dashboard_user-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "dashboard_users": [
    {
      "id": "54350110fedcbaf01234ffee",
      "created_at": "2020-07-24T03:26:02Z",
      "email_address": "user@example.com",
      "status": "active"
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
