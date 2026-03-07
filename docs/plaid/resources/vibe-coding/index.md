---
title: "Resources - Vibe coding tips | Plaid Docs"
source_url: "https://plaid.com/docs/resources/vibe-coding/"
scraped_at: "2026-03-07T22:05:16+00:00"
---

# Vibe coding tips

#### Get the best results when using AI tools with Plaid

When using AI coding tools like Cursor or Claude Code with Plaid, following a few best practices can help you get more accurate results.

#### Ensure you have the latest Plaid client library and SDK installed

If using an AI tool to set up your environment, manually double-check to make sure that the most recent [Plaid client library](/docs/api/libraries/) and, if applicable, the most recent [frontend SDK](/docs/api/libraries/#link-client-sdks) is installed. AI tools love to install outdated versions of these libraries!

When an outdated Plaid library or SDK is installed, not only will you not be able to use the latest Plaid products and features, but the AI tool will tend to use the locally installed documentation, which is itself outdated, and will not discover new products and features.

#### Explicitly include the specific docs pages you want to use in your context

The Plaid docs site is too large to fit in any commercially available LLM's context windows. You may need to prompt the tool with the specific URLs you want to read docs from, rather than simply having it crawl <https://plaid.com/docs>. For best results, use the [LLM-friendly documentation](/docs/resources/#llm-friendly-documentation).

#### Use the integration overviews, Launch Checklist, and/or vibe coding guides to create a high-level task list

Most products' documentation includes a high-level integration overview list of steps. ([Example for Plaid Check Consumer Report](/docs/check/#standard-integration-flow), [example for Transactions](/docs/transactions/#integration-overview)). Once you have Production access, a more comprehensive and personalized list of steps can be found in the [Launch Center](https://dashboard.plaid.com/developers/launch-center). For Signal Transaction Scores, Transfer, and Transactions, Plaid also has [vibe-coding guides](https://github.com/plaid/ai-coding-toolkit/tree/main/rules), designed to be used as input to an AI assistant to help it implement these integrations correctly.

Giving these steps to an AI tool can help make sure it implements the right steps. (Of course, none of this is a substitute for making sure you understand the high level steps yourself.)

#### Use sample apps, requests, and responses as context

Providing your app with existing, known-working sample code from the [Plaid public GitHub](https://github.com/plaid) or sample requests and responses from the [API Reference](https://plaid.com/docs/api/) or [Link docs](/docs/link/) to use as context helps LLMs to better discover the correct interfaces for Plaid API calls. Providing a sample response as context can be particularly helpful when working with endpoints that have complex and deeply nested responses, like those used by Income or Consumer Report. (Providing the OpenAPI specification as context does not work as well as providing a JSON response or code sample.)

#### Don't forget Dashboard setup steps

Many Plaid products require configuration steps in the Dashboard that are not currently exposed to MCP servers and can't be configured by AI. For products like Identity Verification, Monitor, Layer, Signal Transaction Scores, Balance, and Protect, you will need to use the Dashboard to set up rules and/or templates in order for your integration to work.

#### Ensure you are using the correct test data

Several Plaid products, especially Identity Verification, require specific test inputs and configurations in Sandbox to receive a successful result, a fact which AI coding tools often miss. Make sure to read the testing documentation for the product you're building and that the test data sent by your app is consistent with the test data expected by Plaid.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
