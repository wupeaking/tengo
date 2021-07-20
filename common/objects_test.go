package common_test

import (
	"testing"

	"github.com/d5/tengo/v2/common"
	"github.com/d5/tengo/v2/require"
	"github.com/d5/tengo/v2/token"
)

func TestObject_TypeName(t *testing.T) {
	var o common.Object = &common.Int{}
	require.Equal(t, "int", o.TypeName())
	o = &common.Float{}
	require.Equal(t, "float", o.TypeName())
	o = &common.Char{}
	require.Equal(t, "char", o.TypeName())
	o = &common.String{}
	require.Equal(t, "string", o.TypeName())
	o = &common.Bool{}
	require.Equal(t, "bool", o.TypeName())
	o = &common.Array{}
	require.Equal(t, "array", o.TypeName())
	o = &common.Map{}
	require.Equal(t, "map", o.TypeName())
	o = &common.ArrayIterator{}
	require.Equal(t, "array-iterator", o.TypeName())
	o = &common.StringIterator{}
	require.Equal(t, "string-iterator", o.TypeName())
	o = &common.MapIterator{}
	require.Equal(t, "map-iterator", o.TypeName())
	o = &common.BuiltinFunction{Name: "fn"}
	require.Equal(t, "builtin-function:fn", o.TypeName())
	o = &common.UserFunction{Name: "fn"}
	require.Equal(t, "user-function:fn", o.TypeName())
	o = &common.CompiledFunction{}
	require.Equal(t, "compiled-function", o.TypeName())
	o = &common.Undefined{}
	require.Equal(t, "undefined", o.TypeName())
	o = &common.Error{}
	require.Equal(t, "error", o.TypeName())
	o = &common.Bytes{}
	require.Equal(t, "bytes", o.TypeName())
}

