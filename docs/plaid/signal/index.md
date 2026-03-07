---
title: "Signal - Intro to Signal | Plaid Docs"
source_url: "https://plaid.com/docs/signal/"
scraped_at: "2026-03-07T22:05:18+00:00"
---

# Introduction to Signal

#### Evaluate ACH payment risk

Get started with Signal

[API Reference](/docs/api/products/signal/)[Quickstart](/docs/quickstart/)[Demo](https://plaid.coastdemo.com/share/6786ccc5a048a5f1cf748cb5?zoom=100)

#### Signal Overview

Plaid Signal is Plaid's solution for ACH risk management. With Signal, you can use Signal Rules in the Plaid Dashboard to create and manage business logic for handling transactions.

Plaid Signal includes two separate products: [Balance](/docs/balance/), which gets real-time balances; and [Signal Transaction Scores](/docs/signal/#signal-transaction-scores-overview), which uses ML modeling to assess transaction risk using over 80 attributes. You can purchase and use either Balance or Signal Transaction Scores by itself, or combine them for a more comprehensive ACH risk management approach.

Signal products simplify managing transaction risk with the no-code [Signal Rules Dashboard](https://dashboard.plaid.com/signal/risk-profiles), which allows you to easily configure risk rules and react quickly to changing trends.

#### Signal Transaction Scores overview

Prefer to learn by watching? Get an overview of how Signal Transaction Scores works in just 3 minutes!

Signal Transaction Scores applies machine learning to linked bank account data to predict the likelihood that a transaction will result in an ACH return. Signal Transaction Scores simplifies payment risk management by:

- Evaluating transactions at ultra-low latency (p95 < 2 seconds) so you can incorporate risk evaluations into critical user-present interactions, like account funding or purchase flows
- Powering management via Signal Rules, including a rule optimization platform that incorporates industry benchmarks, backtesting, and personalized rule suggestions based on your business's transaction activity, making it easy to tune your thresholds for maximized revenue
- Returning over 80 predictive insights that you can incorporate into your own risk assessment models

The Signal Platform considers over 1,000 risk factors to evaluate proposed transactions. Over time, as you use Signal Transaction Scores, it will provide more customized and refined recommendations.

Signal Transaction scores can evaluate the risk of US domestic transactions over ACH (both Standard and Same Day ACH). Signal products cannot be used to evaluate RTP or RfP transactions, debit card transactions, transactions to or from a non-US bank account, or wire transfers. For these use cases, use [Balance](/docs/balance/) with [`/accounts/balance/get`](/docs/api/products/signal/#accountsbalanceget) instead. For more details, see [Signal Transaction Scores vs. Balance comparison chart](/docs/payments/#balance-and-signal-transaction-scores-comparison).

#### How Signal Transaction Scores works

First, Plaid analyzes and summarizes the level of risk a transaction poses into a [risk score](/docs/signal/#signal-transaction-scores) called a Signal Transaction Score.

Next, a [ruleset](/docs/signal/#signal-rulesets) is applied to turn these scores into actions. You must create and tune these rulesets to match your business's risk tolerance.

##### Signal Transaction Scores

When you call [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate), Plaid generates a score for a proposed transaction, predicting the likelihood of returns due to insufficient funds, closed or frozen accounts, and other administrative bank returns, as well as consumer authorized returns. A higher score indicates a greater likelihood that the transaction will result in an ACH return.

##### Signal rulesets

To turn these risk scores into an action, you will configure a [Signal Ruleset](/docs/signal/signal-rules/). Plaid can suggest an initial set of rules to approve payments below a certain score threshold. Signal Transaction Scores provides you with access to both the aggregate score and over 80 predictive insights, allowing you to set up simple score-based rulesets or to create more complex rules.

![Image of example distribution of Plaid Signal Transaction Scores](/assets/img/docs/signal/performance-page.png)

The shape of this graph (how many transactions are low risk versus high risk) is unique per customer. [Reporting returns](/docs/signal/reporting-returns/) will allow the Signal Dashboard to provide you with personalized recommendations for [adjusting score thresholds](/docs/signal/signal-rules/).

It is recommended that you roll out Signal Transaction Scores in stages to collect data, then use the Dashboard tools to adjust your approval logic at each rollout phase.

#### Integration overview

1. [Create a new Item with Signal](/docs/signal/creating-signal-items/) or [add Signal to an existing Item](/docs/signal/creating-signal-items/#adding-signal-to-existing-items).
2. [Create a Signal ruleset](/docs/signal/signal-rules/) using the Dashboard.
3. Call [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate) and [determine the next steps based on results](/docs/signal/signal-rules/#using-signal-ruleset-results).
4. [Report ACH returns and decisions](/docs/signal/reporting-returns/) to Plaid.
5. After launch, periodically review and [tune your Signal Rules](/docs/signal/tuning-rules/) using the Dashboard.

#### Billing

Signal Transaction Scores is billed on a per-request fee basis based on the number of calls to [`/signal/evaluate`](/docs/api/products/signal/#signalevaluate). For more details, see [per-request flat fee billing](/docs/account/billing/#per-request-flat-fee).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
