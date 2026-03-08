set shell := ["bash", "-euo", "pipefail", "-c"]

# Build the plaid binary into the repo root.
build:
	go build -o plaid .

# Install the CLI binary to ~/.local/bin.
install: build
	mkdir -p "$HOME/.local/bin"
	install -m 0755 plaid "$HOME/.local/bin/plaid"
