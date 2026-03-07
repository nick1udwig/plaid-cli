---
title: "Monitor - Introduction to Monitor | Plaid Docs"
source_url: "https://plaid.com/docs/monitor/"
scraped_at: "2026-03-07T22:05:09+00:00"
---

# Introduction to Monitor

#### Sanction, PEP, and watchlist screening for anti-money laundering compliance

Get started with Monitor

[API Reference](/docs/api/products/monitor/)[Get started](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access)

#### Overview

Monitor detects if your customers are on government watchlists, using modern APIs that reduce false positives and increase efficiency. Monitor screens users against a number of sanction and PEP lists. Sanction lists are updated daily and can be used for daily automated re-scans. Any hits will be exposed via both a user-friendly dashboard UI and via API, for use with either manual or automated review workflows. Monitor also integrates directly with [Identity Verification](/docs/identity-verification/) for an end-to-end verification and KYC solution.

Monitor can be used to screen end users in any country. To integrate with Monitor, your company must be based in the US or Canada.

#### Creating a program

Monitor is not available by default in the Sandbox environment. To obtain Sandbox access, [request full Production access for Monitor](https://dashboard.plaid.com/settings/team/products), which will automatically grant Sandbox access, or [contact Sales](https://plaid.com/products/monitor/#contactForm). To obtain Sandbox access for Monitor if you are already using another Plaid product, you can also contact your account manager or submit an [access request ticket](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access).

The first step to integrating is logging into the dashboard and accessing the **New Program** page. Programs help you to define the settings and workflows for conducting watchlist screenings in your organization. For example, you can define match sensitivities, which lists are being screened, and the behavior of ongoing screening. For more details on the configuration of screenings and how matches are determined, see [Watchlists and Matching Algorithms](/docs/monitor/algorithms-and-lists/).

Monitor allows you to start a customer in one program and then move them to a different one. To move users between programs, call [`/watchlist_screening/individual/update`](/docs/api/products/monitor/#watchlist_screeningindividualupdate) or [`/watchlist_screening/entity/update`](/docs/api/products/monitor/#watchlist_screeningentityupdate) with the `watchlist_screening_id` you want to move and the ID of the program you want to move it to in the `watchlist_program_id` field. Customers cannot be in multiple programs at once.

The number of programs you wish to maintain will vary based on the complexity of your organization. Many companies choose to split up programs such that each program aligns to the compliance ownership within their organization and each reviewer will only see programs they are responsible for. Some common ways of splitting up programs are by product vertical ("Cardholders", "Personal Loans", "FDIC Accounts", etc.), geography ("US Cardholders", "European Cardholders", etc.), risk level ("High Risk Individuals", "Medium Risk Individuals"), or a combination of the above. For ideas about how to split customers up by risk, [OFAC has prepared a helpful risk matrix](https://cognitohq.com/blog/ofac-risk-matrix) that can provide a baseline for setting initial risk level and then modifying the risk level based on ongoing user behavior. You are allowed to change a program after it has been created, although note that increasing the sensitivity of a program after creation may result in a large influx of new hits.

If your use case involves an ongoing relationship with your customers, Plaid recommends creating at least two programs -- one with rescan enabled for active customers, and one with rescan disabled, for customers who have not yet established an ongoing relationship (for example, customers who have not completed the onboarding experience) or who have closed their accounts. For more details, see [Monitor fee model](/docs/account/billing/#monitor-fee-model).

Once your first program is created, you can start your integration.

#### Basic integration overview

Before integrating Monitor, make sure you have [Created a program](/docs/monitor/#creating-a-program).

This integration overview describes the process for using a Dashboard-based review workflow, which is the most common way to use Monitor. The review process can also be performed using the API; see the [API reference](/docs/api/products/monitor/) for details.

If you are already using Plaid's [Identity Verification](/docs/identity-verification/) product:

1. In the Identity Verification template editor, on the Setup tab, enable the **Screen customers** setting to automatically assign users to a Monitor program after they have completed verification.
2. When your user has successfully completed the Identity Verification process, the response you receive from calling [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget) will include a `watchlist_screening_id` value.
3. Call [`/watchlist_screening/entity/get`](/docs/api/products/monitor/#watchlist_screeningentityget) or [`/watchlist_screening/individual/get`](/docs/api/products/monitor/#watchlist_screeningindividualget) with this value, and check the `status` flag on the response.
   - If it is `cleared`, your customer did not appear on any of your target screening lists.
   - If it is `pending_review`, your customer received one or more hits and is pending review.
4. For customers who are pending review, a member of your team will need to review and update the user's status in the Dashboard. When they do (and you have configured your webhooks as [described below](/docs/monitor/#configuring-webhooks)) you will receive a webhook, `SCREENING: STATUS_UPDATED` (or `ENTITY_SCREENING: STATUS_UPDATED` for entities). This webhook will include the watchlist screening id of the individual.
5. You can then make an additional call to [`/watchlist_screening/entity/get`](/docs/api/products/monitor/#watchlist_screeningentityget) or [`/watchlist_screening/individual/get`](/docs/api/products/monitor/#watchlist_screeningindividualget) and check the `status` flag on the response.
   - If it is `cleared`, the screening hit has been dismissed by a reviewer as invalid.
   - If it is `pending_review`, a customer that was previously marked as `cleared` or `rejected` was marked for review.
   - If it is `rejected`, your customer received a hit that has been confirmed by a reviewer as a valid match.

If you are not using Identity Verification, the process is slightly different:

1. Create a customer object using [`/watchlist_screening/individual/create`](/docs/api/products/monitor/#watchlist_screeningindividualcreate) or [`/watchlist_screening/entity/create`](/docs/api/products/monitor/#watchlist_screeningentitycreate), passing along any verified information you have about your user. While not required, it is recommended to specify a `client_user_id`, which should refer to an internal ID in your database, to help associate this screening with your customer in the future.
2. The results of this `create` call will include a `status` flag -- either `cleared`, or `pending_review` as specified above.
3. At this point, you can proceed from step 4 in the above workflow.

#### Configuring webhooks

To complete your integration, you should add a webhook receiver endpoint to your application.

Like Identity Verification, you can configure the webhooks sent by Monitor by visiting the dashboard [webhook configuration](https://dashboard.plaid.com/developers/webhooks) page. Click **New webhook**, and then select which events you want to subscribe to. For Monitor, there are two:

- [`SCREENING: STATUS_UPDATED`](/docs/api/products/monitor/#screening-status_updated) when an individual's screening status has been updated
- [`ENTITY: STATUS_UPDATED`](/docs/api/products/monitor/#entity_screening-status_updated) when an entity's screening status has been updated

Enter the URL of your webhook receiver for the webhook you wish to subscribe to and click **Save**. Plaid will now send an HTTP POST request to the webhook receiver endpoint every time the event occurs.

For more information on webhooks, see the [webhook documentation](/docs/api/webhooks/).

#### Building a reviewer workflow

This guide assumes that you will be using the dashboard for reviewing hits, but you can also re-create any aspect of the dashboard in your internal applications via the [API](/docs/api/products/monitor/).

![Plaid Monitor Dashboard showing programs list with status, type, sensitivity, and monitoring options. Buttons for new programs.](/assets/img/docs/monitor/monitor-screenshot.png)

##### Assigning hits to reviewers

Once your basic integration is complete, any customers who have potential hits will start showing up on your dashboard. You can assign any screenings that are pending review to different reviewers in your organization. If using an API-based review workflow, you can get a full list of users and their associated user IDs by calling [`/dashboard_user/list`](/docs/api/kyc-aml-users/#dashboard_userlist) and setting your screening `assignee` to the desired reviewer.

#### Preparing for ongoing screening

Monitor updates screening lists daily and supports ongoing daily rescans of your entire customer base to alert you when new hits are discovered. This system is designed around the concept of a living `pending_review` queue that is updated whenever new hits are found. We recommend that a reviewer log in daily to check the review status on the dashboard. If you prefer to be alerted when a hit is found or to handle hits via automation, you can set up a daily automated job to check [`/watchlist_screening/individual/hit/list`](/docs/api/products/monitor/#watchlist_screeningindividualhitlist) or [`/watchlist_screening/entity/hit/list`](/docs/api/products/monitor/#watchlist_screeningentityhitlist) once a day to poll any of the screenings that end up in this queue. Then use the associated `client_user_id` or `id` to tie the hit back to your internal database and determine if action is required.

Plaid will also fire [webhooks](https://plaid.com/docs/api/products/monitor/) to alert you when any new hits are found. Webhooks will contain the watchlist screening id, which you can use to go retrieve any watchlist screening hits that are pending review.

#### Generating reports

You can generate and export reports on Monitor activity via the Dashboard, under [Monitor->Reports](https://dashboard.plaid.com/monitor/reports). Reports are provided in .csv format.

You can also generate reports using the Monitor API. For example, you can use [`/watchlist_screening/individual/list`](/docs/api/products/monitor/#watchlist_screeningindividuallist) or [`/watchlist_screening/entity/list`](/docs/api/products/monitor/#watchlist_screeningentitylist) to generate the list of individuals or entities you'd like to report on, then call [`/watchlist_screening/individual/hit/list`](/docs/api/products/monitor/#watchlist_screeningindividualhitlist) or [`/watchlist_screening/entity/hit/list`](/docs/api/products/monitor/#watchlist_screeningentityhitlist) for each target individual or entity to generate a list of their hits.

#### Additional considerations

- Adding Monitor is not the only part of a successful Anti-Money Laundering program. We highly recommend that you consult with an AML professional to help you build all of the policies and procedures required. AML regulations vary widely for different jurisdictions and industries.
- Monitor **does not blanket ban any countries on your behalf**. Some sanction programs effectively prohibit working with certain jurisdictions and it is up to you to ensure that you are only working with individuals from countries that are appropriate for your desired risk levels. For instance, here is an [article from OFAC about a "Country List"](https://www.treasury.gov/resource-center/sanctions/Programs/Pages/faq_10_page.aspx)
- Monitor is built to make it difficult to overconstrain a search and therefore have false negatives. That being said, if your data inputs are not quality controlled to at least a basic extent, this can result in false negatives or false positives (for example, broad searches based on incomplete names).

#### Testing Monitor

Monitor is not available by default in the Sandbox environment. To obtain Sandbox access, [request full Production access for Monitor](https://dashboard.plaid.com/settings/team/products), which will automatically grant Sandbox access, or [contact Sales](https://plaid.com/products/monitor/#contactForm). To obtain Sandbox access for Monitor if you are already using another Plaid product, you can also contact your account manager or submit an [access request ticket](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access).

Sandbox supports a real-world dataset, meaning that sanctioned individuals should show up just as they would in Production. Some example sanctioned individuals and entities you can use for testing purposes include:

| Individual Name | Date of Birth | Location | Document |
| --- | --- | --- | --- |
| Joseph Kony | 1964-09-18 | CF |  |
| Jianshu Zhou | 1971-07-15 | CN | E09598913 |
| Ermias Ghermay | Any day in 1980 | LY |  |

| Entity Name | Location | Document |
| --- | --- | --- |
| Iran Mobin Electronic Development Company | IR | 10103492170 |
| МИНСКИЙ ОТРЯД МИЛИЦИИ ОСОБОГО НАЗНАЧЕНИЯ | BY |  |
| Islamic Radio And Television Union | AF |  |

Testing Monitor alongside Identity Verification can be a little tricky, as Monitor screens against a real-world dataset, while Identity Verification only supports one single test identity, Leslie Knope. If you wish to test the case of Identity Verification succeeding, but a Monitor screening encountering a hit, you can try one of these options:

- From the Monitor section of the Dashboard, select the user from the **Cleared** tab and change their status from "Cleared" to "Pending Review" or "Rejected". This will fire a `STATUS_UPDATED` webhook, and the user's new screening status will be reflected in the response the next time you call [`/watchlist_screening/individual/get`](/docs/api/products/monitor/#watchlist_screeningindividualget).
- You can also create additional watchlist screenings against the same client user ID by calling [`/watchlist_screening/individual/create`](/docs/api/products/monitor/#watchlist_screeningindividualcreate) with the name and birthday of a real sanctioned individual. This screening will appear in the Monitor dashboard in the "Pending review" section and will be included in the array of screening results that get returned if you call [`/watchlist_screening/individual/list`](/docs/api/products/monitor/#watchlist_screeningindividuallist) with this client user ID.

Note that the Sandbox environment sanction and watchlist dataset may be out of date. Make sure all real checks are carried out on the live Production environment.

#### Additional resources

For a step-by-step walkthrough of implementing Monitor, see the [Identity Verification and Monitor solution guide](https://plaid.com/documents/plaid-idv-monitor-solution-guide.pdf)

#### Monitor pricing

Monitor has a base fee for each new user scanned, as well as a separate rescanning fee, determined by the number of users rescanned each month. For more details, see the [Monitor fee model](/docs/account/billing/#monitor-fee-model). To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/).

#### Next steps

To learn more about building with Monitor, see the [Monitor API Reference](/docs/api/products/monitor/).

If you're ready to launch to Production, see the Launch checklist.

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
