---
title: "Transfer - Glossary | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/glossary/"
scraped_at: "2026-03-07T22:05:24+00:00"
---

# Glossary

#### Glossary of Transfer-specific terminology

This glossary contains terms specific to the Plaid Transfer product. Most terms in this glossary are industry terms related to funds transfers.

For a general glossary of Plaid API terminology not related to the functionality of the Transfer product specifically, see the [Plaid Glossary](/docs/quickstart/glossary/).

#### Payment rails

##### ACH

ACH (Automated Clearing House) is the primary US electronic network for processing bank-to-bank payments, such as direct deposits, bill payments, and transfers, typically settled in batches. It enables low-cost, reliable fund transfers between accounts, often taking 1–3 business days.

##### Standard ACH

Standard ACH payments are payments sent over the ACH network without expedited processing. These typically complete within 1-3 business days. For transfers submitted via Plaid, the Standard ACH cutoff time is 8:30 PM Eastern Time. It is recommended to submit a request at least 15 minutes before the cutoff time in order to ensure that it will be processed before the cutoff.

##### Same Day ACH

Same Day ACH uses the ACH network to allow eligible bank-to-bank transfers to be processed and settled on the same business day, typically within a few hours. For transfers submitted via Plaid as Same Day ACH, the Same Day ACH cutoff is 3:00 PM Eastern Time and the Standard ACH cutoff is 8:30 PM Eastern Time. It is recommended to submit a request at least 15 minutes before the cutoff time in order to ensure that it will be processed before the cutoff. If a Same Day ACH transfer is processed after the Same Day ACH cutoff but before the Standard ACH cutoff, it will be sent over Standard ACH rails and will not incur Same Day ACH charges; this will apply to both legs of the transfer if applicable.

##### FedNow

FedNow is a real-time payment service launched by the Federal Reserve that enables instant, 24/7/365 bank-to-bank transfers in the U.S., with immediate funds availability to recipients. Like RTP, FedNow only supports `credit` transfers. FedNow payments are irreversible once sent. Within the Plaid Transfer API, payments designated as `rtp` will be routed via either FedNow or RTP depending on availability. To check whether an institution supports real-time payments, use [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget).

##### RTP

RTP, or Real Time Payments, is a U.S. payment rail administered by The Clearing House. RTP payments provide immediate, 24/7/365 settlement of payments between participating banks, with funds available to the recipient within seconds. Like FedNow, RTP only supports `credit` transfers. RTP payments are irreversible once sent. Within the Plaid Transfer API, payments designated as `rtp` will be routed via either FedNow or RTP depending on availability. To check whether an institution supports real-time payments, use [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget).

##### RfP

RfP, or Request for Payment, is a mechanism by which you can initiate the equivalent of a debit funds transfer using the credit-only realtime RTP or FedNow rails. When an RfP is sent to an end user, they will be prompted to approve the RfP, which will then initiate a credit transfer to you for the given amount. To check whether an institution supports RfP, use [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget). Support for RfP within Plaid Transfer is currently in a closed limited beta. To express interest, contact your Account Manager.

##### Wire

