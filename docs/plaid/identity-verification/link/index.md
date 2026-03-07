---
title: "Identity Verification - Link callbacks | Plaid Docs"
source_url: "https://plaid.com/docs/identity-verification/link/"
scraped_at: "2026-03-07T22:04:55+00:00"
---

# Link callbacks

#### Guide to using Link callbacks in your integration

After integrating Identity Verification using Link, you can use Link’s client-side callbacks to see how your users are progressing through the flow. Note that you will not be able to access these if you use the Plaid-hosted flow with shareable links.

[`onSuccess`](/docs/link/web/#link-web-create-onSuccess) will fire when a user has successfully completed the Link flow. The `link_session_id` returned by `onSuccess` can be passed to [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget) in the `identity_verification_id` field to retrieve data on the verification. Note that receiving `onSuccess` simply means that the user completed the Identity Verification flow, not that their identity was verified. A completed session in which a user fails Identity Verification will still result in `onSuccess`.

[`onExit`](/docs/link/web/#link-web-create-onExit) will fire when a user exits a session without finishing all required steps. If you receive this event, you may want to provide an option for the user to restart the Link flow, and/or provide guidance on manually performing KYC.

[`onEvent`](/docs/link/web/#link-web-create-onEvent) will fire at certain points throughout the IDV flow with updates on the user’s progress at each point. For each step, Plaid will send an [`IDENTITY_VERIFICATION_START_STEP`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-START-STEP) event as well as a [`IDENTITY_VERIFICATION_PASS_STEP`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-PASS-STEP) or [`IDENTITY_VERIFICATION_FAIL_STEP`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-FAIL-STEP) event based on the outcome of the step. Each step will also have an associated `view_name`. The possible view names include:
[`ACCEPT_TOS`](/docs/link/web/#link-web-onevent-metadata-view-name-ACCEPT-TOS), [`VERIFY_SMS`](/docs/link/web/#link-web-onevent-metadata-view-name-VERIFY-SMS),[`KYC_CHECK`](/docs/link/web/#link-web-onevent-metadata-view-name-KYC-CHECK), [`DOCUMENTARY_VERIFICATION`](/docs/link/web/#link-web-onevent-metadata-view-name-DOCUMENTARY-VERIFICATION), [`RISK_CHECK`](/docs/link/web/#link-web-onevent-metadata-view-name-RISK-CHECK), and [`SELFIE_CHECK`](/docs/link/web/#link-web-onevent-metadata-view-name-SELFIE-CHECK).

For each session, Plaid will also inform you of the outcome of the session by sending an [`IDENTITY_VERIFICATION_PASS_SESSION`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-PASS-SESSION) or [`IDENTITY_VERIFICATION_FAIL_SESSION`](/docs/link/web/#link-web-onevent-eventName-IDENTITY-VERIFICATION-FAIL-SESSION) event based on the outcome of the session.

Note that, aside from `IDENTITY_VERIFICATION_PASS_SESSION` and `IDENTITY_VERIFICATION_FAIL_SESSION`, the sequence of view names and `onEvent` events are subject to change without notice, as Plaid adds additional functionality and makes improvements to flows. You should treat view names as informational and not rely on them for critical business logic.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
