---
title: "Identity - Introduction to Identity | Plaid Docs"
source_url: "https://plaid.com/docs/identity/"
scraped_at: "2026-03-07T22:04:56+00:00"
---

# Introduction to Identity

#### Verify users' account ownership and reduce fraud with the Identity product.

Get started with Identity

[API Reference](/docs/api/products/identity/)[Quickstart](/docs/quickstart/)

#### Overview

Plaid's Identity product helps you verify users' identities by accessing information on file with their financial institution. Using Identity, you can access a user's phone number, email address, mailing address, and name. This can be used to fight fraud by verifying that the linked financial institution account is owned by the same user who created an account with you.

Prefer to learn by watching? Get an overview of how Identity works in just 3 minutes!

Plaid provides two endpoints for Identity:

- [`/identity/get`](/docs/api/products/identity/#identityget) retrieves the user's name and contact information from their financial institution
- [`/identity/match`](/docs/api/products/identity/#identitymatch) allows you to provide a name and contact information for a user and returns a set of match scores indicating how well that information matches the information on file with their financial institution, simplifying the process of using identity data for account ownership verification.

Both endpoints can be used to reduce fraud, improve user onboarding and conversion, and to complement Know Your Customer (KYC) compliance checks.

Identity can be used to verify all major account types supported by Plaid, including checking and savings accounts, credit cards, brokerage accounts, and loans. Identity coverage can be used to verify an account before initiating an ACH transaction: 97% of Items initialized with [Auth](/docs/auth/) provide Identity data as well.

Depending on your use case, you may want to verify the identity of all users, or only some. For example, you might want to verify the identity of any user initiating a funds transfer, or you might only verify users who you have identified as being higher risk, based on data such as email address, location, financial institution, or activity patterns.

Plaid Identity can be combined with [Plaid Identity Verification](/docs/identity-verification/) in a single workflow to provide full identity verification and fraud reduction. [Identity Verification](/docs/identity-verification/) is used for KYC, to verify that the user of your app is the person they claim to be, while Identity is used to confirm that the ownership information on their linked bank or credit card account matches this verified identity.

#### Integration overview

1. Create a Link token by calling [`/link/token/create`](/docs/api/link/#linktokencreate).
   - If you are using Identity as a best-effort adjunct to other products, like Auth, put `identity` in the `required_if_supported_products` array.
   - If you want to require Identity and fail sessions where Identity data is not available, put `identity` in the `products` array.
2. Initialize Link with the Link token from the previous step. For more details, see [Link](/docs/link/).
3. Once the user has successfully completed the Link flow, exchange the `public_token` for an `access_token` by calling [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange).
4. Once you have the `access_token`, call [`/identity/get`](/docs/api/products/identity/#identityget) to get data, or [`/identity/match`](/docs/api/products/identity/#identitymatch) to match data. For more information on [`/identity/match`](/docs/api/products/identity/#identitymatch), including optional streamlined integration modes, see [Identity Match](/docs/identity/#identity-match).

#### Getting Identity data

The [`/identity/get`](/docs/api/products/identity/#identityget) endpoint provides you with several pieces of information. The name is guaranteed to be available; the email address, phone number, and address are usually available, but may be `null` otherwise.

##### Sample Identity data

Identity data returned by [`/identity/get`](/docs/api/products/identity/#identityget) includes owners' names, addresses, email addresses, and phone numbers as reported by the financial institution. In the case of an Item containing accounts with multiple owners, all owners' information will be provided. For business accounts, the account owner’s name can be that of a business.

Sample Identity data

```
"owners": [
  {
    "addresses": [
      {
        "data": {
          "city": "Malakoff",
          "country": "US",
          "postal_code": "14236",
          "region": "NY",
          "street": "2992 Cameron Road"
        },
        "primary": true
      }
    ],
    "emails": [
      {
        "data": "accountholder0@example.com",
        "primary": true,
        "type": "primary"
      },
      {
        "data": "accountholdersecondary0@example.com",
        "primary": false,
        "type": "secondary"
      }
    ],
    "names": [
      "Alberta Bobbeth Charleson"
    ],
    "phone_numbers": [
      {
        "data": "1112223333",
        "primary": true,
        "type": "home"
      }
    ]
  }
]
```

##### Typical identity fill rates by field

| Field | Typical fill rate |
| --- | --- |
| Name | 100% |
| Email address | 98% |
| Address | 93% |
| Phone number | 90% |

#### Identity Match

Using [`/identity/match`](/docs/api/products/identity/#identitymatch), Identity data fields such as names, addresses, email addresses, and phone numbers can be matched against the owners’ identity information on connected accounts. Multiple fields are matched in a single request, and a separate score is returned for each field to indicate how closely it matched. Customers using Identity Match can experience onboarding conversion improvements of 20% or more, without increasing fraud rates, when using Identity Match versus using their own matching algorithms.

If you are using [Auth](/docs/auth/), you can optionally integrate Identity Match via the Dashboard with no additional code. For details, see [Integrating Identity Match via the Account Verification Dashboard](/docs/identity/#integrating-identity-match-via-the-account-verification-dashboard).

![](/assets/img/docs/identity-match.png)

Example Identity Match results in the Account Verification Dashboard

If you are already using [Identity Verification](/docs/identity-verification/), you can further enhance your verification process by enabling Financial Account Matching in your Identity Verification template.
This feature allows you to match the data collected during KYC without having to specify the `user` field in your [`/identity/match`](/docs/api/products/identity/#identitymatch) request.
To ensure accurate matching, it's important to maintain consistency in the `client_user_id` used for end users across all products.

##### Sample Identity Match data

Data returned by [`/identity/match`](/docs/api/products/identity/#identitymatch) includes scores from matching a user's name, address, phone number, and email with the account owner's data that was present on the connected account.
A score ranging from 0 to 100 is provided for each identity field. The score for a given field will be missing if the identity information for that field was not provided in the request or unavailable at the connected account. In the case of an Item containing accounts with multiple owners, the highest matched scores are returned.

You should typically not set the match score requirement for a field to 100. For example, if a phone number match score of 100 is required, the presence or absence of a country code, parentheses, or other formatting differences may cause a phone number mismatch. 70 is the default recommended match score threshold for all fields.

###### Example of how to interpret name match score

| Range | Meaning | Example |
| --- | --- | --- |
| 100 | Exact match | Andrew Smith, Andrew Smith |
| 85-99 | Strong match, likely spelling error, nickname, or a missing middle name, prefix or suffix | Andrew Smith, Andrew Simth |
| 70-84 | Possible match, likely alias or nickname and spelling error | Andrew Smith, Andy Simth |
| 50-69 | Unlikely match, likely relative | Andrew Smith, Betty Smith |
| 0-49 | Unlikely match | Andrew Smith, Ray Charles |

###### Example of how to interpret phone number score

| Range | Meaning | Example |
| --- | --- | --- |
| 100 | Exact match | +1-555-867-5309, +1-555-867-5309 |
| 90-99 | Same phone number, likely different formatting | +1-555-867-5309, 1 (555)-867-5309 |
| 70-89 | Same phone number, likely different formatting and/or missing country code | +1-555-867-5309, 5558675309 |
| 0-69 | Unlikely match | +1-555-867-5309, 555-867-5302 |

Sample Identity match data

```
{
  "accounts": [
    {
      ..
      "legal_name": {
        "score": 90,
        "is_nickname_match": true,
        "is_first_name_or_last_name_match": true
      },
      "phone_number": {
        "score": 100
      },
      "email_address": {
        "score": 100
      },
      "address": {
        "score": 100,
        "is_postal_code_match": true
      }
      ..
    }
  ]
}
```

##### Using Identity Match with micro-deposit or database Items

Items verified via [Same Day micro-deposits](/docs/auth/coverage/same-day/), [Instant micro-deposits](/docs/auth/coverage/instant/#instant-micro-deposits), or [Database Auth](/docs/auth/coverage/database-auth/) are not typically compatible with any other Plaid products besides Auth and Transfer. An exception is Identity Match, which can be used with these Items if they were previously seen on the Plaid network (approximately 30% of these Items).

Note that this applies to the Identity Match endpoint only; Items created using these flows are not compatible with [`/identity/get`](/docs/api/products/identity/#identityget).

To use Identity Match with micro-deposit or database Items, you have two options:

- For a low-code option that allows you to run the Identity Match check within the Link flow, use the [Account Verification Dashboard](/docs/identity/#integrating-identity-match-via-the-account-verification-dashboard).
- To run the Identity Match check outside the Link flow, continue with the [integration instructions in this section](/docs/identity/#integration-instructions).

###### Integration instructions

1. To enable the Item for Identity Match, `identity` must be included in the `optional_products`, `required_if_supported_products`, or `additional_consented_products` array during the [`/link/token/create`](/docs/api/link/#linktokencreate) call at Item creation. (For this flow, `identity` should not be included in the `products` array, since including any other product than `auth` in the `products` array can prevent micro-deposit or database verification flows from activating.)
   - If this was not done during Item creation, you can Identity consent to an existing Item by sending it through [update mode](/docs/link/update-mode/#requesting-additional-consented-products).
2. Call [`/identity/match`](/docs/api/products/identity/#identitymatch).

If the Item is supported by Identity Match, you will receive results. If the Item was not previously seen on the Plaid network and therefore cannot be used with Identity Match, you will receive a `PRODUCTS_NOT_SUPPORTED` error and will not be billed for the API call.

#### Integrating Identity Match via the Account Verification Dashboard

If you are using [Auth](/docs/auth/), you can integrate Identity Match via the [Account Verification Dashboard](https://dashboard.plaid.com/account-verification). This is a lower-code option that allows you to rely on Plaid Link for checking match scores and handling failed sessions.

When enabling Identity in this way, the Identity check will be incorporated into the Link flow.

This configuration option can only be used if Auth and Identity are the only required Plaid products you are using in the session. If you need access to other products, use the [standard integration flow](/docs/identity/#integration-overview).

1. In the [Dashboard](https://dashboard.plaid.com/account-verification), select your Link customization and enable the "Configure via Dashboard" toggle.
2. Select the settings for Auth you would like to use, then click "Next".
3. Select the rules to use for Identity Match. Typically, the match threshold for each field should be set to a score of 70 or higher.
4. When calling [`/link/token/create`](/docs/api/link/#linktokencreate):
   - Ensure that `identity` is in the `products` array (or `required_if_supported_products` array, if you want to allow customers to link accounts if their bank does not support Identity).
   - Set the `link_customization_name` to the customization you selected in Step 1 (this is optional if you are using the `default` customization).
   - Provide the identity data you have about the user, such as their `legal_name`, `phone_number`, `address`, and `email_address`.
5. Launch Link using the Link token. When the user goes through the Link flow, the identity data you supplied in the previous step will be compared to the data retrieved from their linked financial account. If the result based on your configuration is a pass, they will see a success screen. If it is a failure, they will see an error.
6. Watch for the `IDENTITY_MATCH_PASSED` or `IDENTITY_MATCH_FAILED` event from Link to detect the result of Identity Match in the Link flow. You can also view the status of a completed session in the [Dashboard](https://dashboard.plaid.com/account-verification).

If the Identity Match check fails, you will not receive a `HANDOFF` event or `onSuccess` callback, but if you want to proceed with creating an Item anyway, you can obtain the `public_token` by calling [`/link/token/get`](/docs/api/link/#linktokenget).

If Identity is enabled via the Account Verification Dashboard, you will be billed for Identity Match as long as Plaid could obtain any Identity data from the Item, even if you did not exchange the `public_token`. If Plaid was unable to obtain any Identity data for the Item, you will not be billed.

/link/token/create

```
const request: LinkTokenCreateRequest = {
  loading_sample: true
};
try {
  const response = await plaidClient.linkTokenCreate(request);
  const linkToken = response.data.link_token;
} catch (error) {
  // handle error
}
```

#### Testing Identity

Identity can be tested in Sandbox against test data. Plaid provides a [GitHub repo](https://github.com/plaid/sandbox-custom-users/) with test data for testing Identity in Sandbox, helping you test configuration options beyond those offered by the default Sandbox user. For more information on configuring custom Sandbox data, see [Configuring the custom user account](/docs/sandbox/user-custom/#configuring-the-custom-user-account).

[`/identity/match`](/docs/api/products/identity/#identitymatch) can be tested in Sandbox using custom user accounts as described above. Note that it is not recommended to send real user data or PII in the Sandbox environment for use with [`/identity/match`](/docs/api/products/identity/#identitymatch), as Sandbox is meant for testing purposes.

For details on testing Identity Match with Same Day Micro-deposits, see [Testing Same-Day Microdeposits](/docs/auth/coverage/testing/#testing-same-day-micro-deposits).

When testing Identity Match using the default `user_good` custom Sandbox user, you can expect the following identity data to be used for matching against:

```
Alberta Bobbeth Charleson
2992 Cameron Rd.
Malakoff, NY 14236
Phone: +1-111-222-3333
Email: accountholder0@example.com
```

#### Identity Document Upload

The [Identity Document Upload](/docs/identity/identity-document-upload/) add-on can be used to verify account ownership based on a bank statement uploaded by an end user. This feature is intended primarily for use with Items created via loginless Auth flows, such as Same-Day Micro-deposits, Instant Micro-deposits, or Database Insights. For more details, see [Identity Document Upload](/docs/identity/identity-document-upload/).

#### Sample app tutorial and code

For a real-life example of an app that incorporates Identity, see the Node-based [Plaid Pattern Account Funding](https://github.com/plaid/pattern-account-funding) sample app. Pattern Account Funding is a sample account funding app that fetches Identity data in order to verify identity prior to a funds transfer. The Identity code can be found in [items.js](https://github.com/plaid/pattern-account-funding/blob/master/server/routes/items.js#L81-L116).

For a tutorial walkthrough of creating a similar app, see [Account funding tutorial](https://github.com/plaid/account-funding-tutorial).

#### Identity pricing

Identity is billed on a [one-time fee model](/docs/account/billing/#one-time-fee), meaning you will be billed once for each Item with Identity added to it. Identity Match is billed on [a per-request flat-fee basis](/docs/account/billing/#per-request-flat-fee). To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/). For more details about pricing and billing models, see [Plaid billing](/docs/account/billing/).

#### Identity availability by country

Identity and Identity Match are available in all supported countries (US, Canada, and Europe, including the UK). In Canada, Identity Match cannot be requested via the Production access form and is available via a Growth or Custom plan only. To request access to Identity Match in Canada, contact your account manager or [Sales](https://plaid.com/contact/).

#### Next steps

To get started building with Identity, see [Add Identity to your App](/docs/identity/add-to-app/).

If you're ready to launch to Production, see the Launch Center.

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
