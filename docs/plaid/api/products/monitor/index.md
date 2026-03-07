---
title: "API - Monitor | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/monitor/"
scraped_at: "2026-03-07T22:04:18+00:00"
---

# Monitor

#### API reference for Monitor endpoints and webhooks

For how-to guidance, see the [Monitor documentation](/docs/monitor/).

| Endpoints |  |
| --- | --- |
| [`/watchlist_screening/individual/create`](/docs/api/products/monitor/#watchlist_screeningindividualcreate) | Create a watchlist screening for a person |
| [`/watchlist_screening/individual/get`](/docs/api/products/monitor/#watchlist_screeningindividualget) | Retrieve an individual watchlist screening |
| [`/watchlist_screening/individual/list`](/docs/api/products/monitor/#watchlist_screeningindividuallist) | List individual watchlist screenings |
| [`/watchlist_screening/individual/update`](/docs/api/products/monitor/#watchlist_screeningindividualupdate) | Update individual watchlist screening |
| [`/watchlist_screening/individual/history/list`](/docs/api/products/monitor/#watchlist_screeningindividualhistorylist) | List history for entity watchlist screenings |
| [`/watchlist_screening/individual/review/create`](/docs/api/products/monitor/#watchlist_screeningindividualreviewcreate) | Create a review for an individual watchlist screening |
| [`/watchlist_screening/individual/review/list`](/docs/api/products/monitor/#watchlist_screeningindividualreviewlist) | List reviews for individual watchlist screenings |
| [`/watchlist_screening/individual/hit/list`](/docs/api/products/monitor/#watchlist_screeningindividualhitlist) | List hits for individual watchlist screenings |
| [`/watchlist_screening/individual/program/get`](/docs/api/products/monitor/#watchlist_screeningindividualprogramget) | Get individual watchlist screening programs |
| [`/watchlist_screening/individual/program/list`](/docs/api/products/monitor/#watchlist_screeningindividualprogramlist) | List individual watchlist screening programs |
| [`/watchlist_screening/entity/create`](/docs/api/products/monitor/#watchlist_screeningentitycreate) | Create a watchlist screening for an entity |
| [`/watchlist_screening/entity/get`](/docs/api/products/monitor/#watchlist_screeningentityget) | Retrieve an individual watchlist screening |
| [`/watchlist_screening/entity/list`](/docs/api/products/monitor/#watchlist_screeningentitylist) | List individual watchlist screenings |
| [`/watchlist_screening/entity/update`](/docs/api/products/monitor/#watchlist_screeningentityupdate) | Update individual watchlist screening |
| [`/watchlist_screening/entity/history/list`](/docs/api/products/monitor/#watchlist_screeningentityhistorylist) | List history for individual watchlist screenings |
| [`/watchlist_screening/entity/review/create`](/docs/api/products/monitor/#watchlist_screeningentityreviewcreate) | Create a review for an individual watchlist screening |
| [`/watchlist_screening/entity/review/list`](/docs/api/products/monitor/#watchlist_screeningentityreviewlist) | List reviews for individual watchlist screenings |
| [`/watchlist_screening/entity/hit/list`](/docs/api/products/monitor/#watchlist_screeningentityhitlist) | List hits for individual watchlist screenings |
| [`/watchlist_screening/entity/program/get`](/docs/api/products/monitor/#watchlist_screeningentityprogramget) | Get individual watchlist screening programs |
| [`/watchlist_screening/entity/program/list`](/docs/api/products/monitor/#watchlist_screeningentityprogramlist) | List individual watchlist screening programs |

| See also |  |
| --- | --- |
| [`/dashboard_user/get`](/docs/api/kyc-aml-users/#dashboard_userget) | Retrieve information about a dashboard user |
| [`/dashboard_user/list`](/docs/api/kyc-aml-users/#dashboard_userlist) | List dashboard users |

| Webhooks |  |
| --- | --- |
| [`SCREENING: STATUS_UPDATED`](/docs/api/products/monitor/#screening-status_updated) | The status of an individual watchlist screening has changed |
| [`ENTITY_SCREENING: STATUS_UPDATED`](/docs/api/products/monitor/#entity_screening-status_updated) | The status of an entity watchlist screening has changed |

### Endpoints

=\*=\*=\*=

#### `/watchlist_screening/individual/create`

#### Create a watchlist screening for a person

Create a new Watchlist Screening to check your customer against watchlists defined in the associated Watchlist Program. If your associated program has ongoing screening enabled, this is the profile information that will be used to monitor your customer over time.

/watchlist\_screening/individual/create

**Request fields**

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-search-terms)

requiredobjectrequired, object

Search inputs for creating a watchlist screening

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-search-terms-watchlist-program-id)

requiredstringrequired, string

ID of the associated program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-search-terms-legal-name)

requiredstringrequired, string

The legal name of the individual being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`date_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-search-terms-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`document_number`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-search-terms-document-number)

stringstring

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`country`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-search-terms-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/watchlist\_screening/individual/create

```
const request: WatchlistScreeningIndividualCreateRequest = {
  search_terms: {
    watchlist_program_id: 'prg_2eRPsDnL66rZ7H',
    legal_name: 'Aleksey Potemkin',
    date_of_birth: '1990-05-29',
    document_number: 'C31195855',
    country: 'US',
  },
  client_user_id: 'example-client-user-id-123',
};

try {
  const response = await client.watchlistScreeningIndividualCreate(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/create

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-id)

stringstring

ID of the associated screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-search-terms)

objectobject

Search terms for creating an individual watchlist screening

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-search-terms-watchlist-program-id)

stringstring

ID of the associated program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-search-terms-legal-name)

stringstring

The legal name of the individual being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`date_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-search-terms-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`document_number`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`country`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`version`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "scr_52xR9LKo77r1Np",
  "search_terms": {
    "watchlist_program_id": "prg_2eRPsDnL66rZ7H",
    "legal_name": "Aleksey Potemkin",
    "date_of_birth": "1990-05-29",
    "document_number": "C31195855",
    "country": "US",
    "version": 1
  },
  "assignee": "54350110fedcbaf01234ffee",
  "status": "cleared",
  "client_user_id": "your-db-id-3b24110",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/get`

#### Retrieve an individual watchlist screening

Retrieve a previously created individual watchlist screening

/watchlist\_screening/individual/get

**Request fields**

[`watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-individual-get-request-watchlist-screening-id)

requiredstringrequired, string

ID of the associated screening.

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

/watchlist\_screening/individual/get

```
const request: WatchlistScreeningIndividualGetRequest = {
  watchlist_screening_id: 'scr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningIndividualGet(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/get

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-id)

stringstring

ID of the associated screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-search-terms)

objectobject

Search terms for creating an individual watchlist screening

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-search-terms-watchlist-program-id)

stringstring

ID of the associated program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-search-terms-legal-name)

stringstring

The legal name of the individual being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`date_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-search-terms-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`document_number`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`country`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`version`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "scr_52xR9LKo77r1Np",
  "search_terms": {
    "watchlist_program_id": "prg_2eRPsDnL66rZ7H",
    "legal_name": "Aleksey Potemkin",
    "date_of_birth": "1990-05-29",
    "document_number": "C31195855",
    "country": "US",
    "version": 1
  },
  "assignee": "54350110fedcbaf01234ffee",
  "status": "cleared",
  "client_user_id": "your-db-id-3b24110",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/list`

#### List Individual Watchlist Screenings

List previously created watchlist screenings for individuals

/watchlist\_screening/individual/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-list-request-watchlist-program-id)

requiredstringrequired, string

ID of the associated program.

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-list-request-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`status`](/docs/api/products/monitor/#watchlist_screening-individual-list-request-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`assignee`](/docs/api/products/monitor/#watchlist_screening-individual-list-request-assignee)

stringstring

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-individual-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/individual/list

```
const request: WatchlistScreeningIndividualListRequest = {
  watchlist_program_id: 'prg_2eRPsDnL66rZ7H',
  client_user_id: 'example-client-user-id-123',
};

try {
  const response = await client.watchlistScreeningIndividualList(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/list

**Response fields**

[`watchlist_screenings`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings)

[object][object]

List of individual watchlist screenings

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-id)

stringstring

ID of the associated screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-search-terms)

objectobject

Search terms for creating an individual watchlist screening

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-search-terms-watchlist-program-id)

stringstring

ID of the associated program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-search-terms-legal-name)

stringstring

The legal name of the individual being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`date_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-search-terms-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`document_number`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`country`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`version`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-watchlist-screenings-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "watchlist_screenings": [
    {
      "id": "scr_52xR9LKo77r1Np",
      "search_terms": {
        "watchlist_program_id": "prg_2eRPsDnL66rZ7H",
        "legal_name": "Aleksey Potemkin",
        "date_of_birth": "1990-05-29",
        "document_number": "C31195855",
        "country": "US",
        "version": 1
      },
      "assignee": "54350110fedcbaf01234ffee",
      "status": "cleared",
      "client_user_id": "your-db-id-3b24110",
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/update`

#### Update individual watchlist screening

Update a specific individual watchlist screening. This endpoint can be used to add additional customer information, correct outdated information, add a reference id, assign the individual to a reviewer, and update which program it is associated with. Please note that you may not update `search_terms` and `status` at the same time since editing `search_terms` may trigger an automatic `status` change.

/watchlist\_screening/individual/update

**Request fields**

[`watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-watchlist-screening-id)

requiredstringrequired, string

ID of the associated screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-search-terms)

objectobject

Search terms for editing an individual watchlist screening

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-search-terms-watchlist-program-id)

stringstring

ID of the associated program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-search-terms-legal-name)

stringstring

The legal name of the individual being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`date_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-search-terms-date-of-birth)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`document_number`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-search-terms-document-number)

stringstring

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`country`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-search-terms-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-assignee)

stringstring

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`reset_fields`](/docs/api/products/monitor/#watchlist_screening-individual-update-request-reset-fields)

[string][string]

A list of fields to reset back to null  
  

Possible values: `assignee`

/watchlist\_screening/individual/update

```
const request: WatchlistScreeningIndividualUpdateRequest = {
  watchlist_screening_id: 'scr_52xR9LKo77r1Np',
  status: 'cleared',
};

try {
  const response = await client.watchlistScreeningIndividualUpdate(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/update

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-id)

stringstring

ID of the associated screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-search-terms)

objectobject

Search terms for creating an individual watchlist screening

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-search-terms-watchlist-program-id)

stringstring

ID of the associated program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-search-terms-legal-name)

stringstring

The legal name of the individual being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`date_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-search-terms-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`document_number`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`country`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`version`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-update-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "scr_52xR9LKo77r1Np",
  "search_terms": {
    "watchlist_program_id": "prg_2eRPsDnL66rZ7H",
    "legal_name": "Aleksey Potemkin",
    "date_of_birth": "1990-05-29",
    "document_number": "C31195855",
    "country": "US",
    "version": 1
  },
  "assignee": "54350110fedcbaf01234ffee",
  "status": "cleared",
  "client_user_id": "your-db-id-3b24110",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/history/list`

#### List history for individual watchlist screenings

List all changes to the individual watchlist screening in reverse-chronological order. If the watchlist screening has not been edited, no history will be returned.

/watchlist\_screening/individual/history/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-request-watchlist-screening-id)

requiredstringrequired, string

ID of the associated screening.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/individual/history/list

```
const request: WatchlistScreeningIndividualHistoryListRequest = {
  watchlist_screening_id: 'scr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningIndividualHistoryList(
    request,
  );
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/history/list

**Response fields**

[`watchlist_screenings`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings)

[object][object]

List of individual watchlist screenings

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-id)

stringstring

ID of the associated screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-search-terms)

objectobject

Search terms for creating an individual watchlist screening

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-search-terms-watchlist-program-id)

stringstring

ID of the associated program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-search-terms-legal-name)

stringstring

The legal name of the individual being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`date_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-search-terms-date-of-birth)

nullablestringnullable, string

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`document_number`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`country`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`version`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-watchlist-screenings-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-history-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "watchlist_screenings": [
    {
      "id": "scr_52xR9LKo77r1Np",
      "search_terms": {
        "watchlist_program_id": "prg_2eRPsDnL66rZ7H",
        "legal_name": "Aleksey Potemkin",
        "date_of_birth": "1990-05-29",
        "document_number": "C31195855",
        "country": "US",
        "version": 1
      },
      "assignee": "54350110fedcbaf01234ffee",
      "status": "cleared",
      "client_user_id": "your-db-id-3b24110",
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/review/create`

#### Create a review for an individual watchlist screening

Create a review for the individual watchlist screening. Reviews are compliance reports created by users in your organization regarding the relevance of potential hits found by Plaid.

/watchlist\_screening/individual/review/create

**Request fields**

[`confirmed_hits`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-request-confirmed-hits)

required[string]required, [string]

Hits to mark as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.

[`dismissed_hits`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-request-dismissed-hits)

required[string]required, [string]

Hits to mark as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.

[`comment`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-request-comment)

stringstring

A comment submitted by a team member as part of reviewing a watchlist screening.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-request-watchlist-screening-id)

requiredstringrequired, string

ID of the associated screening.

/watchlist\_screening/individual/review/create

```
const request: WatchlistScreeningIndividualReviewCreateRequest = {
  confirmed_hits: ['scrhit_52xR9LKo77r1Np'],
  dismissed_hits: [],
  watchlist_screening_id: 'scr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningIndividualReviewCreate(
    request,
  );
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/review/create

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-id)

stringstring

ID of the associated review.

[`confirmed_hits`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-confirmed-hits)

[string][string]

Hits marked as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.

[`dismissed_hits`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-dismissed-hits)

[string][string]

Hits marked as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.

[`comment`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-comment)

nullablestringnullable, string

A comment submitted by a team member as part of reviewing a watchlist screening.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-review-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "rev_aCLNRxK3UVzn2r",
  "confirmed_hits": [
    "scrhit_52xR9LKo77r1Np"
  ],
  "dismissed_hits": [
    "scrhit_52xR9LKo77r1Np"
  ],
  "comment": "These look like legitimate matches, rejecting the customer.",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/review/list`

#### List reviews for individual watchlist screenings

List all reviews for the individual watchlist screening.

/watchlist\_screening/individual/review/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-request-watchlist-screening-id)

requiredstringrequired, string

ID of the associated screening.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/individual/review/list

```
const request: WatchlistScreeningIndividualReviewListRequest = {
  watchlist_screening_id: 'scr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningIndividualReviewList(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/review/list

**Response fields**

[`watchlist_screening_reviews`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews)

[object][object]

List of screening reviews

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews-id)

stringstring

ID of the associated review.

[`confirmed_hits`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews-confirmed-hits)

[string][string]

Hits marked as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.

[`dismissed_hits`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews-dismissed-hits)

[string][string]

Hits marked as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.

[`comment`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews-comment)

nullablestringnullable, string

A comment submitted by a team member as part of reviewing a watchlist screening.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-watchlist-screening-reviews-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-review-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "watchlist_screening_reviews": [
    {
      "id": "rev_aCLNRxK3UVzn2r",
      "confirmed_hits": [
        "scrhit_52xR9LKo77r1Np"
      ],
      "dismissed_hits": [
        "scrhit_52xR9LKo77r1Np"
      ],
      "comment": "These look like legitimate matches, rejecting the customer.",
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/hit/list`

#### List hits for individual watchlist screening

List all hits found by Plaid for a particular individual watchlist screening.

/watchlist\_screening/individual/hit/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-request-watchlist-screening-id)

requiredstringrequired, string

ID of the associated screening.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/individual/hit/list

```
const request: WatchlistScreeningIndividualHitListRequest = {
  watchlist_screening_id: 'scr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningIndividualHitList(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/hit/list

**Response fields**

[`watchlist_screening_hits`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits)

[object][object]

List of individual watchlist screening hits

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-id)

stringstring

ID of the associated screening hit.

[`review_status`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-review-status)

stringstring

The current state of review. All watchlist screening hits begin in a `pending_review` state but can be changed by creating a review. When a hit is in the `pending_review` state, it will always show the latest version of the watchlist data Plaid has available and be compared against the latest customer information saved in the watchlist screening. Once a hit has been marked as `confirmed` or `dismissed` it will no longer be updated so that the state is as it was when the review was first conducted.  
  

Possible values: `confirmed`, `pending_review`, `dismissed`

[`first_active`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-first-active)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`inactive_since`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-inactive-since)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`historical_since`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-historical-since)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`list_code`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-list-code)

stringstring

Shorthand identifier for a specific screening list for individuals.
`AU_CON`: Australia Department of Foreign Affairs and Trade Consolidated List
`CA_CON`: Government of Canada Consolidated List of Sanctions
`EU_CON`: European External Action Service Consolidated List
`IZ_CIA`: CIA List of Chiefs of State and Cabinet Members
`IZ_IPL`: Interpol Red Notices for Wanted Persons List
`IZ_PEP`: Politically Exposed Persons List
`IZ_UNC`: United Nations Consolidated Sanctions
`IZ_WBK`: World Bank Listing of Ineligible Firms and Individuals
`UK_HMC`: UK HM Treasury Consolidated List
`US_DPL`: Bureau of Industry and Security Denied Persons List
`US_DTC`: US Department of State AECA Debarred
`US_FBI`: US Department of Justice FBI Wanted List
`US_FSE`: US OFAC Foreign Sanctions Evaders
`US_ISN`: US Department of State Nonproliferation Sanctions
`US_PLC`: US OFAC Palestinian Legislative Council
`US_SAM`: US System for Award Management Exclusion List
`US_SDN`: US OFAC Specially Designated Nationals List
`US_SSI`: US OFAC Sectoral Sanctions Identifications
`SG_SOF`: Government of Singapore Terrorists and Terrorist Entities
`TR_TWL`: Government of Turkey Terrorist Wanted List
`TR_DFD`: Government of Turkey Domestic Freezing Decisions
`TR_FOR`: Government of Turkey Foreign Freezing Requests
`TR_WMD`: Government of Turkey Weapons of Mass Destruction
`TR_CMB`: Government of Turkey Capital Markets Board  
  

Possible values: `AU_CON`, `CA_CON`, `EU_CON`, `IZ_CIA`, `IZ_IPL`, `IZ_PEP`, `IZ_UNC`, `IZ_WBK`, `UK_HMC`, `US_DPL`, `US_DTC`, `US_FBI`, `US_FSE`, `US_ISN`, `US_MBS`, `US_PLC`, `US_SAM`, `US_SDN`, `US_SSI`, `SG_SOF`, `TR_TWL`, `TR_DFD`, `TR_FOR`, `TR_WMD`, `TR_CMB`

[`plaid_uid`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-plaid-uid)

stringstring

A universal identifier for a watchlist individual that is stable across searches and updates.

[`source_uid`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-source-uid)

nullablestringnullable, string

The identifier provided by the source sanction or watchlist. When one is not provided by the source, this is `null`.

[`analysis`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-analysis)

objectobject

Analysis information describing why a screening hit matched the provided user information

[`dates_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-analysis-dates-of-birth)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`documents`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-analysis-documents)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`locations`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-analysis-locations)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`names`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-analysis-names)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`search_terms_version`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-analysis-search-terms-version)

integerinteger

The version of the screening's `search_terms` that were compared when the screening hit was added. screening hits are immutable once they have been reviewed. If changes are detected due to updates to the screening's `search_terms`, the associated program, or the list's source data prior to review, the screening hit will be updated to reflect those changes.

[`data`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data)

objectobject

Information associated with the watchlist hit

[`dates_of_birth`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-dates-of-birth)

[object][object]

Dates of birth associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-dates-of-birth-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-dates-of-birth-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-dates-of-birth-data)

objectobject

A date range with a start and end date

[`beginning`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-dates-of-birth-data-beginning)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`ending`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-dates-of-birth-data-ending)

stringstring

A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).  
  

Format: `date`

[`documents`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-documents)

[object][object]

Documents associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-documents-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-documents-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-documents-data)

objectobject

An official document, usually issued by a governing body or institution, with an associated identifier.

[`type`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-documents-data-type)

stringstring

The kind of official document represented by this object.  
`birth_certificate` - A certificate of birth  
`drivers_license` - A license to operate a motor vehicle  
`immigration_number` - Immigration or residence documents  
`military_id` - Identification issued by a military group  
`other` - Any document not covered by other categories  
`passport` - An official passport issue by a government  
`personal_identification` - Any generic personal identification that is not covered by other categories  
`ration_card` - Identification that entitles the holder to rations  
`ssn` - United States Social Security Number  
`student_id` - Identification issued by an educational institution  
`tax_id` - Identification issued for the purpose of collecting taxes  
`travel_document` - Visas, entry permits, refugee documents, etc.  
`voter_id` - Identification issued for the purpose of voting  
  

Possible values: `birth_certificate`, `drivers_license`, `immigration_number`, `military_id`, `other`, `passport`, `personal_identification`, `ration_card`, `ssn`, `student_id`, `tax_id`, `travel_document`, `voter_id`

[`number`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-documents-data-number)

stringstring

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`locations`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-locations)

[object][object]

Locations associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-locations-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-locations-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-locations-data)

objectobject

Location information for the associated individual watchlist hit

[`full`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-locations-data-full)

stringstring

The full location string, potentially including elements like street, city, postal codes and country codes. Note that this is not necessarily a complete or well-formatted address.

[`country`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-locations-data-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`names`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-names)

[object][object]

Names associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-names-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-names-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-names-data)

objectobject

Name information for the associated individual watchlist hit

[`full`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-names-data-full)

stringstring

The full name of the individual, including all parts.

[`is_primary`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-names-data-is-primary)

booleanboolean

Primary names are those most commonly used to refer to this person. Only one name will ever be marked as primary.

[`weak_alias_determination`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-watchlist-screening-hits-data-names-data-weak-alias-determination)

stringstring

Names that are explicitly marked as low quality either by their `source` list, or by `plaid` by a series of additional checks done by Plaid. Plaid does not ever surface a hit as a result of a weak name alone. If a name has no quality issues, this value will be `none`.  
  

Possible values: `none`, `source`, `plaid`

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-hit-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "watchlist_screening_hits": [
    {
      "id": "scrhit_52xR9LKo77r1Np",
      "review_status": "pending_review",
      "first_active": "2020-07-24T03:26:02Z",
      "inactive_since": "2020-07-24T03:26:02Z",
      "historical_since": "2020-07-24T03:26:02Z",
      "list_code": "US_SDN",
      "plaid_uid": "uid_3NggckTimGSJHS",
      "source_uid": "26192ABC",
      "analysis": {
        "dates_of_birth": "match",
        "documents": "match",
        "locations": "match",
        "names": "match",
        "search_terms_version": 1
      },
      "data": {
        "dates_of_birth": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "beginning": "1990-05-29",
              "ending": "1990-05-29"
            }
          }
        ],
        "documents": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "type": "passport",
              "number": "C31195855"
            }
          }
        ],
        "locations": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "full": "Florida, US",
              "country": "US"
            }
          }
        ],
        "names": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "full": "Aleksey Potemkin",
              "is_primary": false,
              "weak_alias_determination": "none"
            }
          }
        ]
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/program/get`

#### Get individual watchlist screening program

Get an individual watchlist screening program

/watchlist\_screening/individual/program/get

**Request fields**

[`watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-request-watchlist-program-id)

requiredstringrequired, string

ID of the associated program.

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

/watchlist\_screening/individual/program/get

```
const request: WatchlistScreeningIndividualProgramGetRequest = {
  watchlist_program_id: 'prg_2eRPsDnL66rZ7H',
};

try {
  const response = await client.watchlistScreeningIndividualProgramGet(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/program/get

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-id)

stringstring

ID of the associated program.

[`created_at`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`is_rescanning_enabled`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-is-rescanning-enabled)

booleanboolean

Indicator specifying whether the program is enabled and will perform daily rescans.

[`lists_enabled`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-lists-enabled)

[string][string]

Watchlists enabled for the associated program  
  

Possible values: `AU_CON`, `CA_CON`, `EU_CON`, `IZ_CIA`, `IZ_IPL`, `IZ_PEP`, `IZ_UNC`, `IZ_WBK`, `UK_HMC`, `US_DPL`, `US_DTC`, `US_FBI`, `US_FSE`, `US_ISN`, `US_MBS`, `US_PLC`, `US_SAM`, `US_SDN`, `US_SSI`, `SG_SOF`, `TR_TWL`, `TR_DFD`, `TR_FOR`, `TR_WMD`, `TR_CMB`

[`name`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-name)

stringstring

A name for the program to define its purpose. For example, "High Risk Individuals", "US Cardholders", or "Applicants".

[`name_sensitivity`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-name-sensitivity)

stringstring

The valid name matching sensitivity configurations for a screening program. Note that while certain matching techniques may be more prevalent on less strict settings, all matching algorithms are enabled for every sensitivity.  
`coarse` - See more potential matches. This sensitivity will see more broad phonetic matches across alphabets that make missing a potential hit very unlikely. This setting is noisier and will require more manual review.  
`balanced` - A good default for most companies. This sensitivity is balanced to show high quality hits with reduced noise.  
`strict` - Aggressive false positive reduction. This sensitivity will require names to be more similar than `coarse` and `balanced` settings, relying less on phonetics, while still accounting for character transpositions, missing tokens, and other common permutations.  
`exact` - Matches must be nearly exact. This sensitivity will only show hits with exact or nearly exact name matches with only basic correction such as extraneous symbols and capitalization. This setting is generally not recommended unless you have a very specific use case.  
  

Possible values: `coarse`, `balanced`, `strict`, `exact`

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`is_archived`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-is-archived)

booleanboolean

Archived programs are read-only and cannot screen new customers nor participate in ongoing monitoring.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-program-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "prg_2eRPsDnL66rZ7H",
  "created_at": "2020-07-24T03:26:02Z",
  "is_rescanning_enabled": true,
  "lists_enabled": [
    "US_SDN"
  ],
  "name": "Sample Program",
  "name_sensitivity": "balanced",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "is_archived": false,
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/individual/program/list`

#### List individual watchlist screening programs

List all individual watchlist screening programs

/watchlist\_screening/individual/program/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/individual/program/list

```
try {
  const response = await client.watchlistScreeningIndividualProgramList({});
} catch (error) {
  // handle error
}
```

/watchlist\_screening/individual/program/list

**Response fields**

[`watchlist_programs`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs)

[object][object]

List of individual watchlist screening programs

[`id`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-id)

stringstring

ID of the associated program.

[`created_at`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`is_rescanning_enabled`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-is-rescanning-enabled)

booleanboolean

Indicator specifying whether the program is enabled and will perform daily rescans.

[`lists_enabled`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-lists-enabled)

[string][string]

Watchlists enabled for the associated program  
  

Possible values: `AU_CON`, `CA_CON`, `EU_CON`, `IZ_CIA`, `IZ_IPL`, `IZ_PEP`, `IZ_UNC`, `IZ_WBK`, `UK_HMC`, `US_DPL`, `US_DTC`, `US_FBI`, `US_FSE`, `US_ISN`, `US_MBS`, `US_PLC`, `US_SAM`, `US_SDN`, `US_SSI`, `SG_SOF`, `TR_TWL`, `TR_DFD`, `TR_FOR`, `TR_WMD`, `TR_CMB`

[`name`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-name)

stringstring

A name for the program to define its purpose. For example, "High Risk Individuals", "US Cardholders", or "Applicants".

[`name_sensitivity`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-name-sensitivity)

stringstring

The valid name matching sensitivity configurations for a screening program. Note that while certain matching techniques may be more prevalent on less strict settings, all matching algorithms are enabled for every sensitivity.  
`coarse` - See more potential matches. This sensitivity will see more broad phonetic matches across alphabets that make missing a potential hit very unlikely. This setting is noisier and will require more manual review.  
`balanced` - A good default for most companies. This sensitivity is balanced to show high quality hits with reduced noise.  
`strict` - Aggressive false positive reduction. This sensitivity will require names to be more similar than `coarse` and `balanced` settings, relying less on phonetics, while still accounting for character transpositions, missing tokens, and other common permutations.  
`exact` - Matches must be nearly exact. This sensitivity will only show hits with exact or nearly exact name matches with only basic correction such as extraneous symbols and capitalization. This setting is generally not recommended unless you have a very specific use case.  
  

Possible values: `coarse`, `balanced`, `strict`, `exact`

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`is_archived`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-watchlist-programs-is-archived)

booleanboolean

Archived programs are read-only and cannot screen new customers nor participate in ongoing monitoring.

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-individual-program-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "watchlist_programs": [
    {
      "id": "prg_2eRPsDnL66rZ7H",
      "created_at": "2020-07-24T03:26:02Z",
      "is_rescanning_enabled": true,
      "lists_enabled": [
        "US_SDN"
      ],
      "name": "Sample Program",
      "name_sensitivity": "balanced",
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      },
      "is_archived": false
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/create`

#### Create a watchlist screening for an entity

Create a new entity watchlist screening to check your customer against watchlists defined in the associated entity watchlist program. If your associated program has ongoing screening enabled, this is the profile information that will be used to monitor your customer over time.

/watchlist\_screening/entity/create

**Request fields**

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-search-terms)

requiredobjectrequired, object

Search inputs for creating an entity watchlist screening

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-search-terms-entity-watchlist-program-id)

requiredstringrequired, string

ID of the associated entity program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-search-terms-legal-name)

requiredstringrequired, string

The name of the organization being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`document_number`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-search-terms-document-number)

stringstring

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`email_address`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-search-terms-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`country`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-search-terms-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`phone_number`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-search-terms-phone-number)

stringstring

A phone number in E.164 format.

[`url`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-search-terms-url)

stringstring

An 'http' or 'https' URL (must begin with either of those).  
  

Format: `uri`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

/watchlist\_screening/entity/create

```
const request: WatchlistScreeningEntityCreateRequest = {
  search_terms: {
    entity_watchlist_program_id: 'entprg_2eRPsDnL66rZ7H',
    legal_name: 'Example Screening Entity',
    document_number: 'C31195855',
    email_address: 'user@example.com',
    country: 'US',
    phone_number: '+14025671234',
    url: 'https://example.com',
  },
  client_user_id: 'example-client-user-id-123',
};

try {
  const response = await client.watchlistScreeningEntityCreate(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/create

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-id)

stringstring

ID of the associated entity screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms)

objectobject

Search terms associated with an entity used for searching against watchlists

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms-entity-watchlist-program-id)

stringstring

ID of the associated entity program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms-legal-name)

stringstring

The name of the organization being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`document_number`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`email_address`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`country`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`phone_number`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`url`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms-url)

nullablestringnullable, string

An 'http' or 'https' URL (must begin with either of those).  
  

Format: `uri`

[`version`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "entscr_52xR9LKo77r1Np",
  "search_terms": {
    "entity_watchlist_program_id": "entprg_2eRPsDnL66rZ7H",
    "legal_name": "Al-Qaida",
    "document_number": "C31195855",
    "email_address": "user@example.com",
    "country": "US",
    "phone_number": "+14025671234",
    "url": "https://example.com",
    "version": 1
  },
  "assignee": "54350110fedcbaf01234ffee",
  "status": "cleared",
  "client_user_id": "your-db-id-3b24110",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/get`

#### Get an entity screening

Retrieve an entity watchlist screening.

/watchlist\_screening/entity/get

**Request fields**

[`entity_watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-entity-get-request-entity-watchlist-screening-id)

requiredstringrequired, string

ID of the associated entity screening.

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

/watchlist\_screening/entity/get

```
const request: WatchlistScreeningEntityGetRequest = {
  entity_watchlist_screening_id: 'entscr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningEntityGet(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/get

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-id)

stringstring

ID of the associated entity screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms)

objectobject

Search terms associated with an entity used for searching against watchlists

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms-entity-watchlist-program-id)

stringstring

ID of the associated entity program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms-legal-name)

stringstring

The name of the organization being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`document_number`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`email_address`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`country`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`phone_number`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`url`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms-url)

nullablestringnullable, string

An 'http' or 'https' URL (must begin with either of those).  
  

Format: `uri`

[`version`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "entscr_52xR9LKo77r1Np",
  "search_terms": {
    "entity_watchlist_program_id": "entprg_2eRPsDnL66rZ7H",
    "legal_name": "Al-Qaida",
    "document_number": "C31195855",
    "email_address": "user@example.com",
    "country": "US",
    "phone_number": "+14025671234",
    "url": "https://example.com",
    "version": 1
  },
  "assignee": "54350110fedcbaf01234ffee",
  "status": "cleared",
  "client_user_id": "your-db-id-3b24110",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/list`

#### List entity watchlist screenings

List all entity screenings.

/watchlist\_screening/entity/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-list-request-entity-watchlist-program-id)

requiredstringrequired, string

ID of the associated entity program.

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-list-request-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`status`](/docs/api/products/monitor/#watchlist_screening-entity-list-request-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`assignee`](/docs/api/products/monitor/#watchlist_screening-entity-list-request-assignee)

stringstring

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-entity-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/entity/list

```
const request: WatchlistScreeningEntityListRequest = {
  entity_watchlist_program_id: 'entprg_2eRPsDnL66rZ7H',
};

try {
  const response = await client.watchlistScreeningEntityList(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/list

**Response fields**

[`entity_watchlist_screenings`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings)

[object][object]

List of entity watchlist screening

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-id)

stringstring

ID of the associated entity screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms)

objectobject

Search terms associated with an entity used for searching against watchlists

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms-entity-watchlist-program-id)

stringstring

ID of the associated entity program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms-legal-name)

stringstring

The name of the organization being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`document_number`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`email_address`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`country`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`phone_number`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`url`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms-url)

nullablestringnullable, string

An 'http' or 'https' URL (must begin with either of those).  
  

Format: `uri`

[`version`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-entity-watchlist-screenings-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "entity_watchlist_screenings": [
    {
      "id": "entscr_52xR9LKo77r1Np",
      "search_terms": {
        "entity_watchlist_program_id": "entprg_2eRPsDnL66rZ7H",
        "legal_name": "Al-Qaida",
        "document_number": "C31195855",
        "email_address": "user@example.com",
        "country": "US",
        "phone_number": "+14025671234",
        "url": "https://example.com",
        "version": 1
      },
      "assignee": "54350110fedcbaf01234ffee",
      "status": "cleared",
      "client_user_id": "your-db-id-3b24110",
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/update`

#### Update an entity screening

Update an entity watchlist screening.

/watchlist\_screening/entity/update

**Request fields**

[`entity_watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-entity-watchlist-screening-id)

requiredstringrequired, string

ID of the associated entity screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-search-terms)

objectobject

Search terms for editing an entity watchlist screening

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-search-terms-entity-watchlist-program-id)

requiredstringrequired, string

ID of the associated entity program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-search-terms-legal-name)

stringstring

The name of the organization being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`document_number`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-search-terms-document-number)

