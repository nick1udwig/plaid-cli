---
title: "Payments (Europe) - Account Sweeping | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/virtual-accounts/account-sweeping/"
scraped_at: "2026-03-07T22:05:13+00:00"
---

# Account Sweeping

#### Sweep funds from your virtual account

Account sweeping periodically sweeps the available balance of a virtual account to a designated client-owned bank account. By default, automated account sweeping is not enabled.

Sweeping funds ensures you control how much balance you maintain in your virtual account. If a virtual account used for refunds or payouts also has account sweeping enabled, there may be requests that fail whenever a virtual account does not have sufficient funds.

#### Sweep funds from a Virtual Account

Make sure your virtual account is set up before following these steps. For more information on setting up an account, see [Managing virtual accounts](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/).

There are two ways you can sweep funds from your virtual account:

##### Automated

Automated sweeps are the preferred sweep method for pay-by-bank and pay-in use cases. If your wallet is set up primarily for payouts, you should not enable automated sweeping, as you need to maintain funds in your wallet balance in order to issue payouts.

For each virtual account you want automated account sweeping to be enabled for, provide your Account Manager with:

- The virtual account's `wallet_id`.
- The account details for the designated account which funds should be swept to.

The available balance of each virtual account will be swept once a day at 12:00 AM UTC indefinitely. Once your automated sweeping is set up, you can configure it within the Plaid Dashboard.

In the UK, if the sweep amount is over the Faster Payment Service £1M limit, the sweep will be sent via CHAPS. CHAPS payments are only processed on UK business days, between 6:00 AM and 6:00 PM (06:00–18:00), UK local time.

##### Manual

If you prefer to manually control account sweeping, you will need to manually [execute a Payout](/docs/payment-initiation/virtual-accounts/payouts/) each time you want to sweep funds.

#### Sweep requirements

When a sweep occurs, all transactions since the last sweep will be included in the sweep. For a sweep to occur, your account balance must be above the minimum balance. If this is not the case, no sweep will occur, and any unswept transactions will be rolled into the next sweep.

#### Sweep reporting

Sweep reporting enables you to know exactly which transactions are included in any given sweep. This helps improve reconciliation and avoid issues where transactions occur during the processing of the sweep. Sweep reporting is available for automated sweeps only.

You can find your sweep reports at **Plaid Dashboard > Payments > Accounts > Sweeps**.

[See an example sweep report](https://plaid.com/documents/sample-sweep-report.csv).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
