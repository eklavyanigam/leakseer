# Leakseer

A command-line tool, written in Go, that scans files and folders for leaked secrets such as API keys and tokens. Its focus is AI tooling: besides generic secret formats, it checks `.env` files and MCP/AI tool configs, where keys often end up by accident.

> **Status: in progress.** This is a learning project, built one stage at a time. Not ready for real use yet.

## Design rules
- Standard library only, no third-party dependencies.
- Findings never contain the full secret, only a masked preview.
- No real secrets in the repo. All test values are fake.

## Progress
- [x] Project setup
- [ ] Scan a single file (in progress: line parsing, masking and the first detector are being built)
- [ ] Recursive folder scanning
- [ ] Pluggable detectors
- [ ] Concurrent scanning
- [ ] `.env` and MCP config rules
- [ ] False-positive handling (inline ignore and ignore file)
- [ ] CLI flags, JSON output and exit codes
- [ ] Docs, tests and CI

## Usage
Not available yet.