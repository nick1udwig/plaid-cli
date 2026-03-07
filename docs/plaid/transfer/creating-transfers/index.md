---
title: "Transfer - Creating transfers | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/creating-transfers/"
scraped_at: "2026-03-07T22:05:24+00:00"
---

# Linking accounts and creating transfers

#### Link accounts and initiate transfers

Get started with Creating Transfers

[API Reference](/docs/api/products/transfer/)

Don't feel like reading? Check out this video walkthrough of a working Plaid Transfer application!

#### Account Linking

Before initiating a transfer through Plaid, your end users need to link a bank account to your app using [Link](/docs/link/), Plaid's client-side widget. Link will connect the user's account and obtain and verify the account and routing number required to initiate a transfer. Supported account types are debitable checking, savings, or money-market accounts.

See the [Link documentation](/docs/link/) for more details on setting up a Plaid Link session. At a high level, the steps are:

1. Call [`/link/token/create`](/docs/api/link/#linktokencreate), specifying `transfer` in the `products` parameter.
2. Initialize a Link instance using the `link_token` created in the previous step. For more details for your specific platform, see the [Link documentation](/docs/link/). The user will now go through the Link flow.
3. The `onSuccess` callback will indicate the user has completed the Link flow. Call [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to exchange the `public_token` returned by `onSuccess` for an `access_token`. You will also need to obtain and store the `account_id` of the account you wish to transfer funds to or from; this can also be obtained from the `metadata.accounts` field in the `onSuccess` callback, or by calling [`/accounts/get`](/docs/api/accounts/#accountsget). You will use the `account_id` later, when [authorizing the transfer](/docs/transfer/creating-transfers/#authorizing-a-transfer), to specify which account to credit or debit.

Once a Plaid Item is created through Link, you can then immediately process a transfer utilizing that account or initiate the transfer at a later time.

The `metadata.accounts.mask` field and the [`/accounts/get`](/docs/api/accounts/#accountsget) endpoint will contain an account mask, typically the last 2-4 digits of the user's account number. When showing the linked account to your end user in your UI, you should always identify it using the mask rather than displaying or truncating the account number, as some institutions provide special account numbers that do not match what the end user is familiar with.

Most major financial institutions require OAuth connections. Make sure to complete the OAuth security questionnaire at least six weeks ahead of time to ensure your connections are enabled by launch. For more information, see the [OAuth Guide](/docs/link/oauth/).

##### Optimizing the Link UI for Transfer

The following Link configuration options are commonly used with Transfer:

- [**Account select**](/docs/link/customization/#account-select): The "Account Select: Enabled for one account" setting configures the Link UI so that the end user may only select a single account. If you are not using this setting, you will need to build your own UI to let your end user select which linked account they want to use with Transfer. When using [Transfer UI](/docs/transfer/using-transfer-ui/), this setting is mandatory.
- [**Embedded Institution Search**](/docs/link/embedded-institution-search/): This presentation mode shows the Link institution search screen embedded directly into your UI, before the end user has interacted with Link. Embedded Institution Search increases end user uptake of pay-by-bank payment methods and is strongly recommended when implementing Transfer as part of a pay-by-bank use case where multiple different payment methods are supported.

If you are using Transfer in a pay-by-bank use case where multiple payment methods are supported, see [Increasing pay-by-bank adoption](/docs/auth/pay-by-bank-ux/) for recommendations on increasing the percentage of customers who choose to pay by bank.

##### Managing Proof of Authorization

When making debit transactions, it is important to collect and store Proof of Authorization (POA) for NACHA compliance and to dispute return requests if necessary. If not using Transfer UI, you must collect your own POA and store it for at least two years. For more details on best practices for collecting and storing POA, see [NACHA guidance](https://www.nacha.org/system/files/2022-11/WEB_Proof_of_Authorization_Industry_Practices.pdf) and [Authorization requirements for ACH debits](https://support.plaid.com/hc/en-us/articles/32881513531671-What-are-my-obligations-as-an-ACH-Originator#h_01JYHCBZ492B106J7YF45ECKWR). For exact requirements, [refer to the NACHA Operating Rules](https://support.plaid.com/hc/en-us/articles/32881513531671-What-are-my-obligations-as-an-ACH-Originator#h_01JXZS5HMS270QDFV7PKF9NJCC).

While most customers collect and store their own authorization records, as an alternative, you can use Plaid's [Transfer UI](/docs/transfer/using-transfer-ui/) for collecting ACH authorizations.

Transfer UI is a drop-in module for triggering transfer initiations while collecting the necessary authorization required. Simply invoke the UI when the user gets to the confirmation / submission stage of the payment, and Plaid handles the rest.

If an end user claims that a transaction is unauthorized and initiates a return request, their bank may contact Plaid to request a POA. POA requests will only ever be sent to Plaid, and not directly to you. Plaid will then contact you to request any additional information needed. You can provide this information via the [Tasks pane](https://dashboard.plaid.com/transfer/implementation-checklist) in the Dashboard.

##### Importing account and routing numbers

If you are migrating from another payment processor and would like to import known account and routing numbers into Plaid, planning to implement a custom account linking UI, or intend to use wire transfers as a payment rail, contact your account manager about using the [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) endpoint. This endpoint can be used to import previously-verified account and routing numbers and will return an `access_token` and `account_id` for a given account and routing number pair. Items created in this way will always have an authorization decision rationale code of `MIGRATED_ACCOUNT_ITEM`, since Plaid will be unable to assess transfer risk for these Items. This endpoint is not enabled by default; to request access, contact your Plaid Account Manager or [Support](https://dashboard.plaid.com/support).

##### Expanding institution coverage

To see if an institution supports Transfer, use the institution status page in the Dashboard or the [`/institutions/get`](/docs/api/institutions/#institutionsget) endpoint. If an institution is listed as supporting Auth, it will support Transfer.

Transfer supports all of the same flows as Auth, including the optional micro-deposit and database-based flows, which allow you to increase the number of supported institutions and provide pathways for end users who can't or don't want to log in to their institutions. Items created with Same Day micro-deposits, Database Insights, or Database Match will always have an authorization decision rationale code of `MANUALLY_VERIFIED_ITEM`, since Plaid will be unable to assess transfer risk for these Items. For more details about these flows and instructions on implementing them, see [Full Auth coverage](/docs/auth/coverage/).

#### Authorizing a transfer

Before a transfer is created, it must first be authorized. During the authorization step, Plaid's Signal Payment Risk platform runs a series of risk and compliance checks that determine whether the transfer should proceed or be rerouted to a different payment method.

For more details about the specific checks that Signal runs and how to customize them for your business needs, see [Customizing Signal Rules](/docs/transfer/signal-rules/).

To create a transfer authorization, call [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate). You will be required to specify the `access_token` and `account_id` from the account linking step, as well as a `user.legal_name`, the transaction `amount`, the type of the transaction (`debit` or `credit`), and the transaction `network` (`ach`, `same-day-ach`, `rtp`, or `wire`). To request `wire` transfers, contact your Account Manager.

If you are using a [custom set of Signal Rules](/docs/transfer/signal-rules/#creating-additional-custom-rulesets), you should also provide a `ruleset_key`.

For ACH transfers, an `ach_class` is also required. An `idempotency_key` is also strongly recommended to avoid creating duplicate transfers (or being billed for multiple authorizations). If you are a Transfer for Platforms (beta) customer, you will also include an `originator_client_id`. For more details on these parameters, see [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) in the API Reference.

Failure to provide an `idempotency_key` when calling [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) may result in duplicate transfers when retrying requests. It is only safe to omit idempotency keys if you have already built and thoroughly tested your own logic to avoid creating or processing duplicate authorizations. Even if you have, idempotency keys are still strongly recommended.

Plaid will return `'approved'` or `'declined'` as the authorization decision, along with a `decision_rationale` and `authorization_id`. If the transaction is approved, you can proceed to [Initiate the transfer](/docs/transfer/creating-transfers/#initiating-a-transfer).

To avoid blocking transfers, [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) will authorize transfers as `approved` by default in circumstances where Plaid can't accurately predict the risk of return. Always monitor the `decision_rationale` or customize your [Signal Rules](/docs/transfer/signal-rules/) to assess the full risk of a transfer before proceeding to the submission step.

Approved authorizations are valid for 1 hour by default, unless otherwise configured by Plaid support. You may cancel approved authorizations through the [`/transfer/authorization/cancel`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcancel) endpoint if you no longer intend to use the authorization. Denied authorizations will have a `code` and `description` in the `decision_rationale` object that provide additional insight.

##### ACH SEC codes

The ACH Standard Entry Class (SEC) code is represented in the Plaid API by the `ach_class` field and is required for all ACH transfers. SEC codes indicate the type of authorization the originator received in order to submit an ACH payment instruction into the ACH network. Using the incorrect code can result in a failed transfer authorization, ACH returns outside of the expected return window, or, if not remedied, eventual loss of access to Transfer.

Plaid permits only the following SEC codes to be utilized for ACH Transfers:

|  | ACH Debit | ACH Credit |
| --- | --- | --- |
| Consumer accounts | `web`, `tel`, `ppd` | `ppd` |
| Business accounts | `ccd` | `ccd` |

Transfer customers are approved for SEC codes based on the information supplied in the Transfer application. If an unapproved or unsupported SEC code is submitted, the transfer authorization request will fail with the error [`TRANSFER_FORBIDDEN_ACH_CLASS`](/docs/errors/transfer/#transfer_forbidden_ach_class).

###### SEC code definitions:

`ccd` - Corporate Credit or Debit. The transfer moves funds to or from a business bank account

`ppd` - Prearranged Payment or Deposit. The transfer is part of a pre-existing relationship with a consumer. Authorization was obtained from the consumer in person via writing, or through online authorization, or via an electronic document signing, e.g. Docusign. For example language for online authorization, see the [2025 NACHA Operating Rules](https://support.plaid.com/hc/en-us/articles/32881513531671-What-are-my-obligations-as-an-ACH-Originator#h_01JXZS5HMS270QDFV7PKF9NJCC) — Section 2.3.2, Authorization of Entries via Electronic Means. Can be used for credits or debits.

`tel` - Telephone-Initiated Entry. The transfer debits a consumer's bank account. Debit authorization has been received orally over the telephone via a recorded call.

`web` - Internet-Initiated Entry. The transfer debits a consumer’s bank account. Authorization from the consumer is obtained over the Internet (e.g. a web or mobile application). Can be used for single debits or recurring debits.

###### SEC codes and ACH returns

SEC codes determine how long Receiving Depository Financial Institutions (RDFIs, i.e. counterparty banks) have to submit ACH returns for the transfers. ACH debit transfers labeled with consumer SEC codes (`web`, `tel`) permit RDFIs to submit unauthorized returns (R07, R10, R11, etc.) for up to 60 days after the effective date of the transfer, even if the debit was to a business bank account. If a consumer account is debited using a business SEC code (`ccd`), the RDFI can return the transfer for up to 60 days with an R05 return code because is it not proper to debit consumers using a business SEC code.

##### Handling `user_action_required` as a decision rationale code

Sometimes Plaid needs to collect additional user input in order to properly assess transfer risk. The most common scenario is to fix a stale Item. In this case, Plaid returns `user_action_required` as the decision rationale code.

If it is important for your use case to always have an active connection to an end-user's bank account to check the balance, then you should [set the outcome of this rule](/docs/transfer/signal-rules/) to `REROUTE` and prompt the user to relink their account and restore the Item connection before initiating the transfer (see [Repairing Items in `ITEM_LOGIN_REQUIRED`](/docs/transfer/creating-transfers/#repairing-items-in-item_login_required-state)). Otherwise, as long as you still have a valid ACH authorization to debit the end user account, you can use the default rule outcome of `ACCEPT` and proceed to initiate the transfer. Note that Plaid cannot retrieve the account balance on this Item until the Item connection is restored.

Previously, if an Item’s state was `ITEM_LOGIN_REQUIRED`, [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) would respond with `user_action_required` as the authorization decision. As of August 2025, the authorization will be either `approved` or `declined`, based on how you configure your ruleset.

##### Repairing Items in `ITEM_LOGIN_REQUIRED` state

1. Create a new Link token via [`/link/token/create`](/docs/api/link/#linktokencreate). Instead of providing an `access_token`, you should set [`transfer.authorization_id`](/docs/api/link/#link-token-create-request-transfer-authorization-id) in the request.
2. Initialize Link with the `link_token`. Link will automatically guide the user through the necessary steps.
3. You do not need to repeat the token exchange step in Link's `onSuccess` callback as the Item's `access_token` remains unchanged.

After completing the required user action, you can retry the [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) endpoint with the same request body. You can even reuse the same `idempotency_key` as idempotency does not apply to `user_action_required` authorizations.

##### TRANSFER\_LIMIT\_REACHED

If you are seeing transfer authorizations fail with the `TRANSFER_LIMIT_REACHED` rationale code when using Transfer in a new Plaid integration, ensure you have completed the [Transfer Implementation Checklist](https://dashboard.plaid.com/transfer) in the Dashboard in the Dashboard and have received full production access. If your Implementation Checklist has not been approved, your transfer limits will be set at very low levels intended only to allow testing rather than true production usage (maximum $10/transfer, $100/month).

#### Initiating a transfer

After assessing a transfer's risk using the authorization engine and receiving an approved response, you can proceed to submit the transfer for processing.

Pass the `authorization_id`, `access_token`, `account_id`, and a `description` (a string that will be visible to the user) to the [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) endpoint. To make a transfer for less than the amount authorized, provide an optional `amount`; otherwise the transfer will be made for the full authorized amount. The `authorization_id` will also function similarly to an idempotency key; attempting to re-use an `authorization_id` will not create a new transfer, but will return details about the already created transfer. You can also provide the optional field `metadata` to include internal reference numbers or other information to help you reconcile the transfer.

[`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) will return `transfer_id` as well as the `transfer_status`. All transfers begin in a `pending` status. This status changes as the transfer moves through the payment network. Whenever the status is updated, a Transfer event will be triggered.

To learn more about the different stages of the transfer's lifecycle and tracking the status of the transfer, see [event monitoring](/docs/transfer/reconciling-transfers/#event-monitoring).

If submitting a `credit` transfer, you must have enough funds in your Ledger balance to cover the transfer amount.

##### Description field recommendations

The description you send appears on your end-user's bank statement (after the bank prepends your company name and phone number). A clear, consistent value helps users recognize the payment and reduces "unrecognized charge" disputes and ACH returns.

###### Recommendations

- ACH transactions have a 10-character limit; RTP transactions have a 15-character limit. Be concise and abbreviate if necessary.
- Describe the purpose of the transaction, not your company name; your company name will already be present on the line Item.
- Avoid overly granular values (e.g., an invoice ID) that consume space but don't aid recognition. Put variable data in your internal metadata instead.
- Avoid punctuation, non-ASCII characters, case-sensitive values, or emojis, as many banks will only display all-caps ASCII characters, with punctuation stripped.

| Use case | Suggested description | Notes / variants |
| --- | --- | --- |
| Consumer debit (pay-by-bank checkout, bill pay) | `PAYMENT` | Add abbreviation for cycle if helpful: `PAYMENT01`, `PAYMENTJUN` |
| Consumer credit / payout | `PAYOUT` | Refunds can use `REFUND` |
| Loan repayment | `LOANPAY` | Keeps it distinguishable from generic "payment" |
| Rent collection | `RENT` | Optional month tag: `RENTMAY` |
| Savings or investment sweep | `TRANSFER` |  |
| Internal transfer between user accounts | `TRANSFER` |  |
| Returned-transfer retry | `Retry 1` / `Retry 2` | These values are required for R01/R09 retries |

##### Handling errors in transfer creation

If you receive a retryable error code such as a 500 (Internal Server Error) or 429 (Too Many Requests) when calling [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate), you should retry the request; Plaid uses the `authorization_id` as an idempotency key, so you can be guaranteed that retries to [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) will not result in duplicate transfers.

A request to [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) may fail with a 500 error even if the transfer is successfully created. You should not assume that a 500 error response to a [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) call means that the transfer failed. As a source of truth for the transfer status, use [Transfer events](/docs/transfer/reconciling-transfers/).

##### Initiating an Instant Payout via RTP or FedNow

Find out more about using Transfer for Instant Payouts in this 3-minute overview

Initiating an Instant Payout transfer via RTP or FedNow works the same as initiating an ACH transfer. When initiating an Instant Payout transfer, specify `network=rtp` when calling [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate). `rtp` as the network refers to all real time payment rails; Plaid will automatically route between Real Time Payment rail by TCH or FedNow rails as necessary.

Roughly ~70% of accounts in the United States can receive Instant Payouts. If the account is not eligible to receive an Instant Payout, [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) will return an `INVALID_FIELD` error code with an error message that the account is ineligible for Instant Payouts. If you'd like to see if the account is eligible for Instant Payouts before calling [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate), use the [`/transfer/capabilities/get`](/docs/api/products/transfer/account-linking/#transfercapabilitiesget) endpoint.

Only `credit` style transfers (where you are sending funds to a user) can be sent using Instant Payout transfers.

##### Initiating a payout via wire

Wire transfer capabilities are in early availability. To request access to wire transfers, contact your account manager.

To send a wire to a recipient end-user, you must first create an Item for the recipient using the [`/transfer/migrate_account`](/docs/api/products/transfer/account-linking/#transfermigrate_account) endpoint and provide a `wire_routing_number` in addition to the `routing_number`. Wire routing numbers should be collected directly from end users. While the routing number for a wire payment is often the same as the routing number for an ACH payment, some institutions have different wire routing numbers. [Tokenized account numbers](/docs/auth/#tokenized-account-numbers), such as those used at Chase and PNC, cannot be used for wire transfers, so account numbers must be collected directly from end users at institutions that use TANs.

Authorization requests sent with a network of `wire` will only be attempted via wire and will not fall back to any other payment method. Only transfers of type `credit` can be sent using wire transfers.

Wires received by Plaid are processed hourly from 9 AM ET to 7 PM ET. The same-day cutoff time for wire initiation to Plaid is 6:30 PM ET; wires submitted after that time will be processed on the next business day.

The transaction limit for a wire is $999,999.99. Authorization requests sent with an amount greater than $999,999.99 will fail.

Wire transfers may be subject to an incoming wire transfer fee from the recipient's bank, which is typically deducted by the bank from the amount received by the recipient. In the rare instance that a wire is rejected by the receiving bank (RDFI), you will be returned the original wire transfer amount, less a return fee (if any) from the receiving bank. The status of an RDFI-rejected wire will be `RETURNED` with [failure code](https://plaid.com/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-failure-reason-failure-code) `REVERSAL` and a [failure reason description](https://plaid.com/docs/api/products/transfer/reading-transfers/#transfer-get-response-transfer-failure-reason-description) supplied by the banking partner. If a return fee is charged, the amount will be shown in the `wire_return_fee` field. In the Ledger, a new entry will be created with type `transfer`, description `wire return fee`, and an `amount` equal to the fee amount.

#### ACH processing windows

Transfers that are submitted via the [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) endpoint will be submitted in the next processing window. For a Same Day ACH transfer, the cutoff time for the last window of the day is 3:00PM Eastern Time. For a Standard ACH transfer, the cutoff time is 8:30PM Eastern Time. It is recommended to submit a transfer at least 15 minutes before the cutoff time to guarantee that it will be processed before the cutoff.

| Window | Plaid Cutoff | Settlement Date | Approx. Funds Availability to Plaid |
| --- | --- | --- | --- |
| Same Day #1 | 09:35 AM ET | T+0 (1:00 PM ET) | 1:00 PM ET |
| Same Day #2 | 01:50 PM ET | T+0 (5:00 PM ET) | 5:00 PM ET |
| Same Day EOD | 03:00 PM ET | T+0 (~6:00 PM ET) | End of day (~6:00 PM ET) |
| Standard | 08:30 PM ET | T+1 (~9:00 AM ET next day) | Morning of settlement date (~9:00 AM ET) |

Same Day ACH transfers submitted after 3:00PM Eastern Time and before 8:30PM Eastern Time will be automatically updated to be Standard ACH and submitted to the network in that following ACH processing window. This update applies to sweeps as well. This process minimizes the risk of return due to insufficient funds by reducing the delay between Plaid's balance check and the submission of the transfer to the ACH network. This ensures that the settlement time remains earlier than it would have been if the transfer had been submitted in next Same Day window; once the cutoff for Same Day ACH is missed, Standard ACH payments submitted before the cutoff settle to Plaid before the first Same Day settlement time of the next day.

ACH processing windows are active only on banking days. ACH transfers will not be submitted to the network on weekends or bank holidays. Any transfer created when an ACH processing window is closed will automatically be submitted to the network at the opening of the next processing window.

If a transfer is created when an ACH processing window is closed, the account's available balance or status may change between the [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) call and the time when the transfer is submitted to the network. To reduce your risk of ACH returns, wait until the processing window is open before calling [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) and [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate), especially when submitting ACH credits.

See the [flow of funds](/docs/transfer/flow-of-funds/) overview for more details on how, and when, funds move.

#### Bank statement formatting

Each bank has discretion in how they format an ACH, RTP, or FedNow transaction in their bank statements. The most common pattern used is `[Company Name] [Phone Number] [Transfer Description]`.

- `[Company Name]` is the name provided in your Transfer application. This must match a legally registered name for your company.
- `[Phone Number]` is the phone number that you provided in your Transfer application.
- `[Transfer Description]` is the string that you passed into the description field of your [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) request.

To request a change to your phone number or company name, [file a support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).

#### Canceling transfers

To cancel a previously created transfer, use the Dashboard, or call the [`/transfer/cancel`](/docs/api/products/transfer/initiating-transfers/#transfercancel) endpoint with the appropriate `transfer_id`. Note that once Plaid has sent the transfer to the payment network, it cannot be canceled. Practically speaking, this means that Instant Payouts via RTP/FedNow cannot be canceled, as these transfers are immediately sent to the payment network. You can check use the `cancellable` property found via [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) to determine if a transfer is eligible for cancelation.

If the transfer is not `cancellable`, you may be able to undo it in another way. For undoing debit transfers, see [refunds](/docs/transfer/refunds/). For information on undoing an ACH credit transfer, see [Escalating issues with ACH transfers](/docs/transfer/troubleshooting/#escalating-issues-with-ach-transfers).

#### Sweeping funds to funding accounts

There are two types of bank accounts in the Plaid Transfer system.

The first is a consumer checking, savings, or cash management account connected via Plaid Link, where you are pulling money from or issuing a payout to. In the Transfer API, this account is represented by an `access_token` and `account_id` of a Plaid Item. A "transfer" is an intent to move money to or from this account.

The second type of account involved is your own business checking account, which is linked to your Plaid Ledger. This account is configured with Plaid Transfer during your onboarding by using the details provided in your application. A "sweep" pushes money into, or pulls money from, a business checking account. For example, funding your Plaid Ledger account by calling [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit) will trigger a sweep event.

To learn about all the different ways you can trigger a sweep event, see [Moving money between Plaid Ledger and your bank account](/docs/transfer/flow-of-funds/#moving-money-between-plaid-ledger-and-your-bank-account).

Sweeps can be observed via the [`/transfer/sweep/list`](/docs/api/products/transfer/reading-transfers/#transfersweeplist) and [`/transfer/sweep/get`](/docs/api/products/transfer/reading-transfers/#transfersweepget) endpoints. [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) will also return all events types, including sweep events.

#### Transfer limits

When you sign up with Plaid Transfer, you will be assigned transfer limits, which are placed on your authorization usage. These limits are initially assigned based on the volume expectations you provide in your transfer application. When you successfully create an authorization using [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate), the amount authorized will be counted against your limits. Authorized amounts that aren’t used will stop counting towards your limits if the authorization expires or is canceled (via [`/transfer/authorization/cancel`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcancel)).

Any authorization you attempt to create over your limits will be automatically declined.

| Limit | Debit | Credit |
| --- | --- | --- |
| Single Transfer | The maximum amount of a single debit transfer authorization | The maximum amount of a single credit transfer authorization |
| Daily | The maximum amount of debit transfers you may authorize in a calendar day | The maximum amount of credit transfers you may authorize in a calendar day |
| Monthly | The maximum amount of debit transfers you may authorize in a calendar month | The maximum amount of credit transfers you may authorize in a calendar month |

Daily limits refresh every day and monthly limits refresh on the first day of the month at 12:00 AM EST.

##### Limit monitoring

You can view your current limits on the “Account Details” page in the [Transfer Dashboard](https://dashboard.plaid.com/transfer). You can monitor your usage against your daily and monthly limits via the Account Information widget on the Overview page. You can also monitor your authorization limits and usage through the Transfer APIs.

The [`/transfer/configuration/get`](/docs/api/products/transfer/metrics/#transferconfigurationget) endpoint returns your configurations for each of the transfer limits. The [`/transfer/metrics/get`](/docs/api/products/transfer/metrics/#transfermetricsget) endpoint contains information about your daily and monthly authorization usage for credit and debit transfers.

Plaid will also send you automatic email notifications when your utilization is approaching your limits. If your daily utilization exceeds 85% of your daily limits or your monthly utilization exceeds 80% of your monthly limits, you will receive automated email alerts. These alerts will be sent to your ACH, Technical, and Billing contacts. You can configure those contacts via the [Company Profile](https://dashboard.plaid.com/settings/company/compliance?tab=companyProfile) page on the Plaid Dashboard.

Any call to [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) that will cause you to exceed your limits will return the decision `"declined"` with the decision rationale code set to `TRANSFER_LIMIT_REACHED`. The rationale description identifies which specific limit has been reached.

##### Requesting limit changes

If you need to change your limits for any reason, you can request changes via the [Account Details](https://dashboard.plaid.com/transfer/account-details) page in the Plaid Dashboard.

#### Peer to peer transfers

Plaid Transfer does not support peer to peer transfers or transfers between two accounts held by the same person. For these use cases, see [Auth](/docs/auth/).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
