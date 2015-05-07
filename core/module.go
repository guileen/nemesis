package core

type Context struct {
	Req *Req
	Res *Res
}

type Req struct {
	Context *Context
}

type Res struct {
	Context *Context
}

// Relations between module is tree like
type Module interface {
	Init(interface{})
	GetSubModules() []*Module
	// true to break, false to next
	Process(*Context) bool
}

func LoadModules(conf interface{}) GroupModule {
}

type SimpleModule struct {
	Modules []*Module
}

func (sm *SimpleModule) GetSubModules() []*Module {
	return sm.Modules
}

func (sm *SimpleModule) ProcessSubModuels(ctx *Context) bool {
	for _, m := range sm.GetSubModules() {
		if m.Process(ctx) {
			return true
		}
	}
	return false
}
