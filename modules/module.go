package modules

// Relations between module is tree like
type Module interface {
	Init(Node)
	// true to break, false to next
	Process(*Req, *Res) bool
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