stringstring

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`email_address`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-search-terms-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`country`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-search-terms-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`phone_number`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-search-terms-phone-number)

stringstring

A phone number in E.164 format.

[`url`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-search-terms-url)

stringstring

An 'http' or 'https' URL (must begin with either of those).  
  

Format: `uri`

[`assignee`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-assignee)

stringstring

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-client-user-id)

stringstring

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`reset_fields`](/docs/api/products/monitor/#watchlist_screening-entity-update-request-reset-fields)

[string][string]

A list of fields to reset back to null  
  

Possible values: `assignee`

/watchlist\_screening/entity/update

```
const request: WatchlistScreeningEntityUpdateRequest = {
  entity_watchlist_screening_id: 'entscr_52xR9LKo77r1Np',
  status: 'cleared',
};

try {
  const response = await client.watchlistScreeningEntityUpdate(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/update

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-id)

stringstring

ID of the associated entity screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms)

objectobject

Search terms associated with an entity used for searching against watchlists

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms-entity-watchlist-program-id)

stringstring

ID of the associated entity program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms-legal-name)

stringstring

The name of the organization being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`document_number`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`email_address`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`country`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`phone_number`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`url`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms-url)

nullablestringnullable, string

An 'http' or 'https' URL (must begin with either of those).  
  

Format: `uri`

[`version`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-update-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "entscr_52xR9LKo77r1Np",
  "search_terms": {
    "entity_watchlist_program_id": "entprg_2eRPsDnL66rZ7H",
    "legal_name": "Al-Qaida",
    "document_number": "C31195855",
    "email_address": "user@example.com",
    "country": "US",
    "phone_number": "+14025671234",
    "url": "https://example.com",
    "version": 1
  },
  "assignee": "54350110fedcbaf01234ffee",
  "status": "cleared",
  "client_user_id": "your-db-id-3b24110",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/history/list`

#### List history for entity watchlist screenings

List all changes to the entity watchlist screening in reverse-chronological order. If the watchlist screening has not been edited, no history will be returned.

/watchlist\_screening/entity/history/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`entity_watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-request-entity-watchlist-screening-id)

requiredstringrequired, string

ID of the associated entity screening.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/entity/history/list

```
const request: WatchlistScreeningEntityHistoryListRequest = {
  entity_watchlist_screening_id: 'entscr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningEntityHistoryList(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/history/list

**Response fields**

[`entity_watchlist_screenings`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings)

[object][object]

List of entity watchlist screening

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-id)

stringstring

ID of the associated entity screening.

[`search_terms`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms)

objectobject

Search terms associated with an entity used for searching against watchlists

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms-entity-watchlist-program-id)

