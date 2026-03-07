---
title: "Account - Security | Plaid Docs"
source_url: "https://plaid.com/docs/account/security/"
scraped_at: "2026-03-07T22:03:45+00:00"
---

# Plaid Dashboard account security and account management

#### Learn how to protect and manage your Plaid Dashboard account

#### Password resets

To reset your password, log in to the Plaid Dashboard, then access the [Personal Security page](https://dashboard.plaid.com/personal/security), which contains a password reset screen.

If you ever forget your password to your Plaid Dashboard account, use the **Password reset** button on the login page to reset your password. You will only be able to receive your password reset email if you have verified your Plaid email address. If your email address is unverified, search your email for a message from [no-reply@plaid.com](mailto:no-reply@plaid.com) with the subject line "Confirm your email" and click on the link in that email to verify your address. You should then be able to reset your password.

#### Two-factor authentication

Plaid allows you to enable two-factor authentication for your dashboard account via either text message or app-based authentication. After you enable two-factor authentication, you’ll be prompted to enter an authentication code each time you sign in. When you enable two-factor authentication, Plaid will show you a backup recovery code. Be sure to write down the recovery code and store it in a safe place, since it will only be displayed once.

To check whether your team members have enabled two-factor authentication, you can visit the [Team Members page](https://dashboard.plaid.com/settings/team/members), where your team members' status will be displayed as "2FA ON", "2FA OFF", or "PENDING".

#### SSO login

Plaid offers SSO (single sign-on) login to the Plaid Dashboard for all customers, except those on Pay-as-you-go-plans. We support all major identity providers, including Okta, Google Workspaces, OneLogin, and over 30 others. You can enable your account for SSO via the Dashboard, under [Settings > Team Settings > SSO](https://dashboard.plaid.com/settings/team/sso). There is no additional charge for SSO enablement.

SCIM support is available to customers on Premium Support packages. To learn more, contact your Account Manager.

#### Rotating keys

If your `secret` is ever compromised, you should rotate it from the [Keys](https://dashboard.plaid.com/developers/keys) section of the Dashboard. Clicking the **Rotate secret** button will generate a new secret, which you can then copy and test. The old secret will still remain active until you delete it by clicking the **Delete** button.

#### Managing notification settings

To control which emails you receive from Plaid, you can update your [notification settings](https://dashboard.plaid.com/settings/team/notification-preferences) in the Dashboard.

#### Deleting accounts

To leave or delete a team, see [Leaving or deleting a team](/docs/account/teams/#leaving-or-deleting-a-team).

To delete your Plaid Dashboard login itself, [contact Support](https://dashboard.plaid.com/support/new/product-and-development/account-administration/delete-account). Note that deleting your login is irreversible.

A Plaid Dashboard account is different from a Plaid account at my.plaid.com. To manage the financial information you're sharing via Plaid, visit [my.plaid.com](https://my.plaid.com/).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
