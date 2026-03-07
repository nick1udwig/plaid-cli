---
title: "Consumer Report (by Plaid Check) - Onboard users with Layer | Plaid Docs"
source_url: "https://plaid.com/docs/check/onboard-users-with-layer/"
scraped_at: "2026-03-07T22:04:48+00:00"
---

# Onboard users with Layer

#### A guide to integrating Layer with Plaid Check

#### Overview

This guide will walk you through how to onboard users to Plaid Check with Layer. This process involves creating a user, onboarding the user with Layer, and retrieving the user's Plaid Check data.

#### Integration steps

##### Create a user

1. Create a `user_id` using the [`/user/create`](/docs/api/users/#usercreate) endpoint. At this stage, the `identity` field is optional.

##### Onboard the user with Layer

1. Onboard the user by following Layer's [Integration Overview steps](/docs/layer/#integration-overview).  
    When creating the Link token via [`/session/token/create`](/docs/api/products/layer/#sessiontokencreate), pass in the `user_id` you created in the `user.user_id` field and be sure the template used has CRA products enabled.
2. Call [`/user_account/session/get`](/docs/api/products/layer/#user_accountsessionget) to retrieve user-permissioned identity information.
3. After processing the identity information, update the Plaid user record by calling the [`/user/update`](/docs/api/users/#userupdate) endpoint. Populate the `identity` field with the relevant information. At minimum, the following fields must be provided and non-empty: `name`, `date_of_birth`, `emails`, `phone_numbers`, and `addresses` (with at least one email address, phone number, and address designated as `primary`). Providing at least a partial SSN via the `id_numbers` field is highly recommended, since it improves the accuracy of matching user records during compliance processes such as file disclosure, dispute, or security freeze requests.

##### Retrieve the user's Plaid Check data

1. Now that the identity information has been added to the user record, you can generate a Consumer Report by calling [`/cra/check_report/create`](/docs/api/products/check/#cracheck_reportcreate).
2. Once the Consumer Report has been generated (indicated by a `USER_CHECK_REPORT_READY` webhook), retrieve the report by calling [`/cra/check_report/base_report/get`](/docs/api/products/check/#cracheck_reportbase_reportget), [`/cra/check_report/income_insights/get`](/docs/api/products/check/#cracheck_reportincome_insightsget), [`/cra/check_report/partner_insights/get`](/docs/api/products/check/#cracheck_reportpartner_insightsget), or [`/cra/check_report/pdf/get`](/docs/api/products/check/#cracheck_reportpdfget).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
