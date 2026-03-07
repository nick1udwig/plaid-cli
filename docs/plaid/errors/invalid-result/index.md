---
title: "Errors - Invalid Result errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/invalid-result/"
scraped_at: "2026-03-07T22:04:50+00:00"
---

# Invalid Result errors

#### Guide to troubleshooting invalid result errors

#### **LAST\_UPDATED\_DATETIME\_OUT\_OF\_RANGE**

##### No data is available from the Item within the specified date-time.

##### Common causes

- [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) was called with a parameter specifying the minimum acceptable data freshness, but no balance data meeting those requirements was available from the institution.

##### Troubleshooting steps

This error is not preventable by developer actions. As a workaround, you can use a cached balance from [`/accounts/get`](/docs/api/accounts/#accountsget).

API error response

```
http code 400
{
 "error_type": "INVALID_RESULT",
 "error_code": "LAST_UPDATED_DATETIME_OUT_OF_RANGE",
 "error_message": "requested datetime out of range, most recently updated balance 2021-01-01T00:00:00Z",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PLAID\_DIRECT\_ITEM\_IMPORT\_RETURNED\_INVALID\_MFA**

##### The Plaid Direct Item import resulted in invalid MFA.

##### Common causes

- No known causes.

##### Troubleshooting steps

If you experience this error, contact your Plaid Account Manager or [file a support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

API error response

```
http code 400
{
 "error_type": "INVALID_RESULT",
 "error_code": "PLAID_DIRECT_ITEM_IMPORT_RETURNED_INVALID_MFA",
 "error_message": "Plaid Direct Item Imports should not result in MFA.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
