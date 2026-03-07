---
title: "Account - Teams | Plaid Docs"
source_url: "https://plaid.com/docs/account/teams/"
scraped_at: "2026-03-07T22:03:45+00:00"
---

# Plaid teams

#### Learn about Plaid teams

Your Plaid Dashboard account can create or join multiple teams to enable collaboration across different teams or companies.

You can be a member of multiple Plaid teams, each with their own API keys and billing information.

#### Creating a team

At the top-left of your Plaid Dashboard, click the team name drop-down and select **Create new team**.

![Dropdown displaying team options: WonderWallet, WonderWallet Team 2, WizardCash, and option to create a new team.](/assets/img/docs/team-menu.png)

Once you have submitted your team name and purpose into the form, click **Create team**. Your new team will come with a fresh set of API keys which can be used to build a new Plaid integration.

#### Managing your team membership

You can create a new team, view your existing teams, and leave a team from the [Personal menu](https://dashboard.plaid.com/personal/teams), found at the bottom-left of your Plaid Dashboard:

![Personal menu with options: Profile, Security, Teams, Logout for Plaid User; dropdown indicator present.](/assets/img/docs/personal-menu.png)

#### Team-specific settings

Visit [Settings > Team Settings](https://dashboard.plaid.com/settings/company/profile) in your Plaid Dashboard to view your company information, team members, billing settings, and invoices:

![Company and team settings menu with options for profile, branding, members, products, communication, link migrations, and compliance.](/assets/img/docs/team-settings.png)

#### Team domains

In addition to manually adding and removing team members, which you can do from [Team settings](https://dashboard.plaid.com/settings/company/profile), you can use the [Domains](https://dashboard.plaid.com/settings/team/domains) feature to automatically add new Plaid users to your team when they sign up with a verified email domain. This simplifies team management by preventing you from having to manually add members and also ensures that your co-workers join the correct Plaid team.

To associate a domain with your team, go to the [Domains](https://dashboard.plaid.com/settings/team/domains) tab and go through the WorkOS flow to verify your domain. You must own and control your own domain name to use this feature; public email domains such as gmail.com are not eligible for domain verification.

![](/assets/img/docs/domain-capture.png)

Once domain verification is complete, which can take up to 48 hours, you can enable Domain Capture for that team by returning to the [Domains](https://dashboard.plaid.com/settings/team/domains) tab and clicking the "Enable" button for that team.

When Domain Capture is enabled, Dashboard users who sign up with an email address matching your team's verified domain will be prompted to join after verifying their email.

![](/assets/img/docs/domain-capture-user-flow.png)

You will receive an email notification whenever a user joins your team in this way. New members will be added with the "Default" permissions set.

To accommodate different organizational structures, multiple teams may be associated with a single domain, and a single team may be associated with multiple domains.
If a single domain is associated with multiple different teams, and Domain Capture is enabled for more than one of those teams, the user will be shown all eligible teams and prompted to choose which team(s) to join.

Enabling Domain Capture is a reversible operation; you may disable Domain Capture at any time from the Domains tab.

Managing Domain Verification and Domain Capture requires the Domain Verification permission, which is separate from the Team Management permission.

#### Managing team member permissions

From the Team Settings menu, you can also add and remove team members, as long as you have the Team management permission. You can assign specific permissions to each team member that you add to your Plaid team.

Members with the **Admin** permission are automatically granted all permissions to the entire Plaid Dashboard.

Members with the **Team Management** permission have similar access to the Admin permission – they can see and manage almost everything. However, they can't view or change API keys or view and file Support tickets. Some of what this role can do:

- Make production / product requests
- Change API configurations (OAuth redirect URIs, change pinned API version)
- Configure webhooks
- Configure team’s display settings and name
- Manage billing / invoices
- Manage team membership
- Manage OAuth migrations
- Delete the team

Members with the **Support** permission have access to the Support pages for viewing and filing support tickets.

Members with the **Link Customization Write Access** permission can create or edit Link customizations.

Certain pages are accessible to all Dashboard team members, regardless of their permissions. These pages include:

- Overview
- Status
- Logs
- Usage
- Analytics

You can also use the Team Settings page to control which environments team members have access to. In order to view the team's `client_id`, team members must have access to at least one environment.

##### Identity Verification, Monitor, and Beacon permissions

Members with Admin or Team Management permissions are able to assign granular permissions for Plaid Identity Verification, Monitor, and Beacon. There are numerous permission settings, which allow you to define who should be able to view certain information (e.g. end user PII) and take various actions. These permissions can be assigned while adding a team member or editing permissions for existing team members in the [Team Settings](https://dashboard.plaid.com/settings/team/members) page. Note that 2FA will need to be enabled for a user to be able to access this data.

##### Transfer Dashboard permissions

See the [Transfer Dashboard documentation](/docs/transfer/dashboard/#dashboard-permissions) for details on Transfer-specific Dashboard permissions.

##### Credit Dashboard permissions

- **Credit Dashboard All Permissions** - Grants access to all Credit Dashboard features.
  - **Credit Dashboard Access** - Grants access to view Credit Dashboard
  - **Credit Dashboard Application Creation** - Grants ability to create hosted link URLs
  - **Credit Dashboard User Profile Access** - Grants access to view Credit Dashboard user profiles
  - **Credit Dashboard View PII** - Allows viewing of PII on Credit Dashboard

#### Consolidating multiple accounts using teams

Teams prevent you from having to maintain multiple developer accounts. However, if you have already created multiple accounts, you can still consolidate them into a single account without having to create new keys or re-apply for access. To do so, follow these steps:

1. Select an account to be the new primary account that will contain all the teams.
2. From each account that you will be consolidating into the new primary account, invite the new primary account to the team.
3. From each account that you will be consolidating into the new primary account, create a new empty team.
4. From the new primary account, accept the invitations.
5. From each account that you will be consolidating into the new primary account, leave the old team.

Your teams will now be consolidated under the new account.

#### Leaving or deleting a team

To leave or delete a team, click on your name in the lower-left corner of the Dashboard and select **Teams** to access the [Teams page](https://dashboard.plaid.com/personal/teams). Click on the **...** button on the right side of the team's listing. Select **Leave Team** or **Delete Team**. Deleting the team will remove access to the team for all users, and the API keys for team will be deactivated. You must have the Team Management permission to delete a team. You may only delete a team via the Dashboard if you are on a Pay-as-you-go plan. If you have a dedicated Plaid Account Manager, contact your Account Manager instead to request team deletion.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
