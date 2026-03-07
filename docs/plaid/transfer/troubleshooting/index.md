---
title: "Transfer - Errors and returns | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/troubleshooting/"
scraped_at: "2026-03-07T22:05:26+00:00"
---

# Errors and returns

#### Troubleshooting common errors and ACH returns

#### Overview

ACH returns and Transfer errors can both require investigation and troubleshooting.

An ACH return means that your attempted ACH transfer was returned. This typically not due to a Plaid issue and instead is caused by a problem with the ACH transfer itself, such as alleged fraud, insufficient funds, or a closed account. For more details, see [ACH Returns](/docs/transfer/troubleshooting/#ach-returns).

A Transfer error is an error from the Plaid Transfer API. Transfer errors are related to your interaction with the Plaid API. Causes may include bank API downtime, exceeding Plaid's limits for transfer activity, or attempting to use a Plaid Transfer feature you haven't been enabled for. For more details, see [Transfer errors](/docs/transfer/troubleshooting/#transfer-errors).

#### ACH returns

After an ACH payment has been posted, it can be [returned](https://plaid.com/resources/ach/ach-return/) and clawed back from your account for a number of reasons. For example, a bank could determine that there are insufficient funds to complete the transfer, or a consumer can contact their bank and flag the transaction as unauthorized. If a debit transfer results in an ACH return, the funds will be automatically clawed back in a return sweep. If a credit transfer experiences an ACH return, the funds will be returned to your Ledger available balance.

Most ACH returns (including all R01 "insufficient funds" returns, which are by far the most common return type) must occur within 2 banking days of the transaction entering the `settled` status. For example, if a transaction was settled on Monday, then after close of business on Wednesday, it would be safe from an R01 return. However, for ACH returns in which the consumer alleges that the transaction was unauthorized, including return codes such as R10 or R11, the window is 60 calendar days.

Note that an ACH return is distinct from an ACH reversal. An ACH reversal is performed only to correct an erroneous transfer and will only be done at the request of the originator (i.e. you). For details, see [Escalating issues with ACH transfers](/docs/transfer/troubleshooting/#escalating-issues-with-ach-transfers).

If an ACH transfer is returned, its status will be `returned`. All returned ACH transactions will have an ACH return code. By reading the code, you can troubleshoot returned transactions.

For more detailed troubleshooting information, see [Common ACH Return Codes in the Help Center](https://support.plaid.com/hc/en-us/articles/32881797799575-What-are-the-Common-ACH-Return-Codes).

For a full list of ACH return codes, see [ACH Return Codes](/docs/errors/transfer/#ach-return-codes).

| Return Code | Description | Notes | Troubleshooting |
| --- | --- | --- | --- |
| R01 | Insufficient funds | Available balance is not sufficient to cover the dollar amount of the debit entry. | If you are submitting the transfer on a Friday, we recommend using Same-Day ACH to decrease the likelihood of the user's account balance dipping below the transfer amount over the weekend. |
| R03 | No account or unable to locate account | The account number structure is valid, but the account number does not correspond to an existing account, or the name on the transaction does not match the name on the account. | Prompt the user to link another account via Plaid Link and use this newly linked account to create the transfer. |
| R09 | Uncollected funds | Available balance is sufficient, but the collected balance is not sufficient to cover the entry | Retry the transaction a few days later after any large "holds" on the user's account have been released. |
| R10 | Customer advises not authorized | Member advises not authorized, notice not provided, improper source document, or amount of entry not accurately obtained from source document. | Make sure you add proper WEB debit authorization language to your "pay with bank" UX to protect against exposure of returns of over 60 days. Unlike other ACH return codes, you cannot resubmit a payment after encountering an R10 error. |

A returned transfer can be retried up to 2 times. When you try to reprocess a returned transfer, the `description` field in your [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) request must be `"Retry 1"` or `"Retry 2"` to indicate that it's a retry of a previously returned transfer. Note that retries are only allowed for transfers returned as R01 or R09.

##### ACH return categories

ACH return rates are broken into three categories.

*Administrative*: The transfer could not be processed due to a technical error with the account. Common return codes in this category include R02 (account closed), R03 (no account), and R04 (invalid account number).

*Unauthorized*: The transfer was reported as unauthorized. Common return codes in this category include R10 (customer advises originator is not authorized to debit account), R11 (customer advises originator is authorized to debit account, but amount is incorrect), and R07 (customer advises authorization to debit account was revoked).

*Other*: All other returns. The most common code in this category is R01 (insufficient funds).

###### Return rate limits

You can view your current return rates by category on the Transfer Dashboard. Plaid's maximum limits for ACH return rates are:

- Overall: 12%
- Administrative: 2.5%
- Unauthorized: 0.4%

For more details, see [ACH Debit Return Rate Thresholds](https://support.plaid.com/hc/en-us/articles/32881513531671-What-are-my-obligations-as-an-ACH-Originator#h_01JYHCEY1RT625QMSVWQPDEVPB).

If you exceed these limits in a 30 day period, you will be contacted by Plaid with a warning email. Failure to reduce return rates below these limits within three months may result in a fine. If excessive return rates remain unaddressed for an extended period of time, your access to Plaid Transfer may be revoked. If you are having trouble reducing your return rates below these limits, contact your Plaid Account Manager for help.

#### RTP failures

Plaid performs various checks within [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) to fail upfront and as early as possible if the RTP or FedNow transfer won't succeed. However, due to the asynchronous nature of the RTP and FedNow payment networks, it is still possible for transfers to fail after authorization due to a rejection by the receiving bank or an issue in the payment network. Failures in the real-time networks are rare (<1%), and often indicate that the consumer bank account itself is in an unusable state.

If the RTP transfer failed, the status will be `failed` and an error code will be present in the `return_code` transfer property. For the following failure codes, further attempts to use this bank account for payouts via the real-time network are unlikely to succeed. Instead, we recommend that you prompt your user to link a different bank account via Plaid Link.

For a full list of RTP failure codes, see [RTP / RfP error codes](/docs/errors/transfer/#rtprfp-error-codes).

| Return Code | Description | Notes |
| --- | --- | --- |
| AC03 | Creditor account is invalid | Account number is not recognized as valid by the receiving bank. |
| AC04 | Account closed | The account is closed and no longer active. |
| AC06 | Account is blocked | Receiving bank has placed a block on the account, for various operational or legal reasons. |
| E997 | Timeout clock has expired | The timeout for processing this transaction was exceeded in the FedNow network. You can safely retry the payment with a different idempotency key. |

#### Transfer errors

For error codes specific to transfer and accompanying troubleshooting steps, including for common errors [`TRANSFER_NETWORK_LIMIT_EXCEEDED`](/docs/errors/transfer/#transfer_network_limit_exceeded), [`TRANSFER_ACCOUNT_BLOCKED`](/docs/errors/transfer/#transfer_account_blocked), and [`TRANSFER_FORBIDDEN_ACH_CLASS`](/docs/errors/transfer/#transfer_forbidden_ach_class) see [Transfer error codes](/docs/errors/transfer/).

##### Handling 500 errors in transfer creation

You may occasionally encounter a 500 HTTP error when creating a transfer via [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) or [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate). In this case, it's possible that the authorization or transfer was created successfully, but the response was not received. It is highly recommended to use the `idempotency_key` field in [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) so that retries in this scenario are safe to perform and will not result in duplicate authorizations. Retrying [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) with the same `authorization_id` is guaranteed to only ever produce a single transfer.

Whether you retry the request to [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) or not, it's possible that the first request succeeded, and so your integration should be prepared to consume events from [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) that you may not have seen the `transfer_id` for previously. It is recommended you have a recovery process in place, either by logging these instances for manual review and reconciliation, or automatically reconciling by pulling the transfer details via [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) and using the `transfer.metadata` property to utilize any client-side data you attached to the transfer for reference.

Another way to determine if a transfer creation succeeded after receiving a 500 error from [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) is to follow up with a call to [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) using the `authorization_id` you received from [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate). If the transfer was actually created, you will receive a successful response containing the transfer details. If it wasn't, you will receive a 404 HTTP error with the `NOT_FOUND` error code.

#### Escalating issues with ACH transfers

In the event of a problem with an ACH transfer, you can escalate to Plaid. Typical reasons to request an escalation include:

- You accidentally submitted an erroneous transfer
- Your end-customer claims that they did not receive funds you sent them
- Your end-customer claims that you made an unauthorized debit from their account

Available escalation types include reversals, recalls, trace requests, and written statement of unauthorized debit (WSUD) requests.

To learn more about escalations, including which escalation type is appropriate for your situation and how to submit an escalation, see [How are ACH Transfer Investigations Escalated?](https://support.plaid.com/hc/en-us/articles/32881719199767-How-are-ACH-Transfer-Investigations-Escalated)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
