#!/usr/bin/env python3

from __future__ import annotations

import json
import re
import sys
import time
from concurrent.futures import FIRST_COMPLETED, ThreadPoolExecutor, wait
from dataclasses import dataclass
from datetime import datetime, timezone
from pathlib import Path
from urllib.error import HTTPError, URLError
from urllib.parse import urljoin, urlparse
from urllib.request import Request, urlopen

from bs4 import BeautifulSoup, NavigableString, Tag
from markdownify import markdownify as md


BASE_URL = "https://plaid.com"
SITEMAP_URL = f"{BASE_URL}/sitemap.xml"
DOCS_PREFIX = "/docs/"
OUTPUT_ROOT = Path("docs/plaid")
USER_AGENT = (
    "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 "
    "(KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"
)
HTML_TIMEOUT_SECONDS = 30
MAX_WORKERS = 2
REQUEST_PAUSE_SECONDS = 0.25
RETRYABLE_STATUS_CODES = {429, 500, 502, 503, 504}
MAX_RETRIES = 5

SITEMAP_LINK_RE = re.compile(r"""<loc>https://plaid\.com(/docs/[^<]+)</loc>""")
MULTIBLANK_RE = re.compile(r"\n{3,}")


@dataclass
class PageResult:
    canonical_path: str
    source_url: str
    output_file: str


def canonicalize_path(value: str) -> str | None:
    parsed = urlparse(value)
    path = parsed.path or value
    path = path.split("#", 1)[0].split("?", 1)[0]
    if path.endswith("/index.html.md"):
        path = path[: -len("index.html.md")]
    if path == "/docs":
        path = "/docs/"
    if not path.startswith(DOCS_PREFIX):
        return None
    if path.endswith(".md") or path.endswith(".html"):
        return None
    return path


def build_request(url: str) -> Request:
    return Request(url, headers={"User-Agent": USER_AGENT})


def fetch_text(url: str, timeout_seconds: int) -> str:
    for attempt in range(MAX_RETRIES):
        try:
            with urlopen(build_request(url), timeout=timeout_seconds) as response:
                charset = response.headers.get_content_charset() or "utf-8"
                text = response.read().decode(charset, "replace")
            time.sleep(REQUEST_PAUSE_SECONDS)
            return text
        except HTTPError as exc:
            if exc.code not in RETRYABLE_STATUS_CODES or attempt == MAX_RETRIES - 1:
                raise
            retry_after = exc.headers.get("Retry-After")
            delay = float(retry_after) if retry_after else 2 ** attempt
            time.sleep(delay)
        except URLError:
            if attempt == MAX_RETRIES - 1:
                raise
            time.sleep(2 ** attempt)
    raise RuntimeError(f"Failed to fetch {url}")


def load_sitemap_paths() -> list[str]:
    sitemap = fetch_text(SITEMAP_URL, HTML_TIMEOUT_SECONDS)
    paths = ["/docs/"]
    for raw in SITEMAP_LINK_RE.findall(sitemap):
        canonical = canonicalize_path(raw)
        if canonical:
            paths.append(canonical)
    return sorted(set(paths))


def output_path_for(canonical_path: str) -> Path:
    relative = canonical_path[len(DOCS_PREFIX) :].strip("/")
    if not relative:
        return OUTPUT_ROOT / "index.md"
    return OUTPUT_ROOT / relative / "index.md"


def unwrap_heading_links(container: Tag) -> None:
    for anchor in list(container.find_all("a")):
        children = [
            child
            for child in anchor.children
            if not isinstance(child, NavigableString) or child.strip()
        ]
        if len(children) != 1:
            continue
        only_child = children[0]
        if isinstance(only_child, Tag) and only_child.name in {
            "h1",
            "h2",
            "h3",
            "h4",
            "h5",
            "h6",
        }:
            anchor.replace_with(only_child)


def clean_main(main: Tag) -> Tag:
    for tag in main.find_all(
        [
            "script",
            "style",
            "noscript",
            "svg",
            "button",
            "input",
            "textarea",
            "select",
            "form",
        ]
    ):
        tag.decompose()

    for hidden in main.select("[aria-hidden='true']"):
        hidden.decompose()

    unwrap_heading_links(main)
    return main


