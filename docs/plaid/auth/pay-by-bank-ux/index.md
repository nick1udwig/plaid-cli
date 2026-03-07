---
title: "Auth - Increasing pay-by-bank adoption | Plaid Docs"
source_url: "https://plaid.com/docs/auth/pay-by-bank-ux/"
scraped_at: "2026-03-07T22:04:45+00:00"
---

# Increasing pay-by-bank adoption

#### Enhance discovery and conversion on bank payments via UX design

#### Increasing Pay by bank adoption

Properly presenting 'pay by bank' as a payment method and showcasing the convenience of using Plaid for account authentication can result in up to 2-5x more users choosing to pay by bank.

#### Payment method presentation

Place 'Pay by bank' as the first option in a payment list, have it pre-selected with a radio button, and display [Embedded Institution Search](/docs/link/embedded-institution-search/) inline. This shows that instant verification is available and easy to use.

![Payment page with options for monthly, quarterly, or annual plans. 'Pay by bank' is selected. Bank logos displayed for selection. Total: $47.49/month.](/assets/img/docs/embedded-search-desktop.png)

Example of Embedded Institution Search on Desktop

##### Embedded Institution Select default open

Users are more likely to select 'Pay by bank' if they understand that it will be a secure, open banking-powered experience, rather than having to enter an account and routing number, which is their default expectation. Users are even more likely to use pay by bank if they see their bank's logo within Link.

Plaid will dynamically show the user the most popular institutions for your application in their area, or you can customize this list via the Plaid Dashboard, under [Link->Link Customization->Institution Select](https://dashboard.plaid.com/link/institution-select).

###### Recommendations:

- Show as many logos as possible, following the [breakpoint guidance](/docs/link/embedded-institution-search/#customization-options), in order to increase the likelihood that the user will see their bank.
- By providing the user’s phone number when calling [`/link/token/create`](/docs/api/link/#linktokencreate), you can activate the streamlined version of [returning user experience](/docs/link/returning-user/#pre-filling-phone-numbers-for-faster-account-verification), which has been shown to drive a 2x increase in bank payments adoption for returning users.
- Ensure Embedded Institution Search is visible by default, without requiring the user to first select 'Pay by bank'.

#### Using accordion open

If the Embedded Institution Search inline display option is not available to you due to design or technical constraints, use an "accordion open" technique. Once the user selects 'Pay by bank', render the Embedded Institution Search for easy user comprehension of their next steps to add a bank account.

![Payment page with options for monthly, quarterly, annual plans. Select pay by bank or card. Pet insurance $47.49/month. Pay button.](/assets/img/docs/embedded-link-accordion.gif)

Example of Embedded Institution Search on Desktop

###### Recommendations:

- Labeling the payment method as 'Pay by bank' and adding an appropriate CTA such as 'instantly verify your bank account' will help users understand that your pay by bank flow is powered by open banking.
- Use "accordion open" if necessary to display Embedded Institution Search.

#### Configure Account Select

The [Account Select](/docs/link/customization/#account-select) option "Enabled for one account" setting configures the Link UI so that the end user may only select a single account. This is the appropriate configuration for most pay-by-bank use cases. You can configure this option in the Dashboard, under [Link > Link Customization > Account Select](https://dashboard.plaid.com/link/account-select).

#### Enable Database Auth

Although most users prefer signing into their bank account via open banking, some prefer to provide their account and routing numbers manually.

##### Recommendations:

- To support both populations, enable [Database Auth](/docs/auth/coverage/database-auth/).

Database Auth is appropriate for low-to-medium-risk scenarios, such as recurring bill payments and loan repayments. If you experience high rates of fraud or ACH returns in your payments flow, you should not enable Database Auth.

#### Auth Type Select

For user populations with a higher propensity to link accounts manually, Plaid provides [Auth Type Select](/docs/auth/coverage/flow-options/#adding-manual-verification-entry-points-with-auth-type-select) as an upfront option to the end user to choose between open banking login or manual account connection. Otherwise, this option will only be displayed if the user can't find their bank or encounters an error trying to connect.

Examples of populations that may prefer manually linking include business users (who may not have access to their organization's online banking credentials) or users who are less comfortable with technology.

![Pay by bank: Select payment method with bank search and Auth Type Select. Monthly plan, $47.49/mo summary visible.](/assets/img/docs/embedded-link-ats.gif)

Embedded Institution Search with Auth Type Select enabled

Auth Type Select will increase the percentage of users whose proposed payments cannot be evaluated for risk. Use [Signal Transaction Scores](/docs/signal/signal-rules/#data-availability-limitations) and [Identity Match](/docs/identity/#using-identity-match-with-micro-deposit-or-database-items) with manually verified Items to reduce the risk of return. If you experience high rates of fraud or ACH returns in your payments flow, you should not enable Auth Type Select.

##### Recommendations:

- Enable [Auth Type Select](/docs/auth/coverage/flow-options/#adding-manual-verification-entry-points-with-auth-type-select) if you have reason to believe that your user population prefers to link accounts manually and your risk profile can tolerate lower coverage of anti-fraud and anti-ACH-return checks.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
