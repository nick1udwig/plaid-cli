---
title: "Errors - Overview | Plaid Docs"
source_url: "https://plaid.com/docs/errors/"
scraped_at: "2026-03-07T22:04:49+00:00"
---

# Errors

#### A comprehensive breakdown of all Plaid error codes

=\*=\*=\*=

#### Most common errors

The following are the most common errors that may occur in response to an API call even if your implementation is correct. This list of common errors excludes Item errors that occur only during the Link flow (typically due to bad data entry by the end user), such as [`INVALID_CREDENTIALS`](/docs/errors/item/#invalid_credentials), and errors that can occur only due to sending invalid input, such as [`INVALID_FIELD`](/docs/errors/invalid-request/#invalid_field). It is recommended that your integration should, at minimum, handle each of the errors below that are applicable to your product and/or integration mode.

In the table below, "institution-based products" refers to any product or integration that connects to a bank or other financial institution (i.e., most Plaid products); it excludes products such as Identity Verification, Monitor, Enrich, and Document Income that do not involve making a connection to a financial institution.

| Error | Applies to | Summary |
| --- | --- | --- |
| [`ITEM_LOGIN_REQUIRED`](/docs/errors/item/#item_login_required) | All institution-based products | Item has expired credentials or consent |
| [`PRODUCT_NOT_READY`](/docs/errors/item/#product_not_ready) | Signal, Assets, Income, Check, Auth, [`/transactions/get`](/docs/api/products/transactions/#transactionsget) | Plaid hasn't finished obtaining the data needed to fulfill your request |
| [`PRODUCTS_NOT_SUPPORTED`](/docs/errors/item/#products_not_supported) | All institution-based products | The product endpoint isn't compatible with this Item |
| [`TRANSACTIONS_SYNC_MUTATION_DURING_PAGINATION`](/docs/errors/transactions/#transactions_sync_mutation_during_pagination) | [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) | An update was received during Transactions pagination |
| [`NO_ACCOUNTS`](/docs/errors/item/#no_accounts) | All institution-based products | Couldn't find any open accounts |
| [`NO_AUTH_ACCOUNTS`](/docs/errors/item/#no_auth_accounts-or-no-depository-accounts) | Auth | Couldn't find any debitable checking, savings, or cash management accounts |
| [`NO_LIABILITY_ACCOUNTS`](/docs/errors/item/#no_liability_accounts) | Liabilities | Couldn't find any credit accounts |
| [`NO_INVESTMENT_ACCOUNTS`](/docs/errors/item/#no_investment_accounts) | Investments | Couldn't find any investment accounts |
| [`ACCESS_NOT_GRANTED`](/docs/errors/item/#access_not_granted) | All institution-based products | The end user didn't grant an OAuth permission required for your request |
| [`ADDITIONAL_CONSENT_REQUIRED`](/docs/errors/invalid-input/#additional_consent_required) | Integrations in the US or Canada that [add products to Items after Link](/docs/link/initializing-products/#adding-products-post-link) | The end user didn't grant a data scope required for your request |
| [`INSTITUTION_NOT_RESPONDING`](/docs/errors/institution/#institution_not_responding) | All institution-based products | Temporary financial institution connectivity outage |
| [`INSTITUTION_DOWN`](/docs/errors/institution/#institution_down) | All institution-based products | Temporary financial institution connectivity outage |
| [`RATE_LIMIT_EXCEEDED`](/docs/errors/rate-limit-exceeded/) | Applications that batch-process Plaid API calls or have heavy traffic spikes | Too many requests made too quickly, or Limited Production caps hit |
| [`INTERNAL_SERVER_ERROR`](/docs/errors/api/#internal_server_error-or-plaid-internal-error) | All products | Internal error or financial institution error not otherwise specified |

#### Errors overview

|  |  |
| --- | --- |
| Item errors  Occur when an Item may be invalid or not supported on Plaid's platform. | [ACCESS\_NOT\_GRANTED](/docs/errors/item/#access_not_granted)  [INSTANT\_MATCH\_FAILED](/docs/errors/item/#instant_match_failed)  [INSUFFICIENT\_CREDENTIALS](/docs/errors/item/#insufficient_credentials)  [INVALID\_CREDENTIALS](/docs/errors/item/#invalid_credentials)  [INVALID\_MFA](/docs/errors/item/#invalid_mfa)  [INVALID\_OTP](/docs/errors/item/#invalid_otp)  [INVALID\_PHONE\_NUMBER](/docs/errors/item/#invalid_phone_number)  [INVALID\_SEND\_METHOD](/docs/errors/item/#invalid_send_method)  [INVALID\_UPDATED\_USERNAME](/docs/errors/item/#invalid_updated_username)  [ITEM\_CONCURRENTLY\_DELETED](/docs/errors/item/#item_concurrently_deleted)  [ITEM\_LOCKED](/docs/errors/item/#item_locked)  [ITEM\_LOGIN\_REQUIRED](/docs/errors/item/#item_login_required)  [ITEM\_NOT\_FOUND](/docs/errors/item/#item_not_found)  [ITEM\_NOT\_SUPPORTED](/docs/errors/item/#item_not_supported)  [MANUAL\_VERIFICATION\_REQUIRED](/docs/errors/item/#manual_verification_required)  [MFA\_NOT\_SUPPORTED](/docs/errors/item/#mfa_not_supported)  [NO\_ACCOUNTS](/docs/errors/item/#no_accounts)  [NO\_AUTH\_ACCOUNTS](/docs/errors/item/#no_auth_accounts-or-no-depository-accounts) or no-depository-accounts  [NO\_INVESTMENT\_ACCOUNTS](/docs/errors/item/#no_investment_accounts)  [NO\_INVESTMENT\_AUTH\_ACCOUNTS](/docs/errors/item/#no_investment_auth_accounts)  [NO\_LIABILITY\_ACCOUNTS](/docs/errors/item/#no_liability_accounts)  [PASSWORD\_RESET\_REQUIRED](/docs/errors/item/#password_reset_required)  [PRODUCT\_NOT\_ENABLED](/docs/errors/item/#product_not_enabled)  [PRODUCT\_NOT\_READY](/docs/errors/item/#product_not_ready)  [PRODUCTS\_NOT\_SUPPORTED](/docs/errors/item/#products_not_supported)  [USER\_INPUT\_TIMEOUT](/docs/errors/item/#user_input_timeout)  [USER\_SETUP\_REQUIRED](/docs/errors/item/#user_setup_required) |
| Institution errors  Occur when there are errors for the requested financial institution. | [INSTITUTION\_DOWN](/docs/errors/institution/#institution_down)  [INSTITUTION\_NO\_LONGER\_SUPPORTED](/docs/errors/institution/#institution_no_longer_supported)  [INSTITUTION\_NOT\_AVAILABLE](/docs/errors/institution/#institution_not_available)  [INSTITUTION\_NOT\_ENABLED\_IN\_ENVIRONMENT](/docs/errors/institution/#institution_not_enabled_in_environment)  [INSTITUTION\_NOT\_RESPONDING](/docs/errors/institution/#institution_not_responding)  [INSTITUTION\_REGISTRATION\_REQUIRED](/docs/errors/institution/#institution_registration_required)  [UNAUTHORIZED\_INSTITUTION](/docs/errors/institution/#unauthorized_institution)  [UNSUPPORTED\_RESPONSE](/docs/errors/institution/#unsupported_response) |
| API errors  Occur during planned maintenance and in response to API errors. | [INTERNAL\_SERVER\_ERROR](/docs/errors/api/#internal_server_error-or-plaid-internal-error) or plaid-internal-error  [PLANNED\_MAINTENANCE](/docs/errors/api/#planned_maintenance) |
| Assets errors  Occur for errors related to Asset endpoints. | [PRODUCT\_NOT\_ENABLED](/docs/errors/assets/#product_not_enabled)  [DATA\_UNAVAILABLE](/docs/errors/assets/#data_unavailable)  [PRODUCT\_NOT\_READY](/docs/errors/assets/#product_not_ready)  [ASSET\_REPORT\_GENERATION\_FAILED](/docs/errors/assets/#asset_report_generation_failed)  [INVALID\_PARENT](/docs/errors/assets/#invalid_parent)  [INSIGHTS\_NOT\_ENABLED](/docs/errors/assets/#insights_not_enabled)  [INSIGHTS\_PREVIOUSLY\_NOT\_ENABLED](/docs/errors/assets/#insights_previously_not_enabled)  [DATA\_QUALITY\_CHECK\_FAILED](/docs/errors/assets/#data_quality_check_failed) |
| Payment errors  Occur for errors related to Payment Initiation endpoints. | [PAYMENT\_BLOCKED](/docs/errors/payment/#payment_blocked)  [PAYMENT\_CANCELLED](/docs/errors/payment/#payment_cancelled)  [PAYMENT\_INSUFFICIENT\_FUNDS](/docs/errors/payment/#payment_insufficient_funds)  [PAYMENT\_INVALID\_RECIPIENT](/docs/errors/payment/#payment_invalid_recipient)  [PAYMENT\_INVALID\_REFERENCE](/docs/errors/payment/#payment_invalid_reference)  [PAYMENT\_INVALID\_SCHEDULE](/docs/errors/payment/#payment_invalid_schedule)  [PAYMENT\_REJECTED](/docs/errors/payment/#payment_rejected)  [PAYMENT\_SCHEME\_NOT\_SUPPORTED](/docs/errors/payment/#payment_scheme_not_supported)  [PAYMENT\_CONSENT\_INVALID\_CONSTRAINTS](/docs/errors/payment/#payment_consent_invalid_constraints)  [PAYMENT\_CONSENT\_CANCELLED](/docs/errors/payment/#payment_consent_cancelled) |
| Virtual Account errors  Occur for errors related to Virtual Account endpoints. | [TRANSACTION\_INSUFFICIENT\_FUNDS](/docs/errors/virtual-account/#transaction_insufficient_funds)  [TRANSACTION\_AMOUNT\_EXCEEDED](/docs/errors/virtual-account/#transaction_amount_exceeded)  [TRANSACTION\_ON\_SAME\_ACCOUNT](/docs/errors/virtual-account/#transaction_on_same_account)  [TRANSACTION\_CURRENCY\_MISMATCH](/docs/errors/virtual-account/#transaction_currency_mismatch)  [TRANSACTION\_IBAN\_INVALID](/docs/errors/virtual-account/#transaction_iban_invalid)  [TRANSACTION\_BACS\_INVALID](/docs/errors/virtual-account/#transaction_bacs_invalid)  [TRANSACTION\_FAST\_PAY\_DISABLED](/docs/errors/virtual-account/#transaction_fast_pay_disabled)  [TRANSACTION\_EXECUTION\_FAILED](/docs/errors/virtual-account/#transaction_execution_failed)  [NONIDENTICAL\_REQUEST](/docs/errors/virtual-account/#nonidentical_request)  [REQUEST\_CONFLICT](/docs/errors/virtual-account/#request_conflict) |
| Transactions errors  Occur for errors related to Transactions endpoints. | [TRANSACTIONS\_SYNC\_MUTATION\_DURING\_PAGINATION](/docs/errors/transactions/#transactions_sync_mutation_during_pagination) |
| Transfer errors  Occur for errors related to Transfer endpoints. | [TRANSFER\_NETWORK\_LIMIT\_EXCEEDED](/docs/errors/transfer/#transfer_network_limit_exceeded)  [TRANSFER\_ACCOUNT\_BLOCKED](/docs/errors/transfer/#transfer_account_blocked)  [TRANSFER\_NOT\_CANCELLABLE](/docs/errors/transfer/#transfer_not_cancellable)  [TRANSFER\_UNSUPPORTED\_ACCOUNT\_TYPE](/docs/errors/transfer/#transfer_unsupported_account_type)  [TRANSFER\_FORBIDDEN\_ACH\_CLASS](/docs/errors/transfer/#transfer_forbidden_ach_class)  [TRANSFER\_UI\_UNAUTHORIZED](/docs/errors/transfer/#transfer_ui_unauthorized)  [TRANSFER\_ORIGINATOR\_NOT\_FOUND](/docs/errors/transfer/#transfer_originator_not_found)  [INCOMPLETE\_CUSTOMER\_ONBOARDING](/docs/errors/transfer/#incomplete_customer_onboarding)  [UNAUTHORIZED\_ACCESS](/docs/errors/transfer/#unauthorized_access) |
| Signal errors  Occur for errors related to Signal endpoints. | [ADDENDUM\_NOT\_SIGNED](/docs/errors/signal/#addendum_not_signed)  [CLIENT\_TRANSACTION\_ID\_ALREADY\_IN\_USE](/docs/errors/signal/#client_transaction_id_already_in_use)  [INVALID\_CONFIGURATION\_STATE](/docs/errors/signal/#invalid_configuration_state)  [NOT\_ENABLED\_FOR\_SIGNAL\_TRANSACTION\_SCORE\_RULESETS](/docs/errors/signal/#not_enabled_for_signal_transaction_score_rulesets)  [RULESET\_NOT\_FOUND](/docs/errors/signal/#ruleset_not_found)  [SIGNAL\_TRANSACTION\_NOT\_INITIATED](/docs/errors/signal/#signal_transaction_not_initiated) |
| Income errors  Occur for errors related to Income endpoints. | [INCOME\_VERIFICATION\_DOCUMENT\_NOT\_FOUND](/docs/errors/income/#income_verification_document_not_found)  [INCOME\_VERIFICATION\_FAILED](/docs/errors/income/#income_verification_failed)  [INCOME\_VERIFICATION\_NOT\_FOUND](/docs/errors/income/#income_verification_not_found)  [INCOME\_VERIFICATION\_UPLOAD\_ERROR](/docs/errors/income/#income_verification_upload_error)  [PRODUCT\_NOT\_ENABLED](/docs/errors/income/#product_not_enabled)  [PRODUCT\_NOT\_READY](/docs/errors/income/#product_not_ready)  [VERIFICATION\_STATUS\_PENDING\_APPROVAL](/docs/errors/income/#verification_status_pending_approval)  [EMPLOYMENT\_NOT\_FOUND](/docs/errors/income/#employment_not_found) |
| Sandbox errors  Occur when invalid parameters are supplied in the Sandbox environment. | [SANDBOX\_PRODUCT\_NOT\_ENABLED](/docs/errors/sandbox/#sandbox_product_not_enabled)  [SANDBOX\_WEBHOOK\_INVALID](/docs/errors/sandbox/#sandbox_webhook_invalid)  [SANDBOX\_TRANSFER\_EVENT\_TRANSITION\_INVALID](/docs/errors/sandbox/#sandbox_transfer_event_transition_invalid) |
| Invalid Request errors  Occur when a request is malformed and cannot be processed. | [MISSING\_FIELDS](/docs/errors/invalid-request/#missing_fields)  [UNKNOWN\_FIELDS](/docs/errors/invalid-request/#unknown_fields)  [INVALID\_FIELD](/docs/errors/invalid-request/#invalid_field)  [INVALID\_CONFIGURATION](/docs/errors/invalid-request/#invalid_configuration)  [INCOMPATIBLE\_API\_VERSION](/docs/errors/invalid-request/#incompatible_api_version)  [INVALID\_BODY](/docs/errors/invalid-request/#invalid_body)  [INVALID\_HEADERS](/docs/errors/invalid-request/#invalid_headers)  [NOT\_FOUND](/docs/errors/invalid-request/#not_found)  [NO\_LONGER\_AVAILABLE](/docs/errors/invalid-request/#no_longer_available)  [SANDBOX\_ONLY](/docs/errors/invalid-request/#sandbox_only)  [INVALID\_ACCOUNT\_NUMBER](/docs/errors/invalid-request/#invalid_account_number) |
| Invalid Input errors  Occur when all fields are provided, but the values provided are incorrect in some way. | [ADDITIONAL\_CONSENT\_REQUIRED](/docs/errors/invalid-input/#additional_consent_required)  [DIRECT\_INTEGRATION\_NOT\_ENABLED](/docs/errors/invalid-input/#direct_integration_not_enabled)  [INCORRECT\_DEPOSIT\_VERIFICATION](/docs/errors/invalid-input/#incorrect_deposit_verification)  [INVALID\_ACCESS\_TOKEN](/docs/errors/invalid-input/#invalid_access_token)  [INVALID\_ACCOUNT\_ID](/docs/errors/invalid-input/#invalid_account_id)  [INVALID\_API\_KEYS](/docs/errors/invalid-input/#invalid_api_keys)  [INVALID\_AUDIT\_COPY\_TOKEN](/docs/errors/invalid-input/#invalid_audit_copy_token)  [INVALID\_INSTITUTION](/docs/errors/invalid-input/#invalid_institution)  [INVALID\_LINK\_CUSTOMIZATION](/docs/errors/invalid-input/#invalid_link_customization)  [INVALID\_PROCESSOR\_TOKEN](/docs/errors/invalid-input/#invalid_processor_token)  [INVALID\_PRODUCT](/docs/errors/invalid-input/#invalid_product)  [INVALID\_PUBLIC\_TOKEN](/docs/errors/invalid-input/#invalid_public_token)  [INVALID\_LINK\_TOKEN](/docs/errors/invalid-input/#invalid_link_token)  [INVALID\_STRIPE\_ACCOUNT](/docs/errors/invalid-input/#invalid_stripe_account)  [INVALID\_USER\_ID](/docs/errors/invalid-input/#invalid_user_id)  [INVALID\_USER\_IDENTITY\_DATA](/docs/errors/invalid-input/#invalid_user_identity_data)  [INVALID\_USER\_TOKEN](/docs/errors/invalid-input/#invalid_user_token)  [INVALID\_WEBHOOK\_VERIFICATION\_KEY\_ID](/docs/errors/invalid-input/#invalid_webhook_verification_key_id)  [PROFILE\_AUTHENTICATION\_FAILED](/docs/errors/invalid-input/#profile_authentication_failed)  [UNAUTHORIZED\_ENVIRONMENT](/docs/errors/invalid-input/#unauthorized_environment)  [UNAUTHORIZED\_ROUTE\_ACCESS](/docs/errors/invalid-input/#unauthorized_route_access)  [USER\_PERMISSION\_REVOKED](/docs/errors/invalid-input/#user_permission_revoked)  [TOO\_MANY\_VERIFICATION\_ATTEMPTS](/docs/errors/invalid-input/#too_many_verification_attempts) |
| Invalid Result errors  Occur when a request is valid, but the output would be unusable for any supported flow. | [PLAID\_DIRECT\_ITEM\_IMPORT\_RETURNED\_INVALID\_MFA](/docs/errors/invalid-result/#plaid_direct_item_import_returned_invalid_mfa) |
| Rate Limit Exceeded errors  Occur when an excessive number of requests are made in a short period of time. | [ACCOUNTS\_LIMIT](/docs/errors/rate-limit-exceeded/#accounts_limit)  [ACCOUNTS\_BALANCE\_GET\_LIMIT](/docs/errors/rate-limit-exceeded/#accounts_balance_get_limit)  [AUTH\_LIMIT](/docs/errors/rate-limit-exceeded/#auth_limit)  [BALANCE\_LIMIT](/docs/errors/rate-limit-exceeded/#balance_limit)  [CREDITS\_EXHAUSTED](/docs/errors/rate-limit-exceeded/#credits_exhausted)  [IDENTITY\_LIMIT](/docs/errors/rate-limit-exceeded/#identity_limit)  [INSTITUTIONS\_GET\_LIMIT](/docs/errors/rate-limit-exceeded/#institutions_get_limit)  [INSTITUTIONS\_GET\_BY\_ID\_LIMIT](/docs/errors/rate-limit-exceeded/#institutions_get_by_id_limit)  [INSTITUTION\_RATE\_LIMIT](/docs/errors/rate-limit-exceeded/#institution_rate_limit)  [INVESTMENT\_HOLDINGS\_GET\_LIMIT](/docs/errors/rate-limit-exceeded/#investment_holdings_get_limit)  [INVESTMENT\_TRANSACTIONS\_LIMIT](/docs/errors/rate-limit-exceeded/#investment_transactions_limit)  [ITEM\_GET\_LIMIT](/docs/errors/rate-limit-exceeded/#item_get_limit)  [RATE\_LIMIT](/docs/errors/rate-limit-exceeded/#rate_limit)  [TRANSACTIONS\_LIMIT](/docs/errors/rate-limit-exceeded/#transactions_limit)  [TRANSACTIONS\_REFRESH\_LIMIT](/docs/errors/rate-limit-exceeded/#transactions_refresh_limit)  [TRANSACTIONS\_SYNC\_LIMIT](/docs/errors/rate-limit-exceeded/#transactions_sync_limit) |
| ReCAPTCHA errors  Occur when a ReCAPTCHA challenge has been presented or failed during the link process. | [RECAPTCHA\_REQUIRED](/docs/errors/recaptcha/#recaptcha_required)  [RECAPTCHA\_BAD](/docs/errors/recaptcha/#recaptcha_bad) |
| OAuth errors  Occur when there is an error in OAuth authentication. | [INCORRECT\_OAUTH\_NONCE](/docs/errors/oauth/#incorrect_oauth_nonce)  [INCORRECT\_LINK\_TOKEN](/docs/errors/oauth/#incorrect_link_token)  [OAUTH\_STATE\_ID\_ALREADY\_PROCESSED](/docs/errors/oauth/#oauth_state_id_already_processed)  [OAUTH\_STATE\_ID\_NOT\_FOUND](/docs/errors/oauth/#oauth_state_id_not_found) |
| Micro-deposits errors  Occur when there is an error with micro-deposits. | [BANK\_TRANSFER\_ACCOUNT\_BLOCKED](/docs/errors/microdeposits/#bank_transfer_account_blocked) |
| Partner errors  Occur when there is an error with creating or managing end customers. | [CUSTOMER\_NOT\_FOUND](/docs/errors/partner/#customer_not_found)  [FLOWDOWN\_NOT\_COMPLETE](/docs/errors/partner/#flowdown_not_complete)  [QUESTIONNAIRE\_NOT\_COMPLETE](/docs/errors/partner/#questionnaire_not_complete)  [CUSTOMER\_NOT\_READY\_FOR\_ENABLEMENT](/docs/errors/partner/#customer_not_ready_for_enablement)  [CUSTOMER\_ALREADY\_ENABLED](/docs/errors/partner/#customer_already_enabled)  [CUSTOMER\_ALREADY\_CREATED](/docs/errors/partner/#customer_already_created)  [LOGO\_REQUIRED](/docs/errors/partner/#logo_required)  [INVALID\_LOGO](/docs/errors/partner/#invalid_logo)  [CONTACT\_REQUIRED](/docs/errors/partner/#contact_required)  [ASSETS\_UNDER\_MANAGEMENT\_REQUIRED](/docs/errors/partner/#assets_under_management_required)  [CUSTOMER\_REMOVAL\_NOT\_ALLOWED](/docs/errors/partner/#customer_removal_not_allowed)  [OAUTH\_REGISTRATION\_ERROR](/docs/errors/partner/#oauth_registration_error) |
| Check Report errors  Occur when there is an error with creating or retrieving a Check Report. | [CONSUMER\_REPORT\_EXPIRED](/docs/errors/check-report/#consumer_report_expired)  [DATA\_UNAVAILABLE](/docs/errors/check-report/#data_unavailable)  [PRODUCT\_NOT\_READY](/docs/errors/check-report/#product_not_ready)  [INSTITUTION\_TRANSACTION\_HISTORY\_NOT\_SUPPORTED](/docs/errors/check-report/#institution_transaction_history_not_supported)  [INSUFFICIENT\_TRANSACTION\_DATA](/docs/errors/check-report/#insufficient_transaction_data)  [NO\_ACCOUNTS](/docs/errors/check-report/#no_accounts)  [NETWORK\_CONSENT\_REQUIRED](/docs/errors/check-report/#network_consent_required)  [DATA\_QUALITY\_CHECK\_FAILED](/docs/errors/check-report/#data_quality_check_failed) |
| User errors  Occur when there is an error with creating or managing a user | [USER\_NOT\_FOUND](/docs/errors/user/#user_not_found) |
|  |  |

=\*=\*=\*=

#### Error schema

Errors are identified by `error_code` and categorized by `error_type`. Use these in preference to HTTP status codes to identify and handle specific errors. HTTP status codes are set and provide the broadest categorization of errors: 4xx codes are for developer- or user-related errors, and 5xx codes are for Plaid-related errors, and the status will be 2xx in non-error cases. An Item with a non-`null` error object will only be part of an API response when calling [`/item/get`](/docs/api/items/#itemget) to view Item status. Otherwise, error fields will be `null` if no error has occurred; if an error has occurred, an error code will be returned instead.

**Properties**

[`error_type`](/docs/errors/#PlaidError-error-type)

stringstring

A broad categorization of the error. Safe for programmatic use.  
  

Possible values: `INVALID_REQUEST`, `INVALID_RESULT`, `INVALID_INPUT`, `INSTITUTION_ERROR`, `RATE_LIMIT_EXCEEDED`, `API_ERROR`, `ITEM_ERROR`, `ASSET_REPORT_ERROR`, `RECAPTCHA_ERROR`, `OAUTH_ERROR`, `PAYMENT_ERROR`, `BANK_TRANSFER_ERROR`, `INCOME_VERIFICATION_ERROR`, `MICRODEPOSITS_ERROR`, `SANDBOX_ERROR`, `PARTNER_ERROR`, `SIGNAL_ERROR`, `TRANSACTIONS_ERROR`, `TRANSACTION_ERROR`, `TRANSFER_ERROR`, `CHECK_REPORT_ERROR`, `CONSUMER_REPORT_ERROR`, `USER_ERROR`

[`error_code`](/docs/errors/#PlaidError-error-code)

stringstring

The particular error code. Safe for programmatic use.

[`error_code_reason`](/docs/errors/#PlaidError-error-code-reason)

stringstring

The specific reason for the error code. Currently, reasons are only supported OAuth-based item errors; `null` will be returned otherwise. Safe for programmatic use.  
Possible values:
`OAUTH_INVALID_TOKEN`: The user’s OAuth connection to this institution has been invalidated.  
`OAUTH_CONSENT_EXPIRED`: The user's access consent for this OAuth connection to this institution has expired.  
`OAUTH_USER_REVOKED`: The user’s OAuth connection to this institution is invalid because the user revoked their connection.

[`error_message`](/docs/errors/#PlaidError-error-message)

stringstring

A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.

[`display_message`](/docs/errors/#PlaidError-display-message)

stringstring

A user-friendly representation of the error code. `null` if the error is not related to user action.  
This may change over time and is not safe for programmatic use.

[`request_id`](/docs/errors/#PlaidError-request-id)

stringstring

A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.

[`causes`](/docs/errors/#PlaidError-causes)

arrayarray

In this product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  
`causes` will be provided for the `error_type` `ASSET_REPORT_ERROR` or `CHECK_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.

[`status`](/docs/errors/#PlaidError-status)

integerinteger

The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.

[`documentation_url`](/docs/errors/#PlaidError-documentation-url)

stringstring

The URL of a Plaid documentation page with more information about the error

[`suggested_action`](/docs/errors/#PlaidError-suggested-action)

stringstring

Suggested steps for resolving the error

[`required_account_subtypes`](/docs/errors/#PlaidError-required-account-subtypes)

[string][string]

A list of the account subtypes that were requested via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

[`provided_account_subtypes`](/docs/errors/#PlaidError-provided-account-subtypes)

[string][string]

A list of the account subtypes that were extracted but did not match the requested subtypes via the `account_filters` parameter in `/link/token/create`. Currently only populated for `NO_ACCOUNTS` errors from Items with `investments_auth` as an enabled product.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
