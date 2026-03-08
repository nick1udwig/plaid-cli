---
title: "Resources - MCP Server | Plaid Docs"
source_url: "https://plaid.com/docs/resources/mcp/"
scraped_at: "2026-03-07T22:05:16+00:00"
---

# Dashboard MCP Server

#### Connect your LLM-powered application to Plaid’s developer tools

For AI tools and applications, Plaid hosts a [Model Context Protocol](https://modelcontextprotocol.io/) (MCP) server that provides several tools to help you better understand your Plaid integration health, investigate user-facing issues, and more.

This article is intended for developers looking to build their own applications that can interact with the MCP server.

To use the MCP server with Claude Desktop, Cursor, VS Code, or Zed, see the [AI toolkit GitHub repo documentation](https://github.com/plaid/ai-coding-toolkit/tree/main/sandbox) instead.

For use with Claude.ai, documentation can be found on [Anthropic's website](https://support.anthropic.com/en/articles/10168395-setting-up-integrations-on-claude-ai).

The Dashboard MCP server is a new feature under active development and breaking changes may occur. Plaid currently offers limited official support.

#### Integration Process

##### Requirements

The Dashboard MCP server only works with production data. You must be approved for production access with at least one Plaid product in order to use the MCP server.

##### Authorization

In order to connect to the Dashboard MCP server, you must first create an OAuth token with the scope `mcp:dashboard` and the `client_credentials` grant type via the [`/oauth/token`](/docs/api/oauth/#oauthtoken) API.

Creating an Access Token

```
curl -X POST https://production.plaid.com/oauth/token -H 'Content-Type: application/json' -d '{ "client_id": "${PLAID_CLIENT_ID}", "client_secret": "YOUR_PRODUCTION_SECRET", "grant_type": "client_credentials", "scope": "mcp:dashboard" }'
```

This will return an `access_token` you can use to authenticate into the Dashboard MCP server. This will also return a `refresh_token` that you can use to request a new access token when the old one expires.

##### Creating the MCP connection

Plaid's Dashboard MCP server is available at <https://api.dashboard.plaid.com/mcp>, using the Streamable HTTP protocol. When you communicate with the MCP server, you must pass the `access_token` from the authorization step above via an `Authorization: Bearer <access_token>` header.

The exact method for connecting with the MCP server depends on the LLM client library you are using, but it is typically provided to the model as an entry in the list of tools that it can access. For more information, please refer to the documentation from your model provider, such as [OpenAI](https://platform.openai.com/docs/guides/tools-remote-mcp) or [Anthropic](https://docs.anthropic.com/en/docs/agents-and-tools/mcp-connector). Plaid recommends using the most recent version of your model's client library.

The following is a basic example of how to access the Dashboard MCP server using OpenAI or Anthropic's client libraries for Python.

Select group for content switcher

OpenAI Python Example

```
import json
import openai

def request_or_refresh_access_token():
    # For simplicity, we're just creating a new access token every time
    plaid_client = get_plaid_client()
    oauth_token_request = {
        "grant_type": "client_credentials",
        "scope": "mcp:dashboard",
    }

    response = plaid_client.oauth_token(oauth_token_request)
    return response["access_token"]

def main() -> None:
    """Run the example call and print the structured response."""

    client = openai.OpenAI()
    dashboard_token = request_or_refresh_access_token()

    response = client.responses.create(
        model="gpt-5.2",
        tools=[
            {
                "type": "mcp",
                "server_label": "plaid",
                "server_url": "https://api.dashboard.plaid.com/mcp",
                "require_approval": "never",
                "headers": {"Authorization": "Bearer " + dashboard_token},
            }
        ],
        # input="What is the status of Item ID <item ID>?",
        input="Please examine all Link sessions over the last week and provide me with a detailed report. Call out any patterns you see among Link sessions that were not successful.",
    )

    # Nearly all calls to the openAI client will return the mcp_list_tools object
    # and an assistant message. If the client made a call to the MCP server,
    # it will also return one or more mcp_call objects.
    for item in response.output:
        if item.type == "mcp_list_tools":
            print("\n=== MCP Tools ===")
            for tool in item.tools:
                print(f" • {tool['name']}")

        elif item.type == "mcp_call":
            print("\n=== MCP Call ===")
            print(f"Tool: {item.name}")
            try:
                # The arguments field is a JSON string – parse it for readability.
                args = json.loads(item.arguments)
            except (TypeError, json.JSONDecodeError):
                args = item.arguments
            print("Arguments:")
            print(json.dumps(args, indent=2))

        elif item.type == "message":
            print("\n=== Assistant Message ===")
            # The content field is a list of ResponseOutputText objects.
            texts = []
            for part in item.content:
                # Depending on SDK version, `text` may be an attribute or dict key.
                text = getattr(part, "text", None) or (
                    part.get("text") if isinstance(part, dict) else None
                )
                if text:
                    texts.append(text)
            print("\n".join(texts))

        else:
            print(f"\n=== Unhandled output type: {item.type} ===")
            print(item)

if __name__ == "__main__":
    main()
```

Sample output

```
=== MCP Tools ===
 • plaid_debug_item
 • plaid_get_link_analytics
 • plaid_get_tools_introduction
 • plaid_get_usages
 • plaid_list_teams

=== MCP Call ===
Tool: plaid_list_teams
Arguments:
{}

=== MCP Call ===
Tool: plaid_get_link_analytics
Arguments:
{
  "from_date": "2025-05-01",
  "team_id": "58b9a231bdc6a453f37c338e",
  "to_date": "2025-05-08"
}

=== Assistant Message ===
Here is a detailed analysis of all Plaid Link sessions between 2025-05-01 and 2025-05-08, including successful conversions and error patterns:

## 1. Conversion Funnel Summary

- Link Opens: 58
- Institution Selected: 29
- Handoffs (successfully completed): 18
...
<Message truncated for brevity>
```

##### Expired tokens

Access tokens for the MCP server expire after 15 minutes. If the `access_token` passed to the server has expired, requests will return a 401 status code. You can request a new token either by creating a new one as described above, or by calling the [`/oauth/token`](/docs/api/oauth/#oauthtoken) endpoint with the `refresh_token` from the previous step. Both options are valid.

Refreshing an Access Token

```
curl -X POST https://production.plaid.com/oauth/token -H 'Content-Type: application/json' -d '{ "client_id": "${PLAID_CLIENT_ID}", "secret": "YOUR_PRODUCTION_SECRET", "refresh_token": "YOUR_REFRESH_TOKEN", "grant_type": "refresh_token" }'
```

Please note that if your `access_token` has expired, different client libraries will surface that information in different ways. In OpenAI's Python library, for instance, this is raised as a `openai.APIStatusError`:

Error object for expired token

```
Error code: 424 - {'error': {'message': "Error retrieving tool list from MCP server: 'plaid'. Http status code: 401 (Unauthorized)", 'type': 'external_connector_error', 'param': 'tools', 'code': 'http_error'}}
```

In Anthropic's Python library, this is raised as an `anthropic.BadRequestError:`

Error object for expired token

```
Error code: 400 - {'type': 'error', 'error': {'type': 'invalid_request_error', 'message': "Invalid authorization token for MCP server 'Plaid Dashboard Server'. Please check your token and try again."}}
```

#### Dashboard MCP tools

The Dashboard MCP server currently supports the following tools:

##### `plaid_debug_item`

Diagnose a Plaid item by retrieving related metadata. This tool provides comprehensive information about why an item may not be working properly.

##### `plaid_get_link_analytics`

Retrieves Plaid Link analytics data for analyzing user conversion and error rates. This tool can be used to:

- Analyze Link conversion funnel metrics
- Track user progression through the Link flow
- Monitor error rates over time
- Evaluate Link performance and user experience

Some common use cases for using the `plaid_get_link_analytics` tool include:

- Monitoring Link conversion rates and identifying drop-offs
- Analyzing conversion trends over specific time periods
- Tracking error patterns and frequency (both by type and total via CountEntities)
- Generating comprehensive Link performance reports
- Comparing performance across different stages of the Link flow

##### `plaid_get_usages`

Retrieves usage metrics for Plaid products and services. Use this tool when you need to:

- Get usage data for specific metrics over a time period
- Track product usage and consumption
- Monitor API request volumes

##### `plaid_list_teams`

List all teams associated with the OAuth token. This is often called automatically by the MCP server to ensure that it is retrieving the data from the appropriate team.

For more details about the tools provided by the Dashboard MCP server, connect to the server using a tool such as the [MCP Inspector](https://github.com/modelcontextprotocol/inspector).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
