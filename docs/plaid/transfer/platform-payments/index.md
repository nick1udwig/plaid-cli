---
title: "Transfer - Transfer for Platforms | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/platform-payments/"
scraped_at: "2026-03-07T22:05:25+00:00"
---

# Transfer for Platforms

#### Learn how to use Transfer as a Payments Platform

Get started with Transfer for Platforms

[API Reference](/docs/api/products/transfer/)[Quickstart](/docs/quickstart/)

#### Overview

Transfer for Platforms helps vertical SaaS providers unlock new revenue streams by embedding bill payments - whether collecting funds for items such as rent, tuition, or medical bills, or disbursing payments to customers’ vendors. This results in faster onboarding, higher ACH adoption, and fewer returns.

Transfer for Platforms is for companies who would like to incorporate Plaid Transfer into the financial platforms that they create and then provide to their end-customers. Plaid refers to these companies as Reseller Partners. Supported use cases include payment collection for ad hoc or recurring services provided to individuals or businesses in healthcare, education, home services, real estate, technology, or other professional services industries. Platforms serving customers in these industries can also use Transfer for Platforms to pay out to their customers’ vendors.

![Image showing two types of transfer customers. Originators move money directly to customers. Platforms involve an intermediary business.](/assets/img/docs/transfer/platform_structure_overview.png)

