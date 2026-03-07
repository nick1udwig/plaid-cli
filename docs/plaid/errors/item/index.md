---
title: "Errors - Item errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/item/"
scraped_at: "2026-03-07T22:04:51+00:00"
---

# Item Errors

#### Guide to troubleshooting Item errors

#### **ACCESS\_NOT\_GRANTED**

##### The user did not grant necessary permissions for their account.

##### Sample user-facing error message

Insufficient Sharing Permissions: There was an error connecting to your account. Try linking your account again by selecting the required information to share with this application.

##### Common causes

- This Item's access is affected by institution-hosted access controls.
- The user did not agree to share, or has revoked, access to the data required for the requested product. Note that for some institutions, the end user may need to specifically opt-in during the OAuth flow to share specific details, such as identity data, or account and routing number information, even if they have already opted in to sharing information about a specific account.
- The user does not have permission to share required information about the account. This can happen at some institutions using OAuth connections if the user is not the account owner (for example, they have a role such as trustee, investment advisor, power of attorney holder, or authorized cardholder).

##### Troubleshooting steps

Prompt the end user to allow Plaid to access identity data and/or account and routing number data. The end user should do this during the Link flow if they were unable to successfully complete the Link flow for the Item, or via Link’s [update mode](/docs/link/update-mode/) if the Item has already been added.

