---
title: "Identity Verification - Risk checks | Plaid Docs"
source_url: "https://plaid.com/docs/identity-verification/risk-checks/"
scraped_at: "2026-03-07T22:04:55+00:00"
---

# Risk checks

#### Understand what each Risk Check means and factors that influence results

This page outlines the various checks performed in the Risk Check category to help you understand the results and what Identity Verification evaluates during each check.

For each session's risk check results, the Identity Verification Dashboard will show the largest factors contributing to the risk.

#### Email Risk

Identity Verification does not collect an email address during the Link flow. Email risk will only be assessed if you collect the end user's email address and provide it to Plaid via [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) or [`/link/token/create`](/docs/api/link/#linktokencreate). Only verified email addresses should be sent for email risk assessment.

Attributes that increase email risk include:

- Email provided via a disposable email service, such as Mailinator. This is a strong fraud signal.
- Email associated with a recently-registered (<3 months old) domain.
- Email domain not deliverable. We do a live check on the domain to see if it is configured to receive email. Failing this test means the email is fake, which is a strong fraud signal.
- Email not present on breach lists or only present on newer breaches (presence on older lists indicates an older, established, actively-used email address)
- Email registered with few or no popular services, such as social media platforms. We check over 90 services for connections. In the Dashboard, all of the services that the email is detected as associated with will be displayed.

#### Device risk

Plaid looks at multiple attributes of the end user's device. Attributes that can increase device risk include:

| Risk factor | Contribution to device risk |
| --- | --- |
| Proxy usage | Moderate |
| VPN usage | Moderate |
| Tor usage | High |
| IP address matching known-malicious IP lists | High |
| Datacenter IP address (correlated with abuse) | High |
| IP address geolocation mismatch vs KYC data or device time zone | Moderate |
| Suspicious open ports | Moderate |
| Incognito sessions | Moderate |
| Cookies disabled | Moderate |
| Large number of sessions / devices (indicative of fraud ring / account farm) | High |

#### Identity Verification Network Risk

In Sandbox, when testing, make sure to lower risk thresholds for Network Risk, as normal testing behavior can trigger the Identity Verification Network Risk Check.

Plaid creates a device fingerprint based on device/session attributes, including IP address, location, browser plugins, OS settings, WebGL parameters, user agent, TCP settings, cookies, screen resolution, battery usage, and device memory. The fingerprint allows Plaid to identify when the same device is used for multiple Identity Verification sessions.

Plaid tracks the velocity of sessions per device across all customers. If a device is onboarding with many different Plaid clients in a short timeframe, this suggests coordinated fraud attempts. The dashboard will show counts for sessions in the last 24 hours, 7 days, 3 months, and all-time.

If a device is seen creating multiple accounts in a short period (e.g. several signups in one day), this is also flagged as high risk, since it is often indicative of fraud rings or account farming.

#### Phone risk

Phone risk checks for throwaway / burner phone numbers. Plaid checks for phone number registration with over 90 services. More links indicates lower risk.

#### Stolen identity risk and synthetic identity risk

Stolen identity risk and synthetic identity risk will be assessed only for US end users, and only if a full 9-digit Social Security Number is available.

Note that while stolen and synthetic identity risk are displayed in the Dashboard as percentages, the numbers do not indicate a percentage likelihood that the identity is stolen or synthetic. Instead, they merely represent that the score is normalized to be in the 0-100 range. A score of 70% or higher for synthetic identity, or a score of 90% or higher for stolen identity, indicates high risk.

Risk factors for stolen identity include:

- Multiple different identities associated with the same phone number, email, or SSN
- PII associated with deceased individual (e.g. SSN found in Death Master File)

Risk factors for synthetic identity include:

- Multiple different identities associated with the same phone number, email, or SSN
- PII associated with deceased individual (e.g. SSN found in Death Master File)
- SSN inconsistent with date of birth or address history
- Brief or no established usage of email address, phone number, or SSN (see [phone risk](/docs/identity-verification/risk-checks/#phone-risk) and [email risk](/docs/identity-verification/risk-checks/#email-risk) for more details on how usage history for these values is established)

#### Facial duplicate risk

Facial duplicate risk will be assessed only if either Selfie Checks or Document Verification is enabled.

Facial duplicate risk detects of user whose facial biometrics (from selfies or ID document portraits) match those from previous verification sessions. This is designed to catch repeat fraud attempts, synthetic identity attacks, and incentive abuse.

#### Behavior risk

Behavior risk analyzes a user's behavior to determine whether it appears human and genuine. Factors analyzed include:

- How fast a user types in their PII
- How accurately a user enters their PII
- Whether the data is copied and pasted
- The field order in which a user inputs data
- Mouse movement and scrolling patterns
- The use of autofill
- The frequency and device variety of entering the same or similar PII (see also [network risk](/docs/identity-verification/risk-checks/#identity-verification-network-risk)).

These are categorized into three risk buckets.

**Behavior:** “Risky” means the user’s interaction with the form is statistically anomalous compared to legitimate users, but does not necessarily show signs of being a bot or a fraud ring. For example, users typically enter their own PII fluently, so entering PII with hesitation or multiple corrections may indicate stolen identity.

**Fraud Ring:** "Yes" means the session matches known fraud ring patterns. For example, repeatedly entering and correcting similar PII on a single device or across multiple devices simultaneously can indicate a fraud ring attempting a synthetic identity attack.

**Bot:** "Yes" means the session exhibited patterns associated with non-human, non-autofill, automated data entry.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
