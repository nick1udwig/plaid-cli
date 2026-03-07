# Plaid Go CLI Plan

## Goal

Build an agent-friendly Go CLI over Plaid's API surface so a user can give an agent controlled access to their Plaid-backed financial data and operations. The agent is the app: this is an "app for one", not a reusable end-user product integration.

## Decisions

- Scope: Plaid API only. Do not target the Plaid Dashboard MCP server in this CLI.
- Delivery plan: keep the phase breakdown in this document and execute in that order.
- Command model: keep product/resource subcommands and stable JSON output.
- Capability handling: document required capabilities clearly in command help text, but do not enforce policy in the CLI itself.

## Source Snapshot

- Local docs mirror: `docs/plaid/`
- Canonical docs pages imported: `284`
- API surface sourced from:
  - `docs/plaid/api/`
  - `docs/plaid/api/products/`
  - related product guides under `docs/plaid/*`

## Key Docs Findings

- Plaid's API is JSON over HTTPS and almost every endpoint is a `POST`.
- Item + Link lifecycle endpoints are prerequisite plumbing for most product flows.
- Balance is documented as a product, but its current API reference lives under Signal:
  - `balance` product UX/docs: `docs/plaid/balance/index.md`
  - realtime balance endpoint: `/accounts/balance/get`
  - signal endpoint family: `docs/plaid/api/products/signal/index.md`
- `protect` currently has a product API page in the docs tree, but no documented public endpoints.
- `transfer` is by far the largest command family and should stay grouped under one top-level command with nested action groups.

## Recommended Top-Level Commands

### Core lifecycle and platform commands

- `account`
  - `/accounts/get`
- `item`
  - `get`, `remove`, `webhook-update`, `public-token-exchange`, `access-token-invalidate`
- `link`
  - `token-create`, `token-get`
- `institution`
  - `get`, `get-by-id`, `search`
- `processor`
  - `token-create`, `stripe-bank-account-token-create`, `token-permissions-get`, `token-permissions-set`
- `user`
  - `create`, `get`, `update`, `remove`, `items-get`
- `oauth`
  - `token`, `introspect`, `revoke`
- `consent`
  - `events-get`
- `network`
  - `status-get`
- `sandbox`
  - all `/sandbox/*` endpoints, including transfer simulators and test clocks

### Product commands

- `auth`
  - `get`, `verify`, `bank-transfer-event-list`, `bank-transfer-event-sync`
- `balance`
  - `get` -> `/accounts/balance/get`
- `signal`
  - `evaluate`, `prepare`, `decision-report`, `return-report`
- `identity`
  - `get`, `match`, `documents-uploads-get`
- `transactions`
  - `sync`, `get`, `recurring-get`, `refresh`, `categories-get`
- `investments`
  - `holdings-get`, `transactions-get`, `refresh`
- `investments-move`
  - `auth-get`
- `liabilities`
  - `get`
- `assets`
  - `report-create`, `report-get`, `report-pdf-get`, `report-refresh`, `report-filter`, `report-remove`, `audit-copy-create`, `audit-copy-remove`
- `income`
  - `sessions-get`, `bank-income-get`, `bank-income-pdf-get`, `bank-income-refresh`, `bank-statements-uploads-get`, `payroll-income-get`, `payroll-income-risk-signals-get`, `payroll-income-refresh`, `employment-get`, `parsing-config-update`
- `statements`
  - `list`, `download`, `refresh`
- `enrich`
  - `transactions-enrich`
- `check`
  - CRA report and monitoring endpoints
- `identity-verification`
  - `create`, `get`, `list`, `retry`
- `monitor`
  - nested `individual` and `entity` groups for screenings, reviews, hits, programs
- `beacon`
  - nested `user`, `report`, `report-syndication`, `duplicate`, `dashboard-user`
- `layer`
  - `session-token-create`, `user-account-session-get`
- `payment-initiation`
  - nested `recipient`, `payment`, and `consent` groups
- `transfer`
  - nested groups:
  - `authorization`
  - `transfer`
  - `event`
  - `sweep`
  - `refund`
  - `recurring`
  - `ledger`
  - `capabilities`
  - `intent`
  - `originator`
- `virtual-account`
  - maps to `/wallet/*`

## Capability Model

Each command should declare a required capability set in its help text and generated docs. Enforcement happens outside the CLI.

- `read`
  - `get`, `list`, `search`, `sync`, `download`, `status`, `introspect`
- `write`
  - `create`, `update`, `remove`, `cancel`, `revoke`, `invalidate`, `execute`, `submit`, `retry`, `prepare`, `verify`, `exchange`, `refresh`, `report`
- `sandbox`
  - any `/sandbox/*` endpoint
- `admin`
  - `oauth`, processor token permissions, webhook URL updates, user removal

Examples:

- `auth get` -> `read`
- `balance get` -> `read`
- `item remove` -> `write`
- `link token-create` -> `write`
- `transfer create` -> `write`
- `sandbox transfer-simulate` -> `sandbox`
- `oauth token` -> `admin`

## CLI UX

