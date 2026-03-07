---
title: "Transfer - Reserves | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/reserves/"
scraped_at: "2026-03-07T22:05:26+00:00"
---

# Reserves

#### Learn about Transfer reserves requirements

Reserves are static funds held by Plaid to mitigate the financial risks associated with ACH debit transactions. Holding funds in reserve with Plaid operates on a sliding scale, meaning that increasing the amount in reserves can lead to progressively lower ACH debit hold times and increased limits, with a few exceptions.

#### Why reserves are required

Reserves cover potential losses that Plaid may experience as a result of processing ACH debits. Reserves are a common industry practice used by payment processors and financial institutions to ensure that businesses have sufficient balance to return funds to end-users when necessary.

ACH debits have financial risk because they can fail ("return") after settlement occurs. Even after funds have been transferred to Plaid, they can be returned back to the debited account up to 60 days after the transaction. This can occur due to a variety of reasons, including insufficient funds in the debited account; the debited account being inactive, closed, or blocked; or the consumer claiming the debit was unauthorized. If a customer does not have sufficient funds to cover the return, that will cause a negative balance and a potential loss to Plaid.

#### How Plaid uses reserves

If a Ledger's available balance goes negative due to ACH debit returns, the balance will usually recover after a few days due to pending funds coming off hold. However, if the amounts of the available and pending balances are insufficient to cover the returned amounts, then Plaid will debit your linked funding account for that Ledger to recover the balance.

Plaid will draw from reserves only if the owed amounts could not be recovered via payment processing or auto-initiated debits to the associated funding account.

#### Funding your reserves

During your onboarding process, Plaid will calculate the required amount of reserve funds and options for number of hold days based on your use case, maximum transaction size, and expected volume of transactions.

These options will be presented to you in your [Implementation checklist](https://plaid.com/docs/transfer/application/#implementation-checklist). After selecting the number of hold days, you can see the corresponding reserve amount required.

Your implementation checklist will show you the necessary details to initiate a wire transfer for the required reserve amount. Once Plaid receives the funds, they will be automatically allocated to your reserve balance.

Your current reserve balance is displayed in the [Dashboard](https://dashboard.plaid.com/transfer). See the "Account Information" section to view the total amount currently held.

#### Changes to your reserve balance

The required reserve amount may be adjusted based on changes to ACH debit hold times, transaction limits, or return rates.

If adjustments to your reserve amount are necessary, your Plaid Account Manager will communicate this to you and provide guidance.
For further assistance or to adjust your reserves, please contact your Plaid Account Manager directly.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
