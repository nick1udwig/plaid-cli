---
title: "Errors - Transactions errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/transactions/"
scraped_at: "2026-03-07T22:04:53+00:00"
---

# Transactions errors

#### Guide to troubleshooting Transactions errors

#### **TRANSACTIONS\_SYNC\_MUTATION\_DURING\_PAGINATION**

##### Transaction data was updated during pagination.

##### Common causes

- Transaction data was updated during pagination. This can occur when calling the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) endpoint.

##### Troubleshooting steps

Restart the pagination loop beginning at the original cursor value used for the first request in the pagination loop. Note that the entire loop must be restarted, beginning with the first request whose response included a `true` value for `has_more`. If you attempt to re-run only the requests using the cursor value where this error was encountered, without restarting the entire loop, the error will recur. The request loop must be run until `has_more` is `false`.

To reduce the frequency of this error, you can increase the `count` parameter in the [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) from the default value of 100 to a higher value, such as the maximum of 500, which will reduce the number of pages in the API response.

API error response

```
http code 400
{
 "error_type": "TRANSACTIONS_ERROR",
 "error_code": "TRANSACTIONS_SYNC_MUTATION_DURING_PAGINATION",
 "error_message": "Underlying transaction data changed since last page was fetched. Please restart pagination from last update.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