- JSON to stdout, diagnostics to stderr
- non-zero exit on Plaid API errors
- include `request_id`, `error_type`, `error_code`, and HTTP status in error output
- persistent state lives under `~/.plaid-cli`
  - config
  - credentials
  - cached metadata
  - generated request templates or other local state as needed
- shared flags:
  - `--env sandbox|development|production`
  - `--client-id`
  - `--secret`
  - `--access-token`
  - `--output json|json-pretty`
- every command should expose typed flags for:
  - required request fields
  - the most common optional fields
- every command should also support:
  - `--body @file.json` as an escape hatch for complex or rare nested payloads
  - `--print-request-template` to emit a minimal request skeleton for that endpoint
  - `--print-doc-path` to show the local docs page backing the command
- merge behavior:
  - `--body` provides the base JSON request object
  - explicit flags override fields from `--body`

## Input Strategy

The CLI should not force the agent to memorize every Plaid request schema.

The right model is a hybrid:

- first-class flags for the fields an agent will most often need
- endpoint-aware help text listing required fields and common optional fields
- `--print-request-template` for bootstrapping valid JSON quickly
- `--body` only for long-tail fields or deeply nested structures

Example shape:

- `plaid auth get --access-token ...`
- `plaid balance get --access-token ... --account-id ...`
- `plaid transfer create --body @transfer.json`
- `plaid payment-initiation consent create --print-request-template`

So no, the intended design is not "the agent must know every parameter up front". JSON remains the escape hatch, not the primary ergonomics layer.

## Persistent State

Use `~/.plaid-cli` for all persistent local state.

Proposed layout:

- `~/.plaid-cli/config.json`
- `~/.plaid-cli/app-profile.json`
- `~/.plaid-cli/items/`
- `~/.plaid-cli/cache/`
- `~/.plaid-cli/logs/` if local debug logs are added later

At minimum, writes should create files with restrictive permissions.
Encryption at rest is desirable post-MVP, but not required for the first implementation pass.

### State Model

Do not conflate Plaid app credentials with end-user bank connections.

- app profile
  - Plaid `client_id`
  - Plaid `secret`
  - environment
  - `client_name`
  - default country codes
  - default product sets / Link options
- item
  - `item_id`
  - `access_token`
  - institution metadata
  - selected account metadata
  - product/consent metadata as needed

Important:

- we do **not** store bank login credentials
- we store Plaid tokens and metadata resulting from Link
- multiple Items is the normal case, because one human may link several institutions

Recommended v1 default:

- single local owner only
- one persisted app profile for now
- multiple linked Items under that single local owner

## Bootstrap and Login UX

Plaid app credentials are provisioned out-of-band via the Plaid Dashboard, not created by this CLI.

Recommended bootstrap flow:

1. `plaid init --env production --client-id ... --secret ...`
2. CLI stores app profile in `~/.plaid-cli/app-profile.json`
3. `plaid link connect ...`
4. CLI creates a Link token with Hosted Link enabled
5. CLI opens the Hosted Link URL in the user's browser
6. CLI polls `/link/token/get` until completion, then exchanges `public_token` for `access_token`
7. CLI saves the resulting Item under `~/.plaid-cli/items/...`

Why Hosted Link:

- it is the simplest browser-based flow for a CLI
- Plaid hosts the frontend
- `/link/token/get` can be used to retrieve `public_token` and session results without requiring a frontend callback

For MVP, do not require a local callback server. Polling `/link/token/get` is sufficient and is a better fit for a non-interactive CLI.

## Implementation Recommendation

Use raw HTTP over Plaid's REST API for v1, with a hybrid typed-flags-plus-JSON request builder.

Why:

- Plaid's API is consistently JSON-over-POST.
- The command surface is large and keeps growing.
- agents benefit from stable JSON output and partially typed input, without us hand-modeling every nested request schema on day one.
- we can still keep the docs and OpenAPI as the source of truth for coverage.

The official Go client (`plaid-go`) is still useful as a reference, but using it as the primary command implementation will make broad endpoint coverage slower.

## Suggested Phases

### Phase 1: CLI skeleton and core read flows

- Go module and command framework
- config/env loading
- shared HTTP client
- shared request/response renderer
- capability metadata/helptext
- commands:
  - `link`
  - `item`
  - `account`
  - `institution`
  - `auth`
  - `balance`
  - `identity`
  - `transactions`
  - `sandbox`

### Phase 2: money movement and risk

- `signal`
- `transfer`
- `processor`
- `payment-initiation`
- `oauth`
- `user`

### Phase 3: long-tail products

- `assets`
- `income`
- `statements`
- `check`
- `identity-verification`
- `monitor`
- `beacon`
- `layer`
- `investments`
- `investments-move`
- `virtual-account`

### Phase 4: spec-driven generation

- generate endpoint metadata from Plaid docs or OpenAPI
- reduce hand-maintained command definitions
- add coverage checks against the local docs snapshot

## Testing Strategy

- unit tests for request building, flag parsing, help metadata, and error formatting
- golden tests for JSON stdout/stderr
- sandbox integration tests gated by env vars
- docs/spec drift test against `docs/plaid/manifest.json`
