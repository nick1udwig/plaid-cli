---
title: "Errors - User errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/user/"
scraped_at: "2026-03-07T22:04:53+00:00"
---

# User Errors

#### Guide to troubleshooting User errors

#### **USER\_NOT\_FOUND**

##### The provided `user_id` could not be found.

##### Common causes

- The `user_id` was never created.
- The `user_id` was created in a different environment (Sandbox vs Production).
- The `user_id` was previously deleted via [`/user/remove`](/docs/api/users/#userremove).

##### Troubleshooting steps

Ensure you're using the correct API environment (Sandbox or Production) that matches where the `user_id` was created.

Confirm the `user_id` format is correct (should start with `usr_`)

API error response

```
http code 400
{
 "error_type": "USER_ERROR",
 "error_code": "USER_NOT_FOUND",
 "error_message": "The User you requested cannot be found. This User does not exist or has had access removed by the user.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
