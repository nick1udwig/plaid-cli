---
title: "Errors - Transfer errors | Plaid Docs"
source_url: "https://plaid.com/docs/errors/transfer/"
scraped_at: "2026-03-07T22:04:54+00:00"
---

# Transfer errors

#### Errors specific to the Transfer product

#### **TRANSFER\_NETWORK\_LIMIT\_EXCEEDED**

##### The attempted transfer exceeded the given network's amount limit.

##### Common causes

- The attempted transfer's amount exceeded the network specific limit. The error message will indicate which network limit was exceeded. For reference, see the table below.

##### Troubleshooting steps

Make sure the amount of the transfer is within the network's limit.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "TRANSFER_NETWORK_LIMIT_EXCEEDED",
 "error_message": "[network] transaction amount cannot exceed [amount]",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

| Network | Maximum Transaction Limit |
| --- | --- |
| ACH | $10,000,000.00 |
| Same Day ACH | $1,000,000.00 |
| RTP | $1,000,000.00 |
| RfP | $3,500.00 |
| Wire | $999,999.99 |

#### **TRANSFER\_ACCOUNT\_BLOCKED**

##### The transfer could not be completed because a previous transfer involving the same end-user account resulted in an error

##### Common causes

- Plaid has flagged the end-user's account as not valid for use with Transfer.

##### Troubleshooting steps

Ask the end-user to link a different account with Plaid and re-attempt the transfer with the new account.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "TRANSFER_ACCOUNT_BLOCKED",
 "error_message": "transfer was blocked due to a previous ACH return on this account",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSFER\_NOT\_CANCELLABLE**

##### The transfer could not be canceled

##### Common causes

- An attempt was made to cancel a transfer that has already been sent to the network for execution and cannot be cancelled at this stage.

##### Troubleshooting steps

Use an endpoint such as [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) to check the value of the `cancellable` property before canceling a transfer.

If applicable, contact the counterparty to the transfer and ask them to return the funds after the transfer has executed.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "TRANSFER_NOT_CANCELLABLE",
 "error_message": "transfer is not cancellable",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSFER\_UNSUPPORTED\_ACCOUNT\_TYPE**

##### An attempt was made to transfer funds to or from an unsupported account type

##### Common causes

- An attempt was made to transfer funds to or from an unsupported account type. Only checking, savings, and cash management accounts can be used with Transfer. In addition, if the transfer is a debit transfer, the account must be a debitable account. Common examples of non-debitable depository accounts include savings accounts at Chime or at Navy Federal Credit Union (NFCU).

##### Troubleshooting steps

Ask the user to link a different account to use with Transfer.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "TRANSFER_UNSUPPORTED_ACCOUNT_TYPE",
 "error_message": "transfer account type not supported",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSFER\_FORBIDDEN\_ACH\_CLASS**

##### An attempt was made to create a transfer with a forbidden ACH class (SEC code)

##### Common causes

- Your Plaid account has not been enabled for the ACH class specified in the request.
- The ACH class specified in the transfer request was incorrect.

##### Troubleshooting steps

Verify that the ACH class in the request is correct.

If you have not already done so, contact your Plaid Account Manager to request that the desired ACH class be enabled for your account.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "TRANSFER_FORBIDDEN_ACH_CLASS",
 "error_message": "specified ach_class is forbidden",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSFER\_UI\_UNAUTHORIZED**

##### The client is not enabled for the Transfer UI product

##### Common causes

- Your account is not enabled for use with Transfer UI.

##### Troubleshooting steps

