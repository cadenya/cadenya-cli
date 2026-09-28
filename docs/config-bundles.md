# Workspace configuration bundles

Use `cadenya config validate`, `cadenya config plan`, and `cadenya config apply` to manage workspace resources from YAML.

```sh
cadenya config validate -C examples/basic
cadenya config plan -C examples/basic
cadenya config apply -C examples/basic
```

## How ownership works

Every resource the action writes gets a label: `bundle_key=<your bundle key>`. That label is the whole ownership model. Plan and apply list everything that carries it, compare that with `.cadenya/`, and act on the difference.

**Anything with your bundle key that isn't in `.cadenya/` gets deleted on the next apply.** That's the point: delete a file, apply, and the resource is gone. But it means the bundle key has to be yours alone. Pick one per repository (or one per bundle, if a repository has several) and don't share it.

The label also keeps the action away from everything else in the workspace:

1. It never adopts a resource without your label. If an external ID you're about to create already exists outside the bundle, plan fails and names it.
2. It never deletes a parent with children outside the bundle. An agent with a hand-made variation stays put, and so does a Bare or HTTP tool set holding someone else's tool, or a memory layer holding someone else's entry.
3. It never deletes something a resource outside the bundle depends on: a tool, tool set, memory layer, or sub-agent that another variation assigns, a variation another widget pins, or an agent another widget is bound to. Plan names the resource that's in the way. (Checking this lists every agent, variation, and widget in the workspace, so it only happens when the plan deletes something.)
4. It refuses to apply an empty `.cadenya/` unless you set `allow-empty: true`. An empty tree means "delete the whole bundle." That deserves a flag.

Changing the bundle key doesn't rename or move anything. It starts a new, empty scope, and the old resources keep the old label.

## Bundle layout

```text
cadenya.yaml                     # settings, optional
.cadenya/
  toolSets/
    ask-cadenya.yaml             # MCP
    frontend.yaml                # Bare
    frontend/
      fetch-form-state.yaml      # a tool in the frontend tool set
    petstore.yaml                # OpenAPI
  memoryLayers/
    playbooks.yaml
    playbooks/
      tool-set-setup.yaml        # an entry in the playbooks layer
  agents/
    ask-cadenya.yaml
    ask-cadenya/
      default.yaml               # a variation of the ask-cadenya agent
  widgets/
    support-chat.yaml            # embeds the ask-cadenya agent
```

One resource per file, as `.yaml` or `.yml`. A child directory takes its name from the parent's **filename**, even when the parent sets its own external ID. A child's identity includes its parent, so two agents can each have a `default` variation.

The loader is strict on purpose. It rejects unknown fields (typos inside an adapter included), unknown directories under `.cadenya/`, child directories with no parent file, symlinks, YAML aliases, and duplicate identities. It skips files that aren't YAML, so a `README.md` in there is fine.

A missing `toolSets/`, `memoryLayers/`, `agents/`, or `widgets/` directory means you want none of that type. **So removing `agents/` deletes every agent in the bundle.** The `.cadenya/` directory itself has to exist.

## Settings

`cadenya.yaml` sits in the action's `directory` (the repository root, by default):

```yaml
bundleKey: support-tools
workspaceId: development
baseUrl: https://api.cadenya.com
resourceDir: .cadenya
```

Every key is optional in the file, but `bundleKey` has to come from somewhere: the file, the `bundle-key` input, or `CADENYA_BUNDLE_KEY`. It's a label value: 1 to 63 letters, digits, `-`, `_`, or `.`, starting and ending with a letter or digit.

Changing the bundle key doesn't rename or move anything. It starts a new, empty scope, and the old resources keep the old label.

Keep the API key in a secret. It never belongs in `cadenya.yaml`.

## Resource files

Each file is the SDK's JSON shape, written as YAML: `metadata` and `spec`. The external ID comes from a top-level `externalId`, then `metadata.externalId`, then the filename without its extension. If you set both fields, they have to match. `metadata.name` defaults to the external ID.

**Renaming a file without an explicit external ID changes its identity.** Plan shows a delete and a create, not a rename.

### Tool set

`.cadenya/toolSets/frontend.yaml`:

```yaml
metadata:
  name: Frontend tools
  labels:
    team: frontend
spec:
  description: Tools executed by the frontend application
  adapter:
    type: bare
    bare: {}
```

Your own labels ride along with `bundle_key`. Setting `bundle_key` yourself to anything but the bundle key is an error.

MCP and OpenAPI tool sets need nothing but the adapter. Here's the Swagger Petstore, from `.cadenya/toolSets/petstore.yaml`:

```yaml
metadata:
  name: Swagger Petstore
spec:
  description: Find pets, place orders, and manage the store inventory
  adapter:
    type: openapi
    openapi:
      type: url
      url: https://petstore3.swagger.io/api/v3/openapi.json
```

Cadenya syncs the tools once the tool set exists (19 of them, for the Petstore). Every apply sends an update, and an update starts another sync. A sync that finds nothing new changes nothing.

