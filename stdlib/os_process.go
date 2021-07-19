package stdlib

import (
	"os"
	"syscall"

	"github.com/d5/tengo/v2/common"
)

func makeOSProcessState(state *os.ProcessState) *common.ImmutableMap {
	return &common.ImmutableMap{
		Value: map[string]common.Object{
			"exited": &common.UserFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &common.UserFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &common.UserFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &common.UserFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *common.ImmutableMap {
	return &common.ImmutableMap{
		Value: map[string]common.Object{
			"kill": &common.UserFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &common.UserFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &common.UserFunction{
				Name: "signal",
				Value: func(args ...common.Object) (common.Object, error) {
					if len(args) != 1 {
						return nil, common.ErrWrongNumArguments
					}
					i1, ok := common.ToInt64(args[0])
					if !ok {
						return nil, common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					return wrapError(proc.Signal(syscall.Signal(i1))), nil
				},
			},
			"wait": &common.UserFunction{
				Name: "wait",
				Value: func(args ...common.Object) (common.Object, error) {
					if len(args) != 0 {
						return nil, common.ErrWrongNumArguments
					}
					state, err := proc.Wait()
					if err != nil {
						return wrapError(err), nil
					}
					return makeOSProcessState(state), nil
				},
			},
		},
	}
}
