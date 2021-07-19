package stdlib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/d5/tengo/v2/common"
)

var osModule = map[string]common.Object{
	"o_rdonly":            &common.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &common.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &common.Int{Value: int64(os.O_RDWR)},
	"o_append":            &common.Int{Value: int64(os.O_APPEND)},
	"o_create":            &common.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &common.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &common.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &common.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &common.Int{Value: int64(os.ModeDir)},
	"mode_append":         &common.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &common.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &common.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &common.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &common.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &common.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &common.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &common.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &common.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &common.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &common.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &common.Int{Value: int64(os.ModeType)},
	"mode_perm":           &common.Int{Value: int64(os.ModePerm)},
	"path_separator":      &common.Char{Value: os.PathSeparator},
	"path_list_separator": &common.Char{Value: os.PathListSeparator},
	"dev_null":            &common.String{Value: os.DevNull},
	"seek_set":            &common.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &common.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &common.Int{Value: int64(io.SeekEnd)},
	"args": &common.UserFunction{
		Name:  "args",
		Value: osArgs,
	}, // args() => array(string)
	"chdir": &common.UserFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chown": &common.UserFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	}, // chown(name string, uid int, gid int) => error
	"clearenv": &common.UserFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"environ": &common.UserFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &common.UserFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"expand_env": &common.UserFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &common.UserFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &common.UserFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &common.UserFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &common.UserFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &common.UserFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &common.UserFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &common.UserFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &common.UserFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &common.UserFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &common.UserFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &common.UserFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &common.UserFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &common.UserFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &common.UserFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &common.UserFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &common.UserFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &common.UserFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &common.UserFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &common.UserFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &common.UserFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &common.UserFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &common.UserFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &common.UserFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &common.UserFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &common.UserFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &common.UserFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &common.UserFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &common.UserFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &common.UserFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &common.UserFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &common.UserFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &common.UserFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
}

func osReadFile(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		return nil, common.ErrWrongNumArguments
	}
	fname, ok := common.ToString(args[0])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > common.MaxBytesLen {
		return nil, common.ErrBytesLimit
	}
	return &common.Bytes{Value: bytes}, nil
}

func osStat(args ...common.Object) (ret common.Object, err error) {
	if len(args) != 1 {
		return nil, common.ErrWrongNumArguments
	}
	fname, ok := common.ToString(args[0])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &common.ImmutableMap{
		Value: map[string]common.Object{
			"name":  &common.String{Value: stat.Name()},
			"mtime": &common.Time{Value: stat.ModTime()},
			"size":  &common.Int{Value: stat.Size()},
			"mode":  &common.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = common.TrueValue
	} else {
		fstat.Value["directory"] = common.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...common.Object) (common.Object, error) {
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
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(args ...common.Object) (common.Object, error) {
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
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...common.Object) (common.Object, error) {
	if len(args) != 3 {
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
	i2, ok := common.ToInt(args[1])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := common.ToInt(args[2])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(args ...common.Object) (common.Object, error) {
	if len(args) != 0 {
		return nil, common.ErrWrongNumArguments
	}
	arr := &common.Array{}
	for _, osArg := range os.Args {
		if len(osArg) > common.MaxStringLen {
			return nil, common.ErrStringLimit
		}
		arr.Value = append(arr.Value, &common.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *common.UserFunction {
	return &common.UserFunction{
		Name: name,
		Value: func(args ...common.Object) (common.Object, error) {
			if len(args) != 2 {
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
			i2, ok := common.ToInt64(args[1])
			if !ok {
				return nil, common.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(args ...common.Object) (common.Object, error) {
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
	res, ok := os.LookupEnv(s1)
	if !ok {
		return common.FalseValue, nil
	}
	if len(res) > common.MaxStringLen {
		return nil, common.ErrStringLimit
	}
	return &common.String{Value: res}, nil
}

func osExpandEnv(args ...common.Object) (common.Object, error) {
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
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > common.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > common.MaxStringLen {
		return nil, common.ErrStringLimit
	}
	return &common.String{Value: s}, nil
}

func osExec(args ...common.Object) (common.Object, error) {
	if len(args) == 0 {
		return nil, common.ErrWrongNumArguments
	}
	name, ok := common.ToString(args[0])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := common.ToString(arg)
		if !ok {
			return nil, common.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(args ...common.Object) (common.Object, error) {
	if len(args) != 1 {
		return nil, common.ErrWrongNumArguments
	}
	i1, ok := common.ToInt(args[0])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(args ...common.Object) (common.Object, error) {
	if len(args) != 4 {
		return nil, common.ErrWrongNumArguments
	}
	name, ok := common.ToString(args[0])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *common.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *common.ImmutableArray:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, common.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := common.ToString(args[2])
	if !ok {
		return nil, common.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *common.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *common.ImmutableArray:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, common.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []common.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*common.String)
		if !ok {
			return nil, common.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}
