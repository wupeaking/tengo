package stdlib

import (
	"encoding/hex"

	"github.com/d5/tengo/v2/common"
)

var hexModule = map[string]common.Object{
	"encode": &common.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &common.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