stringstring

ID of the associated entity program.

[`legal_name`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms-legal-name)

stringstring

The name of the organization being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.

[`document_number`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms-document-number)

nullablestringnullable, string

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`email_address`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms-email-address)

nullablestringnullable, string

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`country`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms-country)

nullablestringnullable, string

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`phone_number`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms-phone-number)

nullablestringnullable, string

A phone number in E.164 format.

[`url`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms-url)

nullablestringnullable, string

An 'http' or 'https' URL (must begin with either of those).  
  

Format: `uri`

[`version`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-search-terms-version)

integerinteger

The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.

[`assignee`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-assignee)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`status`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-status)

stringstring

A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.  
  

Possible values: `rejected`, `pending_review`, `cleared`

[`client_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-client-user-id)

nullablestringnullable, string

A unique ID that identifies the end user in your system. Either a `user_id` or the `client_user_id` must be provided. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-entity-watchlist-screenings-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-history-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "entity_watchlist_screenings": [
    {
      "id": "entscr_52xR9LKo77r1Np",
      "search_terms": {
        "entity_watchlist_program_id": "entprg_2eRPsDnL66rZ7H",
        "legal_name": "Al-Qaida",
        "document_number": "C31195855",
        "email_address": "user@example.com",
        "country": "US",
        "phone_number": "+14025671234",
        "url": "https://example.com",
        "version": 1
      },
      "assignee": "54350110fedcbaf01234ffee",
      "status": "cleared",
      "client_user_id": "your-db-id-3b24110",
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/review/create`

#### Create a review for an entity watchlist screening

Create a review for an entity watchlist screening. Reviews are compliance reports created by users in your organization regarding the relevance of potential hits found by Plaid.

/watchlist\_screening/entity/review/create

**Request fields**

[`confirmed_hits`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-request-confirmed-hits)

required[string]required, [string]

Hits to mark as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.

[`dismissed_hits`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-request-dismissed-hits)

required[string]required, [string]

Hits to mark as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.

[`comment`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-request-comment)

stringstring

A comment submitted by a team member as part of reviewing a watchlist screening.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`entity_watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-request-entity-watchlist-screening-id)

requiredstringrequired, string

ID of the associated entity screening.

/watchlist\_screening/entity/review/create

```
const request: WatchlistScreeningEntityReviewCreateRequest = {
  entity_watchlist_screening_id: 'entscr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningEntityReviewCreate(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/review/create

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-id)

stringstring

ID of the associated entity review.

[`confirmed_hits`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-confirmed-hits)

[string][string]

Hits marked as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.

[`dismissed_hits`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-dismissed-hits)

[string][string]

Hits marked as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.

[`comment`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-comment)

nullablestringnullable, string

A comment submitted by a team member as part of reviewing a watchlist screening.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-review-create-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "entrev_aCLNRxK3UVzn2r",
  "confirmed_hits": [
    "enthit_52xR9LKo77r1Np"
  ],
  "dismissed_hits": [
    "enthit_52xR9LKo77r1Np"
  ],
  "comment": "These look like legitimate matches, rejecting the customer.",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/review/list`

#### List reviews for entity watchlist screenings

List all reviews for a particular entity watchlist screening. Reviews are compliance reports created by users in your organization regarding the relevance of potential hits found by Plaid.

/watchlist\_screening/entity/review/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`entity_watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-request-entity-watchlist-screening-id)

