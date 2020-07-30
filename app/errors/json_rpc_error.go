package errors

import (
	"stash.lamoda.ru/gotools/rpc/proto/jsonrpc2"
)

type jsonRpsErrorParamsExtractor interface {
	ExtractJsonRpsErrorParams() (int, string)
}

func newJsonRpcError(code int, msg string) *jsonrpc2.Error {
	return jsonrpc2.NewError(code, msg, nil)
}

func WrapToJsonRpcError(err error) *jsonrpc2.Error {
	if baseErr, ok := err.(jsonRpsErrorParamsExtractor); ok {
		return newJsonRpcError(baseErr.ExtractJsonRpsErrorParams())
	} else {
		return jsonrpc2.NewError(jsonrpc2.E_SERVER, err.Error(), nil)
	}
}
