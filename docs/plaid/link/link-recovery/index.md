---
title: "Link - Link Recovery (beta) | Plaid Docs"
source_url: "https://plaid.com/docs/link/link-recovery/"
scraped_at: "2026-03-07T22:05:06+00:00"
---

# Link Recovery (beta)

#### Notify users impacted by institution downtime to return to your app and link their accounts

#### Overview

Link Recovery (beta) allows you to provide a superior user experience and improved conversion when users run into issues connecting with an institution. When a user attempts to connect an institution that is experiencing a temporary connectivity outage, Link Recovery lets the user opt in to receiving an email notification from Plaid letting them know when the issue has been resolved. The user can then launch Link directly from a link in the email notification and complete the process of connecting their account.

Link Recovery is designed to be easy to implement and can be added to an existing Plaid integration with only a few lines of code.

The user encounters an error because their institution is down.

![The user encounters an error because their institution is down.](/assets/img/docs/link-recovery/lr-1.png)

![They elect to be notified when the problem is resolved.](/assets/img/docs/link-recovery/lr-2.png)

![They exit Link...](/assets/img/docs/link-recovery/lr-3.png)

![...and receive a confirmation email.](/assets/img/docs/link-recovery/lr-4.png)

![When the institution is back up, they are notified...](/assets/img/docs/link-recovery/return-1.png)

![...and can click a link in the email to go directly to Link.](/assets/img/docs/link-recovery/return-2.png)

![They connect to the institution...](/assets/img/docs/link-recovery/return-3.png)

![...select their account...](/assets/img/docs/link-recovery/return-4.png)

![...and receive a success screen with a button to redirect them to your app.](/assets/img/docs/link-recovery/return-5.png)

Link Recovery is currently in beta. To request access, [complete this short Google form](https://docs.google.com/forms/d/e/1FAIpQLSfk98Di4SajVIG0xw5gLhEAKyspNBl8fvUrg32QJ7WnIvb9Cw/viewform).

#### Integration steps

##### Calling /link/token/create

When calling [`/link/token/create`](/docs/api/link/#linktokencreate), specify the following, in addition to your normal parameters:

- A `webhook` URI, where you will receive the `SESSION_FINISHED` webhook used to deliver the `public_token` for Link Recovery sessions.
- (Optional) A `hosted_link.completion_redirect_uri`, indicating the URI that the user should be redirected to after completing a Link Recovery session. This field is optional, but strongly recommended; you can use it to redirect the user to your app so they can see the impact of linking their account.
- If you are not specifying a `hosted_link.completion_redirect_uri`, include a `hosted_link` object (an empty object is fine) in the request in order to enable Hosted Link for your session. Link Recovery requires Hosted Link.

##### Listening for the SESSION\_FINISHED webhook

Next, make sure to listen for the `SESSION_FINISHED` webhook, which will fire when your end user completes their Link Recovery session.

Sample SESSION\_FINISHED webhook

```
{
  "webhook_type": "LINK",
  "webhook_code": "SESSION_FINISHED",
  "status": "SUCCESS",
  "link_session_id": "356dbb28-7f98-44d1-8e6d-0cec580f3171",
  "link_token": "link-sandbox-af1a0311-da53-4636-b754-dd15cc058176",
  "public_tokens": [
    "public-sandbox-b0e2c4ee-a763-4df5-bfe9-46a46bce993d"
  ],
  "environment": "sandbox"
}
```

After retrieving the public token from the `public_tokens` array, you can exchange it as normal.

#### Testing Link Recovery

Plaid will not send email notifications in the Sandbox environment, so the Link Recovery flow cannot be tested end-to-end in Sandbox. Instead, you can test the two parts of the flow separately.

To test that you are enabled for Link Recovery and that users can sign up to receive emails, log in to Link in Sandbox with the following credentials to simulate a temporary institution outage: username: `user_good` password: `error_INSTITUTION_DOWN`. (If you are not yet enabled for Link Recovery, you can also use the following credentials to see the Link Recovery experience in Sandbox: username: `user_link_recovery` password: `{}`.)

When prompted for an email address in the Sandbox Link Recovery flow, enter `example@plaid.com` as the email and `123456` as the security code. Note that in the Sandbox environment, no Link Recovery emails will actually be sent.

To test that the Link Recovery session works correctly, call [`/link/token/create`](/docs/api/link/#linktokencreate) and manually go to the `hosted_link_uri` present in the response. This will launch a Hosted Link session, where you can test going through the Link flow with the `user_good` / `pass_good` credentials and triggering the `SESSION_FINISHED` webhook.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