func TestObject_IsFalsy(t *testing.T) {
	var o common.Object = &common.Int{Value: 0}
	require.True(t, o.IsFalsy())
	o = &common.Int{Value: 1}
	require.False(t, o.IsFalsy())
	o = &common.Float{Value: 0}
	require.False(t, o.IsFalsy())
	o = &common.Float{Value: 1}
	require.False(t, o.IsFalsy())
	o = &common.Char{Value: ' '}
	require.False(t, o.IsFalsy())
	o = &common.Char{Value: 'T'}
	require.False(t, o.IsFalsy())
	o = &common.String{Value: ""}
	require.True(t, o.IsFalsy())
	o = &common.String{Value: " "}
	require.False(t, o.IsFalsy())
	o = &common.Array{Value: nil}
	require.True(t, o.IsFalsy())
	o = &common.Array{Value: []common.Object{nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &common.Map{Value: nil}
	require.True(t, o.IsFalsy())
	o = &common.Map{Value: map[string]common.Object{"a": nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &common.StringIterator{}
	require.True(t, o.IsFalsy())
	o = &common.ArrayIterator{}
	require.True(t, o.IsFalsy())
	o = &common.MapIterator{}
	require.True(t, o.IsFalsy())
	o = &common.BuiltinFunction{}
	require.False(t, o.IsFalsy())
	o = &common.CompiledFunction{}
	require.False(t, o.IsFalsy())
	o = &common.Undefined{}
	require.True(t, o.IsFalsy())
	o = &common.Error{}
	require.True(t, o.IsFalsy())
	o = &common.Bytes{}
	require.True(t, o.IsFalsy())
	o = &common.Bytes{Value: []byte{1, 2}}
	require.False(t, o.IsFalsy())
}

func TestObject_String(t *testing.T) {
	var o common.Object = &common.Int{Value: 0}
	require.Equal(t, "0", o.String())
	o = &common.Int{Value: 1}
	require.Equal(t, "1", o.String())
	o = &common.Float{Value: 0}
	require.Equal(t, "0", o.String())
	o = &common.Float{Value: 1}
	require.Equal(t, "1", o.String())
	o = &common.Char{Value: ' '}
	require.Equal(t, " ", o.String())
	o = &common.Char{Value: 'T'}
	require.Equal(t, "T", o.String())
	o = &common.String{Value: ""}
	require.Equal(t, `""`, o.String())
	o = &common.String{Value: " "}
	require.Equal(t, `" "`, o.String())
	o = &common.Array{Value: nil}
	require.Equal(t, "[]", o.String())
	o = &common.Map{Value: nil}
	require.Equal(t, "{}", o.String())
	o = &common.Error{Value: nil}
	require.Equal(t, "error", o.String())
	o = &common.Error{Value: &common.String{Value: "error 1"}}
	require.Equal(t, `error: "error 1"`, o.String())
	o = &common.StringIterator{}
	require.Equal(t, "<string-iterator>", o.String())
	o = &common.ArrayIterator{}
	require.Equal(t, "<array-iterator>", o.String())
	o = &common.MapIterator{}
	require.Equal(t, "<map-iterator>", o.String())
	o = &common.Undefined{}
	require.Equal(t, "<undefined>", o.String())
	o = &common.Bytes{}
	require.Equal(t, "", o.String())
	o = &common.Bytes{Value: []byte("foo")}
	require.Equal(t, "foo", o.String())
}

func TestObject_BinaryOp(t *testing.T) {
	var o common.Object = &common.Char{}
	_, err := o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.Bool{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.Map{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.ArrayIterator{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.StringIterator{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.MapIterator{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.BuiltinFunction{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.CompiledFunction{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.Undefined{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
	o = &common.Error{}
	_, err = o.BinaryOp(token.Add, common.UndefinedValue)
	require.Error(t, err)
}

func TestArray_BinaryOp(t *testing.T) {
	testBinaryOp(t, &common.Array{Value: nil}, token.Add,
		&common.Array{Value: nil}, &common.Array{Value: nil})
	testBinaryOp(t, &common.Array{Value: nil}, token.Add,
		&common.Array{Value: []common.Object{}}, &common.Array{Value: nil})
	testBinaryOp(t, &common.Array{Value: []common.Object{}}, token.Add,
		&common.Array{Value: nil}, &common.Array{Value: []common.Object{}})
	testBinaryOp(t, &common.Array{Value: []common.Object{}}, token.Add,
		&common.Array{Value: []common.Object{}},
		&common.Array{Value: []common.Object{}})
	testBinaryOp(t, &common.Array{Value: nil}, token.Add,
		&common.Array{Value: []common.Object{
			&common.Int{Value: 1},
		}}, &common.Array{Value: []common.Object{
			&common.Int{Value: 1},
		}})
	testBinaryOp(t, &common.Array{Value: nil}, token.Add,
		&common.Array{Value: []common.Object{
			&common.Int{Value: 1},
			&common.Int{Value: 2},
			&common.Int{Value: 3},
		}}, &common.Array{Value: []common.Object{
			&common.Int{Value: 1},
			&common.Int{Value: 2},
			&common.Int{Value: 3},
		}})
	testBinaryOp(t, &common.Array{Value: []common.Object{
		&common.Int{Value: 1},
		&common.Int{Value: 2},
		&common.Int{Value: 3},
	}}, token.Add, &common.Array{Value: nil},
		&common.Array{Value: []common.Object{
			&common.Int{Value: 1},
			&common.Int{Value: 2},
			&common.Int{Value: 3},
		}})
	testBinaryOp(t, &common.Array{Value: []common.Object{
		&common.Int{Value: 1},
		&common.Int{Value: 2},
		&common.Int{Value: 3},
	}}, token.Add, &common.Array{Value: []common.Object{
		&common.Int{Value: 4},
		&common.Int{Value: 5},
		&common.Int{Value: 6},
	}}, &common.Array{Value: []common.Object{
		&common.Int{Value: 1},
		&common.Int{Value: 2},
		&common.Int{Value: 3},
		&common.Int{Value: 4},
		&common.Int{Value: 5},
		&common.Int{Value: 6},
	}})
}

func TestError_Equals(t *testing.T) {
	err1 := &common.Error{Value: &common.String{Value: "some error"}}
	err2 := err1
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = &common.Error{Value: &common.String{Value: "some error"}}
	require.False(t, err1.Equals(err2))
	require.False(t, err2.Equals(err1))
}

func TestFloat_BinaryOp(t *testing.T) {
	// float + float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &common.Float{Value: l}, token.Add,
				&common.Float{Value: r}, &common.Float{Value: l + r})
		}
	}

	// float - float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &common.Float{Value: l}, token.Sub,
				&common.Float{Value: r}, &common.Float{Value: l - r})
		}
	}

	// float * float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &common.Float{Value: l}, token.Mul,
				&common.Float{Value: r}, &common.Float{Value: l * r})
		}
	}

	// float / float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			if r != 0 {
				testBinaryOp(t, &common.Float{Value: l}, token.Quo,
					&common.Float{Value: r}, &common.Float{Value: l / r})
			}
		}
	}

	// float < float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &common.Float{Value: l}, token.Less,
				&common.Float{Value: r}, boolValue(l < r))
		}
	}

	// float > float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &common.Float{Value: l}, token.Greater,
				&common.Float{Value: r}, boolValue(l > r))
		}
	}

	// float <= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &common.Float{Value: l}, token.LessEq,
				&common.Float{Value: r}, boolValue(l <= r))
		}
	}

	// float >= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &common.Float{Value: l}, token.GreaterEq,
				&common.Float{Value: r}, boolValue(l >= r))
		}
	}

	// float + int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Float{Value: l}, token.Add,
				&common.Int{Value: r}, &common.Float{Value: l + float64(r)})
		}
	}

	// float - int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Float{Value: l}, token.Sub,
				&common.Int{Value: r}, &common.Float{Value: l - float64(r)})
		}
	}

	// float * int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Float{Value: l}, token.Mul,
				&common.Int{Value: r}, &common.Float{Value: l * float64(r)})
		}
	}

	// float / int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &common.Float{Value: l}, token.Quo,
					&common.Int{Value: r},
					&common.Float{Value: l / float64(r)})
			}
		}
	}

	// float < int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Float{Value: l}, token.Less,
				&common.Int{Value: r}, boolValue(l < float64(r)))
		}
	}

	// float > int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Float{Value: l}, token.Greater,
				&common.Int{Value: r}, boolValue(l > float64(r)))
		}
	}

	// float <= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Float{Value: l}, token.LessEq,
				&common.Int{Value: r}, boolValue(l <= float64(r)))
		}
	}

	// float >= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Float{Value: l}, token.GreaterEq,
				&common.Int{Value: r}, boolValue(l >= float64(r)))
		}
	}
}

