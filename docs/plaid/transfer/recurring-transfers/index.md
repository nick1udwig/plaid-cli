---
title: "Transfer - Recurring transfers | Plaid Docs"
source_url: "https://plaid.com/docs/transfer/recurring-transfers/"
scraped_at: "2026-03-07T22:05:25+00:00"
---

# Recurring Transfers

#### Learn how to set up and use recurring ACH transactions

Recurring transfers allow you to automatically originate fixed amount ACH transactions with a regular interval according to a schedule you define. Plaid currently supports intervals with an arbitrary number of weeks or months.

Once you set up the recurring transfer, Plaid automatically originates the ACH transaction on the planned date, defined by the recurring schedule. You can look up and cancel recurring transfers. You may also receive updates about the recurring transfer itself, as well as each individual transfer originated by the recurring transfer.

Recurring transfers cannot be used with [Platform payments](/docs/transfer/platform-payments/).

#### Creating a recurring transfer

Before creating a recurring transfer, you should be familiar with creating one-time transfers. See [Creating Transfers](/docs/transfer/creating-transfers/).

Use [`/transfer/recurring/create`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcreate) to create a new recurring transfer. The request body is very similar to [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate), except that you will provide an additional `schedule` parameter.

The `schedule` defines the start date and the recurring interval. Optionally, you may provide an end date, which is the last day, inclusively, that an ACH transaction can be originated. If the end date is not set, the recurring transfer remains active until it is canceled.

The recurrence interval is calculated by multiplying `interval_unit` by `interval_count.` For example, to set up a recurring transfer that's originated once every 2 weeks, set interval\_unit = "week" and interval\_count = 2.

The `interval_execution_day` parameter indicates which day in the week or month that you expect a new ACH transaction to be originated.

For a weekly recurring schedule, `interval_execution_day` should be an integer ranging from 1 to 5, representing Monday through Friday.

For a monthly recurring schedule, `interval_execution_day` should be either
a positive integer ranging from 1 to 28, indicating the 1st through 28th day of the month;
or a negative integer ranging from -1 to -5, where -1 is the last day of the month, -2 is the second-to-last day of the month, and so on.

For example, a transfer with `interval_execution_day` of `-1` in 2025 would be executed on January 31, February 28, March 31, April 30, June 2 (because the last day of May 2025 is a Saturday; see [Weekend and bank holiday adjustment](/docs/transfer/recurring-transfers/#weekend-and-bank-holiday-adjustment)), June 30, etc.

#### Canceling a recurring transfer

To cancel a recurring transfer, use [`/transfer/recurring/cancel`](/docs/api/products/transfer/recurring-transfers/#transferrecurringcancel) and provide the `recurring_transfer_id` of the recurring transfer you wish to cancel.

If you cancel a recurring transfer on the same day that a new ACH transaction is supposed to be originated, it is not guaranteed that this transaction can be canceled.

Once a recurring transfer has been cancelled, the `status` field associated with the `recurring_transfer_id` will be `cancelled`.

#### Execution of scheduled recurring transfers

Before each instance of a recurring transfer is executed, Plaid will automatically perform the same authorization check performed by [`/transfer/authorization/create`](/docs/api/products/transfer/initiating-transfers/#transferauthorizationcreate). If the check succeeds, the transfer will proceed, and Plaid will notify you via a [`RECURRING_NEW_TRANSFER`](/docs/api/products/transfer/recurring-transfers/#recurring_new_transfer) webhook. If the check fails, the transfer will not be executed, and Plaid will notify you via a [`RECURRING_TRANSFER_SKIPPED`](/docs/api/products/transfer/recurring-transfers/#recurring_transfer_skipped) webhook.

If a recurring transfer instance is skipped due to a failed authorization check, it will not be retried. Subsequent transfers will be attempted as normal.

Once the last scheduled instance of a recurring transfer has been executed (or attempted to be executed), the `status` field associated with the `recurring_transfer_id` will move from `active` to `expired`.

##### Weekend and bank holiday adjustment

If the planned origination date falls on a weekend or a bank holiday, Plaid automatically adjusts it to the next available banking day, provided that the adjusted date is on or before the schedule's `end_date`.

This means if the origination date after adjustment falls after the recurring schedule's `end_date`, it will not be originated. We recommend you always choose a banking day as the `end_date` in the schedule if it's needed.

#### Receiving updates on a recurring transfer

Plaid sends following webhook events regarding a recurring transfer:

- [`RECURRING_NEW_TRANSFER`](/docs/api/products/transfer/recurring-transfers/#recurring_new_transfer) when a new ACH transaction is originated on the planned date.
- [`RECURRING_TRANSFER_SKIPPED`](/docs/api/products/transfer/recurring-transfers/#recurring_transfer_skipped) when Plaid is unable to originate a new ACH transaction on the planned date due to a failed authorization, such the account having insufficient funds.
- [`RECURRING_CANCELLED`](/docs/api/products/transfer/recurring-transfers/#recurring_new_transfer) when the recurring transfer is cancelled by Plaid.
- For all transfers created through recurring transfer, [`TRANSFER_EVENTS_UPDATE`](/docs/api/products/transfer/reading-transfers/#transfer_events_update) webhook events are also sent so that you can receive updates on each individual ACH transaction.

The transfers created through recurring transfer appear in the response of [`/transfer/list`](/docs/api/products/transfer/reading-transfers/#transferlist) and can be queried by [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) given a `transfer_id`. The `recurring_transfer` object also has a `transfer_ids` field, containing the ids of the transfers originated by this recurring transfer. To get details of transfer instances created from a recurring transfer, you can call [`/transfer/recurring/get`](/docs/api/products/transfer/recurring-transfers/#transferrecurringget) on a schedule and store the new `transfer_id`s, then call [`/transfer/get`](/docs/api/products/transfer/reading-transfers/#transferget) on those id(s) to retrieve the transfer details.

#### Recurring transfers and PNC TAN expiration

If you have a recurring transfer that lasts for over a year and where the end user's account is held at PNC Bank, you may need to send the Item through [update mode](/docs/link/update-mode/) at least once a year to avoid disruption of the transfer schedule. For more details, see [PNC TAN expiration](/docs/auth/#pnc-tan-expiration).

#### Testing recurring transfers

See [Simulating recurring transfers](/docs/transfer/sandbox/#simulating-recurring-transfers).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
