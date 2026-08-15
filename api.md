# cadenya CLI reference

JSON-valued flags accept a literal document, `@file`, or `-` (stdin; at most ONE flag per invocation may use it). See README.md for usage patterns.

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
cadenya api-keys create [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya api-keys update <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
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
cadenya workspace-admin create --metadata <JSON> --spec <JSON>
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
cadenya workspace-admin update [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
```
List workspace members

```sh
cadenya workspace-admin list-members [--workspace-id <value>] [--limit <value>] [--cursor <value>]
```
Add a member to a workspace

```sh
cadenya workspace-admin add-member [--workspace-id <value>] [--profile-id <value>] [--email <value>]
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
cadenya agents create [--workspace-id <value>] --metadata <JSON> --spec <JSON> [--default-variation <JSON>]
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
cadenya agents update <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
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
cadenya agents schedules create <agent-id> [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya agents schedules update <agent-id> <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
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
cadenya agents variations create <agent-id> [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya agents variations update <agent-id> <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
```
Add an assignment to a variation

```sh
cadenya agents variations add-assignment <agent-id> <variation-id> [--workspace-id <value>] --body <JSON>
```
Remove an assignment from a variation

```sh
cadenya agents variations remove-assignment <agent-id> <variation-id> <id> [--workspace-id <value>]
```
Attach a memory layer to a variation

```sh
cadenya agents variations add-memory-layer <agent-id> <variation-id> [--workspace-id <value>] --memory-layer-id <value> [--position <value>]
```
Remove a memory layer assignment from a variation

```sh
cadenya agents variations remove-memory-layer <agent-id> <variation-id> <id> [--workspace-id <value>]
```
Update a variation's memory layer assignment

```sh
cadenya agents variations update-memory-layer <agent-id> <variation-id> <id> [--workspace-id <value>] [--position <value>]
```

## cadenya ai-provider-keys

List AI provider keys

```sh
cadenya ai-provider-keys list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--promotional[=true|false]] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new AI provider key

```sh
cadenya ai-provider-keys create [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya ai-provider-keys update <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
```

## cadenya memory-layers

List memory layers

```sh
cadenya memory-layers list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--type <value>] [--agent-id <value>] [--episodic-key-prefix <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new memory layer

```sh
cadenya memory-layers create [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya memory-layers update <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
```

## cadenya memory-layers entries

List memory entries

```sh
cadenya memory-layers entries list <memory-layer-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new memory entry

```sh
cadenya memory-layers entries create <memory-layer-id> [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya memory-layers entries update <memory-layer-id> <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
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
cadenya models swap-on-variations [--workspace-id <value>] [--model-swaps <JSON>]...
```

## cadenya objectives

List objectives

```sh
cadenya objectives list [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--agent-id <value>] [--parent-objective-id <value>] [--state <value>] [--profile-id <value>] [--sort-order <value>] [--include-info[=true|false]] [--agent-schedule-id <value>] [--labels <value>] [--tenant-id <value>] [--subject-id <value>] [--widget-id <value>] [--widget-session-id <value>]
```
Create a new objective

```sh
cadenya objectives create [--workspace-id <value>] --agent-id <value> [--variation-id <value>] [--metadata <JSON>] --system-prompt-data <JSON> [--first-user-message <value>] [--secrets <JSON>]... [--memory-cascade <JSON>]... [--first-user-message-data <JSON>] [--episodic-memory <JSON>] [--tenant <JSON>] [--subject <JSON>] [--pinned-parameters <JSON>]
```
Get an objective by ID

```sh
cadenya objectives retrieve <id> [--workspace-id <value>]
```
List objective context windows

```sh
cadenya objectives list-context-windows <objective-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--include-info[=true|false]] [--labels <value>]
```
Get objective context diagnostics

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
cadenya objectives create-feedback <objective-id> [--workspace-id <value>] --metadata <JSON> --data <JSON>
```
List objective tasks

```sh
cadenya objectives list-tasks <objective-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--sort-order <value>]
```
Get an objective task by ID

```sh
cadenya objectives retrieve-task <objective-id> <id> [--workspace-id <value>]
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
cadenya objectives deny-tool-call <objective-id> <tool-call-id> [--workspace-id <value>] [--memo <value>]
```
Set a bare tool call's content

```sh
cadenya objectives set-tool-call-content <objective-id> <tool-call-id> [--workspace-id <value>] --content <JSON>...
```
List objective tools

```sh
cadenya objectives list-tools <objective-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>]
```
Cancel an objective

```sh
cadenya objectives cancel <objective-id> [--workspace-id <value>] [--reason <value>]
```
Compact an objective

```sh
cadenya objectives compact <objective-id> [--workspace-id <value>] [--compaction-config <JSON>]
```
Continue an objective

```sh
cadenya objectives continue <objective-id> [--workspace-id <value>] --message <value> [--enqueue[=true|false]]
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
cadenya tool-sets create [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya tool-sets update <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
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
cadenya tool-sets secrets create <tool-set-id> [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya tool-sets secrets update <tool-set-id> <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
```

## cadenya tool-sets tools

List tools

```sh
cadenya tool-sets tools list <tool-set-id> [--workspace-id <value>] [--limit <value>] [--cursor <value>] [--prefix <value>] [--query <value>] [--names <value>]... [--states <value>]... [--requires-approval[=true|false]] [--labels <value>] [--sort-order <value>] [--include-info[=true|false]]
```
Create a new tool

```sh
cadenya tool-sets tools create <tool-set-id> [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya tool-sets tools update <tool-set-id> <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
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
cadenya uploads create [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya widget-sessions create [--workspace-id <value>] [--metadata <JSON>] --spec <JSON> [--secrets <JSON>]...
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
cadenya widgets create [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya widgets update <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
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
cadenya workspace-secrets create [--workspace-id <value>] --metadata <JSON> --spec <JSON>
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
cadenya workspace-secrets update <id> [--workspace-id <value>] [--metadata <JSON>] [--spec <JSON>] [--update-mask <value>]
```
