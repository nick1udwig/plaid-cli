---
title: "Income - Document Income | Plaid Docs"
source_url: "https://plaid.com/docs/income/document-income/"
scraped_at: "2026-03-07T22:04:58+00:00"
---

# Document Income

#### Learn about Document Income features and implementation

Get started with Document Income

[API Reference](/docs/api/products/income/)[Quickstart](https://github.com/plaid/income-sample)

#### Overview

Document Income allows you to retrieve gross and/or net income information via user-uploaded documents such as a pay stub, bank statement, W-2, or 1099 form.

Prefer to learn by watching? Get an overview of how Income works in just 3 minutes!

#### Integration process

1. (Optional) Update your [Link customization](https://dashboard.plaid.com/link/manualVerificationOfIncomeUpload) in the Dashboard to configure which types of documents a user can upload to verify their income. By default, if you do not change this setting, only pay stubs are accepted.
2. Call [`/user/create`](/docs/api/users/#usercreate) to create a `user_token` that will represent the end user interacting with your application. This step will only need to be done once per end user. If you are using multiple Income types, do not repeat this step when switching to a different Income type.
3. Call [`/link/token/create`](/docs/api/link/#linktokencreate). In addition to the required parameters, you will need to provide the following:
   - For `user_token`, provide the `user_token` from [`/user/create`](/docs/api/users/#usercreate).
   - For `products`, use `["income_verification"]`. Document Income cannot be used in the same Link session as any other Plaid products, except for Payroll Income.
   - For `income_verification.income_source_types`, use `payroll`.
   - (Optional) If you are only using Document Income and do not want customers to use Payroll Income, for `income_verification.payroll_income.flow_types`, use `["payroll_document_income"]`.
   - Provide a `webhook` URI with the endpoint where you will receive Plaid webhooks.
   - If using Fraud Risk, set `income_verification.payroll_income.parsing_config` to either `['risk_signals']` or `['risk_signals', 'ocr']`. For more details, see [Fraud Risk detection](/docs/income/document-income/#fraud-risk-detection).
4. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see [Link](/docs/link/).
5. Open Link in your web or mobile client and listen to the `onSuccess` and `onExit` callbacks, which will fire once the user has finished or exited the Link session.
6. If using Fraud Risk, wait for the [`INCOME: INCOME_VERIFICATION_RISK_SIGNALS`](/docs/api/products/income/#income_verification_risk_signals) webhook, then call [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget).
   - If the document requires manual review, call [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) to get a URI where the user's original documents can be downloaded.
7. If using Document Parsing, wait for the [`INCOME: INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook, then call [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) and/or [`/credit/bank_statements/uploads/get`](/docs/api/products/income/#creditbank_statementsuploadsget) to obtain parsed income details. You can use [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget) to see what document types were uploaded, which will determine which of the two endpoints to call. For details, see [Document Parsing](/docs/income/document-income/#document-parsing).

#### Fraud risk detection

To detect potential document fraud or document tampering during the Document Income flow, you can use the optional Fraud Risk feature. Fraud Risk's AI-powered analysis scans for over two dozen different fraud signals within categories such as: visual evidence of tampering, suspicious metadata, inconsistent contents, and similarity to known fraudulent documents. If your account is not enabled for this feature, contact sales or your Plaid Account Manager to request access.

To enable Fraud Risk:

1. When calling [`/link/token/create`](/docs/api/link/#linktokencreate), set the `income_verification.payroll_income.parsing_config` array to include `'risk_signals'`. By default, this will disable [Document Parsing](/docs/income/document-income/#document-parsing); to keep it enabled, include `'ocr'` in the array as well.
2. When the risk analysis has been completed, you will receive an [`INCOME_VERIFICATION_RISK_SIGNALS`](/docs/api/products/income/#income_verification_risk_signals) webhook. This webhook may take up to 45 minutes to fire.
3. Once the webhook has been received, call the [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) endpoint.

[`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget) will return a risk score for each document, as well as a detailed set of reasons for any potential risk.

You can automatically reject documents with a high risk score, automatically accept documents with a low risk score, and manually review documents in between. We suggest starting with a threshold of 80 for auto-rejection and 20 for auto-acceptance. As you gather more data points on typical risk scores for your use case, you can tune these parameters to reduce the number of documents undergoing manual review.

To obtain a copy of the original document for manual review, call [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) and use the `document_metadata.download_url`.

To enable Fraud Risk on a verification where the Link flow has already been completed, use the [`/credit/payroll_income/parsing_config/update`](/docs/api/products/income/#creditpayroll_incomeparsing_configupdate) endpoint.

#### Document Parsing

Document Parsing is an optional feature that allows you to obtain a JSON representation of an uploaded document. If your account is not enabled for this feature, contact sales or your Plaid Account Manager to request access.

To enable Document Parsing:

1. Call [`/link/token/create`](/docs/api/link/#linktokencreate) as normal, then launch Link and have the user go through the Link flow. You do not need to specify a `parsing_config` when calling [`/link/token/create`](/docs/api/link/#linktokencreate), as Document Parsing will be enabled by default if this field is omitted. However, if this field is supplied (for example, to enable Fraud Risk), it must include `ocr` to enable Document Parsing.
2. To see which file types a user uploaded, use [`/credit/sessions/get`](/docs/api/products/income/#creditsessionsget). The `document_income_results` field will show how many of each filetype were uploaded.
3. Wait for document parsing to complete, which will be indicated by the [`INCOME_VERIFICATION`](/docs/api/products/income/#income_verification) webhook. This webhook may take up to 45 minutes to fire.
4. Once the webhook has been received, to obtain parsed JSON data from a pay stub, W-2, or 1099, use [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget). To obtain parsed JSON data from a bank statement, use [`/credit/bank_statements/uploads/get`](/docs/api/products/income/#creditbank_statementsuploadsget).

To enable Document Parsing on a verification where the Link flow has already been completed, use the [`/credit/payroll_income/parsing_config/update`](/docs/api/products/income/#creditpayroll_incomeparsing_configupdate) endpoint.

#### Downloading original documents

To download the original user-uploaded documents, use the `document_metadata.download_url` returned by [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget).

#### Customizing document uploads

To configure which documents your users can upload, create a [Link customization in the Dashboard](https://dashboard.plaid.com/link). You must update this setting in order to accept documents other than pay stubs.

Bank statements are only supported with a custom contract. If you are not contracted for bank statements, do not enable bank statement uploads in your Link customization, as it will cause your [`/link/token/create`](/docs/api/link/#linktokencreate) calls to fail. To request access to bank statement uploads, contact Sales or your Account Manager.

##### Supported document types

- Pay stubs
- Bank statements (requires custom contract)
- W-2
- 1099-K
- 1099-MISC

##### Supported file types

- PDF
- PNG
- JPEG
- GIF
- BMP
- TIFF

For more details, see [Document upload settings](/docs/link/customization/#document-income-only-document-upload-settings).

#### Testing Document Income

In the Sandbox environment, when testing Document Income, by default, the contents of the actual document will not be processed and Sandbox will instead use pre-populated test data. You can customize the response you get by uploading JSON files with a custom configuration schema in the Link flow. Each JSON file will represent one document in the response and can include customizations for:

- Either a paystub or a W-2
- Risk signals
- OCR parsing status
- Risk signals status

The customization schema for each document and risk signals is very similar to the objects in the responses from [`/credit/payroll_income/get`](/docs/api/products/income/#creditpayroll_incomeget) and [`/credit/payroll_income/risk_signals/get`](/docs/api/products/income/#creditpayroll_incomerisk_signalsget). Example JSON files that you can adapt can be found in the [Custom Sandbox User GitHub repo](https://github.com/plaid/sandbox-custom-users/tree/main/income/document_income).

#### Document Income pricing

Document Income is billed on a [one-time fee model](/docs/account/billing/#one-time-fee). The fee depends on the number and type of documents processed and which processing options are enabled (i.e. fraud risk, document parsing, or both).

To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

#### Next steps

If you're ready to launch to Production, see the Launch checklist.

[#### Launch Center

See next steps to launch in Production

Launch](https://dashboard.plaid.com/developers/launch-center)

#### Launch Center

See next steps to launch in Production

[Launch](https://dashboard.plaid.com/developers/launch-center)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