requiredstringrequired, string

ID of the associated entity screening.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/entity/review/list

```
const request: WatchlistScreeningEntityReviewListRequest = {
  entity_watchlist_screening_id: 'entscr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningEntityReviewList(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/review/list

**Response fields**

[`entity_watchlist_screening_reviews`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews)

[object][object]

List of entity watchlist screening reviews

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews-id)

stringstring

ID of the associated entity review.

[`confirmed_hits`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews-confirmed-hits)

[string][string]

Hits marked as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.

[`dismissed_hits`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews-dismissed-hits)

[string][string]

Hits marked as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.

[`comment`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews-comment)

nullablestringnullable, string

A comment submitted by a team member as part of reviewing a watchlist screening.

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-entity-watchlist-screening-reviews-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-review-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "entity_watchlist_screening_reviews": [
    {
      "id": "entrev_aCLNRxK3UVzn2r",
      "confirmed_hits": [
        "enthit_52xR9LKo77r1Np"
      ],
      "dismissed_hits": [
        "enthit_52xR9LKo77r1Np"
      ],
      "comment": "These look like legitimate matches, rejecting the customer.",
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/hit/list`

#### List hits for entity watchlist screenings

List all hits for the entity watchlist screening.

/watchlist\_screening/entity/hit/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`entity_watchlist_screening_id`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-request-entity-watchlist-screening-id)