def html_to_markdown(html: str, source_url: str) -> str:
    soup = BeautifulSoup(html, "html.parser")
    title = soup.title.get_text(" ", strip=True) if soup.title else source_url
    main = soup.find("main")
    if main is None:
        raise RuntimeError(f"Missing <main> for {source_url}")

    main = clean_main(main)
    content = md(
        str(main),
        heading_style="ATX",
        bullets="-",
        strip=["span"],
    ).strip()
    content = MULTIBLANK_RE.sub("\n\n", content)

    scraped_at = datetime.now(timezone.utc).replace(microsecond=0).isoformat()
    return (
        "---\n"
        f"title: {json.dumps(title)}\n"
        f"source_url: {json.dumps(source_url)}\n"
        f"scraped_at: {json.dumps(scraped_at)}\n"
        "---\n\n"
        f"{content}\n"
    )


def ensure_output_root() -> None:
    OUTPUT_ROOT.mkdir(parents=True, exist_ok=True)


def fetch_and_write(path: str) -> PageResult:
    canonical_path = canonicalize_path(path)
    if canonical_path is None:
        raise RuntimeError(f"Invalid docs path: {path}")

    source_url = urljoin(BASE_URL, canonical_path)
    html = fetch_text(source_url, HTML_TIMEOUT_SECONDS)
    markdown = html_to_markdown(html, source_url)

    destination = output_path_for(canonical_path)
    destination.parent.mkdir(parents=True, exist_ok=True)
    destination.write_text(markdown, encoding="utf-8")

    return PageResult(
        canonical_path=canonical_path,
        source_url=source_url,
        output_file=str(destination),
    )


def crawl(paths: list[str]) -> tuple[list[dict], list[dict]]:
    ensure_output_root()

    manifest: list[dict] = []
    failures: list[dict] = []
    in_flight = {}
    pending = list(reversed(paths))
    completed = 0

    with ThreadPoolExecutor(max_workers=MAX_WORKERS) as executor:
        while pending or in_flight:
            while pending and len(in_flight) < MAX_WORKERS:
                path = pending.pop()
                future = executor.submit(fetch_and_write, path)
                in_flight[future] = path

            done, _ = wait(in_flight, return_when=FIRST_COMPLETED)
            for future in done:
                path = in_flight.pop(future)
                try:
                    result = future.result()
                except HTTPError as exc:
                    failures.append(
                        {
                            "requested_path": path,
                            "error": f"HTTP {exc.code}",
                            "reason": exc.reason,
                        }
                    )
                    continue
                except URLError as exc:
                    failures.append(
                        {
                            "requested_path": path,
                            "error": "URL error",
                            "reason": str(exc.reason),
                        }
                    )
                    continue
                except Exception as exc:  # pragma: no cover - defensive
                    failures.append(
                        {
                            "requested_path": path,
                            "error": type(exc).__name__,
                            "reason": str(exc),
                        }
                    )
                    continue

                completed += 1
                manifest.append(
                    {
                        "path": result.canonical_path,
                        "source_url": result.source_url,
                        "output_file": result.output_file,
                    }
                )
                if completed % 25 == 0:
                    print(f"saved {completed} pages", file=sys.stderr)

    manifest.sort(key=lambda item: item["path"])
    failures.sort(key=lambda item: item["requested_path"])
    return manifest, failures


def main() -> int:
    paths = load_sitemap_paths()
    manifest, failures = crawl(paths)

    (OUTPUT_ROOT / "manifest.json").write_text(
        json.dumps(manifest, indent=2) + "\n",
        encoding="utf-8",
    )
    (OUTPUT_ROOT / "failures.json").write_text(
        json.dumps(failures, indent=2) + "\n",
        encoding="utf-8",
    )

    print(f"saved {len(manifest)} canonical docs pages")
    if failures:
        print(f"encountered {len(failures)} fetch failures", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
