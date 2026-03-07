---
title: "Transfer - Plaid Ledger flow of funds | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/flow-of-funds/"
scraped_at: "2026-03-07T22:05:24+00:00"
---

# Plaid Ledger flow of funds

#### Understand how Transfer moves money with the Plaid Ledger

#### Plaid Ledger overview

The Plaid Ledger is a balance you keep with Plaid for use with the Transfer product. It is a settlement balance that all debit transfers settle into, and all credits are paid out from. Plaid provides several tools for moving money between your business checking account and your Ledger balance.

![Diagram of Plaid Ledger flow: Business account sweeps to Plaid ledger, transfers to customer account; details ACH credit/debit timing.](/assets/img/docs/transfer/ledger_overview.png)

Overview of the Plaid Ledger flow

By clicking on your Ledger in the Dashboard, you will be directed to the specific Ledger's page.

![Plaid Ledger Balance page showing summary, pending funds, and balance history with search and filters for event ID and date range.](/assets/img/docs/transfer/ledger-details.png)

Plaid Ledger details page

From this page, you can:

- View available and ending balances
- View the release schedule of pending funds
- Deposit funds into the Ledger
- Withdraw funds from the Ledger to a linked funding account
- View a real-time activity table of Ledger events and export the Ledger events as a CSV for reconciliation
- Retrieve Ledger balance events

#### Receiving payments

Your Plaid Ledger account has two balance values: `pending`, which contains funds that are in processing, and `available`, which contain funds that are available for use.

When receiving payments from your customers (via an ACH debit transfer), funds are placed in the `pending` balance as soon as your transfer is created. These funds become `available` after a hold period elapses. The hold period is typically 2-5 days after the transfer has settled; your specific hold period is unique to you and will depend on factors such as the reserve funds level you maintain with Plaid. The hold period will begin on the business day after the transaction emits the `settled` event and enters the `pending` state. Funds are typically made available at approximately 3:00PM ET on the last day of the hold period.

The projected date when funds in the `pending` state will be available is displayed on the Dashboard. It can also be obtained via any endpoint that returns a transfer object, such as [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget), in the `transfer.expected_funds_available_date` field. Once the funds have been made available, the [`funds_available` event](/docs/transfer/reconciling-transfers/) will be emitted.

See below for an example schedule of when an ACH debit is withdrawn from a consumer's account (corresponding to the `settled` event), and when the funds are made available to you.

![Same Day ACH debit schedule: Funds withdrawal on Mon-Wed, available by Fri-Mon, based on time submitted (EST).](/assets/img/docs/transfer/example_sameday_4_day.png)

Example schedule for a Same Day ACH debit.

![Standard ACH debit schedule: Funds withdrawn 1-3 days after submission, available on the 4th day based on submission time.](/assets/img/docs/transfer/example_standard_4_day.png)

Example schedule for a Standard ACH debit.

Note that the funds available date refers to when funds are available in your Ledger, not when the funds are withdrawn to your funding account. Withdrawing these funds to your funding account is performed via Same Day ACH and may take up to one additional business day.

If the ACH debit experiences a return or reversal, the funds will be clawed back from the Plaid Ledger balance. If the funds have already been made available and withdrawn to your business checking account, your available Ledger balance may go negative.

