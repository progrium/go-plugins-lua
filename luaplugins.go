package luaplugins

import (
	"github.com/aarzilli/golua/lua"
	"github.com/stevedonovan/luar"
)


type Runtime struct {
	plugins map[string]*lua.State
}

func GetRuntime() *Runtime {
	return &Runtime{
		plugins: make(map[string]*lua.State),
	}
}

func (r Runtime) FileExtension() string {
	return ".lua"
}

func (r Runtime) InitPlugin(name, source string, implements func(string)) error {
	context := luar.Init()
	luar.Register(context, "", luar.Map{
		"implements": func(interfaceName string) {
			implements(interfaceName)	
		},
	})
	context.DoString(source)
	r.plugins[name] = context
	return nil
}

func (r Runtime) CallPlugin(name, function string, args []interface{}) (interface{}, error) {
	luafn := luar.NewLuaObjectFromName(r.plugins[name], function)
	value, err := luafn.Call(args...)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (r Runtime) LoadEnvironment(environment interface{}) {

}
