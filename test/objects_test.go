package complier

import (
	"testing"

	"github.com/d5/tengo/v2/require"
	"github.com/d5/tengo/v2/token"
)

func TestObject_TypeName(t *testing.T) {
	var o Object = &Int{}
	require.Equal(t, "int", o.TypeName())
	o = &Float{}
	require.Equal(t, "float", o.TypeName())
	o = &Char{}
	require.Equal(t, "char", o.TypeName())
	o = &String{}
	require.Equal(t, "string", o.TypeName())
	o = &Bool{}
	require.Equal(t, "bool", o.TypeName())
	o = &Array{}
	require.Equal(t, "array", o.TypeName())
	o = &Map{}
	require.Equal(t, "map", o.TypeName())
	o = &ArrayIterator{}
	require.Equal(t, "array-iterator", o.TypeName())
	o = &StringIterator{}
	require.Equal(t, "string-iterator", o.TypeName())
	o = &MapIterator{}
	require.Equal(t, "map-iterator", o.TypeName())
	o = &BuiltinFunction{Name: "fn"}
	require.Equal(t, "builtin-function:fn", o.TypeName())
	o = &UserFunction{Name: "fn"}
	require.Equal(t, "user-function:fn", o.TypeName())
	o = &CompiledFunction{}
	require.Equal(t, "compiled-function", o.TypeName())
	o = &Undefined{}
	require.Equal(t, "undefined", o.TypeName())
	o = &Error{}
	require.Equal(t, "error", o.TypeName())
	o = &Bytes{}
	require.Equal(t, "bytes", o.TypeName())
}

