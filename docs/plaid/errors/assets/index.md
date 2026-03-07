---
title: "Errors - Assets errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/assets/"
scraped_at: "2026-03-07T22:04:49+00:00"
---

# Assets Errors

#### Guide to troubleshooting Assets errors

#### **PRODUCT\_NOT\_ENABLED**

##### Common causes

- One or more of the Items in the Asset Report was not initialized with the Assets product. Unlike some products, Assets cannot be initialized "on-the-fly" and must be initialized during the initial link process.

##### Troubleshooting steps

Make sure to include `assets` in the list of products to initialize the Item for during Link, then have your user re-link the Item.

API error response

```
http code 400
{
 "error_type": "ASSET_REPORT_ERROR",
 "error_code": "PRODUCT_NOT_ENABLED",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **DATA\_UNAVAILABLE**

##### Common causes

- One or more of the Items in the Asset Report has experienced an error.

##### Troubleshooting steps

Check the `causes` field for a detailed breakdown of errors, then follow the troubleshooting steps for any errors found.

If the `causes` field is not present, use [`/item/get`](/docs/api/items/#itemget) to query the Items in the Asset Report for errors, then follow the troubleshooting steps for any errors found.

API error response

```
http code 400
{
 "error_type": "ASSET_REPORT_ERROR",
 "error_code": "DATA_UNAVAILABLE",
 "error_message": "unable to pull sufficient data for all items to generate this report. please see the \"causes\" field for more details",
 "display_message": null,
 "causes": [
  {
   "display_message": null,
   "error_code": "ITEM_LOGIN_REQUIRED",
   "error_message": "the login details of this item have changed (credentials, MFA, or required user action) and a user login is required to update this information. use Link's update mode to restore the item to a good state",
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

- [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) or [`/asset_report/pdf/get`](/docs/api/products/assets/#asset_reportpdfget) was called before the Asset Report generation completed.

##### Troubleshooting steps

Listen for the [`PRODUCT_READY`](/docs/api/products/assets/#product_ready) webhook and wait to call [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) or [`/asset_report/pdf/get`](/docs/api/products/assets/#asset_reportpdfget) until that webhook has been fired.

API error response

```
http code 400
{
 "error_type": "ASSET_REPORT_ERROR",
 "error_code": "PRODUCT_NOT_READY",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ASSET\_REPORT\_GENERATION\_FAILED**

##### Common causes

- Plaid is experiencing temporary difficulties generating Asset Reports.
- The Asset Report was too large to generate.

##### Troubleshooting steps

Check the `error_message` field for details on the cause of the error.

If the `error_message` indicates the Asset Report was too large, try requesting a shorter date range in the Asset Report.

Try creating the Asset Report again later.

If the error persists, please submit a [Support](https://dashboard.plaid.com/support/new/) ticket.

API error response

```
http code 400
{
 "error_type": "ASSET_REPORT_ERROR",
 "error_code": "ASSET_REPORT_GENERATION_FAILED",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_PARENT**

##### Common causes

- An endpoint that creates a copy or modified copy of an Asset Report (such as [`/asset_report/filter`](/docs/api/products/assets/#asset_reportfilter) or [`/asset_report/audit_copy/create`](/docs/api/products/assets/#asset_reportaudit_copycreate)) was called, but the original Asset Report does not exist, either because it was never successfully created in the first place or because it was deleted.

Re-create the original Asset Report and re-try the endpoint call on the new Asset Report.

API error response

```
http code 400
{
 "error_type": "ASSET_REPORT_ERROR",
 "error_code": "INVALID_PARENT",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSIGHTS\_NOT\_ENABLED**

##### Common causes

- [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) or [`/asset_report/pdf/get`](/docs/api/products/assets/#asset_reportpdfget) was called with `enable_insights` set to `true` by a Plaid developer account that has not been enabled for access to Asset Reports with Insights.

[Contact sales](https://plaid.com/contact) to enable Asset Reports with Insights for your account.

API error response

```
http code 400
{
 "error_type": "ASSET_REPORT_ERROR",
 "error_code": "INSIGHTS_NOT_ENABLED",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSIGHTS\_PREVIOUSLY\_NOT\_ENABLED**

##### Common causes

- [`/asset_report/get`](/docs/api/products/assets/#asset_reportget) or [`/asset_report/pdf/get`](/docs/api/products/assets/#asset_reportpdfget) was called with `enable_insights` set to `true` by a Plaid developer account that is currently enabled for Asset Reports with Insights, but was not enabled at the time that the report was created.

Re-create the Asset Report with [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) and then fetch the new Asset Report.

API error response

```
http code 400
{
 "error_type": "ASSET_REPORT_ERROR",
 "error_code": "INSIGHTS_PREVIOUSLY_NOT_ENABLED",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **DATA\_QUALITY\_CHECK\_FAILED**

##### Common causes

- The extracted data from the financial institution showed signs of poor data quality, preventing the generation of an accurate report.

Try creating the Asset Report again 1-2 days later.

API error response

```
http code 400
{
 "error_type": "ASSET_REPORT_ERROR",
 "error_code": "DATA_QUALITY_CHECK_FAILED",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
