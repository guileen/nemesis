package modules

type ProxyPassModule struct {
	BaseModule
}

func NewProxyPassModule(node Node) *ProxyPassModule {
	m := &ProxyPassModule{}
	return m
}

func (m *ProxyPassModule) Process(req *Req, res *Res) bool {
	return false
}
