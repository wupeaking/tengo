package stdlib

import (
	"math/rand"

	"github.com/d5/tengo/v2/common"
)

var randModule = map[string]common.Object{
	"int": &common.UserFunction{
		Name:  "int",
		Value: FuncARI64(rand.Int63),
	},
	"float": &common.UserFunction{
		Name:  "float",
		Value: FuncARF(rand.Float64),
	},
	"intn": &common.UserFunction{
		Name:  "intn",
		Value: FuncAI64RI64(rand.Int63n),
	},
	"exp_float": &common.UserFunction{
		Name:  "exp_float",
		Value: FuncARF(rand.ExpFloat64),
	},
	"norm_float": &common.UserFunction{
		Name:  "norm_float",
		Value: FuncARF(rand.NormFloat64),
	},
	"perm": &common.UserFunction{
		Name:  "perm",
		Value: FuncAIRIs(rand.Perm),
	},
	"seed": &common.UserFunction{
		Name:  "seed",
		Value: FuncAI64R(rand.Seed),
	},
	"read": &common.UserFunction{
		Name: "read",
		Value: func(args ...common.Object) (ret common.Object, err error) {
			if len(args) != 1 {
				return nil, common.ErrWrongNumArguments
			}
			y1, ok := args[0].(*common.Bytes)
			if !ok {
				return nil, common.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "bytes",
					Found:    args[0].TypeName(),
				}
			}
			res, err := rand.Read(y1.Value)
			if err != nil {
				ret = wrapError(err)
				return
			}
			return &common.Int{Value: int64(res)}, nil
		},
	},
	"rand": &common.UserFunction{
		Name: "rand",
		Value: func(args ...common.Object) (common.Object, error) {
			if len(args) != 1 {
				return nil, common.ErrWrongNumArguments
			}
			i1, ok := common.ToInt64(args[0])
			if !ok {
				return nil, common.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "int(compatible)",
					Found:    args[0].TypeName(),
				}
			}
			src := rand.NewSource(i1)
			return randRand(rand.New(src)), nil
		},
	},
}

func randRand(r *rand.Rand) *common.ImmutableMap {
	return &common.ImmutableMap{
		Value: map[string]common.Object{
			"int": &common.UserFunction{
				Name:  "int",
				Value: FuncARI64(r.Int63),
			},
			"float": &common.UserFunction{
				Name:  "float",
				Value: FuncARF(r.Float64),
			},
			"intn": &common.UserFunction{
				Name:  "intn",
				Value: FuncAI64RI64(r.Int63n),
			},
			"exp_float": &common.UserFunction{
				Name:  "exp_float",
				Value: FuncARF(r.ExpFloat64),
			},
			"norm_float": &common.UserFunction{
				Name:  "norm_float",
				Value: FuncARF(r.NormFloat64),
			},
			"perm": &common.UserFunction{
				Name:  "perm",
				Value: FuncAIRIs(r.Perm),
			},
			"seed": &common.UserFunction{
				Name:  "seed",
				Value: FuncAI64R(r.Seed),
			},
			"read": &common.UserFunction{
				Name: "read",
				Value: func(args ...common.Object) (
					ret common.Object,
					err error,
				) {
					if len(args) != 1 {
						return nil, common.ErrWrongNumArguments
					}
					y1, ok := args[0].(*common.Bytes)
					if !ok {
						return nil, common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "bytes",
							Found:    args[0].TypeName(),
						}
					}
					res, err := r.Read(y1.Value)
					if err != nil {
						ret = wrapError(err)
						return
					}
					return &common.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
