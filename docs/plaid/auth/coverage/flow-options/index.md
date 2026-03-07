---
title: "Auth - Configuring entry points | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/flow-options/"
scraped_at: "2026-03-07T22:04:31+00:00"
---

# Configuring Auth entry points

#### Configure end users' options for verifying their account with Reroute to Credentials and Auth Type Select

#### Default manual verification flow entry point

When using the default flow, the **Link with account numbers** entry point will appear when the user does any of the following:

- Encounters an error in the Link flow
- Selects an institution whose connection health is poor
- Attempts to close the Link modal without connecting an institution
- Receives a "no results" message when searching for an institution
- Scrolls to the bottom of the institution search results

To trigger this flow via the [Link demo](https://plaid.com/link-demo) specify **Auth** in the Product dropdown menu and then **Launch Demo**. Trigger empty search results (type in a search query like ‘xyz’) or the Exit Pane (by closing Link) and select **Link with account numbers**.

This entry point will direct the user to Database Auth, Instant Micro-deposits, or Same-Day Micro-deposits, depending on which fallback options you have enabled and on whether the institution supports Instant ACH.

You can change these entry points using the features Auth Type Select and Reroute to Credentials.

Auth Type Select adds a prompt for the end user to pick a manual Auth flow upfront, without hitting any of the triggers above.

Reroute to Credentials is the opposite: it guides or forces the end user away from manual Auth flows if Plaid detects that the end user's institution is supported with a non-manual flow.

#### Adding manual verification entry points with Auth Type Select

![Plaid auth manual link flow, select bank link option: Instantly or Manually. Continue button visible at the bottom.](/assets/img/docs/auth/auth_type_select_manual.png)

Auth Type Select is a Link configuration that shows a pane upfront, enabling end users to actively choose between
credential-based authentication and manual account authentication at the start of the Plaid Link flow.

To demo **Auth Type Select** via the [Plaid Link demo](https://plaid.com/demo), specify **United States** and **Auth** in the Country and Product dropdown menu respectively. Toggle on Auth Type Select, then **Launch Demo**.

To enable Auth Type Select, use the [Dashboard Account Verification pane](https://dashboard.plaid.com/account-verification), or, if not using the Dashboard Account Verification pane, set `auth.auth_type_select_enabled: true` when calling [`/link/token/create`](/docs/api/link/#linktokencreate).

![Embedded institution search with Auth Type Select enabled, showing bank options and payment plan selection for PetGuard insurance.](/assets/img/docs/embedded-link-ats.gif)

Embedded Institution Search with Auth Type Select enabled

##### When to use Auth Type Select

Auth Type Select is best for customers who expect that a substantial portion of their consumers would prefer to authenticate manually, and who prefer to
maximize Auth coverage over other data products. Examples may include business account verification use cases where a high percentage of users are not expected to have
access to credentials of a bank account.

Customers that see the greatest success with the Auth Type Select configuration have a substantial portion of users (over half, as a rule
of thumb) who cannot or will not connect by logging into their bank, who have invested in fraud risk mitigations or have a low risk use
case, and/or have very high intent users (such as users connecting a bank account to receive a payment).

Some customers observe an increase in conversion of verified Auth data with this configuration. However, this configuration reduces coverage of other
products. When offered the manual option upfront, more users may choose to link manually. Users who opt to connect via micro-deposits can have lower conversion than Instant Auth because
the flow requires more steps for a user, including potentially returning to Link the next day to verify micro-deposits. It may encourage users to
connect via micro-deposits who would otherwise connect via credential-based Auth types (if the manual option was not available
upfront).

This configuration is best for customers who want to maximize Auth coverage and can tolerate reduced coverage of other products like
Balance and Identity.

#### Removing manual verification entry points with Reroute to Credentials

Reroute to Credentials will prompt users to use a login-based verification flow if they attempt to use a manual verification flow with an institution where login-based flows are supported. [Optional Reroute to Credentials](/docs/auth/coverage/flow-options/#optional-reroute) (the default Link flow) provides a nudge to use a login-based flow, but will not block manual verification, whereas [Forced Reroute to Credentials](/docs/auth/coverage/flow-options/#forced-reroute) will block manual verification if login-based flows are available. Alternatively, you can [disable Reroute to Credentials](/docs/auth/coverage/flow-options/#no-reroute) entirely.

![](/assets/img/docs/auth/reroute-overview.png)

Forced Reroute to Credentials example flow

In some cases, when rerouting a user to a supported institution, Plaid may show a list of institutions instead of a single institution. This can occur when a financial institution has multiple different brands that share a common routing number.

##### Optional Reroute

Optional Reroute to Credentials is the default flow. It is best for customers who would like to make a recommendation for the user to connect via credentials, but not block the user from proceeding with the Same Day Micro-deposits flow. This flow will provide a reminder to the user that the institution is supported, and detail the benefits of instant connection, while also leaving the manual verification option in place. To enable Optional Reroute, you do not need to take any action.

If you have disabled Optional Reroute (by enabling [Forced Reroute](/docs/auth/coverage/flow-options/#forced-reroute) or [No Reroute](/docs/auth/coverage/flow-options/#no-reroute)), you can re-enable it by using the [Dashboard Account Verification pane](https://dashboard.plaid.com/account-verification), or set `auth.reroute_to_credentials: "OPTIONAL"` in the [`/link/token/create`](/docs/api/link/#linktokencreate) call.

| Single Institution Optional Reroute | Multiple Institution Optional Reroute |
| --- | --- |
| Sign in to Gingham Bank via Plaid using routing number 12341234. Tap 'Continue' or 'Enter account number instead.' | Mobile screen with options to select a financial institution. Choices: Gingham Bank, Herringbone Treasury. Option to enter account number. |

##### Forced Reroute

Forced Reroute to Credentials is best for customers who would like to restrict the user to connect via credentials if the institution is supported on a credential-based flow
and stop the user from proceeding with a manual flow. This flow is designed to block a user from connecting manually for institutions that are supported
via credentials in high ACH risk use cases. To enable Forced Reroute, use the [Dashboard Account Verification pane](https://dashboard.plaid.com/account-verification), or set `auth.reroute_to_credentials: "FORCED"` in the [`/link/token/create`](/docs/api/link/#linktokencreate) call.

| Single Institution Forced Reroute | Multiple Institution Forced Reroute |
| --- | --- |
| Sign in to Gingham Bank using Wonderwallet with routing number 12341234. Press 'Continue' to securely connect your account. | Select an institution based on routing number 12345678. Options: Gingham Bank, Herringbone Treasury. Sign in required. |

##### No Reroute

To completely disable Reroute to Credentials, use the [Dashboard Account Verification pane](https://dashboard.plaid.com/account-verification), or set `auth.reroute_to_credentials: "OFF"` in the [`/link/token/create`](/docs/api/link/#linktokencreate) call.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
