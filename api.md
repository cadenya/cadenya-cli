# cadenya CLI reference

Request bodies are built from typed flags; any subtree also takes a YAML/JSON document (`@file`, `-` for stdin, or a literal; at most ONE input per invocation may read stdin) and `-f <doc>` supplies the whole body. Document flags are listed under `schema <command>`. See README.md for usage patterns.

## cadenya accounts

Retrieves the current account for the token accessing the API

```sh
cadenya accounts retrieve
```
Rotates the challenge token for the account

```sh
cadenya accounts rotate-challenge-token
```
Rotates the webhook signing key for the account

```sh
cadenya accounts rotate-webhook-signing-key
```

## cadenya api-keys

Get the global API key

```sh
cadenya api-keys retrieve-global
```
Disable the global API key

```sh
cadenya api-keys disable-global
```
Enable the global API key

```sh
cadenya api-keys enable-global
```
Rotate the global API key

```sh
cadenya api-keys rotate-global
```
List API keys

```sh
cadenya api-keys list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new API key

```sh
cadenya api-keys create [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... [--description <value>] [--permission <value>]... [-f <doc>] [--dry-run]
```
Get an API key by ID

```sh
cadenya api-keys retrieve <id> [--workspace-id <value>]
```
Delete an API key

```sh
cadenya api-keys delete <id> [--workspace-id <value>]
```
Update an API key

```sh
cadenya api-keys update <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--description <value>] [--permission <value>]... [--update-mask <value>] [-f <doc>] [--dry-run]
```
Disable an API key

```sh
cadenya api-keys disable <id> [--workspace-id <value>]
```
Enable an API key

```sh
cadenya api-keys enable <id> [--workspace-id <value>]
```
Rotate an API key

```sh
cadenya api-keys rotate <id> [--workspace-id <value>]
```

## cadenya workspace-admin

Search account profiles

```sh
cadenya workspace-admin list-profiles [--limit <value>] [--cursor <value>] [--query <value>] [--labels <value>]
```
List all workspaces in the account

```sh
cadenya workspace-admin list-account [--limit <value>] [--cursor <value>] [--include-archived[=true|false]] [--labels <value>]
```
Create a workspace

```sh
cadenya workspace-admin create --name <value> [--external-id <value>] [--label KEY=VALUE]... [--description <value>] [-f <doc>] [--dry-run]
```
Get a workspace by ID

```sh
cadenya workspace-admin retrieve [--workspace-id <value>]
```
Archive a workspace

```sh
cadenya workspace-admin archive [--workspace-id <value>]
```
Update a workspace

```sh
cadenya workspace-admin update [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--description <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```
List workspace members

```sh
cadenya workspace-admin list-members [--workspace-id <value>] [--limit <value>] [--cursor <value>]
```
Add a member to a workspace

```sh
cadenya workspace-admin add-member [--workspace-id <value>] [--profile-id <value>] [--email <value>] [-f <doc>] [--dry-run]
```
Remove a member from a workspace

```sh
cadenya workspace-admin remove-member <profile-id> [--workspace-id <value>]
```

## cadenya profiles

Retrieves the profile for the credentials accessing the API

```sh
cadenya profiles whoami
```

## cadenya workspaces

List workspaces

```sh
cadenya workspaces list [--limit <value>] [--cursor <value>] [--sort-order <value>] [--include-info[=true|false]] [--labels <value>]
```

## cadenya agents

List agents

```sh
cadenya agents list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--state <value>] [--variation-selection-mode <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new agent

```sh
cadenya agents create [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... [--description <value>] [--webhook-events-url <value>] --variation-selection-mode <random|weighted> [--system-prompt-data-schema KEY=VALUE|<doc>]... [--output-definition KEY=VALUE|<doc>]... [--enable-episodic-memory[=true|false]] [--episodic-memory-ttl <value>] [--default-variation-name <value>] [--default-variation-external-id <value>] [--default-variation-label KEY=VALUE]... [--default-variation-system-prompt-template <value>] [--default-variation-discovery-max-tools <value>] [--default-variation-discovery-hint <value>]... [--default-variation-constraints-max-tool-calls <value>] [--default-variation-constraints-max-sub-objectives <value>] [--default-variation-constraints-inactivity-timeout <value>] [--default-variation-description <value>] [--default-variation-model-id <value>] [--default-variation-model-temperature <value>] [--default-variation-model-top-p <value>] [--default-variation-model-top-k <value>] [--default-variation-model-stop-sequence <value>]... [--default-variation-model-max-output-tokens <value>] [--default-variation-model-reasoning-effort <none|low|medium|high>] [--default-variation-model-caching-enabled[=true|false]] [--default-variation-compaction-trigger-threshold <value>] [--default-variation-compaction-summarization-instructions <value>] [--default-variation-compaction-tool-result-clearing-preserve-recent-results <value>] [--default-variation-first-user-message-template <value>] [-f <doc>] [--dry-run]
```
List feedback for an agent

```sh
cadenya agents list-feedback <agent-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--query <value>] [--sentiment <value>] [--agent-variation-id <value>] [--created-after <value>] [--created-before <value>] [--labels <value>] [--include-info[=true|false]]
```
List webhook deliveries

```sh
cadenya agents list-webhook-deliveries <agent-id> [--workspace-id <value>] [--cursor <value>] [--limit <value>] [--objective-id <value>] [--event-type <value>] [--labels <value>]
```
Get an agent by ID

```sh
cadenya agents retrieve <id> [--workspace-id <value>]
```
Delete an agent

```sh
cadenya agents delete <id> [--workspace-id <value>]
```
Update an agent

```sh
cadenya agents update <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--description <value>] [--webhook-events-url <value>] [--variation-selection-mode <random|weighted>] [--system-prompt-data-schema KEY=VALUE|<doc>]... [--output-definition KEY=VALUE|<doc>]... [--enable-episodic-memory[=true|false]] [--episodic-memory-ttl <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```
Archive an agent

```sh
cadenya agents archive <id> [--workspace-id <value>]
```
Publish an agent

```sh
cadenya agents publish <id> [--workspace-id <value>]
```
Unarchive an agent

```sh
cadenya agents unarchive <id> [--workspace-id <value>]
```
Unpublish an agent

```sh
cadenya agents unpublish <id> [--workspace-id <value>]
```

## cadenya agents schedules

List schedules

```sh
cadenya agents schedules list <agent-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new schedule

```sh
cadenya agents schedules create <agent-id> [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... [--schedule-calendar <doc>]... [--schedule-interval k=v,...|<doc>]... [--schedule-timezone <value>] [--overlap-policy <allow|skip>] [--first-user-message <value>] [--variation-id <value>] [--system-prompt-data KEY=VALUE|<doc>]... [--first-user-message-data KEY=VALUE|<doc>]... [-f <doc>] [--dry-run]
```
Get a schedule by ID

```sh
cadenya agents schedules retrieve <agent-id> <id> [--workspace-id <value>]
```
Delete a schedule

```sh
cadenya agents schedules delete <agent-id> <id> [--workspace-id <value>]
```
Update a schedule

```sh
cadenya agents schedules update <agent-id> <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--schedule-calendar <doc>]... [--schedule-interval k=v,...|<doc>]... [--schedule-timezone <value>] [--overlap-policy <allow|skip>] [--first-user-message <value>] [--variation-id <value>] [--system-prompt-data KEY=VALUE|<doc>]... [--first-user-message-data KEY=VALUE|<doc>]... [--update-mask <value>] [-f <doc>] [--dry-run]
```
Archive a schedule

```sh
cadenya agents schedules archive <agent-id> <id> [--workspace-id <value>]
```
Pause a schedule

```sh
cadenya agents schedules pause <agent-id> <id> [--workspace-id <value>]
```
Resume a schedule

```sh
cadenya agents schedules resume <agent-id> <id> [--workspace-id <value>]
```

## cadenya agents variations

List variations

```sh
cadenya agents variations list <agent-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--sort-order <value>] [--include-info[=true|false]] [--labels <value>]
```
Create a new variation

```sh
cadenya agents variations create <agent-id> [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... [--system-prompt-template <value>] [--discovery-max-tools <value>] [--discovery-hint <value>]... [--constraints-max-tool-calls <value>] [--constraints-max-sub-objectives <value>] [--constraints-inactivity-timeout <value>] [--description <value>] [--model-id <value>] [--model-temperature <value>] [--model-top-p <value>] [--model-top-k <value>] [--model-stop-sequence <value>]... [--model-max-output-tokens <value>] [--model-reasoning-effort <none|low|medium|high>] [--model-caching-enabled[=true|false]] [--compaction-trigger-threshold <value>] [--compaction-summarization-instructions <value>] [--compaction-tool-result-clearing-preserve-recent-results <value>] [--first-user-message-template <value>] [-f <doc>] [--dry-run]
```
Get a variation by ID

```sh
cadenya agents variations retrieve <agent-id> <id> [--workspace-id <value>]
```
Delete a variation

```sh
cadenya agents variations delete <agent-id> <id> [--workspace-id <value>]
```
Update a variation

```sh
cadenya agents variations update <agent-id> <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--system-prompt-template <value>] [--discovery-max-tools <value>] [--discovery-hint <value>]... [--constraints-max-tool-calls <value>] [--constraints-max-sub-objectives <value>] [--constraints-inactivity-timeout <value>] [--description <value>] [--model-id <value>] [--model-temperature <value>] [--model-top-p <value>] [--model-top-k <value>] [--model-stop-sequence <value>]... [--model-max-output-tokens <value>] [--model-reasoning-effort <none|low|medium|high>] [--model-caching-enabled[=true|false]] [--compaction-trigger-threshold <value>] [--compaction-summarization-instructions <value>] [--compaction-tool-result-clearing-preserve-recent-results <value>] [--first-user-message-template <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```
Add an assignment to a variation

```sh
cadenya agents variations add-assignment <agent-id> <variation-id> [--workspace-id <value>] --type <tool-id|tool-set-id|sub-agent-id> [--tool-id <value>] [--tool-set-id <value>] [--sub-agent-id <value>] [-f <doc>] [--dry-run]
```
Remove an assignment from a variation

```sh
cadenya agents variations remove-assignment <agent-id> <variation-id> <id> [--workspace-id <value>]
```
Attach a memory layer to a variation

```sh
cadenya agents variations add-memory-layer <agent-id> <variation-id> [--workspace-id <value>] --memory-layer-id <value> [--position <value>] [-f <doc>] [--dry-run]
```
Remove a memory layer assignment from a variation

```sh
cadenya agents variations remove-memory-layer <agent-id> <variation-id> <id> [--workspace-id <value>]
```
Update a variation's memory layer assignment

```sh
cadenya agents variations update-memory-layer <agent-id> <variation-id> <id> [--workspace-id <value>] [--position <value>] [-f <doc>] [--dry-run]
```

## cadenya ai-provider-keys

List AI provider keys

```sh
cadenya ai-provider-keys list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--promotional[=true|false]] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new AI provider key

