---
title: "Auth - Additional Auth flows | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/"
scraped_at: "2026-03-07T22:04:31+00:00"
---

# Additional verification flows

#### Support Auth for more US financial institutions and users

#### Overview

Instant Auth covers approximately 95% of users' eligible bank accounts.
For the remaining ~5%, or users who choose not to use instant credential verification, Plaid offers users the ability to
verify their accounts through an optimized set of database validation and micro-deposit flows.

Our video tutorial also provides a comprehensive guide to adding Auth using micro-deposits

The following flows are enabled in your app by default once you have integrated Auth:

- [Instant Auth](/docs/auth/coverage/instant/): User enters their credentials and is authenticated immediately. This is the default flow.
- [Instant Match](/docs/auth/coverage/instant/#instant-match) (US only): User enters their credentials, account number, and routing number. Plaid matches user input against masked values and authenticates immediately.

To increase conversion, you can enable the following additional Auth flows:

- [Automated Micro-deposits](/docs/auth/coverage/automated/) (US only): User enters their credentials, account number, and routing number. Plaid makes a micro-deposit and automatically verifies the deposit in as little as one to two business days.
- [Database Auth](/docs/auth/coverage/database-auth/) (US only): User enters account and routing numbers. Plaid attempts to match this data against known-good account and routing numbers recognized on the Plaid network.
- [Instant Micro-deposits](/docs/auth/coverage/instant/#instant-micro-deposits) (US only): User enters account and routing numbers. Plaid makes a RTP or FedNow micro-deposit and the user manually verifies the code in as little as 5 seconds.
- [Same Day Micro-deposits](/docs/auth/coverage/same-day/) (US only): User enters account and routing numbers. Plaid makes a Same Day ACH micro-deposit and the user manually verifies the code in as little as one business day.

To see each of these flows in action, check out the [Auth workflows demo](https://plaid.coastdemo.com/share/67d0ce0df465686c02cc4fd2?zoom=100).

##### Auth flow method comparison chart

| Method | Best for | Verification timeframe | Requires additional integration work |
| --- | --- | --- | --- |
| [Instant Auth](/docs/auth/coverage/instant/) | All risk profiles | Instant | No (enabled by default) |
| [Instant Match](/docs/auth/coverage/instant/#instant-match) | All risk profiles | Instant | No (enabled by default) |
| [Automated Micro-deposits](/docs/auth/coverage/automated/) | All risk profiles | 1-2 business days | No, but requires "pending verification" UX state in app |
| [Instant Micro-deposits](/docs/auth/coverage/instant/#instant-micro-deposits) | Low and medium risk profiles | Instant | No |
| [Same Day Micro-deposits](/docs/auth/coverage/same-day/) | Low and medium risk profiles | 1-2 business days, plus wait for user interaction | Yes, and requires "pending verification" UX state in app |
| [Database Auth](/docs/auth/coverage/database-auth/) | Low and medium risk profiles | Instant | No, unless used in advanced configuration mode |

Accounts verified via Database Auth, Instant Micro-deposits, or Same-Day Micro-deposits do not have active data connections with a financial institution. These accounts and their associated Items can only be used with Auth and Transfer, not with any other Plaid products (such as Balance or Transactions), with the partial exception of [Identity Match](/docs/identity/#using-identity-match-with-micro-deposit-or-database-items) and [Signal Transaction Scores](/docs/signal/signal-rules/#data-availability-limitations).

Because of this limitation, accounts verified via these flows may be more vulnerable to fraud or return risk. Consider your level of risk exposure and current anti-fraud measures before enabling these methods. For more details, see [Anti-fraud best practices](/docs/auth/coverage/same-day-link-best-practices/).

#### Entry points

To learn more about when Plaid presents Database Auth, Instant Micro-deposits, or Same-Day Micro-deposits as options, and to customize the Link user experience around these flows, see [Configuring entry points](/docs/auth/coverage/flow-options/).

#### Integration instructions

Most additional Auth flows can be enabled in the [Account Verification Dashboard](https://dashboard.plaid.com/account-verification) with no additional integration work, as long as the following requirements are met:

- The `country_codes` array in [`/link/token/create`](/docs/api/link/#linktokencreate) is set to `["US"]`. Additional Auth flows are not available if other countries are selected.
- The `products` array does not contain any products other than `auth` or `transfer`. Any other products you want to use must be configured in the `additional_consented_products`, `required_if_supported_products`, or `optional_products` arrays and will be used on a best-effort basis.

[Same-Day Micro-deposits](/docs/auth/coverage/same-day/) requires additional integration work, as you will need to present a Link UI where the end user can verify the micro-deposit description. For more details, see [Same-Day micro-deposits](/docs/auth/coverage/same-day/).

[Database Auth](/docs/auth/coverage/database-auth/), if used in the advanced mode where fallback rules are not configured in the Dashboard, requires additional integration work. For more details, see [Database Auth](/docs/auth/coverage/database-auth/).

![Alt text: 'Account Verification Dashboard with settings for instant verification methods, additional methods toggle, and micro-deposit options.'](/assets/img/docs/auth/av_dashboard_config.png)

Example of some configuration options available in the Account Verification Dashboard

#### Institution coverage

To see which Auth flows a given institution supports, you can call [`/institutions/get`](/docs/api/institutions/#institutionsget) with [`options.include_auth_metadata`](/docs/api/institutions/#institutions-get-request-options-include-auth-metadata) set to `true`. The results will be returned in the [`auth_metadata.supported_methods` object](/docs/api/institutions/#institutions-get-response-institutions-auth-metadata-supported-methods) in the response. Alternatively, you can see this information on the [Institution Status page](https://dashboard.plaid.com/activity/status) in the Plaid Dashboard. The Same Day Micro-deposits and Database Auth flows will not appear in these results, as they do not depend on institution capabilities and are available at all institutions.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
