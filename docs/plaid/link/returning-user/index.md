---
title: "Link - Returning user experience | Plaid Docs"
source_url: "https://plaid.com/docs/link/returning-user/"
scraped_at: "2026-03-07T22:05:08+00:00"
---

# The returning user experience

#### Learn how Plaid streamlines the user experience for returning users

#### Overview

The returning user experience (formerly known as Remember Me) streamlines onboarding for users who have already connected a financial account with Plaid in the US or Canada. In Link, users can choose to associate their phone number with the accounts they're connecting to a financial app or service. Once users have opted-in to being 'remembered' by Plaid, they'll be able to quickly connect those same accounts to other financial apps and services in the future using a one-time password for thousands of financial institutions, resulting in higher conversion and a simpler user experience.

On the Consent screen in Link, users can input their phone number and verify it using a one-time password sent to their device. Next, the user will proceed to select an institution and connect their account(s). Once account verification is completed, Plaid will associate the institution and accounts with the user's phone number. If a user has previously connected an account, Link may streamline the flow by reusing an existing connection and taking the user directly to account selection. If Link can’t reuse that connection, the user will go through the full Link flow to reconnect.

The returning user experience is automatically enabled for all eligible customers, and you do not need to make any updates to your integration to support it.

The legacy returning user experience flows (Institution Boosting, Pre-Matched RUX, and Pre-Authenticated RUX) have been replaced by the revamped returning user experience (formerly known as the Remember Me flow) as of October 28, 2024. For questions or more details, contact your Plaid Account Manager.

##### The returning user experience for a new user

The user starts a Link session and enters their phone number.

![The user starts a Link session and enters their phone number.](/assets/img/docs/rux/rux_tour/rux-comcon.png)

