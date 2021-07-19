package stdlib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/d5/tengo/v2/common"
)

var textModule = map[string]common.Object{
	"re_match": &common.UserFunction{
		Name:  "re_match",
		Value: textREMatch,
	}, // re_match(pattern, text) => bool/error
	"re_find": &common.UserFunction{
		Name:  "re_find",
		Value: textREFind,
	}, // re_find(pattern, text, count) => [[{text:,begin:,end:}]]/undefined
	"re_replace": &common.UserFunction{
		Name:  "re_replace",
		Value: textREReplace,
	}, // re_replace(pattern, text, repl) => string/error
	"re_split": &common.UserFunction{
		Name:  "re_split",
		Value: textRESplit,
	}, // re_split(pattern, text, count) => [string]/error
	"re_compile": &common.UserFunction{
		Name:  "re_compile",
		Value: textRECompile,
	}, // re_compile(pattern) => Regexp/error
	"compare": &common.UserFunction{
		Name:  "compare",
		Value: FuncASSRI(strings.Compare),
	}, // compare(a, b) => int
	"contains": &common.UserFunction{
		Name:  "contains",
		Value: FuncASSRB(strings.Contains),
	}, // contains(s, substr) => bool
	"contains_any": &common.UserFunction{
		Name:  "contains_any",
		Value: FuncASSRB(strings.ContainsAny),
	}, // contains_any(s, chars) => bool
	"count": &common.UserFunction{
		Name:  "count",
		Value: FuncASSRI(strings.Count),
	}, // count(s, substr) => int
	"equal_fold": &common.UserFunction{
		Name:  "equal_fold",
		Value: FuncASSRB(strings.EqualFold),
	}, // "equal_fold(s, t) => bool
	"fields": &common.UserFunction{
		Name:  "fields",
		Value: FuncASRSs(strings.Fields),
	}, // fields(s) => [string]
	"has_prefix": &common.UserFunction{
		Name:  "has_prefix",
		Value: FuncASSRB(strings.HasPrefix),
	}, // has_prefix(s, prefix) => bool
	"has_suffix": &common.UserFunction{
		Name:  "has_suffix",
		Value: FuncASSRB(strings.HasSuffix),
	}, // has_suffix(s, suffix) => bool
	"index": &common.UserFunction{
		Name:  "index",
		Value: FuncASSRI(strings.Index),
	}, // index(s, substr) => int
	"index_any": &common.UserFunction{
		Name:  "index_any",
		Value: FuncASSRI(strings.IndexAny),
	}, // index_any(s, chars) => int
	"join": &common.UserFunction{
		Name:  "join",
		Value: textJoin,
	}, // join(arr, sep) => string
	"last_index": &common.UserFunction{
		Name:  "last_index",
		Value: FuncASSRI(strings.LastIndex),
	}, // last_index(s, substr) => int
	"last_index_any": &common.UserFunction{
		Name:  "last_index_any",
		Value: FuncASSRI(strings.LastIndexAny),
	}, // last_index_any(s, chars) => int
	"repeat": &common.UserFunction{
		Name:  "repeat",
		Value: textRepeat,
	}, // repeat(s, count) => string
	"replace": &common.UserFunction{
		Name:  "replace",
		Value: textReplace,
	}, // replace(s, old, new, n) => string
	"substr": &common.UserFunction{
		Name:  "substr",
		Value: textSubstring,
	}, // substr(s, lower, upper) => string
	"split": &common.UserFunction{
		Name:  "split",
		Value: FuncASSRSs(strings.Split),
	}, // split(s, sep) => [string]
	"split_after": &common.UserFunction{
		Name:  "split_after",
		Value: FuncASSRSs(strings.SplitAfter),
	}, // split_after(s, sep) => [string]
	"split_after_n": &common.UserFunction{
		Name:  "split_after_n",
		Value: FuncASSIRSs(strings.SplitAfterN),
	}, // split_after_n(s, sep, n) => [string]
	"split_n": &common.UserFunction{
		Name:  "split_n",
		Value: FuncASSIRSs(strings.SplitN),
	}, // split_n(s, sep, n) => [string]
	"title": &common.UserFunction{
		Name:  "title",
		Value: FuncASRS(strings.Title),
	}, // title(s) => string
	"to_lower": &common.UserFunction{
		Name:  "to_lower",
		Value: FuncASRS(strings.ToLower),
	}, // to_lower(s) => string
	"to_title": &common.UserFunction{
		Name:  "to_title",
		Value: FuncASRS(strings.ToTitle),
	}, // to_title(s) => string
	"to_upper": &common.UserFunction{
		Name:  "to_upper",
		Value: FuncASRS(strings.ToUpper),
	}, // to_upper(s) => string
	"pad_left": &common.UserFunction{
		Name:  "pad_left",
		Value: textPadLeft,
	}, // pad_left(s, pad_len, pad_with) => string
	"pad_right": &common.UserFunction{
		Name:  "pad_right",
		Value: textPadRight,
	}, // pad_right(s, pad_len, pad_with) => string
	"trim": &common.UserFunction{
		Name:  "trim",
		Value: FuncASSRS(strings.Trim),
	}, // trim(s, cutset) => string
	"trim_left": &common.UserFunction{
		Name:  "trim_left",
		Value: FuncASSRS(strings.TrimLeft),
	}, // trim_left(s, cutset) => string
	"trim_prefix": &common.UserFunction{
		Name:  "trim_prefix",
		Value: FuncASSRS(strings.TrimPrefix),
	}, // trim_prefix(s, prefix) => string
	"trim_right": &common.UserFunction{
		Name:  "trim_right",
		Value: FuncASSRS(strings.TrimRight),
	}, // trim_right(s, cutset) => string
	"trim_space": &common.UserFunction{
		Name:  "trim_space",
		Value: FuncASRS(strings.TrimSpace),
	}, // trim_space(s) => string
	"trim_suffix": &common.UserFunction{
		Name:  "trim_suffix",
		Value: FuncASSRS(strings.TrimSuffix),
	}, // trim_suffix(s, suffix) => string
	"atoi": &common.UserFunction{
		Name:  "atoi",
		Value: FuncASRIE(strconv.Atoi),
	}, // atoi(str) => int/error
	"format_bool": &common.UserFunction{
		Name:  "format_bool",
		Value: textFormatBool,
	}, // format_bool(b) => string
	"format_float": &common.UserFunction{
		Name:  "format_float",
		Value: textFormatFloat,
	}, // format_float(f, fmt, prec, bits) => string
	"format_int": &common.UserFunction{
		Name:  "format_int",
		Value: textFormatInt,
	}, // format_int(i, base) => string
	"itoa": &common.UserFunction{
		Name:  "itoa",
		Value: FuncAIRS(strconv.Itoa),
	}, // itoa(i) => string
	"parse_bool": &common.UserFunction{
		Name:  "parse_bool",
		Value: textParseBool,
	}, // parse_bool(str) => bool/error
	"parse_float": &common.UserFunction{
		Name:  "parse_float",
		Value: textParseFloat,
	}, // parse_float(str, bits) => float/error
	"parse_int": &common.UserFunction{
		Name:  "parse_int",
		Value: textParseInt,
	}, // parse_int(str, base, bits) => int/error
	"quote": &common.UserFunction{
		Name:  "quote",
		Value: FuncASRS(strconv.Quote),
	}, // quote(str) => string
	"unquote": &common.UserFunction{
		Name:  "unquote",
		Value: FuncASRSE(strconv.Unquote),
	}, // unquote(str) => string/error
}

