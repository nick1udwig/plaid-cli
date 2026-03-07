---
title: "API - Program Metrics | Plaid Docs"
source_url: "https://plaid.com/docs/api/products/transfer/metrics/"
scraped_at: "2026-03-07T22:04:23+00:00"
---

# Program Metrics

#### API reference for Transfer program metrics

For how-to guidance, see the [Transfer documentation](/docs/transfer/).

| Program Metrics |  |
| --- | --- |
| [`/transfer/metrics/get`](/docs/api/products/transfer/metrics/#transfermetricsget) | Get transfer product usage metrics |
| [`/transfer/configuration/get`](/docs/api/products/transfer/metrics/#transferconfigurationget) | Get transfer product configuration |

=\*=\*=\*=

#### `/transfer/metrics/get`

#### Get transfer product usage metrics

Use the [`/transfer/metrics/get`](/docs/api/products/transfer/metrics/#transfermetricsget) endpoint to view your transfer product usage metrics.

/transfer/metrics/get

**Request fields**

[`client_id`](/docs/api/products/transfer/metrics/#transfer-metrics-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/metrics/#transfer-metrics-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`originator_client_id`](/docs/api/products/transfer/metrics/#transfer-metrics-get-request-originator-client-id)

stringstring

The Plaid client ID of the transfer originator. Should only be present if `client_id` is a [Platform customer](https://plaid.com/docs/transfer/application/#originators-vs-platforms).

/transfer/metrics/get

```
const request: TransferMetricsGetRequest = {
  originator_client_id: '61b8f48ded273e001aa8db6d',
};

try {
  const response = await client.transferMetricsGet(request);
} catch (error) {
  // handle error
}
```

/transfer/metrics/get

**Response fields**

[`request_id`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`daily_debit_transfer_volume`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-daily-debit-transfer-volume)

stringstring

Sum of dollar amount of debit transfers in last 24 hours (decimal string with two digits of precision e.g. "10.00").

[`daily_credit_transfer_volume`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-daily-credit-transfer-volume)

stringstring

Sum of dollar amount of credit transfers in last 24 hours (decimal string with two digits of precision e.g. "10.00").

[`monthly_debit_transfer_volume`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-monthly-debit-transfer-volume)

stringstring

Sum of dollar amount of debit transfers in current calendar month (decimal string with two digits of precision e.g. "10.00").

[`monthly_credit_transfer_volume`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-monthly-credit-transfer-volume)

stringstring

Sum of dollar amount of credit transfers in current calendar month (decimal string with two digits of precision e.g. "10.00").

[`iso_currency_code`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-iso-currency-code)

stringstring

The currency of the dollar amount, e.g. "USD".

[`return_rates`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-return-rates)

nullableobjectnullable, object

Details regarding return rates.

[`last_60d`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-return-rates-last-60d)

nullableobjectnullable, object

Details regarding return rates.

[`overall_return_rate`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-return-rates-last-60d-overall-return-rate)

stringstring

The overall return rate.

[`unauthorized_return_rate`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-return-rates-last-60d-unauthorized-return-rate)

stringstring

The unauthorized return rate.

[`administrative_return_rate`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-return-rates-last-60d-administrative-return-rate)

stringstring

The administrative return rate.

[`authorization_usage`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-authorization-usage)

nullableobjectnullable, object

Details regarding authorization usage.

[`daily_credit_utilization`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-authorization-usage-daily-credit-utilization)

stringstring

The daily credit utilization formatted as a decimal.

[`daily_debit_utilization`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-authorization-usage-daily-debit-utilization)

stringstring

The daily debit utilization formatted as a decimal.

[`monthly_credit_utilization`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-authorization-usage-monthly-credit-utilization)

stringstring

The monthly credit utilization formatted as a decimal.

[`monthly_debit_utilization`](/docs/api/products/transfer/metrics/#transfer-metrics-get-response-authorization-usage-monthly-debit-utilization)

stringstring

The monthly debit utilization formatted as a decimal.

Response Object

```
{
  "daily_debit_transfer_volume": "1234.56",
  "daily_credit_transfer_volume": "567.89",
  "monthly_transfer_volume": "",
  "monthly_debit_transfer_volume": "10000.00",
  "monthly_credit_transfer_volume": "2345.67",
  "iso_currency_code": "USD",
  "request_id": "saKrIBuEB9qJZno",
  "return_rates": {
    "last_60d": {
      "overall_return_rate": "0.1023",
      "administrative_return_rate": "0.0160",
      "unauthorized_return_rate": "0.0028"
    }
  },
  "authorization_usage": {
    "daily_credit_utilization": "0.2300",
    "daily_debit_utilization": "0.3401",
    "monthly_credit_utilization": "0.9843",
    "monthly_debit_utilization": "0.3220"
  }
}
```

=\*=\*=\*=

#### `/transfer/configuration/get`

#### Get transfer product configuration

Use the [`/transfer/configuration/get`](/docs/api/products/transfer/metrics/#transferconfigurationget) endpoint to view your transfer product configurations.

/transfer/configuration/get

**Request fields**

[`client_id`](/docs/api/products/transfer/metrics/#transfer-configuration-get-request-client-id)

stringstring

Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.

[`secret`](/docs/api/products/transfer/metrics/#transfer-configuration-get-request-secret)

stringstring

Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.

[`originator_client_id`](/docs/api/products/transfer/metrics/#transfer-configuration-get-request-originator-client-id)

stringstring

The Plaid client ID of the transfer originator. Should only be present if `client_id` is a [Platform customer](https://plaid.com/docs/transfer/application/#originators-vs-platforms).

/transfer/configuration/get

```
const request: TransferConfigurationGetRequest = {
  originator_client_id: '61b8f48ded273e001aa8db6d',
};

try {
  const response = await client.transferConfigurationGet(request);
} catch (error) {
  // handle error
}
```

/transfer/configuration/get

**Response fields**

[`request_id`](/docs/api/products/transfer/metrics/#transfer-configuration-get-response-request-id)

stringstring

A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.

[`max_single_transfer_credit_amount`](/docs/api/products/transfer/metrics/#transfer-configuration-get-response-max-single-transfer-credit-amount)

stringstring

The max limit of dollar amount of a single credit transfer (decimal string with two digits of precision e.g. "10.00").

[`max_single_transfer_debit_amount`](/docs/api/products/transfer/metrics/#transfer-configuration-get-response-max-single-transfer-debit-amount)

stringstring

The max limit of dollar amount of a single debit transfer (decimal string with two digits of precision e.g. "10.00").

[`max_daily_credit_amount`](/docs/api/products/transfer/metrics/#transfer-configuration-get-response-max-daily-credit-amount)

stringstring

The max limit of sum of dollar amount of credit transfers in last 24 hours (decimal string with two digits of precision e.g. "10.00").

[`max_daily_debit_amount`](/docs/api/products/transfer/metrics/#transfer-configuration-get-response-max-daily-debit-amount)

stringstring

The max limit of sum of dollar amount of debit transfers in last 24 hours (decimal string with two digits of precision e.g. "10.00").

[`max_monthly_credit_amount`](/docs/api/products/transfer/metrics/#transfer-configuration-get-response-max-monthly-credit-amount)

stringstring

The max limit of sum of dollar amount of credit transfers in one calendar month (decimal string with two digits of precision e.g. "10.00").

[`max_monthly_debit_amount`](/docs/api/products/transfer/metrics/#transfer-configuration-get-response-max-monthly-debit-amount)

stringstring

The max limit of sum of dollar amount of debit transfers in one calendar month (decimal string with two digits of precision e.g. "10.00").

[`iso_currency_code`](/docs/api/products/transfer/metrics/#transfer-configuration-get-response-iso-currency-code)

stringstring

The currency of the dollar amount, e.g. "USD".

Response Object

```
{
  "max_single_transfer_amount": "",
  "max_single_transfer_credit_amount": "1000.00",
  "max_single_transfer_debit_amount": "1000.00",
  "max_daily_credit_amount": "50000.00",
  "max_daily_debit_amount": "50000.00",
  "max_monthly_amount": "",
  "max_monthly_credit_amount": "500000.00",
  "max_monthly_debit_amount": "500000.00",
  "iso_currency_code": "USD",
  "request_id": "saKrIBuEB9qJZno"
}
```

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
