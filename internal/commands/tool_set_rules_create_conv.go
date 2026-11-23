package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

const bodySchemaToolSetRulesCreate = "{\"$defs\":{\"AutoModeConfig\":{\"properties\":{\"failureMessage\":{\"description\":\"Optional guidance returned to the agent when the decision thresholds fail.\\n When empty, a default denial message is used. Client errors never use this message.\",\"type\":\"string\"},\"passingMode\":{\"$ref\":\"AutoModeConfigPassingMode\"},\"questions\":{\"items\":{\"$ref\":\"AutoModeQuestion\"},\"type\":\"array\"}},\"required\":[\"questions\",\"passingMode\"],\"type\":\"object\"},\"AutoModeConfigPassingMode\":{\"enum\":[\"PASSING_MODE_ALL\",\"PASSING_MODE_NONE\",\"PASSING_MODE_MOST\"],\"enumShort\":{\"all\":\"PASSING_MODE_ALL\",\"most\":\"PASSING_MODE_MOST\",\"none\":\"PASSING_MODE_NONE\"},\"type\":\"string\"},\"AutoModeQuestion\":{\"properties\":{\"question\":{\"description\":\"Question is the prompt evaluated by the decision model.\",\"type\":\"string\"},\"threshold\":{\"description\":\"Threshold is the confidence score that the decision model must exceed to pass.\",\"type\":\"number\"}},\"required\":[\"question\",\"threshold\"],\"type\":\"object\"},\"ToolActions\":{\"properties\":{\"autoMode\":{\"$ref\":\"AutoModeConfig\",\"description\":\"Optional automatic approval policy. Unset means no opinion.\"},\"parameterActions\":{\"description\":\"Pre-call actions, applied in order. A later action on the same parameter\\n path wins.\",\"items\":{\"$ref\":\"ToolParameterAction\"},\"type\":\"array\"},\"resultActions\":{\"description\":\"Post-call actions, applied in order. Result transforms chain: each one\\n reads the previous one's output. Not supported on bare tool sets.\",\"items\":{\"$ref\":\"ToolResultAction\"},\"type\":\"array\"},\"widgetArgumentExposure\":{\"$ref\":\"ToolWidgetArgumentExposure\",\"description\":\"Whether matching tools expose their final call arguments in widget\\n events. Unset means no opinion.\"}},\"type\":\"object\"},\"ToolParameterAction\":{\"discriminator\":{\"propertyName\":\"type\"},\"oneOf\":[{\"$ref\":\"ToolParameterAction_RemoveVariant\"},{\"$ref\":\"ToolParameterAction_SetVariant\"},{\"$ref\":\"ToolParameterAction_PinVariant\"}]},\"ToolParameterActionPinOnMissing\":{\"enum\":[\"ON_MISSING_FAIL\",\"ON_MISSING_SKIP\"],\"enumShort\":{\"fail\":\"ON_MISSING_FAIL\",\"skip\":\"ON_MISSING_SKIP\"},\"type\":\"string\"},\"ToolParameterAction_Pin\":{\"properties\":{\"onMissing\":{\"$ref\":\"ToolParameterActionPinOnMissing\",\"description\":\"Default: ON_MISSING_FAIL.\"},\"path\":{\"type\":\"string\"},\"pinnedParameter\":{\"description\":\"Key into the objective's pinned_parameters map. Need not equal the\\n last segment of `path` — this is how a pinned `orgId` reaches a\\n tool whose parameter is named `organizationId`.\",\"type\":\"string\"}},\"required\":[\"path\",\"pinnedParameter\",\"onMissing\"],\"type\":\"object\"},\"ToolParameterAction_PinVariant\":{\"properties\":{\"pin\":{\"$ref\":\"ToolParameterAction_Pin\"},\"type\":{\"const\":\"pin\"}},\"required\":[\"type\",\"pin\"],\"type\":\"object\"},\"ToolParameterAction_Remove\":{\"properties\":{\"path\":{\"type\":\"string\"}},\"required\":[\"path\"],\"type\":\"object\"},\"ToolParameterAction_RemoveVariant\":{\"properties\":{\"remove\":{\"$ref\":\"ToolParameterAction_Remove\"},\"type\":{\"const\":\"remove\"}},\"required\":[\"type\",\"remove\"],\"type\":\"object\"},\"ToolParameterAction_Set\":{\"properties\":{\"path\":{\"type\":\"string\"},\"valueTemplate\":{\"type\":\"string\"}},\"required\":[\"path\",\"valueTemplate\"],\"type\":\"object\"},\"ToolParameterAction_SetVariant\":{\"properties\":{\"set\":{\"$ref\":\"ToolParameterAction_Set\"},\"type\":{\"const\":\"set\"}},\"required\":[\"type\",\"set\"],\"type\":\"object\"},\"ToolResultAction\":{\"discriminator\":{\"propertyName\":\"type\"},\"oneOf\":[{\"$ref\":\"ToolResultAction_TransformVariant\"}]},\"ToolResultActionTransformOnError\":{\"enum\":[\"ON_ERROR_RAW_CONTENT\",\"ON_ERROR_FAIL\"],\"enumShort\":{\"fail\":\"ON_ERROR_FAIL\",\"raw-content\":\"ON_ERROR_RAW_CONTENT\"},\"type\":\"string\"},\"ToolResultAction_Transform\":{\"properties\":{\"contentTemplate\":{\"type\":\"string\"},\"expectJson\":{\"description\":\"Require the tool result to have text content that parses as JSON\\n before rendering. A non-JSON (or text-less) result is then an error\\n subject to `on_error` even if the template never reads\\n `result.json`. Off by default: `result.json` is simply absent for\\n non-JSON results, and text-less results skip the transform.\",\"type\":\"boolean\"},\"onError\":{\"$ref\":\"ToolResultActionTransformOnError\",\"description\":\"Default: ON_ERROR_RAW_CONTENT.\"}},\"required\":[\"contentTemplate\",\"onError\",\"expectJson\"],\"type\":\"object\"},\"ToolResultAction_TransformVariant\":{\"properties\":{\"transform\":{\"$ref\":\"ToolResultAction_Transform\"},\"type\":{\"const\":\"transform\"}},\"required\":[\"type\",\"transform\"],\"type\":\"object\"},\"ToolWidgetArgumentExposure\":{\"properties\":{\"enabled\":{\"type\":\"boolean\"}},\"required\":[\"enabled\"],\"type\":\"object\"}},\"properties\":{\"actions\":{\"$ref\":\"ToolActions\"},\"description\":{\"type\":\"string\"},\"disabled\":{\"type\":\"boolean\"},\"name\":{\"type\":\"string\"},\"position\":{\"type\":\"integer\"},\"when\":{\"type\":\"string\"}},\"required\":[\"name\"],\"type\":\"object\"}"

