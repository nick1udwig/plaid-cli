# plaid-cli

Single-owner, agent-friendly CLI for the Plaid API.

## Setup

Humans must complete one-time setup before an agent can use the CLI.

See [docs/getting-started.md](docs/getting-started.md) for:

- how to get the required Plaid app credentials
- how to run `plaid init`
- how to connect a bank with Hosted Link

## Examples

Print a minimal request template for a transfer authorization:

```bash
plaid transfer authorization create --print-request-template
```

Check whether a linked account supports Transfer capabilities:

```bash
plaid transfer capabilities get --item YOUR_ITEM_ID --account-id YOUR_ACCOUNT_ID
```

Create a transfer using typed flags plus JSON as an escape hatch for less common fields:

```bash
plaid transfer create \
  --authorization-id YOUR_AUTHORIZATION_ID \
  --description "payment" \
  --body @transfer.json
```

## Local Docs Snapshot

Plaid's docs are mirrored under [docs/plaid/](docs/plaid/) for implementation reference.
