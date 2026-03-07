---
title: "Identity Verification - Managing failed verifications | Plaid Docs"
source_url: "https://plaid.com/docs/identity-verification/failed-verifications/"
scraped_at: "2026-03-07T22:04:54+00:00"
---

# Operational guide

#### Operational guidance for Identity Verification, including how to manage failed verifications

#### Overview

The most common reason for Identity Verification failures is bad actors and attempted fraud. However, sometimes a genuine user may fail a check due to various reasons, such as accidentally entering incorrect information, a lack of records available to confirm their identity, or issues with their identity documents. This page discusses common reasons why legitimate users may fail checks and how you can handle these issues.

#### Common causes for Data Source Verification failures

##### No match

Causes include:

- Recent personal information changes, such as a change of address or a legal name change
- If you are providing any user information via the API rather than Link, it must use the correct format. For example, if you send DOB as month/day/year rather than in ISO format, it will cause date of birth mismatches. See [Input validation rules](/docs/identity-verification/hybrid-input-validation/).

##### No data

Causes include:

- Phone number appears less consistently in data sources than other fields, and is most likely to result in "no data"
- Recent immigrants and younger end users are less likely to have data on file
- Countries other than the US have less complete records and are less likely to have data source matches

#### Common causes for Documentary Verification failures

- Expired documents will fail Documentary Verification, even if they are genuine
- Documents that do not match the name and date of birth entered will fail

#### Common causes for Risk Check failures

See [Risk Checks](/docs/identity-verification/risk-checks/) for more information.

#### Common causes for Selfie Check failures

Causes generally involve low-quality images, which can be caused by:

- Poor quality camera, such as one found on a very old, low-end device
- Insufficient lighting
- Improper framing (head cut off or not facing camera)
- Excessive glare from glasses
- Motion blur while taking photo

Plaid provides guidance in Link to help users take appropriate selfies.

#### Managing failures

If you have reason to believe that the end user's information is correct, you can override a failure via the Dashboard by clicking the **Override Result** button on the verification detail page in the Dashboard.

To allow a user to retry a verification, you can issue a retry. For instructions on the process, see [Retries](/docs/identity-verification/#retries).

#### Reporting issues

If at any point you believe that the Identity Verification result was incorrect, you can use the **Report a Problem** button in the Identity Verification Dashboard to inform Plaid. If you require a response, file a ticket via the Dashboard instead.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