![Plaid does not recognize this number as one that's connected to Plaid before. The user is prompted to confirm their number through an SMS code.](/assets/img/docs/rux/rux_new/rux_new_3.png)

![The user selects an institution to connect with.](/assets/img/docs/rux/rux_new/rux_new_4.png)

![They enter their credentials, or proceed through an OAuth flow.](/assets/img/docs/rux/rux_new/rux_new_5.png)

![...and they select their account.](/assets/img/docs/rux/rux_new/rux_new_6.png)

![This user has successfully connected to this institution, and it's now available for them to connect to the next time they use Plaid Link.](/assets/img/docs/rux/rux_new/rux_new_7.png)

##### The returning user experience for a saved user

The user starts a Link session and enters their phone number.

![The user starts a Link session and enters their phone number.](/assets/img/docs/rux/rux_tour/rux-comcon.png)

![Plaid recognizes this number as one that's connected to Plaid before, and asks the user to confirm a code sent by SMS.](/assets/img/docs/rux/rux_new/rux_new_3.png)

![The user can then instantly connect to an institution that they've previously connected to with Plaid.](/assets/img/docs/rux/rux_tour/rux_4.png)

![They can also choose which saved accounts to connect to your app.](/assets/img/docs/rux/rux_tour/rux_5.png)

##### Returning user experience

When users want to connect their saved institutions and accounts to additional Plaid-powered apps or services, Plaid runs security checks to detect that they are a returning user with the same phone number and device.

Users can choose to be remembered by Plaid, making future connections faster and easier. To enroll, the user will enter their phone number in Link. This flow is available for phone numbers with country codes +1 (covering the United States, Canada, and part of the Caribbean), +44 (UK), and +52 (Mexico); users selecting codes from any other country will be redirected to the standard Link flow. The default country shown in the dropdown will be based on the country code(s) provided in the [`/link/token/create`](/docs/api/link/#linktokencreate) call; for example, if [`/link/token/create`](/docs/api/link/#linktokencreate) is called with a single country code, that country will be shown as the default; if it is called with multiple countries, including `US`, the United States will be shown as the default.

#### Pre-filling phone numbers for faster account verification

Link sessions can be enabled for a more streamlined user experience when Plaid already knows the user’s phone number. When a `user.phone_number` is provided via [`/link/token/create`](/docs/api/link/#linktokencreate), Plaid will pre-fill the phone number in Link for the user. Only the last 4 digits of the associated phone number are shown in order to preserve user privacy.

The user can then verify their phone number. Plaid will deploy a number of security checks to verify that the phone number belongs to the device before the user can select which saved institution to connect to an app or service. Pre-filling phone numbers can help boost conversion while reducing the number of manual inputs from users.

#### Passkeys

Passkeys are a secure alternative to passwords and can provide a streamlined account linking experience for thousands of institutions. Plaid will automatically enable passkeys for Link sessions on iOS devices when you use the Plaid iOS SDK 4.3.0+ or React Native 10.2.0+ SDK.

In Link, users will be prompted to enable passkeys with Face or Touch ID before proceeding with connecting their financial accounts to a chosen app or service. When the same user wants to connect to another app or service, they can use their passkey with Face or Touch ID instead of a one-time password for a streamlined authentication experience.

![Plaid passkey flow for returning users: Connect with customer app, input phone, log in with Face/Touch ID, select saved account.](/assets/img/docs/rux/passkeys.png)

#### Passport

Passport provides a streamlined experience where end users can opt-in to saving their financial account connections and personal information in a Plaid account. They can then securely and easily share this account with some Plaid-powered apps and services.

Passport is available now for [Consumer Report by Plaid Check](/docs/check/). End users can use Passport to share their consumer report with Plaid Check customers and partners to receive credit-related services. This experience is automatically enabled for eligible customers, and you do not need to make any updates to your integration to support it.

To view the Passport experience, you can launch Link with any [Plaid Check](/docs/check/) product in the Sandbox environment.

#### Silent Network Authentication

Silent Network Authentication (SNA) is an aspect of the mobile returning user experience that enables Plaid to verify a consumer's phone number using their mobile network carrier. Eligible users can skip manual OTP verification of their phone number, which removes one step in the Link flow. SNA is currently used for all eligible Layer and Plaid Check Consumer Reports flows. SNA is also occasionally used for eligible returning user experience flows for other products.

SNA is supported on mobile flows using Plaid's native or React Native mobile SDKs using Android SDK version >=5.6.0 and iOS SDK version >=4.7. Customers using Layer or Plaid Check Consumer Reports are strongly recommended to upgrade to iOS SDK version >=6.0.1 (or React Native iOS >=12.0.2) in order to take advantage of substantial performance improvements to the SNA flow. Aside from using a compatible SDK, no work is required on your part to enable SNA.

SNA is compatible with AT&T, T-Mobile, and Verizon, and works by verifying the user's SIM is connected to the mobile network carrier and not spoofed or cloned. If SNA cannot verify the user's number (for example, because verification failed, their network is not compatible, or mobile data is unavailable) Plaid will fall back to the manual OTP flow.

SNA is available only in Production and cannot currently be emulated in Sandbox.

#### Testing in Sandbox

The returning user flow can be tested in the Sandbox or Production environments.

Real phone numbers do not work in Sandbox. Instead, Sandbox has been seeded with a test user whose phone numbers may be used to trigger different scenarios. To explore each scenario, enter the corresponding phone number and correct OTP. For all scenarios, the correct OTP is `123456`.

Returning User: A user who has previously enrolled in the returning user experience by confirming their device and successfully linking an Item.

| Link Returning User Sandbox Scenarios | Seeded Phone Number |
| --- | --- |
| New User | `415-555-0010` |
| Verified Returning User | `415-555-0011` |
| Verified Returning User: linked new account | `415-555-0012` |
| Verified Returning User: linked OAuth institution | `415-555-0013` |
| Verified Returning User + new device | `415-555-0014` |
| Verified Returning User: automatic account selection | `415-555-0015` |

#### Tracking events

Link emits events to indicate whether or not users opt-in to being remembered by Plaid:

##### Events

| Event | Meaning |
| --- | --- |
| `SUBMIT_PHONE` | User has provided their phone number to be remembered by Plaid |
| `SUBMIT_OTP` | User has entered an OTP code to verify their phone number |
| `VERIFY_PHONE` | User has successfully verified their phone number |
| `SKIP_SUBMIT_PHONE` | User chose not provide their phone number to be remembered by Plaid |
| `CONNECT_NEW_INSTITUTION` | User chose to connect a new institution |

##### View names

The following can be found in the `view_name` field in the `TRANSITION_VIEW` event for returning user panes:

| View name | Meaning |
| --- | --- |
| `SUBMIT_PHONE` | User was prompted to provide their phone number to be remembered by Plaid |
| `VERIFY_PHONE` | User was prompted to verify their phone number |
| `SELECT_SAVED_ACCOUNT` | User was prompted to select the underlying account from the saved Item |
| `SELECT_SAVED_INSTITUTION` | User was prompted to select one of multiple saved Items |

##### Match reasons

The `match_reason` field in the `SELECT_INSTITUTION` event has the following values for the returning user flow:

| Match reason | Meaning |
| --- | --- |
| `AUTO_SELECT_SAVED_INSTITUTION` | The `SELECT_SAVED_INSTITUTION` pane was skipped |
| `SAVED_INSTITUTION` | User selected a saved institution |
| `SAVED_ACCOUNT` | User selected a saved account |

##### Error events

Link will emit the `ERROR` event when the user submits an invalid phone number or an invalid OTP. The `error_code` will be `INVALID_PHONE_NUMBER` or `INVALID_OTP`, respectively.

For more details, see [Link SDK documentation](/docs/link/web/). For more information on tracking Link conversion in general, see [Improving Link conversion](/docs/link/best-practices/#improving-link-conversion).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
