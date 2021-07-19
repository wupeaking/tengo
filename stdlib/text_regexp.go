package stdlib

import (
	"regexp"

	"github.com/d5/tengo/v2/common"
)

func makeTextRegexp(re *regexp.Regexp) *common.ImmutableMap {
	return &common.ImmutableMap{
		Value: map[string]common.Object{
			// match(text) => bool
			"match": &common.UserFunction{
				Value: func(args ...common.Object) (
					ret common.Object,
					err error,
				) {
					if len(args) != 1 {
						err = common.ErrWrongNumArguments
						return
					}

					s1, ok := common.ToString(args[0])
					if !ok {
						err = common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if re.MatchString(s1) {
						ret = common.TrueValue
					} else {
						ret = common.FalseValue
					}

					return
				},
			},

			// find(text) 			=> array(array({text:,begin:,end:}))/undefined
			// find(text, maxCount) => array(array({text:,begin:,end:}))/undefined
			"find": &common.UserFunction{
				Value: func(args ...common.Object) (
					ret common.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = common.ErrWrongNumArguments
						return
					}

					s1, ok := common.ToString(args[0])
					if !ok {
						err = common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if numArgs == 1 {
						m := re.FindStringSubmatchIndex(s1)
						if m == nil {
							ret = common.UndefinedValue
							return
						}

						arr := &common.Array{}
						for i := 0; i < len(m); i += 2 {
							arr.Value = append(arr.Value,
								&common.ImmutableMap{
									Value: map[string]common.Object{
										"text": &common.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &common.Int{
											Value: int64(m[i]),
										},
										"end": &common.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						ret = &common.Array{Value: []common.Object{arr}}

						return
					}

					i2, ok := common.ToInt(args[1])
					if !ok {
						err = common.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}
					m := re.FindAllStringSubmatchIndex(s1, i2)
					if m == nil {
						ret = common.UndefinedValue
						return
					}

					arr := &common.Array{}
					for _, m := range m {
						subMatch := &common.Array{}
						for i := 0; i < len(m); i += 2 {
							subMatch.Value = append(subMatch.Value,
								&common.ImmutableMap{
									Value: map[string]common.Object{
										"text": &common.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &common.Int{
											Value: int64(m[i]),
										},
										"end": &common.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						arr.Value = append(arr.Value, subMatch)
					}

					ret = arr

					return
				},
			},

			// replace(src, repl) => string
			"replace": &common.UserFunction{
				Value: func(args ...common.Object) (
					ret common.Object,
					err error,
				) {
					if len(args) != 2 {
						err = common.ErrWrongNumArguments
						return
					}

					s1, ok := common.ToString(args[0])
					if !ok {
						err = common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					s2, ok := common.ToString(args[1])
					if !ok {
						err = common.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "string(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}

					s, ok := doTextRegexpReplace(re, s1, s2)
					if !ok {
						return nil, common.ErrStringLimit
					}

					ret = &common.String{Value: s}

					return
				},
			},

			// split(text) 			 => array(string)
			// split(text, maxCount) => array(string)
			"split": &common.UserFunction{
				Value: func(args ...common.Object) (
					ret common.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = common.ErrWrongNumArguments
						return
					}

					s1, ok := common.ToString(args[0])
					if !ok {
						err = common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					var i2 = -1
					if numArgs > 1 {
						i2, ok = common.ToInt(args[1])
						if !ok {
							err = common.ErrInvalidArgumentType{
								Name:     "second",
								Expected: "int(compatible)",
								Found:    args[1].TypeName(),
							}
							return
						}
					}

					arr := &common.Array{}
					for _, s := range re.Split(s1, i2) {
						arr.Value = append(arr.Value,
							&common.String{Value: s})
					}

					ret = arr

					return
				},
			},
		},
	}
}

// Size-limit checking implementation of regexp.ReplaceAllString.
func doTextRegexpReplace(re *regexp.Regexp, src, repl string) (string, bool) {
	idx := 0
	out := ""
	for _, m := range re.FindAllStringSubmatchIndex(src, -1) {
		var exp []byte
		exp = re.ExpandString(exp, repl, src, m)
		if len(out)+m[0]-idx+len(exp) > common.MaxStringLen {
			return "", false
		}
		out += src[idx:m[0]] + string(exp)
		idx = m[1]
	}
	if idx < len(src) {
		if len(out)+len(src)-idx > common.MaxStringLen {
			return "", false
		}
		out += src[idx:]
	}
	return out, true
}
