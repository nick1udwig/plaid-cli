---
title: "Errors - Sandbox errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/sandbox/"
scraped_at: "2026-03-07T22:04:53+00:00"
---

# Sandbox Errors

#### Guide to troubleshooting Sandbox errors

#### **SANDBOX\_PRODUCT\_NOT\_ENABLED**

##### The requested product is not enabled for an Item

##### Common causes

- A sandbox operation could not be performed because a product has not been enabled on the Sandbox Item.

##### Troubleshooting steps

Verify that you are enabled for the requested product in your [Dashboard](https://dashboard.plaid.com).

If the error persists, submit a [Support](https://dashboard.plaid.com/support/new/) ticket.

API error response

```
http code 400
{
 "error_type": "SANDBOX_ERROR",
 "error_code": "SANDBOX_PRODUCT_NOT_ENABLED",
 "error_message": "The [auth | transactions | ...] product is not enabled on this item",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **SANDBOX\_WEBHOOK\_INVALID**

##### The request to fire a Sandbox webhook failed.

##### Common causes

- The webhook for the Item sent in the [`/sandbox/item/fire_webhook`](/docs/api/sandbox/#sandboxitemfire_webhook) request is not set or is invalid.

##### Troubleshooting steps

Create a new Item with a valid webhook set.

API error response

```
http code 400
{
 "error_type": "SANDBOX_ERROR",
 "error_code": "SANDBOX_WEBHOOK_INVALID",
 "error_message": "Webhook for this item is either not set up, or invalid. Please update the item's webhook and try again.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **SANDBOX\_TRANSFER\_EVENT\_TRANSITION\_INVALID**

##### The `/sandbox/transfer/simulate` endpoint was called with parameters specifying an invalid event transition.

##### Common causes

- The [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate) endpoint was called with parameters specifying an invalid event transition.

##### Troubleshooting steps

Ensure the sequence of events specified via [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate) is valid.

Compatible status --> event type transitions include:

`pending` --> `failed`

`pending` --> `posted`

`posted` --> `reversed`

API error response

```
http code 400
{
 "error_type": "SANDBOX_ERROR",
 "error_code": "SANDBOX_TRANSFER_EVENT_TRANSITION_INVALID",
 "error_message": "The provided simulated event type is incompatible with the current transfer status",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
