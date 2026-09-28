# Maintaining the CLI

This repository owns the `cadenya` command tree. Commands, names, flags, help text, examples, and output belong here. Redwood is not part of this repository's build or release path.

## Applying an API change

1. Start from the public OpenAPI/proto diff in `cadenya/cadenya`. Identify additions, changed request or response fields, and behavior that matters in a terminal.
2. Update the relevant command and its help in this repository. Prefer a task-oriented command over exposing every request field verbatim. Preserve existing flags and JSON output unless a deliberate CLI compatibility change is being released.
3. Update `api.md` and focused examples when a command changes. The initial command inventory was imported from the last Redwood staging output; it is now maintained here.
4. Keep the Go SDK version in `go.mod` pinned. A CLI-only change needs no SDK update. If an API change requires new SDK code, use an available SDK tag or a Go pseudo-version for a pushed SDK commit. A CLI release does not require an SDK version tag. The SDK module must still be reachable by Go tooling.
5. Update `cadenya config` in `internal/config`, `internal/reconcile`, `internal/configcommand`, and `internal/configaction` when bundle behavior changes. Keep the report schema and GitHub Action outputs compatible.

## Releasing

Merge CLI changes directly into this repository. Release Please opens a version and changelog PR; merging it creates a tag. The tagged GoReleaser workflow builds and publishes `cadenya` archives, Linux packages, the Homebrew cask, and Chocolatey when configured. Check the archive checksum and the changed command in a downloaded binary.

After the first CLI release with `cadenya config`, update `cadenya/cadenya-config-loader` examples to the actual CLI tag and decide whether its Action should provide that tag as a default. The Action has no Go package and never publishes a separate executable. Existing Action tags keep using their original binary.
