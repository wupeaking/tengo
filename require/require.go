package require

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/d5/tengo/v2/common"
	"github.com/d5/tengo/v2/complier"
	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/token"
)

// NoError asserts err is not an error.
func NoError(t *testing.T, err error, msg ...interface{}) {
	if err != nil {
		failExpectedActual(t, "no error", err, msg...)
	}
}

// Error asserts err is an error.
func Error(t *testing.T, err error, msg ...interface{}) {
	if err == nil {
		failExpectedActual(t, "error", err, msg...)
	}
}

// Nil asserts v is nil.
func Nil(t *testing.T, v interface{}, msg ...interface{}) {
	if !isNil(v) {
		failExpectedActual(t, "nil", v, msg...)
	}
}

// True asserts v is true.
func True(t *testing.T, v bool, msg ...interface{}) {
	if !v {
		failExpectedActual(t, "true", v, msg...)
	}
}

// False asserts vis false.
func False(t *testing.T, v bool, msg ...interface{}) {
	if v {
		failExpectedActual(t, "false", v, msg...)
	}
}

// NotNil asserts v is not nil.
func NotNil(t *testing.T, v interface{}, msg ...interface{}) {
	if isNil(v) {
		failExpectedActual(t, "not nil", v, msg...)
	}
}

// IsType asserts expected and actual are of the same type.
func IsType(
	t *testing.T,
	expected, actual interface{},
	msg ...interface{},
) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		failExpectedActual(t, reflect.TypeOf(expected),
			reflect.TypeOf(actual), msg...)
	}
}

// Equal asserts expected and actual are equal.
func Equal(
	t *testing.T,
	expected, actual interface{},
	msg ...interface{},
) {
	if isNil(expected) {
		Nil(t, actual, "expected nil, but got not nil")
		return
	}
	NotNil(t, actual, "expected not nil, but got nil")
	IsType(t, expected, actual, msg...)

	switch expected := expected.(type) {
	case int:
		if expected != actual.(int) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case int64:
		if expected != actual.(int64) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case float64:
		if expected != actual.(float64) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case string:
		if expected != actual.(string) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case []byte:
		if !bytes.Equal(expected, actual.([]byte)) {
			failExpectedActual(t, string(expected),
				string(actual.([]byte)), msg...)
		}
	case []string:
		if !equalStringSlice(expected, actual.([]string)) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case []int:
		if !equalIntSlice(expected, actual.([]int)) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case bool:
		if expected != actual.(bool) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case rune:
		if expected != actual.(rune) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case *complier.Symbol:
		if !equalSymbol(expected, actual.(*complier.Symbol)) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case parser.Pos:
		if expected != actual.(parser.Pos) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case token.Token:
		if expected != actual.(token.Token) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case []common.Object:
		equalObjectSlice(t, expected, actual.([]common.Object), msg...)
	case *common.Int:
		Equal(t, expected.Value, actual.(*common.Int).Value, msg...)
	case *common.Float:
		Equal(t, expected.Value, actual.(*common.Float).Value, msg...)
	case *common.String:
		Equal(t, expected.Value, actual.(*common.String).Value, msg...)
	case *common.Char:
		Equal(t, expected.Value, actual.(*common.Char).Value, msg...)
	case *common.Bool:
		if expected != actual {
			failExpectedActual(t, expected, actual, msg...)
		}
	case *common.Array:
		equalObjectSlice(t, expected.Value,
			actual.(*common.Array).Value, msg...)
	case *common.ImmutableArray:
		equalObjectSlice(t, expected.Value,
			actual.(*common.ImmutableArray).Value, msg...)
	case *common.Bytes:
		if !bytes.Equal(expected.Value, actual.(*common.Bytes).Value) {
			failExpectedActual(t, string(expected.Value),
				string(actual.(*common.Bytes).Value), msg...)
		}
	case *common.Map:
		equalObjectMap(t, expected.Value,
			actual.(*common.Map).Value, msg...)
	case *common.ImmutableMap:
		equalObjectMap(t, expected.Value,
			actual.(*common.ImmutableMap).Value, msg...)
	case *common.CompiledFunction:
		equalCompiledFunction(t, expected,
			actual.(*common.CompiledFunction), msg...)
	case *common.Undefined:
		if expected != actual {
			failExpectedActual(t, expected, actual, msg...)
		}
	case *common.Error:
		Equal(t, expected.Value, actual.(*common.Error).Value, msg...)
	case common.Object:
		if !expected.Equals(actual.(common.Object)) {
			failExpectedActual(t, expected, actual, msg...)
		}
	case *parser.SourceFileSet:
		equalFileSet(t, expected, actual.(*parser.SourceFileSet), msg...)
	case *parser.SourceFile:
		Equal(t, expected.Name, actual.(*parser.SourceFile).Name, msg...)
		Equal(t, expected.Base, actual.(*parser.SourceFile).Base, msg...)
		Equal(t, expected.Size, actual.(*parser.SourceFile).Size, msg...)
		True(t, equalIntSlice(expected.Lines,
			actual.(*parser.SourceFile).Lines), msg...)
	case error:
		if expected != actual.(error) {
			failExpectedActual(t, expected, actual, msg...)
		}
	default:
		panic(fmt.Errorf("type not implemented: %T", expected))
	}
}

