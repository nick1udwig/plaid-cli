---
name: explore-plaid-cli
description: Provides concrete `plaid` CLI workflows, one-line commands, and sandbox-first playbooks for Plaid API tasks. Use when a user asks how to try `plaid-cli` by hand, wants runnable examples, wants likely agent use cases, or needs help mapping goals like reading balances, syncing transactions, testing webhooks, or moving money to the right commands.
---

# Explore Plaid CLI

Provide short, runnable playbooks for this repo's `plaid` binary. Prefer sandbox flows unless the user explicitly asks for development or production.

## Quick Start

- Read [references/use-cases.md](references/use-cases.md) for concrete command sequences.
- Point humans to `docs/getting-started.md` for one-time app setup and `docs/live-test-setup.md` for sandbox setup.
- Use `plaid` if it is on `PATH`; otherwise replace `plaid` with `go run .`.
- State when a snippet requires `jq` to capture IDs from JSON output.
- Prefer one-line commands that can be copied directly into a shell.
- Prefer `--state-dir /tmp/plaid-play` or another scratch directory for demos and ad hoc exploration.
- If `.env.sandbox` exists, suggest `set -a; source .env.sandbox; set +a` before running commands.

## Workflow

1. Identify whether the user wants a read-only flow, sandbox mutation flow, browser-based Link flow, webhook test, or money-movement flow.
2. Default to sandbox.
3. If the CLI has not been initialized, start with `plaid init`.
4. For the fastest headless demo, prefer `sandbox public-token-create` plus `item public-token-exchange` over Hosted Link.
5. Reuse saved `item_id` and `account_id` values to build follow-up commands.

## Response Style

- Give concrete commands first, not broad descriptions.
- Separate safe read-only commands from write or move-money commands.
- Call out product or dashboard prerequisites before suggesting a gated flow.
- Prefer runnable sequences over `--help`-only guidance when the command is safe to demonstrate.
- Use `plaid <command> --help` or the mirrored docs under `docs/plaid/` when exact flags matter.
- Do not ask the human to expose secrets in chat; use env vars or `plaid init`.

## Key Paths

- `README.md`
- `docs/getting-started.md`
- `docs/live-test-setup.md`
- `docs/plaid/`
- [references/use-cases.md](references/use-cases.md)
