---
title: "Link - Preventing duplicate Items | Plaid Docs"
source_url: "https://plaid.com/docs/link/duplicate-items/"
scraped_at: "2026-03-07T22:05:04+00:00"
---

# Preventing duplicate Items

#### Prevent unnecessary billing and confusing application behavior

#### How duplicate Items are created

An [Item](https://plaid.com/docs/quickstart/glossary/#item) represents a login at a financial institution. This login typically happens when an end user links an account (or several accounts) using [Plaid Link](https://plaid.com/docs/link/). A duplicate Item will be created if the end user logs into the same institution using the same credentials again using Plaid Link (in the same application), and if an access token is requested for the Item.

Duplicate Items can occur for multiple reasons. For example, a duplicate Item can occur if a user accidentally links the same account more than once, because they do not realize they already linked an account, or because their linked account is no longer working. Duplicate Items can also occur if a user intentionally links multiple Items for abusive purposes (for example, as part of an attempt to receive multiple sign-up bonuses or to evade a ban).

Duplicate Items can create confusing or unwanted behavior in your application, can sometimes be a vector for fraud and abuse, and may cause you to be charged multiple times for the same Item. We recommend building safeguards in your application to help prevent end users from creating duplicate Items. In this article, we'll describe a few ways to prevent and detect Item duplication.

#### Preventing duplicate Items

##### Use the onSuccess callback metadata

The [`onSuccess`](https://plaid.com/docs/link/web/#onsuccess) callback is called when a user successfully links an Item using Plaid Link. This callback provides metadata that you can use to prevent duplicate Items from being created. (Alternatively, if you're using a connection method that does not provide data in the `onSuccess` callback, such as Hosted Link, you can obtain this metadata by calling [`/link/token/get`](/docs/api/link/#linktokenget) and using the `results.item_add_results` array.) One approach is to require a user login prior to launching Plaid Link so that you can retrieve existing Items associated with the user.

Then, before requesting an `access_token`, examine and compare the `onSuccess` callback metadata to the user's existing Items. You can compare a combination of the accounts’ `institution_id`, account `name`, and account `mask` to determine whether an end user has previously linked an account to your application. Do not exchange a public token for an access token if you detect a duplicate Item.

While the `mask` value is usually the same as the last 4 digits of the account number, this is not the case at all institutions. Never detect duplicate Items by attempting to match a `mask` with an account number.

Sample onSuccess metadata schema

```
{
  institution: {
    name: 'Wells Fargo',
    institution_id: 'ins_4'
  },
  accounts: [
    {
      id: 'ygPnJweommTWNr9doD6ZfGR6GGVQy7fyREmWy',
      name: 'Plaid Checking',
      mask: '0000',
      type: 'depository',
      subtype: 'checking',
      verification_status: ''
    },
    {
      id: '9ebEyJAl33FRrZNLBG8ECxD9xxpwWnuRNZ1V4',
      name: 'Plaid Saving',
      mask: '1111',
      type: 'depository',
      subtype: 'savings'
    }
    ...
  ],
  link_session_id: '79e772be-547d-4c9c-8b76-4ac4ed4c441a'
}
```

Note that duplicate Items are not always identical. For example, an Item may have a checking account linked, while its duplicate may have only a savings account linked. While this scenario is rare, if it becomes a problem for your use case, you can switch to using the `institution_id` field on a per-user basis to prevent duplicate Items and enforce that each user of your app not link multiple Items with the same institution. However, because this is a legitimate use case for some applications, it is not recommended as the primary means of detecting duplicate Items.

Duplicate Items at Chase, PNC, NFCU, and Schwab

For Chase, PNC, Navy Federal Credit Union, and Charles Schwab, existing Items may be invalidated if an end user adds a second Item using the same credentials, at the point when the user has completed the institution's OAuth flow. At Chase, this will only happen if the second Item does not have exactly the same set of accounts associated as the first Item (i.e. the user granted permissions to different accounts in the institution's OAuth flow) or if either of the Items is initialized with a [Plaid Check product](/docs/check/).

If this occurs, delete the old Item using [`/item/remove`](/docs/api/items/#itemremove) and use the new Item. Alternatively, the old Item can repaired using [update mode](/docs/link/update-mode/), which will invalidate the new Item. Then you may delete the new Item.

For Auth or Transfer customers, Chase or PNC Item invalidation caused by creating a duplicate Item will not invalidate or change the Item's TAN.

##### Implement Pre-Link messaging

A lightweight but effective method for preventing accidental duplicate Items from being created is by providing relevant messaging and information to end users before they engage with Plaid Link in your application. For example, displaying a list of accounts they've already connected to your application can help prevent end users from inadvertently linking the same accounts again.

##### Use Link's update mode

From time to time, an Item may become unhealthy due to entering the `ITEM_LOGIN_REQUIRED` state. Or, the Item may still be healthy, but you may want your end user to authorize additional accounts or permissions associated with the Item. When either of these scenarios happens, use [Link in update mode](/docs/link/update-mode/) to refresh the Item instead of creating a new Item.

#### Example implementation: Preventing accidental duplicate Items

For an example that demonstrates how to prevent accidental duplicate Items, see the [Plaid Pattern](https://github.com/plaid/pattern) sample app. Plaid Pattern implements simple server-side logic to check whether the user ID and institution ID pair already exist in the application database. If the pair exists, an access token will not be requested for this Item, thereby preventing a duplicate. In addition, a message is displayed to the end user informing them that they've already linked this Item. The relevant code can be found in [/server/routes/items.js](https://github.com/plaid/pattern/blob/master/server/routes/items.js#L41-L49) and [/server/db/queries/items.js](https://github.com/plaid/pattern/blob/master/server/db/queries/items.js#L73-L88).

#### Identifying existing duplicate Items

To identify existing duplicate Items, use the same matching logic as described above, but retrieve this data via the [`/accounts/get`](/docs/api/accounts/#accountsget) endpoint instead of the `onSuccess` callback. Occasionally, the `mask` or `name` fields may be null, in which case you can compare `institution_id` and `client_user_id` as a fallback. After identifying duplicate Items, use the [`/item/remove`](/docs/api/items/#itemremove) endpoint to delete an Item.

If you are using Auth via [`/auth/get`](/docs/api/products/auth/#authget), existing duplicate Items may be detected by using the account and routing number fields and looking for duplicates. This method will not work for Items at institutions that use [Tokenized Account Numbers (TANs)](/docs/auth/#tokenized-account-numbers). To detect existing duplicate Items at these institutions, you can use the [`persistent_account_id`](/docs/api/accounts/#accounts-get-response-accounts-persistent-account-id) field, which will be the same across duplicate instances of the Item.

#### Detecting and preventing duplicate Items across user accounts

When attempting to detect or prevent duplicate Items, your approach should depend on the type of duplicate Item you are trying to prevent. For example, you can often scope your search to either a single user's accounts, or to all accounts known to your application. If you are attempting to prevent accidental duplicate Items, you should scope the search to a single user; if you are attempting to detect or prevent abuse, it may make more sense to expand your search to all accounts that have been linked to your app.

Note that there can be legitimate use cases for the same financial institution account to be linked across multiple different user accounts (for example, family members who share a joint bank account, or employees of the same company who share access to a business account); depending on your use case, you may want to incorporate this data into a broader abuse detection framework rather than blocking all duplicate accounts.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
