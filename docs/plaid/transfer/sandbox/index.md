---
title: "Transfer - Testing in Sandbox | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/sandbox/"
scraped_at: "2026-03-07T22:05:26+00:00"
---

# Transfer Sandbox

#### Use Sandbox to test your transfer integration

While all newly created teams can use Transfer in Sandbox by default, some older teams may need to be manually enabled. If you receive an error indicating that the product is not enabled for your account when attempting to use Transfer in the Sandbox environment, contact Support or your Plaid Account Manager to request access.

#### Introduction to Sandbox

Sandbox is Plaid's environment for testing using fake data. For more details about getting started with Sandbox environment, including how to skip the Link flow in Sandbox with `/sandbox/item/public_token/create`, and testing non-Transfer-specific functionality such as Item errors, see [Sandbox](/docs/sandbox/).

#### Simulating money movement, events, and webhooks

When creating transfers in Sandbox, see the [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) documentation on how to generate approved or declined authorization responses.

Any cutoff times applied in Production will also apply in the Sandbox environment. For example, creating a same-day ACH transfer in Sandbox after 3:00 pm ET will automatically update the transfer to standard ACH.

In the Sandbox environment, no real banks are involved and no events are triggered automatically. By default, all transfers, refunds, and sweeps created in Sandbox remain at the `pending` status until you actively change them. This can be most easily accomplished through the Plaid Dashboard, where you can simulate next steps and common failure scenarios for any transfer that you've created in the Sandbox environment.

![Plaid Dashboard UI: Buttons to simulate sandbox transfer events. Options: 'Next Event,' 'Failure,' and 'Return.'](/assets/img/docs/transfer/sandbox_dashboard.png)

You can simulate many Sandbox events using the Plaid Dashboard.

If you would like to change the status of a transfer through the API, you can call the following endpoints to simulate events and fire transfer webhooks manually.

