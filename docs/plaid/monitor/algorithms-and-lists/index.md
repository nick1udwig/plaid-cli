---
title: "Monitor - Watchlists and matching algorithms | Plaid Docs"
source_url: "https://plaid.com/docs/monitor/algorithms-and-lists/"
scraped_at: "2026-03-07T22:05:10+00:00"
---

# Watchlists and Matching Algorithms

#### Learn about the watchlists and matching algorithms used by Monitor

Plaid Monitor's system supports a wide variety of algorithms and languages to support common inconsistencies across both user inputs and underlying watchlist search data.

#### Supported languages

|  |  |  |
| --- | --- | --- |
| Arabic | German | Korean |
| Spanish | Simplified Chinese | Greek |
| Pashto | Thai | Traditional Chinese |
| Hungarian | Persian | Urdu |
| English | Italian | Portuguese |
| French | Japanese | Russian |
| Hebrew | Burmese | Vietnamese |

#### Supported algorithms

|  |  |
| --- | --- |
| Phonetic similarity | Jesus ↔ Heyzeus ↔ Haezoos |
| Transliteration spelling differences | Abdul Rasheed ↔ Abd al-Rashid |
| Nicknames | William ↔ Will ↔ Bill ↔ Billy |
| Missing spaces or hyphens | MaryEllen ↔ Mary Ellen ↔ Mary-Ellen |
| Titles and honorifics | Dr. ↔ Mr. ↔ Ph.D. |
| Truncated name components | McDonalds ↔ McDonald ↔ McD |
| Missing name components | Phillip Charles Carr ↔ Phillip Carr |
| Out-of-order name components | Diaz, Carlos Alfonzo ↔ Carlos Alfonzo Diaz |
| Initials | J. E. Smith ↔ James Earl Smith |
| Names split inconsistently across database fields | Dick. Van Dyke ↔ Dick Van . Dyke |
| Same name in multiple languages | Mao Zedong ↔ Мао Цзэдун ↔ 毛泽东 ↔ 毛澤東 |
| Semantically similar names | Eagle Pharmaceuticals, Inc. ↔ Eagle Drugs, Co. |
| Semantically similar names across language | Nippon Telegraph and Telephone Corporation ↔ 日本電信電話株式会社 |

##### Name sensitivity

Plaid determines a similarity score based on the result of these comparisons. Each type of matching will have a different impact on the overall similarity score, based on the importance of the discrepancy, the length and composition of the name being verified, and other variables like language. For example, "Mr. John Doe" vs. "John Doe" will receive a negligible penalty, whereas "John Paul Doe" vs "John Doe" will receive a more significant penalty.

The sensitivity you select will determine the minimum similarity score required for a match. We recommend that you start with Balanced sensitivity and adjust as needed.

If you are migrating from another identity verification solution and have configured similarity scores there, you can use the following table to roughly equate Plaid's sensitivity levels with similarity scores:

| Name sensitivity | Minimum acceptable similarity score |
| --- | --- |
| Coarse | 70 |
| Balanced | 80 |
| Strict | 90 |
| Exact | 100 |

Note that setting the scores at a comparable level between different solutions will not yield exactly the same results, since different products use different algorithms.

#### Supported watchlists

**US – Office of Foreign Assets Control**

- Specially Designated Nationals List
- Foreign Sanctions Evaders
- Palestinian Legislative Council
- Sectoral Sanctions Identifications
- Non-SDN Menu-Based Sanctions
- Correspondent Account or Payable-Through Account Sanctions
- Non-SDN Chinese Military-Industrial Complex List

**US - Department of Justice**

- FBI Wanted List

**US – Department of State**

- Nonproliferation Sanctions
- AECA Debarred
- Terrorist Exclusion List

**US – General Services Administration**

- Excluded Party List System

**Bureau of Industry and Security**

- Denied Persons List
- Unverified List

**UK – Her Majesty’s Treasury**

- Consolidated list

**EU – European External Action Service**

- Consolidated list

**AU – Department of Foreign Affairs and Trade**

- Consolidated list

**CA – Government of Canada**

- Consolidated List of Sanctions

**International**

- Interpol Red Notices for Wanted Persons List
- World Bank Listing of Ineligible Firms and Individuals
- United Nations Consolidated Sanctions

**Politically Exposed Persons**

- Politically Exposed Persons List
- State Owned Enterprise List
- CIA List of Chiefs of State and Cabinet Members

**SG - Government of Singapore**

- Terrorists and Terrorist Entities

**TR - Government of Turkey**

- Terrorist Wanted List
- Domestic Freezing Decisions
- Foreign Freezing Requests
- Weapons of Mass Destruction
- Capital Markets Board

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
