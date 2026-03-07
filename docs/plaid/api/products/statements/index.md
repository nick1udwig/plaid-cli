---
title: "API - Statements | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/statements/"
scraped_at: "2026-03-07T22:04:19+00:00"
---

# Statements

#### API reference for Statements endpoints and webhooks

For how-to guidance, see the [Statements documentation](/docs/statements/).

| Endpoint | Description |
| --- | --- |
| [`/statements/list`](/docs/api/products/statements/#statementslist) | Get a list of statements available to download |
| [`/statements/download`](/docs/api/products/statements/#statementsdownload) | Download a single bank statement |
| [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh) | Trigger on-demand statement extractions |

| Webhook Name | Description |
| --- | --- |
| [`STATEMENTS_REFRESH_COMPLETE`](/docs/api/products/statements/#statements_refresh_complete) | Statements refresh completed |

### Endpoints

=\*=\*=\*=

#### `/statements/list`

#### Retrieve a list of all statements associated with an item.

The [`/statements/list`](/docs/api/products/statements/#statementslist) endpoint retrieves a list of all statements associated with an item.

/statements/list

**Request fields**

[`access_token`](/docs/api/products/statements/#statements-list-request-access-token)

requiredstringrequired, string

The access token associated with the Item data is being requested for.

[`client_id`](/docs/api/products/statements/#statements-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/statements/#statements-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/statements/list

```
const listRequest: StatementsListRequest = {
  access_token: access_token,
};
const listResponse = await plaidClient.statementsList(listRequest);
accounts = listResponse.accounts;
statements = listResponse.accounts[0].statements;
```

/statements/list

**Response fields**

[`accounts`](/docs/api/products/statements/#statements-list-response-accounts)

[object][object]

[`account_id`](/docs/api/products/statements/#statements-list-response-accounts-account-id)

stringstring

Plaid's unique identifier for the account.

[`account_mask`](/docs/api/products/statements/#statements-list-response-accounts-account-mask)

stringstring

The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.

[`account_name`](/docs/api/products/statements/#statements-list-response-accounts-account-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself.

[`account_official_name`](/docs/api/products/statements/#statements-list-response-accounts-account-official-name)

stringstring

The official name of the account as given by the financial institution.

[`account_subtype`](/docs/api/products/statements/#statements-list-response-accounts-account-subtype)

stringstring

The subtype of the account. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema).

[`account_type`](/docs/api/products/statements/#statements-list-response-accounts-account-type)

stringstring

The type of account. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema).

[`statements`](/docs/api/products/statements/#statements-list-response-accounts-statements)

[object][object]

The list of statements' metadata associated with this account.

[`statement_id`](/docs/api/products/statements/#statements-list-response-accounts-statements-statement-id)

stringstring

Plaid's unique identifier for the statement.

[`date_posted`](/docs/api/products/statements/#statements-list-response-accounts-statements-date-posted)

nullablestringnullable, string

Date when the statement was posted by the FI, if known  
  

Format: `date`

[`month`](/docs/api/products/statements/#statements-list-response-accounts-statements-month)

integerinteger

Month of the year. Possible values: 1 through 12 (January through December).

[`year`](/docs/api/products/statements/#statements-list-response-accounts-statements-year)

integerinteger

The year of statement.  
  

Minimum: `2010`

[`institution_id`](/docs/api/products/statements/#statements-list-response-institution-id)

stringstring

The Plaid Institution ID associated with the Item.

[`institution_name`](/docs/api/products/statements/#statements-list-response-institution-name)

stringstring

The name of the institution associated with the Item.

[`item_id`](/docs/api/products/statements/#statements-list-response-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`request_id`](/docs/api/products/statements/#statements-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
  "institution_id": "ins_3",
  "institution_name": "Chase",
  "accounts": [
    {
      "account_id": "3gE5gnRzNyfXpBK5wEEKcymJ5albGVUqg77gr",
      "account_mask": "0000",
      "account_name": "Plaid Saving",
      "account_official_name": "Plaid Silver Standard 0.1% Interest Saving",
      "account_subtype": "savings",
      "account_type": "depository",
      "statements": [
        {
          "statement_id": "vzeNDwK7KQIm4yEog683uElbp9GRLEFXGK98D",
          "month": 5,
          "year": 2023,
          "date_posted": "2023-05-01"
        }
      ]
    }
  ],
  "request_id": "eYupqX1mZkEuQRx"
}
```

=\*=\*=\*=

#### `/statements/download`

#### Retrieve a single statement.

The [`/statements/download`](/docs/api/products/statements/#statementsdownload) endpoint retrieves a single statement PDF in binary format. The response will contain a `Plaid-Content-Hash` header containing a SHA 256 checksum of the statement. This can be used to verify that the file being sent by Plaid is the same file that was downloaded to your system.

/statements/download

**Request fields**

[`access_token`](/docs/api/products/statements/#statements-download-request-access-token)

requiredstringrequired, string

The access token associated with the Item data is being requested for.

[`client_id`](/docs/api/products/statements/#statements-download-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/statements/#statements-download-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`statement_id`](/docs/api/products/statements/#statements-download-request-statement-id)

requiredstringrequired, string

Plaid's unique identifier for the statement.

/statements/download

```
let downloadRequest: StatementsDownloadRequest = {
  access_token: accessToken,
  statement_id: statement.statement_id,
};
let downloadResponse = await plaidClient.statementsDownload(
  downloadRequest,
  {responseType: 'arraybuffer'},
);
let pdf = downloadResponse.data.toString('base64');
```

##### Response

This endpoint returns a single statement, exactly as provided by the financial institution, in the form of binary PDF data.

=\*=\*=\*=

#### `/statements/refresh`

#### Refresh statements data.

[`/statements/refresh`](/docs/api/products/statements/#statementsrefresh) initiates an on-demand extraction to fetch the statements for the provided dates.

/statements/refresh

**Request fields**

[`access_token`](/docs/api/products/statements/#statements-refresh-request-access-token)

requiredstringrequired, string

The access token associated with the Item data is being requested for.

[`client_id`](/docs/api/products/statements/#statements-refresh-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/statements/#statements-refresh-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`start_date`](/docs/api/products/statements/#statements-refresh-request-start-date)

requiredstringrequired, string

The start date for statements, in "YYYY-MM-DD" format, e.g. "2023-08-30". To determine whether a statement falls within the specified date range, Plaid will use the statement posted date. The statement posted date is typically either the last day of the statement period, or the following day.  
  

Format: `date`

[`end_date`](/docs/api/products/statements/#statements-refresh-request-end-date)

requiredstringrequired, string

The end date for statements, in "YYYY-MM-DD" format, e.g. "2023-10-30". You can request up to two years of data. To determine whether a statement falls within the specified date range, Plaid will use the statement posted date. The statement posted date is typically either the last day of the statement period, or the following day.  
  

Format: `date`

/statements/refresh

```
const refreshRequest: StatementsRefreshRequest = {
  access_token: accessToken,
  start_date: '2023-11-01',
  end_date: '2024-02-01',
};
const refreshResponse = await plaidClient.statementsRefresh(refreshRequest);
```

/statements/refresh

**Response fields**

[`request_id`](/docs/api/products/statements/#statements-refresh-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "request_id": "eYupqX1mZkEuQRx"
}
```

### Webhooks

Statement webhooks are sent to indicate that statements refresh has finished processing. All webhooks related to statements have a `webhook_type` of `STATEMENTS`.

=\*=\*=\*=

#### `STATEMENTS_REFRESH_COMPLETE`

Fired when refreshed statements extraction is completed or failed to be completed. Triggered by calling [`/statements/refresh`](/docs/api/products/statements/#statementsrefresh).

**Properties**

[`webhook_type`](/docs/api/products/statements/#StatementsRefreshCompleteWebhook-webhook-type)

stringstring

`STATEMENTS`

[`webhook_code`](/docs/api/products/statements/#StatementsRefreshCompleteWebhook-webhook-code)

stringstring

`STATEMENTS_REFRESH_COMPLETE`

[`item_id`](/docs/api/products/statements/#StatementsRefreshCompleteWebhook-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`result`](/docs/api/products/statements/#StatementsRefreshCompleteWebhook-result)

stringstring

The result of the statement refresh extraction  
`SUCCESS`: The statements were successfully extracted and can be listed via `/statements/list/` and downloaded via `/statements/download/`.  
`FAILURE`: The statements failed to be extracted.  
  

Possible values: `SUCCESS`, `FAILURE`

[`environment`](/docs/api/products/statements/#StatementsRefreshCompleteWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "STATEMENTS",
  "webhook_code": "STATEMENTS_REFRESH_COMPLETE",
  "item_id": "eVBnVMp7zdTJLkRNr33Rs6zr7KNJqBFL9DrE6",
  "result": "SUCCESS",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
