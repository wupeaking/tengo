package tengo_test

import (
	"strings"
	"testing"
	"time"

	"github.com/d5/tengo/v2/common"
	"github.com/d5/tengo/v2/complier"
	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/require"
)

func TestInstructions_String(t *testing.T) {
	assertInstructionString(t,
		[][]byte{
			complier.MakeInstruction(parser.OpConstant, 1),
			complier.MakeInstruction(parser.OpConstant, 2),
			complier.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 CONST   1    
0003 CONST   2    
0006 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			complier.MakeInstruction(parser.OpBinaryOp, 11),
			complier.MakeInstruction(parser.OpConstant, 2),
			complier.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 CONST   2    
0005 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			complier.MakeInstruction(parser.OpBinaryOp, 11),
			complier.MakeInstruction(parser.OpGetLocal, 1),
			complier.MakeInstruction(parser.OpConstant, 2),
			complier.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 GETL    1    
0004 CONST   2    
0007 CONST   65535`)
}

func TestMakeInstruction(t *testing.T) {
	makeInstruction(t, []byte{parser.OpConstant, 0, 0},
		parser.OpConstant, 0)
	makeInstruction(t, []byte{parser.OpConstant, 0, 1},
		parser.OpConstant, 1)
	makeInstruction(t, []byte{parser.OpConstant, 255, 254},
		parser.OpConstant, 65534)
	makeInstruction(t, []byte{parser.OpPop}, parser.OpPop)
	makeInstruction(t, []byte{parser.OpTrue}, parser.OpTrue)
	makeInstruction(t, []byte{parser.OpFalse}, parser.OpFalse)
}

func TestNumObjects(t *testing.T) {
	testCountObjects(t, &common.Array{}, 1)
	testCountObjects(t, &common.Array{Value: []common.Object{
		&common.Int{Value: 1},
		&common.Int{Value: 2},
		&common.Array{Value: []common.Object{
			&common.Int{Value: 3},
			&common.Int{Value: 4},
			&common.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, common.TrueValue, 1)
	testCountObjects(t, common.FalseValue, 1)
	testCountObjects(t, &common.BuiltinFunction{}, 1)
	testCountObjects(t, &common.Bytes{Value: []byte("foobar")}, 1)
	testCountObjects(t, &common.Char{Value: '가'}, 1)
	testCountObjects(t, &common.CompiledFunction{}, 1)
	testCountObjects(t, &common.Error{Value: &common.Int{Value: 5}}, 2)
	testCountObjects(t, &common.Float{Value: 19.84}, 1)
	testCountObjects(t, &common.ImmutableArray{Value: []common.Object{
		&common.Int{Value: 1},
		&common.Int{Value: 2},
		&common.ImmutableArray{Value: []common.Object{
			&common.Int{Value: 3},
			&common.Int{Value: 4},
			&common.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &common.ImmutableMap{
		Value: map[string]common.Object{
			"k1": &common.Int{Value: 1},
			"k2": &common.Int{Value: 2},
			"k3": &common.Array{Value: []common.Object{
				&common.Int{Value: 3},
				&common.Int{Value: 4},
				&common.Int{Value: 5},
			}},
		}}, 7)
	testCountObjects(t, &common.Int{Value: 1984}, 1)
	testCountObjects(t, &common.Map{Value: map[string]common.Object{
		"k1": &common.Int{Value: 1},
		"k2": &common.Int{Value: 2},
		"k3": &common.Array{Value: []common.Object{
			&common.Int{Value: 3},
			&common.Int{Value: 4},
			&common.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &common.String{Value: "foo bar"}, 1)
	testCountObjects(t, &common.Time{Value: time.Now()}, 1)
	testCountObjects(t, common.UndefinedValue, 1)
}

func testCountObjects(t *testing.T, o common.Object, expected int) {
	require.Equal(t, expected, common.CountObjects(o))
}

func assertInstructionString(
	t *testing.T,
	instructions [][]byte,
	expected string,
) {
	concatted := make([]byte, 0)
	for _, e := range instructions {
		concatted = append(concatted, e...)
	}
	require.Equal(t, expected, strings.Join(
		complier.FormatInstructions(concatted, 0), "\n"))
}

func makeInstruction(
	t *testing.T,
	expected []byte,
	opcode parser.Opcode,
	operands ...int,
) {
	inst := complier.MakeInstruction(opcode, operands...)
	require.Equal(t, expected, inst)
}
