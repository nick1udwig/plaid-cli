---
title: "Identity - Identity Document Upload | Plaid Docs"
source_url: "https://plaid.com/docs/identity/identity-document-upload/"
scraped_at: "2026-03-07T22:04:57+00:00"
---

# Identity Document Upload

#### Learn about statement upload-based account ownership verification

#### Overview

In order to provide the user an additional way to verify their account ownership, Identity Document Upload can be used to upload a bank statement, verify that statement belongs to the account in question, and verify the ownership information on the statement. This feature is intended primarily for use with Items created via loginless Auth flows, such as Same-Day Micro-deposits, Instant Micro-deposits, or Database Insights.

Identity Document Upload is available as an add-on to Identity. For pricing details, and to request access to Identity Document Upload, contact your Plaid Account Manager or Sales.

To detect potential document fraud or document tampering during the Identity Document Upload flow, you can use the optional Fraud Risk feature, which scans for over two dozen different fraud signals in categories such as visual evidence of tampering, suspicious metadata, inconsistent contents, and similarity to known fraudulent documents.

#### Implementation

To use Document Upload, you will first create the Item with another product such as Auth or Transfer. Then, run Link in [update mode](/docs/link/update-mode/) with the following parameters (in addition to the required fields):

- The `products` array set to `["identity"]`.
- The `identity.account_ids` array should contain the `account_id` of the account to verify. Currently, only one `account_id` can be specified.
- `identity.is_document_upload` should be set to `true`.
- (Optional) to enable Fraud Risk, set `identity.parsing_configs` to `["ocr", "risk_signals"]`

Example /link/token/create call for Identity Document Upload

```
curl -X POST https://sandbox.plaid.com/link/token/create \
-H 'Content-Type: application/json' \
-d '{
  "client_id": "${PLAID_CLIENT_ID}",
  "secret": "${PLAID_SECRET}",
  "client_name": "Insert Client name here",
  "products": ["identity"],
  "access_token": "Insert access token here",
  "identity": 
    {
      "is_document_upload": true, 
      "account_ids": ["ZXEbW7Rkr9iv1qj8abebU1KDJlkexgSgrLAod"], 
      "parsing_configs": ["ocr", "risk_signals"]
    },
  "country_codes": ["US"],
  "language": "en",
  "user": {
    "client_user_id": "unique_user_id"
  }
}'
```

During update mode, the end user will be prompted to upload a bank statement. After the statement has been uploaded and processed, Plaid will send an `IDENTITY: DOCUMENT_UPDATE_AVAILABLE` webhook. A `"document_status": "OCR_PROCESSING_COMPLETE"` field in the webhook body indicates that the statement was successfully parsed.

If the parsing is successful, you can call [`/identity/documents/uploads/get`](/docs/api/products/identity/#identitydocumentsuploadsget), which will return the identity data parsed from the document in the same format as [`/identity/get`](/docs/api/products/identity/#identityget). This endpoint will also return a `documents` array -- for any given document in the array, if the `metadata.is_account_number_match` field is `true`, Plaid has verified that the account number on the document matches the account number known to Plaid. If it is `false`, the document does not substantiate the end user's ownership of the account.

Documents array

```
"documents": [
  {
      "document_id": "1d107b7f-91fe-44c8-b8e9-325494addf5d",
      "metadata": {
          "document_type": "BANK_STATEMENT",
          "is_account_number_match": true,
          "last_updated": "2024-01-29T08:06:46Z",
          "uploaded_at": "2024-01-29T08:06:46Z"
      },
      "risk_insights": {
          "risk_signals": [
              {
                  "has_fraud_risk": true,
                  "page_number": 0,
                  "signal_description": "Creation date and modification date do not match",
                  "type": "METADATA_DATES_OUTSIDE_WINDOW"
              },
              {
                  "has_fraud_risk": true,
                  "page_number": 0,
                  "signal_description": "Adobe Acrobat",
                  "type": "SOFTWARE_BLACKLIST"
              }
          ],
          "risk_summary": {
              "risk_score": 100
          }
      }
  }
],
```

If Fraud Risk was enabled, the `document` object will contain a `risk_insights` object, including details about potential risks detected in the uploaded document. The `risk_insights.risk_summary.risk_score` field will contain a score summarizing the risk of the document. If the score is 80 or higher, we recommend treating the account identity as unverified and potentially high risk and sending the user through a manual verification flow.

##### Testing Identity Document Upload

In the Sandbox, Plaid will not parse the uploaded bank statement and will instead return pre-populated test user data. Because of this, Sandbox will proceed much faster than Production; you should make sure your app can handle the asynchronous behavior of Production, in which the user's statement will be verified after they have completed Link, rather than during the Link session.

In Sandbox, you must provide a valid webhook URI in order to upload documents.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
