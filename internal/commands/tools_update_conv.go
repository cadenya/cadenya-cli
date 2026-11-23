package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

const bodySchemaToolsUpdate = "{\"$defs\":{\"AutoModeConfig\":{\"properties\":{\"failureMessage\":{\"description\":\"Optional guidance returned to the agent when the decision thresholds fail.\\n When empty, a default denial message is used. Client errors never use this message.\",\"type\":\"string\"},\"passingMode\":{\"$ref\":\"AutoModeConfigPassingMode\"},\"questions\":{\"items\":{\"$ref\":\"AutoModeQuestion\"},\"type\":\"array\"}},\"required\":[\"questions\",\"passingMode\"],\"type\":\"object\"},\"AutoModeConfigPassingMode\":{\"enum\":[\"PASSING_MODE_ALL\",\"PASSING_MODE_NONE\",\"PASSING_MODE_MOST\"],\"enumShort\":{\"all\":\"PASSING_MODE_ALL\",\"most\":\"PASSING_MODE_MOST\",\"none\":\"PASSING_MODE_NONE\"},\"type\":\"string\"},\"AutoModeQuestion\":{\"properties\":{\"question\":{\"description\":\"Question is the prompt evaluated by the decision model.\",\"type\":\"string\"},\"threshold\":{\"description\":\"Threshold is the confidence score that the decision model must exceed to pass.\",\"type\":\"number\"}},\"required\":[\"question\",\"threshold\"],\"type\":\"object\"},\"ConfigHttpRequestMethod\":{\"enum\":[\"GET\",\"POST\",\"PUT\",\"PATCH\",\"DELETE\"],\"enumShort\":{\"delete\":\"DELETE\",\"get\":\"GET\",\"patch\":\"PATCH\",\"post\":\"POST\",\"put\":\"PUT\"},\"type\":\"string\"},\"Config_Bare\":{\"properties\":{\"alwaysSetResult\":{\"description\":\"When set, the tool call's result is recorded immediately as this\\n fixed text instead of parking the call to wait for externally\\n supplied content. The tool_called event is still emitted. Useful for\\n tools whose dispatch is the intent (e.g. a frontend renders a\\n component from the call parameters) but whose LLM turn still needs\\n tool-result content.\",\"type\":\"string\"}},\"type\":\"object\"},\"Config_HTTP\":{\"properties\":{\"headers\":{\"additionalProperties\":{\"type\":\"string\"},\"type\":\"object\"},\"path\":{\"type\":\"string\"},\"query\":{\"type\":\"string\"},\"requestBodyContentType\":{\"type\":\"string\"},\"requestBodyTemplate\":{\"description\":\"These are only used when the request method is a POST, PUT, or PATCH\",\"type\":\"string\"},\"requestMethod\":{\"$ref\":\"ConfigHttpRequestMethod\"}},\"required\":[\"requestMethod\"],\"type\":\"object\"},\"Config_MCP\":{\"properties\":{\"annotations\":{\"$ref\":\"MCP_Annotations\",\"description\":\"Tool behavior annotations from the MCP server, captured during sync.\"}},\"type\":\"object\"},\"Config_OpenAPI\":{\"properties\":{\"method\":{\"type\":\"string\"},\"path\":{\"type\":\"string\"}},\"type\":\"object\"},\"MCP_Annotations\":{\"properties\":{\"destructiveHint\":{\"description\":\"If true, the tool may perform destructive updates to its environment.\\n Only meaningful when read_only_hint is false.\",\"type\":\"boolean\"},\"idempotentHint\":{\"description\":\"If true, calling the tool repeatedly with the same arguments has no\\n additional effect. Only meaningful when read_only_hint is false.\",\"type\":\"boolean\"},\"openWorldHint\":{\"description\":\"If true, the tool may interact with an \\\"open world\\\" of external\\n entities (e.g. web search); if false, its domain is closed.\",\"type\":\"boolean\"},\"readOnlyHint\":{\"description\":\"If true, the tool does not modify its environment.\",\"type\":\"boolean\"},\"title\":{\"description\":\"A human-readable title for the tool.\",\"type\":\"string\"}},\"type\":\"object\"},\"ToolActions\":{\"properties\":{\"autoMode\":{\"$ref\":\"AutoModeConfig\",\"description\":\"Optional automatic approval policy. Unset means no opinion.\"},\"parameterActions\":{\"description\":\"Pre-call actions, applied in order. A later action on the same parameter\\n path wins.\",\"items\":{\"$ref\":\"ToolParameterAction\"},\"type\":\"array\"},\"resultActions\":{\"description\":\"Post-call actions, applied in order. Result transforms chain: each one\\n reads the previous one's output. Not supported on bare tool sets.\",\"items\":{\"$ref\":\"ToolResultAction\"},\"type\":\"array\"},\"widgetArgumentExposure\":{\"$ref\":\"ToolWidgetArgumentExposure\",\"description\":\"Whether matching tools expose their final call arguments in widget\\n events. Unset means no opinion.\"}},\"type\":\"object\"},\"ToolParameterAction\":{\"discriminator\":{\"propertyName\":\"type\"},\"oneOf\":[{\"$ref\":\"ToolParameterAction_RemoveVariant\"},{\"$ref\":\"ToolParameterAction_SetVariant\"},{\"$ref\":\"ToolParameterAction_PinVariant\"}]},\"ToolParameterActionPinOnMissing\":{\"enum\":[\"ON_MISSING_FAIL\",\"ON_MISSING_SKIP\"],\"enumShort\":{\"fail\":\"ON_MISSING_FAIL\",\"skip\":\"ON_MISSING_SKIP\"},\"type\":\"string\"},\"ToolParameterAction_Pin\":{\"properties\":{\"onMissing\":{\"$ref\":\"ToolParameterActionPinOnMissing\",\"description\":\"Default: ON_MISSING_FAIL.\"},\"path\":{\"type\":\"string\"},\"pinnedParameter\":{\"description\":\"Key into the objective's pinned_parameters map. Need not equal the\\n last segment of `path` — this is how a pinned `orgId` reaches a\\n tool whose parameter is named `organizationId`.\",\"type\":\"string\"}},\"required\":[\"path\",\"pinnedParameter\",\"onMissing\"],\"type\":\"object\"},\"ToolParameterAction_PinVariant\":{\"properties\":{\"pin\":{\"$ref\":\"ToolParameterAction_Pin\"},\"type\":{\"const\":\"pin\"}},\"required\":[\"type\",\"pin\"],\"type\":\"object\"},\"ToolParameterAction_Remove\":{\"properties\":{\"path\":{\"type\":\"string\"}},\"required\":[\"path\"],\"type\":\"object\"},\"ToolParameterAction_RemoveVariant\":{\"properties\":{\"remove\":{\"$ref\":\"ToolParameterAction_Remove\"},\"type\":{\"const\":\"remove\"}},\"required\":[\"type\",\"remove\"],\"type\":\"object\"},\"ToolParameterAction_Set\":{\"properties\":{\"path\":{\"type\":\"string\"},\"valueTemplate\":{\"type\":\"string\"}},\"required\":[\"path\",\"valueTemplate\"],\"type\":\"object\"},\"ToolParameterAction_SetVariant\":{\"properties\":{\"set\":{\"$ref\":\"ToolParameterAction_Set\"},\"type\":{\"const\":\"set\"}},\"required\":[\"type\",\"set\"],\"type\":\"object\"},\"ToolResultAction\":{\"discriminator\":{\"propertyName\":\"type\"},\"oneOf\":[{\"$ref\":\"ToolResultAction_TransformVariant\"}]},\"ToolResultActionTransformOnError\":{\"enum\":[\"ON_ERROR_RAW_CONTENT\",\"ON_ERROR_FAIL\"],\"enumShort\":{\"fail\":\"ON_ERROR_FAIL\",\"raw-content\":\"ON_ERROR_RAW_CONTENT\"},\"type\":\"string\"},\"ToolResultAction_Transform\":{\"properties\":{\"contentTemplate\":{\"type\":\"string\"},\"expectJson\":{\"description\":\"Require the tool result to have text content that parses as JSON\\n before rendering. A non-JSON (or text-less) result is then an error\\n subject to `on_error` even if the template never reads\\n `result.json`. Off by default: `result.json` is simply absent for\\n non-JSON results, and text-less results skip the transform.\",\"type\":\"boolean\"},\"onError\":{\"$ref\":\"ToolResultActionTransformOnError\",\"description\":\"Default: ON_ERROR_RAW_CONTENT.\"}},\"required\":[\"contentTemplate\",\"onError\",\"expectJson\"],\"type\":\"object\"},\"ToolResultAction_TransformVariant\":{\"properties\":{\"transform\":{\"$ref\":\"ToolResultAction_Transform\"},\"type\":{\"const\":\"transform\"}},\"required\":[\"type\",\"transform\"],\"type\":\"object\"},\"ToolSpec\":{\"properties\":{\"actions\":{\"$ref\":\"ToolActions\",\"description\":\"The tool's own actions. They apply after every rule on the tool set, so\\n they win every conflict; the tool's result transforms run last. Only\\n allowed on tools managed through the API: a tool that comes from a sync\\n or from just-in-time loading is replaced when its upstream definition\\n changes, so use a rule with `tool.name == \\\"...\\\"` for those.\"},\"config\":{\"$ref\":\"ToolSpec_Config\",\"description\":\"Configuration for this specific tool. Its transport is derived from the\\n tool set adapter, while details such as endpoint and method are stored on\\n the tool itself.\\n\\n Required, and exactly one adapter must be set.\"},\"description\":{\"type\":\"string\"},\"llmToolName\":{\"description\":\"The name provided to the LLM, which may differ from the metadata.name on the tool.\\n LLMs have specific length and format requirements, and tool set sources may not comply\\n with them, so Cadenya does its best to format names into a usable format.\",\"type\":\"string\"},\"parameters\":{\"additionalProperties\":{},\"description\":\"The tool's JSON Schema, as handed to the LLM. Required, but may be the\\n empty object `{}` for a tool that takes no arguments. Requiring it rather\\n than defaulting it means a misspelled field name (`inputSchema`, say) is a\\n 400 instead of a silently parameterless tool.\",\"type\":\"object\"},\"requiresApproval\":{\"type\":\"boolean\"}},\"required\":[\"description\",\"requiresApproval\",\"parameters\",\"config\"],\"type\":\"object\"},\"ToolSpec_Config\":{\"discriminator\":{\"propertyName\":\"type\"},\"oneOf\":[{\"$ref\":\"ToolSpec_Config_Http\"},{\"$ref\":\"ToolSpec_Config_Mcp\"},{\"$ref\":\"ToolSpec_Config_Openapi\"},{\"$ref\":\"ToolSpec_Config_Bare\"}]},\"ToolSpec_Config_Bare\":{\"properties\":{\"bare\":{\"$ref\":\"Config_Bare\"},\"type\":{\"const\":\"bare\"}},\"required\":[\"type\",\"bare\"],\"type\":\"object\"},\"ToolSpec_Config_Http\":{\"properties\":{\"http\":{\"$ref\":\"Config_HTTP\"},\"type\":{\"const\":\"http\"}},\"required\":[\"type\",\"http\"],\"type\":\"object\"},\"ToolSpec_Config_Mcp\":{\"properties\":{\"mcp\":{\"$ref\":\"Config_MCP\"},\"type\":{\"const\":\"mcp\"}},\"required\":[\"type\",\"mcp\"],\"type\":\"object\"},\"ToolSpec_Config_Openapi\":{\"properties\":{\"openapi\":{\"$ref\":\"Config_OpenAPI\"},\"type\":{\"const\":\"openapi\"}},\"required\":[\"type\",\"openapi\"],\"type\":\"object\"},\"ToolWidgetArgumentExposure\":{\"properties\":{\"enabled\":{\"type\":\"boolean\"}},\"required\":[\"enabled\"],\"type\":\"object\"},\"UpdateResourceMetadata\":{\"properties\":{\"externalId\":{\"description\":\"External ID for the resource (e.g., a workflow ID from an external system)\",\"type\":\"string\"},\"labels\":{\"additionalProperties\":{\"type\":\"string\"},\"description\":\"Key-value pairs for categorization and filtering. Values are 0-63\\n alphanumeric characters with \\\"-\\\", \\\"_\\\", or \\\".\\\" allowed between; keys\\n follow the same shape and additionally accept an optional DNS-subdomain\\n prefix (e.g. \\\"cadenya.com/\\\") of at most 253 characters.\\n Examples: {\\\"environment\\\": \\\"production\\\", \\\"team\\\": \\\"platform\\\", \\\"version\\\": \\\"v2\\\"}\",\"type\":\"object\"},\"name\":{\"description\":\"Human-readable name for the resource (e.g., \\\"Customer Support Agent\\\", \\\"Email Tool\\\")\",\"type\":\"string\"}},\"required\":[\"name\"],\"type\":\"object\"}},\"properties\":{\"metadata\":{\"$ref\":\"UpdateResourceMetadata\"},\"spec\":{\"$ref\":\"ToolSpec\"},\"updateMask\":{\"type\":\"string\"}},\"type\":\"object\"}"