func TestInt_BinaryOp(t *testing.T) {
	// int + int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Int{Value: l}, token.Add,
				&common.Int{Value: r}, &common.Int{Value: l + r})
		}
	}

	// int - int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Int{Value: l}, token.Sub,
				&common.Int{Value: r}, &common.Int{Value: l - r})
		}
	}

	// int * int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Int{Value: l}, token.Mul,
				&common.Int{Value: r}, &common.Int{Value: l * r})
		}
	}

	// int / int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &common.Int{Value: l}, token.Quo,
					&common.Int{Value: r}, &common.Int{Value: l / r})
			}
		}
	}

	// int % int
	for l := int64(-4); l <= 4; l++ {
		for r := -int64(-4); r <= 4; r++ {
			if r == 0 {
				testBinaryOp(t, &common.Int{Value: l}, token.Rem,
					&common.Int{Value: r}, &common.Int{Value: l % r})
			}
		}
	}

	// int & int
	testBinaryOp(t,
		&common.Int{Value: 0}, token.And, &common.Int{Value: 0},
		&common.Int{Value: int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.And, &common.Int{Value: 0},
		&common.Int{Value: int64(1) & int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 0}, token.And, &common.Int{Value: 1},
		&common.Int{Value: int64(0) & int64(1)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.And, &common.Int{Value: 1},
		&common.Int{Value: int64(1)})
	testBinaryOp(t,
		&common.Int{Value: 0}, token.And, &common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(0) & int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.And, &common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(1) & int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: int64(0xffffffff)}, token.And,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: 1984}, token.And,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(1984) & int64(0xffffffff)})
	testBinaryOp(t, &common.Int{Value: -1984}, token.And,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(-1984) & int64(0xffffffff)})

	// int | int
	testBinaryOp(t,
		&common.Int{Value: 0}, token.Or, &common.Int{Value: 0},
		&common.Int{Value: int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.Or, &common.Int{Value: 0},
		&common.Int{Value: int64(1) | int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 0}, token.Or, &common.Int{Value: 1},
		&common.Int{Value: int64(0) | int64(1)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.Or, &common.Int{Value: 1},
		&common.Int{Value: int64(1)})
	testBinaryOp(t,
		&common.Int{Value: 0}, token.Or, &common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(0) | int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.Or, &common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(1) | int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: int64(0xffffffff)}, token.Or,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: 1984}, token.Or,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(1984) | int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: -1984}, token.Or,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(-1984) | int64(0xffffffff)})

	// int ^ int
	testBinaryOp(t,
		&common.Int{Value: 0}, token.Xor, &common.Int{Value: 0},
		&common.Int{Value: int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.Xor, &common.Int{Value: 0},
		&common.Int{Value: int64(1) ^ int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 0}, token.Xor, &common.Int{Value: 1},
		&common.Int{Value: int64(0) ^ int64(1)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.Xor, &common.Int{Value: 1},
		&common.Int{Value: int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 0}, token.Xor, &common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(0) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.Xor, &common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(1) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: int64(0xffffffff)}, token.Xor,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 1984}, token.Xor,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(1984) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: -1984}, token.Xor,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(-1984) ^ int64(0xffffffff)})

	// int &^ int
	testBinaryOp(t,
		&common.Int{Value: 0}, token.AndNot, &common.Int{Value: 0},
		&common.Int{Value: int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.AndNot, &common.Int{Value: 0},
		&common.Int{Value: int64(1) &^ int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 0}, token.AndNot,
		&common.Int{Value: 1}, &common.Int{Value: int64(0) &^ int64(1)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.AndNot, &common.Int{Value: 1},
		&common.Int{Value: int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 0}, token.AndNot,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(0) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: 1}, token.AndNot,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(1) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: int64(0xffffffff)}, token.AndNot,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(0)})
	testBinaryOp(t,
		&common.Int{Value: 1984}, token.AndNot,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(1984) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&common.Int{Value: -1984}, token.AndNot,
		&common.Int{Value: int64(0xffffffff)},
		&common.Int{Value: int64(-1984) &^ int64(0xffffffff)})

	// int << int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&common.Int{Value: 0}, token.Shl, &common.Int{Value: s},
			&common.Int{Value: int64(0) << uint(s)})
		testBinaryOp(t,
			&common.Int{Value: 1}, token.Shl, &common.Int{Value: s},
			&common.Int{Value: int64(1) << uint(s)})
		testBinaryOp(t,
			&common.Int{Value: 2}, token.Shl, &common.Int{Value: s},
			&common.Int{Value: int64(2) << uint(s)})
		testBinaryOp(t,
			&common.Int{Value: -1}, token.Shl, &common.Int{Value: s},
			&common.Int{Value: int64(-1) << uint(s)})
		testBinaryOp(t,
			&common.Int{Value: -2}, token.Shl, &common.Int{Value: s},
			&common.Int{Value: int64(-2) << uint(s)})
		testBinaryOp(t,
			&common.Int{Value: int64(0xffffffff)}, token.Shl,
			&common.Int{Value: s},
			&common.Int{Value: int64(0xffffffff) << uint(s)})
	}

	// int >> int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&common.Int{Value: 0}, token.Shr, &common.Int{Value: s},
			&common.Int{Value: int64(0) >> uint(s)})
		testBinaryOp(t,
			&common.Int{Value: 1}, token.Shr, &common.Int{Value: s},
			&common.Int{Value: int64(1) >> uint(s)})
		testBinaryOp(t,
			&common.Int{Value: 2}, token.Shr, &common.Int{Value: s},
			&common.Int{Value: int64(2) >> uint(s)})
		testBinaryOp(t,
			&common.Int{Value: -1}, token.Shr, &common.Int{Value: s},
			&common.Int{Value: int64(-1) >> uint(s)})
		testBinaryOp(t,
			&common.Int{Value: -2}, token.Shr, &common.Int{Value: s},
			&common.Int{Value: int64(-2) >> uint(s)})
		testBinaryOp(t,
			&common.Int{Value: int64(0xffffffff)}, token.Shr,
			&common.Int{Value: s},
			&common.Int{Value: int64(0xffffffff) >> uint(s)})
	}

	// int < int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Int{Value: l}, token.Less,
				&common.Int{Value: r}, boolValue(l < r))
		}
	}

	// int > int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Int{Value: l}, token.Greater,
				&common.Int{Value: r}, boolValue(l > r))
		}
	}

	// int <= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Int{Value: l}, token.LessEq,
				&common.Int{Value: r}, boolValue(l <= r))
		}
	}

	// int >= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &common.Int{Value: l}, token.GreaterEq,
				&common.Int{Value: r}, boolValue(l >= r))
		}
	}

	// int + float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &common.Int{Value: l}, token.Add,
				&common.Float{Value: r},
				&common.Float{Value: float64(l) + r})
		}
	}

	// int - float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &common.Int{Value: l}, token.Sub,
				&common.Float{Value: r},
				&common.Float{Value: float64(l) - r})
		}
	}

	// int * float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &common.Int{Value: l}, token.Mul,
				&common.Float{Value: r},
				&common.Float{Value: float64(l) * r})
		}
	}

	// int / float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			if r != 0 {
				testBinaryOp(t, &common.Int{Value: l}, token.Quo,
					&common.Float{Value: r},
					&common.Float{Value: float64(l) / r})
			}
		}
	}

	// int < float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &common.Int{Value: l}, token.Less,
				&common.Float{Value: r}, boolValue(float64(l) < r))
		}
	}

	// int > float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &common.Int{Value: l}, token.Greater,
				&common.Float{Value: r}, boolValue(float64(l) > r))
		}
	}

	// int <= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &common.Int{Value: l}, token.LessEq,
				&common.Float{Value: r}, boolValue(float64(l) <= r))
		}
	}

	// int >= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &common.Int{Value: l}, token.GreaterEq,
				&common.Float{Value: r}, boolValue(float64(l) >= r))
		}
	}
}

