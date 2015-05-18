package modules

type SetUrlModule struct {
	BaseModule
}

func NewSetUrlModule(node Scalar) *SetUrlModule {
	return &SetUrlModule{}
}

func (m *SetUrlModule) Process(req *Req, res *Res) bool {
	// TODO req.SetPath()
	return m.Next().Process(req, res)
}
