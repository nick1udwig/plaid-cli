---
title: "Account - Activity, logs, and status | Plaid Docs"
source_url: "https://plaid.com/docs/account/activity/"
scraped_at: "2026-03-07T22:03:45+00:00"
---

# Dashboard logs and troubleshooting

#### Discover logging and troubleshooting information available in the Plaid Dashboard

#### Logs for webhooks, Link, and API requests

The Plaid Dashboard [Activity Log](https://dashboard.plaid.com/activity/usage) shows the past 14 days of API activity. Using the dashboard, you can see a record of all requests and responses, as well as all webhooks sent by the Plaid API, and all Link events.

![Plaid API logs showing request types, descriptions, institutions, environments, timestamps, and response codes in a dashboard.](/assets/img/docs/activity-logs.png)

You can view the details of any request, response, webhook, or event, and view error information for any failed API request.

![Request and response details for failed API call: invalid public token expired. Error code 400, Halifax, Bank API, success rate 100%.](/assets/img/docs/request-details.png)

##### Link analytics

The [Link analytics](https://dashboard.plaid.com/link-analytics) page in the Dashboard shows a summary of Link conversion, along with top Link errors your users are experiencing and recommendations for increasing conversion. For more details, see [Link conversion](/docs/link/best-practices/).

#### Logs for billable activity

The Plaid Dashboard [Usage Page](https://dashboard.plaid.com/activity/usage) shows billable API usage for most Plaid products. Usage data is also available via the Plaid MCP server, allowing you to interact with this data via chat or an LLM agent. For more details, see the [blog post on the Plaid Dashboard MCP server](https://plaid.com/blog/plaid-mcp-ai-assistant-claude/) and the [Plaid Dashboard MCP server documentation](/docs/resources/mcp/).

#### Logs for Payment Initiation, Transfer, and Signal

The Payment Initiation (UK and Europe), Transfer, and Signal products also have their own logs in the Plaid Dashboard. If you are enabled for these products, links to view activity will appear in the Products menu within the dashboard. Using these logs, you can view the status and history of payment and transfer attempts, returns, and other product-specific information. For Payment Initiation and Transfer, this information is also available via the API; for more details, see the API documentation for [Payment Initiation](/docs/api/products/payment-initiation/) and [Transfer](/docs/api/products/transfer/).

#### Audit logs

Audit logs of activity occurring in the Dashboard is available to admins on teams with Premium Support Packages. To learn more about upgrading to a Premium Support Package, contact your Account Manager.

Dashboard audit logs include user identity, action type, IP address, and timestamps for core Dashboard actions. Audit log support for actions on product-specific Dashboard pages is not yet available.

#### Troubleshooting institution status

The Plaid Dashboard contains an [Institution Status](https://dashboard.plaid.com/activity/status) pane, which allows you to view details and stats about institution connectivity over the past two weeks, as well as any recent downtime or special notes about the institution. You can also use this view to search institutions to see whether they are supported by Plaid and which Plaid products can be used with them. Most of this information is also available via the API; for more information on programmatically retrieving institution status, see [Institutions API endpoints](https://plaid.com/docs/api/institutions/).

##### Institution status details

![Plaid status for a bank: Auth 100%, Identity 100%, Transactions 100%. Historical graph of item add success. All products 100% success rate.](/assets/img/docs/institution-status-detail.png)

The graph at the top of the page shows institution status over the past two weeks, while the boxes below the graph show the current institution status. The time range used for calculating the current institution status may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more.

Plaid displays both Item add success rates (success adding new Items) and Item update success rates (success refreshing data for existing Items). The success rate as shown on the [main status page](https://dashboard.plaid.com/activity/status) is the Item add success rate.

All success rates reflect the percentage of successful attempts to retrieve data. Both Plaid errors and institution errors are combined when calculating the percentage of unsuccessful attempts. User errors, such as an Item add failing due to an incorrect password entered during Link, are not considered when calculating success rates.

All success rates are rounded down to the nearest percentage point. A success rate of 99% or higher indicates that the institution is considered healthy.

##### Institution migration status

Within the Institutions Status page, you can use the [migration pane](https://dashboard.plaid.com/activity/status) to view the status of OAuth migrations. The migration pane shows all institutions with planned or current migrations and allows you to drill in to see the migration timeline and the impact on existing Items.

![Plaid migration status for a bank, in progress, type: Classic to API, timeline displayed.](/assets/img/docs/migration-status-detail.png)

##### Insufficient data and delayed data

If the institution has not had enough Plaid traffic during the window being evaluated to produce meaningful health data, the "Insufficient Data" message will be displayed. This is most likely to affect small institutions. If this occurs, try using the [Item Debugger](https://dashboard.plaid.com/activity/debugger) instead to diagnose the problem.

If Plaid has not been able to update data from the institution for over two days, the institution status may appear as "delayed". If it has been at least two weeks since the last successful institution update, or if Plaid is aware of a problem in updating data that is likely to take over two weeks to resolve, the institution status will appear as "stopped".

##### Alerting

You can set up alerts to be notified of changes to the global Item add success rate of any institution. From [**Settings > Team Settings > Communications**](https://dashboard.plaid.com/settings/team/notification-preferences), open the "Status alerts" tab to create an alert. You can also create an alert directly from the institution's status page. If you choose webhook-based alerting, the webhook that will be sent is the [`INSTITUTION_STATUS_ALERT_TRIGGERED`](/docs/api/institutions/#institution_status_alert_triggered) webhook.

#### Troubleshooting with Item Debugger

To view the status of a specific Item or Link session in the dashboard, you can use the [Item Debugger](https://dashboard.plaid.com/activity/debugger). You can look up Items and Link sessions by `client_user_id`, `item_id`, `access_token`, `link_token`, or `link_session_id`. Troubleshooting information available includes error codes and suggested troubleshooting steps you can take to resolve any errors.

Access to the Item Debugger is also available via Plaid's Dashboard MCP server, allowing you to troubleshoot Item details via LLM chat or an agent. For more details, see the [blog post on the Plaid Dashboard MCP server](https://plaid.com/blog/plaid-mcp-ai-assistant-claude/) and the [Plaid Dashboard MCP server documentation](/docs/resources/mcp/).

![Item debugger showing a healthy item for institution Juno with an ID, transactions updated 5 hours ago.](/assets/img/docs/item-debugger.png)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
