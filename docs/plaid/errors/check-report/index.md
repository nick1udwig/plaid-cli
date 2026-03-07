---
title: "Errors - Check Report errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/check-report/"
scraped_at: "2026-03-07T22:04:49+00:00"
---

# Check Report Errors

#### Guide to troubleshooting Check Report errors

#### **CONSUMER\_REPORT\_EXPIRED**

##### Common causes

- The Check Report has expired as 24 hours have passed since its creation, so it is no longer retrievable.

##### Troubleshooting steps

Create a new report for the user by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).

API error response

```
http code 400
{
 "error_type": "CONSUMER_REPORT_ERROR",
 "error_code": "CONSUMER_REPORT_EXPIRED",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **DATA\_UNAVAILABLE**

##### Common causes

- Plaid Check failed to retrieve data from the user's financial institutions to generate Check Report product data.

##### Troubleshooting steps

Check the `causes` field for a detailed breakdown of errors, then follow the troubleshooting steps for any errors found.

If the `causes` field is not present, use `user/item/get` to query the Items in the Asset Report for errors, then follow the troubleshooting steps for any errors found.

API error response

```
http code 400
{
 "error_type": "CHECK_REPORT_ERROR",
 "error_code": "DATA_UNAVAILABLE",
 "error_message": "The Check Report you are trying to pull does not have sufficient transactions data to generate a report.",
 "display_message": null,
 "causes": [
  {
   "display_message": null,
   "error_code": "ITEM_LOGIN_REQUIRED",
   "error_message": "The login details of this item have changed (credentials, MFA, or required user action) and a user login is required to update this information. use Link's update mode to restore the item to a good state",
   "error_type": "ITEM_ERROR",
   "item_id": "pZ942ZA0npFEa0BgLCV9DAQv3Zq8ErIZhc81F"
  }
 ],
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PRODUCT\_NOT\_READY**

Note that this error should not be confused with the [`PRODUCT_NOT_READY`](/docs/errors/item/#product_not_ready) error of type `ITEM_ERROR`, which has different causes.

##### Common causes

- `/check_report/.../get` was called before report generation completed.

##### Troubleshooting steps

Listen for the `CHECK_REPORT_READY` or `CHECK_REPORT_FAILED` webhooks and wait to call `/check_report/.../get` until that webhook has been fired.

API error response

```
http code 400
{
 "error_type": "CONSUMER_REPORT_ERROR",
 "error_code": "PRODUCT_NOT_READY",
 "error_message": "The consumer report you are trying to pull is not ready. Please wait for a CHECK_REPORT_READY webhook to fetch data",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTITUTION\_TRANSACTION\_HISTORY\_NOT\_SUPPORTED**

##### Common causes

- The user's financial institution does not support the length of transaction history you require

##### Troubleshooting steps

Lower the `days_required` to be compatible with the financial institution's limits.

API error response

```
http code 400
{
 "error_type": "CHECK_REPORT_ERROR",
 "error_code": "INSTITUTION_TRANSACTION_HISTORY_NOT_SUPPORTED",
 "error_message": "The user’s bank does not support the transactions range you require. Lowering the days_required may result in success.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSUFFICIENT\_TRANSACTION\_DATA**

##### Common causes

- The user does not have sufficient transactions data to generate income insights, partner insights, or cashflow insights. The `error_message` will provide more detail on what data is missing.

##### Troubleshooting steps

Have the user connect a different financial institutions that will satisfy the transaction requirements.

API error response

```
http code 400
{
 "error_type": "CHECK_REPORT_ERROR",
 "error_code": "INSUFFICIENT_TRANSACTION_DATA",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NO\_ACCOUNTS**

##### Common causes

- The user does not the correct accounts linked to generate income insights. The `error_message` will provide more detail on what data is missing.

##### Troubleshooting steps

Have the user connect a different financial institutions that will satisfy the transaction requirements.

API error response

```
http code 400
{
 "error_type": "CHECK_REPORT_ERROR",
 "error_code": "NO_ACCOUNTS",
 "error_message": "No depository accounts were found for this user",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NETWORK\_CONSENT\_REQUIRED**

##### Common causes

- The user did not provide consent in link to share Plaid network data.

##### Troubleshooting steps

Have the user go through link again with `cra_network_insights` or `cra_cashflow_insights` so they can give consent to share network data

API error response

```
http code 400
{
 "error_type": "CHECK_REPORT_ERROR",
 "error_code": "NETWORK_CONSENT_REQUIRED",
 "error_message": "User has not provided consent to share network data",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **DATA\_QUALITY\_CHECK\_FAILED**

##### Common causes

- The extracted data from the financial institution showed signs of poor data quality, preventing the generation of an accurate report.

##### Troubleshooting steps

Try creating the Check Report again 1-2 days later.

API error response

```
http code 400
{
 "error_type": "CHECK_REPORT_ERROR",
 "error_code": "DATA_QUALITY_CHECK_FAILED",
 "error_message": "Bank provided inconsistent transaction history data that has a high chance of error. The bank will need to resolve issues for product data to be generated.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
