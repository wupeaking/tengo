package stdlib

import (
	"fmt"

	"github.com/d5/tengo/v2/common"
)

var fmtModule = map[string]common.Object{
	"print":   &common.UserFunction{Name: "print", Value: fmtPrint},
	"printf":  &common.UserFunction{Name: "printf", Value: fmtPrintf},
	"println": &common.UserFunction{Name: "println", Value: fmtPrintln},
	"sprintf": &common.UserFunction{Name: "sprintf", Value: fmtSprintf},
}

func fmtPrint(args ...common.Object) (ret common.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtPrintf(args ...common.Object) (ret common.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, common.ErrWrongNumArguments
	}

	format, ok := args[0].(*common.String)
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		fmt.Print(format)
		return nil, nil
	}

	s, err := common.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	fmt.Print(s)
	return nil, nil
}

func fmtPrintln(args ...common.Object) (ret common.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	printArgs = append(printArgs, "\n")
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtSprintf(args ...common.Object) (ret common.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, common.ErrWrongNumArguments
	}

	format, ok := args[0].(*common.String)
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := common.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &common.String{Value: s}, nil
}

func getPrintArgs(args ...common.Object) ([]interface{}, error) {
	var printArgs []interface{}
	l := 0
	for _, arg := range args {
		s, _ := common.ToString(arg)
		slen := len(s)
		// make sure length does not exceed the limit
		if l+slen > common.MaxStringLen {
			return nil, common.ErrStringLimit
		}
		l += slen
		printArgs = append(printArgs, s)
	}
	return printArgs, nil
}
