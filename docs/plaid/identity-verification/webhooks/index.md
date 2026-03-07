---
title: "Identity Verification - Webhooks | Plaid Docs"
source_url: "https://plaid.com/docs/identity-verification/webhooks/"
scraped_at: "2026-03-07T22:04:56+00:00"
---

# Webhooks

#### Subscribe to events and get real-time alerts for changes

Once you have Link with Identity Verification embedded in your app, you should be able to complete an entire session and review it in the dashboard.

To complete your integration, you need to add a webhook receiver endpoint to your application.

To add a webhook, visit the [dashboard webhook configuration page](https://dashboard.plaid.com/developers/webhooks) and click **New Webhook**. Note that this is the only setting used to configure Identity Verification webhook receivers; webhook URLs set via [`/link/token/create`](/docs/api/link/#linktokencreate) will not be used.

You can select which events you want to subscribe to. For Identity Verification, there are three events:

- [`STEP_UPDATED`](https://plaid.com/docs/api/products/identity-verification/#step_updated)
- [`STATUS_UPDATED`](https://plaid.com/docs/api/products/identity-verification/#status_updated)
- [`RETRIED`](https://plaid.com/docs/api/products/identity-verification/#retried)

Enter your webhook receiver endpoint for the webhook you wish to subscribe to and click **Save**. Plaid will now send an HTTP POST request to the webhook receiver endpoint every time the event occurs, in both the Sandbox and Production environments. If multiple webhook receiver endpoints are configured for an Identity Verification event, webhooks will be sent to all the configured endpoints.

#### Event ordering

Identity Verification does not guarantee that webhooks will be delivered in any particular order. For example, while the logical ordering of webhooks for a Identity Verification session might look like this:

1. `STEP_UPDATED` The user has started the Identity Verification session and is on the first step
2. `STEP_UPDATED`
3. `STATUS_UPDATED` The user has reached a terminal state for their session
4. `RETRIED` A retry has been requested for this user, either via the dashboard or via API
5. `STEP_UPDATED`
6. `STEP_UPDATED`
7. `STATUS_UPDATED` The retry has been completed

you should be prepared to handle these events in any delivery order. For example, consider whether your application will properly handle:

- A `STEP_UPDATED` event being delivered after a `STATUS_UPDATED` event.
- A `STEP_UPDATED` event being delivered before an associated `RETRIED`

In order to properly handle webhook events being delivered out of order, your application should lookup the user's associated session(s) via the [`/identity_verification/list`](/docs/api/products/identity-verification/#identity_verificationlist) API.

#### Webhook Reference

For full webhook information, refer to the [API Documentation](/docs/api/products/identity-verification/).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
