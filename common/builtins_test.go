package common_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/d5/tengo/v2/common"
)

func Test_builtinDelete(t *testing.T) {
	var builtinDelete func(args ...common.Object) (common.Object, error)
	for _, f := range common.GetAllBuiltinFunctions() {
		if f.Name == "delete" {
			builtinDelete = f.Value
			break
		}
	}
	if builtinDelete == nil {
		t.Fatal("builtin delete not found")
	}
	type args struct {
		args []common.Object
	}
	tests := []struct {
		name      string
		args      args
		want      common.Object
		wantErr   bool
		wantedErr error
		target    interface{}
	}{
		{name: "invalid-arg", args: args{[]common.Object{&common.String{},
			&common.String{}}}, wantErr: true,
			wantedErr: common.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "map",
				Found:    "string"},
		},
		{name: "no-args",
			wantErr: true, wantedErr: common.ErrWrongNumArguments},
		{name: "empty-args", args: args{[]common.Object{}}, wantErr: true,
			wantedErr: common.ErrWrongNumArguments,
		},
		{name: "3-args", args: args{[]common.Object{
			(*common.Map)(nil), (*common.String)(nil), (*common.String)(nil)}},
			wantErr: true, wantedErr: common.ErrWrongNumArguments,
		},
		{name: "nil-map-empty-key",
			args: args{[]common.Object{&common.Map{}, &common.String{}}},
			want: common.UndefinedValue,
		},
		{name: "nil-map-nonstr-key",
			args: args{[]common.Object{
				&common.Map{}, &common.Int{}}}, wantErr: true,
			wantedErr: common.ErrInvalidArgumentType{
				Name: "second", Expected: "string", Found: "int"},
		},
		{name: "nil-map-no-key",
			args: args{[]common.Object{&common.Map{}}}, wantErr: true,
			wantedErr: common.ErrWrongNumArguments,
		},
		{name: "map-missing-key",
			args: args{
				[]common.Object{
					&common.Map{Value: map[string]common.Object{
						"key": &common.String{Value: "value"},
					}},
					&common.String{Value: "key1"}}},
			want: common.UndefinedValue,
			target: &common.Map{
				Value: map[string]common.Object{
					"key": &common.String{
						Value: "value"}}},
		},
		{name: "map-emptied",
			args: args{
				[]common.Object{
					&common.Map{Value: map[string]common.Object{
						"key": &common.String{Value: "value"},
					}},
					&common.String{Value: "key"}}},
			want:   common.UndefinedValue,
			target: &common.Map{Value: map[string]common.Object{}},
		},
		{name: "map-multi-keys",
			args: args{
				[]common.Object{
					&common.Map{Value: map[string]common.Object{
						"key1": &common.String{Value: "value1"},
						"key2": &common.Int{Value: 10},
					}},
					&common.String{Value: "key1"}}},
			want: common.UndefinedValue,
			target: &common.Map{Value: map[string]common.Object{
				"key2": &common.Int{Value: 10}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinDelete(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinDelete() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.wantedErr) {
				if err.Error() != tt.wantedErr.Error() {
					t.Errorf("builtinDelete() error = %v, wantedErr %v",
						err, tt.wantedErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("builtinDelete() = %v, want %v", got, tt.want)
				return
			}
			if !tt.wantErr && tt.target != nil {
				switch v := tt.args.args[0].(type) {
				case *common.Map, *common.Array:
					if !reflect.DeepEqual(tt.target, tt.args.args[0]) {
						t.Errorf("builtinDelete() common.Objects are not equal "+
							"got: %+v, want: %+v", tt.args.args[0], tt.target)
					}
				default:
					t.Errorf("builtinDelete() unsuporrted arg[0] type %s",
						v.TypeName())
					return
				}
			}
		})
	}
}

func Test_builtinSplice(t *testing.T) {
	var builtinSplice func(args ...common.Object) (common.Object, error)
	for _, f := range common.GetAllBuiltinFunctions() {
		if f.Name == "splice" {
			builtinSplice = f.Value
			break
		}
	}
	if builtinSplice == nil {
		t.Fatal("builtin splice not found")
	}
	tests := []struct {
		name      string
		args      []common.Object
		deleted   common.Object
		Array     *common.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []common.Object{}, wantErr: true,
			wantedErr: common.ErrWrongNumArguments,
		},
		{name: "invalid args", args: []common.Object{&common.Map{}},
			wantErr: true,
			wantedErr: common.ErrInvalidArgumentType{
				Name: "first", Expected: "array", Found: "map"},
		},
		{name: "invalid args",
			args:    []common.Object{&common.Array{}, &common.String{}},
			wantErr: true,
			wantedErr: common.ErrInvalidArgumentType{
				Name: "second", Expected: "int", Found: "string"},
		},
		{name: "negative index",
			args:      []common.Object{&common.Array{}, &common.Int{Value: -1}},
			wantErr:   true,
			wantedErr: common.ErrIndexOutOfBounds},
		{name: "non int count",
			args: []common.Object{
				&common.Array{}, &common.Int{Value: 0},
				&common.String{Value: ""}},
			wantErr: true,
			wantedErr: common.ErrInvalidArgumentType{
				Name: "third", Expected: "int", Found: "string"},
		},
		{name: "negative count",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 0},
				&common.Int{Value: -1}},
			wantErr:   true,
			wantedErr: common.ErrIndexOutOfBounds,
		},
		{name: "insert with zero count",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 0},
				&common.Int{Value: 0},
				&common.String{Value: "b"}},
			deleted: &common.Array{Value: []common.Object{}},
			Array: &common.Array{Value: []common.Object{
				&common.String{Value: "b"},
				&common.Int{Value: 0},
				&common.Int{Value: 1},
				&common.Int{Value: 2}}},
		},
		{name: "insert",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 1},
				&common.Int{Value: 0},
				&common.String{Value: "c"},
				&common.String{Value: "d"}},
			deleted: &common.Array{Value: []common.Object{}},
			Array: &common.Array{Value: []common.Object{
				&common.Int{Value: 0},
				&common.String{Value: "c"},
				&common.String{Value: "d"},
				&common.Int{Value: 1},
				&common.Int{Value: 2}}},
		},
		{name: "insert with zero count",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 1},
				&common.Int{Value: 0},
				&common.String{Value: "c"},
				&common.String{Value: "d"}},
			deleted: &common.Array{Value: []common.Object{}},
			Array: &common.Array{Value: []common.Object{
				&common.Int{Value: 0},
				&common.String{Value: "c"},
				&common.String{Value: "d"},
				&common.Int{Value: 1},
				&common.Int{Value: 2}}},
		},
		{name: "insert with delete",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 1},
				&common.Int{Value: 1},
				&common.String{Value: "c"},
				&common.String{Value: "d"}},
			deleted: &common.Array{
				Value: []common.Object{&common.Int{Value: 1}}},
			Array: &common.Array{Value: []common.Object{
				&common.Int{Value: 0},
				&common.String{Value: "c"},
				&common.String{Value: "d"},
				&common.Int{Value: 2}}},
		},
		{name: "insert with delete multi",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 1},
				&common.Int{Value: 2},
				&common.String{Value: "c"},
				&common.String{Value: "d"}},
			deleted: &common.Array{Value: []common.Object{
				&common.Int{Value: 1},
				&common.Int{Value: 2}}},
			Array: &common.Array{
				Value: []common.Object{
					&common.Int{Value: 0},
					&common.String{Value: "c"},
					&common.String{Value: "d"}}},
		},
		{name: "delete all with positive count",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 0},
				&common.Int{Value: 3}},
			deleted: &common.Array{Value: []common.Object{
				&common.Int{Value: 0},
				&common.Int{Value: 1},
				&common.Int{Value: 2}}},
			Array: &common.Array{Value: []common.Object{}},
		},
		{name: "delete all with big count",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 0},
				&common.Int{Value: 5}},
			deleted: &common.Array{Value: []common.Object{
				&common.Int{Value: 0},
				&common.Int{Value: 1},
				&common.Int{Value: 2}}},
			Array: &common.Array{Value: []common.Object{}},
		},
		{name: "nothing2",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}}},
			Array: &common.Array{Value: []common.Object{}},
			deleted: &common.Array{Value: []common.Object{
				&common.Int{Value: 0},
				&common.Int{Value: 1},
				&common.Int{Value: 2}}},
		},
		{name: "pop without count",
			args: []common.Object{
				&common.Array{Value: []common.Object{
					&common.Int{Value: 0},
					&common.Int{Value: 1},
					&common.Int{Value: 2}}},
				&common.Int{Value: 2}},
			deleted: &common.Array{Value: []common.Object{&common.Int{Value: 2}}},
			Array: &common.Array{Value: []common.Object{
				&common.Int{Value: 0}, &common.Int{Value: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinSplice(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinSplice() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.deleted) {
				t.Errorf("builtinSplice() = %v, want %v", got, tt.deleted)
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinSplice() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.Array != nil && !reflect.DeepEqual(tt.Array, tt.args[0]) {
				t.Errorf("builtinSplice() arrays are not equal expected"+
					" %s, got %s", tt.Array, tt.args[0].(*common.Array))
			}
		})
	}
}

func Test_builtinRange(t *testing.T) {
	var builtinRange func(args ...common.Object) (common.Object, error)
	for _, f := range common.GetAllBuiltinFunctions() {
		if f.Name == "range" {
			builtinRange = f.Value
			break
		}
	}
	if builtinRange == nil {
		t.Fatal("builtin range not found")
	}
	tests := []struct {
		name      string
		args      []common.Object
		result    *common.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []common.Object{}, wantErr: true,
			wantedErr: common.ErrWrongNumArguments,
		},
		{name: "single args", args: []common.Object{&common.Map{}},
			wantErr:   true,
			wantedErr: common.ErrWrongNumArguments,
		},
		{name: "4 args", args: []common.Object{&common.Map{}, &common.String{}, &common.String{}, &common.String{}},
			wantErr:   true,
			wantedErr: common.ErrWrongNumArguments,
		},
		{name: "invalid start",
			args:    []common.Object{&common.String{}, &common.String{}},
			wantErr: true,
			wantedErr: common.ErrInvalidArgumentType{
				Name: "start", Expected: "int", Found: "string"},
		},
		{name: "invalid stop",
			args:    []common.Object{&common.Int{}, &common.String{}},
			wantErr: true,
			wantedErr: common.ErrInvalidArgumentType{
				Name: "stop", Expected: "int", Found: "string"},
		},
		{name: "invalid step",
			args:    []common.Object{&common.Int{}, &common.Int{}, &common.String{}},
			wantErr: true,
			wantedErr: common.ErrInvalidArgumentType{
				Name: "step", Expected: "int", Found: "string"},
		},
		{name: "zero step",
			args:      []common.Object{&common.Int{}, &common.Int{}, &common.Int{}}, //must greate than 0
			wantErr:   true,
			wantedErr: common.ErrInvalidRangeStep,
		},
		{name: "negative step",
			args:      []common.Object{&common.Int{}, &common.Int{}, intObject(-2)}, //must greate than 0
			wantErr:   true,
			wantedErr: common.ErrInvalidRangeStep,
		},
		{name: "same bound",
			args:    []common.Object{&common.Int{}, &common.Int{}},
			wantErr: false,
			result: &common.Array{
				Value: nil,
			},
		},
		{name: "positive range",
			args:    []common.Object{&common.Int{}, &common.Int{Value: 5}},
			wantErr: false,
			result: &common.Array{
				Value: []common.Object{
					intObject(0),
					intObject(1),
					intObject(2),
					intObject(3),
					intObject(4),
				},
			},
		},
		{name: "negative range",
			args:    []common.Object{&common.Int{}, &common.Int{Value: -5}},
			wantErr: false,
			result: &common.Array{
				Value: []common.Object{
					intObject(0),
					intObject(-1),
					intObject(-2),
					intObject(-3),
					intObject(-4),
				},
			},
		},

		{name: "positive with step",
			args:    []common.Object{&common.Int{}, &common.Int{Value: 5}, &common.Int{Value: 2}},
			wantErr: false,
			result: &common.Array{
				Value: []common.Object{
					intObject(0),
					intObject(2),
					intObject(4),
				},
			},
		},
		{name: "negative with step",
			args:    []common.Object{&common.Int{}, &common.Int{Value: -10}, &common.Int{Value: 2}},
			wantErr: false,
			result: &common.Array{
				Value: []common.Object{
					intObject(0),
					intObject(-2),
					intObject(-4),
					intObject(-6),
					intObject(-8),
				},
			},
		},
		{name: "positive with step",
			args:    []common.Object{&common.Int{}, &common.Int{Value: 5}, &common.Int{Value: 2}},
			wantErr: false,
			result: &common.Array{
				Value: []common.Object{
					intObject(0),
					intObject(2),
					intObject(4),
				},
			},
		},
		{name: "negative with step",
			args:    []common.Object{&common.Int{}, &common.Int{Value: -10}, &common.Int{Value: 2}},
			wantErr: false,
			result: &common.Array{
				Value: []common.Object{
					intObject(0),
					intObject(-2),
					intObject(-4),
					intObject(-6),
					intObject(-8),
				},
			},
		},
		{name: "large range",
			args:    []common.Object{intObject(-10), intObject(10), &common.Int{Value: 3}},
			wantErr: false,
			result: &common.Array{
				Value: []common.Object{
					intObject(-10),
					intObject(-7),
					intObject(-4),
					intObject(-1),
					intObject(2),
					intObject(5),
					intObject(8),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinRange(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinRange() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinRange() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.result != nil && !reflect.DeepEqual(tt.result, got) {
				t.Errorf("builtinRange() arrays are not equal expected"+
					" %s, got %s", tt.result, got.(*common.Array))
			}
		})
	}
}

func intObject(v int64) *common.Int {
	return &common.Int{Value: v}
}
