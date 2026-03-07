---
title: "API - New User API migration | Plaid Docs"
source_url: "https://plaid.com/docs/api/users/user-apis/"
scraped_at: "2026-03-07T22:04:28+00:00"
---

# User APIs

#### Information on new user APIs

Plaid is updating our APIs to support the next generation of user-based products - such as Plaid Protect - and to create a more unified and consistent experience across our platform. These updates improve multi-product compatibility, simplify debugging, and ensure user identifiers behave consistently across all Plaid products. If you're beginning a new Plaid Check Consumer Report (CRA) or Multi-Item Link integration in December 2025 or later, you’ll use these updated APIs to build your integration.

If you are an existing customer using Plaid Check, Plaid Income Verification, or Multi-Item Link as of December 10, 2025, here's what you need to know:

- **No action is required.**
- **Your existing integration remains fully supported.** Plaid is not removing support and your integration will continue to function as expected.
- **We’ll share optional migration steps in Q1 2026.** Only after that point will you be able to migrate.

=\*=\*=\*=

#### What's new

Customers with existing user token integrations cannot migrate to this new flow until April 1, 2026. The API changes below are currently described for informational purposes only.

##### Updates to user creation and identification

- When calling [`/user/create`](/docs/api/users/#usercreate), the response includes a single `user_id` instead of a `user_token` and a `user_id`. This `user_id` is used instead of the `user_token` to identify the user throughout the Plaid API, including when calling API endpoints or when receiving webhooks.
  - A `user_id` created on the new API (prefixed with `usr_`) is not equivalent to a `user_id` (not prefixed with `usr_`) created on the old API. If you have not yet migrated to the updated user APIs, you cannot use a `user_id` in place of a `user_token` for endpoints that accept either identifier.
- [`/user/create`](/docs/api/users/#usercreate) is now idempotent. In the old flow, when [`/user/create`](/docs/api/users/#usercreate) was called on a `client_user_id` more than once, it would return an error; it now returns the same `user_id` as the original call.
- The user schema has an `identity` object (instead of the `consumer_report_user_identity` object), which is used in the [`/user/create`](/docs/api/users/#usercreate) and [`/user/update`](/docs/api/users/#userupdate) request bodies. This `identity` object has a different schema than the `consumer_report_user_identity` object.

##### Changes to user management

- In the old flow, a `client_user_id` could never be re-used to create a new user, even if the user token was deleted with [`/user/remove`](/docs/api/users/#userremove). In the new flow, once [`/user/remove`](/docs/api/users/#userremove) has been called on a `user_id`, a new user can be created for the same `client_user_id` by calling [`/user/create`](/docs/api/users/#usercreate).
- The endpoint [`/user/get`](/docs/api/users/#userget) has been added, allowing you to retrieve identity details about a user that you have previously created.
- Coming soon, the user APIs will include net-new endpoints and additional functionality to simplify user management.

##### Other changes

- The webhooks `CHECK_REPORT_READY` and `CHECK_REPORT_FAILED` have been renamed to `USER_CHECK_REPORT_READY` and `USER_CHECK_REPORT_FAILED`.
- For Cash Flow Insights (beta) customers, the different Insights webhooks have been replaced by a single webhook, `CASH_FLOW_INSIGHTS_UPDATED`, with an `insights` payload field listing all of the insights received.

When optional migration begins on April 1, 2026, existing customers who currently receive the older version of webhooks will begin to receive both the new and old sets of webhooks, in order to allow for migration.

=\*=\*=\*=

#### Who gets the new user APIs

As of December 10, 2025, all Plaid customers will experience the new user API behavior by default, with the following exceptions:

- Any existing Plaid customers who ever used the [`/user/create`](/docs/api/users/#usercreate) endpoint in either Sandbox or Production as of December 10, 2025, will automatically be kept on the old user API behavior, to avoid breaking changes. This group includes all existing and currently integrating customers of Consumer Report, Multi-Item Link, and/or Income Verification.
- After December 10, 2025, any new customers of Plaid Income Verification will need to contact Sales, Support, or their Account Manager to request access to the old user APIs, since the new user APIs will not be ready for Plaid Income Verification until late Q1 2026. Note that this applies only to the legacy [Plaid Income Verification](/docs/income/) product; it does not apply to the Plaid Check Consumer Report Income modules, such as Base Report and Income Insights.

If you aren't sure whether you have the new or old API, call [`/user/create`](/docs/api/users/#usercreate).

- In the new API, the response will not include a `user_token`, and your `user_id` will be formatted with the prefix `usr_`.
- In the old API, the response will include a `user_token`, and the `user_id` will not contain a prefix.

=\*=\*=\*=

#### Client library version requirements

To use the new user APIs with a Plaid client library, the minimum client library versions are:

- Python: 38.0.0
- Go: 41.0.0
- Java: 39.0.0
- Node: 41.0.0
- Ruby: 45.0.0

=\*=\*=\*=

#### Summary

**New clients integrating with Plaid Check or Multi-Item Link** beginning December 10, 2025 or later should use the new `user_id` based implementation currently described in the docs.

**Existing users of other Plaid products** who are integrating with Plaid Check or Multi-Item Link for the first time beginning December 10, 2025 or later should use the new `user_id` based implementation currently described in the docs. They may also need to [update their client library versions](/docs/api/users/user-apis/#client-library-version-requirements).

**New clients using Plaid's legacy (non-CRA) Bank Income product** for the first time beginning December 10, 2025 or later should contact their Account Manager or file a support ticket via the Dashboard to request access to the `user_token` field.

**Existing clients already using Plaid Check or Multi-Item Link products** should take no action at this time. Migration will be optional (though recommended) and available beginning April 1, 2026. At this time, these existing customers will also begin to receive the new renamed webhooks in addition to their existing webhooks, in order to allow for migration.

More information will be provided about the new user APIs in Q1 2026. In the meantime, if you have questions about timing, readiness, or how Plaid's new user APIs might benefit your integration, contact your Plaid Account Manager.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
