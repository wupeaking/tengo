package stdlib

import (
	"os"

	"github.com/d5/tengo/v2/common"
)

func makeOSFile(file *os.File) *common.ImmutableMap {
	return &common.ImmutableMap{
		Value: map[string]common.Object{
			// chdir() => true/error
			"chdir": &common.UserFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &common.UserFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &common.UserFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &common.UserFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &common.UserFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &common.UserFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &common.UserFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &common.UserFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &common.UserFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &common.UserFunction{
				Name: "chmod",
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
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &common.UserFunction{
				Name: "seek",
				Value: func(args ...common.Object) (common.Object, error) {
					if len(args) != 2 {
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
					i2, ok := common.ToInt(args[1])
					if !ok {
						return nil, common.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &common.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &common.UserFunction{
				Name: "stat",
				Value: func(args ...common.Object) (common.Object, error) {
					if len(args) != 0 {
						return nil, common.ErrWrongNumArguments
					}
					return osStat(&common.String{Value: file.Name()})
				},
			},
		},
	}
}
