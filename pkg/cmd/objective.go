// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/option"

	"github.com/cadenya/cadenya-cli/internal/apiquery"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var objectivesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new objective in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
			BodyPath: "agentId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "system-prompt-data",
			Usage:    "Arbitrary data rendered into the selected variation's system_prompt_template\n (liquid) to produce the objective's system prompt. If the agent has a\n system_prompt_data_schema, this must satisfy it.",
			Required: true,
			BodyPath: "systemPromptData",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "episodic-memory",
			Usage:    "Episodic is used to configure the episodic memory for the objective",
			BodyPath: "episodicMemory",
		},
		&requestflag.Flag[string]{
			Name:     "first-user-message",
			Usage:    "Optional explicit first user message for the LLM chat history. When not set,\n the selected variation's first_user_message_template is rendered with\n first_user_message_data instead. If neither this field nor a\n first_user_message_template is present, the request is rejected with InvalidArgument.",
			BodyPath: "firstUserMessage",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "first-user-message-data",
			Usage:    "Arbitrary data rendered into the selected variation's first_user_message_template\n (liquid) to produce the first user message. Separate from `system_prompt_data`,\n which renders the system prompt template.",
			BodyPath: "firstUserMessageData",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "memory-cascade",
			Usage:    "Memory layers/entries layered over the baseline cascade inherited\n from the selected variation — element-level rules over inherited\n styles, in CSS terms.\n\n Array order is resolution order: EARLIER elements are more specific\n and are consulted first. Entries pinned via memory_entry_id behave\n as single-entry layers at their position.\n\n System-managed layers (e.g., episodic) cannot be referenced here;\n they attach themselves automatically based on the episodic key.\n\n Size cap: the TOTAL effective cascade (this field + the variation's\n memory layer assignments) must not exceed 10 entries. A request\n that would produce a larger cascade is rejected with\n InvalidArgument.",
			BodyPath: "memoryCascade",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "CreateOperationMetadata contains the user-provided fields for creating\n an operation. Read-only fields (id, account_id, workspace_id, created_at, profile_id)\n are excluded since they are set by the server.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pinned-parameters",
			Usage:    "Parameters forced onto this objective's tool calls. A pinned parameter\n is an overlay on a tool's JSON schema: the parameter is removed from\n what the LLM sees, and its value is always overwritten server-side with\n the pinned value — the model cannot choose a different value for it.",
			BodyPath: "pinnedParameters",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "secret",
			Usage:    "Secrets that can be used in the headers for tool calls using the secret interpolation format.",
			BodyPath: "secrets",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "subject",
			Usage:    "SubjectAssertion identifies a person within a tenant in the customer's own\n namespace — typically their user id. Asserting a subject upserts the\n subject record under the asserted tenant and associates the created\n resource with it. A subject assertion is only valid alongside a tenant\n assertion: subject identifiers are scoped to their tenant.",
			BodyPath: "subject",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "tenant",
			Usage:    "TenantAssertion identifies a tenant in the customer's own namespace — their\n org, company, or team identifier for an end user. Asserting a tenant\n upserts the tenant record in the workspace (keyed on `id` as the tenant's\n external_id) and associates the created resource with it.",
			BodyPath: "tenant",
		},
		&requestflag.Flag[string]{
			Name:     "variation-id",
			Usage:    "Optional explicit variation selection. Overrides the agent's variation_selection_mode.",
			BodyPath: "variationId",
		},
	},
	Action:          handleObjectivesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"episodic-memory": {
		&requestflag.InnerFlag[string]{
			Name:       "episodic-memory.key",
			Usage:      "The caller-supplied episodic key. Objectives created with the same key\n (for the same agent) share one episodic memory layer.",
			InnerField: "key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "episodic-memory.memory-layer-id",
			Usage:      "The episodic memory layer resolved (created or reused) for this\n objective's key. Populated by the system at objective creation.",
			InnerField: "memoryLayerId",
		},
	},
	"memory-cascade": {
		&requestflag.InnerFlag[string]{
			Name:       "memory-cascade.memory-layer-id",
			InnerField: "memoryLayerId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "memory-cascade.memory-entry-id",
			Usage:      "When set, inserts only this entry from memory_layer_id into the cascade —\n behaves as a single-entry layer (only this key resolves at this\n position). The entry must belong to memory_layer_id; mismatches are\n rejected with InvalidArgument.",
			InnerField: "memoryEntryId",
		},
	},
	"metadata": {
		&requestflag.InnerFlag[string]{
			Name:       "metadata.external-id",
			Usage:      "External ID for the operation (e.g., a workflow ID from an external system)",
			InnerField: "externalId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "metadata.labels",
			Usage:      "Key-value pairs for categorization and filtering. Values are 0-63\n alphanumeric characters with \"-\", \"_\", or \".\" allowed between; keys\n follow the same shape and additionally accept an optional DNS-subdomain\n prefix (e.g. \"cadenya.com/\") of at most 253 characters.\n Examples: {\"priority\": \"high\", \"source\": \"api\", \"workflow\": \"onboarding\"}",
			InnerField: "labels",
		},
	},
	"secret": {
		&requestflag.InnerFlag[string]{
			Name:       "secret.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "secret.value",
			InnerField: "value",
		},
	},
	"subject": {
		&requestflag.InnerFlag[string]{
			Name:       "subject.id",
			Usage:      "The subject identifier in the customer's namespace (e.g. their user id).\n Stored as the subject record's external_id; unique within the tenant.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "subject.name",
			Usage:      "Optional human-readable name for the subject. Updates the subject\n record's name on every assertion that provides it.",
			InnerField: "name",
		},
	},
	"tenant": {
		&requestflag.InnerFlag[string]{
			Name:       "tenant.id",
			Usage:      "The tenant identifier in the customer's namespace (e.g. \"acme-corp\").\n Stored as the tenant record's external_id; stable across requests.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tenant.name",
			Usage:      "Optional human-readable name for the tenant. Updates the tenant record's\n name on every assertion that provides it.",
			InnerField: "name",
		},
	},
})

var objectivesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves an objective by ID from the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleObjectivesRetrieve,
	HideHelpCommand: true,
}

var objectivesList = cli.Command{
	Name:    "list",
	Usage:   "Lists all objectives in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Usage:     "Agent ID for filtering",
			QueryPath: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-schedule-id",
			Usage:     "Filter to objectives produced by a specific AgentSchedule. Accepts\n canonical as_… form or external_id:<value> form.",
			QueryPath: "agentScheduleId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[bool]{
			Name:      "include-info",
			Usage:     "When set to true you may use more of your alloted API rate-limit",
			QueryPath: "includeInfo",
		},
		&requestflag.Flag[string]{
			Name:      "labels",
			Usage:     "Filters by metadata labels. Comma-separated key=value pairs,\n e.g. \"env=prod,team=ai\". A resource matches only if every pair\n matches exactly (AND semantics).",
			QueryPath: "labels",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "parent-objective-id",
			Usage:     "Optional filters",
			QueryPath: "parentObjectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "profile-id",
			QueryPath: "profileId",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort order for results (asc or desc by creation time)",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Filter by state",
			QueryPath: "state",
		},
		&requestflag.Flag[string]{
			Name:      "subject-id",
			Usage:     "Filter to objectives associated with a subject. Accepts the canonical\n `subj_…` form or the `external_id:<value>` form; the external_id form is\n scoped within a tenant and requires `tenant_id` to also be set.",
			QueryPath: "subjectId",
		},
		&requestflag.Flag[string]{
			Name:      "tenant-id",
			Usage:     "Filter to objectives associated with a tenant. Accepts the canonical\n `tenant_…` form or the `external_id:<value>` form.",
			QueryPath: "tenantId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesList,
	HideHelpCommand: true,
}

var objectivesCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancels a running or pending objective. The objective's state will be set to\nSTATE_CANCELLED.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Optional reason for cancellation",
			BodyPath: "reason",
		},
	},
	Action:          handleObjectivesCancel,
	HideHelpCommand: true,
}

var objectivesCompact = requestflag.WithInnerFlags(cli.Command{
	Name:    "compact",
	Usage:   "Triggers compaction on a running objective. Optionally override the variation's\ncompaction config.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "compaction-config",
			Usage:    "CompactionConfig defines how context window compaction behaves for objectives using this variation.",
			BodyPath: "compactionConfig",
		},
	},
	Action:          handleObjectivesCompact,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"compaction-config": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "compaction-config.summarization",
			Usage:      "SummarizationStrategy configures LLM-powered summarization of older conversation turns.",
			InnerField: "summarization",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "compaction-config.tool-result-clearing",
			Usage:      "ToolResultClearingStrategy configures clearing of older tool result content.",
			InnerField: "toolResultClearing",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "compaction-config.trigger-threshold",
			Usage:      "Trigger threshold as a percentage of the model's context window (0.0 to 1.0).\n When input tokens reach this percentage of the model's limit, compaction triggers.\n Default: 0.75 (75%)",
			InnerField: "triggerThreshold",
		},
	},
})