requiredstringrequired, string

ID of the associated entity screening.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/entity/hit/list

```
const request: WatchlistScreeningEntityHitListRequest = {
  entity_watchlist_screening_id: 'entscr_52xR9LKo77r1Np',
};

try {
  const response = await client.watchlistScreeningEntityHitList(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/hit/list

**Response fields**

[`entity_watchlist_screening_hits`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits)

[object][object]

List of entity watchlist screening hits

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-id)

stringstring

ID of the associated entity screening hit.

[`review_status`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-review-status)

stringstring

The current state of review. All watchlist screening hits begin in a `pending_review` state but can be changed by creating a review. When a hit is in the `pending_review` state, it will always show the latest version of the watchlist data Plaid has available and be compared against the latest customer information saved in the watchlist screening. Once a hit has been marked as `confirmed` or `dismissed` it will no longer be updated so that the state is as it was when the review was first conducted.  
  

Possible values: `confirmed`, `pending_review`, `dismissed`

[`first_active`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-first-active)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`inactive_since`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-inactive-since)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`historical_since`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-historical-since)

nullablestringnullable, string

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`list_code`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-list-code)

stringstring

Shorthand identifier for a specific screening list for entities.
`AU_CON`: Australia Department of Foreign Affairs and Trade Consolidated List
`CA_CON`: Government of Canada Consolidated List of Sanctions
`EU_CON`: European External Action Service Consolidated List
`IZ_SOE`: State Owned Enterprise List
`IZ_UNC`: United Nations Consolidated Sanctions
`IZ_WBK`: World Bank Listing of Ineligible Firms and Individuals
`US_CAP`: US OFAC Correspondent Account or Payable-Through Account Sanctions
`US_FSE`: US OFAC Foreign Sanctions Evaders
`US_MBS`: US Non-SDN Menu-Based Sanctions
`US_SDN`: US Specially Designated Nationals List
`US_SSI`: US OFAC Sectoral Sanctions Identifications
`US_CMC`: US OFAC Non-SDN Chinese Military-Industrial Complex List
`US_UVL`: Bureau of Industry and Security Unverified List
`US_SAM`: US System for Award Management Exclusion List
`US_TEL`: US Terrorist Exclusion List
`UK_HMC`: UK HM Treasury Consolidated List  
  

Possible values: `CA_CON`, `EU_CON`, `IZ_SOE`, `IZ_UNC`, `IZ_WBK`, `US_CAP`, `US_FSE`, `US_MBS`, `US_SDN`, `US_SSI`, `US_CMC`, `US_UVL`, `US_SAM`, `US_TEL`, `AU_CON`, `UK_HMC`

[`plaid_uid`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-plaid-uid)

stringstring

A universal identifier for a watchlist individual that is stable across searches and updates.

[`source_uid`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-source-uid)

nullablestringnullable, string

The identifier provided by the source sanction or watchlist. When one is not provided by the source, this is `null`.

[`analysis`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-analysis)

objectobject

Analysis information describing why a screening hit matched the provided entity information

[`documents`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-analysis-documents)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`email_addresses`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-analysis-email-addresses)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`locations`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-analysis-locations)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`names`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-analysis-names)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`phone_numbers`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-analysis-phone-numbers)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`urls`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-analysis-urls)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`search_terms_version`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-analysis-search-terms-version)

integerinteger

The version of the entity screening's `search_terms` that were compared when the entity screening hit was added. entity screening hits are immutable once they have been reviewed. If changes are detected due to updates to the entity screening's `search_terms`, the associated entity program, or the list's source data prior to review, the entity screening hit will be updated to reflect those changes.

[`data`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data)

objectobject

Information associated with the entity watchlist hit

[`documents`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-documents)

[object][object]

Documents associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-documents-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-documents-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-documents-data)

objectobject

An official document, usually issued by a governing body or institution, with an associated identifier.

[`type`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-documents-data-type)

stringstring

The kind of official document represented by this object.  
`bik` - Russian bank code  
`business_number` - A number that uniquely identifies the business within a category of businesses  
`imo` - Number assigned to the entity by the International Maritime Organization  
`other` - Any document not covered by other categories  
`swift` - Number identifying a bank and branch.  
`tax_id` - Identification issued for the purpose of collecting taxes  
  

Possible values: `bik`, `business_number`, `imo`, `other`, `swift`, `tax_id`

[`number`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-documents-data-number)

stringstring

The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.

[`email_addresses`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-email-addresses)

[object][object]

Email addresses associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-email-addresses-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-email-addresses-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-email-addresses-data)

objectobject

Email address information for the associated entity watchlist hit

[`email_address`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-email-addresses-data-email-address)

stringstring

A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).  
  

Format: `email`

[`locations`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-locations)

[object][object]

Locations associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-locations-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-locations-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-locations-data)

objectobject

Location information for the associated individual watchlist hit

[`full`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-locations-data-full)

stringstring

The full location string, potentially including elements like street, city, postal codes and country codes. Note that this is not necessarily a complete or well-formatted address.

[`country`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-locations-data-country)

stringstring

Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.

[`names`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-names)

[object][object]

Names associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-names-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-names-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-names-data)

objectobject

Name information for the associated entity watchlist hit

[`full`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-names-data-full)

stringstring

The full name of the entity.

[`is_primary`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-names-data-is-primary)

booleanboolean

Primary names are those most commonly used to refer to this entity. Only one name will ever be marked as primary.

[`weak_alias_determination`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-names-data-weak-alias-determination)

stringstring

Names that are explicitly marked as low quality either by their `source` list, or by `plaid` by a series of additional checks done by Plaid. Plaid does not ever surface a hit as a result of a weak name alone. If a name has no quality issues, this value will be `none`.  
  

Possible values: `none`, `source`, `plaid`

[`phone_numbers`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-phone-numbers)

[object][object]

Phone numbers associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-phone-numbers-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-phone-numbers-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-phone-numbers-data)

objectobject

Phone number information associated with the entity screening hit

[`type`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-phone-numbers-data-type)

stringstring

An enum indicating whether a phone number is a phone line or a fax line.  
  

Possible values: `phone`, `fax`

[`phone_number`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-phone-numbers-data-phone-number)

stringstring

A phone number in E.164 format.

[`urls`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-urls)

[object][object]

URLs associated with the watchlist hit

[`analysis`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-urls-analysis)

objectobject

Summary object reflecting the match result of the associated data

[`summary`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-urls-analysis-summary)

stringstring

An enum indicating the match type between data provided by user and data checked against an external data source.  
`match` indicates that the provided input data was a strong match against external data.  
`partial_match` indicates the data approximately matched against external data. For example, "Knope" vs. "Knope-Wyatt" for last name.  
`no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  
`no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  
`no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.  
  

