package modules

type StaticModule struct {
	BaseModule
	rootPath string
}

func NewStaticModule(node Node) *StaticModule {
	m := &StaticModule{}
	m.Init(node)
	return m
}

func (m *StaticModule) Init(node Node) {
}

func (m *StaticModule) Process(req *Req, res *Res) bool {
	return false
}
