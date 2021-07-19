package complier

import (
	"bytes"
	"testing"
	"time"

	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/require"
)

type srcfile struct {
	name string
	size int
}

func TestBytecode(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))

	testBytecodeSerialization(t, bytecode(
		concatInsts(), objectsArray(
			&Char{Value: 'y'},
			&Float{Value: 93.11},
			compiledFunction(1, 0,
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpSetLocal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpGetFree, 0)),
			&Float{Value: 39.2},
			&Int{Value: 192},
			&String{Value: "bar"})))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			MakeInstruction(parser.OpConstant, 0),
			MakeInstruction(parser.OpSetGlobal, 0),
			MakeInstruction(parser.OpConstant, 6),
			MakeInstruction(parser.OpPop)),
		objectsArray(
			&Int{Value: 55},
			&Int{Value: 66},
			&Int{Value: 77},
			&Int{Value: 88},
			&ImmutableMap{
				Value: map[string]Object{
					"array": &ImmutableArray{
						Value: []Object{
							&Int{Value: 1},
							&Int{Value: 2},
							&Int{Value: 3},
							TrueValue,
							FalseValue,
							UndefinedValue,
						},
					},
					"true":  TrueValue,
					"false": FalseValue,
					"bytes": &Bytes{Value: make([]byte, 16)},
					"char":  &Char{Value: 'Y'},
					"error": &Error{Value: &String{
						Value: "some error",
					}},
					"float": &Float{Value: -19.84},
					"immutable_array": &ImmutableArray{
						Value: []Object{
							&Int{Value: 1},
							&Int{Value: 2},
							&Int{Value: 3},
							TrueValue,
							FalseValue,
							UndefinedValue,
						},
					},
					"immutable_map": &ImmutableMap{
						Value: map[string]Object{
							"a": &Int{Value: 1},
							"b": &Int{Value: 2},
							"c": &Int{Value: 3},
							"d": TrueValue,
							"e": FalseValue,
							"f": UndefinedValue,
						},
					},
					"int": &Int{Value: 91},
					"map": &Map{
						Value: map[string]Object{
							"a": &Int{Value: 1},
							"b": &Int{Value: 2},
							"c": &Int{Value: 3},
							"d": TrueValue,
							"e": FalseValue,
							"f": UndefinedValue,
						},
					},
					"string":    &String{Value: "foo bar"},
					"time":      &Time{Value: time.Now()},
					"undefined": UndefinedValue,
				},
			},
			compiledFunction(1, 0,
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpSetLocal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpGetFree, 0),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpGetFree, 1),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpGetLocal, 0),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpSetLocal, 0),
				MakeInstruction(parser.OpGetFree, 0),
				MakeInstruction(parser.OpGetLocal, 0),
				MakeInstruction(parser.OpClosure, 4, 2),
				MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpSetLocal, 0),
				MakeInstruction(parser.OpGetLocal, 0),
				MakeInstruction(parser.OpClosure, 5, 1),
				MakeInstruction(parser.OpReturn, 1))),
		fileSet(srcfile{name: "file1", size: 100},
			srcfile{name: "file2", size: 200})))
}

func TestBytecode_RemoveDuplicates(t *testing.T) {
	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(), objectsArray(
				&Char{Value: 'y'},
				&Float{Value: 93.11},
				compiledFunction(1, 0,
					MakeInstruction(parser.OpConstant, 3),
					MakeInstruction(parser.OpSetLocal, 0),
					MakeInstruction(parser.OpGetGlobal, 0),
					MakeInstruction(parser.OpGetFree, 0)),
				&Float{Value: 39.2},
				&Int{Value: 192},
				&String{Value: "bar"})),
		bytecode(
			concatInsts(), objectsArray(
				&Char{Value: 'y'},
				&Float{Value: 93.11},
				compiledFunction(1, 0,
					MakeInstruction(parser.OpConstant, 3),
					MakeInstruction(parser.OpSetLocal, 0),
					MakeInstruction(parser.OpGetGlobal, 0),
					MakeInstruction(parser.OpGetFree, 0)),
				&Float{Value: 39.2},
				&Int{Value: 192},
				&String{Value: "bar"})))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpConstant, 4),
				MakeInstruction(parser.OpConstant, 5),
				MakeInstruction(parser.OpConstant, 6),
				MakeInstruction(parser.OpConstant, 7),
				MakeInstruction(parser.OpConstant, 8),
				MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&Int{Value: 1},
				&Float{Value: 2.0},
				&Char{Value: '3'},
				&String{Value: "four"},
				compiledFunction(1, 0,
					MakeInstruction(parser.OpConstant, 3),
					MakeInstruction(parser.OpConstant, 7),
					MakeInstruction(parser.OpSetLocal, 0),
					MakeInstruction(parser.OpGetGlobal, 0),
					MakeInstruction(parser.OpGetFree, 0)),
				&Int{Value: 1},
				&Float{Value: 2.0},
				&Char{Value: '3'},
				&String{Value: "four"})),
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpConstant, 4),
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&Int{Value: 1},
				&Float{Value: 2.0},
				&Char{Value: '3'},
				&String{Value: "four"},
				compiledFunction(1, 0,
					MakeInstruction(parser.OpConstant, 3),
					MakeInstruction(parser.OpConstant, 2),
					MakeInstruction(parser.OpSetLocal, 0),
					MakeInstruction(parser.OpGetGlobal, 0),
					MakeInstruction(parser.OpGetFree, 0)))))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpConstant, 4)),
			objectsArray(
				&Int{Value: 1},
				&Int{Value: 2},
				&Int{Value: 3},
				&Int{Value: 1},
				&Int{Value: 3})),
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				&Int{Value: 1},
				&Int{Value: 2},
				&Int{Value: 3})))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			&Int{Value: 55},
			&Int{Value: 66},
			&Int{Value: 77},
			&Int{Value: 88},
			compiledFunction(1, 0,
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpReturn, 1))))
	require.Equal(t, 7, b.CountObjects())
}

func fileSet(files ...srcfile) *parser.SourceFileSet {
	fileSet := parser.NewFileSet()
	for _, f := range files {
		fileSet.AddFile(f.name, -1, f.size)
	}
	return fileSet
}

func bytecodeFileSet(
	instructions []byte,
	constants []Object,
	fileSet *parser.SourceFileSet,
) *Bytecode {
	return &Bytecode{
		FileSet:      fileSet,
		MainFunction: &CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func testBytecodeRemoveDuplicates(
	t *testing.T,
	input, expected *Bytecode,
) {
	input.RemoveDuplicates()

	require.Equal(t, expected.FileSet, input.FileSet)
	require.Equal(t, expected.MainFunction, input.MainFunction)
	require.Equal(t, expected.Constants, input.Constants)
}

func testBytecodeSerialization(t *testing.T, b *Bytecode) {
	var buf bytes.Buffer
	err := b.Encode(&buf)
	require.NoError(t, err)

	r := &Bytecode{}
	err = r.Decode(bytes.NewReader(buf.Bytes()), nil)
	require.NoError(t, err)

	require.Equal(t, b.FileSet, r.FileSet)
	require.Equal(t, b.MainFunction, r.MainFunction)
	require.Equal(t, b.Constants, r.Constants)
}
