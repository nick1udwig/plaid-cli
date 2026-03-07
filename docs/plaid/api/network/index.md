---
title: "API - Network | Plaid Docs"
source_url: "https://plaid.com/docs/api/network/"
scraped_at: "2026-03-07T22:03:50+00:00"
---

# Network

#### API reference for the Plaid Network

| Endpoints |  |
| --- | --- |
| [`/network/status/get`](/docs/api/network/#networkstatusget) | Check the status of a user in the Plaid Network |

### Endpoints

=\*=\*=\*=

#### `/network/status/get`

#### Check a user's Plaid Network status

The [`/network/status/get`](/docs/api/network/#networkstatusget) endpoint can be used to check whether Plaid has a matching profile for the user.
This is useful for determining if a user is eligible for a streamlined experience, such as Layer.

Note: it is strongly recommended to check for Layer eligibility in the frontend. [`/network/status/get`](/docs/api/network/#networkstatusget) should only be used for checking Layer eligibility if a frontend check is not possible for your use case.
For instructions on performing a frontend eligibility check, see the [Layer documentation](https://plaid.com/docs/layer/#integration-overview).

/network/status/get

**Request fields**

[`client_id`](/docs/api/network/#network-status-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/network/#network-status-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`user`](/docs/api/network/#network-status-get-request-user)

requiredobjectrequired, object

An object specifying information about the end user for the network status check.

[`phone_number`](/docs/api/network/#network-status-get-request-user-phone-number)

requiredstringrequired, string

The user's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.

[`template_id`](/docs/api/network/#network-status-get-request-template-id)

stringstring

The id of a template defined in Plaid Dashboard. This field is used if you have additional criteria that you want to check against (e.g. Layer eligibility).

/network/status/get

```
const request: NetworkStatusGetRequest = {
  user: {
    phone_number: '+14155550015',
  },
};
try {
  const response = await plaidClient.networkStatusGet(request);
  const networkStatus = response.data.network_status;
} catch (error) {
  // handle error
}
```

/network/status/get

**Response fields**

[`network_status`](/docs/api/network/#network-status-get-response-network-status)

stringstring

Enum representing the overall network status of the user.  
  

Possible values: `UNKNOWN`, `RETURNING_USER`

[`layer`](/docs/api/network/#network-status-get-response-layer)

nullableobjectnullable, object

An object representing Layer-related metadata for the requested user.

[`eligible`](/docs/api/network/#network-status-get-response-layer-eligible)

booleanboolean

Indicates if the user is eligible for a Layer session.

[`request_id`](/docs/api/network/#network-status-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "network_status": "RETURNING_USER",
  "request_id": "m8MDnv9okwxFNBV"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