### Tool

`.cadenya/toolSets/frontend/fetch-form-state.yaml`:

```yaml
metadata:
  name: Fetch form state
spec:
  description: Return the current form fields and their values
  parameters:
    type: object
    properties: {}
    additionalProperties: false
  config:
    type: bare
    bare: {}
```

Tools go under Bare and HTTP tool sets, and a tool's `config` type has to match its parent's adapter. MCP and OpenAPI tool sets sync their own tools, so there's nothing to write for them. `parameters` is required, and `{}` works for a tool that takes no arguments.

An HTTP pair looks like this: the tool set gets `adapter: {type: http, http: {baseUrl: ...}}`, and each tool gets `config: {type: http, http: {requestMethod: GET, path: /forms}}`.

### Agent

`.cadenya/agents/ask-cadenya.yaml`:

```yaml
externalId: cadenya-assistant
metadata:
  name: Ask Cadenya
spec:
  description: Answer questions about Cadenya and the current form
```

In the API, this agent is `cadenya-assistant`. Its variations still live in `agents/ask-cadenya/`, because the directory follows the filename.

### Agent state

An agent file takes one more top-level field, `state`: `draft` or `published`. The example publishes its agent:

```yaml
externalId: cadenya-assistant
state: published
metadata:
  name: Ask Cadenya
spec:
  description: Answer questions about Cadenya and the current form
```

Apply publishes after the agent's variations exist, because Cadenya won't publish an agent without one. **So `state: published` needs at least one variation in the bundle, and validate checks that.** An archived agent with a `state` gets unarchived first.

The same rule runs the other way. If an agent is published in Cadenya and the plan would delete its last variation, plan stops, even when the YAML leaves `state` out. Keep a variation, or set `state: draft`.

Leave `state` out and the action leaves the state alone, the same as any other field you don't write. New agents start as drafts.

### Variation

`.cadenya/agents/ask-cadenya/default.yaml`:

```yaml
metadata:
  name: Default
spec:
  systemPromptTemplate: |
    Help the user configure Cadenya. Consult the documentation tools for API
    questions, and fetch the current form state when needed.
  modelConfig:
    modelId: getting-started-key.openai-gpt-4-1
    temperature: 0
  assignments:
    - type: toolSetId
      toolSetId: external_id:ask-cadenya
    - type: toolId
      toolId: external_id:frontend/fetch-form-state
```

`modelId` is a reference key: your AI provider key's external ID, a dot, then the model's external ID. Or use its canonical `model_` ID. The model's external ID has no dots of its own, so GPT-4.1 is `openai-gpt-4-1`, not `gpt-4.1`. **`openai.gpt-4.1` fails, and validate tells you so before apply can.**

Every Cadenya account starts with a provider key whose external ID is `getting-started-key`, so the example works as written. A provider key you add has its own external ID. To see the models a workspace can use, list them (that takes the `models:read` scope).

### Memory layer

`.cadenya/memoryLayers/playbooks.yaml`:

```yaml
metadata:
  name: Support playbooks
spec:
  type: MEMORY_LAYER_TYPE_SKILLS
  description: How the assistant handles common configuration questions
```

`type` is required, and `MEMORY_LAYER_TYPE_SKILLS` is the one a bundle can create. (Episodic layers belong to the runtime.) **A layer's type can't change after it exists.** Plan stops if you try, so give the layer a new external ID to replace it.

### Memory entry

`.cadenya/memoryLayers/playbooks/tool-set-setup.yaml`:

```yaml
metadata:
  name: Set up a tool set
spec:
  key: skills/tool-set-setup
  description: Use when someone asks how to connect an API or MCP server to an agent.
  content: |
    1. Ask whether the API has an OpenAPI spec or an MCP server.
    2. For OpenAPI, point a tool set at the spec URL. Cadenya syncs it hourly.
    3. For MCP, use the server URL. Tools sync once the tool set exists.
    4. Filter the tools before assigning the tool set to a variation.
```

`key` is what the model passes to `get_memory`, and it defaults to the external ID. Keys can hold slashes and external IDs can't, so a hierarchical key like `skills/tool-set-setup` needs its own `key` field. Validate applies Cadenya's key rules: ASCII letters, digits, and `! - _ . * ' ( ) /`, no leading, trailing, or doubled slash, and nothing under `cadenya/` or `system/`.

For a skills layer, `description` is the "when to use this" line the model sees before it loads the body. Write it for the model.

To put a layer in a variation's memory cascade, add it to `memoryLayerAssignments`:

```yaml
  memoryLayerAssignments:
    - memoryLayerId: external_id:playbooks
      position: 0
```

Lower positions are consulted first.

### Widget

`.cadenya/widgets/support-chat.yaml`:

```yaml
metadata:
  name: Support chat
spec:
  agentId: external_id:cadenya-assistant
  originAllowlist:
    - http://localhost:3000
```

