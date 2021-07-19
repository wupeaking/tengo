package complier

import (
	"fmt"
	"strings"
	"testing"

	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/require"
)

func TestCompiler_Compile(t *testing.T) {
	expectCompile(t, `1 + 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1; 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 - 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 12),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 * 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 13),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `2 / 1`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 14),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(2),
				intObject(1))))

	expectCompile(t, `true`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpTrue),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `false`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpFalse),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `1 > 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 39),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 < 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 39),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(2),
				intObject(1))))

	expectCompile(t, `1 >= 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 44),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 <= 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 44),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(2),
				intObject(1))))

	expectCompile(t, `1 == 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpEqual),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 != 2`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpNotEqual),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `true == false`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpTrue),
				MakeInstruction(parser.OpFalse),
				MakeInstruction(parser.OpEqual),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `true != false`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpTrue),
				MakeInstruction(parser.OpFalse),
				MakeInstruction(parser.OpNotEqual),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `-1`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpMinus),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1))))

	expectCompile(t, `!true`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpTrue),
				MakeInstruction(parser.OpLNot),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `if true { 10 }; 3333`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpTrue),         // 0000
				MakeInstruction(parser.OpJumpFalsy, 8), // 0001
				MakeInstruction(parser.OpConstant, 0),  // 0004
				MakeInstruction(parser.OpPop),          // 0007
				MakeInstruction(parser.OpConstant, 1),  // 0008
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)), // 0011
			objectsArray(
				intObject(10),
				intObject(3333))))

	expectCompile(t, `if (true) { 10 } else { 20 }; 3333;`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpTrue),          // 0000
				MakeInstruction(parser.OpJumpFalsy, 11), // 0001
				MakeInstruction(parser.OpConstant, 0),   // 0004
				MakeInstruction(parser.OpPop),           // 0007
				MakeInstruction(parser.OpJump, 15),      // 0008
				MakeInstruction(parser.OpConstant, 1),   // 0011
				MakeInstruction(parser.OpPop),           // 0014
				MakeInstruction(parser.OpConstant, 2),   // 0015
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)), // 0018
			objectsArray(
				intObject(10),
				intObject(20),
				intObject(3333))))

	expectCompile(t, `"kami"`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("kami"))))

	expectCompile(t, `"ka" + "mi"`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("ka"),
				stringObject("mi"))))

	expectCompile(t, `a := 1; b := 2; a += b`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpSetGlobal, 1),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 1),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `a := 1; b := 2; a /= b`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpSetGlobal, 1),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 1),
				MakeInstruction(parser.OpBinaryOp, 14),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `[]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpArray, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `[1, 2, 3]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpArray, 3),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3))))

	expectCompile(t, `[1 + 2, 3 - 4, 5 * 6]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpBinaryOp, 12),
				MakeInstruction(parser.OpConstant, 4),
				MakeInstruction(parser.OpConstant, 5),
				MakeInstruction(parser.OpBinaryOp, 13),
				MakeInstruction(parser.OpArray, 3),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3),
				intObject(4),
				intObject(5),
				intObject(6))))

	expectCompile(t, `{}`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpMap, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `{a: 2, b: 4, c: 6}`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpConstant, 4),
				MakeInstruction(parser.OpConstant, 5),
				MakeInstruction(parser.OpMap, 6),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("a"),
				intObject(2),
				stringObject("b"),
				intObject(4),
				stringObject("c"),
				intObject(6))))

	expectCompile(t, `{a: 2 + 3, b: 5 * 6}`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpConstant, 4),
				MakeInstruction(parser.OpConstant, 5),
				MakeInstruction(parser.OpBinaryOp, 13),
				MakeInstruction(parser.OpMap, 4),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("a"),
				intObject(2),
				intObject(3),
				stringObject("b"),
				intObject(5),
				intObject(6))))

	expectCompile(t, `[1, 2, 3][1 + 1]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpArray, 3),
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpIndex),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3))))

	expectCompile(t, `{a: 2}[2 - 1]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpMap, 2),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpBinaryOp, 12),
				MakeInstruction(parser.OpIndex),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("a"),
				intObject(2),
				intObject(1))))

	expectCompile(t, `[1, 2, 3][:]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpArray, 3),
				MakeInstruction(parser.OpNull),
				MakeInstruction(parser.OpNull),
				MakeInstruction(parser.OpSliceIndex),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3))))

	expectCompile(t, `[1, 2, 3][0 : 2]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpArray, 3),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpSliceIndex),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3),
				intObject(0))))

	expectCompile(t, `[1, 2, 3][:2]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpArray, 3),
				MakeInstruction(parser.OpNull),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpSliceIndex),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3))))

	expectCompile(t, `[1, 2, 3][0:]`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpArray, 3),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpNull),
				MakeInstruction(parser.OpSliceIndex),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3),
				intObject(0))))

	expectCompile(t, `f1 := func(a) { return a }; f1([1, 2]...);`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpArray, 2),
				MakeInstruction(parser.OpCall, 1, 1),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpReturn, 1)),
				intObject(1),
				intObject(2))))

	expectCompile(t, `func() { return 5 + 10 }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(5),
				intObject(10),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpConstant, 1),
					MakeInstruction(parser.OpBinaryOp, 11),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `func() { 5 + 10 }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(5),
				intObject(10),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpConstant, 1),
					MakeInstruction(parser.OpBinaryOp, 11),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { 1; 2 }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpConstant, 1),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { 1; return 2 }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpConstant, 1),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `func() { if(true) { return 1 } else { return 2 } }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpTrue),         // 0000
					MakeInstruction(parser.OpJumpFalsy, 9), // 0001
					MakeInstruction(parser.OpConstant, 0),  // 0004
					MakeInstruction(parser.OpReturn, 1),    // 0007
					MakeInstruction(parser.OpConstant, 1),  // 0009
					MakeInstruction(parser.OpReturn, 1))))) // 0012

	expectCompile(t, `func() { 1; if(true) { 2 } else { 3 }; 4 }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 4),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3),
				intObject(4),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),   // 0000
					MakeInstruction(parser.OpPop),           // 0003
					MakeInstruction(parser.OpTrue),          // 0004
					MakeInstruction(parser.OpJumpFalsy, 15), // 0005
					MakeInstruction(parser.OpConstant, 1),   // 0008
					MakeInstruction(parser.OpPop),           // 0011
					MakeInstruction(parser.OpJump, 19),      // 0012
					MakeInstruction(parser.OpConstant, 2),   // 0015
					MakeInstruction(parser.OpPop),           // 0018
					MakeInstruction(parser.OpConstant, 3),   // 0019
					MakeInstruction(parser.OpPop),           // 0022
					MakeInstruction(parser.OpReturn, 0)))))  // 0023

	expectCompile(t, `func() { }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(0, 0,
					MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { 24 }()`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpCall, 0, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(24),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { return 24 }()`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpCall, 0, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(24),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `noArg := func() { 24 }; noArg();`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpCall, 0, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(24),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `noArg := func() { return 24 }; noArg();`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpCall, 0, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(24),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `n := 55; func() { n };`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpGetGlobal, 0),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { n := 55; return n }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				compiledFunction(1, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpDefineLocal, 0),
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `func() { a := 55; b := 77; return a + b }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				intObject(77),
				compiledFunction(2, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpDefineLocal, 0),
					MakeInstruction(parser.OpConstant, 1),
					MakeInstruction(parser.OpDefineLocal, 1),
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpGetLocal, 1),
					MakeInstruction(parser.OpBinaryOp, 11),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `f1 := func(a) { return a }; f1(24);`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpCall, 1, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpReturn, 1)),
				intObject(24))))

	expectCompile(t, `varTest := func(...a) { return a }; varTest(1,2,3);`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpCall, 3, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpReturn, 1)),
				intObject(1), intObject(2), intObject(3))))

	expectCompile(t, `f1 := func(a, b, c) { a; b; return c; }; f1(24, 25, 26);`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpCall, 3, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(3, 3,
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpGetLocal, 1),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpGetLocal, 2),
					MakeInstruction(parser.OpReturn, 1)),
				intObject(24),
				intObject(25),
				intObject(26))))

	expectCompile(t, `func() { n := 55; n = 23; return n }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				intObject(23),
				compiledFunction(1, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpDefineLocal, 0),
					MakeInstruction(parser.OpConstant, 1),
					MakeInstruction(parser.OpSetLocal, 0),
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpReturn, 1)))))
	expectCompile(t, `len([]);`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpGetBuiltin, 0),
				MakeInstruction(parser.OpArray, 0),
				MakeInstruction(parser.OpCall, 1, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `func() { return len([]) }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(0, 0,
					MakeInstruction(parser.OpGetBuiltin, 0),
					MakeInstruction(parser.OpArray, 0),
					MakeInstruction(parser.OpCall, 1, 0),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `func(a) { func(b) { return a + b } }`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					MakeInstruction(parser.OpGetFree, 0),
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpBinaryOp, 11),
					MakeInstruction(parser.OpReturn, 1)),
				compiledFunction(1, 1,
					MakeInstruction(parser.OpGetLocalPtr, 0),
					MakeInstruction(parser.OpClosure, 0, 1),
					MakeInstruction(parser.OpPop),
					MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `
func(a) {
	return func(b) {
		return func(c) {
			return a + b + c
		}
	}
}`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					MakeInstruction(parser.OpGetFree, 0),
					MakeInstruction(parser.OpGetFree, 1),
					MakeInstruction(parser.OpBinaryOp, 11),
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpBinaryOp, 11),
					MakeInstruction(parser.OpReturn, 1)),
				compiledFunction(1, 1,
					MakeInstruction(parser.OpGetFreePtr, 0),
					MakeInstruction(parser.OpGetLocalPtr, 0),
					MakeInstruction(parser.OpClosure, 0, 2),
					MakeInstruction(parser.OpReturn, 1)),
				compiledFunction(1, 1,
					MakeInstruction(parser.OpGetLocalPtr, 0),
					MakeInstruction(parser.OpClosure, 1, 1),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `
g := 55;

func() {
	a := 66;

	return func() {
		b := 77;

		return func() {
			c := 88;

			return g + a + b + c;
		}
	}
}`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpConstant, 6),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				intObject(66),
				intObject(77),
				intObject(88),
				compiledFunction(1, 0,
					MakeInstruction(parser.OpConstant, 3),
					MakeInstruction(parser.OpDefineLocal, 0),
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
					MakeInstruction(parser.OpDefineLocal, 0),
					MakeInstruction(parser.OpGetFreePtr, 0),
					MakeInstruction(parser.OpGetLocalPtr, 0),
					MakeInstruction(parser.OpClosure, 4, 2),
					MakeInstruction(parser.OpReturn, 1)),
				compiledFunction(1, 0,
					MakeInstruction(parser.OpConstant, 1),
					MakeInstruction(parser.OpDefineLocal, 0),
					MakeInstruction(parser.OpGetLocalPtr, 0),
					MakeInstruction(parser.OpClosure, 5, 1),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `for i:=0; i<10; i++ {}`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpBinaryOp, 39),
				MakeInstruction(parser.OpJumpFalsy, 31),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpJump, 6),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(0),
				intObject(10),
				intObject(1))))

	expectCompile(t, `m := {}; for k, v in m {}`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpMap, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpIteratorInit),
				MakeInstruction(parser.OpSetGlobal, 1),
				MakeInstruction(parser.OpGetGlobal, 1),
				MakeInstruction(parser.OpIteratorNext),
				MakeInstruction(parser.OpJumpFalsy, 37),
				MakeInstruction(parser.OpGetGlobal, 1),
				MakeInstruction(parser.OpIteratorKey),
				MakeInstruction(parser.OpSetGlobal, 2),
				MakeInstruction(parser.OpGetGlobal, 1),
				MakeInstruction(parser.OpIteratorValue),
				MakeInstruction(parser.OpSetGlobal, 3),
				MakeInstruction(parser.OpJump, 13),
				MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `a := 0; a == 0 && a != 1 || a < 1`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpSetGlobal, 0),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpEqual),
				MakeInstruction(parser.OpAndJump, 23),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpNotEqual),
				MakeInstruction(parser.OpOrJump, 34),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpGetGlobal, 0),
				MakeInstruction(parser.OpBinaryOp, 39),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(0),
				intObject(1))))

	// unknown module name
	expectCompileError(t, `import("user1")`, "module 'user1' not found")

	// too many errors
	expectCompileError(t, `
r["x"] = {
    @a:1,
    @b:1,
    @c:1,
    @d:1,
    @e:1,
    @f:1,
    @g:1,
    @h:1,
    @i:1,
    @j:1,
    @k:1
}
`, "Parse Error: illegal character U+0040 '@'\n\tat test:3:5 (and 10 more errors)")

	expectCompileError(t, `import("")`, "empty module name")

	// https://github.com/d5/tengo/issues/314
	expectCompileError(t, `
(func() {
	fn := fn()
})()	
`, "unresolved reference 'fn")
}

func TestCompilerErrorReport(t *testing.T) {
	expectCompileError(t, `import("user1")`,
		"Compile Error: module 'user1' not found\n\tat test:1:1")

	expectCompileError(t, `a = 1`,
		"Compile Error: unresolved reference 'a'\n\tat test:1:1")
	expectCompileError(t, `a, b := 1, 2`,
		"Compile Error: tuple assignment not allowed\n\tat test:1:1")
	expectCompileError(t, `a.b := 1`,
		"not allowed with selector")
	expectCompileError(t, `a:=1; a:=3`,
		"Compile Error: 'a' redeclared in this block\n\tat test:1:7")

	expectCompileError(t, `return 5`,
		"Compile Error: return not allowed outside function\n\tat test:1:1")
	expectCompileError(t, `func() { break }`,
		"Compile Error: break not allowed outside loop\n\tat test:1:10")
	expectCompileError(t, `func() { continue }`,
		"Compile Error: continue not allowed outside loop\n\tat test:1:10")
	expectCompileError(t, `func() { export 5 }`,
		"Compile Error: export not allowed inside function\n\tat test:1:10")
}

func TestCompilerDeadCode(t *testing.T) {
	expectCompile(t, `
func() {
	a := 4
	return a

	b := 5 // dead code from here
	c := a
	return b
}`,
		bytecode(
			concatInsts(
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(4),
				intObject(5),
				compiledFunction(0, 0,
					MakeInstruction(parser.OpConstant, 0),
					MakeInstruction(parser.OpDefineLocal, 0),
					MakeInstruction(parser.OpGetLocal, 0),
					MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `
func() {
	if true {
		return 5
		a := 4  // dead code from here
		b := a
		return b
	} else {
		return 4
		c := 5  // dead code from here
		d := c
		return d
	}
}`, bytecode(
		concatInsts(
			MakeInstruction(parser.OpConstant, 2),
			MakeInstruction(parser.OpPop),
			MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(5),
			intObject(4),
			compiledFunction(0, 0,
				MakeInstruction(parser.OpTrue),
				MakeInstruction(parser.OpJumpFalsy, 9),
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpReturn, 1),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `
func() {
	a := 1
	for {
		if a == 5 {
			return 10
		}
		5 + 5
		return 20
		b := a
		return b
	}
}`, bytecode(
		concatInsts(
			MakeInstruction(parser.OpConstant, 4),
			MakeInstruction(parser.OpPop),
			MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(1),
			intObject(5),
			intObject(10),
			intObject(20),
			compiledFunction(0, 0,
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpDefineLocal, 0),
				MakeInstruction(parser.OpGetLocal, 0),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpEqual),
				MakeInstruction(parser.OpJumpFalsy, 19),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpReturn, 1),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpBinaryOp, 11),
				MakeInstruction(parser.OpPop),
				MakeInstruction(parser.OpConstant, 3),
				MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `
func() {
	if true {
		return 5
		a := 4  // dead code from here
		b := a
		return b
	} else {
		return 4
		c := 5  // dead code from here
		d := c
		return d
	}
}`, bytecode(
		concatInsts(
			MakeInstruction(parser.OpConstant, 2),
			MakeInstruction(parser.OpPop),
			MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(5),
			intObject(4),
			compiledFunction(0, 0,
				MakeInstruction(parser.OpTrue),
				MakeInstruction(parser.OpJumpFalsy, 9),
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpReturn, 1),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpReturn, 1)))))
}

func TestCompilerScopes(t *testing.T) {
	expectCompile(t, `
if a := 1; a {
    a = 2
	b := a
} else {
    a = 3
	b := a
}`, bytecode(
		concatInsts(
			MakeInstruction(parser.OpConstant, 0),
			MakeInstruction(parser.OpSetGlobal, 0),
			MakeInstruction(parser.OpGetGlobal, 0),
			MakeInstruction(parser.OpJumpFalsy, 27),
			MakeInstruction(parser.OpConstant, 1),
			MakeInstruction(parser.OpSetGlobal, 0),
			MakeInstruction(parser.OpGetGlobal, 0),
			MakeInstruction(parser.OpSetGlobal, 1),
			MakeInstruction(parser.OpJump, 39),
			MakeInstruction(parser.OpConstant, 2),
			MakeInstruction(parser.OpSetGlobal, 0),
			MakeInstruction(parser.OpGetGlobal, 0),
			MakeInstruction(parser.OpSetGlobal, 2),
			MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(1),
			intObject(2),
			intObject(3))))

	expectCompile(t, `
func() {
	if a := 1; a {
    	a = 2
		b := a
	} else {
    	a = 3
		b := a
	}
}`, bytecode(
		concatInsts(
			MakeInstruction(parser.OpConstant, 3),
			MakeInstruction(parser.OpPop),
			MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(1),
			intObject(2),
			intObject(3),
			compiledFunction(0, 0,
				MakeInstruction(parser.OpConstant, 0),
				MakeInstruction(parser.OpDefineLocal, 0),
				MakeInstruction(parser.OpGetLocal, 0),
				MakeInstruction(parser.OpJumpFalsy, 22),
				MakeInstruction(parser.OpConstant, 1),
				MakeInstruction(parser.OpSetLocal, 0),
				MakeInstruction(parser.OpGetLocal, 0),
				MakeInstruction(parser.OpDefineLocal, 1),
				MakeInstruction(parser.OpJump, 31),
				MakeInstruction(parser.OpConstant, 2),
				MakeInstruction(parser.OpSetLocal, 0),
				MakeInstruction(parser.OpGetLocal, 0),
				MakeInstruction(parser.OpDefineLocal, 1),
				MakeInstruction(parser.OpReturn, 0)))))
}

func concatInsts(instructions ...[]byte) []byte {
	var concat []byte
	for _, i := range instructions {
		concat = append(concat, i...)
	}
	return concat
}

func bytecode(
	instructions []byte,
	constants []Object,
) *Bytecode {
	return &Bytecode{
		FileSet:      parser.NewFileSet(),
		MainFunction: &CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func expectCompile(
	t *testing.T,
	input string,
	expected *Bytecode,
) {
	actual, trace, err := traceCompile(input, nil)

	var ok bool
	defer func() {
		if !ok {
			for _, tr := range trace {
				t.Log(tr)
			}
		}
	}()

	require.NoError(t, err)
	equalBytecode(t, expected, actual)
	ok = true
}

func expectCompileError(t *testing.T, input, expected string) {
	_, trace, err := traceCompile(input, nil)

	var ok bool
	defer func() {
		if !ok {
			for _, tr := range trace {
				t.Log(tr)
			}
		}
	}()

	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), expected),
		"expected error string: %s, got: %s", expected, err.Error())
	ok = true
}

func equalBytecode(t *testing.T, expected, actual *Bytecode) {
	require.Equal(t, expected.MainFunction, actual.MainFunction)
	equalConstants(t, expected.Constants, actual.Constants)
}

func equalConstants(t *testing.T, expected, actual []Object) {
	require.Equal(t, len(expected), len(actual))
	for i := 0; i < len(expected); i++ {
		require.Equal(t, expected[i], actual[i])
	}
}

type compileTracer struct {
	Out []string
}

func (o *compileTracer) Write(p []byte) (n int, err error) {
	o.Out = append(o.Out, string(p))
	return len(p), nil
}

func traceCompile(
	input string,
	symbols map[string]Object,
) (res *Bytecode, trace []string, err error) {
	fileSet := parser.NewFileSet()
	file := fileSet.AddFile("test", -1, len(input))

	p := parser.NewParser(file, []byte(input), nil)

	symTable := NewSymbolTable()
	for name := range symbols {
		symTable.Define(name)
	}
	for idx, fn := range GetAllBuiltinFunctions() {
		symTable.DefineBuiltin(idx, fn.Name)
	}

	tr := &compileTracer{}
	c := NewCompiler(file, symTable, nil, nil, tr)
	parsed, err := p.ParseFile()
	if err != nil {
		return
	}

	err = c.Compile(parsed)
	res = c.Bytecode()
	res.RemoveDuplicates()
	{
		trace = append(trace, fmt.Sprintf("Compiler Trace:\n%s",
			strings.Join(tr.Out, "")))
		trace = append(trace, fmt.Sprintf("Compiled Constants:\n%s",
			strings.Join(res.FormatConstants(), "\n")))
		trace = append(trace, fmt.Sprintf("Compiled Instructions:\n%s\n",
			strings.Join(res.FormatInstructions(), "\n")))
	}
	if err != nil {
		return
	}
	return
}

func objectsArray(o ...Object) []Object {
	return o
}

func intObject(v int64) *Int {
	return &Int{Value: v}
}

func stringObject(v string) *String {
	return &String{Value: v}
}

func compiledFunction(
	numLocals, numParams int,
	insts ...[]byte,
) *CompiledFunction {
	return &CompiledFunction{
		Instructions:  concatInsts(insts...),
		NumLocals:     numLocals,
		NumParameters: numParams,
	}
}
