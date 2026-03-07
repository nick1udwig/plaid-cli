---
title: "KYC/AML and anti-fraud | Plaid Docs"
source_url: "https://plaid.com/docs/kyc-aml/"
scraped_at: "2026-03-07T22:05:00+00:00"
---

# KYC, AML, and anti-fraud products

#### Review and compare solutions

This page provides overviews of Plaid's KYC, AML, and anti-fraud products to help you find the right one for your needs.

Plaid offers Identity Verification for KYC (Know Your Customer) compliance, Monitor for AML (Anti Money Laundering) compliance, and Beacon (beta) for anti-fraud. The three products are tightly integrated and managed through a single dashboard but can also be used separately.

Plaid Identity is also an anti-fraud solution targeted specifically at account takeover fraud. Because it's designed for payments use cases, you can learn more about it on the [Payments products](/docs/payments/) page.

#### Identity Verification

[Identity Verification](/docs/identity-verification/) is Plaid's KYC solution. Via an interactive session with the user, Identify Verification allows you to check user-provided identity information such as name, ID number, phone number, and address against high-trust identity databases; check identity documents for expiration, signs of fraud, or mismatch with other user-provided data; run selfie verification to confirm liveness and photo ID matching; verify a user's phone number via SMS; and analyze a user's session, behavior, and identity details for signs of fraud. Identity Verification can also be run as a fully background session without interactivity (for checking pre-collected user data against databases only). Identity Verification can plug directly into Monitor for automatic AML watchlist scanning.

Identity Verification is a separate product from Identity / Identity Match, but can integrate directly with [Identity Match](/docs/identity/#identity-match) to reduce account takeover fraud risk when a user has linked a financial account to your application. Identity Match will verify that the ownership details (such as name, address, and phone number) on the linked account match the data verified via Identity Verification.

#### Monitor

[Monitor](/docs/monitor/) provides AML capabilities, screening users against PEP (Politically Exposed Person) and sanction lists. Monitor is frequently used with Identity Verification -- integrating Monitor into an existing Identity Verification deployment can be as simple as checking a box on your Identity Verification template -- but can also be deployed separately. Any watchlist hits will be exposed via both a user-friendly dashboard UI and via API, for use with either manual or automated review workflows.

#### Monitor and Identity Verification comparison

|  | Identity Verification | Monitor |
| --- | --- | --- |
| Summary | Flexible and configurable KYC compliance to verify identity and detect fraud | Scan users against watchlists for AML compliance |
| Supported countries | 190+ countries | 190+ countries |
| UI languages | English, French, Spanish, Japanese, Portuguese | N/A |
| Billing plans available | Pay-as-you-go or 12-month contract | Pay-as-you-go or 12-month contract |

##### Beacon

Beacon is Plaid's free network for fraud detection. You can scan all new users with Beacon to see whether they have any known fraud alerts from the Beacon Network. You can also be alerted when a new incident of fraud is reported on a user you have already added. These capabilities are available free of charge to all Plaid customers.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