var objectivesContinue = cli.Command{
	Name:    "continue",
	Usage:   "Continues an objective that has completed",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:     "message",
			Usage:    "The message to continue an objective that has completed (or you are enqueing)",
			Required: true,
			BodyPath: "message",
		},
		&requestflag.Flag[bool]{
			Name:     "enqueue",
			Usage:    "When set to true, the message will be enqueued for when the agent loop is available to process it.",
			BodyPath: "enqueue",
		},
	},
	Action:          handleObjectivesContinue,
	HideHelpCommand: true,
}

var objectivesListContextWindows = cli.Command{
	Name:    "list-context-windows",
	Usage:   "Read-only list of the last five windows of execution for this objective, ordered\nby most recent first",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[bool]{
			Name:      "include-info",
			Usage:     "When set to true you may use more of your alloted API rate-limit",
			QueryPath: "includeInfo",
		},
		&requestflag.Flag[string]{
			Name:      "labels",
			Usage:     "Filters by metadata labels. Comma-separated key=value pairs,\n e.g. \"env=prod,team=ai\". A resource matches only if every pair\n matches exactly (AND semantics).",
			QueryPath: "labels",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesListContextWindows,
	HideHelpCommand: true,
}

var objectivesListEvents = cli.Command{
	Name:    "list-events",
	Usage:   "Lists all events for an objective",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[bool]{
			Name:      "include-info",
			Usage:     "When set to true you may use more of your alloted API rate-limit",
			QueryPath: "includeInfo",
		},
		&requestflag.Flag[string]{
			Name:      "labels",
			Usage:     "Filters by metadata labels. Comma-separated key=value pairs,\n e.g. \"env=prod,team=ai\". A resource matches only if every pair\n matches exactly (AND semantics).",
			QueryPath: "labels",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "since-event-id",
			Usage:     "Optional string to fetch events since an ID",
			QueryPath: "sinceEventId",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort order for results (asc or desc by creation time)",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "window-id",
			Usage:     "Optional context window ID to filter events by",
			QueryPath: "windowId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesListEvents,
	HideHelpCommand: true,
}

var objectivesRetrieveDiagnostics = cli.Command{
	Name:    "retrieve-diagnostics",
	Usage:   "Returns the context-usage breakdown measured for the objective's most recent\niteration: character lengths per context component (system prompt, memory\nappendices, tool definitions, messages by role) alongside the iteration's input\ntoken counts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
	},
	Action:          handleObjectivesRetrieveDiagnostics,
	HideHelpCommand: true,
}

var objectivesStreamEvents = cli.Command{
	Name:    "stream-events",
	Usage:   "Streams events for an objective in real-time using server-sent events (SSE)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesStreamEvents,
	HideHelpCommand: true,
}

func handleObjectivesCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveNewParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives create",
		Transform:      transform,
	})
}

func handleObjectivesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveGetParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.Get(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives retrieve",
		Transform:      transform,
	})
}

func handleObjectivesList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveListParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Objectives.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "objectives list",
			Transform:      transform,
		})
	} else {
		iter := client.Objectives.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "objectives list",
			Transform:      transform,
		})
	}
}

func handleObjectivesCancel(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveCancelParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.Cancel(
		ctx,
		cmd.Value("objective-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives cancel",
		Transform:      transform,
	})
}

func handleObjectivesCompact(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveCompactParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.Compact(
		ctx,
		cmd.Value("objective-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives compact",
		Transform:      transform,
	})
}

func handleObjectivesContinue(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveContinueParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.Continue(
		ctx,
		cmd.Value("objective-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives continue",
		Transform:      transform,
	})
}

func handleObjectivesListContextWindows(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveListContextWindowsParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Objectives.ListContextWindows(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "objectives list-context-windows",
			Transform:      transform,
		})
	} else {
		iter := client.Objectives.ListContextWindowsAutoPaging(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "objectives list-context-windows",
			Transform:      transform,
		})
	}
}

func handleObjectivesListEvents(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveListEventsParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Objectives.ListEvents(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "objectives list-events",
			Transform:      transform,
		})
	} else {
		iter := client.Objectives.ListEventsAutoPaging(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "objectives list-events",
			Transform:      transform,
		})
	}
}

func handleObjectivesRetrieveDiagnostics(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveGetDiagnosticsParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.GetDiagnostics(
		ctx,
		cmd.Value("objective-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives retrieve-diagnostics",
		Transform:      transform,
	})
}

func handleObjectivesStreamEvents(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveStreamEventsParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	stream := client.Objectives.StreamEventsStreaming(
		ctx,
		cmd.Value("objective-id").(string),
		params,
		options...,
	)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(stream, maxItems, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives stream-events",
		Transform:      transform,
	})
}