// Fail marks the function as having failed but continues execution.
func Fail(t *testing.T, msg ...interface{}) {
	t.Logf("\nError trace:\n\t%s\n%s", strings.Join(errorTrace(), "\n\t"),
		message(msg...))
	t.Fail()
}

func failExpectedActual(
	t *testing.T,
	expected, actual interface{},
	msg ...interface{},
) {
	var addMsg string
	if len(msg) > 0 {
		addMsg = "\nMessage:  " + message(msg...)
	}

	t.Logf("\nError trace:\n\t%s\nExpected: %v\nActual:   %v%s",
		strings.Join(errorTrace(), "\n\t"),
		expected, actual,
		addMsg)
	t.FailNow()
}

func message(formatArgs ...interface{}) string {
	var format string
	var args []interface{}
	if len(formatArgs) > 0 {
		format = formatArgs[0].(string)
	}
	if len(formatArgs) > 1 {
		args = formatArgs[1:]
	}
	return fmt.Sprintf(format, args...)
}

func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalSymbol(a, b *complier.Symbol) bool {
	return a.Name == b.Name &&
		a.Index == b.Index &&
		a.Scope == b.Scope
}

func equalObjectSlice(
	t *testing.T,
	expected, actual []common.Object,
	msg ...interface{},
) {
	Equal(t, len(expected), len(actual), msg...)
	for i := 0; i < len(expected); i++ {
		Equal(t, expected[i], actual[i], msg...)
	}
}

func equalFileSet(
	t *testing.T,
	expected, actual *parser.SourceFileSet,
	msg ...interface{},
) {
	Equal(t, len(expected.Files), len(actual.Files), msg...)
	for i, f := range expected.Files {
		Equal(t, f, actual.Files[i], msg...)
	}
	Equal(t, expected.Base, actual.Base)
	Equal(t, expected.LastFile, actual.LastFile)
}

func equalObjectMap(
	t *testing.T,
	expected, actual map[string]common.Object,
	msg ...interface{},
) {
	Equal(t, len(expected), len(actual), msg...)
	for key, expectedVal := range expected {
		actualVal := actual[key]
		Equal(t, expectedVal, actualVal, msg...)
	}
}

func equalCompiledFunction(
	t *testing.T,
	expected, actual common.Object,
	msg ...interface{},
) {
	expectedT := expected.(*common.CompiledFunction)
	actualT := actual.(*common.CompiledFunction)
	Equal(t,
		complier.FormatInstructions(expectedT.Instructions, 0),
		complier.FormatInstructions(actualT.Instructions, 0), msg...)
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	kind := value.Kind()
	return kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil()
}

func errorTrace() []string {
	var pc uintptr
	file := ""
	line := 0
	var ok bool
	name := ""

	var callers []string
	for i := 0; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {
			break
		}

		if file == "<autogenerated>" {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}
		name = f.Name()

		if name == "testing.tRunner" {
			break
		}

		parts := strings.Split(file, "/")
		file = parts[len(parts)-1]
		if len(parts) > 1 {
			dir := parts[len(parts)-2]
			if dir != "require" ||
				file == "mock_test.go" {
				callers = append(callers, fmt.Sprintf("%s:%d", file, line))
			}
		}

		// Drop the package
		segments := strings.Split(name, ".")
		name = segments[len(segments)-1]
		if isTest(name, "Test") ||
			isTest(name, "Benchmark") ||
			isTest(name, "Example") {
			break
		}
	}
	return callers
}

func isTest(name, prefix string) bool {
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Test" is ok
		return true
	}
	r, _ := utf8.DecodeRuneInString(name[len(prefix):])
	return !unicode.IsLower(r)
}