func TestObject_IsFalsy(t *testing.T) {
	var o Object = &Int{Value: 0}
	require.True(t, o.IsFalsy())
	o = &Int{Value: 1}
	require.False(t, o.IsFalsy())
	o = &Float{Value: 0}
	require.False(t, o.IsFalsy())
	o = &Float{Value: 1}
	require.False(t, o.IsFalsy())
	o = &Char{Value: ' '}
	require.False(t, o.IsFalsy())
	o = &Char{Value: 'T'}
	require.False(t, o.IsFalsy())
	o = &String{Value: ""}
	require.True(t, o.IsFalsy())
	o = &String{Value: " "}
	require.False(t, o.IsFalsy())
	o = &Array{Value: nil}
	require.True(t, o.IsFalsy())
	o = &Array{Value: []Object{nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &Map{Value: nil}
	require.True(t, o.IsFalsy())
	o = &Map{Value: map[string]Object{"a": nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &StringIterator{}
	require.True(t, o.IsFalsy())
	o = &ArrayIterator{}
	require.True(t, o.IsFalsy())
	o = &MapIterator{}
	require.True(t, o.IsFalsy())
	o = &BuiltinFunction{}
	require.False(t, o.IsFalsy())
	o = &CompiledFunction{}
	require.False(t, o.IsFalsy())
	o = &Undefined{}
	require.True(t, o.IsFalsy())
	o = &Error{}
	require.True(t, o.IsFalsy())
	o = &Bytes{}
	require.True(t, o.IsFalsy())
	o = &Bytes{Value: []byte{1, 2}}
	require.False(t, o.IsFalsy())
}

func TestObject_String(t *testing.T) {
	var o Object = &Int{Value: 0}
	require.Equal(t, "0", o.String())
	o = &Int{Value: 1}
	require.Equal(t, "1", o.String())
	o = &Float{Value: 0}
	require.Equal(t, "0", o.String())
	o = &Float{Value: 1}
	require.Equal(t, "1", o.String())
	o = &Char{Value: ' '}
	require.Equal(t, " ", o.String())
	o = &Char{Value: 'T'}
	require.Equal(t, "T", o.String())
	o = &String{Value: ""}
	require.Equal(t, `""`, o.String())
	o = &String{Value: " "}
	require.Equal(t, `" "`, o.String())
	o = &Array{Value: nil}
	require.Equal(t, "[]", o.String())
	o = &Map{Value: nil}
	require.Equal(t, "{}", o.String())
	o = &Error{Value: nil}
	require.Equal(t, "error", o.String())
	o = &Error{Value: &String{Value: "error 1"}}
	require.Equal(t, `error: "error 1"`, o.String())
	o = &StringIterator{}
	require.Equal(t, "<string-iterator>", o.String())
	o = &ArrayIterator{}
	require.Equal(t, "<array-iterator>", o.String())
	o = &MapIterator{}
	require.Equal(t, "<map-iterator>", o.String())
	o = &Undefined{}
	require.Equal(t, "<undefined>", o.String())
	o = &Bytes{}
	require.Equal(t, "", o.String())
	o = &Bytes{Value: []byte("foo")}
	require.Equal(t, "foo", o.String())
}

func TestObject_BinaryOp(t *testing.T) {
	var o Object = &Char{}
	_, err := o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &Bool{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &Map{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &ArrayIterator{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &StringIterator{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &MapIterator{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &BuiltinFunction{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &CompiledFunction{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &Undefined{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
	o = &Error{}
	_, err = o.BinaryOp(token.Add, UndefinedValue)
	require.Error(t, err)
}

func TestArray_BinaryOp(t *testing.T) {
	testBinaryOp(t, &Array{Value: nil}, token.Add,
		&Array{Value: nil}, &Array{Value: nil})
	testBinaryOp(t, &Array{Value: nil}, token.Add,
		&Array{Value: []Object{}}, &Array{Value: nil})
	testBinaryOp(t, &Array{Value: []Object{}}, token.Add,
		&Array{Value: nil}, &Array{Value: []Object{}})
	testBinaryOp(t, &Array{Value: []Object{}}, token.Add,
		&Array{Value: []Object{}},
		&Array{Value: []Object{}})
	testBinaryOp(t, &Array{Value: nil}, token.Add,
		&Array{Value: []Object{
			&Int{Value: 1},
		}}, &Array{Value: []Object{
			&Int{Value: 1},
		}})
	testBinaryOp(t, &Array{Value: nil}, token.Add,
		&Array{Value: []Object{
			&Int{Value: 1},
			&Int{Value: 2},
			&Int{Value: 3},
		}}, &Array{Value: []Object{
			&Int{Value: 1},
			&Int{Value: 2},
			&Int{Value: 3},
		}})
	testBinaryOp(t, &Array{Value: []Object{
		&Int{Value: 1},
		&Int{Value: 2},
		&Int{Value: 3},
	}}, token.Add, &Array{Value: nil},
		&Array{Value: []Object{
			&Int{Value: 1},
			&Int{Value: 2},
			&Int{Value: 3},
		}})
	testBinaryOp(t, &Array{Value: []Object{
		&Int{Value: 1},
		&Int{Value: 2},
		&Int{Value: 3},
	}}, token.Add, &Array{Value: []Object{
		&Int{Value: 4},
		&Int{Value: 5},
		&Int{Value: 6},
	}}, &Array{Value: []Object{
		&Int{Value: 1},
		&Int{Value: 2},
		&Int{Value: 3},
		&Int{Value: 4},
		&Int{Value: 5},
		&Int{Value: 6},
	}})
}

func TestError_Equals(t *testing.T) {
	err1 := &Error{Value: &String{Value: "some error"}}
	err2 := err1
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = &Error{Value: &String{Value: "some error"}}
	require.False(t, err1.Equals(err2))
	require.False(t, err2.Equals(err1))
}

func TestFloat_BinaryOp(t *testing.T) {
	// float + float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &Float{Value: l}, token.Add,
				&Float{Value: r}, &Float{Value: l + r})
		}
	}

	// float - float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &Float{Value: l}, token.Sub,
				&Float{Value: r}, &Float{Value: l - r})
		}
	}

	// float * float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &Float{Value: l}, token.Mul,
				&Float{Value: r}, &Float{Value: l * r})
		}
	}

	// float / float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			if r != 0 {
				testBinaryOp(t, &Float{Value: l}, token.Quo,
					&Float{Value: r}, &Float{Value: l / r})
			}
		}
	}

	// float < float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &Float{Value: l}, token.Less,
				&Float{Value: r}, boolValue(l < r))
		}
	}

	// float > float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &Float{Value: l}, token.Greater,
				&Float{Value: r}, boolValue(l > r))
		}
	}

	// float <= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &Float{Value: l}, token.LessEq,
				&Float{Value: r}, boolValue(l <= r))
		}
	}

	// float >= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &Float{Value: l}, token.GreaterEq,
				&Float{Value: r}, boolValue(l >= r))
		}
	}

	// float + int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Float{Value: l}, token.Add,
				&Int{Value: r}, &Float{Value: l + float64(r)})
		}
	}

	// float - int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Float{Value: l}, token.Sub,
				&Int{Value: r}, &Float{Value: l - float64(r)})
		}
	}

	// float * int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Float{Value: l}, token.Mul,
				&Int{Value: r}, &Float{Value: l * float64(r)})
		}
	}

	// float / int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &Float{Value: l}, token.Quo,
					&Int{Value: r},
					&Float{Value: l / float64(r)})
			}
		}
	}

	// float < int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Float{Value: l}, token.Less,
				&Int{Value: r}, boolValue(l < float64(r)))
		}
	}

	// float > int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Float{Value: l}, token.Greater,
				&Int{Value: r}, boolValue(l > float64(r)))
		}
	}

	// float <= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Float{Value: l}, token.LessEq,
				&Int{Value: r}, boolValue(l <= float64(r)))
		}
	}

	// float >= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Float{Value: l}, token.GreaterEq,
				&Int{Value: r}, boolValue(l >= float64(r)))
		}
	}
}

