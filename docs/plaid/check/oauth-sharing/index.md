---
title: "Consumer Report (by Plaid Check) - Share data with partners | Plaid Docs"
source_url: "https://plaid.com/docs/check/oauth-sharing/"
scraped_at: "2026-03-07T22:04:48+00:00"
---

# Sharing Consumer Report data with partners

#### Share Consumer Report data with approved partners using OAuth tokens

#### Overview

For eligible product use cases, Plaid Check provides the ability to share Consumer Reporting Agency (CRA) report information with approved partners (Fannie Mae, Freddie Mac, or Experian) using OAuth.

To share CRA report information with partners, you need to create an OAuth access token associated with a previously created `user_id` that your partner can use to access the information.

#### Creating an OAuth token

To create an access token that will be provided to your partner, call [`/oauth/token`](/docs/api/oauth/#oauthtoken) following the example below.

For the `subject_token`, provide the `user_id` that you created using [`/user/create`](/docs/api/users/#usercreate).

The `audience` must match the partner you are sharing with. Valid values are:

- `urn:plaid:params:cra-partner:experian`
- `urn:plaid:params:cra-partner:fannie-mae`
- `urn:plaid:params:cra-partner:freddie-mac`

You can also provide multiple comma-separated values to create a token that will work with multiple partners. For example: `"audience": "urn:plaid:params:cra-partner:experian,urn:plaid:params:cra-partner:fannie-mae"`

/oauth/token/ request

```
curl 'https://sandbox.plaid.com/oauth/token' \
--header 'Content-Type: application/json' \
--data '{
  "client_id": "${PLAID_CLIENT_ID}",
  "secret": "${PLAID_SECRET}",
  "grant_type": "urn:ietf:params:oauth:grant-type:token-exchange",
  "scope": "user:read",
  "subject_token_type":"urn:plaid:params:credit:multi-user",
  "audience": "urn:plaid:params:cra-partner:fannie-mae",
  "subject_token": "usr_9nSp2KuZ2x4JDw"
}'
```

/oauth/token response

```
{
  "access_token": "pda-RDdg0TUCB0FB25_UPIlnhA==",
  "expires_in": 2591999,
  "refresh_token": "Rpdr--viXurkDg88d5zf8m6Wl0g==",
  "request_id": "vPamXI8hYXPl7P2",
  "token_type": "Bearer"
}
```

You can then provide the `access_token` to your partner.

#### Revoking an OAuth token

To revoke your token, call [`/oauth/revoke`](/docs/api/oauth/#oauthrevoke), passing in the `access_token` and/or `refresh_token` as the `token` value.

/oauth/revoke request

```
curl -X POST https://sandbox.plaid.com/oauth/revoke \
-H 'Content-Type: application/json' \
-d '{
  "client_id": "${PLAID_CLIENT_ID}",
  "secret": "${PLAID_SECRET}",
  "token": "pda-RDdg0TUCB0FB25_UPIlnhA=="
}'
```

/oauth/revoke response

```
{
  "request_id": "pCDVCQK8ve2MzhM"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
