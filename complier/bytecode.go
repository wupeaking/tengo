package complier

import (
	"encoding/gob"
	"fmt"
	"io"
	"reflect"

	"github.com/d5/tengo/v2/common"
	"github.com/d5/tengo/v2/parser"
)

// Bytecode is a compiled instructions and constants.
type Bytecode struct {
	FileSet      *parser.SourceFileSet
	MainFunction *common.CompiledFunction
	Constants    []common.Object
}

// Encode writes Bytecode data to the writer.
func (b *Bytecode) Encode(w io.Writer) error {
	enc := gob.NewEncoder(w)
	if err := enc.Encode(b.FileSet); err != nil {
		return err
	}
	if err := enc.Encode(b.MainFunction); err != nil {
		return err
	}
	return enc.Encode(b.Constants)
}

// CountObjects returns the number of objects found in Constants.
func (b *Bytecode) CountObjects() int {
	n := 0
	for _, c := range b.Constants {
		n += common.CountObjects(c)
	}
	return n
}

// FormatInstructions returns human readable string representations of
// compiled instructions.
func (b *Bytecode) FormatInstructions() []string {
	return FormatInstructions(b.MainFunction.Instructions, 0)
}

// FormatConstants returns human readable string representations of
// compiled constants.
func (b *Bytecode) FormatConstants() (output []string) {
	for cidx, cn := range b.Constants {
		switch cn := cn.(type) {
		case *common.CompiledFunction:
			output = append(output, fmt.Sprintf(
				"[% 3d] (Compiled Function|%p)", cidx, &cn))
			for _, l := range FormatInstructions(cn.Instructions, 0) {
				output = append(output, fmt.Sprintf("     %s", l))
			}
		default:
			output = append(output, fmt.Sprintf("[% 3d] %s (%s|%p)",
				cidx, cn, reflect.TypeOf(cn).Elem().Name(), &cn))
		}
	}
	return
}

// Decode reads Bytecode data from the reader.
func (b *Bytecode) Decode(r io.Reader, modules *common.ModuleMap) error {
	if modules == nil {
		modules = common.NewModuleMap()
	}

	dec := gob.NewDecoder(r)
	if err := dec.Decode(&b.FileSet); err != nil {
		return err
	}
	// TODO: files in b.FileSet.File does not have their 'set' field properly
	//  set to b.FileSet as it's private field and not serialized by gob
	//  encoder/decoder.
	if err := dec.Decode(&b.MainFunction); err != nil {
		return err
	}
	if err := dec.Decode(&b.Constants); err != nil {
		return err
	}
	for i, v := range b.Constants {
		fv, err := fixDecodedObject(v, modules)
		if err != nil {
			return err
		}
		b.Constants[i] = fv
	}
	return nil
}

// RemoveDuplicates finds and remove the duplicate values in Constants.
// Note this function mutates Bytecode.
func (b *Bytecode) RemoveDuplicates() {
	var deduped []common.Object

	indexMap := make(map[int]int) // mapping from old constant index to new index
	fns := make(map[*common.CompiledFunction]int)
	ints := make(map[int64]int)
	strings := make(map[string]int)
	floats := make(map[float64]int)
	chars := make(map[rune]int)
	immutableMaps := make(map[string]int) // for modules

	for curIdx, c := range b.Constants {
		switch c := c.(type) {
		case *common.CompiledFunction:
			if newIdx, ok := fns[c]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				fns[c] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *common.ImmutableMap:
			modName := inferModuleName(c)
			newIdx, ok := immutableMaps[modName]
			if modName != "" && ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				immutableMaps[modName] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *common.Int:
			if newIdx, ok := ints[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				ints[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *common.String:
			if newIdx, ok := strings[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				strings[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *common.Float:
			if newIdx, ok := floats[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				floats[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *common.Char:
			if newIdx, ok := chars[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				chars[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		default:
			panic(fmt.Errorf("unsupported top-level constant type: %s",
				c.TypeName()))
		}
	}

	// replace with de-duplicated constants
	b.Constants = deduped

	// update CONST instructions with new indexes
	// main function
	updateConstIndexes(b.MainFunction.Instructions, indexMap)
	// other compiled functions in constants
	for _, c := range b.Constants {
		switch c := c.(type) {
		case *common.CompiledFunction:
			updateConstIndexes(c.Instructions, indexMap)
		}
	}
}

func fixDecodedObject(
	o common.Object,
	modules *common.ModuleMap,
) (common.Object, error) {
	switch o := o.(type) {
	case *common.Bool:
		if o.IsFalsy() {
			return common.FalseValue, nil
		}
		return common.TrueValue, nil
	case *common.Undefined:
		return common.UndefinedValue, nil
	case *common.Array:
		for i, v := range o.Value {
			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Value[i] = fv
		}
	case *common.ImmutableArray:
		for i, v := range o.Value {
			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Value[i] = fv
		}
	case *common.Map:
		for k, v := range o.Value {
			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Value[k] = fv
		}
	case *common.ImmutableMap:
		modName := inferModuleName(o)
		if mod := modules.GetBuiltinModule(modName); mod != nil {
			return mod.AsImmutableMap(modName), nil
		}

		for k, v := range o.Value {
			// encoding of user function not supported
			if _, isUserFunction := v.(*common.UserFunction); isUserFunction {
				return nil, fmt.Errorf("user function not decodable")
			}

			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Value[k] = fv
		}
	}
	return o, nil
}

func updateConstIndexes(insts []byte, indexMap map[int]int) {
	i := 0
	for i < len(insts) {
		op := insts[i]
		numOperands := parser.OpcodeOperands[op]
		_, read := parser.ReadOperands(numOperands, insts[i+1:])

		switch op {
		case parser.OpConstant:
			curIdx := int(insts[i+2]) | int(insts[i+1])<<8
			newIdx, ok := indexMap[curIdx]
			if !ok {
				panic(fmt.Errorf("constant index not found: %d", curIdx))
			}
			copy(insts[i:], MakeInstruction(op, newIdx))
		case parser.OpClosure:
			curIdx := int(insts[i+2]) | int(insts[i+1])<<8
			numFree := int(insts[i+3])
			newIdx, ok := indexMap[curIdx]
			if !ok {
				panic(fmt.Errorf("constant index not found: %d", curIdx))
			}
			copy(insts[i:], MakeInstruction(op, newIdx, numFree))
		}

		i += 1 + read
	}
}

func inferModuleName(mod *common.ImmutableMap) string {
	if modName, ok := mod.Value["__module_name__"].(*common.String); ok {
		return modName.Value
	}
	return ""
}

func init() {
	gob.Register(&parser.SourceFileSet{})
	gob.Register(&parser.SourceFile{})
	gob.Register(&common.Array{})
	gob.Register(&common.Bool{})
	gob.Register(&common.Bytes{})
	gob.Register(&common.Char{})
	gob.Register(&common.CompiledFunction{})
	gob.Register(&common.Error{})
	gob.Register(&common.Float{})
	gob.Register(&common.ImmutableArray{})
	gob.Register(&common.ImmutableMap{})
	gob.Register(&common.Int{})
	gob.Register(&common.Map{})
	gob.Register(&common.String{})
	gob.Register(&common.Time{})
	gob.Register(&common.Undefined{})
	gob.Register(&common.UserFunction{})
}