// ToolSetRulesCreateConversion is the typed request produced from one urfave command.
type ToolSetRulesCreateConversion struct {
	Params sdk.ToolSetRuleCreateParams
	Body   any
}

// ConvertToolSetRulesCreate reads urfave values according to Redwood's operation IR.
func ConvertToolSetRulesCreate(cmd *cli.Command, out *ToolSetRulesCreateConversion) error {
	values := map[string]any{}
	if cmd.IsSet("workspace-id") {
		values["workspaceId"] = cmd.String("workspace-id")
	}
	_schema := parseBodySchema(bodySchemaToolSetRulesCreate)
	_body := newBodyBuilder()
	_strict := cmd.Bool("strict")
	var _rawBody any
	if cmd.IsSet("file") {
		if err := _body.applyFile("file", cmd.String("file"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions") {
		if err := _body.applyDoc("actions", []string{"actions"}, cmd.String("actions"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-widget-argument-exposure") {
		if err := _body.applyDoc("actions-widget-argument-exposure", []string{"actions", "widgetArgumentExposure"}, cmd.String("actions-widget-argument-exposure"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-auto-mode") {
		if err := _body.applyDoc("actions-auto-mode", []string{"actions", "autoMode"}, cmd.String("actions-auto-mode"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("name") {
		_v, err := stringArg("name", cmd.String("name"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("name", []string{"name"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("description") {
		_v, err := stringArg("description", cmd.String("description"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("description", []string{"description"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("when") {
		_v, err := stringArg("when", cmd.String("when"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("when", []string{"when"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-parameter-action") {
		if err := _body.applyDocItems("actions-parameter-action", []string{"actions", "parameterActions"}, cmd.StringSlice("actions-parameter-action")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-result-action") {
		if err := _body.applyDocItems("actions-result-action", []string{"actions", "resultActions"}, cmd.StringSlice("actions-result-action")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-widget-argument-exposure-enabled") {
		if err := _body.set("actions-widget-argument-exposure-enabled", []string{"actions", "widgetArgumentExposure", "enabled"}, cmd.Bool("actions-widget-argument-exposure-enabled")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-auto-mode-question") {
		if err := _body.applyShorthandItems("actions-auto-mode-question", []string{"actions", "autoMode", "questions"}, cmd.StringSlice("actions-auto-mode-question"), shorthandSpec{Fields: []shorthandField{{Wire: "question", Key: "question", Kind: scalarString, Enum: nil, Required: true}, {Wire: "threshold", Key: "threshold", Kind: scalarFloat, Enum: nil, Required: true}}}); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-auto-mode-passing-mode") {
		_v, err := enumSpec{Values: []string{"PASSING_MODE_ALL", "PASSING_MODE_NONE", "PASSING_MODE_MOST"}, Short: []string{"all", "none", "most"}}.parse("actions-auto-mode-passing-mode", cmd.String("actions-auto-mode-passing-mode"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("actions-auto-mode-passing-mode", []string{"actions", "autoMode", "passingMode"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("actions-auto-mode-failure-message") {
		_v, err := stringArg("actions-auto-mode-failure-message", cmd.String("actions-auto-mode-failure-message"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("actions-auto-mode-failure-message", []string{"actions", "autoMode", "failureMessage"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("disabled") {
		if err := _body.set("disabled", []string{"disabled"}, cmd.Bool("disabled")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("position") {
		if err := _body.set("position", []string{"position"}, cmd.Int32("position")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if err := _body.finish(_schema, map[string]string{"name": "--name", "description": "--description", "when": "--when", "actions": "--actions", "actions.parameterActions": "--actions-parameter-action", "actions.resultActions": "--actions-result-action", "actions.widgetArgumentExposure": "--actions-widget-argument-exposure", "actions.widgetArgumentExposure.enabled": "--actions-widget-argument-exposure-enabled", "actions.autoMode": "--actions-auto-mode", "actions.autoMode.questions": "--actions-auto-mode-question", "actions.autoMode.passingMode": "--actions-auto-mode-passing-mode", "actions.autoMode.failureMessage": "--actions-auto-mode-failure-message", "disabled": "--disabled", "position": "--position"}, _strict); err != nil {
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
