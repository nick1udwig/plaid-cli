---
title: "Link - Institution status in Link | Plaid Docs"
source_url: "https://plaid.com/docs/link/institution-status/"
scraped_at: "2026-03-07T22:05:05+00:00"
---

# Institution status in Link

#### Details of institution health status accessible in Link

#### Institution status in Link

Link proactively lets users know if an institution's connection isn't performing well.
Below are the three views a user will see depending on the status of the institution they select.

When the status of `item_logins` of an institution is `DEGRADED`, Link will warn
users that they may experience issues and allow them to continue. Once the
status of `item_logins` becomes `DOWN`, Link will block users from attempting to
log in and suggest they find another bank.

![Plaid Link institution status panels: Healthy allows credential input; Degraded shows connectivity issue; Down indicates connection failure.](/assets/img/docs/institution-status-in-link.png)

Institution Health warnings can be tested in the Sandbox environment by using one of the dedicated ["Unhealthy" Platypus Bank test institutions](/docs/sandbox/institutions/#institution-details).

For a more detailed view of institution status, see the [status dashboard](https://dashboard.plaid.com/activity/status), which provides a browsable view of institutions, supported products, and institution health.

#### Connectivity not supported

![Plaid Link message: 'Connectivity not supported' for simple institution. Advises trying another institution.](/assets/img/docs/link-connectivity-not-supported.png)

If an institution is supported by Plaid, but is not supported by the product set or country codes Link was initialized with, a "Connectivity not supported" error message will appear. This message does not reflect the health of the institution. This message can be resolved by calling [`/link/token/create`](/docs/api/link/#linktokencreate) with a more minimal product set, or by making sure you are Production-enabled for the country you are initializing Link for. For more details, see [Link troubleshooting](/docs/link/troubleshooting/#missing-institutions-or-connectivity-not-supported-error).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
