---
title: "Identity Verification - Metrics and reporting | Plaid Docs"
source_url: "https://plaid.com/docs/identity-verification/reporting/"
scraped_at: "2026-03-07T22:04:55+00:00"
---

# Generating metrics

#### Generating metrics for Identity Verification

You can view the status of any Identity Verification session in the Dashboard. To obtain aggregated data such as the percentage of sessions that were completed or the percentage that passed verification vs. were rejected, you will need to use the API to generate the data.

#### Conversion and success rates

The *conversion rate* is defined as the percentage of sessions begun that were completed, regardless of whether the user passed or failed verification.

The *success rate* is defined as the percentage of completed sessions that resulted in the user passing verification.

#### Calculating success rates

The most comprehensive way to measure overall success rates is to use [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget), since it includes backend-only sessions, sessions generated using shareable links, and manual overrides after the session is completed.

If the status is `success`, the verification succeeded.

If the status is `failed` or `pending review`, the verification did not succeed.

The success rate is the number of sessions with the `success` status divided by the total of all sessions with the `success`, `failed`, or `pending review` status.

Sessions with a status other than `success`, `failed`, or `pending review` should be discarded for the purpose of calculating overall success rates, since they represent sessions that were not completed.

#### Calculating conversion rates

If you want to measure the conversion rate, you can use the `onEvent` callback. A session is started if you receive the `IDENTITY_VERIFICATION_START_STEP` event. It was completed if you receive either the `IDENTITY_VERIFICATION_PASS_SESSION` or `IDENTITY_VERIFICATION_FAIL_SESSION` event, and it was successful only if you receive the `IDENTITY_VERIFICATION_PASS_SESSION`. To correlate different events with the same Link session, use the `link_session_id`. For more details, see [Link callbacks](/docs/identity-verification/link/).

For Identity Verification, `onEvent` callback information is not available via the [`/link/token/get`](/docs/api/link/#linktokenget) endpoint.

Alternatively, you can calculate conversion using [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget). The conversion rate is the number of sessions with a status of `success`, `failed`, or `pending review` divided by the total number of unique identity verification IDs. Because this metric includes sessions where the result was manually overridden, as well as backend-only sessions, it will provide different results from Link-based conversion metrics, especially if you use a combination of backend-based sessions and Link-based sessions. Whether it makes more sense to report on overall conversion or Link-based conversion will depend on your use case.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
