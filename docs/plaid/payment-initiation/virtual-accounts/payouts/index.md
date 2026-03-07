---
title: "Payments (Europe) - Payouts | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/virtual-accounts/payouts/"
scraped_at: "2026-03-07T22:05:14+00:00"
---

# Payouts

#### Send funds from your virtual account

A payout represents the flow of funds from a client’s virtual account to an end user. A payout can be for an arbitrary amount if you have sufficient funds stored in your virtual account to cover the payout. Using Payouts requires integrating with virtual account API routes and receiving transaction webhooks.

#### Key Benefits of Payouts

- **Instant**: Payouts ensures that funds arrive in your users' accounts almost immediately.
- **Easy to Integrate**: A single, unified API for both Payment Initiation and Payouts streamlines the development process.
- **Verified**: In conjunction with Plaid's Auth and Identity products, users' bank details are securely and automatically populated. This pre-verification step ensures you are always sending funds to a valid and correct bank account, minimising payment failures and fraud risk.
- **Low Cost**: Move away from the expensive and variable fees associated with card payments. Payouts offers a low, fixed-fee structure, allowing you to significantly reduce costs.

#### Execute a Payout

Make sure your virtual account is set up before following these steps. For more information on setting up an account, see [Managing virtual accounts](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/).

1. Call [`/wallet/get`](/docs/api/products/virtual-accounts/#walletget) to check your virtual account balance (optional).

   - If you have insufficient funds to make your desired payout, make sure to [fund your virtual account](/docs/payment-initiation/virtual-accounts/managing-virtual-accounts/#fund-your-virtual-account) before proceeding. After funding your virtual account, another request to [`/wallet/get`](/docs/api/products/virtual-accounts/#walletget) will show the updated balance.
2. Call [`/wallet/transaction/execute`](/docs/api/products/virtual-accounts/#wallettransactionexecute) and store the `transaction_id` and status from the response.

   - [Configure transaction webhooks](https://dashboard.plaid.com/developers/webhooks) to receive real-time [status update webhooks](/docs/api/products/virtual-accounts/#wallet_transaction_status_update) for each payout transaction.
   - In addition to using webhooks, you can confirm the transaction has been executed by calling [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) with the `transaction_id`.

#### Verification of Payee

Verification of Payee is a regulatory requirement under the EU Instant Payments Regulation. It requires that PSPs validate the recipient's name and IBAN against their bank account details before execution, reducing fraud and user error. Plaid automatically performs Verification of Payee checks for all Eurozone payouts.

UK Payouts and Refunds are not typically subject to Verification of Payee.

##### Verification statuses

The Verification of Payee check returns one of four statuses: Match, Partial Match, No Match, and Match not possible (indicating that the verification system is unavailable). By default, payments proceed for all statuses except No Match. You can also customize the product behavior for the Partial Match and No Match statuses by contacting your Account Manager.

If the payment is blocked:

- The `WALLET_TRANSACTION_STATUS_UPDATE` webhook will indicate a failed payment.
- Calling [`/wallet/transaction/get`](/docs/api/products/virtual-accounts/#wallettransactionget) with the `transaction_id` will show a the `TRANSACTION_PAYEE_VERIFICATION_NO_MATCH` error.

#### Testing Payouts

You can begin testing Payouts in Sandbox by following the steps listed in the [Add Virtual Accounts to your App](/docs/payment-initiation/virtual-accounts/add-to-app/) guide. For Production access you will first need to [submit a product access request Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