// ToolsUpdateConversion is the typed request produced from one urfave command.
type ToolsUpdateConversion struct {
	Params sdk.ToolUpdateParams
	Body   any
}

// ConvertToolsUpdate reads urfave values according to Redwood's operation IR.
func ConvertToolsUpdate(cmd *cli.Command, out *ToolsUpdateConversion) error {
	values := map[string]any{}
	if cmd.IsSet("workspace-id") {
		values["workspaceId"] = cmd.String("workspace-id")
	}
	_schema := parseBodySchema(bodySchemaToolsUpdate)
	_body := newBodyBuilder()
	_strict := cmd.Bool("strict")
	var _rawBody any
	if cmd.IsSet("file") {
		if err := _body.applyFile("file", cmd.String("file"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("metadata") {
		if err := _body.applyDoc("metadata", []string{"metadata"}, cmd.String("metadata"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("spec") {
		if err := _body.applyDoc("spec", []string{"spec"}, cmd.String("spec"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("config") {
		if err := _body.applyUnionFlag(unionSpec{Flag: "config", Path: []string{"spec", "config"}, Discriminator: "type", Required: false, Inferable: true, Arms: []unionArm{{Tag: "http", Keys: []string{"http"}, Init: []string{"http"}}, {Tag: "mcp", Keys: []string{"mcp"}, Init: []string{"mcp"}}, {Tag: "openapi", Keys: []string{"openapi"}, Init: []string{"openapi"}}, {Tag: "bare", Keys: []string{"bare"}, Init: []string{"bare"}}}}, cmd.String("config"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("http") {
		if err := _body.applyDoc("http", []string{"spec", "config", "http"}, cmd.String("http"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("mcp") {
		if err := _body.applyDoc("mcp", []string{"spec", "config", "mcp"}, cmd.String("mcp"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("mcp-annotations") {
		if err := _body.applyDoc("mcp-annotations", []string{"spec", "config", "mcp", "annotations"}, cmd.String("mcp-annotations"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("openapi") {
		if err := _body.applyDoc("openapi", []string{"spec", "config", "openapi"}, cmd.String("openapi"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("bare") {
		if err := _body.applyDoc("bare", []string{"spec", "config", "bare"}, cmd.String("bare"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions") {
		if err := _body.applyDoc("actions", []string{"spec", "actions"}, cmd.String("actions"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-widget-argument-exposure") {
		if err := _body.applyDoc("actions-widget-argument-exposure", []string{"spec", "actions", "widgetArgumentExposure"}, cmd.String("actions-widget-argument-exposure"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-auto-mode") {
		if err := _body.applyDoc("actions-auto-mode", []string{"spec", "actions", "autoMode"}, cmd.String("actions-auto-mode"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("name") {
		_v, err := stringArg("name", cmd.String("name"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("name", []string{"metadata", "name"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("external-id") {
		_v, err := stringArg("external-id", cmd.String("external-id"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("external-id", []string{"metadata", "externalId"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("label") {
		if err := _body.applyEntries("label", []string{"metadata", "labels"}, cmd.StringSlice("label"), scalarString, nil); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("description") {
		_v, err := stringArg("description", cmd.String("description"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("description", []string{"spec", "description"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("requires-approval") {
		if err := _body.set("requires-approval", []string{"spec", "requiresApproval"}, cmd.Bool("requires-approval")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("parameter") {
		if err := _body.applyEntryDocs("parameter", []string{"spec", "parameters"}, cmd.StringSlice("parameter")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("http-request-method") {
		_v, err := enumSpec{Values: []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, Short: []string{"get", "post", "put", "patch", "delete"}}.parse("http-request-method", cmd.String("http-request-method"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("http-request-method", []string{"spec", "config", "http", "requestMethod"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("http-path") {
		_v, err := stringArg("http-path", cmd.String("http-path"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("http-path", []string{"spec", "config", "http", "path"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("http-query") {
		_v, err := stringArg("http-query", cmd.String("http-query"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("http-query", []string{"spec", "config", "http", "query"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("http-header") {
		if err := _body.applyEntries("http-header", []string{"spec", "config", "http", "headers"}, cmd.StringSlice("http-header"), scalarString, nil); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("http-request-body-template") {
		_v, err := stringArg("http-request-body-template", cmd.String("http-request-body-template"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("http-request-body-template", []string{"spec", "config", "http", "requestBodyTemplate"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("http-request-body-content-type") {
		_v, err := stringArg("http-request-body-content-type", cmd.String("http-request-body-content-type"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("http-request-body-content-type", []string{"spec", "config", "http", "requestBodyContentType"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("mcp-annotations-title") {
		_v, err := stringArg("mcp-annotations-title", cmd.String("mcp-annotations-title"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("mcp-annotations-title", []string{"spec", "config", "mcp", "annotations", "title"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("mcp-annotations-read-only-hint") {
		if err := _body.set("mcp-annotations-read-only-hint", []string{"spec", "config", "mcp", "annotations", "readOnlyHint"}, cmd.Bool("mcp-annotations-read-only-hint")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("mcp-annotations-destructive-hint") {
		if err := _body.set("mcp-annotations-destructive-hint", []string{"spec", "config", "mcp", "annotations", "destructiveHint"}, cmd.Bool("mcp-annotations-destructive-hint")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("mcp-annotations-idempotent-hint") {
		if err := _body.set("mcp-annotations-idempotent-hint", []string{"spec", "config", "mcp", "annotations", "idempotentHint"}, cmd.Bool("mcp-annotations-idempotent-hint")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("mcp-annotations-open-world-hint") {
		if err := _body.set("mcp-annotations-open-world-hint", []string{"spec", "config", "mcp", "annotations", "openWorldHint"}, cmd.Bool("mcp-annotations-open-world-hint")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("openapi-path") {
		_v, err := stringArg("openapi-path", cmd.String("openapi-path"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("openapi-path", []string{"spec", "config", "openapi", "path"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("openapi-method") {
		_v, err := stringArg("openapi-method", cmd.String("openapi-method"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("openapi-method", []string{"spec", "config", "openapi", "method"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("bare-always-set-result") {
		_v, err := stringArg("bare-always-set-result", cmd.String("bare-always-set-result"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("bare-always-set-result", []string{"spec", "config", "bare", "alwaysSetResult"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("llm-tool-name") {
		_v, err := stringArg("llm-tool-name", cmd.String("llm-tool-name"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("llm-tool-name", []string{"spec", "llmToolName"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-parameter-action") {
		if err := _body.applyDocItems("actions-parameter-action", []string{"spec", "actions", "parameterActions"}, cmd.StringSlice("actions-parameter-action")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-result-action") {
		if err := _body.applyDocItems("actions-result-action", []string{"spec", "actions", "resultActions"}, cmd.StringSlice("actions-result-action")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-widget-argument-exposure-enabled") {
		if err := _body.set("actions-widget-argument-exposure-enabled", []string{"spec", "actions", "widgetArgumentExposure", "enabled"}, cmd.Bool("actions-widget-argument-exposure-enabled")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-auto-mode-question") {
		if err := _body.applyShorthandItems("actions-auto-mode-question", []string{"spec", "actions", "autoMode", "questions"}, cmd.StringSlice("actions-auto-mode-question"), shorthandSpec{Fields: []shorthandField{{Wire: "question", Key: "question", Kind: scalarString, Enum: nil, Required: true}, {Wire: "threshold", Key: "threshold", Kind: scalarFloat, Enum: nil, Required: true}}}); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-auto-mode-passing-mode") {
		_v, err := enumSpec{Values: []string{"PASSING_MODE_ALL", "PASSING_MODE_NONE", "PASSING_MODE_MOST"}, Short: []string{"all", "none", "most"}}.parse("actions-auto-mode-passing-mode", cmd.String("actions-auto-mode-passing-mode"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("actions-auto-mode-passing-mode", []string{"spec", "actions", "autoMode", "passingMode"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-auto-mode-failure-message") {
		_v, err := stringArg("actions-auto-mode-failure-message", cmd.String("actions-auto-mode-failure-message"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("actions-auto-mode-failure-message", []string{"spec", "actions", "autoMode", "failureMessage"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("update-mask") {
		_v, err := stringArg("update-mask", cmd.String("update-mask"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("update-mask", []string{"updateMask"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if err := _body.resolveUnion(unionSpec{Flag: "config", Path: []string{"spec", "config"}, Discriminator: "type", Required: false, Inferable: true, Arms: []unionArm{{Tag: "http", Keys: []string{"http"}, Init: []string{"http"}}, {Tag: "mcp", Keys: []string{"mcp"}, Init: []string{"mcp"}}, {Tag: "openapi", Keys: []string{"openapi"}, Init: []string{"openapi"}}, {Tag: "bare", Keys: []string{"bare"}, Init: []string{"bare"}}}}); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	// A partial update names the paths it changes; a mask supplied
	// by flag or document wins.
	if _, _given := _body.lookup([]string{"updateMask"}); !_given {
		if _mask := _body.updateMask("updateMask"); _mask != "" {
			_ = _body.set("update-mask", []string{"updateMask"}, _mask)
		}
	}
	if err := _body.finish(_schema, map[string]string{"metadata": "--metadata", "metadata.name": "--name", "metadata.externalId": "--external-id", "metadata.labels": "--label", "spec": "--spec", "spec.description": "--description", "spec.requiresApproval": "--requires-approval", "spec.parameters": "--parameter", "spec.config": "--config", "spec.config.http": "--http", "spec.config.http.requestMethod": "--http-request-method", "spec.config.http.path": "--http-path", "spec.config.http.query": "--http-query", "spec.config.http.headers": "--http-header", "spec.config.http.requestBodyTemplate": "--http-request-body-template", "spec.config.http.requestBodyContentType": "--http-request-body-content-type", "spec.config.mcp": "--mcp", "spec.config.mcp.annotations": "--mcp-annotations", "spec.config.mcp.annotations.title": "--mcp-annotations-title", "spec.config.mcp.annotations.readOnlyHint": "--mcp-annotations-read-only-hint", "spec.config.mcp.annotations.destructiveHint": "--mcp-annotations-destructive-hint", "spec.config.mcp.annotations.idempotentHint": "--mcp-annotations-idempotent-hint", "spec.config.mcp.annotations.openWorldHint": "--mcp-annotations-open-world-hint", "spec.config.openapi": "--openapi", "spec.config.openapi.path": "--openapi-path", "spec.config.openapi.method": "--openapi-method", "spec.config.bare": "--bare", "spec.config.bare.alwaysSetResult": "--bare-always-set-result", "spec.llmToolName": "--llm-tool-name", "spec.actions": "--actions", "spec.actions.parameterActions": "--actions-parameter-action", "spec.actions.resultActions": "--actions-result-action", "spec.actions.widgetArgumentExposure": "--actions-widget-argument-exposure", "spec.actions.widgetArgumentExposure.enabled": "--actions-widget-argument-exposure-enabled", "spec.actions.autoMode": "--actions-auto-mode", "spec.actions.autoMode.questions": "--actions-auto-mode-question", "spec.actions.autoMode.passingMode": "--actions-auto-mode-passing-mode", "spec.actions.autoMode.failureMessage": "--actions-auto-mode-failure-message", "updateMask": "--update-mask"}, _strict); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	_ = _rawBody
	for _k, _v := range _body.body {
		values[_k] = _v
	}
	if err := decodeParams(values, &out.Params); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	out.Body = _body.body
	return nil
}
