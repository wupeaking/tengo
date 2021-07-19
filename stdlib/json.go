package stdlib

import (
	"bytes"
	gojson "encoding/json"

	"github.com/d5/tengo/v2/common"
	"github.com/d5/tengo/v2/stdlib/json"
)

var jsonModule = map[string]common.Object{
	"decode": &common.UserFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &common.UserFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &common.UserFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &common.UserFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		return nil, common.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *common.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &common.Error{
				Value: &common.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *common.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &common.Error{
				Value: &common.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	default:
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonEncode(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		return nil, common.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &common.Error{Value: &common.String{Value: err.Error()}}, nil
	}

	return &common.Bytes{Value: b}, nil
}

func jsonIndent(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 3 {
		return nil, common.ErrWrongNumArguments
	}

	prefix, ok := common.ToString(args[1])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := common.ToString(args[2])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *common.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &common.Error{
				Value: &common.String{Value: err.Error()},
			}, nil
		}
		return &common.Bytes{Value: dst.Bytes()}, nil
	case *common.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &common.Error{
				Value: &common.String{Value: err.Error()},
			}, nil
		}
		return &common.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		return nil, common.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *common.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &common.Bytes{Value: dst.Bytes()}, nil
	case *common.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &common.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
