package gou

import (
	"strconv"
	"strings"

	"github.com/yaoapp/gou/session"
	"github.com/yaoapp/kun/exception"
)

// NewProcess 创建运行器
func NewProcess(name string, args ...interface{}) *Process {
	process := &Process{Name: name, Args: args, Global: map[string]interface{}{}}
	process.extraProcess()
	return process
}

// Run 运行方法
func (process *Process) Run() interface{} {
	return process.Handler(process)
}

// ProcessOf 创建运行器, 创建成功返回处理器，创建失败返回错误信息
func ProcessOf(name string, args ...interface{}) (*Process, error) {
	process := &Process{Name: name, Args: args, Global: map[string]interface{}{}}
	err := process.make()
	if err != nil {
		return nil, err
	}
	return process, nil
}

// Exec 运行方法, 失败返回错误
func (process *Process) Exec() (value interface{}, err error) {
	defer func() { err = exception.Catch(recover()) }()
	value = process.Handler(process)
	return value, nil
}

// RegisterProcessHandler 注册 ProcessHandler
func RegisterProcessHandler(name string, handler ProcessHandler) {
	name = strings.ToLower(name)
	ThirdHandlers[name] = handler
}

// AliasProcess 设置别名
func AliasProcess(name string, alias string) {
	name = strings.ToLower(name)
	alias = strings.ToLower(alias)
	ThirdHandlers[alias] = ThirdHandlers[name]
}

// WithSID 设定会话ID
func (process *Process) WithSID(sid string) *Process {
	process.Sid = sid
	return process
}

// WithGlobal 设定全局变量
func (process *Process) WithGlobal(global map[string]interface{}) *Process {
	process.Global = global
	return process
}

// 解析方法
func (process *Process) make() (err error) {
	defer func() { err = exception.Catch(recover()) }()
	process.extraProcess()
	return nil
}

// extraProcess 解析执行方法  name = "models.user.Find", name = "plugins.user.Login"
// return type=models, name=login, class=user
// @下一版优化这个函数
func (process *Process) extraProcess() {
	namer := strings.Split(process.Name, ".")
	last := len(namer) - 1
	if last < 2 && namer[0] != "flows" && namer[0] != "session" {
		exception.New("Process:%s format error", 400, process.Name).Throw()
	}

	process.Type = strings.ToLower(namer[0])
	if last > 1 {
		process.Class = strings.ToLower(strings.Join(namer[1:last], "."))
		process.Method = strings.ToLower(namer[last])
	} else {
		process.Class = strings.ToLower(namer[1])
		process.Method = ""
	}

	switch process.Type {
	case "plugins":
		process.Name = strings.ToLower(process.Name)
		process.Handler = processPlugin
		return
	case "flows":
		process.Name = strings.ToLower(process.Name)
		process.Handler = processFlow
		return
	case "scripts":
		process.Class = strings.ToLower(strings.Join(namer[1:last], "."))
		process.Method = namer[last]
		process.Handler = processScript
		return
	case "session":
		process.Method = strings.ToLower(namer[last])
		process.Handler = processSession
		return
	case "stores":
		process.Name = strings.ToLower(process.Name)
		handler, has := StoreHandlers[process.Method]
		if !has {
			exception.New("Store: %s %s does not exist", 404, process.Name, process.Method).Throw()
		}
		process.Handler = handler
		return
	case "models":
		process.Name = strings.ToLower(process.Name)
		handler, has := ModelHandlers[process.Method]
		if !has {
			exception.New("Model: %s %s does not exist", 404, process.Name, process.Method).Throw()
		}
		process.Handler = handler
		return
	default:
		if handler, has := ThirdHandlers[strings.ToLower(process.Name)]; has {
			process.Name = strings.ToLower(process.Name)
			process.Handler = handler
			return
		} else if handler, has := ThirdHandlers[process.Type]; has {
			process.Name = strings.ToLower(process.Name)
			process.Handler = handler
			return
		}
	}

	exception.New("%s does not found", 404, process.Name).Throw()
}

// processPlugin 运行插件中的方法
func processPlugin(process *Process) interface{} {
	plugin := SelectPluginModel(process.Class)
	res, err := plugin.Exec(process.Method, process.Args...)
	if err != nil {
		exception.Err(err, 500).Throw()
	}
	return res.MustValue()
}

// processFlow 运行工作流
func processFlow(process *Process) interface{} {
	name := strings.TrimPrefix(process.Name, "flows.")
	flow := SelectFlow(name).WithGlobal(process.Global).WithSID(process.Sid)
	return flow.Exec(process.Args...)
}

// processScript 运行脚本中定义的处理器
func processScript(process *Process) interface{} {
	res, err := Yao.New(process.Class, process.Method).
		WithGlobal(process.Global).
		WithSid(process.Sid).
		Call(process.Args...)

	if err != nil {
		message := err.Error()

		// JS Exception
		if strings.HasPrefix(message, "Exception|") {
			message = strings.Replace(message, "Exception|", "", -1)
			values := strings.Split(message, ":")
			if len(values) == 2 {
				code := 500
				if v, err := strconv.Atoi(values[0]); err == nil {
					code = v
				}
				message = strings.TrimSpace(values[1])
				exception.New(message, code).Throw()
			}
		}

		// Other
		code := 500
		values := strings.Split(message, "|")
		if len(values) == 2 {
			if v, err := strconv.Atoi(values[0]); err == nil {
				code = v
			}
			message = values[0]
		}

		exception.New(message, code).Throw()
	}
	return res
}

// processSession 运行Session函数
func processSession(process *Process) interface{} {

	if process.Method == "start" {
		process.Sid = session.ID()
		return process.Sid
	}

	if process.Sid == "" {
		return nil
	}

	ss := session.Global().ID(process.Sid)
	switch process.Method {
	case "id":
		return process.Sid
	case "get":
		process.ValidateArgNums(1)
		return ss.MustGet(process.ArgsString(0))
	case "set":
		process.ValidateArgNums(2)
		ss.MustSet(process.ArgsString(0), process.Args[1])
		return nil
	case "setmany":
		process.ValidateArgNums(1)
		ss.MustSetMany(process.ArgsMap(0))
		return nil
	case "dump":
		return ss.MustDump()
	}
	return nil
}
