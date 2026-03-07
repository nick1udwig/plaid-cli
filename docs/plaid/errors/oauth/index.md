---
title: "Errors - OAuth errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/oauth/"
scraped_at: "2026-03-07T22:04:51+00:00"
---

# OAuth Errors

#### Guide to troubleshooting OAuth errors

For information on troubleshooting OAuth errors in Link not related to a specific error code, see [Link troubleshooting: OAuth not working](/docs/link/troubleshooting/#oauth-not-working).

#### **INCORRECT\_OAUTH\_NONCE**

##### An incorrect OAuth nonce was supplied when re-initializing Link.

##### Common causes

- During the OAuth flow, Link must be initialized, the user must be handed off to the institution's OAuth authorization page, and then Link must be re-initialized for the user to complete Link flow. This error can occur if a different nonce is supplied during the re-initialization process than was originally supplied to Link for the first initialization step.

##### Troubleshooting steps

When re-initializing Link, make sure to use the same nonce that was used to originally initialize Link for that Item.

API error response

```
http code 400
{
 "error_type": "OAUTH_ERROR",
 "error_code": "INCORRECT_OAUTH_NONCE",
 "error_message": "nonce does not match",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INCORRECT\_LINK\_TOKEN**

##### An incorrect Link token was supplied when re-initializing Link.

##### Common causes

- During the OAuth flow, Link must be initialized, the user must be handed off to the institution's OAuth authorization page, and then Link must be re-initialized for the user to complete Link flow. This error can occur if a different `link_token` is supplied during the re-initialization process than was originally supplied to Link for the first initialization step.

##### Troubleshooting steps

When re-initializing Link, make sure to use the same `link_token` that was used to originally initialize Link for that Item.

API error response

```
http code 400
{
 "error_type": "OAUTH_ERROR",
 "error_code": "INCORRECT_LINK_TOKEN",
 "error_message": "link token does not match original link token",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **OAUTH\_STATE\_ID\_ALREADY\_PROCESSED**

##### The OAuth state has already been processed.

##### Common causes

- During the OAuth flow, Link must be initialized, the user must be handed off to the institution's OAuth authorization page, and then Link must be re-initialized for the user to complete Link flow. This error can occur if the OAuth state ID used during re-initialization of Link has already been used.

##### Troubleshooting steps

When re-initializing Link, make sure to use a unique OAuth state ID for each Link instance.

When re-initializing Link, make sure to correctly set the `receivedRedirectUri` as described in the [re-initializing Link](https://plaid.com/docs/link/oauth/#reinitializing-link) section of the OAuth Guide. Plaid will automatically extract the OAuth state ID from the `receivedRedirectUri`.

API error response

```
http code 208
{
 "error_type": "OAUTH_ERROR",
 "error_code": "OAUTH_STATE_ID_ALREADY_PROCESSED",
 "error_message": null,
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **OAUTH\_STATE\_ID\_NOT\_FOUND**

##### The OAuth state id could not be found.

##### Common causes

- An internal Plaid error occurred.

##### Troubleshooting steps

[File a Plaid Support ticket](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/oauth-issues).

API error response

```
http code 404
{
 "error_type": "OAUTH_ERROR",
 "error_code": "OAUTH_STATE_ID_NOT_FOUND",
 "error_message": "the provided oauth_state_id wasn't found",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