Wires, or wire transfers, are electronic funds transfers that move money directly between banks, typically within hours. Wires are only supported for credit transfers. Wires are irreversible once sent. The fee for wires is greater than for other payment methods. To request access to wires, contact your Account Manager. For more details on wires, see [Initiating a payout via wire](/docs/transfer/creating-transfers/#initiating-a-payout-via-wire).

#### ACH terms

##### NACHA

NACHA, also styled as Nacha, is the North American Clearinghouse Association. NACHA manages the ACH network and establishes the rules and procedures around ACH, including what constitutes valid authorization for an ACH transaction and under what circumstances an ACH transaction may be returned or reversed. For more details on these topics, see [What are my responsibilities as an ACH originator?](https://support.plaid.com/hc/en-us/articles/32881513531671-What-are-my-obligations-as-an-ACH-Originator).

##### ODFI

The ODFI, or the Originating Depository Financial Institution, is the bank in an ACH transaction that initiates the transfer. The direction of funds movement is expressed from the perspective of the ODFI -- in the case of a debit transaction, the ODFI is the recipient of the funds, and in the case of a credit transaction, the ODFI is the sender. The ODFI's counterparty bank in the transaction is known as the RDFI, or the Receiving Depository Financial Institution. When you use Plaid Transfer to initiate a transaction, the ODFI is Plaid's Banking Partner.

##### RDFI

The RDFI, or the Receiving Depository Financial Institution, is the bank in an ACH transaction that receives the transfer request from the ODFI. In the case of a debit transaction, the RDFI is the sender of the funds, and in the case of a credit transaction, the RDFI is the recipient. When using Plaid Transfer, the RDFI is typically your end user's bank where they linked their account. In the case of funds movement to or from your Plaid Ledger, initiated from a Plaid UI such as the Dashboard, Transfer API, or automatic balance transfers, the RDFI is your bank. (If moving money to your Ledger via an ACH transfer initiated from your bank, the RDFI is Plaid's Banking Partner and the ODFI is your bank.)

##### Recall

A recall is a request submitted by the Originator to undo an ACH transfer that cannot be reversed or canceled. For more details, see [How do I request a recall?](https://support.plaid.com/hc/en-us/articles/32881719199767-How-are-ACH-Transfer-Investigations-Escalated#h_01JXZT0SCV68P41TP5MSG2E6M3).

##### Return

A return, or ACH Return, occurs when an ACH transaction is rejected or reversed by the receiving bank after it has been submitted, due to issues like insufficient funds, invalid account details, unauthorized debits, or other errors. Returns must be initiated within a specific timeframe (e.g., 2 banking days for most reasons, 60 days for unauthorized consumer debits). Each return comes with a reason code (e.g., R01 = NSF, R03 = No Account). For more details on dealing with returns, see [ACH Returns](/docs/transfer/troubleshooting/#ach-returns). Returns will not occur when using RTP, FedNow, or wire transfers, as these funds transfer methods are irreversible.

##### Reversal

An ACH reversal is a request submitted by the Originator to the ODFI to cancel or correct a previously submitted ACH transaction that has already been processed, typically due to an error such as duplicate payment, wrong amount, or incorrect account number. Reversals must be initiated within 5 banking days of the original settlement date. For details, see [How do I request a reversal?](https://support.plaid.com/hc/en-us/articles/32881719199767-How-are-ACH-Transfer-Investigations-Escalated#h_01JXZT18R8C24G31SB3C1P7PMR). Reversals are not allowed when using RTP, FedNow, or wire transfers, as these funds transfer methods are irreversible.

#### Payment statuses

For more information on payment statuses and their timelines, see the [Transfer status lifecycle](/docs/transfer/reconciling-transfers/#transfer-status-lifecycle).

##### Cancelled

The payment was canceled by the requestor while the payment was in a pending state. No money was moved. This is a terminal state. For more information, see [Canceling transfers](/docs/transfer/creating-transfers/#canceling-transfers).

##### Failed

A failed payment was rejected by either Plaid or the financial institution. No money was moved. This is a terminal state. For more information on handling failures, see [Errors and troubleshooting](/docs/transfer/troubleshooting/).

##### Funds available

Funds from the transfer have been released from hold and applied to the Ledger's available balance. This state is only applicable to ACH debits. This is the terminal state of a successful debit transfer.

##### Pending

A pending payment has not yet been submitted to the payment network and may still be canceled. For information on when payments are submitted to the network, see [ACH processing windows](/docs/transfer/creating-transfers/#ach-processing-windows) and [Initiating a payout via wire](/docs/transfer/creating-transfers/#initiating-a-payout-via-wire).

##### Posted

A posted payment has been submitted to the payment network but has not yet settled. For information on settlement timelines, see the [Transfer Status Lifecycle](/docs/transfer/reconciling-transfers/#transfer-status-lifecycle)

##### Settled

A payment is settled when funds have been transferred between the sending financial institution and the receiving financial institution. Settlement does not guarantee that the institution has released the funds to the recipient; a settled payment may still appear as a pending transaction in the recipient's or sender's bank account. Settled is the normal end state of a successful payment and becomes a terminal state after 60 days. Before the 60 day mark, a settled payment can still transition to the [returned](/docs/transfer/glossary/#returned) state; for more details, see [ACH returns](/docs/transfer/troubleshooting/#ach-returns). For information on settlement timelines, see the [Transfer Status Lifecycle](/docs/transfer/reconciling-transfers/#transfer-status-lifecycle).

##### Returned

A payment has experienced an [ACH return](/docs/transfer/glossary/#return). This is a terminal state. For more details on dealing with returned payments, see [ACH Returns](/docs/transfer/troubleshooting/#ach-returns).

#### Other Transfer terms

##### FBO account

Plaid's FBO ("for benefit of") account is used to hold money used in transfer operations. This is the bank account where you send money to fund your Ledger, and where money is held until it is disbursed to your account; it is the bank account where the money that backs your Ledger is actually stored. The FBO designation indicates that while Plaid is the custodian of the funds, the funds are held "for the benefit of" (i.e., belong to) Plaid's Transfer customers.

##### Funding account

The business bank account belonging to your organization and linked to your Ledger. The funding account is used as the ultimate source of funds for credit transfers and the ultimate recipient of the funds for debit transfers.

##### Ledger

The Ledger, or the Plaid Ledger, is a balance you keep with Plaid for use with the Transfer product. It is a settlement balance that all debit transfers settle into and all credits are paid out from. For details on working with the Ledger, including moving money into and out of the Ledger or maintaining multiple Ledgers, see [Plaid Ledger flow of funds](https://plaid.com/docs/transfer/flow-of-funds/).

##### Originator

The Originator is the party who is requesting a funds transfer. When using Plaid Transfer, you are the originator of all transfers you request unless you are using Transfer for Platforms (beta). For more details on the distinction between Originators and Platforms, see [Originators vs Platforms](/docs/transfer/application/#originators-vs-platforms).

##### Sweep

A sweep is a transfer of money, typically automated or pre-programmed, between two accounts held by the same party to optimize the usage or availability of funds. For example, a brokerage account might automatically sweep money from a cash account to a money market account to earn higher interest while maintaining liquidity. In the context of Plaid Transfer, a sweep refers to any money movement between your Plaid Ledger(s) and your bank account. To learn more about sweeps, see [Sweeping funds to funding accounts](/docs/transfer/creating-transfers/#sweeping-funds-to-funding-accounts).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
