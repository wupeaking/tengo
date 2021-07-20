package tengo_test

import (
	"strings"
	"testing"
	"time"

	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/require"
)

func TestInstructions_String(t *testing.T) {
	assertInstructionString(t,
		[][]byte{
			MakeInstruction(parser.OpConstant, 1),
			MakeInstruction(parser.OpConstant, 2),
			MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 CONST   1    
0003 CONST   2    
0006 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			MakeInstruction(parser.OpBinaryOp, 11),
			MakeInstruction(parser.OpConstant, 2),
			MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 CONST   2    
0005 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			MakeInstruction(parser.OpBinaryOp, 11),
			MakeInstruction(parser.OpGetLocal, 1),
			MakeInstruction(parser.OpConstant, 2),
			MakeInstruction(parser.OpConstant, 65535),
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
	testCountObjects(t, &Array{}, 1)
	testCountObjects(t, &Array{Value: []Object{
		&Int{Value: 1},
		&Int{Value: 2},
		&Array{Value: []Object{
			&Int{Value: 3},
			&Int{Value: 4},
			&Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, TrueValue, 1)
	testCountObjects(t, FalseValue, 1)
	testCountObjects(t, &BuiltinFunction{}, 1)
	testCountObjects(t, &Bytes{Value: []byte("foobar")}, 1)
	testCountObjects(t, &Char{Value: '가'}, 1)
	testCountObjects(t, &CompiledFunction{}, 1)
	testCountObjects(t, &Error{Value: &Int{Value: 5}}, 2)
	testCountObjects(t, &Float{Value: 19.84}, 1)
	testCountObjects(t, &ImmutableArray{Value: []Object{
		&Int{Value: 1},
		&Int{Value: 2},
		&ImmutableArray{Value: []Object{
			&Int{Value: 3},
			&Int{Value: 4},
			&Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &ImmutableMap{
		Value: map[string]Object{
			"k1": &Int{Value: 1},
			"k2": &Int{Value: 2},
			"k3": &Array{Value: []Object{
				&Int{Value: 3},
				&Int{Value: 4},
				&Int{Value: 5},
			}},
		}}, 7)
	testCountObjects(t, &Int{Value: 1984}, 1)
	testCountObjects(t, &Map{Value: map[string]Object{
		"k1": &Int{Value: 1},
		"k2": &Int{Value: 2},
		"k3": &Array{Value: []Object{
			&Int{Value: 3},
			&Int{Value: 4},
			&Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &String{Value: "foo bar"}, 1)
	testCountObjects(t, &Time{Value: time.Now()}, 1)
	testCountObjects(t, UndefinedValue, 1)
}

func testCountObjects(t *testing.T, o Object, expected int) {
	require.Equal(t, expected, CountObjects(o))
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
		FormatInstructions(concatted, 0), "\n"))
}

func makeInstruction(
	t *testing.T,
	expected []byte,
	opcode parser.Opcode,
	operands ...int,
) {
	inst := MakeInstruction(opcode, operands...)
	require.Equal(t, expected, inst)
}
