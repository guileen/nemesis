package modules

// Relations between module is tree like
type Module interface {
	// Init(Node)
	// true to break, false to next
	SetName(string)
	GetName() string
	SetNext(Module)
	Next() Module
	// prcess request and response, return true if processed
	Process(*Req, *Res) bool
}

type BaseModule struct {
	name string
	next Module
}

func (base *BaseModule) SetName(name string) {
	base.name = name
}

func (base *BaseModule) GetName() string {
	return base.name
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
