package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

const bodySchemaTrustedEndpointsCreate = "{\"$defs\":{\"CreateAccountResourceMetadata\":{\"properties\":{\"externalId\":{\"description\":\"External ID for the resource (e.g., a workflow ID from an external system)\",\"type\":\"string\"},\"labels\":{\"additionalProperties\":{\"type\":\"string\"},\"description\":\"Key-value pairs for categorization and filtering. Values are 0-63\\n alphanumeric characters with \\\"-\\\", \\\"_\\\", or \\\".\\\" allowed between; keys\\n follow the same shape and additionally accept an optional DNS-subdomain\\n prefix (e.g. \\\"cadenya.com/\\\") of at most 253 characters.\\n Examples: {\\\"environment\\\": \\\"production\\\", \\\"team\\\": \\\"platform\\\", \\\"version\\\": \\\"v2\\\"}\",\"type\":\"object\"},\"name\":{\"description\":\"Human-readable name for the resource (e.g., \\\"Production API Key\\\", \\\"Staging Workspace\\\")\",\"type\":\"string\"}},\"required\":[\"name\"],\"type\":\"object\"},\"TrustedEndpointSpec\":{\"properties\":{\"host\":{\"description\":\"The trusted host, with an optional port (default 443). Either an exact\\n host such as \\\"api.stripe.com\\\", or a wildcard covering exactly one label,\\n such as \\\"*.acme.com\\\", which matches \\\"t1.acme.com\\\" but not \\\"acme.com\\\" or\\n \\\"a.t1.acme.com\\\". A scheme, path or credentials are not allowed. The host\\n must be fully qualified and end in a known top-level domain (.com, .io),\\n and neither it nor a wildcard's domain may be a public suffix such as\\n \\\"com\\\", \\\"co.uk\\\" or \\\"vercel.app\\\".\\n The host is normalized (lowercased, converted to ASCII) when saved.\",\"type\":\"string\"}},\"required\":[\"host\"],\"type\":\"object\"}},\"properties\":{\"metadata\":{\"$ref\":\"CreateAccountResourceMetadata\"},\"spec\":{\"$ref\":\"TrustedEndpointSpec\"}},\"required\":[\"metadata\",\"spec\"],\"type\":\"object\"}"

// TrustedEndpointsCreateConversion is the typed request produced from one urfave command.
type TrustedEndpointsCreateConversion struct {
	Params sdk.TrustedEndpointCreateParams
	Body   any
}

// ConvertTrustedEndpointsCreate reads urfave values according to Redwood's operation IR.
func ConvertTrustedEndpointsCreate(cmd *cli.Command, out *TrustedEndpointsCreateConversion) error {
	values := map[string]any{}
	_schema := parseBodySchema(bodySchemaTrustedEndpointsCreate)
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
	if cmd.IsSet("host") {
		_v, err := stringArg("host", cmd.String("host"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("host", []string{"spec", "host"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if err := _body.finish(_schema, map[string]string{"metadata": "--metadata", "metadata.name": "--name", "metadata.externalId": "--external-id", "metadata.labels": "--label", "spec": "--spec", "spec.host": "--host"}, _strict); err != nil {
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
