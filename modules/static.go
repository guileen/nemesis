package modules

import "net/http"

type StaticModule struct {
	BaseModule
	rootPath   string
	fileServer http.Handler
}

func NewStaticModule(node Node) *StaticModule {
	m := &StaticModule{}
	m.Init(node)
	return m
}

func (m *StaticModule) Init(node Node) {
	m.rootPath = node.(Scalar).String()
	m.fileServer = http.FileServer(http.Dir(m.rootPath))
}

func (m *StaticModule) Process(req *Req, res *Res) bool {
	m.fileServer.ServeHTTP(res.writer, req.request)
	return true
}
