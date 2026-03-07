---
title: "Payments and Funding | Plaid Docs"
source_url: "https://plaid.com/docs/payments/"
scraped_at: "2026-03-07T22:05:14+00:00"
---

# Payment products

#### Review and compare payment and account funding solutions

This page provides overviews of Plaid's payment and funding solutions to help you find the right one for your needs.

#### Auth

[Auth](/docs/auth/) is a flexible solution for account and routing number verification that can be used with either a Plaid payments [processor partner](/docs/auth/partnerships/), or any other processor of your choice, for payments processing.

#### Transfer

[See a live, interactive demo](https://plaid.coastdemo.com/share/66faf2bfb0feffe54db38114?zoom=100) of a Plaid-powered payment flow using Transfer.

[Transfer](/docs/transfer/) is an end-to-end payment processing solution that performs funds transfers.

#### Auth and Transfer comparison

Auth provides the most flexibility in terms of supported use cases and integration modes, while Transfer provides a single, end-to-end solution for funds transfers.

|  | Auth | Transfer |
| --- | --- | --- |
| Summary | An extensible account verification solution that plugs into your payments infrastructure | An out of the box, end-to-end solution for payments |
| Verifies account and routing numbers | Yes | Yes |
| Supports non-credential-based flows for account verification | Yes | Yes |
| Processes ACH transfers | No, requires a [processor partner or other processor](/docs/auth/#using-a-payments-service) | Yes |
| Provides optional out of the box UI for transfer authorization | No | Yes |
| Checks for ACH return risk | When used with Balance and/or Signal | Yes |
| Checks for account takeover risk | When used with Identity or Identity Match | When used with Identity or Identity Match |
| Supported countries | US, CA, UK, Europe | US |
| Supports peer-to-peer and marketplace payments | Yes | No |
| Supported use cases and industries | Wider range of supported use cases and industries | More restricted range of supported use cases and industries |
| Supported payment rails | Standard ACH, Same Day ACH, RTP, RfP, Wire, and international payment schemes | Standard ACH, Same Day ACH, RTP, RfP, Wire (for domestic credit transfers only) |
| Billing plans available | Pay-as-you-go or 12-month contract | 12-month contract (Custom only) |

#### Plaid Signal

Plaid Signal is Plaid's solution for ACH risk management. With Signal, you can use Signal Rules in the Plaid Dashboard to create and manage business logic for handling transactions.

Plaid Signal includes two separate products: Balance, which gets real-time balances; and Signal Transaction Scores, which uses ML modeling to assess transaction risk using over 80 attributes. You can purchase and use either Balance or Signal Transaction Scores by itself, or combine them for a more comprehensive ACH risk management approach.

#### Balance

[Balance](/docs/balance/) gets real-time balances to reduce ACH risk. Using Signal Rules, you create your own business logic to indicate how to process transactions based on comparing real-time balance and transaction amount. Balance is frequently used with Auth and also automatically included as part of Transfer; you can customize the Balance checks used via the Dashboard.

For non-ACH risk assessment use cases (e.g. treasury management or personal finance management), Balance can also be used without Signal Rules.

#### Signal Transaction Scores

[Signal Transaction Scores](/docs/signal/) uses ML modeling to predict the risk of ACH returns with the lowest latency and can be used with either Auth or Transfer. Using Signal Rules, you create your own business logic to indicate how to process transactions based on over 80 attributes. The Signal Dashboard also provides backtesting, rules analytics, and rules suggestions that optimize your revenue based on personalized business data, industry trends, and your desired risk profile.

#### Balance and Signal Transaction Scores comparison

Balance fetches an end user's real-time balance and is a simple and affordable solution for predicting insufficient funds ACH returns (over 75% of returns). It can also be used for non-ACH use cases, such as financial management apps that require displaying real-time balance information.

Signal Transaction Scores is a premium, higher-priced solution that provides greater accuracy and much lower latency in assessing return risk to reduce abandon rates during critical flows in the user funnel, such as checkout and account funding.

Signal Transaction Scores is recommended for higher-risk use cases, or for friction-sensitive user-present flows where avoiding user abandons during the risk check is critical.

Balance and Signal Transaction Scores both use Signal Rules and share a single integration path, using the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) endpoint. They are designed to be easily be used together in the same integration. Switching between the two can be done via no-code configuration the Signal Rules Dashboard, or by changing a single parameter in the [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) call.

|  | Balance | Signal Transaction Scores |
| --- | --- | --- |
| Summary | Get real-time balance data to reduce risk of ACH returns | Use ML modeling to quickly predict the risk of ACH returns |
| Advantages | Low cost, guarantees real-time balance data, supports low-code integration | Lowest latency, most accurate results, supports low-code integration as well as integrating with your existing risk engine |
| Can predict wide variety of ACH return codes | No, R01 only | Yes |
| Rules logic can be created and managed via Signal Rules Dashboard | Yes | Yes |
| Can be used for non-payments use cases requiring realtime balance | Yes | No |
| Supported countries | US, CA, UK, Europe | US |
| Supported payment methods | Any: ACH, wires, RTP, etc. | ACH only |
| Balance data freshness | Real-time | Cached, typically updated daily |
| p50 latency | ~ 3 seconds | ~ 1 second |
| p95 latency | ~ 11 seconds | < 2 seconds |
| Billing plans available | Pay-as-you-go or 12-month contract | 12-month contract (Custom only) |

#### Identity

[Identity](/docs/identity/) compares account info you hold about your customers to bank account ownership details to reduce account takeover risk. It is commonly used with Auth and Transfer, especially in high risk use cases.

#### Identity and Identity Match comparison

[Identity Match](/docs/identity/#identity-match) is an optional add-on to [Identity](/docs/identity/) that provides greater ease of integration and data minimization. Customers using Identity Match can experience onboarding conversion improvements of 20% or more, without increasing fraud rates, when using Identity Match versus using their own matching algorithms.

|  | Identity | Identity Match |
| --- | --- | --- |
| Summary | Get account holder info on file with bank | Compare account holder info with info on file at bank |
| Avoids the need to build complex string matching logic | No | Yes |
| Minimizes PII stored in system | No | Yes |
| Integrates with [Identity Verification](/docs/identity-verification) | No | Yes |
| Works with accounts verified via loginless flows such as [Same-Day Micro-deposits](/docs/auth/coverage/same-day) or [Database Insights](/docs/auth/coverage/database) | No | Supports ~30% of accounts verified by these flows |
| Supported countries | US, CA, UK, Europe | US, UK, Europe (CA: Early availability only, [contact sales](https://plaid.com/contact/)) |
| Billing plans available | Pay-as-you-go or 12-month contract | Pay-as-you-go or 12-month contract |

#### Investments Move

[Investments Move](/docs/investments-move/) (US/CA only) is designed specifically for automating data entry and reducing user friction for brokerage-to-brokerage ACATS and ATON transfers, helping reduce abandon rates and failed transfers for new users moving their assets to your brokerage. Unlike other products on this page, it cannot be used to move money between bank accounts, only to move holdings from one brokerage to another.

#### Payments (Europe)

Plaid's [Payments suite](/docs/payment-initiation/) enables your users to make real-time payments without manually entering their account number and sort code, or leaving your app.

#### Virtual Accounts (Europe)

[Virtual Accounts](/docs/payment-initiation/virtual-accounts/) enables wallet-based features for your payments. Virtual Accounts can be used alongside other Payments components to enhance its capabilities with returns, settlement status visibility, and payouts, and can also be used on its own to add payout capabilities to apps that accept bank transfers from other sources.

#### Payment Initiation and Virtual Accounts comparison

|  | Payments (without Virtual Accounts) | Virtual Accounts |
| --- | --- | --- |
| Summary | Enable end users to make real-time, in-app payments | Manage and track payments made by end users |
| Allows end-users to make payments via Link | Yes | No |
| Supports recurring payments | UK only, via [VRP](/docs/payment-initiation/variable-recurring-payments/) | UK only, via VRP, only if used with [Payments](/docs/payment-initiation/variable-recurring-payments/) |
| Supports issuing payouts | No | Yes |
| Supports issuing returns | No | Yes, only if used with Payments |
| Provides payment status updates | Yes | Yes, only if used with Payments |
| Indicates whether a payment has settled | No | Yes, only if used with Payments |
| Supported countries | [18 European countries](https://support.plaid.com/hc/en-us/articles/27895826947735-What-Plaid-products-are-supported-in-each-country-and-region) (including UK) | [18 European countries](https://support.plaid.com/hc/en-us/articles/27895826947735-What-Plaid-products-are-supported-in-each-country-and-region) (including UK) |
| Supports non-Eurozone local payments | Yes | Yes |
| Billing plans available | 12-month contract (Custom only) | 12-month contract (Custom only) |

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
