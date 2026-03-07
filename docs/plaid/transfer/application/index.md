---
title: "Transfer - Transfer Application | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/application/"
scraped_at: "2026-03-07T22:05:23+00:00"
---

# Transfer Application

#### Get approved to create transfers

[Request Access](https://dashboard.plaid.com/overview/production)[Create Transfers](/docs/transfer/creating-transfers/)

When you [request access to Transfer in Production](https://dashboard.plaid.com/overview/request-products), Plaid will collect information from you regarding your business and customer verification processes. If your company is privately held, you will also need to submit identity verification details for a "control person" (such as an executive or senior manager).

#### Originators vs Platforms

As part of your application, Plaid will confirm what type of Transfer customer you are - either an Originator or a Platform.

![Image showing two types of transfer customers. Originators move money directly to customers. Platforms involve an intermediary business.](/assets/img/docs/transfer/originator_platform.png)

Types of Transfer customers

Transfer for Platforms is still in beta. If you are interested in Transfer for Platforms, [contact sales](https://plaid.com/contact/) or your Plaid account manager, or see the [documentation](/docs/transfer/platform-payments/) to learn more.

##### What is an originator?

An Originator initiates ACH payments and is the main beneficiary of these funds. Originators are a party to every transfer and have a direct contractual relationship with the consumer. If you are not contractually obligated to be the main beneficiary of the funds, you are likely a [Platform](/docs/transfer/application/#what-is-a-platform).

###### Examples of Originators:

- Neobank funding a new user account
- E-commerce website offering ACH as a payment method
- Lender disbursing loan funds to a user

##### What is a Platform?

A Platform onboards originators of their own and facilitates money movement between them and their consumers. Platforms do not have a direct contractual relationship with their originator's consumers and are not obligated to be paid; they are simply an intermediary.

###### Examples of Platforms:

- Marketplace website that brings buyers and sellers together
- Any reseller partner, such as a payment processor that provides an SDK to their customers using Plaid transfer functionality
- SaaS provider that sells account funding tools or payment services
- Crowdfunding website for nonprofit organizations

To learn more about Plaid's Platforms solution, see the [Transfer for Platforms documentation](/docs/transfer/platform-payments/).

#### Prohibited industries

There are certain types of businesses, products and industries that can't use Plaid's Transfer services. You must not use Plaid's Transfer services for any illegal activities or for the businesses, product types, or industries listed below. The types of industries listed here are representative of prohibited categories, but this is not an exhaustive list.

- Adult content or services
- Marijuana dispensaries and related products
- Payday lending
- Weapons or munitions

To use Plaid's Transfer services, you must remain compliant with these industry limitations. There are additional restricted industries that are subject to enhanced due diligence for approval for Transfer usage.

If you have questions about prohibited or restricted industries for Transfer, please [contact us](https://plaid.com/contact/).

#### Application process

Requesting access to Transfer involves completing the following forms:

- The [Production Request form](https://dashboard.plaid.com/onboarding/) (if you already have Plaid Production access for another product, request Transfer via the [Products page](https://dashboard.plaid.com/settings/team/products) instead)
- The [Transfer Application](https://dashboard.plaid.com/transfer/questionnaire), also referred to as the Transfer Questionnaire.
- [Transfer Document collection](https://dashboard.plaid.com/transfer/documents).

The Transfer Application and Transfer Document Collection cannot be started until the Production Request form is complete. However, you can begin these forms while your Production Request is in review.

If any of these forms are started but not complete, a link to complete them will appear on the main homepage of your Dashboard. You can also access them directly by clicking the links above.

If you don't have all the required information available, you can save any of these forms and complete them later. You can also send links to these forms with other members of your Plaid team to complete. The status of the forms will be shared across the team, so that multiple different team members can each fill out different parts.

##### General information

- Description of business
- Industry
- General use case(s) (dropdown selection)
- Business entity and individual information
- Requested Plaid products

##### Transfer-specific information

- Specific description of Transfer use case
- ACH processing history
- ACH processing projections

If Plaid is unable to validate individual or entity information through primary methods, we may request supplemental documents. Requests may include documentation to verify:

- Name (Business entity or individual)
- Address (Business entity or individual)
- SSN / EIN
- DOB

Certain restricted industries will require Enhanced Due Diligence (EDD) specific to that industry. Materials requested during EDD may include:

- Proof of relevant licensing or registration for regulated activities
- Copies of policies and procedures related to compliance with regulatory requirements
- Regulatory audit records
- Consumer complaint policies
- Financial statements
- Terms and conditions
- Responses to industry-specific questions

Depending on the industry, these materials may be requested either as part of the Transfer application, or later, during the application review process.

#### Implementation checklist

After your Transfer access has been approved, additional operational tasks required for implementation can be found under [“Transfer > Tasks”](https://dashboard.plaid.com/transfer/implementation-checklist) in the left-hand menu of the Dashboard.

For additional tasks required as part of implementation, see the [Launch Center](https://dashboard.plaid.com/developers/launch-center).

#### Transfer billing

Transfer is typically billed as a fee for each transfer, plus a flat fee for each Transfer Item and for each call to [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate). For details of the transfer pricing model, see the [Transfer billing documentation](/docs/account/billing/#transfer-fee-model). Exact transfer pricing will vary depending on the details of your business, payments activity, and Plaid integration. After you have completed the Transfer Application form, Plaid will contact you with more information about the pricing that applies to you.

Transfer is available only via a Custom plan with a 12-month minimum contract and is not available via Pay-as-you-go.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