```sh
cadenya ai-provider-keys create [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... [--provider <openrouter|openai|anthropic|gemini|openai-compatible>] [--credentials <api-key|headers>] [--api-key <value>] [--header KEY=VALUE]... [--config <openrouter|openai|openai-compatible>] [--openrouter-region <value>] [--openai-organization-id <value>] [--openai-project-id <value>] [--openai-compatible-base-url <value>] [-f <doc>] [--dry-run]
```
Get an AI provider key by ID

```sh
cadenya ai-provider-keys retrieve <id> [--workspace-id <value>]
```
Delete an AI provider key

```sh
cadenya ai-provider-keys delete <id> [--workspace-id <value>]
```
Update an AI provider key

```sh
cadenya ai-provider-keys update <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--provider <openrouter|openai|anthropic|gemini|openai-compatible>] [--credentials <api-key|headers>] [--api-key <value>] [--header KEY=VALUE]... [--config <openrouter|openai|openai-compatible>] [--openrouter-region <value>] [--openai-organization-id <value>] [--openai-project-id <value>] [--openai-compatible-base-url <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```

## cadenya memory-layers

List memory layers

```sh
cadenya memory-layers list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--type <value>] [--agent-id <value>] [--episodic-key-prefix <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new memory layer

```sh
cadenya memory-layers create [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... --type <episodic|skills> [--description <value>] [-f <doc>] [--dry-run]
```
Get a memory layer by ID

```sh
cadenya memory-layers retrieve <id> [--workspace-id <value>]
```
Delete a memory layer

```sh
cadenya memory-layers delete <id> [--workspace-id <value>]
```
Update a memory layer

```sh
cadenya memory-layers update <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--type <episodic|skills>] [--description <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```

## cadenya memory-layers entries

List memory entries

```sh
cadenya memory-layers entries list <memory-layer-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new memory entry

```sh
cadenya memory-layers entries create <memory-layer-id> [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... --type <content|upload-id> [--content <value>] [--key <value>] [--description <value>] [--upload-id <value>] [-f <doc>] [--dry-run]
```
Get a memory entry by ID

```sh
cadenya memory-layers entries retrieve <memory-layer-id> <id> [--workspace-id <value>]
```
Delete a memory entry

```sh
cadenya memory-layers entries delete <memory-layer-id> <id> [--workspace-id <value>]
```
Update a memory entry

```sh
cadenya memory-layers entries update <memory-layer-id> <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--key <value>] [--description <value>] [--content <value>] [--upload-id <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```

## cadenya models

List models

```sh
cadenya models list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--state <value>] [--ai-provider-key-id <value>] [--is-assigned[=true|false]] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Get a model by ID

```sh
cadenya models retrieve <id> [--workspace-id <value>]
```
Disable a model

```sh
cadenya models disable <id> [--workspace-id <value>]
```
Enable a model

```sh
cadenya models enable <id> [--workspace-id <value>]
```
Swap models on agent variations

```sh
cadenya models swap-on-variations [--workspace-id <value>] [--model-swap k=v,...|<doc>]... [-f <doc>] [--dry-run]
```

## cadenya objectives

List objectives

```sh
cadenya objectives list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--agent-id <value>] [--parent-objective-id <value>] [--state <value>] [--profile-id <value>] [--sort-order <value>] [--include-info[=true|false]] [--agent-schedule-id <value>] [--labels <value>] [--tenant-id <value>] [--subject-id <value>] [--widget-id <value>] [--widget-session-id <value>]
```
Create a new objective

```sh
cadenya objectives create [--workspace-id <value>] --agent-id <value> [--variation-id <value>] [--label KEY=VALUE]... [--external-id <value>] --system-prompt-data KEY=VALUE|<doc>... [--first-user-message <value>] [--secret k=v,...|<doc>]... [--memory-cascade k=v,...|<doc>]... [--first-user-message-data KEY=VALUE|<doc>]... [--episodic-memory-key <value>] [--tenant-id <value>] [--tenant-name <value>] [--subject-id <value>] [--subject-name <value>] [--pinned-parameter KEY=VALUE]... [-f <doc>] [--dry-run]
```
Get an objective by ID

```sh
cadenya objectives retrieve <id> [--workspace-id <value>]
```
List objective context windows

```sh
cadenya objectives list-context-windows <objective-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--include-info[=true|false]] [--labels <value>]
```
Get objective context usage

```sh
cadenya objectives retrieve-diagnostics <objective-id> [--workspace-id <value>]
```
List objective events

```sh
cadenya objectives list-events <objective-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--sort-order <value>] [--include-info[=true|false]] [--window-id <value>] [--since-event-id <value>] [--labels <value>]
```
Stream objective events

```sh
cadenya objectives stream-events <objective-id> [--workspace-id <value>] [--last-event-id <id>]
```
List feedback for an objective

```sh
cadenya objectives list-feedback <objective-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--labels <value>]
```
Submit feedback for an objective

```sh
cadenya objectives create-feedback <objective-id> [--workspace-id <value>] [--label KEY=VALUE]... [--external-id <value>] [--data-score <value>] [--data-comment <value>] [-f <doc>] [--dry-run]
```
List objective tool calls

```sh
cadenya objectives list-tool-calls <objective-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--status <value>] [--include-info[=true|false]] [--execution-status <value>] [--labels <value>]
```
Get an objective tool call by ID

```sh
cadenya objectives retrieve-tool-call <objective-id> <tool-call-id> [--workspace-id <value>]
```
Approve a tool call

```sh
cadenya objectives approve-tool-call <objective-id> <tool-call-id> [--workspace-id <value>]
```
Deny a tool call

```sh
cadenya objectives deny-tool-call <objective-id> <tool-call-id> [--workspace-id <value>] [--memo <value>] [-f <doc>] [--dry-run]
```
Set a bare tool call's content

```sh
cadenya objectives set-tool-call-content <objective-id> <tool-call-id> [--workspace-id <value>] --content <doc>... [-f <doc>] [--dry-run]
```
List objective tools

```sh
cadenya objectives list-tools <objective-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>]
```
Cancel an objective

```sh
cadenya objectives cancel <objective-id> [--workspace-id <value>] [--reason <value>] [-f <doc>] [--dry-run]
```
Compact an objective

```sh
cadenya objectives compact <objective-id> [--workspace-id <value>] [--compaction-config-trigger-threshold <value>] [--compaction-config-summarization-instructions <value>] [--compaction-config-tool-result-clearing-preserve-recent-results <value>] [-f <doc>] [--dry-run]
```
Continue an objective

```sh
cadenya objectives continue <objective-id> [--workspace-id <value>] --message <value> [--enqueue[=true|false]] [-f <doc>] [--dry-run]
```

## cadenya tool-search

Search for tools or tool sets

```sh
cadenya tool-search search-or-sets [--workspace-id <value>] --query <value>
```

## cadenya tenants

List tenants

```sh
cadenya tenants list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--query <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Get a tenant by ID

