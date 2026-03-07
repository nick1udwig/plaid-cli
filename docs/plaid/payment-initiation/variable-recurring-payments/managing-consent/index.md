---
title: "Payments (Europe) - Managing Consent | Plaid Docs"
source_url: "https://plaid.com/docs/payment-initiation/variable-recurring-payments/managing-consent/"
scraped_at: "2026-03-07T22:05:12+00:00"
---

# Managing and Revoking Consent

#### Learn how to track payment consent status and allow users to revoke consent

#### Tracking consent status

Whenever a consent status changes, Plaid will send a [`CONSENT_STATUS_UPDATE`](/docs/api/products/payment-initiation/#consent_status_update) webhook to the webhook listener endpoint that you specified during the [`/link/token/create`](/docs/api/link/#linktokencreate) call.

You can also retrieve consent status using [`/payment_initiation/consent/get`](/docs/api/products/payment-initiation/#payment_initiationconsentget).

#### Consent authorisation

All consents begin in the `UNAUTHORISED` status.

Once a user has [completed the consent flow in Link](/docs/payment-initiation/variable-recurring-payments/add-to-app/#launch-the-payment-flow-in-link), the consent status will update to `AUTHORISED`.

At this point you can make payments within the consent parameters, with no user input required.

If the user exits Link without authorising the payment, or rejects the payment within Link, the consent will still remain in the `UNAUTHORISED` status.

#### Consent revocation

The end user may decide to revoke their consent, which can be done via their bank, your app or the [Plaid Portal](http://my.plaid.com). After consent has been revoked, the consent status will be updated to `REVOKED`, and the `consent_id` can no longer be used to make payments. There is no way to restore a revoked `consent_id`; you will need to create a new `consent_id` and send the user back through Link to grant consent.

To allow end users to revoke consent via your app, implement the [`/payment_initiation/consent/revoke`](/docs/api/products/payment-initiation/#payment_initiationconsentrevoke) endpoint and create an entry point for it in your UI.

/payment\_initiation/consent/revoke request

```
const request: PaymentInitiationConsentRevokeRequest = {
  consent_id: consentID,
};
try {
  const response = await plaidClient.paymentInitiationConsentRevoke(request);
} catch (error) {
  // handle error
}
```

#### Consent rejection

A consent will enter the `REJECTED` status only if it is rejected by the bank. Common reasons may include the bank not supporting the scope of the VRP; for example, not all banks support `COMMERCIAL` scopes.

#### Consent expiration

If the consent was created with an expiration date-time when calling [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate), the consent will move into the `EXPIRED` status once that date-time has passed. An expired consent cannot be refreshed; you must create a new one instead.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
