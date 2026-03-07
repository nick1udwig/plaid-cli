# Plaid Docs Snapshot

This directory mirrors Plaid's documentation site as local Markdown for reference while building the CLI.

- Source site: `https://plaid.com/docs/`
- Imported from Plaid sitemap plus `/docs/`
- Snapshot format: one `index.md` per docs path
- Manifest: `docs/plaid/manifest.json`
- Failures: `docs/plaid/failures.json`

Refresh the snapshot with:

```bash
uv run --with beautifulsoup4 --with markdownify python scripts/scrape_plaid_docs.py
```