```sh
cadenya tenants retrieve <id> [--workspace-id <value>] [--include-info[=true|false]]
```
Erase a tenant

```sh
cadenya tenants delete <id> [--workspace-id <value>]
```
List a tenant's subjects

```sh
cadenya tenants list-subjects <tenant-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--query <value>] [--sort-order <value>] [--include-info[=true|false]]
```

## cadenya tool-sets

List tool sets

```sh
cadenya tool-sets list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--state <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new tool set

```sh
cadenya tool-sets create [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... [--description <value>] --adapter <mcp|http|openapi|bare> [--mcp-url <value>] [--mcp-header KEY=VALUE]... [--mcp-include-tools-filter <doc>]... [--mcp-include-tools-operator <and|or>] [--mcp-exclude-tools-filter <doc>]... [--mcp-exclude-tools-operator <and|or>] [--mcp-tool-approvals <always|only>] [--mcp-tool-approvals-always[=true|false]] [--mcp-tool-approvals-only-filter <doc>]... [--mcp-tool-approvals-only-operator <and|or>] [--mcp-just-in-time-enabled[=true|false]] [--mcp-just-in-time-fail-objective-on-tool-list-error[=true|false]] [--http-base-url <value>] [--http-header KEY=VALUE]... [--openapi <url|upload-id>] [--openapi-url <value>] [--openapi-header KEY=VALUE]... [--openapi-include-tools-filter <doc>]... [--openapi-include-tools-operator <and|or>] [--openapi-exclude-tools-filter <doc>]... [--openapi-exclude-tools-operator <and|or>] [--openapi-tool-approvals <always|only>] [--openapi-tool-approvals-always[=true|false]] [--openapi-tool-approvals-only-filter <doc>]... [--openapi-tool-approvals-only-operator <and|or>] [--openapi-base-url <value>] [--openapi-server-name <value>] [--openapi-upload-id <value>] [--bare-content-timeout <value>] [--overlay <doc>]... [-f <doc>] [--dry-run]
```
Get a tool set by ID

```sh
cadenya tool-sets retrieve <id> [--workspace-id <value>]
```
Delete a tool set

```sh
cadenya tool-sets delete <id> [--workspace-id <value>]
```
Update a tool set

```sh
cadenya tool-sets update <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--description <value>] [--adapter <mcp|http|openapi|bare>] [--mcp-url <value>] [--mcp-header KEY=VALUE]... [--mcp-include-tools-filter <doc>]... [--mcp-include-tools-operator <and|or>] [--mcp-exclude-tools-filter <doc>]... [--mcp-exclude-tools-operator <and|or>] [--mcp-tool-approvals <always|only>] [--mcp-tool-approvals-always[=true|false]] [--mcp-tool-approvals-only-filter <doc>]... [--mcp-tool-approvals-only-operator <and|or>] [--mcp-just-in-time-enabled[=true|false]] [--mcp-just-in-time-fail-objective-on-tool-list-error[=true|false]] [--http-base-url <value>] [--http-header KEY=VALUE]... [--openapi <url|upload-id>] [--openapi-url <value>] [--openapi-header KEY=VALUE]... [--openapi-include-tools-filter <doc>]... [--openapi-include-tools-operator <and|or>] [--openapi-exclude-tools-filter <doc>]... [--openapi-exclude-tools-operator <and|or>] [--openapi-tool-approvals <always|only>] [--openapi-tool-approvals-always[=true|false]] [--openapi-tool-approvals-only-filter <doc>]... [--openapi-tool-approvals-only-operator <and|or>] [--openapi-base-url <value>] [--openapi-server-name <value>] [--openapi-upload-id <value>] [--bare-content-timeout <value>] [--overlay <doc>]... [--update-mask <value>] [-f <doc>] [--dry-run]
```
Archive a tool set

```sh
cadenya tool-sets archive <id> [--workspace-id <value>]
```
Unarchive a tool set

```sh
cadenya tool-sets unarchive <id> [--workspace-id <value>]
```
List tool set events

```sh
cadenya tool-sets list-events <tool-set-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--sort-order <value>] [--include-info[=true|false]] [--labels <value>]
```
Get consumed OpenAPI spec

```sh
cadenya tool-sets retrieve-open-api-spec <tool-set-id> [--workspace-id <value>]
```
List tool set usage

```sh
cadenya tool-sets list-usage <tool-set-id> [--workspace-id <value>] [--tool-id <value>] [--limit <value>] [--cursor <value>] [--sort-order <value>]
```

## cadenya tool-sets secrets

List tool set secrets

```sh
cadenya tool-sets secrets list <tool-set-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new tool set secret

