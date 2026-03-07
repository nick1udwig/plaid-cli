---
title: "Investments Move - Introduction to Investments Move | Plaid Docs"
source_url: "https://plaid.com/docs/investments-move/"
scraped_at: "2026-03-07T22:04:59+00:00"
---

# Introduction to Investments Move

#### Automate ACATS and reduce rejections with broker-sourced data

Get started with Investments Move

[API Reference](/docs/api/products/investments-move/)

#### Overview

Investments Move (US and CA only) allows you to obtain broker-sourced data to automate brokerage account transfers, reduce operational complexity, and avoid rejections. You can use Investments Move to replace your existing manual ACATS data entry flow or build a new, high-performing flow. Investments Move is designed to increase both the number of users who submit an ACATS transfer and the rate at which these requests successfully settle.

Using Investments Move, a user accesses their brokerage account via a Plaid Link-initiated flow, and Plaid returns the source of truth data needed to submit an ACATS transfer, such as account number, account type, account holder name, holdings data, and DTC codes. For brokerages in Canada, Investments Move can retrieve the data necessary for an ATON transfer.

Investments Move is currently in Early Availability and available on a limited basis. To request access, if you are new to Plaid, [contact Sales](https://plaid.com/contact/); if you are already a Plaid customer, contact your Account Manager.

Looking for investment holdings or trading activity, without initiating ACATS transfers? See [Investments](/docs/investments/) instead.

#### Integration overview

The steps below show an overview of how to integrate Investments Move.

1. Call [`/link/token/create`](/docs/api/link/#linktokencreate). Along with any other parameters you specify, make sure to include the following:
   - The `products` array should be set to `["investments_auth"]`.
   - (Optional) To enable [fallback flows](/docs/investments-move/#fallback-flows), set the corresponding flag within the `investments_auth` object. For details, see [Fallback flows](/docs/investments-move/#fallback-flows).
2. On the client side, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see the [Link documentation](/docs/link/).
3. Once the end user has completed the Link flow, exchange the `public_token` for an `access_token` by calling [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange).
4. Call the [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) endpoint with the `access_token` you received from the token exchange. This endpoint will return all the information needed to initiate an ACATS transfer, including account holder name, account number, and DTC number.

#### Fallback flows

By default, an end user will only be able to select institutions where Plaid can obtain the full account number, account holder name, and holdings data. You can optionally choose to expand the number of institutions available by enabling one or more of the flows below, which allow the end user to directly enter data that Plaid was not able to obtain directly from the broker.

To know what data came directly from the brokerage, what was entered by the user and verified against partial data from the brokerage, and what was entered by the user but not verified by the brokerage because no data was available, you can use the `data_sources` object in the [`/investments/auth/get`](/docs/api/products/investments-move/#investmentsauthget) response.

##### Masked Number Match

Masked Number Match supports institutions where Plaid is able to obtain broker-sourced holdings data and account holder information but can only obtain a partial (masked) account number, typically the last four digits of the account number. During the Link flow, the user will be prompted to enter their full account number, which will be verified against the partial account number; if the numbers do not match, the Link attempt will fail and the user will be required to re-enter the account number.

To enable Masked Number Match, in the [`/link/token/create`](/docs/api/link/#linktokencreate) request, set `investments_auth.masked_number_match_enabled` to `true`.

##### Stated Account Number

Stated Account Number supports institutions where Plaid is able to obtain broker-sourced holdings data, and may or may not be able to obtain account holder information, but cannot obtain any part of an account number. During the Link flow, the user will be prompted to enter their full account number. If Plaid was unable to obtain account holder information, the user will also be prompted to enter their full name.

To enable Stated Account Number, in the [`/link/token/create`](/docs/api/link/#linktokencreate) request, set `investments_auth.stated_account_number_enabled` to `true`.

##### Manual Entry

Manual Entry supports institutions where Plaid cannot return any information. During the Link flow, users will be able to manually enter the information needed to submit the ACATS transfer. Plaid will still attempt to standardize and enrich the data they provide.

To enable Manual Entry, in the [`/link/token/create`](/docs/api/link/#linktokencreate) request, set `investments_auth.manual_entry_enabled` to `true`.

#### Testing Investments Move

Investments Move can be tested in [Sandbox](/docs/sandbox/) without any additional permissions.

To test fallback flows in Sandbox, use the institution Houndstooth Bank (`ins_109512`). The first fallback flow that is explicitly enabled in the [`/link/token/create`](/docs/api/link/#linktokencreate) call will be used. The flows will be attempted in the following order: Masked Number Match, Stated Account Number, and Manual Entry. (So for example, if both Masked Number Match and Manual Entry are enabled, Masked Number Match will be used.)

#### Investments Move pricing

Investments Move is billed on a [per-request flat fee model](/docs/account/billing/#per-request-flat-fee). To view the exact pricing you may be eligible for, [contact Sales](https://plaid.com/contact/).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