For more details on why a payment could be returned, see [Troubleshooting ACH Returns](/docs/transfer/troubleshooting/#ach-returns).

#### Issuing payouts

When issuing payouts to your customers (via ACH credit and RTP/FedNow transfers), the funds for the payout are deducted from your available balance and are sent out to your consumer in the next available ACH window, or immediately for RTP/FedNow. If there are not enough funds in your available balance to cover a transfer, then [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) will decline the transfer authorization with `NSF` (insufficient funds) as the `decision_rationale`.

You can check the value of the Plaid Ledger balance any time via the Dashboard or by calling [`/transfer/ledger/get`](/docs/api/products/transfer/ledger/#transferledgerget).

In order to avoid declined transfers, you should discuss with your company the level of funds you want to keep in the balance, and how you intend to refill the balance once it has been used.

How much money to keep in your Ledger will depend on your business cash flow constraints. It is recommended to keep a minimum of three days' worth of funds in your balance to cover activity taking place over the weekend, when ACH credits and wires are unavailable.

Even if your use case is for debit transfers only, you should retain enough funds in your Ledger to handle refunds and returns. Maintaining a zero or negative Ledger balance will prevent you from issuing refunds.

#### Moving money between Plaid Ledger and your bank account

It is important to consider your strategy for moving funds between your Plaid Ledger balance and your business checking bank account. The Plaid Ledger balance cannot be used to store funds indefinitely.

Plaid provides four methods to manage your Plaid Ledger balance. Sweeps to or from your Ledger balance are priced the same as a transfer (for example, a `network=rtp` Ledger balance withdrawal is priced as an RTP transfer). If using Same Day ACH and your sweep amount is over the Same Day ACH transaction limit of $1M, Plaid will automatically break your sweep request into multiple transfers.

Each of the following methods will create a sweep object, viewable in [`/transfer/sweep/list`](/docs/api/products/transfer/reading-transfers/#transfersweeplist), and will emit `sweep.*` events for you to observe. See [Reconciling Transfers](/docs/transfer/reconciling-transfers/) for more details.

You must ensure that your accounting department does not reject pulls from Plaid to deposit money into your Ledger balance. Contact your bank to ensure you accept incoming ACH debits from the following company or Originator ID: 1460820571.

| Method | Optimizes for... |
| --- | --- |
| [Dashboard](/docs/transfer/flow-of-funds/#moving-money-via-the-dashboard) | Ad-hoc manual deposits/withdrawals |
| [API endpoints](/docs/transfer/flow-of-funds/#moving-money-via-api) | Custom deposit/withdrawal logic |
| [Automatic balance thresholds](/docs/transfer/flow-of-funds/#automatically-managing-the-balance-level) | Ongoing hands-off management |
| [Pushing funds to Ledger via wire or ACH](/docs/transfer/flow-of-funds/#pushing-funds-to-your-ledger-via-ach-rtp-or-wire) | Immediate funds availability |

All methods of moving money to your Plaid Ledger, except for pushing funds via wire or ACH, are subject to a three business day hold. Funds will be made available at approximately 3:00 PM on the third business day after funds [settle](/docs/transfer/creating-transfers/#ach-processing-windows) to Plaid. Note that this duration may not be the same as the hold time applied to funds collected from debit transactions. Funds pushed by wire or ACH are available immediately upon settlement.

If you are a [Transfer for Platforms (beta)](/docs/transfer/platform-payments/) customer, the API is the only method available for moving funds. For more details on moving money as a Platform, see [the Transfer for Platforms documentation](/docs/transfer/platform-payments/#payment-collection).

##### Moving money via the Dashboard

You can use the Plaid Dashboard to manually trigger a deposit or withdrawal. Withdrawals will push the funds to your business checking account in the next Same Day ACH window. Deposits into your Plaid Ledger balance will pull from (debit) your business checking account using Same Day ACH and put the funds in your pending balance.

To move money via the Dashboard, you must have the "Manage Transfer Ledger Funds" permission, available via the [Team member setting page](https://dashboard.plaid.com/settings/team/members).

Deposits that are pulled (debited) from your business checking account will be made available three business days after funds settle to Plaid. The deposit will display the `funds_available` status in the Dashboard once the funds are available to use in your Ledger.

###### Adding or changing funding accounts

To add or change a funding account, navigate to the “Account Details” page of the Transfer dashboard and click on “Add account”.

![Dashboard UI showing 'Account details' section with a highlighted 'Add account' button for managing funding accounts.](/assets/img/docs/transfer/example_add_funding_account.png)

You will then be able to use Plaid Link to automatically connect your account and verify that it is open, active, and belongs to you.  For accounts that are unable to be linked via Plaid, there will be an option to manually enter account and routing information, and submit two bank statements to confirm your ownership of the account. Manually linked accounts will be reviewed by the Plaid team, and you will receive an email confirmation upon approval.

When new funding accounts are added, an email notification is sent to admins on the Plaid account notifying them of the change.

Within the dashboard, you can also specify which funding account to use as the default account for sweeps, give the account a friendly name, or remove the account.

##### Automatically managing the balance level

You can set up automatic balance thresholds, allowing you to operate in a "set and forget" model. You can set both a minimum balance and a maximum balance. For the reasons described in [Issuing payouts](/docs/transfer/flow-of-funds/#issuing-payouts), it is not recommended to set a zero dollar minimum balance.

If you don't set a minimum balance, or set a minimum balance of zero dollars, Plaid will automatically debit your linked funding account if you have a negative available balance that will not be covered by pending incoming funds.

For example, if you have an available balance of -$500, and $300 in pending funds being released over the next three days, Plaid will debit your linked funding account by $200.

Every business day at 3:00 PM EST, Plaid will attempt to withdraw or deposit funds, depending on your configuration. If your available Ledger balance exceeds the maximum balance defined, Plaid will move money from your Ledger to your bank account using Same Day ACH. These funds should be available in your business checking account by 6:00 PM EST.

If your overall Ledger balance (including both pending and available balances) goes below your minimum balance, Plaid will move funds from your bank account to your Ledger using Same Day ACH. These funds will be placed in the pending balance, and made available 3 business days after the funds settle to Plaid.

To enable automatic refills and withdrawals, click the "settings" button on the [Ledger page](https://dashboard.plaid.com/transfer/ledgers/). To change this setting, you must have the "Manage Transfer Ledger Funds" permission, available via the [Team member setting page](https://dashboard.plaid.com/settings/team/members).

![Plaid dashboard showing Ledger summary, balance history, and settings for maintaining minimum balance in transfer flow.](/assets/img/docs/transfer/transfer-ledgers.png)

##### Moving money via API

The API is the most flexible method for moving funds, as you can determine when and how your money moves: via Same Day ACH, standard ACH, or instantly via RTP/FedNow.

To deposit funds into your Plaid Ledger balance via API, use [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit). This endpoint will issue a [sweep](/docs/transfer/creating-transfers/#sweeping-funds-to-funding-accounts) to withdraw funds from your funding account. Once Plaid has received the funds (you may specify either standard or Same Day ACH), they will be made available after a three business day hold.

To see when the transfer is available in your Ledger, you can use the [`/transfer/event/list`](/docs/api/products/transfer/reading-transfers/#transfereventlist) or [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) endpoint to check for the `sweep.funds_available status`. You can also check the value of the Plaid Ledger balance via API by calling [`/transfer/ledger/get`](/docs/api/products/transfer/ledger/#transferledgerget).

You can retrieve Plaid Ledger balance events by calling the `transfer/ledger/event/list` endpoint with specified filter criteria. For example, you could search for all events for a specific `ledger_id`.

To withdraw funds from your Plaid Ledger balance via API, use [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw). You can choose the rails to use for the transfer: Same Day ACH, standard ACH, or instant.

##### Pushing funds to your Ledger via ACH, RTP, or wire

To move funds into your Ledger balance, you can an initiate a push of funds from your bank account to Plaid. In this approach, you initiate the funds transfer through your bank, typically via their website. To obtain the account and routing number for your Ledger, in the Transfer Dashboard, navigate to the Ledger details page and click the "Settings" icon in the upper right.

Transfer Ledger balances can receive funds pushed via ACH credit, RTP, or wire. Pushed funds become available for use in your Plaid Ledger as soon as they are settled to Plaid. Incoming wires pushed to the Plaid Ledger are charged a separate fee, billed to your account. Incoming ACH and RTP credits pushed to the Plaid Ledger incur no fee. For more details, contact your Plaid Account Manager.

#### Maintaining multiple Ledger balances

You can maintain multiple Ledger balances with Plaid. This is useful when you need separation between certain types of transactions for accounting purposes, or to connect different transactions with different bank accounts. For instance, you can use multiple Ledgers to:

- Use separate funding accounts to issue credits and collect debits.
- Keep accounting separate for different use cases, business units, or states where you operate.

Every Plaid client has one Ledger created and ready to use as their default Ledger to originate transfers or sweeps.

You can create additional Ledgers in the Plaid Transfer dashboard. From the [Ledgers page](https://dashboard.plaid.com/transfer/ledgers), click "add". You will be prompted to assign a name for the Ledger and select one of your active funding accounts as the default funding account for this new Ledger. When you select a funding account for your Ledger, all of your automatic deposits and withdrawals flow to and from that funding account.

![Create a new ledger with fields for 'Ledger name' and 'Default funding account'. Checkbox to set as default ledger.](/assets/img/docs/transfer/add-ledger.png)

Adding a Ledger

If you need to create more than ten total Ledgers, submit a [support ticket](https://dashboard.plaid.com/support/new/admin/account-administration).

You can also view the detailed info of each Ledger via the Dashboard, including the Ledger’s id, name, balance, transaction history, and pending balance.

![Default Ledger Balance page showing a summary. Pending funds timeline and balance history are also visible.](/assets/img/docs/transfer/multiple-ledgers.png)

List of multiple Ledgers

Each Ledger has its own separate balance settings, including the funding account and minimum and maximum balance rules. You can share a single funding account across multiple Ledgers, or give each Ledger its own funding account.

Once created, additional Ledgers cannot be deleted. This is because Ledgers function as audit logs for your transaction history. If you no longer wish to use a Ledger, you can stop assigning transactions to it.

You can also move funds between Ledgers, by calling [`/transfer/ledger/distribute`](/docs/api/products/transfer/ledger/#transferledgerdistribute). Note that only funds in a Ledger's available balance can be distributed to another Ledger.

##### Working with multiple Ledgers in the API

If you have created multiple Ledgers, you should indicate which Ledger to use by specifying a `ledger_id` when calling [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate), [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit), and [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw). If no `ledger_id` is specified, Plaid will use your default Ledger. You can change which Ledger is considered the default in the Dashboard.

Each Ledger's balance history is reported separately.

Transfer objects, including `authorization`, `transfer`, `transfer_event`, `sweep` and `ledger`, will contain a `ledger_id` field, indicating which Ledger the object is associated with.

![Multi-ledger flow: Debit transfers to Pay-ins ledger, credit sweep to Funding account 1; Debit sweep from Funding account 2 to Pay-outs ledger, credit transfers to user accounts.](/assets/img/docs/transfer/multi-ledger-funds-flow.png)

Example multi-Ledger funds flow, using separate funds for issuing credits and collecting debits

![Flow diagram showing debit transfers from end user accounts to Ledgers 1, 2, and 3, then credits to Funding accounts 1 and 2.](/assets/img/docs/transfer/multi-ledger-funds-flow2.png)

Example multi-Ledger funds flow, using separate funds for different states

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