```sh
cadenya tool-sets secrets create <tool-set-id> [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... [--value <value>] [-f <doc>] [--dry-run]
```
Get a tool set secret by ID

```sh
cadenya tool-sets secrets retrieve <tool-set-id> <id> [--workspace-id <value>]
```
Delete a tool set secret

```sh
cadenya tool-sets secrets delete <tool-set-id> <id> [--workspace-id <value>]
```
Update a tool set secret

```sh
cadenya tool-sets secrets update <tool-set-id> <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--value <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```

## cadenya tool-sets tools

List tools

```sh
cadenya tool-sets tools list <tool-set-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--names <value>]... [--states <value>]... [--requires-approval[=true|false]] [--overlays <value>]... [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new tool

```sh
cadenya tool-sets tools create <tool-set-id> [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... --description <value> --requires-approval[=true|false] --parameter KEY=VALUE|<doc>... --config <http|mcp|openapi|bare> [--http-request-method <get|post|put|patch|delete>] [--http-path <value>] [--http-query <value>] [--http-header KEY=VALUE]... [--http-request-body-template <value>] [--http-request-body-content-type <value>] [--mcp-annotations-title <value>] [--mcp-annotations-read-only-hint[=true|false]] [--mcp-annotations-destructive-hint[=true|false]] [--mcp-annotations-idempotent-hint[=true|false]] [--mcp-annotations-open-world-hint[=true|false]] [--openapi-path <value>] [--openapi-method <value>] [--bare-always-set-result <value>] [--llm-tool-name <value>] [-f <doc>] [--dry-run]
```
Get a tool by ID

```sh
cadenya tool-sets tools retrieve <tool-set-id> <id> [--workspace-id <value>]
```
Delete a tool

```sh
cadenya tool-sets tools delete <tool-set-id> <id> [--workspace-id <value>]
```
Update a tool

```sh
cadenya tool-sets tools update <tool-set-id> <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--description <value>] [--requires-approval[=true|false]] [--parameter KEY=VALUE|<doc>]... [--config <http|mcp|openapi|bare>] [--http-request-method <get|post|put|patch|delete>] [--http-path <value>] [--http-query <value>] [--http-header KEY=VALUE]... [--http-request-body-template <value>] [--http-request-body-content-type <value>] [--mcp-annotations-title <value>] [--mcp-annotations-read-only-hint[=true|false]] [--mcp-annotations-destructive-hint[=true|false]] [--mcp-annotations-idempotent-hint[=true|false]] [--mcp-annotations-open-world-hint[=true|false]] [--openapi-path <value>] [--openapi-method <value>] [--bare-always-set-result <value>] [--llm-tool-name <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```
Omit a tool

```sh
cadenya tool-sets tools omit <tool-set-id> <id> [--workspace-id <value>]
```
Restore a tool

```sh
cadenya tool-sets tools restore <tool-set-id> <id> [--workspace-id <value>]
```

## cadenya uploads

Create an upload

```sh
cadenya uploads create [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... --filename <value> --content-type <value> --size-bytes <value> [-f <doc>] [--dry-run]
```
Get an upload by ID

```sh
cadenya uploads retrieve <id> [--workspace-id <value>]
```

## cadenya widget-sessions

List widget sessions

```sh
cadenya widget-sessions list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--widget-id <value>] [--tenant-id <value>] [--subject-id <value>] [--state <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a widget session

```sh
cadenya widget-sessions create [--workspace-id <value>] [--label KEY=VALUE]... [--external-id <value>] --widget-id <value> [--tenant-id <value>] [--tenant-name <value>] [--subject-id <value>] [--subject-name <value>] [--expires-at <value>] [--pinned-parameter KEY=VALUE]... [--secret k=v,...|<doc>]... [-f <doc>] [--dry-run]
```
Delete all of a tenant's widget sessions

```sh
cadenya widget-sessions delete-tenant [--workspace-id <value>] [--tenant-id <value>]
```
Get a widget session by ID

```sh
cadenya widget-sessions retrieve <id> [--workspace-id <value>]
```
Delete a widget session

```sh
cadenya widget-sessions delete <id> [--workspace-id <value>]
```
Revoke a widget session

```sh
cadenya widget-sessions revoke <id> [--workspace-id <value>]
```

## cadenya widgets

List widgets

```sh
cadenya widgets list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--agent-id <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new widget

```sh
cadenya widgets create [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... --agent-id <value> [--variation-id <value>] [--origin-allowlist <value>]... [-f <doc>] [--dry-run]
```
Get a widget by ID

```sh
cadenya widgets retrieve <id> [--workspace-id <value>]
```
Delete a widget

```sh
cadenya widgets delete <id> [--workspace-id <value>]
```
Update a widget

```sh
cadenya widgets update <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--agent-id <value>] [--variation-id <value>] [--origin-allowlist <value>]... [--update-mask <value>] [-f <doc>] [--dry-run]
```
Archive a widget

```sh
cadenya widgets archive <id> [--workspace-id <value>]
```
Unarchive a widget

```sh
cadenya widgets unarchive <id> [--workspace-id <value>]
```

## cadenya workspace-secrets

List workspace secrets

```sh
cadenya workspace-secrets list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new workspace secret

```sh
cadenya workspace-secrets create [--workspace-id <value>] --name <value> [--external-id <value>] [--label KEY=VALUE]... [--value <value>] [-f <doc>] [--dry-run]
```
Get a workspace secret by ID

```sh
cadenya workspace-secrets retrieve <id> [--workspace-id <value>]
```
Delete a workspace secret

```sh
cadenya workspace-secrets delete <id> [--workspace-id <value>]
```
Update a workspace secret

```sh
cadenya workspace-secrets update <id> [--workspace-id <value>] [--name <value>] [--external-id <value>] [--label KEY=VALUE]... [--value <value>] [--update-mask <value>] [-f <doc>] [--dry-run]
```
