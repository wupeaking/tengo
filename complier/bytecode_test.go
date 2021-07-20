package complier_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/d5/tengo/v2/common"
	"github.com/d5/tengo/v2/complier"
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
			&common.Char{Value: 'y'},
			&common.Float{Value: 93.11},
			compiledFunction(1, 0,
				complier.MakeInstruction(parser.OpConstant, 3),
				complier.MakeInstruction(parser.OpSetLocal, 0),
				complier.MakeInstruction(parser.OpGetGlobal, 0),
				complier.MakeInstruction(parser.OpGetFree, 0)),
			&common.Float{Value: 39.2},
			&common.Int{Value: 192},
			&common.String{Value: "bar"})))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			complier.MakeInstruction(parser.OpConstant, 0),
			complier.MakeInstruction(parser.OpSetGlobal, 0),
			complier.MakeInstruction(parser.OpConstant, 6),
			complier.MakeInstruction(parser.OpPop)),
		objectsArray(
			&common.Int{Value: 55},
			&common.Int{Value: 66},
			&common.Int{Value: 77},
			&common.Int{Value: 88},
			&common.ImmutableMap{
				Value: map[string]common.Object{
					"array": &common.ImmutableArray{
						Value: []common.Object{
							&common.Int{Value: 1},
							&common.Int{Value: 2},
							&common.Int{Value: 3},
							common.TrueValue,
							common.FalseValue,
							common.UndefinedValue,
						},
					},
					"true":  common.TrueValue,
					"false": common.FalseValue,
					"bytes": &common.Bytes{Value: make([]byte, 16)},
					"char":  &common.Char{Value: 'Y'},
					"error": &common.Error{Value: &common.String{
						Value: "some error",
					}},
					"float": &common.Float{Value: -19.84},
					"immutable_array": &common.ImmutableArray{
						Value: []common.Object{
							&common.Int{Value: 1},
							&common.Int{Value: 2},
							&common.Int{Value: 3},
							common.TrueValue,
							common.FalseValue,
							common.UndefinedValue,
						},
					},
					"immutable_map": &common.ImmutableMap{
						Value: map[string]common.Object{
							"a": &common.Int{Value: 1},
							"b": &common.Int{Value: 2},
							"c": &common.Int{Value: 3},
							"d": common.TrueValue,
							"e": common.FalseValue,
							"f": common.UndefinedValue,
						},
					},
					"int": &common.Int{Value: 91},
					"map": &common.Map{
						Value: map[string]common.Object{
							"a": &common.Int{Value: 1},
							"b": &common.Int{Value: 2},
							"c": &common.Int{Value: 3},
							"d": common.TrueValue,
							"e": common.FalseValue,
							"f": common.UndefinedValue,
						},
					},
					"string":    &common.String{Value: "foo bar"},
					"time":      &common.Time{Value: time.Now()},
					"undefined": common.UndefinedValue,
				},
			},
			compiledFunction(1, 0,
				complier.MakeInstruction(parser.OpConstant, 3),
				complier.MakeInstruction(parser.OpSetLocal, 0),
				complier.MakeInstruction(parser.OpGetGlobal, 0),
				complier.MakeInstruction(parser.OpGetFree, 0),
				complier.MakeInstruction(parser.OpBinaryOp, 11),
				complier.MakeInstruction(parser.OpGetFree, 1),
				complier.MakeInstruction(parser.OpBinaryOp, 11),
				complier.MakeInstruction(parser.OpGetLocal, 0),
				complier.MakeInstruction(parser.OpBinaryOp, 11),
				complier.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				complier.MakeInstruction(parser.OpConstant, 2),
				complier.MakeInstruction(parser.OpSetLocal, 0),
				complier.MakeInstruction(parser.OpGetFree, 0),
				complier.MakeInstruction(parser.OpGetLocal, 0),
				complier.MakeInstruction(parser.OpClosure, 4, 2),
				complier.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				complier.MakeInstruction(parser.OpConstant, 1),
				complier.MakeInstruction(parser.OpSetLocal, 0),
				complier.MakeInstruction(parser.OpGetLocal, 0),
				complier.MakeInstruction(parser.OpClosure, 5, 1),
				complier.MakeInstruction(parser.OpReturn, 1))),
		fileSet(srcfile{name: "file1", size: 100},
			srcfile{name: "file2", size: 200})))
}

