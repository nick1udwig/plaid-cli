---
title: "API - Consent | Plaid Docs"
source_url: "https://plaid.com/docs/api/consent/"
scraped_at: "2026-03-07T22:03:46+00:00"
---

# Consent

#### API reference for managing consent

| Endpoints |  |
| --- | --- |
| [`/consent/events/get`](/docs/api/consent/#consenteventsget) | Retrieve consent events |

| See also |  |
| --- | --- |
| [`/item/get`](/docs/api/items/#itemget) | Retrieve an Item (includes Item consent details) |

=\*=\*=\*=

#### `/consent/events/get`

#### List a historical log of item consent events

List a historical log of Item consent events. Consent logs are only available for events occurring on or after November 7, 2024. Extremely recent events (occurring within the past 12 hours) may not be available via this endpoint. Up to three years of consent logs will be available via the endpoint.

/consent/events/get

**Request fields**

[`client_id`](/docs/api/consent/#consent-events-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/consent/#consent-events-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`access_token`](/docs/api/consent/#consent-events-get-request-access-token)

requiredstringrequired, string

The access token associated with the Item data is being requested for.

/consent/events/get

```
const request: ConsentEventsGetRequest = {
  access_token: accessToken,
};
try {
  const response = await plaidClient.consentEventsGet(request);
  const consentEvents = response.data.consent_events;
  const status = response.data.status;
} catch (error) {
  // handle error
}
```

/consent/events/get

**Response fields**

[`request_id`](/docs/api/consent/#consent-events-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`consent_events`](/docs/api/consent/#consent-events-get-response-consent-events)

[object][object]

A list of consent events.

[`item_id`](/docs/api/consent/#consent-events-get-response-consent-events-item-id)

stringstring

The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.

[`created_at`](/docs/api/consent/#consent-events-get-response-consent-events-created-at)

stringstring

The date and time when the consent event occurred, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.  
  

Format: `date-time`

[`event_type`](/docs/api/consent/#consent-events-get-response-consent-events-event-type)

stringstring

A broad categorization of the consent event.  
  

Possible values: `CONSENT_GRANTED`, `CONSENT_REVOKED`, `CONSENT_UPDATED`

[`event_code`](/docs/api/consent/#consent-events-get-response-consent-events-event-code)

stringstring

Codes describing the object of a consent event.  
  

Possible values: `USER_AGREEMENT`, `USE_CASES`, `DATA_SCOPES`, `ACCOUNT_SCOPES`, `REVOCATION`

[`institution_id`](/docs/api/consent/#consent-events-get-response-consent-events-institution-id)

nullablestringnullable, string

Unique identifier for the institution associated with the Item. Field is `null` for Items created via Same Day Micro-deposits.

[`institution_name`](/docs/api/consent/#consent-events-get-response-consent-events-institution-name)

nullablestringnullable, string

The full name of the institution associated with the Item. Field is `null` for Items created via Same Day Micro-deposits.

[`initiator`](/docs/api/consent/#consent-events-get-response-consent-events-initiator)

stringstring

The entity that initiated collection of consent.  
  

Possible values: `PLAID`, `DATA_PROVIDER`, `CUSTOMER`, `END_USER`

[`consented_use_cases`](/docs/api/consent/#consent-events-get-response-consent-events-consented-use-cases)

[string][string]

A list of strings containing the full list of use cases the end user has consented to for the Item.  
See the [full list](https://plaid.com/docs/link/data-transparency-messaging-migration-guide/#updating-link-customizations) of use cases.

[`consented_data_scopes`](/docs/api/consent/#consent-events-get-response-consent-events-consented-data-scopes)

[string][string]

A list of strings containing the full list of data scopes the end user has consented to for the Item. These correspond to consented products; see the [full mapping](https://plaid.com/docs/link/data-transparency-messaging-migration-guide/#data-scopes-by-product) of data scopes and products.

[`consented_accounts`](/docs/api/consent/#consent-events-get-response-consent-events-consented-accounts)

[object][object]

An array containing the accounts associated with the Item for which authorizations are granted.

[`account_id`](/docs/api/consent/#consent-events-get-response-consent-events-consented-accounts-account-id)

stringstring

Plaid’s unique identifier for the account. Like all Plaid identifiers, the `account_id` is case sensitive.

[`mask`](/docs/api/consent/#consent-events-get-response-consent-events-consented-accounts-mask)

stringstring

The last 2-4 alphanumeric characters of an account's official account number

[`name`](/docs/api/consent/#consent-events-get-response-consent-events-consented-accounts-name)

stringstring

The name of the account, either assigned by the user or by the financial institution itself

[`official_name`](/docs/api/consent/#consent-events-get-response-consent-events-consented-accounts-official-name)

stringstring

The official name of the account as given by the financial institution

[`type`](/docs/api/consent/#consent-events-get-response-consent-events-consented-accounts-type)

stringstring

`investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  
`credit:` Credit card  
`depository:` Depository account  
`loan:` Loan account  
`other:` Non-specified account type  
See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `investment`, `credit`, `depository`, `loan`, `brokerage`, `other`

[`subtype`](/docs/api/consent/#consent-events-get-response-consent-events-consented-accounts-subtype)

nullablestringnullable, string

See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.  
  

Possible values: `401a`, `401k`, `403B`, `457b`, `529`, `auto`, `brokerage`, `business`, `cash isa`, `cash management`, `cd`, `checking`, `commercial`, `construction`, `consumer`, `credit card`, `crypto exchange`, `ebt`, `education savings account`, `fixed annuity`, `gic`, `health reimbursement arrangement`, `home equity`, `hsa`, `isa`, `ira`, `keogh`, `lif`, `life insurance`, `line of credit`, `lira`, `loan`, `lrif`, `lrsp`, `money market`, `mortgage`, `mutual fund`, `non-custodial wallet`, `non-taxable brokerage account`, `other`, `other insurance`, `other annuity`, `overdraft`, `paypal`, `payroll`, `pension`, `prepaid`, `prif`, `profit sharing plan`, `rdsp`, `resp`, `retirement`, `rlif`, `roth`, `roth 401k`, `rrif`, `rrsp`, `sarsep`, `savings`, `sep ira`, `simple ira`, `sipp`, `stock plan`, `student`, `thrift savings plan`, `tfsa`, `trust`, `ugma`, `utma`, `variable annuity`

Response Object

```
{
  "request_id": "m8MDnv9okwxFNBV",
  "consent_events": [
    {
      "item_id": "Ed6bjNrDLJfGvZWwnkQlfxwoNz54B5C97ejBr",
      "event_type": "CONSENT_GRANTED",
      "event_code": "USER_AGREEMENT",
      "institution_id": "ins_123456",
      "institution_name": "Platypus bank",
      "initiator": "END_USER",
      "created_at": "2019-02-15T15:51:39Z",
      "consented_use_cases": [],
      "consented_data_scopes": [],
      "consented_accounts": []
    },
    {
      "item_id": "Ed6bjNrDLJfGvZWwnkQlfxwoNz54B5C97ejBr",
      "event_type": "CONSENT_GRANTED",
      "event_code": "USE_CASES",
      "institution_id": "ins_123456",
      "institution_name": "Platypus bank",
      "initiator": "END_USER",
      "created_at": "2019-02-15T15:52:39Z",
      "consented_use_cases": [
        "Send and receive money",
        "Track and manage your finances"
      ],
      "consented_data_scopes": [],
      "consented_accounts": []
    },
    {
      "item_id": "Ed6bjNrDLJfGvZWwnkQlfxwoNz54B5C97ejBr",
      "event_type": "CONSENT_GRANTED",
      "event_code": "DATA_SCOPES",
      "institution_id": "ins_123456",
      "institution_name": "Platypus bank",
      "initiator": "END_USER",
      "created_at": "2019-02-15T15:52:39Z",
      "consented_use_cases": [],
      "consented_data_scopes": [
        "account_balance_info",
        "contact_info",
        "account_routing_number"
      ],
      "consented_accounts": []
    },
    {
      "item_id": "Ed6bjNrDLJfGvZWwnkQlfxwoNz54B5C97ejBr",
      "event_type": "CONSENT_GRANTED",
      "event_code": "ACCOUNT_SCOPES",
      "institution_id": "ins_123456",
      "institution_name": "Platypus bank",
      "initiator": "END_USER",
      "created_at": "2019-02-15T15:53:39Z",
      "consented_use_cases": [],
      "consented_data_scopes": [],
      "consented_accounts": [
        {
          "account_id": "blgvvBlXw3cq5GMPwqB6s6q4dLKB9WcVqGDGo",
          "mask": "0000",
          "name": "Plaid Checking",
          "official_name": "Plaid Gold Standard 0% Interest Checking",
          "type": "depository",
          "subtype": "checking"
        }
      ]
    },
    {
      "item_id": "Ed6bjNrDLJfGvZWwnkQlfxwoNz54B5C97ejBr",
      "event_type": "CONSENT_REVOKED",
      "event_code": "REVOCATION",
      "institution_id": "ins_123456",
      "institution_name": "Platypus bank",
      "initiator": "END_USER",
      "created_at": "2020-02-20T15:53:39Z",
      "consented_use_cases": [],
      "consented_data_scopes": [],
      "consented_accounts": []
    }
  ]
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
