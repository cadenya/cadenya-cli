package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

const bodySchemaTrustedEndpointsVerify = "{\"$defs\":{},\"properties\":{\"url\":{\"type\":\"string\"}},\"required\":[\"url\"],\"type\":\"object\"}"

// TrustedEndpointsVerifyConversion is the typed request produced from one urfave command.
type TrustedEndpointsVerifyConversion struct {
	Params sdk.TrustedEndpointVerifyParams
	Body   any
}

// ConvertTrustedEndpointsVerify reads urfave values according to Redwood's operation IR.
func ConvertTrustedEndpointsVerify(cmd *cli.Command, out *TrustedEndpointsVerifyConversion) error {
	values := map[string]any{}
	_schema := parseBodySchema(bodySchemaTrustedEndpointsVerify)
	_body := newBodyBuilder()
	_strict := cmd.Bool("strict")
	var _rawBody any
	if cmd.IsSet("file") {
		if err := _body.applyFile("file", cmd.String("file"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("url") {
		_v, err := stringArg("url", cmd.String("url"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("url", []string{"url"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if err := _body.finish(_schema, map[string]string{"url": "--url"}, _strict); err != nil {
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
