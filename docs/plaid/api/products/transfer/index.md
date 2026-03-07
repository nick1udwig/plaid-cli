---
title: "API - Transfer | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/transfer/"
scraped_at: "2026-03-07T22:04:20+00:00"
---

# Transfer

#### API reference for Transfer endpoints and webhooks

For how-to guidance, see the [Transfer documentation](/docs/transfer/).

| Initiating Transfers |  |
| --- | --- |
| [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) | Create a transfer authorization |
| [`/transfer/authorization/cancel`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcancel) | Cancel a transfer authorization |
| [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) | Create a transfer |
| [`/transfer/cancel`](/docs/api/products/transfer/initiating-transfers/#transfercancel) | Cancel a transfer |

| Reading Transfers |  |
| --- | --- |
| [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) | Retrieve information about a transfer |
| [`/transfer/list`](/docs/api/products/transfer/reading-transfers/#transferlist) | Retrieve a list of transfers and their statuses |
| [`/transfer/event/list`](/docs/api/products/transfer/reading-transfers/#transfereventlist) | Retrieve a list of transfer events |
| [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) | Sync transfer events |
| [`/transfer/sweep/get`](/docs/api/products/transfer/reading-transfers/#transfersweepget) | Retrieve information about a sweep |
| [`/transfer/sweep/list`](/docs/api/products/transfer/reading-transfers/#transfersweeplist) | Retrieve a list of sweeps |

| Account Linking |  |
| --- | --- |
| [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget) | Determine RTP eligibility for a Plaid Item |
| [`/transfer/intent/create`](/docs/api/products/transfer/account-linking/#transferintentcreate) | Create a transfer intent and invoke Transfer UI (Transfer UI only) |
| [`/transfer/intent/get`](/docs/api/products/transfer/account-linking/#transferintentget) | Retrieve information about a transfer intent (Transfer UI only) |
| [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) | Create an Item to use with Transfer from known account and routing numbers |

| Recurring Transfers |  |
| --- | --- |
| [`/transfer/recurring/create`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcreate) | Create a recurring transfer |
| [`/transfer/recurring/cancel`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcancel) | Cancel a recurring transfer |
| [`/transfer/recurring/get`](/docs/api/products/transfer/recurring-transfers/#transferrecurringget) | Retrieve information about a recurring transfer |
| [`/transfer/recurring/list`](/docs/api/products/transfer/recurring-transfers/#transferrecurringlist) | Retrieve a list of recurring transfers |

| Refunds |  |
| --- | --- |
| [`/transfer/refund/create`](/docs/api/products/transfer/refunds/#transferrefundcreate) | Create a refund for a transfer |
| [`/transfer/refund/cancel`](/docs/api/products/transfer/refunds/#transferrefundcancel) | Cancel a refund |
| [`/transfer/refund/get`](/docs/api/products/transfer/refunds/#transferrefundget) | Retrieve information about a refund |

| Transfer for Platforms |  |
| --- | --- |
| [`/transfer/platform/originator/create`](/docs/api/products/transfer/platform-payments/#transferplatformoriginatorcreate) | Pass transfer specific onboarding info for the originator |
| [`/transfer/platform/person/create`](/docs/api/products/transfer/platform-payments/#transferplatformpersoncreate) | Create each individual who is a beneficial owner or control person of the business |
| [`/transfer/platform/requirement/submit`](/docs/api/products/transfer/platform-payments/#transferplatformrequirementsubmit) | Pass additional data Plaid needs to make an onboarding decision for the originator |
| [`/transfer/platform/document/submit`](/docs/api/products/transfer/platform-payments/#transferplatformdocumentsubmit) | Submit documents Plaid needs to verify information about the originator |
| [`/transfer/originator/get`](/docs/api/products/transfer/platform-payments/#transferoriginatorget) | Get the status of an originator's onboarding |
| [`/transfer/originator/list`](/docs/api/products/transfer/platform-payments/#transferoriginatorlist) | Get the status of all originators' onboarding |
| [`/transfer/originator/funding_account/create`](/docs/api/products/transfer/platform-payments/#transferoriginatorfunding_accountcreate) | Create a new funding account for an originator |

| Plaid Ledger |  |
| --- | --- |
| [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit) | Deposit funds into a ledger balance held with Plaid |
| [`/transfer/ledger/distribute`](/docs/api/products/transfer/ledger/#transferledgerdistribute) | Move available balance between platform and its originator |
| [`/transfer/ledger/get`](/docs/api/products/transfer/ledger/#transferledgerget) | Retrieve information about the ledger balance held with Plaid |
| [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw) | Withdraw funds from a ledger balance held with Plaid |
| [`/transfer/ledger/event/list`](/docs/api/products/transfer/ledger/#transferledgereventlist) | Retrieve a list of ledger balance events |

| Program Metrics |  |
| --- | --- |
| [`/transfer/metrics/get`](/docs/api/products/transfer/metrics/#transfermetricsget) | Get transfer product usage metrics |
| [`/transfer/configuration/get`](/docs/api/products/transfer/metrics/#transferconfigurationget) | Get transfer product configuration |

| Sandbox |  |
| --- | --- |
| [`/sandbox/transfer/simulate`](/docs/api/sandbox/#sandboxtransfersimulate) | Simulate a transfer event |
| [`/sandbox/transfer/refund/simulate`](/docs/api/sandbox/#sandboxtransferrefundsimulate) | Simulate a refund event |
| [`/sandbox/transfer/sweep/simulate`](/docs/api/sandbox/#sandboxtransfersweepsimulate) | Simulate creating a sweep |
| [`/sandbox/transfer/fire_webhook`](/docs/api/sandbox/#sandboxtransferfire_webhook) | Simulate a transfer webhook |
| [`/sandbox/transfer/ledger/deposit/simulate`](/docs/api/sandbox/#sandboxtransferledgerdepositsimulate) | Simulate a deposit sweep event |
| [`/sandbox/transfer/ledger/simulate_available`](/docs/api/sandbox/#sandboxtransferledgersimulate_available) | Simulate converting pending balance into available balance |
| [`/sandbox/transfer/ledger/withdraw/simulate`](/docs/api/sandbox/#sandboxtransferledgerwithdrawsimulate) | Simulate a withdrawal sweep event |
| [`/sandbox/transfer/test_clock/create`](/docs/api/sandbox/#sandboxtransfertest_clockcreate) | Create a test clock |
| [`/sandbox/transfer/test_clock/advance`](/docs/api/sandbox/#sandboxtransfertest_clockadvance) | Advance a test clock |
| [`/sandbox/transfer/test_clock/get`](/docs/api/sandbox/#sandboxtransfertest_clockget) | Retrieve information about a test clock |
| [`/sandbox/transfer/test_clock/list`](/docs/api/sandbox/#sandboxtransfertest_clocklist) | Retrieve a list of test clocks |

| Webhooks |  |
| --- | --- |
| [`TRANSFER_EVENTS_UPDATE`](/docs/api/products/transfer/reading-transfers/#transfer_events_update) | New transfer events available |
| [`RECURRING_CANCELLED`](/docs/api/products/transfer/recurring-transfers/#recurring_cancelled) | A recurring transfer has been cancelled by Plaid |
| [`RECURRING_NEW_TRANSFER`](/docs/api/products/transfer/recurring-transfers/#recurring_new_transfer) | A new transfer of a recurring transfer has been originated |
| [`RECURRING_TRANSFER_SKIPPED`](/docs/api/products/transfer/recurring-transfers/#recurring_transfer_skipped) | An instance of a scheduled recurring transfer could not be created |

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