func TestMap_Index(t *testing.T) {
	m := &common.Map{Value: make(map[string]common.Object)}
	k := &common.Int{Value: 1}
	v := &common.String{Value: "abcdef"}
	err := m.IndexSet(k, v)

	require.NoError(t, err)

	res, err := m.IndexGet(k)
	require.NoError(t, err)
	require.Equal(t, v, res)
}

func TestString_BinaryOp(t *testing.T) {
	lstr := "abcde"
	rstr := "01234"
	for l := 0; l < len(lstr); l++ {
		for r := 0; r < len(rstr); r++ {
			ls := lstr[l:]
			rs := rstr[r:]
			testBinaryOp(t, &common.String{Value: ls}, token.Add,
				&common.String{Value: rs},
				&common.String{Value: ls + rs})

			rc := []rune(rstr)[r]
			testBinaryOp(t, &common.String{Value: ls}, token.Add,
				&common.Char{Value: rc},
				&common.String{Value: ls + string(rc)})
		}
	}
}

func testBinaryOp(
	t *testing.T,
	lhs common.Object,
	op token.Token,
	rhs common.Object,
	expected common.Object,
) {
	t.Helper()
	actual, err := lhs.BinaryOp(op, rhs)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func boolValue(b bool) common.Object {
	if b {
		return common.TrueValue
	}
	return common.FalseValue
}