Possible values: `match`, `partial_match`, `no_match`, `no_data`, `no_input`

[`data`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-urls-data)

objectobject

URLs associated with the entity screening hit

[`url`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-entity-watchlist-screening-hits-data-urls-data-url)

stringstring

An 'http' or 'https' URL (must begin with either of those).  
  

Format: `uri`

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-hit-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "entity_watchlist_screening_hits": [
    {
      "id": "enthit_52xR9LKo77r1Np",
      "review_status": "pending_review",
      "first_active": "2020-07-24T03:26:02Z",
      "inactive_since": "2020-07-24T03:26:02Z",
      "historical_since": "2020-07-24T03:26:02Z",
      "list_code": "EU_CON",
      "plaid_uid": "uid_3NggckTimGSJHS",
      "source_uid": "26192ABC",
      "analysis": {
        "documents": "match",
        "email_addresses": "match",
        "locations": "match",
        "names": "match",
        "phone_numbers": "match",
        "urls": "match",
        "search_terms_version": 1
      },
      "data": {
        "documents": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "type": "swift",
              "number": "C31195855"
            }
          }
        ],
        "email_addresses": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "email_address": "user@example.com"
            }
          }
        ],
        "locations": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "full": "Florida, US",
              "country": "US"
            }
          }
        ],
        "names": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "full": "Al Qaida",
              "is_primary": false,
              "weak_alias_determination": "none"
            }
          }
        ],
        "phone_numbers": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "type": "phone",
              "phone_number": "+14025671234"
            }
          }
        ],
        "urls": [
          {
            "analysis": {
              "summary": "match"
            },
            "data": {
              "url": "https://example.com"
            }
          }
        ]
      }
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/program/get`

#### Get entity watchlist screening program

Get an entity watchlist screening program

/watchlist\_screening/entity/program/get

**Request fields**

[`entity_watchlist_program_id`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-request-entity-watchlist-program-id)

requiredstringrequired, string

ID of the associated entity program.

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

/watchlist\_screening/entity/program/get

```
const request: WatchlistScreeningEntityProgramGetRequest = {
  entity_watchlist_program_id: 'entprg_2eRPsDnL66rZ7H',
};