[Contact Sales](https://plaid.com/contact/) or your Plaid Account Manager to request approval for Transfer products.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "TRANSFER_UI_UNAUTHORIZED",
 "error_message": "client is not authorized for transfer UI",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **TRANSFER\_ORIGINATOR\_NOT\_FOUND**

##### An association between the sender and the originator client doesn't exist

##### Common causes

- There is a typo in the `originator_client_id`
- You are a Transfer for Platforms customer (not a partner) and did not call [`/transfer/platform/originator/create`](/docs/api/products/transfer/platform-payments/#transferplatformoriginatorcreate) for this end-customer

##### Troubleshooting steps

If you are a Transfer for Platforms customer, call [`/transfer/platform/originator/create`](/docs/api/products/transfer/platform-payments/#transferplatformoriginatorcreate) to create a valid `originator_client_id` for your end-customer first.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "TRANSFER_ORIGINATOR_NOT_FOUND",
 "error_message": "the association between the sender and the originator client doesn't exist",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **INCOMPLETE\_CUSTOMER\_ONBOARDING**

##### The end customer has not completed onboarding. Their diligence status must be `approved` before moving funds

##### Common causes

- You tried to move funds for an originator that hasn't been approved by Plaid yet

##### Troubleshooting steps

Check the diligence status of this end customer by calling [`/transfer/originator/get`](/docs/api/products/transfer/platform-payments/#transferoriginatorget). Wait until the originator is marked as `approved` before attempting to move funds.

Contact your Plaid Account Manager for more specific information regarding the onboarding status of your end customer.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "INCOMPLETE_CUSTOMER_ONBOARDING",
 "error_message": "end customer has not completed onboarding. their diligence status must be `approved` before moving funds",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### **UNAUTHORIZED\_ACCESS**

##### You are not authorized to access this endpoint for your use case

##### Common causes

- You are not a Transfer for Platforms customer and tried to call a platform-only endpoint
- You are a Plaid Partner and tried to create an end-customer using [`/transfer/platform/originator/create`](/docs/api/products/transfer/platform-payments/#transferplatformoriginatorcreate) (instead of [`/partner/customer/create`](/docs/api/partner/#partnercustomercreate))

##### Troubleshooting steps

Contact your Plaid Account Manager if you believe you should have access to this route.

API error response

```
http code 400
{
 "error_type": "TRANSFER_ERROR",
 "error_code": "UNAUTHORIZED_ACCESS",
 "error_message": "you are not authorized to access this endpoint for your use case.",
 "display_message": null,
 "request_id": "HNTDNrA8F1shFEW"
}
```

#### ACH return codes

Prefer to learn by watching? A [video guide](https://www.youtube.com/watch?v=7UqU2TGLiMc) is available for this topic.

All `returned` ACH transactions will have an ACH return code. By reading the code, you can troubleshoot and debug returned transactions.

[Transfer](/docs/transfer/) customers should see [Common ACH Return Codes in the Help Center](https://support.plaid.com/hc/en-us/articles/32881797799575-What-are-the-Common-ACH-Return-Codes) for more detailed information on troubleshooting ACH return codes.

| Return Code | Description | Notes |
| --- | --- | --- |
| R01 | Insufficient funds | Available balance is not sufficient to cover the dollar amount of the debit entry |
| R02 | Account closed | A previously open account is now closed |
| R03 | No account or unable to locate account | The account number structure is valid, but the account number does not correspond to an existing account, or the name on the transaction does not match the name on the account. |
| R04 | Invalid account number | The account number fails the check digit validation or may contain an incorrect number of digits. Commonly caused by invalidated TANs; see [Tokenized account numbers](/docs/auth/#tokenized-account-numbers) for more details. |
| R05 | Unauthorized debit to consumer account | A business debit entry was transmitted to a member’s consumer account, and the member had not authorized the entry |
| R06 | Returned per ODFI's request | The ODFI has requested that the RDFI return the entry |
| R07 | Authorization revoked by customer | Member who previously authorized an entry has revoked authorization with the originator |
| R08 | Payment stopped or stop payment on item | Member had previously requested a stop payment of a single or recurring entry |
| R09 | Uncollected funds | Available balance is sufficient, but the collected balance is not sufficient to cover the entry |
| R10 | Customer advises not authorized | Member advises not authorized, notice not provided, improper source document, or amount of entry not accurately obtained from source document |
| R11 | Check truncation entry return | To be used when returning a check truncation entry |
| R12 | Branch sold to another DFI | RDFI unable to post entry destined for a bank account maintained at a branch sold to another financial institution |
| R13 | Invalid ACH routing number | Financial institution does not receive commercial ACH entries |
| R14 | Representative payee deceased or unable to continue in that capacity | Representative payee is deceased or unable to continue in that capacity, beneficiary is not deceased |
| R15 | Beneficiary of account holder deceased | Beneficiary or Account Holder Deceased |
| R16 | Account frozen | Access to account is restricted due to a specific action taken by the RDFI or by legal action |
| R17 | File record edit criteria | Fields rejected by RDFI processing (identified in return addenda) |
| R18 | Improper effective entry | Entries have been presented prior to the first available processing window for the effective date |
| R19 | Amount field error | Improper formatting of the amount field |
| R20 | Non-transaction account | Policies or regulations (such as Regulation D) prohibit or limit activity to the account indicated |
| R21 | Invalid company identification | The company ID information not valid (normally CIE entries) |
| R22 | Invalid individual ID number | Individual id used by receiver is incorrect (CIE entries) |
| R23 | Credit entry refused by receiver | Receiver returned entry because minimum or exact amount not remitted, bank account is subject to litigation, or payment represents an overpayment, originator is not known to receiver or receiver has not authorized this credit entry to this bank account |
| R24 | Duplicate entry | RDFI has received a duplicate entry |
| R25 | Addenda error | Improper formatting of the addenda record information |
| R26 | Mandatory field error | Improper information in one of the mandatory fields |
| R27 | Trace number error | Original entry trace number is not valid for return entry; or addenda trace numbers do not correspond with entry detail record |
| R28 | Routing number or check digit | Check digit for transit routing number is incorrect |
| R29 | Corporate customer advises not authorized | RDFI has been notified by business account holder that a specific transaction is unauthorized. Business accounts often automatically decline debits for security purposes. Your customer needs to inform their bank to allow debits that use your tax identifier. |
| R30 | RDFI not participant in check truncation program | Financial institution not participating in automated check safekeeping application |
| R31 | Permissible return entry | RDFI has been notified by business account holder that a specific transaction is unauthorized |
| R32 | RDFI non settlement | RDFI is not able to settle the entry |
| R33 | Return of XCK entry | RDFI determines at its sole discretion to return an XCK entry; an XCK return entry may be initiated by midnight of the sixtieth day following the settlement date if the XCK entry |
| R34 | Limited participation DFI | RDFI participation has been limited by a federal or state supervisor |
| R35 | Return of improper debit entry | ACH debit not permitted for use with the CIE standard entry class code (except for reversals) |
| R36 | Return of improper credit entry |  |
| R37 | Source Document Presented for Payment | Check used for an ARC, BOC or POP entry has also been presented for payment |
| R38 | Stop payment on source document | Stop payment has been placed on a check used for an ARC entry |
| R40 | Return of ENR entry by federal government agency (ENR only) |  |
| R41 | Invalid transaction code (ENR only) |  |
| R42 | Routing number or check digit error (ENR only) |  |
| R43 | Invalid DFI account number (ENR only) |  |
| R44 | Invalid individual ID number (ENR only) |  |
| R45 | Invalid individual name/company name (ENR only) |  |
| R46 | Invalid representative payee indicator (ENR only) |  |
| R47 | Duplicate enrollment |  |
| R50 | State law affecting RCK acceptance |  |
| R51 | Item is ineligible, notice not provided, signature not genuine |  |
| R52 | Stop payment on item |  |
| R61 | Misrouted return | Return entry was sent by RDFI to an incorrect ODFI routing/transit number |
| R62 | Incorrect trace number |  |
| R63 | Incorrect dollar amount |  |
| R64 | Incorrect individual identification |  |
| R65 | Incorrect transaction code |  |
| R66 | Incorrect company identification |  |
| R67 | Duplicate return | ODFI has received more than one return entry for the same original entry |
| R68 | Untimely return | Return entry did not meet the return deadline |
| R69 | Multiple errors |  |
| R70 | Permissible return entry not accepted |  |
| R71 | Misrouted dishonored return |  |
| R72 | Untimely dishonored return |  |
| R73 | Timely original return |  |
| R74 | Corrected return |  |
| R80 | Cross-border payment coding error |  |
| R81 | Nonparticipant in cross-border program |  |
| R82 | Invalid foreign receiving DFI identification |  |

#### RTP/RfP error codes

`failed` RTP/RfP transactions will have a failure description. By reading the description, you can troubleshoot and debug failed transactions.

This table can also be [viewed as a Google doc](https://docs.google.com/spreadsheets/d/12jHY-VZYtYqzSHQE6eEzy_I7ZbN4SLZcWTcdjoMAiuQ/edit?gid=123470926#gid=123470926).

| Failure Code | Description | Rail(s) | Condensed Error Category | Retryable | Suggested Action |
| --- | --- | --- | --- | --- | --- |
| 1100 | Other Reasons - Not specified | Both | Beneficiary or compliance information issue | No | Investigate beneficiary data or compliance issues; do not auto‑retry. |
| 9909 | Central Switch system malfunction | RTP | Network issue (retryable) | Yes - auto‑retry | Wait briefly and retry automatically; consider fallback rail if urgent. |
| 9910 | Instructed Agent signed-off | RTP | Network issue (retryable) | Yes - auto‑retry | Wait briefly and retry automatically; consider fallback rail if urgent. |
| 9912 | Recipient connection unavailable | RTP | Network issue (retryable) | Yes - auto‑retry | Wait briefly and retry automatically; consider fallback rail if urgent. |
| 9934 | Instructing Agent signed-off | RTP | Network issue (retryable) | Yes - auto‑retry | Wait briefly and retry automatically; consider fallback rail if urgent. |
| 9946 | Instructing Agent suspended | RTP | Network issue (retryable) | Yes - auto‑retry | Wait briefly and retry automatically; consider fallback rail if urgent. |
| 9947 | Instructed Agent suspended | RTP | Network issue (retryable) | Yes - auto‑retry | Wait briefly and retry automatically; consider fallback rail if urgent. |
| 9948 | Central Switch service suspended | RTP | Network issue (retryable) | Yes - auto‑retry | Wait briefly and retry automatically; consider fallback rail if urgent. |
| AC02 | Debtor account is invalid | Both | Issues with the sending account | No | Ask the sender to correct or switch accounts, then resubmit. |
| AC03 | Creditor account is invalid | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| AC04 | Account closed | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| AC06 | Account is blocked | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| AC07 | Creditor account closed | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| AC13 | Debtor account type invalid | Both | Issues with the sending account | No | Ask the sender to correct or switch accounts, then resubmit. |
| AC14 | Creditor account type invalid | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| AG01 | Transaction forbidden on this type of account | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| AG03 | Transaction type not supported on this account | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| AM02 | Transaction amount exceeds allowed maximum | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| AM04 | Insufficient funds | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| AM09 | Incorrect amount received | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| AM12 | Amount invalid or missing | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| AM13 | Amount exceeds clearing system limits | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| AM14 | Amount exceeds bank-client limit | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| BE04 | Missing or incorrect creditor address | Both | Beneficiary or compliance information issue | No | Investigate beneficiary data or compliance issues; do not auto‑retry. |
| BE06 | End customer not known | Both | Beneficiary or compliance information issue | No | Investigate beneficiary data or compliance issues; do not auto‑retry. |
| BE07 | Missing or incorrect debtor address | Both | Beneficiary or compliance information issue | No | Investigate beneficiary data or compliance issues; do not auto‑retry. |
| BE16 | Debtor identification invalid | Both | Beneficiary or compliance information issue | No | Investigate beneficiary data or compliance issues; do not auto‑retry. |
| BE17 | Creditor identification invalid | Both | Beneficiary or compliance information issue | No | Investigate beneficiary data or compliance issues; do not auto‑retry. |
| BLKD | Payment has been blocked | Both | Beneficiary or compliance information issue | No | Investigate beneficiary data or compliance issues; do not auto‑retry. |
| CH11 | Creditor identifier unknown | RFP | Authorization declined or customer action required | No | Respect customer’s decision or obtain new authorization. |
| CUST | Customer declined payment | Both | Authorization declined or customer action required | No | Respect customer’s decision or obtain new authorization. |
| DS04 | Order rejected – content issues | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| DS0H | Signer not authorized for this account | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| DS24 | Waiting time expired (incomplete order) | Both | Authorization declined or customer action required | Yes - after correction | Respect customer’s decision or obtain new authorization. |
| DT04 | Future date not supported | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| DUPL | Duplicate payment detected | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| E997 | Timeout clock has expired | Both | Network issue (retryable) | Yes - after correction | Retry the payment with a different idempotency key. |
| FF02 | Syntax error in narrative information | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| FF03 | Invalid payment type information | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| FF08 | End‑to‑End ID missing or invalid | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| INSF | Insufficient funds for outbound message | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| MD07 | End customer is deceased | Both | Beneficiary or compliance information issue | No | Investigate beneficiary data or compliance issues; do not auto‑retry. |
| NARR | Narrative reason provided | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| NOAT | Account does not support this message type | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| RC01 | Bank identifier format incorrect | Both | Routing or bank identifier issue | Yes - after correction | Supply a valid routing/BIC and retry. |
| RC02 | Bank identifier invalid or missing | Both | Routing or bank identifier issue | Yes - after correction | Supply a valid routing/BIC and retry. |
| RC03 | Debtor FI identifier invalid or missing | Both | Routing or bank identifier issue | Yes - after correction | Supply a valid routing/BIC and retry. |
| RC04 | Creditor FI identifier invalid or missing | Both | Issues with the receiving account (invalid, closed, blocked) | No | Request updated recipient account details or use another payout account. |
| SL12 | Debtor opted out of RFPs | RFP | Authorization declined or customer action required | No | Respect customer’s decision or obtain new authorization. |
| TK01 | Invalid token | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| TK02 | Sender token not found | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| TK03 | Receiver token not found | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| TK04 | Token expired | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| TK05 | Token counterparty mismatch | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| TK06 | Token value‑limit rule violation | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| TK07 | Single‑use token already used | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |
| TK08 | Token suspended | Both | Issues with payment details (amount, currency, token, duplicate, format) | Yes - after correction | Correct the payment amount, currency, token or format, then retry. |

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
