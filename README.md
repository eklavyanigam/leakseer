# Leakseer

A command-line tool, written in Go, that scans files for leaked secrets such as API keys, tokens and private keys. The goal is to go further than generic scanners by adding rules for AI tooling: MCP/AI tool configs, `.env` files and Jupyter notebooks.

> **Status: in progress.** This is a learning project, built one stage at a time. Not ready for real use yet.

## Design rules
- Standard library only, no third-party dependencies.
- Findings never contain the full secret, only a masked preview.
- No real secrets in the repo. All test values are fake.

## Progress
- [x] Project setup
- [ ] Scan a single file (in progress: line parsing and masking done, detection next)
- [ ] Recursive folder scanning
- [ ] Pluggable detectors
- [ ] Concurrent scanning
- [ ] AI/MCP config, `.env` and notebook rules
- [ ] False-positive handling
- [ ] CLI flags, JSON output and exit codes
- [ ] Release: docs, CI, binaries

## Usage
Not available yet.