try {
  const response = await client.watchlistScreeningEntityProgramGet(request);
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/program/get

**Response fields**

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-id)

stringstring

ID of the associated entity program.

[`created_at`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`is_rescanning_enabled`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-is-rescanning-enabled)

booleanboolean

Indicator specifying whether the program is enabled and will perform daily rescans.

[`lists_enabled`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-lists-enabled)

[string][string]

Watchlists enabled for the associated program  
  

Possible values: `CA_CON`, `EU_CON`, `IZ_SOE`, `IZ_UNC`, `IZ_WBK`, `US_CAP`, `US_FSE`, `US_MBS`, `US_SDN`, `US_SSI`, `US_CMC`, `US_UVL`, `US_SAM`, `US_TEL`, `AU_CON`, `UK_HMC`

[`name`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-name)

stringstring

A name for the entity program to define its purpose. For example, "High Risk Organizations" or "Applicants".

[`name_sensitivity`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-name-sensitivity)

stringstring

The valid name matching sensitivity configurations for a screening program. Note that while certain matching techniques may be more prevalent on less strict settings, all matching algorithms are enabled for every sensitivity.  
`coarse` - See more potential matches. This sensitivity will see more broad phonetic matches across alphabets that make missing a potential hit very unlikely. This setting is noisier and will require more manual review.  
`balanced` - A good default for most companies. This sensitivity is balanced to show high quality hits with reduced noise.  
`strict` - Aggressive false positive reduction. This sensitivity will require names to be more similar than `coarse` and `balanced` settings, relying less on phonetics, while still accounting for character transpositions, missing tokens, and other common permutations.  
`exact` - Matches must be nearly exact. This sensitivity will only show hits with exact or nearly exact name matches with only basic correction such as extraneous symbols and capitalization. This setting is generally not recommended unless you have a very specific use case.  
  

