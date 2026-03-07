---
title: "Resources | Plaid Docs"
source_url: "https://plaid.com/docs/resources/"
scraped_at: "2026-03-07T22:05:16+00:00"
---

# Plaid resources

#### External resources to help build your Plaid integration

In addition to the docs here, Plaid has a collection of tools, sample code,
SDKs, and videos to help you build your Plaid integration.

#### Sample apps

Our [GitHub](https://github.com/plaid) contains a number of sample apps to help
you get started using Plaid!

**Quickstart**

[![](/assets/img/docs/quickstart.jpeg)](https://github.com/plaid/quickstart)

Official Plaid Quickstart, using React, available in NodeJS, Python, Ruby, Java, and Go

**Tiny Quickstart**

[![](/assets/img/docs/tiny-quickstart.png)](https://github.com/plaid/tiny-quickstart)

A Plaid implementation in 150 lines. Available in JavaScript, Next.Js, React, and React Native frontends.

**Personal Finances**

[![](/assets/img/docs/pattern-square.png)](https://github.com/plaid/pattern)

Usages of Transactions, Assets, and Liabilities APIs, using NodeJS and React

**Account Funding**

[![](/assets/img/docs/pattern_account_funding_screenshot.png)](https://github.com/plaid/pattern-account-funding/)

Auth, Identity, and Balance APIs for money movement, in NodeJS and React

**Account Funding in Europe**

[![](/assets/img/docs/pattern_account_funding_europe_screenshot.png)](https://github.com/plaid/payment-initiation-pattern-app/)

Account funding using the Payment Initiation API, using NodeJS and React

**Transfer Quickstart**

[![](/assets/img/docs/transfer-quickstart-screenshot.png)](https://github.com/plaid/transfer-quickstart)

Use Transfer to receive funds from customers, using NodeJS and JavaScript

**Identity Verification Quickstart**

[![](/assets/img/docs/idv_quickstart_screenshot.png)](https://github.com/plaid/idv-quickstart)

See Plaid Identity Verification in action! (NodeJS only)

**Layer Quickstart**

[![](/assets/img/docs/layer-quickstart.png)](https://github.com/plaid/layer-quickstart)

Use Layer to quickly onboard customers, using NodeJS and React

**Income**

[![](/assets/img/docs/income_sample_screenshot.png)](https://github.com/plaid/income-sample)

A lending app using the Income API, using NodeJS and React

#### Demos, video tutorials, and guides

##### Demos

[demo.plaid.com](https://demo.plaid.com) features interactive demos showing API calls alongside the end user experience for many Plaid products and use cases, including account funding and verification, personal financial management with Transactions, credit and underwriting, pay-by-bank, and more.

For the older demo page that just shows the Link UI rather than a full integration experience, see [Link demo](https://plaid.com/link-demo).

##### Video tutorials

[Plaid Academy](https://www.youtube.com/playlist?list=PLyKH4ZiEQ1bH5wpCt9SiyVfHlV2HecFBq) features detailed tutorial walkthroughs for common use cases and key concepts, several with accompanying repos.

[Plaid in 3 Minutes](https://www.youtube.com/playlist?list=PLyKH4ZiEQ1bE7XBcpX81BQWLy1olem1wf) contains overviews of different Plaid APIs, as well as short tutorials for busy people.

[Plaid Quick Tips](https://www.youtube.com/playlist?list=PLyKH4ZiEQ1bEAnnIiFMRGat7tFQl3gcJv) has snack-sized solutions to pressing Plaid problems.

The playlist below contains all of Plaid's up-to-date educational videos.

##### Launch center

The [Launch Center](https://dashboard.plaid.com/developers/launch-center) provides a list of steps to help
make sure your app is Plaid-optimized and production ready.

#### Developer tools

Developer tools include our SDKs, Postman collection, and OpenAPI file.

##### SDKs

For information about official and community supported libraries for both the
Plaid API and the Link frontend component, see [libraries](/docs/api/libraries/).

##### Postman

The [Postman Collection](https://github.com/plaid/plaid-postman) provides an
easy, no-code way to explore the API.

##### OpenAPI file

The [Plaid OpenAPI file](https://github.com/plaid/plaid-openapi) is available on
GitHub. This file describes the Plaid API in a standardized way suitable for
usage with tools for testing, client library generation, and other purposes that
ingest OpenAPI/Swagger definition files.

##### Test data libraries

Plaid maintains a repo of
[configuration objects](https://github.com/plaid/sandbox-custom-users) suitable
for usage as [custom Sandbox users](https://plaid.com/docs/sandbox/user-custom/)
to help you test apps in Sandbox.

##### Enrich test transaction library

For testing Enrich, Plaid provides a list of [Sandbox-compatible Enrich transactions](https://plaid.com/documents/enrich_sandbox_preset_transactions.csv).

##### Credit attributes library

For the Assets and Income products, Plaid maintains a
[Credit attributes library](https://github.com/plaid/credit-attributes) of helper
scripts that can be used to derive various useful attributes from an Asset
Report or Bank Income.

#### AI Developer Toolkit

Plaid's AI toolkit includes a number of tools to help you incorporate AI tooling into your development flow.

##### Vibe coding tips

Our [vibe coding tips](/docs/resources/vibe-coding/) article contains recommendations on how to get your favorite AI coding assistant or IDE to work best with Plaid, including common mistakes AI tools will make and how to fix them.

##### Dashboard MCP Server

Plaid's [Dashboard MCP server](/docs/resources/mcp/) allows you to access Plaid's diagnostics
and analytics tools to retrieve information about Link analytics, Item health, and more,
using your favorite LLM-powered application or client library.

##### LLM-friendly documentation

In accordance with the [llms-txt](https://llmstxt.org/) proposal, an index of all
of Plaid's documentation can be found in LLM-friendly Markdown at
<https://plaid.com/llms.txt>. You can find a Markdown equivalent of
Plaid documentation pages or API reference sections by clicking the document icon next to the header. You can also access Markdown pages by adding `index.html.md` to the end of the URL. For example, a Markdown version of <https://plaid.com/docs/auth> can be found at
<https://plaid.com/docs/auth/index.html.md>.

Consult the documentation for your favorite LLM-powered coding assistant
for the best way to incorporate this documentation into your workflow, as best
practices are changing frequently. Often, adding a prompt like the following
in your helper file can be useful.

Sample prompt

```
## Instructions on using the Plaid API

For instructions on how to use the Plaid API, please go to 
https://plaid.com/docs/llms.txt. There you can find a list of other
documentation pages that you can retrieve to obtain the necessary 
information. If you need to search for additional documentation, you
should first try to use a link that is listed in the llms.txt file.
```

A single file containing all of Plaid's documentation can be found
at <https://plaid.com/docs/llms-full.txt>, although it is generally too large
to fit into most models' context windows.

##### "Vibe Coding" guides

Trying to prompt your favorite LLM-powered coding tool to build a Plaid-powered
application? Plaid provides several [detailed guides](https://github.com/plaid/ai-coding-toolkit/tree/main/rules) you
can reference within your LLM coding session that will provide your agent with
additional context. You can use these guides to more quickly build proof of concept
applications, or to kick-start your development process.

##### Sandbox MCP server

Plaid also provides a [Sandbox MCP server](https://github.com/plaid/ai-coding-toolkit/tree/main/sandbox),
which, when run locally, provides helpful utilities for building and testing your Plaid-powered applications.
You can use the Sandbox MCP server to generate realistic looking [custom user data](/docs/sandbox/user-custom/),
simulate webhooks, or create public tokens in the [Sandbox environment](/docs/sandbox/).

##### Ask Bill

[Bill](https://plaid.com/docs/support/?showChat=true) is our helpful robot platypus that reads our docs for fun! You can [ask Bill](https://plaid.com/docs/support/?showChat=true) any time you need help with the Plaid API. Remember that Bill (like all of us) isn't perfect, even though he tries his best -- he can sometimes make mistakes, and he can only provide answers based on Plaid's published documentation and other publicly available information sources.

#### Community

You can join the
[Plaid Developer Community on Discord](https://discord.gg/sf57M8DW3y) or ask
questions on
[Stack Overflow](https://stackoverflow.com/questions/tagged/plaid?tab=Newest).

#### Integration support

For small businesses and startups, [HumbleDevs](https://www.humbledevs.com/) is a popular Plaid integration partner offering affordable integration solutions. HumbleDevs offers several [pre-specced Plaid solutions packages](https://v1.humbledevs.com/solutions) with known timelines and costs for popular Plaid-powered solutions, and can also build a custom Plaid integration.

For larger businesses or those with other requirements, Plaid's partnerships team will work with you to understand your needs and recommend a partner that is a good fit for your business. For more information, contact your account manager or the [sales team](https://plaid.com/contact/). Or, if you already have production access, you can [contact Support](https://dashboard.plaid.com/support/new/product-and-development/developer-lifecycle/developer-resources) and selecting "Building with Plaid" followed by "Plaid Partner Implementation Assistance".

At the enterprise level, Plaid also offers integration packages. For more details, [contact Sales](https://plaid.com/contact/) or your account manager.

#### Partner directory

For more information about Plaid partners, see the
[Partner directory](https://plaid.com/partner-directory/).

The Partner directory is not a comprehensive list of all Plaid partners; if you don't see the partnership you're looking for, contact your Account Manager or the [sales team](https://plaid.com/contact/).

#### Security portal, audits, and certifications

For a high-level overview of information on Plaid's privacy and security practices for developers, see [Plaid Safety](https://plaid.com/safety/).

For details on Plaid's current certifications, audit results, and the security knowledge base, see the [Security Portal](https://security.plaid.com/).

Plaid maintains a [HackerOne bug bounty program](https://hackerone.com/plaid) for security researchers.

#### Bank Coverage Explorer

For a detailed chart showing Plaid's coverage by bank and product, see the [US and Canada Bank Coverage Explorer](/docs/institutions/) and the [European Bank Coverage Explorer](/docs/institutions/europe/).

#### Third-party resources

Several third party developers have built resources to help you integrate with
Plaid. Note that these resources are not built or supported by Plaid, and
inclusion of a resource on the list below does not constitute an endorsement by
Plaid or indicate that Plaid has evaluated the resource for correctness,
security, or anything else.

If you have a third-party resource you'd like to be featured on this list, DM us
[@plaiddev](https://twitter.com/plaiddev) on X, or reach out to an admin
on the [Plaid Developer Discord](https://discord.gg/sf57M8DW3y).

- [How to Build a Fintech app on AWS using the Plaid API](https://aws.amazon.com/blogs/apn/how-to-build-a-fintech-app-on-aws-using-the-plaid-api/)
  and the accompanying
  [sample app](https://github.com/aws-samples/aws-plaid-demo-app) demonstrates
  how to build a Plaid app on AWS, using the latest Transactions Sync API. The
  app and tutorial incorporate AWS Lambda for webhook handling and Amazon
  Cognito for authentication.
- The [Plaid tag on GitHub](https://github.com/topics/plaid) contains multiple example apps using Plaid.
- [CodeSandbox](https://codesandbox.io/examples/package/plaid) has several
  sample apps that use Plaid.
- [WebhookDB](https://webhookdb.com/docs/plaid) has a tutorial showing how to
  use their service to manage Plaid webhooks.
- The [CloudQuery Plaid plugin](https://github.com/cloudquery/cq-source-plaid)
  extracts data from Plaid and loads it into any supported CloudQuery
  destination, such as PostgreSQL, Snowflake, BigQuery, or S3.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