func TestInt_BinaryOp(t *testing.T) {
	// int + int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Int{Value: l}, token.Add,
				&Int{Value: r}, &Int{Value: l + r})
		}
	}

	// int - int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Int{Value: l}, token.Sub,
				&Int{Value: r}, &Int{Value: l - r})
		}
	}

	// int * int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Int{Value: l}, token.Mul,
				&Int{Value: r}, &Int{Value: l * r})
		}
	}

	// int / int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &Int{Value: l}, token.Quo,
					&Int{Value: r}, &Int{Value: l / r})
			}
		}
	}

	// int % int
	for l := int64(-4); l <= 4; l++ {
		for r := -int64(-4); r <= 4; r++ {
			if r == 0 {
				testBinaryOp(t, &Int{Value: l}, token.Rem,
					&Int{Value: r}, &Int{Value: l % r})
			}
		}
	}

	// int & int
	testBinaryOp(t,
		&Int{Value: 0}, token.And, &Int{Value: 0},
		&Int{Value: int64(0)})
	testBinaryOp(t,
		&Int{Value: 1}, token.And, &Int{Value: 0},
		&Int{Value: int64(1) & int64(0)})
	testBinaryOp(t,
		&Int{Value: 0}, token.And, &Int{Value: 1},
		&Int{Value: int64(0) & int64(1)})
	testBinaryOp(t,
		&Int{Value: 1}, token.And, &Int{Value: 1},
		&Int{Value: int64(1)})
	testBinaryOp(t,
		&Int{Value: 0}, token.And, &Int{Value: int64(0xffffffff)},
		&Int{Value: int64(0) & int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: 1}, token.And, &Int{Value: int64(0xffffffff)},
		&Int{Value: int64(1) & int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: int64(0xffffffff)}, token.And,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: 1984}, token.And,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(1984) & int64(0xffffffff)})
	testBinaryOp(t, &Int{Value: -1984}, token.And,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(-1984) & int64(0xffffffff)})

	// int | int
	testBinaryOp(t,
		&Int{Value: 0}, token.Or, &Int{Value: 0},
		&Int{Value: int64(0)})
	testBinaryOp(t,
		&Int{Value: 1}, token.Or, &Int{Value: 0},
		&Int{Value: int64(1) | int64(0)})
	testBinaryOp(t,
		&Int{Value: 0}, token.Or, &Int{Value: 1},
		&Int{Value: int64(0) | int64(1)})
	testBinaryOp(t,
		&Int{Value: 1}, token.Or, &Int{Value: 1},
		&Int{Value: int64(1)})
	testBinaryOp(t,
		&Int{Value: 0}, token.Or, &Int{Value: int64(0xffffffff)},
		&Int{Value: int64(0) | int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: 1}, token.Or, &Int{Value: int64(0xffffffff)},
		&Int{Value: int64(1) | int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: int64(0xffffffff)}, token.Or,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: 1984}, token.Or,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(1984) | int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: -1984}, token.Or,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(-1984) | int64(0xffffffff)})

	// int ^ int
	testBinaryOp(t,
		&Int{Value: 0}, token.Xor, &Int{Value: 0},
		&Int{Value: int64(0)})
	testBinaryOp(t,
		&Int{Value: 1}, token.Xor, &Int{Value: 0},
		&Int{Value: int64(1) ^ int64(0)})
	testBinaryOp(t,
		&Int{Value: 0}, token.Xor, &Int{Value: 1},
		&Int{Value: int64(0) ^ int64(1)})
	testBinaryOp(t,
		&Int{Value: 1}, token.Xor, &Int{Value: 1},
		&Int{Value: int64(0)})
	testBinaryOp(t,
		&Int{Value: 0}, token.Xor, &Int{Value: int64(0xffffffff)},
		&Int{Value: int64(0) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: 1}, token.Xor, &Int{Value: int64(0xffffffff)},
		&Int{Value: int64(1) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: int64(0xffffffff)}, token.Xor,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(0)})
	testBinaryOp(t,
		&Int{Value: 1984}, token.Xor,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(1984) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: -1984}, token.Xor,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(-1984) ^ int64(0xffffffff)})

	// int &^ int
	testBinaryOp(t,
		&Int{Value: 0}, token.AndNot, &Int{Value: 0},
		&Int{Value: int64(0)})
	testBinaryOp(t,
		&Int{Value: 1}, token.AndNot, &Int{Value: 0},
		&Int{Value: int64(1) &^ int64(0)})
	testBinaryOp(t,
		&Int{Value: 0}, token.AndNot,
		&Int{Value: 1}, &Int{Value: int64(0) &^ int64(1)})
	testBinaryOp(t,
		&Int{Value: 1}, token.AndNot, &Int{Value: 1},
		&Int{Value: int64(0)})
	testBinaryOp(t,
		&Int{Value: 0}, token.AndNot,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(0) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: 1}, token.AndNot,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(1) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: int64(0xffffffff)}, token.AndNot,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(0)})
	testBinaryOp(t,
		&Int{Value: 1984}, token.AndNot,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(1984) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&Int{Value: -1984}, token.AndNot,
		&Int{Value: int64(0xffffffff)},
		&Int{Value: int64(-1984) &^ int64(0xffffffff)})

	// int << int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&Int{Value: 0}, token.Shl, &Int{Value: s},
			&Int{Value: int64(0) << uint(s)})
		testBinaryOp(t,
			&Int{Value: 1}, token.Shl, &Int{Value: s},
			&Int{Value: int64(1) << uint(s)})
		testBinaryOp(t,
			&Int{Value: 2}, token.Shl, &Int{Value: s},
			&Int{Value: int64(2) << uint(s)})
		testBinaryOp(t,
			&Int{Value: -1}, token.Shl, &Int{Value: s},
			&Int{Value: int64(-1) << uint(s)})
		testBinaryOp(t,
			&Int{Value: -2}, token.Shl, &Int{Value: s},
			&Int{Value: int64(-2) << uint(s)})
		testBinaryOp(t,
			&Int{Value: int64(0xffffffff)}, token.Shl,
			&Int{Value: s},
			&Int{Value: int64(0xffffffff) << uint(s)})
	}

	// int >> int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&Int{Value: 0}, token.Shr, &Int{Value: s},
			&Int{Value: int64(0) >> uint(s)})
		testBinaryOp(t,
			&Int{Value: 1}, token.Shr, &Int{Value: s},
			&Int{Value: int64(1) >> uint(s)})
		testBinaryOp(t,
			&Int{Value: 2}, token.Shr, &Int{Value: s},
			&Int{Value: int64(2) >> uint(s)})
		testBinaryOp(t,
			&Int{Value: -1}, token.Shr, &Int{Value: s},
			&Int{Value: int64(-1) >> uint(s)})
		testBinaryOp(t,
			&Int{Value: -2}, token.Shr, &Int{Value: s},
			&Int{Value: int64(-2) >> uint(s)})
		testBinaryOp(t,
			&Int{Value: int64(0xffffffff)}, token.Shr,
			&Int{Value: s},
			&Int{Value: int64(0xffffffff) >> uint(s)})
	}

	// int < int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Int{Value: l}, token.Less,
				&Int{Value: r}, boolValue(l < r))
		}
	}

	// int > int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Int{Value: l}, token.Greater,
				&Int{Value: r}, boolValue(l > r))
		}
	}

	// int <= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Int{Value: l}, token.LessEq,
				&Int{Value: r}, boolValue(l <= r))
		}
	}

	// int >= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &Int{Value: l}, token.GreaterEq,
				&Int{Value: r}, boolValue(l >= r))
		}
	}

	// int + float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &Int{Value: l}, token.Add,
				&Float{Value: r},
				&Float{Value: float64(l) + r})
		}
	}

	// int - float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &Int{Value: l}, token.Sub,
				&Float{Value: r},
				&Float{Value: float64(l) - r})
		}
	}

	// int * float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &Int{Value: l}, token.Mul,
				&Float{Value: r},
				&Float{Value: float64(l) * r})
		}
	}

	// int / float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			if r != 0 {
				testBinaryOp(t, &Int{Value: l}, token.Quo,
					&Float{Value: r},
					&Float{Value: float64(l) / r})
			}
		}
	}

	// int < float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &Int{Value: l}, token.Less,
				&Float{Value: r}, boolValue(float64(l) < r))
		}
	}

	// int > float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &Int{Value: l}, token.Greater,
				&Float{Value: r}, boolValue(float64(l) > r))
		}
	}

	// int <= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &Int{Value: l}, token.LessEq,
				&Float{Value: r}, boolValue(float64(l) <= r))
		}
	}

	// int >= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &Int{Value: l}, token.GreaterEq,
				&Float{Value: r}, boolValue(float64(l) >= r))
		}
	}
}

func TestMap_Index(t *testing.T) {
	m := &Map{Value: make(map[string]Object)}
	k := &Int{Value: 1}
	v := &String{Value: "abcdef"}
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
			testBinaryOp(t, &String{Value: ls}, token.Add,
				&String{Value: rs},
				&String{Value: ls + rs})

			rc := []rune(rstr)[r]
			testBinaryOp(t, &String{Value: ls}, token.Add,
				&Char{Value: rc},
				&String{Value: ls + string(rc)})
		}
	}
}

func testBinaryOp(
	t *testing.T,
	lhs Object,
	op token.Token,
	rhs Object,
	expected Object,
) {
	t.Helper()
	actual, err := lhs.BinaryOp(op, rhs)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func boolValue(b bool) Object {
	if b {
		return TrueValue
	}
	return FalseValue
}
