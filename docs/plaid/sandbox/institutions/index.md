---
title: "Sandbox - Sandbox institutions | Plaid Docs"
source_url: "https://plaid.com/docs/sandbox/institutions/"
scraped_at: "2026-03-07T22:05:16+00:00"
---

# Sandbox institutions

#### View institutions and institution test configurations available in the Sandbox

#### Institution details

All of the institutions that are available in the Plaid Production environment are also available in Sandbox. Plaid also provides several Sandbox-only institutions to write integration tests against:

| Institution name | ID |
| --- | --- |
| First Platypus Bank (non-OAuth bank) | `ins_109508` |
| First Platypus Balance Bank (non-OAuth bank) | `ins_130016` |
| First Gingham Credit Union (non-OAuth bank) | `ins_109509` |
| Tattersall Federal Credit Union (non-OAuth bank) | `ins_109510` |
| Tartan Bank (non-OAuth bank) | `ins_109511` |
| Houndstooth Bank (for Instant Match and Automated Micro-deposit testing) | `ins_109512` |
| Tartan-Dominion Bank of Canada (Canadian bank) | `ins_43` |
| Flexible Platypus Open Banking (UK Bank) | `ins_116834` |
| Royal Bank of Plaid (UK Bank) | `ins_117650` |
| Platypus OAuth Bank (for [OAuth testing](/docs/link/oauth/#testing-oauth)) | `ins_127287` |
| First Platypus OAuth App2App Bank (for [App-to-App OAuth testing](/docs/link/oauth/#app-to-app-authentication)) | `ins_132241` |
| Flexible Platypus Open Banking (for [OAuth QR code authentication](/docs/link/oauth/#qr-code-authentication) testing) | `ins_117181` |
| Windowpane Bank (for [Instant Micro-deposit testing](/docs/auth/coverage/instant/#instant-micro-deposits)) | `ins_135858` |
| Unhealthy Platypus Bank - Degraded | `ins_132363` |
| Unhealthy Platypus Bank - Down | `ins_132361` |
| Unsupported Platypus Bank - Institution Not Supported | `ins_133402` |
| Platypus Bank RUX Auth (formerly for testing legacy RUX flows) | `ins_133502` |
| Platypus Bank RUX Match (formerly for testing legacy RUX flows) | `ins_133503` |

All Sandbox institutions can be accessed using the [Sandbox test credentials](/docs/sandbox/test-credentials/).

Note that when testing OAuth in Sandbox, the OAuth flow for Platypus will be shown instead of institution-specific flows, and institution-specific behaviors will not occur, except at test institutions as shown above.

None of the institutions in the list above use OAuth flows except for those that are explicitly marked as OAuth or European / UK institutions.

#### Institution details for Auth testing

You can trigger specific Auth flows in the Sandbox environment by using the following values.
For detailed instructions on testing Auth in the Sandbox, see [Testing Auth in the Sandbox](/docs/auth/coverage/testing/).

| Step | Instant Match | Automated Micro-deposits | Same-Day Micro-deposits | Instant Micro-deposits |
| --- | --- | --- | --- | --- |
| Institution | Houndstooth Bank | Houndstooth Bank | N/A | Windowpane Bank |
| Username | user\_good | user\_good | N/A | N/A |
| Password | pass\_good | microdeposits\_good | N/A | N/A |
| Account | Plaid Saving (\*\*\*\*1111) | Plaid Checking (\*\*\*\*0000) | Checking or Savings | Checking or Savings |
| Routing number | 021000021 | 021000021 | 110000000 | 333333334 |
| Account number | 1111222233331111 | 1111222233330000 | 1111222233330000 | 1111222233330000 |
| Micro-deposit code | N/A | N/A | ABC | ABC |
| Wait before Sandbox verification | Instant | About 24 hours | Verify interactively within minutes | Instant |

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
