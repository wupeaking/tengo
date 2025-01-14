package stdlib

import (
	"github.com/d5/tengo/v2/common"
)

//go:generate go run gensrcmods.go

// AllModuleNames returns a list of all default module names.
func AllModuleNames() []string {
	var names []string
	for name := range BuiltinModules {
		names = append(names, name)
	}
	for name := range SourceModules {
		names = append(names, name)
	}
	return names
}

// GetModuleMap returns the module map that includes all modules
// for the given module names.
func GetModuleMap(names ...string) *common.ModuleMap {
	modules := common.NewModuleMap()
	for _, name := range names {
		if mod := BuiltinModules[name]; mod != nil {
			modules.AddBuiltinModule(name, mod)
		}
		if mod := SourceModules[name]; mod != "" {
			modules.AddSourceModule(name, []byte(mod))
		}
	}
	return modules
}
