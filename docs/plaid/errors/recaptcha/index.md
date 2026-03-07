---
title: "Errors - Recaptcha errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/recaptcha/"
scraped_at: "2026-03-07T22:04:52+00:00"
---

# ReCAPTCHA Errors

#### Guide to troubleshooting reCAPTCHA errors

#### **RECAPTCHA\_REQUIRED**

##### The request was flagged by Plaid's fraud system, and requires additional verification to ensure they are not a bot.

##### Link user experience

Your user will be prompted to solve a [Google reCAPTCHA](https://www.google.com/recaptcha/intro/v3.html) challenge in the Link `Recaptcha` pane.
If they solve the challenge successfully, the user's request is resubmitted and they are directed to the next Item creation step.

##### Common causes

- Plaid's fraud system detects abusive traffic and considers a variety of parameters throughout Item creation requests.
  When a request is considered risky or possibly fraudulent, Link presents a reCAPTCHA for the user to solve.

##### Troubleshooting steps

Link will automatically guide your user through reCAPTCHA verification. As a general rule, we recommend instrumenting basic fraud monitoring to detect and protect your website from spam and abuse.

If your user cannot verify their session, please submit a [Support](https://dashboard.plaid.com/support/new) ticket with the following identifiers: `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "RECAPTCHA_ERROR",
 "error_code": "RECAPTCHA_REQUIRED",
 "error_message": "This request requires additional verification. Please resubmit the request after completing the challenge",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **RECAPTCHA\_BAD**

##### The provided challenge response was denied.

##### Sample user-facing error message

Please try connecting a different account: There was a problem processing your request. Your account could not be connected at this time

##### Link user experience

If after several failed attempts your user is still unable to solve the reCAPTCHA challenge,
the `RECAPTCHA_BAD` error is returned and they are directed to the `InstitutionSelect` pane to connect a new account.

##### Common causes

- The user was unable to successfully solve the presented reCAPTCHA problem after several attempts.
- The current session is a bot or other malicious software.

##### Troubleshooting steps

Verify your user's session -- reCAPTCHA is built to allow valid users to pass through with ease. If a user was unable to solve the challenge, they may be a bad actor.

If your user cannot verify their session, please submit a [Support](https://dashboard.plaid.com/support/new) ticket with the following identifiers: `link_session_id` or `request_id`.

API error response

```
http code 400
{
 "error_type": "RECAPTCHA_ERROR",
 "error_code": "RECAPTCHA_BAD",
 "error_message": "The provided challenge response was denied. Please try again",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
