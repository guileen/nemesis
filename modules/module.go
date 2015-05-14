package modules

// Relations between module is tree like
type Module interface {
	// Init(Node)
	// true to break, false to next
	SetNext(Module)
	Next() Module
	Process(*Req, *Res) bool
}

type BaseModule struct {
	next Module
}

func (base *BaseModule) SetNext(next Module) {
	base.next = next
}

func (base *BaseModule) Next() Module {
	return base.next
}

type SimpleModule struct {
	Modules []Module
}

func (sm *SimpleModule) GetSubModules() []Module {
	return sm.Modules
}

func (sm *SimpleModule) ProcessSubModuels(req *Req, res *Res) bool {
	for _, m := range sm.GetSubModules() {
		if m.Process(req, res) {
			return true
		}
	}
	return false
}