func textREMatch(args ...common.Object) (ret common.Object, err error) {
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

	matched, err := regexp.MatchString(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	if matched {
		ret = common.TrueValue
	} else {
		ret = common.FalseValue
	}

	return
}

func textREFind(args ...common.Object) (ret common.Object, err error) {
	numArgs := len(args)
	if numArgs != 2 && numArgs != 3 {
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

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
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

	if numArgs < 3 {
		m := re.FindStringSubmatchIndex(s2)
		if m == nil {
			ret = common.UndefinedValue
			return
		}

		arr := &common.Array{}
		for i := 0; i < len(m); i += 2 {
			arr.Value = append(arr.Value,
				&common.ImmutableMap{Value: map[string]common.Object{
					"text":  &common.String{Value: s2[m[i]:m[i+1]]},
					"begin": &common.Int{Value: int64(m[i])},
					"end":   &common.Int{Value: int64(m[i+1])},
				}})
		}

		ret = &common.Array{Value: []common.Object{arr}}

		return
	}

	i3, ok := common.ToInt(args[2])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	m := re.FindAllStringSubmatchIndex(s2, i3)
	if m == nil {
		ret = common.UndefinedValue
		return
	}

	arr := &common.Array{}
	for _, m := range m {
		subMatch := &common.Array{}
		for i := 0; i < len(m); i += 2 {
			subMatch.Value = append(subMatch.Value,
				&common.ImmutableMap{Value: map[string]common.Object{
					"text":  &common.String{Value: s2[m[i]:m[i+1]]},
					"begin": &common.Int{Value: int64(m[i])},
					"end":   &common.Int{Value: int64(m[i+1])},
				}})
		}

		arr.Value = append(arr.Value, subMatch)
	}

	ret = arr

	return
}

func textREReplace(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 3 {
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

	s3, ok := common.ToString(args[2])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
	} else {
		s, ok := doTextRegexpReplace(re, s2, s3)
		if !ok {
			return nil, common.ErrStringLimit
		}

		ret = &common.String{Value: s}
	}

	return
}

func textRESplit(args ...common.Object) (ret common.Object, err error) {
	numArgs := len(args)
	if numArgs != 2 && numArgs != 3 {
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

	var i3 = -1
	if numArgs > 2 {
		i3, ok = common.ToInt(args[2])
		if !ok {
			err = common.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	arr := &common.Array{}
	for _, s := range re.Split(s2, i3) {
		arr.Value = append(arr.Value, &common.String{Value: s})
	}

	ret = arr

	return
}

func textRECompile(args ...common.Object) (ret common.Object, err error) {
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

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
	} else {
		ret = makeTextRegexp(re)
	}

	return
}

func textReplace(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 4 {
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

	s3, ok := common.ToString(args[2])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := common.ToInt(args[3])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	s, ok := doTextReplace(s1, s2, s3, i4)
	if !ok {
		err = common.ErrStringLimit
		return
	}

	ret = &common.String{Value: s}

	return
}

func textSubstring(args ...common.Object) (ret common.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
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

	i2, ok := common.ToInt(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	strlen := len(s1)
	i3 := strlen
	if argslen == 3 {
		i3, ok = common.ToInt(args[2])
		if !ok {
			err = common.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	if i2 > i3 {
		err = common.ErrInvalidIndexType
		return
	}

	if i2 < 0 {
		i2 = 0
	} else if i2 > strlen {
		i2 = strlen
	}

	if i3 < 0 {
		i3 = 0
	} else if i3 > strlen {
		i3 = strlen
	}

	ret = &common.String{Value: s1[i2:i3]}

	return
}

func textPadLeft(args ...common.Object) (ret common.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
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

	i2, ok := common.ToInt(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if i2 > common.MaxStringLen {
		return nil, common.ErrStringLimit
	}

	sLen := len(s1)
	if sLen >= i2 {
		ret = &common.String{Value: s1}
		return
	}

	s3 := " "
	if argslen == 3 {
		s3, ok = common.ToString(args[2])
		if !ok {
			err = common.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "string(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	padStrLen := len(s3)
	if padStrLen == 0 {
		ret = &common.String{Value: s1}
		return
	}

	padCount := ((i2 - padStrLen) / padStrLen) + 1
	retStr := strings.Repeat(s3, padCount) + s1
	ret = &common.String{Value: retStr[len(retStr)-i2:]}

	return
}

func textPadRight(args ...common.Object) (ret common.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
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

	i2, ok := common.ToInt(args[1])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if i2 > common.MaxStringLen {
		return nil, common.ErrStringLimit
	}

	sLen := len(s1)
	if sLen >= i2 {
		ret = &common.String{Value: s1}
		return
	}

	s3 := " "
	if argslen == 3 {
		s3, ok = common.ToString(args[2])
		if !ok {
			err = common.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "string(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	padStrLen := len(s3)
	if padStrLen == 0 {
		ret = &common.String{Value: s1}
		return
	}

	padCount := ((i2 - padStrLen) / padStrLen) + 1
	retStr := s1 + strings.Repeat(s3, padCount)
	ret = &common.String{Value: retStr[:i2]}

	return
}

func textRepeat(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		return nil, common.ErrWrongNumArguments
	}

	s1, ok := common.ToString(args[0])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	i2, ok := common.ToInt(args[1])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	if len(s1)*i2 > common.MaxStringLen {
		return nil, common.ErrStringLimit
	}

	return &common.String{Value: strings.Repeat(s1, i2)}, nil
}

func textJoin(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		return nil, common.ErrWrongNumArguments
	}

	var slen int
	var ss1 []string
	switch arg0 := args[0].(type) {
	case *common.Array:
		for idx, a := range arg0.Value {
			as, ok := common.ToString(a)
			if !ok {
				return nil, common.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("first[%d]", idx),
					Expected: "string(compatible)",
					Found:    a.TypeName(),
				}
			}
			slen += len(as)
			ss1 = append(ss1, as)
		}
	case *common.ImmutableArray:
		for idx, a := range arg0.Value {
			as, ok := common.ToString(a)
			if !ok {
				return nil, common.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("first[%d]", idx),
					Expected: "string(compatible)",
					Found:    a.TypeName(),
				}
			}
			slen += len(as)
			ss1 = append(ss1, as)
		}
	default:
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}

	s2, ok := common.ToString(args[1])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	// make sure output length does not exceed the limit
	if slen+len(s2)*(len(ss1)-1) > common.MaxStringLen {
		return nil, common.ErrStringLimit
	}

	return &common.String{Value: strings.Join(ss1, s2)}, nil
}

func textFormatBool(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	b1, ok := args[0].(*common.Bool)
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bool",
			Found:    args[0].TypeName(),
		}
		return
	}

	if b1 == common.TrueValue {
		ret = &common.String{Value: "true"}
	} else {
		ret = &common.String{Value: "false"}
	}

	return
}

func textFormatFloat(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 4 {
		err = common.ErrWrongNumArguments
		return
	}

	f1, ok := args[0].(*common.Float)
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "float",
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

	i3, ok := common.ToInt(args[2])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := common.ToInt(args[3])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &common.String{Value: strconv.FormatFloat(f1.Value, s2[0], i3, i4)}

	return
}

func textFormatInt(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].(*common.Int)
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int",
			Found:    args[0].TypeName(),
		}
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

	ret = &common.String{Value: strconv.FormatInt(i1.Value, i2)}

	return
}

func textParseBool(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		err = common.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].(*common.String)
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
		return
	}

	parsed, err := strconv.ParseBool(s1.Value)
	if err != nil {
		ret = wrapError(err)
		return
	}

	if parsed {
		ret = common.TrueValue
	} else {
		ret = common.FalseValue
	}

	return
}

func textParseFloat(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 2 {
		err = common.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].(*common.String)
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
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

	parsed, err := strconv.ParseFloat(s1.Value, i2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &common.Float{Value: parsed}

	return
}

func textParseInt(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 3 {
		err = common.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].(*common.String)
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
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

	i3, ok := common.ToInt(args[2])
	if !ok {
		err = common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	parsed, err := strconv.ParseInt(s1.Value, i2, i3)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &common.Int{Value: parsed}

	return
}

// Modified implementation of strings.Replace
// to limit the maximum length of output string.
func doTextReplace(s, old, new string, n int) (string, bool) {
	if old == new || n == 0 {
		return s, true // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return s, true // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}

		ssj := s[start:j]
		if w+len(ssj)+len(new) > common.MaxStringLen {
			return "", false
		}

		w += copy(t[w:], ssj)
		w += copy(t[w:], new)
		start = j + len(old)
	}

	ss := s[start:]
	if w+len(ss) > common.MaxStringLen {
		return "", false
	}

	w += copy(t[w:], ss)

	return string(t[0:w]), true
}
