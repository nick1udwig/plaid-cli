---
title: "Auth - Test in Sandbox | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/testing/"
scraped_at: "2026-03-07T22:04:33+00:00"
---

# Testing Auth flows

#### Testing Instant Match and micro-deposit flows

#### Sandbox configuration

To test in the [Sandbox environment](/docs/sandbox/), you need to set the Link configuration environment to `"sandbox"`:

App.js

```
Plaid.create({
  // ...
  env: 'sandbox',
});
```

You will also want to direct any backend API requests to the sandbox URL:

API host

```
https://sandbox.plaid.com
```

#### Testing Instant Match

##### Test credentials

| Sandbox input | Successful credentials | Erroneous credentials |
| --- | --- | --- |
| Institution Name | Houndstooth Bank (`ins_109512`) | –– |
| Username | `user_good` | –– |
| Password | `pass_good` | –– |
| Account Selection | `Plaid Savings (****1111)` | –– |
| Routing number | `021000021` | Any other routing number |
| Account number | `1111222233331111` | Any other account number |

##### Entering the Link flow

- Search for “Houndstooth Bank” in Link.
- Use the input data in the table above to simulate the desired outcome.

#### Testing Instant Micro-deposits

You can also test this flow in Sandbox with no code by using the [Link Demo](/demo/).

| Sandbox input | Successful credentials | Erroneous credentials |
| --- | --- | --- |
| Institution name | Windowpane Bank (`ins_135858`) | –– |
| Routing number | `333333334` | Any other value (if testing upgrade path) |
| Account number | `1111222233330000` | –– |
| Deposit code | `ABC` | Any other value |

##### Entering the primary Link flow

- Search for "Windowpane Bank" in Link.
- When prompted, use the input data in the table above to simulate the desired outcome.
- For all other fields, you can use any values.

##### Testing the upgrade path from Same Day Micro-deposits to Instant Micro-deposits

- If you have enabled Same Day Micro-deposits, enter gibberish into the search bar to ensure no institutions will be found, then click on Link with account numbers.
- If you have configured Same Day Micro-deposits with Auth Type Select, choose "Manually" from the menu.
- When prompted, enter the input data in the table above. For all other fields, you can use any values.
- Note that, unlike the primary Instant Micro-deposits flow, the routing and account number entries will be on separate panes.

#### Testing Database Auth or Database Insights

You can also test this flow in Sandbox with no code by using the [Link Demo](/demo/).

##### Test credentials (US-based sessions)

To simulate a `database_insights_pass` result:

| Sandbox input | Credentials |
| --- | --- |
| Routing number | Any routing number (e.g. `110000000`) |
| Account number | `1111222233331111` |

To simulate a `database_insights_pass_with_caution` result:

| Sandbox input | Credentials |
| --- | --- |
| Routing number | Any routing number (e.g. `110000000`) |
| Account number | `1111222233333333` |

Any other account number will simulate a `database_insights_fail` result, e.g.:

| Sandbox input | Credentials |
| --- | --- |
| Routing number | Any routing number (e.g. `110000000`) |
| Account number | `1111222233332222` |

To simulate a 100 name match score, use the name "John Smith".

##### Test credentials (CA-based sessions)

To simulate a `database_insights_pass` result:

| Sandbox input | Credentials |
| --- | --- |
| Institution number | `110` |
| Branch number | `11111` |
| Account number | `11223311` |

To simulate a `database_insights_pass_with_caution` result:

| Sandbox input | Credentials |
| --- | --- |
| Institution number | `110` |
| Branch number | `11111` |
| Account number | `11223333` |

Any other input will simulate a `database_insights_fail` result, e.g.:

| Sandbox input | Credentials |
| --- | --- |
| Institution number | `110` |
| Branch number | `11111` |
| Account number | `11223322` |

To simulate a 100 name match score, use the name "John Smith".

##### Entering the Link flow

- If you have configured Database Auth or Database Insights with Auth Type Select, choose "Manually" from the menu.
- If you have configured Database Auth or Database Insights with Embedded Search, choose "Connect Manually".
- If you have enabled Database Auth or Database Insights without either of the above options, enter gibberish into the search bar to ensure no institutions will be found, then click on "Link with account numbers".
- Use input data from the table above to simulate various outcomes.

#### Testing Automated Micro-deposits

