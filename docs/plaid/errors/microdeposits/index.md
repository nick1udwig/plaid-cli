---
title: "Errors - Microdeposits errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/microdeposits/"
scraped_at: "2026-03-07T22:04:51+00:00"
---

# Micro-deposits Errors

#### Guide to troubleshooting micro-deposits errors

#### **BANK\_TRANSFER\_ACCOUNT\_BLOCKED**

##### The bank transfer could not be completed because a previous transfer involving the same end-user account resulted in an error.

##### Common causes

- Plaid cannot send micro-deposits to verify this account because it has flagged the account as not valid for use. This may happen, for example, because a previous attempt by Plaid to transfer funds to this end user account resulted in an "account frozen" or "invalid account number" error.

##### Troubleshooting steps

Ask the end user to link a different account with Plaid and re-attempt the transfer with the new account.

API error response

```
http code 400
{
 "error_type": "BANK_TRANSFER_ERROR",
 "error_code": "BANK_TRANSFER_ACCOUNT_BLOCKED",
 "error_message": "bank transfer was blocked due to a previous ACH return on this account",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
