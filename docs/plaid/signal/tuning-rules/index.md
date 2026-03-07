---
title: "Signal - Tuning the Signal Rules | Plaid Docs"
source_url: "https://plaid.com/docs/signal/tuning-rules/"
scraped_at: "2026-03-07T22:05:19+00:00"
---

# Tuning your Signal Rules

#### Learn about tuning your Signal Rules logic

#### Overview

The set of initial rules Plaid suggests when creating a ruleset is meant to be a starting point. Tuning the logic within your [Signal Rules Dashboard](https://dashboard.plaid.com/signal/risk-profiles) is an important step in achieving the right outcomes for your business.

This guide outlines how to incrementally and systematically tune your ruleset to optimize payment acceptances while managing risk. This guide assumes you have integrated using [Signal Rules](/docs/signal/signal-rules/).

##### Step 1: Confirm your ruleset has enough data

Before you get started adjusting rules, ensure your ACH returns are [reported to Plaid](https://plaid.com/docs/signal/reporting-returns/#reporting-returns). You can tune your rules at any time, but having significant data (i.e. at least hundreds of ACH returns) will be most effective.

##### Step 2: Review performance of current ruleset

To review an active ruleset’s performance, locate the [ruleset](https://dashboard.plaid.com/signal/risk-profiles), and click into it. Review the metrics shown for accepted, reviewed, and rerouted transactions, as well as your ACH return rate. Adjust the timeframe in the top right to relevant traffic.

![Performance page of a ruleset](/assets/img/docs/signal/tuning-rules/signal-rules-tuning-overall-performance.png)

Performance page of a ruleset

##### Step 3: Set a tuning objective

After analyzing the current performance, establish the right goal for your business (usually either reducing returned transactions, or increasing accepted transactions).

For example, do you want to:

- *Reduce* ACH return rates?
- *Increase* ACH acceptance rates?

As with all risk-based approaches in payments, there is a fundamental tradeoff between the percentage of payments accepted, and the amount of risk taken on in doing so. If you want to increase your acceptance rate, you’ll likely see a higher percentage of returned transactions. If you want to reduce the percentage of returned transactions, this often means lowering your acceptance rate.

##### Step 4: Choose a rule to adjust

Once you’ve established your goal, it’s time to try editing a rule. We recommend starting by selecting a rule that uses the Bank Score as a parameter and `REROUTE` as the result. By lowering the risk score threshold needed to `REROUTE` the payment, you’ll lower your acceptances, and your ACH return rate. By raising the risk score threshold needed to `REROUTE` the payment, you’ll increase your acceptance rate and your return rate.

You can check which rules are matching the most traffic from the Ruleset rule page. The number of transactions that matched each rule, since the last revision of the ruleset was published, is listed under the ‘Transactions’ column.

![Ruleset overview page, showing transaction counts](/assets/img/docs/signal/tuning-rules/signal-rules-tuning-ruleset.png)

Ruleset overview page, showing transaction counts

To see the percentage of transactions that currently match a rule, click “edit” on the rule. The right-hand side of the edit screen will show “Matching Transactions.” In the example below, we can see that 1.2% of transactions are deemed as too high risk to proceed, and thus `REROUTE`’d to another payment method.

![Rule detail page](/assets/img/docs/signal/tuning-rules/signal-rules-tuning-bank-score-rule.png)

Rule detail page

##### Step 5: Tune the rule

Next, adjust the risk score threshold for the rule. Use the graph icon 📈 beside the rule conditions to understand the distribution of your traffic with a given score.

In this view, we can see estimates on how much traffic is assigned scores near this datapoint. Drag the slider from right to left to see estimates detailing the percentage of traffic that would be `REROUTE`d. In particular, it’s useful to see red transactions in this graph. These are transactions that previously were processed but resulted in an ACH return.

![Performance page, showing distribution of scores](/assets/img/docs/signal/tuning-rules/signal-rules-tuning-graph-performance.png)

Performance page, showing distribution of scores

If you are looking to reduce ACH returns, you should lower the risk score threshold in the REROUTE rule (example: go from 85 -> 80). This will mean *more* transactions are rerouted. If you are looking to increase acceptances, you should *raise* the risk score threshold (example: go from 85 to 90).

Once you have found a viable new risk score threshold to use, save it in the rule.

##### Step 6: Save and review adjustment impact

Once you have found a threshold that looks compelling, “Save” this working edit to return the ruleset overview page. Here, you can get an estimate of the overall impact of all staged edits. Look for projected changes in accepted transactions, and return rate.

![Ruleset overview page showing backtested results](/assets/img/docs/signal/tuning-rules/signal-rules-tuning-delta-view.png)

Ruleset overview page showing backtested results

For the example rule above, you can see if you publish this change, you will `REROUTE` more transactions to other payment methods, but in return you’ll lower your returns meaningfully. If this is a worthwhile tradeoff for your business, click “Review and publish” to push these edits to your live traffic.

##### Step 7: Publish the rule

Click “Review and publish” to push edits to live traffic. This will take effect immediately.

![Dialog that allows you to push the edits live](/assets/img/docs/signal/tuning-rules/signal-rules-tuning-publish-edits.png)

Dialog that allows you to push the edits live

##### Step 8: Observe results over time

Continue monitoring rule performance for return rate and acceptance rate shifts. Plaid recommends visiting the dashboard regularly to measure the effects of published changes to rules. Iterate and adjust your thresholds as needed.

**Note on increasing acceptance rate** - If you are increasing the percentage of accepted transactions, this means you are going to process transactions that would have previously been rerouted to another payment method. Because those prior transactions were not processed, and thus there was no insight on if they resulted in an ACH return, Plaid cannot provide a concrete estimated return rate. For this reason, it’s recommended to increase your acceptance rate slowly over a period of time.

#### Advanced / Larger-scale rule tuning

Once you’ve observed over 5,000 ACH returns, Plaid can more intelligently analyze your traffic on your behalf, such as finding what additional datapoints may strongly correlate with risk for your particular traffic, and provide updated rules and risk thresholds.

This process is currently in beta and is offered with your existing Signal Transaction Scores pricing. Please [file a support ticket](https://dashboard.plaid.com/support/new/admin/account-administration) if you have enough traffic and would like to use it.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
