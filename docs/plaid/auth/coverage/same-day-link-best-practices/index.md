---
title: "Auth - Anti-fraud best practices | Plaid Docs"
source_url: "https://plaid.com/docs/auth/coverage/same-day-link-best-practices/"
scraped_at: "2026-03-07T22:04:32+00:00"
---

# Anti-fraud best practices

#### Optimally configure manual entry verification flows to reduce fraud

Plaid provides a suite of fraud prevention products that assist your application in catching bad actors and ACH returns. You can verify the source
of funds with [Identity](/docs/identity/), confirm the real-time [Balance](/docs/balance/) prior to a transfer, and leverage Plaid's ML-based [Signal Transaction Scores](/docs/signal/) to prevent returns and release funds earlier. These features are fully compatible with accounts connected via Instant Auth, Instant Match, and Automated Micro-deposits.

However, if an account is connected via a manual entry method such as Same Day Micro-deposits, Instant Micro-deposits, or Database Auth, these features are not always available, which could increase the likelihood that you experience fraud and ACH returns. Following the recommendations below can help mitigate these risks.

##### Use Identity Match, Signal, and/or Identity Document Upload

Approximately 30% of Items created with manual entry methods (including 100% of Items created by Database Auth that have a `database_insights_pass` verification status) are supported by [`/identity/match`](/docs/api/products/identity/#identitymatch), which allows you to determine the likelihood that the user's identity details, such as name and address, on file with their financial institution match identity information held by you. For more details on this feature, see [Identity](/docs/identity/#identity-match).

The same Items that are supported by [`/identity/match`](/docs/api/products/identity/#identitymatch) are also supported by Signal Transaction Scores. Signal Transaction Scores can assess the return risk of a transaction based on machine learning analysis and alert you to high-risk transactions. For more details, see [Signal Transaction Scores](/docs/signal/).

[Identity Document Upload](/docs/identity/identity-document-upload/) verifies account owner identity based on bank statements, and is compatible with Items that don't support [`/identity/match`](/docs/api/products/identity/#identitymatch) or [`/identity/get`](/docs/api/products/identity/#identityget). After creating the Item, you can use update mode to send the user through a Link session where they will be asked to upload a bank statement. Plaid will then analyze this statement for an account number match and will parse identity data from the statement. With the optional Fraud Risk feature, you can also check the uploaded statement for signs of fraud. For more details, see [Identity Document Upload](/docs/identity/identity-document-upload/).

##### Adjust a user’s Link experience based on their risk profile

In order to reduce fraud upstream on your application, you can use [Plaid Identity Verification](/docs/identity-verification/) to verify a
government ID or match with a selfie of the document holder.

If your application does not have an identity verification solution or Plaid Link is not gated from the general public with fraud prevention and user verification checks in place, we do not recommend adopting manual entry verification flows as it may introduce an unnecessary fraud vector onto your platform.

If you identify a user to be riskier, consider disabling manual entry verification flows for those users.

Another option for riskier users is to leave manual entry verification flows enabled, but enable [Reroute to Credentials in Forced mode](/docs/auth/coverage/flow-options/#forced-reroute), which will only allow
the user to link via manual entry verification flows when using a routing number not supported by other authentication methods.

You may also consider changing your user’s experience with your service based on their connection method. For example, if a user connected via a manual entry verification flow, you may consider enforcing a lower transfer threshold than for users where it was possible to verify identity and increasing hold times
on those funds.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