Possible values: `coarse`, `balanced`, `strict`, `exact`

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`is_archived`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-is-archived)

booleanboolean

Archived programs are read-only and cannot screen new customers nor participate in ongoing monitoring.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-program-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "id": "entprg_2eRPsDnL66rZ7H",
  "created_at": "2020-07-24T03:26:02Z",
  "is_rescanning_enabled": true,
  "lists_enabled": [
    "EU_CON"
  ],
  "name": "Sample Program",
  "name_sensitivity": "balanced",
  "audit_trail": {
    "source": "dashboard",
    "dashboard_user_id": "54350110fedcbaf01234ffee",
    "timestamp": "2020-07-24T03:26:02Z"
  },
  "is_archived": false,
  "request_id": "saKrIBuEB9qJZng"
}
```

=\*=\*=\*=

#### `/watchlist_screening/entity/program/list`

#### List entity watchlist screening programs

List all entity watchlist screening programs

/watchlist\_screening/entity/program/list

**Request fields**

[`secret`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`client_id`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`cursor`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-request-cursor)

stringstring

An identifier that determines which page of results you receive.

/watchlist\_screening/entity/program/list

```
try {
  const response = await client.watchlistScreeningEntityProgramList({});
} catch (error) {
  // handle error
}
```

/watchlist\_screening/entity/program/list

**Response fields**

[`entity_watchlist_programs`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs)

[object][object]

List of entity watchlist screening programs

[`id`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-id)

stringstring

ID of the associated entity program.

[`created_at`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-created-at)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`is_rescanning_enabled`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-is-rescanning-enabled)

booleanboolean

Indicator specifying whether the program is enabled and will perform daily rescans.

[`lists_enabled`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-lists-enabled)

[string][string]

Watchlists enabled for the associated program  
  

Possible values: `CA_CON`, `EU_CON`, `IZ_SOE`, `IZ_UNC`, `IZ_WBK`, `US_CAP`, `US_FSE`, `US_MBS`, `US_SDN`, `US_SSI`, `US_CMC`, `US_UVL`, `US_SAM`, `US_TEL`, `AU_CON`, `UK_HMC`

[`name`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-name)

stringstring

A name for the entity program to define its purpose. For example, "High Risk Organizations" or "Applicants".

[`name_sensitivity`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-name-sensitivity)

stringstring

The valid name matching sensitivity configurations for a screening program. Note that while certain matching techniques may be more prevalent on less strict settings, all matching algorithms are enabled for every sensitivity.  
`coarse` - See more potential matches. This sensitivity will see more broad phonetic matches across alphabets that make missing a potential hit very unlikely. This setting is noisier and will require more manual review.  
`balanced` - A good default for most companies. This sensitivity is balanced to show high quality hits with reduced noise.  
`strict` - Aggressive false positive reduction. This sensitivity will require names to be more similar than `coarse` and `balanced` settings, relying less on phonetics, while still accounting for character transpositions, missing tokens, and other common permutations.  
`exact` - Matches must be nearly exact. This sensitivity will only show hits with exact or nearly exact name matches with only basic correction such as extraneous symbols and capitalization. This setting is generally not recommended unless you have a very specific use case.  
  

Possible values: `coarse`, `balanced`, `strict`, `exact`

[`audit_trail`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-audit-trail)

objectobject

Information about the last change made to the parent object specifying what caused the change as well as when it occurred.

[`source`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-audit-trail-source)

stringstring

A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.  
  

Possible values: `dashboard`, `link`, `api`, `system`

[`dashboard_user_id`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-audit-trail-dashboard-user-id)

nullablestringnullable, string

ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.

[`timestamp`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-audit-trail-timestamp)

stringstring

An ISO8601 formatted timestamp.  
  

Format: `date-time`

[`is_archived`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-entity-watchlist-programs-is-archived)

booleanboolean

Archived programs are read-only and cannot screen new customers nor participate in ongoing monitoring.

[`next_cursor`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-next-cursor)

nullablestringnullable, string

An identifier that determines which page of results you receive.

[`request_id`](/docs/api/products/monitor/#watchlist_screening-entity-program-list-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

Response Object

```
{
  "entity_watchlist_programs": [
    {
      "id": "entprg_2eRPsDnL66rZ7H",
      "created_at": "2020-07-24T03:26:02Z",
      "is_rescanning_enabled": true,
      "lists_enabled": [
        "EU_CON"
      ],
      "name": "Sample Program",
      "name_sensitivity": "balanced",
      "audit_trail": {
        "source": "dashboard",
        "dashboard_user_id": "54350110fedcbaf01234ffee",
        "timestamp": "2020-07-24T03:26:02Z"
      },
      "is_archived": false
    }
  ],
  "next_cursor": "eyJkaXJlY3Rpb24iOiJuZXh0Iiwib2Zmc2V0IjoiMTU5NDM",
  "request_id": "saKrIBuEB9qJZng"
}
```

### Webhooks

=\*=\*=\*=

#### `SCREENING: STATUS_UPDATED`

Fired when an individual screening status has changed, which can occur manually via the dashboard or during ongoing monitoring.

**Properties**

[`webhook_type`](/docs/api/products/monitor/#ScreeningStatusUpdatedWebhook-webhook-type)

stringstring

`SCREENING`

[`webhook_code`](/docs/api/products/monitor/#ScreeningStatusUpdatedWebhook-webhook-code)

stringstring

`STATUS_UPDATED`

[`screening_id`](/docs/api/products/monitor/#ScreeningStatusUpdatedWebhook-screening-id)

stringstring

The ID of the associated screening.

[`environment`](/docs/api/products/monitor/#ScreeningStatusUpdatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "SCREENING",
  "webhook_code": "STATUS_UPDATED",
  "screening_id": "scr_52xR9LKo77r1Np",
  "environment": "production"
}
```

=\*=\*=\*=

#### `ENTITY_SCREENING: STATUS_UPDATED`

Fired when an entity screening status has changed, which can occur manually via the dashboard or during ongoing monitoring.

**Properties**

[`webhook_type`](/docs/api/products/monitor/#EntityScreeningStatusUpdatedWebhook-webhook-type)

stringstring

`ENTITY_SCREENING`

[`webhook_code`](/docs/api/products/monitor/#EntityScreeningStatusUpdatedWebhook-webhook-code)

stringstring

`STATUS_UPDATED`

[`entity_screening_id`](/docs/api/products/monitor/#EntityScreeningStatusUpdatedWebhook-entity-screening-id)

stringstring

The ID of the associated entity screening.

[`environment`](/docs/api/products/monitor/#EntityScreeningStatusUpdatedWebhook-environment)

stringstring

The Plaid environment the webhook was sent from  
  

Possible values: `sandbox`, `production`

API Object

```
{
  "webhook_type": "ENTITY_SCREENING",
  "webhook_code": "STATUS_UPDATED",
  "screening_id": "entscr_52xR9LKo77r1Np",
  "environment": "production"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