| Action to simulate | Simulation endpoint |
| --- | --- |
| Transfer events | [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate) |
| Refund events | [`/sandbox/transfer/refund/simulate`](/docs/api/sandbox/#sandboxtransferrefundsimulate) |
| Ledger withdrawals sweep events | [`/sandbox/transfer/ledger/withdraw/simulate`](/docs/api/sandbox/#sandboxtransferledgerwithdrawsimulate) |
| Ledger deposits sweep events | [`/sandbox/transfer/ledger/deposit/simulate`](/docs/api/sandbox/#sandboxtransferledgerdepositsimulate) |
| Ledger pending-to-available funds movements | [`/sandbox/transfer/ledger/simulate_available`](/docs/api/sandbox/#sandboxtransferledgersimulate_available) |
| Webhooks | [`/sandbox/transfer/fire_webhook`](/docs/api/sandbox/#sandboxtransferfire_webhook) |

When you change the status of a Sandbox transfer, either through the Dashboard or the API, the corresponding webhook will not fire. You must use the [`/sandbox/transfer/fire_webhook`](/docs/api/sandbox/#sandboxtransferfire_webhook) API to send a Transfer webhook in Sandbox.

For a full list of transfer simulations available in Sandbox, see [Sandbox endpoints](/docs/api/sandbox/).

For a general overview of the Plaid Sandbox environment, see [Sandbox overview](/docs/sandbox/).

#### Automatic Sandbox simulations

In addition to manual simulations, the Sandbox environment supports automatic state transitions for transfers, refunds, and sweeps (deposits and withdrawals) when using specific test amounts. These special amounts trigger predefined state transitions without requiring manual intervention through the Dashboard or API.

Once you initiate a transfer with any of these specific amount values, the transfer's stages will be immediately simulated. The appropriate webhooks for Sandbox transfer events will be fired, and you can fetch the events through [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync). You can also view the events' progression on your Transfer Dashboard, on the details page for the specific transfer.

In Production, transfer state changes can take minutes to hours depending on processing time. Sandbox simulations provide immediate state transitions for testing purposes.

##### Transfer state transitions

Creating Sandbox transfers with the specified amounts via [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) will result in automatic state transitions.

| Amount | State Transition Path |
| --- | --- |
| $11.11 | pending → posted → settled → funds available (ACH debits only) |
| $22.22 | pending → failed |
| $33.33 | pending → posted → settled → funds available (ACH debits only) → returned (R01) |
| $44.44 | pending → posted → settled → funds available (ACH debits only) → returned (R02) |
| $55.55 | pending → posted → settled → funds available (ACH debits only) → returned (R16) |
| $66.66 | pending → posted → settled → returned |

##### Refund state transitions

Creating Sandbox refunds with the specified amounts via [`/transfer/refund/create`](/docs/api/products/transfer/refunds/#transferrefundcreate) will result in automatic state transitions.

| Amount | State Transition Path |
| --- | --- |
| $1.11 | pending → posted → settled |
| $2.22 | pending → failed |
| $3.33 | pending → posted → settled → returned |

##### Sweep State Transitions

Creating Sandbox sweeps with the specified amounts via [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit) and [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw) will result in automatic state transitions.

| Amount | State Transition Path |
| --- | --- |
| $111.11 | pending → posted → settled → funds available (deposits only) |
| $222.22 | pending → failed |
| $333.33 | pending → posted → settled → funds available (deposits only) → returned |
| $444.44 | pending → posted → settled → returned |

#### Testing Plaid Ledger flow of funds

While testing Plaid Ledger in Sandbox, you can always call the following endpoints at any time to verify expected behavior:

- [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) to validate the transfer status change
- [`/transfer/sweep/get`](/docs/api/products/transfer/reading-transfers/#transfersweepget) to validate the deposit/withdrawal status change
- [`/transfer/ledger/get`](/docs/api/products/transfer/ledger/#transferledgerget) to validate the Ledger balance change

#### Wire transfers

In Sandbox, all customers are enabled for wire transfers in order to facilitate testing. Having access to wire transfers in Sandbox does not mean you are enabled for wires in Production. To request access to wire transfers in Production, contact your Account Manager.

##### Issuing payouts with a Plaid Ledger

###### Adding funds to the Plaid Ledger

All new Plaid Ledgers have a starting balance of $100 in Sandbox. Once this is depleted, or if you wish to simulate a larger transfer, you will need to fund the Ledger before testing the payout.

1. Call [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit) to create a sweep that adds funds to a Plaid Ledger balance. The funds will immediately show up in the `pending` balance.
2. Call [`/sandbox/transfer/ledger/deposit/simulate`](/docs/api/sandbox/#sandboxtransferledgerdepositsimulate) with `sweep.posted` as the `event_type`.
3. Call [`/sandbox/transfer/ledger/deposit/simulate`](/docs/api/sandbox/#sandboxtransferledgerdepositsimulate) with `sweep.settled` as the `event_type`. This will move your sweep to `settled` status.
4. Call [`/sandbox/transfer/ledger/simulate_available`](/docs/api/sandbox/#sandboxtransferledgersimulate_available) in order to simulate the passage of time and transition the funds from `pending` to `available`.

###### Issuing a payout

1. Call [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) with `type=credit` and your desired network. To simulate an insufficient funds failure, call [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) with an amount larger than your current available balance.
2. Call [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate). This will immediately create a transfer and decrement the available balance.

##### Receiving payments via Plaid Ledger

1. Call [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) with `type=debit`.
2. Call [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate). This will immediately create a transfer and increment your pending balance in the Plaid Ledger.
3. Call [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate) with `posted` as the `event_type`.
4. Call [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate) with `settled` as the `event_type`. This will move your transfer to `settled` status. The funds will remain in `pending` until the hold is elapsed.
5. Call [`/sandbox/transfer/ledger/simulate_available`](/docs/api/sandbox/#sandboxtransferledgersimulate_available) to simulate the passage of time and convert the Ledger balance to `available`. If you are a Transfer for Platforms customer, this will convert all pending balances in all Plaid Ledgers.
6. Call [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw) to create a sweep that withdraws funds from the Plaid Ledger balance. The endpoint will immediately decrement the available balance and create a sweep with `pending` status.
7. Call [`/sandbox/transfer/ledger/withdraw/simulate`](/docs/api/sandbox/#sandboxtransferledgerwithdrawsimulate) with `sweep.posted` as the `event_type`.
8. Call [`/sandbox/transfer/ledger/withdraw/simulate`](/docs/api/sandbox/#sandboxtransferledgerwithdrawsimulate) with `sweep.settled` as the `event_type`. This will move your sweep to `settled` status.

In Production, transfer and sweep status will be updated automatically to reflect the real processing status, and the `pending` balance will automatically become `available` after the client-specific hold day.

##### Simulating a return with Plaid Ledger

1. Follow steps 1-3 in [Receiving payments via Plaid Ledger](/docs/transfer/sandbox/#receiving-payments-via-plaid-ledger) section to create a debit and move it to `posted`.
2. (Optional) To simulate a return after the funds were made available, follow steps 4-5 as well. If you would like to simulate a return before the funds were made available, omit this step.
3. Call [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate) with `returned` as the `event_type`. This will move the transfer status to returned and decrement pending balance, which you can verify by calling [`/transfer/ledger/get`](/docs/api/products/transfer/ledger/#transferledgerget).

###### Testing Instant Payouts

Testing Instants Payouts works the same way as testing ACH Transfers.

To test the [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget) endpoint in Sandbox, log in using the `user_good` user Sandbox account (see [Sandbox simple test credentials](/docs/sandbox/test-credentials/#sandbox-simple-test-credentials) for more information), and use the first 2 checking and savings accounts in the "First Platypus Bank" institution (ending in 0000 or 1111), which will return `true`. Any other account will return `false`. If using a [custom Sandbox user](/docs/sandbox/user-custom/), set `numbers.ach_routing` to `322271627` in order to return `true`.

#### Setting balances

You can customize the starting balance and certain other data for a linked account by using a Sandbox [custom user](/docs/sandbox/user-custom/). The easiest way to set a custom balance is to use a test user in the [custom users repo](https://github.com/plaid/sandbox-custom-users/). Using one of the example files, modify the starting balance, then follow the instructions in the repo's README to add the user to your Sandbox.

#### Simulating recurring transfers

In the Sandbox environment, recurring transfers can be simulated by using a `test_clock` object.

A `test_clock` is a mock clock that has a `virtual_time` field, indicating the current timestamp on this test clock. You can attach a `test_clock` to a `recurring_transfer` in Sandbox by providing a `test_clock_id` when calling [`/transfer/recurring/create`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcreate).

You can advance the `virtual_time` on a `test_clock` to a higher value by calling [`/sandbox/transfer/test_clock/advance`](/docs/api/sandbox/#sandboxtransfertest_clockadvance), but you can never decrease the `virtual_time`.

When a test clock is advanced, all active recurring transfers attached to this clock will generate new originations as if the time had elapsed in Production. For instance, assuming a test clock is attached to a weekly recurring transfer, if the test clock is advanced by 14 days, you should see two new transfers being created.

Note that advancing a test clock does not change the status of the transfer objects created by a recurring transfer. Transfers will stay in `pending` status unless you call [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate) to simulate a transfer status update.

##### Sample Sandbox recurring transfer scenarios

1. Create a test clock using [`/sandbox/transfer/test_clock/create`](/docs/api/sandbox/#sandboxtransfertest_clockcreate), with a `virtual_time` of `"2022-11-14T07:00:00-08:00"`, which is 2022-11-14 7AM PST (Monday).
2. Create a weekly recurring transfer on every Tuesday using [`/transfer/recurring/create`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcreate) with the `test_clock_id` returned from step 1. The recurring schedule starts on 2022-11-15, and ends on 2022-11-30.
3. Advance the test clock to `"2022-11-15T23:59:00-08:00"`, which is the end of day 2022-11-15 PST (Tuesday).
4. Since we advanced the test clock to the last minute of Tuesday and the recurring transfer is also scheduled on every Tuesday, we expect to see 1 transfer being created. Inspect the recurring transfer created in step 2 with [`/transfer/recurring/get`](/docs/api/products/transfer/recurring-transfers/#transferrecurringget). Confirm the `transfer_ids` field now has 1 element, and the `status` field is `"active"`.
5. Advance the test clock to `"2022-11-29T23:59:00-08:00"`, which is the end of day 2022-11-29 PST (Tuesday).
6. Inspect the recurring transfer created in step 2 with [`/transfer/recurring/get`](/docs/api/products/transfer/recurring-transfers/#transferrecurringget). Confirm that the `transfer_ids` field now has 3 elements, and the `status` field is now `"expired"`.
7. Advance the test clock to `"2022-12-06T23:59:00-08:00"`, which is the end of day 2022-12-06 PST (Tuesday).
8. Inspect the recurring transfer created in step 2 with [`/transfer/recurring/get`](/docs/api/products/transfer/recurring-transfers/#transferrecurringget). Confirm that the `transfer_ids` field still has 3 elements, and the status field remains `"expired"`.

#### Testing Transfer for Platforms end-user onboarding

In Sandbox, [Transfer for Platforms](/docs/transfer/platform-payments/) customers must use [`/transfer/platform/originator/create`](/docs/api/products/transfer/platform-payments/#transferplatformoriginatorcreate) instead of the [`/partner/customer/create`](/docs/api/partner/#partnercustomercreate) and [`/partner/customer/enable`](/docs/api/partner/#partnercustomerenable) endpoints that they would use in Production.

[`/transfer/platform/originator/create`](/docs/api/products/transfer/platform-payments/#transferplatformoriginatorcreate), [`/transfer/platform/requirement/submit`](/docs/api/products/transfer/platform-payments/#transferplatformrequirementsubmit), and `/transfer/platform/document/submit` will all succeed on valid requests. [`/transfer/platform/person/create`](/docs/api/products/transfer/platform-payments/#transferplatformpersoncreate) will succeed and return a random `person_id`.

In addition to working with `originator_id`s created by [`/transfer/platform/originator/create`](/docs/api/products/transfer/platform-payments/#transferplatformoriginatorcreate), [`/transfer/originator/get`](/docs/api/products/transfer/platform-payments/#transferoriginatorget) will return the following static responses based on which `client_id` is passed as the `originator_client_id`:

| `client_id` | Response |
| --- | --- |
| `111111111111111111111111` | `approved` |
| `222222222222222222222222` | `denied` |
| `333333333333333333333333` | `under_review` |
| `444444444444444444444444` | `more_information_required` |

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