A widget embeds one agent. `originAllowlist` holds exact origins: a scheme, a host, and an optional port. No paths, no wildcards, and validate says so before Cadenya has to.

`variationId` pins every conversation to one of the agent's variations, as `external_id:<agent>/<variation>`. Leave it out and the agent's selection mode picks. To unpin, set `variationId: null`. **Delete a pinned variation and plan asks you to set `variationId` first**, since Cadenya won't delete a variation out from under its widget.

Widgets have no children, so `widgets/` holds files only. There's no `state` for widgets.

### References

Variations and widgets point at resources in the same bundle by external ID. Apply swaps in the canonical IDs as it creates things, so a variation can reference a tool created earlier in the same run:

| Assignment | Reference |
| --- | --- |
| `toolSetId` | `external_id:<tool set>` |
| `toolId` | `external_id:<tool set>/<tool>` |
| `subAgentId` | `external_id:<agent>` |
| `memoryLayerAssignments[].memoryLayerId` | `external_id:<memory layer>` |
| Widget `agentId` | `external_id:<agent>` |
| Widget `variationId` | `external_id:<agent>/<variation>` |

Validate checks every one of these before any API call. For a resource outside the bundle, use its canonical ID. Agent pools work the same way: reference them by ID. Bundles don't create them.

## What an update changes

Every update sends the name, the external ID, and the complete label map. For `spec`, the top-level fields in your YAML become the update mask:

1. A field you leave out keeps whatever value it has in Cadenya.
2. A field you include replaces the remote value whole, nested objects and all.
3. An empty value clears the field. `assignments: []` removes every assignment, and `{}`, `null`, `""`, `false`, and `0` clear the same way.

Memory entries follow the same rule, and `content` counts as a field: leave it out and the entry's body stays as it is.

Nothing is interpolated. `${SECRET_NAME}` references and Liquid templates reach Cadenya as written.

### Why does a clean plan list every resource?

Because every resource that exists gets an update on every apply, changed or not. Run plan right after an apply and you see:

```
update    toolSet ask-cadenya
update    toolSet frontend
update    toolSet petstore
update    tool frontend/fetch-form-state
update    memoryLayer playbooks
update    memoryEntry playbooks/tool-set-setup
update    agent cadenya-assistant
update    agentVariation cadenya-assistant/default
update    widget support-chat
Plan: 0 create, 9 update, 0 delete, 0 detach
```

There's no second `publish`. State changes only show up when the state differs.

**Plan lists operations, not a field-by-field diff.** That's the tradeoff. You can't see which fields would change. In return, the action never compares its idea of a field with the server's, and every field you wrote matches the YAML after each apply.

## What apply does

Plan and apply read everything before they write anything:

1. Load and validate the whole bundle.
2. List every resource with your label, archived ones included, across every page. (A server that repeats a pagination cursor fails the plan instead of looping forever.)
3. Check that nothing you're about to create exists outside the bundle, and that nothing you're about to delete has children outside it.

Then apply writes, in this order:

4. Detach: remove assignments to resources that are about to be deleted.
5. Unarchive or unpublish agents.
6. Delete, children first: widgets, variations, agents, memory entries, memory layers, tools, then tool sets.
7. Create or update: tool sets, tools, memory layers, memory entries, agents, variations, then widgets.
8. Publish agents.
9. Delete what had to wait: variations removed from agents that stay, and agents a widget is moving off of.

Why the split between steps 6 and 9? Because Cadenya won't delete the last variation of a published agent, or an agent a widget is bound to. Swap a published agent's only variation for a new one and the new one has to exist first. Move a widget to a new agent and the widget has to move (step 7) before the old agent can go. A published agent you delete outright gets unpublished in step 5, so its variations can go.

Say you delete `frontend/fetch-form-state.yaml` and take its assignment out of `default.yaml`. Plan shows the detach first:

```
detach    agentVariation cadenya-assistant/default
delete    tool frontend/fetch-form-state
update    toolSet ask-cadenya
update    toolSet frontend
update    toolSet petstore
update    memoryLayer playbooks
update    memoryEntry playbooks/tool-set-setup
update    agent cadenya-assistant
update    agentVariation cadenya-assistant/default
update    widget support-chat
Plan: 0 create, 8 update, 1 delete, 1 detach
```

A detach needs the field it changes in the variation's YAML: `assignments` for tools, tool sets, and sub-agents, `memoryLayerAssignments` for memory layers. Without it, plan stops and names the field to add (`[]` works). The detach matters for memory layers in particular: Cadenya refuses to delete a layer that a variation still assigns.

**Apply isn't transactional.** It stops at the first failed request and reports what finished. It can't roll back writes that already landed. Fix the cause and run it again: every run builds a fresh plan from what's in Cadenya now.

Two applies against the same bundle at the same time race each other. Give each workspace and bundle one `concurrency` group, as the quick start does.

The action manages configuration and stops there. Beyond `state`, lifecycle changes (archiving, for one) are out of scope, along with secrets and waiting for an MCP or OpenAPI tool set to finish syncing.
