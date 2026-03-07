---
title: "Errors - Payment errors (Europe) | Plaid Docs"
source_url: "https://plaid.com/docs/errors/payment/"
scraped_at: "2026-03-07T22:04:52+00:00"
---

# Payment Errors

#### Errors specific to the Payment Initiation product

#### **PAYMENT\_BLOCKED**

##### The payment was blocked for violating compliance rules.

##### Common causes

- The payment amount value when calling [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) was too high.
- Too many payments were created in a short period of time.

##### Troubleshooting steps

Contact your Plaid Account Manager for your compliance rules.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_BLOCKED",
 "error_message": "payment blocked",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_CANCELLED**

##### The payment was cancelled.

##### Sample user-facing error message

Payment cancelled: Try making your payment again or select a different bank to continue.

##### Common causes

- The payment was cancelled by the user during the authorisation process

##### Troubleshooting steps

Try again or choose another institution to initiate a payment.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_CANCELLED",
 "error_message": "user cancelled the payment",
 "display_message": "Try making your payment again or select a different bank to continue.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_INSUFFICIENT\_FUNDS**

##### The account has insufficient funds to complete the payment.

##### Sample user-facing error message

Insufficient funds: There isn't enough money in this account to complete the payment. Try again, or select another account or bank.

##### Common causes

- The account selected has insufficient funds to complete the payment.

##### Troubleshooting steps

Try again or choose another account or institution to initiate a payment.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_INSUFFICIENT_FUNDS",
 "error_message": "insufficient funds to complete the request",
 "display_message": "There isn't enough money in this account to complete the payment. Try again, or select another account or bank.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_INVALID\_RECIPIENT**

##### The recipient was rejected by the chosen institution.

##### Sample user-facing error message

Payment failed: Try making your payment again or select a different bank to continue.

##### Common causes

- The recipient name is too long or contains special characters.
- The address is too long or contains special characters.
- The account number is invalid.

##### Troubleshooting steps

Try again with a different recipient.

Create a new recipient with a shorter name and address and/or without special characters.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_INVALID_RECIPIENT",
 "error_message": "payment recipient invalid",
 "display_message": "The payment recipient is invalid for the selected institution. Create a new payment with a different recipient.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_INVALID\_REFERENCE**

##### The reference was rejected by the chosen institution.

##### Sample user-facing error message

Payment failed: Try making your payment again or select a different bank to continue.

##### Common causes

- The reference is too long or contains special characters.

##### Troubleshooting steps

Try again with a different reference.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_INVALID_REFERENCE",
 "error_message": "payment reference invalid",
 "display_message": "The payment reference is invalid for the selected institution. Create a new payment with a different reference.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_INVALID\_SCHEDULE**

##### The schedule was rejected by the chosen institution.

##### Sample user-facing error message

Payment failed: Try making your payment again or select a different bank to continue.

##### Common causes

- The chosen institution does not support negative payment execution days.
- The first payment date is a holiday or on a weekend.

##### Troubleshooting steps

Try again with a different schedule.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_INVALID_SCHEDULE",
 "error_message": "payment schedule invalid",
 "display_message": "The payment schedule is invalid for the selected institution. Create a new payment with a different schedule.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_REJECTED**

##### The payment was rejected by the chosen institution.

##### Sample user-facing error message

Payment failed: Try making your payment again or select a different bank to continue.

##### Common causes

- The amount was too large.
- The payment was considered suspicious by the institution.

##### Troubleshooting steps

Try again with different payment parameters.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_REJECTED",
 "error_message": "payment rejected",
 "display_message": "The payment was rejected by the institution. Try again, or select another account or bank.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_SCHEME\_NOT\_SUPPORTED**

##### The requested payment scheme is not supported by the chosen institution.

##### Sample user-facing error message

Payment failed: Try making your payment again or select a different bank to continue.

##### Common causes

- The payment scheme specified when calling [`/payment_initiation/payment/create`](/docs/api/products/payment-initiation/#payment_initiationpaymentcreate) is not supported by the institution.
- Scheme automatic downgrade failed.

##### Troubleshooting steps

Try again with a different payment scheme.

Try again with a different institution.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_SCHEME_NOT_SUPPORTED",
 "error_message": "payment scheme not supported",
 "display_message": "The payment scheme is not supported by the institution. Either change scheme or select another institution.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_CONSENT\_INVALID\_CONSTRAINTS**

##### The payment consent constraints are missing or not supported by the institution.

##### Sample user-facing error message

Payment failed: Try making your payment again or select a different bank to continue.

##### Common causes

- The payment consent constraints specified when calling [`/payment_initiation/consent/create`](/docs/api/products/payment-initiation/#payment_initiationconsentcreate) are not supported by the institution.

##### Troubleshooting steps

Recreate the consent with different constraints.

Try again with a different institution.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_CONSENT_INVALID_CONSTRAINTS",
 "error_message": "payment consent constraints invalid",
 "display_message": "The payment consent constraints are missing or not supported by the institution. Either update constraints or select another institution.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **PAYMENT\_CONSENT\_CANCELLED**

##### The payment consent was cancelled.

##### Sample user-facing error message

Payment cancelled: Try setting up your payment again or select a different bank to continue

##### Common causes

- The payment consent was cancelled by the user during the authorisation process.

##### Troubleshooting steps

Try again or choose another institution to authorise the payment consent.

API error response

```
http code 400
{
 "error_type": "PAYMENT_ERROR",
 "error_code": "PAYMENT_CONSENT_CANCELLED",
 "error_message": "user cancelled the payment consent",
 "display_message": "Authorise your payment consent again or select a different bank to continue.",
 "request_id": "HNTDNrA8F1shFEW"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