Transfer for Platforms is in beta. If you are interested in Transfer for Platforms, [contact sales](https://plaid.com/contact) or your Plaid account manager.

Financial services use cases, such as crowdfunding, peer-to-peer payments, or lending, are not currently eligible to join the Transfer for Platforms beta program.

If you are not a Platform, see [Transfer](/docs/transfer/) for Plaid's general transfer documentation. If you're unsure whether you're a Platform, see [Originators vs Platforms](/docs/transfer/application/#originators-vs-platforms).

##### Benefits of Transfer for Platforms

- Instant, API-first onboarding: The entire onboarding flow stays within your product and customers don’t have to enroll through a separate portal.
- Operational tools built for scale: Transfer for Platforms supports sub-ledgers for each customer, configurable hold times and payout speeds, built-in risk tooling, and integrated fee collection.
- Visibility and control via the Plaid Dashboard: Track onboarding progress, payment activity, and performance indicators from a single dashboard, with the ability to view all data at the customer level.

#### Application process

To apply for Transfer for Platforms, submit a Transfer application through the [Plaid Dashboard](https://dashboard.plaid.com/settings/team/products). After review, Plaid will reach out to collect any additional information and/or documents needed to verify your use case.

As part of this review, a default configuration will be established for your payments program that is applied to each customer after risk approval. This configuration includes:

- Limits (per transaction, daily, monthly)
- Hold times for funds from debit transfers (if applicable)
- ACH SEC codes (how ACH debits are authorized, if applicable)
- Payment methods

You can [request changes to limits and hold times](/docs/transfer/dashboard/#limits-reserves-and-hold-time) through the Plaid Dashboard.

#### Integration process

Transfer for Platforms is built on top of Transfer. It's recommended you review the Transfer [integration steps](/docs/transfer/) and use the Transfer for Platforms documentation as a supplement.

At a high level, the instructions for integrating with Transfer for Platforms are the same as for integrating with Transfer, with two exceptions:

- You will need to onboard your customers (also known as originators).
- When calling most endpoints in the Transfer API, you will specify an `originator_client_id`.

You must use your end-customer's `client_id` and `production_secret` when calling [`/link/token/create`](/docs/api/link/#linktokencreate) and [`/item/public_token/exchange`](/docs/api/items/#itempublic_tokenexchange) to link end-user Items for Transfer. You should use your own `client_id` and `secret` when calling other Transfer endpoints.

###### Customers vs Users

**End-customers** (also known as **originators** or simply **customers**) are your customers, on whose behalf funds are being transferred. These are identified by an `originator_client_id`.

**Users** or (**end-users**) are the customers of your end-customers; the people who are paying in to Transfer. These are associated with an `access_token`.

#### Onboarding end-customers

To set up a customer for Transfer, you will first need to submit them to Plaid's risk review process. You will submit basic KYB/KYC details about the customer that you already collect as part of your onboarding flow.

Plaid then runs automated checks on these customer details and will let you know if any additional documentation is required to verify the business. Once a customer is approved, they are automatically enabled for processing, and you are notified via webhook.

The process for onboarding an end-customer is as follows:

##### Step 1: Ensure customers accept Plaid's Terms of Service

As part of your onboarding process, ensure your customers accept Plaid Transfer’s Terms of Service.

In addition to your own ToS, include a link to [Plaid’s Terms of Service for Transfer End-customers](https://plaid.com/documents/plaid_end_customer_payments_msa.pdf) and collect acknowledgement of these terms. Make sure to store the IP address your customer was using when they accepted the ToS as well a timestamp, as you will need this information in Step 3.

##### Step 2: Create an end-customer

Call [`/partner/customer/create`](/docs/api/partner/#partnercustomercreate) to create a new client ID for the customer. Make sure to specify `auth` (**not** `transfer`) in the products array.

You will receive an `end_customer.client_id` as part of the response object. Make sure to save this value -- you will use this value as the `originator_client_id` to identify this end-customer in future calls.

Next, call [`/partner/customer/enable`](/docs/api/partner/#partnercustomerenable) to enable this customer in Production. Make sure to store the `production_secret` that is returned, as you will need this when creating a Link token for your customer's end-users.

After your customer has been enabled, you may proceed to the next step.

##### Step 3: Submit required information for risk approval

To comply with Know Your Customer (KYC) obligations, Plaid must collect, verify, and maintain information about your customers. Plaid will collect information about their business entities, their Control Persons (CPs), and qualifying Ultimate Beneficial Owners (UBOs). You can use the Plaid API to submit this required information for automated verification attempts. If the API cannot affirmatively verify the validity of the business entity, CP, or UBO, then Plaid will require additional documentary evidence to complete the validation.

First, call [`/transfer/platform/originator/create`](/docs/api/products/transfer/platform-payments/#transferplatformoriginatorcreate) to pass Transfer-specific onboarding info about the originator. This will include metadata related to your customer's accepting Plaid's terms of service.

Next, call [`/transfer/platform/person/create`](/docs/api/products/transfer/platform-payments/#transferplatformpersoncreate) for each individual who is a Control Person or Ultimate Beneficial Owner. The endpoint will return a unique `person_id` for each person. Make sure to store these IDs, as they will be needed to identify requirements associated with these individuals in the future.

###### Control Persons and Ultimate Beneficial Owners

A Control Person is an individual that has significant responsibility to control, manage, or direct a legal entity. Control Persons may include an executive officer or senior manager (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating Officer, Managing Member, General Partner, President, Vice President, or Treasurer) or any other individual who regularly performs similar functions.

An Ultimate Beneficial Owner is any individual who owns 25% or more of the company. It is generally expected that there will be one or more ultimate beneficial owners for each end-customer you submit for onboarding review.

At least one Control Person must be submitted for every end-customer. All beneficial owners with 25%+ ownership are required to be identified.

For business entities designated as sole proprietorships for the `business_org_type`, only one person needs to be created and associated with the business entity. This person should be indicated as the `CONTROL_PERSON_AND_BENEFICIAL_OWNER` for the `person_relationship` value. In this situation, is not necessary to also create a Beneficial Owner.

##### Step 4: Supply additional information

At this point, if you call [`/transfer/originator/get`](/docs/api/products/transfer/platform-payments/#transferoriginatorget), your customer’s onboarding status will be `more_information_required`, and the `outstanding_requirements` property will include a list of information that you need to submit.

You can submit this additional data by calling [`/transfer/platform/requirement/submit`](/docs/api/products/transfer/platform-payments/#transferplatformrequirementsubmit). You may call this endpoint multiple times -- it is not necessary to submit the additional data all at once.

###### Information collected to verify a business entity:

- Legal name
- EIN or Tax ID
- Legal address
- Industry
- Organization type
- Website
- Product Description
- Bank account (business bank account owned by the customer)

###### Information collected to verify a person:

- Legal name
- Legal address
- Identification number (SSN or ITIN)
- Date of birth
- Email
- Phone
- Relationship
- Percent ownership (if beneficial owner)
- Title (if Control Person)

##### Step 5: Wait for updates on risk approval status

Once you’ve submitted all requested information, your customer should move into the `under_review` status. You can then either listen for the `PLATFORM_ONBOARDING_UPDATE` webhook or poll [`/transfer/originator/get`](/docs/api/products/transfer/platform-payments/#transferoriginatorget) until the `status` is a value other than `under_review`. This process typically takes 15 minutes.

If the new status is `approved`, your end-customer is enabled for processing. If the status is `denied`, reach out to your Plaid account manager to determine if the customer can be re-submitted.

If the status is `more_information_required`, the response will include an `outstanding_requirements` array that specifies additional documents that must be submitted.

##### Step 6: (If needed) Submit additional documentation

If Plaid requests additional documentation, you will need to call `/transfer/platform/document/submit` for each document that is required.

After you have submitted all required documents, you can listen for the `PLATFORM_ONBOARDING_UPDATE` webhook or poll [`/transfer/originator/get`](/docs/api/products/transfer/platform-payments/#transferoriginatorget) as described in step 5 to listen for the new status of the customer. Note that documents submitted during this step are reviewed manually, so it can take up to two business days for the status to be updated.

###### Documents that may be required to verify a business entity:

- **Name:** Either the Articles of Incorporation for the company or a Recent Certificate of Good Standing. For companies that have undergone a name change, the original Secretary of State formation documents must be appended with the name change amendment.
- **Address:** A document that proves the address provided belongs to the business. May be a utility bill, bank statement, or property tax bill. Must be from the past 3 months. For utility bills, the service address must match the submitted address for the business if the bill was mailed to an alternative address (e.g. a P.O. box). If the statement or bill is in the DBA name of the business, append a Proof of DBA document to the address validation document file.
- **EIN or Tax ID:** IRS-issued EIN Letter (Form CP 575) or a completed business tax filing. Note that SS-4 documents are not acceptable.

###### Documents that may be required to verify a person:

- **Name or date of birth:** Driver’s license, passport, military ID, or other government issued ID
- **Identification number:** SSN or ITIN card
- **Address:** A document that proves that the residential address belongs to the associated person. May be a utility bill, bank statement, or government ID. Must be from the past 3 months.

#### Payment collection

This section discusses moving money for your customers' accounts. You also have your own Plaid Ledger, which is associated with your `client_id`; to learn about funds transfers between your own Ledger and funding account, see [Moving money between Plaid Ledger and your bank account](/docs/transfer/flow-of-funds/#moving-money-between-plaid-ledger-and-your-bank-account).

![Image showing two types of transfer customers. Originators move money directly to customers. Platforms involve an intermediary business.](/assets/img/docs/transfer/platform_payment_collection.png)

Flow of funds overview for payment collection

A [Plaid Ledger](/docs/transfer/flow-of-funds/) will automatically be created for each of your customers. Each Ledger is linked to a funding account, which is a regular business checking account. For your customers, the account details for the account they would like to use as a funding account are collected during onboarding associated with their `originator_client_id`.

##### Collecting payment from end-users

To collect payment for a bill, membership, or subscription from a customer’s end-user, initiate a debit transfer by calling [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) with a `type` of `debit`, followed by a call to [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate).

Funds from ACH debits are initially placed in the pending balance of the customer's Ledger. Once the hold period (typically 2-3 business days) elapses, the funds will shift to the customer's available balance at or before 3pm EST on the final release date. Once the funds are in your customer's available balance, they can be withdrawn to the customer’s linked funding account.

Refunds can be issued for debit transfers when necessary as long as the customer has sufficient funds in their available balance. For more details, see [Refunds](/docs/transfer/refunds/).

Once a debit transfer is created, you can find the expected date when funds will be made available (assuming the transfer is not returned during the hold period) in the `expected_funds_available_date` value in the [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) response object. This date is also visible in the customer’s Transfer Dashboard. If the transfer does get returned before this date, this field will be null.

To withdraw funds from the Plaid Ledger to the customer's linked funding account, call [`/transfer/ledger/withdraw`](/docs/api/products/transfer/ledger/#transferledgerwithdraw), specifying the `originator_client_id`, `amount`, and payment method (e.g. `ach`, `same-day-ach`, `rtp`). This will create a credit sweep and move the funds from your customer's Ledger balance to their funding account during the next processing window (or instantly if done via RTP).

##### Returns on ACH debits

Returns on ACH debits will decrease a customer’s Ledger balance. If funds are still in the pending balance when the return occurs, the pending balance will decrease. If the funds have transitioned to the available balance, the available balance will decrease.

###### Negative customer balances

Returns on ACH debits can result in a negative available balance if there are insufficient funds in the customer’s available Ledger balance at the time of the return. This occurs most commonly due to unauthorized returns from debits to consumers, which have a 60-calendar day return window.

To avoid negative balances, you can set minimum and maximum available balance thresholds in the Dashboard for each customer. Setting a minimum balance greater than $0 helps ensure there are enough funds to cover refunds and returns.

If a customer’s available balance becomes negative, Plaid will automatically debit the customer’s linked funding account to restore the balance. Ensure that your customers do not have a block on their accounts for these auto-debits. See [white-listing Plaid’s originator ID](/docs/transfer/flow-of-funds/#pushing-funds-to-your-ledger-via-ach-rtp-or-wire).

Although Plaid will auto-debit customers’ linked funding accounts in the event of a negative balance, this auto-debit can fail (e.g. due to insufficient funds or a closed bank account). If this occurs, Plaid will reach out to you. If an end-customer’s balance is persistently negative, Plaid may place limits on their ability to use Transfer.

You are liable for your customers’ negative balances. You should have strategies in place to recover funds from them if automatic attempts to debit customer funding accounts fail.

#### Payouts to customers’ vendors

![Image showing two types of transfer customers. Originators move money directly to customers. Platforms involve an intermediary business.](/assets/img/docs/transfer/platform_vendor_payout.png)

Flow of funds overview for vendor payouts

##### Issuing payouts to customers’ vendors

Transfer for Platforms can be used to help your customers manage their bills. Payments can be made to customers’ vendors for goods or services rendered to the customer.

Payouts to individuals are not supported unless related to payment collection (e.g. customers can issue refunds to bill payers, but cannot issue stand-alone payments to consumers).

To issue a vendor payment, funds must be available in your end-customer's Ledger balance. To add funds, create a deposit by calling [`/transfer/ledger/deposit`](/docs/api/products/transfer/ledger/#transferledgerdeposit) and providing the `originator_client_id`, the `amount`, and payment method (e.g. `ach`, `same-day-ach`, `rtp`). This will move funds from the customer's funding account to their Ledger balance at the next available window. These funds will land first into the Ledger's pending balance and become available for use at 3pm ET three business days after the `sweep.settled` event is emitted. For example, a same-day ACH deposit initiated at 9am ET on a Monday is available at 3pm ET on Thursday.

You can also set a minimum available balance threshold in the Plaid dashboard for each customer. Setting a minimum balance greater than $0 helps ensure there are consistently enough funds to cover payouts.

After funding your customer's Ledger balance, you can create payouts by calling [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate) with a `type` of `credit`, followed by [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate), which will decrease their available Ledger balance and issue the payout.

Any ACH returns or processing failures will return the funds back into the available balance.

#### Collecting fees as a Platform

If you collect a per-transfer fee for incoming payments, you can collect that fee using the facilitator fee feature. When calling [`/transfer/create`](/docs/api/products/transfer/initiating-transfers/#transfercreate) with `type=debit`, specify the `facilitator_fee` parameter. The amount specified in the `facilitator_fee` will be allocated towards your Plaid Ledger balance instead of the customer's.

For example, if you specify `facilitator_fee: "2.00"` on a $10 ACH debit, $8 will be placed in the end-customer's Ledger pending balance, and $2 will be placed in your Platform pending balance. Both amounts will be placed in the Ledgers' pending balances and will be converted to available balances after the hold time expires. If the debit results in an ACH return, funds are clawed back from both Ledgers.

Facilitator fees are available only on ACH debit transfers. Facilitator fees are not supported on end-customer deposits or withdrawals to or from a Ledger balance. To support a custom fee structure (such as taking a fixed monthly fee from your end-customer), you can use [`/transfer/ledger/distribute`](/docs/api/products/transfer/ledger/#transferledgerdistribute).

#### Moving funds between Ledgers

You can move money between your own Platform-level Ledger balance and the Ledger balance of any of your customers by using [`/transfer/ledger/distribute`](/docs/api/products/transfer/ledger/#transferledgerdistribute). Currently, [`/transfer/event/sync`](/docs/api/products/transfer/reading-transfers/#transfereventsync) does not return any events for Ledger distributions. You can track ledger distribution events via [`/transfer/ledger/event/list`](/docs/api/products/transfer/ledger/#transferledgereventlist).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