You can also test this flow in Sandbox with no code by using the [Link Demo](/demo/).

##### Test credentials

| Sandbox input | Successful credentials | Erroneous credentials |
| --- | --- | --- |
| Institution Name | Houndstooth Bank (`ins_109512`) | –– |
| Username | `user_good` | –– |
| Password | `microdeposits_good` | –– |
| Account Selection | `Plaid Checking (****0000)` | –– |
| Routing number | `021000021` | Any other routing number |
| Account number | `1111222233330000` | Any other account number |

##### Entering the Link flow

- Search for “Houndstooth Bank” in Link.
- Use the input data in the table above to simulate the desired outcome.

The micro-deposit verification will automatically succeed after twenty-four hours. To test a failed micro-deposit, or to skip the twenty-four hour waiting period, use the [`/sandbox/item/set_verification_status`](/docs/api/sandbox/#sandboxitemset_verification_status) endpoint to manually control the Item's micro-deposit verification status.

#### Testing Same Day Micro-deposits

You can also test this flow with no code by using the [Link Demo](/demo/).

##### Test credentials

| Sandbox input | Successful credentials | Erroneous credentials |
| --- | --- | --- |
| Routing number | `110000000` | –– |
| Account number | `1111222233330000` | –– |
| Deposit code | `ABC` | Any other value |

##### Initiating micro-deposits in Link

- If you have configured Same Day Micro-deposits with Auth Type Select, choose "Manually" from the menu
- If you have enabled Same Day Micro-deposits without Auto Type Select, enter gibberish into the search bar to ensure no institutions will be found, then click on "Link with account numbers".
- Use the input data in the table above to simulate the desired outcome.

##### Verifying micro-deposits in Link

- Call [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) with your `public_token` from the Link session in previous step to receive an `access_token`.
- Call [`/link/token/create`](/docs/api/link/#linktokencreate) and provide the `access_token` from the previous step to receive a `link_token`.
- Open Link with your `link_token`.
- In the deposit code field, enter `ABC`

#### Testing micro-deposit events

Micro-deposits that are generated in Sandbox will never be posted by default.
In order to generate a posted event that you can see when you query [`/bank_transfer/event/sync`](/docs/api/products/auth/#bank_transfereventsync), you can use the [`/sandbox/bank_transfer/simulate`](/docs/bank-transfers/reference/#sandboxbank_transfersimulate) endpoint.

Simulating a posted event in Sandbox does not generate a webhook. You will need to call [`/sandbox/bank_transfer/fire_webhook`](/docs/bank-transfers/reference/#sandboxbank_transferfire_webhook) each time you want a
webhook to be published for testing.

#### Testing Identity Match with Same Day micro-deposit Items

| Sandbox input | Successful credentials | Erroneous credentials |
| --- | --- | --- |
| Routing number | `011401533` | `110000000` |
| Account number | `1111222233330000` | `1111222233330000` |
| Deposit code | `ABC` | `ABC` |

- Complete the Same Day Micro-deposits flow as detailed above
- Call [`/identity/match`](/docs/api/products/identity/#identitymatch) using any variation of data in order to test the matching algorithm
- The identity data of the test user is as follows:

| Field | Sandbox identity |
| --- | --- |
| Name | Alberta Bobbeth Charleson |
| Email | [accountholder0@example.com](mailto:accountholder0@example.com) |
| Phone | 1112223333 |
| Street | 2992 Cameron Road |
| City | Malakoff |
| Region | NY |
| Country | US |
| Postal Code | 14236 |

#### Testing Database Match

You can also test this flow in Sandbox with no code by using the [Link Demo](/demo/).

##### Test credentials

To simulate a `database_matched` result:

| Sandbox input | Credentials |
| --- | --- |
| Routing number | `110000000` |
| Account number | `1111222233331111` |
| Name | `John Smith` |

Any other input will simulate no match and proceed with Same Day Micro-deposits, e.g.:

| Sandbox input | Credentials |
| --- | --- |
| Routing number | `110000000` |
| Account number | `1111222233332222` |
| Name | `John Smith` |

##### Entering the Link flow

- If you have enabled Database Match, enter gibberish into the search bar to ensure no institutions will be found, then click on Link with account numbers.
- If you have configured Database Match with Auth Type Select, choose "Manually" from the menu.
- Use the input data in the table above to simulate different outcomes.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