func TestBytecode_RemoveDuplicates(t *testing.T) {
	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(), objectsArray(
				&common.Char{Value: 'y'},
				&common.Float{Value: 93.11},
				compiledFunction(1, 0,
					complier.MakeInstruction(parser.OpConstant, 3),
					complier.MakeInstruction(parser.OpSetLocal, 0),
					complier.MakeInstruction(parser.OpGetGlobal, 0),
					complier.MakeInstruction(parser.OpGetFree, 0)),
				&common.Float{Value: 39.2},
				&common.Int{Value: 192},
				&common.String{Value: "bar"})),
		bytecode(
			concatInsts(), objectsArray(
				&common.Char{Value: 'y'},
				&common.Float{Value: 93.11},
				compiledFunction(1, 0,
					complier.MakeInstruction(parser.OpConstant, 3),
					complier.MakeInstruction(parser.OpSetLocal, 0),
					complier.MakeInstruction(parser.OpGetGlobal, 0),
					complier.MakeInstruction(parser.OpGetFree, 0)),
				&common.Float{Value: 39.2},
				&common.Int{Value: 192},
				&common.String{Value: "bar"})))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				complier.MakeInstruction(parser.OpConstant, 0),
				complier.MakeInstruction(parser.OpConstant, 1),
				complier.MakeInstruction(parser.OpConstant, 2),
				complier.MakeInstruction(parser.OpConstant, 3),
				complier.MakeInstruction(parser.OpConstant, 4),
				complier.MakeInstruction(parser.OpConstant, 5),
				complier.MakeInstruction(parser.OpConstant, 6),
				complier.MakeInstruction(parser.OpConstant, 7),
				complier.MakeInstruction(parser.OpConstant, 8),
				complier.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&common.Int{Value: 1},
				&common.Float{Value: 2.0},
				&common.Char{Value: '3'},
				&common.String{Value: "four"},
				compiledFunction(1, 0,
					complier.MakeInstruction(parser.OpConstant, 3),
					complier.MakeInstruction(parser.OpConstant, 7),
					complier.MakeInstruction(parser.OpSetLocal, 0),
					complier.MakeInstruction(parser.OpGetGlobal, 0),
					complier.MakeInstruction(parser.OpGetFree, 0)),
				&common.Int{Value: 1},
				&common.Float{Value: 2.0},
				&common.Char{Value: '3'},
				&common.String{Value: "four"})),
		bytecode(
			concatInsts(
				complier.MakeInstruction(parser.OpConstant, 0),
				complier.MakeInstruction(parser.OpConstant, 1),
				complier.MakeInstruction(parser.OpConstant, 2),
				complier.MakeInstruction(parser.OpConstant, 3),
				complier.MakeInstruction(parser.OpConstant, 4),
				complier.MakeInstruction(parser.OpConstant, 0),
				complier.MakeInstruction(parser.OpConstant, 1),
				complier.MakeInstruction(parser.OpConstant, 2),
				complier.MakeInstruction(parser.OpConstant, 3),
				complier.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&common.Int{Value: 1},
				&common.Float{Value: 2.0},
				&common.Char{Value: '3'},
				&common.String{Value: "four"},
				compiledFunction(1, 0,
					complier.MakeInstruction(parser.OpConstant, 3),
					complier.MakeInstruction(parser.OpConstant, 2),
					complier.MakeInstruction(parser.OpSetLocal, 0),
					complier.MakeInstruction(parser.OpGetGlobal, 0),
					complier.MakeInstruction(parser.OpGetFree, 0)))))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				complier.MakeInstruction(parser.OpConstant, 0),
				complier.MakeInstruction(parser.OpConstant, 1),
				complier.MakeInstruction(parser.OpConstant, 2),
				complier.MakeInstruction(parser.OpConstant, 3),
				complier.MakeInstruction(parser.OpConstant, 4)),
			objectsArray(
				&common.Int{Value: 1},
				&common.Int{Value: 2},
				&common.Int{Value: 3},
				&common.Int{Value: 1},
				&common.Int{Value: 3})),
		bytecode(
			concatInsts(
				complier.MakeInstruction(parser.OpConstant, 0),
				complier.MakeInstruction(parser.OpConstant, 1),
				complier.MakeInstruction(parser.OpConstant, 2),
				complier.MakeInstruction(parser.OpConstant, 0),
				complier.MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				&common.Int{Value: 1},
				&common.Int{Value: 2},
				&common.Int{Value: 3})))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			&common.Int{Value: 55},
			&common.Int{Value: 66},
			&common.Int{Value: 77},
			&common.Int{Value: 88},
			compiledFunction(1, 0,
				complier.MakeInstruction(parser.OpConstant, 3),
				complier.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				complier.MakeInstruction(parser.OpConstant, 2),
				complier.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				complier.MakeInstruction(parser.OpConstant, 1),
				complier.MakeInstruction(parser.OpReturn, 1))))
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
	constants []common.Object,
	fileSet *parser.SourceFileSet,
) *complier.Bytecode {
	return &complier.Bytecode{
		FileSet:      fileSet,
		MainFunction: &common.CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func testBytecodeRemoveDuplicates(
	t *testing.T,
	input, expected *complier.Bytecode,
) {
	input.RemoveDuplicates()

	require.Equal(t, expected.FileSet, input.FileSet)
	require.Equal(t, expected.MainFunction, input.MainFunction)
	require.Equal(t, expected.Constants, input.Constants)
}

func testBytecodeSerialization(t *testing.T, b *complier.Bytecode) {
	var buf bytes.Buffer
	err := b.Encode(&buf)
	require.NoError(t, err)

	r := &complier.Bytecode{}
	err = r.Decode(bytes.NewReader(buf.Bytes()), nil)
	require.NoError(t, err)

	require.Equal(t, b.FileSet, r.FileSet)
	require.Equal(t, b.MainFunction, r.MainFunction)
	require.Equal(t, b.Constants, r.Constants)
}
