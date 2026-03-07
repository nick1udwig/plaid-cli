# Getting Started

This CLI is intended for a single human operator who wants to let an agent work with the Plaid API on their behalf.

## What You Need

Before the CLI can do anything, the human operator must obtain the Plaid app credentials from the Plaid Dashboard:

- `client_id`
- `secret`
- target environment: `sandbox`, `development`, or `production`
- a `client_name` to display in Link

Plaid provisions these values through the Dashboard, not through a CLI bootstrap API.

## One-Time Setup

Initialize the CLI with your Plaid app credentials:

```bash
plaid init \
  --env production \
  --client-id YOUR_CLIENT_ID \
  --secret YOUR_SECRET \
  --client-name "Plaid CLI"
```

This stores persistent state under `~/.plaid-cli`.

## Connect a Bank with Hosted Link

Open a Hosted Link session in the browser and save the resulting Item locally:

```bash
plaid link connect \
  --product auth \
  --product transactions
```

The CLI will:

1. Create a Link token using your saved app profile
2. Open the Hosted Link URL in your browser
3. Poll Plaid for session completion
4. Exchange the returned `public_token` for an `access_token`
5. Save the linked Item under `~/.plaid-cli/items/`

## Notes

- The CLI stores Plaid tokens and metadata, not your bank username/password.
- Files written under `~/.plaid-cli` should use restrictive permissions.
- Encrypting sensitive state at rest is desirable post-MVP, but MVP uses filesystem permissions first.
- For MVP, the CLI is intentionally non-interactive apart from opening the browser for Hosted Link.
