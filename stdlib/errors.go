package stdlib

import "github.com/d5/tengo/v2/common"

func wrapError(err error) common.Object {
	if err == nil {
		return common.TrueValue
	}
	return &common.Error{Value: &common.String{Value: err.Error()}}
}
