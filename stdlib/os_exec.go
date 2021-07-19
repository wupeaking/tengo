package stdlib

import (
	"os/exec"

	"github.com/d5/tengo/v2/common"
)

func makeOSExecCommand(cmd *exec.Cmd) *common.ImmutableMap {
	return &common.ImmutableMap{
		Value: map[string]common.Object{
			// combined_output() => bytes/error
			"combined_output": &common.UserFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &common.UserFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &common.UserFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &common.UserFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &common.UserFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &common.UserFunction{
				Name: "set_path",
				Value: func(args ...common.Object) (common.Object, error) {
					if len(args) != 1 {
						return nil, common.ErrWrongNumArguments
					}
					s1, ok := common.ToString(args[0])
					if !ok {
						return nil, common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Path = s1
					return common.UndefinedValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &common.UserFunction{
				Name: "set_dir",
				Value: func(args ...common.Object) (common.Object, error) {
					if len(args) != 1 {
						return nil, common.ErrWrongNumArguments
					}
					s1, ok := common.ToString(args[0])
					if !ok {
						return nil, common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Dir = s1
					return common.UndefinedValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &common.UserFunction{
				Name: "set_env",
				Value: func(args ...common.Object) (common.Object, error) {
					if len(args) != 1 {
						return nil, common.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *common.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *common.ImmutableArray:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					default:
						return nil, common.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "array",
							Found:    arg0.TypeName(),
						}
					}
					cmd.Env = env
					return common.UndefinedValue, nil
				},
			},
			// process() => imap(process)
			"process": &common.UserFunction{
				Name: "process",
				Value: func(args ...common.Object) (common.Object, error) {
					if len(args) != 0 {
						return nil, common.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
