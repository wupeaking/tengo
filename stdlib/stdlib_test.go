package stdlib_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/d5/tengo/v2/common"
	"github.com/d5/tengo/v2/require"
	"github.com/d5/tengo/v2/scripts"
	"github.com/d5/tengo/v2/stdlib"
)

type ARR = []interface{}
type MAP = map[string]interface{}
type IARR []interface{}
type IMAP map[string]interface{}

func TestAllModuleNames(t *testing.T) {
	names := stdlib.AllModuleNames()
	require.Equal(t,
		len(stdlib.BuiltinModules)+len(stdlib.SourceModules),
		len(names))
}

func TestModulesRun(t *testing.T) {
	// os.File
	expect(t, `
os := import("os")
out := ""

write_file := func(filename, data) {
	file := os.create(filename)
	if !file { return file }

	if res := file.write(bytes(data)); is_error(res) {
		return res
	}

	return file.close()
}

read_file := func(filename) {
	file := os.open(filename)
	if !file { return file }

	data := bytes(100)
	cnt := file.read(data)
	if  is_error(cnt) {
		return cnt
	}

	file.close()
	return data[:cnt]
}

if write_file("./temp", "foobar") {
	out = string(read_file("./temp"))
}

os.remove("./temp")
`, "foobar")

	// exec.command
	expect(t, `
out := ""
os := import("os")
cmd := os.exec("echo", "foo", "bar")
if !is_error(cmd) {
	out = cmd.output()
}
`, []byte("foo bar\n"))

}

func TestGetModules(t *testing.T) {
	mods := stdlib.GetModuleMap()
	require.Equal(t, 0, mods.Len())

	mods = stdlib.GetModuleMap("os")
	require.Equal(t, 1, mods.Len())
	require.NotNil(t, mods.Get("os"))

	mods = stdlib.GetModuleMap("os", "rand")
	require.Equal(t, 2, mods.Len())
	require.NotNil(t, mods.Get("os"))
	require.NotNil(t, mods.Get("rand"))

	mods = stdlib.GetModuleMap("text", "text")
	require.Equal(t, 1, mods.Len())
	require.NotNil(t, mods.Get("text"))

	mods = stdlib.GetModuleMap("nonexisting", "text")
	require.Equal(t, 1, mods.Len())
	require.NotNil(t, mods.Get("text"))
}

type callres struct {
	t *testing.T
	o interface{}
	e error
}

func (c callres) call(funcName string, args ...interface{}) callres {
	if c.e != nil {
		return c
	}

	var oargs []common.Object
	for _, v := range args {
		oargs = append(oargs, object(v))
	}

	switch o := c.o.(type) {
	case *common.BuiltinModule:
		m, ok := o.Attrs[funcName]
		if !ok {
			return callres{t: c.t, e: fmt.Errorf(
				"function not found: %s", funcName)}
		}

		f, ok := m.(*common.UserFunction)
		if !ok {
			return callres{t: c.t, e: fmt.Errorf(
				"non-callable: %s", funcName)}
		}

		res, err := f.Value(oargs...)
		return callres{t: c.t, o: res, e: err}
	case *common.UserFunction:
		res, err := o.Value(oargs...)
		return callres{t: c.t, o: res, e: err}
	case *common.ImmutableMap:
		m, ok := o.Value[funcName]
		if !ok {
			return callres{t: c.t, e: fmt.Errorf("function not found: %s", funcName)}
		}

		f, ok := m.(*common.UserFunction)
		if !ok {
			return callres{t: c.t, e: fmt.Errorf("non-callable: %s", funcName)}
		}

		res, err := f.Value(oargs...)
		return callres{t: c.t, o: res, e: err}
	default:
		panic(fmt.Errorf("unexpected object: %v (%T)", o, o))
	}
}

func (c callres) expect(expected interface{}, msgAndArgs ...interface{}) {
	require.NoError(c.t, c.e, msgAndArgs...)
	require.Equal(c.t, object(expected), c.o, msgAndArgs...)
}

func (c callres) expectError() {
	require.Error(c.t, c.e)
}

func module(t *testing.T, moduleName string) callres {
	mod := stdlib.GetModuleMap(moduleName).GetBuiltinModule(moduleName)
	if mod == nil {
		return callres{t: t, e: fmt.Errorf("module not found: %s", moduleName)}
	}

	return callres{t: t, o: mod}
}

func object(v interface{}) common.Object {
	switch v := v.(type) {
	case common.Object:
		return v
	case string:
		return &common.String{Value: v}
	case int64:
		return &common.Int{Value: v}
	case int: // for convenience
		return &common.Int{Value: int64(v)}
	case bool:
		if v {
			return common.TrueValue
		}
		return common.FalseValue
	case rune:
		return &common.Char{Value: v}
	case byte: // for convenience
		return &common.Char{Value: rune(v)}
	case float64:
		return &common.Float{Value: v}
	case []byte:
		return &common.Bytes{Value: v}
	case MAP:
		objs := make(map[string]common.Object)
		for k, v := range v {
			objs[k] = object(v)
		}

		return &common.Map{Value: objs}
	case ARR:
		var objs []common.Object
		for _, e := range v {
			objs = append(objs, object(e))
		}

		return &common.Array{Value: objs}
	case IMAP:
		objs := make(map[string]common.Object)
		for k, v := range v {
			objs[k] = object(v)
		}

		return &common.ImmutableMap{Value: objs}
	case IARR:
		var objs []common.Object
		for _, e := range v {
			objs = append(objs, object(e))
		}

		return &common.ImmutableArray{Value: objs}
	case time.Time:
		return &common.Time{Value: v}
	case []int:
		var objs []common.Object
		for _, e := range v {
			objs = append(objs, &common.Int{Value: int64(e)})
		}

		return &common.Array{Value: objs}
	}

	panic(fmt.Errorf("unknown type: %T", v))
}

func expect(t *testing.T, input string, expected interface{}) {
	s := scripts.NewScript([]byte(input))
	s.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	c, err := s.Run()
	require.NoError(t, err)
	require.NotNil(t, c)
	v := c.Get("out")
	require.NotNil(t, v)
	require.Equal(t, expected, v.Value())
}