If you frequently experience this error, you may want to ask your Account Manager to enable [Update Mode with Product Validations (UMPV)](/docs/link/update-mode/#resolving-access_not_granted-or-no_auth_accounts-errors-via-product-validations) for your account. Sending the end user through UMPV will ensure that they provide the correct permissions during the OAuth flow.

If your Plaid integration involves adding products to Items after Link (instead of specifying these products in the [`/link/token/create`](/docs/api/link/#linktokencreate) `products` array) consider using [Required if Supported Products](/docs/link/initializing-products/#required-if-supported-products) instead. This will ensure that your user has selected the required permissions for these additional products during Link, reducing the frequency with which `ACCESS_NOT_GRANTED` errors will occur.

If there are other security settings on the user's account that prevent sharing data with aggregators, they should adjust their security preferences on their institution's online banking portal. It may take up to 48 hours for changes to take effect.

Confirm that the user is the owner of the account.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "ACCESS_NOT_GRANTED",
 "error_message": "access to this product was not granted for the account",
 "display_message": "The user did not grant the necessary permissions for this product on their account.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSTANT\_MATCH\_FAILED**

##### Instant Match could not be initialized for the Item.

##### Common causes

- Instant Auth could not be used for the Item, and Instant Match has been enabled for your account, but a country other than `US` is specified in Link's country code initialization.
- The Item does not support Instant Auth or Instant Match. If this is the case, Plaid will automatically attempt to enter a micro-deposit based verification flow.

##### Troubleshooting steps

Update the countries used to initialize Link. Instant Match can only be used when Link is initialized with `US` as the only country code.

Review Link activity logs to verify whether a micro-deposit verification flow was launched in Link after this error occurred. If it was not launched, see [Add institution coverage](https://plaid.com/docs/auth/coverage/) for more information on enabling micro-deposit verification flows. If it was launched successfully, no further troubleshooting action is required.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "INSTANT_MATCH_FAILED",
 "error_message": "Item cannot be verified through Instant Match. Ensure you are correctly enabling all auth features in Link.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INSUFFICIENT\_CREDENTIALS**

##### The user did not provide sufficient authorization in order to link their account via an OAuth login flow.

##### Sample user-facing error message

Couldn't connect to Platypus Bank: Please try to sign in again, or connect another institution.

##### Common causes

- Your user abandoned the bank OAuth flow without completing it.

##### Troubleshooting steps

Have your user attempt to link their account again.

If this error persists, please submit a
[Support](https://dashboard.plaid.com/support/new) ticket with the following
identifiers: `access_token`, `institution_id`, and either `link_session_id` or
`request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "INSUFFICIENT_CREDENTIALS",
 "error_message": "insufficient authorization was provided to complete the request",
 "display_message": "INSUFFICIENT_CREDENTIALS",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_CREDENTIALS**

##### The financial institution indicated that the credentials provided were invalid.

##### Sample user-facing error message

The credentials you provided were incorrect: Check that your credentials are the same that you use for this institution

##### Link user experience

Your user will be redirected to the `Credentials` pane to retry entering correct credentials.

##### Common causes

- The user entered incorrect credentials at their selected institution.
  - Extra spaces, capitalization, and punctuation errors are common causes of `INVALID_CREDENTIALS`.
- The institution requires special configuration steps before the user can link their account with Plaid. KeyBank, Interactive Brokers, Morgan Stanley, and Betterment are examples of institutions that require this setup.
- The user selected the wrong institution.
  - Plaid supports institutions that have multiple login portals for the various products they offer, and it is common for users to confuse a different selection for the one which their credentials would actually be accepted.
  - This confusion is particularly common between Vanguard (brokerage accounts) and My Vanguard Plan (retirement accounts). This is also common for users attempting to link prepaid stored-value cards, as many institutions have separate login portals specifically for those cards.

##### Troubleshooting steps

Prompt your user to retry entering their credentials.

Note: The Institution may lock a user out of their account after 3-5 repeated attempts, resulting in an [`ITEM_LOCKED`](/docs/errors/item/#item_locked) error.

Confirm that the credentials being entered are correct by asking the user to test logging in to their financial institution website using the same set of credentials.

The user should check their financial institution website for special settings to allow access for third-party apps, such as a "third party application password" or "allow third party access" setting.

Verify that the user is selecting the correct institution in Link’s Institution Select view.

If the credentials are confirmed to be legitimate, and a user still cannot authenticate, please submit an 'Invalid credentials errors' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/invalid-creds) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "INVALID_CREDENTIALS",
 "error_message": "the provided credentials were not correct",
 "display_message": "The provided credentials were not correct. Please try again.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_MFA**

##### The institution indicated that the provided response for MFA was invalid

##### Sample user-facing error message

Your answers were incorrect: For security reasons, your account may be locked after several unsuccessful attempts

##### Link user experience

Your user will be redirected to the MFA pane to retry entering the correct value.

##### Common causes

- The user entered an incorrect answer for the security question presented by the selected institution.
- The user selected an MFA device that is not active.
- The institution failed to send the one-time code for the user's selected device.

##### Troubleshooting steps

If the user still cannot log in despite providing correct information, or if they cannot receive an MFA token despite having the correct device, please submit a 'Invalid credentials errors' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/invalid-creds) ticket with the following identifiers: `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "INVALID_MFA",
 "error_message": "the provided MFA response(s) were not correct",
 "display_message": "The provided MFA responses were not correct. Please try again.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_SEND\_METHOD**

##### Returned when the method used to send MFA credentials was deemed invalid by the institution.

##### Link user experience

Your user will be shown an error message indicating that an internal error occurred and will be prompted to close Link.

##### Common causes

- The institution is experiencing login issues.
- The integration between Plaid and the financial institution is experiencing errors.

##### Troubleshooting steps

If the error persists, submit a 'Multi-factor authentication issues' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/mfa-issue) ticket with the following identifiers: `institution_id` and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "INVALID_SEND_METHOD",
 "error_message": "the provided MFA send_method was invalid",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_PHONE\_NUMBER**

##### The submitted phone number was invalid.

##### Sample user-facing error message

Invalid phone number: We couldn't verify that 4151231234 is a valid number. Please re-enter your phone number to try again.

##### Link user experience

Your user will be redirected to the phone number input pane to retry submitting a valid phone number.

##### Common causes

- The user entered an invalid phone number. Only US and CA phone numbers are accepted for the [returning user experience](/docs/link/returning-user/).

##### Troubleshooting steps

If the user still cannot log in despite providing correct information, instruct the user to select "Maybe later" to continue in Link without the returning user experience.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "INVALID_PHONE_NUMBER",
 "error_message": "the provided phone number was invalid",
 "display_message": "The provided phone number was invalid. Please try again.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_OTP**

##### The submitted OTP was invalid.

##### Sample user-facing error message

Security code incorrect: Check that the code you entered is the same code that was sent to you

##### Link user experience

Your user will be redirected to the OTP input pane to retry submitting a valid OTP. After three `INVALID_OTP` errors, Link will stop accepting OTP inputs and the user will be prompted to exit Link or select another institution. If the flow was the Layer or Returning User Experience (Remember Me) flow, after the third invalid OTP code, instead of `INVALID_OTP`, the `PROFILE_AUTHENTICATION_FAILED` error will trigger, and Link will stop accepting OTP inputs.

##### Common causes

- The user entered an invalid OTP.

##### Troubleshooting steps

If the user still cannot log in despite providing correct information, instruct the user to select "Maybe later" to continue in Link without the returning user experience.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "INVALID_OTP",
 "error_message": "the provided OTP was invalid",
 "display_message": "The provided OTP was invalid. Please try again.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INVALID\_UPDATED\_USERNAME**

##### The username entered during update mode did not match the original username.

##### Sample user-facing error message

Username incorrect: Try entering your bank account username again. If you recently changed it, you may need to un-link your account and then re-link.

##### Link user experience

Your user will be directed to enter a different username.

##### Common causes

- While updating an Item in [update mode](/docs/link/update-mode/), the user provided a username that doesn't match the original username provided when they originally linked the account.
- The user was updating an Item in update mode, but Plaid cannot verify that the new Item is the same as the original Item. This may be due to changes to the user's selections during the OAuth UI flow, changes in the `/link/token/create/` call parameters, or internal system changes at the financial institution.

##### Troubleshooting steps

If your user entered the wrong username, or the username for a different account, they should enter the correct, original username.

Have your user ensure that the capitalization for their username is the same as it was when they first logged in to Plaid. The username entered during update mode must case-match with the original username, even if the institution does not consider usernames case-sensitive.

If your user actually changed the username for the account, you should delete the original Item and direct your user to the regular Link flow to link their account as a new Item.

If your user is using the correct and original login credentials, delete the original Item and direct your user to the regular Link flow to link their account as a new Item.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "INVALID_UPDATED_USERNAME",
 "error_message": "the username provided to /item/credentials/update does not match the original username for the item",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ITEM\_CONCURRENTLY\_DELETED**

##### This item was deleted while the operation was in progress.

##### Common causes

- An Item was deleted via [`/item/remove`](/docs/api/items/#itemremove) while a request for its data was in process.

##### Troubleshooting steps

If you plan to delete an Item immediately after retrieving data from it, make sure to wait until your request has successfully returned the data you need before calling [`/item/remove`](/docs/api/items/#itemremove).

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "ITEM_CONCURRENTLY_DELETED",
 "error_message": "This item was deleted while the operation was in progress",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ITEM\_LOCKED**

##### The financial institution indicated that the user's account is locked. The user will need to work with the institution directly to unlock their account.

##### Sample user-facing error message

Too many attempts: We're unable to complete your request. Please log into your account on the financial institution's website to unlock or reset access and then try again.

##### Link user experience

Your user will be directed to a new window with the institution's homepage to unlock their account. Link will then display the Institution Select pane for the user to connect a different account.

##### Common causes

- The user entered their credentials incorrectly after more than 3-5 attempts, triggering the institution’s fraud protection systems and locking the user’s account.

##### Troubleshooting steps

Request that the user log in directly to their institution. Attempting to log in is a reliable way of confirming whether the user’s account is legitimately locked with a given institution. If the user cannot log in due to their account being locked, the website will usually note it as such, giving supplemental information on what they can do to resolve their account.

If the account is locked, ask the user to work with the financial institution to unlock their account.

Steps on unlocking an account are usually provided when a login attempt fails with the institution directly, in the event that their account is actually locked. Once unlocked, they should then try to re-authenticate using Plaid Link.

If the user is persistently locked out of their institution, submit an 'Invalid credentials errors' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/invalid-creds) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "ITEM_LOCKED",
 "error_message": "the account is locked. prompt the user to visit the institution's site and unlock their account",
 "display_message": "The given account has been locked by the financial institution. Please visit your financial institution's website to unlock your account.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ITEM\_LOGIN\_REQUIRED**

##### Additional input from the user is required to continue getting data for this Item.

##### Sample user-facing error message

Username or password incorrect: If you've recently updated your account with this institution, be sure you're entering your updated credentials

##### Common causes

- The institution does not use an OAuth-based connection and the user changed their password.
- The institution does not use an OAuth-based connection and the user changed their multi-factor settings, or their multi-factor authentication has expired.
- The institution has undergone a migration from a non-OAuth-based connection to an OAuth-based connection.
- The institution uses an OAuth-based connection and the user's consent is no longer valid, either because it has expired, or because the user revoked access.
- The institution uses an OAuth-based connection that does not allow duplicate Items, but the user connected a duplicate Item to the same application.
- (Sandbox only) The Item was created in the Sandbox environment and is over 30 days old.

##### Troubleshooting steps

Send the item through Link’s [update mode](/docs/link/update-mode/). Update mode will then automatically prompt the user for whatever input (such as a new password, or refreshed consent) is necessary to fix the Item.

If the Item is on an OAuth-based connection and has an expired `consent_expiration_time`, you may be able to reduce the frequency of this error by listening for the [`PENDING_DISCONNECT`](/docs/api/items/#pending_disconnect) webhook (US/CA Items) or the [`PENDING_EXPIRATION`](/docs/api/items/#pending_expiration) webhook (UK/EU Items) and proactively sending the user through update mode before the Item expires.

If the Item is on an OAuth-based connection and the error was caused by the user adding a duplicate Item, remove the old Item, and see [preventing duplicate Items](/docs/link/duplicate-items/) for more details on preventing and remediating duplicate Items.

To recover from this error more quickly, listen for the [`LOGIN_REPAIRED`](/docs/api/items/#login_repaired) webhook, which will fire when an Item exits `ITEM_LOGIN_REQUIRED` without going through update mode in your app.

If the error is legitimate, having the user authenticate again should generally return their Item to a healthy state without further intervention.

If this error persists, please submit a [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "ITEM_LOGIN_REQUIRED",
 "error_message": "the login details of this item have changed (credentials, MFA, or required user action) and a user login is required to update this information. use Link's update mode to restore the item to a good state",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ITEM\_NOT\_FOUND**

##### The Item you requested cannot be found. This Item does not exist, has been previously removed via /item/remove, or has had access removed by the user

##### Common causes

- Item was previously removed via [`/item/remove`](/docs/api/items/#itemremove).
- The user has depermissioned or deleted their Item via [my.plaid.com](https://my.plaid.com).
- Plaid support has deleted the Item in response to a data deletion request from the user.

##### Troubleshooting steps

Launch a new instance of Plaid Link and prompt the user to create a new Item.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "ITEM_NOT_FOUND",
 "error_message": "The Item you requested cannot be found. This Item does not exist, has been previously removed via /item/remove, or has had access removed by the user.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **ITEM\_NOT\_SUPPORTED**

##### Plaid is unable to support this user's accounts due to restrictions in place at the financial institution.

##### Sample user-facing error message

Account not currently supported: Your account is not currently supported. Please log in using a different account

##### Link user experience

Your user will be redirected to the Institution Select pane to connect a different account.

##### Common causes

- Plaid does not currently support the types of accounts for the connected Item, due to restrictions in place at the selected institution.
- Plaid does not currently support the specific type of multi-factor authentication in place at the selected institution.
- The credentials provided are for a 'guest' account or other account type with limited account access.
- A processor partner token is being created or used, but the country, products, or account types associated with the Item are not supported by the processor partner.

##### Troubleshooting steps

Prompt the user to connect a different account and institution.

If the error persists and none of the common causes above seem to apply, [contact Support](https://dashboard.plaid.com/support/new/).

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "ITEM_NOT_SUPPORTED",
 "error_message": "this account is currently not supported",
 "display_message": "The given account is not currently supported for this financial institution. We apologize for the inconvenience.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **MANUAL\_VERIFICATION\_REQUIRED**

##### Returned when a data request has been made for a product that requires the user to manually verify their Item via Link's update mode.

##### Common causes

- [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) was called on an Item that could not be automatically verified. This can occur when the institution only supports verification through one of the [Investments Move fallback flows](/docs/investments-move/#fallback-flows).

##### Troubleshooting steps

(If using Investments Move) Verify the Item manually by having the user go through a [Link update mode](/docs/link/update-mode/#requesting-additional-products-accounts-permissions-or-use-cases) session using the existing access token and with `investments_auth` initialized in the `products` array.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "MANUAL_VERIFICATION_REQUIRED",
 "error_message": "This Item could not be automatically verified. The user must manually verify their account through a separate Link session using the existing access token and with 'investments_auth' initialized in the 'products' array.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **MFA\_NOT\_SUPPORTED**

##### Returned when the user requires a form of MFA that Plaid does not support for a given financial institution.

##### Sample user-facing error message

Your account settings are incompatible: Your account could not be connected because the multi-factor authentication method it uses is not currently supported. Please try a different account.

##### Link user experience

Your user will be redirected to the Institution Select pane to connect a different account.

##### Common causes

- Plaid does not currently support the specific type of multi-factor authentication in place at the selected institution.
- The user's multi-factor authentication setting is configured not to remember trusted devices and instead to present a multi-factor challenge on every login attempt. This prevents Plaid from refreshing data asynchronously, which many products (especially Transactions) require.

##### Troubleshooting steps

Prompt the user to connect a different account and institution.

If your application does not require asynchronous data refresh to work properly, [contact Support](https://dashboard.plaid.com/support/new/) to explore your options for enabling user access.

If user's settings are not configured to present a multi-factor challenge on every login attempt but this error is still appearing, submit a 'Multi-factor authentication issues' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/mfa-issue) ticket with the following identifiers: `institution_id` and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "MFA_NOT_SUPPORTED",
 "error_message": "this account requires a MFA type that we do not currently support for the institution",
 "display_message": "The multi-factor security features enabled on this account are not currently supported for this financial institution. We apologize for the inconvenience.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NO\_ACCOUNTS**

##### Returned when there are no open accounts associated with the Item.

##### Sample user-facing error message

No compatible accounts: Your credentials are correct, but we couldn’t find any accounts with this institution that are compatible with this application. Try another account, financial institution, or check for another connection method.

##### Link user experience

Your user will be redirected to the Institution Select pane to connect a different account.

##### Common causes

- The user successfully logged into their account, but Plaid was unable to retrieve any open, active, or eligible accounts in the connected Item.
- The user closed their account.
- The user revoked access to the account.
- The account experienced [account churn](/docs/errors/invalid-input/#account-churn), which happens when Plaid can no longer recognize an account as the same one that the user granted you access to. When this happens, your permissions to the account may be revoked.
- There is a problem with Plaid's connection to the user's financial institution.

##### Troubleshooting steps

If you are using the `account_filters` parameter when calling [`/link/token/create`](/docs/api/link/#linktokencreate), ensure that you are not filtering out valid accounts that could be used with your app.

If the Item is at institution that uses OAuth, ensure that the user has not denied or revoked OAuth access to all of their eligible accounts.

Ensure the user has not closed their account. If the user no longer has any open accounts associated with the Item, remove the Item via [`/item/remove`](/docs/api/items/#itemremove).

If you previously had access to the account, send your user through [update mode with account selection enabled](/docs/link/update-mode/#using-update-mode-to-request-new-accounts) to allow them to re-grant access to the account.

Prompt the user to connect a different account and institution.

If this error persists and the user does have accounts at the institution that you should be able to access, please submit a [Support](https://dashboard.plaid.com/support/new) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "NO_ACCOUNTS",
 "error_message": "no valid accounts were found for this item",
 "display_message": "No valid accounts were found at the financial institution. Please visit your financial institution's website to confirm accounts are available.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NO\_AUTH\_ACCOUNTS or no-depository-accounts**

##### Returned from POST /auth/get when there are no valid accounts for which account and routing numbers could be retrieved.

##### Sample user-facing error message

No eligible accounts: We didn't find any accounts eligible for money movement at this institution. Please try linking another institution

##### Common causes

- [`/auth/get`](/docs/api/products/auth/#authget) was called on an Item with no accounts that can support Auth, or accounts that do support Auth were filtered out. Only debitable checking, savings, and cash management accounts can be used with Auth.
- Plaid's ability to retrieve Auth data from the institution has been temporarily disrupted.
- The end user is connecting to an institution that uses OAuth-based flows but did not grant permission in the OAuth flow for the institution to share details for any compatible accounts. Note that for some institutions, the end user may need to specifically opt-in during the OAuth flow to share account and routing number information even if they have already opted in to sharing information about their checking or savings account.
- The end user revoked access to the account.
- The account experienced [account churn](/docs/errors/invalid-input/#account-churn), which happens when Plaid can no longer recognize an account as the same one that the user granted you access to. When this happens, your permissions to the account may be revoked.

##### Troubleshooting steps

Ensure that any `account_id` specified in the `options` filter for [`/auth/get`](/docs/api/products/auth/#authget) belongs to a debitable checking, savings, or cash management account.

Ensure that the end user has a debitable checking, savings, or cash management account at the institution. Not all accounts permit ACH debits. Common examples of non-debitable accounts include savings accounts at Chime or at Navy Federal Credit Union (NFCU).

If the error occurred after attempting to add Auth to a previously healthy Item, send the user through [Update Mode with Product Validations](/docs/link/update-mode/#resolving-access_not_granted-or-no_auth_accounts-errors-via-product-validations).

If the Item is at an OAuth-based institution, prompt the end user to allow Plaid to access a debitable checking, savings, or cash management account, along with its account and routing number data. The end user should do this during the Link flow if they were unable to successfully complete the Link flow for the Item, or at their institution's online banking portal if the Item has already been added.

If you previously had access to the account, send your user through the [update mode flow](/docs/link/update-mode/) to allow them to re-grant access to the account.

Return your user to the Link flow and prompt them to link a different account.

If the problem persists even though the end user has confirmed that they have a debitable checking, savings, or cash management account at the institution, [open a support ticket](https://dashboard.plaid.com/support/new).

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "NO_AUTH_ACCOUNTS",
 "error_message": "There are no valid checking or savings account(s) associated with this Item. See https://plaid.com/docs/api/#item-errors for more.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NO\_INVESTMENT\_ACCOUNTS**

##### Returned from POST /investments/holdings/get, POST /investments/transactions/get, or Link initialized with the Investments product, when there are no valid investment account(s) for which holdings or transactions could be retrieved.

##### Sample user-facing error message

No investment accounts: None of your accounts are investment accounts. Please connect using a different bank

##### Common causes

- Link was initialized with the Investments product, but the user attempted to link an account with no investment accounts.
- The [`/investments/holdings/get`](/docs/api/products/investments/#investmentsholdingsget) or [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) endpoint was called, but there are no investment accounts associated with the Item.
- The end user is connecting to an institution that uses OAuth-based flows but did not grant permission in the OAuth flow for the institution to share details for any investment accounts.

##### Troubleshooting steps

Have the user open an investment account at the institution and then re-link, or link an Item that already has an investment account.

If your end user is connecting to an institution that uses OAuth-based flows (one for which the `oauth` field in the institution record is `true`), ensure that your end user consented to share details for an investment account.

If the problem persists, Plaid may be erroneously categorizing the account. If this is the case, [open a support ticket](https://dashboard.plaid.com/support/new).

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "NO_INVESTMENT_ACCOUNTS",
 "error_message": "There are no valid investment account(s) associated with this Item. See https://plaid.com/docs/api/#item-errors for more information.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NO\_INVESTMENT\_AUTH\_ACCOUNTS**

##### Returned from POST /investments/holdings/get or POST /investments/transactions/get when there are no valid investment account(s) for which holdings or transactions could be retrieved.

##### Common causes

- The [`/investments/holdings/get`](/docs/api/products/investments/#investmentsholdingsget) or [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) endpoint was called, but there are no investment accounts associated with the Item.
- The end user is connecting to an institution that uses OAuth-based flows but did not grant permission in the OAuth flow for the institution to share details for any investment accounts.

##### Troubleshooting steps

Have the user open an investment account at the institution and then re-link, or link an Item that already has an investment account. If the user links a new Item, delete the old one via [`/item/remove`](/docs/api/items/#itemremove).

If your end user is connecting to an institution that uses OAuth-based flows (one for which the `oauth` field in the institution record is `true`), ensure that your end user consented to share details for an investment account.

If the problem persists, Plaid may be erroneously categorizing the account. If this is the case, [open a support ticket](https://dashboard.plaid.com/support/new).

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "NO_INVESTMENT_AUTH_ACCOUNTS",
 "error_message": "There are no valid investment account(s) associated with this Item. See https://plaid.com/docs/api/#item-errors for more information.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **NO\_LIABILITY\_ACCOUNTS**

##### Returned from POST /liabilities/get when there are no valid liability account(s) for which liabilities could be retrieved.

##### Sample user-facing error message

No liability accounts: None of your accounts are liability accounts. Please connect using a different bank

##### Common causes

- The [`/liabilities/get`](/docs/api/products/liabilities/#liabilitiesget) endpoint was called, but there are no supported liability accounts associated with the Item.
- The end user is connecting to an institution that uses OAuth-based flows but did not grant permission in the OAuth flow for the institution to share details for any liabilities accounts.

##### Troubleshooting steps

Make sure the user has linked an Item with a supported account type and subtype. The account types supported for the Liabilities product are `credit` accounts with the subtype of `credit card` or `paypal`, and `loan` accounts with the subtype of `student loan` or `mortgage`.

If your end user is connecting to an institution that uses OAuth-based flows (one for which the `oauth` field in the institution record is `true`), ensure that your end user consented to share details for a credit card, student loan, PayPal credit account, or mortgage.

If the problem persists, Plaid may be erroneously categorizing the account. If this is the case, [open a support ticket](https://dashboard.plaid.com/support/new).

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "NO_LIABILITY_ACCOUNTS",
 "error_message": "There are no valid liability account(s) associated with this Item. See https://plaid.com/docs/api/#item-errors for more information.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PASSWORD\_RESET\_REQUIRED**

##### The user must log in directly to the financial institution and reset their password.

##### Common causes

- The institution is blocking access because the user needs to reset their password.

##### Troubleshooting steps

Request that the user log in and reset their password directly at their institution.

Ask that the user log in to the bank, and confirm whether they are presented with a message to reset their password.
If they are, they should follow the instructions to reset their password.
Once the password is reset, they should attempt to link their account again.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "PASSWORD_RESET_REQUIRED",
 "error_message": "user must reset their password",
 "display_message": "The user needs to reset their password for their account",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PRODUCT\_NOT\_ENABLED**

##### A requested product was not enabled for the current access token. Please ensure it is included when when initializing Link and create the Item again.

##### Common causes

- You requested a product that was not enabled for the current access token. Ensure it is included when when calling [`/link/token/create`](/docs/api/link/#linktokencreate) and create the Item again.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "PRODUCT_NOT_ENABLED",
 "error_message": "A requested product was not enabled for the current access token. Please ensure it is included when when initializing Link and create the Item again.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PRODUCT\_NOT\_READY**

##### Returned when a data request has been made for a product that is not yet ready.

For the Assets version of this error, see [`ASSET_REPORT_ERROR: PRODUCT_NOT_READY`](/docs/errors/assets/#product_not_ready). For the Income Verification version, see [`INCOME_VERIFICATION_ERROR: PRODUCT_NOT_READY`](/docs/errors/income/#product_not_ready).

##### Common causes

- [`/transactions/get`](/docs/api/products/transactions/#transactionsget) was called before the first 30 days of transactions data could be extracted. This typically happens if the endpoint was called within a few seconds of linking the Item. It will also happen if [`/transactions/get`](/docs/api/products/transactions/#transactionsget) was called for the first time on an Item that was not initialized with `transactions` in the [`/link/token/create`](/docs/api/link/#linktokencreate) call. Note that this error does not occur when using [`/transactions/sync`](/docs/api/products/transactions/#transactionssync); if called too early, [`/transactions/sync`](/docs/api/products/transactions/#transactionssync) will simply fail to return any transactions.
- Occasionally, this error can occur upon calling [`/transactions/get`](/docs/api/products/transactions/#transactionsget) if Plaid's attempt to extract transactions for the Item has failed, due to a connectivity error with the financial institution, and Plaid has never successfully extracted transactions for the Item in the past. If this happens, Plaid will continue to retry the extraction at least once a day.
- [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) was called for the first time on a new Item before any balance data could be extracted for the Item. When called, [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will return a `PRODUCT_NOT_READY` error if cached balance data does not exist and new balance data could not be obtained after 15 seconds. For some customers who participated in the early phases of the Signal beta, [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) will return `PRODUCT_NOT_READY` immediately if cached balance data is not available when [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) is called.
- [`/auth/get`](/docs/api/products/auth/#authget) was called on an Item that hasn't been verified, which is possible when using [micro-deposit based verification](/docs/auth/coverage/).
- [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) was called before the investments data could be extracted. This typically happens when the endpoint is called with the option for `async_update` set to true, and called again within a few seconds of linking the Item. It will also happen if [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) (with `async_update` set to true) was called for the first time on an Item that was not initialized with `investments` in the [`/link/token/create`](/docs/api/link/#linktokencreate) call.

##### Troubleshooting steps

If you know at the point of Link initialization that you will want to use Transactions with the linked Item, initialize Link with `transactions` in order to start the product initialization process ahead of time. For more details, see [Choosing how to initialize products](/docs/link/initializing-products/).

Listen for the `INITIAL_UPDATE` webhook fired when the Transactions product is ready and only call [`/transactions/get`](/docs/api/products/transactions/#transactionsget) after that webhook has been fired.

(If using Auth) Verify the Item manually, or wait for automated verification to succeed.

(For failures to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate)) Wait a few seconds, then retry.

Listen for the `HISTORICAL_UPDATE` webhook fired when the Investments product is ready and only call [`/investments/transactions/get`](/docs/api/products/investments/#investmentstransactionsget) after that webhook has been fired.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "PRODUCT_NOT_READY",
 "error_message": "the requested product is not yet ready. please provide a webhook or try the request again later",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PRODUCTS\_NOT\_SUPPORTED**

##### Returned when a data request has been made for an Item for a product that it does not support. Use the /item/get endpoint to find out which products an Item supports.

##### Common causes

- A product endpoint request was made for an Item that does not support that product. This can happen when trying to call a product endpoint on an Item if the product was not included in either the `products` or `optional_products` arrays when calling [`/link/token/create`](/docs/api/link/#linktokencreate).
- Updated accounts have been requested for an Item initialized with the Assets product, which does not support adding or updating accounts after the initial Link.
- You are attempting to call [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) on a non-US financial institution.
- You are attempting to call [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) on a manually linked item (Database Insights, Database Match, Instant Micro-deposits, Same Day Micro-deposits).
- You are attempting to call [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) on a Capital One Item that contains only credit cards or other non-depository products. Capital One does not support the [`/transactions/refresh`](/docs/api/products/transactions/#transactionsrefresh) endpoint for non-depository products.
- You are attempting to call [`/identity/match`](/docs/api/products/identity/#identitymatch) on an Item that was manually linked (using Database Insights, Database Match, Instant Micro-deposits, or Same Day Micro-deposits) that has not been seen on the Plaid network before.

##### Troubleshooting steps

Use the [`/item/get`](/docs/api/items/#itemget) endpoint to determine which products a given Item supports.

Use the [`/institutions/get_by_id`](/docs/api/institutions/#institutionsget_by_id) endpoint to determine which products a given institution supports. Receiving this error is expected if the user's institution does not support the requested product.

If the product is critical to your use case, put it the `products` array when calling [`/link/token/create`](/docs/api/link/#linktokencreate). Warning: making this change may reduce conversion, as it will block customers from linking their accounts if their institution doesn't support the product. For more details, see [Choosing how to initialize products](/docs/link/initializing-products/).

If the Item is at an OAuth-based institution, prompt the end user to allow Plaid to access the specific types of data your app needs. The end user should do this during the Link flow if they were unable to successfully complete the Link flow for the Item, or If the Item is at an OAuth-based institution, prompt the end user to allow Plaid to access identity data and/or account and routing number data. The end user should do this during the Link flow if they were unable to successfully complete the Link flow for the Item, or via Link's [update mode](/docs/link/update-mode/) if the Item has already been added.

If receiving this error during [`/link/token/create`](/docs/api/link/#linktokencreate) for update mode, and you are using the Item with the Assets product, you must instead create a new Item rather than updating the existing Item.

Call [`/institutions/get`](/docs/api/institutions/#institutionsget) and ensure the `country_code` is US before calling [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate).

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "PRODUCTS_NOT_SUPPORTED",
 "error_message": "",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **USER\_SETUP\_REQUIRED**

##### The user must log in directly to the financial institution and take some action before Plaid can access their accounts.

##### Sample user-facing error message

Action required with your account: Log in to your bank and update your account. Then, return to Plaid to continue.

##### Link user experience

Your user will be directed to a new window with the institution's homepage to unlock their account. Link will then display the Institution Select pane for the user to connect a different account.

##### Common causes

- The institution requires special configuration steps before the user can link their account with Plaid. KeyBank, Morgan Stanley, Interactive Brokers, and Betterment are examples of institutions that require this setup.
- The user’s account is not fully set up at their institution.
- The institution is blocking access due to an administrative task that requires completion. This error can arise for a number of reasons, the most common being:
  - The user must agree to updated terms and conditions.
  - The user must reset their password.
  - The user must enter additional account information.

##### Troubleshooting steps

Request that the user log in and complete their account setup directly at their institution.

Ask that the user log in to the bank, and confirm whether they are presented with some form of agreement page, modal, or some other actionable task. The user should also check their bank website for special settings to allow access for third-party apps, such as a "third party application password" or "allow third party access" setting.

This can usually be done completely online, but there are some financial institutions that require the user to call a phone number, or access a physical branch.

Once completed, prompt the user to re-authenticate with Plaid Link.

If the user is still unable to log in to their institution, please submit a 'Invalid credentials errors' [Support](https://dashboard.plaid.com/support/new/financial-institutions/authentication-issues/invalid-creds) ticket with the following identifiers: `access_token`, `institution_id`, and either `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "USER_SETUP_REQUIRED",
 "error_message": "the account has not been fully set up. prompt the user to visit the issuing institution's site and finish the setup process",
 "display_message": "The given account is not fully setup. Please visit your financial institution's website to setup your account.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **USER\_INPUT\_TIMEOUT**

##### The user did not complete a step in the Link flow, and it timed out.

##### Sample user-facing error message

Session expired: Please re-enter your information to link your accounts.

##### Common causes

- Your user did not complete the account selection flow within five minutes.

##### Troubleshooting steps

Have your user attempt to link their account again.

API error response

```
http code 400
{
 "error_type": "ITEM_ERROR",
 "error_code": "USER_INPUT_TIMEOUT",
 "error_message": "user must retry this operation",
 "display_message": "The application timed out waiting for user input",
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
