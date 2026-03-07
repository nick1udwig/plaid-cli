---
title: "Errors - Invalid Input errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/invalid-input/"
scraped_at: "2026-03-07T22:04:50+00:00"
---

# Invalid Input Errors

#### Guide to troubleshooting invalid input errors

#### **DIRECT\_INTEGRATION\_NOT\_ENABLED**

##### An attempt was made to create an Item without using Link.

##### Common causes

- `/item/create` was called directly, without using Link.

##### Troubleshooting steps

In the Production environment, use Link to create the Item. In Sandbox, use [`/sandbox/public_token/create`](/docs/api/sandbox/#sandboxpublic_tokencreate)

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "DIRECT_INTEGRATION_NOT_ENABLED",
 "error_message": "your client ID is only authorized to use Plaid Link. head to the docs (https://plaid.com/docs) to get started.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INCORRECT\_DEPOSIT\_VERIFICATION**

##### The user submitted an incorrect Manual Same-Day micro-deposit verification input during Item verification in Link.

##### Sample user-facing error message

2 attempts remaining: The code you entered is incorrect. Check your bank statement to find the code in front of ACCTVERIFY.

##### Common causes

- Your user submitted an incorrect micro-deposit verification input when verifying an account via Manual Same-Day micro-deposits.

##### Troubleshooting steps

Have your user attempt to enter the micro-deposit verification input again.

If your user enters an incorrect micro-deposit verification input three times, the Item will be permanently locked. In this case, you must restart the Link flow from the beginning and have the user re-link their account.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INCORRECT_DEPOSIT_VERIFICATION",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_ACCESS\_TOKEN**

##### Common causes

- Access tokens are in the format: `access-<environment>-<identifier>`
- This error can happen when the `access_token` you provided is invalid or pertains to a different API environment

##### Troubleshooting steps

Make sure you are not using a token created in one environment in a different environment (for example, using a Sandbox token in the Production environment).

Ensure that the `client_id`, `secret`, and `access_token` are all associated with the same Plaid developer account.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_ACCESS_TOKEN",
 "error_message": "could not find matching access token",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_ACCOUNT\_ID**

##### The supplied `account_id` is not valid

##### Common causes

One of the `account_id`(s) specified in the API call's `account_ids` object is invalid or does not exist.

- Your integration is passing a correctly formatted, but invalid `account_id` for the Item in question.
- The underlying account may have been closed at the bank, and thus removed from our API.
- The Item affected is at an institution that uses OAuth-based connections, and the user revoked access to the specific account.
- The `account_id` was removed from our API, either completely or a new `account_id` was assigned to the same
  underlying account.
- You are requesting an account that your user has de-selected in the Account Select v2 update flow.

##### Troubleshooting steps

Verify that your integration is passing in correctly formatted and valid `account_id`(s)

Ensure that your integration only uses `account_id`(s) that belong to the Item in question. Early on in your development it is
important to verify that your integration only uses `account_id`(s), and other Plaid identifiers like `item_id`, for the Item that they belong to.

Also be sure to preserve the case of any non-numeric characters in Plaid identifiers, as they are case sensitive.

Verify the Item's currently active accounts and their `account_id`(s).

The user may have revoked access to the account. If this is the case, it is a security best practice to give the user a choice between restoring their account and having your app delete all data for that account. If your user wants to restore access to the account, you can put them through [update mode](/docs/link/update-mode/), which will give them the option to grant access to the account again. Note that doing so will result in an account with a new and different `account_id`, which can be obtained by calling [`/accounts/get`](/docs/api/accounts/#accountsget).

Verify that after completing update mode, your implementation checks for the current `account_id` information associated with the Item, instead of re-using the pre-update mode `account_id`(s). Updated `account_id` data can be found in the `onSuccess` Link event, or by calling certain endpoints, such as [`/accounts/get`](/docs/api/accounts/#accountsget).

Verify that the `account_id` was not changed or removed from the API.

##### Account churn

If the underlying account has not been closed or changed at the bank and the `account_id` no longer appears, Plaid may have
removed the account entirely or assigned the account a new `account_id`, a situation known as "account churn".

Some common causes for account churn are:

- The Item was in an unhealthy state for an extended period of time. If an Item has remained in an error state for over a year, its underlying data may be removed. If the Item is then later refreshed, the Item data will be re-generated, resulting in new `account_id` data.
- The bank or user drastically changing the name of the account, e.g. an account named "Savings account" becomes "Jane's vacation fund".
- The account's mask is changed by the bank, which can occur when banks change their backend systems.

Account churn caused by the latter two reasons is unexpected API behavior. If you experience account churn on an Item that was otherwise healthy, [file a Support ticket](https://dashboard.plaid.com/support/new/financial-institutions/missing-data/missing-accounts).

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_ACCOUNT_ID",
 "error_message": "failed to find requested account ID for requested item",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_API\_KEYS**

##### The client ID and secret included in the request body were invalid. Find your API keys in the Dashboard.

##### Common causes

- The API keys are not valid for the environment being used, which can commonly happen when switching between development environments and forgetting to switch API keys

##### Troubleshooting steps

Find your API keys in the [Dashboard](https://dashboard.plaid.com/developers/keys).

Make sure you are using the secret that corresponds to the environment you are using (Sandbox or Production).

Make sure you are not using a token created in one environment in a different environment (for example, using a Sandbox token in the Production environment).

Visit the Plaid [Dashboard](https://dashboard.plaid.com) to verify that you are enabled for the environment you are using.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_API_KEYS",
 "error_message": "invalid client_id or secret provided",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_AUDIT\_COPY\_TOKEN**

##### The audit copy token supplied to the server was invalid.

##### Common causes

- You attempted to access an Asset Report using an `audit_copy_token` that is invalid or was revoked using [`/asset_report/audit_copy/remove`](/docs/api/products/assets/#asset_reportaudit_copyremove) or [`/asset_report/remove`](/docs/api/products/assets/#asset_reportremove).

##### Troubleshooting steps

Generate a new `audit_copy_token` via [`/asset_report/audit_copy/create`](/docs/api/products/assets/#asset_reportaudit_copycreate).

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_AUDIT_COPY_TOKEN",
 "error_message": null,
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_INSTITUTION**

##### The `institution_id` specified is invalid or does not exist.

##### Common causes

- The `institution_id` specified is invalid or does not exist.

##### Troubleshooting steps

Check the `institution_id` to ensure it is valid.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_INSTITUTION",
 "error_message": "invalid institution_id provided",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_LINK\_CUSTOMIZATION**

##### The Link customization is not valid for the request.

##### Common causes

- The Link customization is missing a use case and the session is enabled for Data Transparency Messaging.
- This error can happen when requesting to update account selections with a Link customization that does not enable [Account Select v2](/docs/link/customization/#account-select).

##### Troubleshooting steps

In the Dashboard, under [Link > Link Customization > Data Transparency Messaging](https://dashboard.plaid.com/link/data-transparency-v5), ensure at least one use case is selected. After selecting a use case, make sure to click "publish changes".

Update your Link customization to enable Account Select v2 or use a Link customization with Account Select v2 enabled.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_LINK_CUSTOMIZATION",
 "error_message": "requested link customization is not set to update account selection",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_LINK\_TOKEN**

##### The `link_token` provided to initialize Link was invalid.

##### Sample user-facing error message

The credentials you provided were incorrect: For security reasons, your account may be locked after several unsuccessful attempts

##### Common causes

- The `link_token` has expired. A `link_token` lasts at least 30 minutes before expiring.
- The `link_token` was already used. A `link_token` can only be used once, except when working in the Sandbox test environment.
- The `link_token` was created in a different environment than the one it was used with. For example, a Sandbox `link_token` was used in Production.
- A user entered invalid credentials too many times during the Link flow, invalidating the `link_token`.

##### Troubleshooting steps

Confirm that the `link_token` is from the correct environment.

Generate a new `link_token` for initializing Link with and re-launch Link.

For more detailed instructions on handling this error, see [Handling invalid Link Tokens](/docs/link/handle-invalid-link-token/).

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_LINK_TOKEN",
 "error_message": "invalid link_token provided",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_PROCESSOR\_TOKEN**

##### The `processor_token` provided to initialize Link was invalid.

##### Common causes

- The `processor_token` used to initialize Link was invalid.

##### Troubleshooting steps

If you are testing in Sandbox, make sure that your `processor_token` was created using the Sandbox-specific endpoint [`/sandbox/processor_token/create`](/docs/api/sandbox/#sandboxprocessor_tokencreate) instead of [`/processor/token/create`](/docs/api/processors/#processortokencreate). Likewise, if testing in Production, make sure that your `processor_token` was created using [`/processor/token/create`](/docs/api/processors/#processortokencreate) rather than [`/sandbox/processor_token/create`](/docs/api/sandbox/#sandboxprocessor_tokencreate).

Make sure you are not using a `processor_token` created in one environment in a different environment (for example, using a Sandbox token in the Production environment).

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_PROCESSOR_TOKEN",
 "error_message": null,
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_PRODUCT**

##### The product is not a valid configuration value, or your client ID does not have access to this product.

##### Common causes

- The [`/link/token/create`](/docs/api/link/#linktokencreate) call specified a product name that does not exist.
- The [`/link/token/create`](/docs/api/link/#linktokencreate) call specified a product that should be omitted from the [`/link/token/create`](/docs/api/link/#linktokencreate) call, such as `balance` or `monitor`.
- The endpoint you are trying to access is not enabled for your account in the environment where you are trying to use it. For example, Identity Verification access is only available in Sandbox after you have received Production access.
- Your integration is using a partner endpoint integration that has not yet been enabled in the Dashboard.
- Your integration is attempting to call a processor endpoint on an Item that was initialized with products that are not compatible with processor endpoints.

##### Troubleshooting steps

If you included `balance`, this product should not be specified during the [`/link/token/create`](/docs/api/link/#linktokencreate) call. Instead, it is initialized automatically when any other product is specified. Include a different product, such as `auth` or `transactions`, and remove `balance` from the API call.

See the [products field documentation](/docs/api/link/#link-token-create-request-products) to ensure the product you specified exists, does not contain a typo, and is valid for [`/link/token/create`](/docs/api/link/#linktokencreate). Examples of product names that should not be passed to [`/link/token/create`](/docs/api/link/#linktokencreate) include `monitor`, `recurring_transactions`, `balance`, and `investments_refresh`.

If you requested a product you don't yet have access to in the specified environment, submit a [product request](https://dashboard.plaid.com/overview/request-products) or [contact sales](/products/protect/#contact-form) to request access. If you are waiting for Production access, you may be able to use Sandbox instead of Production to test in the meantime.

If you are using a partner integration, check the Dashboard [Integrations page](https://dashboard.plaid.com/developers/integrations) to make sure it is enabled.

If you are using a partner integration, when specifying `products` during the [`/link/token/create`](/docs/api/link/#linktokencreate) call, make sure you are not initializing the Item with any Plaid products other than the ones the partner has indicated that it supports.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_PRODUCT",
 "error_message": "client is not authorized to access the following products: [\"identity_verification\"]",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_USER\_ID**

##### The user id is not in a valid format

##### Common causes

- A different identifier (such as a `client_user_id` or `user_token`) was sent instead of a `user_id`.

##### Troubleshooting steps

Confirm the `user_id` format is correct (should start with `usr_`).

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_USER_ID",
 "error_message": "user.user_id must be a non-empty string",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_USER\_IDENTITY\_DATA**

##### The user identity provided failed Plaid’s validation checks. The error message will contain specific details about which field(s) failed validation.

##### Common causes

- Invalid phone number. Plaid phone number validation checks not only for the correct number of digits, but also for a valid area code, prefix, etc., even in the Sandbox environment.
- Required identity fields are empty or contain only whitespace
- Invalid email address format
- Invalid date of birth format (must be YYYY-MM-DD)
- Incomplete address data (missing required fields when any address field is provided)

##### Troubleshooting steps

Review the specific error message in the response for details on which field failed validation.

Verify that the phone number passes [libphonenumber validation](https://libphonenumber.appspot.com/).

API error response

```
http code 403
{
 "error_type": "USER_ERROR",
 "error_code": "INVALID_USER_IDENTITY_DATA",
 "error_message": "Invalid email format.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ADDITIONAL\_CONSENT\_REQUIRED**

##### The end user has not provided consent to the requested product

##### Common causes

- You are using a Link flow that is enabled for [Data Transparency Messaging](/docs/link/data-transparency-messaging-migration-guide/) and are trying to access an endpoint you did not collect consent for.

##### Troubleshooting steps

If your Link flow is enabled for [Data Transparency Messaging](/docs/link/data-transparency-messaging-migration-guide/), check that the `products`, `required_if_supported_products`, `optional_products`, or `additional_consented_products` parameters passed to [`/link/token/create`](/docs/api/link/#linktokencreate) include the product you are requesting.
You can fix any Items with missing consent by using [update mode](/docs/link/update-mode/#data-transparency-messaging).

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "ADDITIONAL_CONSENT_REQUIRED",
 "error_message": "client does not have user consent to access the PRODUCT_AUTH product",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_PUBLIC\_TOKEN**

##### Common causes

- Public tokens are in the format: `public-<environment>-<identifier>`
- This error can happen when the `public_token` you provided is invalid, pertains to a different API environment, or has expired.

##### Troubleshooting steps

Make sure you are not using a token created in one environment in a different environment (for example, using a Sandbox token in the Production environment).

Ensure that the `client_id`, `secret`, and `public_token` are all associated with the same Plaid developer account.

The `public_token` expires after 30 minutes. If your `public_token` has expired, send your user to the Link flow to generate a new `public_token`.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_PUBLIC_TOKEN",
 "error_message": "could not find matching public token",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_STRIPE\_ACCOUNT**

##### The supplied Stripe account is invalid

##### Common causes

After [`/processor/stripe/bank_account_token/create`](/docs/api/processors/#processorstripebank_account_tokencreate) was called, Plaid received a response from Stripe indicating that the Stripe account specified in the API call's `account_id` is invalid.

##### Troubleshooting steps

See the returned `error_message`, which contains information from Stripe regarding why the account was deemed invalid.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_STRIPE_ACCOUNT",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_USER\_TOKEN**

##### The supplied user token is invalid

##### Common causes

- The user token is not associated with the given user ID.
- The user token is invalid or pertains to a different API environment.

##### Troubleshooting steps

Check that the user token is entered correctly.

Make a call to [`/user/create`](/docs/api/users/#usercreate) to create a user token for the given user ID.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_USER_TOKEN",
 "error_message": "could not find matching user token",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_WEBHOOK\_VERIFICATION\_KEY\_ID**

##### The `key_id` provided to the webhook verification endpoint was invalid.

##### Common causes

- A request was made to [`/webhook_verification_key/get`](/docs/api/webhooks/webhook-verification/#get-webhook-verification-key) using an invalid `key_id`.
- The call to [`/webhook_verification_key/get`](/docs/api/webhooks/webhook-verification/#get-webhook-verification-key) was made from an environment different than the one the webhook was sent from (for example, verification of a Sandbox webhook was attempted against Production).

##### Troubleshooting steps

Ensure that the `key_id` argument provided to [`/webhook_verification_key/get`](/docs/api/webhooks/webhook-verification/#get-webhook-verification-key) is in fact the `kid` extracted from the JWT headers. See [webhook verification](/docs/api/webhooks/webhook-verification/) for detailed instructions.

Ensure that the webhook is being verified against the same environment as it originated from.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "INVALID_WEBHOOK_VERIFICATION_KEY_ID",
 "error_message": "invalid key_id provided. note that key_ids are specific to Plaid environments, and verification requests must be made to the same environment that the webhook was sent from",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PROFILE\_AUTHENTICATION\_FAILED**

##### The end user could not authenticate the device associated with their profile

##### Common causes

- When prompted to verify their phone number by entering a code delivered to their number via SMS, the end user entered three incorrect codes. (The first and second incorrect codes will result in an `INVALID_OTP` error, which is recoverable within Link and will not prompt the user to end the Link session.)

##### Troubleshooting steps

Re-route the user to a non-Layer flow.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "PROFILE_AUTHENTICATION_FAILED",
 "error_message": "the profile could not be authenticated",
 "display_message": "We were not able to authenticate your access to this Plaid account.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TOO\_MANY\_VERIFICATION\_ATTEMPTS**

##### The user attempted to verify their Manual Same-Day micro-deposit codes more than 3 times and their Item is now permanently locked. The user must retry submitting their account information in Link.

##### Sample user-facing error message

No attempts remaining: The code you entered is incorrect. You have no more attempts left.

##### Common causes

- Your user repeatedly submitted incorrect micro-deposit codes when verifying an account via Manual Same-Day micro-deposits.

##### Troubleshooting steps

Re-initiate the Link flow and have your user attempt to verify their account again.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "TOO_MANY_VERIFICATION_ATTEMPTS",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **UNAUTHORIZED\_ENVIRONMENT**

##### Your client ID does not have access to this API environment. See which environments you are enabled for from the Dashboard.

##### Sample user-facing error message

Unauthorized Environment: Your Client ID is not authorized to access this API environment. Contact Support to gain access

##### Common causes

- You may not be enabled for the environment you are using.
- Your code may be calling a deprecated endpoint.

##### Troubleshooting steps

Visit the Plaid [Dashboard](https://dashboard.plaid.com) to verify that you are enabled for the environment you are using.

Make sure that your code is not calling deprecated endpoints. Actively supported endpoints are listed in the [API reference](/docs/api/).

Find your API keys in the [Dashboard](https://dashboard.plaid.com/developers/keys).

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "UNAUTHORIZED_ENVIRONMENT",
 "error_message": "you are not authorized to create items in this api environment. Go to the Dashboard (https://dashboard.plaid.com) to see which environments you are authorized for.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **UNAUTHORIZED\_ROUTE\_ACCESS**

##### Your client ID does not have access to this route.

##### Common causes

- The endpoint you are trying to access must be manually enabled for your account.

##### Troubleshooting steps

[Contact Sales](https://plaid.com/contact) to gain access.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "UNAUTHORIZED_ROUTE_ACCESS",
 "error_message": "you are not authorized to access this route in this api environment.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **USER\_PERMISSION\_REVOKED**

##### The end user has revoked access to their data.

##### Common causes

- The end user revoked access to their data via the Plaid consumer portal at my.plaid.com.

##### Troubleshooting steps

Delete the item using [`/item/remove`](/docs/api/items/#itemremove) and prompt your user to re-enter the Link flow to re-authorize access to their data. Note that if the user re-authorizes access, a new Item will be created, and the old Item will not be re-activated.

If applicable, direct your user to a fallback, manual flow for gathering account data.

API error response

```
http code 400
{
 "error_type": "INVALID_INPUT",
 "error_code": "USER_PERMISSION_REVOKED",
 "error_message": "the holder of this account has revoked their permission for your application to access it",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
