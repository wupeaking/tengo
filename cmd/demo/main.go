package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/d5/tengo/v2/common"
	"github.com/d5/tengo/v2/complier"
	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/stdlib"
	"github.com/d5/tengo/v2/vm"
)

func main() {
	modules := stdlib.GetModuleMap(stdlib.AllModuleNames()...)
	inputFile := "example.tengo"
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr,
			"Error reading input file: %s\n", err.Error())
		os.Exit(1)
	}
	err = CompileOnly(modules, inputData, inputFile, "")
	if err != nil {
		panic(err)
	}

	byteCodes, err := ioutil.ReadFile("example.tengo.out")
	if err != nil {
		panic(err)
	}
	err = RunCompiled(modules, byteCodes)
	if err != nil {
		panic(err)
	}
}

func CompileOnly(
	modules *common.ModuleMap,
	data []byte,
	inputFile, outputFile string,
) (err error) {
	bytecode, err := compileSrc(modules, data, inputFile)
	if err != nil {
		return
	}

	if outputFile == "" {
		outputFile = inputFile + ".out"
	}

	out, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = out.Close()
		} else {
			err = out.Close()
		}
	}()

	err = bytecode.Encode(out)
	if err != nil {
		return
	}
	fmt.Println(outputFile)

	return
}

func compileSrc(
	modules *common.ModuleMap,
	src []byte,
	inputFile string,
) (*complier.Bytecode, error) {
	fileSet := parser.NewFileSet()
	srcFile := fileSet.AddFile(filepath.Base(inputFile), -1, len(src))

	p := parser.NewParser(srcFile, src, nil)
	file, err := p.ParseFile()
	if err != nil {
		return nil, err
	}

	c := complier.NewCompiler(srcFile, nil, nil, modules, nil)
	c.EnableFileImport(true)

	if err := c.Compile(file); err != nil {
		return nil, err
	}

	bytecode := c.Bytecode()
	bytecode.RemoveDuplicates()
	return bytecode, nil
}

func RunCompiled(modules *common.ModuleMap, data []byte) (err error) {
	bytecode := &complier.Bytecode{}
	err = bytecode.Decode(bytes.NewReader(data), modules)
	if err != nil {
		return
	}

	machine := vm.NewVM(bytecode, nil, -1)
	err = machine.Run()
	return
}
